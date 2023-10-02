// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd/build/toolchain"
	"github.com/primecitizens/pcz/cmd/dev"
	"github.com/primecitizens/pcz/cmd/dev/internal"
	"github.com/primecitizens/pcz/cmd/dev/internal/js"
)

func init() {
	dev.AddEnvironment(&_cmd)
}

var (
	_serveOpts = ServeOptions{}
	_cmd       = cli.Cmd{
		Pattern:    "web [options] PACKAGE-DIR",
		BriefUsage: "Development help for web applications (js/wasm)",
		LocalFlags: cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &_serveOpts),
		Run:        runServe,
		Extra: &cli.CmdHelp{
			Example: "Serve the js/wasm app at :8080 and rebuild on file changes:\n\n" +
				"  $ pcz dev web --port 8080 ./main-pkg\n\n" +
				"Serve custom wasm blob and bindings:\n\n" +
				"  $ pcz dev web --port 8080 --wasm ./app.wasm --bindings ./app.js",
			Extra: &_serveOpts,
		},
		Children: []*cli.Cmd{
			js.BindgenCmd(),
		},
	}
)

type ServeOptions struct {
	WASM     string `cli:"wasm,#set path to a wasm blob (setting this disables automatic rebuild)"`
	Bindings string `cli:"bindings,#set path to a custom js bindings file (required when --wasm is set)"`
	HTML     string `cli:"html,#set path to a custom html page template to load bindings and wasm"`

	// bindgen options

	ES         string `cli:"es,def=5,comp=3,comp=5,comp=6,#set ES language target"`
	ModuleName string `cli:"name,def=bindings,#set module name"`

	// build options

	Debug       bool               `cli:"debug,#debug build process"`
	NoCache     bool               `cli:"no-cache,#build without cache"`
	Trimpath    bool               `cli:"trimpath,#remove all file system paths from the resulting executable"`
	Platform    toolchain.Platform `cli:"p|platform,def=js/wasm,#\"os/arch\" pair"`
	BuildDir    string             `cli:"build-dir,def=.pcz/build,#set directory to store intermediate build outputs"`
	EntrySymbol string             `cli:"entry,def=rt0,#set entry symbol name"`
	Tags        []string           `cli:"t|tag,#go build tags"`
	Defs        []string           `cli:"X|define,#link-time string value definition of the form importpath.name=value"`
	ASMFlags    []string           `cli:"A|asmflag,#add flags to go tool asm"`
	GCFlags     []string           `cli:"G|gcflag,#add flags to go tool compile (not including gccgo)"`
	LDFlags     []string           `cli:"L|ldflag,#add flags to go tool link"`

	// serve options

	ManualRun bool     `cli:"manual,#disable running wasm automatically"`
	CacheDir  string   `cli:"cache-dir,def=.pcz/cache/dev-js,#set directory to store built wasm, generated html & bindings triggered by file change"`
	Port      uint16   `cli:"port,#set web server listening port"`
	Env       []string `cli:"env,#add environment variables for wasm app in key=value format"`
	TLSCert   string   `cli:"tls-cert,#set path to a tls cert (pem file)"`
	TLSKey    string   `cli:"tls-key,#set path to the key to the tls-cert (pem file)"`
}

func runServe(copts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) (err error) {
	opts := route.Target().Extra.(*cli.CmdHelp).Extra.(*ServeOptions)

	if len(opts.Bindings) == 0 {
		if len(opts.WASM) != 0 {
			return fmt.Errorf("A custom bindings file (js) is required for a custom wasm blob, as pcz cannot figure out how to generate bindings for the custom wasm blob")
		}

		if len(posArgs) != 1 {
			return fmt.Errorf("expecting exactly one main package as positional arg, got %d", len(posArgs))
		}
	}

	cacheDir, err := filepath.Abs(opts.CacheDir)
	if err != nil {
		panic(err)
	}

	var dw *internal.DepWatcher
	if len(opts.WASM) == 0 {
		x := js.BindgenCmd().Extra.(*js.BindgenFlags).BuildOpts
		x.Platform = opts.Platform
		x.BuildDir = opts.BuildDir
		x.EntrySymbol = opts.EntrySymbol
		x.Tags = opts.Tags
		x.Defs = opts.Defs
		x.ASMFlags = opts.ASMFlags
		x.GCFlags = opts.GCFlags
		x.LDFlags = opts.LDFlags
		x.Debug = opts.Debug
		x.NoCache = opts.NoCache
		x.Trimpath = opts.Trimpath

		tc := toolchain.NewToolchain(&x)

		pkgs, err := tc.List(posArgs[0])
		if err != nil {
			return fmt.Errorf("failed to list package dependencies of %q for the first time", posArgs[0])
		}

		dw = internal.NewDepWatch(tc, pkgs, posArgs[0])
	}

	stopSig := make(chan struct{})
	updateCh := make(chan resourceUpdate, 1)
	ctx := &serveContext{
		stderr: copts.PickStderr(os.Stderr),
		stop:   stopSig,
		update: updateCh,
	}

	l, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv6loopback,
		Port: int(opts.Port),
		Zone: "",
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		l.Close()

		close(stopSig)
		close(updateCh)
	}()

	tdCh := make(chan *TemplateData, 1)

	go ctx.routineHTML(opts.HTML, tdCh, cacheDir)
	go ctx.routineWASM(opts.WASM, dw, cacheDir, js.BindgenOptions{
		Mode:          "umd",
		ES:            opts.ES,
		ModuleName:    opts.ModuleName,
		CustomWrapper: "", // no customization
	})
	if len(opts.Bindings) != 0 {
		go ctx.routineCustomBindings(opts.Bindings)
	}

	tdCh <- &TemplateData{
		OS:         opts.Platform.OS(),
		ModuleName: opts.ModuleName,
		ManualRun:  opts.ManualRun,
		Args:       dashArgs,
		Environ:    opts.Env,
	}

	res := &serveResouces{}

	go func() {
		for upd := range updateCh {
			res.Update(upd)
			ctx.Stderr("%s updated\n", upd.Kind)
		}
	}()

	srv := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var data []byte

			switch path := r.URL.Path; path {
			case "/bindings.js":
				w.Header().Set("Content-Type", "application/javascript")
				data = res.Bindings()
			case "/app.wasm":
				w.Header().Set("Content-Type", "application/wasm")
				data = res.WASM()
			case "/":
				w.Header().Set("Content-Type", "text/html")
				data = res.HTML()
			default:
				http.Error(w, "path not found: "+path, http.StatusNotFound)
				return
			}

			w.Header().Add("Cache-Control", "no-cache")
			http.ServeContent(w, r, r.URL.Path, time.Now(), bytes.NewReader(data))
		}),
	}

	tls := len(opts.TLSCert) != 0
	schema := "http"
	if tls {
		schema = "https"
	}

	ctx.Stderr("serving at %s://%s\n", schema, l.Addr().String())

	if tls {
		err = srv.ServeTLS(l, opts.TLSCert, opts.TLSKey)
	} else {
		err = srv.Serve(l)
	}
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

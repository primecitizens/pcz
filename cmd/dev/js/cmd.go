// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd"
	"github.com/primecitizens/pcz/cmd/buildcfg"
	"github.com/primecitizens/pcz/cmd/dev"
)

func init() {
	dev.AddEnvironment(&_cmd)
}

var (
	_serveOpts = ServeOptions{}
	_cmd       = cli.Cmd{
		Pattern:    "js",
		BriefUsage: "Development tools for js/wasm",
		LocalFlags: cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &_serveOpts),
		Run:        runServe,
		Extra:      &_serveOpts,
		Children: []*cli.Cmd{
			&_bindgenCmd,
		},
	}
)

type ServeOptions struct {
	Wasm     string `cli:"wasm,#set path to the wasm blob, if set, automatic build is disabled"`
	Bindings string `cli:"bindings,#set path to a custom js bindings file"`
	HTML     string `cli:"html,#set path to a custom html page template to load bindings and wasm"`

	// bindgen options

	ES         string `cli:"es,def=5,comp=3,comp=5,comp=6,#set ES language target"`
	ModuleName string `cli:"name,def=bindings,#set es module name"`

	// build options

	Platform    buildcfg.Platform `cli:"p|platform,def=js/wasm,#\"os/arch\" pair"`
	BuildDir    string            `cli:"build-dir,def=.pcz/build,#set directory to store intermediate build outputs"`
	EntrySymbol string            `cli:"entry,def=rt0,#set entry symbol name"`
	Tags        []string          `cli:"t|tag,#go build tags"`
	Defs        []string          `cli:"X|define,#link-time string value definition of the form importpath.name=value"`
	ASMFlags    []string          `cli:"A|asmflag,#add flags to go tool asm"`
	GCFlags     []string          `cli:"G|gcflag,#add flags to go tool compile (not including gccgo)"`
	LDFlags     []string          `cli:"L|ldflag,#add flags to go tool link"`

	// serve options

	CacheDir string   `cli:"cache-dir,def=.pcz/cache,#set directory to store built wasm, generated html & bindings triggered by file change"`
	Port     uint16   `cli:"port,#set port to listen"`
	Env      []string `cli:"env,#add environment variables in key=value format"`
	TLSCert  string   `cli:"tls-cert,#set path to a tls cert (pem file)"`
	TLSKey   string   `cli:"tls-key,#set path to the key to the tls-cert (pem file)"`
}

func runServe(copts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) (err error) {
	opts := route.Target().Extra.(*ServeOptions)

	if len(posArgs) != 1 {
		return fmt.Errorf("expecting exactly one main package as positional arg, got %d", len(posArgs))
	}

	if len(opts.Bindings) == 0 {
		err = createBindings(io.Discard, nil, bindgenSpec{
			mode: "raw",
		})
		if err != nil {
			_, ok := err.(*cmd.ErrFeatureNotBuilt)
			if ok {
				return fmt.Errorf("pcz was not built with bindgen feature, thus a custom bindings file (js) is required: %w", err)
			}
		}
	}

	cacheDir, err := filepath.Abs(opts.CacheDir)
	if err != nil {
		panic(err)
	}

	var dw *cmd.DepWatcher
	if len(opts.Bindings) == 0 || len(opts.Wasm) == 0 {
		x := &_bindgenOpts.buildOpts
		x.Platform = opts.Platform
		x.BuildDir = opts.BuildDir
		x.EntrySymbol = opts.EntrySymbol
		x.Tags = opts.Tags
		x.Defs = opts.Defs
		x.ASMFlags = opts.ASMFlags
		x.GCFlags = opts.GCFlags
		x.LDFlags = opts.LDFlags

		tc := buildcfg.NewToolchain(x)

		pkgs, err := tc.List(posArgs[0])
		if err != nil {
			return fmt.Errorf("failed to list package dependencies of %q for the first time", posArgs[0])
		}

		dw = cmd.NewDepWatch(tc, pkgs, posArgs[0])
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
	go ctx.routineWASM(opts.Wasm, dw, cacheDir)
	go ctx.routineBindings(opts.Bindings, dw, cacheDir, bindgenSpec{
		mode:       "umd",
		es:         opts.ES,
		moduleName: opts.ModuleName,
		template:   "", // no customization
	})

	tdCh <- &TemplateData{
		OS:         opts.Platform.OS(),
		ModuleName: opts.ModuleName,
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

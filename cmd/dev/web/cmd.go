// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"bytes"
	"errors"
	"fmt"
	"go/build"
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
	dev.AddEnvironment(&webCmd)
}

var (
	serveFlags = ServeFlags{}
	webCmd     = cli.Cmd{
		Pattern:    "web [options] PACKAGE-DIR",
		BriefUsage: "Development help for web applications (js/wasm)",
		LocalFlags: cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &serveFlags),
		Run:        runServe,
		Extra: &cli.CmdHelp{
			Example: "Serve the app at :8080 and rebuild on source code update:\n\n" +
				"  $ pcz dev web --port 8080 ./main-pkg\n\n" +
				"Serve custom wasm blob and bindings:\n\n" +
				"  $ pcz dev web --port 8080 --wasm ./app.wasm --bindings ./app.js",
			Extra: &serveFlags,
		},
		Children: []*cli.Cmd{
			js.BindgenCmd(),
		},
	}
)

type ServeFlags struct {
	WASM     string `cli:"wasm,#set path to a wasm blob (setting this disables automatic rebuild)"`
	Bindings string `cli:"bindings,#set path to a custom js bindings file (required when --wasm is set)"`
	HTML     string `cli:"html,#set path to a custom html page template to load bindings and wasm"`

	// serve options

	ManualRun bool     `cli:"manual,#disable running wasm automatically"`
	CacheDir  string   `cli:"cache-dir,def=.pcz/cache/dev-js,#set directory to store built wasm, generated html & bindings triggered by file change"`
	Port      uint16   `cli:"port,#set web server listening port"`
	Env       []string `cli:"env,#add environment variables for wasm app in key=value format"`
	TLSCert   string   `cli:"tls-cert,#set path to a tls cert (pem file)"`
	TLSKey    string   `cli:"tls-key,#set path to the key to the tls-cert (pem file)"`

	// bindgen options

	ES         string `cli:"es,def=5,comp=5,comp=6,#set ES language when transpiling typescript"`
	ModuleName string `cli:"module-name,def=bindings,#set module name for generated bindings"`

	// build options

	Debug    bool   `cli:"debug,#debug the wasm build process"`
	BuildDir string `cli:"build-dir,def=.pcz/build,#set directory to store intermediate build outputs"`
	NoCache  bool   `cli:"no-cache,#build without cache"`
	Trimpath bool   `cli:"trimpath,#remove all file system paths from the resulting executable"`

	Defs        []string           `cli:"X|define,#link-time string value definition of the form importpath.name=value"`
	EntrySymbol string             `cli:"entry,def=rt0,#set entry symbol name"`
	Platform    toolchain.Platform `cli:"p|platform,def=js/wasm,#\"os/arch\" pair"`
	Tags        []string           `cli:"t|tag,#add build tags"`

	LDFlags  []string `cli:"L|ldflag,#add flags to go tool link"`
	GCFlags  []string `cli:"G|gcflag,#add flags to go tool compile (not including gccgo)"`
	ASMFlags []string `cli:"A|asmflag,#add flags to go tool asm"`
}

func runServe(opts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) (err error) {
	flags := &serveFlags

	if len(flags.Bindings) == 0 {
		if len(flags.WASM) != 0 {
			return fmt.Errorf(
				"Please add `--bindings <path>` for `--wasm %q`: "+
					"pcz cannot generate bindings for custom wasm blob",
				flags.WASM,
			)
		}

		if len(posArgs) != 1 {
			return fmt.Errorf(
				"Please set exactly one package dir as positional arg, got %d", len(posArgs),
			)
		}
	}

	cacheDir, err := filepath.Abs(flags.CacheDir)
	if err != nil {
		panic(err)
	}

	var dw *internal.DepWatcher
	if len(flags.WASM) == 0 {
		tc := toolchain.NewToolchain(&toolchain.Options{
			Context: &build.Context{
				GOARCH:        "wasm",
				GOOS:          "js",
				GOROOT:        toolchain.Context.GOROOT,
				GOPATH:        toolchain.Context.GOPATH,
				Dir:           toolchain.Context.Dir,
				CgoEnabled:    toolchain.Context.CgoEnabled,
				UseAllFiles:   toolchain.Context.UseAllFiles,
				Compiler:      toolchain.Context.Compiler,
				BuildTags:     toolchain.Context.BuildTags,
				ToolTags:      toolchain.Context.ToolTags,
				ReleaseTags:   toolchain.Context.ReleaseTags,
				InstallSuffix: toolchain.Context.InstallSuffix,
			},
			Debug:       flags.Debug,
			BuildDir:    flags.BuildDir,
			NoCache:     flags.NoCache,
			Trimpath:    flags.Trimpath,
			Defs:        flags.Defs,
			EntrySymbol: flags.EntrySymbol,
			Platform:    flags.Platform,
			Tags:        flags.Tags,
			LDFlags:     flags.LDFlags,
			GCFlags:     flags.GCFlags,
			ASMFlags:    flags.ASMFlags,
		})

		pkgs, err := tc.List(posArgs[0])
		if err != nil {
			return fmt.Errorf("failed to list package dependencies of %q for the first time", posArgs[0])
		}

		dw = internal.NewDepWatch(tc, pkgs, posArgs[0])
	}

	stopSig := make(chan struct{})
	updateCh := make(chan resourceUpdate, 1)
	ctx := &serveContext{
		stderr: opts.PickStderr(os.Stderr),
		stop:   stopSig,
		update: updateCh,
	}

	l, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv6loopback,
		Port: int(flags.Port),
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

	go ctx.routineHTML(flags.HTML, tdCh, cacheDir)
	go ctx.routineWASM(flags.WASM, dw, cacheDir, js.BindgenSpec{
		ModuleSystem:  "umd",
		ES:            flags.ES,
		ModuleName:    flags.ModuleName,
		CustomWrapper: "", // no customization
	})
	if len(flags.Bindings) != 0 {
		go ctx.routineCustomBindings(flags.Bindings)
	}

	tdCh <- &TemplateData{
		OS:         flags.Platform.OS(),
		ModuleName: flags.ModuleName,
		ManualRun:  flags.ManualRun,
		Args:       dashArgs,
		Environ:    flags.Env,
	}

	res := &serveResouces{}

	go func() {
		for upd := range updateCh {
			res.Update(upd)
			ctx.Stderr("[%s] updated\n", upd.Kind)
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

	tls := len(flags.TLSCert) != 0
	schema := "http"
	if tls {
		schema = "https"
	}

	ctx.Stderr("serving at %s://%s\n", schema, l.Addr().String())

	if tls {
		err = srv.ServeTLS(l, flags.TLSCert, flags.TLSKey)
	} else {
		err = srv.Serve(l)
	}
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

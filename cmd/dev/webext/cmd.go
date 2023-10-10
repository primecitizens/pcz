// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webext

import (
	"bytes"
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"slices"
	"sync"
	"time"
	"unsafe"

	"github.com/fsnotify/fsnotify"
	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd/build/toolchain"
	"github.com/primecitizens/pcz/cmd/dev"
	"github.com/primecitizens/pcz/cmd/dev/internal"
	"github.com/primecitizens/pcz/cmd/dev/internal/js"
	"github.com/primecitizens/pcz/cmd/embed"
)

func init() {
	dev.AddEnvironment(&webextCmd)
}

var (
	webextFlags = WebextFlags{
		PackMatch: defaultPackMatch,
	}
	webextCmd = cli.Cmd{
		Pattern:    "webext [options] PACKAGE-DIR",
		BriefUsage: "Dev environment for WebExtension development (js/wasm)",
		LocalFlags: cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &webextFlags),
		Run:        runWebext,
		Extra: &cli.CmdHelp{
			Example: "continuous wasm build and bindgen generation on source file change:\n\n" +
				"  $ pcz dev webext --bindings ./bindngs.js --wasm ./main.wasm ./main-pkg",
		},
		Children: []*cli.Cmd{
			&cmdNew,
			&cmdPack,
			js.BindgenCmd(),
		},
	}
)

type WebextFlags struct {
	WASM      string `cli:"wasm,#set output path for the built wasm"`
	Bindings  string `cli:"bindings,#set output path for the generated bindings"`
	Embedding string `cli:"embedding,#set output path for the generated js embedding"`
	Zip       string `cli:"zip,#set output path for the generated zip file"`

	// embed flags

	EmbedMode  string            `cli:"embed-mode,def=uint8array,comp=uint8array,comp=base64,comp=data-url,#set embedding mode"`
	EmbedFiles map[string]string `cli:"embed,#add embedded file in format \"[export] {var|const|let} <var-name>=<file-path>\""`

	// bindgen flags

	ES         string `cli:"es,def=5,comp=5,comp=6,#set ES language when transpiling typescript"`
	ModuleName string `cli:"module-name,def=bindings,#set module name for generated bindings"`

	// pack flags

	PackDir   string   `cli:"pack-base,#set base dir for matching pack glob patterns"`
	PackMatch []string `cli:"pack,#add glob pattern to include matched files in zip"`

	// build flags

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

func runWebext(opts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) error {
	if len(posArgs) != 1 {
		return fmt.Errorf("expecting exactly one main package as positional arg, got %d", len(posArgs))
	}

	flags := &webextFlags
	if len(flags.Bindings) == 0 {
		return fmt.Errorf("please add `--bindings <path>` to set path to write generated bindings.")
	}

	if len(flags.WASM) == 0 {
		return fmt.Errorf("please add `--wasm <path>` to set path to write the built wasm bolb.")
	}

	if len(flags.EmbedFiles) != 0 && len(flags.Embedding) == 0 {
		return fmt.Errorf("please add `--embedding <path>` to set path to write your embedded files")
	}

	if len(flags.PackMatch) != 0 && len(flags.Zip) == 0 {
		return fmt.Errorf("please add `--zip <path>` to set path to write zip file of matched files")
	}

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

	w, err := fsnotify.NewWatcher()
	if err != nil {
		panic(fmt.Errorf("failed to create fs watcher: %w", err))
	}

	dw := internal.NewDepWatch(tc, pkgs, posArgs[0])

	err = dw.AddWatcher(w)
	if err != nil {
		panic(fmt.Errorf("failed to register fs watcher to watch source code: %w", err))
	}

	_stderr := opts.PickStderr(os.Stderr)
	Stderr := func(format string, args ...any) {
		fmt.Fprintf(_stderr, format, args...)
	}

	if len(flags.Embedding) == 0 && len(flags.EmbedFiles) != 0 {
		Stderr("WARNING: embedding output path is set, but no files to embed: missing `--embed` flags?\n")
	}

	if len(flags.Zip) == 0 && len(flags.PackMatch) != 0 {
		Stderr("WARNING: zip output path is set, but no glob pattern is provided: missing `--pack` flags?\n")
	}

	var (
		packTsk      *internal.SingleTask
		embeddingTsk *internal.SingleTask
	)

	if globs := flags.PackMatch; len(globs) != 0 {
		packTsk = internal.NewSingleTask(time.Second, func() {
			Stderr("[zip] updating...\n")
			var buf bytes.Buffer

			entries, err := Pack(&buf, flags.PackDir, globs)
			if err != nil {
				Stderr("[zip] failed to create zip content: %v\n", err)
				return
			}

			if len(entries) != 0 {
				Stderr("[zip] packed:\n")
				for _, f := range entries {
					Stderr("[zip-archive] path = %q, size = %d\n", f.Path, f.Size)
				}
			} else {
				Stderr("[zip] no files packed.\n")
			}

			err = os.WriteFile(flags.Zip, buf.Bytes(), 0640)
			if err != nil {
				Stderr("[zip] failed to write to file %q: %v\n", flags.Zip, err)
				return
			}

			Stderr("[zip] updated\n")
		})
	}

	if len(flags.EmbedFiles) != 0 {
		embedWatch, err := fsnotify.NewWatcher()
		if err != nil {
			panic(fmt.Errorf("[embedding] failed to create fs watcher for embeded files: %w", err))
		}
		defer embedWatch.Close()

		missingFiles := &sync.Map{}

		watchFiles := make(map[string]string, len(flags.EmbedFiles))
		for name, file := range flags.EmbedFiles {
			absPath, err := filepath.Abs(file)
			if err != nil {
				panic(fmt.Errorf("[embedding] failed to get absoulte path of %q", file))
			}

			err = embedWatch.Add(absPath)
			if err != nil {
				if os.IsNotExist(err) {
					// record missing
					missingFiles.Store(absPath, nil)
				} else {
					panic(fmt.Errorf("[embedding] failed to watch embed path %q: %w", absPath, err))
				}
			}

			watchFiles[name] = absPath
		}

		embeddingTsk = internal.NewSingleTask(time.Second, func() {
			skip := false
			missingFiles.Range(func(key, value any) bool {
				absPath := key.(string)
				if _, err := os.Stat(absPath); err != nil {
					skip = true
					Stderr("[embedding] missing file %q: %v\n", absPath, err)
					return true
				}

				if err = embedWatch.Add(absPath); err != nil {
					// failed to add, but the file does exit, do not skip
					Stderr("[embedding] failed adding %q to file watch: %v", absPath, err)
				} else {
					// added to file watcher, no longer needed
					missingFiles.Delete(absPath)
				}

				return true
			})

			if skip {
				return
			}

			Stderr("[embedding] updating...\n")
			var buf bytes.Buffer

			err := embed.EmbedJS(&buf, embed.JSEmbedOptions{
				Mode:  flags.EmbedMode,
				Files: watchFiles,
			})
			if err != nil {
				Stderr("[embedding] failed to create js embedding: %v\n", err)
				return
			}

			err = os.WriteFile(flags.Embedding, buf.Bytes(), 0640)
			if err != nil {
				Stderr("[embedding] failed to write to file %q: %v\n", flags.Embedding, err)
				return
			}

			Stderr("[embedding] updated\n")
			if packTsk != nil {
				go packTsk.Run()
			}
		})

		go func() {
			embeddingTsk.Run()
			select {
			case evt := <-embedWatch.Events:
				switch evt.Op {
				case fsnotify.Create, fsnotify.Write:
					if flags.Debug {
						Stderr("[debug] request js embedding generation for %q\n", evt.Name)
					}
					go embeddingTsk.Run()
				}
			case err = <-embedWatch.Errors:
				Stderr("[fswatch] error: %v\n", err)
			}
		}()
	}

	bindingsTsk := internal.NewSingleTask(time.Second, func() {
		Stderr("[bindings] updating...\n")
		data, err := os.ReadFile(flags.WASM)
		if err != nil {
			Stderr("[bindings] failed to read the built wasm: %v\n", err)
			return
		}

		filter, err := js.CreateSimpleFilterFromWasm(data)
		if err != nil {
			Stderr("[bindings] failed to collect imported functions from the built wasm: %v\n", err)
			return
		}

		code, err := js.CreateJSBindings(js.BindgenSpec{
			ModuleSystem: "umd",
			ES:           flags.ES,
			ModuleName:   flags.ModuleName,
			Deps:         dw.CurrentDeps(),
			Filter:       filter,
		})
		if err != nil {
			Stderr("[bindings] failed to generate: %v\n", err)
			return
		}

		data = unsafe.Slice(unsafe.StringData(code), len(code))
		err = os.WriteFile(flags.Bindings, data, 0640)
		if err != nil {
			Stderr("[bindings] failed to write to file %q: %v\n", flags.Bindings, err)
			return
		}

		Stderr("[bindings] updated\n")
		if embeddingTsk != nil {
			go embeddingTsk.Run()
		}
	})

	wasmTsk := internal.NewSingleTask(time.Second, func() {
		Stderr("[wasm] building...\n")
		err = dw.Build(flags.WASM)
		if err != nil {
			Stderr("[wasm] failed to build: %v\n", err)
			return
		}

		Stderr("[wasm] updated\n")

		bindingsTsk.Run()

		if packTsk != nil {
			go packTsk.Run()
		}
	})

	wasmTsk.Run()
	for {
		select {
		case ev := <-w.Events:
			switch ev.Op {
			case fsnotify.Remove, fsnotify.Write:
			default:
				continue
			}

			pkg := dw.Contains(filepath.Dir(ev.Name))
			if pkg == nil {
				continue
			}

			switch file := filepath.Base(ev.Name); {
			case file == "ffi_bindings.ts":
				if flags.Debug {
					Stderr("[debug] request bindings generation for %q", ev.Name)
				}
				go bindingsTsk.Run()
				continue
			case slices.Contains(pkg.GoFiles, file):
				// TODO: diff package deps in dw?
				// may have changed import statements
				dw.Refresh()
			case slices.Contains(pkg.SFiles, file),
				slices.Contains(pkg.EmbedFiles, file):
			default:
				continue
			}

			if flags.Debug {
				Stderr("[debug] request wasm rebuild for %q", ev.Name)
			}
			go wasmTsk.Run()
		case err = <-w.Errors:
			Stderr("[fswatch] error: %v\n", err)
		}
	}

	return nil
}

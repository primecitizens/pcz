// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webext

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"time"
	"unsafe"

	"github.com/fsnotify/fsnotify"
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
	_cmd = cli.Cmd{
		Pattern:    "webext [options] PACKAGE-DIR",
		BriefUsage: "Dev environment for WebExtension development (js/wasm)",
		LocalFlags: _Flags{},
		Run:        runServe,
		Extra: &cli.CmdHelp{
			Example: "continuous wasm rebuild and bindgen generation on source file change:\n\n" +
				"  $ pcz dev webext --bindings ./bindngs.js --wasm ./main.wasm ./main-pkg",
		},
		Children: []*cli.Cmd{
			js.BindgenCmd(),
		},
	}

	_baseFlags           = cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, js.GetBindgenFlags())
	_flag_OutputBindings = cli.String{
		BriefUsage: "set output path for the generated bindings",
		Value:      &js.GetBindgenFlags().OutputFile,
	}
	_flag_OutputWasm = cli.String{
		BriefUsage: "set output path for the built wasm",
		Value:      &js.GetBindgenFlags().WASM,
	}
)

var _ cli.FlagIndexer = _Flags{}

type _Flags struct{}

// FindFlag implements cli.FlagIndexer.
func (_Flags) FindFlag(name string) (cli.Flag, bool) {
	switch name {
	case "bindings":
		return &_flag_OutputBindings, true
	case "wasm":
		return &_flag_OutputWasm, true
	case "output", "o": // prevent using BindgenFlags.OutputFile
		return nil, false
	default:
		return _baseFlags.FindFlag(name)
	}
}

// NthFlag implements cli.FlagIndexer.
func (_Flags) NthFlag(i int) (cli.FlagInfo, bool) {
	switch i {
	case 0:
		return cli.FlagInfo{Name: "bindings"}, true
	case 1:
		return cli.FlagInfo{Name: "wasm"}, true
	default:
		// it skips the first 2 flags
		return _baseFlags.NthFlag(i)
	}
}

func runServe(opts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) error {
	if len(posArgs) != 1 {
		return fmt.Errorf("expecting exactly one main package as positional arg, got %d", len(posArgs))
	}

	flags := js.GetBindgenFlags()
	if len(flags.OutputFile) == 0 {
		return fmt.Errorf("please set path to write generated bindings.")
	}

	if len(flags.WASM) == 0 {
		return fmt.Errorf("please set path to write the wasm bolb.")
	}

	tc := toolchain.NewToolchain(&flags.BuildOpts)
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

	bindingsTsk := internal.NewSingleTask(time.Second, func() {
		Stderr("updating bindings...\n")
		data, err := os.ReadFile(flags.WASM)
		if err != nil {
			Stderr("failed to read the built wasm for bindings: %v\n", err)
			return
		}

		filter, err := js.CreateSimpleFilterFromWasm(data)
		if err != nil {
			Stderr("failed to collect imported functions from the built wasm: %v\n", err)
			return
		}

		code, err := js.CreateJSBindings(js.BindgenOptions{
			Mode:          "umd",
			ES:            flags.ES,
			ModuleName:    flags.ModuleName,
			CustomWrapper: flags.CustomTemplate,
			Deps:          dw.CurrentDeps(),
			Filter:        filter,
		})
		if err != nil {
			Stderr("failed to generate bindings: %v\n", err)
			return
		}

		data = unsafe.Slice(unsafe.StringData(code), len(code))
		_ = os.WriteFile(flags.OutputFile, data, 0640)
		Stderr("bindings updated.\n")
	})

	wasmTsk := internal.NewSingleTask(time.Second, func() {
		Stderr("building wasm blob...\n")
		err = dw.Build(flags.WASM)
		if err != nil {
			Stderr("failed to build wasm blob: %v\n", err)
			return
		}

		bindingsTsk.Run()
		Stderr("wasm updated.\n")
	})

	wasmTsk.Run()
	for {
		select {
		case ev := <-w.Events:
			switch ev.Op {
			case fsnotify.Remove,
				fsnotify.Write:
			default:
				continue
			}

			pkg := dw.Contains(filepath.Dir(ev.Name))
			if pkg == nil {
				continue
			}

			switch file := filepath.Base(ev.Name); {
			case file == "ffi_bindings.ts":
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

			go wasmTsk.Run()
		case <-w.Errors:
		}
	}

	return nil
}

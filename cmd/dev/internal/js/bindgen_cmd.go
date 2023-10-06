// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"fmt"
	"go/build"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd/build/toolchain"
	"github.com/primecitizens/pcz/cmd/utils/fsutils"
	"github.com/primecitizens/pcz/cmd/utils/jsutils"
)

func BindgenCmd() *cli.Cmd {
	return &_bindgenCmd
}

var (
	_bindgenOpts = BindgenFlags{
		BuildOpts: toolchain.Options{
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
			EntrySymbol: "rt0",
			Platform:    "js/wasm",
		},
		OutputFile:     "",
		CustomTemplate: "",
		Mode:           "umd",
		ES:             "5",
		ModuleName:     "bindings",
		Platform:       "js/wasm",
	}

	_bindgenCmd = cli.Cmd{
		Pattern:    "bindgen --wasm /path/to/myapp.wasm PACKAGE-DIR",
		BriefUsage: "Generate bindings by collecting ffi_bindings.ts in all dependent go packages",
		LocalFlags: cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &_bindgenOpts),
		Run:        runBindgen,
		Extra:      &_bindgenOpts,
	}
)

type BindgenFlags struct {
	BuildOpts toolchain.Options

	WASM           string             `cli:"wasm,#path to the wasm blob (required, for import filtering)"`
	OutputFile     string             `cli:"o|output,#set file path to write the generated output, if not set, write generated bindings to stdout"`
	Mode           string             `cli:"m|mode,def=raw,comp=raw,comp=cjs,comp=amd,comp=umd,#set target module system to transpile typescript"`
	Minify         bool               `cli:"minify,#write minified code and source map along with the bindgen output file"`
	Platform       toolchain.Platform `cli:"p|platform,def=js/wasm,#set os/arch pair"`
	Tags           []string           `cli:"t|tag,#go build tags"`
	CustomTemplate string             `cli:"T|template,#set path to a custom bindgen template file"`
	ES             string             `cli:"es,def=5,comp=3,comp=5,comp=6,#set ES language when transpiling typescript"`
	ModuleName     string             `cli:"n|name,def=bindings,#set module name for generated bindings"`
}

func runBindgen(opts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) (err error) {
	flags := route.Target().Extra.(*BindgenFlags)
	if len(posArgs) != 1 {
		panic(fmt.Errorf("expecting exactly one positional arg for main package (got %d)", len(posArgs)))
	}

	spec := BindgenOptions{
		Mode:          flags.Mode,
		ES:            flags.ES,
		ModuleName:    flags.ModuleName,
		CustomWrapper: "",  // see below
		Deps:          nil, // see below
		Filter:        nil, // see below
	}

	if tfile := flags.CustomTemplate; len(tfile) != 0 {
		data, err := os.ReadFile(tfile)
		if err != nil {
			panic(err)
		}

		spec.CustomWrapper = unsafe.String(unsafe.SliceData(data), len(data))
	}

	flags.BuildOpts.Platform = flags.Platform
	flags.BuildOpts.Tags = flags.Tags

	tc := toolchain.NewToolchain(&flags.BuildOpts)
	spec.Deps, err = tc.List(posArgs[0])
	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(flags.WASM)
	spec.Filter, err = CreateSimpleFilterFromWasm(file)
	if err != nil {
		panic(err)
	}

	code, err := CreateJSBindings(spec)
	if err != nil {
		panic(err)
	}

	outputFile := flags.OutputFile
	err = fsutils.WriteToFileOrStdout(opts, outputFile, func(w io.Writer) error {
		_, err := w.Write(unsafe.Slice(unsafe.StringData(code), len(code)))
		return err
	})
	if err != nil {
		panic(err)
	}

	if flags.Minify {
		var discardSourceMap bool
		switch spec.Mode {
		case "raw", "":
			discardSourceMap = true
			if len(spec.ModuleName) == 0 {
				spec.ModuleName = "bindings" // ensure there is a module name
			}

			spec.Mode = "umd"
			code, err = transpileTS(code, spec)
			if err != nil {
				panic(err)
			}
		}

		var minified, sourceMap []byte
		minified, sourceMap, err = jsutils.Minify(jsutils.MinifyOptions{
			SourceCode: code,
			ES:         translateES(spec.ES),
		})
		if err != nil {
			panic(err)
		}

		var minifiedFile, sourceMapFile string
		if len(outputFile) > 0 {
			name := filepath.Base(outputFile)
			i := strings.LastIndexByte(name, '.')
			if i >= 0 {
				outputFile = outputFile[:len(outputFile)-len(name)+i]
			}

			minifiedFile = outputFile + ".min.js"
			sourceMapFile = minifiedFile + ".map"
		}

		err = fsutils.WriteToFileOrStdout(opts, minifiedFile, func(w io.Writer) error {
			_, err := w.Write(minified)
			return err
		})
		if err != nil {
			panic(err)
		}

		if !discardSourceMap {
			err = fsutils.WriteToFileOrStdout(opts, sourceMapFile, func(w io.Writer) error {
				_, err = w.Write(sourceMap)
				return err
			})
			if err != nil {
				panic(err)
			}
		}
	}

	return nil
}

func translateES(es string) api.Target {
	switch es {
	case "next":
		return api.ESNext
	case "5":
		return api.ES5
	case "2015", "6":
		return api.ES2015
	case "2016":
		return api.ES2016
	case "2017":
		return api.ES2017
	case "2018":
		return api.ES2018
	case "2019":
		return api.ES2019
	case "2020":
		return api.ES2020
	case "2021":
		return api.ES2021
	case "2022":
		return api.ES2022
	default:
		return api.ESNext
	}
}

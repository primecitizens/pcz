// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"go/build"

	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd/buildcfg"
)

var (
	_bindgenOpts = BindgenOptions{
		buildOpts: buildcfg.BuildOptions{
			Context: &build.Context{
				GOARCH:        "wasm",
				GOOS:          "js",
				GOROOT:        buildcfg.Context.GOROOT,
				GOPATH:        buildcfg.Context.GOPATH,
				Dir:           buildcfg.Context.Dir,
				CgoEnabled:    buildcfg.Context.CgoEnabled,
				UseAllFiles:   buildcfg.Context.UseAllFiles,
				Compiler:      buildcfg.Context.Compiler,
				BuildTags:     buildcfg.Context.BuildTags,
				ToolTags:      buildcfg.Context.ToolTags,
				ReleaseTags:   buildcfg.Context.ReleaseTags,
				InstallSuffix: buildcfg.Context.InstallSuffix,
			},
			EntrySymbol: "rt0",
			Platform:    "js/wasm",
		},
		OutputFile:   "",
		TemplateFile: "",
		Mode:         "umd",
		ES:           "5",
		ModuleName:   "bindings",
		Platform:     "js/wasm",
	}

	_bindgenCmd = cli.Cmd{
		Pattern:    "bindgen",
		BriefUsage: "Generate bindings by collecting ffi_bindings.ts in all dependent go packages",
		LocalFlags: cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &_bindgenOpts),
		Run:        runBindgen,
		Extra:      &_bindgenOpts,
	}
)

type BindgenOptions struct {
	buildOpts buildcfg.BuildOptions

	OutputFile   string            `cli:"o|output,#set file path to write the generated output"`
	TemplateFile string            `cli:"T|template,#set path to custom bindgen template file"`
	Mode         string            `cli:"m|mode,def=raw,comp=raw,comp=cjs,comp=amd,comp=umd,#transpile typescript to"`
	ES           string            `cli:"es,def=5,comp=3,comp=5,comp=6,#set ES language target"`
	ModuleName   string            `cli:"n|name,def=bindings,#set es module name"`
	Tags         []string          `cli:"t|tag,#go build tags"`
	Platform     buildcfg.Platform `cli:"p|platform,def=js/wasm,#set os/arch pair"`
}

type bindgenSpec struct {
	mode       string
	es         string
	moduleName string
	template   string
}

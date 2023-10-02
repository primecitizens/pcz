// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package build

import (
	"fmt"
	"path/filepath"

	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd/build/toolchain"
)

func Cmd() *cli.Cmd {
	return &_build
}

var (
	_buildOpts = toolchain.Options{
		Context: &toolchain.Context,
	}
	_opts  = Options{}
	_build = cli.Cmd{
		Pattern:    "build PACKAGE",
		BriefUsage: "",
		Flags:      cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &_buildOpts),
		LocalFlags: cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &_opts),
		Run:        runBuild,
		Extra:      &_buildOpts,
	}
)

type Options struct {
	Output string `cli:"o|output,#write output to file"`
}

func runBuild(opts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) error {
	if len(posArgs) != 1 {
		return fmt.Errorf("expecting exactly one package to build, got %d", len(posArgs))
	}

	tc := toolchain.NewToolchain(route.Target().Extra.(*toolchain.Options))
	if err := tc.VerifyVersion(tc.GoVersion()); err != nil {
		panic(err)
	}

	output, err := filepath.Abs(_opts.Output)
	if err != nil {
		panic(err)
	}

	pkgs, err := tc.List(posArgs[0])
	if err != nil {
		panic(err)
	}

	err = tc.Build(pkgs, output)
	if err != nil {
		panic(err)
	}

	return nil
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !bindgen

package js

import (
	"io"
	"os"

	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd"
	"github.com/primecitizens/pcz/cmd/buildcfg"
)

var errBindgenNotBuilt = &cmd.ErrFeatureNotBuilt{
	Feature: "dev/wasm/bindgen",
	Tag:     "bindgen",
}

func runBindgen(opts *cli.CmdOptions, route cli.Route, posArgs, dashArgs []string) error {
	_, err2 := io.WriteString(opts.PickStderr(os.Stderr), errBindgenNotBuilt.Error())
	if err2 != nil {
		panic(err2)
	}

	return errBindgenNotBuilt
}

func createBindings(out io.Writer, pkgs []buildcfg.Package, spec bindgenSpec) error {
	return errBindgenNotBuilt
}

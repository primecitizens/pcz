// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webext

import (
	"fmt"

	"github.com/primecitizens/cli"
)

// TODO: implement tool for updating manifest.json

var (
	_cmdManifest = cli.Cmd{
		Pattern:    "manifest",
		BriefUsage: "",
		Flags:      nil,
		LocalFlags: nil,
		FlagRule:   nil,
		Run:        updateManifest,
		Completion: nil,
		Extra:      nil,
		Children:   []*cli.Cmd{},
		State:      0,
	}
)

func updateManifest(opts *cli.CmdOptions, route cli.Route, posArgs []string, dashArgs []string) error {
	panic(fmt.Errorf("unimplemented"))
}

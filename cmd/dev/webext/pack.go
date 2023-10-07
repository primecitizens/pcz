// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webext

import (
	"fmt"

	"github.com/primecitizens/cli"
)

// TODO: implement web extension packing

var (
	_cmdPack = cli.Cmd{
		Pattern:    "pack",
		BriefUsage: "Create a new web extension project",
		Flags:      nil,
		LocalFlags: nil,
		FlagRule:   nil,
		Run:        packWebExtension,
		Completion: nil,
		Extra:      nil,
		Children:   []*cli.Cmd{},
		State:      0,
	}
)

func packWebExtension(opts *cli.CmdOptions, route cli.Route, posArgs []string, dashArgs []string) error {
	panic(fmt.Errorf("unimplemented"))
}

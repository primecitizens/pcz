// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package dev

import (
	"github.com/primecitizens/cli"
)

func AddEnvironment(cmd *cli.Cmd) {
	_cmd.Children = append(_cmd.Children, cmd)
}

type Options struct{}

var (
	_opts    = Options{}
	_subcmds [5]*cli.Cmd
	_cmd     = cli.Cmd{
		Pattern:    "dev",
		BriefUsage: "Development helper",
		Flags:      cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &_opts),
		Children:   _subcmds[:0:len(_subcmds)],
	}
)

func Cmd() *cli.Cmd {
	return &_cmd
}

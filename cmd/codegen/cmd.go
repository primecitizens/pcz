// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package codegen

import (
	"path/filepath"

	"github.com/primecitizens/cli"
)

func Cmd() *cli.Cmd {
	return &_cmd
}

var (
	_opts    = Options{}
	_subcmds = [8]*cli.Cmd{}
	_cmd     = cli.Cmd{
		Pattern:    "codegen PLATFORM",
		BriefUsage: "Generate platform api spec and code",
		Flags:      cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, &_opts),
		PreRun:     prerun,
		Extra:      &_opts,
		Children:   _subcmds[:0:len(_subcmds)],
	}
)

func prerun(opts *cli.CmdOptions, route cli.Route, prerunAt int, posArgs, dashArgs []string) (err error) {
	copts := route[prerunAt].Extra.(*Options)
	copts.OutputDir, err = filepath.Abs(copts.OutputDir)
	return
}

type Options struct {
	OutputDir string `cli:"o|output-dir,#set output dir"`
}

func AddPlatform(cmd *cli.Cmd) {
	_cmd.Children = append(_cmd.Children, cmd)
}

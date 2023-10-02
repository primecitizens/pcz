// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	_ "embed" // for go:embed
	"errors"
	"os"
	"time"

	"github.com/primecitizens/cli"
	"github.com/primecitizens/pcz/cmd/build"
	"github.com/primecitizens/pcz/cmd/codegen"
	"github.com/primecitizens/pcz/cmd/dev"

	// codegen targets
	_ "github.com/primecitizens/pcz/cmd/codegen/android"
	_ "github.com/primecitizens/pcz/cmd/codegen/darwin"
	_ "github.com/primecitizens/pcz/cmd/codegen/js"
	_ "github.com/primecitizens/pcz/cmd/codegen/win32"

	// dev tool platforms
	_ "github.com/primecitizens/pcz/cmd/dev/android"
	_ "github.com/primecitizens/pcz/cmd/dev/ios"
	_ "github.com/primecitizens/pcz/cmd/dev/nodejs"
	_ "github.com/primecitizens/pcz/cmd/dev/web"
	_ "github.com/primecitizens/pcz/cmd/dev/webext"
)

func main() {
	err := root.Exec(&opts, os.Args[1:]...)
	if err != nil && !errors.Is(err, cli.ErrHelpHandled{}) {
		panic(err)
	}
}

var (
	root = cli.Cmd{
		Pattern:    "pcz",
		BriefUsage: "",
		Children:   _subcmds[:],
	}

	_subcmds = [...]*cli.Cmd{
		dev.Cmd(),
		build.Cmd(),
		codegen.Cmd(),
	}

	_routeBuf  [8]*cli.Cmd
	_posArgBuf [16]string
	opts       = cli.CmdOptions{
		ParseOptions: &cli.ParseOptions{
			StartTime:        time.Now(),
			HandleParseError: nil,
			PosArgsBuf:       _posArgBuf[:0:len(_posArgBuf)],
			HelpArgs:         nil,
			Extra:            nil,
		},
		RouteBuf:          _routeBuf[:0:len(_routeBuf)],
		Stdin:             os.Stdin,
		Stdout:            os.Stdout,
		Stderr:            os.Stderr,
		HandleArgError:    cli.HandleArgErrorAsHelpRequest,
		HandleHelpRequest: cli.HandleHelpRequest,
		Extra:             nil,
		SkipPostRun:       false,
		DoNotSetFlags:     false,
	}
)

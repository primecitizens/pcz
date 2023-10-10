// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package embed

import (
	"github.com/primecitizens/cli"
)

func Cmd() *cli.Cmd {
	return &embedCmd
}

var (
	embedCmd = cli.Cmd{
		Pattern:    "embed",
		BriefUsage: "Create file with data embedded",
		Children: []*cli.Cmd{
			&jsEmbedCmd,
		},
	}
)

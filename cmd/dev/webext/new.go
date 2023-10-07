// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webext

import (
	"fmt"

	"github.com/primecitizens/cli"
)

// TODO: implement creating web extension from directory template

var (
	_cmdNew = cli.Cmd{
		Pattern:    "new",
		BriefUsage: "Create a new web extension project",
		// LocalFlags: cli.NewReflectIndexer(cli.DefaultReflectVPFactory{}, nil),
		Run: createNewProject,
	}
)

type projectOptions struct {
	Popup         bool `cli:"popup,#add popup template"`
	ServiceWorker bool `cli:"service-worker|bg,#add service-worker template"`
}

func createNewProject(opts *cli.CmdOptions, route cli.Route, posArgs []string, dashArgs []string) error {
	panic(fmt.Errorf("unimplemented"))
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package main

import (
	"context"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {
	app, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	rt := wazero.NewRuntimeWithConfig(
		ctx,
		wazero.NewRuntimeConfigInterpreter().
			WithDebugInfoEnabled(true).
			WithCoreFeatures(api.CoreFeaturesV2),
	)
	defer rt.Close(ctx)

	_, err = wasi_snapshot_preview1.Instantiate(ctx, rt)
	if err != nil {
		panic(err)
	}

	_, err = rt.InstantiateWithConfig(
		ctx,
		app,
		wazero.NewModuleConfig().
			WithArgs(os.Args[2:]...).
			WithStdin(os.Stdin).
			WithStdout(os.Stdout).
			WithStderr(os.Stderr),
	)
	if err != nil {
		panic(err)
	}
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && js

package rt0

import (
	_ "github.com/primecitizens/std/ffi/js" // import symbol "wasm_export_run" and "wasm_export_getsp"

	"github.com/primecitizens/std/ffi/js/bindings"
)

func exit0() {
	bindings.Exit(0)
}

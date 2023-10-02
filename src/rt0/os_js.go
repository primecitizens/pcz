// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && js

package rt0

import (
	_ "github.com/primecitizens/std/ffi/js" // import symbol "wasm_export_run" and "wasm_export_getsp"

	"github.com/primecitizens/std/core/alloc"
	"github.com/primecitizens/std/core/alloc/sbrkalloc"
	"github.com/primecitizens/std/core/yield"
	"github.com/primecitizens/std/ffi/js/bindings"
)

var _alloc sbrkalloc.T

func malloc() alloc.M { return &_alloc }
func palloc() alloc.P { return &_alloc }

func exit0() {
	bindings.Exit(0)
	yield.Thread()
}

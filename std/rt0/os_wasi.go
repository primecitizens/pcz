// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && wasm && wasip1

package rt0

import (
	"github.com/primecitizens/pcz/std/core/alloc"
	"github.com/primecitizens/pcz/std/core/alloc/sbrkalloc"
	"github.com/primecitizens/pcz/std/ffi/wasm/wasi"
)

var _alloc sbrkalloc.T

func malloc() alloc.M { return &_alloc }
func palloc() alloc.P { return &_alloc }

func exit0() {
	wasi.Exit(0)
}

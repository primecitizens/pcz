// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm && js

package sysalloc

import (
	_ "unsafe" // for go:linkname

	"github.com/primecitizens/std/ffi/js/bindings"
)

func resetMemoryDataView() {
	bindings.ResetMemoryDataView()
}

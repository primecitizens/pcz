// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && wasm && wasip1

package rt0

import (
	"github.com/primecitizens/std/ffi/wasm/wasi"
)

func exit0() {
	wasi.Exit(0)
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package alloc

import "unsafe"

var (
	//go:linkname zerobase runtime.zerobase
	zerobase uintptr
)

func ZeroSized() unsafe.Pointer {
	return unsafe.Pointer(&zerobase)
}

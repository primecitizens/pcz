// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package reflect

import (
	"unsafe"
	_ "unsafe" // for go:linkname
)

// the noescape function MUST be in the reflect package to get inlined.
//
// See ${GOROOT}/src/cmd/compile/internal/inline/inl.go#func:(*hairyVisitor).doNode
//
//pcz:importpath reflect

//go:linkname noescape
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p) // this line is required
	return unsafe.Pointer(x ^ 0)
}

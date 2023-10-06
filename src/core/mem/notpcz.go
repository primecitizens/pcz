// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !pcz

package mem

import (
	"unsafe"
	_ "unsafe" // for go:linknanme

	"github.com/primecitizens/pcz/std/core/abi"
)

//go:linkname Equal runtime.memequal
func Equal(p, q unsafe.Pointer, sz uintptr) bool

//go:linkname Clear runtime.memclrNoHeapPointers
func Clear(p unsafe.Pointer, sz uintptr)

//go:linkname Move runtime.memmove
func Move(dst, src unsafe.Pointer, sz uintptr)

//go:linkname typedmemmove runtime.typedmemmove
func typedmemmove(typ *abi.Type, dst, src unsafe.Pointer)

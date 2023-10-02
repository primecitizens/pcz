// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package mem

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
)

// Move copies n bytes from "src" to "dst".
//
// Move ensures that any pointer in "from" is written to "to" with
// an indivisible write, so that racy reads cannot observe a
// half-written pointer. This is necessary to prevent the garbage
// collector from observing invalid pointers, and differs from memmove
// in unmanaged languages. However, memmove is only required to do
// this if "from" and "to" may contain pointers, which can only be the
// case if "from", "to", and "n" are all be word-aligned.
//
// Implementations are in move_*.s.
//
// See ${GOROOT}/src/runtime/stubs.go#func:memmove
//
//go:noescape
func Move(dst, src unsafe.Pointer, n uintptr)

//go:linkname typedmemmove runtime.typedmemmove
func typedmemmove(typ *abi.Type, dst, src unsafe.Pointer)

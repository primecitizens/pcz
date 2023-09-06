// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mem

import (
	"unsafe"

	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/num"
)

func memmove(to, from, n uintptr)

// memmove copies n bytes from "from" to "to".
//
// memmove ensures that any pointer in "from" is written to "to" with
// an indivisible write, so that racy reads cannot observe a
// half-written pointer. This is necessary to prevent the garbage
// collector from observing invalid pointers, and differs from memmove
// in unmanaged languages. However, memmove is only required to do
// this if "from" and "to" may contain pointers, which can only be the
// case if "from", "to", and "n" are all be word-aligned.
//
// Implementations are in memmove_*.s.
func Memmove[T1, T2 num.Pointer](to T1, from T2, n uintptr) {
	memmove(uintptr(to), uintptr(from), n)
}

func TypedMemmove(typ *abi.Type, dst, src unsafe.Pointer)

func TypedSliceCopy(typ *abi.Type, dstPtr unsafe.Pointer, dstLen int, srcPtr unsafe.Pointer, srcLen int) int

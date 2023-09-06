// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package runtime

import (
	"unsafe"

	stdslice "github.com/primecitizens/std/builtin/slice"
	"github.com/primecitizens/std/core/abi"
)

// see $GOROOT/src/runtime/slice.go

func makeslice(et *abi.Type, len, cap int) unsafe.Pointer {
	return stdslice.Make(et, len, cap)
}

func makeslice64(et *abi.Type, len64, cap64 int64) unsafe.Pointer {
	len := int(len64)
	if int64(len) != len64 {
		panicmakeslicelen()
	}

	cap := int(cap64)
	if int64(cap) != cap64 {
		panicmakeslicecap()
	}

	return stdslice.Make(et, len, cap)
}

func makeslicecopy(et *abi.Type, tolen int, fromlen int, from unsafe.Pointer) unsafe.Pointer {
	return stdslice.MakeCopy(et, tolen, fromlen, from)
}

func growslice(oldPtr unsafe.Pointer, newLen, oldCap, num int, et *abi.Type) stdslice.Header {
	return stdslice.Grow(oldPtr, newLen, oldCap, num, et)
}

func slicecopy(toPtr unsafe.Pointer, toLen int, fromPtr unsafe.Pointer, fromLen int, width uintptr) int {
	return stdslice.Copy(toPtr, toLen, fromPtr, fromLen, width)
}

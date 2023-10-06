// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package stdprint

import (
	"unsafe"

	stdslice "github.com/primecitizens/pcz/std/builtin/slice"
	stdstring "github.com/primecitizens/pcz/std/builtin/string"
	"github.com/primecitizens/pcz/std/plat/android"
	"github.com/primecitizens/pcz/std/plat/libc"
)

func gwrite(b []byte) {
	if len(b) == 0 {
		return
	}

	ptr, sz := stdslice.ToListAndSize[int](b)
	android.Log(android.LogPriority_Info, "print\x00", stdstring.SliceByteToStringTmp(ptr, sz))

	libc.Write(2, uintptr(unsafe.Pointer(&b[0])), int32(len(b)))
}

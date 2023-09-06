// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !(android || js || wasip1)

package stdprint

import (
	"unsafe"

	"github.com/primecitizens/std/plat/libc"
)

func gwrite(b []byte) {
	if len(b) == 0 {
		return
	}

	libc.Write(2, uintptr(unsafe.Pointer(&b[0])), int32(len(b)))
}

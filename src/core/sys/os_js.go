// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && js

package sys

import (
	"unsafe"
	_ "unsafe" // for go:linkname

	"github.com/primecitizens/std/core/alloc/sysalloc"
	"github.com/primecitizens/std/core/mark"
	"github.com/primecitizens/std/ffi/js/bindings"
)

var (
	args []string
	envs []string
)

func ntharg(i int) string {
	return args[i]
}

func nthenv(i int) string {
	return envs[i]
}

func init() {
	var (
		count     uint32
		totalSize = uintptr(bindings.ArgsEnvsSize(
			uint32(uintptr(unsafe.Pointer(mark.NoEscape(&count)))),
		))

		i uint32
		n int32

		buf []byte
	)

	if count == 0 || totalSize == 0 {
		return
	}

	totalSize += unsafe.Sizeof(string("")) * uintptr(count)

	buf = unsafe.Slice((*byte)(sysalloc.Sbrk(totalSize)), totalSize)

	args = unsafe.Slice(
		(*string)(unsafe.Pointer(unsafe.SliceData(buf))),
		count,
	)
	buf = buf[unsafe.Sizeof(string(""))*uintptr(count):]

	for i = 0; ; i++ {
		n = bindings.NthSysInto(
			i,
			uint32(uintptr(unsafe.Pointer(unsafe.SliceData(buf)))),
			uint32(len(buf)),
		)
		if n == -1 {
			break
		}

		args[i] = unsafe.String(unsafe.SliceData(buf), n)
		buf = buf[n:]
	}

	envs = args[i:]
	args = args[:i]

	for {
		i++
		n = bindings.NthSysInto(
			i,
			uint32(uintptr(unsafe.Pointer(unsafe.SliceData(buf)))),
			uint32(len(buf)),
		)
		if n == -1 {
			break
		}

		envs[i-uint32(len(args))-1] = unsafe.String(unsafe.SliceData(buf), n)
		buf = buf[n:]
	}
}

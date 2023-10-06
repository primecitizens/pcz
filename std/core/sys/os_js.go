// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && js

package sys

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/core/sys/bindings"
	"github.com/primecitizens/pcz/std/ffi/js"
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
		totalSize = uintptr(
			bindings.Sizes(
				js.Pointer(&count),
			),
		)

		i uint32
		n int32
	)

	if count == 0 || totalSize == 0 {
		return
	}

	totalSize += unsafe.Sizeof(string("")) * uintptr(count)

	buf := make([]byte, totalSize)
	args = unsafe.Slice(
		(*string)(unsafe.Pointer(unsafe.SliceData(buf))),
		count,
	)
	buf = buf[unsafe.Sizeof(string(""))*uintptr(count):]

	for i = 0; ; i++ {
		n = bindings.Nth(
			i,
			js.SliceData(buf),
			js.SizeU(len(buf)),
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
		n = bindings.Nth(
			i,
			js.SliceData(buf),
			js.SizeU(len(buf)),
		)
		if n == -1 {
			break
		}

		envs[i-uint32(len(args))-1] = unsafe.String(unsafe.SliceData(buf), n)
		buf = buf[n:]
	}
}

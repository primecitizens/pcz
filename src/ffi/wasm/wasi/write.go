// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package wasi

import (
	"unsafe"

	"github.com/primecitizens/std/core/mark"
)

//go:wasmimport wasi_snapshot_preview1 fd_write
//go:noescape
func fd_write(fd FD, iovs unsafe.Pointer, iovsLen Size, nwritten unsafe.Pointer) Errno

func Write(fd FD, iovs ...IOBuffer) (Size, Errno) {
	if len(iovs) > 0 {
		iovs = mark.NoEscapeSlice(iovs)
		var n Size
		errno := fd_write(
			fd,
			unsafe.Pointer(mark.NoEscape(&iovs[0])),
			Size(len(iovs)),
			unsafe.Pointer(mark.NoEscape(&n)),
		)
		return n, errno
	}

	var (
		v0 IOBuffer
		n  Size
	)
	errno := fd_write(
		fd,
		unsafe.Pointer(mark.NoEscape(&v0)),
		1,
		unsafe.Pointer(mark.NoEscape(&n)),
	)
	return n, errno
}

//go:wasmimport wasi_snapshot_preview1 fd_pwrite
//go:noescape
func fd_pwrite(fd FD, iovs unsafe.Pointer, iovsLen Size, offset FileSize, nwritten unsafe.Pointer) Errno

func Pwrite(fd FD, offset int64, iovs ...IOBuffer) (Size, Errno) {
	if len(iovs) > 0 {
		iovs = mark.NoEscapeSlice(iovs)
		var n Size
		errno := fd_pwrite(
			fd,
			unsafe.Pointer(mark.NoEscape(&iovs[0])),
			Size(len(iovs)),
			FileSize(offset),
			unsafe.Pointer(mark.NoEscape(&n)),
		)
		return n, errno
	}

	var (
		v0 IOBuffer
		n  Size
	)
	errno := fd_pwrite(
		fd,
		unsafe.Pointer(mark.NoEscape(&v0)),
		1,
		FileSize(offset),
		unsafe.Pointer(mark.NoEscape(&n)),
	)
	return n, errno
}

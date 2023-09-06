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

//go:wasmimport wasi_snapshot_preview1 fd_read
//go:noescape
func fd_read(fd FD, iovs unsafe.Pointer, iovsLen Size, nread unsafe.Pointer) Errno

func Read(fd FD, iovec ...IOBuffer) (n Size, errno Errno) {
	if len(iovec) > 0 {
		iovec = mark.NoEscapeSlice(iovec)
		errno = fd_read(
			fd,
			unsafe.Pointer(mark.NoEscape(&iovec[0])),
			Size(len(iovec)),
			unsafe.Pointer(mark.NoEscape(&n)),
		)
		return
	}

	var v0 IOBuffer
	errno = fd_read(
		fd,
		unsafe.Pointer(mark.NoEscape(&v0)),
		1,
		unsafe.Pointer(mark.NoEscape(&n)),
	)
	return
}

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_preadfd-fd-iovs-iovec_array-offset-filesize---resultsize-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_pread
//go:noescape
func fd_pread(fd FD, iovs unsafe.Pointer, iovsLen Size, offset FileSize, nread unsafe.Pointer) Errno

func Pread(fd FD, offset int64, iovec ...IOBuffer) (n Size, errno Errno) {
	if len(iovec) > 0 {
		iovec = mark.NoEscapeSlice(iovec)
		errno = fd_pread(
			fd,
			unsafe.Pointer(mark.NoEscape(&iovec[0])),
			Size(len(iovec)),
			FileSize(offset),
			unsafe.Pointer(&n),
		)
		return
	}

	var v0 IOBuffer
	errno = fd_pread(
		fd,
		unsafe.Pointer(mark.NoEscape(&v0)),
		Size(len(iovec)),
		FileSize(offset),
		unsafe.Pointer(&n),
	)
	return
}

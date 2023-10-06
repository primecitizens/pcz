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
)

//go:wasmimport wasi_snapshot_preview1 fd_read
//go:noescape
func Read(
	fd FD,
	iovs unsafe.Pointer,
	iovsLen Size,
	nread unsafe.Pointer,
) Errno

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_preadfd-fd-iovs-iovec_array-offset-filesize---resultsize-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_pread
//go:noescape
func Pread(
	fd FD,
	iovs unsafe.Pointer,
	iovsLen Size,
	offset Filesize,
	nread unsafe.Pointer,
) Errno

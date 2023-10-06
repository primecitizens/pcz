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

type Stat_t struct {
	Dev      uint64
	Ino      uint64
	Filetype Filetype
	Nlink    uint64
	Size_    uint64
	Atime    uint64
	Mtime    uint64
	Ctime    uint64

	Mode int

	// Uid and Gid are always zero on wasip1 platforms
	Uid uint32
	Gid uint32
}

//go:wasmimport wasi_snapshot_preview1 path_filestat_get
//go:noescape
func StatPath(
	dirFD FD,
	flags LookupFlags,
	path unsafe.Pointer,
	pathLen Size,
	buf unsafe.Pointer,
) Errno

//go:wasmimport wasi_snapshot_preview1 fd_filestat_get
//go:noescape
func StatFD(fd FD, buf unsafe.Pointer) Errno

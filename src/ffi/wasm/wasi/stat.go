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

	"github.com/primecitizens/std/core/assert"
)

type Stat_t struct {
	Dev      uint64
	Ino      uint64
	Filetype Filetype
	Nlink    uint64
	Size     uint64
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
func path_filestat_get(fd FD, flags LookupFlags, path unsafe.Pointer, pathLen Size, buf unsafe.Pointer) Errno

func Stat(path string, st *Stat_t) Errno {
	if len(path) == 0 {
		return EINVAL
	}
	assert.TODO()
	return 0
	// dirFd, pathPtr, pathLen := preparePath(path)
	// errno = path_filestat_get(dirFd, LOOKUP_SYMLINK_FOLLOW, pathPtr, pathLen, unsafe.Pointer(st))
	// setDefaultMode(st)
	// return
}

func Lstat(path string, st *Stat_t) Errno {
	if len(path) == 0 {
		return EINVAL
	}
	assert.TODO()
	return 0
	// dirFd, pathPtr, pathLen := preparePath(path)
	// errno = path_filestat_get(dirFd, 0, pathPtr, pathLen, unsafe.Pointer(st))
	// setDefaultMode(st)
	// return
}

//go:wasmimport wasi_snapshot_preview1 fd_filestat_get
//go:noescape
func fd_filestat_get(fd FD, buf unsafe.Pointer) Errno

func Fstat(fd FD, st *Stat_t) Errno {
	errno := fd_filestat_get(fd, unsafe.Pointer(st))
	setDefaultMode(st)
	return errno
}

func setDefaultMode(st *Stat_t) {
	// WASI does not support unix-like permissions, but Go programs are likely
	// to expect the permission bits to not be zero so we set defaults to help
	// avoid breaking applications that are migrating to WASM.
	if st.Filetype == Filetype_DIRECTORY {
		st.Mode = 0700
	} else {
		st.Mode = 0600
	}
}

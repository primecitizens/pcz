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

//go:wasmimport wasi_snapshot_preview1 path_link
//go:noescape
func path_link(oldFd int32, oldFlags LookupFlags, oldPath unsafe.Pointer, oldPathLen Size, newFd int32, newPath unsafe.Pointer, newPathLen Size) Errno

func Link(path, link string) Errno {
	if len(path) == 0 || len(link) == 0 {
		return EINVAL
	}

	assert.TODO()
	return 0
	// oldDirFd, oldPathPtr, oldPathLen := preparePath(path)
	// newDirFd, newPathPtr, newPathLen := preparePath(link)
	// return path_link(
	// 	oldDirFd,
	// 	0,
	// 	oldPathPtr,
	// 	oldPathLen,
	// 	newDirFd,
	// 	newPathPtr,
	// 	newPathLen,
	// )
}

//go:wasmimport wasi_snapshot_preview1 path_unlink_file
//go:noescape
func path_unlink_file(fd FD, path unsafe.Pointer, pathLen Size) Errno

func Unlink(path string) Errno {
	if len(path) == 0 {
		return EINVAL
	}
	assert.TODO()
	return 0
	// dirFd, pathPtr, pathLen := preparePath(path)
	// return path_unlink_file(dirFd, pathPtr, pathLen)
}

//go:wasmimport wasi_snapshot_preview1 path_symlink
//go:noescape
func path_symlink(oldPath unsafe.Pointer, oldPathLen Size, fd FD, newPath unsafe.Pointer, newPathLen Size) Errno

func Symlink(path, link string) Errno {
	if len(path) == 0 || len(link) == 0 {
		return EINVAL
	}
	assert.TODO()
	return 0
	// dirFd, pathPtr, pathlen := preparePath(link)
	// return path_symlink(
	// 	stringPointer(path),
	// 	Size(len(path)),
	// 	dirFd,
	// 	pathPtr,
	// 	pathlen,
	// )
}

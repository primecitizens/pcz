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

//go:wasmimport wasi_snapshot_preview1 path_open
//go:noescape
func path_open(rootFD int32, dirflags LookupFlags, path unsafe.Pointer, pathLen Size, oflags OpenFlags, fsRightsBase Rights, fsRightsInheriting Rights, fsFlags FDflag, fd unsafe.Pointer) Errno

func Open(path string, openmode int, perm uint32) (FD, Errno) {
	if len(path) == 0 {
		return -1, EINVAL
	}
	assert.TODO()
	return 0, 0
	// dirFd, pathPtr, pathLen := preparePath(path)
	//
	// var oflags OpenFlags
	//
	//	if (openmode & O_CREATE) != 0 {
	//		oflags |= OpenFlag_CREATE
	//	}
	//
	//	if (openmode & O_TRUNC) != 0 {
	//		oflags |= OpenFlag_TRUNC
	//	}
	//
	//	if (openmode & O_EXCL) != 0 {
	//		oflags |= OpenFlag_EXCL
	//	}
	//
	// var rights Rights
	// switch openmode & (O_RDONLY | O_WRONLY | O_RDWR) {
	// case O_RDONLY:
	//
	//	rights = fileRights & ^writeRights
	//
	// case O_WRONLY:
	//
	//	rights = fileRights & ^readRights
	//
	// case O_RDWR:
	//
	//		rights = fileRights
	//	}
	//
	// var fdflags FDflags
	//
	//	if (openmode & O_APPEND) != 0 {
	//		fdflags |= FDFLAG_APPEND
	//	}
	//
	//	if (openmode & O_SYNC) != 0 {
	//		fdflags |= FDFLAG_SYNC
	//	}
	//
	// var fd FD
	// errno := path_open(
	//
	//	dirFd,
	//	LOOKUP_SYMLINK_FOLLOW,
	//	pathPtr,
	//	pathLen,
	//	oflags,
	//	rights,
	//	fileRights,
	//	fdflags,
	//	unsafe.Pointer(&fd),
	//
	// )
	//
	//	if errno == EISDIR && oflags == 0 && fdflags == 0 && ((rights & writeRights) == 0) {
	//		// wasmtime and wasmedge will error if attempting to open a directory
	//		// because we are asking for too many rights. However, we cannot
	//		// determine ahread of time if the path we are about to open is a
	//		// directory, so instead we fallback to a second call to path_open with
	//		// a more limited set of rights.
	//		//
	//		// This approach is subject to a race if the file system is modified
	//		// concurrently, so we also inject OFLAG_DIRECTORY to ensure that we do
	//		// not accidentally open a file which is not a directory.
	//		errno = path_open(
	//			dirFd,
	//			LOOKUP_SYMLINK_FOLLOW,
	//			pathPtr,
	//			pathLen,
	//			oflags|OFLAG_DIRECTORY,
	//			rights&dirRights,
	//			fileRights,
	//			fdflags,
	//			unsafe.Pointer(&fd),
	//		)
	//	}
	//
	// return fd, errno
}

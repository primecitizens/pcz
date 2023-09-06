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

//go:wasmimport wasi_snapshot_preview1 fd_seek
//go:noescape
func fd_seek(fd FD, offset FileDelta, whence uint32, newoffset unsafe.Pointer) Errno

func Seek(fd FD, offset FileDelta, whence uint32) (FileSize, Errno) {
	var newoffset FileSize
	errno := fd_seek(
		fd, offset, whence, unsafe.Pointer(mark.NoEscape(&newoffset)),
	)
	return newoffset, errno
}

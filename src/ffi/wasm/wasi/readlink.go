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

//go:wasmimport wasi_snapshot_preview1 path_readlink
//go:noescape
func path_readlink(fd FD, path unsafe.Pointer, pathLen Size, buf unsafe.Pointer, bufLen Size, nwritten unsafe.Pointer) Errno

// For some reason wasmtime returns ERANGE when the output buffer is
// shorter than the symbolic link value.
func Readlink(path string, buf []byte) (n Size, errno Errno) {
	if len(path) == 0 {
		return 0, EINVAL
	}
	if len(buf) == 0 {
		return 0, 0
	}

	assert.TODO()
	return 0, 0
	// dirFd, pathPtr, pathLen := preparePath(path)
	// errno = path_readlink(
	// 	dirFd,
	// 	pathPtr,
	// 	pathLen,
	// 	unsafe.Pointer(&buf[0]),
	// 	Size(len(buf)),
	// 	unsafe.Pointer(&n),
	// )
	// return
}

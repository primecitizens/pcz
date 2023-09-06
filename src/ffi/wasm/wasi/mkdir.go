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

//go:wasmimport wasi_snapshot_preview1 path_create_directory
//go:noescape
func path_create_directory(fd FD, path unsafe.Pointer, pathLen Size) Errno

func Mkdir(path string, perm uint32) Errno {
	if len(path) == 0 {
		return EINVAL
	}
	assert.TODO()
	return 0
	// dirFd, pathPtr, pathLen := preparePath(path)
	// errno := path_create_directory(dirFd, pathPtr, pathLen)
	// return errno
}

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

//go:wasmimport wasi_snapshot_preview1 path_remove_directory
//go:noescape
func path_remove_directory(fd FD, path unsafe.Pointer, pathLen Size) Errno

func Rmdir(path string) Errno {
	if len(path) == 0 {
		return EINVAL
	}
	assert.TODO()
	return 0
	// dirFd, pathPtr, pathLen := preparePath(path)
	// return path_remove_directory(dirFd, pathPtr, pathLen)
}

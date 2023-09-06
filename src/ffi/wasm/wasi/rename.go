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

//go:wasmimport wasi_snapshot_preview1 path_rename
//go:noescape
func path_rename(
	oldFd FD, oldPath unsafe.Pointer, oldPathLen Size,
	newFd FD, newPath unsafe.Pointer, newPathLen Size,
) Errno

func Rename(from, to string) Errno {
	if len(from) == 0 || len(to) == 0 {
		return EINVAL
	}

	assert.TODO()
	return 0
	// oldDirFd, oldPathPtr, oldPathLen := preparePath(from)
	// newDirFd, newPathPtr, newPathLen := preparePath(to)
	// return path_rename(
	// 	oldDirFd,
	// 	oldPathPtr,
	// 	oldPathLen,
	// 	newDirFd,
	// 	newPathPtr,
	// 	newPathLen,
	// )
}

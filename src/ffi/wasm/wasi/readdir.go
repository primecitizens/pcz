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

//go:wasmimport wasi_snapshot_preview1 fd_readdir
//go:noescape
func fd_readdir(
	fd FD, buf unsafe.Pointer, bufLen Size,
	cookie DirCookie,
	nwritten unsafe.Pointer,
) Errno

func ReadDir(fd FD, cookie DirCookie, buf []byte) (Size, Errno) {
	if len(buf) == 0 {
		return 0, 0
	}
	buf = mark.NoEscapeSlice(buf)

	var nwritten Size
	errno := fd_readdir(
		fd,
		unsafe.Pointer(&buf[0]),
		Size(len(buf)),
		cookie,
		unsafe.Pointer(&nwritten),
	)
	return nwritten, errno
}

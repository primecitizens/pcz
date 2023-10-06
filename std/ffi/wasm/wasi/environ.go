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

// Environ depends on values returned by EnvironSizes
//
// envv is [count]uint32
// environBuf is [totalBytes]byte
//
//go:wasmimport wasi_snapshot_preview1 environ_get
//go:noescape
func Environ(envv, environBuf unsafe.Pointer) Errno

//go:wasmimport wasi_snapshot_preview1 environ_sizes_get
//go:noescape
func environ_sizes_get(environCount, environBufLen unsafe.Pointer) Errno

func EnvironSizes() (envc, totalBytes Size, errno Errno) {
	var n, sz Size
	errno = environ_sizes_get(unsafe.Pointer(&n), unsafe.Pointer(&sz))
	return n, sz, errno
}

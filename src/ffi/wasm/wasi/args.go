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

//go:wasmimport wasi_snapshot_preview1 args_sizes_get
//go:noescape
func args_sizes_get(argc, argvBufLen unsafe.Pointer) Errno

func ArgsSizes() (argc, totalBytes Size, errno Errno) {
	var n, sz Size
	errno = args_sizes_get(unsafe.Pointer(&n), unsafe.Pointer(&sz))
	return n, sz, errno
}

// ArgsGet depends on values returned by args_sizes_get
// argv is [argc]uint32
// argv is [totalBytes]byte
//
//go:wasmimport wasi_snapshot_preview1 args_get
//go:noescape
func ArgsGet(argv, argvBuf unsafe.Pointer) Errno

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

type ShutdownFlags uint32

const (
	Shutdown_RD ShutdownFlags = 0x1
	Shutdown_WR ShutdownFlags = 0x2
	Shutdown_RW ShutdownFlags = Shutdown_RD | Shutdown_WR
)

func SockAccept(fd FD) (newfd FD, errno Errno) {
	errno = sock_accept(fd, 0, unsafe.Pointer(&newfd))
	return
}

//go:wasmimport wasi_snapshot_preview1 sock_accept
//go:noescape
func sock_accept(fd FD, flags FDflag, newfd unsafe.Pointer) Errno

//go:wasmimport wasi_snapshot_preview1 sock_shutdown
//go:noescape
func SockShutdown(fd FD, how ShutdownFlags) Errno

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

//go:wasmimport wasi_snapshot_preview1 sock_accept
//go:noescape
func SockAccept(fd FD, flags FDflags, newfd unsafe.Pointer) Errno

type Riflags uint16

const (
	Riflag_RECV_PEEK    Riflags = 0x1 // Returns the message without removing it from the socket's receive queue.
	Riflag_RECV_WAITALL               // On byte-stream sockets, block until the full amount of data can be returned.
)

type Roflags uint16

const (
	Roflags_RECV_DATA_TRUNCATED Roflags = 0x1 // Message data has been truncated.
)

type SockRecvResult struct {
	NRead   Size
	Roflags Roflags
}

// TODO: go1.21 doesn't support uint16 args

// //go:wasmimport wasi_snapshot_preview1 sock_recv
// //go:noescape
// func SockRecv(fd FD, iovs unsafe.Pointer, n Size, riflags Riflags, result unsafe.Pointer) Errno

type Siflags uint16

// //go:wasmimport wasi_snapshot_preview1 sock_send
// //go:noescape
// func SockSend(fd FD, iovs unsafe.Pointer, n Size, siflags Siflags, nwritten Size) Errno

type SDflags uint32

const (
	SDflag_RD SDflags = 0x1 // Disables further receive operations.
	SDflag_WR SDflags = 0x2 // Disables further send operations.
	SDflag_RW SDflags = SDflag_RD | SDflag_WR
)

//go:wasmimport wasi_snapshot_preview1 sock_shutdown
//go:noescape
func SockShutdown(fd FD, how SDflags) Errno

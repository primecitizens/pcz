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

type FDpreopenType = uint8

const (
	FDpreopenType_DIR FDpreopenType = iota
)

type FDprestatDir struct {
	PrestateNameLen Size
}

type FDprestat_t struct {
	Type FDpreopenType
	Dir  FDprestatDir
}

//go:wasmimport wasi_snapshot_preview1 fd_prestat_get
//go:noescape
func fd_prestat_get(fd FD, prestat unsafe.Pointer) Errno

func FDprestat(fd FD, out *FDprestat_t) Errno {
	return fd_prestat_get(fd, unsafe.Pointer(mark.NoEscape(out)))
}

//go:wasmimport wasi_snapshot_preview1 fd_prestat_dir_name
//go:noescape
func fd_prestat_dir_name(fd FD, path unsafe.Pointer, pathLen Size) Errno

func FDprestatDirname(fd FD, buf []byte) Errno {
	return fd_prestat_dir_name(fd, bytesPointer(buf), Size(len(buf)))
}

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-FDstat_t-record
// fdflags must be at offset 2, hence the uint16 type rather than the
// fdflags (uint32) type.
type FDstat_t struct {
	Type             Filetype
	FDflags          uint16
	RightsBase       Rights
	RightsInheriting Rights
}

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-fd_fdstat_set_rightsfd-fd-fs_rights_base-rights-fs_rights_inheriting-rights---result-errno
//
//go:wasmimport wasi_snapshot_preview1 fd_fdstat_set_rights
//go:noescape
func FDstatSetRights(fd FD, rightsBase, rightsInheriting Rights) Errno

//go:wasmimport wasi_snapshot_preview1 fd_fdstat_get
//go:noescape
func fdstat_get(fd FD, buf unsafe.Pointer) Errno

func FDstat(fd FD, out *FDstat_t) Errno {
	return fdstat_get(fd, unsafe.Pointer(mark.NoEscape(out)))
}

//go:wasmimport wasi_snapshot_preview1 fd_fdstat_set_flags
//go:noescape
func FDstatSetFlags(fd FD, flags FDflag) Errno

func FDstatGetFlags(fd FD) (FDflag, Errno) {
	var stat FDstat_t
	errno := fdstat_get(fd, unsafe.Pointer(mark.NoEscape(&stat)))
	return FDflag(stat.FDflags), errno
}

func FDstatGetType(fd FD) (Filetype, Errno) {
	var stat FDstat_t
	errno := fdstat_get(fd, unsafe.Pointer(mark.NoEscape(&stat)))
	return stat.Type, errno
}

func SetNonblock(fd FD, nonblocking bool) Errno {
	flags, errno := FDstatGetFlags(fd)
	if errno != 0 {
		return errno
	}
	if nonblocking {
		flags |= FDflag_NONBLOCK
	} else {
		flags &^= FDflag_NONBLOCK
	}
	return FDstatSetFlags(fd, flags)
}

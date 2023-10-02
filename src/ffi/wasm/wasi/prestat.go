// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasip1

package wasi

import (
	"unsafe"
)

type PrestatType = uint8

const (
	PreopenType_DIR PrestatType = iota
)

type PrestatDir struct {
	PrestateNameLen Size
}

type Prestat_t struct {
	Type PrestatType
	Dir  PrestatDir
}

//go:wasmimport wasi_snapshot_preview1 fd_prestat_get
//go:noescape
func Prestat(fd FD, prestat unsafe.Pointer) Errno

//go:wasmimport wasi_snapshot_preview1 fd_prestat_dir_name
//go:noescape
func PrestatDirname(fd FD, buf unsafe.Pointer, len Size) Errno

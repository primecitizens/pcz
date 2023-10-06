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

	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/core/mark"
)

type Dirent struct {
	// The offset of the next directory entry stored in this directory.
	Next Dircookie
	// The serial number of the file referred to by this directory entry.
	Ino uint64
	// The length of the name of the directory entry.
	Namlen uint32

	// The type of the file referred to by this directory entry.
	Type Filetype
	_    [3]byte
}

func (ent *Dirent) Decode(buf []byte) (name string, next []byte, ok bool) {
	const (
		direntOffsetsOK = true &&
			unsafe.Offsetof(Dirent{}.Next) == 0 &&
			unsafe.Offsetof(Dirent{}.Ino) == 8 &&
			unsafe.Offsetof(Dirent{}.Namlen) == 16 &&
			unsafe.Offsetof(Dirent{}.Type) == 20

		direntSize = unsafe.Sizeof(Dirent{})
	)

	if !direntOffsetsOK || direntSize != 24 {
		assert.Throw("dirent", "size", "or", "offsets", "not", "match")
	}

	if uintptr(len(buf)) < direntSize {
		next = buf[:0]
		return
	}

	*ent = *(*Dirent)(mark.NoEscapeSliceDataPointer(buf))
	buf = buf[direntSize:]
	if len(buf) < int(ent.Namlen) {
		next = buf[:0]
		return
	}

	name = mark.NoEscapeBytesString(buf)[:ent.Namlen]
	next = buf[ent.Namlen:]
	ok = true
	return
}

//go:wasmimport wasi_snapshot_preview1 fd_readdir
//go:noescape
func Readdir(
	fd FD,
	buf unsafe.Pointer,
	bufLen Size,
	cookie Dircookie,
	nwritten unsafe.Pointer,
) Errno

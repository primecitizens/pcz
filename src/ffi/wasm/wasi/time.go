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
	"github.com/primecitizens/std/core/mark"
)

type ClockID = uint32

const (
	ClockRealtime ClockID = iota
	ClockMonotonic
	ClockProcessCPUTimeID
	ClockThreadCPUTimeID
)

//go:wasmimport wasi_snapshot_preview1 clock_time_get
//go:noescape
func ClockTimeGet(id ClockID, precision Timestamp, time unsafe.Pointer) Errno

func Walltime() (sec int64, nsec int32, errno Errno) {
	var time Timestamp
	errno = ClockTimeGet(ClockRealtime, 0, unsafe.Pointer(mark.NoEscape(&time)))
	if errno == 0 {
		sec = int64(time / 1000000000)
		nsec = int32(time % 1000000000)
	}
	return
}

func Nanotime() (mono int64, errno Errno) {
	var time Timestamp
	errno = ClockTimeGet(ClockMonotonic, 0, unsafe.Pointer(mark.NoEscape(&time)))
	if errno == 0 {
		mono = int64(time)
	}
	return
}

// Now is an alias of Walltime, but throws on failed ClockTimeGet call.
func Now() (sec int64, nsec int32) {
	sec, nsec, errno := Walltime()
	if errno != 0 {
		assert.Throw("clock_time_get failed")
	}
	return
}

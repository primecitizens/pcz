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

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

func bytesPointer(b []byte) unsafe.Pointer {
	return unsafe.Pointer(mark.NoEscape(unsafe.SliceData(b)))
}

func stringPointer(s string) unsafe.Pointer {
	return unsafe.Pointer(mark.NoEscape(unsafe.StringData(s)))
}

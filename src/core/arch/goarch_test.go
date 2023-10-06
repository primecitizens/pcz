// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arch_test

import (
	"testing"
	"unsafe"

	"github.com/primecitizens/pcz/std/core/arch"
)

func TestUintSize(t *testing.T) {
	var x uint
	if want := unsafe.Sizeof(x) * 8; arch.UintBits != want {
		t.Fatalf("UintSize = %d; want %d", arch.UintBits, want)
	}
}

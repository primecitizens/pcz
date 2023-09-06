// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package rt0

import (
	"github.com/primecitizens/std/core/alloc/sysalloc"
	"github.com/primecitizens/std/core/arch"
	"github.com/primecitizens/std/core/debug"
	"github.com/primecitizens/std/core/stack"
)

// rt0stack is typed to make $rt0stack__size available to assembly.
type rt0stack struct {
	_  [7]uintptr
	pc uintptr
}

var tmpStack rt0stack

//go:nosplit
func _rt0_wasm() {
	if stacksize < 32 {
		stacksize = arch.DefaultPhysPageSize // 64KB
	}

	sp := sysalloc.Sbrk(stacksize)
	if sp == nil {
		debug.Abort()
	}

	// rt0() requires 32 bytes of stack (4 * sizeof uintptr) to run
	// stacksize -= 32
	stack.SetSP(uintptr(sp))
	rt0(0, nil)
}

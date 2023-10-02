// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz

package debug

import (
	_ "unsafe" // for go:linkname

	"github.com/primecitizens/std/core/arch"
)

const minFrameSize = arch.MinFrameSize

// Breakpoint executes a breakpoint trap.
func Breakpoint()

func Abort()

// Return0 is a stub used to return 0 from deferproc.
// It is called at the very end of deferproc to signal
// the calling Go function that it should not jump
// to deferreturn.
func Return0()

// The implementation may be a compiler intrinsic; there is not
// necessarily code implementing this on every platform.
//
// For example:
//
//	func f(arg1, arg2, arg3 int) {
//		pc := getcallerpc()
//		sp := getcallersp()
//	}
//
// These two lines find the PC and SP immediately following
// the call to f (where f will return).
//
// The call to getcallerpc and getcallersp must be done in the
// frame being asked about.
//
// The result of getcallersp is correct at the time of the return,
// but it may be invalidated by any subsequent call to a function
// that might relocate the stack in order to grow or shrink it.
// A general rule is that the result of getcallersp should be used
// immediately and can only be passed to nosplit functions.

// GetCallerPC returns the program counter (PC) of its caller's caller.
func GetCallerPC() uintptr

// GetCallerSP returns the stack pointer (SP) of its caller's caller.
func GetCallerSP() uintptr

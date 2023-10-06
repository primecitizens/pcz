// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package stack

import (
	"github.com/primecitizens/pcz/std/core/abi"
)

// A Frame holds information about a single stack frame.
//
// see $GOROOT/src/runtime/symtab.go#stkframe
type Frame struct {
	// fn is the function being run in this frame. If there is
	// inlining, this is the outermost function.
	fn abi.FuncInfo

	// pc is the program counter within fn.
	//
	// The meaning of this is subtle:
	//
	// - Typically, this frame performed a regular function call
	//   and this is the return PC (just after the CALL
	//   instruction). In this case, pc-1 reflects the CALL
	//   instruction itself and is the correct source of symbolic
	//   information.
	//
	// - If this frame "called" sigpanic, then pc is the
	//   instruction that panicked, and pc is the correct address
	//   to use for symbolic information.
	//
	// - If this is the innermost frame, then PC is where
	//   execution will continue, but it may not be the
	//   instruction following a CALL. This may be from
	//   cooperative preemption, in which case this is the
	//   instruction after the call to morestack. Or this may be
	//   from a signal or an un-started goroutine, in which case
	//   PC could be any instruction, including the first
	//   instruction in a function. Conventionally, we use pc-1
	//   for symbolic information, unless pc == fn.entry(), in
	//   which case we use pc.
	pc uintptr

	// continpc is the PC where execution will continue in fn, or
	// 0 if execution will not continue in this frame.
	//
	// This is usually the same as pc, unless this frame "called"
	// sigpanic, in which case it's either the address of
	// deferreturn or 0 if this frame will never execute again.
	//
	// This is the PC to use to look up GC liveness for this frame.
	continpc uintptr

	lr   uintptr // program counter at caller aka link register
	sp   uintptr // stack pointer at pc
	fp   uintptr // stack pointer at caller aka frame pointer
	varp uintptr // top of local variables
	argp uintptr // pointer to function arguments
}

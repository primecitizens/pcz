// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && riscv64

#include "textflag.h"

// X2 is the stack pointer

TEXT ·GetSP<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVD X2, X10
	RET

TEXT ·SetSP<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVD X10, X2
	RET

// Called during function prolog when more stack is needed.
// Called with return address (i.e. caller's PC) in X5 (aka T0),
// and the LR register contains the caller's LR.
//
// The traceback routines see morestack on a g0 as being
// the top of a stack (for example, morestack calling newstack
// calling the scheduler calling newm calling gc), so we must
// record an argument size. For that purpose, it has no arguments.
TEXT runtime·morestack(SB),NOSPLIT|NOFRAME,$0-0
	// TODO: implement
	UNDEF // crash

// morestack but not preserving ctxt.
TEXT runtime·morestack_noctxt(SB),NOSPLIT|NOFRAME,$0-0
	// Force SPWRITE. This function doesn't actually write SP,
	// but it is called with a special calling convention where
	// the caller doesn't save LR on stack but passes it as a
	// register, and the unwinder currently doesn't understand.
	// Make it SPWRITE to stop unwinding. (See issue 54332)
	MOV X2, X2

	MOV ZERO, CTXT
	JMP runtime·morestack(SB)

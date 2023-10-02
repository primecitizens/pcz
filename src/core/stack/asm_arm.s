// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && arm

#include "textflag.h"

TEXT ·GetSP(SB),NOSPLIT,$0-4
	MOVW R13, R0
	ADD $8, R0
	MOVW R0, sp+0(FP)
	RET

TEXT ·SetSP(SB),NOSPLIT,$0-4
	MOVW sp+0(FP), R13
	RET

// Called during function prolog when more stack is needed.
// R3 prolog's LR
// using NOFRAME means do not save LR on stack.
//
// The traceback routines see morestack on a g0 as being
// the top of a stack (for example, morestack calling newstack
// calling the scheduler calling newm calling gc), so we must
// record an argument size. For that purpose, it has no arguments.
TEXT runtime·morestack(SB),NOSPLIT|NOFRAME,$0-0
	// TODO: implement
	MOVW $0, R0
	MOVW (R0), R1 // crash

TEXT runtime·morestack_noctxt(SB),NOSPLIT|NOFRAME,$0-0
	// Force SPWRITE. This function doesn't actually write SP,
	// but it is called with a special calling convention where
	// the caller doesn't save LR on stack but passes it as a
	// register (R3), and the unwinder currently doesn't understand.
	// Make it SPWRITE to stop unwinding. (See issue 54332)
	MOVW R13, R13

	MOVW $0, R7
	B runtime·morestack(SB)

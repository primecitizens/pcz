// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build mips64 || mips64le

#include "textflag.h"

TEXT ·GetSP(SB),NOSPLIT,$0-8
	MOVV R29, R0
	ADDV $16, R0
	MOVV R0, sp+0(FP)
	RET

TEXT ·SetSP(SB),NOSPLIT,$8-8
	MOVV sp+0(FP), R29
	RET

// Called during function prolog when more stack is needed.
// Caller has already loaded:
// R1: framesize, R2: argsize, R3: LR
//
// The traceback routines see morestack on a g0 as being
// the top of a stack (for example, morestack calling newstack
// calling the scheduler calling newm calling gc), so we must
// record an argument size. For that purpose, it has no arguments.
TEXT runtime·morestack(SB),NOSPLIT|NOFRAME,$0
	// TODO: implement
	UNDEF // crash

// morestack but not preserving ctxt.
TEXT runtime·morestack_noctxt(SB),NOSPLIT|NOFRAME,$0
	// Force SPWRITE. This function doesn't actually write SP,
	// but it is called with a special calling convention where
	// the caller doesn't save LR on stack but passes it as a
	// register (R3), and the unwinder currently doesn't understand.
	// Make it SPWRITE to stop unwinding. (See issue 54332)
	MOVV R29, R29

	MOVV R0, REGCTXT
	JMP runtime·morestack(SB)

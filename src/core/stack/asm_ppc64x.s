// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && (ppc64 || ppc64le)

#include "textflag.h"

// R1 is the stack pointer

TEXT ·GetSP<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVD R1, R3
	RET

TEXT ·SetSP<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVD R3, R1
	RET

// Called during function prolog when more stack is needed.
// Caller has already loaded:
// R3: framesize, R4: argsize, R5: LR
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
	// register (R5), and the unwinder currently doesn't understand.
	// Make it SPWRITE to stop unwinding. (See issue 54332)
	// Use OR R0, R1 instead of MOVD R1, R1 as the MOVD instruction
	// has a special affect on Power8,9,10 by lowering the thread 
	// priority and causing a slowdown in execution time

	OR R0, R1
	MOVD R0, R11
	BR runtime·morestack(SB)

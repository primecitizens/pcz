// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && amd64

#include "textflag.h"

TEXT ·GetSP<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVQ SP, AX
	RET

TEXT ·SetSP<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	NOP SP // tell vet SP changed - stop checking offsets
	MOVQ SP, AX
	RET

TEXT ·GetFP<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVQ BP, AX
	RET

// Called during function prolog when more stack is needed.
//
// The traceback routines see morestack on a g0 as being
// the top of a stack (for example, morestack calling newstack
// calling the scheduler calling newm calling gc), so we must
// record an argument size. For that purpose, it has no arguments.
TEXT runtime·morestack(SB),NOSPLIT|NOFRAME,$0-0
	// TODO: implement
	INT $3 // crash
loop:
	JMP loop

// morestack but not preserving ctxt.
TEXT runtime·morestack_noctxt(SB),NOSPLIT|NOFRAME,$0-0
	MOVL $0, DX
	JMP runtime·morestack(SB)

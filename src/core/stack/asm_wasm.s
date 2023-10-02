// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && wasm

#include "textflag.h"

TEXT ·GetSP(SB),NOSPLIT,$0-8
	Get SP
	I64ExtendI32U
	Set R0
	MOVD R0, ret+0(FP)
	RET

TEXT ·SetSP(SB),NOSPLIT,$8-8
	NOP SP // tell vet SP changed - stop checking offsets
	MOVD sp+0(FP), SP
	RET

// Called during function prolog when more stack is needed.
//
// The traceback routines see morestack on a g0 as being
// the top of a stack (for example, morestack calling newstack
// calling the scheduler calling newm calling gc), so we must
// record an argument size. For that purpose, it has no arguments.
TEXT runtime·morestack(SB),NOSPLIT,$0-0
	// TODO: implement
	UNDEF // crash

// morestack but not preserving ctxt.
TEXT runtime·morestack_noctxt(SB),NOSPLIT,$0
	MOVD $0, CTXT
	JMP runtime·morestack(SB)

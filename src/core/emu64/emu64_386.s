// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "funcdata.h"
#include "textflag.h"

TEXT ·Uint32ToFloat64(SB),NOSPLIT,$8-12
	MOVL a+0(FP), AX
	MOVL AX, 0(SP)
	MOVL $0, 4(SP)
	FMOVV 0(SP), F0
	FMOVDP F0, ret+4(FP)
	RET

TEXT ·Float64ToUint32(SB),NOSPLIT,$12-12
	FMOVD a+0(FP), F0
	FSTCW 0(SP)
	FLDCW ·controlWord64trunc(SB)
	FMOVVP F0, 4(SP)
	FLDCW 0(SP)
	MOVL 4(SP), AX
	MOVL AX, ret+8(FP)
	RET

// _mul64x32(lo64 *uint64, a uint64, b uint32) (hi32 uint32)
// sets *lo64 = low 64 bits of 96-bit product a*b; returns high 32 bits.
TEXT ·_mul64by32(SB), NOSPLIT, $0
	MOVL lo64+0(FP), CX
	MOVL a_lo+4(FP), AX
	MULL b+12(FP)
	MOVL AX, 0(CX)
	MOVL DX, BX
	MOVL a_hi+8(FP), AX
	MULL b+12(FP)
	ADDL AX, BX
	ADCL $0, DX
	MOVL BX, 4(CX)
	MOVL DX, AX
	MOVL AX, hi32+16(FP)
	RET

TEXT ·_div64by32(SB), NOSPLIT, $0
	MOVL r+12(FP), CX
	MOVL a_lo+0(FP), AX
	MOVL a_hi+4(FP), DX
	DIVL b+8(FP)
	MOVL DX, 0(CX)
	MOVL AX, q+16(FP)
	RET

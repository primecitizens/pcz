// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ppc64 || ppc64le

#include "textflag.h"

// gcWriteBarrier informs the GC about heap pointer writes.
//
// gcWriteBarrier does NOT follow the Go ABI. It accepts the
// number of bytes of buffer needed in R29, and returns a pointer
// to the buffer space in R29.
// It clobbers condition codes.
// It does not clobber R0 through R17 (except special registers),
// but may clobber any other register, *including* R31.
TEXT gcWriteBarrier<>(SB),NOSPLIT,$120
	// The standard prologue clobbers R31.
	// We use R18, R19, and R31 as scratch registers.
retry:
	MOVD g_m(g), R18
	MOVD m_p(R18), R18
	MOVD (p_wbBuf+wbBuf_next)(R18), R19
	MOVD (p_wbBuf+wbBuf_end)(R18), R31
	// Increment wbBuf.next position.
	ADD R29, R19
	// Is the buffer full?
	CMPU R31, R19
	BLT flush
	// Commit to the larger buffer.
	MOVD R19, (p_wbBuf+wbBuf_next)(R18)
	// Make return value (the original next position)
	SUB R29, R19, R29
	RET

flush:
	// Save registers R0 through R15 since these were not saved by the caller.
	// We don't save all registers on ppc64 because it takes too much space.
	MOVD R20, (FIXED_FRAME+0)(R1)
	MOVD R21, (FIXED_FRAME+8)(R1)
	// R0 is always 0, so no need to spill.
	// R1 is SP.
	// R2 is SB.
	MOVD R3, (FIXED_FRAME+16)(R1)
	MOVD R4, (FIXED_FRAME+24)(R1)
	MOVD R5, (FIXED_FRAME+32)(R1)
	MOVD R6, (FIXED_FRAME+40)(R1)
	MOVD R7, (FIXED_FRAME+48)(R1)
	MOVD R8, (FIXED_FRAME+56)(R1)
	MOVD R9, (FIXED_FRAME+64)(R1)
	MOVD R10, (FIXED_FRAME+72)(R1)
	// R11, R12 may be clobbered by external-linker-inserted trampoline
	// R13 is REGTLS
	MOVD R14, (FIXED_FRAME+80)(R1)
	MOVD R15, (FIXED_FRAME+88)(R1)
	MOVD R16, (FIXED_FRAME+96)(R1)
	MOVD R17, (FIXED_FRAME+104)(R1)
	MOVD R29, (FIXED_FRAME+112)(R1)

	CALL runtime·wbBufFlush(SB)

	MOVD (FIXED_FRAME+0)(R1), R20
	MOVD (FIXED_FRAME+8)(R1), R21
	MOVD (FIXED_FRAME+16)(R1), R3
	MOVD (FIXED_FRAME+24)(R1), R4
	MOVD (FIXED_FRAME+32)(R1), R5
	MOVD (FIXED_FRAME+40)(R1), R6
	MOVD (FIXED_FRAME+48)(R1), R7
	MOVD (FIXED_FRAME+56)(R1), R8
	MOVD (FIXED_FRAME+64)(R1), R9
	MOVD (FIXED_FRAME+72)(R1), R10
	MOVD (FIXED_FRAME+80)(R1), R14
	MOVD (FIXED_FRAME+88)(R1), R15
	MOVD (FIXED_FRAME+96)(R1), R16
	MOVD (FIXED_FRAME+104)(R1), R17
	MOVD (FIXED_FRAME+112)(R1), R29
	JMP retry

TEXT runtime·gcWriteBarrier1<ABIInternal>(SB),NOSPLIT,$0
	MOVD $8, R29
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier2<ABIInternal>(SB),NOSPLIT,$0
	MOVD $16, R29
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier3<ABIInternal>(SB),NOSPLIT,$0
	MOVD $24, R29
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier4<ABIInternal>(SB),NOSPLIT,$0
	MOVD $32, R29
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier5<ABIInternal>(SB),NOSPLIT,$0
	MOVD $40, R29
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier6<ABIInternal>(SB),NOSPLIT,$0
	MOVD $48, R29
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier7<ABIInternal>(SB),NOSPLIT,$0
	MOVD $56, R29
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier8<ABIInternal>(SB),NOSPLIT,$0
	MOVD $64, R29
	JMP gcWriteBarrier<>(SB)

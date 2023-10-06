// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build mips || mipsle

#include "textflag.h"

// gcWriteBarrier informs the GC about heap pointer writes.
//
// gcWriteBarrier does NOT follow the Go ABI. It accepts the
// number of bytes of buffer needed in R25, and returns a pointer
// to the buffer space in R25.
// It clobbers R23 (the linker temp register).
// The act of CALLing gcWriteBarrier will clobber R31 (LR).
// It does not clobber any other general-purpose registers,
// but may clobber others (e.g., floating point registers).
TEXT gcWriteBarrier<>(SB),NOSPLIT,$104
	// Save the registers clobbered by the fast path.
	MOVW R1, 100(R29)
	MOVW R2, 104(R29)
retry:
	MOVW g_m(g), R1
	MOVW m_p(R1), R1
	MOVW (p_wbBuf+wbBuf_next)(R1), R2
	MOVW (p_wbBuf+wbBuf_end)(R1), R23 // R23 is linker temp register
	// Increment wbBuf.next position.
	ADD R25, R2
	// Is the buffer full?
	SGTU R2, R23, R23
	BNE R23, flush
	// Commit to the larger buffer.
	MOVW R2, (p_wbBuf+wbBuf_next)(R1)
	// Make return value (the original next position)
	SUB R25, R2, R25
	// Restore registers.
	MOVW 100(R29), R1
	MOVW 104(R29), R2
	RET

flush:
	// Save all general purpose registers since these could be
	// clobbered by wbBufFlush and were not saved by the caller.
	MOVW R20, 4(R29)
	MOVW R21, 8(R29)
	// R1 already saved
	// R2 already saved
	MOVW R3, 12(R29)
	MOVW R4, 16(R29)
	MOVW R5, 20(R29)
	MOVW R6, 24(R29)
	MOVW R7, 28(R29)
	MOVW R8, 32(R29)
	MOVW R9, 36(R29)
	MOVW R10, 40(R29)
	MOVW R11, 44(R29)
	MOVW R12, 48(R29)
	MOVW R13, 52(R29)
	MOVW R14, 56(R29)
	MOVW R15, 60(R29)
	MOVW R16, 64(R29)
	MOVW R17, 68(R29)
	MOVW R18, 72(R29)
	MOVW R19, 76(R29)
	MOVW R20, 80(R29)
	// R21 already saved
	// R22 already saved.
	MOVW R22, 84(R29)
	// R23 is tmp register.
	MOVW R24, 88(R29)
	MOVW R25, 92(R29)
	// R26 is reserved by kernel.
	// R27 is reserved by kernel.
	MOVW R28, 96(R29)
	// R29 is SP.
	// R30 is g.
	// R31 is LR, which was saved by the prologue.

	CALL runtime·wbBufFlush(SB)

	MOVW 4(R29), R20
	MOVW 8(R29), R21
	MOVW 12(R29), R3
	MOVW 16(R29), R4
	MOVW 20(R29), R5
	MOVW 24(R29), R6
	MOVW 28(R29), R7
	MOVW 32(R29), R8
	MOVW 36(R29), R9
	MOVW 40(R29), R10
	MOVW 44(R29), R11
	MOVW 48(R29), R12
	MOVW 52(R29), R13
	MOVW 56(R29), R14
	MOVW 60(R29), R15
	MOVW 64(R29), R16
	MOVW 68(R29), R17
	MOVW 72(R29), R18
	MOVW 76(R29), R19
	MOVW 80(R29), R20
	MOVW 84(R29), R22
	MOVW 88(R29), R24
	MOVW 92(R29), R25
	MOVW 96(R29), R28
	JMP retry

TEXT runtime·gcWriteBarrier1<ABIInternal>(SB),NOSPLIT,$0
	MOVW $4, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier2<ABIInternal>(SB),NOSPLIT,$0
	MOVW $8, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier3<ABIInternal>(SB),NOSPLIT,$0
	MOVW $12, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier4<ABIInternal>(SB),NOSPLIT,$0
	MOVW $16, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier5<ABIInternal>(SB),NOSPLIT,$0
	MOVW $20, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier6<ABIInternal>(SB),NOSPLIT,$0
	MOVW $24, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier7<ABIInternal>(SB),NOSPLIT,$0
	MOVW $28, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier8<ABIInternal>(SB),NOSPLIT,$0
	MOVW $32, R25
	JMP gcWriteBarrier<>(SB)

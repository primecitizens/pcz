// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build mips64 || mips64le

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
TEXT gcWriteBarrier<>(SB),NOSPLIT,$192
	// Save the registers clobbered by the fast path.
	MOVV R1, 184(R29)
	MOVV R2, 192(R29)
retry:
	MOVV g_m(g), R1
	MOVV m_p(R1), R1
	MOVV (p_wbBuf+wbBuf_next)(R1), R2
	MOVV (p_wbBuf+wbBuf_end)(R1), R23 // R23 is linker temp register
	// Increment wbBuf.next position.
	ADDV R25, R2
	// Is the buffer full?
	SGTU R2, R23, R23
	BNE R23, flush
	// Commit to the larger buffer.
	MOVV R2, (p_wbBuf+wbBuf_next)(R1)
	// Make return value (the original next position)
	SUBV R25, R2, R25
	// Restore registers.
	MOVV 184(R29), R1
	MOVV 192(R29), R2
	RET

flush:
	// Save all general purpose registers since these could be
	// clobbered by wbBufFlush and were not saved by the caller.
	MOVV R20, 8(R29)
	MOVV R21, 16(R29)
	// R1 already saved
	// R2 already saved
	MOVV R3, 24(R29)
	MOVV R4, 32(R29)
	MOVV R5, 40(R29)
	MOVV R6, 48(R29)
	MOVV R7, 56(R29)
	MOVV R8, 64(R29)
	MOVV R9, 72(R29)
	MOVV R10, 80(R29)
	MOVV R11, 88(R29)
	MOVV R12, 96(R29)
	MOVV R13, 104(R29)
	MOVV R14, 112(R29)
	MOVV R15, 120(R29)
	MOVV R16, 128(R29)
	MOVV R17, 136(R29)
	MOVV R18, 144(R29)
	MOVV R19, 152(R29)
	// R20 already saved
	// R21 already saved.
	MOVV R22, 160(R29)
	// R23 is tmp register.
	MOVV R24, 168(R29)
	MOVV R25, 176(R29)
	// R26 is reserved by kernel.
	// R27 is reserved by kernel.
	// R28 is REGSB (not modified by Go code).
	// R29 is SP.
	// R30 is g.
	// R31 is LR, which was saved by the prologue.

	CALL runtime·wbBufFlush(SB)

	MOVV 8(R29), R20
	MOVV 16(R29), R21
	MOVV 24(R29), R3
	MOVV 32(R29), R4
	MOVV 40(R29), R5
	MOVV 48(R29), R6
	MOVV 56(R29), R7
	MOVV 64(R29), R8
	MOVV 72(R29), R9
	MOVV 80(R29), R10
	MOVV 88(R29), R11
	MOVV 96(R29), R12
	MOVV 104(R29), R13
	MOVV 112(R29), R14
	MOVV 120(R29), R15
	MOVV 128(R29), R16
	MOVV 136(R29), R17
	MOVV 144(R29), R18
	MOVV 152(R29), R19
	MOVV 160(R29), R22
	MOVV 168(R29), R24
	MOVV 176(R29), R25
	JMP retry

TEXT runtime·gcWriteBarrier1<ABIInternal>(SB),NOSPLIT,$0
	MOVV $8, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier2<ABIInternal>(SB),NOSPLIT,$0
	MOVV $16, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier3<ABIInternal>(SB),NOSPLIT,$0
	MOVV $24, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier4<ABIInternal>(SB),NOSPLIT,$0
	MOVV $32, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier5<ABIInternal>(SB),NOSPLIT,$0
	MOVV $40, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier6<ABIInternal>(SB),NOSPLIT,$0
	MOVV $48, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier7<ABIInternal>(SB),NOSPLIT,$0
	MOVV $56, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier8<ABIInternal>(SB),NOSPLIT,$0
	MOVV $64, R25
	JMP gcWriteBarrier<>(SB)

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm64

#include "textflag.h"

// gcWriteBarrier informs the GC about heap pointer writes.
//
// gcWriteBarrier does NOT follow the Go ABI. It accepts the
// number of bytes of buffer needed in R25, and returns a pointer
// to the buffer space in R25.
// It clobbers condition codes.
// It does not clobber any general-purpose registers except R27,
// but may clobber others (e.g., floating point registers)
// The act of CALLing gcWriteBarrier will clobber R30 (LR).
TEXT gcWriteBarrier<>(SB),NOSPLIT,$200
	// Save the registers clobbered by the fast path.
	STP (R0, R1), 184(RSP)
retry:
	MOVD g_m(g), R0
	MOVD m_p(R0), R0
	MOVD (p_wbBuf+wbBuf_next)(R0), R1
	MOVD (p_wbBuf+wbBuf_end)(R0), R27
	// Increment wbBuf.next position.
	ADD R25, R1
	// Is the buffer full?
	CMP R27, R1
	BHI flush
	// Commit to the larger buffer.
	MOVD R1, (p_wbBuf+wbBuf_next)(R0)
	// Make return value (the original next position)
	SUB R25, R1, R25
	// Restore registers.
	LDP 184(RSP), (R0, R1)
	RET

flush:
	// Save all general purpose registers since these could be
	// clobbered by wbBufFlush and were not saved by the caller.
	// R0 and R1 already saved
	STP (R2, R3), 1*8(RSP)
	STP (R4, R5), 3*8(RSP)
	STP (R6, R7), 5*8(RSP)
	STP (R8, R9), 7*8(RSP)
	STP (R10, R11), 9*8(RSP)
	STP (R12, R13), 11*8(RSP)
	STP (R14, R15), 13*8(RSP)
	// R16, R17 may be clobbered by linker trampoline
	// R18 is unused.
	STP (R19, R20), 15*8(RSP)
	STP (R21, R22), 17*8(RSP)
	STP (R23, R24), 19*8(RSP)
	STP (R25, R26), 21*8(RSP)
	// R27 is temp register.
	// R28 is g.
	// R29 is frame pointer (unused).
	// R30 is LR, which was saved by the prologue.
	// R31 is SP.

	CALL runtime·wbBufFlush(SB)
	LDP 1*8(RSP), (R2, R3)
	LDP 3*8(RSP), (R4, R5)
	LDP 5*8(RSP), (R6, R7)
	LDP 7*8(RSP), (R8, R9)
	LDP 9*8(RSP), (R10, R11)
	LDP 11*8(RSP), (R12, R13)
	LDP 13*8(RSP), (R14, R15)
	LDP 15*8(RSP), (R19, R20)
	LDP 17*8(RSP), (R21, R22)
	LDP 19*8(RSP), (R23, R24)
	LDP 21*8(RSP), (R25, R26)
	JMP retry

TEXT runtime·gcWriteBarrier1<ABIInternal>(SB),NOSPLIT,$0
	MOVD $8, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier2<ABIInternal>(SB),NOSPLIT,$0
	MOVD $16, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier3<ABIInternal>(SB),NOSPLIT,$0
	MOVD $24, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier4<ABIInternal>(SB),NOSPLIT,$0
	MOVD $32, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier5<ABIInternal>(SB),NOSPLIT,$0
	MOVD $40, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier6<ABIInternal>(SB),NOSPLIT,$0
	MOVD $48, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier7<ABIInternal>(SB),NOSPLIT,$0
	MOVD $56, R25
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier8<ABIInternal>(SB),NOSPLIT,$0
	MOVD $64, R25
	JMP gcWriteBarrier<>(SB)

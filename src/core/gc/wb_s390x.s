// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// gcWriteBarrier informs the GC about heap pointer writes.
//
// gcWriteBarrier does NOT follow the Go ABI. It accepts the
// number of bytes of buffer needed in R9, and returns a pointer
// to the buffer space in R9.
// It clobbers R10 (the temp register) and R1 (used by PLT stub).
// It does not clobber any other general-purpose registers,
// but may clobber others (e.g., floating point registers).
TEXT gcWriteBarrier<>(SB),NOSPLIT,$96
	// Save the registers clobbered by the fast path.
	MOVD R4, 96(R15)
retry:
	MOVD g_m(g), R1
	MOVD m_p(R1), R1
	// Increment wbBuf.next position.
	MOVD R9, R4
	ADD (p_wbBuf+wbBuf_next)(R1), R4
	// Is the buffer full?
	MOVD (p_wbBuf+wbBuf_end)(R1), R10
	CMPUBGT R4, R10, flush
	// Commit to the larger buffer.
	MOVD R4, (p_wbBuf+wbBuf_next)(R1)
	// Make return value (the original next position)
	SUB R9, R4, R9
	// Restore registers.
	MOVD 96(R15), R4
	RET

flush:
	// Save all general purpose registers since these could be
	// clobbered by wbBufFlush and were not saved by the caller.
	STMG R2, R3, 8(R15)
	MOVD R0, 24(R15)
	// R1 already saved.
	// R4 already saved.
	STMG R5, R12, 32(R15) // save R5 - R12
	// R13 is g.
	// R14 is LR.
	// R15 is SP.

	CALL runtime·wbBufFlush(SB)

	LMG 8(R15), R2, R3   // restore R2 - R3
	MOVD 24(R15), R0      // restore R0
	LMG 32(R15), R5, R12 // restore R5 - R12
	JMP retry

TEXT runtime·gcWriteBarrier1<ABIInternal>(SB),NOSPLIT,$0
	MOVD $8, R9
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier2<ABIInternal>(SB),NOSPLIT,$0
	MOVD $16, R9
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier3<ABIInternal>(SB),NOSPLIT,$0
	MOVD $24, R9
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier4<ABIInternal>(SB),NOSPLIT,$0
	MOVD $32, R9
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier5<ABIInternal>(SB),NOSPLIT,$0
	MOVD $40, R9
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier6<ABIInternal>(SB),NOSPLIT,$0
	MOVD $48, R9
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier7<ABIInternal>(SB),NOSPLIT,$0
	MOVD $56, R9
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier8<ABIInternal>(SB),NOSPLIT,$0
	MOVD $64, R9
	JMP gcWriteBarrier<>(SB)

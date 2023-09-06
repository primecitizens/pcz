// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// gcWriteBarrier informs the GC about heap pointer writes.
//
// gcWriteBarrier does NOT follow the Go ABI. It accepts the
// number of bytes of buffer needed in R8, and returns a pointer
// to the buffer space in R8.
// It clobbers condition codes.
// It does not clobber any other general-purpose registers,
// but may clobber others (e.g., floating point registers).
// The act of CALLing gcWriteBarrier will clobber R14 (LR).
TEXT gcWriteBarrier<>(SB),NOSPLIT|NOFRAME,$0
	// Save the registers clobbered by the fast path.
	MOVM.DB.W [R0,R1], (R13)
retry:
	MOVW g_m(g), R0
	MOVW m_p(R0), R0
	MOVW (p_wbBuf+wbBuf_next)(R0), R1
	MOVW (p_wbBuf+wbBuf_end)(R0), R11
	// Increment wbBuf.next position.
	ADD R8, R1
	// Is the buffer full?
	CMP R11, R1
	BHI flush
	// Commit to the larger buffer.
	MOVW R1, (p_wbBuf+wbBuf_next)(R0)
	// Make return value (the original next position)
	SUB R8, R1, R8
	// Restore registers.
	MOVM.IA.W (R13), [R0,R1]
	RET

flush:
	// Save all general purpose registers since these could be
	// clobbered by wbBufFlush and were not saved by the caller.
	//
	// R0 and R1 were saved at entry.
	// R10 is g, so preserved.
	// R11 is linker temp, so no need to save.
	// R13 is stack pointer.
	// R15 is PC.
	MOVM.DB.W [R2-R9,R12], (R13)
	// Save R14 (LR) because the fast path above doesn't save it,
	// but needs it to RET.
	MOVM.DB.W [R14], (R13)

	CALL runtime·wbBufFlush(SB)

	MOVM.IA.W (R13), [R14]
	MOVM.IA.W (R13), [R2-R9,R12]
	JMP retry

TEXT runtime·gcWriteBarrier1<ABIInternal>(SB),NOSPLIT,$0
	MOVW $4, R8
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier2<ABIInternal>(SB),NOSPLIT,$0
	MOVW $8, R8
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier3<ABIInternal>(SB),NOSPLIT,$0
	MOVW $12, R8
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier4<ABIInternal>(SB),NOSPLIT,$0
	MOVW $16, R8
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier5<ABIInternal>(SB),NOSPLIT,$0
	MOVW $20, R8
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier6<ABIInternal>(SB),NOSPLIT,$0
	MOVW $24, R8
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier7<ABIInternal>(SB),NOSPLIT,$0
	MOVW $28, R8
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier8<ABIInternal>(SB),NOSPLIT,$0
	MOVW $32, R8
	JMP gcWriteBarrier<>(SB)

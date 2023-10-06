// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build amd64

#include "textflag.h"

// gcWriteBarrier informs the GC about heap pointer writes.
//
// gcWriteBarrier returns space in a write barrier buffer which
// should be filled in by the caller.
// gcWriteBarrier does NOT follow the Go ABI. It accepts the
// number of bytes of buffer needed in R11, and returns a pointer
// to the buffer space in R11.
// It clobbers FLAGS. It does not clobber any general-purpose registers,
// but may clobber others (e.g., SSE registers).
// Typical use would be, when doing *(CX+88) = AX
//     CMPL    $0, runtime.writeBarrier(SB)
//     JEQ     dowrite
//     CALL    runtime.gcBatchBarrier2(SB)
//     MOVQ    AX, (R11)
//     MOVQ    88(CX), DX
//     MOVQ    DX, 8(R11)
// dowrite:
//     MOVQ    AX, 88(CX)
TEXT gcWriteBarrier<>(SB),NOSPLIT,$112
	// Save the registers clobbered by the fast path. This is slightly
	// faster than having the caller spill these.
	MOVQ R12, 96(SP)
	MOVQ R13, 104(SP)
retry:
	// TODO: Consider passing g.m.p in as an argument so they can be shared
	// across a sequence of write barriers.
	MOVQ g_m(R14), R13
	MOVQ m_p(R13), R13
	// Get current buffer write position.
	MOVQ (p_wbBuf+wbBuf_next)(R13), R12 // original next position
	ADDQ R11, R12 // new next position
	// Is the buffer full?
	CMPQ R12, (p_wbBuf+wbBuf_end)(R13)
	JA flush
	// Commit to the larger buffer.
	MOVQ R12, (p_wbBuf+wbBuf_next)(R13)
	// Make return value (the original next position)
	SUBQ R11, R12
	MOVQ R12, R11
	// Restore registers.
	MOVQ 96(SP), R12
	MOVQ 104(SP), R13
	RET

flush:
	// Save all general purpose registers since these could be
	// clobbered by wbBufFlush and were not saved by the caller.
	// It is possible for wbBufFlush to clobber other registers
	// (e.g., SSE registers), but the compiler takes care of saving
	// those in the caller if necessary. This strikes a balance
	// with registers that are likely to be used.
	//
	// We don't have type information for these, but all code under
	// here is NOSPLIT, so nothing will observe these.
	//
	// TODO: We could strike a different balance; e.g., saving X0
	// and not saving GP registers that are less likely to be used.
	MOVQ DI, 0(SP)
	MOVQ AX, 8(SP)
	MOVQ BX, 16(SP)
	MOVQ CX, 24(SP)
	MOVQ DX, 32(SP)
	// DI already saved
	MOVQ SI, 40(SP)
	MOVQ BP, 48(SP)
	MOVQ R8, 56(SP)
	MOVQ R9, 64(SP)
	MOVQ R10, 72(SP)
	MOVQ R11, 80(SP)
	// R12 already saved
	// R13 already saved
	// R14 is g
	MOVQ R15, 88(SP)

	CALL runtime·wbBufFlush(SB)

	MOVQ 0(SP), DI
	MOVQ 8(SP), AX
	MOVQ 16(SP), BX
	MOVQ 24(SP), CX
	MOVQ 32(SP), DX
	MOVQ 40(SP), SI
	MOVQ 48(SP), BP
	MOVQ 56(SP), R8
	MOVQ 64(SP), R9
	MOVQ 72(SP), R10
	MOVQ 80(SP), R11
	MOVQ 88(SP), R15
	JMP retry

TEXT runtime·gcWriteBarrier1<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVL   $8, R11
	JMP     gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier2<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVL   $16, R11
	JMP     gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier3<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVL   $24, R11
	JMP     gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier4<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVL   $32, R11
	JMP     gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier5<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVL   $40, R11
	JMP     gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier6<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVL   $48, R11
	JMP     gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier7<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVL   $56, R11
	JMP     gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier8<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVL   $64, R11
	JMP     gcWriteBarrier<>(SB)

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386

#include "textflag.h"

// gcWriteBarrier informs the GC about heap pointer writes.
//
// gcWriteBarrier returns space in a write barrier buffer which
// should be filled in by the caller.
// gcWriteBarrier does NOT follow the Go ABI. It accepts the
// number of bytes of buffer needed in DI, and returns a pointer
// to the buffer space in DI.
// It clobbers FLAGS. It does not clobber any general-purpose registers,
// but may clobber others (e.g., SSE registers).
// Typical use would be, when doing *(CX+88) = AX
//     CMPL    $0, runtime.writeBarrier(SB)
//     JEQ     dowrite
//     CALL    runtime.gcBatchBarrier2(SB)
//     MOVL    AX, (DI)
//     MOVL    88(CX), DX
//     MOVL    DX, 4(DI)
// dowrite:
//     MOVL    AX, 88(CX)
TEXT gcWriteBarrier<>(SB),NOSPLIT,$28
	// Save the registers clobbered by the fast path. This is slightly
	// faster than having the caller spill these.
	MOVL CX, 20(SP)
	MOVL BX, 24(SP)
retry:
	// TODO: Consider passing g.m.p in as an argument so they can be shared
	// across a sequence of write barriers.
	get_tls(BX)
	MOVL g(BX), BX
	MOVL g_m(BX), BX
	MOVL m_p(BX), BX
	// Get current buffer write position.
	MOVL (p_wbBuf+wbBuf_next)(BX), CX // original next position
	ADDL DI, CX // new next position
	// Is the buffer full?
	CMPL CX, (p_wbBuf+wbBuf_end)(BX)
	JA flush
	// Commit to the larger buffer.
	MOVL CX, (p_wbBuf+wbBuf_next)(BX)
	// Make return value (the original next position)
	SUBL DI, CX
	MOVL CX, DI
	// Restore registers.
	MOVL 20(SP), CX
	MOVL 24(SP), BX
	RET

flush:
	// Save all general purpose registers since these could be
	// clobbered by wbBufFlush and were not saved by the caller.
	MOVL DI, 0(SP)
	MOVL AX, 4(SP)
	// BX already saved
	// CX already saved
	MOVL DX, 8(SP)
	MOVL BP, 12(SP)
	MOVL SI, 16(SP)
	// DI already saved

	CALL runtime·wbBufFlush(SB)

	MOVL 0(SP), DI
	MOVL 4(SP), AX
	MOVL 8(SP), DX
	MOVL 12(SP), BP
	MOVL 16(SP), SI
	JMP retry

TEXT runtime·gcWriteBarrier1<ABIInternal>(SB),NOSPLIT,$0
	MOVL $4, DI
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier2<ABIInternal>(SB),NOSPLIT,$0
	MOVL $8, DI
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier3<ABIInternal>(SB),NOSPLIT,$0
	MOVL $12, DI
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier4<ABIInternal>(SB),NOSPLIT,$0
	MOVL $16, DI
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier5<ABIInternal>(SB),NOSPLIT,$0
	MOVL $20, DI
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier6<ABIInternal>(SB),NOSPLIT,$0
	MOVL $24, DI
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier7<ABIInternal>(SB),NOSPLIT,$0
	MOVL $28, DI
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier8<ABIInternal>(SB),NOSPLIT,$0
	MOVL $32, DI
	JMP gcWriteBarrier<>(SB)

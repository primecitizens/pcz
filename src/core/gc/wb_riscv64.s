// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build riscv64

#include "textflag.h"

// gcWriteBarrier informs the GC about heap pointer writes.
//
// gcWriteBarrier does NOT follow the Go ABI. It accepts the
// number of bytes of buffer needed in X24, and returns a pointer
// to the buffer space in X24.
// It clobbers X31 aka T6 (the linker temp register - REG_TMP).
// The act of CALLing gcWriteBarrier will clobber RA (LR).
// It does not clobber any other general-purpose registers,
// but may clobber others (e.g., floating point registers).
TEXT gcWriteBarrier<>(SB),NOSPLIT,$208
	// Save the registers clobbered by the fast path.
	MOV A0, 24*8(X2)
	MOV A1, 25*8(X2)
retry:
	MOV g_m(g), A0
	MOV m_p(A0), A0
	MOV (p_wbBuf+wbBuf_next)(A0), A1
	MOV (p_wbBuf+wbBuf_end)(A0), T6 // T6 is linker temp register (REG_TMP)
	// Increment wbBuf.next position.
	ADD X24, A1
	// Is the buffer full?
	BLTU T6, A1, flush
	// Commit to the larger buffer.
	MOV A1, (p_wbBuf+wbBuf_next)(A0)
	// Make the return value (the original next position)
	SUB X24, A1, X24
	// Restore registers.
	MOV 24*8(X2), A0
	MOV 25*8(X2), A1
	RET

flush:
	// Save all general purpose registers since these could be
	// clobbered by wbBufFlush and were not saved by the caller.
	MOV T0, 1*8(X2)
	MOV T1, 2*8(X2)
	// X0 is zero register
	// X1 is LR, saved by prologue
	// X2 is SP
	// X3 is GP
	// X4 is TP
	MOV X7, 3*8(X2)
	MOV X8, 4*8(X2)
	MOV X9, 5*8(X2)
	// X10 already saved (A0)
	// X11 already saved (A1)
	MOV X12, 6*8(X2)
	MOV X13, 7*8(X2)
	MOV X14, 8*8(X2)
	MOV X15, 9*8(X2)
	MOV X16, 10*8(X2)
	MOV X17, 11*8(X2)
	MOV X18, 12*8(X2)
	MOV X19, 13*8(X2)
	MOV X20, 14*8(X2)
	MOV X21, 15*8(X2)
	MOV X22, 16*8(X2)
	MOV X23, 17*8(X2)
	MOV X24, 18*8(X2)
	MOV X25, 19*8(X2)
	MOV X26, 20*8(X2)
	// X27 is g.
	MOV X28, 21*8(X2)
	MOV X29, 22*8(X2)
	MOV X30, 23*8(X2)
	// X31 is tmp register.

	CALL runtime·wbBufFlush(SB)

	MOV 1*8(X2), T0
	MOV 2*8(X2), T1
	MOV 3*8(X2), X7
	MOV 4*8(X2), X8
	MOV 5*8(X2), X9
	MOV 6*8(X2), X12
	MOV 7*8(X2), X13
	MOV 8*8(X2), X14
	MOV 9*8(X2), X15
	MOV 10*8(X2), X16
	MOV 11*8(X2), X17
	MOV 12*8(X2), X18
	MOV 13*8(X2), X19
	MOV 14*8(X2), X20
	MOV 15*8(X2), X21
	MOV 16*8(X2), X22
	MOV 17*8(X2), X23
	MOV 18*8(X2), X24
	MOV 19*8(X2), X25
	MOV 20*8(X2), X26
	MOV 21*8(X2), X28
	MOV 22*8(X2), X29
	MOV 23*8(X2), X30

	JMP retry

TEXT runtime·gcWriteBarrier1<ABIInternal>(SB),NOSPLIT,$0
	MOV $8, X24
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier2<ABIInternal>(SB),NOSPLIT,$0
	MOV $16, X24
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier3<ABIInternal>(SB),NOSPLIT,$0
	MOV $24, X24
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier4<ABIInternal>(SB),NOSPLIT,$0
	MOV $32, X24
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier5<ABIInternal>(SB),NOSPLIT,$0
	MOV $40, X24
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier6<ABIInternal>(SB),NOSPLIT,$0
	MOV $48, X24
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier7<ABIInternal>(SB),NOSPLIT,$0
	MOV $56, X24
	JMP gcWriteBarrier<>(SB)
TEXT runtime·gcWriteBarrier8<ABIInternal>(SB),NOSPLIT,$0
	MOV $64, X24
	JMP gcWriteBarrier<>(SB)

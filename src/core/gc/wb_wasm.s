// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasm

#include "textflag.h"

// gcWriteBarrier informs the GC about heap pointer writes.
//
// gcWriteBarrier does NOT follow the Go ABI. It accepts the
// number of bytes of buffer needed as a wasm argument
// (put on the TOS by the caller, lives in local R0 in this body)
// and returns a pointer to the buffer space as a wasm result
// (left on the TOS in this body, appears on the wasm stack
// in the caller).
TEXT gcWriteBarrier<>(SB), NOSPLIT, $0
	Loop
		// R3 = g.m
		MOVD g_m(g), R3
		// R4 = p
		MOVD m_p(R3), R4
		// R5 = wbBuf.next
		MOVD p_wbBuf+wbBuf_next(R4), R5

		// Increment wbBuf.next
		Get R5
		Get R0
		I64Add
		Set R5

		// Is the buffer full?
		Get R5
		I64Load (p_wbBuf+wbBuf_end)(R4)
		I64LeU
		If
			// Commit to the larger buffer.
			MOVD R5, p_wbBuf+wbBuf_next(R4)

			// Make return value (the original next position)
			Get R5
			Get R0
			I64Sub

			Return
		End

		// Flush
		CALLNORESUME runtime·wbBufFlush(SB)

		// Retry
		Br $0
	End

TEXT runtime·gcWriteBarrier1<ABIInternal>(SB),NOSPLIT,$0
	I64Const $8
	Call gcWriteBarrier<>(SB)
	Return
TEXT runtime·gcWriteBarrier2<ABIInternal>(SB),NOSPLIT,$0
	I64Const $16
	Call gcWriteBarrier<>(SB)
	Return
TEXT runtime·gcWriteBarrier3<ABIInternal>(SB),NOSPLIT,$0
	I64Const $24
	Call gcWriteBarrier<>(SB)
	Return
TEXT runtime·gcWriteBarrier4<ABIInternal>(SB),NOSPLIT,$0
	I64Const $32
	Call gcWriteBarrier<>(SB)
	Return
TEXT runtime·gcWriteBarrier5<ABIInternal>(SB),NOSPLIT,$0
	I64Const $40
	Call gcWriteBarrier<>(SB)
	Return
TEXT runtime·gcWriteBarrier6<ABIInternal>(SB),NOSPLIT,$0
	I64Const $48
	Call gcWriteBarrier<>(SB)
	Return
TEXT runtime·gcWriteBarrier7<ABIInternal>(SB),NOSPLIT,$0
	I64Const $56
	Call gcWriteBarrier<>(SB)
	Return
TEXT runtime·gcWriteBarrier8<ABIInternal>(SB),NOSPLIT,$0
	I64Const $64
	Call gcWriteBarrier<>(SB)
	Return

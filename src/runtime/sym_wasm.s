// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT wasm_pc_f_loop(SB),NOSPLIT,$0
// Call the function for the current PC_F. Repeat until PAUSE != 0 indicates pause or exit.
// The WebAssembly stack may unwind, e.g. when switching goroutines.
// The Go stack on the linear memory is then used to jump to the correct functions
// with this loop, without having to restore the full WebAssembly stack.
// It is expected to have a pending call before entering the loop, so check PAUSE first.
	Get PAUSE
	I32Eqz
	If
	loop:
	  Loop
	    // Get PC_B & PC_F from -8(SP)
	    Get SP
	    I32Const $8
	    I32Sub
	    I32Load16U $0 // PC_B

	    Get SP
	    I32Const $8
	    I32Sub
	    I32Load16U $2 // PC_F

	    CallIndirect $0
	    Drop

	    Get PAUSE
	    I32Eqz
	    BrIf loop
	  End
	End

	I32Const $0
	Set PAUSE

	Return

TEXT runtime·wasmDiv(SB), NOSPLIT, $0-0
	Get R0
	I64Const $-0x8000000000000000
	I64Eq
	If
	  Get R1
	  I64Const $-1
	  I64Eq
	  If
	    I64Const $-0x8000000000000000
	    Return
	  End
	End
	Get R0
	Get R1
	I64DivS
	Return

TEXT runtime·wasmTruncS(SB), NOSPLIT, $0-0
	Get R0
	Get R0
	F64Ne // NaN
	If
	  I64Const $0x8000000000000000
	  Return
	End

	Get R0
	F64Const $0x7ffffffffffffc00p0 // Maximum truncated representation of 0x7fffffffffffffff
	F64Gt
	If
	  I64Const $0x8000000000000000
	  Return
	End

	Get R0
	F64Const $-0x7ffffffffffffc00p0 // Minimum truncated representation of -0x8000000000000000
	F64Lt
	If
	  I64Const $0x8000000000000000
	  Return
	End

	Get R0
	I64TruncF64S
	Return

TEXT runtime·wasmTruncU(SB), NOSPLIT, $0-0
	Get R0
	Get R0
	F64Ne // NaN
	If
	  I64Const $0x8000000000000000
	  Return
	End

	Get R0
	F64Const $0xfffffffffffff800p0 // Maximum truncated representation of 0xffffffffffffffff
	F64Gt
	If
	  I64Const $0x8000000000000000
	  Return
	End

	Get R0
	F64Const $0.
	F64Lt
	If
	  I64Const $0x8000000000000000
	  Return
	End

	Get R0
	I64TruncF64U
	Return

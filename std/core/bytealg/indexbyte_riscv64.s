// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
// 
// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && riscv64

#include "textflag.h"

TEXT ·IndexSliceByte<ABIInternal>(SB),NOSPLIT,$0-40
	// X10 = b_base
	// X11 = b_len
	// X12 = b_cap (unused)
	// X13 = byte to find
	AND $0xff, X13
	MOV X10, X12 // store base for later
	ADD X10, X11 // end
	ADD $-1, X10

loop:
	ADD $1, X10
	BEQ X10, X11, notfound
	MOVBU (X10), X14
	BNE X13, X14, loop

	SUB X12, X10 // remove base
	RET

notfound:
	MOV $-1, X10
	RET

TEXT ·IndexByte<ABIInternal>(SB),NOSPLIT,$0-32
	// X10 = b_base
	// X11 = b_len
	// X12 = byte to find
	AND $0xff, X12
	MOV X10, X13 // store base for later
	ADD X10, X11 // end
	ADD $-1, X10

loop:
	ADD $1, X10
	BEQ X10, X11, notfound
	MOVBU (X10), X14
	BNE X12, X14, loop

	SUB X13, X10 // remove base
	RET

notfound:
	MOV $-1, X10
	RET

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·Breakpoint(SB),NOSPLIT,$0-0
	INT $3
	RET

TEXT ·Abort(SB),NOSPLIT,$0-0
	INT $3
loop:
	JMP loop

TEXT ·Return0(SB),NOSPLIT,$0
	MOVL $0, AX
	RET

TEXT ·GetCallerPC(SB),NOSPLIT,$4-8
	MOVL argp+0(FP), AX  // addr of first arg
	MOVL -4(AX), AX      // get calling pc
	MOVL AX, ret+4(FP)
	RET

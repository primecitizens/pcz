// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && amd64

#include "textflag.h"

TEXT ·Breakpoint(SB),NOSPLIT,$0-0
	BYTE $0xcc
	RET

TEXT ·Abort(SB),NOSPLIT,$0-0
	INT $3
loop:
	JMP loop

TEXT ·Return0(SB),NOSPLIT,$0
	MOVL $0, AX
	RET

TEXT ·GetCallerPC<ABIInternal>(SB), NOSPLIT|NOFRAME,$8-8
	MOVQ SP, AX  // addr of first arg
	MOVQ -8(AX), AX      // get calling pc
	RET

TEXT ·GetCallerSP<ABIInternal>(SB), NOSPLIT|NOFRAME, $0-0
	LEAQ sp+8(FP), AX
	RET

TEXT ·GetClosurePtr<ABIInternal>(SB), NOSPLIT|NOFRAME, $0-0
	MOVQ DX, AX
	RET

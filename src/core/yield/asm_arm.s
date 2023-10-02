// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm

#include "textflag.h"

TEXT ·Proc(SB),NOSPLIT|NOFRAME,$0
	MOVW cycles+0(FP), R1
	MOVW $0, R0
LOOP:
	WORD $0xe320f001 // YIELD (NOP pre-ARMv6K)
	CMP R0, R1
	B.NE 2(PC)
	RET
	SUB $1, R1
	B LOOP

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·Breakpoint(SB),NOSPLIT,$0-0
	// gdb won't skip this breakpoint instruction automatically,
	// so you must manually "set $pc+=4" to skip it and continue.
#ifdef GOOS_plan9
	WORD $0xD1200070 // undefined instruction used as armv5 breakpoint in Plan 9
#else
	WORD $0xe7f001f0 // undefined instruction that gdb understands is a software breakpoint
#endif
	RET

TEXT ·Abort(SB),NOSPLIT|NOFRAME,$0-0
	MOVW $0, R0
	MOVW (R0), R1

TEXT ·Return0(SB),NOSPLIT,$0
	MOVW $0, R0
	RET

TEXT ·GetCallerPC(SB),NOSPLIT,$-4-4
	MOVW 0(R13), R0 // LR saved by caller
	MOVW R0, ret+0(FP)
	RET

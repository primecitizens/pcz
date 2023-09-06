// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build mips || mipsle

#include "textflag.h"

TEXT ·Breakpoint(SB),NOSPLIT,$0-0
	BREAK
	RET

TEXT ·Abort(SB),NOSPLIT,$0-0
	UNDEF

TEXT ·Return0(SB),NOSPLIT,$0
	MOVW $0, R1
	RET

TEXT ·GetCallerPC(SB),NOSPLIT,$-4-4
	MOVW 0(R29), R1 // LR saved by caller
	MOVW R1, ret+0(FP)
	RET

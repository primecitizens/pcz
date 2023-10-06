// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && (mips64 || mips64le)

#include "textflag.h"

TEXT ·Breakpoint(SB),NOSPLIT|NOFRAME,$0-0
	MOVV R0, 2(R0) // TODO: TD
	RET

TEXT ·Abort(SB),NOSPLIT|NOFRAME,$0-0
	MOVW (R0), R0
	UNDEF

TEXT ·Return0(SB),NOSPLIT,$0
	MOVW $0, R1
	RET

TEXT ·GetCallerPC(SB),NOSPLIT,$-8-8
	MOVV 0(R29), R1    // LR saved by caller
	MOVV R1, ret+0(FP)
	RET

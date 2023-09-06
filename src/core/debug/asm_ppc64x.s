// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ppc64 || ppc64le

#include "textflag.h"

TEXT ·Breakpoint(SB),NOSPLIT|NOFRAME,$0-0
	TW $31, R0, R0
	RET

TEXT ·Abort(SB),NOSPLIT|NOFRAME,$0-0
	MOVW (R0), R0
	UNDEF

TEXT ·Return0(SB),NOSPLIT,$0
	MOVW $0, R3
	RET

TEXT ·GetCallerPC<ABIInternal>(SB),NOSPLIT|NOFRAME,$0-0
	MOVD 0(R1), R3  // LR saved by caller
	RET

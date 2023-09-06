// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·Breakpoint(SB),NOSPLIT|NOFRAME,$0-0
	EBREAK
	RET

TEXT ·Abort(SB),NOSPLIT|NOFRAME,$0-0
	EBREAK
	RET

TEXT ·Return0(SB),NOSPLIT,$0
	MOV $0, A0
	RET

TEXT ·GetCallerPC<ABIInternal>(SB),NOSPLIT|NOFRAME,$0-0
	MOV 0(X2), X10 // LR saved by caller
	RET

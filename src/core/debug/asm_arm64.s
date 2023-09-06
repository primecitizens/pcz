// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"


TEXT ·Breakpoint(SB),NOSPLIT|NOFRAME,$0-0
// Windows ARM64 needs an immediate 0xf000 argument.
// See go.dev/issues/53837.
#ifdef GOOS_windows
	BRK $0xf000
#else
	BRK
#endif
	RET

TEXT ·Abort(SB),NOSPLIT|NOFRAME,$0-0
	MOVD ZR, R0
	MOVD (R0), R0
	UNDEF

TEXT ·Return0(SB), NOSPLIT, $0
	MOVW $0, R0
	RET

TEXT ·GetCallerPC<ABIInternal>(SB),NOSPLIT|NOFRAME,$-8-0
	MOVD 0(RSP), R0 // LR saved by caller
	RET

TEXT ·GetCallerSP<ABIInternal>(SB), NOSPLIT|NOFRAME, $0-0
	MOVD $callersp-8(FP), R0
	RET (R30)
	
TEXT ·GetClosurePtr<ABIInternal>(SB), NOSPLIT|NOFRAME, $0-0
	MOVD R26, R0
	RET (R30)

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"


TEXT ·GetTLSBaseAddress<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
#ifdef GOOS_darwin
	WORD $0xd53bd060 // MRS TPIDRRO_EL0, R0
	// Darwin sometimes returns unaligned pointers
	AND $0xfffffffffffffff8, R0
#else
	WORD $0xd53bd040 // MRS TPIDR_EL0, R0
#endif
	RET

// func SetG(*stdgo.GHead)
TEXT ·SetG<ABIInternal>(SB),NOSPLIT|NOFRAME,$0
	MOVD R0, g
	RET

TEXT ·Topframe<ABIInternal>(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVD R1, RSP
	JMP (R2)
	RET

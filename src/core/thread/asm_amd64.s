// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

#include "textflag.h"

TEXT ·GetTLSBaseAddress<ABIInternal>(SB),NOSPLIT,$0
	MOVQ TLS, AX

#ifdef GOOS_darwin
	// Darwin sometimes returns unaligned pointers
	ANDQ $0xfffffffffffffff8, AX
#endif

	RET

// func SetG(*stdgo.GHead)
TEXT ·SetG<ABIInternal>(SB),NOSPLIT,$0
	MOVQ AX, R14
	// Initialize TLS entry.
	// See cmd/link/internal/ld/sym.go:computeTLSOffset.
	MOVQ AX, 0x30(GS)
	RET

TEXT ·Topframe<ABIInternal>(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	JMP (CX)
	RET

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && riscv64

#include "textflag.h"

#ifdef GOOS_linux

TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOV 0(X2), A0  // argc
	ADD $8, X2, A1 // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

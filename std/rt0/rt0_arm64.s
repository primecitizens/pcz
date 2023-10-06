// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && arm64

#include "textflag.h"

#ifdef GOOS_linux

TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	MOVD 0(RSP), R0 // argc
	ADD $8, RSP, R1 // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_android

TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	MOVD 0(RSP), R0 // argc
	ADD $8, RSP, R1 // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_openbsd

TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	MOVD 0(RSP), R0 // argc
	ADD $8, RSP, R1 // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_netbsd

TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	MOVD 0(RSP), R0 // argc
	ADD $8, RSP, R1 // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_freebsd

// On FreeBSD argc/argv are passed in R0, not RSP
TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	ADD $8, R0, R1 // argv
	MOVD 0(R0), R0 // argc
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_darwin

TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	// R0: argc
	// R1: argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_ios

TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	// R0: argc
	// R1: argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_windows

TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	MOVD 0(RSP), R0 // argc
	ADD $8, RSP, R1 // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

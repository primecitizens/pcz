// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && amd64

#include "textflag.h"

#ifdef GOOS_linux

TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVQ 0(SP), AX // argc
	LEAQ 8(SP), BX // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_android

TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVQ 0(SP), AX // argc
	LEAQ 8(SP), BX // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_freebsd
// On FreeBSD argc/argv are passed in DI, not SP
TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVQ 0(DI), AX // argc
	LEAQ 8(DI), BX // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_netbsd

TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVQ 0(SP), AX // argc
	LEAQ 8(SP), BX // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_openbsd

TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVQ 0(SP), AX // argc
	LEAQ 8(SP), BX // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_dragonfly
// On Dragonfly argc/argv are passed in DI, not SP
TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVQ 0(DI), AX // argc
	LEAQ 8(DI), BX // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_solaris

TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVQ 0(SP), AX // argc
	LEAQ 8(SP), BX // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_darwin

TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVQ 0(SP), AX // argc
	LEAQ 8(SP), BX // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_ios

TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVQ 0(SP), AX // argc
	MOVQ 8(SP), BX // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

#ifdef GOOS_windows

TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVQ 0(SP), AX // argc
	LEAQ 8(SP), BX // argv
	JMP ·rt0<ABIInternal>(SB)

#endif

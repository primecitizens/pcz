// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && (ppc64 || ppc64le)

#include "textflag.h"
#include "funcdata.h" // for NO_LOCAL_POINTERS

// spillArgs stores return values from registers to a *internal/abi.RegArgs in R20.
TEXT ·spillArgs(SB),NOSPLIT,$0-0
	MOVD R3, 0(R20)
	MOVD R4, 8(R20)
	MOVD R5, 16(R20)
	MOVD R6, 24(R20)
	MOVD R7, 32(R20)
	MOVD R8, 40(R20)
	MOVD R9, 48(R20)
	MOVD R10, 56(R20)
	MOVD R14, 64(R20)
	MOVD R15, 72(R20)
	MOVD R16, 80(R20)
	MOVD R17, 88(R20)
	FMOVD F1, 96(R20)
	FMOVD F2, 104(R20)
	FMOVD F3, 112(R20)
	FMOVD F4, 120(R20)
	FMOVD F5, 128(R20)
	FMOVD F6, 136(R20)
	FMOVD F7, 144(R20)
	FMOVD F8, 152(R20)
	FMOVD F9, 160(R20)
	FMOVD F10, 168(R20)
	FMOVD F11, 176(R20)
	FMOVD F12, 184(R20)
	RET

// unspillArgs loads args into registers from a *internal/abi.RegArgs in R20.
TEXT ·unspillArgs(SB),NOSPLIT,$0-0
	MOVD 0(R20), R3
	MOVD 8(R20), R4
	MOVD 16(R20), R5
	MOVD 24(R20), R6
	MOVD 32(R20), R7
	MOVD 40(R20), R8
	MOVD 48(R20), R9
	MOVD 56(R20), R10
	MOVD 64(R20), R14
	MOVD 72(R20), R15
	MOVD 80(R20), R16
	MOVD 88(R20), R17
	FMOVD 96(R20), F1
	FMOVD 104(R20), F2
	FMOVD 112(R20), F3
	FMOVD 120(R20), F4
	FMOVD 128(R20), F5
	FMOVD 136(R20), F6
	FMOVD 144(R20), F7
	FMOVD 152(R20), F8
	FMOVD 160(R20), F9
	FMOVD 168(R20), F10
	FMOVD 176(R20), F11
	FMOVD 184(R20), F12
	RET

// reflectcall: call a function with the given argument list
// func call(stackArgsType *_type, f *FuncVal, stackArgs *byte, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs).
// we don't have variable-sized frames, so we use a small number
// of constant-sized-frame functions to encode a few bits of size in the pc.
// Caution: ugly multiline assembly macros in your future!

#define DISPATCH(NAME,MAXSIZE)		\
	MOVD $MAXSIZE, R31;		\
	CMP R3, R31;		\
	BGT 4(PC);			\
	MOVD $NAME(SB), R12;		\
	MOVD R12, CTR;		\
	BR (CTR)
// Note: can't just "BR NAME(SB)" - bad inlining results.

TEXT ·Call(SB), NOSPLIT|NOFRAME, $0-48
	MOVWZ frameSize+32(FP), R3
	DISPATCH(·call16, 16)
	DISPATCH(·call32, 32)
	DISPATCH(·call64, 64)
	DISPATCH(·call128, 128)
	DISPATCH(·call256, 256)
	DISPATCH(·call512, 512)
	DISPATCH(·call1024, 1024)
	DISPATCH(·call2048, 2048)
	DISPATCH(·call4096, 4096)
	DISPATCH(·call8192, 8192)
	DISPATCH(·call16384, 16384)
	DISPATCH(·call32768, 32768)
	DISPATCH(·call65536, 65536)
	DISPATCH(·call131072, 131072)
	DISPATCH(·call262144, 262144)
	DISPATCH(·call524288, 524288)
	DISPATCH(·call1048576, 1048576)
	DISPATCH(·call2097152, 2097152)
	DISPATCH(·call4194304, 4194304)
	DISPATCH(·call8388608, 8388608)
	DISPATCH(·call16777216, 16777216)
	DISPATCH(·call33554432, 33554432)
	DISPATCH(·call67108864, 67108864)
	DISPATCH(·call134217728, 134217728)
	DISPATCH(·call268435456, 268435456)
	DISPATCH(·call536870912, 536870912)
	DISPATCH(·call1073741824, 1073741824)
	MOVD $·badreflectcall(SB), R12
	MOVD R12, CTR
	BR (CTR)

#define CALLFN(NAME,MAXSIZE)			\
TEXT NAME(SB), WRAPPER, $MAXSIZE-48;		\
	NO_LOCAL_POINTERS;			\
	/* copy arguments to stack */		\
	MOVD stackArgs+16(FP), R3;			\
	MOVWZ stackArgsSize+24(FP), R4;			\
	MOVD R1, R5;				\
	CMP R4, $8;				\
	BLT tailsetup;			\
	/* copy 8 at a time if possible */	\
	ADD $(FIXED_FRAME-8), R5;			\
	SUB $8, R3;				\
top: \
	MOVDU 8(R3), R7;			\
	MOVDU R7, 8(R5);			\
	SUB $8, R4;				\
	CMP R4, $8;				\
	BGE top;				\
	/* handle remaining bytes */	\
	CMP $0, R4;			\
	BEQ callfn;			\
	ADD $7, R3;			\
	ADD $7, R5;			\
	BR tail;			\
tailsetup: \
	CMP $0, R4;			\
	BEQ callfn;			\
	ADD     $(FIXED_FRAME-1), R5;	\
	SUB     $1, R3;			\
tail: \
	MOVBU 1(R3), R6;		\
	MOVBU R6, 1(R5);		\
	SUB $1, R4;			\
	CMP $0, R4;			\
	BGT tail;			\
callfn: \
	/* call function */			\
	MOVD f+8(FP), R11;			\
#ifdef GOOS_aix \
	/* AIX won't trigger a SIGSEGV if R11 = nil */	\
	/* So it manually triggers it */	\
	CMP R0, R11 \
	BNE 2(PC)				\
	MOVD R0, 0(R0)			\
#endif \
	MOVD regArgs+40(FP), R20;    \
	BL      ·unspillArgs(SB);        \
	MOVD (R11), R12;			\
	MOVD R12, CTR;			\
	PCDATA  $PCDATA_StackMapIndex, $0;	\
	BL (CTR);				\
#ifndef GOOS_aix \
	MOVD 24(R1), R2;			\
#endif \
	/* copy return values back */		\
	MOVD regArgs+40(FP), R20;		\
	BL ·spillArgs(SB);			\
	MOVD stackArgsType+0(FP), R7;		\
	MOVD stackArgs+16(FP), R3;			\
	MOVWZ stackArgsSize+24(FP), R4;			\
	MOVWZ stackRetOffset+28(FP), R6;		\
	ADD $FIXED_FRAME, R1, R5;		\
	ADD R6, R5; 			\
	ADD R6, R3;				\
	SUB R6, R4;				\
	BL callRet<>(SB);			\
	RET

// callRet copies return values back at the end of call*. This is a
// separate function so it can allocate stack space for the arguments
// to reflectcallmove. It does not follow the Go ABI; it expects its
// arguments in registers.
TEXT callRet<>(SB), NOSPLIT, $40-0
	NO_LOCAL_POINTERS
	MOVD R7, FIXED_FRAME+0(R1)
	MOVD R3, FIXED_FRAME+8(R1)
	MOVD R5, FIXED_FRAME+16(R1)
	MOVD R4, FIXED_FRAME+24(R1)
	MOVD R20, FIXED_FRAME+32(R1)
	BL ·reflectcallmove(SB)
	RET

CALLFN(·call16, 16)
CALLFN(·call32, 32)
CALLFN(·call64, 64)
CALLFN(·call128, 128)
CALLFN(·call256, 256)
CALLFN(·call512, 512)
CALLFN(·call1024, 1024)
CALLFN(·call2048, 2048)
CALLFN(·call4096, 4096)
CALLFN(·call8192, 8192)
CALLFN(·call16384, 16384)
CALLFN(·call32768, 32768)
CALLFN(·call65536, 65536)
CALLFN(·call131072, 131072)
CALLFN(·call262144, 262144)
CALLFN(·call524288, 524288)
CALLFN(·call1048576, 1048576)
CALLFN(·call2097152, 2097152)
CALLFN(·call4194304, 4194304)
CALLFN(·call8388608, 8388608)
CALLFN(·call16777216, 16777216)
CALLFN(·call33554432, 33554432)
CALLFN(·call67108864, 67108864)
CALLFN(·call134217728, 134217728)
CALLFN(·call268435456, 268435456)
CALLFN(·call536870912, 536870912)
CALLFN(·call1073741824, 1073741824)

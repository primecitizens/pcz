// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h" // for NO_LOCAL_POINTERS

// spillArgs stores return values from registers to a *internal/abi.RegArgs in R20.
TEXT ·spillArgs(SB),NOSPLIT,$0-0
	STP (R0, R1), (0*8)(R20)
	STP (R2, R3), (2*8)(R20)
	STP (R4, R5), (4*8)(R20)
	STP (R6, R7), (6*8)(R20)
	STP (R8, R9), (8*8)(R20)
	STP (R10, R11), (10*8)(R20)
	STP (R12, R13), (12*8)(R20)
	STP (R14, R15), (14*8)(R20)
	FSTPD (F0, F1), (16*8)(R20)
	FSTPD (F2, F3), (18*8)(R20)
	FSTPD (F4, F5), (20*8)(R20)
	FSTPD (F6, F7), (22*8)(R20)
	FSTPD (F8, F9), (24*8)(R20)
	FSTPD (F10, F11), (26*8)(R20)
	FSTPD (F12, F13), (28*8)(R20)
	FSTPD (F14, F15), (30*8)(R20)
	RET

// unspillArgs loads args into registers from a *internal/abi.RegArgs in R20.
TEXT ·unspillArgs(SB),NOSPLIT,$0-0
	LDP (0*8)(R20), (R0, R1)
	LDP (2*8)(R20), (R2, R3)
	LDP (4*8)(R20), (R4, R5)
	LDP (6*8)(R20), (R6, R7)
	LDP (8*8)(R20), (R8, R9)
	LDP (10*8)(R20), (R10, R11)
	LDP (12*8)(R20), (R12, R13)
	LDP (14*8)(R20), (R14, R15)
	FLDPD (16*8)(R20), (F0, F1)
	FLDPD (18*8)(R20), (F2, F3)
	FLDPD (20*8)(R20), (F4, F5)
	FLDPD (22*8)(R20), (F6, F7)
	FLDPD (24*8)(R20), (F8, F9)
	FLDPD (26*8)(R20), (F10, F11)
	FLDPD (28*8)(R20), (F12, F13)
	FLDPD (30*8)(R20), (F14, F15)
	RET

// reflectcall: call a function with the given argument list
// func call(stackArgsType *_type, f *FuncVal, stackArgs *byte, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs).
// we don't have variable-sized frames, so we use a small number
// of constant-sized-frame functions to encode a few bits of size in the pc.
// Caution: ugly multiline assembly macros in your future!

#define DISPATCH(NAME,MAXSIZE)		\
	MOVD $MAXSIZE, R27;		\
	CMP R27, R16;		\
	BGT 3(PC);			\
	MOVD $NAME(SB), R27;	\
	B (R27)
// Note: can't just "B NAME(SB)" - bad inlining results.

TEXT ·Call(SB), NOSPLIT|NOFRAME, $0-48
	MOVWU frameSize+32(FP), R16
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
	MOVD $·badreflectcall(SB), R0
	B (R0)

#define CALLFN(NAME,MAXSIZE)			\
TEXT NAME(SB), WRAPPER, $MAXSIZE-48;		\
	NO_LOCAL_POINTERS;			\
	/* copy arguments to stack */		\
	MOVD stackArgs+16(FP), R3;			\
	MOVWU stackArgsSize+24(FP), R4;		\
	ADD $8, RSP, R5;			\
	BIC $0xf, R4, R6;			\
	CBZ R6, 6(PC);			\
	/* if R6=(argsize&~15) != 0 */		\
	ADD R6, R5, R6;			\
	/* copy 16 bytes a time */		\
	LDP.P 16(R3), (R7, R8);		\
	STP.P (R7, R8), 16(R5);		\
	CMP R5, R6;				\
	BNE -3(PC);				\
	AND $0xf, R4, R6;			\
	CBZ R6, 6(PC);			\
	/* if R6=(argsize&15) != 0 */		\
	ADD R6, R5, R6;			\
	/* copy 1 byte a time for the rest */	\
	MOVBU.P 1(R3), R7;			\
	MOVBU.P R7, 1(R5);			\
	CMP R5, R6;				\
	BNE -3(PC);				\
	/* set up argument registers */		\
	MOVD regArgs+40(FP), R20;		\
	CALL ·unspillArgs(SB);		\
	/* call function */			\
	MOVD f+8(FP), R26;			\
	MOVD (R26), R20;			\
	PCDATA $PCDATA_StackMapIndex, $0;	\
	BL (R20);				\
	/* copy return values back */		\
	MOVD regArgs+40(FP), R20;		\
	CALL ·spillArgs(SB);		\
	MOVD stackArgsType+0(FP), R7;		\
	MOVD stackArgs+16(FP), R3;			\
	MOVWU stackArgsSize+24(FP), R4;			\
	MOVWU stackRetOffset+28(FP), R6;		\
	ADD $8, RSP, R5;			\
	ADD R6, R5; 			\
	ADD R6, R3;				\
	SUB R6, R4;				\
	BL callRet<>(SB);			\
	RET

// callRet copies return values back at the end of call*. This is a
// separate function so it can allocate stack space for the arguments
// to reflectcallmove. It does not follow the Go ABI; it expects its
// arguments in registers.
TEXT callRet<>(SB), NOSPLIT, $48-0
	NO_LOCAL_POINTERS
	STP (R7, R3), 8(RSP)
	STP (R5, R4), 24(RSP)
	MOVD R20, 40(RSP)
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

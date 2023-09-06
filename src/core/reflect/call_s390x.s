// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h" // for NO_LOCAL_POINTERS

// reflectcall: call a function with the given argument list
// func call(stackArgsType *_type, f *FuncVal, stackArgs *byte, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs).
// we don't have variable-sized frames, so we use a small number
// of constant-sized-frame functions to encode a few bits of size in the pc.
// Caution: ugly multiline assembly macros in your future!

#define DISPATCH(NAME,MAXSIZE)		\
	MOVD $MAXSIZE, R4;		\
	CMP R3, R4;		\
	BGT 3(PC);			\
	MOVD $NAME(SB), R5;	\
	BR (R5)
// Note: can't just "BR NAME(SB)" - bad inlining results.

TEXT ·Call(SB), NOSPLIT, $-8-48
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
	MOVD $·badreflectcall(SB), R5
	BR (R5)

#define CALLFN(NAME,MAXSIZE)			\
TEXT NAME(SB), WRAPPER, $MAXSIZE-48;		\
	NO_LOCAL_POINTERS;			\
	/* copy arguments to stack */		\
	MOVD stackArgs+16(FP), R4;			\
	MOVWZ stackArgsSize+24(FP), R5;		\
	MOVD $stack-MAXSIZE(SP), R6;		\
loopArgs: /* copy 256 bytes at a time */	\
	CMP R5, $256;			\
	BLT tailArgs;			\
	SUB $256, R5;			\
	MVC $256, 0(R4), 0(R6);		\
	MOVD $256(R4), R4;			\
	MOVD $256(R6), R6;			\
	BR loopArgs;			\
tailArgs: /* copy remaining bytes */		\
	CMP R5, $0;				\
	BEQ callFunction;			\
	SUB $1, R5;				\
	EXRL $callfnMVC<>(SB), R5;		\
callFunction:					\
	MOVD f+8(FP), R12;			\
	MOVD (R12), R8;			\
	PCDATA  $PCDATA_StackMapIndex, $0;	\
	BL (R8);				\
	/* copy return values back */		\
	MOVD stackArgsType+0(FP), R7;		\
	MOVD stackArgs+16(FP), R6;			\
	MOVWZ stackArgsSize+24(FP), R5;			\
	MOVD $stack-MAXSIZE(SP), R4;		\
	MOVWZ stackRetOffset+28(FP), R1;		\
	ADD R1, R4;				\
	ADD R1, R6;				\
	SUB R1, R5;				\
	BL callRet<>(SB);			\
	RET

// callRet copies return values back at the end of call*. This is a
// separate function so it can allocate stack space for the arguments
// to reflectcallmove. It does not follow the Go ABI; it expects its
// arguments in registers.
TEXT callRet<>(SB), NOSPLIT, $40-0
	MOVD R7, 8(R15)
	MOVD R6, 16(R15)
	MOVD R4, 24(R15)
	MOVD R5, 32(R15)
	MOVD $0, 40(R15)
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

// Not a function: target for EXRL (execute relative long) instruction.
TEXT callfnMVC<>(SB),NOSPLIT|NOFRAME,$0-0
	MVC $1, 0(R4), 0(R6)

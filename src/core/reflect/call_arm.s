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
	CMP $MAXSIZE, R0;		\
	B.HI 3(PC);			\
	MOVW $NAME(SB), R1;		\
	B (R1)

TEXT ·Call(SB),NOSPLIT|NOFRAME,$0-28
	MOVW frameSize+20(FP), R0
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
	MOVW $·badreflectcall(SB), R1
	B (R1)

#define CALLFN(NAME,MAXSIZE)			\
TEXT NAME(SB), WRAPPER, $MAXSIZE-28;		\
	NO_LOCAL_POINTERS;			\
	/* copy arguments to stack */		\
	MOVW stackArgs+8(FP), R0;		\
	MOVW stackArgsSize+12(FP), R2;		\
	ADD $4, R13, R1;			\
	CMP $0, R2;				\
	B.EQ 5(PC);				\
	MOVBU.P 1(R0), R5;			\
	MOVBU.P R5, 1(R1);			\
	SUB $1, R2, R2;			\
	B -5(PC);				\
	/* call function */			\
	MOVW f+4(FP), R7;			\
	MOVW (R7), R0;			\
	PCDATA  $PCDATA_StackMapIndex, $0;	\
	BL (R0);				\
	/* copy return values back */		\
	MOVW stackArgsType+0(FP), R4;		\
	MOVW stackArgs+8(FP), R0;		\
	MOVW stackArgsSize+12(FP), R2;		\
	MOVW stackArgsRetOffset+16(FP), R3;		\
	ADD $4, R13, R1;			\
	ADD R3, R1;				\
	ADD R3, R0;				\
	SUB R3, R2;				\
	BL callRet<>(SB);			\
	RET

// callRet copies return values back at the end of call*. This is a
// separate function so it can allocate stack space for the arguments
// to reflectcallmove. It does not follow the Go ABI; it expects its
// arguments in registers.
TEXT callRet<>(SB), NOSPLIT, $20-0
	MOVW R4, 4(R13)
	MOVW R0, 8(R13)
	MOVW R1, 12(R13)
	MOVW R2, 16(R13)
	MOVW $0, R7
	MOVW R7, 20(R13)
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

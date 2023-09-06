// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h" // for NO_LOCAL_POINTERS

// spillArgs stores return values from registers to a *internal/abi.RegArgs in X25.
TEXT ·spillArgs(SB),NOSPLIT,$0-0
	MOV X10, (0*8)(X25)
	MOV X11, (1*8)(X25)
	MOV X12, (2*8)(X25)
	MOV X13, (3*8)(X25)
	MOV X14, (4*8)(X25)
	MOV X15, (5*8)(X25)
	MOV X16, (6*8)(X25)
	MOV X17, (7*8)(X25)
	MOV X8,  (8*8)(X25)
	MOV X9,  (9*8)(X25)
	MOV X18, (10*8)(X25)
	MOV X19, (11*8)(X25)
	MOV X20, (12*8)(X25)
	MOV X21, (13*8)(X25)
	MOV X22, (14*8)(X25)
	MOV X23, (15*8)(X25)
	MOVD F10, (16*8)(X25)
	MOVD F11, (17*8)(X25)
	MOVD F12, (18*8)(X25)
	MOVD F13, (19*8)(X25)
	MOVD F14, (20*8)(X25)
	MOVD F15, (21*8)(X25)
	MOVD F16, (22*8)(X25)
	MOVD F17, (23*8)(X25)
	MOVD F8,  (24*8)(X25)
	MOVD F9,  (25*8)(X25)
	MOVD F18, (26*8)(X25)
	MOVD F19, (27*8)(X25)
	MOVD F20, (28*8)(X25)
	MOVD F21, (29*8)(X25)
	MOVD F22, (30*8)(X25)
	MOVD F23, (31*8)(X25)
	RET

// unspillArgs loads args into registers from a *internal/abi.RegArgs in X25.
TEXT ·unspillArgs(SB),NOSPLIT,$0-0
	MOV (0*8)(X25), X10
	MOV (1*8)(X25), X11
	MOV (2*8)(X25), X12
	MOV (3*8)(X25), X13
	MOV (4*8)(X25), X14
	MOV (5*8)(X25), X15
	MOV (6*8)(X25), X16
	MOV (7*8)(X25), X17
	MOV (8*8)(X25), X8
	MOV (9*8)(X25), X9
	MOV (10*8)(X25), X18
	MOV (11*8)(X25), X19
	MOV (12*8)(X25), X20
	MOV (13*8)(X25), X21
	MOV (14*8)(X25), X22
	MOV (15*8)(X25), X23
	MOVD (16*8)(X25), F10
	MOVD (17*8)(X25), F11
	MOVD (18*8)(X25), F12
	MOVD (19*8)(X25), F13
	MOVD (20*8)(X25), F14
	MOVD (21*8)(X25), F15
	MOVD (22*8)(X25), F16
	MOVD (23*8)(X25), F17
	MOVD (24*8)(X25), F8
	MOVD (25*8)(X25), F9
	MOVD (26*8)(X25), F18
	MOVD (27*8)(X25), F19
	MOVD (28*8)(X25), F20
	MOVD (29*8)(X25), F21
	MOVD (30*8)(X25), F22
	MOVD (31*8)(X25), F23
	RET

// reflectcall: call a function with the given argument list
// func call(stackArgsType *_type, f *FuncVal, stackArgs *byte, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs).
// we don't have variable-sized frames, so we use a small number
// of constant-sized-frame functions to encode a few bits of size in the pc.
// Caution: ugly multiline assembly macros in your future!

#define DISPATCH(NAME,MAXSIZE)	\
	MOV $MAXSIZE, T1 \
	BLTU T1, T0, 3(PC)	\
	MOV $NAME(SB), T2;	\
	JALR ZERO, T2
// Note: can't just "BR NAME(SB)" - bad inlining results.

// func call(stackArgsType *_type, fn, stackArgs unsafe.Pointer, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs).
TEXT ·Call(SB), NOSPLIT|NOFRAME, $0-48
	MOVWU frameSize+32(FP), T0
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
	MOV $·badreflectcall(SB), T2
	JALR ZERO, T2

#define CALLFN(NAME,MAXSIZE)			\
TEXT NAME(SB), WRAPPER, $MAXSIZE-48;		\
	NO_LOCAL_POINTERS;			\
	/* copy arguments to stack */		\
	MOV stackArgs+16(FP), A1;			\
	MOVWU stackArgsSize+24(FP), A2;		\
	MOV X2, A3;				\
	ADD $8, A3;				\
	ADD A3, A2;				\
	BEQ A3, A2, 6(PC);			\
	MOVBU (A1), A4;			\
	ADD $1, A1;				\
	MOVB A4, (A3);			\
	ADD $1, A3;				\
	JMP -5(PC);				\
	/* set up argument registers */		\
	MOV regArgs+40(FP), X25;		\
	CALL ·unspillArgs(SB);		\
	/* call function */			\
	MOV f+8(FP), CTXT;			\
	MOV (CTXT), X25;			\
	PCDATA  $PCDATA_StackMapIndex, $0;	\
	JALR RA, X25;				\
	/* copy return values back */		\
	MOV regArgs+40(FP), X25;		\
	CALL ·spillArgs(SB);		\
	MOV stackArgsType+0(FP), A5;		\
	MOV stackArgs+16(FP), A1;			\
	MOVWU stackArgsSize+24(FP), A2;			\
	MOVWU stackRetOffset+28(FP), A4;		\
	ADD $8, X2, A3;			\
	ADD A4, A3; 			\
	ADD A4, A1;				\
	SUB A4, A2;				\
	CALL callRet<>(SB);			\
	RET

// callRet copies return values back at the end of call*. This is a
// separate function so it can allocate stack space for the arguments
// to reflectcallmove. It does not follow the Go ABI; it expects its
// arguments in registers.
TEXT callRet<>(SB), NOSPLIT, $40-0
	NO_LOCAL_POINTERS
	MOV A5, 8(X2)
	MOV A1, 16(X2)
	MOV A3, 24(X2)
	MOV A2, 32(X2)
	MOV X25, 40(X2)
	CALL ·reflectcallmove(SB)
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

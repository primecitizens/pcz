// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && amd64

#include "textflag.h"
#include "funcdata.h" // for NO_LOCAL_POINTERS

// spillArgs stores return values from registers to a *core/abi.RegArgs in R12.
TEXT ·spillArgs(SB),NOSPLIT,$0-0
	MOVQ AX, 0(R12)
	MOVQ BX, 8(R12)
	MOVQ CX, 16(R12)
	MOVQ DI, 24(R12)
	MOVQ SI, 32(R12)
	MOVQ R8, 40(R12)
	MOVQ R9, 48(R12)
	MOVQ R10, 56(R12)
	MOVQ R11, 64(R12)
	MOVQ X0, 72(R12)
	MOVQ X1, 80(R12)
	MOVQ X2, 88(R12)
	MOVQ X3, 96(R12)
	MOVQ X4, 104(R12)
	MOVQ X5, 112(R12)
	MOVQ X6, 120(R12)
	MOVQ X7, 128(R12)
	MOVQ X8, 136(R12)
	MOVQ X9, 144(R12)
	MOVQ X10, 152(R12)
	MOVQ X11, 160(R12)
	MOVQ X12, 168(R12)
	MOVQ X13, 176(R12)
	MOVQ X14, 184(R12)
	RET

// unspillArgs loads args into registers from a *core/abi.RegArgs in R12.
TEXT ·unspillArgs(SB),NOSPLIT,$0-0
	MOVQ 0(R12), AX
	MOVQ 8(R12), BX
	MOVQ 16(R12), CX
	MOVQ 24(R12), DI
	MOVQ 32(R12), SI
	MOVQ 40(R12), R8
	MOVQ 48(R12), R9
	MOVQ 56(R12), R10
	MOVQ 64(R12), R11
	MOVQ 72(R12), X0
	MOVQ 80(R12), X1
	MOVQ 88(R12), X2
	MOVQ 96(R12), X3
	MOVQ 104(R12), X4
	MOVQ 112(R12), X5
	MOVQ 120(R12), X6
	MOVQ 128(R12), X7
	MOVQ 136(R12), X8
	MOVQ 144(R12), X9
	MOVQ 152(R12), X10
	MOVQ 160(R12), X11
	MOVQ 168(R12), X12
	MOVQ 176(R12), X13
	MOVQ 184(R12), X14
	RET

// reflectcall: call a function with the given argument list
// func call(stackArgsType *_type, f *FuncVal, stackArgs *byte, stackArgsSize, stackRetOffset, frameSize uint32, regArgs *abi.RegArgs).
// we don't have variable-sized frames, so we use a small number
// of constant-sized-frame functions to encode a few bits of size in the pc.
// Caution: ugly multiline assembly macros in your future!

#define DISPATCH(NAME,MAXSIZE)		\
	CMPQ CX, $MAXSIZE;		\
	JA 3(PC);			\
	MOVQ $NAME(SB), AX;		\
	JMP AX
// Note: can't just "JMP NAME(SB)" - bad inlining results.

TEXT ·Call(SB), NOSPLIT, $0-48
	MOVLQZX frameSize+32(FP), CX
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
	MOVQ $·badreflectcall(SB), AX
	JMP AX

#define CALLFN(NAME,MAXSIZE)			\
TEXT NAME(SB), WRAPPER, $MAXSIZE-48;		\
	NO_LOCAL_POINTERS;			\
	/* copy arguments to stack */		\
	MOVQ stackArgs+16(FP), SI;		\
	MOVLQZX stackArgsSize+24(FP), CX;		\
	MOVQ SP, DI;				\
	REP;MOVSB;				\
	/* set up argument registers */		\
	MOVQ    regArgs+40(FP), R12;		\
	CALL    ·unspillArgs(SB);		\
	/* call function */			\
	MOVQ f+8(FP), DX;			\
	PCDATA  $PCDATA_StackMapIndex, $0;	\
	MOVQ (DX), R12;			\
	CALL R12;				\
	/* copy register return values back */		\
	MOVQ    regArgs+40(FP), R12;		\
	CALL    ·spillArgs(SB);		\
	MOVLQZX stackArgsSize+24(FP), CX;		\
	MOVLQZX stackRetOffset+28(FP), BX;		\
	MOVQ stackArgs+16(FP), DI;		\
	MOVQ stackArgsType+0(FP), DX;		\
	MOVQ SP, SI;				\
	ADDQ BX, DI;				\
	ADDQ BX, SI;				\
	SUBQ BX, CX;				\
	CALL callRet<>(SB);			\
	RET

// callRet copies return values back at the end of call*. This is a
// separate function so it can allocate stack space for the arguments
// to reflectcallmove. It does not follow the Go ABI; it expects its
// arguments in registers.
TEXT callRet<>(SB), NOSPLIT, $40-0
	NO_LOCAL_POINTERS
	MOVQ DX, 0(SP)
	MOVQ DI, 8(SP)
	MOVQ SI, 16(SP)
	MOVQ CX, 24(SP)
	MOVQ R12, 32(SP)
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

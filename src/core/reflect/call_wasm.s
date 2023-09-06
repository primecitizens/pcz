// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"
#include "funcdata.h" // for NO_LOCAL_POINTERS

#define DISPATCH(NAME, MAXSIZE) \
	Get R0; \
	I64Const $MAXSIZE; \
	I64LeU; \
	If; \
		JMP NAME(SB); \
	End

TEXT ·Call(SB),NOSPLIT,$0-48
	I64Load fn+8(FP)
	I64Eqz
	If
		CALLNORESUME ·sigpanic<ABIInternal>(SB)
	End

	MOVW frameSize+32(FP), R0

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
	JMP ·badreflectcall(SB)

#define CALLFN(NAME, MAXSIZE) \
TEXT NAME(SB), WRAPPER, $MAXSIZE-48; \
	NO_LOCAL_POINTERS; \
	MOVW stackArgsSize+24(FP), R0; \
	\
	Get R0; \
	I64Eqz; \
	Not; \
	If; \
		Get SP; \
		I64Load stackArgs+16(FP); \
		I32WrapI64; \
		I64Load stackArgsSize+24(FP); \
		I32WrapI64; \
		MemoryCopy; \
	End; \
	\
	MOVD f+8(FP), CTXT; \
	Get CTXT; \
	I32WrapI64; \
	I64Load $0; \
	CALL; \
	\
	I64LoadUint32U stackRetOffset+28(FP); \
	Set R0; \
	\
	MOVD stackArgsType+0(FP), RET0; \
	\
	I64Load stackArgs+16(FP); \
	Get R0; \
	I64Add; \
	Set RET1; \
	\
	Get SP; \
	I64ExtendI32U; \
	Get R0; \
	I64Add; \
	Set RET2; \
	\
	I64LoadUint32U stackArgsSize+24(FP); \
	Get R0; \
	I64Sub; \
	Set RET3; \
	\
	CALL callRet<>(SB); \
	RET

// callRet copies return values back at the end of call*. This is a
// separate function so it can allocate stack space for the arguments
// to reflectcallmove. It does not follow the Go ABI; it expects its
// arguments in registers.
TEXT callRet<>(SB), NOSPLIT, $40-0
	NO_LOCAL_POINTERS
	MOVD RET0, 0(SP)
	MOVD RET1, 8(SP)
	MOVD RET2, 16(SP)
	MOVD RET3, 24(SP)
	MOVD $0,   32(SP)
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

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build pcz && wasm

#include "textflag.h"

// func SetG(*stdgo.GHead)
TEXT ·SetG(SB),NOSPLIT,$8-0
	MOVD gp+0(FP), g
	RET

// func Topframe(arg, sp uintptr, fn func(uintptr))
TEXT ·Topframe(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	MOVD arg+0(FP), R0
	MOVD fn+16(FP), R1

	I64Load sp+8(FP)
	I32WrapI64
	Set SP

	MOVD R1, 0(SP)

	Get SP
	I32Const $8
	I32Add
	Set SP
	
	MOVD R0, 0(SP) // arg

	Get R1
	CALL // fn(arg)
	RET

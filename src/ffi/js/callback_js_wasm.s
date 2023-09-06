// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && js && wasm

#include "textflag.h"

// wasm_export_run gets called from JavaScript.
//
// In the official go std, it is the entrypoint of the js/wasm program.
// But in ours, it is used to call go function directly.
//
// R0: argc (i32), repurposed as Ref<CallbackArgs>
// R1: argv (i32), unused.
TEXT wasm_export_run(SB),NOSPLIT,$0
	Get R1
	I32Eqz
	Not
	If
		Get R1
		Set SP
	End

	// prepare args for the go func
	Get SP
	I32Const $16
	I32Sub
	Set SP

	Get SP
	Get R0
	I32Store $8 // argc

	Get SP
	Get R1
	I32Store $12 // argv

	I32Const $0
	Call ·handleCallback(SB)

	Return

// wasm_export_getsp gets called from JavaScript to retrieve the SP.
TEXT wasm_export_getsp(SB),NOSPLIT,$0
	Get SP
	Return

// func callHandler(recv, targetPC uintptr, ctx *CallbackContext, fn cbFunc[unsafe.Pointer])
TEXT ·callHandler(SB),NOSPLIT,$0-32
	// The go assembler will insert:
	//
	//   SP -= 8
	//   0(SP) = $·callHandler+X(SB) // PC_F = index of callhandler, PC_B = X (block label)

	// MOVD fn+24(FP), 0(SP)
	MOVD fn+24(FP), R0

	// SP += 8 to avoid copy stack.
	Get SP
	I32Const $8
	I32Add
	Set SP

	Get R0
	CALL

	// here we use wasm return instead of go function return (RET) to avoid SP += 8
	// since we have done it before CALL.
	I32Const $0
	Return

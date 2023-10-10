// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build !noos && js && wasm

#include "textflag.h"

// wasm_export_run has wasm signature of (argc i32, argv i32) -> ()
//
// it gets called from JavaScript.
//
// In the official go std, it is the entrypoint of the js/wasm program.
// But in pcz, it is used to handle callbacks.
//
// R0: argc (i32), heap.Ref<CallbackContext>, if 0, continue paused execution.
// R1: argv (i32), optional SP.
TEXT wasm_export_run(SB),NOSPLIT|NOFRAME,$0
	Get R1
	I32Eqz
	Not
	If
	  Get R1
	  Set SP
	End

	Get R0
	I32Eqz
	If
	  // the program is awaiting.
	  Call wasm_pc_f_loop(SB)
	  Return
	End

	// prepare arg ctx and pLR for handleCallback
	Get SP
	I32Const $16 // 8 bytes for args, 8 bytes for LR
	I32Sub
	Set SP

	Get SP
	Get R0
	I64ExtendI32U
	I64Store32 $8 // 8(SP) = R0 (= ref ctx)

	Get SP
	I64Const $wasm_pc_f_loop(SB)
	I64Store // LR

	I32Const $0
	Call ·handleCallback(SB)
	If
	  // the callback issued awaiting, return immediately as PAUSE is set to 1.
	  Return
	End

	// only loop for async callback, synchronous callbacks (go -> js -> go)
	// are expected to resume the program on its own.
	Get PAUSE
	If
	  Call wasm_pc_f_loop(SB)
	Else
	  // pop args for `Call ·handleCallback(SB)` (LR already poped by it)
	  Get SP
	  I32Const $8
	  I32Add
	  Set SP
	End

	Return

// wasm_pc_f_loop has wasm signature of () -> ()
//
// NOTE: this function is different from the `wasm_pc_f_loop` found in
// the standard go runtime.
TEXT wasm_pc_f_loop(SB),NOSPLIT|NOFRAME,$0
	I32Const $0
	Set PAUSE

loop:
	Loop
	  Get SP
	  I32Const $8
	  I32Sub
	  I64Load $0
	  I64Const $wasm_pc_f_loop(SB)
	  I64Eq
	  If
	    // reached end of the loop (callback done), pop the LR and return.
	    Get SP
	    I32Const $8
	    I32Add
	    Set SP

	    Return
	  End

	  // Get PC_B & PC_F from -8(SP)
	  Get SP
	  I32Const $8
	  I32Sub
	  I32Load16U $0 // PC_B

	  Get SP
	  I32Const $8
	  I32Sub
	  I32Load16U $2 // PC_F

	  CallIndirect $0
	  I32Eqz
	  BrIf loop
	End

	Return

// wasm_export_getsp gets called from JavaScript to retrieve the SP.
TEXT wasm_export_getsp(SB),NOSPLIT|NOFRAME,$0
	Get SP
	Return

// func callDispatcher(recv, targetPC uintptr, ctx *CallbackContext, fn uintptr)
TEXT ·callDispatcher(SB),NOSPLIT|NOFRAME|TOPFRAME,$0-32
	skip:
	Block // $2
	  done:
	  Block // $1
	    entry:
	    Block // $0
	      // TODO(BrTable): currently go asm doesn't support BrTable in assembly text
	      // but emits BrTable automatically.
	      Get PC_B
	      I32Eqz
	      BrIf entry

	      Get PC_B
	      I32Const $1
	      I32Eq
	      BrIf done

	      Get PC_B
	      I32Const $2
	      I32Eq
	      BrIf skip
	    End // $0

	    Get SP
	    I64Const $·callDispatcher+1(SB)
	    I64Store // LR

	    I32Const $0 // PC_B

	    I64Load 32(SP) // PC = (8/* LR */ + 24 /* offset to last arg fn */)(SP)
	    I32WrapI64
	    I32Const $16
	    I32ShrU // PC_F

	    CallIndirect $0
	    BrIf skip
	  End // $1

	  // A callback can only be dispatched once, change LR to FIXUP 
	  // in handleCallback
	  // so wasm_pc_f_loop doesnt't loop indefinitely
	  Get SP
	  I32Const $8
	  I32Sub
	  // TODO: is there a way to get resume point in go code?
	  // currently we have to compile and dump the assembly to find target
	  // PC_B:
	  // 	Block_7 (from inner most to out most) calling cb.resolveFn
	  // 	PC_B = 16 (index of Block_7 in the BrTable).
	  I64Const $·handleCallback+16(SB)
	  I64Store // LR

	  I32Const $0
	  Return
	End // $2

	I32Const $1
	Return

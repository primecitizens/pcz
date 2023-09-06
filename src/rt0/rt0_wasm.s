// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !noos && wasm

#include "textflag.h"
#include "go_asm.h" // for rt0stack__size

#ifdef GOOS_wasip1

TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	MOVD $·tmpStack+(rt0stack__size)(SB), SP
	I32Const $0
	Call ·_rt0_wasm(SB)
	Return

#endif

#ifdef GOOS_js

// rt0 for js/wasm only exists to mark the exported functions as alive.
TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	I32Const $·_rt0_wasm(SB)
	Drop
	I32Const $wasm_export_run(SB)
	Drop
	I32Const $wasm_export_resume(SB)
	Drop
	I32Const $wasm_export_getsp(SB)
	Drop

// wasm_export_resume() gets called from JavaScript.
//
// In the official go std, it resumes the execution of Go code until it
// needs to wait for an event.
//
// But in this std, it is the entrypoint of the program.
//
// NOTE: the symbol name is known to the go tool link, only change when the
// linker changes its name or signature.
TEXT wasm_export_resume(SB),NOSPLIT,$0
	MOVD $·tmpStack+(rt0stack__size)(SB), SP
	I32Const $0
	Call ·_rt0_wasm(SB)
	Return

#endif

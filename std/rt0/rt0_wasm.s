// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !noos && wasm

#include "textflag.h"

#ifdef GOOS_wasip1

// NOTE: the name is not "rt0" because the go1.21 linker hard coded 
// "_rt0_wasm_wasip1" being exported as "_start".
//
// See ${GOROOT}/src/cmd/link/internal/wasm/asm.go#func:writeExportSec
TEXT _rt0_wasm_wasip1(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	// See wasm_export_resume
	MOVD $(4096+8192-(8/* argc */+8 /* argv */ + 8 /* LR */)), SP

	I32Const $0
	Call ·rt0(SB)
	Drop
	Return

#endif

#ifdef GOOS_js

// rt0 for js/wasm only exists to mark the exported functions as alive.
TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
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
// But in pcz, we use it as the entrypoint of the program.
//
// NOTE: the symbol name is known to the go tool link, only change when the
// linker changes its name or signature.
TEXT wasm_export_resume(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	// go linker reserves 4096 bytes for zero page, then 8192 bytes for wasm_exec.js
	// to set argv and envv, but we don't do that, use the reserved 8192 bytes as
	// the initial stack.
	//
	// see ${GOROOT}/src/cmd/link/internal/ld/data.go#const:wasmMinDataAddr
	// 
	// NOTE: argc (uint32) is 8 bytes in wasm.
	MOVD $(4096+8192-(8/* argc */+8 /* argv */ + 8 /* LR */)), SP

	I32Const $0
	Call ·rt0(SB)
	Drop
	Return

#endif

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// When building with the "noos" tag, the GOOS is actually "js".
//
//go:build noos && wasm

#include "textflag.h"

TEXT rt0(SB),NOSPLIT|NOFRAME|TOPFRAME,$0
	// See comments inside wasm_export_resume
	MOVD $(4096+8192-(8/* argc */+8 /* argv */ + 8 /* LR */)), SP

	I32Const $0
	Call ·rt0(SB)
	Drop
	Return

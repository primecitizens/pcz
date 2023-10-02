// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && wasm

#include "textflag.h"

TEXT ·Breakpoint(SB),NOSPLIT,$0-0
	UNDEF

TEXT ·Abort(SB),NOSPLIT,$0-0
	UNDEF

TEXT ·Return0(SB),NOSPLIT,$0-0
	MOVD $0, RET0
	RET

TEXT ·GetCallerPC(SB),NOSPLIT|NOFRAME,$0-8
	Get SP
	I64ExtendI32U
	I64Const $8
	I64Add
	I32WrapI64

	Get SP
	I64Load $0
	I64Store $0

	Get SP
	I32Const $8
	I32Add
	Set SP

	I32Const $0
	Return

TEXT ·GetCallerSP(SB),NOSPLIT|NOFRAME,$0-8
	Get SP
	I64ExtendI32U
	I64Const $8
	I64Add
	I32WrapI64

	Get SP
	I64ExtendI32U
	I64Const $8
	I64Add
	I64Store $0

	Get SP
	I32Const $8
	I32Add
	Set SP

	I32Const $0
	Return

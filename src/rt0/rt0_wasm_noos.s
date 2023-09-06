// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build noos && wasm

#include "go_asm.h" // for rt0stack__size
#include "textflag.h"

TEXT rt0(SB),NOSPLIT|NOFRAME,$0
	MOVD $·tmpStack+(rt0stack__size)(SB), SP
	I32Const $0
	Call ·_rt0_wasm(SB)
	Return

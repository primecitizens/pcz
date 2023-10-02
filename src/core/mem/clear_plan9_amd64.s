// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && amd64 && plan9

#include "textflag.h"

// See memclrNoHeapPointers Go doc for important implementation constraints.

// func Clear(ptr unsafe.Pointer, n uintptr)
TEXT ·Clear(SB),NOSPLIT,$0-16
	MOVQ ptr+0(FP), DI
	MOVQ n+8(FP), CX
	MOVQ CX, BX
	ANDQ $7, BX
	SHRQ $3, CX
	MOVQ $0, AX
	CLD
	REP
	STOSQ
	MOVQ BX, CX
	REP
	STOSB
	RET

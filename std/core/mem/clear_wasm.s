// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && wasm

#include "textflag.h"

// See memclrNoHeapPointers Go doc for important implementation constraints.

// func Clear(ptr unsafe.Pointer, n uintptr)
TEXT ·Clear(SB), NOSPLIT, $0-16
	MOVD ptr+0(FP), R0
	MOVD n+8(FP), R1

	Get R0
	I32WrapI64
	I32Const $0
	Get R1
	I32WrapI64
	MemoryFill
	RET

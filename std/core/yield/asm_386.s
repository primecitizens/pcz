// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386

#include "textflag.h"

TEXT ·Proc(SB),NOSPLIT,$0-0
	MOVL cycles+0(FP), AX
LOOP:
	PAUSE
	SUBL $1, AX
	JNZ LOOP
	RET

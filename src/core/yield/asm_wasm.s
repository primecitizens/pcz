// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasm

#include "textflag.h"

TEXT ·Proc(SB),NOSPLIT,$0-0 // FIXME
	RET

TEXT ·Thread(SB),NOSPLIT,$0-0
	I32Const $1
	Set PAUSE
	RETUNWIND

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build riscv64

#include "textflag.h"

// func Proc(cycles uint32)
TEXT ·Proc<ABIInternal>(SB),NOSPLIT,$0
	RET

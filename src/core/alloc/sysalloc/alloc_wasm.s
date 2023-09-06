// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·growMemory(SB),NOSPLIT,$0
	Get SP
	I32Load pages+0(FP)
	GrowMemory
	I32Store ret+8(FP)
	RET

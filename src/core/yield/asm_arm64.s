// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·Proc<ABIInternal>(SB),NOSPLIT,$0
LOOP:
	YIELD
	SUBW $1, R0
	CBNZ R0, LOOP
	RET

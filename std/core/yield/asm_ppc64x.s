// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ppc64 || ppc64le

TEXT ·Proc(SB),NOSPLIT|NOFRAME,$0-4
	MOVW cycles+0(FP), R7
	// POWER does not have a pause/yield instruction equivalent.
	// Instead, we can lower the program priority by setting the
	// Program Priority Register prior to the wait loop and set it
	// back to default afterwards. On Linux, the default priority is
	// medium-low. For details, see page 837 of the ISA 3.0.
	OR R1, R1, R1 // Set PPR priority to low
LOOP:
	SUB $1, R7
	CMP $0, R7
	BNE LOOP
	OR R6, R6, R6 // Set PPR priority back to medium-low
	RET

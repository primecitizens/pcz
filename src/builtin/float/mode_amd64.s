// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

#include "textflag.h"

// func SetFPUMode(uint32)
TEXT ·SetFPUMode(SB),NOSPLIT,$0
	LDMXCSR mode+0(FP)
	RET

// func GetFPUMode() uint32
TEXT ·GetFPUMode(SB),NOSPLIT,$0
	STMXCSR mode+0(FP)
	RET

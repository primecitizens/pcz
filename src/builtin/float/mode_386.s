// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

#include "textflag.h"

// func SetFPUMode(uint16)
TEXT ·SetFPUMode(SB),NOSPLIT,$2-2
	FLDCW mode+0(FP)
	RET

// func GetFPUMode() uint16
TEXT ·GetFPUMode(SB),NOSPLIT,$2-2
	FSTCW mode+0(FP)
	RET

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

#include "textflag.h"

TEXT ·SetFPUMode<ABIInternal>(SB),NOSPLIT,$0
	MSR R0, FPCR
	RET

TEXT ·GetFPUMode<ABIInternal>(SB),NOSPLIT,$0
	MRS FPCR, R0
	RET

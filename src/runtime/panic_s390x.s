// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && s390x

#include "textflag.h"

// Note: these functions use a special calling convention to save generated code space.
// Arguments are passed in registers, but the space for those arguments are allocated
// in the caller's stack frame. These stubs write the args into that stack space and
// then tail call to the corresponding runtime handler.
// The tail call makes these stubs disappear in backtraces.

TEXT runtime·panicIndex(SB),NOSPLIT,$0-16
	MOVD R0, x+0(FP)
	MOVD R1, y+8(FP)
	JMP runtime·goPanicIndex(SB)

TEXT runtime·panicIndexU(SB),NOSPLIT,$0-16
	MOVD R0, x+0(FP)
	MOVD R1, y+8(FP)
	JMP runtime·goPanicIndexU(SB)

TEXT runtime·panicSliceAlen(SB),NOSPLIT,$0-16
	MOVD R1, x+0(FP)
	MOVD R2, y+8(FP)
	JMP runtime·goPanicSliceAlen(SB)

TEXT runtime·panicSliceAlenU(SB),NOSPLIT,$0-16
	MOVD R1, x+0(FP)
	MOVD R2, y+8(FP)
	JMP runtime·goPanicSliceAlenU(SB)

TEXT runtime·panicSliceAcap(SB),NOSPLIT,$0-16
	MOVD R1, x+0(FP)
	MOVD R2, y+8(FP)
	JMP runtime·goPanicSliceAcap(SB)

TEXT runtime·panicSliceAcapU(SB),NOSPLIT,$0-16
	MOVD R1, x+0(FP)
	MOVD R2, y+8(FP)
	JMP runtime·goPanicSliceAcapU(SB)

TEXT runtime·panicSliceB(SB),NOSPLIT,$0-16
	MOVD R0, x+0(FP)
	MOVD R1, y+8(FP)
	JMP runtime·goPanicSliceB(SB)

TEXT runtime·panicSliceBU(SB),NOSPLIT,$0-16
	MOVD R0, x+0(FP)
	MOVD R1, y+8(FP)
	JMP runtime·goPanicSliceBU(SB)

TEXT runtime·panicSlice3Alen(SB),NOSPLIT,$0-16
	MOVD R2, x+0(FP)
	MOVD R3, y+8(FP)
	JMP runtime·goPanicSlice3Alen(SB)

TEXT runtime·panicSlice3AlenU(SB),NOSPLIT,$0-16
	MOVD R2, x+0(FP)
	MOVD R3, y+8(FP)
	JMP runtime·goPanicSlice3AlenU(SB)

TEXT runtime·panicSlice3Acap(SB),NOSPLIT,$0-16
	MOVD R2, x+0(FP)
	MOVD R3, y+8(FP)
	JMP runtime·goPanicSlice3Acap(SB)

TEXT runtime·panicSlice3AcapU(SB),NOSPLIT,$0-16
	MOVD R2, x+0(FP)
	MOVD R3, y+8(FP)
	JMP runtime·goPanicSlice3AcapU(SB)

TEXT runtime·panicSlice3B(SB),NOSPLIT,$0-16
	MOVD R1, x+0(FP)
	MOVD R2, y+8(FP)
	JMP runtime·goPanicSlice3B(SB)

TEXT runtime·panicSlice3BU(SB),NOSPLIT,$0-16
	MOVD R1, x+0(FP)
	MOVD R2, y+8(FP)
	JMP runtime·goPanicSlice3BU(SB)

TEXT runtime·panicSlice3C(SB),NOSPLIT,$0-16
	MOVD R0, x+0(FP)
	MOVD R1, y+8(FP)
	JMP runtime·goPanicSlice3C(SB)

TEXT runtime·panicSlice3CU(SB),NOSPLIT,$0-16
	MOVD R0, x+0(FP)
	MOVD R1, y+8(FP)
	JMP runtime·goPanicSlice3CU(SB)

TEXT runtime·panicSliceConvert(SB),NOSPLIT,$0-16
	MOVD R2, x+0(FP)
	MOVD R3, y+8(FP)
	JMP runtime·goPanicSliceConvert(SB)

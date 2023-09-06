// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// Note: these functions use a special calling convention to save generated code space.
// Arguments are passed in registers, but the space for those arguments are allocated
// in the caller's stack frame. These stubs write the args into that stack space and
// then tail call to the corresponding runtime handler.
// The tail call makes these stubs disappear in backtraces.

TEXT runtime·panicIndex(SB),NOSPLIT,$0-8
	MOVW R0, x+0(FP)
	MOVW R1, y+4(FP)
	JMP runtime·goPanicIndex(SB)

TEXT runtime·panicIndexU(SB),NOSPLIT,$0-8
	MOVW R0, x+0(FP)
	MOVW R1, y+4(FP)
	JMP runtime·goPanicIndexU(SB)

TEXT runtime·panicSliceAlen(SB),NOSPLIT,$0-8
	MOVW R1, x+0(FP)
	MOVW R2, y+4(FP)
	JMP runtime·goPanicSliceAlen(SB)

TEXT runtime·panicSliceAlenU(SB),NOSPLIT,$0-8
	MOVW R1, x+0(FP)
	MOVW R2, y+4(FP)
	JMP runtime·goPanicSliceAlenU(SB)

TEXT runtime·panicSliceAcap(SB),NOSPLIT,$0-8
	MOVW R1, x+0(FP)
	MOVW R2, y+4(FP)
	JMP runtime·goPanicSliceAcap(SB)

TEXT runtime·panicSliceAcapU(SB),NOSPLIT,$0-8
	MOVW R1, x+0(FP)
	MOVW R2, y+4(FP)
	JMP runtime·goPanicSliceAcapU(SB)

TEXT runtime·panicSliceB(SB),NOSPLIT,$0-8
	MOVW R0, x+0(FP)
	MOVW R1, y+4(FP)
	JMP runtime·goPanicSliceB(SB)

TEXT runtime·panicSliceBU(SB),NOSPLIT,$0-8
	MOVW R0, x+0(FP)
	MOVW R1, y+4(FP)
	JMP runtime·goPanicSliceBU(SB)

TEXT runtime·panicSlice3Alen(SB),NOSPLIT,$0-8
	MOVW R2, x+0(FP)
	MOVW R3, y+4(FP)
	JMP runtime·goPanicSlice3Alen(SB)

TEXT runtime·panicSlice3AlenU(SB),NOSPLIT,$0-8
	MOVW R2, x+0(FP)
	MOVW R3, y+4(FP)
	JMP runtime·goPanicSlice3AlenU(SB)

TEXT runtime·panicSlice3Acap(SB),NOSPLIT,$0-8
	MOVW R2, x+0(FP)
	MOVW R3, y+4(FP)
	JMP runtime·goPanicSlice3Acap(SB)

TEXT runtime·panicSlice3AcapU(SB),NOSPLIT,$0-8
	MOVW R2, x+0(FP)
	MOVW R3, y+4(FP)
	JMP runtime·goPanicSlice3AcapU(SB)

TEXT runtime·panicSlice3B(SB),NOSPLIT,$0-8
	MOVW R1, x+0(FP)
	MOVW R2, y+4(FP)
	JMP runtime·goPanicSlice3B(SB)

TEXT runtime·panicSlice3BU(SB),NOSPLIT,$0-8
	MOVW R1, x+0(FP)
	MOVW R2, y+4(FP)
	JMP runtime·goPanicSlice3BU(SB)

TEXT runtime·panicSlice3C(SB),NOSPLIT,$0-8
	MOVW R0, x+0(FP)
	MOVW R1, y+4(FP)
	JMP runtime·goPanicSlice3C(SB)

TEXT runtime·panicSlice3CU(SB),NOSPLIT,$0-8
	MOVW R0, x+0(FP)
	MOVW R1, y+4(FP)
	JMP runtime·goPanicSlice3CU(SB)

TEXT runtime·panicSliceConvert(SB),NOSPLIT,$0-8
	MOVW R2, x+0(FP)
	MOVW R3, y+4(FP)
	JMP runtime·goPanicSliceConvert(SB)

// Extended versions for 64-bit indexes.
TEXT runtime·panicExtendIndex(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R0, lo+4(FP)
	MOVW R1, y+8(FP)
	JMP runtime·goPanicExtendIndex(SB)

TEXT runtime·panicExtendIndexU(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R0, lo+4(FP)
	MOVW R1, y+8(FP)
	JMP runtime·goPanicExtendIndexU(SB)

TEXT runtime·panicExtendSliceAlen(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R1, lo+4(FP)
	MOVW R2, y+8(FP)
	JMP runtime·goPanicExtendSliceAlen(SB)

TEXT runtime·panicExtendSliceAlenU(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R1, lo+4(FP)
	MOVW R2, y+8(FP)
	JMP runtime·goPanicExtendSliceAlenU(SB)

TEXT runtime·panicExtendSliceAcap(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R1, lo+4(FP)
	MOVW R2, y+8(FP)
	JMP runtime·goPanicExtendSliceAcap(SB)

TEXT runtime·panicExtendSliceAcapU(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R1, lo+4(FP)
	MOVW R2, y+8(FP)
	JMP runtime·goPanicExtendSliceAcapU(SB)

TEXT runtime·panicExtendSliceB(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R0, lo+4(FP)
	MOVW R1, y+8(FP)
	JMP runtime·goPanicExtendSliceB(SB)

TEXT runtime·panicExtendSliceBU(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R0, lo+4(FP)
	MOVW R1, y+8(FP)
	JMP runtime·goPanicExtendSliceBU(SB)

TEXT runtime·panicExtendSlice3Alen(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R2, lo+4(FP)
	MOVW R3, y+8(FP)
	JMP runtime·goPanicExtendSlice3Alen(SB)

TEXT runtime·panicExtendSlice3AlenU(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R2, lo+4(FP)
	MOVW R3, y+8(FP)
	JMP runtime·goPanicExtendSlice3AlenU(SB)

TEXT runtime·panicExtendSlice3Acap(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R2, lo+4(FP)
	MOVW R3, y+8(FP)
	JMP runtime·goPanicExtendSlice3Acap(SB)

TEXT runtime·panicExtendSlice3AcapU(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R2, lo+4(FP)
	MOVW R3, y+8(FP)
	JMP runtime·goPanicExtendSlice3AcapU(SB)

TEXT runtime·panicExtendSlice3B(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R1, lo+4(FP)
	MOVW R2, y+8(FP)
	JMP runtime·goPanicExtendSlice3B(SB)

TEXT runtime·panicExtendSlice3BU(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R1, lo+4(FP)
	MOVW R2, y+8(FP)
	JMP runtime·goPanicExtendSlice3BU(SB)

TEXT runtime·panicExtendSlice3C(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R0, lo+4(FP)
	MOVW R1, y+8(FP)
	JMP runtime·goPanicExtendSlice3C(SB)

TEXT runtime·panicExtendSlice3CU(SB),NOSPLIT,$0-12
	MOVW R4, hi+0(FP)
	MOVW R0, lo+4(FP)
	MOVW R1, y+8(FP)
	JMP runtime·goPanicExtendSlice3CU(SB)

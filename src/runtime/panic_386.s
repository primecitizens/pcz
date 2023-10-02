// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && 386

#include "textflag.h"

// Note: these functions use a special calling convention to save generated code space.
// Arguments are passed in registers, but the space for those arguments are allocated
// in the caller's stack frame. These stubs write the args into that stack space and
// then tail call to the corresponding runtime handler.
// The tail call makes these stubs disappear in backtraces.

TEXT runtim·panicIndex(SB),NOSPLIT,$0-8
	MOVL AX, x+0(FP)
	MOVL CX, y+4(FP)
	JMP runtime·goPanicIndex(SB)

TEXT runtim·panicIndexU(SB),NOSPLIT,$0-8
	MOVL AX, x+0(FP)
	MOVL CX, y+4(FP)
	JMP runtime·goPanicIndexU(SB)

TEXT runtim·panicSliceAlen(SB),NOSPLIT,$0-8
	MOVL CX, x+0(FP)
	MOVL DX, y+4(FP)
	JMP runtime·goPanicSliceAlen(SB)

TEXT runtim·panicSliceAlenU(SB),NOSPLIT,$0-8
	MOVL CX, x+0(FP)
	MOVL DX, y+4(FP)
	JMP runtime·goPanicSliceAlenU(SB)

TEXT runtim·panicSliceAcap(SB),NOSPLIT,$0-8
	MOVL CX, x+0(FP)
	MOVL DX, y+4(FP)
	JMP runtime·goPanicSliceAcap(SB)

TEXT runtim·panicSliceAcapU(SB),NOSPLIT,$0-8
	MOVL CX, x+0(FP)
	MOVL DX, y+4(FP)
	JMP runtime·goPanicSliceAcapU(SB)

TEXT runtim·panicSliceB(SB),NOSPLIT,$0-8
	MOVL AX, x+0(FP)
	MOVL CX, y+4(FP)
	JMP runtime·goPanicSliceB(SB)

TEXT runtim·panicSliceBU(SB),NOSPLIT,$0-8
	MOVL AX, x+0(FP)
	MOVL CX, y+4(FP)
	JMP runtime·goPanicSliceBU(SB)

TEXT runtim·panicSlice3Alen(SB),NOSPLIT,$0-8
	MOVL DX, x+0(FP)
	MOVL BX, y+4(FP)
	JMP runtime·goPanicSlice3Alen(SB)

TEXT runtim·panicSlice3AlenU(SB),NOSPLIT,$0-8
	MOVL DX, x+0(FP)
	MOVL BX, y+4(FP)
	JMP runtime·goPanicSlice3AlenU(SB)

TEXT runtim·panicSlice3Acap(SB),NOSPLIT,$0-8
	MOVL DX, x+0(FP)
	MOVL BX, y+4(FP)
	JMP runtime·goPanicSlice3Acap(SB)

TEXT runtim·panicSlice3AcapU(SB),NOSPLIT,$0-8
	MOVL DX, x+0(FP)
	MOVL BX, y+4(FP)
	JMP runtime·goPanicSlice3AcapU(SB)

TEXT runtim·panicSlice3B(SB),NOSPLIT,$0-8
	MOVL CX, x+0(FP)
	MOVL DX, y+4(FP)
	JMP runtime·goPanicSlice3B(SB)

TEXT runtim·panicSlice3BU(SB),NOSPLIT,$0-8
	MOVL CX, x+0(FP)
	MOVL DX, y+4(FP)
	JMP runtime·goPanicSlice3BU(SB)

TEXT runtim·panicSlice3C(SB),NOSPLIT,$0-8
	MOVL AX, x+0(FP)
	MOVL CX, y+4(FP)
	JMP runtime·goPanicSlice3C(SB)

TEXT runtim·panicSlice3CU(SB),NOSPLIT,$0-8
	MOVL AX, x+0(FP)
	MOVL CX, y+4(FP)
	JMP runtime·goPanicSlice3CU(SB)

TEXT runtim·panicSliceConvert(SB),NOSPLIT,$0-8
	MOVL DX, x+0(FP)
	MOVL BX, y+4(FP)
	JMP runtime·goPanicSliceConvert(SB)

// Extended versions for 64-bit indexes.
TEXT runtim·panicExtendIndex(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL AX, lo+4(FP)
	MOVL CX, y+8(FP)
	JMP runtime·goPanicExtendIndex(SB)

TEXT runtim·panicExtendIndexU(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL AX, lo+4(FP)
	MOVL CX, y+8(FP)
	JMP runtime·goPanicExtendIndexU(SB)

TEXT runtim·panicExtendSliceAlen(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL CX, lo+4(FP)
	MOVL DX, y+8(FP)
	JMP runtime·goPanicExtendSliceAlen(SB)

TEXT runtim·panicExtendSliceAlenU(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL CX, lo+4(FP)
	MOVL DX, y+8(FP)
	JMP runtime·goPanicExtendSliceAlenU(SB)

TEXT runtim·panicExtendSliceAcap(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL CX, lo+4(FP)
	MOVL DX, y+8(FP)
	JMP runtime·goPanicExtendSliceAcap(SB)

TEXT runtim·panicExtendSliceAcapU(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL CX, lo+4(FP)
	MOVL DX, y+8(FP)
	JMP runtime·goPanicExtendSliceAcapU(SB)

TEXT runtim·panicExtendSliceB(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL AX, lo+4(FP)
	MOVL CX, y+8(FP)
	JMP runtime·goPanicExtendSliceB(SB)

TEXT runtim·panicExtendSliceBU(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL AX, lo+4(FP)
	MOVL CX, y+8(FP)
	JMP runtime·goPanicExtendSliceBU(SB)

TEXT runtim·panicExtendSlice3Alen(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL DX, lo+4(FP)
	MOVL BX, y+8(FP)
	JMP runtime·goPanicExtendSlice3Alen(SB)

TEXT runtim·panicExtendSlice3AlenU(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL DX, lo+4(FP)
	MOVL BX, y+8(FP)
	JMP runtime·goPanicExtendSlice3AlenU(SB)

TEXT runtim·panicExtendSlice3Acap(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL DX, lo+4(FP)
	MOVL BX, y+8(FP)
	JMP runtime·goPanicExtendSlice3Acap(SB)

TEXT runtim·panicExtendSlice3AcapU(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL DX, lo+4(FP)
	MOVL BX, y+8(FP)
	JMP runtime·goPanicExtendSlice3AcapU(SB)

TEXT runtim·panicExtendSlice3B(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL CX, lo+4(FP)
	MOVL DX, y+8(FP)
	JMP runtime·goPanicExtendSlice3B(SB)

TEXT runtim·panicExtendSlice3BU(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL CX, lo+4(FP)
	MOVL DX, y+8(FP)
	JMP runtime·goPanicExtendSlice3BU(SB)

TEXT runtim·panicExtendSlice3C(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL AX, lo+4(FP)
	MOVL CX, y+8(FP)
	JMP runtime·goPanicExtendSlice3C(SB)

TEXT runtim·panicExtendSlice3CU(SB),NOSPLIT,$0-12
	MOVL SI, hi+0(FP)
	MOVL AX, lo+4(FP)
	MOVL CX, y+8(FP)
	JMP runtime·goPanicExtendSlice3CU(SB)

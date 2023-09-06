// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ppc64 || ppc64le

#include "textflag.h"

// Note: these functions use a special calling convention to save generated code space.
// Arguments are passed in registers, but the space for those arguments are allocated
// in the caller's stack frame. These stubs write the args into that stack space and
// then tail call to the corresponding runtime handler.
// The tail call makes these stubs disappear in backtraces.

TEXT runtime·panicIndex<ABIInternal>(SB),NOSPLIT,$0-16
	JMP runtime·goPanicIndex<ABIInternal>(SB)

TEXT runtime·panicIndexU<ABIInternal>(SB),NOSPLIT,$0-16
	JMP runtime·goPanicIndexU<ABIInternal>(SB)

TEXT runtime·panicSliceAlen<ABIInternal>(SB),NOSPLIT,$0-16
	MOVD R4, R3
	MOVD R5, R4
	JMP runtime·goPanicSliceAlen<ABIInternal>(SB)

TEXT runtime·panicSliceAlenU<ABIInternal>(SB),NOSPLIT,$0-16
	MOVD R4, R3
	MOVD R5, R4
	JMP runtime·goPanicSliceAlenU<ABIInternal>(SB)

TEXT runtime·panicSliceAcap<ABIInternal>(SB),NOSPLIT,$0-16
	MOVD R4, R3
	MOVD R5, R4
	JMP runtime·goPanicSliceAcap<ABIInternal>(SB)

TEXT runtime·panicSliceAcapU<ABIInternal>(SB),NOSPLIT,$0-16
	MOVD R4, R3
	MOVD R5, R4
	JMP runtime·goPanicSliceAcapU<ABIInternal>(SB)

TEXT runtime·panicSliceB<ABIInternal>(SB),NOSPLIT,$0-16
	JMP runtime·goPanicSliceB<ABIInternal>(SB)

TEXT runtime·panicSliceBU<ABIInternal>(SB),NOSPLIT,$0-16
	JMP runtime·goPanicSliceBU<ABIInternal>(SB)

TEXT runtime·panicSlice3Alen<ABIInternal>(SB),NOSPLIT,$0-16
	MOVD R5, R3
	MOVD R6, R4
	JMP runtime·goPanicSlice3Alen<ABIInternal>(SB)

TEXT runtime·panicSlice3AlenU<ABIInternal>(SB),NOSPLIT,$0-16
	MOVD R5, R3
	MOVD R6, R4
	JMP runtime·goPanicSlice3AlenU<ABIInternal>(SB)

TEXT runtime·panicSlice3Acap<ABIInternal>(SB),NOSPLIT,$0-16
	MOVD R5, R3
	MOVD R6, R4
	JMP runtime·goPanicSlice3Acap<ABIInternal>(SB)

TEXT runtime·panicSlice3AcapU<ABIInternal>(SB),NOSPLIT,$0-16
	MOVD R5, R3
	MOVD R6, R4
	JMP runtime·goPanicSlice3AcapU<ABIInternal>(SB)

TEXT runtime·panicSlice3B<ABIInternal>(SB),NOSPLIT,$0-16
	MOVD R4, R3
	MOVD R5, R4
	JMP runtime·goPanicSlice3B<ABIInternal>(SB)

TEXT runtime·panicSlice3BU<ABIInternal>(SB),NOSPLIT,$0-16
	MOVD R4, R3
	MOVD R5, R4
	JMP runtime·goPanicSlice3BU<ABIInternal>(SB)

TEXT runtime·panicSlice3C<ABIInternal>(SB),NOSPLIT,$0-16
	JMP runtime·goPanicSlice3C<ABIInternal>(SB)

TEXT runtime·panicSlice3CU<ABIInternal>(SB),NOSPLIT,$0-16
	JMP runtime·goPanicSlice3CU<ABIInternal>(SB)

TEXT runtime·panicSliceConvert<ABIInternal>(SB),NOSPLIT,$0-16
	MOVD R5, R3
	MOVD R6, R4
	JMP runtime·goPanicSliceConvert<ABIInternal>(SB)

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm

#include "textflag.h"
#include "funcdata.h"

// armPublicationBarrier is a native store/store barrier for ARMv7+.
// On earlier ARM revisions, armPublicationBarrier is a no-op.
// This will not work on SMP ARMv6 machines, if any are in use.
// To implement publicationBarrier in sys_$GOOS_arm.s using the native
// instructions, use:
//
//	TEXT ·PublicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
//		B runtime·armPublicationBarrier(SB)
//
TEXT ·PublicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	MOVB runtime·goarm(SB), R11
	CMP $7, R11
	BLT 2(PC)
	DMB MB_ST
	RET

// bool armcas(int32 *val, int32 old, int32 new)
//
// To implement ·cas in sys_$GOOS_arm.s
// using the native instructions, use:
//
//	TEXT ·cas(SB),NOSPLIT,$0
//		B ·armcas(SB)
//
TEXT ·armcas(SB),NOSPLIT,$0-13
	MOVW ptr+0(FP), R1
	MOVW old+4(FP), R2
	MOVW new+8(FP), R3
casl:
	LDREX (R1), R0
	CMP R0, R2
	BNE casfail

	MOVB runtime·goarm(SB), R8
	CMP $7, R8
	BLT 2(PC)
	DMB MB_ISHST

	STREX R3, (R1), R0
	CMP $0, R0
	BNE casl
	MOVW $1, R0

	CMP $7, R8
	BLT 2(PC)
	DMB MB_ISH

	MOVB R0, ret+12(FP)
	RET
casfail:
	MOVW $0, R0
	MOVB R0, ret+12(FP)
	RET

// 64-bit atomics
// The native ARM implementations use LDREXD/STREXD, which are
// available on ARMv6k or later. We use them only on ARMv7.
// On older ARM, we use Go implementations which simulate 64-bit
// atomics with locks.
TEXT armCas64<>(SB),NOSPLIT,$0-21
	// addr is already in R1
	MOVW old_lo+4(FP), R2
	MOVW old_hi+8(FP), R3
	MOVW new_lo+12(FP), R4
	MOVW new_hi+16(FP), R5
cas64loop:
	LDREXD (R1), R6 // loads R6 and R7
	CMP R2, R6
	BNE cas64fail
	CMP R3, R7
	BNE cas64fail

	DMB MB_ISHST

	STREXD R4, (R1), R0 // stores R4 and R5
	CMP $0, R0
	BNE cas64loop
	MOVW $1, R0

	DMB MB_ISH

	MOVBU R0, swapped+20(FP)
	RET
cas64fail:
	MOVW $0, R0
	MOVBU R0, swapped+20(FP)
	RET

TEXT armAdd64<>(SB),NOSPLIT,$0-20
	// addr is already in R1
	MOVW delta_lo+4(FP), R2
	MOVW delta_hi+8(FP), R3

add64loop:
	LDREXD (R1), R4 // loads R4 and R5
	ADD.S R2, R4
	ADC R3, R5

	DMB MB_ISHST

	STREXD R4, (R1), R0 // stores R4 and R5
	CMP $0, R0
	BNE add64loop

	DMB MB_ISH

	MOVW R4, new_lo+12(FP)
	MOVW R5, new_hi+16(FP)
	RET

TEXT armSwap64<>(SB),NOSPLIT,$0-20
	// addr is already in R1
	MOVW new_lo+4(FP), R2
	MOVW new_hi+8(FP), R3

swap64loop:
	LDREXD (R1), R4 // loads R4 and R5

	DMB MB_ISHST

	STREXD R2, (R1), R0 // stores R2 and R3
	CMP $0, R0
	BNE swap64loop

	DMB MB_ISH

	MOVW R4, old_lo+12(FP)
	MOVW R5, old_hi+16(FP)
	RET

TEXT armLoad64<>(SB),NOSPLIT,$0-12
	// addr is already in R1

	LDREXD (R1), R2 // loads R2 and R3
	DMB MB_ISH

	MOVW R2, val_lo+4(FP)
	MOVW R3, val_hi+8(FP)
	RET

TEXT armStore64<>(SB),NOSPLIT,$0-12
	// addr is already in R1
	MOVW val_lo+4(FP), R2
	MOVW val_hi+8(FP), R3

store64loop:
	LDREXD (R1), R4 // loads R4 and R5

	DMB MB_ISHST

	STREXD R2, (R1), R0 // stores R2 and R3
	CMP $0, R0
	BNE store64loop

	DMB MB_ISH
	RET

// The following functions all panic if their address argument isn't
// 8-byte aligned. Since we're calling back into Go code to do this,
// we have to cooperate with stack unwinding. In the normal case, the
// functions tail-call into the appropriate implementation, which
// means they must not open a frame. Hence, when they go down the
// panic path, at that point they push the LR to create a real frame
// (they don't need to pop it because panic won't return; however, we
// do need to set the SP delta back).

// Check if R1 is 8-byte aligned, panic if not.
// Clobbers R2.
#define CHECK_ALIGN \
	AND.S $7, R1, R2 \
	BEQ 	4(PC) \
	MOVW.W R14, -4(R13) /* prepare a real frame */ \
	BL ·panicUnaligned(SB) \
	ADD $4, R13 /* compensate SP delta */

//
// Store
//

TEXT ·Store64(SB),NOSPLIT,$-4-12
	NO_LOCAL_POINTERS
	MOVW addr+0(FP), R1
	CHECK_ALIGN

	MOVB runtime·goarm(SB), R11
	CMP $7, R11
	BLT 2(PC)
	JMP armStore64<>(SB)
	JMP ·goStore64(SB)

TEXT ·StoreUintptr(SB),NOSPLIT,$0-8
	B ·Store(SB)

TEXT ·StorePointer(SB),NOSPLIT,$0-8
	B ·Store(SB)

TEXT ·StoreInt32(SB),NOSPLIT,$0-8
	B ·Store(SB)

TEXT ·StoreInt64(SB),NOSPLIT,$0-12
	B ·Store64(SB)

//
// StoreRel
//

TEXT ·StoreRel32(SB),NOSPLIT,$0-8
	B ·Store(SB)

TEXT ·StoreRelUintptr(SB),NOSPLIT,$0-8
	B ·Store(SB)

//
// Load
//

TEXT ·Load64(SB),NOSPLIT,$-4-12
	NO_LOCAL_POINTERS
	MOVW addr+0(FP), R1
	CHECK_ALIGN

	MOVB runtime·goarm(SB), R11
	CMP $7, R11
	BLT 2(PC)
	JMP armLoad64<>(SB)
	JMP ·goLoad64(SB)

TEXT ·LoadUintptr(SB),NOSPLIT,$0-8
	B ·Load32(SB)

TEXT ·LoadPointer(SB),NOSPLIT|NOFRAME,$0-8
	B ·Load32(SB)

TEXT ·LoadUint(SB),NOSPLIT,$0-8
	B ·Load32(SB)

TEXT ·LoadInt32(SB),NOSPLIT,$0-8
	B ·Load32(SB)

TEXT ·LoadInt64(SB),NOSPLIT,$-4-12
	B ·Load64(SB)

//
// LoadAcq
//

TEXT ·LoadAcq32(SB),NOSPLIT|NOFRAME,$0-8
	B ·Load32(SB)

TEXT ·LoadAcqUintptr(SB),NOSPLIT|NOFRAME,$0-8
	B 	·Load32(SB)

//
// bitwise
//

//
// Swap
//

TEXT ·Swap64(SB),NOSPLIT,$-4-20
	NO_LOCAL_POINTERS
	MOVW addr+0(FP), R1
	CHECK_ALIGN

	MOVB runtime·goarm(SB), R11
	CMP $7, R11
	BLT 2(PC)
	JMP armSwap64<>(SB)
	JMP ·goSwap64(SB)

TEXT ·SwapInt32(SB),NOSPLIT,$0-12
	B ·Swap32(SB)

TEXT ·SwapInt64(SB),NOSPLIT,$-4-20
	B ·Swap64(SB)

//
// Add
//

TEXT ·Add64(SB),NOSPLIT,$-4-20
	NO_LOCAL_POINTERS
	MOVW addr+0(FP), R1
	CHECK_ALIGN

	MOVB runtime·goarm(SB), R11
	CMP $7, R11
	BLT 2(PC)
	JMP armAdd64<>(SB)
	JMP ·goAdd64(SB)

TEXT ·AddUintptr(SB),NOSPLIT,$0-12
	B ·Add32(SB)

TEXT ·AddInt32(SB),NOSPLIT,$0-12
	B ·Add32(SB)

TEXT ·AddInt64(SB),NOSPLIT,$-4-20
	B ·Add64(SB)

//
// Compare and swap
//

TEXT ·Cas64(SB),NOSPLIT,$-4-21
	NO_LOCAL_POINTERS
	MOVW addr+0(FP), R1
	CHECK_ALIGN

	MOVB runtime·goarm(SB), R11
	CMP $7, R11
	BLT 2(PC)
	JMP armCas64<>(SB)
	JMP ·goCas64(SB)

TEXT ·CasInt32(SB),NOSPLIT,$0-13
	B ·Cas(SB)

TEXT ·CasInt64(SB),NOSPLIT,$-4-21
	B ·Cas64(SB)

TEXT ·CasUintptr(SB),NOSPLIT,$0-13
	B ·Cas(SB)

TEXT ·CasUnsafePointer(SB),NOSPLIT,$0-13
	B ·Cas(SB)

TEXT ·CasRel32(SB),NOSPLIT,$0-13
	B ·Cas(SB)

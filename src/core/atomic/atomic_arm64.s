// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_asm.h"
#include "textflag.h"

#define HAS_ATOMICS_R4 \
	MOVW ·useARM64Atomics(SB), R4

TEXT ·PublicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	DMB $0xe // DMB ST
	RET

//
// Store
//

TEXT ·Store8(SB),NOSPLIT,$0-9
	MOVD ptr+0(FP), R0
	MOVB val+8(FP), R1
	STLRB R1, (R0)
	RET

TEXT ·Store32(SB),NOSPLIT,$0-12
	MOVD ptr+0(FP), R0
	MOVW val+8(FP), R1
	STLRW R1, (R0)
	RET

TEXT ·Store64(SB),NOSPLIT,$0-16
	MOVD ptr+0(FP), R0
	MOVD val+8(FP), R1
	STLR R1, (R0)
	RET

TEXT ·StoreUintptr(SB),NOSPLIT,$0-16
	B ·Store64(SB)

TEXT ·StorePointer(SB),NOSPLIT,$0-16
	B ·Store64(SB)

TEXT ·StoreInt32(SB),NOSPLIT,$0-12
	B ·Store32(SB)

TEXT ·StoreInt64(SB),NOSPLIT,$0-16
	B ·Store64(SB)

//
// StoreRel
//


TEXT ·StoreRelUintptr(SB),NOSPLIT,$0-16
	B ·Store64(SB)

TEXT ·StoreRel32(SB),NOSPLIT,$0-12
	B ·Store32(SB)

TEXT ·StoreRel64(SB),NOSPLIT,$0-16
	B ·Store64(SB)

//
// Load
//

// uint8 ·Load8(uint8 volatile* addr)
TEXT ·Load8(SB),NOSPLIT,$0-9
	MOVD ptr+0(FP), R0
	LDARB (R0), R0
	MOVB R0, ret+8(FP)
	RET

// uint32 ·Load32(uint32 volatile* addr)
TEXT ·Load32(SB),NOSPLIT,$0-12
	MOVD ptr+0(FP), R0
	LDARW (R0), R0
	MOVW R0, ret+8(FP)
	RET


// uint64 ·Load64(uint64 volatile* addr)
TEXT ·Load64(SB),NOSPLIT,$0-16
	MOVD ptr+0(FP), R0
	LDAR (R0), R0
	MOVD R0, ret+8(FP)
	RET

TEXT ·LoadUintptr(SB),NOSPLIT,$0-16
	B ·Load64(SB)

// void *·LoadPointer(void *volatile *addr)
TEXT ·LoadPointer(SB),NOSPLIT,$0-16
	MOVD ptr+0(FP), R0
	LDAR (R0), R0
	MOVD R0, ret+8(FP)
	RET

TEXT ·LoadUint(SB),NOSPLIT,$0-16
	B ·Load64(SB)

TEXT ·LoadInt32(SB),NOSPLIT,$0-12
	B ·Load32(SB)

TEXT ·LoadInt64(SB),NOSPLIT,$0-16
	B ·Load64(SB)

//
// LoadAcq
//

// uint32 ·LoadAcq32(uint32 volatile* addr)
TEXT ·LoadAcq32(SB),NOSPLIT,$0-12
	B ·Load32(SB)

// uint64 ·LoadAcqUintptr(uint64 volatile* addr)
TEXT ·LoadAcq64(SB),NOSPLIT,$0-16
	B ·Load64(SB)

// uintptr ·LoadAcq64(uintptr volatile* addr)
TEXT ·LoadAcqUintptr(SB),NOSPLIT,$0-16
	B ·Load64(SB)

//
// bitwise
//

TEXT ·And8(SB),NOSPLIT,$0-9
	MOVD ptr+0(FP), R0
	MOVB val+8(FP), R1
	HAS_ATOMICS_R4
	CBZ 	R4, load_store_loop
	MVN 	R1, R2
	LDCLRALB R2, (R0), R3
	RET
load_store_loop:
	LDAXRB (R0), R2
	AND R1, R2
	STLXRB R2, (R0), R3
	CBNZ R3, load_store_loop
	RET

// func And32(addr *uint32, v uint32)
TEXT ·And32(SB),NOSPLIT,$0-12
	MOVD ptr+0(FP), R0
	MOVW val+8(FP), R1
	HAS_ATOMICS_R4
	CBZ 	R4, load_store_loop
	MVN 	R1, R2
	LDCLRALW R2, (R0), R3
	RET
load_store_loop:
	LDAXRW (R0), R2
	AND R1, R2
	STLXRW R2, (R0), R3
	CBNZ R3, load_store_loop
	RET

TEXT ·Or8(SB),NOSPLIT,$0-9
	MOVD ptr+0(FP), R0
	MOVB val+8(FP), R1
	HAS_ATOMICS_R4
	CBZ 	R4, load_store_loop
	LDORALB R1, (R0), R2
	RET
load_store_loop:
	LDAXRB (R0), R2
	ORR R1, R2
	STLXRB R2, (R0), R3
	CBNZ R3, load_store_loop
	RET

// func Or32(addr *uint32, v uint32)
TEXT ·Or32(SB),NOSPLIT,$0-12
	MOVD ptr+0(FP), R0
	MOVW val+8(FP), R1
	HAS_ATOMICS_R4
	CBZ 	R4, load_store_loop
	LDORALW R1, (R0), R2
	RET
load_store_loop:
	LDAXRW (R0), R2
	ORR R1, R2
	STLXRW R2, (R0), R3
	CBNZ R3, load_store_loop
	RET

//
// Swap
//

// uint32 Swap32(ptr *uint32, new uint32)
TEXT ·Swap32(SB),NOSPLIT,$0-20
	MOVD ptr+0(FP), R0
	MOVW new+8(FP), R1
	HAS_ATOMICS_R4
	CBZ 	R4, load_store_loop
	SWPALW R1, (R0), R2
	MOVW R2, ret+16(FP)
	RET
load_store_loop:
	LDAXRW (R0), R2
	STLXRW R1, (R0), R3
	CBNZ R3, load_store_loop
	MOVW R2, ret+16(FP)
	RET

// uint64 Swap64(ptr *uint64, new uint64)
TEXT ·Swap64(SB),NOSPLIT,$0-24
	MOVD ptr+0(FP), R0
	MOVD new+8(FP), R1
	HAS_ATOMICS_R4
	CBZ 	R4, load_store_loop
	SWPALD R1, (R0), R2
	MOVD R2, ret+16(FP)
	RET
load_store_loop:
	LDAXR (R0), R2
	STLXR R1, (R0), R3
	CBNZ R3, load_store_loop
	MOVD R2, ret+16(FP)
	RET

TEXT ·SwapUintptr(SB),NOSPLIT,$0-24
	B ·Swap64(SB)

TEXT ·SwapInt32(SB),NOSPLIT,$0-20
	B ·Swap32(SB)

TEXT ·SwapInt64(SB),NOSPLIT,$0-24
	B ·Swap64(SB)

//
// Add
//

// uint32 Add32(uint32 volatile *ptr, int32 delta)
TEXT ·Add32(SB),NOSPLIT,$0-20
	MOVD ptr+0(FP), R0
	MOVW delta+8(FP), R1
	HAS_ATOMICS_R4
	CBZ 	R4, load_store_loop
	LDADDALW R1, (R0), R2
	ADD 	R1, R2
	MOVW R2, ret+16(FP)
	RET
load_store_loop:
	LDAXRW (R0), R2
	ADDW R2, R1, R2
	STLXRW R2, (R0), R3
	CBNZ R3, load_store_loop
	MOVW R2, ret+16(FP)
	RET

// uint64 Add64(uint64 volatile *ptr, int64 delta)
TEXT ·Add64(SB),NOSPLIT,$0-24
	MOVD ptr+0(FP), R0
	MOVD delta+8(FP), R1
	HAS_ATOMICS_R4
	CBZ 	R4, load_store_loop
	LDADDALD R1, (R0), R2
	ADD 	R1, R2
	MOVD R2, ret+16(FP)
	RET
load_store_loop:
	LDAXR (R0), R2
	ADD R2, R1, R2
	STLXR R2, (R0), R3
	CBNZ R3, load_store_loop
	MOVD R2, ret+16(FP)
	RET

TEXT ·AddUintptr(SB),NOSPLIT,$0-24
	B ·Add64(SB)

TEXT ·AddInt32(SB),NOSPLIT,$0-20
	B ·Add32(SB)

TEXT ·AddInt64(SB),NOSPLIT,$0-24
	B ·Add64(SB)

//
// Compare and swap
//

// bool Cas32(uint32 *ptr, uint32 old, uint32 new)
TEXT ·Cas32(SB),NOSPLIT,$0-17
	MOVD ptr+0(FP), R0
	MOVW old+8(FP), R1
	MOVW new+12(FP), R2
	HAS_ATOMICS_R4
	CBZ 	R4, load_store_loop
	MOVD R1, R3
	CASALW R3, (R0), R2
	CMP 	R1, R3
	CSET EQ, R0
	MOVB R0, ret+16(FP)
	RET
load_store_loop:
	LDAXRW (R0), R3
	CMPW R1, R3
	BNE ok
	STLXRW R2, (R0), R3
	CBNZ R3, load_store_loop
ok:
	CSET EQ, R0
	MOVB R0, ret+16(FP)
	RET

// bool ·Cas64(uint64 *ptr, uint64 old, uint64 new)
TEXT ·Cas64(SB),NOSPLIT,$0-25
	MOVD ptr+0(FP), R0
	MOVD old+8(FP), R1
	MOVD new+16(FP), R2
	HAS_ATOMICS_R4
	CBZ 	R4, load_store_loop
	MOVD R1, R3
	CASALD R3, (R0), R2
	CMP 	R1, R3
	CSET EQ, R0
	MOVB R0, ret+24(FP)
	RET
load_store_loop:
	LDAXR (R0), R3
	CMP R1, R3
	BNE ok
	STLXR R2, (R0), R3
	CBNZ R3, load_store_loop
ok:
	CSET EQ, R0
	MOVB R0, ret+24(FP)
	RET

TEXT ·CasUintptr(SB),NOSPLIT,$0-25
	B ·Cas64(SB)

TEXT ·CasUnsafePointer(SB),NOSPLIT,$0-25
	B ·Cas64(SB)

TEXT ·CasInt32(SB),NOSPLIT,$0-17
	B ·Cas32(SB)

TEXT ·CasInt64(SB),NOSPLIT,$0-25
	B ·Cas64(SB)

//
// CasRel
//

TEXT ·CasRel32(SB),NOSPLIT,$0-17
	B ·Cas32(SB)

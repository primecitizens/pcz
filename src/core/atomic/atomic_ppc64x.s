// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ppc64 || ppc64le

#include "textflag.h"

// For more details about how various memory models are
// enforced on POWER, the following paper provides more
// details about how they enforce C/C++ like models. This
// gives context about why the strange looking code
// sequences below work.
//
// http://www.rdrop.com/users/paulmck/scalability/paper/N2745r.2011.03.04a.html

TEXT ·PublicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	// LWSYNC is the "export" barrier recommended by Power ISA
	// v2.07 book II, appendix B.2.2.2.
	// LWSYNC is a load/load, load/store, and store/store barrier.
	LWSYNC
	RET

//
// Store
//

TEXT ·Store8(SB), NOSPLIT, $0-9
	MOVD ptr+0(FP), R3
	MOVB val+8(FP), R4
	SYNC
	MOVB R4, 0(R3)
	RET

TEXT ·Store32(SB), NOSPLIT, $0-12
	MOVD ptr+0(FP), R3
	MOVW val+8(FP), R4
	SYNC
	MOVW R4, 0(R3)
	RET

TEXT ·Store64(SB), NOSPLIT, $0-16
	MOVD ptr+0(FP), R3
	MOVD val+8(FP), R4
	SYNC
	MOVD R4, 0(R3)
	RET

TEXT ·StoreUintptr(SB), NOSPLIT, $0-16
	BR ·Store64(SB)

TEXT ·StorePointer(SB), NOSPLIT, $0-16
	BR ·Store64(SB)

TEXT ·StoreInt32(SB), NOSPLIT, $0-12
	BR ·Store(SB)

TEXT ·StoreInt64(SB), NOSPLIT, $0-16
	BR ·Store64(SB)

//
// StoreRel
//

TEXT ·StoreRel32(SB), NOSPLIT, $0-12
	MOVD ptr+0(FP), R3
	MOVW val+8(FP), R4
	LWSYNC
	MOVW R4, 0(R3)
	RET

TEXT ·StoreRel64(SB), NOSPLIT, $0-16
	MOVD ptr+0(FP), R3
	MOVD val+8(FP), R4
	LWSYNC
	MOVD R4, 0(R3)
	RET

TEXT ·StoreRelUintptr(SB), NOSPLIT, $0-16
	BR ·StoreRel64(SB)

//
// Load
//

// uint8 ·Load8(uint8 volatile* ptr)
TEXT ·Load8(SB),NOSPLIT|NOFRAME,$-8-9
	MOVD ptr+0(FP), R3
	SYNC
	MOVBZ 0(R3), R3
	CMP R3, R3, CR7
	BC 4, 30, 1(PC) // bne- cr7,0x4
	ISYNC
	MOVB R3, ret+8(FP)
	RET

// uint32 ·Load32(uint32 volatile* ptr)
TEXT ·Load32(SB),NOSPLIT|NOFRAME,$-8-12
	MOVD ptr+0(FP), R3
	SYNC
	MOVWZ 0(R3), R3
	CMPW R3, R3, CR7
	BC 4, 30, 1(PC) // bne- cr7,0x4
	ISYNC
	MOVW R3, ret+8(FP)
	RET

// uint64 ·Load64(uint64 volatile* ptr)
TEXT ·Load64(SB),NOSPLIT|NOFRAME,$-8-16
	MOVD ptr+0(FP), R3
	SYNC
	MOVD 0(R3), R3
	CMP R3, R3, CR7
	BC 4, 30, 1(PC) // bne- cr7,0x4
	ISYNC
	MOVD R3, ret+8(FP)
	RET

TEXT ·LoadUintptr(SB),  NOSPLIT|NOFRAME, $0-16
	BR ·Load64(SB)

// void *·LoadPointer(void *volatile *ptr)
TEXT ·LoadPointer(SB),NOSPLIT|NOFRAME,$-8-16
	MOVD ptr+0(FP), R3
	SYNC
	MOVD 0(R3), R3
	CMP R3, R3, CR7
	BC 4, 30, 1(PC) // bne- cr7,0x4
	ISYNC
	MOVD R3, ret+8(FP)
	RET

TEXT ·LoadUint(SB), NOSPLIT|NOFRAME, $0-16
	BR ·Load64(SB)

TEXT ·LoadInt32(SB), NOSPLIT, $0-12
	BR ·Load32(SB)

TEXT ·LoadInt64(SB), NOSPLIT, $0-16
	BR ·Load64(SB)

//
// LoadAcq
//

// uint32 ·LoadAcq32(uint32 volatile* ptr)
TEXT ·LoadAcq32(SB),NOSPLIT|NOFRAME,$-8-12
	MOVD   ptr+0(FP), R3
	MOVWZ  0(R3), R3
	CMPW   R3, R3, CR7
	BC     4, 30, 1(PC) // bne- cr7, 0x4
	ISYNC
	MOVW   R3, ret+8(FP)
	RET

// uint64 ·LoadAcq64(uint64 volatile* ptr)
TEXT ·LoadAcq64(SB),NOSPLIT|NOFRAME,$-8-16
	MOVD   ptr+0(FP), R3
	MOVD   0(R3), R3
	CMP    R3, R3, CR7
	BC     4, 30, 1(PC) // bne- cr7, 0x4
	ISYNC
	MOVD   R3, ret+8(FP)
	RET

TEXT ·LoadAcqUintptr(SB),  NOSPLIT|NOFRAME, $0-16
	BR ·LoadAcq64(SB)

//
// bitwise
//

// void ·Or8(byte volatile*, byte);
TEXT ·Or8(SB), NOSPLIT, $0-9
	MOVD ptr+0(FP), R3
	MOVBZ val+8(FP), R4
	LWSYNC
again:
	LBAR (R3), R6
	OR R4, R6
	STBCCC R6, (R3)
	BNE again
	RET

// func Or32(addr *uint32, v uint32)
TEXT ·Or32(SB), NOSPLIT, $0-12
	MOVD ptr+0(FP), R3
	MOVW val+8(FP), R4
	LWSYNC
again:
	LWAR (R3), R6
	OR R4, R6
	STWCCC R6, (R3)
	BNE again
	RET

// void ·And8(byte volatile*, byte);
TEXT ·And8(SB), NOSPLIT, $0-9
	MOVD ptr+0(FP), R3
	MOVBZ val+8(FP), R4
	LWSYNC
again:
	LBAR (R3), R6
	AND R4, R6
	STBCCC R6, (R3)
	BNE again
	RET

// func And32(addr *uint32, v uint32)
TEXT ·And32(SB), NOSPLIT, $0-12
	MOVD ptr+0(FP), R3
	MOVW val+8(FP), R4
	LWSYNC
again:
	LWAR (R3),R6
	AND R4, R6
	STWCCC R6, (R3)
	BNE again
	RET

//
// Swap
//

// uint32 Swap32(ptr *uint32, new uint32)
TEXT ·Swap32(SB), NOSPLIT, $0-20
	MOVD ptr+0(FP), R4
	MOVW new+8(FP), R5
	LWSYNC
	LWAR (R4), R3
	STWCCC R5, (R4)
	BNE -2(PC)
	ISYNC
	MOVW R3, ret+16(FP)
	RET

// uint64 Swap64(ptr *uint64, new uint64)
TEXT ·Swap64(SB), NOSPLIT, $0-24
	MOVD ptr+0(FP), R4
	MOVD new+8(FP), R5
	LWSYNC
	LDAR (R4), R3
	STDCCC R5, (R4)
	BNE -2(PC)
	ISYNC
	MOVD R3, ret+16(FP)
	RET

TEXT ·SwapUintptr(SB), NOSPLIT, $0-24
	BR ·Swap64(SB)

TEXT ·SwapInt32(SB), NOSPLIT, $0-20
	BR ·Swap32(SB)

TEXT ·SwapInt64(SB), NOSPLIT, $0-24
	BR ·Swap64(SB)

//
// Add
//

// uint32 Add32(uint32 volatile *ptr, int32 delta)
TEXT ·Add32(SB), NOSPLIT, $0-20
	MOVD ptr+0(FP), R4
	MOVW delta+8(FP), R5
	LWSYNC
	LWAR (R4), R3
	ADD R5, R3
	STWCCC R3, (R4)
	BNE -3(PC)
	MOVW R3, ret+16(FP)
	RET

// uint64 Add64(uint64 volatile *val, int64 delta)
TEXT ·Add64(SB), NOSPLIT, $0-24
	MOVD ptr+0(FP), R4
	MOVD delta+8(FP), R5
	LWSYNC
	LDAR (R4), R3
	ADD R5, R3
	STDCCC R3, (R4)
	BNE -3(PC)
	MOVD R3, ret+16(FP)
	RET

TEXT ·AddUintptr(SB), NOSPLIT, $0-24
	BR ·Add64(SB)

TEXT ·AddInt32(SB), NOSPLIT, $0-20
	BR ·Add32(SB)

TEXT ·AddInt64(SB), NOSPLIT, $0-24
	BR ·Add64(SB)

//
// Compare and swap
//

// bool Cas32(uint32 *ptr, uint32 old, uint32 new)
TEXT ·Cas32(SB), NOSPLIT, $0-17
	MOVD ptr+0(FP), R3
	MOVWZ old+8(FP), R4
	MOVWZ new+12(FP), R5
	LWSYNC
cas_again:
	LWAR (R3), R6
	CMPW R6, R4
	BNE cas_fail
	STWCCC R5, (R3)
	BNE cas_again
	MOVD $1, R3
	LWSYNC
	MOVB R3, ret+16(FP)
	RET
cas_fail:
	MOVB R0, ret+16(FP)
	RET

// bool ·Cas64(uint64 *ptr, uint64 old, uint64 new)
TEXT ·Cas64(SB), NOSPLIT, $0-25
	MOVD ptr+0(FP), R3
	MOVD old+8(FP), R4
	MOVD new+16(FP), R5
	LWSYNC
cas64_again:
	LDAR (R3), R6
	CMP R6, R4
	BNE cas64_fail
	STDCCC R5, (R3)
	BNE cas64_again
	MOVD $1, R3
	LWSYNC
	MOVB R3, ret+24(FP)
	RET
cas64_fail:
	MOVB R0, ret+24(FP)
	RET

TEXT ·CasUintptr(SB), NOSPLIT, $0-25
	BR ·Cas64(SB)

// bool CasUnsafePointer(void **val, void *old, void *new)
TEXT ·CasUnsafePointer(SB), NOSPLIT, $0-25
	BR ·Cas64(SB)

TEXT ·CasInt32(SB), NOSPLIT, $0-17
	BR ·Cas32(SB)

TEXT ·CasInt64(SB), NOSPLIT, $0-25
	BR ·Cas64(SB)

//
// CasRel
//

TEXT ·CasRel32(SB), NOSPLIT, $0-17
	MOVD    ptr+0(FP), R3
	MOVWZ   old+8(FP), R4
	MOVWZ   new+12(FP), R5
	LWSYNC
cas_again:
	LWAR    (R3), $0, R6        // 0 = Mutex release hint
	CMPW    R6, R4
	BNE     cas_fail
	STWCCC  R5, (R3)
	BNE     cas_again
	MOVD    $1, R3
	MOVB    R3, ret+16(FP)
	RET
cas_fail:
	MOVB    R0, ret+16(FP)
	RET

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Note: some of these functions are semantically inlined
// by the compiler (in src/cmd/compile/internal/gc/ssa.go).

#include "textflag.h"

TEXT ·PublicationBarrier(SB),NOSPLIT,$0-0
	// Stores are already ordered on x86, so this is just a
	// compile barrier.
	RET

TEXT ·Store32(SB), NOSPLIT, $0-12
	MOVQ ptr+0(FP), BX
	MOVL val+8(FP), AX
	XCHGL AX, 0(BX)
	RET

TEXT ·Store8(SB), NOSPLIT, $0-9
	MOVQ ptr+0(FP), BX
	MOVB val+8(FP), AX
	XCHGB AX, 0(BX)
	RET

TEXT ·Store64(SB), NOSPLIT, $0-16
	MOVQ ptr+0(FP), BX
	MOVQ val+8(FP), AX
	XCHGQ AX, 0(BX)
	RET

TEXT ·StoreUintptr(SB), NOSPLIT, $0-16
	JMP ·Store64(SB)

TEXT ·StorePointer(SB), NOSPLIT, $0-16
	MOVQ ptr+0(FP), BX
	MOVQ val+8(FP), AX
	XCHGQ AX, 0(BX)
	RET

TEXT ·StoreInt32(SB), NOSPLIT, $0-12
	JMP ·Store32(SB)

TEXT ·StoreInt64(SB), NOSPLIT, $0-16
	JMP ·Store64(SB)

//
// StoreRel
//

TEXT ·StoreRel32(SB), NOSPLIT, $0-12
	JMP ·Store32(SB)

TEXT ·StoreRel64(SB), NOSPLIT, $0-16
	JMP ·Store64(SB)

TEXT ·StoreRelUintptr(SB), NOSPLIT, $0-16
	JMP ·Store64(SB)

//
// Load
//

TEXT ·LoadUintptr(SB), NOSPLIT, $0-16
	JMP ·Load64(SB)

TEXT ·LoadUint(SB), NOSPLIT, $0-16
	JMP ·Load64(SB)

TEXT ·LoadInt32(SB), NOSPLIT, $0-12
	JMP ·Load32(SB)

TEXT ·LoadInt64(SB), NOSPLIT, $0-16
	JMP ·Load64(SB)

//
// bitwise
//

// void ·Or8(byte volatile*, byte);
TEXT ·Or8(SB), NOSPLIT, $0-9
	MOVQ ptr+0(FP), AX
	MOVB val+8(FP), BX
	LOCK
	ORB BX, (AX)
	RET

// func Or32(addr *uint32, v uint32)
TEXT ·Or32(SB), NOSPLIT, $0-12
	MOVQ ptr+0(FP), AX
	MOVL val+8(FP), BX
	LOCK
	ORL BX, (AX)
	RET

// void ·And8(byte volatile*, byte);
TEXT ·And8(SB), NOSPLIT, $0-9
	MOVQ ptr+0(FP), AX
	MOVB val+8(FP), BX
	LOCK
	ANDB BX, (AX)
	RET

// func And32(addr *uint32, v uint32)
TEXT ·And32(SB), NOSPLIT, $0-12
	MOVQ ptr+0(FP), AX
	MOVL val+8(FP), BX
	LOCK
	ANDL BX, (AX)
	RET

//
// Swap
//

// uint32 Swap32(ptr *uint32, new uint32)
TEXT ·Swap32(SB), NOSPLIT, $0-20
	MOVQ ptr+0(FP), BX
	MOVL new+8(FP), AX
	XCHGL AX, 0(BX)
	MOVL AX, ret+16(FP)
	RET

// uint64 Swap64(ptr *uint64, new uint64)
TEXT ·Swap64(SB), NOSPLIT, $0-24
	MOVQ ptr+0(FP), BX
	MOVQ new+8(FP), AX
	XCHGQ AX, 0(BX)
	MOVQ AX, ret+16(FP)
	RET

TEXT ·SwapUintptr(SB), NOSPLIT, $0-24
	JMP ·Swap64(SB)

TEXT ·SwapInt32(SB), NOSPLIT, $0-20
	JMP ·Swap32(SB)

TEXT ·SwapInt64(SB), NOSPLIT, $0-24
	JMP ·Swap64(SB)

//
// Add
//

// uint32 Add32(uint32 volatile *val, int32 delta)
TEXT ·Add32(SB), NOSPLIT, $0-20
	MOVQ ptr+0(FP), BX
	MOVL delta+8(FP), AX
	MOVL AX, CX
	LOCK
	XADDL AX, 0(BX)
	ADDL CX, AX
	MOVL AX, ret+16(FP)
	RET

// uint64 Add64(uint64 volatile *val, int64 delta)
TEXT ·Add64(SB), NOSPLIT, $0-24
	MOVQ ptr+0(FP), BX
	MOVQ delta+8(FP), AX
	MOVQ AX, CX
	LOCK
	XADDQ AX, 0(BX)
	ADDQ CX, AX
	MOVQ AX, ret+16(FP)
	RET

TEXT ·AddUintptr(SB), NOSPLIT, $0-24
	JMP ·Add64(SB)

TEXT ·AddInt32(SB), NOSPLIT, $0-20
	JMP ·Add32(SB)

TEXT ·AddInt64(SB), NOSPLIT, $0-24
	JMP ·Add64(SB)

//
// Compare and swap
//

// bool Cas32(int32 *val, int32 old, int32 new)
TEXT ·Cas32(SB),NOSPLIT,$0-17
	MOVQ ptr+0(FP), BX
	MOVL old+8(FP), AX
	MOVL new+12(FP), CX
	LOCK
	CMPXCHGL CX, 0(BX)
	SETEQ ret+16(FP)
	RET

// bool ·Cas64(uint64 *val, uint64 old, uint64 new)
TEXT ·Cas64(SB), NOSPLIT, $0-25
	MOVQ ptr+0(FP), BX
	MOVQ old+8(FP), AX
	MOVQ new+16(FP), CX
	LOCK
	CMPXCHGQ CX, 0(BX)
	SETEQ ret+24(FP)
	RET

TEXT ·CasUintptr(SB), NOSPLIT, $0-25
	JMP ·Cas64(SB)

// bool CasUnsafePointer(void **val, void *old, void *new)
TEXT ·CasUnsafePointer(SB), NOSPLIT, $0-25
	MOVQ ptr+0(FP), BX
	MOVQ old+8(FP), AX
	MOVQ new+16(FP), CX
	LOCK
	CMPXCHGQ CX, 0(BX)
	SETEQ ret+24(FP)
	RET

TEXT ·CasInt32(SB), NOSPLIT, $0-17
	JMP ·Cas32(SB)

TEXT ·CasInt64(SB), NOSPLIT, $0-25
	JMP ·Cas64(SB)

//
// CasRel
//

TEXT ·CasRel32(SB), NOSPLIT, $0-17
	JMP ·Cas32(SB)

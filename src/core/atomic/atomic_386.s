// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386

#include "textflag.h"
#include "funcdata.h"

TEXT ·PublicationBarrier(SB),NOSPLIT,$0-0
	// Stores are already ordered on x86, so this is just a
	// compile barrier.
	RET

//
// Store
//

TEXT ·Store8(SB), NOSPLIT, $0-5
	MOVL ptr+0(FP), BX
	MOVB val+4(FP), AX
	XCHGB AX, 0(BX)
	RET

TEXT ·Store32(SB), NOSPLIT, $0-8
	MOVL ptr+0(FP), BX
	MOVL val+4(FP), AX
	XCHGL AX, 0(BX)
	RET

// void ·Store64(uint64 volatile* addr, uint64 v);
TEXT ·Store64(SB), NOSPLIT, $0-12
	NO_LOCAL_POINTERS
	MOVL ptr+0(FP), AX
	TESTL $7, AX
	JZ 2(PC)
	CALL ·panicUnaligned(SB)
	// MOVQ and EMMS were introduced on the Pentium MMX.
	MOVQ val+4(FP), M0
	MOVQ M0, (AX)
	EMMS
	// This is essentially a no-op, but it provides required memory fencing.
	// It can be replaced with MFENCE, but MFENCE was introduced only on the Pentium4 (SSE2).
	XORL AX, AX
	LOCK
	XADDL AX, (SP)
	RET

TEXT ·StoreUintptr(SB), NOSPLIT, $0-8
	JMP ·Store32(SB)

TEXT ·StorePointer(SB), NOSPLIT, $0-8
	MOVL ptr+0(FP), BX
	MOVL val+4(FP), AX
	XCHGL AX, 0(BX)
	RET

TEXT ·StoreInt32(SB), NOSPLIT, $0-8
	JMP ·Store32(SB)

TEXT ·StoreInt64(SB), NOSPLIT, $0-12
	JMP ·Store64(SB)

//
// StoreRel
//

TEXT ·StoreRel32(SB), NOSPLIT, $0-8
	JMP ·Store32(SB)

TEXT ·StoreRelUintptr(SB), NOSPLIT, $0-8
	JMP ·Store32(SB)

//
// Load
//

// uint64 atomicload64(uint64 volatile* addr);
TEXT ·Load64(SB), NOSPLIT, $0-12
	NO_LOCAL_POINTERS
	MOVL ptr+0(FP), AX
	TESTL $7, AX
	JZ 2(PC)
	CALL ·panicUnaligned(SB)
	MOVQ (AX), M0
	MOVQ M0, ret+4(FP)
	EMMS
	RET

TEXT ·LoadUintptr(SB), NOSPLIT, $0-8
	JMP ·Load32(SB)

TEXT ·LoadUint(SB), NOSPLIT, $0-8
	JMP ·Load32(SB)

TEXT ·LoadInt32(SB), NOSPLIT, $0-8
	JMP ·Load32(SB)

TEXT ·LoadInt64(SB), NOSPLIT, $0-12
	JMP ·Load64(SB)

//
// bitwise
//

// void ·Or8(byte volatile*, byte);
TEXT ·Or8(SB), NOSPLIT, $0-5
	MOVL ptr+0(FP), AX
	MOVB val+4(FP), BX
	LOCK
	ORB BX, (AX)
	RET

// func Or32(addr *uint32, v uint32)
TEXT ·Or32(SB), NOSPLIT, $0-8
	MOVL ptr+0(FP), AX
	MOVL val+4(FP), BX
	LOCK
	ORL BX, (AX)
	RET

// void ·And8(byte volatile*, byte);
TEXT ·And8(SB), NOSPLIT, $0-5
	MOVL ptr+0(FP), AX
	MOVB val+4(FP), BX
	LOCK
	ANDB BX, (AX)
	RET

// func And32(addr *uint32, v uint32)
TEXT ·And32(SB), NOSPLIT, $0-8
	MOVL ptr+0(FP), AX
	MOVL val+4(FP), BX
	LOCK
	ANDL BX, (AX)
	RET

//
// Swap
//

TEXT ·Swap32(SB), NOSPLIT, $0-12
	MOVL ptr+0(FP), BX
	MOVL new+4(FP), AX
	XCHGL AX, 0(BX)
	MOVL AX, ret+8(FP)
	RET

TEXT ·Swap64(SB),NOSPLIT,$0-20
	NO_LOCAL_POINTERS
	// no XCHGQ so use CMPXCHG8B loop
	MOVL ptr+0(FP), BP
	TESTL $7, BP
	JZ 2(PC)
	CALL ·panicUnaligned(SB)
	// CX:BX = new
	MOVL new_lo+4(FP), BX
	MOVL new_hi+8(FP), CX
	// DX:AX = *addr
	MOVL 0(BP), AX
	MOVL 4(BP), DX
swaploop:
	// if *addr == DX:AX
	//	*addr = CX:BX
	// else
	//	DX:AX = *addr
	// all in one instruction
	LOCK
	CMPXCHG8B 0(BP)
	JNZ swaploop

	// success
	// return DX:AX
	MOVL AX, ret_lo+12(FP)
	MOVL DX, ret_hi+16(FP)
	RET

TEXT ·SwapUintptr(SB), NOSPLIT, $0-12
	JMP ·Swap32(SB)

TEXT ·SwapInt32(SB), NOSPLIT, $0-12
	JMP ·Swap32(SB)

TEXT ·SwapInt64(SB), NOSPLIT, $0-20
	JMP ·Swap64(SB)

//
// Add
//

// uint32 Add32(uint32 volatile *val, int32 delta)
TEXT ·Add32(SB), NOSPLIT, $0-12
	MOVL ptr+0(FP), BX
	MOVL delta+4(FP), AX
	MOVL AX, CX
	LOCK
	XADDL AX, 0(BX)
	ADDL CX, AX
	MOVL AX, ret+8(FP)
	RET

TEXT ·Add64(SB), NOSPLIT, $0-20
	NO_LOCAL_POINTERS
	// no XADDQ so use CMPXCHG8B loop
	MOVL ptr+0(FP), BP
	TESTL $7, BP
	JZ 2(PC)
	CALL ·panicUnaligned(SB)
	// DI:SI = delta
	MOVL delta_lo+4(FP), SI
	MOVL delta_hi+8(FP), DI
	// DX:AX = *addr
	MOVL 0(BP), AX
	MOVL 4(BP), DX
addloop:
	// CX:BX = DX:AX (*addr) + DI:SI (delta)
	MOVL AX, BX
	MOVL DX, CX
	ADDL SI, BX
	ADCL DI, CX

	// if *addr == DX:AX {
	//	*addr = CX:BX
	// } else {
	//	DX:AX = *addr
	// }
	// all in one instruction
	LOCK
	CMPXCHG8B 0(BP)

	JNZ addloop

	// success
	// return CX:BX
	MOVL BX, ret_lo+12(FP)
	MOVL CX, ret_hi+16(FP)
	RET

TEXT ·AddUintptr(SB), NOSPLIT, $0-12
	JMP ·Add32(SB)

TEXT ·AddInt32(SB), NOSPLIT, $0-12
	JMP ·Add32(SB)

TEXT ·AddInt64(SB), NOSPLIT, $0-20
	JMP ·Add64(SB)

//
// Compare and swap
//

// bool Cas32(int32 *val, int32 old, int32 new)
TEXT ·Cas32(SB), NOSPLIT, $0-13
	MOVL ptr+0(FP), BX
	MOVL old+4(FP), AX
	MOVL new+8(FP), CX
	LOCK
	CMPXCHGL CX, 0(BX)
	SETEQ ret+12(FP)
	RET

// bool ·Cas64(uint64 *val, uint64 old, uint64 new)
TEXT ·Cas64(SB), NOSPLIT, $0-21
	NO_LOCAL_POINTERS
	MOVL ptr+0(FP), BP
	TESTL $7, BP
	JZ 2(PC)
	CALL ·panicUnaligned(SB)
	MOVL old_lo+4(FP), AX
	MOVL old_hi+8(FP), DX
	MOVL new_lo+12(FP), BX
	MOVL new_hi+16(FP), CX
	LOCK
	CMPXCHG8B 0(BP)
	SETEQ ret+20(FP)
	RET

TEXT ·CasUintptr(SB), NOSPLIT, $0-13
	JMP ·Cas32(SB)

// bool CasUnsafePointer(void **p, void *old, void *new)
TEXT ·CasUnsafePointer(SB), NOSPLIT, $0-13
	MOVL ptr+0(FP), BX
	MOVL old+4(FP), AX
	MOVL new+8(FP), CX
	LOCK
	CMPXCHGL CX, 0(BX)
	SETEQ ret+12(FP)
	RET

TEXT ·CasInt32(SB), NOSPLIT, $0-13
	JMP ·Cas32(SB)

TEXT ·CasInt64(SB), NOSPLIT, $0-21
	JMP ·Cas64(SB)

//
// CasRel
//

TEXT ·CasRel32(SB), NOSPLIT, $0-13
	JMP ·Cas32(SB)

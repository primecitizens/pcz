// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT ·PublicationBarrier(SB),NOSPLIT|NOFRAME,$0-0
	// Stores are already ordered on s390x, so this is just a
	// compile barrier.
	RET

// func Store8(ptr *uint8, val uint8)
TEXT ·Store8(SB), NOSPLIT, $0
	MOVD ptr+0(FP), R2
	MOVB val+8(FP), R3
	MOVB R3, 0(R2)
	SYNC
	RET

// func Store32(ptr *uint32, val uint32)
TEXT ·Store32(SB), NOSPLIT, $0
	MOVD ptr+0(FP), R2
	MOVWZ val+8(FP), R3
	MOVW R3, 0(R2)
	SYNC
	RET

// func Store64(ptr *uint64, val uint64)
TEXT ·Store64(SB), NOSPLIT, $0
	MOVD ptr+0(FP), R2
	MOVD val+8(FP), R3
	MOVD R3, 0(R2)
	SYNC
	RET

// func StorePointer(ptr unsafe.Pointer, val unsafe.Pointer)
TEXT ·StorePointer(SB), NOSPLIT, $0
	MOVD ptr+0(FP), R2
	MOVD val+8(FP), R3
	MOVD R3, 0(R2)
	SYNC
	RET

// func StoreUintptr(ptr *uintptr, new uintptr)
TEXT ·StoreUintptr(SB), NOSPLIT, $0-16
	BR ·Store64(SB)

// func StoreInt32(ptr *int32, new int32)
TEXT ·StoreInt32(SB), NOSPLIT, $0-12
	BR ·Store32(SB)

// func StoreInt64(ptr *int64, new int64)
TEXT ·StoreInt64(SB), NOSPLIT, $0-16
	BR ·Store64(SB)

// 
// Load
//

// func LoadUintptr(ptr *uintptr) uintptr
TEXT ·LoadUintptr(SB), NOSPLIT, $0-16
	BR ·Load64(SB)

// func LoadUint(ptr *uint) uint
TEXT ·LoadUint(SB), NOSPLIT, $0-16
	BR ·Load64(SB)

// func LoadInt32(ptr *int32) int32
TEXT ·LoadInt32(SB), NOSPLIT, $0-12
	BR ·Load32(SB)

// func LoadInt64(ptr *int64) int64
TEXT ·LoadInt64(SB), NOSPLIT, $0-16
	BR ·Load64(SB)

//
// bitwise
//

// func Or8(addr *uint8, v uint8)
TEXT ·Or8(SB), NOSPLIT, $0-9
	MOVD ptr+0(FP), R3
	MOVBZ val+8(FP), R4
	// We don't have atomic operations that work on individual bytes so we
	// need to align addr down to a word boundary and create a mask
	// containing v to OR with the entire word atomically.
	MOVD $(3<<3), R5
	RXSBG $59, $60, $3, R3, R5 // R5 = 24 - ((addr % 4) * 8) = ((addr & 3) << 3) ^ (3 << 3)
	ANDW $~3, R3              // R3 = floor(addr, 4) = addr &^ 3
	SLW R5, R4               // R4 = uint32(v) << R5
	LAO R4, R6, 0(R3)        // R6 = *R3; *R3 |= R4; (atomic)
	RET

// func Or32(addr *uint32, v uint32)
TEXT ·Or32(SB), NOSPLIT, $0-12
	MOVD ptr+0(FP), R3
	MOVW val+8(FP), R4
	LAO R4, R6, 0(R3)        // R6 = *R3; *R3 |= R4; (atomic)
	RET

// func And8(addr *uint8, v uint8)
TEXT ·And8(SB), NOSPLIT, $0-9
	MOVD ptr+0(FP), R3
	MOVBZ val+8(FP), R4
	// We don't have atomic operations that work on individual bytes so we
	// need to align addr down to a word boundary and create a mask
	// containing v to AND with the entire word atomically.
	ORW $~0xff, R4           // R4 = uint32(v) | 0xffffff00
	MOVD $(3<<3), R5
	RXSBG $59, $60, $3, R3, R5 // R5 = 24 - ((addr % 4) * 8) = ((addr & 3) << 3) ^ (3 << 3)
	ANDW $~3, R3              // R3 = floor(addr, 4) = addr &^ 3
	RLL R5, R4, R4           // R4 = rotl(R4, R5)
	LAN R4, R6, 0(R3)        // R6 = *R3; *R3 &= R4; (atomic)
	RET

// func And32(addr *uint32, v uint32)
TEXT ·And32(SB), NOSPLIT, $0-12
	MOVD ptr+0(FP), R3
	MOVW val+8(FP), R4
	LAN R4, R6, 0(R3)        // R6 = *R3; *R3 &= R4; (atomic)
	RET


// 
// Swap
//

// func Swap32(ptr *uint32, new uint32) uint32
TEXT ·Swap32(SB), NOSPLIT, $0-20
	MOVD ptr+0(FP), R4
	MOVW new+8(FP), R3
	MOVW (R4), R6
repeat:
	CS R6, R3, (R4) // if R6==(R4) then (R4)=R3 else R6=(R4)
	BNE repeat
	MOVW R6, ret+16(FP)
	RET

// func Swap64(ptr *uint64, new uint64) uint64
TEXT ·Swap64(SB), NOSPLIT, $0-24
	MOVD ptr+0(FP), R4
	MOVD new+8(FP), R3
	MOVD (R4), R6
repeat:
	CSG R6, R3, (R4) // if R6==(R4) then (R4)=R3 else R6=(R4)
	BNE repeat
	MOVD R6, ret+16(FP)
	RET

// func SwapInt32(ptr *int32, new int32) int32
TEXT ·SwapInt32(SB), NOSPLIT, $0-20
	BR ·Swap32(SB)

// func SwapInt64(ptr *int64, new int64) int64
TEXT ·SwapInt64(SB), NOSPLIT, $0-24
	BR ·Swap64(SB)

// func SwapUintptr(ptr *uintptr, new uintptr) uintptr
TEXT ·SwapUintptr(SB), NOSPLIT, $0-24
	BR ·Swap64(SB)

// 
// Add
//

// func AddUintptr(ptr *uintptr, delta uintptr) uintptr
TEXT ·AddUintptr(SB), NOSPLIT, $0-24
	BR ·Add64(SB)

// func AddInt32(ptr *int32, delta int32) int32
TEXT ·AddInt32(SB), NOSPLIT, $0-20
	BR ·Add32(SB)

// func AddInt64(ptr *int64, delta int64) int64
TEXT ·AddInt64(SB), NOSPLIT, $0-24
	BR ·Add64(SB)

// func Add32(ptr *uint32, delta int32) uint32
// Atomically:
//	*ptr += delta
//	return *ptr
TEXT ·Add32(SB), NOSPLIT, $0-20
	MOVD ptr+0(FP), R4
	MOVW delta+8(FP), R5
	MOVW (R4), R3
repeat:
	ADD R5, R3, R6
	CS R3, R6, (R4) // if R3==(R4) then (R4)=R6 else R3=(R4)
	BNE repeat
	MOVW R6, ret+16(FP)
	RET

// func Add64(ptr *uint64, delta int64) uint64
TEXT ·Add64(SB), NOSPLIT, $0-24
	MOVD ptr+0(FP), R4
	MOVD delta+8(FP), R5
	MOVD (R4), R3
repeat:
	ADD R5, R3, R6
	CSG R3, R6, (R4) // if R3==(R4) then (R4)=R6 else R3=(R4)
	BNE repeat
	MOVD R6, ret+16(FP)
	RET

// 
// Compare and swap
// 

// func Cas32(ptr *uint32, old, new uint32) bool
TEXT ·Cas32(SB), NOSPLIT, $0-17
	MOVD ptr+0(FP), R3
	MOVWZ old+8(FP), R4
	MOVWZ new+12(FP), R5
	CS R4, R5, 0(R3)    //  if (R4 == 0(R3)) then 0(R3)= R5
	BNE cas_fail
	MOVB $1, ret+16(FP)
	RET
cas_fail:
	MOVB $0, ret+16(FP)
	RET

// func Cas64(ptr *uint64, old, new uint64) bool
TEXT ·Cas64(SB), NOSPLIT, $0-25
	MOVD ptr+0(FP), R3
	MOVD old+8(FP), R4
	MOVD new+16(FP), R5
	CSG R4, R5, 0(R3)    //  if (R4 == 0(R3)) then 0(R3)= R5
	BNE cas64_fail
	MOVB $1, ret+24(FP)
	RET
cas64_fail:
	MOVB $0, ret+24(FP)
	RET

// func CasUintptr(ptr *uintptr, old, new uintptr) bool
TEXT ·CasUintptr(SB), NOSPLIT, $0-25
	BR ·Cas64(SB)

// func CasUnsafePointer(ptr *unsafe.Pointer, old, new unsafe.Pointer) bool
TEXT ·CasUnsafePointer(SB), NOSPLIT, $0-25
	BR ·Cas64(SB)

// func CasInt32(ptr *int32, old, new int32) bool
TEXT ·CasInt32(SB), NOSPLIT, $0-17
	BR ·Cas32(SB)

// func CasInt64(ptr *int64, old, new int64) bool
TEXT ·CasInt64(SB), NOSPLIT, $0-25
	BR ·Cas64(SB)

// func CasRel32(ptr *uint32, old, new uint32) bool
TEXT ·CasRel32(SB), NOSPLIT, $0-17
	BR ·Cas32(SB)

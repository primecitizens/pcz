// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
// 
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build pcz && (ppc64 || ppc64le)

#include "textflag.h"

TEXT ·IndexSliceByte<ABIInternal>(SB),NOSPLIT|NOFRAME,$0-40
	// R3 = byte array pointer
	// R4 = length
	MOVD R6, R5 // R5 = byte
	MOVBZ ·isPOWER9(SB), R16
	BR indexbytebody<>(SB)

TEXT ·IndexByte<ABIInternal>(SB),NOSPLIT|NOFRAME,$0-32
	// R3 = string
	// R4 = length
	// R5 = byte
	MOVBZ ·isPOWER9(SB), R16
	BR indexbytebody<>(SB)

// R3 = addr of string
// R4 = len of string
// R5 = byte to find
// R16 = 1 if running on a POWER9 system, 0 otherwise
// On exit:
// R3 = return value
TEXT indexbytebody<>(SB),NOSPLIT|NOFRAME,$0-0
	MOVD R3,R17 // Save base address for calculating the index later.
	RLDICR $0,R3,$60,R8 // Align address to doubleword boundary in R8.
	RLDIMI $8,R5,$48,R5 // Replicating the byte across the register.
	ADD R4,R3,R7 // Last acceptable address in R7.

	RLDIMI $16,R5,$32,R5
	CMPU R4,$32 // Check if it's a small string (≤32 bytes). Those will be processed differently.
	MOVD $-1,R9
	RLWNM $3,R3,$26,$28,R6 // shift amount for mask (r3&0x7)*8
	RLDIMI $32,R5,$0,R5
	MOVD R7,R10 // Save last acceptable address in R10 for later.
	ADD $-1,R7,R7
#ifdef GOARCH_ppc64le
	SLD R6,R9,R9 // Prepare mask for Little Endian
#else
	SRD R6,R9,R9 // Same for Big Endian
#endif
	BLT small_string // Jump to the small string case if it's <32 bytes.
	CMP R16,$1 // optimize for power8 v power9
	BNE power8
	VSPLTISB $3,V10 // Use V10 as control for VBPERMQ
	MTVRD R5,V1
	LVSL (R0+R0),V11 // set up the permute vector such that V10 has {0x78, .., 0x8, 0x0}
	VSLB V11,V10,V10 // to extract the first bit of match result into GPR
	VSPLTB $7,V1,V1 // Replicate byte across V1
	CMP R4,$64
	MOVD $16,R11
	MOVD R3,R8
	BLT cmp32
	MOVD $32,R12
	MOVD $48,R6

loop64:
	LXVB16X (R0)(R8),V2 // scan 64 bytes at a time
	VCMPEQUBCC V2,V1,V6
	BNE CR6,foundat0 // match found at R8, jump out

	LXVB16X (R8)(R11),V2
	VCMPEQUBCC V2,V1,V6
	BNE CR6,foundat1 // match found at R8+16 bytes, jump out

	LXVB16X (R8)(R12),V2
	VCMPEQUBCC V2,V1,V6
	BNE CR6,foundat2 // match found at R8+32 bytes, jump out

	LXVB16X (R8)(R6),V2
	VCMPEQUBCC V2,V1,V6
	BNE CR6,foundat3 // match found at R8+48 bytes, jump out
	ADD $64,R8
	ADD $-64,R4
	CMP R4,$64 // >=64 bytes left to scan?
	BGE loop64
	CMP R4,$32
	BLT rem // jump to rem if there are < 32 bytes left
cmp32:
	LXVB16X (R0)(R8),V2 // 32-63 bytes left
	VCMPEQUBCC V2,V1,V6
	BNE CR6,foundat0 // match found at R8

	LXVB16X (R11)(R8),V2
	VCMPEQUBCC V2,V1,V6
	BNE CR6,foundat1 // match found at R8+16

	ADD $32,R8
	ADD $-32,R4
rem:
	RLDICR $0,R8,$60,R8 // align address to reuse code for tail end processing
	BR small_string

foundat3:
	ADD $16,R8
foundat2:
	ADD $16,R8
foundat1:
	ADD $16,R8
foundat0:
	// Compress the result into a single doubleword and
	// move it to a GPR for the final calculation.
	VBPERMQ V6,V10,V6
	MFVRD V6,R3
	// count leading zeroes upto the match that ends up in low 16 bits
	// in both endian modes, compute index by subtracting the number by 16
	CNTLZW R3,R11
	ADD $-16,R11
	ADD R8,R11,R3 // Calculate byte address
	SUB R17,R3
	RET
power8:
	// If we are 64-byte aligned, branch to qw_align just to get the auxiliary values
	// in V0, V1 and V10, then branch to the preloop.
	ANDCC $63,R3,R11
	BEQ CR0,qw_align
	RLDICL $0,R3,$61,R11

	MOVD 0(R8),R12 // Load one doubleword from the aligned address in R8.
	CMPB R12,R5,R3 // Check for a match.
	AND R9,R3,R3 // Mask bytes below s_base
	RLDICR $0,R7,$60,R7 // Last doubleword in R7
	CMPU R3,$0,CR7 // If we have a match, jump to the final computation
	BNE CR7,done
	ADD $8,R8,R8
	ADD $-8,R4,R4
	ADD R4,R11,R4

	// Check for quadword alignment
	ANDCC $15,R8,R11
	BEQ CR0,qw_align

	// Not aligned, so handle the next doubleword
	MOVD 0(R8),R12
	CMPB R12,R5,R3
	CMPU R3,$0,CR7
	BNE CR7,done
	ADD $8,R8,R8
	ADD $-8,R4,R4

	// Either quadword aligned or 64-byte at this point. We can use LVX.
qw_align:

	// Set up auxiliary data for the vectorized algorithm.
	VSPLTISB  $0,V0 // Replicate 0 across V0
	VSPLTISB  $3,V10 // Use V10 as control for VBPERMQ
	MTVRD   R5,V1
	LVSL   (R0+R0),V11
	VSLB   V11,V10,V10
	VSPLTB   $7,V1,V1 // Replicate byte across V1
	CMPU   R4, $64 // If len ≤ 64, don't use the vectorized loop
	BLE   tail

	// We will load 4 quardwords per iteration in the loop, so check for
	// 64-byte alignment. If 64-byte aligned, then branch to the preloop.
	ANDCC   $63,R8,R11
	BEQ   CR0,preloop

	// Not 64-byte aligned. Load one quadword at a time until aligned.
	LVX     (R8+R0),V4
	VCMPEQUBCC  V1,V4,V6 // Check for byte in V4
	BNE     CR6,found_qw_align
	ADD     $16,R8,R8
	ADD     $-16,R4,R4

	ANDCC     $63,R8,R11
	BEQ     CR0,preloop
	LVX     (R8+R0),V4
	VCMPEQUBCC  V1,V4,V6 // Check for byte in V4
	BNE     CR6,found_qw_align
	ADD     $16,R8,R8
	ADD     $-16,R4,R4

	ANDCC     $63,R8,R11
	BEQ     CR0,preloop
	LVX     (R8+R0),V4
	VCMPEQUBCC  V1,V4,V6 // Check for byte in V4
	BNE     CR6,found_qw_align
	ADD     $-16,R4,R4
	ADD     $16,R8,R8

	// 64-byte aligned. Prepare for the main loop.
preloop:
	CMPU R4,$64
	BLE tail       // If len ≤ 64, don't use the vectorized loop

	// We are now aligned to a 64-byte boundary. We will load 4 quadwords
	// per loop iteration. The last doubleword is in R10, so our loop counter
	// starts at (R10-R8)/64.
	SUB R8,R10,R6
	SRD $6,R6,R9      // Loop counter in R9
	MOVD R9,CTR

	ADD $-64,R8,R8   // Adjust index for loop entry
	MOVD $16,R11      // Load offsets for the vector loads
	MOVD $32,R9
	MOVD $48,R7

	// Main loop we will load 64 bytes per iteration
loop:
	ADD     $64,R8,R8       // Fuse addi+lvx for performance
	LVX     (R8+R0),V2       // Load 4 16-byte vectors
	LVX     (R8+R11),V3
	VCMPEQUB    V1,V2,V6       // Look for byte in each vector
	VCMPEQUB    V1,V3,V7

	LVX     (R8+R9),V4
	LVX     (R8+R7),V5
	VCMPEQUB    V1,V4,V8
	VCMPEQUB    V1,V5,V9

	VOR     V6,V7,V11       // Compress the result in a single vector
	VOR     V8,V9,V12
	VOR     V11,V12,V13
	VCMPEQUBCC  V0,V13,V14       // Check for byte
	BGE     CR6,found
	BC     16,0,loop       // bdnz loop

	// Handle the tailing bytes or R4 ≤ 64
	RLDICL $0,R6,$58,R4
	ADD $64,R8,R8
tail:
	CMPU     R4,$0
	BEQ     notfound
	LVX     (R8+R0),V4
	VCMPEQUBCC  V1,V4,V6
	BNE     CR6,found_qw_align
	ADD     $16,R8,R8
	CMPU     R4,$16,CR6
	BLE     CR6,notfound
	ADD     $-16,R4,R4

	LVX     (R8+R0),V4
	VCMPEQUBCC  V1,V4,V6
	BNE     CR6,found_qw_align
	ADD     $16,R8,R8
	CMPU     R4,$16,CR6
	BLE     CR6,notfound
	ADD     $-16,R4,R4

	LVX     (R8+R0),V4
	VCMPEQUBCC  V1,V4,V6
	BNE     CR6,found_qw_align
	ADD     $16,R8,R8
	CMPU     R4,$16,CR6
	BLE     CR6,notfound
	ADD     $-16,R4,R4

	LVX     (R8+R0),V4
	VCMPEQUBCC  V1,V4,V6
	BNE     CR6,found_qw_align

notfound:
	MOVD $-1, R3
	RET

found:
	// We will now compress the results into a single doubleword,
	// so it can be moved to a GPR for the final index calculation.

	// The bytes in V6-V9 are either 0x00 or 0xFF. So, permute the
	// first bit of each byte into bits 48-63.
	VBPERMQ   V6,V10,V6
	VBPERMQ   V7,V10,V7
	VBPERMQ   V8,V10,V8
	VBPERMQ   V9,V10,V9

	// Shift each 16-bit component into its correct position for
	// merging into a single doubleword.
#ifdef GOARCH_ppc64le
	VSLDOI   $2,V7,V7,V7
	VSLDOI   $4,V8,V8,V8
	VSLDOI   $6,V9,V9,V9
#else
	VSLDOI   $6,V6,V6,V6
	VSLDOI   $4,V7,V7,V7
	VSLDOI   $2,V8,V8,V8
#endif

	// Merge V6-V9 into a single doubleword and move to a GPR.
	VOR V6,V7,V11
	VOR V8,V9,V4
	VOR V4,V11,V4
	MFVRD V4,R3

#ifdef GOARCH_ppc64le
	ADD   $-1,R3,R11
	ANDN   R3,R11,R11
	POPCNTD   R11,R11 // Count trailing zeros (Little Endian).
#else
	CNTLZD R3,R11 // Count leading zeros (Big Endian).
#endif
	ADD R8,R11,R3 // Calculate byte address

return:
	SUB R17, R3
	RET

found_qw_align:
	// Use the same algorithm as above. Compress the result into
	// a single doubleword and move it to a GPR for the final
	// calculation.
	VBPERMQ   V6,V10,V6

#ifdef GOARCH_ppc64le
	MFVRD   V6,R3
	ADD   $-1,R3,R11
	ANDN   R3,R11,R11
	POPCNTD   R11,R11
#else
	VSLDOI   $6,V6,V6,V6
	MFVRD   V6,R3
	CNTLZD   R3,R11
#endif
	ADD   R8,R11,R3
	CMPU   R11,R4
	BLT   return
	BR   notfound
	PCALIGN   $16

done:
	ADD $-1,R10,R6
	// Offset of last index for the final
	// doubleword comparison
	RLDICL $0,R6,$61,R6
	// At this point, R3 has 0xFF in the same position as the byte we are
	// looking for in the doubleword. Use that to calculate the exact index
	// of the byte.
#ifdef GOARCH_ppc64le
	ADD $-1,R3,R11
	ANDN R3,R11,R11
	POPCNTD R11,R11 // Count trailing zeros (Little Endian).
#else
	CNTLZD R3,R11 // Count leading zeros (Big Endian).
#endif
	CMPU R8,R7 // Check if we are at the last doubleword.
	SRD $3,R11 // Convert trailing zeros to bytes.
	ADD R11,R8,R3
	CMPU R11,R6,CR7 // If at the last doubleword, check the byte offset.
	BNE return
	BLE CR7,return
	BR notfound

small_string:
	// process string of length < 32 bytes
	// We unroll this loop for better performance.
	CMPU R4,$0 // Check for length=0
	BEQ notfound

	MOVD 0(R8),R12 // Load one doubleword from the aligned address in R8.
	CMPB R12,R5,R3 // Check for a match.
	AND R9,R3,R3 // Mask bytes below s_base.
	CMPU R3,$0,CR7 // If we have a match, jump to the final computation.
	RLDICR $0,R7,$60,R7 // Last doubleword in R7.
	CMPU R8,R7
	BNE CR7,done
	BEQ notfound // Hit length.

	MOVDU 8(R8),R12
	CMPB R12,R5,R3
	CMPU R3,$0,CR6
	CMPU R8,R7
	BNE CR6,done
	BEQ notfound

	MOVDU 8(R8),R12
	CMPB R12,R5,R3
	CMPU R3,$0,CR6
	CMPU R8,R7
	BNE CR6,done
	BEQ notfound

	MOVDU 8(R8),R12
	CMPB R12,R5,R3
	CMPU R3,$0,CR6
	CMPU R8,R7
	BNE CR6,done
	BEQ notfound

	MOVDU 8(R8),R12
	CMPB R12,R5,R3
	CMPU R3,$0,CR6
	BNE CR6,done
	BR notfound


// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pe

// Reloc represents a PE COFF relocation.
// Each section contains its own relocation list.
type Reloc struct {
	VirtualAddress   uint32
	SymbolTableIndex uint32
	Type             uint16
}

const (
	IMAGE_REL_I386_DIR32  = 0x0006
	IMAGE_REL_I386_SECREL = 0x000B
	IMAGE_REL_I386_REL32  = 0x0014

	IMAGE_REL_AMD64_ADDR64 = 0x0001
	IMAGE_REL_AMD64_ADDR32 = 0x0002
	IMAGE_REL_AMD64_REL32  = 0x0004
	IMAGE_REL_AMD64_SECREL = 0x000B

	IMAGE_REL_ARM_ABSOLUTE = 0x0000
	IMAGE_REL_ARM_ADDR32   = 0x0001
	IMAGE_REL_ARM_ADDR32NB = 0x0002
	IMAGE_REL_ARM_BRANCH24 = 0x0003
	IMAGE_REL_ARM_BRANCH11 = 0x0004
	IMAGE_REL_ARM_SECREL   = 0x000F

	IMAGE_REL_ARM64_ABSOLUTE       = 0x0000
	IMAGE_REL_ARM64_ADDR32         = 0x0001
	IMAGE_REL_ARM64_ADDR32NB       = 0x0002
	IMAGE_REL_ARM64_BRANCH26       = 0x0003
	IMAGE_REL_ARM64_PAGEBASE_REL21 = 0x0004
	IMAGE_REL_ARM64_REL21          = 0x0005
	IMAGE_REL_ARM64_PAGEOFFSET_12A = 0x0006
	IMAGE_REL_ARM64_PAGEOFFSET_12L = 0x0007
	IMAGE_REL_ARM64_SECREL         = 0x0008
	IMAGE_REL_ARM64_SECREL_LOW12A  = 0x0009
	IMAGE_REL_ARM64_SECREL_HIGH12A = 0x000A
	IMAGE_REL_ARM64_SECREL_LOW12L  = 0x000B
	IMAGE_REL_ARM64_TOKEN          = 0x000C
	IMAGE_REL_ARM64_SECTION        = 0x000D
	IMAGE_REL_ARM64_ADDR64         = 0x000E
	IMAGE_REL_ARM64_BRANCH19       = 0x000F
	IMAGE_REL_ARM64_BRANCH14       = 0x0010
	IMAGE_REL_ARM64_REL32          = 0x0011

	IMAGE_REL_BASED_HIGHLOW = 3
	IMAGE_REL_BASED_DIR64   = 10
)

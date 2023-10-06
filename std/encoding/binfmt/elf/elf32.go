// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elf

// ELF32 File header.
type Header32 struct {
	Ident     [EI_NIDENT]byte /* File identification. */
	Type      uint16          /* File type. */
	Machine   uint16          /* Machine architecture. */
	Version   uint32          /* ELF format version. */
	Entry     uint32          /* Entry point. */
	Phoff     uint32          /* Program header file offset. */
	Shoff     uint32          /* Section header file offset. */
	Flags     uint32          /* Architecture-specific flags. */
	Ehsize    uint16          /* Size of ELF header in bytes. */
	Phentsize uint16          /* Size of program header entry. */
	Phnum     uint16          /* Number of program header entries. */
	Shentsize uint16          /* Size of section header entry. */
	Shnum     uint16          /* Number of section header entries. */
	Shstrndx  uint16          /* Section name strings section. */
}

// ELF32 Section header.
type Section32 struct {
	Name      uint32 /* Section name (index into the section header string table). */
	Type      uint32 /* Section type. */
	Flags     uint32 /* Section flags. */
	Addr      uint32 /* Address in memory image. */
	Off       uint32 /* Offset in file. */
	Size      uint32 /* Size in bytes. */
	Link      uint32 /* Index of a related section. */
	Info      uint32 /* Depends on section type. */
	Addralign uint32 /* Alignment in bytes. */
	Entsize   uint32 /* Size of each entry in section. */
}

// ELF32 Program header.
type Prog32 struct {
	Type   uint32 /* Entry type. */
	Off    uint32 /* File offset of contents. */
	Vaddr  uint32 /* Virtual address in memory image. */
	Paddr  uint32 /* Physical address (not used). */
	Filesz uint32 /* Size of contents in file. */
	Memsz  uint32 /* Size of contents in memory. */
	Flags  uint32 /* Access permission flags. */
	Align  uint32 /* Alignment in memory and file. */
}

// ELF32 Dynamic structure. The ".dynamic" section contains an array of them.
type Dyn32 struct {
	Tag int32  /* Entry type. */
	Val uint32 /* Integer/Address value. */
}

// ELF32 Compression header.
type Chdr32 struct {
	Type      uint32
	Size      uint32
	Addralign uint32
}

/*
 * Relocation entries.
 */

// ELF32 Relocations that don't need an addend field.
type Rel32 struct {
	Off  uint32 /* Location to be relocated. */
	Info uint32 /* Relocation type and symbol index. */
}

// ELF32 Relocations that need an addend field.
type Rela32 struct {
	Off    uint32 /* Location to be relocated. */
	Info   uint32 /* Relocation type and symbol index. */
	Addend int32  /* Addend. */
}

func R_SYM32(info uint32) uint32      { return info >> 8 }
func R_TYPE32(info uint32) uint32     { return info & 0xff }
func R_INFO32(sym, typ uint32) uint32 { return sym<<8 | typ }

// ELF32 Symbol.
type Sym32 struct {
	Name  uint32
	Value uint32
	Size  uint32
	Info  uint8
	Other uint8
	Shndx uint16
}

const Sym32Size = 16

func ST_BIND(info uint8) SymBind { return SymBind(info >> 4) }
func ST_TYPE(info uint8) SymType { return SymType(info & 0xF) }
func ST_INFO(bind SymBind, typ SymType) uint8 {
	return uint8(bind)<<4 | uint8(typ)&0xf
}
func ST_VISIBILITY(other uint8) SymVis { return SymVis(other & 3) }

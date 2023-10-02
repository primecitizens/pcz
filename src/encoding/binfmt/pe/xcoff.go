// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pe

const (
	// Total amount of space to reserve at the start of the file
	// for File Header, Auxiliary Header, and Section Headers.
	// May waste some.
	XCOFFHDRRESERVE = FILHSZ_64 + AOUTHSZ_EXEC64 + SCNHSZ_64*23

	// base on dump -o, then rounded from 32B to 64B to
	// match worst case elf text section alignment on ppc64.
	XCOFFSECTALIGN int64 = 64

	// XCOFF binaries should normally have all its sections position-independent.
	// However, this is not yet possible for .text because of some R_ADDR relocations
	// inside RODATA symbols.
	// .data and .bss are position-independent so their address start inside a unreachable
	// segment during execution to force segfault if something is wrong.
	XCOFFTEXTBASE = 0x100000000 // Start of text address
	XCOFFDATABASE = 0x200000000 // Start of data address
)

// File Header
type XcoffFileHdr64 struct {
	Fmagic   uint16 // Target machine
	Fnscns   uint16 // Number of sections
	Ftimedat int32  // Time and date of file creation
	Fsymptr  uint64 // Byte offset to symbol table start
	Fopthdr  uint16 // Number of bytes in optional header
	Fflags   uint16 // Flags
	Fnsyms   int32  // Number of entries in symbol table
}

const (
	U64_TOCMAGIC = 0767 // AIX 64-bit XCOFF
)

// Flags that describe the type of the object file.
const (
	F_RELFLG    = 0x0001
	F_EXEC      = 0x0002
	F_LNNO      = 0x0004
	F_FDPR_PROF = 0x0010
	F_FDPR_OPTI = 0x0020
	F_DSA       = 0x0040
	F_VARPG     = 0x0100
	F_DYNLOAD   = 0x1000
	F_SHROBJ    = 0x2000
	F_LOADONLY  = 0x4000
)

// Auxiliary Header
type XcoffAoutHdr64 struct {
	Omagic      int16    // Flags - Ignored If Vstamp Is 1
	Ovstamp     int16    // Version
	Odebugger   uint32   // Reserved For Debugger
	Otextstart  uint64   // Virtual Address Of Text
	Odatastart  uint64   // Virtual Address Of Data
	Otoc        uint64   // Toc Address
	Osnentry    int16    // Section Number For Entry Point
	Osntext     int16    // Section Number For Text
	Osndata     int16    // Section Number For Data
	Osntoc      int16    // Section Number For Toc
	Osnloader   int16    // Section Number For Loader
	Osnbss      int16    // Section Number For Bss
	Oalgntext   int16    // Max Text Alignment
	Oalgndata   int16    // Max Data Alignment
	Omodtype    [2]byte  // Module Type Field
	Ocpuflag    uint8    // Bit Flags - Cputypes Of Objects
	Ocputype    uint8    // Reserved for CPU type
	Otextpsize  uint8    // Requested text page size
	Odatapsize  uint8    // Requested data page size
	Ostackpsize uint8    // Requested stack page size
	Oflags      uint8    // Flags And TLS Alignment
	Otsize      uint64   // Text Size In Bytes
	Odsize      uint64   // Data Size In Bytes
	Obsize      uint64   // Bss Size In Bytes
	Oentry      uint64   // Entry Point Address
	Omaxstack   uint64   // Max Stack Size Allowed
	Omaxdata    uint64   // Max Data Size Allowed
	Osntdata    int16    // Section Number For Tdata Section
	Osntbss     int16    // Section Number For Tbss Section
	Ox64flags   uint16   // Additional Flags For 64-Bit Objects
	Oresv3a     int16    // Reserved
	Oresv3      [2]int32 // Reserved
}

// Section Header
type XcoffScnHdr64 struct {
	Sname    [8]byte // Section Name
	Spaddr   uint64  // Physical Address
	Svaddr   uint64  // Virtual Address
	Ssize    uint64  // Section Size
	Sscnptr  uint64  // File Offset To Raw Data
	Srelptr  uint64  // File Offset To Relocation
	Slnnoptr uint64  // File Offset To Line Numbers
	Snreloc  uint32  // Number Of Relocation Entries
	Snlnno   uint32  // Number Of Line Number Entries
	Sflags   uint32  // flags
}

// Flags defining the section type.
const (
	STYP_DWARF  = 0x0010
	STYP_TEXT   = 0x0020
	STYP_DATA   = 0x0040
	STYP_BSS    = 0x0080
	STYP_EXCEPT = 0x0100
	STYP_INFO   = 0x0200
	STYP_TDATA  = 0x0400
	STYP_TBSS   = 0x0800
	STYP_LOADER = 0x1000
	STYP_DEBUG  = 0x2000
	STYP_TYPCHK = 0x4000
	STYP_OVRFLO = 0x8000
)
const (
	SSUBTYP_DWINFO  = 0x10000 // DWARF info section
	SSUBTYP_DWLINE  = 0x20000 // DWARF line-number section
	SSUBTYP_DWPBNMS = 0x30000 // DWARF public names section
	SSUBTYP_DWPBTYP = 0x40000 // DWARF public types section
	SSUBTYP_DWARNGE = 0x50000 // DWARF aranges section
	SSUBTYP_DWABREV = 0x60000 // DWARF abbreviation section
	SSUBTYP_DWSTR   = 0x70000 // DWARF strings section
	SSUBTYP_DWRNGES = 0x80000 // DWARF ranges section
	SSUBTYP_DWLOC   = 0x90000 // DWARF location lists section
	SSUBTYP_DWFRAME = 0xA0000 // DWARF frames section
	SSUBTYP_DWMAC   = 0xB0000 // DWARF macros section
)

// Headers size
const (
	FILHSZ_32      = 20
	FILHSZ_64      = 24
	AOUTHSZ_EXEC32 = 72
	AOUTHSZ_EXEC64 = 120
	SCNHSZ_32      = 40
	SCNHSZ_64      = 72
	LDHDRSZ_32     = 32
	LDHDRSZ_64     = 56
	LDSYMSZ_64     = 24
	RELSZ_64       = 14
)

// Type representing all XCOFF symbols.
type xcoffSym interface {
}

// Symbol Table Entry
type XcoffSymEnt64 struct {
	Nvalue  uint64 // Symbol value
	Noffset uint32 // Offset of the name in string table or .debug section
	Nscnum  int16  // Section number of symbol
	Ntype   uint16 // Basic and derived type specification
	Nsclass uint8  // Storage class of symbol
	Nnumaux int8   // Number of auxiliary entries
}

const SYMESZ = 18

const (
	// Nscnum
	N_DEBUG = -2
	N_ABS   = -1
	N_UNDEF = 0

	//Ntype
	SYM_V_INTERNAL  = 0x1000
	SYM_V_HIDDEN    = 0x2000
	SYM_V_PROTECTED = 0x3000
	SYM_V_EXPORTED  = 0x4000
	SYM_TYPE_FUNC   = 0x0020 // is function
)

// Storage Class.
const (
	C_NULL    = 0   // Symbol table entry marked for deletion
	C_EXT     = 2   // External symbol
	C_STAT    = 3   // Static symbol
	C_BLOCK   = 100 // Beginning or end of inner block
	C_FCN     = 101 // Beginning or end of function
	C_FILE    = 103 // Source file name and compiler information
	C_HIDEXT  = 107 // Unnamed external symbol
	C_BINCL   = 108 // Beginning of include file
	C_EINCL   = 109 // End of include file
	C_WEAKEXT = 111 // Weak external symbol
	C_DWARF   = 112 // DWARF symbol
	C_GSYM    = 128 // Global variable
	C_LSYM    = 129 // Automatic variable allocated on stack
	C_PSYM    = 130 // Argument to subroutine allocated on stack
	C_RSYM    = 131 // Register variable
	C_RPSYM   = 132 // Argument to function or procedure stored in register
	C_STSYM   = 133 // Statically allocated symbol
	C_BCOMM   = 135 // Beginning of common block
	C_ECOML   = 136 // Local member of common block
	C_ECOMM   = 137 // End of common block
	C_DECL    = 140 // Declaration of object
	C_ENTRY   = 141 // Alternate entry
	C_FUN     = 142 // Function or procedure
	C_BSTAT   = 143 // Beginning of static block
	C_ESTAT   = 144 // End of static block
	C_GTLS    = 145 // Global thread-local variable
	C_STTLS   = 146 // Static thread-local variable
)

// File Auxiliary Entry
type XcoffAuxFile64 struct {
	Xzeroes  uint32 // The name is always in the string table
	Xoffset  uint32 // Offset in the string table
	X_pad1   [6]byte
	Xftype   uint8 // Source file string type
	X_pad2   [2]byte
	Xauxtype uint8 // Type of auxiliary entry
}

// Function Auxiliary Entry
type XcoffAuxFcn64 struct {
	Xlnnoptr uint64 // File pointer to line number
	Xfsize   uint32 // Size of function in bytes
	Xendndx  uint32 // Symbol table index of next entry
	Xpad     uint8  // Unused
	Xauxtype uint8  // Type of auxiliary entry
}

// csect Auxiliary Entry.
type XcoffAuxCSect64 struct {
	Xscnlenlo uint32 // Lower 4 bytes of length or symbol table index
	Xparmhash uint32 // Offset of parameter type-check string
	Xsnhash   uint16 // .typchk section number
	Xsmtyp    uint8  // Symbol alignment and type
	Xsmclas   uint8  // Storage-mapping class
	Xscnlenhi uint32 // Upper 4 bytes of length or symbol table index
	Xpad      uint8  // Unused
	Xauxtype  uint8  // Type of auxiliary entry
}

// DWARF Auxiliary Entry
type XcoffAuxDWARF64 struct {
	Xscnlen  uint64 // Length of this symbol section
	X_pad    [9]byte
	Xauxtype uint8 // Type of auxiliary entry
}

// Auxiliary type
const (
	_AUX_EXCEPT = 255
	_AUX_FCN    = 254
	_AUX_SYM    = 253
	_AUX_FILE   = 252
	_AUX_CSECT  = 251
	_AUX_SECT   = 250
)

// Xftype field
const (
	XFT_FN = 0   // Source File Name
	XFT_CT = 1   // Compile Time Stamp
	XFT_CV = 2   // Compiler Version Number
	XFT_CD = 128 // Compiler Defined Information/

)

// Symbol type field.
const (
	XTY_ER  = 0    // External reference
	XTY_SD  = 1    // Section definition
	XTY_LD  = 2    // Label definition
	XTY_CM  = 3    // Common csect definition
	XTY_WK  = 0x8  // Weak symbol
	XTY_EXP = 0x10 // Exported symbol
	XTY_ENT = 0x20 // Entry point symbol
	XTY_IMP = 0x40 // Imported symbol
)

// Storage-mapping class.
const (
	XMC_PR     = 0  // Program code
	XMC_RO     = 1  // Read-only constant
	XMC_DB     = 2  // Debug dictionary table
	XMC_TC     = 3  // TOC entry
	XMC_UA     = 4  // Unclassified
	XMC_RW     = 5  // Read/Write data
	XMC_GL     = 6  // Global linkage
	XMC_XO     = 7  // Extended operation
	XMC_SV     = 8  // 32-bit supervisor call descriptor
	XMC_BS     = 9  // BSS class
	XMC_DS     = 10 // Function descriptor
	XMC_UC     = 11 // Unnamed FORTRAN common
	XMC_TC0    = 15 // TOC anchor
	XMC_TD     = 16 // Scalar data entry in the TOC
	XMC_SV64   = 17 // 64-bit supervisor call descriptor
	XMC_SV3264 = 18 // Supervisor call descriptor for both 32-bit and 64-bit
	XMC_TL     = 20 // Read/Write thread-local data
	XMC_UL     = 21 // Read/Write thread-local data (.tbss)
	XMC_TE     = 22 // TOC entry
)

// Loader Header
type XcoffLdHdr64 struct {
	Lversion int32  // Loader section version number
	Lnsyms   int32  // Number of symbol table entries
	Lnreloc  int32  // Number of relocation table entries
	Listlen  uint32 // Length of import file ID string table
	Lnimpid  int32  // Number of import file IDs
	Lstlen   uint32 // Length of string table
	Limpoff  uint64 // Offset to start of import file IDs
	Lstoff   uint64 // Offset to start of string table
	Lsymoff  uint64 // Offset to start of symbol table
	Lrldoff  uint64 // Offset to start of relocation entries
}

// Loader Symbol
type XcoffLdSym64 struct {
	Lvalue  uint64 // Address field
	Loffset uint32 // Byte offset into string table of symbol name
	Lscnum  int16  // Section number containing symbol
	Lsmtype int8   // Symbol type, export, import flags
	Lsmclas int8   // Symbol storage class
	Lifile  int32  // Import file ID; ordinal of import file IDs
	Lparm   uint32 // Parameter type-check field
}

type XcoffLdRel64 struct {
	Lvaddr  uint64 // Address Field
	Lrtype  uint16 // Relocation Size and Type
	Lrsecnm int16  // Section Number being relocated
	Lsymndx int32  // Loader-Section symbol table index
}

const (
	XCOFF_R_POS = 0x00 // A(sym) Positive Relocation
	XCOFF_R_NEG = 0x01 // -A(sym) Negative Relocation
	XCOFF_R_REL = 0x02 // A(sym-*) Relative to self
	XCOFF_R_TOC = 0x03 // A(sym-TOC) Relative to TOC
	XCOFF_R_TRL = 0x12 // A(sym-TOC) TOC Relative indirect load.

	XCOFF_R_TRLA = 0x13 // A(sym-TOC) TOC Rel load address. modifiable inst
	XCOFF_R_GL   = 0x05 // A(external TOC of sym) Global Linkage
	XCOFF_R_TCL  = 0x06 // A(local TOC of sym) Local object TOC address
	XCOFF_R_RL   = 0x0C // A(sym) Pos indirect load. modifiable instruction
	XCOFF_R_RLA  = 0x0D // A(sym) Pos Load Address. modifiable instruction
	XCOFF_R_REF  = 0x0F // AL0(sym) Non relocating ref. No garbage collect
	XCOFF_R_BA   = 0x08 // A(sym) Branch absolute. Cannot modify instruction
	XCOFF_R_RBA  = 0x18 // A(sym) Branch absolute. modifiable instruction
	XCOFF_R_BR   = 0x0A // A(sym-*) Branch rel to self. non modifiable
	XCOFF_R_RBR  = 0x1A // A(sym-*) Branch rel to self. modifiable instr

	XCOFF_R_TLS    = 0x20 // General-dynamic reference to TLS symbol
	XCOFF_R_TLS_IE = 0x21 // Initial-exec reference to TLS symbol
	XCOFF_R_TLS_LD = 0x22 // Local-dynamic reference to TLS symbol
	XCOFF_R_TLS_LE = 0x23 // Local-exec reference to TLS symbol
	XCOFF_R_TLSM   = 0x24 // Module reference to TLS symbol
	XCOFF_R_TLSML  = 0x25 // Module reference to local (own) module

	XCOFF_R_TOCU = 0x30 // Relative to TOC - high order bits
	XCOFF_R_TOCL = 0x31 // Relative to TOC - low order bits
)

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pe

const COFFSymbolSize = 18

// COFFSymbol represents single COFF symbol table record.
type COFFSymbol struct {
	Name               [8]uint8
	Value              uint32
	SectionNumber      int16
	Type               uint16
	StorageClass       uint8
	NumberOfAuxSymbols uint8
}

// COFFSymbolAuxFormat5 describes the expected form of an aux symbol
// attached to a section definition symbol. The PE format defines a
// number of different aux symbol formats: format 1 for function
// definitions, format 2 for .be and .ef symbols, and so on. Format 5
// holds extra info associated with a section definition, including
// number of relocations + line numbers, as well as COMDAT info. See
// https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#auxiliary-format-5-section-definitions
// for more on what's going on here.
type COFFSymbolAuxFormat5 struct {
	Size           uint32
	NumRelocs      uint16
	NumLineNumbers uint16
	Checksum       uint32
	SecNum         uint16
	Selection      uint8
	_              [3]uint8 // padding
}

// These constants make up the possible values for the 'Selection'
// field in an AuxFormat5.
const (
	IMAGE_COMDAT_SELECT_NODUPLICATES = 1
	IMAGE_COMDAT_SELECT_ANY          = 2
	IMAGE_COMDAT_SELECT_SAME_SIZE    = 3
	IMAGE_COMDAT_SELECT_EXACT_MATCH  = 4
	IMAGE_COMDAT_SELECT_ASSOCIATIVE  = 5
	IMAGE_COMDAT_SELECT_LARGEST      = 6
)

// See https://docs.microsoft.com/en-us/windows/win32/debug/pe-format.
const (
	// TODO: the Microsoft doco says IMAGE_SYM_DTYPE_ARRAY is 3 and IMAGE_SYM_DTYPE_FUNCTION is 2
	IMAGE_SYM_TYPE_NULL      = 0
	IMAGE_SYM_TYPE_STRUCT    = 8
	IMAGE_SYM_DTYPE_FUNCTION = 0x20
	IMAGE_SYM_DTYPE_ARRAY    = 0x30
	IMAGE_SYM_CLASS_EXTERNAL = 2
	IMAGE_SYM_CLASS_STATIC   = 3
)

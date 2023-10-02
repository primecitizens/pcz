// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package wasm

import (
	"unsafe"

	"github.com/primecitizens/std/encoding/binary"
)

const (
	// Magic is the wasm file magic number.
	Magic = 0x6d736100 // \0asm
)

type SectionID byte

const (
	Section_CUSTOM SectionID = iota
	Section_TYPE
	Section_IMPORT
	Section_FUNCTION
	Section_TABLE
	Section_MEMORY
	Section_GLOBAL
	Section_EXPORT
	Section_START
	Section_ELEMENT
	Section_CODE
	Section_DATA
	Section_DATA_COUNT
)

type Header struct {
	Magic   binary.U32le[uint32]
	Version binary.U32le[uint32]
}

func (h *Header) Decode(data []byte) (next []byte, ok bool) {
	if len(data) < 8 {
		return
	}

	h.Magic = *(*binary.U32le[uint32])(unsafe.Pointer(&data[0]))
	h.Version = *(*binary.U32le[uint32])(unsafe.Pointer(&data[4]))
	return data[8:], true
}

type Section struct {
	ID   SectionID
	Data []byte
}

func (sect *Section) Decode(data []byte) (next []byte, ok bool) {
	if len(data) <= 4 {
		return
	}

	sect.ID, data, ok = readU8[SectionID](data)
	if !ok {
		return
	}

	sz, data, ok := readU32(data)
	if ok = ok && int(sz) <= len(data); !ok {
		return
	}

	sect.Data = data[:sz]
	return data[sz:], true
}

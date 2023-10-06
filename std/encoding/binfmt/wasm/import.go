// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package wasm

import (
	"unsafe"
)

type ImportSection struct {
	// Entries is the count of import entries to decode.
	Entries uint32
}

func (sect *ImportSection) Decode(data []byte) (next []byte, ok bool) {
	sect.Entries, next, ok = readU32(data)
	return
}

type ImportKind uint8

const (
	ImportKindFunction ImportKind = iota
	ImportKindTable
	ImportKindMemory
	ImportKindGlobal
)

type ImportEntry struct {
	Module string
	Field  string

	Kind ImportKind

	Function FunctionImport
	Table    TableImport
	Memory   MemoryImport
	Global   GlobalImport
}

func (e *ImportEntry) Parse(data []byte) (next []byte, ok bool) {
	sz, data, ok := readU32(data)
	if ok = ok && int(sz) < len(data); !ok {
		return
	}

	e.Module = unsafe.String(unsafe.SliceData(data), sz)

	sz, data, ok = readU32(data[sz:])
	if ok = ok && int(sz) < len(data); !ok {
		return
	}

	e.Field = unsafe.String(unsafe.SliceData(data), sz)

	if e.Kind, data, ok = readU8[ImportKind](data[sz:]); !ok {
		return
	}

	switch e.Kind {
	case ImportKindFunction:
		e.Function.Index, next, ok = readU32(data)
	case ImportKindTable:
		if e.Table.ElemType, data, ok = readU8[int8](data); !ok {
			return
		}

		next, ok = e.Table.Limits.Parse(data)
	case ImportKindMemory:
		next, ok = e.Memory.Limits.Parse(data)
	case ImportKindGlobal:
		if e.Global.ContentType, data, ok = readU8[int8](data); !ok {
			return
		}
		var mutable uint8
		mutable, next, ok = readU8[uint8](data)
		e.Global.Mutable = mutable == 1
	default:
		return
	}

	return
}

type FunctionImport struct {
	Index uint32
}

type MemoryImport struct {
	Limits ResizableLimits
}

type TableImport struct {
	ElemType int8
	Limits   ResizableLimits
}

type GlobalImport struct {
	ContentType int8
	Mutable     bool
}

type ResizableLimits struct {
	Initial uint32
	Maximum uint32
}

func (l *ResizableLimits) Parse(data []byte) (next []byte, ok bool) {
	var hasMax uint8
	if hasMax, data, ok = readU8[uint8](data); !ok {
		return
	}

	if l.Initial, data, ok = readU32(data); !ok {
		return
	}

	if hasMax == 0 {
		return data, true
	}

	l.Maximum, next, ok = readU32(data)
	return
}

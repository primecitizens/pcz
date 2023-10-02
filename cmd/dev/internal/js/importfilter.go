// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package js

import (
	"fmt"

	"github.com/primecitizens/std/encoding/binfmt/wasm"
)

type WasmImportKey struct {
	Module string
	Field  string
}

type WasmImportFilter interface {
	Keep(module, field string) bool
}

type SimpleWasmImportFilter map[WasmImportKey]bool

func (m SimpleWasmImportFilter) Keep(module, field string) bool {
	return m[WasmImportKey{
		Module: module,
		Field:  field,
	}]
}

func CreateSimpleFilterFromWasm(blob []byte) (filter SimpleWasmImportFilter, err error) {
	var (
		header wasm.Header
		sect   wasm.Section
		imp    wasm.ImportSection
		ent    wasm.ImportEntry
	)

	blob, ok := header.Decode(blob)
	if !ok {
		return nil, fmt.Errorf("failed to get header of wasm blob")
	}

	filter = make(SimpleWasmImportFilter)
	for ok {
		blob, ok = sect.Decode(blob)
		if sect.ID != wasm.Section_IMPORT {
			continue
		}

		sect.Data, ok = imp.Decode(sect.Data)
		for i := 0; ok && i < int(imp.Entries); i++ {
			sect.Data, ok = ent.Parse(sect.Data)
			if !ok {
				break
			}

			if ent.Kind != wasm.ImportKindFunction {
				continue
			}

			filter[WasmImportKey{
				Module: ent.Module,
				Field:  ent.Field,
			}] = true
		}
	}

	return
}

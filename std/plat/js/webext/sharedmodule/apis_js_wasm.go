// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package sharedmodule

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/sharedmodule/bindings"
)

type Export struct {
	// Allowlist is "Export.allowlist"
	//
	// Optional
	Allowlist js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Export with all fields set.
func (p Export) FromRef(ref js.Ref) Export {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Export in the application heap.
func (p Export) New() js.Ref {
	return bindings.ExportJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Export) UpdateFrom(ref js.Ref) {
	bindings.ExportJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Export) Update(ref js.Ref) {
	bindings.ExportJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Export) FreeMembers(recursive bool) {
	js.Free(
		p.Allowlist.Ref(),
	)
	p.Allowlist = p.Allowlist.FromRef(js.Undefined)
}

type Import struct {
	// Id is "Import.id"
	//
	// Optional
	Id js.String
	// MinimumVersion is "Import.minimum_version"
	//
	// Optional
	MinimumVersion js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Import with all fields set.
func (p Import) FromRef(ref js.Ref) Import {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Import in the application heap.
func (p Import) New() js.Ref {
	return bindings.ImportJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Import) UpdateFrom(ref js.Ref) {
	bindings.ImportJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Import) Update(ref js.Ref) {
	bindings.ImportJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Import) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.MinimumVersion.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.MinimumVersion = p.MinimumVersion.FromRef(js.Undefined)
}

type ManifestKeys struct {
	// Import is "ManifestKeys.import"
	//
	// Optional
	Import js.Array[Import]
	// Export is "ManifestKeys.export"
	//
	// Optional
	//
	// NOTE: Export.FFI_USE MUST be set to true to get Export used.
	Export Export

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ManifestKeys with all fields set.
func (p ManifestKeys) FromRef(ref js.Ref) ManifestKeys {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ManifestKeys in the application heap.
func (p ManifestKeys) New() js.Ref {
	return bindings.ManifestKeysJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ManifestKeys) UpdateFrom(ref js.Ref) {
	bindings.ManifestKeysJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ManifestKeys) Update(ref js.Ref) {
	bindings.ManifestKeysJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ManifestKeys) FreeMembers(recursive bool) {
	js.Free(
		p.Import.Ref(),
	)
	p.Import = p.Import.FromRef(js.Undefined)
	if recursive {
		p.Export.FreeMembers(true)
	}
}

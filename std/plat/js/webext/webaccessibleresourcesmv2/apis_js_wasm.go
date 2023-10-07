// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webaccessibleresourcesmv2

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/webaccessibleresourcesmv2/bindings"
)

type ManifestKeys struct {
	// WebAccessibleResources is "ManifestKeys.web_accessible_resources"
	//
	// Optional
	WebAccessibleResources js.Array[js.String]

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
		p.WebAccessibleResources.Ref(),
	)
	p.WebAccessibleResources = p.WebAccessibleResources.FromRef(js.Undefined)
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webaccessibleresources

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/webaccessibleresources/bindings"
)

type WebAccessibleResource struct {
	// Resources is "WebAccessibleResource.resources"
	//
	// Optional
	Resources js.Array[js.String]
	// Matches is "WebAccessibleResource.matches"
	//
	// Optional
	Matches js.Array[js.String]
	// ExtensionIds is "WebAccessibleResource.extension_ids"
	//
	// Optional
	ExtensionIds js.Array[js.String]
	// UseDynamicUrl is "WebAccessibleResource.use_dynamic_url"
	//
	// Optional
	//
	// NOTE: FFI_USE_UseDynamicUrl MUST be set to true to make this field effective.
	UseDynamicUrl bool

	FFI_USE_UseDynamicUrl bool // for UseDynamicUrl.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebAccessibleResource with all fields set.
func (p WebAccessibleResource) FromRef(ref js.Ref) WebAccessibleResource {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebAccessibleResource in the application heap.
func (p WebAccessibleResource) New() js.Ref {
	return bindings.WebAccessibleResourceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WebAccessibleResource) UpdateFrom(ref js.Ref) {
	bindings.WebAccessibleResourceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebAccessibleResource) Update(ref js.Ref) {
	bindings.WebAccessibleResourceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebAccessibleResource) FreeMembers(recursive bool) {
	js.Free(
		p.Resources.Ref(),
		p.Matches.Ref(),
		p.ExtensionIds.Ref(),
	)
	p.Resources = p.Resources.FromRef(js.Undefined)
	p.Matches = p.Matches.FromRef(js.Undefined)
	p.ExtensionIds = p.ExtensionIds.FromRef(js.Undefined)
}

type ManifestKeys struct {
	// WebAccessibleResources is "ManifestKeys.web_accessible_resources"
	//
	// Optional
	WebAccessibleResources js.Array[WebAccessibleResource]

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

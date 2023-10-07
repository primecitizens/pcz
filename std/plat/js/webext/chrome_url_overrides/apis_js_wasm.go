// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package chrome_url_overrides

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/chrome_url_overrides/bindings"
)

type UrlOverrideInfo struct {
	// Newtab is "UrlOverrideInfo.newtab"
	//
	// Optional
	Newtab js.String
	// Bookmarks is "UrlOverrideInfo.bookmarks"
	//
	// Optional
	Bookmarks js.String
	// History is "UrlOverrideInfo.history"
	//
	// Optional
	History js.String
	// Activationmessage is "UrlOverrideInfo.activationmessage"
	//
	// Optional
	Activationmessage js.String
	// Keyboard is "UrlOverrideInfo.keyboard"
	//
	// Optional
	Keyboard js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UrlOverrideInfo with all fields set.
func (p UrlOverrideInfo) FromRef(ref js.Ref) UrlOverrideInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UrlOverrideInfo in the application heap.
func (p UrlOverrideInfo) New() js.Ref {
	return bindings.UrlOverrideInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UrlOverrideInfo) UpdateFrom(ref js.Ref) {
	bindings.UrlOverrideInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UrlOverrideInfo) Update(ref js.Ref) {
	bindings.UrlOverrideInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UrlOverrideInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Newtab.Ref(),
		p.Bookmarks.Ref(),
		p.History.Ref(),
		p.Activationmessage.Ref(),
		p.Keyboard.Ref(),
	)
	p.Newtab = p.Newtab.FromRef(js.Undefined)
	p.Bookmarks = p.Bookmarks.FromRef(js.Undefined)
	p.History = p.History.FromRef(js.Undefined)
	p.Activationmessage = p.Activationmessage.FromRef(js.Undefined)
	p.Keyboard = p.Keyboard.FromRef(js.Undefined)
}

type ManifestKeys struct {
	// ChromeUrlOverrides is "ManifestKeys.chrome_url_overrides"
	//
	// Optional
	//
	// NOTE: ChromeUrlOverrides.FFI_USE MUST be set to true to get ChromeUrlOverrides used.
	ChromeUrlOverrides UrlOverrideInfo

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
	if recursive {
		p.ChromeUrlOverrides.FreeMembers(true)
	}
}

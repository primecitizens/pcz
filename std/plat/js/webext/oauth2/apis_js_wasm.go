// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package oauth2

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/oauth2/bindings"
)

type OAuth2Info struct {
	// AutoApprove is "OAuth2Info.auto_approve"
	//
	// Optional
	//
	// NOTE: FFI_USE_AutoApprove MUST be set to true to make this field effective.
	AutoApprove bool
	// ClientId is "OAuth2Info.client_id"
	//
	// Optional
	ClientId js.String
	// Scopes is "OAuth2Info.scopes"
	//
	// Optional
	Scopes js.Array[js.String]

	FFI_USE_AutoApprove bool // for AutoApprove.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OAuth2Info with all fields set.
func (p OAuth2Info) FromRef(ref js.Ref) OAuth2Info {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OAuth2Info in the application heap.
func (p OAuth2Info) New() js.Ref {
	return bindings.OAuth2InfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OAuth2Info) UpdateFrom(ref js.Ref) {
	bindings.OAuth2InfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OAuth2Info) Update(ref js.Ref) {
	bindings.OAuth2InfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OAuth2Info) FreeMembers(recursive bool) {
	js.Free(
		p.ClientId.Ref(),
		p.Scopes.Ref(),
	)
	p.ClientId = p.ClientId.FromRef(js.Undefined)
	p.Scopes = p.Scopes.FromRef(js.Undefined)
}

type ManifestKeys struct {
	// Oauth2 is "ManifestKeys.oauth2"
	//
	// Optional
	//
	// NOTE: Oauth2.FFI_USE MUST be set to true to get Oauth2 used.
	Oauth2 OAuth2Info

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
		p.Oauth2.FreeMembers(true)
	}
}

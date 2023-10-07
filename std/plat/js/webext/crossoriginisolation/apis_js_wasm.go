// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package crossoriginisolation

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/crossoriginisolation/bindings"
)

type ResponseHeader struct {
	// Value is "ResponseHeader.value"
	//
	// Optional
	Value js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ResponseHeader with all fields set.
func (p ResponseHeader) FromRef(ref js.Ref) ResponseHeader {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ResponseHeader in the application heap.
func (p ResponseHeader) New() js.Ref {
	return bindings.ResponseHeaderJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ResponseHeader) UpdateFrom(ref js.Ref) {
	bindings.ResponseHeaderJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ResponseHeader) Update(ref js.Ref) {
	bindings.ResponseHeaderJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ResponseHeader) FreeMembers(recursive bool) {
	js.Free(
		p.Value.Ref(),
	)
	p.Value = p.Value.FromRef(js.Undefined)
}

type ManifestKeys struct {
	// CrossOriginEmbedderPolicy is "ManifestKeys.cross_origin_embedder_policy"
	//
	// Optional
	//
	// NOTE: CrossOriginEmbedderPolicy.FFI_USE MUST be set to true to get CrossOriginEmbedderPolicy used.
	CrossOriginEmbedderPolicy ResponseHeader
	// CrossOriginOpenerPolicy is "ManifestKeys.cross_origin_opener_policy"
	//
	// Optional
	//
	// NOTE: CrossOriginOpenerPolicy.FFI_USE MUST be set to true to get CrossOriginOpenerPolicy used.
	CrossOriginOpenerPolicy ResponseHeader

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
		p.CrossOriginEmbedderPolicy.FreeMembers(true)
		p.CrossOriginOpenerPolicy.FreeMembers(true)
	}
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package topsites

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/topsites/bindings"
)

type MostVisitedURL struct {
	// Title is "MostVisitedURL.title"
	//
	// Required
	Title js.String
	// Url is "MostVisitedURL.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MostVisitedURL with all fields set.
func (p MostVisitedURL) FromRef(ref js.Ref) MostVisitedURL {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MostVisitedURL in the application heap.
func (p MostVisitedURL) New() js.Ref {
	return bindings.MostVisitedURLJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MostVisitedURL) UpdateFrom(ref js.Ref) {
	bindings.MostVisitedURLJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MostVisitedURL) Update(ref js.Ref) {
	bindings.MostVisitedURLJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MostVisitedURL) FreeMembers(recursive bool) {
	js.Free(
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

// HasFuncGet returns true if the function "WEBEXT.topSites.get" exists.
func HasFuncGet() bool {
	return js.True == bindings.HasFuncGet()
}

// FuncGet returns the function "WEBEXT.topSites.get".
func FuncGet() (fn js.Func[func() js.Promise[js.Array[MostVisitedURL]]]) {
	bindings.FuncGet(
		js.Pointer(&fn),
	)
	return
}

// Get calls the function "WEBEXT.topSites.get" directly.
func Get() (ret js.Promise[js.Array[MostVisitedURL]]) {
	bindings.CallGet(
		js.Pointer(&ret),
	)

	return
}

// TryGet calls the function "WEBEXT.topSites.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGet() (ret js.Promise[js.Array[MostVisitedURL]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGet(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

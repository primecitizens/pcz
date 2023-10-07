// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package guestviewinternal

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/guestviewinternal/bindings"
)

type Size struct {
	// Height is "Size.height"
	//
	// Required
	Height int64
	// Width is "Size.width"
	//
	// Required
	Width int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Size with all fields set.
func (p Size) FromRef(ref js.Ref) Size {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Size in the application heap.
func (p Size) New() js.Ref {
	return bindings.SizeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Size) UpdateFrom(ref js.Ref) {
	bindings.SizeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Size) Update(ref js.Ref) {
	bindings.SizeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Size) FreeMembers(recursive bool) {
}

type SizeParams struct {
	// EnableAutoSize is "SizeParams.enableAutoSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_EnableAutoSize MUST be set to true to make this field effective.
	EnableAutoSize bool
	// Max is "SizeParams.max"
	//
	// Optional
	//
	// NOTE: Max.FFI_USE MUST be set to true to get Max used.
	Max Size
	// Min is "SizeParams.min"
	//
	// Optional
	//
	// NOTE: Min.FFI_USE MUST be set to true to get Min used.
	Min Size
	// Normal is "SizeParams.normal"
	//
	// Optional
	//
	// NOTE: Normal.FFI_USE MUST be set to true to get Normal used.
	Normal Size

	FFI_USE_EnableAutoSize bool // for EnableAutoSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SizeParams with all fields set.
func (p SizeParams) FromRef(ref js.Ref) SizeParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SizeParams in the application heap.
func (p SizeParams) New() js.Ref {
	return bindings.SizeParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SizeParams) UpdateFrom(ref js.Ref) {
	bindings.SizeParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SizeParams) Update(ref js.Ref) {
	bindings.SizeParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SizeParams) FreeMembers(recursive bool) {
	if recursive {
		p.Max.FreeMembers(true)
		p.Min.FreeMembers(true)
		p.Normal.FreeMembers(true)
	}
}

// HasFuncCreateGuest returns true if the function "WEBEXT.guestViewInternal.createGuest" exists.
func HasFuncCreateGuest() bool {
	return js.True == bindings.HasFuncCreateGuest()
}

// FuncCreateGuest returns the function "WEBEXT.guestViewInternal.createGuest".
func FuncCreateGuest() (fn js.Func[func(viewType js.String, ownerRoutingId int64, createParams js.Any) js.Promise[js.BigInt[int64]]]) {
	bindings.FuncCreateGuest(
		js.Pointer(&fn),
	)
	return
}

// CreateGuest calls the function "WEBEXT.guestViewInternal.createGuest" directly.
func CreateGuest(viewType js.String, ownerRoutingId int64, createParams js.Any) (ret js.Promise[js.BigInt[int64]]) {
	bindings.CallCreateGuest(
		js.Pointer(&ret),
		viewType.Ref(),
		float64(ownerRoutingId),
		createParams.Ref(),
	)

	return
}

// TryCreateGuest calls the function "WEBEXT.guestViewInternal.createGuest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreateGuest(viewType js.String, ownerRoutingId int64, createParams js.Any) (ret js.Promise[js.BigInt[int64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreateGuest(
		js.Pointer(&ret), js.Pointer(&exception),
		viewType.Ref(),
		float64(ownerRoutingId),
		createParams.Ref(),
	)

	return
}

// HasFuncDestroyUnattachedGuest returns true if the function "WEBEXT.guestViewInternal.destroyUnattachedGuest" exists.
func HasFuncDestroyUnattachedGuest() bool {
	return js.True == bindings.HasFuncDestroyUnattachedGuest()
}

// FuncDestroyUnattachedGuest returns the function "WEBEXT.guestViewInternal.destroyUnattachedGuest".
func FuncDestroyUnattachedGuest() (fn js.Func[func(instanceId int64)]) {
	bindings.FuncDestroyUnattachedGuest(
		js.Pointer(&fn),
	)
	return
}

// DestroyUnattachedGuest calls the function "WEBEXT.guestViewInternal.destroyUnattachedGuest" directly.
func DestroyUnattachedGuest(instanceId int64) (ret js.Void) {
	bindings.CallDestroyUnattachedGuest(
		js.Pointer(&ret),
		float64(instanceId),
	)

	return
}

// TryDestroyUnattachedGuest calls the function "WEBEXT.guestViewInternal.destroyUnattachedGuest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDestroyUnattachedGuest(instanceId int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDestroyUnattachedGuest(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
	)

	return
}

// HasFuncSetSize returns true if the function "WEBEXT.guestViewInternal.setSize" exists.
func HasFuncSetSize() bool {
	return js.True == bindings.HasFuncSetSize()
}

// FuncSetSize returns the function "WEBEXT.guestViewInternal.setSize".
func FuncSetSize() (fn js.Func[func(instanceId int64, params SizeParams) js.Promise[js.Void]]) {
	bindings.FuncSetSize(
		js.Pointer(&fn),
	)
	return
}

// SetSize calls the function "WEBEXT.guestViewInternal.setSize" directly.
func SetSize(instanceId int64, params SizeParams) (ret js.Promise[js.Void]) {
	bindings.CallSetSize(
		js.Pointer(&ret),
		float64(instanceId),
		js.Pointer(&params),
	)

	return
}

// TrySetSize calls the function "WEBEXT.guestViewInternal.setSize"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetSize(instanceId int64, params SizeParams) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetSize(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		js.Pointer(&params),
	)

	return
}

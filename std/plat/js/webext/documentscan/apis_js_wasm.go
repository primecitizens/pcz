// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package documentscan

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/documentscan/bindings"
)

type ScanCallbackFunc func(this js.Ref, result *ScanResults) js.Ref

func (fn ScanCallbackFunc) Register() js.Func[func(result *ScanResults)] {
	return js.RegisterCallback[func(result *ScanResults)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ScanCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ScanResults
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ScanCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *ScanResults) js.Ref
	Arg T
}

func (cb *ScanCallback[T]) Register() js.Func[func(result *ScanResults)] {
	return js.RegisterCallback[func(result *ScanResults)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ScanCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 ScanResults
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ScanResults struct {
	// DataUrls is "ScanResults.dataUrls"
	//
	// Optional
	DataUrls js.Array[js.String]
	// MimeType is "ScanResults.mimeType"
	//
	// Optional
	MimeType js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScanResults with all fields set.
func (p ScanResults) FromRef(ref js.Ref) ScanResults {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScanResults in the application heap.
func (p ScanResults) New() js.Ref {
	return bindings.ScanResultsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScanResults) UpdateFrom(ref js.Ref) {
	bindings.ScanResultsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScanResults) Update(ref js.Ref) {
	bindings.ScanResultsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScanResults) FreeMembers(recursive bool) {
	js.Free(
		p.DataUrls.Ref(),
		p.MimeType.Ref(),
	)
	p.DataUrls = p.DataUrls.FromRef(js.Undefined)
	p.MimeType = p.MimeType.FromRef(js.Undefined)
}

type ScanOptions struct {
	// MimeTypes is "ScanOptions.mimeTypes"
	//
	// Optional
	MimeTypes js.Array[js.String]
	// MaxImages is "ScanOptions.maxImages"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxImages MUST be set to true to make this field effective.
	MaxImages int32

	FFI_USE_MaxImages bool // for MaxImages.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScanOptions with all fields set.
func (p ScanOptions) FromRef(ref js.Ref) ScanOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScanOptions in the application heap.
func (p ScanOptions) New() js.Ref {
	return bindings.ScanOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScanOptions) UpdateFrom(ref js.Ref) {
	bindings.ScanOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScanOptions) Update(ref js.Ref) {
	bindings.ScanOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScanOptions) FreeMembers(recursive bool) {
	js.Free(
		p.MimeTypes.Ref(),
	)
	p.MimeTypes = p.MimeTypes.FromRef(js.Undefined)
}

// HasFuncScan returns true if the function "WEBEXT.documentScan.scan" exists.
func HasFuncScan() bool {
	return js.True == bindings.HasFuncScan()
}

// FuncScan returns the function "WEBEXT.documentScan.scan".
func FuncScan() (fn js.Func[func(options ScanOptions) js.Promise[ScanResults]]) {
	bindings.FuncScan(
		js.Pointer(&fn),
	)
	return
}

// Scan calls the function "WEBEXT.documentScan.scan" directly.
func Scan(options ScanOptions) (ret js.Promise[ScanResults]) {
	bindings.CallScan(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScan calls the function "WEBEXT.documentScan.scan"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryScan(options ScanOptions) (ret js.Promise[ScanResults], exception js.Any, ok bool) {
	ok = js.True == bindings.TryScan(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

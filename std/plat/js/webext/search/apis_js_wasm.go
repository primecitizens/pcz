// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package search

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/search/bindings"
)

type Disposition uint32

const (
	_ Disposition = iota

	Disposition_CURRENT_TAB
	Disposition_NEW_TAB
	Disposition_NEW_WINDOW
)

func (Disposition) FromRef(str js.Ref) Disposition {
	return Disposition(bindings.ConstOfDisposition(str))
}

func (x Disposition) String() (string, bool) {
	switch x {
	case Disposition_CURRENT_TAB:
		return "CURRENT_TAB", true
	case Disposition_NEW_TAB:
		return "NEW_TAB", true
	case Disposition_NEW_WINDOW:
		return "NEW_WINDOW", true
	default:
		return "", false
	}
}

type QueryCallbackFunc func(this js.Ref) js.Ref

func (fn QueryCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn QueryCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type QueryCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *QueryCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *QueryCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type QueryInfo struct {
	// Text is "QueryInfo.text"
	//
	// Optional
	Text js.String
	// Disposition is "QueryInfo.disposition"
	//
	// Optional
	Disposition Disposition
	// TabId is "QueryInfo.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32

	FFI_USE_TabId bool // for TabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a QueryInfo with all fields set.
func (p QueryInfo) FromRef(ref js.Ref) QueryInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new QueryInfo in the application heap.
func (p QueryInfo) New() js.Ref {
	return bindings.QueryInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *QueryInfo) UpdateFrom(ref js.Ref) {
	bindings.QueryInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *QueryInfo) Update(ref js.Ref) {
	bindings.QueryInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *QueryInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Text.Ref(),
	)
	p.Text = p.Text.FromRef(js.Undefined)
}

// HasFuncQuery returns true if the function "WEBEXT.search.query" exists.
func HasFuncQuery() bool {
	return js.True == bindings.HasFuncQuery()
}

// FuncQuery returns the function "WEBEXT.search.query".
func FuncQuery() (fn js.Func[func(queryInfo QueryInfo) js.Promise[js.Void]]) {
	bindings.FuncQuery(
		js.Pointer(&fn),
	)
	return
}

// Query calls the function "WEBEXT.search.query" directly.
func Query(queryInfo QueryInfo) (ret js.Promise[js.Void]) {
	bindings.CallQuery(
		js.Pointer(&ret),
		js.Pointer(&queryInfo),
	)

	return
}

// TryQuery calls the function "WEBEXT.search.query"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryQuery(queryInfo QueryInfo) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryQuery(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&queryInfo),
	)

	return
}

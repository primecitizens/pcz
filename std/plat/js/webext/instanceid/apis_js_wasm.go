// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package instanceid

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/instanceid/bindings"
)

type DeleteTokenArgDeleteTokenParams struct {
	// AuthorizedEntity is "DeleteTokenArgDeleteTokenParams.authorizedEntity"
	//
	// Required
	AuthorizedEntity js.String
	// Scope is "DeleteTokenArgDeleteTokenParams.scope"
	//
	// Required
	Scope js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeleteTokenArgDeleteTokenParams with all fields set.
func (p DeleteTokenArgDeleteTokenParams) FromRef(ref js.Ref) DeleteTokenArgDeleteTokenParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeleteTokenArgDeleteTokenParams in the application heap.
func (p DeleteTokenArgDeleteTokenParams) New() js.Ref {
	return bindings.DeleteTokenArgDeleteTokenParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeleteTokenArgDeleteTokenParams) UpdateFrom(ref js.Ref) {
	bindings.DeleteTokenArgDeleteTokenParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeleteTokenArgDeleteTokenParams) Update(ref js.Ref) {
	bindings.DeleteTokenArgDeleteTokenParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeleteTokenArgDeleteTokenParams) FreeMembers(recursive bool) {
	js.Free(
		p.AuthorizedEntity.Ref(),
		p.Scope.Ref(),
	)
	p.AuthorizedEntity = p.AuthorizedEntity.FromRef(js.Undefined)
	p.Scope = p.Scope.FromRef(js.Undefined)
}

type GetTokenArgGetTokenParams struct {
	// AuthorizedEntity is "GetTokenArgGetTokenParams.authorizedEntity"
	//
	// Required
	AuthorizedEntity js.String
	// Options is "GetTokenArgGetTokenParams.options"
	//
	// Optional
	Options js.String
	// Scope is "GetTokenArgGetTokenParams.scope"
	//
	// Required
	Scope js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetTokenArgGetTokenParams with all fields set.
func (p GetTokenArgGetTokenParams) FromRef(ref js.Ref) GetTokenArgGetTokenParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetTokenArgGetTokenParams in the application heap.
func (p GetTokenArgGetTokenParams) New() js.Ref {
	return bindings.GetTokenArgGetTokenParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetTokenArgGetTokenParams) UpdateFrom(ref js.Ref) {
	bindings.GetTokenArgGetTokenParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetTokenArgGetTokenParams) Update(ref js.Ref) {
	bindings.GetTokenArgGetTokenParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetTokenArgGetTokenParams) FreeMembers(recursive bool) {
	js.Free(
		p.AuthorizedEntity.Ref(),
		p.Options.Ref(),
		p.Scope.Ref(),
	)
	p.AuthorizedEntity = p.AuthorizedEntity.FromRef(js.Undefined)
	p.Options = p.Options.FromRef(js.Undefined)
	p.Scope = p.Scope.FromRef(js.Undefined)
}

// HasFuncDeleteID returns true if the function "WEBEXT.instanceID.deleteID" exists.
func HasFuncDeleteID() bool {
	return js.True == bindings.HasFuncDeleteID()
}

// FuncDeleteID returns the function "WEBEXT.instanceID.deleteID".
func FuncDeleteID() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncDeleteID(
		js.Pointer(&fn),
	)
	return
}

// DeleteID calls the function "WEBEXT.instanceID.deleteID" directly.
func DeleteID() (ret js.Promise[js.Void]) {
	bindings.CallDeleteID(
		js.Pointer(&ret),
	)

	return
}

// TryDeleteID calls the function "WEBEXT.instanceID.deleteID"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteID() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteID(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteToken returns true if the function "WEBEXT.instanceID.deleteToken" exists.
func HasFuncDeleteToken() bool {
	return js.True == bindings.HasFuncDeleteToken()
}

// FuncDeleteToken returns the function "WEBEXT.instanceID.deleteToken".
func FuncDeleteToken() (fn js.Func[func(deleteTokenParams DeleteTokenArgDeleteTokenParams) js.Promise[js.Void]]) {
	bindings.FuncDeleteToken(
		js.Pointer(&fn),
	)
	return
}

// DeleteToken calls the function "WEBEXT.instanceID.deleteToken" directly.
func DeleteToken(deleteTokenParams DeleteTokenArgDeleteTokenParams) (ret js.Promise[js.Void]) {
	bindings.CallDeleteToken(
		js.Pointer(&ret),
		js.Pointer(&deleteTokenParams),
	)

	return
}

// TryDeleteToken calls the function "WEBEXT.instanceID.deleteToken"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteToken(deleteTokenParams DeleteTokenArgDeleteTokenParams) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteToken(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&deleteTokenParams),
	)

	return
}

// HasFuncGetCreationTime returns true if the function "WEBEXT.instanceID.getCreationTime" exists.
func HasFuncGetCreationTime() bool {
	return js.True == bindings.HasFuncGetCreationTime()
}

// FuncGetCreationTime returns the function "WEBEXT.instanceID.getCreationTime".
func FuncGetCreationTime() (fn js.Func[func() js.Promise[js.Number[float64]]]) {
	bindings.FuncGetCreationTime(
		js.Pointer(&fn),
	)
	return
}

// GetCreationTime calls the function "WEBEXT.instanceID.getCreationTime" directly.
func GetCreationTime() (ret js.Promise[js.Number[float64]]) {
	bindings.CallGetCreationTime(
		js.Pointer(&ret),
	)

	return
}

// TryGetCreationTime calls the function "WEBEXT.instanceID.getCreationTime"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCreationTime() (ret js.Promise[js.Number[float64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCreationTime(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetID returns true if the function "WEBEXT.instanceID.getID" exists.
func HasFuncGetID() bool {
	return js.True == bindings.HasFuncGetID()
}

// FuncGetID returns the function "WEBEXT.instanceID.getID".
func FuncGetID() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetID(
		js.Pointer(&fn),
	)
	return
}

// GetID calls the function "WEBEXT.instanceID.getID" directly.
func GetID() (ret js.Promise[js.String]) {
	bindings.CallGetID(
		js.Pointer(&ret),
	)

	return
}

// TryGetID calls the function "WEBEXT.instanceID.getID"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetID() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetID(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetToken returns true if the function "WEBEXT.instanceID.getToken" exists.
func HasFuncGetToken() bool {
	return js.True == bindings.HasFuncGetToken()
}

// FuncGetToken returns the function "WEBEXT.instanceID.getToken".
func FuncGetToken() (fn js.Func[func(getTokenParams GetTokenArgGetTokenParams) js.Promise[js.String]]) {
	bindings.FuncGetToken(
		js.Pointer(&fn),
	)
	return
}

// GetToken calls the function "WEBEXT.instanceID.getToken" directly.
func GetToken(getTokenParams GetTokenArgGetTokenParams) (ret js.Promise[js.String]) {
	bindings.CallGetToken(
		js.Pointer(&ret),
		js.Pointer(&getTokenParams),
	)

	return
}

// TryGetToken calls the function "WEBEXT.instanceID.getToken"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetToken(getTokenParams GetTokenArgGetTokenParams) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetToken(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&getTokenParams),
	)

	return
}

type OnTokenRefreshEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnTokenRefreshEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnTokenRefreshEventCallbackFunc) DispatchCallback(
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

type OnTokenRefreshEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnTokenRefreshEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnTokenRefreshEventCallback[T]) DispatchCallback(
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

// HasFuncOnTokenRefresh returns true if the function "WEBEXT.instanceID.onTokenRefresh.addListener" exists.
func HasFuncOnTokenRefresh() bool {
	return js.True == bindings.HasFuncOnTokenRefresh()
}

// FuncOnTokenRefresh returns the function "WEBEXT.instanceID.onTokenRefresh.addListener".
func FuncOnTokenRefresh() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnTokenRefresh(
		js.Pointer(&fn),
	)
	return
}

// OnTokenRefresh calls the function "WEBEXT.instanceID.onTokenRefresh.addListener" directly.
func OnTokenRefresh(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnTokenRefresh(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnTokenRefresh calls the function "WEBEXT.instanceID.onTokenRefresh.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnTokenRefresh(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnTokenRefresh(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffTokenRefresh returns true if the function "WEBEXT.instanceID.onTokenRefresh.removeListener" exists.
func HasFuncOffTokenRefresh() bool {
	return js.True == bindings.HasFuncOffTokenRefresh()
}

// FuncOffTokenRefresh returns the function "WEBEXT.instanceID.onTokenRefresh.removeListener".
func FuncOffTokenRefresh() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffTokenRefresh(
		js.Pointer(&fn),
	)
	return
}

// OffTokenRefresh calls the function "WEBEXT.instanceID.onTokenRefresh.removeListener" directly.
func OffTokenRefresh(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffTokenRefresh(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffTokenRefresh calls the function "WEBEXT.instanceID.onTokenRefresh.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffTokenRefresh(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffTokenRefresh(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnTokenRefresh returns true if the function "WEBEXT.instanceID.onTokenRefresh.hasListener" exists.
func HasFuncHasOnTokenRefresh() bool {
	return js.True == bindings.HasFuncHasOnTokenRefresh()
}

// FuncHasOnTokenRefresh returns the function "WEBEXT.instanceID.onTokenRefresh.hasListener".
func FuncHasOnTokenRefresh() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnTokenRefresh(
		js.Pointer(&fn),
	)
	return
}

// HasOnTokenRefresh calls the function "WEBEXT.instanceID.onTokenRefresh.hasListener" directly.
func HasOnTokenRefresh(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnTokenRefresh(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnTokenRefresh calls the function "WEBEXT.instanceID.onTokenRefresh.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnTokenRefresh(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnTokenRefresh(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

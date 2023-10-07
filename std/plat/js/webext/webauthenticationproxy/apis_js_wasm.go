// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webauthenticationproxy

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/webauthenticationproxy/bindings"
)

type CreateRequest struct {
	// RequestId is "CreateRequest.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// RequestDetailsJson is "CreateRequest.requestDetailsJson"
	//
	// Optional
	RequestDetailsJson js.String

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateRequest with all fields set.
func (p CreateRequest) FromRef(ref js.Ref) CreateRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateRequest in the application heap.
func (p CreateRequest) New() js.Ref {
	return bindings.CreateRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateRequest) UpdateFrom(ref js.Ref) {
	bindings.CreateRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateRequest) Update(ref js.Ref) {
	bindings.CreateRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateRequest) FreeMembers(recursive bool) {
	js.Free(
		p.RequestDetailsJson.Ref(),
	)
	p.RequestDetailsJson = p.RequestDetailsJson.FromRef(js.Undefined)
}

type DOMExceptionDetails struct {
	// Name is "DOMExceptionDetails.name"
	//
	// Optional
	Name js.String
	// Message is "DOMExceptionDetails.message"
	//
	// Optional
	Message js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DOMExceptionDetails with all fields set.
func (p DOMExceptionDetails) FromRef(ref js.Ref) DOMExceptionDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DOMExceptionDetails in the application heap.
func (p DOMExceptionDetails) New() js.Ref {
	return bindings.DOMExceptionDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DOMExceptionDetails) UpdateFrom(ref js.Ref) {
	bindings.DOMExceptionDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DOMExceptionDetails) Update(ref js.Ref) {
	bindings.DOMExceptionDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DOMExceptionDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
		p.Message.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	p.Message = p.Message.FromRef(js.Undefined)
}

type CreateResponseDetails struct {
	// RequestId is "CreateResponseDetails.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// Error is "CreateResponseDetails.error"
	//
	// Optional
	//
	// NOTE: Error.FFI_USE MUST be set to true to get Error used.
	Error DOMExceptionDetails
	// ResponseJson is "CreateResponseDetails.responseJson"
	//
	// Optional
	ResponseJson js.String

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CreateResponseDetails with all fields set.
func (p CreateResponseDetails) FromRef(ref js.Ref) CreateResponseDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CreateResponseDetails in the application heap.
func (p CreateResponseDetails) New() js.Ref {
	return bindings.CreateResponseDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CreateResponseDetails) UpdateFrom(ref js.Ref) {
	bindings.CreateResponseDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CreateResponseDetails) Update(ref js.Ref) {
	bindings.CreateResponseDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CreateResponseDetails) FreeMembers(recursive bool) {
	js.Free(
		p.ResponseJson.Ref(),
	)
	p.ResponseJson = p.ResponseJson.FromRef(js.Undefined)
	if recursive {
		p.Error.FreeMembers(true)
	}
}

type ErrorCallbackFunc func(this js.Ref, err js.String) js.Ref

func (fn ErrorCallbackFunc) Register() js.Func[func(err js.String)] {
	return js.RegisterCallback[func(err js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ErrorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ErrorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, err js.String) js.Ref
	Arg T
}

func (cb *ErrorCallback[T]) Register() js.Func[func(err js.String)] {
	return js.RegisterCallback[func(err js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ErrorCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetRequest struct {
	// RequestId is "GetRequest.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// RequestDetailsJson is "GetRequest.requestDetailsJson"
	//
	// Optional
	RequestDetailsJson js.String

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetRequest with all fields set.
func (p GetRequest) FromRef(ref js.Ref) GetRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetRequest in the application heap.
func (p GetRequest) New() js.Ref {
	return bindings.GetRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetRequest) UpdateFrom(ref js.Ref) {
	bindings.GetRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetRequest) Update(ref js.Ref) {
	bindings.GetRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetRequest) FreeMembers(recursive bool) {
	js.Free(
		p.RequestDetailsJson.Ref(),
	)
	p.RequestDetailsJson = p.RequestDetailsJson.FromRef(js.Undefined)
}

type GetResponseDetails struct {
	// RequestId is "GetResponseDetails.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// Error is "GetResponseDetails.error"
	//
	// Optional
	//
	// NOTE: Error.FFI_USE MUST be set to true to get Error used.
	Error DOMExceptionDetails
	// ResponseJson is "GetResponseDetails.responseJson"
	//
	// Optional
	ResponseJson js.String

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetResponseDetails with all fields set.
func (p GetResponseDetails) FromRef(ref js.Ref) GetResponseDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetResponseDetails in the application heap.
func (p GetResponseDetails) New() js.Ref {
	return bindings.GetResponseDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetResponseDetails) UpdateFrom(ref js.Ref) {
	bindings.GetResponseDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetResponseDetails) Update(ref js.Ref) {
	bindings.GetResponseDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetResponseDetails) FreeMembers(recursive bool) {
	js.Free(
		p.ResponseJson.Ref(),
	)
	p.ResponseJson = p.ResponseJson.FromRef(js.Undefined)
	if recursive {
		p.Error.FreeMembers(true)
	}
}

type IsUvpaaRequest struct {
	// RequestId is "IsUvpaaRequest.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32

	FFI_USE_RequestId bool // for RequestId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IsUvpaaRequest with all fields set.
func (p IsUvpaaRequest) FromRef(ref js.Ref) IsUvpaaRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IsUvpaaRequest in the application heap.
func (p IsUvpaaRequest) New() js.Ref {
	return bindings.IsUvpaaRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IsUvpaaRequest) UpdateFrom(ref js.Ref) {
	bindings.IsUvpaaRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IsUvpaaRequest) Update(ref js.Ref) {
	bindings.IsUvpaaRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IsUvpaaRequest) FreeMembers(recursive bool) {
}

type IsUvpaaResponseDetails struct {
	// RequestId is "IsUvpaaResponseDetails.requestId"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestId MUST be set to true to make this field effective.
	RequestId int32
	// IsUvpaa is "IsUvpaaResponseDetails.isUvpaa"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsUvpaa MUST be set to true to make this field effective.
	IsUvpaa bool

	FFI_USE_RequestId bool // for RequestId.
	FFI_USE_IsUvpaa   bool // for IsUvpaa.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IsUvpaaResponseDetails with all fields set.
func (p IsUvpaaResponseDetails) FromRef(ref js.Ref) IsUvpaaResponseDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IsUvpaaResponseDetails in the application heap.
func (p IsUvpaaResponseDetails) New() js.Ref {
	return bindings.IsUvpaaResponseDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *IsUvpaaResponseDetails) UpdateFrom(ref js.Ref) {
	bindings.IsUvpaaResponseDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IsUvpaaResponseDetails) Update(ref js.Ref) {
	bindings.IsUvpaaResponseDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IsUvpaaResponseDetails) FreeMembers(recursive bool) {
}

type VoidCallbackFunc func(this js.Ref) js.Ref

func (fn VoidCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn VoidCallbackFunc) DispatchCallback(
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

type VoidCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *VoidCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *VoidCallback[T]) DispatchCallback(
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

// HasFuncAttach returns true if the function "WEBEXT.webAuthenticationProxy.attach" exists.
func HasFuncAttach() bool {
	return js.True == bindings.HasFuncAttach()
}

// FuncAttach returns the function "WEBEXT.webAuthenticationProxy.attach".
func FuncAttach() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncAttach(
		js.Pointer(&fn),
	)
	return
}

// Attach calls the function "WEBEXT.webAuthenticationProxy.attach" directly.
func Attach() (ret js.Promise[js.String]) {
	bindings.CallAttach(
		js.Pointer(&ret),
	)

	return
}

// TryAttach calls the function "WEBEXT.webAuthenticationProxy.attach"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAttach() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAttach(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCompleteCreateRequest returns true if the function "WEBEXT.webAuthenticationProxy.completeCreateRequest" exists.
func HasFuncCompleteCreateRequest() bool {
	return js.True == bindings.HasFuncCompleteCreateRequest()
}

// FuncCompleteCreateRequest returns the function "WEBEXT.webAuthenticationProxy.completeCreateRequest".
func FuncCompleteCreateRequest() (fn js.Func[func(details CreateResponseDetails) js.Promise[js.Void]]) {
	bindings.FuncCompleteCreateRequest(
		js.Pointer(&fn),
	)
	return
}

// CompleteCreateRequest calls the function "WEBEXT.webAuthenticationProxy.completeCreateRequest" directly.
func CompleteCreateRequest(details CreateResponseDetails) (ret js.Promise[js.Void]) {
	bindings.CallCompleteCreateRequest(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryCompleteCreateRequest calls the function "WEBEXT.webAuthenticationProxy.completeCreateRequest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCompleteCreateRequest(details CreateResponseDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompleteCreateRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncCompleteGetRequest returns true if the function "WEBEXT.webAuthenticationProxy.completeGetRequest" exists.
func HasFuncCompleteGetRequest() bool {
	return js.True == bindings.HasFuncCompleteGetRequest()
}

// FuncCompleteGetRequest returns the function "WEBEXT.webAuthenticationProxy.completeGetRequest".
func FuncCompleteGetRequest() (fn js.Func[func(details GetResponseDetails) js.Promise[js.Void]]) {
	bindings.FuncCompleteGetRequest(
		js.Pointer(&fn),
	)
	return
}

// CompleteGetRequest calls the function "WEBEXT.webAuthenticationProxy.completeGetRequest" directly.
func CompleteGetRequest(details GetResponseDetails) (ret js.Promise[js.Void]) {
	bindings.CallCompleteGetRequest(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryCompleteGetRequest calls the function "WEBEXT.webAuthenticationProxy.completeGetRequest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCompleteGetRequest(details GetResponseDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompleteGetRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncCompleteIsUvpaaRequest returns true if the function "WEBEXT.webAuthenticationProxy.completeIsUvpaaRequest" exists.
func HasFuncCompleteIsUvpaaRequest() bool {
	return js.True == bindings.HasFuncCompleteIsUvpaaRequest()
}

// FuncCompleteIsUvpaaRequest returns the function "WEBEXT.webAuthenticationProxy.completeIsUvpaaRequest".
func FuncCompleteIsUvpaaRequest() (fn js.Func[func(details IsUvpaaResponseDetails) js.Promise[js.Void]]) {
	bindings.FuncCompleteIsUvpaaRequest(
		js.Pointer(&fn),
	)
	return
}

// CompleteIsUvpaaRequest calls the function "WEBEXT.webAuthenticationProxy.completeIsUvpaaRequest" directly.
func CompleteIsUvpaaRequest(details IsUvpaaResponseDetails) (ret js.Promise[js.Void]) {
	bindings.CallCompleteIsUvpaaRequest(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryCompleteIsUvpaaRequest calls the function "WEBEXT.webAuthenticationProxy.completeIsUvpaaRequest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCompleteIsUvpaaRequest(details IsUvpaaResponseDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCompleteIsUvpaaRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncDetach returns true if the function "WEBEXT.webAuthenticationProxy.detach" exists.
func HasFuncDetach() bool {
	return js.True == bindings.HasFuncDetach()
}

// FuncDetach returns the function "WEBEXT.webAuthenticationProxy.detach".
func FuncDetach() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncDetach(
		js.Pointer(&fn),
	)
	return
}

// Detach calls the function "WEBEXT.webAuthenticationProxy.detach" directly.
func Detach() (ret js.Promise[js.String]) {
	bindings.CallDetach(
		js.Pointer(&ret),
	)

	return
}

// TryDetach calls the function "WEBEXT.webAuthenticationProxy.detach"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDetach() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDetach(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnCreateRequestEventCallbackFunc func(this js.Ref, requestInfo *CreateRequest) js.Ref

func (fn OnCreateRequestEventCallbackFunc) Register() js.Func[func(requestInfo *CreateRequest)] {
	return js.RegisterCallback[func(requestInfo *CreateRequest)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnCreateRequestEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CreateRequest
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

type OnCreateRequestEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestInfo *CreateRequest) js.Ref
	Arg T
}

func (cb *OnCreateRequestEventCallback[T]) Register() js.Func[func(requestInfo *CreateRequest)] {
	return js.RegisterCallback[func(requestInfo *CreateRequest)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnCreateRequestEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CreateRequest
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

// HasFuncOnCreateRequest returns true if the function "WEBEXT.webAuthenticationProxy.onCreateRequest.addListener" exists.
func HasFuncOnCreateRequest() bool {
	return js.True == bindings.HasFuncOnCreateRequest()
}

// FuncOnCreateRequest returns the function "WEBEXT.webAuthenticationProxy.onCreateRequest.addListener".
func FuncOnCreateRequest() (fn js.Func[func(callback js.Func[func(requestInfo *CreateRequest)])]) {
	bindings.FuncOnCreateRequest(
		js.Pointer(&fn),
	)
	return
}

// OnCreateRequest calls the function "WEBEXT.webAuthenticationProxy.onCreateRequest.addListener" directly.
func OnCreateRequest(callback js.Func[func(requestInfo *CreateRequest)]) (ret js.Void) {
	bindings.CallOnCreateRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnCreateRequest calls the function "WEBEXT.webAuthenticationProxy.onCreateRequest.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnCreateRequest(callback js.Func[func(requestInfo *CreateRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnCreateRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffCreateRequest returns true if the function "WEBEXT.webAuthenticationProxy.onCreateRequest.removeListener" exists.
func HasFuncOffCreateRequest() bool {
	return js.True == bindings.HasFuncOffCreateRequest()
}

// FuncOffCreateRequest returns the function "WEBEXT.webAuthenticationProxy.onCreateRequest.removeListener".
func FuncOffCreateRequest() (fn js.Func[func(callback js.Func[func(requestInfo *CreateRequest)])]) {
	bindings.FuncOffCreateRequest(
		js.Pointer(&fn),
	)
	return
}

// OffCreateRequest calls the function "WEBEXT.webAuthenticationProxy.onCreateRequest.removeListener" directly.
func OffCreateRequest(callback js.Func[func(requestInfo *CreateRequest)]) (ret js.Void) {
	bindings.CallOffCreateRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffCreateRequest calls the function "WEBEXT.webAuthenticationProxy.onCreateRequest.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffCreateRequest(callback js.Func[func(requestInfo *CreateRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffCreateRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnCreateRequest returns true if the function "WEBEXT.webAuthenticationProxy.onCreateRequest.hasListener" exists.
func HasFuncHasOnCreateRequest() bool {
	return js.True == bindings.HasFuncHasOnCreateRequest()
}

// FuncHasOnCreateRequest returns the function "WEBEXT.webAuthenticationProxy.onCreateRequest.hasListener".
func FuncHasOnCreateRequest() (fn js.Func[func(callback js.Func[func(requestInfo *CreateRequest)]) bool]) {
	bindings.FuncHasOnCreateRequest(
		js.Pointer(&fn),
	)
	return
}

// HasOnCreateRequest calls the function "WEBEXT.webAuthenticationProxy.onCreateRequest.hasListener" directly.
func HasOnCreateRequest(callback js.Func[func(requestInfo *CreateRequest)]) (ret bool) {
	bindings.CallHasOnCreateRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnCreateRequest calls the function "WEBEXT.webAuthenticationProxy.onCreateRequest.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnCreateRequest(callback js.Func[func(requestInfo *CreateRequest)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnCreateRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnGetRequestEventCallbackFunc func(this js.Ref, requestInfo *GetRequest) js.Ref

func (fn OnGetRequestEventCallbackFunc) Register() js.Func[func(requestInfo *GetRequest)] {
	return js.RegisterCallback[func(requestInfo *GetRequest)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnGetRequestEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetRequest
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

type OnGetRequestEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestInfo *GetRequest) js.Ref
	Arg T
}

func (cb *OnGetRequestEventCallback[T]) Register() js.Func[func(requestInfo *GetRequest)] {
	return js.RegisterCallback[func(requestInfo *GetRequest)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnGetRequestEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 GetRequest
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

// HasFuncOnGetRequest returns true if the function "WEBEXT.webAuthenticationProxy.onGetRequest.addListener" exists.
func HasFuncOnGetRequest() bool {
	return js.True == bindings.HasFuncOnGetRequest()
}

// FuncOnGetRequest returns the function "WEBEXT.webAuthenticationProxy.onGetRequest.addListener".
func FuncOnGetRequest() (fn js.Func[func(callback js.Func[func(requestInfo *GetRequest)])]) {
	bindings.FuncOnGetRequest(
		js.Pointer(&fn),
	)
	return
}

// OnGetRequest calls the function "WEBEXT.webAuthenticationProxy.onGetRequest.addListener" directly.
func OnGetRequest(callback js.Func[func(requestInfo *GetRequest)]) (ret js.Void) {
	bindings.CallOnGetRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnGetRequest calls the function "WEBEXT.webAuthenticationProxy.onGetRequest.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnGetRequest(callback js.Func[func(requestInfo *GetRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnGetRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffGetRequest returns true if the function "WEBEXT.webAuthenticationProxy.onGetRequest.removeListener" exists.
func HasFuncOffGetRequest() bool {
	return js.True == bindings.HasFuncOffGetRequest()
}

// FuncOffGetRequest returns the function "WEBEXT.webAuthenticationProxy.onGetRequest.removeListener".
func FuncOffGetRequest() (fn js.Func[func(callback js.Func[func(requestInfo *GetRequest)])]) {
	bindings.FuncOffGetRequest(
		js.Pointer(&fn),
	)
	return
}

// OffGetRequest calls the function "WEBEXT.webAuthenticationProxy.onGetRequest.removeListener" directly.
func OffGetRequest(callback js.Func[func(requestInfo *GetRequest)]) (ret js.Void) {
	bindings.CallOffGetRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffGetRequest calls the function "WEBEXT.webAuthenticationProxy.onGetRequest.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffGetRequest(callback js.Func[func(requestInfo *GetRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffGetRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnGetRequest returns true if the function "WEBEXT.webAuthenticationProxy.onGetRequest.hasListener" exists.
func HasFuncHasOnGetRequest() bool {
	return js.True == bindings.HasFuncHasOnGetRequest()
}

// FuncHasOnGetRequest returns the function "WEBEXT.webAuthenticationProxy.onGetRequest.hasListener".
func FuncHasOnGetRequest() (fn js.Func[func(callback js.Func[func(requestInfo *GetRequest)]) bool]) {
	bindings.FuncHasOnGetRequest(
		js.Pointer(&fn),
	)
	return
}

// HasOnGetRequest calls the function "WEBEXT.webAuthenticationProxy.onGetRequest.hasListener" directly.
func HasOnGetRequest(callback js.Func[func(requestInfo *GetRequest)]) (ret bool) {
	bindings.CallHasOnGetRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnGetRequest calls the function "WEBEXT.webAuthenticationProxy.onGetRequest.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnGetRequest(callback js.Func[func(requestInfo *GetRequest)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnGetRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnIsUvpaaRequestEventCallbackFunc func(this js.Ref, requestInfo *IsUvpaaRequest) js.Ref

func (fn OnIsUvpaaRequestEventCallbackFunc) Register() js.Func[func(requestInfo *IsUvpaaRequest)] {
	return js.RegisterCallback[func(requestInfo *IsUvpaaRequest)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnIsUvpaaRequestEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 IsUvpaaRequest
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

type OnIsUvpaaRequestEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestInfo *IsUvpaaRequest) js.Ref
	Arg T
}

func (cb *OnIsUvpaaRequestEventCallback[T]) Register() js.Func[func(requestInfo *IsUvpaaRequest)] {
	return js.RegisterCallback[func(requestInfo *IsUvpaaRequest)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnIsUvpaaRequestEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 IsUvpaaRequest
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

// HasFuncOnIsUvpaaRequest returns true if the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.addListener" exists.
func HasFuncOnIsUvpaaRequest() bool {
	return js.True == bindings.HasFuncOnIsUvpaaRequest()
}

// FuncOnIsUvpaaRequest returns the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.addListener".
func FuncOnIsUvpaaRequest() (fn js.Func[func(callback js.Func[func(requestInfo *IsUvpaaRequest)])]) {
	bindings.FuncOnIsUvpaaRequest(
		js.Pointer(&fn),
	)
	return
}

// OnIsUvpaaRequest calls the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.addListener" directly.
func OnIsUvpaaRequest(callback js.Func[func(requestInfo *IsUvpaaRequest)]) (ret js.Void) {
	bindings.CallOnIsUvpaaRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnIsUvpaaRequest calls the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnIsUvpaaRequest(callback js.Func[func(requestInfo *IsUvpaaRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnIsUvpaaRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffIsUvpaaRequest returns true if the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.removeListener" exists.
func HasFuncOffIsUvpaaRequest() bool {
	return js.True == bindings.HasFuncOffIsUvpaaRequest()
}

// FuncOffIsUvpaaRequest returns the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.removeListener".
func FuncOffIsUvpaaRequest() (fn js.Func[func(callback js.Func[func(requestInfo *IsUvpaaRequest)])]) {
	bindings.FuncOffIsUvpaaRequest(
		js.Pointer(&fn),
	)
	return
}

// OffIsUvpaaRequest calls the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.removeListener" directly.
func OffIsUvpaaRequest(callback js.Func[func(requestInfo *IsUvpaaRequest)]) (ret js.Void) {
	bindings.CallOffIsUvpaaRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffIsUvpaaRequest calls the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffIsUvpaaRequest(callback js.Func[func(requestInfo *IsUvpaaRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffIsUvpaaRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnIsUvpaaRequest returns true if the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.hasListener" exists.
func HasFuncHasOnIsUvpaaRequest() bool {
	return js.True == bindings.HasFuncHasOnIsUvpaaRequest()
}

// FuncHasOnIsUvpaaRequest returns the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.hasListener".
func FuncHasOnIsUvpaaRequest() (fn js.Func[func(callback js.Func[func(requestInfo *IsUvpaaRequest)]) bool]) {
	bindings.FuncHasOnIsUvpaaRequest(
		js.Pointer(&fn),
	)
	return
}

// HasOnIsUvpaaRequest calls the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.hasListener" directly.
func HasOnIsUvpaaRequest(callback js.Func[func(requestInfo *IsUvpaaRequest)]) (ret bool) {
	bindings.CallHasOnIsUvpaaRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnIsUvpaaRequest calls the function "WEBEXT.webAuthenticationProxy.onIsUvpaaRequest.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnIsUvpaaRequest(callback js.Func[func(requestInfo *IsUvpaaRequest)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnIsUvpaaRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRemoteSessionStateChangeEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnRemoteSessionStateChangeEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRemoteSessionStateChangeEventCallbackFunc) DispatchCallback(
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

type OnRemoteSessionStateChangeEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnRemoteSessionStateChangeEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRemoteSessionStateChangeEventCallback[T]) DispatchCallback(
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

// HasFuncOnRemoteSessionStateChange returns true if the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.addListener" exists.
func HasFuncOnRemoteSessionStateChange() bool {
	return js.True == bindings.HasFuncOnRemoteSessionStateChange()
}

// FuncOnRemoteSessionStateChange returns the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.addListener".
func FuncOnRemoteSessionStateChange() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnRemoteSessionStateChange(
		js.Pointer(&fn),
	)
	return
}

// OnRemoteSessionStateChange calls the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.addListener" directly.
func OnRemoteSessionStateChange(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnRemoteSessionStateChange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRemoteSessionStateChange calls the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRemoteSessionStateChange(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRemoteSessionStateChange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRemoteSessionStateChange returns true if the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.removeListener" exists.
func HasFuncOffRemoteSessionStateChange() bool {
	return js.True == bindings.HasFuncOffRemoteSessionStateChange()
}

// FuncOffRemoteSessionStateChange returns the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.removeListener".
func FuncOffRemoteSessionStateChange() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffRemoteSessionStateChange(
		js.Pointer(&fn),
	)
	return
}

// OffRemoteSessionStateChange calls the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.removeListener" directly.
func OffRemoteSessionStateChange(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffRemoteSessionStateChange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRemoteSessionStateChange calls the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRemoteSessionStateChange(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRemoteSessionStateChange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRemoteSessionStateChange returns true if the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.hasListener" exists.
func HasFuncHasOnRemoteSessionStateChange() bool {
	return js.True == bindings.HasFuncHasOnRemoteSessionStateChange()
}

// FuncHasOnRemoteSessionStateChange returns the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.hasListener".
func FuncHasOnRemoteSessionStateChange() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnRemoteSessionStateChange(
		js.Pointer(&fn),
	)
	return
}

// HasOnRemoteSessionStateChange calls the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.hasListener" directly.
func HasOnRemoteSessionStateChange(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnRemoteSessionStateChange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRemoteSessionStateChange calls the function "WEBEXT.webAuthenticationProxy.onRemoteSessionStateChange.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRemoteSessionStateChange(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRemoteSessionStateChange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRequestCanceledEventCallbackFunc func(this js.Ref, requestId int32) js.Ref

func (fn OnRequestCanceledEventCallbackFunc) Register() js.Func[func(requestId int32)] {
	return js.RegisterCallback[func(requestId int32)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRequestCanceledEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnRequestCanceledEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, requestId int32) js.Ref
	Arg T
}

func (cb *OnRequestCanceledEventCallback[T]) Register() js.Func[func(requestId int32)] {
	return js.RegisterCallback[func(requestId int32)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRequestCanceledEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[int32]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnRequestCanceled returns true if the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.addListener" exists.
func HasFuncOnRequestCanceled() bool {
	return js.True == bindings.HasFuncOnRequestCanceled()
}

// FuncOnRequestCanceled returns the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.addListener".
func FuncOnRequestCanceled() (fn js.Func[func(callback js.Func[func(requestId int32)])]) {
	bindings.FuncOnRequestCanceled(
		js.Pointer(&fn),
	)
	return
}

// OnRequestCanceled calls the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.addListener" directly.
func OnRequestCanceled(callback js.Func[func(requestId int32)]) (ret js.Void) {
	bindings.CallOnRequestCanceled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRequestCanceled calls the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRequestCanceled(callback js.Func[func(requestId int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRequestCanceled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRequestCanceled returns true if the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.removeListener" exists.
func HasFuncOffRequestCanceled() bool {
	return js.True == bindings.HasFuncOffRequestCanceled()
}

// FuncOffRequestCanceled returns the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.removeListener".
func FuncOffRequestCanceled() (fn js.Func[func(callback js.Func[func(requestId int32)])]) {
	bindings.FuncOffRequestCanceled(
		js.Pointer(&fn),
	)
	return
}

// OffRequestCanceled calls the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.removeListener" directly.
func OffRequestCanceled(callback js.Func[func(requestId int32)]) (ret js.Void) {
	bindings.CallOffRequestCanceled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRequestCanceled calls the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRequestCanceled(callback js.Func[func(requestId int32)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRequestCanceled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRequestCanceled returns true if the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.hasListener" exists.
func HasFuncHasOnRequestCanceled() bool {
	return js.True == bindings.HasFuncHasOnRequestCanceled()
}

// FuncHasOnRequestCanceled returns the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.hasListener".
func FuncHasOnRequestCanceled() (fn js.Func[func(callback js.Func[func(requestId int32)]) bool]) {
	bindings.FuncHasOnRequestCanceled(
		js.Pointer(&fn),
	)
	return
}

// HasOnRequestCanceled calls the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.hasListener" directly.
func HasOnRequestCanceled(callback js.Func[func(requestId int32)]) (ret bool) {
	bindings.CallHasOnRequestCanceled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRequestCanceled calls the function "WEBEXT.webAuthenticationProxy.onRequestCanceled.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRequestCanceled(callback js.Func[func(requestId int32)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRequestCanceled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

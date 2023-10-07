// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webrequestinternal

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/webrequest"
	"github.com/primecitizens/pcz/std/plat/js/webext/webrequestinternal/bindings"
)

type AddEventListenerArgCallbackFunc func(this js.Ref) js.Ref

func (fn AddEventListenerArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AddEventListenerArgCallbackFunc) DispatchCallback(
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

type AddEventListenerArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *AddEventListenerArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AddEventListenerArgCallback[T]) DispatchCallback(
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

type AddEventListenerOptions uint32

const (
	_ AddEventListenerOptions = iota

	AddEventListenerOptions_REQUEST_HEADERS
	AddEventListenerOptions_RESPONSE_HEADERS
	AddEventListenerOptions_BLOCKING
	AddEventListenerOptions_ASYNC_BLOCKING
	AddEventListenerOptions_REQUEST_BODY
	AddEventListenerOptions_EXTRA_HEADERS
)

func (AddEventListenerOptions) FromRef(str js.Ref) AddEventListenerOptions {
	return AddEventListenerOptions(bindings.ConstOfAddEventListenerOptions(str))
}

func (x AddEventListenerOptions) String() (string, bool) {
	switch x {
	case AddEventListenerOptions_REQUEST_HEADERS:
		return "requestHeaders", true
	case AddEventListenerOptions_RESPONSE_HEADERS:
		return "responseHeaders", true
	case AddEventListenerOptions_BLOCKING:
		return "blocking", true
	case AddEventListenerOptions_ASYNC_BLOCKING:
		return "asyncBlocking", true
	case AddEventListenerOptions_REQUEST_BODY:
		return "requestBody", true
	case AddEventListenerOptions_EXTRA_HEADERS:
		return "extraHeaders", true
	default:
		return "", false
	}
}

// HasFuncAddEventListener returns true if the function "WEBEXT.webRequestInternal.addEventListener" exists.
func HasFuncAddEventListener() bool {
	return js.True == bindings.HasFuncAddEventListener()
}

// FuncAddEventListener returns the function "WEBEXT.webRequestInternal.addEventListener".
func FuncAddEventListener() (fn js.Func[func(callback js.Func[func()], filter webrequest.RequestFilter, extraInfoSpec js.Array[AddEventListenerOptions], eventName js.String, subEventName js.String, webViewInstanceId int64)]) {
	bindings.FuncAddEventListener(
		js.Pointer(&fn),
	)
	return
}

// AddEventListener calls the function "WEBEXT.webRequestInternal.addEventListener" directly.
func AddEventListener(callback js.Func[func()], filter webrequest.RequestFilter, extraInfoSpec js.Array[AddEventListenerOptions], eventName js.String, subEventName js.String, webViewInstanceId int64) (ret js.Void) {
	bindings.CallAddEventListener(
		js.Pointer(&ret),
		callback.Ref(),
		js.Pointer(&filter),
		extraInfoSpec.Ref(),
		eventName.Ref(),
		subEventName.Ref(),
		float64(webViewInstanceId),
	)

	return
}

// TryAddEventListener calls the function "WEBEXT.webRequestInternal.addEventListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddEventListener(callback js.Func[func()], filter webrequest.RequestFilter, extraInfoSpec js.Array[AddEventListenerOptions], eventName js.String, subEventName js.String, webViewInstanceId int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddEventListener(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
		js.Pointer(&filter),
		extraInfoSpec.Ref(),
		eventName.Ref(),
		subEventName.Ref(),
		float64(webViewInstanceId),
	)

	return
}

// HasFuncEventHandled returns true if the function "WEBEXT.webRequestInternal.eventHandled" exists.
func HasFuncEventHandled() bool {
	return js.True == bindings.HasFuncEventHandled()
}

// FuncEventHandled returns the function "WEBEXT.webRequestInternal.eventHandled".
func FuncEventHandled() (fn js.Func[func(eventName js.String, subEventName js.String, requestId js.String, webViewInstanceId int64, response webrequest.BlockingResponse)]) {
	bindings.FuncEventHandled(
		js.Pointer(&fn),
	)
	return
}

// EventHandled calls the function "WEBEXT.webRequestInternal.eventHandled" directly.
func EventHandled(eventName js.String, subEventName js.String, requestId js.String, webViewInstanceId int64, response webrequest.BlockingResponse) (ret js.Void) {
	bindings.CallEventHandled(
		js.Pointer(&ret),
		eventName.Ref(),
		subEventName.Ref(),
		requestId.Ref(),
		float64(webViewInstanceId),
		js.Pointer(&response),
	)

	return
}

// TryEventHandled calls the function "WEBEXT.webRequestInternal.eventHandled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryEventHandled(eventName js.String, subEventName js.String, requestId js.String, webViewInstanceId int64, response webrequest.BlockingResponse) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventHandled(
		js.Pointer(&ret), js.Pointer(&exception),
		eventName.Ref(),
		subEventName.Ref(),
		requestId.Ref(),
		float64(webViewInstanceId),
		js.Pointer(&response),
	)

	return
}

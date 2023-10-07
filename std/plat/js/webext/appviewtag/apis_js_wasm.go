// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package appviewtag

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/appviewtag/bindings"
)

type EmbedRequest struct {
	ref js.Ref
}

func (this EmbedRequest) Once() EmbedRequest {
	this.ref.Once()
	return this
}

func (this EmbedRequest) Ref() js.Ref {
	return this.ref
}

func (this EmbedRequest) FromRef(ref js.Ref) EmbedRequest {
	this.ref = ref
	return this
}

func (this EmbedRequest) Free() {
	this.ref.Free()
}

// EmbedderId returns the value of property "EmbedRequest.embedderId".
//
// It returns ok=false if there is no such property.
func (this EmbedRequest) EmbedderId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetEmbedRequestEmbedderId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEmbedderId sets the value of property "EmbedRequest.embedderId" to val.
//
// It returns false if the property cannot be set.
func (this EmbedRequest) SetEmbedderId(val js.String) bool {
	return js.True == bindings.SetEmbedRequestEmbedderId(
		this.ref,
		val.Ref(),
	)
}

// Data returns the value of property "EmbedRequest.data".
//
// It returns ok=false if there is no such property.
func (this EmbedRequest) Data() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetEmbedRequestData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetData sets the value of property "EmbedRequest.data" to val.
//
// It returns false if the property cannot be set.
func (this EmbedRequest) SetData(val js.Object) bool {
	return js.True == bindings.SetEmbedRequestData(
		this.ref,
		val.Ref(),
	)
}

// HasFuncAllow returns true if the method "EmbedRequest.allow" exists.
func (this EmbedRequest) HasFuncAllow() bool {
	return js.True == bindings.HasFuncEmbedRequestAllow(
		this.ref,
	)
}

// FuncAllow returns the method "EmbedRequest.allow".
func (this EmbedRequest) FuncAllow() (fn js.Func[func(url js.String)]) {
	bindings.FuncEmbedRequestAllow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Allow calls the method "EmbedRequest.allow".
func (this EmbedRequest) Allow(url js.String) (ret js.Void) {
	bindings.CallEmbedRequestAllow(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryAllow calls the method "EmbedRequest.allow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EmbedRequest) TryAllow(url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEmbedRequestAllow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncDeny returns true if the method "EmbedRequest.deny" exists.
func (this EmbedRequest) HasFuncDeny() bool {
	return js.True == bindings.HasFuncEmbedRequestDeny(
		this.ref,
	)
}

// FuncDeny returns the method "EmbedRequest.deny".
func (this EmbedRequest) FuncDeny() (fn js.Func[func()]) {
	bindings.FuncEmbedRequestDeny(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Deny calls the method "EmbedRequest.deny".
func (this EmbedRequest) Deny() (ret js.Void) {
	bindings.CallEmbedRequestDeny(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeny calls the method "EmbedRequest.deny"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EmbedRequest) TryDeny() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEmbedRequestDeny(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type EmbeddingCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn EmbeddingCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EmbeddingCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EmbeddingCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *EmbeddingCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EmbeddingCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncConnect returns true if the function "WEBEXT.appviewTag.connect" exists.
func HasFuncConnect() bool {
	return js.True == bindings.HasFuncConnect()
}

// FuncConnect returns the function "WEBEXT.appviewTag.connect".
func FuncConnect() (fn js.Func[func(app js.String, data js.Any, callback js.Func[func(success bool)])]) {
	bindings.FuncConnect(
		js.Pointer(&fn),
	)
	return
}

// Connect calls the function "WEBEXT.appviewTag.connect" directly.
func Connect(app js.String, data js.Any, callback js.Func[func(success bool)]) (ret js.Void) {
	bindings.CallConnect(
		js.Pointer(&ret),
		app.Ref(),
		data.Ref(),
		callback.Ref(),
	)

	return
}

// TryConnect calls the function "WEBEXT.appviewTag.connect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryConnect(app js.String, data js.Any, callback js.Func[func(success bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryConnect(
		js.Pointer(&ret), js.Pointer(&exception),
		app.Ref(),
		data.Ref(),
		callback.Ref(),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package loginscreenstorage

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/loginscreenstorage/bindings"
)

type RetrieveCallbackFunc func(this js.Ref, data js.String) js.Ref

func (fn RetrieveCallbackFunc) Register() js.Func[func(data js.String)] {
	return js.RegisterCallback[func(data js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RetrieveCallbackFunc) DispatchCallback(
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

type RetrieveCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.String) js.Ref
	Arg T
}

func (cb *RetrieveCallback[T]) Register() js.Func[func(data js.String)] {
	return js.RegisterCallback[func(data js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RetrieveCallback[T]) DispatchCallback(
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

type StoreCallbackFunc func(this js.Ref) js.Ref

func (fn StoreCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn StoreCallbackFunc) DispatchCallback(
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

type StoreCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *StoreCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *StoreCallback[T]) DispatchCallback(
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

// HasFuncRetrieveCredentials returns true if the function "WEBEXT.loginScreenStorage.retrieveCredentials" exists.
func HasFuncRetrieveCredentials() bool {
	return js.True == bindings.HasFuncRetrieveCredentials()
}

// FuncRetrieveCredentials returns the function "WEBEXT.loginScreenStorage.retrieveCredentials".
func FuncRetrieveCredentials() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncRetrieveCredentials(
		js.Pointer(&fn),
	)
	return
}

// RetrieveCredentials calls the function "WEBEXT.loginScreenStorage.retrieveCredentials" directly.
func RetrieveCredentials() (ret js.Promise[js.String]) {
	bindings.CallRetrieveCredentials(
		js.Pointer(&ret),
	)

	return
}

// TryRetrieveCredentials calls the function "WEBEXT.loginScreenStorage.retrieveCredentials"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRetrieveCredentials() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRetrieveCredentials(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRetrievePersistentData returns true if the function "WEBEXT.loginScreenStorage.retrievePersistentData" exists.
func HasFuncRetrievePersistentData() bool {
	return js.True == bindings.HasFuncRetrievePersistentData()
}

// FuncRetrievePersistentData returns the function "WEBEXT.loginScreenStorage.retrievePersistentData".
func FuncRetrievePersistentData() (fn js.Func[func(ownerId js.String) js.Promise[js.String]]) {
	bindings.FuncRetrievePersistentData(
		js.Pointer(&fn),
	)
	return
}

// RetrievePersistentData calls the function "WEBEXT.loginScreenStorage.retrievePersistentData" directly.
func RetrievePersistentData(ownerId js.String) (ret js.Promise[js.String]) {
	bindings.CallRetrievePersistentData(
		js.Pointer(&ret),
		ownerId.Ref(),
	)

	return
}

// TryRetrievePersistentData calls the function "WEBEXT.loginScreenStorage.retrievePersistentData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRetrievePersistentData(ownerId js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRetrievePersistentData(
		js.Pointer(&ret), js.Pointer(&exception),
		ownerId.Ref(),
	)

	return
}

// HasFuncStoreCredentials returns true if the function "WEBEXT.loginScreenStorage.storeCredentials" exists.
func HasFuncStoreCredentials() bool {
	return js.True == bindings.HasFuncStoreCredentials()
}

// FuncStoreCredentials returns the function "WEBEXT.loginScreenStorage.storeCredentials".
func FuncStoreCredentials() (fn js.Func[func(extensionId js.String, credentials js.String) js.Promise[js.Void]]) {
	bindings.FuncStoreCredentials(
		js.Pointer(&fn),
	)
	return
}

// StoreCredentials calls the function "WEBEXT.loginScreenStorage.storeCredentials" directly.
func StoreCredentials(extensionId js.String, credentials js.String) (ret js.Promise[js.Void]) {
	bindings.CallStoreCredentials(
		js.Pointer(&ret),
		extensionId.Ref(),
		credentials.Ref(),
	)

	return
}

// TryStoreCredentials calls the function "WEBEXT.loginScreenStorage.storeCredentials"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStoreCredentials(extensionId js.String, credentials js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStoreCredentials(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
		credentials.Ref(),
	)

	return
}

// HasFuncStorePersistentData returns true if the function "WEBEXT.loginScreenStorage.storePersistentData" exists.
func HasFuncStorePersistentData() bool {
	return js.True == bindings.HasFuncStorePersistentData()
}

// FuncStorePersistentData returns the function "WEBEXT.loginScreenStorage.storePersistentData".
func FuncStorePersistentData() (fn js.Func[func(extensionIds js.Array[js.String], data js.String) js.Promise[js.Void]]) {
	bindings.FuncStorePersistentData(
		js.Pointer(&fn),
	)
	return
}

// StorePersistentData calls the function "WEBEXT.loginScreenStorage.storePersistentData" directly.
func StorePersistentData(extensionIds js.Array[js.String], data js.String) (ret js.Promise[js.Void]) {
	bindings.CallStorePersistentData(
		js.Pointer(&ret),
		extensionIds.Ref(),
		data.Ref(),
	)

	return
}

// TryStorePersistentData calls the function "WEBEXT.loginScreenStorage.storePersistentData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStorePersistentData(extensionIds js.Array[js.String], data js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorePersistentData(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionIds.Ref(),
		data.Ref(),
	)

	return
}

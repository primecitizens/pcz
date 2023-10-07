// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package deviceattributes

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/enterprise/deviceattributes/bindings"
)

type GetDeviceAnnotatedLocationCallbackFunc func(this js.Ref, annotatedLocation js.String) js.Ref

func (fn GetDeviceAnnotatedLocationCallbackFunc) Register() js.Func[func(annotatedLocation js.String)] {
	return js.RegisterCallback[func(annotatedLocation js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeviceAnnotatedLocationCallbackFunc) DispatchCallback(
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

type GetDeviceAnnotatedLocationCallback[T any] struct {
	Fn  func(arg T, this js.Ref, annotatedLocation js.String) js.Ref
	Arg T
}

func (cb *GetDeviceAnnotatedLocationCallback[T]) Register() js.Func[func(annotatedLocation js.String)] {
	return js.RegisterCallback[func(annotatedLocation js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeviceAnnotatedLocationCallback[T]) DispatchCallback(
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

type GetDeviceAssetIdCallbackFunc func(this js.Ref, assetId js.String) js.Ref

func (fn GetDeviceAssetIdCallbackFunc) Register() js.Func[func(assetId js.String)] {
	return js.RegisterCallback[func(assetId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeviceAssetIdCallbackFunc) DispatchCallback(
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

type GetDeviceAssetIdCallback[T any] struct {
	Fn  func(arg T, this js.Ref, assetId js.String) js.Ref
	Arg T
}

func (cb *GetDeviceAssetIdCallback[T]) Register() js.Func[func(assetId js.String)] {
	return js.RegisterCallback[func(assetId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeviceAssetIdCallback[T]) DispatchCallback(
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

type GetDeviceHostnameCallbackFunc func(this js.Ref, hostname js.String) js.Ref

func (fn GetDeviceHostnameCallbackFunc) Register() js.Func[func(hostname js.String)] {
	return js.RegisterCallback[func(hostname js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeviceHostnameCallbackFunc) DispatchCallback(
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

type GetDeviceHostnameCallback[T any] struct {
	Fn  func(arg T, this js.Ref, hostname js.String) js.Ref
	Arg T
}

func (cb *GetDeviceHostnameCallback[T]) Register() js.Func[func(hostname js.String)] {
	return js.RegisterCallback[func(hostname js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeviceHostnameCallback[T]) DispatchCallback(
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

type GetDeviceSerialNumberCallbackFunc func(this js.Ref, serialNumber js.String) js.Ref

func (fn GetDeviceSerialNumberCallbackFunc) Register() js.Func[func(serialNumber js.String)] {
	return js.RegisterCallback[func(serialNumber js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDeviceSerialNumberCallbackFunc) DispatchCallback(
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

type GetDeviceSerialNumberCallback[T any] struct {
	Fn  func(arg T, this js.Ref, serialNumber js.String) js.Ref
	Arg T
}

func (cb *GetDeviceSerialNumberCallback[T]) Register() js.Func[func(serialNumber js.String)] {
	return js.RegisterCallback[func(serialNumber js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDeviceSerialNumberCallback[T]) DispatchCallback(
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

type GetDirectoryDeviceIdCallbackFunc func(this js.Ref, deviceId js.String) js.Ref

func (fn GetDirectoryDeviceIdCallbackFunc) Register() js.Func[func(deviceId js.String)] {
	return js.RegisterCallback[func(deviceId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetDirectoryDeviceIdCallbackFunc) DispatchCallback(
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

type GetDirectoryDeviceIdCallback[T any] struct {
	Fn  func(arg T, this js.Ref, deviceId js.String) js.Ref
	Arg T
}

func (cb *GetDirectoryDeviceIdCallback[T]) Register() js.Func[func(deviceId js.String)] {
	return js.RegisterCallback[func(deviceId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetDirectoryDeviceIdCallback[T]) DispatchCallback(
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

// HasFuncGetDeviceAnnotatedLocation returns true if the function "WEBEXT.enterprise.deviceAttributes.getDeviceAnnotatedLocation" exists.
func HasFuncGetDeviceAnnotatedLocation() bool {
	return js.True == bindings.HasFuncGetDeviceAnnotatedLocation()
}

// FuncGetDeviceAnnotatedLocation returns the function "WEBEXT.enterprise.deviceAttributes.getDeviceAnnotatedLocation".
func FuncGetDeviceAnnotatedLocation() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetDeviceAnnotatedLocation(
		js.Pointer(&fn),
	)
	return
}

// GetDeviceAnnotatedLocation calls the function "WEBEXT.enterprise.deviceAttributes.getDeviceAnnotatedLocation" directly.
func GetDeviceAnnotatedLocation() (ret js.Promise[js.String]) {
	bindings.CallGetDeviceAnnotatedLocation(
		js.Pointer(&ret),
	)

	return
}

// TryGetDeviceAnnotatedLocation calls the function "WEBEXT.enterprise.deviceAttributes.getDeviceAnnotatedLocation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeviceAnnotatedLocation() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeviceAnnotatedLocation(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDeviceAssetId returns true if the function "WEBEXT.enterprise.deviceAttributes.getDeviceAssetId" exists.
func HasFuncGetDeviceAssetId() bool {
	return js.True == bindings.HasFuncGetDeviceAssetId()
}

// FuncGetDeviceAssetId returns the function "WEBEXT.enterprise.deviceAttributes.getDeviceAssetId".
func FuncGetDeviceAssetId() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetDeviceAssetId(
		js.Pointer(&fn),
	)
	return
}

// GetDeviceAssetId calls the function "WEBEXT.enterprise.deviceAttributes.getDeviceAssetId" directly.
func GetDeviceAssetId() (ret js.Promise[js.String]) {
	bindings.CallGetDeviceAssetId(
		js.Pointer(&ret),
	)

	return
}

// TryGetDeviceAssetId calls the function "WEBEXT.enterprise.deviceAttributes.getDeviceAssetId"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeviceAssetId() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeviceAssetId(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDeviceHostname returns true if the function "WEBEXT.enterprise.deviceAttributes.getDeviceHostname" exists.
func HasFuncGetDeviceHostname() bool {
	return js.True == bindings.HasFuncGetDeviceHostname()
}

// FuncGetDeviceHostname returns the function "WEBEXT.enterprise.deviceAttributes.getDeviceHostname".
func FuncGetDeviceHostname() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetDeviceHostname(
		js.Pointer(&fn),
	)
	return
}

// GetDeviceHostname calls the function "WEBEXT.enterprise.deviceAttributes.getDeviceHostname" directly.
func GetDeviceHostname() (ret js.Promise[js.String]) {
	bindings.CallGetDeviceHostname(
		js.Pointer(&ret),
	)

	return
}

// TryGetDeviceHostname calls the function "WEBEXT.enterprise.deviceAttributes.getDeviceHostname"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeviceHostname() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeviceHostname(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDeviceSerialNumber returns true if the function "WEBEXT.enterprise.deviceAttributes.getDeviceSerialNumber" exists.
func HasFuncGetDeviceSerialNumber() bool {
	return js.True == bindings.HasFuncGetDeviceSerialNumber()
}

// FuncGetDeviceSerialNumber returns the function "WEBEXT.enterprise.deviceAttributes.getDeviceSerialNumber".
func FuncGetDeviceSerialNumber() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetDeviceSerialNumber(
		js.Pointer(&fn),
	)
	return
}

// GetDeviceSerialNumber calls the function "WEBEXT.enterprise.deviceAttributes.getDeviceSerialNumber" directly.
func GetDeviceSerialNumber() (ret js.Promise[js.String]) {
	bindings.CallGetDeviceSerialNumber(
		js.Pointer(&ret),
	)

	return
}

// TryGetDeviceSerialNumber calls the function "WEBEXT.enterprise.deviceAttributes.getDeviceSerialNumber"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDeviceSerialNumber() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDeviceSerialNumber(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetDirectoryDeviceId returns true if the function "WEBEXT.enterprise.deviceAttributes.getDirectoryDeviceId" exists.
func HasFuncGetDirectoryDeviceId() bool {
	return js.True == bindings.HasFuncGetDirectoryDeviceId()
}

// FuncGetDirectoryDeviceId returns the function "WEBEXT.enterprise.deviceAttributes.getDirectoryDeviceId".
func FuncGetDirectoryDeviceId() (fn js.Func[func() js.Promise[js.String]]) {
	bindings.FuncGetDirectoryDeviceId(
		js.Pointer(&fn),
	)
	return
}

// GetDirectoryDeviceId calls the function "WEBEXT.enterprise.deviceAttributes.getDirectoryDeviceId" directly.
func GetDirectoryDeviceId() (ret js.Promise[js.String]) {
	bindings.CallGetDirectoryDeviceId(
		js.Pointer(&ret),
	)

	return
}

// TryGetDirectoryDeviceId calls the function "WEBEXT.enterprise.deviceAttributes.getDirectoryDeviceId"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDirectoryDeviceId() (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDirectoryDeviceId(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

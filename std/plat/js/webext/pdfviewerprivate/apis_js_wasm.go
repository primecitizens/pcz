// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package pdfviewerprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/pdfviewerprivate/bindings"
)

type IsAllowedLocalFileAccessCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn IsAllowedLocalFileAccessCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsAllowedLocalFileAccessCallbackFunc) DispatchCallback(
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

type IsAllowedLocalFileAccessCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *IsAllowedLocalFileAccessCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsAllowedLocalFileAccessCallback[T]) DispatchCallback(
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

type IsPdfOcrAlwaysActiveCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn IsPdfOcrAlwaysActiveCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsPdfOcrAlwaysActiveCallbackFunc) DispatchCallback(
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

type IsPdfOcrAlwaysActiveCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *IsPdfOcrAlwaysActiveCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsPdfOcrAlwaysActiveCallback[T]) DispatchCallback(
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

type OnPdfOcrPrefSetCallbackFunc func(this js.Ref, result bool) js.Ref

func (fn OnPdfOcrPrefSetCallbackFunc) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPdfOcrPrefSetCallbackFunc) DispatchCallback(
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

type OnPdfOcrPrefSetCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result bool) js.Ref
	Arg T
}

func (cb *OnPdfOcrPrefSetCallback[T]) Register() js.Func[func(result bool)] {
	return js.RegisterCallback[func(result bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPdfOcrPrefSetCallback[T]) DispatchCallback(
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

// HasFuncIsAllowedLocalFileAccess returns true if the function "WEBEXT.pdfViewerPrivate.isAllowedLocalFileAccess" exists.
func HasFuncIsAllowedLocalFileAccess() bool {
	return js.True == bindings.HasFuncIsAllowedLocalFileAccess()
}

// FuncIsAllowedLocalFileAccess returns the function "WEBEXT.pdfViewerPrivate.isAllowedLocalFileAccess".
func FuncIsAllowedLocalFileAccess() (fn js.Func[func(url js.String) js.Promise[js.Boolean]]) {
	bindings.FuncIsAllowedLocalFileAccess(
		js.Pointer(&fn),
	)
	return
}

// IsAllowedLocalFileAccess calls the function "WEBEXT.pdfViewerPrivate.isAllowedLocalFileAccess" directly.
func IsAllowedLocalFileAccess(url js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallIsAllowedLocalFileAccess(
		js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryIsAllowedLocalFileAccess calls the function "WEBEXT.pdfViewerPrivate.isAllowedLocalFileAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsAllowedLocalFileAccess(url js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsAllowedLocalFileAccess(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncIsPdfOcrAlwaysActive returns true if the function "WEBEXT.pdfViewerPrivate.isPdfOcrAlwaysActive" exists.
func HasFuncIsPdfOcrAlwaysActive() bool {
	return js.True == bindings.HasFuncIsPdfOcrAlwaysActive()
}

// FuncIsPdfOcrAlwaysActive returns the function "WEBEXT.pdfViewerPrivate.isPdfOcrAlwaysActive".
func FuncIsPdfOcrAlwaysActive() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsPdfOcrAlwaysActive(
		js.Pointer(&fn),
	)
	return
}

// IsPdfOcrAlwaysActive calls the function "WEBEXT.pdfViewerPrivate.isPdfOcrAlwaysActive" directly.
func IsPdfOcrAlwaysActive() (ret js.Promise[js.Boolean]) {
	bindings.CallIsPdfOcrAlwaysActive(
		js.Pointer(&ret),
	)

	return
}

// TryIsPdfOcrAlwaysActive calls the function "WEBEXT.pdfViewerPrivate.isPdfOcrAlwaysActive"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsPdfOcrAlwaysActive() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsPdfOcrAlwaysActive(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnPdfOcrPrefChangedEventCallbackFunc func(this js.Ref, value bool) js.Ref

func (fn OnPdfOcrPrefChangedEventCallbackFunc) Register() js.Func[func(value bool)] {
	return js.RegisterCallback[func(value bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPdfOcrPrefChangedEventCallbackFunc) DispatchCallback(
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

type OnPdfOcrPrefChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, value bool) js.Ref
	Arg T
}

func (cb *OnPdfOcrPrefChangedEventCallback[T]) Register() js.Func[func(value bool)] {
	return js.RegisterCallback[func(value bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPdfOcrPrefChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnPdfOcrPrefChanged returns true if the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.addListener" exists.
func HasFuncOnPdfOcrPrefChanged() bool {
	return js.True == bindings.HasFuncOnPdfOcrPrefChanged()
}

// FuncOnPdfOcrPrefChanged returns the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.addListener".
func FuncOnPdfOcrPrefChanged() (fn js.Func[func(callback js.Func[func(value bool)])]) {
	bindings.FuncOnPdfOcrPrefChanged(
		js.Pointer(&fn),
	)
	return
}

// OnPdfOcrPrefChanged calls the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.addListener" directly.
func OnPdfOcrPrefChanged(callback js.Func[func(value bool)]) (ret js.Void) {
	bindings.CallOnPdfOcrPrefChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPdfOcrPrefChanged calls the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPdfOcrPrefChanged(callback js.Func[func(value bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPdfOcrPrefChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPdfOcrPrefChanged returns true if the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.removeListener" exists.
func HasFuncOffPdfOcrPrefChanged() bool {
	return js.True == bindings.HasFuncOffPdfOcrPrefChanged()
}

// FuncOffPdfOcrPrefChanged returns the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.removeListener".
func FuncOffPdfOcrPrefChanged() (fn js.Func[func(callback js.Func[func(value bool)])]) {
	bindings.FuncOffPdfOcrPrefChanged(
		js.Pointer(&fn),
	)
	return
}

// OffPdfOcrPrefChanged calls the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.removeListener" directly.
func OffPdfOcrPrefChanged(callback js.Func[func(value bool)]) (ret js.Void) {
	bindings.CallOffPdfOcrPrefChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPdfOcrPrefChanged calls the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPdfOcrPrefChanged(callback js.Func[func(value bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPdfOcrPrefChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPdfOcrPrefChanged returns true if the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.hasListener" exists.
func HasFuncHasOnPdfOcrPrefChanged() bool {
	return js.True == bindings.HasFuncHasOnPdfOcrPrefChanged()
}

// FuncHasOnPdfOcrPrefChanged returns the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.hasListener".
func FuncHasOnPdfOcrPrefChanged() (fn js.Func[func(callback js.Func[func(value bool)]) bool]) {
	bindings.FuncHasOnPdfOcrPrefChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnPdfOcrPrefChanged calls the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.hasListener" directly.
func HasOnPdfOcrPrefChanged(callback js.Func[func(value bool)]) (ret bool) {
	bindings.CallHasOnPdfOcrPrefChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPdfOcrPrefChanged calls the function "WEBEXT.pdfViewerPrivate.onPdfOcrPrefChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPdfOcrPrefChanged(callback js.Func[func(value bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPdfOcrPrefChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetPdfOcrPref returns true if the function "WEBEXT.pdfViewerPrivate.setPdfOcrPref" exists.
func HasFuncSetPdfOcrPref() bool {
	return js.True == bindings.HasFuncSetPdfOcrPref()
}

// FuncSetPdfOcrPref returns the function "WEBEXT.pdfViewerPrivate.setPdfOcrPref".
func FuncSetPdfOcrPref() (fn js.Func[func(value bool) js.Promise[js.Boolean]]) {
	bindings.FuncSetPdfOcrPref(
		js.Pointer(&fn),
	)
	return
}

// SetPdfOcrPref calls the function "WEBEXT.pdfViewerPrivate.setPdfOcrPref" directly.
func SetPdfOcrPref(value bool) (ret js.Promise[js.Boolean]) {
	bindings.CallSetPdfOcrPref(
		js.Pointer(&ret),
		js.Bool(bool(value)),
	)

	return
}

// TrySetPdfOcrPref calls the function "WEBEXT.pdfViewerPrivate.setPdfOcrPref"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPdfOcrPref(value bool) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPdfOcrPref(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(value)),
	)

	return
}

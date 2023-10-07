// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package imageloaderprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/imageloaderprivate/bindings"
)

type GetThumbnailCallbackFunc func(this js.Ref, thumbnailDataUrl js.String) js.Ref

func (fn GetThumbnailCallbackFunc) Register() js.Func[func(thumbnailDataUrl js.String)] {
	return js.RegisterCallback[func(thumbnailDataUrl js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetThumbnailCallbackFunc) DispatchCallback(
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

type GetThumbnailCallback[T any] struct {
	Fn  func(arg T, this js.Ref, thumbnailDataUrl js.String) js.Ref
	Arg T
}

func (cb *GetThumbnailCallback[T]) Register() js.Func[func(thumbnailDataUrl js.String)] {
	return js.RegisterCallback[func(thumbnailDataUrl js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetThumbnailCallback[T]) DispatchCallback(
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

// HasFuncGetArcDocumentsProviderThumbnail returns true if the function "WEBEXT.imageLoaderPrivate.getArcDocumentsProviderThumbnail" exists.
func HasFuncGetArcDocumentsProviderThumbnail() bool {
	return js.True == bindings.HasFuncGetArcDocumentsProviderThumbnail()
}

// FuncGetArcDocumentsProviderThumbnail returns the function "WEBEXT.imageLoaderPrivate.getArcDocumentsProviderThumbnail".
func FuncGetArcDocumentsProviderThumbnail() (fn js.Func[func(url js.String, widthHint int32, heightHint int32, callback js.Func[func(thumbnailDataUrl js.String)])]) {
	bindings.FuncGetArcDocumentsProviderThumbnail(
		js.Pointer(&fn),
	)
	return
}

// GetArcDocumentsProviderThumbnail calls the function "WEBEXT.imageLoaderPrivate.getArcDocumentsProviderThumbnail" directly.
func GetArcDocumentsProviderThumbnail(url js.String, widthHint int32, heightHint int32, callback js.Func[func(thumbnailDataUrl js.String)]) (ret js.Void) {
	bindings.CallGetArcDocumentsProviderThumbnail(
		js.Pointer(&ret),
		url.Ref(),
		int32(widthHint),
		int32(heightHint),
		callback.Ref(),
	)

	return
}

// TryGetArcDocumentsProviderThumbnail calls the function "WEBEXT.imageLoaderPrivate.getArcDocumentsProviderThumbnail"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetArcDocumentsProviderThumbnail(url js.String, widthHint int32, heightHint int32, callback js.Func[func(thumbnailDataUrl js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetArcDocumentsProviderThumbnail(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		int32(widthHint),
		int32(heightHint),
		callback.Ref(),
	)

	return
}

// HasFuncGetDriveThumbnail returns true if the function "WEBEXT.imageLoaderPrivate.getDriveThumbnail" exists.
func HasFuncGetDriveThumbnail() bool {
	return js.True == bindings.HasFuncGetDriveThumbnail()
}

// FuncGetDriveThumbnail returns the function "WEBEXT.imageLoaderPrivate.getDriveThumbnail".
func FuncGetDriveThumbnail() (fn js.Func[func(url js.String, cropToSquare bool, callback js.Func[func(thumbnailDataUrl js.String)])]) {
	bindings.FuncGetDriveThumbnail(
		js.Pointer(&fn),
	)
	return
}

// GetDriveThumbnail calls the function "WEBEXT.imageLoaderPrivate.getDriveThumbnail" directly.
func GetDriveThumbnail(url js.String, cropToSquare bool, callback js.Func[func(thumbnailDataUrl js.String)]) (ret js.Void) {
	bindings.CallGetDriveThumbnail(
		js.Pointer(&ret),
		url.Ref(),
		js.Bool(bool(cropToSquare)),
		callback.Ref(),
	)

	return
}

// TryGetDriveThumbnail calls the function "WEBEXT.imageLoaderPrivate.getDriveThumbnail"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetDriveThumbnail(url js.String, cropToSquare bool, callback js.Func[func(thumbnailDataUrl js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetDriveThumbnail(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		js.Bool(bool(cropToSquare)),
		callback.Ref(),
	)

	return
}

// HasFuncGetPdfThumbnail returns true if the function "WEBEXT.imageLoaderPrivate.getPdfThumbnail" exists.
func HasFuncGetPdfThumbnail() bool {
	return js.True == bindings.HasFuncGetPdfThumbnail()
}

// FuncGetPdfThumbnail returns the function "WEBEXT.imageLoaderPrivate.getPdfThumbnail".
func FuncGetPdfThumbnail() (fn js.Func[func(url js.String, width int32, height int32, callback js.Func[func(thumbnailDataUrl js.String)])]) {
	bindings.FuncGetPdfThumbnail(
		js.Pointer(&fn),
	)
	return
}

// GetPdfThumbnail calls the function "WEBEXT.imageLoaderPrivate.getPdfThumbnail" directly.
func GetPdfThumbnail(url js.String, width int32, height int32, callback js.Func[func(thumbnailDataUrl js.String)]) (ret js.Void) {
	bindings.CallGetPdfThumbnail(
		js.Pointer(&ret),
		url.Ref(),
		int32(width),
		int32(height),
		callback.Ref(),
	)

	return
}

// TryGetPdfThumbnail calls the function "WEBEXT.imageLoaderPrivate.getPdfThumbnail"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetPdfThumbnail(url js.String, width int32, height int32, callback js.Func[func(thumbnailDataUrl js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetPdfThumbnail(
		js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		int32(width),
		int32(height),
		callback.Ref(),
	)

	return
}

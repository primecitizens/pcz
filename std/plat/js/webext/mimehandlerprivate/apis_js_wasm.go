// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package mimehandlerprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/mimehandlerprivate/bindings"
)

type GetStreamDetailsCallbackFunc func(this js.Ref, streamInfo *StreamInfo) js.Ref

func (fn GetStreamDetailsCallbackFunc) Register() js.Func[func(streamInfo *StreamInfo)] {
	return js.RegisterCallback[func(streamInfo *StreamInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetStreamDetailsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StreamInfo
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

type GetStreamDetailsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, streamInfo *StreamInfo) js.Ref
	Arg T
}

func (cb *GetStreamDetailsCallback[T]) Register() js.Func[func(streamInfo *StreamInfo)] {
	return js.RegisterCallback[func(streamInfo *StreamInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetStreamDetailsCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 StreamInfo
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

type StreamInfo struct {
	// MimeType is "StreamInfo.mimeType"
	//
	// Optional
	MimeType js.String
	// OriginalUrl is "StreamInfo.originalUrl"
	//
	// Optional
	OriginalUrl js.String
	// StreamUrl is "StreamInfo.streamUrl"
	//
	// Optional
	StreamUrl js.String
	// TabId is "StreamInfo.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32
	// ResponseHeaders is "StreamInfo.responseHeaders"
	//
	// Optional
	ResponseHeaders js.Object
	// Embedded is "StreamInfo.embedded"
	//
	// Optional
	//
	// NOTE: FFI_USE_Embedded MUST be set to true to make this field effective.
	Embedded bool

	FFI_USE_TabId    bool // for TabId.
	FFI_USE_Embedded bool // for Embedded.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StreamInfo with all fields set.
func (p StreamInfo) FromRef(ref js.Ref) StreamInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StreamInfo in the application heap.
func (p StreamInfo) New() js.Ref {
	return bindings.StreamInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StreamInfo) UpdateFrom(ref js.Ref) {
	bindings.StreamInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StreamInfo) Update(ref js.Ref) {
	bindings.StreamInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StreamInfo) FreeMembers(recursive bool) {
	js.Free(
		p.MimeType.Ref(),
		p.OriginalUrl.Ref(),
		p.StreamUrl.Ref(),
		p.ResponseHeaders.Ref(),
	)
	p.MimeType = p.MimeType.FromRef(js.Undefined)
	p.OriginalUrl = p.OriginalUrl.FromRef(js.Undefined)
	p.StreamUrl = p.StreamUrl.FromRef(js.Undefined)
	p.ResponseHeaders = p.ResponseHeaders.FromRef(js.Undefined)
}

type PdfPluginAttributes struct {
	// BackgroundColor is "PdfPluginAttributes.backgroundColor"
	//
	// Optional
	//
	// NOTE: FFI_USE_BackgroundColor MUST be set to true to make this field effective.
	BackgroundColor float64
	// AllowJavascript is "PdfPluginAttributes.allowJavascript"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllowJavascript MUST be set to true to make this field effective.
	AllowJavascript bool

	FFI_USE_BackgroundColor bool // for BackgroundColor.
	FFI_USE_AllowJavascript bool // for AllowJavascript.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PdfPluginAttributes with all fields set.
func (p PdfPluginAttributes) FromRef(ref js.Ref) PdfPluginAttributes {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PdfPluginAttributes in the application heap.
func (p PdfPluginAttributes) New() js.Ref {
	return bindings.PdfPluginAttributesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PdfPluginAttributes) UpdateFrom(ref js.Ref) {
	bindings.PdfPluginAttributesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PdfPluginAttributes) Update(ref js.Ref) {
	bindings.PdfPluginAttributesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PdfPluginAttributes) FreeMembers(recursive bool) {
}

type SetShowBeforeUnloadDialogCallbackFunc func(this js.Ref) js.Ref

func (fn SetShowBeforeUnloadDialogCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetShowBeforeUnloadDialogCallbackFunc) DispatchCallback(
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

type SetShowBeforeUnloadDialogCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SetShowBeforeUnloadDialogCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetShowBeforeUnloadDialogCallback[T]) DispatchCallback(
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

// HasFuncGetStreamInfo returns true if the function "WEBEXT.mimeHandlerPrivate.getStreamInfo" exists.
func HasFuncGetStreamInfo() bool {
	return js.True == bindings.HasFuncGetStreamInfo()
}

// FuncGetStreamInfo returns the function "WEBEXT.mimeHandlerPrivate.getStreamInfo".
func FuncGetStreamInfo() (fn js.Func[func(callback js.Func[func(streamInfo *StreamInfo)])]) {
	bindings.FuncGetStreamInfo(
		js.Pointer(&fn),
	)
	return
}

// GetStreamInfo calls the function "WEBEXT.mimeHandlerPrivate.getStreamInfo" directly.
func GetStreamInfo(callback js.Func[func(streamInfo *StreamInfo)]) (ret js.Void) {
	bindings.CallGetStreamInfo(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetStreamInfo calls the function "WEBEXT.mimeHandlerPrivate.getStreamInfo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetStreamInfo(callback js.Func[func(streamInfo *StreamInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetStreamInfo(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSaveEventCallbackFunc func(this js.Ref, streamUrl js.String) js.Ref

func (fn OnSaveEventCallbackFunc) Register() js.Func[func(streamUrl js.String)] {
	return js.RegisterCallback[func(streamUrl js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSaveEventCallbackFunc) DispatchCallback(
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

type OnSaveEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, streamUrl js.String) js.Ref
	Arg T
}

func (cb *OnSaveEventCallback[T]) Register() js.Func[func(streamUrl js.String)] {
	return js.RegisterCallback[func(streamUrl js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSaveEventCallback[T]) DispatchCallback(
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

// HasFuncOnSave returns true if the function "WEBEXT.mimeHandlerPrivate.onSave.addListener" exists.
func HasFuncOnSave() bool {
	return js.True == bindings.HasFuncOnSave()
}

// FuncOnSave returns the function "WEBEXT.mimeHandlerPrivate.onSave.addListener".
func FuncOnSave() (fn js.Func[func(callback js.Func[func(streamUrl js.String)])]) {
	bindings.FuncOnSave(
		js.Pointer(&fn),
	)
	return
}

// OnSave calls the function "WEBEXT.mimeHandlerPrivate.onSave.addListener" directly.
func OnSave(callback js.Func[func(streamUrl js.String)]) (ret js.Void) {
	bindings.CallOnSave(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSave calls the function "WEBEXT.mimeHandlerPrivate.onSave.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSave(callback js.Func[func(streamUrl js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSave(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSave returns true if the function "WEBEXT.mimeHandlerPrivate.onSave.removeListener" exists.
func HasFuncOffSave() bool {
	return js.True == bindings.HasFuncOffSave()
}

// FuncOffSave returns the function "WEBEXT.mimeHandlerPrivate.onSave.removeListener".
func FuncOffSave() (fn js.Func[func(callback js.Func[func(streamUrl js.String)])]) {
	bindings.FuncOffSave(
		js.Pointer(&fn),
	)
	return
}

// OffSave calls the function "WEBEXT.mimeHandlerPrivate.onSave.removeListener" directly.
func OffSave(callback js.Func[func(streamUrl js.String)]) (ret js.Void) {
	bindings.CallOffSave(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSave calls the function "WEBEXT.mimeHandlerPrivate.onSave.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSave(callback js.Func[func(streamUrl js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSave(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSave returns true if the function "WEBEXT.mimeHandlerPrivate.onSave.hasListener" exists.
func HasFuncHasOnSave() bool {
	return js.True == bindings.HasFuncHasOnSave()
}

// FuncHasOnSave returns the function "WEBEXT.mimeHandlerPrivate.onSave.hasListener".
func FuncHasOnSave() (fn js.Func[func(callback js.Func[func(streamUrl js.String)]) bool]) {
	bindings.FuncHasOnSave(
		js.Pointer(&fn),
	)
	return
}

// HasOnSave calls the function "WEBEXT.mimeHandlerPrivate.onSave.hasListener" directly.
func HasOnSave(callback js.Func[func(streamUrl js.String)]) (ret bool) {
	bindings.CallHasOnSave(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSave calls the function "WEBEXT.mimeHandlerPrivate.onSave.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSave(callback js.Func[func(streamUrl js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSave(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetPdfPluginAttributes returns true if the function "WEBEXT.mimeHandlerPrivate.setPdfPluginAttributes" exists.
func HasFuncSetPdfPluginAttributes() bool {
	return js.True == bindings.HasFuncSetPdfPluginAttributes()
}

// FuncSetPdfPluginAttributes returns the function "WEBEXT.mimeHandlerPrivate.setPdfPluginAttributes".
func FuncSetPdfPluginAttributes() (fn js.Func[func(pdfPluginAttributes PdfPluginAttributes)]) {
	bindings.FuncSetPdfPluginAttributes(
		js.Pointer(&fn),
	)
	return
}

// SetPdfPluginAttributes calls the function "WEBEXT.mimeHandlerPrivate.setPdfPluginAttributes" directly.
func SetPdfPluginAttributes(pdfPluginAttributes PdfPluginAttributes) (ret js.Void) {
	bindings.CallSetPdfPluginAttributes(
		js.Pointer(&ret),
		js.Pointer(&pdfPluginAttributes),
	)

	return
}

// TrySetPdfPluginAttributes calls the function "WEBEXT.mimeHandlerPrivate.setPdfPluginAttributes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPdfPluginAttributes(pdfPluginAttributes PdfPluginAttributes) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPdfPluginAttributes(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&pdfPluginAttributes),
	)

	return
}

// HasFuncSetShowBeforeUnloadDialog returns true if the function "WEBEXT.mimeHandlerPrivate.setShowBeforeUnloadDialog" exists.
func HasFuncSetShowBeforeUnloadDialog() bool {
	return js.True == bindings.HasFuncSetShowBeforeUnloadDialog()
}

// FuncSetShowBeforeUnloadDialog returns the function "WEBEXT.mimeHandlerPrivate.setShowBeforeUnloadDialog".
func FuncSetShowBeforeUnloadDialog() (fn js.Func[func(showDialog bool, callback js.Func[func()])]) {
	bindings.FuncSetShowBeforeUnloadDialog(
		js.Pointer(&fn),
	)
	return
}

// SetShowBeforeUnloadDialog calls the function "WEBEXT.mimeHandlerPrivate.setShowBeforeUnloadDialog" directly.
func SetShowBeforeUnloadDialog(showDialog bool, callback js.Func[func()]) (ret js.Void) {
	bindings.CallSetShowBeforeUnloadDialog(
		js.Pointer(&ret),
		js.Bool(bool(showDialog)),
		callback.Ref(),
	)

	return
}

// TrySetShowBeforeUnloadDialog calls the function "WEBEXT.mimeHandlerPrivate.setShowBeforeUnloadDialog"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetShowBeforeUnloadDialog(showDialog bool, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetShowBeforeUnloadDialog(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(showDialog)),
		callback.Ref(),
	)

	return
}

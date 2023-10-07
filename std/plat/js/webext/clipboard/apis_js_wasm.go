// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package clipboard

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/clipboard/bindings"
)

type DataItemType uint32

const (
	_ DataItemType = iota

	DataItemType_TEXT_PLAIN
	DataItemType_TEXT_HTML
)

func (DataItemType) FromRef(str js.Ref) DataItemType {
	return DataItemType(bindings.ConstOfDataItemType(str))
}

func (x DataItemType) String() (string, bool) {
	switch x {
	case DataItemType_TEXT_PLAIN:
		return "textPlain", true
	case DataItemType_TEXT_HTML:
		return "textHtml", true
	default:
		return "", false
	}
}

type AdditionalDataItem struct {
	// Type is "AdditionalDataItem.type"
	//
	// Optional
	Type DataItemType
	// Data is "AdditionalDataItem.data"
	//
	// Optional
	Data js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AdditionalDataItem with all fields set.
func (p AdditionalDataItem) FromRef(ref js.Ref) AdditionalDataItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AdditionalDataItem in the application heap.
func (p AdditionalDataItem) New() js.Ref {
	return bindings.AdditionalDataItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AdditionalDataItem) UpdateFrom(ref js.Ref) {
	bindings.AdditionalDataItemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AdditionalDataItem) Update(ref js.Ref) {
	bindings.AdditionalDataItemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AdditionalDataItem) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

type ImageType uint32

const (
	_ ImageType = iota

	ImageType_PNG
	ImageType_JPEG
)

func (ImageType) FromRef(str js.Ref) ImageType {
	return ImageType(bindings.ConstOfImageType(str))
}

func (x ImageType) String() (string, bool) {
	switch x {
	case ImageType_PNG:
		return "png", true
	case ImageType_JPEG:
		return "jpeg", true
	default:
		return "", false
	}
}

type SetImageDataCallbackFunc func(this js.Ref) js.Ref

func (fn SetImageDataCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetImageDataCallbackFunc) DispatchCallback(
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

type SetImageDataCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SetImageDataCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetImageDataCallback[T]) DispatchCallback(
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

type OnClipboardDataChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnClipboardDataChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnClipboardDataChangedEventCallbackFunc) DispatchCallback(
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

type OnClipboardDataChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnClipboardDataChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnClipboardDataChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnClipboardDataChanged returns true if the function "WEBEXT.clipboard.onClipboardDataChanged.addListener" exists.
func HasFuncOnClipboardDataChanged() bool {
	return js.True == bindings.HasFuncOnClipboardDataChanged()
}

// FuncOnClipboardDataChanged returns the function "WEBEXT.clipboard.onClipboardDataChanged.addListener".
func FuncOnClipboardDataChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnClipboardDataChanged(
		js.Pointer(&fn),
	)
	return
}

// OnClipboardDataChanged calls the function "WEBEXT.clipboard.onClipboardDataChanged.addListener" directly.
func OnClipboardDataChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnClipboardDataChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClipboardDataChanged calls the function "WEBEXT.clipboard.onClipboardDataChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClipboardDataChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClipboardDataChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClipboardDataChanged returns true if the function "WEBEXT.clipboard.onClipboardDataChanged.removeListener" exists.
func HasFuncOffClipboardDataChanged() bool {
	return js.True == bindings.HasFuncOffClipboardDataChanged()
}

// FuncOffClipboardDataChanged returns the function "WEBEXT.clipboard.onClipboardDataChanged.removeListener".
func FuncOffClipboardDataChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffClipboardDataChanged(
		js.Pointer(&fn),
	)
	return
}

// OffClipboardDataChanged calls the function "WEBEXT.clipboard.onClipboardDataChanged.removeListener" directly.
func OffClipboardDataChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffClipboardDataChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClipboardDataChanged calls the function "WEBEXT.clipboard.onClipboardDataChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClipboardDataChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClipboardDataChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClipboardDataChanged returns true if the function "WEBEXT.clipboard.onClipboardDataChanged.hasListener" exists.
func HasFuncHasOnClipboardDataChanged() bool {
	return js.True == bindings.HasFuncHasOnClipboardDataChanged()
}

// FuncHasOnClipboardDataChanged returns the function "WEBEXT.clipboard.onClipboardDataChanged.hasListener".
func FuncHasOnClipboardDataChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnClipboardDataChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnClipboardDataChanged calls the function "WEBEXT.clipboard.onClipboardDataChanged.hasListener" directly.
func HasOnClipboardDataChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnClipboardDataChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClipboardDataChanged calls the function "WEBEXT.clipboard.onClipboardDataChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClipboardDataChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClipboardDataChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetImageData returns true if the function "WEBEXT.clipboard.setImageData" exists.
func HasFuncSetImageData() bool {
	return js.True == bindings.HasFuncSetImageData()
}

// FuncSetImageData returns the function "WEBEXT.clipboard.setImageData".
func FuncSetImageData() (fn js.Func[func(imageData js.ArrayBuffer, typ ImageType, additionalItems js.Array[AdditionalDataItem]) js.Promise[js.Void]]) {
	bindings.FuncSetImageData(
		js.Pointer(&fn),
	)
	return
}

// SetImageData calls the function "WEBEXT.clipboard.setImageData" directly.
func SetImageData(imageData js.ArrayBuffer, typ ImageType, additionalItems js.Array[AdditionalDataItem]) (ret js.Promise[js.Void]) {
	bindings.CallSetImageData(
		js.Pointer(&ret),
		imageData.Ref(),
		uint32(typ),
		additionalItems.Ref(),
	)

	return
}

// TrySetImageData calls the function "WEBEXT.clipboard.setImageData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetImageData(imageData js.ArrayBuffer, typ ImageType, additionalItems js.Array[AdditionalDataItem]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetImageData(
		js.Pointer(&ret), js.Pointer(&exception),
		imageData.Ref(),
		uint32(typ),
		additionalItems.Ref(),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webrtcdesktopcaptureprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/webrtcdesktopcaptureprivate/bindings"
)

type DesktopCaptureSourceType uint32

const (
	_ DesktopCaptureSourceType = iota

	DesktopCaptureSourceType_SCREEN
	DesktopCaptureSourceType_WINDOW
	DesktopCaptureSourceType_TAB
)

func (DesktopCaptureSourceType) FromRef(str js.Ref) DesktopCaptureSourceType {
	return DesktopCaptureSourceType(bindings.ConstOfDesktopCaptureSourceType(str))
}

func (x DesktopCaptureSourceType) String() (string, bool) {
	switch x {
	case DesktopCaptureSourceType_SCREEN:
		return "screen", true
	case DesktopCaptureSourceType_WINDOW:
		return "window", true
	case DesktopCaptureSourceType_TAB:
		return "tab", true
	default:
		return "", false
	}
}

type RequestInfo struct {
	// GuestProcessId is "RequestInfo.guestProcessId"
	//
	// Optional
	//
	// NOTE: FFI_USE_GuestProcessId MUST be set to true to make this field effective.
	GuestProcessId int32
	// GuestRenderFrameId is "RequestInfo.guestRenderFrameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_GuestRenderFrameId MUST be set to true to make this field effective.
	GuestRenderFrameId int32

	FFI_USE_GuestProcessId     bool // for GuestProcessId.
	FFI_USE_GuestRenderFrameId bool // for GuestRenderFrameId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestInfo with all fields set.
func (p RequestInfo) FromRef(ref js.Ref) RequestInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestInfo in the application heap.
func (p RequestInfo) New() js.Ref {
	return bindings.RequestInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestInfo) UpdateFrom(ref js.Ref) {
	bindings.RequestInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestInfo) Update(ref js.Ref) {
	bindings.RequestInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestInfo) FreeMembers(recursive bool) {
}

// HasFuncCancelChooseDesktopMedia returns true if the function "WEBEXT.webrtcDesktopCapturePrivate.cancelChooseDesktopMedia" exists.
func HasFuncCancelChooseDesktopMedia() bool {
	return js.True == bindings.HasFuncCancelChooseDesktopMedia()
}

// FuncCancelChooseDesktopMedia returns the function "WEBEXT.webrtcDesktopCapturePrivate.cancelChooseDesktopMedia".
func FuncCancelChooseDesktopMedia() (fn js.Func[func(desktopMediaRequestId int32)]) {
	bindings.FuncCancelChooseDesktopMedia(
		js.Pointer(&fn),
	)
	return
}

// CancelChooseDesktopMedia calls the function "WEBEXT.webrtcDesktopCapturePrivate.cancelChooseDesktopMedia" directly.
func CancelChooseDesktopMedia(desktopMediaRequestId int32) (ret js.Void) {
	bindings.CallCancelChooseDesktopMedia(
		js.Pointer(&ret),
		int32(desktopMediaRequestId),
	)

	return
}

// TryCancelChooseDesktopMedia calls the function "WEBEXT.webrtcDesktopCapturePrivate.cancelChooseDesktopMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCancelChooseDesktopMedia(desktopMediaRequestId int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCancelChooseDesktopMedia(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(desktopMediaRequestId),
	)

	return
}

type ChooseDesktopMediaCallbackFunc func(this js.Ref, streamId js.String) js.Ref

func (fn ChooseDesktopMediaCallbackFunc) Register() js.Func[func(streamId js.String)] {
	return js.RegisterCallback[func(streamId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ChooseDesktopMediaCallbackFunc) DispatchCallback(
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

type ChooseDesktopMediaCallback[T any] struct {
	Fn  func(arg T, this js.Ref, streamId js.String) js.Ref
	Arg T
}

func (cb *ChooseDesktopMediaCallback[T]) Register() js.Func[func(streamId js.String)] {
	return js.RegisterCallback[func(streamId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ChooseDesktopMediaCallback[T]) DispatchCallback(
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

// HasFuncChooseDesktopMedia returns true if the function "WEBEXT.webrtcDesktopCapturePrivate.chooseDesktopMedia" exists.
func HasFuncChooseDesktopMedia() bool {
	return js.True == bindings.HasFuncChooseDesktopMedia()
}

// FuncChooseDesktopMedia returns the function "WEBEXT.webrtcDesktopCapturePrivate.chooseDesktopMedia".
func FuncChooseDesktopMedia() (fn js.Func[func(sources js.Array[DesktopCaptureSourceType], request RequestInfo, callback js.Func[func(streamId js.String)]) int32]) {
	bindings.FuncChooseDesktopMedia(
		js.Pointer(&fn),
	)
	return
}

// ChooseDesktopMedia calls the function "WEBEXT.webrtcDesktopCapturePrivate.chooseDesktopMedia" directly.
func ChooseDesktopMedia(sources js.Array[DesktopCaptureSourceType], request RequestInfo, callback js.Func[func(streamId js.String)]) (ret int32) {
	bindings.CallChooseDesktopMedia(
		js.Pointer(&ret),
		sources.Ref(),
		js.Pointer(&request),
		callback.Ref(),
	)

	return
}

// TryChooseDesktopMedia calls the function "WEBEXT.webrtcDesktopCapturePrivate.chooseDesktopMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryChooseDesktopMedia(sources js.Array[DesktopCaptureSourceType], request RequestInfo, callback js.Func[func(streamId js.String)]) (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryChooseDesktopMedia(
		js.Pointer(&ret), js.Pointer(&exception),
		sources.Ref(),
		js.Pointer(&request),
		callback.Ref(),
	)

	return
}

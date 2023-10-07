// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package desktopcapture

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/desktopcapture/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabs"
)

type ChooseDesktopMediaArgCallbackFunc func(this js.Ref, streamId js.String, options *ChooseDesktopMediaArgCallbackArgOptions) js.Ref

func (fn ChooseDesktopMediaArgCallbackFunc) Register() js.Func[func(streamId js.String, options *ChooseDesktopMediaArgCallbackArgOptions)] {
	return js.RegisterCallback[func(streamId js.String, options *ChooseDesktopMediaArgCallbackArgOptions)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ChooseDesktopMediaArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 ChooseDesktopMediaArgCallbackArgOptions
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ChooseDesktopMediaArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, streamId js.String, options *ChooseDesktopMediaArgCallbackArgOptions) js.Ref
	Arg T
}

func (cb *ChooseDesktopMediaArgCallback[T]) Register() js.Func[func(streamId js.String, options *ChooseDesktopMediaArgCallbackArgOptions)] {
	return js.RegisterCallback[func(streamId js.String, options *ChooseDesktopMediaArgCallbackArgOptions)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ChooseDesktopMediaArgCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 ChooseDesktopMediaArgCallbackArgOptions
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ChooseDesktopMediaArgCallbackArgOptions struct {
	// CanRequestAudioTrack is "ChooseDesktopMediaArgCallbackArgOptions.canRequestAudioTrack"
	//
	// Required
	CanRequestAudioTrack bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ChooseDesktopMediaArgCallbackArgOptions with all fields set.
func (p ChooseDesktopMediaArgCallbackArgOptions) FromRef(ref js.Ref) ChooseDesktopMediaArgCallbackArgOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ChooseDesktopMediaArgCallbackArgOptions in the application heap.
func (p ChooseDesktopMediaArgCallbackArgOptions) New() js.Ref {
	return bindings.ChooseDesktopMediaArgCallbackArgOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ChooseDesktopMediaArgCallbackArgOptions) UpdateFrom(ref js.Ref) {
	bindings.ChooseDesktopMediaArgCallbackArgOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ChooseDesktopMediaArgCallbackArgOptions) Update(ref js.Ref) {
	bindings.ChooseDesktopMediaArgCallbackArgOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ChooseDesktopMediaArgCallbackArgOptions) FreeMembers(recursive bool) {
}

type SelfCapturePreferenceEnum uint32

const (
	_ SelfCapturePreferenceEnum = iota

	SelfCapturePreferenceEnum_INCLUDE
	SelfCapturePreferenceEnum_EXCLUDE
)

func (SelfCapturePreferenceEnum) FromRef(str js.Ref) SelfCapturePreferenceEnum {
	return SelfCapturePreferenceEnum(bindings.ConstOfSelfCapturePreferenceEnum(str))
}

func (x SelfCapturePreferenceEnum) String() (string, bool) {
	switch x {
	case SelfCapturePreferenceEnum_INCLUDE:
		return "include", true
	case SelfCapturePreferenceEnum_EXCLUDE:
		return "exclude", true
	default:
		return "", false
	}
}

type SystemAudioPreferenceEnum uint32

const (
	_ SystemAudioPreferenceEnum = iota

	SystemAudioPreferenceEnum_INCLUDE
	SystemAudioPreferenceEnum_EXCLUDE
)

func (SystemAudioPreferenceEnum) FromRef(str js.Ref) SystemAudioPreferenceEnum {
	return SystemAudioPreferenceEnum(bindings.ConstOfSystemAudioPreferenceEnum(str))
}

func (x SystemAudioPreferenceEnum) String() (string, bool) {
	switch x {
	case SystemAudioPreferenceEnum_INCLUDE:
		return "include", true
	case SystemAudioPreferenceEnum_EXCLUDE:
		return "exclude", true
	default:
		return "", false
	}
}

type ChooseDesktopMediaArgOptions struct {
	// SelfBrowserSurface is "ChooseDesktopMediaArgOptions.selfBrowserSurface"
	//
	// Optional
	SelfBrowserSurface SelfCapturePreferenceEnum
	// SuppressLocalAudioPlaybackIntended is "ChooseDesktopMediaArgOptions.suppressLocalAudioPlaybackIntended"
	//
	// Optional
	//
	// NOTE: FFI_USE_SuppressLocalAudioPlaybackIntended MUST be set to true to make this field effective.
	SuppressLocalAudioPlaybackIntended bool
	// SystemAudio is "ChooseDesktopMediaArgOptions.systemAudio"
	//
	// Optional
	SystemAudio SystemAudioPreferenceEnum

	FFI_USE_SuppressLocalAudioPlaybackIntended bool // for SuppressLocalAudioPlaybackIntended.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ChooseDesktopMediaArgOptions with all fields set.
func (p ChooseDesktopMediaArgOptions) FromRef(ref js.Ref) ChooseDesktopMediaArgOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ChooseDesktopMediaArgOptions in the application heap.
func (p ChooseDesktopMediaArgOptions) New() js.Ref {
	return bindings.ChooseDesktopMediaArgOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ChooseDesktopMediaArgOptions) UpdateFrom(ref js.Ref) {
	bindings.ChooseDesktopMediaArgOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ChooseDesktopMediaArgOptions) Update(ref js.Ref) {
	bindings.ChooseDesktopMediaArgOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ChooseDesktopMediaArgOptions) FreeMembers(recursive bool) {
}

type DesktopCaptureSourceType uint32

const (
	_ DesktopCaptureSourceType = iota

	DesktopCaptureSourceType_SCREEN
	DesktopCaptureSourceType_WINDOW
	DesktopCaptureSourceType_TAB
	DesktopCaptureSourceType_AUDIO
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
	case DesktopCaptureSourceType_AUDIO:
		return "audio", true
	default:
		return "", false
	}
}

// HasFuncCancelChooseDesktopMedia returns true if the function "WEBEXT.desktopCapture.cancelChooseDesktopMedia" exists.
func HasFuncCancelChooseDesktopMedia() bool {
	return js.True == bindings.HasFuncCancelChooseDesktopMedia()
}

// FuncCancelChooseDesktopMedia returns the function "WEBEXT.desktopCapture.cancelChooseDesktopMedia".
func FuncCancelChooseDesktopMedia() (fn js.Func[func(desktopMediaRequestId int64)]) {
	bindings.FuncCancelChooseDesktopMedia(
		js.Pointer(&fn),
	)
	return
}

// CancelChooseDesktopMedia calls the function "WEBEXT.desktopCapture.cancelChooseDesktopMedia" directly.
func CancelChooseDesktopMedia(desktopMediaRequestId int64) (ret js.Void) {
	bindings.CallCancelChooseDesktopMedia(
		js.Pointer(&ret),
		float64(desktopMediaRequestId),
	)

	return
}

// TryCancelChooseDesktopMedia calls the function "WEBEXT.desktopCapture.cancelChooseDesktopMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCancelChooseDesktopMedia(desktopMediaRequestId int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCancelChooseDesktopMedia(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(desktopMediaRequestId),
	)

	return
}

// HasFuncChooseDesktopMedia returns true if the function "WEBEXT.desktopCapture.chooseDesktopMedia" exists.
func HasFuncChooseDesktopMedia() bool {
	return js.True == bindings.HasFuncChooseDesktopMedia()
}

// FuncChooseDesktopMedia returns the function "WEBEXT.desktopCapture.chooseDesktopMedia".
func FuncChooseDesktopMedia() (fn js.Func[func(sources js.Array[DesktopCaptureSourceType], targetTab tabs.Tab, options ChooseDesktopMediaArgOptions, callback js.Func[func(streamId js.String, options *ChooseDesktopMediaArgCallbackArgOptions)]) int64]) {
	bindings.FuncChooseDesktopMedia(
		js.Pointer(&fn),
	)
	return
}

// ChooseDesktopMedia calls the function "WEBEXT.desktopCapture.chooseDesktopMedia" directly.
func ChooseDesktopMedia(sources js.Array[DesktopCaptureSourceType], targetTab tabs.Tab, options ChooseDesktopMediaArgOptions, callback js.Func[func(streamId js.String, options *ChooseDesktopMediaArgCallbackArgOptions)]) (ret int64) {
	bindings.CallChooseDesktopMedia(
		js.Pointer(&ret),
		sources.Ref(),
		js.Pointer(&targetTab),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TryChooseDesktopMedia calls the function "WEBEXT.desktopCapture.chooseDesktopMedia"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryChooseDesktopMedia(sources js.Array[DesktopCaptureSourceType], targetTab tabs.Tab, options ChooseDesktopMediaArgOptions, callback js.Func[func(streamId js.String, options *ChooseDesktopMediaArgCallbackArgOptions)]) (ret int64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryChooseDesktopMedia(
		js.Pointer(&ret), js.Pointer(&exception),
		sources.Ref(),
		js.Pointer(&targetTab),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

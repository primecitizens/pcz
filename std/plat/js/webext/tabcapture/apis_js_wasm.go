// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package tabcapture

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/tabcapture/bindings"
)

type TabCaptureState uint32

const (
	_ TabCaptureState = iota

	TabCaptureState_PENDING
	TabCaptureState_ACTIVE
	TabCaptureState_STOPPED
	TabCaptureState_ERROR
)

func (TabCaptureState) FromRef(str js.Ref) TabCaptureState {
	return TabCaptureState(bindings.ConstOfTabCaptureState(str))
}

func (x TabCaptureState) String() (string, bool) {
	switch x {
	case TabCaptureState_PENDING:
		return "pending", true
	case TabCaptureState_ACTIVE:
		return "active", true
	case TabCaptureState_STOPPED:
		return "stopped", true
	case TabCaptureState_ERROR:
		return "error", true
	default:
		return "", false
	}
}

type CaptureInfo struct {
	// TabId is "CaptureInfo.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32
	// Status is "CaptureInfo.status"
	//
	// Optional
	Status TabCaptureState
	// Fullscreen is "CaptureInfo.fullscreen"
	//
	// Optional
	//
	// NOTE: FFI_USE_Fullscreen MUST be set to true to make this field effective.
	Fullscreen bool

	FFI_USE_TabId      bool // for TabId.
	FFI_USE_Fullscreen bool // for Fullscreen.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CaptureInfo with all fields set.
func (p CaptureInfo) FromRef(ref js.Ref) CaptureInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CaptureInfo in the application heap.
func (p CaptureInfo) New() js.Ref {
	return bindings.CaptureInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CaptureInfo) UpdateFrom(ref js.Ref) {
	bindings.CaptureInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CaptureInfo) Update(ref js.Ref) {
	bindings.CaptureInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CaptureInfo) FreeMembers(recursive bool) {
}

type MediaStreamConstraint struct {
	// Mandatory is "MediaStreamConstraint.mandatory"
	//
	// Optional
	Mandatory js.Object
	// Optional is "MediaStreamConstraint.optional"
	//
	// Optional
	Optional js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaStreamConstraint with all fields set.
func (p MediaStreamConstraint) FromRef(ref js.Ref) MediaStreamConstraint {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaStreamConstraint in the application heap.
func (p MediaStreamConstraint) New() js.Ref {
	return bindings.MediaStreamConstraintJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *MediaStreamConstraint) UpdateFrom(ref js.Ref) {
	bindings.MediaStreamConstraintJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaStreamConstraint) Update(ref js.Ref) {
	bindings.MediaStreamConstraintJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaStreamConstraint) FreeMembers(recursive bool) {
	js.Free(
		p.Mandatory.Ref(),
		p.Optional.Ref(),
	)
	p.Mandatory = p.Mandatory.FromRef(js.Undefined)
	p.Optional = p.Optional.FromRef(js.Undefined)
}

type CaptureOptions struct {
	// Audio is "CaptureOptions.audio"
	//
	// Optional
	//
	// NOTE: FFI_USE_Audio MUST be set to true to make this field effective.
	Audio bool
	// Video is "CaptureOptions.video"
	//
	// Optional
	//
	// NOTE: FFI_USE_Video MUST be set to true to make this field effective.
	Video bool
	// AudioConstraints is "CaptureOptions.audioConstraints"
	//
	// Optional
	//
	// NOTE: AudioConstraints.FFI_USE MUST be set to true to get AudioConstraints used.
	AudioConstraints MediaStreamConstraint
	// VideoConstraints is "CaptureOptions.videoConstraints"
	//
	// Optional
	//
	// NOTE: VideoConstraints.FFI_USE MUST be set to true to get VideoConstraints used.
	VideoConstraints MediaStreamConstraint
	// PresentationId is "CaptureOptions.presentationId"
	//
	// Optional
	PresentationId js.String

	FFI_USE_Audio bool // for Audio.
	FFI_USE_Video bool // for Video.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CaptureOptions with all fields set.
func (p CaptureOptions) FromRef(ref js.Ref) CaptureOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CaptureOptions in the application heap.
func (p CaptureOptions) New() js.Ref {
	return bindings.CaptureOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CaptureOptions) UpdateFrom(ref js.Ref) {
	bindings.CaptureOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CaptureOptions) Update(ref js.Ref) {
	bindings.CaptureOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CaptureOptions) FreeMembers(recursive bool) {
	js.Free(
		p.PresentationId.Ref(),
	)
	p.PresentationId = p.PresentationId.FromRef(js.Undefined)
	if recursive {
		p.AudioConstraints.FreeMembers(true)
		p.VideoConstraints.FreeMembers(true)
	}
}

type GetCapturedTabsCallbackFunc func(this js.Ref, result js.Array[CaptureInfo]) js.Ref

func (fn GetCapturedTabsCallbackFunc) Register() js.Func[func(result js.Array[CaptureInfo])] {
	return js.RegisterCallback[func(result js.Array[CaptureInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetCapturedTabsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[CaptureInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetCapturedTabsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[CaptureInfo]) js.Ref
	Arg T
}

func (cb *GetCapturedTabsCallback[T]) Register() js.Func[func(result js.Array[CaptureInfo])] {
	return js.RegisterCallback[func(result js.Array[CaptureInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetCapturedTabsCallback[T]) DispatchCallback(
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

		js.Array[CaptureInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetMediaStreamIdCallbackFunc func(this js.Ref, streamId js.String) js.Ref

func (fn GetMediaStreamIdCallbackFunc) Register() js.Func[func(streamId js.String)] {
	return js.RegisterCallback[func(streamId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetMediaStreamIdCallbackFunc) DispatchCallback(
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

type GetMediaStreamIdCallback[T any] struct {
	Fn  func(arg T, this js.Ref, streamId js.String) js.Ref
	Arg T
}

func (cb *GetMediaStreamIdCallback[T]) Register() js.Func[func(streamId js.String)] {
	return js.RegisterCallback[func(streamId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetMediaStreamIdCallback[T]) DispatchCallback(
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

type GetMediaStreamOptions struct {
	// ConsumerTabId is "GetMediaStreamOptions.consumerTabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ConsumerTabId MUST be set to true to make this field effective.
	ConsumerTabId int32
	// TargetTabId is "GetMediaStreamOptions.targetTabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TargetTabId MUST be set to true to make this field effective.
	TargetTabId int32

	FFI_USE_ConsumerTabId bool // for ConsumerTabId.
	FFI_USE_TargetTabId   bool // for TargetTabId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetMediaStreamOptions with all fields set.
func (p GetMediaStreamOptions) FromRef(ref js.Ref) GetMediaStreamOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetMediaStreamOptions in the application heap.
func (p GetMediaStreamOptions) New() js.Ref {
	return bindings.GetMediaStreamOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetMediaStreamOptions) UpdateFrom(ref js.Ref) {
	bindings.GetMediaStreamOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetMediaStreamOptions) Update(ref js.Ref) {
	bindings.GetMediaStreamOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetMediaStreamOptions) FreeMembers(recursive bool) {
}

type GetTabMediaCallbackFunc func(this js.Ref, stream js.Object) js.Ref

func (fn GetTabMediaCallbackFunc) Register() js.Func[func(stream js.Object)] {
	return js.RegisterCallback[func(stream js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetTabMediaCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetTabMediaCallback[T any] struct {
	Fn  func(arg T, this js.Ref, stream js.Object) js.Ref
	Arg T
}

func (cb *GetTabMediaCallback[T]) Register() js.Func[func(stream js.Object)] {
	return js.RegisterCallback[func(stream js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetTabMediaCallback[T]) DispatchCallback(
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

		js.Object{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncCapture returns true if the function "WEBEXT.tabCapture.capture" exists.
func HasFuncCapture() bool {
	return js.True == bindings.HasFuncCapture()
}

// FuncCapture returns the function "WEBEXT.tabCapture.capture".
func FuncCapture() (fn js.Func[func(options CaptureOptions, callback js.Func[func(stream js.Object)])]) {
	bindings.FuncCapture(
		js.Pointer(&fn),
	)
	return
}

// Capture calls the function "WEBEXT.tabCapture.capture" directly.
func Capture(options CaptureOptions, callback js.Func[func(stream js.Object)]) (ret js.Void) {
	bindings.CallCapture(
		js.Pointer(&ret),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TryCapture calls the function "WEBEXT.tabCapture.capture"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCapture(options CaptureOptions, callback js.Func[func(stream js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCapture(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// HasFuncGetCapturedTabs returns true if the function "WEBEXT.tabCapture.getCapturedTabs" exists.
func HasFuncGetCapturedTabs() bool {
	return js.True == bindings.HasFuncGetCapturedTabs()
}

// FuncGetCapturedTabs returns the function "WEBEXT.tabCapture.getCapturedTabs".
func FuncGetCapturedTabs() (fn js.Func[func() js.Promise[js.Array[CaptureInfo]]]) {
	bindings.FuncGetCapturedTabs(
		js.Pointer(&fn),
	)
	return
}

// GetCapturedTabs calls the function "WEBEXT.tabCapture.getCapturedTabs" directly.
func GetCapturedTabs() (ret js.Promise[js.Array[CaptureInfo]]) {
	bindings.CallGetCapturedTabs(
		js.Pointer(&ret),
	)

	return
}

// TryGetCapturedTabs calls the function "WEBEXT.tabCapture.getCapturedTabs"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetCapturedTabs() (ret js.Promise[js.Array[CaptureInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetCapturedTabs(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetMediaStreamId returns true if the function "WEBEXT.tabCapture.getMediaStreamId" exists.
func HasFuncGetMediaStreamId() bool {
	return js.True == bindings.HasFuncGetMediaStreamId()
}

// FuncGetMediaStreamId returns the function "WEBEXT.tabCapture.getMediaStreamId".
func FuncGetMediaStreamId() (fn js.Func[func(options GetMediaStreamOptions) js.Promise[js.String]]) {
	bindings.FuncGetMediaStreamId(
		js.Pointer(&fn),
	)
	return
}

// GetMediaStreamId calls the function "WEBEXT.tabCapture.getMediaStreamId" directly.
func GetMediaStreamId(options GetMediaStreamOptions) (ret js.Promise[js.String]) {
	bindings.CallGetMediaStreamId(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryGetMediaStreamId calls the function "WEBEXT.tabCapture.getMediaStreamId"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetMediaStreamId(options GetMediaStreamOptions) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetMediaStreamId(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

type OnStatusChangedEventCallbackFunc func(this js.Ref, info *CaptureInfo) js.Ref

func (fn OnStatusChangedEventCallbackFunc) Register() js.Func[func(info *CaptureInfo)] {
	return js.RegisterCallback[func(info *CaptureInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnStatusChangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CaptureInfo
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

type OnStatusChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, info *CaptureInfo) js.Ref
	Arg T
}

func (cb *OnStatusChangedEventCallback[T]) Register() js.Func[func(info *CaptureInfo)] {
	return js.RegisterCallback[func(info *CaptureInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnStatusChangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 CaptureInfo
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

// HasFuncOnStatusChanged returns true if the function "WEBEXT.tabCapture.onStatusChanged.addListener" exists.
func HasFuncOnStatusChanged() bool {
	return js.True == bindings.HasFuncOnStatusChanged()
}

// FuncOnStatusChanged returns the function "WEBEXT.tabCapture.onStatusChanged.addListener".
func FuncOnStatusChanged() (fn js.Func[func(callback js.Func[func(info *CaptureInfo)])]) {
	bindings.FuncOnStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// OnStatusChanged calls the function "WEBEXT.tabCapture.onStatusChanged.addListener" directly.
func OnStatusChanged(callback js.Func[func(info *CaptureInfo)]) (ret js.Void) {
	bindings.CallOnStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnStatusChanged calls the function "WEBEXT.tabCapture.onStatusChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnStatusChanged(callback js.Func[func(info *CaptureInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffStatusChanged returns true if the function "WEBEXT.tabCapture.onStatusChanged.removeListener" exists.
func HasFuncOffStatusChanged() bool {
	return js.True == bindings.HasFuncOffStatusChanged()
}

// FuncOffStatusChanged returns the function "WEBEXT.tabCapture.onStatusChanged.removeListener".
func FuncOffStatusChanged() (fn js.Func[func(callback js.Func[func(info *CaptureInfo)])]) {
	bindings.FuncOffStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// OffStatusChanged calls the function "WEBEXT.tabCapture.onStatusChanged.removeListener" directly.
func OffStatusChanged(callback js.Func[func(info *CaptureInfo)]) (ret js.Void) {
	bindings.CallOffStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffStatusChanged calls the function "WEBEXT.tabCapture.onStatusChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffStatusChanged(callback js.Func[func(info *CaptureInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnStatusChanged returns true if the function "WEBEXT.tabCapture.onStatusChanged.hasListener" exists.
func HasFuncHasOnStatusChanged() bool {
	return js.True == bindings.HasFuncHasOnStatusChanged()
}

// FuncHasOnStatusChanged returns the function "WEBEXT.tabCapture.onStatusChanged.hasListener".
func FuncHasOnStatusChanged() (fn js.Func[func(callback js.Func[func(info *CaptureInfo)]) bool]) {
	bindings.FuncHasOnStatusChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnStatusChanged calls the function "WEBEXT.tabCapture.onStatusChanged.hasListener" directly.
func HasOnStatusChanged(callback js.Func[func(info *CaptureInfo)]) (ret bool) {
	bindings.CallHasOnStatusChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnStatusChanged calls the function "WEBEXT.tabCapture.onStatusChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnStatusChanged(callback js.Func[func(info *CaptureInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnStatusChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

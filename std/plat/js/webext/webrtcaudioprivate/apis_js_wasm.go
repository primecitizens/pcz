// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webrtcaudioprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/webrtcaudioprivate/bindings"
)

type GetSinksCallbackFunc func(this js.Ref, sinkInfo js.Array[SinkInfo]) js.Ref

func (fn GetSinksCallbackFunc) Register() js.Func[func(sinkInfo js.Array[SinkInfo])] {
	return js.RegisterCallback[func(sinkInfo js.Array[SinkInfo])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetSinksCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[SinkInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetSinksCallback[T any] struct {
	Fn  func(arg T, this js.Ref, sinkInfo js.Array[SinkInfo]) js.Ref
	Arg T
}

func (cb *GetSinksCallback[T]) Register() js.Func[func(sinkInfo js.Array[SinkInfo])] {
	return js.RegisterCallback[func(sinkInfo js.Array[SinkInfo])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetSinksCallback[T]) DispatchCallback(
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

		js.Array[SinkInfo]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SinkInfo struct {
	// SinkId is "SinkInfo.sinkId"
	//
	// Optional
	SinkId js.String
	// SinkLabel is "SinkInfo.sinkLabel"
	//
	// Optional
	SinkLabel js.String
	// SampleRate is "SinkInfo.sampleRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_SampleRate MUST be set to true to make this field effective.
	SampleRate int32
	// IsReady is "SinkInfo.isReady"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsReady MUST be set to true to make this field effective.
	IsReady bool
	// IsDefault is "SinkInfo.isDefault"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDefault MUST be set to true to make this field effective.
	IsDefault bool

	FFI_USE_SampleRate bool // for SampleRate.
	FFI_USE_IsReady    bool // for IsReady.
	FFI_USE_IsDefault  bool // for IsDefault.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SinkInfo with all fields set.
func (p SinkInfo) FromRef(ref js.Ref) SinkInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SinkInfo in the application heap.
func (p SinkInfo) New() js.Ref {
	return bindings.SinkInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SinkInfo) UpdateFrom(ref js.Ref) {
	bindings.SinkInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SinkInfo) Update(ref js.Ref) {
	bindings.SinkInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SinkInfo) FreeMembers(recursive bool) {
	js.Free(
		p.SinkId.Ref(),
		p.SinkLabel.Ref(),
	)
	p.SinkId = p.SinkId.FromRef(js.Undefined)
	p.SinkLabel = p.SinkLabel.FromRef(js.Undefined)
}

type SinkIdCallbackFunc func(this js.Ref, sinkId js.String) js.Ref

func (fn SinkIdCallbackFunc) Register() js.Func[func(sinkId js.String)] {
	return js.RegisterCallback[func(sinkId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SinkIdCallbackFunc) DispatchCallback(
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

type SinkIdCallback[T any] struct {
	Fn  func(arg T, this js.Ref, sinkId js.String) js.Ref
	Arg T
}

func (cb *SinkIdCallback[T]) Register() js.Func[func(sinkId js.String)] {
	return js.RegisterCallback[func(sinkId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SinkIdCallback[T]) DispatchCallback(
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

// HasFuncGetAssociatedSink returns true if the function "WEBEXT.webrtcAudioPrivate.getAssociatedSink" exists.
func HasFuncGetAssociatedSink() bool {
	return js.True == bindings.HasFuncGetAssociatedSink()
}

// FuncGetAssociatedSink returns the function "WEBEXT.webrtcAudioPrivate.getAssociatedSink".
func FuncGetAssociatedSink() (fn js.Func[func(securityOrigin js.String, sourceIdInOrigin js.String) js.Promise[js.String]]) {
	bindings.FuncGetAssociatedSink(
		js.Pointer(&fn),
	)
	return
}

// GetAssociatedSink calls the function "WEBEXT.webrtcAudioPrivate.getAssociatedSink" directly.
func GetAssociatedSink(securityOrigin js.String, sourceIdInOrigin js.String) (ret js.Promise[js.String]) {
	bindings.CallGetAssociatedSink(
		js.Pointer(&ret),
		securityOrigin.Ref(),
		sourceIdInOrigin.Ref(),
	)

	return
}

// TryGetAssociatedSink calls the function "WEBEXT.webrtcAudioPrivate.getAssociatedSink"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAssociatedSink(securityOrigin js.String, sourceIdInOrigin js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAssociatedSink(
		js.Pointer(&ret), js.Pointer(&exception),
		securityOrigin.Ref(),
		sourceIdInOrigin.Ref(),
	)

	return
}

// HasFuncGetSinks returns true if the function "WEBEXT.webrtcAudioPrivate.getSinks" exists.
func HasFuncGetSinks() bool {
	return js.True == bindings.HasFuncGetSinks()
}

// FuncGetSinks returns the function "WEBEXT.webrtcAudioPrivate.getSinks".
func FuncGetSinks() (fn js.Func[func() js.Promise[js.Array[SinkInfo]]]) {
	bindings.FuncGetSinks(
		js.Pointer(&fn),
	)
	return
}

// GetSinks calls the function "WEBEXT.webrtcAudioPrivate.getSinks" directly.
func GetSinks() (ret js.Promise[js.Array[SinkInfo]]) {
	bindings.CallGetSinks(
		js.Pointer(&ret),
	)

	return
}

// TryGetSinks calls the function "WEBEXT.webrtcAudioPrivate.getSinks"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSinks() (ret js.Promise[js.Array[SinkInfo]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSinks(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnSinksChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnSinksChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSinksChangedEventCallbackFunc) DispatchCallback(
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

type OnSinksChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnSinksChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSinksChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnSinksChanged returns true if the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.addListener" exists.
func HasFuncOnSinksChanged() bool {
	return js.True == bindings.HasFuncOnSinksChanged()
}

// FuncOnSinksChanged returns the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.addListener".
func FuncOnSinksChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnSinksChanged(
		js.Pointer(&fn),
	)
	return
}

// OnSinksChanged calls the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.addListener" directly.
func OnSinksChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnSinksChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSinksChanged calls the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSinksChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSinksChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSinksChanged returns true if the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.removeListener" exists.
func HasFuncOffSinksChanged() bool {
	return js.True == bindings.HasFuncOffSinksChanged()
}

// FuncOffSinksChanged returns the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.removeListener".
func FuncOffSinksChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffSinksChanged(
		js.Pointer(&fn),
	)
	return
}

// OffSinksChanged calls the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.removeListener" directly.
func OffSinksChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffSinksChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSinksChanged calls the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSinksChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSinksChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSinksChanged returns true if the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.hasListener" exists.
func HasFuncHasOnSinksChanged() bool {
	return js.True == bindings.HasFuncHasOnSinksChanged()
}

// FuncHasOnSinksChanged returns the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.hasListener".
func FuncHasOnSinksChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnSinksChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSinksChanged calls the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.hasListener" directly.
func HasOnSinksChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnSinksChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSinksChanged calls the function "WEBEXT.webrtcAudioPrivate.onSinksChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSinksChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSinksChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

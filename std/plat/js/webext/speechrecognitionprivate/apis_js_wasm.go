// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package speechrecognitionprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/speechrecognitionprivate/bindings"
)

type OnStartCallbackFunc func(this js.Ref, typ SpeechRecognitionType) js.Ref

func (fn OnStartCallbackFunc) Register() js.Func[func(typ SpeechRecognitionType)] {
	return js.RegisterCallback[func(typ SpeechRecognitionType)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnStartCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		SpeechRecognitionType(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnStartCallback[T any] struct {
	Fn  func(arg T, this js.Ref, typ SpeechRecognitionType) js.Ref
	Arg T
}

func (cb *OnStartCallback[T]) Register() js.Func[func(typ SpeechRecognitionType)] {
	return js.RegisterCallback[func(typ SpeechRecognitionType)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnStartCallback[T]) DispatchCallback(
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

		SpeechRecognitionType(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SpeechRecognitionType uint32

const (
	_ SpeechRecognitionType = iota

	SpeechRecognitionType_ON_DEVICE
	SpeechRecognitionType_NETWORK
)

func (SpeechRecognitionType) FromRef(str js.Ref) SpeechRecognitionType {
	return SpeechRecognitionType(bindings.ConstOfSpeechRecognitionType(str))
}

func (x SpeechRecognitionType) String() (string, bool) {
	switch x {
	case SpeechRecognitionType_ON_DEVICE:
		return "onDevice", true
	case SpeechRecognitionType_NETWORK:
		return "network", true
	default:
		return "", false
	}
}

type OnStopCallbackFunc func(this js.Ref) js.Ref

func (fn OnStopCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnStopCallbackFunc) DispatchCallback(
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

type OnStopCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnStopCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnStopCallback[T]) DispatchCallback(
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

type SpeechRecognitionErrorEvent struct {
	// ClientId is "SpeechRecognitionErrorEvent.clientId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ClientId MUST be set to true to make this field effective.
	ClientId int32
	// Message is "SpeechRecognitionErrorEvent.message"
	//
	// Optional
	Message js.String

	FFI_USE_ClientId bool // for ClientId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SpeechRecognitionErrorEvent with all fields set.
func (p SpeechRecognitionErrorEvent) FromRef(ref js.Ref) SpeechRecognitionErrorEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SpeechRecognitionErrorEvent in the application heap.
func (p SpeechRecognitionErrorEvent) New() js.Ref {
	return bindings.SpeechRecognitionErrorEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SpeechRecognitionErrorEvent) UpdateFrom(ref js.Ref) {
	bindings.SpeechRecognitionErrorEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SpeechRecognitionErrorEvent) Update(ref js.Ref) {
	bindings.SpeechRecognitionErrorEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SpeechRecognitionErrorEvent) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
}

type SpeechRecognitionResultEvent struct {
	// ClientId is "SpeechRecognitionResultEvent.clientId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ClientId MUST be set to true to make this field effective.
	ClientId int32
	// Transcript is "SpeechRecognitionResultEvent.transcript"
	//
	// Optional
	Transcript js.String
	// IsFinal is "SpeechRecognitionResultEvent.isFinal"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsFinal MUST be set to true to make this field effective.
	IsFinal bool

	FFI_USE_ClientId bool // for ClientId.
	FFI_USE_IsFinal  bool // for IsFinal.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SpeechRecognitionResultEvent with all fields set.
func (p SpeechRecognitionResultEvent) FromRef(ref js.Ref) SpeechRecognitionResultEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SpeechRecognitionResultEvent in the application heap.
func (p SpeechRecognitionResultEvent) New() js.Ref {
	return bindings.SpeechRecognitionResultEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SpeechRecognitionResultEvent) UpdateFrom(ref js.Ref) {
	bindings.SpeechRecognitionResultEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SpeechRecognitionResultEvent) Update(ref js.Ref) {
	bindings.SpeechRecognitionResultEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SpeechRecognitionResultEvent) FreeMembers(recursive bool) {
	js.Free(
		p.Transcript.Ref(),
	)
	p.Transcript = p.Transcript.FromRef(js.Undefined)
}

type SpeechRecognitionStopEvent struct {
	// ClientId is "SpeechRecognitionStopEvent.clientId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ClientId MUST be set to true to make this field effective.
	ClientId int32

	FFI_USE_ClientId bool // for ClientId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SpeechRecognitionStopEvent with all fields set.
func (p SpeechRecognitionStopEvent) FromRef(ref js.Ref) SpeechRecognitionStopEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SpeechRecognitionStopEvent in the application heap.
func (p SpeechRecognitionStopEvent) New() js.Ref {
	return bindings.SpeechRecognitionStopEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SpeechRecognitionStopEvent) UpdateFrom(ref js.Ref) {
	bindings.SpeechRecognitionStopEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SpeechRecognitionStopEvent) Update(ref js.Ref) {
	bindings.SpeechRecognitionStopEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SpeechRecognitionStopEvent) FreeMembers(recursive bool) {
}

type StartOptions struct {
	// ClientId is "StartOptions.clientId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ClientId MUST be set to true to make this field effective.
	ClientId int32
	// Locale is "StartOptions.locale"
	//
	// Optional
	Locale js.String
	// InterimResults is "StartOptions.interimResults"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterimResults MUST be set to true to make this field effective.
	InterimResults bool

	FFI_USE_ClientId       bool // for ClientId.
	FFI_USE_InterimResults bool // for InterimResults.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StartOptions with all fields set.
func (p StartOptions) FromRef(ref js.Ref) StartOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StartOptions in the application heap.
func (p StartOptions) New() js.Ref {
	return bindings.StartOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StartOptions) UpdateFrom(ref js.Ref) {
	bindings.StartOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StartOptions) Update(ref js.Ref) {
	bindings.StartOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StartOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Locale.Ref(),
	)
	p.Locale = p.Locale.FromRef(js.Undefined)
}

type StopOptions struct {
	// ClientId is "StopOptions.clientId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ClientId MUST be set to true to make this field effective.
	ClientId int32

	FFI_USE_ClientId bool // for ClientId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StopOptions with all fields set.
func (p StopOptions) FromRef(ref js.Ref) StopOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StopOptions in the application heap.
func (p StopOptions) New() js.Ref {
	return bindings.StopOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StopOptions) UpdateFrom(ref js.Ref) {
	bindings.StopOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StopOptions) Update(ref js.Ref) {
	bindings.StopOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StopOptions) FreeMembers(recursive bool) {
}

type OnErrorEventCallbackFunc func(this js.Ref, event *SpeechRecognitionErrorEvent) js.Ref

func (fn OnErrorEventCallbackFunc) Register() js.Func[func(event *SpeechRecognitionErrorEvent)] {
	return js.RegisterCallback[func(event *SpeechRecognitionErrorEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnErrorEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SpeechRecognitionErrorEvent
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

type OnErrorEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *SpeechRecognitionErrorEvent) js.Ref
	Arg T
}

func (cb *OnErrorEventCallback[T]) Register() js.Func[func(event *SpeechRecognitionErrorEvent)] {
	return js.RegisterCallback[func(event *SpeechRecognitionErrorEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnErrorEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SpeechRecognitionErrorEvent
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

// HasFuncOnError returns true if the function "WEBEXT.speechRecognitionPrivate.onError.addListener" exists.
func HasFuncOnError() bool {
	return js.True == bindings.HasFuncOnError()
}

// FuncOnError returns the function "WEBEXT.speechRecognitionPrivate.onError.addListener".
func FuncOnError() (fn js.Func[func(callback js.Func[func(event *SpeechRecognitionErrorEvent)])]) {
	bindings.FuncOnError(
		js.Pointer(&fn),
	)
	return
}

// OnError calls the function "WEBEXT.speechRecognitionPrivate.onError.addListener" directly.
func OnError(callback js.Func[func(event *SpeechRecognitionErrorEvent)]) (ret js.Void) {
	bindings.CallOnError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnError calls the function "WEBEXT.speechRecognitionPrivate.onError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnError(callback js.Func[func(event *SpeechRecognitionErrorEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffError returns true if the function "WEBEXT.speechRecognitionPrivate.onError.removeListener" exists.
func HasFuncOffError() bool {
	return js.True == bindings.HasFuncOffError()
}

// FuncOffError returns the function "WEBEXT.speechRecognitionPrivate.onError.removeListener".
func FuncOffError() (fn js.Func[func(callback js.Func[func(event *SpeechRecognitionErrorEvent)])]) {
	bindings.FuncOffError(
		js.Pointer(&fn),
	)
	return
}

// OffError calls the function "WEBEXT.speechRecognitionPrivate.onError.removeListener" directly.
func OffError(callback js.Func[func(event *SpeechRecognitionErrorEvent)]) (ret js.Void) {
	bindings.CallOffError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffError calls the function "WEBEXT.speechRecognitionPrivate.onError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffError(callback js.Func[func(event *SpeechRecognitionErrorEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnError returns true if the function "WEBEXT.speechRecognitionPrivate.onError.hasListener" exists.
func HasFuncHasOnError() bool {
	return js.True == bindings.HasFuncHasOnError()
}

// FuncHasOnError returns the function "WEBEXT.speechRecognitionPrivate.onError.hasListener".
func FuncHasOnError() (fn js.Func[func(callback js.Func[func(event *SpeechRecognitionErrorEvent)]) bool]) {
	bindings.FuncHasOnError(
		js.Pointer(&fn),
	)
	return
}

// HasOnError calls the function "WEBEXT.speechRecognitionPrivate.onError.hasListener" directly.
func HasOnError(callback js.Func[func(event *SpeechRecognitionErrorEvent)]) (ret bool) {
	bindings.CallHasOnError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnError calls the function "WEBEXT.speechRecognitionPrivate.onError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnError(callback js.Func[func(event *SpeechRecognitionErrorEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnResultEventCallbackFunc func(this js.Ref, event *SpeechRecognitionResultEvent) js.Ref

func (fn OnResultEventCallbackFunc) Register() js.Func[func(event *SpeechRecognitionResultEvent)] {
	return js.RegisterCallback[func(event *SpeechRecognitionResultEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnResultEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SpeechRecognitionResultEvent
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

type OnResultEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *SpeechRecognitionResultEvent) js.Ref
	Arg T
}

func (cb *OnResultEventCallback[T]) Register() js.Func[func(event *SpeechRecognitionResultEvent)] {
	return js.RegisterCallback[func(event *SpeechRecognitionResultEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnResultEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SpeechRecognitionResultEvent
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

// HasFuncOnResult returns true if the function "WEBEXT.speechRecognitionPrivate.onResult.addListener" exists.
func HasFuncOnResult() bool {
	return js.True == bindings.HasFuncOnResult()
}

// FuncOnResult returns the function "WEBEXT.speechRecognitionPrivate.onResult.addListener".
func FuncOnResult() (fn js.Func[func(callback js.Func[func(event *SpeechRecognitionResultEvent)])]) {
	bindings.FuncOnResult(
		js.Pointer(&fn),
	)
	return
}

// OnResult calls the function "WEBEXT.speechRecognitionPrivate.onResult.addListener" directly.
func OnResult(callback js.Func[func(event *SpeechRecognitionResultEvent)]) (ret js.Void) {
	bindings.CallOnResult(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnResult calls the function "WEBEXT.speechRecognitionPrivate.onResult.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnResult(callback js.Func[func(event *SpeechRecognitionResultEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnResult(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffResult returns true if the function "WEBEXT.speechRecognitionPrivate.onResult.removeListener" exists.
func HasFuncOffResult() bool {
	return js.True == bindings.HasFuncOffResult()
}

// FuncOffResult returns the function "WEBEXT.speechRecognitionPrivate.onResult.removeListener".
func FuncOffResult() (fn js.Func[func(callback js.Func[func(event *SpeechRecognitionResultEvent)])]) {
	bindings.FuncOffResult(
		js.Pointer(&fn),
	)
	return
}

// OffResult calls the function "WEBEXT.speechRecognitionPrivate.onResult.removeListener" directly.
func OffResult(callback js.Func[func(event *SpeechRecognitionResultEvent)]) (ret js.Void) {
	bindings.CallOffResult(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffResult calls the function "WEBEXT.speechRecognitionPrivate.onResult.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffResult(callback js.Func[func(event *SpeechRecognitionResultEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffResult(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnResult returns true if the function "WEBEXT.speechRecognitionPrivate.onResult.hasListener" exists.
func HasFuncHasOnResult() bool {
	return js.True == bindings.HasFuncHasOnResult()
}

// FuncHasOnResult returns the function "WEBEXT.speechRecognitionPrivate.onResult.hasListener".
func FuncHasOnResult() (fn js.Func[func(callback js.Func[func(event *SpeechRecognitionResultEvent)]) bool]) {
	bindings.FuncHasOnResult(
		js.Pointer(&fn),
	)
	return
}

// HasOnResult calls the function "WEBEXT.speechRecognitionPrivate.onResult.hasListener" directly.
func HasOnResult(callback js.Func[func(event *SpeechRecognitionResultEvent)]) (ret bool) {
	bindings.CallHasOnResult(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnResult calls the function "WEBEXT.speechRecognitionPrivate.onResult.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnResult(callback js.Func[func(event *SpeechRecognitionResultEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnResult(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnStopEventCallbackFunc func(this js.Ref, event *SpeechRecognitionStopEvent) js.Ref

func (fn OnStopEventCallbackFunc) Register() js.Func[func(event *SpeechRecognitionStopEvent)] {
	return js.RegisterCallback[func(event *SpeechRecognitionStopEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnStopEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SpeechRecognitionStopEvent
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

type OnStopEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *SpeechRecognitionStopEvent) js.Ref
	Arg T
}

func (cb *OnStopEventCallback[T]) Register() js.Func[func(event *SpeechRecognitionStopEvent)] {
	return js.RegisterCallback[func(event *SpeechRecognitionStopEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnStopEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SpeechRecognitionStopEvent
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

// HasFuncOnStop returns true if the function "WEBEXT.speechRecognitionPrivate.onStop.addListener" exists.
func HasFuncOnStop() bool {
	return js.True == bindings.HasFuncOnStop()
}

// FuncOnStop returns the function "WEBEXT.speechRecognitionPrivate.onStop.addListener".
func FuncOnStop() (fn js.Func[func(callback js.Func[func(event *SpeechRecognitionStopEvent)])]) {
	bindings.FuncOnStop(
		js.Pointer(&fn),
	)
	return
}

// OnStop calls the function "WEBEXT.speechRecognitionPrivate.onStop.addListener" directly.
func OnStop(callback js.Func[func(event *SpeechRecognitionStopEvent)]) (ret js.Void) {
	bindings.CallOnStop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnStop calls the function "WEBEXT.speechRecognitionPrivate.onStop.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnStop(callback js.Func[func(event *SpeechRecognitionStopEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnStop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffStop returns true if the function "WEBEXT.speechRecognitionPrivate.onStop.removeListener" exists.
func HasFuncOffStop() bool {
	return js.True == bindings.HasFuncOffStop()
}

// FuncOffStop returns the function "WEBEXT.speechRecognitionPrivate.onStop.removeListener".
func FuncOffStop() (fn js.Func[func(callback js.Func[func(event *SpeechRecognitionStopEvent)])]) {
	bindings.FuncOffStop(
		js.Pointer(&fn),
	)
	return
}

// OffStop calls the function "WEBEXT.speechRecognitionPrivate.onStop.removeListener" directly.
func OffStop(callback js.Func[func(event *SpeechRecognitionStopEvent)]) (ret js.Void) {
	bindings.CallOffStop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffStop calls the function "WEBEXT.speechRecognitionPrivate.onStop.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffStop(callback js.Func[func(event *SpeechRecognitionStopEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffStop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnStop returns true if the function "WEBEXT.speechRecognitionPrivate.onStop.hasListener" exists.
func HasFuncHasOnStop() bool {
	return js.True == bindings.HasFuncHasOnStop()
}

// FuncHasOnStop returns the function "WEBEXT.speechRecognitionPrivate.onStop.hasListener".
func FuncHasOnStop() (fn js.Func[func(callback js.Func[func(event *SpeechRecognitionStopEvent)]) bool]) {
	bindings.FuncHasOnStop(
		js.Pointer(&fn),
	)
	return
}

// HasOnStop calls the function "WEBEXT.speechRecognitionPrivate.onStop.hasListener" directly.
func HasOnStop(callback js.Func[func(event *SpeechRecognitionStopEvent)]) (ret bool) {
	bindings.CallHasOnStop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnStop calls the function "WEBEXT.speechRecognitionPrivate.onStop.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnStop(callback js.Func[func(event *SpeechRecognitionStopEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnStop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncStart returns true if the function "WEBEXT.speechRecognitionPrivate.start" exists.
func HasFuncStart() bool {
	return js.True == bindings.HasFuncStart()
}

// FuncStart returns the function "WEBEXT.speechRecognitionPrivate.start".
func FuncStart() (fn js.Func[func(options StartOptions) js.Promise[SpeechRecognitionType]]) {
	bindings.FuncStart(
		js.Pointer(&fn),
	)
	return
}

// Start calls the function "WEBEXT.speechRecognitionPrivate.start" directly.
func Start(options StartOptions) (ret js.Promise[SpeechRecognitionType]) {
	bindings.CallStart(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryStart calls the function "WEBEXT.speechRecognitionPrivate.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStart(options StartOptions) (ret js.Promise[SpeechRecognitionType], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStart(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncStop returns true if the function "WEBEXT.speechRecognitionPrivate.stop" exists.
func HasFuncStop() bool {
	return js.True == bindings.HasFuncStop()
}

// FuncStop returns the function "WEBEXT.speechRecognitionPrivate.stop".
func FuncStop() (fn js.Func[func(options StopOptions) js.Promise[js.Void]]) {
	bindings.FuncStop(
		js.Pointer(&fn),
	)
	return
}

// Stop calls the function "WEBEXT.speechRecognitionPrivate.stop" directly.
func Stop(options StopOptions) (ret js.Promise[js.Void]) {
	bindings.CallStop(
		js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryStop calls the function "WEBEXT.speechRecognitionPrivate.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStop(options StopOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryStop(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package tts

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/tts/bindings"
)

type EventType uint32

const (
	_ EventType = iota

	EventType_START
	EventType_END
	EventType_WORD
	EventType_SENTENCE
	EventType_MARKER
	EventType_INTERRUPTED
	EventType_CANCELLED
	EventType_ERROR
	EventType_PAUSE
	EventType_RESUME
)

func (EventType) FromRef(str js.Ref) EventType {
	return EventType(bindings.ConstOfEventType(str))
}

func (x EventType) String() (string, bool) {
	switch x {
	case EventType_START:
		return "start", true
	case EventType_END:
		return "end", true
	case EventType_WORD:
		return "word", true
	case EventType_SENTENCE:
		return "sentence", true
	case EventType_MARKER:
		return "marker", true
	case EventType_INTERRUPTED:
		return "interrupted", true
	case EventType_CANCELLED:
		return "cancelled", true
	case EventType_ERROR:
		return "error", true
	case EventType_PAUSE:
		return "pause", true
	case EventType_RESUME:
		return "resume", true
	default:
		return "", false
	}
}

type TtsEvent struct {
	// CharIndex is "TtsEvent.charIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_CharIndex MUST be set to true to make this field effective.
	CharIndex int64
	// ErrorMessage is "TtsEvent.errorMessage"
	//
	// Optional
	ErrorMessage js.String
	// IsFinalEvent is "TtsEvent.isFinalEvent"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsFinalEvent MUST be set to true to make this field effective.
	IsFinalEvent bool
	// Length is "TtsEvent.length"
	//
	// Optional
	//
	// NOTE: FFI_USE_Length MUST be set to true to make this field effective.
	Length int64
	// SrcId is "TtsEvent.srcId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SrcId MUST be set to true to make this field effective.
	SrcId float64
	// Type is "TtsEvent.type"
	//
	// Required
	Type EventType

	FFI_USE_CharIndex    bool // for CharIndex.
	FFI_USE_IsFinalEvent bool // for IsFinalEvent.
	FFI_USE_Length       bool // for Length.
	FFI_USE_SrcId        bool // for SrcId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TtsEvent with all fields set.
func (p TtsEvent) FromRef(ref js.Ref) TtsEvent {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TtsEvent in the application heap.
func (p TtsEvent) New() js.Ref {
	return bindings.TtsEventJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TtsEvent) UpdateFrom(ref js.Ref) {
	bindings.TtsEventJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TtsEvent) Update(ref js.Ref) {
	bindings.TtsEventJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TtsEvent) FreeMembers(recursive bool) {
	js.Free(
		p.ErrorMessage.Ref(),
	)
	p.ErrorMessage = p.ErrorMessage.FromRef(js.Undefined)
}

type VoiceGender uint32

const (
	_ VoiceGender = iota

	VoiceGender_MALE
	VoiceGender_FEMALE
)

func (VoiceGender) FromRef(str js.Ref) VoiceGender {
	return VoiceGender(bindings.ConstOfVoiceGender(str))
}

func (x VoiceGender) String() (string, bool) {
	switch x {
	case VoiceGender_MALE:
		return "male", true
	case VoiceGender_FEMALE:
		return "female", true
	default:
		return "", false
	}
}

type TtsOptionsFieldOnEventFunc func(this js.Ref, event *TtsEvent) js.Ref

func (fn TtsOptionsFieldOnEventFunc) Register() js.Func[func(event *TtsEvent)] {
	return js.RegisterCallback[func(event *TtsEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn TtsOptionsFieldOnEventFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TtsEvent
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

type TtsOptionsFieldOnEvent[T any] struct {
	Fn  func(arg T, this js.Ref, event *TtsEvent) js.Ref
	Arg T
}

func (cb *TtsOptionsFieldOnEvent[T]) Register() js.Func[func(event *TtsEvent)] {
	return js.RegisterCallback[func(event *TtsEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *TtsOptionsFieldOnEvent[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TtsEvent
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

type TtsOptions struct {
	// DesiredEventTypes is "TtsOptions.desiredEventTypes"
	//
	// Optional
	DesiredEventTypes js.Array[js.String]
	// Enqueue is "TtsOptions.enqueue"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enqueue MUST be set to true to make this field effective.
	Enqueue bool
	// ExtensionId is "TtsOptions.extensionId"
	//
	// Optional
	ExtensionId js.String
	// Gender is "TtsOptions.gender"
	//
	// Optional
	Gender VoiceGender
	// Lang is "TtsOptions.lang"
	//
	// Optional
	Lang js.String
	// OnEvent is "TtsOptions.onEvent"
	//
	// Optional
	OnEvent js.Func[func(event *TtsEvent)]
	// Pitch is "TtsOptions.pitch"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pitch MUST be set to true to make this field effective.
	Pitch float64
	// Rate is "TtsOptions.rate"
	//
	// Optional
	//
	// NOTE: FFI_USE_Rate MUST be set to true to make this field effective.
	Rate float64
	// RequiredEventTypes is "TtsOptions.requiredEventTypes"
	//
	// Optional
	RequiredEventTypes js.Array[js.String]
	// VoiceName is "TtsOptions.voiceName"
	//
	// Optional
	VoiceName js.String
	// Volume is "TtsOptions.volume"
	//
	// Optional
	//
	// NOTE: FFI_USE_Volume MUST be set to true to make this field effective.
	Volume float64

	FFI_USE_Enqueue bool // for Enqueue.
	FFI_USE_Pitch   bool // for Pitch.
	FFI_USE_Rate    bool // for Rate.
	FFI_USE_Volume  bool // for Volume.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TtsOptions with all fields set.
func (p TtsOptions) FromRef(ref js.Ref) TtsOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TtsOptions in the application heap.
func (p TtsOptions) New() js.Ref {
	return bindings.TtsOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TtsOptions) UpdateFrom(ref js.Ref) {
	bindings.TtsOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TtsOptions) Update(ref js.Ref) {
	bindings.TtsOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TtsOptions) FreeMembers(recursive bool) {
	js.Free(
		p.DesiredEventTypes.Ref(),
		p.ExtensionId.Ref(),
		p.Lang.Ref(),
		p.OnEvent.Ref(),
		p.RequiredEventTypes.Ref(),
		p.VoiceName.Ref(),
	)
	p.DesiredEventTypes = p.DesiredEventTypes.FromRef(js.Undefined)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.Lang = p.Lang.FromRef(js.Undefined)
	p.OnEvent = p.OnEvent.FromRef(js.Undefined)
	p.RequiredEventTypes = p.RequiredEventTypes.FromRef(js.Undefined)
	p.VoiceName = p.VoiceName.FromRef(js.Undefined)
}

type TtsVoice struct {
	// EventTypes is "TtsVoice.eventTypes"
	//
	// Optional
	EventTypes js.Array[EventType]
	// ExtensionId is "TtsVoice.extensionId"
	//
	// Optional
	ExtensionId js.String
	// Gender is "TtsVoice.gender"
	//
	// Optional
	Gender VoiceGender
	// Lang is "TtsVoice.lang"
	//
	// Optional
	Lang js.String
	// Remote is "TtsVoice.remote"
	//
	// Optional
	//
	// NOTE: FFI_USE_Remote MUST be set to true to make this field effective.
	Remote bool
	// VoiceName is "TtsVoice.voiceName"
	//
	// Optional
	VoiceName js.String

	FFI_USE_Remote bool // for Remote.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a TtsVoice with all fields set.
func (p TtsVoice) FromRef(ref js.Ref) TtsVoice {
	p.UpdateFrom(ref)
	return p
}

// New creates a new TtsVoice in the application heap.
func (p TtsVoice) New() js.Ref {
	return bindings.TtsVoiceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *TtsVoice) UpdateFrom(ref js.Ref) {
	bindings.TtsVoiceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *TtsVoice) Update(ref js.Ref) {
	bindings.TtsVoiceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *TtsVoice) FreeMembers(recursive bool) {
	js.Free(
		p.EventTypes.Ref(),
		p.ExtensionId.Ref(),
		p.Lang.Ref(),
		p.VoiceName.Ref(),
	)
	p.EventTypes = p.EventTypes.FromRef(js.Undefined)
	p.ExtensionId = p.ExtensionId.FromRef(js.Undefined)
	p.Lang = p.Lang.FromRef(js.Undefined)
	p.VoiceName = p.VoiceName.FromRef(js.Undefined)
}

// HasFuncGetVoices returns true if the function "WEBEXT.tts.getVoices" exists.
func HasFuncGetVoices() bool {
	return js.True == bindings.HasFuncGetVoices()
}

// FuncGetVoices returns the function "WEBEXT.tts.getVoices".
func FuncGetVoices() (fn js.Func[func() js.Promise[js.Array[TtsVoice]]]) {
	bindings.FuncGetVoices(
		js.Pointer(&fn),
	)
	return
}

// GetVoices calls the function "WEBEXT.tts.getVoices" directly.
func GetVoices() (ret js.Promise[js.Array[TtsVoice]]) {
	bindings.CallGetVoices(
		js.Pointer(&ret),
	)

	return
}

// TryGetVoices calls the function "WEBEXT.tts.getVoices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetVoices() (ret js.Promise[js.Array[TtsVoice]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetVoices(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsSpeaking returns true if the function "WEBEXT.tts.isSpeaking" exists.
func HasFuncIsSpeaking() bool {
	return js.True == bindings.HasFuncIsSpeaking()
}

// FuncIsSpeaking returns the function "WEBEXT.tts.isSpeaking".
func FuncIsSpeaking() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsSpeaking(
		js.Pointer(&fn),
	)
	return
}

// IsSpeaking calls the function "WEBEXT.tts.isSpeaking" directly.
func IsSpeaking() (ret js.Promise[js.Boolean]) {
	bindings.CallIsSpeaking(
		js.Pointer(&ret),
	)

	return
}

// TryIsSpeaking calls the function "WEBEXT.tts.isSpeaking"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsSpeaking() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsSpeaking(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnEventEventCallbackFunc func(this js.Ref, event *TtsEvent) js.Ref

func (fn OnEventEventCallbackFunc) Register() js.Func[func(event *TtsEvent)] {
	return js.RegisterCallback[func(event *TtsEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TtsEvent
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

type OnEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event *TtsEvent) js.Ref
	Arg T
}

func (cb *OnEventEventCallback[T]) Register() js.Func[func(event *TtsEvent)] {
	return js.RegisterCallback[func(event *TtsEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 TtsEvent
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

// HasFuncOnEvent returns true if the function "WEBEXT.tts.onEvent.addListener" exists.
func HasFuncOnEvent() bool {
	return js.True == bindings.HasFuncOnEvent()
}

// FuncOnEvent returns the function "WEBEXT.tts.onEvent.addListener".
func FuncOnEvent() (fn js.Func[func(callback js.Func[func(event *TtsEvent)])]) {
	bindings.FuncOnEvent(
		js.Pointer(&fn),
	)
	return
}

// OnEvent calls the function "WEBEXT.tts.onEvent.addListener" directly.
func OnEvent(callback js.Func[func(event *TtsEvent)]) (ret js.Void) {
	bindings.CallOnEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnEvent calls the function "WEBEXT.tts.onEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnEvent(callback js.Func[func(event *TtsEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffEvent returns true if the function "WEBEXT.tts.onEvent.removeListener" exists.
func HasFuncOffEvent() bool {
	return js.True == bindings.HasFuncOffEvent()
}

// FuncOffEvent returns the function "WEBEXT.tts.onEvent.removeListener".
func FuncOffEvent() (fn js.Func[func(callback js.Func[func(event *TtsEvent)])]) {
	bindings.FuncOffEvent(
		js.Pointer(&fn),
	)
	return
}

// OffEvent calls the function "WEBEXT.tts.onEvent.removeListener" directly.
func OffEvent(callback js.Func[func(event *TtsEvent)]) (ret js.Void) {
	bindings.CallOffEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffEvent calls the function "WEBEXT.tts.onEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffEvent(callback js.Func[func(event *TtsEvent)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnEvent returns true if the function "WEBEXT.tts.onEvent.hasListener" exists.
func HasFuncHasOnEvent() bool {
	return js.True == bindings.HasFuncHasOnEvent()
}

// FuncHasOnEvent returns the function "WEBEXT.tts.onEvent.hasListener".
func FuncHasOnEvent() (fn js.Func[func(callback js.Func[func(event *TtsEvent)]) bool]) {
	bindings.FuncHasOnEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnEvent calls the function "WEBEXT.tts.onEvent.hasListener" directly.
func HasOnEvent(callback js.Func[func(event *TtsEvent)]) (ret bool) {
	bindings.CallHasOnEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnEvent calls the function "WEBEXT.tts.onEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnEvent(callback js.Func[func(event *TtsEvent)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncPause returns true if the function "WEBEXT.tts.pause" exists.
func HasFuncPause() bool {
	return js.True == bindings.HasFuncPause()
}

// FuncPause returns the function "WEBEXT.tts.pause".
func FuncPause() (fn js.Func[func()]) {
	bindings.FuncPause(
		js.Pointer(&fn),
	)
	return
}

// Pause calls the function "WEBEXT.tts.pause" directly.
func Pause() (ret js.Void) {
	bindings.CallPause(
		js.Pointer(&ret),
	)

	return
}

// TryPause calls the function "WEBEXT.tts.pause"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPause() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPause(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncResume returns true if the function "WEBEXT.tts.resume" exists.
func HasFuncResume() bool {
	return js.True == bindings.HasFuncResume()
}

// FuncResume returns the function "WEBEXT.tts.resume".
func FuncResume() (fn js.Func[func()]) {
	bindings.FuncResume(
		js.Pointer(&fn),
	)
	return
}

// Resume calls the function "WEBEXT.tts.resume" directly.
func Resume() (ret js.Void) {
	bindings.CallResume(
		js.Pointer(&ret),
	)

	return
}

// TryResume calls the function "WEBEXT.tts.resume"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryResume() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryResume(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSpeak returns true if the function "WEBEXT.tts.speak" exists.
func HasFuncSpeak() bool {
	return js.True == bindings.HasFuncSpeak()
}

// FuncSpeak returns the function "WEBEXT.tts.speak".
func FuncSpeak() (fn js.Func[func(utterance js.String, options TtsOptions) js.Promise[js.Void]]) {
	bindings.FuncSpeak(
		js.Pointer(&fn),
	)
	return
}

// Speak calls the function "WEBEXT.tts.speak" directly.
func Speak(utterance js.String, options TtsOptions) (ret js.Promise[js.Void]) {
	bindings.CallSpeak(
		js.Pointer(&ret),
		utterance.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySpeak calls the function "WEBEXT.tts.speak"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySpeak(utterance js.String, options TtsOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeak(
		js.Pointer(&ret), js.Pointer(&exception),
		utterance.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncStop returns true if the function "WEBEXT.tts.stop" exists.
func HasFuncStop() bool {
	return js.True == bindings.HasFuncStop()
}

// FuncStop returns the function "WEBEXT.tts.stop".
func FuncStop() (fn js.Func[func()]) {
	bindings.FuncStop(
		js.Pointer(&fn),
	)
	return
}

// Stop calls the function "WEBEXT.tts.stop" directly.
func Stop() (ret js.Void) {
	bindings.CallStop(
		js.Pointer(&ret),
	)

	return
}

// TryStop calls the function "WEBEXT.tts.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStop(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

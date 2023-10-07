// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package ttsengine

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/tts"
	"github.com/primecitizens/pcz/std/plat/js/webext/ttsengine/bindings"
)

type AudioBuffer struct {
	// AudioBuffer is "AudioBuffer.audioBuffer"
	//
	// Required
	AudioBuffer js.TypedArray[uint8]
	// CharIndex is "AudioBuffer.charIndex"
	//
	// Optional
	//
	// NOTE: FFI_USE_CharIndex MUST be set to true to make this field effective.
	CharIndex int64
	// IsLastBuffer is "AudioBuffer.isLastBuffer"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsLastBuffer MUST be set to true to make this field effective.
	IsLastBuffer bool

	FFI_USE_CharIndex    bool // for CharIndex.
	FFI_USE_IsLastBuffer bool // for IsLastBuffer.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioBuffer with all fields set.
func (p AudioBuffer) FromRef(ref js.Ref) AudioBuffer {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioBuffer in the application heap.
func (p AudioBuffer) New() js.Ref {
	return bindings.AudioBufferJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioBuffer) UpdateFrom(ref js.Ref) {
	bindings.AudioBufferJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioBuffer) Update(ref js.Ref) {
	bindings.AudioBufferJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioBuffer) FreeMembers(recursive bool) {
	js.Free(
		p.AudioBuffer.Ref(),
	)
	p.AudioBuffer = p.AudioBuffer.FromRef(js.Undefined)
}

type AudioStreamOptions struct {
	// BufferSize is "AudioStreamOptions.bufferSize"
	//
	// Required
	BufferSize int64
	// SampleRate is "AudioStreamOptions.sampleRate"
	//
	// Required
	SampleRate int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioStreamOptions with all fields set.
func (p AudioStreamOptions) FromRef(ref js.Ref) AudioStreamOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioStreamOptions in the application heap.
func (p AudioStreamOptions) New() js.Ref {
	return bindings.AudioStreamOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AudioStreamOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioStreamOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioStreamOptions) Update(ref js.Ref) {
	bindings.AudioStreamOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioStreamOptions) FreeMembers(recursive bool) {
}

type OnSpeakArgSendTtsEventFunc func(this js.Ref, event *tts.TtsEvent) js.Ref

func (fn OnSpeakArgSendTtsEventFunc) Register() js.Func[func(event *tts.TtsEvent)] {
	return js.RegisterCallback[func(event *tts.TtsEvent)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSpeakArgSendTtsEventFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 tts.TtsEvent
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

type OnSpeakArgSendTtsEvent[T any] struct {
	Fn  func(arg T, this js.Ref, event *tts.TtsEvent) js.Ref
	Arg T
}

func (cb *OnSpeakArgSendTtsEvent[T]) Register() js.Func[func(event *tts.TtsEvent)] {
	return js.RegisterCallback[func(event *tts.TtsEvent)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSpeakArgSendTtsEvent[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 tts.TtsEvent
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

type OnSpeakWithAudioStreamArgSendErrorFunc func(this js.Ref, errorMessage js.String) js.Ref

func (fn OnSpeakWithAudioStreamArgSendErrorFunc) Register() js.Func[func(errorMessage js.String)] {
	return js.RegisterCallback[func(errorMessage js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSpeakWithAudioStreamArgSendErrorFunc) DispatchCallback(
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

type OnSpeakWithAudioStreamArgSendError[T any] struct {
	Fn  func(arg T, this js.Ref, errorMessage js.String) js.Ref
	Arg T
}

func (cb *OnSpeakWithAudioStreamArgSendError[T]) Register() js.Func[func(errorMessage js.String)] {
	return js.RegisterCallback[func(errorMessage js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSpeakWithAudioStreamArgSendError[T]) DispatchCallback(
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

type OnSpeakWithAudioStreamArgSendTtsAudioFunc func(this js.Ref, audioBufferParams *AudioBuffer) js.Ref

func (fn OnSpeakWithAudioStreamArgSendTtsAudioFunc) Register() js.Func[func(audioBufferParams *AudioBuffer)] {
	return js.RegisterCallback[func(audioBufferParams *AudioBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSpeakWithAudioStreamArgSendTtsAudioFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AudioBuffer
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

type OnSpeakWithAudioStreamArgSendTtsAudio[T any] struct {
	Fn  func(arg T, this js.Ref, audioBufferParams *AudioBuffer) js.Ref
	Arg T
}

func (cb *OnSpeakWithAudioStreamArgSendTtsAudio[T]) Register() js.Func[func(audioBufferParams *AudioBuffer)] {
	return js.RegisterCallback[func(audioBufferParams *AudioBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSpeakWithAudioStreamArgSendTtsAudio[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 AudioBuffer
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

type SpeakOptions struct {
	// Gender is "SpeakOptions.gender"
	//
	// Optional
	Gender VoiceGender
	// Lang is "SpeakOptions.lang"
	//
	// Optional
	Lang js.String
	// Pitch is "SpeakOptions.pitch"
	//
	// Optional
	//
	// NOTE: FFI_USE_Pitch MUST be set to true to make this field effective.
	Pitch float64
	// Rate is "SpeakOptions.rate"
	//
	// Optional
	//
	// NOTE: FFI_USE_Rate MUST be set to true to make this field effective.
	Rate float64
	// VoiceName is "SpeakOptions.voiceName"
	//
	// Optional
	VoiceName js.String
	// Volume is "SpeakOptions.volume"
	//
	// Optional
	//
	// NOTE: FFI_USE_Volume MUST be set to true to make this field effective.
	Volume float64

	FFI_USE_Pitch  bool // for Pitch.
	FFI_USE_Rate   bool // for Rate.
	FFI_USE_Volume bool // for Volume.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SpeakOptions with all fields set.
func (p SpeakOptions) FromRef(ref js.Ref) SpeakOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SpeakOptions in the application heap.
func (p SpeakOptions) New() js.Ref {
	return bindings.SpeakOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SpeakOptions) UpdateFrom(ref js.Ref) {
	bindings.SpeakOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SpeakOptions) Update(ref js.Ref) {
	bindings.SpeakOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SpeakOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Lang.Ref(),
		p.VoiceName.Ref(),
	)
	p.Lang = p.Lang.FromRef(js.Undefined)
	p.VoiceName = p.VoiceName.FromRef(js.Undefined)
}

type OnPauseEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnPauseEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPauseEventCallbackFunc) DispatchCallback(
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

type OnPauseEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnPauseEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPauseEventCallback[T]) DispatchCallback(
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

// HasFuncOnPause returns true if the function "WEBEXT.ttsEngine.onPause.addListener" exists.
func HasFuncOnPause() bool {
	return js.True == bindings.HasFuncOnPause()
}

// FuncOnPause returns the function "WEBEXT.ttsEngine.onPause.addListener".
func FuncOnPause() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnPause(
		js.Pointer(&fn),
	)
	return
}

// OnPause calls the function "WEBEXT.ttsEngine.onPause.addListener" directly.
func OnPause(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnPause(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPause calls the function "WEBEXT.ttsEngine.onPause.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPause(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPause(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPause returns true if the function "WEBEXT.ttsEngine.onPause.removeListener" exists.
func HasFuncOffPause() bool {
	return js.True == bindings.HasFuncOffPause()
}

// FuncOffPause returns the function "WEBEXT.ttsEngine.onPause.removeListener".
func FuncOffPause() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffPause(
		js.Pointer(&fn),
	)
	return
}

// OffPause calls the function "WEBEXT.ttsEngine.onPause.removeListener" directly.
func OffPause(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffPause(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPause calls the function "WEBEXT.ttsEngine.onPause.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPause(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPause(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPause returns true if the function "WEBEXT.ttsEngine.onPause.hasListener" exists.
func HasFuncHasOnPause() bool {
	return js.True == bindings.HasFuncHasOnPause()
}

// FuncHasOnPause returns the function "WEBEXT.ttsEngine.onPause.hasListener".
func FuncHasOnPause() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnPause(
		js.Pointer(&fn),
	)
	return
}

// HasOnPause calls the function "WEBEXT.ttsEngine.onPause.hasListener" directly.
func HasOnPause(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnPause(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPause calls the function "WEBEXT.ttsEngine.onPause.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPause(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPause(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnResumeEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnResumeEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnResumeEventCallbackFunc) DispatchCallback(
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

type OnResumeEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnResumeEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnResumeEventCallback[T]) DispatchCallback(
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

// HasFuncOnResume returns true if the function "WEBEXT.ttsEngine.onResume.addListener" exists.
func HasFuncOnResume() bool {
	return js.True == bindings.HasFuncOnResume()
}

// FuncOnResume returns the function "WEBEXT.ttsEngine.onResume.addListener".
func FuncOnResume() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnResume(
		js.Pointer(&fn),
	)
	return
}

// OnResume calls the function "WEBEXT.ttsEngine.onResume.addListener" directly.
func OnResume(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnResume(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnResume calls the function "WEBEXT.ttsEngine.onResume.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnResume(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnResume(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffResume returns true if the function "WEBEXT.ttsEngine.onResume.removeListener" exists.
func HasFuncOffResume() bool {
	return js.True == bindings.HasFuncOffResume()
}

// FuncOffResume returns the function "WEBEXT.ttsEngine.onResume.removeListener".
func FuncOffResume() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffResume(
		js.Pointer(&fn),
	)
	return
}

// OffResume calls the function "WEBEXT.ttsEngine.onResume.removeListener" directly.
func OffResume(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffResume(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffResume calls the function "WEBEXT.ttsEngine.onResume.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffResume(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffResume(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnResume returns true if the function "WEBEXT.ttsEngine.onResume.hasListener" exists.
func HasFuncHasOnResume() bool {
	return js.True == bindings.HasFuncHasOnResume()
}

// FuncHasOnResume returns the function "WEBEXT.ttsEngine.onResume.hasListener".
func FuncHasOnResume() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnResume(
		js.Pointer(&fn),
	)
	return
}

// HasOnResume calls the function "WEBEXT.ttsEngine.onResume.hasListener" directly.
func HasOnResume(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnResume(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnResume calls the function "WEBEXT.ttsEngine.onResume.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnResume(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnResume(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSpeakEventCallbackFunc func(this js.Ref, utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)]) js.Ref

func (fn OnSpeakEventCallbackFunc) Register() js.Func[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])] {
	return js.RegisterCallback[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSpeakEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 SpeakOptions
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func(event *tts.TtsEvent)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSpeakEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)]) js.Ref
	Arg T
}

func (cb *OnSpeakEventCallback[T]) Register() js.Func[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])] {
	return js.RegisterCallback[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSpeakEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 SpeakOptions
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func(event *tts.TtsEvent)]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSpeak returns true if the function "WEBEXT.ttsEngine.onSpeak.addListener" exists.
func HasFuncOnSpeak() bool {
	return js.True == bindings.HasFuncOnSpeak()
}

// FuncOnSpeak returns the function "WEBEXT.ttsEngine.onSpeak.addListener".
func FuncOnSpeak() (fn js.Func[func(callback js.Func[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])])]) {
	bindings.FuncOnSpeak(
		js.Pointer(&fn),
	)
	return
}

// OnSpeak calls the function "WEBEXT.ttsEngine.onSpeak.addListener" directly.
func OnSpeak(callback js.Func[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])]) (ret js.Void) {
	bindings.CallOnSpeak(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSpeak calls the function "WEBEXT.ttsEngine.onSpeak.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSpeak(callback js.Func[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSpeak(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSpeak returns true if the function "WEBEXT.ttsEngine.onSpeak.removeListener" exists.
func HasFuncOffSpeak() bool {
	return js.True == bindings.HasFuncOffSpeak()
}

// FuncOffSpeak returns the function "WEBEXT.ttsEngine.onSpeak.removeListener".
func FuncOffSpeak() (fn js.Func[func(callback js.Func[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])])]) {
	bindings.FuncOffSpeak(
		js.Pointer(&fn),
	)
	return
}

// OffSpeak calls the function "WEBEXT.ttsEngine.onSpeak.removeListener" directly.
func OffSpeak(callback js.Func[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])]) (ret js.Void) {
	bindings.CallOffSpeak(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSpeak calls the function "WEBEXT.ttsEngine.onSpeak.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSpeak(callback js.Func[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSpeak(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSpeak returns true if the function "WEBEXT.ttsEngine.onSpeak.hasListener" exists.
func HasFuncHasOnSpeak() bool {
	return js.True == bindings.HasFuncHasOnSpeak()
}

// FuncHasOnSpeak returns the function "WEBEXT.ttsEngine.onSpeak.hasListener".
func FuncHasOnSpeak() (fn js.Func[func(callback js.Func[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])]) bool]) {
	bindings.FuncHasOnSpeak(
		js.Pointer(&fn),
	)
	return
}

// HasOnSpeak calls the function "WEBEXT.ttsEngine.onSpeak.hasListener" directly.
func HasOnSpeak(callback js.Func[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])]) (ret bool) {
	bindings.CallHasOnSpeak(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSpeak calls the function "WEBEXT.ttsEngine.onSpeak.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSpeak(callback js.Func[func(utterance js.String, options *SpeakOptions, sendTtsEvent js.Func[func(event *tts.TtsEvent)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSpeak(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSpeakWithAudioStreamEventCallbackFunc func(this js.Ref, utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)]) js.Ref

func (fn OnSpeakWithAudioStreamEventCallbackFunc) Register() js.Func[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])] {
	return js.RegisterCallback[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSpeakWithAudioStreamEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 5+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 SpeakOptions
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)
	var arg2 AudioStreamOptions
	arg2.UpdateFrom(args[2+1])
	defer arg2.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		mark.NoEscape(&arg2),
		js.Func[func(audioBufferParams *AudioBuffer)]{}.FromRef(args[3+1]),
		js.Func[func(errorMessage js.String)]{}.FromRef(args[4+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnSpeakWithAudioStreamEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)]) js.Ref
	Arg T
}

func (cb *OnSpeakWithAudioStreamEventCallback[T]) Register() js.Func[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])] {
	return js.RegisterCallback[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSpeakWithAudioStreamEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 5+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 SpeakOptions
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)
	var arg2 AudioStreamOptions
	arg2.UpdateFrom(args[2+1])
	defer arg2.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		mark.NoEscape(&arg2),
		js.Func[func(audioBufferParams *AudioBuffer)]{}.FromRef(args[3+1]),
		js.Func[func(errorMessage js.String)]{}.FromRef(args[4+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSpeakWithAudioStream returns true if the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.addListener" exists.
func HasFuncOnSpeakWithAudioStream() bool {
	return js.True == bindings.HasFuncOnSpeakWithAudioStream()
}

// FuncOnSpeakWithAudioStream returns the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.addListener".
func FuncOnSpeakWithAudioStream() (fn js.Func[func(callback js.Func[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])])]) {
	bindings.FuncOnSpeakWithAudioStream(
		js.Pointer(&fn),
	)
	return
}

// OnSpeakWithAudioStream calls the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.addListener" directly.
func OnSpeakWithAudioStream(callback js.Func[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])]) (ret js.Void) {
	bindings.CallOnSpeakWithAudioStream(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSpeakWithAudioStream calls the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSpeakWithAudioStream(callback js.Func[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSpeakWithAudioStream(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSpeakWithAudioStream returns true if the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.removeListener" exists.
func HasFuncOffSpeakWithAudioStream() bool {
	return js.True == bindings.HasFuncOffSpeakWithAudioStream()
}

// FuncOffSpeakWithAudioStream returns the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.removeListener".
func FuncOffSpeakWithAudioStream() (fn js.Func[func(callback js.Func[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])])]) {
	bindings.FuncOffSpeakWithAudioStream(
		js.Pointer(&fn),
	)
	return
}

// OffSpeakWithAudioStream calls the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.removeListener" directly.
func OffSpeakWithAudioStream(callback js.Func[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])]) (ret js.Void) {
	bindings.CallOffSpeakWithAudioStream(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSpeakWithAudioStream calls the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSpeakWithAudioStream(callback js.Func[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSpeakWithAudioStream(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSpeakWithAudioStream returns true if the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.hasListener" exists.
func HasFuncHasOnSpeakWithAudioStream() bool {
	return js.True == bindings.HasFuncHasOnSpeakWithAudioStream()
}

// FuncHasOnSpeakWithAudioStream returns the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.hasListener".
func FuncHasOnSpeakWithAudioStream() (fn js.Func[func(callback js.Func[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])]) bool]) {
	bindings.FuncHasOnSpeakWithAudioStream(
		js.Pointer(&fn),
	)
	return
}

// HasOnSpeakWithAudioStream calls the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.hasListener" directly.
func HasOnSpeakWithAudioStream(callback js.Func[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])]) (ret bool) {
	bindings.CallHasOnSpeakWithAudioStream(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSpeakWithAudioStream calls the function "WEBEXT.ttsEngine.onSpeakWithAudioStream.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSpeakWithAudioStream(callback js.Func[func(utterance js.String, options *SpeakOptions, audioStreamOptions *AudioStreamOptions, sendTtsAudio js.Func[func(audioBufferParams *AudioBuffer)], sendError js.Func[func(errorMessage js.String)])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSpeakWithAudioStream(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnStopEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnStopEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnStopEventCallbackFunc) DispatchCallback(
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

type OnStopEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnStopEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnStopEventCallback[T]) DispatchCallback(
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

// HasFuncOnStop returns true if the function "WEBEXT.ttsEngine.onStop.addListener" exists.
func HasFuncOnStop() bool {
	return js.True == bindings.HasFuncOnStop()
}

// FuncOnStop returns the function "WEBEXT.ttsEngine.onStop.addListener".
func FuncOnStop() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnStop(
		js.Pointer(&fn),
	)
	return
}

// OnStop calls the function "WEBEXT.ttsEngine.onStop.addListener" directly.
func OnStop(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnStop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnStop calls the function "WEBEXT.ttsEngine.onStop.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnStop(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnStop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffStop returns true if the function "WEBEXT.ttsEngine.onStop.removeListener" exists.
func HasFuncOffStop() bool {
	return js.True == bindings.HasFuncOffStop()
}

// FuncOffStop returns the function "WEBEXT.ttsEngine.onStop.removeListener".
func FuncOffStop() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffStop(
		js.Pointer(&fn),
	)
	return
}

// OffStop calls the function "WEBEXT.ttsEngine.onStop.removeListener" directly.
func OffStop(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffStop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffStop calls the function "WEBEXT.ttsEngine.onStop.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffStop(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffStop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnStop returns true if the function "WEBEXT.ttsEngine.onStop.hasListener" exists.
func HasFuncHasOnStop() bool {
	return js.True == bindings.HasFuncHasOnStop()
}

// FuncHasOnStop returns the function "WEBEXT.ttsEngine.onStop.hasListener".
func FuncHasOnStop() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnStop(
		js.Pointer(&fn),
	)
	return
}

// HasOnStop calls the function "WEBEXT.ttsEngine.onStop.hasListener" directly.
func HasOnStop(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnStop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnStop calls the function "WEBEXT.ttsEngine.onStop.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnStop(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnStop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSendTtsAudio returns true if the function "WEBEXT.ttsEngine.sendTtsAudio" exists.
func HasFuncSendTtsAudio() bool {
	return js.True == bindings.HasFuncSendTtsAudio()
}

// FuncSendTtsAudio returns the function "WEBEXT.ttsEngine.sendTtsAudio".
func FuncSendTtsAudio() (fn js.Func[func(requestId int64, audio AudioBuffer)]) {
	bindings.FuncSendTtsAudio(
		js.Pointer(&fn),
	)
	return
}

// SendTtsAudio calls the function "WEBEXT.ttsEngine.sendTtsAudio" directly.
func SendTtsAudio(requestId int64, audio AudioBuffer) (ret js.Void) {
	bindings.CallSendTtsAudio(
		js.Pointer(&ret),
		float64(requestId),
		js.Pointer(&audio),
	)

	return
}

// TrySendTtsAudio calls the function "WEBEXT.ttsEngine.sendTtsAudio"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendTtsAudio(requestId int64, audio AudioBuffer) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendTtsAudio(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(requestId),
		js.Pointer(&audio),
	)

	return
}

// HasFuncSendTtsEvent returns true if the function "WEBEXT.ttsEngine.sendTtsEvent" exists.
func HasFuncSendTtsEvent() bool {
	return js.True == bindings.HasFuncSendTtsEvent()
}

// FuncSendTtsEvent returns the function "WEBEXT.ttsEngine.sendTtsEvent".
func FuncSendTtsEvent() (fn js.Func[func(requestId int64, event tts.TtsEvent)]) {
	bindings.FuncSendTtsEvent(
		js.Pointer(&fn),
	)
	return
}

// SendTtsEvent calls the function "WEBEXT.ttsEngine.sendTtsEvent" directly.
func SendTtsEvent(requestId int64, event tts.TtsEvent) (ret js.Void) {
	bindings.CallSendTtsEvent(
		js.Pointer(&ret),
		float64(requestId),
		js.Pointer(&event),
	)

	return
}

// TrySendTtsEvent calls the function "WEBEXT.ttsEngine.sendTtsEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendTtsEvent(requestId int64, event tts.TtsEvent) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendTtsEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(requestId),
		js.Pointer(&event),
	)

	return
}

// HasFuncUpdateVoices returns true if the function "WEBEXT.ttsEngine.updateVoices" exists.
func HasFuncUpdateVoices() bool {
	return js.True == bindings.HasFuncUpdateVoices()
}

// FuncUpdateVoices returns the function "WEBEXT.ttsEngine.updateVoices".
func FuncUpdateVoices() (fn js.Func[func(voices js.Array[tts.TtsVoice])]) {
	bindings.FuncUpdateVoices(
		js.Pointer(&fn),
	)
	return
}

// UpdateVoices calls the function "WEBEXT.ttsEngine.updateVoices" directly.
func UpdateVoices(voices js.Array[tts.TtsVoice]) (ret js.Void) {
	bindings.CallUpdateVoices(
		js.Pointer(&ret),
		voices.Ref(),
	)

	return
}

// TryUpdateVoices calls the function "WEBEXT.ttsEngine.updateVoices"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateVoices(voices js.Array[tts.TtsVoice]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateVoices(
		js.Pointer(&ret), js.Pointer(&exception),
		voices.Ref(),
	)

	return
}

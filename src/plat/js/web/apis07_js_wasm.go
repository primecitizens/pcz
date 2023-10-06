// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type DecodeSuccessCallbackFunc func(this js.Ref, decodedData AudioBuffer) js.Ref

func (fn DecodeSuccessCallbackFunc) Register() js.Func[func(decodedData AudioBuffer)] {
	return js.RegisterCallback[func(decodedData AudioBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DecodeSuccessCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		AudioBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DecodeSuccessCallback[T any] struct {
	Fn  func(arg T, this js.Ref, decodedData AudioBuffer) js.Ref
	Arg T
}

func (cb *DecodeSuccessCallback[T]) Register() js.Func[func(decodedData AudioBuffer)] {
	return js.RegisterCallback[func(decodedData AudioBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DecodeSuccessCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		AudioBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DecodeErrorCallbackFunc func(this js.Ref, err DOMException) js.Ref

func (fn DecodeErrorCallbackFunc) Register() js.Func[func(err DOMException)] {
	return js.RegisterCallback[func(err DOMException)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DecodeErrorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		DOMException{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DecodeErrorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, err DOMException) js.Ref
	Arg T
}

func (cb *DecodeErrorCallback[T]) Register() js.Func[func(err DOMException)] {
	return js.RegisterCallback[func(err DOMException)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DecodeErrorCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		DOMException{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

const (
	DOMException_INDEX_SIZE_ERR              uint16 = 1
	DOMException_DOMSTRING_SIZE_ERR          uint16 = 2
	DOMException_HIERARCHY_REQUEST_ERR       uint16 = 3
	DOMException_WRONG_DOCUMENT_ERR          uint16 = 4
	DOMException_INVALID_CHARACTER_ERR       uint16 = 5
	DOMException_NO_DATA_ALLOWED_ERR         uint16 = 6
	DOMException_NO_MODIFICATION_ALLOWED_ERR uint16 = 7
	DOMException_NOT_FOUND_ERR               uint16 = 8
	DOMException_NOT_SUPPORTED_ERR           uint16 = 9
	DOMException_INUSE_ATTRIBUTE_ERR         uint16 = 10
	DOMException_INVALID_STATE_ERR           uint16 = 11
	DOMException_SYNTAX_ERR                  uint16 = 12
	DOMException_INVALID_MODIFICATION_ERR    uint16 = 13
	DOMException_NAMESPACE_ERR               uint16 = 14
	DOMException_INVALID_ACCESS_ERR          uint16 = 15
	DOMException_VALIDATION_ERR              uint16 = 16
	DOMException_TYPE_MISMATCH_ERR           uint16 = 17
	DOMException_SECURITY_ERR                uint16 = 18
	DOMException_NETWORK_ERR                 uint16 = 19
	DOMException_ABORT_ERR                   uint16 = 20
	DOMException_URL_MISMATCH_ERR            uint16 = 21
	DOMException_QUOTA_EXCEEDED_ERR          uint16 = 22
	DOMException_TIMEOUT_ERR                 uint16 = 23
	DOMException_INVALID_NODE_TYPE_ERR       uint16 = 24
	DOMException_DATA_CLONE_ERR              uint16 = 25
)

func NewDOMException(message js.String, name js.String) (ret DOMException) {
	ret.ref = bindings.NewDOMExceptionByDOMException(
		message.Ref(),
		name.Ref())
	return
}

func NewDOMExceptionByDOMException1(message js.String) (ret DOMException) {
	ret.ref = bindings.NewDOMExceptionByDOMException1(
		message.Ref())
	return
}

func NewDOMExceptionByDOMException2() (ret DOMException) {
	ret.ref = bindings.NewDOMExceptionByDOMException2()
	return
}

type DOMException struct {
	ref js.Ref
}

func (this DOMException) Once() DOMException {
	this.Ref().Once()
	return this
}

func (this DOMException) Ref() js.Ref {
	return this.ref
}

func (this DOMException) FromRef(ref js.Ref) DOMException {
	this.ref = ref
	return this
}

func (this DOMException) Free() {
	this.Ref().Free()
}

// Name returns the value of property "DOMException.name".
//
// It returns ok=false if there is no such property.
func (this DOMException) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDOMExceptionName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "DOMException.message".
//
// It returns ok=false if there is no such property.
func (this DOMException) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDOMExceptionMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Code returns the value of property "DOMException.code".
//
// It returns ok=false if there is no such property.
func (this DOMException) Code() (ret uint16, ok bool) {
	ok = js.True == bindings.GetDOMExceptionCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type AudioDestinationNode struct {
	AudioNode
}

func (this AudioDestinationNode) Once() AudioDestinationNode {
	this.Ref().Once()
	return this
}

func (this AudioDestinationNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this AudioDestinationNode) FromRef(ref js.Ref) AudioDestinationNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this AudioDestinationNode) Free() {
	this.Ref().Free()
}

// MaxChannelCount returns the value of property "AudioDestinationNode.maxChannelCount".
//
// It returns ok=false if there is no such property.
func (this AudioDestinationNode) MaxChannelCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioDestinationNodeMaxChannelCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type AudioListener struct {
	ref js.Ref
}

func (this AudioListener) Once() AudioListener {
	this.Ref().Once()
	return this
}

func (this AudioListener) Ref() js.Ref {
	return this.ref
}

func (this AudioListener) FromRef(ref js.Ref) AudioListener {
	this.ref = ref
	return this
}

func (this AudioListener) Free() {
	this.Ref().Free()
}

// PositionX returns the value of property "AudioListener.positionX".
//
// It returns ok=false if there is no such property.
func (this AudioListener) PositionX() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerPositionX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PositionY returns the value of property "AudioListener.positionY".
//
// It returns ok=false if there is no such property.
func (this AudioListener) PositionY() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerPositionY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PositionZ returns the value of property "AudioListener.positionZ".
//
// It returns ok=false if there is no such property.
func (this AudioListener) PositionZ() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerPositionZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ForwardX returns the value of property "AudioListener.forwardX".
//
// It returns ok=false if there is no such property.
func (this AudioListener) ForwardX() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerForwardX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ForwardY returns the value of property "AudioListener.forwardY".
//
// It returns ok=false if there is no such property.
func (this AudioListener) ForwardY() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerForwardY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ForwardZ returns the value of property "AudioListener.forwardZ".
//
// It returns ok=false if there is no such property.
func (this AudioListener) ForwardZ() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerForwardZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UpX returns the value of property "AudioListener.upX".
//
// It returns ok=false if there is no such property.
func (this AudioListener) UpX() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerUpX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UpY returns the value of property "AudioListener.upY".
//
// It returns ok=false if there is no such property.
func (this AudioListener) UpY() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerUpY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UpZ returns the value of property "AudioListener.upZ".
//
// It returns ok=false if there is no such property.
func (this AudioListener) UpZ() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerUpZ(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSetPosition returns true if the method "AudioListener.setPosition" exists.
func (this AudioListener) HasSetPosition() bool {
	return js.True == bindings.HasAudioListenerSetPosition(
		this.Ref(),
	)
}

// SetPositionFunc returns the method "AudioListener.setPosition".
func (this AudioListener) SetPositionFunc() (fn js.Func[func(x float32, y float32, z float32)]) {
	return fn.FromRef(
		bindings.AudioListenerSetPositionFunc(
			this.Ref(),
		),
	)
}

// SetPosition calls the method "AudioListener.setPosition".
func (this AudioListener) SetPosition(x float32, y float32, z float32) (ret js.Void) {
	bindings.CallAudioListenerSetPosition(
		this.Ref(), js.Pointer(&ret),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// TrySetPosition calls the method "AudioListener.setPosition"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioListener) TrySetPosition(x float32, y float32, z float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioListenerSetPosition(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// HasSetOrientation returns true if the method "AudioListener.setOrientation" exists.
func (this AudioListener) HasSetOrientation() bool {
	return js.True == bindings.HasAudioListenerSetOrientation(
		this.Ref(),
	)
}

// SetOrientationFunc returns the method "AudioListener.setOrientation".
func (this AudioListener) SetOrientationFunc() (fn js.Func[func(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32)]) {
	return fn.FromRef(
		bindings.AudioListenerSetOrientationFunc(
			this.Ref(),
		),
	)
}

// SetOrientation calls the method "AudioListener.setOrientation".
func (this AudioListener) SetOrientation(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32) (ret js.Void) {
	bindings.CallAudioListenerSetOrientation(
		this.Ref(), js.Pointer(&ret),
		float32(x),
		float32(y),
		float32(z),
		float32(xUp),
		float32(yUp),
		float32(zUp),
	)

	return
}

// TrySetOrientation calls the method "AudioListener.setOrientation"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AudioListener) TrySetOrientation(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioListenerSetOrientation(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float32(x),
		float32(y),
		float32(z),
		float32(xUp),
		float32(yUp),
		float32(zUp),
	)

	return
}

type AudioContextState uint32

const (
	_ AudioContextState = iota

	AudioContextState_SUSPENDED
	AudioContextState_RUNNING
	AudioContextState_CLOSED
)

func (AudioContextState) FromRef(str js.Ref) AudioContextState {
	return AudioContextState(bindings.ConstOfAudioContextState(str))
}

func (x AudioContextState) String() (string, bool) {
	switch x {
	case AudioContextState_SUSPENDED:
		return "suspended", true
	case AudioContextState_RUNNING:
		return "running", true
	case AudioContextState_CLOSED:
		return "closed", true
	default:
		return "", false
	}
}

type StructuredSerializeOptions struct {
	// Transfer is "StructuredSerializeOptions.transfer"
	//
	// Optional, defaults to [].
	Transfer js.Array[js.Object]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StructuredSerializeOptions with all fields set.
func (p StructuredSerializeOptions) FromRef(ref js.Ref) StructuredSerializeOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StructuredSerializeOptions in the application heap.
func (p StructuredSerializeOptions) New() js.Ref {
	return bindings.StructuredSerializeOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p StructuredSerializeOptions) UpdateFrom(ref js.Ref) {
	bindings.StructuredSerializeOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p StructuredSerializeOptions) Update(ref js.Ref) {
	bindings.StructuredSerializeOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MessagePort struct {
	EventTarget
}

func (this MessagePort) Once() MessagePort {
	this.Ref().Once()
	return this
}

func (this MessagePort) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this MessagePort) FromRef(ref js.Ref) MessagePort {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this MessagePort) Free() {
	this.Ref().Free()
}

// HasPostMessage returns true if the method "MessagePort.postMessage" exists.
func (this MessagePort) HasPostMessage() bool {
	return js.True == bindings.HasMessagePortPostMessage(
		this.Ref(),
	)
}

// PostMessageFunc returns the method "MessagePort.postMessage".
func (this MessagePort) PostMessageFunc() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.MessagePortPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "MessagePort.postMessage".
func (this MessagePort) PostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void) {
	bindings.CallMessagePortPostMessage(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// TryPostMessage calls the method "MessagePort.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessagePort) TryPostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessagePortPostMessage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// HasPostMessage1 returns true if the method "MessagePort.postMessage" exists.
func (this MessagePort) HasPostMessage1() bool {
	return js.True == bindings.HasMessagePortPostMessage1(
		this.Ref(),
	)
}

// PostMessage1Func returns the method "MessagePort.postMessage".
func (this MessagePort) PostMessage1Func() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.MessagePortPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "MessagePort.postMessage".
func (this MessagePort) PostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void) {
	bindings.CallMessagePortPostMessage1(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage1 calls the method "MessagePort.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessagePort) TryPostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessagePortPostMessage1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasPostMessage2 returns true if the method "MessagePort.postMessage" exists.
func (this MessagePort) HasPostMessage2() bool {
	return js.True == bindings.HasMessagePortPostMessage2(
		this.Ref(),
	)
}

// PostMessage2Func returns the method "MessagePort.postMessage".
func (this MessagePort) PostMessage2Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.MessagePortPostMessage2Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "MessagePort.postMessage".
func (this MessagePort) PostMessage2(message js.Any) (ret js.Void) {
	bindings.CallMessagePortPostMessage2(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage2 calls the method "MessagePort.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessagePort) TryPostMessage2(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessagePortPostMessage2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasStart returns true if the method "MessagePort.start" exists.
func (this MessagePort) HasStart() bool {
	return js.True == bindings.HasMessagePortStart(
		this.Ref(),
	)
}

// StartFunc returns the method "MessagePort.start".
func (this MessagePort) StartFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MessagePortStartFunc(
			this.Ref(),
		),
	)
}

// Start calls the method "MessagePort.start".
func (this MessagePort) Start() (ret js.Void) {
	bindings.CallMessagePortStart(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStart calls the method "MessagePort.start"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessagePort) TryStart() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessagePortStart(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "MessagePort.close" exists.
func (this MessagePort) HasClose() bool {
	return js.True == bindings.HasMessagePortClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "MessagePort.close".
func (this MessagePort) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MessagePortCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "MessagePort.close".
func (this MessagePort) Close() (ret js.Void) {
	bindings.CallMessagePortClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "MessagePort.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessagePort) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessagePortClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type AudioWorklet struct {
	Worklet
}

func (this AudioWorklet) Once() AudioWorklet {
	this.Ref().Once()
	return this
}

func (this AudioWorklet) Ref() js.Ref {
	return this.Worklet.Ref()
}

func (this AudioWorklet) FromRef(ref js.Ref) AudioWorklet {
	this.Worklet = this.Worklet.FromRef(ref)
	return this
}

func (this AudioWorklet) Free() {
	this.Ref().Free()
}

// Port returns the value of property "AudioWorklet.port".
//
// It returns ok=false if there is no such property.
func (this AudioWorklet) Port() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetAudioWorkletPort(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type BaseAudioContext struct {
	EventTarget
}

func (this BaseAudioContext) Once() BaseAudioContext {
	this.Ref().Once()
	return this
}

func (this BaseAudioContext) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this BaseAudioContext) FromRef(ref js.Ref) BaseAudioContext {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this BaseAudioContext) Free() {
	this.Ref().Free()
}

// Destination returns the value of property "BaseAudioContext.destination".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) Destination() (ret AudioDestinationNode, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextDestination(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SampleRate returns the value of property "BaseAudioContext.sampleRate".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) SampleRate() (ret float32, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextSampleRate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CurrentTime returns the value of property "BaseAudioContext.currentTime".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) CurrentTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextCurrentTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Listener returns the value of property "BaseAudioContext.listener".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) Listener() (ret AudioListener, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextListener(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// State returns the value of property "BaseAudioContext.state".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) State() (ret AudioContextState, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AudioWorklet returns the value of property "BaseAudioContext.audioWorklet".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) AudioWorklet() (ret AudioWorklet, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextAudioWorklet(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasCreateAnalyser returns true if the method "BaseAudioContext.createAnalyser" exists.
func (this BaseAudioContext) HasCreateAnalyser() bool {
	return js.True == bindings.HasBaseAudioContextCreateAnalyser(
		this.Ref(),
	)
}

// CreateAnalyserFunc returns the method "BaseAudioContext.createAnalyser".
func (this BaseAudioContext) CreateAnalyserFunc() (fn js.Func[func() AnalyserNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateAnalyserFunc(
			this.Ref(),
		),
	)
}

// CreateAnalyser calls the method "BaseAudioContext.createAnalyser".
func (this BaseAudioContext) CreateAnalyser() (ret AnalyserNode) {
	bindings.CallBaseAudioContextCreateAnalyser(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateAnalyser calls the method "BaseAudioContext.createAnalyser"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateAnalyser() (ret AnalyserNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateAnalyser(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateBiquadFilter returns true if the method "BaseAudioContext.createBiquadFilter" exists.
func (this BaseAudioContext) HasCreateBiquadFilter() bool {
	return js.True == bindings.HasBaseAudioContextCreateBiquadFilter(
		this.Ref(),
	)
}

// CreateBiquadFilterFunc returns the method "BaseAudioContext.createBiquadFilter".
func (this BaseAudioContext) CreateBiquadFilterFunc() (fn js.Func[func() BiquadFilterNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateBiquadFilterFunc(
			this.Ref(),
		),
	)
}

// CreateBiquadFilter calls the method "BaseAudioContext.createBiquadFilter".
func (this BaseAudioContext) CreateBiquadFilter() (ret BiquadFilterNode) {
	bindings.CallBaseAudioContextCreateBiquadFilter(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateBiquadFilter calls the method "BaseAudioContext.createBiquadFilter"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateBiquadFilter() (ret BiquadFilterNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateBiquadFilter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateBuffer returns true if the method "BaseAudioContext.createBuffer" exists.
func (this BaseAudioContext) HasCreateBuffer() bool {
	return js.True == bindings.HasBaseAudioContextCreateBuffer(
		this.Ref(),
	)
}

// CreateBufferFunc returns the method "BaseAudioContext.createBuffer".
func (this BaseAudioContext) CreateBufferFunc() (fn js.Func[func(numberOfChannels uint32, length uint32, sampleRate float32) AudioBuffer]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateBufferFunc(
			this.Ref(),
		),
	)
}

// CreateBuffer calls the method "BaseAudioContext.createBuffer".
func (this BaseAudioContext) CreateBuffer(numberOfChannels uint32, length uint32, sampleRate float32) (ret AudioBuffer) {
	bindings.CallBaseAudioContextCreateBuffer(
		this.Ref(), js.Pointer(&ret),
		uint32(numberOfChannels),
		uint32(length),
		float32(sampleRate),
	)

	return
}

// TryCreateBuffer calls the method "BaseAudioContext.createBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateBuffer(numberOfChannels uint32, length uint32, sampleRate float32) (ret AudioBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(numberOfChannels),
		uint32(length),
		float32(sampleRate),
	)

	return
}

// HasCreateBufferSource returns true if the method "BaseAudioContext.createBufferSource" exists.
func (this BaseAudioContext) HasCreateBufferSource() bool {
	return js.True == bindings.HasBaseAudioContextCreateBufferSource(
		this.Ref(),
	)
}

// CreateBufferSourceFunc returns the method "BaseAudioContext.createBufferSource".
func (this BaseAudioContext) CreateBufferSourceFunc() (fn js.Func[func() AudioBufferSourceNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateBufferSourceFunc(
			this.Ref(),
		),
	)
}

// CreateBufferSource calls the method "BaseAudioContext.createBufferSource".
func (this BaseAudioContext) CreateBufferSource() (ret AudioBufferSourceNode) {
	bindings.CallBaseAudioContextCreateBufferSource(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateBufferSource calls the method "BaseAudioContext.createBufferSource"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateBufferSource() (ret AudioBufferSourceNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateBufferSource(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateChannelMerger returns true if the method "BaseAudioContext.createChannelMerger" exists.
func (this BaseAudioContext) HasCreateChannelMerger() bool {
	return js.True == bindings.HasBaseAudioContextCreateChannelMerger(
		this.Ref(),
	)
}

// CreateChannelMergerFunc returns the method "BaseAudioContext.createChannelMerger".
func (this BaseAudioContext) CreateChannelMergerFunc() (fn js.Func[func(numberOfInputs uint32) ChannelMergerNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateChannelMergerFunc(
			this.Ref(),
		),
	)
}

// CreateChannelMerger calls the method "BaseAudioContext.createChannelMerger".
func (this BaseAudioContext) CreateChannelMerger(numberOfInputs uint32) (ret ChannelMergerNode) {
	bindings.CallBaseAudioContextCreateChannelMerger(
		this.Ref(), js.Pointer(&ret),
		uint32(numberOfInputs),
	)

	return
}

// TryCreateChannelMerger calls the method "BaseAudioContext.createChannelMerger"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateChannelMerger(numberOfInputs uint32) (ret ChannelMergerNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateChannelMerger(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(numberOfInputs),
	)

	return
}

// HasCreateChannelMerger1 returns true if the method "BaseAudioContext.createChannelMerger" exists.
func (this BaseAudioContext) HasCreateChannelMerger1() bool {
	return js.True == bindings.HasBaseAudioContextCreateChannelMerger1(
		this.Ref(),
	)
}

// CreateChannelMerger1Func returns the method "BaseAudioContext.createChannelMerger".
func (this BaseAudioContext) CreateChannelMerger1Func() (fn js.Func[func() ChannelMergerNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateChannelMerger1Func(
			this.Ref(),
		),
	)
}

// CreateChannelMerger1 calls the method "BaseAudioContext.createChannelMerger".
func (this BaseAudioContext) CreateChannelMerger1() (ret ChannelMergerNode) {
	bindings.CallBaseAudioContextCreateChannelMerger1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateChannelMerger1 calls the method "BaseAudioContext.createChannelMerger"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateChannelMerger1() (ret ChannelMergerNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateChannelMerger1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateChannelSplitter returns true if the method "BaseAudioContext.createChannelSplitter" exists.
func (this BaseAudioContext) HasCreateChannelSplitter() bool {
	return js.True == bindings.HasBaseAudioContextCreateChannelSplitter(
		this.Ref(),
	)
}

// CreateChannelSplitterFunc returns the method "BaseAudioContext.createChannelSplitter".
func (this BaseAudioContext) CreateChannelSplitterFunc() (fn js.Func[func(numberOfOutputs uint32) ChannelSplitterNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateChannelSplitterFunc(
			this.Ref(),
		),
	)
}

// CreateChannelSplitter calls the method "BaseAudioContext.createChannelSplitter".
func (this BaseAudioContext) CreateChannelSplitter(numberOfOutputs uint32) (ret ChannelSplitterNode) {
	bindings.CallBaseAudioContextCreateChannelSplitter(
		this.Ref(), js.Pointer(&ret),
		uint32(numberOfOutputs),
	)

	return
}

// TryCreateChannelSplitter calls the method "BaseAudioContext.createChannelSplitter"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateChannelSplitter(numberOfOutputs uint32) (ret ChannelSplitterNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateChannelSplitter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(numberOfOutputs),
	)

	return
}

// HasCreateChannelSplitter1 returns true if the method "BaseAudioContext.createChannelSplitter" exists.
func (this BaseAudioContext) HasCreateChannelSplitter1() bool {
	return js.True == bindings.HasBaseAudioContextCreateChannelSplitter1(
		this.Ref(),
	)
}

// CreateChannelSplitter1Func returns the method "BaseAudioContext.createChannelSplitter".
func (this BaseAudioContext) CreateChannelSplitter1Func() (fn js.Func[func() ChannelSplitterNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateChannelSplitter1Func(
			this.Ref(),
		),
	)
}

// CreateChannelSplitter1 calls the method "BaseAudioContext.createChannelSplitter".
func (this BaseAudioContext) CreateChannelSplitter1() (ret ChannelSplitterNode) {
	bindings.CallBaseAudioContextCreateChannelSplitter1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateChannelSplitter1 calls the method "BaseAudioContext.createChannelSplitter"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateChannelSplitter1() (ret ChannelSplitterNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateChannelSplitter1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateConstantSource returns true if the method "BaseAudioContext.createConstantSource" exists.
func (this BaseAudioContext) HasCreateConstantSource() bool {
	return js.True == bindings.HasBaseAudioContextCreateConstantSource(
		this.Ref(),
	)
}

// CreateConstantSourceFunc returns the method "BaseAudioContext.createConstantSource".
func (this BaseAudioContext) CreateConstantSourceFunc() (fn js.Func[func() ConstantSourceNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateConstantSourceFunc(
			this.Ref(),
		),
	)
}

// CreateConstantSource calls the method "BaseAudioContext.createConstantSource".
func (this BaseAudioContext) CreateConstantSource() (ret ConstantSourceNode) {
	bindings.CallBaseAudioContextCreateConstantSource(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateConstantSource calls the method "BaseAudioContext.createConstantSource"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateConstantSource() (ret ConstantSourceNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateConstantSource(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateConvolver returns true if the method "BaseAudioContext.createConvolver" exists.
func (this BaseAudioContext) HasCreateConvolver() bool {
	return js.True == bindings.HasBaseAudioContextCreateConvolver(
		this.Ref(),
	)
}

// CreateConvolverFunc returns the method "BaseAudioContext.createConvolver".
func (this BaseAudioContext) CreateConvolverFunc() (fn js.Func[func() ConvolverNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateConvolverFunc(
			this.Ref(),
		),
	)
}

// CreateConvolver calls the method "BaseAudioContext.createConvolver".
func (this BaseAudioContext) CreateConvolver() (ret ConvolverNode) {
	bindings.CallBaseAudioContextCreateConvolver(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateConvolver calls the method "BaseAudioContext.createConvolver"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateConvolver() (ret ConvolverNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateConvolver(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateDelay returns true if the method "BaseAudioContext.createDelay" exists.
func (this BaseAudioContext) HasCreateDelay() bool {
	return js.True == bindings.HasBaseAudioContextCreateDelay(
		this.Ref(),
	)
}

// CreateDelayFunc returns the method "BaseAudioContext.createDelay".
func (this BaseAudioContext) CreateDelayFunc() (fn js.Func[func(maxDelayTime float64) DelayNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateDelayFunc(
			this.Ref(),
		),
	)
}

// CreateDelay calls the method "BaseAudioContext.createDelay".
func (this BaseAudioContext) CreateDelay(maxDelayTime float64) (ret DelayNode) {
	bindings.CallBaseAudioContextCreateDelay(
		this.Ref(), js.Pointer(&ret),
		float64(maxDelayTime),
	)

	return
}

// TryCreateDelay calls the method "BaseAudioContext.createDelay"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateDelay(maxDelayTime float64) (ret DelayNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateDelay(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(maxDelayTime),
	)

	return
}

// HasCreateDelay1 returns true if the method "BaseAudioContext.createDelay" exists.
func (this BaseAudioContext) HasCreateDelay1() bool {
	return js.True == bindings.HasBaseAudioContextCreateDelay1(
		this.Ref(),
	)
}

// CreateDelay1Func returns the method "BaseAudioContext.createDelay".
func (this BaseAudioContext) CreateDelay1Func() (fn js.Func[func() DelayNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateDelay1Func(
			this.Ref(),
		),
	)
}

// CreateDelay1 calls the method "BaseAudioContext.createDelay".
func (this BaseAudioContext) CreateDelay1() (ret DelayNode) {
	bindings.CallBaseAudioContextCreateDelay1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateDelay1 calls the method "BaseAudioContext.createDelay"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateDelay1() (ret DelayNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateDelay1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateDynamicsCompressor returns true if the method "BaseAudioContext.createDynamicsCompressor" exists.
func (this BaseAudioContext) HasCreateDynamicsCompressor() bool {
	return js.True == bindings.HasBaseAudioContextCreateDynamicsCompressor(
		this.Ref(),
	)
}

// CreateDynamicsCompressorFunc returns the method "BaseAudioContext.createDynamicsCompressor".
func (this BaseAudioContext) CreateDynamicsCompressorFunc() (fn js.Func[func() DynamicsCompressorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateDynamicsCompressorFunc(
			this.Ref(),
		),
	)
}

// CreateDynamicsCompressor calls the method "BaseAudioContext.createDynamicsCompressor".
func (this BaseAudioContext) CreateDynamicsCompressor() (ret DynamicsCompressorNode) {
	bindings.CallBaseAudioContextCreateDynamicsCompressor(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateDynamicsCompressor calls the method "BaseAudioContext.createDynamicsCompressor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateDynamicsCompressor() (ret DynamicsCompressorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateDynamicsCompressor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateGain returns true if the method "BaseAudioContext.createGain" exists.
func (this BaseAudioContext) HasCreateGain() bool {
	return js.True == bindings.HasBaseAudioContextCreateGain(
		this.Ref(),
	)
}

// CreateGainFunc returns the method "BaseAudioContext.createGain".
func (this BaseAudioContext) CreateGainFunc() (fn js.Func[func() GainNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateGainFunc(
			this.Ref(),
		),
	)
}

// CreateGain calls the method "BaseAudioContext.createGain".
func (this BaseAudioContext) CreateGain() (ret GainNode) {
	bindings.CallBaseAudioContextCreateGain(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateGain calls the method "BaseAudioContext.createGain"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateGain() (ret GainNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateGain(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateIIRFilter returns true if the method "BaseAudioContext.createIIRFilter" exists.
func (this BaseAudioContext) HasCreateIIRFilter() bool {
	return js.True == bindings.HasBaseAudioContextCreateIIRFilter(
		this.Ref(),
	)
}

// CreateIIRFilterFunc returns the method "BaseAudioContext.createIIRFilter".
func (this BaseAudioContext) CreateIIRFilterFunc() (fn js.Func[func(feedforward js.Array[float64], feedback js.Array[float64]) IIRFilterNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateIIRFilterFunc(
			this.Ref(),
		),
	)
}

// CreateIIRFilter calls the method "BaseAudioContext.createIIRFilter".
func (this BaseAudioContext) CreateIIRFilter(feedforward js.Array[float64], feedback js.Array[float64]) (ret IIRFilterNode) {
	bindings.CallBaseAudioContextCreateIIRFilter(
		this.Ref(), js.Pointer(&ret),
		feedforward.Ref(),
		feedback.Ref(),
	)

	return
}

// TryCreateIIRFilter calls the method "BaseAudioContext.createIIRFilter"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateIIRFilter(feedforward js.Array[float64], feedback js.Array[float64]) (ret IIRFilterNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateIIRFilter(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		feedforward.Ref(),
		feedback.Ref(),
	)

	return
}

// HasCreateOscillator returns true if the method "BaseAudioContext.createOscillator" exists.
func (this BaseAudioContext) HasCreateOscillator() bool {
	return js.True == bindings.HasBaseAudioContextCreateOscillator(
		this.Ref(),
	)
}

// CreateOscillatorFunc returns the method "BaseAudioContext.createOscillator".
func (this BaseAudioContext) CreateOscillatorFunc() (fn js.Func[func() OscillatorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateOscillatorFunc(
			this.Ref(),
		),
	)
}

// CreateOscillator calls the method "BaseAudioContext.createOscillator".
func (this BaseAudioContext) CreateOscillator() (ret OscillatorNode) {
	bindings.CallBaseAudioContextCreateOscillator(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateOscillator calls the method "BaseAudioContext.createOscillator"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateOscillator() (ret OscillatorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateOscillator(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreatePanner returns true if the method "BaseAudioContext.createPanner" exists.
func (this BaseAudioContext) HasCreatePanner() bool {
	return js.True == bindings.HasBaseAudioContextCreatePanner(
		this.Ref(),
	)
}

// CreatePannerFunc returns the method "BaseAudioContext.createPanner".
func (this BaseAudioContext) CreatePannerFunc() (fn js.Func[func() PannerNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreatePannerFunc(
			this.Ref(),
		),
	)
}

// CreatePanner calls the method "BaseAudioContext.createPanner".
func (this BaseAudioContext) CreatePanner() (ret PannerNode) {
	bindings.CallBaseAudioContextCreatePanner(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreatePanner calls the method "BaseAudioContext.createPanner"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreatePanner() (ret PannerNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreatePanner(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreatePeriodicWave returns true if the method "BaseAudioContext.createPeriodicWave" exists.
func (this BaseAudioContext) HasCreatePeriodicWave() bool {
	return js.True == bindings.HasBaseAudioContextCreatePeriodicWave(
		this.Ref(),
	)
}

// CreatePeriodicWaveFunc returns the method "BaseAudioContext.createPeriodicWave".
func (this BaseAudioContext) CreatePeriodicWaveFunc() (fn js.Func[func(real js.Array[float32], imag js.Array[float32], constraints PeriodicWaveConstraints) PeriodicWave]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreatePeriodicWaveFunc(
			this.Ref(),
		),
	)
}

// CreatePeriodicWave calls the method "BaseAudioContext.createPeriodicWave".
func (this BaseAudioContext) CreatePeriodicWave(real js.Array[float32], imag js.Array[float32], constraints PeriodicWaveConstraints) (ret PeriodicWave) {
	bindings.CallBaseAudioContextCreatePeriodicWave(
		this.Ref(), js.Pointer(&ret),
		real.Ref(),
		imag.Ref(),
		js.Pointer(&constraints),
	)

	return
}

// TryCreatePeriodicWave calls the method "BaseAudioContext.createPeriodicWave"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreatePeriodicWave(real js.Array[float32], imag js.Array[float32], constraints PeriodicWaveConstraints) (ret PeriodicWave, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreatePeriodicWave(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		real.Ref(),
		imag.Ref(),
		js.Pointer(&constraints),
	)

	return
}

// HasCreatePeriodicWave1 returns true if the method "BaseAudioContext.createPeriodicWave" exists.
func (this BaseAudioContext) HasCreatePeriodicWave1() bool {
	return js.True == bindings.HasBaseAudioContextCreatePeriodicWave1(
		this.Ref(),
	)
}

// CreatePeriodicWave1Func returns the method "BaseAudioContext.createPeriodicWave".
func (this BaseAudioContext) CreatePeriodicWave1Func() (fn js.Func[func(real js.Array[float32], imag js.Array[float32]) PeriodicWave]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreatePeriodicWave1Func(
			this.Ref(),
		),
	)
}

// CreatePeriodicWave1 calls the method "BaseAudioContext.createPeriodicWave".
func (this BaseAudioContext) CreatePeriodicWave1(real js.Array[float32], imag js.Array[float32]) (ret PeriodicWave) {
	bindings.CallBaseAudioContextCreatePeriodicWave1(
		this.Ref(), js.Pointer(&ret),
		real.Ref(),
		imag.Ref(),
	)

	return
}

// TryCreatePeriodicWave1 calls the method "BaseAudioContext.createPeriodicWave"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreatePeriodicWave1(real js.Array[float32], imag js.Array[float32]) (ret PeriodicWave, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreatePeriodicWave1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		real.Ref(),
		imag.Ref(),
	)

	return
}

// HasCreateScriptProcessor returns true if the method "BaseAudioContext.createScriptProcessor" exists.
func (this BaseAudioContext) HasCreateScriptProcessor() bool {
	return js.True == bindings.HasBaseAudioContextCreateScriptProcessor(
		this.Ref(),
	)
}

// CreateScriptProcessorFunc returns the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessorFunc() (fn js.Func[func(bufferSize uint32, numberOfInputChannels uint32, numberOfOutputChannels uint32) ScriptProcessorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateScriptProcessorFunc(
			this.Ref(),
		),
	)
}

// CreateScriptProcessor calls the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessor(bufferSize uint32, numberOfInputChannels uint32, numberOfOutputChannels uint32) (ret ScriptProcessorNode) {
	bindings.CallBaseAudioContextCreateScriptProcessor(
		this.Ref(), js.Pointer(&ret),
		uint32(bufferSize),
		uint32(numberOfInputChannels),
		uint32(numberOfOutputChannels),
	)

	return
}

// TryCreateScriptProcessor calls the method "BaseAudioContext.createScriptProcessor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateScriptProcessor(bufferSize uint32, numberOfInputChannels uint32, numberOfOutputChannels uint32) (ret ScriptProcessorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateScriptProcessor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(bufferSize),
		uint32(numberOfInputChannels),
		uint32(numberOfOutputChannels),
	)

	return
}

// HasCreateScriptProcessor1 returns true if the method "BaseAudioContext.createScriptProcessor" exists.
func (this BaseAudioContext) HasCreateScriptProcessor1() bool {
	return js.True == bindings.HasBaseAudioContextCreateScriptProcessor1(
		this.Ref(),
	)
}

// CreateScriptProcessor1Func returns the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessor1Func() (fn js.Func[func(bufferSize uint32, numberOfInputChannels uint32) ScriptProcessorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateScriptProcessor1Func(
			this.Ref(),
		),
	)
}

// CreateScriptProcessor1 calls the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessor1(bufferSize uint32, numberOfInputChannels uint32) (ret ScriptProcessorNode) {
	bindings.CallBaseAudioContextCreateScriptProcessor1(
		this.Ref(), js.Pointer(&ret),
		uint32(bufferSize),
		uint32(numberOfInputChannels),
	)

	return
}

// TryCreateScriptProcessor1 calls the method "BaseAudioContext.createScriptProcessor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateScriptProcessor1(bufferSize uint32, numberOfInputChannels uint32) (ret ScriptProcessorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateScriptProcessor1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(bufferSize),
		uint32(numberOfInputChannels),
	)

	return
}

// HasCreateScriptProcessor2 returns true if the method "BaseAudioContext.createScriptProcessor" exists.
func (this BaseAudioContext) HasCreateScriptProcessor2() bool {
	return js.True == bindings.HasBaseAudioContextCreateScriptProcessor2(
		this.Ref(),
	)
}

// CreateScriptProcessor2Func returns the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessor2Func() (fn js.Func[func(bufferSize uint32) ScriptProcessorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateScriptProcessor2Func(
			this.Ref(),
		),
	)
}

// CreateScriptProcessor2 calls the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessor2(bufferSize uint32) (ret ScriptProcessorNode) {
	bindings.CallBaseAudioContextCreateScriptProcessor2(
		this.Ref(), js.Pointer(&ret),
		uint32(bufferSize),
	)

	return
}

// TryCreateScriptProcessor2 calls the method "BaseAudioContext.createScriptProcessor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateScriptProcessor2(bufferSize uint32) (ret ScriptProcessorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateScriptProcessor2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(bufferSize),
	)

	return
}

// HasCreateScriptProcessor3 returns true if the method "BaseAudioContext.createScriptProcessor" exists.
func (this BaseAudioContext) HasCreateScriptProcessor3() bool {
	return js.True == bindings.HasBaseAudioContextCreateScriptProcessor3(
		this.Ref(),
	)
}

// CreateScriptProcessor3Func returns the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessor3Func() (fn js.Func[func() ScriptProcessorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateScriptProcessor3Func(
			this.Ref(),
		),
	)
}

// CreateScriptProcessor3 calls the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessor3() (ret ScriptProcessorNode) {
	bindings.CallBaseAudioContextCreateScriptProcessor3(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateScriptProcessor3 calls the method "BaseAudioContext.createScriptProcessor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateScriptProcessor3() (ret ScriptProcessorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateScriptProcessor3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateStereoPanner returns true if the method "BaseAudioContext.createStereoPanner" exists.
func (this BaseAudioContext) HasCreateStereoPanner() bool {
	return js.True == bindings.HasBaseAudioContextCreateStereoPanner(
		this.Ref(),
	)
}

// CreateStereoPannerFunc returns the method "BaseAudioContext.createStereoPanner".
func (this BaseAudioContext) CreateStereoPannerFunc() (fn js.Func[func() StereoPannerNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateStereoPannerFunc(
			this.Ref(),
		),
	)
}

// CreateStereoPanner calls the method "BaseAudioContext.createStereoPanner".
func (this BaseAudioContext) CreateStereoPanner() (ret StereoPannerNode) {
	bindings.CallBaseAudioContextCreateStereoPanner(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateStereoPanner calls the method "BaseAudioContext.createStereoPanner"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateStereoPanner() (ret StereoPannerNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateStereoPanner(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateWaveShaper returns true if the method "BaseAudioContext.createWaveShaper" exists.
func (this BaseAudioContext) HasCreateWaveShaper() bool {
	return js.True == bindings.HasBaseAudioContextCreateWaveShaper(
		this.Ref(),
	)
}

// CreateWaveShaperFunc returns the method "BaseAudioContext.createWaveShaper".
func (this BaseAudioContext) CreateWaveShaperFunc() (fn js.Func[func() WaveShaperNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateWaveShaperFunc(
			this.Ref(),
		),
	)
}

// CreateWaveShaper calls the method "BaseAudioContext.createWaveShaper".
func (this BaseAudioContext) CreateWaveShaper() (ret WaveShaperNode) {
	bindings.CallBaseAudioContextCreateWaveShaper(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateWaveShaper calls the method "BaseAudioContext.createWaveShaper"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryCreateWaveShaper() (ret WaveShaperNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateWaveShaper(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDecodeAudioData returns true if the method "BaseAudioContext.decodeAudioData" exists.
func (this BaseAudioContext) HasDecodeAudioData() bool {
	return js.True == bindings.HasBaseAudioContextDecodeAudioData(
		this.Ref(),
	)
}

// DecodeAudioDataFunc returns the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) DecodeAudioDataFunc() (fn js.Func[func(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)], errorCallback js.Func[func(err DOMException)]) js.Promise[AudioBuffer]]) {
	return fn.FromRef(
		bindings.BaseAudioContextDecodeAudioDataFunc(
			this.Ref(),
		),
	)
}

// DecodeAudioData calls the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) DecodeAudioData(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)], errorCallback js.Func[func(err DOMException)]) (ret js.Promise[AudioBuffer]) {
	bindings.CallBaseAudioContextDecodeAudioData(
		this.Ref(), js.Pointer(&ret),
		audioData.Ref(),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// TryDecodeAudioData calls the method "BaseAudioContext.decodeAudioData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryDecodeAudioData(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)], errorCallback js.Func[func(err DOMException)]) (ret js.Promise[AudioBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextDecodeAudioData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		audioData.Ref(),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasDecodeAudioData1 returns true if the method "BaseAudioContext.decodeAudioData" exists.
func (this BaseAudioContext) HasDecodeAudioData1() bool {
	return js.True == bindings.HasBaseAudioContextDecodeAudioData1(
		this.Ref(),
	)
}

// DecodeAudioData1Func returns the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) DecodeAudioData1Func() (fn js.Func[func(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)]) js.Promise[AudioBuffer]]) {
	return fn.FromRef(
		bindings.BaseAudioContextDecodeAudioData1Func(
			this.Ref(),
		),
	)
}

// DecodeAudioData1 calls the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) DecodeAudioData1(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)]) (ret js.Promise[AudioBuffer]) {
	bindings.CallBaseAudioContextDecodeAudioData1(
		this.Ref(), js.Pointer(&ret),
		audioData.Ref(),
		successCallback.Ref(),
	)

	return
}

// TryDecodeAudioData1 calls the method "BaseAudioContext.decodeAudioData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryDecodeAudioData1(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)]) (ret js.Promise[AudioBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextDecodeAudioData1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		audioData.Ref(),
		successCallback.Ref(),
	)

	return
}

// HasDecodeAudioData2 returns true if the method "BaseAudioContext.decodeAudioData" exists.
func (this BaseAudioContext) HasDecodeAudioData2() bool {
	return js.True == bindings.HasBaseAudioContextDecodeAudioData2(
		this.Ref(),
	)
}

// DecodeAudioData2Func returns the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) DecodeAudioData2Func() (fn js.Func[func(audioData js.ArrayBuffer) js.Promise[AudioBuffer]]) {
	return fn.FromRef(
		bindings.BaseAudioContextDecodeAudioData2Func(
			this.Ref(),
		),
	)
}

// DecodeAudioData2 calls the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) DecodeAudioData2(audioData js.ArrayBuffer) (ret js.Promise[AudioBuffer]) {
	bindings.CallBaseAudioContextDecodeAudioData2(
		this.Ref(), js.Pointer(&ret),
		audioData.Ref(),
	)

	return
}

// TryDecodeAudioData2 calls the method "BaseAudioContext.decodeAudioData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this BaseAudioContext) TryDecodeAudioData2(audioData js.ArrayBuffer) (ret js.Promise[AudioBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextDecodeAudioData2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		audioData.Ref(),
	)

	return
}

type AnalyserOptions struct {
	// FftSize is "AnalyserOptions.fftSize"
	//
	// Optional, defaults to 2048.
	//
	// NOTE: FFI_USE_FftSize MUST be set to true to make this field effective.
	FftSize uint32
	// MaxDecibels is "AnalyserOptions.maxDecibels"
	//
	// Optional, defaults to -30.
	//
	// NOTE: FFI_USE_MaxDecibels MUST be set to true to make this field effective.
	MaxDecibels float64
	// MinDecibels is "AnalyserOptions.minDecibels"
	//
	// Optional, defaults to -100.
	//
	// NOTE: FFI_USE_MinDecibels MUST be set to true to make this field effective.
	MinDecibels float64
	// SmoothingTimeConstant is "AnalyserOptions.smoothingTimeConstant"
	//
	// Optional, defaults to 0.8.
	//
	// NOTE: FFI_USE_SmoothingTimeConstant MUST be set to true to make this field effective.
	SmoothingTimeConstant float64
	// ChannelCount is "AnalyserOptions.channelCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChannelCount MUST be set to true to make this field effective.
	ChannelCount uint32
	// ChannelCountMode is "AnalyserOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "AnalyserOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_FftSize               bool // for FftSize.
	FFI_USE_MaxDecibels           bool // for MaxDecibels.
	FFI_USE_MinDecibels           bool // for MinDecibels.
	FFI_USE_SmoothingTimeConstant bool // for SmoothingTimeConstant.
	FFI_USE_ChannelCount          bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AnalyserOptions with all fields set.
func (p AnalyserOptions) FromRef(ref js.Ref) AnalyserOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AnalyserOptions in the application heap.
func (p AnalyserOptions) New() js.Ref {
	return bindings.AnalyserOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AnalyserOptions) UpdateFrom(ref js.Ref) {
	bindings.AnalyserOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AnalyserOptions) Update(ref js.Ref) {
	bindings.AnalyserOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewAnalyserNode(context BaseAudioContext, options AnalyserOptions) (ret AnalyserNode) {
	ret.ref = bindings.NewAnalyserNodeByAnalyserNode(
		context.Ref(),
		js.Pointer(&options))
	return
}

func NewAnalyserNodeByAnalyserNode1(context BaseAudioContext) (ret AnalyserNode) {
	ret.ref = bindings.NewAnalyserNodeByAnalyserNode1(
		context.Ref())
	return
}

type AnalyserNode struct {
	AudioNode
}

func (this AnalyserNode) Once() AnalyserNode {
	this.Ref().Once()
	return this
}

func (this AnalyserNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this AnalyserNode) FromRef(ref js.Ref) AnalyserNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this AnalyserNode) Free() {
	this.Ref().Free()
}

// FftSize returns the value of property "AnalyserNode.fftSize".
//
// It returns ok=false if there is no such property.
func (this AnalyserNode) FftSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAnalyserNodeFftSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFftSize sets the value of property "AnalyserNode.fftSize" to val.
//
// It returns false if the property cannot be set.
func (this AnalyserNode) SetFftSize(val uint32) bool {
	return js.True == bindings.SetAnalyserNodeFftSize(
		this.Ref(),
		uint32(val),
	)
}

// FrequencyBinCount returns the value of property "AnalyserNode.frequencyBinCount".
//
// It returns ok=false if there is no such property.
func (this AnalyserNode) FrequencyBinCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAnalyserNodeFrequencyBinCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MinDecibels returns the value of property "AnalyserNode.minDecibels".
//
// It returns ok=false if there is no such property.
func (this AnalyserNode) MinDecibels() (ret float64, ok bool) {
	ok = js.True == bindings.GetAnalyserNodeMinDecibels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMinDecibels sets the value of property "AnalyserNode.minDecibels" to val.
//
// It returns false if the property cannot be set.
func (this AnalyserNode) SetMinDecibels(val float64) bool {
	return js.True == bindings.SetAnalyserNodeMinDecibels(
		this.Ref(),
		float64(val),
	)
}

// MaxDecibels returns the value of property "AnalyserNode.maxDecibels".
//
// It returns ok=false if there is no such property.
func (this AnalyserNode) MaxDecibels() (ret float64, ok bool) {
	ok = js.True == bindings.GetAnalyserNodeMaxDecibels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMaxDecibels sets the value of property "AnalyserNode.maxDecibels" to val.
//
// It returns false if the property cannot be set.
func (this AnalyserNode) SetMaxDecibels(val float64) bool {
	return js.True == bindings.SetAnalyserNodeMaxDecibels(
		this.Ref(),
		float64(val),
	)
}

// SmoothingTimeConstant returns the value of property "AnalyserNode.smoothingTimeConstant".
//
// It returns ok=false if there is no such property.
func (this AnalyserNode) SmoothingTimeConstant() (ret float64, ok bool) {
	ok = js.True == bindings.GetAnalyserNodeSmoothingTimeConstant(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSmoothingTimeConstant sets the value of property "AnalyserNode.smoothingTimeConstant" to val.
//
// It returns false if the property cannot be set.
func (this AnalyserNode) SetSmoothingTimeConstant(val float64) bool {
	return js.True == bindings.SetAnalyserNodeSmoothingTimeConstant(
		this.Ref(),
		float64(val),
	)
}

// HasGetFloatFrequencyData returns true if the method "AnalyserNode.getFloatFrequencyData" exists.
func (this AnalyserNode) HasGetFloatFrequencyData() bool {
	return js.True == bindings.HasAnalyserNodeGetFloatFrequencyData(
		this.Ref(),
	)
}

// GetFloatFrequencyDataFunc returns the method "AnalyserNode.getFloatFrequencyData".
func (this AnalyserNode) GetFloatFrequencyDataFunc() (fn js.Func[func(array js.TypedArray[float32])]) {
	return fn.FromRef(
		bindings.AnalyserNodeGetFloatFrequencyDataFunc(
			this.Ref(),
		),
	)
}

// GetFloatFrequencyData calls the method "AnalyserNode.getFloatFrequencyData".
func (this AnalyserNode) GetFloatFrequencyData(array js.TypedArray[float32]) (ret js.Void) {
	bindings.CallAnalyserNodeGetFloatFrequencyData(
		this.Ref(), js.Pointer(&ret),
		array.Ref(),
	)

	return
}

// TryGetFloatFrequencyData calls the method "AnalyserNode.getFloatFrequencyData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnalyserNode) TryGetFloatFrequencyData(array js.TypedArray[float32]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnalyserNodeGetFloatFrequencyData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		array.Ref(),
	)

	return
}

// HasGetByteFrequencyData returns true if the method "AnalyserNode.getByteFrequencyData" exists.
func (this AnalyserNode) HasGetByteFrequencyData() bool {
	return js.True == bindings.HasAnalyserNodeGetByteFrequencyData(
		this.Ref(),
	)
}

// GetByteFrequencyDataFunc returns the method "AnalyserNode.getByteFrequencyData".
func (this AnalyserNode) GetByteFrequencyDataFunc() (fn js.Func[func(array js.TypedArray[uint8])]) {
	return fn.FromRef(
		bindings.AnalyserNodeGetByteFrequencyDataFunc(
			this.Ref(),
		),
	)
}

// GetByteFrequencyData calls the method "AnalyserNode.getByteFrequencyData".
func (this AnalyserNode) GetByteFrequencyData(array js.TypedArray[uint8]) (ret js.Void) {
	bindings.CallAnalyserNodeGetByteFrequencyData(
		this.Ref(), js.Pointer(&ret),
		array.Ref(),
	)

	return
}

// TryGetByteFrequencyData calls the method "AnalyserNode.getByteFrequencyData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnalyserNode) TryGetByteFrequencyData(array js.TypedArray[uint8]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnalyserNodeGetByteFrequencyData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		array.Ref(),
	)

	return
}

// HasGetFloatTimeDomainData returns true if the method "AnalyserNode.getFloatTimeDomainData" exists.
func (this AnalyserNode) HasGetFloatTimeDomainData() bool {
	return js.True == bindings.HasAnalyserNodeGetFloatTimeDomainData(
		this.Ref(),
	)
}

// GetFloatTimeDomainDataFunc returns the method "AnalyserNode.getFloatTimeDomainData".
func (this AnalyserNode) GetFloatTimeDomainDataFunc() (fn js.Func[func(array js.TypedArray[float32])]) {
	return fn.FromRef(
		bindings.AnalyserNodeGetFloatTimeDomainDataFunc(
			this.Ref(),
		),
	)
}

// GetFloatTimeDomainData calls the method "AnalyserNode.getFloatTimeDomainData".
func (this AnalyserNode) GetFloatTimeDomainData(array js.TypedArray[float32]) (ret js.Void) {
	bindings.CallAnalyserNodeGetFloatTimeDomainData(
		this.Ref(), js.Pointer(&ret),
		array.Ref(),
	)

	return
}

// TryGetFloatTimeDomainData calls the method "AnalyserNode.getFloatTimeDomainData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnalyserNode) TryGetFloatTimeDomainData(array js.TypedArray[float32]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnalyserNodeGetFloatTimeDomainData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		array.Ref(),
	)

	return
}

// HasGetByteTimeDomainData returns true if the method "AnalyserNode.getByteTimeDomainData" exists.
func (this AnalyserNode) HasGetByteTimeDomainData() bool {
	return js.True == bindings.HasAnalyserNodeGetByteTimeDomainData(
		this.Ref(),
	)
}

// GetByteTimeDomainDataFunc returns the method "AnalyserNode.getByteTimeDomainData".
func (this AnalyserNode) GetByteTimeDomainDataFunc() (fn js.Func[func(array js.TypedArray[uint8])]) {
	return fn.FromRef(
		bindings.AnalyserNodeGetByteTimeDomainDataFunc(
			this.Ref(),
		),
	)
}

// GetByteTimeDomainData calls the method "AnalyserNode.getByteTimeDomainData".
func (this AnalyserNode) GetByteTimeDomainData(array js.TypedArray[uint8]) (ret js.Void) {
	bindings.CallAnalyserNodeGetByteTimeDomainData(
		this.Ref(), js.Pointer(&ret),
		array.Ref(),
	)

	return
}

// TryGetByteTimeDomainData calls the method "AnalyserNode.getByteTimeDomainData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnalyserNode) TryGetByteTimeDomainData(array js.TypedArray[uint8]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnalyserNodeGetByteTimeDomainData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		array.Ref(),
	)

	return
}

type AnimationEventInit struct {
	// AnimationName is "AnimationEventInit.animationName"
	//
	// Optional, defaults to "".
	AnimationName js.String
	// ElapsedTime is "AnimationEventInit.elapsedTime"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_ElapsedTime MUST be set to true to make this field effective.
	ElapsedTime float64
	// PseudoElement is "AnimationEventInit.pseudoElement"
	//
	// Optional, defaults to "".
	PseudoElement js.String
	// Bubbles is "AnimationEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "AnimationEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "AnimationEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_ElapsedTime bool // for ElapsedTime.
	FFI_USE_Bubbles     bool // for Bubbles.
	FFI_USE_Cancelable  bool // for Cancelable.
	FFI_USE_Composed    bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AnimationEventInit with all fields set.
func (p AnimationEventInit) FromRef(ref js.Ref) AnimationEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AnimationEventInit in the application heap.
func (p AnimationEventInit) New() js.Ref {
	return bindings.AnimationEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AnimationEventInit) UpdateFrom(ref js.Ref) {
	bindings.AnimationEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AnimationEventInit) Update(ref js.Ref) {
	bindings.AnimationEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewAnimationEvent(typ js.String, animationEventInitDict AnimationEventInit) (ret AnimationEvent) {
	ret.ref = bindings.NewAnimationEventByAnimationEvent(
		typ.Ref(),
		js.Pointer(&animationEventInitDict))
	return
}

func NewAnimationEventByAnimationEvent1(typ js.String) (ret AnimationEvent) {
	ret.ref = bindings.NewAnimationEventByAnimationEvent1(
		typ.Ref())
	return
}

type AnimationEvent struct {
	Event
}

func (this AnimationEvent) Once() AnimationEvent {
	this.Ref().Once()
	return this
}

func (this AnimationEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this AnimationEvent) FromRef(ref js.Ref) AnimationEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this AnimationEvent) Free() {
	this.Ref().Free()
}

// AnimationName returns the value of property "AnimationEvent.animationName".
//
// It returns ok=false if there is no such property.
func (this AnimationEvent) AnimationName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAnimationEventAnimationName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ElapsedTime returns the value of property "AnimationEvent.elapsedTime".
//
// It returns ok=false if there is no such property.
func (this AnimationEvent) ElapsedTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetAnimationEventElapsedTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PseudoElement returns the value of property "AnimationEvent.pseudoElement".
//
// It returns ok=false if there is no such property.
func (this AnimationEvent) PseudoElement() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAnimationEventPseudoElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type AnimationPlaybackEventInit struct {
	// CurrentTime is "AnimationPlaybackEventInit.currentTime"
	//
	// Optional, defaults to null.
	CurrentTime CSSNumberish
	// TimelineTime is "AnimationPlaybackEventInit.timelineTime"
	//
	// Optional, defaults to null.
	TimelineTime CSSNumberish
	// Bubbles is "AnimationPlaybackEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "AnimationPlaybackEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "AnimationPlaybackEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AnimationPlaybackEventInit with all fields set.
func (p AnimationPlaybackEventInit) FromRef(ref js.Ref) AnimationPlaybackEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AnimationPlaybackEventInit in the application heap.
func (p AnimationPlaybackEventInit) New() js.Ref {
	return bindings.AnimationPlaybackEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AnimationPlaybackEventInit) UpdateFrom(ref js.Ref) {
	bindings.AnimationPlaybackEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AnimationPlaybackEventInit) Update(ref js.Ref) {
	bindings.AnimationPlaybackEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewAnimationPlaybackEvent(typ js.String, eventInitDict AnimationPlaybackEventInit) (ret AnimationPlaybackEvent) {
	ret.ref = bindings.NewAnimationPlaybackEventByAnimationPlaybackEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewAnimationPlaybackEventByAnimationPlaybackEvent1(typ js.String) (ret AnimationPlaybackEvent) {
	ret.ref = bindings.NewAnimationPlaybackEventByAnimationPlaybackEvent1(
		typ.Ref())
	return
}

type AnimationPlaybackEvent struct {
	Event
}

func (this AnimationPlaybackEvent) Once() AnimationPlaybackEvent {
	this.Ref().Once()
	return this
}

func (this AnimationPlaybackEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this AnimationPlaybackEvent) FromRef(ref js.Ref) AnimationPlaybackEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this AnimationPlaybackEvent) Free() {
	this.Ref().Free()
}

// CurrentTime returns the value of property "AnimationPlaybackEvent.currentTime".
//
// It returns ok=false if there is no such property.
func (this AnimationPlaybackEvent) CurrentTime() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationPlaybackEventCurrentTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TimelineTime returns the value of property "AnimationPlaybackEvent.timelineTime".
//
// It returns ok=false if there is no such property.
func (this AnimationPlaybackEvent) TimelineTime() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationPlaybackEventTimelineTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type AnimatorInstanceConstructorFunc func(this js.Ref, options js.Any, state js.Any) js.Ref

func (fn AnimatorInstanceConstructorFunc) Register() js.Func[func(options js.Any, state js.Any) js.Any] {
	return js.RegisterCallback[func(options js.Any, state js.Any) js.Any](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AnimatorInstanceConstructorFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
		js.Any{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AnimatorInstanceConstructor[T any] struct {
	Fn  func(arg T, this js.Ref, options js.Any, state js.Any) js.Ref
	Arg T
}

func (cb *AnimatorInstanceConstructor[T]) Register() js.Func[func(options js.Any, state js.Any) js.Any] {
	return js.RegisterCallback[func(options js.Any, state js.Any) js.Any](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AnimatorInstanceConstructor[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Any{}.FromRef(args[0+1]),
		js.Any{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AnimationWorkletGlobalScope struct {
	WorkletGlobalScope
}

func (this AnimationWorkletGlobalScope) Once() AnimationWorkletGlobalScope {
	this.Ref().Once()
	return this
}

func (this AnimationWorkletGlobalScope) Ref() js.Ref {
	return this.WorkletGlobalScope.Ref()
}

func (this AnimationWorkletGlobalScope) FromRef(ref js.Ref) AnimationWorkletGlobalScope {
	this.WorkletGlobalScope = this.WorkletGlobalScope.FromRef(ref)
	return this
}

func (this AnimationWorkletGlobalScope) Free() {
	this.Ref().Free()
}

// HasRegisterAnimator returns true if the method "AnimationWorkletGlobalScope.registerAnimator" exists.
func (this AnimationWorkletGlobalScope) HasRegisterAnimator() bool {
	return js.True == bindings.HasAnimationWorkletGlobalScopeRegisterAnimator(
		this.Ref(),
	)
}

// RegisterAnimatorFunc returns the method "AnimationWorkletGlobalScope.registerAnimator".
func (this AnimationWorkletGlobalScope) RegisterAnimatorFunc() (fn js.Func[func(name js.String, animatorCtor js.Func[func(options js.Any, state js.Any) js.Any])]) {
	return fn.FromRef(
		bindings.AnimationWorkletGlobalScopeRegisterAnimatorFunc(
			this.Ref(),
		),
	)
}

// RegisterAnimator calls the method "AnimationWorkletGlobalScope.registerAnimator".
func (this AnimationWorkletGlobalScope) RegisterAnimator(name js.String, animatorCtor js.Func[func(options js.Any, state js.Any) js.Any]) (ret js.Void) {
	bindings.CallAnimationWorkletGlobalScopeRegisterAnimator(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		animatorCtor.Ref(),
	)

	return
}

// TryRegisterAnimator calls the method "AnimationWorkletGlobalScope.registerAnimator"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this AnimationWorkletGlobalScope) TryRegisterAnimator(name js.String, animatorCtor js.Func[func(options js.Any, state js.Any) js.Any]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationWorkletGlobalScopeRegisterAnimator(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		animatorCtor.Ref(),
	)

	return
}

type AppBannerPromptOutcome uint32

const (
	_ AppBannerPromptOutcome = iota

	AppBannerPromptOutcome_ACCEPTED
	AppBannerPromptOutcome_DISMISSED
)

func (AppBannerPromptOutcome) FromRef(str js.Ref) AppBannerPromptOutcome {
	return AppBannerPromptOutcome(bindings.ConstOfAppBannerPromptOutcome(str))
}

func (x AppBannerPromptOutcome) String() (string, bool) {
	switch x {
	case AppBannerPromptOutcome_ACCEPTED:
		return "accepted", true
	case AppBannerPromptOutcome_DISMISSED:
		return "dismissed", true
	default:
		return "", false
	}
}

type AppendMode uint32

const (
	_ AppendMode = iota

	AppendMode_SEGMENTS
	AppendMode_SEQUENCE
)

func (AppendMode) FromRef(str js.Ref) AppendMode {
	return AppendMode(bindings.ConstOfAppendMode(str))
}

func (x AppendMode) String() (string, bool) {
	switch x {
	case AppendMode_SEGMENTS:
		return "segments", true
	case AppendMode_SEQUENCE:
		return "sequence", true
	default:
		return "", false
	}
}

type OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView struct {
	ref js.Ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) FromRef(ref js.Ref) OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView {
	return OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView{
		ref: ref,
	}
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

type ArrayBufferView = OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView

type AttestationConveyancePreference uint32

const (
	_ AttestationConveyancePreference = iota

	AttestationConveyancePreference_NONE
	AttestationConveyancePreference_INDIRECT
	AttestationConveyancePreference_DIRECT
	AttestationConveyancePreference_ENTERPRISE
)

func (AttestationConveyancePreference) FromRef(str js.Ref) AttestationConveyancePreference {
	return AttestationConveyancePreference(bindings.ConstOfAttestationConveyancePreference(str))
}

func (x AttestationConveyancePreference) String() (string, bool) {
	switch x {
	case AttestationConveyancePreference_NONE:
		return "none", true
	case AttestationConveyancePreference_INDIRECT:
		return "indirect", true
	case AttestationConveyancePreference_DIRECT:
		return "direct", true
	case AttestationConveyancePreference_ENTERPRISE:
		return "enterprise", true
	default:
		return "", false
	}
}

type AttributionReportingRequestOptions struct {
	// EventSourceEligible is "AttributionReportingRequestOptions.eventSourceEligible"
	//
	// Required
	EventSourceEligible bool
	// TriggerEligible is "AttributionReportingRequestOptions.triggerEligible"
	//
	// Required
	TriggerEligible bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AttributionReportingRequestOptions with all fields set.
func (p AttributionReportingRequestOptions) FromRef(ref js.Ref) AttributionReportingRequestOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AttributionReportingRequestOptions in the application heap.
func (p AttributionReportingRequestOptions) New() js.Ref {
	return bindings.AttributionReportingRequestOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AttributionReportingRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.AttributionReportingRequestOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AttributionReportingRequestOptions) Update(ref js.Ref) {
	bindings.AttributionReportingRequestOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AuctionAd struct {
	// RenderURL is "AuctionAd.renderURL"
	//
	// Required
	RenderURL js.String
	// Metadata is "AuctionAd.metadata"
	//
	// Optional
	Metadata js.Any
	// BuyerReportingId is "AuctionAd.buyerReportingId"
	//
	// Optional
	BuyerReportingId js.String
	// BuyerAndSellerReportingId is "AuctionAd.buyerAndSellerReportingId"
	//
	// Optional
	BuyerAndSellerReportingId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuctionAd with all fields set.
func (p AuctionAd) FromRef(ref js.Ref) AuctionAd {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuctionAd in the application heap.
func (p AuctionAd) New() js.Ref {
	return bindings.AuctionAdJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AuctionAd) UpdateFrom(ref js.Ref) {
	bindings.AuctionAdJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuctionAd) Update(ref js.Ref) {
	bindings.AuctionAdJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AuctionAdConfig struct {
	// Seller is "AuctionAdConfig.seller"
	//
	// Required
	Seller js.String
	// DecisionLogicURL is "AuctionAdConfig.decisionLogicURL"
	//
	// Required
	DecisionLogicURL js.String
	// TrustedScoringSignalsURL is "AuctionAdConfig.trustedScoringSignalsURL"
	//
	// Optional
	TrustedScoringSignalsURL js.String
	// InterestGroupBuyers is "AuctionAdConfig.interestGroupBuyers"
	//
	// Optional
	InterestGroupBuyers js.Array[js.String]
	// AuctionSignals is "AuctionAdConfig.auctionSignals"
	//
	// Optional
	AuctionSignals js.Promise[js.Any]
	// SellerSignals is "AuctionAdConfig.sellerSignals"
	//
	// Optional
	SellerSignals js.Promise[js.Any]
	// DirectFromSellerSignals is "AuctionAdConfig.directFromSellerSignals"
	//
	// Optional
	DirectFromSellerSignals js.Promise[js.String]
	// SellerTimeout is "AuctionAdConfig.sellerTimeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_SellerTimeout MUST be set to true to make this field effective.
	SellerTimeout uint64
	// SellerExperimentGroupId is "AuctionAdConfig.sellerExperimentGroupId"
	//
	// Optional
	//
	// NOTE: FFI_USE_SellerExperimentGroupId MUST be set to true to make this field effective.
	SellerExperimentGroupId uint16
	// SellerCurrency is "AuctionAdConfig.sellerCurrency"
	//
	// Optional
	SellerCurrency js.String
	// PerBuyerSignals is "AuctionAdConfig.perBuyerSignals"
	//
	// Optional
	PerBuyerSignals js.Promise[js.Record[js.Any]]
	// PerBuyerTimeouts is "AuctionAdConfig.perBuyerTimeouts"
	//
	// Optional
	PerBuyerTimeouts js.Promise[js.Record[uint64]]
	// PerBuyerGroupLimits is "AuctionAdConfig.perBuyerGroupLimits"
	//
	// Optional
	PerBuyerGroupLimits js.Record[uint16]
	// PerBuyerExperimentGroupIds is "AuctionAdConfig.perBuyerExperimentGroupIds"
	//
	// Optional
	PerBuyerExperimentGroupIds js.Record[uint16]
	// PerBuyerPrioritySignals is "AuctionAdConfig.perBuyerPrioritySignals"
	//
	// Optional
	PerBuyerPrioritySignals js.Record[js.Record[float64]]
	// PerBuyerCurrencies is "AuctionAdConfig.perBuyerCurrencies"
	//
	// Optional
	PerBuyerCurrencies js.Promise[js.Record[js.String]]
	// ComponentAuctions is "AuctionAdConfig.componentAuctions"
	//
	// Optional, defaults to [].
	ComponentAuctions js.Array[AuctionAdConfig]
	// Signal is "AuctionAdConfig.signal"
	//
	// Optional
	Signal AbortSignal
	// ResolveToConfig is "AuctionAdConfig.resolveToConfig"
	//
	// Optional
	ResolveToConfig js.Promise[js.Boolean]

	FFI_USE_SellerTimeout           bool // for SellerTimeout.
	FFI_USE_SellerExperimentGroupId bool // for SellerExperimentGroupId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuctionAdConfig with all fields set.
func (p AuctionAdConfig) FromRef(ref js.Ref) AuctionAdConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuctionAdConfig in the application heap.
func (p AuctionAdConfig) New() js.Ref {
	return bindings.AuctionAdConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AuctionAdConfig) UpdateFrom(ref js.Ref) {
	bindings.AuctionAdConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuctionAdConfig) Update(ref js.Ref) {
	bindings.AuctionAdConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AuctionAdInterestGroup struct {
	// Priority is "AuctionAdInterestGroup.priority"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_Priority MUST be set to true to make this field effective.
	Priority float64
	// PrioritySignalsOverrides is "AuctionAdInterestGroup.prioritySignalsOverrides"
	//
	// Optional
	PrioritySignalsOverrides js.Record[float64]
	// Owner is "AuctionAdInterestGroup.owner"
	//
	// Required
	Owner js.String
	// Name is "AuctionAdInterestGroup.name"
	//
	// Required
	Name js.String
	// LifetimeMs is "AuctionAdInterestGroup.lifetimeMs"
	//
	// Required
	LifetimeMs float64
	// EnableBiddingSignalsPrioritization is "AuctionAdInterestGroup.enableBiddingSignalsPrioritization"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_EnableBiddingSignalsPrioritization MUST be set to true to make this field effective.
	EnableBiddingSignalsPrioritization bool
	// PriorityVector is "AuctionAdInterestGroup.priorityVector"
	//
	// Optional
	PriorityVector js.Record[float64]
	// ExecutionMode is "AuctionAdInterestGroup.executionMode"
	//
	// Optional, defaults to "compatibility".
	ExecutionMode js.String
	// BiddingLogicURL is "AuctionAdInterestGroup.biddingLogicURL"
	//
	// Optional
	BiddingLogicURL js.String
	// BiddingWasmHelperURL is "AuctionAdInterestGroup.biddingWasmHelperURL"
	//
	// Optional
	BiddingWasmHelperURL js.String
	// UpdateURL is "AuctionAdInterestGroup.updateURL"
	//
	// Optional
	UpdateURL js.String
	// TrustedBiddingSignalsURL is "AuctionAdInterestGroup.trustedBiddingSignalsURL"
	//
	// Optional
	TrustedBiddingSignalsURL js.String
	// TrustedBiddingSignalsKeys is "AuctionAdInterestGroup.trustedBiddingSignalsKeys"
	//
	// Optional
	TrustedBiddingSignalsKeys js.Array[js.String]
	// UserBiddingSignals is "AuctionAdInterestGroup.userBiddingSignals"
	//
	// Optional
	UserBiddingSignals js.Any
	// Ads is "AuctionAdInterestGroup.ads"
	//
	// Optional
	Ads js.Array[AuctionAd]
	// AdComponents is "AuctionAdInterestGroup.adComponents"
	//
	// Optional
	AdComponents js.Array[AuctionAd]

	FFI_USE_Priority                           bool // for Priority.
	FFI_USE_EnableBiddingSignalsPrioritization bool // for EnableBiddingSignalsPrioritization.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuctionAdInterestGroup with all fields set.
func (p AuctionAdInterestGroup) FromRef(ref js.Ref) AuctionAdInterestGroup {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuctionAdInterestGroup in the application heap.
func (p AuctionAdInterestGroup) New() js.Ref {
	return bindings.AuctionAdInterestGroupJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AuctionAdInterestGroup) UpdateFrom(ref js.Ref) {
	bindings.AuctionAdInterestGroupJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuctionAdInterestGroup) Update(ref js.Ref) {
	bindings.AuctionAdInterestGroupJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AuctionAdInterestGroupKey struct {
	// Owner is "AuctionAdInterestGroupKey.owner"
	//
	// Required
	Owner js.String
	// Name is "AuctionAdInterestGroupKey.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AuctionAdInterestGroupKey with all fields set.
func (p AuctionAdInterestGroupKey) FromRef(ref js.Ref) AuctionAdInterestGroupKey {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AuctionAdInterestGroupKey in the application heap.
func (p AuctionAdInterestGroupKey) New() js.Ref {
	return bindings.AuctionAdInterestGroupKeyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AuctionAdInterestGroupKey) UpdateFrom(ref js.Ref) {
	bindings.AuctionAdInterestGroupKeyJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AuctionAdInterestGroupKey) Update(ref js.Ref) {
	bindings.AuctionAdInterestGroupKeyJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AudioConfiguration struct {
	// ContentType is "AudioConfiguration.contentType"
	//
	// Required
	ContentType js.String
	// Channels is "AudioConfiguration.channels"
	//
	// Optional
	Channels js.String
	// Bitrate is "AudioConfiguration.bitrate"
	//
	// Optional
	//
	// NOTE: FFI_USE_Bitrate MUST be set to true to make this field effective.
	Bitrate uint64
	// Samplerate is "AudioConfiguration.samplerate"
	//
	// Optional
	//
	// NOTE: FFI_USE_Samplerate MUST be set to true to make this field effective.
	Samplerate uint32
	// SpatialRendering is "AudioConfiguration.spatialRendering"
	//
	// Optional
	//
	// NOTE: FFI_USE_SpatialRendering MUST be set to true to make this field effective.
	SpatialRendering bool

	FFI_USE_Bitrate          bool // for Bitrate.
	FFI_USE_Samplerate       bool // for Samplerate.
	FFI_USE_SpatialRendering bool // for SpatialRendering.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioConfiguration with all fields set.
func (p AudioConfiguration) FromRef(ref js.Ref) AudioConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioConfiguration in the application heap.
func (p AudioConfiguration) New() js.Ref {
	return bindings.AudioConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioConfiguration) UpdateFrom(ref js.Ref) {
	bindings.AudioConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioConfiguration) Update(ref js.Ref) {
	bindings.AudioConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AudioContextLatencyCategory uint32

const (
	_ AudioContextLatencyCategory = iota

	AudioContextLatencyCategory_BALANCED
	AudioContextLatencyCategory_INTERACTIVE
	AudioContextLatencyCategory_PLAYBACK
)

func (AudioContextLatencyCategory) FromRef(str js.Ref) AudioContextLatencyCategory {
	return AudioContextLatencyCategory(bindings.ConstOfAudioContextLatencyCategory(str))
}

func (x AudioContextLatencyCategory) String() (string, bool) {
	switch x {
	case AudioContextLatencyCategory_BALANCED:
		return "balanced", true
	case AudioContextLatencyCategory_INTERACTIVE:
		return "interactive", true
	case AudioContextLatencyCategory_PLAYBACK:
		return "playback", true
	default:
		return "", false
	}
}

type OneOf_AudioContextLatencyCategory_Float64 struct {
	ref js.Ref
}

func (x OneOf_AudioContextLatencyCategory_Float64) Ref() js.Ref {
	return x.ref
}

func (x OneOf_AudioContextLatencyCategory_Float64) Free() {
	x.ref.Free()
}

func (x OneOf_AudioContextLatencyCategory_Float64) FromRef(ref js.Ref) OneOf_AudioContextLatencyCategory_Float64 {
	return OneOf_AudioContextLatencyCategory_Float64{
		ref: ref,
	}
}

func (x OneOf_AudioContextLatencyCategory_Float64) AudioContextLatencyCategory() AudioContextLatencyCategory {
	return AudioContextLatencyCategory(0).FromRef(x.ref)
}

func (x OneOf_AudioContextLatencyCategory_Float64) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

type AudioSinkType uint32

const (
	_ AudioSinkType = iota

	AudioSinkType_NONE
)

func (AudioSinkType) FromRef(str js.Ref) AudioSinkType {
	return AudioSinkType(bindings.ConstOfAudioSinkType(str))
}

func (x AudioSinkType) String() (string, bool) {
	switch x {
	case AudioSinkType_NONE:
		return "none", true
	default:
		return "", false
	}
}

type AudioSinkOptions struct {
	// Type is "AudioSinkOptions.type"
	//
	// Required
	Type AudioSinkType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioSinkOptions with all fields set.
func (p AudioSinkOptions) FromRef(ref js.Ref) AudioSinkOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioSinkOptions in the application heap.
func (p AudioSinkOptions) New() js.Ref {
	return bindings.AudioSinkOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioSinkOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioSinkOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioSinkOptions) Update(ref js.Ref) {
	bindings.AudioSinkOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_String_AudioSinkOptions struct {
	ref js.Ref
}

func (x OneOf_String_AudioSinkOptions) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_AudioSinkOptions) Free() {
	x.ref.Free()
}

func (x OneOf_String_AudioSinkOptions) FromRef(ref js.Ref) OneOf_String_AudioSinkOptions {
	return OneOf_String_AudioSinkOptions{
		ref: ref,
	}
}

func (x OneOf_String_AudioSinkOptions) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_AudioSinkOptions) AudioSinkOptions() AudioSinkOptions {
	var ret AudioSinkOptions
	ret.UpdateFrom(x.ref)
	return ret
}

type AudioContextOptions struct {
	// LatencyHint is "AudioContextOptions.latencyHint"
	//
	// Optional, defaults to "interactive".
	LatencyHint OneOf_AudioContextLatencyCategory_Float64
	// SampleRate is "AudioContextOptions.sampleRate"
	//
	// Optional
	//
	// NOTE: FFI_USE_SampleRate MUST be set to true to make this field effective.
	SampleRate float32
	// SinkId is "AudioContextOptions.sinkId"
	//
	// Optional
	SinkId OneOf_String_AudioSinkOptions

	FFI_USE_SampleRate bool // for SampleRate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioContextOptions with all fields set.
func (p AudioContextOptions) FromRef(ref js.Ref) AudioContextOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioContextOptions in the application heap.
func (p AudioContextOptions) New() js.Ref {
	return bindings.AudioContextOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioContextOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioContextOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioContextOptions) Update(ref js.Ref) {
	bindings.AudioContextOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AudioTimestamp struct {
	// ContextTime is "AudioTimestamp.contextTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_ContextTime MUST be set to true to make this field effective.
	ContextTime float64
	// PerformanceTime is "AudioTimestamp.performanceTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_PerformanceTime MUST be set to true to make this field effective.
	PerformanceTime DOMHighResTimeStamp

	FFI_USE_ContextTime     bool // for ContextTime.
	FFI_USE_PerformanceTime bool // for PerformanceTime.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AudioTimestamp with all fields set.
func (p AudioTimestamp) FromRef(ref js.Ref) AudioTimestamp {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AudioTimestamp in the application heap.
func (p AudioTimestamp) New() js.Ref {
	return bindings.AudioTimestampJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AudioTimestamp) UpdateFrom(ref js.Ref) {
	bindings.AudioTimestampJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AudioTimestamp) Update(ref js.Ref) {
	bindings.AudioTimestampJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

const (
	HTMLMediaElement_NETWORK_EMPTY     uint16 = 0
	HTMLMediaElement_NETWORK_IDLE      uint16 = 1
	HTMLMediaElement_NETWORK_LOADING   uint16 = 2
	HTMLMediaElement_NETWORK_NO_SOURCE uint16 = 3
	HTMLMediaElement_HAVE_NOTHING      uint16 = 0
	HTMLMediaElement_HAVE_METADATA     uint16 = 1
	HTMLMediaElement_HAVE_CURRENT_DATA uint16 = 2
	HTMLMediaElement_HAVE_FUTURE_DATA  uint16 = 3
	HTMLMediaElement_HAVE_ENOUGH_DATA  uint16 = 4
)

type CanPlayTypeResult uint32

const (
	_ CanPlayTypeResult = iota

	CanPlayTypeResult_
	CanPlayTypeResult_MAYBE
	CanPlayTypeResult_PROBABLY
)

func (CanPlayTypeResult) FromRef(str js.Ref) CanPlayTypeResult {
	return CanPlayTypeResult(bindings.ConstOfCanPlayTypeResult(str))
}

func (x CanPlayTypeResult) String() (string, bool) {
	switch x {
	case CanPlayTypeResult_:
		return "", true
	case CanPlayTypeResult_MAYBE:
		return "maybe", true
	case CanPlayTypeResult_PROBABLY:
		return "probably", true
	default:
		return "", false
	}
}

type TextTrackCue struct {
	EventTarget
}

func (this TextTrackCue) Once() TextTrackCue {
	this.Ref().Once()
	return this
}

func (this TextTrackCue) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this TextTrackCue) FromRef(ref js.Ref) TextTrackCue {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this TextTrackCue) Free() {
	this.Ref().Free()
}

// Track returns the value of property "TextTrackCue.track".
//
// It returns ok=false if there is no such property.
func (this TextTrackCue) Track() (ret TextTrack, ok bool) {
	ok = js.True == bindings.GetTextTrackCueTrack(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "TextTrackCue.id".
//
// It returns ok=false if there is no such property.
func (this TextTrackCue) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextTrackCueId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetId sets the value of property "TextTrackCue.id" to val.
//
// It returns false if the property cannot be set.
func (this TextTrackCue) SetId(val js.String) bool {
	return js.True == bindings.SetTextTrackCueId(
		this.Ref(),
		val.Ref(),
	)
}

// StartTime returns the value of property "TextTrackCue.startTime".
//
// It returns ok=false if there is no such property.
func (this TextTrackCue) StartTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextTrackCueStartTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetStartTime sets the value of property "TextTrackCue.startTime" to val.
//
// It returns false if the property cannot be set.
func (this TextTrackCue) SetStartTime(val float64) bool {
	return js.True == bindings.SetTextTrackCueStartTime(
		this.Ref(),
		float64(val),
	)
}

// EndTime returns the value of property "TextTrackCue.endTime".
//
// It returns ok=false if there is no such property.
func (this TextTrackCue) EndTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextTrackCueEndTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetEndTime sets the value of property "TextTrackCue.endTime" to val.
//
// It returns false if the property cannot be set.
func (this TextTrackCue) SetEndTime(val float64) bool {
	return js.True == bindings.SetTextTrackCueEndTime(
		this.Ref(),
		float64(val),
	)
}

// PauseOnExit returns the value of property "TextTrackCue.pauseOnExit".
//
// It returns ok=false if there is no such property.
func (this TextTrackCue) PauseOnExit() (ret bool, ok bool) {
	ok = js.True == bindings.GetTextTrackCuePauseOnExit(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPauseOnExit sets the value of property "TextTrackCue.pauseOnExit" to val.
//
// It returns false if the property cannot be set.
func (this TextTrackCue) SetPauseOnExit(val bool) bool {
	return js.True == bindings.SetTextTrackCuePauseOnExit(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

type TextTrackKind uint32

const (
	_ TextTrackKind = iota

	TextTrackKind_SUBTITLES
	TextTrackKind_CAPTIONS
	TextTrackKind_DESCRIPTIONS
	TextTrackKind_CHAPTERS
	TextTrackKind_METADATA
)

func (TextTrackKind) FromRef(str js.Ref) TextTrackKind {
	return TextTrackKind(bindings.ConstOfTextTrackKind(str))
}

func (x TextTrackKind) String() (string, bool) {
	switch x {
	case TextTrackKind_SUBTITLES:
		return "subtitles", true
	case TextTrackKind_CAPTIONS:
		return "captions", true
	case TextTrackKind_DESCRIPTIONS:
		return "descriptions", true
	case TextTrackKind_CHAPTERS:
		return "chapters", true
	case TextTrackKind_METADATA:
		return "metadata", true
	default:
		return "", false
	}
}

type TextTrackMode uint32

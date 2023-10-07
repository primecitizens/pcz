// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
		js.ThrowInvalidCallbackInvocation()
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "DOMException.name".
//
// It returns ok=false if there is no such property.
func (this DOMException) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDOMExceptionName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "DOMException.message".
//
// It returns ok=false if there is no such property.
func (this DOMException) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDOMExceptionMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Code returns the value of property "DOMException.code".
//
// It returns ok=false if there is no such property.
func (this DOMException) Code() (ret uint16, ok bool) {
	ok = js.True == bindings.GetDOMExceptionCode(
		this.ref, js.Pointer(&ret),
	)
	return
}

type AudioDestinationNode struct {
	AudioNode
}

func (this AudioDestinationNode) Once() AudioDestinationNode {
	this.ref.Once()
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
	this.ref.Free()
}

// MaxChannelCount returns the value of property "AudioDestinationNode.maxChannelCount".
//
// It returns ok=false if there is no such property.
func (this AudioDestinationNode) MaxChannelCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAudioDestinationNodeMaxChannelCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

type AudioListener struct {
	ref js.Ref
}

func (this AudioListener) Once() AudioListener {
	this.ref.Once()
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
	this.ref.Free()
}

// PositionX returns the value of property "AudioListener.positionX".
//
// It returns ok=false if there is no such property.
func (this AudioListener) PositionX() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerPositionX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PositionY returns the value of property "AudioListener.positionY".
//
// It returns ok=false if there is no such property.
func (this AudioListener) PositionY() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerPositionY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PositionZ returns the value of property "AudioListener.positionZ".
//
// It returns ok=false if there is no such property.
func (this AudioListener) PositionZ() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerPositionZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ForwardX returns the value of property "AudioListener.forwardX".
//
// It returns ok=false if there is no such property.
func (this AudioListener) ForwardX() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerForwardX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ForwardY returns the value of property "AudioListener.forwardY".
//
// It returns ok=false if there is no such property.
func (this AudioListener) ForwardY() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerForwardY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ForwardZ returns the value of property "AudioListener.forwardZ".
//
// It returns ok=false if there is no such property.
func (this AudioListener) ForwardZ() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerForwardZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UpX returns the value of property "AudioListener.upX".
//
// It returns ok=false if there is no such property.
func (this AudioListener) UpX() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerUpX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UpY returns the value of property "AudioListener.upY".
//
// It returns ok=false if there is no such property.
func (this AudioListener) UpY() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerUpY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// UpZ returns the value of property "AudioListener.upZ".
//
// It returns ok=false if there is no such property.
func (this AudioListener) UpZ() (ret AudioParam, ok bool) {
	ok = js.True == bindings.GetAudioListenerUpZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSetPosition returns true if the method "AudioListener.setPosition" exists.
func (this AudioListener) HasFuncSetPosition() bool {
	return js.True == bindings.HasFuncAudioListenerSetPosition(
		this.ref,
	)
}

// FuncSetPosition returns the method "AudioListener.setPosition".
func (this AudioListener) FuncSetPosition() (fn js.Func[func(x float32, y float32, z float32)]) {
	bindings.FuncAudioListenerSetPosition(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetPosition calls the method "AudioListener.setPosition".
func (this AudioListener) SetPosition(x float32, y float32, z float32) (ret js.Void) {
	bindings.CallAudioListenerSetPosition(
		this.ref, js.Pointer(&ret),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// TrySetPosition calls the method "AudioListener.setPosition"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioListener) TrySetPosition(x float32, y float32, z float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioListenerSetPosition(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float32(x),
		float32(y),
		float32(z),
	)

	return
}

// HasFuncSetOrientation returns true if the method "AudioListener.setOrientation" exists.
func (this AudioListener) HasFuncSetOrientation() bool {
	return js.True == bindings.HasFuncAudioListenerSetOrientation(
		this.ref,
	)
}

// FuncSetOrientation returns the method "AudioListener.setOrientation".
func (this AudioListener) FuncSetOrientation() (fn js.Func[func(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32)]) {
	bindings.FuncAudioListenerSetOrientation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetOrientation calls the method "AudioListener.setOrientation".
func (this AudioListener) SetOrientation(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32) (ret js.Void) {
	bindings.CallAudioListenerSetOrientation(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AudioListener) TrySetOrientation(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAudioListenerSetOrientation(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *StructuredSerializeOptions) UpdateFrom(ref js.Ref) {
	bindings.StructuredSerializeOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StructuredSerializeOptions) Update(ref js.Ref) {
	bindings.StructuredSerializeOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StructuredSerializeOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Transfer.Ref(),
	)
	p.Transfer = p.Transfer.FromRef(js.Undefined)
}

type MessagePort struct {
	EventTarget
}

func (this MessagePort) Once() MessagePort {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncPostMessage returns true if the method "MessagePort.postMessage" exists.
func (this MessagePort) HasFuncPostMessage() bool {
	return js.True == bindings.HasFuncMessagePortPostMessage(
		this.ref,
	)
}

// FuncPostMessage returns the method "MessagePort.postMessage".
func (this MessagePort) FuncPostMessage() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	bindings.FuncMessagePortPostMessage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage calls the method "MessagePort.postMessage".
func (this MessagePort) PostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void) {
	bindings.CallMessagePortPostMessage(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// TryPostMessage calls the method "MessagePort.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessagePort) TryPostMessage(message js.Any, transfer js.Array[js.Object]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessagePortPostMessage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		transfer.Ref(),
	)

	return
}

// HasFuncPostMessage1 returns true if the method "MessagePort.postMessage" exists.
func (this MessagePort) HasFuncPostMessage1() bool {
	return js.True == bindings.HasFuncMessagePortPostMessage1(
		this.ref,
	)
}

// FuncPostMessage1 returns the method "MessagePort.postMessage".
func (this MessagePort) FuncPostMessage1() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	bindings.FuncMessagePortPostMessage1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage1 calls the method "MessagePort.postMessage".
func (this MessagePort) PostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void) {
	bindings.CallMessagePortPostMessage1(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage1 calls the method "MessagePort.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessagePort) TryPostMessage1(message js.Any, options StructuredSerializeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessagePortPostMessage1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncPostMessage2 returns true if the method "MessagePort.postMessage" exists.
func (this MessagePort) HasFuncPostMessage2() bool {
	return js.True == bindings.HasFuncMessagePortPostMessage2(
		this.ref,
	)
}

// FuncPostMessage2 returns the method "MessagePort.postMessage".
func (this MessagePort) FuncPostMessage2() (fn js.Func[func(message js.Any)]) {
	bindings.FuncMessagePortPostMessage2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage2 calls the method "MessagePort.postMessage".
func (this MessagePort) PostMessage2(message js.Any) (ret js.Void) {
	bindings.CallMessagePortPostMessage2(
		this.ref, js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage2 calls the method "MessagePort.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessagePort) TryPostMessage2(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessagePortPostMessage2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncStart returns true if the method "MessagePort.start" exists.
func (this MessagePort) HasFuncStart() bool {
	return js.True == bindings.HasFuncMessagePortStart(
		this.ref,
	)
}

// FuncStart returns the method "MessagePort.start".
func (this MessagePort) FuncStart() (fn js.Func[func()]) {
	bindings.FuncMessagePortStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start calls the method "MessagePort.start".
func (this MessagePort) Start() (ret js.Void) {
	bindings.CallMessagePortStart(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStart calls the method "MessagePort.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessagePort) TryStart() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessagePortStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "MessagePort.close" exists.
func (this MessagePort) HasFuncClose() bool {
	return js.True == bindings.HasFuncMessagePortClose(
		this.ref,
	)
}

// FuncClose returns the method "MessagePort.close".
func (this MessagePort) FuncClose() (fn js.Func[func()]) {
	bindings.FuncMessagePortClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "MessagePort.close".
func (this MessagePort) Close() (ret js.Void) {
	bindings.CallMessagePortClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "MessagePort.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessagePort) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessagePortClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type AudioWorklet struct {
	Worklet
}

func (this AudioWorklet) Once() AudioWorklet {
	this.ref.Once()
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
	this.ref.Free()
}

// Port returns the value of property "AudioWorklet.port".
//
// It returns ok=false if there is no such property.
func (this AudioWorklet) Port() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetAudioWorkletPort(
		this.ref, js.Pointer(&ret),
	)
	return
}

type BaseAudioContext struct {
	EventTarget
}

func (this BaseAudioContext) Once() BaseAudioContext {
	this.ref.Once()
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
	this.ref.Free()
}

// Destination returns the value of property "BaseAudioContext.destination".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) Destination() (ret AudioDestinationNode, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextDestination(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SampleRate returns the value of property "BaseAudioContext.sampleRate".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) SampleRate() (ret float32, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextSampleRate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CurrentTime returns the value of property "BaseAudioContext.currentTime".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) CurrentTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextCurrentTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Listener returns the value of property "BaseAudioContext.listener".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) Listener() (ret AudioListener, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextListener(
		this.ref, js.Pointer(&ret),
	)
	return
}

// State returns the value of property "BaseAudioContext.state".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) State() (ret AudioContextState, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AudioWorklet returns the value of property "BaseAudioContext.audioWorklet".
//
// It returns ok=false if there is no such property.
func (this BaseAudioContext) AudioWorklet() (ret AudioWorklet, ok bool) {
	ok = js.True == bindings.GetBaseAudioContextAudioWorklet(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncCreateAnalyser returns true if the method "BaseAudioContext.createAnalyser" exists.
func (this BaseAudioContext) HasFuncCreateAnalyser() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateAnalyser(
		this.ref,
	)
}

// FuncCreateAnalyser returns the method "BaseAudioContext.createAnalyser".
func (this BaseAudioContext) FuncCreateAnalyser() (fn js.Func[func() AnalyserNode]) {
	bindings.FuncBaseAudioContextCreateAnalyser(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateAnalyser calls the method "BaseAudioContext.createAnalyser".
func (this BaseAudioContext) CreateAnalyser() (ret AnalyserNode) {
	bindings.CallBaseAudioContextCreateAnalyser(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateAnalyser calls the method "BaseAudioContext.createAnalyser"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateAnalyser() (ret AnalyserNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateAnalyser(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateBiquadFilter returns true if the method "BaseAudioContext.createBiquadFilter" exists.
func (this BaseAudioContext) HasFuncCreateBiquadFilter() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateBiquadFilter(
		this.ref,
	)
}

// FuncCreateBiquadFilter returns the method "BaseAudioContext.createBiquadFilter".
func (this BaseAudioContext) FuncCreateBiquadFilter() (fn js.Func[func() BiquadFilterNode]) {
	bindings.FuncBaseAudioContextCreateBiquadFilter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateBiquadFilter calls the method "BaseAudioContext.createBiquadFilter".
func (this BaseAudioContext) CreateBiquadFilter() (ret BiquadFilterNode) {
	bindings.CallBaseAudioContextCreateBiquadFilter(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateBiquadFilter calls the method "BaseAudioContext.createBiquadFilter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateBiquadFilter() (ret BiquadFilterNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateBiquadFilter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateBuffer returns true if the method "BaseAudioContext.createBuffer" exists.
func (this BaseAudioContext) HasFuncCreateBuffer() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateBuffer(
		this.ref,
	)
}

// FuncCreateBuffer returns the method "BaseAudioContext.createBuffer".
func (this BaseAudioContext) FuncCreateBuffer() (fn js.Func[func(numberOfChannels uint32, length uint32, sampleRate float32) AudioBuffer]) {
	bindings.FuncBaseAudioContextCreateBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateBuffer calls the method "BaseAudioContext.createBuffer".
func (this BaseAudioContext) CreateBuffer(numberOfChannels uint32, length uint32, sampleRate float32) (ret AudioBuffer) {
	bindings.CallBaseAudioContextCreateBuffer(
		this.ref, js.Pointer(&ret),
		uint32(numberOfChannels),
		uint32(length),
		float32(sampleRate),
	)

	return
}

// TryCreateBuffer calls the method "BaseAudioContext.createBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateBuffer(numberOfChannels uint32, length uint32, sampleRate float32) (ret AudioBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(numberOfChannels),
		uint32(length),
		float32(sampleRate),
	)

	return
}

// HasFuncCreateBufferSource returns true if the method "BaseAudioContext.createBufferSource" exists.
func (this BaseAudioContext) HasFuncCreateBufferSource() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateBufferSource(
		this.ref,
	)
}

// FuncCreateBufferSource returns the method "BaseAudioContext.createBufferSource".
func (this BaseAudioContext) FuncCreateBufferSource() (fn js.Func[func() AudioBufferSourceNode]) {
	bindings.FuncBaseAudioContextCreateBufferSource(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateBufferSource calls the method "BaseAudioContext.createBufferSource".
func (this BaseAudioContext) CreateBufferSource() (ret AudioBufferSourceNode) {
	bindings.CallBaseAudioContextCreateBufferSource(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateBufferSource calls the method "BaseAudioContext.createBufferSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateBufferSource() (ret AudioBufferSourceNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateBufferSource(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateChannelMerger returns true if the method "BaseAudioContext.createChannelMerger" exists.
func (this BaseAudioContext) HasFuncCreateChannelMerger() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateChannelMerger(
		this.ref,
	)
}

// FuncCreateChannelMerger returns the method "BaseAudioContext.createChannelMerger".
func (this BaseAudioContext) FuncCreateChannelMerger() (fn js.Func[func(numberOfInputs uint32) ChannelMergerNode]) {
	bindings.FuncBaseAudioContextCreateChannelMerger(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateChannelMerger calls the method "BaseAudioContext.createChannelMerger".
func (this BaseAudioContext) CreateChannelMerger(numberOfInputs uint32) (ret ChannelMergerNode) {
	bindings.CallBaseAudioContextCreateChannelMerger(
		this.ref, js.Pointer(&ret),
		uint32(numberOfInputs),
	)

	return
}

// TryCreateChannelMerger calls the method "BaseAudioContext.createChannelMerger"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateChannelMerger(numberOfInputs uint32) (ret ChannelMergerNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateChannelMerger(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(numberOfInputs),
	)

	return
}

// HasFuncCreateChannelMerger1 returns true if the method "BaseAudioContext.createChannelMerger" exists.
func (this BaseAudioContext) HasFuncCreateChannelMerger1() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateChannelMerger1(
		this.ref,
	)
}

// FuncCreateChannelMerger1 returns the method "BaseAudioContext.createChannelMerger".
func (this BaseAudioContext) FuncCreateChannelMerger1() (fn js.Func[func() ChannelMergerNode]) {
	bindings.FuncBaseAudioContextCreateChannelMerger1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateChannelMerger1 calls the method "BaseAudioContext.createChannelMerger".
func (this BaseAudioContext) CreateChannelMerger1() (ret ChannelMergerNode) {
	bindings.CallBaseAudioContextCreateChannelMerger1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateChannelMerger1 calls the method "BaseAudioContext.createChannelMerger"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateChannelMerger1() (ret ChannelMergerNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateChannelMerger1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateChannelSplitter returns true if the method "BaseAudioContext.createChannelSplitter" exists.
func (this BaseAudioContext) HasFuncCreateChannelSplitter() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateChannelSplitter(
		this.ref,
	)
}

// FuncCreateChannelSplitter returns the method "BaseAudioContext.createChannelSplitter".
func (this BaseAudioContext) FuncCreateChannelSplitter() (fn js.Func[func(numberOfOutputs uint32) ChannelSplitterNode]) {
	bindings.FuncBaseAudioContextCreateChannelSplitter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateChannelSplitter calls the method "BaseAudioContext.createChannelSplitter".
func (this BaseAudioContext) CreateChannelSplitter(numberOfOutputs uint32) (ret ChannelSplitterNode) {
	bindings.CallBaseAudioContextCreateChannelSplitter(
		this.ref, js.Pointer(&ret),
		uint32(numberOfOutputs),
	)

	return
}

// TryCreateChannelSplitter calls the method "BaseAudioContext.createChannelSplitter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateChannelSplitter(numberOfOutputs uint32) (ret ChannelSplitterNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateChannelSplitter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(numberOfOutputs),
	)

	return
}

// HasFuncCreateChannelSplitter1 returns true if the method "BaseAudioContext.createChannelSplitter" exists.
func (this BaseAudioContext) HasFuncCreateChannelSplitter1() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateChannelSplitter1(
		this.ref,
	)
}

// FuncCreateChannelSplitter1 returns the method "BaseAudioContext.createChannelSplitter".
func (this BaseAudioContext) FuncCreateChannelSplitter1() (fn js.Func[func() ChannelSplitterNode]) {
	bindings.FuncBaseAudioContextCreateChannelSplitter1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateChannelSplitter1 calls the method "BaseAudioContext.createChannelSplitter".
func (this BaseAudioContext) CreateChannelSplitter1() (ret ChannelSplitterNode) {
	bindings.CallBaseAudioContextCreateChannelSplitter1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateChannelSplitter1 calls the method "BaseAudioContext.createChannelSplitter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateChannelSplitter1() (ret ChannelSplitterNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateChannelSplitter1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateConstantSource returns true if the method "BaseAudioContext.createConstantSource" exists.
func (this BaseAudioContext) HasFuncCreateConstantSource() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateConstantSource(
		this.ref,
	)
}

// FuncCreateConstantSource returns the method "BaseAudioContext.createConstantSource".
func (this BaseAudioContext) FuncCreateConstantSource() (fn js.Func[func() ConstantSourceNode]) {
	bindings.FuncBaseAudioContextCreateConstantSource(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateConstantSource calls the method "BaseAudioContext.createConstantSource".
func (this BaseAudioContext) CreateConstantSource() (ret ConstantSourceNode) {
	bindings.CallBaseAudioContextCreateConstantSource(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateConstantSource calls the method "BaseAudioContext.createConstantSource"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateConstantSource() (ret ConstantSourceNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateConstantSource(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateConvolver returns true if the method "BaseAudioContext.createConvolver" exists.
func (this BaseAudioContext) HasFuncCreateConvolver() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateConvolver(
		this.ref,
	)
}

// FuncCreateConvolver returns the method "BaseAudioContext.createConvolver".
func (this BaseAudioContext) FuncCreateConvolver() (fn js.Func[func() ConvolverNode]) {
	bindings.FuncBaseAudioContextCreateConvolver(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateConvolver calls the method "BaseAudioContext.createConvolver".
func (this BaseAudioContext) CreateConvolver() (ret ConvolverNode) {
	bindings.CallBaseAudioContextCreateConvolver(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateConvolver calls the method "BaseAudioContext.createConvolver"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateConvolver() (ret ConvolverNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateConvolver(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateDelay returns true if the method "BaseAudioContext.createDelay" exists.
func (this BaseAudioContext) HasFuncCreateDelay() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateDelay(
		this.ref,
	)
}

// FuncCreateDelay returns the method "BaseAudioContext.createDelay".
func (this BaseAudioContext) FuncCreateDelay() (fn js.Func[func(maxDelayTime float64) DelayNode]) {
	bindings.FuncBaseAudioContextCreateDelay(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateDelay calls the method "BaseAudioContext.createDelay".
func (this BaseAudioContext) CreateDelay(maxDelayTime float64) (ret DelayNode) {
	bindings.CallBaseAudioContextCreateDelay(
		this.ref, js.Pointer(&ret),
		float64(maxDelayTime),
	)

	return
}

// TryCreateDelay calls the method "BaseAudioContext.createDelay"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateDelay(maxDelayTime float64) (ret DelayNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateDelay(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(maxDelayTime),
	)

	return
}

// HasFuncCreateDelay1 returns true if the method "BaseAudioContext.createDelay" exists.
func (this BaseAudioContext) HasFuncCreateDelay1() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateDelay1(
		this.ref,
	)
}

// FuncCreateDelay1 returns the method "BaseAudioContext.createDelay".
func (this BaseAudioContext) FuncCreateDelay1() (fn js.Func[func() DelayNode]) {
	bindings.FuncBaseAudioContextCreateDelay1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateDelay1 calls the method "BaseAudioContext.createDelay".
func (this BaseAudioContext) CreateDelay1() (ret DelayNode) {
	bindings.CallBaseAudioContextCreateDelay1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateDelay1 calls the method "BaseAudioContext.createDelay"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateDelay1() (ret DelayNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateDelay1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateDynamicsCompressor returns true if the method "BaseAudioContext.createDynamicsCompressor" exists.
func (this BaseAudioContext) HasFuncCreateDynamicsCompressor() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateDynamicsCompressor(
		this.ref,
	)
}

// FuncCreateDynamicsCompressor returns the method "BaseAudioContext.createDynamicsCompressor".
func (this BaseAudioContext) FuncCreateDynamicsCompressor() (fn js.Func[func() DynamicsCompressorNode]) {
	bindings.FuncBaseAudioContextCreateDynamicsCompressor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateDynamicsCompressor calls the method "BaseAudioContext.createDynamicsCompressor".
func (this BaseAudioContext) CreateDynamicsCompressor() (ret DynamicsCompressorNode) {
	bindings.CallBaseAudioContextCreateDynamicsCompressor(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateDynamicsCompressor calls the method "BaseAudioContext.createDynamicsCompressor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateDynamicsCompressor() (ret DynamicsCompressorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateDynamicsCompressor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateGain returns true if the method "BaseAudioContext.createGain" exists.
func (this BaseAudioContext) HasFuncCreateGain() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateGain(
		this.ref,
	)
}

// FuncCreateGain returns the method "BaseAudioContext.createGain".
func (this BaseAudioContext) FuncCreateGain() (fn js.Func[func() GainNode]) {
	bindings.FuncBaseAudioContextCreateGain(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateGain calls the method "BaseAudioContext.createGain".
func (this BaseAudioContext) CreateGain() (ret GainNode) {
	bindings.CallBaseAudioContextCreateGain(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateGain calls the method "BaseAudioContext.createGain"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateGain() (ret GainNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateGain(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateIIRFilter returns true if the method "BaseAudioContext.createIIRFilter" exists.
func (this BaseAudioContext) HasFuncCreateIIRFilter() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateIIRFilter(
		this.ref,
	)
}

// FuncCreateIIRFilter returns the method "BaseAudioContext.createIIRFilter".
func (this BaseAudioContext) FuncCreateIIRFilter() (fn js.Func[func(feedforward js.Array[float64], feedback js.Array[float64]) IIRFilterNode]) {
	bindings.FuncBaseAudioContextCreateIIRFilter(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateIIRFilter calls the method "BaseAudioContext.createIIRFilter".
func (this BaseAudioContext) CreateIIRFilter(feedforward js.Array[float64], feedback js.Array[float64]) (ret IIRFilterNode) {
	bindings.CallBaseAudioContextCreateIIRFilter(
		this.ref, js.Pointer(&ret),
		feedforward.Ref(),
		feedback.Ref(),
	)

	return
}

// TryCreateIIRFilter calls the method "BaseAudioContext.createIIRFilter"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateIIRFilter(feedforward js.Array[float64], feedback js.Array[float64]) (ret IIRFilterNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateIIRFilter(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		feedforward.Ref(),
		feedback.Ref(),
	)

	return
}

// HasFuncCreateOscillator returns true if the method "BaseAudioContext.createOscillator" exists.
func (this BaseAudioContext) HasFuncCreateOscillator() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateOscillator(
		this.ref,
	)
}

// FuncCreateOscillator returns the method "BaseAudioContext.createOscillator".
func (this BaseAudioContext) FuncCreateOscillator() (fn js.Func[func() OscillatorNode]) {
	bindings.FuncBaseAudioContextCreateOscillator(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateOscillator calls the method "BaseAudioContext.createOscillator".
func (this BaseAudioContext) CreateOscillator() (ret OscillatorNode) {
	bindings.CallBaseAudioContextCreateOscillator(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateOscillator calls the method "BaseAudioContext.createOscillator"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateOscillator() (ret OscillatorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateOscillator(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreatePanner returns true if the method "BaseAudioContext.createPanner" exists.
func (this BaseAudioContext) HasFuncCreatePanner() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreatePanner(
		this.ref,
	)
}

// FuncCreatePanner returns the method "BaseAudioContext.createPanner".
func (this BaseAudioContext) FuncCreatePanner() (fn js.Func[func() PannerNode]) {
	bindings.FuncBaseAudioContextCreatePanner(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreatePanner calls the method "BaseAudioContext.createPanner".
func (this BaseAudioContext) CreatePanner() (ret PannerNode) {
	bindings.CallBaseAudioContextCreatePanner(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreatePanner calls the method "BaseAudioContext.createPanner"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreatePanner() (ret PannerNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreatePanner(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreatePeriodicWave returns true if the method "BaseAudioContext.createPeriodicWave" exists.
func (this BaseAudioContext) HasFuncCreatePeriodicWave() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreatePeriodicWave(
		this.ref,
	)
}

// FuncCreatePeriodicWave returns the method "BaseAudioContext.createPeriodicWave".
func (this BaseAudioContext) FuncCreatePeriodicWave() (fn js.Func[func(real js.Array[float32], imag js.Array[float32], constraints PeriodicWaveConstraints) PeriodicWave]) {
	bindings.FuncBaseAudioContextCreatePeriodicWave(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreatePeriodicWave calls the method "BaseAudioContext.createPeriodicWave".
func (this BaseAudioContext) CreatePeriodicWave(real js.Array[float32], imag js.Array[float32], constraints PeriodicWaveConstraints) (ret PeriodicWave) {
	bindings.CallBaseAudioContextCreatePeriodicWave(
		this.ref, js.Pointer(&ret),
		real.Ref(),
		imag.Ref(),
		js.Pointer(&constraints),
	)

	return
}

// TryCreatePeriodicWave calls the method "BaseAudioContext.createPeriodicWave"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreatePeriodicWave(real js.Array[float32], imag js.Array[float32], constraints PeriodicWaveConstraints) (ret PeriodicWave, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreatePeriodicWave(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		real.Ref(),
		imag.Ref(),
		js.Pointer(&constraints),
	)

	return
}

// HasFuncCreatePeriodicWave1 returns true if the method "BaseAudioContext.createPeriodicWave" exists.
func (this BaseAudioContext) HasFuncCreatePeriodicWave1() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreatePeriodicWave1(
		this.ref,
	)
}

// FuncCreatePeriodicWave1 returns the method "BaseAudioContext.createPeriodicWave".
func (this BaseAudioContext) FuncCreatePeriodicWave1() (fn js.Func[func(real js.Array[float32], imag js.Array[float32]) PeriodicWave]) {
	bindings.FuncBaseAudioContextCreatePeriodicWave1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreatePeriodicWave1 calls the method "BaseAudioContext.createPeriodicWave".
func (this BaseAudioContext) CreatePeriodicWave1(real js.Array[float32], imag js.Array[float32]) (ret PeriodicWave) {
	bindings.CallBaseAudioContextCreatePeriodicWave1(
		this.ref, js.Pointer(&ret),
		real.Ref(),
		imag.Ref(),
	)

	return
}

// TryCreatePeriodicWave1 calls the method "BaseAudioContext.createPeriodicWave"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreatePeriodicWave1(real js.Array[float32], imag js.Array[float32]) (ret PeriodicWave, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreatePeriodicWave1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		real.Ref(),
		imag.Ref(),
	)

	return
}

// HasFuncCreateScriptProcessor returns true if the method "BaseAudioContext.createScriptProcessor" exists.
func (this BaseAudioContext) HasFuncCreateScriptProcessor() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateScriptProcessor(
		this.ref,
	)
}

// FuncCreateScriptProcessor returns the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) FuncCreateScriptProcessor() (fn js.Func[func(bufferSize uint32, numberOfInputChannels uint32, numberOfOutputChannels uint32) ScriptProcessorNode]) {
	bindings.FuncBaseAudioContextCreateScriptProcessor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateScriptProcessor calls the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessor(bufferSize uint32, numberOfInputChannels uint32, numberOfOutputChannels uint32) (ret ScriptProcessorNode) {
	bindings.CallBaseAudioContextCreateScriptProcessor(
		this.ref, js.Pointer(&ret),
		uint32(bufferSize),
		uint32(numberOfInputChannels),
		uint32(numberOfOutputChannels),
	)

	return
}

// TryCreateScriptProcessor calls the method "BaseAudioContext.createScriptProcessor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateScriptProcessor(bufferSize uint32, numberOfInputChannels uint32, numberOfOutputChannels uint32) (ret ScriptProcessorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateScriptProcessor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(bufferSize),
		uint32(numberOfInputChannels),
		uint32(numberOfOutputChannels),
	)

	return
}

// HasFuncCreateScriptProcessor1 returns true if the method "BaseAudioContext.createScriptProcessor" exists.
func (this BaseAudioContext) HasFuncCreateScriptProcessor1() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateScriptProcessor1(
		this.ref,
	)
}

// FuncCreateScriptProcessor1 returns the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) FuncCreateScriptProcessor1() (fn js.Func[func(bufferSize uint32, numberOfInputChannels uint32) ScriptProcessorNode]) {
	bindings.FuncBaseAudioContextCreateScriptProcessor1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateScriptProcessor1 calls the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessor1(bufferSize uint32, numberOfInputChannels uint32) (ret ScriptProcessorNode) {
	bindings.CallBaseAudioContextCreateScriptProcessor1(
		this.ref, js.Pointer(&ret),
		uint32(bufferSize),
		uint32(numberOfInputChannels),
	)

	return
}

// TryCreateScriptProcessor1 calls the method "BaseAudioContext.createScriptProcessor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateScriptProcessor1(bufferSize uint32, numberOfInputChannels uint32) (ret ScriptProcessorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateScriptProcessor1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(bufferSize),
		uint32(numberOfInputChannels),
	)

	return
}

// HasFuncCreateScriptProcessor2 returns true if the method "BaseAudioContext.createScriptProcessor" exists.
func (this BaseAudioContext) HasFuncCreateScriptProcessor2() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateScriptProcessor2(
		this.ref,
	)
}

// FuncCreateScriptProcessor2 returns the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) FuncCreateScriptProcessor2() (fn js.Func[func(bufferSize uint32) ScriptProcessorNode]) {
	bindings.FuncBaseAudioContextCreateScriptProcessor2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateScriptProcessor2 calls the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessor2(bufferSize uint32) (ret ScriptProcessorNode) {
	bindings.CallBaseAudioContextCreateScriptProcessor2(
		this.ref, js.Pointer(&ret),
		uint32(bufferSize),
	)

	return
}

// TryCreateScriptProcessor2 calls the method "BaseAudioContext.createScriptProcessor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateScriptProcessor2(bufferSize uint32) (ret ScriptProcessorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateScriptProcessor2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(bufferSize),
	)

	return
}

// HasFuncCreateScriptProcessor3 returns true if the method "BaseAudioContext.createScriptProcessor" exists.
func (this BaseAudioContext) HasFuncCreateScriptProcessor3() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateScriptProcessor3(
		this.ref,
	)
}

// FuncCreateScriptProcessor3 returns the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) FuncCreateScriptProcessor3() (fn js.Func[func() ScriptProcessorNode]) {
	bindings.FuncBaseAudioContextCreateScriptProcessor3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateScriptProcessor3 calls the method "BaseAudioContext.createScriptProcessor".
func (this BaseAudioContext) CreateScriptProcessor3() (ret ScriptProcessorNode) {
	bindings.CallBaseAudioContextCreateScriptProcessor3(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateScriptProcessor3 calls the method "BaseAudioContext.createScriptProcessor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateScriptProcessor3() (ret ScriptProcessorNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateScriptProcessor3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateStereoPanner returns true if the method "BaseAudioContext.createStereoPanner" exists.
func (this BaseAudioContext) HasFuncCreateStereoPanner() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateStereoPanner(
		this.ref,
	)
}

// FuncCreateStereoPanner returns the method "BaseAudioContext.createStereoPanner".
func (this BaseAudioContext) FuncCreateStereoPanner() (fn js.Func[func() StereoPannerNode]) {
	bindings.FuncBaseAudioContextCreateStereoPanner(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateStereoPanner calls the method "BaseAudioContext.createStereoPanner".
func (this BaseAudioContext) CreateStereoPanner() (ret StereoPannerNode) {
	bindings.CallBaseAudioContextCreateStereoPanner(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateStereoPanner calls the method "BaseAudioContext.createStereoPanner"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateStereoPanner() (ret StereoPannerNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateStereoPanner(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateWaveShaper returns true if the method "BaseAudioContext.createWaveShaper" exists.
func (this BaseAudioContext) HasFuncCreateWaveShaper() bool {
	return js.True == bindings.HasFuncBaseAudioContextCreateWaveShaper(
		this.ref,
	)
}

// FuncCreateWaveShaper returns the method "BaseAudioContext.createWaveShaper".
func (this BaseAudioContext) FuncCreateWaveShaper() (fn js.Func[func() WaveShaperNode]) {
	bindings.FuncBaseAudioContextCreateWaveShaper(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateWaveShaper calls the method "BaseAudioContext.createWaveShaper".
func (this BaseAudioContext) CreateWaveShaper() (ret WaveShaperNode) {
	bindings.CallBaseAudioContextCreateWaveShaper(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateWaveShaper calls the method "BaseAudioContext.createWaveShaper"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryCreateWaveShaper() (ret WaveShaperNode, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextCreateWaveShaper(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDecodeAudioData returns true if the method "BaseAudioContext.decodeAudioData" exists.
func (this BaseAudioContext) HasFuncDecodeAudioData() bool {
	return js.True == bindings.HasFuncBaseAudioContextDecodeAudioData(
		this.ref,
	)
}

// FuncDecodeAudioData returns the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) FuncDecodeAudioData() (fn js.Func[func(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)], errorCallback js.Func[func(err DOMException)]) js.Promise[AudioBuffer]]) {
	bindings.FuncBaseAudioContextDecodeAudioData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DecodeAudioData calls the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) DecodeAudioData(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)], errorCallback js.Func[func(err DOMException)]) (ret js.Promise[AudioBuffer]) {
	bindings.CallBaseAudioContextDecodeAudioData(
		this.ref, js.Pointer(&ret),
		audioData.Ref(),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// TryDecodeAudioData calls the method "BaseAudioContext.decodeAudioData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryDecodeAudioData(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)], errorCallback js.Func[func(err DOMException)]) (ret js.Promise[AudioBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextDecodeAudioData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		audioData.Ref(),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasFuncDecodeAudioData1 returns true if the method "BaseAudioContext.decodeAudioData" exists.
func (this BaseAudioContext) HasFuncDecodeAudioData1() bool {
	return js.True == bindings.HasFuncBaseAudioContextDecodeAudioData1(
		this.ref,
	)
}

// FuncDecodeAudioData1 returns the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) FuncDecodeAudioData1() (fn js.Func[func(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)]) js.Promise[AudioBuffer]]) {
	bindings.FuncBaseAudioContextDecodeAudioData1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DecodeAudioData1 calls the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) DecodeAudioData1(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)]) (ret js.Promise[AudioBuffer]) {
	bindings.CallBaseAudioContextDecodeAudioData1(
		this.ref, js.Pointer(&ret),
		audioData.Ref(),
		successCallback.Ref(),
	)

	return
}

// TryDecodeAudioData1 calls the method "BaseAudioContext.decodeAudioData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryDecodeAudioData1(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)]) (ret js.Promise[AudioBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextDecodeAudioData1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		audioData.Ref(),
		successCallback.Ref(),
	)

	return
}

// HasFuncDecodeAudioData2 returns true if the method "BaseAudioContext.decodeAudioData" exists.
func (this BaseAudioContext) HasFuncDecodeAudioData2() bool {
	return js.True == bindings.HasFuncBaseAudioContextDecodeAudioData2(
		this.ref,
	)
}

// FuncDecodeAudioData2 returns the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) FuncDecodeAudioData2() (fn js.Func[func(audioData js.ArrayBuffer) js.Promise[AudioBuffer]]) {
	bindings.FuncBaseAudioContextDecodeAudioData2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DecodeAudioData2 calls the method "BaseAudioContext.decodeAudioData".
func (this BaseAudioContext) DecodeAudioData2(audioData js.ArrayBuffer) (ret js.Promise[AudioBuffer]) {
	bindings.CallBaseAudioContextDecodeAudioData2(
		this.ref, js.Pointer(&ret),
		audioData.Ref(),
	)

	return
}

// TryDecodeAudioData2 calls the method "BaseAudioContext.decodeAudioData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this BaseAudioContext) TryDecodeAudioData2(audioData js.ArrayBuffer) (ret js.Promise[AudioBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryBaseAudioContextDecodeAudioData2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *AnalyserOptions) UpdateFrom(ref js.Ref) {
	bindings.AnalyserOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AnalyserOptions) Update(ref js.Ref) {
	bindings.AnalyserOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AnalyserOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// FftSize returns the value of property "AnalyserNode.fftSize".
//
// It returns ok=false if there is no such property.
func (this AnalyserNode) FftSize() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAnalyserNodeFftSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFftSize sets the value of property "AnalyserNode.fftSize" to val.
//
// It returns false if the property cannot be set.
func (this AnalyserNode) SetFftSize(val uint32) bool {
	return js.True == bindings.SetAnalyserNodeFftSize(
		this.ref,
		uint32(val),
	)
}

// FrequencyBinCount returns the value of property "AnalyserNode.frequencyBinCount".
//
// It returns ok=false if there is no such property.
func (this AnalyserNode) FrequencyBinCount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetAnalyserNodeFrequencyBinCount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MinDecibels returns the value of property "AnalyserNode.minDecibels".
//
// It returns ok=false if there is no such property.
func (this AnalyserNode) MinDecibels() (ret float64, ok bool) {
	ok = js.True == bindings.GetAnalyserNodeMinDecibels(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMinDecibels sets the value of property "AnalyserNode.minDecibels" to val.
//
// It returns false if the property cannot be set.
func (this AnalyserNode) SetMinDecibels(val float64) bool {
	return js.True == bindings.SetAnalyserNodeMinDecibels(
		this.ref,
		float64(val),
	)
}

// MaxDecibels returns the value of property "AnalyserNode.maxDecibels".
//
// It returns ok=false if there is no such property.
func (this AnalyserNode) MaxDecibels() (ret float64, ok bool) {
	ok = js.True == bindings.GetAnalyserNodeMaxDecibels(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMaxDecibels sets the value of property "AnalyserNode.maxDecibels" to val.
//
// It returns false if the property cannot be set.
func (this AnalyserNode) SetMaxDecibels(val float64) bool {
	return js.True == bindings.SetAnalyserNodeMaxDecibels(
		this.ref,
		float64(val),
	)
}

// SmoothingTimeConstant returns the value of property "AnalyserNode.smoothingTimeConstant".
//
// It returns ok=false if there is no such property.
func (this AnalyserNode) SmoothingTimeConstant() (ret float64, ok bool) {
	ok = js.True == bindings.GetAnalyserNodeSmoothingTimeConstant(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSmoothingTimeConstant sets the value of property "AnalyserNode.smoothingTimeConstant" to val.
//
// It returns false if the property cannot be set.
func (this AnalyserNode) SetSmoothingTimeConstant(val float64) bool {
	return js.True == bindings.SetAnalyserNodeSmoothingTimeConstant(
		this.ref,
		float64(val),
	)
}

// HasFuncGetFloatFrequencyData returns true if the method "AnalyserNode.getFloatFrequencyData" exists.
func (this AnalyserNode) HasFuncGetFloatFrequencyData() bool {
	return js.True == bindings.HasFuncAnalyserNodeGetFloatFrequencyData(
		this.ref,
	)
}

// FuncGetFloatFrequencyData returns the method "AnalyserNode.getFloatFrequencyData".
func (this AnalyserNode) FuncGetFloatFrequencyData() (fn js.Func[func(array js.TypedArray[float32])]) {
	bindings.FuncAnalyserNodeGetFloatFrequencyData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFloatFrequencyData calls the method "AnalyserNode.getFloatFrequencyData".
func (this AnalyserNode) GetFloatFrequencyData(array js.TypedArray[float32]) (ret js.Void) {
	bindings.CallAnalyserNodeGetFloatFrequencyData(
		this.ref, js.Pointer(&ret),
		array.Ref(),
	)

	return
}

// TryGetFloatFrequencyData calls the method "AnalyserNode.getFloatFrequencyData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnalyserNode) TryGetFloatFrequencyData(array js.TypedArray[float32]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnalyserNodeGetFloatFrequencyData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		array.Ref(),
	)

	return
}

// HasFuncGetByteFrequencyData returns true if the method "AnalyserNode.getByteFrequencyData" exists.
func (this AnalyserNode) HasFuncGetByteFrequencyData() bool {
	return js.True == bindings.HasFuncAnalyserNodeGetByteFrequencyData(
		this.ref,
	)
}

// FuncGetByteFrequencyData returns the method "AnalyserNode.getByteFrequencyData".
func (this AnalyserNode) FuncGetByteFrequencyData() (fn js.Func[func(array js.TypedArray[uint8])]) {
	bindings.FuncAnalyserNodeGetByteFrequencyData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetByteFrequencyData calls the method "AnalyserNode.getByteFrequencyData".
func (this AnalyserNode) GetByteFrequencyData(array js.TypedArray[uint8]) (ret js.Void) {
	bindings.CallAnalyserNodeGetByteFrequencyData(
		this.ref, js.Pointer(&ret),
		array.Ref(),
	)

	return
}

// TryGetByteFrequencyData calls the method "AnalyserNode.getByteFrequencyData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnalyserNode) TryGetByteFrequencyData(array js.TypedArray[uint8]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnalyserNodeGetByteFrequencyData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		array.Ref(),
	)

	return
}

// HasFuncGetFloatTimeDomainData returns true if the method "AnalyserNode.getFloatTimeDomainData" exists.
func (this AnalyserNode) HasFuncGetFloatTimeDomainData() bool {
	return js.True == bindings.HasFuncAnalyserNodeGetFloatTimeDomainData(
		this.ref,
	)
}

// FuncGetFloatTimeDomainData returns the method "AnalyserNode.getFloatTimeDomainData".
func (this AnalyserNode) FuncGetFloatTimeDomainData() (fn js.Func[func(array js.TypedArray[float32])]) {
	bindings.FuncAnalyserNodeGetFloatTimeDomainData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetFloatTimeDomainData calls the method "AnalyserNode.getFloatTimeDomainData".
func (this AnalyserNode) GetFloatTimeDomainData(array js.TypedArray[float32]) (ret js.Void) {
	bindings.CallAnalyserNodeGetFloatTimeDomainData(
		this.ref, js.Pointer(&ret),
		array.Ref(),
	)

	return
}

// TryGetFloatTimeDomainData calls the method "AnalyserNode.getFloatTimeDomainData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnalyserNode) TryGetFloatTimeDomainData(array js.TypedArray[float32]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnalyserNodeGetFloatTimeDomainData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		array.Ref(),
	)

	return
}

// HasFuncGetByteTimeDomainData returns true if the method "AnalyserNode.getByteTimeDomainData" exists.
func (this AnalyserNode) HasFuncGetByteTimeDomainData() bool {
	return js.True == bindings.HasFuncAnalyserNodeGetByteTimeDomainData(
		this.ref,
	)
}

// FuncGetByteTimeDomainData returns the method "AnalyserNode.getByteTimeDomainData".
func (this AnalyserNode) FuncGetByteTimeDomainData() (fn js.Func[func(array js.TypedArray[uint8])]) {
	bindings.FuncAnalyserNodeGetByteTimeDomainData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetByteTimeDomainData calls the method "AnalyserNode.getByteTimeDomainData".
func (this AnalyserNode) GetByteTimeDomainData(array js.TypedArray[uint8]) (ret js.Void) {
	bindings.CallAnalyserNodeGetByteTimeDomainData(
		this.ref, js.Pointer(&ret),
		array.Ref(),
	)

	return
}

// TryGetByteTimeDomainData calls the method "AnalyserNode.getByteTimeDomainData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnalyserNode) TryGetByteTimeDomainData(array js.TypedArray[uint8]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnalyserNodeGetByteTimeDomainData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *AnimationEventInit) UpdateFrom(ref js.Ref) {
	bindings.AnimationEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AnimationEventInit) Update(ref js.Ref) {
	bindings.AnimationEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AnimationEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.AnimationName.Ref(),
		p.PseudoElement.Ref(),
	)
	p.AnimationName = p.AnimationName.FromRef(js.Undefined)
	p.PseudoElement = p.PseudoElement.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// AnimationName returns the value of property "AnimationEvent.animationName".
//
// It returns ok=false if there is no such property.
func (this AnimationEvent) AnimationName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAnimationEventAnimationName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ElapsedTime returns the value of property "AnimationEvent.elapsedTime".
//
// It returns ok=false if there is no such property.
func (this AnimationEvent) ElapsedTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetAnimationEventElapsedTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PseudoElement returns the value of property "AnimationEvent.pseudoElement".
//
// It returns ok=false if there is no such property.
func (this AnimationEvent) PseudoElement() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAnimationEventPseudoElement(
		this.ref, js.Pointer(&ret),
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
func (p *AnimationPlaybackEventInit) UpdateFrom(ref js.Ref) {
	bindings.AnimationPlaybackEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AnimationPlaybackEventInit) Update(ref js.Ref) {
	bindings.AnimationPlaybackEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AnimationPlaybackEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.CurrentTime.Ref(),
		p.TimelineTime.Ref(),
	)
	p.CurrentTime = p.CurrentTime.FromRef(js.Undefined)
	p.TimelineTime = p.TimelineTime.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// CurrentTime returns the value of property "AnimationPlaybackEvent.currentTime".
//
// It returns ok=false if there is no such property.
func (this AnimationPlaybackEvent) CurrentTime() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationPlaybackEventCurrentTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TimelineTime returns the value of property "AnimationPlaybackEvent.timelineTime".
//
// It returns ok=false if there is no such property.
func (this AnimationPlaybackEvent) TimelineTime() (ret CSSNumberish, ok bool) {
	ok = js.True == bindings.GetAnimationPlaybackEventTimelineTime(
		this.ref, js.Pointer(&ret),
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncRegisterAnimator returns true if the method "AnimationWorkletGlobalScope.registerAnimator" exists.
func (this AnimationWorkletGlobalScope) HasFuncRegisterAnimator() bool {
	return js.True == bindings.HasFuncAnimationWorkletGlobalScopeRegisterAnimator(
		this.ref,
	)
}

// FuncRegisterAnimator returns the method "AnimationWorkletGlobalScope.registerAnimator".
func (this AnimationWorkletGlobalScope) FuncRegisterAnimator() (fn js.Func[func(name js.String, animatorCtor js.Func[func(options js.Any, state js.Any) js.Any])]) {
	bindings.FuncAnimationWorkletGlobalScopeRegisterAnimator(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RegisterAnimator calls the method "AnimationWorkletGlobalScope.registerAnimator".
func (this AnimationWorkletGlobalScope) RegisterAnimator(name js.String, animatorCtor js.Func[func(options js.Any, state js.Any) js.Any]) (ret js.Void) {
	bindings.CallAnimationWorkletGlobalScopeRegisterAnimator(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		animatorCtor.Ref(),
	)

	return
}

// TryRegisterAnimator calls the method "AnimationWorkletGlobalScope.registerAnimator"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AnimationWorkletGlobalScope) TryRegisterAnimator(name js.String, animatorCtor js.Func[func(options js.Any, state js.Any) js.Any]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAnimationWorkletGlobalScopeRegisterAnimator(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *AttributionReportingRequestOptions) UpdateFrom(ref js.Ref) {
	bindings.AttributionReportingRequestOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AttributionReportingRequestOptions) Update(ref js.Ref) {
	bindings.AttributionReportingRequestOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AttributionReportingRequestOptions) FreeMembers(recursive bool) {
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
func (p *AuctionAd) UpdateFrom(ref js.Ref) {
	bindings.AuctionAdJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuctionAd) Update(ref js.Ref) {
	bindings.AuctionAdJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuctionAd) FreeMembers(recursive bool) {
	js.Free(
		p.RenderURL.Ref(),
		p.Metadata.Ref(),
		p.BuyerReportingId.Ref(),
		p.BuyerAndSellerReportingId.Ref(),
	)
	p.RenderURL = p.RenderURL.FromRef(js.Undefined)
	p.Metadata = p.Metadata.FromRef(js.Undefined)
	p.BuyerReportingId = p.BuyerReportingId.FromRef(js.Undefined)
	p.BuyerAndSellerReportingId = p.BuyerAndSellerReportingId.FromRef(js.Undefined)
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
func (p *AuctionAdConfig) UpdateFrom(ref js.Ref) {
	bindings.AuctionAdConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuctionAdConfig) Update(ref js.Ref) {
	bindings.AuctionAdConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuctionAdConfig) FreeMembers(recursive bool) {
	js.Free(
		p.Seller.Ref(),
		p.DecisionLogicURL.Ref(),
		p.TrustedScoringSignalsURL.Ref(),
		p.InterestGroupBuyers.Ref(),
		p.AuctionSignals.Ref(),
		p.SellerSignals.Ref(),
		p.DirectFromSellerSignals.Ref(),
		p.SellerCurrency.Ref(),
		p.PerBuyerSignals.Ref(),
		p.PerBuyerTimeouts.Ref(),
		p.PerBuyerGroupLimits.Ref(),
		p.PerBuyerExperimentGroupIds.Ref(),
		p.PerBuyerPrioritySignals.Ref(),
		p.PerBuyerCurrencies.Ref(),
		p.ComponentAuctions.Ref(),
		p.Signal.Ref(),
		p.ResolveToConfig.Ref(),
	)
	p.Seller = p.Seller.FromRef(js.Undefined)
	p.DecisionLogicURL = p.DecisionLogicURL.FromRef(js.Undefined)
	p.TrustedScoringSignalsURL = p.TrustedScoringSignalsURL.FromRef(js.Undefined)
	p.InterestGroupBuyers = p.InterestGroupBuyers.FromRef(js.Undefined)
	p.AuctionSignals = p.AuctionSignals.FromRef(js.Undefined)
	p.SellerSignals = p.SellerSignals.FromRef(js.Undefined)
	p.DirectFromSellerSignals = p.DirectFromSellerSignals.FromRef(js.Undefined)
	p.SellerCurrency = p.SellerCurrency.FromRef(js.Undefined)
	p.PerBuyerSignals = p.PerBuyerSignals.FromRef(js.Undefined)
	p.PerBuyerTimeouts = p.PerBuyerTimeouts.FromRef(js.Undefined)
	p.PerBuyerGroupLimits = p.PerBuyerGroupLimits.FromRef(js.Undefined)
	p.PerBuyerExperimentGroupIds = p.PerBuyerExperimentGroupIds.FromRef(js.Undefined)
	p.PerBuyerPrioritySignals = p.PerBuyerPrioritySignals.FromRef(js.Undefined)
	p.PerBuyerCurrencies = p.PerBuyerCurrencies.FromRef(js.Undefined)
	p.ComponentAuctions = p.ComponentAuctions.FromRef(js.Undefined)
	p.Signal = p.Signal.FromRef(js.Undefined)
	p.ResolveToConfig = p.ResolveToConfig.FromRef(js.Undefined)
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
func (p *AuctionAdInterestGroup) UpdateFrom(ref js.Ref) {
	bindings.AuctionAdInterestGroupJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuctionAdInterestGroup) Update(ref js.Ref) {
	bindings.AuctionAdInterestGroupJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuctionAdInterestGroup) FreeMembers(recursive bool) {
	js.Free(
		p.PrioritySignalsOverrides.Ref(),
		p.Owner.Ref(),
		p.Name.Ref(),
		p.PriorityVector.Ref(),
		p.ExecutionMode.Ref(),
		p.BiddingLogicURL.Ref(),
		p.BiddingWasmHelperURL.Ref(),
		p.UpdateURL.Ref(),
		p.TrustedBiddingSignalsURL.Ref(),
		p.TrustedBiddingSignalsKeys.Ref(),
		p.UserBiddingSignals.Ref(),
		p.Ads.Ref(),
		p.AdComponents.Ref(),
	)
	p.PrioritySignalsOverrides = p.PrioritySignalsOverrides.FromRef(js.Undefined)
	p.Owner = p.Owner.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.PriorityVector = p.PriorityVector.FromRef(js.Undefined)
	p.ExecutionMode = p.ExecutionMode.FromRef(js.Undefined)
	p.BiddingLogicURL = p.BiddingLogicURL.FromRef(js.Undefined)
	p.BiddingWasmHelperURL = p.BiddingWasmHelperURL.FromRef(js.Undefined)
	p.UpdateURL = p.UpdateURL.FromRef(js.Undefined)
	p.TrustedBiddingSignalsURL = p.TrustedBiddingSignalsURL.FromRef(js.Undefined)
	p.TrustedBiddingSignalsKeys = p.TrustedBiddingSignalsKeys.FromRef(js.Undefined)
	p.UserBiddingSignals = p.UserBiddingSignals.FromRef(js.Undefined)
	p.Ads = p.Ads.FromRef(js.Undefined)
	p.AdComponents = p.AdComponents.FromRef(js.Undefined)
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
func (p *AuctionAdInterestGroupKey) UpdateFrom(ref js.Ref) {
	bindings.AuctionAdInterestGroupKeyJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AuctionAdInterestGroupKey) Update(ref js.Ref) {
	bindings.AuctionAdInterestGroupKeyJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AuctionAdInterestGroupKey) FreeMembers(recursive bool) {
	js.Free(
		p.Owner.Ref(),
		p.Name.Ref(),
	)
	p.Owner = p.Owner.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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
func (p *AudioConfiguration) UpdateFrom(ref js.Ref) {
	bindings.AudioConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioConfiguration) Update(ref js.Ref) {
	bindings.AudioConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioConfiguration) FreeMembers(recursive bool) {
	js.Free(
		p.ContentType.Ref(),
		p.Channels.Ref(),
	)
	p.ContentType = p.ContentType.FromRef(js.Undefined)
	p.Channels = p.Channels.FromRef(js.Undefined)
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
func (p *AudioSinkOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioSinkOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioSinkOptions) Update(ref js.Ref) {
	bindings.AudioSinkOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioSinkOptions) FreeMembers(recursive bool) {
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
func (p *AudioContextOptions) UpdateFrom(ref js.Ref) {
	bindings.AudioContextOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioContextOptions) Update(ref js.Ref) {
	bindings.AudioContextOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioContextOptions) FreeMembers(recursive bool) {
	js.Free(
		p.LatencyHint.Ref(),
		p.SinkId.Ref(),
	)
	p.LatencyHint = p.LatencyHint.FromRef(js.Undefined)
	p.SinkId = p.SinkId.FromRef(js.Undefined)
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
func (p *AudioTimestamp) UpdateFrom(ref js.Ref) {
	bindings.AudioTimestampJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AudioTimestamp) Update(ref js.Ref) {
	bindings.AudioTimestampJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AudioTimestamp) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Track returns the value of property "TextTrackCue.track".
//
// It returns ok=false if there is no such property.
func (this TextTrackCue) Track() (ret TextTrack, ok bool) {
	ok = js.True == bindings.GetTextTrackCueTrack(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "TextTrackCue.id".
//
// It returns ok=false if there is no such property.
func (this TextTrackCue) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTextTrackCueId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetId sets the value of property "TextTrackCue.id" to val.
//
// It returns false if the property cannot be set.
func (this TextTrackCue) SetId(val js.String) bool {
	return js.True == bindings.SetTextTrackCueId(
		this.ref,
		val.Ref(),
	)
}

// StartTime returns the value of property "TextTrackCue.startTime".
//
// It returns ok=false if there is no such property.
func (this TextTrackCue) StartTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextTrackCueStartTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetStartTime sets the value of property "TextTrackCue.startTime" to val.
//
// It returns false if the property cannot be set.
func (this TextTrackCue) SetStartTime(val float64) bool {
	return js.True == bindings.SetTextTrackCueStartTime(
		this.ref,
		float64(val),
	)
}

// EndTime returns the value of property "TextTrackCue.endTime".
//
// It returns ok=false if there is no such property.
func (this TextTrackCue) EndTime() (ret float64, ok bool) {
	ok = js.True == bindings.GetTextTrackCueEndTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEndTime sets the value of property "TextTrackCue.endTime" to val.
//
// It returns false if the property cannot be set.
func (this TextTrackCue) SetEndTime(val float64) bool {
	return js.True == bindings.SetTextTrackCueEndTime(
		this.ref,
		float64(val),
	)
}

// PauseOnExit returns the value of property "TextTrackCue.pauseOnExit".
//
// It returns ok=false if there is no such property.
func (this TextTrackCue) PauseOnExit() (ret bool, ok bool) {
	ok = js.True == bindings.GetTextTrackCuePauseOnExit(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPauseOnExit sets the value of property "TextTrackCue.pauseOnExit" to val.
//
// It returns false if the property cannot be set.
func (this TextTrackCue) SetPauseOnExit(val bool) bool {
	return js.True == bindings.SetTextTrackCuePauseOnExit(
		this.ref,
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

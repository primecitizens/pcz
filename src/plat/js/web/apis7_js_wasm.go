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

type WaveShaperOptions struct {
	// Curve is "WaveShaperOptions.curve"
	//
	// Optional
	Curve js.Array[float32]
	// Oversample is "WaveShaperOptions.oversample"
	//
	// Optional, defaults to "none".
	Oversample OverSampleType
	// ChannelCount is "WaveShaperOptions.channelCount"
	//
	// Optional
	ChannelCount uint32
	// ChannelCountMode is "WaveShaperOptions.channelCountMode"
	//
	// Optional
	ChannelCountMode ChannelCountMode
	// ChannelInterpretation is "WaveShaperOptions.channelInterpretation"
	//
	// Optional
	ChannelInterpretation ChannelInterpretation

	FFI_USE_ChannelCount bool // for ChannelCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WaveShaperOptions with all fields set.
func (p WaveShaperOptions) FromRef(ref js.Ref) WaveShaperOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 WaveShaperOptions WaveShaperOptions [// WaveShaperOptions] [0x140001823c0 0x14000182460 0x14000182500 0x14000182640 0x140001826e0 0x140001825a0] 0x14000aa32c0 {0 0}} in the application heap.
func (p WaveShaperOptions) New() js.Ref {
	return bindings.WaveShaperOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p WaveShaperOptions) UpdateFrom(ref js.Ref) {
	bindings.WaveShaperOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p WaveShaperOptions) Update(ref js.Ref) {
	bindings.WaveShaperOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewWaveShaperNode(context BaseAudioContext, options WaveShaperOptions) WaveShaperNode {
	return WaveShaperNode{}.FromRef(
		bindings.NewWaveShaperNodeByWaveShaperNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewWaveShaperNodeByWaveShaperNode1(context BaseAudioContext) WaveShaperNode {
	return WaveShaperNode{}.FromRef(
		bindings.NewWaveShaperNodeByWaveShaperNode1(
			context.Ref()),
	)
}

type WaveShaperNode struct {
	AudioNode
}

func (this WaveShaperNode) Once() WaveShaperNode {
	this.Ref().Once()
	return this
}

func (this WaveShaperNode) Ref() js.Ref {
	return this.AudioNode.Ref()
}

func (this WaveShaperNode) FromRef(ref js.Ref) WaveShaperNode {
	this.AudioNode = this.AudioNode.FromRef(ref)
	return this
}

func (this WaveShaperNode) Free() {
	this.Ref().Free()
}

// Curve returns the value of property "WaveShaperNode.curve".
//
// The returned bool will be false if there is no such property.
func (this WaveShaperNode) Curve() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetWaveShaperNodeCurve(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// Curve sets the value of property "WaveShaperNode.curve" to val.
//
// It returns false if the property cannot be set.
func (this WaveShaperNode) SetCurve(val js.TypedArray[float32]) bool {
	return js.True == bindings.SetWaveShaperNodeCurve(
		this.Ref(),
		val.Ref(),
	)
}

// Oversample returns the value of property "WaveShaperNode.oversample".
//
// The returned bool will be false if there is no such property.
func (this WaveShaperNode) Oversample() (OverSampleType, bool) {
	var _ok bool
	_ret := bindings.GetWaveShaperNodeOversample(
		this.Ref(), js.Pointer(&_ok),
	)
	return OverSampleType(_ret), _ok
}

// Oversample sets the value of property "WaveShaperNode.oversample" to val.
//
// It returns false if the property cannot be set.
func (this WaveShaperNode) SetOversample(val OverSampleType) bool {
	return js.True == bindings.SetWaveShaperNodeOversample(
		this.Ref(),
		uint32(val),
	)
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

func NewDOMException(message js.String, name js.String) DOMException {
	return DOMException{}.FromRef(
		bindings.NewDOMExceptionByDOMException(
			message.Ref(),
			name.Ref()),
	)
}

func NewDOMExceptionByDOMException1(message js.String) DOMException {
	return DOMException{}.FromRef(
		bindings.NewDOMExceptionByDOMException1(
			message.Ref()),
	)
}

func NewDOMExceptionByDOMException2() DOMException {
	return DOMException{}.FromRef(
		bindings.NewDOMExceptionByDOMException2(),
	)
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
// The returned bool will be false if there is no such property.
func (this DOMException) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDOMExceptionName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Message returns the value of property "DOMException.message".
//
// The returned bool will be false if there is no such property.
func (this DOMException) Message() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetDOMExceptionMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Code returns the value of property "DOMException.code".
//
// The returned bool will be false if there is no such property.
func (this DOMException) Code() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetDOMExceptionCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this AudioDestinationNode) MaxChannelCount() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAudioDestinationNodeMaxChannelCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this AudioListener) PositionX() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetAudioListenerPositionX(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// PositionY returns the value of property "AudioListener.positionY".
//
// The returned bool will be false if there is no such property.
func (this AudioListener) PositionY() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetAudioListenerPositionY(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// PositionZ returns the value of property "AudioListener.positionZ".
//
// The returned bool will be false if there is no such property.
func (this AudioListener) PositionZ() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetAudioListenerPositionZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// ForwardX returns the value of property "AudioListener.forwardX".
//
// The returned bool will be false if there is no such property.
func (this AudioListener) ForwardX() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetAudioListenerForwardX(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// ForwardY returns the value of property "AudioListener.forwardY".
//
// The returned bool will be false if there is no such property.
func (this AudioListener) ForwardY() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetAudioListenerForwardY(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// ForwardZ returns the value of property "AudioListener.forwardZ".
//
// The returned bool will be false if there is no such property.
func (this AudioListener) ForwardZ() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetAudioListenerForwardZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// UpX returns the value of property "AudioListener.upX".
//
// The returned bool will be false if there is no such property.
func (this AudioListener) UpX() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetAudioListenerUpX(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// UpY returns the value of property "AudioListener.upY".
//
// The returned bool will be false if there is no such property.
func (this AudioListener) UpY() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetAudioListenerUpY(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// UpZ returns the value of property "AudioListener.upZ".
//
// The returned bool will be false if there is no such property.
func (this AudioListener) UpZ() (AudioParam, bool) {
	var _ok bool
	_ret := bindings.GetAudioListenerUpZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioParam{}.FromRef(_ret), _ok
}

// SetPosition calls the method "AudioListener.setPosition".
//
// The returned bool will be false if there is no such method.
func (this AudioListener) SetPosition(x float32, y float32, z float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioListenerSetPosition(
		this.Ref(), js.Pointer(&_ok),
		float32(x),
		float32(y),
		float32(z),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPositionFunc returns the method "AudioListener.setPosition".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioListener) SetPositionFunc() (fn js.Func[func(x float32, y float32, z float32)]) {
	return fn.FromRef(
		bindings.AudioListenerSetPositionFunc(
			this.Ref(),
		),
	)
}

// SetOrientation calls the method "AudioListener.setOrientation".
//
// The returned bool will be false if there is no such method.
func (this AudioListener) SetOrientation(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAudioListenerSetOrientation(
		this.Ref(), js.Pointer(&_ok),
		float32(x),
		float32(y),
		float32(z),
		float32(xUp),
		float32(yUp),
		float32(zUp),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetOrientationFunc returns the method "AudioListener.setOrientation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AudioListener) SetOrientationFunc() (fn js.Func[func(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32)]) {
	return fn.FromRef(
		bindings.AudioListenerSetOrientationFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 StructuredSerializeOptions StructuredSerializeOptions [// StructuredSerializeOptions] [0x14000182a00] 0x14000aa33f8 {0 0}} in the application heap.
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

// PostMessage calls the method "MessagePort.postMessage".
//
// The returned bool will be false if there is no such method.
func (this MessagePort) PostMessage(message js.Any, transfer js.Array[js.Object]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessagePortPostMessage(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		transfer.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessageFunc returns the method "MessagePort.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessagePort) PostMessageFunc() (fn js.Func[func(message js.Any, transfer js.Array[js.Object])]) {
	return fn.FromRef(
		bindings.MessagePortPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "MessagePort.postMessage".
//
// The returned bool will be false if there is no such method.
func (this MessagePort) PostMessage1(message js.Any, options StructuredSerializeOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessagePortPostMessage1(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage1Func returns the method "MessagePort.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessagePort) PostMessage1Func() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.MessagePortPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage2 calls the method "MessagePort.postMessage".
//
// The returned bool will be false if there is no such method.
func (this MessagePort) PostMessage2(message js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessagePortPostMessage2(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage2Func returns the method "MessagePort.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessagePort) PostMessage2Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.MessagePortPostMessage2Func(
			this.Ref(),
		),
	)
}

// Start calls the method "MessagePort.start".
//
// The returned bool will be false if there is no such method.
func (this MessagePort) Start() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessagePortStart(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StartFunc returns the method "MessagePort.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessagePort) StartFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MessagePortStartFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "MessagePort.close".
//
// The returned bool will be false if there is no such method.
func (this MessagePort) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessagePortClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "MessagePort.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessagePort) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MessagePortCloseFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this AudioWorklet) Port() (MessagePort, bool) {
	var _ok bool
	_ret := bindings.GetAudioWorkletPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return MessagePort{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this BaseAudioContext) Destination() (AudioDestinationNode, bool) {
	var _ok bool
	_ret := bindings.GetBaseAudioContextDestination(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioDestinationNode{}.FromRef(_ret), _ok
}

// SampleRate returns the value of property "BaseAudioContext.sampleRate".
//
// The returned bool will be false if there is no such property.
func (this BaseAudioContext) SampleRate() (float32, bool) {
	var _ok bool
	_ret := bindings.GetBaseAudioContextSampleRate(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// CurrentTime returns the value of property "BaseAudioContext.currentTime".
//
// The returned bool will be false if there is no such property.
func (this BaseAudioContext) CurrentTime() (float64, bool) {
	var _ok bool
	_ret := bindings.GetBaseAudioContextCurrentTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Listener returns the value of property "BaseAudioContext.listener".
//
// The returned bool will be false if there is no such property.
func (this BaseAudioContext) Listener() (AudioListener, bool) {
	var _ok bool
	_ret := bindings.GetBaseAudioContextListener(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioListener{}.FromRef(_ret), _ok
}

// State returns the value of property "BaseAudioContext.state".
//
// The returned bool will be false if there is no such property.
func (this BaseAudioContext) State() (AudioContextState, bool) {
	var _ok bool
	_ret := bindings.GetBaseAudioContextState(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioContextState(_ret), _ok
}

// AudioWorklet returns the value of property "BaseAudioContext.audioWorklet".
//
// The returned bool will be false if there is no such property.
func (this BaseAudioContext) AudioWorklet() (AudioWorklet, bool) {
	var _ok bool
	_ret := bindings.GetBaseAudioContextAudioWorklet(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioWorklet{}.FromRef(_ret), _ok
}

// CreateAnalyser calls the method "BaseAudioContext.createAnalyser".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateAnalyser() (AnalyserNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateAnalyser(
		this.Ref(), js.Pointer(&_ok),
	)

	return AnalyserNode{}.FromRef(_ret), _ok
}

// CreateAnalyserFunc returns the method "BaseAudioContext.createAnalyser".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateAnalyserFunc() (fn js.Func[func() AnalyserNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateAnalyserFunc(
			this.Ref(),
		),
	)
}

// CreateBiquadFilter calls the method "BaseAudioContext.createBiquadFilter".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateBiquadFilter() (BiquadFilterNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateBiquadFilter(
		this.Ref(), js.Pointer(&_ok),
	)

	return BiquadFilterNode{}.FromRef(_ret), _ok
}

// CreateBiquadFilterFunc returns the method "BaseAudioContext.createBiquadFilter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateBiquadFilterFunc() (fn js.Func[func() BiquadFilterNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateBiquadFilterFunc(
			this.Ref(),
		),
	)
}

// CreateBuffer calls the method "BaseAudioContext.createBuffer".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateBuffer(numberOfChannels uint32, length uint32, sampleRate float32) (AudioBuffer, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateBuffer(
		this.Ref(), js.Pointer(&_ok),
		uint32(numberOfChannels),
		uint32(length),
		float32(sampleRate),
	)

	return AudioBuffer{}.FromRef(_ret), _ok
}

// CreateBufferFunc returns the method "BaseAudioContext.createBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateBufferFunc() (fn js.Func[func(numberOfChannels uint32, length uint32, sampleRate float32) AudioBuffer]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateBufferFunc(
			this.Ref(),
		),
	)
}

// CreateBufferSource calls the method "BaseAudioContext.createBufferSource".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateBufferSource() (AudioBufferSourceNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateBufferSource(
		this.Ref(), js.Pointer(&_ok),
	)

	return AudioBufferSourceNode{}.FromRef(_ret), _ok
}

// CreateBufferSourceFunc returns the method "BaseAudioContext.createBufferSource".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateBufferSourceFunc() (fn js.Func[func() AudioBufferSourceNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateBufferSourceFunc(
			this.Ref(),
		),
	)
}

// CreateChannelMerger calls the method "BaseAudioContext.createChannelMerger".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateChannelMerger(numberOfInputs uint32) (ChannelMergerNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateChannelMerger(
		this.Ref(), js.Pointer(&_ok),
		uint32(numberOfInputs),
	)

	return ChannelMergerNode{}.FromRef(_ret), _ok
}

// CreateChannelMergerFunc returns the method "BaseAudioContext.createChannelMerger".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateChannelMergerFunc() (fn js.Func[func(numberOfInputs uint32) ChannelMergerNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateChannelMergerFunc(
			this.Ref(),
		),
	)
}

// CreateChannelMerger1 calls the method "BaseAudioContext.createChannelMerger".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateChannelMerger1() (ChannelMergerNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateChannelMerger1(
		this.Ref(), js.Pointer(&_ok),
	)

	return ChannelMergerNode{}.FromRef(_ret), _ok
}

// CreateChannelMerger1Func returns the method "BaseAudioContext.createChannelMerger".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateChannelMerger1Func() (fn js.Func[func() ChannelMergerNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateChannelMerger1Func(
			this.Ref(),
		),
	)
}

// CreateChannelSplitter calls the method "BaseAudioContext.createChannelSplitter".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateChannelSplitter(numberOfOutputs uint32) (ChannelSplitterNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateChannelSplitter(
		this.Ref(), js.Pointer(&_ok),
		uint32(numberOfOutputs),
	)

	return ChannelSplitterNode{}.FromRef(_ret), _ok
}

// CreateChannelSplitterFunc returns the method "BaseAudioContext.createChannelSplitter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateChannelSplitterFunc() (fn js.Func[func(numberOfOutputs uint32) ChannelSplitterNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateChannelSplitterFunc(
			this.Ref(),
		),
	)
}

// CreateChannelSplitter1 calls the method "BaseAudioContext.createChannelSplitter".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateChannelSplitter1() (ChannelSplitterNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateChannelSplitter1(
		this.Ref(), js.Pointer(&_ok),
	)

	return ChannelSplitterNode{}.FromRef(_ret), _ok
}

// CreateChannelSplitter1Func returns the method "BaseAudioContext.createChannelSplitter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateChannelSplitter1Func() (fn js.Func[func() ChannelSplitterNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateChannelSplitter1Func(
			this.Ref(),
		),
	)
}

// CreateConstantSource calls the method "BaseAudioContext.createConstantSource".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateConstantSource() (ConstantSourceNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateConstantSource(
		this.Ref(), js.Pointer(&_ok),
	)

	return ConstantSourceNode{}.FromRef(_ret), _ok
}

// CreateConstantSourceFunc returns the method "BaseAudioContext.createConstantSource".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateConstantSourceFunc() (fn js.Func[func() ConstantSourceNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateConstantSourceFunc(
			this.Ref(),
		),
	)
}

// CreateConvolver calls the method "BaseAudioContext.createConvolver".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateConvolver() (ConvolverNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateConvolver(
		this.Ref(), js.Pointer(&_ok),
	)

	return ConvolverNode{}.FromRef(_ret), _ok
}

// CreateConvolverFunc returns the method "BaseAudioContext.createConvolver".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateConvolverFunc() (fn js.Func[func() ConvolverNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateConvolverFunc(
			this.Ref(),
		),
	)
}

// CreateDelay calls the method "BaseAudioContext.createDelay".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateDelay(maxDelayTime float64) (DelayNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateDelay(
		this.Ref(), js.Pointer(&_ok),
		float64(maxDelayTime),
	)

	return DelayNode{}.FromRef(_ret), _ok
}

// CreateDelayFunc returns the method "BaseAudioContext.createDelay".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateDelayFunc() (fn js.Func[func(maxDelayTime float64) DelayNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateDelayFunc(
			this.Ref(),
		),
	)
}

// CreateDelay1 calls the method "BaseAudioContext.createDelay".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateDelay1() (DelayNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateDelay1(
		this.Ref(), js.Pointer(&_ok),
	)

	return DelayNode{}.FromRef(_ret), _ok
}

// CreateDelay1Func returns the method "BaseAudioContext.createDelay".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateDelay1Func() (fn js.Func[func() DelayNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateDelay1Func(
			this.Ref(),
		),
	)
}

// CreateDynamicsCompressor calls the method "BaseAudioContext.createDynamicsCompressor".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateDynamicsCompressor() (DynamicsCompressorNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateDynamicsCompressor(
		this.Ref(), js.Pointer(&_ok),
	)

	return DynamicsCompressorNode{}.FromRef(_ret), _ok
}

// CreateDynamicsCompressorFunc returns the method "BaseAudioContext.createDynamicsCompressor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateDynamicsCompressorFunc() (fn js.Func[func() DynamicsCompressorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateDynamicsCompressorFunc(
			this.Ref(),
		),
	)
}

// CreateGain calls the method "BaseAudioContext.createGain".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateGain() (GainNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateGain(
		this.Ref(), js.Pointer(&_ok),
	)

	return GainNode{}.FromRef(_ret), _ok
}

// CreateGainFunc returns the method "BaseAudioContext.createGain".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateGainFunc() (fn js.Func[func() GainNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateGainFunc(
			this.Ref(),
		),
	)
}

// CreateIIRFilter calls the method "BaseAudioContext.createIIRFilter".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateIIRFilter(feedforward js.Array[float64], feedback js.Array[float64]) (IIRFilterNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateIIRFilter(
		this.Ref(), js.Pointer(&_ok),
		feedforward.Ref(),
		feedback.Ref(),
	)

	return IIRFilterNode{}.FromRef(_ret), _ok
}

// CreateIIRFilterFunc returns the method "BaseAudioContext.createIIRFilter".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateIIRFilterFunc() (fn js.Func[func(feedforward js.Array[float64], feedback js.Array[float64]) IIRFilterNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateIIRFilterFunc(
			this.Ref(),
		),
	)
}

// CreateOscillator calls the method "BaseAudioContext.createOscillator".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateOscillator() (OscillatorNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateOscillator(
		this.Ref(), js.Pointer(&_ok),
	)

	return OscillatorNode{}.FromRef(_ret), _ok
}

// CreateOscillatorFunc returns the method "BaseAudioContext.createOscillator".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateOscillatorFunc() (fn js.Func[func() OscillatorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateOscillatorFunc(
			this.Ref(),
		),
	)
}

// CreatePanner calls the method "BaseAudioContext.createPanner".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreatePanner() (PannerNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreatePanner(
		this.Ref(), js.Pointer(&_ok),
	)

	return PannerNode{}.FromRef(_ret), _ok
}

// CreatePannerFunc returns the method "BaseAudioContext.createPanner".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreatePannerFunc() (fn js.Func[func() PannerNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreatePannerFunc(
			this.Ref(),
		),
	)
}

// CreatePeriodicWave calls the method "BaseAudioContext.createPeriodicWave".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreatePeriodicWave(real js.Array[float32], imag js.Array[float32], constraints PeriodicWaveConstraints) (PeriodicWave, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreatePeriodicWave(
		this.Ref(), js.Pointer(&_ok),
		real.Ref(),
		imag.Ref(),
		js.Pointer(&constraints),
	)

	return PeriodicWave{}.FromRef(_ret), _ok
}

// CreatePeriodicWaveFunc returns the method "BaseAudioContext.createPeriodicWave".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreatePeriodicWaveFunc() (fn js.Func[func(real js.Array[float32], imag js.Array[float32], constraints PeriodicWaveConstraints) PeriodicWave]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreatePeriodicWaveFunc(
			this.Ref(),
		),
	)
}

// CreatePeriodicWave1 calls the method "BaseAudioContext.createPeriodicWave".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreatePeriodicWave1(real js.Array[float32], imag js.Array[float32]) (PeriodicWave, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreatePeriodicWave1(
		this.Ref(), js.Pointer(&_ok),
		real.Ref(),
		imag.Ref(),
	)

	return PeriodicWave{}.FromRef(_ret), _ok
}

// CreatePeriodicWave1Func returns the method "BaseAudioContext.createPeriodicWave".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreatePeriodicWave1Func() (fn js.Func[func(real js.Array[float32], imag js.Array[float32]) PeriodicWave]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreatePeriodicWave1Func(
			this.Ref(),
		),
	)
}

// CreateScriptProcessor calls the method "BaseAudioContext.createScriptProcessor".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateScriptProcessor(bufferSize uint32, numberOfInputChannels uint32, numberOfOutputChannels uint32) (ScriptProcessorNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateScriptProcessor(
		this.Ref(), js.Pointer(&_ok),
		uint32(bufferSize),
		uint32(numberOfInputChannels),
		uint32(numberOfOutputChannels),
	)

	return ScriptProcessorNode{}.FromRef(_ret), _ok
}

// CreateScriptProcessorFunc returns the method "BaseAudioContext.createScriptProcessor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateScriptProcessorFunc() (fn js.Func[func(bufferSize uint32, numberOfInputChannels uint32, numberOfOutputChannels uint32) ScriptProcessorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateScriptProcessorFunc(
			this.Ref(),
		),
	)
}

// CreateScriptProcessor1 calls the method "BaseAudioContext.createScriptProcessor".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateScriptProcessor1(bufferSize uint32, numberOfInputChannels uint32) (ScriptProcessorNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateScriptProcessor1(
		this.Ref(), js.Pointer(&_ok),
		uint32(bufferSize),
		uint32(numberOfInputChannels),
	)

	return ScriptProcessorNode{}.FromRef(_ret), _ok
}

// CreateScriptProcessor1Func returns the method "BaseAudioContext.createScriptProcessor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateScriptProcessor1Func() (fn js.Func[func(bufferSize uint32, numberOfInputChannels uint32) ScriptProcessorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateScriptProcessor1Func(
			this.Ref(),
		),
	)
}

// CreateScriptProcessor2 calls the method "BaseAudioContext.createScriptProcessor".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateScriptProcessor2(bufferSize uint32) (ScriptProcessorNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateScriptProcessor2(
		this.Ref(), js.Pointer(&_ok),
		uint32(bufferSize),
	)

	return ScriptProcessorNode{}.FromRef(_ret), _ok
}

// CreateScriptProcessor2Func returns the method "BaseAudioContext.createScriptProcessor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateScriptProcessor2Func() (fn js.Func[func(bufferSize uint32) ScriptProcessorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateScriptProcessor2Func(
			this.Ref(),
		),
	)
}

// CreateScriptProcessor3 calls the method "BaseAudioContext.createScriptProcessor".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateScriptProcessor3() (ScriptProcessorNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateScriptProcessor3(
		this.Ref(), js.Pointer(&_ok),
	)

	return ScriptProcessorNode{}.FromRef(_ret), _ok
}

// CreateScriptProcessor3Func returns the method "BaseAudioContext.createScriptProcessor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateScriptProcessor3Func() (fn js.Func[func() ScriptProcessorNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateScriptProcessor3Func(
			this.Ref(),
		),
	)
}

// CreateStereoPanner calls the method "BaseAudioContext.createStereoPanner".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateStereoPanner() (StereoPannerNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateStereoPanner(
		this.Ref(), js.Pointer(&_ok),
	)

	return StereoPannerNode{}.FromRef(_ret), _ok
}

// CreateStereoPannerFunc returns the method "BaseAudioContext.createStereoPanner".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateStereoPannerFunc() (fn js.Func[func() StereoPannerNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateStereoPannerFunc(
			this.Ref(),
		),
	)
}

// CreateWaveShaper calls the method "BaseAudioContext.createWaveShaper".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) CreateWaveShaper() (WaveShaperNode, bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextCreateWaveShaper(
		this.Ref(), js.Pointer(&_ok),
	)

	return WaveShaperNode{}.FromRef(_ret), _ok
}

// CreateWaveShaperFunc returns the method "BaseAudioContext.createWaveShaper".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) CreateWaveShaperFunc() (fn js.Func[func() WaveShaperNode]) {
	return fn.FromRef(
		bindings.BaseAudioContextCreateWaveShaperFunc(
			this.Ref(),
		),
	)
}

// DecodeAudioData calls the method "BaseAudioContext.decodeAudioData".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) DecodeAudioData(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)], errorCallback js.Func[func(err DOMException)]) (js.Promise[AudioBuffer], bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextDecodeAudioData(
		this.Ref(), js.Pointer(&_ok),
		audioData.Ref(),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return js.Promise[AudioBuffer]{}.FromRef(_ret), _ok
}

// DecodeAudioDataFunc returns the method "BaseAudioContext.decodeAudioData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) DecodeAudioDataFunc() (fn js.Func[func(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)], errorCallback js.Func[func(err DOMException)]) js.Promise[AudioBuffer]]) {
	return fn.FromRef(
		bindings.BaseAudioContextDecodeAudioDataFunc(
			this.Ref(),
		),
	)
}

// DecodeAudioData1 calls the method "BaseAudioContext.decodeAudioData".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) DecodeAudioData1(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)]) (js.Promise[AudioBuffer], bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextDecodeAudioData1(
		this.Ref(), js.Pointer(&_ok),
		audioData.Ref(),
		successCallback.Ref(),
	)

	return js.Promise[AudioBuffer]{}.FromRef(_ret), _ok
}

// DecodeAudioData1Func returns the method "BaseAudioContext.decodeAudioData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) DecodeAudioData1Func() (fn js.Func[func(audioData js.ArrayBuffer, successCallback js.Func[func(decodedData AudioBuffer)]) js.Promise[AudioBuffer]]) {
	return fn.FromRef(
		bindings.BaseAudioContextDecodeAudioData1Func(
			this.Ref(),
		),
	)
}

// DecodeAudioData2 calls the method "BaseAudioContext.decodeAudioData".
//
// The returned bool will be false if there is no such method.
func (this BaseAudioContext) DecodeAudioData2(audioData js.ArrayBuffer) (js.Promise[AudioBuffer], bool) {
	var _ok bool
	_ret := bindings.CallBaseAudioContextDecodeAudioData2(
		this.Ref(), js.Pointer(&_ok),
		audioData.Ref(),
	)

	return js.Promise[AudioBuffer]{}.FromRef(_ret), _ok
}

// DecodeAudioData2Func returns the method "BaseAudioContext.decodeAudioData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this BaseAudioContext) DecodeAudioData2Func() (fn js.Func[func(audioData js.ArrayBuffer) js.Promise[AudioBuffer]]) {
	return fn.FromRef(
		bindings.BaseAudioContextDecodeAudioData2Func(
			this.Ref(),
		),
	)
}

type AnalyserOptions struct {
	// FftSize is "AnalyserOptions.fftSize"
	//
	// Optional, defaults to 2048.
	FftSize uint32
	// MaxDecibels is "AnalyserOptions.maxDecibels"
	//
	// Optional, defaults to -30.
	MaxDecibels float64
	// MinDecibels is "AnalyserOptions.minDecibels"
	//
	// Optional, defaults to -100.
	MinDecibels float64
	// SmoothingTimeConstant is "AnalyserOptions.smoothingTimeConstant"
	//
	// Optional, defaults to 0.8.
	SmoothingTimeConstant float64
	// ChannelCount is "AnalyserOptions.channelCount"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 AnalyserOptions AnalyserOptions [// AnalyserOptions] [0x14000182aa0 0x14000182be0 0x14000182d20 0x14000182e60 0x14000182fa0 0x140001830e0 0x14000183180 0x14000182b40 0x14000182c80 0x14000182dc0 0x14000182f00 0x14000183040] 0x14000aa3590 {0 0}} in the application heap.
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

func NewAnalyserNode(context BaseAudioContext, options AnalyserOptions) AnalyserNode {
	return AnalyserNode{}.FromRef(
		bindings.NewAnalyserNodeByAnalyserNode(
			context.Ref(),
			js.Pointer(&options)),
	)
}

func NewAnalyserNodeByAnalyserNode1(context BaseAudioContext) AnalyserNode {
	return AnalyserNode{}.FromRef(
		bindings.NewAnalyserNodeByAnalyserNode1(
			context.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this AnalyserNode) FftSize() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAnalyserNodeFftSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// FftSize sets the value of property "AnalyserNode.fftSize" to val.
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
// The returned bool will be false if there is no such property.
func (this AnalyserNode) FrequencyBinCount() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetAnalyserNodeFrequencyBinCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MinDecibels returns the value of property "AnalyserNode.minDecibels".
//
// The returned bool will be false if there is no such property.
func (this AnalyserNode) MinDecibels() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAnalyserNodeMinDecibels(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// MinDecibels sets the value of property "AnalyserNode.minDecibels" to val.
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
// The returned bool will be false if there is no such property.
func (this AnalyserNode) MaxDecibels() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAnalyserNodeMaxDecibels(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// MaxDecibels sets the value of property "AnalyserNode.maxDecibels" to val.
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
// The returned bool will be false if there is no such property.
func (this AnalyserNode) SmoothingTimeConstant() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAnalyserNodeSmoothingTimeConstant(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// SmoothingTimeConstant sets the value of property "AnalyserNode.smoothingTimeConstant" to val.
//
// It returns false if the property cannot be set.
func (this AnalyserNode) SetSmoothingTimeConstant(val float64) bool {
	return js.True == bindings.SetAnalyserNodeSmoothingTimeConstant(
		this.Ref(),
		float64(val),
	)
}

// GetFloatFrequencyData calls the method "AnalyserNode.getFloatFrequencyData".
//
// The returned bool will be false if there is no such method.
func (this AnalyserNode) GetFloatFrequencyData(array js.TypedArray[float32]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnalyserNodeGetFloatFrequencyData(
		this.Ref(), js.Pointer(&_ok),
		array.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetFloatFrequencyDataFunc returns the method "AnalyserNode.getFloatFrequencyData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnalyserNode) GetFloatFrequencyDataFunc() (fn js.Func[func(array js.TypedArray[float32])]) {
	return fn.FromRef(
		bindings.AnalyserNodeGetFloatFrequencyDataFunc(
			this.Ref(),
		),
	)
}

// GetByteFrequencyData calls the method "AnalyserNode.getByteFrequencyData".
//
// The returned bool will be false if there is no such method.
func (this AnalyserNode) GetByteFrequencyData(array js.TypedArray[uint8]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnalyserNodeGetByteFrequencyData(
		this.Ref(), js.Pointer(&_ok),
		array.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetByteFrequencyDataFunc returns the method "AnalyserNode.getByteFrequencyData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnalyserNode) GetByteFrequencyDataFunc() (fn js.Func[func(array js.TypedArray[uint8])]) {
	return fn.FromRef(
		bindings.AnalyserNodeGetByteFrequencyDataFunc(
			this.Ref(),
		),
	)
}

// GetFloatTimeDomainData calls the method "AnalyserNode.getFloatTimeDomainData".
//
// The returned bool will be false if there is no such method.
func (this AnalyserNode) GetFloatTimeDomainData(array js.TypedArray[float32]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnalyserNodeGetFloatTimeDomainData(
		this.Ref(), js.Pointer(&_ok),
		array.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetFloatTimeDomainDataFunc returns the method "AnalyserNode.getFloatTimeDomainData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnalyserNode) GetFloatTimeDomainDataFunc() (fn js.Func[func(array js.TypedArray[float32])]) {
	return fn.FromRef(
		bindings.AnalyserNodeGetFloatTimeDomainDataFunc(
			this.Ref(),
		),
	)
}

// GetByteTimeDomainData calls the method "AnalyserNode.getByteTimeDomainData".
//
// The returned bool will be false if there is no such method.
func (this AnalyserNode) GetByteTimeDomainData(array js.TypedArray[uint8]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnalyserNodeGetByteTimeDomainData(
		this.Ref(), js.Pointer(&_ok),
		array.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// GetByteTimeDomainDataFunc returns the method "AnalyserNode.getByteTimeDomainData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnalyserNode) GetByteTimeDomainDataFunc() (fn js.Func[func(array js.TypedArray[uint8])]) {
	return fn.FromRef(
		bindings.AnalyserNodeGetByteTimeDomainDataFunc(
			this.Ref(),
		),
	)
}

type AnimationEventInit struct {
	// AnimationName is "AnimationEventInit.animationName"
	//
	// Optional, defaults to "".
	AnimationName js.String
	// ElapsedTime is "AnimationEventInit.elapsedTime"
	//
	// Optional, defaults to 0.0.
	ElapsedTime float64
	// PseudoElement is "AnimationEventInit.pseudoElement"
	//
	// Optional, defaults to "".
	PseudoElement js.String
	// Bubbles is "AnimationEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "AnimationEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "AnimationEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 AnimationEventInit AnimationEventInit [// AnimationEventInit] [0x140001832c0 0x14000183360 0x140001834a0 0x14000183540 0x14000183680 0x140001837c0 0x14000183400 0x140001835e0 0x14000183720 0x14000183860] 0x14000aa3620 {0 0}} in the application heap.
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

func NewAnimationEvent(typ js.String, animationEventInitDict AnimationEventInit) AnimationEvent {
	return AnimationEvent{}.FromRef(
		bindings.NewAnimationEventByAnimationEvent(
			typ.Ref(),
			js.Pointer(&animationEventInitDict)),
	)
}

func NewAnimationEventByAnimationEvent1(typ js.String) AnimationEvent {
	return AnimationEvent{}.FromRef(
		bindings.NewAnimationEventByAnimationEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this AnimationEvent) AnimationName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAnimationEventAnimationName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ElapsedTime returns the value of property "AnimationEvent.elapsedTime".
//
// The returned bool will be false if there is no such property.
func (this AnimationEvent) ElapsedTime() (float64, bool) {
	var _ok bool
	_ret := bindings.GetAnimationEventElapsedTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PseudoElement returns the value of property "AnimationEvent.pseudoElement".
//
// The returned bool will be false if there is no such property.
func (this AnimationEvent) PseudoElement() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAnimationEventPseudoElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
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
	Bubbles bool
	// Cancelable is "AnimationPlaybackEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "AnimationPlaybackEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 AnimationPlaybackEventInit AnimationPlaybackEventInit [// AnimationPlaybackEventInit] [0x140001839a0 0x14000183a40 0x14000183ae0 0x14000183c20 0x14000183d60 0x14000183b80 0x14000183cc0 0x14000183e00] 0x14000aa3680 {0 0}} in the application heap.
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

func NewAnimationPlaybackEvent(typ js.String, eventInitDict AnimationPlaybackEventInit) AnimationPlaybackEvent {
	return AnimationPlaybackEvent{}.FromRef(
		bindings.NewAnimationPlaybackEventByAnimationPlaybackEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewAnimationPlaybackEventByAnimationPlaybackEvent1(typ js.String) AnimationPlaybackEvent {
	return AnimationPlaybackEvent{}.FromRef(
		bindings.NewAnimationPlaybackEventByAnimationPlaybackEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this AnimationPlaybackEvent) CurrentTime() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetAnimationPlaybackEventCurrentTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
}

// TimelineTime returns the value of property "AnimationPlaybackEvent.timelineTime".
//
// The returned bool will be false if there is no such property.
func (this AnimationPlaybackEvent) TimelineTime() (CSSNumberish, bool) {
	var _ok bool
	_ret := bindings.GetAnimationPlaybackEventTimelineTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumberish{}.FromRef(_ret), _ok
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

// RegisterAnimator calls the method "AnimationWorkletGlobalScope.registerAnimator".
//
// The returned bool will be false if there is no such method.
func (this AnimationWorkletGlobalScope) RegisterAnimator(name js.String, animatorCtor js.Func[func(options js.Any, state js.Any) js.Any]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAnimationWorkletGlobalScopeRegisterAnimator(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		animatorCtor.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RegisterAnimatorFunc returns the method "AnimationWorkletGlobalScope.registerAnimator".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AnimationWorkletGlobalScope) RegisterAnimatorFunc() (fn js.Func[func(name js.String, animatorCtor js.Func[func(options js.Any, state js.Any) js.Any])]) {
	return fn.FromRef(
		bindings.AnimationWorkletGlobalScopeRegisterAnimatorFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 AttributionReportingRequestOptions AttributionReportingRequestOptions [// AttributionReportingRequestOptions] [0x14000183f40 0x1400019e000] 0x14000aa36c8 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 AuctionAd AuctionAd [// AuctionAd] [0x1400019e0a0 0x1400019e140 0x1400019e1e0 0x1400019e280] 0x14000aa36e0 {0 0}} in the application heap.
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
	SellerTimeout uint64
	// SellerExperimentGroupId is "AuctionAdConfig.sellerExperimentGroupId"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 AuctionAdConfig AuctionAdConfig [// AuctionAdConfig] [0x1400019e3c0 0x1400019e460 0x1400019e500 0x1400019e5a0 0x1400019e640 0x1400019e6e0 0x1400019e780 0x1400019e820 0x1400019e960 0x1400019eaa0 0x1400019eb40 0x1400019ebe0 0x1400019ec80 0x1400019ed20 0x1400019edc0 0x1400019ee60 0x1400019ef00 0x1400019efa0 0x1400019f040 0x1400019e8c0 0x1400019ea00] 0x14000aa3710 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 AuctionAdInterestGroup AuctionAdInterestGroup [// AuctionAdInterestGroup] [0x1400019f0e0 0x1400019f220 0x1400019f2c0 0x1400019f360 0x1400019f400 0x1400019f4a0 0x1400019f5e0 0x1400019f680 0x1400019f720 0x1400019f7c0 0x1400019f860 0x1400019f900 0x1400019f9a0 0x1400019fa40 0x1400019fae0 0x1400019fb80 0x1400019f180 0x1400019f540] 0x14000aa3740 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 AuctionAdInterestGroupKey AuctionAdInterestGroupKey [// AuctionAdInterestGroupKey] [0x1400019fc20 0x1400019fcc0] 0x14000aa3788 {0 0}} in the application heap.
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
	Bitrate uint64
	// Samplerate is "AudioConfiguration.samplerate"
	//
	// Optional
	Samplerate uint32
	// SpatialRendering is "AudioConfiguration.spatialRendering"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 AudioConfiguration AudioConfiguration [// AudioConfiguration] [0x1400019fd60 0x1400019fe00 0x1400019fea0 0x1400020c000 0x1400020c140 0x1400019ff40 0x1400020c0a0 0x1400020c1e0] 0x14000aa37d0 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 AudioSinkOptions AudioSinkOptions [// AudioSinkOptions] [0x1400020c460] 0x14000aa3890 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 AudioContextOptions AudioContextOptions [// AudioContextOptions] [0x1400020c280 0x1400020c320 0x1400020c500 0x1400020c3c0] 0x14000aa3848 {0 0}} in the application heap.
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
	ContextTime float64
	// PerformanceTime is "AudioTimestamp.performanceTime"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 AudioTimestamp AudioTimestamp [// AudioTimestamp] [0x1400020c5a0 0x1400020c6e0 0x1400020c640 0x1400020c780] 0x14000aa38d8 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this TextTrackCue) Track() (TextTrack, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackCueTrack(
		this.Ref(), js.Pointer(&_ok),
	)
	return TextTrack{}.FromRef(_ret), _ok
}

// Id returns the value of property "TextTrackCue.id".
//
// The returned bool will be false if there is no such property.
func (this TextTrackCue) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackCueId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Id sets the value of property "TextTrackCue.id" to val.
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
// The returned bool will be false if there is no such property.
func (this TextTrackCue) StartTime() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackCueStartTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// StartTime sets the value of property "TextTrackCue.startTime" to val.
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
// The returned bool will be false if there is no such property.
func (this TextTrackCue) EndTime() (float64, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackCueEndTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// EndTime sets the value of property "TextTrackCue.endTime" to val.
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
// The returned bool will be false if there is no such property.
func (this TextTrackCue) PauseOnExit() (bool, bool) {
	var _ok bool
	_ret := bindings.GetTextTrackCuePauseOnExit(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// PauseOnExit sets the value of property "TextTrackCue.pauseOnExit" to val.
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

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

const (
	_ RecordingState = iota

	RecordingState_INACTIVE
	RecordingState_RECORDING
	RecordingState_PAUSED
)

func (RecordingState) FromRef(str js.Ref) RecordingState {
	return RecordingState(bindings.ConstOfRecordingState(str))
}

func (x RecordingState) String() (string, bool) {
	switch x {
	case RecordingState_INACTIVE:
		return "inactive", true
	case RecordingState_RECORDING:
		return "recording", true
	case RecordingState_PAUSED:
		return "paused", true
	default:
		return "", false
	}
}

func NewMediaRecorder(stream MediaStream, options MediaRecorderOptions) MediaRecorder {
	return MediaRecorder{}.FromRef(
		bindings.NewMediaRecorderByMediaRecorder(
			stream.Ref(),
			js.Pointer(&options)),
	)
}

func NewMediaRecorderByMediaRecorder1(stream MediaStream) MediaRecorder {
	return MediaRecorder{}.FromRef(
		bindings.NewMediaRecorderByMediaRecorder1(
			stream.Ref()),
	)
}

type MediaRecorder struct {
	EventTarget
}

func (this MediaRecorder) Once() MediaRecorder {
	this.Ref().Once()
	return this
}

func (this MediaRecorder) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this MediaRecorder) FromRef(ref js.Ref) MediaRecorder {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this MediaRecorder) Free() {
	this.Ref().Free()
}

// Stream returns the value of property "MediaRecorder.stream".
//
// The returned bool will be false if there is no such property.
func (this MediaRecorder) Stream() (MediaStream, bool) {
	var _ok bool
	_ret := bindings.GetMediaRecorderStream(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaStream{}.FromRef(_ret), _ok
}

// MimeType returns the value of property "MediaRecorder.mimeType".
//
// The returned bool will be false if there is no such property.
func (this MediaRecorder) MimeType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaRecorderMimeType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// State returns the value of property "MediaRecorder.state".
//
// The returned bool will be false if there is no such property.
func (this MediaRecorder) State() (RecordingState, bool) {
	var _ok bool
	_ret := bindings.GetMediaRecorderState(
		this.Ref(), js.Pointer(&_ok),
	)
	return RecordingState(_ret), _ok
}

// VideoBitsPerSecond returns the value of property "MediaRecorder.videoBitsPerSecond".
//
// The returned bool will be false if there is no such property.
func (this MediaRecorder) VideoBitsPerSecond() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetMediaRecorderVideoBitsPerSecond(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// AudioBitsPerSecond returns the value of property "MediaRecorder.audioBitsPerSecond".
//
// The returned bool will be false if there is no such property.
func (this MediaRecorder) AudioBitsPerSecond() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetMediaRecorderAudioBitsPerSecond(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// AudioBitrateMode returns the value of property "MediaRecorder.audioBitrateMode".
//
// The returned bool will be false if there is no such property.
func (this MediaRecorder) AudioBitrateMode() (BitrateMode, bool) {
	var _ok bool
	_ret := bindings.GetMediaRecorderAudioBitrateMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return BitrateMode(_ret), _ok
}

// Start calls the method "MediaRecorder.start".
//
// The returned bool will be false if there is no such method.
func (this MediaRecorder) Start(timeslice uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaRecorderStart(
		this.Ref(), js.Pointer(&_ok),
		uint32(timeslice),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StartFunc returns the method "MediaRecorder.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaRecorder) StartFunc() (fn js.Func[func(timeslice uint32)]) {
	return fn.FromRef(
		bindings.MediaRecorderStartFunc(
			this.Ref(),
		),
	)
}

// Start1 calls the method "MediaRecorder.start".
//
// The returned bool will be false if there is no such method.
func (this MediaRecorder) Start1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaRecorderStart1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Start1Func returns the method "MediaRecorder.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaRecorder) Start1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaRecorderStart1Func(
			this.Ref(),
		),
	)
}

// Stop calls the method "MediaRecorder.stop".
//
// The returned bool will be false if there is no such method.
func (this MediaRecorder) Stop() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaRecorderStop(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StopFunc returns the method "MediaRecorder.stop".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaRecorder) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaRecorderStopFunc(
			this.Ref(),
		),
	)
}

// Pause calls the method "MediaRecorder.pause".
//
// The returned bool will be false if there is no such method.
func (this MediaRecorder) Pause() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaRecorderPause(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PauseFunc returns the method "MediaRecorder.pause".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaRecorder) PauseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaRecorderPauseFunc(
			this.Ref(),
		),
	)
}

// Resume calls the method "MediaRecorder.resume".
//
// The returned bool will be false if there is no such method.
func (this MediaRecorder) Resume() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaRecorderResume(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResumeFunc returns the method "MediaRecorder.resume".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaRecorder) ResumeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaRecorderResumeFunc(
			this.Ref(),
		),
	)
}

// RequestData calls the method "MediaRecorder.requestData".
//
// The returned bool will be false if there is no such method.
func (this MediaRecorder) RequestData() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMediaRecorderRequestData(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RequestDataFunc returns the method "MediaRecorder.requestData".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaRecorder) RequestDataFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaRecorderRequestDataFunc(
			this.Ref(),
		),
	)
}

// IsTypeSupported calls the staticmethod "MediaRecorder.isTypeSupported".
//
// The returned bool will be false if there is no such method.
func (this MediaRecorder) IsTypeSupported(typ js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallMediaRecorderIsTypeSupported(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return _ret == js.True, _ok
}

// IsTypeSupportedFunc returns the staticmethod "MediaRecorder.isTypeSupported".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaRecorder) IsTypeSupportedFunc() (fn js.Func[func(typ js.String) bool]) {
	return fn.FromRef(
		bindings.MediaRecorderIsTypeSupportedFunc(
			this.Ref(),
		),
	)
}

type MediaStreamTrackEventInit struct {
	// Track is "MediaStreamTrackEventInit.track"
	//
	// Required
	Track MediaStreamTrack
	// Bubbles is "MediaStreamTrackEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "MediaStreamTrackEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "MediaStreamTrackEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaStreamTrackEventInit with all fields set.
func (p MediaStreamTrackEventInit) FromRef(ref js.Ref) MediaStreamTrackEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MediaStreamTrackEventInit MediaStreamTrackEventInit [// MediaStreamTrackEventInit] [0x140008eda40 0x140008edae0 0x140008edc20 0x140008edd60 0x140008edb80 0x140008edcc0 0x140008ede00] 0x140008bf3e0 {0 0}} in the application heap.
func (p MediaStreamTrackEventInit) New() js.Ref {
	return bindings.MediaStreamTrackEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaStreamTrackEventInit) UpdateFrom(ref js.Ref) {
	bindings.MediaStreamTrackEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaStreamTrackEventInit) Update(ref js.Ref) {
	bindings.MediaStreamTrackEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMediaStreamTrackEvent(typ js.String, eventInitDict MediaStreamTrackEventInit) MediaStreamTrackEvent {
	return MediaStreamTrackEvent{}.FromRef(
		bindings.NewMediaStreamTrackEventByMediaStreamTrackEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type MediaStreamTrackEvent struct {
	Event
}

func (this MediaStreamTrackEvent) Once() MediaStreamTrackEvent {
	this.Ref().Once()
	return this
}

func (this MediaStreamTrackEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this MediaStreamTrackEvent) FromRef(ref js.Ref) MediaStreamTrackEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this MediaStreamTrackEvent) Free() {
	this.Ref().Free()
}

// Track returns the value of property "MediaStreamTrackEvent.track".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamTrackEvent) Track() (MediaStreamTrack, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamTrackEventTrack(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaStreamTrack{}.FromRef(_ret), _ok
}

type MediaStreamTrackProcessorInit struct {
	// Track is "MediaStreamTrackProcessorInit.track"
	//
	// Required
	Track MediaStreamTrack
	// MaxBufferSize is "MediaStreamTrackProcessorInit.maxBufferSize"
	//
	// Optional
	MaxBufferSize uint16

	FFI_USE_MaxBufferSize bool // for MaxBufferSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaStreamTrackProcessorInit with all fields set.
func (p MediaStreamTrackProcessorInit) FromRef(ref js.Ref) MediaStreamTrackProcessorInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MediaStreamTrackProcessorInit MediaStreamTrackProcessorInit [// MediaStreamTrackProcessorInit] [0x140008edea0 0x140008edf40 0x140008fe000] 0x140008bf410 {0 0}} in the application heap.
func (p MediaStreamTrackProcessorInit) New() js.Ref {
	return bindings.MediaStreamTrackProcessorInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaStreamTrackProcessorInit) UpdateFrom(ref js.Ref) {
	bindings.MediaStreamTrackProcessorInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaStreamTrackProcessorInit) Update(ref js.Ref) {
	bindings.MediaStreamTrackProcessorInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMediaStreamTrackProcessor(init MediaStreamTrackProcessorInit) MediaStreamTrackProcessor {
	return MediaStreamTrackProcessor{}.FromRef(
		bindings.NewMediaStreamTrackProcessorByMediaStreamTrackProcessor(
			js.Pointer(&init)),
	)
}

type MediaStreamTrackProcessor struct {
	ref js.Ref
}

func (this MediaStreamTrackProcessor) Once() MediaStreamTrackProcessor {
	this.Ref().Once()
	return this
}

func (this MediaStreamTrackProcessor) Ref() js.Ref {
	return this.ref
}

func (this MediaStreamTrackProcessor) FromRef(ref js.Ref) MediaStreamTrackProcessor {
	this.ref = ref
	return this
}

func (this MediaStreamTrackProcessor) Free() {
	this.Ref().Free()
}

// Readable returns the value of property "MediaStreamTrackProcessor.readable".
//
// The returned bool will be false if there is no such property.
func (this MediaStreamTrackProcessor) Readable() (ReadableStream, bool) {
	var _ok bool
	_ret := bindings.GetMediaStreamTrackProcessorReadable(
		this.Ref(), js.Pointer(&_ok),
	)
	return ReadableStream{}.FromRef(_ret), _ok
}

// Readable sets the value of property "MediaStreamTrackProcessor.readable" to val.
//
// It returns false if the property cannot be set.
func (this MediaStreamTrackProcessor) SetReadable(val ReadableStream) bool {
	return js.True == bindings.SetMediaStreamTrackProcessorReadable(
		this.Ref(),
		val.Ref(),
	)
}

type MemoryDescriptor struct {
	// Initial is "MemoryDescriptor.initial"
	//
	// Required
	Initial uint32
	// Maximum is "MemoryDescriptor.maximum"
	//
	// Optional
	Maximum uint32

	FFI_USE_Maximum bool // for Maximum.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MemoryDescriptor with all fields set.
func (p MemoryDescriptor) FromRef(ref js.Ref) MemoryDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MemoryDescriptor MemoryDescriptor [// MemoryDescriptor] [0x140008fe0a0 0x140008fe140 0x140008fe1e0] 0x140008bf458 {0 0}} in the application heap.
func (p MemoryDescriptor) New() js.Ref {
	return bindings.MemoryDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MemoryDescriptor) UpdateFrom(ref js.Ref) {
	bindings.MemoryDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MemoryDescriptor) Update(ref js.Ref) {
	bindings.MemoryDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMemory(descriptor MemoryDescriptor) Memory {
	return Memory{}.FromRef(
		bindings.NewMemoryByMemory(
			js.Pointer(&descriptor)),
	)
}

type Memory struct {
	ref js.Ref
}

func (this Memory) Once() Memory {
	this.Ref().Once()
	return this
}

func (this Memory) Ref() js.Ref {
	return this.ref
}

func (this Memory) FromRef(ref js.Ref) Memory {
	this.ref = ref
	return this
}

func (this Memory) Free() {
	this.Ref().Free()
}

// Buffer returns the value of property "Memory.buffer".
//
// The returned bool will be false if there is no such property.
func (this Memory) Buffer() (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.GetMemoryBuffer(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

// Grow calls the method "Memory.grow".
//
// The returned bool will be false if there is no such method.
func (this Memory) Grow(delta uint32) (uint32, bool) {
	var _ok bool
	_ret := bindings.CallMemoryGrow(
		this.Ref(), js.Pointer(&_ok),
		uint32(delta),
	)

	return uint32(_ret), _ok
}

// GrowFunc returns the method "Memory.grow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Memory) GrowFunc() (fn js.Func[func(delta uint32) uint32]) {
	return fn.FromRef(
		bindings.MemoryGrowFunc(
			this.Ref(),
		),
	)
}

type MessageChannel struct {
	ref js.Ref
}

func (this MessageChannel) Once() MessageChannel {
	this.Ref().Once()
	return this
}

func (this MessageChannel) Ref() js.Ref {
	return this.ref
}

func (this MessageChannel) FromRef(ref js.Ref) MessageChannel {
	this.ref = ref
	return this
}

func (this MessageChannel) Free() {
	this.Ref().Free()
}

// Port1 returns the value of property "MessageChannel.port1".
//
// The returned bool will be false if there is no such property.
func (this MessageChannel) Port1() (MessagePort, bool) {
	var _ok bool
	_ret := bindings.GetMessageChannelPort1(
		this.Ref(), js.Pointer(&_ok),
	)
	return MessagePort{}.FromRef(_ret), _ok
}

// Port2 returns the value of property "MessageChannel.port2".
//
// The returned bool will be false if there is no such property.
func (this MessageChannel) Port2() (MessagePort, bool) {
	var _ok bool
	_ret := bindings.GetMessageChannelPort2(
		this.Ref(), js.Pointer(&_ok),
	)
	return MessagePort{}.FromRef(_ret), _ok
}

type OneOf_Object_MessagePort_ServiceWorker struct {
	ref js.Ref
}

func (x OneOf_Object_MessagePort_ServiceWorker) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Object_MessagePort_ServiceWorker) Free() {
	x.ref.Free()
}

func (x OneOf_Object_MessagePort_ServiceWorker) FromRef(ref js.Ref) OneOf_Object_MessagePort_ServiceWorker {
	return OneOf_Object_MessagePort_ServiceWorker{
		ref: ref,
	}
}

func (x OneOf_Object_MessagePort_ServiceWorker) Object() js.Object {
	return js.Object{}.FromRef(x.ref)
}

func (x OneOf_Object_MessagePort_ServiceWorker) MessagePort() MessagePort {
	return MessagePort{}.FromRef(x.ref)
}

func (x OneOf_Object_MessagePort_ServiceWorker) ServiceWorker() ServiceWorker {
	return ServiceWorker{}.FromRef(x.ref)
}

type MessageEventSource = OneOf_Object_MessagePort_ServiceWorker

type MessageEventInit struct {
	// Data is "MessageEventInit.data"
	//
	// Optional, defaults to null.
	Data js.Any
	// Origin is "MessageEventInit.origin"
	//
	// Optional, defaults to "".
	Origin js.String
	// LastEventId is "MessageEventInit.lastEventId"
	//
	// Optional, defaults to "".
	LastEventId js.String
	// Source is "MessageEventInit.source"
	//
	// Optional, defaults to null.
	Source MessageEventSource
	// Ports is "MessageEventInit.ports"
	//
	// Optional, defaults to [].
	Ports js.Array[MessagePort]
	// Bubbles is "MessageEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "MessageEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "MessageEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MessageEventInit with all fields set.
func (p MessageEventInit) FromRef(ref js.Ref) MessageEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MessageEventInit MessageEventInit [// MessageEventInit] [0x140008fe320 0x140008fe3c0 0x140008fe460 0x140008fe500 0x140008fe5a0 0x140008fe640 0x140008fe780 0x140008fe8c0 0x140008fe6e0 0x140008fe820 0x140008fe960] 0x140008bf4d0 {0 0}} in the application heap.
func (p MessageEventInit) New() js.Ref {
	return bindings.MessageEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MessageEventInit) UpdateFrom(ref js.Ref) {
	bindings.MessageEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MessageEventInit) Update(ref js.Ref) {
	bindings.MessageEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMessageEvent(typ js.String, eventInitDict MessageEventInit) MessageEvent {
	return MessageEvent{}.FromRef(
		bindings.NewMessageEventByMessageEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewMessageEventByMessageEvent1(typ js.String) MessageEvent {
	return MessageEvent{}.FromRef(
		bindings.NewMessageEventByMessageEvent1(
			typ.Ref()),
	)
}

type MessageEvent struct {
	Event
}

func (this MessageEvent) Once() MessageEvent {
	this.Ref().Once()
	return this
}

func (this MessageEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this MessageEvent) FromRef(ref js.Ref) MessageEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this MessageEvent) Free() {
	this.Ref().Free()
}

// Data returns the value of property "MessageEvent.data".
//
// The returned bool will be false if there is no such property.
func (this MessageEvent) Data() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetMessageEventData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Origin returns the value of property "MessageEvent.origin".
//
// The returned bool will be false if there is no such property.
func (this MessageEvent) Origin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMessageEventOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LastEventId returns the value of property "MessageEvent.lastEventId".
//
// The returned bool will be false if there is no such property.
func (this MessageEvent) LastEventId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMessageEventLastEventId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Source returns the value of property "MessageEvent.source".
//
// The returned bool will be false if there is no such property.
func (this MessageEvent) Source() (MessageEventSource, bool) {
	var _ok bool
	_ret := bindings.GetMessageEventSource(
		this.Ref(), js.Pointer(&_ok),
	)
	return MessageEventSource{}.FromRef(_ret), _ok
}

// Ports returns the value of property "MessageEvent.ports".
//
// The returned bool will be false if there is no such property.
func (this MessageEvent) Ports() (js.FrozenArray[MessagePort], bool) {
	var _ok bool
	_ret := bindings.GetMessageEventPorts(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[MessagePort]{}.FromRef(_ret), _ok
}

// InitMessageEvent calls the method "MessageEvent.initMessageEvent".
//
// The returned bool will be false if there is no such method.
func (this MessageEvent) InitMessageEvent(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource, ports js.Array[MessagePort]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessageEventInitMessageEvent(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
		lastEventId.Ref(),
		source.Ref(),
		ports.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMessageEventFunc returns the method "MessageEvent.initMessageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessageEvent) InitMessageEventFunc() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource, ports js.Array[MessagePort])]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEventFunc(
			this.Ref(),
		),
	)
}

// InitMessageEvent1 calls the method "MessageEvent.initMessageEvent".
//
// The returned bool will be false if there is no such method.
func (this MessageEvent) InitMessageEvent1(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessageEventInitMessageEvent1(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
		lastEventId.Ref(),
		source.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMessageEvent1Func returns the method "MessageEvent.initMessageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessageEvent) InitMessageEvent1Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent1Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent2 calls the method "MessageEvent.initMessageEvent".
//
// The returned bool will be false if there is no such method.
func (this MessageEvent) InitMessageEvent2(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessageEventInitMessageEvent2(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
		lastEventId.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMessageEvent2Func returns the method "MessageEvent.initMessageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessageEvent) InitMessageEvent2Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent2Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent3 calls the method "MessageEvent.initMessageEvent".
//
// The returned bool will be false if there is no such method.
func (this MessageEvent) InitMessageEvent3(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessageEventInitMessageEvent3(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMessageEvent3Func returns the method "MessageEvent.initMessageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessageEvent) InitMessageEvent3Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent3Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent4 calls the method "MessageEvent.initMessageEvent".
//
// The returned bool will be false if there is no such method.
func (this MessageEvent) InitMessageEvent4(typ js.String, bubbles bool, cancelable bool, data js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessageEventInitMessageEvent4(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMessageEvent4Func returns the method "MessageEvent.initMessageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessageEvent) InitMessageEvent4Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent4Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent5 calls the method "MessageEvent.initMessageEvent".
//
// The returned bool will be false if there is no such method.
func (this MessageEvent) InitMessageEvent5(typ js.String, bubbles bool, cancelable bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessageEventInitMessageEvent5(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMessageEvent5Func returns the method "MessageEvent.initMessageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessageEvent) InitMessageEvent5Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent5Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent6 calls the method "MessageEvent.initMessageEvent".
//
// The returned bool will be false if there is no such method.
func (this MessageEvent) InitMessageEvent6(typ js.String, bubbles bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessageEventInitMessageEvent6(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMessageEvent6Func returns the method "MessageEvent.initMessageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessageEvent) InitMessageEvent6Func() (fn js.Func[func(typ js.String, bubbles bool)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent6Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent7 calls the method "MessageEvent.initMessageEvent".
//
// The returned bool will be false if there is no such method.
func (this MessageEvent) InitMessageEvent7(typ js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMessageEventInitMessageEvent7(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMessageEvent7Func returns the method "MessageEvent.initMessageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MessageEvent) InitMessageEvent7Func() (fn js.Func[func(typ js.String)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent7Func(
			this.Ref(),
		),
	)
}

type MeteringMode uint32

const (
	_ MeteringMode = iota

	MeteringMode_NONE
	MeteringMode_MANUAL
	MeteringMode_SINGLE_SHOT
	MeteringMode_CONTINUOUS
)

func (MeteringMode) FromRef(str js.Ref) MeteringMode {
	return MeteringMode(bindings.ConstOfMeteringMode(str))
}

func (x MeteringMode) String() (string, bool) {
	switch x {
	case MeteringMode_NONE:
		return "none", true
	case MeteringMode_MANUAL:
		return "manual", true
	case MeteringMode_SINGLE_SHOT:
		return "single-shot", true
	case MeteringMode_CONTINUOUS:
		return "continuous", true
	default:
		return "", false
	}
}

type MidiPermissionDescriptor struct {
	// Sysex is "MidiPermissionDescriptor.sysex"
	//
	// Optional, defaults to false.
	Sysex bool
	// Name is "MidiPermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE_Sysex bool // for Sysex.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MidiPermissionDescriptor with all fields set.
func (p MidiPermissionDescriptor) FromRef(ref js.Ref) MidiPermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MidiPermissionDescriptor MidiPermissionDescriptor [// MidiPermissionDescriptor] [0x140008feaa0 0x140008febe0 0x140008feb40] 0x140008bf860 {0 0}} in the application heap.
func (p MidiPermissionDescriptor) New() js.Ref {
	return bindings.MidiPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MidiPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.MidiPermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MidiPermissionDescriptor) Update(ref js.Ref) {
	bindings.MidiPermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MockCameraConfiguration struct {
	// DefaultFrameRate is "MockCameraConfiguration.defaultFrameRate"
	//
	// Optional, defaults to 30.
	DefaultFrameRate float64
	// FacingMode is "MockCameraConfiguration.facingMode"
	//
	// Optional, defaults to "user".
	FacingMode js.String
	// Label is "MockCameraConfiguration.label"
	//
	// Optional
	Label js.String
	// DeviceId is "MockCameraConfiguration.deviceId"
	//
	// Optional
	DeviceId js.String
	// GroupId is "MockCameraConfiguration.groupId"
	//
	// Optional
	GroupId js.String

	FFI_USE_DefaultFrameRate bool // for DefaultFrameRate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MockCameraConfiguration with all fields set.
func (p MockCameraConfiguration) FromRef(ref js.Ref) MockCameraConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MockCameraConfiguration MockCameraConfiguration [// MockCameraConfiguration] [0x140008fec80 0x140008fedc0 0x140008fee60 0x140008fef00 0x140008fefa0 0x140008fed20] 0x140008bf890 {0 0}} in the application heap.
func (p MockCameraConfiguration) New() js.Ref {
	return bindings.MockCameraConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MockCameraConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MockCameraConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MockCameraConfiguration) Update(ref js.Ref) {
	bindings.MockCameraConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MockCaptureDeviceConfiguration struct {
	// Label is "MockCaptureDeviceConfiguration.label"
	//
	// Optional
	Label js.String
	// DeviceId is "MockCaptureDeviceConfiguration.deviceId"
	//
	// Optional
	DeviceId js.String
	// GroupId is "MockCaptureDeviceConfiguration.groupId"
	//
	// Optional
	GroupId js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MockCaptureDeviceConfiguration with all fields set.
func (p MockCaptureDeviceConfiguration) FromRef(ref js.Ref) MockCaptureDeviceConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MockCaptureDeviceConfiguration MockCaptureDeviceConfiguration [// MockCaptureDeviceConfiguration] [0x140008ff040 0x140008ff0e0 0x140008ff180] 0x140008bf8c0 {0 0}} in the application heap.
func (p MockCaptureDeviceConfiguration) New() js.Ref {
	return bindings.MockCaptureDeviceConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MockCaptureDeviceConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MockCaptureDeviceConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MockCaptureDeviceConfiguration) Update(ref js.Ref) {
	bindings.MockCaptureDeviceConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MockCapturePromptResult uint32

const (
	_ MockCapturePromptResult = iota

	MockCapturePromptResult_GRANTED
	MockCapturePromptResult_DENIED
)

func (MockCapturePromptResult) FromRef(str js.Ref) MockCapturePromptResult {
	return MockCapturePromptResult(bindings.ConstOfMockCapturePromptResult(str))
}

func (x MockCapturePromptResult) String() (string, bool) {
	switch x {
	case MockCapturePromptResult_GRANTED:
		return "granted", true
	case MockCapturePromptResult_DENIED:
		return "denied", true
	default:
		return "", false
	}
}

type MockCapturePromptResultConfiguration struct {
	// GetUserMedia is "MockCapturePromptResultConfiguration.getUserMedia"
	//
	// Optional
	GetUserMedia MockCapturePromptResult
	// GetDisplayMedia is "MockCapturePromptResultConfiguration.getDisplayMedia"
	//
	// Optional
	GetDisplayMedia MockCapturePromptResult

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MockCapturePromptResultConfiguration with all fields set.
func (p MockCapturePromptResultConfiguration) FromRef(ref js.Ref) MockCapturePromptResultConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MockCapturePromptResultConfiguration MockCapturePromptResultConfiguration [// MockCapturePromptResultConfiguration] [0x140008ff220 0x140008ff2c0] 0x140008bf8d8 {0 0}} in the application heap.
func (p MockCapturePromptResultConfiguration) New() js.Ref {
	return bindings.MockCapturePromptResultConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MockCapturePromptResultConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MockCapturePromptResultConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MockCapturePromptResultConfiguration) Update(ref js.Ref) {
	bindings.MockCapturePromptResultConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MockMicrophoneConfiguration struct {
	// DefaultSampleRate is "MockMicrophoneConfiguration.defaultSampleRate"
	//
	// Optional, defaults to 44100.
	DefaultSampleRate uint32
	// Label is "MockMicrophoneConfiguration.label"
	//
	// Optional
	Label js.String
	// DeviceId is "MockMicrophoneConfiguration.deviceId"
	//
	// Optional
	DeviceId js.String
	// GroupId is "MockMicrophoneConfiguration.groupId"
	//
	// Optional
	GroupId js.String

	FFI_USE_DefaultSampleRate bool // for DefaultSampleRate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MockMicrophoneConfiguration with all fields set.
func (p MockMicrophoneConfiguration) FromRef(ref js.Ref) MockMicrophoneConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MockMicrophoneConfiguration MockMicrophoneConfiguration [// MockMicrophoneConfiguration] [0x140008ff360 0x140008ff4a0 0x140008ff540 0x140008ff5e0 0x140008ff400] 0x140008bf8f0 {0 0}} in the application heap.
func (p MockMicrophoneConfiguration) New() js.Ref {
	return bindings.MockMicrophoneConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MockMicrophoneConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MockMicrophoneConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MockMicrophoneConfiguration) Update(ref js.Ref) {
	bindings.MockMicrophoneConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MockSensor struct {
	// MaxSamplingFrequency is "MockSensor.maxSamplingFrequency"
	//
	// Optional
	MaxSamplingFrequency float64
	// MinSamplingFrequency is "MockSensor.minSamplingFrequency"
	//
	// Optional
	MinSamplingFrequency float64
	// RequestedSamplingFrequency is "MockSensor.requestedSamplingFrequency"
	//
	// Optional
	RequestedSamplingFrequency float64

	FFI_USE_MaxSamplingFrequency       bool // for MaxSamplingFrequency.
	FFI_USE_MinSamplingFrequency       bool // for MinSamplingFrequency.
	FFI_USE_RequestedSamplingFrequency bool // for RequestedSamplingFrequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MockSensor with all fields set.
func (p MockSensor) FromRef(ref js.Ref) MockSensor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MockSensor MockSensor [// MockSensor] [0x140008ff680 0x140008ff7c0 0x140008ff900 0x140008ff720 0x140008ff860 0x140008ff9a0] 0x140008bf908 {0 0}} in the application heap.
func (p MockSensor) New() js.Ref {
	return bindings.MockSensorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MockSensor) UpdateFrom(ref js.Ref) {
	bindings.MockSensorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MockSensor) Update(ref js.Ref) {
	bindings.MockSensorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MockSensorType uint32

const (
	_ MockSensorType = iota

	MockSensorType_AMBIENT_LIGHT
	MockSensorType_ACCELEROMETER
	MockSensorType_LINEAR_ACCELERATION
	MockSensorType_GRAVITY
	MockSensorType_GYROSCOPE
	MockSensorType_MAGNETOMETER
	MockSensorType_UNCALIBRATED_MAGNETOMETER
	MockSensorType_ABSOLUTE_ORIENTATION
	MockSensorType_RELATIVE_ORIENTATION
	MockSensorType_GEOLOCATION
	MockSensorType_PROXIMITY
)

func (MockSensorType) FromRef(str js.Ref) MockSensorType {
	return MockSensorType(bindings.ConstOfMockSensorType(str))
}

func (x MockSensorType) String() (string, bool) {
	switch x {
	case MockSensorType_AMBIENT_LIGHT:
		return "ambient-light", true
	case MockSensorType_ACCELEROMETER:
		return "accelerometer", true
	case MockSensorType_LINEAR_ACCELERATION:
		return "linear-acceleration", true
	case MockSensorType_GRAVITY:
		return "gravity", true
	case MockSensorType_GYROSCOPE:
		return "gyroscope", true
	case MockSensorType_MAGNETOMETER:
		return "magnetometer", true
	case MockSensorType_UNCALIBRATED_MAGNETOMETER:
		return "uncalibrated-magnetometer", true
	case MockSensorType_ABSOLUTE_ORIENTATION:
		return "absolute-orientation", true
	case MockSensorType_RELATIVE_ORIENTATION:
		return "relative-orientation", true
	case MockSensorType_GEOLOCATION:
		return "geolocation", true
	case MockSensorType_PROXIMITY:
		return "proximity", true
	default:
		return "", false
	}
}

type MockSensorConfiguration struct {
	// MockSensorType is "MockSensorConfiguration.mockSensorType"
	//
	// Required
	MockSensorType MockSensorType
	// Connected is "MockSensorConfiguration.connected"
	//
	// Optional, defaults to true.
	Connected bool
	// MaxSamplingFrequency is "MockSensorConfiguration.maxSamplingFrequency"
	//
	// Optional
	MaxSamplingFrequency float64
	// MinSamplingFrequency is "MockSensorConfiguration.minSamplingFrequency"
	//
	// Optional
	MinSamplingFrequency float64

	FFI_USE_Connected            bool // for Connected.
	FFI_USE_MaxSamplingFrequency bool // for MaxSamplingFrequency.
	FFI_USE_MinSamplingFrequency bool // for MinSamplingFrequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MockSensorConfiguration with all fields set.
func (p MockSensorConfiguration) FromRef(ref js.Ref) MockSensorConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MockSensorConfiguration MockSensorConfiguration [// MockSensorConfiguration] [0x140008ffa40 0x140008ffae0 0x140008ffc20 0x140008ffd60 0x140008ffb80 0x140008ffcc0 0x140008ffe00] 0x140008bf920 {0 0}} in the application heap.
func (p MockSensorConfiguration) New() js.Ref {
	return bindings.MockSensorConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MockSensorConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MockSensorConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MockSensorConfiguration) Update(ref js.Ref) {
	bindings.MockSensorConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MockSensorReadingValues struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MockSensorReadingValues with all fields set.
func (p MockSensorReadingValues) FromRef(ref js.Ref) MockSensorReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MockSensorReadingValues MockSensorReadingValues [// MockSensorReadingValues] [] 0x140008bfaa0 {0 0}} in the application heap.
func (p MockSensorReadingValues) New() js.Ref {
	return bindings.MockSensorReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MockSensorReadingValues) UpdateFrom(ref js.Ref) {
	bindings.MockSensorReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MockSensorReadingValues) Update(ref js.Ref) {
	bindings.MockSensorReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MouseEventInit struct {
	// ScreenX is "MouseEventInit.screenX"
	//
	// Optional, defaults to 0.
	ScreenX int32
	// ScreenY is "MouseEventInit.screenY"
	//
	// Optional, defaults to 0.
	ScreenY int32
	// ClientX is "MouseEventInit.clientX"
	//
	// Optional, defaults to 0.
	ClientX int32
	// ClientY is "MouseEventInit.clientY"
	//
	// Optional, defaults to 0.
	ClientY int32
	// Button is "MouseEventInit.button"
	//
	// Optional, defaults to 0.
	Button int16
	// Buttons is "MouseEventInit.buttons"
	//
	// Optional, defaults to 0.
	Buttons uint16
	// RelatedTarget is "MouseEventInit.relatedTarget"
	//
	// Optional, defaults to null.
	RelatedTarget EventTarget
	// CtrlKey is "MouseEventInit.ctrlKey"
	//
	// Optional, defaults to false.
	CtrlKey bool
	// ShiftKey is "MouseEventInit.shiftKey"
	//
	// Optional, defaults to false.
	ShiftKey bool
	// AltKey is "MouseEventInit.altKey"
	//
	// Optional, defaults to false.
	AltKey bool
	// MetaKey is "MouseEventInit.metaKey"
	//
	// Optional, defaults to false.
	MetaKey bool
	// ModifierAltGraph is "MouseEventInit.modifierAltGraph"
	//
	// Optional, defaults to false.
	ModifierAltGraph bool
	// ModifierCapsLock is "MouseEventInit.modifierCapsLock"
	//
	// Optional, defaults to false.
	ModifierCapsLock bool
	// ModifierFn is "MouseEventInit.modifierFn"
	//
	// Optional, defaults to false.
	ModifierFn bool
	// ModifierFnLock is "MouseEventInit.modifierFnLock"
	//
	// Optional, defaults to false.
	ModifierFnLock bool
	// ModifierHyper is "MouseEventInit.modifierHyper"
	//
	// Optional, defaults to false.
	ModifierHyper bool
	// ModifierNumLock is "MouseEventInit.modifierNumLock"
	//
	// Optional, defaults to false.
	ModifierNumLock bool
	// ModifierScrollLock is "MouseEventInit.modifierScrollLock"
	//
	// Optional, defaults to false.
	ModifierScrollLock bool
	// ModifierSuper is "MouseEventInit.modifierSuper"
	//
	// Optional, defaults to false.
	ModifierSuper bool
	// ModifierSymbol is "MouseEventInit.modifierSymbol"
	//
	// Optional, defaults to false.
	ModifierSymbol bool
	// ModifierSymbolLock is "MouseEventInit.modifierSymbolLock"
	//
	// Optional, defaults to false.
	ModifierSymbolLock bool
	// View is "MouseEventInit.view"
	//
	// Optional, defaults to null.
	View Window
	// Detail is "MouseEventInit.detail"
	//
	// Optional, defaults to 0.
	Detail int32
	// Bubbles is "MouseEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "MouseEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "MouseEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool
	// MovementX is "MouseEventInit.movementX"
	//
	// Optional, defaults to 0.
	MovementX float64
	// MovementY is "MouseEventInit.movementY"
	//
	// Optional, defaults to 0.
	MovementY float64

	FFI_USE_ScreenX            bool // for ScreenX.
	FFI_USE_ScreenY            bool // for ScreenY.
	FFI_USE_ClientX            bool // for ClientX.
	FFI_USE_ClientY            bool // for ClientY.
	FFI_USE_Button             bool // for Button.
	FFI_USE_Buttons            bool // for Buttons.
	FFI_USE_CtrlKey            bool // for CtrlKey.
	FFI_USE_ShiftKey           bool // for ShiftKey.
	FFI_USE_AltKey             bool // for AltKey.
	FFI_USE_MetaKey            bool // for MetaKey.
	FFI_USE_ModifierAltGraph   bool // for ModifierAltGraph.
	FFI_USE_ModifierCapsLock   bool // for ModifierCapsLock.
	FFI_USE_ModifierFn         bool // for ModifierFn.
	FFI_USE_ModifierFnLock     bool // for ModifierFnLock.
	FFI_USE_ModifierHyper      bool // for ModifierHyper.
	FFI_USE_ModifierNumLock    bool // for ModifierNumLock.
	FFI_USE_ModifierScrollLock bool // for ModifierScrollLock.
	FFI_USE_ModifierSuper      bool // for ModifierSuper.
	FFI_USE_ModifierSymbol     bool // for ModifierSymbol.
	FFI_USE_ModifierSymbolLock bool // for ModifierSymbolLock.
	FFI_USE_Detail             bool // for Detail.
	FFI_USE_Bubbles            bool // for Bubbles.
	FFI_USE_Cancelable         bool // for Cancelable.
	FFI_USE_Composed           bool // for Composed.
	FFI_USE_MovementX          bool // for MovementX.
	FFI_USE_MovementY          bool // for MovementY.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MouseEventInit with all fields set.
func (p MouseEventInit) FromRef(ref js.Ref) MouseEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MouseEventInit MouseEventInit [// MouseEventInit] [0x140008ffea0 0x14000904000 0x14000904140 0x14000904280 0x140009043c0 0x14000904500 0x14000904640 0x140009046e0 0x14000904820 0x14000904960 0x14000904aa0 0x14000904be0 0x14000904d20 0x14000904e60 0x14000904fa0 0x140009050e0 0x14000905220 0x14000905360 0x140009054a0 0x140009055e0 0x14000905720 0x14000905860 0x14000905900 0x14000905a40 0x14000905b80 0x14000905cc0 0x14000905e00 0x14000905f40 0x140008fff40 0x140009040a0 0x140009041e0 0x14000904320 0x14000904460 0x140009045a0 0x14000904780 0x140009048c0 0x14000904a00 0x14000904b40 0x14000904c80 0x14000904dc0 0x14000904f00 0x14000905040 0x14000905180 0x140009052c0 0x14000905400 0x14000905540 0x14000905680 0x140009057c0 0x140009059a0 0x14000905ae0 0x14000905c20 0x14000905d60 0x14000905ea0 0x1400090a000] 0x140008bfae8 {0 0}} in the application heap.
func (p MouseEventInit) New() js.Ref {
	return bindings.MouseEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MouseEventInit) UpdateFrom(ref js.Ref) {
	bindings.MouseEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MouseEventInit) Update(ref js.Ref) {
	bindings.MouseEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMouseEvent(typ js.String, eventInitDict MouseEventInit) MouseEvent {
	return MouseEvent{}.FromRef(
		bindings.NewMouseEventByMouseEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewMouseEventByMouseEvent1(typ js.String) MouseEvent {
	return MouseEvent{}.FromRef(
		bindings.NewMouseEventByMouseEvent1(
			typ.Ref()),
	)
}

type MouseEvent struct {
	UIEvent
}

func (this MouseEvent) Once() MouseEvent {
	this.Ref().Once()
	return this
}

func (this MouseEvent) Ref() js.Ref {
	return this.UIEvent.Ref()
}

func (this MouseEvent) FromRef(ref js.Ref) MouseEvent {
	this.UIEvent = this.UIEvent.FromRef(ref)
	return this
}

func (this MouseEvent) Free() {
	this.Ref().Free()
}

// ScreenX returns the value of property "MouseEvent.screenX".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) ScreenX() (int32, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventScreenX(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ScreenY returns the value of property "MouseEvent.screenY".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) ScreenY() (int32, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventScreenY(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ClientX returns the value of property "MouseEvent.clientX".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) ClientX() (int32, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventClientX(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ClientY returns the value of property "MouseEvent.clientY".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) ClientY() (int32, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventClientY(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// CtrlKey returns the value of property "MouseEvent.ctrlKey".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) CtrlKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventCtrlKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ShiftKey returns the value of property "MouseEvent.shiftKey".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) ShiftKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventShiftKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// AltKey returns the value of property "MouseEvent.altKey".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) AltKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventAltKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// MetaKey returns the value of property "MouseEvent.metaKey".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) MetaKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventMetaKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Button returns the value of property "MouseEvent.button".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) Button() (int16, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventButton(
		this.Ref(), js.Pointer(&_ok),
	)
	return int16(_ret), _ok
}

// Buttons returns the value of property "MouseEvent.buttons".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) Buttons() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventButtons(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// RelatedTarget returns the value of property "MouseEvent.relatedTarget".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) RelatedTarget() (EventTarget, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventRelatedTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return EventTarget{}.FromRef(_ret), _ok
}

// MovementX returns the value of property "MouseEvent.movementX".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) MovementX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventMovementX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// MovementY returns the value of property "MouseEvent.movementY".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) MovementY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventMovementY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PageX returns the value of property "MouseEvent.pageX".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) PageX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventPageX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// PageY returns the value of property "MouseEvent.pageY".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) PageY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventPageY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// X returns the value of property "MouseEvent.x".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) X() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Y returns the value of property "MouseEvent.y".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) Y() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// OffsetX returns the value of property "MouseEvent.offsetX".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) OffsetX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventOffsetX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// OffsetY returns the value of property "MouseEvent.offsetY".
//
// The returned bool will be false if there is no such property.
func (this MouseEvent) OffsetY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetMouseEventOffsetY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// GetModifierState calls the method "MouseEvent.getModifierState".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) GetModifierState(keyArg js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventGetModifierState(
		this.Ref(), js.Pointer(&_ok),
		keyArg.Ref(),
	)

	return _ret == js.True, _ok
}

// GetModifierStateFunc returns the method "MouseEvent.getModifierState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) GetModifierStateFunc() (fn js.Func[func(keyArg js.String) bool]) {
	return fn.FromRef(
		bindings.MouseEventGetModifierStateFunc(
			this.Ref(),
		),
	)
}

// InitMouseEvent calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16, relatedTargetArg EventTarget) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
		int32(clientXArg),
		int32(clientYArg),
		js.Bool(bool(ctrlKeyArg)),
		js.Bool(bool(altKeyArg)),
		js.Bool(bool(shiftKeyArg)),
		js.Bool(bool(metaKeyArg)),
		int32(buttonArg),
		relatedTargetArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEventFunc returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEventFunc() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16, relatedTargetArg EventTarget)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEventFunc(
			this.Ref(),
		),
	)
}

// InitMouseEvent1 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent1(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
		int32(clientXArg),
		int32(clientYArg),
		js.Bool(bool(ctrlKeyArg)),
		js.Bool(bool(altKeyArg)),
		js.Bool(bool(shiftKeyArg)),
		js.Bool(bool(metaKeyArg)),
		int32(buttonArg),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent1Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent1Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent1Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent2 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent2(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
		int32(clientXArg),
		int32(clientYArg),
		js.Bool(bool(ctrlKeyArg)),
		js.Bool(bool(altKeyArg)),
		js.Bool(bool(shiftKeyArg)),
		js.Bool(bool(metaKeyArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent2Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent2Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent2Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent3 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent3(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
		int32(clientXArg),
		int32(clientYArg),
		js.Bool(bool(ctrlKeyArg)),
		js.Bool(bool(altKeyArg)),
		js.Bool(bool(shiftKeyArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent3Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent3Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent3Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent4 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent4(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
		int32(clientXArg),
		int32(clientYArg),
		js.Bool(bool(ctrlKeyArg)),
		js.Bool(bool(altKeyArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent4Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent4Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent4Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent5 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent5(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
		int32(clientXArg),
		int32(clientYArg),
		js.Bool(bool(ctrlKeyArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent5Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent5Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent5Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent6 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent6(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent6(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
		int32(clientXArg),
		int32(clientYArg),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent6Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent6Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent6Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent7 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent7(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent7(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
		int32(clientXArg),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent7Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent7Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent7Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent8 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent8(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent8(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent8Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent8Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent8Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent9 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent9(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent9(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent9Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent9Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent9Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent10 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent10(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent10(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent10Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent10Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent10Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent11 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent11(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent11(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent11Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent11Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent11Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent12 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent12(typeArg js.String, bubblesArg bool, cancelableArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent12(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent12Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent12Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent12Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent13 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent13(typeArg js.String, bubblesArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent13(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent13Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent13Func() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent13Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent14 calls the method "MouseEvent.initMouseEvent".
//
// The returned bool will be false if there is no such method.
func (this MouseEvent) InitMouseEvent14(typeArg js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMouseEventInitMouseEvent14(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMouseEvent14Func returns the method "MouseEvent.initMouseEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MouseEvent) InitMouseEvent14Func() (fn js.Func[func(typeArg js.String)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent14Func(
			this.Ref(),
		),
	)
}

type MutationCallbackFunc func(this js.Ref, mutations js.Array[MutationRecord], observer MutationObserver) js.Ref

func (fn MutationCallbackFunc) Register() js.Func[func(mutations js.Array[MutationRecord], observer MutationObserver)] {
	return js.RegisterCallback[func(mutations js.Array[MutationRecord], observer MutationObserver)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn MutationCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[MutationRecord]{}.FromRef(args[0+1]),
		MutationObserver{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type MutationCallback[T any] struct {
	Fn  func(arg T, this js.Ref, mutations js.Array[MutationRecord], observer MutationObserver) js.Ref
	Arg T
}

func (cb *MutationCallback[T]) Register() js.Func[func(mutations js.Array[MutationRecord], observer MutationObserver)] {
	return js.RegisterCallback[func(mutations js.Array[MutationRecord], observer MutationObserver)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *MutationCallback[T]) DispatchCallback(
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

		js.Array[MutationRecord]{}.FromRef(args[0+1]),
		MutationObserver{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type MutationRecord struct {
	ref js.Ref
}

func (this MutationRecord) Once() MutationRecord {
	this.Ref().Once()
	return this
}

func (this MutationRecord) Ref() js.Ref {
	return this.ref
}

func (this MutationRecord) FromRef(ref js.Ref) MutationRecord {
	this.ref = ref
	return this
}

func (this MutationRecord) Free() {
	this.Ref().Free()
}

// Type returns the value of property "MutationRecord.type".
//
// The returned bool will be false if there is no such property.
func (this MutationRecord) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMutationRecordType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Target returns the value of property "MutationRecord.target".
//
// The returned bool will be false if there is no such property.
func (this MutationRecord) Target() (Node, bool) {
	var _ok bool
	_ret := bindings.GetMutationRecordTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// AddedNodes returns the value of property "MutationRecord.addedNodes".
//
// The returned bool will be false if there is no such property.
func (this MutationRecord) AddedNodes() (NodeList, bool) {
	var _ok bool
	_ret := bindings.GetMutationRecordAddedNodes(
		this.Ref(), js.Pointer(&_ok),
	)
	return NodeList{}.FromRef(_ret), _ok
}

// RemovedNodes returns the value of property "MutationRecord.removedNodes".
//
// The returned bool will be false if there is no such property.
func (this MutationRecord) RemovedNodes() (NodeList, bool) {
	var _ok bool
	_ret := bindings.GetMutationRecordRemovedNodes(
		this.Ref(), js.Pointer(&_ok),
	)
	return NodeList{}.FromRef(_ret), _ok
}

// PreviousSibling returns the value of property "MutationRecord.previousSibling".
//
// The returned bool will be false if there is no such property.
func (this MutationRecord) PreviousSibling() (Node, bool) {
	var _ok bool
	_ret := bindings.GetMutationRecordPreviousSibling(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// NextSibling returns the value of property "MutationRecord.nextSibling".
//
// The returned bool will be false if there is no such property.
func (this MutationRecord) NextSibling() (Node, bool) {
	var _ok bool
	_ret := bindings.GetMutationRecordNextSibling(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// AttributeName returns the value of property "MutationRecord.attributeName".
//
// The returned bool will be false if there is no such property.
func (this MutationRecord) AttributeName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMutationRecordAttributeName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AttributeNamespace returns the value of property "MutationRecord.attributeNamespace".
//
// The returned bool will be false if there is no such property.
func (this MutationRecord) AttributeNamespace() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMutationRecordAttributeNamespace(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// OldValue returns the value of property "MutationRecord.oldValue".
//
// The returned bool will be false if there is no such property.
func (this MutationRecord) OldValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMutationRecordOldValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type MutationObserverInit struct {
	// ChildList is "MutationObserverInit.childList"
	//
	// Optional, defaults to false.
	ChildList bool
	// Attributes is "MutationObserverInit.attributes"
	//
	// Optional
	Attributes bool
	// CharacterData is "MutationObserverInit.characterData"
	//
	// Optional
	CharacterData bool
	// Subtree is "MutationObserverInit.subtree"
	//
	// Optional, defaults to false.
	Subtree bool
	// AttributeOldValue is "MutationObserverInit.attributeOldValue"
	//
	// Optional
	AttributeOldValue bool
	// CharacterDataOldValue is "MutationObserverInit.characterDataOldValue"
	//
	// Optional
	CharacterDataOldValue bool
	// AttributeFilter is "MutationObserverInit.attributeFilter"
	//
	// Optional
	AttributeFilter js.Array[js.String]

	FFI_USE_ChildList             bool // for ChildList.
	FFI_USE_Attributes            bool // for Attributes.
	FFI_USE_CharacterData         bool // for CharacterData.
	FFI_USE_Subtree               bool // for Subtree.
	FFI_USE_AttributeOldValue     bool // for AttributeOldValue.
	FFI_USE_CharacterDataOldValue bool // for CharacterDataOldValue.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MutationObserverInit with all fields set.
func (p MutationObserverInit) FromRef(ref js.Ref) MutationObserverInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MutationObserverInit MutationObserverInit [// MutationObserverInit] [0x1400090a1e0 0x1400090a320 0x1400090a460 0x1400090a5a0 0x1400090a6e0 0x1400090a820 0x1400090a960 0x1400090a280 0x1400090a3c0 0x1400090a500 0x1400090a640 0x1400090a780 0x1400090a8c0] 0x140008bfd10 {0 0}} in the application heap.
func (p MutationObserverInit) New() js.Ref {
	return bindings.MutationObserverInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MutationObserverInit) UpdateFrom(ref js.Ref) {
	bindings.MutationObserverInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MutationObserverInit) Update(ref js.Ref) {
	bindings.MutationObserverInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewMutationObserver(callback js.Func[func(mutations js.Array[MutationRecord], observer MutationObserver)]) MutationObserver {
	return MutationObserver{}.FromRef(
		bindings.NewMutationObserverByMutationObserver(
			callback.Ref()),
	)
}

type MutationObserver struct {
	ref js.Ref
}

func (this MutationObserver) Once() MutationObserver {
	this.Ref().Once()
	return this
}

func (this MutationObserver) Ref() js.Ref {
	return this.ref
}

func (this MutationObserver) FromRef(ref js.Ref) MutationObserver {
	this.ref = ref
	return this
}

func (this MutationObserver) Free() {
	this.Ref().Free()
}

// Observe calls the method "MutationObserver.observe".
//
// The returned bool will be false if there is no such method.
func (this MutationObserver) Observe(target Node, options MutationObserverInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMutationObserverObserve(
		this.Ref(), js.Pointer(&_ok),
		target.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ObserveFunc returns the method "MutationObserver.observe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationObserver) ObserveFunc() (fn js.Func[func(target Node, options MutationObserverInit)]) {
	return fn.FromRef(
		bindings.MutationObserverObserveFunc(
			this.Ref(),
		),
	)
}

// Observe1 calls the method "MutationObserver.observe".
//
// The returned bool will be false if there is no such method.
func (this MutationObserver) Observe1(target Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMutationObserverObserve1(
		this.Ref(), js.Pointer(&_ok),
		target.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Observe1Func returns the method "MutationObserver.observe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationObserver) Observe1Func() (fn js.Func[func(target Node)]) {
	return fn.FromRef(
		bindings.MutationObserverObserve1Func(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "MutationObserver.disconnect".
//
// The returned bool will be false if there is no such method.
func (this MutationObserver) Disconnect() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMutationObserverDisconnect(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisconnectFunc returns the method "MutationObserver.disconnect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationObserver) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MutationObserverDisconnectFunc(
			this.Ref(),
		),
	)
}

// TakeRecords calls the method "MutationObserver.takeRecords".
//
// The returned bool will be false if there is no such method.
func (this MutationObserver) TakeRecords() (js.Array[MutationRecord], bool) {
	var _ok bool
	_ret := bindings.CallMutationObserverTakeRecords(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[MutationRecord]{}.FromRef(_ret), _ok
}

// TakeRecordsFunc returns the method "MutationObserver.takeRecords".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationObserver) TakeRecordsFunc() (fn js.Func[func() js.Array[MutationRecord]]) {
	return fn.FromRef(
		bindings.MutationObserverTakeRecordsFunc(
			this.Ref(),
		),
	)
}

const (
	MutationEvent_MODIFICATION uint16 = 1
	MutationEvent_ADDITION     uint16 = 2
	MutationEvent_REMOVAL      uint16 = 3
)

func NewMutationEvent(typ js.String, eventInitDict EventInit) MutationEvent {
	return MutationEvent{}.FromRef(
		bindings.NewMutationEventByMutationEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewMutationEventByMutationEvent1(typ js.String) MutationEvent {
	return MutationEvent{}.FromRef(
		bindings.NewMutationEventByMutationEvent1(
			typ.Ref()),
	)
}

type MutationEvent struct {
	Event
}

func (this MutationEvent) Once() MutationEvent {
	this.Ref().Once()
	return this
}

func (this MutationEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this MutationEvent) FromRef(ref js.Ref) MutationEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this MutationEvent) Free() {
	this.Ref().Free()
}

// RelatedNode returns the value of property "MutationEvent.relatedNode".
//
// The returned bool will be false if there is no such property.
func (this MutationEvent) RelatedNode() (Node, bool) {
	var _ok bool
	_ret := bindings.GetMutationEventRelatedNode(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// PrevValue returns the value of property "MutationEvent.prevValue".
//
// The returned bool will be false if there is no such property.
func (this MutationEvent) PrevValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMutationEventPrevValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// NewValue returns the value of property "MutationEvent.newValue".
//
// The returned bool will be false if there is no such property.
func (this MutationEvent) NewValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMutationEventNewValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AttrName returns the value of property "MutationEvent.attrName".
//
// The returned bool will be false if there is no such property.
func (this MutationEvent) AttrName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMutationEventAttrName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AttrChange returns the value of property "MutationEvent.attrChange".
//
// The returned bool will be false if there is no such property.
func (this MutationEvent) AttrChange() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetMutationEventAttrChange(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// InitMutationEvent calls the method "MutationEvent.initMutationEvent".
//
// The returned bool will be false if there is no such method.
func (this MutationEvent) InitMutationEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String, attrChangeArg uint16) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMutationEventInitMutationEvent(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
		newValueArg.Ref(),
		attrNameArg.Ref(),
		uint32(attrChangeArg),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMutationEventFunc returns the method "MutationEvent.initMutationEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationEvent) InitMutationEventFunc() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String, attrChangeArg uint16)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEventFunc(
			this.Ref(),
		),
	)
}

// InitMutationEvent1 calls the method "MutationEvent.initMutationEvent".
//
// The returned bool will be false if there is no such method.
func (this MutationEvent) InitMutationEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMutationEventInitMutationEvent1(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
		newValueArg.Ref(),
		attrNameArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMutationEvent1Func returns the method "MutationEvent.initMutationEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationEvent) InitMutationEvent1Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent1Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent2 calls the method "MutationEvent.initMutationEvent".
//
// The returned bool will be false if there is no such method.
func (this MutationEvent) InitMutationEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMutationEventInitMutationEvent2(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
		newValueArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMutationEvent2Func returns the method "MutationEvent.initMutationEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationEvent) InitMutationEvent2Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent2Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent3 calls the method "MutationEvent.initMutationEvent".
//
// The returned bool will be false if there is no such method.
func (this MutationEvent) InitMutationEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMutationEventInitMutationEvent3(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMutationEvent3Func returns the method "MutationEvent.initMutationEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationEvent) InitMutationEvent3Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent3Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent4 calls the method "MutationEvent.initMutationEvent".
//
// The returned bool will be false if there is no such method.
func (this MutationEvent) InitMutationEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMutationEventInitMutationEvent4(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMutationEvent4Func returns the method "MutationEvent.initMutationEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationEvent) InitMutationEvent4Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent4Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent5 calls the method "MutationEvent.initMutationEvent".
//
// The returned bool will be false if there is no such method.
func (this MutationEvent) InitMutationEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMutationEventInitMutationEvent5(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMutationEvent5Func returns the method "MutationEvent.initMutationEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationEvent) InitMutationEvent5Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent5Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent6 calls the method "MutationEvent.initMutationEvent".
//
// The returned bool will be false if there is no such method.
func (this MutationEvent) InitMutationEvent6(typeArg js.String, bubblesArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMutationEventInitMutationEvent6(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMutationEvent6Func returns the method "MutationEvent.initMutationEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationEvent) InitMutationEvent6Func() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent6Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent7 calls the method "MutationEvent.initMutationEvent".
//
// The returned bool will be false if there is no such method.
func (this MutationEvent) InitMutationEvent7(typeArg js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallMutationEventInitMutationEvent7(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitMutationEvent7Func returns the method "MutationEvent.initMutationEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MutationEvent) InitMutationEvent7Func() (fn js.Func[func(typeArg js.String)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent7Func(
			this.Ref(),
		),
	)
}

type NDEFMakeReadOnlyOptions struct {
	// Signal is "NDEFMakeReadOnlyOptions.signal"
	//
	// Optional
	Signal AbortSignal

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NDEFMakeReadOnlyOptions with all fields set.
func (p NDEFMakeReadOnlyOptions) FromRef(ref js.Ref) NDEFMakeReadOnlyOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NDEFMakeReadOnlyOptions NDEFMakeReadOnlyOptions [// NDEFMakeReadOnlyOptions] [0x1400090aaa0] 0x140009200d8 {0 0}} in the application heap.
func (p NDEFMakeReadOnlyOptions) New() js.Ref {
	return bindings.NDEFMakeReadOnlyOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NDEFMakeReadOnlyOptions) UpdateFrom(ref js.Ref) {
	bindings.NDEFMakeReadOnlyOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NDEFMakeReadOnlyOptions) Update(ref js.Ref) {
	bindings.NDEFMakeReadOnlyOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type NDEFRecordInit struct {
	// RecordType is "NDEFRecordInit.recordType"
	//
	// Required
	RecordType js.String
	// MediaType is "NDEFRecordInit.mediaType"
	//
	// Optional
	MediaType js.String
	// Id is "NDEFRecordInit.id"
	//
	// Optional
	Id js.String
	// Encoding is "NDEFRecordInit.encoding"
	//
	// Optional
	Encoding js.String
	// Lang is "NDEFRecordInit.lang"
	//
	// Optional
	Lang js.String
	// Data is "NDEFRecordInit.data"
	//
	// Optional
	Data js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NDEFRecordInit with all fields set.
func (p NDEFRecordInit) FromRef(ref js.Ref) NDEFRecordInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NDEFRecordInit NDEFRecordInit [// NDEFRecordInit] [0x1400090ab40 0x1400090abe0 0x1400090ac80 0x1400090ad20 0x1400090adc0 0x1400090ae60] 0x14000920138 {0 0}} in the application heap.
func (p NDEFRecordInit) New() js.Ref {
	return bindings.NDEFRecordInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NDEFRecordInit) UpdateFrom(ref js.Ref) {
	bindings.NDEFRecordInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NDEFRecordInit) Update(ref js.Ref) {
	bindings.NDEFRecordInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type NDEFMessageInit struct {
	// Records is "NDEFMessageInit.records"
	//
	// Required
	Records js.Array[NDEFRecordInit]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NDEFMessageInit with all fields set.
func (p NDEFMessageInit) FromRef(ref js.Ref) NDEFMessageInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NDEFMessageInit NDEFMessageInit [// NDEFMessageInit] [0x1400090af00] 0x14000920108 {0 0}} in the application heap.
func (p NDEFMessageInit) New() js.Ref {
	return bindings.NDEFMessageInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NDEFMessageInit) UpdateFrom(ref js.Ref) {
	bindings.NDEFMessageInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NDEFMessageInit) Update(ref js.Ref) {
	bindings.NDEFMessageInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewNDEFRecord(recordInit NDEFRecordInit) NDEFRecord {
	return NDEFRecord{}.FromRef(
		bindings.NewNDEFRecordByNDEFRecord(
			js.Pointer(&recordInit)),
	)
}

type NDEFRecord struct {
	ref js.Ref
}

func (this NDEFRecord) Once() NDEFRecord {
	this.Ref().Once()
	return this
}

func (this NDEFRecord) Ref() js.Ref {
	return this.ref
}

func (this NDEFRecord) FromRef(ref js.Ref) NDEFRecord {
	this.ref = ref
	return this
}

func (this NDEFRecord) Free() {
	this.Ref().Free()
}

// RecordType returns the value of property "NDEFRecord.recordType".
//
// The returned bool will be false if there is no such property.
func (this NDEFRecord) RecordType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNDEFRecordRecordType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// MediaType returns the value of property "NDEFRecord.mediaType".
//
// The returned bool will be false if there is no such property.
func (this NDEFRecord) MediaType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNDEFRecordMediaType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Id returns the value of property "NDEFRecord.id".
//
// The returned bool will be false if there is no such property.
func (this NDEFRecord) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNDEFRecordId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Data returns the value of property "NDEFRecord.data".
//
// The returned bool will be false if there is no such property.
func (this NDEFRecord) Data() (js.DataView, bool) {
	var _ok bool
	_ret := bindings.GetNDEFRecordData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.DataView{}.FromRef(_ret), _ok
}

// Encoding returns the value of property "NDEFRecord.encoding".
//
// The returned bool will be false if there is no such property.
func (this NDEFRecord) Encoding() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNDEFRecordEncoding(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Lang returns the value of property "NDEFRecord.lang".
//
// The returned bool will be false if there is no such property.
func (this NDEFRecord) Lang() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNDEFRecordLang(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ToRecords calls the method "NDEFRecord.toRecords".
//
// The returned bool will be false if there is no such method.
func (this NDEFRecord) ToRecords() (js.Array[NDEFRecord], bool) {
	var _ok bool
	_ret := bindings.CallNDEFRecordToRecords(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[NDEFRecord]{}.FromRef(_ret), _ok
}

// ToRecordsFunc returns the method "NDEFRecord.toRecords".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NDEFRecord) ToRecordsFunc() (fn js.Func[func() js.Array[NDEFRecord]]) {
	return fn.FromRef(
		bindings.NDEFRecordToRecordsFunc(
			this.Ref(),
		),
	)
}

func NewNDEFMessage(messageInit NDEFMessageInit) NDEFMessage {
	return NDEFMessage{}.FromRef(
		bindings.NewNDEFMessageByNDEFMessage(
			js.Pointer(&messageInit)),
	)
}

type NDEFMessage struct {
	ref js.Ref
}

func (this NDEFMessage) Once() NDEFMessage {
	this.Ref().Once()
	return this
}

func (this NDEFMessage) Ref() js.Ref {
	return this.ref
}

func (this NDEFMessage) FromRef(ref js.Ref) NDEFMessage {
	this.ref = ref
	return this
}

func (this NDEFMessage) Free() {
	this.Ref().Free()
}

// Records returns the value of property "NDEFMessage.records".
//
// The returned bool will be false if there is no such property.
func (this NDEFMessage) Records() (js.FrozenArray[NDEFRecord], bool) {
	var _ok bool
	_ret := bindings.GetNDEFMessageRecords(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[NDEFRecord]{}.FromRef(_ret), _ok
}

type OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit struct {
	ref js.Ref
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) Free() {
	x.ref.Free()
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) FromRef(ref js.Ref) OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit {
	return OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit{
		ref: ref,
	}
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) TypedArrayInt8() js.TypedArray[int8] {
	return js.TypedArray[int8]{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) TypedArrayInt16() js.TypedArray[int16] {
	return js.TypedArray[int16]{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) TypedArrayInt32() js.TypedArray[int32] {
	return js.TypedArray[int32]{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) TypedArrayUint16() js.TypedArray[uint16] {
	return js.TypedArray[uint16]{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) TypedArrayUint32() js.TypedArray[uint32] {
	return js.TypedArray[uint32]{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) TypedArrayInt64() js.TypedArray[int64] {
	return js.TypedArray[int64]{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) TypedArrayUint64() js.TypedArray[uint64] {
	return js.TypedArray[uint64]{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) DataView() js.DataView {
	return js.DataView{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

func (x OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit) NDEFMessageInit() NDEFMessageInit {
	var ret NDEFMessageInit
	ret.UpdateFrom(x.ref)
	return ret
}

type NDEFMessageSource = OneOf_String_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_NDEFMessageInit

type NDEFScanOptions struct {
	// Signal is "NDEFScanOptions.signal"
	//
	// Optional
	Signal AbortSignal

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NDEFScanOptions with all fields set.
func (p NDEFScanOptions) FromRef(ref js.Ref) NDEFScanOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NDEFScanOptions NDEFScanOptions [// NDEFScanOptions] [0x1400090b040] 0x14000920258 {0 0}} in the application heap.
func (p NDEFScanOptions) New() js.Ref {
	return bindings.NDEFScanOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NDEFScanOptions) UpdateFrom(ref js.Ref) {
	bindings.NDEFScanOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NDEFScanOptions) Update(ref js.Ref) {
	bindings.NDEFScanOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type NDEFWriteOptions struct {
	// Overwrite is "NDEFWriteOptions.overwrite"
	//
	// Optional, defaults to true.
	Overwrite bool
	// Signal is "NDEFWriteOptions.signal"
	//
	// Optional
	Signal AbortSignal

	FFI_USE_Overwrite bool // for Overwrite.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NDEFWriteOptions with all fields set.
func (p NDEFWriteOptions) FromRef(ref js.Ref) NDEFWriteOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NDEFWriteOptions NDEFWriteOptions [// NDEFWriteOptions] [0x1400090b0e0 0x1400090b220 0x1400090b180] 0x140009202a0 {0 0}} in the application heap.
func (p NDEFWriteOptions) New() js.Ref {
	return bindings.NDEFWriteOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NDEFWriteOptions) UpdateFrom(ref js.Ref) {
	bindings.NDEFWriteOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NDEFWriteOptions) Update(ref js.Ref) {
	bindings.NDEFWriteOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type NDEFReader struct {
	EventTarget
}

func (this NDEFReader) Once() NDEFReader {
	this.Ref().Once()
	return this
}

func (this NDEFReader) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this NDEFReader) FromRef(ref js.Ref) NDEFReader {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this NDEFReader) Free() {
	this.Ref().Free()
}

// Scan calls the method "NDEFReader.scan".
//
// The returned bool will be false if there is no such method.
func (this NDEFReader) Scan(options NDEFScanOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNDEFReaderScan(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ScanFunc returns the method "NDEFReader.scan".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NDEFReader) ScanFunc() (fn js.Func[func(options NDEFScanOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderScanFunc(
			this.Ref(),
		),
	)
}

// Scan1 calls the method "NDEFReader.scan".
//
// The returned bool will be false if there is no such method.
func (this NDEFReader) Scan1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNDEFReaderScan1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Scan1Func returns the method "NDEFReader.scan".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NDEFReader) Scan1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderScan1Func(
			this.Ref(),
		),
	)
}

// Write calls the method "NDEFReader.write".
//
// The returned bool will be false if there is no such method.
func (this NDEFReader) Write(message NDEFMessageSource, options NDEFWriteOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNDEFReaderWrite(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// WriteFunc returns the method "NDEFReader.write".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NDEFReader) WriteFunc() (fn js.Func[func(message NDEFMessageSource, options NDEFWriteOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderWriteFunc(
			this.Ref(),
		),
	)
}

// Write1 calls the method "NDEFReader.write".
//
// The returned bool will be false if there is no such method.
func (this NDEFReader) Write1(message NDEFMessageSource) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNDEFReaderWrite1(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Write1Func returns the method "NDEFReader.write".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NDEFReader) Write1Func() (fn js.Func[func(message NDEFMessageSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderWrite1Func(
			this.Ref(),
		),
	)
}

// MakeReadOnly calls the method "NDEFReader.makeReadOnly".
//
// The returned bool will be false if there is no such method.
func (this NDEFReader) MakeReadOnly(options NDEFMakeReadOnlyOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNDEFReaderMakeReadOnly(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// MakeReadOnlyFunc returns the method "NDEFReader.makeReadOnly".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NDEFReader) MakeReadOnlyFunc() (fn js.Func[func(options NDEFMakeReadOnlyOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderMakeReadOnlyFunc(
			this.Ref(),
		),
	)
}

// MakeReadOnly1 calls the method "NDEFReader.makeReadOnly".
//
// The returned bool will be false if there is no such method.
func (this NDEFReader) MakeReadOnly1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallNDEFReaderMakeReadOnly1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// MakeReadOnly1Func returns the method "NDEFReader.makeReadOnly".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NDEFReader) MakeReadOnly1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderMakeReadOnly1Func(
			this.Ref(),
		),
	)
}

type NDEFReadingEventInit struct {
	// SerialNumber is "NDEFReadingEventInit.serialNumber"
	//
	// Optional, defaults to "".
	SerialNumber js.String
	// Message is "NDEFReadingEventInit.message"
	//
	// Required
	Message NDEFMessageInit
	// Bubbles is "NDEFReadingEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "NDEFReadingEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "NDEFReadingEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NDEFReadingEventInit with all fields set.
func (p NDEFReadingEventInit) FromRef(ref js.Ref) NDEFReadingEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NDEFReadingEventInit NDEFReadingEventInit [// NDEFReadingEventInit] [0x1400090b2c0 0x1400090b360 0x1400090b400 0x1400090b540 0x1400090b680 0x1400090b4a0 0x1400090b5e0 0x1400090b720] 0x14000920300 {0 0}} in the application heap.
func (p NDEFReadingEventInit) New() js.Ref {
	return bindings.NDEFReadingEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NDEFReadingEventInit) UpdateFrom(ref js.Ref) {
	bindings.NDEFReadingEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NDEFReadingEventInit) Update(ref js.Ref) {
	bindings.NDEFReadingEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewNDEFReadingEvent(typ js.String, readingEventInitDict NDEFReadingEventInit) NDEFReadingEvent {
	return NDEFReadingEvent{}.FromRef(
		bindings.NewNDEFReadingEventByNDEFReadingEvent(
			typ.Ref(),
			js.Pointer(&readingEventInitDict)),
	)
}

type NDEFReadingEvent struct {
	Event
}

func (this NDEFReadingEvent) Once() NDEFReadingEvent {
	this.Ref().Once()
	return this
}

func (this NDEFReadingEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this NDEFReadingEvent) FromRef(ref js.Ref) NDEFReadingEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this NDEFReadingEvent) Free() {
	this.Ref().Free()
}

// SerialNumber returns the value of property "NDEFReadingEvent.serialNumber".
//
// The returned bool will be false if there is no such property.
func (this NDEFReadingEvent) SerialNumber() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNDEFReadingEventSerialNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Message returns the value of property "NDEFReadingEvent.message".
//
// The returned bool will be false if there is no such property.
func (this NDEFReadingEvent) Message() (NDEFMessage, bool) {
	var _ok bool
	_ret := bindings.GetNDEFReadingEventMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return NDEFMessage{}.FromRef(_ret), _ok
}

type NamedFlow struct {
	EventTarget
}

func (this NamedFlow) Once() NamedFlow {
	this.Ref().Once()
	return this
}

func (this NamedFlow) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this NamedFlow) FromRef(ref js.Ref) NamedFlow {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this NamedFlow) Free() {
	this.Ref().Free()
}

// Name returns the value of property "NamedFlow.name".
//
// The returned bool will be false if there is no such property.
func (this NamedFlow) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNamedFlowName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Overset returns the value of property "NamedFlow.overset".
//
// The returned bool will be false if there is no such property.
func (this NamedFlow) Overset() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNamedFlowOverset(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// FirstEmptyRegionIndex returns the value of property "NamedFlow.firstEmptyRegionIndex".
//
// The returned bool will be false if there is no such property.
func (this NamedFlow) FirstEmptyRegionIndex() (int16, bool) {
	var _ok bool
	_ret := bindings.GetNamedFlowFirstEmptyRegionIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int16(_ret), _ok
}

// GetRegions calls the method "NamedFlow.getRegions".
//
// The returned bool will be false if there is no such method.
func (this NamedFlow) GetRegions() (js.Array[Element], bool) {
	var _ok bool
	_ret := bindings.CallNamedFlowGetRegions(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[Element]{}.FromRef(_ret), _ok
}

// GetRegionsFunc returns the method "NamedFlow.getRegions".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NamedFlow) GetRegionsFunc() (fn js.Func[func() js.Array[Element]]) {
	return fn.FromRef(
		bindings.NamedFlowGetRegionsFunc(
			this.Ref(),
		),
	)
}

// GetContent calls the method "NamedFlow.getContent".
//
// The returned bool will be false if there is no such method.
func (this NamedFlow) GetContent() (js.Array[Node], bool) {
	var _ok bool
	_ret := bindings.CallNamedFlowGetContent(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[Node]{}.FromRef(_ret), _ok
}

// GetContentFunc returns the method "NamedFlow.getContent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NamedFlow) GetContentFunc() (fn js.Func[func() js.Array[Node]]) {
	return fn.FromRef(
		bindings.NamedFlowGetContentFunc(
			this.Ref(),
		),
	)
}

// GetRegionsByContent calls the method "NamedFlow.getRegionsByContent".
//
// The returned bool will be false if there is no such method.
func (this NamedFlow) GetRegionsByContent(node Node) (js.Array[Element], bool) {
	var _ok bool
	_ret := bindings.CallNamedFlowGetRegionsByContent(
		this.Ref(), js.Pointer(&_ok),
		node.Ref(),
	)

	return js.Array[Element]{}.FromRef(_ret), _ok
}

// GetRegionsByContentFunc returns the method "NamedFlow.getRegionsByContent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NamedFlow) GetRegionsByContentFunc() (fn js.Func[func(node Node) js.Array[Element]]) {
	return fn.FromRef(
		bindings.NamedFlowGetRegionsByContentFunc(
			this.Ref(),
		),
	)
}

type NavigationDestination struct {
	ref js.Ref
}

func (this NavigationDestination) Once() NavigationDestination {
	this.Ref().Once()
	return this
}

func (this NavigationDestination) Ref() js.Ref {
	return this.ref
}

func (this NavigationDestination) FromRef(ref js.Ref) NavigationDestination {
	this.ref = ref
	return this
}

func (this NavigationDestination) Free() {
	this.Ref().Free()
}

// Url returns the value of property "NavigationDestination.url".
//
// The returned bool will be false if there is no such property.
func (this NavigationDestination) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigationDestinationUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Key returns the value of property "NavigationDestination.key".
//
// The returned bool will be false if there is no such property.
func (this NavigationDestination) Key() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigationDestinationKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Id returns the value of property "NavigationDestination.id".
//
// The returned bool will be false if there is no such property.
func (this NavigationDestination) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigationDestinationId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Index returns the value of property "NavigationDestination.index".
//
// The returned bool will be false if there is no such property.
func (this NavigationDestination) Index() (int64, bool) {
	var _ok bool
	_ret := bindings.GetNavigationDestinationIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int64(_ret), _ok
}

// SameDocument returns the value of property "NavigationDestination.sameDocument".
//
// The returned bool will be false if there is no such property.
func (this NavigationDestination) SameDocument() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigationDestinationSameDocument(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// GetState calls the method "NavigationDestination.getState".
//
// The returned bool will be false if there is no such method.
func (this NavigationDestination) GetState() (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallNavigationDestinationGetState(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// GetStateFunc returns the method "NavigationDestination.getState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NavigationDestination) GetStateFunc() (fn js.Func[func() js.Any]) {
	return fn.FromRef(
		bindings.NavigationDestinationGetStateFunc(
			this.Ref(),
		),
	)
}

type NavigateEventInit struct {
	// NavigationType is "NavigateEventInit.navigationType"
	//
	// Optional, defaults to "push".
	NavigationType NavigationType
	// Destination is "NavigateEventInit.destination"
	//
	// Required
	Destination NavigationDestination
	// CanIntercept is "NavigateEventInit.canIntercept"
	//
	// Optional, defaults to false.
	CanIntercept bool
	// UserInitiated is "NavigateEventInit.userInitiated"
	//
	// Optional, defaults to false.
	UserInitiated bool
	// HashChange is "NavigateEventInit.hashChange"
	//
	// Optional, defaults to false.
	HashChange bool
	// Signal is "NavigateEventInit.signal"
	//
	// Required
	Signal AbortSignal
	// FormData is "NavigateEventInit.formData"
	//
	// Optional, defaults to null.
	FormData FormData
	// DownloadRequest is "NavigateEventInit.downloadRequest"
	//
	// Optional, defaults to null.
	DownloadRequest js.String
	// Info is "NavigateEventInit.info"
	//
	// Optional
	Info js.Any
	// HasUAVisualTransition is "NavigateEventInit.hasUAVisualTransition"
	//
	// Optional, defaults to false.
	HasUAVisualTransition bool
	// Bubbles is "NavigateEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "NavigateEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "NavigateEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_CanIntercept          bool // for CanIntercept.
	FFI_USE_UserInitiated         bool // for UserInitiated.
	FFI_USE_HashChange            bool // for HashChange.
	FFI_USE_HasUAVisualTransition bool // for HasUAVisualTransition.
	FFI_USE_Bubbles               bool // for Bubbles.
	FFI_USE_Cancelable            bool // for Cancelable.
	FFI_USE_Composed              bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NavigateEventInit with all fields set.
func (p NavigateEventInit) FromRef(ref js.Ref) NavigateEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NavigateEventInit NavigateEventInit [// NavigateEventInit] [0x1400090b900 0x1400090ba40 0x1400090bae0 0x1400090bc20 0x1400090bd60 0x1400090bea0 0x1400090bf40 0x14000934000 0x140009340a0 0x14000934140 0x14000934280 0x140009343c0 0x14000934500 0x1400090bb80 0x1400090bcc0 0x1400090be00 0x140009341e0 0x14000934320 0x14000934460 0x140009345a0] 0x14000920348 {0 0}} in the application heap.
func (p NavigateEventInit) New() js.Ref {
	return bindings.NavigateEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NavigateEventInit) UpdateFrom(ref js.Ref) {
	bindings.NavigateEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NavigateEventInit) Update(ref js.Ref) {
	bindings.NavigateEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

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

func NewMediaRecorder(stream MediaStream, options MediaRecorderOptions) (ret MediaRecorder) {
	ret.ref = bindings.NewMediaRecorderByMediaRecorder(
		stream.Ref(),
		js.Pointer(&options))
	return
}

func NewMediaRecorderByMediaRecorder1(stream MediaStream) (ret MediaRecorder) {
	ret.ref = bindings.NewMediaRecorderByMediaRecorder1(
		stream.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this MediaRecorder) Stream() (ret MediaStream, ok bool) {
	ok = js.True == bindings.GetMediaRecorderStream(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MimeType returns the value of property "MediaRecorder.mimeType".
//
// It returns ok=false if there is no such property.
func (this MediaRecorder) MimeType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaRecorderMimeType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// State returns the value of property "MediaRecorder.state".
//
// It returns ok=false if there is no such property.
func (this MediaRecorder) State() (ret RecordingState, ok bool) {
	ok = js.True == bindings.GetMediaRecorderState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// VideoBitsPerSecond returns the value of property "MediaRecorder.videoBitsPerSecond".
//
// It returns ok=false if there is no such property.
func (this MediaRecorder) VideoBitsPerSecond() (ret uint32, ok bool) {
	ok = js.True == bindings.GetMediaRecorderVideoBitsPerSecond(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AudioBitsPerSecond returns the value of property "MediaRecorder.audioBitsPerSecond".
//
// It returns ok=false if there is no such property.
func (this MediaRecorder) AudioBitsPerSecond() (ret uint32, ok bool) {
	ok = js.True == bindings.GetMediaRecorderAudioBitsPerSecond(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AudioBitrateMode returns the value of property "MediaRecorder.audioBitrateMode".
//
// It returns ok=false if there is no such property.
func (this MediaRecorder) AudioBitrateMode() (ret BitrateMode, ok bool) {
	ok = js.True == bindings.GetMediaRecorderAudioBitrateMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasStart returns true if the method "MediaRecorder.start" exists.
func (this MediaRecorder) HasStart() bool {
	return js.True == bindings.HasMediaRecorderStart(
		this.Ref(),
	)
}

// StartFunc returns the method "MediaRecorder.start".
func (this MediaRecorder) StartFunc() (fn js.Func[func(timeslice uint32)]) {
	return fn.FromRef(
		bindings.MediaRecorderStartFunc(
			this.Ref(),
		),
	)
}

// Start calls the method "MediaRecorder.start".
func (this MediaRecorder) Start(timeslice uint32) (ret js.Void) {
	bindings.CallMediaRecorderStart(
		this.Ref(), js.Pointer(&ret),
		uint32(timeslice),
	)

	return
}

// TryStart calls the method "MediaRecorder.start"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaRecorder) TryStart(timeslice uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderStart(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(timeslice),
	)

	return
}

// HasStart1 returns true if the method "MediaRecorder.start" exists.
func (this MediaRecorder) HasStart1() bool {
	return js.True == bindings.HasMediaRecorderStart1(
		this.Ref(),
	)
}

// Start1Func returns the method "MediaRecorder.start".
func (this MediaRecorder) Start1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaRecorderStart1Func(
			this.Ref(),
		),
	)
}

// Start1 calls the method "MediaRecorder.start".
func (this MediaRecorder) Start1() (ret js.Void) {
	bindings.CallMediaRecorderStart1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStart1 calls the method "MediaRecorder.start"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaRecorder) TryStart1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderStart1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStop returns true if the method "MediaRecorder.stop" exists.
func (this MediaRecorder) HasStop() bool {
	return js.True == bindings.HasMediaRecorderStop(
		this.Ref(),
	)
}

// StopFunc returns the method "MediaRecorder.stop".
func (this MediaRecorder) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaRecorderStopFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "MediaRecorder.stop".
func (this MediaRecorder) Stop() (ret js.Void) {
	bindings.CallMediaRecorderStop(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "MediaRecorder.stop"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaRecorder) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderStop(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPause returns true if the method "MediaRecorder.pause" exists.
func (this MediaRecorder) HasPause() bool {
	return js.True == bindings.HasMediaRecorderPause(
		this.Ref(),
	)
}

// PauseFunc returns the method "MediaRecorder.pause".
func (this MediaRecorder) PauseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaRecorderPauseFunc(
			this.Ref(),
		),
	)
}

// Pause calls the method "MediaRecorder.pause".
func (this MediaRecorder) Pause() (ret js.Void) {
	bindings.CallMediaRecorderPause(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryPause calls the method "MediaRecorder.pause"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaRecorder) TryPause() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderPause(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasResume returns true if the method "MediaRecorder.resume" exists.
func (this MediaRecorder) HasResume() bool {
	return js.True == bindings.HasMediaRecorderResume(
		this.Ref(),
	)
}

// ResumeFunc returns the method "MediaRecorder.resume".
func (this MediaRecorder) ResumeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaRecorderResumeFunc(
			this.Ref(),
		),
	)
}

// Resume calls the method "MediaRecorder.resume".
func (this MediaRecorder) Resume() (ret js.Void) {
	bindings.CallMediaRecorderResume(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryResume calls the method "MediaRecorder.resume"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaRecorder) TryResume() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderResume(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRequestData returns true if the method "MediaRecorder.requestData" exists.
func (this MediaRecorder) HasRequestData() bool {
	return js.True == bindings.HasMediaRecorderRequestData(
		this.Ref(),
	)
}

// RequestDataFunc returns the method "MediaRecorder.requestData".
func (this MediaRecorder) RequestDataFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MediaRecorderRequestDataFunc(
			this.Ref(),
		),
	)
}

// RequestData calls the method "MediaRecorder.requestData".
func (this MediaRecorder) RequestData() (ret js.Void) {
	bindings.CallMediaRecorderRequestData(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRequestData calls the method "MediaRecorder.requestData"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaRecorder) TryRequestData() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderRequestData(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsTypeSupported returns true if the staticmethod "MediaRecorder.isTypeSupported" exists.
func (this MediaRecorder) HasIsTypeSupported() bool {
	return js.True == bindings.HasMediaRecorderIsTypeSupported(
		this.Ref(),
	)
}

// IsTypeSupportedFunc returns the staticmethod "MediaRecorder.isTypeSupported".
func (this MediaRecorder) IsTypeSupportedFunc() (fn js.Func[func(typ js.String) bool]) {
	return fn.FromRef(
		bindings.MediaRecorderIsTypeSupportedFunc(
			this.Ref(),
		),
	)
}

// IsTypeSupported calls the staticmethod "MediaRecorder.isTypeSupported".
func (this MediaRecorder) IsTypeSupported(typ js.String) (ret bool) {
	bindings.CallMediaRecorderIsTypeSupported(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryIsTypeSupported calls the staticmethod "MediaRecorder.isTypeSupported"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MediaRecorder) TryIsTypeSupported(typ js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderIsTypeSupported(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

type MediaStreamTrackEventInit struct {
	// Track is "MediaStreamTrackEventInit.track"
	//
	// Required
	Track MediaStreamTrack
	// Bubbles is "MediaStreamTrackEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "MediaStreamTrackEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "MediaStreamTrackEventInit.composed"
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

// FromRef calls UpdateFrom and returns a MediaStreamTrackEventInit with all fields set.
func (p MediaStreamTrackEventInit) FromRef(ref js.Ref) MediaStreamTrackEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaStreamTrackEventInit in the application heap.
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

func NewMediaStreamTrackEvent(typ js.String, eventInitDict MediaStreamTrackEventInit) (ret MediaStreamTrackEvent) {
	ret.ref = bindings.NewMediaStreamTrackEventByMediaStreamTrackEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this MediaStreamTrackEvent) Track() (ret MediaStreamTrack, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackEventTrack(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type MediaStreamTrackProcessorInit struct {
	// Track is "MediaStreamTrackProcessorInit.track"
	//
	// Required
	Track MediaStreamTrack
	// MaxBufferSize is "MediaStreamTrackProcessorInit.maxBufferSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxBufferSize MUST be set to true to make this field effective.
	MaxBufferSize uint16

	FFI_USE_MaxBufferSize bool // for MaxBufferSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaStreamTrackProcessorInit with all fields set.
func (p MediaStreamTrackProcessorInit) FromRef(ref js.Ref) MediaStreamTrackProcessorInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaStreamTrackProcessorInit in the application heap.
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

func NewMediaStreamTrackProcessor(init MediaStreamTrackProcessorInit) (ret MediaStreamTrackProcessor) {
	ret.ref = bindings.NewMediaStreamTrackProcessorByMediaStreamTrackProcessor(
		js.Pointer(&init))
	return
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
// It returns ok=false if there is no such property.
func (this MediaStreamTrackProcessor) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackProcessorReadable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReadable sets the value of property "MediaStreamTrackProcessor.readable" to val.
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
	//
	// NOTE: FFI_USE_Maximum MUST be set to true to make this field effective.
	Maximum uint32

	FFI_USE_Maximum bool // for Maximum.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MemoryDescriptor with all fields set.
func (p MemoryDescriptor) FromRef(ref js.Ref) MemoryDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MemoryDescriptor in the application heap.
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

func NewMemory(descriptor MemoryDescriptor) (ret Memory) {
	ret.ref = bindings.NewMemoryByMemory(
		js.Pointer(&descriptor))
	return
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
// It returns ok=false if there is no such property.
func (this Memory) Buffer() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetMemoryBuffer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGrow returns true if the method "Memory.grow" exists.
func (this Memory) HasGrow() bool {
	return js.True == bindings.HasMemoryGrow(
		this.Ref(),
	)
}

// GrowFunc returns the method "Memory.grow".
func (this Memory) GrowFunc() (fn js.Func[func(delta uint32) uint32]) {
	return fn.FromRef(
		bindings.MemoryGrowFunc(
			this.Ref(),
		),
	)
}

// Grow calls the method "Memory.grow".
func (this Memory) Grow(delta uint32) (ret uint32) {
	bindings.CallMemoryGrow(
		this.Ref(), js.Pointer(&ret),
		uint32(delta),
	)

	return
}

// TryGrow calls the method "Memory.grow"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this Memory) TryGrow(delta uint32) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMemoryGrow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(delta),
	)

	return
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
// It returns ok=false if there is no such property.
func (this MessageChannel) Port1() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetMessageChannelPort1(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Port2 returns the value of property "MessageChannel.port2".
//
// It returns ok=false if there is no such property.
func (this MessageChannel) Port2() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetMessageChannelPort2(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "MessageEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "MessageEventInit.composed"
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

// FromRef calls UpdateFrom and returns a MessageEventInit with all fields set.
func (p MessageEventInit) FromRef(ref js.Ref) MessageEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MessageEventInit in the application heap.
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

func NewMessageEvent(typ js.String, eventInitDict MessageEventInit) (ret MessageEvent) {
	ret.ref = bindings.NewMessageEventByMessageEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewMessageEventByMessageEvent1(typ js.String) (ret MessageEvent) {
	ret.ref = bindings.NewMessageEventByMessageEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this MessageEvent) Data() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetMessageEventData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Origin returns the value of property "MessageEvent.origin".
//
// It returns ok=false if there is no such property.
func (this MessageEvent) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMessageEventOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LastEventId returns the value of property "MessageEvent.lastEventId".
//
// It returns ok=false if there is no such property.
func (this MessageEvent) LastEventId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMessageEventLastEventId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Source returns the value of property "MessageEvent.source".
//
// It returns ok=false if there is no such property.
func (this MessageEvent) Source() (ret MessageEventSource, ok bool) {
	ok = js.True == bindings.GetMessageEventSource(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ports returns the value of property "MessageEvent.ports".
//
// It returns ok=false if there is no such property.
func (this MessageEvent) Ports() (ret js.FrozenArray[MessagePort], ok bool) {
	ok = js.True == bindings.GetMessageEventPorts(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasInitMessageEvent returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasInitMessageEvent() bool {
	return js.True == bindings.HasMessageEventInitMessageEvent(
		this.Ref(),
	)
}

// InitMessageEventFunc returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEventFunc() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource, ports js.Array[MessagePort])]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEventFunc(
			this.Ref(),
		),
	)
}

// InitMessageEvent calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource, ports js.Array[MessagePort]) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
		lastEventId.Ref(),
		source.Ref(),
		ports.Ref(),
	)

	return
}

// TryInitMessageEvent calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessageEvent) TryInitMessageEvent(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource, ports js.Array[MessagePort]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
		lastEventId.Ref(),
		source.Ref(),
		ports.Ref(),
	)

	return
}

// HasInitMessageEvent1 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasInitMessageEvent1() bool {
	return js.True == bindings.HasMessageEventInitMessageEvent1(
		this.Ref(),
	)
}

// InitMessageEvent1Func returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent1Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent1Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent1 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent1(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent1(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
		lastEventId.Ref(),
		source.Ref(),
	)

	return
}

// TryInitMessageEvent1 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessageEvent) TryInitMessageEvent1(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
		lastEventId.Ref(),
		source.Ref(),
	)

	return
}

// HasInitMessageEvent2 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasInitMessageEvent2() bool {
	return js.True == bindings.HasMessageEventInitMessageEvent2(
		this.Ref(),
	)
}

// InitMessageEvent2Func returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent2Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent2Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent2 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent2(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent2(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
		lastEventId.Ref(),
	)

	return
}

// TryInitMessageEvent2 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessageEvent) TryInitMessageEvent2(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
		lastEventId.Ref(),
	)

	return
}

// HasInitMessageEvent3 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasInitMessageEvent3() bool {
	return js.True == bindings.HasMessageEventInitMessageEvent3(
		this.Ref(),
	)
}

// InitMessageEvent3Func returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent3Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent3Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent3 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent3(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent3(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
	)

	return
}

// TryInitMessageEvent3 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessageEvent) TryInitMessageEvent3(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
	)

	return
}

// HasInitMessageEvent4 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasInitMessageEvent4() bool {
	return js.True == bindings.HasMessageEventInitMessageEvent4(
		this.Ref(),
	)
}

// InitMessageEvent4Func returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent4Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent4Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent4 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent4(typ js.String, bubbles bool, cancelable bool, data js.Any) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent4(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
	)

	return
}

// TryInitMessageEvent4 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessageEvent) TryInitMessageEvent4(typ js.String, bubbles bool, cancelable bool, data js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
	)

	return
}

// HasInitMessageEvent5 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasInitMessageEvent5() bool {
	return js.True == bindings.HasMessageEventInitMessageEvent5(
		this.Ref(),
	)
}

// InitMessageEvent5Func returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent5Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent5Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent5 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent5(typ js.String, bubbles bool, cancelable bool) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent5(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// TryInitMessageEvent5 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessageEvent) TryInitMessageEvent5(typ js.String, bubbles bool, cancelable bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent5(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// HasInitMessageEvent6 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasInitMessageEvent6() bool {
	return js.True == bindings.HasMessageEventInitMessageEvent6(
		this.Ref(),
	)
}

// InitMessageEvent6Func returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent6Func() (fn js.Func[func(typ js.String, bubbles bool)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent6Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent6 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent6(typ js.String, bubbles bool) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent6(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// TryInitMessageEvent6 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessageEvent) TryInitMessageEvent6(typ js.String, bubbles bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent6(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// HasInitMessageEvent7 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasInitMessageEvent7() bool {
	return js.True == bindings.HasMessageEventInitMessageEvent7(
		this.Ref(),
	)
}

// InitMessageEvent7Func returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent7Func() (fn js.Func[func(typ js.String)]) {
	return fn.FromRef(
		bindings.MessageEventInitMessageEvent7Func(
			this.Ref(),
		),
	)
}

// InitMessageEvent7 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent7(typ js.String) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent7(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryInitMessageEvent7 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MessageEvent) TryInitMessageEvent7(typ js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent7(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
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
	//
	// NOTE: FFI_USE_Sysex MUST be set to true to make this field effective.
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

// New creates a new MidiPermissionDescriptor in the application heap.
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
	//
	// NOTE: FFI_USE_DefaultFrameRate MUST be set to true to make this field effective.
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

// New creates a new MockCameraConfiguration in the application heap.
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

// New creates a new MockCaptureDeviceConfiguration in the application heap.
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

// New creates a new MockCapturePromptResultConfiguration in the application heap.
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
	//
	// NOTE: FFI_USE_DefaultSampleRate MUST be set to true to make this field effective.
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

// New creates a new MockMicrophoneConfiguration in the application heap.
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
	//
	// NOTE: FFI_USE_MaxSamplingFrequency MUST be set to true to make this field effective.
	MaxSamplingFrequency float64
	// MinSamplingFrequency is "MockSensor.minSamplingFrequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinSamplingFrequency MUST be set to true to make this field effective.
	MinSamplingFrequency float64
	// RequestedSamplingFrequency is "MockSensor.requestedSamplingFrequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_RequestedSamplingFrequency MUST be set to true to make this field effective.
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

// New creates a new MockSensor in the application heap.
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
	//
	// NOTE: FFI_USE_Connected MUST be set to true to make this field effective.
	Connected bool
	// MaxSamplingFrequency is "MockSensorConfiguration.maxSamplingFrequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxSamplingFrequency MUST be set to true to make this field effective.
	MaxSamplingFrequency float64
	// MinSamplingFrequency is "MockSensorConfiguration.minSamplingFrequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinSamplingFrequency MUST be set to true to make this field effective.
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

// New creates a new MockSensorConfiguration in the application heap.
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

// New creates a new MockSensorReadingValues in the application heap.
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
	//
	// NOTE: FFI_USE_ScreenX MUST be set to true to make this field effective.
	ScreenX int32
	// ScreenY is "MouseEventInit.screenY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ScreenY MUST be set to true to make this field effective.
	ScreenY int32
	// ClientX is "MouseEventInit.clientX"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ClientX MUST be set to true to make this field effective.
	ClientX int32
	// ClientY is "MouseEventInit.clientY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ClientY MUST be set to true to make this field effective.
	ClientY int32
	// Button is "MouseEventInit.button"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Button MUST be set to true to make this field effective.
	Button int16
	// Buttons is "MouseEventInit.buttons"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Buttons MUST be set to true to make this field effective.
	Buttons uint16
	// RelatedTarget is "MouseEventInit.relatedTarget"
	//
	// Optional, defaults to null.
	RelatedTarget EventTarget
	// CtrlKey is "MouseEventInit.ctrlKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_CtrlKey MUST be set to true to make this field effective.
	CtrlKey bool
	// ShiftKey is "MouseEventInit.shiftKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ShiftKey MUST be set to true to make this field effective.
	ShiftKey bool
	// AltKey is "MouseEventInit.altKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AltKey MUST be set to true to make this field effective.
	AltKey bool
	// MetaKey is "MouseEventInit.metaKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_MetaKey MUST be set to true to make this field effective.
	MetaKey bool
	// ModifierAltGraph is "MouseEventInit.modifierAltGraph"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierAltGraph MUST be set to true to make this field effective.
	ModifierAltGraph bool
	// ModifierCapsLock is "MouseEventInit.modifierCapsLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierCapsLock MUST be set to true to make this field effective.
	ModifierCapsLock bool
	// ModifierFn is "MouseEventInit.modifierFn"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierFn MUST be set to true to make this field effective.
	ModifierFn bool
	// ModifierFnLock is "MouseEventInit.modifierFnLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierFnLock MUST be set to true to make this field effective.
	ModifierFnLock bool
	// ModifierHyper is "MouseEventInit.modifierHyper"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierHyper MUST be set to true to make this field effective.
	ModifierHyper bool
	// ModifierNumLock is "MouseEventInit.modifierNumLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierNumLock MUST be set to true to make this field effective.
	ModifierNumLock bool
	// ModifierScrollLock is "MouseEventInit.modifierScrollLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierScrollLock MUST be set to true to make this field effective.
	ModifierScrollLock bool
	// ModifierSuper is "MouseEventInit.modifierSuper"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSuper MUST be set to true to make this field effective.
	ModifierSuper bool
	// ModifierSymbol is "MouseEventInit.modifierSymbol"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSymbol MUST be set to true to make this field effective.
	ModifierSymbol bool
	// ModifierSymbolLock is "MouseEventInit.modifierSymbolLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSymbolLock MUST be set to true to make this field effective.
	ModifierSymbolLock bool
	// View is "MouseEventInit.view"
	//
	// Optional, defaults to null.
	View Window
	// Detail is "MouseEventInit.detail"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Detail MUST be set to true to make this field effective.
	Detail int32
	// Bubbles is "MouseEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "MouseEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "MouseEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool
	// MovementX is "MouseEventInit.movementX"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_MovementX MUST be set to true to make this field effective.
	MovementX float64
	// MovementY is "MouseEventInit.movementY"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_MovementY MUST be set to true to make this field effective.
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

// New creates a new MouseEventInit in the application heap.
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

func NewMouseEvent(typ js.String, eventInitDict MouseEventInit) (ret MouseEvent) {
	ret.ref = bindings.NewMouseEventByMouseEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewMouseEventByMouseEvent1(typ js.String) (ret MouseEvent) {
	ret.ref = bindings.NewMouseEventByMouseEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this MouseEvent) ScreenX() (ret int32, ok bool) {
	ok = js.True == bindings.GetMouseEventScreenX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ScreenY returns the value of property "MouseEvent.screenY".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) ScreenY() (ret int32, ok bool) {
	ok = js.True == bindings.GetMouseEventScreenY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ClientX returns the value of property "MouseEvent.clientX".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) ClientX() (ret int32, ok bool) {
	ok = js.True == bindings.GetMouseEventClientX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ClientY returns the value of property "MouseEvent.clientY".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) ClientY() (ret int32, ok bool) {
	ok = js.True == bindings.GetMouseEventClientY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CtrlKey returns the value of property "MouseEvent.ctrlKey".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) CtrlKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetMouseEventCtrlKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ShiftKey returns the value of property "MouseEvent.shiftKey".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) ShiftKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetMouseEventShiftKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AltKey returns the value of property "MouseEvent.altKey".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) AltKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetMouseEventAltKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MetaKey returns the value of property "MouseEvent.metaKey".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) MetaKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetMouseEventMetaKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Button returns the value of property "MouseEvent.button".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) Button() (ret int16, ok bool) {
	ok = js.True == bindings.GetMouseEventButton(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Buttons returns the value of property "MouseEvent.buttons".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) Buttons() (ret uint16, ok bool) {
	ok = js.True == bindings.GetMouseEventButtons(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RelatedTarget returns the value of property "MouseEvent.relatedTarget".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) RelatedTarget() (ret EventTarget, ok bool) {
	ok = js.True == bindings.GetMouseEventRelatedTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MovementX returns the value of property "MouseEvent.movementX".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) MovementX() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventMovementX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MovementY returns the value of property "MouseEvent.movementY".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) MovementY() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventMovementY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PageX returns the value of property "MouseEvent.pageX".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) PageX() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventPageX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PageY returns the value of property "MouseEvent.pageY".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) PageY() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventPageY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// X returns the value of property "MouseEvent.x".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "MouseEvent.y".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OffsetX returns the value of property "MouseEvent.offsetX".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) OffsetX() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventOffsetX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OffsetY returns the value of property "MouseEvent.offsetY".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) OffsetY() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventOffsetY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetModifierState returns true if the method "MouseEvent.getModifierState" exists.
func (this MouseEvent) HasGetModifierState() bool {
	return js.True == bindings.HasMouseEventGetModifierState(
		this.Ref(),
	)
}

// GetModifierStateFunc returns the method "MouseEvent.getModifierState".
func (this MouseEvent) GetModifierStateFunc() (fn js.Func[func(keyArg js.String) bool]) {
	return fn.FromRef(
		bindings.MouseEventGetModifierStateFunc(
			this.Ref(),
		),
	)
}

// GetModifierState calls the method "MouseEvent.getModifierState".
func (this MouseEvent) GetModifierState(keyArg js.String) (ret bool) {
	bindings.CallMouseEventGetModifierState(
		this.Ref(), js.Pointer(&ret),
		keyArg.Ref(),
	)

	return
}

// TryGetModifierState calls the method "MouseEvent.getModifierState"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryGetModifierState(keyArg js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventGetModifierState(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		keyArg.Ref(),
	)

	return
}

// HasInitMouseEvent returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent(
		this.Ref(),
	)
}

// InitMouseEventFunc returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEventFunc() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16, relatedTargetArg EventTarget)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEventFunc(
			this.Ref(),
		),
	)
}

// InitMouseEvent calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16, relatedTargetArg EventTarget) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryInitMouseEvent calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16, relatedTargetArg EventTarget) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasInitMouseEvent1 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent1() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent1(
		this.Ref(),
	)
}

// InitMouseEvent1Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent1Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent1Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent1 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent1(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryInitMouseEvent1 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasInitMouseEvent2 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent2() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent2(
		this.Ref(),
	)
}

// InitMouseEvent2Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent2Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent2Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent2 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent2(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryInitMouseEvent2 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasInitMouseEvent3 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent3() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent3(
		this.Ref(),
	)
}

// InitMouseEvent3Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent3Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent3Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent3 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent3(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryInitMouseEvent3 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasInitMouseEvent4 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent4() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent4(
		this.Ref(),
	)
}

// InitMouseEvent4Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent4Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent4Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent4 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent4(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryInitMouseEvent4 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasInitMouseEvent5 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent5() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent5(
		this.Ref(),
	)
}

// InitMouseEvent5Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent5Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent5Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent5 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent5(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryInitMouseEvent5 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent5(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasInitMouseEvent6 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent6() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent6(
		this.Ref(),
	)
}

// InitMouseEvent6Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent6Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent6Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent6 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent6(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent6(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryInitMouseEvent6 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent6(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent6(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasInitMouseEvent7 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent7() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent7(
		this.Ref(),
	)
}

// InitMouseEvent7Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent7Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent7Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent7 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent7(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent7(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
		int32(clientXArg),
	)

	return
}

// TryInitMouseEvent7 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent7(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent7(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
		int32(clientXArg),
	)

	return
}

// HasInitMouseEvent8 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent8() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent8(
		this.Ref(),
	)
}

// InitMouseEvent8Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent8Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent8Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent8 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent8(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent8(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
	)

	return
}

// TryInitMouseEvent8 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent8(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent8(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
		int32(screenYArg),
	)

	return
}

// HasInitMouseEvent9 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent9() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent9(
		this.Ref(),
	)
}

// InitMouseEvent9Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent9Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent9Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent9 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent9(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent9(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
	)

	return
}

// TryInitMouseEvent9 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent9(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent9(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
	)

	return
}

// HasInitMouseEvent10 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent10() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent10(
		this.Ref(),
	)
}

// InitMouseEvent10Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent10Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent10Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent10 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent10(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent10(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
	)

	return
}

// TryInitMouseEvent10 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent10(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent10(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
	)

	return
}

// HasInitMouseEvent11 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent11() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent11(
		this.Ref(),
	)
}

// InitMouseEvent11Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent11Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent11Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent11 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent11(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent11(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// TryInitMouseEvent11 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent11(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent11(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// HasInitMouseEvent12 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent12() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent12(
		this.Ref(),
	)
}

// InitMouseEvent12Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent12Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent12Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent12 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent12(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent12(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// TryInitMouseEvent12 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent12(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent12(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// HasInitMouseEvent13 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent13() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent13(
		this.Ref(),
	)
}

// InitMouseEvent13Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent13Func() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent13Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent13 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent13(typeArg js.String, bubblesArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent13(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// TryInitMouseEvent13 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent13(typeArg js.String, bubblesArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent13(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// HasInitMouseEvent14 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasInitMouseEvent14() bool {
	return js.True == bindings.HasMouseEventInitMouseEvent14(
		this.Ref(),
	)
}

// InitMouseEvent14Func returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent14Func() (fn js.Func[func(typeArg js.String)]) {
	return fn.FromRef(
		bindings.MouseEventInitMouseEvent14Func(
			this.Ref(),
		),
	)
}

// InitMouseEvent14 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent14(typeArg js.String) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent14(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
	)

	return
}

// TryInitMouseEvent14 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MouseEvent) TryInitMouseEvent14(typeArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent14(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this MutationRecord) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationRecordType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Target returns the value of property "MutationRecord.target".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) Target() (ret Node, ok bool) {
	ok = js.True == bindings.GetMutationRecordTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AddedNodes returns the value of property "MutationRecord.addedNodes".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) AddedNodes() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetMutationRecordAddedNodes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RemovedNodes returns the value of property "MutationRecord.removedNodes".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) RemovedNodes() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetMutationRecordRemovedNodes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreviousSibling returns the value of property "MutationRecord.previousSibling".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) PreviousSibling() (ret Node, ok bool) {
	ok = js.True == bindings.GetMutationRecordPreviousSibling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NextSibling returns the value of property "MutationRecord.nextSibling".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) NextSibling() (ret Node, ok bool) {
	ok = js.True == bindings.GetMutationRecordNextSibling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AttributeName returns the value of property "MutationRecord.attributeName".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) AttributeName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationRecordAttributeName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AttributeNamespace returns the value of property "MutationRecord.attributeNamespace".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) AttributeNamespace() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationRecordAttributeNamespace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// OldValue returns the value of property "MutationRecord.oldValue".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) OldValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationRecordOldValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type MutationObserverInit struct {
	// ChildList is "MutationObserverInit.childList"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ChildList MUST be set to true to make this field effective.
	ChildList bool
	// Attributes is "MutationObserverInit.attributes"
	//
	// Optional
	//
	// NOTE: FFI_USE_Attributes MUST be set to true to make this field effective.
	Attributes bool
	// CharacterData is "MutationObserverInit.characterData"
	//
	// Optional
	//
	// NOTE: FFI_USE_CharacterData MUST be set to true to make this field effective.
	CharacterData bool
	// Subtree is "MutationObserverInit.subtree"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Subtree MUST be set to true to make this field effective.
	Subtree bool
	// AttributeOldValue is "MutationObserverInit.attributeOldValue"
	//
	// Optional
	//
	// NOTE: FFI_USE_AttributeOldValue MUST be set to true to make this field effective.
	AttributeOldValue bool
	// CharacterDataOldValue is "MutationObserverInit.characterDataOldValue"
	//
	// Optional
	//
	// NOTE: FFI_USE_CharacterDataOldValue MUST be set to true to make this field effective.
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

// New creates a new MutationObserverInit in the application heap.
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

func NewMutationObserver(callback js.Func[func(mutations js.Array[MutationRecord], observer MutationObserver)]) (ret MutationObserver) {
	ret.ref = bindings.NewMutationObserverByMutationObserver(
		callback.Ref())
	return
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

// HasObserve returns true if the method "MutationObserver.observe" exists.
func (this MutationObserver) HasObserve() bool {
	return js.True == bindings.HasMutationObserverObserve(
		this.Ref(),
	)
}

// ObserveFunc returns the method "MutationObserver.observe".
func (this MutationObserver) ObserveFunc() (fn js.Func[func(target Node, options MutationObserverInit)]) {
	return fn.FromRef(
		bindings.MutationObserverObserveFunc(
			this.Ref(),
		),
	)
}

// Observe calls the method "MutationObserver.observe".
func (this MutationObserver) Observe(target Node, options MutationObserverInit) (ret js.Void) {
	bindings.CallMutationObserverObserve(
		this.Ref(), js.Pointer(&ret),
		target.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryObserve calls the method "MutationObserver.observe"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationObserver) TryObserve(target Node, options MutationObserverInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationObserverObserve(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		target.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasObserve1 returns true if the method "MutationObserver.observe" exists.
func (this MutationObserver) HasObserve1() bool {
	return js.True == bindings.HasMutationObserverObserve1(
		this.Ref(),
	)
}

// Observe1Func returns the method "MutationObserver.observe".
func (this MutationObserver) Observe1Func() (fn js.Func[func(target Node)]) {
	return fn.FromRef(
		bindings.MutationObserverObserve1Func(
			this.Ref(),
		),
	)
}

// Observe1 calls the method "MutationObserver.observe".
func (this MutationObserver) Observe1(target Node) (ret js.Void) {
	bindings.CallMutationObserverObserve1(
		this.Ref(), js.Pointer(&ret),
		target.Ref(),
	)

	return
}

// TryObserve1 calls the method "MutationObserver.observe"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationObserver) TryObserve1(target Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationObserverObserve1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		target.Ref(),
	)

	return
}

// HasDisconnect returns true if the method "MutationObserver.disconnect" exists.
func (this MutationObserver) HasDisconnect() bool {
	return js.True == bindings.HasMutationObserverDisconnect(
		this.Ref(),
	)
}

// DisconnectFunc returns the method "MutationObserver.disconnect".
func (this MutationObserver) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.MutationObserverDisconnectFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "MutationObserver.disconnect".
func (this MutationObserver) Disconnect() (ret js.Void) {
	bindings.CallMutationObserverDisconnect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDisconnect calls the method "MutationObserver.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationObserver) TryDisconnect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationObserverDisconnect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasTakeRecords returns true if the method "MutationObserver.takeRecords" exists.
func (this MutationObserver) HasTakeRecords() bool {
	return js.True == bindings.HasMutationObserverTakeRecords(
		this.Ref(),
	)
}

// TakeRecordsFunc returns the method "MutationObserver.takeRecords".
func (this MutationObserver) TakeRecordsFunc() (fn js.Func[func() js.Array[MutationRecord]]) {
	return fn.FromRef(
		bindings.MutationObserverTakeRecordsFunc(
			this.Ref(),
		),
	)
}

// TakeRecords calls the method "MutationObserver.takeRecords".
func (this MutationObserver) TakeRecords() (ret js.Array[MutationRecord]) {
	bindings.CallMutationObserverTakeRecords(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTakeRecords calls the method "MutationObserver.takeRecords"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationObserver) TryTakeRecords() (ret js.Array[MutationRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationObserverTakeRecords(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

const (
	MutationEvent_MODIFICATION uint16 = 1
	MutationEvent_ADDITION     uint16 = 2
	MutationEvent_REMOVAL      uint16 = 3
)

func NewMutationEvent(typ js.String, eventInitDict EventInit) (ret MutationEvent) {
	ret.ref = bindings.NewMutationEventByMutationEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewMutationEventByMutationEvent1(typ js.String) (ret MutationEvent) {
	ret.ref = bindings.NewMutationEventByMutationEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this MutationEvent) RelatedNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetMutationEventRelatedNode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PrevValue returns the value of property "MutationEvent.prevValue".
//
// It returns ok=false if there is no such property.
func (this MutationEvent) PrevValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationEventPrevValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NewValue returns the value of property "MutationEvent.newValue".
//
// It returns ok=false if there is no such property.
func (this MutationEvent) NewValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationEventNewValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AttrName returns the value of property "MutationEvent.attrName".
//
// It returns ok=false if there is no such property.
func (this MutationEvent) AttrName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationEventAttrName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AttrChange returns the value of property "MutationEvent.attrChange".
//
// It returns ok=false if there is no such property.
func (this MutationEvent) AttrChange() (ret uint16, ok bool) {
	ok = js.True == bindings.GetMutationEventAttrChange(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasInitMutationEvent returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasInitMutationEvent() bool {
	return js.True == bindings.HasMutationEventInitMutationEvent(
		this.Ref(),
	)
}

// InitMutationEventFunc returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEventFunc() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String, attrChangeArg uint16)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEventFunc(
			this.Ref(),
		),
	)
}

// InitMutationEvent calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String, attrChangeArg uint16) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
		newValueArg.Ref(),
		attrNameArg.Ref(),
		uint32(attrChangeArg),
	)

	return
}

// TryInitMutationEvent calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationEvent) TryInitMutationEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String, attrChangeArg uint16) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
		newValueArg.Ref(),
		attrNameArg.Ref(),
		uint32(attrChangeArg),
	)

	return
}

// HasInitMutationEvent1 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasInitMutationEvent1() bool {
	return js.True == bindings.HasMutationEventInitMutationEvent1(
		this.Ref(),
	)
}

// InitMutationEvent1Func returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent1Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent1Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent1 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent1(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
		newValueArg.Ref(),
		attrNameArg.Ref(),
	)

	return
}

// TryInitMutationEvent1 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationEvent) TryInitMutationEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
		newValueArg.Ref(),
		attrNameArg.Ref(),
	)

	return
}

// HasInitMutationEvent2 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasInitMutationEvent2() bool {
	return js.True == bindings.HasMutationEventInitMutationEvent2(
		this.Ref(),
	)
}

// InitMutationEvent2Func returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent2Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent2Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent2 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent2(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
		newValueArg.Ref(),
	)

	return
}

// TryInitMutationEvent2 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationEvent) TryInitMutationEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
		newValueArg.Ref(),
	)

	return
}

// HasInitMutationEvent3 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasInitMutationEvent3() bool {
	return js.True == bindings.HasMutationEventInitMutationEvent3(
		this.Ref(),
	)
}

// InitMutationEvent3Func returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent3Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent3Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent3 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent3(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
	)

	return
}

// TryInitMutationEvent3 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationEvent) TryInitMutationEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
	)

	return
}

// HasInitMutationEvent4 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasInitMutationEvent4() bool {
	return js.True == bindings.HasMutationEventInitMutationEvent4(
		this.Ref(),
	)
}

// InitMutationEvent4Func returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent4Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent4Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent4 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent4(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
	)

	return
}

// TryInitMutationEvent4 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationEvent) TryInitMutationEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
	)

	return
}

// HasInitMutationEvent5 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasInitMutationEvent5() bool {
	return js.True == bindings.HasMutationEventInitMutationEvent5(
		this.Ref(),
	)
}

// InitMutationEvent5Func returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent5Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent5Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent5 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent5(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// TryInitMutationEvent5 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationEvent) TryInitMutationEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent5(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// HasInitMutationEvent6 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasInitMutationEvent6() bool {
	return js.True == bindings.HasMutationEventInitMutationEvent6(
		this.Ref(),
	)
}

// InitMutationEvent6Func returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent6Func() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent6Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent6 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent6(typeArg js.String, bubblesArg bool) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent6(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// TryInitMutationEvent6 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationEvent) TryInitMutationEvent6(typeArg js.String, bubblesArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent6(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// HasInitMutationEvent7 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasInitMutationEvent7() bool {
	return js.True == bindings.HasMutationEventInitMutationEvent7(
		this.Ref(),
	)
}

// InitMutationEvent7Func returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent7Func() (fn js.Func[func(typeArg js.String)]) {
	return fn.FromRef(
		bindings.MutationEventInitMutationEvent7Func(
			this.Ref(),
		),
	)
}

// InitMutationEvent7 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent7(typeArg js.String) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent7(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
	)

	return
}

// TryInitMutationEvent7 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this MutationEvent) TryInitMutationEvent7(typeArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent7(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
	)

	return
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

// New creates a new NDEFMakeReadOnlyOptions in the application heap.
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

// New creates a new NDEFRecordInit in the application heap.
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

// New creates a new NDEFMessageInit in the application heap.
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

func NewNDEFRecord(recordInit NDEFRecordInit) (ret NDEFRecord) {
	ret.ref = bindings.NewNDEFRecordByNDEFRecord(
		js.Pointer(&recordInit))
	return
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
// It returns ok=false if there is no such property.
func (this NDEFRecord) RecordType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFRecordRecordType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MediaType returns the value of property "NDEFRecord.mediaType".
//
// It returns ok=false if there is no such property.
func (this NDEFRecord) MediaType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFRecordMediaType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "NDEFRecord.id".
//
// It returns ok=false if there is no such property.
func (this NDEFRecord) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFRecordId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Data returns the value of property "NDEFRecord.data".
//
// It returns ok=false if there is no such property.
func (this NDEFRecord) Data() (ret js.DataView, ok bool) {
	ok = js.True == bindings.GetNDEFRecordData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Encoding returns the value of property "NDEFRecord.encoding".
//
// It returns ok=false if there is no such property.
func (this NDEFRecord) Encoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFRecordEncoding(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Lang returns the value of property "NDEFRecord.lang".
//
// It returns ok=false if there is no such property.
func (this NDEFRecord) Lang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFRecordLang(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToRecords returns true if the method "NDEFRecord.toRecords" exists.
func (this NDEFRecord) HasToRecords() bool {
	return js.True == bindings.HasNDEFRecordToRecords(
		this.Ref(),
	)
}

// ToRecordsFunc returns the method "NDEFRecord.toRecords".
func (this NDEFRecord) ToRecordsFunc() (fn js.Func[func() js.Array[NDEFRecord]]) {
	return fn.FromRef(
		bindings.NDEFRecordToRecordsFunc(
			this.Ref(),
		),
	)
}

// ToRecords calls the method "NDEFRecord.toRecords".
func (this NDEFRecord) ToRecords() (ret js.Array[NDEFRecord]) {
	bindings.CallNDEFRecordToRecords(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToRecords calls the method "NDEFRecord.toRecords"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NDEFRecord) TryToRecords() (ret js.Array[NDEFRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFRecordToRecords(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewNDEFMessage(messageInit NDEFMessageInit) (ret NDEFMessage) {
	ret.ref = bindings.NewNDEFMessageByNDEFMessage(
		js.Pointer(&messageInit))
	return
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
// It returns ok=false if there is no such property.
func (this NDEFMessage) Records() (ret js.FrozenArray[NDEFRecord], ok bool) {
	ok = js.True == bindings.GetNDEFMessageRecords(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// New creates a new NDEFScanOptions in the application heap.
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
	//
	// NOTE: FFI_USE_Overwrite MUST be set to true to make this field effective.
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

// New creates a new NDEFWriteOptions in the application heap.
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

// HasScan returns true if the method "NDEFReader.scan" exists.
func (this NDEFReader) HasScan() bool {
	return js.True == bindings.HasNDEFReaderScan(
		this.Ref(),
	)
}

// ScanFunc returns the method "NDEFReader.scan".
func (this NDEFReader) ScanFunc() (fn js.Func[func(options NDEFScanOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderScanFunc(
			this.Ref(),
		),
	)
}

// Scan calls the method "NDEFReader.scan".
func (this NDEFReader) Scan(options NDEFScanOptions) (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderScan(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScan calls the method "NDEFReader.scan"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NDEFReader) TryScan(options NDEFScanOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderScan(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasScan1 returns true if the method "NDEFReader.scan" exists.
func (this NDEFReader) HasScan1() bool {
	return js.True == bindings.HasNDEFReaderScan1(
		this.Ref(),
	)
}

// Scan1Func returns the method "NDEFReader.scan".
func (this NDEFReader) Scan1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderScan1Func(
			this.Ref(),
		),
	)
}

// Scan1 calls the method "NDEFReader.scan".
func (this NDEFReader) Scan1() (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderScan1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScan1 calls the method "NDEFReader.scan"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NDEFReader) TryScan1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderScan1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasWrite returns true if the method "NDEFReader.write" exists.
func (this NDEFReader) HasWrite() bool {
	return js.True == bindings.HasNDEFReaderWrite(
		this.Ref(),
	)
}

// WriteFunc returns the method "NDEFReader.write".
func (this NDEFReader) WriteFunc() (fn js.Func[func(message NDEFMessageSource, options NDEFWriteOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderWriteFunc(
			this.Ref(),
		),
	)
}

// Write calls the method "NDEFReader.write".
func (this NDEFReader) Write(message NDEFMessageSource, options NDEFWriteOptions) (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderWrite(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryWrite calls the method "NDEFReader.write"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NDEFReader) TryWrite(message NDEFMessageSource, options NDEFWriteOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderWrite(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasWrite1 returns true if the method "NDEFReader.write" exists.
func (this NDEFReader) HasWrite1() bool {
	return js.True == bindings.HasNDEFReaderWrite1(
		this.Ref(),
	)
}

// Write1Func returns the method "NDEFReader.write".
func (this NDEFReader) Write1Func() (fn js.Func[func(message NDEFMessageSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderWrite1Func(
			this.Ref(),
		),
	)
}

// Write1 calls the method "NDEFReader.write".
func (this NDEFReader) Write1(message NDEFMessageSource) (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderWrite1(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryWrite1 calls the method "NDEFReader.write"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NDEFReader) TryWrite1(message NDEFMessageSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderWrite1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasMakeReadOnly returns true if the method "NDEFReader.makeReadOnly" exists.
func (this NDEFReader) HasMakeReadOnly() bool {
	return js.True == bindings.HasNDEFReaderMakeReadOnly(
		this.Ref(),
	)
}

// MakeReadOnlyFunc returns the method "NDEFReader.makeReadOnly".
func (this NDEFReader) MakeReadOnlyFunc() (fn js.Func[func(options NDEFMakeReadOnlyOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderMakeReadOnlyFunc(
			this.Ref(),
		),
	)
}

// MakeReadOnly calls the method "NDEFReader.makeReadOnly".
func (this NDEFReader) MakeReadOnly(options NDEFMakeReadOnlyOptions) (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderMakeReadOnly(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryMakeReadOnly calls the method "NDEFReader.makeReadOnly"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NDEFReader) TryMakeReadOnly(options NDEFMakeReadOnlyOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderMakeReadOnly(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasMakeReadOnly1 returns true if the method "NDEFReader.makeReadOnly" exists.
func (this NDEFReader) HasMakeReadOnly1() bool {
	return js.True == bindings.HasNDEFReaderMakeReadOnly1(
		this.Ref(),
	)
}

// MakeReadOnly1Func returns the method "NDEFReader.makeReadOnly".
func (this NDEFReader) MakeReadOnly1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.NDEFReaderMakeReadOnly1Func(
			this.Ref(),
		),
	)
}

// MakeReadOnly1 calls the method "NDEFReader.makeReadOnly".
func (this NDEFReader) MakeReadOnly1() (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderMakeReadOnly1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryMakeReadOnly1 calls the method "NDEFReader.makeReadOnly"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NDEFReader) TryMakeReadOnly1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderMakeReadOnly1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type NDEFReadingEventInit struct {
	// SerialNumber is "NDEFReadingEventInit.serialNumber"
	//
	// Optional, defaults to "".
	SerialNumber js.String
	// Message is "NDEFReadingEventInit.message"
	//
	// Required
	//
	// NOTE: Message.FFI_USE MUST be set to true to get Message used.
	Message NDEFMessageInit
	// Bubbles is "NDEFReadingEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "NDEFReadingEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "NDEFReadingEventInit.composed"
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

// FromRef calls UpdateFrom and returns a NDEFReadingEventInit with all fields set.
func (p NDEFReadingEventInit) FromRef(ref js.Ref) NDEFReadingEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NDEFReadingEventInit in the application heap.
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

func NewNDEFReadingEvent(typ js.String, readingEventInitDict NDEFReadingEventInit) (ret NDEFReadingEvent) {
	ret.ref = bindings.NewNDEFReadingEventByNDEFReadingEvent(
		typ.Ref(),
		js.Pointer(&readingEventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this NDEFReadingEvent) SerialNumber() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFReadingEventSerialNumber(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "NDEFReadingEvent.message".
//
// It returns ok=false if there is no such property.
func (this NDEFReadingEvent) Message() (ret NDEFMessage, ok bool) {
	ok = js.True == bindings.GetNDEFReadingEventMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this NamedFlow) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNamedFlowName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Overset returns the value of property "NamedFlow.overset".
//
// It returns ok=false if there is no such property.
func (this NamedFlow) Overset() (ret bool, ok bool) {
	ok = js.True == bindings.GetNamedFlowOverset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FirstEmptyRegionIndex returns the value of property "NamedFlow.firstEmptyRegionIndex".
//
// It returns ok=false if there is no such property.
func (this NamedFlow) FirstEmptyRegionIndex() (ret int16, ok bool) {
	ok = js.True == bindings.GetNamedFlowFirstEmptyRegionIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetRegions returns true if the method "NamedFlow.getRegions" exists.
func (this NamedFlow) HasGetRegions() bool {
	return js.True == bindings.HasNamedFlowGetRegions(
		this.Ref(),
	)
}

// GetRegionsFunc returns the method "NamedFlow.getRegions".
func (this NamedFlow) GetRegionsFunc() (fn js.Func[func() js.Array[Element]]) {
	return fn.FromRef(
		bindings.NamedFlowGetRegionsFunc(
			this.Ref(),
		),
	)
}

// GetRegions calls the method "NamedFlow.getRegions".
func (this NamedFlow) GetRegions() (ret js.Array[Element]) {
	bindings.CallNamedFlowGetRegions(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetRegions calls the method "NamedFlow.getRegions"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NamedFlow) TryGetRegions() (ret js.Array[Element], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedFlowGetRegions(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetContent returns true if the method "NamedFlow.getContent" exists.
func (this NamedFlow) HasGetContent() bool {
	return js.True == bindings.HasNamedFlowGetContent(
		this.Ref(),
	)
}

// GetContentFunc returns the method "NamedFlow.getContent".
func (this NamedFlow) GetContentFunc() (fn js.Func[func() js.Array[Node]]) {
	return fn.FromRef(
		bindings.NamedFlowGetContentFunc(
			this.Ref(),
		),
	)
}

// GetContent calls the method "NamedFlow.getContent".
func (this NamedFlow) GetContent() (ret js.Array[Node]) {
	bindings.CallNamedFlowGetContent(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetContent calls the method "NamedFlow.getContent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NamedFlow) TryGetContent() (ret js.Array[Node], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedFlowGetContent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetRegionsByContent returns true if the method "NamedFlow.getRegionsByContent" exists.
func (this NamedFlow) HasGetRegionsByContent() bool {
	return js.True == bindings.HasNamedFlowGetRegionsByContent(
		this.Ref(),
	)
}

// GetRegionsByContentFunc returns the method "NamedFlow.getRegionsByContent".
func (this NamedFlow) GetRegionsByContentFunc() (fn js.Func[func(node Node) js.Array[Element]]) {
	return fn.FromRef(
		bindings.NamedFlowGetRegionsByContentFunc(
			this.Ref(),
		),
	)
}

// GetRegionsByContent calls the method "NamedFlow.getRegionsByContent".
func (this NamedFlow) GetRegionsByContent(node Node) (ret js.Array[Element]) {
	bindings.CallNamedFlowGetRegionsByContent(
		this.Ref(), js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryGetRegionsByContent calls the method "NamedFlow.getRegionsByContent"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NamedFlow) TryGetRegionsByContent(node Node) (ret js.Array[Element], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedFlowGetRegionsByContent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this NavigationDestination) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationDestinationUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Key returns the value of property "NavigationDestination.key".
//
// It returns ok=false if there is no such property.
func (this NavigationDestination) Key() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationDestinationKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "NavigationDestination.id".
//
// It returns ok=false if there is no such property.
func (this NavigationDestination) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationDestinationId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Index returns the value of property "NavigationDestination.index".
//
// It returns ok=false if there is no such property.
func (this NavigationDestination) Index() (ret int64, ok bool) {
	ok = js.True == bindings.GetNavigationDestinationIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SameDocument returns the value of property "NavigationDestination.sameDocument".
//
// It returns ok=false if there is no such property.
func (this NavigationDestination) SameDocument() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigationDestinationSameDocument(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetState returns true if the method "NavigationDestination.getState" exists.
func (this NavigationDestination) HasGetState() bool {
	return js.True == bindings.HasNavigationDestinationGetState(
		this.Ref(),
	)
}

// GetStateFunc returns the method "NavigationDestination.getState".
func (this NavigationDestination) GetStateFunc() (fn js.Func[func() js.Any]) {
	return fn.FromRef(
		bindings.NavigationDestinationGetStateFunc(
			this.Ref(),
		),
	)
}

// GetState calls the method "NavigationDestination.getState".
func (this NavigationDestination) GetState() (ret js.Any) {
	bindings.CallNavigationDestinationGetState(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetState calls the method "NavigationDestination.getState"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NavigationDestination) TryGetState() (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationDestinationGetState(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
	//
	// NOTE: FFI_USE_CanIntercept MUST be set to true to make this field effective.
	CanIntercept bool
	// UserInitiated is "NavigateEventInit.userInitiated"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_UserInitiated MUST be set to true to make this field effective.
	UserInitiated bool
	// HashChange is "NavigateEventInit.hashChange"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_HashChange MUST be set to true to make this field effective.
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
	//
	// NOTE: FFI_USE_HasUAVisualTransition MUST be set to true to make this field effective.
	HasUAVisualTransition bool
	// Bubbles is "NavigateEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "NavigateEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "NavigateEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new NavigateEventInit in the application heap.
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

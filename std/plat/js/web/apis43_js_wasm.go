// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
	this.ref.Once()
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
	this.ref.Free()
}

// Stream returns the value of property "MediaRecorder.stream".
//
// It returns ok=false if there is no such property.
func (this MediaRecorder) Stream() (ret MediaStream, ok bool) {
	ok = js.True == bindings.GetMediaRecorderStream(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MimeType returns the value of property "MediaRecorder.mimeType".
//
// It returns ok=false if there is no such property.
func (this MediaRecorder) MimeType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaRecorderMimeType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// State returns the value of property "MediaRecorder.state".
//
// It returns ok=false if there is no such property.
func (this MediaRecorder) State() (ret RecordingState, ok bool) {
	ok = js.True == bindings.GetMediaRecorderState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VideoBitsPerSecond returns the value of property "MediaRecorder.videoBitsPerSecond".
//
// It returns ok=false if there is no such property.
func (this MediaRecorder) VideoBitsPerSecond() (ret uint32, ok bool) {
	ok = js.True == bindings.GetMediaRecorderVideoBitsPerSecond(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AudioBitsPerSecond returns the value of property "MediaRecorder.audioBitsPerSecond".
//
// It returns ok=false if there is no such property.
func (this MediaRecorder) AudioBitsPerSecond() (ret uint32, ok bool) {
	ok = js.True == bindings.GetMediaRecorderAudioBitsPerSecond(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AudioBitrateMode returns the value of property "MediaRecorder.audioBitrateMode".
//
// It returns ok=false if there is no such property.
func (this MediaRecorder) AudioBitrateMode() (ret BitrateMode, ok bool) {
	ok = js.True == bindings.GetMediaRecorderAudioBitrateMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncStart returns true if the method "MediaRecorder.start" exists.
func (this MediaRecorder) HasFuncStart() bool {
	return js.True == bindings.HasFuncMediaRecorderStart(
		this.ref,
	)
}

// FuncStart returns the method "MediaRecorder.start".
func (this MediaRecorder) FuncStart() (fn js.Func[func(timeslice uint32)]) {
	bindings.FuncMediaRecorderStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start calls the method "MediaRecorder.start".
func (this MediaRecorder) Start(timeslice uint32) (ret js.Void) {
	bindings.CallMediaRecorderStart(
		this.ref, js.Pointer(&ret),
		uint32(timeslice),
	)

	return
}

// TryStart calls the method "MediaRecorder.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaRecorder) TryStart(timeslice uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(timeslice),
	)

	return
}

// HasFuncStart1 returns true if the method "MediaRecorder.start" exists.
func (this MediaRecorder) HasFuncStart1() bool {
	return js.True == bindings.HasFuncMediaRecorderStart1(
		this.ref,
	)
}

// FuncStart1 returns the method "MediaRecorder.start".
func (this MediaRecorder) FuncStart1() (fn js.Func[func()]) {
	bindings.FuncMediaRecorderStart1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start1 calls the method "MediaRecorder.start".
func (this MediaRecorder) Start1() (ret js.Void) {
	bindings.CallMediaRecorderStart1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStart1 calls the method "MediaRecorder.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaRecorder) TryStart1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderStart1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStop returns true if the method "MediaRecorder.stop" exists.
func (this MediaRecorder) HasFuncStop() bool {
	return js.True == bindings.HasFuncMediaRecorderStop(
		this.ref,
	)
}

// FuncStop returns the method "MediaRecorder.stop".
func (this MediaRecorder) FuncStop() (fn js.Func[func()]) {
	bindings.FuncMediaRecorderStop(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stop calls the method "MediaRecorder.stop".
func (this MediaRecorder) Stop() (ret js.Void) {
	bindings.CallMediaRecorderStop(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "MediaRecorder.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaRecorder) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderStop(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPause returns true if the method "MediaRecorder.pause" exists.
func (this MediaRecorder) HasFuncPause() bool {
	return js.True == bindings.HasFuncMediaRecorderPause(
		this.ref,
	)
}

// FuncPause returns the method "MediaRecorder.pause".
func (this MediaRecorder) FuncPause() (fn js.Func[func()]) {
	bindings.FuncMediaRecorderPause(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Pause calls the method "MediaRecorder.pause".
func (this MediaRecorder) Pause() (ret js.Void) {
	bindings.CallMediaRecorderPause(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPause calls the method "MediaRecorder.pause"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaRecorder) TryPause() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderPause(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncResume returns true if the method "MediaRecorder.resume" exists.
func (this MediaRecorder) HasFuncResume() bool {
	return js.True == bindings.HasFuncMediaRecorderResume(
		this.ref,
	)
}

// FuncResume returns the method "MediaRecorder.resume".
func (this MediaRecorder) FuncResume() (fn js.Func[func()]) {
	bindings.FuncMediaRecorderResume(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Resume calls the method "MediaRecorder.resume".
func (this MediaRecorder) Resume() (ret js.Void) {
	bindings.CallMediaRecorderResume(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryResume calls the method "MediaRecorder.resume"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaRecorder) TryResume() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderResume(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRequestData returns true if the method "MediaRecorder.requestData" exists.
func (this MediaRecorder) HasFuncRequestData() bool {
	return js.True == bindings.HasFuncMediaRecorderRequestData(
		this.ref,
	)
}

// FuncRequestData returns the method "MediaRecorder.requestData".
func (this MediaRecorder) FuncRequestData() (fn js.Func[func()]) {
	bindings.FuncMediaRecorderRequestData(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RequestData calls the method "MediaRecorder.requestData".
func (this MediaRecorder) RequestData() (ret js.Void) {
	bindings.CallMediaRecorderRequestData(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRequestData calls the method "MediaRecorder.requestData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaRecorder) TryRequestData() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderRequestData(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsTypeSupported returns true if the static method "MediaRecorder.isTypeSupported" exists.
func (this MediaRecorder) HasFuncIsTypeSupported() bool {
	return js.True == bindings.HasFuncMediaRecorderIsTypeSupported(
		this.ref,
	)
}

// FuncIsTypeSupported returns the static method "MediaRecorder.isTypeSupported".
func (this MediaRecorder) FuncIsTypeSupported() (fn js.Func[func(typ js.String) bool]) {
	bindings.FuncMediaRecorderIsTypeSupported(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsTypeSupported calls the static method "MediaRecorder.isTypeSupported".
func (this MediaRecorder) IsTypeSupported(typ js.String) (ret bool) {
	bindings.CallMediaRecorderIsTypeSupported(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryIsTypeSupported calls the static method "MediaRecorder.isTypeSupported"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaRecorder) TryIsTypeSupported(typ js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaRecorderIsTypeSupported(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *MediaStreamTrackEventInit) UpdateFrom(ref js.Ref) {
	bindings.MediaStreamTrackEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaStreamTrackEventInit) Update(ref js.Ref) {
	bindings.MediaStreamTrackEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaStreamTrackEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Track.Ref(),
	)
	p.Track = p.Track.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Track returns the value of property "MediaStreamTrackEvent.track".
//
// It returns ok=false if there is no such property.
func (this MediaStreamTrackEvent) Track() (ret MediaStreamTrack, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackEventTrack(
		this.ref, js.Pointer(&ret),
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
func (p *MediaStreamTrackProcessorInit) UpdateFrom(ref js.Ref) {
	bindings.MediaStreamTrackProcessorInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaStreamTrackProcessorInit) Update(ref js.Ref) {
	bindings.MediaStreamTrackProcessorInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaStreamTrackProcessorInit) FreeMembers(recursive bool) {
	js.Free(
		p.Track.Ref(),
	)
	p.Track = p.Track.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Readable returns the value of property "MediaStreamTrackProcessor.readable".
//
// It returns ok=false if there is no such property.
func (this MediaStreamTrackProcessor) Readable() (ret ReadableStream, ok bool) {
	ok = js.True == bindings.GetMediaStreamTrackProcessorReadable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReadable sets the value of property "MediaStreamTrackProcessor.readable" to val.
//
// It returns false if the property cannot be set.
func (this MediaStreamTrackProcessor) SetReadable(val ReadableStream) bool {
	return js.True == bindings.SetMediaStreamTrackProcessorReadable(
		this.ref,
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
func (p *MemoryDescriptor) UpdateFrom(ref js.Ref) {
	bindings.MemoryDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MemoryDescriptor) Update(ref js.Ref) {
	bindings.MemoryDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MemoryDescriptor) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Buffer returns the value of property "Memory.buffer".
//
// It returns ok=false if there is no such property.
func (this Memory) Buffer() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetMemoryBuffer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGrow returns true if the method "Memory.grow" exists.
func (this Memory) HasFuncGrow() bool {
	return js.True == bindings.HasFuncMemoryGrow(
		this.ref,
	)
}

// FuncGrow returns the method "Memory.grow".
func (this Memory) FuncGrow() (fn js.Func[func(delta uint32) uint32]) {
	bindings.FuncMemoryGrow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Grow calls the method "Memory.grow".
func (this Memory) Grow(delta uint32) (ret uint32) {
	bindings.CallMemoryGrow(
		this.ref, js.Pointer(&ret),
		uint32(delta),
	)

	return
}

// TryGrow calls the method "Memory.grow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Memory) TryGrow(delta uint32) (ret uint32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMemoryGrow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(delta),
	)

	return
}

type MessageChannel struct {
	ref js.Ref
}

func (this MessageChannel) Once() MessageChannel {
	this.ref.Once()
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
	this.ref.Free()
}

// Port1 returns the value of property "MessageChannel.port1".
//
// It returns ok=false if there is no such property.
func (this MessageChannel) Port1() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetMessageChannelPort1(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Port2 returns the value of property "MessageChannel.port2".
//
// It returns ok=false if there is no such property.
func (this MessageChannel) Port2() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetMessageChannelPort2(
		this.ref, js.Pointer(&ret),
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
func (p *MessageEventInit) UpdateFrom(ref js.Ref) {
	bindings.MessageEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MessageEventInit) Update(ref js.Ref) {
	bindings.MessageEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MessageEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
		p.Origin.Ref(),
		p.LastEventId.Ref(),
		p.Source.Ref(),
		p.Ports.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
	p.Origin = p.Origin.FromRef(js.Undefined)
	p.LastEventId = p.LastEventId.FromRef(js.Undefined)
	p.Source = p.Source.FromRef(js.Undefined)
	p.Ports = p.Ports.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Data returns the value of property "MessageEvent.data".
//
// It returns ok=false if there is no such property.
func (this MessageEvent) Data() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetMessageEventData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Origin returns the value of property "MessageEvent.origin".
//
// It returns ok=false if there is no such property.
func (this MessageEvent) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMessageEventOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LastEventId returns the value of property "MessageEvent.lastEventId".
//
// It returns ok=false if there is no such property.
func (this MessageEvent) LastEventId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMessageEventLastEventId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Source returns the value of property "MessageEvent.source".
//
// It returns ok=false if there is no such property.
func (this MessageEvent) Source() (ret MessageEventSource, ok bool) {
	ok = js.True == bindings.GetMessageEventSource(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Ports returns the value of property "MessageEvent.ports".
//
// It returns ok=false if there is no such property.
func (this MessageEvent) Ports() (ret js.FrozenArray[MessagePort], ok bool) {
	ok = js.True == bindings.GetMessageEventPorts(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncInitMessageEvent returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasFuncInitMessageEvent() bool {
	return js.True == bindings.HasFuncMessageEventInitMessageEvent(
		this.ref,
	)
}

// FuncInitMessageEvent returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) FuncInitMessageEvent() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource, ports js.Array[MessagePort])]) {
	bindings.FuncMessageEventInitMessageEvent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMessageEvent calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource, ports js.Array[MessagePort]) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessageEvent) TryInitMessageEvent(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource, ports js.Array[MessagePort]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMessageEvent1 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasFuncInitMessageEvent1() bool {
	return js.True == bindings.HasFuncMessageEventInitMessageEvent1(
		this.ref,
	)
}

// FuncInitMessageEvent1 returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) FuncInitMessageEvent1() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource)]) {
	bindings.FuncMessageEventInitMessageEvent1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMessageEvent1 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent1(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent1(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessageEvent) TryInitMessageEvent1(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String, source MessageEventSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMessageEvent2 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasFuncInitMessageEvent2() bool {
	return js.True == bindings.HasFuncMessageEventInitMessageEvent2(
		this.ref,
	)
}

// FuncInitMessageEvent2 returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) FuncInitMessageEvent2() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String)]) {
	bindings.FuncMessageEventInitMessageEvent2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMessageEvent2 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent2(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent2(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessageEvent) TryInitMessageEvent2(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String, lastEventId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
		lastEventId.Ref(),
	)

	return
}

// HasFuncInitMessageEvent3 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasFuncInitMessageEvent3() bool {
	return js.True == bindings.HasFuncMessageEventInitMessageEvent3(
		this.ref,
	)
}

// FuncInitMessageEvent3 returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) FuncInitMessageEvent3() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String)]) {
	bindings.FuncMessageEventInitMessageEvent3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMessageEvent3 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent3(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent3(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
	)

	return
}

// TryInitMessageEvent3 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessageEvent) TryInitMessageEvent3(typ js.String, bubbles bool, cancelable bool, data js.Any, origin js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
		origin.Ref(),
	)

	return
}

// HasFuncInitMessageEvent4 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasFuncInitMessageEvent4() bool {
	return js.True == bindings.HasFuncMessageEventInitMessageEvent4(
		this.ref,
	)
}

// FuncInitMessageEvent4 returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) FuncInitMessageEvent4() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, data js.Any)]) {
	bindings.FuncMessageEventInitMessageEvent4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMessageEvent4 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent4(typ js.String, bubbles bool, cancelable bool, data js.Any) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent4(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
	)

	return
}

// TryInitMessageEvent4 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessageEvent) TryInitMessageEvent4(typ js.String, bubbles bool, cancelable bool, data js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		data.Ref(),
	)

	return
}

// HasFuncInitMessageEvent5 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasFuncInitMessageEvent5() bool {
	return js.True == bindings.HasFuncMessageEventInitMessageEvent5(
		this.ref,
	)
}

// FuncInitMessageEvent5 returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) FuncInitMessageEvent5() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool)]) {
	bindings.FuncMessageEventInitMessageEvent5(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMessageEvent5 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent5(typ js.String, bubbles bool, cancelable bool) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent5(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// TryInitMessageEvent5 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessageEvent) TryInitMessageEvent5(typ js.String, bubbles bool, cancelable bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent5(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// HasFuncInitMessageEvent6 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasFuncInitMessageEvent6() bool {
	return js.True == bindings.HasFuncMessageEventInitMessageEvent6(
		this.ref,
	)
}

// FuncInitMessageEvent6 returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) FuncInitMessageEvent6() (fn js.Func[func(typ js.String, bubbles bool)]) {
	bindings.FuncMessageEventInitMessageEvent6(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMessageEvent6 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent6(typ js.String, bubbles bool) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent6(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// TryInitMessageEvent6 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessageEvent) TryInitMessageEvent6(typ js.String, bubbles bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent6(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// HasFuncInitMessageEvent7 returns true if the method "MessageEvent.initMessageEvent" exists.
func (this MessageEvent) HasFuncInitMessageEvent7() bool {
	return js.True == bindings.HasFuncMessageEventInitMessageEvent7(
		this.ref,
	)
}

// FuncInitMessageEvent7 returns the method "MessageEvent.initMessageEvent".
func (this MessageEvent) FuncInitMessageEvent7() (fn js.Func[func(typ js.String)]) {
	bindings.FuncMessageEventInitMessageEvent7(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMessageEvent7 calls the method "MessageEvent.initMessageEvent".
func (this MessageEvent) InitMessageEvent7(typ js.String) (ret js.Void) {
	bindings.CallMessageEventInitMessageEvent7(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryInitMessageEvent7 calls the method "MessageEvent.initMessageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MessageEvent) TryInitMessageEvent7(typ js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMessageEventInitMessageEvent7(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *MidiPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.MidiPermissionDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MidiPermissionDescriptor) Update(ref js.Ref) {
	bindings.MidiPermissionDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MidiPermissionDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
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
func (p *MockCameraConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MockCameraConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MockCameraConfiguration) Update(ref js.Ref) {
	bindings.MockCameraConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MockCameraConfiguration) FreeMembers(recursive bool) {
	js.Free(
		p.FacingMode.Ref(),
		p.Label.Ref(),
		p.DeviceId.Ref(),
		p.GroupId.Ref(),
	)
	p.FacingMode = p.FacingMode.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
	p.DeviceId = p.DeviceId.FromRef(js.Undefined)
	p.GroupId = p.GroupId.FromRef(js.Undefined)
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
func (p *MockCaptureDeviceConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MockCaptureDeviceConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MockCaptureDeviceConfiguration) Update(ref js.Ref) {
	bindings.MockCaptureDeviceConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MockCaptureDeviceConfiguration) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
		p.DeviceId.Ref(),
		p.GroupId.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
	p.DeviceId = p.DeviceId.FromRef(js.Undefined)
	p.GroupId = p.GroupId.FromRef(js.Undefined)
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
func (p *MockCapturePromptResultConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MockCapturePromptResultConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MockCapturePromptResultConfiguration) Update(ref js.Ref) {
	bindings.MockCapturePromptResultConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MockCapturePromptResultConfiguration) FreeMembers(recursive bool) {
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
func (p *MockMicrophoneConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MockMicrophoneConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MockMicrophoneConfiguration) Update(ref js.Ref) {
	bindings.MockMicrophoneConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MockMicrophoneConfiguration) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
		p.DeviceId.Ref(),
		p.GroupId.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
	p.DeviceId = p.DeviceId.FromRef(js.Undefined)
	p.GroupId = p.GroupId.FromRef(js.Undefined)
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
func (p *MockSensor) UpdateFrom(ref js.Ref) {
	bindings.MockSensorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MockSensor) Update(ref js.Ref) {
	bindings.MockSensorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MockSensor) FreeMembers(recursive bool) {
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
func (p *MockSensorConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MockSensorConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MockSensorConfiguration) Update(ref js.Ref) {
	bindings.MockSensorConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MockSensorConfiguration) FreeMembers(recursive bool) {
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
func (p *MockSensorReadingValues) UpdateFrom(ref js.Ref) {
	bindings.MockSensorReadingValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MockSensorReadingValues) Update(ref js.Ref) {
	bindings.MockSensorReadingValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MockSensorReadingValues) FreeMembers(recursive bool) {
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
func (p *MouseEventInit) UpdateFrom(ref js.Ref) {
	bindings.MouseEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MouseEventInit) Update(ref js.Ref) {
	bindings.MouseEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MouseEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.RelatedTarget.Ref(),
		p.View.Ref(),
	)
	p.RelatedTarget = p.RelatedTarget.FromRef(js.Undefined)
	p.View = p.View.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// ScreenX returns the value of property "MouseEvent.screenX".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) ScreenX() (ret int32, ok bool) {
	ok = js.True == bindings.GetMouseEventScreenX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ScreenY returns the value of property "MouseEvent.screenY".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) ScreenY() (ret int32, ok bool) {
	ok = js.True == bindings.GetMouseEventScreenY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ClientX returns the value of property "MouseEvent.clientX".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) ClientX() (ret int32, ok bool) {
	ok = js.True == bindings.GetMouseEventClientX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ClientY returns the value of property "MouseEvent.clientY".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) ClientY() (ret int32, ok bool) {
	ok = js.True == bindings.GetMouseEventClientY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CtrlKey returns the value of property "MouseEvent.ctrlKey".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) CtrlKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetMouseEventCtrlKey(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ShiftKey returns the value of property "MouseEvent.shiftKey".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) ShiftKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetMouseEventShiftKey(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AltKey returns the value of property "MouseEvent.altKey".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) AltKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetMouseEventAltKey(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MetaKey returns the value of property "MouseEvent.metaKey".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) MetaKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetMouseEventMetaKey(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Button returns the value of property "MouseEvent.button".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) Button() (ret int16, ok bool) {
	ok = js.True == bindings.GetMouseEventButton(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Buttons returns the value of property "MouseEvent.buttons".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) Buttons() (ret uint16, ok bool) {
	ok = js.True == bindings.GetMouseEventButtons(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RelatedTarget returns the value of property "MouseEvent.relatedTarget".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) RelatedTarget() (ret EventTarget, ok bool) {
	ok = js.True == bindings.GetMouseEventRelatedTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MovementX returns the value of property "MouseEvent.movementX".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) MovementX() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventMovementX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MovementY returns the value of property "MouseEvent.movementY".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) MovementY() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventMovementY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PageX returns the value of property "MouseEvent.pageX".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) PageX() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventPageX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PageY returns the value of property "MouseEvent.pageY".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) PageY() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventPageY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// X returns the value of property "MouseEvent.x".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "MouseEvent.y".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OffsetX returns the value of property "MouseEvent.offsetX".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) OffsetX() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventOffsetX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OffsetY returns the value of property "MouseEvent.offsetY".
//
// It returns ok=false if there is no such property.
func (this MouseEvent) OffsetY() (ret float64, ok bool) {
	ok = js.True == bindings.GetMouseEventOffsetY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetModifierState returns true if the method "MouseEvent.getModifierState" exists.
func (this MouseEvent) HasFuncGetModifierState() bool {
	return js.True == bindings.HasFuncMouseEventGetModifierState(
		this.ref,
	)
}

// FuncGetModifierState returns the method "MouseEvent.getModifierState".
func (this MouseEvent) FuncGetModifierState() (fn js.Func[func(keyArg js.String) bool]) {
	bindings.FuncMouseEventGetModifierState(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetModifierState calls the method "MouseEvent.getModifierState".
func (this MouseEvent) GetModifierState(keyArg js.String) (ret bool) {
	bindings.CallMouseEventGetModifierState(
		this.ref, js.Pointer(&ret),
		keyArg.Ref(),
	)

	return
}

// TryGetModifierState calls the method "MouseEvent.getModifierState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryGetModifierState(keyArg js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventGetModifierState(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		keyArg.Ref(),
	)

	return
}

// HasFuncInitMouseEvent returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent(
		this.ref,
	)
}

// FuncInitMouseEvent returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16, relatedTargetArg EventTarget)]) {
	bindings.FuncMouseEventInitMouseEvent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16, relatedTargetArg EventTarget) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16, relatedTargetArg EventTarget) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMouseEvent1 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent1() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent1(
		this.ref,
	)
}

// FuncInitMouseEvent1 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent1() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16)]) {
	bindings.FuncMouseEventInitMouseEvent1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent1 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent1(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool, buttonArg int16) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMouseEvent2 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent2() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent2(
		this.ref,
	)
}

// FuncInitMouseEvent2 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent2() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool)]) {
	bindings.FuncMouseEventInitMouseEvent2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent2 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent2(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool, metaKeyArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMouseEvent3 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent3() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent3(
		this.ref,
	)
}

// FuncInitMouseEvent3 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent3() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool)]) {
	bindings.FuncMouseEventInitMouseEvent3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent3 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent3(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool, shiftKeyArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMouseEvent4 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent4() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent4(
		this.ref,
	)
}

// FuncInitMouseEvent4 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent4() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool)]) {
	bindings.FuncMouseEventInitMouseEvent4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent4 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent4(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool, altKeyArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMouseEvent5 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent5() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent5(
		this.ref,
	)
}

// FuncInitMouseEvent5 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent5() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool)]) {
	bindings.FuncMouseEventInitMouseEvent5(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent5 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent5(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32, ctrlKeyArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent5(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMouseEvent6 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent6() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent6(
		this.ref,
	)
}

// FuncInitMouseEvent6 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent6() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32)]) {
	bindings.FuncMouseEventInitMouseEvent6(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent6 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent6(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent6(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent6(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32, clientYArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent6(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMouseEvent7 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent7() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent7(
		this.ref,
	)
}

// FuncInitMouseEvent7 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent7() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32)]) {
	bindings.FuncMouseEventInitMouseEvent7(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent7 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent7(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent7(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent7(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32, clientXArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent7(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMouseEvent8 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent8() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent8(
		this.ref,
	)
}

// FuncInitMouseEvent8 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent8() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32)]) {
	bindings.FuncMouseEventInitMouseEvent8(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent8 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent8(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent8(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent8(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32, screenYArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent8(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMouseEvent9 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent9() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent9(
		this.ref,
	)
}

// FuncInitMouseEvent9 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent9() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32)]) {
	bindings.FuncMouseEventInitMouseEvent9(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent9 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent9(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent9(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent9(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32, screenXArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent9(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
		int32(screenXArg),
	)

	return
}

// HasFuncInitMouseEvent10 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent10() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent10(
		this.ref,
	)
}

// FuncInitMouseEvent10 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent10() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32)]) {
	bindings.FuncMouseEventInitMouseEvent10(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent10 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent10(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent10(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
	)

	return
}

// TryInitMouseEvent10 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent10(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, detailArg int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent10(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		int32(detailArg),
	)

	return
}

// HasFuncInitMouseEvent11 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent11() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent11(
		this.ref,
	)
}

// FuncInitMouseEvent11 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent11() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window)]) {
	bindings.FuncMouseEventInitMouseEvent11(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent11 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent11(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent11(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// TryInitMouseEvent11 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent11(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent11(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// HasFuncInitMouseEvent12 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent12() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent12(
		this.ref,
	)
}

// FuncInitMouseEvent12 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent12() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	bindings.FuncMouseEventInitMouseEvent12(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent12 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent12(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent12(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// TryInitMouseEvent12 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent12(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent12(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// HasFuncInitMouseEvent13 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent13() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent13(
		this.ref,
	)
}

// FuncInitMouseEvent13 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent13() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	bindings.FuncMouseEventInitMouseEvent13(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent13 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent13(typeArg js.String, bubblesArg bool) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent13(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// TryInitMouseEvent13 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent13(typeArg js.String, bubblesArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent13(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// HasFuncInitMouseEvent14 returns true if the method "MouseEvent.initMouseEvent" exists.
func (this MouseEvent) HasFuncInitMouseEvent14() bool {
	return js.True == bindings.HasFuncMouseEventInitMouseEvent14(
		this.ref,
	)
}

// FuncInitMouseEvent14 returns the method "MouseEvent.initMouseEvent".
func (this MouseEvent) FuncInitMouseEvent14() (fn js.Func[func(typeArg js.String)]) {
	bindings.FuncMouseEventInitMouseEvent14(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMouseEvent14 calls the method "MouseEvent.initMouseEvent".
func (this MouseEvent) InitMouseEvent14(typeArg js.String) (ret js.Void) {
	bindings.CallMouseEventInitMouseEvent14(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
	)

	return
}

// TryInitMouseEvent14 calls the method "MouseEvent.initMouseEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MouseEvent) TryInitMouseEvent14(typeArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMouseEventInitMouseEvent14(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "MutationRecord.type".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationRecordType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Target returns the value of property "MutationRecord.target".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) Target() (ret Node, ok bool) {
	ok = js.True == bindings.GetMutationRecordTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AddedNodes returns the value of property "MutationRecord.addedNodes".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) AddedNodes() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetMutationRecordAddedNodes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// RemovedNodes returns the value of property "MutationRecord.removedNodes".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) RemovedNodes() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetMutationRecordRemovedNodes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PreviousSibling returns the value of property "MutationRecord.previousSibling".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) PreviousSibling() (ret Node, ok bool) {
	ok = js.True == bindings.GetMutationRecordPreviousSibling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NextSibling returns the value of property "MutationRecord.nextSibling".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) NextSibling() (ret Node, ok bool) {
	ok = js.True == bindings.GetMutationRecordNextSibling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AttributeName returns the value of property "MutationRecord.attributeName".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) AttributeName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationRecordAttributeName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AttributeNamespace returns the value of property "MutationRecord.attributeNamespace".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) AttributeNamespace() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationRecordAttributeNamespace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OldValue returns the value of property "MutationRecord.oldValue".
//
// It returns ok=false if there is no such property.
func (this MutationRecord) OldValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationRecordOldValue(
		this.ref, js.Pointer(&ret),
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
func (p *MutationObserverInit) UpdateFrom(ref js.Ref) {
	bindings.MutationObserverInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MutationObserverInit) Update(ref js.Ref) {
	bindings.MutationObserverInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MutationObserverInit) FreeMembers(recursive bool) {
	js.Free(
		p.AttributeFilter.Ref(),
	)
	p.AttributeFilter = p.AttributeFilter.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncObserve returns true if the method "MutationObserver.observe" exists.
func (this MutationObserver) HasFuncObserve() bool {
	return js.True == bindings.HasFuncMutationObserverObserve(
		this.ref,
	)
}

// FuncObserve returns the method "MutationObserver.observe".
func (this MutationObserver) FuncObserve() (fn js.Func[func(target Node, options MutationObserverInit)]) {
	bindings.FuncMutationObserverObserve(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Observe calls the method "MutationObserver.observe".
func (this MutationObserver) Observe(target Node, options MutationObserverInit) (ret js.Void) {
	bindings.CallMutationObserverObserve(
		this.ref, js.Pointer(&ret),
		target.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryObserve calls the method "MutationObserver.observe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationObserver) TryObserve(target Node, options MutationObserverInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationObserverObserve(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		target.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncObserve1 returns true if the method "MutationObserver.observe" exists.
func (this MutationObserver) HasFuncObserve1() bool {
	return js.True == bindings.HasFuncMutationObserverObserve1(
		this.ref,
	)
}

// FuncObserve1 returns the method "MutationObserver.observe".
func (this MutationObserver) FuncObserve1() (fn js.Func[func(target Node)]) {
	bindings.FuncMutationObserverObserve1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Observe1 calls the method "MutationObserver.observe".
func (this MutationObserver) Observe1(target Node) (ret js.Void) {
	bindings.CallMutationObserverObserve1(
		this.ref, js.Pointer(&ret),
		target.Ref(),
	)

	return
}

// TryObserve1 calls the method "MutationObserver.observe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationObserver) TryObserve1(target Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationObserverObserve1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		target.Ref(),
	)

	return
}

// HasFuncDisconnect returns true if the method "MutationObserver.disconnect" exists.
func (this MutationObserver) HasFuncDisconnect() bool {
	return js.True == bindings.HasFuncMutationObserverDisconnect(
		this.ref,
	)
}

// FuncDisconnect returns the method "MutationObserver.disconnect".
func (this MutationObserver) FuncDisconnect() (fn js.Func[func()]) {
	bindings.FuncMutationObserverDisconnect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disconnect calls the method "MutationObserver.disconnect".
func (this MutationObserver) Disconnect() (ret js.Void) {
	bindings.CallMutationObserverDisconnect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDisconnect calls the method "MutationObserver.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationObserver) TryDisconnect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationObserverDisconnect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncTakeRecords returns true if the method "MutationObserver.takeRecords" exists.
func (this MutationObserver) HasFuncTakeRecords() bool {
	return js.True == bindings.HasFuncMutationObserverTakeRecords(
		this.ref,
	)
}

// FuncTakeRecords returns the method "MutationObserver.takeRecords".
func (this MutationObserver) FuncTakeRecords() (fn js.Func[func() js.Array[MutationRecord]]) {
	bindings.FuncMutationObserverTakeRecords(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TakeRecords calls the method "MutationObserver.takeRecords".
func (this MutationObserver) TakeRecords() (ret js.Array[MutationRecord]) {
	bindings.CallMutationObserverTakeRecords(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTakeRecords calls the method "MutationObserver.takeRecords"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationObserver) TryTakeRecords() (ret js.Array[MutationRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationObserverTakeRecords(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// RelatedNode returns the value of property "MutationEvent.relatedNode".
//
// It returns ok=false if there is no such property.
func (this MutationEvent) RelatedNode() (ret Node, ok bool) {
	ok = js.True == bindings.GetMutationEventRelatedNode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PrevValue returns the value of property "MutationEvent.prevValue".
//
// It returns ok=false if there is no such property.
func (this MutationEvent) PrevValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationEventPrevValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NewValue returns the value of property "MutationEvent.newValue".
//
// It returns ok=false if there is no such property.
func (this MutationEvent) NewValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationEventNewValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AttrName returns the value of property "MutationEvent.attrName".
//
// It returns ok=false if there is no such property.
func (this MutationEvent) AttrName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMutationEventAttrName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AttrChange returns the value of property "MutationEvent.attrChange".
//
// It returns ok=false if there is no such property.
func (this MutationEvent) AttrChange() (ret uint16, ok bool) {
	ok = js.True == bindings.GetMutationEventAttrChange(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncInitMutationEvent returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasFuncInitMutationEvent() bool {
	return js.True == bindings.HasFuncMutationEventInitMutationEvent(
		this.ref,
	)
}

// FuncInitMutationEvent returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) FuncInitMutationEvent() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String, attrChangeArg uint16)]) {
	bindings.FuncMutationEventInitMutationEvent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMutationEvent calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String, attrChangeArg uint16) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationEvent) TryInitMutationEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String, attrChangeArg uint16) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMutationEvent1 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasFuncInitMutationEvent1() bool {
	return js.True == bindings.HasFuncMutationEventInitMutationEvent1(
		this.ref,
	)
}

// FuncInitMutationEvent1 returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) FuncInitMutationEvent1() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String)]) {
	bindings.FuncMutationEventInitMutationEvent1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMutationEvent1 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent1(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationEvent) TryInitMutationEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String, attrNameArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

// HasFuncInitMutationEvent2 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasFuncInitMutationEvent2() bool {
	return js.True == bindings.HasFuncMutationEventInitMutationEvent2(
		this.ref,
	)
}

// FuncInitMutationEvent2 returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) FuncInitMutationEvent2() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String)]) {
	bindings.FuncMutationEventInitMutationEvent2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMutationEvent2 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent2(
		this.ref, js.Pointer(&ret),
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
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationEvent) TryInitMutationEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String, newValueArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
		newValueArg.Ref(),
	)

	return
}

// HasFuncInitMutationEvent3 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasFuncInitMutationEvent3() bool {
	return js.True == bindings.HasFuncMutationEventInitMutationEvent3(
		this.ref,
	)
}

// FuncInitMutationEvent3 returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) FuncInitMutationEvent3() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String)]) {
	bindings.FuncMutationEventInitMutationEvent3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMutationEvent3 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent3(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
	)

	return
}

// TryInitMutationEvent3 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationEvent) TryInitMutationEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node, prevValueArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
		prevValueArg.Ref(),
	)

	return
}

// HasFuncInitMutationEvent4 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasFuncInitMutationEvent4() bool {
	return js.True == bindings.HasFuncMutationEventInitMutationEvent4(
		this.ref,
	)
}

// FuncInitMutationEvent4 returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) FuncInitMutationEvent4() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node)]) {
	bindings.FuncMutationEventInitMutationEvent4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMutationEvent4 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent4(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
	)

	return
}

// TryInitMutationEvent4 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationEvent) TryInitMutationEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, relatedNodeArg Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		relatedNodeArg.Ref(),
	)

	return
}

// HasFuncInitMutationEvent5 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasFuncInitMutationEvent5() bool {
	return js.True == bindings.HasFuncMutationEventInitMutationEvent5(
		this.ref,
	)
}

// FuncInitMutationEvent5 returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) FuncInitMutationEvent5() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	bindings.FuncMutationEventInitMutationEvent5(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMutationEvent5 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent5(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// TryInitMutationEvent5 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationEvent) TryInitMutationEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent5(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// HasFuncInitMutationEvent6 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasFuncInitMutationEvent6() bool {
	return js.True == bindings.HasFuncMutationEventInitMutationEvent6(
		this.ref,
	)
}

// FuncInitMutationEvent6 returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) FuncInitMutationEvent6() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	bindings.FuncMutationEventInitMutationEvent6(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMutationEvent6 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent6(typeArg js.String, bubblesArg bool) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent6(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// TryInitMutationEvent6 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationEvent) TryInitMutationEvent6(typeArg js.String, bubblesArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent6(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// HasFuncInitMutationEvent7 returns true if the method "MutationEvent.initMutationEvent" exists.
func (this MutationEvent) HasFuncInitMutationEvent7() bool {
	return js.True == bindings.HasFuncMutationEventInitMutationEvent7(
		this.ref,
	)
}

// FuncInitMutationEvent7 returns the method "MutationEvent.initMutationEvent".
func (this MutationEvent) FuncInitMutationEvent7() (fn js.Func[func(typeArg js.String)]) {
	bindings.FuncMutationEventInitMutationEvent7(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitMutationEvent7 calls the method "MutationEvent.initMutationEvent".
func (this MutationEvent) InitMutationEvent7(typeArg js.String) (ret js.Void) {
	bindings.CallMutationEventInitMutationEvent7(
		this.ref, js.Pointer(&ret),
		typeArg.Ref(),
	)

	return
}

// TryInitMutationEvent7 calls the method "MutationEvent.initMutationEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MutationEvent) TryInitMutationEvent7(typeArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMutationEventInitMutationEvent7(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *NDEFMakeReadOnlyOptions) UpdateFrom(ref js.Ref) {
	bindings.NDEFMakeReadOnlyOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NDEFMakeReadOnlyOptions) Update(ref js.Ref) {
	bindings.NDEFMakeReadOnlyOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NDEFMakeReadOnlyOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
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
func (p *NDEFRecordInit) UpdateFrom(ref js.Ref) {
	bindings.NDEFRecordInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NDEFRecordInit) Update(ref js.Ref) {
	bindings.NDEFRecordInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NDEFRecordInit) FreeMembers(recursive bool) {
	js.Free(
		p.RecordType.Ref(),
		p.MediaType.Ref(),
		p.Id.Ref(),
		p.Encoding.Ref(),
		p.Lang.Ref(),
		p.Data.Ref(),
	)
	p.RecordType = p.RecordType.FromRef(js.Undefined)
	p.MediaType = p.MediaType.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Encoding = p.Encoding.FromRef(js.Undefined)
	p.Lang = p.Lang.FromRef(js.Undefined)
	p.Data = p.Data.FromRef(js.Undefined)
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
func (p *NDEFMessageInit) UpdateFrom(ref js.Ref) {
	bindings.NDEFMessageInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NDEFMessageInit) Update(ref js.Ref) {
	bindings.NDEFMessageInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NDEFMessageInit) FreeMembers(recursive bool) {
	js.Free(
		p.Records.Ref(),
	)
	p.Records = p.Records.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// RecordType returns the value of property "NDEFRecord.recordType".
//
// It returns ok=false if there is no such property.
func (this NDEFRecord) RecordType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFRecordRecordType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MediaType returns the value of property "NDEFRecord.mediaType".
//
// It returns ok=false if there is no such property.
func (this NDEFRecord) MediaType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFRecordMediaType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "NDEFRecord.id".
//
// It returns ok=false if there is no such property.
func (this NDEFRecord) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFRecordId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Data returns the value of property "NDEFRecord.data".
//
// It returns ok=false if there is no such property.
func (this NDEFRecord) Data() (ret js.DataView, ok bool) {
	ok = js.True == bindings.GetNDEFRecordData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Encoding returns the value of property "NDEFRecord.encoding".
//
// It returns ok=false if there is no such property.
func (this NDEFRecord) Encoding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFRecordEncoding(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Lang returns the value of property "NDEFRecord.lang".
//
// It returns ok=false if there is no such property.
func (this NDEFRecord) Lang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFRecordLang(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToRecords returns true if the method "NDEFRecord.toRecords" exists.
func (this NDEFRecord) HasFuncToRecords() bool {
	return js.True == bindings.HasFuncNDEFRecordToRecords(
		this.ref,
	)
}

// FuncToRecords returns the method "NDEFRecord.toRecords".
func (this NDEFRecord) FuncToRecords() (fn js.Func[func() js.Array[NDEFRecord]]) {
	bindings.FuncNDEFRecordToRecords(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToRecords calls the method "NDEFRecord.toRecords".
func (this NDEFRecord) ToRecords() (ret js.Array[NDEFRecord]) {
	bindings.CallNDEFRecordToRecords(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToRecords calls the method "NDEFRecord.toRecords"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NDEFRecord) TryToRecords() (ret js.Array[NDEFRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFRecordToRecords(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Records returns the value of property "NDEFMessage.records".
//
// It returns ok=false if there is no such property.
func (this NDEFMessage) Records() (ret js.FrozenArray[NDEFRecord], ok bool) {
	ok = js.True == bindings.GetNDEFMessageRecords(
		this.ref, js.Pointer(&ret),
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
func (p *NDEFScanOptions) UpdateFrom(ref js.Ref) {
	bindings.NDEFScanOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NDEFScanOptions) Update(ref js.Ref) {
	bindings.NDEFScanOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NDEFScanOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
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
func (p *NDEFWriteOptions) UpdateFrom(ref js.Ref) {
	bindings.NDEFWriteOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NDEFWriteOptions) Update(ref js.Ref) {
	bindings.NDEFWriteOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NDEFWriteOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
}

type NDEFReader struct {
	EventTarget
}

func (this NDEFReader) Once() NDEFReader {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncScan returns true if the method "NDEFReader.scan" exists.
func (this NDEFReader) HasFuncScan() bool {
	return js.True == bindings.HasFuncNDEFReaderScan(
		this.ref,
	)
}

// FuncScan returns the method "NDEFReader.scan".
func (this NDEFReader) FuncScan() (fn js.Func[func(options NDEFScanOptions) js.Promise[js.Void]]) {
	bindings.FuncNDEFReaderScan(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scan calls the method "NDEFReader.scan".
func (this NDEFReader) Scan(options NDEFScanOptions) (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderScan(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryScan calls the method "NDEFReader.scan"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NDEFReader) TryScan(options NDEFScanOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderScan(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncScan1 returns true if the method "NDEFReader.scan" exists.
func (this NDEFReader) HasFuncScan1() bool {
	return js.True == bindings.HasFuncNDEFReaderScan1(
		this.ref,
	)
}

// FuncScan1 returns the method "NDEFReader.scan".
func (this NDEFReader) FuncScan1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncNDEFReaderScan1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Scan1 calls the method "NDEFReader.scan".
func (this NDEFReader) Scan1() (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderScan1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryScan1 calls the method "NDEFReader.scan"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NDEFReader) TryScan1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderScan1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncWrite returns true if the method "NDEFReader.write" exists.
func (this NDEFReader) HasFuncWrite() bool {
	return js.True == bindings.HasFuncNDEFReaderWrite(
		this.ref,
	)
}

// FuncWrite returns the method "NDEFReader.write".
func (this NDEFReader) FuncWrite() (fn js.Func[func(message NDEFMessageSource, options NDEFWriteOptions) js.Promise[js.Void]]) {
	bindings.FuncNDEFReaderWrite(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Write calls the method "NDEFReader.write".
func (this NDEFReader) Write(message NDEFMessageSource, options NDEFWriteOptions) (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderWrite(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryWrite calls the method "NDEFReader.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NDEFReader) TryWrite(message NDEFMessageSource, options NDEFWriteOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderWrite(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncWrite1 returns true if the method "NDEFReader.write" exists.
func (this NDEFReader) HasFuncWrite1() bool {
	return js.True == bindings.HasFuncNDEFReaderWrite1(
		this.ref,
	)
}

// FuncWrite1 returns the method "NDEFReader.write".
func (this NDEFReader) FuncWrite1() (fn js.Func[func(message NDEFMessageSource) js.Promise[js.Void]]) {
	bindings.FuncNDEFReaderWrite1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Write1 calls the method "NDEFReader.write".
func (this NDEFReader) Write1(message NDEFMessageSource) (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderWrite1(
		this.ref, js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryWrite1 calls the method "NDEFReader.write"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NDEFReader) TryWrite1(message NDEFMessageSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderWrite1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

// HasFuncMakeReadOnly returns true if the method "NDEFReader.makeReadOnly" exists.
func (this NDEFReader) HasFuncMakeReadOnly() bool {
	return js.True == bindings.HasFuncNDEFReaderMakeReadOnly(
		this.ref,
	)
}

// FuncMakeReadOnly returns the method "NDEFReader.makeReadOnly".
func (this NDEFReader) FuncMakeReadOnly() (fn js.Func[func(options NDEFMakeReadOnlyOptions) js.Promise[js.Void]]) {
	bindings.FuncNDEFReaderMakeReadOnly(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MakeReadOnly calls the method "NDEFReader.makeReadOnly".
func (this NDEFReader) MakeReadOnly(options NDEFMakeReadOnlyOptions) (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderMakeReadOnly(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryMakeReadOnly calls the method "NDEFReader.makeReadOnly"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NDEFReader) TryMakeReadOnly(options NDEFMakeReadOnlyOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderMakeReadOnly(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncMakeReadOnly1 returns true if the method "NDEFReader.makeReadOnly" exists.
func (this NDEFReader) HasFuncMakeReadOnly1() bool {
	return js.True == bindings.HasFuncNDEFReaderMakeReadOnly1(
		this.ref,
	)
}

// FuncMakeReadOnly1 returns the method "NDEFReader.makeReadOnly".
func (this NDEFReader) FuncMakeReadOnly1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncNDEFReaderMakeReadOnly1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// MakeReadOnly1 calls the method "NDEFReader.makeReadOnly".
func (this NDEFReader) MakeReadOnly1() (ret js.Promise[js.Void]) {
	bindings.CallNDEFReaderMakeReadOnly1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryMakeReadOnly1 calls the method "NDEFReader.makeReadOnly"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NDEFReader) TryMakeReadOnly1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNDEFReaderMakeReadOnly1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *NDEFReadingEventInit) UpdateFrom(ref js.Ref) {
	bindings.NDEFReadingEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NDEFReadingEventInit) Update(ref js.Ref) {
	bindings.NDEFReadingEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NDEFReadingEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.SerialNumber.Ref(),
	)
	p.SerialNumber = p.SerialNumber.FromRef(js.Undefined)
	if recursive {
		p.Message.FreeMembers(true)
	}
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
	this.ref.Once()
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
	this.ref.Free()
}

// SerialNumber returns the value of property "NDEFReadingEvent.serialNumber".
//
// It returns ok=false if there is no such property.
func (this NDEFReadingEvent) SerialNumber() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNDEFReadingEventSerialNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "NDEFReadingEvent.message".
//
// It returns ok=false if there is no such property.
func (this NDEFReadingEvent) Message() (ret NDEFMessage, ok bool) {
	ok = js.True == bindings.GetNDEFReadingEventMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

type NamedFlow struct {
	EventTarget
}

func (this NamedFlow) Once() NamedFlow {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "NamedFlow.name".
//
// It returns ok=false if there is no such property.
func (this NamedFlow) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNamedFlowName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Overset returns the value of property "NamedFlow.overset".
//
// It returns ok=false if there is no such property.
func (this NamedFlow) Overset() (ret bool, ok bool) {
	ok = js.True == bindings.GetNamedFlowOverset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FirstEmptyRegionIndex returns the value of property "NamedFlow.firstEmptyRegionIndex".
//
// It returns ok=false if there is no such property.
func (this NamedFlow) FirstEmptyRegionIndex() (ret int16, ok bool) {
	ok = js.True == bindings.GetNamedFlowFirstEmptyRegionIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetRegions returns true if the method "NamedFlow.getRegions" exists.
func (this NamedFlow) HasFuncGetRegions() bool {
	return js.True == bindings.HasFuncNamedFlowGetRegions(
		this.ref,
	)
}

// FuncGetRegions returns the method "NamedFlow.getRegions".
func (this NamedFlow) FuncGetRegions() (fn js.Func[func() js.Array[Element]]) {
	bindings.FuncNamedFlowGetRegions(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRegions calls the method "NamedFlow.getRegions".
func (this NamedFlow) GetRegions() (ret js.Array[Element]) {
	bindings.CallNamedFlowGetRegions(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetRegions calls the method "NamedFlow.getRegions"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NamedFlow) TryGetRegions() (ret js.Array[Element], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedFlowGetRegions(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetContent returns true if the method "NamedFlow.getContent" exists.
func (this NamedFlow) HasFuncGetContent() bool {
	return js.True == bindings.HasFuncNamedFlowGetContent(
		this.ref,
	)
}

// FuncGetContent returns the method "NamedFlow.getContent".
func (this NamedFlow) FuncGetContent() (fn js.Func[func() js.Array[Node]]) {
	bindings.FuncNamedFlowGetContent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetContent calls the method "NamedFlow.getContent".
func (this NamedFlow) GetContent() (ret js.Array[Node]) {
	bindings.CallNamedFlowGetContent(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetContent calls the method "NamedFlow.getContent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NamedFlow) TryGetContent() (ret js.Array[Node], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedFlowGetContent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetRegionsByContent returns true if the method "NamedFlow.getRegionsByContent" exists.
func (this NamedFlow) HasFuncGetRegionsByContent() bool {
	return js.True == bindings.HasFuncNamedFlowGetRegionsByContent(
		this.ref,
	)
}

// FuncGetRegionsByContent returns the method "NamedFlow.getRegionsByContent".
func (this NamedFlow) FuncGetRegionsByContent() (fn js.Func[func(node Node) js.Array[Element]]) {
	bindings.FuncNamedFlowGetRegionsByContent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRegionsByContent calls the method "NamedFlow.getRegionsByContent".
func (this NamedFlow) GetRegionsByContent(node Node) (ret js.Array[Element]) {
	bindings.CallNamedFlowGetRegionsByContent(
		this.ref, js.Pointer(&ret),
		node.Ref(),
	)

	return
}

// TryGetRegionsByContent calls the method "NamedFlow.getRegionsByContent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NamedFlow) TryGetRegionsByContent(node Node) (ret js.Array[Element], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNamedFlowGetRegionsByContent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		node.Ref(),
	)

	return
}

type NavigationDestination struct {
	ref js.Ref
}

func (this NavigationDestination) Once() NavigationDestination {
	this.ref.Once()
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
	this.ref.Free()
}

// Url returns the value of property "NavigationDestination.url".
//
// It returns ok=false if there is no such property.
func (this NavigationDestination) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationDestinationUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Key returns the value of property "NavigationDestination.key".
//
// It returns ok=false if there is no such property.
func (this NavigationDestination) Key() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationDestinationKey(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "NavigationDestination.id".
//
// It returns ok=false if there is no such property.
func (this NavigationDestination) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigationDestinationId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Index returns the value of property "NavigationDestination.index".
//
// It returns ok=false if there is no such property.
func (this NavigationDestination) Index() (ret int64, ok bool) {
	ok = js.True == bindings.GetNavigationDestinationIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SameDocument returns the value of property "NavigationDestination.sameDocument".
//
// It returns ok=false if there is no such property.
func (this NavigationDestination) SameDocument() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigationDestinationSameDocument(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetState returns true if the method "NavigationDestination.getState" exists.
func (this NavigationDestination) HasFuncGetState() bool {
	return js.True == bindings.HasFuncNavigationDestinationGetState(
		this.ref,
	)
}

// FuncGetState returns the method "NavigationDestination.getState".
func (this NavigationDestination) FuncGetState() (fn js.Func[func() js.Any]) {
	bindings.FuncNavigationDestinationGetState(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetState calls the method "NavigationDestination.getState".
func (this NavigationDestination) GetState() (ret js.Any) {
	bindings.CallNavigationDestinationGetState(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetState calls the method "NavigationDestination.getState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NavigationDestination) TryGetState() (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationDestinationGetState(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *NavigateEventInit) UpdateFrom(ref js.Ref) {
	bindings.NavigateEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NavigateEventInit) Update(ref js.Ref) {
	bindings.NavigateEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NavigateEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Destination.Ref(),
		p.Signal.Ref(),
		p.FormData.Ref(),
		p.DownloadRequest.Ref(),
		p.Info.Ref(),
	)
	p.Destination = p.Destination.FromRef(js.Undefined)
	p.Signal = p.Signal.FromRef(js.Undefined)
	p.FormData = p.FormData.FromRef(js.Undefined)
	p.DownloadRequest = p.DownloadRequest.FromRef(js.Undefined)
	p.Info = p.Info.FromRef(js.Undefined)
}

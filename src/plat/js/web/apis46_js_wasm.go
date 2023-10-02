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

func NewPresentationConnectionAvailableEvent(typ js.String, eventInitDict PresentationConnectionAvailableEventInit) PresentationConnectionAvailableEvent {
	return PresentationConnectionAvailableEvent{}.FromRef(
		bindings.NewPresentationConnectionAvailableEventByPresentationConnectionAvailableEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type PresentationConnectionAvailableEvent struct {
	Event
}

func (this PresentationConnectionAvailableEvent) Once() PresentationConnectionAvailableEvent {
	this.Ref().Once()
	return this
}

func (this PresentationConnectionAvailableEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this PresentationConnectionAvailableEvent) FromRef(ref js.Ref) PresentationConnectionAvailableEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this PresentationConnectionAvailableEvent) Free() {
	this.Ref().Free()
}

// Connection returns the value of property "PresentationConnectionAvailableEvent.connection".
//
// The returned bool will be false if there is no such property.
func (this PresentationConnectionAvailableEvent) Connection() (PresentationConnection, bool) {
	var _ok bool
	_ret := bindings.GetPresentationConnectionAvailableEventConnection(
		this.Ref(), js.Pointer(&_ok),
	)
	return PresentationConnection{}.FromRef(_ret), _ok
}

type PresentationConnectionCloseReason uint32

const (
	_ PresentationConnectionCloseReason = iota

	PresentationConnectionCloseReason_ERROR
	PresentationConnectionCloseReason_CLOSED
	PresentationConnectionCloseReason_WENTAWAY
)

func (PresentationConnectionCloseReason) FromRef(str js.Ref) PresentationConnectionCloseReason {
	return PresentationConnectionCloseReason(bindings.ConstOfPresentationConnectionCloseReason(str))
}

func (x PresentationConnectionCloseReason) String() (string, bool) {
	switch x {
	case PresentationConnectionCloseReason_ERROR:
		return "error", true
	case PresentationConnectionCloseReason_CLOSED:
		return "closed", true
	case PresentationConnectionCloseReason_WENTAWAY:
		return "wentaway", true
	default:
		return "", false
	}
}

type PresentationConnectionCloseEventInit struct {
	// Reason is "PresentationConnectionCloseEventInit.reason"
	//
	// Required
	Reason PresentationConnectionCloseReason
	// Message is "PresentationConnectionCloseEventInit.message"
	//
	// Optional, defaults to "".
	Message js.String
	// Bubbles is "PresentationConnectionCloseEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "PresentationConnectionCloseEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PresentationConnectionCloseEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PresentationConnectionCloseEventInit with all fields set.
func (p PresentationConnectionCloseEventInit) FromRef(ref js.Ref) PresentationConnectionCloseEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PresentationConnectionCloseEventInit PresentationConnectionCloseEventInit [// PresentationConnectionCloseEventInit] [0x14000960460 0x14000960500 0x140009605a0 0x140009606e0 0x14000960820 0x14000960640 0x14000960780 0x140009608c0] 0x14000920f78 {0 0}} in the application heap.
func (p PresentationConnectionCloseEventInit) New() js.Ref {
	return bindings.PresentationConnectionCloseEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PresentationConnectionCloseEventInit) UpdateFrom(ref js.Ref) {
	bindings.PresentationConnectionCloseEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PresentationConnectionCloseEventInit) Update(ref js.Ref) {
	bindings.PresentationConnectionCloseEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPresentationConnectionCloseEvent(typ js.String, eventInitDict PresentationConnectionCloseEventInit) PresentationConnectionCloseEvent {
	return PresentationConnectionCloseEvent{}.FromRef(
		bindings.NewPresentationConnectionCloseEventByPresentationConnectionCloseEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type PresentationConnectionCloseEvent struct {
	Event
}

func (this PresentationConnectionCloseEvent) Once() PresentationConnectionCloseEvent {
	this.Ref().Once()
	return this
}

func (this PresentationConnectionCloseEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this PresentationConnectionCloseEvent) FromRef(ref js.Ref) PresentationConnectionCloseEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this PresentationConnectionCloseEvent) Free() {
	this.Ref().Free()
}

// Reason returns the value of property "PresentationConnectionCloseEvent.reason".
//
// The returned bool will be false if there is no such property.
func (this PresentationConnectionCloseEvent) Reason() (PresentationConnectionCloseReason, bool) {
	var _ok bool
	_ret := bindings.GetPresentationConnectionCloseEventReason(
		this.Ref(), js.Pointer(&_ok),
	)
	return PresentationConnectionCloseReason(_ret), _ok
}

// Message returns the value of property "PresentationConnectionCloseEvent.message".
//
// The returned bool will be false if there is no such property.
func (this PresentationConnectionCloseEvent) Message() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPresentationConnectionCloseEventMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type PressureUpdateCallbackFunc func(this js.Ref, changes js.Array[PressureRecord], observer PressureObserver) js.Ref

func (fn PressureUpdateCallbackFunc) Register() js.Func[func(changes js.Array[PressureRecord], observer PressureObserver)] {
	return js.RegisterCallback[func(changes js.Array[PressureRecord], observer PressureObserver)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PressureUpdateCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[PressureRecord]{}.FromRef(args[0+1]),
		PressureObserver{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PressureUpdateCallback[T any] struct {
	Fn  func(arg T, this js.Ref, changes js.Array[PressureRecord], observer PressureObserver) js.Ref
	Arg T
}

func (cb *PressureUpdateCallback[T]) Register() js.Func[func(changes js.Array[PressureRecord], observer PressureObserver)] {
	return js.RegisterCallback[func(changes js.Array[PressureRecord], observer PressureObserver)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PressureUpdateCallback[T]) DispatchCallback(
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

		js.Array[PressureRecord]{}.FromRef(args[0+1]),
		PressureObserver{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PressureSource uint32

const (
	_ PressureSource = iota

	PressureSource_THERMALS
	PressureSource_CPU
)

func (PressureSource) FromRef(str js.Ref) PressureSource {
	return PressureSource(bindings.ConstOfPressureSource(str))
}

func (x PressureSource) String() (string, bool) {
	switch x {
	case PressureSource_THERMALS:
		return "thermals", true
	case PressureSource_CPU:
		return "cpu", true
	default:
		return "", false
	}
}

type PressureState uint32

const (
	_ PressureState = iota

	PressureState_NOMINAL
	PressureState_FAIR
	PressureState_SERIOUS
	PressureState_CRITICAL
)

func (PressureState) FromRef(str js.Ref) PressureState {
	return PressureState(bindings.ConstOfPressureState(str))
}

func (x PressureState) String() (string, bool) {
	switch x {
	case PressureState_NOMINAL:
		return "nominal", true
	case PressureState_FAIR:
		return "fair", true
	case PressureState_SERIOUS:
		return "serious", true
	case PressureState_CRITICAL:
		return "critical", true
	default:
		return "", false
	}
}

type PressureRecord struct {
	ref js.Ref
}

func (this PressureRecord) Once() PressureRecord {
	this.Ref().Once()
	return this
}

func (this PressureRecord) Ref() js.Ref {
	return this.ref
}

func (this PressureRecord) FromRef(ref js.Ref) PressureRecord {
	this.ref = ref
	return this
}

func (this PressureRecord) Free() {
	this.Ref().Free()
}

// Source returns the value of property "PressureRecord.source".
//
// The returned bool will be false if there is no such property.
func (this PressureRecord) Source() (PressureSource, bool) {
	var _ok bool
	_ret := bindings.GetPressureRecordSource(
		this.Ref(), js.Pointer(&_ok),
	)
	return PressureSource(_ret), _ok
}

// State returns the value of property "PressureRecord.state".
//
// The returned bool will be false if there is no such property.
func (this PressureRecord) State() (PressureState, bool) {
	var _ok bool
	_ret := bindings.GetPressureRecordState(
		this.Ref(), js.Pointer(&_ok),
	)
	return PressureState(_ret), _ok
}

// Time returns the value of property "PressureRecord.time".
//
// The returned bool will be false if there is no such property.
func (this PressureRecord) Time() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPressureRecordTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// ToJSON calls the method "PressureRecord.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PressureRecord) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPressureRecordToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PressureRecord.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PressureRecord) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PressureRecordToJSONFunc(
			this.Ref(),
		),
	)
}

type PressureObserverOptions struct {
	// SampleRate is "PressureObserverOptions.sampleRate"
	//
	// Optional, defaults to 1.0.
	SampleRate float64

	FFI_USE_SampleRate bool // for SampleRate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PressureObserverOptions with all fields set.
func (p PressureObserverOptions) FromRef(ref js.Ref) PressureObserverOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PressureObserverOptions PressureObserverOptions [// PressureObserverOptions] [0x14000960aa0 0x14000960b40] 0x14000920fc0 {0 0}} in the application heap.
func (p PressureObserverOptions) New() js.Ref {
	return bindings.PressureObserverOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PressureObserverOptions) UpdateFrom(ref js.Ref) {
	bindings.PressureObserverOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PressureObserverOptions) Update(ref js.Ref) {
	bindings.PressureObserverOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPressureObserver(callback js.Func[func(changes js.Array[PressureRecord], observer PressureObserver)], options PressureObserverOptions) PressureObserver {
	return PressureObserver{}.FromRef(
		bindings.NewPressureObserverByPressureObserver(
			callback.Ref(),
			js.Pointer(&options)),
	)
}

func NewPressureObserverByPressureObserver1(callback js.Func[func(changes js.Array[PressureRecord], observer PressureObserver)]) PressureObserver {
	return PressureObserver{}.FromRef(
		bindings.NewPressureObserverByPressureObserver1(
			callback.Ref()),
	)
}

type PressureObserver struct {
	ref js.Ref
}

func (this PressureObserver) Once() PressureObserver {
	this.Ref().Once()
	return this
}

func (this PressureObserver) Ref() js.Ref {
	return this.ref
}

func (this PressureObserver) FromRef(ref js.Ref) PressureObserver {
	this.ref = ref
	return this
}

func (this PressureObserver) Free() {
	this.Ref().Free()
}

// SupportedSources returns the value of property "PressureObserver.supportedSources".
//
// The returned bool will be false if there is no such property.
func (this PressureObserver) SupportedSources() (js.FrozenArray[PressureSource], bool) {
	var _ok bool
	_ret := bindings.GetPressureObserverSupportedSources(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[PressureSource]{}.FromRef(_ret), _ok
}

// Observe calls the method "PressureObserver.observe".
//
// The returned bool will be false if there is no such method.
func (this PressureObserver) Observe(source PressureSource) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallPressureObserverObserve(
		this.Ref(), js.Pointer(&_ok),
		uint32(source),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ObserveFunc returns the method "PressureObserver.observe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PressureObserver) ObserveFunc() (fn js.Func[func(source PressureSource) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PressureObserverObserveFunc(
			this.Ref(),
		),
	)
}

// Unobserve calls the method "PressureObserver.unobserve".
//
// The returned bool will be false if there is no such method.
func (this PressureObserver) Unobserve(source PressureSource) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPressureObserverUnobserve(
		this.Ref(), js.Pointer(&_ok),
		uint32(source),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UnobserveFunc returns the method "PressureObserver.unobserve".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PressureObserver) UnobserveFunc() (fn js.Func[func(source PressureSource)]) {
	return fn.FromRef(
		bindings.PressureObserverUnobserveFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "PressureObserver.disconnect".
//
// The returned bool will be false if there is no such method.
func (this PressureObserver) Disconnect() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPressureObserverDisconnect(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisconnectFunc returns the method "PressureObserver.disconnect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PressureObserver) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PressureObserverDisconnectFunc(
			this.Ref(),
		),
	)
}

// TakeRecords calls the method "PressureObserver.takeRecords".
//
// The returned bool will be false if there is no such method.
func (this PressureObserver) TakeRecords() (js.Array[PressureRecord], bool) {
	var _ok bool
	_ret := bindings.CallPressureObserverTakeRecords(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[PressureRecord]{}.FromRef(_ret), _ok
}

// TakeRecordsFunc returns the method "PressureObserver.takeRecords".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PressureObserver) TakeRecordsFunc() (fn js.Func[func() js.Array[PressureRecord]]) {
	return fn.FromRef(
		bindings.PressureObserverTakeRecordsFunc(
			this.Ref(),
		),
	)
}

type PrivateNetworkAccessPermissionDescriptor struct {
	// Id is "PrivateNetworkAccessPermissionDescriptor.id"
	//
	// Optional
	Id js.String
	// Name is "PrivateNetworkAccessPermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PrivateNetworkAccessPermissionDescriptor with all fields set.
func (p PrivateNetworkAccessPermissionDescriptor) FromRef(ref js.Ref) PrivateNetworkAccessPermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PrivateNetworkAccessPermissionDescriptor PrivateNetworkAccessPermissionDescriptor [// PrivateNetworkAccessPermissionDescriptor] [0x14000960be0 0x14000960c80] 0x14000921008 {0 0}} in the application heap.
func (p PrivateNetworkAccessPermissionDescriptor) New() js.Ref {
	return bindings.PrivateNetworkAccessPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PrivateNetworkAccessPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.PrivateNetworkAccessPermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PrivateNetworkAccessPermissionDescriptor) Update(ref js.Ref) {
	bindings.PrivateNetworkAccessPermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ProfilerInitOptions struct {
	// SampleInterval is "ProfilerInitOptions.sampleInterval"
	//
	// Required
	SampleInterval DOMHighResTimeStamp
	// MaxBufferSize is "ProfilerInitOptions.maxBufferSize"
	//
	// Required
	MaxBufferSize uint32

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProfilerInitOptions with all fields set.
func (p ProfilerInitOptions) FromRef(ref js.Ref) ProfilerInitOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ProfilerInitOptions ProfilerInitOptions [// ProfilerInitOptions] [0x14000960d20 0x14000960dc0] 0x14000921068 {0 0}} in the application heap.
func (p ProfilerInitOptions) New() js.Ref {
	return bindings.ProfilerInitOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ProfilerInitOptions) UpdateFrom(ref js.Ref) {
	bindings.ProfilerInitOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ProfilerInitOptions) Update(ref js.Ref) {
	bindings.ProfilerInitOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ProfilerResource = js.String

type ProfilerFrame struct {
	// Name is "ProfilerFrame.name"
	//
	// Required
	Name js.String
	// ResourceId is "ProfilerFrame.resourceId"
	//
	// Optional
	ResourceId uint64
	// Line is "ProfilerFrame.line"
	//
	// Optional
	Line uint64
	// Column is "ProfilerFrame.column"
	//
	// Optional
	Column uint64

	FFI_USE_ResourceId bool // for ResourceId.
	FFI_USE_Line       bool // for Line.
	FFI_USE_Column     bool // for Column.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProfilerFrame with all fields set.
func (p ProfilerFrame) FromRef(ref js.Ref) ProfilerFrame {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ProfilerFrame ProfilerFrame [// ProfilerFrame] [0x14000960f00 0x14000960fa0 0x140009610e0 0x14000961220 0x14000961040 0x14000961180 0x140009612c0] 0x140009210b0 {0 0}} in the application heap.
func (p ProfilerFrame) New() js.Ref {
	return bindings.ProfilerFrameJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ProfilerFrame) UpdateFrom(ref js.Ref) {
	bindings.ProfilerFrameJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ProfilerFrame) Update(ref js.Ref) {
	bindings.ProfilerFrameJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ProfilerStack struct {
	// ParentId is "ProfilerStack.parentId"
	//
	// Optional
	ParentId uint64
	// FrameId is "ProfilerStack.frameId"
	//
	// Required
	FrameId uint64

	FFI_USE_ParentId bool // for ParentId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProfilerStack with all fields set.
func (p ProfilerStack) FromRef(ref js.Ref) ProfilerStack {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ProfilerStack ProfilerStack [// ProfilerStack] [0x14000961400 0x14000961540 0x140009614a0] 0x14000921110 {0 0}} in the application heap.
func (p ProfilerStack) New() js.Ref {
	return bindings.ProfilerStackJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ProfilerStack) UpdateFrom(ref js.Ref) {
	bindings.ProfilerStackJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ProfilerStack) Update(ref js.Ref) {
	bindings.ProfilerStackJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ProfilerSample struct {
	// Timestamp is "ProfilerSample.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// StackId is "ProfilerSample.stackId"
	//
	// Optional
	StackId uint64

	FFI_USE_StackId bool // for StackId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProfilerSample with all fields set.
func (p ProfilerSample) FromRef(ref js.Ref) ProfilerSample {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ProfilerSample ProfilerSample [// ProfilerSample] [0x14000961680 0x14000961720 0x140009617c0] 0x14000921140 {0 0}} in the application heap.
func (p ProfilerSample) New() js.Ref {
	return bindings.ProfilerSampleJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ProfilerSample) UpdateFrom(ref js.Ref) {
	bindings.ProfilerSampleJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ProfilerSample) Update(ref js.Ref) {
	bindings.ProfilerSampleJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ProfilerTrace struct {
	// Resources is "ProfilerTrace.resources"
	//
	// Required
	Resources js.Array[ProfilerResource]
	// Frames is "ProfilerTrace.frames"
	//
	// Required
	Frames js.Array[ProfilerFrame]
	// Stacks is "ProfilerTrace.stacks"
	//
	// Required
	Stacks js.Array[ProfilerStack]
	// Samples is "ProfilerTrace.samples"
	//
	// Required
	Samples js.Array[ProfilerSample]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProfilerTrace with all fields set.
func (p ProfilerTrace) FromRef(ref js.Ref) ProfilerTrace {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ProfilerTrace ProfilerTrace [// ProfilerTrace] [0x14000960e60 0x14000961360 0x140009615e0 0x14000961860] 0x14000921098 {0 0}} in the application heap.
func (p ProfilerTrace) New() js.Ref {
	return bindings.ProfilerTraceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ProfilerTrace) UpdateFrom(ref js.Ref) {
	bindings.ProfilerTraceJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ProfilerTrace) Update(ref js.Ref) {
	bindings.ProfilerTraceJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewProfiler(options ProfilerInitOptions) Profiler {
	return Profiler{}.FromRef(
		bindings.NewProfilerByProfiler(
			js.Pointer(&options)),
	)
}

type Profiler struct {
	EventTarget
}

func (this Profiler) Once() Profiler {
	this.Ref().Once()
	return this
}

func (this Profiler) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Profiler) FromRef(ref js.Ref) Profiler {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Profiler) Free() {
	this.Ref().Free()
}

// SampleInterval returns the value of property "Profiler.sampleInterval".
//
// The returned bool will be false if there is no such property.
func (this Profiler) SampleInterval() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetProfilerSampleInterval(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// Stopped returns the value of property "Profiler.stopped".
//
// The returned bool will be false if there is no such property.
func (this Profiler) Stopped() (bool, bool) {
	var _ok bool
	_ret := bindings.GetProfilerStopped(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Stop calls the method "Profiler.stop".
//
// The returned bool will be false if there is no such method.
func (this Profiler) Stop() (js.Promise[ProfilerTrace], bool) {
	var _ok bool
	_ret := bindings.CallProfilerStop(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[ProfilerTrace]{}.FromRef(_ret), _ok
}

// StopFunc returns the method "Profiler.stop".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Profiler) StopFunc() (fn js.Func[func() js.Promise[ProfilerTrace]]) {
	return fn.FromRef(
		bindings.ProfilerStopFunc(
			this.Ref(),
		),
	)
}

type ProgressEventInit struct {
	// LengthComputable is "ProgressEventInit.lengthComputable"
	//
	// Optional, defaults to false.
	LengthComputable bool
	// Loaded is "ProgressEventInit.loaded"
	//
	// Optional, defaults to 0.
	Loaded uint64
	// Total is "ProgressEventInit.total"
	//
	// Optional, defaults to 0.
	Total uint64
	// Bubbles is "ProgressEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "ProgressEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "ProgressEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_LengthComputable bool // for LengthComputable.
	FFI_USE_Loaded           bool // for Loaded.
	FFI_USE_Total            bool // for Total.
	FFI_USE_Bubbles          bool // for Bubbles.
	FFI_USE_Cancelable       bool // for Cancelable.
	FFI_USE_Composed         bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProgressEventInit with all fields set.
func (p ProgressEventInit) FromRef(ref js.Ref) ProgressEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ProgressEventInit ProgressEventInit [// ProgressEventInit] [0x140009619a0 0x14000961ae0 0x14000961c20 0x14000961d60 0x14000961ea0 0x14000966000 0x14000961a40 0x14000961b80 0x14000961cc0 0x14000961e00 0x14000961f40 0x140009660a0] 0x14000921188 {0 0}} in the application heap.
func (p ProgressEventInit) New() js.Ref {
	return bindings.ProgressEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ProgressEventInit) UpdateFrom(ref js.Ref) {
	bindings.ProgressEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ProgressEventInit) Update(ref js.Ref) {
	bindings.ProgressEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewProgressEvent(typ js.String, eventInitDict ProgressEventInit) ProgressEvent {
	return ProgressEvent{}.FromRef(
		bindings.NewProgressEventByProgressEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewProgressEventByProgressEvent1(typ js.String) ProgressEvent {
	return ProgressEvent{}.FromRef(
		bindings.NewProgressEventByProgressEvent1(
			typ.Ref()),
	)
}

type ProgressEvent struct {
	Event
}

func (this ProgressEvent) Once() ProgressEvent {
	this.Ref().Once()
	return this
}

func (this ProgressEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this ProgressEvent) FromRef(ref js.Ref) ProgressEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this ProgressEvent) Free() {
	this.Ref().Free()
}

// LengthComputable returns the value of property "ProgressEvent.lengthComputable".
//
// The returned bool will be false if there is no such property.
func (this ProgressEvent) LengthComputable() (bool, bool) {
	var _ok bool
	_ret := bindings.GetProgressEventLengthComputable(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Loaded returns the value of property "ProgressEvent.loaded".
//
// The returned bool will be false if there is no such property.
func (this ProgressEvent) Loaded() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetProgressEventLoaded(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// Total returns the value of property "ProgressEvent.total".
//
// The returned bool will be false if there is no such property.
func (this ProgressEvent) Total() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetProgressEventTotal(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

type PromiseRejectionEventInit struct {
	// Promise is "PromiseRejectionEventInit.promise"
	//
	// Required
	Promise js.Promise[js.Any]
	// Reason is "PromiseRejectionEventInit.reason"
	//
	// Optional
	Reason js.Any
	// Bubbles is "PromiseRejectionEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "PromiseRejectionEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PromiseRejectionEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PromiseRejectionEventInit with all fields set.
func (p PromiseRejectionEventInit) FromRef(ref js.Ref) PromiseRejectionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PromiseRejectionEventInit PromiseRejectionEventInit [// PromiseRejectionEventInit] [0x140009661e0 0x14000966280 0x14000966320 0x14000966460 0x140009665a0 0x140009663c0 0x14000966500 0x14000966640] 0x140009211d0 {0 0}} in the application heap.
func (p PromiseRejectionEventInit) New() js.Ref {
	return bindings.PromiseRejectionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PromiseRejectionEventInit) UpdateFrom(ref js.Ref) {
	bindings.PromiseRejectionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PromiseRejectionEventInit) Update(ref js.Ref) {
	bindings.PromiseRejectionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPromiseRejectionEvent(typ js.String, eventInitDict PromiseRejectionEventInit) PromiseRejectionEvent {
	return PromiseRejectionEvent{}.FromRef(
		bindings.NewPromiseRejectionEventByPromiseRejectionEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type PromiseRejectionEvent struct {
	Event
}

func (this PromiseRejectionEvent) Once() PromiseRejectionEvent {
	this.Ref().Once()
	return this
}

func (this PromiseRejectionEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this PromiseRejectionEvent) FromRef(ref js.Ref) PromiseRejectionEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this PromiseRejectionEvent) Free() {
	this.Ref().Free()
}

// Promise returns the value of property "PromiseRejectionEvent.promise".
//
// The returned bool will be false if there is no such property.
func (this PromiseRejectionEvent) Promise() (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.GetPromiseRejectionEventPromise(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// Reason returns the value of property "PromiseRejectionEvent.reason".
//
// The returned bool will be false if there is no such property.
func (this PromiseRejectionEvent) Reason() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetPromiseRejectionEventReason(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

type ProximityReadingValues struct {
	// Distance is "ProximityReadingValues.distance"
	//
	// Required
	Distance float64
	// Max is "ProximityReadingValues.max"
	//
	// Required
	Max float64
	// Near is "ProximityReadingValues.near"
	//
	// Required
	Near bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProximityReadingValues with all fields set.
func (p ProximityReadingValues) FromRef(ref js.Ref) ProximityReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ProximityReadingValues ProximityReadingValues [// ProximityReadingValues] [0x14000966780 0x14000966820 0x140009668c0] 0x14000921200 {0 0}} in the application heap.
func (p ProximityReadingValues) New() js.Ref {
	return bindings.ProximityReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ProximityReadingValues) UpdateFrom(ref js.Ref) {
	bindings.ProximityReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ProximityReadingValues) Update(ref js.Ref) {
	bindings.ProximityReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewProximitySensor(sensorOptions SensorOptions) ProximitySensor {
	return ProximitySensor{}.FromRef(
		bindings.NewProximitySensorByProximitySensor(
			js.Pointer(&sensorOptions)),
	)
}

func NewProximitySensorByProximitySensor1() ProximitySensor {
	return ProximitySensor{}.FromRef(
		bindings.NewProximitySensorByProximitySensor1(),
	)
}

type ProximitySensor struct {
	Sensor
}

func (this ProximitySensor) Once() ProximitySensor {
	this.Ref().Once()
	return this
}

func (this ProximitySensor) Ref() js.Ref {
	return this.Sensor.Ref()
}

func (this ProximitySensor) FromRef(ref js.Ref) ProximitySensor {
	this.Sensor = this.Sensor.FromRef(ref)
	return this
}

func (this ProximitySensor) Free() {
	this.Ref().Free()
}

// Distance returns the value of property "ProximitySensor.distance".
//
// The returned bool will be false if there is no such property.
func (this ProximitySensor) Distance() (float64, bool) {
	var _ok bool
	_ret := bindings.GetProximitySensorDistance(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Max returns the value of property "ProximitySensor.max".
//
// The returned bool will be false if there is no such property.
func (this ProximitySensor) Max() (float64, bool) {
	var _ok bool
	_ret := bindings.GetProximitySensorMax(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Near returns the value of property "ProximitySensor.near".
//
// The returned bool will be false if there is no such property.
func (this ProximitySensor) Near() (bool, bool) {
	var _ok bool
	_ret := bindings.GetProximitySensorNear(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type RegistrationResponseJSON struct {
	// Id is "RegistrationResponseJSON.id"
	//
	// Required
	Id Base64URLString
	// RawId is "RegistrationResponseJSON.rawId"
	//
	// Required
	RawId Base64URLString
	// Response is "RegistrationResponseJSON.response"
	//
	// Required
	Response AuthenticatorAttestationResponseJSON
	// AuthenticatorAttachment is "RegistrationResponseJSON.authenticatorAttachment"
	//
	// Optional
	AuthenticatorAttachment js.String
	// ClientExtensionResults is "RegistrationResponseJSON.clientExtensionResults"
	//
	// Required
	ClientExtensionResults AuthenticationExtensionsClientOutputsJSON
	// Type is "RegistrationResponseJSON.type"
	//
	// Required
	Type js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RegistrationResponseJSON with all fields set.
func (p RegistrationResponseJSON) FromRef(ref js.Ref) RegistrationResponseJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RegistrationResponseJSON RegistrationResponseJSON [// RegistrationResponseJSON] [0x14000966a00 0x14000966aa0 0x14000966b40 0x14000966be0 0x14000966c80 0x14000966d20] 0x14000921260 {0 0}} in the application heap.
func (p RegistrationResponseJSON) New() js.Ref {
	return bindings.RegistrationResponseJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RegistrationResponseJSON) UpdateFrom(ref js.Ref) {
	bindings.RegistrationResponseJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RegistrationResponseJSON) Update(ref js.Ref) {
	bindings.RegistrationResponseJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_RegistrationResponseJSON_AuthenticationResponseJSON struct {
	ref js.Ref
}

func (x OneOf_RegistrationResponseJSON_AuthenticationResponseJSON) Ref() js.Ref {
	return x.ref
}

func (x OneOf_RegistrationResponseJSON_AuthenticationResponseJSON) Free() {
	x.ref.Free()
}

func (x OneOf_RegistrationResponseJSON_AuthenticationResponseJSON) FromRef(ref js.Ref) OneOf_RegistrationResponseJSON_AuthenticationResponseJSON {
	return OneOf_RegistrationResponseJSON_AuthenticationResponseJSON{
		ref: ref,
	}
}

func (x OneOf_RegistrationResponseJSON_AuthenticationResponseJSON) RegistrationResponseJSON() RegistrationResponseJSON {
	var ret RegistrationResponseJSON
	ret.UpdateFrom(x.ref)
	return ret
}

func (x OneOf_RegistrationResponseJSON_AuthenticationResponseJSON) AuthenticationResponseJSON() AuthenticationResponseJSON {
	var ret AuthenticationResponseJSON
	ret.UpdateFrom(x.ref)
	return ret
}

type PublicKeyCredentialJSON = OneOf_RegistrationResponseJSON_AuthenticationResponseJSON

type PublicKeyCredentialDescriptorJSON struct {
	// Id is "PublicKeyCredentialDescriptorJSON.id"
	//
	// Required
	Id Base64URLString
	// Type is "PublicKeyCredentialDescriptorJSON.type"
	//
	// Required
	Type js.String
	// Transports is "PublicKeyCredentialDescriptorJSON.transports"
	//
	// Optional
	Transports js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialDescriptorJSON with all fields set.
func (p PublicKeyCredentialDescriptorJSON) FromRef(ref js.Ref) PublicKeyCredentialDescriptorJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PublicKeyCredentialDescriptorJSON PublicKeyCredentialDescriptorJSON [// PublicKeyCredentialDescriptorJSON] [0x14000967040 0x140009670e0 0x14000967180] 0x140009212d8 {0 0}} in the application heap.
func (p PublicKeyCredentialDescriptorJSON) New() js.Ref {
	return bindings.PublicKeyCredentialDescriptorJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PublicKeyCredentialDescriptorJSON) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialDescriptorJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PublicKeyCredentialDescriptorJSON) Update(ref js.Ref) {
	bindings.PublicKeyCredentialDescriptorJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PublicKeyCredentialRequestOptionsJSON struct {
	// Challenge is "PublicKeyCredentialRequestOptionsJSON.challenge"
	//
	// Required
	Challenge Base64URLString
	// Timeout is "PublicKeyCredentialRequestOptionsJSON.timeout"
	//
	// Optional
	Timeout uint32
	// RpId is "PublicKeyCredentialRequestOptionsJSON.rpId"
	//
	// Optional
	RpId js.String
	// AllowCredentials is "PublicKeyCredentialRequestOptionsJSON.allowCredentials"
	//
	// Optional, defaults to [].
	AllowCredentials js.Array[PublicKeyCredentialDescriptorJSON]
	// UserVerification is "PublicKeyCredentialRequestOptionsJSON.userVerification"
	//
	// Optional, defaults to "preferred".
	UserVerification js.String
	// Hints is "PublicKeyCredentialRequestOptionsJSON.hints"
	//
	// Optional, defaults to [].
	Hints js.Array[js.String]
	// Attestation is "PublicKeyCredentialRequestOptionsJSON.attestation"
	//
	// Optional, defaults to "none".
	Attestation js.String
	// AttestationFormats is "PublicKeyCredentialRequestOptionsJSON.attestationFormats"
	//
	// Optional, defaults to [].
	AttestationFormats js.Array[js.String]
	// Extensions is "PublicKeyCredentialRequestOptionsJSON.extensions"
	//
	// Optional
	Extensions AuthenticationExtensionsClientInputsJSON

	FFI_USE_Timeout bool // for Timeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialRequestOptionsJSON with all fields set.
func (p PublicKeyCredentialRequestOptionsJSON) FromRef(ref js.Ref) PublicKeyCredentialRequestOptionsJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PublicKeyCredentialRequestOptionsJSON PublicKeyCredentialRequestOptionsJSON [// PublicKeyCredentialRequestOptionsJSON] [0x14000966dc0 0x14000966e60 0x14000966fa0 0x14000967220 0x140009672c0 0x14000967360 0x14000967400 0x140009674a0 0x14000967540 0x14000966f00] 0x140009212a8 {0 0}} in the application heap.
func (p PublicKeyCredentialRequestOptionsJSON) New() js.Ref {
	return bindings.PublicKeyCredentialRequestOptionsJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PublicKeyCredentialRequestOptionsJSON) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialRequestOptionsJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PublicKeyCredentialRequestOptionsJSON) Update(ref js.Ref) {
	bindings.PublicKeyCredentialRequestOptionsJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PublicKeyCredentialUserEntityJSON struct {
	// Id is "PublicKeyCredentialUserEntityJSON.id"
	//
	// Required
	Id Base64URLString
	// Name is "PublicKeyCredentialUserEntityJSON.name"
	//
	// Required
	Name js.String
	// DisplayName is "PublicKeyCredentialUserEntityJSON.displayName"
	//
	// Required
	DisplayName js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialUserEntityJSON with all fields set.
func (p PublicKeyCredentialUserEntityJSON) FromRef(ref js.Ref) PublicKeyCredentialUserEntityJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PublicKeyCredentialUserEntityJSON PublicKeyCredentialUserEntityJSON [// PublicKeyCredentialUserEntityJSON] [0x14000967680 0x14000967720 0x140009677c0] 0x14000921368 {0 0}} in the application heap.
func (p PublicKeyCredentialUserEntityJSON) New() js.Ref {
	return bindings.PublicKeyCredentialUserEntityJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PublicKeyCredentialUserEntityJSON) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialUserEntityJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PublicKeyCredentialUserEntityJSON) Update(ref js.Ref) {
	bindings.PublicKeyCredentialUserEntityJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PublicKeyCredentialCreationOptionsJSON struct {
	// Rp is "PublicKeyCredentialCreationOptionsJSON.rp"
	//
	// Required
	Rp PublicKeyCredentialRpEntity
	// User is "PublicKeyCredentialCreationOptionsJSON.user"
	//
	// Required
	User PublicKeyCredentialUserEntityJSON
	// Challenge is "PublicKeyCredentialCreationOptionsJSON.challenge"
	//
	// Required
	Challenge Base64URLString
	// PubKeyCredParams is "PublicKeyCredentialCreationOptionsJSON.pubKeyCredParams"
	//
	// Required
	PubKeyCredParams js.Array[PublicKeyCredentialParameters]
	// Timeout is "PublicKeyCredentialCreationOptionsJSON.timeout"
	//
	// Optional
	Timeout uint32
	// ExcludeCredentials is "PublicKeyCredentialCreationOptionsJSON.excludeCredentials"
	//
	// Optional, defaults to [].
	ExcludeCredentials js.Array[PublicKeyCredentialDescriptorJSON]
	// AuthenticatorSelection is "PublicKeyCredentialCreationOptionsJSON.authenticatorSelection"
	//
	// Optional
	AuthenticatorSelection AuthenticatorSelectionCriteria
	// Hints is "PublicKeyCredentialCreationOptionsJSON.hints"
	//
	// Optional, defaults to [].
	Hints js.Array[js.String]
	// Attestation is "PublicKeyCredentialCreationOptionsJSON.attestation"
	//
	// Optional, defaults to "none".
	Attestation js.String
	// AttestationFormats is "PublicKeyCredentialCreationOptionsJSON.attestationFormats"
	//
	// Optional, defaults to [].
	AttestationFormats js.Array[js.String]
	// Extensions is "PublicKeyCredentialCreationOptionsJSON.extensions"
	//
	// Optional
	Extensions AuthenticationExtensionsClientInputsJSON

	FFI_USE_Timeout bool // for Timeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialCreationOptionsJSON with all fields set.
func (p PublicKeyCredentialCreationOptionsJSON) FromRef(ref js.Ref) PublicKeyCredentialCreationOptionsJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PublicKeyCredentialCreationOptionsJSON PublicKeyCredentialCreationOptionsJSON [// PublicKeyCredentialCreationOptionsJSON] [0x140009675e0 0x14000967860 0x14000967900 0x140009679a0 0x14000967a40 0x14000967b80 0x14000967c20 0x14000967cc0 0x14000967d60 0x14000967e00 0x14000967ea0 0x14000967ae0] 0x14000921320 {0 0}} in the application heap.
func (p PublicKeyCredentialCreationOptionsJSON) New() js.Ref {
	return bindings.PublicKeyCredentialCreationOptionsJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PublicKeyCredentialCreationOptionsJSON) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialCreationOptionsJSONJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PublicKeyCredentialCreationOptionsJSON) Update(ref js.Ref) {
	bindings.PublicKeyCredentialCreationOptionsJSONJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PublicKeyCredential struct {
	Credential
}

func (this PublicKeyCredential) Once() PublicKeyCredential {
	this.Ref().Once()
	return this
}

func (this PublicKeyCredential) Ref() js.Ref {
	return this.Credential.Ref()
}

func (this PublicKeyCredential) FromRef(ref js.Ref) PublicKeyCredential {
	this.Credential = this.Credential.FromRef(ref)
	return this
}

func (this PublicKeyCredential) Free() {
	this.Ref().Free()
}

// RawId returns the value of property "PublicKeyCredential.rawId".
//
// The returned bool will be false if there is no such property.
func (this PublicKeyCredential) RawId() (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.GetPublicKeyCredentialRawId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

// Response returns the value of property "PublicKeyCredential.response".
//
// The returned bool will be false if there is no such property.
func (this PublicKeyCredential) Response() (AuthenticatorResponse, bool) {
	var _ok bool
	_ret := bindings.GetPublicKeyCredentialResponse(
		this.Ref(), js.Pointer(&_ok),
	)
	return AuthenticatorResponse{}.FromRef(_ret), _ok
}

// AuthenticatorAttachment returns the value of property "PublicKeyCredential.authenticatorAttachment".
//
// The returned bool will be false if there is no such property.
func (this PublicKeyCredential) AuthenticatorAttachment() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPublicKeyCredentialAuthenticatorAttachment(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// GetClientExtensionResults calls the method "PublicKeyCredential.getClientExtensionResults".
//
// The returned bool will be false if there is no such method.
func (this PublicKeyCredential) GetClientExtensionResults() (AuthenticationExtensionsClientOutputs, bool) {
	var _ret AuthenticationExtensionsClientOutputs
	_ok := js.True == bindings.CallPublicKeyCredentialGetClientExtensionResults(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetClientExtensionResultsFunc returns the method "PublicKeyCredential.getClientExtensionResults".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PublicKeyCredential) GetClientExtensionResultsFunc() (fn js.Func[func() AuthenticationExtensionsClientOutputs]) {
	return fn.FromRef(
		bindings.PublicKeyCredentialGetClientExtensionResultsFunc(
			this.Ref(),
		),
	)
}

// IsConditionalMediationAvailable calls the staticmethod "PublicKeyCredential.isConditionalMediationAvailable".
//
// The returned bool will be false if there is no such method.
func (this PublicKeyCredential) IsConditionalMediationAvailable() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallPublicKeyCredentialIsConditionalMediationAvailable(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// IsConditionalMediationAvailableFunc returns the staticmethod "PublicKeyCredential.isConditionalMediationAvailable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PublicKeyCredential) IsConditionalMediationAvailableFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.PublicKeyCredentialIsConditionalMediationAvailableFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PublicKeyCredential.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PublicKeyCredential) ToJSON() (PublicKeyCredentialJSON, bool) {
	var _ok bool
	_ret := bindings.CallPublicKeyCredentialToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return PublicKeyCredentialJSON{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PublicKeyCredential.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PublicKeyCredential) ToJSONFunc() (fn js.Func[func() PublicKeyCredentialJSON]) {
	return fn.FromRef(
		bindings.PublicKeyCredentialToJSONFunc(
			this.Ref(),
		),
	)
}

// ParseRequestOptionsFromJSON calls the staticmethod "PublicKeyCredential.parseRequestOptionsFromJSON".
//
// The returned bool will be false if there is no such method.
func (this PublicKeyCredential) ParseRequestOptionsFromJSON(options PublicKeyCredentialRequestOptionsJSON) (PublicKeyCredentialRequestOptions, bool) {
	var _ret PublicKeyCredentialRequestOptions
	_ok := js.True == bindings.CallPublicKeyCredentialParseRequestOptionsFromJSON(
		this.Ref(), js.Pointer(&_ret),
		js.Pointer(&options),
	)

	return _ret, _ok
}

// ParseRequestOptionsFromJSONFunc returns the staticmethod "PublicKeyCredential.parseRequestOptionsFromJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PublicKeyCredential) ParseRequestOptionsFromJSONFunc() (fn js.Func[func(options PublicKeyCredentialRequestOptionsJSON) PublicKeyCredentialRequestOptions]) {
	return fn.FromRef(
		bindings.PublicKeyCredentialParseRequestOptionsFromJSONFunc(
			this.Ref(),
		),
	)
}

// IsPasskeyPlatformAuthenticatorAvailable calls the staticmethod "PublicKeyCredential.isPasskeyPlatformAuthenticatorAvailable".
//
// The returned bool will be false if there is no such method.
func (this PublicKeyCredential) IsPasskeyPlatformAuthenticatorAvailable() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallPublicKeyCredentialIsPasskeyPlatformAuthenticatorAvailable(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// IsPasskeyPlatformAuthenticatorAvailableFunc returns the staticmethod "PublicKeyCredential.isPasskeyPlatformAuthenticatorAvailable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PublicKeyCredential) IsPasskeyPlatformAuthenticatorAvailableFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.PublicKeyCredentialIsPasskeyPlatformAuthenticatorAvailableFunc(
			this.Ref(),
		),
	)
}

// ParseCreationOptionsFromJSON calls the staticmethod "PublicKeyCredential.parseCreationOptionsFromJSON".
//
// The returned bool will be false if there is no such method.
func (this PublicKeyCredential) ParseCreationOptionsFromJSON(options PublicKeyCredentialCreationOptionsJSON) (PublicKeyCredentialCreationOptions, bool) {
	var _ret PublicKeyCredentialCreationOptions
	_ok := js.True == bindings.CallPublicKeyCredentialParseCreationOptionsFromJSON(
		this.Ref(), js.Pointer(&_ret),
		js.Pointer(&options),
	)

	return _ret, _ok
}

// ParseCreationOptionsFromJSONFunc returns the staticmethod "PublicKeyCredential.parseCreationOptionsFromJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PublicKeyCredential) ParseCreationOptionsFromJSONFunc() (fn js.Func[func(options PublicKeyCredentialCreationOptionsJSON) PublicKeyCredentialCreationOptions]) {
	return fn.FromRef(
		bindings.PublicKeyCredentialParseCreationOptionsFromJSONFunc(
			this.Ref(),
		),
	)
}

// IsUserVerifyingPlatformAuthenticatorAvailable calls the staticmethod "PublicKeyCredential.isUserVerifyingPlatformAuthenticatorAvailable".
//
// The returned bool will be false if there is no such method.
func (this PublicKeyCredential) IsUserVerifyingPlatformAuthenticatorAvailable() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallPublicKeyCredentialIsUserVerifyingPlatformAuthenticatorAvailable(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// IsUserVerifyingPlatformAuthenticatorAvailableFunc returns the staticmethod "PublicKeyCredential.isUserVerifyingPlatformAuthenticatorAvailable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PublicKeyCredential) IsUserVerifyingPlatformAuthenticatorAvailableFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.PublicKeyCredentialIsUserVerifyingPlatformAuthenticatorAvailableFunc(
			this.Ref(),
		),
	)
}

type PublicKeyCredentialEntity struct {
	// Name is "PublicKeyCredentialEntity.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialEntity with all fields set.
func (p PublicKeyCredentialEntity) FromRef(ref js.Ref) PublicKeyCredentialEntity {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PublicKeyCredentialEntity PublicKeyCredentialEntity [// PublicKeyCredentialEntity] [0x14000972000] 0x140009213b0 {0 0}} in the application heap.
func (p PublicKeyCredentialEntity) New() js.Ref {
	return bindings.PublicKeyCredentialEntityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PublicKeyCredentialEntity) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialEntityJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PublicKeyCredentialEntity) Update(ref js.Ref) {
	bindings.PublicKeyCredentialEntityJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PublicKeyCredentialHints uint32

const (
	_ PublicKeyCredentialHints = iota

	PublicKeyCredentialHints_SECURITY_KEY
	PublicKeyCredentialHints_CLIENT_DEVICE
	PublicKeyCredentialHints_HYBRID
)

func (PublicKeyCredentialHints) FromRef(str js.Ref) PublicKeyCredentialHints {
	return PublicKeyCredentialHints(bindings.ConstOfPublicKeyCredentialHints(str))
}

func (x PublicKeyCredentialHints) String() (string, bool) {
	switch x {
	case PublicKeyCredentialHints_SECURITY_KEY:
		return "security-key", true
	case PublicKeyCredentialHints_CLIENT_DEVICE:
		return "client-device", true
	case PublicKeyCredentialHints_HYBRID:
		return "hybrid", true
	default:
		return "", false
	}
}

type PublicKeyCredentialType uint32

const (
	_ PublicKeyCredentialType = iota

	PublicKeyCredentialType_PUBLIC_KEY
)

func (PublicKeyCredentialType) FromRef(str js.Ref) PublicKeyCredentialType {
	return PublicKeyCredentialType(bindings.ConstOfPublicKeyCredentialType(str))
}

func (x PublicKeyCredentialType) String() (string, bool) {
	switch x {
	case PublicKeyCredentialType_PUBLIC_KEY:
		return "public-key", true
	default:
		return "", false
	}
}

type PushMessageDataInit = OneOf_TypedArrayInt8_TypedArrayInt16_TypedArrayInt32_TypedArrayUint8_TypedArrayUint16_TypedArrayUint32_TypedArrayInt64_TypedArrayUint64_TypedArrayFloat32_TypedArrayFloat64_DataView_ArrayBuffer_String

type PushEventInit struct {
	// Data is "PushEventInit.data"
	//
	// Optional
	Data PushMessageDataInit
	// Bubbles is "PushEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "PushEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PushEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PushEventInit with all fields set.
func (p PushEventInit) FromRef(ref js.Ref) PushEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PushEventInit PushEventInit [// PushEventInit] [0x140009720a0 0x14000972140 0x14000972280 0x140009723c0 0x140009721e0 0x14000972320 0x14000972460] 0x14000921458 {0 0}} in the application heap.
func (p PushEventInit) New() js.Ref {
	return bindings.PushEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PushEventInit) UpdateFrom(ref js.Ref) {
	bindings.PushEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PushEventInit) Update(ref js.Ref) {
	bindings.PushEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PushMessageData struct {
	ref js.Ref
}

func (this PushMessageData) Once() PushMessageData {
	this.Ref().Once()
	return this
}

func (this PushMessageData) Ref() js.Ref {
	return this.ref
}

func (this PushMessageData) FromRef(ref js.Ref) PushMessageData {
	this.ref = ref
	return this
}

func (this PushMessageData) Free() {
	this.Ref().Free()
}

// ArrayBuffer calls the method "PushMessageData.arrayBuffer".
//
// The returned bool will be false if there is no such method.
func (this PushMessageData) ArrayBuffer() (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.CallPushMessageDataArrayBuffer(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

// ArrayBufferFunc returns the method "PushMessageData.arrayBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushMessageData) ArrayBufferFunc() (fn js.Func[func() js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.PushMessageDataArrayBufferFunc(
			this.Ref(),
		),
	)
}

// Blob calls the method "PushMessageData.blob".
//
// The returned bool will be false if there is no such method.
func (this PushMessageData) Blob() (Blob, bool) {
	var _ok bool
	_ret := bindings.CallPushMessageDataBlob(
		this.Ref(), js.Pointer(&_ok),
	)

	return Blob{}.FromRef(_ret), _ok
}

// BlobFunc returns the method "PushMessageData.blob".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushMessageData) BlobFunc() (fn js.Func[func() Blob]) {
	return fn.FromRef(
		bindings.PushMessageDataBlobFunc(
			this.Ref(),
		),
	)
}

// Json calls the method "PushMessageData.json".
//
// The returned bool will be false if there is no such method.
func (this PushMessageData) Json() (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallPushMessageDataJson(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// JsonFunc returns the method "PushMessageData.json".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushMessageData) JsonFunc() (fn js.Func[func() js.Any]) {
	return fn.FromRef(
		bindings.PushMessageDataJsonFunc(
			this.Ref(),
		),
	)
}

// Text calls the method "PushMessageData.text".
//
// The returned bool will be false if there is no such method.
func (this PushMessageData) Text() (js.String, bool) {
	var _ok bool
	_ret := bindings.CallPushMessageDataText(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.String{}.FromRef(_ret), _ok
}

// TextFunc returns the method "PushMessageData.text".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PushMessageData) TextFunc() (fn js.Func[func() js.String]) {
	return fn.FromRef(
		bindings.PushMessageDataTextFunc(
			this.Ref(),
		),
	)
}

func NewPushEvent(typ js.String, eventInitDict PushEventInit) PushEvent {
	return PushEvent{}.FromRef(
		bindings.NewPushEventByPushEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewPushEventByPushEvent1(typ js.String) PushEvent {
	return PushEvent{}.FromRef(
		bindings.NewPushEventByPushEvent1(
			typ.Ref()),
	)
}

type PushEvent struct {
	ExtendableEvent
}

func (this PushEvent) Once() PushEvent {
	this.Ref().Once()
	return this
}

func (this PushEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this PushEvent) FromRef(ref js.Ref) PushEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this PushEvent) Free() {
	this.Ref().Free()
}

// Data returns the value of property "PushEvent.data".
//
// The returned bool will be false if there is no such property.
func (this PushEvent) Data() (PushMessageData, bool) {
	var _ok bool
	_ret := bindings.GetPushEventData(
		this.Ref(), js.Pointer(&_ok),
	)
	return PushMessageData{}.FromRef(_ret), _ok
}

type PushPermissionDescriptor struct {
	// UserVisibleOnly is "PushPermissionDescriptor.userVisibleOnly"
	//
	// Optional, defaults to false.
	UserVisibleOnly bool
	// Name is "PushPermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE_UserVisibleOnly bool // for UserVisibleOnly.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PushPermissionDescriptor with all fields set.
func (p PushPermissionDescriptor) FromRef(ref js.Ref) PushPermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PushPermissionDescriptor PushPermissionDescriptor [// PushPermissionDescriptor] [0x14000972500 0x14000972640 0x140009725a0] 0x14000921578 {0 0}} in the application heap.
func (p PushPermissionDescriptor) New() js.Ref {
	return bindings.PushPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PushPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.PushPermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PushPermissionDescriptor) Update(ref js.Ref) {
	bindings.PushPermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PushSubscriptionChangeEventInit struct {
	// NewSubscription is "PushSubscriptionChangeEventInit.newSubscription"
	//
	// Optional, defaults to null.
	NewSubscription PushSubscription
	// OldSubscription is "PushSubscriptionChangeEventInit.oldSubscription"
	//
	// Optional, defaults to null.
	OldSubscription PushSubscription
	// Bubbles is "PushSubscriptionChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "PushSubscriptionChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PushSubscriptionChangeEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PushSubscriptionChangeEventInit with all fields set.
func (p PushSubscriptionChangeEventInit) FromRef(ref js.Ref) PushSubscriptionChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PushSubscriptionChangeEventInit PushSubscriptionChangeEventInit [// PushSubscriptionChangeEventInit] [0x140009726e0 0x14000972780 0x14000972820 0x14000972960 0x14000972aa0 0x140009728c0 0x14000972a00 0x14000972b40] 0x140009215d8 {0 0}} in the application heap.
func (p PushSubscriptionChangeEventInit) New() js.Ref {
	return bindings.PushSubscriptionChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PushSubscriptionChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.PushSubscriptionChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PushSubscriptionChangeEventInit) Update(ref js.Ref) {
	bindings.PushSubscriptionChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPushSubscriptionChangeEvent(typ js.String, eventInitDict PushSubscriptionChangeEventInit) PushSubscriptionChangeEvent {
	return PushSubscriptionChangeEvent{}.FromRef(
		bindings.NewPushSubscriptionChangeEventByPushSubscriptionChangeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewPushSubscriptionChangeEventByPushSubscriptionChangeEvent1(typ js.String) PushSubscriptionChangeEvent {
	return PushSubscriptionChangeEvent{}.FromRef(
		bindings.NewPushSubscriptionChangeEventByPushSubscriptionChangeEvent1(
			typ.Ref()),
	)
}

type PushSubscriptionChangeEvent struct {
	ExtendableEvent
}

func (this PushSubscriptionChangeEvent) Once() PushSubscriptionChangeEvent {
	this.Ref().Once()
	return this
}

func (this PushSubscriptionChangeEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this PushSubscriptionChangeEvent) FromRef(ref js.Ref) PushSubscriptionChangeEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this PushSubscriptionChangeEvent) Free() {
	this.Ref().Free()
}

// NewSubscription returns the value of property "PushSubscriptionChangeEvent.newSubscription".
//
// The returned bool will be false if there is no such property.
func (this PushSubscriptionChangeEvent) NewSubscription() (PushSubscription, bool) {
	var _ok bool
	_ret := bindings.GetPushSubscriptionChangeEventNewSubscription(
		this.Ref(), js.Pointer(&_ok),
	)
	return PushSubscription{}.FromRef(_ret), _ok
}

// OldSubscription returns the value of property "PushSubscriptionChangeEvent.oldSubscription".
//
// The returned bool will be false if there is no such property.
func (this PushSubscriptionChangeEvent) OldSubscription() (PushSubscription, bool) {
	var _ok bool
	_ret := bindings.GetPushSubscriptionChangeEventOldSubscription(
		this.Ref(), js.Pointer(&_ok),
	)
	return PushSubscription{}.FromRef(_ret), _ok
}

type RTCAnswerOptions struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCAnswerOptions with all fields set.
func (p RTCAnswerOptions) FromRef(ref js.Ref) RTCAnswerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 RTCAnswerOptions RTCAnswerOptions [// RTCAnswerOptions] [] 0x14000921620 {0 0}} in the application heap.
func (p RTCAnswerOptions) New() js.Ref {
	return bindings.RTCAnswerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCAnswerOptions) UpdateFrom(ref js.Ref) {
	bindings.RTCAnswerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCAnswerOptions) Update(ref js.Ref) {
	bindings.RTCAnswerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCStatsType uint32

const (
	_ RTCStatsType = iota

	RTCStatsType_CODEC
	RTCStatsType_INBOUND_RTP
	RTCStatsType_OUTBOUND_RTP
	RTCStatsType_REMOTE_INBOUND_RTP
	RTCStatsType_REMOTE_OUTBOUND_RTP
	RTCStatsType_MEDIA_SOURCE
	RTCStatsType_MEDIA_PLAYOUT
	RTCStatsType_PEER_CONNECTION
	RTCStatsType_DATA_CHANNEL
	RTCStatsType_TRANSPORT
	RTCStatsType_CANDIDATE_PAIR
	RTCStatsType_LOCAL_CANDIDATE
	RTCStatsType_REMOTE_CANDIDATE
	RTCStatsType_CERTIFICATE
)

func (RTCStatsType) FromRef(str js.Ref) RTCStatsType {
	return RTCStatsType(bindings.ConstOfRTCStatsType(str))
}

func (x RTCStatsType) String() (string, bool) {
	switch x {
	case RTCStatsType_CODEC:
		return "codec", true
	case RTCStatsType_INBOUND_RTP:
		return "inbound-rtp", true
	case RTCStatsType_OUTBOUND_RTP:
		return "outbound-rtp", true
	case RTCStatsType_REMOTE_INBOUND_RTP:
		return "remote-inbound-rtp", true
	case RTCStatsType_REMOTE_OUTBOUND_RTP:
		return "remote-outbound-rtp", true
	case RTCStatsType_MEDIA_SOURCE:
		return "media-source", true
	case RTCStatsType_MEDIA_PLAYOUT:
		return "media-playout", true
	case RTCStatsType_PEER_CONNECTION:
		return "peer-connection", true
	case RTCStatsType_DATA_CHANNEL:
		return "data-channel", true
	case RTCStatsType_TRANSPORT:
		return "transport", true
	case RTCStatsType_CANDIDATE_PAIR:
		return "candidate-pair", true
	case RTCStatsType_LOCAL_CANDIDATE:
		return "local-candidate", true
	case RTCStatsType_REMOTE_CANDIDATE:
		return "remote-candidate", true
	case RTCStatsType_CERTIFICATE:
		return "certificate", true
	default:
		return "", false
	}
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

func NewPresentationConnectionAvailableEvent(typ js.String, eventInitDict PresentationConnectionAvailableEventInit) (ret PresentationConnectionAvailableEvent) {
	ret.ref = bindings.NewPresentationConnectionAvailableEventByPresentationConnectionAvailableEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type PresentationConnectionAvailableEvent struct {
	Event
}

func (this PresentationConnectionAvailableEvent) Once() PresentationConnectionAvailableEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Connection returns the value of property "PresentationConnectionAvailableEvent.connection".
//
// It returns ok=false if there is no such property.
func (this PresentationConnectionAvailableEvent) Connection() (ret PresentationConnection, ok bool) {
	ok = js.True == bindings.GetPresentationConnectionAvailableEventConnection(
		this.ref, js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PresentationConnectionCloseEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PresentationConnectionCloseEventInit.composed"
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

// FromRef calls UpdateFrom and returns a PresentationConnectionCloseEventInit with all fields set.
func (p PresentationConnectionCloseEventInit) FromRef(ref js.Ref) PresentationConnectionCloseEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PresentationConnectionCloseEventInit in the application heap.
func (p PresentationConnectionCloseEventInit) New() js.Ref {
	return bindings.PresentationConnectionCloseEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PresentationConnectionCloseEventInit) UpdateFrom(ref js.Ref) {
	bindings.PresentationConnectionCloseEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PresentationConnectionCloseEventInit) Update(ref js.Ref) {
	bindings.PresentationConnectionCloseEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PresentationConnectionCloseEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
}

func NewPresentationConnectionCloseEvent(typ js.String, eventInitDict PresentationConnectionCloseEventInit) (ret PresentationConnectionCloseEvent) {
	ret.ref = bindings.NewPresentationConnectionCloseEventByPresentationConnectionCloseEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type PresentationConnectionCloseEvent struct {
	Event
}

func (this PresentationConnectionCloseEvent) Once() PresentationConnectionCloseEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Reason returns the value of property "PresentationConnectionCloseEvent.reason".
//
// It returns ok=false if there is no such property.
func (this PresentationConnectionCloseEvent) Reason() (ret PresentationConnectionCloseReason, ok bool) {
	ok = js.True == bindings.GetPresentationConnectionCloseEventReason(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "PresentationConnectionCloseEvent.message".
//
// It returns ok=false if there is no such property.
func (this PresentationConnectionCloseEvent) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPresentationConnectionCloseEventMessage(
		this.ref, js.Pointer(&ret),
	)
	return
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// Source returns the value of property "PressureRecord.source".
//
// It returns ok=false if there is no such property.
func (this PressureRecord) Source() (ret PressureSource, ok bool) {
	ok = js.True == bindings.GetPressureRecordSource(
		this.ref, js.Pointer(&ret),
	)
	return
}

// State returns the value of property "PressureRecord.state".
//
// It returns ok=false if there is no such property.
func (this PressureRecord) State() (ret PressureState, ok bool) {
	ok = js.True == bindings.GetPressureRecordState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Time returns the value of property "PressureRecord.time".
//
// It returns ok=false if there is no such property.
func (this PressureRecord) Time() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPressureRecordTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncToJSON returns true if the method "PressureRecord.toJSON" exists.
func (this PressureRecord) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncPressureRecordToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "PressureRecord.toJSON".
func (this PressureRecord) FuncToJSON() (fn js.Func[func() js.Object]) {
	bindings.FuncPressureRecordToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "PressureRecord.toJSON".
func (this PressureRecord) ToJSON() (ret js.Object) {
	bindings.CallPressureRecordToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PressureRecord.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PressureRecord) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPressureRecordToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PressureObserverOptions struct {
	// SampleRate is "PressureObserverOptions.sampleRate"
	//
	// Optional, defaults to 1.0.
	//
	// NOTE: FFI_USE_SampleRate MUST be set to true to make this field effective.
	SampleRate float64

	FFI_USE_SampleRate bool // for SampleRate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PressureObserverOptions with all fields set.
func (p PressureObserverOptions) FromRef(ref js.Ref) PressureObserverOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PressureObserverOptions in the application heap.
func (p PressureObserverOptions) New() js.Ref {
	return bindings.PressureObserverOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PressureObserverOptions) UpdateFrom(ref js.Ref) {
	bindings.PressureObserverOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PressureObserverOptions) Update(ref js.Ref) {
	bindings.PressureObserverOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PressureObserverOptions) FreeMembers(recursive bool) {
}

func NewPressureObserver(callback js.Func[func(changes js.Array[PressureRecord], observer PressureObserver)], options PressureObserverOptions) (ret PressureObserver) {
	ret.ref = bindings.NewPressureObserverByPressureObserver(
		callback.Ref(),
		js.Pointer(&options))
	return
}

func NewPressureObserverByPressureObserver1(callback js.Func[func(changes js.Array[PressureRecord], observer PressureObserver)]) (ret PressureObserver) {
	ret.ref = bindings.NewPressureObserverByPressureObserver1(
		callback.Ref())
	return
}

type PressureObserver struct {
	ref js.Ref
}

func (this PressureObserver) Once() PressureObserver {
	this.ref.Once()
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
	this.ref.Free()
}

// SupportedSources returns the value of property "PressureObserver.supportedSources".
//
// It returns ok=false if there is no such property.
func (this PressureObserver) SupportedSources() (ret js.FrozenArray[PressureSource], ok bool) {
	ok = js.True == bindings.GetPressureObserverSupportedSources(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncObserve returns true if the method "PressureObserver.observe" exists.
func (this PressureObserver) HasFuncObserve() bool {
	return js.True == bindings.HasFuncPressureObserverObserve(
		this.ref,
	)
}

// FuncObserve returns the method "PressureObserver.observe".
func (this PressureObserver) FuncObserve() (fn js.Func[func(source PressureSource) js.Promise[js.Void]]) {
	bindings.FuncPressureObserverObserve(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Observe calls the method "PressureObserver.observe".
func (this PressureObserver) Observe(source PressureSource) (ret js.Promise[js.Void]) {
	bindings.CallPressureObserverObserve(
		this.ref, js.Pointer(&ret),
		uint32(source),
	)

	return
}

// TryObserve calls the method "PressureObserver.observe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PressureObserver) TryObserve(source PressureSource) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPressureObserverObserve(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(source),
	)

	return
}

// HasFuncUnobserve returns true if the method "PressureObserver.unobserve" exists.
func (this PressureObserver) HasFuncUnobserve() bool {
	return js.True == bindings.HasFuncPressureObserverUnobserve(
		this.ref,
	)
}

// FuncUnobserve returns the method "PressureObserver.unobserve".
func (this PressureObserver) FuncUnobserve() (fn js.Func[func(source PressureSource)]) {
	bindings.FuncPressureObserverUnobserve(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Unobserve calls the method "PressureObserver.unobserve".
func (this PressureObserver) Unobserve(source PressureSource) (ret js.Void) {
	bindings.CallPressureObserverUnobserve(
		this.ref, js.Pointer(&ret),
		uint32(source),
	)

	return
}

// TryUnobserve calls the method "PressureObserver.unobserve"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PressureObserver) TryUnobserve(source PressureSource) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPressureObserverUnobserve(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(source),
	)

	return
}

// HasFuncDisconnect returns true if the method "PressureObserver.disconnect" exists.
func (this PressureObserver) HasFuncDisconnect() bool {
	return js.True == bindings.HasFuncPressureObserverDisconnect(
		this.ref,
	)
}

// FuncDisconnect returns the method "PressureObserver.disconnect".
func (this PressureObserver) FuncDisconnect() (fn js.Func[func()]) {
	bindings.FuncPressureObserverDisconnect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Disconnect calls the method "PressureObserver.disconnect".
func (this PressureObserver) Disconnect() (ret js.Void) {
	bindings.CallPressureObserverDisconnect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDisconnect calls the method "PressureObserver.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PressureObserver) TryDisconnect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPressureObserverDisconnect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncTakeRecords returns true if the method "PressureObserver.takeRecords" exists.
func (this PressureObserver) HasFuncTakeRecords() bool {
	return js.True == bindings.HasFuncPressureObserverTakeRecords(
		this.ref,
	)
}

// FuncTakeRecords returns the method "PressureObserver.takeRecords".
func (this PressureObserver) FuncTakeRecords() (fn js.Func[func() js.Array[PressureRecord]]) {
	bindings.FuncPressureObserverTakeRecords(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TakeRecords calls the method "PressureObserver.takeRecords".
func (this PressureObserver) TakeRecords() (ret js.Array[PressureRecord]) {
	bindings.CallPressureObserverTakeRecords(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryTakeRecords calls the method "PressureObserver.takeRecords"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PressureObserver) TryTakeRecords() (ret js.Array[PressureRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPressureObserverTakeRecords(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// New creates a new PrivateNetworkAccessPermissionDescriptor in the application heap.
func (p PrivateNetworkAccessPermissionDescriptor) New() js.Ref {
	return bindings.PrivateNetworkAccessPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PrivateNetworkAccessPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.PrivateNetworkAccessPermissionDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PrivateNetworkAccessPermissionDescriptor) Update(ref js.Ref) {
	bindings.PrivateNetworkAccessPermissionDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PrivateNetworkAccessPermissionDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Name.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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

// New creates a new ProfilerInitOptions in the application heap.
func (p ProfilerInitOptions) New() js.Ref {
	return bindings.ProfilerInitOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProfilerInitOptions) UpdateFrom(ref js.Ref) {
	bindings.ProfilerInitOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProfilerInitOptions) Update(ref js.Ref) {
	bindings.ProfilerInitOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProfilerInitOptions) FreeMembers(recursive bool) {
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
	//
	// NOTE: FFI_USE_ResourceId MUST be set to true to make this field effective.
	ResourceId uint64
	// Line is "ProfilerFrame.line"
	//
	// Optional
	//
	// NOTE: FFI_USE_Line MUST be set to true to make this field effective.
	Line uint64
	// Column is "ProfilerFrame.column"
	//
	// Optional
	//
	// NOTE: FFI_USE_Column MUST be set to true to make this field effective.
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

// New creates a new ProfilerFrame in the application heap.
func (p ProfilerFrame) New() js.Ref {
	return bindings.ProfilerFrameJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProfilerFrame) UpdateFrom(ref js.Ref) {
	bindings.ProfilerFrameJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProfilerFrame) Update(ref js.Ref) {
	bindings.ProfilerFrameJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProfilerFrame) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
}

type ProfilerStack struct {
	// ParentId is "ProfilerStack.parentId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ParentId MUST be set to true to make this field effective.
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

// New creates a new ProfilerStack in the application heap.
func (p ProfilerStack) New() js.Ref {
	return bindings.ProfilerStackJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProfilerStack) UpdateFrom(ref js.Ref) {
	bindings.ProfilerStackJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProfilerStack) Update(ref js.Ref) {
	bindings.ProfilerStackJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProfilerStack) FreeMembers(recursive bool) {
}

type ProfilerSample struct {
	// Timestamp is "ProfilerSample.timestamp"
	//
	// Required
	Timestamp DOMHighResTimeStamp
	// StackId is "ProfilerSample.stackId"
	//
	// Optional
	//
	// NOTE: FFI_USE_StackId MUST be set to true to make this field effective.
	StackId uint64

	FFI_USE_StackId bool // for StackId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProfilerSample with all fields set.
func (p ProfilerSample) FromRef(ref js.Ref) ProfilerSample {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProfilerSample in the application heap.
func (p ProfilerSample) New() js.Ref {
	return bindings.ProfilerSampleJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProfilerSample) UpdateFrom(ref js.Ref) {
	bindings.ProfilerSampleJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProfilerSample) Update(ref js.Ref) {
	bindings.ProfilerSampleJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProfilerSample) FreeMembers(recursive bool) {
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

// New creates a new ProfilerTrace in the application heap.
func (p ProfilerTrace) New() js.Ref {
	return bindings.ProfilerTraceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProfilerTrace) UpdateFrom(ref js.Ref) {
	bindings.ProfilerTraceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProfilerTrace) Update(ref js.Ref) {
	bindings.ProfilerTraceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProfilerTrace) FreeMembers(recursive bool) {
	js.Free(
		p.Resources.Ref(),
		p.Frames.Ref(),
		p.Stacks.Ref(),
		p.Samples.Ref(),
	)
	p.Resources = p.Resources.FromRef(js.Undefined)
	p.Frames = p.Frames.FromRef(js.Undefined)
	p.Stacks = p.Stacks.FromRef(js.Undefined)
	p.Samples = p.Samples.FromRef(js.Undefined)
}

func NewProfiler(options ProfilerInitOptions) (ret Profiler) {
	ret.ref = bindings.NewProfilerByProfiler(
		js.Pointer(&options))
	return
}

type Profiler struct {
	EventTarget
}

func (this Profiler) Once() Profiler {
	this.ref.Once()
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
	this.ref.Free()
}

// SampleInterval returns the value of property "Profiler.sampleInterval".
//
// It returns ok=false if there is no such property.
func (this Profiler) SampleInterval() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetProfilerSampleInterval(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Stopped returns the value of property "Profiler.stopped".
//
// It returns ok=false if there is no such property.
func (this Profiler) Stopped() (ret bool, ok bool) {
	ok = js.True == bindings.GetProfilerStopped(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncStop returns true if the method "Profiler.stop" exists.
func (this Profiler) HasFuncStop() bool {
	return js.True == bindings.HasFuncProfilerStop(
		this.ref,
	)
}

// FuncStop returns the method "Profiler.stop".
func (this Profiler) FuncStop() (fn js.Func[func() js.Promise[ProfilerTrace]]) {
	bindings.FuncProfilerStop(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stop calls the method "Profiler.stop".
func (this Profiler) Stop() (ret js.Promise[ProfilerTrace]) {
	bindings.CallProfilerStop(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "Profiler.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Profiler) TryStop() (ret js.Promise[ProfilerTrace], exception js.Any, ok bool) {
	ok = js.True == bindings.TryProfilerStop(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ProgressEventInit struct {
	// LengthComputable is "ProgressEventInit.lengthComputable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_LengthComputable MUST be set to true to make this field effective.
	LengthComputable bool
	// Loaded is "ProgressEventInit.loaded"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Loaded MUST be set to true to make this field effective.
	Loaded uint64
	// Total is "ProgressEventInit.total"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Total MUST be set to true to make this field effective.
	Total uint64
	// Bubbles is "ProgressEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "ProgressEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "ProgressEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new ProgressEventInit in the application heap.
func (p ProgressEventInit) New() js.Ref {
	return bindings.ProgressEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProgressEventInit) UpdateFrom(ref js.Ref) {
	bindings.ProgressEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProgressEventInit) Update(ref js.Ref) {
	bindings.ProgressEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProgressEventInit) FreeMembers(recursive bool) {
}

func NewProgressEvent(typ js.String, eventInitDict ProgressEventInit) (ret ProgressEvent) {
	ret.ref = bindings.NewProgressEventByProgressEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewProgressEventByProgressEvent1(typ js.String) (ret ProgressEvent) {
	ret.ref = bindings.NewProgressEventByProgressEvent1(
		typ.Ref())
	return
}

type ProgressEvent struct {
	Event
}

func (this ProgressEvent) Once() ProgressEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// LengthComputable returns the value of property "ProgressEvent.lengthComputable".
//
// It returns ok=false if there is no such property.
func (this ProgressEvent) LengthComputable() (ret bool, ok bool) {
	ok = js.True == bindings.GetProgressEventLengthComputable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Loaded returns the value of property "ProgressEvent.loaded".
//
// It returns ok=false if there is no such property.
func (this ProgressEvent) Loaded() (ret uint64, ok bool) {
	ok = js.True == bindings.GetProgressEventLoaded(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Total returns the value of property "ProgressEvent.total".
//
// It returns ok=false if there is no such property.
func (this ProgressEvent) Total() (ret uint64, ok bool) {
	ok = js.True == bindings.GetProgressEventTotal(
		this.ref, js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PromiseRejectionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PromiseRejectionEventInit.composed"
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

// FromRef calls UpdateFrom and returns a PromiseRejectionEventInit with all fields set.
func (p PromiseRejectionEventInit) FromRef(ref js.Ref) PromiseRejectionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PromiseRejectionEventInit in the application heap.
func (p PromiseRejectionEventInit) New() js.Ref {
	return bindings.PromiseRejectionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PromiseRejectionEventInit) UpdateFrom(ref js.Ref) {
	bindings.PromiseRejectionEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PromiseRejectionEventInit) Update(ref js.Ref) {
	bindings.PromiseRejectionEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PromiseRejectionEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Promise.Ref(),
		p.Reason.Ref(),
	)
	p.Promise = p.Promise.FromRef(js.Undefined)
	p.Reason = p.Reason.FromRef(js.Undefined)
}

func NewPromiseRejectionEvent(typ js.String, eventInitDict PromiseRejectionEventInit) (ret PromiseRejectionEvent) {
	ret.ref = bindings.NewPromiseRejectionEventByPromiseRejectionEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type PromiseRejectionEvent struct {
	Event
}

func (this PromiseRejectionEvent) Once() PromiseRejectionEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Promise returns the value of property "PromiseRejectionEvent.promise".
//
// It returns ok=false if there is no such property.
func (this PromiseRejectionEvent) Promise() (ret js.Promise[js.Any], ok bool) {
	ok = js.True == bindings.GetPromiseRejectionEventPromise(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Reason returns the value of property "PromiseRejectionEvent.reason".
//
// It returns ok=false if there is no such property.
func (this PromiseRejectionEvent) Reason() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetPromiseRejectionEventReason(
		this.ref, js.Pointer(&ret),
	)
	return
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

// New creates a new ProximityReadingValues in the application heap.
func (p ProximityReadingValues) New() js.Ref {
	return bindings.ProximityReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProximityReadingValues) UpdateFrom(ref js.Ref) {
	bindings.ProximityReadingValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProximityReadingValues) Update(ref js.Ref) {
	bindings.ProximityReadingValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProximityReadingValues) FreeMembers(recursive bool) {
}

func NewProximitySensor(sensorOptions SensorOptions) (ret ProximitySensor) {
	ret.ref = bindings.NewProximitySensorByProximitySensor(
		js.Pointer(&sensorOptions))
	return
}

func NewProximitySensorByProximitySensor1() (ret ProximitySensor) {
	ret.ref = bindings.NewProximitySensorByProximitySensor1()
	return
}

type ProximitySensor struct {
	Sensor
}

func (this ProximitySensor) Once() ProximitySensor {
	this.ref.Once()
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
	this.ref.Free()
}

// Distance returns the value of property "ProximitySensor.distance".
//
// It returns ok=false if there is no such property.
func (this ProximitySensor) Distance() (ret float64, ok bool) {
	ok = js.True == bindings.GetProximitySensorDistance(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Max returns the value of property "ProximitySensor.max".
//
// It returns ok=false if there is no such property.
func (this ProximitySensor) Max() (ret float64, ok bool) {
	ok = js.True == bindings.GetProximitySensorMax(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Near returns the value of property "ProximitySensor.near".
//
// It returns ok=false if there is no such property.
func (this ProximitySensor) Near() (ret bool, ok bool) {
	ok = js.True == bindings.GetProximitySensorNear(
		this.ref, js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: Response.FFI_USE MUST be set to true to get Response used.
	Response AuthenticatorAttestationResponseJSON
	// AuthenticatorAttachment is "RegistrationResponseJSON.authenticatorAttachment"
	//
	// Optional
	AuthenticatorAttachment js.String
	// ClientExtensionResults is "RegistrationResponseJSON.clientExtensionResults"
	//
	// Required
	//
	// NOTE: ClientExtensionResults.FFI_USE MUST be set to true to get ClientExtensionResults used.
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

// New creates a new RegistrationResponseJSON in the application heap.
func (p RegistrationResponseJSON) New() js.Ref {
	return bindings.RegistrationResponseJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RegistrationResponseJSON) UpdateFrom(ref js.Ref) {
	bindings.RegistrationResponseJSONJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RegistrationResponseJSON) Update(ref js.Ref) {
	bindings.RegistrationResponseJSONJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RegistrationResponseJSON) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.RawId.Ref(),
		p.AuthenticatorAttachment.Ref(),
		p.Type.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.RawId = p.RawId.FromRef(js.Undefined)
	p.AuthenticatorAttachment = p.AuthenticatorAttachment.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
	if recursive {
		p.Response.FreeMembers(true)
		p.ClientExtensionResults.FreeMembers(true)
	}
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

// New creates a new PublicKeyCredentialDescriptorJSON in the application heap.
func (p PublicKeyCredentialDescriptorJSON) New() js.Ref {
	return bindings.PublicKeyCredentialDescriptorJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKeyCredentialDescriptorJSON) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialDescriptorJSONJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKeyCredentialDescriptorJSON) Update(ref js.Ref) {
	bindings.PublicKeyCredentialDescriptorJSONJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKeyCredentialDescriptorJSON) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Type.Ref(),
		p.Transports.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
	p.Transports = p.Transports.FromRef(js.Undefined)
}

type PublicKeyCredentialRequestOptionsJSON struct {
	// Challenge is "PublicKeyCredentialRequestOptionsJSON.challenge"
	//
	// Required
	Challenge Base64URLString
	// Timeout is "PublicKeyCredentialRequestOptionsJSON.timeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timeout MUST be set to true to make this field effective.
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
	//
	// NOTE: Extensions.FFI_USE MUST be set to true to get Extensions used.
	Extensions AuthenticationExtensionsClientInputsJSON

	FFI_USE_Timeout bool // for Timeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialRequestOptionsJSON with all fields set.
func (p PublicKeyCredentialRequestOptionsJSON) FromRef(ref js.Ref) PublicKeyCredentialRequestOptionsJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PublicKeyCredentialRequestOptionsJSON in the application heap.
func (p PublicKeyCredentialRequestOptionsJSON) New() js.Ref {
	return bindings.PublicKeyCredentialRequestOptionsJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKeyCredentialRequestOptionsJSON) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialRequestOptionsJSONJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKeyCredentialRequestOptionsJSON) Update(ref js.Ref) {
	bindings.PublicKeyCredentialRequestOptionsJSONJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKeyCredentialRequestOptionsJSON) FreeMembers(recursive bool) {
	js.Free(
		p.Challenge.Ref(),
		p.RpId.Ref(),
		p.AllowCredentials.Ref(),
		p.UserVerification.Ref(),
		p.Hints.Ref(),
		p.Attestation.Ref(),
		p.AttestationFormats.Ref(),
	)
	p.Challenge = p.Challenge.FromRef(js.Undefined)
	p.RpId = p.RpId.FromRef(js.Undefined)
	p.AllowCredentials = p.AllowCredentials.FromRef(js.Undefined)
	p.UserVerification = p.UserVerification.FromRef(js.Undefined)
	p.Hints = p.Hints.FromRef(js.Undefined)
	p.Attestation = p.Attestation.FromRef(js.Undefined)
	p.AttestationFormats = p.AttestationFormats.FromRef(js.Undefined)
	if recursive {
		p.Extensions.FreeMembers(true)
	}
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

// New creates a new PublicKeyCredentialUserEntityJSON in the application heap.
func (p PublicKeyCredentialUserEntityJSON) New() js.Ref {
	return bindings.PublicKeyCredentialUserEntityJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKeyCredentialUserEntityJSON) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialUserEntityJSONJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKeyCredentialUserEntityJSON) Update(ref js.Ref) {
	bindings.PublicKeyCredentialUserEntityJSONJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKeyCredentialUserEntityJSON) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Name.Ref(),
		p.DisplayName.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.DisplayName = p.DisplayName.FromRef(js.Undefined)
}

type PublicKeyCredentialCreationOptionsJSON struct {
	// Rp is "PublicKeyCredentialCreationOptionsJSON.rp"
	//
	// Required
	//
	// NOTE: Rp.FFI_USE MUST be set to true to get Rp used.
	Rp PublicKeyCredentialRpEntity
	// User is "PublicKeyCredentialCreationOptionsJSON.user"
	//
	// Required
	//
	// NOTE: User.FFI_USE MUST be set to true to get User used.
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
	//
	// NOTE: FFI_USE_Timeout MUST be set to true to make this field effective.
	Timeout uint32
	// ExcludeCredentials is "PublicKeyCredentialCreationOptionsJSON.excludeCredentials"
	//
	// Optional, defaults to [].
	ExcludeCredentials js.Array[PublicKeyCredentialDescriptorJSON]
	// AuthenticatorSelection is "PublicKeyCredentialCreationOptionsJSON.authenticatorSelection"
	//
	// Optional
	//
	// NOTE: AuthenticatorSelection.FFI_USE MUST be set to true to get AuthenticatorSelection used.
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
	//
	// NOTE: Extensions.FFI_USE MUST be set to true to get Extensions used.
	Extensions AuthenticationExtensionsClientInputsJSON

	FFI_USE_Timeout bool // for Timeout.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PublicKeyCredentialCreationOptionsJSON with all fields set.
func (p PublicKeyCredentialCreationOptionsJSON) FromRef(ref js.Ref) PublicKeyCredentialCreationOptionsJSON {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PublicKeyCredentialCreationOptionsJSON in the application heap.
func (p PublicKeyCredentialCreationOptionsJSON) New() js.Ref {
	return bindings.PublicKeyCredentialCreationOptionsJSONJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKeyCredentialCreationOptionsJSON) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialCreationOptionsJSONJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKeyCredentialCreationOptionsJSON) Update(ref js.Ref) {
	bindings.PublicKeyCredentialCreationOptionsJSONJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKeyCredentialCreationOptionsJSON) FreeMembers(recursive bool) {
	js.Free(
		p.Challenge.Ref(),
		p.PubKeyCredParams.Ref(),
		p.ExcludeCredentials.Ref(),
		p.Hints.Ref(),
		p.Attestation.Ref(),
		p.AttestationFormats.Ref(),
	)
	p.Challenge = p.Challenge.FromRef(js.Undefined)
	p.PubKeyCredParams = p.PubKeyCredParams.FromRef(js.Undefined)
	p.ExcludeCredentials = p.ExcludeCredentials.FromRef(js.Undefined)
	p.Hints = p.Hints.FromRef(js.Undefined)
	p.Attestation = p.Attestation.FromRef(js.Undefined)
	p.AttestationFormats = p.AttestationFormats.FromRef(js.Undefined)
	if recursive {
		p.Rp.FreeMembers(true)
		p.User.FreeMembers(true)
		p.AuthenticatorSelection.FreeMembers(true)
		p.Extensions.FreeMembers(true)
	}
}

type PublicKeyCredential struct {
	Credential
}

func (this PublicKeyCredential) Once() PublicKeyCredential {
	this.ref.Once()
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
	this.ref.Free()
}

// RawId returns the value of property "PublicKeyCredential.rawId".
//
// It returns ok=false if there is no such property.
func (this PublicKeyCredential) RawId() (ret js.ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetPublicKeyCredentialRawId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Response returns the value of property "PublicKeyCredential.response".
//
// It returns ok=false if there is no such property.
func (this PublicKeyCredential) Response() (ret AuthenticatorResponse, ok bool) {
	ok = js.True == bindings.GetPublicKeyCredentialResponse(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AuthenticatorAttachment returns the value of property "PublicKeyCredential.authenticatorAttachment".
//
// It returns ok=false if there is no such property.
func (this PublicKeyCredential) AuthenticatorAttachment() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPublicKeyCredentialAuthenticatorAttachment(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetClientExtensionResults returns true if the method "PublicKeyCredential.getClientExtensionResults" exists.
func (this PublicKeyCredential) HasFuncGetClientExtensionResults() bool {
	return js.True == bindings.HasFuncPublicKeyCredentialGetClientExtensionResults(
		this.ref,
	)
}

// FuncGetClientExtensionResults returns the method "PublicKeyCredential.getClientExtensionResults".
func (this PublicKeyCredential) FuncGetClientExtensionResults() (fn js.Func[func() AuthenticationExtensionsClientOutputs]) {
	bindings.FuncPublicKeyCredentialGetClientExtensionResults(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetClientExtensionResults calls the method "PublicKeyCredential.getClientExtensionResults".
func (this PublicKeyCredential) GetClientExtensionResults() (ret AuthenticationExtensionsClientOutputs) {
	bindings.CallPublicKeyCredentialGetClientExtensionResults(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetClientExtensionResults calls the method "PublicKeyCredential.getClientExtensionResults"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PublicKeyCredential) TryGetClientExtensionResults() (ret AuthenticationExtensionsClientOutputs, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPublicKeyCredentialGetClientExtensionResults(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsConditionalMediationAvailable returns true if the static method "PublicKeyCredential.isConditionalMediationAvailable" exists.
func (this PublicKeyCredential) HasFuncIsConditionalMediationAvailable() bool {
	return js.True == bindings.HasFuncPublicKeyCredentialIsConditionalMediationAvailable(
		this.ref,
	)
}

// FuncIsConditionalMediationAvailable returns the static method "PublicKeyCredential.isConditionalMediationAvailable".
func (this PublicKeyCredential) FuncIsConditionalMediationAvailable() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncPublicKeyCredentialIsConditionalMediationAvailable(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsConditionalMediationAvailable calls the static method "PublicKeyCredential.isConditionalMediationAvailable".
func (this PublicKeyCredential) IsConditionalMediationAvailable() (ret js.Promise[js.Boolean]) {
	bindings.CallPublicKeyCredentialIsConditionalMediationAvailable(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsConditionalMediationAvailable calls the static method "PublicKeyCredential.isConditionalMediationAvailable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PublicKeyCredential) TryIsConditionalMediationAvailable() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPublicKeyCredentialIsConditionalMediationAvailable(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncToJSON returns true if the method "PublicKeyCredential.toJSON" exists.
func (this PublicKeyCredential) HasFuncToJSON() bool {
	return js.True == bindings.HasFuncPublicKeyCredentialToJSON(
		this.ref,
	)
}

// FuncToJSON returns the method "PublicKeyCredential.toJSON".
func (this PublicKeyCredential) FuncToJSON() (fn js.Func[func() PublicKeyCredentialJSON]) {
	bindings.FuncPublicKeyCredentialToJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToJSON calls the method "PublicKeyCredential.toJSON".
func (this PublicKeyCredential) ToJSON() (ret PublicKeyCredentialJSON) {
	bindings.CallPublicKeyCredentialToJSON(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PublicKeyCredential.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PublicKeyCredential) TryToJSON() (ret PublicKeyCredentialJSON, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPublicKeyCredentialToJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncParseRequestOptionsFromJSON returns true if the static method "PublicKeyCredential.parseRequestOptionsFromJSON" exists.
func (this PublicKeyCredential) HasFuncParseRequestOptionsFromJSON() bool {
	return js.True == bindings.HasFuncPublicKeyCredentialParseRequestOptionsFromJSON(
		this.ref,
	)
}

// FuncParseRequestOptionsFromJSON returns the static method "PublicKeyCredential.parseRequestOptionsFromJSON".
func (this PublicKeyCredential) FuncParseRequestOptionsFromJSON() (fn js.Func[func(options PublicKeyCredentialRequestOptionsJSON) PublicKeyCredentialRequestOptions]) {
	bindings.FuncPublicKeyCredentialParseRequestOptionsFromJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ParseRequestOptionsFromJSON calls the static method "PublicKeyCredential.parseRequestOptionsFromJSON".
func (this PublicKeyCredential) ParseRequestOptionsFromJSON(options PublicKeyCredentialRequestOptionsJSON) (ret PublicKeyCredentialRequestOptions) {
	bindings.CallPublicKeyCredentialParseRequestOptionsFromJSON(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryParseRequestOptionsFromJSON calls the static method "PublicKeyCredential.parseRequestOptionsFromJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PublicKeyCredential) TryParseRequestOptionsFromJSON(options PublicKeyCredentialRequestOptionsJSON) (ret PublicKeyCredentialRequestOptions, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPublicKeyCredentialParseRequestOptionsFromJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncIsPasskeyPlatformAuthenticatorAvailable returns true if the static method "PublicKeyCredential.isPasskeyPlatformAuthenticatorAvailable" exists.
func (this PublicKeyCredential) HasFuncIsPasskeyPlatformAuthenticatorAvailable() bool {
	return js.True == bindings.HasFuncPublicKeyCredentialIsPasskeyPlatformAuthenticatorAvailable(
		this.ref,
	)
}

// FuncIsPasskeyPlatformAuthenticatorAvailable returns the static method "PublicKeyCredential.isPasskeyPlatformAuthenticatorAvailable".
func (this PublicKeyCredential) FuncIsPasskeyPlatformAuthenticatorAvailable() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncPublicKeyCredentialIsPasskeyPlatformAuthenticatorAvailable(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsPasskeyPlatformAuthenticatorAvailable calls the static method "PublicKeyCredential.isPasskeyPlatformAuthenticatorAvailable".
func (this PublicKeyCredential) IsPasskeyPlatformAuthenticatorAvailable() (ret js.Promise[js.Boolean]) {
	bindings.CallPublicKeyCredentialIsPasskeyPlatformAuthenticatorAvailable(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsPasskeyPlatformAuthenticatorAvailable calls the static method "PublicKeyCredential.isPasskeyPlatformAuthenticatorAvailable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PublicKeyCredential) TryIsPasskeyPlatformAuthenticatorAvailable() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPublicKeyCredentialIsPasskeyPlatformAuthenticatorAvailable(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncParseCreationOptionsFromJSON returns true if the static method "PublicKeyCredential.parseCreationOptionsFromJSON" exists.
func (this PublicKeyCredential) HasFuncParseCreationOptionsFromJSON() bool {
	return js.True == bindings.HasFuncPublicKeyCredentialParseCreationOptionsFromJSON(
		this.ref,
	)
}

// FuncParseCreationOptionsFromJSON returns the static method "PublicKeyCredential.parseCreationOptionsFromJSON".
func (this PublicKeyCredential) FuncParseCreationOptionsFromJSON() (fn js.Func[func(options PublicKeyCredentialCreationOptionsJSON) PublicKeyCredentialCreationOptions]) {
	bindings.FuncPublicKeyCredentialParseCreationOptionsFromJSON(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ParseCreationOptionsFromJSON calls the static method "PublicKeyCredential.parseCreationOptionsFromJSON".
func (this PublicKeyCredential) ParseCreationOptionsFromJSON(options PublicKeyCredentialCreationOptionsJSON) (ret PublicKeyCredentialCreationOptions) {
	bindings.CallPublicKeyCredentialParseCreationOptionsFromJSON(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryParseCreationOptionsFromJSON calls the static method "PublicKeyCredential.parseCreationOptionsFromJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PublicKeyCredential) TryParseCreationOptionsFromJSON(options PublicKeyCredentialCreationOptionsJSON) (ret PublicKeyCredentialCreationOptions, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPublicKeyCredentialParseCreationOptionsFromJSON(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncIsUserVerifyingPlatformAuthenticatorAvailable returns true if the static method "PublicKeyCredential.isUserVerifyingPlatformAuthenticatorAvailable" exists.
func (this PublicKeyCredential) HasFuncIsUserVerifyingPlatformAuthenticatorAvailable() bool {
	return js.True == bindings.HasFuncPublicKeyCredentialIsUserVerifyingPlatformAuthenticatorAvailable(
		this.ref,
	)
}

// FuncIsUserVerifyingPlatformAuthenticatorAvailable returns the static method "PublicKeyCredential.isUserVerifyingPlatformAuthenticatorAvailable".
func (this PublicKeyCredential) FuncIsUserVerifyingPlatformAuthenticatorAvailable() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncPublicKeyCredentialIsUserVerifyingPlatformAuthenticatorAvailable(
		this.ref, js.Pointer(&fn),
	)
	return
}

// IsUserVerifyingPlatformAuthenticatorAvailable calls the static method "PublicKeyCredential.isUserVerifyingPlatformAuthenticatorAvailable".
func (this PublicKeyCredential) IsUserVerifyingPlatformAuthenticatorAvailable() (ret js.Promise[js.Boolean]) {
	bindings.CallPublicKeyCredentialIsUserVerifyingPlatformAuthenticatorAvailable(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryIsUserVerifyingPlatformAuthenticatorAvailable calls the static method "PublicKeyCredential.isUserVerifyingPlatformAuthenticatorAvailable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PublicKeyCredential) TryIsUserVerifyingPlatformAuthenticatorAvailable() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPublicKeyCredentialIsUserVerifyingPlatformAuthenticatorAvailable(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// New creates a new PublicKeyCredentialEntity in the application heap.
func (p PublicKeyCredentialEntity) New() js.Ref {
	return bindings.PublicKeyCredentialEntityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PublicKeyCredentialEntity) UpdateFrom(ref js.Ref) {
	bindings.PublicKeyCredentialEntityJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PublicKeyCredentialEntity) Update(ref js.Ref) {
	bindings.PublicKeyCredentialEntityJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PublicKeyCredentialEntity) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PushEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PushEventInit.composed"
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

// FromRef calls UpdateFrom and returns a PushEventInit with all fields set.
func (p PushEventInit) FromRef(ref js.Ref) PushEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PushEventInit in the application heap.
func (p PushEventInit) New() js.Ref {
	return bindings.PushEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PushEventInit) UpdateFrom(ref js.Ref) {
	bindings.PushEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PushEventInit) Update(ref js.Ref) {
	bindings.PushEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PushEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
}

type PushMessageData struct {
	ref js.Ref
}

func (this PushMessageData) Once() PushMessageData {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncArrayBuffer returns true if the method "PushMessageData.arrayBuffer" exists.
func (this PushMessageData) HasFuncArrayBuffer() bool {
	return js.True == bindings.HasFuncPushMessageDataArrayBuffer(
		this.ref,
	)
}

// FuncArrayBuffer returns the method "PushMessageData.arrayBuffer".
func (this PushMessageData) FuncArrayBuffer() (fn js.Func[func() js.ArrayBuffer]) {
	bindings.FuncPushMessageDataArrayBuffer(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ArrayBuffer calls the method "PushMessageData.arrayBuffer".
func (this PushMessageData) ArrayBuffer() (ret js.ArrayBuffer) {
	bindings.CallPushMessageDataArrayBuffer(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryArrayBuffer calls the method "PushMessageData.arrayBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushMessageData) TryArrayBuffer() (ret js.ArrayBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushMessageDataArrayBuffer(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncBlob returns true if the method "PushMessageData.blob" exists.
func (this PushMessageData) HasFuncBlob() bool {
	return js.True == bindings.HasFuncPushMessageDataBlob(
		this.ref,
	)
}

// FuncBlob returns the method "PushMessageData.blob".
func (this PushMessageData) FuncBlob() (fn js.Func[func() Blob]) {
	bindings.FuncPushMessageDataBlob(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Blob calls the method "PushMessageData.blob".
func (this PushMessageData) Blob() (ret Blob) {
	bindings.CallPushMessageDataBlob(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBlob calls the method "PushMessageData.blob"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushMessageData) TryBlob() (ret Blob, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushMessageDataBlob(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncJson returns true if the method "PushMessageData.json" exists.
func (this PushMessageData) HasFuncJson() bool {
	return js.True == bindings.HasFuncPushMessageDataJson(
		this.ref,
	)
}

// FuncJson returns the method "PushMessageData.json".
func (this PushMessageData) FuncJson() (fn js.Func[func() js.Any]) {
	bindings.FuncPushMessageDataJson(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Json calls the method "PushMessageData.json".
func (this PushMessageData) Json() (ret js.Any) {
	bindings.CallPushMessageDataJson(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryJson calls the method "PushMessageData.json"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushMessageData) TryJson() (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushMessageDataJson(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncText returns true if the method "PushMessageData.text" exists.
func (this PushMessageData) HasFuncText() bool {
	return js.True == bindings.HasFuncPushMessageDataText(
		this.ref,
	)
}

// FuncText returns the method "PushMessageData.text".
func (this PushMessageData) FuncText() (fn js.Func[func() js.String]) {
	bindings.FuncPushMessageDataText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Text calls the method "PushMessageData.text".
func (this PushMessageData) Text() (ret js.String) {
	bindings.CallPushMessageDataText(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryText calls the method "PushMessageData.text"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PushMessageData) TryText() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPushMessageDataText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewPushEvent(typ js.String, eventInitDict PushEventInit) (ret PushEvent) {
	ret.ref = bindings.NewPushEventByPushEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewPushEventByPushEvent1(typ js.String) (ret PushEvent) {
	ret.ref = bindings.NewPushEventByPushEvent1(
		typ.Ref())
	return
}

type PushEvent struct {
	ExtendableEvent
}

func (this PushEvent) Once() PushEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Data returns the value of property "PushEvent.data".
//
// It returns ok=false if there is no such property.
func (this PushEvent) Data() (ret PushMessageData, ok bool) {
	ok = js.True == bindings.GetPushEventData(
		this.ref, js.Pointer(&ret),
	)
	return
}

type PushPermissionDescriptor struct {
	// UserVisibleOnly is "PushPermissionDescriptor.userVisibleOnly"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_UserVisibleOnly MUST be set to true to make this field effective.
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

// New creates a new PushPermissionDescriptor in the application heap.
func (p PushPermissionDescriptor) New() js.Ref {
	return bindings.PushPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PushPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.PushPermissionDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PushPermissionDescriptor) Update(ref js.Ref) {
	bindings.PushPermissionDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PushPermissionDescriptor) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PushSubscriptionChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PushSubscriptionChangeEventInit.composed"
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

// FromRef calls UpdateFrom and returns a PushSubscriptionChangeEventInit with all fields set.
func (p PushSubscriptionChangeEventInit) FromRef(ref js.Ref) PushSubscriptionChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PushSubscriptionChangeEventInit in the application heap.
func (p PushSubscriptionChangeEventInit) New() js.Ref {
	return bindings.PushSubscriptionChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PushSubscriptionChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.PushSubscriptionChangeEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PushSubscriptionChangeEventInit) Update(ref js.Ref) {
	bindings.PushSubscriptionChangeEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PushSubscriptionChangeEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.NewSubscription.Ref(),
		p.OldSubscription.Ref(),
	)
	p.NewSubscription = p.NewSubscription.FromRef(js.Undefined)
	p.OldSubscription = p.OldSubscription.FromRef(js.Undefined)
}

func NewPushSubscriptionChangeEvent(typ js.String, eventInitDict PushSubscriptionChangeEventInit) (ret PushSubscriptionChangeEvent) {
	ret.ref = bindings.NewPushSubscriptionChangeEventByPushSubscriptionChangeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewPushSubscriptionChangeEventByPushSubscriptionChangeEvent1(typ js.String) (ret PushSubscriptionChangeEvent) {
	ret.ref = bindings.NewPushSubscriptionChangeEventByPushSubscriptionChangeEvent1(
		typ.Ref())
	return
}

type PushSubscriptionChangeEvent struct {
	ExtendableEvent
}

func (this PushSubscriptionChangeEvent) Once() PushSubscriptionChangeEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// NewSubscription returns the value of property "PushSubscriptionChangeEvent.newSubscription".
//
// It returns ok=false if there is no such property.
func (this PushSubscriptionChangeEvent) NewSubscription() (ret PushSubscription, ok bool) {
	ok = js.True == bindings.GetPushSubscriptionChangeEventNewSubscription(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OldSubscription returns the value of property "PushSubscriptionChangeEvent.oldSubscription".
//
// It returns ok=false if there is no such property.
func (this PushSubscriptionChangeEvent) OldSubscription() (ret PushSubscription, ok bool) {
	ok = js.True == bindings.GetPushSubscriptionChangeEventOldSubscription(
		this.ref, js.Pointer(&ret),
	)
	return
}

type RTCAnswerOptions struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCAnswerOptions with all fields set.
func (p RTCAnswerOptions) FromRef(ref js.Ref) RTCAnswerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCAnswerOptions in the application heap.
func (p RTCAnswerOptions) New() js.Ref {
	return bindings.RTCAnswerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCAnswerOptions) UpdateFrom(ref js.Ref) {
	bindings.RTCAnswerOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCAnswerOptions) Update(ref js.Ref) {
	bindings.RTCAnswerOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCAnswerOptions) FreeMembers(recursive bool) {
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

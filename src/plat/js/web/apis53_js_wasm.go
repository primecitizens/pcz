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

type ScrollOptions struct {
	// Behavior is "ScrollOptions.behavior"
	//
	// Optional, defaults to "auto".
	Behavior ScrollBehavior

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScrollOptions with all fields set.
func (p ScrollOptions) FromRef(ref js.Ref) ScrollOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ScrollOptions ScrollOptions [// ScrollOptions] [0x14000a43b80] 0x14000a02228 {0 0}} in the application heap.
func (p ScrollOptions) New() js.Ref {
	return bindings.ScrollOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ScrollOptions) UpdateFrom(ref js.Ref) {
	bindings.ScrollOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ScrollOptions) Update(ref js.Ref) {
	bindings.ScrollOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ScrollSetting uint32

const (
	_ ScrollSetting = iota

	ScrollSetting_
	ScrollSetting_UP
)

func (ScrollSetting) FromRef(str js.Ref) ScrollSetting {
	return ScrollSetting(bindings.ConstOfScrollSetting(str))
}

func (x ScrollSetting) String() (string, bool) {
	switch x {
	case ScrollSetting_:
		return "", true
	case ScrollSetting_UP:
		return "up", true
	default:
		return "", false
	}
}

type ScrollTimelineOptions struct {
	// Source is "ScrollTimelineOptions.source"
	//
	// Optional
	Source Element
	// Axis is "ScrollTimelineOptions.axis"
	//
	// Optional, defaults to "block".
	Axis ScrollAxis

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScrollTimelineOptions with all fields set.
func (p ScrollTimelineOptions) FromRef(ref js.Ref) ScrollTimelineOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ScrollTimelineOptions ScrollTimelineOptions [// ScrollTimelineOptions] [0x14000a43c20 0x14000a43cc0] 0x14000a02270 {0 0}} in the application heap.
func (p ScrollTimelineOptions) New() js.Ref {
	return bindings.ScrollTimelineOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ScrollTimelineOptions) UpdateFrom(ref js.Ref) {
	bindings.ScrollTimelineOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ScrollTimelineOptions) Update(ref js.Ref) {
	bindings.ScrollTimelineOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewScrollTimeline(options ScrollTimelineOptions) ScrollTimeline {
	return ScrollTimeline{}.FromRef(
		bindings.NewScrollTimelineByScrollTimeline(
			js.Pointer(&options)),
	)
}

func NewScrollTimelineByScrollTimeline1() ScrollTimeline {
	return ScrollTimeline{}.FromRef(
		bindings.NewScrollTimelineByScrollTimeline1(),
	)
}

type ScrollTimeline struct {
	AnimationTimeline
}

func (this ScrollTimeline) Once() ScrollTimeline {
	this.Ref().Once()
	return this
}

func (this ScrollTimeline) Ref() js.Ref {
	return this.AnimationTimeline.Ref()
}

func (this ScrollTimeline) FromRef(ref js.Ref) ScrollTimeline {
	this.AnimationTimeline = this.AnimationTimeline.FromRef(ref)
	return this
}

func (this ScrollTimeline) Free() {
	this.Ref().Free()
}

// Source returns the value of property "ScrollTimeline.source".
//
// The returned bool will be false if there is no such property.
func (this ScrollTimeline) Source() (Element, bool) {
	var _ok bool
	_ret := bindings.GetScrollTimelineSource(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// Axis returns the value of property "ScrollTimeline.axis".
//
// The returned bool will be false if there is no such property.
func (this ScrollTimeline) Axis() (ScrollAxis, bool) {
	var _ok bool
	_ret := bindings.GetScrollTimelineAxis(
		this.Ref(), js.Pointer(&_ok),
	)
	return ScrollAxis(_ret), _ok
}

type SecurePaymentConfirmationRequest struct {
	// Challenge is "SecurePaymentConfirmationRequest.challenge"
	//
	// Required
	Challenge BufferSource
	// RpId is "SecurePaymentConfirmationRequest.rpId"
	//
	// Required
	RpId js.String
	// CredentialIds is "SecurePaymentConfirmationRequest.credentialIds"
	//
	// Required
	CredentialIds js.Array[BufferSource]
	// Instrument is "SecurePaymentConfirmationRequest.instrument"
	//
	// Required
	Instrument PaymentCredentialInstrument
	// Timeout is "SecurePaymentConfirmationRequest.timeout"
	//
	// Optional
	Timeout uint32
	// PayeeName is "SecurePaymentConfirmationRequest.payeeName"
	//
	// Optional
	PayeeName js.String
	// PayeeOrigin is "SecurePaymentConfirmationRequest.payeeOrigin"
	//
	// Optional
	PayeeOrigin js.String
	// Extensions is "SecurePaymentConfirmationRequest.extensions"
	//
	// Optional
	Extensions AuthenticationExtensionsClientInputs
	// Locale is "SecurePaymentConfirmationRequest.locale"
	//
	// Optional
	Locale js.Array[js.String]
	// ShowOptOut is "SecurePaymentConfirmationRequest.showOptOut"
	//
	// Optional
	ShowOptOut bool

	FFI_USE_Timeout    bool // for Timeout.
	FFI_USE_ShowOptOut bool // for ShowOptOut.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SecurePaymentConfirmationRequest with all fields set.
func (p SecurePaymentConfirmationRequest) FromRef(ref js.Ref) SecurePaymentConfirmationRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SecurePaymentConfirmationRequest SecurePaymentConfirmationRequest [// SecurePaymentConfirmationRequest] [0x14000a43e00 0x14000a43ea0 0x14000a43f40 0x14000a50000 0x14000a500a0 0x14000a50280 0x14000a50320 0x14000a503c0 0x14000a50460 0x14000a50500 0x14000a50140 0x14000a505a0] 0x14000a022a0 {0 0}} in the application heap.
func (p SecurePaymentConfirmationRequest) New() js.Ref {
	return bindings.SecurePaymentConfirmationRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SecurePaymentConfirmationRequest) UpdateFrom(ref js.Ref) {
	bindings.SecurePaymentConfirmationRequestJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SecurePaymentConfirmationRequest) Update(ref js.Ref) {
	bindings.SecurePaymentConfirmationRequestJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SecurityPolicyViolationEventInit struct {
	// DocumentURI is "SecurityPolicyViolationEventInit.documentURI"
	//
	// Required
	DocumentURI js.String
	// Referrer is "SecurityPolicyViolationEventInit.referrer"
	//
	// Optional, defaults to "".
	Referrer js.String
	// BlockedURI is "SecurityPolicyViolationEventInit.blockedURI"
	//
	// Optional, defaults to "".
	BlockedURI js.String
	// ViolatedDirective is "SecurityPolicyViolationEventInit.violatedDirective"
	//
	// Required
	ViolatedDirective js.String
	// EffectiveDirective is "SecurityPolicyViolationEventInit.effectiveDirective"
	//
	// Required
	EffectiveDirective js.String
	// OriginalPolicy is "SecurityPolicyViolationEventInit.originalPolicy"
	//
	// Required
	OriginalPolicy js.String
	// SourceFile is "SecurityPolicyViolationEventInit.sourceFile"
	//
	// Optional, defaults to "".
	SourceFile js.String
	// Sample is "SecurityPolicyViolationEventInit.sample"
	//
	// Optional, defaults to "".
	Sample js.String
	// Disposition is "SecurityPolicyViolationEventInit.disposition"
	//
	// Required
	Disposition SecurityPolicyViolationEventDisposition
	// StatusCode is "SecurityPolicyViolationEventInit.statusCode"
	//
	// Required
	StatusCode uint16
	// LineNumber is "SecurityPolicyViolationEventInit.lineNumber"
	//
	// Optional, defaults to 0.
	LineNumber uint32
	// ColumnNumber is "SecurityPolicyViolationEventInit.columnNumber"
	//
	// Optional, defaults to 0.
	ColumnNumber uint32
	// Bubbles is "SecurityPolicyViolationEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "SecurityPolicyViolationEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "SecurityPolicyViolationEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_LineNumber   bool // for LineNumber.
	FFI_USE_ColumnNumber bool // for ColumnNumber.
	FFI_USE_Bubbles      bool // for Bubbles.
	FFI_USE_Cancelable   bool // for Cancelable.
	FFI_USE_Composed     bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SecurityPolicyViolationEventInit with all fields set.
func (p SecurityPolicyViolationEventInit) FromRef(ref js.Ref) SecurityPolicyViolationEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SecurityPolicyViolationEventInit SecurityPolicyViolationEventInit [// SecurityPolicyViolationEventInit] [0x14000a50640 0x14000a506e0 0x14000a50780 0x14000a50820 0x14000a508c0 0x14000a50960 0x14000a50a00 0x14000a50aa0 0x14000a50b40 0x14000a50be0 0x14000a50c80 0x14000a50dc0 0x14000a50f00 0x14000a51040 0x14000a51180 0x14000a50d20 0x14000a50e60 0x14000a50fa0 0x14000a510e0 0x14000a51220] 0x14000a02300 {0 0}} in the application heap.
func (p SecurityPolicyViolationEventInit) New() js.Ref {
	return bindings.SecurityPolicyViolationEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SecurityPolicyViolationEventInit) UpdateFrom(ref js.Ref) {
	bindings.SecurityPolicyViolationEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SecurityPolicyViolationEventInit) Update(ref js.Ref) {
	bindings.SecurityPolicyViolationEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewSecurityPolicyViolationEvent(typ js.String, eventInitDict SecurityPolicyViolationEventInit) SecurityPolicyViolationEvent {
	return SecurityPolicyViolationEvent{}.FromRef(
		bindings.NewSecurityPolicyViolationEventBySecurityPolicyViolationEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewSecurityPolicyViolationEventBySecurityPolicyViolationEvent1(typ js.String) SecurityPolicyViolationEvent {
	return SecurityPolicyViolationEvent{}.FromRef(
		bindings.NewSecurityPolicyViolationEventBySecurityPolicyViolationEvent1(
			typ.Ref()),
	)
}

type SecurityPolicyViolationEvent struct {
	Event
}

func (this SecurityPolicyViolationEvent) Once() SecurityPolicyViolationEvent {
	this.Ref().Once()
	return this
}

func (this SecurityPolicyViolationEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this SecurityPolicyViolationEvent) FromRef(ref js.Ref) SecurityPolicyViolationEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this SecurityPolicyViolationEvent) Free() {
	this.Ref().Free()
}

// DocumentURI returns the value of property "SecurityPolicyViolationEvent.documentURI".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) DocumentURI() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventDocumentURI(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Referrer returns the value of property "SecurityPolicyViolationEvent.referrer".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) Referrer() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventReferrer(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BlockedURI returns the value of property "SecurityPolicyViolationEvent.blockedURI".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) BlockedURI() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventBlockedURI(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// EffectiveDirective returns the value of property "SecurityPolicyViolationEvent.effectiveDirective".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) EffectiveDirective() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventEffectiveDirective(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ViolatedDirective returns the value of property "SecurityPolicyViolationEvent.violatedDirective".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) ViolatedDirective() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventViolatedDirective(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// OriginalPolicy returns the value of property "SecurityPolicyViolationEvent.originalPolicy".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) OriginalPolicy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventOriginalPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SourceFile returns the value of property "SecurityPolicyViolationEvent.sourceFile".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) SourceFile() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventSourceFile(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Sample returns the value of property "SecurityPolicyViolationEvent.sample".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) Sample() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventSample(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Disposition returns the value of property "SecurityPolicyViolationEvent.disposition".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) Disposition() (SecurityPolicyViolationEventDisposition, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventDisposition(
		this.Ref(), js.Pointer(&_ok),
	)
	return SecurityPolicyViolationEventDisposition(_ret), _ok
}

// StatusCode returns the value of property "SecurityPolicyViolationEvent.statusCode".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) StatusCode() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventStatusCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// LineNumber returns the value of property "SecurityPolicyViolationEvent.lineNumber".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) LineNumber() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventLineNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ColumnNumber returns the value of property "SecurityPolicyViolationEvent.columnNumber".
//
// The returned bool will be false if there is no such property.
func (this SecurityPolicyViolationEvent) ColumnNumber() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSecurityPolicyViolationEventColumnNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

type Sensor struct {
	EventTarget
}

func (this Sensor) Once() Sensor {
	this.Ref().Once()
	return this
}

func (this Sensor) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Sensor) FromRef(ref js.Ref) Sensor {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Sensor) Free() {
	this.Ref().Free()
}

// Activated returns the value of property "Sensor.activated".
//
// The returned bool will be false if there is no such property.
func (this Sensor) Activated() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSensorActivated(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// HasReading returns the value of property "Sensor.hasReading".
//
// The returned bool will be false if there is no such property.
func (this Sensor) HasReading() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSensorHasReading(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Timestamp returns the value of property "Sensor.timestamp".
//
// The returned bool will be false if there is no such property.
func (this Sensor) Timestamp() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetSensorTimestamp(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// Start calls the method "Sensor.start".
//
// The returned bool will be false if there is no such method.
func (this Sensor) Start() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSensorStart(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StartFunc returns the method "Sensor.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Sensor) StartFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SensorStartFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "Sensor.stop".
//
// The returned bool will be false if there is no such method.
func (this Sensor) Stop() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSensorStop(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StopFunc returns the method "Sensor.stop".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Sensor) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SensorStopFunc(
			this.Ref(),
		),
	)
}

type SensorErrorEventInit struct {
	// Error is "SensorErrorEventInit.error"
	//
	// Required
	Error DOMException
	// Bubbles is "SensorErrorEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "SensorErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "SensorErrorEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SensorErrorEventInit with all fields set.
func (p SensorErrorEventInit) FromRef(ref js.Ref) SensorErrorEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SensorErrorEventInit SensorErrorEventInit [// SensorErrorEventInit] [0x14000a51400 0x14000a514a0 0x14000a515e0 0x14000a51720 0x14000a51540 0x14000a51680 0x14000a517c0] 0x14000a02390 {0 0}} in the application heap.
func (p SensorErrorEventInit) New() js.Ref {
	return bindings.SensorErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SensorErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.SensorErrorEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SensorErrorEventInit) Update(ref js.Ref) {
	bindings.SensorErrorEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewSensorErrorEvent(typ js.String, errorEventInitDict SensorErrorEventInit) SensorErrorEvent {
	return SensorErrorEvent{}.FromRef(
		bindings.NewSensorErrorEventBySensorErrorEvent(
			typ.Ref(),
			js.Pointer(&errorEventInitDict)),
	)
}

type SensorErrorEvent struct {
	Event
}

func (this SensorErrorEvent) Once() SensorErrorEvent {
	this.Ref().Once()
	return this
}

func (this SensorErrorEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this SensorErrorEvent) FromRef(ref js.Ref) SensorErrorEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this SensorErrorEvent) Free() {
	this.Ref().Free()
}

// Error returns the value of property "SensorErrorEvent.error".
//
// The returned bool will be false if there is no such property.
func (this SensorErrorEvent) Error() (DOMException, bool) {
	var _ok bool
	_ret := bindings.GetSensorErrorEventError(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMException{}.FromRef(_ret), _ok
}

func NewSequenceEffect(children js.Array[AnimationEffect], timing OneOf_Float64_EffectTiming) SequenceEffect {
	return SequenceEffect{}.FromRef(
		bindings.NewSequenceEffectBySequenceEffect(
			children.Ref(),
			timing.Ref()),
	)
}

func NewSequenceEffectBySequenceEffect1(children js.Array[AnimationEffect]) SequenceEffect {
	return SequenceEffect{}.FromRef(
		bindings.NewSequenceEffectBySequenceEffect1(
			children.Ref()),
	)
}

type SequenceEffect struct {
	GroupEffect
}

func (this SequenceEffect) Once() SequenceEffect {
	this.Ref().Once()
	return this
}

func (this SequenceEffect) Ref() js.Ref {
	return this.GroupEffect.Ref()
}

func (this SequenceEffect) FromRef(ref js.Ref) SequenceEffect {
	this.GroupEffect = this.GroupEffect.FromRef(ref)
	return this
}

func (this SequenceEffect) Free() {
	this.Ref().Free()
}

// Clone calls the method "SequenceEffect.clone".
//
// The returned bool will be false if there is no such method.
func (this SequenceEffect) Clone() (SequenceEffect, bool) {
	var _ok bool
	_ret := bindings.CallSequenceEffectClone(
		this.Ref(), js.Pointer(&_ok),
	)

	return SequenceEffect{}.FromRef(_ret), _ok
}

// CloneFunc returns the method "SequenceEffect.clone".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SequenceEffect) CloneFunc() (fn js.Func[func() SequenceEffect]) {
	return fn.FromRef(
		bindings.SequenceEffectCloneFunc(
			this.Ref(),
		),
	)
}

type ServiceWorkerGlobalScope struct {
	WorkerGlobalScope
}

func (this ServiceWorkerGlobalScope) Once() ServiceWorkerGlobalScope {
	this.Ref().Once()
	return this
}

func (this ServiceWorkerGlobalScope) Ref() js.Ref {
	return this.WorkerGlobalScope.Ref()
}

func (this ServiceWorkerGlobalScope) FromRef(ref js.Ref) ServiceWorkerGlobalScope {
	this.WorkerGlobalScope = this.WorkerGlobalScope.FromRef(ref)
	return this
}

func (this ServiceWorkerGlobalScope) Free() {
	this.Ref().Free()
}

// Clients returns the value of property "ServiceWorkerGlobalScope.clients".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerGlobalScope) Clients() (Clients, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerGlobalScopeClients(
		this.Ref(), js.Pointer(&_ok),
	)
	return Clients{}.FromRef(_ret), _ok
}

// Registration returns the value of property "ServiceWorkerGlobalScope.registration".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerGlobalScope) Registration() (ServiceWorkerRegistration, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerGlobalScopeRegistration(
		this.Ref(), js.Pointer(&_ok),
	)
	return ServiceWorkerRegistration{}.FromRef(_ret), _ok
}

// ServiceWorker returns the value of property "ServiceWorkerGlobalScope.serviceWorker".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerGlobalScope) ServiceWorker() (ServiceWorker, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerGlobalScopeServiceWorker(
		this.Ref(), js.Pointer(&_ok),
	)
	return ServiceWorker{}.FromRef(_ret), _ok
}

// CookieStore returns the value of property "ServiceWorkerGlobalScope.cookieStore".
//
// The returned bool will be false if there is no such property.
func (this ServiceWorkerGlobalScope) CookieStore() (CookieStore, bool) {
	var _ok bool
	_ret := bindings.GetServiceWorkerGlobalScopeCookieStore(
		this.Ref(), js.Pointer(&_ok),
	)
	return CookieStore{}.FromRef(_ret), _ok
}

// SkipWaiting calls the method "ServiceWorkerGlobalScope.skipWaiting".
//
// The returned bool will be false if there is no such method.
func (this ServiceWorkerGlobalScope) SkipWaiting() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallServiceWorkerGlobalScopeSkipWaiting(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SkipWaitingFunc returns the method "ServiceWorkerGlobalScope.skipWaiting".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ServiceWorkerGlobalScope) SkipWaitingFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.ServiceWorkerGlobalScopeSkipWaitingFunc(
			this.Ref(),
		),
	)
}

func NewShadowAnimation(source Animation, newTarget OneOf_Element_CSSPseudoElement) ShadowAnimation {
	return ShadowAnimation{}.FromRef(
		bindings.NewShadowAnimationByShadowAnimation(
			source.Ref(),
			newTarget.Ref()),
	)
}

type ShadowAnimation struct {
	Animation
}

func (this ShadowAnimation) Once() ShadowAnimation {
	this.Ref().Once()
	return this
}

func (this ShadowAnimation) Ref() js.Ref {
	return this.Animation.Ref()
}

func (this ShadowAnimation) FromRef(ref js.Ref) ShadowAnimation {
	this.Animation = this.Animation.FromRef(ref)
	return this
}

func (this ShadowAnimation) Free() {
	this.Ref().Free()
}

// SourceAnimation returns the value of property "ShadowAnimation.sourceAnimation".
//
// The returned bool will be false if there is no such property.
func (this ShadowAnimation) SourceAnimation() (Animation, bool) {
	var _ok bool
	_ret := bindings.GetShadowAnimationSourceAnimation(
		this.Ref(), js.Pointer(&_ok),
	)
	return Animation{}.FromRef(_ret), _ok
}

type SharedStorageSetMethodOptions struct {
	// IgnoreIfPresent is "SharedStorageSetMethodOptions.ignoreIfPresent"
	//
	// Optional, defaults to false.
	IgnoreIfPresent bool

	FFI_USE_IgnoreIfPresent bool // for IgnoreIfPresent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SharedStorageSetMethodOptions with all fields set.
func (p SharedStorageSetMethodOptions) FromRef(ref js.Ref) SharedStorageSetMethodOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SharedStorageSetMethodOptions SharedStorageSetMethodOptions [// SharedStorageSetMethodOptions] [0x14000a51900 0x14000a519a0] 0x14000a02408 {0 0}} in the application heap.
func (p SharedStorageSetMethodOptions) New() js.Ref {
	return bindings.SharedStorageSetMethodOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SharedStorageSetMethodOptions) UpdateFrom(ref js.Ref) {
	bindings.SharedStorageSetMethodOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SharedStorageSetMethodOptions) Update(ref js.Ref) {
	bindings.SharedStorageSetMethodOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type SharedStorage struct {
	ref js.Ref
}

func (this SharedStorage) Once() SharedStorage {
	this.Ref().Once()
	return this
}

func (this SharedStorage) Ref() js.Ref {
	return this.ref
}

func (this SharedStorage) FromRef(ref js.Ref) SharedStorage {
	this.ref = ref
	return this
}

func (this SharedStorage) Free() {
	this.Ref().Free()
}

// Set calls the method "SharedStorage.set".
//
// The returned bool will be false if there is no such method.
func (this SharedStorage) Set(key js.String, value js.String, options SharedStorageSetMethodOptions) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSharedStorageSet(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
		value.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// SetFunc returns the method "SharedStorage.set".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SharedStorage) SetFunc() (fn js.Func[func(key js.String, value js.String, options SharedStorageSetMethodOptions) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SharedStorageSetFunc(
			this.Ref(),
		),
	)
}

// Set1 calls the method "SharedStorage.set".
//
// The returned bool will be false if there is no such method.
func (this SharedStorage) Set1(key js.String, value js.String) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSharedStorageSet1(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
		value.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// Set1Func returns the method "SharedStorage.set".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SharedStorage) Set1Func() (fn js.Func[func(key js.String, value js.String) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SharedStorageSet1Func(
			this.Ref(),
		),
	)
}

// Append calls the method "SharedStorage.append".
//
// The returned bool will be false if there is no such method.
func (this SharedStorage) Append(key js.String, value js.String) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSharedStorageAppend(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
		value.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// AppendFunc returns the method "SharedStorage.append".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SharedStorage) AppendFunc() (fn js.Func[func(key js.String, value js.String) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SharedStorageAppendFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "SharedStorage.delete".
//
// The returned bool will be false if there is no such method.
func (this SharedStorage) Delete(key js.String) (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSharedStorageDelete(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// DeleteFunc returns the method "SharedStorage.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SharedStorage) DeleteFunc() (fn js.Func[func(key js.String) js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SharedStorageDeleteFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "SharedStorage.clear".
//
// The returned bool will be false if there is no such method.
func (this SharedStorage) Clear() (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.CallSharedStorageClear(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// ClearFunc returns the method "SharedStorage.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SharedStorage) ClearFunc() (fn js.Func[func() js.Promise[js.Any]]) {
	return fn.FromRef(
		bindings.SharedStorageClearFunc(
			this.Ref(),
		),
	)
}

type SharedStorageOperation struct {
	ref js.Ref
}

func (this SharedStorageOperation) Once() SharedStorageOperation {
	this.Ref().Once()
	return this
}

func (this SharedStorageOperation) Ref() js.Ref {
	return this.ref
}

func (this SharedStorageOperation) FromRef(ref js.Ref) SharedStorageOperation {
	this.ref = ref
	return this
}

func (this SharedStorageOperation) Free() {
	this.Ref().Free()
}

type SharedStorageOperationConstructorFunc func(this js.Ref, options SharedStorageRunOperationMethodOptions) js.Ref

func (fn SharedStorageOperationConstructorFunc) Register() js.Func[func(options SharedStorageRunOperationMethodOptions) SharedStorageOperation] {
	return js.RegisterCallback[func(options SharedStorageRunOperationMethodOptions) SharedStorageOperation](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SharedStorageOperationConstructorFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		SharedStorageRunOperationMethodOptions{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SharedStorageOperationConstructor[T any] struct {
	Fn  func(arg T, this js.Ref, options SharedStorageRunOperationMethodOptions) js.Ref
	Arg T
}

func (cb *SharedStorageOperationConstructor[T]) Register() js.Func[func(options SharedStorageRunOperationMethodOptions) SharedStorageOperation] {
	return js.RegisterCallback[func(options SharedStorageRunOperationMethodOptions) SharedStorageOperation](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SharedStorageOperationConstructor[T]) DispatchCallback(
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

		SharedStorageRunOperationMethodOptions{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SharedStorageRunOperation struct {
	SharedStorageOperation
}

func (this SharedStorageRunOperation) Once() SharedStorageRunOperation {
	this.Ref().Once()
	return this
}

func (this SharedStorageRunOperation) Ref() js.Ref {
	return this.SharedStorageOperation.Ref()
}

func (this SharedStorageRunOperation) FromRef(ref js.Ref) SharedStorageRunOperation {
	this.SharedStorageOperation = this.SharedStorageOperation.FromRef(ref)
	return this
}

func (this SharedStorageRunOperation) Free() {
	this.Ref().Free()
}

// Run calls the method "SharedStorageRunOperation.run".
//
// The returned bool will be false if there is no such method.
func (this SharedStorageRunOperation) Run(data js.Object) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallSharedStorageRunOperationRun(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RunFunc returns the method "SharedStorageRunOperation.run".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SharedStorageRunOperation) RunFunc() (fn js.Func[func(data js.Object) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.SharedStorageRunOperationRunFunc(
			this.Ref(),
		),
	)
}

type SharedStorageSelectURLOperation struct {
	SharedStorageOperation
}

func (this SharedStorageSelectURLOperation) Once() SharedStorageSelectURLOperation {
	this.Ref().Once()
	return this
}

func (this SharedStorageSelectURLOperation) Ref() js.Ref {
	return this.SharedStorageOperation.Ref()
}

func (this SharedStorageSelectURLOperation) FromRef(ref js.Ref) SharedStorageSelectURLOperation {
	this.SharedStorageOperation = this.SharedStorageOperation.FromRef(ref)
	return this
}

func (this SharedStorageSelectURLOperation) Free() {
	this.Ref().Free()
}

// Run calls the method "SharedStorageSelectURLOperation.run".
//
// The returned bool will be false if there is no such method.
func (this SharedStorageSelectURLOperation) Run(data js.Object, urls js.FrozenArray[SharedStorageUrlWithMetadata]) (js.Promise[js.Number[int32]], bool) {
	var _ok bool
	_ret := bindings.CallSharedStorageSelectURLOperationRun(
		this.Ref(), js.Pointer(&_ok),
		data.Ref(),
		urls.Ref(),
	)

	return js.Promise[js.Number[int32]]{}.FromRef(_ret), _ok
}

// RunFunc returns the method "SharedStorageSelectURLOperation.run".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SharedStorageSelectURLOperation) RunFunc() (fn js.Func[func(data js.Object, urls js.FrozenArray[SharedStorageUrlWithMetadata]) js.Promise[js.Number[int32]]]) {
	return fn.FromRef(
		bindings.SharedStorageSelectURLOperationRunFunc(
			this.Ref(),
		),
	)
}

type WorkletSharedStorage struct {
	SharedStorage
}

func (this WorkletSharedStorage) Once() WorkletSharedStorage {
	this.Ref().Once()
	return this
}

func (this WorkletSharedStorage) Ref() js.Ref {
	return this.SharedStorage.Ref()
}

func (this WorkletSharedStorage) FromRef(ref js.Ref) WorkletSharedStorage {
	this.SharedStorage = this.SharedStorage.FromRef(ref)
	return this
}

func (this WorkletSharedStorage) Free() {
	this.Ref().Free()
}

// Get calls the method "WorkletSharedStorage.get".
//
// The returned bool will be false if there is no such method.
func (this WorkletSharedStorage) Get(key js.String) (js.Promise[js.String], bool) {
	var _ok bool
	_ret := bindings.CallWorkletSharedStorageGet(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
	)

	return js.Promise[js.String]{}.FromRef(_ret), _ok
}

// GetFunc returns the method "WorkletSharedStorage.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkletSharedStorage) GetFunc() (fn js.Func[func(key js.String) js.Promise[js.String]]) {
	return fn.FromRef(
		bindings.WorkletSharedStorageGetFunc(
			this.Ref(),
		),
	)
}

// Length calls the method "WorkletSharedStorage.length".
//
// The returned bool will be false if there is no such method.
func (this WorkletSharedStorage) Length() (js.Promise[js.Number[uint32]], bool) {
	var _ok bool
	_ret := bindings.CallWorkletSharedStorageLength(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Number[uint32]]{}.FromRef(_ret), _ok
}

// LengthFunc returns the method "WorkletSharedStorage.length".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkletSharedStorage) LengthFunc() (fn js.Func[func() js.Promise[js.Number[uint32]]]) {
	return fn.FromRef(
		bindings.WorkletSharedStorageLengthFunc(
			this.Ref(),
		),
	)
}

// RemainingBudget calls the method "WorkletSharedStorage.remainingBudget".
//
// The returned bool will be false if there is no such method.
func (this WorkletSharedStorage) RemainingBudget() (js.Promise[js.Number[float64]], bool) {
	var _ok bool
	_ret := bindings.CallWorkletSharedStorageRemainingBudget(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Number[float64]]{}.FromRef(_ret), _ok
}

// RemainingBudgetFunc returns the method "WorkletSharedStorage.remainingBudget".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this WorkletSharedStorage) RemainingBudgetFunc() (fn js.Func[func() js.Promise[js.Number[float64]]]) {
	return fn.FromRef(
		bindings.WorkletSharedStorageRemainingBudgetFunc(
			this.Ref(),
		),
	)
}

type SharedStorageWorkletGlobalScope struct {
	WorkletGlobalScope
}

func (this SharedStorageWorkletGlobalScope) Once() SharedStorageWorkletGlobalScope {
	this.Ref().Once()
	return this
}

func (this SharedStorageWorkletGlobalScope) Ref() js.Ref {
	return this.WorkletGlobalScope.Ref()
}

func (this SharedStorageWorkletGlobalScope) FromRef(ref js.Ref) SharedStorageWorkletGlobalScope {
	this.WorkletGlobalScope = this.WorkletGlobalScope.FromRef(ref)
	return this
}

func (this SharedStorageWorkletGlobalScope) Free() {
	this.Ref().Free()
}

// SharedStorage returns the value of property "SharedStorageWorkletGlobalScope.sharedStorage".
//
// The returned bool will be false if there is no such property.
func (this SharedStorageWorkletGlobalScope) SharedStorage() (WorkletSharedStorage, bool) {
	var _ok bool
	_ret := bindings.GetSharedStorageWorkletGlobalScopeSharedStorage(
		this.Ref(), js.Pointer(&_ok),
	)
	return WorkletSharedStorage{}.FromRef(_ret), _ok
}

// Register calls the method "SharedStorageWorkletGlobalScope.register".
//
// The returned bool will be false if there is no such method.
func (this SharedStorageWorkletGlobalScope) Register(name js.String, operationCtor js.Func[func(options SharedStorageRunOperationMethodOptions) SharedStorageOperation]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSharedStorageWorkletGlobalScopeRegister(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		operationCtor.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RegisterFunc returns the method "SharedStorageWorkletGlobalScope.register".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SharedStorageWorkletGlobalScope) RegisterFunc() (fn js.Func[func(name js.String, operationCtor js.Func[func(options SharedStorageRunOperationMethodOptions) SharedStorageOperation])]) {
	return fn.FromRef(
		bindings.SharedStorageWorkletGlobalScopeRegisterFunc(
			this.Ref(),
		),
	)
}

type OneOf_String_WorkerOptions struct {
	ref js.Ref
}

func (x OneOf_String_WorkerOptions) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_WorkerOptions) Free() {
	x.ref.Free()
}

func (x OneOf_String_WorkerOptions) FromRef(ref js.Ref) OneOf_String_WorkerOptions {
	return OneOf_String_WorkerOptions{
		ref: ref,
	}
}

func (x OneOf_String_WorkerOptions) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_WorkerOptions) WorkerOptions() WorkerOptions {
	var ret WorkerOptions
	ret.UpdateFrom(x.ref)
	return ret
}

func NewSharedWorker(scriptURL js.String, options OneOf_String_WorkerOptions) SharedWorker {
	return SharedWorker{}.FromRef(
		bindings.NewSharedWorkerBySharedWorker(
			scriptURL.Ref(),
			options.Ref()),
	)
}

func NewSharedWorkerBySharedWorker1(scriptURL js.String) SharedWorker {
	return SharedWorker{}.FromRef(
		bindings.NewSharedWorkerBySharedWorker1(
			scriptURL.Ref()),
	)
}

type SharedWorker struct {
	EventTarget
}

func (this SharedWorker) Once() SharedWorker {
	this.Ref().Once()
	return this
}

func (this SharedWorker) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this SharedWorker) FromRef(ref js.Ref) SharedWorker {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this SharedWorker) Free() {
	this.Ref().Free()
}

// Port returns the value of property "SharedWorker.port".
//
// The returned bool will be false if there is no such property.
func (this SharedWorker) Port() (MessagePort, bool) {
	var _ok bool
	_ret := bindings.GetSharedWorkerPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return MessagePort{}.FromRef(_ret), _ok
}

type SharedWorkerGlobalScope struct {
	WorkerGlobalScope
}

func (this SharedWorkerGlobalScope) Once() SharedWorkerGlobalScope {
	this.Ref().Once()
	return this
}

func (this SharedWorkerGlobalScope) Ref() js.Ref {
	return this.WorkerGlobalScope.Ref()
}

func (this SharedWorkerGlobalScope) FromRef(ref js.Ref) SharedWorkerGlobalScope {
	this.WorkerGlobalScope = this.WorkerGlobalScope.FromRef(ref)
	return this
}

func (this SharedWorkerGlobalScope) Free() {
	this.Ref().Free()
}

// Name returns the value of property "SharedWorkerGlobalScope.name".
//
// The returned bool will be false if there is no such property.
func (this SharedWorkerGlobalScope) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSharedWorkerGlobalScopeName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Close calls the method "SharedWorkerGlobalScope.close".
//
// The returned bool will be false if there is no such method.
func (this SharedWorkerGlobalScope) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSharedWorkerGlobalScopeClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "SharedWorkerGlobalScope.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SharedWorkerGlobalScope) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SharedWorkerGlobalScopeCloseFunc(
			this.Ref(),
		),
	)
}

type SmallCryptoKeyID uint64

type SpeechGrammar struct {
	ref js.Ref
}

func (this SpeechGrammar) Once() SpeechGrammar {
	this.Ref().Once()
	return this
}

func (this SpeechGrammar) Ref() js.Ref {
	return this.ref
}

func (this SpeechGrammar) FromRef(ref js.Ref) SpeechGrammar {
	this.ref = ref
	return this
}

func (this SpeechGrammar) Free() {
	this.Ref().Free()
}

// Src returns the value of property "SpeechGrammar.src".
//
// The returned bool will be false if there is no such property.
func (this SpeechGrammar) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSpeechGrammarSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src sets the value of property "SpeechGrammar.src" to val.
//
// It returns false if the property cannot be set.
func (this SpeechGrammar) SetSrc(val js.String) bool {
	return js.True == bindings.SetSpeechGrammarSrc(
		this.Ref(),
		val.Ref(),
	)
}

// Weight returns the value of property "SpeechGrammar.weight".
//
// The returned bool will be false if there is no such property.
func (this SpeechGrammar) Weight() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechGrammarWeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Weight sets the value of property "SpeechGrammar.weight" to val.
//
// It returns false if the property cannot be set.
func (this SpeechGrammar) SetWeight(val float32) bool {
	return js.True == bindings.SetSpeechGrammarWeight(
		this.Ref(),
		float32(val),
	)
}

type SpeechGrammarList struct {
	ref js.Ref
}

func (this SpeechGrammarList) Once() SpeechGrammarList {
	this.Ref().Once()
	return this
}

func (this SpeechGrammarList) Ref() js.Ref {
	return this.ref
}

func (this SpeechGrammarList) FromRef(ref js.Ref) SpeechGrammarList {
	this.ref = ref
	return this
}

func (this SpeechGrammarList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "SpeechGrammarList.length".
//
// The returned bool will be false if there is no such property.
func (this SpeechGrammarList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechGrammarListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "SpeechGrammarList.item".
//
// The returned bool will be false if there is no such method.
func (this SpeechGrammarList) Item(index uint32) (SpeechGrammar, bool) {
	var _ok bool
	_ret := bindings.CallSpeechGrammarListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return SpeechGrammar{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "SpeechGrammarList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechGrammarList) ItemFunc() (fn js.Func[func(index uint32) SpeechGrammar]) {
	return fn.FromRef(
		bindings.SpeechGrammarListItemFunc(
			this.Ref(),
		),
	)
}

// AddFromURI calls the method "SpeechGrammarList.addFromURI".
//
// The returned bool will be false if there is no such method.
func (this SpeechGrammarList) AddFromURI(src js.String, weight float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSpeechGrammarListAddFromURI(
		this.Ref(), js.Pointer(&_ok),
		src.Ref(),
		float32(weight),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddFromURIFunc returns the method "SpeechGrammarList.addFromURI".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechGrammarList) AddFromURIFunc() (fn js.Func[func(src js.String, weight float32)]) {
	return fn.FromRef(
		bindings.SpeechGrammarListAddFromURIFunc(
			this.Ref(),
		),
	)
}

// AddFromURI1 calls the method "SpeechGrammarList.addFromURI".
//
// The returned bool will be false if there is no such method.
func (this SpeechGrammarList) AddFromURI1(src js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSpeechGrammarListAddFromURI1(
		this.Ref(), js.Pointer(&_ok),
		src.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddFromURI1Func returns the method "SpeechGrammarList.addFromURI".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechGrammarList) AddFromURI1Func() (fn js.Func[func(src js.String)]) {
	return fn.FromRef(
		bindings.SpeechGrammarListAddFromURI1Func(
			this.Ref(),
		),
	)
}

// AddFromString calls the method "SpeechGrammarList.addFromString".
//
// The returned bool will be false if there is no such method.
func (this SpeechGrammarList) AddFromString(string js.String, weight float32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSpeechGrammarListAddFromString(
		this.Ref(), js.Pointer(&_ok),
		string.Ref(),
		float32(weight),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddFromStringFunc returns the method "SpeechGrammarList.addFromString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechGrammarList) AddFromStringFunc() (fn js.Func[func(string js.String, weight float32)]) {
	return fn.FromRef(
		bindings.SpeechGrammarListAddFromStringFunc(
			this.Ref(),
		),
	)
}

// AddFromString1 calls the method "SpeechGrammarList.addFromString".
//
// The returned bool will be false if there is no such method.
func (this SpeechGrammarList) AddFromString1(string js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSpeechGrammarListAddFromString1(
		this.Ref(), js.Pointer(&_ok),
		string.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddFromString1Func returns the method "SpeechGrammarList.addFromString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechGrammarList) AddFromString1Func() (fn js.Func[func(string js.String)]) {
	return fn.FromRef(
		bindings.SpeechGrammarListAddFromString1Func(
			this.Ref(),
		),
	)
}

type SpeechRecognition struct {
	EventTarget
}

func (this SpeechRecognition) Once() SpeechRecognition {
	this.Ref().Once()
	return this
}

func (this SpeechRecognition) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this SpeechRecognition) FromRef(ref js.Ref) SpeechRecognition {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this SpeechRecognition) Free() {
	this.Ref().Free()
}

// Grammars returns the value of property "SpeechRecognition.grammars".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognition) Grammars() (SpeechGrammarList, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionGrammars(
		this.Ref(), js.Pointer(&_ok),
	)
	return SpeechGrammarList{}.FromRef(_ret), _ok
}

// Grammars sets the value of property "SpeechRecognition.grammars" to val.
//
// It returns false if the property cannot be set.
func (this SpeechRecognition) SetGrammars(val SpeechGrammarList) bool {
	return js.True == bindings.SetSpeechRecognitionGrammars(
		this.Ref(),
		val.Ref(),
	)
}

// Lang returns the value of property "SpeechRecognition.lang".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognition) Lang() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionLang(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Lang sets the value of property "SpeechRecognition.lang" to val.
//
// It returns false if the property cannot be set.
func (this SpeechRecognition) SetLang(val js.String) bool {
	return js.True == bindings.SetSpeechRecognitionLang(
		this.Ref(),
		val.Ref(),
	)
}

// Continuous returns the value of property "SpeechRecognition.continuous".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognition) Continuous() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionContinuous(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Continuous sets the value of property "SpeechRecognition.continuous" to val.
//
// It returns false if the property cannot be set.
func (this SpeechRecognition) SetContinuous(val bool) bool {
	return js.True == bindings.SetSpeechRecognitionContinuous(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// InterimResults returns the value of property "SpeechRecognition.interimResults".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognition) InterimResults() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionInterimResults(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// InterimResults sets the value of property "SpeechRecognition.interimResults" to val.
//
// It returns false if the property cannot be set.
func (this SpeechRecognition) SetInterimResults(val bool) bool {
	return js.True == bindings.SetSpeechRecognitionInterimResults(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// MaxAlternatives returns the value of property "SpeechRecognition.maxAlternatives".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognition) MaxAlternatives() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionMaxAlternatives(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// MaxAlternatives sets the value of property "SpeechRecognition.maxAlternatives" to val.
//
// It returns false if the property cannot be set.
func (this SpeechRecognition) SetMaxAlternatives(val uint32) bool {
	return js.True == bindings.SetSpeechRecognitionMaxAlternatives(
		this.Ref(),
		uint32(val),
	)
}

// Start calls the method "SpeechRecognition.start".
//
// The returned bool will be false if there is no such method.
func (this SpeechRecognition) Start() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSpeechRecognitionStart(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StartFunc returns the method "SpeechRecognition.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechRecognition) StartFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SpeechRecognitionStartFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "SpeechRecognition.stop".
//
// The returned bool will be false if there is no such method.
func (this SpeechRecognition) Stop() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSpeechRecognitionStop(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StopFunc returns the method "SpeechRecognition.stop".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechRecognition) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SpeechRecognitionStopFunc(
			this.Ref(),
		),
	)
}

// Abort calls the method "SpeechRecognition.abort".
//
// The returned bool will be false if there is no such method.
func (this SpeechRecognition) Abort() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallSpeechRecognitionAbort(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AbortFunc returns the method "SpeechRecognition.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechRecognition) AbortFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.SpeechRecognitionAbortFunc(
			this.Ref(),
		),
	)
}

type SpeechRecognitionAlternative struct {
	ref js.Ref
}

func (this SpeechRecognitionAlternative) Once() SpeechRecognitionAlternative {
	this.Ref().Once()
	return this
}

func (this SpeechRecognitionAlternative) Ref() js.Ref {
	return this.ref
}

func (this SpeechRecognitionAlternative) FromRef(ref js.Ref) SpeechRecognitionAlternative {
	this.ref = ref
	return this
}

func (this SpeechRecognitionAlternative) Free() {
	this.Ref().Free()
}

// Transcript returns the value of property "SpeechRecognitionAlternative.transcript".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognitionAlternative) Transcript() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionAlternativeTranscript(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Confidence returns the value of property "SpeechRecognitionAlternative.confidence".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognitionAlternative) Confidence() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionAlternativeConfidence(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

type SpeechRecognitionErrorCode uint32

const (
	_ SpeechRecognitionErrorCode = iota

	SpeechRecognitionErrorCode_NO_SPEECH
	SpeechRecognitionErrorCode_ABORTED
	SpeechRecognitionErrorCode_AUDIO_CAPTURE
	SpeechRecognitionErrorCode_NETWORK
	SpeechRecognitionErrorCode_NOT_ALLOWED
	SpeechRecognitionErrorCode_SERVICE_NOT_ALLOWED
	SpeechRecognitionErrorCode_BAD_GRAMMAR
	SpeechRecognitionErrorCode_LANGUAGE_NOT_SUPPORTED
)

func (SpeechRecognitionErrorCode) FromRef(str js.Ref) SpeechRecognitionErrorCode {
	return SpeechRecognitionErrorCode(bindings.ConstOfSpeechRecognitionErrorCode(str))
}

func (x SpeechRecognitionErrorCode) String() (string, bool) {
	switch x {
	case SpeechRecognitionErrorCode_NO_SPEECH:
		return "no-speech", true
	case SpeechRecognitionErrorCode_ABORTED:
		return "aborted", true
	case SpeechRecognitionErrorCode_AUDIO_CAPTURE:
		return "audio-capture", true
	case SpeechRecognitionErrorCode_NETWORK:
		return "network", true
	case SpeechRecognitionErrorCode_NOT_ALLOWED:
		return "not-allowed", true
	case SpeechRecognitionErrorCode_SERVICE_NOT_ALLOWED:
		return "service-not-allowed", true
	case SpeechRecognitionErrorCode_BAD_GRAMMAR:
		return "bad-grammar", true
	case SpeechRecognitionErrorCode_LANGUAGE_NOT_SUPPORTED:
		return "language-not-supported", true
	default:
		return "", false
	}
}

type SpeechRecognitionErrorEventInit struct {
	// Error is "SpeechRecognitionErrorEventInit.error"
	//
	// Required
	Error SpeechRecognitionErrorCode
	// Message is "SpeechRecognitionErrorEventInit.message"
	//
	// Optional, defaults to "".
	Message js.String
	// Bubbles is "SpeechRecognitionErrorEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "SpeechRecognitionErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "SpeechRecognitionErrorEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SpeechRecognitionErrorEventInit with all fields set.
func (p SpeechRecognitionErrorEventInit) FromRef(ref js.Ref) SpeechRecognitionErrorEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SpeechRecognitionErrorEventInit SpeechRecognitionErrorEventInit [// SpeechRecognitionErrorEventInit] [0x14000a51c20 0x14000a51cc0 0x14000a51d60 0x14000a51ea0 0x14000a64000 0x14000a51e00 0x14000a51f40 0x14000a640a0] 0x14000a025b8 {0 0}} in the application heap.
func (p SpeechRecognitionErrorEventInit) New() js.Ref {
	return bindings.SpeechRecognitionErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SpeechRecognitionErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.SpeechRecognitionErrorEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SpeechRecognitionErrorEventInit) Update(ref js.Ref) {
	bindings.SpeechRecognitionErrorEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewSpeechRecognitionErrorEvent(typ js.String, eventInitDict SpeechRecognitionErrorEventInit) SpeechRecognitionErrorEvent {
	return SpeechRecognitionErrorEvent{}.FromRef(
		bindings.NewSpeechRecognitionErrorEventBySpeechRecognitionErrorEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type SpeechRecognitionErrorEvent struct {
	Event
}

func (this SpeechRecognitionErrorEvent) Once() SpeechRecognitionErrorEvent {
	this.Ref().Once()
	return this
}

func (this SpeechRecognitionErrorEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this SpeechRecognitionErrorEvent) FromRef(ref js.Ref) SpeechRecognitionErrorEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this SpeechRecognitionErrorEvent) Free() {
	this.Ref().Free()
}

// Error returns the value of property "SpeechRecognitionErrorEvent.error".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognitionErrorEvent) Error() (SpeechRecognitionErrorCode, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionErrorEventError(
		this.Ref(), js.Pointer(&_ok),
	)
	return SpeechRecognitionErrorCode(_ret), _ok
}

// Message returns the value of property "SpeechRecognitionErrorEvent.message".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognitionErrorEvent) Message() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionErrorEventMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type SpeechRecognitionResult struct {
	ref js.Ref
}

func (this SpeechRecognitionResult) Once() SpeechRecognitionResult {
	this.Ref().Once()
	return this
}

func (this SpeechRecognitionResult) Ref() js.Ref {
	return this.ref
}

func (this SpeechRecognitionResult) FromRef(ref js.Ref) SpeechRecognitionResult {
	this.ref = ref
	return this
}

func (this SpeechRecognitionResult) Free() {
	this.Ref().Free()
}

// Length returns the value of property "SpeechRecognitionResult.length".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognitionResult) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionResultLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// IsFinal returns the value of property "SpeechRecognitionResult.isFinal".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognitionResult) IsFinal() (bool, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionResultIsFinal(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Item calls the method "SpeechRecognitionResult.item".
//
// The returned bool will be false if there is no such method.
func (this SpeechRecognitionResult) Item(index uint32) (SpeechRecognitionAlternative, bool) {
	var _ok bool
	_ret := bindings.CallSpeechRecognitionResultItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return SpeechRecognitionAlternative{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "SpeechRecognitionResult.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechRecognitionResult) ItemFunc() (fn js.Func[func(index uint32) SpeechRecognitionAlternative]) {
	return fn.FromRef(
		bindings.SpeechRecognitionResultItemFunc(
			this.Ref(),
		),
	)
}

type SpeechRecognitionResultList struct {
	ref js.Ref
}

func (this SpeechRecognitionResultList) Once() SpeechRecognitionResultList {
	this.Ref().Once()
	return this
}

func (this SpeechRecognitionResultList) Ref() js.Ref {
	return this.ref
}

func (this SpeechRecognitionResultList) FromRef(ref js.Ref) SpeechRecognitionResultList {
	this.ref = ref
	return this
}

func (this SpeechRecognitionResultList) Free() {
	this.Ref().Free()
}

// Length returns the value of property "SpeechRecognitionResultList.length".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognitionResultList) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionResultListLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Item calls the method "SpeechRecognitionResultList.item".
//
// The returned bool will be false if there is no such method.
func (this SpeechRecognitionResultList) Item(index uint32) (SpeechRecognitionResult, bool) {
	var _ok bool
	_ret := bindings.CallSpeechRecognitionResultListItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return SpeechRecognitionResult{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "SpeechRecognitionResultList.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this SpeechRecognitionResultList) ItemFunc() (fn js.Func[func(index uint32) SpeechRecognitionResult]) {
	return fn.FromRef(
		bindings.SpeechRecognitionResultListItemFunc(
			this.Ref(),
		),
	)
}

type SpeechRecognitionEventInit struct {
	// ResultIndex is "SpeechRecognitionEventInit.resultIndex"
	//
	// Optional, defaults to 0.
	ResultIndex uint32
	// Results is "SpeechRecognitionEventInit.results"
	//
	// Required
	Results SpeechRecognitionResultList
	// Bubbles is "SpeechRecognitionEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "SpeechRecognitionEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "SpeechRecognitionEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_ResultIndex bool // for ResultIndex.
	FFI_USE_Bubbles     bool // for Bubbles.
	FFI_USE_Cancelable  bool // for Cancelable.
	FFI_USE_Composed    bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SpeechRecognitionEventInit with all fields set.
func (p SpeechRecognitionEventInit) FromRef(ref js.Ref) SpeechRecognitionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SpeechRecognitionEventInit SpeechRecognitionEventInit [// SpeechRecognitionEventInit] [0x14000a641e0 0x14000a643c0 0x14000a64460 0x14000a645a0 0x14000a646e0 0x14000a64280 0x14000a64500 0x14000a64640 0x14000a64780] 0x14000a025e8 {0 0}} in the application heap.
func (p SpeechRecognitionEventInit) New() js.Ref {
	return bindings.SpeechRecognitionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SpeechRecognitionEventInit) UpdateFrom(ref js.Ref) {
	bindings.SpeechRecognitionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SpeechRecognitionEventInit) Update(ref js.Ref) {
	bindings.SpeechRecognitionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewSpeechRecognitionEvent(typ js.String, eventInitDict SpeechRecognitionEventInit) SpeechRecognitionEvent {
	return SpeechRecognitionEvent{}.FromRef(
		bindings.NewSpeechRecognitionEventBySpeechRecognitionEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type SpeechRecognitionEvent struct {
	Event
}

func (this SpeechRecognitionEvent) Once() SpeechRecognitionEvent {
	this.Ref().Once()
	return this
}

func (this SpeechRecognitionEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this SpeechRecognitionEvent) FromRef(ref js.Ref) SpeechRecognitionEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this SpeechRecognitionEvent) Free() {
	this.Ref().Free()
}

// ResultIndex returns the value of property "SpeechRecognitionEvent.resultIndex".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognitionEvent) ResultIndex() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionEventResultIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Results returns the value of property "SpeechRecognitionEvent.results".
//
// The returned bool will be false if there is no such property.
func (this SpeechRecognitionEvent) Results() (SpeechRecognitionResultList, bool) {
	var _ok bool
	_ret := bindings.GetSpeechRecognitionEventResults(
		this.Ref(), js.Pointer(&_ok),
	)
	return SpeechRecognitionResultList{}.FromRef(_ret), _ok
}

type SpeechSynthesisErrorCode uint32

const (
	_ SpeechSynthesisErrorCode = iota

	SpeechSynthesisErrorCode_CANCELED
	SpeechSynthesisErrorCode_INTERRUPTED
	SpeechSynthesisErrorCode_AUDIO_BUSY
	SpeechSynthesisErrorCode_AUDIO_HARDWARE
	SpeechSynthesisErrorCode_NETWORK
	SpeechSynthesisErrorCode_SYNTHESIS_UNAVAILABLE
	SpeechSynthesisErrorCode_SYNTHESIS_FAILED
	SpeechSynthesisErrorCode_LANGUAGE_UNAVAILABLE
	SpeechSynthesisErrorCode_VOICE_UNAVAILABLE
	SpeechSynthesisErrorCode_TEXT_TOO_LONG
	SpeechSynthesisErrorCode_INVALID_ARGUMENT
	SpeechSynthesisErrorCode_NOT_ALLOWED
)

func (SpeechSynthesisErrorCode) FromRef(str js.Ref) SpeechSynthesisErrorCode {
	return SpeechSynthesisErrorCode(bindings.ConstOfSpeechSynthesisErrorCode(str))
}

func (x SpeechSynthesisErrorCode) String() (string, bool) {
	switch x {
	case SpeechSynthesisErrorCode_CANCELED:
		return "canceled", true
	case SpeechSynthesisErrorCode_INTERRUPTED:
		return "interrupted", true
	case SpeechSynthesisErrorCode_AUDIO_BUSY:
		return "audio-busy", true
	case SpeechSynthesisErrorCode_AUDIO_HARDWARE:
		return "audio-hardware", true
	case SpeechSynthesisErrorCode_NETWORK:
		return "network", true
	case SpeechSynthesisErrorCode_SYNTHESIS_UNAVAILABLE:
		return "synthesis-unavailable", true
	case SpeechSynthesisErrorCode_SYNTHESIS_FAILED:
		return "synthesis-failed", true
	case SpeechSynthesisErrorCode_LANGUAGE_UNAVAILABLE:
		return "language-unavailable", true
	case SpeechSynthesisErrorCode_VOICE_UNAVAILABLE:
		return "voice-unavailable", true
	case SpeechSynthesisErrorCode_TEXT_TOO_LONG:
		return "text-too-long", true
	case SpeechSynthesisErrorCode_INVALID_ARGUMENT:
		return "invalid-argument", true
	case SpeechSynthesisErrorCode_NOT_ALLOWED:
		return "not-allowed", true
	default:
		return "", false
	}
}

type SpeechSynthesisErrorEventInit struct {
	// Error is "SpeechSynthesisErrorEventInit.error"
	//
	// Required
	Error SpeechSynthesisErrorCode
	// Utterance is "SpeechSynthesisErrorEventInit.utterance"
	//
	// Required
	Utterance SpeechSynthesisUtterance
	// CharIndex is "SpeechSynthesisErrorEventInit.charIndex"
	//
	// Optional, defaults to 0.
	CharIndex uint32
	// CharLength is "SpeechSynthesisErrorEventInit.charLength"
	//
	// Optional, defaults to 0.
	CharLength uint32
	// ElapsedTime is "SpeechSynthesisErrorEventInit.elapsedTime"
	//
	// Optional, defaults to 0.
	ElapsedTime float32
	// Name is "SpeechSynthesisErrorEventInit.name"
	//
	// Optional, defaults to "".
	Name js.String
	// Bubbles is "SpeechSynthesisErrorEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "SpeechSynthesisErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "SpeechSynthesisErrorEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_CharIndex   bool // for CharIndex.
	FFI_USE_CharLength  bool // for CharLength.
	FFI_USE_ElapsedTime bool // for ElapsedTime.
	FFI_USE_Bubbles     bool // for Bubbles.
	FFI_USE_Cancelable  bool // for Cancelable.
	FFI_USE_Composed    bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SpeechSynthesisErrorEventInit with all fields set.
func (p SpeechSynthesisErrorEventInit) FromRef(ref js.Ref) SpeechSynthesisErrorEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SpeechSynthesisErrorEventInit SpeechSynthesisErrorEventInit [// SpeechSynthesisErrorEventInit] [0x14000a648c0 0x14000a64960 0x14000a64a00 0x14000a64b40 0x14000a64c80 0x14000a64dc0 0x14000a64e60 0x14000a64fa0 0x14000a650e0 0x14000a64aa0 0x14000a64be0 0x14000a64d20 0x14000a64f00 0x14000a65040 0x14000a65180] 0x14000a02828 {0 0}} in the application heap.
func (p SpeechSynthesisErrorEventInit) New() js.Ref {
	return bindings.SpeechSynthesisErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SpeechSynthesisErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.SpeechSynthesisErrorEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SpeechSynthesisErrorEventInit) Update(ref js.Ref) {
	bindings.SpeechSynthesisErrorEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewSpeechSynthesisErrorEvent(typ js.String, eventInitDict SpeechSynthesisErrorEventInit) SpeechSynthesisErrorEvent {
	return SpeechSynthesisErrorEvent{}.FromRef(
		bindings.NewSpeechSynthesisErrorEventBySpeechSynthesisErrorEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type SpeechSynthesisErrorEvent struct {
	SpeechSynthesisEvent
}

func (this SpeechSynthesisErrorEvent) Once() SpeechSynthesisErrorEvent {
	this.Ref().Once()
	return this
}

func (this SpeechSynthesisErrorEvent) Ref() js.Ref {
	return this.SpeechSynthesisEvent.Ref()
}

func (this SpeechSynthesisErrorEvent) FromRef(ref js.Ref) SpeechSynthesisErrorEvent {
	this.SpeechSynthesisEvent = this.SpeechSynthesisEvent.FromRef(ref)
	return this
}

func (this SpeechSynthesisErrorEvent) Free() {
	this.Ref().Free()
}

// Error returns the value of property "SpeechSynthesisErrorEvent.error".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisErrorEvent) Error() (SpeechSynthesisErrorCode, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisErrorEventError(
		this.Ref(), js.Pointer(&_ok),
	)
	return SpeechSynthesisErrorCode(_ret), _ok
}

type SpeechSynthesisEventInit struct {
	// Utterance is "SpeechSynthesisEventInit.utterance"
	//
	// Required
	Utterance SpeechSynthesisUtterance
	// CharIndex is "SpeechSynthesisEventInit.charIndex"
	//
	// Optional, defaults to 0.
	CharIndex uint32
	// CharLength is "SpeechSynthesisEventInit.charLength"
	//
	// Optional, defaults to 0.
	CharLength uint32
	// ElapsedTime is "SpeechSynthesisEventInit.elapsedTime"
	//
	// Optional, defaults to 0.
	ElapsedTime float32
	// Name is "SpeechSynthesisEventInit.name"
	//
	// Optional, defaults to "".
	Name js.String
	// Bubbles is "SpeechSynthesisEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "SpeechSynthesisEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "SpeechSynthesisEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_CharIndex   bool // for CharIndex.
	FFI_USE_CharLength  bool // for CharLength.
	FFI_USE_ElapsedTime bool // for ElapsedTime.
	FFI_USE_Bubbles     bool // for Bubbles.
	FFI_USE_Cancelable  bool // for Cancelable.
	FFI_USE_Composed    bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SpeechSynthesisEventInit with all fields set.
func (p SpeechSynthesisEventInit) FromRef(ref js.Ref) SpeechSynthesisEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SpeechSynthesisEventInit SpeechSynthesisEventInit [// SpeechSynthesisEventInit] [0x14000a65220 0x14000a652c0 0x14000a65400 0x14000a65540 0x14000a65680 0x14000a65720 0x14000a65860 0x14000a659a0 0x14000a65360 0x14000a654a0 0x14000a655e0 0x14000a657c0 0x14000a65900 0x14000a65a40] 0x14000a028b8 {0 0}} in the application heap.
func (p SpeechSynthesisEventInit) New() js.Ref {
	return bindings.SpeechSynthesisEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SpeechSynthesisEventInit) UpdateFrom(ref js.Ref) {
	bindings.SpeechSynthesisEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SpeechSynthesisEventInit) Update(ref js.Ref) {
	bindings.SpeechSynthesisEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewSpeechSynthesisEvent(typ js.String, eventInitDict SpeechSynthesisEventInit) SpeechSynthesisEvent {
	return SpeechSynthesisEvent{}.FromRef(
		bindings.NewSpeechSynthesisEventBySpeechSynthesisEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type SpeechSynthesisEvent struct {
	Event
}

func (this SpeechSynthesisEvent) Once() SpeechSynthesisEvent {
	this.Ref().Once()
	return this
}

func (this SpeechSynthesisEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this SpeechSynthesisEvent) FromRef(ref js.Ref) SpeechSynthesisEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this SpeechSynthesisEvent) Free() {
	this.Ref().Free()
}

// Utterance returns the value of property "SpeechSynthesisEvent.utterance".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisEvent) Utterance() (SpeechSynthesisUtterance, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisEventUtterance(
		this.Ref(), js.Pointer(&_ok),
	)
	return SpeechSynthesisUtterance{}.FromRef(_ret), _ok
}

// CharIndex returns the value of property "SpeechSynthesisEvent.charIndex".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisEvent) CharIndex() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisEventCharIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// CharLength returns the value of property "SpeechSynthesisEvent.charLength".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisEvent) CharLength() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisEventCharLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ElapsedTime returns the value of property "SpeechSynthesisEvent.elapsedTime".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisEvent) ElapsedTime() (float32, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisEventElapsedTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return float32(_ret), _ok
}

// Name returns the value of property "SpeechSynthesisEvent.name".
//
// The returned bool will be false if there is no such property.
func (this SpeechSynthesisEvent) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetSpeechSynthesisEventName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type StorageEventInit struct {
	// Key is "StorageEventInit.key"
	//
	// Optional, defaults to null.
	Key js.String
	// OldValue is "StorageEventInit.oldValue"
	//
	// Optional, defaults to null.
	OldValue js.String
	// NewValue is "StorageEventInit.newValue"
	//
	// Optional, defaults to null.
	NewValue js.String
	// Url is "StorageEventInit.url"
	//
	// Optional, defaults to "".
	Url js.String
	// StorageArea is "StorageEventInit.storageArea"
	//
	// Optional, defaults to null.
	StorageArea Storage
	// Bubbles is "StorageEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "StorageEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "StorageEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a StorageEventInit with all fields set.
func (p StorageEventInit) FromRef(ref js.Ref) StorageEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 StorageEventInit StorageEventInit [// StorageEventInit] [0x14000a65b80 0x14000a65c20 0x14000a65cc0 0x14000a65d60 0x14000a65e00 0x14000a65ea0 0x14000a6a000 0x14000a6a140 0x14000a65f40 0x14000a6a0a0 0x14000a6a1e0] 0x14000a02978 {0 0}} in the application heap.
func (p StorageEventInit) New() js.Ref {
	return bindings.StorageEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p StorageEventInit) UpdateFrom(ref js.Ref) {
	bindings.StorageEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p StorageEventInit) Update(ref js.Ref) {
	bindings.StorageEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewStorageEvent(typ js.String, eventInitDict StorageEventInit) StorageEvent {
	return StorageEvent{}.FromRef(
		bindings.NewStorageEventByStorageEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewStorageEventByStorageEvent1(typ js.String) StorageEvent {
	return StorageEvent{}.FromRef(
		bindings.NewStorageEventByStorageEvent1(
			typ.Ref()),
	)
}

type StorageEvent struct {
	Event
}

func (this StorageEvent) Once() StorageEvent {
	this.Ref().Once()
	return this
}

func (this StorageEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this StorageEvent) FromRef(ref js.Ref) StorageEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this StorageEvent) Free() {
	this.Ref().Free()
}

// Key returns the value of property "StorageEvent.key".
//
// The returned bool will be false if there is no such property.
func (this StorageEvent) Key() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetStorageEventKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// OldValue returns the value of property "StorageEvent.oldValue".
//
// The returned bool will be false if there is no such property.
func (this StorageEvent) OldValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetStorageEventOldValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// NewValue returns the value of property "StorageEvent.newValue".
//
// The returned bool will be false if there is no such property.
func (this StorageEvent) NewValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetStorageEventNewValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Url returns the value of property "StorageEvent.url".
//
// The returned bool will be false if there is no such property.
func (this StorageEvent) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetStorageEventUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// StorageArea returns the value of property "StorageEvent.storageArea".
//
// The returned bool will be false if there is no such property.
func (this StorageEvent) StorageArea() (Storage, bool) {
	var _ok bool
	_ret := bindings.GetStorageEventStorageArea(
		this.Ref(), js.Pointer(&_ok),
	)
	return Storage{}.FromRef(_ret), _ok
}

// InitStorageEvent calls the method "StorageEvent.initStorageEvent".
//
// The returned bool will be false if there is no such method.
func (this StorageEvent) InitStorageEvent(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String, url js.String, storageArea Storage) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStorageEventInitStorageEvent(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
		newValue.Ref(),
		url.Ref(),
		storageArea.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitStorageEventFunc returns the method "StorageEvent.initStorageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageEvent) InitStorageEventFunc() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String, url js.String, storageArea Storage)]) {
	return fn.FromRef(
		bindings.StorageEventInitStorageEventFunc(
			this.Ref(),
		),
	)
}

// InitStorageEvent1 calls the method "StorageEvent.initStorageEvent".
//
// The returned bool will be false if there is no such method.
func (this StorageEvent) InitStorageEvent1(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String, url js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStorageEventInitStorageEvent1(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
		newValue.Ref(),
		url.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitStorageEvent1Func returns the method "StorageEvent.initStorageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageEvent) InitStorageEvent1Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String, url js.String)]) {
	return fn.FromRef(
		bindings.StorageEventInitStorageEvent1Func(
			this.Ref(),
		),
	)
}

// InitStorageEvent2 calls the method "StorageEvent.initStorageEvent".
//
// The returned bool will be false if there is no such method.
func (this StorageEvent) InitStorageEvent2(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStorageEventInitStorageEvent2(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
		newValue.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitStorageEvent2Func returns the method "StorageEvent.initStorageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageEvent) InitStorageEvent2Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String)]) {
	return fn.FromRef(
		bindings.StorageEventInitStorageEvent2Func(
			this.Ref(),
		),
	)
}

// InitStorageEvent3 calls the method "StorageEvent.initStorageEvent".
//
// The returned bool will be false if there is no such method.
func (this StorageEvent) InitStorageEvent3(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStorageEventInitStorageEvent3(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitStorageEvent3Func returns the method "StorageEvent.initStorageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageEvent) InitStorageEvent3Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String)]) {
	return fn.FromRef(
		bindings.StorageEventInitStorageEvent3Func(
			this.Ref(),
		),
	)
}

// InitStorageEvent4 calls the method "StorageEvent.initStorageEvent".
//
// The returned bool will be false if there is no such method.
func (this StorageEvent) InitStorageEvent4(typ js.String, bubbles bool, cancelable bool, key js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStorageEventInitStorageEvent4(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitStorageEvent4Func returns the method "StorageEvent.initStorageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageEvent) InitStorageEvent4Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, key js.String)]) {
	return fn.FromRef(
		bindings.StorageEventInitStorageEvent4Func(
			this.Ref(),
		),
	)
}

// InitStorageEvent5 calls the method "StorageEvent.initStorageEvent".
//
// The returned bool will be false if there is no such method.
func (this StorageEvent) InitStorageEvent5(typ js.String, bubbles bool, cancelable bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStorageEventInitStorageEvent5(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitStorageEvent5Func returns the method "StorageEvent.initStorageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageEvent) InitStorageEvent5Func() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool)]) {
	return fn.FromRef(
		bindings.StorageEventInitStorageEvent5Func(
			this.Ref(),
		),
	)
}

// InitStorageEvent6 calls the method "StorageEvent.initStorageEvent".
//
// The returned bool will be false if there is no such method.
func (this StorageEvent) InitStorageEvent6(typ js.String, bubbles bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStorageEventInitStorageEvent6(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitStorageEvent6Func returns the method "StorageEvent.initStorageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageEvent) InitStorageEvent6Func() (fn js.Func[func(typ js.String, bubbles bool)]) {
	return fn.FromRef(
		bindings.StorageEventInitStorageEvent6Func(
			this.Ref(),
		),
	)
}

// InitStorageEvent7 calls the method "StorageEvent.initStorageEvent".
//
// The returned bool will be false if there is no such method.
func (this StorageEvent) InitStorageEvent7(typ js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallStorageEventInitStorageEvent7(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitStorageEvent7Func returns the method "StorageEvent.initStorageEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this StorageEvent) InitStorageEvent7Func() (fn js.Func[func(typ js.String)]) {
	return fn.FromRef(
		bindings.StorageEventInitStorageEvent7Func(
			this.Ref(),
		),
	)
}

type OneOf_Element_ProcessingInstruction struct {
	ref js.Ref
}

func (x OneOf_Element_ProcessingInstruction) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Element_ProcessingInstruction) Free() {
	x.ref.Free()
}

func (x OneOf_Element_ProcessingInstruction) FromRef(ref js.Ref) OneOf_Element_ProcessingInstruction {
	return OneOf_Element_ProcessingInstruction{
		ref: ref,
	}
}

func (x OneOf_Element_ProcessingInstruction) Element() Element {
	return Element{}.FromRef(x.ref)
}

func (x OneOf_Element_ProcessingInstruction) ProcessingInstruction() ProcessingInstruction {
	return ProcessingInstruction{}.FromRef(x.ref)
}

type StyleSheet struct {
	ref js.Ref
}

func (this StyleSheet) Once() StyleSheet {
	this.Ref().Once()
	return this
}

func (this StyleSheet) Ref() js.Ref {
	return this.ref
}

func (this StyleSheet) FromRef(ref js.Ref) StyleSheet {
	this.ref = ref
	return this
}

func (this StyleSheet) Free() {
	this.Ref().Free()
}

// Type returns the value of property "StyleSheet.type".
//
// The returned bool will be false if there is no such property.
func (this StyleSheet) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetStyleSheetType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Href returns the value of property "StyleSheet.href".
//
// The returned bool will be false if there is no such property.
func (this StyleSheet) Href() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetStyleSheetHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// OwnerNode returns the value of property "StyleSheet.ownerNode".
//
// The returned bool will be false if there is no such property.
func (this StyleSheet) OwnerNode() (OneOf_Element_ProcessingInstruction, bool) {
	var _ok bool
	_ret := bindings.GetStyleSheetOwnerNode(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_Element_ProcessingInstruction{}.FromRef(_ret), _ok
}

// ParentStyleSheet returns the value of property "StyleSheet.parentStyleSheet".
//
// The returned bool will be false if there is no such property.
func (this StyleSheet) ParentStyleSheet() (CSSStyleSheet, bool) {
	var _ok bool
	_ret := bindings.GetStyleSheetParentStyleSheet(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleSheet{}.FromRef(_ret), _ok
}

// Title returns the value of property "StyleSheet.title".
//
// The returned bool will be false if there is no such property.
func (this StyleSheet) Title() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetStyleSheetTitle(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Media returns the value of property "StyleSheet.media".
//
// The returned bool will be false if there is no such property.
func (this StyleSheet) Media() (MediaList, bool) {
	var _ok bool
	_ret := bindings.GetStyleSheetMedia(
		this.Ref(), js.Pointer(&_ok),
	)
	return MediaList{}.FromRef(_ret), _ok
}

// Disabled returns the value of property "StyleSheet.disabled".
//
// The returned bool will be false if there is no such property.
func (this StyleSheet) Disabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetStyleSheetDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Disabled sets the value of property "StyleSheet.disabled" to val.
//
// It returns false if the property cannot be set.
func (this StyleSheet) SetDisabled(val bool) bool {
	return js.True == bindings.SetStyleSheetDisabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

type SubmitEventInit struct {
	// Submitter is "SubmitEventInit.submitter"
	//
	// Optional, defaults to null.
	Submitter HTMLElement
	// Bubbles is "SubmitEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "SubmitEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "SubmitEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SubmitEventInit with all fields set.
func (p SubmitEventInit) FromRef(ref js.Ref) SubmitEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 SubmitEventInit SubmitEventInit [// SubmitEventInit] [0x14000a6a3c0 0x14000a6a460 0x14000a6a5a0 0x14000a6a6e0 0x14000a6a500 0x14000a6a640 0x14000a6a780] 0x14000a02d20 {0 0}} in the application heap.
func (p SubmitEventInit) New() js.Ref {
	return bindings.SubmitEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p SubmitEventInit) UpdateFrom(ref js.Ref) {
	bindings.SubmitEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p SubmitEventInit) Update(ref js.Ref) {
	bindings.SubmitEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewSubmitEvent(typ js.String, eventInitDict SubmitEventInit) SubmitEvent {
	return SubmitEvent{}.FromRef(
		bindings.NewSubmitEventBySubmitEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewSubmitEventBySubmitEvent1(typ js.String) SubmitEvent {
	return SubmitEvent{}.FromRef(
		bindings.NewSubmitEventBySubmitEvent1(
			typ.Ref()),
	)
}

type SubmitEvent struct {
	Event
}

func (this SubmitEvent) Once() SubmitEvent {
	this.Ref().Once()
	return this
}

func (this SubmitEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this SubmitEvent) FromRef(ref js.Ref) SubmitEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this SubmitEvent) Free() {
	this.Ref().Free()
}

// Submitter returns the value of property "SubmitEvent.submitter".
//
// The returned bool will be false if there is no such property.
func (this SubmitEvent) Submitter() (HTMLElement, bool) {
	var _ok bool
	_ret := bindings.GetSubmitEventSubmitter(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLElement{}.FromRef(_ret), _ok
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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

// New creates a new ScrollOptions in the application heap.
func (p ScrollOptions) New() js.Ref {
	return bindings.ScrollOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScrollOptions) UpdateFrom(ref js.Ref) {
	bindings.ScrollOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScrollOptions) Update(ref js.Ref) {
	bindings.ScrollOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScrollOptions) FreeMembers(recursive bool) {
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

// New creates a new ScrollTimelineOptions in the application heap.
func (p ScrollTimelineOptions) New() js.Ref {
	return bindings.ScrollTimelineOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScrollTimelineOptions) UpdateFrom(ref js.Ref) {
	bindings.ScrollTimelineOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScrollTimelineOptions) Update(ref js.Ref) {
	bindings.ScrollTimelineOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScrollTimelineOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Source.Ref(),
	)
	p.Source = p.Source.FromRef(js.Undefined)
}

func NewScrollTimeline(options ScrollTimelineOptions) (ret ScrollTimeline) {
	ret.ref = bindings.NewScrollTimelineByScrollTimeline(
		js.Pointer(&options))
	return
}

func NewScrollTimelineByScrollTimeline1() (ret ScrollTimeline) {
	ret.ref = bindings.NewScrollTimelineByScrollTimeline1()
	return
}

type ScrollTimeline struct {
	AnimationTimeline
}

func (this ScrollTimeline) Once() ScrollTimeline {
	this.ref.Once()
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
	this.ref.Free()
}

// Source returns the value of property "ScrollTimeline.source".
//
// It returns ok=false if there is no such property.
func (this ScrollTimeline) Source() (ret Element, ok bool) {
	ok = js.True == bindings.GetScrollTimelineSource(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Axis returns the value of property "ScrollTimeline.axis".
//
// It returns ok=false if there is no such property.
func (this ScrollTimeline) Axis() (ret ScrollAxis, ok bool) {
	ok = js.True == bindings.GetScrollTimelineAxis(
		this.ref, js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: Instrument.FFI_USE MUST be set to true to get Instrument used.
	Instrument PaymentCredentialInstrument
	// Timeout is "SecurePaymentConfirmationRequest.timeout"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timeout MUST be set to true to make this field effective.
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
	//
	// NOTE: Extensions.FFI_USE MUST be set to true to get Extensions used.
	Extensions AuthenticationExtensionsClientInputs
	// Locale is "SecurePaymentConfirmationRequest.locale"
	//
	// Optional
	Locale js.Array[js.String]
	// ShowOptOut is "SecurePaymentConfirmationRequest.showOptOut"
	//
	// Optional
	//
	// NOTE: FFI_USE_ShowOptOut MUST be set to true to make this field effective.
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

// New creates a new SecurePaymentConfirmationRequest in the application heap.
func (p SecurePaymentConfirmationRequest) New() js.Ref {
	return bindings.SecurePaymentConfirmationRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SecurePaymentConfirmationRequest) UpdateFrom(ref js.Ref) {
	bindings.SecurePaymentConfirmationRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SecurePaymentConfirmationRequest) Update(ref js.Ref) {
	bindings.SecurePaymentConfirmationRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SecurePaymentConfirmationRequest) FreeMembers(recursive bool) {
	js.Free(
		p.Challenge.Ref(),
		p.RpId.Ref(),
		p.CredentialIds.Ref(),
		p.PayeeName.Ref(),
		p.PayeeOrigin.Ref(),
		p.Locale.Ref(),
	)
	p.Challenge = p.Challenge.FromRef(js.Undefined)
	p.RpId = p.RpId.FromRef(js.Undefined)
	p.CredentialIds = p.CredentialIds.FromRef(js.Undefined)
	p.PayeeName = p.PayeeName.FromRef(js.Undefined)
	p.PayeeOrigin = p.PayeeOrigin.FromRef(js.Undefined)
	p.Locale = p.Locale.FromRef(js.Undefined)
	if recursive {
		p.Instrument.FreeMembers(true)
		p.Extensions.FreeMembers(true)
	}
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
	//
	// NOTE: FFI_USE_LineNumber MUST be set to true to make this field effective.
	LineNumber uint32
	// ColumnNumber is "SecurityPolicyViolationEventInit.columnNumber"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ColumnNumber MUST be set to true to make this field effective.
	ColumnNumber uint32
	// Bubbles is "SecurityPolicyViolationEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "SecurityPolicyViolationEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "SecurityPolicyViolationEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new SecurityPolicyViolationEventInit in the application heap.
func (p SecurityPolicyViolationEventInit) New() js.Ref {
	return bindings.SecurityPolicyViolationEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SecurityPolicyViolationEventInit) UpdateFrom(ref js.Ref) {
	bindings.SecurityPolicyViolationEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SecurityPolicyViolationEventInit) Update(ref js.Ref) {
	bindings.SecurityPolicyViolationEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SecurityPolicyViolationEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.DocumentURI.Ref(),
		p.Referrer.Ref(),
		p.BlockedURI.Ref(),
		p.ViolatedDirective.Ref(),
		p.EffectiveDirective.Ref(),
		p.OriginalPolicy.Ref(),
		p.SourceFile.Ref(),
		p.Sample.Ref(),
	)
	p.DocumentURI = p.DocumentURI.FromRef(js.Undefined)
	p.Referrer = p.Referrer.FromRef(js.Undefined)
	p.BlockedURI = p.BlockedURI.FromRef(js.Undefined)
	p.ViolatedDirective = p.ViolatedDirective.FromRef(js.Undefined)
	p.EffectiveDirective = p.EffectiveDirective.FromRef(js.Undefined)
	p.OriginalPolicy = p.OriginalPolicy.FromRef(js.Undefined)
	p.SourceFile = p.SourceFile.FromRef(js.Undefined)
	p.Sample = p.Sample.FromRef(js.Undefined)
}

func NewSecurityPolicyViolationEvent(typ js.String, eventInitDict SecurityPolicyViolationEventInit) (ret SecurityPolicyViolationEvent) {
	ret.ref = bindings.NewSecurityPolicyViolationEventBySecurityPolicyViolationEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewSecurityPolicyViolationEventBySecurityPolicyViolationEvent1(typ js.String) (ret SecurityPolicyViolationEvent) {
	ret.ref = bindings.NewSecurityPolicyViolationEventBySecurityPolicyViolationEvent1(
		typ.Ref())
	return
}

type SecurityPolicyViolationEvent struct {
	Event
}

func (this SecurityPolicyViolationEvent) Once() SecurityPolicyViolationEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// DocumentURI returns the value of property "SecurityPolicyViolationEvent.documentURI".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) DocumentURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventDocumentURI(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Referrer returns the value of property "SecurityPolicyViolationEvent.referrer".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) Referrer() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventReferrer(
		this.ref, js.Pointer(&ret),
	)
	return
}

// BlockedURI returns the value of property "SecurityPolicyViolationEvent.blockedURI".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) BlockedURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventBlockedURI(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EffectiveDirective returns the value of property "SecurityPolicyViolationEvent.effectiveDirective".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) EffectiveDirective() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventEffectiveDirective(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ViolatedDirective returns the value of property "SecurityPolicyViolationEvent.violatedDirective".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) ViolatedDirective() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventViolatedDirective(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OriginalPolicy returns the value of property "SecurityPolicyViolationEvent.originalPolicy".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) OriginalPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventOriginalPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SourceFile returns the value of property "SecurityPolicyViolationEvent.sourceFile".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) SourceFile() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventSourceFile(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Sample returns the value of property "SecurityPolicyViolationEvent.sample".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) Sample() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventSample(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Disposition returns the value of property "SecurityPolicyViolationEvent.disposition".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) Disposition() (ret SecurityPolicyViolationEventDisposition, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventDisposition(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StatusCode returns the value of property "SecurityPolicyViolationEvent.statusCode".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) StatusCode() (ret uint16, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventStatusCode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LineNumber returns the value of property "SecurityPolicyViolationEvent.lineNumber".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) LineNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventLineNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ColumnNumber returns the value of property "SecurityPolicyViolationEvent.columnNumber".
//
// It returns ok=false if there is no such property.
func (this SecurityPolicyViolationEvent) ColumnNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSecurityPolicyViolationEventColumnNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

type Sensor struct {
	EventTarget
}

func (this Sensor) Once() Sensor {
	this.ref.Once()
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
	this.ref.Free()
}

// Activated returns the value of property "Sensor.activated".
//
// It returns ok=false if there is no such property.
func (this Sensor) Activated() (ret bool, ok bool) {
	ok = js.True == bindings.GetSensorActivated(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasReading returns the value of property "Sensor.hasReading".
//
// It returns ok=false if there is no such property.
func (this Sensor) HasReading() (ret bool, ok bool) {
	ok = js.True == bindings.GetSensorHasReading(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Timestamp returns the value of property "Sensor.timestamp".
//
// It returns ok=false if there is no such property.
func (this Sensor) Timestamp() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetSensorTimestamp(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncStart returns true if the method "Sensor.start" exists.
func (this Sensor) HasFuncStart() bool {
	return js.True == bindings.HasFuncSensorStart(
		this.ref,
	)
}

// FuncStart returns the method "Sensor.start".
func (this Sensor) FuncStart() (fn js.Func[func()]) {
	bindings.FuncSensorStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start calls the method "Sensor.start".
func (this Sensor) Start() (ret js.Void) {
	bindings.CallSensorStart(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStart calls the method "Sensor.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Sensor) TryStart() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySensorStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStop returns true if the method "Sensor.stop" exists.
func (this Sensor) HasFuncStop() bool {
	return js.True == bindings.HasFuncSensorStop(
		this.ref,
	)
}

// FuncStop returns the method "Sensor.stop".
func (this Sensor) FuncStop() (fn js.Func[func()]) {
	bindings.FuncSensorStop(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stop calls the method "Sensor.stop".
func (this Sensor) Stop() (ret js.Void) {
	bindings.CallSensorStop(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "Sensor.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Sensor) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySensorStop(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SensorErrorEventInit struct {
	// Error is "SensorErrorEventInit.error"
	//
	// Required
	Error DOMException
	// Bubbles is "SensorErrorEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "SensorErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "SensorErrorEventInit.composed"
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

// FromRef calls UpdateFrom and returns a SensorErrorEventInit with all fields set.
func (p SensorErrorEventInit) FromRef(ref js.Ref) SensorErrorEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SensorErrorEventInit in the application heap.
func (p SensorErrorEventInit) New() js.Ref {
	return bindings.SensorErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SensorErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.SensorErrorEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SensorErrorEventInit) Update(ref js.Ref) {
	bindings.SensorErrorEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SensorErrorEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Error.Ref(),
	)
	p.Error = p.Error.FromRef(js.Undefined)
}

func NewSensorErrorEvent(typ js.String, errorEventInitDict SensorErrorEventInit) (ret SensorErrorEvent) {
	ret.ref = bindings.NewSensorErrorEventBySensorErrorEvent(
		typ.Ref(),
		js.Pointer(&errorEventInitDict))
	return
}

type SensorErrorEvent struct {
	Event
}

func (this SensorErrorEvent) Once() SensorErrorEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Error returns the value of property "SensorErrorEvent.error".
//
// It returns ok=false if there is no such property.
func (this SensorErrorEvent) Error() (ret DOMException, ok bool) {
	ok = js.True == bindings.GetSensorErrorEventError(
		this.ref, js.Pointer(&ret),
	)
	return
}

func NewSequenceEffect(children js.Array[AnimationEffect], timing OneOf_Float64_EffectTiming) (ret SequenceEffect) {
	ret.ref = bindings.NewSequenceEffectBySequenceEffect(
		children.Ref(),
		timing.Ref())
	return
}

func NewSequenceEffectBySequenceEffect1(children js.Array[AnimationEffect]) (ret SequenceEffect) {
	ret.ref = bindings.NewSequenceEffectBySequenceEffect1(
		children.Ref())
	return
}

type SequenceEffect struct {
	GroupEffect
}

func (this SequenceEffect) Once() SequenceEffect {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncClone returns true if the method "SequenceEffect.clone" exists.
func (this SequenceEffect) HasFuncClone() bool {
	return js.True == bindings.HasFuncSequenceEffectClone(
		this.ref,
	)
}

// FuncClone returns the method "SequenceEffect.clone".
func (this SequenceEffect) FuncClone() (fn js.Func[func() SequenceEffect]) {
	bindings.FuncSequenceEffectClone(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clone calls the method "SequenceEffect.clone".
func (this SequenceEffect) Clone() (ret SequenceEffect) {
	bindings.CallSequenceEffectClone(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClone calls the method "SequenceEffect.clone"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SequenceEffect) TryClone() (ret SequenceEffect, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySequenceEffectClone(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ServiceWorkerGlobalScope struct {
	WorkerGlobalScope
}

func (this ServiceWorkerGlobalScope) Once() ServiceWorkerGlobalScope {
	this.ref.Once()
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
	this.ref.Free()
}

// Clients returns the value of property "ServiceWorkerGlobalScope.clients".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerGlobalScope) Clients() (ret Clients, ok bool) {
	ok = js.True == bindings.GetServiceWorkerGlobalScopeClients(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Registration returns the value of property "ServiceWorkerGlobalScope.registration".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerGlobalScope) Registration() (ret ServiceWorkerRegistration, ok bool) {
	ok = js.True == bindings.GetServiceWorkerGlobalScopeRegistration(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ServiceWorker returns the value of property "ServiceWorkerGlobalScope.serviceWorker".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerGlobalScope) ServiceWorker() (ret ServiceWorker, ok bool) {
	ok = js.True == bindings.GetServiceWorkerGlobalScopeServiceWorker(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CookieStore returns the value of property "ServiceWorkerGlobalScope.cookieStore".
//
// It returns ok=false if there is no such property.
func (this ServiceWorkerGlobalScope) CookieStore() (ret CookieStore, ok bool) {
	ok = js.True == bindings.GetServiceWorkerGlobalScopeCookieStore(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSkipWaiting returns true if the method "ServiceWorkerGlobalScope.skipWaiting" exists.
func (this ServiceWorkerGlobalScope) HasFuncSkipWaiting() bool {
	return js.True == bindings.HasFuncServiceWorkerGlobalScopeSkipWaiting(
		this.ref,
	)
}

// FuncSkipWaiting returns the method "ServiceWorkerGlobalScope.skipWaiting".
func (this ServiceWorkerGlobalScope) FuncSkipWaiting() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncServiceWorkerGlobalScopeSkipWaiting(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SkipWaiting calls the method "ServiceWorkerGlobalScope.skipWaiting".
func (this ServiceWorkerGlobalScope) SkipWaiting() (ret js.Promise[js.Void]) {
	bindings.CallServiceWorkerGlobalScopeSkipWaiting(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySkipWaiting calls the method "ServiceWorkerGlobalScope.skipWaiting"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ServiceWorkerGlobalScope) TrySkipWaiting() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryServiceWorkerGlobalScopeSkipWaiting(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewShadowAnimation(source Animation, newTarget OneOf_Element_CSSPseudoElement) (ret ShadowAnimation) {
	ret.ref = bindings.NewShadowAnimationByShadowAnimation(
		source.Ref(),
		newTarget.Ref())
	return
}

type ShadowAnimation struct {
	Animation
}

func (this ShadowAnimation) Once() ShadowAnimation {
	this.ref.Once()
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
	this.ref.Free()
}

// SourceAnimation returns the value of property "ShadowAnimation.sourceAnimation".
//
// It returns ok=false if there is no such property.
func (this ShadowAnimation) SourceAnimation() (ret Animation, ok bool) {
	ok = js.True == bindings.GetShadowAnimationSourceAnimation(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SharedStorageSetMethodOptions struct {
	// IgnoreIfPresent is "SharedStorageSetMethodOptions.ignoreIfPresent"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IgnoreIfPresent MUST be set to true to make this field effective.
	IgnoreIfPresent bool

	FFI_USE_IgnoreIfPresent bool // for IgnoreIfPresent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SharedStorageSetMethodOptions with all fields set.
func (p SharedStorageSetMethodOptions) FromRef(ref js.Ref) SharedStorageSetMethodOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SharedStorageSetMethodOptions in the application heap.
func (p SharedStorageSetMethodOptions) New() js.Ref {
	return bindings.SharedStorageSetMethodOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SharedStorageSetMethodOptions) UpdateFrom(ref js.Ref) {
	bindings.SharedStorageSetMethodOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SharedStorageSetMethodOptions) Update(ref js.Ref) {
	bindings.SharedStorageSetMethodOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SharedStorageSetMethodOptions) FreeMembers(recursive bool) {
}

type SharedStorage struct {
	ref js.Ref
}

func (this SharedStorage) Once() SharedStorage {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncSet returns true if the method "SharedStorage.set" exists.
func (this SharedStorage) HasFuncSet() bool {
	return js.True == bindings.HasFuncSharedStorageSet(
		this.ref,
	)
}

// FuncSet returns the method "SharedStorage.set".
func (this SharedStorage) FuncSet() (fn js.Func[func(key js.String, value js.String, options SharedStorageSetMethodOptions) js.Promise[js.Any]]) {
	bindings.FuncSharedStorageSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "SharedStorage.set".
func (this SharedStorage) Set(key js.String, value js.String, options SharedStorageSetMethodOptions) (ret js.Promise[js.Any]) {
	bindings.CallSharedStorageSet(
		this.ref, js.Pointer(&ret),
		key.Ref(),
		value.Ref(),
		js.Pointer(&options),
	)

	return
}

// TrySet calls the method "SharedStorage.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SharedStorage) TrySet(key js.String, value js.String, options SharedStorageSetMethodOptions) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharedStorageSet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
		value.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncSet1 returns true if the method "SharedStorage.set" exists.
func (this SharedStorage) HasFuncSet1() bool {
	return js.True == bindings.HasFuncSharedStorageSet1(
		this.ref,
	)
}

// FuncSet1 returns the method "SharedStorage.set".
func (this SharedStorage) FuncSet1() (fn js.Func[func(key js.String, value js.String) js.Promise[js.Any]]) {
	bindings.FuncSharedStorageSet1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set1 calls the method "SharedStorage.set".
func (this SharedStorage) Set1(key js.String, value js.String) (ret js.Promise[js.Any]) {
	bindings.CallSharedStorageSet1(
		this.ref, js.Pointer(&ret),
		key.Ref(),
		value.Ref(),
	)

	return
}

// TrySet1 calls the method "SharedStorage.set"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SharedStorage) TrySet1(key js.String, value js.String) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharedStorageSet1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncAppend returns true if the method "SharedStorage.append" exists.
func (this SharedStorage) HasFuncAppend() bool {
	return js.True == bindings.HasFuncSharedStorageAppend(
		this.ref,
	)
}

// FuncAppend returns the method "SharedStorage.append".
func (this SharedStorage) FuncAppend() (fn js.Func[func(key js.String, value js.String) js.Promise[js.Any]]) {
	bindings.FuncSharedStorageAppend(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Append calls the method "SharedStorage.append".
func (this SharedStorage) Append(key js.String, value js.String) (ret js.Promise[js.Any]) {
	bindings.CallSharedStorageAppend(
		this.ref, js.Pointer(&ret),
		key.Ref(),
		value.Ref(),
	)

	return
}

// TryAppend calls the method "SharedStorage.append"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SharedStorage) TryAppend(key js.String, value js.String) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharedStorageAppend(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
		value.Ref(),
	)

	return
}

// HasFuncDelete returns true if the method "SharedStorage.delete" exists.
func (this SharedStorage) HasFuncDelete() bool {
	return js.True == bindings.HasFuncSharedStorageDelete(
		this.ref,
	)
}

// FuncDelete returns the method "SharedStorage.delete".
func (this SharedStorage) FuncDelete() (fn js.Func[func(key js.String) js.Promise[js.Any]]) {
	bindings.FuncSharedStorageDelete(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Delete calls the method "SharedStorage.delete".
func (this SharedStorage) Delete(key js.String) (ret js.Promise[js.Any]) {
	bindings.CallSharedStorageDelete(
		this.ref, js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryDelete calls the method "SharedStorage.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SharedStorage) TryDelete(key js.String) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharedStorageDelete(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

// HasFuncClear returns true if the method "SharedStorage.clear" exists.
func (this SharedStorage) HasFuncClear() bool {
	return js.True == bindings.HasFuncSharedStorageClear(
		this.ref,
	)
}

// FuncClear returns the method "SharedStorage.clear".
func (this SharedStorage) FuncClear() (fn js.Func[func() js.Promise[js.Any]]) {
	bindings.FuncSharedStorageClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "SharedStorage.clear".
func (this SharedStorage) Clear() (ret js.Promise[js.Any]) {
	bindings.CallSharedStorageClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "SharedStorage.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SharedStorage) TryClear() (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharedStorageClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SharedStorageOperation struct {
	ref js.Ref
}

func (this SharedStorageOperation) Once() SharedStorageOperation {
	this.ref.Once()
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
	this.ref.Free()
}

type SharedStorageOperationConstructorFunc func(this js.Ref, options *SharedStorageRunOperationMethodOptions) js.Ref

func (fn SharedStorageOperationConstructorFunc) Register() js.Func[func(options *SharedStorageRunOperationMethodOptions) SharedStorageOperation] {
	return js.RegisterCallback[func(options *SharedStorageRunOperationMethodOptions) SharedStorageOperation](
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
	var arg0 SharedStorageRunOperationMethodOptions
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

type SharedStorageOperationConstructor[T any] struct {
	Fn  func(arg T, this js.Ref, options *SharedStorageRunOperationMethodOptions) js.Ref
	Arg T
}

func (cb *SharedStorageOperationConstructor[T]) Register() js.Func[func(options *SharedStorageRunOperationMethodOptions) SharedStorageOperation] {
	return js.RegisterCallback[func(options *SharedStorageRunOperationMethodOptions) SharedStorageOperation](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SharedStorageOperationConstructor[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 SharedStorageRunOperationMethodOptions
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

type SharedStorageRunOperation struct {
	SharedStorageOperation
}

func (this SharedStorageRunOperation) Once() SharedStorageRunOperation {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncRun returns true if the method "SharedStorageRunOperation.run" exists.
func (this SharedStorageRunOperation) HasFuncRun() bool {
	return js.True == bindings.HasFuncSharedStorageRunOperationRun(
		this.ref,
	)
}

// FuncRun returns the method "SharedStorageRunOperation.run".
func (this SharedStorageRunOperation) FuncRun() (fn js.Func[func(data js.Object) js.Promise[js.Void]]) {
	bindings.FuncSharedStorageRunOperationRun(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Run calls the method "SharedStorageRunOperation.run".
func (this SharedStorageRunOperation) Run(data js.Object) (ret js.Promise[js.Void]) {
	bindings.CallSharedStorageRunOperationRun(
		this.ref, js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TryRun calls the method "SharedStorageRunOperation.run"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SharedStorageRunOperation) TryRun(data js.Object) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharedStorageRunOperationRun(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

type SharedStorageSelectURLOperation struct {
	SharedStorageOperation
}

func (this SharedStorageSelectURLOperation) Once() SharedStorageSelectURLOperation {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncRun returns true if the method "SharedStorageSelectURLOperation.run" exists.
func (this SharedStorageSelectURLOperation) HasFuncRun() bool {
	return js.True == bindings.HasFuncSharedStorageSelectURLOperationRun(
		this.ref,
	)
}

// FuncRun returns the method "SharedStorageSelectURLOperation.run".
func (this SharedStorageSelectURLOperation) FuncRun() (fn js.Func[func(data js.Object, urls js.FrozenArray[SharedStorageUrlWithMetadata]) js.Promise[js.Number[int32]]]) {
	bindings.FuncSharedStorageSelectURLOperationRun(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Run calls the method "SharedStorageSelectURLOperation.run".
func (this SharedStorageSelectURLOperation) Run(data js.Object, urls js.FrozenArray[SharedStorageUrlWithMetadata]) (ret js.Promise[js.Number[int32]]) {
	bindings.CallSharedStorageSelectURLOperationRun(
		this.ref, js.Pointer(&ret),
		data.Ref(),
		urls.Ref(),
	)

	return
}

// TryRun calls the method "SharedStorageSelectURLOperation.run"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SharedStorageSelectURLOperation) TryRun(data js.Object, urls js.FrozenArray[SharedStorageUrlWithMetadata]) (ret js.Promise[js.Number[int32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharedStorageSelectURLOperationRun(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
		urls.Ref(),
	)

	return
}

type WorkletSharedStorage struct {
	SharedStorage
}

func (this WorkletSharedStorage) Once() WorkletSharedStorage {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncGet returns true if the method "WorkletSharedStorage.get" exists.
func (this WorkletSharedStorage) HasFuncGet() bool {
	return js.True == bindings.HasFuncWorkletSharedStorageGet(
		this.ref,
	)
}

// FuncGet returns the method "WorkletSharedStorage.get".
func (this WorkletSharedStorage) FuncGet() (fn js.Func[func(key js.String) js.Promise[js.String]]) {
	bindings.FuncWorkletSharedStorageGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "WorkletSharedStorage.get".
func (this WorkletSharedStorage) Get(key js.String) (ret js.Promise[js.String]) {
	bindings.CallWorkletSharedStorageGet(
		this.ref, js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryGet calls the method "WorkletSharedStorage.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkletSharedStorage) TryGet(key js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletSharedStorageGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

// HasFuncLength returns true if the method "WorkletSharedStorage.length" exists.
func (this WorkletSharedStorage) HasFuncLength() bool {
	return js.True == bindings.HasFuncWorkletSharedStorageLength(
		this.ref,
	)
}

// FuncLength returns the method "WorkletSharedStorage.length".
func (this WorkletSharedStorage) FuncLength() (fn js.Func[func() js.Promise[js.Number[uint32]]]) {
	bindings.FuncWorkletSharedStorageLength(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Length calls the method "WorkletSharedStorage.length".
func (this WorkletSharedStorage) Length() (ret js.Promise[js.Number[uint32]]) {
	bindings.CallWorkletSharedStorageLength(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryLength calls the method "WorkletSharedStorage.length"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkletSharedStorage) TryLength() (ret js.Promise[js.Number[uint32]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletSharedStorageLength(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemainingBudget returns true if the method "WorkletSharedStorage.remainingBudget" exists.
func (this WorkletSharedStorage) HasFuncRemainingBudget() bool {
	return js.True == bindings.HasFuncWorkletSharedStorageRemainingBudget(
		this.ref,
	)
}

// FuncRemainingBudget returns the method "WorkletSharedStorage.remainingBudget".
func (this WorkletSharedStorage) FuncRemainingBudget() (fn js.Func[func() js.Promise[js.Number[float64]]]) {
	bindings.FuncWorkletSharedStorageRemainingBudget(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemainingBudget calls the method "WorkletSharedStorage.remainingBudget".
func (this WorkletSharedStorage) RemainingBudget() (ret js.Promise[js.Number[float64]]) {
	bindings.CallWorkletSharedStorageRemainingBudget(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRemainingBudget calls the method "WorkletSharedStorage.remainingBudget"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this WorkletSharedStorage) TryRemainingBudget() (ret js.Promise[js.Number[float64]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryWorkletSharedStorageRemainingBudget(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SharedStorageWorkletGlobalScope struct {
	WorkletGlobalScope
}

func (this SharedStorageWorkletGlobalScope) Once() SharedStorageWorkletGlobalScope {
	this.ref.Once()
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
	this.ref.Free()
}

// SharedStorage returns the value of property "SharedStorageWorkletGlobalScope.sharedStorage".
//
// It returns ok=false if there is no such property.
func (this SharedStorageWorkletGlobalScope) SharedStorage() (ret WorkletSharedStorage, ok bool) {
	ok = js.True == bindings.GetSharedStorageWorkletGlobalScopeSharedStorage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRegister returns true if the method "SharedStorageWorkletGlobalScope.register" exists.
func (this SharedStorageWorkletGlobalScope) HasFuncRegister() bool {
	return js.True == bindings.HasFuncSharedStorageWorkletGlobalScopeRegister(
		this.ref,
	)
}

// FuncRegister returns the method "SharedStorageWorkletGlobalScope.register".
func (this SharedStorageWorkletGlobalScope) FuncRegister() (fn js.Func[func(name js.String, operationCtor js.Func[func(options *SharedStorageRunOperationMethodOptions) SharedStorageOperation])]) {
	bindings.FuncSharedStorageWorkletGlobalScopeRegister(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Register calls the method "SharedStorageWorkletGlobalScope.register".
func (this SharedStorageWorkletGlobalScope) Register(name js.String, operationCtor js.Func[func(options *SharedStorageRunOperationMethodOptions) SharedStorageOperation]) (ret js.Void) {
	bindings.CallSharedStorageWorkletGlobalScopeRegister(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		operationCtor.Ref(),
	)

	return
}

// TryRegister calls the method "SharedStorageWorkletGlobalScope.register"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SharedStorageWorkletGlobalScope) TryRegister(name js.String, operationCtor js.Func[func(options *SharedStorageRunOperationMethodOptions) SharedStorageOperation]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharedStorageWorkletGlobalScopeRegister(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		operationCtor.Ref(),
	)

	return
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

func NewSharedWorker(scriptURL js.String, options OneOf_String_WorkerOptions) (ret SharedWorker) {
	ret.ref = bindings.NewSharedWorkerBySharedWorker(
		scriptURL.Ref(),
		options.Ref())
	return
}

func NewSharedWorkerBySharedWorker1(scriptURL js.String) (ret SharedWorker) {
	ret.ref = bindings.NewSharedWorkerBySharedWorker1(
		scriptURL.Ref())
	return
}

type SharedWorker struct {
	EventTarget
}

func (this SharedWorker) Once() SharedWorker {
	this.ref.Once()
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
	this.ref.Free()
}

// Port returns the value of property "SharedWorker.port".
//
// It returns ok=false if there is no such property.
func (this SharedWorker) Port() (ret MessagePort, ok bool) {
	ok = js.True == bindings.GetSharedWorkerPort(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SharedWorkerGlobalScope struct {
	WorkerGlobalScope
}

func (this SharedWorkerGlobalScope) Once() SharedWorkerGlobalScope {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "SharedWorkerGlobalScope.name".
//
// It returns ok=false if there is no such property.
func (this SharedWorkerGlobalScope) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSharedWorkerGlobalScopeName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncClose returns true if the method "SharedWorkerGlobalScope.close" exists.
func (this SharedWorkerGlobalScope) HasFuncClose() bool {
	return js.True == bindings.HasFuncSharedWorkerGlobalScopeClose(
		this.ref,
	)
}

// FuncClose returns the method "SharedWorkerGlobalScope.close".
func (this SharedWorkerGlobalScope) FuncClose() (fn js.Func[func()]) {
	bindings.FuncSharedWorkerGlobalScopeClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "SharedWorkerGlobalScope.close".
func (this SharedWorkerGlobalScope) Close() (ret js.Void) {
	bindings.CallSharedWorkerGlobalScopeClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "SharedWorkerGlobalScope.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SharedWorkerGlobalScope) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySharedWorkerGlobalScopeClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SmallCryptoKeyID uint64

type SpeechGrammar struct {
	ref js.Ref
}

func (this SpeechGrammar) Once() SpeechGrammar {
	this.ref.Once()
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
	this.ref.Free()
}

// Src returns the value of property "SpeechGrammar.src".
//
// It returns ok=false if there is no such property.
func (this SpeechGrammar) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSpeechGrammarSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "SpeechGrammar.src" to val.
//
// It returns false if the property cannot be set.
func (this SpeechGrammar) SetSrc(val js.String) bool {
	return js.True == bindings.SetSpeechGrammarSrc(
		this.ref,
		val.Ref(),
	)
}

// Weight returns the value of property "SpeechGrammar.weight".
//
// It returns ok=false if there is no such property.
func (this SpeechGrammar) Weight() (ret float32, ok bool) {
	ok = js.True == bindings.GetSpeechGrammarWeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWeight sets the value of property "SpeechGrammar.weight" to val.
//
// It returns false if the property cannot be set.
func (this SpeechGrammar) SetWeight(val float32) bool {
	return js.True == bindings.SetSpeechGrammarWeight(
		this.ref,
		float32(val),
	)
}

type SpeechGrammarList struct {
	ref js.Ref
}

func (this SpeechGrammarList) Once() SpeechGrammarList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "SpeechGrammarList.length".
//
// It returns ok=false if there is no such property.
func (this SpeechGrammarList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSpeechGrammarListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "SpeechGrammarList.item" exists.
func (this SpeechGrammarList) HasFuncItem() bool {
	return js.True == bindings.HasFuncSpeechGrammarListItem(
		this.ref,
	)
}

// FuncItem returns the method "SpeechGrammarList.item".
func (this SpeechGrammarList) FuncItem() (fn js.Func[func(index uint32) SpeechGrammar]) {
	bindings.FuncSpeechGrammarListItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "SpeechGrammarList.item".
func (this SpeechGrammarList) Item(index uint32) (ret SpeechGrammar) {
	bindings.CallSpeechGrammarListItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "SpeechGrammarList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SpeechGrammarList) TryItem(index uint32) (ret SpeechGrammar, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechGrammarListItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncAddFromURI returns true if the method "SpeechGrammarList.addFromURI" exists.
func (this SpeechGrammarList) HasFuncAddFromURI() bool {
	return js.True == bindings.HasFuncSpeechGrammarListAddFromURI(
		this.ref,
	)
}

// FuncAddFromURI returns the method "SpeechGrammarList.addFromURI".
func (this SpeechGrammarList) FuncAddFromURI() (fn js.Func[func(src js.String, weight float32)]) {
	bindings.FuncSpeechGrammarListAddFromURI(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddFromURI calls the method "SpeechGrammarList.addFromURI".
func (this SpeechGrammarList) AddFromURI(src js.String, weight float32) (ret js.Void) {
	bindings.CallSpeechGrammarListAddFromURI(
		this.ref, js.Pointer(&ret),
		src.Ref(),
		float32(weight),
	)

	return
}

// TryAddFromURI calls the method "SpeechGrammarList.addFromURI"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SpeechGrammarList) TryAddFromURI(src js.String, weight float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechGrammarListAddFromURI(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		src.Ref(),
		float32(weight),
	)

	return
}

// HasFuncAddFromURI1 returns true if the method "SpeechGrammarList.addFromURI" exists.
func (this SpeechGrammarList) HasFuncAddFromURI1() bool {
	return js.True == bindings.HasFuncSpeechGrammarListAddFromURI1(
		this.ref,
	)
}

// FuncAddFromURI1 returns the method "SpeechGrammarList.addFromURI".
func (this SpeechGrammarList) FuncAddFromURI1() (fn js.Func[func(src js.String)]) {
	bindings.FuncSpeechGrammarListAddFromURI1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddFromURI1 calls the method "SpeechGrammarList.addFromURI".
func (this SpeechGrammarList) AddFromURI1(src js.String) (ret js.Void) {
	bindings.CallSpeechGrammarListAddFromURI1(
		this.ref, js.Pointer(&ret),
		src.Ref(),
	)

	return
}

// TryAddFromURI1 calls the method "SpeechGrammarList.addFromURI"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SpeechGrammarList) TryAddFromURI1(src js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechGrammarListAddFromURI1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		src.Ref(),
	)

	return
}

// HasFuncAddFromString returns true if the method "SpeechGrammarList.addFromString" exists.
func (this SpeechGrammarList) HasFuncAddFromString() bool {
	return js.True == bindings.HasFuncSpeechGrammarListAddFromString(
		this.ref,
	)
}

// FuncAddFromString returns the method "SpeechGrammarList.addFromString".
func (this SpeechGrammarList) FuncAddFromString() (fn js.Func[func(string js.String, weight float32)]) {
	bindings.FuncSpeechGrammarListAddFromString(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddFromString calls the method "SpeechGrammarList.addFromString".
func (this SpeechGrammarList) AddFromString(string js.String, weight float32) (ret js.Void) {
	bindings.CallSpeechGrammarListAddFromString(
		this.ref, js.Pointer(&ret),
		string.Ref(),
		float32(weight),
	)

	return
}

// TryAddFromString calls the method "SpeechGrammarList.addFromString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SpeechGrammarList) TryAddFromString(string js.String, weight float32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechGrammarListAddFromString(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		string.Ref(),
		float32(weight),
	)

	return
}

// HasFuncAddFromString1 returns true if the method "SpeechGrammarList.addFromString" exists.
func (this SpeechGrammarList) HasFuncAddFromString1() bool {
	return js.True == bindings.HasFuncSpeechGrammarListAddFromString1(
		this.ref,
	)
}

// FuncAddFromString1 returns the method "SpeechGrammarList.addFromString".
func (this SpeechGrammarList) FuncAddFromString1() (fn js.Func[func(string js.String)]) {
	bindings.FuncSpeechGrammarListAddFromString1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddFromString1 calls the method "SpeechGrammarList.addFromString".
func (this SpeechGrammarList) AddFromString1(string js.String) (ret js.Void) {
	bindings.CallSpeechGrammarListAddFromString1(
		this.ref, js.Pointer(&ret),
		string.Ref(),
	)

	return
}

// TryAddFromString1 calls the method "SpeechGrammarList.addFromString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SpeechGrammarList) TryAddFromString1(string js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechGrammarListAddFromString1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		string.Ref(),
	)

	return
}

type SpeechRecognition struct {
	EventTarget
}

func (this SpeechRecognition) Once() SpeechRecognition {
	this.ref.Once()
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
	this.ref.Free()
}

// Grammars returns the value of property "SpeechRecognition.grammars".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognition) Grammars() (ret SpeechGrammarList, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionGrammars(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetGrammars sets the value of property "SpeechRecognition.grammars" to val.
//
// It returns false if the property cannot be set.
func (this SpeechRecognition) SetGrammars(val SpeechGrammarList) bool {
	return js.True == bindings.SetSpeechRecognitionGrammars(
		this.ref,
		val.Ref(),
	)
}

// Lang returns the value of property "SpeechRecognition.lang".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognition) Lang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionLang(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLang sets the value of property "SpeechRecognition.lang" to val.
//
// It returns false if the property cannot be set.
func (this SpeechRecognition) SetLang(val js.String) bool {
	return js.True == bindings.SetSpeechRecognitionLang(
		this.ref,
		val.Ref(),
	)
}

// Continuous returns the value of property "SpeechRecognition.continuous".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognition) Continuous() (ret bool, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionContinuous(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetContinuous sets the value of property "SpeechRecognition.continuous" to val.
//
// It returns false if the property cannot be set.
func (this SpeechRecognition) SetContinuous(val bool) bool {
	return js.True == bindings.SetSpeechRecognitionContinuous(
		this.ref,
		js.Bool(bool(val)),
	)
}

// InterimResults returns the value of property "SpeechRecognition.interimResults".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognition) InterimResults() (ret bool, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionInterimResults(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetInterimResults sets the value of property "SpeechRecognition.interimResults" to val.
//
// It returns false if the property cannot be set.
func (this SpeechRecognition) SetInterimResults(val bool) bool {
	return js.True == bindings.SetSpeechRecognitionInterimResults(
		this.ref,
		js.Bool(bool(val)),
	)
}

// MaxAlternatives returns the value of property "SpeechRecognition.maxAlternatives".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognition) MaxAlternatives() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionMaxAlternatives(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMaxAlternatives sets the value of property "SpeechRecognition.maxAlternatives" to val.
//
// It returns false if the property cannot be set.
func (this SpeechRecognition) SetMaxAlternatives(val uint32) bool {
	return js.True == bindings.SetSpeechRecognitionMaxAlternatives(
		this.ref,
		uint32(val),
	)
}

// HasFuncStart returns true if the method "SpeechRecognition.start" exists.
func (this SpeechRecognition) HasFuncStart() bool {
	return js.True == bindings.HasFuncSpeechRecognitionStart(
		this.ref,
	)
}

// FuncStart returns the method "SpeechRecognition.start".
func (this SpeechRecognition) FuncStart() (fn js.Func[func()]) {
	bindings.FuncSpeechRecognitionStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start calls the method "SpeechRecognition.start".
func (this SpeechRecognition) Start() (ret js.Void) {
	bindings.CallSpeechRecognitionStart(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStart calls the method "SpeechRecognition.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SpeechRecognition) TryStart() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechRecognitionStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStop returns true if the method "SpeechRecognition.stop" exists.
func (this SpeechRecognition) HasFuncStop() bool {
	return js.True == bindings.HasFuncSpeechRecognitionStop(
		this.ref,
	)
}

// FuncStop returns the method "SpeechRecognition.stop".
func (this SpeechRecognition) FuncStop() (fn js.Func[func()]) {
	bindings.FuncSpeechRecognitionStop(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stop calls the method "SpeechRecognition.stop".
func (this SpeechRecognition) Stop() (ret js.Void) {
	bindings.CallSpeechRecognitionStop(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "SpeechRecognition.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SpeechRecognition) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechRecognitionStop(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAbort returns true if the method "SpeechRecognition.abort" exists.
func (this SpeechRecognition) HasFuncAbort() bool {
	return js.True == bindings.HasFuncSpeechRecognitionAbort(
		this.ref,
	)
}

// FuncAbort returns the method "SpeechRecognition.abort".
func (this SpeechRecognition) FuncAbort() (fn js.Func[func()]) {
	bindings.FuncSpeechRecognitionAbort(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort calls the method "SpeechRecognition.abort".
func (this SpeechRecognition) Abort() (ret js.Void) {
	bindings.CallSpeechRecognitionAbort(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAbort calls the method "SpeechRecognition.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SpeechRecognition) TryAbort() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechRecognitionAbort(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SpeechRecognitionAlternative struct {
	ref js.Ref
}

func (this SpeechRecognitionAlternative) Once() SpeechRecognitionAlternative {
	this.ref.Once()
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
	this.ref.Free()
}

// Transcript returns the value of property "SpeechRecognitionAlternative.transcript".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognitionAlternative) Transcript() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionAlternativeTranscript(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Confidence returns the value of property "SpeechRecognitionAlternative.confidence".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognitionAlternative) Confidence() (ret float32, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionAlternativeConfidence(
		this.ref, js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "SpeechRecognitionErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "SpeechRecognitionErrorEventInit.composed"
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

// FromRef calls UpdateFrom and returns a SpeechRecognitionErrorEventInit with all fields set.
func (p SpeechRecognitionErrorEventInit) FromRef(ref js.Ref) SpeechRecognitionErrorEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SpeechRecognitionErrorEventInit in the application heap.
func (p SpeechRecognitionErrorEventInit) New() js.Ref {
	return bindings.SpeechRecognitionErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SpeechRecognitionErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.SpeechRecognitionErrorEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SpeechRecognitionErrorEventInit) Update(ref js.Ref) {
	bindings.SpeechRecognitionErrorEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SpeechRecognitionErrorEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
}

func NewSpeechRecognitionErrorEvent(typ js.String, eventInitDict SpeechRecognitionErrorEventInit) (ret SpeechRecognitionErrorEvent) {
	ret.ref = bindings.NewSpeechRecognitionErrorEventBySpeechRecognitionErrorEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type SpeechRecognitionErrorEvent struct {
	Event
}

func (this SpeechRecognitionErrorEvent) Once() SpeechRecognitionErrorEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Error returns the value of property "SpeechRecognitionErrorEvent.error".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognitionErrorEvent) Error() (ret SpeechRecognitionErrorCode, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionErrorEventError(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "SpeechRecognitionErrorEvent.message".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognitionErrorEvent) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionErrorEventMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SpeechRecognitionResult struct {
	ref js.Ref
}

func (this SpeechRecognitionResult) Once() SpeechRecognitionResult {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "SpeechRecognitionResult.length".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognitionResult) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionResultLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsFinal returns the value of property "SpeechRecognitionResult.isFinal".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognitionResult) IsFinal() (ret bool, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionResultIsFinal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "SpeechRecognitionResult.item" exists.
func (this SpeechRecognitionResult) HasFuncItem() bool {
	return js.True == bindings.HasFuncSpeechRecognitionResultItem(
		this.ref,
	)
}

// FuncItem returns the method "SpeechRecognitionResult.item".
func (this SpeechRecognitionResult) FuncItem() (fn js.Func[func(index uint32) SpeechRecognitionAlternative]) {
	bindings.FuncSpeechRecognitionResultItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "SpeechRecognitionResult.item".
func (this SpeechRecognitionResult) Item(index uint32) (ret SpeechRecognitionAlternative) {
	bindings.CallSpeechRecognitionResultItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "SpeechRecognitionResult.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SpeechRecognitionResult) TryItem(index uint32) (ret SpeechRecognitionAlternative, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechRecognitionResultItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type SpeechRecognitionResultList struct {
	ref js.Ref
}

func (this SpeechRecognitionResultList) Once() SpeechRecognitionResultList {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "SpeechRecognitionResultList.length".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognitionResultList) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionResultListLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "SpeechRecognitionResultList.item" exists.
func (this SpeechRecognitionResultList) HasFuncItem() bool {
	return js.True == bindings.HasFuncSpeechRecognitionResultListItem(
		this.ref,
	)
}

// FuncItem returns the method "SpeechRecognitionResultList.item".
func (this SpeechRecognitionResultList) FuncItem() (fn js.Func[func(index uint32) SpeechRecognitionResult]) {
	bindings.FuncSpeechRecognitionResultListItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "SpeechRecognitionResultList.item".
func (this SpeechRecognitionResultList) Item(index uint32) (ret SpeechRecognitionResult) {
	bindings.CallSpeechRecognitionResultListItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "SpeechRecognitionResultList.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this SpeechRecognitionResultList) TryItem(index uint32) (ret SpeechRecognitionResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySpeechRecognitionResultListItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

type SpeechRecognitionEventInit struct {
	// ResultIndex is "SpeechRecognitionEventInit.resultIndex"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ResultIndex MUST be set to true to make this field effective.
	ResultIndex uint32
	// Results is "SpeechRecognitionEventInit.results"
	//
	// Required
	Results SpeechRecognitionResultList
	// Bubbles is "SpeechRecognitionEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "SpeechRecognitionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "SpeechRecognitionEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new SpeechRecognitionEventInit in the application heap.
func (p SpeechRecognitionEventInit) New() js.Ref {
	return bindings.SpeechRecognitionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SpeechRecognitionEventInit) UpdateFrom(ref js.Ref) {
	bindings.SpeechRecognitionEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SpeechRecognitionEventInit) Update(ref js.Ref) {
	bindings.SpeechRecognitionEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SpeechRecognitionEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Results.Ref(),
	)
	p.Results = p.Results.FromRef(js.Undefined)
}

func NewSpeechRecognitionEvent(typ js.String, eventInitDict SpeechRecognitionEventInit) (ret SpeechRecognitionEvent) {
	ret.ref = bindings.NewSpeechRecognitionEventBySpeechRecognitionEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type SpeechRecognitionEvent struct {
	Event
}

func (this SpeechRecognitionEvent) Once() SpeechRecognitionEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// ResultIndex returns the value of property "SpeechRecognitionEvent.resultIndex".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognitionEvent) ResultIndex() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionEventResultIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Results returns the value of property "SpeechRecognitionEvent.results".
//
// It returns ok=false if there is no such property.
func (this SpeechRecognitionEvent) Results() (ret SpeechRecognitionResultList, ok bool) {
	ok = js.True == bindings.GetSpeechRecognitionEventResults(
		this.ref, js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_CharIndex MUST be set to true to make this field effective.
	CharIndex uint32
	// CharLength is "SpeechSynthesisErrorEventInit.charLength"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_CharLength MUST be set to true to make this field effective.
	CharLength uint32
	// ElapsedTime is "SpeechSynthesisErrorEventInit.elapsedTime"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ElapsedTime MUST be set to true to make this field effective.
	ElapsedTime float32
	// Name is "SpeechSynthesisErrorEventInit.name"
	//
	// Optional, defaults to "".
	Name js.String
	// Bubbles is "SpeechSynthesisErrorEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "SpeechSynthesisErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "SpeechSynthesisErrorEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new SpeechSynthesisErrorEventInit in the application heap.
func (p SpeechSynthesisErrorEventInit) New() js.Ref {
	return bindings.SpeechSynthesisErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SpeechSynthesisErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.SpeechSynthesisErrorEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SpeechSynthesisErrorEventInit) Update(ref js.Ref) {
	bindings.SpeechSynthesisErrorEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SpeechSynthesisErrorEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Utterance.Ref(),
		p.Name.Ref(),
	)
	p.Utterance = p.Utterance.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

func NewSpeechSynthesisErrorEvent(typ js.String, eventInitDict SpeechSynthesisErrorEventInit) (ret SpeechSynthesisErrorEvent) {
	ret.ref = bindings.NewSpeechSynthesisErrorEventBySpeechSynthesisErrorEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type SpeechSynthesisErrorEvent struct {
	SpeechSynthesisEvent
}

func (this SpeechSynthesisErrorEvent) Once() SpeechSynthesisErrorEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Error returns the value of property "SpeechSynthesisErrorEvent.error".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisErrorEvent) Error() (ret SpeechSynthesisErrorCode, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisErrorEventError(
		this.ref, js.Pointer(&ret),
	)
	return
}

type SpeechSynthesisEventInit struct {
	// Utterance is "SpeechSynthesisEventInit.utterance"
	//
	// Required
	Utterance SpeechSynthesisUtterance
	// CharIndex is "SpeechSynthesisEventInit.charIndex"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_CharIndex MUST be set to true to make this field effective.
	CharIndex uint32
	// CharLength is "SpeechSynthesisEventInit.charLength"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_CharLength MUST be set to true to make this field effective.
	CharLength uint32
	// ElapsedTime is "SpeechSynthesisEventInit.elapsedTime"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_ElapsedTime MUST be set to true to make this field effective.
	ElapsedTime float32
	// Name is "SpeechSynthesisEventInit.name"
	//
	// Optional, defaults to "".
	Name js.String
	// Bubbles is "SpeechSynthesisEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "SpeechSynthesisEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "SpeechSynthesisEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new SpeechSynthesisEventInit in the application heap.
func (p SpeechSynthesisEventInit) New() js.Ref {
	return bindings.SpeechSynthesisEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SpeechSynthesisEventInit) UpdateFrom(ref js.Ref) {
	bindings.SpeechSynthesisEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SpeechSynthesisEventInit) Update(ref js.Ref) {
	bindings.SpeechSynthesisEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SpeechSynthesisEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Utterance.Ref(),
		p.Name.Ref(),
	)
	p.Utterance = p.Utterance.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
}

func NewSpeechSynthesisEvent(typ js.String, eventInitDict SpeechSynthesisEventInit) (ret SpeechSynthesisEvent) {
	ret.ref = bindings.NewSpeechSynthesisEventBySpeechSynthesisEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type SpeechSynthesisEvent struct {
	Event
}

func (this SpeechSynthesisEvent) Once() SpeechSynthesisEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Utterance returns the value of property "SpeechSynthesisEvent.utterance".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisEvent) Utterance() (ret SpeechSynthesisUtterance, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisEventUtterance(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CharIndex returns the value of property "SpeechSynthesisEvent.charIndex".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisEvent) CharIndex() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisEventCharIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CharLength returns the value of property "SpeechSynthesisEvent.charLength".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisEvent) CharLength() (ret uint32, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisEventCharLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ElapsedTime returns the value of property "SpeechSynthesisEvent.elapsedTime".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisEvent) ElapsedTime() (ret float32, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisEventElapsedTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "SpeechSynthesisEvent.name".
//
// It returns ok=false if there is no such property.
func (this SpeechSynthesisEvent) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetSpeechSynthesisEventName(
		this.ref, js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "StorageEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "StorageEventInit.composed"
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

// FromRef calls UpdateFrom and returns a StorageEventInit with all fields set.
func (p StorageEventInit) FromRef(ref js.Ref) StorageEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new StorageEventInit in the application heap.
func (p StorageEventInit) New() js.Ref {
	return bindings.StorageEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *StorageEventInit) UpdateFrom(ref js.Ref) {
	bindings.StorageEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *StorageEventInit) Update(ref js.Ref) {
	bindings.StorageEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *StorageEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Key.Ref(),
		p.OldValue.Ref(),
		p.NewValue.Ref(),
		p.Url.Ref(),
		p.StorageArea.Ref(),
	)
	p.Key = p.Key.FromRef(js.Undefined)
	p.OldValue = p.OldValue.FromRef(js.Undefined)
	p.NewValue = p.NewValue.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.StorageArea = p.StorageArea.FromRef(js.Undefined)
}

func NewStorageEvent(typ js.String, eventInitDict StorageEventInit) (ret StorageEvent) {
	ret.ref = bindings.NewStorageEventByStorageEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewStorageEventByStorageEvent1(typ js.String) (ret StorageEvent) {
	ret.ref = bindings.NewStorageEventByStorageEvent1(
		typ.Ref())
	return
}

type StorageEvent struct {
	Event
}

func (this StorageEvent) Once() StorageEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Key returns the value of property "StorageEvent.key".
//
// It returns ok=false if there is no such property.
func (this StorageEvent) Key() (ret js.String, ok bool) {
	ok = js.True == bindings.GetStorageEventKey(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OldValue returns the value of property "StorageEvent.oldValue".
//
// It returns ok=false if there is no such property.
func (this StorageEvent) OldValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetStorageEventOldValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NewValue returns the value of property "StorageEvent.newValue".
//
// It returns ok=false if there is no such property.
func (this StorageEvent) NewValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetStorageEventNewValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "StorageEvent.url".
//
// It returns ok=false if there is no such property.
func (this StorageEvent) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetStorageEventUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// StorageArea returns the value of property "StorageEvent.storageArea".
//
// It returns ok=false if there is no such property.
func (this StorageEvent) StorageArea() (ret Storage, ok bool) {
	ok = js.True == bindings.GetStorageEventStorageArea(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncInitStorageEvent returns true if the method "StorageEvent.initStorageEvent" exists.
func (this StorageEvent) HasFuncInitStorageEvent() bool {
	return js.True == bindings.HasFuncStorageEventInitStorageEvent(
		this.ref,
	)
}

// FuncInitStorageEvent returns the method "StorageEvent.initStorageEvent".
func (this StorageEvent) FuncInitStorageEvent() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String, url js.String, storageArea Storage)]) {
	bindings.FuncStorageEventInitStorageEvent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitStorageEvent calls the method "StorageEvent.initStorageEvent".
func (this StorageEvent) InitStorageEvent(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String, url js.String, storageArea Storage) (ret js.Void) {
	bindings.CallStorageEventInitStorageEvent(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
		newValue.Ref(),
		url.Ref(),
		storageArea.Ref(),
	)

	return
}

// TryInitStorageEvent calls the method "StorageEvent.initStorageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageEvent) TryInitStorageEvent(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String, url js.String, storageArea Storage) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageEventInitStorageEvent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
		newValue.Ref(),
		url.Ref(),
		storageArea.Ref(),
	)

	return
}

// HasFuncInitStorageEvent1 returns true if the method "StorageEvent.initStorageEvent" exists.
func (this StorageEvent) HasFuncInitStorageEvent1() bool {
	return js.True == bindings.HasFuncStorageEventInitStorageEvent1(
		this.ref,
	)
}

// FuncInitStorageEvent1 returns the method "StorageEvent.initStorageEvent".
func (this StorageEvent) FuncInitStorageEvent1() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String, url js.String)]) {
	bindings.FuncStorageEventInitStorageEvent1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitStorageEvent1 calls the method "StorageEvent.initStorageEvent".
func (this StorageEvent) InitStorageEvent1(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String, url js.String) (ret js.Void) {
	bindings.CallStorageEventInitStorageEvent1(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
		newValue.Ref(),
		url.Ref(),
	)

	return
}

// TryInitStorageEvent1 calls the method "StorageEvent.initStorageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageEvent) TryInitStorageEvent1(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String, url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageEventInitStorageEvent1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
		newValue.Ref(),
		url.Ref(),
	)

	return
}

// HasFuncInitStorageEvent2 returns true if the method "StorageEvent.initStorageEvent" exists.
func (this StorageEvent) HasFuncInitStorageEvent2() bool {
	return js.True == bindings.HasFuncStorageEventInitStorageEvent2(
		this.ref,
	)
}

// FuncInitStorageEvent2 returns the method "StorageEvent.initStorageEvent".
func (this StorageEvent) FuncInitStorageEvent2() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String)]) {
	bindings.FuncStorageEventInitStorageEvent2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitStorageEvent2 calls the method "StorageEvent.initStorageEvent".
func (this StorageEvent) InitStorageEvent2(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String) (ret js.Void) {
	bindings.CallStorageEventInitStorageEvent2(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
		newValue.Ref(),
	)

	return
}

// TryInitStorageEvent2 calls the method "StorageEvent.initStorageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageEvent) TryInitStorageEvent2(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String, newValue js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageEventInitStorageEvent2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
		newValue.Ref(),
	)

	return
}

// HasFuncInitStorageEvent3 returns true if the method "StorageEvent.initStorageEvent" exists.
func (this StorageEvent) HasFuncInitStorageEvent3() bool {
	return js.True == bindings.HasFuncStorageEventInitStorageEvent3(
		this.ref,
	)
}

// FuncInitStorageEvent3 returns the method "StorageEvent.initStorageEvent".
func (this StorageEvent) FuncInitStorageEvent3() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String)]) {
	bindings.FuncStorageEventInitStorageEvent3(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitStorageEvent3 calls the method "StorageEvent.initStorageEvent".
func (this StorageEvent) InitStorageEvent3(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String) (ret js.Void) {
	bindings.CallStorageEventInitStorageEvent3(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
	)

	return
}

// TryInitStorageEvent3 calls the method "StorageEvent.initStorageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageEvent) TryInitStorageEvent3(typ js.String, bubbles bool, cancelable bool, key js.String, oldValue js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageEventInitStorageEvent3(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
		oldValue.Ref(),
	)

	return
}

// HasFuncInitStorageEvent4 returns true if the method "StorageEvent.initStorageEvent" exists.
func (this StorageEvent) HasFuncInitStorageEvent4() bool {
	return js.True == bindings.HasFuncStorageEventInitStorageEvent4(
		this.ref,
	)
}

// FuncInitStorageEvent4 returns the method "StorageEvent.initStorageEvent".
func (this StorageEvent) FuncInitStorageEvent4() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool, key js.String)]) {
	bindings.FuncStorageEventInitStorageEvent4(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitStorageEvent4 calls the method "StorageEvent.initStorageEvent".
func (this StorageEvent) InitStorageEvent4(typ js.String, bubbles bool, cancelable bool, key js.String) (ret js.Void) {
	bindings.CallStorageEventInitStorageEvent4(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
	)

	return
}

// TryInitStorageEvent4 calls the method "StorageEvent.initStorageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageEvent) TryInitStorageEvent4(typ js.String, bubbles bool, cancelable bool, key js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageEventInitStorageEvent4(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
		key.Ref(),
	)

	return
}

// HasFuncInitStorageEvent5 returns true if the method "StorageEvent.initStorageEvent" exists.
func (this StorageEvent) HasFuncInitStorageEvent5() bool {
	return js.True == bindings.HasFuncStorageEventInitStorageEvent5(
		this.ref,
	)
}

// FuncInitStorageEvent5 returns the method "StorageEvent.initStorageEvent".
func (this StorageEvent) FuncInitStorageEvent5() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool)]) {
	bindings.FuncStorageEventInitStorageEvent5(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitStorageEvent5 calls the method "StorageEvent.initStorageEvent".
func (this StorageEvent) InitStorageEvent5(typ js.String, bubbles bool, cancelable bool) (ret js.Void) {
	bindings.CallStorageEventInitStorageEvent5(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// TryInitStorageEvent5 calls the method "StorageEvent.initStorageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageEvent) TryInitStorageEvent5(typ js.String, bubbles bool, cancelable bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageEventInitStorageEvent5(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// HasFuncInitStorageEvent6 returns true if the method "StorageEvent.initStorageEvent" exists.
func (this StorageEvent) HasFuncInitStorageEvent6() bool {
	return js.True == bindings.HasFuncStorageEventInitStorageEvent6(
		this.ref,
	)
}

// FuncInitStorageEvent6 returns the method "StorageEvent.initStorageEvent".
func (this StorageEvent) FuncInitStorageEvent6() (fn js.Func[func(typ js.String, bubbles bool)]) {
	bindings.FuncStorageEventInitStorageEvent6(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitStorageEvent6 calls the method "StorageEvent.initStorageEvent".
func (this StorageEvent) InitStorageEvent6(typ js.String, bubbles bool) (ret js.Void) {
	bindings.CallStorageEventInitStorageEvent6(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// TryInitStorageEvent6 calls the method "StorageEvent.initStorageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageEvent) TryInitStorageEvent6(typ js.String, bubbles bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageEventInitStorageEvent6(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// HasFuncInitStorageEvent7 returns true if the method "StorageEvent.initStorageEvent" exists.
func (this StorageEvent) HasFuncInitStorageEvent7() bool {
	return js.True == bindings.HasFuncStorageEventInitStorageEvent7(
		this.ref,
	)
}

// FuncInitStorageEvent7 returns the method "StorageEvent.initStorageEvent".
func (this StorageEvent) FuncInitStorageEvent7() (fn js.Func[func(typ js.String)]) {
	bindings.FuncStorageEventInitStorageEvent7(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitStorageEvent7 calls the method "StorageEvent.initStorageEvent".
func (this StorageEvent) InitStorageEvent7(typ js.String) (ret js.Void) {
	bindings.CallStorageEventInitStorageEvent7(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryInitStorageEvent7 calls the method "StorageEvent.initStorageEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this StorageEvent) TryInitStorageEvent7(typ js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStorageEventInitStorageEvent7(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
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
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "StyleSheet.type".
//
// It returns ok=false if there is no such property.
func (this StyleSheet) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetStyleSheetType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Href returns the value of property "StyleSheet.href".
//
// It returns ok=false if there is no such property.
func (this StyleSheet) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetStyleSheetHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

// OwnerNode returns the value of property "StyleSheet.ownerNode".
//
// It returns ok=false if there is no such property.
func (this StyleSheet) OwnerNode() (ret OneOf_Element_ProcessingInstruction, ok bool) {
	ok = js.True == bindings.GetStyleSheetOwnerNode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ParentStyleSheet returns the value of property "StyleSheet.parentStyleSheet".
//
// It returns ok=false if there is no such property.
func (this StyleSheet) ParentStyleSheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetStyleSheetParentStyleSheet(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Title returns the value of property "StyleSheet.title".
//
// It returns ok=false if there is no such property.
func (this StyleSheet) Title() (ret js.String, ok bool) {
	ok = js.True == bindings.GetStyleSheetTitle(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Media returns the value of property "StyleSheet.media".
//
// It returns ok=false if there is no such property.
func (this StyleSheet) Media() (ret MediaList, ok bool) {
	ok = js.True == bindings.GetStyleSheetMedia(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Disabled returns the value of property "StyleSheet.disabled".
//
// It returns ok=false if there is no such property.
func (this StyleSheet) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetStyleSheetDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "StyleSheet.disabled" to val.
//
// It returns false if the property cannot be set.
func (this StyleSheet) SetDisabled(val bool) bool {
	return js.True == bindings.SetStyleSheetDisabled(
		this.ref,
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "SubmitEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "SubmitEventInit.composed"
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

// FromRef calls UpdateFrom and returns a SubmitEventInit with all fields set.
func (p SubmitEventInit) FromRef(ref js.Ref) SubmitEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SubmitEventInit in the application heap.
func (p SubmitEventInit) New() js.Ref {
	return bindings.SubmitEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SubmitEventInit) UpdateFrom(ref js.Ref) {
	bindings.SubmitEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SubmitEventInit) Update(ref js.Ref) {
	bindings.SubmitEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SubmitEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Submitter.Ref(),
	)
	p.Submitter = p.Submitter.FromRef(js.Undefined)
}

func NewSubmitEvent(typ js.String, eventInitDict SubmitEventInit) (ret SubmitEvent) {
	ret.ref = bindings.NewSubmitEventBySubmitEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewSubmitEventBySubmitEvent1(typ js.String) (ret SubmitEvent) {
	ret.ref = bindings.NewSubmitEventBySubmitEvent1(
		typ.Ref())
	return
}

type SubmitEvent struct {
	Event
}

func (this SubmitEvent) Once() SubmitEvent {
	this.ref.Once()
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
	this.ref.Free()
}

// Submitter returns the value of property "SubmitEvent.submitter".
//
// It returns ok=false if there is no such property.
func (this SubmitEvent) Submitter() (ret HTMLElement, ok bool) {
	ok = js.True == bindings.GetSubmitEventSubmitter(
		this.ref, js.Pointer(&ret),
	)
	return
}

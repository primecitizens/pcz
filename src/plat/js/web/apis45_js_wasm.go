// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type PaymentDetailsModifier struct {
	// SupportedMethods is "PaymentDetailsModifier.supportedMethods"
	//
	// Required
	SupportedMethods js.String
	// Total is "PaymentDetailsModifier.total"
	//
	// Optional
	//
	// NOTE: Total.FFI_USE MUST be set to true to get Total used.
	Total PaymentItem
	// AdditionalDisplayItems is "PaymentDetailsModifier.additionalDisplayItems"
	//
	// Optional
	AdditionalDisplayItems js.Array[PaymentItem]
	// Data is "PaymentDetailsModifier.data"
	//
	// Optional
	Data js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentDetailsModifier with all fields set.
func (p PaymentDetailsModifier) FromRef(ref js.Ref) PaymentDetailsModifier {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentDetailsModifier in the application heap.
func (p PaymentDetailsModifier) New() js.Ref {
	return bindings.PaymentDetailsModifierJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentDetailsModifier) UpdateFrom(ref js.Ref) {
	bindings.PaymentDetailsModifierJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentDetailsModifier) Update(ref js.Ref) {
	bindings.PaymentDetailsModifierJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaymentDetailsBase struct {
	// DisplayItems is "PaymentDetailsBase.displayItems"
	//
	// Optional
	DisplayItems js.Array[PaymentItem]
	// Modifiers is "PaymentDetailsBase.modifiers"
	//
	// Optional
	Modifiers js.Array[PaymentDetailsModifier]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentDetailsBase with all fields set.
func (p PaymentDetailsBase) FromRef(ref js.Ref) PaymentDetailsBase {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentDetailsBase in the application heap.
func (p PaymentDetailsBase) New() js.Ref {
	return bindings.PaymentDetailsBaseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentDetailsBase) UpdateFrom(ref js.Ref) {
	bindings.PaymentDetailsBaseJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentDetailsBase) Update(ref js.Ref) {
	bindings.PaymentDetailsBaseJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaymentDetailsInit struct {
	// Id is "PaymentDetailsInit.id"
	//
	// Optional
	Id js.String
	// Total is "PaymentDetailsInit.total"
	//
	// Required
	//
	// NOTE: Total.FFI_USE MUST be set to true to get Total used.
	Total PaymentItem
	// DisplayItems is "PaymentDetailsInit.displayItems"
	//
	// Optional
	DisplayItems js.Array[PaymentItem]
	// Modifiers is "PaymentDetailsInit.modifiers"
	//
	// Optional
	Modifiers js.Array[PaymentDetailsModifier]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentDetailsInit with all fields set.
func (p PaymentDetailsInit) FromRef(ref js.Ref) PaymentDetailsInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentDetailsInit in the application heap.
func (p PaymentDetailsInit) New() js.Ref {
	return bindings.PaymentDetailsInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentDetailsInit) UpdateFrom(ref js.Ref) {
	bindings.PaymentDetailsInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentDetailsInit) Update(ref js.Ref) {
	bindings.PaymentDetailsInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaymentDetailsUpdate struct {
	// Total is "PaymentDetailsUpdate.total"
	//
	// Optional
	//
	// NOTE: Total.FFI_USE MUST be set to true to get Total used.
	Total PaymentItem
	// PaymentMethodErrors is "PaymentDetailsUpdate.paymentMethodErrors"
	//
	// Optional
	PaymentMethodErrors js.Object
	// DisplayItems is "PaymentDetailsUpdate.displayItems"
	//
	// Optional
	DisplayItems js.Array[PaymentItem]
	// Modifiers is "PaymentDetailsUpdate.modifiers"
	//
	// Optional
	Modifiers js.Array[PaymentDetailsModifier]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentDetailsUpdate with all fields set.
func (p PaymentDetailsUpdate) FromRef(ref js.Ref) PaymentDetailsUpdate {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentDetailsUpdate in the application heap.
func (p PaymentDetailsUpdate) New() js.Ref {
	return bindings.PaymentDetailsUpdateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentDetailsUpdate) UpdateFrom(ref js.Ref) {
	bindings.PaymentDetailsUpdateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentDetailsUpdate) Update(ref js.Ref) {
	bindings.PaymentDetailsUpdateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaymentHandlerResponse struct {
	// MethodName is "PaymentHandlerResponse.methodName"
	//
	// Optional
	MethodName js.String
	// Details is "PaymentHandlerResponse.details"
	//
	// Optional
	Details js.Object
	// PayerName is "PaymentHandlerResponse.payerName"
	//
	// Optional
	PayerName js.String
	// PayerEmail is "PaymentHandlerResponse.payerEmail"
	//
	// Optional
	PayerEmail js.String
	// PayerPhone is "PaymentHandlerResponse.payerPhone"
	//
	// Optional
	PayerPhone js.String
	// ShippingAddress is "PaymentHandlerResponse.shippingAddress"
	//
	// Optional
	//
	// NOTE: ShippingAddress.FFI_USE MUST be set to true to get ShippingAddress used.
	ShippingAddress AddressInit
	// ShippingOption is "PaymentHandlerResponse.shippingOption"
	//
	// Optional
	ShippingOption js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentHandlerResponse with all fields set.
func (p PaymentHandlerResponse) FromRef(ref js.Ref) PaymentHandlerResponse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentHandlerResponse in the application heap.
func (p PaymentHandlerResponse) New() js.Ref {
	return bindings.PaymentHandlerResponseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentHandlerResponse) UpdateFrom(ref js.Ref) {
	bindings.PaymentHandlerResponseJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentHandlerResponse) Update(ref js.Ref) {
	bindings.PaymentHandlerResponseJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaymentMethodChangeEventInit struct {
	// MethodName is "PaymentMethodChangeEventInit.methodName"
	//
	// Optional, defaults to "".
	MethodName js.String
	// MethodDetails is "PaymentMethodChangeEventInit.methodDetails"
	//
	// Optional, defaults to null.
	MethodDetails js.Object
	// Bubbles is "PaymentMethodChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PaymentMethodChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PaymentMethodChangeEventInit.composed"
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

// FromRef calls UpdateFrom and returns a PaymentMethodChangeEventInit with all fields set.
func (p PaymentMethodChangeEventInit) FromRef(ref js.Ref) PaymentMethodChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentMethodChangeEventInit in the application heap.
func (p PaymentMethodChangeEventInit) New() js.Ref {
	return bindings.PaymentMethodChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentMethodChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.PaymentMethodChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentMethodChangeEventInit) Update(ref js.Ref) {
	bindings.PaymentMethodChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPaymentMethodChangeEvent(typ js.String, eventInitDict PaymentMethodChangeEventInit) (ret PaymentMethodChangeEvent) {
	ret.ref = bindings.NewPaymentMethodChangeEventByPaymentMethodChangeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewPaymentMethodChangeEventByPaymentMethodChangeEvent1(typ js.String) (ret PaymentMethodChangeEvent) {
	ret.ref = bindings.NewPaymentMethodChangeEventByPaymentMethodChangeEvent1(
		typ.Ref())
	return
}

type PaymentMethodChangeEvent struct {
	PaymentRequestUpdateEvent
}

func (this PaymentMethodChangeEvent) Once() PaymentMethodChangeEvent {
	this.Ref().Once()
	return this
}

func (this PaymentMethodChangeEvent) Ref() js.Ref {
	return this.PaymentRequestUpdateEvent.Ref()
}

func (this PaymentMethodChangeEvent) FromRef(ref js.Ref) PaymentMethodChangeEvent {
	this.PaymentRequestUpdateEvent = this.PaymentRequestUpdateEvent.FromRef(ref)
	return this
}

func (this PaymentMethodChangeEvent) Free() {
	this.Ref().Free()
}

// MethodName returns the value of property "PaymentMethodChangeEvent.methodName".
//
// It returns ok=false if there is no such property.
func (this PaymentMethodChangeEvent) MethodName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPaymentMethodChangeEventMethodName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MethodDetails returns the value of property "PaymentMethodChangeEvent.methodDetails".
//
// It returns ok=false if there is no such property.
func (this PaymentMethodChangeEvent) MethodDetails() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetPaymentMethodChangeEventMethodDetails(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type PaymentMethodData struct {
	// SupportedMethods is "PaymentMethodData.supportedMethods"
	//
	// Required
	SupportedMethods js.String
	// Data is "PaymentMethodData.data"
	//
	// Optional
	Data js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentMethodData with all fields set.
func (p PaymentMethodData) FromRef(ref js.Ref) PaymentMethodData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentMethodData in the application heap.
func (p PaymentMethodData) New() js.Ref {
	return bindings.PaymentMethodDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentMethodData) UpdateFrom(ref js.Ref) {
	bindings.PaymentMethodDataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentMethodData) Update(ref js.Ref) {
	bindings.PaymentMethodDataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaymentShippingType uint32

const (
	_ PaymentShippingType = iota

	PaymentShippingType_SHIPPING
	PaymentShippingType_DELIVERY
	PaymentShippingType_PICKUP
)

func (PaymentShippingType) FromRef(str js.Ref) PaymentShippingType {
	return PaymentShippingType(bindings.ConstOfPaymentShippingType(str))
}

func (x PaymentShippingType) String() (string, bool) {
	switch x {
	case PaymentShippingType_SHIPPING:
		return "shipping", true
	case PaymentShippingType_DELIVERY:
		return "delivery", true
	case PaymentShippingType_PICKUP:
		return "pickup", true
	default:
		return "", false
	}
}

type PaymentOptions struct {
	// RequestPayerName is "PaymentOptions.requestPayerName"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_RequestPayerName MUST be set to true to make this field effective.
	RequestPayerName bool
	// RequestBillingAddress is "PaymentOptions.requestBillingAddress"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_RequestBillingAddress MUST be set to true to make this field effective.
	RequestBillingAddress bool
	// RequestPayerEmail is "PaymentOptions.requestPayerEmail"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_RequestPayerEmail MUST be set to true to make this field effective.
	RequestPayerEmail bool
	// RequestPayerPhone is "PaymentOptions.requestPayerPhone"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_RequestPayerPhone MUST be set to true to make this field effective.
	RequestPayerPhone bool
	// RequestShipping is "PaymentOptions.requestShipping"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_RequestShipping MUST be set to true to make this field effective.
	RequestShipping bool
	// ShippingType is "PaymentOptions.shippingType"
	//
	// Optional, defaults to "shipping".
	ShippingType PaymentShippingType

	FFI_USE_RequestPayerName      bool // for RequestPayerName.
	FFI_USE_RequestBillingAddress bool // for RequestBillingAddress.
	FFI_USE_RequestPayerEmail     bool // for RequestPayerEmail.
	FFI_USE_RequestPayerPhone     bool // for RequestPayerPhone.
	FFI_USE_RequestShipping       bool // for RequestShipping.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentOptions with all fields set.
func (p PaymentOptions) FromRef(ref js.Ref) PaymentOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentOptions in the application heap.
func (p PaymentOptions) New() js.Ref {
	return bindings.PaymentOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentOptions) UpdateFrom(ref js.Ref) {
	bindings.PaymentOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentOptions) Update(ref js.Ref) {
	bindings.PaymentOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaymentValidationErrors struct {
	// Error is "PaymentValidationErrors.error"
	//
	// Optional
	Error js.String
	// PaymentMethod is "PaymentValidationErrors.paymentMethod"
	//
	// Optional
	PaymentMethod js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentValidationErrors with all fields set.
func (p PaymentValidationErrors) FromRef(ref js.Ref) PaymentValidationErrors {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentValidationErrors in the application heap.
func (p PaymentValidationErrors) New() js.Ref {
	return bindings.PaymentValidationErrorsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentValidationErrors) UpdateFrom(ref js.Ref) {
	bindings.PaymentValidationErrorsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentValidationErrors) Update(ref js.Ref) {
	bindings.PaymentValidationErrorsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaymentResponse struct {
	EventTarget
}

func (this PaymentResponse) Once() PaymentResponse {
	this.Ref().Once()
	return this
}

func (this PaymentResponse) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this PaymentResponse) FromRef(ref js.Ref) PaymentResponse {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this PaymentResponse) Free() {
	this.Ref().Free()
}

// RequestId returns the value of property "PaymentResponse.requestId".
//
// It returns ok=false if there is no such property.
func (this PaymentResponse) RequestId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPaymentResponseRequestId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MethodName returns the value of property "PaymentResponse.methodName".
//
// It returns ok=false if there is no such property.
func (this PaymentResponse) MethodName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPaymentResponseMethodName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Details returns the value of property "PaymentResponse.details".
//
// It returns ok=false if there is no such property.
func (this PaymentResponse) Details() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetPaymentResponseDetails(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "PaymentResponse.toJSON" exists.
func (this PaymentResponse) HasToJSON() bool {
	return js.True == bindings.HasPaymentResponseToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "PaymentResponse.toJSON".
func (this PaymentResponse) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PaymentResponseToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PaymentResponse.toJSON".
func (this PaymentResponse) ToJSON() (ret js.Object) {
	bindings.CallPaymentResponseToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PaymentResponse.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentResponse) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentResponseToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasComplete returns true if the method "PaymentResponse.complete" exists.
func (this PaymentResponse) HasComplete() bool {
	return js.True == bindings.HasPaymentResponseComplete(
		this.Ref(),
	)
}

// CompleteFunc returns the method "PaymentResponse.complete".
func (this PaymentResponse) CompleteFunc() (fn js.Func[func(result PaymentComplete, details PaymentCompleteDetails) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentResponseCompleteFunc(
			this.Ref(),
		),
	)
}

// Complete calls the method "PaymentResponse.complete".
func (this PaymentResponse) Complete(result PaymentComplete, details PaymentCompleteDetails) (ret js.Promise[js.Void]) {
	bindings.CallPaymentResponseComplete(
		this.Ref(), js.Pointer(&ret),
		uint32(result),
		js.Pointer(&details),
	)

	return
}

// TryComplete calls the method "PaymentResponse.complete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentResponse) TryComplete(result PaymentComplete, details PaymentCompleteDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentResponseComplete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(result),
		js.Pointer(&details),
	)

	return
}

// HasComplete1 returns true if the method "PaymentResponse.complete" exists.
func (this PaymentResponse) HasComplete1() bool {
	return js.True == bindings.HasPaymentResponseComplete1(
		this.Ref(),
	)
}

// Complete1Func returns the method "PaymentResponse.complete".
func (this PaymentResponse) Complete1Func() (fn js.Func[func(result PaymentComplete) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentResponseComplete1Func(
			this.Ref(),
		),
	)
}

// Complete1 calls the method "PaymentResponse.complete".
func (this PaymentResponse) Complete1(result PaymentComplete) (ret js.Promise[js.Void]) {
	bindings.CallPaymentResponseComplete1(
		this.Ref(), js.Pointer(&ret),
		uint32(result),
	)

	return
}

// TryComplete1 calls the method "PaymentResponse.complete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentResponse) TryComplete1(result PaymentComplete) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentResponseComplete1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(result),
	)

	return
}

// HasComplete2 returns true if the method "PaymentResponse.complete" exists.
func (this PaymentResponse) HasComplete2() bool {
	return js.True == bindings.HasPaymentResponseComplete2(
		this.Ref(),
	)
}

// Complete2Func returns the method "PaymentResponse.complete".
func (this PaymentResponse) Complete2Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentResponseComplete2Func(
			this.Ref(),
		),
	)
}

// Complete2 calls the method "PaymentResponse.complete".
func (this PaymentResponse) Complete2() (ret js.Promise[js.Void]) {
	bindings.CallPaymentResponseComplete2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryComplete2 calls the method "PaymentResponse.complete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentResponse) TryComplete2() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentResponseComplete2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRetry returns true if the method "PaymentResponse.retry" exists.
func (this PaymentResponse) HasRetry() bool {
	return js.True == bindings.HasPaymentResponseRetry(
		this.Ref(),
	)
}

// RetryFunc returns the method "PaymentResponse.retry".
func (this PaymentResponse) RetryFunc() (fn js.Func[func(errorFields PaymentValidationErrors) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentResponseRetryFunc(
			this.Ref(),
		),
	)
}

// Retry calls the method "PaymentResponse.retry".
func (this PaymentResponse) Retry(errorFields PaymentValidationErrors) (ret js.Promise[js.Void]) {
	bindings.CallPaymentResponseRetry(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&errorFields),
	)

	return
}

// TryRetry calls the method "PaymentResponse.retry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentResponse) TryRetry(errorFields PaymentValidationErrors) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentResponseRetry(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&errorFields),
	)

	return
}

// HasRetry1 returns true if the method "PaymentResponse.retry" exists.
func (this PaymentResponse) HasRetry1() bool {
	return js.True == bindings.HasPaymentResponseRetry1(
		this.Ref(),
	)
}

// Retry1Func returns the method "PaymentResponse.retry".
func (this PaymentResponse) Retry1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentResponseRetry1Func(
			this.Ref(),
		),
	)
}

// Retry1 calls the method "PaymentResponse.retry".
func (this PaymentResponse) Retry1() (ret js.Promise[js.Void]) {
	bindings.CallPaymentResponseRetry1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRetry1 calls the method "PaymentResponse.retry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentResponse) TryRetry1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentResponseRetry1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewPaymentRequest(methodData js.Array[PaymentMethodData], details PaymentDetailsInit) (ret PaymentRequest) {
	ret.ref = bindings.NewPaymentRequestByPaymentRequest(
		methodData.Ref(),
		js.Pointer(&details))
	return
}

type PaymentRequest struct {
	EventTarget
}

func (this PaymentRequest) Once() PaymentRequest {
	this.Ref().Once()
	return this
}

func (this PaymentRequest) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this PaymentRequest) FromRef(ref js.Ref) PaymentRequest {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this PaymentRequest) Free() {
	this.Ref().Free()
}

// Id returns the value of property "PaymentRequest.id".
//
// It returns ok=false if there is no such property.
func (this PaymentRequest) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPaymentRequestId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasShow returns true if the method "PaymentRequest.show" exists.
func (this PaymentRequest) HasShow() bool {
	return js.True == bindings.HasPaymentRequestShow(
		this.Ref(),
	)
}

// ShowFunc returns the method "PaymentRequest.show".
func (this PaymentRequest) ShowFunc() (fn js.Func[func(detailsPromise js.Promise[PaymentDetailsUpdate]) js.Promise[PaymentResponse]]) {
	return fn.FromRef(
		bindings.PaymentRequestShowFunc(
			this.Ref(),
		),
	)
}

// Show calls the method "PaymentRequest.show".
func (this PaymentRequest) Show(detailsPromise js.Promise[PaymentDetailsUpdate]) (ret js.Promise[PaymentResponse]) {
	bindings.CallPaymentRequestShow(
		this.Ref(), js.Pointer(&ret),
		detailsPromise.Ref(),
	)

	return
}

// TryShow calls the method "PaymentRequest.show"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequest) TryShow(detailsPromise js.Promise[PaymentDetailsUpdate]) (ret js.Promise[PaymentResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestShow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		detailsPromise.Ref(),
	)

	return
}

// HasShow1 returns true if the method "PaymentRequest.show" exists.
func (this PaymentRequest) HasShow1() bool {
	return js.True == bindings.HasPaymentRequestShow1(
		this.Ref(),
	)
}

// Show1Func returns the method "PaymentRequest.show".
func (this PaymentRequest) Show1Func() (fn js.Func[func() js.Promise[PaymentResponse]]) {
	return fn.FromRef(
		bindings.PaymentRequestShow1Func(
			this.Ref(),
		),
	)
}

// Show1 calls the method "PaymentRequest.show".
func (this PaymentRequest) Show1() (ret js.Promise[PaymentResponse]) {
	bindings.CallPaymentRequestShow1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryShow1 calls the method "PaymentRequest.show"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequest) TryShow1() (ret js.Promise[PaymentResponse], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestShow1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAbort returns true if the method "PaymentRequest.abort" exists.
func (this PaymentRequest) HasAbort() bool {
	return js.True == bindings.HasPaymentRequestAbort(
		this.Ref(),
	)
}

// AbortFunc returns the method "PaymentRequest.abort".
func (this PaymentRequest) AbortFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentRequestAbortFunc(
			this.Ref(),
		),
	)
}

// Abort calls the method "PaymentRequest.abort".
func (this PaymentRequest) Abort() (ret js.Promise[js.Void]) {
	bindings.CallPaymentRequestAbort(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAbort calls the method "PaymentRequest.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequest) TryAbort() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestAbort(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCanMakePayment returns true if the method "PaymentRequest.canMakePayment" exists.
func (this PaymentRequest) HasCanMakePayment() bool {
	return js.True == bindings.HasPaymentRequestCanMakePayment(
		this.Ref(),
	)
}

// CanMakePaymentFunc returns the method "PaymentRequest.canMakePayment".
func (this PaymentRequest) CanMakePaymentFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.PaymentRequestCanMakePaymentFunc(
			this.Ref(),
		),
	)
}

// CanMakePayment calls the method "PaymentRequest.canMakePayment".
func (this PaymentRequest) CanMakePayment() (ret js.Promise[js.Boolean]) {
	bindings.CallPaymentRequestCanMakePayment(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCanMakePayment calls the method "PaymentRequest.canMakePayment"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequest) TryCanMakePayment() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestCanMakePayment(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsSecurePaymentConfirmationAvailable returns true if the staticmethod "PaymentRequest.isSecurePaymentConfirmationAvailable" exists.
func (this PaymentRequest) HasIsSecurePaymentConfirmationAvailable() bool {
	return js.True == bindings.HasPaymentRequestIsSecurePaymentConfirmationAvailable(
		this.Ref(),
	)
}

// IsSecurePaymentConfirmationAvailableFunc returns the staticmethod "PaymentRequest.isSecurePaymentConfirmationAvailable".
func (this PaymentRequest) IsSecurePaymentConfirmationAvailableFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.PaymentRequestIsSecurePaymentConfirmationAvailableFunc(
			this.Ref(),
		),
	)
}

// IsSecurePaymentConfirmationAvailable calls the staticmethod "PaymentRequest.isSecurePaymentConfirmationAvailable".
func (this PaymentRequest) IsSecurePaymentConfirmationAvailable() (ret js.Promise[js.Boolean]) {
	bindings.CallPaymentRequestIsSecurePaymentConfirmationAvailable(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIsSecurePaymentConfirmationAvailable calls the staticmethod "PaymentRequest.isSecurePaymentConfirmationAvailable"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequest) TryIsSecurePaymentConfirmationAvailable() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestIsSecurePaymentConfirmationAvailable(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PaymentShippingOption struct {
	// Id is "PaymentShippingOption.id"
	//
	// Required
	Id js.String
	// Label is "PaymentShippingOption.label"
	//
	// Required
	Label js.String
	// Amount is "PaymentShippingOption.amount"
	//
	// Required
	//
	// NOTE: Amount.FFI_USE MUST be set to true to get Amount used.
	Amount PaymentCurrencyAmount
	// Selected is "PaymentShippingOption.selected"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Selected MUST be set to true to make this field effective.
	Selected bool

	FFI_USE_Selected bool // for Selected.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentShippingOption with all fields set.
func (p PaymentShippingOption) FromRef(ref js.Ref) PaymentShippingOption {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentShippingOption in the application heap.
func (p PaymentShippingOption) New() js.Ref {
	return bindings.PaymentShippingOptionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentShippingOption) UpdateFrom(ref js.Ref) {
	bindings.PaymentShippingOptionJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentShippingOption) Update(ref js.Ref) {
	bindings.PaymentShippingOptionJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaymentRequestDetailsUpdate struct {
	// Error is "PaymentRequestDetailsUpdate.error"
	//
	// Optional
	Error js.String
	// Total is "PaymentRequestDetailsUpdate.total"
	//
	// Optional
	//
	// NOTE: Total.FFI_USE MUST be set to true to get Total used.
	Total PaymentCurrencyAmount
	// Modifiers is "PaymentRequestDetailsUpdate.modifiers"
	//
	// Optional
	Modifiers js.Array[PaymentDetailsModifier]
	// ShippingOptions is "PaymentRequestDetailsUpdate.shippingOptions"
	//
	// Optional
	ShippingOptions js.Array[PaymentShippingOption]
	// PaymentMethodErrors is "PaymentRequestDetailsUpdate.paymentMethodErrors"
	//
	// Optional
	PaymentMethodErrors js.Object
	// ShippingAddressErrors is "PaymentRequestDetailsUpdate.shippingAddressErrors"
	//
	// Optional
	//
	// NOTE: ShippingAddressErrors.FFI_USE MUST be set to true to get ShippingAddressErrors used.
	ShippingAddressErrors AddressErrors

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentRequestDetailsUpdate with all fields set.
func (p PaymentRequestDetailsUpdate) FromRef(ref js.Ref) PaymentRequestDetailsUpdate {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentRequestDetailsUpdate in the application heap.
func (p PaymentRequestDetailsUpdate) New() js.Ref {
	return bindings.PaymentRequestDetailsUpdateJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentRequestDetailsUpdate) UpdateFrom(ref js.Ref) {
	bindings.PaymentRequestDetailsUpdateJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentRequestDetailsUpdate) Update(ref js.Ref) {
	bindings.PaymentRequestDetailsUpdateJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaymentRequestEventInit struct {
	// TopOrigin is "PaymentRequestEventInit.topOrigin"
	//
	// Optional
	TopOrigin js.String
	// PaymentRequestOrigin is "PaymentRequestEventInit.paymentRequestOrigin"
	//
	// Optional
	PaymentRequestOrigin js.String
	// PaymentRequestId is "PaymentRequestEventInit.paymentRequestId"
	//
	// Optional
	PaymentRequestId js.String
	// MethodData is "PaymentRequestEventInit.methodData"
	//
	// Optional
	MethodData js.Array[PaymentMethodData]
	// Total is "PaymentRequestEventInit.total"
	//
	// Optional
	//
	// NOTE: Total.FFI_USE MUST be set to true to get Total used.
	Total PaymentCurrencyAmount
	// Modifiers is "PaymentRequestEventInit.modifiers"
	//
	// Optional
	Modifiers js.Array[PaymentDetailsModifier]
	// PaymentOptions is "PaymentRequestEventInit.paymentOptions"
	//
	// Optional
	//
	// NOTE: PaymentOptions.FFI_USE MUST be set to true to get PaymentOptions used.
	PaymentOptions PaymentOptions
	// ShippingOptions is "PaymentRequestEventInit.shippingOptions"
	//
	// Optional
	ShippingOptions js.Array[PaymentShippingOption]
	// Bubbles is "PaymentRequestEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PaymentRequestEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PaymentRequestEventInit.composed"
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

// FromRef calls UpdateFrom and returns a PaymentRequestEventInit with all fields set.
func (p PaymentRequestEventInit) FromRef(ref js.Ref) PaymentRequestEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentRequestEventInit in the application heap.
func (p PaymentRequestEventInit) New() js.Ref {
	return bindings.PaymentRequestEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentRequestEventInit) UpdateFrom(ref js.Ref) {
	bindings.PaymentRequestEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentRequestEventInit) Update(ref js.Ref) {
	bindings.PaymentRequestEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPaymentRequestEvent(typ js.String, eventInitDict PaymentRequestEventInit) (ret PaymentRequestEvent) {
	ret.ref = bindings.NewPaymentRequestEventByPaymentRequestEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewPaymentRequestEventByPaymentRequestEvent1(typ js.String) (ret PaymentRequestEvent) {
	ret.ref = bindings.NewPaymentRequestEventByPaymentRequestEvent1(
		typ.Ref())
	return
}

type PaymentRequestEvent struct {
	ExtendableEvent
}

func (this PaymentRequestEvent) Once() PaymentRequestEvent {
	this.Ref().Once()
	return this
}

func (this PaymentRequestEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this PaymentRequestEvent) FromRef(ref js.Ref) PaymentRequestEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this PaymentRequestEvent) Free() {
	this.Ref().Free()
}

// TopOrigin returns the value of property "PaymentRequestEvent.topOrigin".
//
// It returns ok=false if there is no such property.
func (this PaymentRequestEvent) TopOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPaymentRequestEventTopOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PaymentRequestOrigin returns the value of property "PaymentRequestEvent.paymentRequestOrigin".
//
// It returns ok=false if there is no such property.
func (this PaymentRequestEvent) PaymentRequestOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPaymentRequestEventPaymentRequestOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PaymentRequestId returns the value of property "PaymentRequestEvent.paymentRequestId".
//
// It returns ok=false if there is no such property.
func (this PaymentRequestEvent) PaymentRequestId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPaymentRequestEventPaymentRequestId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MethodData returns the value of property "PaymentRequestEvent.methodData".
//
// It returns ok=false if there is no such property.
func (this PaymentRequestEvent) MethodData() (ret js.FrozenArray[PaymentMethodData], ok bool) {
	ok = js.True == bindings.GetPaymentRequestEventMethodData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Total returns the value of property "PaymentRequestEvent.total".
//
// It returns ok=false if there is no such property.
func (this PaymentRequestEvent) Total() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetPaymentRequestEventTotal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Modifiers returns the value of property "PaymentRequestEvent.modifiers".
//
// It returns ok=false if there is no such property.
func (this PaymentRequestEvent) Modifiers() (ret js.FrozenArray[PaymentDetailsModifier], ok bool) {
	ok = js.True == bindings.GetPaymentRequestEventModifiers(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PaymentOptions returns the value of property "PaymentRequestEvent.paymentOptions".
//
// It returns ok=false if there is no such property.
func (this PaymentRequestEvent) PaymentOptions() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetPaymentRequestEventPaymentOptions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ShippingOptions returns the value of property "PaymentRequestEvent.shippingOptions".
//
// It returns ok=false if there is no such property.
func (this PaymentRequestEvent) ShippingOptions() (ret js.FrozenArray[PaymentShippingOption], ok bool) {
	ok = js.True == bindings.GetPaymentRequestEventShippingOptions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasOpenWindow returns true if the method "PaymentRequestEvent.openWindow" exists.
func (this PaymentRequestEvent) HasOpenWindow() bool {
	return js.True == bindings.HasPaymentRequestEventOpenWindow(
		this.Ref(),
	)
}

// OpenWindowFunc returns the method "PaymentRequestEvent.openWindow".
func (this PaymentRequestEvent) OpenWindowFunc() (fn js.Func[func(url js.String) js.Promise[WindowClient]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventOpenWindowFunc(
			this.Ref(),
		),
	)
}

// OpenWindow calls the method "PaymentRequestEvent.openWindow".
func (this PaymentRequestEvent) OpenWindow(url js.String) (ret js.Promise[WindowClient]) {
	bindings.CallPaymentRequestEventOpenWindow(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryOpenWindow calls the method "PaymentRequestEvent.openWindow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequestEvent) TryOpenWindow(url js.String) (ret js.Promise[WindowClient], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestEventOpenWindow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasChangePaymentMethod returns true if the method "PaymentRequestEvent.changePaymentMethod" exists.
func (this PaymentRequestEvent) HasChangePaymentMethod() bool {
	return js.True == bindings.HasPaymentRequestEventChangePaymentMethod(
		this.Ref(),
	)
}

// ChangePaymentMethodFunc returns the method "PaymentRequestEvent.changePaymentMethod".
func (this PaymentRequestEvent) ChangePaymentMethodFunc() (fn js.Func[func(methodName js.String, methodDetails js.Object) js.Promise[PaymentRequestDetailsUpdate]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventChangePaymentMethodFunc(
			this.Ref(),
		),
	)
}

// ChangePaymentMethod calls the method "PaymentRequestEvent.changePaymentMethod".
func (this PaymentRequestEvent) ChangePaymentMethod(methodName js.String, methodDetails js.Object) (ret js.Promise[PaymentRequestDetailsUpdate]) {
	bindings.CallPaymentRequestEventChangePaymentMethod(
		this.Ref(), js.Pointer(&ret),
		methodName.Ref(),
		methodDetails.Ref(),
	)

	return
}

// TryChangePaymentMethod calls the method "PaymentRequestEvent.changePaymentMethod"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequestEvent) TryChangePaymentMethod(methodName js.String, methodDetails js.Object) (ret js.Promise[PaymentRequestDetailsUpdate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestEventChangePaymentMethod(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		methodName.Ref(),
		methodDetails.Ref(),
	)

	return
}

// HasChangePaymentMethod1 returns true if the method "PaymentRequestEvent.changePaymentMethod" exists.
func (this PaymentRequestEvent) HasChangePaymentMethod1() bool {
	return js.True == bindings.HasPaymentRequestEventChangePaymentMethod1(
		this.Ref(),
	)
}

// ChangePaymentMethod1Func returns the method "PaymentRequestEvent.changePaymentMethod".
func (this PaymentRequestEvent) ChangePaymentMethod1Func() (fn js.Func[func(methodName js.String) js.Promise[PaymentRequestDetailsUpdate]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventChangePaymentMethod1Func(
			this.Ref(),
		),
	)
}

// ChangePaymentMethod1 calls the method "PaymentRequestEvent.changePaymentMethod".
func (this PaymentRequestEvent) ChangePaymentMethod1(methodName js.String) (ret js.Promise[PaymentRequestDetailsUpdate]) {
	bindings.CallPaymentRequestEventChangePaymentMethod1(
		this.Ref(), js.Pointer(&ret),
		methodName.Ref(),
	)

	return
}

// TryChangePaymentMethod1 calls the method "PaymentRequestEvent.changePaymentMethod"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequestEvent) TryChangePaymentMethod1(methodName js.String) (ret js.Promise[PaymentRequestDetailsUpdate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestEventChangePaymentMethod1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		methodName.Ref(),
	)

	return
}

// HasChangeShippingAddress returns true if the method "PaymentRequestEvent.changeShippingAddress" exists.
func (this PaymentRequestEvent) HasChangeShippingAddress() bool {
	return js.True == bindings.HasPaymentRequestEventChangeShippingAddress(
		this.Ref(),
	)
}

// ChangeShippingAddressFunc returns the method "PaymentRequestEvent.changeShippingAddress".
func (this PaymentRequestEvent) ChangeShippingAddressFunc() (fn js.Func[func(shippingAddress AddressInit) js.Promise[PaymentRequestDetailsUpdate]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventChangeShippingAddressFunc(
			this.Ref(),
		),
	)
}

// ChangeShippingAddress calls the method "PaymentRequestEvent.changeShippingAddress".
func (this PaymentRequestEvent) ChangeShippingAddress(shippingAddress AddressInit) (ret js.Promise[PaymentRequestDetailsUpdate]) {
	bindings.CallPaymentRequestEventChangeShippingAddress(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&shippingAddress),
	)

	return
}

// TryChangeShippingAddress calls the method "PaymentRequestEvent.changeShippingAddress"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequestEvent) TryChangeShippingAddress(shippingAddress AddressInit) (ret js.Promise[PaymentRequestDetailsUpdate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestEventChangeShippingAddress(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&shippingAddress),
	)

	return
}

// HasChangeShippingAddress1 returns true if the method "PaymentRequestEvent.changeShippingAddress" exists.
func (this PaymentRequestEvent) HasChangeShippingAddress1() bool {
	return js.True == bindings.HasPaymentRequestEventChangeShippingAddress1(
		this.Ref(),
	)
}

// ChangeShippingAddress1Func returns the method "PaymentRequestEvent.changeShippingAddress".
func (this PaymentRequestEvent) ChangeShippingAddress1Func() (fn js.Func[func() js.Promise[PaymentRequestDetailsUpdate]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventChangeShippingAddress1Func(
			this.Ref(),
		),
	)
}

// ChangeShippingAddress1 calls the method "PaymentRequestEvent.changeShippingAddress".
func (this PaymentRequestEvent) ChangeShippingAddress1() (ret js.Promise[PaymentRequestDetailsUpdate]) {
	bindings.CallPaymentRequestEventChangeShippingAddress1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryChangeShippingAddress1 calls the method "PaymentRequestEvent.changeShippingAddress"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequestEvent) TryChangeShippingAddress1() (ret js.Promise[PaymentRequestDetailsUpdate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestEventChangeShippingAddress1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasChangeShippingOption returns true if the method "PaymentRequestEvent.changeShippingOption" exists.
func (this PaymentRequestEvent) HasChangeShippingOption() bool {
	return js.True == bindings.HasPaymentRequestEventChangeShippingOption(
		this.Ref(),
	)
}

// ChangeShippingOptionFunc returns the method "PaymentRequestEvent.changeShippingOption".
func (this PaymentRequestEvent) ChangeShippingOptionFunc() (fn js.Func[func(shippingOption js.String) js.Promise[PaymentRequestDetailsUpdate]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventChangeShippingOptionFunc(
			this.Ref(),
		),
	)
}

// ChangeShippingOption calls the method "PaymentRequestEvent.changeShippingOption".
func (this PaymentRequestEvent) ChangeShippingOption(shippingOption js.String) (ret js.Promise[PaymentRequestDetailsUpdate]) {
	bindings.CallPaymentRequestEventChangeShippingOption(
		this.Ref(), js.Pointer(&ret),
		shippingOption.Ref(),
	)

	return
}

// TryChangeShippingOption calls the method "PaymentRequestEvent.changeShippingOption"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequestEvent) TryChangeShippingOption(shippingOption js.String) (ret js.Promise[PaymentRequestDetailsUpdate], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestEventChangeShippingOption(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		shippingOption.Ref(),
	)

	return
}

// HasRespondWith returns true if the method "PaymentRequestEvent.respondWith" exists.
func (this PaymentRequestEvent) HasRespondWith() bool {
	return js.True == bindings.HasPaymentRequestEventRespondWith(
		this.Ref(),
	)
}

// RespondWithFunc returns the method "PaymentRequestEvent.respondWith".
func (this PaymentRequestEvent) RespondWithFunc() (fn js.Func[func(handlerResponsePromise js.Promise[PaymentHandlerResponse])]) {
	return fn.FromRef(
		bindings.PaymentRequestEventRespondWithFunc(
			this.Ref(),
		),
	)
}

// RespondWith calls the method "PaymentRequestEvent.respondWith".
func (this PaymentRequestEvent) RespondWith(handlerResponsePromise js.Promise[PaymentHandlerResponse]) (ret js.Void) {
	bindings.CallPaymentRequestEventRespondWith(
		this.Ref(), js.Pointer(&ret),
		handlerResponsePromise.Ref(),
	)

	return
}

// TryRespondWith calls the method "PaymentRequestEvent.respondWith"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequestEvent) TryRespondWith(handlerResponsePromise js.Promise[PaymentHandlerResponse]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestEventRespondWith(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		handlerResponsePromise.Ref(),
	)

	return
}

type PaymentRequestUpdateEventInit struct {
	// Bubbles is "PaymentRequestUpdateEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PaymentRequestUpdateEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PaymentRequestUpdateEventInit.composed"
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

// FromRef calls UpdateFrom and returns a PaymentRequestUpdateEventInit with all fields set.
func (p PaymentRequestUpdateEventInit) FromRef(ref js.Ref) PaymentRequestUpdateEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentRequestUpdateEventInit in the application heap.
func (p PaymentRequestUpdateEventInit) New() js.Ref {
	return bindings.PaymentRequestUpdateEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentRequestUpdateEventInit) UpdateFrom(ref js.Ref) {
	bindings.PaymentRequestUpdateEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentRequestUpdateEventInit) Update(ref js.Ref) {
	bindings.PaymentRequestUpdateEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPaymentRequestUpdateEvent(typ js.String, eventInitDict PaymentRequestUpdateEventInit) (ret PaymentRequestUpdateEvent) {
	ret.ref = bindings.NewPaymentRequestUpdateEventByPaymentRequestUpdateEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewPaymentRequestUpdateEventByPaymentRequestUpdateEvent1(typ js.String) (ret PaymentRequestUpdateEvent) {
	ret.ref = bindings.NewPaymentRequestUpdateEventByPaymentRequestUpdateEvent1(
		typ.Ref())
	return
}

type PaymentRequestUpdateEvent struct {
	Event
}

func (this PaymentRequestUpdateEvent) Once() PaymentRequestUpdateEvent {
	this.Ref().Once()
	return this
}

func (this PaymentRequestUpdateEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this PaymentRequestUpdateEvent) FromRef(ref js.Ref) PaymentRequestUpdateEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this PaymentRequestUpdateEvent) Free() {
	this.Ref().Free()
}

// HasUpdateWith returns true if the method "PaymentRequestUpdateEvent.updateWith" exists.
func (this PaymentRequestUpdateEvent) HasUpdateWith() bool {
	return js.True == bindings.HasPaymentRequestUpdateEventUpdateWith(
		this.Ref(),
	)
}

// UpdateWithFunc returns the method "PaymentRequestUpdateEvent.updateWith".
func (this PaymentRequestUpdateEvent) UpdateWithFunc() (fn js.Func[func(detailsPromise js.Promise[PaymentDetailsUpdate])]) {
	return fn.FromRef(
		bindings.PaymentRequestUpdateEventUpdateWithFunc(
			this.Ref(),
		),
	)
}

// UpdateWith calls the method "PaymentRequestUpdateEvent.updateWith".
func (this PaymentRequestUpdateEvent) UpdateWith(detailsPromise js.Promise[PaymentDetailsUpdate]) (ret js.Void) {
	bindings.CallPaymentRequestUpdateEventUpdateWith(
		this.Ref(), js.Pointer(&ret),
		detailsPromise.Ref(),
	)

	return
}

// TryUpdateWith calls the method "PaymentRequestUpdateEvent.updateWith"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PaymentRequestUpdateEvent) TryUpdateWith(detailsPromise js.Promise[PaymentDetailsUpdate]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaymentRequestUpdateEventUpdateWith(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		detailsPromise.Ref(),
	)

	return
}

type Pbkdf2Params struct {
	// Salt is "Pbkdf2Params.salt"
	//
	// Required
	Salt BufferSource
	// Iterations is "Pbkdf2Params.iterations"
	//
	// Required
	Iterations uint32
	// Hash is "Pbkdf2Params.hash"
	//
	// Required
	Hash HashAlgorithmIdentifier
	// Name is "Pbkdf2Params.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Pbkdf2Params with all fields set.
func (p Pbkdf2Params) FromRef(ref js.Ref) Pbkdf2Params {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Pbkdf2Params in the application heap.
func (p Pbkdf2Params) New() js.Ref {
	return bindings.Pbkdf2ParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p Pbkdf2Params) UpdateFrom(ref js.Ref) {
	bindings.Pbkdf2ParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p Pbkdf2Params) Update(ref js.Ref) {
	bindings.Pbkdf2ParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PerformanceElementTiming struct {
	PerformanceEntry
}

func (this PerformanceElementTiming) Once() PerformanceElementTiming {
	this.Ref().Once()
	return this
}

func (this PerformanceElementTiming) Ref() js.Ref {
	return this.PerformanceEntry.Ref()
}

func (this PerformanceElementTiming) FromRef(ref js.Ref) PerformanceElementTiming {
	this.PerformanceEntry = this.PerformanceEntry.FromRef(ref)
	return this
}

func (this PerformanceElementTiming) Free() {
	this.Ref().Free()
}

// RenderTime returns the value of property "PerformanceElementTiming.renderTime".
//
// It returns ok=false if there is no such property.
func (this PerformanceElementTiming) RenderTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceElementTimingRenderTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LoadTime returns the value of property "PerformanceElementTiming.loadTime".
//
// It returns ok=false if there is no such property.
func (this PerformanceElementTiming) LoadTime() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceElementTimingLoadTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IntersectionRect returns the value of property "PerformanceElementTiming.intersectionRect".
//
// It returns ok=false if there is no such property.
func (this PerformanceElementTiming) IntersectionRect() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetPerformanceElementTimingIntersectionRect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Identifier returns the value of property "PerformanceElementTiming.identifier".
//
// It returns ok=false if there is no such property.
func (this PerformanceElementTiming) Identifier() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceElementTimingIdentifier(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NaturalWidth returns the value of property "PerformanceElementTiming.naturalWidth".
//
// It returns ok=false if there is no such property.
func (this PerformanceElementTiming) NaturalWidth() (ret uint32, ok bool) {
	ok = js.True == bindings.GetPerformanceElementTimingNaturalWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NaturalHeight returns the value of property "PerformanceElementTiming.naturalHeight".
//
// It returns ok=false if there is no such property.
func (this PerformanceElementTiming) NaturalHeight() (ret uint32, ok bool) {
	ok = js.True == bindings.GetPerformanceElementTimingNaturalHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Id returns the value of property "PerformanceElementTiming.id".
//
// It returns ok=false if there is no such property.
func (this PerformanceElementTiming) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceElementTimingId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Element returns the value of property "PerformanceElementTiming.element".
//
// It returns ok=false if there is no such property.
func (this PerformanceElementTiming) Element() (ret Element, ok bool) {
	ok = js.True == bindings.GetPerformanceElementTimingElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Url returns the value of property "PerformanceElementTiming.url".
//
// It returns ok=false if there is no such property.
func (this PerformanceElementTiming) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceElementTimingUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "PerformanceElementTiming.toJSON" exists.
func (this PerformanceElementTiming) HasToJSON() bool {
	return js.True == bindings.HasPerformanceElementTimingToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "PerformanceElementTiming.toJSON".
func (this PerformanceElementTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceElementTimingToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PerformanceElementTiming.toJSON".
func (this PerformanceElementTiming) ToJSON() (ret js.Object) {
	bindings.CallPerformanceElementTimingToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceElementTiming.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceElementTiming) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceElementTimingToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PerformanceEventTiming struct {
	PerformanceEntry
}

func (this PerformanceEventTiming) Once() PerformanceEventTiming {
	this.Ref().Once()
	return this
}

func (this PerformanceEventTiming) Ref() js.Ref {
	return this.PerformanceEntry.Ref()
}

func (this PerformanceEventTiming) FromRef(ref js.Ref) PerformanceEventTiming {
	this.PerformanceEntry = this.PerformanceEntry.FromRef(ref)
	return this
}

func (this PerformanceEventTiming) Free() {
	this.Ref().Free()
}

// ProcessingStart returns the value of property "PerformanceEventTiming.processingStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceEventTiming) ProcessingStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceEventTimingProcessingStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ProcessingEnd returns the value of property "PerformanceEventTiming.processingEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceEventTiming) ProcessingEnd() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceEventTimingProcessingEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Cancelable returns the value of property "PerformanceEventTiming.cancelable".
//
// It returns ok=false if there is no such property.
func (this PerformanceEventTiming) Cancelable() (ret bool, ok bool) {
	ok = js.True == bindings.GetPerformanceEventTimingCancelable(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Target returns the value of property "PerformanceEventTiming.target".
//
// It returns ok=false if there is no such property.
func (this PerformanceEventTiming) Target() (ret Node, ok bool) {
	ok = js.True == bindings.GetPerformanceEventTimingTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// InteractionId returns the value of property "PerformanceEventTiming.interactionId".
//
// It returns ok=false if there is no such property.
func (this PerformanceEventTiming) InteractionId() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceEventTimingInteractionId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "PerformanceEventTiming.toJSON" exists.
func (this PerformanceEventTiming) HasToJSON() bool {
	return js.True == bindings.HasPerformanceEventTimingToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "PerformanceEventTiming.toJSON".
func (this PerformanceEventTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceEventTimingToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PerformanceEventTiming.toJSON".
func (this PerformanceEventTiming) ToJSON() (ret js.Object) {
	bindings.CallPerformanceEventTimingToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceEventTiming.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceEventTiming) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceEventTimingToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type TaskAttributionTiming struct {
	PerformanceEntry
}

func (this TaskAttributionTiming) Once() TaskAttributionTiming {
	this.Ref().Once()
	return this
}

func (this TaskAttributionTiming) Ref() js.Ref {
	return this.PerformanceEntry.Ref()
}

func (this TaskAttributionTiming) FromRef(ref js.Ref) TaskAttributionTiming {
	this.PerformanceEntry = this.PerformanceEntry.FromRef(ref)
	return this
}

func (this TaskAttributionTiming) Free() {
	this.Ref().Free()
}

// ContainerType returns the value of property "TaskAttributionTiming.containerType".
//
// It returns ok=false if there is no such property.
func (this TaskAttributionTiming) ContainerType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTaskAttributionTimingContainerType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContainerSrc returns the value of property "TaskAttributionTiming.containerSrc".
//
// It returns ok=false if there is no such property.
func (this TaskAttributionTiming) ContainerSrc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTaskAttributionTimingContainerSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContainerId returns the value of property "TaskAttributionTiming.containerId".
//
// It returns ok=false if there is no such property.
func (this TaskAttributionTiming) ContainerId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTaskAttributionTimingContainerId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContainerName returns the value of property "TaskAttributionTiming.containerName".
//
// It returns ok=false if there is no such property.
func (this TaskAttributionTiming) ContainerName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetTaskAttributionTimingContainerName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "TaskAttributionTiming.toJSON" exists.
func (this TaskAttributionTiming) HasToJSON() bool {
	return js.True == bindings.HasTaskAttributionTimingToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "TaskAttributionTiming.toJSON".
func (this TaskAttributionTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.TaskAttributionTimingToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "TaskAttributionTiming.toJSON".
func (this TaskAttributionTiming) ToJSON() (ret js.Object) {
	bindings.CallTaskAttributionTimingToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "TaskAttributionTiming.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this TaskAttributionTiming) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTaskAttributionTimingToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PerformanceLongTaskTiming struct {
	PerformanceEntry
}

func (this PerformanceLongTaskTiming) Once() PerformanceLongTaskTiming {
	this.Ref().Once()
	return this
}

func (this PerformanceLongTaskTiming) Ref() js.Ref {
	return this.PerformanceEntry.Ref()
}

func (this PerformanceLongTaskTiming) FromRef(ref js.Ref) PerformanceLongTaskTiming {
	this.PerformanceEntry = this.PerformanceEntry.FromRef(ref)
	return this
}

func (this PerformanceLongTaskTiming) Free() {
	this.Ref().Free()
}

// Attribution returns the value of property "PerformanceLongTaskTiming.attribution".
//
// It returns ok=false if there is no such property.
func (this PerformanceLongTaskTiming) Attribution() (ret js.FrozenArray[TaskAttributionTiming], ok bool) {
	ok = js.True == bindings.GetPerformanceLongTaskTimingAttribution(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "PerformanceLongTaskTiming.toJSON" exists.
func (this PerformanceLongTaskTiming) HasToJSON() bool {
	return js.True == bindings.HasPerformanceLongTaskTimingToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "PerformanceLongTaskTiming.toJSON".
func (this PerformanceLongTaskTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceLongTaskTimingToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PerformanceLongTaskTiming.toJSON".
func (this PerformanceLongTaskTiming) ToJSON() (ret js.Object) {
	bindings.CallPerformanceLongTaskTimingToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceLongTaskTiming.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceLongTaskTiming) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceLongTaskTimingToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PerformanceNavigationTiming struct {
	PerformanceResourceTiming
}

func (this PerformanceNavigationTiming) Once() PerformanceNavigationTiming {
	this.Ref().Once()
	return this
}

func (this PerformanceNavigationTiming) Ref() js.Ref {
	return this.PerformanceResourceTiming.Ref()
}

func (this PerformanceNavigationTiming) FromRef(ref js.Ref) PerformanceNavigationTiming {
	this.PerformanceResourceTiming = this.PerformanceResourceTiming.FromRef(ref)
	return this
}

func (this PerformanceNavigationTiming) Free() {
	this.Ref().Free()
}

// UnloadEventStart returns the value of property "PerformanceNavigationTiming.unloadEventStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) UnloadEventStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingUnloadEventStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UnloadEventEnd returns the value of property "PerformanceNavigationTiming.unloadEventEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) UnloadEventEnd() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingUnloadEventEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomInteractive returns the value of property "PerformanceNavigationTiming.domInteractive".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) DomInteractive() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingDomInteractive(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomContentLoadedEventStart returns the value of property "PerformanceNavigationTiming.domContentLoadedEventStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) DomContentLoadedEventStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingDomContentLoadedEventStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomContentLoadedEventEnd returns the value of property "PerformanceNavigationTiming.domContentLoadedEventEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) DomContentLoadedEventEnd() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingDomContentLoadedEventEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomComplete returns the value of property "PerformanceNavigationTiming.domComplete".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) DomComplete() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingDomComplete(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LoadEventStart returns the value of property "PerformanceNavigationTiming.loadEventStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) LoadEventStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingLoadEventStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LoadEventEnd returns the value of property "PerformanceNavigationTiming.loadEventEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) LoadEventEnd() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingLoadEventEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "PerformanceNavigationTiming.type".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) Type() (ret NavigationTimingType, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RedirectCount returns the value of property "PerformanceNavigationTiming.redirectCount".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) RedirectCount() (ret uint16, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingRedirectCount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CriticalCHRestart returns the value of property "PerformanceNavigationTiming.criticalCHRestart".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) CriticalCHRestart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingCriticalCHRestart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ActivationStart returns the value of property "PerformanceNavigationTiming.activationStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceNavigationTiming) ActivationStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceNavigationTimingActivationStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "PerformanceNavigationTiming.toJSON" exists.
func (this PerformanceNavigationTiming) HasToJSON() bool {
	return js.True == bindings.HasPerformanceNavigationTimingToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "PerformanceNavigationTiming.toJSON".
func (this PerformanceNavigationTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceNavigationTimingToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PerformanceNavigationTiming.toJSON".
func (this PerformanceNavigationTiming) ToJSON() (ret js.Object) {
	bindings.CallPerformanceNavigationTimingToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceNavigationTiming.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceNavigationTiming) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceNavigationTimingToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PerformanceObserverCallbackFunc func(this js.Ref, entries PerformanceObserverEntryList, observer PerformanceObserver, options PerformanceObserverCallbackOptions) js.Ref

func (fn PerformanceObserverCallbackFunc) Register() js.Func[func(entries PerformanceObserverEntryList, observer PerformanceObserver, options PerformanceObserverCallbackOptions)] {
	return js.RegisterCallback[func(entries PerformanceObserverEntryList, observer PerformanceObserver, options PerformanceObserverCallbackOptions)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PerformanceObserverCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		PerformanceObserverEntryList{}.FromRef(args[0+1]),
		PerformanceObserver{}.FromRef(args[1+1]),
		PerformanceObserverCallbackOptions{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PerformanceObserverCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries PerformanceObserverEntryList, observer PerformanceObserver, options PerformanceObserverCallbackOptions) js.Ref
	Arg T
}

func (cb *PerformanceObserverCallback[T]) Register() js.Func[func(entries PerformanceObserverEntryList, observer PerformanceObserver, options PerformanceObserverCallbackOptions)] {
	return js.RegisterCallback[func(entries PerformanceObserverEntryList, observer PerformanceObserver, options PerformanceObserverCallbackOptions)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PerformanceObserverCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		PerformanceObserverEntryList{}.FromRef(args[0+1]),
		PerformanceObserver{}.FromRef(args[1+1]),
		PerformanceObserverCallbackOptions{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PerformanceObserverEntryList struct {
	ref js.Ref
}

func (this PerformanceObserverEntryList) Once() PerformanceObserverEntryList {
	this.Ref().Once()
	return this
}

func (this PerformanceObserverEntryList) Ref() js.Ref {
	return this.ref
}

func (this PerformanceObserverEntryList) FromRef(ref js.Ref) PerformanceObserverEntryList {
	this.ref = ref
	return this
}

func (this PerformanceObserverEntryList) Free() {
	this.Ref().Free()
}

// HasGetEntries returns true if the method "PerformanceObserverEntryList.getEntries" exists.
func (this PerformanceObserverEntryList) HasGetEntries() bool {
	return js.True == bindings.HasPerformanceObserverEntryListGetEntries(
		this.Ref(),
	)
}

// GetEntriesFunc returns the method "PerformanceObserverEntryList.getEntries".
func (this PerformanceObserverEntryList) GetEntriesFunc() (fn js.Func[func() PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceObserverEntryListGetEntriesFunc(
			this.Ref(),
		),
	)
}

// GetEntries calls the method "PerformanceObserverEntryList.getEntries".
func (this PerformanceObserverEntryList) GetEntries() (ret PerformanceEntryList) {
	bindings.CallPerformanceObserverEntryListGetEntries(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetEntries calls the method "PerformanceObserverEntryList.getEntries"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceObserverEntryList) TryGetEntries() (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceObserverEntryListGetEntries(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetEntriesByType returns true if the method "PerformanceObserverEntryList.getEntriesByType" exists.
func (this PerformanceObserverEntryList) HasGetEntriesByType() bool {
	return js.True == bindings.HasPerformanceObserverEntryListGetEntriesByType(
		this.Ref(),
	)
}

// GetEntriesByTypeFunc returns the method "PerformanceObserverEntryList.getEntriesByType".
func (this PerformanceObserverEntryList) GetEntriesByTypeFunc() (fn js.Func[func(typ js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceObserverEntryListGetEntriesByTypeFunc(
			this.Ref(),
		),
	)
}

// GetEntriesByType calls the method "PerformanceObserverEntryList.getEntriesByType".
func (this PerformanceObserverEntryList) GetEntriesByType(typ js.String) (ret PerformanceEntryList) {
	bindings.CallPerformanceObserverEntryListGetEntriesByType(
		this.Ref(), js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryGetEntriesByType calls the method "PerformanceObserverEntryList.getEntriesByType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceObserverEntryList) TryGetEntriesByType(typ js.String) (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceObserverEntryListGetEntriesByType(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

// HasGetEntriesByName returns true if the method "PerformanceObserverEntryList.getEntriesByName" exists.
func (this PerformanceObserverEntryList) HasGetEntriesByName() bool {
	return js.True == bindings.HasPerformanceObserverEntryListGetEntriesByName(
		this.Ref(),
	)
}

// GetEntriesByNameFunc returns the method "PerformanceObserverEntryList.getEntriesByName".
func (this PerformanceObserverEntryList) GetEntriesByNameFunc() (fn js.Func[func(name js.String, typ js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceObserverEntryListGetEntriesByNameFunc(
			this.Ref(),
		),
	)
}

// GetEntriesByName calls the method "PerformanceObserverEntryList.getEntriesByName".
func (this PerformanceObserverEntryList) GetEntriesByName(name js.String, typ js.String) (ret PerformanceEntryList) {
	bindings.CallPerformanceObserverEntryListGetEntriesByName(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		typ.Ref(),
	)

	return
}

// TryGetEntriesByName calls the method "PerformanceObserverEntryList.getEntriesByName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceObserverEntryList) TryGetEntriesByName(name js.String, typ js.String) (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceObserverEntryListGetEntriesByName(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		typ.Ref(),
	)

	return
}

// HasGetEntriesByName1 returns true if the method "PerformanceObserverEntryList.getEntriesByName" exists.
func (this PerformanceObserverEntryList) HasGetEntriesByName1() bool {
	return js.True == bindings.HasPerformanceObserverEntryListGetEntriesByName1(
		this.Ref(),
	)
}

// GetEntriesByName1Func returns the method "PerformanceObserverEntryList.getEntriesByName".
func (this PerformanceObserverEntryList) GetEntriesByName1Func() (fn js.Func[func(name js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceObserverEntryListGetEntriesByName1Func(
			this.Ref(),
		),
	)
}

// GetEntriesByName1 calls the method "PerformanceObserverEntryList.getEntriesByName".
func (this PerformanceObserverEntryList) GetEntriesByName1(name js.String) (ret PerformanceEntryList) {
	bindings.CallPerformanceObserverEntryListGetEntriesByName1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGetEntriesByName1 calls the method "PerformanceObserverEntryList.getEntriesByName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceObserverEntryList) TryGetEntriesByName1(name js.String) (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceObserverEntryListGetEntriesByName1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type PerformanceObserverCallbackOptions struct {
	// DroppedEntriesCount is "PerformanceObserverCallbackOptions.droppedEntriesCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_DroppedEntriesCount MUST be set to true to make this field effective.
	DroppedEntriesCount uint64

	FFI_USE_DroppedEntriesCount bool // for DroppedEntriesCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PerformanceObserverCallbackOptions with all fields set.
func (p PerformanceObserverCallbackOptions) FromRef(ref js.Ref) PerformanceObserverCallbackOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PerformanceObserverCallbackOptions in the application heap.
func (p PerformanceObserverCallbackOptions) New() js.Ref {
	return bindings.PerformanceObserverCallbackOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PerformanceObserverCallbackOptions) UpdateFrom(ref js.Ref) {
	bindings.PerformanceObserverCallbackOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PerformanceObserverCallbackOptions) Update(ref js.Ref) {
	bindings.PerformanceObserverCallbackOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PerformanceObserverInit struct {
	// EntryTypes is "PerformanceObserverInit.entryTypes"
	//
	// Optional
	EntryTypes js.Array[js.String]
	// Type is "PerformanceObserverInit.type"
	//
	// Optional
	Type js.String
	// Buffered is "PerformanceObserverInit.buffered"
	//
	// Optional
	//
	// NOTE: FFI_USE_Buffered MUST be set to true to make this field effective.
	Buffered bool
	// DurationThreshold is "PerformanceObserverInit.durationThreshold"
	//
	// Optional
	//
	// NOTE: FFI_USE_DurationThreshold MUST be set to true to make this field effective.
	DurationThreshold DOMHighResTimeStamp

	FFI_USE_Buffered          bool // for Buffered.
	FFI_USE_DurationThreshold bool // for DurationThreshold.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PerformanceObserverInit with all fields set.
func (p PerformanceObserverInit) FromRef(ref js.Ref) PerformanceObserverInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PerformanceObserverInit in the application heap.
func (p PerformanceObserverInit) New() js.Ref {
	return bindings.PerformanceObserverInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PerformanceObserverInit) UpdateFrom(ref js.Ref) {
	bindings.PerformanceObserverInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PerformanceObserverInit) Update(ref js.Ref) {
	bindings.PerformanceObserverInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPerformanceObserver(callback js.Func[func(entries PerformanceObserverEntryList, observer PerformanceObserver, options PerformanceObserverCallbackOptions)]) (ret PerformanceObserver) {
	ret.ref = bindings.NewPerformanceObserverByPerformanceObserver(
		callback.Ref())
	return
}

type PerformanceObserver struct {
	ref js.Ref
}

func (this PerformanceObserver) Once() PerformanceObserver {
	this.Ref().Once()
	return this
}

func (this PerformanceObserver) Ref() js.Ref {
	return this.ref
}

func (this PerformanceObserver) FromRef(ref js.Ref) PerformanceObserver {
	this.ref = ref
	return this
}

func (this PerformanceObserver) Free() {
	this.Ref().Free()
}

// SupportedEntryTypes returns the value of property "PerformanceObserver.supportedEntryTypes".
//
// It returns ok=false if there is no such property.
func (this PerformanceObserver) SupportedEntryTypes() (ret js.FrozenArray[js.String], ok bool) {
	ok = js.True == bindings.GetPerformanceObserverSupportedEntryTypes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasObserve returns true if the method "PerformanceObserver.observe" exists.
func (this PerformanceObserver) HasObserve() bool {
	return js.True == bindings.HasPerformanceObserverObserve(
		this.Ref(),
	)
}

// ObserveFunc returns the method "PerformanceObserver.observe".
func (this PerformanceObserver) ObserveFunc() (fn js.Func[func(options PerformanceObserverInit)]) {
	return fn.FromRef(
		bindings.PerformanceObserverObserveFunc(
			this.Ref(),
		),
	)
}

// Observe calls the method "PerformanceObserver.observe".
func (this PerformanceObserver) Observe(options PerformanceObserverInit) (ret js.Void) {
	bindings.CallPerformanceObserverObserve(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryObserve calls the method "PerformanceObserver.observe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceObserver) TryObserve(options PerformanceObserverInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceObserverObserve(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasObserve1 returns true if the method "PerformanceObserver.observe" exists.
func (this PerformanceObserver) HasObserve1() bool {
	return js.True == bindings.HasPerformanceObserverObserve1(
		this.Ref(),
	)
}

// Observe1Func returns the method "PerformanceObserver.observe".
func (this PerformanceObserver) Observe1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PerformanceObserverObserve1Func(
			this.Ref(),
		),
	)
}

// Observe1 calls the method "PerformanceObserver.observe".
func (this PerformanceObserver) Observe1() (ret js.Void) {
	bindings.CallPerformanceObserverObserve1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryObserve1 calls the method "PerformanceObserver.observe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceObserver) TryObserve1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceObserverObserve1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDisconnect returns true if the method "PerformanceObserver.disconnect" exists.
func (this PerformanceObserver) HasDisconnect() bool {
	return js.True == bindings.HasPerformanceObserverDisconnect(
		this.Ref(),
	)
}

// DisconnectFunc returns the method "PerformanceObserver.disconnect".
func (this PerformanceObserver) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PerformanceObserverDisconnectFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "PerformanceObserver.disconnect".
func (this PerformanceObserver) Disconnect() (ret js.Void) {
	bindings.CallPerformanceObserverDisconnect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDisconnect calls the method "PerformanceObserver.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceObserver) TryDisconnect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceObserverDisconnect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasTakeRecords returns true if the method "PerformanceObserver.takeRecords" exists.
func (this PerformanceObserver) HasTakeRecords() bool {
	return js.True == bindings.HasPerformanceObserverTakeRecords(
		this.Ref(),
	)
}

// TakeRecordsFunc returns the method "PerformanceObserver.takeRecords".
func (this PerformanceObserver) TakeRecordsFunc() (fn js.Func[func() PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceObserverTakeRecordsFunc(
			this.Ref(),
		),
	)
}

// TakeRecords calls the method "PerformanceObserver.takeRecords".
func (this PerformanceObserver) TakeRecords() (ret PerformanceEntryList) {
	bindings.CallPerformanceObserverTakeRecords(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTakeRecords calls the method "PerformanceObserver.takeRecords"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceObserver) TryTakeRecords() (ret PerformanceEntryList, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceObserverTakeRecords(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PerformancePaintTiming struct {
	PerformanceEntry
}

func (this PerformancePaintTiming) Once() PerformancePaintTiming {
	this.Ref().Once()
	return this
}

func (this PerformancePaintTiming) Ref() js.Ref {
	return this.PerformanceEntry.Ref()
}

func (this PerformancePaintTiming) FromRef(ref js.Ref) PerformancePaintTiming {
	this.PerformanceEntry = this.PerformanceEntry.FromRef(ref)
	return this
}

func (this PerformancePaintTiming) Free() {
	this.Ref().Free()
}

type RenderBlockingStatusType uint32

const (
	_ RenderBlockingStatusType = iota

	RenderBlockingStatusType_BLOCKING
	RenderBlockingStatusType_NON_BLOCKING
)

func (RenderBlockingStatusType) FromRef(str js.Ref) RenderBlockingStatusType {
	return RenderBlockingStatusType(bindings.ConstOfRenderBlockingStatusType(str))
}

func (x RenderBlockingStatusType) String() (string, bool) {
	switch x {
	case RenderBlockingStatusType_BLOCKING:
		return "blocking", true
	case RenderBlockingStatusType_NON_BLOCKING:
		return "non-blocking", true
	default:
		return "", false
	}
}

type PerformanceServerTiming struct {
	ref js.Ref
}

func (this PerformanceServerTiming) Once() PerformanceServerTiming {
	this.Ref().Once()
	return this
}

func (this PerformanceServerTiming) Ref() js.Ref {
	return this.ref
}

func (this PerformanceServerTiming) FromRef(ref js.Ref) PerformanceServerTiming {
	this.ref = ref
	return this
}

func (this PerformanceServerTiming) Free() {
	this.Ref().Free()
}

// Name returns the value of property "PerformanceServerTiming.name".
//
// It returns ok=false if there is no such property.
func (this PerformanceServerTiming) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceServerTimingName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Duration returns the value of property "PerformanceServerTiming.duration".
//
// It returns ok=false if there is no such property.
func (this PerformanceServerTiming) Duration() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceServerTimingDuration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Description returns the value of property "PerformanceServerTiming.description".
//
// It returns ok=false if there is no such property.
func (this PerformanceServerTiming) Description() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceServerTimingDescription(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "PerformanceServerTiming.toJSON" exists.
func (this PerformanceServerTiming) HasToJSON() bool {
	return js.True == bindings.HasPerformanceServerTimingToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "PerformanceServerTiming.toJSON".
func (this PerformanceServerTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceServerTimingToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PerformanceServerTiming.toJSON".
func (this PerformanceServerTiming) ToJSON() (ret js.Object) {
	bindings.CallPerformanceServerTimingToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceServerTiming.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceServerTiming) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceServerTimingToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PerformanceResourceTiming struct {
	PerformanceEntry
}

func (this PerformanceResourceTiming) Once() PerformanceResourceTiming {
	this.Ref().Once()
	return this
}

func (this PerformanceResourceTiming) Ref() js.Ref {
	return this.PerformanceEntry.Ref()
}

func (this PerformanceResourceTiming) FromRef(ref js.Ref) PerformanceResourceTiming {
	this.PerformanceEntry = this.PerformanceEntry.FromRef(ref)
	return this
}

func (this PerformanceResourceTiming) Free() {
	this.Ref().Free()
}

// InitiatorType returns the value of property "PerformanceResourceTiming.initiatorType".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) InitiatorType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingInitiatorType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DeliveryType returns the value of property "PerformanceResourceTiming.deliveryType".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) DeliveryType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingDeliveryType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NextHopProtocol returns the value of property "PerformanceResourceTiming.nextHopProtocol".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) NextHopProtocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingNextHopProtocol(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WorkerStart returns the value of property "PerformanceResourceTiming.workerStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) WorkerStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingWorkerStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RedirectStart returns the value of property "PerformanceResourceTiming.redirectStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) RedirectStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingRedirectStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RedirectEnd returns the value of property "PerformanceResourceTiming.redirectEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) RedirectEnd() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingRedirectEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FetchStart returns the value of property "PerformanceResourceTiming.fetchStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) FetchStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingFetchStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomainLookupStart returns the value of property "PerformanceResourceTiming.domainLookupStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) DomainLookupStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingDomainLookupStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DomainLookupEnd returns the value of property "PerformanceResourceTiming.domainLookupEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) DomainLookupEnd() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingDomainLookupEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ConnectStart returns the value of property "PerformanceResourceTiming.connectStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) ConnectStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingConnectStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ConnectEnd returns the value of property "PerformanceResourceTiming.connectEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) ConnectEnd() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingConnectEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SecureConnectionStart returns the value of property "PerformanceResourceTiming.secureConnectionStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) SecureConnectionStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingSecureConnectionStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RequestStart returns the value of property "PerformanceResourceTiming.requestStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) RequestStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingRequestStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FirstInterimResponseStart returns the value of property "PerformanceResourceTiming.firstInterimResponseStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) FirstInterimResponseStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingFirstInterimResponseStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ResponseStart returns the value of property "PerformanceResourceTiming.responseStart".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) ResponseStart() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingResponseStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ResponseEnd returns the value of property "PerformanceResourceTiming.responseEnd".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) ResponseEnd() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingResponseEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TransferSize returns the value of property "PerformanceResourceTiming.transferSize".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) TransferSize() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingTransferSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// EncodedBodySize returns the value of property "PerformanceResourceTiming.encodedBodySize".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) EncodedBodySize() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingEncodedBodySize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DecodedBodySize returns the value of property "PerformanceResourceTiming.decodedBodySize".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) DecodedBodySize() (ret uint64, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingDecodedBodySize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ResponseStatus returns the value of property "PerformanceResourceTiming.responseStatus".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) ResponseStatus() (ret uint16, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingResponseStatus(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RenderBlockingStatus returns the value of property "PerformanceResourceTiming.renderBlockingStatus".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) RenderBlockingStatus() (ret RenderBlockingStatusType, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingRenderBlockingStatus(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContentType returns the value of property "PerformanceResourceTiming.contentType".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) ContentType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingContentType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ServerTiming returns the value of property "PerformanceResourceTiming.serverTiming".
//
// It returns ok=false if there is no such property.
func (this PerformanceResourceTiming) ServerTiming() (ret js.FrozenArray[PerformanceServerTiming], ok bool) {
	ok = js.True == bindings.GetPerformanceResourceTimingServerTiming(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "PerformanceResourceTiming.toJSON" exists.
func (this PerformanceResourceTiming) HasToJSON() bool {
	return js.True == bindings.HasPerformanceResourceTimingToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "PerformanceResourceTiming.toJSON".
func (this PerformanceResourceTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceResourceTimingToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "PerformanceResourceTiming.toJSON".
func (this PerformanceResourceTiming) ToJSON() (ret js.Object) {
	bindings.CallPerformanceResourceTimingToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "PerformanceResourceTiming.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PerformanceResourceTiming) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPerformanceResourceTimingToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PeriodicSyncEventInit struct {
	// Tag is "PeriodicSyncEventInit.tag"
	//
	// Required
	Tag js.String
	// Bubbles is "PeriodicSyncEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PeriodicSyncEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PeriodicSyncEventInit.composed"
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

// FromRef calls UpdateFrom and returns a PeriodicSyncEventInit with all fields set.
func (p PeriodicSyncEventInit) FromRef(ref js.Ref) PeriodicSyncEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PeriodicSyncEventInit in the application heap.
func (p PeriodicSyncEventInit) New() js.Ref {
	return bindings.PeriodicSyncEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PeriodicSyncEventInit) UpdateFrom(ref js.Ref) {
	bindings.PeriodicSyncEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PeriodicSyncEventInit) Update(ref js.Ref) {
	bindings.PeriodicSyncEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPeriodicSyncEvent(typ js.String, init PeriodicSyncEventInit) (ret PeriodicSyncEvent) {
	ret.ref = bindings.NewPeriodicSyncEventByPeriodicSyncEvent(
		typ.Ref(),
		js.Pointer(&init))
	return
}

type PeriodicSyncEvent struct {
	ExtendableEvent
}

func (this PeriodicSyncEvent) Once() PeriodicSyncEvent {
	this.Ref().Once()
	return this
}

func (this PeriodicSyncEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this PeriodicSyncEvent) FromRef(ref js.Ref) PeriodicSyncEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this PeriodicSyncEvent) Free() {
	this.Ref().Free()
}

// Tag returns the value of property "PeriodicSyncEvent.tag".
//
// It returns ok=false if there is no such property.
func (this PeriodicSyncEvent) Tag() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPeriodicSyncEventTag(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type PermissionDescriptor struct {
	// Name is "PermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PermissionDescriptor with all fields set.
func (p PermissionDescriptor) FromRef(ref js.Ref) PermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PermissionDescriptor in the application heap.
func (p PermissionDescriptor) New() js.Ref {
	return bindings.PermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.PermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PermissionDescriptor) Update(ref js.Ref) {
	bindings.PermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PermissionSetParameters struct {
	// Descriptor is "PermissionSetParameters.descriptor"
	//
	// Required
	//
	// NOTE: Descriptor.FFI_USE MUST be set to true to get Descriptor used.
	Descriptor PermissionDescriptor
	// State is "PermissionSetParameters.state"
	//
	// Required
	State PermissionState

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PermissionSetParameters with all fields set.
func (p PermissionSetParameters) FromRef(ref js.Ref) PermissionSetParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PermissionSetParameters in the application heap.
func (p PermissionSetParameters) New() js.Ref {
	return bindings.PermissionSetParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PermissionSetParameters) UpdateFrom(ref js.Ref) {
	bindings.PermissionSetParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PermissionSetParameters) Update(ref js.Ref) {
	bindings.PermissionSetParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PermissionsPolicyViolationReportBody struct {
	ReportBody
}

func (this PermissionsPolicyViolationReportBody) Once() PermissionsPolicyViolationReportBody {
	this.Ref().Once()
	return this
}

func (this PermissionsPolicyViolationReportBody) Ref() js.Ref {
	return this.ReportBody.Ref()
}

func (this PermissionsPolicyViolationReportBody) FromRef(ref js.Ref) PermissionsPolicyViolationReportBody {
	this.ReportBody = this.ReportBody.FromRef(ref)
	return this
}

func (this PermissionsPolicyViolationReportBody) Free() {
	this.Ref().Free()
}

// FeatureId returns the value of property "PermissionsPolicyViolationReportBody.featureId".
//
// It returns ok=false if there is no such property.
func (this PermissionsPolicyViolationReportBody) FeatureId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPermissionsPolicyViolationReportBodyFeatureId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SourceFile returns the value of property "PermissionsPolicyViolationReportBody.sourceFile".
//
// It returns ok=false if there is no such property.
func (this PermissionsPolicyViolationReportBody) SourceFile() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPermissionsPolicyViolationReportBodySourceFile(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LineNumber returns the value of property "PermissionsPolicyViolationReportBody.lineNumber".
//
// It returns ok=false if there is no such property.
func (this PermissionsPolicyViolationReportBody) LineNumber() (ret int32, ok bool) {
	ok = js.True == bindings.GetPermissionsPolicyViolationReportBodyLineNumber(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ColumnNumber returns the value of property "PermissionsPolicyViolationReportBody.columnNumber".
//
// It returns ok=false if there is no such property.
func (this PermissionsPolicyViolationReportBody) ColumnNumber() (ret int32, ok bool) {
	ok = js.True == bindings.GetPermissionsPolicyViolationReportBodyColumnNumber(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Disposition returns the value of property "PermissionsPolicyViolationReportBody.disposition".
//
// It returns ok=false if there is no such property.
func (this PermissionsPolicyViolationReportBody) Disposition() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPermissionsPolicyViolationReportBodyDisposition(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type PictureInPictureEventInit struct {
	// PictureInPictureWindow is "PictureInPictureEventInit.pictureInPictureWindow"
	//
	// Required
	PictureInPictureWindow PictureInPictureWindow
	// Bubbles is "PictureInPictureEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PictureInPictureEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PictureInPictureEventInit.composed"
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

// FromRef calls UpdateFrom and returns a PictureInPictureEventInit with all fields set.
func (p PictureInPictureEventInit) FromRef(ref js.Ref) PictureInPictureEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PictureInPictureEventInit in the application heap.
func (p PictureInPictureEventInit) New() js.Ref {
	return bindings.PictureInPictureEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PictureInPictureEventInit) UpdateFrom(ref js.Ref) {
	bindings.PictureInPictureEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PictureInPictureEventInit) Update(ref js.Ref) {
	bindings.PictureInPictureEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPictureInPictureEvent(typ js.String, eventInitDict PictureInPictureEventInit) (ret PictureInPictureEvent) {
	ret.ref = bindings.NewPictureInPictureEventByPictureInPictureEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type PictureInPictureEvent struct {
	Event
}

func (this PictureInPictureEvent) Once() PictureInPictureEvent {
	this.Ref().Once()
	return this
}

func (this PictureInPictureEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this PictureInPictureEvent) FromRef(ref js.Ref) PictureInPictureEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this PictureInPictureEvent) Free() {
	this.Ref().Free()
}

// PictureInPictureWindow returns the value of property "PictureInPictureEvent.pictureInPictureWindow".
//
// It returns ok=false if there is no such property.
func (this PictureInPictureEvent) PictureInPictureWindow() (ret PictureInPictureWindow, ok bool) {
	ok = js.True == bindings.GetPictureInPictureEventPictureInPictureWindow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type PopStateEventInit struct {
	// State is "PopStateEventInit.state"
	//
	// Optional, defaults to null.
	State js.Any
	// HasUAVisualTransition is "PopStateEventInit.hasUAVisualTransition"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_HasUAVisualTransition MUST be set to true to make this field effective.
	HasUAVisualTransition bool
	// Bubbles is "PopStateEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PopStateEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PopStateEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_HasUAVisualTransition bool // for HasUAVisualTransition.
	FFI_USE_Bubbles               bool // for Bubbles.
	FFI_USE_Cancelable            bool // for Cancelable.
	FFI_USE_Composed              bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PopStateEventInit with all fields set.
func (p PopStateEventInit) FromRef(ref js.Ref) PopStateEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PopStateEventInit in the application heap.
func (p PopStateEventInit) New() js.Ref {
	return bindings.PopStateEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PopStateEventInit) UpdateFrom(ref js.Ref) {
	bindings.PopStateEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PopStateEventInit) Update(ref js.Ref) {
	bindings.PopStateEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPopStateEvent(typ js.String, eventInitDict PopStateEventInit) (ret PopStateEvent) {
	ret.ref = bindings.NewPopStateEventByPopStateEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewPopStateEventByPopStateEvent1(typ js.String) (ret PopStateEvent) {
	ret.ref = bindings.NewPopStateEventByPopStateEvent1(
		typ.Ref())
	return
}

type PopStateEvent struct {
	Event
}

func (this PopStateEvent) Once() PopStateEvent {
	this.Ref().Once()
	return this
}

func (this PopStateEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this PopStateEvent) FromRef(ref js.Ref) PopStateEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this PopStateEvent) Free() {
	this.Ref().Free()
}

// State returns the value of property "PopStateEvent.state".
//
// It returns ok=false if there is no such property.
func (this PopStateEvent) State() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetPopStateEventState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasUAVisualTransition returns the value of property "PopStateEvent.hasUAVisualTransition".
//
// It returns ok=false if there is no such property.
func (this PopStateEvent) HasUAVisualTransition() (ret bool, ok bool) {
	ok = js.True == bindings.GetPopStateEventHasUAVisualTransition(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type PortalActivateEventInit struct {
	// Data is "PortalActivateEventInit.data"
	//
	// Optional, defaults to null.
	Data js.Any
	// Bubbles is "PortalActivateEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PortalActivateEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PortalActivateEventInit.composed"
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

// FromRef calls UpdateFrom and returns a PortalActivateEventInit with all fields set.
func (p PortalActivateEventInit) FromRef(ref js.Ref) PortalActivateEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PortalActivateEventInit in the application heap.
func (p PortalActivateEventInit) New() js.Ref {
	return bindings.PortalActivateEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PortalActivateEventInit) UpdateFrom(ref js.Ref) {
	bindings.PortalActivateEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PortalActivateEventInit) Update(ref js.Ref) {
	bindings.PortalActivateEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPortalActivateEvent(typ js.String, eventInitDict PortalActivateEventInit) (ret PortalActivateEvent) {
	ret.ref = bindings.NewPortalActivateEventByPortalActivateEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewPortalActivateEventByPortalActivateEvent1(typ js.String) (ret PortalActivateEvent) {
	ret.ref = bindings.NewPortalActivateEventByPortalActivateEvent1(
		typ.Ref())
	return
}

type PortalActivateEvent struct {
	Event
}

func (this PortalActivateEvent) Once() PortalActivateEvent {
	this.Ref().Once()
	return this
}

func (this PortalActivateEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this PortalActivateEvent) FromRef(ref js.Ref) PortalActivateEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this PortalActivateEvent) Free() {
	this.Ref().Free()
}

// Data returns the value of property "PortalActivateEvent.data".
//
// It returns ok=false if there is no such property.
func (this PortalActivateEvent) Data() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetPortalActivateEventData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAdoptPredecessor returns true if the method "PortalActivateEvent.adoptPredecessor" exists.
func (this PortalActivateEvent) HasAdoptPredecessor() bool {
	return js.True == bindings.HasPortalActivateEventAdoptPredecessor(
		this.Ref(),
	)
}

// AdoptPredecessorFunc returns the method "PortalActivateEvent.adoptPredecessor".
func (this PortalActivateEvent) AdoptPredecessorFunc() (fn js.Func[func() HTMLPortalElement]) {
	return fn.FromRef(
		bindings.PortalActivateEventAdoptPredecessorFunc(
			this.Ref(),
		),
	)
}

// AdoptPredecessor calls the method "PortalActivateEvent.adoptPredecessor".
func (this PortalActivateEvent) AdoptPredecessor() (ret HTMLPortalElement) {
	bindings.CallPortalActivateEventAdoptPredecessor(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAdoptPredecessor calls the method "PortalActivateEvent.adoptPredecessor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PortalActivateEvent) TryAdoptPredecessor() (ret HTMLPortalElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPortalActivateEventAdoptPredecessor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PositionAlignSetting uint32

const (
	_ PositionAlignSetting = iota

	PositionAlignSetting_LINE_LEFT
	PositionAlignSetting_CENTER
	PositionAlignSetting_LINE_RIGHT
	PositionAlignSetting_AUTO
)

func (PositionAlignSetting) FromRef(str js.Ref) PositionAlignSetting {
	return PositionAlignSetting(bindings.ConstOfPositionAlignSetting(str))
}

func (x PositionAlignSetting) String() (string, bool) {
	switch x {
	case PositionAlignSetting_LINE_LEFT:
		return "line-left", true
	case PositionAlignSetting_CENTER:
		return "center", true
	case PositionAlignSetting_LINE_RIGHT:
		return "line-right", true
	case PositionAlignSetting_AUTO:
		return "auto", true
	default:
		return "", false
	}
}

type PresentationConnectionAvailableEventInit struct {
	// Connection is "PresentationConnectionAvailableEventInit.connection"
	//
	// Required
	Connection PresentationConnection
	// Bubbles is "PresentationConnectionAvailableEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PresentationConnectionAvailableEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PresentationConnectionAvailableEventInit.composed"
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

// FromRef calls UpdateFrom and returns a PresentationConnectionAvailableEventInit with all fields set.
func (p PresentationConnectionAvailableEventInit) FromRef(ref js.Ref) PresentationConnectionAvailableEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PresentationConnectionAvailableEventInit in the application heap.
func (p PresentationConnectionAvailableEventInit) New() js.Ref {
	return bindings.PresentationConnectionAvailableEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PresentationConnectionAvailableEventInit) UpdateFrom(ref js.Ref) {
	bindings.PresentationConnectionAvailableEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PresentationConnectionAvailableEventInit) Update(ref js.Ref) {
	bindings.PresentationConnectionAvailableEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

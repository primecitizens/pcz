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

type PaymentDetailsModifier struct {
	// SupportedMethods is "PaymentDetailsModifier.supportedMethods"
	//
	// Required
	SupportedMethods js.String
	// Total is "PaymentDetailsModifier.total"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 PaymentDetailsModifier PaymentDetailsModifier [// PaymentDetailsModifier] [0x1400093ee60 0x1400093ef00 0x1400093efa0 0x1400093f040] 0x14000920900 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 PaymentDetailsBase PaymentDetailsBase [// PaymentDetailsBase] [0x1400093edc0 0x1400093f0e0] 0x140009208d0 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 PaymentDetailsInit PaymentDetailsInit [// PaymentDetailsInit] [0x1400093f180 0x1400093f220 0x1400093f2c0 0x1400093f360] 0x14000920948 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 PaymentDetailsUpdate PaymentDetailsUpdate [// PaymentDetailsUpdate] [0x1400093f400 0x1400093f4a0 0x1400093f540 0x1400093f5e0] 0x14000920990 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 PaymentHandlerResponse PaymentHandlerResponse [// PaymentHandlerResponse] [0x1400093f680 0x1400093f720 0x1400093f7c0 0x1400093f860 0x1400093f900 0x1400093f9a0 0x1400093fa40] 0x140009209a8 {0 0}} in the application heap.
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
	Bubbles bool
	// Cancelable is "PaymentMethodChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PaymentMethodChangeEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 PaymentMethodChangeEventInit PaymentMethodChangeEventInit [// PaymentMethodChangeEventInit] [0x1400093fae0 0x1400093fb80 0x1400093fc20 0x1400093fd60 0x1400093fea0 0x1400093fcc0 0x1400093fe00 0x1400093ff40] 0x140009209d8 {0 0}} in the application heap.
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

func NewPaymentMethodChangeEvent(typ js.String, eventInitDict PaymentMethodChangeEventInit) PaymentMethodChangeEvent {
	return PaymentMethodChangeEvent{}.FromRef(
		bindings.NewPaymentMethodChangeEventByPaymentMethodChangeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewPaymentMethodChangeEventByPaymentMethodChangeEvent1(typ js.String) PaymentMethodChangeEvent {
	return PaymentMethodChangeEvent{}.FromRef(
		bindings.NewPaymentMethodChangeEventByPaymentMethodChangeEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this PaymentMethodChangeEvent) MethodName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPaymentMethodChangeEventMethodName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// MethodDetails returns the value of property "PaymentMethodChangeEvent.methodDetails".
//
// The returned bool will be false if there is no such property.
func (this PaymentMethodChangeEvent) MethodDetails() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetPaymentMethodChangeEventMethodDetails(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
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

// New creates a new {0x140004cc0e0 PaymentMethodData PaymentMethodData [// PaymentMethodData] [0x140009520a0 0x14000952140] 0x14000920a20 {0 0}} in the application heap.
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
	RequestPayerName bool
	// RequestBillingAddress is "PaymentOptions.requestBillingAddress"
	//
	// Optional, defaults to false.
	RequestBillingAddress bool
	// RequestPayerEmail is "PaymentOptions.requestPayerEmail"
	//
	// Optional, defaults to false.
	RequestPayerEmail bool
	// RequestPayerPhone is "PaymentOptions.requestPayerPhone"
	//
	// Optional, defaults to false.
	RequestPayerPhone bool
	// RequestShipping is "PaymentOptions.requestShipping"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 PaymentOptions PaymentOptions [// PaymentOptions] [0x140009521e0 0x14000952320 0x14000952460 0x140009525a0 0x140009526e0 0x14000952820 0x14000952280 0x140009523c0 0x14000952500 0x14000952640 0x14000952780] 0x14000920a68 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 PaymentValidationErrors PaymentValidationErrors [// PaymentValidationErrors] [0x140009528c0 0x14000952960] 0x14000920af8 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this PaymentResponse) RequestId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPaymentResponseRequestId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// MethodName returns the value of property "PaymentResponse.methodName".
//
// The returned bool will be false if there is no such property.
func (this PaymentResponse) MethodName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPaymentResponseMethodName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Details returns the value of property "PaymentResponse.details".
//
// The returned bool will be false if there is no such property.
func (this PaymentResponse) Details() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetPaymentResponseDetails(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// ToJSON calls the method "PaymentResponse.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PaymentResponse) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPaymentResponseToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PaymentResponse.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentResponse) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PaymentResponseToJSONFunc(
			this.Ref(),
		),
	)
}

// Complete calls the method "PaymentResponse.complete".
//
// The returned bool will be false if there is no such method.
func (this PaymentResponse) Complete(result PaymentComplete, details PaymentCompleteDetails) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallPaymentResponseComplete(
		this.Ref(), js.Pointer(&_ok),
		uint32(result),
		js.Pointer(&details),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// CompleteFunc returns the method "PaymentResponse.complete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentResponse) CompleteFunc() (fn js.Func[func(result PaymentComplete, details PaymentCompleteDetails) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentResponseCompleteFunc(
			this.Ref(),
		),
	)
}

// Complete1 calls the method "PaymentResponse.complete".
//
// The returned bool will be false if there is no such method.
func (this PaymentResponse) Complete1(result PaymentComplete) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallPaymentResponseComplete1(
		this.Ref(), js.Pointer(&_ok),
		uint32(result),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Complete1Func returns the method "PaymentResponse.complete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentResponse) Complete1Func() (fn js.Func[func(result PaymentComplete) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentResponseComplete1Func(
			this.Ref(),
		),
	)
}

// Complete2 calls the method "PaymentResponse.complete".
//
// The returned bool will be false if there is no such method.
func (this PaymentResponse) Complete2() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallPaymentResponseComplete2(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Complete2Func returns the method "PaymentResponse.complete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentResponse) Complete2Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentResponseComplete2Func(
			this.Ref(),
		),
	)
}

// Retry calls the method "PaymentResponse.retry".
//
// The returned bool will be false if there is no such method.
func (this PaymentResponse) Retry(errorFields PaymentValidationErrors) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallPaymentResponseRetry(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&errorFields),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RetryFunc returns the method "PaymentResponse.retry".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentResponse) RetryFunc() (fn js.Func[func(errorFields PaymentValidationErrors) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentResponseRetryFunc(
			this.Ref(),
		),
	)
}

// Retry1 calls the method "PaymentResponse.retry".
//
// The returned bool will be false if there is no such method.
func (this PaymentResponse) Retry1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallPaymentResponseRetry1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Retry1Func returns the method "PaymentResponse.retry".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentResponse) Retry1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentResponseRetry1Func(
			this.Ref(),
		),
	)
}

func NewPaymentRequest(methodData js.Array[PaymentMethodData], details PaymentDetailsInit) PaymentRequest {
	return PaymentRequest{}.FromRef(
		bindings.NewPaymentRequestByPaymentRequest(
			methodData.Ref(),
			js.Pointer(&details)),
	)
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
// The returned bool will be false if there is no such property.
func (this PaymentRequest) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPaymentRequestId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Show calls the method "PaymentRequest.show".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequest) Show(detailsPromise js.Promise[PaymentDetailsUpdate]) (js.Promise[PaymentResponse], bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestShow(
		this.Ref(), js.Pointer(&_ok),
		detailsPromise.Ref(),
	)

	return js.Promise[PaymentResponse]{}.FromRef(_ret), _ok
}

// ShowFunc returns the method "PaymentRequest.show".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequest) ShowFunc() (fn js.Func[func(detailsPromise js.Promise[PaymentDetailsUpdate]) js.Promise[PaymentResponse]]) {
	return fn.FromRef(
		bindings.PaymentRequestShowFunc(
			this.Ref(),
		),
	)
}

// Show1 calls the method "PaymentRequest.show".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequest) Show1() (js.Promise[PaymentResponse], bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestShow1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PaymentResponse]{}.FromRef(_ret), _ok
}

// Show1Func returns the method "PaymentRequest.show".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequest) Show1Func() (fn js.Func[func() js.Promise[PaymentResponse]]) {
	return fn.FromRef(
		bindings.PaymentRequestShow1Func(
			this.Ref(),
		),
	)
}

// Abort calls the method "PaymentRequest.abort".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequest) Abort() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestAbort(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// AbortFunc returns the method "PaymentRequest.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequest) AbortFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.PaymentRequestAbortFunc(
			this.Ref(),
		),
	)
}

// CanMakePayment calls the method "PaymentRequest.canMakePayment".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequest) CanMakePayment() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestCanMakePayment(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// CanMakePaymentFunc returns the method "PaymentRequest.canMakePayment".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequest) CanMakePaymentFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.PaymentRequestCanMakePaymentFunc(
			this.Ref(),
		),
	)
}

// IsSecurePaymentConfirmationAvailable calls the staticmethod "PaymentRequest.isSecurePaymentConfirmationAvailable".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequest) IsSecurePaymentConfirmationAvailable() (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestIsSecurePaymentConfirmationAvailable(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// IsSecurePaymentConfirmationAvailableFunc returns the staticmethod "PaymentRequest.isSecurePaymentConfirmationAvailable".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequest) IsSecurePaymentConfirmationAvailableFunc() (fn js.Func[func() js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.PaymentRequestIsSecurePaymentConfirmationAvailableFunc(
			this.Ref(),
		),
	)
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
	Amount PaymentCurrencyAmount
	// Selected is "PaymentShippingOption.selected"
	//
	// Optional, defaults to false.
	Selected bool

	FFI_USE_Selected bool // for Selected.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentShippingOption with all fields set.
func (p PaymentShippingOption) FromRef(ref js.Ref) PaymentShippingOption {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PaymentShippingOption PaymentShippingOption [// PaymentShippingOption] [0x14000952c80 0x14000952d20 0x14000952dc0 0x14000952e60 0x14000952f00] 0x14000920b40 {0 0}} in the application heap.
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
	ShippingAddressErrors AddressErrors

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentRequestDetailsUpdate with all fields set.
func (p PaymentRequestDetailsUpdate) FromRef(ref js.Ref) PaymentRequestDetailsUpdate {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PaymentRequestDetailsUpdate PaymentRequestDetailsUpdate [// PaymentRequestDetailsUpdate] [0x14000952aa0 0x14000952b40 0x14000952be0 0x14000952fa0 0x14000953040 0x140009530e0] 0x14000920b10 {0 0}} in the application heap.
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
	Total PaymentCurrencyAmount
	// Modifiers is "PaymentRequestEventInit.modifiers"
	//
	// Optional
	Modifiers js.Array[PaymentDetailsModifier]
	// PaymentOptions is "PaymentRequestEventInit.paymentOptions"
	//
	// Optional
	PaymentOptions PaymentOptions
	// ShippingOptions is "PaymentRequestEventInit.shippingOptions"
	//
	// Optional
	ShippingOptions js.Array[PaymentShippingOption]
	// Bubbles is "PaymentRequestEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "PaymentRequestEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PaymentRequestEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 PaymentRequestEventInit PaymentRequestEventInit [// PaymentRequestEventInit] [0x14000953180 0x14000953220 0x140009532c0 0x14000953360 0x14000953400 0x140009534a0 0x14000953540 0x140009535e0 0x14000953680 0x140009537c0 0x14000953900 0x14000953720 0x14000953860 0x140009539a0] 0x14000920b88 {0 0}} in the application heap.
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

func NewPaymentRequestEvent(typ js.String, eventInitDict PaymentRequestEventInit) PaymentRequestEvent {
	return PaymentRequestEvent{}.FromRef(
		bindings.NewPaymentRequestEventByPaymentRequestEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewPaymentRequestEventByPaymentRequestEvent1(typ js.String) PaymentRequestEvent {
	return PaymentRequestEvent{}.FromRef(
		bindings.NewPaymentRequestEventByPaymentRequestEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this PaymentRequestEvent) TopOrigin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPaymentRequestEventTopOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// PaymentRequestOrigin returns the value of property "PaymentRequestEvent.paymentRequestOrigin".
//
// The returned bool will be false if there is no such property.
func (this PaymentRequestEvent) PaymentRequestOrigin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPaymentRequestEventPaymentRequestOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// PaymentRequestId returns the value of property "PaymentRequestEvent.paymentRequestId".
//
// The returned bool will be false if there is no such property.
func (this PaymentRequestEvent) PaymentRequestId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPaymentRequestEventPaymentRequestId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// MethodData returns the value of property "PaymentRequestEvent.methodData".
//
// The returned bool will be false if there is no such property.
func (this PaymentRequestEvent) MethodData() (js.FrozenArray[PaymentMethodData], bool) {
	var _ok bool
	_ret := bindings.GetPaymentRequestEventMethodData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[PaymentMethodData]{}.FromRef(_ret), _ok
}

// Total returns the value of property "PaymentRequestEvent.total".
//
// The returned bool will be false if there is no such property.
func (this PaymentRequestEvent) Total() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetPaymentRequestEventTotal(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// Modifiers returns the value of property "PaymentRequestEvent.modifiers".
//
// The returned bool will be false if there is no such property.
func (this PaymentRequestEvent) Modifiers() (js.FrozenArray[PaymentDetailsModifier], bool) {
	var _ok bool
	_ret := bindings.GetPaymentRequestEventModifiers(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[PaymentDetailsModifier]{}.FromRef(_ret), _ok
}

// PaymentOptions returns the value of property "PaymentRequestEvent.paymentOptions".
//
// The returned bool will be false if there is no such property.
func (this PaymentRequestEvent) PaymentOptions() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetPaymentRequestEventPaymentOptions(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// ShippingOptions returns the value of property "PaymentRequestEvent.shippingOptions".
//
// The returned bool will be false if there is no such property.
func (this PaymentRequestEvent) ShippingOptions() (js.FrozenArray[PaymentShippingOption], bool) {
	var _ok bool
	_ret := bindings.GetPaymentRequestEventShippingOptions(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[PaymentShippingOption]{}.FromRef(_ret), _ok
}

// OpenWindow calls the method "PaymentRequestEvent.openWindow".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequestEvent) OpenWindow(url js.String) (js.Promise[WindowClient], bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestEventOpenWindow(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
	)

	return js.Promise[WindowClient]{}.FromRef(_ret), _ok
}

// OpenWindowFunc returns the method "PaymentRequestEvent.openWindow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequestEvent) OpenWindowFunc() (fn js.Func[func(url js.String) js.Promise[WindowClient]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventOpenWindowFunc(
			this.Ref(),
		),
	)
}

// ChangePaymentMethod calls the method "PaymentRequestEvent.changePaymentMethod".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequestEvent) ChangePaymentMethod(methodName js.String, methodDetails js.Object) (js.Promise[PaymentRequestDetailsUpdate], bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestEventChangePaymentMethod(
		this.Ref(), js.Pointer(&_ok),
		methodName.Ref(),
		methodDetails.Ref(),
	)

	return js.Promise[PaymentRequestDetailsUpdate]{}.FromRef(_ret), _ok
}

// ChangePaymentMethodFunc returns the method "PaymentRequestEvent.changePaymentMethod".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequestEvent) ChangePaymentMethodFunc() (fn js.Func[func(methodName js.String, methodDetails js.Object) js.Promise[PaymentRequestDetailsUpdate]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventChangePaymentMethodFunc(
			this.Ref(),
		),
	)
}

// ChangePaymentMethod1 calls the method "PaymentRequestEvent.changePaymentMethod".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequestEvent) ChangePaymentMethod1(methodName js.String) (js.Promise[PaymentRequestDetailsUpdate], bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestEventChangePaymentMethod1(
		this.Ref(), js.Pointer(&_ok),
		methodName.Ref(),
	)

	return js.Promise[PaymentRequestDetailsUpdate]{}.FromRef(_ret), _ok
}

// ChangePaymentMethod1Func returns the method "PaymentRequestEvent.changePaymentMethod".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequestEvent) ChangePaymentMethod1Func() (fn js.Func[func(methodName js.String) js.Promise[PaymentRequestDetailsUpdate]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventChangePaymentMethod1Func(
			this.Ref(),
		),
	)
}

// ChangeShippingAddress calls the method "PaymentRequestEvent.changeShippingAddress".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequestEvent) ChangeShippingAddress(shippingAddress AddressInit) (js.Promise[PaymentRequestDetailsUpdate], bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestEventChangeShippingAddress(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&shippingAddress),
	)

	return js.Promise[PaymentRequestDetailsUpdate]{}.FromRef(_ret), _ok
}

// ChangeShippingAddressFunc returns the method "PaymentRequestEvent.changeShippingAddress".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequestEvent) ChangeShippingAddressFunc() (fn js.Func[func(shippingAddress AddressInit) js.Promise[PaymentRequestDetailsUpdate]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventChangeShippingAddressFunc(
			this.Ref(),
		),
	)
}

// ChangeShippingAddress1 calls the method "PaymentRequestEvent.changeShippingAddress".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequestEvent) ChangeShippingAddress1() (js.Promise[PaymentRequestDetailsUpdate], bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestEventChangeShippingAddress1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[PaymentRequestDetailsUpdate]{}.FromRef(_ret), _ok
}

// ChangeShippingAddress1Func returns the method "PaymentRequestEvent.changeShippingAddress".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequestEvent) ChangeShippingAddress1Func() (fn js.Func[func() js.Promise[PaymentRequestDetailsUpdate]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventChangeShippingAddress1Func(
			this.Ref(),
		),
	)
}

// ChangeShippingOption calls the method "PaymentRequestEvent.changeShippingOption".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequestEvent) ChangeShippingOption(shippingOption js.String) (js.Promise[PaymentRequestDetailsUpdate], bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestEventChangeShippingOption(
		this.Ref(), js.Pointer(&_ok),
		shippingOption.Ref(),
	)

	return js.Promise[PaymentRequestDetailsUpdate]{}.FromRef(_ret), _ok
}

// ChangeShippingOptionFunc returns the method "PaymentRequestEvent.changeShippingOption".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequestEvent) ChangeShippingOptionFunc() (fn js.Func[func(shippingOption js.String) js.Promise[PaymentRequestDetailsUpdate]]) {
	return fn.FromRef(
		bindings.PaymentRequestEventChangeShippingOptionFunc(
			this.Ref(),
		),
	)
}

// RespondWith calls the method "PaymentRequestEvent.respondWith".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequestEvent) RespondWith(handlerResponsePromise js.Promise[PaymentHandlerResponse]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestEventRespondWith(
		this.Ref(), js.Pointer(&_ok),
		handlerResponsePromise.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RespondWithFunc returns the method "PaymentRequestEvent.respondWith".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequestEvent) RespondWithFunc() (fn js.Func[func(handlerResponsePromise js.Promise[PaymentHandlerResponse])]) {
	return fn.FromRef(
		bindings.PaymentRequestEventRespondWithFunc(
			this.Ref(),
		),
	)
}

type PaymentRequestUpdateEventInit struct {
	// Bubbles is "PaymentRequestUpdateEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "PaymentRequestUpdateEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PaymentRequestUpdateEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 PaymentRequestUpdateEventInit PaymentRequestUpdateEventInit [// PaymentRequestUpdateEventInit] [0x14000953ae0 0x14000953c20 0x14000953d60 0x14000953b80 0x14000953cc0 0x14000953e00] 0x14000920c48 {0 0}} in the application heap.
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

func NewPaymentRequestUpdateEvent(typ js.String, eventInitDict PaymentRequestUpdateEventInit) PaymentRequestUpdateEvent {
	return PaymentRequestUpdateEvent{}.FromRef(
		bindings.NewPaymentRequestUpdateEventByPaymentRequestUpdateEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewPaymentRequestUpdateEventByPaymentRequestUpdateEvent1(typ js.String) PaymentRequestUpdateEvent {
	return PaymentRequestUpdateEvent{}.FromRef(
		bindings.NewPaymentRequestUpdateEventByPaymentRequestUpdateEvent1(
			typ.Ref()),
	)
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

// UpdateWith calls the method "PaymentRequestUpdateEvent.updateWith".
//
// The returned bool will be false if there is no such method.
func (this PaymentRequestUpdateEvent) UpdateWith(detailsPromise js.Promise[PaymentDetailsUpdate]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaymentRequestUpdateEventUpdateWith(
		this.Ref(), js.Pointer(&_ok),
		detailsPromise.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateWithFunc returns the method "PaymentRequestUpdateEvent.updateWith".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaymentRequestUpdateEvent) UpdateWithFunc() (fn js.Func[func(detailsPromise js.Promise[PaymentDetailsUpdate])]) {
	return fn.FromRef(
		bindings.PaymentRequestUpdateEventUpdateWithFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 Pbkdf2Params Pbkdf2Params [// Pbkdf2Params] [0x14000953ea0 0x14000953f40 0x1400095a000 0x1400095a0a0] 0x14000920c78 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this PerformanceElementTiming) RenderTime() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceElementTimingRenderTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// LoadTime returns the value of property "PerformanceElementTiming.loadTime".
//
// The returned bool will be false if there is no such property.
func (this PerformanceElementTiming) LoadTime() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceElementTimingLoadTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// IntersectionRect returns the value of property "PerformanceElementTiming.intersectionRect".
//
// The returned bool will be false if there is no such property.
func (this PerformanceElementTiming) IntersectionRect() (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceElementTimingIntersectionRect(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

// Identifier returns the value of property "PerformanceElementTiming.identifier".
//
// The returned bool will be false if there is no such property.
func (this PerformanceElementTiming) Identifier() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceElementTimingIdentifier(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// NaturalWidth returns the value of property "PerformanceElementTiming.naturalWidth".
//
// The returned bool will be false if there is no such property.
func (this PerformanceElementTiming) NaturalWidth() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceElementTimingNaturalWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// NaturalHeight returns the value of property "PerformanceElementTiming.naturalHeight".
//
// The returned bool will be false if there is no such property.
func (this PerformanceElementTiming) NaturalHeight() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceElementTimingNaturalHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Id returns the value of property "PerformanceElementTiming.id".
//
// The returned bool will be false if there is no such property.
func (this PerformanceElementTiming) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceElementTimingId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Element returns the value of property "PerformanceElementTiming.element".
//
// The returned bool will be false if there is no such property.
func (this PerformanceElementTiming) Element() (Element, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceElementTimingElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// Url returns the value of property "PerformanceElementTiming.url".
//
// The returned bool will be false if there is no such property.
func (this PerformanceElementTiming) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceElementTimingUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ToJSON calls the method "PerformanceElementTiming.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PerformanceElementTiming) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceElementTimingToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PerformanceElementTiming.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceElementTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceElementTimingToJSONFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this PerformanceEventTiming) ProcessingStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceEventTimingProcessingStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// ProcessingEnd returns the value of property "PerformanceEventTiming.processingEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceEventTiming) ProcessingEnd() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceEventTimingProcessingEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// Cancelable returns the value of property "PerformanceEventTiming.cancelable".
//
// The returned bool will be false if there is no such property.
func (this PerformanceEventTiming) Cancelable() (bool, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceEventTimingCancelable(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Target returns the value of property "PerformanceEventTiming.target".
//
// The returned bool will be false if there is no such property.
func (this PerformanceEventTiming) Target() (Node, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceEventTimingTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return Node{}.FromRef(_ret), _ok
}

// InteractionId returns the value of property "PerformanceEventTiming.interactionId".
//
// The returned bool will be false if there is no such property.
func (this PerformanceEventTiming) InteractionId() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceEventTimingInteractionId(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// ToJSON calls the method "PerformanceEventTiming.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PerformanceEventTiming) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceEventTimingToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PerformanceEventTiming.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceEventTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceEventTimingToJSONFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this TaskAttributionTiming) ContainerType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTaskAttributionTimingContainerType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ContainerSrc returns the value of property "TaskAttributionTiming.containerSrc".
//
// The returned bool will be false if there is no such property.
func (this TaskAttributionTiming) ContainerSrc() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTaskAttributionTimingContainerSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ContainerId returns the value of property "TaskAttributionTiming.containerId".
//
// The returned bool will be false if there is no such property.
func (this TaskAttributionTiming) ContainerId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTaskAttributionTimingContainerId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ContainerName returns the value of property "TaskAttributionTiming.containerName".
//
// The returned bool will be false if there is no such property.
func (this TaskAttributionTiming) ContainerName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetTaskAttributionTimingContainerName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ToJSON calls the method "TaskAttributionTiming.toJSON".
//
// The returned bool will be false if there is no such method.
func (this TaskAttributionTiming) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallTaskAttributionTimingToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "TaskAttributionTiming.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this TaskAttributionTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.TaskAttributionTimingToJSONFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this PerformanceLongTaskTiming) Attribution() (js.FrozenArray[TaskAttributionTiming], bool) {
	var _ok bool
	_ret := bindings.GetPerformanceLongTaskTimingAttribution(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[TaskAttributionTiming]{}.FromRef(_ret), _ok
}

// ToJSON calls the method "PerformanceLongTaskTiming.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PerformanceLongTaskTiming) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceLongTaskTimingToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PerformanceLongTaskTiming.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceLongTaskTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceLongTaskTimingToJSONFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) UnloadEventStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingUnloadEventStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// UnloadEventEnd returns the value of property "PerformanceNavigationTiming.unloadEventEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) UnloadEventEnd() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingUnloadEventEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// DomInteractive returns the value of property "PerformanceNavigationTiming.domInteractive".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) DomInteractive() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingDomInteractive(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// DomContentLoadedEventStart returns the value of property "PerformanceNavigationTiming.domContentLoadedEventStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) DomContentLoadedEventStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingDomContentLoadedEventStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// DomContentLoadedEventEnd returns the value of property "PerformanceNavigationTiming.domContentLoadedEventEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) DomContentLoadedEventEnd() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingDomContentLoadedEventEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// DomComplete returns the value of property "PerformanceNavigationTiming.domComplete".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) DomComplete() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingDomComplete(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// LoadEventStart returns the value of property "PerformanceNavigationTiming.loadEventStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) LoadEventStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingLoadEventStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// LoadEventEnd returns the value of property "PerformanceNavigationTiming.loadEventEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) LoadEventEnd() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingLoadEventEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// Type returns the value of property "PerformanceNavigationTiming.type".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) Type() (NavigationTimingType, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingType(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigationTimingType(_ret), _ok
}

// RedirectCount returns the value of property "PerformanceNavigationTiming.redirectCount".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) RedirectCount() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingRedirectCount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// CriticalCHRestart returns the value of property "PerformanceNavigationTiming.criticalCHRestart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) CriticalCHRestart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingCriticalCHRestart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// ActivationStart returns the value of property "PerformanceNavigationTiming.activationStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceNavigationTiming) ActivationStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceNavigationTimingActivationStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// ToJSON calls the method "PerformanceNavigationTiming.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PerformanceNavigationTiming) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceNavigationTimingToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PerformanceNavigationTiming.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceNavigationTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceNavigationTimingToJSONFunc(
			this.Ref(),
		),
	)
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

// GetEntries calls the method "PerformanceObserverEntryList.getEntries".
//
// The returned bool will be false if there is no such method.
func (this PerformanceObserverEntryList) GetEntries() (PerformanceEntryList, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceObserverEntryListGetEntries(
		this.Ref(), js.Pointer(&_ok),
	)

	return PerformanceEntryList{}.FromRef(_ret), _ok
}

// GetEntriesFunc returns the method "PerformanceObserverEntryList.getEntries".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceObserverEntryList) GetEntriesFunc() (fn js.Func[func() PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceObserverEntryListGetEntriesFunc(
			this.Ref(),
		),
	)
}

// GetEntriesByType calls the method "PerformanceObserverEntryList.getEntriesByType".
//
// The returned bool will be false if there is no such method.
func (this PerformanceObserverEntryList) GetEntriesByType(typ js.String) (PerformanceEntryList, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceObserverEntryListGetEntriesByType(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	return PerformanceEntryList{}.FromRef(_ret), _ok
}

// GetEntriesByTypeFunc returns the method "PerformanceObserverEntryList.getEntriesByType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceObserverEntryList) GetEntriesByTypeFunc() (fn js.Func[func(typ js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceObserverEntryListGetEntriesByTypeFunc(
			this.Ref(),
		),
	)
}

// GetEntriesByName calls the method "PerformanceObserverEntryList.getEntriesByName".
//
// The returned bool will be false if there is no such method.
func (this PerformanceObserverEntryList) GetEntriesByName(name js.String, typ js.String) (PerformanceEntryList, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceObserverEntryListGetEntriesByName(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		typ.Ref(),
	)

	return PerformanceEntryList{}.FromRef(_ret), _ok
}

// GetEntriesByNameFunc returns the method "PerformanceObserverEntryList.getEntriesByName".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceObserverEntryList) GetEntriesByNameFunc() (fn js.Func[func(name js.String, typ js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceObserverEntryListGetEntriesByNameFunc(
			this.Ref(),
		),
	)
}

// GetEntriesByName1 calls the method "PerformanceObserverEntryList.getEntriesByName".
//
// The returned bool will be false if there is no such method.
func (this PerformanceObserverEntryList) GetEntriesByName1(name js.String) (PerformanceEntryList, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceObserverEntryListGetEntriesByName1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return PerformanceEntryList{}.FromRef(_ret), _ok
}

// GetEntriesByName1Func returns the method "PerformanceObserverEntryList.getEntriesByName".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceObserverEntryList) GetEntriesByName1Func() (fn js.Func[func(name js.String) PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceObserverEntryListGetEntriesByName1Func(
			this.Ref(),
		),
	)
}

type PerformanceObserverCallbackOptions struct {
	// DroppedEntriesCount is "PerformanceObserverCallbackOptions.droppedEntriesCount"
	//
	// Optional
	DroppedEntriesCount uint64

	FFI_USE_DroppedEntriesCount bool // for DroppedEntriesCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PerformanceObserverCallbackOptions with all fields set.
func (p PerformanceObserverCallbackOptions) FromRef(ref js.Ref) PerformanceObserverCallbackOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PerformanceObserverCallbackOptions PerformanceObserverCallbackOptions [// PerformanceObserverCallbackOptions] [0x1400095a3c0 0x1400095a460] 0x14000920d08 {0 0}} in the application heap.
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
	Buffered bool
	// DurationThreshold is "PerformanceObserverInit.durationThreshold"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 PerformanceObserverInit PerformanceObserverInit [// PerformanceObserverInit] [0x1400095a500 0x1400095a5a0 0x1400095a640 0x1400095a780 0x1400095a6e0 0x1400095a820] 0x14000920d38 {0 0}} in the application heap.
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

func NewPerformanceObserver(callback js.Func[func(entries PerformanceObserverEntryList, observer PerformanceObserver, options PerformanceObserverCallbackOptions)]) PerformanceObserver {
	return PerformanceObserver{}.FromRef(
		bindings.NewPerformanceObserverByPerformanceObserver(
			callback.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this PerformanceObserver) SupportedEntryTypes() (js.FrozenArray[js.String], bool) {
	var _ok bool
	_ret := bindings.GetPerformanceObserverSupportedEntryTypes(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[js.String]{}.FromRef(_ret), _ok
}

// Observe calls the method "PerformanceObserver.observe".
//
// The returned bool will be false if there is no such method.
func (this PerformanceObserver) Observe(options PerformanceObserverInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceObserverObserve(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ObserveFunc returns the method "PerformanceObserver.observe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceObserver) ObserveFunc() (fn js.Func[func(options PerformanceObserverInit)]) {
	return fn.FromRef(
		bindings.PerformanceObserverObserveFunc(
			this.Ref(),
		),
	)
}

// Observe1 calls the method "PerformanceObserver.observe".
//
// The returned bool will be false if there is no such method.
func (this PerformanceObserver) Observe1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceObserverObserve1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Observe1Func returns the method "PerformanceObserver.observe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceObserver) Observe1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PerformanceObserverObserve1Func(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "PerformanceObserver.disconnect".
//
// The returned bool will be false if there is no such method.
func (this PerformanceObserver) Disconnect() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceObserverDisconnect(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisconnectFunc returns the method "PerformanceObserver.disconnect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceObserver) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PerformanceObserverDisconnectFunc(
			this.Ref(),
		),
	)
}

// TakeRecords calls the method "PerformanceObserver.takeRecords".
//
// The returned bool will be false if there is no such method.
func (this PerformanceObserver) TakeRecords() (PerformanceEntryList, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceObserverTakeRecords(
		this.Ref(), js.Pointer(&_ok),
	)

	return PerformanceEntryList{}.FromRef(_ret), _ok
}

// TakeRecordsFunc returns the method "PerformanceObserver.takeRecords".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceObserver) TakeRecordsFunc() (fn js.Func[func() PerformanceEntryList]) {
	return fn.FromRef(
		bindings.PerformanceObserverTakeRecordsFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this PerformanceServerTiming) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceServerTimingName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Duration returns the value of property "PerformanceServerTiming.duration".
//
// The returned bool will be false if there is no such property.
func (this PerformanceServerTiming) Duration() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceServerTimingDuration(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// Description returns the value of property "PerformanceServerTiming.description".
//
// The returned bool will be false if there is no such property.
func (this PerformanceServerTiming) Description() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceServerTimingDescription(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ToJSON calls the method "PerformanceServerTiming.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PerformanceServerTiming) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceServerTimingToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PerformanceServerTiming.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceServerTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceServerTimingToJSONFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) InitiatorType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingInitiatorType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DeliveryType returns the value of property "PerformanceResourceTiming.deliveryType".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) DeliveryType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingDeliveryType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// NextHopProtocol returns the value of property "PerformanceResourceTiming.nextHopProtocol".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) NextHopProtocol() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingNextHopProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// WorkerStart returns the value of property "PerformanceResourceTiming.workerStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) WorkerStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingWorkerStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// RedirectStart returns the value of property "PerformanceResourceTiming.redirectStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) RedirectStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingRedirectStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// RedirectEnd returns the value of property "PerformanceResourceTiming.redirectEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) RedirectEnd() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingRedirectEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// FetchStart returns the value of property "PerformanceResourceTiming.fetchStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) FetchStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingFetchStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// DomainLookupStart returns the value of property "PerformanceResourceTiming.domainLookupStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) DomainLookupStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingDomainLookupStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// DomainLookupEnd returns the value of property "PerformanceResourceTiming.domainLookupEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) DomainLookupEnd() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingDomainLookupEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// ConnectStart returns the value of property "PerformanceResourceTiming.connectStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) ConnectStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingConnectStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// ConnectEnd returns the value of property "PerformanceResourceTiming.connectEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) ConnectEnd() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingConnectEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// SecureConnectionStart returns the value of property "PerformanceResourceTiming.secureConnectionStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) SecureConnectionStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingSecureConnectionStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// RequestStart returns the value of property "PerformanceResourceTiming.requestStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) RequestStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingRequestStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// FirstInterimResponseStart returns the value of property "PerformanceResourceTiming.firstInterimResponseStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) FirstInterimResponseStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingFirstInterimResponseStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// ResponseStart returns the value of property "PerformanceResourceTiming.responseStart".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) ResponseStart() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingResponseStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// ResponseEnd returns the value of property "PerformanceResourceTiming.responseEnd".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) ResponseEnd() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingResponseEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// TransferSize returns the value of property "PerformanceResourceTiming.transferSize".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) TransferSize() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingTransferSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// EncodedBodySize returns the value of property "PerformanceResourceTiming.encodedBodySize".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) EncodedBodySize() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingEncodedBodySize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// DecodedBodySize returns the value of property "PerformanceResourceTiming.decodedBodySize".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) DecodedBodySize() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingDecodedBodySize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// ResponseStatus returns the value of property "PerformanceResourceTiming.responseStatus".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) ResponseStatus() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingResponseStatus(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// RenderBlockingStatus returns the value of property "PerformanceResourceTiming.renderBlockingStatus".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) RenderBlockingStatus() (RenderBlockingStatusType, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingRenderBlockingStatus(
		this.Ref(), js.Pointer(&_ok),
	)
	return RenderBlockingStatusType(_ret), _ok
}

// ContentType returns the value of property "PerformanceResourceTiming.contentType".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) ContentType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingContentType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ServerTiming returns the value of property "PerformanceResourceTiming.serverTiming".
//
// The returned bool will be false if there is no such property.
func (this PerformanceResourceTiming) ServerTiming() (js.FrozenArray[PerformanceServerTiming], bool) {
	var _ok bool
	_ret := bindings.GetPerformanceResourceTimingServerTiming(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[PerformanceServerTiming]{}.FromRef(_ret), _ok
}

// ToJSON calls the method "PerformanceResourceTiming.toJSON".
//
// The returned bool will be false if there is no such method.
func (this PerformanceResourceTiming) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallPerformanceResourceTimingToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "PerformanceResourceTiming.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PerformanceResourceTiming) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.PerformanceResourceTimingToJSONFunc(
			this.Ref(),
		),
	)
}

type PeriodicSyncEventInit struct {
	// Tag is "PeriodicSyncEventInit.tag"
	//
	// Required
	Tag js.String
	// Bubbles is "PeriodicSyncEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "PeriodicSyncEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PeriodicSyncEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 PeriodicSyncEventInit PeriodicSyncEventInit [// PeriodicSyncEventInit] [0x1400095aa00 0x1400095aaa0 0x1400095abe0 0x1400095ad20 0x1400095ab40 0x1400095ac80 0x1400095adc0] 0x14000920d98 {0 0}} in the application heap.
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

func NewPeriodicSyncEvent(typ js.String, init PeriodicSyncEventInit) PeriodicSyncEvent {
	return PeriodicSyncEvent{}.FromRef(
		bindings.NewPeriodicSyncEventByPeriodicSyncEvent(
			typ.Ref(),
			js.Pointer(&init)),
	)
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
// The returned bool will be false if there is no such property.
func (this PeriodicSyncEvent) Tag() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPeriodicSyncEventTag(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
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

// New creates a new {0x140004cc0e0 PermissionDescriptor PermissionDescriptor [// PermissionDescriptor] [0x1400095ae60] 0x14000920df8 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 PermissionSetParameters PermissionSetParameters [// PermissionSetParameters] [0x1400095af00 0x1400095afa0] 0x14000920e28 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this PermissionsPolicyViolationReportBody) FeatureId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPermissionsPolicyViolationReportBodyFeatureId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SourceFile returns the value of property "PermissionsPolicyViolationReportBody.sourceFile".
//
// The returned bool will be false if there is no such property.
func (this PermissionsPolicyViolationReportBody) SourceFile() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPermissionsPolicyViolationReportBodySourceFile(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LineNumber returns the value of property "PermissionsPolicyViolationReportBody.lineNumber".
//
// The returned bool will be false if there is no such property.
func (this PermissionsPolicyViolationReportBody) LineNumber() (int32, bool) {
	var _ok bool
	_ret := bindings.GetPermissionsPolicyViolationReportBodyLineNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// ColumnNumber returns the value of property "PermissionsPolicyViolationReportBody.columnNumber".
//
// The returned bool will be false if there is no such property.
func (this PermissionsPolicyViolationReportBody) ColumnNumber() (int32, bool) {
	var _ok bool
	_ret := bindings.GetPermissionsPolicyViolationReportBodyColumnNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Disposition returns the value of property "PermissionsPolicyViolationReportBody.disposition".
//
// The returned bool will be false if there is no such property.
func (this PermissionsPolicyViolationReportBody) Disposition() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPermissionsPolicyViolationReportBodyDisposition(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type PictureInPictureEventInit struct {
	// PictureInPictureWindow is "PictureInPictureEventInit.pictureInPictureWindow"
	//
	// Required
	PictureInPictureWindow PictureInPictureWindow
	// Bubbles is "PictureInPictureEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "PictureInPictureEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PictureInPictureEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 PictureInPictureEventInit PictureInPictureEventInit [// PictureInPictureEventInit] [0x1400095b0e0 0x1400095b180 0x1400095b2c0 0x1400095b400 0x1400095b220 0x1400095b360 0x1400095b4a0] 0x14000920e40 {0 0}} in the application heap.
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

func NewPictureInPictureEvent(typ js.String, eventInitDict PictureInPictureEventInit) PictureInPictureEvent {
	return PictureInPictureEvent{}.FromRef(
		bindings.NewPictureInPictureEventByPictureInPictureEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
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
// The returned bool will be false if there is no such property.
func (this PictureInPictureEvent) PictureInPictureWindow() (PictureInPictureWindow, bool) {
	var _ok bool
	_ret := bindings.GetPictureInPictureEventPictureInPictureWindow(
		this.Ref(), js.Pointer(&_ok),
	)
	return PictureInPictureWindow{}.FromRef(_ret), _ok
}

type PopStateEventInit struct {
	// State is "PopStateEventInit.state"
	//
	// Optional, defaults to null.
	State js.Any
	// HasUAVisualTransition is "PopStateEventInit.hasUAVisualTransition"
	//
	// Optional, defaults to false.
	HasUAVisualTransition bool
	// Bubbles is "PopStateEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "PopStateEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PopStateEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 PopStateEventInit PopStateEventInit [// PopStateEventInit] [0x1400095b540 0x1400095b5e0 0x1400095b720 0x1400095b860 0x1400095b9a0 0x1400095b680 0x1400095b7c0 0x1400095b900 0x1400095ba40] 0x14000920ea0 {0 0}} in the application heap.
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

func NewPopStateEvent(typ js.String, eventInitDict PopStateEventInit) PopStateEvent {
	return PopStateEvent{}.FromRef(
		bindings.NewPopStateEventByPopStateEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewPopStateEventByPopStateEvent1(typ js.String) PopStateEvent {
	return PopStateEvent{}.FromRef(
		bindings.NewPopStateEventByPopStateEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this PopStateEvent) State() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetPopStateEventState(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// HasUAVisualTransition returns the value of property "PopStateEvent.hasUAVisualTransition".
//
// The returned bool will be false if there is no such property.
func (this PopStateEvent) HasUAVisualTransition() (bool, bool) {
	var _ok bool
	_ret := bindings.GetPopStateEventHasUAVisualTransition(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type PortalActivateEventInit struct {
	// Data is "PortalActivateEventInit.data"
	//
	// Optional, defaults to null.
	Data js.Any
	// Bubbles is "PortalActivateEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "PortalActivateEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PortalActivateEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 PortalActivateEventInit PortalActivateEventInit [// PortalActivateEventInit] [0x1400095bb80 0x1400095bc20 0x1400095bd60 0x1400095bea0 0x1400095bcc0 0x1400095be00 0x1400095bf40] 0x14000920ee8 {0 0}} in the application heap.
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

func NewPortalActivateEvent(typ js.String, eventInitDict PortalActivateEventInit) PortalActivateEvent {
	return PortalActivateEvent{}.FromRef(
		bindings.NewPortalActivateEventByPortalActivateEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewPortalActivateEventByPortalActivateEvent1(typ js.String) PortalActivateEvent {
	return PortalActivateEvent{}.FromRef(
		bindings.NewPortalActivateEventByPortalActivateEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this PortalActivateEvent) Data() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetPortalActivateEventData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// AdoptPredecessor calls the method "PortalActivateEvent.adoptPredecessor".
//
// The returned bool will be false if there is no such method.
func (this PortalActivateEvent) AdoptPredecessor() (HTMLPortalElement, bool) {
	var _ok bool
	_ret := bindings.CallPortalActivateEventAdoptPredecessor(
		this.Ref(), js.Pointer(&_ok),
	)

	return HTMLPortalElement{}.FromRef(_ret), _ok
}

// AdoptPredecessorFunc returns the method "PortalActivateEvent.adoptPredecessor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PortalActivateEvent) AdoptPredecessorFunc() (fn js.Func[func() HTMLPortalElement]) {
	return fn.FromRef(
		bindings.PortalActivateEventAdoptPredecessorFunc(
			this.Ref(),
		),
	)
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
	Bubbles bool
	// Cancelable is "PresentationConnectionAvailableEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PresentationConnectionAvailableEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 PresentationConnectionAvailableEventInit PresentationConnectionAvailableEventInit [// PresentationConnectionAvailableEventInit] [0x14000960000 0x140009600a0 0x140009601e0 0x14000960320 0x14000960140 0x14000960280 0x140009603c0] 0x14000920f48 {0 0}} in the application heap.
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

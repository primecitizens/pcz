// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web store_PaymentDetailsModifier
//go:noescape
func PaymentDetailsModifierJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentDetailsModifier
//go:noescape
func PaymentDetailsModifierJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaymentDetailsBase
//go:noescape
func PaymentDetailsBaseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentDetailsBase
//go:noescape
func PaymentDetailsBaseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaymentDetailsInit
//go:noescape
func PaymentDetailsInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentDetailsInit
//go:noescape
func PaymentDetailsInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaymentDetailsUpdate
//go:noescape
func PaymentDetailsUpdateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentDetailsUpdate
//go:noescape
func PaymentDetailsUpdateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaymentHandlerResponse
//go:noescape
func PaymentHandlerResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentHandlerResponse
//go:noescape
func PaymentHandlerResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaymentMethodChangeEventInit
//go:noescape
func PaymentMethodChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentMethodChangeEventInit
//go:noescape
func PaymentMethodChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PaymentMethodChangeEvent_PaymentMethodChangeEvent
//go:noescape
func NewPaymentMethodChangeEventByPaymentMethodChangeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PaymentMethodChangeEvent_PaymentMethodChangeEvent1
//go:noescape
func NewPaymentMethodChangeEventByPaymentMethodChangeEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_PaymentMethodChangeEvent_MethodName
//go:noescape
func GetPaymentMethodChangeEventMethodName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PaymentMethodChangeEvent_MethodDetails
//go:noescape
func GetPaymentMethodChangeEventMethodDetails(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PaymentMethodData
//go:noescape
func PaymentMethodDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentMethodData
//go:noescape
func PaymentMethodDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_PaymentShippingType
//go:noescape
func ConstOfPaymentShippingType(str js.Ref) uint32

//go:wasmimport plat/js/web store_PaymentOptions
//go:noescape
func PaymentOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentOptions
//go:noescape
func PaymentOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaymentValidationErrors
//go:noescape
func PaymentValidationErrorsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentValidationErrors
//go:noescape
func PaymentValidationErrorsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_PaymentResponse_RequestId
//go:noescape
func GetPaymentResponseRequestId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PaymentResponse_MethodName
//go:noescape
func GetPaymentResponseMethodName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PaymentResponse_Details
//go:noescape
func GetPaymentResponseDetails(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentResponse_ToJSON
//go:noescape
func HasPaymentResponseToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentResponse_ToJSON
//go:noescape
func PaymentResponseToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentResponse_ToJSON
//go:noescape
func CallPaymentResponseToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PaymentResponse_ToJSON
//go:noescape
func TryPaymentResponseToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentResponse_Complete
//go:noescape
func HasPaymentResponseComplete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentResponse_Complete
//go:noescape
func PaymentResponseCompleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentResponse_Complete
//go:noescape
func CallPaymentResponseComplete(
	this js.Ref, retPtr unsafe.Pointer,
	result uint32,
	details unsafe.Pointer)

//go:wasmimport plat/js/web try_PaymentResponse_Complete
//go:noescape
func TryPaymentResponseComplete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	result uint32,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentResponse_Complete1
//go:noescape
func HasPaymentResponseComplete1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentResponse_Complete1
//go:noescape
func PaymentResponseComplete1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentResponse_Complete1
//go:noescape
func CallPaymentResponseComplete1(
	this js.Ref, retPtr unsafe.Pointer,
	result uint32)

//go:wasmimport plat/js/web try_PaymentResponse_Complete1
//go:noescape
func TryPaymentResponseComplete1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	result uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentResponse_Complete2
//go:noescape
func HasPaymentResponseComplete2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentResponse_Complete2
//go:noescape
func PaymentResponseComplete2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentResponse_Complete2
//go:noescape
func CallPaymentResponseComplete2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PaymentResponse_Complete2
//go:noescape
func TryPaymentResponseComplete2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentResponse_Retry
//go:noescape
func HasPaymentResponseRetry(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentResponse_Retry
//go:noescape
func PaymentResponseRetryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentResponse_Retry
//go:noescape
func CallPaymentResponseRetry(
	this js.Ref, retPtr unsafe.Pointer,
	errorFields unsafe.Pointer)

//go:wasmimport plat/js/web try_PaymentResponse_Retry
//go:noescape
func TryPaymentResponseRetry(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	errorFields unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentResponse_Retry1
//go:noescape
func HasPaymentResponseRetry1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentResponse_Retry1
//go:noescape
func PaymentResponseRetry1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentResponse_Retry1
//go:noescape
func CallPaymentResponseRetry1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PaymentResponse_Retry1
//go:noescape
func TryPaymentResponseRetry1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_PaymentRequest_PaymentRequest
//go:noescape
func NewPaymentRequestByPaymentRequest(
	methodData js.Ref,
	details unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_PaymentRequest_Id
//go:noescape
func GetPaymentRequestId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequest_Show
//go:noescape
func HasPaymentRequestShow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequest_Show
//go:noescape
func PaymentRequestShowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequest_Show
//go:noescape
func CallPaymentRequestShow(
	this js.Ref, retPtr unsafe.Pointer,
	detailsPromise js.Ref)

//go:wasmimport plat/js/web try_PaymentRequest_Show
//go:noescape
func TryPaymentRequestShow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	detailsPromise js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequest_Show1
//go:noescape
func HasPaymentRequestShow1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequest_Show1
//go:noescape
func PaymentRequestShow1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequest_Show1
//go:noescape
func CallPaymentRequestShow1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PaymentRequest_Show1
//go:noescape
func TryPaymentRequestShow1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequest_Abort
//go:noescape
func HasPaymentRequestAbort(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequest_Abort
//go:noescape
func PaymentRequestAbortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequest_Abort
//go:noescape
func CallPaymentRequestAbort(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PaymentRequest_Abort
//go:noescape
func TryPaymentRequestAbort(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequest_CanMakePayment
//go:noescape
func HasPaymentRequestCanMakePayment(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequest_CanMakePayment
//go:noescape
func PaymentRequestCanMakePaymentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequest_CanMakePayment
//go:noescape
func CallPaymentRequestCanMakePayment(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PaymentRequest_CanMakePayment
//go:noescape
func TryPaymentRequestCanMakePayment(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequest_IsSecurePaymentConfirmationAvailable
//go:noescape
func HasPaymentRequestIsSecurePaymentConfirmationAvailable(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequest_IsSecurePaymentConfirmationAvailable
//go:noescape
func PaymentRequestIsSecurePaymentConfirmationAvailableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequest_IsSecurePaymentConfirmationAvailable
//go:noescape
func CallPaymentRequestIsSecurePaymentConfirmationAvailable(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PaymentRequest_IsSecurePaymentConfirmationAvailable
//go:noescape
func TryPaymentRequestIsSecurePaymentConfirmationAvailable(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PaymentShippingOption
//go:noescape
func PaymentShippingOptionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentShippingOption
//go:noescape
func PaymentShippingOptionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaymentRequestDetailsUpdate
//go:noescape
func PaymentRequestDetailsUpdateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentRequestDetailsUpdate
//go:noescape
func PaymentRequestDetailsUpdateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PaymentRequestEventInit
//go:noescape
func PaymentRequestEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentRequestEventInit
//go:noescape
func PaymentRequestEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PaymentRequestEvent_PaymentRequestEvent
//go:noescape
func NewPaymentRequestEventByPaymentRequestEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PaymentRequestEvent_PaymentRequestEvent1
//go:noescape
func NewPaymentRequestEventByPaymentRequestEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_PaymentRequestEvent_TopOrigin
//go:noescape
func GetPaymentRequestEventTopOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PaymentRequestEvent_PaymentRequestOrigin
//go:noescape
func GetPaymentRequestEventPaymentRequestOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PaymentRequestEvent_PaymentRequestId
//go:noescape
func GetPaymentRequestEventPaymentRequestId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PaymentRequestEvent_MethodData
//go:noescape
func GetPaymentRequestEventMethodData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PaymentRequestEvent_Total
//go:noescape
func GetPaymentRequestEventTotal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PaymentRequestEvent_Modifiers
//go:noescape
func GetPaymentRequestEventModifiers(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PaymentRequestEvent_PaymentOptions
//go:noescape
func GetPaymentRequestEventPaymentOptions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PaymentRequestEvent_ShippingOptions
//go:noescape
func GetPaymentRequestEventShippingOptions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequestEvent_OpenWindow
//go:noescape
func HasPaymentRequestEventOpenWindow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequestEvent_OpenWindow
//go:noescape
func PaymentRequestEventOpenWindowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequestEvent_OpenWindow
//go:noescape
func CallPaymentRequestEventOpenWindow(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_PaymentRequestEvent_OpenWindow
//go:noescape
func TryPaymentRequestEventOpenWindow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequestEvent_ChangePaymentMethod
//go:noescape
func HasPaymentRequestEventChangePaymentMethod(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequestEvent_ChangePaymentMethod
//go:noescape
func PaymentRequestEventChangePaymentMethodFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequestEvent_ChangePaymentMethod
//go:noescape
func CallPaymentRequestEventChangePaymentMethod(
	this js.Ref, retPtr unsafe.Pointer,
	methodName js.Ref,
	methodDetails js.Ref)

//go:wasmimport plat/js/web try_PaymentRequestEvent_ChangePaymentMethod
//go:noescape
func TryPaymentRequestEventChangePaymentMethod(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	methodName js.Ref,
	methodDetails js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequestEvent_ChangePaymentMethod1
//go:noescape
func HasPaymentRequestEventChangePaymentMethod1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequestEvent_ChangePaymentMethod1
//go:noescape
func PaymentRequestEventChangePaymentMethod1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequestEvent_ChangePaymentMethod1
//go:noescape
func CallPaymentRequestEventChangePaymentMethod1(
	this js.Ref, retPtr unsafe.Pointer,
	methodName js.Ref)

//go:wasmimport plat/js/web try_PaymentRequestEvent_ChangePaymentMethod1
//go:noescape
func TryPaymentRequestEventChangePaymentMethod1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	methodName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequestEvent_ChangeShippingAddress
//go:noescape
func HasPaymentRequestEventChangeShippingAddress(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequestEvent_ChangeShippingAddress
//go:noescape
func PaymentRequestEventChangeShippingAddressFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequestEvent_ChangeShippingAddress
//go:noescape
func CallPaymentRequestEventChangeShippingAddress(
	this js.Ref, retPtr unsafe.Pointer,
	shippingAddress unsafe.Pointer)

//go:wasmimport plat/js/web try_PaymentRequestEvent_ChangeShippingAddress
//go:noescape
func TryPaymentRequestEventChangeShippingAddress(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	shippingAddress unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequestEvent_ChangeShippingAddress1
//go:noescape
func HasPaymentRequestEventChangeShippingAddress1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequestEvent_ChangeShippingAddress1
//go:noescape
func PaymentRequestEventChangeShippingAddress1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequestEvent_ChangeShippingAddress1
//go:noescape
func CallPaymentRequestEventChangeShippingAddress1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PaymentRequestEvent_ChangeShippingAddress1
//go:noescape
func TryPaymentRequestEventChangeShippingAddress1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequestEvent_ChangeShippingOption
//go:noescape
func HasPaymentRequestEventChangeShippingOption(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequestEvent_ChangeShippingOption
//go:noescape
func PaymentRequestEventChangeShippingOptionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequestEvent_ChangeShippingOption
//go:noescape
func CallPaymentRequestEventChangeShippingOption(
	this js.Ref, retPtr unsafe.Pointer,
	shippingOption js.Ref)

//go:wasmimport plat/js/web try_PaymentRequestEvent_ChangeShippingOption
//go:noescape
func TryPaymentRequestEventChangeShippingOption(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	shippingOption js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_PaymentRequestEvent_RespondWith
//go:noescape
func HasPaymentRequestEventRespondWith(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequestEvent_RespondWith
//go:noescape
func PaymentRequestEventRespondWithFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequestEvent_RespondWith
//go:noescape
func CallPaymentRequestEventRespondWith(
	this js.Ref, retPtr unsafe.Pointer,
	handlerResponsePromise js.Ref)

//go:wasmimport plat/js/web try_PaymentRequestEvent_RespondWith
//go:noescape
func TryPaymentRequestEventRespondWith(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	handlerResponsePromise js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_PaymentRequestUpdateEventInit
//go:noescape
func PaymentRequestUpdateEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PaymentRequestUpdateEventInit
//go:noescape
func PaymentRequestUpdateEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PaymentRequestUpdateEvent_PaymentRequestUpdateEvent
//go:noescape
func NewPaymentRequestUpdateEventByPaymentRequestUpdateEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PaymentRequestUpdateEvent_PaymentRequestUpdateEvent1
//go:noescape
func NewPaymentRequestUpdateEventByPaymentRequestUpdateEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web has_PaymentRequestUpdateEvent_UpdateWith
//go:noescape
func HasPaymentRequestUpdateEventUpdateWith(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PaymentRequestUpdateEvent_UpdateWith
//go:noescape
func PaymentRequestUpdateEventUpdateWithFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PaymentRequestUpdateEvent_UpdateWith
//go:noescape
func CallPaymentRequestUpdateEventUpdateWith(
	this js.Ref, retPtr unsafe.Pointer,
	detailsPromise js.Ref)

//go:wasmimport plat/js/web try_PaymentRequestUpdateEvent_UpdateWith
//go:noescape
func TryPaymentRequestUpdateEventUpdateWith(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	detailsPromise js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_Pbkdf2Params
//go:noescape
func Pbkdf2ParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_Pbkdf2Params
//go:noescape
func Pbkdf2ParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_PerformanceElementTiming_RenderTime
//go:noescape
func GetPerformanceElementTimingRenderTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceElementTiming_LoadTime
//go:noescape
func GetPerformanceElementTimingLoadTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceElementTiming_IntersectionRect
//go:noescape
func GetPerformanceElementTimingIntersectionRect(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceElementTiming_Identifier
//go:noescape
func GetPerformanceElementTimingIdentifier(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceElementTiming_NaturalWidth
//go:noescape
func GetPerformanceElementTimingNaturalWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceElementTiming_NaturalHeight
//go:noescape
func GetPerformanceElementTimingNaturalHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceElementTiming_Id
//go:noescape
func GetPerformanceElementTimingId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceElementTiming_Element
//go:noescape
func GetPerformanceElementTimingElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceElementTiming_Url
//go:noescape
func GetPerformanceElementTimingUrl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceElementTiming_ToJSON
//go:noescape
func HasPerformanceElementTimingToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceElementTiming_ToJSON
//go:noescape
func PerformanceElementTimingToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceElementTiming_ToJSON
//go:noescape
func CallPerformanceElementTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceElementTiming_ToJSON
//go:noescape
func TryPerformanceElementTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceEventTiming_ProcessingStart
//go:noescape
func GetPerformanceEventTimingProcessingStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceEventTiming_ProcessingEnd
//go:noescape
func GetPerformanceEventTimingProcessingEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceEventTiming_Cancelable
//go:noescape
func GetPerformanceEventTimingCancelable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceEventTiming_Target
//go:noescape
func GetPerformanceEventTimingTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceEventTiming_InteractionId
//go:noescape
func GetPerformanceEventTimingInteractionId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceEventTiming_ToJSON
//go:noescape
func HasPerformanceEventTimingToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceEventTiming_ToJSON
//go:noescape
func PerformanceEventTimingToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceEventTiming_ToJSON
//go:noescape
func CallPerformanceEventTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceEventTiming_ToJSON
//go:noescape
func TryPerformanceEventTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TaskAttributionTiming_ContainerType
//go:noescape
func GetTaskAttributionTimingContainerType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TaskAttributionTiming_ContainerSrc
//go:noescape
func GetTaskAttributionTimingContainerSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TaskAttributionTiming_ContainerId
//go:noescape
func GetTaskAttributionTimingContainerId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_TaskAttributionTiming_ContainerName
//go:noescape
func GetTaskAttributionTimingContainerName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_TaskAttributionTiming_ToJSON
//go:noescape
func HasTaskAttributionTimingToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_TaskAttributionTiming_ToJSON
//go:noescape
func TaskAttributionTimingToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TaskAttributionTiming_ToJSON
//go:noescape
func CallTaskAttributionTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_TaskAttributionTiming_ToJSON
//go:noescape
func TryTaskAttributionTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceLongTaskTiming_Attribution
//go:noescape
func GetPerformanceLongTaskTimingAttribution(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceLongTaskTiming_ToJSON
//go:noescape
func HasPerformanceLongTaskTimingToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceLongTaskTiming_ToJSON
//go:noescape
func PerformanceLongTaskTimingToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceLongTaskTiming_ToJSON
//go:noescape
func CallPerformanceLongTaskTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceLongTaskTiming_ToJSON
//go:noescape
func TryPerformanceLongTaskTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_UnloadEventStart
//go:noescape
func GetPerformanceNavigationTimingUnloadEventStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_UnloadEventEnd
//go:noescape
func GetPerformanceNavigationTimingUnloadEventEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_DomInteractive
//go:noescape
func GetPerformanceNavigationTimingDomInteractive(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_DomContentLoadedEventStart
//go:noescape
func GetPerformanceNavigationTimingDomContentLoadedEventStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_DomContentLoadedEventEnd
//go:noescape
func GetPerformanceNavigationTimingDomContentLoadedEventEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_DomComplete
//go:noescape
func GetPerformanceNavigationTimingDomComplete(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_LoadEventStart
//go:noescape
func GetPerformanceNavigationTimingLoadEventStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_LoadEventEnd
//go:noescape
func GetPerformanceNavigationTimingLoadEventEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_Type
//go:noescape
func GetPerformanceNavigationTimingType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_RedirectCount
//go:noescape
func GetPerformanceNavigationTimingRedirectCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_CriticalCHRestart
//go:noescape
func GetPerformanceNavigationTimingCriticalCHRestart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceNavigationTiming_ActivationStart
//go:noescape
func GetPerformanceNavigationTimingActivationStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceNavigationTiming_ToJSON
//go:noescape
func HasPerformanceNavigationTimingToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceNavigationTiming_ToJSON
//go:noescape
func PerformanceNavigationTimingToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceNavigationTiming_ToJSON
//go:noescape
func CallPerformanceNavigationTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceNavigationTiming_ToJSON
//go:noescape
func TryPerformanceNavigationTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceObserverEntryList_GetEntries
//go:noescape
func HasPerformanceObserverEntryListGetEntries(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceObserverEntryList_GetEntries
//go:noescape
func PerformanceObserverEntryListGetEntriesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceObserverEntryList_GetEntries
//go:noescape
func CallPerformanceObserverEntryListGetEntries(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceObserverEntryList_GetEntries
//go:noescape
func TryPerformanceObserverEntryListGetEntries(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceObserverEntryList_GetEntriesByType
//go:noescape
func HasPerformanceObserverEntryListGetEntriesByType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceObserverEntryList_GetEntriesByType
//go:noescape
func PerformanceObserverEntryListGetEntriesByTypeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceObserverEntryList_GetEntriesByType
//go:noescape
func CallPerformanceObserverEntryListGetEntriesByType(
	this js.Ref, retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/web try_PerformanceObserverEntryList_GetEntriesByType
//go:noescape
func TryPerformanceObserverEntryListGetEntriesByType(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceObserverEntryList_GetEntriesByName
//go:noescape
func HasPerformanceObserverEntryListGetEntriesByName(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceObserverEntryList_GetEntriesByName
//go:noescape
func PerformanceObserverEntryListGetEntriesByNameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceObserverEntryList_GetEntriesByName
//go:noescape
func CallPerformanceObserverEntryListGetEntriesByName(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	typ js.Ref)

//go:wasmimport plat/js/web try_PerformanceObserverEntryList_GetEntriesByName
//go:noescape
func TryPerformanceObserverEntryListGetEntriesByName(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceObserverEntryList_GetEntriesByName1
//go:noescape
func HasPerformanceObserverEntryListGetEntriesByName1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceObserverEntryList_GetEntriesByName1
//go:noescape
func PerformanceObserverEntryListGetEntriesByName1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceObserverEntryList_GetEntriesByName1
//go:noescape
func CallPerformanceObserverEntryListGetEntriesByName1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_PerformanceObserverEntryList_GetEntriesByName1
//go:noescape
func TryPerformanceObserverEntryListGetEntriesByName1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_PerformanceObserverCallbackOptions
//go:noescape
func PerformanceObserverCallbackOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PerformanceObserverCallbackOptions
//go:noescape
func PerformanceObserverCallbackOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PerformanceObserverInit
//go:noescape
func PerformanceObserverInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PerformanceObserverInit
//go:noescape
func PerformanceObserverInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PerformanceObserver_PerformanceObserver
//go:noescape
func NewPerformanceObserverByPerformanceObserver(
	callback js.Ref) js.Ref

//go:wasmimport plat/js/web get_PerformanceObserver_SupportedEntryTypes
//go:noescape
func GetPerformanceObserverSupportedEntryTypes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceObserver_Observe
//go:noescape
func HasPerformanceObserverObserve(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceObserver_Observe
//go:noescape
func PerformanceObserverObserveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceObserver_Observe
//go:noescape
func CallPerformanceObserverObserve(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceObserver_Observe
//go:noescape
func TryPerformanceObserverObserve(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceObserver_Observe1
//go:noescape
func HasPerformanceObserverObserve1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceObserver_Observe1
//go:noescape
func PerformanceObserverObserve1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceObserver_Observe1
//go:noescape
func CallPerformanceObserverObserve1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceObserver_Observe1
//go:noescape
func TryPerformanceObserverObserve1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceObserver_Disconnect
//go:noescape
func HasPerformanceObserverDisconnect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceObserver_Disconnect
//go:noescape
func PerformanceObserverDisconnectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceObserver_Disconnect
//go:noescape
func CallPerformanceObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceObserver_Disconnect
//go:noescape
func TryPerformanceObserverDisconnect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceObserver_TakeRecords
//go:noescape
func HasPerformanceObserverTakeRecords(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceObserver_TakeRecords
//go:noescape
func PerformanceObserverTakeRecordsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceObserver_TakeRecords
//go:noescape
func CallPerformanceObserverTakeRecords(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceObserver_TakeRecords
//go:noescape
func TryPerformanceObserverTakeRecords(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_RenderBlockingStatusType
//go:noescape
func ConstOfRenderBlockingStatusType(str js.Ref) uint32

//go:wasmimport plat/js/web get_PerformanceServerTiming_Name
//go:noescape
func GetPerformanceServerTimingName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceServerTiming_Duration
//go:noescape
func GetPerformanceServerTimingDuration(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceServerTiming_Description
//go:noescape
func GetPerformanceServerTimingDescription(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceServerTiming_ToJSON
//go:noescape
func HasPerformanceServerTimingToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceServerTiming_ToJSON
//go:noescape
func PerformanceServerTimingToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceServerTiming_ToJSON
//go:noescape
func CallPerformanceServerTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceServerTiming_ToJSON
//go:noescape
func TryPerformanceServerTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_InitiatorType
//go:noescape
func GetPerformanceResourceTimingInitiatorType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_DeliveryType
//go:noescape
func GetPerformanceResourceTimingDeliveryType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_NextHopProtocol
//go:noescape
func GetPerformanceResourceTimingNextHopProtocol(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_WorkerStart
//go:noescape
func GetPerformanceResourceTimingWorkerStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_RedirectStart
//go:noescape
func GetPerformanceResourceTimingRedirectStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_RedirectEnd
//go:noescape
func GetPerformanceResourceTimingRedirectEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_FetchStart
//go:noescape
func GetPerformanceResourceTimingFetchStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_DomainLookupStart
//go:noescape
func GetPerformanceResourceTimingDomainLookupStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_DomainLookupEnd
//go:noescape
func GetPerformanceResourceTimingDomainLookupEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_ConnectStart
//go:noescape
func GetPerformanceResourceTimingConnectStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_ConnectEnd
//go:noescape
func GetPerformanceResourceTimingConnectEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_SecureConnectionStart
//go:noescape
func GetPerformanceResourceTimingSecureConnectionStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_RequestStart
//go:noescape
func GetPerformanceResourceTimingRequestStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_FirstInterimResponseStart
//go:noescape
func GetPerformanceResourceTimingFirstInterimResponseStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_ResponseStart
//go:noescape
func GetPerformanceResourceTimingResponseStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_ResponseEnd
//go:noescape
func GetPerformanceResourceTimingResponseEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_TransferSize
//go:noescape
func GetPerformanceResourceTimingTransferSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_EncodedBodySize
//go:noescape
func GetPerformanceResourceTimingEncodedBodySize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_DecodedBodySize
//go:noescape
func GetPerformanceResourceTimingDecodedBodySize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_ResponseStatus
//go:noescape
func GetPerformanceResourceTimingResponseStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_RenderBlockingStatus
//go:noescape
func GetPerformanceResourceTimingRenderBlockingStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_ContentType
//go:noescape
func GetPerformanceResourceTimingContentType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PerformanceResourceTiming_ServerTiming
//go:noescape
func GetPerformanceResourceTimingServerTiming(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PerformanceResourceTiming_ToJSON
//go:noescape
func HasPerformanceResourceTimingToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PerformanceResourceTiming_ToJSON
//go:noescape
func PerformanceResourceTimingToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PerformanceResourceTiming_ToJSON
//go:noescape
func CallPerformanceResourceTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PerformanceResourceTiming_ToJSON
//go:noescape
func TryPerformanceResourceTimingToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PeriodicSyncEventInit
//go:noescape
func PeriodicSyncEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PeriodicSyncEventInit
//go:noescape
func PeriodicSyncEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PeriodicSyncEvent_PeriodicSyncEvent
//go:noescape
func NewPeriodicSyncEventByPeriodicSyncEvent(
	typ js.Ref,
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_PeriodicSyncEvent_Tag
//go:noescape
func GetPeriodicSyncEventTag(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PermissionDescriptor
//go:noescape
func PermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PermissionDescriptor
//go:noescape
func PermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PermissionSetParameters
//go:noescape
func PermissionSetParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PermissionSetParameters
//go:noescape
func PermissionSetParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_PermissionsPolicyViolationReportBody_FeatureId
//go:noescape
func GetPermissionsPolicyViolationReportBodyFeatureId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PermissionsPolicyViolationReportBody_SourceFile
//go:noescape
func GetPermissionsPolicyViolationReportBodySourceFile(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PermissionsPolicyViolationReportBody_LineNumber
//go:noescape
func GetPermissionsPolicyViolationReportBodyLineNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PermissionsPolicyViolationReportBody_ColumnNumber
//go:noescape
func GetPermissionsPolicyViolationReportBodyColumnNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PermissionsPolicyViolationReportBody_Disposition
//go:noescape
func GetPermissionsPolicyViolationReportBodyDisposition(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PictureInPictureEventInit
//go:noescape
func PictureInPictureEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PictureInPictureEventInit
//go:noescape
func PictureInPictureEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PictureInPictureEvent_PictureInPictureEvent
//go:noescape
func NewPictureInPictureEventByPictureInPictureEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_PictureInPictureEvent_PictureInPictureWindow
//go:noescape
func GetPictureInPictureEventPictureInPictureWindow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PopStateEventInit
//go:noescape
func PopStateEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PopStateEventInit
//go:noescape
func PopStateEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PopStateEvent_PopStateEvent
//go:noescape
func NewPopStateEventByPopStateEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PopStateEvent_PopStateEvent1
//go:noescape
func NewPopStateEventByPopStateEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_PopStateEvent_State
//go:noescape
func GetPopStateEventState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PopStateEvent_HasUAVisualTransition
//go:noescape
func GetPopStateEventHasUAVisualTransition(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PortalActivateEventInit
//go:noescape
func PortalActivateEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PortalActivateEventInit
//go:noescape
func PortalActivateEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_PortalActivateEvent_PortalActivateEvent
//go:noescape
func NewPortalActivateEventByPortalActivateEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_PortalActivateEvent_PortalActivateEvent1
//go:noescape
func NewPortalActivateEventByPortalActivateEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_PortalActivateEvent_Data
//go:noescape
func GetPortalActivateEventData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PortalActivateEvent_AdoptPredecessor
//go:noescape
func HasPortalActivateEventAdoptPredecessor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PortalActivateEvent_AdoptPredecessor
//go:noescape
func PortalActivateEventAdoptPredecessorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PortalActivateEvent_AdoptPredecessor
//go:noescape
func CallPortalActivateEventAdoptPredecessor(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PortalActivateEvent_AdoptPredecessor
//go:noescape
func TryPortalActivateEventAdoptPredecessor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_PositionAlignSetting
//go:noescape
func ConstOfPositionAlignSetting(str js.Ref) uint32

//go:wasmimport plat/js/web store_PresentationConnectionAvailableEventInit
//go:noescape
func PresentationConnectionAvailableEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PresentationConnectionAvailableEventInit
//go:noescape
func PresentationConnectionAvailableEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref
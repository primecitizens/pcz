// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/webext/certificateprovider constof_Algorithm
//go:noescape
func ConstOfAlgorithm(str js.Ref) uint32

//go:wasmimport plat/js/webext/certificateprovider constof_Hash
//go:noescape
func ConstOfHash(str js.Ref) uint32

//go:wasmimport plat/js/webext/certificateprovider store_CertificateInfo
//go:noescape
func CertificateInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/certificateprovider load_CertificateInfo
//go:noescape
func CertificateInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/certificateprovider store_CertificatesUpdateRequest
//go:noescape
func CertificatesUpdateRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/certificateprovider load_CertificatesUpdateRequest
//go:noescape
func CertificatesUpdateRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/certificateprovider store_ClientCertificateInfo
//go:noescape
func ClientCertificateInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/certificateprovider load_ClientCertificateInfo
//go:noescape
func ClientCertificateInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/certificateprovider constof_Error
//go:noescape
func ConstOfError(str js.Ref) uint32

//go:wasmimport plat/js/webext/certificateprovider constof_PinRequestErrorType
//go:noescape
func ConstOfPinRequestErrorType(str js.Ref) uint32

//go:wasmimport plat/js/webext/certificateprovider constof_PinRequestType
//go:noescape
func ConstOfPinRequestType(str js.Ref) uint32

//go:wasmimport plat/js/webext/certificateprovider store_PinResponseDetails
//go:noescape
func PinResponseDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/certificateprovider load_PinResponseDetails
//go:noescape
func PinResponseDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/certificateprovider store_ReportSignatureDetails
//go:noescape
func ReportSignatureDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/certificateprovider load_ReportSignatureDetails
//go:noescape
func ReportSignatureDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/certificateprovider store_RequestPinDetails
//go:noescape
func RequestPinDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/certificateprovider load_RequestPinDetails
//go:noescape
func RequestPinDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/certificateprovider store_SetCertificatesDetails
//go:noescape
func SetCertificatesDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/certificateprovider load_SetCertificatesDetails
//go:noescape
func SetCertificatesDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/certificateprovider store_SignRequest
//go:noescape
func SignRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/certificateprovider load_SignRequest
//go:noescape
func SignRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/certificateprovider store_SignatureRequest
//go:noescape
func SignatureRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/certificateprovider load_SignatureRequest
//go:noescape
func SignatureRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/certificateprovider store_StopPinRequestDetails
//go:noescape
func StopPinRequestDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/certificateprovider load_StopPinRequestDetails
//go:noescape
func StopPinRequestDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/certificateprovider has_OnCertificatesRequested
//go:noescape
func HasFuncOnCertificatesRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_OnCertificatesRequested
//go:noescape
func FuncOnCertificatesRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_OnCertificatesRequested
//go:noescape
func CallOnCertificatesRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_OnCertificatesRequested
//go:noescape
func TryOnCertificatesRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_OffCertificatesRequested
//go:noescape
func HasFuncOffCertificatesRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_OffCertificatesRequested
//go:noescape
func FuncOffCertificatesRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_OffCertificatesRequested
//go:noescape
func CallOffCertificatesRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_OffCertificatesRequested
//go:noescape
func TryOffCertificatesRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_HasOnCertificatesRequested
//go:noescape
func HasFuncHasOnCertificatesRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_HasOnCertificatesRequested
//go:noescape
func FuncHasOnCertificatesRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_HasOnCertificatesRequested
//go:noescape
func CallHasOnCertificatesRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_HasOnCertificatesRequested
//go:noescape
func TryHasOnCertificatesRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_OnCertificatesUpdateRequested
//go:noescape
func HasFuncOnCertificatesUpdateRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_OnCertificatesUpdateRequested
//go:noescape
func FuncOnCertificatesUpdateRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_OnCertificatesUpdateRequested
//go:noescape
func CallOnCertificatesUpdateRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_OnCertificatesUpdateRequested
//go:noescape
func TryOnCertificatesUpdateRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_OffCertificatesUpdateRequested
//go:noescape
func HasFuncOffCertificatesUpdateRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_OffCertificatesUpdateRequested
//go:noescape
func FuncOffCertificatesUpdateRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_OffCertificatesUpdateRequested
//go:noescape
func CallOffCertificatesUpdateRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_OffCertificatesUpdateRequested
//go:noescape
func TryOffCertificatesUpdateRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_HasOnCertificatesUpdateRequested
//go:noescape
func HasFuncHasOnCertificatesUpdateRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_HasOnCertificatesUpdateRequested
//go:noescape
func FuncHasOnCertificatesUpdateRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_HasOnCertificatesUpdateRequested
//go:noescape
func CallHasOnCertificatesUpdateRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_HasOnCertificatesUpdateRequested
//go:noescape
func TryHasOnCertificatesUpdateRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_OnSignDigestRequested
//go:noescape
func HasFuncOnSignDigestRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_OnSignDigestRequested
//go:noescape
func FuncOnSignDigestRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_OnSignDigestRequested
//go:noescape
func CallOnSignDigestRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_OnSignDigestRequested
//go:noescape
func TryOnSignDigestRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_OffSignDigestRequested
//go:noescape
func HasFuncOffSignDigestRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_OffSignDigestRequested
//go:noescape
func FuncOffSignDigestRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_OffSignDigestRequested
//go:noescape
func CallOffSignDigestRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_OffSignDigestRequested
//go:noescape
func TryOffSignDigestRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_HasOnSignDigestRequested
//go:noescape
func HasFuncHasOnSignDigestRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_HasOnSignDigestRequested
//go:noescape
func FuncHasOnSignDigestRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_HasOnSignDigestRequested
//go:noescape
func CallHasOnSignDigestRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_HasOnSignDigestRequested
//go:noescape
func TryHasOnSignDigestRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_OnSignatureRequested
//go:noescape
func HasFuncOnSignatureRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_OnSignatureRequested
//go:noescape
func FuncOnSignatureRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_OnSignatureRequested
//go:noescape
func CallOnSignatureRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_OnSignatureRequested
//go:noescape
func TryOnSignatureRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_OffSignatureRequested
//go:noescape
func HasFuncOffSignatureRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_OffSignatureRequested
//go:noescape
func FuncOffSignatureRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_OffSignatureRequested
//go:noescape
func CallOffSignatureRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_OffSignatureRequested
//go:noescape
func TryOffSignatureRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_HasOnSignatureRequested
//go:noescape
func HasFuncHasOnSignatureRequested() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_HasOnSignatureRequested
//go:noescape
func FuncHasOnSignatureRequested(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_HasOnSignatureRequested
//go:noescape
func CallHasOnSignatureRequested(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/certificateprovider try_HasOnSignatureRequested
//go:noescape
func TryHasOnSignatureRequested(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_ReportSignature
//go:noescape
func HasFuncReportSignature() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_ReportSignature
//go:noescape
func FuncReportSignature(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_ReportSignature
//go:noescape
func CallReportSignature(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider try_ReportSignature
//go:noescape
func TryReportSignature(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_RequestPin
//go:noescape
func HasFuncRequestPin() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_RequestPin
//go:noescape
func FuncRequestPin(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_RequestPin
//go:noescape
func CallRequestPin(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider try_RequestPin
//go:noescape
func TryRequestPin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_SetCertificates
//go:noescape
func HasFuncSetCertificates() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_SetCertificates
//go:noescape
func FuncSetCertificates(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_SetCertificates
//go:noescape
func CallSetCertificates(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider try_SetCertificates
//go:noescape
func TrySetCertificates(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/certificateprovider has_StopPinRequest
//go:noescape
func HasFuncStopPinRequest() js.Ref

//go:wasmimport plat/js/webext/certificateprovider func_StopPinRequest
//go:noescape
func FuncStopPinRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider call_StopPinRequest
//go:noescape
func CallStopPinRequest(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/certificateprovider try_StopPinRequest
//go:noescape
func TryStopPinRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

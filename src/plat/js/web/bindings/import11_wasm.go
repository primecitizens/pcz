// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bindings

import (
	"unsafe"

	"github.com/primecitizens/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web get_AuthenticatorAttestationResponse_AttestationObject
//go:noescape

func GetAuthenticatorAttestationResponseAttestationObject(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_AuthenticatorAttestationResponse_GetTransports
//go:noescape

func CallAuthenticatorAttestationResponseGetTransports(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AuthenticatorAttestationResponse_GetTransports
//go:noescape

func AuthenticatorAttestationResponseGetTransportsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AuthenticatorAttestationResponse_GetAuthenticatorData
//go:noescape

func CallAuthenticatorAttestationResponseGetAuthenticatorData(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AuthenticatorAttestationResponse_GetAuthenticatorData
//go:noescape

func AuthenticatorAttestationResponseGetAuthenticatorDataFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AuthenticatorAttestationResponse_GetPublicKey
//go:noescape

func CallAuthenticatorAttestationResponseGetPublicKey(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_AuthenticatorAttestationResponse_GetPublicKey
//go:noescape

func AuthenticatorAttestationResponseGetPublicKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_AuthenticatorAttestationResponse_GetPublicKeyAlgorithm
//go:noescape

func CallAuthenticatorAttestationResponseGetPublicKeyAlgorithm(
	this js.Ref,
	ptr unsafe.Pointer,

) int32

//go:wasmimport plat/js/web func_AuthenticatorAttestationResponse_GetPublicKeyAlgorithm
//go:noescape

func AuthenticatorAttestationResponseGetPublicKeyAlgorithmFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_AuthenticatorAttestationResponseJSON
//go:noescape

func AuthenticatorAttestationResponseJSONJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticatorAttestationResponseJSON
//go:noescape

func AuthenticatorAttestationResponseJSONJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_AuthenticatorResponse_ClientDataJSON
//go:noescape

func GetAuthenticatorResponseClientDataJSON(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web store_AuthenticatorSelectionCriteria
//go:noescape

func AuthenticatorSelectionCriteriaJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AuthenticatorSelectionCriteria
//go:noescape

func AuthenticatorSelectionCriteriaJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_AuthenticatorTransport
//go:noescape

func ConstOfAuthenticatorTransport(str js.Ref) uint32

//go:wasmimport plat/js/web constof_AutoKeyword
//go:noescape

func ConstOfAutoKeyword(str js.Ref) uint32

//go:wasmimport plat/js/web constof_AutoplayPolicy
//go:noescape

func ConstOfAutoplayPolicy(str js.Ref) uint32

//go:wasmimport plat/js/web constof_AutoplayPolicyMediaType
//go:noescape

func ConstOfAutoplayPolicyMediaType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_AvcBitstreamFormat
//go:noescape

func ConstOfAvcBitstreamFormat(str js.Ref) uint32

//go:wasmimport plat/js/web store_AvcEncoderConfig
//go:noescape

func AvcEncoderConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_AvcEncoderConfig
//go:noescape

func AvcEncoderConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_URLSearchParams_URLSearchParams
//go:noescape

func NewURLSearchParamsByURLSearchParams(
	init js.Ref) js.Ref

//go:wasmimport plat/js/web new_URLSearchParams_URLSearchParams1
//go:noescape

func NewURLSearchParamsByURLSearchParams1() js.Ref

//go:wasmimport plat/js/web get_URLSearchParams_Size
//go:noescape

func GetURLSearchParamsSize(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_URLSearchParams_Append
//go:noescape

func CallURLSearchParamsAppend(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Append
//go:noescape

func URLSearchParamsAppendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_URLSearchParams_Delete
//go:noescape

func CallURLSearchParamsDelete(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Delete
//go:noescape

func URLSearchParamsDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_URLSearchParams_Delete1
//go:noescape

func CallURLSearchParamsDelete1(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Delete1
//go:noescape

func URLSearchParamsDelete1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_URLSearchParams_Get
//go:noescape

func CallURLSearchParamsGet(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Get
//go:noescape

func URLSearchParamsGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_URLSearchParams_GetAll
//go:noescape

func CallURLSearchParamsGetAll(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_GetAll
//go:noescape

func URLSearchParamsGetAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_URLSearchParams_Has
//go:noescape

func CallURLSearchParamsHas(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Has
//go:noescape

func URLSearchParamsHasFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_URLSearchParams_Has1
//go:noescape

func CallURLSearchParamsHas1(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Has1
//go:noescape

func URLSearchParamsHas1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_URLSearchParams_Set
//go:noescape

func CallURLSearchParamsSet(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Set
//go:noescape

func URLSearchParamsSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_URLSearchParams_Sort
//go:noescape

func CallURLSearchParamsSort(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Sort
//go:noescape

func URLSearchParamsSortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_URLSearchParams_ToString
//go:noescape

func CallURLSearchParamsToString(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_ToString
//go:noescape

func URLSearchParamsToStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ReferrerPolicy
//go:noescape

func ConstOfReferrerPolicy(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RequestMode
//go:noescape

func ConstOfRequestMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RequestCredentials
//go:noescape

func ConstOfRequestCredentials(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RequestCache
//go:noescape

func ConstOfRequestCache(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RequestRedirect
//go:noescape

func ConstOfRequestRedirect(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RequestDuplex
//go:noescape

func ConstOfRequestDuplex(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RequestPriority
//go:noescape

func ConstOfRequestPriority(str js.Ref) uint32

//go:wasmimport plat/js/web constof_TokenVersion
//go:noescape

func ConstOfTokenVersion(str js.Ref) uint32

//go:wasmimport plat/js/web constof_OperationType
//go:noescape

func ConstOfOperationType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_RefreshPolicy
//go:noescape

func ConstOfRefreshPolicy(str js.Ref) uint32

//go:wasmimport plat/js/web store_PrivateToken
//go:noescape

func PrivateTokenJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PrivateToken
//go:noescape

func PrivateTokenJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_RequestTargetAddressSpace
//go:noescape

func ConstOfRequestTargetAddressSpace(str js.Ref) uint32

//go:wasmimport plat/js/web store_RequestInit
//go:noescape

func RequestInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RequestInit
//go:noescape

func RequestInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_Headers_Headers
//go:noescape

func NewHeadersByHeaders(
	init js.Ref) js.Ref

//go:wasmimport plat/js/web new_Headers_Headers1
//go:noescape

func NewHeadersByHeaders1() js.Ref

//go:wasmimport plat/js/web call_Headers_Append
//go:noescape

func CallHeadersAppend(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Headers_Append
//go:noescape

func HeadersAppendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Headers_Delete
//go:noescape

func CallHeadersDelete(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Headers_Delete
//go:noescape

func HeadersDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Headers_Get
//go:noescape

func CallHeadersGet(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Headers_Get
//go:noescape

func HeadersGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Headers_GetSetCookie
//go:noescape

func CallHeadersGetSetCookie(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Headers_GetSetCookie
//go:noescape

func HeadersGetSetCookieFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Headers_Has
//go:noescape

func CallHeadersHas(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Headers_Has
//go:noescape

func HeadersHasFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Headers_Set
//go:noescape

func CallHeadersSet(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Headers_Set
//go:noescape

func HeadersSetFunc(this js.Ref) js.Ref

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

//go:wasmimport plat/js/web get_AuthenticatorAttestationResponse_AttestationObject
//go:noescape
func GetAuthenticatorAttestationResponseAttestationObject(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AuthenticatorAttestationResponse_GetTransports
//go:noescape
func HasFuncAuthenticatorAttestationResponseGetTransports(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AuthenticatorAttestationResponse_GetTransports
//go:noescape
func FuncAuthenticatorAttestationResponseGetTransports(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AuthenticatorAttestationResponse_GetTransports
//go:noescape
func CallAuthenticatorAttestationResponseGetTransports(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AuthenticatorAttestationResponse_GetTransports
//go:noescape
func TryAuthenticatorAttestationResponseGetTransports(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AuthenticatorAttestationResponse_GetAuthenticatorData
//go:noescape
func HasFuncAuthenticatorAttestationResponseGetAuthenticatorData(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AuthenticatorAttestationResponse_GetAuthenticatorData
//go:noescape
func FuncAuthenticatorAttestationResponseGetAuthenticatorData(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AuthenticatorAttestationResponse_GetAuthenticatorData
//go:noescape
func CallAuthenticatorAttestationResponseGetAuthenticatorData(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AuthenticatorAttestationResponse_GetAuthenticatorData
//go:noescape
func TryAuthenticatorAttestationResponseGetAuthenticatorData(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AuthenticatorAttestationResponse_GetPublicKey
//go:noescape
func HasFuncAuthenticatorAttestationResponseGetPublicKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AuthenticatorAttestationResponse_GetPublicKey
//go:noescape
func FuncAuthenticatorAttestationResponseGetPublicKey(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AuthenticatorAttestationResponse_GetPublicKey
//go:noescape
func CallAuthenticatorAttestationResponseGetPublicKey(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AuthenticatorAttestationResponse_GetPublicKey
//go:noescape
func TryAuthenticatorAttestationResponseGetPublicKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_AuthenticatorAttestationResponse_GetPublicKeyAlgorithm
//go:noescape
func HasFuncAuthenticatorAttestationResponseGetPublicKeyAlgorithm(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_AuthenticatorAttestationResponse_GetPublicKeyAlgorithm
//go:noescape
func FuncAuthenticatorAttestationResponseGetPublicKeyAlgorithm(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_AuthenticatorAttestationResponse_GetPublicKeyAlgorithm
//go:noescape
func CallAuthenticatorAttestationResponseGetPublicKeyAlgorithm(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_AuthenticatorAttestationResponse_GetPublicKeyAlgorithm
//go:noescape
func TryAuthenticatorAttestationResponseGetPublicKeyAlgorithm(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_URLSearchParams_Append
//go:noescape
func HasFuncURLSearchParamsAppend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Append
//go:noescape
func FuncURLSearchParamsAppend(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLSearchParams_Append
//go:noescape
func CallURLSearchParamsAppend(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_URLSearchParams_Append
//go:noescape
func TryURLSearchParamsAppend(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLSearchParams_Delete
//go:noescape
func HasFuncURLSearchParamsDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Delete
//go:noescape
func FuncURLSearchParamsDelete(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLSearchParams_Delete
//go:noescape
func CallURLSearchParamsDelete(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_URLSearchParams_Delete
//go:noescape
func TryURLSearchParamsDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLSearchParams_Delete1
//go:noescape
func HasFuncURLSearchParamsDelete1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Delete1
//go:noescape
func FuncURLSearchParamsDelete1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLSearchParams_Delete1
//go:noescape
func CallURLSearchParamsDelete1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_URLSearchParams_Delete1
//go:noescape
func TryURLSearchParamsDelete1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLSearchParams_Get
//go:noescape
func HasFuncURLSearchParamsGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Get
//go:noescape
func FuncURLSearchParamsGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLSearchParams_Get
//go:noescape
func CallURLSearchParamsGet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_URLSearchParams_Get
//go:noescape
func TryURLSearchParamsGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLSearchParams_GetAll
//go:noescape
func HasFuncURLSearchParamsGetAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_GetAll
//go:noescape
func FuncURLSearchParamsGetAll(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLSearchParams_GetAll
//go:noescape
func CallURLSearchParamsGetAll(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_URLSearchParams_GetAll
//go:noescape
func TryURLSearchParamsGetAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLSearchParams_Has
//go:noescape
func HasFuncURLSearchParamsHas(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Has
//go:noescape
func FuncURLSearchParamsHas(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLSearchParams_Has
//go:noescape
func CallURLSearchParamsHas(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_URLSearchParams_Has
//go:noescape
func TryURLSearchParamsHas(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLSearchParams_Has1
//go:noescape
func HasFuncURLSearchParamsHas1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Has1
//go:noescape
func FuncURLSearchParamsHas1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLSearchParams_Has1
//go:noescape
func CallURLSearchParamsHas1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_URLSearchParams_Has1
//go:noescape
func TryURLSearchParamsHas1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLSearchParams_Set
//go:noescape
func HasFuncURLSearchParamsSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Set
//go:noescape
func FuncURLSearchParamsSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLSearchParams_Set
//go:noescape
func CallURLSearchParamsSet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_URLSearchParams_Set
//go:noescape
func TryURLSearchParamsSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_URLSearchParams_Sort
//go:noescape
func HasFuncURLSearchParamsSort(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_Sort
//go:noescape
func FuncURLSearchParamsSort(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLSearchParams_Sort
//go:noescape
func CallURLSearchParamsSort(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_URLSearchParams_Sort
//go:noescape
func TryURLSearchParamsSort(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_URLSearchParams_ToString
//go:noescape
func HasFuncURLSearchParamsToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_URLSearchParams_ToString
//go:noescape
func FuncURLSearchParamsToString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_URLSearchParams_ToString
//go:noescape
func CallURLSearchParamsToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_URLSearchParams_ToString
//go:noescape
func TryURLSearchParamsToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/web has_Headers_Append
//go:noescape
func HasFuncHeadersAppend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Headers_Append
//go:noescape
func FuncHeadersAppend(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Headers_Append
//go:noescape
func CallHeadersAppend(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_Headers_Append
//go:noescape
func TryHeadersAppend(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Headers_Delete
//go:noescape
func HasFuncHeadersDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Headers_Delete
//go:noescape
func FuncHeadersDelete(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Headers_Delete
//go:noescape
func CallHeadersDelete(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_Headers_Delete
//go:noescape
func TryHeadersDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Headers_Get
//go:noescape
func HasFuncHeadersGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Headers_Get
//go:noescape
func FuncHeadersGet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Headers_Get
//go:noescape
func CallHeadersGet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_Headers_Get
//go:noescape
func TryHeadersGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Headers_GetSetCookie
//go:noescape
func HasFuncHeadersGetSetCookie(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Headers_GetSetCookie
//go:noescape
func FuncHeadersGetSetCookie(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Headers_GetSetCookie
//go:noescape
func CallHeadersGetSetCookie(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Headers_GetSetCookie
//go:noescape
func TryHeadersGetSetCookie(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Headers_Has
//go:noescape
func HasFuncHeadersHas(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Headers_Has
//go:noescape
func FuncHeadersHas(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Headers_Has
//go:noescape
func CallHeadersHas(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_Headers_Has
//go:noescape
func TryHeadersHas(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Headers_Set
//go:noescape
func HasFuncHeadersSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Headers_Set
//go:noescape
func FuncHeadersSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Headers_Set
//go:noescape
func CallHeadersSet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_Headers_Set
//go:noescape
func TryHeadersSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref) (ok js.Ref)

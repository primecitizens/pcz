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

//go:wasmimport plat/js/webext/echoprivate store_GetUserConsentArgConsentRequester
//go:noescape
func GetUserConsentArgConsentRequesterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/echoprivate load_GetUserConsentArgConsentRequester
//go:noescape
func GetUserConsentArgConsentRequesterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/echoprivate has_GetOfferInfo
//go:noescape
func HasFuncGetOfferInfo() js.Ref

//go:wasmimport plat/js/webext/echoprivate func_GetOfferInfo
//go:noescape
func FuncGetOfferInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/echoprivate call_GetOfferInfo
//go:noescape
func CallGetOfferInfo(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/echoprivate try_GetOfferInfo
//go:noescape
func TryGetOfferInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/echoprivate has_GetOobeTimestamp
//go:noescape
func HasFuncGetOobeTimestamp() js.Ref

//go:wasmimport plat/js/webext/echoprivate func_GetOobeTimestamp
//go:noescape
func FuncGetOobeTimestamp(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/echoprivate call_GetOobeTimestamp
//go:noescape
func CallGetOobeTimestamp(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/echoprivate try_GetOobeTimestamp
//go:noescape
func TryGetOobeTimestamp(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/echoprivate has_GetRegistrationCode
//go:noescape
func HasFuncGetRegistrationCode() js.Ref

//go:wasmimport plat/js/webext/echoprivate func_GetRegistrationCode
//go:noescape
func FuncGetRegistrationCode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/echoprivate call_GetRegistrationCode
//go:noescape
func CallGetRegistrationCode(
	retPtr unsafe.Pointer,
	typ js.Ref)

//go:wasmimport plat/js/webext/echoprivate try_GetRegistrationCode
//go:noescape
func TryGetRegistrationCode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/echoprivate has_GetUserConsent
//go:noescape
func HasFuncGetUserConsent() js.Ref

//go:wasmimport plat/js/webext/echoprivate func_GetUserConsent
//go:noescape
func FuncGetUserConsent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/echoprivate call_GetUserConsent
//go:noescape
func CallGetUserConsent(
	retPtr unsafe.Pointer,
	consentRequester unsafe.Pointer)

//go:wasmimport plat/js/webext/echoprivate try_GetUserConsent
//go:noescape
func TryGetUserConsent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	consentRequester unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/echoprivate has_SetOfferInfo
//go:noescape
func HasFuncSetOfferInfo() js.Ref

//go:wasmimport plat/js/webext/echoprivate func_SetOfferInfo
//go:noescape
func FuncSetOfferInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/echoprivate call_SetOfferInfo
//go:noescape
func CallSetOfferInfo(
	retPtr unsafe.Pointer,
	id js.Ref,
	offerInfo js.Ref)

//go:wasmimport plat/js/webext/echoprivate try_SetOfferInfo
//go:noescape
func TrySetOfferInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	offerInfo js.Ref) (ok js.Ref)

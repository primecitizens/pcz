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

//go:wasmimport plat/js/webext/platformkeys constof_ClientCertificateType
//go:noescape
func ConstOfClientCertificateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/platformkeys store_ClientCertificateRequest
//go:noescape
func ClientCertificateRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/platformkeys load_ClientCertificateRequest
//go:noescape
func ClientCertificateRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/platformkeys store_Match
//go:noescape
func MatchJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/platformkeys load_Match
//go:noescape
func MatchJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/platformkeys store_SelectDetails
//go:noescape
func SelectDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/platformkeys load_SelectDetails
//go:noescape
func SelectDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/platformkeys store_VerificationResult
//go:noescape
func VerificationResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/platformkeys load_VerificationResult
//go:noescape
func VerificationResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/platformkeys store_VerificationDetails
//go:noescape
func VerificationDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/platformkeys load_VerificationDetails
//go:noescape
func VerificationDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/platformkeys has_GetKeyPair
//go:noescape
func HasFuncGetKeyPair() js.Ref

//go:wasmimport plat/js/webext/platformkeys func_GetKeyPair
//go:noescape
func FuncGetKeyPair(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/platformkeys call_GetKeyPair
//go:noescape
func CallGetKeyPair(
	retPtr unsafe.Pointer,
	certificate js.Ref,
	parameters js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/platformkeys try_GetKeyPair
//go:noescape
func TryGetKeyPair(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	certificate js.Ref,
	parameters js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/platformkeys has_GetKeyPairBySpki
//go:noescape
func HasFuncGetKeyPairBySpki() js.Ref

//go:wasmimport plat/js/webext/platformkeys func_GetKeyPairBySpki
//go:noescape
func FuncGetKeyPairBySpki(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/platformkeys call_GetKeyPairBySpki
//go:noescape
func CallGetKeyPairBySpki(
	retPtr unsafe.Pointer,
	publicKeySpkiDer js.Ref,
	parameters js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/platformkeys try_GetKeyPairBySpki
//go:noescape
func TryGetKeyPairBySpki(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	publicKeySpkiDer js.Ref,
	parameters js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/platformkeys has_SelectClientCertificates
//go:noescape
func HasFuncSelectClientCertificates() js.Ref

//go:wasmimport plat/js/webext/platformkeys func_SelectClientCertificates
//go:noescape
func FuncSelectClientCertificates(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/platformkeys call_SelectClientCertificates
//go:noescape
func CallSelectClientCertificates(
	retPtr unsafe.Pointer,
	details unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/platformkeys try_SelectClientCertificates
//go:noescape
func TrySelectClientCertificates(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/platformkeys has_SubtleCrypto
//go:noescape
func HasFuncSubtleCrypto() js.Ref

//go:wasmimport plat/js/webext/platformkeys func_SubtleCrypto
//go:noescape
func FuncSubtleCrypto(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/platformkeys call_SubtleCrypto
//go:noescape
func CallSubtleCrypto(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/platformkeys try_SubtleCrypto
//go:noescape
func TrySubtleCrypto(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/platformkeys has_VerifyTLSServerCertificate
//go:noescape
func HasFuncVerifyTLSServerCertificate() js.Ref

//go:wasmimport plat/js/webext/platformkeys func_VerifyTLSServerCertificate
//go:noescape
func FuncVerifyTLSServerCertificate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/platformkeys call_VerifyTLSServerCertificate
//go:noescape
func CallVerifyTLSServerCertificate(
	retPtr unsafe.Pointer,
	details unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/platformkeys try_VerifyTLSServerCertificate
//go:noescape
func TryVerifyTLSServerCertificate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

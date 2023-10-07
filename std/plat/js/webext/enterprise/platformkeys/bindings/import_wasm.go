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

//go:wasmimport plat/js/webext/enterprise/platformkeys constof_Algorithm
//go:noescape
func ConstOfAlgorithm(str js.Ref) uint32

//go:wasmimport plat/js/webext/enterprise/platformkeys store_RegisterKeyOptions
//go:noescape
func RegisterKeyOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys load_RegisterKeyOptions
//go:noescape
func RegisterKeyOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeys constof_Scope
//go:noescape
func ConstOfScope(str js.Ref) uint32

//go:wasmimport plat/js/webext/enterprise/platformkeys store_ChallengeKeyOptions
//go:noescape
func ChallengeKeyOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys load_ChallengeKeyOptions
//go:noescape
func ChallengeKeyOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeys store_Token
//go:noescape
func TokenJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys load_Token
//go:noescape
func TokenJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeys has_ChallengeKey
//go:noescape
func HasFuncChallengeKey() js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeys func_ChallengeKey
//go:noescape
func FuncChallengeKey(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/platformkeys call_ChallengeKey
//go:noescape
func CallChallengeKey(
	retPtr unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys try_ChallengeKey
//go:noescape
func TryChallengeKey(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys has_ChallengeMachineKey
//go:noescape
func HasFuncChallengeMachineKey() js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeys func_ChallengeMachineKey
//go:noescape
func FuncChallengeMachineKey(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/platformkeys call_ChallengeMachineKey
//go:noescape
func CallChallengeMachineKey(
	retPtr unsafe.Pointer,
	challenge js.Ref,
	registerKey js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys try_ChallengeMachineKey
//go:noescape
func TryChallengeMachineKey(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	challenge js.Ref,
	registerKey js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys has_ChallengeUserKey
//go:noescape
func HasFuncChallengeUserKey() js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeys func_ChallengeUserKey
//go:noescape
func FuncChallengeUserKey(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/platformkeys call_ChallengeUserKey
//go:noescape
func CallChallengeUserKey(
	retPtr unsafe.Pointer,
	challenge js.Ref,
	registerKey js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys try_ChallengeUserKey
//go:noescape
func TryChallengeUserKey(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	challenge js.Ref,
	registerKey js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys has_GetCertificates
//go:noescape
func HasFuncGetCertificates() js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeys func_GetCertificates
//go:noescape
func FuncGetCertificates(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/platformkeys call_GetCertificates
//go:noescape
func CallGetCertificates(
	retPtr unsafe.Pointer,
	tokenId js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys try_GetCertificates
//go:noescape
func TryGetCertificates(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tokenId js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys has_GetTokens
//go:noescape
func HasFuncGetTokens() js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeys func_GetTokens
//go:noescape
func FuncGetTokens(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/platformkeys call_GetTokens
//go:noescape
func CallGetTokens(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys try_GetTokens
//go:noescape
func TryGetTokens(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys has_ImportCertificate
//go:noescape
func HasFuncImportCertificate() js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeys func_ImportCertificate
//go:noescape
func FuncImportCertificate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/platformkeys call_ImportCertificate
//go:noescape
func CallImportCertificate(
	retPtr unsafe.Pointer,
	tokenId js.Ref,
	certificate js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys try_ImportCertificate
//go:noescape
func TryImportCertificate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tokenId js.Ref,
	certificate js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys has_RemoveCertificate
//go:noescape
func HasFuncRemoveCertificate() js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeys func_RemoveCertificate
//go:noescape
func FuncRemoveCertificate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/platformkeys call_RemoveCertificate
//go:noescape
func CallRemoveCertificate(
	retPtr unsafe.Pointer,
	tokenId js.Ref,
	certificate js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeys try_RemoveCertificate
//go:noescape
func TryRemoveCertificate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tokenId js.Ref,
	certificate js.Ref,
	callback js.Ref) (ok js.Ref)

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

//go:wasmimport plat/js/webext/platformkeysinternal has_GetPublicKey
//go:noescape
func HasFuncGetPublicKey() js.Ref

//go:wasmimport plat/js/webext/platformkeysinternal func_GetPublicKey
//go:noescape
func FuncGetPublicKey(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/platformkeysinternal call_GetPublicKey
//go:noescape
func CallGetPublicKey(
	retPtr unsafe.Pointer,
	certificate js.Ref,
	algorithmName js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/platformkeysinternal try_GetPublicKey
//go:noescape
func TryGetPublicKey(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	certificate js.Ref,
	algorithmName js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/platformkeysinternal has_GetPublicKeyBySpki
//go:noescape
func HasFuncGetPublicKeyBySpki() js.Ref

//go:wasmimport plat/js/webext/platformkeysinternal func_GetPublicKeyBySpki
//go:noescape
func FuncGetPublicKeyBySpki(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/platformkeysinternal call_GetPublicKeyBySpki
//go:noescape
func CallGetPublicKeyBySpki(
	retPtr unsafe.Pointer,
	publicKeySpkiDer js.Ref,
	algorithmName js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/platformkeysinternal try_GetPublicKeyBySpki
//go:noescape
func TryGetPublicKeyBySpki(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	publicKeySpkiDer js.Ref,
	algorithmName js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/platformkeysinternal has_SelectClientCertificates
//go:noescape
func HasFuncSelectClientCertificates() js.Ref

//go:wasmimport plat/js/webext/platformkeysinternal func_SelectClientCertificates
//go:noescape
func FuncSelectClientCertificates(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/platformkeysinternal call_SelectClientCertificates
//go:noescape
func CallSelectClientCertificates(
	retPtr unsafe.Pointer,
	details unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/platformkeysinternal try_SelectClientCertificates
//go:noescape
func TrySelectClientCertificates(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/platformkeysinternal has_Sign
//go:noescape
func HasFuncSign() js.Ref

//go:wasmimport plat/js/webext/platformkeysinternal func_Sign
//go:noescape
func FuncSign(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/platformkeysinternal call_Sign
//go:noescape
func CallSign(
	retPtr unsafe.Pointer,
	tokenId js.Ref,
	publicKey js.Ref,
	algorithmName js.Ref,
	hashAlgorithmName js.Ref,
	data js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/platformkeysinternal try_Sign
//go:noescape
func TrySign(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tokenId js.Ref,
	publicKey js.Ref,
	algorithmName js.Ref,
	hashAlgorithmName js.Ref,
	data js.Ref,
	callback js.Ref) (ok js.Ref)

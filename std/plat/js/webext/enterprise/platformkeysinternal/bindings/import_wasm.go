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

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal store_Hash
//go:noescape
func HashJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal load_Hash
//go:noescape
func HashJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal store_Algorithm
//go:noescape
func AlgorithmJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal load_Algorithm
//go:noescape
func AlgorithmJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal has_GenerateKey
//go:noescape
func HasFuncGenerateKey() js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal func_GenerateKey
//go:noescape
func FuncGenerateKey(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal call_GenerateKey
//go:noescape
func CallGenerateKey(
	retPtr unsafe.Pointer,
	tokenId js.Ref,
	algorithm unsafe.Pointer,
	softwareBacked js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal try_GenerateKey
//go:noescape
func TryGenerateKey(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tokenId js.Ref,
	algorithm unsafe.Pointer,
	softwareBacked js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal has_GetTokens
//go:noescape
func HasFuncGetTokens() js.Ref

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal func_GetTokens
//go:noescape
func FuncGetTokens(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal call_GetTokens
//go:noescape
func CallGetTokens(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/enterprise/platformkeysinternal try_GetTokens
//go:noescape
func TryGetTokens(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

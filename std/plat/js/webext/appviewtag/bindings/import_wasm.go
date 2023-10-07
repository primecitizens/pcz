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

//go:wasmimport plat/js/webext/appviewtag get_EmbedRequest_EmbedderId
//go:noescape
func GetEmbedRequestEmbedderId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/appviewtag set_EmbedRequest_EmbedderId
//go:noescape
func SetEmbedRequestEmbedderId(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/appviewtag get_EmbedRequest_Data
//go:noescape
func GetEmbedRequestData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/appviewtag set_EmbedRequest_Data
//go:noescape
func SetEmbedRequestData(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/appviewtag has_EmbedRequest_Allow
//go:noescape
func HasFuncEmbedRequestAllow(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/appviewtag func_EmbedRequest_Allow
//go:noescape
func FuncEmbedRequestAllow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/appviewtag call_EmbedRequest_Allow
//go:noescape
func CallEmbedRequestAllow(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/appviewtag try_EmbedRequest_Allow
//go:noescape
func TryEmbedRequestAllow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/appviewtag has_EmbedRequest_Deny
//go:noescape
func HasFuncEmbedRequestDeny(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/appviewtag func_EmbedRequest_Deny
//go:noescape
func FuncEmbedRequestDeny(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/appviewtag call_EmbedRequest_Deny
//go:noescape
func CallEmbedRequestDeny(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/appviewtag try_EmbedRequest_Deny
//go:noescape
func TryEmbedRequestDeny(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/appviewtag has_Connect
//go:noescape
func HasFuncConnect() js.Ref

//go:wasmimport plat/js/webext/appviewtag func_Connect
//go:noescape
func FuncConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/appviewtag call_Connect
//go:noescape
func CallConnect(
	retPtr unsafe.Pointer,
	app js.Ref,
	data js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/appviewtag try_Connect
//go:noescape
func TryConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	app js.Ref,
	data js.Ref,
	callback js.Ref) (ok js.Ref)

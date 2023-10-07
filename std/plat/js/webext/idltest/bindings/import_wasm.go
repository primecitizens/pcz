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

//go:wasmimport plat/js/webext/idltest has_GetArrayBuffer
//go:noescape
func HasFuncGetArrayBuffer() js.Ref

//go:wasmimport plat/js/webext/idltest func_GetArrayBuffer
//go:noescape
func FuncGetArrayBuffer(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/idltest call_GetArrayBuffer
//go:noescape
func CallGetArrayBuffer(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/idltest try_GetArrayBuffer
//go:noescape
func TryGetArrayBuffer(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/idltest has_NocompileFunc
//go:noescape
func HasFuncNocompileFunc() js.Ref

//go:wasmimport plat/js/webext/idltest func_NocompileFunc
//go:noescape
func FuncNocompileFunc(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/idltest call_NocompileFunc
//go:noescape
func CallNocompileFunc(
	retPtr unsafe.Pointer,
	switch_ int32)

//go:wasmimport plat/js/webext/idltest try_NocompileFunc
//go:noescape
func TryNocompileFunc(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	switch_ int32) (ok js.Ref)

//go:wasmimport plat/js/webext/idltest has_SendArrayBuffer
//go:noescape
func HasFuncSendArrayBuffer() js.Ref

//go:wasmimport plat/js/webext/idltest func_SendArrayBuffer
//go:noescape
func FuncSendArrayBuffer(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/idltest call_SendArrayBuffer
//go:noescape
func CallSendArrayBuffer(
	retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/webext/idltest try_SendArrayBuffer
//go:noescape
func TrySendArrayBuffer(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/idltest has_SendArrayBufferView
//go:noescape
func HasFuncSendArrayBufferView() js.Ref

//go:wasmimport plat/js/webext/idltest func_SendArrayBufferView
//go:noescape
func FuncSendArrayBufferView(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/idltest call_SendArrayBufferView
//go:noescape
func CallSendArrayBufferView(
	retPtr unsafe.Pointer,
	input js.Ref)

//go:wasmimport plat/js/webext/idltest try_SendArrayBufferView
//go:noescape
func TrySendArrayBufferView(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	input js.Ref) (ok js.Ref)

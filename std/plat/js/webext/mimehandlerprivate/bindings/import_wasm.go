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

//go:wasmimport plat/js/webext/mimehandlerprivate store_StreamInfo
//go:noescape
func StreamInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate load_StreamInfo
//go:noescape
func StreamInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mimehandlerprivate store_PdfPluginAttributes
//go:noescape
func PdfPluginAttributesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate load_PdfPluginAttributes
//go:noescape
func PdfPluginAttributesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/mimehandlerprivate has_GetStreamInfo
//go:noescape
func HasFuncGetStreamInfo() js.Ref

//go:wasmimport plat/js/webext/mimehandlerprivate func_GetStreamInfo
//go:noescape
func FuncGetStreamInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mimehandlerprivate call_GetStreamInfo
//go:noescape
func CallGetStreamInfo(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate try_GetStreamInfo
//go:noescape
func TryGetStreamInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate has_OnSave
//go:noescape
func HasFuncOnSave() js.Ref

//go:wasmimport plat/js/webext/mimehandlerprivate func_OnSave
//go:noescape
func FuncOnSave(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mimehandlerprivate call_OnSave
//go:noescape
func CallOnSave(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate try_OnSave
//go:noescape
func TryOnSave(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate has_OffSave
//go:noescape
func HasFuncOffSave() js.Ref

//go:wasmimport plat/js/webext/mimehandlerprivate func_OffSave
//go:noescape
func FuncOffSave(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mimehandlerprivate call_OffSave
//go:noescape
func CallOffSave(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate try_OffSave
//go:noescape
func TryOffSave(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate has_HasOnSave
//go:noescape
func HasFuncHasOnSave() js.Ref

//go:wasmimport plat/js/webext/mimehandlerprivate func_HasOnSave
//go:noescape
func FuncHasOnSave(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mimehandlerprivate call_HasOnSave
//go:noescape
func CallHasOnSave(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate try_HasOnSave
//go:noescape
func TryHasOnSave(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate has_SetPdfPluginAttributes
//go:noescape
func HasFuncSetPdfPluginAttributes() js.Ref

//go:wasmimport plat/js/webext/mimehandlerprivate func_SetPdfPluginAttributes
//go:noescape
func FuncSetPdfPluginAttributes(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mimehandlerprivate call_SetPdfPluginAttributes
//go:noescape
func CallSetPdfPluginAttributes(
	retPtr unsafe.Pointer,
	pdfPluginAttributes unsafe.Pointer)

//go:wasmimport plat/js/webext/mimehandlerprivate try_SetPdfPluginAttributes
//go:noescape
func TrySetPdfPluginAttributes(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	pdfPluginAttributes unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate has_SetShowBeforeUnloadDialog
//go:noescape
func HasFuncSetShowBeforeUnloadDialog() js.Ref

//go:wasmimport plat/js/webext/mimehandlerprivate func_SetShowBeforeUnloadDialog
//go:noescape
func FuncSetShowBeforeUnloadDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/mimehandlerprivate call_SetShowBeforeUnloadDialog
//go:noescape
func CallSetShowBeforeUnloadDialog(
	retPtr unsafe.Pointer,
	showDialog js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/mimehandlerprivate try_SetShowBeforeUnloadDialog
//go:noescape
func TrySetShowBeforeUnloadDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	showDialog js.Ref,
	callback js.Ref) (ok js.Ref)

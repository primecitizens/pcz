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

//go:wasmimport plat/js/webext/clipboard constof_DataItemType
//go:noescape
func ConstOfDataItemType(str js.Ref) uint32

//go:wasmimport plat/js/webext/clipboard store_AdditionalDataItem
//go:noescape
func AdditionalDataItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/clipboard load_AdditionalDataItem
//go:noescape
func AdditionalDataItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/clipboard constof_ImageType
//go:noescape
func ConstOfImageType(str js.Ref) uint32

//go:wasmimport plat/js/webext/clipboard has_OnClipboardDataChanged
//go:noescape
func HasFuncOnClipboardDataChanged() js.Ref

//go:wasmimport plat/js/webext/clipboard func_OnClipboardDataChanged
//go:noescape
func FuncOnClipboardDataChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/clipboard call_OnClipboardDataChanged
//go:noescape
func CallOnClipboardDataChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/clipboard try_OnClipboardDataChanged
//go:noescape
func TryOnClipboardDataChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/clipboard has_OffClipboardDataChanged
//go:noescape
func HasFuncOffClipboardDataChanged() js.Ref

//go:wasmimport plat/js/webext/clipboard func_OffClipboardDataChanged
//go:noescape
func FuncOffClipboardDataChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/clipboard call_OffClipboardDataChanged
//go:noescape
func CallOffClipboardDataChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/clipboard try_OffClipboardDataChanged
//go:noescape
func TryOffClipboardDataChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/clipboard has_HasOnClipboardDataChanged
//go:noescape
func HasFuncHasOnClipboardDataChanged() js.Ref

//go:wasmimport plat/js/webext/clipboard func_HasOnClipboardDataChanged
//go:noescape
func FuncHasOnClipboardDataChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/clipboard call_HasOnClipboardDataChanged
//go:noescape
func CallHasOnClipboardDataChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/clipboard try_HasOnClipboardDataChanged
//go:noescape
func TryHasOnClipboardDataChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/clipboard has_SetImageData
//go:noescape
func HasFuncSetImageData() js.Ref

//go:wasmimport plat/js/webext/clipboard func_SetImageData
//go:noescape
func FuncSetImageData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/clipboard call_SetImageData
//go:noescape
func CallSetImageData(
	retPtr unsafe.Pointer,
	imageData js.Ref,
	typ uint32,
	additionalItems js.Ref)

//go:wasmimport plat/js/webext/clipboard try_SetImageData
//go:noescape
func TrySetImageData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	imageData js.Ref,
	typ uint32,
	additionalItems js.Ref) (ok js.Ref)

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

//go:wasmimport plat/js/webext/pdfviewerprivate has_IsAllowedLocalFileAccess
//go:noescape
func HasFuncIsAllowedLocalFileAccess() js.Ref

//go:wasmimport plat/js/webext/pdfviewerprivate func_IsAllowedLocalFileAccess
//go:noescape
func FuncIsAllowedLocalFileAccess(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pdfviewerprivate call_IsAllowedLocalFileAccess
//go:noescape
func CallIsAllowedLocalFileAccess(
	retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/webext/pdfviewerprivate try_IsAllowedLocalFileAccess
//go:noescape
func TryIsAllowedLocalFileAccess(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/pdfviewerprivate has_IsPdfOcrAlwaysActive
//go:noescape
func HasFuncIsPdfOcrAlwaysActive() js.Ref

//go:wasmimport plat/js/webext/pdfviewerprivate func_IsPdfOcrAlwaysActive
//go:noescape
func FuncIsPdfOcrAlwaysActive(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pdfviewerprivate call_IsPdfOcrAlwaysActive
//go:noescape
func CallIsPdfOcrAlwaysActive(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/pdfviewerprivate try_IsPdfOcrAlwaysActive
//go:noescape
func TryIsPdfOcrAlwaysActive(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/pdfviewerprivate has_OnPdfOcrPrefChanged
//go:noescape
func HasFuncOnPdfOcrPrefChanged() js.Ref

//go:wasmimport plat/js/webext/pdfviewerprivate func_OnPdfOcrPrefChanged
//go:noescape
func FuncOnPdfOcrPrefChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pdfviewerprivate call_OnPdfOcrPrefChanged
//go:noescape
func CallOnPdfOcrPrefChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/pdfviewerprivate try_OnPdfOcrPrefChanged
//go:noescape
func TryOnPdfOcrPrefChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/pdfviewerprivate has_OffPdfOcrPrefChanged
//go:noescape
func HasFuncOffPdfOcrPrefChanged() js.Ref

//go:wasmimport plat/js/webext/pdfviewerprivate func_OffPdfOcrPrefChanged
//go:noescape
func FuncOffPdfOcrPrefChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pdfviewerprivate call_OffPdfOcrPrefChanged
//go:noescape
func CallOffPdfOcrPrefChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/pdfviewerprivate try_OffPdfOcrPrefChanged
//go:noescape
func TryOffPdfOcrPrefChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/pdfviewerprivate has_HasOnPdfOcrPrefChanged
//go:noescape
func HasFuncHasOnPdfOcrPrefChanged() js.Ref

//go:wasmimport plat/js/webext/pdfviewerprivate func_HasOnPdfOcrPrefChanged
//go:noescape
func FuncHasOnPdfOcrPrefChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pdfviewerprivate call_HasOnPdfOcrPrefChanged
//go:noescape
func CallHasOnPdfOcrPrefChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/pdfviewerprivate try_HasOnPdfOcrPrefChanged
//go:noescape
func TryHasOnPdfOcrPrefChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/pdfviewerprivate has_SetPdfOcrPref
//go:noescape
func HasFuncSetPdfOcrPref() js.Ref

//go:wasmimport plat/js/webext/pdfviewerprivate func_SetPdfOcrPref
//go:noescape
func FuncSetPdfOcrPref(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/pdfviewerprivate call_SetPdfOcrPref
//go:noescape
func CallSetPdfOcrPref(
	retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/webext/pdfviewerprivate try_SetPdfOcrPref
//go:noescape
func TrySetPdfOcrPref(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

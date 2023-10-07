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

//go:wasmimport plat/js/webext/imageloaderprivate has_GetArcDocumentsProviderThumbnail
//go:noescape
func HasFuncGetArcDocumentsProviderThumbnail() js.Ref

//go:wasmimport plat/js/webext/imageloaderprivate func_GetArcDocumentsProviderThumbnail
//go:noescape
func FuncGetArcDocumentsProviderThumbnail(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imageloaderprivate call_GetArcDocumentsProviderThumbnail
//go:noescape
func CallGetArcDocumentsProviderThumbnail(
	retPtr unsafe.Pointer,
	url js.Ref,
	widthHint int32,
	heightHint int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/imageloaderprivate try_GetArcDocumentsProviderThumbnail
//go:noescape
func TryGetArcDocumentsProviderThumbnail(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	widthHint int32,
	heightHint int32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imageloaderprivate has_GetDriveThumbnail
//go:noescape
func HasFuncGetDriveThumbnail() js.Ref

//go:wasmimport plat/js/webext/imageloaderprivate func_GetDriveThumbnail
//go:noescape
func FuncGetDriveThumbnail(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imageloaderprivate call_GetDriveThumbnail
//go:noescape
func CallGetDriveThumbnail(
	retPtr unsafe.Pointer,
	url js.Ref,
	cropToSquare js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/imageloaderprivate try_GetDriveThumbnail
//go:noescape
func TryGetDriveThumbnail(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	cropToSquare js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/imageloaderprivate has_GetPdfThumbnail
//go:noescape
func HasFuncGetPdfThumbnail() js.Ref

//go:wasmimport plat/js/webext/imageloaderprivate func_GetPdfThumbnail
//go:noescape
func FuncGetPdfThumbnail(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/imageloaderprivate call_GetPdfThumbnail
//go:noescape
func CallGetPdfThumbnail(
	retPtr unsafe.Pointer,
	url js.Ref,
	width int32,
	height int32,
	callback js.Ref)

//go:wasmimport plat/js/webext/imageloaderprivate try_GetPdfThumbnail
//go:noescape
func TryGetPdfThumbnail(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	width int32,
	height int32,
	callback js.Ref) (ok js.Ref)

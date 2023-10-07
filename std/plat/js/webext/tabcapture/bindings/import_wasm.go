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

//go:wasmimport plat/js/webext/tabcapture constof_TabCaptureState
//go:noescape
func ConstOfTabCaptureState(str js.Ref) uint32

//go:wasmimport plat/js/webext/tabcapture store_CaptureInfo
//go:noescape
func CaptureInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabcapture load_CaptureInfo
//go:noescape
func CaptureInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabcapture store_MediaStreamConstraint
//go:noescape
func MediaStreamConstraintJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabcapture load_MediaStreamConstraint
//go:noescape
func MediaStreamConstraintJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabcapture store_CaptureOptions
//go:noescape
func CaptureOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabcapture load_CaptureOptions
//go:noescape
func CaptureOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabcapture store_GetMediaStreamOptions
//go:noescape
func GetMediaStreamOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabcapture load_GetMediaStreamOptions
//go:noescape
func GetMediaStreamOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabcapture has_Capture
//go:noescape
func HasFuncCapture() js.Ref

//go:wasmimport plat/js/webext/tabcapture func_Capture
//go:noescape
func FuncCapture(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabcapture call_Capture
//go:noescape
func CallCapture(
	retPtr unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabcapture try_Capture
//go:noescape
func TryCapture(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabcapture has_GetCapturedTabs
//go:noescape
func HasFuncGetCapturedTabs() js.Ref

//go:wasmimport plat/js/webext/tabcapture func_GetCapturedTabs
//go:noescape
func FuncGetCapturedTabs(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabcapture call_GetCapturedTabs
//go:noescape
func CallGetCapturedTabs(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/tabcapture try_GetCapturedTabs
//go:noescape
func TryGetCapturedTabs(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabcapture has_GetMediaStreamId
//go:noescape
func HasFuncGetMediaStreamId() js.Ref

//go:wasmimport plat/js/webext/tabcapture func_GetMediaStreamId
//go:noescape
func FuncGetMediaStreamId(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabcapture call_GetMediaStreamId
//go:noescape
func CallGetMediaStreamId(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/tabcapture try_GetMediaStreamId
//go:noescape
func TryGetMediaStreamId(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabcapture has_OnStatusChanged
//go:noescape
func HasFuncOnStatusChanged() js.Ref

//go:wasmimport plat/js/webext/tabcapture func_OnStatusChanged
//go:noescape
func FuncOnStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabcapture call_OnStatusChanged
//go:noescape
func CallOnStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabcapture try_OnStatusChanged
//go:noescape
func TryOnStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabcapture has_OffStatusChanged
//go:noescape
func HasFuncOffStatusChanged() js.Ref

//go:wasmimport plat/js/webext/tabcapture func_OffStatusChanged
//go:noescape
func FuncOffStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabcapture call_OffStatusChanged
//go:noescape
func CallOffStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabcapture try_OffStatusChanged
//go:noescape
func TryOffStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabcapture has_HasOnStatusChanged
//go:noescape
func HasFuncHasOnStatusChanged() js.Ref

//go:wasmimport plat/js/webext/tabcapture func_HasOnStatusChanged
//go:noescape
func FuncHasOnStatusChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabcapture call_HasOnStatusChanged
//go:noescape
func CallHasOnStatusChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabcapture try_HasOnStatusChanged
//go:noescape
func TryHasOnStatusChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

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

//go:wasmimport plat/js/webext/webrtcdesktopcaptureprivate constof_DesktopCaptureSourceType
//go:noescape
func ConstOfDesktopCaptureSourceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/webrtcdesktopcaptureprivate store_RequestInfo
//go:noescape
func RequestInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webrtcdesktopcaptureprivate load_RequestInfo
//go:noescape
func RequestInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webrtcdesktopcaptureprivate has_CancelChooseDesktopMedia
//go:noescape
func HasFuncCancelChooseDesktopMedia() js.Ref

//go:wasmimport plat/js/webext/webrtcdesktopcaptureprivate func_CancelChooseDesktopMedia
//go:noescape
func FuncCancelChooseDesktopMedia(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcdesktopcaptureprivate call_CancelChooseDesktopMedia
//go:noescape
func CallCancelChooseDesktopMedia(
	retPtr unsafe.Pointer,
	desktopMediaRequestId int32)

//go:wasmimport plat/js/webext/webrtcdesktopcaptureprivate try_CancelChooseDesktopMedia
//go:noescape
func TryCancelChooseDesktopMedia(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	desktopMediaRequestId int32) (ok js.Ref)

//go:wasmimport plat/js/webext/webrtcdesktopcaptureprivate has_ChooseDesktopMedia
//go:noescape
func HasFuncChooseDesktopMedia() js.Ref

//go:wasmimport plat/js/webext/webrtcdesktopcaptureprivate func_ChooseDesktopMedia
//go:noescape
func FuncChooseDesktopMedia(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webrtcdesktopcaptureprivate call_ChooseDesktopMedia
//go:noescape
func CallChooseDesktopMedia(
	retPtr unsafe.Pointer,
	sources js.Ref,
	request unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webrtcdesktopcaptureprivate try_ChooseDesktopMedia
//go:noescape
func TryChooseDesktopMedia(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sources js.Ref,
	request unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

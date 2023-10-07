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

//go:wasmimport plat/js/webext/desktopcapture store_ChooseDesktopMediaArgCallbackArgOptions
//go:noescape
func ChooseDesktopMediaArgCallbackArgOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/desktopcapture load_ChooseDesktopMediaArgCallbackArgOptions
//go:noescape
func ChooseDesktopMediaArgCallbackArgOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/desktopcapture constof_SelfCapturePreferenceEnum
//go:noescape
func ConstOfSelfCapturePreferenceEnum(str js.Ref) uint32

//go:wasmimport plat/js/webext/desktopcapture constof_SystemAudioPreferenceEnum
//go:noescape
func ConstOfSystemAudioPreferenceEnum(str js.Ref) uint32

//go:wasmimport plat/js/webext/desktopcapture store_ChooseDesktopMediaArgOptions
//go:noescape
func ChooseDesktopMediaArgOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/desktopcapture load_ChooseDesktopMediaArgOptions
//go:noescape
func ChooseDesktopMediaArgOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/desktopcapture constof_DesktopCaptureSourceType
//go:noescape
func ConstOfDesktopCaptureSourceType(str js.Ref) uint32

//go:wasmimport plat/js/webext/desktopcapture has_CancelChooseDesktopMedia
//go:noescape
func HasFuncCancelChooseDesktopMedia() js.Ref

//go:wasmimport plat/js/webext/desktopcapture func_CancelChooseDesktopMedia
//go:noescape
func FuncCancelChooseDesktopMedia(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/desktopcapture call_CancelChooseDesktopMedia
//go:noescape
func CallCancelChooseDesktopMedia(
	retPtr unsafe.Pointer,
	desktopMediaRequestId float64)

//go:wasmimport plat/js/webext/desktopcapture try_CancelChooseDesktopMedia
//go:noescape
func TryCancelChooseDesktopMedia(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	desktopMediaRequestId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/desktopcapture has_ChooseDesktopMedia
//go:noescape
func HasFuncChooseDesktopMedia() js.Ref

//go:wasmimport plat/js/webext/desktopcapture func_ChooseDesktopMedia
//go:noescape
func FuncChooseDesktopMedia(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/desktopcapture call_ChooseDesktopMedia
//go:noescape
func CallChooseDesktopMedia(
	retPtr unsafe.Pointer,
	sources js.Ref,
	targetTab unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/desktopcapture try_ChooseDesktopMedia
//go:noescape
func TryChooseDesktopMedia(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sources js.Ref,
	targetTab unsafe.Pointer,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

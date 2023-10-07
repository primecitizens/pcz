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

//go:wasmimport plat/js/webext/sidepanel store_GetPanelOptions
//go:noescape
func GetPanelOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sidepanel load_GetPanelOptions
//go:noescape
func GetPanelOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sidepanel store_SidePanel
//go:noescape
func SidePanelJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sidepanel load_SidePanel
//go:noescape
func SidePanelJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sidepanel store_ManifestKeys
//go:noescape
func ManifestKeysJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sidepanel load_ManifestKeys
//go:noescape
func ManifestKeysJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sidepanel store_OpenOptions
//go:noescape
func OpenOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sidepanel load_OpenOptions
//go:noescape
func OpenOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sidepanel store_PanelBehavior
//go:noescape
func PanelBehaviorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sidepanel load_PanelBehavior
//go:noescape
func PanelBehaviorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sidepanel store_PanelOptions
//go:noescape
func PanelOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/sidepanel load_PanelOptions
//go:noescape
func PanelOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/sidepanel has_GetOptions
//go:noescape
func HasFuncGetOptions() js.Ref

//go:wasmimport plat/js/webext/sidepanel func_GetOptions
//go:noescape
func FuncGetOptions(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sidepanel call_GetOptions
//go:noescape
func CallGetOptions(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/sidepanel try_GetOptions
//go:noescape
func TryGetOptions(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/sidepanel has_GetPanelBehavior
//go:noescape
func HasFuncGetPanelBehavior() js.Ref

//go:wasmimport plat/js/webext/sidepanel func_GetPanelBehavior
//go:noescape
func FuncGetPanelBehavior(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sidepanel call_GetPanelBehavior
//go:noescape
func CallGetPanelBehavior(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/sidepanel try_GetPanelBehavior
//go:noescape
func TryGetPanelBehavior(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/sidepanel has_Open
//go:noescape
func HasFuncOpen() js.Ref

//go:wasmimport plat/js/webext/sidepanel func_Open
//go:noescape
func FuncOpen(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sidepanel call_Open
//go:noescape
func CallOpen(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/sidepanel try_Open
//go:noescape
func TryOpen(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/sidepanel has_SetOptions
//go:noescape
func HasFuncSetOptions() js.Ref

//go:wasmimport plat/js/webext/sidepanel func_SetOptions
//go:noescape
func FuncSetOptions(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sidepanel call_SetOptions
//go:noescape
func CallSetOptions(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/sidepanel try_SetOptions
//go:noescape
func TrySetOptions(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/sidepanel has_SetPanelBehavior
//go:noescape
func HasFuncSetPanelBehavior() js.Ref

//go:wasmimport plat/js/webext/sidepanel func_SetPanelBehavior
//go:noescape
func FuncSetPanelBehavior(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/sidepanel call_SetPanelBehavior
//go:noescape
func CallSetPanelBehavior(
	retPtr unsafe.Pointer,
	behavior unsafe.Pointer)

//go:wasmimport plat/js/webext/sidepanel try_SetPanelBehavior
//go:noescape
func TrySetPanelBehavior(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	behavior unsafe.Pointer) (ok js.Ref)

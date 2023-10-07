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

//go:wasmimport plat/js/webext/app/currentwindowinternal store_Bounds
//go:noescape
func BoundsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal load_Bounds
//go:noescape
func BoundsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal store_RegionRect
//go:noescape
func RegionRectJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal load_RegionRect
//go:noescape
func RegionRectJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal store_Region
//go:noescape
func RegionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal load_Region
//go:noescape
func RegionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal store_SizeConstraints
//go:noescape
func SizeConstraintsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal load_SizeConstraints
//go:noescape
func SizeConstraintsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal has_ClearAttention
//go:noescape
func HasFuncClearAttention() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_ClearAttention
//go:noescape
func FuncClearAttention(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_ClearAttention
//go:noescape
func CallClearAttention(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_ClearAttention
//go:noescape
func TryClearAttention(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_DrawAttention
//go:noescape
func HasFuncDrawAttention() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_DrawAttention
//go:noescape
func FuncDrawAttention(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_DrawAttention
//go:noescape
func CallDrawAttention(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_DrawAttention
//go:noescape
func TryDrawAttention(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_Focus
//go:noescape
func HasFuncFocus() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_Focus
//go:noescape
func FuncFocus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_Focus
//go:noescape
func CallFocus(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_Focus
//go:noescape
func TryFocus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_Fullscreen
//go:noescape
func HasFuncFullscreen() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_Fullscreen
//go:noescape
func FuncFullscreen(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_Fullscreen
//go:noescape
func CallFullscreen(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_Fullscreen
//go:noescape
func TryFullscreen(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_Hide
//go:noescape
func HasFuncHide() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_Hide
//go:noescape
func FuncHide(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_Hide
//go:noescape
func CallHide(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_Hide
//go:noescape
func TryHide(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_Maximize
//go:noescape
func HasFuncMaximize() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_Maximize
//go:noescape
func FuncMaximize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_Maximize
//go:noescape
func CallMaximize(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_Maximize
//go:noescape
func TryMaximize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_Minimize
//go:noescape
func HasFuncMinimize() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_Minimize
//go:noescape
func FuncMinimize(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_Minimize
//go:noescape
func CallMinimize(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_Minimize
//go:noescape
func TryMinimize(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OnAlphaEnabledChanged
//go:noescape
func HasFuncOnAlphaEnabledChanged() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OnAlphaEnabledChanged
//go:noescape
func FuncOnAlphaEnabledChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OnAlphaEnabledChanged
//go:noescape
func CallOnAlphaEnabledChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OnAlphaEnabledChanged
//go:noescape
func TryOnAlphaEnabledChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OffAlphaEnabledChanged
//go:noescape
func HasFuncOffAlphaEnabledChanged() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OffAlphaEnabledChanged
//go:noescape
func FuncOffAlphaEnabledChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OffAlphaEnabledChanged
//go:noescape
func CallOffAlphaEnabledChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OffAlphaEnabledChanged
//go:noescape
func TryOffAlphaEnabledChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_HasOnAlphaEnabledChanged
//go:noescape
func HasFuncHasOnAlphaEnabledChanged() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_HasOnAlphaEnabledChanged
//go:noescape
func FuncHasOnAlphaEnabledChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_HasOnAlphaEnabledChanged
//go:noescape
func CallHasOnAlphaEnabledChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_HasOnAlphaEnabledChanged
//go:noescape
func TryHasOnAlphaEnabledChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OnBoundsChanged
//go:noescape
func HasFuncOnBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OnBoundsChanged
//go:noescape
func FuncOnBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OnBoundsChanged
//go:noescape
func CallOnBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OnBoundsChanged
//go:noescape
func TryOnBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OffBoundsChanged
//go:noescape
func HasFuncOffBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OffBoundsChanged
//go:noescape
func FuncOffBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OffBoundsChanged
//go:noescape
func CallOffBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OffBoundsChanged
//go:noescape
func TryOffBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_HasOnBoundsChanged
//go:noescape
func HasFuncHasOnBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_HasOnBoundsChanged
//go:noescape
func FuncHasOnBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_HasOnBoundsChanged
//go:noescape
func CallHasOnBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_HasOnBoundsChanged
//go:noescape
func TryHasOnBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OnClosed
//go:noescape
func HasFuncOnClosed() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OnClosed
//go:noescape
func FuncOnClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OnClosed
//go:noescape
func CallOnClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OnClosed
//go:noescape
func TryOnClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OffClosed
//go:noescape
func HasFuncOffClosed() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OffClosed
//go:noescape
func FuncOffClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OffClosed
//go:noescape
func CallOffClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OffClosed
//go:noescape
func TryOffClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_HasOnClosed
//go:noescape
func HasFuncHasOnClosed() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_HasOnClosed
//go:noescape
func FuncHasOnClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_HasOnClosed
//go:noescape
func CallHasOnClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_HasOnClosed
//go:noescape
func TryHasOnClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OnFullscreened
//go:noescape
func HasFuncOnFullscreened() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OnFullscreened
//go:noescape
func FuncOnFullscreened(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OnFullscreened
//go:noescape
func CallOnFullscreened(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OnFullscreened
//go:noescape
func TryOnFullscreened(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OffFullscreened
//go:noescape
func HasFuncOffFullscreened() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OffFullscreened
//go:noescape
func FuncOffFullscreened(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OffFullscreened
//go:noescape
func CallOffFullscreened(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OffFullscreened
//go:noescape
func TryOffFullscreened(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_HasOnFullscreened
//go:noescape
func HasFuncHasOnFullscreened() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_HasOnFullscreened
//go:noescape
func FuncHasOnFullscreened(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_HasOnFullscreened
//go:noescape
func CallHasOnFullscreened(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_HasOnFullscreened
//go:noescape
func TryHasOnFullscreened(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OnMaximized
//go:noescape
func HasFuncOnMaximized() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OnMaximized
//go:noescape
func FuncOnMaximized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OnMaximized
//go:noescape
func CallOnMaximized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OnMaximized
//go:noescape
func TryOnMaximized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OffMaximized
//go:noescape
func HasFuncOffMaximized() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OffMaximized
//go:noescape
func FuncOffMaximized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OffMaximized
//go:noescape
func CallOffMaximized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OffMaximized
//go:noescape
func TryOffMaximized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_HasOnMaximized
//go:noescape
func HasFuncHasOnMaximized() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_HasOnMaximized
//go:noescape
func FuncHasOnMaximized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_HasOnMaximized
//go:noescape
func CallHasOnMaximized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_HasOnMaximized
//go:noescape
func TryHasOnMaximized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OnMinimized
//go:noescape
func HasFuncOnMinimized() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OnMinimized
//go:noescape
func FuncOnMinimized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OnMinimized
//go:noescape
func CallOnMinimized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OnMinimized
//go:noescape
func TryOnMinimized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OffMinimized
//go:noescape
func HasFuncOffMinimized() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OffMinimized
//go:noescape
func FuncOffMinimized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OffMinimized
//go:noescape
func CallOffMinimized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OffMinimized
//go:noescape
func TryOffMinimized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_HasOnMinimized
//go:noescape
func HasFuncHasOnMinimized() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_HasOnMinimized
//go:noescape
func FuncHasOnMinimized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_HasOnMinimized
//go:noescape
func CallHasOnMinimized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_HasOnMinimized
//go:noescape
func TryHasOnMinimized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OnRestored
//go:noescape
func HasFuncOnRestored() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OnRestored
//go:noescape
func FuncOnRestored(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OnRestored
//go:noescape
func CallOnRestored(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OnRestored
//go:noescape
func TryOnRestored(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_OffRestored
//go:noescape
func HasFuncOffRestored() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_OffRestored
//go:noescape
func FuncOffRestored(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_OffRestored
//go:noescape
func CallOffRestored(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_OffRestored
//go:noescape
func TryOffRestored(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_HasOnRestored
//go:noescape
func HasFuncHasOnRestored() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_HasOnRestored
//go:noescape
func FuncHasOnRestored(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_HasOnRestored
//go:noescape
func CallHasOnRestored(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_HasOnRestored
//go:noescape
func TryHasOnRestored(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_Restore
//go:noescape
func HasFuncRestore() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_Restore
//go:noescape
func FuncRestore(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_Restore
//go:noescape
func CallRestore(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_Restore
//go:noescape
func TryRestore(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_SetActivateOnPointer
//go:noescape
func HasFuncSetActivateOnPointer() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_SetActivateOnPointer
//go:noescape
func FuncSetActivateOnPointer(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_SetActivateOnPointer
//go:noescape
func CallSetActivateOnPointer(
	retPtr unsafe.Pointer,
	activate_on_pointer js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_SetActivateOnPointer
//go:noescape
func TrySetActivateOnPointer(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	activate_on_pointer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_SetAlwaysOnTop
//go:noescape
func HasFuncSetAlwaysOnTop() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_SetAlwaysOnTop
//go:noescape
func FuncSetAlwaysOnTop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_SetAlwaysOnTop
//go:noescape
func CallSetAlwaysOnTop(
	retPtr unsafe.Pointer,
	always_on_top js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_SetAlwaysOnTop
//go:noescape
func TrySetAlwaysOnTop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	always_on_top js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_SetBounds
//go:noescape
func HasFuncSetBounds() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_SetBounds
//go:noescape
func FuncSetBounds(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_SetBounds
//go:noescape
func CallSetBounds(
	retPtr unsafe.Pointer,
	boundsType js.Ref,
	bounds unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_SetBounds
//go:noescape
func TrySetBounds(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	boundsType js.Ref,
	bounds unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_SetIcon
//go:noescape
func HasFuncSetIcon() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_SetIcon
//go:noescape
func FuncSetIcon(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_SetIcon
//go:noescape
func CallSetIcon(
	retPtr unsafe.Pointer,
	icon_url js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_SetIcon
//go:noescape
func TrySetIcon(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	icon_url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_SetShape
//go:noescape
func HasFuncSetShape() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_SetShape
//go:noescape
func FuncSetShape(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_SetShape
//go:noescape
func CallSetShape(
	retPtr unsafe.Pointer,
	region unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_SetShape
//go:noescape
func TrySetShape(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	region unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_SetSizeConstraints
//go:noescape
func HasFuncSetSizeConstraints() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_SetSizeConstraints
//go:noescape
func FuncSetSizeConstraints(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_SetSizeConstraints
//go:noescape
func CallSetSizeConstraints(
	retPtr unsafe.Pointer,
	boundsType js.Ref,
	constraints unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_SetSizeConstraints
//go:noescape
func TrySetSizeConstraints(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	boundsType js.Ref,
	constraints unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_SetVisibleOnAllWorkspaces
//go:noescape
func HasFuncSetVisibleOnAllWorkspaces() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_SetVisibleOnAllWorkspaces
//go:noescape
func FuncSetVisibleOnAllWorkspaces(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_SetVisibleOnAllWorkspaces
//go:noescape
func CallSetVisibleOnAllWorkspaces(
	retPtr unsafe.Pointer,
	always_visible js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_SetVisibleOnAllWorkspaces
//go:noescape
func TrySetVisibleOnAllWorkspaces(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	always_visible js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal has_Show
//go:noescape
func HasFuncShow() js.Ref

//go:wasmimport plat/js/webext/app/currentwindowinternal func_Show
//go:noescape
func FuncShow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/currentwindowinternal call_Show
//go:noescape
func CallShow(
	retPtr unsafe.Pointer,
	focused js.Ref)

//go:wasmimport plat/js/webext/app/currentwindowinternal try_Show
//go:noescape
func TryShow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	focused js.Ref) (ok js.Ref)

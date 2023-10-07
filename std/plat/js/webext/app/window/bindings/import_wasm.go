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

//go:wasmimport plat/js/webext/app/window store_ContentBounds
//go:noescape
func ContentBoundsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app/window load_ContentBounds
//go:noescape
func ContentBoundsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window get_Bounds_Left
//go:noescape
func GetBoundsLeft(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_Bounds_Left
//go:noescape
func SetBoundsLeft(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_Bounds_Top
//go:noescape
func GetBoundsTop(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_Bounds_Top
//go:noescape
func SetBoundsTop(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_Bounds_Width
//go:noescape
func GetBoundsWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_Bounds_Width
//go:noescape
func SetBoundsWidth(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_Bounds_Height
//go:noescape
func GetBoundsHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_Bounds_Height
//go:noescape
func SetBoundsHeight(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_Bounds_MinWidth
//go:noescape
func GetBoundsMinWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_Bounds_MinWidth
//go:noescape
func SetBoundsMinWidth(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_Bounds_MinHeight
//go:noescape
func GetBoundsMinHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_Bounds_MinHeight
//go:noescape
func SetBoundsMinHeight(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_Bounds_MaxWidth
//go:noescape
func GetBoundsMaxWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_Bounds_MaxWidth
//go:noescape
func SetBoundsMaxWidth(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_Bounds_MaxHeight
//go:noescape
func GetBoundsMaxHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_Bounds_MaxHeight
//go:noescape
func SetBoundsMaxHeight(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/app/window has_Bounds_SetPosition
//go:noescape
func HasFuncBoundsSetPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_Bounds_SetPosition
//go:noescape
func FuncBoundsSetPosition(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_Bounds_SetPosition
//go:noescape
func CallBoundsSetPosition(
	this js.Ref, retPtr unsafe.Pointer,
	left int32,
	top int32)

//go:wasmimport plat/js/webext/app/window try_Bounds_SetPosition
//go:noescape
func TryBoundsSetPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	left int32,
	top int32) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_Bounds_SetSize
//go:noescape
func HasFuncBoundsSetSize(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_Bounds_SetSize
//go:noescape
func FuncBoundsSetSize(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_Bounds_SetSize
//go:noescape
func CallBoundsSetSize(
	this js.Ref, retPtr unsafe.Pointer,
	width int32,
	height int32)

//go:wasmimport plat/js/webext/app/window try_Bounds_SetSize
//go:noescape
func TryBoundsSetSize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	width int32,
	height int32) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_Bounds_SetMinimumSize
//go:noescape
func HasFuncBoundsSetMinimumSize(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_Bounds_SetMinimumSize
//go:noescape
func FuncBoundsSetMinimumSize(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_Bounds_SetMinimumSize
//go:noescape
func CallBoundsSetMinimumSize(
	this js.Ref, retPtr unsafe.Pointer,
	minWidth int32,
	minHeight int32)

//go:wasmimport plat/js/webext/app/window try_Bounds_SetMinimumSize
//go:noescape
func TryBoundsSetMinimumSize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	minWidth int32,
	minHeight int32) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_Bounds_SetMaximumSize
//go:noescape
func HasFuncBoundsSetMaximumSize(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_Bounds_SetMaximumSize
//go:noescape
func FuncBoundsSetMaximumSize(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_Bounds_SetMaximumSize
//go:noescape
func CallBoundsSetMaximumSize(
	this js.Ref, retPtr unsafe.Pointer,
	maxWidth int32,
	maxHeight int32)

//go:wasmimport plat/js/webext/app/window try_Bounds_SetMaximumSize
//go:noescape
func TryBoundsSetMaximumSize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	maxWidth int32,
	maxHeight int32) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window get_AppWindow_HasFrameColor
//go:noescape
func GetAppWindowHasFrameColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_AppWindow_HasFrameColor
//go:noescape
func SetAppWindowHasFrameColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_AppWindow_ActiveFrameColor
//go:noescape
func GetAppWindowActiveFrameColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_AppWindow_ActiveFrameColor
//go:noescape
func SetAppWindowActiveFrameColor(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_AppWindow_InactiveFrameColor
//go:noescape
func GetAppWindowInactiveFrameColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_AppWindow_InactiveFrameColor
//go:noescape
func SetAppWindowInactiveFrameColor(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_AppWindow_ContentWindow
//go:noescape
func GetAppWindowContentWindow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_AppWindow_ContentWindow
//go:noescape
func SetAppWindowContentWindow(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_AppWindow_Id
//go:noescape
func GetAppWindowId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_AppWindow_Id
//go:noescape
func SetAppWindowId(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_AppWindow_InnerBounds
//go:noescape
func GetAppWindowInnerBounds(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_AppWindow_InnerBounds
//go:noescape
func SetAppWindowInnerBounds(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/app/window get_AppWindow_OuterBounds
//go:noescape
func GetAppWindowOuterBounds(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window set_AppWindow_OuterBounds
//go:noescape
func SetAppWindowOuterBounds(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/webext/app/window has_AppWindow_Focus
//go:noescape
func HasFuncAppWindowFocus(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_Focus
//go:noescape
func FuncAppWindowFocus(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_Focus
//go:noescape
func CallAppWindowFocus(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_Focus
//go:noescape
func TryAppWindowFocus(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_Fullscreen
//go:noescape
func HasFuncAppWindowFullscreen(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_Fullscreen
//go:noescape
func FuncAppWindowFullscreen(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_Fullscreen
//go:noescape
func CallAppWindowFullscreen(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_Fullscreen
//go:noescape
func TryAppWindowFullscreen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_IsFullscreen
//go:noescape
func HasFuncAppWindowIsFullscreen(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_IsFullscreen
//go:noescape
func FuncAppWindowIsFullscreen(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_IsFullscreen
//go:noescape
func CallAppWindowIsFullscreen(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_IsFullscreen
//go:noescape
func TryAppWindowIsFullscreen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_Minimize
//go:noescape
func HasFuncAppWindowMinimize(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_Minimize
//go:noescape
func FuncAppWindowMinimize(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_Minimize
//go:noescape
func CallAppWindowMinimize(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_Minimize
//go:noescape
func TryAppWindowMinimize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_IsMinimized
//go:noescape
func HasFuncAppWindowIsMinimized(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_IsMinimized
//go:noescape
func FuncAppWindowIsMinimized(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_IsMinimized
//go:noescape
func CallAppWindowIsMinimized(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_IsMinimized
//go:noescape
func TryAppWindowIsMinimized(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_Maximize
//go:noescape
func HasFuncAppWindowMaximize(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_Maximize
//go:noescape
func FuncAppWindowMaximize(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_Maximize
//go:noescape
func CallAppWindowMaximize(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_Maximize
//go:noescape
func TryAppWindowMaximize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_IsMaximized
//go:noescape
func HasFuncAppWindowIsMaximized(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_IsMaximized
//go:noescape
func FuncAppWindowIsMaximized(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_IsMaximized
//go:noescape
func CallAppWindowIsMaximized(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_IsMaximized
//go:noescape
func TryAppWindowIsMaximized(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_Restore
//go:noescape
func HasFuncAppWindowRestore(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_Restore
//go:noescape
func FuncAppWindowRestore(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_Restore
//go:noescape
func CallAppWindowRestore(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_Restore
//go:noescape
func TryAppWindowRestore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_MoveTo
//go:noescape
func HasFuncAppWindowMoveTo(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_MoveTo
//go:noescape
func FuncAppWindowMoveTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_MoveTo
//go:noescape
func CallAppWindowMoveTo(
	this js.Ref, retPtr unsafe.Pointer,
	left int32,
	top int32)

//go:wasmimport plat/js/webext/app/window try_AppWindow_MoveTo
//go:noescape
func TryAppWindowMoveTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	left int32,
	top int32) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_ResizeTo
//go:noescape
func HasFuncAppWindowResizeTo(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_ResizeTo
//go:noescape
func FuncAppWindowResizeTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_ResizeTo
//go:noescape
func CallAppWindowResizeTo(
	this js.Ref, retPtr unsafe.Pointer,
	width int32,
	height int32)

//go:wasmimport plat/js/webext/app/window try_AppWindow_ResizeTo
//go:noescape
func TryAppWindowResizeTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	width int32,
	height int32) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_DrawAttention
//go:noescape
func HasFuncAppWindowDrawAttention(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_DrawAttention
//go:noescape
func FuncAppWindowDrawAttention(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_DrawAttention
//go:noescape
func CallAppWindowDrawAttention(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_DrawAttention
//go:noescape
func TryAppWindowDrawAttention(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_ClearAttention
//go:noescape
func HasFuncAppWindowClearAttention(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_ClearAttention
//go:noescape
func FuncAppWindowClearAttention(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_ClearAttention
//go:noescape
func CallAppWindowClearAttention(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_ClearAttention
//go:noescape
func TryAppWindowClearAttention(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_Close
//go:noescape
func HasFuncAppWindowClose(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_Close
//go:noescape
func FuncAppWindowClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_Close
//go:noescape
func CallAppWindowClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_Close
//go:noescape
func TryAppWindowClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_Show
//go:noescape
func HasFuncAppWindowShow(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_Show
//go:noescape
func FuncAppWindowShow(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_Show
//go:noescape
func CallAppWindowShow(
	this js.Ref, retPtr unsafe.Pointer,
	focused js.Ref)

//go:wasmimport plat/js/webext/app/window try_AppWindow_Show
//go:noescape
func TryAppWindowShow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	focused js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_Show1
//go:noescape
func HasFuncAppWindowShow1(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_Show1
//go:noescape
func FuncAppWindowShow1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_Show1
//go:noescape
func CallAppWindowShow1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_Show1
//go:noescape
func TryAppWindowShow1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_Hide
//go:noescape
func HasFuncAppWindowHide(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_Hide
//go:noescape
func FuncAppWindowHide(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_Hide
//go:noescape
func CallAppWindowHide(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_Hide
//go:noescape
func TryAppWindowHide(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_GetBounds
//go:noescape
func HasFuncAppWindowGetBounds(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_GetBounds
//go:noescape
func FuncAppWindowGetBounds(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_GetBounds
//go:noescape
func CallAppWindowGetBounds(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_GetBounds
//go:noescape
func TryAppWindowGetBounds(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_SetBounds
//go:noescape
func HasFuncAppWindowSetBounds(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_SetBounds
//go:noescape
func FuncAppWindowSetBounds(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_SetBounds
//go:noescape
func CallAppWindowSetBounds(
	this js.Ref, retPtr unsafe.Pointer,
	bounds unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_SetBounds
//go:noescape
func TryAppWindowSetBounds(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bounds unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_SetIcon
//go:noescape
func HasFuncAppWindowSetIcon(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_SetIcon
//go:noescape
func FuncAppWindowSetIcon(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_SetIcon
//go:noescape
func CallAppWindowSetIcon(
	this js.Ref, retPtr unsafe.Pointer,
	iconUrl js.Ref)

//go:wasmimport plat/js/webext/app/window try_AppWindow_SetIcon
//go:noescape
func TryAppWindowSetIcon(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	iconUrl js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_IsAlwaysOnTop
//go:noescape
func HasFuncAppWindowIsAlwaysOnTop(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_IsAlwaysOnTop
//go:noescape
func FuncAppWindowIsAlwaysOnTop(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_IsAlwaysOnTop
//go:noescape
func CallAppWindowIsAlwaysOnTop(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_IsAlwaysOnTop
//go:noescape
func TryAppWindowIsAlwaysOnTop(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_SetAlwaysOnTop
//go:noescape
func HasFuncAppWindowSetAlwaysOnTop(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_SetAlwaysOnTop
//go:noescape
func FuncAppWindowSetAlwaysOnTop(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_SetAlwaysOnTop
//go:noescape
func CallAppWindowSetAlwaysOnTop(
	this js.Ref, retPtr unsafe.Pointer,
	alwaysOnTop js.Ref)

//go:wasmimport plat/js/webext/app/window try_AppWindow_SetAlwaysOnTop
//go:noescape
func TryAppWindowSetAlwaysOnTop(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	alwaysOnTop js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_AlphaEnabled
//go:noescape
func HasFuncAppWindowAlphaEnabled(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_AlphaEnabled
//go:noescape
func FuncAppWindowAlphaEnabled(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_AlphaEnabled
//go:noescape
func CallAppWindowAlphaEnabled(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_AppWindow_AlphaEnabled
//go:noescape
func TryAppWindowAlphaEnabled(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_AppWindow_SetVisibleOnAllWorkspaces
//go:noescape
func HasFuncAppWindowSetVisibleOnAllWorkspaces(this js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window func_AppWindow_SetVisibleOnAllWorkspaces
//go:noescape
func FuncAppWindowSetVisibleOnAllWorkspaces(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_AppWindow_SetVisibleOnAllWorkspaces
//go:noescape
func CallAppWindowSetVisibleOnAllWorkspaces(
	this js.Ref, retPtr unsafe.Pointer,
	alwaysVisible js.Ref)

//go:wasmimport plat/js/webext/app/window try_AppWindow_SetVisibleOnAllWorkspaces
//go:noescape
func TryAppWindowSetVisibleOnAllWorkspaces(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	alwaysVisible js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window store_BoundsSpecification
//go:noescape
func BoundsSpecificationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app/window load_BoundsSpecification
//go:noescape
func BoundsSpecificationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window constof_WindowType
//go:noescape
func ConstOfWindowType(str js.Ref) uint32

//go:wasmimport plat/js/webext/app/window store_FrameOptions
//go:noescape
func FrameOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app/window load_FrameOptions
//go:noescape
func FrameOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window constof_State
//go:noescape
func ConstOfState(str js.Ref) uint32

//go:wasmimport plat/js/webext/app/window store_CreateWindowOptions
//go:noescape
func CreateWindowOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/app/window load_CreateWindowOptions
//go:noescape
func CreateWindowOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/app/window has_CanSetVisibleOnAllWorkspaces
//go:noescape
func HasFuncCanSetVisibleOnAllWorkspaces() js.Ref

//go:wasmimport plat/js/webext/app/window func_CanSetVisibleOnAllWorkspaces
//go:noescape
func FuncCanSetVisibleOnAllWorkspaces(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_CanSetVisibleOnAllWorkspaces
//go:noescape
func CallCanSetVisibleOnAllWorkspaces(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_CanSetVisibleOnAllWorkspaces
//go:noescape
func TryCanSetVisibleOnAllWorkspaces(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/app/window func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	url js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_Current
//go:noescape
func HasFuncCurrent() js.Ref

//go:wasmimport plat/js/webext/app/window func_Current
//go:noescape
func FuncCurrent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_Current
//go:noescape
func CallCurrent(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_Current
//go:noescape
func TryCurrent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/app/window func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/app/window try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_GetAll
//go:noescape
func HasFuncGetAll() js.Ref

//go:wasmimport plat/js/webext/app/window func_GetAll
//go:noescape
func FuncGetAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_GetAll
//go:noescape
func CallGetAll(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window try_GetAll
//go:noescape
func TryGetAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_InitializeAppWindow
//go:noescape
func HasFuncInitializeAppWindow() js.Ref

//go:wasmimport plat/js/webext/app/window func_InitializeAppWindow
//go:noescape
func FuncInitializeAppWindow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_InitializeAppWindow
//go:noescape
func CallInitializeAppWindow(
	retPtr unsafe.Pointer,
	state js.Ref)

//go:wasmimport plat/js/webext/app/window try_InitializeAppWindow
//go:noescape
func TryInitializeAppWindow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	state js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OnAlphaEnabledChanged
//go:noescape
func HasFuncOnAlphaEnabledChanged() js.Ref

//go:wasmimport plat/js/webext/app/window func_OnAlphaEnabledChanged
//go:noescape
func FuncOnAlphaEnabledChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OnAlphaEnabledChanged
//go:noescape
func CallOnAlphaEnabledChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OnAlphaEnabledChanged
//go:noescape
func TryOnAlphaEnabledChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OffAlphaEnabledChanged
//go:noescape
func HasFuncOffAlphaEnabledChanged() js.Ref

//go:wasmimport plat/js/webext/app/window func_OffAlphaEnabledChanged
//go:noescape
func FuncOffAlphaEnabledChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OffAlphaEnabledChanged
//go:noescape
func CallOffAlphaEnabledChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OffAlphaEnabledChanged
//go:noescape
func TryOffAlphaEnabledChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_HasOnAlphaEnabledChanged
//go:noescape
func HasFuncHasOnAlphaEnabledChanged() js.Ref

//go:wasmimport plat/js/webext/app/window func_HasOnAlphaEnabledChanged
//go:noescape
func FuncHasOnAlphaEnabledChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_HasOnAlphaEnabledChanged
//go:noescape
func CallHasOnAlphaEnabledChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_HasOnAlphaEnabledChanged
//go:noescape
func TryHasOnAlphaEnabledChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OnBoundsChanged
//go:noescape
func HasFuncOnBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/app/window func_OnBoundsChanged
//go:noescape
func FuncOnBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OnBoundsChanged
//go:noescape
func CallOnBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OnBoundsChanged
//go:noescape
func TryOnBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OffBoundsChanged
//go:noescape
func HasFuncOffBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/app/window func_OffBoundsChanged
//go:noescape
func FuncOffBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OffBoundsChanged
//go:noescape
func CallOffBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OffBoundsChanged
//go:noescape
func TryOffBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_HasOnBoundsChanged
//go:noescape
func HasFuncHasOnBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/app/window func_HasOnBoundsChanged
//go:noescape
func FuncHasOnBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_HasOnBoundsChanged
//go:noescape
func CallHasOnBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_HasOnBoundsChanged
//go:noescape
func TryHasOnBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OnClosed
//go:noescape
func HasFuncOnClosed() js.Ref

//go:wasmimport plat/js/webext/app/window func_OnClosed
//go:noescape
func FuncOnClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OnClosed
//go:noescape
func CallOnClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OnClosed
//go:noescape
func TryOnClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OffClosed
//go:noescape
func HasFuncOffClosed() js.Ref

//go:wasmimport plat/js/webext/app/window func_OffClosed
//go:noescape
func FuncOffClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OffClosed
//go:noescape
func CallOffClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OffClosed
//go:noescape
func TryOffClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_HasOnClosed
//go:noescape
func HasFuncHasOnClosed() js.Ref

//go:wasmimport plat/js/webext/app/window func_HasOnClosed
//go:noescape
func FuncHasOnClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_HasOnClosed
//go:noescape
func CallHasOnClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_HasOnClosed
//go:noescape
func TryHasOnClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OnFullscreened
//go:noescape
func HasFuncOnFullscreened() js.Ref

//go:wasmimport plat/js/webext/app/window func_OnFullscreened
//go:noescape
func FuncOnFullscreened(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OnFullscreened
//go:noescape
func CallOnFullscreened(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OnFullscreened
//go:noescape
func TryOnFullscreened(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OffFullscreened
//go:noescape
func HasFuncOffFullscreened() js.Ref

//go:wasmimport plat/js/webext/app/window func_OffFullscreened
//go:noescape
func FuncOffFullscreened(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OffFullscreened
//go:noescape
func CallOffFullscreened(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OffFullscreened
//go:noescape
func TryOffFullscreened(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_HasOnFullscreened
//go:noescape
func HasFuncHasOnFullscreened() js.Ref

//go:wasmimport plat/js/webext/app/window func_HasOnFullscreened
//go:noescape
func FuncHasOnFullscreened(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_HasOnFullscreened
//go:noescape
func CallHasOnFullscreened(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_HasOnFullscreened
//go:noescape
func TryHasOnFullscreened(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OnMaximized
//go:noescape
func HasFuncOnMaximized() js.Ref

//go:wasmimport plat/js/webext/app/window func_OnMaximized
//go:noescape
func FuncOnMaximized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OnMaximized
//go:noescape
func CallOnMaximized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OnMaximized
//go:noescape
func TryOnMaximized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OffMaximized
//go:noescape
func HasFuncOffMaximized() js.Ref

//go:wasmimport plat/js/webext/app/window func_OffMaximized
//go:noescape
func FuncOffMaximized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OffMaximized
//go:noescape
func CallOffMaximized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OffMaximized
//go:noescape
func TryOffMaximized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_HasOnMaximized
//go:noescape
func HasFuncHasOnMaximized() js.Ref

//go:wasmimport plat/js/webext/app/window func_HasOnMaximized
//go:noescape
func FuncHasOnMaximized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_HasOnMaximized
//go:noescape
func CallHasOnMaximized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_HasOnMaximized
//go:noescape
func TryHasOnMaximized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OnMinimized
//go:noescape
func HasFuncOnMinimized() js.Ref

//go:wasmimport plat/js/webext/app/window func_OnMinimized
//go:noescape
func FuncOnMinimized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OnMinimized
//go:noescape
func CallOnMinimized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OnMinimized
//go:noescape
func TryOnMinimized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OffMinimized
//go:noescape
func HasFuncOffMinimized() js.Ref

//go:wasmimport plat/js/webext/app/window func_OffMinimized
//go:noescape
func FuncOffMinimized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OffMinimized
//go:noescape
func CallOffMinimized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OffMinimized
//go:noescape
func TryOffMinimized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_HasOnMinimized
//go:noescape
func HasFuncHasOnMinimized() js.Ref

//go:wasmimport plat/js/webext/app/window func_HasOnMinimized
//go:noescape
func FuncHasOnMinimized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_HasOnMinimized
//go:noescape
func CallHasOnMinimized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_HasOnMinimized
//go:noescape
func TryHasOnMinimized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OnRestored
//go:noescape
func HasFuncOnRestored() js.Ref

//go:wasmimport plat/js/webext/app/window func_OnRestored
//go:noescape
func FuncOnRestored(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OnRestored
//go:noescape
func CallOnRestored(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OnRestored
//go:noescape
func TryOnRestored(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_OffRestored
//go:noescape
func HasFuncOffRestored() js.Ref

//go:wasmimport plat/js/webext/app/window func_OffRestored
//go:noescape
func FuncOffRestored(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_OffRestored
//go:noescape
func CallOffRestored(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_OffRestored
//go:noescape
func TryOffRestored(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/app/window has_HasOnRestored
//go:noescape
func HasFuncHasOnRestored() js.Ref

//go:wasmimport plat/js/webext/app/window func_HasOnRestored
//go:noescape
func FuncHasOnRestored(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/app/window call_HasOnRestored
//go:noescape
func CallHasOnRestored(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/app/window try_HasOnRestored
//go:noescape
func TryHasOnRestored(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

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

//go:wasmimport plat/js/webext/virtualkeyboardprivate store_Bounds
//go:noescape
func BoundsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate load_Bounds
//go:noescape
func BoundsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate constof_DisplayFormat
//go:noescape
func ConstOfDisplayFormat(str js.Ref) uint32

//go:wasmimport plat/js/webext/virtualkeyboardprivate store_ClipboardItem
//go:noescape
func ClipboardItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate load_ClipboardItem
//go:noescape
func ClipboardItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate constof_KeyboardMode
//go:noescape
func ConstOfKeyboardMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/virtualkeyboardprivate store_ContainerBehaviorOptions
//go:noescape
func ContainerBehaviorOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate load_ContainerBehaviorOptions
//go:noescape
func ContainerBehaviorOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate store_GetClipboardHistoryArgOptions
//go:noescape
func GetClipboardHistoryArgOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate load_GetClipboardHistoryArgOptions
//go:noescape
func GetClipboardHistoryArgOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate store_KeyboardConfig
//go:noescape
func KeyboardConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate load_KeyboardConfig
//go:noescape
func KeyboardConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate constof_KeyboardState
//go:noescape
func ConstOfKeyboardState(str js.Ref) uint32

//go:wasmimport plat/js/webext/virtualkeyboardprivate constof_VirtualKeyboardEventType
//go:noescape
func ConstOfVirtualKeyboardEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/virtualkeyboardprivate store_VirtualKeyboardEvent
//go:noescape
func VirtualKeyboardEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate load_VirtualKeyboardEvent
//go:noescape
func VirtualKeyboardEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_DeleteClipboardItem
//go:noescape
func HasFuncDeleteClipboardItem() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_DeleteClipboardItem
//go:noescape
func FuncDeleteClipboardItem(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_DeleteClipboardItem
//go:noescape
func CallDeleteClipboardItem(
	retPtr unsafe.Pointer,
	itemId js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_DeleteClipboardItem
//go:noescape
func TryDeleteClipboardItem(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	itemId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_GetClipboardHistory
//go:noescape
func HasFuncGetClipboardHistory() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_GetClipboardHistory
//go:noescape
func FuncGetClipboardHistory(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_GetClipboardHistory
//go:noescape
func CallGetClipboardHistory(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_GetClipboardHistory
//go:noescape
func TryGetClipboardHistory(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_GetKeyboardConfig
//go:noescape
func HasFuncGetKeyboardConfig() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_GetKeyboardConfig
//go:noescape
func FuncGetKeyboardConfig(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_GetKeyboardConfig
//go:noescape
func CallGetKeyboardConfig(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_GetKeyboardConfig
//go:noescape
func TryGetKeyboardConfig(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_HideKeyboard
//go:noescape
func HasFuncHideKeyboard() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_HideKeyboard
//go:noescape
func FuncHideKeyboard(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_HideKeyboard
//go:noescape
func CallHideKeyboard(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_HideKeyboard
//go:noescape
func TryHideKeyboard(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_InsertText
//go:noescape
func HasFuncInsertText() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_InsertText
//go:noescape
func FuncInsertText(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_InsertText
//go:noescape
func CallInsertText(
	retPtr unsafe.Pointer,
	text js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_InsertText
//go:noescape
func TryInsertText(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	text js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_KeyboardLoaded
//go:noescape
func HasFuncKeyboardLoaded() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_KeyboardLoaded
//go:noescape
func FuncKeyboardLoaded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_KeyboardLoaded
//go:noescape
func CallKeyboardLoaded(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_KeyboardLoaded
//go:noescape
func TryKeyboardLoaded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_LockKeyboard
//go:noescape
func HasFuncLockKeyboard() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_LockKeyboard
//go:noescape
func FuncLockKeyboard(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_LockKeyboard
//go:noescape
func CallLockKeyboard(
	retPtr unsafe.Pointer,
	lock js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_LockKeyboard
//go:noescape
func TryLockKeyboard(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	lock js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OnBoundsChanged
//go:noescape
func HasFuncOnBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OnBoundsChanged
//go:noescape
func FuncOnBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OnBoundsChanged
//go:noescape
func CallOnBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OnBoundsChanged
//go:noescape
func TryOnBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OffBoundsChanged
//go:noescape
func HasFuncOffBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OffBoundsChanged
//go:noescape
func FuncOffBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OffBoundsChanged
//go:noescape
func CallOffBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OffBoundsChanged
//go:noescape
func TryOffBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_HasOnBoundsChanged
//go:noescape
func HasFuncHasOnBoundsChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_HasOnBoundsChanged
//go:noescape
func FuncHasOnBoundsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_HasOnBoundsChanged
//go:noescape
func CallHasOnBoundsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_HasOnBoundsChanged
//go:noescape
func TryHasOnBoundsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OnClipboardHistoryChanged
//go:noescape
func HasFuncOnClipboardHistoryChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OnClipboardHistoryChanged
//go:noescape
func FuncOnClipboardHistoryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OnClipboardHistoryChanged
//go:noescape
func CallOnClipboardHistoryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OnClipboardHistoryChanged
//go:noescape
func TryOnClipboardHistoryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OffClipboardHistoryChanged
//go:noescape
func HasFuncOffClipboardHistoryChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OffClipboardHistoryChanged
//go:noescape
func FuncOffClipboardHistoryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OffClipboardHistoryChanged
//go:noescape
func CallOffClipboardHistoryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OffClipboardHistoryChanged
//go:noescape
func TryOffClipboardHistoryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_HasOnClipboardHistoryChanged
//go:noescape
func HasFuncHasOnClipboardHistoryChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_HasOnClipboardHistoryChanged
//go:noescape
func FuncHasOnClipboardHistoryChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_HasOnClipboardHistoryChanged
//go:noescape
func CallHasOnClipboardHistoryChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_HasOnClipboardHistoryChanged
//go:noescape
func TryHasOnClipboardHistoryChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OnClipboardItemUpdated
//go:noescape
func HasFuncOnClipboardItemUpdated() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OnClipboardItemUpdated
//go:noescape
func FuncOnClipboardItemUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OnClipboardItemUpdated
//go:noescape
func CallOnClipboardItemUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OnClipboardItemUpdated
//go:noescape
func TryOnClipboardItemUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OffClipboardItemUpdated
//go:noescape
func HasFuncOffClipboardItemUpdated() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OffClipboardItemUpdated
//go:noescape
func FuncOffClipboardItemUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OffClipboardItemUpdated
//go:noescape
func CallOffClipboardItemUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OffClipboardItemUpdated
//go:noescape
func TryOffClipboardItemUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_HasOnClipboardItemUpdated
//go:noescape
func HasFuncHasOnClipboardItemUpdated() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_HasOnClipboardItemUpdated
//go:noescape
func FuncHasOnClipboardItemUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_HasOnClipboardItemUpdated
//go:noescape
func CallHasOnClipboardItemUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_HasOnClipboardItemUpdated
//go:noescape
func TryHasOnClipboardItemUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OnColorProviderChanged
//go:noescape
func HasFuncOnColorProviderChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OnColorProviderChanged
//go:noescape
func FuncOnColorProviderChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OnColorProviderChanged
//go:noescape
func CallOnColorProviderChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OnColorProviderChanged
//go:noescape
func TryOnColorProviderChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OffColorProviderChanged
//go:noescape
func HasFuncOffColorProviderChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OffColorProviderChanged
//go:noescape
func FuncOffColorProviderChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OffColorProviderChanged
//go:noescape
func CallOffColorProviderChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OffColorProviderChanged
//go:noescape
func TryOffColorProviderChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_HasOnColorProviderChanged
//go:noescape
func HasFuncHasOnColorProviderChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_HasOnColorProviderChanged
//go:noescape
func FuncHasOnColorProviderChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_HasOnColorProviderChanged
//go:noescape
func CallHasOnColorProviderChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_HasOnColorProviderChanged
//go:noescape
func TryHasOnColorProviderChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OnKeyboardClosed
//go:noescape
func HasFuncOnKeyboardClosed() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OnKeyboardClosed
//go:noescape
func FuncOnKeyboardClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OnKeyboardClosed
//go:noescape
func CallOnKeyboardClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OnKeyboardClosed
//go:noescape
func TryOnKeyboardClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OffKeyboardClosed
//go:noescape
func HasFuncOffKeyboardClosed() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OffKeyboardClosed
//go:noescape
func FuncOffKeyboardClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OffKeyboardClosed
//go:noescape
func CallOffKeyboardClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OffKeyboardClosed
//go:noescape
func TryOffKeyboardClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_HasOnKeyboardClosed
//go:noescape
func HasFuncHasOnKeyboardClosed() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_HasOnKeyboardClosed
//go:noescape
func FuncHasOnKeyboardClosed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_HasOnKeyboardClosed
//go:noescape
func CallHasOnKeyboardClosed(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_HasOnKeyboardClosed
//go:noescape
func TryHasOnKeyboardClosed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OnKeyboardConfigChanged
//go:noescape
func HasFuncOnKeyboardConfigChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OnKeyboardConfigChanged
//go:noescape
func FuncOnKeyboardConfigChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OnKeyboardConfigChanged
//go:noescape
func CallOnKeyboardConfigChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OnKeyboardConfigChanged
//go:noescape
func TryOnKeyboardConfigChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OffKeyboardConfigChanged
//go:noescape
func HasFuncOffKeyboardConfigChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OffKeyboardConfigChanged
//go:noescape
func FuncOffKeyboardConfigChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OffKeyboardConfigChanged
//go:noescape
func CallOffKeyboardConfigChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OffKeyboardConfigChanged
//go:noescape
func TryOffKeyboardConfigChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_HasOnKeyboardConfigChanged
//go:noescape
func HasFuncHasOnKeyboardConfigChanged() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_HasOnKeyboardConfigChanged
//go:noescape
func FuncHasOnKeyboardConfigChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_HasOnKeyboardConfigChanged
//go:noescape
func CallHasOnKeyboardConfigChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_HasOnKeyboardConfigChanged
//go:noescape
func TryHasOnKeyboardConfigChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OpenSettings
//go:noescape
func HasFuncOpenSettings() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OpenSettings
//go:noescape
func FuncOpenSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OpenSettings
//go:noescape
func CallOpenSettings(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OpenSettings
//go:noescape
func TryOpenSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_OpenSuggestionSettings
//go:noescape
func HasFuncOpenSuggestionSettings() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_OpenSuggestionSettings
//go:noescape
func FuncOpenSuggestionSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_OpenSuggestionSettings
//go:noescape
func CallOpenSuggestionSettings(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_OpenSuggestionSettings
//go:noescape
func TryOpenSuggestionSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_PasteClipboardItem
//go:noescape
func HasFuncPasteClipboardItem() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_PasteClipboardItem
//go:noescape
func FuncPasteClipboardItem(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_PasteClipboardItem
//go:noescape
func CallPasteClipboardItem(
	retPtr unsafe.Pointer,
	itemId js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_PasteClipboardItem
//go:noescape
func TryPasteClipboardItem(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	itemId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_SendKeyEvent
//go:noescape
func HasFuncSendKeyEvent() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_SendKeyEvent
//go:noescape
func FuncSendKeyEvent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_SendKeyEvent
//go:noescape
func CallSendKeyEvent(
	retPtr unsafe.Pointer,
	keyEvent unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_SendKeyEvent
//go:noescape
func TrySendKeyEvent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keyEvent unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_SetAreaToRemainOnScreen
//go:noescape
func HasFuncSetAreaToRemainOnScreen() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_SetAreaToRemainOnScreen
//go:noescape
func FuncSetAreaToRemainOnScreen(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_SetAreaToRemainOnScreen
//go:noescape
func CallSetAreaToRemainOnScreen(
	retPtr unsafe.Pointer,
	bounds unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_SetAreaToRemainOnScreen
//go:noescape
func TrySetAreaToRemainOnScreen(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bounds unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_SetContainerBehavior
//go:noescape
func HasFuncSetContainerBehavior() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_SetContainerBehavior
//go:noescape
func FuncSetContainerBehavior(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_SetContainerBehavior
//go:noescape
func CallSetContainerBehavior(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_SetContainerBehavior
//go:noescape
func TrySetContainerBehavior(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_SetDraggableArea
//go:noescape
func HasFuncSetDraggableArea() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_SetDraggableArea
//go:noescape
func FuncSetDraggableArea(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_SetDraggableArea
//go:noescape
func CallSetDraggableArea(
	retPtr unsafe.Pointer,
	bounds unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_SetDraggableArea
//go:noescape
func TrySetDraggableArea(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bounds unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_SetHitTestBounds
//go:noescape
func HasFuncSetHitTestBounds() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_SetHitTestBounds
//go:noescape
func FuncSetHitTestBounds(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_SetHitTestBounds
//go:noescape
func CallSetHitTestBounds(
	retPtr unsafe.Pointer,
	boundsList js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_SetHitTestBounds
//go:noescape
func TrySetHitTestBounds(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	boundsList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_SetHotrodKeyboard
//go:noescape
func HasFuncSetHotrodKeyboard() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_SetHotrodKeyboard
//go:noescape
func FuncSetHotrodKeyboard(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_SetHotrodKeyboard
//go:noescape
func CallSetHotrodKeyboard(
	retPtr unsafe.Pointer,
	enable js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_SetHotrodKeyboard
//go:noescape
func TrySetHotrodKeyboard(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enable js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_SetKeyboardState
//go:noescape
func HasFuncSetKeyboardState() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_SetKeyboardState
//go:noescape
func FuncSetKeyboardState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_SetKeyboardState
//go:noescape
func CallSetKeyboardState(
	retPtr unsafe.Pointer,
	state uint32)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_SetKeyboardState
//go:noescape
func TrySetKeyboardState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	state uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_SetOccludedBounds
//go:noescape
func HasFuncSetOccludedBounds() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_SetOccludedBounds
//go:noescape
func FuncSetOccludedBounds(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_SetOccludedBounds
//go:noescape
func CallSetOccludedBounds(
	retPtr unsafe.Pointer,
	boundsList js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_SetOccludedBounds
//go:noescape
func TrySetOccludedBounds(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	boundsList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/virtualkeyboardprivate has_SetWindowBoundsInScreen
//go:noescape
func HasFuncSetWindowBoundsInScreen() js.Ref

//go:wasmimport plat/js/webext/virtualkeyboardprivate func_SetWindowBoundsInScreen
//go:noescape
func FuncSetWindowBoundsInScreen(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate call_SetWindowBoundsInScreen
//go:noescape
func CallSetWindowBoundsInScreen(
	retPtr unsafe.Pointer,
	bounds unsafe.Pointer)

//go:wasmimport plat/js/webext/virtualkeyboardprivate try_SetWindowBoundsInScreen
//go:noescape
func TrySetWindowBoundsInScreen(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bounds unsafe.Pointer) (ok js.Ref)

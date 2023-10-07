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

//go:wasmimport plat/js/webext/tabs store_ConnectArgConnectInfo
//go:noescape
func ConnectArgConnectInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_ConnectArgConnectInfo
//go:noescape
func ConnectArgConnectInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs constof_MutedInfoReason
//go:noescape
func ConstOfMutedInfoReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/tabs store_MutedInfo
//go:noescape
func MutedInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_MutedInfo
//go:noescape
func MutedInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs constof_TabStatus
//go:noescape
func ConstOfTabStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/tabs store_Tab
//go:noescape
func TabJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_Tab
//go:noescape
func TabJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_CreateArgCreateProperties
//go:noescape
func CreateArgCreatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_CreateArgCreateProperties
//go:noescape
func CreateArgCreatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_GroupArgOptionsFieldCreateProperties
//go:noescape
func GroupArgOptionsFieldCreatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_GroupArgOptionsFieldCreateProperties
//go:noescape
func GroupArgOptionsFieldCreatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_GroupArgOptions
//go:noescape
func GroupArgOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_GroupArgOptions
//go:noescape
func GroupArgOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_HighlightArgHighlightInfo
//go:noescape
func HighlightArgHighlightInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_HighlightArgHighlightInfo
//go:noescape
func HighlightArgHighlightInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs get_MAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND
//go:noescape
func GetMAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/tabs set_MAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND
//go:noescape
func SetMAX_CAPTURE_VISIBLE_TAB_CALLS_PER_SECOND(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_MoveArgMoveProperties
//go:noescape
func MoveArgMovePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_MoveArgMoveProperties
//go:noescape
func MoveArgMovePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_OnActivatedArgActiveInfo
//go:noescape
func OnActivatedArgActiveInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_OnActivatedArgActiveInfo
//go:noescape
func OnActivatedArgActiveInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_OnActiveChangedArgSelectInfo
//go:noescape
func OnActiveChangedArgSelectInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_OnActiveChangedArgSelectInfo
//go:noescape
func OnActiveChangedArgSelectInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_OnAttachedArgAttachInfo
//go:noescape
func OnAttachedArgAttachInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_OnAttachedArgAttachInfo
//go:noescape
func OnAttachedArgAttachInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_OnDetachedArgDetachInfo
//go:noescape
func OnDetachedArgDetachInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_OnDetachedArgDetachInfo
//go:noescape
func OnDetachedArgDetachInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_OnHighlightChangedArgSelectInfo
//go:noescape
func OnHighlightChangedArgSelectInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_OnHighlightChangedArgSelectInfo
//go:noescape
func OnHighlightChangedArgSelectInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_OnHighlightedArgHighlightInfo
//go:noescape
func OnHighlightedArgHighlightInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_OnHighlightedArgHighlightInfo
//go:noescape
func OnHighlightedArgHighlightInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_OnMovedArgMoveInfo
//go:noescape
func OnMovedArgMoveInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_OnMovedArgMoveInfo
//go:noescape
func OnMovedArgMoveInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_OnRemovedArgRemoveInfo
//go:noescape
func OnRemovedArgRemoveInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_OnRemovedArgRemoveInfo
//go:noescape
func OnRemovedArgRemoveInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_OnSelectionChangedArgSelectInfo
//go:noescape
func OnSelectionChangedArgSelectInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_OnSelectionChangedArgSelectInfo
//go:noescape
func OnSelectionChangedArgSelectInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_OnUpdatedArgChangeInfo
//go:noescape
func OnUpdatedArgChangeInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_OnUpdatedArgChangeInfo
//go:noescape
func OnUpdatedArgChangeInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs constof_ZoomSettingsMode
//go:noescape
func ConstOfZoomSettingsMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/tabs constof_ZoomSettingsScope
//go:noescape
func ConstOfZoomSettingsScope(str js.Ref) uint32

//go:wasmimport plat/js/webext/tabs store_ZoomSettings
//go:noescape
func ZoomSettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_ZoomSettings
//go:noescape
func ZoomSettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_OnZoomChangeArgZoomChangeInfo
//go:noescape
func OnZoomChangeArgZoomChangeInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_OnZoomChangeArgZoomChangeInfo
//go:noescape
func OnZoomChangeArgZoomChangeInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs constof_WindowType
//go:noescape
func ConstOfWindowType(str js.Ref) uint32

//go:wasmimport plat/js/webext/tabs store_QueryArgQueryInfo
//go:noescape
func QueryArgQueryInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_QueryArgQueryInfo
//go:noescape
func QueryArgQueryInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_ReloadArgReloadProperties
//go:noescape
func ReloadArgReloadPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_ReloadArgReloadProperties
//go:noescape
func ReloadArgReloadPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_SendMessageArgOptions
//go:noescape
func SendMessageArgOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_SendMessageArgOptions
//go:noescape
func SendMessageArgOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs get_TAB_ID_NONE
//go:noescape
func GetTAB_ID_NONE(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/tabs set_TAB_ID_NONE
//go:noescape
func SetTAB_ID_NONE(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs store_UpdateArgUpdateProperties
//go:noescape
func UpdateArgUpdatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabs load_UpdateArgUpdateProperties
//go:noescape
func UpdateArgUpdatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabs has_CaptureVisibleTab
//go:noescape
func HasFuncCaptureVisibleTab() js.Ref

//go:wasmimport plat/js/webext/tabs func_CaptureVisibleTab
//go:noescape
func FuncCaptureVisibleTab(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_CaptureVisibleTab
//go:noescape
func CallCaptureVisibleTab(
	retPtr unsafe.Pointer,
	windowId float64,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_CaptureVisibleTab
//go:noescape
func TryCaptureVisibleTab(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	windowId float64,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/tabs func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	createProperties unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	createProperties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_DetectLanguage
//go:noescape
func HasFuncDetectLanguage() js.Ref

//go:wasmimport plat/js/webext/tabs func_DetectLanguage
//go:noescape
func FuncDetectLanguage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_DetectLanguage
//go:noescape
func CallDetectLanguage(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/tabs try_DetectLanguage
//go:noescape
func TryDetectLanguage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_Discard
//go:noescape
func HasFuncDiscard() js.Ref

//go:wasmimport plat/js/webext/tabs func_Discard
//go:noescape
func FuncDiscard(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_Discard
//go:noescape
func CallDiscard(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/tabs try_Discard
//go:noescape
func TryDiscard(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_Duplicate
//go:noescape
func HasFuncDuplicate() js.Ref

//go:wasmimport plat/js/webext/tabs func_Duplicate
//go:noescape
func FuncDuplicate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_Duplicate
//go:noescape
func CallDuplicate(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/tabs try_Duplicate
//go:noescape
func TryDuplicate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_ExecuteScript
//go:noescape
func HasFuncExecuteScript() js.Ref

//go:wasmimport plat/js/webext/tabs func_ExecuteScript
//go:noescape
func FuncExecuteScript(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_ExecuteScript
//go:noescape
func CallExecuteScript(
	retPtr unsafe.Pointer,
	tabId float64,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_ExecuteScript
//go:noescape
func TryExecuteScript(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/tabs func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/tabs try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_GetAllInWindow
//go:noescape
func HasFuncGetAllInWindow() js.Ref

//go:wasmimport plat/js/webext/tabs func_GetAllInWindow
//go:noescape
func FuncGetAllInWindow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_GetAllInWindow
//go:noescape
func CallGetAllInWindow(
	retPtr unsafe.Pointer,
	windowId float64)

//go:wasmimport plat/js/webext/tabs try_GetAllInWindow
//go:noescape
func TryGetAllInWindow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	windowId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_GetCurrent
//go:noescape
func HasFuncGetCurrent() js.Ref

//go:wasmimport plat/js/webext/tabs func_GetCurrent
//go:noescape
func FuncGetCurrent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_GetCurrent
//go:noescape
func CallGetCurrent(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_GetCurrent
//go:noescape
func TryGetCurrent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_GetSelected
//go:noescape
func HasFuncGetSelected() js.Ref

//go:wasmimport plat/js/webext/tabs func_GetSelected
//go:noescape
func FuncGetSelected(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_GetSelected
//go:noescape
func CallGetSelected(
	retPtr unsafe.Pointer,
	windowId float64)

//go:wasmimport plat/js/webext/tabs try_GetSelected
//go:noescape
func TryGetSelected(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	windowId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_GetZoom
//go:noescape
func HasFuncGetZoom() js.Ref

//go:wasmimport plat/js/webext/tabs func_GetZoom
//go:noescape
func FuncGetZoom(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_GetZoom
//go:noescape
func CallGetZoom(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/tabs try_GetZoom
//go:noescape
func TryGetZoom(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_GetZoomSettings
//go:noescape
func HasFuncGetZoomSettings() js.Ref

//go:wasmimport plat/js/webext/tabs func_GetZoomSettings
//go:noescape
func FuncGetZoomSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_GetZoomSettings
//go:noescape
func CallGetZoomSettings(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/tabs try_GetZoomSettings
//go:noescape
func TryGetZoomSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_GoBack
//go:noescape
func HasFuncGoBack() js.Ref

//go:wasmimport plat/js/webext/tabs func_GoBack
//go:noescape
func FuncGoBack(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_GoBack
//go:noescape
func CallGoBack(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/tabs try_GoBack
//go:noescape
func TryGoBack(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_GoForward
//go:noescape
func HasFuncGoForward() js.Ref

//go:wasmimport plat/js/webext/tabs func_GoForward
//go:noescape
func FuncGoForward(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_GoForward
//go:noescape
func CallGoForward(
	retPtr unsafe.Pointer,
	tabId float64)

//go:wasmimport plat/js/webext/tabs try_GoForward
//go:noescape
func TryGoForward(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_Group
//go:noescape
func HasFuncGroup() js.Ref

//go:wasmimport plat/js/webext/tabs func_Group
//go:noescape
func FuncGroup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_Group
//go:noescape
func CallGroup(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_Group
//go:noescape
func TryGroup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_InsertCSS
//go:noescape
func HasFuncInsertCSS() js.Ref

//go:wasmimport plat/js/webext/tabs func_InsertCSS
//go:noescape
func FuncInsertCSS(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_InsertCSS
//go:noescape
func CallInsertCSS(
	retPtr unsafe.Pointer,
	tabId float64,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_InsertCSS
//go:noescape
func TryInsertCSS(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_Move
//go:noescape
func HasFuncMove() js.Ref

//go:wasmimport plat/js/webext/tabs func_Move
//go:noescape
func FuncMove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_Move
//go:noescape
func CallMove(
	retPtr unsafe.Pointer,
	tabIds js.Ref,
	moveProperties unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_Move
//go:noescape
func TryMove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabIds js.Ref,
	moveProperties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnActivated
//go:noescape
func HasFuncOnActivated() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnActivated
//go:noescape
func FuncOnActivated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnActivated
//go:noescape
func CallOnActivated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnActivated
//go:noescape
func TryOnActivated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffActivated
//go:noescape
func HasFuncOffActivated() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffActivated
//go:noescape
func FuncOffActivated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffActivated
//go:noescape
func CallOffActivated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffActivated
//go:noescape
func TryOffActivated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnActivated
//go:noescape
func HasFuncHasOnActivated() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnActivated
//go:noescape
func FuncHasOnActivated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnActivated
//go:noescape
func CallHasOnActivated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnActivated
//go:noescape
func TryHasOnActivated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnActiveChanged
//go:noescape
func HasFuncOnActiveChanged() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnActiveChanged
//go:noescape
func FuncOnActiveChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnActiveChanged
//go:noescape
func CallOnActiveChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnActiveChanged
//go:noescape
func TryOnActiveChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffActiveChanged
//go:noescape
func HasFuncOffActiveChanged() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffActiveChanged
//go:noescape
func FuncOffActiveChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffActiveChanged
//go:noescape
func CallOffActiveChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffActiveChanged
//go:noescape
func TryOffActiveChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnActiveChanged
//go:noescape
func HasFuncHasOnActiveChanged() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnActiveChanged
//go:noescape
func FuncHasOnActiveChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnActiveChanged
//go:noescape
func CallHasOnActiveChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnActiveChanged
//go:noescape
func TryHasOnActiveChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnAttached
//go:noescape
func HasFuncOnAttached() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnAttached
//go:noescape
func FuncOnAttached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnAttached
//go:noescape
func CallOnAttached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnAttached
//go:noescape
func TryOnAttached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffAttached
//go:noescape
func HasFuncOffAttached() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffAttached
//go:noescape
func FuncOffAttached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffAttached
//go:noescape
func CallOffAttached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffAttached
//go:noescape
func TryOffAttached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnAttached
//go:noescape
func HasFuncHasOnAttached() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnAttached
//go:noescape
func FuncHasOnAttached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnAttached
//go:noescape
func CallHasOnAttached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnAttached
//go:noescape
func TryHasOnAttached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnCreated
//go:noescape
func HasFuncOnCreated() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnCreated
//go:noescape
func FuncOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnCreated
//go:noescape
func CallOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnCreated
//go:noescape
func TryOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffCreated
//go:noescape
func HasFuncOffCreated() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffCreated
//go:noescape
func FuncOffCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffCreated
//go:noescape
func CallOffCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffCreated
//go:noescape
func TryOffCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnCreated
//go:noescape
func HasFuncHasOnCreated() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnCreated
//go:noescape
func FuncHasOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnCreated
//go:noescape
func CallHasOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnCreated
//go:noescape
func TryHasOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnDetached
//go:noescape
func HasFuncOnDetached() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnDetached
//go:noescape
func FuncOnDetached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnDetached
//go:noescape
func CallOnDetached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnDetached
//go:noescape
func TryOnDetached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffDetached
//go:noescape
func HasFuncOffDetached() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffDetached
//go:noescape
func FuncOffDetached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffDetached
//go:noescape
func CallOffDetached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffDetached
//go:noescape
func TryOffDetached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnDetached
//go:noescape
func HasFuncHasOnDetached() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnDetached
//go:noescape
func FuncHasOnDetached(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnDetached
//go:noescape
func CallHasOnDetached(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnDetached
//go:noescape
func TryHasOnDetached(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnHighlightChanged
//go:noescape
func HasFuncOnHighlightChanged() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnHighlightChanged
//go:noescape
func FuncOnHighlightChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnHighlightChanged
//go:noescape
func CallOnHighlightChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnHighlightChanged
//go:noescape
func TryOnHighlightChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffHighlightChanged
//go:noescape
func HasFuncOffHighlightChanged() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffHighlightChanged
//go:noescape
func FuncOffHighlightChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffHighlightChanged
//go:noescape
func CallOffHighlightChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffHighlightChanged
//go:noescape
func TryOffHighlightChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnHighlightChanged
//go:noescape
func HasFuncHasOnHighlightChanged() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnHighlightChanged
//go:noescape
func FuncHasOnHighlightChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnHighlightChanged
//go:noescape
func CallHasOnHighlightChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnHighlightChanged
//go:noescape
func TryHasOnHighlightChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnHighlighted
//go:noescape
func HasFuncOnHighlighted() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnHighlighted
//go:noescape
func FuncOnHighlighted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnHighlighted
//go:noescape
func CallOnHighlighted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnHighlighted
//go:noescape
func TryOnHighlighted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffHighlighted
//go:noescape
func HasFuncOffHighlighted() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffHighlighted
//go:noescape
func FuncOffHighlighted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffHighlighted
//go:noescape
func CallOffHighlighted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffHighlighted
//go:noescape
func TryOffHighlighted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnHighlighted
//go:noescape
func HasFuncHasOnHighlighted() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnHighlighted
//go:noescape
func FuncHasOnHighlighted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnHighlighted
//go:noescape
func CallHasOnHighlighted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnHighlighted
//go:noescape
func TryHasOnHighlighted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnMoved
//go:noescape
func HasFuncOnMoved() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnMoved
//go:noescape
func FuncOnMoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnMoved
//go:noescape
func CallOnMoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnMoved
//go:noescape
func TryOnMoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffMoved
//go:noescape
func HasFuncOffMoved() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffMoved
//go:noescape
func FuncOffMoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffMoved
//go:noescape
func CallOffMoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffMoved
//go:noescape
func TryOffMoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnMoved
//go:noescape
func HasFuncHasOnMoved() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnMoved
//go:noescape
func FuncHasOnMoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnMoved
//go:noescape
func CallHasOnMoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnMoved
//go:noescape
func TryHasOnMoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnRemoved
//go:noescape
func HasFuncOnRemoved() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnRemoved
//go:noescape
func FuncOnRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnRemoved
//go:noescape
func CallOnRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnRemoved
//go:noescape
func TryOnRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffRemoved
//go:noescape
func HasFuncOffRemoved() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffRemoved
//go:noescape
func FuncOffRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffRemoved
//go:noescape
func CallOffRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffRemoved
//go:noescape
func TryOffRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnRemoved
//go:noescape
func HasFuncHasOnRemoved() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnRemoved
//go:noescape
func FuncHasOnRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnRemoved
//go:noescape
func CallHasOnRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnRemoved
//go:noescape
func TryHasOnRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnReplaced
//go:noescape
func HasFuncOnReplaced() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnReplaced
//go:noescape
func FuncOnReplaced(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnReplaced
//go:noescape
func CallOnReplaced(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnReplaced
//go:noescape
func TryOnReplaced(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffReplaced
//go:noescape
func HasFuncOffReplaced() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffReplaced
//go:noescape
func FuncOffReplaced(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffReplaced
//go:noescape
func CallOffReplaced(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffReplaced
//go:noescape
func TryOffReplaced(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnReplaced
//go:noescape
func HasFuncHasOnReplaced() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnReplaced
//go:noescape
func FuncHasOnReplaced(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnReplaced
//go:noescape
func CallHasOnReplaced(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnReplaced
//go:noescape
func TryHasOnReplaced(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnSelectionChanged
//go:noescape
func HasFuncOnSelectionChanged() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnSelectionChanged
//go:noescape
func FuncOnSelectionChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnSelectionChanged
//go:noescape
func CallOnSelectionChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnSelectionChanged
//go:noescape
func TryOnSelectionChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffSelectionChanged
//go:noescape
func HasFuncOffSelectionChanged() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffSelectionChanged
//go:noescape
func FuncOffSelectionChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffSelectionChanged
//go:noescape
func CallOffSelectionChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffSelectionChanged
//go:noescape
func TryOffSelectionChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnSelectionChanged
//go:noescape
func HasFuncHasOnSelectionChanged() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnSelectionChanged
//go:noescape
func FuncHasOnSelectionChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnSelectionChanged
//go:noescape
func CallHasOnSelectionChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnSelectionChanged
//go:noescape
func TryHasOnSelectionChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnUpdated
//go:noescape
func HasFuncOnUpdated() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnUpdated
//go:noescape
func FuncOnUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnUpdated
//go:noescape
func CallOnUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnUpdated
//go:noescape
func TryOnUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffUpdated
//go:noescape
func HasFuncOffUpdated() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffUpdated
//go:noescape
func FuncOffUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffUpdated
//go:noescape
func CallOffUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffUpdated
//go:noescape
func TryOffUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnUpdated
//go:noescape
func HasFuncHasOnUpdated() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnUpdated
//go:noescape
func FuncHasOnUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnUpdated
//go:noescape
func CallHasOnUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnUpdated
//go:noescape
func TryHasOnUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OnZoomChange
//go:noescape
func HasFuncOnZoomChange() js.Ref

//go:wasmimport plat/js/webext/tabs func_OnZoomChange
//go:noescape
func FuncOnZoomChange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OnZoomChange
//go:noescape
func CallOnZoomChange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OnZoomChange
//go:noescape
func TryOnZoomChange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_OffZoomChange
//go:noescape
func HasFuncOffZoomChange() js.Ref

//go:wasmimport plat/js/webext/tabs func_OffZoomChange
//go:noescape
func FuncOffZoomChange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_OffZoomChange
//go:noescape
func CallOffZoomChange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_OffZoomChange
//go:noescape
func TryOffZoomChange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_HasOnZoomChange
//go:noescape
func HasFuncHasOnZoomChange() js.Ref

//go:wasmimport plat/js/webext/tabs func_HasOnZoomChange
//go:noescape
func FuncHasOnZoomChange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_HasOnZoomChange
//go:noescape
func CallHasOnZoomChange(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabs try_HasOnZoomChange
//go:noescape
func TryHasOnZoomChange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_Query
//go:noescape
func HasFuncQuery() js.Ref

//go:wasmimport plat/js/webext/tabs func_Query
//go:noescape
func FuncQuery(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_Query
//go:noescape
func CallQuery(
	retPtr unsafe.Pointer,
	queryInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_Query
//go:noescape
func TryQuery(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	queryInfo unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_Reload
//go:noescape
func HasFuncReload() js.Ref

//go:wasmimport plat/js/webext/tabs func_Reload
//go:noescape
func FuncReload(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_Reload
//go:noescape
func CallReload(
	retPtr unsafe.Pointer,
	tabId float64,
	reloadProperties unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_Reload
//go:noescape
func TryReload(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64,
	reloadProperties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_Remove
//go:noescape
func HasFuncRemove() js.Ref

//go:wasmimport plat/js/webext/tabs func_Remove
//go:noescape
func FuncRemove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_Remove
//go:noescape
func CallRemove(
	retPtr unsafe.Pointer,
	tabIds js.Ref)

//go:wasmimport plat/js/webext/tabs try_Remove
//go:noescape
func TryRemove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabIds js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_RemoveCSS
//go:noescape
func HasFuncRemoveCSS() js.Ref

//go:wasmimport plat/js/webext/tabs func_RemoveCSS
//go:noescape
func FuncRemoveCSS(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_RemoveCSS
//go:noescape
func CallRemoveCSS(
	retPtr unsafe.Pointer,
	tabId float64,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_RemoveCSS
//go:noescape
func TryRemoveCSS(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_SendMessage
//go:noescape
func HasFuncSendMessage() js.Ref

//go:wasmimport plat/js/webext/tabs func_SendMessage
//go:noescape
func FuncSendMessage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_SendMessage
//go:noescape
func CallSendMessage(
	retPtr unsafe.Pointer,
	tabId float64,
	message js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_SendMessage
//go:noescape
func TrySendMessage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64,
	message js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_SendRequest
//go:noescape
func HasFuncSendRequest() js.Ref

//go:wasmimport plat/js/webext/tabs func_SendRequest
//go:noescape
func FuncSendRequest(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_SendRequest
//go:noescape
func CallSendRequest(
	retPtr unsafe.Pointer,
	tabId float64,
	request js.Ref)

//go:wasmimport plat/js/webext/tabs try_SendRequest
//go:noescape
func TrySendRequest(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64,
	request js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_SetZoom
//go:noescape
func HasFuncSetZoom() js.Ref

//go:wasmimport plat/js/webext/tabs func_SetZoom
//go:noescape
func FuncSetZoom(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_SetZoom
//go:noescape
func CallSetZoom(
	retPtr unsafe.Pointer,
	tabId float64,
	zoomFactor float64)

//go:wasmimport plat/js/webext/tabs try_SetZoom
//go:noescape
func TrySetZoom(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64,
	zoomFactor float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_SetZoomSettings
//go:noescape
func HasFuncSetZoomSettings() js.Ref

//go:wasmimport plat/js/webext/tabs func_SetZoomSettings
//go:noescape
func FuncSetZoomSettings(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_SetZoomSettings
//go:noescape
func CallSetZoomSettings(
	retPtr unsafe.Pointer,
	tabId float64,
	zoomSettings unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_SetZoomSettings
//go:noescape
func TrySetZoomSettings(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64,
	zoomSettings unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_Ungroup
//go:noescape
func HasFuncUngroup() js.Ref

//go:wasmimport plat/js/webext/tabs func_Ungroup
//go:noescape
func FuncUngroup(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_Ungroup
//go:noescape
func CallUngroup(
	retPtr unsafe.Pointer,
	tabIds js.Ref)

//go:wasmimport plat/js/webext/tabs try_Ungroup
//go:noescape
func TryUngroup(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabIds js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabs has_Update
//go:noescape
func HasFuncUpdate() js.Ref

//go:wasmimport plat/js/webext/tabs func_Update
//go:noescape
func FuncUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs call_Update
//go:noescape
func CallUpdate(
	retPtr unsafe.Pointer,
	tabId float64,
	updateProperties unsafe.Pointer)

//go:wasmimport plat/js/webext/tabs try_Update
//go:noescape
func TryUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tabId float64,
	updateProperties unsafe.Pointer) (ok js.Ref)

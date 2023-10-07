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

//go:wasmimport plat/js/webext/webviewinternal store_InjectionItems
//go:noescape
func InjectionItemsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewinternal load_InjectionItems
//go:noescape
func InjectionItemsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewinternal store_ContentScriptDetails
//go:noescape
func ContentScriptDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewinternal load_ContentScriptDetails
//go:noescape
func ContentScriptDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewinternal store_DataTypeSet
//go:noescape
func DataTypeSetJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewinternal load_DataTypeSet
//go:noescape
func DataTypeSetJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewinternal store_FindArgCallbackArgResultsFieldSelectionRect
//go:noescape
func FindArgCallbackArgResultsFieldSelectionRectJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewinternal load_FindArgCallbackArgResultsFieldSelectionRect
//go:noescape
func FindArgCallbackArgResultsFieldSelectionRectJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewinternal store_FindArgCallbackArgResults
//go:noescape
func FindArgCallbackArgResultsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewinternal load_FindArgCallbackArgResults
//go:noescape
func FindArgCallbackArgResultsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewinternal store_FindArgOptions
//go:noescape
func FindArgOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewinternal load_FindArgOptions
//go:noescape
func FindArgOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewinternal constof_ZoomMode
//go:noescape
func ConstOfZoomMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/webviewinternal get_MAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND
//go:noescape
func GetMAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/webviewinternal set_MAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND
//go:noescape
func SetMAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewinternal store_RemovalOptions
//go:noescape
func RemovalOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/webviewinternal load_RemovalOptions
//go:noescape
func RemovalOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/webviewinternal constof_SetPermissionAction
//go:noescape
func ConstOfSetPermissionAction(str js.Ref) uint32

//go:wasmimport plat/js/webext/webviewinternal constof_StopFindingAction
//go:noescape
func ConstOfStopFindingAction(str js.Ref) uint32

//go:wasmimport plat/js/webext/webviewinternal has_AddContentScripts
//go:noescape
func HasFuncAddContentScripts() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_AddContentScripts
//go:noescape
func FuncAddContentScripts(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_AddContentScripts
//go:noescape
func CallAddContentScripts(
	retPtr unsafe.Pointer,
	instanceId float64,
	contentScriptList js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_AddContentScripts
//go:noescape
func TryAddContentScripts(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	contentScriptList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_CaptureVisibleRegion
//go:noescape
func HasFuncCaptureVisibleRegion() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_CaptureVisibleRegion
//go:noescape
func FuncCaptureVisibleRegion(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_CaptureVisibleRegion
//go:noescape
func CallCaptureVisibleRegion(
	retPtr unsafe.Pointer,
	instanceId float64,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_CaptureVisibleRegion
//go:noescape
func TryCaptureVisibleRegion(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_ClearData
//go:noescape
func HasFuncClearData() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_ClearData
//go:noescape
func FuncClearData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_ClearData
//go:noescape
func CallClearData(
	retPtr unsafe.Pointer,
	instanceId float64,
	options unsafe.Pointer,
	dataToRemove unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_ClearData
//go:noescape
func TryClearData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	options unsafe.Pointer,
	dataToRemove unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_ExecuteScript
//go:noescape
func HasFuncExecuteScript() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_ExecuteScript
//go:noescape
func FuncExecuteScript(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_ExecuteScript
//go:noescape
func CallExecuteScript(
	retPtr unsafe.Pointer,
	instanceId float64,
	src js.Ref,
	details unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_ExecuteScript
//go:noescape
func TryExecuteScript(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	src js.Ref,
	details unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_Find
//go:noescape
func HasFuncFind() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_Find
//go:noescape
func FuncFind(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_Find
//go:noescape
func CallFind(
	retPtr unsafe.Pointer,
	instanceId float64,
	searchText js.Ref,
	options unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_Find
//go:noescape
func TryFind(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	searchText js.Ref,
	options unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_GetAudioState
//go:noescape
func HasFuncGetAudioState() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_GetAudioState
//go:noescape
func FuncGetAudioState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_GetAudioState
//go:noescape
func CallGetAudioState(
	retPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_GetAudioState
//go:noescape
func TryGetAudioState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_GetZoom
//go:noescape
func HasFuncGetZoom() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_GetZoom
//go:noescape
func FuncGetZoom(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_GetZoom
//go:noescape
func CallGetZoom(
	retPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_GetZoom
//go:noescape
func TryGetZoom(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_GetZoomMode
//go:noescape
func HasFuncGetZoomMode() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_GetZoomMode
//go:noescape
func FuncGetZoomMode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_GetZoomMode
//go:noescape
func CallGetZoomMode(
	retPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_GetZoomMode
//go:noescape
func TryGetZoomMode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_Go
//go:noescape
func HasFuncGo() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_Go
//go:noescape
func FuncGo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_Go
//go:noescape
func CallGo(
	retPtr unsafe.Pointer,
	instanceId float64,
	relativeIndex float64,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_Go
//go:noescape
func TryGo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	relativeIndex float64,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_InsertCSS
//go:noescape
func HasFuncInsertCSS() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_InsertCSS
//go:noescape
func FuncInsertCSS(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_InsertCSS
//go:noescape
func CallInsertCSS(
	retPtr unsafe.Pointer,
	instanceId float64,
	src js.Ref,
	details unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_InsertCSS
//go:noescape
func TryInsertCSS(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	src js.Ref,
	details unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_IsAudioMuted
//go:noescape
func HasFuncIsAudioMuted() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_IsAudioMuted
//go:noescape
func FuncIsAudioMuted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_IsAudioMuted
//go:noescape
func CallIsAudioMuted(
	retPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_IsAudioMuted
//go:noescape
func TryIsAudioMuted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_IsSpatialNavigationEnabled
//go:noescape
func HasFuncIsSpatialNavigationEnabled() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_IsSpatialNavigationEnabled
//go:noescape
func FuncIsSpatialNavigationEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_IsSpatialNavigationEnabled
//go:noescape
func CallIsSpatialNavigationEnabled(
	retPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_IsSpatialNavigationEnabled
//go:noescape
func TryIsSpatialNavigationEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_LoadDataWithBaseUrl
//go:noescape
func HasFuncLoadDataWithBaseUrl() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_LoadDataWithBaseUrl
//go:noescape
func FuncLoadDataWithBaseUrl(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_LoadDataWithBaseUrl
//go:noescape
func CallLoadDataWithBaseUrl(
	retPtr unsafe.Pointer,
	instanceId float64,
	dataUrl js.Ref,
	baseUrl js.Ref,
	virtualUrl js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_LoadDataWithBaseUrl
//go:noescape
func TryLoadDataWithBaseUrl(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	dataUrl js.Ref,
	baseUrl js.Ref,
	virtualUrl js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_Navigate
//go:noescape
func HasFuncNavigate() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_Navigate
//go:noescape
func FuncNavigate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_Navigate
//go:noescape
func CallNavigate(
	retPtr unsafe.Pointer,
	instanceId float64,
	src js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_Navigate
//go:noescape
func TryNavigate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	src js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_OverrideUserAgent
//go:noescape
func HasFuncOverrideUserAgent() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_OverrideUserAgent
//go:noescape
func FuncOverrideUserAgent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_OverrideUserAgent
//go:noescape
func CallOverrideUserAgent(
	retPtr unsafe.Pointer,
	instanceId float64,
	userAgentOverride js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_OverrideUserAgent
//go:noescape
func TryOverrideUserAgent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	userAgentOverride js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_Reload
//go:noescape
func HasFuncReload() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_Reload
//go:noescape
func FuncReload(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_Reload
//go:noescape
func CallReload(
	retPtr unsafe.Pointer,
	instanceId float64)

//go:wasmimport plat/js/webext/webviewinternal try_Reload
//go:noescape
func TryReload(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_RemoveContentScripts
//go:noescape
func HasFuncRemoveContentScripts() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_RemoveContentScripts
//go:noescape
func FuncRemoveContentScripts(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_RemoveContentScripts
//go:noescape
func CallRemoveContentScripts(
	retPtr unsafe.Pointer,
	instanceId float64,
	scriptNameList js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_RemoveContentScripts
//go:noescape
func TryRemoveContentScripts(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	scriptNameList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_SetAllowScaling
//go:noescape
func HasFuncSetAllowScaling() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_SetAllowScaling
//go:noescape
func FuncSetAllowScaling(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_SetAllowScaling
//go:noescape
func CallSetAllowScaling(
	retPtr unsafe.Pointer,
	instanceId float64,
	allow js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_SetAllowScaling
//go:noescape
func TrySetAllowScaling(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	allow js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_SetAllowTransparency
//go:noescape
func HasFuncSetAllowTransparency() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_SetAllowTransparency
//go:noescape
func FuncSetAllowTransparency(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_SetAllowTransparency
//go:noescape
func CallSetAllowTransparency(
	retPtr unsafe.Pointer,
	instanceId float64,
	allow js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_SetAllowTransparency
//go:noescape
func TrySetAllowTransparency(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	allow js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_SetAudioMuted
//go:noescape
func HasFuncSetAudioMuted() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_SetAudioMuted
//go:noescape
func FuncSetAudioMuted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_SetAudioMuted
//go:noescape
func CallSetAudioMuted(
	retPtr unsafe.Pointer,
	instanceId float64,
	mute js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_SetAudioMuted
//go:noescape
func TrySetAudioMuted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	mute js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_SetName
//go:noescape
func HasFuncSetName() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_SetName
//go:noescape
func FuncSetName(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_SetName
//go:noescape
func CallSetName(
	retPtr unsafe.Pointer,
	instanceId float64,
	frameName js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_SetName
//go:noescape
func TrySetName(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	frameName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_SetPermission
//go:noescape
func HasFuncSetPermission() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_SetPermission
//go:noescape
func FuncSetPermission(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_SetPermission
//go:noescape
func CallSetPermission(
	retPtr unsafe.Pointer,
	instanceId float64,
	requestId float64,
	action uint32,
	userInput js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_SetPermission
//go:noescape
func TrySetPermission(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	requestId float64,
	action uint32,
	userInput js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_SetSpatialNavigationEnabled
//go:noescape
func HasFuncSetSpatialNavigationEnabled() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_SetSpatialNavigationEnabled
//go:noescape
func FuncSetSpatialNavigationEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_SetSpatialNavigationEnabled
//go:noescape
func CallSetSpatialNavigationEnabled(
	retPtr unsafe.Pointer,
	instanceId float64,
	spatialNavEnabled js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_SetSpatialNavigationEnabled
//go:noescape
func TrySetSpatialNavigationEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	spatialNavEnabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_SetZoom
//go:noescape
func HasFuncSetZoom() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_SetZoom
//go:noescape
func FuncSetZoom(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_SetZoom
//go:noescape
func CallSetZoom(
	retPtr unsafe.Pointer,
	instanceId float64,
	zoomFactor float64,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_SetZoom
//go:noescape
func TrySetZoom(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	zoomFactor float64,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_SetZoomMode
//go:noescape
func HasFuncSetZoomMode() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_SetZoomMode
//go:noescape
func FuncSetZoomMode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_SetZoomMode
//go:noescape
func CallSetZoomMode(
	retPtr unsafe.Pointer,
	instanceId float64,
	ZoomMode uint32,
	callback js.Ref)

//go:wasmimport plat/js/webext/webviewinternal try_SetZoomMode
//go:noescape
func TrySetZoomMode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	ZoomMode uint32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_Stop
//go:noescape
func HasFuncStop() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_Stop
//go:noescape
func FuncStop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_Stop
//go:noescape
func CallStop(
	retPtr unsafe.Pointer,
	instanceId float64)

//go:wasmimport plat/js/webext/webviewinternal try_Stop
//go:noescape
func TryStop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_StopFinding
//go:noescape
func HasFuncStopFinding() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_StopFinding
//go:noescape
func FuncStopFinding(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_StopFinding
//go:noescape
func CallStopFinding(
	retPtr unsafe.Pointer,
	instanceId float64,
	action uint32)

//go:wasmimport plat/js/webext/webviewinternal try_StopFinding
//go:noescape
func TryStopFinding(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64,
	action uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/webviewinternal has_Terminate
//go:noescape
func HasFuncTerminate() js.Ref

//go:wasmimport plat/js/webext/webviewinternal func_Terminate
//go:noescape
func FuncTerminate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/webviewinternal call_Terminate
//go:noescape
func CallTerminate(
	retPtr unsafe.Pointer,
	instanceId float64)

//go:wasmimport plat/js/webext/webviewinternal try_Terminate
//go:noescape
func TryTerminate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	instanceId float64) (ok js.Ref)

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

//go:wasmimport plat/js/webext/system/display constof_ActiveState
//go:noescape
func ConstOfActiveState(str js.Ref) uint32

//go:wasmimport plat/js/webext/system/display store_Bounds
//go:noescape
func BoundsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_Bounds
//go:noescape
func BoundsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display store_Edid
//go:noescape
func EdidJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_Edid
//go:noescape
func EdidJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display store_Insets
//go:noescape
func InsetsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_Insets
//go:noescape
func InsetsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display store_DisplayMode
//go:noescape
func DisplayModeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_DisplayMode
//go:noescape
func DisplayModeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display store_DisplayUnitInfo
//go:noescape
func DisplayUnitInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_DisplayUnitInfo
//go:noescape
func DisplayUnitInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display constof_LayoutPosition
//go:noescape
func ConstOfLayoutPosition(str js.Ref) uint32

//go:wasmimport plat/js/webext/system/display store_DisplayLayout
//go:noescape
func DisplayLayoutJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_DisplayLayout
//go:noescape
func DisplayLayoutJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display store_DisplayProperties
//go:noescape
func DisplayPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_DisplayProperties
//go:noescape
func DisplayPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display store_GetInfoFlags
//go:noescape
func GetInfoFlagsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_GetInfoFlags
//go:noescape
func GetInfoFlagsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display constof_MirrorMode
//go:noescape
func ConstOfMirrorMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/system/display store_MirrorModeInfo
//go:noescape
func MirrorModeInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_MirrorModeInfo
//go:noescape
func MirrorModeInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display store_Point
//go:noescape
func PointJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_Point
//go:noescape
func PointJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display store_TouchCalibrationPair
//go:noescape
func TouchCalibrationPairJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_TouchCalibrationPair
//go:noescape
func TouchCalibrationPairJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display store_TouchCalibrationPairQuad
//go:noescape
func TouchCalibrationPairQuadJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/system/display load_TouchCalibrationPairQuad
//go:noescape
func TouchCalibrationPairQuadJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/system/display has_ClearTouchCalibration
//go:noescape
func HasFuncClearTouchCalibration() js.Ref

//go:wasmimport plat/js/webext/system/display func_ClearTouchCalibration
//go:noescape
func FuncClearTouchCalibration(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_ClearTouchCalibration
//go:noescape
func CallClearTouchCalibration(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/system/display try_ClearTouchCalibration
//go:noescape
func TryClearTouchCalibration(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_CompleteCustomTouchCalibration
//go:noescape
func HasFuncCompleteCustomTouchCalibration() js.Ref

//go:wasmimport plat/js/webext/system/display func_CompleteCustomTouchCalibration
//go:noescape
func FuncCompleteCustomTouchCalibration(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_CompleteCustomTouchCalibration
//go:noescape
func CallCompleteCustomTouchCalibration(
	retPtr unsafe.Pointer,
	pairs unsafe.Pointer,
	bounds unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display try_CompleteCustomTouchCalibration
//go:noescape
func TryCompleteCustomTouchCalibration(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	pairs unsafe.Pointer,
	bounds unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_EnableUnifiedDesktop
//go:noescape
func HasFuncEnableUnifiedDesktop() js.Ref

//go:wasmimport plat/js/webext/system/display func_EnableUnifiedDesktop
//go:noescape
func FuncEnableUnifiedDesktop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_EnableUnifiedDesktop
//go:noescape
func CallEnableUnifiedDesktop(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/system/display try_EnableUnifiedDesktop
//go:noescape
func TryEnableUnifiedDesktop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_GetDisplayLayout
//go:noescape
func HasFuncGetDisplayLayout() js.Ref

//go:wasmimport plat/js/webext/system/display func_GetDisplayLayout
//go:noescape
func FuncGetDisplayLayout(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_GetDisplayLayout
//go:noescape
func CallGetDisplayLayout(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display try_GetDisplayLayout
//go:noescape
func TryGetDisplayLayout(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_GetInfo
//go:noescape
func HasFuncGetInfo() js.Ref

//go:wasmimport plat/js/webext/system/display func_GetInfo
//go:noescape
func FuncGetInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_GetInfo
//go:noescape
func CallGetInfo(
	retPtr unsafe.Pointer,
	flags unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display try_GetInfo
//go:noescape
func TryGetInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	flags unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_OnDisplayChanged
//go:noescape
func HasFuncOnDisplayChanged() js.Ref

//go:wasmimport plat/js/webext/system/display func_OnDisplayChanged
//go:noescape
func FuncOnDisplayChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_OnDisplayChanged
//go:noescape
func CallOnDisplayChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/system/display try_OnDisplayChanged
//go:noescape
func TryOnDisplayChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_OffDisplayChanged
//go:noescape
func HasFuncOffDisplayChanged() js.Ref

//go:wasmimport plat/js/webext/system/display func_OffDisplayChanged
//go:noescape
func FuncOffDisplayChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_OffDisplayChanged
//go:noescape
func CallOffDisplayChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/system/display try_OffDisplayChanged
//go:noescape
func TryOffDisplayChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_HasOnDisplayChanged
//go:noescape
func HasFuncHasOnDisplayChanged() js.Ref

//go:wasmimport plat/js/webext/system/display func_HasOnDisplayChanged
//go:noescape
func FuncHasOnDisplayChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_HasOnDisplayChanged
//go:noescape
func CallHasOnDisplayChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/system/display try_HasOnDisplayChanged
//go:noescape
func TryHasOnDisplayChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_OverscanCalibrationAdjust
//go:noescape
func HasFuncOverscanCalibrationAdjust() js.Ref

//go:wasmimport plat/js/webext/system/display func_OverscanCalibrationAdjust
//go:noescape
func FuncOverscanCalibrationAdjust(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_OverscanCalibrationAdjust
//go:noescape
func CallOverscanCalibrationAdjust(
	retPtr unsafe.Pointer,
	id js.Ref,
	delta unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display try_OverscanCalibrationAdjust
//go:noescape
func TryOverscanCalibrationAdjust(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	delta unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_OverscanCalibrationComplete
//go:noescape
func HasFuncOverscanCalibrationComplete() js.Ref

//go:wasmimport plat/js/webext/system/display func_OverscanCalibrationComplete
//go:noescape
func FuncOverscanCalibrationComplete(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_OverscanCalibrationComplete
//go:noescape
func CallOverscanCalibrationComplete(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/system/display try_OverscanCalibrationComplete
//go:noescape
func TryOverscanCalibrationComplete(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_OverscanCalibrationReset
//go:noescape
func HasFuncOverscanCalibrationReset() js.Ref

//go:wasmimport plat/js/webext/system/display func_OverscanCalibrationReset
//go:noescape
func FuncOverscanCalibrationReset(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_OverscanCalibrationReset
//go:noescape
func CallOverscanCalibrationReset(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/system/display try_OverscanCalibrationReset
//go:noescape
func TryOverscanCalibrationReset(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_OverscanCalibrationStart
//go:noescape
func HasFuncOverscanCalibrationStart() js.Ref

//go:wasmimport plat/js/webext/system/display func_OverscanCalibrationStart
//go:noescape
func FuncOverscanCalibrationStart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_OverscanCalibrationStart
//go:noescape
func CallOverscanCalibrationStart(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/system/display try_OverscanCalibrationStart
//go:noescape
func TryOverscanCalibrationStart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_SetDisplayLayout
//go:noescape
func HasFuncSetDisplayLayout() js.Ref

//go:wasmimport plat/js/webext/system/display func_SetDisplayLayout
//go:noescape
func FuncSetDisplayLayout(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_SetDisplayLayout
//go:noescape
func CallSetDisplayLayout(
	retPtr unsafe.Pointer,
	layouts js.Ref)

//go:wasmimport plat/js/webext/system/display try_SetDisplayLayout
//go:noescape
func TrySetDisplayLayout(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	layouts js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_SetDisplayProperties
//go:noescape
func HasFuncSetDisplayProperties() js.Ref

//go:wasmimport plat/js/webext/system/display func_SetDisplayProperties
//go:noescape
func FuncSetDisplayProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_SetDisplayProperties
//go:noescape
func CallSetDisplayProperties(
	retPtr unsafe.Pointer,
	id js.Ref,
	info unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display try_SetDisplayProperties
//go:noescape
func TrySetDisplayProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	info unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_SetMirrorMode
//go:noescape
func HasFuncSetMirrorMode() js.Ref

//go:wasmimport plat/js/webext/system/display func_SetMirrorMode
//go:noescape
func FuncSetMirrorMode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_SetMirrorMode
//go:noescape
func CallSetMirrorMode(
	retPtr unsafe.Pointer,
	info unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display try_SetMirrorMode
//go:noescape
func TrySetMirrorMode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	info unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_ShowNativeTouchCalibration
//go:noescape
func HasFuncShowNativeTouchCalibration() js.Ref

//go:wasmimport plat/js/webext/system/display func_ShowNativeTouchCalibration
//go:noescape
func FuncShowNativeTouchCalibration(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_ShowNativeTouchCalibration
//go:noescape
func CallShowNativeTouchCalibration(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/system/display try_ShowNativeTouchCalibration
//go:noescape
func TryShowNativeTouchCalibration(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/system/display has_StartCustomTouchCalibration
//go:noescape
func HasFuncStartCustomTouchCalibration() js.Ref

//go:wasmimport plat/js/webext/system/display func_StartCustomTouchCalibration
//go:noescape
func FuncStartCustomTouchCalibration(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/system/display call_StartCustomTouchCalibration
//go:noescape
func CallStartCustomTouchCalibration(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/system/display try_StartCustomTouchCalibration
//go:noescape
func TryStartCustomTouchCalibration(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

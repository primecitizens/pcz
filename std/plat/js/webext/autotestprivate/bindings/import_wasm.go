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

//go:wasmimport plat/js/webext/autotestprivate store_Accelerator
//go:noescape
func AcceleratorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_Accelerator
//go:noescape
func AcceleratorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate constof_AppType
//go:noescape
func ConstOfAppType(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate constof_AppInstallSource
//go:noescape
func ConstOfAppInstallSource(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate constof_AppReadiness
//go:noescape
func ConstOfAppReadiness(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate store_App
//go:noescape
func AppJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_App
//go:noescape
func AppJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate constof_AppWindowType
//go:noescape
func ConstOfAppWindowType(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate constof_WindowStateType
//go:noescape
func ConstOfWindowStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate store_Bounds
//go:noescape
func BoundsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_Bounds
//go:noescape
func BoundsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate constof_FrameMode
//go:noescape
func ConstOfFrameMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate store_OverviewInfo
//go:noescape
func OverviewInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_OverviewInfo
//go:noescape
func OverviewInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_AppWindowInfo
//go:noescape
func AppWindowInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_AppWindowInfo
//go:noescape
func AppWindowInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ArcAppDict
//go:noescape
func ArcAppDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ArcAppDict
//go:noescape
func ArcAppDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ArcAppKillsDict
//go:noescape
func ArcAppKillsDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ArcAppKillsDict
//go:noescape
func ArcAppKillsDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ArcAppTracingInfo
//go:noescape
func ArcAppTracingInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ArcAppTracingInfo
//go:noescape
func ArcAppTracingInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ArcPackageDict
//go:noescape
func ArcPackageDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ArcPackageDict
//go:noescape
func ArcPackageDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ArcState
//go:noescape
func ArcStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ArcState
//go:noescape
func ArcStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_AssistantQueryResponse
//go:noescape
func AssistantQueryResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_AssistantQueryResponse
//go:noescape
func AssistantQueryResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_AssistantQueryStatus
//go:noescape
func AssistantQueryStatusJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_AssistantQueryStatus
//go:noescape
func AssistantQueryStatusJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_CryptohomeRecoveryDataDict
//go:noescape
func CryptohomeRecoveryDataDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_CryptohomeRecoveryDataDict
//go:noescape
func CryptohomeRecoveryDataDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_DesksInfo
//go:noescape
func DesksInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_DesksInfo
//go:noescape
func DesksInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_DisplaySmoothnessData
//go:noescape
func DisplaySmoothnessDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_DisplaySmoothnessData
//go:noescape
func DisplaySmoothnessDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ExtensionInfoDict
//go:noescape
func ExtensionInfoDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ExtensionInfoDict
//go:noescape
func ExtensionInfoDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ExtensionsInfoArray
//go:noescape
func ExtensionsInfoArrayJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ExtensionsInfoArray
//go:noescape
func ExtensionsInfoArrayJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_FrameCountingPerSinkData
//go:noescape
func FrameCountingPerSinkDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_FrameCountingPerSinkData
//go:noescape
func FrameCountingPerSinkDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_GetAccessTokenData
//go:noescape
func GetAccessTokenDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_GetAccessTokenData
//go:noescape
func GetAccessTokenDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_GetAccessTokenParams
//go:noescape
func GetAccessTokenParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_GetAccessTokenParams
//go:noescape
func GetAccessTokenParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_GetCurrentInputMethodDescriptorData
//go:noescape
func GetCurrentInputMethodDescriptorDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_GetCurrentInputMethodDescriptorData
//go:noescape
func GetCurrentInputMethodDescriptorDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate constof_LacrosState
//go:noescape
func ConstOfLacrosState(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate constof_LacrosMode
//go:noescape
func ConstOfLacrosMode(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate store_LacrosInfo
//go:noescape
func LacrosInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_LacrosInfo
//go:noescape
func LacrosInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_LauncherSearchBoxState
//go:noescape
func LauncherSearchBoxStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_LauncherSearchBoxState
//go:noescape
func LauncherSearchBoxStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_LoginEventRecorderData
//go:noescape
func LoginEventRecorderDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_LoginEventRecorderData
//go:noescape
func LoginEventRecorderDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_SystemWebApp
//go:noescape
func SystemWebAppJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_SystemWebApp
//go:noescape
func SystemWebAppJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ScrollableShelfInfo
//go:noescape
func ScrollableShelfInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ScrollableShelfInfo
//go:noescape
func ScrollableShelfInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate constof_ShelfAlignmentType
//go:noescape
func ConstOfShelfAlignmentType(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate constof_ShelfItemType
//go:noescape
func ConstOfShelfItemType(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate constof_ShelfItemStatus
//go:noescape
func ConstOfShelfItemStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate store_ShelfItem
//go:noescape
func ShelfItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ShelfItem
//go:noescape
func ShelfItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_Location
//go:noescape
func LocationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_Location
//go:noescape
func LocationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_HotseatSwipeDescriptor
//go:noescape
func HotseatSwipeDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_HotseatSwipeDescriptor
//go:noescape
func HotseatSwipeDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate constof_HotseatState
//go:noescape
func ConstOfHotseatState(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate store_HotseatInfo
//go:noescape
func HotseatInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_HotseatInfo
//go:noescape
func HotseatInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ShelfUIInfo
//go:noescape
func ShelfUIInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ShelfUIInfo
//go:noescape
func ShelfUIInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ThroughputTrackerAnimationData
//go:noescape
func ThroughputTrackerAnimationDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ThroughputTrackerAnimationData
//go:noescape
func ThroughputTrackerAnimationDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate constof_LauncherStateType
//go:noescape
func ConstOfLauncherStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate store_LoginStatusDict
//go:noescape
func LoginStatusDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_LoginStatusDict
//go:noescape
func LoginStatusDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_MakeFuseboxTempDirData
//go:noescape
func MakeFuseboxTempDirDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_MakeFuseboxTempDirData
//go:noescape
func MakeFuseboxTempDirDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate constof_MouseButton
//go:noescape
func ConstOfMouseButton(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate store_Notification
//go:noescape
func NotificationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_Notification
//go:noescape
func NotificationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate constof_OverviewStateType
//go:noescape
func ConstOfOverviewStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate store_PlayStoreState
//go:noescape
func PlayStoreStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_PlayStoreState
//go:noescape
func PlayStoreStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_Printer
//go:noescape
func PrinterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_Printer
//go:noescape
func PrinterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ResetHoldingSpaceOptions
//go:noescape
func ResetHoldingSpaceOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ResetHoldingSpaceOptions
//go:noescape
func ResetHoldingSpaceOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate constof_RotationType
//go:noescape
func ConstOfRotationType(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate store_ScrollableShelfState
//go:noescape
func ScrollableShelfStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ScrollableShelfState
//go:noescape
func ScrollableShelfStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_SetWindowBoundsResult
//go:noescape
func SetWindowBoundsResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_SetWindowBoundsResult
//go:noescape
func SetWindowBoundsResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ShelfIconPinUpdateParam
//go:noescape
func ShelfIconPinUpdateParamJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ShelfIconPinUpdateParam
//go:noescape
func ShelfIconPinUpdateParamJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate store_ShelfState
//go:noescape
func ShelfStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_ShelfState
//go:noescape
func ShelfStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate constof_ThemeStyle
//go:noescape
func ConstOfThemeStyle(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate constof_WMEventType
//go:noescape
func ConstOfWMEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/autotestprivate store_WindowStateChangeDict
//go:noescape
func WindowStateChangeDictJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/autotestprivate load_WindowStateChangeDict
//go:noescape
func WindowStateChangeDictJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/autotestprivate has_ActivateAccelerator
//go:noescape
func HasFuncActivateAccelerator() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ActivateAccelerator
//go:noescape
func FuncActivateAccelerator(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ActivateAccelerator
//go:noescape
func CallActivateAccelerator(
	retPtr unsafe.Pointer,
	accelerator unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_ActivateAccelerator
//go:noescape
func TryActivateAccelerator(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	accelerator unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_ActivateAdjacentDesksToTargetIndex
//go:noescape
func HasFuncActivateAdjacentDesksToTargetIndex() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ActivateAdjacentDesksToTargetIndex
//go:noescape
func FuncActivateAdjacentDesksToTargetIndex(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ActivateAdjacentDesksToTargetIndex
//go:noescape
func CallActivateAdjacentDesksToTargetIndex(
	retPtr unsafe.Pointer,
	index int32)

//go:wasmimport plat/js/webext/autotestprivate try_ActivateAdjacentDesksToTargetIndex
//go:noescape
func TryActivateAdjacentDesksToTargetIndex(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_ActivateAppWindow
//go:noescape
func HasFuncActivateAppWindow() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ActivateAppWindow
//go:noescape
func FuncActivateAppWindow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ActivateAppWindow
//go:noescape
func CallActivateAppWindow(
	retPtr unsafe.Pointer,
	id int32)

//go:wasmimport plat/js/webext/autotestprivate try_ActivateAppWindow
//go:noescape
func TryActivateAppWindow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_ActivateDeskAtIndex
//go:noescape
func HasFuncActivateDeskAtIndex() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ActivateDeskAtIndex
//go:noescape
func FuncActivateDeskAtIndex(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ActivateDeskAtIndex
//go:noescape
func CallActivateDeskAtIndex(
	retPtr unsafe.Pointer,
	index int32)

//go:wasmimport plat/js/webext/autotestprivate try_ActivateDeskAtIndex
//go:noescape
func TryActivateDeskAtIndex(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_AddLoginEventForTesting
//go:noescape
func HasFuncAddLoginEventForTesting() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_AddLoginEventForTesting
//go:noescape
func FuncAddLoginEventForTesting(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_AddLoginEventForTesting
//go:noescape
func CallAddLoginEventForTesting(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_AddLoginEventForTesting
//go:noescape
func TryAddLoginEventForTesting(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_ArcAppTracingStart
//go:noescape
func HasFuncArcAppTracingStart() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ArcAppTracingStart
//go:noescape
func FuncArcAppTracingStart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ArcAppTracingStart
//go:noescape
func CallArcAppTracingStart(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_ArcAppTracingStart
//go:noescape
func TryArcAppTracingStart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_ArcAppTracingStopAndAnalyze
//go:noescape
func HasFuncArcAppTracingStopAndAnalyze() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ArcAppTracingStopAndAnalyze
//go:noescape
func FuncArcAppTracingStopAndAnalyze(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ArcAppTracingStopAndAnalyze
//go:noescape
func CallArcAppTracingStopAndAnalyze(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_ArcAppTracingStopAndAnalyze
//go:noescape
func TryArcAppTracingStopAndAnalyze(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_BootstrapMachineLearningService
//go:noescape
func HasFuncBootstrapMachineLearningService() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_BootstrapMachineLearningService
//go:noescape
func FuncBootstrapMachineLearningService(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_BootstrapMachineLearningService
//go:noescape
func CallBootstrapMachineLearningService(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_BootstrapMachineLearningService
//go:noescape
func TryBootstrapMachineLearningService(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_CloseApp
//go:noescape
func HasFuncCloseApp() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_CloseApp
//go:noescape
func FuncCloseApp(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_CloseApp
//go:noescape
func CallCloseApp(
	retPtr unsafe.Pointer,
	appId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_CloseApp
//go:noescape
func TryCloseApp(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	appId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_CloseAppWindow
//go:noescape
func HasFuncCloseAppWindow() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_CloseAppWindow
//go:noescape
func FuncCloseAppWindow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_CloseAppWindow
//go:noescape
func CallCloseAppWindow(
	retPtr unsafe.Pointer,
	id int32)

//go:wasmimport plat/js/webext/autotestprivate try_CloseAppWindow
//go:noescape
func TryCloseAppWindow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_CouldAllowCrostini
//go:noescape
func HasFuncCouldAllowCrostini() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_CouldAllowCrostini
//go:noescape
func FuncCouldAllowCrostini(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_CouldAllowCrostini
//go:noescape
func CallCouldAllowCrostini(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_CouldAllowCrostini
//go:noescape
func TryCouldAllowCrostini(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_CreateNewDesk
//go:noescape
func HasFuncCreateNewDesk() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_CreateNewDesk
//go:noescape
func FuncCreateNewDesk(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_CreateNewDesk
//go:noescape
func CallCreateNewDesk(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_CreateNewDesk
//go:noescape
func TryCreateNewDesk(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_DisableAutomation
//go:noescape
func HasFuncDisableAutomation() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_DisableAutomation
//go:noescape
func FuncDisableAutomation(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_DisableAutomation
//go:noescape
func CallDisableAutomation(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_DisableAutomation
//go:noescape
func TryDisableAutomation(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_DisableSwitchAccessDialog
//go:noescape
func HasFuncDisableSwitchAccessDialog() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_DisableSwitchAccessDialog
//go:noescape
func FuncDisableSwitchAccessDialog(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_DisableSwitchAccessDialog
//go:noescape
func CallDisableSwitchAccessDialog(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_DisableSwitchAccessDialog
//go:noescape
func TryDisableSwitchAccessDialog(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_EnableAssistantAndWaitForReady
//go:noescape
func HasFuncEnableAssistantAndWaitForReady() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_EnableAssistantAndWaitForReady
//go:noescape
func FuncEnableAssistantAndWaitForReady(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_EnableAssistantAndWaitForReady
//go:noescape
func CallEnableAssistantAndWaitForReady(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_EnableAssistantAndWaitForReady
//go:noescape
func TryEnableAssistantAndWaitForReady(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_ExportCrostini
//go:noescape
func HasFuncExportCrostini() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ExportCrostini
//go:noescape
func FuncExportCrostini(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ExportCrostini
//go:noescape
func CallExportCrostini(
	retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_ExportCrostini
//go:noescape
func TryExportCrostini(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_ForceAutoThemeMode
//go:noescape
func HasFuncForceAutoThemeMode() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ForceAutoThemeMode
//go:noescape
func FuncForceAutoThemeMode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ForceAutoThemeMode
//go:noescape
func CallForceAutoThemeMode(
	retPtr unsafe.Pointer,
	darkModeEnabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_ForceAutoThemeMode
//go:noescape
func TryForceAutoThemeMode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	darkModeEnabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetAccessToken
//go:noescape
func HasFuncGetAccessToken() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetAccessToken
//go:noescape
func FuncGetAccessToken(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetAccessToken
//go:noescape
func CallGetAccessToken(
	retPtr unsafe.Pointer,
	accessTokenParams unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetAccessToken
//go:noescape
func TryGetAccessToken(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	accessTokenParams unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetAllEnterprisePolicies
//go:noescape
func HasFuncGetAllEnterprisePolicies() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetAllEnterprisePolicies
//go:noescape
func FuncGetAllEnterprisePolicies(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetAllEnterprisePolicies
//go:noescape
func CallGetAllEnterprisePolicies(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetAllEnterprisePolicies
//go:noescape
func TryGetAllEnterprisePolicies(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetAllInstalledApps
//go:noescape
func HasFuncGetAllInstalledApps() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetAllInstalledApps
//go:noescape
func FuncGetAllInstalledApps(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetAllInstalledApps
//go:noescape
func CallGetAllInstalledApps(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetAllInstalledApps
//go:noescape
func TryGetAllInstalledApps(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetAppWindowList
//go:noescape
func HasFuncGetAppWindowList() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetAppWindowList
//go:noescape
func FuncGetAppWindowList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetAppWindowList
//go:noescape
func CallGetAppWindowList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetAppWindowList
//go:noescape
func TryGetAppWindowList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetArcApp
//go:noescape
func HasFuncGetArcApp() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetArcApp
//go:noescape
func FuncGetArcApp(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetArcApp
//go:noescape
func CallGetArcApp(
	retPtr unsafe.Pointer,
	appId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_GetArcApp
//go:noescape
func TryGetArcApp(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	appId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetArcAppKills
//go:noescape
func HasFuncGetArcAppKills() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetArcAppKills
//go:noescape
func FuncGetArcAppKills(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetArcAppKills
//go:noescape
func CallGetArcAppKills(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetArcAppKills
//go:noescape
func TryGetArcAppKills(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetArcPackage
//go:noescape
func HasFuncGetArcPackage() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetArcPackage
//go:noescape
func FuncGetArcPackage(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetArcPackage
//go:noescape
func CallGetArcPackage(
	retPtr unsafe.Pointer,
	packageName js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_GetArcPackage
//go:noescape
func TryGetArcPackage(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	packageName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetArcStartTime
//go:noescape
func HasFuncGetArcStartTime() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetArcStartTime
//go:noescape
func FuncGetArcStartTime(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetArcStartTime
//go:noescape
func CallGetArcStartTime(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetArcStartTime
//go:noescape
func TryGetArcStartTime(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetArcState
//go:noescape
func HasFuncGetArcState() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetArcState
//go:noescape
func FuncGetArcState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetArcState
//go:noescape
func CallGetArcState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetArcState
//go:noescape
func TryGetArcState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetClipboardTextData
//go:noescape
func HasFuncGetClipboardTextData() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetClipboardTextData
//go:noescape
func FuncGetClipboardTextData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetClipboardTextData
//go:noescape
func CallGetClipboardTextData(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetClipboardTextData
//go:noescape
func TryGetClipboardTextData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetCryptohomeRecoveryData
//go:noescape
func HasFuncGetCryptohomeRecoveryData() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetCryptohomeRecoveryData
//go:noescape
func FuncGetCryptohomeRecoveryData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetCryptohomeRecoveryData
//go:noescape
func CallGetCryptohomeRecoveryData(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetCryptohomeRecoveryData
//go:noescape
func TryGetCryptohomeRecoveryData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetCurrentInputMethodDescriptor
//go:noescape
func HasFuncGetCurrentInputMethodDescriptor() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetCurrentInputMethodDescriptor
//go:noescape
func FuncGetCurrentInputMethodDescriptor(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetCurrentInputMethodDescriptor
//go:noescape
func CallGetCurrentInputMethodDescriptor(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetCurrentInputMethodDescriptor
//go:noescape
func TryGetCurrentInputMethodDescriptor(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetDefaultPinnedAppIds
//go:noescape
func HasFuncGetDefaultPinnedAppIds() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetDefaultPinnedAppIds
//go:noescape
func FuncGetDefaultPinnedAppIds(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetDefaultPinnedAppIds
//go:noescape
func CallGetDefaultPinnedAppIds(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetDefaultPinnedAppIds
//go:noescape
func TryGetDefaultPinnedAppIds(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetDeskCount
//go:noescape
func HasFuncGetDeskCount() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetDeskCount
//go:noescape
func FuncGetDeskCount(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetDeskCount
//go:noescape
func CallGetDeskCount(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetDeskCount
//go:noescape
func TryGetDeskCount(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetDesksInfo
//go:noescape
func HasFuncGetDesksInfo() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetDesksInfo
//go:noescape
func FuncGetDesksInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetDesksInfo
//go:noescape
func CallGetDesksInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetDesksInfo
//go:noescape
func TryGetDesksInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetDisplaySmoothness
//go:noescape
func HasFuncGetDisplaySmoothness() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetDisplaySmoothness
//go:noescape
func FuncGetDisplaySmoothness(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetDisplaySmoothness
//go:noescape
func CallGetDisplaySmoothness(
	retPtr unsafe.Pointer,
	displayId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_GetDisplaySmoothness
//go:noescape
func TryGetDisplaySmoothness(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	displayId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetExtensionsInfo
//go:noescape
func HasFuncGetExtensionsInfo() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetExtensionsInfo
//go:noescape
func FuncGetExtensionsInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetExtensionsInfo
//go:noescape
func CallGetExtensionsInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetExtensionsInfo
//go:noescape
func TryGetExtensionsInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetLacrosInfo
//go:noescape
func HasFuncGetLacrosInfo() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetLacrosInfo
//go:noescape
func FuncGetLacrosInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetLacrosInfo
//go:noescape
func CallGetLacrosInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetLacrosInfo
//go:noescape
func TryGetLacrosInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetLauncherSearchBoxState
//go:noescape
func HasFuncGetLauncherSearchBoxState() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetLauncherSearchBoxState
//go:noescape
func FuncGetLauncherSearchBoxState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetLauncherSearchBoxState
//go:noescape
func CallGetLauncherSearchBoxState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetLauncherSearchBoxState
//go:noescape
func TryGetLauncherSearchBoxState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetLoginEventRecorderLoginEvents
//go:noescape
func HasFuncGetLoginEventRecorderLoginEvents() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetLoginEventRecorderLoginEvents
//go:noescape
func FuncGetLoginEventRecorderLoginEvents(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetLoginEventRecorderLoginEvents
//go:noescape
func CallGetLoginEventRecorderLoginEvents(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetLoginEventRecorderLoginEvents
//go:noescape
func TryGetLoginEventRecorderLoginEvents(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetPlayStoreState
//go:noescape
func HasFuncGetPlayStoreState() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetPlayStoreState
//go:noescape
func FuncGetPlayStoreState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetPlayStoreState
//go:noescape
func CallGetPlayStoreState(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetPlayStoreState
//go:noescape
func TryGetPlayStoreState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetPrimaryDisplayScaleFactor
//go:noescape
func HasFuncGetPrimaryDisplayScaleFactor() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetPrimaryDisplayScaleFactor
//go:noescape
func FuncGetPrimaryDisplayScaleFactor(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetPrimaryDisplayScaleFactor
//go:noescape
func CallGetPrimaryDisplayScaleFactor(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetPrimaryDisplayScaleFactor
//go:noescape
func TryGetPrimaryDisplayScaleFactor(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetPrinterList
//go:noescape
func HasFuncGetPrinterList() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetPrinterList
//go:noescape
func FuncGetPrinterList(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetPrinterList
//go:noescape
func CallGetPrinterList(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetPrinterList
//go:noescape
func TryGetPrinterList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetRegisteredSystemWebApps
//go:noescape
func HasFuncGetRegisteredSystemWebApps() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetRegisteredSystemWebApps
//go:noescape
func FuncGetRegisteredSystemWebApps(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetRegisteredSystemWebApps
//go:noescape
func CallGetRegisteredSystemWebApps(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetRegisteredSystemWebApps
//go:noescape
func TryGetRegisteredSystemWebApps(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetScrollableShelfInfoForState
//go:noescape
func HasFuncGetScrollableShelfInfoForState() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetScrollableShelfInfoForState
//go:noescape
func FuncGetScrollableShelfInfoForState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetScrollableShelfInfoForState
//go:noescape
func CallGetScrollableShelfInfoForState(
	retPtr unsafe.Pointer,
	state unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetScrollableShelfInfoForState
//go:noescape
func TryGetScrollableShelfInfoForState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	state unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetShelfAlignment
//go:noescape
func HasFuncGetShelfAlignment() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetShelfAlignment
//go:noescape
func FuncGetShelfAlignment(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetShelfAlignment
//go:noescape
func CallGetShelfAlignment(
	retPtr unsafe.Pointer,
	displayId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_GetShelfAlignment
//go:noescape
func TryGetShelfAlignment(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	displayId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetShelfAutoHideBehavior
//go:noescape
func HasFuncGetShelfAutoHideBehavior() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetShelfAutoHideBehavior
//go:noescape
func FuncGetShelfAutoHideBehavior(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetShelfAutoHideBehavior
//go:noescape
func CallGetShelfAutoHideBehavior(
	retPtr unsafe.Pointer,
	displayId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_GetShelfAutoHideBehavior
//go:noescape
func TryGetShelfAutoHideBehavior(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	displayId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetShelfItems
//go:noescape
func HasFuncGetShelfItems() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetShelfItems
//go:noescape
func FuncGetShelfItems(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetShelfItems
//go:noescape
func CallGetShelfItems(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetShelfItems
//go:noescape
func TryGetShelfItems(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetShelfUIInfoForState
//go:noescape
func HasFuncGetShelfUIInfoForState() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetShelfUIInfoForState
//go:noescape
func FuncGetShelfUIInfoForState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetShelfUIInfoForState
//go:noescape
func CallGetShelfUIInfoForState(
	retPtr unsafe.Pointer,
	state unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetShelfUIInfoForState
//go:noescape
func TryGetShelfUIInfoForState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	state unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetThroughputTrackerData
//go:noescape
func HasFuncGetThroughputTrackerData() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetThroughputTrackerData
//go:noescape
func FuncGetThroughputTrackerData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetThroughputTrackerData
//go:noescape
func CallGetThroughputTrackerData(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetThroughputTrackerData
//go:noescape
func TryGetThroughputTrackerData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_GetVisibleNotifications
//go:noescape
func HasFuncGetVisibleNotifications() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_GetVisibleNotifications
//go:noescape
func FuncGetVisibleNotifications(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_GetVisibleNotifications
//go:noescape
func CallGetVisibleNotifications(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_GetVisibleNotifications
//go:noescape
func TryGetVisibleNotifications(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_ImportCrostini
//go:noescape
func HasFuncImportCrostini() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ImportCrostini
//go:noescape
func FuncImportCrostini(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ImportCrostini
//go:noescape
func CallImportCrostini(
	retPtr unsafe.Pointer,
	path js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_ImportCrostini
//go:noescape
func TryImportCrostini(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_InitializeEvents
//go:noescape
func HasFuncInitializeEvents() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_InitializeEvents
//go:noescape
func FuncInitializeEvents(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_InitializeEvents
//go:noescape
func CallInitializeEvents(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_InitializeEvents
//go:noescape
func TryInitializeEvents(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_InstallBorealis
//go:noescape
func HasFuncInstallBorealis() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_InstallBorealis
//go:noescape
func FuncInstallBorealis(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_InstallBorealis
//go:noescape
func CallInstallBorealis(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_InstallBorealis
//go:noescape
func TryInstallBorealis(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_InstallBruschetta
//go:noescape
func HasFuncInstallBruschetta() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_InstallBruschetta
//go:noescape
func FuncInstallBruschetta(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_InstallBruschetta
//go:noescape
func CallInstallBruschetta(
	retPtr unsafe.Pointer,
	vm_name js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_InstallBruschetta
//go:noescape
func TryInstallBruschetta(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vm_name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_InstallPWAForCurrentURL
//go:noescape
func HasFuncInstallPWAForCurrentURL() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_InstallPWAForCurrentURL
//go:noescape
func FuncInstallPWAForCurrentURL(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_InstallPWAForCurrentURL
//go:noescape
func CallInstallPWAForCurrentURL(
	retPtr unsafe.Pointer,
	timeout_ms int32)

//go:wasmimport plat/js/webext/autotestprivate try_InstallPWAForCurrentURL
//go:noescape
func TryInstallPWAForCurrentURL(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	timeout_ms int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_IsAppShown
//go:noescape
func HasFuncIsAppShown() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_IsAppShown
//go:noescape
func FuncIsAppShown(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_IsAppShown
//go:noescape
func CallIsAppShown(
	retPtr unsafe.Pointer,
	appId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_IsAppShown
//go:noescape
func TryIsAppShown(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	appId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_IsArcPackageListInitialRefreshed
//go:noescape
func HasFuncIsArcPackageListInitialRefreshed() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_IsArcPackageListInitialRefreshed
//go:noescape
func FuncIsArcPackageListInitialRefreshed(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_IsArcPackageListInitialRefreshed
//go:noescape
func CallIsArcPackageListInitialRefreshed(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_IsArcPackageListInitialRefreshed
//go:noescape
func TryIsArcPackageListInitialRefreshed(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_IsArcProvisioned
//go:noescape
func HasFuncIsArcProvisioned() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_IsArcProvisioned
//go:noescape
func FuncIsArcProvisioned(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_IsArcProvisioned
//go:noescape
func CallIsArcProvisioned(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_IsArcProvisioned
//go:noescape
func TryIsArcProvisioned(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_IsFeatureEnabled
//go:noescape
func HasFuncIsFeatureEnabled() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_IsFeatureEnabled
//go:noescape
func FuncIsFeatureEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_IsFeatureEnabled
//go:noescape
func CallIsFeatureEnabled(
	retPtr unsafe.Pointer,
	feature_name js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_IsFeatureEnabled
//go:noescape
func TryIsFeatureEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	feature_name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_IsInputMethodReadyForTesting
//go:noescape
func HasFuncIsInputMethodReadyForTesting() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_IsInputMethodReadyForTesting
//go:noescape
func FuncIsInputMethodReadyForTesting(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_IsInputMethodReadyForTesting
//go:noescape
func CallIsInputMethodReadyForTesting(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_IsInputMethodReadyForTesting
//go:noescape
func TryIsInputMethodReadyForTesting(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_IsSystemWebAppOpen
//go:noescape
func HasFuncIsSystemWebAppOpen() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_IsSystemWebAppOpen
//go:noescape
func FuncIsSystemWebAppOpen(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_IsSystemWebAppOpen
//go:noescape
func CallIsSystemWebAppOpen(
	retPtr unsafe.Pointer,
	appId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_IsSystemWebAppOpen
//go:noescape
func TryIsSystemWebAppOpen(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	appId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_IsTabletModeEnabled
//go:noescape
func HasFuncIsTabletModeEnabled() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_IsTabletModeEnabled
//go:noescape
func FuncIsTabletModeEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_IsTabletModeEnabled
//go:noescape
func CallIsTabletModeEnabled(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_IsTabletModeEnabled
//go:noescape
func TryIsTabletModeEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_LaunchApp
//go:noescape
func HasFuncLaunchApp() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_LaunchApp
//go:noescape
func FuncLaunchApp(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_LaunchApp
//go:noescape
func CallLaunchApp(
	retPtr unsafe.Pointer,
	appId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_LaunchApp
//go:noescape
func TryLaunchApp(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	appId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_LaunchFilesAppToPath
//go:noescape
func HasFuncLaunchFilesAppToPath() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_LaunchFilesAppToPath
//go:noescape
func FuncLaunchFilesAppToPath(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_LaunchFilesAppToPath
//go:noescape
func CallLaunchFilesAppToPath(
	retPtr unsafe.Pointer,
	absolutePath js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_LaunchFilesAppToPath
//go:noescape
func TryLaunchFilesAppToPath(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	absolutePath js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_LaunchSystemWebApp
//go:noescape
func HasFuncLaunchSystemWebApp() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_LaunchSystemWebApp
//go:noescape
func FuncLaunchSystemWebApp(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_LaunchSystemWebApp
//go:noescape
func CallLaunchSystemWebApp(
	retPtr unsafe.Pointer,
	appName js.Ref,
	url js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_LaunchSystemWebApp
//go:noescape
func TryLaunchSystemWebApp(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	appName js.Ref,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_LoadSmartDimComponent
//go:noescape
func HasFuncLoadSmartDimComponent() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_LoadSmartDimComponent
//go:noescape
func FuncLoadSmartDimComponent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_LoadSmartDimComponent
//go:noescape
func CallLoadSmartDimComponent(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_LoadSmartDimComponent
//go:noescape
func TryLoadSmartDimComponent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_LockScreen
//go:noescape
func HasFuncLockScreen() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_LockScreen
//go:noescape
func FuncLockScreen(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_LockScreen
//go:noescape
func CallLockScreen(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_LockScreen
//go:noescape
func TryLockScreen(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_LoginStatus
//go:noescape
func HasFuncLoginStatus() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_LoginStatus
//go:noescape
func FuncLoginStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_LoginStatus
//go:noescape
func CallLoginStatus(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_LoginStatus
//go:noescape
func TryLoginStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_Logout
//go:noescape
func HasFuncLogout() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_Logout
//go:noescape
func FuncLogout(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_Logout
//go:noescape
func CallLogout(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_Logout
//go:noescape
func TryLogout(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_MakeFuseboxTempDir
//go:noescape
func HasFuncMakeFuseboxTempDir() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_MakeFuseboxTempDir
//go:noescape
func FuncMakeFuseboxTempDir(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_MakeFuseboxTempDir
//go:noescape
func CallMakeFuseboxTempDir(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_MakeFuseboxTempDir
//go:noescape
func TryMakeFuseboxTempDir(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_MouseClick
//go:noescape
func HasFuncMouseClick() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_MouseClick
//go:noescape
func FuncMouseClick(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_MouseClick
//go:noescape
func CallMouseClick(
	retPtr unsafe.Pointer,
	button uint32)

//go:wasmimport plat/js/webext/autotestprivate try_MouseClick
//go:noescape
func TryMouseClick(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	button uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_MouseMove
//go:noescape
func HasFuncMouseMove() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_MouseMove
//go:noescape
func FuncMouseMove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_MouseMove
//go:noescape
func CallMouseMove(
	retPtr unsafe.Pointer,
	location unsafe.Pointer,
	duration_in_ms float64)

//go:wasmimport plat/js/webext/autotestprivate try_MouseMove
//go:noescape
func TryMouseMove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	location unsafe.Pointer,
	duration_in_ms float64) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_MousePress
//go:noescape
func HasFuncMousePress() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_MousePress
//go:noescape
func FuncMousePress(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_MousePress
//go:noescape
func CallMousePress(
	retPtr unsafe.Pointer,
	button uint32)

//go:wasmimport plat/js/webext/autotestprivate try_MousePress
//go:noescape
func TryMousePress(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	button uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_MouseRelease
//go:noescape
func HasFuncMouseRelease() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_MouseRelease
//go:noescape
func FuncMouseRelease(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_MouseRelease
//go:noescape
func CallMouseRelease(
	retPtr unsafe.Pointer,
	button uint32)

//go:wasmimport plat/js/webext/autotestprivate try_MouseRelease
//go:noescape
func TryMouseRelease(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	button uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_OnClipboardDataChanged
//go:noescape
func HasFuncOnClipboardDataChanged() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_OnClipboardDataChanged
//go:noescape
func FuncOnClipboardDataChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_OnClipboardDataChanged
//go:noescape
func CallOnClipboardDataChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_OnClipboardDataChanged
//go:noescape
func TryOnClipboardDataChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_OffClipboardDataChanged
//go:noescape
func HasFuncOffClipboardDataChanged() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_OffClipboardDataChanged
//go:noescape
func FuncOffClipboardDataChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_OffClipboardDataChanged
//go:noescape
func CallOffClipboardDataChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_OffClipboardDataChanged
//go:noescape
func TryOffClipboardDataChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_HasOnClipboardDataChanged
//go:noescape
func HasFuncHasOnClipboardDataChanged() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_HasOnClipboardDataChanged
//go:noescape
func FuncHasOnClipboardDataChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_HasOnClipboardDataChanged
//go:noescape
func CallHasOnClipboardDataChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_HasOnClipboardDataChanged
//go:noescape
func TryHasOnClipboardDataChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_PinShelfIcon
//go:noescape
func HasFuncPinShelfIcon() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_PinShelfIcon
//go:noescape
func FuncPinShelfIcon(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_PinShelfIcon
//go:noescape
func CallPinShelfIcon(
	retPtr unsafe.Pointer,
	appId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_PinShelfIcon
//go:noescape
func TryPinShelfIcon(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	appId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_RefreshEnterprisePolicies
//go:noescape
func HasFuncRefreshEnterprisePolicies() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_RefreshEnterprisePolicies
//go:noescape
func FuncRefreshEnterprisePolicies(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_RefreshEnterprisePolicies
//go:noescape
func CallRefreshEnterprisePolicies(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_RefreshEnterprisePolicies
//go:noescape
func TryRefreshEnterprisePolicies(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_RefreshRemoteCommands
//go:noescape
func HasFuncRefreshRemoteCommands() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_RefreshRemoteCommands
//go:noescape
func FuncRefreshRemoteCommands(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_RefreshRemoteCommands
//go:noescape
func CallRefreshRemoteCommands(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_RefreshRemoteCommands
//go:noescape
func TryRefreshRemoteCommands(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_RegisterComponent
//go:noescape
func HasFuncRegisterComponent() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_RegisterComponent
//go:noescape
func FuncRegisterComponent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_RegisterComponent
//go:noescape
func CallRegisterComponent(
	retPtr unsafe.Pointer,
	name js.Ref,
	path js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_RegisterComponent
//go:noescape
func TryRegisterComponent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	path js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_RemoveActiveDesk
//go:noescape
func HasFuncRemoveActiveDesk() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_RemoveActiveDesk
//go:noescape
func FuncRemoveActiveDesk(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_RemoveActiveDesk
//go:noescape
func CallRemoveActiveDesk(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_RemoveActiveDesk
//go:noescape
func TryRemoveActiveDesk(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_RemoveAllNotifications
//go:noescape
func HasFuncRemoveAllNotifications() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_RemoveAllNotifications
//go:noescape
func FuncRemoveAllNotifications(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_RemoveAllNotifications
//go:noescape
func CallRemoveAllNotifications(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_RemoveAllNotifications
//go:noescape
func TryRemoveAllNotifications(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_RemoveBruschetta
//go:noescape
func HasFuncRemoveBruschetta() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_RemoveBruschetta
//go:noescape
func FuncRemoveBruschetta(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_RemoveBruschetta
//go:noescape
func CallRemoveBruschetta(
	retPtr unsafe.Pointer,
	vm_name js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_RemoveBruschetta
//go:noescape
func TryRemoveBruschetta(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	vm_name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_RemoveComponentExtension
//go:noescape
func HasFuncRemoveComponentExtension() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_RemoveComponentExtension
//go:noescape
func FuncRemoveComponentExtension(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_RemoveComponentExtension
//go:noescape
func CallRemoveComponentExtension(
	retPtr unsafe.Pointer,
	extensionId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_RemoveComponentExtension
//go:noescape
func TryRemoveComponentExtension(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	extensionId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_RemoveFuseboxTempDir
//go:noescape
func HasFuncRemoveFuseboxTempDir() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_RemoveFuseboxTempDir
//go:noescape
func FuncRemoveFuseboxTempDir(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_RemoveFuseboxTempDir
//go:noescape
func CallRemoveFuseboxTempDir(
	retPtr unsafe.Pointer,
	fuseboxFilePath js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_RemoveFuseboxTempDir
//go:noescape
func TryRemoveFuseboxTempDir(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	fuseboxFilePath js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_RemovePrinter
//go:noescape
func HasFuncRemovePrinter() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_RemovePrinter
//go:noescape
func FuncRemovePrinter(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_RemovePrinter
//go:noescape
func CallRemovePrinter(
	retPtr unsafe.Pointer,
	printerId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_RemovePrinter
//go:noescape
func TryRemovePrinter(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	printerId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_ResetHoldingSpace
//go:noescape
func HasFuncResetHoldingSpace() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ResetHoldingSpace
//go:noescape
func FuncResetHoldingSpace(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ResetHoldingSpace
//go:noescape
func CallResetHoldingSpace(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_ResetHoldingSpace
//go:noescape
func TryResetHoldingSpace(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_Restart
//go:noescape
func HasFuncRestart() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_Restart
//go:noescape
func FuncRestart(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_Restart
//go:noescape
func CallRestart(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_Restart
//go:noescape
func TryRestart(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_RunCrostiniInstaller
//go:noescape
func HasFuncRunCrostiniInstaller() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_RunCrostiniInstaller
//go:noescape
func FuncRunCrostiniInstaller(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_RunCrostiniInstaller
//go:noescape
func CallRunCrostiniInstaller(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_RunCrostiniInstaller
//go:noescape
func TryRunCrostiniInstaller(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_RunCrostiniUninstaller
//go:noescape
func HasFuncRunCrostiniUninstaller() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_RunCrostiniUninstaller
//go:noescape
func FuncRunCrostiniUninstaller(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_RunCrostiniUninstaller
//go:noescape
func CallRunCrostiniUninstaller(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_RunCrostiniUninstaller
//go:noescape
func TryRunCrostiniUninstaller(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SendArcOverlayColor
//go:noescape
func HasFuncSendArcOverlayColor() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SendArcOverlayColor
//go:noescape
func FuncSendArcOverlayColor(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SendArcOverlayColor
//go:noescape
func CallSendArcOverlayColor(
	retPtr unsafe.Pointer,
	color int32,
	theme uint32)

//go:wasmimport plat/js/webext/autotestprivate try_SendArcOverlayColor
//go:noescape
func TrySendArcOverlayColor(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	color int32,
	theme uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SendAssistantTextQuery
//go:noescape
func HasFuncSendAssistantTextQuery() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SendAssistantTextQuery
//go:noescape
func FuncSendAssistantTextQuery(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SendAssistantTextQuery
//go:noescape
func CallSendAssistantTextQuery(
	retPtr unsafe.Pointer,
	query js.Ref,
	timeout_ms int32)

//go:wasmimport plat/js/webext/autotestprivate try_SendAssistantTextQuery
//go:noescape
func TrySendAssistantTextQuery(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	timeout_ms int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetAllowedPref
//go:noescape
func HasFuncSetAllowedPref() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetAllowedPref
//go:noescape
func FuncSetAllowedPref(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetAllowedPref
//go:noescape
func CallSetAllowedPref(
	retPtr unsafe.Pointer,
	pref_name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetAllowedPref
//go:noescape
func TrySetAllowedPref(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	pref_name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetAppWindowState
//go:noescape
func HasFuncSetAppWindowState() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetAppWindowState
//go:noescape
func FuncSetAppWindowState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetAppWindowState
//go:noescape
func CallSetAppWindowState(
	retPtr unsafe.Pointer,
	id int32,
	change unsafe.Pointer,
	wait js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetAppWindowState
//go:noescape
func TrySetAppWindowState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32,
	change unsafe.Pointer,
	wait js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetArcAppWindowFocus
//go:noescape
func HasFuncSetArcAppWindowFocus() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetArcAppWindowFocus
//go:noescape
func FuncSetArcAppWindowFocus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetArcAppWindowFocus
//go:noescape
func CallSetArcAppWindowFocus(
	retPtr unsafe.Pointer,
	packageName js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetArcAppWindowFocus
//go:noescape
func TrySetArcAppWindowFocus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	packageName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetArcInteractiveState
//go:noescape
func HasFuncSetArcInteractiveState() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetArcInteractiveState
//go:noescape
func FuncSetArcInteractiveState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetArcInteractiveState
//go:noescape
func CallSetArcInteractiveState(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetArcInteractiveState
//go:noescape
func TrySetArcInteractiveState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetArcTouchMode
//go:noescape
func HasFuncSetArcTouchMode() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetArcTouchMode
//go:noescape
func FuncSetArcTouchMode(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetArcTouchMode
//go:noescape
func CallSetArcTouchMode(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetArcTouchMode
//go:noescape
func TrySetArcTouchMode(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetAssistantEnabled
//go:noescape
func HasFuncSetAssistantEnabled() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetAssistantEnabled
//go:noescape
func FuncSetAssistantEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetAssistantEnabled
//go:noescape
func CallSetAssistantEnabled(
	retPtr unsafe.Pointer,
	enabled js.Ref,
	timeout_ms int32)

//go:wasmimport plat/js/webext/autotestprivate try_SetAssistantEnabled
//go:noescape
func TrySetAssistantEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref,
	timeout_ms int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetClipboardTextData
//go:noescape
func HasFuncSetClipboardTextData() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetClipboardTextData
//go:noescape
func FuncSetClipboardTextData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetClipboardTextData
//go:noescape
func CallSetClipboardTextData(
	retPtr unsafe.Pointer,
	data js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetClipboardTextData
//go:noescape
func TrySetClipboardTextData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetCrostiniAppScaled
//go:noescape
func HasFuncSetCrostiniAppScaled() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetCrostiniAppScaled
//go:noescape
func FuncSetCrostiniAppScaled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetCrostiniAppScaled
//go:noescape
func CallSetCrostiniAppScaled(
	retPtr unsafe.Pointer,
	appId js.Ref,
	scaled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetCrostiniAppScaled
//go:noescape
func TrySetCrostiniAppScaled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	appId js.Ref,
	scaled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetCrostiniEnabled
//go:noescape
func HasFuncSetCrostiniEnabled() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetCrostiniEnabled
//go:noescape
func FuncSetCrostiniEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetCrostiniEnabled
//go:noescape
func CallSetCrostiniEnabled(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetCrostiniEnabled
//go:noescape
func TrySetCrostiniEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetMetricsEnabled
//go:noescape
func HasFuncSetMetricsEnabled() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetMetricsEnabled
//go:noescape
func FuncSetMetricsEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetMetricsEnabled
//go:noescape
func CallSetMetricsEnabled(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetMetricsEnabled
//go:noescape
func TrySetMetricsEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetMouseReverseScroll
//go:noescape
func HasFuncSetMouseReverseScroll() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetMouseReverseScroll
//go:noescape
func FuncSetMouseReverseScroll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetMouseReverseScroll
//go:noescape
func CallSetMouseReverseScroll(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetMouseReverseScroll
//go:noescape
func TrySetMouseReverseScroll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetMouseSensitivity
//go:noescape
func HasFuncSetMouseSensitivity() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetMouseSensitivity
//go:noescape
func FuncSetMouseSensitivity(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetMouseSensitivity
//go:noescape
func CallSetMouseSensitivity(
	retPtr unsafe.Pointer,
	value int32)

//go:wasmimport plat/js/webext/autotestprivate try_SetMouseSensitivity
//go:noescape
func TrySetMouseSensitivity(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetNaturalScroll
//go:noescape
func HasFuncSetNaturalScroll() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetNaturalScroll
//go:noescape
func FuncSetNaturalScroll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetNaturalScroll
//go:noescape
func CallSetNaturalScroll(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetNaturalScroll
//go:noescape
func TrySetNaturalScroll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetOverviewModeState
//go:noescape
func HasFuncSetOverviewModeState() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetOverviewModeState
//go:noescape
func FuncSetOverviewModeState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetOverviewModeState
//go:noescape
func CallSetOverviewModeState(
	retPtr unsafe.Pointer,
	start js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetOverviewModeState
//go:noescape
func TrySetOverviewModeState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	start js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetPlayStoreEnabled
//go:noescape
func HasFuncSetPlayStoreEnabled() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetPlayStoreEnabled
//go:noescape
func FuncSetPlayStoreEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetPlayStoreEnabled
//go:noescape
func CallSetPlayStoreEnabled(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetPlayStoreEnabled
//go:noescape
func TrySetPlayStoreEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetPluginVMPolicy
//go:noescape
func HasFuncSetPluginVMPolicy() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetPluginVMPolicy
//go:noescape
func FuncSetPluginVMPolicy(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetPluginVMPolicy
//go:noescape
func CallSetPluginVMPolicy(
	retPtr unsafe.Pointer,
	imageUrl js.Ref,
	imageHash js.Ref,
	licenseKey js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetPluginVMPolicy
//go:noescape
func TrySetPluginVMPolicy(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	imageUrl js.Ref,
	imageHash js.Ref,
	licenseKey js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetPrimaryButtonRight
//go:noescape
func HasFuncSetPrimaryButtonRight() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetPrimaryButtonRight
//go:noescape
func FuncSetPrimaryButtonRight(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetPrimaryButtonRight
//go:noescape
func CallSetPrimaryButtonRight(
	retPtr unsafe.Pointer,
	right js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetPrimaryButtonRight
//go:noescape
func TrySetPrimaryButtonRight(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	right js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetShelfAlignment
//go:noescape
func HasFuncSetShelfAlignment() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetShelfAlignment
//go:noescape
func FuncSetShelfAlignment(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetShelfAlignment
//go:noescape
func CallSetShelfAlignment(
	retPtr unsafe.Pointer,
	displayId js.Ref,
	alignment uint32)

//go:wasmimport plat/js/webext/autotestprivate try_SetShelfAlignment
//go:noescape
func TrySetShelfAlignment(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	displayId js.Ref,
	alignment uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetShelfAutoHideBehavior
//go:noescape
func HasFuncSetShelfAutoHideBehavior() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetShelfAutoHideBehavior
//go:noescape
func FuncSetShelfAutoHideBehavior(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetShelfAutoHideBehavior
//go:noescape
func CallSetShelfAutoHideBehavior(
	retPtr unsafe.Pointer,
	displayId js.Ref,
	behavior js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetShelfAutoHideBehavior
//go:noescape
func TrySetShelfAutoHideBehavior(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	displayId js.Ref,
	behavior js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetShelfIconPin
//go:noescape
func HasFuncSetShelfIconPin() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetShelfIconPin
//go:noescape
func FuncSetShelfIconPin(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetShelfIconPin
//go:noescape
func CallSetShelfIconPin(
	retPtr unsafe.Pointer,
	updateParams js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetShelfIconPin
//go:noescape
func TrySetShelfIconPin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	updateParams js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetTabletModeEnabled
//go:noescape
func HasFuncSetTabletModeEnabled() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetTabletModeEnabled
//go:noescape
func FuncSetTabletModeEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetTabletModeEnabled
//go:noescape
func CallSetTabletModeEnabled(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetTabletModeEnabled
//go:noescape
func TrySetTabletModeEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetTapDragging
//go:noescape
func HasFuncSetTapDragging() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetTapDragging
//go:noescape
func FuncSetTapDragging(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetTapDragging
//go:noescape
func CallSetTapDragging(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetTapDragging
//go:noescape
func TrySetTapDragging(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetTapToClick
//go:noescape
func HasFuncSetTapToClick() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetTapToClick
//go:noescape
func FuncSetTapToClick(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetTapToClick
//go:noescape
func CallSetTapToClick(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetTapToClick
//go:noescape
func TrySetTapToClick(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetThreeFingerClick
//go:noescape
func HasFuncSetThreeFingerClick() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetThreeFingerClick
//go:noescape
func FuncSetThreeFingerClick(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetThreeFingerClick
//go:noescape
func CallSetThreeFingerClick(
	retPtr unsafe.Pointer,
	enabled js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetThreeFingerClick
//go:noescape
func TrySetThreeFingerClick(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	enabled js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetTouchpadSensitivity
//go:noescape
func HasFuncSetTouchpadSensitivity() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetTouchpadSensitivity
//go:noescape
func FuncSetTouchpadSensitivity(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetTouchpadSensitivity
//go:noescape
func CallSetTouchpadSensitivity(
	retPtr unsafe.Pointer,
	value int32)

//go:wasmimport plat/js/webext/autotestprivate try_SetTouchpadSensitivity
//go:noescape
func TrySetTouchpadSensitivity(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetWhitelistedPref
//go:noescape
func HasFuncSetWhitelistedPref() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetWhitelistedPref
//go:noescape
func FuncSetWhitelistedPref(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetWhitelistedPref
//go:noescape
func CallSetWhitelistedPref(
	retPtr unsafe.Pointer,
	pref_name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetWhitelistedPref
//go:noescape
func TrySetWhitelistedPref(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	pref_name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SetWindowBounds
//go:noescape
func HasFuncSetWindowBounds() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SetWindowBounds
//go:noescape
func FuncSetWindowBounds(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SetWindowBounds
//go:noescape
func CallSetWindowBounds(
	retPtr unsafe.Pointer,
	id int32,
	bounds unsafe.Pointer,
	displayId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_SetWindowBounds
//go:noescape
func TrySetWindowBounds(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id int32,
	bounds unsafe.Pointer,
	displayId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_ShowPluginVMInstaller
//go:noescape
func HasFuncShowPluginVMInstaller() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ShowPluginVMInstaller
//go:noescape
func FuncShowPluginVMInstaller(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ShowPluginVMInstaller
//go:noescape
func CallShowPluginVMInstaller(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_ShowPluginVMInstaller
//go:noescape
func TryShowPluginVMInstaller(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_ShowVirtualKeyboardIfEnabled
//go:noescape
func HasFuncShowVirtualKeyboardIfEnabled() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_ShowVirtualKeyboardIfEnabled
//go:noescape
func FuncShowVirtualKeyboardIfEnabled(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_ShowVirtualKeyboardIfEnabled
//go:noescape
func CallShowVirtualKeyboardIfEnabled(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_ShowVirtualKeyboardIfEnabled
//go:noescape
func TryShowVirtualKeyboardIfEnabled(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_Shutdown
//go:noescape
func HasFuncShutdown() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_Shutdown
//go:noescape
func FuncShutdown(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_Shutdown
//go:noescape
func CallShutdown(
	retPtr unsafe.Pointer,
	force js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_Shutdown
//go:noescape
func TryShutdown(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	force js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SimulateAsanMemoryBug
//go:noescape
func HasFuncSimulateAsanMemoryBug() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SimulateAsanMemoryBug
//go:noescape
func FuncSimulateAsanMemoryBug(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SimulateAsanMemoryBug
//go:noescape
func CallSimulateAsanMemoryBug(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_SimulateAsanMemoryBug
//go:noescape
func TrySimulateAsanMemoryBug(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_StartArc
//go:noescape
func HasFuncStartArc() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_StartArc
//go:noescape
func FuncStartArc(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_StartArc
//go:noescape
func CallStartArc(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_StartArc
//go:noescape
func TryStartArc(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_StartFrameCounting
//go:noescape
func HasFuncStartFrameCounting() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_StartFrameCounting
//go:noescape
func FuncStartFrameCounting(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_StartFrameCounting
//go:noescape
func CallStartFrameCounting(
	retPtr unsafe.Pointer,
	bucketSizeInSeconds int32)

//go:wasmimport plat/js/webext/autotestprivate try_StartFrameCounting
//go:noescape
func TryStartFrameCounting(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bucketSizeInSeconds int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_StartLoginEventRecorderDataCollection
//go:noescape
func HasFuncStartLoginEventRecorderDataCollection() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_StartLoginEventRecorderDataCollection
//go:noescape
func FuncStartLoginEventRecorderDataCollection(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_StartLoginEventRecorderDataCollection
//go:noescape
func CallStartLoginEventRecorderDataCollection(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_StartLoginEventRecorderDataCollection
//go:noescape
func TryStartLoginEventRecorderDataCollection(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_StartSmoothnessTracking
//go:noescape
func HasFuncStartSmoothnessTracking() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_StartSmoothnessTracking
//go:noescape
func FuncStartSmoothnessTracking(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_StartSmoothnessTracking
//go:noescape
func CallStartSmoothnessTracking(
	retPtr unsafe.Pointer,
	displayId js.Ref,
	throughputIntervalMs int32)

//go:wasmimport plat/js/webext/autotestprivate try_StartSmoothnessTracking
//go:noescape
func TryStartSmoothnessTracking(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	displayId js.Ref,
	throughputIntervalMs int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_StartThroughputTrackerDataCollection
//go:noescape
func HasFuncStartThroughputTrackerDataCollection() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_StartThroughputTrackerDataCollection
//go:noescape
func FuncStartThroughputTrackerDataCollection(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_StartThroughputTrackerDataCollection
//go:noescape
func CallStartThroughputTrackerDataCollection(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_StartThroughputTrackerDataCollection
//go:noescape
func TryStartThroughputTrackerDataCollection(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_StopArc
//go:noescape
func HasFuncStopArc() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_StopArc
//go:noescape
func FuncStopArc(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_StopArc
//go:noescape
func CallStopArc(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_StopArc
//go:noescape
func TryStopArc(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_StopFrameCounting
//go:noescape
func HasFuncStopFrameCounting() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_StopFrameCounting
//go:noescape
func FuncStopFrameCounting(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_StopFrameCounting
//go:noescape
func CallStopFrameCounting(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_StopFrameCounting
//go:noescape
func TryStopFrameCounting(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_StopSmoothnessTracking
//go:noescape
func HasFuncStopSmoothnessTracking() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_StopSmoothnessTracking
//go:noescape
func FuncStopSmoothnessTracking(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_StopSmoothnessTracking
//go:noescape
func CallStopSmoothnessTracking(
	retPtr unsafe.Pointer,
	displayId js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_StopSmoothnessTracking
//go:noescape
func TryStopSmoothnessTracking(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	displayId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_StopThroughputTrackerDataCollection
//go:noescape
func HasFuncStopThroughputTrackerDataCollection() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_StopThroughputTrackerDataCollection
//go:noescape
func FuncStopThroughputTrackerDataCollection(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_StopThroughputTrackerDataCollection
//go:noescape
func CallStopThroughputTrackerDataCollection(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_StopThroughputTrackerDataCollection
//go:noescape
func TryStopThroughputTrackerDataCollection(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_SwapWindowsInSplitView
//go:noescape
func HasFuncSwapWindowsInSplitView() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_SwapWindowsInSplitView
//go:noescape
func FuncSwapWindowsInSplitView(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_SwapWindowsInSplitView
//go:noescape
func CallSwapWindowsInSplitView(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_SwapWindowsInSplitView
//go:noescape
func TrySwapWindowsInSplitView(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_TakeScreenshot
//go:noescape
func HasFuncTakeScreenshot() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_TakeScreenshot
//go:noescape
func FuncTakeScreenshot(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_TakeScreenshot
//go:noescape
func CallTakeScreenshot(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_TakeScreenshot
//go:noescape
func TryTakeScreenshot(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_TakeScreenshotForDisplay
//go:noescape
func HasFuncTakeScreenshotForDisplay() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_TakeScreenshotForDisplay
//go:noescape
func FuncTakeScreenshotForDisplay(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_TakeScreenshotForDisplay
//go:noescape
func CallTakeScreenshotForDisplay(
	retPtr unsafe.Pointer,
	display_id js.Ref)

//go:wasmimport plat/js/webext/autotestprivate try_TakeScreenshotForDisplay
//go:noescape
func TryTakeScreenshotForDisplay(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	display_id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_UpdatePrinter
//go:noescape
func HasFuncUpdatePrinter() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_UpdatePrinter
//go:noescape
func FuncUpdatePrinter(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_UpdatePrinter
//go:noescape
func CallUpdatePrinter(
	retPtr unsafe.Pointer,
	printer unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_UpdatePrinter
//go:noescape
func TryUpdatePrinter(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	printer unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_WaitForAmbientPhotoAnimation
//go:noescape
func HasFuncWaitForAmbientPhotoAnimation() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_WaitForAmbientPhotoAnimation
//go:noescape
func FuncWaitForAmbientPhotoAnimation(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_WaitForAmbientPhotoAnimation
//go:noescape
func CallWaitForAmbientPhotoAnimation(
	retPtr unsafe.Pointer,
	numCompletions int32,
	timeout int32)

//go:wasmimport plat/js/webext/autotestprivate try_WaitForAmbientPhotoAnimation
//go:noescape
func TryWaitForAmbientPhotoAnimation(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	numCompletions int32,
	timeout int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_WaitForAmbientVideo
//go:noescape
func HasFuncWaitForAmbientVideo() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_WaitForAmbientVideo
//go:noescape
func FuncWaitForAmbientVideo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_WaitForAmbientVideo
//go:noescape
func CallWaitForAmbientVideo(
	retPtr unsafe.Pointer,
	timeout int32)

//go:wasmimport plat/js/webext/autotestprivate try_WaitForAmbientVideo
//go:noescape
func TryWaitForAmbientVideo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	timeout int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_WaitForAssistantQueryStatus
//go:noescape
func HasFuncWaitForAssistantQueryStatus() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_WaitForAssistantQueryStatus
//go:noescape
func FuncWaitForAssistantQueryStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_WaitForAssistantQueryStatus
//go:noescape
func CallWaitForAssistantQueryStatus(
	retPtr unsafe.Pointer,
	timeout_s int32)

//go:wasmimport plat/js/webext/autotestprivate try_WaitForAssistantQueryStatus
//go:noescape
func TryWaitForAssistantQueryStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	timeout_s int32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_WaitForDisplayRotation
//go:noescape
func HasFuncWaitForDisplayRotation() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_WaitForDisplayRotation
//go:noescape
func FuncWaitForDisplayRotation(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_WaitForDisplayRotation
//go:noescape
func CallWaitForDisplayRotation(
	retPtr unsafe.Pointer,
	displayId js.Ref,
	rotation uint32)

//go:wasmimport plat/js/webext/autotestprivate try_WaitForDisplayRotation
//go:noescape
func TryWaitForDisplayRotation(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	displayId js.Ref,
	rotation uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_WaitForLauncherState
//go:noescape
func HasFuncWaitForLauncherState() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_WaitForLauncherState
//go:noescape
func FuncWaitForLauncherState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_WaitForLauncherState
//go:noescape
func CallWaitForLauncherState(
	retPtr unsafe.Pointer,
	launcherState uint32)

//go:wasmimport plat/js/webext/autotestprivate try_WaitForLauncherState
//go:noescape
func TryWaitForLauncherState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	launcherState uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_WaitForOverviewState
//go:noescape
func HasFuncWaitForOverviewState() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_WaitForOverviewState
//go:noescape
func FuncWaitForOverviewState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_WaitForOverviewState
//go:noescape
func CallWaitForOverviewState(
	retPtr unsafe.Pointer,
	overviewState uint32)

//go:wasmimport plat/js/webext/autotestprivate try_WaitForOverviewState
//go:noescape
func TryWaitForOverviewState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	overviewState uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/autotestprivate has_WaitForSystemWebAppsInstall
//go:noescape
func HasFuncWaitForSystemWebAppsInstall() js.Ref

//go:wasmimport plat/js/webext/autotestprivate func_WaitForSystemWebAppsInstall
//go:noescape
func FuncWaitForSystemWebAppsInstall(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate call_WaitForSystemWebAppsInstall
//go:noescape
func CallWaitForSystemWebAppsInstall(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/autotestprivate try_WaitForSystemWebAppsInstall
//go:noescape
func TryWaitForSystemWebAppsInstall(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

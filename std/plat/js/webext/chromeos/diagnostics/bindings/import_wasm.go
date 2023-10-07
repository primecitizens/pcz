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

//go:wasmimport plat/js/webext/chromeos/diagnostics constof_AcPowerStatus
//go:noescape
func ConstOfAcPowerStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/diagnostics store_CancelRoutineRequest
//go:noescape
func CancelRoutineRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_CancelRoutineRequest
//go:noescape
func CancelRoutineRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_CreateMemoryRoutineResponse
//go:noescape
func CreateMemoryRoutineResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_CreateMemoryRoutineResponse
//go:noescape
func CreateMemoryRoutineResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics constof_DiskReadRoutineType
//go:noescape
func ConstOfDiskReadRoutineType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/diagnostics constof_ExceptionReason
//go:noescape
func ConstOfExceptionReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/diagnostics store_ExceptionInfo
//go:noescape
func ExceptionInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_ExceptionInfo
//go:noescape
func ExceptionInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics constof_RoutineType
//go:noescape
func ConstOfRoutineType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/diagnostics store_GetAvailableRoutinesResponse
//go:noescape
func GetAvailableRoutinesResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_GetAvailableRoutinesResponse
//go:noescape
func GetAvailableRoutinesResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics constof_RoutineStatus
//go:noescape
func ConstOfRoutineStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/diagnostics constof_UserMessageType
//go:noescape
func ConstOfUserMessageType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/diagnostics store_GetRoutineUpdateResponse
//go:noescape
func GetRoutineUpdateResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_GetRoutineUpdateResponse
//go:noescape
func GetRoutineUpdateResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics constof_RoutineCommandType
//go:noescape
func ConstOfRoutineCommandType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/diagnostics store_GetRoutineUpdateRequest
//go:noescape
func GetRoutineUpdateRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_GetRoutineUpdateRequest
//go:noescape
func GetRoutineUpdateRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics constof_MemtesterTestItemEnum
//go:noescape
func ConstOfMemtesterTestItemEnum(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/diagnostics store_MemtesterResult
//go:noescape
func MemtesterResultJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_MemtesterResult
//go:noescape
func MemtesterResultJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_MemoryRoutineFinishedInfo
//go:noescape
func MemoryRoutineFinishedInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_MemoryRoutineFinishedInfo
//go:noescape
func MemoryRoutineFinishedInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics constof_NvmeSelfTestType
//go:noescape
func ConstOfNvmeSelfTestType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RoutineInitializedInfo
//go:noescape
func RoutineInitializedInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RoutineInitializedInfo
//go:noescape
func RoutineInitializedInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RoutineRunningInfo
//go:noescape
func RoutineRunningInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RoutineRunningInfo
//go:noescape
func RoutineRunningInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics constof_RoutineSupportStatus
//go:noescape
func ConstOfRoutineSupportStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RoutineSupportStatusInfo
//go:noescape
func RoutineSupportStatusInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RoutineSupportStatusInfo
//go:noescape
func RoutineSupportStatusInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics constof_RoutineWaitingReason
//go:noescape
func ConstOfRoutineWaitingReason(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RoutineWaitingInfo
//go:noescape
func RoutineWaitingInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RoutineWaitingInfo
//go:noescape
func RoutineWaitingInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunAcPowerRoutineRequest
//go:noescape
func RunAcPowerRoutineRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunAcPowerRoutineRequest
//go:noescape
func RunAcPowerRoutineRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunBatteryChargeRoutineRequest
//go:noescape
func RunBatteryChargeRoutineRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunBatteryChargeRoutineRequest
//go:noescape
func RunBatteryChargeRoutineRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunBatteryDischargeRoutineRequest
//go:noescape
func RunBatteryDischargeRoutineRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunBatteryDischargeRoutineRequest
//go:noescape
func RunBatteryDischargeRoutineRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunBluetoothPairingRoutineRequest
//go:noescape
func RunBluetoothPairingRoutineRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunBluetoothPairingRoutineRequest
//go:noescape
func RunBluetoothPairingRoutineRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunBluetoothScanningRoutineRequest
//go:noescape
func RunBluetoothScanningRoutineRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunBluetoothScanningRoutineRequest
//go:noescape
func RunBluetoothScanningRoutineRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunCpuRoutineRequest
//go:noescape
func RunCpuRoutineRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunCpuRoutineRequest
//go:noescape
func RunCpuRoutineRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunDiskReadRequest
//go:noescape
func RunDiskReadRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunDiskReadRequest
//go:noescape
func RunDiskReadRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunMemoryRoutineArguments
//go:noescape
func RunMemoryRoutineArgumentsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunMemoryRoutineArguments
//go:noescape
func RunMemoryRoutineArgumentsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunNvmeSelfTestRequest
//go:noescape
func RunNvmeSelfTestRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunNvmeSelfTestRequest
//go:noescape
func RunNvmeSelfTestRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunNvmeWearLevelRequest
//go:noescape
func RunNvmeWearLevelRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunNvmeWearLevelRequest
//go:noescape
func RunNvmeWearLevelRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunPowerButtonRequest
//go:noescape
func RunPowerButtonRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunPowerButtonRequest
//go:noescape
func RunPowerButtonRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunRoutineResponse
//go:noescape
func RunRoutineResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunRoutineResponse
//go:noescape
func RunRoutineResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_RunSmartctlCheckRequest
//go:noescape
func RunSmartctlCheckRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_RunSmartctlCheckRequest
//go:noescape
func RunSmartctlCheckRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics store_StartRoutineRequest
//go:noescape
func StartRoutineRequestJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics load_StartRoutineRequest
//go:noescape
func StartRoutineRequestJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics has_CancelRoutine
//go:noescape
func HasFuncCancelRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_CancelRoutine
//go:noescape
func FuncCancelRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_CancelRoutine
//go:noescape
func CallCancelRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_CancelRoutine
//go:noescape
func TryCancelRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_CreateMemoryRoutine
//go:noescape
func HasFuncCreateMemoryRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_CreateMemoryRoutine
//go:noescape
func FuncCreateMemoryRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_CreateMemoryRoutine
//go:noescape
func CallCreateMemoryRoutine(
	retPtr unsafe.Pointer,
	args unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_CreateMemoryRoutine
//go:noescape
func TryCreateMemoryRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	args unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_GetAvailableRoutines
//go:noescape
func HasFuncGetAvailableRoutines() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_GetAvailableRoutines
//go:noescape
func FuncGetAvailableRoutines(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_GetAvailableRoutines
//go:noescape
func CallGetAvailableRoutines(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_GetAvailableRoutines
//go:noescape
func TryGetAvailableRoutines(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_GetRoutineUpdate
//go:noescape
func HasFuncGetRoutineUpdate() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_GetRoutineUpdate
//go:noescape
func FuncGetRoutineUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_GetRoutineUpdate
//go:noescape
func CallGetRoutineUpdate(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_GetRoutineUpdate
//go:noescape
func TryGetRoutineUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_IsMemoryRoutineArgumentSupported
//go:noescape
func HasFuncIsMemoryRoutineArgumentSupported() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_IsMemoryRoutineArgumentSupported
//go:noescape
func FuncIsMemoryRoutineArgumentSupported(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_IsMemoryRoutineArgumentSupported
//go:noescape
func CallIsMemoryRoutineArgumentSupported(
	retPtr unsafe.Pointer,
	args unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_IsMemoryRoutineArgumentSupported
//go:noescape
func TryIsMemoryRoutineArgumentSupported(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	args unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_OnMemoryRoutineFinished
//go:noescape
func HasFuncOnMemoryRoutineFinished() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_OnMemoryRoutineFinished
//go:noescape
func FuncOnMemoryRoutineFinished(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_OnMemoryRoutineFinished
//go:noescape
func CallOnMemoryRoutineFinished(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_OnMemoryRoutineFinished
//go:noescape
func TryOnMemoryRoutineFinished(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_OffMemoryRoutineFinished
//go:noescape
func HasFuncOffMemoryRoutineFinished() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_OffMemoryRoutineFinished
//go:noescape
func FuncOffMemoryRoutineFinished(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_OffMemoryRoutineFinished
//go:noescape
func CallOffMemoryRoutineFinished(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_OffMemoryRoutineFinished
//go:noescape
func TryOffMemoryRoutineFinished(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_HasOnMemoryRoutineFinished
//go:noescape
func HasFuncHasOnMemoryRoutineFinished() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_HasOnMemoryRoutineFinished
//go:noescape
func FuncHasOnMemoryRoutineFinished(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_HasOnMemoryRoutineFinished
//go:noescape
func CallHasOnMemoryRoutineFinished(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_HasOnMemoryRoutineFinished
//go:noescape
func TryHasOnMemoryRoutineFinished(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_OnRoutineException
//go:noescape
func HasFuncOnRoutineException() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_OnRoutineException
//go:noescape
func FuncOnRoutineException(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_OnRoutineException
//go:noescape
func CallOnRoutineException(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_OnRoutineException
//go:noescape
func TryOnRoutineException(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_OffRoutineException
//go:noescape
func HasFuncOffRoutineException() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_OffRoutineException
//go:noescape
func FuncOffRoutineException(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_OffRoutineException
//go:noescape
func CallOffRoutineException(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_OffRoutineException
//go:noescape
func TryOffRoutineException(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_HasOnRoutineException
//go:noescape
func HasFuncHasOnRoutineException() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_HasOnRoutineException
//go:noescape
func FuncHasOnRoutineException(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_HasOnRoutineException
//go:noescape
func CallHasOnRoutineException(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_HasOnRoutineException
//go:noescape
func TryHasOnRoutineException(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_OnRoutineInitialized
//go:noescape
func HasFuncOnRoutineInitialized() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_OnRoutineInitialized
//go:noescape
func FuncOnRoutineInitialized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_OnRoutineInitialized
//go:noescape
func CallOnRoutineInitialized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_OnRoutineInitialized
//go:noescape
func TryOnRoutineInitialized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_OffRoutineInitialized
//go:noescape
func HasFuncOffRoutineInitialized() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_OffRoutineInitialized
//go:noescape
func FuncOffRoutineInitialized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_OffRoutineInitialized
//go:noescape
func CallOffRoutineInitialized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_OffRoutineInitialized
//go:noescape
func TryOffRoutineInitialized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_HasOnRoutineInitialized
//go:noescape
func HasFuncHasOnRoutineInitialized() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_HasOnRoutineInitialized
//go:noescape
func FuncHasOnRoutineInitialized(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_HasOnRoutineInitialized
//go:noescape
func CallHasOnRoutineInitialized(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_HasOnRoutineInitialized
//go:noescape
func TryHasOnRoutineInitialized(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_OnRoutineRunning
//go:noescape
func HasFuncOnRoutineRunning() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_OnRoutineRunning
//go:noescape
func FuncOnRoutineRunning(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_OnRoutineRunning
//go:noescape
func CallOnRoutineRunning(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_OnRoutineRunning
//go:noescape
func TryOnRoutineRunning(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_OffRoutineRunning
//go:noescape
func HasFuncOffRoutineRunning() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_OffRoutineRunning
//go:noescape
func FuncOffRoutineRunning(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_OffRoutineRunning
//go:noescape
func CallOffRoutineRunning(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_OffRoutineRunning
//go:noescape
func TryOffRoutineRunning(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_HasOnRoutineRunning
//go:noescape
func HasFuncHasOnRoutineRunning() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_HasOnRoutineRunning
//go:noescape
func FuncHasOnRoutineRunning(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_HasOnRoutineRunning
//go:noescape
func CallHasOnRoutineRunning(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_HasOnRoutineRunning
//go:noescape
func TryHasOnRoutineRunning(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_OnRoutineWaiting
//go:noescape
func HasFuncOnRoutineWaiting() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_OnRoutineWaiting
//go:noescape
func FuncOnRoutineWaiting(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_OnRoutineWaiting
//go:noescape
func CallOnRoutineWaiting(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_OnRoutineWaiting
//go:noescape
func TryOnRoutineWaiting(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_OffRoutineWaiting
//go:noescape
func HasFuncOffRoutineWaiting() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_OffRoutineWaiting
//go:noescape
func FuncOffRoutineWaiting(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_OffRoutineWaiting
//go:noescape
func CallOffRoutineWaiting(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_OffRoutineWaiting
//go:noescape
func TryOffRoutineWaiting(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_HasOnRoutineWaiting
//go:noescape
func HasFuncHasOnRoutineWaiting() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_HasOnRoutineWaiting
//go:noescape
func FuncHasOnRoutineWaiting(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_HasOnRoutineWaiting
//go:noescape
func CallHasOnRoutineWaiting(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_HasOnRoutineWaiting
//go:noescape
func TryHasOnRoutineWaiting(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunAcPowerRoutine
//go:noescape
func HasFuncRunAcPowerRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunAcPowerRoutine
//go:noescape
func FuncRunAcPowerRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunAcPowerRoutine
//go:noescape
func CallRunAcPowerRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunAcPowerRoutine
//go:noescape
func TryRunAcPowerRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunAudioDriverRoutine
//go:noescape
func HasFuncRunAudioDriverRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunAudioDriverRoutine
//go:noescape
func FuncRunAudioDriverRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunAudioDriverRoutine
//go:noescape
func CallRunAudioDriverRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunAudioDriverRoutine
//go:noescape
func TryRunAudioDriverRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunBatteryCapacityRoutine
//go:noescape
func HasFuncRunBatteryCapacityRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunBatteryCapacityRoutine
//go:noescape
func FuncRunBatteryCapacityRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunBatteryCapacityRoutine
//go:noescape
func CallRunBatteryCapacityRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunBatteryCapacityRoutine
//go:noescape
func TryRunBatteryCapacityRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunBatteryChargeRoutine
//go:noescape
func HasFuncRunBatteryChargeRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunBatteryChargeRoutine
//go:noescape
func FuncRunBatteryChargeRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunBatteryChargeRoutine
//go:noescape
func CallRunBatteryChargeRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunBatteryChargeRoutine
//go:noescape
func TryRunBatteryChargeRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunBatteryDischargeRoutine
//go:noescape
func HasFuncRunBatteryDischargeRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunBatteryDischargeRoutine
//go:noescape
func FuncRunBatteryDischargeRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunBatteryDischargeRoutine
//go:noescape
func CallRunBatteryDischargeRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunBatteryDischargeRoutine
//go:noescape
func TryRunBatteryDischargeRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunBatteryHealthRoutine
//go:noescape
func HasFuncRunBatteryHealthRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunBatteryHealthRoutine
//go:noescape
func FuncRunBatteryHealthRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunBatteryHealthRoutine
//go:noescape
func CallRunBatteryHealthRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunBatteryHealthRoutine
//go:noescape
func TryRunBatteryHealthRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunBluetoothDiscoveryRoutine
//go:noescape
func HasFuncRunBluetoothDiscoveryRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunBluetoothDiscoveryRoutine
//go:noescape
func FuncRunBluetoothDiscoveryRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunBluetoothDiscoveryRoutine
//go:noescape
func CallRunBluetoothDiscoveryRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunBluetoothDiscoveryRoutine
//go:noescape
func TryRunBluetoothDiscoveryRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunBluetoothPairingRoutine
//go:noescape
func HasFuncRunBluetoothPairingRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunBluetoothPairingRoutine
//go:noescape
func FuncRunBluetoothPairingRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunBluetoothPairingRoutine
//go:noescape
func CallRunBluetoothPairingRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunBluetoothPairingRoutine
//go:noescape
func TryRunBluetoothPairingRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunBluetoothPowerRoutine
//go:noescape
func HasFuncRunBluetoothPowerRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunBluetoothPowerRoutine
//go:noescape
func FuncRunBluetoothPowerRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunBluetoothPowerRoutine
//go:noescape
func CallRunBluetoothPowerRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunBluetoothPowerRoutine
//go:noescape
func TryRunBluetoothPowerRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunBluetoothScanningRoutine
//go:noescape
func HasFuncRunBluetoothScanningRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunBluetoothScanningRoutine
//go:noescape
func FuncRunBluetoothScanningRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunBluetoothScanningRoutine
//go:noescape
func CallRunBluetoothScanningRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunBluetoothScanningRoutine
//go:noescape
func TryRunBluetoothScanningRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunCpuCacheRoutine
//go:noescape
func HasFuncRunCpuCacheRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunCpuCacheRoutine
//go:noescape
func FuncRunCpuCacheRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunCpuCacheRoutine
//go:noescape
func CallRunCpuCacheRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunCpuCacheRoutine
//go:noescape
func TryRunCpuCacheRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunCpuFloatingPointAccuracyRoutine
//go:noescape
func HasFuncRunCpuFloatingPointAccuracyRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunCpuFloatingPointAccuracyRoutine
//go:noescape
func FuncRunCpuFloatingPointAccuracyRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunCpuFloatingPointAccuracyRoutine
//go:noescape
func CallRunCpuFloatingPointAccuracyRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunCpuFloatingPointAccuracyRoutine
//go:noescape
func TryRunCpuFloatingPointAccuracyRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunCpuPrimeSearchRoutine
//go:noescape
func HasFuncRunCpuPrimeSearchRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunCpuPrimeSearchRoutine
//go:noescape
func FuncRunCpuPrimeSearchRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunCpuPrimeSearchRoutine
//go:noescape
func CallRunCpuPrimeSearchRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunCpuPrimeSearchRoutine
//go:noescape
func TryRunCpuPrimeSearchRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunCpuStressRoutine
//go:noescape
func HasFuncRunCpuStressRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunCpuStressRoutine
//go:noescape
func FuncRunCpuStressRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunCpuStressRoutine
//go:noescape
func CallRunCpuStressRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunCpuStressRoutine
//go:noescape
func TryRunCpuStressRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunDiskReadRoutine
//go:noescape
func HasFuncRunDiskReadRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunDiskReadRoutine
//go:noescape
func FuncRunDiskReadRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunDiskReadRoutine
//go:noescape
func CallRunDiskReadRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunDiskReadRoutine
//go:noescape
func TryRunDiskReadRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunDnsResolutionRoutine
//go:noescape
func HasFuncRunDnsResolutionRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunDnsResolutionRoutine
//go:noescape
func FuncRunDnsResolutionRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunDnsResolutionRoutine
//go:noescape
func CallRunDnsResolutionRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunDnsResolutionRoutine
//go:noescape
func TryRunDnsResolutionRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunDnsResolverPresentRoutine
//go:noescape
func HasFuncRunDnsResolverPresentRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunDnsResolverPresentRoutine
//go:noescape
func FuncRunDnsResolverPresentRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunDnsResolverPresentRoutine
//go:noescape
func CallRunDnsResolverPresentRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunDnsResolverPresentRoutine
//go:noescape
func TryRunDnsResolverPresentRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunEmmcLifetimeRoutine
//go:noescape
func HasFuncRunEmmcLifetimeRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunEmmcLifetimeRoutine
//go:noescape
func FuncRunEmmcLifetimeRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunEmmcLifetimeRoutine
//go:noescape
func CallRunEmmcLifetimeRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunEmmcLifetimeRoutine
//go:noescape
func TryRunEmmcLifetimeRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunFingerprintAliveRoutine
//go:noescape
func HasFuncRunFingerprintAliveRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunFingerprintAliveRoutine
//go:noescape
func FuncRunFingerprintAliveRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunFingerprintAliveRoutine
//go:noescape
func CallRunFingerprintAliveRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunFingerprintAliveRoutine
//go:noescape
func TryRunFingerprintAliveRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunGatewayCanBePingedRoutine
//go:noescape
func HasFuncRunGatewayCanBePingedRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunGatewayCanBePingedRoutine
//go:noescape
func FuncRunGatewayCanBePingedRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunGatewayCanBePingedRoutine
//go:noescape
func CallRunGatewayCanBePingedRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunGatewayCanBePingedRoutine
//go:noescape
func TryRunGatewayCanBePingedRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunLanConnectivityRoutine
//go:noescape
func HasFuncRunLanConnectivityRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunLanConnectivityRoutine
//go:noescape
func FuncRunLanConnectivityRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunLanConnectivityRoutine
//go:noescape
func CallRunLanConnectivityRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunLanConnectivityRoutine
//go:noescape
func TryRunLanConnectivityRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunMemoryRoutine
//go:noescape
func HasFuncRunMemoryRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunMemoryRoutine
//go:noescape
func FuncRunMemoryRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunMemoryRoutine
//go:noescape
func CallRunMemoryRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunMemoryRoutine
//go:noescape
func TryRunMemoryRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunNvmeSelfTestRoutine
//go:noescape
func HasFuncRunNvmeSelfTestRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunNvmeSelfTestRoutine
//go:noescape
func FuncRunNvmeSelfTestRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunNvmeSelfTestRoutine
//go:noescape
func CallRunNvmeSelfTestRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunNvmeSelfTestRoutine
//go:noescape
func TryRunNvmeSelfTestRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunNvmeWearLevelRoutine
//go:noescape
func HasFuncRunNvmeWearLevelRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunNvmeWearLevelRoutine
//go:noescape
func FuncRunNvmeWearLevelRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunNvmeWearLevelRoutine
//go:noescape
func CallRunNvmeWearLevelRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunNvmeWearLevelRoutine
//go:noescape
func TryRunNvmeWearLevelRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunPowerButtonRoutine
//go:noescape
func HasFuncRunPowerButtonRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunPowerButtonRoutine
//go:noescape
func FuncRunPowerButtonRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunPowerButtonRoutine
//go:noescape
func CallRunPowerButtonRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunPowerButtonRoutine
//go:noescape
func TryRunPowerButtonRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunSensitiveSensorRoutine
//go:noescape
func HasFuncRunSensitiveSensorRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunSensitiveSensorRoutine
//go:noescape
func FuncRunSensitiveSensorRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunSensitiveSensorRoutine
//go:noescape
func CallRunSensitiveSensorRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunSensitiveSensorRoutine
//go:noescape
func TryRunSensitiveSensorRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunSignalStrengthRoutine
//go:noescape
func HasFuncRunSignalStrengthRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunSignalStrengthRoutine
//go:noescape
func FuncRunSignalStrengthRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunSignalStrengthRoutine
//go:noescape
func CallRunSignalStrengthRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunSignalStrengthRoutine
//go:noescape
func TryRunSignalStrengthRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunSmartctlCheckRoutine
//go:noescape
func HasFuncRunSmartctlCheckRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunSmartctlCheckRoutine
//go:noescape
func FuncRunSmartctlCheckRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunSmartctlCheckRoutine
//go:noescape
func CallRunSmartctlCheckRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunSmartctlCheckRoutine
//go:noescape
func TryRunSmartctlCheckRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_RunUfsLifetimeRoutine
//go:noescape
func HasFuncRunUfsLifetimeRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_RunUfsLifetimeRoutine
//go:noescape
func FuncRunUfsLifetimeRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_RunUfsLifetimeRoutine
//go:noescape
func CallRunUfsLifetimeRoutine(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_RunUfsLifetimeRoutine
//go:noescape
func TryRunUfsLifetimeRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/diagnostics has_StartRoutine
//go:noescape
func HasFuncStartRoutine() js.Ref

//go:wasmimport plat/js/webext/chromeos/diagnostics func_StartRoutine
//go:noescape
func FuncStartRoutine(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics call_StartRoutine
//go:noescape
func CallStartRoutine(
	retPtr unsafe.Pointer,
	request unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/diagnostics try_StartRoutine
//go:noescape
func TryStartRoutine(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	request unsafe.Pointer) (ok js.Ref)

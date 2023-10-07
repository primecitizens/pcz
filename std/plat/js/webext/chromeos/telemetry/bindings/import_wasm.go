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

//go:wasmimport plat/js/webext/chromeos/telemetry store_AudioOutputNodeInfo
//go:noescape
func AudioOutputNodeInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_AudioOutputNodeInfo
//go:noescape
func AudioOutputNodeInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_AudioInputNodeInfo
//go:noescape
func AudioInputNodeInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_AudioInputNodeInfo
//go:noescape
func AudioInputNodeInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_AudioInfo
//go:noescape
func AudioInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_AudioInfo
//go:noescape
func AudioInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_BatteryInfo
//go:noescape
func BatteryInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_BatteryInfo
//go:noescape
func BatteryInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry constof_CpuArchitectureEnum
//go:noescape
func ConstOfCpuArchitectureEnum(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/telemetry store_CpuCStateInfo
//go:noescape
func CpuCStateInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_CpuCStateInfo
//go:noescape
func CpuCStateInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_LogicalCpuInfo
//go:noescape
func LogicalCpuInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_LogicalCpuInfo
//go:noescape
func LogicalCpuInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_PhysicalCpuInfo
//go:noescape
func PhysicalCpuInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_PhysicalCpuInfo
//go:noescape
func PhysicalCpuInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_CpuInfo
//go:noescape
func CpuInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_CpuInfo
//go:noescape
func CpuInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry constof_DisplayInputType
//go:noescape
func ConstOfDisplayInputType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/telemetry store_EmbeddedDisplayInfo
//go:noescape
func EmbeddedDisplayInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_EmbeddedDisplayInfo
//go:noescape
func EmbeddedDisplayInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_ExternalDisplayInfo
//go:noescape
func ExternalDisplayInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_ExternalDisplayInfo
//go:noescape
func ExternalDisplayInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_DisplayInfo
//go:noescape
func DisplayInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_DisplayInfo
//go:noescape
func DisplayInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry constof_FwupdVersionFormat
//go:noescape
func ConstOfFwupdVersionFormat(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/telemetry store_FwupdFirmwareVersionInfo
//go:noescape
func FwupdFirmwareVersionInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_FwupdFirmwareVersionInfo
//go:noescape
func FwupdFirmwareVersionInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry constof_NetworkType
//go:noescape
func ConstOfNetworkType(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/telemetry constof_NetworkState
//go:noescape
func ConstOfNetworkState(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/telemetry store_NetworkInfo
//go:noescape
func NetworkInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_NetworkInfo
//go:noescape
func NetworkInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_InternetConnectivityInfo
//go:noescape
func InternetConnectivityInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_InternetConnectivityInfo
//go:noescape
func InternetConnectivityInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_MarketingInfo
//go:noescape
func MarketingInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_MarketingInfo
//go:noescape
func MarketingInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_MemoryInfo
//go:noescape
func MemoryInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_MemoryInfo
//go:noescape
func MemoryInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_NonRemovableBlockDeviceInfo
//go:noescape
func NonRemovableBlockDeviceInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_NonRemovableBlockDeviceInfo
//go:noescape
func NonRemovableBlockDeviceInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_NonRemovableBlockDeviceInfoResponse
//go:noescape
func NonRemovableBlockDeviceInfoResponseJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_NonRemovableBlockDeviceInfoResponse
//go:noescape
func NonRemovableBlockDeviceInfoResponseJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_OemData
//go:noescape
func OemDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_OemData
//go:noescape
func OemDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_OsVersionInfo
//go:noescape
func OsVersionInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_OsVersionInfo
//go:noescape
func OsVersionInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_StatefulPartitionInfo
//go:noescape
func StatefulPartitionInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_StatefulPartitionInfo
//go:noescape
func StatefulPartitionInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_TpmDictionaryAttack
//go:noescape
func TpmDictionaryAttackJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_TpmDictionaryAttack
//go:noescape
func TpmDictionaryAttackJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry constof_TpmGSCVersion
//go:noescape
func ConstOfTpmGSCVersion(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/telemetry store_TpmVersion
//go:noescape
func TpmVersionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_TpmVersion
//go:noescape
func TpmVersionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_TpmStatus
//go:noescape
func TpmStatusJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_TpmStatus
//go:noescape
func TpmStatusJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_TpmInfo
//go:noescape
func TpmInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_TpmInfo
//go:noescape
func TpmInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_UsbBusInterfaceInfo
//go:noescape
func UsbBusInterfaceInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_UsbBusInterfaceInfo
//go:noescape
func UsbBusInterfaceInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry constof_UsbVersion
//go:noescape
func ConstOfUsbVersion(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/telemetry constof_UsbSpecSpeed
//go:noescape
func ConstOfUsbSpecSpeed(str js.Ref) uint32

//go:wasmimport plat/js/webext/chromeos/telemetry store_UsbBusInfo
//go:noescape
func UsbBusInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_UsbBusInfo
//go:noescape
func UsbBusInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_UsbBusDevices
//go:noescape
func UsbBusDevicesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_UsbBusDevices
//go:noescape
func UsbBusDevicesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry store_VpdInfo
//go:noescape
func VpdInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry load_VpdInfo
//go:noescape
func VpdInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetAudioInfo
//go:noescape
func HasFuncGetAudioInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetAudioInfo
//go:noescape
func FuncGetAudioInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetAudioInfo
//go:noescape
func CallGetAudioInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetAudioInfo
//go:noescape
func TryGetAudioInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetBatteryInfo
//go:noescape
func HasFuncGetBatteryInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetBatteryInfo
//go:noescape
func FuncGetBatteryInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetBatteryInfo
//go:noescape
func CallGetBatteryInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetBatteryInfo
//go:noescape
func TryGetBatteryInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetCpuInfo
//go:noescape
func HasFuncGetCpuInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetCpuInfo
//go:noescape
func FuncGetCpuInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetCpuInfo
//go:noescape
func CallGetCpuInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetCpuInfo
//go:noescape
func TryGetCpuInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetDisplayInfo
//go:noescape
func HasFuncGetDisplayInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetDisplayInfo
//go:noescape
func FuncGetDisplayInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetDisplayInfo
//go:noescape
func CallGetDisplayInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetDisplayInfo
//go:noescape
func TryGetDisplayInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetInternetConnectivityInfo
//go:noescape
func HasFuncGetInternetConnectivityInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetInternetConnectivityInfo
//go:noescape
func FuncGetInternetConnectivityInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetInternetConnectivityInfo
//go:noescape
func CallGetInternetConnectivityInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetInternetConnectivityInfo
//go:noescape
func TryGetInternetConnectivityInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetMarketingInfo
//go:noescape
func HasFuncGetMarketingInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetMarketingInfo
//go:noescape
func FuncGetMarketingInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetMarketingInfo
//go:noescape
func CallGetMarketingInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetMarketingInfo
//go:noescape
func TryGetMarketingInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetMemoryInfo
//go:noescape
func HasFuncGetMemoryInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetMemoryInfo
//go:noescape
func FuncGetMemoryInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetMemoryInfo
//go:noescape
func CallGetMemoryInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetMemoryInfo
//go:noescape
func TryGetMemoryInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetNonRemovableBlockDevicesInfo
//go:noescape
func HasFuncGetNonRemovableBlockDevicesInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetNonRemovableBlockDevicesInfo
//go:noescape
func FuncGetNonRemovableBlockDevicesInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetNonRemovableBlockDevicesInfo
//go:noescape
func CallGetNonRemovableBlockDevicesInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetNonRemovableBlockDevicesInfo
//go:noescape
func TryGetNonRemovableBlockDevicesInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetOemData
//go:noescape
func HasFuncGetOemData() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetOemData
//go:noescape
func FuncGetOemData(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetOemData
//go:noescape
func CallGetOemData(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetOemData
//go:noescape
func TryGetOemData(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetOsVersionInfo
//go:noescape
func HasFuncGetOsVersionInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetOsVersionInfo
//go:noescape
func FuncGetOsVersionInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetOsVersionInfo
//go:noescape
func CallGetOsVersionInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetOsVersionInfo
//go:noescape
func TryGetOsVersionInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetStatefulPartitionInfo
//go:noescape
func HasFuncGetStatefulPartitionInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetStatefulPartitionInfo
//go:noescape
func FuncGetStatefulPartitionInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetStatefulPartitionInfo
//go:noescape
func CallGetStatefulPartitionInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetStatefulPartitionInfo
//go:noescape
func TryGetStatefulPartitionInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetTpmInfo
//go:noescape
func HasFuncGetTpmInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetTpmInfo
//go:noescape
func FuncGetTpmInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetTpmInfo
//go:noescape
func CallGetTpmInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetTpmInfo
//go:noescape
func TryGetTpmInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetUsbBusInfo
//go:noescape
func HasFuncGetUsbBusInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetUsbBusInfo
//go:noescape
func FuncGetUsbBusInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetUsbBusInfo
//go:noescape
func CallGetUsbBusInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetUsbBusInfo
//go:noescape
func TryGetUsbBusInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/chromeos/telemetry has_GetVpdInfo
//go:noescape
func HasFuncGetVpdInfo() js.Ref

//go:wasmimport plat/js/webext/chromeos/telemetry func_GetVpdInfo
//go:noescape
func FuncGetVpdInfo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry call_GetVpdInfo
//go:noescape
func CallGetVpdInfo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/chromeos/telemetry try_GetVpdInfo
//go:noescape
func TryGetVpdInfo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

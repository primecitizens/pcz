// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web has_MLCommandEncoder_Dispatch
//go:noescape
func HasMLCommandEncoderDispatch(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLCommandEncoder_Dispatch
//go:noescape
func MLCommandEncoderDispatchFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MLCommandEncoder_Dispatch
//go:noescape
func CallMLCommandEncoderDispatch(
	this js.Ref, retPtr unsafe.Pointer,
	graph js.Ref,
	inputs js.Ref,
	outputs js.Ref)

//go:wasmimport plat/js/web try_MLCommandEncoder_Dispatch
//go:noescape
func TryMLCommandEncoderDispatch(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	graph js.Ref,
	inputs js.Ref,
	outputs js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLCommandEncoder_Finish
//go:noescape
func HasMLCommandEncoderFinish(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLCommandEncoder_Finish
//go:noescape
func MLCommandEncoderFinishFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MLCommandEncoder_Finish
//go:noescape
func CallMLCommandEncoderFinish(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_MLCommandEncoder_Finish
//go:noescape
func TryMLCommandEncoderFinish(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLCommandEncoder_Finish1
//go:noescape
func HasMLCommandEncoderFinish1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLCommandEncoder_Finish1
//go:noescape
func MLCommandEncoderFinish1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MLCommandEncoder_Finish1
//go:noescape
func CallMLCommandEncoderFinish1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLCommandEncoder_Finish1
//go:noescape
func TryMLCommandEncoderFinish1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MLCommandEncoder_InitializeGraph
//go:noescape
func HasMLCommandEncoderInitializeGraph(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLCommandEncoder_InitializeGraph
//go:noescape
func MLCommandEncoderInitializeGraphFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MLCommandEncoder_InitializeGraph
//go:noescape
func CallMLCommandEncoderInitializeGraph(
	this js.Ref, retPtr unsafe.Pointer,
	graph js.Ref)

//go:wasmimport plat/js/web try_MLCommandEncoder_InitializeGraph
//go:noescape
func TryMLCommandEncoderInitializeGraph(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	graph js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLContext_Compute
//go:noescape
func HasMLContextCompute(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLContext_Compute
//go:noescape
func MLContextComputeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MLContext_Compute
//go:noescape
func CallMLContextCompute(
	this js.Ref, retPtr unsafe.Pointer,
	graph js.Ref,
	inputs js.Ref,
	outputs js.Ref)

//go:wasmimport plat/js/web try_MLContext_Compute
//go:noescape
func TryMLContextCompute(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	graph js.Ref,
	inputs js.Ref,
	outputs js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLContext_ComputeSync
//go:noescape
func HasMLContextComputeSync(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLContext_ComputeSync
//go:noescape
func MLContextComputeSyncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MLContext_ComputeSync
//go:noescape
func CallMLContextComputeSync(
	this js.Ref, retPtr unsafe.Pointer,
	graph js.Ref,
	inputs js.Ref,
	outputs js.Ref)

//go:wasmimport plat/js/web try_MLContext_ComputeSync
//go:noescape
func TryMLContextComputeSync(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	graph js.Ref,
	inputs js.Ref,
	outputs js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MLContext_CreateCommandEncoder
//go:noescape
func HasMLContextCreateCommandEncoder(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MLContext_CreateCommandEncoder
//go:noescape
func MLContextCreateCommandEncoderFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MLContext_CreateCommandEncoder
//go:noescape
func CallMLContextCreateCommandEncoder(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MLContext_CreateCommandEncoder
//go:noescape
func TryMLContextCreateCommandEncoder(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_MLDeviceType
//go:noescape
func ConstOfMLDeviceType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_MLPowerPreference
//go:noescape
func ConstOfMLPowerPreference(str js.Ref) uint32

//go:wasmimport plat/js/web store_MLContextOptions
//go:noescape
func MLContextOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MLContextOptions
//go:noescape
func MLContextOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_ML_CreateContext
//go:noescape
func HasMLCreateContext(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ML_CreateContext
//go:noescape
func MLCreateContextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ML_CreateContext
//go:noescape
func CallMLCreateContext(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_ML_CreateContext
//go:noescape
func TryMLCreateContext(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ML_CreateContext1
//go:noescape
func HasMLCreateContext1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ML_CreateContext1
//go:noescape
func MLCreateContext1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ML_CreateContext1
//go:noescape
func CallMLCreateContext1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ML_CreateContext1
//go:noescape
func TryMLCreateContext1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ML_CreateContext2
//go:noescape
func HasMLCreateContext2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ML_CreateContext2
//go:noescape
func MLCreateContext2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ML_CreateContext2
//go:noescape
func CallMLCreateContext2(
	this js.Ref, retPtr unsafe.Pointer,
	gpuDevice js.Ref)

//go:wasmimport plat/js/web try_ML_CreateContext2
//go:noescape
func TryMLCreateContext2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	gpuDevice js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_ML_CreateContextSync
//go:noescape
func HasMLCreateContextSync(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ML_CreateContextSync
//go:noescape
func MLCreateContextSyncFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ML_CreateContextSync
//go:noescape
func CallMLCreateContextSync(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_ML_CreateContextSync
//go:noescape
func TryMLCreateContextSync(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ML_CreateContextSync1
//go:noescape
func HasMLCreateContextSync1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ML_CreateContextSync1
//go:noescape
func MLCreateContextSync1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ML_CreateContextSync1
//go:noescape
func CallMLCreateContextSync1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ML_CreateContextSync1
//go:noescape
func TryMLCreateContextSync1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ML_CreateContextSync2
//go:noescape
func HasMLCreateContextSync2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ML_CreateContextSync2
//go:noescape
func MLCreateContextSync2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ML_CreateContextSync2
//go:noescape
func CallMLCreateContextSync2(
	this js.Ref, retPtr unsafe.Pointer,
	gpuDevice js.Ref)

//go:wasmimport plat/js/web try_ML_CreateContextSync2
//go:noescape
func TryMLCreateContextSync2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	gpuDevice js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_ConnectionType
//go:noescape
func ConstOfConnectionType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_EffectiveConnectionType
//go:noescape
func ConstOfEffectiveConnectionType(str js.Ref) uint32

//go:wasmimport plat/js/web get_NetworkInformation_Type
//go:noescape
func GetNetworkInformationType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NetworkInformation_EffectiveType
//go:noescape
func GetNetworkInformationEffectiveType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NetworkInformation_DownlinkMax
//go:noescape
func GetNetworkInformationDownlinkMax(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NetworkInformation_Downlink
//go:noescape
func GetNetworkInformationDownlink(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NetworkInformation_Rtt
//go:noescape
func GetNetworkInformationRtt(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_NetworkInformation_SaveData
//go:noescape
func GetNetworkInformationSaveData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_GPUFeatureName
//go:noescape
func ConstOfGPUFeatureName(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPUQueueDescriptor
//go:noescape
func GPUQueueDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUQueueDescriptor
//go:noescape
func GPUQueueDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GPUDeviceDescriptor
//go:noescape
func GPUDeviceDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPUDeviceDescriptor
//go:noescape
func GPUDeviceDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPUAdapterInfo_Vendor
//go:noescape
func GetGPUAdapterInfoVendor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUAdapterInfo_Architecture
//go:noescape
func GetGPUAdapterInfoArchitecture(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUAdapterInfo_Device
//go:noescape
func GetGPUAdapterInfoDevice(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUAdapterInfo_Description
//go:noescape
func GetGPUAdapterInfoDescription(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUAdapter_Features
//go:noescape
func GetGPUAdapterFeatures(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUAdapter_Limits
//go:noescape
func GetGPUAdapterLimits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GPUAdapter_IsFallbackAdapter
//go:noescape
func GetGPUAdapterIsFallbackAdapter(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUAdapter_RequestDevice
//go:noescape
func HasGPUAdapterRequestDevice(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUAdapter_RequestDevice
//go:noescape
func GPUAdapterRequestDeviceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUAdapter_RequestDevice
//go:noescape
func CallGPUAdapterRequestDevice(
	this js.Ref, retPtr unsafe.Pointer,
	descriptor unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUAdapter_RequestDevice
//go:noescape
func TryGPUAdapterRequestDevice(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	descriptor unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUAdapter_RequestDevice1
//go:noescape
func HasGPUAdapterRequestDevice1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUAdapter_RequestDevice1
//go:noescape
func GPUAdapterRequestDevice1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUAdapter_RequestDevice1
//go:noescape
func CallGPUAdapterRequestDevice1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUAdapter_RequestDevice1
//go:noescape
func TryGPUAdapterRequestDevice1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUAdapter_RequestAdapterInfo
//go:noescape
func HasGPUAdapterRequestAdapterInfo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUAdapter_RequestAdapterInfo
//go:noescape
func GPUAdapterRequestAdapterInfoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUAdapter_RequestAdapterInfo
//go:noescape
func CallGPUAdapterRequestAdapterInfo(
	this js.Ref, retPtr unsafe.Pointer,
	unmaskHints js.Ref)

//go:wasmimport plat/js/web try_GPUAdapter_RequestAdapterInfo
//go:noescape
func TryGPUAdapterRequestAdapterInfo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	unmaskHints js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_GPUAdapter_RequestAdapterInfo1
//go:noescape
func HasGPUAdapterRequestAdapterInfo1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPUAdapter_RequestAdapterInfo1
//go:noescape
func GPUAdapterRequestAdapterInfo1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPUAdapter_RequestAdapterInfo1
//go:noescape
func CallGPUAdapterRequestAdapterInfo1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPUAdapter_RequestAdapterInfo1
//go:noescape
func TryGPUAdapterRequestAdapterInfo1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_GPUPowerPreference
//go:noescape
func ConstOfGPUPowerPreference(str js.Ref) uint32

//go:wasmimport plat/js/web store_GPURequestAdapterOptions
//go:noescape
func GPURequestAdapterOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GPURequestAdapterOptions
//go:noescape
func GPURequestAdapterOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_GPU_WgslLanguageFeatures
//go:noescape
func GetGPUWgslLanguageFeatures(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPU_RequestAdapter
//go:noescape
func HasGPURequestAdapter(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPU_RequestAdapter
//go:noescape
func GPURequestAdapterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPU_RequestAdapter
//go:noescape
func CallGPURequestAdapter(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_GPU_RequestAdapter
//go:noescape
func TryGPURequestAdapter(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPU_RequestAdapter1
//go:noescape
func HasGPURequestAdapter1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPU_RequestAdapter1
//go:noescape
func GPURequestAdapter1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPU_RequestAdapter1
//go:noescape
func CallGPURequestAdapter1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPU_RequestAdapter1
//go:noescape
func TryGPURequestAdapter1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_GPU_GetPreferredCanvasFormat
//go:noescape
func HasGPUGetPreferredCanvasFormat(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_GPU_GetPreferredCanvasFormat
//go:noescape
func GPUGetPreferredCanvasFormatFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_GPU_GetPreferredCanvasFormat
//go:noescape
func CallGPUGetPreferredCanvasFormat(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_GPU_GetPreferredCanvasFormat
//go:noescape
func TryGPUGetPreferredCanvasFormat(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Hid
//go:noescape
func GetNavigatorHid(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_WindowControlsOverlay
//go:noescape
func GetNavigatorWindowControlsOverlay(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Credentials
//go:noescape
func GetNavigatorCredentials(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Clipboard
//go:noescape
func GetNavigatorClipboard(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Geolocation
//go:noescape
func GetNavigatorGeolocation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Usb
//go:noescape
func GetNavigatorUsb(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_EpubReadingSystem
//go:noescape
func GetNavigatorEpubReadingSystem(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Xr
//go:noescape
func GetNavigatorXr(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_ServiceWorker
//go:noescape
func GetNavigatorServiceWorker(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_MediaDevices
//go:noescape
func GetNavigatorMediaDevices(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Bluetooth
//go:noescape
func GetNavigatorBluetooth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Serial
//go:noescape
func GetNavigatorSerial(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_MediaCapabilities
//go:noescape
func GetNavigatorMediaCapabilities(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_UserActivation
//go:noescape
func GetNavigatorUserActivation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Permissions
//go:noescape
func GetNavigatorPermissions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Contacts
//go:noescape
func GetNavigatorContacts(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Keyboard
//go:noescape
func GetNavigatorKeyboard(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_MediaSession
//go:noescape
func GetNavigatorMediaSession(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_DevicePosture
//go:noescape
func GetNavigatorDevicePosture(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_MaxTouchPoints
//go:noescape
func GetNavigatorMaxTouchPoints(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Scheduling
//go:noescape
func GetNavigatorScheduling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_WakeLock
//go:noescape
func GetNavigatorWakeLock(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Ink
//go:noescape
func GetNavigatorInk(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Presentation
//go:noescape
func GetNavigatorPresentation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_VirtualKeyboard
//go:noescape
func GetNavigatorVirtualKeyboard(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_UserAgentData
//go:noescape
func GetNavigatorUserAgentData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_HardwareConcurrency
//go:noescape
func GetNavigatorHardwareConcurrency(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_DeviceMemory
//go:noescape
func GetNavigatorDeviceMemory(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Storage
//go:noescape
func GetNavigatorStorage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_StorageBuckets
//go:noescape
func GetNavigatorStorageBuckets(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Locks
//go:noescape
func GetNavigatorLocks(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_AppCodeName
//go:noescape
func GetNavigatorAppCodeName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_AppName
//go:noescape
func GetNavigatorAppName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_AppVersion
//go:noescape
func GetNavigatorAppVersion(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Platform
//go:noescape
func GetNavigatorPlatform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Product
//go:noescape
func GetNavigatorProduct(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_ProductSub
//go:noescape
func GetNavigatorProductSub(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_UserAgent
//go:noescape
func GetNavigatorUserAgent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Vendor
//go:noescape
func GetNavigatorVendor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_VendorSub
//go:noescape
func GetNavigatorVendorSub(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Oscpu
//go:noescape
func GetNavigatorOscpu(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Language
//go:noescape
func GetNavigatorLanguage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Languages
//go:noescape
func GetNavigatorLanguages(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_OnLine
//go:noescape
func GetNavigatorOnLine(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_CookieEnabled
//go:noescape
func GetNavigatorCookieEnabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Plugins
//go:noescape
func GetNavigatorPlugins(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_MimeTypes
//go:noescape
func GetNavigatorMimeTypes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_PdfViewerEnabled
//go:noescape
func GetNavigatorPdfViewerEnabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Webdriver
//go:noescape
func GetNavigatorWebdriver(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Ml
//go:noescape
func GetNavigatorMl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Connection
//go:noescape
func GetNavigatorConnection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Navigator_Gpu
//go:noescape
func GetNavigatorGpu(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_UpdateAdInterestGroups
//go:noescape
func HasNavigatorUpdateAdInterestGroups(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_UpdateAdInterestGroups
//go:noescape
func NavigatorUpdateAdInterestGroupsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_UpdateAdInterestGroups
//go:noescape
func CallNavigatorUpdateAdInterestGroups(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_UpdateAdInterestGroups
//go:noescape
func TryNavigatorUpdateAdInterestGroups(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_SendBeacon
//go:noescape
func HasNavigatorSendBeacon(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_SendBeacon
//go:noescape
func NavigatorSendBeaconFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_SendBeacon
//go:noescape
func CallNavigatorSendBeacon(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref,
	data js.Ref)

//go:wasmimport plat/js/web try_Navigator_SendBeacon
//go:noescape
func TryNavigatorSendBeacon(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_SendBeacon1
//go:noescape
func HasNavigatorSendBeacon1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_SendBeacon1
//go:noescape
func NavigatorSendBeacon1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_SendBeacon1
//go:noescape
func CallNavigatorSendBeacon1(
	this js.Ref, retPtr unsafe.Pointer,
	url js.Ref)

//go:wasmimport plat/js/web try_Navigator_SendBeacon1
//go:noescape
func TryNavigatorSendBeacon1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_GetBattery
//go:noescape
func HasNavigatorGetBattery(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_GetBattery
//go:noescape
func NavigatorGetBatteryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_GetBattery
//go:noescape
func CallNavigatorGetBattery(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_GetBattery
//go:noescape
func TryNavigatorGetBattery(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_Vibrate
//go:noescape
func HasNavigatorVibrate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_Vibrate
//go:noescape
func NavigatorVibrateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_Vibrate
//go:noescape
func CallNavigatorVibrate(
	this js.Ref, retPtr unsafe.Pointer,
	pattern js.Ref)

//go:wasmimport plat/js/web try_Navigator_Vibrate
//go:noescape
func TryNavigatorVibrate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	pattern js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_GetGamepads
//go:noescape
func HasNavigatorGetGamepads(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_GetGamepads
//go:noescape
func NavigatorGetGamepadsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_GetGamepads
//go:noescape
func CallNavigatorGetGamepads(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_GetGamepads
//go:noescape
func TryNavigatorGetGamepads(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_GetInstalledRelatedApps
//go:noescape
func HasNavigatorGetInstalledRelatedApps(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_GetInstalledRelatedApps
//go:noescape
func NavigatorGetInstalledRelatedAppsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_GetInstalledRelatedApps
//go:noescape
func CallNavigatorGetInstalledRelatedApps(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_GetInstalledRelatedApps
//go:noescape
func TryNavigatorGetInstalledRelatedApps(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_RequestMediaKeySystemAccess
//go:noescape
func HasNavigatorRequestMediaKeySystemAccess(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_RequestMediaKeySystemAccess
//go:noescape
func NavigatorRequestMediaKeySystemAccessFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_RequestMediaKeySystemAccess
//go:noescape
func CallNavigatorRequestMediaKeySystemAccess(
	this js.Ref, retPtr unsafe.Pointer,
	keySystem js.Ref,
	supportedConfigurations js.Ref)

//go:wasmimport plat/js/web try_Navigator_RequestMediaKeySystemAccess
//go:noescape
func TryNavigatorRequestMediaKeySystemAccess(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keySystem js.Ref,
	supportedConfigurations js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_JoinAdInterestGroup
//go:noescape
func HasNavigatorJoinAdInterestGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_JoinAdInterestGroup
//go:noescape
func NavigatorJoinAdInterestGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_JoinAdInterestGroup
//go:noescape
func CallNavigatorJoinAdInterestGroup(
	this js.Ref, retPtr unsafe.Pointer,
	group unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_JoinAdInterestGroup
//go:noescape
func TryNavigatorJoinAdInterestGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	group unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_GetUserMedia
//go:noescape
func HasNavigatorGetUserMedia(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_GetUserMedia
//go:noescape
func NavigatorGetUserMediaFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_GetUserMedia
//go:noescape
func CallNavigatorGetUserMedia(
	this js.Ref, retPtr unsafe.Pointer,
	constraints unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref)

//go:wasmimport plat/js/web try_Navigator_GetUserMedia
//go:noescape
func TryNavigatorGetUserMedia(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	constraints unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_LeaveAdInterestGroup
//go:noescape
func HasNavigatorLeaveAdInterestGroup(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_LeaveAdInterestGroup
//go:noescape
func NavigatorLeaveAdInterestGroupFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_LeaveAdInterestGroup
//go:noescape
func CallNavigatorLeaveAdInterestGroup(
	this js.Ref, retPtr unsafe.Pointer,
	group unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_LeaveAdInterestGroup
//go:noescape
func TryNavigatorLeaveAdInterestGroup(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	group unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_LeaveAdInterestGroup1
//go:noescape
func HasNavigatorLeaveAdInterestGroup1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_LeaveAdInterestGroup1
//go:noescape
func NavigatorLeaveAdInterestGroup1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_LeaveAdInterestGroup1
//go:noescape
func CallNavigatorLeaveAdInterestGroup1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_LeaveAdInterestGroup1
//go:noescape
func TryNavigatorLeaveAdInterestGroup1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_RunAdAuction
//go:noescape
func HasNavigatorRunAdAuction(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_RunAdAuction
//go:noescape
func NavigatorRunAdAuctionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_RunAdAuction
//go:noescape
func CallNavigatorRunAdAuction(
	this js.Ref, retPtr unsafe.Pointer,
	config unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_RunAdAuction
//go:noescape
func TryNavigatorRunAdAuction(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	config unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_Share
//go:noescape
func HasNavigatorShare(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_Share
//go:noescape
func NavigatorShareFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_Share
//go:noescape
func CallNavigatorShare(
	this js.Ref, retPtr unsafe.Pointer,
	data unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_Share
//go:noescape
func TryNavigatorShare(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_Share1
//go:noescape
func HasNavigatorShare1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_Share1
//go:noescape
func NavigatorShare1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_Share1
//go:noescape
func CallNavigatorShare1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_Share1
//go:noescape
func TryNavigatorShare1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_CanShare
//go:noescape
func HasNavigatorCanShare(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_CanShare
//go:noescape
func NavigatorCanShareFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_CanShare
//go:noescape
func CallNavigatorCanShare(
	this js.Ref, retPtr unsafe.Pointer,
	data unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_CanShare
//go:noescape
func TryNavigatorCanShare(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	data unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_CanShare1
//go:noescape
func HasNavigatorCanShare1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_CanShare1
//go:noescape
func NavigatorCanShare1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_CanShare1
//go:noescape
func CallNavigatorCanShare1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_CanShare1
//go:noescape
func TryNavigatorCanShare1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_RequestMIDIAccess
//go:noescape
func HasNavigatorRequestMIDIAccess(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_RequestMIDIAccess
//go:noescape
func NavigatorRequestMIDIAccessFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_RequestMIDIAccess
//go:noescape
func CallNavigatorRequestMIDIAccess(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_RequestMIDIAccess
//go:noescape
func TryNavigatorRequestMIDIAccess(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_RequestMIDIAccess1
//go:noescape
func HasNavigatorRequestMIDIAccess1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_RequestMIDIAccess1
//go:noescape
func NavigatorRequestMIDIAccess1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_RequestMIDIAccess1
//go:noescape
func CallNavigatorRequestMIDIAccess1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_RequestMIDIAccess1
//go:noescape
func TryNavigatorRequestMIDIAccess1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_GetAutoplayPolicy
//go:noescape
func HasNavigatorGetAutoplayPolicy(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_GetAutoplayPolicy
//go:noescape
func NavigatorGetAutoplayPolicyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_GetAutoplayPolicy
//go:noescape
func CallNavigatorGetAutoplayPolicy(
	this js.Ref, retPtr unsafe.Pointer,
	typ uint32)

//go:wasmimport plat/js/web try_Navigator_GetAutoplayPolicy
//go:noescape
func TryNavigatorGetAutoplayPolicy(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_GetAutoplayPolicy1
//go:noescape
func HasNavigatorGetAutoplayPolicy1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_GetAutoplayPolicy1
//go:noescape
func NavigatorGetAutoplayPolicy1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_GetAutoplayPolicy1
//go:noescape
func CallNavigatorGetAutoplayPolicy1(
	this js.Ref, retPtr unsafe.Pointer,
	element js.Ref)

//go:wasmimport plat/js/web try_Navigator_GetAutoplayPolicy1
//go:noescape
func TryNavigatorGetAutoplayPolicy1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_GetAutoplayPolicy2
//go:noescape
func HasNavigatorGetAutoplayPolicy2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_GetAutoplayPolicy2
//go:noescape
func NavigatorGetAutoplayPolicy2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_GetAutoplayPolicy2
//go:noescape
func CallNavigatorGetAutoplayPolicy2(
	this js.Ref, retPtr unsafe.Pointer,
	context js.Ref)

//go:wasmimport plat/js/web try_Navigator_GetAutoplayPolicy2
//go:noescape
func TryNavigatorGetAutoplayPolicy2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	context js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_TaintEnabled
//go:noescape
func HasNavigatorTaintEnabled(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_TaintEnabled
//go:noescape
func NavigatorTaintEnabledFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_TaintEnabled
//go:noescape
func CallNavigatorTaintEnabled(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_TaintEnabled
//go:noescape
func TryNavigatorTaintEnabled(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_RegisterProtocolHandler
//go:noescape
func HasNavigatorRegisterProtocolHandler(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_RegisterProtocolHandler
//go:noescape
func NavigatorRegisterProtocolHandlerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_RegisterProtocolHandler
//go:noescape
func CallNavigatorRegisterProtocolHandler(
	this js.Ref, retPtr unsafe.Pointer,
	scheme js.Ref,
	url js.Ref)

//go:wasmimport plat/js/web try_Navigator_RegisterProtocolHandler
//go:noescape
func TryNavigatorRegisterProtocolHandler(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scheme js.Ref,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_UnregisterProtocolHandler
//go:noescape
func HasNavigatorUnregisterProtocolHandler(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_UnregisterProtocolHandler
//go:noescape
func NavigatorUnregisterProtocolHandlerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_UnregisterProtocolHandler
//go:noescape
func CallNavigatorUnregisterProtocolHandler(
	this js.Ref, retPtr unsafe.Pointer,
	scheme js.Ref,
	url js.Ref)

//go:wasmimport plat/js/web try_Navigator_UnregisterProtocolHandler
//go:noescape
func TryNavigatorUnregisterProtocolHandler(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scheme js.Ref,
	url js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_JavaEnabled
//go:noescape
func HasNavigatorJavaEnabled(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_JavaEnabled
//go:noescape
func NavigatorJavaEnabledFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_JavaEnabled
//go:noescape
func CallNavigatorJavaEnabled(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_JavaEnabled
//go:noescape
func TryNavigatorJavaEnabled(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_SetAppBadge
//go:noescape
func HasNavigatorSetAppBadge(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_SetAppBadge
//go:noescape
func NavigatorSetAppBadgeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_SetAppBadge
//go:noescape
func CallNavigatorSetAppBadge(
	this js.Ref, retPtr unsafe.Pointer,
	contents float64)

//go:wasmimport plat/js/web try_Navigator_SetAppBadge
//go:noescape
func TryNavigatorSetAppBadge(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	contents float64) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_SetAppBadge1
//go:noescape
func HasNavigatorSetAppBadge1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_SetAppBadge1
//go:noescape
func NavigatorSetAppBadge1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_SetAppBadge1
//go:noescape
func CallNavigatorSetAppBadge1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_SetAppBadge1
//go:noescape
func TryNavigatorSetAppBadge1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Navigator_ClearAppBadge
//go:noescape
func HasNavigatorClearAppBadge(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Navigator_ClearAppBadge
//go:noescape
func NavigatorClearAppBadgeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Navigator_ClearAppBadge
//go:noescape
func CallNavigatorClearAppBadge(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Navigator_ClearAppBadge
//go:noescape
func TryNavigatorClearAppBadge(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LaunchParams_TargetURL
//go:noescape
func GetLaunchParamsTargetURL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LaunchParams_Files
//go:noescape
func GetLaunchParamsFiles(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_LaunchQueue_SetConsumer
//go:noescape
func HasLaunchQueueSetConsumer(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_LaunchQueue_SetConsumer
//go:noescape
func LaunchQueueSetConsumerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_LaunchQueue_SetConsumer
//go:noescape
func CallLaunchQueueSetConsumer(
	this js.Ref, retPtr unsafe.Pointer,
	consumer js.Ref)

//go:wasmimport plat/js/web try_LaunchQueue_SetConsumer
//go:noescape
func TryLaunchQueueSetConsumer(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	consumer js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_FenceReportingDestination
//go:noescape
func ConstOfFenceReportingDestination(str js.Ref) uint32

//go:wasmimport plat/js/web store_FenceEvent
//go:noescape
func FenceEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FenceEvent
//go:noescape
func FenceEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_Fence_ReportEvent
//go:noescape
func HasFenceReportEvent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Fence_ReportEvent
//go:noescape
func FenceReportEventFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Fence_ReportEvent
//go:noescape
func CallFenceReportEvent(
	this js.Ref, retPtr unsafe.Pointer,
	event js.Ref)

//go:wasmimport plat/js/web try_Fence_ReportEvent
//go:noescape
func TryFenceReportEvent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	event js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Fence_SetReportEventDataForAutomaticBeacons
//go:noescape
func HasFenceSetReportEventDataForAutomaticBeacons(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Fence_SetReportEventDataForAutomaticBeacons
//go:noescape
func FenceSetReportEventDataForAutomaticBeaconsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Fence_SetReportEventDataForAutomaticBeacons
//go:noescape
func CallFenceSetReportEventDataForAutomaticBeacons(
	this js.Ref, retPtr unsafe.Pointer,
	event unsafe.Pointer)

//go:wasmimport plat/js/web try_Fence_SetReportEventDataForAutomaticBeacons
//go:noescape
func TryFenceSetReportEventDataForAutomaticBeacons(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	event unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Fence_GetNestedConfigs
//go:noescape
func HasFenceGetNestedConfigs(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Fence_GetNestedConfigs
//go:noescape
func FenceGetNestedConfigsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Fence_GetNestedConfigs
//go:noescape
func CallFenceGetNestedConfigs(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Fence_GetNestedConfigs
//go:noescape
func TryFenceGetNestedConfigs(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PortalHost_PostMessage
//go:noescape
func HasPortalHostPostMessage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PortalHost_PostMessage
//go:noescape
func PortalHostPostMessageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PortalHost_PostMessage
//go:noescape
func CallPortalHostPostMessage(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_PortalHost_PostMessage
//go:noescape
func TryPortalHostPostMessage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PortalHost_PostMessage1
//go:noescape
func HasPortalHostPostMessage1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PortalHost_PostMessage1
//go:noescape
func PortalHostPostMessage1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PortalHost_PostMessage1
//go:noescape
func CallPortalHostPostMessage1(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_PortalHost_PostMessage1
//go:noescape
func TryPortalHostPostMessage1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_OrientationLockType
//go:noescape
func ConstOfOrientationLockType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_OrientationType
//go:noescape
func ConstOfOrientationType(str js.Ref) uint32

//go:wasmimport plat/js/web get_ScreenOrientation_Type
//go:noescape
func GetScreenOrientationType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ScreenOrientation_Angle
//go:noescape
func GetScreenOrientationAngle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ScreenOrientation_Lock
//go:noescape
func HasScreenOrientationLock(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ScreenOrientation_Lock
//go:noescape
func ScreenOrientationLockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ScreenOrientation_Lock
//go:noescape
func CallScreenOrientationLock(
	this js.Ref, retPtr unsafe.Pointer,
	orientation uint32)

//go:wasmimport plat/js/web try_ScreenOrientation_Lock
//go:noescape
func TryScreenOrientationLock(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	orientation uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_ScreenOrientation_Unlock
//go:noescape
func HasScreenOrientationUnlock(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ScreenOrientation_Unlock
//go:noescape
func ScreenOrientationUnlockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ScreenOrientation_Unlock
//go:noescape
func CallScreenOrientationUnlock(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ScreenOrientation_Unlock
//go:noescape
func TryScreenOrientationUnlock(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Screen_AvailWidth
//go:noescape
func GetScreenAvailWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Screen_AvailHeight
//go:noescape
func GetScreenAvailHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Screen_Width
//go:noescape
func GetScreenWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Screen_Height
//go:noescape
func GetScreenHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Screen_ColorDepth
//go:noescape
func GetScreenColorDepth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Screen_PixelDepth
//go:noescape
func GetScreenPixelDepth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Screen_IsExtended
//go:noescape
func GetScreenIsExtended(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Screen_Orientation
//go:noescape
func GetScreenOrientation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VisualViewport_OffsetLeft
//go:noescape
func GetVisualViewportOffsetLeft(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VisualViewport_OffsetTop
//go:noescape
func GetVisualViewportOffsetTop(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VisualViewport_PageLeft
//go:noescape
func GetVisualViewportPageLeft(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VisualViewport_PageTop
//go:noescape
func GetVisualViewportPageTop(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VisualViewport_Width
//go:noescape
func GetVisualViewportWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VisualViewport_Height
//go:noescape
func GetVisualViewportHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_VisualViewport_Scale
//go:noescape
func GetVisualViewportScale(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_SharedStorageRunOperationMethodOptions
//go:noescape
func SharedStorageRunOperationMethodOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SharedStorageRunOperationMethodOptions
//go:noescape
func SharedStorageRunOperationMethodOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_SharedStorageUrlWithMetadata
//go:noescape
func SharedStorageUrlWithMetadataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SharedStorageUrlWithMetadata
//go:noescape
func SharedStorageUrlWithMetadataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

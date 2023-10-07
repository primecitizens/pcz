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

//go:wasmimport plat/js/web store_BluetoothServiceDataFilterInit
//go:noescape
func BluetoothServiceDataFilterInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothServiceDataFilterInit
//go:noescape
func BluetoothServiceDataFilterInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_BluetoothLEScanFilterInit
//go:noescape
func BluetoothLEScanFilterInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothLEScanFilterInit
//go:noescape
func BluetoothLEScanFilterInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_RequestDeviceOptions
//go:noescape
func RequestDeviceOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_RequestDeviceOptions
//go:noescape
func RequestDeviceOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_Bluetooth_ReferringDevice
//go:noescape
func GetBluetoothReferringDevice(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Bluetooth_GetAvailability
//go:noescape
func HasFuncBluetoothGetAvailability(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_GetAvailability
//go:noescape
func FuncBluetoothGetAvailability(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Bluetooth_GetAvailability
//go:noescape
func CallBluetoothGetAvailability(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Bluetooth_GetAvailability
//go:noescape
func TryBluetoothGetAvailability(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Bluetooth_GetDevices
//go:noescape
func HasFuncBluetoothGetDevices(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_GetDevices
//go:noescape
func FuncBluetoothGetDevices(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Bluetooth_GetDevices
//go:noescape
func CallBluetoothGetDevices(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Bluetooth_GetDevices
//go:noescape
func TryBluetoothGetDevices(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Bluetooth_RequestDevice
//go:noescape
func HasFuncBluetoothRequestDevice(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_RequestDevice
//go:noescape
func FuncBluetoothRequestDevice(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Bluetooth_RequestDevice
//go:noescape
func CallBluetoothRequestDevice(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Bluetooth_RequestDevice
//go:noescape
func TryBluetoothRequestDevice(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Bluetooth_RequestDevice1
//go:noescape
func HasFuncBluetoothRequestDevice1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_RequestDevice1
//go:noescape
func FuncBluetoothRequestDevice1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Bluetooth_RequestDevice1
//go:noescape
func CallBluetoothRequestDevice1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Bluetooth_RequestDevice1
//go:noescape
func TryBluetoothRequestDevice1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_BluetoothAdvertisingEventInit
//go:noescape
func BluetoothAdvertisingEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothAdvertisingEventInit
//go:noescape
func BluetoothAdvertisingEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_BluetoothAdvertisingEvent_BluetoothAdvertisingEvent
//go:noescape
func NewBluetoothAdvertisingEventByBluetoothAdvertisingEvent(
	typ js.Ref,
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_Device
//go:noescape
func GetBluetoothAdvertisingEventDevice(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_Uuids
//go:noescape
func GetBluetoothAdvertisingEventUuids(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_Name
//go:noescape
func GetBluetoothAdvertisingEventName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_Appearance
//go:noescape
func GetBluetoothAdvertisingEventAppearance(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_TxPower
//go:noescape
func GetBluetoothAdvertisingEventTxPower(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_Rssi
//go:noescape
func GetBluetoothAdvertisingEventRssi(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_ManufacturerData
//go:noescape
func GetBluetoothAdvertisingEventManufacturerData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BluetoothAdvertisingEvent_ServiceData
//go:noescape
func GetBluetoothAdvertisingEventServiceData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_BluetoothDataFilterInit
//go:noescape
func BluetoothDataFilterInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothDataFilterInit
//go:noescape
func BluetoothDataFilterInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_BluetoothPermissionDescriptor
//go:noescape
func BluetoothPermissionDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothPermissionDescriptor
//go:noescape
func BluetoothPermissionDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_BluetoothPermissionResult_Devices
//go:noescape
func GetBluetoothPermissionResultDevices(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_BluetoothPermissionResult_Devices
//go:noescape
func SetBluetoothPermissionResultDevices(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_BluetoothPermissionStorage
//go:noescape
func BluetoothPermissionStorageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BluetoothPermissionStorage
//go:noescape
func BluetoothPermissionStorageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_BluetoothUUID_GetService
//go:noescape
func HasFuncBluetoothUUIDGetService(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_GetService
//go:noescape
func FuncBluetoothUUIDGetService(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BluetoothUUID_GetService
//go:noescape
func CallBluetoothUUIDGetService(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_BluetoothUUID_GetService
//go:noescape
func TryBluetoothUUIDGetService(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothUUID_GetCharacteristic
//go:noescape
func HasFuncBluetoothUUIDGetCharacteristic(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_GetCharacteristic
//go:noescape
func FuncBluetoothUUIDGetCharacteristic(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BluetoothUUID_GetCharacteristic
//go:noescape
func CallBluetoothUUIDGetCharacteristic(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_BluetoothUUID_GetCharacteristic
//go:noescape
func TryBluetoothUUIDGetCharacteristic(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothUUID_GetDescriptor
//go:noescape
func HasFuncBluetoothUUIDGetDescriptor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_GetDescriptor
//go:noescape
func FuncBluetoothUUIDGetDescriptor(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BluetoothUUID_GetDescriptor
//go:noescape
func CallBluetoothUUIDGetDescriptor(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_BluetoothUUID_GetDescriptor
//go:noescape
func TryBluetoothUUIDGetDescriptor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BluetoothUUID_CanonicalUUID
//go:noescape
func HasFuncBluetoothUUIDCanonicalUUID(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_CanonicalUUID
//go:noescape
func FuncBluetoothUUIDCanonicalUUID(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BluetoothUUID_CanonicalUUID
//go:noescape
func CallBluetoothUUIDCanonicalUUID(
	this js.Ref, retPtr unsafe.Pointer,
	alias uint32)

//go:wasmimport plat/js/web try_BluetoothUUID_CanonicalUUID
//go:noescape
func TryBluetoothUUIDCanonicalUUID(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	alias uint32) (ok js.Ref)

//go:wasmimport plat/js/web constof_BreakType
//go:noescape
func ConstOfBreakType(str js.Ref) uint32

//go:wasmimport plat/js/web get_IntrinsicSizes_MinContentSize
//go:noescape
func GetIntrinsicSizesMinContentSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IntrinsicSizes_MaxContentSize
//go:noescape
func GetIntrinsicSizesMaxContentSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutFragment_InlineSize
//go:noescape
func GetLayoutFragmentInlineSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutFragment_BlockSize
//go:noescape
func GetLayoutFragmentBlockSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutFragment_InlineOffset
//go:noescape
func GetLayoutFragmentInlineOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_LayoutFragment_InlineOffset
//go:noescape
func SetLayoutFragmentInlineOffset(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_LayoutFragment_BlockOffset
//go:noescape
func GetLayoutFragmentBlockOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_LayoutFragment_BlockOffset
//go:noescape
func SetLayoutFragmentBlockOffset(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_LayoutFragment_Data
//go:noescape
func GetLayoutFragmentData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_LayoutFragment_BreakToken
//go:noescape
func GetLayoutFragmentBreakToken(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_LayoutConstraintsOptions
//go:noescape
func LayoutConstraintsOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_LayoutConstraintsOptions
//go:noescape
func LayoutConstraintsOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_LayoutChild_StyleMap
//go:noescape
func GetLayoutChildStyleMap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_LayoutChild_IntrinsicSizes
//go:noescape
func HasFuncLayoutChildIntrinsicSizes(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_LayoutChild_IntrinsicSizes
//go:noescape
func FuncLayoutChildIntrinsicSizes(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_LayoutChild_IntrinsicSizes
//go:noescape
func CallLayoutChildIntrinsicSizes(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_LayoutChild_IntrinsicSizes
//go:noescape
func TryLayoutChildIntrinsicSizes(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_LayoutChild_LayoutNextFragment
//go:noescape
func HasFuncLayoutChildLayoutNextFragment(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_LayoutChild_LayoutNextFragment
//go:noescape
func FuncLayoutChildLayoutNextFragment(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_LayoutChild_LayoutNextFragment
//go:noescape
func CallLayoutChildLayoutNextFragment(
	this js.Ref, retPtr unsafe.Pointer,
	constraints unsafe.Pointer,
	breakToken js.Ref)

//go:wasmimport plat/js/web try_LayoutChild_LayoutNextFragment
//go:noescape
func TryLayoutChildLayoutNextFragment(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	constraints unsafe.Pointer,
	breakToken js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_ChildBreakToken_BreakType
//go:noescape
func GetChildBreakTokenBreakType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ChildBreakToken_Child
//go:noescape
func GetChildBreakTokenChild(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BreakToken_ChildBreakTokens
//go:noescape
func GetBreakTokenChildBreakTokens(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_BreakToken_Data
//go:noescape
func GetBreakTokenData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_BreakTokenOptions
//go:noescape
func BreakTokenOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BreakTokenOptions
//go:noescape
func BreakTokenOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_BroadcastChannel_BroadcastChannel
//go:noescape
func NewBroadcastChannelByBroadcastChannel(
	name js.Ref) js.Ref

//go:wasmimport plat/js/web get_BroadcastChannel_Name
//go:noescape
func GetBroadcastChannelName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BroadcastChannel_PostMessage
//go:noescape
func HasFuncBroadcastChannelPostMessage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BroadcastChannel_PostMessage
//go:noescape
func FuncBroadcastChannelPostMessage(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BroadcastChannel_PostMessage
//go:noescape
func CallBroadcastChannelPostMessage(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_BroadcastChannel_PostMessage
//go:noescape
func TryBroadcastChannelPostMessage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BroadcastChannel_Close
//go:noescape
func HasFuncBroadcastChannelClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BroadcastChannel_Close
//go:noescape
func FuncBroadcastChannelClose(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BroadcastChannel_Close
//go:noescape
func CallBroadcastChannelClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BroadcastChannel_Close
//go:noescape
func TryBroadcastChannelClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CropTarget_FromElement
//go:noescape
func HasFuncCropTargetFromElement(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CropTarget_FromElement
//go:noescape
func FuncCropTargetFromElement(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CropTarget_FromElement
//go:noescape
func CallCropTargetFromElement(
	this js.Ref, retPtr unsafe.Pointer,
	element js.Ref)

//go:wasmimport plat/js/web try_CropTarget_FromElement
//go:noescape
func TryCropTargetFromElement(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_RestrictionTarget_FromElement
//go:noescape
func HasFuncRestrictionTargetFromElement(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RestrictionTarget_FromElement
//go:noescape
func FuncRestrictionTargetFromElement(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_RestrictionTarget_FromElement
//go:noescape
func CallRestrictionTargetFromElement(
	this js.Ref, retPtr unsafe.Pointer,
	element js.Ref)

//go:wasmimport plat/js/web try_RestrictionTarget_FromElement
//go:noescape
func TryRestrictionTargetFromElement(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BrowserCaptureMediaStreamTrack_CropTo
//go:noescape
func HasFuncBrowserCaptureMediaStreamTrackCropTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BrowserCaptureMediaStreamTrack_CropTo
//go:noescape
func FuncBrowserCaptureMediaStreamTrackCropTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BrowserCaptureMediaStreamTrack_CropTo
//go:noescape
func CallBrowserCaptureMediaStreamTrackCropTo(
	this js.Ref, retPtr unsafe.Pointer,
	cropTarget js.Ref)

//go:wasmimport plat/js/web try_BrowserCaptureMediaStreamTrack_CropTo
//go:noescape
func TryBrowserCaptureMediaStreamTrackCropTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	cropTarget js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_BrowserCaptureMediaStreamTrack_Clone
//go:noescape
func HasFuncBrowserCaptureMediaStreamTrackClone(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BrowserCaptureMediaStreamTrack_Clone
//go:noescape
func FuncBrowserCaptureMediaStreamTrackClone(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BrowserCaptureMediaStreamTrack_Clone
//go:noescape
func CallBrowserCaptureMediaStreamTrackClone(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_BrowserCaptureMediaStreamTrack_Clone
//go:noescape
func TryBrowserCaptureMediaStreamTrackClone(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_BrowserCaptureMediaStreamTrack_RestrictTo
//go:noescape
func HasFuncBrowserCaptureMediaStreamTrackRestrictTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BrowserCaptureMediaStreamTrack_RestrictTo
//go:noescape
func FuncBrowserCaptureMediaStreamTrackRestrictTo(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_BrowserCaptureMediaStreamTrack_RestrictTo
//go:noescape
func CallBrowserCaptureMediaStreamTrackRestrictTo(
	this js.Ref, retPtr unsafe.Pointer,
	RestrictionTarget js.Ref)

//go:wasmimport plat/js/web try_BrowserCaptureMediaStreamTrack_RestrictTo
//go:noescape
func TryBrowserCaptureMediaStreamTrackRestrictTo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	RestrictionTarget js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_QueuingStrategyInit
//go:noescape
func QueuingStrategyInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_QueuingStrategyInit
//go:noescape
func QueuingStrategyInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_ByteLengthQueuingStrategy_ByteLengthQueuingStrategy
//go:noescape
func NewByteLengthQueuingStrategyByByteLengthQueuingStrategy(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_ByteLengthQueuingStrategy_HighWaterMark
//go:noescape
func GetByteLengthQueuingStrategyHighWaterMark(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ByteLengthQueuingStrategy_Size
//go:noescape
func GetByteLengthQueuingStrategySize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_SecurityPolicyViolationEventDisposition
//go:noescape
func ConstOfSecurityPolicyViolationEventDisposition(str js.Ref) uint32

//go:wasmimport plat/js/web get_CSPViolationReportBody_DocumentURL
//go:noescape
func GetCSPViolationReportBodyDocumentURL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSPViolationReportBody_Referrer
//go:noescape
func GetCSPViolationReportBodyReferrer(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSPViolationReportBody_BlockedURL
//go:noescape
func GetCSPViolationReportBodyBlockedURL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSPViolationReportBody_EffectiveDirective
//go:noescape
func GetCSPViolationReportBodyEffectiveDirective(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSPViolationReportBody_OriginalPolicy
//go:noescape
func GetCSPViolationReportBodyOriginalPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSPViolationReportBody_SourceFile
//go:noescape
func GetCSPViolationReportBodySourceFile(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSPViolationReportBody_Sample
//go:noescape
func GetCSPViolationReportBodySample(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSPViolationReportBody_Disposition
//go:noescape
func GetCSPViolationReportBodyDisposition(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSPViolationReportBody_StatusCode
//go:noescape
func GetCSPViolationReportBodyStatusCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSPViolationReportBody_LineNumber
//go:noescape
func GetCSPViolationReportBodyLineNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSPViolationReportBody_ColumnNumber
//go:noescape
func GetCSPViolationReportBodyColumnNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSPViolationReportBody_ToJSON
//go:noescape
func HasFuncCSPViolationReportBodyToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSPViolationReportBody_ToJSON
//go:noescape
func FuncCSPViolationReportBodyToJSON(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSPViolationReportBody_ToJSON
//go:noescape
func CallCSPViolationReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CSPViolationReportBody_ToJSON
//go:noescape
func TryCSPViolationReportBodyToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PropertyDefinition
//go:noescape
func PropertyDefinitionJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PropertyDefinition
//go:noescape
func PropertyDefinitionJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CSSParserOptions
//go:noescape
func CSSParserOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CSSParserOptions
//go:noescape
func CSSParserOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSParserDeclaration_CSSParserDeclaration
//go:noescape
func NewCSSParserDeclarationByCSSParserDeclaration(
	name js.Ref,
	body js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSParserDeclaration_CSSParserDeclaration1
//go:noescape
func NewCSSParserDeclarationByCSSParserDeclaration1(
	name js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSParserDeclaration_Name
//go:noescape
func GetCSSParserDeclarationName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSParserDeclaration_Body
//go:noescape
func GetCSSParserDeclarationBody(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSParserDeclaration_ToString
//go:noescape
func HasFuncCSSParserDeclarationToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSParserDeclaration_ToString
//go:noescape
func FuncCSSParserDeclarationToString(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSSParserDeclaration_ToString
//go:noescape
func CallCSSParserDeclarationToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CSSParserDeclaration_ToString
//go:noescape
func TryCSSParserDeclarationToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_WorkletOptions
//go:noescape
func WorkletOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_WorkletOptions
//go:noescape
func WorkletOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_Worklet_AddModule
//go:noescape
func HasFuncWorkletAddModule(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Worklet_AddModule
//go:noescape
func FuncWorkletAddModule(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Worklet_AddModule
//go:noescape
func CallWorkletAddModule(
	this js.Ref, retPtr unsafe.Pointer,
	moduleURL js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Worklet_AddModule
//go:noescape
func TryWorkletAddModule(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	moduleURL js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Worklet_AddModule1
//go:noescape
func HasFuncWorkletAddModule1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Worklet_AddModule1
//go:noescape
func FuncWorkletAddModule1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_Worklet_AddModule1
//go:noescape
func CallWorkletAddModule1(
	this js.Ref, retPtr unsafe.Pointer,
	moduleURL js.Ref)

//go:wasmimport plat/js/web try_Worklet_AddModule1
//go:noescape
func TryWorkletAddModule1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	moduleURL js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_CSS_ElementSources
//go:noescape
func GetCSSElementSources(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_CSS_AnimationWorklet
//go:noescape
func GetCSSAnimationWorklet(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_CSS_PaintWorklet
//go:noescape
func GetCSSPaintWorklet(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_CSS_LayoutWorklet
//go:noescape
func GetCSSLayoutWorklet(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_CSS_Highlights
//go:noescape
func GetCSSHighlights(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_CSS_Escape
//go:noescape
func HasFuncCSSEscape() js.Ref

//go:wasmimport plat/js/web func_CSS_Escape
//go:noescape
func FuncCSSEscape(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Escape
//go:noescape
func CallCSSEscape(
	retPtr unsafe.Pointer,
	ident js.Ref)

//go:wasmimport plat/js/web try_CSS_Escape
//go:noescape
func TryCSSEscape(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	ident js.Ref) js.Ref

//go:wasmimport plat/js/web has_CSS_RegisterProperty
//go:noescape
func HasFuncCSSRegisterProperty() js.Ref

//go:wasmimport plat/js/web func_CSS_RegisterProperty
//go:noescape
func FuncCSSRegisterProperty(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_RegisterProperty
//go:noescape
func CallCSSRegisterProperty(
	retPtr unsafe.Pointer,
	definition unsafe.Pointer)

//go:wasmimport plat/js/web try_CSS_RegisterProperty
//go:noescape
func TryCSSRegisterProperty(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	definition unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_CSS_Supports
//go:noescape
func HasFuncCSSSupports() js.Ref

//go:wasmimport plat/js/web func_CSS_Supports
//go:noescape
func FuncCSSSupports(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Supports
//go:noescape
func CallCSSSupports(
	retPtr unsafe.Pointer,
	property js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_CSS_Supports
//go:noescape
func TryCSSSupports(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	property js.Ref,
	value js.Ref) js.Ref

//go:wasmimport plat/js/web has_CSS_Supports1
//go:noescape
func HasFuncCSSSupports1() js.Ref

//go:wasmimport plat/js/web func_CSS_Supports1
//go:noescape
func FuncCSSSupports1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Supports1
//go:noescape
func CallCSSSupports1(
	retPtr unsafe.Pointer,
	conditionText js.Ref)

//go:wasmimport plat/js/web try_CSS_Supports1
//go:noescape
func TryCSSSupports1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	conditionText js.Ref) js.Ref

//go:wasmimport plat/js/web has_CSS_Number
//go:noescape
func HasFuncCSSNumber() js.Ref

//go:wasmimport plat/js/web func_CSS_Number
//go:noescape
func FuncCSSNumber(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Number
//go:noescape
func CallCSSNumber(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Number
//go:noescape
func TryCSSNumber(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Percent
//go:noescape
func HasFuncCSSPercent() js.Ref

//go:wasmimport plat/js/web func_CSS_Percent
//go:noescape
func FuncCSSPercent(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Percent
//go:noescape
func CallCSSPercent(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Percent
//go:noescape
func TryCSSPercent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Cap
//go:noescape
func HasFuncCSSCap() js.Ref

//go:wasmimport plat/js/web func_CSS_Cap
//go:noescape
func FuncCSSCap(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Cap
//go:noescape
func CallCSSCap(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cap
//go:noescape
func TryCSSCap(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Ch
//go:noescape
func HasFuncCSSCh() js.Ref

//go:wasmimport plat/js/web func_CSS_Ch
//go:noescape
func FuncCSSCh(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Ch
//go:noescape
func CallCSSCh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Ch
//go:noescape
func TryCSSCh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Em
//go:noescape
func HasFuncCSSEm() js.Ref

//go:wasmimport plat/js/web func_CSS_Em
//go:noescape
func FuncCSSEm(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Em
//go:noescape
func CallCSSEm(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Em
//go:noescape
func TryCSSEm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Ex
//go:noescape
func HasFuncCSSEx() js.Ref

//go:wasmimport plat/js/web func_CSS_Ex
//go:noescape
func FuncCSSEx(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Ex
//go:noescape
func CallCSSEx(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Ex
//go:noescape
func TryCSSEx(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Ic
//go:noescape
func HasFuncCSSIc() js.Ref

//go:wasmimport plat/js/web func_CSS_Ic
//go:noescape
func FuncCSSIc(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Ic
//go:noescape
func CallCSSIc(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Ic
//go:noescape
func TryCSSIc(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Lh
//go:noescape
func HasFuncCSSLh() js.Ref

//go:wasmimport plat/js/web func_CSS_Lh
//go:noescape
func FuncCSSLh(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Lh
//go:noescape
func CallCSSLh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lh
//go:noescape
func TryCSSLh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Rcap
//go:noescape
func HasFuncCSSRcap() js.Ref

//go:wasmimport plat/js/web func_CSS_Rcap
//go:noescape
func FuncCSSRcap(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Rcap
//go:noescape
func CallCSSRcap(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rcap
//go:noescape
func TryCSSRcap(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Rch
//go:noescape
func HasFuncCSSRch() js.Ref

//go:wasmimport plat/js/web func_CSS_Rch
//go:noescape
func FuncCSSRch(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Rch
//go:noescape
func CallCSSRch(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rch
//go:noescape
func TryCSSRch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Rem
//go:noescape
func HasFuncCSSRem() js.Ref

//go:wasmimport plat/js/web func_CSS_Rem
//go:noescape
func FuncCSSRem(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Rem
//go:noescape
func CallCSSRem(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rem
//go:noescape
func TryCSSRem(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Rex
//go:noescape
func HasFuncCSSRex() js.Ref

//go:wasmimport plat/js/web func_CSS_Rex
//go:noescape
func FuncCSSRex(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Rex
//go:noescape
func CallCSSRex(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rex
//go:noescape
func TryCSSRex(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Ric
//go:noescape
func HasFuncCSSRic() js.Ref

//go:wasmimport plat/js/web func_CSS_Ric
//go:noescape
func FuncCSSRic(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Ric
//go:noescape
func CallCSSRic(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Ric
//go:noescape
func TryCSSRic(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Rlh
//go:noescape
func HasFuncCSSRlh() js.Ref

//go:wasmimport plat/js/web func_CSS_Rlh
//go:noescape
func FuncCSSRlh(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Rlh
//go:noescape
func CallCSSRlh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rlh
//go:noescape
func TryCSSRlh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Vw
//go:noescape
func HasFuncCSSVw() js.Ref

//go:wasmimport plat/js/web func_CSS_Vw
//go:noescape
func FuncCSSVw(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Vw
//go:noescape
func CallCSSVw(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vw
//go:noescape
func TryCSSVw(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Vh
//go:noescape
func HasFuncCSSVh() js.Ref

//go:wasmimport plat/js/web func_CSS_Vh
//go:noescape
func FuncCSSVh(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Vh
//go:noescape
func CallCSSVh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vh
//go:noescape
func TryCSSVh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Vi
//go:noescape
func HasFuncCSSVi() js.Ref

//go:wasmimport plat/js/web func_CSS_Vi
//go:noescape
func FuncCSSVi(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Vi
//go:noescape
func CallCSSVi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vi
//go:noescape
func TryCSSVi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Vb
//go:noescape
func HasFuncCSSVb() js.Ref

//go:wasmimport plat/js/web func_CSS_Vb
//go:noescape
func FuncCSSVb(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Vb
//go:noescape
func CallCSSVb(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vb
//go:noescape
func TryCSSVb(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Vmin
//go:noescape
func HasFuncCSSVmin() js.Ref

//go:wasmimport plat/js/web func_CSS_Vmin
//go:noescape
func FuncCSSVmin(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Vmin
//go:noescape
func CallCSSVmin(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vmin
//go:noescape
func TryCSSVmin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Vmax
//go:noescape
func HasFuncCSSVmax() js.Ref

//go:wasmimport plat/js/web func_CSS_Vmax
//go:noescape
func FuncCSSVmax(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Vmax
//go:noescape
func CallCSSVmax(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vmax
//go:noescape
func TryCSSVmax(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Svw
//go:noescape
func HasFuncCSSSvw() js.Ref

//go:wasmimport plat/js/web func_CSS_Svw
//go:noescape
func FuncCSSSvw(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Svw
//go:noescape
func CallCSSSvw(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svw
//go:noescape
func TryCSSSvw(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Svh
//go:noescape
func HasFuncCSSSvh() js.Ref

//go:wasmimport plat/js/web func_CSS_Svh
//go:noescape
func FuncCSSSvh(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Svh
//go:noescape
func CallCSSSvh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svh
//go:noescape
func TryCSSSvh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Svi
//go:noescape
func HasFuncCSSSvi() js.Ref

//go:wasmimport plat/js/web func_CSS_Svi
//go:noescape
func FuncCSSSvi(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Svi
//go:noescape
func CallCSSSvi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svi
//go:noescape
func TryCSSSvi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Svb
//go:noescape
func HasFuncCSSSvb() js.Ref

//go:wasmimport plat/js/web func_CSS_Svb
//go:noescape
func FuncCSSSvb(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Svb
//go:noescape
func CallCSSSvb(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svb
//go:noescape
func TryCSSSvb(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Svmin
//go:noescape
func HasFuncCSSSvmin() js.Ref

//go:wasmimport plat/js/web func_CSS_Svmin
//go:noescape
func FuncCSSSvmin(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Svmin
//go:noescape
func CallCSSSvmin(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svmin
//go:noescape
func TryCSSSvmin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Svmax
//go:noescape
func HasFuncCSSSvmax() js.Ref

//go:wasmimport plat/js/web func_CSS_Svmax
//go:noescape
func FuncCSSSvmax(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Svmax
//go:noescape
func CallCSSSvmax(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svmax
//go:noescape
func TryCSSSvmax(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Lvw
//go:noescape
func HasFuncCSSLvw() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvw
//go:noescape
func FuncCSSLvw(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Lvw
//go:noescape
func CallCSSLvw(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvw
//go:noescape
func TryCSSLvw(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Lvh
//go:noescape
func HasFuncCSSLvh() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvh
//go:noescape
func FuncCSSLvh(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Lvh
//go:noescape
func CallCSSLvh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvh
//go:noescape
func TryCSSLvh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Lvi
//go:noescape
func HasFuncCSSLvi() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvi
//go:noescape
func FuncCSSLvi(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Lvi
//go:noescape
func CallCSSLvi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvi
//go:noescape
func TryCSSLvi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Lvb
//go:noescape
func HasFuncCSSLvb() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvb
//go:noescape
func FuncCSSLvb(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Lvb
//go:noescape
func CallCSSLvb(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvb
//go:noescape
func TryCSSLvb(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Lvmin
//go:noescape
func HasFuncCSSLvmin() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvmin
//go:noescape
func FuncCSSLvmin(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Lvmin
//go:noescape
func CallCSSLvmin(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvmin
//go:noescape
func TryCSSLvmin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Lvmax
//go:noescape
func HasFuncCSSLvmax() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvmax
//go:noescape
func FuncCSSLvmax(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Lvmax
//go:noescape
func CallCSSLvmax(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvmax
//go:noescape
func TryCSSLvmax(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Dvw
//go:noescape
func HasFuncCSSDvw() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvw
//go:noescape
func FuncCSSDvw(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Dvw
//go:noescape
func CallCSSDvw(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvw
//go:noescape
func TryCSSDvw(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Dvh
//go:noescape
func HasFuncCSSDvh() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvh
//go:noescape
func FuncCSSDvh(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Dvh
//go:noescape
func CallCSSDvh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvh
//go:noescape
func TryCSSDvh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Dvi
//go:noescape
func HasFuncCSSDvi() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvi
//go:noescape
func FuncCSSDvi(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Dvi
//go:noescape
func CallCSSDvi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvi
//go:noescape
func TryCSSDvi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Dvb
//go:noescape
func HasFuncCSSDvb() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvb
//go:noescape
func FuncCSSDvb(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Dvb
//go:noescape
func CallCSSDvb(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvb
//go:noescape
func TryCSSDvb(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Dvmin
//go:noescape
func HasFuncCSSDvmin() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvmin
//go:noescape
func FuncCSSDvmin(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Dvmin
//go:noescape
func CallCSSDvmin(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvmin
//go:noescape
func TryCSSDvmin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Dvmax
//go:noescape
func HasFuncCSSDvmax() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvmax
//go:noescape
func FuncCSSDvmax(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Dvmax
//go:noescape
func CallCSSDvmax(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvmax
//go:noescape
func TryCSSDvmax(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Cqw
//go:noescape
func HasFuncCSSCqw() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqw
//go:noescape
func FuncCSSCqw(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Cqw
//go:noescape
func CallCSSCqw(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqw
//go:noescape
func TryCSSCqw(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Cqh
//go:noescape
func HasFuncCSSCqh() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqh
//go:noescape
func FuncCSSCqh(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Cqh
//go:noescape
func CallCSSCqh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqh
//go:noescape
func TryCSSCqh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Cqi
//go:noescape
func HasFuncCSSCqi() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqi
//go:noescape
func FuncCSSCqi(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Cqi
//go:noescape
func CallCSSCqi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqi
//go:noescape
func TryCSSCqi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Cqb
//go:noescape
func HasFuncCSSCqb() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqb
//go:noescape
func FuncCSSCqb(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Cqb
//go:noescape
func CallCSSCqb(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqb
//go:noescape
func TryCSSCqb(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Cqmin
//go:noescape
func HasFuncCSSCqmin() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqmin
//go:noescape
func FuncCSSCqmin(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Cqmin
//go:noescape
func CallCSSCqmin(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqmin
//go:noescape
func TryCSSCqmin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Cqmax
//go:noescape
func HasFuncCSSCqmax() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqmax
//go:noescape
func FuncCSSCqmax(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Cqmax
//go:noescape
func CallCSSCqmax(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqmax
//go:noescape
func TryCSSCqmax(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Cm
//go:noescape
func HasFuncCSSCm() js.Ref

//go:wasmimport plat/js/web func_CSS_Cm
//go:noescape
func FuncCSSCm(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Cm
//go:noescape
func CallCSSCm(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cm
//go:noescape
func TryCSSCm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Mm
//go:noescape
func HasFuncCSSMm() js.Ref

//go:wasmimport plat/js/web func_CSS_Mm
//go:noescape
func FuncCSSMm(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Mm
//go:noescape
func CallCSSMm(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Mm
//go:noescape
func TryCSSMm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Q
//go:noescape
func HasFuncCSSQ() js.Ref

//go:wasmimport plat/js/web func_CSS_Q
//go:noescape
func FuncCSSQ(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Q
//go:noescape
func CallCSSQ(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Q
//go:noescape
func TryCSSQ(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_In
//go:noescape
func HasFuncCSSIn() js.Ref

//go:wasmimport plat/js/web func_CSS_In
//go:noescape
func FuncCSSIn(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_In
//go:noescape
func CallCSSIn(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_In
//go:noescape
func TryCSSIn(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Pt
//go:noescape
func HasFuncCSSPt() js.Ref

//go:wasmimport plat/js/web func_CSS_Pt
//go:noescape
func FuncCSSPt(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Pt
//go:noescape
func CallCSSPt(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Pt
//go:noescape
func TryCSSPt(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Pc
//go:noescape
func HasFuncCSSPc() js.Ref

//go:wasmimport plat/js/web func_CSS_Pc
//go:noescape
func FuncCSSPc(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Pc
//go:noescape
func CallCSSPc(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Pc
//go:noescape
func TryCSSPc(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Px
//go:noescape
func HasFuncCSSPx() js.Ref

//go:wasmimport plat/js/web func_CSS_Px
//go:noescape
func FuncCSSPx(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Px
//go:noescape
func CallCSSPx(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Px
//go:noescape
func TryCSSPx(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Deg
//go:noescape
func HasFuncCSSDeg() js.Ref

//go:wasmimport plat/js/web func_CSS_Deg
//go:noescape
func FuncCSSDeg(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Deg
//go:noescape
func CallCSSDeg(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Deg
//go:noescape
func TryCSSDeg(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Grad
//go:noescape
func HasFuncCSSGrad() js.Ref

//go:wasmimport plat/js/web func_CSS_Grad
//go:noescape
func FuncCSSGrad(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Grad
//go:noescape
func CallCSSGrad(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Grad
//go:noescape
func TryCSSGrad(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Rad
//go:noescape
func HasFuncCSSRad() js.Ref

//go:wasmimport plat/js/web func_CSS_Rad
//go:noescape
func FuncCSSRad(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Rad
//go:noescape
func CallCSSRad(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rad
//go:noescape
func TryCSSRad(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Turn
//go:noescape
func HasFuncCSSTurn() js.Ref

//go:wasmimport plat/js/web func_CSS_Turn
//go:noescape
func FuncCSSTurn(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Turn
//go:noescape
func CallCSSTurn(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Turn
//go:noescape
func TryCSSTurn(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_S
//go:noescape
func HasFuncCSSS() js.Ref

//go:wasmimport plat/js/web func_CSS_S
//go:noescape
func FuncCSSS(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_S
//go:noescape
func CallCSSS(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_S
//go:noescape
func TryCSSS(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Ms
//go:noescape
func HasFuncCSSMs() js.Ref

//go:wasmimport plat/js/web func_CSS_Ms
//go:noescape
func FuncCSSMs(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Ms
//go:noescape
func CallCSSMs(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Ms
//go:noescape
func TryCSSMs(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Hz
//go:noescape
func HasFuncCSSHz() js.Ref

//go:wasmimport plat/js/web func_CSS_Hz
//go:noescape
func FuncCSSHz(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Hz
//go:noescape
func CallCSSHz(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Hz
//go:noescape
func TryCSSHz(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_KHz
//go:noescape
func HasFuncCSSKHz() js.Ref

//go:wasmimport plat/js/web func_CSS_KHz
//go:noescape
func FuncCSSKHz(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_KHz
//go:noescape
func CallCSSKHz(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_KHz
//go:noescape
func TryCSSKHz(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Dpi
//go:noescape
func HasFuncCSSDpi() js.Ref

//go:wasmimport plat/js/web func_CSS_Dpi
//go:noescape
func FuncCSSDpi(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Dpi
//go:noescape
func CallCSSDpi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dpi
//go:noescape
func TryCSSDpi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Dpcm
//go:noescape
func HasFuncCSSDpcm() js.Ref

//go:wasmimport plat/js/web func_CSS_Dpcm
//go:noescape
func FuncCSSDpcm(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Dpcm
//go:noescape
func CallCSSDpcm(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dpcm
//go:noescape
func TryCSSDpcm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Dppx
//go:noescape
func HasFuncCSSDppx() js.Ref

//go:wasmimport plat/js/web func_CSS_Dppx
//go:noescape
func FuncCSSDppx(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Dppx
//go:noescape
func CallCSSDppx(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dppx
//go:noescape
func TryCSSDppx(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_Fr
//go:noescape
func HasFuncCSSFr() js.Ref

//go:wasmimport plat/js/web func_CSS_Fr
//go:noescape
func FuncCSSFr(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_Fr
//go:noescape
func CallCSSFr(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Fr
//go:noescape
func TryCSSFr(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseStylesheet
//go:noescape
func HasFuncCSSParseStylesheet() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseStylesheet
//go:noescape
func FuncCSSParseStylesheet(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseStylesheet
//go:noescape
func CallCSSParseStylesheet(
	retPtr unsafe.Pointer,
	css js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CSS_ParseStylesheet
//go:noescape
func TryCSSParseStylesheet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseStylesheet1
//go:noescape
func HasFuncCSSParseStylesheet1() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseStylesheet1
//go:noescape
func FuncCSSParseStylesheet1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseStylesheet1
//go:noescape
func CallCSSParseStylesheet1(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseStylesheet1
//go:noescape
func TryCSSParseStylesheet1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseRuleList
//go:noescape
func HasFuncCSSParseRuleList() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRuleList
//go:noescape
func FuncCSSParseRuleList(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseRuleList
//go:noescape
func CallCSSParseRuleList(
	retPtr unsafe.Pointer,
	css js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CSS_ParseRuleList
//go:noescape
func TryCSSParseRuleList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseRuleList1
//go:noescape
func HasFuncCSSParseRuleList1() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRuleList1
//go:noescape
func FuncCSSParseRuleList1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseRuleList1
//go:noescape
func CallCSSParseRuleList1(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseRuleList1
//go:noescape
func TryCSSParseRuleList1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseRule
//go:noescape
func HasFuncCSSParseRule() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRule
//go:noescape
func FuncCSSParseRule(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseRule
//go:noescape
func CallCSSParseRule(
	retPtr unsafe.Pointer,
	css js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CSS_ParseRule
//go:noescape
func TryCSSParseRule(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseRule1
//go:noescape
func HasFuncCSSParseRule1() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRule1
//go:noescape
func FuncCSSParseRule1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseRule1
//go:noescape
func CallCSSParseRule1(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseRule1
//go:noescape
func TryCSSParseRule1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseDeclarationList
//go:noescape
func HasFuncCSSParseDeclarationList() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclarationList
//go:noescape
func FuncCSSParseDeclarationList(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseDeclarationList
//go:noescape
func CallCSSParseDeclarationList(
	retPtr unsafe.Pointer,
	css js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CSS_ParseDeclarationList
//go:noescape
func TryCSSParseDeclarationList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseDeclarationList1
//go:noescape
func HasFuncCSSParseDeclarationList1() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclarationList1
//go:noescape
func FuncCSSParseDeclarationList1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseDeclarationList1
//go:noescape
func CallCSSParseDeclarationList1(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseDeclarationList1
//go:noescape
func TryCSSParseDeclarationList1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseDeclaration
//go:noescape
func HasFuncCSSParseDeclaration() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclaration
//go:noescape
func FuncCSSParseDeclaration(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseDeclaration
//go:noescape
func CallCSSParseDeclaration(
	retPtr unsafe.Pointer,
	css js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CSS_ParseDeclaration
//go:noescape
func TryCSSParseDeclaration(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref,
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseDeclaration1
//go:noescape
func HasFuncCSSParseDeclaration1() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclaration1
//go:noescape
func FuncCSSParseDeclaration1(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseDeclaration1
//go:noescape
func CallCSSParseDeclaration1(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseDeclaration1
//go:noescape
func TryCSSParseDeclaration1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseValue
//go:noescape
func HasFuncCSSParseValue() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseValue
//go:noescape
func FuncCSSParseValue(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseValue
//go:noescape
func CallCSSParseValue(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseValue
//go:noescape
func TryCSSParseValue(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseValueList
//go:noescape
func HasFuncCSSParseValueList() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseValueList
//go:noescape
func FuncCSSParseValueList(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseValueList
//go:noescape
func CallCSSParseValueList(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseValueList
//go:noescape
func TryCSSParseValueList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) js.Ref

//go:wasmimport plat/js/web has_CSS_ParseCommaValueList
//go:noescape
func HasFuncCSSParseCommaValueList() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseCommaValueList
//go:noescape
func FuncCSSParseCommaValueList(fn unsafe.Pointer)

//go:wasmimport plat/js/web call_CSS_ParseCommaValueList
//go:noescape
func CallCSSParseCommaValueList(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseCommaValueList
//go:noescape
func TryCSSParseCommaValueList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSAnimation_CSSAnimation
//go:noescape
func NewCSSAnimationByCSSAnimation(
	effect js.Ref,
	timeline js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSAnimation_CSSAnimation1
//go:noescape
func NewCSSAnimationByCSSAnimation1(
	effect js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSAnimation_CSSAnimation2
//go:noescape
func NewCSSAnimationByCSSAnimation2() js.Ref

//go:wasmimport plat/js/web get_CSSAnimation_AnimationName
//go:noescape
func GetCSSAnimationAnimationName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

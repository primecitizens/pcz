// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

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
func HasBluetoothGetAvailability(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_GetAvailability
//go:noescape
func BluetoothGetAvailabilityFunc(this js.Ref) js.Ref

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
func HasBluetoothGetDevices(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_GetDevices
//go:noescape
func BluetoothGetDevicesFunc(this js.Ref) js.Ref

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
func HasBluetoothRequestDevice(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_RequestDevice
//go:noescape
func BluetoothRequestDeviceFunc(this js.Ref) js.Ref

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
func HasBluetoothRequestDevice1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Bluetooth_RequestDevice1
//go:noescape
func BluetoothRequestDevice1Func(this js.Ref) js.Ref

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
func HasBluetoothUUIDGetService(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_GetService
//go:noescape
func BluetoothUUIDGetServiceFunc(this js.Ref) js.Ref

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
func HasBluetoothUUIDGetCharacteristic(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_GetCharacteristic
//go:noescape
func BluetoothUUIDGetCharacteristicFunc(this js.Ref) js.Ref

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
func HasBluetoothUUIDGetDescriptor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_GetDescriptor
//go:noescape
func BluetoothUUIDGetDescriptorFunc(this js.Ref) js.Ref

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
func HasBluetoothUUIDCanonicalUUID(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BluetoothUUID_CanonicalUUID
//go:noescape
func BluetoothUUIDCanonicalUUIDFunc(this js.Ref) js.Ref

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
func HasLayoutChildIntrinsicSizes(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_LayoutChild_IntrinsicSizes
//go:noescape
func LayoutChildIntrinsicSizesFunc(this js.Ref) js.Ref

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
func HasLayoutChildLayoutNextFragment(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_LayoutChild_LayoutNextFragment
//go:noescape
func LayoutChildLayoutNextFragmentFunc(this js.Ref) js.Ref

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
func HasBroadcastChannelPostMessage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BroadcastChannel_PostMessage
//go:noescape
func BroadcastChannelPostMessageFunc(this js.Ref) js.Ref

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
func HasBroadcastChannelClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BroadcastChannel_Close
//go:noescape
func BroadcastChannelCloseFunc(this js.Ref) js.Ref

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
func HasCropTargetFromElement(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CropTarget_FromElement
//go:noescape
func CropTargetFromElementFunc(this js.Ref) js.Ref

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
func HasRestrictionTargetFromElement(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_RestrictionTarget_FromElement
//go:noescape
func RestrictionTargetFromElementFunc(this js.Ref) js.Ref

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
func HasBrowserCaptureMediaStreamTrackCropTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BrowserCaptureMediaStreamTrack_CropTo
//go:noescape
func BrowserCaptureMediaStreamTrackCropToFunc(this js.Ref) js.Ref

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
func HasBrowserCaptureMediaStreamTrackClone(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BrowserCaptureMediaStreamTrack_Clone
//go:noescape
func BrowserCaptureMediaStreamTrackCloneFunc(this js.Ref) js.Ref

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
func HasBrowserCaptureMediaStreamTrackRestrictTo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_BrowserCaptureMediaStreamTrack_RestrictTo
//go:noescape
func BrowserCaptureMediaStreamTrackRestrictToFunc(this js.Ref) js.Ref

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
func HasCSPViolationReportBodyToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSPViolationReportBody_ToJSON
//go:noescape
func CSPViolationReportBodyToJSONFunc(this js.Ref) js.Ref

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
func HasCSSParserDeclarationToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSParserDeclaration_ToString
//go:noescape
func CSSParserDeclarationToStringFunc(this js.Ref) js.Ref

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
func HasWorkletAddModule(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Worklet_AddModule
//go:noescape
func WorkletAddModuleFunc(this js.Ref) js.Ref

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
func HasWorkletAddModule1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Worklet_AddModule1
//go:noescape
func WorkletAddModule1Func(this js.Ref) js.Ref

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
func GetCSSElementSources(
	retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_CSS_AnimationWorklet
//go:noescape
func GetCSSAnimationWorklet(
	retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_CSS_PaintWorklet
//go:noescape
func GetCSSPaintWorklet(
	retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_CSS_LayoutWorklet
//go:noescape
func GetCSSLayoutWorklet(
	retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_CSS_Highlights
//go:noescape
func GetCSSHighlights(
	retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web has_CSS_Escape
//go:noescape
func HasCSSEscape() js.Ref

//go:wasmimport plat/js/web func_CSS_Escape
//go:noescape
func CSSEscapeFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Escape
//go:noescape
func CallCSSEscape(
	retPtr unsafe.Pointer,
	ident js.Ref)

//go:wasmimport plat/js/web try_CSS_Escape
//go:noescape
func TryCSSEscape(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	ident js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_RegisterProperty
//go:noescape
func HasCSSRegisterProperty() js.Ref

//go:wasmimport plat/js/web func_CSS_RegisterProperty
//go:noescape
func CSSRegisterPropertyFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_RegisterProperty
//go:noescape
func CallCSSRegisterProperty(
	retPtr unsafe.Pointer,
	definition unsafe.Pointer)

//go:wasmimport plat/js/web try_CSS_RegisterProperty
//go:noescape
func TryCSSRegisterProperty(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	definition unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Supports
//go:noescape
func HasCSSSupports() js.Ref

//go:wasmimport plat/js/web func_CSS_Supports
//go:noescape
func CSSSupportsFunc() js.Ref

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
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Supports1
//go:noescape
func HasCSSSupports1() js.Ref

//go:wasmimport plat/js/web func_CSS_Supports1
//go:noescape
func CSSSupports1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_Supports1
//go:noescape
func CallCSSSupports1(
	retPtr unsafe.Pointer,
	conditionText js.Ref)

//go:wasmimport plat/js/web try_CSS_Supports1
//go:noescape
func TryCSSSupports1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	conditionText js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Number
//go:noescape
func HasCSSNumber() js.Ref

//go:wasmimport plat/js/web func_CSS_Number
//go:noescape
func CSSNumberFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Number
//go:noescape
func CallCSSNumber(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Number
//go:noescape
func TryCSSNumber(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Percent
//go:noescape
func HasCSSPercent() js.Ref

//go:wasmimport plat/js/web func_CSS_Percent
//go:noescape
func CSSPercentFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Percent
//go:noescape
func CallCSSPercent(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Percent
//go:noescape
func TryCSSPercent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Cap
//go:noescape
func HasCSSCap() js.Ref

//go:wasmimport plat/js/web func_CSS_Cap
//go:noescape
func CSSCapFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cap
//go:noescape
func CallCSSCap(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cap
//go:noescape
func TryCSSCap(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Ch
//go:noescape
func HasCSSCh() js.Ref

//go:wasmimport plat/js/web func_CSS_Ch
//go:noescape
func CSSChFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Ch
//go:noescape
func CallCSSCh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Ch
//go:noescape
func TryCSSCh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Em
//go:noescape
func HasCSSEm() js.Ref

//go:wasmimport plat/js/web func_CSS_Em
//go:noescape
func CSSEmFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Em
//go:noescape
func CallCSSEm(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Em
//go:noescape
func TryCSSEm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Ex
//go:noescape
func HasCSSEx() js.Ref

//go:wasmimport plat/js/web func_CSS_Ex
//go:noescape
func CSSExFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Ex
//go:noescape
func CallCSSEx(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Ex
//go:noescape
func TryCSSEx(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Ic
//go:noescape
func HasCSSIc() js.Ref

//go:wasmimport plat/js/web func_CSS_Ic
//go:noescape
func CSSIcFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Ic
//go:noescape
func CallCSSIc(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Ic
//go:noescape
func TryCSSIc(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Lh
//go:noescape
func HasCSSLh() js.Ref

//go:wasmimport plat/js/web func_CSS_Lh
//go:noescape
func CSSLhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lh
//go:noescape
func CallCSSLh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lh
//go:noescape
func TryCSSLh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Rcap
//go:noescape
func HasCSSRcap() js.Ref

//go:wasmimport plat/js/web func_CSS_Rcap
//go:noescape
func CSSRcapFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rcap
//go:noescape
func CallCSSRcap(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rcap
//go:noescape
func TryCSSRcap(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Rch
//go:noescape
func HasCSSRch() js.Ref

//go:wasmimport plat/js/web func_CSS_Rch
//go:noescape
func CSSRchFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rch
//go:noescape
func CallCSSRch(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rch
//go:noescape
func TryCSSRch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Rem
//go:noescape
func HasCSSRem() js.Ref

//go:wasmimport plat/js/web func_CSS_Rem
//go:noescape
func CSSRemFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rem
//go:noescape
func CallCSSRem(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rem
//go:noescape
func TryCSSRem(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Rex
//go:noescape
func HasCSSRex() js.Ref

//go:wasmimport plat/js/web func_CSS_Rex
//go:noescape
func CSSRexFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rex
//go:noescape
func CallCSSRex(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rex
//go:noescape
func TryCSSRex(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Ric
//go:noescape
func HasCSSRic() js.Ref

//go:wasmimport plat/js/web func_CSS_Ric
//go:noescape
func CSSRicFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Ric
//go:noescape
func CallCSSRic(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Ric
//go:noescape
func TryCSSRic(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Rlh
//go:noescape
func HasCSSRlh() js.Ref

//go:wasmimport plat/js/web func_CSS_Rlh
//go:noescape
func CSSRlhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rlh
//go:noescape
func CallCSSRlh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rlh
//go:noescape
func TryCSSRlh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Vw
//go:noescape
func HasCSSVw() js.Ref

//go:wasmimport plat/js/web func_CSS_Vw
//go:noescape
func CSSVwFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vw
//go:noescape
func CallCSSVw(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vw
//go:noescape
func TryCSSVw(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Vh
//go:noescape
func HasCSSVh() js.Ref

//go:wasmimport plat/js/web func_CSS_Vh
//go:noescape
func CSSVhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vh
//go:noescape
func CallCSSVh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vh
//go:noescape
func TryCSSVh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Vi
//go:noescape
func HasCSSVi() js.Ref

//go:wasmimport plat/js/web func_CSS_Vi
//go:noescape
func CSSViFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vi
//go:noescape
func CallCSSVi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vi
//go:noescape
func TryCSSVi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Vb
//go:noescape
func HasCSSVb() js.Ref

//go:wasmimport plat/js/web func_CSS_Vb
//go:noescape
func CSSVbFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vb
//go:noescape
func CallCSSVb(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vb
//go:noescape
func TryCSSVb(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Vmin
//go:noescape
func HasCSSVmin() js.Ref

//go:wasmimport plat/js/web func_CSS_Vmin
//go:noescape
func CSSVminFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vmin
//go:noescape
func CallCSSVmin(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vmin
//go:noescape
func TryCSSVmin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Vmax
//go:noescape
func HasCSSVmax() js.Ref

//go:wasmimport plat/js/web func_CSS_Vmax
//go:noescape
func CSSVmaxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Vmax
//go:noescape
func CallCSSVmax(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Vmax
//go:noescape
func TryCSSVmax(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Svw
//go:noescape
func HasCSSSvw() js.Ref

//go:wasmimport plat/js/web func_CSS_Svw
//go:noescape
func CSSSvwFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svw
//go:noescape
func CallCSSSvw(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svw
//go:noescape
func TryCSSSvw(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Svh
//go:noescape
func HasCSSSvh() js.Ref

//go:wasmimport plat/js/web func_CSS_Svh
//go:noescape
func CSSSvhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svh
//go:noescape
func CallCSSSvh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svh
//go:noescape
func TryCSSSvh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Svi
//go:noescape
func HasCSSSvi() js.Ref

//go:wasmimport plat/js/web func_CSS_Svi
//go:noescape
func CSSSviFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svi
//go:noescape
func CallCSSSvi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svi
//go:noescape
func TryCSSSvi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Svb
//go:noescape
func HasCSSSvb() js.Ref

//go:wasmimport plat/js/web func_CSS_Svb
//go:noescape
func CSSSvbFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svb
//go:noescape
func CallCSSSvb(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svb
//go:noescape
func TryCSSSvb(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Svmin
//go:noescape
func HasCSSSvmin() js.Ref

//go:wasmimport plat/js/web func_CSS_Svmin
//go:noescape
func CSSSvminFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svmin
//go:noescape
func CallCSSSvmin(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svmin
//go:noescape
func TryCSSSvmin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Svmax
//go:noescape
func HasCSSSvmax() js.Ref

//go:wasmimport plat/js/web func_CSS_Svmax
//go:noescape
func CSSSvmaxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Svmax
//go:noescape
func CallCSSSvmax(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Svmax
//go:noescape
func TryCSSSvmax(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Lvw
//go:noescape
func HasCSSLvw() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvw
//go:noescape
func CSSLvwFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvw
//go:noescape
func CallCSSLvw(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvw
//go:noescape
func TryCSSLvw(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Lvh
//go:noescape
func HasCSSLvh() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvh
//go:noescape
func CSSLvhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvh
//go:noescape
func CallCSSLvh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvh
//go:noescape
func TryCSSLvh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Lvi
//go:noescape
func HasCSSLvi() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvi
//go:noescape
func CSSLviFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvi
//go:noescape
func CallCSSLvi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvi
//go:noescape
func TryCSSLvi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Lvb
//go:noescape
func HasCSSLvb() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvb
//go:noescape
func CSSLvbFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvb
//go:noescape
func CallCSSLvb(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvb
//go:noescape
func TryCSSLvb(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Lvmin
//go:noescape
func HasCSSLvmin() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvmin
//go:noescape
func CSSLvminFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvmin
//go:noescape
func CallCSSLvmin(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvmin
//go:noescape
func TryCSSLvmin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Lvmax
//go:noescape
func HasCSSLvmax() js.Ref

//go:wasmimport plat/js/web func_CSS_Lvmax
//go:noescape
func CSSLvmaxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Lvmax
//go:noescape
func CallCSSLvmax(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Lvmax
//go:noescape
func TryCSSLvmax(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Dvw
//go:noescape
func HasCSSDvw() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvw
//go:noescape
func CSSDvwFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvw
//go:noescape
func CallCSSDvw(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvw
//go:noescape
func TryCSSDvw(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Dvh
//go:noescape
func HasCSSDvh() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvh
//go:noescape
func CSSDvhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvh
//go:noescape
func CallCSSDvh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvh
//go:noescape
func TryCSSDvh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Dvi
//go:noescape
func HasCSSDvi() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvi
//go:noescape
func CSSDviFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvi
//go:noescape
func CallCSSDvi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvi
//go:noescape
func TryCSSDvi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Dvb
//go:noescape
func HasCSSDvb() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvb
//go:noescape
func CSSDvbFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvb
//go:noescape
func CallCSSDvb(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvb
//go:noescape
func TryCSSDvb(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Dvmin
//go:noescape
func HasCSSDvmin() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvmin
//go:noescape
func CSSDvminFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvmin
//go:noescape
func CallCSSDvmin(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvmin
//go:noescape
func TryCSSDvmin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Dvmax
//go:noescape
func HasCSSDvmax() js.Ref

//go:wasmimport plat/js/web func_CSS_Dvmax
//go:noescape
func CSSDvmaxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dvmax
//go:noescape
func CallCSSDvmax(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dvmax
//go:noescape
func TryCSSDvmax(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Cqw
//go:noescape
func HasCSSCqw() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqw
//go:noescape
func CSSCqwFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqw
//go:noescape
func CallCSSCqw(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqw
//go:noescape
func TryCSSCqw(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Cqh
//go:noescape
func HasCSSCqh() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqh
//go:noescape
func CSSCqhFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqh
//go:noescape
func CallCSSCqh(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqh
//go:noescape
func TryCSSCqh(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Cqi
//go:noescape
func HasCSSCqi() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqi
//go:noescape
func CSSCqiFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqi
//go:noescape
func CallCSSCqi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqi
//go:noescape
func TryCSSCqi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Cqb
//go:noescape
func HasCSSCqb() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqb
//go:noescape
func CSSCqbFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqb
//go:noescape
func CallCSSCqb(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqb
//go:noescape
func TryCSSCqb(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Cqmin
//go:noescape
func HasCSSCqmin() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqmin
//go:noescape
func CSSCqminFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqmin
//go:noescape
func CallCSSCqmin(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqmin
//go:noescape
func TryCSSCqmin(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Cqmax
//go:noescape
func HasCSSCqmax() js.Ref

//go:wasmimport plat/js/web func_CSS_Cqmax
//go:noescape
func CSSCqmaxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cqmax
//go:noescape
func CallCSSCqmax(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cqmax
//go:noescape
func TryCSSCqmax(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Cm
//go:noescape
func HasCSSCm() js.Ref

//go:wasmimport plat/js/web func_CSS_Cm
//go:noescape
func CSSCmFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Cm
//go:noescape
func CallCSSCm(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Cm
//go:noescape
func TryCSSCm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Mm
//go:noescape
func HasCSSMm() js.Ref

//go:wasmimport plat/js/web func_CSS_Mm
//go:noescape
func CSSMmFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Mm
//go:noescape
func CallCSSMm(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Mm
//go:noescape
func TryCSSMm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Q
//go:noescape
func HasCSSQ() js.Ref

//go:wasmimport plat/js/web func_CSS_Q
//go:noescape
func CSSQFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Q
//go:noescape
func CallCSSQ(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Q
//go:noescape
func TryCSSQ(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_In
//go:noescape
func HasCSSIn() js.Ref

//go:wasmimport plat/js/web func_CSS_In
//go:noescape
func CSSInFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_In
//go:noescape
func CallCSSIn(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_In
//go:noescape
func TryCSSIn(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Pt
//go:noescape
func HasCSSPt() js.Ref

//go:wasmimport plat/js/web func_CSS_Pt
//go:noescape
func CSSPtFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Pt
//go:noescape
func CallCSSPt(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Pt
//go:noescape
func TryCSSPt(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Pc
//go:noescape
func HasCSSPc() js.Ref

//go:wasmimport plat/js/web func_CSS_Pc
//go:noescape
func CSSPcFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Pc
//go:noescape
func CallCSSPc(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Pc
//go:noescape
func TryCSSPc(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Px
//go:noescape
func HasCSSPx() js.Ref

//go:wasmimport plat/js/web func_CSS_Px
//go:noescape
func CSSPxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Px
//go:noescape
func CallCSSPx(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Px
//go:noescape
func TryCSSPx(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Deg
//go:noescape
func HasCSSDeg() js.Ref

//go:wasmimport plat/js/web func_CSS_Deg
//go:noescape
func CSSDegFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Deg
//go:noescape
func CallCSSDeg(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Deg
//go:noescape
func TryCSSDeg(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Grad
//go:noescape
func HasCSSGrad() js.Ref

//go:wasmimport plat/js/web func_CSS_Grad
//go:noescape
func CSSGradFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Grad
//go:noescape
func CallCSSGrad(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Grad
//go:noescape
func TryCSSGrad(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Rad
//go:noescape
func HasCSSRad() js.Ref

//go:wasmimport plat/js/web func_CSS_Rad
//go:noescape
func CSSRadFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Rad
//go:noescape
func CallCSSRad(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Rad
//go:noescape
func TryCSSRad(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Turn
//go:noescape
func HasCSSTurn() js.Ref

//go:wasmimport plat/js/web func_CSS_Turn
//go:noescape
func CSSTurnFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Turn
//go:noescape
func CallCSSTurn(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Turn
//go:noescape
func TryCSSTurn(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_S
//go:noescape
func HasCSSS() js.Ref

//go:wasmimport plat/js/web func_CSS_S
//go:noescape
func CSSSFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_S
//go:noescape
func CallCSSS(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_S
//go:noescape
func TryCSSS(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Ms
//go:noescape
func HasCSSMs() js.Ref

//go:wasmimport plat/js/web func_CSS_Ms
//go:noescape
func CSSMsFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Ms
//go:noescape
func CallCSSMs(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Ms
//go:noescape
func TryCSSMs(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Hz
//go:noescape
func HasCSSHz() js.Ref

//go:wasmimport plat/js/web func_CSS_Hz
//go:noescape
func CSSHzFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Hz
//go:noescape
func CallCSSHz(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Hz
//go:noescape
func TryCSSHz(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_KHz
//go:noescape
func HasCSSKHz() js.Ref

//go:wasmimport plat/js/web func_CSS_KHz
//go:noescape
func CSSKHzFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_KHz
//go:noescape
func CallCSSKHz(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_KHz
//go:noescape
func TryCSSKHz(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Dpi
//go:noescape
func HasCSSDpi() js.Ref

//go:wasmimport plat/js/web func_CSS_Dpi
//go:noescape
func CSSDpiFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dpi
//go:noescape
func CallCSSDpi(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dpi
//go:noescape
func TryCSSDpi(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Dpcm
//go:noescape
func HasCSSDpcm() js.Ref

//go:wasmimport plat/js/web func_CSS_Dpcm
//go:noescape
func CSSDpcmFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dpcm
//go:noescape
func CallCSSDpcm(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dpcm
//go:noescape
func TryCSSDpcm(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Dppx
//go:noescape
func HasCSSDppx() js.Ref

//go:wasmimport plat/js/web func_CSS_Dppx
//go:noescape
func CSSDppxFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Dppx
//go:noescape
func CallCSSDppx(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Dppx
//go:noescape
func TryCSSDppx(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_Fr
//go:noescape
func HasCSSFr() js.Ref

//go:wasmimport plat/js/web func_CSS_Fr
//go:noescape
func CSSFrFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_Fr
//go:noescape
func CallCSSFr(
	retPtr unsafe.Pointer,
	value float64)

//go:wasmimport plat/js/web try_CSS_Fr
//go:noescape
func TryCSSFr(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value float64) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseStylesheet
//go:noescape
func HasCSSParseStylesheet() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseStylesheet
//go:noescape
func CSSParseStylesheetFunc() js.Ref

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
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseStylesheet1
//go:noescape
func HasCSSParseStylesheet1() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseStylesheet1
//go:noescape
func CSSParseStylesheet1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseStylesheet1
//go:noescape
func CallCSSParseStylesheet1(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseStylesheet1
//go:noescape
func TryCSSParseStylesheet1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseRuleList
//go:noescape
func HasCSSParseRuleList() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRuleList
//go:noescape
func CSSParseRuleListFunc() js.Ref

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
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseRuleList1
//go:noescape
func HasCSSParseRuleList1() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRuleList1
//go:noescape
func CSSParseRuleList1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseRuleList1
//go:noescape
func CallCSSParseRuleList1(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseRuleList1
//go:noescape
func TryCSSParseRuleList1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseRule
//go:noescape
func HasCSSParseRule() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRule
//go:noescape
func CSSParseRuleFunc() js.Ref

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
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseRule1
//go:noescape
func HasCSSParseRule1() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseRule1
//go:noescape
func CSSParseRule1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseRule1
//go:noescape
func CallCSSParseRule1(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseRule1
//go:noescape
func TryCSSParseRule1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseDeclarationList
//go:noescape
func HasCSSParseDeclarationList() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclarationList
//go:noescape
func CSSParseDeclarationListFunc() js.Ref

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
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseDeclarationList1
//go:noescape
func HasCSSParseDeclarationList1() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclarationList1
//go:noescape
func CSSParseDeclarationList1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseDeclarationList1
//go:noescape
func CallCSSParseDeclarationList1(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseDeclarationList1
//go:noescape
func TryCSSParseDeclarationList1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseDeclaration
//go:noescape
func HasCSSParseDeclaration() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclaration
//go:noescape
func CSSParseDeclarationFunc() js.Ref

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
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseDeclaration1
//go:noescape
func HasCSSParseDeclaration1() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseDeclaration1
//go:noescape
func CSSParseDeclaration1Func() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseDeclaration1
//go:noescape
func CallCSSParseDeclaration1(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseDeclaration1
//go:noescape
func TryCSSParseDeclaration1(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseValue
//go:noescape
func HasCSSParseValue() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseValue
//go:noescape
func CSSParseValueFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseValue
//go:noescape
func CallCSSParseValue(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseValue
//go:noescape
func TryCSSParseValue(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseValueList
//go:noescape
func HasCSSParseValueList() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseValueList
//go:noescape
func CSSParseValueListFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseValueList
//go:noescape
func CallCSSParseValueList(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseValueList
//go:noescape
func TryCSSParseValueList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSS_ParseCommaValueList
//go:noescape
func HasCSSParseCommaValueList() js.Ref

//go:wasmimport plat/js/web func_CSS_ParseCommaValueList
//go:noescape
func CSSParseCommaValueListFunc() js.Ref

//go:wasmimport plat/js/web call_CSS_ParseCommaValueList
//go:noescape
func CallCSSParseCommaValueList(
	retPtr unsafe.Pointer,
	css js.Ref)

//go:wasmimport plat/js/web try_CSS_ParseCommaValueList
//go:noescape
func TryCSSParseCommaValueList(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	css js.Ref) (ok js.Ref)

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

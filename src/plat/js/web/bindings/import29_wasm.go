// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

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

//go:wasmimport plat/js/web constof_FlowControlType
//go:noescape

func ConstOfFlowControlType(str js.Ref) uint32

//go:wasmimport plat/js/web store_SerialOptions
//go:noescape

func SerialOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SerialOptions
//go:noescape

func SerialOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_SerialOutputSignals
//go:noescape

func SerialOutputSignalsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SerialOutputSignals
//go:noescape

func SerialOutputSignalsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_SerialInputSignals
//go:noescape

func SerialInputSignalsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SerialInputSignals
//go:noescape

func SerialInputSignalsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_SerialPort_Readable
//go:noescape

func GetSerialPortReadable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SerialPort_Writable
//go:noescape

func GetSerialPortWritable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_SerialPort_GetInfo
//go:noescape

func CallSerialPortGetInfo(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SerialPort_GetInfo
//go:noescape

func SerialPortGetInfoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_Open
//go:noescape

func CallSerialPortOpen(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_SerialPort_Open
//go:noescape

func SerialPortOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_SetSignals
//go:noescape

func CallSerialPortSetSignals(
	this js.Ref,
	ptr unsafe.Pointer,

	signals unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_SerialPort_SetSignals
//go:noescape

func SerialPortSetSignalsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_SetSignals1
//go:noescape

func CallSerialPortSetSignals1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SerialPort_SetSignals1
//go:noescape

func SerialPortSetSignals1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_GetSignals
//go:noescape

func CallSerialPortGetSignals(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SerialPort_GetSignals
//go:noescape

func SerialPortGetSignalsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_Close
//go:noescape

func CallSerialPortClose(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SerialPort_Close
//go:noescape

func SerialPortCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_Forget
//go:noescape

func CallSerialPortForget(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SerialPort_Forget
//go:noescape

func SerialPortForgetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_SerialPortFilter
//go:noescape

func SerialPortFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SerialPortFilter
//go:noescape

func SerialPortFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_SerialPortRequestOptions
//go:noescape

func SerialPortRequestOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SerialPortRequestOptions
//go:noescape

func SerialPortRequestOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_Serial_GetPorts
//go:noescape

func CallSerialGetPorts(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Serial_GetPorts
//go:noescape

func SerialGetPortsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Serial_RequestPort
//go:noescape

func CallSerialRequestPort(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Serial_RequestPort
//go:noescape

func SerialRequestPortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Serial_RequestPort1
//go:noescape

func CallSerialRequestPort1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Serial_RequestPort1
//go:noescape

func SerialRequestPort1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MediaDecodingType
//go:noescape

func ConstOfMediaDecodingType(str js.Ref) uint32

//go:wasmimport plat/js/web store_KeySystemTrackConfiguration
//go:noescape

func KeySystemTrackConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_KeySystemTrackConfiguration
//go:noescape

func KeySystemTrackConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaCapabilitiesKeySystemConfiguration
//go:noescape

func MediaCapabilitiesKeySystemConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaCapabilitiesKeySystemConfiguration
//go:noescape

func MediaCapabilitiesKeySystemConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_HdrMetadataType
//go:noescape

func ConstOfHdrMetadataType(str js.Ref) uint32

//go:wasmimport plat/js/web constof_TransferFunction
//go:noescape

func ConstOfTransferFunction(str js.Ref) uint32

//go:wasmimport plat/js/web store_VideoConfiguration
//go:noescape

func VideoConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_VideoConfiguration
//go:noescape

func VideoConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaDecodingConfiguration
//go:noescape

func MediaDecodingConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaDecodingConfiguration
//go:noescape

func MediaDecodingConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaCapabilitiesDecodingInfo
//go:noescape

func MediaCapabilitiesDecodingInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaCapabilitiesDecodingInfo
//go:noescape

func MediaCapabilitiesDecodingInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MediaEncodingType
//go:noescape

func ConstOfMediaEncodingType(str js.Ref) uint32

//go:wasmimport plat/js/web store_MediaEncodingConfiguration
//go:noescape

func MediaEncodingConfigurationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaEncodingConfiguration
//go:noescape

func MediaEncodingConfigurationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaCapabilitiesEncodingInfo
//go:noescape

func MediaCapabilitiesEncodingInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaCapabilitiesEncodingInfo
//go:noescape

func MediaCapabilitiesEncodingInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaCapabilities_DecodingInfo
//go:noescape

func CallMediaCapabilitiesDecodingInfo(
	this js.Ref,
	ptr unsafe.Pointer,

	configuration unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_MediaCapabilities_DecodingInfo
//go:noescape

func MediaCapabilitiesDecodingInfoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaCapabilities_EncodingInfo
//go:noescape

func CallMediaCapabilitiesEncodingInfo(
	this js.Ref,
	ptr unsafe.Pointer,

	configuration unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_MediaCapabilities_EncodingInfo
//go:noescape

func MediaCapabilitiesEncodingInfoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_UserActivation_HasBeenActive
//go:noescape

func GetUserActivationHasBeenActive(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_UserActivation_IsActive
//go:noescape

func GetUserActivationIsActive(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_PermissionStatus_State
//go:noescape

func GetPermissionStatusState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_PermissionStatus_Name
//go:noescape

func GetPermissionStatusName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Permissions_Query
//go:noescape

func CallPermissionsQuery(
	this js.Ref,
	ptr unsafe.Pointer,

	permissionDesc js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Permissions_Query
//go:noescape

func PermissionsQueryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Permissions_Revoke
//go:noescape

func CallPermissionsRevoke(
	this js.Ref,
	ptr unsafe.Pointer,

	permissionDesc js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Permissions_Revoke
//go:noescape

func PermissionsRevokeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Permissions_Request
//go:noescape

func CallPermissionsRequest(
	this js.Ref,
	ptr unsafe.Pointer,

	permissionDesc js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Permissions_Request
//go:noescape

func PermissionsRequestFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ContactProperty
//go:noescape

func ConstOfContactProperty(str js.Ref) uint32

//go:wasmimport plat/js/web get_ContactAddress_City
//go:noescape

func GetContactAddressCity(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ContactAddress_Country
//go:noescape

func GetContactAddressCountry(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ContactAddress_DependentLocality
//go:noescape

func GetContactAddressDependentLocality(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ContactAddress_Organization
//go:noescape

func GetContactAddressOrganization(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ContactAddress_Phone
//go:noescape

func GetContactAddressPhone(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ContactAddress_PostalCode
//go:noescape

func GetContactAddressPostalCode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ContactAddress_Recipient
//go:noescape

func GetContactAddressRecipient(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ContactAddress_Region
//go:noescape

func GetContactAddressRegion(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ContactAddress_SortingCode
//go:noescape

func GetContactAddressSortingCode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ContactAddress_AddressLine
//go:noescape

func GetContactAddressAddressLine(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_ContactAddress_ToJSON
//go:noescape

func CallContactAddressToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ContactAddress_ToJSON
//go:noescape

func ContactAddressToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_ContactInfo
//go:noescape

func ContactInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ContactInfo
//go:noescape

func ContactInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ContactsSelectOptions
//go:noescape

func ContactsSelectOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ContactsSelectOptions
//go:noescape

func ContactsSelectOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContactsManager_GetProperties
//go:noescape

func CallContactsManagerGetProperties(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ContactsManager_GetProperties
//go:noescape

func ContactsManagerGetPropertiesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContactsManager_Select
//go:noescape

func CallContactsManagerSelect(
	this js.Ref,
	ptr unsafe.Pointer,

	properties js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_ContactsManager_Select
//go:noescape

func ContactsManagerSelectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContactsManager_Select1
//go:noescape

func CallContactsManagerSelect1(
	this js.Ref,
	ptr unsafe.Pointer,

	properties js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_ContactsManager_Select1
//go:noescape

func ContactsManagerSelect1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Keyboard_Lock
//go:noescape

func CallKeyboardLock(
	this js.Ref,
	ptr unsafe.Pointer,

	keyCodes js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Keyboard_Lock
//go:noescape

func KeyboardLockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Keyboard_Lock1
//go:noescape

func CallKeyboardLock1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Keyboard_Lock1
//go:noescape

func KeyboardLock1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Keyboard_Unlock
//go:noescape

func CallKeyboardUnlock(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Keyboard_Unlock
//go:noescape

func KeyboardUnlockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Keyboard_GetLayoutMap
//go:noescape

func CallKeyboardGetLayoutMap(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Keyboard_GetLayoutMap
//go:noescape

func KeyboardGetLayoutMapFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_MediaSessionAction
//go:noescape

func ConstOfMediaSessionAction(str js.Ref) uint32

//go:wasmimport plat/js/web store_MediaSessionActionDetails
//go:noescape

func MediaSessionActionDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaSessionActionDetails
//go:noescape

func MediaSessionActionDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaPositionState
//go:noescape

func MediaPositionStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaPositionState
//go:noescape

func MediaPositionStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaImage
//go:noescape

func MediaImageJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaImage
//go:noescape

func MediaImageJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_MediaMetadataInit
//go:noescape

func MediaMetadataInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MediaMetadataInit
//go:noescape

func MediaMetadataInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_MediaMetadata_MediaMetadata
//go:noescape

func NewMediaMetadataByMediaMetadata(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_MediaMetadata_MediaMetadata1
//go:noescape

func NewMediaMetadataByMediaMetadata1() js.Ref

//go:wasmimport plat/js/web get_MediaMetadata_Title
//go:noescape

func GetMediaMetadataTitle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_MediaMetadata_Title
//go:noescape

func SetMediaMetadataTitle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaMetadata_Artist
//go:noescape

func GetMediaMetadataArtist(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_MediaMetadata_Artist
//go:noescape

func SetMediaMetadataArtist(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaMetadata_Album
//go:noescape

func GetMediaMetadataAlbum(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_MediaMetadata_Album
//go:noescape

func SetMediaMetadataAlbum(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaMetadata_Artwork
//go:noescape

func GetMediaMetadataArtwork(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_MediaMetadata_Artwork
//go:noescape

func SetMediaMetadataArtwork(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web constof_MediaSessionPlaybackState
//go:noescape

func ConstOfMediaSessionPlaybackState(str js.Ref) uint32

//go:wasmimport plat/js/web get_MediaSession_Metadata
//go:noescape

func GetMediaSessionMetadata(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_MediaSession_Metadata
//go:noescape

func SetMediaSessionMetadata(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaSession_PlaybackState
//go:noescape

func GetMediaSessionPlaybackState(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_MediaSession_PlaybackState
//go:noescape

func SetMediaSessionPlaybackState(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web call_MediaSession_SetActionHandler
//go:noescape

func CallMediaSessionSetActionHandler(
	this js.Ref,
	ptr unsafe.Pointer,

	action uint32,
	handler js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaSession_SetActionHandler
//go:noescape

func MediaSessionSetActionHandlerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSession_SetPositionState
//go:noescape

func CallMediaSessionSetPositionState(
	this js.Ref,
	ptr unsafe.Pointer,

	state unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_MediaSession_SetPositionState
//go:noescape

func MediaSessionSetPositionStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSession_SetPositionState1
//go:noescape

func CallMediaSessionSetPositionState1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_MediaSession_SetPositionState1
//go:noescape

func MediaSessionSetPositionState1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSession_SetMicrophoneActive
//go:noescape

func CallMediaSessionSetMicrophoneActive(
	this js.Ref,
	ptr unsafe.Pointer,

	active js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaSession_SetMicrophoneActive
//go:noescape

func MediaSessionSetMicrophoneActiveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSession_SetCameraActive
//go:noescape

func CallMediaSessionSetCameraActive(
	this js.Ref,
	ptr unsafe.Pointer,

	active js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_MediaSession_SetCameraActive
//go:noescape

func MediaSessionSetCameraActiveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web constof_DevicePostureType
//go:noescape

func ConstOfDevicePostureType(str js.Ref) uint32

//go:wasmimport plat/js/web get_DevicePosture_Type
//go:noescape

func GetDevicePostureType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

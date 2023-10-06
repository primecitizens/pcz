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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SerialPort_Writable
//go:noescape
func GetSerialPortWritable(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SerialPort_GetInfo
//go:noescape
func HasSerialPortGetInfo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SerialPort_GetInfo
//go:noescape
func SerialPortGetInfoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_GetInfo
//go:noescape
func CallSerialPortGetInfo(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SerialPort_GetInfo
//go:noescape
func TrySerialPortGetInfo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SerialPort_Open
//go:noescape
func HasSerialPortOpen(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SerialPort_Open
//go:noescape
func SerialPortOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_Open
//go:noescape
func CallSerialPortOpen(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_SerialPort_Open
//go:noescape
func TrySerialPortOpen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SerialPort_SetSignals
//go:noescape
func HasSerialPortSetSignals(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SerialPort_SetSignals
//go:noescape
func SerialPortSetSignalsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_SetSignals
//go:noescape
func CallSerialPortSetSignals(
	this js.Ref, retPtr unsafe.Pointer,
	signals unsafe.Pointer)

//go:wasmimport plat/js/web try_SerialPort_SetSignals
//go:noescape
func TrySerialPortSetSignals(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	signals unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SerialPort_SetSignals1
//go:noescape
func HasSerialPortSetSignals1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SerialPort_SetSignals1
//go:noescape
func SerialPortSetSignals1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_SetSignals1
//go:noescape
func CallSerialPortSetSignals1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SerialPort_SetSignals1
//go:noescape
func TrySerialPortSetSignals1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SerialPort_GetSignals
//go:noescape
func HasSerialPortGetSignals(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SerialPort_GetSignals
//go:noescape
func SerialPortGetSignalsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_GetSignals
//go:noescape
func CallSerialPortGetSignals(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SerialPort_GetSignals
//go:noescape
func TrySerialPortGetSignals(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SerialPort_Close
//go:noescape
func HasSerialPortClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SerialPort_Close
//go:noescape
func SerialPortCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_Close
//go:noescape
func CallSerialPortClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SerialPort_Close
//go:noescape
func TrySerialPortClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SerialPort_Forget
//go:noescape
func HasSerialPortForget(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SerialPort_Forget
//go:noescape
func SerialPortForgetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SerialPort_Forget
//go:noescape
func CallSerialPortForget(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SerialPort_Forget
//go:noescape
func TrySerialPortForget(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/web has_Serial_GetPorts
//go:noescape
func HasSerialGetPorts(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Serial_GetPorts
//go:noescape
func SerialGetPortsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Serial_GetPorts
//go:noescape
func CallSerialGetPorts(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Serial_GetPorts
//go:noescape
func TrySerialGetPorts(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Serial_RequestPort
//go:noescape
func HasSerialRequestPort(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Serial_RequestPort
//go:noescape
func SerialRequestPortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Serial_RequestPort
//go:noescape
func CallSerialRequestPort(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Serial_RequestPort
//go:noescape
func TrySerialRequestPort(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Serial_RequestPort1
//go:noescape
func HasSerialRequestPort1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Serial_RequestPort1
//go:noescape
func SerialRequestPort1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Serial_RequestPort1
//go:noescape
func CallSerialRequestPort1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Serial_RequestPort1
//go:noescape
func TrySerialRequestPort1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/web has_MediaCapabilities_DecodingInfo
//go:noescape
func HasMediaCapabilitiesDecodingInfo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaCapabilities_DecodingInfo
//go:noescape
func MediaCapabilitiesDecodingInfoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaCapabilities_DecodingInfo
//go:noescape
func CallMediaCapabilitiesDecodingInfo(
	this js.Ref, retPtr unsafe.Pointer,
	configuration unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaCapabilities_DecodingInfo
//go:noescape
func TryMediaCapabilitiesDecodingInfo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	configuration unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaCapabilities_EncodingInfo
//go:noescape
func HasMediaCapabilitiesEncodingInfo(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaCapabilities_EncodingInfo
//go:noescape
func MediaCapabilitiesEncodingInfoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaCapabilities_EncodingInfo
//go:noescape
func CallMediaCapabilitiesEncodingInfo(
	this js.Ref, retPtr unsafe.Pointer,
	configuration unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaCapabilities_EncodingInfo
//go:noescape
func TryMediaCapabilitiesEncodingInfo(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	configuration unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_UserActivation_HasBeenActive
//go:noescape
func GetUserActivationHasBeenActive(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_UserActivation_IsActive
//go:noescape
func GetUserActivationIsActive(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PermissionStatus_State
//go:noescape
func GetPermissionStatusState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_PermissionStatus_Name
//go:noescape
func GetPermissionStatusName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Permissions_Query
//go:noescape
func HasPermissionsQuery(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Permissions_Query
//go:noescape
func PermissionsQueryFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Permissions_Query
//go:noescape
func CallPermissionsQuery(
	this js.Ref, retPtr unsafe.Pointer,
	permissionDesc js.Ref)

//go:wasmimport plat/js/web try_Permissions_Query
//go:noescape
func TryPermissionsQuery(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	permissionDesc js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Permissions_Revoke
//go:noescape
func HasPermissionsRevoke(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Permissions_Revoke
//go:noescape
func PermissionsRevokeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Permissions_Revoke
//go:noescape
func CallPermissionsRevoke(
	this js.Ref, retPtr unsafe.Pointer,
	permissionDesc js.Ref)

//go:wasmimport plat/js/web try_Permissions_Revoke
//go:noescape
func TryPermissionsRevoke(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	permissionDesc js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Permissions_Request
//go:noescape
func HasPermissionsRequest(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Permissions_Request
//go:noescape
func PermissionsRequestFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Permissions_Request
//go:noescape
func CallPermissionsRequest(
	this js.Ref, retPtr unsafe.Pointer,
	permissionDesc js.Ref)

//go:wasmimport plat/js/web try_Permissions_Request
//go:noescape
func TryPermissionsRequest(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	permissionDesc js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_ContactProperty
//go:noescape
func ConstOfContactProperty(str js.Ref) uint32

//go:wasmimport plat/js/web get_ContactAddress_City
//go:noescape
func GetContactAddressCity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ContactAddress_Country
//go:noescape
func GetContactAddressCountry(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ContactAddress_DependentLocality
//go:noescape
func GetContactAddressDependentLocality(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ContactAddress_Organization
//go:noescape
func GetContactAddressOrganization(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ContactAddress_Phone
//go:noescape
func GetContactAddressPhone(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ContactAddress_PostalCode
//go:noescape
func GetContactAddressPostalCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ContactAddress_Recipient
//go:noescape
func GetContactAddressRecipient(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ContactAddress_Region
//go:noescape
func GetContactAddressRegion(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ContactAddress_SortingCode
//go:noescape
func GetContactAddressSortingCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_ContactAddress_AddressLine
//go:noescape
func GetContactAddressAddressLine(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ContactAddress_ToJSON
//go:noescape
func HasContactAddressToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ContactAddress_ToJSON
//go:noescape
func ContactAddressToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContactAddress_ToJSON
//go:noescape
func CallContactAddressToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ContactAddress_ToJSON
//go:noescape
func TryContactAddressToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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

//go:wasmimport plat/js/web has_ContactsManager_GetProperties
//go:noescape
func HasContactsManagerGetProperties(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ContactsManager_GetProperties
//go:noescape
func ContactsManagerGetPropertiesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContactsManager_GetProperties
//go:noescape
func CallContactsManagerGetProperties(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_ContactsManager_GetProperties
//go:noescape
func TryContactsManagerGetProperties(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ContactsManager_Select
//go:noescape
func HasContactsManagerSelect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ContactsManager_Select
//go:noescape
func ContactsManagerSelectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContactsManager_Select
//go:noescape
func CallContactsManagerSelect(
	this js.Ref, retPtr unsafe.Pointer,
	properties js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_ContactsManager_Select
//go:noescape
func TryContactsManagerSelect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	properties js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_ContactsManager_Select1
//go:noescape
func HasContactsManagerSelect1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_ContactsManager_Select1
//go:noescape
func ContactsManagerSelect1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_ContactsManager_Select1
//go:noescape
func CallContactsManagerSelect1(
	this js.Ref, retPtr unsafe.Pointer,
	properties js.Ref)

//go:wasmimport plat/js/web try_ContactsManager_Select1
//go:noescape
func TryContactsManagerSelect1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	properties js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Keyboard_Lock
//go:noescape
func HasKeyboardLock(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Keyboard_Lock
//go:noescape
func KeyboardLockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Keyboard_Lock
//go:noescape
func CallKeyboardLock(
	this js.Ref, retPtr unsafe.Pointer,
	keyCodes js.Ref)

//go:wasmimport plat/js/web try_Keyboard_Lock
//go:noescape
func TryKeyboardLock(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	keyCodes js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Keyboard_Lock1
//go:noescape
func HasKeyboardLock1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Keyboard_Lock1
//go:noescape
func KeyboardLock1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Keyboard_Lock1
//go:noescape
func CallKeyboardLock1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Keyboard_Lock1
//go:noescape
func TryKeyboardLock1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Keyboard_Unlock
//go:noescape
func HasKeyboardUnlock(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Keyboard_Unlock
//go:noescape
func KeyboardUnlockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Keyboard_Unlock
//go:noescape
func CallKeyboardUnlock(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Keyboard_Unlock
//go:noescape
func TryKeyboardUnlock(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Keyboard_GetLayoutMap
//go:noescape
func HasKeyboardGetLayoutMap(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Keyboard_GetLayoutMap
//go:noescape
func KeyboardGetLayoutMapFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Keyboard_GetLayoutMap
//go:noescape
func CallKeyboardGetLayoutMap(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Keyboard_GetLayoutMap
//go:noescape
func TryKeyboardGetLayoutMap(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_MediaMetadata_Title
//go:noescape
func SetMediaMetadataTitle(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaMetadata_Artist
//go:noescape
func GetMediaMetadataArtist(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_MediaMetadata_Artist
//go:noescape
func SetMediaMetadataArtist(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaMetadata_Album
//go:noescape
func GetMediaMetadataAlbum(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_MediaMetadata_Album
//go:noescape
func SetMediaMetadataAlbum(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaMetadata_Artwork
//go:noescape
func GetMediaMetadataArtwork(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

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
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_MediaSession_Metadata
//go:noescape
func SetMediaSessionMetadata(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_MediaSession_PlaybackState
//go:noescape
func GetMediaSessionPlaybackState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_MediaSession_PlaybackState
//go:noescape
func SetMediaSessionPlaybackState(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web has_MediaSession_SetActionHandler
//go:noescape
func HasMediaSessionSetActionHandler(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSession_SetActionHandler
//go:noescape
func MediaSessionSetActionHandlerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSession_SetActionHandler
//go:noescape
func CallMediaSessionSetActionHandler(
	this js.Ref, retPtr unsafe.Pointer,
	action uint32,
	handler js.Ref)

//go:wasmimport plat/js/web try_MediaSession_SetActionHandler
//go:noescape
func TryMediaSessionSetActionHandler(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	action uint32,
	handler js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaSession_SetPositionState
//go:noescape
func HasMediaSessionSetPositionState(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSession_SetPositionState
//go:noescape
func MediaSessionSetPositionStateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSession_SetPositionState
//go:noescape
func CallMediaSessionSetPositionState(
	this js.Ref, retPtr unsafe.Pointer,
	state unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaSession_SetPositionState
//go:noescape
func TryMediaSessionSetPositionState(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	state unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaSession_SetPositionState1
//go:noescape
func HasMediaSessionSetPositionState1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSession_SetPositionState1
//go:noescape
func MediaSessionSetPositionState1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSession_SetPositionState1
//go:noescape
func CallMediaSessionSetPositionState1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_MediaSession_SetPositionState1
//go:noescape
func TryMediaSessionSetPositionState1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaSession_SetMicrophoneActive
//go:noescape
func HasMediaSessionSetMicrophoneActive(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSession_SetMicrophoneActive
//go:noescape
func MediaSessionSetMicrophoneActiveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSession_SetMicrophoneActive
//go:noescape
func CallMediaSessionSetMicrophoneActive(
	this js.Ref, retPtr unsafe.Pointer,
	active js.Ref)

//go:wasmimport plat/js/web try_MediaSession_SetMicrophoneActive
//go:noescape
func TryMediaSessionSetMicrophoneActive(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	active js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_MediaSession_SetCameraActive
//go:noescape
func HasMediaSessionSetCameraActive(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_MediaSession_SetCameraActive
//go:noescape
func MediaSessionSetCameraActiveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_MediaSession_SetCameraActive
//go:noescape
func CallMediaSessionSetCameraActive(
	this js.Ref, retPtr unsafe.Pointer,
	active js.Ref)

//go:wasmimport plat/js/web try_MediaSession_SetCameraActive
//go:noescape
func TryMediaSessionSetCameraActive(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	active js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web constof_DevicePostureType
//go:noescape
func ConstOfDevicePostureType(str js.Ref) uint32

//go:wasmimport plat/js/web get_DevicePosture_Type
//go:noescape
func GetDevicePostureType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

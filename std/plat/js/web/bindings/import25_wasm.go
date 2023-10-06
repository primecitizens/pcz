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

//go:wasmimport plat/js/web store_MIDIOptions
//go:noescape
func MIDIOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_MIDIOptions
//go:noescape
func MIDIOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_HIDUnitSystem
//go:noescape
func ConstOfHIDUnitSystem(str js.Ref) uint32

//go:wasmimport plat/js/web store_HIDReportItem
//go:noescape
func HIDReportItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HIDReportItem
//go:noescape
func HIDReportItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_HIDReportInfo
//go:noescape
func HIDReportInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HIDReportInfo
//go:noescape
func HIDReportInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_HIDCollectionInfo
//go:noescape
func HIDCollectionInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HIDCollectionInfo
//go:noescape
func HIDCollectionInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_HIDDevice_Opened
//go:noescape
func GetHIDDeviceOpened(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HIDDevice_VendorId
//go:noescape
func GetHIDDeviceVendorId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HIDDevice_ProductId
//go:noescape
func GetHIDDeviceProductId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HIDDevice_ProductName
//go:noescape
func GetHIDDeviceProductName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HIDDevice_Collections
//go:noescape
func GetHIDDeviceCollections(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HIDDevice_Open
//go:noescape
func HasHIDDeviceOpen(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HIDDevice_Open
//go:noescape
func HIDDeviceOpenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HIDDevice_Open
//go:noescape
func CallHIDDeviceOpen(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HIDDevice_Open
//go:noescape
func TryHIDDeviceOpen(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HIDDevice_Close
//go:noescape
func HasHIDDeviceClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HIDDevice_Close
//go:noescape
func HIDDeviceCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HIDDevice_Close
//go:noescape
func CallHIDDeviceClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HIDDevice_Close
//go:noescape
func TryHIDDeviceClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HIDDevice_Forget
//go:noescape
func HasHIDDeviceForget(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HIDDevice_Forget
//go:noescape
func HIDDeviceForgetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HIDDevice_Forget
//go:noescape
func CallHIDDeviceForget(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HIDDevice_Forget
//go:noescape
func TryHIDDeviceForget(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HIDDevice_SendReport
//go:noescape
func HasHIDDeviceSendReport(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HIDDevice_SendReport
//go:noescape
func HIDDeviceSendReportFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HIDDevice_SendReport
//go:noescape
func CallHIDDeviceSendReport(
	this js.Ref, retPtr unsafe.Pointer,
	reportId uint32,
	data js.Ref)

//go:wasmimport plat/js/web try_HIDDevice_SendReport
//go:noescape
func TryHIDDeviceSendReport(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	reportId uint32,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HIDDevice_SendFeatureReport
//go:noescape
func HasHIDDeviceSendFeatureReport(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HIDDevice_SendFeatureReport
//go:noescape
func HIDDeviceSendFeatureReportFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HIDDevice_SendFeatureReport
//go:noescape
func CallHIDDeviceSendFeatureReport(
	this js.Ref, retPtr unsafe.Pointer,
	reportId uint32,
	data js.Ref)

//go:wasmimport plat/js/web try_HIDDevice_SendFeatureReport
//go:noescape
func TryHIDDeviceSendFeatureReport(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	reportId uint32,
	data js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HIDDevice_ReceiveFeatureReport
//go:noescape
func HasHIDDeviceReceiveFeatureReport(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HIDDevice_ReceiveFeatureReport
//go:noescape
func HIDDeviceReceiveFeatureReportFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HIDDevice_ReceiveFeatureReport
//go:noescape
func CallHIDDeviceReceiveFeatureReport(
	this js.Ref, retPtr unsafe.Pointer,
	reportId uint32)

//go:wasmimport plat/js/web try_HIDDevice_ReceiveFeatureReport
//go:noescape
func TryHIDDeviceReceiveFeatureReport(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	reportId uint32) (ok js.Ref)

//go:wasmimport plat/js/web store_HIDDeviceFilter
//go:noescape
func HIDDeviceFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HIDDeviceFilter
//go:noescape
func HIDDeviceFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_HIDDeviceRequestOptions
//go:noescape
func HIDDeviceRequestOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HIDDeviceRequestOptions
//go:noescape
func HIDDeviceRequestOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_HID_GetDevices
//go:noescape
func HasHIDGetDevices(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HID_GetDevices
//go:noescape
func HIDGetDevicesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HID_GetDevices
//go:noescape
func CallHIDGetDevices(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HID_GetDevices
//go:noescape
func TryHIDGetDevices(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HID_RequestDevice
//go:noescape
func HasHIDRequestDevice(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HID_RequestDevice
//go:noescape
func HIDRequestDeviceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HID_RequestDevice
//go:noescape
func CallHIDRequestDevice(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_HID_RequestDevice
//go:noescape
func TryHIDRequestDevice(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_WindowControlsOverlay_Visible
//go:noescape
func GetWindowControlsOverlayVisible(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_WindowControlsOverlay_GetTitlebarAreaRect
//go:noescape
func HasWindowControlsOverlayGetTitlebarAreaRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_WindowControlsOverlay_GetTitlebarAreaRect
//go:noescape
func WindowControlsOverlayGetTitlebarAreaRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_WindowControlsOverlay_GetTitlebarAreaRect
//go:noescape
func CallWindowControlsOverlayGetTitlebarAreaRect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_WindowControlsOverlay_GetTitlebarAreaRect
//go:noescape
func TryWindowControlsOverlayGetTitlebarAreaRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Credential_Id
//go:noescape
func GetCredentialId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Credential_Type
//go:noescape
func GetCredentialType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Credential_IsConditionalMediationAvailable
//go:noescape
func HasCredentialIsConditionalMediationAvailable(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Credential_IsConditionalMediationAvailable
//go:noescape
func CredentialIsConditionalMediationAvailableFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Credential_IsConditionalMediationAvailable
//go:noescape
func CallCredentialIsConditionalMediationAvailable(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Credential_IsConditionalMediationAvailable
//go:noescape
func TryCredentialIsConditionalMediationAvailable(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_CredentialMediationRequirement
//go:noescape
func ConstOfCredentialMediationRequirement(str js.Ref) uint32

//go:wasmimport plat/js/web constof_OTPCredentialTransportType
//go:noescape
func ConstOfOTPCredentialTransportType(str js.Ref) uint32

//go:wasmimport plat/js/web store_OTPCredentialRequestOptions
//go:noescape
func OTPCredentialRequestOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_OTPCredentialRequestOptions
//go:noescape
func OTPCredentialRequestOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PublicKeyCredentialDescriptor
//go:noescape
func PublicKeyCredentialDescriptorJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PublicKeyCredentialDescriptor
//go:noescape
func PublicKeyCredentialDescriptorJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PublicKeyCredentialRequestOptions
//go:noescape
func PublicKeyCredentialRequestOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PublicKeyCredentialRequestOptions
//go:noescape
func PublicKeyCredentialRequestOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_IdentityProviderConfig
//go:noescape
func IdentityProviderConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdentityProviderConfig
//go:noescape
func IdentityProviderConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_IdentityCredentialRequestOptionsContext
//go:noescape
func ConstOfIdentityCredentialRequestOptionsContext(str js.Ref) uint32

//go:wasmimport plat/js/web store_IdentityCredentialRequestOptions
//go:noescape
func IdentityCredentialRequestOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IdentityCredentialRequestOptions
//go:noescape
func IdentityCredentialRequestOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_FederatedCredentialRequestOptions
//go:noescape
func FederatedCredentialRequestOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FederatedCredentialRequestOptions
//go:noescape
func FederatedCredentialRequestOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CredentialRequestOptions
//go:noescape
func CredentialRequestOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CredentialRequestOptions
//go:noescape
func CredentialRequestOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PublicKeyCredentialRpEntity
//go:noescape
func PublicKeyCredentialRpEntityJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PublicKeyCredentialRpEntity
//go:noescape
func PublicKeyCredentialRpEntityJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PublicKeyCredentialUserEntity
//go:noescape
func PublicKeyCredentialUserEntityJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PublicKeyCredentialUserEntity
//go:noescape
func PublicKeyCredentialUserEntityJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PublicKeyCredentialParameters
//go:noescape
func PublicKeyCredentialParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PublicKeyCredentialParameters
//go:noescape
func PublicKeyCredentialParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PublicKeyCredentialCreationOptions
//go:noescape
func PublicKeyCredentialCreationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PublicKeyCredentialCreationOptions
//go:noescape
func PublicKeyCredentialCreationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_PasswordCredentialData
//go:noescape
func PasswordCredentialDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PasswordCredentialData
//go:noescape
func PasswordCredentialDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_FederatedCredentialInit
//go:noescape
func FederatedCredentialInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FederatedCredentialInit
//go:noescape
func FederatedCredentialInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_CredentialCreationOptions
//go:noescape
func CredentialCreationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CredentialCreationOptions
//go:noescape
func CredentialCreationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_CredentialsContainer_Get
//go:noescape
func HasCredentialsContainerGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CredentialsContainer_Get
//go:noescape
func CredentialsContainerGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CredentialsContainer_Get
//go:noescape
func CallCredentialsContainerGet(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CredentialsContainer_Get
//go:noescape
func TryCredentialsContainerGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CredentialsContainer_Get1
//go:noescape
func HasCredentialsContainerGet1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CredentialsContainer_Get1
//go:noescape
func CredentialsContainerGet1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CredentialsContainer_Get1
//go:noescape
func CallCredentialsContainerGet1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CredentialsContainer_Get1
//go:noescape
func TryCredentialsContainerGet1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CredentialsContainer_Store
//go:noescape
func HasCredentialsContainerStore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CredentialsContainer_Store
//go:noescape
func CredentialsContainerStoreFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CredentialsContainer_Store
//go:noescape
func CallCredentialsContainerStore(
	this js.Ref, retPtr unsafe.Pointer,
	credential js.Ref)

//go:wasmimport plat/js/web try_CredentialsContainer_Store
//go:noescape
func TryCredentialsContainerStore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	credential js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CredentialsContainer_Create
//go:noescape
func HasCredentialsContainerCreate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CredentialsContainer_Create
//go:noescape
func CredentialsContainerCreateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CredentialsContainer_Create
//go:noescape
func CallCredentialsContainerCreate(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_CredentialsContainer_Create
//go:noescape
func TryCredentialsContainerCreate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CredentialsContainer_Create1
//go:noescape
func HasCredentialsContainerCreate1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CredentialsContainer_Create1
//go:noescape
func CredentialsContainerCreate1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CredentialsContainer_Create1
//go:noescape
func CallCredentialsContainerCreate1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CredentialsContainer_Create1
//go:noescape
func TryCredentialsContainerCreate1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CredentialsContainer_PreventSilentAccess
//go:noescape
func HasCredentialsContainerPreventSilentAccess(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CredentialsContainer_PreventSilentAccess
//go:noescape
func CredentialsContainerPreventSilentAccessFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CredentialsContainer_PreventSilentAccess
//go:noescape
func CallCredentialsContainerPreventSilentAccess(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CredentialsContainer_PreventSilentAccess
//go:noescape
func TryCredentialsContainerPreventSilentAccess(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationCoordinates_Accuracy
//go:noescape
func GetGeolocationCoordinatesAccuracy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationCoordinates_Latitude
//go:noescape
func GetGeolocationCoordinatesLatitude(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationCoordinates_Longitude
//go:noescape
func GetGeolocationCoordinatesLongitude(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationCoordinates_Altitude
//go:noescape
func GetGeolocationCoordinatesAltitude(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationCoordinates_AltitudeAccuracy
//go:noescape
func GetGeolocationCoordinatesAltitudeAccuracy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationCoordinates_Heading
//go:noescape
func GetGeolocationCoordinatesHeading(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationCoordinates_Speed
//go:noescape
func GetGeolocationCoordinatesSpeed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationPosition_Coords
//go:noescape
func GetGeolocationPositionCoords(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationPosition_Timestamp
//go:noescape
func GetGeolocationPositionTimestamp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationPositionError_Code
//go:noescape
func GetGeolocationPositionErrorCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_GeolocationPositionError_Message
//go:noescape
func GetGeolocationPositionErrorMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_PositionOptions
//go:noescape
func PositionOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PositionOptions
//go:noescape
func PositionOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web has_Geolocation_GetCurrentPosition
//go:noescape
func HasGeolocationGetCurrentPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Geolocation_GetCurrentPosition
//go:noescape
func GeolocationGetCurrentPositionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Geolocation_GetCurrentPosition
//go:noescape
func CallGeolocationGetCurrentPosition(
	this js.Ref, retPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Geolocation_GetCurrentPosition
//go:noescape
func TryGeolocationGetCurrentPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Geolocation_GetCurrentPosition1
//go:noescape
func HasGeolocationGetCurrentPosition1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Geolocation_GetCurrentPosition1
//go:noescape
func GeolocationGetCurrentPosition1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Geolocation_GetCurrentPosition1
//go:noescape
func CallGeolocationGetCurrentPosition1(
	this js.Ref, retPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref)

//go:wasmimport plat/js/web try_Geolocation_GetCurrentPosition1
//go:noescape
func TryGeolocationGetCurrentPosition1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Geolocation_GetCurrentPosition2
//go:noescape
func HasGeolocationGetCurrentPosition2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Geolocation_GetCurrentPosition2
//go:noescape
func GeolocationGetCurrentPosition2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Geolocation_GetCurrentPosition2
//go:noescape
func CallGeolocationGetCurrentPosition2(
	this js.Ref, retPtr unsafe.Pointer,
	successCallback js.Ref)

//go:wasmimport plat/js/web try_Geolocation_GetCurrentPosition2
//go:noescape
func TryGeolocationGetCurrentPosition2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	successCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Geolocation_WatchPosition
//go:noescape
func HasGeolocationWatchPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Geolocation_WatchPosition
//go:noescape
func GeolocationWatchPositionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Geolocation_WatchPosition
//go:noescape
func CallGeolocationWatchPosition(
	this js.Ref, retPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_Geolocation_WatchPosition
//go:noescape
func TryGeolocationWatchPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Geolocation_WatchPosition1
//go:noescape
func HasGeolocationWatchPosition1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Geolocation_WatchPosition1
//go:noescape
func GeolocationWatchPosition1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Geolocation_WatchPosition1
//go:noescape
func CallGeolocationWatchPosition1(
	this js.Ref, retPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref)

//go:wasmimport plat/js/web try_Geolocation_WatchPosition1
//go:noescape
func TryGeolocationWatchPosition1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	successCallback js.Ref,
	errorCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Geolocation_WatchPosition2
//go:noescape
func HasGeolocationWatchPosition2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Geolocation_WatchPosition2
//go:noescape
func GeolocationWatchPosition2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Geolocation_WatchPosition2
//go:noescape
func CallGeolocationWatchPosition2(
	this js.Ref, retPtr unsafe.Pointer,
	successCallback js.Ref)

//go:wasmimport plat/js/web try_Geolocation_WatchPosition2
//go:noescape
func TryGeolocationWatchPosition2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	successCallback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Geolocation_ClearWatch
//go:noescape
func HasGeolocationClearWatch(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Geolocation_ClearWatch
//go:noescape
func GeolocationClearWatchFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Geolocation_ClearWatch
//go:noescape
func CallGeolocationClearWatch(
	this js.Ref, retPtr unsafe.Pointer,
	watchId int32)

//go:wasmimport plat/js/web try_Geolocation_ClearWatch
//go:noescape
func TryGeolocationClearWatch(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	watchId int32) (ok js.Ref)

//go:wasmimport plat/js/web constof_USBTransferStatus
//go:noescape
func ConstOfUSBTransferStatus(str js.Ref) uint32

//go:wasmimport plat/js/web new_USBInTransferResult_USBInTransferResult
//go:noescape
func NewUSBInTransferResultByUSBInTransferResult(
	status uint32,
	data js.Ref) js.Ref

//go:wasmimport plat/js/web new_USBInTransferResult_USBInTransferResult1
//go:noescape
func NewUSBInTransferResultByUSBInTransferResult1(
	status uint32) js.Ref

//go:wasmimport plat/js/web get_USBInTransferResult_Data
//go:noescape
func GetUSBInTransferResultData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_USBInTransferResult_Status
//go:noescape
func GetUSBInTransferResultStatus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_USBRequestType
//go:noescape
func ConstOfUSBRequestType(str js.Ref) uint32

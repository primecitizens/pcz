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

//go:wasmimport plat/js/webext/networkingprivate store_APNProperties
//go:noescape
func APNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_APNProperties
//go:noescape
func APNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate constof_ActivationStateType
//go:noescape
func ConstOfActivationStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networkingprivate constof_CaptivePortalStatus
//go:noescape
func ConstOfCaptivePortalStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/networkingprivate store_FoundNetworkProperties
//go:noescape
func FoundNetworkPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_FoundNetworkProperties
//go:noescape
func FoundNetworkPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_CellularProviderProperties
//go:noescape
func CellularProviderPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_CellularProviderProperties
//go:noescape
func CellularProviderPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_PaymentPortal
//go:noescape
func PaymentPortalJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_PaymentPortal
//go:noescape
func PaymentPortalJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_SIMLockStatus
//go:noescape
func SIMLockStatusJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_SIMLockStatus
//go:noescape
func SIMLockStatusJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_CellularProperties
//go:noescape
func CellularPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_CellularProperties
//go:noescape
func CellularPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_CellularSimState
//go:noescape
func CellularSimStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_CellularSimState
//go:noescape
func CellularSimStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_CellularStateProperties
//go:noescape
func CellularStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_CellularStateProperties
//go:noescape
func CellularStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_Certificate
//go:noescape
func CertificateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_Certificate
//go:noescape
func CertificateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_CertificateLists
//go:noescape
func CertificateListsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_CertificateLists
//go:noescape
func CertificateListsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_IssuerSubjectPattern
//go:noescape
func IssuerSubjectPatternJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_IssuerSubjectPattern
//go:noescape
func IssuerSubjectPatternJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_CertificatePattern
//go:noescape
func CertificatePatternJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_CertificatePattern
//go:noescape
func CertificatePatternJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate constof_ConnectionStateType
//go:noescape
func ConstOfConnectionStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networkingprivate constof_DeviceStateType
//go:noescape
func ConstOfDeviceStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networkingprivate constof_NetworkType
//go:noescape
func ConstOfNetworkType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networkingprivate store_DeviceStateProperties
//go:noescape
func DeviceStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_DeviceStateProperties
//go:noescape
func DeviceStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_EAPProperties
//go:noescape
func EAPPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_EAPProperties
//go:noescape
func EAPPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_EAPStateProperties
//go:noescape
func EAPStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_EAPStateProperties
//go:noescape
func EAPStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_EthernetProperties
//go:noescape
func EthernetPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_EthernetProperties
//go:noescape
func EthernetPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_EthernetStateProperties
//go:noescape
func EthernetStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_EthernetStateProperties
//go:noescape
func EthernetStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_GlobalPolicy
//go:noescape
func GlobalPolicyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_GlobalPolicy
//go:noescape
func GlobalPolicyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedBoolean
//go:noescape
func ManagedBooleanJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedBoolean
//go:noescape
func ManagedBooleanJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedDOMString
//go:noescape
func ManagedDOMStringJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedDOMString
//go:noescape
func ManagedDOMStringJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedAPNProperties
//go:noescape
func ManagedAPNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedAPNProperties
//go:noescape
func ManagedAPNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedAPNList
//go:noescape
func ManagedAPNListJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedAPNList
//go:noescape
func ManagedAPNListJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedCellularProperties
//go:noescape
func ManagedCellularPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedCellularProperties
//go:noescape
func ManagedCellularPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedDOMStringList
//go:noescape
func ManagedDOMStringListJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedDOMStringList
//go:noescape
func ManagedDOMStringListJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedIssuerSubjectPattern
//go:noescape
func ManagedIssuerSubjectPatternJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedIssuerSubjectPattern
//go:noescape
func ManagedIssuerSubjectPatternJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedCertificatePattern
//go:noescape
func ManagedCertificatePatternJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedCertificatePattern
//go:noescape
func ManagedCertificatePatternJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedEAPProperties
//go:noescape
func ManagedEAPPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedEAPProperties
//go:noescape
func ManagedEAPPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedEthernetProperties
//go:noescape
func ManagedEthernetPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedEthernetProperties
//go:noescape
func ManagedEthernetPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate constof_IPConfigType
//go:noescape
func ConstOfIPConfigType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networkingprivate store_ManagedIPConfigType
//go:noescape
func ManagedIPConfigTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedIPConfigType
//go:noescape
func ManagedIPConfigTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_IPConfigProperties
//go:noescape
func IPConfigPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_IPConfigProperties
//go:noescape
func IPConfigPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedLong
//go:noescape
func ManagedLongJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedLong
//go:noescape
func ManagedLongJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate constof_ProxySettingsType
//go:noescape
func ConstOfProxySettingsType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networkingprivate store_ManagedProxySettingsType
//go:noescape
func ManagedProxySettingsTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedProxySettingsType
//go:noescape
func ManagedProxySettingsTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedProxyLocation
//go:noescape
func ManagedProxyLocationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedProxyLocation
//go:noescape
func ManagedProxyLocationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedManualProxySettings
//go:noescape
func ManagedManualProxySettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedManualProxySettings
//go:noescape
func ManagedManualProxySettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedProxySettings
//go:noescape
func ManagedProxySettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedProxySettings
//go:noescape
func ManagedProxySettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedIPConfigProperties
//go:noescape
func ManagedIPConfigPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedIPConfigProperties
//go:noescape
func ManagedIPConfigPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_TetherProperties
//go:noescape
func TetherPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_TetherProperties
//go:noescape
func TetherPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedXAUTHProperties
//go:noescape
func ManagedXAUTHPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedXAUTHProperties
//go:noescape
func ManagedXAUTHPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedIPSecProperties
//go:noescape
func ManagedIPSecPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedIPSecProperties
//go:noescape
func ManagedIPSecPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedL2TPProperties
//go:noescape
func ManagedL2TPPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedL2TPProperties
//go:noescape
func ManagedL2TPPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedVerifyX509
//go:noescape
func ManagedVerifyX509JSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedVerifyX509
//go:noescape
func ManagedVerifyX509JSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedOpenVPNProperties
//go:noescape
func ManagedOpenVPNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedOpenVPNProperties
//go:noescape
func ManagedOpenVPNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedThirdPartyVPNProperties
//go:noescape
func ManagedThirdPartyVPNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedThirdPartyVPNProperties
//go:noescape
func ManagedThirdPartyVPNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedVPNProperties
//go:noescape
func ManagedVPNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedVPNProperties
//go:noescape
func ManagedVPNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedWiFiProperties
//go:noescape
func ManagedWiFiPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedWiFiProperties
//go:noescape
func ManagedWiFiPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManagedProperties
//go:noescape
func ManagedPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManagedProperties
//go:noescape
func ManagedPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_XAUTHProperties
//go:noescape
func XAUTHPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_XAUTHProperties
//go:noescape
func XAUTHPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_IPSecProperties
//go:noescape
func IPSecPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_IPSecProperties
//go:noescape
func IPSecPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ThirdPartyVPNProperties
//go:noescape
func ThirdPartyVPNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ThirdPartyVPNProperties
//go:noescape
func ThirdPartyVPNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_VPNStateProperties
//go:noescape
func VPNStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_VPNStateProperties
//go:noescape
func VPNStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_WiFiStateProperties
//go:noescape
func WiFiStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_WiFiStateProperties
//go:noescape
func WiFiStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_NetworkStateProperties
//go:noescape
func NetworkStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_NetworkStateProperties
//go:noescape
func NetworkStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ProxyLocation
//go:noescape
func ProxyLocationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ProxyLocation
//go:noescape
func ProxyLocationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ManualProxySettings
//go:noescape
func ManualProxySettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ManualProxySettings
//go:noescape
func ManualProxySettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_ProxySettings
//go:noescape
func ProxySettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_ProxySettings
//go:noescape
func ProxySettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_L2TPProperties
//go:noescape
func L2TPPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_L2TPProperties
//go:noescape
func L2TPPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_VerifyX509
//go:noescape
func VerifyX509JSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_VerifyX509
//go:noescape
func VerifyX509JSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_OpenVPNProperties
//go:noescape
func OpenVPNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_OpenVPNProperties
//go:noescape
func OpenVPNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_VPNProperties
//go:noescape
func VPNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_VPNProperties
//go:noescape
func VPNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_WiFiProperties
//go:noescape
func WiFiPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_WiFiProperties
//go:noescape
func WiFiPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_NetworkProperties
//go:noescape
func NetworkPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_NetworkProperties
//go:noescape
func NetworkPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_NetworkConfigProperties
//go:noescape
func NetworkConfigPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_NetworkConfigProperties
//go:noescape
func NetworkConfigPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate store_NetworkFilter
//go:noescape
func NetworkFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networkingprivate load_NetworkFilter
//go:noescape
func NetworkFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networkingprivate has_CreateNetwork
//go:noescape
func HasFuncCreateNetwork() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_CreateNetwork
//go:noescape
func FuncCreateNetwork(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_CreateNetwork
//go:noescape
func CallCreateNetwork(
	retPtr unsafe.Pointer,
	shared js.Ref,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate try_CreateNetwork
//go:noescape
func TryCreateNetwork(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	shared js.Ref,
	properties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_DisableNetworkType
//go:noescape
func HasFuncDisableNetworkType() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_DisableNetworkType
//go:noescape
func FuncDisableNetworkType(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_DisableNetworkType
//go:noescape
func CallDisableNetworkType(
	retPtr unsafe.Pointer,
	networkType uint32)

//go:wasmimport plat/js/webext/networkingprivate try_DisableNetworkType
//go:noescape
func TryDisableNetworkType(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_EnableNetworkType
//go:noescape
func HasFuncEnableNetworkType() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_EnableNetworkType
//go:noescape
func FuncEnableNetworkType(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_EnableNetworkType
//go:noescape
func CallEnableNetworkType(
	retPtr unsafe.Pointer,
	networkType uint32)

//go:wasmimport plat/js/webext/networkingprivate try_EnableNetworkType
//go:noescape
func TryEnableNetworkType(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_ForgetNetwork
//go:noescape
func HasFuncForgetNetwork() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_ForgetNetwork
//go:noescape
func FuncForgetNetwork(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_ForgetNetwork
//go:noescape
func CallForgetNetwork(
	retPtr unsafe.Pointer,
	networkGuid js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_ForgetNetwork
//go:noescape
func TryForgetNetwork(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_GetCaptivePortalStatus
//go:noescape
func HasFuncGetCaptivePortalStatus() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_GetCaptivePortalStatus
//go:noescape
func FuncGetCaptivePortalStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_GetCaptivePortalStatus
//go:noescape
func CallGetCaptivePortalStatus(
	retPtr unsafe.Pointer,
	networkGuid js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_GetCaptivePortalStatus
//go:noescape
func TryGetCaptivePortalStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_GetCertificateLists
//go:noescape
func HasFuncGetCertificateLists() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_GetCertificateLists
//go:noescape
func FuncGetCertificateLists(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_GetCertificateLists
//go:noescape
func CallGetCertificateLists(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate try_GetCertificateLists
//go:noescape
func TryGetCertificateLists(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_GetDeviceStates
//go:noescape
func HasFuncGetDeviceStates() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_GetDeviceStates
//go:noescape
func FuncGetDeviceStates(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_GetDeviceStates
//go:noescape
func CallGetDeviceStates(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate try_GetDeviceStates
//go:noescape
func TryGetDeviceStates(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_GetEnabledNetworkTypes
//go:noescape
func HasFuncGetEnabledNetworkTypes() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_GetEnabledNetworkTypes
//go:noescape
func FuncGetEnabledNetworkTypes(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_GetEnabledNetworkTypes
//go:noescape
func CallGetEnabledNetworkTypes(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_GetEnabledNetworkTypes
//go:noescape
func TryGetEnabledNetworkTypes(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_GetGlobalPolicy
//go:noescape
func HasFuncGetGlobalPolicy() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_GetGlobalPolicy
//go:noescape
func FuncGetGlobalPolicy(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_GetGlobalPolicy
//go:noescape
func CallGetGlobalPolicy(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate try_GetGlobalPolicy
//go:noescape
func TryGetGlobalPolicy(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_GetManagedProperties
//go:noescape
func HasFuncGetManagedProperties() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_GetManagedProperties
//go:noescape
func FuncGetManagedProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_GetManagedProperties
//go:noescape
func CallGetManagedProperties(
	retPtr unsafe.Pointer,
	networkGuid js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_GetManagedProperties
//go:noescape
func TryGetManagedProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_GetNetworks
//go:noescape
func HasFuncGetNetworks() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_GetNetworks
//go:noescape
func FuncGetNetworks(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_GetNetworks
//go:noescape
func CallGetNetworks(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate try_GetNetworks
//go:noescape
func TryGetNetworks(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_GetProperties
//go:noescape
func HasFuncGetProperties() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_GetProperties
//go:noescape
func FuncGetProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_GetProperties
//go:noescape
func CallGetProperties(
	retPtr unsafe.Pointer,
	networkGuid js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_GetProperties
//go:noescape
func TryGetProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_GetState
//go:noescape
func HasFuncGetState() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_GetState
//go:noescape
func FuncGetState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_GetState
//go:noescape
func CallGetState(
	retPtr unsafe.Pointer,
	networkGuid js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_GetState
//go:noescape
func TryGetState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_GetVisibleNetworks
//go:noescape
func HasFuncGetVisibleNetworks() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_GetVisibleNetworks
//go:noescape
func FuncGetVisibleNetworks(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_GetVisibleNetworks
//go:noescape
func CallGetVisibleNetworks(
	retPtr unsafe.Pointer,
	networkType uint32,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_GetVisibleNetworks
//go:noescape
func TryGetVisibleNetworks(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkType uint32,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_OnCertificateListsChanged
//go:noescape
func HasFuncOnCertificateListsChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_OnCertificateListsChanged
//go:noescape
func FuncOnCertificateListsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_OnCertificateListsChanged
//go:noescape
func CallOnCertificateListsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_OnCertificateListsChanged
//go:noescape
func TryOnCertificateListsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_OffCertificateListsChanged
//go:noescape
func HasFuncOffCertificateListsChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_OffCertificateListsChanged
//go:noescape
func FuncOffCertificateListsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_OffCertificateListsChanged
//go:noescape
func CallOffCertificateListsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_OffCertificateListsChanged
//go:noescape
func TryOffCertificateListsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_HasOnCertificateListsChanged
//go:noescape
func HasFuncHasOnCertificateListsChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_HasOnCertificateListsChanged
//go:noescape
func FuncHasOnCertificateListsChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_HasOnCertificateListsChanged
//go:noescape
func CallHasOnCertificateListsChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_HasOnCertificateListsChanged
//go:noescape
func TryHasOnCertificateListsChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_OnDeviceStateListChanged
//go:noescape
func HasFuncOnDeviceStateListChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_OnDeviceStateListChanged
//go:noescape
func FuncOnDeviceStateListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_OnDeviceStateListChanged
//go:noescape
func CallOnDeviceStateListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_OnDeviceStateListChanged
//go:noescape
func TryOnDeviceStateListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_OffDeviceStateListChanged
//go:noescape
func HasFuncOffDeviceStateListChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_OffDeviceStateListChanged
//go:noescape
func FuncOffDeviceStateListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_OffDeviceStateListChanged
//go:noescape
func CallOffDeviceStateListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_OffDeviceStateListChanged
//go:noescape
func TryOffDeviceStateListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_HasOnDeviceStateListChanged
//go:noescape
func HasFuncHasOnDeviceStateListChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_HasOnDeviceStateListChanged
//go:noescape
func FuncHasOnDeviceStateListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_HasOnDeviceStateListChanged
//go:noescape
func CallHasOnDeviceStateListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_HasOnDeviceStateListChanged
//go:noescape
func TryHasOnDeviceStateListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_OnNetworkListChanged
//go:noescape
func HasFuncOnNetworkListChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_OnNetworkListChanged
//go:noescape
func FuncOnNetworkListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_OnNetworkListChanged
//go:noescape
func CallOnNetworkListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_OnNetworkListChanged
//go:noescape
func TryOnNetworkListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_OffNetworkListChanged
//go:noescape
func HasFuncOffNetworkListChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_OffNetworkListChanged
//go:noescape
func FuncOffNetworkListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_OffNetworkListChanged
//go:noescape
func CallOffNetworkListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_OffNetworkListChanged
//go:noescape
func TryOffNetworkListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_HasOnNetworkListChanged
//go:noescape
func HasFuncHasOnNetworkListChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_HasOnNetworkListChanged
//go:noescape
func FuncHasOnNetworkListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_HasOnNetworkListChanged
//go:noescape
func CallHasOnNetworkListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_HasOnNetworkListChanged
//go:noescape
func TryHasOnNetworkListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_OnNetworksChanged
//go:noescape
func HasFuncOnNetworksChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_OnNetworksChanged
//go:noescape
func FuncOnNetworksChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_OnNetworksChanged
//go:noescape
func CallOnNetworksChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_OnNetworksChanged
//go:noescape
func TryOnNetworksChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_OffNetworksChanged
//go:noescape
func HasFuncOffNetworksChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_OffNetworksChanged
//go:noescape
func FuncOffNetworksChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_OffNetworksChanged
//go:noescape
func CallOffNetworksChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_OffNetworksChanged
//go:noescape
func TryOffNetworksChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_HasOnNetworksChanged
//go:noescape
func HasFuncHasOnNetworksChanged() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_HasOnNetworksChanged
//go:noescape
func FuncHasOnNetworksChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_HasOnNetworksChanged
//go:noescape
func CallHasOnNetworksChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_HasOnNetworksChanged
//go:noescape
func TryHasOnNetworksChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_OnPortalDetectionCompleted
//go:noescape
func HasFuncOnPortalDetectionCompleted() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_OnPortalDetectionCompleted
//go:noescape
func FuncOnPortalDetectionCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_OnPortalDetectionCompleted
//go:noescape
func CallOnPortalDetectionCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_OnPortalDetectionCompleted
//go:noescape
func TryOnPortalDetectionCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_OffPortalDetectionCompleted
//go:noescape
func HasFuncOffPortalDetectionCompleted() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_OffPortalDetectionCompleted
//go:noescape
func FuncOffPortalDetectionCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_OffPortalDetectionCompleted
//go:noescape
func CallOffPortalDetectionCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_OffPortalDetectionCompleted
//go:noescape
func TryOffPortalDetectionCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_HasOnPortalDetectionCompleted
//go:noescape
func HasFuncHasOnPortalDetectionCompleted() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_HasOnPortalDetectionCompleted
//go:noescape
func FuncHasOnPortalDetectionCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_HasOnPortalDetectionCompleted
//go:noescape
func CallHasOnPortalDetectionCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_HasOnPortalDetectionCompleted
//go:noescape
func TryHasOnPortalDetectionCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_RequestNetworkScan
//go:noescape
func HasFuncRequestNetworkScan() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_RequestNetworkScan
//go:noescape
func FuncRequestNetworkScan(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_RequestNetworkScan
//go:noescape
func CallRequestNetworkScan(
	retPtr unsafe.Pointer,
	networkType uint32)

//go:wasmimport plat/js/webext/networkingprivate try_RequestNetworkScan
//go:noescape
func TryRequestNetworkScan(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_SelectCellularMobileNetwork
//go:noescape
func HasFuncSelectCellularMobileNetwork() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_SelectCellularMobileNetwork
//go:noescape
func FuncSelectCellularMobileNetwork(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_SelectCellularMobileNetwork
//go:noescape
func CallSelectCellularMobileNetwork(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	networkId js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_SelectCellularMobileNetwork
//go:noescape
func TrySelectCellularMobileNetwork(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	networkId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_SetCellularSimState
//go:noescape
func HasFuncSetCellularSimState() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_SetCellularSimState
//go:noescape
func FuncSetCellularSimState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_SetCellularSimState
//go:noescape
func CallSetCellularSimState(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	simState unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate try_SetCellularSimState
//go:noescape
func TrySetCellularSimState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	simState unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_SetProperties
//go:noescape
func HasFuncSetProperties() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_SetProperties
//go:noescape
func FuncSetProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_SetProperties
//go:noescape
func CallSetProperties(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	properties unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate try_SetProperties
//go:noescape
func TrySetProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	properties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_StartActivate
//go:noescape
func HasFuncStartActivate() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_StartActivate
//go:noescape
func FuncStartActivate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_StartActivate
//go:noescape
func CallStartActivate(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	carrier js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_StartActivate
//go:noescape
func TryStartActivate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	carrier js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_StartConnect
//go:noescape
func HasFuncStartConnect() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_StartConnect
//go:noescape
func FuncStartConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_StartConnect
//go:noescape
func CallStartConnect(
	retPtr unsafe.Pointer,
	networkGuid js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_StartConnect
//go:noescape
func TryStartConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_StartDisconnect
//go:noescape
func HasFuncStartDisconnect() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_StartDisconnect
//go:noescape
func FuncStartDisconnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_StartDisconnect
//go:noescape
func CallStartDisconnect(
	retPtr unsafe.Pointer,
	networkGuid js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_StartDisconnect
//go:noescape
func TryStartDisconnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networkingprivate has_UnlockCellularSim
//go:noescape
func HasFuncUnlockCellularSim() js.Ref

//go:wasmimport plat/js/webext/networkingprivate func_UnlockCellularSim
//go:noescape
func FuncUnlockCellularSim(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networkingprivate call_UnlockCellularSim
//go:noescape
func CallUnlockCellularSim(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	pin js.Ref,
	puk js.Ref)

//go:wasmimport plat/js/webext/networkingprivate try_UnlockCellularSim
//go:noescape
func TryUnlockCellularSim(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	pin js.Ref,
	puk js.Ref) (ok js.Ref)

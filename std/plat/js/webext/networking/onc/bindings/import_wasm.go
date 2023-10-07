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

//go:wasmimport plat/js/webext/networking/onc constof_ActivationStateType
//go:noescape
func ConstOfActivationStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networking/onc constof_CaptivePortalStatus
//go:noescape
func ConstOfCaptivePortalStatus(str js.Ref) uint32

//go:wasmimport plat/js/webext/networking/onc store_FoundNetworkProperties
//go:noescape
func FoundNetworkPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_FoundNetworkProperties
//go:noescape
func FoundNetworkPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_CellularProviderProperties
//go:noescape
func CellularProviderPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_CellularProviderProperties
//go:noescape
func CellularProviderPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_PaymentPortal
//go:noescape
func PaymentPortalJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_PaymentPortal
//go:noescape
func PaymentPortalJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_SIMLockStatus
//go:noescape
func SIMLockStatusJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_SIMLockStatus
//go:noescape
func SIMLockStatusJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_CellularProperties
//go:noescape
func CellularPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_CellularProperties
//go:noescape
func CellularPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_CellularStateProperties
//go:noescape
func CellularStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_CellularStateProperties
//go:noescape
func CellularStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_IssuerSubjectPattern
//go:noescape
func IssuerSubjectPatternJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_IssuerSubjectPattern
//go:noescape
func IssuerSubjectPatternJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_CertificatePattern
//go:noescape
func CertificatePatternJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_CertificatePattern
//go:noescape
func CertificatePatternJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc constof_ClientCertificateType
//go:noescape
func ConstOfClientCertificateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networking/onc constof_ConnectionStateType
//go:noescape
func ConstOfConnectionStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networking/onc constof_DeviceStateType
//go:noescape
func ConstOfDeviceStateType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networking/onc constof_NetworkType
//go:noescape
func ConstOfNetworkType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networking/onc store_DeviceStateProperties
//go:noescape
func DeviceStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_DeviceStateProperties
//go:noescape
func DeviceStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedDOMString
//go:noescape
func ManagedDOMStringJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedDOMString
//go:noescape
func ManagedDOMStringJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_EAPProperties
//go:noescape
func EAPPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_EAPProperties
//go:noescape
func EAPPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_EthernetProperties
//go:noescape
func EthernetPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_EthernetProperties
//go:noescape
func EthernetPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_EthernetStateProperties
//go:noescape
func EthernetStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_EthernetStateProperties
//go:noescape
func EthernetStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_GlobalPolicy
//go:noescape
func GlobalPolicyJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_GlobalPolicy
//go:noescape
func GlobalPolicyJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedBoolean
//go:noescape
func ManagedBooleanJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedBoolean
//go:noescape
func ManagedBooleanJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedCellularProperties
//go:noescape
func ManagedCellularPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedCellularProperties
//go:noescape
func ManagedCellularPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedEthernetProperties
//go:noescape
func ManagedEthernetPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedEthernetProperties
//go:noescape
func ManagedEthernetPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc constof_IPConfigType
//go:noescape
func ConstOfIPConfigType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networking/onc store_ManagedIPConfigType
//go:noescape
func ManagedIPConfigTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedIPConfigType
//go:noescape
func ManagedIPConfigTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_IPConfigProperties
//go:noescape
func IPConfigPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_IPConfigProperties
//go:noescape
func IPConfigPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedLong
//go:noescape
func ManagedLongJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedLong
//go:noescape
func ManagedLongJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc constof_ProxySettingsType
//go:noescape
func ConstOfProxySettingsType(str js.Ref) uint32

//go:wasmimport plat/js/webext/networking/onc store_ManagedProxySettingsType
//go:noescape
func ManagedProxySettingsTypeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedProxySettingsType
//go:noescape
func ManagedProxySettingsTypeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedProxyLocation
//go:noescape
func ManagedProxyLocationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedProxyLocation
//go:noescape
func ManagedProxyLocationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedManualProxySettings
//go:noescape
func ManagedManualProxySettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedManualProxySettings
//go:noescape
func ManagedManualProxySettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedDOMStringList
//go:noescape
func ManagedDOMStringListJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedDOMStringList
//go:noescape
func ManagedDOMStringListJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedProxySettings
//go:noescape
func ManagedProxySettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedProxySettings
//go:noescape
func ManagedProxySettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedIPConfigProperties
//go:noescape
func ManagedIPConfigPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedIPConfigProperties
//go:noescape
func ManagedIPConfigPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedVPNProperties
//go:noescape
func ManagedVPNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedVPNProperties
//go:noescape
func ManagedVPNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedWiFiProperties
//go:noescape
func ManagedWiFiPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedWiFiProperties
//go:noescape
func ManagedWiFiPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedProperties
//go:noescape
func ManagedPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedProperties
//go:noescape
func ManagedPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_VPNStateProperties
//go:noescape
func VPNStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_VPNStateProperties
//go:noescape
func VPNStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_WiFiStateProperties
//go:noescape
func WiFiStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_WiFiStateProperties
//go:noescape
func WiFiStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_NetworkStateProperties
//go:noescape
func NetworkStatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_NetworkStateProperties
//go:noescape
func NetworkStatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ProxyLocation
//go:noescape
func ProxyLocationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ProxyLocation
//go:noescape
func ProxyLocationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManualProxySettings
//go:noescape
func ManualProxySettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManualProxySettings
//go:noescape
func ManualProxySettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ProxySettings
//go:noescape
func ProxySettingsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ProxySettings
//go:noescape
func ProxySettingsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_VPNProperties
//go:noescape
func VPNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_VPNProperties
//go:noescape
func VPNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_WiFiProperties
//go:noescape
func WiFiPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_WiFiProperties
//go:noescape
func WiFiPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_NetworkProperties
//go:noescape
func NetworkPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_NetworkProperties
//go:noescape
func NetworkPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ManagedThirdPartyVPNProperties
//go:noescape
func ManagedThirdPartyVPNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ManagedThirdPartyVPNProperties
//go:noescape
func ManagedThirdPartyVPNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_WiMAXProperties
//go:noescape
func WiMAXPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_WiMAXProperties
//go:noescape
func WiMAXPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_NetworkConfigProperties
//go:noescape
func NetworkConfigPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_NetworkConfigProperties
//go:noescape
func NetworkConfigPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_NetworkFilter
//go:noescape
func NetworkFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_NetworkFilter
//go:noescape
func NetworkFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc store_ThirdPartyVPNProperties
//go:noescape
func ThirdPartyVPNPropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/networking/onc load_ThirdPartyVPNProperties
//go:noescape
func ThirdPartyVPNPropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/networking/onc has_CreateNetwork
//go:noescape
func HasFuncCreateNetwork() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_CreateNetwork
//go:noescape
func FuncCreateNetwork(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_CreateNetwork
//go:noescape
func CallCreateNetwork(
	retPtr unsafe.Pointer,
	shared js.Ref,
	properties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_CreateNetwork
//go:noescape
func TryCreateNetwork(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	shared js.Ref,
	properties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_DisableNetworkType
//go:noescape
func HasFuncDisableNetworkType() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_DisableNetworkType
//go:noescape
func FuncDisableNetworkType(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_DisableNetworkType
//go:noescape
func CallDisableNetworkType(
	retPtr unsafe.Pointer,
	networkType uint32)

//go:wasmimport plat/js/webext/networking/onc try_DisableNetworkType
//go:noescape
func TryDisableNetworkType(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_EnableNetworkType
//go:noescape
func HasFuncEnableNetworkType() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_EnableNetworkType
//go:noescape
func FuncEnableNetworkType(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_EnableNetworkType
//go:noescape
func CallEnableNetworkType(
	retPtr unsafe.Pointer,
	networkType uint32)

//go:wasmimport plat/js/webext/networking/onc try_EnableNetworkType
//go:noescape
func TryEnableNetworkType(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_ForgetNetwork
//go:noescape
func HasFuncForgetNetwork() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_ForgetNetwork
//go:noescape
func FuncForgetNetwork(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_ForgetNetwork
//go:noescape
func CallForgetNetwork(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_ForgetNetwork
//go:noescape
func TryForgetNetwork(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_GetCaptivePortalStatus
//go:noescape
func HasFuncGetCaptivePortalStatus() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_GetCaptivePortalStatus
//go:noescape
func FuncGetCaptivePortalStatus(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_GetCaptivePortalStatus
//go:noescape
func CallGetCaptivePortalStatus(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_GetCaptivePortalStatus
//go:noescape
func TryGetCaptivePortalStatus(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_GetDeviceStates
//go:noescape
func HasFuncGetDeviceStates() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_GetDeviceStates
//go:noescape
func FuncGetDeviceStates(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_GetDeviceStates
//go:noescape
func CallGetDeviceStates(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_GetDeviceStates
//go:noescape
func TryGetDeviceStates(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_GetGlobalPolicy
//go:noescape
func HasFuncGetGlobalPolicy() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_GetGlobalPolicy
//go:noescape
func FuncGetGlobalPolicy(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_GetGlobalPolicy
//go:noescape
func CallGetGlobalPolicy(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_GetGlobalPolicy
//go:noescape
func TryGetGlobalPolicy(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_GetManagedProperties
//go:noescape
func HasFuncGetManagedProperties() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_GetManagedProperties
//go:noescape
func FuncGetManagedProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_GetManagedProperties
//go:noescape
func CallGetManagedProperties(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_GetManagedProperties
//go:noescape
func TryGetManagedProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_GetNetworks
//go:noescape
func HasFuncGetNetworks() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_GetNetworks
//go:noescape
func FuncGetNetworks(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_GetNetworks
//go:noescape
func CallGetNetworks(
	retPtr unsafe.Pointer,
	filter unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_GetNetworks
//go:noescape
func TryGetNetworks(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	filter unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_GetProperties
//go:noescape
func HasFuncGetProperties() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_GetProperties
//go:noescape
func FuncGetProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_GetProperties
//go:noescape
func CallGetProperties(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_GetProperties
//go:noescape
func TryGetProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_GetState
//go:noescape
func HasFuncGetState() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_GetState
//go:noescape
func FuncGetState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_GetState
//go:noescape
func CallGetState(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_GetState
//go:noescape
func TryGetState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_OnDeviceStateListChanged
//go:noescape
func HasFuncOnDeviceStateListChanged() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_OnDeviceStateListChanged
//go:noescape
func FuncOnDeviceStateListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_OnDeviceStateListChanged
//go:noescape
func CallOnDeviceStateListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_OnDeviceStateListChanged
//go:noescape
func TryOnDeviceStateListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_OffDeviceStateListChanged
//go:noescape
func HasFuncOffDeviceStateListChanged() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_OffDeviceStateListChanged
//go:noescape
func FuncOffDeviceStateListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_OffDeviceStateListChanged
//go:noescape
func CallOffDeviceStateListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_OffDeviceStateListChanged
//go:noescape
func TryOffDeviceStateListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_HasOnDeviceStateListChanged
//go:noescape
func HasFuncHasOnDeviceStateListChanged() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_HasOnDeviceStateListChanged
//go:noescape
func FuncHasOnDeviceStateListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_HasOnDeviceStateListChanged
//go:noescape
func CallHasOnDeviceStateListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_HasOnDeviceStateListChanged
//go:noescape
func TryHasOnDeviceStateListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_OnNetworkListChanged
//go:noescape
func HasFuncOnNetworkListChanged() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_OnNetworkListChanged
//go:noescape
func FuncOnNetworkListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_OnNetworkListChanged
//go:noescape
func CallOnNetworkListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_OnNetworkListChanged
//go:noescape
func TryOnNetworkListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_OffNetworkListChanged
//go:noescape
func HasFuncOffNetworkListChanged() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_OffNetworkListChanged
//go:noescape
func FuncOffNetworkListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_OffNetworkListChanged
//go:noescape
func CallOffNetworkListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_OffNetworkListChanged
//go:noescape
func TryOffNetworkListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_HasOnNetworkListChanged
//go:noescape
func HasFuncHasOnNetworkListChanged() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_HasOnNetworkListChanged
//go:noescape
func FuncHasOnNetworkListChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_HasOnNetworkListChanged
//go:noescape
func CallHasOnNetworkListChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_HasOnNetworkListChanged
//go:noescape
func TryHasOnNetworkListChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_OnNetworksChanged
//go:noescape
func HasFuncOnNetworksChanged() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_OnNetworksChanged
//go:noescape
func FuncOnNetworksChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_OnNetworksChanged
//go:noescape
func CallOnNetworksChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_OnNetworksChanged
//go:noescape
func TryOnNetworksChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_OffNetworksChanged
//go:noescape
func HasFuncOffNetworksChanged() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_OffNetworksChanged
//go:noescape
func FuncOffNetworksChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_OffNetworksChanged
//go:noescape
func CallOffNetworksChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_OffNetworksChanged
//go:noescape
func TryOffNetworksChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_HasOnNetworksChanged
//go:noescape
func HasFuncHasOnNetworksChanged() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_HasOnNetworksChanged
//go:noescape
func FuncHasOnNetworksChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_HasOnNetworksChanged
//go:noescape
func CallHasOnNetworksChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_HasOnNetworksChanged
//go:noescape
func TryHasOnNetworksChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_OnPortalDetectionCompleted
//go:noescape
func HasFuncOnPortalDetectionCompleted() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_OnPortalDetectionCompleted
//go:noescape
func FuncOnPortalDetectionCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_OnPortalDetectionCompleted
//go:noescape
func CallOnPortalDetectionCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_OnPortalDetectionCompleted
//go:noescape
func TryOnPortalDetectionCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_OffPortalDetectionCompleted
//go:noescape
func HasFuncOffPortalDetectionCompleted() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_OffPortalDetectionCompleted
//go:noescape
func FuncOffPortalDetectionCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_OffPortalDetectionCompleted
//go:noescape
func CallOffPortalDetectionCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_OffPortalDetectionCompleted
//go:noescape
func TryOffPortalDetectionCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_HasOnPortalDetectionCompleted
//go:noescape
func HasFuncHasOnPortalDetectionCompleted() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_HasOnPortalDetectionCompleted
//go:noescape
func FuncHasOnPortalDetectionCompleted(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_HasOnPortalDetectionCompleted
//go:noescape
func CallHasOnPortalDetectionCompleted(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_HasOnPortalDetectionCompleted
//go:noescape
func TryHasOnPortalDetectionCompleted(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_RequestNetworkScan
//go:noescape
func HasFuncRequestNetworkScan() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_RequestNetworkScan
//go:noescape
func FuncRequestNetworkScan(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_RequestNetworkScan
//go:noescape
func CallRequestNetworkScan(
	retPtr unsafe.Pointer,
	networkType uint32)

//go:wasmimport plat/js/webext/networking/onc try_RequestNetworkScan
//go:noescape
func TryRequestNetworkScan(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkType uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_SetProperties
//go:noescape
func HasFuncSetProperties() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_SetProperties
//go:noescape
func FuncSetProperties(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_SetProperties
//go:noescape
func CallSetProperties(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	properties unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_SetProperties
//go:noescape
func TrySetProperties(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	properties unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_StartConnect
//go:noescape
func HasFuncStartConnect() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_StartConnect
//go:noescape
func FuncStartConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_StartConnect
//go:noescape
func CallStartConnect(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_StartConnect
//go:noescape
func TryStartConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/networking/onc has_StartDisconnect
//go:noescape
func HasFuncStartDisconnect() js.Ref

//go:wasmimport plat/js/webext/networking/onc func_StartDisconnect
//go:noescape
func FuncStartDisconnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/networking/onc call_StartDisconnect
//go:noescape
func CallStartDisconnect(
	retPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref)

//go:wasmimport plat/js/webext/networking/onc try_StartDisconnect
//go:noescape
func TryStartDisconnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	networkGuid js.Ref,
	callback js.Ref) (ok js.Ref)

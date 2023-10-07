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

//go:wasmimport plat/js/webext/bluetoothprivate constof_ConnectResultType
//go:noescape
func ConstOfConnectResultType(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetoothprivate constof_TransportType
//go:noescape
func ConstOfTransportType(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetoothprivate store_DiscoveryFilter
//go:noescape
func DiscoveryFilterJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate load_DiscoveryFilter
//go:noescape
func DiscoveryFilterJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate store_NewAdapterState
//go:noescape
func NewAdapterStateJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate load_NewAdapterState
//go:noescape
func NewAdapterStateJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate constof_PairingEventType
//go:noescape
func ConstOfPairingEventType(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetoothprivate store_PairingEvent
//go:noescape
func PairingEventJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate load_PairingEvent
//go:noescape
func PairingEventJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate constof_PairingResponse
//go:noescape
func ConstOfPairingResponse(str js.Ref) uint32

//go:wasmimport plat/js/webext/bluetoothprivate store_SetPairingResponseOptions
//go:noescape
func SetPairingResponseOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate load_SetPairingResponseOptions
//go:noescape
func SetPairingResponseOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate has_Connect
//go:noescape
func HasFuncConnect() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_Connect
//go:noescape
func FuncConnect(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_Connect
//go:noescape
func CallConnect(
	retPtr unsafe.Pointer,
	deviceAddress js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate try_Connect
//go:noescape
func TryConnect(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deviceAddress js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_DisconnectAll
//go:noescape
func HasFuncDisconnectAll() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_DisconnectAll
//go:noescape
func FuncDisconnectAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_DisconnectAll
//go:noescape
func CallDisconnectAll(
	retPtr unsafe.Pointer,
	deviceAddress js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate try_DisconnectAll
//go:noescape
func TryDisconnectAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deviceAddress js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_ForgetDevice
//go:noescape
func HasFuncForgetDevice() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_ForgetDevice
//go:noescape
func FuncForgetDevice(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_ForgetDevice
//go:noescape
func CallForgetDevice(
	retPtr unsafe.Pointer,
	deviceAddress js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate try_ForgetDevice
//go:noescape
func TryForgetDevice(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deviceAddress js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_OnDeviceAddressChanged
//go:noescape
func HasFuncOnDeviceAddressChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_OnDeviceAddressChanged
//go:noescape
func FuncOnDeviceAddressChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_OnDeviceAddressChanged
//go:noescape
func CallOnDeviceAddressChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate try_OnDeviceAddressChanged
//go:noescape
func TryOnDeviceAddressChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_OffDeviceAddressChanged
//go:noescape
func HasFuncOffDeviceAddressChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_OffDeviceAddressChanged
//go:noescape
func FuncOffDeviceAddressChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_OffDeviceAddressChanged
//go:noescape
func CallOffDeviceAddressChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate try_OffDeviceAddressChanged
//go:noescape
func TryOffDeviceAddressChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_HasOnDeviceAddressChanged
//go:noescape
func HasFuncHasOnDeviceAddressChanged() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_HasOnDeviceAddressChanged
//go:noescape
func FuncHasOnDeviceAddressChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_HasOnDeviceAddressChanged
//go:noescape
func CallHasOnDeviceAddressChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate try_HasOnDeviceAddressChanged
//go:noescape
func TryHasOnDeviceAddressChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_OnPairing
//go:noescape
func HasFuncOnPairing() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_OnPairing
//go:noescape
func FuncOnPairing(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_OnPairing
//go:noescape
func CallOnPairing(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate try_OnPairing
//go:noescape
func TryOnPairing(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_OffPairing
//go:noescape
func HasFuncOffPairing() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_OffPairing
//go:noescape
func FuncOffPairing(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_OffPairing
//go:noescape
func CallOffPairing(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate try_OffPairing
//go:noescape
func TryOffPairing(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_HasOnPairing
//go:noescape
func HasFuncHasOnPairing() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_HasOnPairing
//go:noescape
func FuncHasOnPairing(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_HasOnPairing
//go:noescape
func CallHasOnPairing(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate try_HasOnPairing
//go:noescape
func TryHasOnPairing(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_Pair
//go:noescape
func HasFuncPair() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_Pair
//go:noescape
func FuncPair(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_Pair
//go:noescape
func CallPair(
	retPtr unsafe.Pointer,
	deviceAddress js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate try_Pair
//go:noescape
func TryPair(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	deviceAddress js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_RecordDeviceSelection
//go:noescape
func HasFuncRecordDeviceSelection() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_RecordDeviceSelection
//go:noescape
func FuncRecordDeviceSelection(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_RecordDeviceSelection
//go:noescape
func CallRecordDeviceSelection(
	retPtr unsafe.Pointer,
	selectionDurationMs int32,
	wasPaired js.Ref,
	transport uint32)

//go:wasmimport plat/js/webext/bluetoothprivate try_RecordDeviceSelection
//go:noescape
func TryRecordDeviceSelection(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	selectionDurationMs int32,
	wasPaired js.Ref,
	transport uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_RecordPairing
//go:noescape
func HasFuncRecordPairing() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_RecordPairing
//go:noescape
func FuncRecordPairing(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_RecordPairing
//go:noescape
func CallRecordPairing(
	retPtr unsafe.Pointer,
	transport uint32,
	pairingDurationMs int32,
	result uint32)

//go:wasmimport plat/js/webext/bluetoothprivate try_RecordPairing
//go:noescape
func TryRecordPairing(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	transport uint32,
	pairingDurationMs int32,
	result uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_RecordReconnection
//go:noescape
func HasFuncRecordReconnection() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_RecordReconnection
//go:noescape
func FuncRecordReconnection(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_RecordReconnection
//go:noescape
func CallRecordReconnection(
	retPtr unsafe.Pointer,
	result uint32)

//go:wasmimport plat/js/webext/bluetoothprivate try_RecordReconnection
//go:noescape
func TryRecordReconnection(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	result uint32) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_SetAdapterState
//go:noescape
func HasFuncSetAdapterState() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_SetAdapterState
//go:noescape
func FuncSetAdapterState(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_SetAdapterState
//go:noescape
func CallSetAdapterState(
	retPtr unsafe.Pointer,
	adapterState unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate try_SetAdapterState
//go:noescape
func TrySetAdapterState(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	adapterState unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_SetDiscoveryFilter
//go:noescape
func HasFuncSetDiscoveryFilter() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_SetDiscoveryFilter
//go:noescape
func FuncSetDiscoveryFilter(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_SetDiscoveryFilter
//go:noescape
func CallSetDiscoveryFilter(
	retPtr unsafe.Pointer,
	discoveryFilter unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate try_SetDiscoveryFilter
//go:noescape
func TrySetDiscoveryFilter(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	discoveryFilter unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bluetoothprivate has_SetPairingResponse
//go:noescape
func HasFuncSetPairingResponse() js.Ref

//go:wasmimport plat/js/webext/bluetoothprivate func_SetPairingResponse
//go:noescape
func FuncSetPairingResponse(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate call_SetPairingResponse
//go:noescape
func CallSetPairingResponse(
	retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/webext/bluetoothprivate try_SetPairingResponse
//go:noescape
func TrySetPairingResponse(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

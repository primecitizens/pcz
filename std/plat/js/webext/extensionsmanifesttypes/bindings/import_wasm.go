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

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_AutomationChoice1
//go:noescape
func AutomationChoice1JSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_AutomationChoice1
//go:noescape
func AutomationChoice1JSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_Bluetooth
//go:noescape
func BluetoothJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_Bluetooth
//go:noescape
func BluetoothJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_ContentCapabilities
//go:noescape
func ContentCapabilitiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_ContentCapabilities
//go:noescape
func ContentCapabilitiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_ExternallyConnectable
//go:noescape
func ExternallyConnectableJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_ExternallyConnectable
//go:noescape
func ExternallyConnectableJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_KioskSecondaryAppsElem
//go:noescape
func KioskSecondaryAppsElemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_KioskSecondaryAppsElem
//go:noescape
func KioskSecondaryAppsElemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_OptionsUI
//go:noescape
func OptionsUIJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_OptionsUI
//go:noescape
func OptionsUIJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_SocketsFieldTcp
//go:noescape
func SocketsFieldTcpJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_SocketsFieldTcp
//go:noescape
func SocketsFieldTcpJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_SocketsFieldTcpServer
//go:noescape
func SocketsFieldTcpServerJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_SocketsFieldTcpServer
//go:noescape
func SocketsFieldTcpServerJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_SocketsFieldUdp
//go:noescape
func SocketsFieldUdpJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_SocketsFieldUdp
//go:noescape
func SocketsFieldUdpJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_Sockets
//go:noescape
func SocketsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_Sockets
//go:noescape
func SocketsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_UsbPrintersFieldFiltersElem
//go:noescape
func UsbPrintersFieldFiltersElemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_UsbPrintersFieldFiltersElem
//go:noescape
func UsbPrintersFieldFiltersElemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/extensionsmanifesttypes store_UsbPrinters
//go:noescape
func UsbPrintersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/extensionsmanifesttypes load_UsbPrinters
//go:noescape
func UsbPrintersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package extensionsmanifesttypes

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensionsmanifesttypes/bindings"
)

type AutomationChoice1 struct {
	// Desktop is "AutomationChoice1.desktop"
	//
	// Optional
	//
	// NOTE: FFI_USE_Desktop MUST be set to true to make this field effective.
	Desktop bool
	// Interact is "AutomationChoice1.interact"
	//
	// Optional
	//
	// NOTE: FFI_USE_Interact MUST be set to true to make this field effective.
	Interact bool
	// Matches is "AutomationChoice1.matches"
	//
	// Optional
	Matches js.Array[js.String]

	FFI_USE_Desktop  bool // for Desktop.
	FFI_USE_Interact bool // for Interact.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AutomationChoice1 with all fields set.
func (p AutomationChoice1) FromRef(ref js.Ref) AutomationChoice1 {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AutomationChoice1 in the application heap.
func (p AutomationChoice1) New() js.Ref {
	return bindings.AutomationChoice1JSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AutomationChoice1) UpdateFrom(ref js.Ref) {
	bindings.AutomationChoice1JSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AutomationChoice1) Update(ref js.Ref) {
	bindings.AutomationChoice1JSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AutomationChoice1) FreeMembers(recursive bool) {
	js.Free(
		p.Matches.Ref(),
	)
	p.Matches = p.Matches.FromRef(js.Undefined)
}

type OneOf_Bool_AutomationChoice1 struct {
	ref js.Ref
}

func (x OneOf_Bool_AutomationChoice1) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Bool_AutomationChoice1) Free() {
	x.ref.Free()
}

func (x OneOf_Bool_AutomationChoice1) FromRef(ref js.Ref) OneOf_Bool_AutomationChoice1 {
	return OneOf_Bool_AutomationChoice1{
		ref: ref,
	}
}

func (x OneOf_Bool_AutomationChoice1) Bool() bool {
	return x.ref == js.True
}

func (x OneOf_Bool_AutomationChoice1) AutomationChoice1() AutomationChoice1 {
	var ret AutomationChoice1
	ret.UpdateFrom(x.ref)
	return ret
}

type Automation = OneOf_Bool_AutomationChoice1

type Bluetooth struct {
	// LowEnergy is "Bluetooth.low_energy"
	//
	// Optional
	//
	// NOTE: FFI_USE_LowEnergy MUST be set to true to make this field effective.
	LowEnergy bool
	// Peripheral is "Bluetooth.peripheral"
	//
	// Optional
	//
	// NOTE: FFI_USE_Peripheral MUST be set to true to make this field effective.
	Peripheral bool
	// Socket is "Bluetooth.socket"
	//
	// Optional
	//
	// NOTE: FFI_USE_Socket MUST be set to true to make this field effective.
	Socket bool
	// Uuids is "Bluetooth.uuids"
	//
	// Optional
	Uuids js.Array[js.String]

	FFI_USE_LowEnergy  bool // for LowEnergy.
	FFI_USE_Peripheral bool // for Peripheral.
	FFI_USE_Socket     bool // for Socket.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Bluetooth with all fields set.
func (p Bluetooth) FromRef(ref js.Ref) Bluetooth {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Bluetooth in the application heap.
func (p Bluetooth) New() js.Ref {
	return bindings.BluetoothJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Bluetooth) UpdateFrom(ref js.Ref) {
	bindings.BluetoothJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Bluetooth) Update(ref js.Ref) {
	bindings.BluetoothJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Bluetooth) FreeMembers(recursive bool) {
	js.Free(
		p.Uuids.Ref(),
	)
	p.Uuids = p.Uuids.FromRef(js.Undefined)
}

type ContentCapabilities struct {
	// Matches is "ContentCapabilities.matches"
	//
	// Required
	Matches js.Array[js.String]
	// Permissions is "ContentCapabilities.permissions"
	//
	// Required
	Permissions js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContentCapabilities with all fields set.
func (p ContentCapabilities) FromRef(ref js.Ref) ContentCapabilities {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContentCapabilities in the application heap.
func (p ContentCapabilities) New() js.Ref {
	return bindings.ContentCapabilitiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContentCapabilities) UpdateFrom(ref js.Ref) {
	bindings.ContentCapabilitiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContentCapabilities) Update(ref js.Ref) {
	bindings.ContentCapabilitiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContentCapabilities) FreeMembers(recursive bool) {
	js.Free(
		p.Matches.Ref(),
		p.Permissions.Ref(),
	)
	p.Matches = p.Matches.FromRef(js.Undefined)
	p.Permissions = p.Permissions.FromRef(js.Undefined)
}

type ExternallyConnectable struct {
	// AcceptsTlsChannelId is "ExternallyConnectable.accepts_tls_channel_id"
	//
	// Optional
	//
	// NOTE: FFI_USE_AcceptsTlsChannelId MUST be set to true to make this field effective.
	AcceptsTlsChannelId bool
	// Ids is "ExternallyConnectable.ids"
	//
	// Optional
	Ids js.Array[js.String]
	// Matches is "ExternallyConnectable.matches"
	//
	// Optional
	Matches js.Array[js.String]

	FFI_USE_AcceptsTlsChannelId bool // for AcceptsTlsChannelId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ExternallyConnectable with all fields set.
func (p ExternallyConnectable) FromRef(ref js.Ref) ExternallyConnectable {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExternallyConnectable in the application heap.
func (p ExternallyConnectable) New() js.Ref {
	return bindings.ExternallyConnectableJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ExternallyConnectable) UpdateFrom(ref js.Ref) {
	bindings.ExternallyConnectableJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ExternallyConnectable) Update(ref js.Ref) {
	bindings.ExternallyConnectableJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ExternallyConnectable) FreeMembers(recursive bool) {
	js.Free(
		p.Ids.Ref(),
		p.Matches.Ref(),
	)
	p.Ids = p.Ids.FromRef(js.Undefined)
	p.Matches = p.Matches.FromRef(js.Undefined)
}

type KioskSecondaryAppsElem struct {
	// EnabledOnLaunch is "KioskSecondaryAppsElem.enabled_on_launch"
	//
	// Optional
	//
	// NOTE: FFI_USE_EnabledOnLaunch MUST be set to true to make this field effective.
	EnabledOnLaunch bool
	// Id is "KioskSecondaryAppsElem.id"
	//
	// Required
	Id js.String

	FFI_USE_EnabledOnLaunch bool // for EnabledOnLaunch.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a KioskSecondaryAppsElem with all fields set.
func (p KioskSecondaryAppsElem) FromRef(ref js.Ref) KioskSecondaryAppsElem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new KioskSecondaryAppsElem in the application heap.
func (p KioskSecondaryAppsElem) New() js.Ref {
	return bindings.KioskSecondaryAppsElemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *KioskSecondaryAppsElem) UpdateFrom(ref js.Ref) {
	bindings.KioskSecondaryAppsElemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *KioskSecondaryAppsElem) Update(ref js.Ref) {
	bindings.KioskSecondaryAppsElemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *KioskSecondaryAppsElem) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
}

type KioskSecondaryApps = js.Array[KioskSecondaryAppsElem]

type OptionsUI struct {
	// ChromeStyle is "OptionsUI.chrome_style"
	//
	// Optional
	//
	// NOTE: FFI_USE_ChromeStyle MUST be set to true to make this field effective.
	ChromeStyle bool
	// OpenInTab is "OptionsUI.open_in_tab"
	//
	// Optional
	//
	// NOTE: FFI_USE_OpenInTab MUST be set to true to make this field effective.
	OpenInTab bool
	// Page is "OptionsUI.page"
	//
	// Required
	Page js.String

	FFI_USE_ChromeStyle bool // for ChromeStyle.
	FFI_USE_OpenInTab   bool // for OpenInTab.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OptionsUI with all fields set.
func (p OptionsUI) FromRef(ref js.Ref) OptionsUI {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OptionsUI in the application heap.
func (p OptionsUI) New() js.Ref {
	return bindings.OptionsUIJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OptionsUI) UpdateFrom(ref js.Ref) {
	bindings.OptionsUIJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OptionsUI) Update(ref js.Ref) {
	bindings.OptionsUIJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OptionsUI) FreeMembers(recursive bool) {
	js.Free(
		p.Page.Ref(),
	)
	p.Page = p.Page.FromRef(js.Undefined)
}

type OneOf_String_ArrayString struct {
	ref js.Ref
}

func (x OneOf_String_ArrayString) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayString) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayString) FromRef(ref js.Ref) OneOf_String_ArrayString {
	return OneOf_String_ArrayString{
		ref: ref,
	}
}

func (x OneOf_String_ArrayString) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayString) ArrayString() js.Array[js.String] {
	return js.Array[js.String]{}.FromRef(x.ref)
}

type SocketHostPatterns = OneOf_String_ArrayString

type SocketsFieldTcp struct {
	// Connect is "SocketsFieldTcp.connect"
	//
	// Optional
	Connect SocketHostPatterns

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SocketsFieldTcp with all fields set.
func (p SocketsFieldTcp) FromRef(ref js.Ref) SocketsFieldTcp {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SocketsFieldTcp in the application heap.
func (p SocketsFieldTcp) New() js.Ref {
	return bindings.SocketsFieldTcpJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SocketsFieldTcp) UpdateFrom(ref js.Ref) {
	bindings.SocketsFieldTcpJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SocketsFieldTcp) Update(ref js.Ref) {
	bindings.SocketsFieldTcpJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SocketsFieldTcp) FreeMembers(recursive bool) {
	js.Free(
		p.Connect.Ref(),
	)
	p.Connect = p.Connect.FromRef(js.Undefined)
}

type SocketsFieldTcpServer struct {
	// Listen is "SocketsFieldTcpServer.listen"
	//
	// Optional
	Listen SocketHostPatterns

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SocketsFieldTcpServer with all fields set.
func (p SocketsFieldTcpServer) FromRef(ref js.Ref) SocketsFieldTcpServer {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SocketsFieldTcpServer in the application heap.
func (p SocketsFieldTcpServer) New() js.Ref {
	return bindings.SocketsFieldTcpServerJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SocketsFieldTcpServer) UpdateFrom(ref js.Ref) {
	bindings.SocketsFieldTcpServerJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SocketsFieldTcpServer) Update(ref js.Ref) {
	bindings.SocketsFieldTcpServerJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SocketsFieldTcpServer) FreeMembers(recursive bool) {
	js.Free(
		p.Listen.Ref(),
	)
	p.Listen = p.Listen.FromRef(js.Undefined)
}

type SocketsFieldUdp struct {
	// Bind is "SocketsFieldUdp.bind"
	//
	// Optional
	Bind SocketHostPatterns
	// MulticastMembership is "SocketsFieldUdp.multicastMembership"
	//
	// Optional
	MulticastMembership SocketHostPatterns
	// Send is "SocketsFieldUdp.send"
	//
	// Optional
	Send SocketHostPatterns

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SocketsFieldUdp with all fields set.
func (p SocketsFieldUdp) FromRef(ref js.Ref) SocketsFieldUdp {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SocketsFieldUdp in the application heap.
func (p SocketsFieldUdp) New() js.Ref {
	return bindings.SocketsFieldUdpJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SocketsFieldUdp) UpdateFrom(ref js.Ref) {
	bindings.SocketsFieldUdpJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SocketsFieldUdp) Update(ref js.Ref) {
	bindings.SocketsFieldUdpJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SocketsFieldUdp) FreeMembers(recursive bool) {
	js.Free(
		p.Bind.Ref(),
		p.MulticastMembership.Ref(),
		p.Send.Ref(),
	)
	p.Bind = p.Bind.FromRef(js.Undefined)
	p.MulticastMembership = p.MulticastMembership.FromRef(js.Undefined)
	p.Send = p.Send.FromRef(js.Undefined)
}

type Sockets struct {
	// Tcp is "Sockets.tcp"
	//
	// Optional
	//
	// NOTE: Tcp.FFI_USE MUST be set to true to get Tcp used.
	Tcp SocketsFieldTcp
	// TcpServer is "Sockets.tcpServer"
	//
	// Optional
	//
	// NOTE: TcpServer.FFI_USE MUST be set to true to get TcpServer used.
	TcpServer SocketsFieldTcpServer
	// Udp is "Sockets.udp"
	//
	// Optional
	//
	// NOTE: Udp.FFI_USE MUST be set to true to get Udp used.
	Udp SocketsFieldUdp

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Sockets with all fields set.
func (p Sockets) FromRef(ref js.Ref) Sockets {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Sockets in the application heap.
func (p Sockets) New() js.Ref {
	return bindings.SocketsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Sockets) UpdateFrom(ref js.Ref) {
	bindings.SocketsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Sockets) Update(ref js.Ref) {
	bindings.SocketsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Sockets) FreeMembers(recursive bool) {
	if recursive {
		p.Tcp.FreeMembers(true)
		p.TcpServer.FreeMembers(true)
		p.Udp.FreeMembers(true)
	}
}

type UsbPrintersFieldFiltersElem struct {
	// InterfaceClass is "UsbPrintersFieldFiltersElem.interfaceClass"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceClass MUST be set to true to make this field effective.
	InterfaceClass int64
	// InterfaceProtocol is "UsbPrintersFieldFiltersElem.interfaceProtocol"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceProtocol MUST be set to true to make this field effective.
	InterfaceProtocol int64
	// InterfaceSubclass is "UsbPrintersFieldFiltersElem.interfaceSubclass"
	//
	// Optional
	//
	// NOTE: FFI_USE_InterfaceSubclass MUST be set to true to make this field effective.
	InterfaceSubclass int64
	// ProductId is "UsbPrintersFieldFiltersElem.productId"
	//
	// Optional
	//
	// NOTE: FFI_USE_ProductId MUST be set to true to make this field effective.
	ProductId int64
	// VendorId is "UsbPrintersFieldFiltersElem.vendorId"
	//
	// Required
	VendorId int64

	FFI_USE_InterfaceClass    bool // for InterfaceClass.
	FFI_USE_InterfaceProtocol bool // for InterfaceProtocol.
	FFI_USE_InterfaceSubclass bool // for InterfaceSubclass.
	FFI_USE_ProductId         bool // for ProductId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UsbPrintersFieldFiltersElem with all fields set.
func (p UsbPrintersFieldFiltersElem) FromRef(ref js.Ref) UsbPrintersFieldFiltersElem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UsbPrintersFieldFiltersElem in the application heap.
func (p UsbPrintersFieldFiltersElem) New() js.Ref {
	return bindings.UsbPrintersFieldFiltersElemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UsbPrintersFieldFiltersElem) UpdateFrom(ref js.Ref) {
	bindings.UsbPrintersFieldFiltersElemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UsbPrintersFieldFiltersElem) Update(ref js.Ref) {
	bindings.UsbPrintersFieldFiltersElemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UsbPrintersFieldFiltersElem) FreeMembers(recursive bool) {
}

type UsbPrinters struct {
	// Filters is "UsbPrinters.filters"
	//
	// Required
	Filters js.Array[UsbPrintersFieldFiltersElem]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UsbPrinters with all fields set.
func (p UsbPrinters) FromRef(ref js.Ref) UsbPrinters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UsbPrinters in the application heap.
func (p UsbPrinters) New() js.Ref {
	return bindings.UsbPrintersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UsbPrinters) UpdateFrom(ref js.Ref) {
	bindings.UsbPrintersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UsbPrinters) Update(ref js.Ref) {
	bindings.UsbPrintersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UsbPrinters) FreeMembers(recursive bool) {
	js.Free(
		p.Filters.Ref(),
	)
	p.Filters = p.Filters.FromRef(js.Undefined)
}

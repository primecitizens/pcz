// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package vpnprovider

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/vpnprovider/bindings"
)

type CallCompleteCallbackFunc func(this js.Ref) js.Ref

func (fn CallCompleteCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CallCompleteCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CallCompleteCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CallCompleteCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CallCompleteCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateConfigCompleteCallbackFunc func(this js.Ref, id js.String) js.Ref

func (fn CreateConfigCompleteCallbackFunc) Register() js.Func[func(id js.String)] {
	return js.RegisterCallback[func(id js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateConfigCompleteCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CreateConfigCompleteCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String) js.Ref
	Arg T
}

func (cb *CreateConfigCompleteCallback[T]) Register() js.Func[func(id js.String)] {
	return js.RegisterCallback[func(id js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateConfigCompleteCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Parameters struct {
	// Address is "Parameters.address"
	//
	// Optional
	Address js.String
	// BroadcastAddress is "Parameters.broadcastAddress"
	//
	// Optional
	BroadcastAddress js.String
	// Mtu is "Parameters.mtu"
	//
	// Optional
	Mtu js.String
	// ExclusionList is "Parameters.exclusionList"
	//
	// Optional
	ExclusionList js.Array[js.String]
	// InclusionList is "Parameters.inclusionList"
	//
	// Optional
	InclusionList js.Array[js.String]
	// DomainSearch is "Parameters.domainSearch"
	//
	// Optional
	DomainSearch js.Array[js.String]
	// DnsServers is "Parameters.dnsServers"
	//
	// Optional
	DnsServers js.Array[js.String]
	// Reconnect is "Parameters.reconnect"
	//
	// Optional
	Reconnect js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Parameters with all fields set.
func (p Parameters) FromRef(ref js.Ref) Parameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Parameters in the application heap.
func (p Parameters) New() js.Ref {
	return bindings.ParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Parameters) UpdateFrom(ref js.Ref) {
	bindings.ParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Parameters) Update(ref js.Ref) {
	bindings.ParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Parameters) FreeMembers(recursive bool) {
	js.Free(
		p.Address.Ref(),
		p.BroadcastAddress.Ref(),
		p.Mtu.Ref(),
		p.ExclusionList.Ref(),
		p.InclusionList.Ref(),
		p.DomainSearch.Ref(),
		p.DnsServers.Ref(),
		p.Reconnect.Ref(),
	)
	p.Address = p.Address.FromRef(js.Undefined)
	p.BroadcastAddress = p.BroadcastAddress.FromRef(js.Undefined)
	p.Mtu = p.Mtu.FromRef(js.Undefined)
	p.ExclusionList = p.ExclusionList.FromRef(js.Undefined)
	p.InclusionList = p.InclusionList.FromRef(js.Undefined)
	p.DomainSearch = p.DomainSearch.FromRef(js.Undefined)
	p.DnsServers = p.DnsServers.FromRef(js.Undefined)
	p.Reconnect = p.Reconnect.FromRef(js.Undefined)
}

type PlatformMessage uint32

const (
	_ PlatformMessage = iota

	PlatformMessage_CONNECTED
	PlatformMessage_DISCONNECTED
	PlatformMessage_ERROR
	PlatformMessage_LINK_DOWN
	PlatformMessage_LINK_UP
	PlatformMessage_LINK_CHANGED
	PlatformMessage_SUSPEND
	PlatformMessage_RESUME
)

func (PlatformMessage) FromRef(str js.Ref) PlatformMessage {
	return PlatformMessage(bindings.ConstOfPlatformMessage(str))
}

func (x PlatformMessage) String() (string, bool) {
	switch x {
	case PlatformMessage_CONNECTED:
		return "connected", true
	case PlatformMessage_DISCONNECTED:
		return "disconnected", true
	case PlatformMessage_ERROR:
		return "error", true
	case PlatformMessage_LINK_DOWN:
		return "linkDown", true
	case PlatformMessage_LINK_UP:
		return "linkUp", true
	case PlatformMessage_LINK_CHANGED:
		return "linkChanged", true
	case PlatformMessage_SUSPEND:
		return "suspend", true
	case PlatformMessage_RESUME:
		return "resume", true
	default:
		return "", false
	}
}

type UIEvent uint32

const (
	_ UIEvent = iota

	UIEvent_SHOW_ADD_DIALOG
	UIEvent_SHOW_CONFIGURE_DIALOG
)

func (UIEvent) FromRef(str js.Ref) UIEvent {
	return UIEvent(bindings.ConstOfUIEvent(str))
}

func (x UIEvent) String() (string, bool) {
	switch x {
	case UIEvent_SHOW_ADD_DIALOG:
		return "showAddDialog", true
	case UIEvent_SHOW_CONFIGURE_DIALOG:
		return "showConfigureDialog", true
	default:
		return "", false
	}
}

type VpnConnectionState uint32

const (
	_ VpnConnectionState = iota

	VpnConnectionState_CONNECTED
	VpnConnectionState_FAILURE
)

func (VpnConnectionState) FromRef(str js.Ref) VpnConnectionState {
	return VpnConnectionState(bindings.ConstOfVpnConnectionState(str))
}

func (x VpnConnectionState) String() (string, bool) {
	switch x {
	case VpnConnectionState_CONNECTED:
		return "connected", true
	case VpnConnectionState_FAILURE:
		return "failure", true
	default:
		return "", false
	}
}

// HasFuncCreateConfig returns true if the function "WEBEXT.vpnProvider.createConfig" exists.
func HasFuncCreateConfig() bool {
	return js.True == bindings.HasFuncCreateConfig()
}

// FuncCreateConfig returns the function "WEBEXT.vpnProvider.createConfig".
func FuncCreateConfig() (fn js.Func[func(name js.String) js.Promise[js.String]]) {
	bindings.FuncCreateConfig(
		js.Pointer(&fn),
	)
	return
}

// CreateConfig calls the function "WEBEXT.vpnProvider.createConfig" directly.
func CreateConfig(name js.String) (ret js.Promise[js.String]) {
	bindings.CallCreateConfig(
		js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryCreateConfig calls the function "WEBEXT.vpnProvider.createConfig"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCreateConfig(name js.String) (ret js.Promise[js.String], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCreateConfig(
		js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncDestroyConfig returns true if the function "WEBEXT.vpnProvider.destroyConfig" exists.
func HasFuncDestroyConfig() bool {
	return js.True == bindings.HasFuncDestroyConfig()
}

// FuncDestroyConfig returns the function "WEBEXT.vpnProvider.destroyConfig".
func FuncDestroyConfig() (fn js.Func[func(id js.String) js.Promise[js.Void]]) {
	bindings.FuncDestroyConfig(
		js.Pointer(&fn),
	)
	return
}

// DestroyConfig calls the function "WEBEXT.vpnProvider.destroyConfig" directly.
func DestroyConfig(id js.String) (ret js.Promise[js.Void]) {
	bindings.CallDestroyConfig(
		js.Pointer(&ret),
		id.Ref(),
	)

	return
}

// TryDestroyConfig calls the function "WEBEXT.vpnProvider.destroyConfig"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDestroyConfig(id js.String) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDestroyConfig(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
	)

	return
}

// HasFuncNotifyConnectionStateChanged returns true if the function "WEBEXT.vpnProvider.notifyConnectionStateChanged" exists.
func HasFuncNotifyConnectionStateChanged() bool {
	return js.True == bindings.HasFuncNotifyConnectionStateChanged()
}

// FuncNotifyConnectionStateChanged returns the function "WEBEXT.vpnProvider.notifyConnectionStateChanged".
func FuncNotifyConnectionStateChanged() (fn js.Func[func(state VpnConnectionState) js.Promise[js.Void]]) {
	bindings.FuncNotifyConnectionStateChanged(
		js.Pointer(&fn),
	)
	return
}

// NotifyConnectionStateChanged calls the function "WEBEXT.vpnProvider.notifyConnectionStateChanged" directly.
func NotifyConnectionStateChanged(state VpnConnectionState) (ret js.Promise[js.Void]) {
	bindings.CallNotifyConnectionStateChanged(
		js.Pointer(&ret),
		uint32(state),
	)

	return
}

// TryNotifyConnectionStateChanged calls the function "WEBEXT.vpnProvider.notifyConnectionStateChanged"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryNotifyConnectionStateChanged(state VpnConnectionState) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNotifyConnectionStateChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(state),
	)

	return
}

type OnConfigCreatedEventCallbackFunc func(this js.Ref, id js.String, name js.String, data js.Object) js.Ref

func (fn OnConfigCreatedEventCallbackFunc) Register() js.Func[func(id js.String, name js.String, data js.Object)] {
	return js.RegisterCallback[func(id js.String, name js.String, data js.Object)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnConfigCreatedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		js.Object{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnConfigCreatedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String, name js.String, data js.Object) js.Ref
	Arg T
}

func (cb *OnConfigCreatedEventCallback[T]) Register() js.Func[func(id js.String, name js.String, data js.Object)] {
	return js.RegisterCallback[func(id js.String, name js.String, data js.Object)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnConfigCreatedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		js.Object{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnConfigCreated returns true if the function "WEBEXT.vpnProvider.onConfigCreated.addListener" exists.
func HasFuncOnConfigCreated() bool {
	return js.True == bindings.HasFuncOnConfigCreated()
}

// FuncOnConfigCreated returns the function "WEBEXT.vpnProvider.onConfigCreated.addListener".
func FuncOnConfigCreated() (fn js.Func[func(callback js.Func[func(id js.String, name js.String, data js.Object)])]) {
	bindings.FuncOnConfigCreated(
		js.Pointer(&fn),
	)
	return
}

// OnConfigCreated calls the function "WEBEXT.vpnProvider.onConfigCreated.addListener" directly.
func OnConfigCreated(callback js.Func[func(id js.String, name js.String, data js.Object)]) (ret js.Void) {
	bindings.CallOnConfigCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnConfigCreated calls the function "WEBEXT.vpnProvider.onConfigCreated.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnConfigCreated(callback js.Func[func(id js.String, name js.String, data js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnConfigCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffConfigCreated returns true if the function "WEBEXT.vpnProvider.onConfigCreated.removeListener" exists.
func HasFuncOffConfigCreated() bool {
	return js.True == bindings.HasFuncOffConfigCreated()
}

// FuncOffConfigCreated returns the function "WEBEXT.vpnProvider.onConfigCreated.removeListener".
func FuncOffConfigCreated() (fn js.Func[func(callback js.Func[func(id js.String, name js.String, data js.Object)])]) {
	bindings.FuncOffConfigCreated(
		js.Pointer(&fn),
	)
	return
}

// OffConfigCreated calls the function "WEBEXT.vpnProvider.onConfigCreated.removeListener" directly.
func OffConfigCreated(callback js.Func[func(id js.String, name js.String, data js.Object)]) (ret js.Void) {
	bindings.CallOffConfigCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffConfigCreated calls the function "WEBEXT.vpnProvider.onConfigCreated.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffConfigCreated(callback js.Func[func(id js.String, name js.String, data js.Object)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffConfigCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnConfigCreated returns true if the function "WEBEXT.vpnProvider.onConfigCreated.hasListener" exists.
func HasFuncHasOnConfigCreated() bool {
	return js.True == bindings.HasFuncHasOnConfigCreated()
}

// FuncHasOnConfigCreated returns the function "WEBEXT.vpnProvider.onConfigCreated.hasListener".
func FuncHasOnConfigCreated() (fn js.Func[func(callback js.Func[func(id js.String, name js.String, data js.Object)]) bool]) {
	bindings.FuncHasOnConfigCreated(
		js.Pointer(&fn),
	)
	return
}

// HasOnConfigCreated calls the function "WEBEXT.vpnProvider.onConfigCreated.hasListener" directly.
func HasOnConfigCreated(callback js.Func[func(id js.String, name js.String, data js.Object)]) (ret bool) {
	bindings.CallHasOnConfigCreated(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnConfigCreated calls the function "WEBEXT.vpnProvider.onConfigCreated.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnConfigCreated(callback js.Func[func(id js.String, name js.String, data js.Object)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnConfigCreated(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnConfigRemovedEventCallbackFunc func(this js.Ref, id js.String) js.Ref

func (fn OnConfigRemovedEventCallbackFunc) Register() js.Func[func(id js.String)] {
	return js.RegisterCallback[func(id js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnConfigRemovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnConfigRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String) js.Ref
	Arg T
}

func (cb *OnConfigRemovedEventCallback[T]) Register() js.Func[func(id js.String)] {
	return js.RegisterCallback[func(id js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnConfigRemovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnConfigRemoved returns true if the function "WEBEXT.vpnProvider.onConfigRemoved.addListener" exists.
func HasFuncOnConfigRemoved() bool {
	return js.True == bindings.HasFuncOnConfigRemoved()
}

// FuncOnConfigRemoved returns the function "WEBEXT.vpnProvider.onConfigRemoved.addListener".
func FuncOnConfigRemoved() (fn js.Func[func(callback js.Func[func(id js.String)])]) {
	bindings.FuncOnConfigRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnConfigRemoved calls the function "WEBEXT.vpnProvider.onConfigRemoved.addListener" directly.
func OnConfigRemoved(callback js.Func[func(id js.String)]) (ret js.Void) {
	bindings.CallOnConfigRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnConfigRemoved calls the function "WEBEXT.vpnProvider.onConfigRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnConfigRemoved(callback js.Func[func(id js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnConfigRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffConfigRemoved returns true if the function "WEBEXT.vpnProvider.onConfigRemoved.removeListener" exists.
func HasFuncOffConfigRemoved() bool {
	return js.True == bindings.HasFuncOffConfigRemoved()
}

// FuncOffConfigRemoved returns the function "WEBEXT.vpnProvider.onConfigRemoved.removeListener".
func FuncOffConfigRemoved() (fn js.Func[func(callback js.Func[func(id js.String)])]) {
	bindings.FuncOffConfigRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffConfigRemoved calls the function "WEBEXT.vpnProvider.onConfigRemoved.removeListener" directly.
func OffConfigRemoved(callback js.Func[func(id js.String)]) (ret js.Void) {
	bindings.CallOffConfigRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffConfigRemoved calls the function "WEBEXT.vpnProvider.onConfigRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffConfigRemoved(callback js.Func[func(id js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffConfigRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnConfigRemoved returns true if the function "WEBEXT.vpnProvider.onConfigRemoved.hasListener" exists.
func HasFuncHasOnConfigRemoved() bool {
	return js.True == bindings.HasFuncHasOnConfigRemoved()
}

// FuncHasOnConfigRemoved returns the function "WEBEXT.vpnProvider.onConfigRemoved.hasListener".
func FuncHasOnConfigRemoved() (fn js.Func[func(callback js.Func[func(id js.String)]) bool]) {
	bindings.FuncHasOnConfigRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnConfigRemoved calls the function "WEBEXT.vpnProvider.onConfigRemoved.hasListener" directly.
func HasOnConfigRemoved(callback js.Func[func(id js.String)]) (ret bool) {
	bindings.CallHasOnConfigRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnConfigRemoved calls the function "WEBEXT.vpnProvider.onConfigRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnConfigRemoved(callback js.Func[func(id js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnConfigRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPacketReceivedEventCallbackFunc func(this js.Ref, data js.ArrayBuffer) js.Ref

func (fn OnPacketReceivedEventCallbackFunc) Register() js.Func[func(data js.ArrayBuffer)] {
	return js.RegisterCallback[func(data js.ArrayBuffer)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPacketReceivedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPacketReceivedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, data js.ArrayBuffer) js.Ref
	Arg T
}

func (cb *OnPacketReceivedEventCallback[T]) Register() js.Func[func(data js.ArrayBuffer)] {
	return js.RegisterCallback[func(data js.ArrayBuffer)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPacketReceivedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.ArrayBuffer{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPacketReceived returns true if the function "WEBEXT.vpnProvider.onPacketReceived.addListener" exists.
func HasFuncOnPacketReceived() bool {
	return js.True == bindings.HasFuncOnPacketReceived()
}

// FuncOnPacketReceived returns the function "WEBEXT.vpnProvider.onPacketReceived.addListener".
func FuncOnPacketReceived() (fn js.Func[func(callback js.Func[func(data js.ArrayBuffer)])]) {
	bindings.FuncOnPacketReceived(
		js.Pointer(&fn),
	)
	return
}

// OnPacketReceived calls the function "WEBEXT.vpnProvider.onPacketReceived.addListener" directly.
func OnPacketReceived(callback js.Func[func(data js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallOnPacketReceived(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPacketReceived calls the function "WEBEXT.vpnProvider.onPacketReceived.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPacketReceived(callback js.Func[func(data js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPacketReceived(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPacketReceived returns true if the function "WEBEXT.vpnProvider.onPacketReceived.removeListener" exists.
func HasFuncOffPacketReceived() bool {
	return js.True == bindings.HasFuncOffPacketReceived()
}

// FuncOffPacketReceived returns the function "WEBEXT.vpnProvider.onPacketReceived.removeListener".
func FuncOffPacketReceived() (fn js.Func[func(callback js.Func[func(data js.ArrayBuffer)])]) {
	bindings.FuncOffPacketReceived(
		js.Pointer(&fn),
	)
	return
}

// OffPacketReceived calls the function "WEBEXT.vpnProvider.onPacketReceived.removeListener" directly.
func OffPacketReceived(callback js.Func[func(data js.ArrayBuffer)]) (ret js.Void) {
	bindings.CallOffPacketReceived(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPacketReceived calls the function "WEBEXT.vpnProvider.onPacketReceived.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPacketReceived(callback js.Func[func(data js.ArrayBuffer)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPacketReceived(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPacketReceived returns true if the function "WEBEXT.vpnProvider.onPacketReceived.hasListener" exists.
func HasFuncHasOnPacketReceived() bool {
	return js.True == bindings.HasFuncHasOnPacketReceived()
}

// FuncHasOnPacketReceived returns the function "WEBEXT.vpnProvider.onPacketReceived.hasListener".
func FuncHasOnPacketReceived() (fn js.Func[func(callback js.Func[func(data js.ArrayBuffer)]) bool]) {
	bindings.FuncHasOnPacketReceived(
		js.Pointer(&fn),
	)
	return
}

// HasOnPacketReceived calls the function "WEBEXT.vpnProvider.onPacketReceived.hasListener" directly.
func HasOnPacketReceived(callback js.Func[func(data js.ArrayBuffer)]) (ret bool) {
	bindings.CallHasOnPacketReceived(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPacketReceived calls the function "WEBEXT.vpnProvider.onPacketReceived.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPacketReceived(callback js.Func[func(data js.ArrayBuffer)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPacketReceived(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPlatformMessageEventCallbackFunc func(this js.Ref, id js.String, message PlatformMessage, err js.String) js.Ref

func (fn OnPlatformMessageEventCallbackFunc) Register() js.Func[func(id js.String, message PlatformMessage, err js.String)] {
	return js.RegisterCallback[func(id js.String, message PlatformMessage, err js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPlatformMessageEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		PlatformMessage(0).FromRef(args[1+1]),
		js.String{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnPlatformMessageEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, id js.String, message PlatformMessage, err js.String) js.Ref
	Arg T
}

func (cb *OnPlatformMessageEventCallback[T]) Register() js.Func[func(id js.String, message PlatformMessage, err js.String)] {
	return js.RegisterCallback[func(id js.String, message PlatformMessage, err js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPlatformMessageEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		PlatformMessage(0).FromRef(args[1+1]),
		js.String{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPlatformMessage returns true if the function "WEBEXT.vpnProvider.onPlatformMessage.addListener" exists.
func HasFuncOnPlatformMessage() bool {
	return js.True == bindings.HasFuncOnPlatformMessage()
}

// FuncOnPlatformMessage returns the function "WEBEXT.vpnProvider.onPlatformMessage.addListener".
func FuncOnPlatformMessage() (fn js.Func[func(callback js.Func[func(id js.String, message PlatformMessage, err js.String)])]) {
	bindings.FuncOnPlatformMessage(
		js.Pointer(&fn),
	)
	return
}

// OnPlatformMessage calls the function "WEBEXT.vpnProvider.onPlatformMessage.addListener" directly.
func OnPlatformMessage(callback js.Func[func(id js.String, message PlatformMessage, err js.String)]) (ret js.Void) {
	bindings.CallOnPlatformMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPlatformMessage calls the function "WEBEXT.vpnProvider.onPlatformMessage.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPlatformMessage(callback js.Func[func(id js.String, message PlatformMessage, err js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPlatformMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPlatformMessage returns true if the function "WEBEXT.vpnProvider.onPlatformMessage.removeListener" exists.
func HasFuncOffPlatformMessage() bool {
	return js.True == bindings.HasFuncOffPlatformMessage()
}

// FuncOffPlatformMessage returns the function "WEBEXT.vpnProvider.onPlatformMessage.removeListener".
func FuncOffPlatformMessage() (fn js.Func[func(callback js.Func[func(id js.String, message PlatformMessage, err js.String)])]) {
	bindings.FuncOffPlatformMessage(
		js.Pointer(&fn),
	)
	return
}

// OffPlatformMessage calls the function "WEBEXT.vpnProvider.onPlatformMessage.removeListener" directly.
func OffPlatformMessage(callback js.Func[func(id js.String, message PlatformMessage, err js.String)]) (ret js.Void) {
	bindings.CallOffPlatformMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPlatformMessage calls the function "WEBEXT.vpnProvider.onPlatformMessage.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPlatformMessage(callback js.Func[func(id js.String, message PlatformMessage, err js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPlatformMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPlatformMessage returns true if the function "WEBEXT.vpnProvider.onPlatformMessage.hasListener" exists.
func HasFuncHasOnPlatformMessage() bool {
	return js.True == bindings.HasFuncHasOnPlatformMessage()
}

// FuncHasOnPlatformMessage returns the function "WEBEXT.vpnProvider.onPlatformMessage.hasListener".
func FuncHasOnPlatformMessage() (fn js.Func[func(callback js.Func[func(id js.String, message PlatformMessage, err js.String)]) bool]) {
	bindings.FuncHasOnPlatformMessage(
		js.Pointer(&fn),
	)
	return
}

// HasOnPlatformMessage calls the function "WEBEXT.vpnProvider.onPlatformMessage.hasListener" directly.
func HasOnPlatformMessage(callback js.Func[func(id js.String, message PlatformMessage, err js.String)]) (ret bool) {
	bindings.CallHasOnPlatformMessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPlatformMessage calls the function "WEBEXT.vpnProvider.onPlatformMessage.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPlatformMessage(callback js.Func[func(id js.String, message PlatformMessage, err js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPlatformMessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnUIEventEventCallbackFunc func(this js.Ref, event UIEvent, id js.String) js.Ref

func (fn OnUIEventEventCallbackFunc) Register() js.Func[func(event UIEvent, id js.String)] {
	return js.RegisterCallback[func(event UIEvent, id js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnUIEventEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		UIEvent(0).FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnUIEventEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, event UIEvent, id js.String) js.Ref
	Arg T
}

func (cb *OnUIEventEventCallback[T]) Register() js.Func[func(event UIEvent, id js.String)] {
	return js.RegisterCallback[func(event UIEvent, id js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnUIEventEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		UIEvent(0).FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnUIEvent returns true if the function "WEBEXT.vpnProvider.onUIEvent.addListener" exists.
func HasFuncOnUIEvent() bool {
	return js.True == bindings.HasFuncOnUIEvent()
}

// FuncOnUIEvent returns the function "WEBEXT.vpnProvider.onUIEvent.addListener".
func FuncOnUIEvent() (fn js.Func[func(callback js.Func[func(event UIEvent, id js.String)])]) {
	bindings.FuncOnUIEvent(
		js.Pointer(&fn),
	)
	return
}

// OnUIEvent calls the function "WEBEXT.vpnProvider.onUIEvent.addListener" directly.
func OnUIEvent(callback js.Func[func(event UIEvent, id js.String)]) (ret js.Void) {
	bindings.CallOnUIEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUIEvent calls the function "WEBEXT.vpnProvider.onUIEvent.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUIEvent(callback js.Func[func(event UIEvent, id js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUIEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUIEvent returns true if the function "WEBEXT.vpnProvider.onUIEvent.removeListener" exists.
func HasFuncOffUIEvent() bool {
	return js.True == bindings.HasFuncOffUIEvent()
}

// FuncOffUIEvent returns the function "WEBEXT.vpnProvider.onUIEvent.removeListener".
func FuncOffUIEvent() (fn js.Func[func(callback js.Func[func(event UIEvent, id js.String)])]) {
	bindings.FuncOffUIEvent(
		js.Pointer(&fn),
	)
	return
}

// OffUIEvent calls the function "WEBEXT.vpnProvider.onUIEvent.removeListener" directly.
func OffUIEvent(callback js.Func[func(event UIEvent, id js.String)]) (ret js.Void) {
	bindings.CallOffUIEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUIEvent calls the function "WEBEXT.vpnProvider.onUIEvent.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUIEvent(callback js.Func[func(event UIEvent, id js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUIEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUIEvent returns true if the function "WEBEXT.vpnProvider.onUIEvent.hasListener" exists.
func HasFuncHasOnUIEvent() bool {
	return js.True == bindings.HasFuncHasOnUIEvent()
}

// FuncHasOnUIEvent returns the function "WEBEXT.vpnProvider.onUIEvent.hasListener".
func FuncHasOnUIEvent() (fn js.Func[func(callback js.Func[func(event UIEvent, id js.String)]) bool]) {
	bindings.FuncHasOnUIEvent(
		js.Pointer(&fn),
	)
	return
}

// HasOnUIEvent calls the function "WEBEXT.vpnProvider.onUIEvent.hasListener" directly.
func HasOnUIEvent(callback js.Func[func(event UIEvent, id js.String)]) (ret bool) {
	bindings.CallHasOnUIEvent(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUIEvent calls the function "WEBEXT.vpnProvider.onUIEvent.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUIEvent(callback js.Func[func(event UIEvent, id js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUIEvent(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSendPacket returns true if the function "WEBEXT.vpnProvider.sendPacket" exists.
func HasFuncSendPacket() bool {
	return js.True == bindings.HasFuncSendPacket()
}

// FuncSendPacket returns the function "WEBEXT.vpnProvider.sendPacket".
func FuncSendPacket() (fn js.Func[func(data js.ArrayBuffer) js.Promise[js.Void]]) {
	bindings.FuncSendPacket(
		js.Pointer(&fn),
	)
	return
}

// SendPacket calls the function "WEBEXT.vpnProvider.sendPacket" directly.
func SendPacket(data js.ArrayBuffer) (ret js.Promise[js.Void]) {
	bindings.CallSendPacket(
		js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySendPacket calls the function "WEBEXT.vpnProvider.sendPacket"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendPacket(data js.ArrayBuffer) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendPacket(
		js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

// HasFuncSetParameters returns true if the function "WEBEXT.vpnProvider.setParameters" exists.
func HasFuncSetParameters() bool {
	return js.True == bindings.HasFuncSetParameters()
}

// FuncSetParameters returns the function "WEBEXT.vpnProvider.setParameters".
func FuncSetParameters() (fn js.Func[func(parameters Parameters) js.Promise[js.Void]]) {
	bindings.FuncSetParameters(
		js.Pointer(&fn),
	)
	return
}

// SetParameters calls the function "WEBEXT.vpnProvider.setParameters" directly.
func SetParameters(parameters Parameters) (ret js.Promise[js.Void]) {
	bindings.CallSetParameters(
		js.Pointer(&ret),
		js.Pointer(&parameters),
	)

	return
}

// TrySetParameters calls the function "WEBEXT.vpnProvider.setParameters"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetParameters(parameters Parameters) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetParameters(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&parameters),
	)

	return
}

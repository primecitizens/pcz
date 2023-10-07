// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package proxy

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/proxy/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/types"
)

type Mode uint32

const (
	_ Mode = iota

	Mode_DIRECT
	Mode_AUTO_DETECT
	Mode_PAC_SCRIPT
	Mode_FIXED_SERVERS
	Mode_SYSTEM
)

func (Mode) FromRef(str js.Ref) Mode {
	return Mode(bindings.ConstOfMode(str))
}

func (x Mode) String() (string, bool) {
	switch x {
	case Mode_DIRECT:
		return "direct", true
	case Mode_AUTO_DETECT:
		return "auto_detect", true
	case Mode_PAC_SCRIPT:
		return "pac_script", true
	case Mode_FIXED_SERVERS:
		return "fixed_servers", true
	case Mode_SYSTEM:
		return "system", true
	default:
		return "", false
	}
}

type OnProxyErrorArgDetails struct {
	// Details is "OnProxyErrorArgDetails.details"
	//
	// Required
	Details js.String
	// Error is "OnProxyErrorArgDetails.error"
	//
	// Required
	Error js.String
	// Fatal is "OnProxyErrorArgDetails.fatal"
	//
	// Required
	Fatal bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnProxyErrorArgDetails with all fields set.
func (p OnProxyErrorArgDetails) FromRef(ref js.Ref) OnProxyErrorArgDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnProxyErrorArgDetails in the application heap.
func (p OnProxyErrorArgDetails) New() js.Ref {
	return bindings.OnProxyErrorArgDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnProxyErrorArgDetails) UpdateFrom(ref js.Ref) {
	bindings.OnProxyErrorArgDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnProxyErrorArgDetails) Update(ref js.Ref) {
	bindings.OnProxyErrorArgDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnProxyErrorArgDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Details.Ref(),
		p.Error.Ref(),
	)
	p.Details = p.Details.FromRef(js.Undefined)
	p.Error = p.Error.FromRef(js.Undefined)
}

type PacScript struct {
	// Data is "PacScript.data"
	//
	// Optional
	Data js.String
	// Mandatory is "PacScript.mandatory"
	//
	// Optional
	//
	// NOTE: FFI_USE_Mandatory MUST be set to true to make this field effective.
	Mandatory bool
	// Url is "PacScript.url"
	//
	// Optional
	Url js.String

	FFI_USE_Mandatory bool // for Mandatory.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PacScript with all fields set.
func (p PacScript) FromRef(ref js.Ref) PacScript {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PacScript in the application heap.
func (p PacScript) New() js.Ref {
	return bindings.PacScriptJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PacScript) UpdateFrom(ref js.Ref) {
	bindings.PacScriptJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PacScript) Update(ref js.Ref) {
	bindings.PacScriptJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PacScript) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
		p.Url.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type Scheme uint32

const (
	_ Scheme = iota

	Scheme_HTTP
	Scheme_HTTPS
	Scheme_QUIC
	Scheme_SOCKS4
	Scheme_SOCKS5
)

func (Scheme) FromRef(str js.Ref) Scheme {
	return Scheme(bindings.ConstOfScheme(str))
}

func (x Scheme) String() (string, bool) {
	switch x {
	case Scheme_HTTP:
		return "http", true
	case Scheme_HTTPS:
		return "https", true
	case Scheme_QUIC:
		return "quic", true
	case Scheme_SOCKS4:
		return "socks4", true
	case Scheme_SOCKS5:
		return "socks5", true
	default:
		return "", false
	}
}

type ProxyServer struct {
	// Host is "ProxyServer.host"
	//
	// Required
	Host js.String
	// Port is "ProxyServer.port"
	//
	// Optional
	//
	// NOTE: FFI_USE_Port MUST be set to true to make this field effective.
	Port int64
	// Scheme is "ProxyServer.scheme"
	//
	// Optional
	Scheme Scheme

	FFI_USE_Port bool // for Port.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProxyServer with all fields set.
func (p ProxyServer) FromRef(ref js.Ref) ProxyServer {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProxyServer in the application heap.
func (p ProxyServer) New() js.Ref {
	return bindings.ProxyServerJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProxyServer) UpdateFrom(ref js.Ref) {
	bindings.ProxyServerJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProxyServer) Update(ref js.Ref) {
	bindings.ProxyServerJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProxyServer) FreeMembers(recursive bool) {
	js.Free(
		p.Host.Ref(),
	)
	p.Host = p.Host.FromRef(js.Undefined)
}

type ProxyRules struct {
	// BypassList is "ProxyRules.bypassList"
	//
	// Optional
	BypassList js.Array[js.String]
	// FallbackProxy is "ProxyRules.fallbackProxy"
	//
	// Optional
	//
	// NOTE: FallbackProxy.FFI_USE MUST be set to true to get FallbackProxy used.
	FallbackProxy ProxyServer
	// ProxyForFtp is "ProxyRules.proxyForFtp"
	//
	// Optional
	//
	// NOTE: ProxyForFtp.FFI_USE MUST be set to true to get ProxyForFtp used.
	ProxyForFtp ProxyServer
	// ProxyForHttp is "ProxyRules.proxyForHttp"
	//
	// Optional
	//
	// NOTE: ProxyForHttp.FFI_USE MUST be set to true to get ProxyForHttp used.
	ProxyForHttp ProxyServer
	// ProxyForHttps is "ProxyRules.proxyForHttps"
	//
	// Optional
	//
	// NOTE: ProxyForHttps.FFI_USE MUST be set to true to get ProxyForHttps used.
	ProxyForHttps ProxyServer
	// SingleProxy is "ProxyRules.singleProxy"
	//
	// Optional
	//
	// NOTE: SingleProxy.FFI_USE MUST be set to true to get SingleProxy used.
	SingleProxy ProxyServer

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProxyRules with all fields set.
func (p ProxyRules) FromRef(ref js.Ref) ProxyRules {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProxyRules in the application heap.
func (p ProxyRules) New() js.Ref {
	return bindings.ProxyRulesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProxyRules) UpdateFrom(ref js.Ref) {
	bindings.ProxyRulesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProxyRules) Update(ref js.Ref) {
	bindings.ProxyRulesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProxyRules) FreeMembers(recursive bool) {
	js.Free(
		p.BypassList.Ref(),
	)
	p.BypassList = p.BypassList.FromRef(js.Undefined)
	if recursive {
		p.FallbackProxy.FreeMembers(true)
		p.ProxyForFtp.FreeMembers(true)
		p.ProxyForHttp.FreeMembers(true)
		p.ProxyForHttps.FreeMembers(true)
		p.SingleProxy.FreeMembers(true)
	}
}

type ProxyConfig struct {
	// Mode is "ProxyConfig.mode"
	//
	// Required
	Mode Mode
	// PacScript is "ProxyConfig.pacScript"
	//
	// Optional
	//
	// NOTE: PacScript.FFI_USE MUST be set to true to get PacScript used.
	PacScript PacScript
	// Rules is "ProxyConfig.rules"
	//
	// Optional
	//
	// NOTE: Rules.FFI_USE MUST be set to true to get Rules used.
	Rules ProxyRules

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ProxyConfig with all fields set.
func (p ProxyConfig) FromRef(ref js.Ref) ProxyConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ProxyConfig in the application heap.
func (p ProxyConfig) New() js.Ref {
	return bindings.ProxyConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ProxyConfig) UpdateFrom(ref js.Ref) {
	bindings.ProxyConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ProxyConfig) Update(ref js.Ref) {
	bindings.ProxyConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ProxyConfig) FreeMembers(recursive bool) {
	if recursive {
		p.PacScript.FreeMembers(true)
		p.Rules.FreeMembers(true)
	}
}

type OnProxyErrorEventCallbackFunc func(this js.Ref, details *OnProxyErrorArgDetails) js.Ref

func (fn OnProxyErrorEventCallbackFunc) Register() js.Func[func(details *OnProxyErrorArgDetails)] {
	return js.RegisterCallback[func(details *OnProxyErrorArgDetails)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnProxyErrorEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnProxyErrorArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnProxyErrorEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, details *OnProxyErrorArgDetails) js.Ref
	Arg T
}

func (cb *OnProxyErrorEventCallback[T]) Register() js.Func[func(details *OnProxyErrorArgDetails)] {
	return js.RegisterCallback[func(details *OnProxyErrorArgDetails)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnProxyErrorEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnProxyErrorArgDetails
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnProxyError returns true if the function "WEBEXT.proxy.onProxyError.addListener" exists.
func HasFuncOnProxyError() bool {
	return js.True == bindings.HasFuncOnProxyError()
}

// FuncOnProxyError returns the function "WEBEXT.proxy.onProxyError.addListener".
func FuncOnProxyError() (fn js.Func[func(callback js.Func[func(details *OnProxyErrorArgDetails)])]) {
	bindings.FuncOnProxyError(
		js.Pointer(&fn),
	)
	return
}

// OnProxyError calls the function "WEBEXT.proxy.onProxyError.addListener" directly.
func OnProxyError(callback js.Func[func(details *OnProxyErrorArgDetails)]) (ret js.Void) {
	bindings.CallOnProxyError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnProxyError calls the function "WEBEXT.proxy.onProxyError.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnProxyError(callback js.Func[func(details *OnProxyErrorArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnProxyError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffProxyError returns true if the function "WEBEXT.proxy.onProxyError.removeListener" exists.
func HasFuncOffProxyError() bool {
	return js.True == bindings.HasFuncOffProxyError()
}

// FuncOffProxyError returns the function "WEBEXT.proxy.onProxyError.removeListener".
func FuncOffProxyError() (fn js.Func[func(callback js.Func[func(details *OnProxyErrorArgDetails)])]) {
	bindings.FuncOffProxyError(
		js.Pointer(&fn),
	)
	return
}

// OffProxyError calls the function "WEBEXT.proxy.onProxyError.removeListener" directly.
func OffProxyError(callback js.Func[func(details *OnProxyErrorArgDetails)]) (ret js.Void) {
	bindings.CallOffProxyError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffProxyError calls the function "WEBEXT.proxy.onProxyError.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffProxyError(callback js.Func[func(details *OnProxyErrorArgDetails)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffProxyError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnProxyError returns true if the function "WEBEXT.proxy.onProxyError.hasListener" exists.
func HasFuncHasOnProxyError() bool {
	return js.True == bindings.HasFuncHasOnProxyError()
}

// FuncHasOnProxyError returns the function "WEBEXT.proxy.onProxyError.hasListener".
func FuncHasOnProxyError() (fn js.Func[func(callback js.Func[func(details *OnProxyErrorArgDetails)]) bool]) {
	bindings.FuncHasOnProxyError(
		js.Pointer(&fn),
	)
	return
}

// HasOnProxyError calls the function "WEBEXT.proxy.onProxyError.hasListener" directly.
func HasOnProxyError(callback js.Func[func(details *OnProxyErrorArgDetails)]) (ret bool) {
	bindings.CallHasOnProxyError(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnProxyError calls the function "WEBEXT.proxy.onProxyError.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnProxyError(callback js.Func[func(details *OnProxyErrorArgDetails)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnProxyError(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// Settings returns the value of property "WEBEXT.proxy.settings".
//
// The returned bool will be false if there is no such property.
func Settings() (ret types.ChromeSetting, ok bool) {
	ok = js.True == bindings.GetSettings(
		js.Pointer(&ret),
	)

	return
}

// SetSettings sets the value of property "WEBEXT.proxy.settings" to val.
//
// It returns false if the property cannot be set.
func SetSettings(val types.ChromeSetting) bool {
	return js.True == bindings.SetSettings(
		val.Ref())
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package safebrowsingprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/safebrowsingprivate/bindings"
)

type DangerousDownloadInfo struct {
	// Url is "DangerousDownloadInfo.url"
	//
	// Optional
	Url js.String
	// FileName is "DangerousDownloadInfo.fileName"
	//
	// Optional
	FileName js.String
	// DownloadDigestSha256 is "DangerousDownloadInfo.downloadDigestSha256"
	//
	// Optional
	DownloadDigestSha256 js.String
	// UserName is "DangerousDownloadInfo.userName"
	//
	// Optional
	UserName js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DangerousDownloadInfo with all fields set.
func (p DangerousDownloadInfo) FromRef(ref js.Ref) DangerousDownloadInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DangerousDownloadInfo in the application heap.
func (p DangerousDownloadInfo) New() js.Ref {
	return bindings.DangerousDownloadInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DangerousDownloadInfo) UpdateFrom(ref js.Ref) {
	bindings.DangerousDownloadInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DangerousDownloadInfo) Update(ref js.Ref) {
	bindings.DangerousDownloadInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DangerousDownloadInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.FileName.Ref(),
		p.DownloadDigestSha256.Ref(),
		p.UserName.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.FileName = p.FileName.FromRef(js.Undefined)
	p.DownloadDigestSha256 = p.DownloadDigestSha256.FromRef(js.Undefined)
	p.UserName = p.UserName.FromRef(js.Undefined)
}

type GetReferrerChainCallbackFunc func(this js.Ref, entries js.Array[ReferrerChainEntry]) js.Ref

func (fn GetReferrerChainCallbackFunc) Register() js.Func[func(entries js.Array[ReferrerChainEntry])] {
	return js.RegisterCallback[func(entries js.Array[ReferrerChainEntry])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetReferrerChainCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[ReferrerChainEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetReferrerChainCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[ReferrerChainEntry]) js.Ref
	Arg T
}

func (cb *GetReferrerChainCallback[T]) Register() js.Func[func(entries js.Array[ReferrerChainEntry])] {
	return js.RegisterCallback[func(entries js.Array[ReferrerChainEntry])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetReferrerChainCallback[T]) DispatchCallback(
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

		js.Array[ReferrerChainEntry]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type URLType uint32

const (
	_ URLType = iota

	URLType_EVENT_URL
	URLType_LANDING_PAGE
	URLType_LANDING_REFERRER
	URLType_CLIENT_REDIRECT
	URLType_RECENT_NAVIGATION
	URLType_REFERRER
)

func (URLType) FromRef(str js.Ref) URLType {
	return URLType(bindings.ConstOfURLType(str))
}

func (x URLType) String() (string, bool) {
	switch x {
	case URLType_EVENT_URL:
		return "EVENT_URL", true
	case URLType_LANDING_PAGE:
		return "LANDING_PAGE", true
	case URLType_LANDING_REFERRER:
		return "LANDING_REFERRER", true
	case URLType_CLIENT_REDIRECT:
		return "CLIENT_REDIRECT", true
	case URLType_RECENT_NAVIGATION:
		return "RECENT_NAVIGATION", true
	case URLType_REFERRER:
		return "REFERRER", true
	default:
		return "", false
	}
}

type ServerRedirect struct {
	// Url is "ServerRedirect.url"
	//
	// Optional
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ServerRedirect with all fields set.
func (p ServerRedirect) FromRef(ref js.Ref) ServerRedirect {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ServerRedirect in the application heap.
func (p ServerRedirect) New() js.Ref {
	return bindings.ServerRedirectJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ServerRedirect) UpdateFrom(ref js.Ref) {
	bindings.ServerRedirectJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ServerRedirect) Update(ref js.Ref) {
	bindings.ServerRedirectJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ServerRedirect) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type NavigationInitiation uint32

const (
	_ NavigationInitiation = iota

	NavigationInitiation_BROWSER_INITIATED
	NavigationInitiation_RENDERER_INITIATED_WITHOUT_USER_GESTURE
	NavigationInitiation_RENDERER_INITIATED_WITH_USER_GESTURE
	NavigationInitiation_COPY_PASTE_USER_INITIATED
	NavigationInitiation_NOTIFICATION_INITIATED
)

func (NavigationInitiation) FromRef(str js.Ref) NavigationInitiation {
	return NavigationInitiation(bindings.ConstOfNavigationInitiation(str))
}

func (x NavigationInitiation) String() (string, bool) {
	switch x {
	case NavigationInitiation_BROWSER_INITIATED:
		return "BROWSER_INITIATED", true
	case NavigationInitiation_RENDERER_INITIATED_WITHOUT_USER_GESTURE:
		return "RENDERER_INITIATED_WITHOUT_USER_GESTURE", true
	case NavigationInitiation_RENDERER_INITIATED_WITH_USER_GESTURE:
		return "RENDERER_INITIATED_WITH_USER_GESTURE", true
	case NavigationInitiation_COPY_PASTE_USER_INITIATED:
		return "COPY_PASTE_USER_INITIATED", true
	case NavigationInitiation_NOTIFICATION_INITIATED:
		return "NOTIFICATION_INITIATED", true
	default:
		return "", false
	}
}

type ReferrerChainEntry struct {
	// Url is "ReferrerChainEntry.url"
	//
	// Optional
	Url js.String
	// MainFrameUrl is "ReferrerChainEntry.mainFrameUrl"
	//
	// Optional
	MainFrameUrl js.String
	// UrlType is "ReferrerChainEntry.urlType"
	//
	// Optional
	UrlType URLType
	// IpAddresses is "ReferrerChainEntry.ipAddresses"
	//
	// Optional
	IpAddresses js.Array[js.String]
	// ReferrerUrl is "ReferrerChainEntry.referrerUrl"
	//
	// Optional
	ReferrerUrl js.String
	// ReferrerMainFrameUrl is "ReferrerChainEntry.referrerMainFrameUrl"
	//
	// Optional
	ReferrerMainFrameUrl js.String
	// IsRetargeting is "ReferrerChainEntry.isRetargeting"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsRetargeting MUST be set to true to make this field effective.
	IsRetargeting bool
	// NavigationTimeMs is "ReferrerChainEntry.navigationTimeMs"
	//
	// Optional
	//
	// NOTE: FFI_USE_NavigationTimeMs MUST be set to true to make this field effective.
	NavigationTimeMs float64
	// ServerRedirectChain is "ReferrerChainEntry.serverRedirectChain"
	//
	// Optional
	ServerRedirectChain js.Array[ServerRedirect]
	// NavigationInitiation is "ReferrerChainEntry.navigationInitiation"
	//
	// Optional
	NavigationInitiation NavigationInitiation
	// MaybeLaunchedByExternalApp is "ReferrerChainEntry.maybeLaunchedByExternalApp"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaybeLaunchedByExternalApp MUST be set to true to make this field effective.
	MaybeLaunchedByExternalApp bool
	// IsSubframeUrlRemoved is "ReferrerChainEntry.isSubframeUrlRemoved"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsSubframeUrlRemoved MUST be set to true to make this field effective.
	IsSubframeUrlRemoved bool
	// IsSubframeReferrerUrlRemoved is "ReferrerChainEntry.isSubframeReferrerUrlRemoved"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsSubframeReferrerUrlRemoved MUST be set to true to make this field effective.
	IsSubframeReferrerUrlRemoved bool
	// IsUrlRemovedByPolicy is "ReferrerChainEntry.isUrlRemovedByPolicy"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsUrlRemovedByPolicy MUST be set to true to make this field effective.
	IsUrlRemovedByPolicy bool

	FFI_USE_IsRetargeting                bool // for IsRetargeting.
	FFI_USE_NavigationTimeMs             bool // for NavigationTimeMs.
	FFI_USE_MaybeLaunchedByExternalApp   bool // for MaybeLaunchedByExternalApp.
	FFI_USE_IsSubframeUrlRemoved         bool // for IsSubframeUrlRemoved.
	FFI_USE_IsSubframeReferrerUrlRemoved bool // for IsSubframeReferrerUrlRemoved.
	FFI_USE_IsUrlRemovedByPolicy         bool // for IsUrlRemovedByPolicy.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReferrerChainEntry with all fields set.
func (p ReferrerChainEntry) FromRef(ref js.Ref) ReferrerChainEntry {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReferrerChainEntry in the application heap.
func (p ReferrerChainEntry) New() js.Ref {
	return bindings.ReferrerChainEntryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReferrerChainEntry) UpdateFrom(ref js.Ref) {
	bindings.ReferrerChainEntryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReferrerChainEntry) Update(ref js.Ref) {
	bindings.ReferrerChainEntryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReferrerChainEntry) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.MainFrameUrl.Ref(),
		p.IpAddresses.Ref(),
		p.ReferrerUrl.Ref(),
		p.ReferrerMainFrameUrl.Ref(),
		p.ServerRedirectChain.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.MainFrameUrl = p.MainFrameUrl.FromRef(js.Undefined)
	p.IpAddresses = p.IpAddresses.FromRef(js.Undefined)
	p.ReferrerUrl = p.ReferrerUrl.FromRef(js.Undefined)
	p.ReferrerMainFrameUrl = p.ReferrerMainFrameUrl.FromRef(js.Undefined)
	p.ServerRedirectChain = p.ServerRedirectChain.FromRef(js.Undefined)
}

type InterstitialInfo struct {
	// Url is "InterstitialInfo.url"
	//
	// Optional
	Url js.String
	// Reason is "InterstitialInfo.reason"
	//
	// Optional
	Reason js.String
	// NetErrorCode is "InterstitialInfo.netErrorCode"
	//
	// Optional
	NetErrorCode js.String
	// UserName is "InterstitialInfo.userName"
	//
	// Optional
	UserName js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InterstitialInfo with all fields set.
func (p InterstitialInfo) FromRef(ref js.Ref) InterstitialInfo {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InterstitialInfo in the application heap.
func (p InterstitialInfo) New() js.Ref {
	return bindings.InterstitialInfoJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InterstitialInfo) UpdateFrom(ref js.Ref) {
	bindings.InterstitialInfoJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InterstitialInfo) Update(ref js.Ref) {
	bindings.InterstitialInfoJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InterstitialInfo) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.Reason.Ref(),
		p.NetErrorCode.Ref(),
		p.UserName.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Reason = p.Reason.FromRef(js.Undefined)
	p.NetErrorCode = p.NetErrorCode.FromRef(js.Undefined)
	p.UserName = p.UserName.FromRef(js.Undefined)
}

type PolicySpecifiedPasswordReuse struct {
	// Url is "PolicySpecifiedPasswordReuse.url"
	//
	// Optional
	Url js.String
	// UserName is "PolicySpecifiedPasswordReuse.userName"
	//
	// Optional
	UserName js.String
	// IsPhishingUrl is "PolicySpecifiedPasswordReuse.isPhishingUrl"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsPhishingUrl MUST be set to true to make this field effective.
	IsPhishingUrl bool

	FFI_USE_IsPhishingUrl bool // for IsPhishingUrl.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PolicySpecifiedPasswordReuse with all fields set.
func (p PolicySpecifiedPasswordReuse) FromRef(ref js.Ref) PolicySpecifiedPasswordReuse {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PolicySpecifiedPasswordReuse in the application heap.
func (p PolicySpecifiedPasswordReuse) New() js.Ref {
	return bindings.PolicySpecifiedPasswordReuseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PolicySpecifiedPasswordReuse) UpdateFrom(ref js.Ref) {
	bindings.PolicySpecifiedPasswordReuseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PolicySpecifiedPasswordReuse) Update(ref js.Ref) {
	bindings.PolicySpecifiedPasswordReuseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PolicySpecifiedPasswordReuse) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
		p.UserName.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
	p.UserName = p.UserName.FromRef(js.Undefined)
}

// HasFuncGetReferrerChain returns true if the function "WEBEXT.safeBrowsingPrivate.getReferrerChain" exists.
func HasFuncGetReferrerChain() bool {
	return js.True == bindings.HasFuncGetReferrerChain()
}

// FuncGetReferrerChain returns the function "WEBEXT.safeBrowsingPrivate.getReferrerChain".
func FuncGetReferrerChain() (fn js.Func[func(tabId int32) js.Promise[js.Array[ReferrerChainEntry]]]) {
	bindings.FuncGetReferrerChain(
		js.Pointer(&fn),
	)
	return
}

// GetReferrerChain calls the function "WEBEXT.safeBrowsingPrivate.getReferrerChain" directly.
func GetReferrerChain(tabId int32) (ret js.Promise[js.Array[ReferrerChainEntry]]) {
	bindings.CallGetReferrerChain(
		js.Pointer(&ret),
		int32(tabId),
	)

	return
}

// TryGetReferrerChain calls the function "WEBEXT.safeBrowsingPrivate.getReferrerChain"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetReferrerChain(tabId int32) (ret js.Promise[js.Array[ReferrerChainEntry]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetReferrerChain(
		js.Pointer(&ret), js.Pointer(&exception),
		int32(tabId),
	)

	return
}

type OnDangerousDownloadOpenedEventCallbackFunc func(this js.Ref, dict *DangerousDownloadInfo) js.Ref

func (fn OnDangerousDownloadOpenedEventCallbackFunc) Register() js.Func[func(dict *DangerousDownloadInfo)] {
	return js.RegisterCallback[func(dict *DangerousDownloadInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDangerousDownloadOpenedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DangerousDownloadInfo
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

type OnDangerousDownloadOpenedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, dict *DangerousDownloadInfo) js.Ref
	Arg T
}

func (cb *OnDangerousDownloadOpenedEventCallback[T]) Register() js.Func[func(dict *DangerousDownloadInfo)] {
	return js.RegisterCallback[func(dict *DangerousDownloadInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDangerousDownloadOpenedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 DangerousDownloadInfo
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

// HasFuncOnDangerousDownloadOpened returns true if the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.addListener" exists.
func HasFuncOnDangerousDownloadOpened() bool {
	return js.True == bindings.HasFuncOnDangerousDownloadOpened()
}

// FuncOnDangerousDownloadOpened returns the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.addListener".
func FuncOnDangerousDownloadOpened() (fn js.Func[func(callback js.Func[func(dict *DangerousDownloadInfo)])]) {
	bindings.FuncOnDangerousDownloadOpened(
		js.Pointer(&fn),
	)
	return
}

// OnDangerousDownloadOpened calls the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.addListener" directly.
func OnDangerousDownloadOpened(callback js.Func[func(dict *DangerousDownloadInfo)]) (ret js.Void) {
	bindings.CallOnDangerousDownloadOpened(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDangerousDownloadOpened calls the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDangerousDownloadOpened(callback js.Func[func(dict *DangerousDownloadInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDangerousDownloadOpened(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDangerousDownloadOpened returns true if the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.removeListener" exists.
func HasFuncOffDangerousDownloadOpened() bool {
	return js.True == bindings.HasFuncOffDangerousDownloadOpened()
}

// FuncOffDangerousDownloadOpened returns the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.removeListener".
func FuncOffDangerousDownloadOpened() (fn js.Func[func(callback js.Func[func(dict *DangerousDownloadInfo)])]) {
	bindings.FuncOffDangerousDownloadOpened(
		js.Pointer(&fn),
	)
	return
}

// OffDangerousDownloadOpened calls the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.removeListener" directly.
func OffDangerousDownloadOpened(callback js.Func[func(dict *DangerousDownloadInfo)]) (ret js.Void) {
	bindings.CallOffDangerousDownloadOpened(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDangerousDownloadOpened calls the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDangerousDownloadOpened(callback js.Func[func(dict *DangerousDownloadInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDangerousDownloadOpened(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDangerousDownloadOpened returns true if the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.hasListener" exists.
func HasFuncHasOnDangerousDownloadOpened() bool {
	return js.True == bindings.HasFuncHasOnDangerousDownloadOpened()
}

// FuncHasOnDangerousDownloadOpened returns the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.hasListener".
func FuncHasOnDangerousDownloadOpened() (fn js.Func[func(callback js.Func[func(dict *DangerousDownloadInfo)]) bool]) {
	bindings.FuncHasOnDangerousDownloadOpened(
		js.Pointer(&fn),
	)
	return
}

// HasOnDangerousDownloadOpened calls the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.hasListener" directly.
func HasOnDangerousDownloadOpened(callback js.Func[func(dict *DangerousDownloadInfo)]) (ret bool) {
	bindings.CallHasOnDangerousDownloadOpened(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDangerousDownloadOpened calls the function "WEBEXT.safeBrowsingPrivate.onDangerousDownloadOpened.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDangerousDownloadOpened(callback js.Func[func(dict *DangerousDownloadInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDangerousDownloadOpened(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPolicySpecifiedPasswordChangedEventCallbackFunc func(this js.Ref, userName js.String) js.Ref

func (fn OnPolicySpecifiedPasswordChangedEventCallbackFunc) Register() js.Func[func(userName js.String)] {
	return js.RegisterCallback[func(userName js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPolicySpecifiedPasswordChangedEventCallbackFunc) DispatchCallback(
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

type OnPolicySpecifiedPasswordChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, userName js.String) js.Ref
	Arg T
}

func (cb *OnPolicySpecifiedPasswordChangedEventCallback[T]) Register() js.Func[func(userName js.String)] {
	return js.RegisterCallback[func(userName js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPolicySpecifiedPasswordChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnPolicySpecifiedPasswordChanged returns true if the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.addListener" exists.
func HasFuncOnPolicySpecifiedPasswordChanged() bool {
	return js.True == bindings.HasFuncOnPolicySpecifiedPasswordChanged()
}

// FuncOnPolicySpecifiedPasswordChanged returns the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.addListener".
func FuncOnPolicySpecifiedPasswordChanged() (fn js.Func[func(callback js.Func[func(userName js.String)])]) {
	bindings.FuncOnPolicySpecifiedPasswordChanged(
		js.Pointer(&fn),
	)
	return
}

// OnPolicySpecifiedPasswordChanged calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.addListener" directly.
func OnPolicySpecifiedPasswordChanged(callback js.Func[func(userName js.String)]) (ret js.Void) {
	bindings.CallOnPolicySpecifiedPasswordChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPolicySpecifiedPasswordChanged calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPolicySpecifiedPasswordChanged(callback js.Func[func(userName js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPolicySpecifiedPasswordChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPolicySpecifiedPasswordChanged returns true if the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.removeListener" exists.
func HasFuncOffPolicySpecifiedPasswordChanged() bool {
	return js.True == bindings.HasFuncOffPolicySpecifiedPasswordChanged()
}

// FuncOffPolicySpecifiedPasswordChanged returns the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.removeListener".
func FuncOffPolicySpecifiedPasswordChanged() (fn js.Func[func(callback js.Func[func(userName js.String)])]) {
	bindings.FuncOffPolicySpecifiedPasswordChanged(
		js.Pointer(&fn),
	)
	return
}

// OffPolicySpecifiedPasswordChanged calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.removeListener" directly.
func OffPolicySpecifiedPasswordChanged(callback js.Func[func(userName js.String)]) (ret js.Void) {
	bindings.CallOffPolicySpecifiedPasswordChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPolicySpecifiedPasswordChanged calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPolicySpecifiedPasswordChanged(callback js.Func[func(userName js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPolicySpecifiedPasswordChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPolicySpecifiedPasswordChanged returns true if the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.hasListener" exists.
func HasFuncHasOnPolicySpecifiedPasswordChanged() bool {
	return js.True == bindings.HasFuncHasOnPolicySpecifiedPasswordChanged()
}

// FuncHasOnPolicySpecifiedPasswordChanged returns the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.hasListener".
func FuncHasOnPolicySpecifiedPasswordChanged() (fn js.Func[func(callback js.Func[func(userName js.String)]) bool]) {
	bindings.FuncHasOnPolicySpecifiedPasswordChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnPolicySpecifiedPasswordChanged calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.hasListener" directly.
func HasOnPolicySpecifiedPasswordChanged(callback js.Func[func(userName js.String)]) (ret bool) {
	bindings.CallHasOnPolicySpecifiedPasswordChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPolicySpecifiedPasswordChanged calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPolicySpecifiedPasswordChanged(callback js.Func[func(userName js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPolicySpecifiedPasswordChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnPolicySpecifiedPasswordReuseDetectedEventCallbackFunc func(this js.Ref, reuseDetails *PolicySpecifiedPasswordReuse) js.Ref

func (fn OnPolicySpecifiedPasswordReuseDetectedEventCallbackFunc) Register() js.Func[func(reuseDetails *PolicySpecifiedPasswordReuse)] {
	return js.RegisterCallback[func(reuseDetails *PolicySpecifiedPasswordReuse)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPolicySpecifiedPasswordReuseDetectedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PolicySpecifiedPasswordReuse
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

type OnPolicySpecifiedPasswordReuseDetectedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, reuseDetails *PolicySpecifiedPasswordReuse) js.Ref
	Arg T
}

func (cb *OnPolicySpecifiedPasswordReuseDetectedEventCallback[T]) Register() js.Func[func(reuseDetails *PolicySpecifiedPasswordReuse)] {
	return js.RegisterCallback[func(reuseDetails *PolicySpecifiedPasswordReuse)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPolicySpecifiedPasswordReuseDetectedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 PolicySpecifiedPasswordReuse
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

// HasFuncOnPolicySpecifiedPasswordReuseDetected returns true if the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.addListener" exists.
func HasFuncOnPolicySpecifiedPasswordReuseDetected() bool {
	return js.True == bindings.HasFuncOnPolicySpecifiedPasswordReuseDetected()
}

// FuncOnPolicySpecifiedPasswordReuseDetected returns the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.addListener".
func FuncOnPolicySpecifiedPasswordReuseDetected() (fn js.Func[func(callback js.Func[func(reuseDetails *PolicySpecifiedPasswordReuse)])]) {
	bindings.FuncOnPolicySpecifiedPasswordReuseDetected(
		js.Pointer(&fn),
	)
	return
}

// OnPolicySpecifiedPasswordReuseDetected calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.addListener" directly.
func OnPolicySpecifiedPasswordReuseDetected(callback js.Func[func(reuseDetails *PolicySpecifiedPasswordReuse)]) (ret js.Void) {
	bindings.CallOnPolicySpecifiedPasswordReuseDetected(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPolicySpecifiedPasswordReuseDetected calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPolicySpecifiedPasswordReuseDetected(callback js.Func[func(reuseDetails *PolicySpecifiedPasswordReuse)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPolicySpecifiedPasswordReuseDetected(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPolicySpecifiedPasswordReuseDetected returns true if the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.removeListener" exists.
func HasFuncOffPolicySpecifiedPasswordReuseDetected() bool {
	return js.True == bindings.HasFuncOffPolicySpecifiedPasswordReuseDetected()
}

// FuncOffPolicySpecifiedPasswordReuseDetected returns the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.removeListener".
func FuncOffPolicySpecifiedPasswordReuseDetected() (fn js.Func[func(callback js.Func[func(reuseDetails *PolicySpecifiedPasswordReuse)])]) {
	bindings.FuncOffPolicySpecifiedPasswordReuseDetected(
		js.Pointer(&fn),
	)
	return
}

// OffPolicySpecifiedPasswordReuseDetected calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.removeListener" directly.
func OffPolicySpecifiedPasswordReuseDetected(callback js.Func[func(reuseDetails *PolicySpecifiedPasswordReuse)]) (ret js.Void) {
	bindings.CallOffPolicySpecifiedPasswordReuseDetected(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPolicySpecifiedPasswordReuseDetected calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPolicySpecifiedPasswordReuseDetected(callback js.Func[func(reuseDetails *PolicySpecifiedPasswordReuse)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPolicySpecifiedPasswordReuseDetected(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPolicySpecifiedPasswordReuseDetected returns true if the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.hasListener" exists.
func HasFuncHasOnPolicySpecifiedPasswordReuseDetected() bool {
	return js.True == bindings.HasFuncHasOnPolicySpecifiedPasswordReuseDetected()
}

// FuncHasOnPolicySpecifiedPasswordReuseDetected returns the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.hasListener".
func FuncHasOnPolicySpecifiedPasswordReuseDetected() (fn js.Func[func(callback js.Func[func(reuseDetails *PolicySpecifiedPasswordReuse)]) bool]) {
	bindings.FuncHasOnPolicySpecifiedPasswordReuseDetected(
		js.Pointer(&fn),
	)
	return
}

// HasOnPolicySpecifiedPasswordReuseDetected calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.hasListener" directly.
func HasOnPolicySpecifiedPasswordReuseDetected(callback js.Func[func(reuseDetails *PolicySpecifiedPasswordReuse)]) (ret bool) {
	bindings.CallHasOnPolicySpecifiedPasswordReuseDetected(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPolicySpecifiedPasswordReuseDetected calls the function "WEBEXT.safeBrowsingPrivate.onPolicySpecifiedPasswordReuseDetected.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPolicySpecifiedPasswordReuseDetected(callback js.Func[func(reuseDetails *PolicySpecifiedPasswordReuse)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPolicySpecifiedPasswordReuseDetected(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSecurityInterstitialProceededEventCallbackFunc func(this js.Ref, dict *InterstitialInfo) js.Ref

func (fn OnSecurityInterstitialProceededEventCallbackFunc) Register() js.Func[func(dict *InterstitialInfo)] {
	return js.RegisterCallback[func(dict *InterstitialInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSecurityInterstitialProceededEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 InterstitialInfo
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

type OnSecurityInterstitialProceededEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, dict *InterstitialInfo) js.Ref
	Arg T
}

func (cb *OnSecurityInterstitialProceededEventCallback[T]) Register() js.Func[func(dict *InterstitialInfo)] {
	return js.RegisterCallback[func(dict *InterstitialInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSecurityInterstitialProceededEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 InterstitialInfo
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

// HasFuncOnSecurityInterstitialProceeded returns true if the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.addListener" exists.
func HasFuncOnSecurityInterstitialProceeded() bool {
	return js.True == bindings.HasFuncOnSecurityInterstitialProceeded()
}

// FuncOnSecurityInterstitialProceeded returns the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.addListener".
func FuncOnSecurityInterstitialProceeded() (fn js.Func[func(callback js.Func[func(dict *InterstitialInfo)])]) {
	bindings.FuncOnSecurityInterstitialProceeded(
		js.Pointer(&fn),
	)
	return
}

// OnSecurityInterstitialProceeded calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.addListener" directly.
func OnSecurityInterstitialProceeded(callback js.Func[func(dict *InterstitialInfo)]) (ret js.Void) {
	bindings.CallOnSecurityInterstitialProceeded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSecurityInterstitialProceeded calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSecurityInterstitialProceeded(callback js.Func[func(dict *InterstitialInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSecurityInterstitialProceeded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSecurityInterstitialProceeded returns true if the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.removeListener" exists.
func HasFuncOffSecurityInterstitialProceeded() bool {
	return js.True == bindings.HasFuncOffSecurityInterstitialProceeded()
}

// FuncOffSecurityInterstitialProceeded returns the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.removeListener".
func FuncOffSecurityInterstitialProceeded() (fn js.Func[func(callback js.Func[func(dict *InterstitialInfo)])]) {
	bindings.FuncOffSecurityInterstitialProceeded(
		js.Pointer(&fn),
	)
	return
}

// OffSecurityInterstitialProceeded calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.removeListener" directly.
func OffSecurityInterstitialProceeded(callback js.Func[func(dict *InterstitialInfo)]) (ret js.Void) {
	bindings.CallOffSecurityInterstitialProceeded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSecurityInterstitialProceeded calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSecurityInterstitialProceeded(callback js.Func[func(dict *InterstitialInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSecurityInterstitialProceeded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSecurityInterstitialProceeded returns true if the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.hasListener" exists.
func HasFuncHasOnSecurityInterstitialProceeded() bool {
	return js.True == bindings.HasFuncHasOnSecurityInterstitialProceeded()
}

// FuncHasOnSecurityInterstitialProceeded returns the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.hasListener".
func FuncHasOnSecurityInterstitialProceeded() (fn js.Func[func(callback js.Func[func(dict *InterstitialInfo)]) bool]) {
	bindings.FuncHasOnSecurityInterstitialProceeded(
		js.Pointer(&fn),
	)
	return
}

// HasOnSecurityInterstitialProceeded calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.hasListener" directly.
func HasOnSecurityInterstitialProceeded(callback js.Func[func(dict *InterstitialInfo)]) (ret bool) {
	bindings.CallHasOnSecurityInterstitialProceeded(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSecurityInterstitialProceeded calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialProceeded.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSecurityInterstitialProceeded(callback js.Func[func(dict *InterstitialInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSecurityInterstitialProceeded(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnSecurityInterstitialShownEventCallbackFunc func(this js.Ref, dict *InterstitialInfo) js.Ref

func (fn OnSecurityInterstitialShownEventCallbackFunc) Register() js.Func[func(dict *InterstitialInfo)] {
	return js.RegisterCallback[func(dict *InterstitialInfo)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnSecurityInterstitialShownEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 InterstitialInfo
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

type OnSecurityInterstitialShownEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, dict *InterstitialInfo) js.Ref
	Arg T
}

func (cb *OnSecurityInterstitialShownEventCallback[T]) Register() js.Func[func(dict *InterstitialInfo)] {
	return js.RegisterCallback[func(dict *InterstitialInfo)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnSecurityInterstitialShownEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 InterstitialInfo
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

// HasFuncOnSecurityInterstitialShown returns true if the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.addListener" exists.
func HasFuncOnSecurityInterstitialShown() bool {
	return js.True == bindings.HasFuncOnSecurityInterstitialShown()
}

// FuncOnSecurityInterstitialShown returns the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.addListener".
func FuncOnSecurityInterstitialShown() (fn js.Func[func(callback js.Func[func(dict *InterstitialInfo)])]) {
	bindings.FuncOnSecurityInterstitialShown(
		js.Pointer(&fn),
	)
	return
}

// OnSecurityInterstitialShown calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.addListener" directly.
func OnSecurityInterstitialShown(callback js.Func[func(dict *InterstitialInfo)]) (ret js.Void) {
	bindings.CallOnSecurityInterstitialShown(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSecurityInterstitialShown calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSecurityInterstitialShown(callback js.Func[func(dict *InterstitialInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSecurityInterstitialShown(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSecurityInterstitialShown returns true if the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.removeListener" exists.
func HasFuncOffSecurityInterstitialShown() bool {
	return js.True == bindings.HasFuncOffSecurityInterstitialShown()
}

// FuncOffSecurityInterstitialShown returns the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.removeListener".
func FuncOffSecurityInterstitialShown() (fn js.Func[func(callback js.Func[func(dict *InterstitialInfo)])]) {
	bindings.FuncOffSecurityInterstitialShown(
		js.Pointer(&fn),
	)
	return
}

// OffSecurityInterstitialShown calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.removeListener" directly.
func OffSecurityInterstitialShown(callback js.Func[func(dict *InterstitialInfo)]) (ret js.Void) {
	bindings.CallOffSecurityInterstitialShown(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSecurityInterstitialShown calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSecurityInterstitialShown(callback js.Func[func(dict *InterstitialInfo)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSecurityInterstitialShown(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSecurityInterstitialShown returns true if the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.hasListener" exists.
func HasFuncHasOnSecurityInterstitialShown() bool {
	return js.True == bindings.HasFuncHasOnSecurityInterstitialShown()
}

// FuncHasOnSecurityInterstitialShown returns the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.hasListener".
func FuncHasOnSecurityInterstitialShown() (fn js.Func[func(callback js.Func[func(dict *InterstitialInfo)]) bool]) {
	bindings.FuncHasOnSecurityInterstitialShown(
		js.Pointer(&fn),
	)
	return
}

// HasOnSecurityInterstitialShown calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.hasListener" directly.
func HasOnSecurityInterstitialShown(callback js.Func[func(dict *InterstitialInfo)]) (ret bool) {
	bindings.CallHasOnSecurityInterstitialShown(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSecurityInterstitialShown calls the function "WEBEXT.safeBrowsingPrivate.onSecurityInterstitialShown.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSecurityInterstitialShown(callback js.Func[func(dict *InterstitialInfo)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSecurityInterstitialShown(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

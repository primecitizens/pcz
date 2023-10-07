// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package declarativecontent

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/declarativecontent/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/events"
)

type ImageDataType = js.TypedArray[uint8]

type PageStateMatcherInstanceType uint32

const (
	_ PageStateMatcherInstanceType = iota

	PageStateMatcherInstanceType_DECLARATIVE_CONTENT_PAGE_STATE_MATCHER
)

func (PageStateMatcherInstanceType) FromRef(str js.Ref) PageStateMatcherInstanceType {
	return PageStateMatcherInstanceType(bindings.ConstOfPageStateMatcherInstanceType(str))
}

func (x PageStateMatcherInstanceType) String() (string, bool) {
	switch x {
	case PageStateMatcherInstanceType_DECLARATIVE_CONTENT_PAGE_STATE_MATCHER:
		return "declarativeContent.PageStateMatcher", true
	default:
		return "", false
	}
}

type PageStateMatcher struct {
	// Css is "PageStateMatcher.css"
	//
	// Optional
	Css js.Array[js.String]
	// InstanceType is "PageStateMatcher.instanceType"
	//
	// Required
	InstanceType PageStateMatcherInstanceType
	// IsBookmarked is "PageStateMatcher.isBookmarked"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsBookmarked MUST be set to true to make this field effective.
	IsBookmarked bool
	// PageUrl is "PageStateMatcher.pageUrl"
	//
	// Optional
	//
	// NOTE: PageUrl.FFI_USE MUST be set to true to get PageUrl used.
	PageUrl events.UrlFilter

	FFI_USE_IsBookmarked bool // for IsBookmarked.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PageStateMatcher with all fields set.
func (p PageStateMatcher) FromRef(ref js.Ref) PageStateMatcher {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PageStateMatcher in the application heap.
func (p PageStateMatcher) New() js.Ref {
	return bindings.PageStateMatcherJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PageStateMatcher) UpdateFrom(ref js.Ref) {
	bindings.PageStateMatcherJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PageStateMatcher) Update(ref js.Ref) {
	bindings.PageStateMatcherJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PageStateMatcher) FreeMembers(recursive bool) {
	js.Free(
		p.Css.Ref(),
	)
	p.Css = p.Css.FromRef(js.Undefined)
	if recursive {
		p.PageUrl.FreeMembers(true)
	}
}

type RequestContentScriptInstanceType uint32

const (
	_ RequestContentScriptInstanceType = iota

	RequestContentScriptInstanceType_DECLARATIVE_CONTENT_REQUEST_CONTENT_SCRIPT
)

func (RequestContentScriptInstanceType) FromRef(str js.Ref) RequestContentScriptInstanceType {
	return RequestContentScriptInstanceType(bindings.ConstOfRequestContentScriptInstanceType(str))
}

func (x RequestContentScriptInstanceType) String() (string, bool) {
	switch x {
	case RequestContentScriptInstanceType_DECLARATIVE_CONTENT_REQUEST_CONTENT_SCRIPT:
		return "declarativeContent.RequestContentScript", true
	default:
		return "", false
	}
}

type RequestContentScript struct {
	// AllFrames is "RequestContentScript.allFrames"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllFrames MUST be set to true to make this field effective.
	AllFrames bool
	// Css is "RequestContentScript.css"
	//
	// Optional
	Css js.Array[js.String]
	// InstanceType is "RequestContentScript.instanceType"
	//
	// Required
	InstanceType RequestContentScriptInstanceType
	// Js is "RequestContentScript.js"
	//
	// Optional
	Js js.Array[js.String]
	// MatchAboutBlank is "RequestContentScript.matchAboutBlank"
	//
	// Optional
	//
	// NOTE: FFI_USE_MatchAboutBlank MUST be set to true to make this field effective.
	MatchAboutBlank bool

	FFI_USE_AllFrames       bool // for AllFrames.
	FFI_USE_MatchAboutBlank bool // for MatchAboutBlank.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RequestContentScript with all fields set.
func (p RequestContentScript) FromRef(ref js.Ref) RequestContentScript {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RequestContentScript in the application heap.
func (p RequestContentScript) New() js.Ref {
	return bindings.RequestContentScriptJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RequestContentScript) UpdateFrom(ref js.Ref) {
	bindings.RequestContentScriptJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RequestContentScript) Update(ref js.Ref) {
	bindings.RequestContentScriptJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RequestContentScript) FreeMembers(recursive bool) {
	js.Free(
		p.Css.Ref(),
		p.Js.Ref(),
	)
	p.Css = p.Css.FromRef(js.Undefined)
	p.Js = p.Js.FromRef(js.Undefined)
}

type OneOf_TypedArrayUint8_Any struct {
	ref js.Ref
}

func (x OneOf_TypedArrayUint8_Any) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayUint8_Any) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayUint8_Any) FromRef(ref js.Ref) OneOf_TypedArrayUint8_Any {
	return OneOf_TypedArrayUint8_Any{
		ref: ref,
	}
}

func (x OneOf_TypedArrayUint8_Any) TypedArrayUint8() js.TypedArray[uint8] {
	return js.TypedArray[uint8]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayUint8_Any) Any() js.Any {
	return js.Any{}.FromRef(x.ref)
}

type SetIconInstanceType uint32

const (
	_ SetIconInstanceType = iota

	SetIconInstanceType_DECLARATIVE_CONTENT_SET_ICON
)

func (SetIconInstanceType) FromRef(str js.Ref) SetIconInstanceType {
	return SetIconInstanceType(bindings.ConstOfSetIconInstanceType(str))
}

func (x SetIconInstanceType) String() (string, bool) {
	switch x {
	case SetIconInstanceType_DECLARATIVE_CONTENT_SET_ICON:
		return "declarativeContent.SetIcon", true
	default:
		return "", false
	}
}

type SetIcon struct {
	// ImageData is "SetIcon.imageData"
	//
	// Optional
	ImageData OneOf_TypedArrayUint8_Any
	// InstanceType is "SetIcon.instanceType"
	//
	// Required
	InstanceType SetIconInstanceType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SetIcon with all fields set.
func (p SetIcon) FromRef(ref js.Ref) SetIcon {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SetIcon in the application heap.
func (p SetIcon) New() js.Ref {
	return bindings.SetIconJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SetIcon) UpdateFrom(ref js.Ref) {
	bindings.SetIconJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SetIcon) Update(ref js.Ref) {
	bindings.SetIconJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SetIcon) FreeMembers(recursive bool) {
	js.Free(
		p.ImageData.Ref(),
	)
	p.ImageData = p.ImageData.FromRef(js.Undefined)
}

type ShowActionInstanceType uint32

const (
	_ ShowActionInstanceType = iota

	ShowActionInstanceType_DECLARATIVE_CONTENT_SHOW_ACTION
)

func (ShowActionInstanceType) FromRef(str js.Ref) ShowActionInstanceType {
	return ShowActionInstanceType(bindings.ConstOfShowActionInstanceType(str))
}

func (x ShowActionInstanceType) String() (string, bool) {
	switch x {
	case ShowActionInstanceType_DECLARATIVE_CONTENT_SHOW_ACTION:
		return "declarativeContent.ShowAction", true
	default:
		return "", false
	}
}

type ShowAction struct {
	// InstanceType is "ShowAction.instanceType"
	//
	// Required
	InstanceType ShowActionInstanceType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ShowAction with all fields set.
func (p ShowAction) FromRef(ref js.Ref) ShowAction {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ShowAction in the application heap.
func (p ShowAction) New() js.Ref {
	return bindings.ShowActionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ShowAction) UpdateFrom(ref js.Ref) {
	bindings.ShowActionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ShowAction) Update(ref js.Ref) {
	bindings.ShowActionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ShowAction) FreeMembers(recursive bool) {
}

type ShowPageActionInstanceType uint32

const (
	_ ShowPageActionInstanceType = iota

	ShowPageActionInstanceType_DECLARATIVE_CONTENT_SHOW_PAGE_ACTION
)

func (ShowPageActionInstanceType) FromRef(str js.Ref) ShowPageActionInstanceType {
	return ShowPageActionInstanceType(bindings.ConstOfShowPageActionInstanceType(str))
}

func (x ShowPageActionInstanceType) String() (string, bool) {
	switch x {
	case ShowPageActionInstanceType_DECLARATIVE_CONTENT_SHOW_PAGE_ACTION:
		return "declarativeContent.ShowPageAction", true
	default:
		return "", false
	}
}

type ShowPageAction struct {
	// InstanceType is "ShowPageAction.instanceType"
	//
	// Required
	InstanceType ShowPageActionInstanceType

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ShowPageAction with all fields set.
func (p ShowPageAction) FromRef(ref js.Ref) ShowPageAction {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ShowPageAction in the application heap.
func (p ShowPageAction) New() js.Ref {
	return bindings.ShowPageActionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ShowPageAction) UpdateFrom(ref js.Ref) {
	bindings.ShowPageActionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ShowPageAction) Update(ref js.Ref) {
	bindings.ShowPageActionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ShowPageAction) FreeMembers(recursive bool) {
}

type OnPageChangedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnPageChangedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnPageChangedEventCallbackFunc) DispatchCallback(
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

type OnPageChangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnPageChangedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnPageChangedEventCallback[T]) DispatchCallback(
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

// HasFuncOnPageChanged returns true if the function "WEBEXT.declarativeContent.onPageChanged.addListener" exists.
func HasFuncOnPageChanged() bool {
	return js.True == bindings.HasFuncOnPageChanged()
}

// FuncOnPageChanged returns the function "WEBEXT.declarativeContent.onPageChanged.addListener".
func FuncOnPageChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnPageChanged(
		js.Pointer(&fn),
	)
	return
}

// OnPageChanged calls the function "WEBEXT.declarativeContent.onPageChanged.addListener" directly.
func OnPageChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnPageChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPageChanged calls the function "WEBEXT.declarativeContent.onPageChanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPageChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPageChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPageChanged returns true if the function "WEBEXT.declarativeContent.onPageChanged.removeListener" exists.
func HasFuncOffPageChanged() bool {
	return js.True == bindings.HasFuncOffPageChanged()
}

// FuncOffPageChanged returns the function "WEBEXT.declarativeContent.onPageChanged.removeListener".
func FuncOffPageChanged() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffPageChanged(
		js.Pointer(&fn),
	)
	return
}

// OffPageChanged calls the function "WEBEXT.declarativeContent.onPageChanged.removeListener" directly.
func OffPageChanged(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffPageChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPageChanged calls the function "WEBEXT.declarativeContent.onPageChanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPageChanged(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPageChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPageChanged returns true if the function "WEBEXT.declarativeContent.onPageChanged.hasListener" exists.
func HasFuncHasOnPageChanged() bool {
	return js.True == bindings.HasFuncHasOnPageChanged()
}

// FuncHasOnPageChanged returns the function "WEBEXT.declarativeContent.onPageChanged.hasListener".
func FuncHasOnPageChanged() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnPageChanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnPageChanged calls the function "WEBEXT.declarativeContent.onPageChanged.hasListener" directly.
func HasOnPageChanged(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnPageChanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPageChanged calls the function "WEBEXT.declarativeContent.onPageChanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPageChanged(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPageChanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

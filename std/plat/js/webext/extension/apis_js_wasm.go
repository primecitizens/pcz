// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package extension

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/extension/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/runtime"
)

type ViewType uint32

const (
	_ ViewType = iota

	ViewType_TAB
	ViewType_POPUP
)

func (ViewType) FromRef(str js.Ref) ViewType {
	return ViewType(bindings.ConstOfViewType(str))
}

func (x ViewType) String() (string, bool) {
	switch x {
	case ViewType_TAB:
		return "tab", true
	case ViewType_POPUP:
		return "popup", true
	default:
		return "", false
	}
}

type GetViewsArgFetchProperties struct {
	// TabId is "GetViewsArgFetchProperties.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int64
	// Type is "GetViewsArgFetchProperties.type"
	//
	// Optional
	Type ViewType
	// WindowId is "GetViewsArgFetchProperties.windowId"
	//
	// Optional
	//
	// NOTE: FFI_USE_WindowId MUST be set to true to make this field effective.
	WindowId int64

	FFI_USE_TabId    bool // for TabId.
	FFI_USE_WindowId bool // for WindowId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetViewsArgFetchProperties with all fields set.
func (p GetViewsArgFetchProperties) FromRef(ref js.Ref) GetViewsArgFetchProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetViewsArgFetchProperties in the application heap.
func (p GetViewsArgFetchProperties) New() js.Ref {
	return bindings.GetViewsArgFetchPropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GetViewsArgFetchProperties) UpdateFrom(ref js.Ref) {
	bindings.GetViewsArgFetchPropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetViewsArgFetchProperties) Update(ref js.Ref) {
	bindings.GetViewsArgFetchPropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetViewsArgFetchProperties) FreeMembers(recursive bool) {
}

type LastErrorProperty struct {
	// Message is "LastErrorProperty.message"
	//
	// Required
	Message js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LastErrorProperty with all fields set.
func (p LastErrorProperty) FromRef(ref js.Ref) LastErrorProperty {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LastErrorProperty in the application heap.
func (p LastErrorProperty) New() js.Ref {
	return bindings.LastErrorPropertyJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LastErrorProperty) UpdateFrom(ref js.Ref) {
	bindings.LastErrorPropertyJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LastErrorProperty) Update(ref js.Ref) {
	bindings.LastErrorPropertyJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LastErrorProperty) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
}

type OnRequestArgSendResponseFunc func(this js.Ref) js.Ref

func (fn OnRequestArgSendResponseFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRequestArgSendResponseFunc) DispatchCallback(
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

type OnRequestArgSendResponse[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnRequestArgSendResponse[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRequestArgSendResponse[T]) DispatchCallback(
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

type OnRequestExternalArgSendResponseFunc func(this js.Ref) js.Ref

func (fn OnRequestExternalArgSendResponseFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRequestExternalArgSendResponseFunc) DispatchCallback(
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

type OnRequestExternalArgSendResponse[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnRequestExternalArgSendResponse[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRequestExternalArgSendResponse[T]) DispatchCallback(
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

// HasFuncGetBackgroundPage returns true if the function "WEBEXT.extension.getBackgroundPage" exists.
func HasFuncGetBackgroundPage() bool {
	return js.True == bindings.HasFuncGetBackgroundPage()
}

// FuncGetBackgroundPage returns the function "WEBEXT.extension.getBackgroundPage".
func FuncGetBackgroundPage() (fn js.Func[func() js.Any]) {
	bindings.FuncGetBackgroundPage(
		js.Pointer(&fn),
	)
	return
}

// GetBackgroundPage calls the function "WEBEXT.extension.getBackgroundPage" directly.
func GetBackgroundPage() (ret js.Any) {
	bindings.CallGetBackgroundPage(
		js.Pointer(&ret),
	)

	return
}

// TryGetBackgroundPage calls the function "WEBEXT.extension.getBackgroundPage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetBackgroundPage() (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetBackgroundPage(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetExtensionTabs returns true if the function "WEBEXT.extension.getExtensionTabs" exists.
func HasFuncGetExtensionTabs() bool {
	return js.True == bindings.HasFuncGetExtensionTabs()
}

// FuncGetExtensionTabs returns the function "WEBEXT.extension.getExtensionTabs".
func FuncGetExtensionTabs() (fn js.Func[func(windowId int64) js.Array[js.Any]]) {
	bindings.FuncGetExtensionTabs(
		js.Pointer(&fn),
	)
	return
}

// GetExtensionTabs calls the function "WEBEXT.extension.getExtensionTabs" directly.
func GetExtensionTabs(windowId int64) (ret js.Array[js.Any]) {
	bindings.CallGetExtensionTabs(
		js.Pointer(&ret),
		float64(windowId),
	)

	return
}

// TryGetExtensionTabs calls the function "WEBEXT.extension.getExtensionTabs"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetExtensionTabs(windowId int64) (ret js.Array[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetExtensionTabs(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(windowId),
	)

	return
}

// HasFuncGetURL returns true if the function "WEBEXT.extension.getURL" exists.
func HasFuncGetURL() bool {
	return js.True == bindings.HasFuncGetURL()
}

// FuncGetURL returns the function "WEBEXT.extension.getURL".
func FuncGetURL() (fn js.Func[func(path js.String) js.String]) {
	bindings.FuncGetURL(
		js.Pointer(&fn),
	)
	return
}

// GetURL calls the function "WEBEXT.extension.getURL" directly.
func GetURL(path js.String) (ret js.String) {
	bindings.CallGetURL(
		js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryGetURL calls the function "WEBEXT.extension.getURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetURL(path js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetURL(
		js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasFuncGetViews returns true if the function "WEBEXT.extension.getViews" exists.
func HasFuncGetViews() bool {
	return js.True == bindings.HasFuncGetViews()
}

// FuncGetViews returns the function "WEBEXT.extension.getViews".
func FuncGetViews() (fn js.Func[func(fetchProperties GetViewsArgFetchProperties) js.Array[js.Any]]) {
	bindings.FuncGetViews(
		js.Pointer(&fn),
	)
	return
}

// GetViews calls the function "WEBEXT.extension.getViews" directly.
func GetViews(fetchProperties GetViewsArgFetchProperties) (ret js.Array[js.Any]) {
	bindings.CallGetViews(
		js.Pointer(&ret),
		js.Pointer(&fetchProperties),
	)

	return
}

// TryGetViews calls the function "WEBEXT.extension.getViews"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetViews(fetchProperties GetViewsArgFetchProperties) (ret js.Array[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetViews(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&fetchProperties),
	)

	return
}

// InIncognitoContext returns the value of property "WEBEXT.extension.inIncognitoContext".
//
// The returned bool will be false if there is no such property.
func InIncognitoContext() (ret bool, ok bool) {
	ok = js.True == bindings.GetInIncognitoContext(
		js.Pointer(&ret),
	)

	return
}

// SetInIncognitoContext sets the value of property "WEBEXT.extension.inIncognitoContext" to val.
//
// It returns false if the property cannot be set.
func SetInIncognitoContext(val bool) bool {
	return js.True == bindings.SetInIncognitoContext(
		js.Bool(bool(val)))
}

// HasFuncIsAllowedFileSchemeAccess returns true if the function "WEBEXT.extension.isAllowedFileSchemeAccess" exists.
func HasFuncIsAllowedFileSchemeAccess() bool {
	return js.True == bindings.HasFuncIsAllowedFileSchemeAccess()
}

// FuncIsAllowedFileSchemeAccess returns the function "WEBEXT.extension.isAllowedFileSchemeAccess".
func FuncIsAllowedFileSchemeAccess() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsAllowedFileSchemeAccess(
		js.Pointer(&fn),
	)
	return
}

// IsAllowedFileSchemeAccess calls the function "WEBEXT.extension.isAllowedFileSchemeAccess" directly.
func IsAllowedFileSchemeAccess() (ret js.Promise[js.Boolean]) {
	bindings.CallIsAllowedFileSchemeAccess(
		js.Pointer(&ret),
	)

	return
}

// TryIsAllowedFileSchemeAccess calls the function "WEBEXT.extension.isAllowedFileSchemeAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsAllowedFileSchemeAccess() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsAllowedFileSchemeAccess(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIsAllowedIncognitoAccess returns true if the function "WEBEXT.extension.isAllowedIncognitoAccess" exists.
func HasFuncIsAllowedIncognitoAccess() bool {
	return js.True == bindings.HasFuncIsAllowedIncognitoAccess()
}

// FuncIsAllowedIncognitoAccess returns the function "WEBEXT.extension.isAllowedIncognitoAccess".
func FuncIsAllowedIncognitoAccess() (fn js.Func[func() js.Promise[js.Boolean]]) {
	bindings.FuncIsAllowedIncognitoAccess(
		js.Pointer(&fn),
	)
	return
}

// IsAllowedIncognitoAccess calls the function "WEBEXT.extension.isAllowedIncognitoAccess" directly.
func IsAllowedIncognitoAccess() (ret js.Promise[js.Boolean]) {
	bindings.CallIsAllowedIncognitoAccess(
		js.Pointer(&ret),
	)

	return
}

// TryIsAllowedIncognitoAccess calls the function "WEBEXT.extension.isAllowedIncognitoAccess"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsAllowedIncognitoAccess() (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsAllowedIncognitoAccess(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// LastError returns the value of property "WEBEXT.extension.lastError".
//
// The returned bool will be false if there is no such property.
func LastError() (ret LastErrorProperty, ok bool) {
	ok = js.True == bindings.GetLastError(
		js.Pointer(&ret),
	)

	return
}

// SetLastError sets the value of property "WEBEXT.extension.lastError" to val.
//
// It returns false if the property cannot be set.
func SetLastError(val LastErrorProperty) bool {
	return js.True == bindings.SetLastError(
		js.Pointer(&val))
}

type OnRequestEventCallbackFunc func(this js.Ref, request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()]) js.Ref

func (fn OnRequestEventCallbackFunc) Register() js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])] {
	return js.RegisterCallback[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRequestEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 runtime.MessageSender
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func()]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnRequestEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()]) js.Ref
	Arg T
}

func (cb *OnRequestEventCallback[T]) Register() js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])] {
	return js.RegisterCallback[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRequestEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 runtime.MessageSender
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Any{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func()]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnRequest returns true if the function "WEBEXT.extension.onRequest.addListener" exists.
func HasFuncOnRequest() bool {
	return js.True == bindings.HasFuncOnRequest()
}

// FuncOnRequest returns the function "WEBEXT.extension.onRequest.addListener".
func FuncOnRequest() (fn js.Func[func(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])])]) {
	bindings.FuncOnRequest(
		js.Pointer(&fn),
	)
	return
}

// OnRequest calls the function "WEBEXT.extension.onRequest.addListener" directly.
func OnRequest(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret js.Void) {
	bindings.CallOnRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRequest calls the function "WEBEXT.extension.onRequest.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRequest(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRequest returns true if the function "WEBEXT.extension.onRequest.removeListener" exists.
func HasFuncOffRequest() bool {
	return js.True == bindings.HasFuncOffRequest()
}

// FuncOffRequest returns the function "WEBEXT.extension.onRequest.removeListener".
func FuncOffRequest() (fn js.Func[func(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])])]) {
	bindings.FuncOffRequest(
		js.Pointer(&fn),
	)
	return
}

// OffRequest calls the function "WEBEXT.extension.onRequest.removeListener" directly.
func OffRequest(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret js.Void) {
	bindings.CallOffRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRequest calls the function "WEBEXT.extension.onRequest.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRequest(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRequest returns true if the function "WEBEXT.extension.onRequest.hasListener" exists.
func HasFuncHasOnRequest() bool {
	return js.True == bindings.HasFuncHasOnRequest()
}

// FuncHasOnRequest returns the function "WEBEXT.extension.onRequest.hasListener".
func FuncHasOnRequest() (fn js.Func[func(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) bool]) {
	bindings.FuncHasOnRequest(
		js.Pointer(&fn),
	)
	return
}

// HasOnRequest calls the function "WEBEXT.extension.onRequest.hasListener" directly.
func HasOnRequest(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret bool) {
	bindings.CallHasOnRequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRequest calls the function "WEBEXT.extension.onRequest.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRequest(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRequestExternalEventCallbackFunc func(this js.Ref, request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()]) js.Ref

func (fn OnRequestExternalEventCallbackFunc) Register() js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])] {
	return js.RegisterCallback[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRequestExternalEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 runtime.MessageSender
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.Any{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func()]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnRequestExternalEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()]) js.Ref
	Arg T
}

func (cb *OnRequestExternalEventCallback[T]) Register() js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])] {
	return js.RegisterCallback[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRequestExternalEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 runtime.MessageSender
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Any{}.FromRef(args[0+1]),
		mark.NoEscape(&arg1),
		js.Func[func()]{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnRequestExternal returns true if the function "WEBEXT.extension.onRequestExternal.addListener" exists.
func HasFuncOnRequestExternal() bool {
	return js.True == bindings.HasFuncOnRequestExternal()
}

// FuncOnRequestExternal returns the function "WEBEXT.extension.onRequestExternal.addListener".
func FuncOnRequestExternal() (fn js.Func[func(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])])]) {
	bindings.FuncOnRequestExternal(
		js.Pointer(&fn),
	)
	return
}

// OnRequestExternal calls the function "WEBEXT.extension.onRequestExternal.addListener" directly.
func OnRequestExternal(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret js.Void) {
	bindings.CallOnRequestExternal(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRequestExternal calls the function "WEBEXT.extension.onRequestExternal.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRequestExternal(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRequestExternal(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRequestExternal returns true if the function "WEBEXT.extension.onRequestExternal.removeListener" exists.
func HasFuncOffRequestExternal() bool {
	return js.True == bindings.HasFuncOffRequestExternal()
}

// FuncOffRequestExternal returns the function "WEBEXT.extension.onRequestExternal.removeListener".
func FuncOffRequestExternal() (fn js.Func[func(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])])]) {
	bindings.FuncOffRequestExternal(
		js.Pointer(&fn),
	)
	return
}

// OffRequestExternal calls the function "WEBEXT.extension.onRequestExternal.removeListener" directly.
func OffRequestExternal(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret js.Void) {
	bindings.CallOffRequestExternal(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRequestExternal calls the function "WEBEXT.extension.onRequestExternal.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRequestExternal(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRequestExternal(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRequestExternal returns true if the function "WEBEXT.extension.onRequestExternal.hasListener" exists.
func HasFuncHasOnRequestExternal() bool {
	return js.True == bindings.HasFuncHasOnRequestExternal()
}

// FuncHasOnRequestExternal returns the function "WEBEXT.extension.onRequestExternal.hasListener".
func FuncHasOnRequestExternal() (fn js.Func[func(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) bool]) {
	bindings.FuncHasOnRequestExternal(
		js.Pointer(&fn),
	)
	return
}

// HasOnRequestExternal calls the function "WEBEXT.extension.onRequestExternal.hasListener" directly.
func HasOnRequestExternal(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret bool) {
	bindings.CallHasOnRequestExternal(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRequestExternal calls the function "WEBEXT.extension.onRequestExternal.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRequestExternal(callback js.Func[func(request js.Any, sender *runtime.MessageSender, sendResponse js.Func[func()])]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRequestExternal(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSendRequest returns true if the function "WEBEXT.extension.sendRequest" exists.
func HasFuncSendRequest() bool {
	return js.True == bindings.HasFuncSendRequest()
}

// FuncSendRequest returns the function "WEBEXT.extension.sendRequest".
func FuncSendRequest() (fn js.Func[func(extensionId js.String, request js.Any) js.Promise[js.Any]]) {
	bindings.FuncSendRequest(
		js.Pointer(&fn),
	)
	return
}

// SendRequest calls the function "WEBEXT.extension.sendRequest" directly.
func SendRequest(extensionId js.String, request js.Any) (ret js.Promise[js.Any]) {
	bindings.CallSendRequest(
		js.Pointer(&ret),
		extensionId.Ref(),
		request.Ref(),
	)

	return
}

// TrySendRequest calls the function "WEBEXT.extension.sendRequest"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySendRequest(extensionId js.String, request js.Any) (ret js.Promise[js.Any], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySendRequest(
		js.Pointer(&ret), js.Pointer(&exception),
		extensionId.Ref(),
		request.Ref(),
	)

	return
}

// HasFuncSetUpdateUrlData returns true if the function "WEBEXT.extension.setUpdateUrlData" exists.
func HasFuncSetUpdateUrlData() bool {
	return js.True == bindings.HasFuncSetUpdateUrlData()
}

// FuncSetUpdateUrlData returns the function "WEBEXT.extension.setUpdateUrlData".
func FuncSetUpdateUrlData() (fn js.Func[func(data js.String)]) {
	bindings.FuncSetUpdateUrlData(
		js.Pointer(&fn),
	)
	return
}

// SetUpdateUrlData calls the function "WEBEXT.extension.setUpdateUrlData" directly.
func SetUpdateUrlData(data js.String) (ret js.Void) {
	bindings.CallSetUpdateUrlData(
		js.Pointer(&ret),
		data.Ref(),
	)

	return
}

// TrySetUpdateUrlData calls the function "WEBEXT.extension.setUpdateUrlData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetUpdateUrlData(data js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetUpdateUrlData(
		js.Pointer(&ret), js.Pointer(&exception),
		data.Ref(),
	)

	return
}

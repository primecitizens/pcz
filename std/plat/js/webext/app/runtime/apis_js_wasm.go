// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package runtime

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/app/runtime/bindings"
)

type ActionType uint32

const (
	_ ActionType = iota

	ActionType_NEW_NOTE
)

func (ActionType) FromRef(str js.Ref) ActionType {
	return ActionType(bindings.ConstOfActionType(str))
}

func (x ActionType) String() (string, bool) {
	switch x {
	case ActionType_NEW_NOTE:
		return "new_note", true
	default:
		return "", false
	}
}

type ActionData struct {
	// ActionType is "ActionData.actionType"
	//
	// Optional
	ActionType ActionType
	// IsLockScreenAction is "ActionData.isLockScreenAction"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsLockScreenAction MUST be set to true to make this field effective.
	IsLockScreenAction bool
	// RestoreLastActionState is "ActionData.restoreLastActionState"
	//
	// Optional
	//
	// NOTE: FFI_USE_RestoreLastActionState MUST be set to true to make this field effective.
	RestoreLastActionState bool

	FFI_USE_IsLockScreenAction     bool // for IsLockScreenAction.
	FFI_USE_RestoreLastActionState bool // for RestoreLastActionState.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ActionData with all fields set.
func (p ActionData) FromRef(ref js.Ref) ActionData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ActionData in the application heap.
func (p ActionData) New() js.Ref {
	return bindings.ActionDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ActionData) UpdateFrom(ref js.Ref) {
	bindings.ActionDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ActionData) Update(ref js.Ref) {
	bindings.ActionDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ActionData) FreeMembers(recursive bool) {
}

type EmbedRequest struct {
	ref js.Ref
}

func (this EmbedRequest) Once() EmbedRequest {
	this.ref.Once()
	return this
}

func (this EmbedRequest) Ref() js.Ref {
	return this.ref
}

func (this EmbedRequest) FromRef(ref js.Ref) EmbedRequest {
	this.ref = ref
	return this
}

func (this EmbedRequest) Free() {
	this.ref.Free()
}

// EmbedderId returns the value of property "EmbedRequest.embedderId".
//
// It returns ok=false if there is no such property.
func (this EmbedRequest) EmbedderId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetEmbedRequestEmbedderId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetEmbedderId sets the value of property "EmbedRequest.embedderId" to val.
//
// It returns false if the property cannot be set.
func (this EmbedRequest) SetEmbedderId(val js.String) bool {
	return js.True == bindings.SetEmbedRequestEmbedderId(
		this.ref,
		val.Ref(),
	)
}

// Data returns the value of property "EmbedRequest.data".
//
// It returns ok=false if there is no such property.
func (this EmbedRequest) Data() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetEmbedRequestData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetData sets the value of property "EmbedRequest.data" to val.
//
// It returns false if the property cannot be set.
func (this EmbedRequest) SetData(val js.Any) bool {
	return js.True == bindings.SetEmbedRequestData(
		this.ref,
		val.Ref(),
	)
}

// HasFuncAllow returns true if the method "EmbedRequest.allow" exists.
func (this EmbedRequest) HasFuncAllow() bool {
	return js.True == bindings.HasFuncEmbedRequestAllow(
		this.ref,
	)
}

// FuncAllow returns the method "EmbedRequest.allow".
func (this EmbedRequest) FuncAllow() (fn js.Func[func(url js.String)]) {
	bindings.FuncEmbedRequestAllow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Allow calls the method "EmbedRequest.allow".
func (this EmbedRequest) Allow(url js.String) (ret js.Void) {
	bindings.CallEmbedRequestAllow(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryAllow calls the method "EmbedRequest.allow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EmbedRequest) TryAllow(url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEmbedRequestAllow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncDeny returns true if the method "EmbedRequest.deny" exists.
func (this EmbedRequest) HasFuncDeny() bool {
	return js.True == bindings.HasFuncEmbedRequestDeny(
		this.ref,
	)
}

// FuncDeny returns the method "EmbedRequest.deny".
func (this EmbedRequest) FuncDeny() (fn js.Func[func()]) {
	bindings.FuncEmbedRequestDeny(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Deny calls the method "EmbedRequest.deny".
func (this EmbedRequest) Deny() (ret js.Void) {
	bindings.CallEmbedRequestDeny(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeny calls the method "EmbedRequest.deny"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EmbedRequest) TryDeny() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEmbedRequestDeny(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type LaunchItem struct {
	// Entry is "LaunchItem.entry"
	//
	// Optional
	Entry js.Object
	// Type is "LaunchItem.type"
	//
	// Optional
	Type js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LaunchItem with all fields set.
func (p LaunchItem) FromRef(ref js.Ref) LaunchItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LaunchItem in the application heap.
func (p LaunchItem) New() js.Ref {
	return bindings.LaunchItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LaunchItem) UpdateFrom(ref js.Ref) {
	bindings.LaunchItemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LaunchItem) Update(ref js.Ref) {
	bindings.LaunchItemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LaunchItem) FreeMembers(recursive bool) {
	js.Free(
		p.Entry.Ref(),
		p.Type.Ref(),
	)
	p.Entry = p.Entry.FromRef(js.Undefined)
	p.Type = p.Type.FromRef(js.Undefined)
}

type LaunchSource uint32

const (
	_ LaunchSource = iota

	LaunchSource_UNTRACKED
	LaunchSource_APP_LAUNCHER
	LaunchSource_NEW_TAB_PAGE
	LaunchSource_RELOAD
	LaunchSource_RESTART
	LaunchSource_LOAD_AND_LAUNCH
	LaunchSource_COMMAND_LINE
	LaunchSource_FILE_HANDLER
	LaunchSource_URL_HANDLER
	LaunchSource_SYSTEM_TRAY
	LaunchSource_ABOUT_PAGE
	LaunchSource_KEYBOARD
	LaunchSource_EXTENSIONS_PAGE
	LaunchSource_MANAGEMENT_API
	LaunchSource_EPHEMERAL_APP
	LaunchSource_BACKGROUND
	LaunchSource_KIOSK
	LaunchSource_CHROME_INTERNAL
	LaunchSource_TEST
	LaunchSource_INSTALLED_NOTIFICATION
	LaunchSource_CONTEXT_MENU
	LaunchSource_ARC
	LaunchSource_INTENT_URL
	LaunchSource_APP_HOME_PAGE
)

func (LaunchSource) FromRef(str js.Ref) LaunchSource {
	return LaunchSource(bindings.ConstOfLaunchSource(str))
}

func (x LaunchSource) String() (string, bool) {
	switch x {
	case LaunchSource_UNTRACKED:
		return "untracked", true
	case LaunchSource_APP_LAUNCHER:
		return "app_launcher", true
	case LaunchSource_NEW_TAB_PAGE:
		return "new_tab_page", true
	case LaunchSource_RELOAD:
		return "reload", true
	case LaunchSource_RESTART:
		return "restart", true
	case LaunchSource_LOAD_AND_LAUNCH:
		return "load_and_launch", true
	case LaunchSource_COMMAND_LINE:
		return "command_line", true
	case LaunchSource_FILE_HANDLER:
		return "file_handler", true
	case LaunchSource_URL_HANDLER:
		return "url_handler", true
	case LaunchSource_SYSTEM_TRAY:
		return "system_tray", true
	case LaunchSource_ABOUT_PAGE:
		return "about_page", true
	case LaunchSource_KEYBOARD:
		return "keyboard", true
	case LaunchSource_EXTENSIONS_PAGE:
		return "extensions_page", true
	case LaunchSource_MANAGEMENT_API:
		return "management_api", true
	case LaunchSource_EPHEMERAL_APP:
		return "ephemeral_app", true
	case LaunchSource_BACKGROUND:
		return "background", true
	case LaunchSource_KIOSK:
		return "kiosk", true
	case LaunchSource_CHROME_INTERNAL:
		return "chrome_internal", true
	case LaunchSource_TEST:
		return "test", true
	case LaunchSource_INSTALLED_NOTIFICATION:
		return "installed_notification", true
	case LaunchSource_CONTEXT_MENU:
		return "context_menu", true
	case LaunchSource_ARC:
		return "arc", true
	case LaunchSource_INTENT_URL:
		return "intent_url", true
	case LaunchSource_APP_HOME_PAGE:
		return "app_home_page", true
	default:
		return "", false
	}
}

type LaunchData struct {
	// Id is "LaunchData.id"
	//
	// Optional
	Id js.String
	// Items is "LaunchData.items"
	//
	// Optional
	Items js.Array[LaunchItem]
	// Url is "LaunchData.url"
	//
	// Optional
	Url js.String
	// ReferrerUrl is "LaunchData.referrerUrl"
	//
	// Optional
	ReferrerUrl js.String
	// IsDemoSession is "LaunchData.isDemoSession"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsDemoSession MUST be set to true to make this field effective.
	IsDemoSession bool
	// IsKioskSession is "LaunchData.isKioskSession"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsKioskSession MUST be set to true to make this field effective.
	IsKioskSession bool
	// IsPublicSession is "LaunchData.isPublicSession"
	//
	// Optional
	//
	// NOTE: FFI_USE_IsPublicSession MUST be set to true to make this field effective.
	IsPublicSession bool
	// Source is "LaunchData.source"
	//
	// Optional
	Source LaunchSource
	// ActionData is "LaunchData.actionData"
	//
	// Optional
	//
	// NOTE: ActionData.FFI_USE MUST be set to true to get ActionData used.
	ActionData ActionData

	FFI_USE_IsDemoSession   bool // for IsDemoSession.
	FFI_USE_IsKioskSession  bool // for IsKioskSession.
	FFI_USE_IsPublicSession bool // for IsPublicSession.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LaunchData with all fields set.
func (p LaunchData) FromRef(ref js.Ref) LaunchData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LaunchData in the application heap.
func (p LaunchData) New() js.Ref {
	return bindings.LaunchDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *LaunchData) UpdateFrom(ref js.Ref) {
	bindings.LaunchDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *LaunchData) Update(ref js.Ref) {
	bindings.LaunchDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *LaunchData) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Items.Ref(),
		p.Url.Ref(),
		p.ReferrerUrl.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Items = p.Items.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.ReferrerUrl = p.ReferrerUrl.FromRef(js.Undefined)
	if recursive {
		p.ActionData.FreeMembers(true)
	}
}

type OnEmbedRequestedEventCallbackFunc func(this js.Ref, request EmbedRequest) js.Ref

func (fn OnEmbedRequestedEventCallbackFunc) Register() js.Func[func(request EmbedRequest)] {
	return js.RegisterCallback[func(request EmbedRequest)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnEmbedRequestedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		EmbedRequest{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OnEmbedRequestedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, request EmbedRequest) js.Ref
	Arg T
}

func (cb *OnEmbedRequestedEventCallback[T]) Register() js.Func[func(request EmbedRequest)] {
	return js.RegisterCallback[func(request EmbedRequest)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnEmbedRequestedEventCallback[T]) DispatchCallback(
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

		EmbedRequest{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnEmbedRequested returns true if the function "WEBEXT.app.runtime.onEmbedRequested.addListener" exists.
func HasFuncOnEmbedRequested() bool {
	return js.True == bindings.HasFuncOnEmbedRequested()
}

// FuncOnEmbedRequested returns the function "WEBEXT.app.runtime.onEmbedRequested.addListener".
func FuncOnEmbedRequested() (fn js.Func[func(callback js.Func[func(request EmbedRequest)])]) {
	bindings.FuncOnEmbedRequested(
		js.Pointer(&fn),
	)
	return
}

// OnEmbedRequested calls the function "WEBEXT.app.runtime.onEmbedRequested.addListener" directly.
func OnEmbedRequested(callback js.Func[func(request EmbedRequest)]) (ret js.Void) {
	bindings.CallOnEmbedRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnEmbedRequested calls the function "WEBEXT.app.runtime.onEmbedRequested.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnEmbedRequested(callback js.Func[func(request EmbedRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnEmbedRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffEmbedRequested returns true if the function "WEBEXT.app.runtime.onEmbedRequested.removeListener" exists.
func HasFuncOffEmbedRequested() bool {
	return js.True == bindings.HasFuncOffEmbedRequested()
}

// FuncOffEmbedRequested returns the function "WEBEXT.app.runtime.onEmbedRequested.removeListener".
func FuncOffEmbedRequested() (fn js.Func[func(callback js.Func[func(request EmbedRequest)])]) {
	bindings.FuncOffEmbedRequested(
		js.Pointer(&fn),
	)
	return
}

// OffEmbedRequested calls the function "WEBEXT.app.runtime.onEmbedRequested.removeListener" directly.
func OffEmbedRequested(callback js.Func[func(request EmbedRequest)]) (ret js.Void) {
	bindings.CallOffEmbedRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffEmbedRequested calls the function "WEBEXT.app.runtime.onEmbedRequested.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffEmbedRequested(callback js.Func[func(request EmbedRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffEmbedRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnEmbedRequested returns true if the function "WEBEXT.app.runtime.onEmbedRequested.hasListener" exists.
func HasFuncHasOnEmbedRequested() bool {
	return js.True == bindings.HasFuncHasOnEmbedRequested()
}

// FuncHasOnEmbedRequested returns the function "WEBEXT.app.runtime.onEmbedRequested.hasListener".
func FuncHasOnEmbedRequested() (fn js.Func[func(callback js.Func[func(request EmbedRequest)]) bool]) {
	bindings.FuncHasOnEmbedRequested(
		js.Pointer(&fn),
	)
	return
}

// HasOnEmbedRequested calls the function "WEBEXT.app.runtime.onEmbedRequested.hasListener" directly.
func HasOnEmbedRequested(callback js.Func[func(request EmbedRequest)]) (ret bool) {
	bindings.CallHasOnEmbedRequested(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnEmbedRequested calls the function "WEBEXT.app.runtime.onEmbedRequested.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnEmbedRequested(callback js.Func[func(request EmbedRequest)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnEmbedRequested(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnLaunchedEventCallbackFunc func(this js.Ref, launchData *LaunchData) js.Ref

func (fn OnLaunchedEventCallbackFunc) Register() js.Func[func(launchData *LaunchData)] {
	return js.RegisterCallback[func(launchData *LaunchData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnLaunchedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LaunchData
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

type OnLaunchedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, launchData *LaunchData) js.Ref
	Arg T
}

func (cb *OnLaunchedEventCallback[T]) Register() js.Func[func(launchData *LaunchData)] {
	return js.RegisterCallback[func(launchData *LaunchData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnLaunchedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 LaunchData
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

// HasFuncOnLaunched returns true if the function "WEBEXT.app.runtime.onLaunched.addListener" exists.
func HasFuncOnLaunched() bool {
	return js.True == bindings.HasFuncOnLaunched()
}

// FuncOnLaunched returns the function "WEBEXT.app.runtime.onLaunched.addListener".
func FuncOnLaunched() (fn js.Func[func(callback js.Func[func(launchData *LaunchData)])]) {
	bindings.FuncOnLaunched(
		js.Pointer(&fn),
	)
	return
}

// OnLaunched calls the function "WEBEXT.app.runtime.onLaunched.addListener" directly.
func OnLaunched(callback js.Func[func(launchData *LaunchData)]) (ret js.Void) {
	bindings.CallOnLaunched(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnLaunched calls the function "WEBEXT.app.runtime.onLaunched.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnLaunched(callback js.Func[func(launchData *LaunchData)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnLaunched(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffLaunched returns true if the function "WEBEXT.app.runtime.onLaunched.removeListener" exists.
func HasFuncOffLaunched() bool {
	return js.True == bindings.HasFuncOffLaunched()
}

// FuncOffLaunched returns the function "WEBEXT.app.runtime.onLaunched.removeListener".
func FuncOffLaunched() (fn js.Func[func(callback js.Func[func(launchData *LaunchData)])]) {
	bindings.FuncOffLaunched(
		js.Pointer(&fn),
	)
	return
}

// OffLaunched calls the function "WEBEXT.app.runtime.onLaunched.removeListener" directly.
func OffLaunched(callback js.Func[func(launchData *LaunchData)]) (ret js.Void) {
	bindings.CallOffLaunched(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffLaunched calls the function "WEBEXT.app.runtime.onLaunched.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffLaunched(callback js.Func[func(launchData *LaunchData)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffLaunched(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnLaunched returns true if the function "WEBEXT.app.runtime.onLaunched.hasListener" exists.
func HasFuncHasOnLaunched() bool {
	return js.True == bindings.HasFuncHasOnLaunched()
}

// FuncHasOnLaunched returns the function "WEBEXT.app.runtime.onLaunched.hasListener".
func FuncHasOnLaunched() (fn js.Func[func(callback js.Func[func(launchData *LaunchData)]) bool]) {
	bindings.FuncHasOnLaunched(
		js.Pointer(&fn),
	)
	return
}

// HasOnLaunched calls the function "WEBEXT.app.runtime.onLaunched.hasListener" directly.
func HasOnLaunched(callback js.Func[func(launchData *LaunchData)]) (ret bool) {
	bindings.CallHasOnLaunched(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnLaunched calls the function "WEBEXT.app.runtime.onLaunched.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnLaunched(callback js.Func[func(launchData *LaunchData)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnLaunched(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnRestartedEventCallbackFunc func(this js.Ref) js.Ref

func (fn OnRestartedEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnRestartedEventCallbackFunc) DispatchCallback(
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

type OnRestartedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *OnRestartedEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnRestartedEventCallback[T]) DispatchCallback(
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

// HasFuncOnRestarted returns true if the function "WEBEXT.app.runtime.onRestarted.addListener" exists.
func HasFuncOnRestarted() bool {
	return js.True == bindings.HasFuncOnRestarted()
}

// FuncOnRestarted returns the function "WEBEXT.app.runtime.onRestarted.addListener".
func FuncOnRestarted() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnRestarted(
		js.Pointer(&fn),
	)
	return
}

// OnRestarted calls the function "WEBEXT.app.runtime.onRestarted.addListener" directly.
func OnRestarted(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnRestarted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnRestarted calls the function "WEBEXT.app.runtime.onRestarted.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnRestarted(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnRestarted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffRestarted returns true if the function "WEBEXT.app.runtime.onRestarted.removeListener" exists.
func HasFuncOffRestarted() bool {
	return js.True == bindings.HasFuncOffRestarted()
}

// FuncOffRestarted returns the function "WEBEXT.app.runtime.onRestarted.removeListener".
func FuncOffRestarted() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffRestarted(
		js.Pointer(&fn),
	)
	return
}

// OffRestarted calls the function "WEBEXT.app.runtime.onRestarted.removeListener" directly.
func OffRestarted(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffRestarted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffRestarted calls the function "WEBEXT.app.runtime.onRestarted.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffRestarted(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffRestarted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnRestarted returns true if the function "WEBEXT.app.runtime.onRestarted.hasListener" exists.
func HasFuncHasOnRestarted() bool {
	return js.True == bindings.HasFuncHasOnRestarted()
}

// FuncHasOnRestarted returns the function "WEBEXT.app.runtime.onRestarted.hasListener".
func FuncHasOnRestarted() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnRestarted(
		js.Pointer(&fn),
	)
	return
}

// HasOnRestarted calls the function "WEBEXT.app.runtime.onRestarted.hasListener" directly.
func HasOnRestarted(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnRestarted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnRestarted calls the function "WEBEXT.app.runtime.onRestarted.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnRestarted(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnRestarted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

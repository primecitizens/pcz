// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

type NavigationInterceptHandlerFunc func(this js.Ref) js.Ref

func (fn NavigationInterceptHandlerFunc) Register() js.Func[func() js.Promise[js.Void]] {
	return js.RegisterCallback[func() js.Promise[js.Void]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn NavigationInterceptHandlerFunc) DispatchCallback(
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

type NavigationInterceptHandler[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *NavigationInterceptHandler[T]) Register() js.Func[func() js.Promise[js.Void]] {
	return js.RegisterCallback[func() js.Promise[js.Void]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *NavigationInterceptHandler[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 0+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NavigationFocusReset uint32

const (
	_ NavigationFocusReset = iota

	NavigationFocusReset_AFTER_TRANSITION
	NavigationFocusReset_MANUAL
)

func (NavigationFocusReset) FromRef(str js.Ref) NavigationFocusReset {
	return NavigationFocusReset(bindings.ConstOfNavigationFocusReset(str))
}

func (x NavigationFocusReset) String() (string, bool) {
	switch x {
	case NavigationFocusReset_AFTER_TRANSITION:
		return "after-transition", true
	case NavigationFocusReset_MANUAL:
		return "manual", true
	default:
		return "", false
	}
}

type NavigationScrollBehavior uint32

const (
	_ NavigationScrollBehavior = iota

	NavigationScrollBehavior_AFTER_TRANSITION
	NavigationScrollBehavior_MANUAL
)

func (NavigationScrollBehavior) FromRef(str js.Ref) NavigationScrollBehavior {
	return NavigationScrollBehavior(bindings.ConstOfNavigationScrollBehavior(str))
}

func (x NavigationScrollBehavior) String() (string, bool) {
	switch x {
	case NavigationScrollBehavior_AFTER_TRANSITION:
		return "after-transition", true
	case NavigationScrollBehavior_MANUAL:
		return "manual", true
	default:
		return "", false
	}
}

type NavigationInterceptOptions struct {
	// Handler is "NavigationInterceptOptions.handler"
	//
	// Optional
	Handler js.Func[func() js.Promise[js.Void]]
	// FocusReset is "NavigationInterceptOptions.focusReset"
	//
	// Optional
	FocusReset NavigationFocusReset
	// Scroll is "NavigationInterceptOptions.scroll"
	//
	// Optional
	Scroll NavigationScrollBehavior

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NavigationInterceptOptions with all fields set.
func (p NavigationInterceptOptions) FromRef(ref js.Ref) NavigationInterceptOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NavigationInterceptOptions NavigationInterceptOptions [// NavigationInterceptOptions] [0x14000934640 0x140009346e0 0x14000934780] 0x140009203f0 {0 0}} in the application heap.
func (p NavigationInterceptOptions) New() js.Ref {
	return bindings.NavigationInterceptOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NavigationInterceptOptions) UpdateFrom(ref js.Ref) {
	bindings.NavigationInterceptOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NavigationInterceptOptions) Update(ref js.Ref) {
	bindings.NavigationInterceptOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewNavigateEvent(typ js.String, eventInitDict NavigateEventInit) NavigateEvent {
	return NavigateEvent{}.FromRef(
		bindings.NewNavigateEventByNavigateEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type NavigateEvent struct {
	Event
}

func (this NavigateEvent) Once() NavigateEvent {
	this.Ref().Once()
	return this
}

func (this NavigateEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this NavigateEvent) FromRef(ref js.Ref) NavigateEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this NavigateEvent) Free() {
	this.Ref().Free()
}

// NavigationType returns the value of property "NavigateEvent.navigationType".
//
// The returned bool will be false if there is no such property.
func (this NavigateEvent) NavigationType() (NavigationType, bool) {
	var _ok bool
	_ret := bindings.GetNavigateEventNavigationType(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigationType(_ret), _ok
}

// Destination returns the value of property "NavigateEvent.destination".
//
// The returned bool will be false if there is no such property.
func (this NavigateEvent) Destination() (NavigationDestination, bool) {
	var _ok bool
	_ret := bindings.GetNavigateEventDestination(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigationDestination{}.FromRef(_ret), _ok
}

// CanIntercept returns the value of property "NavigateEvent.canIntercept".
//
// The returned bool will be false if there is no such property.
func (this NavigateEvent) CanIntercept() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigateEventCanIntercept(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// UserInitiated returns the value of property "NavigateEvent.userInitiated".
//
// The returned bool will be false if there is no such property.
func (this NavigateEvent) UserInitiated() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigateEventUserInitiated(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// HashChange returns the value of property "NavigateEvent.hashChange".
//
// The returned bool will be false if there is no such property.
func (this NavigateEvent) HashChange() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigateEventHashChange(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Signal returns the value of property "NavigateEvent.signal".
//
// The returned bool will be false if there is no such property.
func (this NavigateEvent) Signal() (AbortSignal, bool) {
	var _ok bool
	_ret := bindings.GetNavigateEventSignal(
		this.Ref(), js.Pointer(&_ok),
	)
	return AbortSignal{}.FromRef(_ret), _ok
}

// FormData returns the value of property "NavigateEvent.formData".
//
// The returned bool will be false if there is no such property.
func (this NavigateEvent) FormData() (FormData, bool) {
	var _ok bool
	_ret := bindings.GetNavigateEventFormData(
		this.Ref(), js.Pointer(&_ok),
	)
	return FormData{}.FromRef(_ret), _ok
}

// DownloadRequest returns the value of property "NavigateEvent.downloadRequest".
//
// The returned bool will be false if there is no such property.
func (this NavigateEvent) DownloadRequest() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNavigateEventDownloadRequest(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Info returns the value of property "NavigateEvent.info".
//
// The returned bool will be false if there is no such property.
func (this NavigateEvent) Info() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetNavigateEventInfo(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// HasUAVisualTransition returns the value of property "NavigateEvent.hasUAVisualTransition".
//
// The returned bool will be false if there is no such property.
func (this NavigateEvent) HasUAVisualTransition() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigateEventHasUAVisualTransition(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Intercept calls the method "NavigateEvent.intercept".
//
// The returned bool will be false if there is no such method.
func (this NavigateEvent) Intercept(options NavigationInterceptOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallNavigateEventIntercept(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InterceptFunc returns the method "NavigateEvent.intercept".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NavigateEvent) InterceptFunc() (fn js.Func[func(options NavigationInterceptOptions)]) {
	return fn.FromRef(
		bindings.NavigateEventInterceptFunc(
			this.Ref(),
		),
	)
}

// Intercept1 calls the method "NavigateEvent.intercept".
//
// The returned bool will be false if there is no such method.
func (this NavigateEvent) Intercept1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallNavigateEventIntercept1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Intercept1Func returns the method "NavigateEvent.intercept".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NavigateEvent) Intercept1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.NavigateEventIntercept1Func(
			this.Ref(),
		),
	)
}

// Scroll calls the method "NavigateEvent.scroll".
//
// The returned bool will be false if there is no such method.
func (this NavigateEvent) Scroll() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallNavigateEventScroll(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScrollFunc returns the method "NavigateEvent.scroll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this NavigateEvent) ScrollFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.NavigateEventScrollFunc(
			this.Ref(),
		),
	)
}

type NavigationCurrentEntryChangeEventInit struct {
	// NavigationType is "NavigationCurrentEntryChangeEventInit.navigationType"
	//
	// Optional, defaults to null.
	NavigationType NavigationType
	// From is "NavigationCurrentEntryChangeEventInit.from"
	//
	// Required
	From NavigationHistoryEntry
	// Bubbles is "NavigationCurrentEntryChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "NavigationCurrentEntryChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "NavigationCurrentEntryChangeEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NavigationCurrentEntryChangeEventInit with all fields set.
func (p NavigationCurrentEntryChangeEventInit) FromRef(ref js.Ref) NavigationCurrentEntryChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NavigationCurrentEntryChangeEventInit NavigationCurrentEntryChangeEventInit [// NavigationCurrentEntryChangeEventInit] [0x140009348c0 0x14000934960 0x14000934a00 0x14000934b40 0x14000934c80 0x14000934aa0 0x14000934be0 0x14000934d20] 0x14000920498 {0 0}} in the application heap.
func (p NavigationCurrentEntryChangeEventInit) New() js.Ref {
	return bindings.NavigationCurrentEntryChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NavigationCurrentEntryChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.NavigationCurrentEntryChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NavigationCurrentEntryChangeEventInit) Update(ref js.Ref) {
	bindings.NavigationCurrentEntryChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewNavigationCurrentEntryChangeEvent(typ js.String, eventInitDict NavigationCurrentEntryChangeEventInit) NavigationCurrentEntryChangeEvent {
	return NavigationCurrentEntryChangeEvent{}.FromRef(
		bindings.NewNavigationCurrentEntryChangeEventByNavigationCurrentEntryChangeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type NavigationCurrentEntryChangeEvent struct {
	Event
}

func (this NavigationCurrentEntryChangeEvent) Once() NavigationCurrentEntryChangeEvent {
	this.Ref().Once()
	return this
}

func (this NavigationCurrentEntryChangeEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this NavigationCurrentEntryChangeEvent) FromRef(ref js.Ref) NavigationCurrentEntryChangeEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this NavigationCurrentEntryChangeEvent) Free() {
	this.Ref().Free()
}

// NavigationType returns the value of property "NavigationCurrentEntryChangeEvent.navigationType".
//
// The returned bool will be false if there is no such property.
func (this NavigationCurrentEntryChangeEvent) NavigationType() (NavigationType, bool) {
	var _ok bool
	_ret := bindings.GetNavigationCurrentEntryChangeEventNavigationType(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigationType(_ret), _ok
}

// From returns the value of property "NavigationCurrentEntryChangeEvent.from".
//
// The returned bool will be false if there is no such property.
func (this NavigationCurrentEntryChangeEvent) From() (NavigationHistoryEntry, bool) {
	var _ok bool
	_ret := bindings.GetNavigationCurrentEntryChangeEventFrom(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigationHistoryEntry{}.FromRef(_ret), _ok
}

type NavigationEventInit struct {
	// Dir is "NavigationEventInit.dir"
	//
	// Optional
	Dir SpatialNavigationDirection
	// RelatedTarget is "NavigationEventInit.relatedTarget"
	//
	// Optional, defaults to null.
	RelatedTarget EventTarget
	// View is "NavigationEventInit.view"
	//
	// Optional, defaults to null.
	View Window
	// Detail is "NavigationEventInit.detail"
	//
	// Optional, defaults to 0.
	Detail int32
	// Bubbles is "NavigationEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "NavigationEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "NavigationEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Detail     bool // for Detail.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NavigationEventInit with all fields set.
func (p NavigationEventInit) FromRef(ref js.Ref) NavigationEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NavigationEventInit NavigationEventInit [// NavigationEventInit] [0x14000934e60 0x14000934f00 0x14000934fa0 0x14000935040 0x14000935180 0x140009352c0 0x14000935400 0x140009350e0 0x14000935220 0x14000935360 0x140009354a0] 0x14000920510 {0 0}} in the application heap.
func (p NavigationEventInit) New() js.Ref {
	return bindings.NavigationEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NavigationEventInit) UpdateFrom(ref js.Ref) {
	bindings.NavigationEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NavigationEventInit) Update(ref js.Ref) {
	bindings.NavigationEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewNavigationEvent(typ js.String, eventInitDict NavigationEventInit) NavigationEvent {
	return NavigationEvent{}.FromRef(
		bindings.NewNavigationEventByNavigationEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewNavigationEventByNavigationEvent1(typ js.String) NavigationEvent {
	return NavigationEvent{}.FromRef(
		bindings.NewNavigationEventByNavigationEvent1(
			typ.Ref()),
	)
}

type NavigationEvent struct {
	UIEvent
}

func (this NavigationEvent) Once() NavigationEvent {
	this.Ref().Once()
	return this
}

func (this NavigationEvent) Ref() js.Ref {
	return this.UIEvent.Ref()
}

func (this NavigationEvent) FromRef(ref js.Ref) NavigationEvent {
	this.UIEvent = this.UIEvent.FromRef(ref)
	return this
}

func (this NavigationEvent) Free() {
	this.Ref().Free()
}

// Dir returns the value of property "NavigationEvent.dir".
//
// The returned bool will be false if there is no such property.
func (this NavigationEvent) Dir() (SpatialNavigationDirection, bool) {
	var _ok bool
	_ret := bindings.GetNavigationEventDir(
		this.Ref(), js.Pointer(&_ok),
	)
	return SpatialNavigationDirection(_ret), _ok
}

// RelatedTarget returns the value of property "NavigationEvent.relatedTarget".
//
// The returned bool will be false if there is no such property.
func (this NavigationEvent) RelatedTarget() (EventTarget, bool) {
	var _ok bool
	_ret := bindings.GetNavigationEventRelatedTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return EventTarget{}.FromRef(_ret), _ok
}

type NavigationTimingType uint32

const (
	_ NavigationTimingType = iota

	NavigationTimingType_NAVIGATE
	NavigationTimingType_RELOAD
	NavigationTimingType_BACK_FORWARD
	NavigationTimingType_PRERENDER
)

func (NavigationTimingType) FromRef(str js.Ref) NavigationTimingType {
	return NavigationTimingType(bindings.ConstOfNavigationTimingType(str))
}

func (x NavigationTimingType) String() (string, bool) {
	switch x {
	case NavigationTimingType_NAVIGATE:
		return "navigate", true
	case NavigationTimingType_RELOAD:
		return "reload", true
	case NavigationTimingType_BACK_FORWARD:
		return "back_forward", true
	case NavigationTimingType_PRERENDER:
		return "prerender", true
	default:
		return "", false
	}
}

type NotificationEventInit struct {
	// Notification is "NotificationEventInit.notification"
	//
	// Required
	Notification Notification
	// Action is "NotificationEventInit.action"
	//
	// Optional, defaults to "".
	Action js.String
	// Bubbles is "NotificationEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "NotificationEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "NotificationEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NotificationEventInit with all fields set.
func (p NotificationEventInit) FromRef(ref js.Ref) NotificationEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 NotificationEventInit NotificationEventInit [// NotificationEventInit] [0x140009355e0 0x14000935680 0x14000935720 0x14000935860 0x140009359a0 0x140009357c0 0x14000935900 0x14000935a40] 0x140009205a0 {0 0}} in the application heap.
func (p NotificationEventInit) New() js.Ref {
	return bindings.NotificationEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NotificationEventInit) UpdateFrom(ref js.Ref) {
	bindings.NotificationEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NotificationEventInit) Update(ref js.Ref) {
	bindings.NotificationEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewNotificationEvent(typ js.String, eventInitDict NotificationEventInit) NotificationEvent {
	return NotificationEvent{}.FromRef(
		bindings.NewNotificationEventByNotificationEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type NotificationEvent struct {
	ExtendableEvent
}

func (this NotificationEvent) Once() NotificationEvent {
	this.Ref().Once()
	return this
}

func (this NotificationEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this NotificationEvent) FromRef(ref js.Ref) NotificationEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this NotificationEvent) Free() {
	this.Ref().Free()
}

// Notification returns the value of property "NotificationEvent.notification".
//
// The returned bool will be false if there is no such property.
func (this NotificationEvent) Notification() (Notification, bool) {
	var _ok bool
	_ret := bindings.GetNotificationEventNotification(
		this.Ref(), js.Pointer(&_ok),
	)
	return Notification{}.FromRef(_ret), _ok
}

// Action returns the value of property "NotificationEvent.action".
//
// The returned bool will be false if there is no such property.
func (this NotificationEvent) Action() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetNotificationEventAction(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type OES_draw_buffers_indexed struct {
	ref js.Ref
}

func (this OES_draw_buffers_indexed) Once() OES_draw_buffers_indexed {
	this.Ref().Once()
	return this
}

func (this OES_draw_buffers_indexed) Ref() js.Ref {
	return this.ref
}

func (this OES_draw_buffers_indexed) FromRef(ref js.Ref) OES_draw_buffers_indexed {
	this.ref = ref
	return this
}

func (this OES_draw_buffers_indexed) Free() {
	this.Ref().Free()
}

// EnableiOES calls the method "OES_draw_buffers_indexed.enableiOES".
//
// The returned bool will be false if there is no such method.
func (this OES_draw_buffers_indexed) EnableiOES(target GLenum, index GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOES_draw_buffers_indexedEnableiOES(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EnableiOESFunc returns the method "OES_draw_buffers_indexed.enableiOES".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OES_draw_buffers_indexed) EnableiOESFunc() (fn js.Func[func(target GLenum, index GLuint)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedEnableiOESFunc(
			this.Ref(),
		),
	)
}

// DisableiOES calls the method "OES_draw_buffers_indexed.disableiOES".
//
// The returned bool will be false if there is no such method.
func (this OES_draw_buffers_indexed) DisableiOES(target GLenum, index GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOES_draw_buffers_indexedDisableiOES(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisableiOESFunc returns the method "OES_draw_buffers_indexed.disableiOES".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OES_draw_buffers_indexed) DisableiOESFunc() (fn js.Func[func(target GLenum, index GLuint)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedDisableiOESFunc(
			this.Ref(),
		),
	)
}

// BlendEquationiOES calls the method "OES_draw_buffers_indexed.blendEquationiOES".
//
// The returned bool will be false if there is no such method.
func (this OES_draw_buffers_indexed) BlendEquationiOES(buf GLuint, mode GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOES_draw_buffers_indexedBlendEquationiOES(
		this.Ref(), js.Pointer(&_ok),
		uint32(buf),
		uint32(mode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendEquationiOESFunc returns the method "OES_draw_buffers_indexed.blendEquationiOES".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OES_draw_buffers_indexed) BlendEquationiOESFunc() (fn js.Func[func(buf GLuint, mode GLenum)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedBlendEquationiOESFunc(
			this.Ref(),
		),
	)
}

// BlendEquationSeparateiOES calls the method "OES_draw_buffers_indexed.blendEquationSeparateiOES".
//
// The returned bool will be false if there is no such method.
func (this OES_draw_buffers_indexed) BlendEquationSeparateiOES(buf GLuint, modeRGB GLenum, modeAlpha GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOES_draw_buffers_indexedBlendEquationSeparateiOES(
		this.Ref(), js.Pointer(&_ok),
		uint32(buf),
		uint32(modeRGB),
		uint32(modeAlpha),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendEquationSeparateiOESFunc returns the method "OES_draw_buffers_indexed.blendEquationSeparateiOES".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OES_draw_buffers_indexed) BlendEquationSeparateiOESFunc() (fn js.Func[func(buf GLuint, modeRGB GLenum, modeAlpha GLenum)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedBlendEquationSeparateiOESFunc(
			this.Ref(),
		),
	)
}

// BlendFunciOES calls the method "OES_draw_buffers_indexed.blendFunciOES".
//
// The returned bool will be false if there is no such method.
func (this OES_draw_buffers_indexed) BlendFunciOES(buf GLuint, src GLenum, dst GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOES_draw_buffers_indexedBlendFunciOES(
		this.Ref(), js.Pointer(&_ok),
		uint32(buf),
		uint32(src),
		uint32(dst),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendFunciOESFunc returns the method "OES_draw_buffers_indexed.blendFunciOES".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OES_draw_buffers_indexed) BlendFunciOESFunc() (fn js.Func[func(buf GLuint, src GLenum, dst GLenum)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedBlendFunciOESFunc(
			this.Ref(),
		),
	)
}

// BlendFuncSeparateiOES calls the method "OES_draw_buffers_indexed.blendFuncSeparateiOES".
//
// The returned bool will be false if there is no such method.
func (this OES_draw_buffers_indexed) BlendFuncSeparateiOES(buf GLuint, srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOES_draw_buffers_indexedBlendFuncSeparateiOES(
		this.Ref(), js.Pointer(&_ok),
		uint32(buf),
		uint32(srcRGB),
		uint32(dstRGB),
		uint32(srcAlpha),
		uint32(dstAlpha),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BlendFuncSeparateiOESFunc returns the method "OES_draw_buffers_indexed.blendFuncSeparateiOES".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OES_draw_buffers_indexed) BlendFuncSeparateiOESFunc() (fn js.Func[func(buf GLuint, srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedBlendFuncSeparateiOESFunc(
			this.Ref(),
		),
	)
}

// ColorMaskiOES calls the method "OES_draw_buffers_indexed.colorMaskiOES".
//
// The returned bool will be false if there is no such method.
func (this OES_draw_buffers_indexed) ColorMaskiOES(buf GLuint, r GLboolean, g GLboolean, b GLboolean, a GLboolean) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOES_draw_buffers_indexedColorMaskiOES(
		this.Ref(), js.Pointer(&_ok),
		uint32(buf),
		js.Bool(bool(r)),
		js.Bool(bool(g)),
		js.Bool(bool(b)),
		js.Bool(bool(a)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ColorMaskiOESFunc returns the method "OES_draw_buffers_indexed.colorMaskiOES".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OES_draw_buffers_indexed) ColorMaskiOESFunc() (fn js.Func[func(buf GLuint, r GLboolean, g GLboolean, b GLboolean, a GLboolean)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedColorMaskiOESFunc(
			this.Ref(),
		),
	)
}

type OES_element_index_uint struct {
	ref js.Ref
}

func (this OES_element_index_uint) Once() OES_element_index_uint {
	this.Ref().Once()
	return this
}

func (this OES_element_index_uint) Ref() js.Ref {
	return this.ref
}

func (this OES_element_index_uint) FromRef(ref js.Ref) OES_element_index_uint {
	this.ref = ref
	return this
}

func (this OES_element_index_uint) Free() {
	this.Ref().Free()
}

type OES_fbo_render_mipmap struct {
	ref js.Ref
}

func (this OES_fbo_render_mipmap) Once() OES_fbo_render_mipmap {
	this.Ref().Once()
	return this
}

func (this OES_fbo_render_mipmap) Ref() js.Ref {
	return this.ref
}

func (this OES_fbo_render_mipmap) FromRef(ref js.Ref) OES_fbo_render_mipmap {
	this.ref = ref
	return this
}

func (this OES_fbo_render_mipmap) Free() {
	this.Ref().Free()
}

const (
	OES_standard_derivatives_FRAGMENT_SHADER_DERIVATIVE_HINT_OES GLenum = 0x8B8B
)

type OES_standard_derivatives struct {
	ref js.Ref
}

func (this OES_standard_derivatives) Once() OES_standard_derivatives {
	this.Ref().Once()
	return this
}

func (this OES_standard_derivatives) Ref() js.Ref {
	return this.ref
}

func (this OES_standard_derivatives) FromRef(ref js.Ref) OES_standard_derivatives {
	this.ref = ref
	return this
}

func (this OES_standard_derivatives) Free() {
	this.Ref().Free()
}

type OES_texture_float struct {
	ref js.Ref
}

func (this OES_texture_float) Once() OES_texture_float {
	this.Ref().Once()
	return this
}

func (this OES_texture_float) Ref() js.Ref {
	return this.ref
}

func (this OES_texture_float) FromRef(ref js.Ref) OES_texture_float {
	this.ref = ref
	return this
}

func (this OES_texture_float) Free() {
	this.Ref().Free()
}

type OES_texture_float_linear struct {
	ref js.Ref
}

func (this OES_texture_float_linear) Once() OES_texture_float_linear {
	this.Ref().Once()
	return this
}

func (this OES_texture_float_linear) Ref() js.Ref {
	return this.ref
}

func (this OES_texture_float_linear) FromRef(ref js.Ref) OES_texture_float_linear {
	this.ref = ref
	return this
}

func (this OES_texture_float_linear) Free() {
	this.Ref().Free()
}

const (
	OES_texture_half_float_HALF_FLOAT_OES GLenum = 0x8D61
)

type OES_texture_half_float struct {
	ref js.Ref
}

func (this OES_texture_half_float) Once() OES_texture_half_float {
	this.Ref().Once()
	return this
}

func (this OES_texture_half_float) Ref() js.Ref {
	return this.ref
}

func (this OES_texture_half_float) FromRef(ref js.Ref) OES_texture_half_float {
	this.ref = ref
	return this
}

func (this OES_texture_half_float) Free() {
	this.Ref().Free()
}

type OES_texture_half_float_linear struct {
	ref js.Ref
}

func (this OES_texture_half_float_linear) Once() OES_texture_half_float_linear {
	this.Ref().Once()
	return this
}

func (this OES_texture_half_float_linear) Ref() js.Ref {
	return this.ref
}

func (this OES_texture_half_float_linear) FromRef(ref js.Ref) OES_texture_half_float_linear {
	this.ref = ref
	return this
}

func (this OES_texture_half_float_linear) Free() {
	this.Ref().Free()
}

const (
	OES_vertex_array_object_VERTEX_ARRAY_BINDING_OES GLenum = 0x85B5
)

type WebGLVertexArrayObjectOES struct {
	WebGLObject
}

func (this WebGLVertexArrayObjectOES) Once() WebGLVertexArrayObjectOES {
	this.Ref().Once()
	return this
}

func (this WebGLVertexArrayObjectOES) Ref() js.Ref {
	return this.WebGLObject.Ref()
}

func (this WebGLVertexArrayObjectOES) FromRef(ref js.Ref) WebGLVertexArrayObjectOES {
	this.WebGLObject = this.WebGLObject.FromRef(ref)
	return this
}

func (this WebGLVertexArrayObjectOES) Free() {
	this.Ref().Free()
}

type OES_vertex_array_object struct {
	ref js.Ref
}

func (this OES_vertex_array_object) Once() OES_vertex_array_object {
	this.Ref().Once()
	return this
}

func (this OES_vertex_array_object) Ref() js.Ref {
	return this.ref
}

func (this OES_vertex_array_object) FromRef(ref js.Ref) OES_vertex_array_object {
	this.ref = ref
	return this
}

func (this OES_vertex_array_object) Free() {
	this.Ref().Free()
}

// CreateVertexArrayOES calls the method "OES_vertex_array_object.createVertexArrayOES".
//
// The returned bool will be false if there is no such method.
func (this OES_vertex_array_object) CreateVertexArrayOES() (WebGLVertexArrayObjectOES, bool) {
	var _ok bool
	_ret := bindings.CallOES_vertex_array_objectCreateVertexArrayOES(
		this.Ref(), js.Pointer(&_ok),
	)

	return WebGLVertexArrayObjectOES{}.FromRef(_ret), _ok
}

// CreateVertexArrayOESFunc returns the method "OES_vertex_array_object.createVertexArrayOES".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OES_vertex_array_object) CreateVertexArrayOESFunc() (fn js.Func[func() WebGLVertexArrayObjectOES]) {
	return fn.FromRef(
		bindings.OES_vertex_array_objectCreateVertexArrayOESFunc(
			this.Ref(),
		),
	)
}

// DeleteVertexArrayOES calls the method "OES_vertex_array_object.deleteVertexArrayOES".
//
// The returned bool will be false if there is no such method.
func (this OES_vertex_array_object) DeleteVertexArrayOES(arrayObject WebGLVertexArrayObjectOES) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOES_vertex_array_objectDeleteVertexArrayOES(
		this.Ref(), js.Pointer(&_ok),
		arrayObject.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteVertexArrayOESFunc returns the method "OES_vertex_array_object.deleteVertexArrayOES".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OES_vertex_array_object) DeleteVertexArrayOESFunc() (fn js.Func[func(arrayObject WebGLVertexArrayObjectOES)]) {
	return fn.FromRef(
		bindings.OES_vertex_array_objectDeleteVertexArrayOESFunc(
			this.Ref(),
		),
	)
}

// IsVertexArrayOES calls the method "OES_vertex_array_object.isVertexArrayOES".
//
// The returned bool will be false if there is no such method.
func (this OES_vertex_array_object) IsVertexArrayOES(arrayObject WebGLVertexArrayObjectOES) (GLboolean, bool) {
	var _ok bool
	_ret := bindings.CallOES_vertex_array_objectIsVertexArrayOES(
		this.Ref(), js.Pointer(&_ok),
		arrayObject.Ref(),
	)

	return _ret == js.True, _ok
}

// IsVertexArrayOESFunc returns the method "OES_vertex_array_object.isVertexArrayOES".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OES_vertex_array_object) IsVertexArrayOESFunc() (fn js.Func[func(arrayObject WebGLVertexArrayObjectOES) GLboolean]) {
	return fn.FromRef(
		bindings.OES_vertex_array_objectIsVertexArrayOESFunc(
			this.Ref(),
		),
	)
}

// BindVertexArrayOES calls the method "OES_vertex_array_object.bindVertexArrayOES".
//
// The returned bool will be false if there is no such method.
func (this OES_vertex_array_object) BindVertexArrayOES(arrayObject WebGLVertexArrayObjectOES) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOES_vertex_array_objectBindVertexArrayOES(
		this.Ref(), js.Pointer(&_ok),
		arrayObject.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BindVertexArrayOESFunc returns the method "OES_vertex_array_object.bindVertexArrayOES".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OES_vertex_array_object) BindVertexArrayOESFunc() (fn js.Func[func(arrayObject WebGLVertexArrayObjectOES)]) {
	return fn.FromRef(
		bindings.OES_vertex_array_objectBindVertexArrayOESFunc(
			this.Ref(),
		),
	)
}

type OTPCredential struct {
	Credential
}

func (this OTPCredential) Once() OTPCredential {
	this.Ref().Once()
	return this
}

func (this OTPCredential) Ref() js.Ref {
	return this.Credential.Ref()
}

func (this OTPCredential) FromRef(ref js.Ref) OTPCredential {
	this.Credential = this.Credential.FromRef(ref)
	return this
}

func (this OTPCredential) Free() {
	this.Ref().Free()
}

// Code returns the value of property "OTPCredential.code".
//
// The returned bool will be false if there is no such property.
func (this OTPCredential) Code() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetOTPCredentialCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

const (
	OVR_multiview2_FRAMEBUFFER_ATTACHMENT_TEXTURE_NUM_VIEWS_OVR       GLenum = 0x9630
	OVR_multiview2_FRAMEBUFFER_ATTACHMENT_TEXTURE_BASE_VIEW_INDEX_OVR GLenum = 0x9632
	OVR_multiview2_MAX_VIEWS_OVR                                      GLenum = 0x9631
	OVR_multiview2_FRAMEBUFFER_INCOMPLETE_VIEW_TARGETS_OVR            GLenum = 0x9633
)

type OVR_multiview2 struct {
	ref js.Ref
}

func (this OVR_multiview2) Once() OVR_multiview2 {
	this.Ref().Once()
	return this
}

func (this OVR_multiview2) Ref() js.Ref {
	return this.ref
}

func (this OVR_multiview2) FromRef(ref js.Ref) OVR_multiview2 {
	this.ref = ref
	return this
}

func (this OVR_multiview2) Free() {
	this.Ref().Free()
}

// FramebufferTextureMultiviewOVR calls the method "OVR_multiview2.framebufferTextureMultiviewOVR".
//
// The returned bool will be false if there is no such method.
func (this OVR_multiview2) FramebufferTextureMultiviewOVR(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, baseViewIndex GLint, numViews GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOVR_multiview2FramebufferTextureMultiviewOVR(
		this.Ref(), js.Pointer(&_ok),
		uint32(target),
		uint32(attachment),
		texture.Ref(),
		int32(level),
		int32(baseViewIndex),
		int32(numViews),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FramebufferTextureMultiviewOVRFunc returns the method "OVR_multiview2.framebufferTextureMultiviewOVR".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OVR_multiview2) FramebufferTextureMultiviewOVRFunc() (fn js.Func[func(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, baseViewIndex GLint, numViews GLsizei)]) {
	return fn.FromRef(
		bindings.OVR_multiview2FramebufferTextureMultiviewOVRFunc(
			this.Ref(),
		),
	)
}

type OfflineAudioCompletionEventInit struct {
	// RenderedBuffer is "OfflineAudioCompletionEventInit.renderedBuffer"
	//
	// Required
	RenderedBuffer AudioBuffer
	// Bubbles is "OfflineAudioCompletionEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "OfflineAudioCompletionEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "OfflineAudioCompletionEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OfflineAudioCompletionEventInit with all fields set.
func (p OfflineAudioCompletionEventInit) FromRef(ref js.Ref) OfflineAudioCompletionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 OfflineAudioCompletionEventInit OfflineAudioCompletionEventInit [// OfflineAudioCompletionEventInit] [0x14000935b80 0x14000935c20 0x14000935d60 0x14000935ea0 0x14000935cc0 0x14000935e00 0x14000935f40] 0x140009205d0 {0 0}} in the application heap.
func (p OfflineAudioCompletionEventInit) New() js.Ref {
	return bindings.OfflineAudioCompletionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p OfflineAudioCompletionEventInit) UpdateFrom(ref js.Ref) {
	bindings.OfflineAudioCompletionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p OfflineAudioCompletionEventInit) Update(ref js.Ref) {
	bindings.OfflineAudioCompletionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewOfflineAudioCompletionEvent(typ js.String, eventInitDict OfflineAudioCompletionEventInit) OfflineAudioCompletionEvent {
	return OfflineAudioCompletionEvent{}.FromRef(
		bindings.NewOfflineAudioCompletionEventByOfflineAudioCompletionEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type OfflineAudioCompletionEvent struct {
	Event
}

func (this OfflineAudioCompletionEvent) Once() OfflineAudioCompletionEvent {
	this.Ref().Once()
	return this
}

func (this OfflineAudioCompletionEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this OfflineAudioCompletionEvent) FromRef(ref js.Ref) OfflineAudioCompletionEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this OfflineAudioCompletionEvent) Free() {
	this.Ref().Free()
}

// RenderedBuffer returns the value of property "OfflineAudioCompletionEvent.renderedBuffer".
//
// The returned bool will be false if there is no such property.
func (this OfflineAudioCompletionEvent) RenderedBuffer() (AudioBuffer, bool) {
	var _ok bool
	_ret := bindings.GetOfflineAudioCompletionEventRenderedBuffer(
		this.Ref(), js.Pointer(&_ok),
	)
	return AudioBuffer{}.FromRef(_ret), _ok
}

type OfflineAudioContextOptions struct {
	// NumberOfChannels is "OfflineAudioContextOptions.numberOfChannels"
	//
	// Optional, defaults to 1.
	NumberOfChannels uint32
	// Length is "OfflineAudioContextOptions.length"
	//
	// Required
	Length uint32
	// SampleRate is "OfflineAudioContextOptions.sampleRate"
	//
	// Required
	SampleRate float32

	FFI_USE_NumberOfChannels bool // for NumberOfChannels.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OfflineAudioContextOptions with all fields set.
func (p OfflineAudioContextOptions) FromRef(ref js.Ref) OfflineAudioContextOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 OfflineAudioContextOptions OfflineAudioContextOptions [// OfflineAudioContextOptions] [0x1400093e000 0x1400093e140 0x1400093e1e0 0x1400093e0a0] 0x14000920600 {0 0}} in the application heap.
func (p OfflineAudioContextOptions) New() js.Ref {
	return bindings.OfflineAudioContextOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p OfflineAudioContextOptions) UpdateFrom(ref js.Ref) {
	bindings.OfflineAudioContextOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p OfflineAudioContextOptions) Update(ref js.Ref) {
	bindings.OfflineAudioContextOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewOfflineAudioContext(contextOptions OfflineAudioContextOptions) OfflineAudioContext {
	return OfflineAudioContext{}.FromRef(
		bindings.NewOfflineAudioContextByOfflineAudioContext(
			js.Pointer(&contextOptions)),
	)
}

func NewOfflineAudioContextByOfflineAudioContext1(numberOfChannels uint32, length uint32, sampleRate float32) OfflineAudioContext {
	return OfflineAudioContext{}.FromRef(
		bindings.NewOfflineAudioContextByOfflineAudioContext1(
			uint32(numberOfChannels),
			uint32(length),
			float32(sampleRate)),
	)
}

type OfflineAudioContext struct {
	BaseAudioContext
}

func (this OfflineAudioContext) Once() OfflineAudioContext {
	this.Ref().Once()
	return this
}

func (this OfflineAudioContext) Ref() js.Ref {
	return this.BaseAudioContext.Ref()
}

func (this OfflineAudioContext) FromRef(ref js.Ref) OfflineAudioContext {
	this.BaseAudioContext = this.BaseAudioContext.FromRef(ref)
	return this
}

func (this OfflineAudioContext) Free() {
	this.Ref().Free()
}

// Length returns the value of property "OfflineAudioContext.length".
//
// The returned bool will be false if there is no such property.
func (this OfflineAudioContext) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetOfflineAudioContextLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// StartRendering calls the method "OfflineAudioContext.startRendering".
//
// The returned bool will be false if there is no such method.
func (this OfflineAudioContext) StartRendering() (js.Promise[AudioBuffer], bool) {
	var _ok bool
	_ret := bindings.CallOfflineAudioContextStartRendering(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[AudioBuffer]{}.FromRef(_ret), _ok
}

// StartRenderingFunc returns the method "OfflineAudioContext.startRendering".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OfflineAudioContext) StartRenderingFunc() (fn js.Func[func() js.Promise[AudioBuffer]]) {
	return fn.FromRef(
		bindings.OfflineAudioContextStartRenderingFunc(
			this.Ref(),
		),
	)
}

// Resume calls the method "OfflineAudioContext.resume".
//
// The returned bool will be false if there is no such method.
func (this OfflineAudioContext) Resume() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallOfflineAudioContextResume(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ResumeFunc returns the method "OfflineAudioContext.resume".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OfflineAudioContext) ResumeFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.OfflineAudioContextResumeFunc(
			this.Ref(),
		),
	)
}

// Suspend calls the method "OfflineAudioContext.suspend".
//
// The returned bool will be false if there is no such method.
func (this OfflineAudioContext) Suspend(suspendTime float64) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallOfflineAudioContextSuspend(
		this.Ref(), js.Pointer(&_ok),
		float64(suspendTime),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// SuspendFunc returns the method "OfflineAudioContext.suspend".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OfflineAudioContext) SuspendFunc() (fn js.Func[func(suspendTime float64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.OfflineAudioContextSuspendFunc(
			this.Ref(),
		),
	)
}

type OneOf_TypedArrayFloat32_TypedArrayFloat64_DOMMatrix struct {
	ref js.Ref
}

func (x OneOf_TypedArrayFloat32_TypedArrayFloat64_DOMMatrix) Ref() js.Ref {
	return x.ref
}

func (x OneOf_TypedArrayFloat32_TypedArrayFloat64_DOMMatrix) Free() {
	x.ref.Free()
}

func (x OneOf_TypedArrayFloat32_TypedArrayFloat64_DOMMatrix) FromRef(ref js.Ref) OneOf_TypedArrayFloat32_TypedArrayFloat64_DOMMatrix {
	return OneOf_TypedArrayFloat32_TypedArrayFloat64_DOMMatrix{
		ref: ref,
	}
}

func (x OneOf_TypedArrayFloat32_TypedArrayFloat64_DOMMatrix) TypedArrayFloat32() js.TypedArray[float32] {
	return js.TypedArray[float32]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayFloat32_TypedArrayFloat64_DOMMatrix) TypedArrayFloat64() js.TypedArray[float64] {
	return js.TypedArray[float64]{}.FromRef(x.ref)
}

func (x OneOf_TypedArrayFloat32_TypedArrayFloat64_DOMMatrix) DOMMatrix() DOMMatrix {
	return DOMMatrix{}.FromRef(x.ref)
}

type RotationMatrixType = OneOf_TypedArrayFloat32_TypedArrayFloat64_DOMMatrix

type OrientationSensor struct {
	Sensor
}

func (this OrientationSensor) Once() OrientationSensor {
	this.Ref().Once()
	return this
}

func (this OrientationSensor) Ref() js.Ref {
	return this.Sensor.Ref()
}

func (this OrientationSensor) FromRef(ref js.Ref) OrientationSensor {
	this.Sensor = this.Sensor.FromRef(ref)
	return this
}

func (this OrientationSensor) Free() {
	this.Ref().Free()
}

// Quaternion returns the value of property "OrientationSensor.quaternion".
//
// The returned bool will be false if there is no such property.
func (this OrientationSensor) Quaternion() (js.FrozenArray[float64], bool) {
	var _ok bool
	_ret := bindings.GetOrientationSensorQuaternion(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[float64]{}.FromRef(_ret), _ok
}

// PopulateMatrix calls the method "OrientationSensor.populateMatrix".
//
// The returned bool will be false if there is no such method.
func (this OrientationSensor) PopulateMatrix(targetMatrix RotationMatrixType) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallOrientationSensorPopulateMatrix(
		this.Ref(), js.Pointer(&_ok),
		targetMatrix.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PopulateMatrixFunc returns the method "OrientationSensor.populateMatrix".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this OrientationSensor) PopulateMatrixFunc() (fn js.Func[func(targetMatrix RotationMatrixType)]) {
	return fn.FromRef(
		bindings.OrientationSensorPopulateMatrixFunc(
			this.Ref(),
		),
	)
}

func NewOverconstrainedError(constraint js.String, message js.String) OverconstrainedError {
	return OverconstrainedError{}.FromRef(
		bindings.NewOverconstrainedErrorByOverconstrainedError(
			constraint.Ref(),
			message.Ref()),
	)
}

func NewOverconstrainedErrorByOverconstrainedError1(constraint js.String) OverconstrainedError {
	return OverconstrainedError{}.FromRef(
		bindings.NewOverconstrainedErrorByOverconstrainedError1(
			constraint.Ref()),
	)
}

type OverconstrainedError struct {
	DOMException
}

func (this OverconstrainedError) Once() OverconstrainedError {
	this.Ref().Once()
	return this
}

func (this OverconstrainedError) Ref() js.Ref {
	return this.DOMException.Ref()
}

func (this OverconstrainedError) FromRef(ref js.Ref) OverconstrainedError {
	this.DOMException = this.DOMException.FromRef(ref)
	return this
}

func (this OverconstrainedError) Free() {
	this.Ref().Free()
}

// Constraint returns the value of property "OverconstrainedError.constraint".
//
// The returned bool will be false if there is no such property.
func (this OverconstrainedError) Constraint() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetOverconstrainedErrorConstraint(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type PageTransitionEventInit struct {
	// Persisted is "PageTransitionEventInit.persisted"
	//
	// Optional, defaults to false.
	Persisted bool
	// Bubbles is "PageTransitionEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "PageTransitionEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "PageTransitionEventInit.composed"
	//
	// Optional, defaults to false.
	Composed bool

	FFI_USE_Persisted  bool // for Persisted.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PageTransitionEventInit with all fields set.
func (p PageTransitionEventInit) FromRef(ref js.Ref) PageTransitionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PageTransitionEventInit PageTransitionEventInit [// PageTransitionEventInit] [0x1400093e280 0x1400093e3c0 0x1400093e500 0x1400093e640 0x1400093e320 0x1400093e460 0x1400093e5a0 0x1400093e6e0] 0x14000920720 {0 0}} in the application heap.
func (p PageTransitionEventInit) New() js.Ref {
	return bindings.PageTransitionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PageTransitionEventInit) UpdateFrom(ref js.Ref) {
	bindings.PageTransitionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PageTransitionEventInit) Update(ref js.Ref) {
	bindings.PageTransitionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewPageTransitionEvent(typ js.String, eventInitDict PageTransitionEventInit) PageTransitionEvent {
	return PageTransitionEvent{}.FromRef(
		bindings.NewPageTransitionEventByPageTransitionEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewPageTransitionEventByPageTransitionEvent1(typ js.String) PageTransitionEvent {
	return PageTransitionEvent{}.FromRef(
		bindings.NewPageTransitionEventByPageTransitionEvent1(
			typ.Ref()),
	)
}

type PageTransitionEvent struct {
	Event
}

func (this PageTransitionEvent) Once() PageTransitionEvent {
	this.Ref().Once()
	return this
}

func (this PageTransitionEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this PageTransitionEvent) FromRef(ref js.Ref) PageTransitionEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this PageTransitionEvent) Free() {
	this.Ref().Free()
}

// Persisted returns the value of property "PageTransitionEvent.persisted".
//
// The returned bool will be false if there is no such property.
func (this PageTransitionEvent) Persisted() (bool, bool) {
	var _ok bool
	_ret := bindings.GetPageTransitionEventPersisted(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type PaintRenderingContext2D struct {
	ref js.Ref
}

func (this PaintRenderingContext2D) Once() PaintRenderingContext2D {
	this.Ref().Once()
	return this
}

func (this PaintRenderingContext2D) Ref() js.Ref {
	return this.ref
}

func (this PaintRenderingContext2D) FromRef(ref js.Ref) PaintRenderingContext2D {
	this.ref = ref
	return this
}

func (this PaintRenderingContext2D) Free() {
	this.Ref().Free()
}

// LineWidth returns the value of property "PaintRenderingContext2D.lineWidth".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) LineWidth() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DLineWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// LineWidth sets the value of property "PaintRenderingContext2D.lineWidth" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetLineWidth(val float64) bool {
	return js.True == bindings.SetPaintRenderingContext2DLineWidth(
		this.Ref(),
		float64(val),
	)
}

// LineCap returns the value of property "PaintRenderingContext2D.lineCap".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) LineCap() (CanvasLineCap, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DLineCap(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasLineCap(_ret), _ok
}

// LineCap sets the value of property "PaintRenderingContext2D.lineCap" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetLineCap(val CanvasLineCap) bool {
	return js.True == bindings.SetPaintRenderingContext2DLineCap(
		this.Ref(),
		uint32(val),
	)
}

// LineJoin returns the value of property "PaintRenderingContext2D.lineJoin".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) LineJoin() (CanvasLineJoin, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DLineJoin(
		this.Ref(), js.Pointer(&_ok),
	)
	return CanvasLineJoin(_ret), _ok
}

// LineJoin sets the value of property "PaintRenderingContext2D.lineJoin" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetLineJoin(val CanvasLineJoin) bool {
	return js.True == bindings.SetPaintRenderingContext2DLineJoin(
		this.Ref(),
		uint32(val),
	)
}

// MiterLimit returns the value of property "PaintRenderingContext2D.miterLimit".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) MiterLimit() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DMiterLimit(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// MiterLimit sets the value of property "PaintRenderingContext2D.miterLimit" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetMiterLimit(val float64) bool {
	return js.True == bindings.SetPaintRenderingContext2DMiterLimit(
		this.Ref(),
		float64(val),
	)
}

// LineDashOffset returns the value of property "PaintRenderingContext2D.lineDashOffset".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) LineDashOffset() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DLineDashOffset(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// LineDashOffset sets the value of property "PaintRenderingContext2D.lineDashOffset" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetLineDashOffset(val float64) bool {
	return js.True == bindings.SetPaintRenderingContext2DLineDashOffset(
		this.Ref(),
		float64(val),
	)
}

// ShadowOffsetX returns the value of property "PaintRenderingContext2D.shadowOffsetX".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) ShadowOffsetX() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DShadowOffsetX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ShadowOffsetX sets the value of property "PaintRenderingContext2D.shadowOffsetX" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetShadowOffsetX(val float64) bool {
	return js.True == bindings.SetPaintRenderingContext2DShadowOffsetX(
		this.Ref(),
		float64(val),
	)
}

// ShadowOffsetY returns the value of property "PaintRenderingContext2D.shadowOffsetY".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) ShadowOffsetY() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DShadowOffsetY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ShadowOffsetY sets the value of property "PaintRenderingContext2D.shadowOffsetY" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetShadowOffsetY(val float64) bool {
	return js.True == bindings.SetPaintRenderingContext2DShadowOffsetY(
		this.Ref(),
		float64(val),
	)
}

// ShadowBlur returns the value of property "PaintRenderingContext2D.shadowBlur".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) ShadowBlur() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DShadowBlur(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ShadowBlur sets the value of property "PaintRenderingContext2D.shadowBlur" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetShadowBlur(val float64) bool {
	return js.True == bindings.SetPaintRenderingContext2DShadowBlur(
		this.Ref(),
		float64(val),
	)
}

// ShadowColor returns the value of property "PaintRenderingContext2D.shadowColor".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) ShadowColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DShadowColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ShadowColor sets the value of property "PaintRenderingContext2D.shadowColor" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetShadowColor(val js.String) bool {
	return js.True == bindings.SetPaintRenderingContext2DShadowColor(
		this.Ref(),
		val.Ref(),
	)
}

// StrokeStyle returns the value of property "PaintRenderingContext2D.strokeStyle".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) StrokeStyle() (OneOf_String_CanvasGradient_CanvasPattern, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DStrokeStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_String_CanvasGradient_CanvasPattern{}.FromRef(_ret), _ok
}

// StrokeStyle sets the value of property "PaintRenderingContext2D.strokeStyle" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetStrokeStyle(val OneOf_String_CanvasGradient_CanvasPattern) bool {
	return js.True == bindings.SetPaintRenderingContext2DStrokeStyle(
		this.Ref(),
		val.Ref(),
	)
}

// FillStyle returns the value of property "PaintRenderingContext2D.fillStyle".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) FillStyle() (OneOf_String_CanvasGradient_CanvasPattern, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DFillStyle(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_String_CanvasGradient_CanvasPattern{}.FromRef(_ret), _ok
}

// FillStyle sets the value of property "PaintRenderingContext2D.fillStyle" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetFillStyle(val OneOf_String_CanvasGradient_CanvasPattern) bool {
	return js.True == bindings.SetPaintRenderingContext2DFillStyle(
		this.Ref(),
		val.Ref(),
	)
}

// ImageSmoothingEnabled returns the value of property "PaintRenderingContext2D.imageSmoothingEnabled".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) ImageSmoothingEnabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DImageSmoothingEnabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ImageSmoothingEnabled sets the value of property "PaintRenderingContext2D.imageSmoothingEnabled" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetImageSmoothingEnabled(val bool) bool {
	return js.True == bindings.SetPaintRenderingContext2DImageSmoothingEnabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// ImageSmoothingQuality returns the value of property "PaintRenderingContext2D.imageSmoothingQuality".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) ImageSmoothingQuality() (ImageSmoothingQuality, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DImageSmoothingQuality(
		this.Ref(), js.Pointer(&_ok),
	)
	return ImageSmoothingQuality(_ret), _ok
}

// ImageSmoothingQuality sets the value of property "PaintRenderingContext2D.imageSmoothingQuality" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetImageSmoothingQuality(val ImageSmoothingQuality) bool {
	return js.True == bindings.SetPaintRenderingContext2DImageSmoothingQuality(
		this.Ref(),
		uint32(val),
	)
}

// GlobalAlpha returns the value of property "PaintRenderingContext2D.globalAlpha".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) GlobalAlpha() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DGlobalAlpha(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// GlobalAlpha sets the value of property "PaintRenderingContext2D.globalAlpha" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetGlobalAlpha(val float64) bool {
	return js.True == bindings.SetPaintRenderingContext2DGlobalAlpha(
		this.Ref(),
		float64(val),
	)
}

// GlobalCompositeOperation returns the value of property "PaintRenderingContext2D.globalCompositeOperation".
//
// The returned bool will be false if there is no such property.
func (this PaintRenderingContext2D) GlobalCompositeOperation() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPaintRenderingContext2DGlobalCompositeOperation(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// GlobalCompositeOperation sets the value of property "PaintRenderingContext2D.globalCompositeOperation" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetGlobalCompositeOperation(val js.String) bool {
	return js.True == bindings.SetPaintRenderingContext2DGlobalCompositeOperation(
		this.Ref(),
		val.Ref(),
	)
}

// ClosePath calls the method "PaintRenderingContext2D.closePath".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) ClosePath() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DClosePath(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClosePathFunc returns the method "PaintRenderingContext2D.closePath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) ClosePathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClosePathFunc(
			this.Ref(),
		),
	)
}

// MoveTo calls the method "PaintRenderingContext2D.moveTo".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) MoveTo(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DMoveTo(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// MoveToFunc returns the method "PaintRenderingContext2D.moveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) MoveToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DMoveToFunc(
			this.Ref(),
		),
	)
}

// LineTo calls the method "PaintRenderingContext2D.lineTo".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) LineTo(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DLineTo(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// LineToFunc returns the method "PaintRenderingContext2D.lineTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) LineToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DLineToFunc(
			this.Ref(),
		),
	)
}

// QuadraticCurveTo calls the method "PaintRenderingContext2D.quadraticCurveTo".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DQuadraticCurveTo(
		this.Ref(), js.Pointer(&_ok),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// QuadraticCurveToFunc returns the method "PaintRenderingContext2D.quadraticCurveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) QuadraticCurveToFunc() (fn js.Func[func(cpx float64, cpy float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DQuadraticCurveToFunc(
			this.Ref(),
		),
	)
}

// BezierCurveTo calls the method "PaintRenderingContext2D.bezierCurveTo".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DBezierCurveTo(
		this.Ref(), js.Pointer(&_ok),
		float64(cp1x),
		float64(cp1y),
		float64(cp2x),
		float64(cp2y),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BezierCurveToFunc returns the method "PaintRenderingContext2D.bezierCurveTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) BezierCurveToFunc() (fn js.Func[func(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DBezierCurveToFunc(
			this.Ref(),
		),
	)
}

// ArcTo calls the method "PaintRenderingContext2D.arcTo".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DArcTo(
		this.Ref(), js.Pointer(&_ok),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ArcToFunc returns the method "PaintRenderingContext2D.arcTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) ArcToFunc() (fn js.Func[func(x1 float64, y1 float64, x2 float64, y2 float64, radius float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DArcToFunc(
			this.Ref(),
		),
	)
}

// Rect calls the method "PaintRenderingContext2D.rect".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Rect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RectFunc returns the method "PaintRenderingContext2D.rect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) RectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DRectFunc(
			this.Ref(),
		),
	)
}

// RoundRect calls the method "PaintRenderingContext2D.roundRect".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) RoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DRoundRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RoundRectFunc returns the method "PaintRenderingContext2D.roundRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) RoundRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DRoundRectFunc(
			this.Ref(),
		),
	)
}

// RoundRect1 calls the method "PaintRenderingContext2D.roundRect".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) RoundRect1(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DRoundRect1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RoundRect1Func returns the method "PaintRenderingContext2D.roundRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) RoundRect1Func() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DRoundRect1Func(
			this.Ref(),
		),
	)
}

// Arc calls the method "PaintRenderingContext2D.arc".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DArc(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ArcFunc returns the method "PaintRenderingContext2D.arc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) ArcFunc() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DArcFunc(
			this.Ref(),
		),
	)
}

// Arc1 calls the method "PaintRenderingContext2D.arc".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Arc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DArc1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Arc1Func returns the method "PaintRenderingContext2D.arc".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) Arc1Func() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DArc1Func(
			this.Ref(),
		),
	)
}

// Ellipse calls the method "PaintRenderingContext2D.ellipse".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Ellipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DEllipse(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// EllipseFunc returns the method "PaintRenderingContext2D.ellipse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) EllipseFunc() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DEllipseFunc(
			this.Ref(),
		),
	)
}

// Ellipse1 calls the method "PaintRenderingContext2D.ellipse".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Ellipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DEllipse1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Ellipse1Func returns the method "PaintRenderingContext2D.ellipse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) Ellipse1Func() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DEllipse1Func(
			this.Ref(),
		),
	)
}

// SetLineDash calls the method "PaintRenderingContext2D.setLineDash".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) SetLineDash(segments js.Array[float64]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DSetLineDash(
		this.Ref(), js.Pointer(&_ok),
		segments.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetLineDashFunc returns the method "PaintRenderingContext2D.setLineDash".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) SetLineDashFunc() (fn js.Func[func(segments js.Array[float64])]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DSetLineDashFunc(
			this.Ref(),
		),
	)
}

// GetLineDash calls the method "PaintRenderingContext2D.getLineDash".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) GetLineDash() (js.Array[float64], bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DGetLineDash(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[float64]{}.FromRef(_ret), _ok
}

// GetLineDashFunc returns the method "PaintRenderingContext2D.getLineDash".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) GetLineDashFunc() (fn js.Func[func() js.Array[float64]]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DGetLineDashFunc(
			this.Ref(),
		),
	)
}

// DrawImage calls the method "PaintRenderingContext2D.drawImage".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) DrawImage(image CanvasImageSource, dx float64, dy float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DDrawImage(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		float64(dx),
		float64(dy),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawImageFunc returns the method "PaintRenderingContext2D.drawImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) DrawImageFunc() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DDrawImageFunc(
			this.Ref(),
		),
	)
}

// DrawImage1 calls the method "PaintRenderingContext2D.drawImage".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) DrawImage1(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DDrawImage1(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		float64(dx),
		float64(dy),
		float64(dw),
		float64(dh),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawImage1Func returns the method "PaintRenderingContext2D.drawImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) DrawImage1Func() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DDrawImage1Func(
			this.Ref(),
		),
	)
}

// DrawImage2 calls the method "PaintRenderingContext2D.drawImage".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) DrawImage2(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DDrawImage2(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		float64(sx),
		float64(sy),
		float64(sw),
		float64(sh),
		float64(dx),
		float64(dy),
		float64(dw),
		float64(dh),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawImage2Func returns the method "PaintRenderingContext2D.drawImage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) DrawImage2Func() (fn js.Func[func(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DDrawImage2Func(
			this.Ref(),
		),
	)
}

// BeginPath calls the method "PaintRenderingContext2D.beginPath".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) BeginPath() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DBeginPath(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// BeginPathFunc returns the method "PaintRenderingContext2D.beginPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) BeginPathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DBeginPathFunc(
			this.Ref(),
		),
	)
}

// Fill calls the method "PaintRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Fill(fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DFill(
		this.Ref(), js.Pointer(&_ok),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FillFunc returns the method "PaintRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) FillFunc() (fn js.Func[func(fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DFillFunc(
			this.Ref(),
		),
	)
}

// Fill1 calls the method "PaintRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Fill1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DFill1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Fill1Func returns the method "PaintRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) Fill1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DFill1Func(
			this.Ref(),
		),
	)
}

// Fill2 calls the method "PaintRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Fill2(path Path2D, fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DFill2(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Fill2Func returns the method "PaintRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) Fill2Func() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DFill2Func(
			this.Ref(),
		),
	)
}

// Fill3 calls the method "PaintRenderingContext2D.fill".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Fill3(path Path2D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DFill3(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Fill3Func returns the method "PaintRenderingContext2D.fill".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) Fill3Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DFill3Func(
			this.Ref(),
		),
	)
}

// Stroke calls the method "PaintRenderingContext2D.stroke".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Stroke() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DStroke(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StrokeFunc returns the method "PaintRenderingContext2D.stroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) StrokeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DStrokeFunc(
			this.Ref(),
		),
	)
}

// Stroke1 calls the method "PaintRenderingContext2D.stroke".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Stroke1(path Path2D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DStroke1(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Stroke1Func returns the method "PaintRenderingContext2D.stroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) Stroke1Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DStroke1Func(
			this.Ref(),
		),
	)
}

// Clip calls the method "PaintRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Clip(fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DClip(
		this.Ref(), js.Pointer(&_ok),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClipFunc returns the method "PaintRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) ClipFunc() (fn js.Func[func(fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClipFunc(
			this.Ref(),
		),
	)
}

// Clip1 calls the method "PaintRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Clip1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DClip1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Clip1Func returns the method "PaintRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) Clip1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClip1Func(
			this.Ref(),
		),
	)
}

// Clip2 calls the method "PaintRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Clip2(path Path2D, fillRule CanvasFillRule) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DClip2(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		uint32(fillRule),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Clip2Func returns the method "PaintRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) Clip2Func() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClip2Func(
			this.Ref(),
		),
	)
}

// Clip3 calls the method "PaintRenderingContext2D.clip".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Clip3(path Path2D) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DClip3(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Clip3Func returns the method "PaintRenderingContext2D.clip".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) Clip3Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClip3Func(
			this.Ref(),
		),
	)
}

// IsPointInPath calls the method "PaintRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) IsPointInPath(x float64, y float64, fillRule CanvasFillRule) (bool, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DIsPointInPath(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return _ret == js.True, _ok
}

// IsPointInPathFunc returns the method "PaintRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) IsPointInPathFunc() (fn js.Func[func(x float64, y float64, fillRule CanvasFillRule) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInPathFunc(
			this.Ref(),
		),
	)
}

// IsPointInPath1 calls the method "PaintRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) IsPointInPath1(x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DIsPointInPath1(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInPath1Func returns the method "PaintRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) IsPointInPath1Func() (fn js.Func[func(x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInPath1Func(
			this.Ref(),
		),
	)
}

// IsPointInPath2 calls the method "PaintRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) IsPointInPath2(path Path2D, x float64, y float64, fillRule CanvasFillRule) (bool, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DIsPointInPath2(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return _ret == js.True, _ok
}

// IsPointInPath2Func returns the method "PaintRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) IsPointInPath2Func() (fn js.Func[func(path Path2D, x float64, y float64, fillRule CanvasFillRule) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInPath2Func(
			this.Ref(),
		),
	)
}

// IsPointInPath3 calls the method "PaintRenderingContext2D.isPointInPath".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) IsPointInPath3(path Path2D, x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DIsPointInPath3(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInPath3Func returns the method "PaintRenderingContext2D.isPointInPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) IsPointInPath3Func() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInPath3Func(
			this.Ref(),
		),
	)
}

// IsPointInStroke calls the method "PaintRenderingContext2D.isPointInStroke".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) IsPointInStroke(x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DIsPointInStroke(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInStrokeFunc returns the method "PaintRenderingContext2D.isPointInStroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) IsPointInStrokeFunc() (fn js.Func[func(x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInStrokeFunc(
			this.Ref(),
		),
	)
}

// IsPointInStroke1 calls the method "PaintRenderingContext2D.isPointInStroke".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) IsPointInStroke1(path Path2D, x float64, y float64) (bool, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DIsPointInStroke1(
		this.Ref(), js.Pointer(&_ok),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return _ret == js.True, _ok
}

// IsPointInStroke1Func returns the method "PaintRenderingContext2D.isPointInStroke".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) IsPointInStroke1Func() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInStroke1Func(
			this.Ref(),
		),
	)
}

// ClearRect calls the method "PaintRenderingContext2D.clearRect".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) ClearRect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DClearRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ClearRectFunc returns the method "PaintRenderingContext2D.clearRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) ClearRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClearRectFunc(
			this.Ref(),
		),
	)
}

// FillRect calls the method "PaintRenderingContext2D.fillRect".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) FillRect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DFillRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FillRectFunc returns the method "PaintRenderingContext2D.fillRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) FillRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DFillRectFunc(
			this.Ref(),
		),
	)
}

// StrokeRect calls the method "PaintRenderingContext2D.strokeRect".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) StrokeRect(x float64, y float64, w float64, h float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DStrokeRect(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StrokeRectFunc returns the method "PaintRenderingContext2D.strokeRect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) StrokeRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DStrokeRectFunc(
			this.Ref(),
		),
	)
}

// CreateLinearGradient calls the method "PaintRenderingContext2D.createLinearGradient".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) CreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) (CanvasGradient, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DCreateLinearGradient(
		this.Ref(), js.Pointer(&_ok),
		float64(x0),
		float64(y0),
		float64(x1),
		float64(y1),
	)

	return CanvasGradient{}.FromRef(_ret), _ok
}

// CreateLinearGradientFunc returns the method "PaintRenderingContext2D.createLinearGradient".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) CreateLinearGradientFunc() (fn js.Func[func(x0 float64, y0 float64, x1 float64, y1 float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DCreateLinearGradientFunc(
			this.Ref(),
		),
	)
}

// CreateRadialGradient calls the method "PaintRenderingContext2D.createRadialGradient".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) CreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) (CanvasGradient, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DCreateRadialGradient(
		this.Ref(), js.Pointer(&_ok),
		float64(x0),
		float64(y0),
		float64(r0),
		float64(x1),
		float64(y1),
		float64(r1),
	)

	return CanvasGradient{}.FromRef(_ret), _ok
}

// CreateRadialGradientFunc returns the method "PaintRenderingContext2D.createRadialGradient".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) CreateRadialGradientFunc() (fn js.Func[func(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DCreateRadialGradientFunc(
			this.Ref(),
		),
	)
}

// CreateConicGradient calls the method "PaintRenderingContext2D.createConicGradient".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) CreateConicGradient(startAngle float64, x float64, y float64) (CanvasGradient, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DCreateConicGradient(
		this.Ref(), js.Pointer(&_ok),
		float64(startAngle),
		float64(x),
		float64(y),
	)

	return CanvasGradient{}.FromRef(_ret), _ok
}

// CreateConicGradientFunc returns the method "PaintRenderingContext2D.createConicGradient".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) CreateConicGradientFunc() (fn js.Func[func(startAngle float64, x float64, y float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DCreateConicGradientFunc(
			this.Ref(),
		),
	)
}

// CreatePattern calls the method "PaintRenderingContext2D.createPattern".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) CreatePattern(image CanvasImageSource, repetition js.String) (CanvasPattern, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DCreatePattern(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
		repetition.Ref(),
	)

	return CanvasPattern{}.FromRef(_ret), _ok
}

// CreatePatternFunc returns the method "PaintRenderingContext2D.createPattern".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) CreatePatternFunc() (fn js.Func[func(image CanvasImageSource, repetition js.String) CanvasPattern]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DCreatePatternFunc(
			this.Ref(),
		),
	)
}

// Scale calls the method "PaintRenderingContext2D.scale".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Scale(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DScale(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ScaleFunc returns the method "PaintRenderingContext2D.scale".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) ScaleFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DScaleFunc(
			this.Ref(),
		),
	)
}

// Rotate calls the method "PaintRenderingContext2D.rotate".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Rotate(angle float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DRotate(
		this.Ref(), js.Pointer(&_ok),
		float64(angle),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RotateFunc returns the method "PaintRenderingContext2D.rotate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) RotateFunc() (fn js.Func[func(angle float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DRotateFunc(
			this.Ref(),
		),
	)
}

// Translate calls the method "PaintRenderingContext2D.translate".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Translate(x float64, y float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DTranslate(
		this.Ref(), js.Pointer(&_ok),
		float64(x),
		float64(y),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TranslateFunc returns the method "PaintRenderingContext2D.translate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) TranslateFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DTranslateFunc(
			this.Ref(),
		),
	)
}

// Transform calls the method "PaintRenderingContext2D.transform".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Transform(a float64, b float64, c float64, d float64, e float64, f float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DTransform(
		this.Ref(), js.Pointer(&_ok),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	_ = _ret
	return js.Void{}, _ok
}

// TransformFunc returns the method "PaintRenderingContext2D.transform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) TransformFunc() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DTransformFunc(
			this.Ref(),
		),
	)
}

// GetTransform calls the method "PaintRenderingContext2D.getTransform".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) GetTransform() (DOMMatrix, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DGetTransform(
		this.Ref(), js.Pointer(&_ok),
	)

	return DOMMatrix{}.FromRef(_ret), _ok
}

// GetTransformFunc returns the method "PaintRenderingContext2D.getTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) GetTransformFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DGetTransformFunc(
			this.Ref(),
		),
	)
}

// SetTransform calls the method "PaintRenderingContext2D.setTransform".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) SetTransform(a float64, b float64, c float64, d float64, e float64, f float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DSetTransform(
		this.Ref(), js.Pointer(&_ok),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetTransformFunc returns the method "PaintRenderingContext2D.setTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) SetTransformFunc() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DSetTransformFunc(
			this.Ref(),
		),
	)
}

// SetTransform1 calls the method "PaintRenderingContext2D.setTransform".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) SetTransform1(transform DOMMatrix2DInit) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DSetTransform1(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&transform),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetTransform1Func returns the method "PaintRenderingContext2D.setTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) SetTransform1Func() (fn js.Func[func(transform DOMMatrix2DInit)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DSetTransform1Func(
			this.Ref(),
		),
	)
}

// SetTransform2 calls the method "PaintRenderingContext2D.setTransform".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) SetTransform2() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DSetTransform2(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetTransform2Func returns the method "PaintRenderingContext2D.setTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) SetTransform2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DSetTransform2Func(
			this.Ref(),
		),
	)
}

// ResetTransform calls the method "PaintRenderingContext2D.resetTransform".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) ResetTransform() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DResetTransform(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResetTransformFunc returns the method "PaintRenderingContext2D.resetTransform".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) ResetTransformFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DResetTransformFunc(
			this.Ref(),
		),
	)
}

// Save calls the method "PaintRenderingContext2D.save".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Save() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DSave(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SaveFunc returns the method "PaintRenderingContext2D.save".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) SaveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DSaveFunc(
			this.Ref(),
		),
	)
}

// Restore calls the method "PaintRenderingContext2D.restore".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Restore() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DRestore(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RestoreFunc returns the method "PaintRenderingContext2D.restore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) RestoreFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DRestoreFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "PaintRenderingContext2D.reset".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) Reset() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DReset(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ResetFunc returns the method "PaintRenderingContext2D.reset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DResetFunc(
			this.Ref(),
		),
	)
}

// IsContextLost calls the method "PaintRenderingContext2D.isContextLost".
//
// The returned bool will be false if there is no such method.
func (this PaintRenderingContext2D) IsContextLost() (bool, bool) {
	var _ok bool
	_ret := bindings.CallPaintRenderingContext2DIsContextLost(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// IsContextLostFunc returns the method "PaintRenderingContext2D.isContextLost".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintRenderingContext2D) IsContextLostFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsContextLostFunc(
			this.Ref(),
		),
	)
}

type PaintRenderingContext2DSettings struct {
	// Alpha is "PaintRenderingContext2DSettings.alpha"
	//
	// Optional, defaults to true.
	Alpha bool

	FFI_USE_Alpha bool // for Alpha.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaintRenderingContext2DSettings with all fields set.
func (p PaintRenderingContext2DSettings) FromRef(ref js.Ref) PaintRenderingContext2DSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PaintRenderingContext2DSettings PaintRenderingContext2DSettings [// PaintRenderingContext2DSettings] [0x1400093e820 0x1400093e8c0] 0x14000920858 {0 0}} in the application heap.
func (p PaintRenderingContext2DSettings) New() js.Ref {
	return bindings.PaintRenderingContext2DSettingsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaintRenderingContext2DSettings) UpdateFrom(ref js.Ref) {
	bindings.PaintRenderingContext2DSettingsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaintRenderingContext2DSettings) Update(ref js.Ref) {
	bindings.PaintRenderingContext2DSettingsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaintSize struct {
	ref js.Ref
}

func (this PaintSize) Once() PaintSize {
	this.Ref().Once()
	return this
}

func (this PaintSize) Ref() js.Ref {
	return this.ref
}

func (this PaintSize) FromRef(ref js.Ref) PaintSize {
	this.ref = ref
	return this
}

func (this PaintSize) Free() {
	this.Ref().Free()
}

// Width returns the value of property "PaintSize.width".
//
// The returned bool will be false if there is no such property.
func (this PaintSize) Width() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPaintSizeWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Height returns the value of property "PaintSize.height".
//
// The returned bool will be false if there is no such property.
func (this PaintSize) Height() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPaintSizeHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type PaintWorkletGlobalScope struct {
	WorkletGlobalScope
}

func (this PaintWorkletGlobalScope) Once() PaintWorkletGlobalScope {
	this.Ref().Once()
	return this
}

func (this PaintWorkletGlobalScope) Ref() js.Ref {
	return this.WorkletGlobalScope.Ref()
}

func (this PaintWorkletGlobalScope) FromRef(ref js.Ref) PaintWorkletGlobalScope {
	this.WorkletGlobalScope = this.WorkletGlobalScope.FromRef(ref)
	return this
}

func (this PaintWorkletGlobalScope) Free() {
	this.Ref().Free()
}

// DevicePixelRatio returns the value of property "PaintWorkletGlobalScope.devicePixelRatio".
//
// The returned bool will be false if there is no such property.
func (this PaintWorkletGlobalScope) DevicePixelRatio() (float64, bool) {
	var _ok bool
	_ret := bindings.GetPaintWorkletGlobalScopeDevicePixelRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// RegisterPaint calls the method "PaintWorkletGlobalScope.registerPaint".
//
// The returned bool will be false if there is no such method.
func (this PaintWorkletGlobalScope) RegisterPaint(name js.String, paintCtor js.Func[func()]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallPaintWorkletGlobalScopeRegisterPaint(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		paintCtor.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RegisterPaintFunc returns the method "PaintWorkletGlobalScope.registerPaint".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this PaintWorkletGlobalScope) RegisterPaintFunc() (fn js.Func[func(name js.String, paintCtor js.Func[func()])]) {
	return fn.FromRef(
		bindings.PaintWorkletGlobalScopeRegisterPaintFunc(
			this.Ref(),
		),
	)
}

func NewPasswordCredential(form HTMLFormElement) PasswordCredential {
	return PasswordCredential{}.FromRef(
		bindings.NewPasswordCredentialByPasswordCredential(
			form.Ref()),
	)
}

func NewPasswordCredentialByPasswordCredential1(data PasswordCredentialData) PasswordCredential {
	return PasswordCredential{}.FromRef(
		bindings.NewPasswordCredentialByPasswordCredential1(
			js.Pointer(&data)),
	)
}

type PasswordCredential struct {
	Credential
}

func (this PasswordCredential) Once() PasswordCredential {
	this.Ref().Once()
	return this
}

func (this PasswordCredential) Ref() js.Ref {
	return this.Credential.Ref()
}

func (this PasswordCredential) FromRef(ref js.Ref) PasswordCredential {
	this.Credential = this.Credential.FromRef(ref)
	return this
}

func (this PasswordCredential) Free() {
	this.Ref().Free()
}

// Password returns the value of property "PasswordCredential.password".
//
// The returned bool will be false if there is no such property.
func (this PasswordCredential) Password() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPasswordCredentialPassword(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name returns the value of property "PasswordCredential.name".
//
// The returned bool will be false if there is no such property.
func (this PasswordCredential) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPasswordCredentialName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// IconURL returns the value of property "PasswordCredential.iconURL".
//
// The returned bool will be false if there is no such property.
func (this PasswordCredential) IconURL() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetPasswordCredentialIconURL(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type PaymentComplete uint32

const (
	_ PaymentComplete = iota

	PaymentComplete_FAIL
	PaymentComplete_SUCCESS
	PaymentComplete_UNKNOWN
)

func (PaymentComplete) FromRef(str js.Ref) PaymentComplete {
	return PaymentComplete(bindings.ConstOfPaymentComplete(str))
}

func (x PaymentComplete) String() (string, bool) {
	switch x {
	case PaymentComplete_FAIL:
		return "fail", true
	case PaymentComplete_SUCCESS:
		return "success", true
	case PaymentComplete_UNKNOWN:
		return "unknown", true
	default:
		return "", false
	}
}

type PaymentCompleteDetails struct {
	// Data is "PaymentCompleteDetails.data"
	//
	// Optional, defaults to null.
	Data js.Object

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentCompleteDetails with all fields set.
func (p PaymentCompleteDetails) FromRef(ref js.Ref) PaymentCompleteDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PaymentCompleteDetails PaymentCompleteDetails [// PaymentCompleteDetails] [0x1400093eaa0] 0x14000920888 {0 0}} in the application heap.
func (p PaymentCompleteDetails) New() js.Ref {
	return bindings.PaymentCompleteDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentCompleteDetails) UpdateFrom(ref js.Ref) {
	bindings.PaymentCompleteDetailsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentCompleteDetails) Update(ref js.Ref) {
	bindings.PaymentCompleteDetailsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type PaymentItem struct {
	// Label is "PaymentItem.label"
	//
	// Required
	Label js.String
	// Amount is "PaymentItem.amount"
	//
	// Required
	Amount PaymentCurrencyAmount
	// Pending is "PaymentItem.pending"
	//
	// Optional, defaults to false.
	Pending bool

	FFI_USE_Pending bool // for Pending.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentItem with all fields set.
func (p PaymentItem) FromRef(ref js.Ref) PaymentItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 PaymentItem PaymentItem [// PaymentItem] [0x1400093eb40 0x1400093ebe0 0x1400093ec80 0x1400093ed20] 0x140009208e8 {0 0}} in the application heap.
func (p PaymentItem) New() js.Ref {
	return bindings.PaymentItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PaymentItem) UpdateFrom(ref js.Ref) {
	bindings.PaymentItemJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PaymentItem) Update(ref js.Ref) {
	bindings.PaymentItemJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

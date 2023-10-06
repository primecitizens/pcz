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

// New creates a new NavigationInterceptOptions in the application heap.
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

func NewNavigateEvent(typ js.String, eventInitDict NavigateEventInit) (ret NavigateEvent) {
	ret.ref = bindings.NewNavigateEventByNavigateEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this NavigateEvent) NavigationType() (ret NavigationType, ok bool) {
	ok = js.True == bindings.GetNavigateEventNavigationType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Destination returns the value of property "NavigateEvent.destination".
//
// It returns ok=false if there is no such property.
func (this NavigateEvent) Destination() (ret NavigationDestination, ok bool) {
	ok = js.True == bindings.GetNavigateEventDestination(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CanIntercept returns the value of property "NavigateEvent.canIntercept".
//
// It returns ok=false if there is no such property.
func (this NavigateEvent) CanIntercept() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigateEventCanIntercept(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// UserInitiated returns the value of property "NavigateEvent.userInitiated".
//
// It returns ok=false if there is no such property.
func (this NavigateEvent) UserInitiated() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigateEventUserInitiated(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HashChange returns the value of property "NavigateEvent.hashChange".
//
// It returns ok=false if there is no such property.
func (this NavigateEvent) HashChange() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigateEventHashChange(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Signal returns the value of property "NavigateEvent.signal".
//
// It returns ok=false if there is no such property.
func (this NavigateEvent) Signal() (ret AbortSignal, ok bool) {
	ok = js.True == bindings.GetNavigateEventSignal(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FormData returns the value of property "NavigateEvent.formData".
//
// It returns ok=false if there is no such property.
func (this NavigateEvent) FormData() (ret FormData, ok bool) {
	ok = js.True == bindings.GetNavigateEventFormData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DownloadRequest returns the value of property "NavigateEvent.downloadRequest".
//
// It returns ok=false if there is no such property.
func (this NavigateEvent) DownloadRequest() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNavigateEventDownloadRequest(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Info returns the value of property "NavigateEvent.info".
//
// It returns ok=false if there is no such property.
func (this NavigateEvent) Info() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetNavigateEventInfo(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasUAVisualTransition returns the value of property "NavigateEvent.hasUAVisualTransition".
//
// It returns ok=false if there is no such property.
func (this NavigateEvent) HasUAVisualTransition() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigateEventHasUAVisualTransition(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasIntercept returns true if the method "NavigateEvent.intercept" exists.
func (this NavigateEvent) HasIntercept() bool {
	return js.True == bindings.HasNavigateEventIntercept(
		this.Ref(),
	)
}

// InterceptFunc returns the method "NavigateEvent.intercept".
func (this NavigateEvent) InterceptFunc() (fn js.Func[func(options NavigationInterceptOptions)]) {
	return fn.FromRef(
		bindings.NavigateEventInterceptFunc(
			this.Ref(),
		),
	)
}

// Intercept calls the method "NavigateEvent.intercept".
func (this NavigateEvent) Intercept(options NavigationInterceptOptions) (ret js.Void) {
	bindings.CallNavigateEventIntercept(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryIntercept calls the method "NavigateEvent.intercept"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NavigateEvent) TryIntercept(options NavigationInterceptOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigateEventIntercept(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasIntercept1 returns true if the method "NavigateEvent.intercept" exists.
func (this NavigateEvent) HasIntercept1() bool {
	return js.True == bindings.HasNavigateEventIntercept1(
		this.Ref(),
	)
}

// Intercept1Func returns the method "NavigateEvent.intercept".
func (this NavigateEvent) Intercept1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.NavigateEventIntercept1Func(
			this.Ref(),
		),
	)
}

// Intercept1 calls the method "NavigateEvent.intercept".
func (this NavigateEvent) Intercept1() (ret js.Void) {
	bindings.CallNavigateEventIntercept1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIntercept1 calls the method "NavigateEvent.intercept"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NavigateEvent) TryIntercept1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigateEventIntercept1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasScroll returns true if the method "NavigateEvent.scroll" exists.
func (this NavigateEvent) HasScroll() bool {
	return js.True == bindings.HasNavigateEventScroll(
		this.Ref(),
	)
}

// ScrollFunc returns the method "NavigateEvent.scroll".
func (this NavigateEvent) ScrollFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.NavigateEventScrollFunc(
			this.Ref(),
		),
	)
}

// Scroll calls the method "NavigateEvent.scroll".
func (this NavigateEvent) Scroll() (ret js.Void) {
	bindings.CallNavigateEventScroll(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryScroll calls the method "NavigateEvent.scroll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this NavigateEvent) TryScroll() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigateEventScroll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "NavigationCurrentEntryChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "NavigationCurrentEntryChangeEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new NavigationCurrentEntryChangeEventInit in the application heap.
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

func NewNavigationCurrentEntryChangeEvent(typ js.String, eventInitDict NavigationCurrentEntryChangeEventInit) (ret NavigationCurrentEntryChangeEvent) {
	ret.ref = bindings.NewNavigationCurrentEntryChangeEventByNavigationCurrentEntryChangeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this NavigationCurrentEntryChangeEvent) NavigationType() (ret NavigationType, ok bool) {
	ok = js.True == bindings.GetNavigationCurrentEntryChangeEventNavigationType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// From returns the value of property "NavigationCurrentEntryChangeEvent.from".
//
// It returns ok=false if there is no such property.
func (this NavigationCurrentEntryChangeEvent) From() (ret NavigationHistoryEntry, ok bool) {
	ok = js.True == bindings.GetNavigationCurrentEntryChangeEventFrom(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_Detail MUST be set to true to make this field effective.
	Detail int32
	// Bubbles is "NavigationEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "NavigationEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "NavigationEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new NavigationEventInit in the application heap.
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

func NewNavigationEvent(typ js.String, eventInitDict NavigationEventInit) (ret NavigationEvent) {
	ret.ref = bindings.NewNavigationEventByNavigationEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewNavigationEventByNavigationEvent1(typ js.String) (ret NavigationEvent) {
	ret.ref = bindings.NewNavigationEventByNavigationEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this NavigationEvent) Dir() (ret SpatialNavigationDirection, ok bool) {
	ok = js.True == bindings.GetNavigationEventDir(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RelatedTarget returns the value of property "NavigationEvent.relatedTarget".
//
// It returns ok=false if there is no such property.
func (this NavigationEvent) RelatedTarget() (ret EventTarget, ok bool) {
	ok = js.True == bindings.GetNavigationEventRelatedTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "NotificationEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "NotificationEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new NotificationEventInit in the application heap.
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

func NewNotificationEvent(typ js.String, eventInitDict NotificationEventInit) (ret NotificationEvent) {
	ret.ref = bindings.NewNotificationEventByNotificationEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this NotificationEvent) Notification() (ret Notification, ok bool) {
	ok = js.True == bindings.GetNotificationEventNotification(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Action returns the value of property "NotificationEvent.action".
//
// It returns ok=false if there is no such property.
func (this NotificationEvent) Action() (ret js.String, ok bool) {
	ok = js.True == bindings.GetNotificationEventAction(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasEnableiOES returns true if the method "OES_draw_buffers_indexed.enableiOES" exists.
func (this OES_draw_buffers_indexed) HasEnableiOES() bool {
	return js.True == bindings.HasOES_draw_buffers_indexedEnableiOES(
		this.Ref(),
	)
}

// EnableiOESFunc returns the method "OES_draw_buffers_indexed.enableiOES".
func (this OES_draw_buffers_indexed) EnableiOESFunc() (fn js.Func[func(target GLenum, index GLuint)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedEnableiOESFunc(
			this.Ref(),
		),
	)
}

// EnableiOES calls the method "OES_draw_buffers_indexed.enableiOES".
func (this OES_draw_buffers_indexed) EnableiOES(target GLenum, index GLuint) (ret js.Void) {
	bindings.CallOES_draw_buffers_indexedEnableiOES(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(index),
	)

	return
}

// TryEnableiOES calls the method "OES_draw_buffers_indexed.enableiOES"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OES_draw_buffers_indexed) TryEnableiOES(target GLenum, index GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOES_draw_buffers_indexedEnableiOES(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(index),
	)

	return
}

// HasDisableiOES returns true if the method "OES_draw_buffers_indexed.disableiOES" exists.
func (this OES_draw_buffers_indexed) HasDisableiOES() bool {
	return js.True == bindings.HasOES_draw_buffers_indexedDisableiOES(
		this.Ref(),
	)
}

// DisableiOESFunc returns the method "OES_draw_buffers_indexed.disableiOES".
func (this OES_draw_buffers_indexed) DisableiOESFunc() (fn js.Func[func(target GLenum, index GLuint)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedDisableiOESFunc(
			this.Ref(),
		),
	)
}

// DisableiOES calls the method "OES_draw_buffers_indexed.disableiOES".
func (this OES_draw_buffers_indexed) DisableiOES(target GLenum, index GLuint) (ret js.Void) {
	bindings.CallOES_draw_buffers_indexedDisableiOES(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(index),
	)

	return
}

// TryDisableiOES calls the method "OES_draw_buffers_indexed.disableiOES"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OES_draw_buffers_indexed) TryDisableiOES(target GLenum, index GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOES_draw_buffers_indexedDisableiOES(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(index),
	)

	return
}

// HasBlendEquationiOES returns true if the method "OES_draw_buffers_indexed.blendEquationiOES" exists.
func (this OES_draw_buffers_indexed) HasBlendEquationiOES() bool {
	return js.True == bindings.HasOES_draw_buffers_indexedBlendEquationiOES(
		this.Ref(),
	)
}

// BlendEquationiOESFunc returns the method "OES_draw_buffers_indexed.blendEquationiOES".
func (this OES_draw_buffers_indexed) BlendEquationiOESFunc() (fn js.Func[func(buf GLuint, mode GLenum)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedBlendEquationiOESFunc(
			this.Ref(),
		),
	)
}

// BlendEquationiOES calls the method "OES_draw_buffers_indexed.blendEquationiOES".
func (this OES_draw_buffers_indexed) BlendEquationiOES(buf GLuint, mode GLenum) (ret js.Void) {
	bindings.CallOES_draw_buffers_indexedBlendEquationiOES(
		this.Ref(), js.Pointer(&ret),
		uint32(buf),
		uint32(mode),
	)

	return
}

// TryBlendEquationiOES calls the method "OES_draw_buffers_indexed.blendEquationiOES"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OES_draw_buffers_indexed) TryBlendEquationiOES(buf GLuint, mode GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOES_draw_buffers_indexedBlendEquationiOES(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buf),
		uint32(mode),
	)

	return
}

// HasBlendEquationSeparateiOES returns true if the method "OES_draw_buffers_indexed.blendEquationSeparateiOES" exists.
func (this OES_draw_buffers_indexed) HasBlendEquationSeparateiOES() bool {
	return js.True == bindings.HasOES_draw_buffers_indexedBlendEquationSeparateiOES(
		this.Ref(),
	)
}

// BlendEquationSeparateiOESFunc returns the method "OES_draw_buffers_indexed.blendEquationSeparateiOES".
func (this OES_draw_buffers_indexed) BlendEquationSeparateiOESFunc() (fn js.Func[func(buf GLuint, modeRGB GLenum, modeAlpha GLenum)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedBlendEquationSeparateiOESFunc(
			this.Ref(),
		),
	)
}

// BlendEquationSeparateiOES calls the method "OES_draw_buffers_indexed.blendEquationSeparateiOES".
func (this OES_draw_buffers_indexed) BlendEquationSeparateiOES(buf GLuint, modeRGB GLenum, modeAlpha GLenum) (ret js.Void) {
	bindings.CallOES_draw_buffers_indexedBlendEquationSeparateiOES(
		this.Ref(), js.Pointer(&ret),
		uint32(buf),
		uint32(modeRGB),
		uint32(modeAlpha),
	)

	return
}

// TryBlendEquationSeparateiOES calls the method "OES_draw_buffers_indexed.blendEquationSeparateiOES"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OES_draw_buffers_indexed) TryBlendEquationSeparateiOES(buf GLuint, modeRGB GLenum, modeAlpha GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOES_draw_buffers_indexedBlendEquationSeparateiOES(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buf),
		uint32(modeRGB),
		uint32(modeAlpha),
	)

	return
}

// HasBlendFunciOES returns true if the method "OES_draw_buffers_indexed.blendFunciOES" exists.
func (this OES_draw_buffers_indexed) HasBlendFunciOES() bool {
	return js.True == bindings.HasOES_draw_buffers_indexedBlendFunciOES(
		this.Ref(),
	)
}

// BlendFunciOESFunc returns the method "OES_draw_buffers_indexed.blendFunciOES".
func (this OES_draw_buffers_indexed) BlendFunciOESFunc() (fn js.Func[func(buf GLuint, src GLenum, dst GLenum)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedBlendFunciOESFunc(
			this.Ref(),
		),
	)
}

// BlendFunciOES calls the method "OES_draw_buffers_indexed.blendFunciOES".
func (this OES_draw_buffers_indexed) BlendFunciOES(buf GLuint, src GLenum, dst GLenum) (ret js.Void) {
	bindings.CallOES_draw_buffers_indexedBlendFunciOES(
		this.Ref(), js.Pointer(&ret),
		uint32(buf),
		uint32(src),
		uint32(dst),
	)

	return
}

// TryBlendFunciOES calls the method "OES_draw_buffers_indexed.blendFunciOES"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OES_draw_buffers_indexed) TryBlendFunciOES(buf GLuint, src GLenum, dst GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOES_draw_buffers_indexedBlendFunciOES(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buf),
		uint32(src),
		uint32(dst),
	)

	return
}

// HasBlendFuncSeparateiOES returns true if the method "OES_draw_buffers_indexed.blendFuncSeparateiOES" exists.
func (this OES_draw_buffers_indexed) HasBlendFuncSeparateiOES() bool {
	return js.True == bindings.HasOES_draw_buffers_indexedBlendFuncSeparateiOES(
		this.Ref(),
	)
}

// BlendFuncSeparateiOESFunc returns the method "OES_draw_buffers_indexed.blendFuncSeparateiOES".
func (this OES_draw_buffers_indexed) BlendFuncSeparateiOESFunc() (fn js.Func[func(buf GLuint, srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedBlendFuncSeparateiOESFunc(
			this.Ref(),
		),
	)
}

// BlendFuncSeparateiOES calls the method "OES_draw_buffers_indexed.blendFuncSeparateiOES".
func (this OES_draw_buffers_indexed) BlendFuncSeparateiOES(buf GLuint, srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) (ret js.Void) {
	bindings.CallOES_draw_buffers_indexedBlendFuncSeparateiOES(
		this.Ref(), js.Pointer(&ret),
		uint32(buf),
		uint32(srcRGB),
		uint32(dstRGB),
		uint32(srcAlpha),
		uint32(dstAlpha),
	)

	return
}

// TryBlendFuncSeparateiOES calls the method "OES_draw_buffers_indexed.blendFuncSeparateiOES"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OES_draw_buffers_indexed) TryBlendFuncSeparateiOES(buf GLuint, srcRGB GLenum, dstRGB GLenum, srcAlpha GLenum, dstAlpha GLenum) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOES_draw_buffers_indexedBlendFuncSeparateiOES(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buf),
		uint32(srcRGB),
		uint32(dstRGB),
		uint32(srcAlpha),
		uint32(dstAlpha),
	)

	return
}

// HasColorMaskiOES returns true if the method "OES_draw_buffers_indexed.colorMaskiOES" exists.
func (this OES_draw_buffers_indexed) HasColorMaskiOES() bool {
	return js.True == bindings.HasOES_draw_buffers_indexedColorMaskiOES(
		this.Ref(),
	)
}

// ColorMaskiOESFunc returns the method "OES_draw_buffers_indexed.colorMaskiOES".
func (this OES_draw_buffers_indexed) ColorMaskiOESFunc() (fn js.Func[func(buf GLuint, r GLboolean, g GLboolean, b GLboolean, a GLboolean)]) {
	return fn.FromRef(
		bindings.OES_draw_buffers_indexedColorMaskiOESFunc(
			this.Ref(),
		),
	)
}

// ColorMaskiOES calls the method "OES_draw_buffers_indexed.colorMaskiOES".
func (this OES_draw_buffers_indexed) ColorMaskiOES(buf GLuint, r GLboolean, g GLboolean, b GLboolean, a GLboolean) (ret js.Void) {
	bindings.CallOES_draw_buffers_indexedColorMaskiOES(
		this.Ref(), js.Pointer(&ret),
		uint32(buf),
		js.Bool(bool(r)),
		js.Bool(bool(g)),
		js.Bool(bool(b)),
		js.Bool(bool(a)),
	)

	return
}

// TryColorMaskiOES calls the method "OES_draw_buffers_indexed.colorMaskiOES"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OES_draw_buffers_indexed) TryColorMaskiOES(buf GLuint, r GLboolean, g GLboolean, b GLboolean, a GLboolean) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOES_draw_buffers_indexedColorMaskiOES(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(buf),
		js.Bool(bool(r)),
		js.Bool(bool(g)),
		js.Bool(bool(b)),
		js.Bool(bool(a)),
	)

	return
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

// HasCreateVertexArrayOES returns true if the method "OES_vertex_array_object.createVertexArrayOES" exists.
func (this OES_vertex_array_object) HasCreateVertexArrayOES() bool {
	return js.True == bindings.HasOES_vertex_array_objectCreateVertexArrayOES(
		this.Ref(),
	)
}

// CreateVertexArrayOESFunc returns the method "OES_vertex_array_object.createVertexArrayOES".
func (this OES_vertex_array_object) CreateVertexArrayOESFunc() (fn js.Func[func() WebGLVertexArrayObjectOES]) {
	return fn.FromRef(
		bindings.OES_vertex_array_objectCreateVertexArrayOESFunc(
			this.Ref(),
		),
	)
}

// CreateVertexArrayOES calls the method "OES_vertex_array_object.createVertexArrayOES".
func (this OES_vertex_array_object) CreateVertexArrayOES() (ret WebGLVertexArrayObjectOES) {
	bindings.CallOES_vertex_array_objectCreateVertexArrayOES(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateVertexArrayOES calls the method "OES_vertex_array_object.createVertexArrayOES"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OES_vertex_array_object) TryCreateVertexArrayOES() (ret WebGLVertexArrayObjectOES, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOES_vertex_array_objectCreateVertexArrayOES(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteVertexArrayOES returns true if the method "OES_vertex_array_object.deleteVertexArrayOES" exists.
func (this OES_vertex_array_object) HasDeleteVertexArrayOES() bool {
	return js.True == bindings.HasOES_vertex_array_objectDeleteVertexArrayOES(
		this.Ref(),
	)
}

// DeleteVertexArrayOESFunc returns the method "OES_vertex_array_object.deleteVertexArrayOES".
func (this OES_vertex_array_object) DeleteVertexArrayOESFunc() (fn js.Func[func(arrayObject WebGLVertexArrayObjectOES)]) {
	return fn.FromRef(
		bindings.OES_vertex_array_objectDeleteVertexArrayOESFunc(
			this.Ref(),
		),
	)
}

// DeleteVertexArrayOES calls the method "OES_vertex_array_object.deleteVertexArrayOES".
func (this OES_vertex_array_object) DeleteVertexArrayOES(arrayObject WebGLVertexArrayObjectOES) (ret js.Void) {
	bindings.CallOES_vertex_array_objectDeleteVertexArrayOES(
		this.Ref(), js.Pointer(&ret),
		arrayObject.Ref(),
	)

	return
}

// TryDeleteVertexArrayOES calls the method "OES_vertex_array_object.deleteVertexArrayOES"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OES_vertex_array_object) TryDeleteVertexArrayOES(arrayObject WebGLVertexArrayObjectOES) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOES_vertex_array_objectDeleteVertexArrayOES(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		arrayObject.Ref(),
	)

	return
}

// HasIsVertexArrayOES returns true if the method "OES_vertex_array_object.isVertexArrayOES" exists.
func (this OES_vertex_array_object) HasIsVertexArrayOES() bool {
	return js.True == bindings.HasOES_vertex_array_objectIsVertexArrayOES(
		this.Ref(),
	)
}

// IsVertexArrayOESFunc returns the method "OES_vertex_array_object.isVertexArrayOES".
func (this OES_vertex_array_object) IsVertexArrayOESFunc() (fn js.Func[func(arrayObject WebGLVertexArrayObjectOES) GLboolean]) {
	return fn.FromRef(
		bindings.OES_vertex_array_objectIsVertexArrayOESFunc(
			this.Ref(),
		),
	)
}

// IsVertexArrayOES calls the method "OES_vertex_array_object.isVertexArrayOES".
func (this OES_vertex_array_object) IsVertexArrayOES(arrayObject WebGLVertexArrayObjectOES) (ret GLboolean) {
	bindings.CallOES_vertex_array_objectIsVertexArrayOES(
		this.Ref(), js.Pointer(&ret),
		arrayObject.Ref(),
	)

	return
}

// TryIsVertexArrayOES calls the method "OES_vertex_array_object.isVertexArrayOES"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OES_vertex_array_object) TryIsVertexArrayOES(arrayObject WebGLVertexArrayObjectOES) (ret GLboolean, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOES_vertex_array_objectIsVertexArrayOES(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		arrayObject.Ref(),
	)

	return
}

// HasBindVertexArrayOES returns true if the method "OES_vertex_array_object.bindVertexArrayOES" exists.
func (this OES_vertex_array_object) HasBindVertexArrayOES() bool {
	return js.True == bindings.HasOES_vertex_array_objectBindVertexArrayOES(
		this.Ref(),
	)
}

// BindVertexArrayOESFunc returns the method "OES_vertex_array_object.bindVertexArrayOES".
func (this OES_vertex_array_object) BindVertexArrayOESFunc() (fn js.Func[func(arrayObject WebGLVertexArrayObjectOES)]) {
	return fn.FromRef(
		bindings.OES_vertex_array_objectBindVertexArrayOESFunc(
			this.Ref(),
		),
	)
}

// BindVertexArrayOES calls the method "OES_vertex_array_object.bindVertexArrayOES".
func (this OES_vertex_array_object) BindVertexArrayOES(arrayObject WebGLVertexArrayObjectOES) (ret js.Void) {
	bindings.CallOES_vertex_array_objectBindVertexArrayOES(
		this.Ref(), js.Pointer(&ret),
		arrayObject.Ref(),
	)

	return
}

// TryBindVertexArrayOES calls the method "OES_vertex_array_object.bindVertexArrayOES"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OES_vertex_array_object) TryBindVertexArrayOES(arrayObject WebGLVertexArrayObjectOES) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOES_vertex_array_objectBindVertexArrayOES(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		arrayObject.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this OTPCredential) Code() (ret js.String, ok bool) {
	ok = js.True == bindings.GetOTPCredentialCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasFramebufferTextureMultiviewOVR returns true if the method "OVR_multiview2.framebufferTextureMultiviewOVR" exists.
func (this OVR_multiview2) HasFramebufferTextureMultiviewOVR() bool {
	return js.True == bindings.HasOVR_multiview2FramebufferTextureMultiviewOVR(
		this.Ref(),
	)
}

// FramebufferTextureMultiviewOVRFunc returns the method "OVR_multiview2.framebufferTextureMultiviewOVR".
func (this OVR_multiview2) FramebufferTextureMultiviewOVRFunc() (fn js.Func[func(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, baseViewIndex GLint, numViews GLsizei)]) {
	return fn.FromRef(
		bindings.OVR_multiview2FramebufferTextureMultiviewOVRFunc(
			this.Ref(),
		),
	)
}

// FramebufferTextureMultiviewOVR calls the method "OVR_multiview2.framebufferTextureMultiviewOVR".
func (this OVR_multiview2) FramebufferTextureMultiviewOVR(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, baseViewIndex GLint, numViews GLsizei) (ret js.Void) {
	bindings.CallOVR_multiview2FramebufferTextureMultiviewOVR(
		this.Ref(), js.Pointer(&ret),
		uint32(target),
		uint32(attachment),
		texture.Ref(),
		int32(level),
		int32(baseViewIndex),
		int32(numViews),
	)

	return
}

// TryFramebufferTextureMultiviewOVR calls the method "OVR_multiview2.framebufferTextureMultiviewOVR"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OVR_multiview2) TryFramebufferTextureMultiviewOVR(target GLenum, attachment GLenum, texture WebGLTexture, level GLint, baseViewIndex GLint, numViews GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOVR_multiview2FramebufferTextureMultiviewOVR(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(target),
		uint32(attachment),
		texture.Ref(),
		int32(level),
		int32(baseViewIndex),
		int32(numViews),
	)

	return
}

type OfflineAudioCompletionEventInit struct {
	// RenderedBuffer is "OfflineAudioCompletionEventInit.renderedBuffer"
	//
	// Required
	RenderedBuffer AudioBuffer
	// Bubbles is "OfflineAudioCompletionEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "OfflineAudioCompletionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "OfflineAudioCompletionEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new OfflineAudioCompletionEventInit in the application heap.
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

func NewOfflineAudioCompletionEvent(typ js.String, eventInitDict OfflineAudioCompletionEventInit) (ret OfflineAudioCompletionEvent) {
	ret.ref = bindings.NewOfflineAudioCompletionEventByOfflineAudioCompletionEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this OfflineAudioCompletionEvent) RenderedBuffer() (ret AudioBuffer, ok bool) {
	ok = js.True == bindings.GetOfflineAudioCompletionEventRenderedBuffer(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type OfflineAudioContextOptions struct {
	// NumberOfChannels is "OfflineAudioContextOptions.numberOfChannels"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_NumberOfChannels MUST be set to true to make this field effective.
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

// New creates a new OfflineAudioContextOptions in the application heap.
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

func NewOfflineAudioContext(contextOptions OfflineAudioContextOptions) (ret OfflineAudioContext) {
	ret.ref = bindings.NewOfflineAudioContextByOfflineAudioContext(
		js.Pointer(&contextOptions))
	return
}

func NewOfflineAudioContextByOfflineAudioContext1(numberOfChannels uint32, length uint32, sampleRate float32) (ret OfflineAudioContext) {
	ret.ref = bindings.NewOfflineAudioContextByOfflineAudioContext1(
		uint32(numberOfChannels),
		uint32(length),
		float32(sampleRate))
	return
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
// It returns ok=false if there is no such property.
func (this OfflineAudioContext) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetOfflineAudioContextLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasStartRendering returns true if the method "OfflineAudioContext.startRendering" exists.
func (this OfflineAudioContext) HasStartRendering() bool {
	return js.True == bindings.HasOfflineAudioContextStartRendering(
		this.Ref(),
	)
}

// StartRenderingFunc returns the method "OfflineAudioContext.startRendering".
func (this OfflineAudioContext) StartRenderingFunc() (fn js.Func[func() js.Promise[AudioBuffer]]) {
	return fn.FromRef(
		bindings.OfflineAudioContextStartRenderingFunc(
			this.Ref(),
		),
	)
}

// StartRendering calls the method "OfflineAudioContext.startRendering".
func (this OfflineAudioContext) StartRendering() (ret js.Promise[AudioBuffer]) {
	bindings.CallOfflineAudioContextStartRendering(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStartRendering calls the method "OfflineAudioContext.startRendering"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OfflineAudioContext) TryStartRendering() (ret js.Promise[AudioBuffer], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOfflineAudioContextStartRendering(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasResume returns true if the method "OfflineAudioContext.resume" exists.
func (this OfflineAudioContext) HasResume() bool {
	return js.True == bindings.HasOfflineAudioContextResume(
		this.Ref(),
	)
}

// ResumeFunc returns the method "OfflineAudioContext.resume".
func (this OfflineAudioContext) ResumeFunc() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.OfflineAudioContextResumeFunc(
			this.Ref(),
		),
	)
}

// Resume calls the method "OfflineAudioContext.resume".
func (this OfflineAudioContext) Resume() (ret js.Promise[js.Void]) {
	bindings.CallOfflineAudioContextResume(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryResume calls the method "OfflineAudioContext.resume"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OfflineAudioContext) TryResume() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOfflineAudioContextResume(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSuspend returns true if the method "OfflineAudioContext.suspend" exists.
func (this OfflineAudioContext) HasSuspend() bool {
	return js.True == bindings.HasOfflineAudioContextSuspend(
		this.Ref(),
	)
}

// SuspendFunc returns the method "OfflineAudioContext.suspend".
func (this OfflineAudioContext) SuspendFunc() (fn js.Func[func(suspendTime float64) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.OfflineAudioContextSuspendFunc(
			this.Ref(),
		),
	)
}

// Suspend calls the method "OfflineAudioContext.suspend".
func (this OfflineAudioContext) Suspend(suspendTime float64) (ret js.Promise[js.Void]) {
	bindings.CallOfflineAudioContextSuspend(
		this.Ref(), js.Pointer(&ret),
		float64(suspendTime),
	)

	return
}

// TrySuspend calls the method "OfflineAudioContext.suspend"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OfflineAudioContext) TrySuspend(suspendTime float64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryOfflineAudioContextSuspend(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(suspendTime),
	)

	return
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
// It returns ok=false if there is no such property.
func (this OrientationSensor) Quaternion() (ret js.FrozenArray[float64], ok bool) {
	ok = js.True == bindings.GetOrientationSensorQuaternion(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasPopulateMatrix returns true if the method "OrientationSensor.populateMatrix" exists.
func (this OrientationSensor) HasPopulateMatrix() bool {
	return js.True == bindings.HasOrientationSensorPopulateMatrix(
		this.Ref(),
	)
}

// PopulateMatrixFunc returns the method "OrientationSensor.populateMatrix".
func (this OrientationSensor) PopulateMatrixFunc() (fn js.Func[func(targetMatrix RotationMatrixType)]) {
	return fn.FromRef(
		bindings.OrientationSensorPopulateMatrixFunc(
			this.Ref(),
		),
	)
}

// PopulateMatrix calls the method "OrientationSensor.populateMatrix".
func (this OrientationSensor) PopulateMatrix(targetMatrix RotationMatrixType) (ret js.Void) {
	bindings.CallOrientationSensorPopulateMatrix(
		this.Ref(), js.Pointer(&ret),
		targetMatrix.Ref(),
	)

	return
}

// TryPopulateMatrix calls the method "OrientationSensor.populateMatrix"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this OrientationSensor) TryPopulateMatrix(targetMatrix RotationMatrixType) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOrientationSensorPopulateMatrix(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		targetMatrix.Ref(),
	)

	return
}

func NewOverconstrainedError(constraint js.String, message js.String) (ret OverconstrainedError) {
	ret.ref = bindings.NewOverconstrainedErrorByOverconstrainedError(
		constraint.Ref(),
		message.Ref())
	return
}

func NewOverconstrainedErrorByOverconstrainedError1(constraint js.String) (ret OverconstrainedError) {
	ret.ref = bindings.NewOverconstrainedErrorByOverconstrainedError1(
		constraint.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this OverconstrainedError) Constraint() (ret js.String, ok bool) {
	ok = js.True == bindings.GetOverconstrainedErrorConstraint(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type PageTransitionEventInit struct {
	// Persisted is "PageTransitionEventInit.persisted"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Persisted MUST be set to true to make this field effective.
	Persisted bool
	// Bubbles is "PageTransitionEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "PageTransitionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "PageTransitionEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
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

// New creates a new PageTransitionEventInit in the application heap.
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

func NewPageTransitionEvent(typ js.String, eventInitDict PageTransitionEventInit) (ret PageTransitionEvent) {
	ret.ref = bindings.NewPageTransitionEventByPageTransitionEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewPageTransitionEventByPageTransitionEvent1(typ js.String) (ret PageTransitionEvent) {
	ret.ref = bindings.NewPageTransitionEventByPageTransitionEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this PageTransitionEvent) Persisted() (ret bool, ok bool) {
	ok = js.True == bindings.GetPageTransitionEventPersisted(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) LineWidth() (ret float64, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DLineWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLineWidth sets the value of property "PaintRenderingContext2D.lineWidth" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) LineCap() (ret CanvasLineCap, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DLineCap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLineCap sets the value of property "PaintRenderingContext2D.lineCap" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) LineJoin() (ret CanvasLineJoin, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DLineJoin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLineJoin sets the value of property "PaintRenderingContext2D.lineJoin" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) MiterLimit() (ret float64, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DMiterLimit(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMiterLimit sets the value of property "PaintRenderingContext2D.miterLimit" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) LineDashOffset() (ret float64, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DLineDashOffset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLineDashOffset sets the value of property "PaintRenderingContext2D.lineDashOffset" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) ShadowOffsetX() (ret float64, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DShadowOffsetX(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetShadowOffsetX sets the value of property "PaintRenderingContext2D.shadowOffsetX" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) ShadowOffsetY() (ret float64, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DShadowOffsetY(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetShadowOffsetY sets the value of property "PaintRenderingContext2D.shadowOffsetY" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) ShadowBlur() (ret float64, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DShadowBlur(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetShadowBlur sets the value of property "PaintRenderingContext2D.shadowBlur" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) ShadowColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DShadowColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetShadowColor sets the value of property "PaintRenderingContext2D.shadowColor" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) StrokeStyle() (ret OneOf_String_CanvasGradient_CanvasPattern, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DStrokeStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetStrokeStyle sets the value of property "PaintRenderingContext2D.strokeStyle" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) FillStyle() (ret OneOf_String_CanvasGradient_CanvasPattern, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DFillStyle(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFillStyle sets the value of property "PaintRenderingContext2D.fillStyle" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) ImageSmoothingEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DImageSmoothingEnabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetImageSmoothingEnabled sets the value of property "PaintRenderingContext2D.imageSmoothingEnabled" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) ImageSmoothingQuality() (ret ImageSmoothingQuality, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DImageSmoothingQuality(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetImageSmoothingQuality sets the value of property "PaintRenderingContext2D.imageSmoothingQuality" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) GlobalAlpha() (ret float64, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DGlobalAlpha(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetGlobalAlpha sets the value of property "PaintRenderingContext2D.globalAlpha" to val.
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
// It returns ok=false if there is no such property.
func (this PaintRenderingContext2D) GlobalCompositeOperation() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPaintRenderingContext2DGlobalCompositeOperation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetGlobalCompositeOperation sets the value of property "PaintRenderingContext2D.globalCompositeOperation" to val.
//
// It returns false if the property cannot be set.
func (this PaintRenderingContext2D) SetGlobalCompositeOperation(val js.String) bool {
	return js.True == bindings.SetPaintRenderingContext2DGlobalCompositeOperation(
		this.Ref(),
		val.Ref(),
	)
}

// HasClosePath returns true if the method "PaintRenderingContext2D.closePath" exists.
func (this PaintRenderingContext2D) HasClosePath() bool {
	return js.True == bindings.HasPaintRenderingContext2DClosePath(
		this.Ref(),
	)
}

// ClosePathFunc returns the method "PaintRenderingContext2D.closePath".
func (this PaintRenderingContext2D) ClosePathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClosePathFunc(
			this.Ref(),
		),
	)
}

// ClosePath calls the method "PaintRenderingContext2D.closePath".
func (this PaintRenderingContext2D) ClosePath() (ret js.Void) {
	bindings.CallPaintRenderingContext2DClosePath(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClosePath calls the method "PaintRenderingContext2D.closePath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryClosePath() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DClosePath(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasMoveTo returns true if the method "PaintRenderingContext2D.moveTo" exists.
func (this PaintRenderingContext2D) HasMoveTo() bool {
	return js.True == bindings.HasPaintRenderingContext2DMoveTo(
		this.Ref(),
	)
}

// MoveToFunc returns the method "PaintRenderingContext2D.moveTo".
func (this PaintRenderingContext2D) MoveToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DMoveToFunc(
			this.Ref(),
		),
	)
}

// MoveTo calls the method "PaintRenderingContext2D.moveTo".
func (this PaintRenderingContext2D) MoveTo(x float64, y float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DMoveTo(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryMoveTo calls the method "PaintRenderingContext2D.moveTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryMoveTo(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DMoveTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasLineTo returns true if the method "PaintRenderingContext2D.lineTo" exists.
func (this PaintRenderingContext2D) HasLineTo() bool {
	return js.True == bindings.HasPaintRenderingContext2DLineTo(
		this.Ref(),
	)
}

// LineToFunc returns the method "PaintRenderingContext2D.lineTo".
func (this PaintRenderingContext2D) LineToFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DLineToFunc(
			this.Ref(),
		),
	)
}

// LineTo calls the method "PaintRenderingContext2D.lineTo".
func (this PaintRenderingContext2D) LineTo(x float64, y float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DLineTo(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryLineTo calls the method "PaintRenderingContext2D.lineTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryLineTo(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DLineTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasQuadraticCurveTo returns true if the method "PaintRenderingContext2D.quadraticCurveTo" exists.
func (this PaintRenderingContext2D) HasQuadraticCurveTo() bool {
	return js.True == bindings.HasPaintRenderingContext2DQuadraticCurveTo(
		this.Ref(),
	)
}

// QuadraticCurveToFunc returns the method "PaintRenderingContext2D.quadraticCurveTo".
func (this PaintRenderingContext2D) QuadraticCurveToFunc() (fn js.Func[func(cpx float64, cpy float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DQuadraticCurveToFunc(
			this.Ref(),
		),
	)
}

// QuadraticCurveTo calls the method "PaintRenderingContext2D.quadraticCurveTo".
func (this PaintRenderingContext2D) QuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DQuadraticCurveTo(
		this.Ref(), js.Pointer(&ret),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	return
}

// TryQuadraticCurveTo calls the method "PaintRenderingContext2D.quadraticCurveTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryQuadraticCurveTo(cpx float64, cpy float64, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DQuadraticCurveTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(cpx),
		float64(cpy),
		float64(x),
		float64(y),
	)

	return
}

// HasBezierCurveTo returns true if the method "PaintRenderingContext2D.bezierCurveTo" exists.
func (this PaintRenderingContext2D) HasBezierCurveTo() bool {
	return js.True == bindings.HasPaintRenderingContext2DBezierCurveTo(
		this.Ref(),
	)
}

// BezierCurveToFunc returns the method "PaintRenderingContext2D.bezierCurveTo".
func (this PaintRenderingContext2D) BezierCurveToFunc() (fn js.Func[func(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DBezierCurveToFunc(
			this.Ref(),
		),
	)
}

// BezierCurveTo calls the method "PaintRenderingContext2D.bezierCurveTo".
func (this PaintRenderingContext2D) BezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DBezierCurveTo(
		this.Ref(), js.Pointer(&ret),
		float64(cp1x),
		float64(cp1y),
		float64(cp2x),
		float64(cp2y),
		float64(x),
		float64(y),
	)

	return
}

// TryBezierCurveTo calls the method "PaintRenderingContext2D.bezierCurveTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryBezierCurveTo(cp1x float64, cp1y float64, cp2x float64, cp2y float64, x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DBezierCurveTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(cp1x),
		float64(cp1y),
		float64(cp2x),
		float64(cp2y),
		float64(x),
		float64(y),
	)

	return
}

// HasArcTo returns true if the method "PaintRenderingContext2D.arcTo" exists.
func (this PaintRenderingContext2D) HasArcTo() bool {
	return js.True == bindings.HasPaintRenderingContext2DArcTo(
		this.Ref(),
	)
}

// ArcToFunc returns the method "PaintRenderingContext2D.arcTo".
func (this PaintRenderingContext2D) ArcToFunc() (fn js.Func[func(x1 float64, y1 float64, x2 float64, y2 float64, radius float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DArcToFunc(
			this.Ref(),
		),
	)
}

// ArcTo calls the method "PaintRenderingContext2D.arcTo".
func (this PaintRenderingContext2D) ArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DArcTo(
		this.Ref(), js.Pointer(&ret),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	return
}

// TryArcTo calls the method "PaintRenderingContext2D.arcTo"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryArcTo(x1 float64, y1 float64, x2 float64, y2 float64, radius float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DArcTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x1),
		float64(y1),
		float64(x2),
		float64(y2),
		float64(radius),
	)

	return
}

// HasRect returns true if the method "PaintRenderingContext2D.rect" exists.
func (this PaintRenderingContext2D) HasRect() bool {
	return js.True == bindings.HasPaintRenderingContext2DRect(
		this.Ref(),
	)
}

// RectFunc returns the method "PaintRenderingContext2D.rect".
func (this PaintRenderingContext2D) RectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DRectFunc(
			this.Ref(),
		),
	)
}

// Rect calls the method "PaintRenderingContext2D.rect".
func (this PaintRenderingContext2D) Rect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryRect calls the method "PaintRenderingContext2D.rect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasRoundRect returns true if the method "PaintRenderingContext2D.roundRect" exists.
func (this PaintRenderingContext2D) HasRoundRect() bool {
	return js.True == bindings.HasPaintRenderingContext2DRoundRect(
		this.Ref(),
	)
}

// RoundRectFunc returns the method "PaintRenderingContext2D.roundRect".
func (this PaintRenderingContext2D) RoundRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DRoundRectFunc(
			this.Ref(),
		),
	)
}

// RoundRect calls the method "PaintRenderingContext2D.roundRect".
func (this PaintRenderingContext2D) RoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (ret js.Void) {
	bindings.CallPaintRenderingContext2DRoundRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	return
}

// TryRoundRect calls the method "PaintRenderingContext2D.roundRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryRoundRect(x float64, y float64, w float64, h float64, radii OneOf_Float64_DOMPointInit_ArrayOneOf_Float64_DOMPointInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DRoundRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
		radii.Ref(),
	)

	return
}

// HasRoundRect1 returns true if the method "PaintRenderingContext2D.roundRect" exists.
func (this PaintRenderingContext2D) HasRoundRect1() bool {
	return js.True == bindings.HasPaintRenderingContext2DRoundRect1(
		this.Ref(),
	)
}

// RoundRect1Func returns the method "PaintRenderingContext2D.roundRect".
func (this PaintRenderingContext2D) RoundRect1Func() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DRoundRect1Func(
			this.Ref(),
		),
	)
}

// RoundRect1 calls the method "PaintRenderingContext2D.roundRect".
func (this PaintRenderingContext2D) RoundRect1(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DRoundRect1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryRoundRect1 calls the method "PaintRenderingContext2D.roundRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryRoundRect1(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DRoundRect1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasArc returns true if the method "PaintRenderingContext2D.arc" exists.
func (this PaintRenderingContext2D) HasArc() bool {
	return js.True == bindings.HasPaintRenderingContext2DArc(
		this.Ref(),
	)
}

// ArcFunc returns the method "PaintRenderingContext2D.arc".
func (this PaintRenderingContext2D) ArcFunc() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DArcFunc(
			this.Ref(),
		),
	)
}

// Arc calls the method "PaintRenderingContext2D.arc".
func (this PaintRenderingContext2D) Arc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void) {
	bindings.CallPaintRenderingContext2DArc(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// TryArc calls the method "PaintRenderingContext2D.arc"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryArc(x float64, y float64, radius float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DArc(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// HasArc1 returns true if the method "PaintRenderingContext2D.arc" exists.
func (this PaintRenderingContext2D) HasArc1() bool {
	return js.True == bindings.HasPaintRenderingContext2DArc1(
		this.Ref(),
	)
}

// Arc1Func returns the method "PaintRenderingContext2D.arc".
func (this PaintRenderingContext2D) Arc1Func() (fn js.Func[func(x float64, y float64, radius float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DArc1Func(
			this.Ref(),
		),
	)
}

// Arc1 calls the method "PaintRenderingContext2D.arc".
func (this PaintRenderingContext2D) Arc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DArc1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// TryArc1 calls the method "PaintRenderingContext2D.arc"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryArc1(x float64, y float64, radius float64, startAngle float64, endAngle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DArc1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radius),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// HasEllipse returns true if the method "PaintRenderingContext2D.ellipse" exists.
func (this PaintRenderingContext2D) HasEllipse() bool {
	return js.True == bindings.HasPaintRenderingContext2DEllipse(
		this.Ref(),
	)
}

// EllipseFunc returns the method "PaintRenderingContext2D.ellipse".
func (this PaintRenderingContext2D) EllipseFunc() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DEllipseFunc(
			this.Ref(),
		),
	)
}

// Ellipse calls the method "PaintRenderingContext2D.ellipse".
func (this PaintRenderingContext2D) Ellipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void) {
	bindings.CallPaintRenderingContext2DEllipse(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// TryEllipse calls the method "PaintRenderingContext2D.ellipse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryEllipse(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64, counterclockwise bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DEllipse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
		js.Bool(bool(counterclockwise)),
	)

	return
}

// HasEllipse1 returns true if the method "PaintRenderingContext2D.ellipse" exists.
func (this PaintRenderingContext2D) HasEllipse1() bool {
	return js.True == bindings.HasPaintRenderingContext2DEllipse1(
		this.Ref(),
	)
}

// Ellipse1Func returns the method "PaintRenderingContext2D.ellipse".
func (this PaintRenderingContext2D) Ellipse1Func() (fn js.Func[func(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DEllipse1Func(
			this.Ref(),
		),
	)
}

// Ellipse1 calls the method "PaintRenderingContext2D.ellipse".
func (this PaintRenderingContext2D) Ellipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DEllipse1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// TryEllipse1 calls the method "PaintRenderingContext2D.ellipse"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryEllipse1(x float64, y float64, radiusX float64, radiusY float64, rotation float64, startAngle float64, endAngle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DEllipse1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(radiusX),
		float64(radiusY),
		float64(rotation),
		float64(startAngle),
		float64(endAngle),
	)

	return
}

// HasSetLineDash returns true if the method "PaintRenderingContext2D.setLineDash" exists.
func (this PaintRenderingContext2D) HasSetLineDash() bool {
	return js.True == bindings.HasPaintRenderingContext2DSetLineDash(
		this.Ref(),
	)
}

// SetLineDashFunc returns the method "PaintRenderingContext2D.setLineDash".
func (this PaintRenderingContext2D) SetLineDashFunc() (fn js.Func[func(segments js.Array[float64])]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DSetLineDashFunc(
			this.Ref(),
		),
	)
}

// SetLineDash calls the method "PaintRenderingContext2D.setLineDash".
func (this PaintRenderingContext2D) SetLineDash(segments js.Array[float64]) (ret js.Void) {
	bindings.CallPaintRenderingContext2DSetLineDash(
		this.Ref(), js.Pointer(&ret),
		segments.Ref(),
	)

	return
}

// TrySetLineDash calls the method "PaintRenderingContext2D.setLineDash"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TrySetLineDash(segments js.Array[float64]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DSetLineDash(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		segments.Ref(),
	)

	return
}

// HasGetLineDash returns true if the method "PaintRenderingContext2D.getLineDash" exists.
func (this PaintRenderingContext2D) HasGetLineDash() bool {
	return js.True == bindings.HasPaintRenderingContext2DGetLineDash(
		this.Ref(),
	)
}

// GetLineDashFunc returns the method "PaintRenderingContext2D.getLineDash".
func (this PaintRenderingContext2D) GetLineDashFunc() (fn js.Func[func() js.Array[float64]]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DGetLineDashFunc(
			this.Ref(),
		),
	)
}

// GetLineDash calls the method "PaintRenderingContext2D.getLineDash".
func (this PaintRenderingContext2D) GetLineDash() (ret js.Array[float64]) {
	bindings.CallPaintRenderingContext2DGetLineDash(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetLineDash calls the method "PaintRenderingContext2D.getLineDash"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryGetLineDash() (ret js.Array[float64], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DGetLineDash(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDrawImage returns true if the method "PaintRenderingContext2D.drawImage" exists.
func (this PaintRenderingContext2D) HasDrawImage() bool {
	return js.True == bindings.HasPaintRenderingContext2DDrawImage(
		this.Ref(),
	)
}

// DrawImageFunc returns the method "PaintRenderingContext2D.drawImage".
func (this PaintRenderingContext2D) DrawImageFunc() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DDrawImageFunc(
			this.Ref(),
		),
	)
}

// DrawImage calls the method "PaintRenderingContext2D.drawImage".
func (this PaintRenderingContext2D) DrawImage(image CanvasImageSource, dx float64, dy float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DDrawImage(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		float64(dx),
		float64(dy),
	)

	return
}

// TryDrawImage calls the method "PaintRenderingContext2D.drawImage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryDrawImage(image CanvasImageSource, dx float64, dy float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DDrawImage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		float64(dx),
		float64(dy),
	)

	return
}

// HasDrawImage1 returns true if the method "PaintRenderingContext2D.drawImage" exists.
func (this PaintRenderingContext2D) HasDrawImage1() bool {
	return js.True == bindings.HasPaintRenderingContext2DDrawImage1(
		this.Ref(),
	)
}

// DrawImage1Func returns the method "PaintRenderingContext2D.drawImage".
func (this PaintRenderingContext2D) DrawImage1Func() (fn js.Func[func(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DDrawImage1Func(
			this.Ref(),
		),
	)
}

// DrawImage1 calls the method "PaintRenderingContext2D.drawImage".
func (this PaintRenderingContext2D) DrawImage1(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DDrawImage1(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		float64(dx),
		float64(dy),
		float64(dw),
		float64(dh),
	)

	return
}

// TryDrawImage1 calls the method "PaintRenderingContext2D.drawImage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryDrawImage1(image CanvasImageSource, dx float64, dy float64, dw float64, dh float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DDrawImage1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		float64(dx),
		float64(dy),
		float64(dw),
		float64(dh),
	)

	return
}

// HasDrawImage2 returns true if the method "PaintRenderingContext2D.drawImage" exists.
func (this PaintRenderingContext2D) HasDrawImage2() bool {
	return js.True == bindings.HasPaintRenderingContext2DDrawImage2(
		this.Ref(),
	)
}

// DrawImage2Func returns the method "PaintRenderingContext2D.drawImage".
func (this PaintRenderingContext2D) DrawImage2Func() (fn js.Func[func(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DDrawImage2Func(
			this.Ref(),
		),
	)
}

// DrawImage2 calls the method "PaintRenderingContext2D.drawImage".
func (this PaintRenderingContext2D) DrawImage2(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DDrawImage2(
		this.Ref(), js.Pointer(&ret),
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

	return
}

// TryDrawImage2 calls the method "PaintRenderingContext2D.drawImage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryDrawImage2(image CanvasImageSource, sx float64, sy float64, sw float64, sh float64, dx float64, dy float64, dw float64, dh float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DDrawImage2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
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

	return
}

// HasBeginPath returns true if the method "PaintRenderingContext2D.beginPath" exists.
func (this PaintRenderingContext2D) HasBeginPath() bool {
	return js.True == bindings.HasPaintRenderingContext2DBeginPath(
		this.Ref(),
	)
}

// BeginPathFunc returns the method "PaintRenderingContext2D.beginPath".
func (this PaintRenderingContext2D) BeginPathFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DBeginPathFunc(
			this.Ref(),
		),
	)
}

// BeginPath calls the method "PaintRenderingContext2D.beginPath".
func (this PaintRenderingContext2D) BeginPath() (ret js.Void) {
	bindings.CallPaintRenderingContext2DBeginPath(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBeginPath calls the method "PaintRenderingContext2D.beginPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryBeginPath() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DBeginPath(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFill returns true if the method "PaintRenderingContext2D.fill" exists.
func (this PaintRenderingContext2D) HasFill() bool {
	return js.True == bindings.HasPaintRenderingContext2DFill(
		this.Ref(),
	)
}

// FillFunc returns the method "PaintRenderingContext2D.fill".
func (this PaintRenderingContext2D) FillFunc() (fn js.Func[func(fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DFillFunc(
			this.Ref(),
		),
	)
}

// Fill calls the method "PaintRenderingContext2D.fill".
func (this PaintRenderingContext2D) Fill(fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallPaintRenderingContext2DFill(
		this.Ref(), js.Pointer(&ret),
		uint32(fillRule),
	)

	return
}

// TryFill calls the method "PaintRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryFill(fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DFill(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(fillRule),
	)

	return
}

// HasFill1 returns true if the method "PaintRenderingContext2D.fill" exists.
func (this PaintRenderingContext2D) HasFill1() bool {
	return js.True == bindings.HasPaintRenderingContext2DFill1(
		this.Ref(),
	)
}

// Fill1Func returns the method "PaintRenderingContext2D.fill".
func (this PaintRenderingContext2D) Fill1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DFill1Func(
			this.Ref(),
		),
	)
}

// Fill1 calls the method "PaintRenderingContext2D.fill".
func (this PaintRenderingContext2D) Fill1() (ret js.Void) {
	bindings.CallPaintRenderingContext2DFill1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryFill1 calls the method "PaintRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryFill1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DFill1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFill2 returns true if the method "PaintRenderingContext2D.fill" exists.
func (this PaintRenderingContext2D) HasFill2() bool {
	return js.True == bindings.HasPaintRenderingContext2DFill2(
		this.Ref(),
	)
}

// Fill2Func returns the method "PaintRenderingContext2D.fill".
func (this PaintRenderingContext2D) Fill2Func() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DFill2Func(
			this.Ref(),
		),
	)
}

// Fill2 calls the method "PaintRenderingContext2D.fill".
func (this PaintRenderingContext2D) Fill2(path Path2D, fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallPaintRenderingContext2DFill2(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// TryFill2 calls the method "PaintRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryFill2(path Path2D, fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DFill2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// HasFill3 returns true if the method "PaintRenderingContext2D.fill" exists.
func (this PaintRenderingContext2D) HasFill3() bool {
	return js.True == bindings.HasPaintRenderingContext2DFill3(
		this.Ref(),
	)
}

// Fill3Func returns the method "PaintRenderingContext2D.fill".
func (this PaintRenderingContext2D) Fill3Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DFill3Func(
			this.Ref(),
		),
	)
}

// Fill3 calls the method "PaintRenderingContext2D.fill".
func (this PaintRenderingContext2D) Fill3(path Path2D) (ret js.Void) {
	bindings.CallPaintRenderingContext2DFill3(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryFill3 calls the method "PaintRenderingContext2D.fill"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryFill3(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DFill3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasStroke returns true if the method "PaintRenderingContext2D.stroke" exists.
func (this PaintRenderingContext2D) HasStroke() bool {
	return js.True == bindings.HasPaintRenderingContext2DStroke(
		this.Ref(),
	)
}

// StrokeFunc returns the method "PaintRenderingContext2D.stroke".
func (this PaintRenderingContext2D) StrokeFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DStrokeFunc(
			this.Ref(),
		),
	)
}

// Stroke calls the method "PaintRenderingContext2D.stroke".
func (this PaintRenderingContext2D) Stroke() (ret js.Void) {
	bindings.CallPaintRenderingContext2DStroke(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStroke calls the method "PaintRenderingContext2D.stroke"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryStroke() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DStroke(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStroke1 returns true if the method "PaintRenderingContext2D.stroke" exists.
func (this PaintRenderingContext2D) HasStroke1() bool {
	return js.True == bindings.HasPaintRenderingContext2DStroke1(
		this.Ref(),
	)
}

// Stroke1Func returns the method "PaintRenderingContext2D.stroke".
func (this PaintRenderingContext2D) Stroke1Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DStroke1Func(
			this.Ref(),
		),
	)
}

// Stroke1 calls the method "PaintRenderingContext2D.stroke".
func (this PaintRenderingContext2D) Stroke1(path Path2D) (ret js.Void) {
	bindings.CallPaintRenderingContext2DStroke1(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryStroke1 calls the method "PaintRenderingContext2D.stroke"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryStroke1(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DStroke1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasClip returns true if the method "PaintRenderingContext2D.clip" exists.
func (this PaintRenderingContext2D) HasClip() bool {
	return js.True == bindings.HasPaintRenderingContext2DClip(
		this.Ref(),
	)
}

// ClipFunc returns the method "PaintRenderingContext2D.clip".
func (this PaintRenderingContext2D) ClipFunc() (fn js.Func[func(fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClipFunc(
			this.Ref(),
		),
	)
}

// Clip calls the method "PaintRenderingContext2D.clip".
func (this PaintRenderingContext2D) Clip(fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallPaintRenderingContext2DClip(
		this.Ref(), js.Pointer(&ret),
		uint32(fillRule),
	)

	return
}

// TryClip calls the method "PaintRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryClip(fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DClip(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(fillRule),
	)

	return
}

// HasClip1 returns true if the method "PaintRenderingContext2D.clip" exists.
func (this PaintRenderingContext2D) HasClip1() bool {
	return js.True == bindings.HasPaintRenderingContext2DClip1(
		this.Ref(),
	)
}

// Clip1Func returns the method "PaintRenderingContext2D.clip".
func (this PaintRenderingContext2D) Clip1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClip1Func(
			this.Ref(),
		),
	)
}

// Clip1 calls the method "PaintRenderingContext2D.clip".
func (this PaintRenderingContext2D) Clip1() (ret js.Void) {
	bindings.CallPaintRenderingContext2DClip1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClip1 calls the method "PaintRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryClip1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DClip1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClip2 returns true if the method "PaintRenderingContext2D.clip" exists.
func (this PaintRenderingContext2D) HasClip2() bool {
	return js.True == bindings.HasPaintRenderingContext2DClip2(
		this.Ref(),
	)
}

// Clip2Func returns the method "PaintRenderingContext2D.clip".
func (this PaintRenderingContext2D) Clip2Func() (fn js.Func[func(path Path2D, fillRule CanvasFillRule)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClip2Func(
			this.Ref(),
		),
	)
}

// Clip2 calls the method "PaintRenderingContext2D.clip".
func (this PaintRenderingContext2D) Clip2(path Path2D, fillRule CanvasFillRule) (ret js.Void) {
	bindings.CallPaintRenderingContext2DClip2(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// TryClip2 calls the method "PaintRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryClip2(path Path2D, fillRule CanvasFillRule) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DClip2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		uint32(fillRule),
	)

	return
}

// HasClip3 returns true if the method "PaintRenderingContext2D.clip" exists.
func (this PaintRenderingContext2D) HasClip3() bool {
	return js.True == bindings.HasPaintRenderingContext2DClip3(
		this.Ref(),
	)
}

// Clip3Func returns the method "PaintRenderingContext2D.clip".
func (this PaintRenderingContext2D) Clip3Func() (fn js.Func[func(path Path2D)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClip3Func(
			this.Ref(),
		),
	)
}

// Clip3 calls the method "PaintRenderingContext2D.clip".
func (this PaintRenderingContext2D) Clip3(path Path2D) (ret js.Void) {
	bindings.CallPaintRenderingContext2DClip3(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
	)

	return
}

// TryClip3 calls the method "PaintRenderingContext2D.clip"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryClip3(path Path2D) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DClip3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
	)

	return
}

// HasIsPointInPath returns true if the method "PaintRenderingContext2D.isPointInPath" exists.
func (this PaintRenderingContext2D) HasIsPointInPath() bool {
	return js.True == bindings.HasPaintRenderingContext2DIsPointInPath(
		this.Ref(),
	)
}

// IsPointInPathFunc returns the method "PaintRenderingContext2D.isPointInPath".
func (this PaintRenderingContext2D) IsPointInPathFunc() (fn js.Func[func(x float64, y float64, fillRule CanvasFillRule) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInPathFunc(
			this.Ref(),
		),
	)
}

// IsPointInPath calls the method "PaintRenderingContext2D.isPointInPath".
func (this PaintRenderingContext2D) IsPointInPath(x float64, y float64, fillRule CanvasFillRule) (ret bool) {
	bindings.CallPaintRenderingContext2DIsPointInPath(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// TryIsPointInPath calls the method "PaintRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryIsPointInPath(x float64, y float64, fillRule CanvasFillRule) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DIsPointInPath(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// HasIsPointInPath1 returns true if the method "PaintRenderingContext2D.isPointInPath" exists.
func (this PaintRenderingContext2D) HasIsPointInPath1() bool {
	return js.True == bindings.HasPaintRenderingContext2DIsPointInPath1(
		this.Ref(),
	)
}

// IsPointInPath1Func returns the method "PaintRenderingContext2D.isPointInPath".
func (this PaintRenderingContext2D) IsPointInPath1Func() (fn js.Func[func(x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInPath1Func(
			this.Ref(),
		),
	)
}

// IsPointInPath1 calls the method "PaintRenderingContext2D.isPointInPath".
func (this PaintRenderingContext2D) IsPointInPath1(x float64, y float64) (ret bool) {
	bindings.CallPaintRenderingContext2DIsPointInPath1(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInPath1 calls the method "PaintRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryIsPointInPath1(x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DIsPointInPath1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasIsPointInPath2 returns true if the method "PaintRenderingContext2D.isPointInPath" exists.
func (this PaintRenderingContext2D) HasIsPointInPath2() bool {
	return js.True == bindings.HasPaintRenderingContext2DIsPointInPath2(
		this.Ref(),
	)
}

// IsPointInPath2Func returns the method "PaintRenderingContext2D.isPointInPath".
func (this PaintRenderingContext2D) IsPointInPath2Func() (fn js.Func[func(path Path2D, x float64, y float64, fillRule CanvasFillRule) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInPath2Func(
			this.Ref(),
		),
	)
}

// IsPointInPath2 calls the method "PaintRenderingContext2D.isPointInPath".
func (this PaintRenderingContext2D) IsPointInPath2(path Path2D, x float64, y float64, fillRule CanvasFillRule) (ret bool) {
	bindings.CallPaintRenderingContext2DIsPointInPath2(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// TryIsPointInPath2 calls the method "PaintRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryIsPointInPath2(path Path2D, x float64, y float64, fillRule CanvasFillRule) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DIsPointInPath2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		float64(x),
		float64(y),
		uint32(fillRule),
	)

	return
}

// HasIsPointInPath3 returns true if the method "PaintRenderingContext2D.isPointInPath" exists.
func (this PaintRenderingContext2D) HasIsPointInPath3() bool {
	return js.True == bindings.HasPaintRenderingContext2DIsPointInPath3(
		this.Ref(),
	)
}

// IsPointInPath3Func returns the method "PaintRenderingContext2D.isPointInPath".
func (this PaintRenderingContext2D) IsPointInPath3Func() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInPath3Func(
			this.Ref(),
		),
	)
}

// IsPointInPath3 calls the method "PaintRenderingContext2D.isPointInPath".
func (this PaintRenderingContext2D) IsPointInPath3(path Path2D, x float64, y float64) (ret bool) {
	bindings.CallPaintRenderingContext2DIsPointInPath3(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInPath3 calls the method "PaintRenderingContext2D.isPointInPath"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryIsPointInPath3(path Path2D, x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DIsPointInPath3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// HasIsPointInStroke returns true if the method "PaintRenderingContext2D.isPointInStroke" exists.
func (this PaintRenderingContext2D) HasIsPointInStroke() bool {
	return js.True == bindings.HasPaintRenderingContext2DIsPointInStroke(
		this.Ref(),
	)
}

// IsPointInStrokeFunc returns the method "PaintRenderingContext2D.isPointInStroke".
func (this PaintRenderingContext2D) IsPointInStrokeFunc() (fn js.Func[func(x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInStrokeFunc(
			this.Ref(),
		),
	)
}

// IsPointInStroke calls the method "PaintRenderingContext2D.isPointInStroke".
func (this PaintRenderingContext2D) IsPointInStroke(x float64, y float64) (ret bool) {
	bindings.CallPaintRenderingContext2DIsPointInStroke(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInStroke calls the method "PaintRenderingContext2D.isPointInStroke"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryIsPointInStroke(x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DIsPointInStroke(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasIsPointInStroke1 returns true if the method "PaintRenderingContext2D.isPointInStroke" exists.
func (this PaintRenderingContext2D) HasIsPointInStroke1() bool {
	return js.True == bindings.HasPaintRenderingContext2DIsPointInStroke1(
		this.Ref(),
	)
}

// IsPointInStroke1Func returns the method "PaintRenderingContext2D.isPointInStroke".
func (this PaintRenderingContext2D) IsPointInStroke1Func() (fn js.Func[func(path Path2D, x float64, y float64) bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsPointInStroke1Func(
			this.Ref(),
		),
	)
}

// IsPointInStroke1 calls the method "PaintRenderingContext2D.isPointInStroke".
func (this PaintRenderingContext2D) IsPointInStroke1(path Path2D, x float64, y float64) (ret bool) {
	bindings.CallPaintRenderingContext2DIsPointInStroke1(
		this.Ref(), js.Pointer(&ret),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// TryIsPointInStroke1 calls the method "PaintRenderingContext2D.isPointInStroke"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryIsPointInStroke1(path Path2D, x float64, y float64) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DIsPointInStroke1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		path.Ref(),
		float64(x),
		float64(y),
	)

	return
}

// HasClearRect returns true if the method "PaintRenderingContext2D.clearRect" exists.
func (this PaintRenderingContext2D) HasClearRect() bool {
	return js.True == bindings.HasPaintRenderingContext2DClearRect(
		this.Ref(),
	)
}

// ClearRectFunc returns the method "PaintRenderingContext2D.clearRect".
func (this PaintRenderingContext2D) ClearRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DClearRectFunc(
			this.Ref(),
		),
	)
}

// ClearRect calls the method "PaintRenderingContext2D.clearRect".
func (this PaintRenderingContext2D) ClearRect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DClearRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryClearRect calls the method "PaintRenderingContext2D.clearRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryClearRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DClearRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasFillRect returns true if the method "PaintRenderingContext2D.fillRect" exists.
func (this PaintRenderingContext2D) HasFillRect() bool {
	return js.True == bindings.HasPaintRenderingContext2DFillRect(
		this.Ref(),
	)
}

// FillRectFunc returns the method "PaintRenderingContext2D.fillRect".
func (this PaintRenderingContext2D) FillRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DFillRectFunc(
			this.Ref(),
		),
	)
}

// FillRect calls the method "PaintRenderingContext2D.fillRect".
func (this PaintRenderingContext2D) FillRect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DFillRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryFillRect calls the method "PaintRenderingContext2D.fillRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryFillRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DFillRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasStrokeRect returns true if the method "PaintRenderingContext2D.strokeRect" exists.
func (this PaintRenderingContext2D) HasStrokeRect() bool {
	return js.True == bindings.HasPaintRenderingContext2DStrokeRect(
		this.Ref(),
	)
}

// StrokeRectFunc returns the method "PaintRenderingContext2D.strokeRect".
func (this PaintRenderingContext2D) StrokeRectFunc() (fn js.Func[func(x float64, y float64, w float64, h float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DStrokeRectFunc(
			this.Ref(),
		),
	)
}

// StrokeRect calls the method "PaintRenderingContext2D.strokeRect".
func (this PaintRenderingContext2D) StrokeRect(x float64, y float64, w float64, h float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DStrokeRect(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// TryStrokeRect calls the method "PaintRenderingContext2D.strokeRect"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryStrokeRect(x float64, y float64, w float64, h float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DStrokeRect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
		float64(w),
		float64(h),
	)

	return
}

// HasCreateLinearGradient returns true if the method "PaintRenderingContext2D.createLinearGradient" exists.
func (this PaintRenderingContext2D) HasCreateLinearGradient() bool {
	return js.True == bindings.HasPaintRenderingContext2DCreateLinearGradient(
		this.Ref(),
	)
}

// CreateLinearGradientFunc returns the method "PaintRenderingContext2D.createLinearGradient".
func (this PaintRenderingContext2D) CreateLinearGradientFunc() (fn js.Func[func(x0 float64, y0 float64, x1 float64, y1 float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DCreateLinearGradientFunc(
			this.Ref(),
		),
	)
}

// CreateLinearGradient calls the method "PaintRenderingContext2D.createLinearGradient".
func (this PaintRenderingContext2D) CreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) (ret CanvasGradient) {
	bindings.CallPaintRenderingContext2DCreateLinearGradient(
		this.Ref(), js.Pointer(&ret),
		float64(x0),
		float64(y0),
		float64(x1),
		float64(y1),
	)

	return
}

// TryCreateLinearGradient calls the method "PaintRenderingContext2D.createLinearGradient"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryCreateLinearGradient(x0 float64, y0 float64, x1 float64, y1 float64) (ret CanvasGradient, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DCreateLinearGradient(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x0),
		float64(y0),
		float64(x1),
		float64(y1),
	)

	return
}

// HasCreateRadialGradient returns true if the method "PaintRenderingContext2D.createRadialGradient" exists.
func (this PaintRenderingContext2D) HasCreateRadialGradient() bool {
	return js.True == bindings.HasPaintRenderingContext2DCreateRadialGradient(
		this.Ref(),
	)
}

// CreateRadialGradientFunc returns the method "PaintRenderingContext2D.createRadialGradient".
func (this PaintRenderingContext2D) CreateRadialGradientFunc() (fn js.Func[func(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DCreateRadialGradientFunc(
			this.Ref(),
		),
	)
}

// CreateRadialGradient calls the method "PaintRenderingContext2D.createRadialGradient".
func (this PaintRenderingContext2D) CreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) (ret CanvasGradient) {
	bindings.CallPaintRenderingContext2DCreateRadialGradient(
		this.Ref(), js.Pointer(&ret),
		float64(x0),
		float64(y0),
		float64(r0),
		float64(x1),
		float64(y1),
		float64(r1),
	)

	return
}

// TryCreateRadialGradient calls the method "PaintRenderingContext2D.createRadialGradient"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryCreateRadialGradient(x0 float64, y0 float64, r0 float64, x1 float64, y1 float64, r1 float64) (ret CanvasGradient, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DCreateRadialGradient(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x0),
		float64(y0),
		float64(r0),
		float64(x1),
		float64(y1),
		float64(r1),
	)

	return
}

// HasCreateConicGradient returns true if the method "PaintRenderingContext2D.createConicGradient" exists.
func (this PaintRenderingContext2D) HasCreateConicGradient() bool {
	return js.True == bindings.HasPaintRenderingContext2DCreateConicGradient(
		this.Ref(),
	)
}

// CreateConicGradientFunc returns the method "PaintRenderingContext2D.createConicGradient".
func (this PaintRenderingContext2D) CreateConicGradientFunc() (fn js.Func[func(startAngle float64, x float64, y float64) CanvasGradient]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DCreateConicGradientFunc(
			this.Ref(),
		),
	)
}

// CreateConicGradient calls the method "PaintRenderingContext2D.createConicGradient".
func (this PaintRenderingContext2D) CreateConicGradient(startAngle float64, x float64, y float64) (ret CanvasGradient) {
	bindings.CallPaintRenderingContext2DCreateConicGradient(
		this.Ref(), js.Pointer(&ret),
		float64(startAngle),
		float64(x),
		float64(y),
	)

	return
}

// TryCreateConicGradient calls the method "PaintRenderingContext2D.createConicGradient"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryCreateConicGradient(startAngle float64, x float64, y float64) (ret CanvasGradient, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DCreateConicGradient(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(startAngle),
		float64(x),
		float64(y),
	)

	return
}

// HasCreatePattern returns true if the method "PaintRenderingContext2D.createPattern" exists.
func (this PaintRenderingContext2D) HasCreatePattern() bool {
	return js.True == bindings.HasPaintRenderingContext2DCreatePattern(
		this.Ref(),
	)
}

// CreatePatternFunc returns the method "PaintRenderingContext2D.createPattern".
func (this PaintRenderingContext2D) CreatePatternFunc() (fn js.Func[func(image CanvasImageSource, repetition js.String) CanvasPattern]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DCreatePatternFunc(
			this.Ref(),
		),
	)
}

// CreatePattern calls the method "PaintRenderingContext2D.createPattern".
func (this PaintRenderingContext2D) CreatePattern(image CanvasImageSource, repetition js.String) (ret CanvasPattern) {
	bindings.CallPaintRenderingContext2DCreatePattern(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
		repetition.Ref(),
	)

	return
}

// TryCreatePattern calls the method "PaintRenderingContext2D.createPattern"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryCreatePattern(image CanvasImageSource, repetition js.String) (ret CanvasPattern, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DCreatePattern(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
		repetition.Ref(),
	)

	return
}

// HasScale returns true if the method "PaintRenderingContext2D.scale" exists.
func (this PaintRenderingContext2D) HasScale() bool {
	return js.True == bindings.HasPaintRenderingContext2DScale(
		this.Ref(),
	)
}

// ScaleFunc returns the method "PaintRenderingContext2D.scale".
func (this PaintRenderingContext2D) ScaleFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DScaleFunc(
			this.Ref(),
		),
	)
}

// Scale calls the method "PaintRenderingContext2D.scale".
func (this PaintRenderingContext2D) Scale(x float64, y float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DScale(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryScale calls the method "PaintRenderingContext2D.scale"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryScale(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DScale(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasRotate returns true if the method "PaintRenderingContext2D.rotate" exists.
func (this PaintRenderingContext2D) HasRotate() bool {
	return js.True == bindings.HasPaintRenderingContext2DRotate(
		this.Ref(),
	)
}

// RotateFunc returns the method "PaintRenderingContext2D.rotate".
func (this PaintRenderingContext2D) RotateFunc() (fn js.Func[func(angle float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DRotateFunc(
			this.Ref(),
		),
	)
}

// Rotate calls the method "PaintRenderingContext2D.rotate".
func (this PaintRenderingContext2D) Rotate(angle float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DRotate(
		this.Ref(), js.Pointer(&ret),
		float64(angle),
	)

	return
}

// TryRotate calls the method "PaintRenderingContext2D.rotate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryRotate(angle float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DRotate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(angle),
	)

	return
}

// HasTranslate returns true if the method "PaintRenderingContext2D.translate" exists.
func (this PaintRenderingContext2D) HasTranslate() bool {
	return js.True == bindings.HasPaintRenderingContext2DTranslate(
		this.Ref(),
	)
}

// TranslateFunc returns the method "PaintRenderingContext2D.translate".
func (this PaintRenderingContext2D) TranslateFunc() (fn js.Func[func(x float64, y float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DTranslateFunc(
			this.Ref(),
		),
	)
}

// Translate calls the method "PaintRenderingContext2D.translate".
func (this PaintRenderingContext2D) Translate(x float64, y float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DTranslate(
		this.Ref(), js.Pointer(&ret),
		float64(x),
		float64(y),
	)

	return
}

// TryTranslate calls the method "PaintRenderingContext2D.translate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryTranslate(x float64, y float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DTranslate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(x),
		float64(y),
	)

	return
}

// HasTransform returns true if the method "PaintRenderingContext2D.transform" exists.
func (this PaintRenderingContext2D) HasTransform() bool {
	return js.True == bindings.HasPaintRenderingContext2DTransform(
		this.Ref(),
	)
}

// TransformFunc returns the method "PaintRenderingContext2D.transform".
func (this PaintRenderingContext2D) TransformFunc() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DTransformFunc(
			this.Ref(),
		),
	)
}

// Transform calls the method "PaintRenderingContext2D.transform".
func (this PaintRenderingContext2D) Transform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DTransform(
		this.Ref(), js.Pointer(&ret),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	return
}

// TryTransform calls the method "PaintRenderingContext2D.transform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryTransform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DTransform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	return
}

// HasGetTransform returns true if the method "PaintRenderingContext2D.getTransform" exists.
func (this PaintRenderingContext2D) HasGetTransform() bool {
	return js.True == bindings.HasPaintRenderingContext2DGetTransform(
		this.Ref(),
	)
}

// GetTransformFunc returns the method "PaintRenderingContext2D.getTransform".
func (this PaintRenderingContext2D) GetTransformFunc() (fn js.Func[func() DOMMatrix]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DGetTransformFunc(
			this.Ref(),
		),
	)
}

// GetTransform calls the method "PaintRenderingContext2D.getTransform".
func (this PaintRenderingContext2D) GetTransform() (ret DOMMatrix) {
	bindings.CallPaintRenderingContext2DGetTransform(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetTransform calls the method "PaintRenderingContext2D.getTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryGetTransform() (ret DOMMatrix, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DGetTransform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetTransform returns true if the method "PaintRenderingContext2D.setTransform" exists.
func (this PaintRenderingContext2D) HasSetTransform() bool {
	return js.True == bindings.HasPaintRenderingContext2DSetTransform(
		this.Ref(),
	)
}

// SetTransformFunc returns the method "PaintRenderingContext2D.setTransform".
func (this PaintRenderingContext2D) SetTransformFunc() (fn js.Func[func(a float64, b float64, c float64, d float64, e float64, f float64)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DSetTransformFunc(
			this.Ref(),
		),
	)
}

// SetTransform calls the method "PaintRenderingContext2D.setTransform".
func (this PaintRenderingContext2D) SetTransform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void) {
	bindings.CallPaintRenderingContext2DSetTransform(
		this.Ref(), js.Pointer(&ret),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	return
}

// TrySetTransform calls the method "PaintRenderingContext2D.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TrySetTransform(a float64, b float64, c float64, d float64, e float64, f float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DSetTransform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(a),
		float64(b),
		float64(c),
		float64(d),
		float64(e),
		float64(f),
	)

	return
}

// HasSetTransform1 returns true if the method "PaintRenderingContext2D.setTransform" exists.
func (this PaintRenderingContext2D) HasSetTransform1() bool {
	return js.True == bindings.HasPaintRenderingContext2DSetTransform1(
		this.Ref(),
	)
}

// SetTransform1Func returns the method "PaintRenderingContext2D.setTransform".
func (this PaintRenderingContext2D) SetTransform1Func() (fn js.Func[func(transform DOMMatrix2DInit)]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DSetTransform1Func(
			this.Ref(),
		),
	)
}

// SetTransform1 calls the method "PaintRenderingContext2D.setTransform".
func (this PaintRenderingContext2D) SetTransform1(transform DOMMatrix2DInit) (ret js.Void) {
	bindings.CallPaintRenderingContext2DSetTransform1(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&transform),
	)

	return
}

// TrySetTransform1 calls the method "PaintRenderingContext2D.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TrySetTransform1(transform DOMMatrix2DInit) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DSetTransform1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&transform),
	)

	return
}

// HasSetTransform2 returns true if the method "PaintRenderingContext2D.setTransform" exists.
func (this PaintRenderingContext2D) HasSetTransform2() bool {
	return js.True == bindings.HasPaintRenderingContext2DSetTransform2(
		this.Ref(),
	)
}

// SetTransform2Func returns the method "PaintRenderingContext2D.setTransform".
func (this PaintRenderingContext2D) SetTransform2Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DSetTransform2Func(
			this.Ref(),
		),
	)
}

// SetTransform2 calls the method "PaintRenderingContext2D.setTransform".
func (this PaintRenderingContext2D) SetTransform2() (ret js.Void) {
	bindings.CallPaintRenderingContext2DSetTransform2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetTransform2 calls the method "PaintRenderingContext2D.setTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TrySetTransform2() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DSetTransform2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasResetTransform returns true if the method "PaintRenderingContext2D.resetTransform" exists.
func (this PaintRenderingContext2D) HasResetTransform() bool {
	return js.True == bindings.HasPaintRenderingContext2DResetTransform(
		this.Ref(),
	)
}

// ResetTransformFunc returns the method "PaintRenderingContext2D.resetTransform".
func (this PaintRenderingContext2D) ResetTransformFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DResetTransformFunc(
			this.Ref(),
		),
	)
}

// ResetTransform calls the method "PaintRenderingContext2D.resetTransform".
func (this PaintRenderingContext2D) ResetTransform() (ret js.Void) {
	bindings.CallPaintRenderingContext2DResetTransform(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryResetTransform calls the method "PaintRenderingContext2D.resetTransform"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryResetTransform() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DResetTransform(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSave returns true if the method "PaintRenderingContext2D.save" exists.
func (this PaintRenderingContext2D) HasSave() bool {
	return js.True == bindings.HasPaintRenderingContext2DSave(
		this.Ref(),
	)
}

// SaveFunc returns the method "PaintRenderingContext2D.save".
func (this PaintRenderingContext2D) SaveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DSaveFunc(
			this.Ref(),
		),
	)
}

// Save calls the method "PaintRenderingContext2D.save".
func (this PaintRenderingContext2D) Save() (ret js.Void) {
	bindings.CallPaintRenderingContext2DSave(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySave calls the method "PaintRenderingContext2D.save"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TrySave() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DSave(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRestore returns true if the method "PaintRenderingContext2D.restore" exists.
func (this PaintRenderingContext2D) HasRestore() bool {
	return js.True == bindings.HasPaintRenderingContext2DRestore(
		this.Ref(),
	)
}

// RestoreFunc returns the method "PaintRenderingContext2D.restore".
func (this PaintRenderingContext2D) RestoreFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DRestoreFunc(
			this.Ref(),
		),
	)
}

// Restore calls the method "PaintRenderingContext2D.restore".
func (this PaintRenderingContext2D) Restore() (ret js.Void) {
	bindings.CallPaintRenderingContext2DRestore(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRestore calls the method "PaintRenderingContext2D.restore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryRestore() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DRestore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReset returns true if the method "PaintRenderingContext2D.reset" exists.
func (this PaintRenderingContext2D) HasReset() bool {
	return js.True == bindings.HasPaintRenderingContext2DReset(
		this.Ref(),
	)
}

// ResetFunc returns the method "PaintRenderingContext2D.reset".
func (this PaintRenderingContext2D) ResetFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DResetFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "PaintRenderingContext2D.reset".
func (this PaintRenderingContext2D) Reset() (ret js.Void) {
	bindings.CallPaintRenderingContext2DReset(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "PaintRenderingContext2D.reset"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryReset() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DReset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIsContextLost returns true if the method "PaintRenderingContext2D.isContextLost" exists.
func (this PaintRenderingContext2D) HasIsContextLost() bool {
	return js.True == bindings.HasPaintRenderingContext2DIsContextLost(
		this.Ref(),
	)
}

// IsContextLostFunc returns the method "PaintRenderingContext2D.isContextLost".
func (this PaintRenderingContext2D) IsContextLostFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.PaintRenderingContext2DIsContextLostFunc(
			this.Ref(),
		),
	)
}

// IsContextLost calls the method "PaintRenderingContext2D.isContextLost".
func (this PaintRenderingContext2D) IsContextLost() (ret bool) {
	bindings.CallPaintRenderingContext2DIsContextLost(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryIsContextLost calls the method "PaintRenderingContext2D.isContextLost"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintRenderingContext2D) TryIsContextLost() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintRenderingContext2DIsContextLost(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type PaintRenderingContext2DSettings struct {
	// Alpha is "PaintRenderingContext2DSettings.alpha"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_Alpha MUST be set to true to make this field effective.
	Alpha bool

	FFI_USE_Alpha bool // for Alpha.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaintRenderingContext2DSettings with all fields set.
func (p PaintRenderingContext2DSettings) FromRef(ref js.Ref) PaintRenderingContext2DSettings {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaintRenderingContext2DSettings in the application heap.
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
// It returns ok=false if there is no such property.
func (this PaintSize) Width() (ret float64, ok bool) {
	ok = js.True == bindings.GetPaintSizeWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Height returns the value of property "PaintSize.height".
//
// It returns ok=false if there is no such property.
func (this PaintSize) Height() (ret float64, ok bool) {
	ok = js.True == bindings.GetPaintSizeHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this PaintWorkletGlobalScope) DevicePixelRatio() (ret float64, ok bool) {
	ok = js.True == bindings.GetPaintWorkletGlobalScopeDevicePixelRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRegisterPaint returns true if the method "PaintWorkletGlobalScope.registerPaint" exists.
func (this PaintWorkletGlobalScope) HasRegisterPaint() bool {
	return js.True == bindings.HasPaintWorkletGlobalScopeRegisterPaint(
		this.Ref(),
	)
}

// RegisterPaintFunc returns the method "PaintWorkletGlobalScope.registerPaint".
func (this PaintWorkletGlobalScope) RegisterPaintFunc() (fn js.Func[func(name js.String, paintCtor js.Func[func()])]) {
	return fn.FromRef(
		bindings.PaintWorkletGlobalScopeRegisterPaintFunc(
			this.Ref(),
		),
	)
}

// RegisterPaint calls the method "PaintWorkletGlobalScope.registerPaint".
func (this PaintWorkletGlobalScope) RegisterPaint(name js.String, paintCtor js.Func[func()]) (ret js.Void) {
	bindings.CallPaintWorkletGlobalScopeRegisterPaint(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		paintCtor.Ref(),
	)

	return
}

// TryRegisterPaint calls the method "PaintWorkletGlobalScope.registerPaint"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this PaintWorkletGlobalScope) TryRegisterPaint(name js.String, paintCtor js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaintWorkletGlobalScopeRegisterPaint(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		paintCtor.Ref(),
	)

	return
}

func NewPasswordCredential(form HTMLFormElement) (ret PasswordCredential) {
	ret.ref = bindings.NewPasswordCredentialByPasswordCredential(
		form.Ref())
	return
}

func NewPasswordCredentialByPasswordCredential1(data PasswordCredentialData) (ret PasswordCredential) {
	ret.ref = bindings.NewPasswordCredentialByPasswordCredential1(
		js.Pointer(&data))
	return
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
// It returns ok=false if there is no such property.
func (this PasswordCredential) Password() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPasswordCredentialPassword(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "PasswordCredential.name".
//
// It returns ok=false if there is no such property.
func (this PasswordCredential) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPasswordCredentialName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IconURL returns the value of property "PasswordCredential.iconURL".
//
// It returns ok=false if there is no such property.
func (this PasswordCredential) IconURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPasswordCredentialIconURL(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// New creates a new PaymentCompleteDetails in the application heap.
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
	//
	// NOTE: Amount.FFI_USE MUST be set to true to get Amount used.
	Amount PaymentCurrencyAmount
	// Pending is "PaymentItem.pending"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Pending MUST be set to true to make this field effective.
	Pending bool

	FFI_USE_Pending bool // for Pending.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PaymentItem with all fields set.
func (p PaymentItem) FromRef(ref js.Ref) PaymentItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PaymentItem in the application heap.
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

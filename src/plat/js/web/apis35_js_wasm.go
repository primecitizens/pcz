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

type EncodedVideoChunkMetadata struct {
	// DecoderConfig is "EncodedVideoChunkMetadata.decoderConfig"
	//
	// Optional
	DecoderConfig VideoDecoderConfig
	// Svc is "EncodedVideoChunkMetadata.svc"
	//
	// Optional
	Svc SvcOutputMetadata
	// AlphaSideData is "EncodedVideoChunkMetadata.alphaSideData"
	//
	// Optional
	AlphaSideData BufferSource

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EncodedVideoChunkMetadata with all fields set.
func (p EncodedVideoChunkMetadata) FromRef(ref js.Ref) EncodedVideoChunkMetadata {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EncodedVideoChunkMetadata in the application heap.
func (p EncodedVideoChunkMetadata) New() js.Ref {
	return bindings.EncodedVideoChunkMetadataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p EncodedVideoChunkMetadata) UpdateFrom(ref js.Ref) {
	bindings.EncodedVideoChunkMetadataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p EncodedVideoChunkMetadata) Update(ref js.Ref) {
	bindings.EncodedVideoChunkMetadataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type EncodedVideoChunkOutputCallbackFunc func(this js.Ref, chunk EncodedVideoChunk, metadata EncodedVideoChunkMetadata) js.Ref

func (fn EncodedVideoChunkOutputCallbackFunc) Register() js.Func[func(chunk EncodedVideoChunk, metadata EncodedVideoChunkMetadata)] {
	return js.RegisterCallback[func(chunk EncodedVideoChunk, metadata EncodedVideoChunkMetadata)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EncodedVideoChunkOutputCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		EncodedVideoChunk{}.FromRef(args[0+1]),
		EncodedVideoChunkMetadata{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EncodedVideoChunkOutputCallback[T any] struct {
	Fn  func(arg T, this js.Ref, chunk EncodedVideoChunk, metadata EncodedVideoChunkMetadata) js.Ref
	Arg T
}

func (cb *EncodedVideoChunkOutputCallback[T]) Register() js.Func[func(chunk EncodedVideoChunk, metadata EncodedVideoChunkMetadata)] {
	return js.RegisterCallback[func(chunk EncodedVideoChunk, metadata EncodedVideoChunkMetadata)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EncodedVideoChunkOutputCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		EncodedVideoChunk{}.FromRef(args[0+1]),
		EncodedVideoChunkMetadata{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ErrorEventInit struct {
	// Message is "ErrorEventInit.message"
	//
	// Optional, defaults to "".
	Message js.String
	// Filename is "ErrorEventInit.filename"
	//
	// Optional, defaults to "".
	Filename js.String
	// Lineno is "ErrorEventInit.lineno"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Lineno MUST be set to true to make this field effective.
	Lineno uint32
	// Colno is "ErrorEventInit.colno"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Colno MUST be set to true to make this field effective.
	Colno uint32
	// Error is "ErrorEventInit.error"
	//
	// Optional
	Error js.Any
	// Bubbles is "ErrorEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "ErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "ErrorEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Lineno     bool // for Lineno.
	FFI_USE_Colno      bool // for Colno.
	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ErrorEventInit with all fields set.
func (p ErrorEventInit) FromRef(ref js.Ref) ErrorEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ErrorEventInit in the application heap.
func (p ErrorEventInit) New() js.Ref {
	return bindings.ErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.ErrorEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ErrorEventInit) Update(ref js.Ref) {
	bindings.ErrorEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewErrorEvent(typ js.String, eventInitDict ErrorEventInit) ErrorEvent {
	return ErrorEvent{}.FromRef(
		bindings.NewErrorEventByErrorEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewErrorEventByErrorEvent1(typ js.String) ErrorEvent {
	return ErrorEvent{}.FromRef(
		bindings.NewErrorEventByErrorEvent1(
			typ.Ref()),
	)
}

type ErrorEvent struct {
	Event
}

func (this ErrorEvent) Once() ErrorEvent {
	this.Ref().Once()
	return this
}

func (this ErrorEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this ErrorEvent) FromRef(ref js.Ref) ErrorEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this ErrorEvent) Free() {
	this.Ref().Free()
}

// Message returns the value of property "ErrorEvent.message".
//
// The returned bool will be false if there is no such property.
func (this ErrorEvent) Message() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetErrorEventMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Filename returns the value of property "ErrorEvent.filename".
//
// The returned bool will be false if there is no such property.
func (this ErrorEvent) Filename() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetErrorEventFilename(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Lineno returns the value of property "ErrorEvent.lineno".
//
// The returned bool will be false if there is no such property.
func (this ErrorEvent) Lineno() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetErrorEventLineno(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Colno returns the value of property "ErrorEvent.colno".
//
// The returned bool will be false if there is no such property.
func (this ErrorEvent) Colno() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetErrorEventColno(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Error returns the value of property "ErrorEvent.error".
//
// The returned bool will be false if there is no such property.
func (this ErrorEvent) Error() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetErrorEventError(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

type EventModifierInit struct {
	// CtrlKey is "EventModifierInit.ctrlKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_CtrlKey MUST be set to true to make this field effective.
	CtrlKey bool
	// ShiftKey is "EventModifierInit.shiftKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ShiftKey MUST be set to true to make this field effective.
	ShiftKey bool
	// AltKey is "EventModifierInit.altKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AltKey MUST be set to true to make this field effective.
	AltKey bool
	// MetaKey is "EventModifierInit.metaKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_MetaKey MUST be set to true to make this field effective.
	MetaKey bool
	// ModifierAltGraph is "EventModifierInit.modifierAltGraph"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierAltGraph MUST be set to true to make this field effective.
	ModifierAltGraph bool
	// ModifierCapsLock is "EventModifierInit.modifierCapsLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierCapsLock MUST be set to true to make this field effective.
	ModifierCapsLock bool
	// ModifierFn is "EventModifierInit.modifierFn"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierFn MUST be set to true to make this field effective.
	ModifierFn bool
	// ModifierFnLock is "EventModifierInit.modifierFnLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierFnLock MUST be set to true to make this field effective.
	ModifierFnLock bool
	// ModifierHyper is "EventModifierInit.modifierHyper"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierHyper MUST be set to true to make this field effective.
	ModifierHyper bool
	// ModifierNumLock is "EventModifierInit.modifierNumLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierNumLock MUST be set to true to make this field effective.
	ModifierNumLock bool
	// ModifierScrollLock is "EventModifierInit.modifierScrollLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierScrollLock MUST be set to true to make this field effective.
	ModifierScrollLock bool
	// ModifierSuper is "EventModifierInit.modifierSuper"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSuper MUST be set to true to make this field effective.
	ModifierSuper bool
	// ModifierSymbol is "EventModifierInit.modifierSymbol"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSymbol MUST be set to true to make this field effective.
	ModifierSymbol bool
	// ModifierSymbolLock is "EventModifierInit.modifierSymbolLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSymbolLock MUST be set to true to make this field effective.
	ModifierSymbolLock bool
	// View is "EventModifierInit.view"
	//
	// Optional, defaults to null.
	View Window
	// Detail is "EventModifierInit.detail"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Detail MUST be set to true to make this field effective.
	Detail int32
	// Bubbles is "EventModifierInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "EventModifierInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "EventModifierInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_CtrlKey            bool // for CtrlKey.
	FFI_USE_ShiftKey           bool // for ShiftKey.
	FFI_USE_AltKey             bool // for AltKey.
	FFI_USE_MetaKey            bool // for MetaKey.
	FFI_USE_ModifierAltGraph   bool // for ModifierAltGraph.
	FFI_USE_ModifierCapsLock   bool // for ModifierCapsLock.
	FFI_USE_ModifierFn         bool // for ModifierFn.
	FFI_USE_ModifierFnLock     bool // for ModifierFnLock.
	FFI_USE_ModifierHyper      bool // for ModifierHyper.
	FFI_USE_ModifierNumLock    bool // for ModifierNumLock.
	FFI_USE_ModifierScrollLock bool // for ModifierScrollLock.
	FFI_USE_ModifierSuper      bool // for ModifierSuper.
	FFI_USE_ModifierSymbol     bool // for ModifierSymbol.
	FFI_USE_ModifierSymbolLock bool // for ModifierSymbolLock.
	FFI_USE_Detail             bool // for Detail.
	FFI_USE_Bubbles            bool // for Bubbles.
	FFI_USE_Cancelable         bool // for Cancelable.
	FFI_USE_Composed           bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EventModifierInit with all fields set.
func (p EventModifierInit) FromRef(ref js.Ref) EventModifierInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EventModifierInit in the application heap.
func (p EventModifierInit) New() js.Ref {
	return bindings.EventModifierInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p EventModifierInit) UpdateFrom(ref js.Ref) {
	bindings.EventModifierInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p EventModifierInit) Update(ref js.Ref) {
	bindings.EventModifierInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

const (
	EventSource_CONNECTING uint16 = 0
	EventSource_OPEN       uint16 = 1
	EventSource_CLOSED     uint16 = 2
)

type EventSourceInit struct {
	// WithCredentials is "EventSourceInit.withCredentials"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_WithCredentials MUST be set to true to make this field effective.
	WithCredentials bool

	FFI_USE_WithCredentials bool // for WithCredentials.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EventSourceInit with all fields set.
func (p EventSourceInit) FromRef(ref js.Ref) EventSourceInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EventSourceInit in the application heap.
func (p EventSourceInit) New() js.Ref {
	return bindings.EventSourceInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p EventSourceInit) UpdateFrom(ref js.Ref) {
	bindings.EventSourceInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p EventSourceInit) Update(ref js.Ref) {
	bindings.EventSourceInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewEventSource(url js.String, eventSourceInitDict EventSourceInit) EventSource {
	return EventSource{}.FromRef(
		bindings.NewEventSourceByEventSource(
			url.Ref(),
			js.Pointer(&eventSourceInitDict)),
	)
}

func NewEventSourceByEventSource1(url js.String) EventSource {
	return EventSource{}.FromRef(
		bindings.NewEventSourceByEventSource1(
			url.Ref()),
	)
}

type EventSource struct {
	EventTarget
}

func (this EventSource) Once() EventSource {
	this.Ref().Once()
	return this
}

func (this EventSource) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this EventSource) FromRef(ref js.Ref) EventSource {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this EventSource) Free() {
	this.Ref().Free()
}

// Url returns the value of property "EventSource.url".
//
// The returned bool will be false if there is no such property.
func (this EventSource) Url() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetEventSourceUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// WithCredentials returns the value of property "EventSource.withCredentials".
//
// The returned bool will be false if there is no such property.
func (this EventSource) WithCredentials() (bool, bool) {
	var _ok bool
	_ret := bindings.GetEventSourceWithCredentials(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ReadyState returns the value of property "EventSource.readyState".
//
// The returned bool will be false if there is no such property.
func (this EventSource) ReadyState() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetEventSourceReadyState(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Close calls the method "EventSource.close".
//
// The returned bool will be false if there is no such method.
func (this EventSource) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEventSourceClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "EventSource.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EventSource) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.EventSourceCloseFunc(
			this.Ref(),
		),
	)
}

type ExtendableCookieChangeEventInit struct {
	// Changed is "ExtendableCookieChangeEventInit.changed"
	//
	// Optional
	Changed CookieList
	// Deleted is "ExtendableCookieChangeEventInit.deleted"
	//
	// Optional
	Deleted CookieList
	// Bubbles is "ExtendableCookieChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "ExtendableCookieChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "ExtendableCookieChangeEventInit.composed"
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

// FromRef calls UpdateFrom and returns a ExtendableCookieChangeEventInit with all fields set.
func (p ExtendableCookieChangeEventInit) FromRef(ref js.Ref) ExtendableCookieChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtendableCookieChangeEventInit in the application heap.
func (p ExtendableCookieChangeEventInit) New() js.Ref {
	return bindings.ExtendableCookieChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ExtendableCookieChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.ExtendableCookieChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ExtendableCookieChangeEventInit) Update(ref js.Ref) {
	bindings.ExtendableCookieChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewExtendableCookieChangeEvent(typ js.String, eventInitDict ExtendableCookieChangeEventInit) ExtendableCookieChangeEvent {
	return ExtendableCookieChangeEvent{}.FromRef(
		bindings.NewExtendableCookieChangeEventByExtendableCookieChangeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewExtendableCookieChangeEventByExtendableCookieChangeEvent1(typ js.String) ExtendableCookieChangeEvent {
	return ExtendableCookieChangeEvent{}.FromRef(
		bindings.NewExtendableCookieChangeEventByExtendableCookieChangeEvent1(
			typ.Ref()),
	)
}

type ExtendableCookieChangeEvent struct {
	ExtendableEvent
}

func (this ExtendableCookieChangeEvent) Once() ExtendableCookieChangeEvent {
	this.Ref().Once()
	return this
}

func (this ExtendableCookieChangeEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this ExtendableCookieChangeEvent) FromRef(ref js.Ref) ExtendableCookieChangeEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this ExtendableCookieChangeEvent) Free() {
	this.Ref().Free()
}

// Changed returns the value of property "ExtendableCookieChangeEvent.changed".
//
// The returned bool will be false if there is no such property.
func (this ExtendableCookieChangeEvent) Changed() (js.FrozenArray[CookieListItem], bool) {
	var _ok bool
	_ret := bindings.GetExtendableCookieChangeEventChanged(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[CookieListItem]{}.FromRef(_ret), _ok
}

// Deleted returns the value of property "ExtendableCookieChangeEvent.deleted".
//
// The returned bool will be false if there is no such property.
func (this ExtendableCookieChangeEvent) Deleted() (js.FrozenArray[CookieListItem], bool) {
	var _ok bool
	_ret := bindings.GetExtendableCookieChangeEventDeleted(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[CookieListItem]{}.FromRef(_ret), _ok
}

type ExtendableEventInit struct {
	// Bubbles is "ExtendableEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "ExtendableEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "ExtendableEventInit.composed"
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

// FromRef calls UpdateFrom and returns a ExtendableEventInit with all fields set.
func (p ExtendableEventInit) FromRef(ref js.Ref) ExtendableEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtendableEventInit in the application heap.
func (p ExtendableEventInit) New() js.Ref {
	return bindings.ExtendableEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ExtendableEventInit) UpdateFrom(ref js.Ref) {
	bindings.ExtendableEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ExtendableEventInit) Update(ref js.Ref) {
	bindings.ExtendableEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewExtendableEvent(typ js.String, eventInitDict ExtendableEventInit) ExtendableEvent {
	return ExtendableEvent{}.FromRef(
		bindings.NewExtendableEventByExtendableEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewExtendableEventByExtendableEvent1(typ js.String) ExtendableEvent {
	return ExtendableEvent{}.FromRef(
		bindings.NewExtendableEventByExtendableEvent1(
			typ.Ref()),
	)
}

type ExtendableEvent struct {
	Event
}

func (this ExtendableEvent) Once() ExtendableEvent {
	this.Ref().Once()
	return this
}

func (this ExtendableEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this ExtendableEvent) FromRef(ref js.Ref) ExtendableEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this ExtendableEvent) Free() {
	this.Ref().Free()
}

// WaitUntil calls the method "ExtendableEvent.waitUntil".
//
// The returned bool will be false if there is no such method.
func (this ExtendableEvent) WaitUntil(f js.Promise[js.Any]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallExtendableEventWaitUntil(
		this.Ref(), js.Pointer(&_ok),
		f.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// WaitUntilFunc returns the method "ExtendableEvent.waitUntil".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ExtendableEvent) WaitUntilFunc() (fn js.Func[func(f js.Promise[js.Any])]) {
	return fn.FromRef(
		bindings.ExtendableEventWaitUntilFunc(
			this.Ref(),
		),
	)
}

type OneOf_Client_ServiceWorker_MessagePort struct {
	ref js.Ref
}

func (x OneOf_Client_ServiceWorker_MessagePort) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Client_ServiceWorker_MessagePort) Free() {
	x.ref.Free()
}

func (x OneOf_Client_ServiceWorker_MessagePort) FromRef(ref js.Ref) OneOf_Client_ServiceWorker_MessagePort {
	return OneOf_Client_ServiceWorker_MessagePort{
		ref: ref,
	}
}

func (x OneOf_Client_ServiceWorker_MessagePort) Client() Client {
	return Client{}.FromRef(x.ref)
}

func (x OneOf_Client_ServiceWorker_MessagePort) ServiceWorker() ServiceWorker {
	return ServiceWorker{}.FromRef(x.ref)
}

func (x OneOf_Client_ServiceWorker_MessagePort) MessagePort() MessagePort {
	return MessagePort{}.FromRef(x.ref)
}

type ExtendableMessageEventInit struct {
	// Data is "ExtendableMessageEventInit.data"
	//
	// Optional, defaults to null.
	Data js.Any
	// Origin is "ExtendableMessageEventInit.origin"
	//
	// Optional, defaults to "".
	Origin js.String
	// LastEventId is "ExtendableMessageEventInit.lastEventId"
	//
	// Optional, defaults to "".
	LastEventId js.String
	// Source is "ExtendableMessageEventInit.source"
	//
	// Optional, defaults to null.
	Source OneOf_Client_ServiceWorker_MessagePort
	// Ports is "ExtendableMessageEventInit.ports"
	//
	// Optional, defaults to [].
	Ports js.Array[MessagePort]
	// Bubbles is "ExtendableMessageEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "ExtendableMessageEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "ExtendableMessageEventInit.composed"
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

// FromRef calls UpdateFrom and returns a ExtendableMessageEventInit with all fields set.
func (p ExtendableMessageEventInit) FromRef(ref js.Ref) ExtendableMessageEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ExtendableMessageEventInit in the application heap.
func (p ExtendableMessageEventInit) New() js.Ref {
	return bindings.ExtendableMessageEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ExtendableMessageEventInit) UpdateFrom(ref js.Ref) {
	bindings.ExtendableMessageEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ExtendableMessageEventInit) Update(ref js.Ref) {
	bindings.ExtendableMessageEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewExtendableMessageEvent(typ js.String, eventInitDict ExtendableMessageEventInit) ExtendableMessageEvent {
	return ExtendableMessageEvent{}.FromRef(
		bindings.NewExtendableMessageEventByExtendableMessageEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewExtendableMessageEventByExtendableMessageEvent1(typ js.String) ExtendableMessageEvent {
	return ExtendableMessageEvent{}.FromRef(
		bindings.NewExtendableMessageEventByExtendableMessageEvent1(
			typ.Ref()),
	)
}

type ExtendableMessageEvent struct {
	ExtendableEvent
}

func (this ExtendableMessageEvent) Once() ExtendableMessageEvent {
	this.Ref().Once()
	return this
}

func (this ExtendableMessageEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this ExtendableMessageEvent) FromRef(ref js.Ref) ExtendableMessageEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this ExtendableMessageEvent) Free() {
	this.Ref().Free()
}

// Data returns the value of property "ExtendableMessageEvent.data".
//
// The returned bool will be false if there is no such property.
func (this ExtendableMessageEvent) Data() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetExtendableMessageEventData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Origin returns the value of property "ExtendableMessageEvent.origin".
//
// The returned bool will be false if there is no such property.
func (this ExtendableMessageEvent) Origin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetExtendableMessageEventOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LastEventId returns the value of property "ExtendableMessageEvent.lastEventId".
//
// The returned bool will be false if there is no such property.
func (this ExtendableMessageEvent) LastEventId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetExtendableMessageEventLastEventId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Source returns the value of property "ExtendableMessageEvent.source".
//
// The returned bool will be false if there is no such property.
func (this ExtendableMessageEvent) Source() (OneOf_Client_ServiceWorker_MessagePort, bool) {
	var _ok bool
	_ret := bindings.GetExtendableMessageEventSource(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_Client_ServiceWorker_MessagePort{}.FromRef(_ret), _ok
}

// Ports returns the value of property "ExtendableMessageEvent.ports".
//
// The returned bool will be false if there is no such property.
func (this ExtendableMessageEvent) Ports() (js.FrozenArray[MessagePort], bool) {
	var _ok bool
	_ret := bindings.GetExtendableMessageEventPorts(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[MessagePort]{}.FromRef(_ret), _ok
}

type EyeDropper struct {
	ref js.Ref
}

func (this EyeDropper) Once() EyeDropper {
	this.Ref().Once()
	return this
}

func (this EyeDropper) Ref() js.Ref {
	return this.ref
}

func (this EyeDropper) FromRef(ref js.Ref) EyeDropper {
	this.ref = ref
	return this
}

func (this EyeDropper) Free() {
	this.Ref().Free()
}

// Open calls the method "EyeDropper.open".
//
// The returned bool will be false if there is no such method.
func (this EyeDropper) Open(options ColorSelectionOptions) (js.Promise[ColorSelectionResult], bool) {
	var _ok bool
	_ret := bindings.CallEyeDropperOpen(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[ColorSelectionResult]{}.FromRef(_ret), _ok
}

// OpenFunc returns the method "EyeDropper.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EyeDropper) OpenFunc() (fn js.Func[func(options ColorSelectionOptions) js.Promise[ColorSelectionResult]]) {
	return fn.FromRef(
		bindings.EyeDropperOpenFunc(
			this.Ref(),
		),
	)
}

// Open1 calls the method "EyeDropper.open".
//
// The returned bool will be false if there is no such method.
func (this EyeDropper) Open1() (js.Promise[ColorSelectionResult], bool) {
	var _ok bool
	_ret := bindings.CallEyeDropperOpen1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[ColorSelectionResult]{}.FromRef(_ret), _ok
}

// Open1Func returns the method "EyeDropper.open".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EyeDropper) Open1Func() (fn js.Func[func() js.Promise[ColorSelectionResult]]) {
	return fn.FromRef(
		bindings.EyeDropperOpen1Func(
			this.Ref(),
		),
	)
}

type FaceDetectorOptions struct {
	// MaxDetectedFaces is "FaceDetectorOptions.maxDetectedFaces"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxDetectedFaces MUST be set to true to make this field effective.
	MaxDetectedFaces uint16
	// FastMode is "FaceDetectorOptions.fastMode"
	//
	// Optional
	//
	// NOTE: FFI_USE_FastMode MUST be set to true to make this field effective.
	FastMode bool

	FFI_USE_MaxDetectedFaces bool // for MaxDetectedFaces.
	FFI_USE_FastMode         bool // for FastMode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FaceDetectorOptions with all fields set.
func (p FaceDetectorOptions) FromRef(ref js.Ref) FaceDetectorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FaceDetectorOptions in the application heap.
func (p FaceDetectorOptions) New() js.Ref {
	return bindings.FaceDetectorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FaceDetectorOptions) UpdateFrom(ref js.Ref) {
	bindings.FaceDetectorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FaceDetectorOptions) Update(ref js.Ref) {
	bindings.FaceDetectorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewFaceDetector(faceDetectorOptions FaceDetectorOptions) FaceDetector {
	return FaceDetector{}.FromRef(
		bindings.NewFaceDetectorByFaceDetector(
			js.Pointer(&faceDetectorOptions)),
	)
}

func NewFaceDetectorByFaceDetector1() FaceDetector {
	return FaceDetector{}.FromRef(
		bindings.NewFaceDetectorByFaceDetector1(),
	)
}

type FaceDetector struct {
	ref js.Ref
}

func (this FaceDetector) Once() FaceDetector {
	this.Ref().Once()
	return this
}

func (this FaceDetector) Ref() js.Ref {
	return this.ref
}

func (this FaceDetector) FromRef(ref js.Ref) FaceDetector {
	this.ref = ref
	return this
}

func (this FaceDetector) Free() {
	this.Ref().Free()
}

// Detect calls the method "FaceDetector.detect".
//
// The returned bool will be false if there is no such method.
func (this FaceDetector) Detect(image ImageBitmapSource) (js.Promise[js.Array[DetectedFace]], bool) {
	var _ok bool
	_ret := bindings.CallFaceDetectorDetect(
		this.Ref(), js.Pointer(&_ok),
		image.Ref(),
	)

	return js.Promise[js.Array[DetectedFace]]{}.FromRef(_ret), _ok
}

// DetectFunc returns the method "FaceDetector.detect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FaceDetector) DetectFunc() (fn js.Func[func(image ImageBitmapSource) js.Promise[js.Array[DetectedFace]]]) {
	return fn.FromRef(
		bindings.FaceDetectorDetectFunc(
			this.Ref(),
		),
	)
}

func NewFederatedCredential(data FederatedCredentialInit) FederatedCredential {
	return FederatedCredential{}.FromRef(
		bindings.NewFederatedCredentialByFederatedCredential(
			js.Pointer(&data)),
	)
}

type FederatedCredential struct {
	Credential
}

func (this FederatedCredential) Once() FederatedCredential {
	this.Ref().Once()
	return this
}

func (this FederatedCredential) Ref() js.Ref {
	return this.Credential.Ref()
}

func (this FederatedCredential) FromRef(ref js.Ref) FederatedCredential {
	this.Credential = this.Credential.FromRef(ref)
	return this
}

func (this FederatedCredential) Free() {
	this.Ref().Free()
}

// Provider returns the value of property "FederatedCredential.provider".
//
// The returned bool will be false if there is no such property.
func (this FederatedCredential) Provider() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFederatedCredentialProvider(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol returns the value of property "FederatedCredential.protocol".
//
// The returned bool will be false if there is no such property.
func (this FederatedCredential) Protocol() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFederatedCredentialProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name returns the value of property "FederatedCredential.name".
//
// The returned bool will be false if there is no such property.
func (this FederatedCredential) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFederatedCredentialName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// IconURL returns the value of property "FederatedCredential.iconURL".
//
// The returned bool will be false if there is no such property.
func (this FederatedCredential) IconURL() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFederatedCredentialIconURL(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type FencedFrameConfigURL = js.String

type FetchEventInit struct {
	// Request is "FetchEventInit.request"
	//
	// Required
	Request Request
	// PreloadResponse is "FetchEventInit.preloadResponse"
	//
	// Optional
	PreloadResponse js.Promise[js.Any]
	// ClientId is "FetchEventInit.clientId"
	//
	// Optional, defaults to "".
	ClientId js.String
	// ResultingClientId is "FetchEventInit.resultingClientId"
	//
	// Optional, defaults to "".
	ResultingClientId js.String
	// ReplacesClientId is "FetchEventInit.replacesClientId"
	//
	// Optional, defaults to "".
	ReplacesClientId js.String
	// Handled is "FetchEventInit.handled"
	//
	// Optional
	Handled js.Promise[js.Void]
	// Bubbles is "FetchEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "FetchEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "FetchEventInit.composed"
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

// FromRef calls UpdateFrom and returns a FetchEventInit with all fields set.
func (p FetchEventInit) FromRef(ref js.Ref) FetchEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FetchEventInit in the application heap.
func (p FetchEventInit) New() js.Ref {
	return bindings.FetchEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FetchEventInit) UpdateFrom(ref js.Ref) {
	bindings.FetchEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FetchEventInit) Update(ref js.Ref) {
	bindings.FetchEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewFetchEvent(typ js.String, eventInitDict FetchEventInit) FetchEvent {
	return FetchEvent{}.FromRef(
		bindings.NewFetchEventByFetchEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type FetchEvent struct {
	ExtendableEvent
}

func (this FetchEvent) Once() FetchEvent {
	this.Ref().Once()
	return this
}

func (this FetchEvent) Ref() js.Ref {
	return this.ExtendableEvent.Ref()
}

func (this FetchEvent) FromRef(ref js.Ref) FetchEvent {
	this.ExtendableEvent = this.ExtendableEvent.FromRef(ref)
	return this
}

func (this FetchEvent) Free() {
	this.Ref().Free()
}

// Request returns the value of property "FetchEvent.request".
//
// The returned bool will be false if there is no such property.
func (this FetchEvent) Request() (Request, bool) {
	var _ok bool
	_ret := bindings.GetFetchEventRequest(
		this.Ref(), js.Pointer(&_ok),
	)
	return Request{}.FromRef(_ret), _ok
}

// PreloadResponse returns the value of property "FetchEvent.preloadResponse".
//
// The returned bool will be false if there is no such property.
func (this FetchEvent) PreloadResponse() (js.Promise[js.Any], bool) {
	var _ok bool
	_ret := bindings.GetFetchEventPreloadResponse(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Any]{}.FromRef(_ret), _ok
}

// ClientId returns the value of property "FetchEvent.clientId".
//
// The returned bool will be false if there is no such property.
func (this FetchEvent) ClientId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFetchEventClientId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ResultingClientId returns the value of property "FetchEvent.resultingClientId".
//
// The returned bool will be false if there is no such property.
func (this FetchEvent) ResultingClientId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFetchEventResultingClientId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReplacesClientId returns the value of property "FetchEvent.replacesClientId".
//
// The returned bool will be false if there is no such property.
func (this FetchEvent) ReplacesClientId() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFetchEventReplacesClientId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Handled returns the value of property "FetchEvent.handled".
//
// The returned bool will be false if there is no such property.
func (this FetchEvent) Handled() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetFetchEventHandled(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// RespondWith calls the method "FetchEvent.respondWith".
//
// The returned bool will be false if there is no such method.
func (this FetchEvent) RespondWith(r js.Promise[Response]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFetchEventRespondWith(
		this.Ref(), js.Pointer(&_ok),
		r.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RespondWithFunc returns the method "FetchEvent.respondWith".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FetchEvent) RespondWithFunc() (fn js.Func[func(r js.Promise[Response])]) {
	return fn.FromRef(
		bindings.FetchEventRespondWithFunc(
			this.Ref(),
		),
	)
}

type FileCallbackFunc func(this js.Ref, file File) js.Ref

func (fn FileCallbackFunc) Register() js.Func[func(file File)] {
	return js.RegisterCallback[func(file File)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FileCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		File{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FileCallback[T any] struct {
	Fn  func(arg T, this js.Ref, file File) js.Ref
	Arg T
}

func (cb *FileCallback[T]) Register() js.Func[func(file File)] {
	return js.RegisterCallback[func(file File)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FileCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		File{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FilePickerOptions struct {
	// Types is "FilePickerOptions.types"
	//
	// Optional
	Types js.Array[FilePickerAcceptType]
	// ExcludeAcceptAllOption is "FilePickerOptions.excludeAcceptAllOption"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ExcludeAcceptAllOption MUST be set to true to make this field effective.
	ExcludeAcceptAllOption bool
	// Id is "FilePickerOptions.id"
	//
	// Optional
	Id js.String
	// StartIn is "FilePickerOptions.startIn"
	//
	// Optional
	StartIn StartInDirectory

	FFI_USE_ExcludeAcceptAllOption bool // for ExcludeAcceptAllOption.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FilePickerOptions with all fields set.
func (p FilePickerOptions) FromRef(ref js.Ref) FilePickerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FilePickerOptions in the application heap.
func (p FilePickerOptions) New() js.Ref {
	return bindings.FilePickerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FilePickerOptions) UpdateFrom(ref js.Ref) {
	bindings.FilePickerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FilePickerOptions) Update(ref js.Ref) {
	bindings.FilePickerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

const (
	FileReader_EMPTY   uint16 = 0
	FileReader_LOADING uint16 = 1
	FileReader_DONE    uint16 = 2
)

type OneOf_String_ArrayBuffer struct {
	ref js.Ref
}

func (x OneOf_String_ArrayBuffer) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_ArrayBuffer) Free() {
	x.ref.Free()
}

func (x OneOf_String_ArrayBuffer) FromRef(ref js.Ref) OneOf_String_ArrayBuffer {
	return OneOf_String_ArrayBuffer{
		ref: ref,
	}
}

func (x OneOf_String_ArrayBuffer) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_ArrayBuffer) ArrayBuffer() js.ArrayBuffer {
	return js.ArrayBuffer{}.FromRef(x.ref)
}

type FileReader struct {
	EventTarget
}

func (this FileReader) Once() FileReader {
	this.Ref().Once()
	return this
}

func (this FileReader) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this FileReader) FromRef(ref js.Ref) FileReader {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this FileReader) Free() {
	this.Ref().Free()
}

// ReadyState returns the value of property "FileReader.readyState".
//
// The returned bool will be false if there is no such property.
func (this FileReader) ReadyState() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetFileReaderReadyState(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Result returns the value of property "FileReader.result".
//
// The returned bool will be false if there is no such property.
func (this FileReader) Result() (OneOf_String_ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.GetFileReaderResult(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_String_ArrayBuffer{}.FromRef(_ret), _ok
}

// Error returns the value of property "FileReader.error".
//
// The returned bool will be false if there is no such property.
func (this FileReader) Error() (DOMException, bool) {
	var _ok bool
	_ret := bindings.GetFileReaderError(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMException{}.FromRef(_ret), _ok
}

// ReadAsArrayBuffer calls the method "FileReader.readAsArrayBuffer".
//
// The returned bool will be false if there is no such method.
func (this FileReader) ReadAsArrayBuffer(blob Blob) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileReaderReadAsArrayBuffer(
		this.Ref(), js.Pointer(&_ok),
		blob.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadAsArrayBufferFunc returns the method "FileReader.readAsArrayBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileReader) ReadAsArrayBufferFunc() (fn js.Func[func(blob Blob)]) {
	return fn.FromRef(
		bindings.FileReaderReadAsArrayBufferFunc(
			this.Ref(),
		),
	)
}

// ReadAsBinaryString calls the method "FileReader.readAsBinaryString".
//
// The returned bool will be false if there is no such method.
func (this FileReader) ReadAsBinaryString(blob Blob) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileReaderReadAsBinaryString(
		this.Ref(), js.Pointer(&_ok),
		blob.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadAsBinaryStringFunc returns the method "FileReader.readAsBinaryString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileReader) ReadAsBinaryStringFunc() (fn js.Func[func(blob Blob)]) {
	return fn.FromRef(
		bindings.FileReaderReadAsBinaryStringFunc(
			this.Ref(),
		),
	)
}

// ReadAsText calls the method "FileReader.readAsText".
//
// The returned bool will be false if there is no such method.
func (this FileReader) ReadAsText(blob Blob, encoding js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileReaderReadAsText(
		this.Ref(), js.Pointer(&_ok),
		blob.Ref(),
		encoding.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadAsTextFunc returns the method "FileReader.readAsText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileReader) ReadAsTextFunc() (fn js.Func[func(blob Blob, encoding js.String)]) {
	return fn.FromRef(
		bindings.FileReaderReadAsTextFunc(
			this.Ref(),
		),
	)
}

// ReadAsText1 calls the method "FileReader.readAsText".
//
// The returned bool will be false if there is no such method.
func (this FileReader) ReadAsText1(blob Blob) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileReaderReadAsText1(
		this.Ref(), js.Pointer(&_ok),
		blob.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadAsText1Func returns the method "FileReader.readAsText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileReader) ReadAsText1Func() (fn js.Func[func(blob Blob)]) {
	return fn.FromRef(
		bindings.FileReaderReadAsText1Func(
			this.Ref(),
		),
	)
}

// ReadAsDataURL calls the method "FileReader.readAsDataURL".
//
// The returned bool will be false if there is no such method.
func (this FileReader) ReadAsDataURL(blob Blob) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileReaderReadAsDataURL(
		this.Ref(), js.Pointer(&_ok),
		blob.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ReadAsDataURLFunc returns the method "FileReader.readAsDataURL".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileReader) ReadAsDataURLFunc() (fn js.Func[func(blob Blob)]) {
	return fn.FromRef(
		bindings.FileReaderReadAsDataURLFunc(
			this.Ref(),
		),
	)
}

// Abort calls the method "FileReader.abort".
//
// The returned bool will be false if there is no such method.
func (this FileReader) Abort() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileReaderAbort(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AbortFunc returns the method "FileReader.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileReader) AbortFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileReaderAbortFunc(
			this.Ref(),
		),
	)
}

type FileReaderSync struct {
	ref js.Ref
}

func (this FileReaderSync) Once() FileReaderSync {
	this.Ref().Once()
	return this
}

func (this FileReaderSync) Ref() js.Ref {
	return this.ref
}

func (this FileReaderSync) FromRef(ref js.Ref) FileReaderSync {
	this.ref = ref
	return this
}

func (this FileReaderSync) Free() {
	this.Ref().Free()
}

// ReadAsArrayBuffer calls the method "FileReaderSync.readAsArrayBuffer".
//
// The returned bool will be false if there is no such method.
func (this FileReaderSync) ReadAsArrayBuffer(blob Blob) (js.ArrayBuffer, bool) {
	var _ok bool
	_ret := bindings.CallFileReaderSyncReadAsArrayBuffer(
		this.Ref(), js.Pointer(&_ok),
		blob.Ref(),
	)

	return js.ArrayBuffer{}.FromRef(_ret), _ok
}

// ReadAsArrayBufferFunc returns the method "FileReaderSync.readAsArrayBuffer".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileReaderSync) ReadAsArrayBufferFunc() (fn js.Func[func(blob Blob) js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.FileReaderSyncReadAsArrayBufferFunc(
			this.Ref(),
		),
	)
}

// ReadAsBinaryString calls the method "FileReaderSync.readAsBinaryString".
//
// The returned bool will be false if there is no such method.
func (this FileReaderSync) ReadAsBinaryString(blob Blob) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallFileReaderSyncReadAsBinaryString(
		this.Ref(), js.Pointer(&_ok),
		blob.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ReadAsBinaryStringFunc returns the method "FileReaderSync.readAsBinaryString".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileReaderSync) ReadAsBinaryStringFunc() (fn js.Func[func(blob Blob) js.String]) {
	return fn.FromRef(
		bindings.FileReaderSyncReadAsBinaryStringFunc(
			this.Ref(),
		),
	)
}

// ReadAsText calls the method "FileReaderSync.readAsText".
//
// The returned bool will be false if there is no such method.
func (this FileReaderSync) ReadAsText(blob Blob, encoding js.String) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallFileReaderSyncReadAsText(
		this.Ref(), js.Pointer(&_ok),
		blob.Ref(),
		encoding.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ReadAsTextFunc returns the method "FileReaderSync.readAsText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileReaderSync) ReadAsTextFunc() (fn js.Func[func(blob Blob, encoding js.String) js.String]) {
	return fn.FromRef(
		bindings.FileReaderSyncReadAsTextFunc(
			this.Ref(),
		),
	)
}

// ReadAsText1 calls the method "FileReaderSync.readAsText".
//
// The returned bool will be false if there is no such method.
func (this FileReaderSync) ReadAsText1(blob Blob) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallFileReaderSyncReadAsText1(
		this.Ref(), js.Pointer(&_ok),
		blob.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ReadAsText1Func returns the method "FileReaderSync.readAsText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileReaderSync) ReadAsText1Func() (fn js.Func[func(blob Blob) js.String]) {
	return fn.FromRef(
		bindings.FileReaderSyncReadAsText1Func(
			this.Ref(),
		),
	)
}

// ReadAsDataURL calls the method "FileReaderSync.readAsDataURL".
//
// The returned bool will be false if there is no such method.
func (this FileReaderSync) ReadAsDataURL(blob Blob) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallFileReaderSyncReadAsDataURL(
		this.Ref(), js.Pointer(&_ok),
		blob.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// ReadAsDataURLFunc returns the method "FileReaderSync.readAsDataURL".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileReaderSync) ReadAsDataURLFunc() (fn js.Func[func(blob Blob) js.String]) {
	return fn.FromRef(
		bindings.FileReaderSyncReadAsDataURLFunc(
			this.Ref(),
		),
	)
}

type FileSystemFileEntry struct {
	FileSystemEntry
}

func (this FileSystemFileEntry) Once() FileSystemFileEntry {
	this.Ref().Once()
	return this
}

func (this FileSystemFileEntry) Ref() js.Ref {
	return this.FileSystemEntry.Ref()
}

func (this FileSystemFileEntry) FromRef(ref js.Ref) FileSystemFileEntry {
	this.FileSystemEntry = this.FileSystemEntry.FromRef(ref)
	return this
}

func (this FileSystemFileEntry) Free() {
	this.Ref().Free()
}

// File calls the method "FileSystemFileEntry.file".
//
// The returned bool will be false if there is no such method.
func (this FileSystemFileEntry) File(successCallback js.Func[func(file File)], errorCallback js.Func[func(err DOMException)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemFileEntryFile(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// FileFunc returns the method "FileSystemFileEntry.file".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemFileEntry) FileFunc() (fn js.Func[func(successCallback js.Func[func(file File)], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.FileSystemFileEntryFileFunc(
			this.Ref(),
		),
	)
}

// File1 calls the method "FileSystemFileEntry.file".
//
// The returned bool will be false if there is no such method.
func (this FileSystemFileEntry) File1(successCallback js.Func[func(file File)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFileSystemFileEntryFile1(
		this.Ref(), js.Pointer(&_ok),
		successCallback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// File1Func returns the method "FileSystemFileEntry.file".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FileSystemFileEntry) File1Func() (fn js.Func[func(successCallback js.Func[func(file File)])]) {
	return fn.FromRef(
		bindings.FileSystemFileEntryFile1Func(
			this.Ref(),
		),
	)
}

type FileSystemPermissionDescriptor struct {
	// Handle is "FileSystemPermissionDescriptor.handle"
	//
	// Required
	Handle FileSystemHandle
	// Mode is "FileSystemPermissionDescriptor.mode"
	//
	// Optional, defaults to "read".
	Mode FileSystemPermissionMode
	// Name is "FileSystemPermissionDescriptor.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FileSystemPermissionDescriptor with all fields set.
func (p FileSystemPermissionDescriptor) FromRef(ref js.Ref) FileSystemPermissionDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FileSystemPermissionDescriptor in the application heap.
func (p FileSystemPermissionDescriptor) New() js.Ref {
	return bindings.FileSystemPermissionDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FileSystemPermissionDescriptor) UpdateFrom(ref js.Ref) {
	bindings.FileSystemPermissionDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FileSystemPermissionDescriptor) Update(ref js.Ref) {
	bindings.FileSystemPermissionDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type FillLightMode uint32

const (
	_ FillLightMode = iota

	FillLightMode_AUTO
	FillLightMode_OFF
	FillLightMode_FLASH
)

func (FillLightMode) FromRef(str js.Ref) FillLightMode {
	return FillLightMode(bindings.ConstOfFillLightMode(str))
}

func (x FillLightMode) String() (string, bool) {
	switch x {
	case FillLightMode_AUTO:
		return "auto", true
	case FillLightMode_OFF:
		return "off", true
	case FillLightMode_FLASH:
		return "flash", true
	default:
		return "", false
	}
}

type FocusEventInit struct {
	// RelatedTarget is "FocusEventInit.relatedTarget"
	//
	// Optional, defaults to null.
	RelatedTarget EventTarget
	// View is "FocusEventInit.view"
	//
	// Optional, defaults to null.
	View Window
	// Detail is "FocusEventInit.detail"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Detail MUST be set to true to make this field effective.
	Detail int32
	// Bubbles is "FocusEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "FocusEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "FocusEventInit.composed"
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

// FromRef calls UpdateFrom and returns a FocusEventInit with all fields set.
func (p FocusEventInit) FromRef(ref js.Ref) FocusEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FocusEventInit in the application heap.
func (p FocusEventInit) New() js.Ref {
	return bindings.FocusEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FocusEventInit) UpdateFrom(ref js.Ref) {
	bindings.FocusEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FocusEventInit) Update(ref js.Ref) {
	bindings.FocusEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewFocusEvent(typ js.String, eventInitDict FocusEventInit) FocusEvent {
	return FocusEvent{}.FromRef(
		bindings.NewFocusEventByFocusEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewFocusEventByFocusEvent1(typ js.String) FocusEvent {
	return FocusEvent{}.FromRef(
		bindings.NewFocusEventByFocusEvent1(
			typ.Ref()),
	)
}

type FocusEvent struct {
	UIEvent
}

func (this FocusEvent) Once() FocusEvent {
	this.Ref().Once()
	return this
}

func (this FocusEvent) Ref() js.Ref {
	return this.UIEvent.Ref()
}

func (this FocusEvent) FromRef(ref js.Ref) FocusEvent {
	this.UIEvent = this.UIEvent.FromRef(ref)
	return this
}

func (this FocusEvent) Free() {
	this.Ref().Free()
}

// RelatedTarget returns the value of property "FocusEvent.relatedTarget".
//
// The returned bool will be false if there is no such property.
func (this FocusEvent) RelatedTarget() (EventTarget, bool) {
	var _ok bool
	_ret := bindings.GetFocusEventRelatedTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return EventTarget{}.FromRef(_ret), _ok
}

type FontFaceSetLoadEventInit struct {
	// Fontfaces is "FontFaceSetLoadEventInit.fontfaces"
	//
	// Optional, defaults to [].
	Fontfaces js.Array[FontFace]
	// Bubbles is "FontFaceSetLoadEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "FontFaceSetLoadEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "FontFaceSetLoadEventInit.composed"
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

// FromRef calls UpdateFrom and returns a FontFaceSetLoadEventInit with all fields set.
func (p FontFaceSetLoadEventInit) FromRef(ref js.Ref) FontFaceSetLoadEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FontFaceSetLoadEventInit in the application heap.
func (p FontFaceSetLoadEventInit) New() js.Ref {
	return bindings.FontFaceSetLoadEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FontFaceSetLoadEventInit) UpdateFrom(ref js.Ref) {
	bindings.FontFaceSetLoadEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FontFaceSetLoadEventInit) Update(ref js.Ref) {
	bindings.FontFaceSetLoadEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewFontFaceSetLoadEvent(typ js.String, eventInitDict FontFaceSetLoadEventInit) FontFaceSetLoadEvent {
	return FontFaceSetLoadEvent{}.FromRef(
		bindings.NewFontFaceSetLoadEventByFontFaceSetLoadEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewFontFaceSetLoadEventByFontFaceSetLoadEvent1(typ js.String) FontFaceSetLoadEvent {
	return FontFaceSetLoadEvent{}.FromRef(
		bindings.NewFontFaceSetLoadEventByFontFaceSetLoadEvent1(
			typ.Ref()),
	)
}

type FontFaceSetLoadEvent struct {
	Event
}

func (this FontFaceSetLoadEvent) Once() FontFaceSetLoadEvent {
	this.Ref().Once()
	return this
}

func (this FontFaceSetLoadEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this FontFaceSetLoadEvent) FromRef(ref js.Ref) FontFaceSetLoadEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this FontFaceSetLoadEvent) Free() {
	this.Ref().Free()
}

// Fontfaces returns the value of property "FontFaceSetLoadEvent.fontfaces".
//
// The returned bool will be false if there is no such property.
func (this FontFaceSetLoadEvent) Fontfaces() (js.FrozenArray[FontFace], bool) {
	var _ok bool
	_ret := bindings.GetFontFaceSetLoadEventFontfaces(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[FontFace]{}.FromRef(_ret), _ok
}

type FontFaceVariationAxis struct {
	ref js.Ref
}

func (this FontFaceVariationAxis) Once() FontFaceVariationAxis {
	this.Ref().Once()
	return this
}

func (this FontFaceVariationAxis) Ref() js.Ref {
	return this.ref
}

func (this FontFaceVariationAxis) FromRef(ref js.Ref) FontFaceVariationAxis {
	this.ref = ref
	return this
}

func (this FontFaceVariationAxis) Free() {
	this.Ref().Free()
}

// Name returns the value of property "FontFaceVariationAxis.name".
//
// The returned bool will be false if there is no such property.
func (this FontFaceVariationAxis) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceVariationAxisName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AxisTag returns the value of property "FontFaceVariationAxis.axisTag".
//
// The returned bool will be false if there is no such property.
func (this FontFaceVariationAxis) AxisTag() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceVariationAxisAxisTag(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// MinimumValue returns the value of property "FontFaceVariationAxis.minimumValue".
//
// The returned bool will be false if there is no such property.
func (this FontFaceVariationAxis) MinimumValue() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceVariationAxisMinimumValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// MaximumValue returns the value of property "FontFaceVariationAxis.maximumValue".
//
// The returned bool will be false if there is no such property.
func (this FontFaceVariationAxis) MaximumValue() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceVariationAxisMaximumValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// DefaultValue returns the value of property "FontFaceVariationAxis.defaultValue".
//
// The returned bool will be false if there is no such property.
func (this FontFaceVariationAxis) DefaultValue() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFontFaceVariationAxisDefaultValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type FormDataEventInit struct {
	// FormData is "FormDataEventInit.formData"
	//
	// Required
	FormData FormData
	// Bubbles is "FormDataEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "FormDataEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "FormDataEventInit.composed"
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

// FromRef calls UpdateFrom and returns a FormDataEventInit with all fields set.
func (p FormDataEventInit) FromRef(ref js.Ref) FormDataEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FormDataEventInit in the application heap.
func (p FormDataEventInit) New() js.Ref {
	return bindings.FormDataEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FormDataEventInit) UpdateFrom(ref js.Ref) {
	bindings.FormDataEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FormDataEventInit) Update(ref js.Ref) {
	bindings.FormDataEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewFormDataEvent(typ js.String, eventInitDict FormDataEventInit) FormDataEvent {
	return FormDataEvent{}.FromRef(
		bindings.NewFormDataEventByFormDataEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type FormDataEvent struct {
	Event
}

func (this FormDataEvent) Once() FormDataEvent {
	this.Ref().Once()
	return this
}

func (this FormDataEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this FormDataEvent) FromRef(ref js.Ref) FormDataEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this FormDataEvent) Free() {
	this.Ref().Free()
}

// FormData returns the value of property "FormDataEvent.formData".
//
// The returned bool will be false if there is no such property.
func (this FormDataEvent) FormData() (FormData, bool) {
	var _ok bool
	_ret := bindings.GetFormDataEventFormData(
		this.Ref(), js.Pointer(&_ok),
	)
	return FormData{}.FromRef(_ret), _ok
}

type FragmentResultOptions struct {
	// InlineSize is "FragmentResultOptions.inlineSize"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_InlineSize MUST be set to true to make this field effective.
	InlineSize float64
	// BlockSize is "FragmentResultOptions.blockSize"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_BlockSize MUST be set to true to make this field effective.
	BlockSize float64
	// AutoBlockSize is "FragmentResultOptions.autoBlockSize"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_AutoBlockSize MUST be set to true to make this field effective.
	AutoBlockSize float64
	// ChildFragments is "FragmentResultOptions.childFragments"
	//
	// Optional, defaults to [].
	ChildFragments js.Array[LayoutFragment]
	// Data is "FragmentResultOptions.data"
	//
	// Optional, defaults to null.
	Data js.Any
	// BreakToken is "FragmentResultOptions.breakToken"
	//
	// Optional, defaults to null.
	BreakToken BreakTokenOptions

	FFI_USE_InlineSize    bool // for InlineSize.
	FFI_USE_BlockSize     bool // for BlockSize.
	FFI_USE_AutoBlockSize bool // for AutoBlockSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FragmentResultOptions with all fields set.
func (p FragmentResultOptions) FromRef(ref js.Ref) FragmentResultOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FragmentResultOptions in the application heap.
func (p FragmentResultOptions) New() js.Ref {
	return bindings.FragmentResultOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p FragmentResultOptions) UpdateFrom(ref js.Ref) {
	bindings.FragmentResultOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p FragmentResultOptions) Update(ref js.Ref) {
	bindings.FragmentResultOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewFragmentResult(options FragmentResultOptions) FragmentResult {
	return FragmentResult{}.FromRef(
		bindings.NewFragmentResultByFragmentResult(
			js.Pointer(&options)),
	)
}

func NewFragmentResultByFragmentResult1() FragmentResult {
	return FragmentResult{}.FromRef(
		bindings.NewFragmentResultByFragmentResult1(),
	)
}

type FragmentResult struct {
	ref js.Ref
}

func (this FragmentResult) Once() FragmentResult {
	this.Ref().Once()
	return this
}

func (this FragmentResult) Ref() js.Ref {
	return this.ref
}

func (this FragmentResult) FromRef(ref js.Ref) FragmentResult {
	this.ref = ref
	return this
}

func (this FragmentResult) Free() {
	this.Ref().Free()
}

// InlineSize returns the value of property "FragmentResult.inlineSize".
//
// The returned bool will be false if there is no such property.
func (this FragmentResult) InlineSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFragmentResultInlineSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// BlockSize returns the value of property "FragmentResult.blockSize".
//
// The returned bool will be false if there is no such property.
func (this FragmentResult) BlockSize() (float64, bool) {
	var _ok bool
	_ret := bindings.GetFragmentResultBlockSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

type GLbyte int8

type GLshort int16

type GLubyte uint8

type GLuint64EXT uint64

type GLushort uint16

const (
	GPUBufferUsage_MAP_READ      GPUFlagsConstant = 0x0001
	GPUBufferUsage_MAP_WRITE     GPUFlagsConstant = 0x0002
	GPUBufferUsage_COPY_SRC      GPUFlagsConstant = 0x0004
	GPUBufferUsage_COPY_DST      GPUFlagsConstant = 0x0008
	GPUBufferUsage_INDEX         GPUFlagsConstant = 0x0010
	GPUBufferUsage_VERTEX        GPUFlagsConstant = 0x0020
	GPUBufferUsage_UNIFORM       GPUFlagsConstant = 0x0040
	GPUBufferUsage_STORAGE       GPUFlagsConstant = 0x0080
	GPUBufferUsage_INDIRECT      GPUFlagsConstant = 0x0100
	GPUBufferUsage_QUERY_RESOLVE GPUFlagsConstant = 0x0200
)

type GPUBufferUsage struct{}

const (
	GPUColorWrite_RED   GPUFlagsConstant = 0x1
	GPUColorWrite_GREEN GPUFlagsConstant = 0x2
	GPUColorWrite_BLUE  GPUFlagsConstant = 0x4
	GPUColorWrite_ALPHA GPUFlagsConstant = 0x8
	GPUColorWrite_ALL   GPUFlagsConstant = 0xF
)

type GPUColorWrite struct{}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
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
	//
	// NOTE: DecoderConfig.FFI_USE MUST be set to true to get DecoderConfig used.
	DecoderConfig VideoDecoderConfig
	// Svc is "EncodedVideoChunkMetadata.svc"
	//
	// Optional
	//
	// NOTE: Svc.FFI_USE MUST be set to true to get Svc used.
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

func NewErrorEvent(typ js.String, eventInitDict ErrorEventInit) (ret ErrorEvent) {
	ret.ref = bindings.NewErrorEventByErrorEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewErrorEventByErrorEvent1(typ js.String) (ret ErrorEvent) {
	ret.ref = bindings.NewErrorEventByErrorEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this ErrorEvent) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetErrorEventMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Filename returns the value of property "ErrorEvent.filename".
//
// It returns ok=false if there is no such property.
func (this ErrorEvent) Filename() (ret js.String, ok bool) {
	ok = js.True == bindings.GetErrorEventFilename(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Lineno returns the value of property "ErrorEvent.lineno".
//
// It returns ok=false if there is no such property.
func (this ErrorEvent) Lineno() (ret uint32, ok bool) {
	ok = js.True == bindings.GetErrorEventLineno(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Colno returns the value of property "ErrorEvent.colno".
//
// It returns ok=false if there is no such property.
func (this ErrorEvent) Colno() (ret uint32, ok bool) {
	ok = js.True == bindings.GetErrorEventColno(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Error returns the value of property "ErrorEvent.error".
//
// It returns ok=false if there is no such property.
func (this ErrorEvent) Error() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetErrorEventError(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewEventSource(url js.String, eventSourceInitDict EventSourceInit) (ret EventSource) {
	ret.ref = bindings.NewEventSourceByEventSource(
		url.Ref(),
		js.Pointer(&eventSourceInitDict))
	return
}

func NewEventSourceByEventSource1(url js.String) (ret EventSource) {
	ret.ref = bindings.NewEventSourceByEventSource1(
		url.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this EventSource) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetEventSourceUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WithCredentials returns the value of property "EventSource.withCredentials".
//
// It returns ok=false if there is no such property.
func (this EventSource) WithCredentials() (ret bool, ok bool) {
	ok = js.True == bindings.GetEventSourceWithCredentials(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReadyState returns the value of property "EventSource.readyState".
//
// It returns ok=false if there is no such property.
func (this EventSource) ReadyState() (ret uint16, ok bool) {
	ok = js.True == bindings.GetEventSourceReadyState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasClose returns true if the method "EventSource.close" exists.
func (this EventSource) HasClose() bool {
	return js.True == bindings.HasEventSourceClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "EventSource.close".
func (this EventSource) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.EventSourceCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "EventSource.close".
func (this EventSource) Close() (ret js.Void) {
	bindings.CallEventSourceClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "EventSource.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EventSource) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventSourceClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewExtendableCookieChangeEvent(typ js.String, eventInitDict ExtendableCookieChangeEventInit) (ret ExtendableCookieChangeEvent) {
	ret.ref = bindings.NewExtendableCookieChangeEventByExtendableCookieChangeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewExtendableCookieChangeEventByExtendableCookieChangeEvent1(typ js.String) (ret ExtendableCookieChangeEvent) {
	ret.ref = bindings.NewExtendableCookieChangeEventByExtendableCookieChangeEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this ExtendableCookieChangeEvent) Changed() (ret js.FrozenArray[CookieListItem], ok bool) {
	ok = js.True == bindings.GetExtendableCookieChangeEventChanged(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Deleted returns the value of property "ExtendableCookieChangeEvent.deleted".
//
// It returns ok=false if there is no such property.
func (this ExtendableCookieChangeEvent) Deleted() (ret js.FrozenArray[CookieListItem], ok bool) {
	ok = js.True == bindings.GetExtendableCookieChangeEventDeleted(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewExtendableEvent(typ js.String, eventInitDict ExtendableEventInit) (ret ExtendableEvent) {
	ret.ref = bindings.NewExtendableEventByExtendableEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewExtendableEventByExtendableEvent1(typ js.String) (ret ExtendableEvent) {
	ret.ref = bindings.NewExtendableEventByExtendableEvent1(
		typ.Ref())
	return
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

// HasWaitUntil returns true if the method "ExtendableEvent.waitUntil" exists.
func (this ExtendableEvent) HasWaitUntil() bool {
	return js.True == bindings.HasExtendableEventWaitUntil(
		this.Ref(),
	)
}

// WaitUntilFunc returns the method "ExtendableEvent.waitUntil".
func (this ExtendableEvent) WaitUntilFunc() (fn js.Func[func(f js.Promise[js.Any])]) {
	return fn.FromRef(
		bindings.ExtendableEventWaitUntilFunc(
			this.Ref(),
		),
	)
}

// WaitUntil calls the method "ExtendableEvent.waitUntil".
func (this ExtendableEvent) WaitUntil(f js.Promise[js.Any]) (ret js.Void) {
	bindings.CallExtendableEventWaitUntil(
		this.Ref(), js.Pointer(&ret),
		f.Ref(),
	)

	return
}

// TryWaitUntil calls the method "ExtendableEvent.waitUntil"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ExtendableEvent) TryWaitUntil(f js.Promise[js.Any]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryExtendableEventWaitUntil(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		f.Ref(),
	)

	return
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

func NewExtendableMessageEvent(typ js.String, eventInitDict ExtendableMessageEventInit) (ret ExtendableMessageEvent) {
	ret.ref = bindings.NewExtendableMessageEventByExtendableMessageEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewExtendableMessageEventByExtendableMessageEvent1(typ js.String) (ret ExtendableMessageEvent) {
	ret.ref = bindings.NewExtendableMessageEventByExtendableMessageEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this ExtendableMessageEvent) Data() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetExtendableMessageEventData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Origin returns the value of property "ExtendableMessageEvent.origin".
//
// It returns ok=false if there is no such property.
func (this ExtendableMessageEvent) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetExtendableMessageEventOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LastEventId returns the value of property "ExtendableMessageEvent.lastEventId".
//
// It returns ok=false if there is no such property.
func (this ExtendableMessageEvent) LastEventId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetExtendableMessageEventLastEventId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Source returns the value of property "ExtendableMessageEvent.source".
//
// It returns ok=false if there is no such property.
func (this ExtendableMessageEvent) Source() (ret OneOf_Client_ServiceWorker_MessagePort, ok bool) {
	ok = js.True == bindings.GetExtendableMessageEventSource(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Ports returns the value of property "ExtendableMessageEvent.ports".
//
// It returns ok=false if there is no such property.
func (this ExtendableMessageEvent) Ports() (ret js.FrozenArray[MessagePort], ok bool) {
	ok = js.True == bindings.GetExtendableMessageEventPorts(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

// HasOpen returns true if the method "EyeDropper.open" exists.
func (this EyeDropper) HasOpen() bool {
	return js.True == bindings.HasEyeDropperOpen(
		this.Ref(),
	)
}

// OpenFunc returns the method "EyeDropper.open".
func (this EyeDropper) OpenFunc() (fn js.Func[func(options ColorSelectionOptions) js.Promise[ColorSelectionResult]]) {
	return fn.FromRef(
		bindings.EyeDropperOpenFunc(
			this.Ref(),
		),
	)
}

// Open calls the method "EyeDropper.open".
func (this EyeDropper) Open(options ColorSelectionOptions) (ret js.Promise[ColorSelectionResult]) {
	bindings.CallEyeDropperOpen(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryOpen calls the method "EyeDropper.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EyeDropper) TryOpen(options ColorSelectionOptions) (ret js.Promise[ColorSelectionResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEyeDropperOpen(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasOpen1 returns true if the method "EyeDropper.open" exists.
func (this EyeDropper) HasOpen1() bool {
	return js.True == bindings.HasEyeDropperOpen1(
		this.Ref(),
	)
}

// Open1Func returns the method "EyeDropper.open".
func (this EyeDropper) Open1Func() (fn js.Func[func() js.Promise[ColorSelectionResult]]) {
	return fn.FromRef(
		bindings.EyeDropperOpen1Func(
			this.Ref(),
		),
	)
}

// Open1 calls the method "EyeDropper.open".
func (this EyeDropper) Open1() (ret js.Promise[ColorSelectionResult]) {
	bindings.CallEyeDropperOpen1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryOpen1 calls the method "EyeDropper.open"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EyeDropper) TryOpen1() (ret js.Promise[ColorSelectionResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEyeDropperOpen1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

func NewFaceDetector(faceDetectorOptions FaceDetectorOptions) (ret FaceDetector) {
	ret.ref = bindings.NewFaceDetectorByFaceDetector(
		js.Pointer(&faceDetectorOptions))
	return
}

func NewFaceDetectorByFaceDetector1() (ret FaceDetector) {
	ret.ref = bindings.NewFaceDetectorByFaceDetector1()
	return
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

// HasDetect returns true if the method "FaceDetector.detect" exists.
func (this FaceDetector) HasDetect() bool {
	return js.True == bindings.HasFaceDetectorDetect(
		this.Ref(),
	)
}

// DetectFunc returns the method "FaceDetector.detect".
func (this FaceDetector) DetectFunc() (fn js.Func[func(image ImageBitmapSource) js.Promise[js.Array[DetectedFace]]]) {
	return fn.FromRef(
		bindings.FaceDetectorDetectFunc(
			this.Ref(),
		),
	)
}

// Detect calls the method "FaceDetector.detect".
func (this FaceDetector) Detect(image ImageBitmapSource) (ret js.Promise[js.Array[DetectedFace]]) {
	bindings.CallFaceDetectorDetect(
		this.Ref(), js.Pointer(&ret),
		image.Ref(),
	)

	return
}

// TryDetect calls the method "FaceDetector.detect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FaceDetector) TryDetect(image ImageBitmapSource) (ret js.Promise[js.Array[DetectedFace]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryFaceDetectorDetect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		image.Ref(),
	)

	return
}

func NewFederatedCredential(data FederatedCredentialInit) (ret FederatedCredential) {
	ret.ref = bindings.NewFederatedCredentialByFederatedCredential(
		js.Pointer(&data))
	return
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
// It returns ok=false if there is no such property.
func (this FederatedCredential) Provider() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFederatedCredentialProvider(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "FederatedCredential.protocol".
//
// It returns ok=false if there is no such property.
func (this FederatedCredential) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFederatedCredentialProtocol(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "FederatedCredential.name".
//
// It returns ok=false if there is no such property.
func (this FederatedCredential) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFederatedCredentialName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IconURL returns the value of property "FederatedCredential.iconURL".
//
// It returns ok=false if there is no such property.
func (this FederatedCredential) IconURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFederatedCredentialIconURL(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewFetchEvent(typ js.String, eventInitDict FetchEventInit) (ret FetchEvent) {
	ret.ref = bindings.NewFetchEventByFetchEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this FetchEvent) Request() (ret Request, ok bool) {
	ok = js.True == bindings.GetFetchEventRequest(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PreloadResponse returns the value of property "FetchEvent.preloadResponse".
//
// It returns ok=false if there is no such property.
func (this FetchEvent) PreloadResponse() (ret js.Promise[js.Any], ok bool) {
	ok = js.True == bindings.GetFetchEventPreloadResponse(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ClientId returns the value of property "FetchEvent.clientId".
//
// It returns ok=false if there is no such property.
func (this FetchEvent) ClientId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFetchEventClientId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ResultingClientId returns the value of property "FetchEvent.resultingClientId".
//
// It returns ok=false if there is no such property.
func (this FetchEvent) ResultingClientId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFetchEventResultingClientId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ReplacesClientId returns the value of property "FetchEvent.replacesClientId".
//
// It returns ok=false if there is no such property.
func (this FetchEvent) ReplacesClientId() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFetchEventReplacesClientId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Handled returns the value of property "FetchEvent.handled".
//
// It returns ok=false if there is no such property.
func (this FetchEvent) Handled() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetFetchEventHandled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasRespondWith returns true if the method "FetchEvent.respondWith" exists.
func (this FetchEvent) HasRespondWith() bool {
	return js.True == bindings.HasFetchEventRespondWith(
		this.Ref(),
	)
}

// RespondWithFunc returns the method "FetchEvent.respondWith".
func (this FetchEvent) RespondWithFunc() (fn js.Func[func(r js.Promise[Response])]) {
	return fn.FromRef(
		bindings.FetchEventRespondWithFunc(
			this.Ref(),
		),
	)
}

// RespondWith calls the method "FetchEvent.respondWith".
func (this FetchEvent) RespondWith(r js.Promise[Response]) (ret js.Void) {
	bindings.CallFetchEventRespondWith(
		this.Ref(), js.Pointer(&ret),
		r.Ref(),
	)

	return
}

// TryRespondWith calls the method "FetchEvent.respondWith"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FetchEvent) TryRespondWith(r js.Promise[Response]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFetchEventRespondWith(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		r.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this FileReader) ReadyState() (ret uint16, ok bool) {
	ok = js.True == bindings.GetFileReaderReadyState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Result returns the value of property "FileReader.result".
//
// It returns ok=false if there is no such property.
func (this FileReader) Result() (ret OneOf_String_ArrayBuffer, ok bool) {
	ok = js.True == bindings.GetFileReaderResult(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Error returns the value of property "FileReader.error".
//
// It returns ok=false if there is no such property.
func (this FileReader) Error() (ret DOMException, ok bool) {
	ok = js.True == bindings.GetFileReaderError(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasReadAsArrayBuffer returns true if the method "FileReader.readAsArrayBuffer" exists.
func (this FileReader) HasReadAsArrayBuffer() bool {
	return js.True == bindings.HasFileReaderReadAsArrayBuffer(
		this.Ref(),
	)
}

// ReadAsArrayBufferFunc returns the method "FileReader.readAsArrayBuffer".
func (this FileReader) ReadAsArrayBufferFunc() (fn js.Func[func(blob Blob)]) {
	return fn.FromRef(
		bindings.FileReaderReadAsArrayBufferFunc(
			this.Ref(),
		),
	)
}

// ReadAsArrayBuffer calls the method "FileReader.readAsArrayBuffer".
func (this FileReader) ReadAsArrayBuffer(blob Blob) (ret js.Void) {
	bindings.CallFileReaderReadAsArrayBuffer(
		this.Ref(), js.Pointer(&ret),
		blob.Ref(),
	)

	return
}

// TryReadAsArrayBuffer calls the method "FileReader.readAsArrayBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileReader) TryReadAsArrayBuffer(blob Blob) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileReaderReadAsArrayBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		blob.Ref(),
	)

	return
}

// HasReadAsBinaryString returns true if the method "FileReader.readAsBinaryString" exists.
func (this FileReader) HasReadAsBinaryString() bool {
	return js.True == bindings.HasFileReaderReadAsBinaryString(
		this.Ref(),
	)
}

// ReadAsBinaryStringFunc returns the method "FileReader.readAsBinaryString".
func (this FileReader) ReadAsBinaryStringFunc() (fn js.Func[func(blob Blob)]) {
	return fn.FromRef(
		bindings.FileReaderReadAsBinaryStringFunc(
			this.Ref(),
		),
	)
}

// ReadAsBinaryString calls the method "FileReader.readAsBinaryString".
func (this FileReader) ReadAsBinaryString(blob Blob) (ret js.Void) {
	bindings.CallFileReaderReadAsBinaryString(
		this.Ref(), js.Pointer(&ret),
		blob.Ref(),
	)

	return
}

// TryReadAsBinaryString calls the method "FileReader.readAsBinaryString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileReader) TryReadAsBinaryString(blob Blob) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileReaderReadAsBinaryString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		blob.Ref(),
	)

	return
}

// HasReadAsText returns true if the method "FileReader.readAsText" exists.
func (this FileReader) HasReadAsText() bool {
	return js.True == bindings.HasFileReaderReadAsText(
		this.Ref(),
	)
}

// ReadAsTextFunc returns the method "FileReader.readAsText".
func (this FileReader) ReadAsTextFunc() (fn js.Func[func(blob Blob, encoding js.String)]) {
	return fn.FromRef(
		bindings.FileReaderReadAsTextFunc(
			this.Ref(),
		),
	)
}

// ReadAsText calls the method "FileReader.readAsText".
func (this FileReader) ReadAsText(blob Blob, encoding js.String) (ret js.Void) {
	bindings.CallFileReaderReadAsText(
		this.Ref(), js.Pointer(&ret),
		blob.Ref(),
		encoding.Ref(),
	)

	return
}

// TryReadAsText calls the method "FileReader.readAsText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileReader) TryReadAsText(blob Blob, encoding js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileReaderReadAsText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		blob.Ref(),
		encoding.Ref(),
	)

	return
}

// HasReadAsText1 returns true if the method "FileReader.readAsText" exists.
func (this FileReader) HasReadAsText1() bool {
	return js.True == bindings.HasFileReaderReadAsText1(
		this.Ref(),
	)
}

// ReadAsText1Func returns the method "FileReader.readAsText".
func (this FileReader) ReadAsText1Func() (fn js.Func[func(blob Blob)]) {
	return fn.FromRef(
		bindings.FileReaderReadAsText1Func(
			this.Ref(),
		),
	)
}

// ReadAsText1 calls the method "FileReader.readAsText".
func (this FileReader) ReadAsText1(blob Blob) (ret js.Void) {
	bindings.CallFileReaderReadAsText1(
		this.Ref(), js.Pointer(&ret),
		blob.Ref(),
	)

	return
}

// TryReadAsText1 calls the method "FileReader.readAsText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileReader) TryReadAsText1(blob Blob) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileReaderReadAsText1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		blob.Ref(),
	)

	return
}

// HasReadAsDataURL returns true if the method "FileReader.readAsDataURL" exists.
func (this FileReader) HasReadAsDataURL() bool {
	return js.True == bindings.HasFileReaderReadAsDataURL(
		this.Ref(),
	)
}

// ReadAsDataURLFunc returns the method "FileReader.readAsDataURL".
func (this FileReader) ReadAsDataURLFunc() (fn js.Func[func(blob Blob)]) {
	return fn.FromRef(
		bindings.FileReaderReadAsDataURLFunc(
			this.Ref(),
		),
	)
}

// ReadAsDataURL calls the method "FileReader.readAsDataURL".
func (this FileReader) ReadAsDataURL(blob Blob) (ret js.Void) {
	bindings.CallFileReaderReadAsDataURL(
		this.Ref(), js.Pointer(&ret),
		blob.Ref(),
	)

	return
}

// TryReadAsDataURL calls the method "FileReader.readAsDataURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileReader) TryReadAsDataURL(blob Blob) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileReaderReadAsDataURL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		blob.Ref(),
	)

	return
}

// HasAbort returns true if the method "FileReader.abort" exists.
func (this FileReader) HasAbort() bool {
	return js.True == bindings.HasFileReaderAbort(
		this.Ref(),
	)
}

// AbortFunc returns the method "FileReader.abort".
func (this FileReader) AbortFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.FileReaderAbortFunc(
			this.Ref(),
		),
	)
}

// Abort calls the method "FileReader.abort".
func (this FileReader) Abort() (ret js.Void) {
	bindings.CallFileReaderAbort(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAbort calls the method "FileReader.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileReader) TryAbort() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileReaderAbort(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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

// HasReadAsArrayBuffer returns true if the method "FileReaderSync.readAsArrayBuffer" exists.
func (this FileReaderSync) HasReadAsArrayBuffer() bool {
	return js.True == bindings.HasFileReaderSyncReadAsArrayBuffer(
		this.Ref(),
	)
}

// ReadAsArrayBufferFunc returns the method "FileReaderSync.readAsArrayBuffer".
func (this FileReaderSync) ReadAsArrayBufferFunc() (fn js.Func[func(blob Blob) js.ArrayBuffer]) {
	return fn.FromRef(
		bindings.FileReaderSyncReadAsArrayBufferFunc(
			this.Ref(),
		),
	)
}

// ReadAsArrayBuffer calls the method "FileReaderSync.readAsArrayBuffer".
func (this FileReaderSync) ReadAsArrayBuffer(blob Blob) (ret js.ArrayBuffer) {
	bindings.CallFileReaderSyncReadAsArrayBuffer(
		this.Ref(), js.Pointer(&ret),
		blob.Ref(),
	)

	return
}

// TryReadAsArrayBuffer calls the method "FileReaderSync.readAsArrayBuffer"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileReaderSync) TryReadAsArrayBuffer(blob Blob) (ret js.ArrayBuffer, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileReaderSyncReadAsArrayBuffer(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		blob.Ref(),
	)

	return
}

// HasReadAsBinaryString returns true if the method "FileReaderSync.readAsBinaryString" exists.
func (this FileReaderSync) HasReadAsBinaryString() bool {
	return js.True == bindings.HasFileReaderSyncReadAsBinaryString(
		this.Ref(),
	)
}

// ReadAsBinaryStringFunc returns the method "FileReaderSync.readAsBinaryString".
func (this FileReaderSync) ReadAsBinaryStringFunc() (fn js.Func[func(blob Blob) js.String]) {
	return fn.FromRef(
		bindings.FileReaderSyncReadAsBinaryStringFunc(
			this.Ref(),
		),
	)
}

// ReadAsBinaryString calls the method "FileReaderSync.readAsBinaryString".
func (this FileReaderSync) ReadAsBinaryString(blob Blob) (ret js.String) {
	bindings.CallFileReaderSyncReadAsBinaryString(
		this.Ref(), js.Pointer(&ret),
		blob.Ref(),
	)

	return
}

// TryReadAsBinaryString calls the method "FileReaderSync.readAsBinaryString"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileReaderSync) TryReadAsBinaryString(blob Blob) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileReaderSyncReadAsBinaryString(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		blob.Ref(),
	)

	return
}

// HasReadAsText returns true if the method "FileReaderSync.readAsText" exists.
func (this FileReaderSync) HasReadAsText() bool {
	return js.True == bindings.HasFileReaderSyncReadAsText(
		this.Ref(),
	)
}

// ReadAsTextFunc returns the method "FileReaderSync.readAsText".
func (this FileReaderSync) ReadAsTextFunc() (fn js.Func[func(blob Blob, encoding js.String) js.String]) {
	return fn.FromRef(
		bindings.FileReaderSyncReadAsTextFunc(
			this.Ref(),
		),
	)
}

// ReadAsText calls the method "FileReaderSync.readAsText".
func (this FileReaderSync) ReadAsText(blob Blob, encoding js.String) (ret js.String) {
	bindings.CallFileReaderSyncReadAsText(
		this.Ref(), js.Pointer(&ret),
		blob.Ref(),
		encoding.Ref(),
	)

	return
}

// TryReadAsText calls the method "FileReaderSync.readAsText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileReaderSync) TryReadAsText(blob Blob, encoding js.String) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileReaderSyncReadAsText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		blob.Ref(),
		encoding.Ref(),
	)

	return
}

// HasReadAsText1 returns true if the method "FileReaderSync.readAsText" exists.
func (this FileReaderSync) HasReadAsText1() bool {
	return js.True == bindings.HasFileReaderSyncReadAsText1(
		this.Ref(),
	)
}

// ReadAsText1Func returns the method "FileReaderSync.readAsText".
func (this FileReaderSync) ReadAsText1Func() (fn js.Func[func(blob Blob) js.String]) {
	return fn.FromRef(
		bindings.FileReaderSyncReadAsText1Func(
			this.Ref(),
		),
	)
}

// ReadAsText1 calls the method "FileReaderSync.readAsText".
func (this FileReaderSync) ReadAsText1(blob Blob) (ret js.String) {
	bindings.CallFileReaderSyncReadAsText1(
		this.Ref(), js.Pointer(&ret),
		blob.Ref(),
	)

	return
}

// TryReadAsText1 calls the method "FileReaderSync.readAsText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileReaderSync) TryReadAsText1(blob Blob) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileReaderSyncReadAsText1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		blob.Ref(),
	)

	return
}

// HasReadAsDataURL returns true if the method "FileReaderSync.readAsDataURL" exists.
func (this FileReaderSync) HasReadAsDataURL() bool {
	return js.True == bindings.HasFileReaderSyncReadAsDataURL(
		this.Ref(),
	)
}

// ReadAsDataURLFunc returns the method "FileReaderSync.readAsDataURL".
func (this FileReaderSync) ReadAsDataURLFunc() (fn js.Func[func(blob Blob) js.String]) {
	return fn.FromRef(
		bindings.FileReaderSyncReadAsDataURLFunc(
			this.Ref(),
		),
	)
}

// ReadAsDataURL calls the method "FileReaderSync.readAsDataURL".
func (this FileReaderSync) ReadAsDataURL(blob Blob) (ret js.String) {
	bindings.CallFileReaderSyncReadAsDataURL(
		this.Ref(), js.Pointer(&ret),
		blob.Ref(),
	)

	return
}

// TryReadAsDataURL calls the method "FileReaderSync.readAsDataURL"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileReaderSync) TryReadAsDataURL(blob Blob) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileReaderSyncReadAsDataURL(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		blob.Ref(),
	)

	return
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

// HasFile returns true if the method "FileSystemFileEntry.file" exists.
func (this FileSystemFileEntry) HasFile() bool {
	return js.True == bindings.HasFileSystemFileEntryFile(
		this.Ref(),
	)
}

// FileFunc returns the method "FileSystemFileEntry.file".
func (this FileSystemFileEntry) FileFunc() (fn js.Func[func(successCallback js.Func[func(file File)], errorCallback js.Func[func(err DOMException)])]) {
	return fn.FromRef(
		bindings.FileSystemFileEntryFileFunc(
			this.Ref(),
		),
	)
}

// File calls the method "FileSystemFileEntry.file".
func (this FileSystemFileEntry) File(successCallback js.Func[func(file File)], errorCallback js.Func[func(err DOMException)]) (ret js.Void) {
	bindings.CallFileSystemFileEntryFile(
		this.Ref(), js.Pointer(&ret),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// TryFile calls the method "FileSystemFileEntry.file"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemFileEntry) TryFile(successCallback js.Func[func(file File)], errorCallback js.Func[func(err DOMException)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemFileEntryFile(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
		errorCallback.Ref(),
	)

	return
}

// HasFile1 returns true if the method "FileSystemFileEntry.file" exists.
func (this FileSystemFileEntry) HasFile1() bool {
	return js.True == bindings.HasFileSystemFileEntryFile1(
		this.Ref(),
	)
}

// File1Func returns the method "FileSystemFileEntry.file".
func (this FileSystemFileEntry) File1Func() (fn js.Func[func(successCallback js.Func[func(file File)])]) {
	return fn.FromRef(
		bindings.FileSystemFileEntryFile1Func(
			this.Ref(),
		),
	)
}

// File1 calls the method "FileSystemFileEntry.file".
func (this FileSystemFileEntry) File1(successCallback js.Func[func(file File)]) (ret js.Void) {
	bindings.CallFileSystemFileEntryFile1(
		this.Ref(), js.Pointer(&ret),
		successCallback.Ref(),
	)

	return
}

// TryFile1 calls the method "FileSystemFileEntry.file"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemFileEntry) TryFile1(successCallback js.Func[func(file File)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemFileEntryFile1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		successCallback.Ref(),
	)

	return
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

func NewFocusEvent(typ js.String, eventInitDict FocusEventInit) (ret FocusEvent) {
	ret.ref = bindings.NewFocusEventByFocusEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewFocusEventByFocusEvent1(typ js.String) (ret FocusEvent) {
	ret.ref = bindings.NewFocusEventByFocusEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this FocusEvent) RelatedTarget() (ret EventTarget, ok bool) {
	ok = js.True == bindings.GetFocusEventRelatedTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewFontFaceSetLoadEvent(typ js.String, eventInitDict FontFaceSetLoadEventInit) (ret FontFaceSetLoadEvent) {
	ret.ref = bindings.NewFontFaceSetLoadEventByFontFaceSetLoadEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewFontFaceSetLoadEventByFontFaceSetLoadEvent1(typ js.String) (ret FontFaceSetLoadEvent) {
	ret.ref = bindings.NewFontFaceSetLoadEventByFontFaceSetLoadEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this FontFaceSetLoadEvent) Fontfaces() (ret js.FrozenArray[FontFace], ok bool) {
	ok = js.True == bindings.GetFontFaceSetLoadEventFontfaces(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
// It returns ok=false if there is no such property.
func (this FontFaceVariationAxis) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceVariationAxisName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AxisTag returns the value of property "FontFaceVariationAxis.axisTag".
//
// It returns ok=false if there is no such property.
func (this FontFaceVariationAxis) AxisTag() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFontFaceVariationAxisAxisTag(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MinimumValue returns the value of property "FontFaceVariationAxis.minimumValue".
//
// It returns ok=false if there is no such property.
func (this FontFaceVariationAxis) MinimumValue() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontFaceVariationAxisMinimumValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaximumValue returns the value of property "FontFaceVariationAxis.maximumValue".
//
// It returns ok=false if there is no such property.
func (this FontFaceVariationAxis) MaximumValue() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontFaceVariationAxisMaximumValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DefaultValue returns the value of property "FontFaceVariationAxis.defaultValue".
//
// It returns ok=false if there is no such property.
func (this FontFaceVariationAxis) DefaultValue() (ret float64, ok bool) {
	ok = js.True == bindings.GetFontFaceVariationAxisDefaultValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewFormDataEvent(typ js.String, eventInitDict FormDataEventInit) (ret FormDataEvent) {
	ret.ref = bindings.NewFormDataEventByFormDataEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
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
// It returns ok=false if there is no such property.
func (this FormDataEvent) FormData() (ret FormData, ok bool) {
	ok = js.True == bindings.GetFormDataEventFormData(
		this.Ref(), js.Pointer(&ret),
	)
	return
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
	//
	// NOTE: BreakToken.FFI_USE MUST be set to true to get BreakToken used.
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

func NewFragmentResult(options FragmentResultOptions) (ret FragmentResult) {
	ret.ref = bindings.NewFragmentResultByFragmentResult(
		js.Pointer(&options))
	return
}

func NewFragmentResultByFragmentResult1() (ret FragmentResult) {
	ret.ref = bindings.NewFragmentResultByFragmentResult1()
	return
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
// It returns ok=false if there is no such property.
func (this FragmentResult) InlineSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetFragmentResultInlineSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BlockSize returns the value of property "FragmentResult.blockSize".
//
// It returns ok=false if there is no such property.
func (this FragmentResult) BlockSize() (ret float64, ok bool) {
	ok = js.True == bindings.GetFragmentResultBlockSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

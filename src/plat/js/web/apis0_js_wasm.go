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

type GLenum uint32

const (
	ANGLE_instanced_arrays_VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE GLenum = 0x88FE
)

type GLint int32

type GLsizei int32

type GLintptr int64

type GLuint uint32

type ANGLE_instanced_arrays struct {
	ref js.Ref
}

func (this ANGLE_instanced_arrays) Once() ANGLE_instanced_arrays {
	this.Ref().Once()
	return this
}

func (this ANGLE_instanced_arrays) Ref() js.Ref {
	return this.ref
}

func (this ANGLE_instanced_arrays) FromRef(ref js.Ref) ANGLE_instanced_arrays {
	this.ref = ref
	return this
}

func (this ANGLE_instanced_arrays) Free() {
	this.Ref().Free()
}

// DrawArraysInstancedANGLE calls the method "ANGLE_instanced_arrays.drawArraysInstancedANGLE".
//
// The returned bool will be false if there is no such method.
func (this ANGLE_instanced_arrays) DrawArraysInstancedANGLE(mode GLenum, first GLint, count GLsizei, primcount GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallANGLE_instanced_arraysDrawArraysInstancedANGLE(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		int32(first),
		int32(count),
		int32(primcount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawArraysInstancedANGLEFunc returns the method "ANGLE_instanced_arrays.drawArraysInstancedANGLE".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ANGLE_instanced_arrays) DrawArraysInstancedANGLEFunc() (fn js.Func[func(mode GLenum, first GLint, count GLsizei, primcount GLsizei)]) {
	return fn.FromRef(
		bindings.ANGLE_instanced_arraysDrawArraysInstancedANGLEFunc(
			this.Ref(),
		),
	)
}

// DrawElementsInstancedANGLE calls the method "ANGLE_instanced_arrays.drawElementsInstancedANGLE".
//
// The returned bool will be false if there is no such method.
func (this ANGLE_instanced_arrays) DrawElementsInstancedANGLE(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, primcount GLsizei) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallANGLE_instanced_arraysDrawElementsInstancedANGLE(
		this.Ref(), js.Pointer(&_ok),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
		int32(primcount),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DrawElementsInstancedANGLEFunc returns the method "ANGLE_instanced_arrays.drawElementsInstancedANGLE".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ANGLE_instanced_arrays) DrawElementsInstancedANGLEFunc() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, primcount GLsizei)]) {
	return fn.FromRef(
		bindings.ANGLE_instanced_arraysDrawElementsInstancedANGLEFunc(
			this.Ref(),
		),
	)
}

// VertexAttribDivisorANGLE calls the method "ANGLE_instanced_arrays.vertexAttribDivisorANGLE".
//
// The returned bool will be false if there is no such method.
func (this ANGLE_instanced_arrays) VertexAttribDivisorANGLE(index GLuint, divisor GLuint) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallANGLE_instanced_arraysVertexAttribDivisorANGLE(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		uint32(divisor),
	)

	_ = _ret
	return js.Void{}, _ok
}

// VertexAttribDivisorANGLEFunc returns the method "ANGLE_instanced_arrays.vertexAttribDivisorANGLE".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this ANGLE_instanced_arrays) VertexAttribDivisorANGLEFunc() (fn js.Func[func(index GLuint, divisor GLuint)]) {
	return fn.FromRef(
		bindings.ANGLE_instanced_arraysVertexAttribDivisorANGLEFunc(
			this.Ref(),
		),
	)
}

type AV1EncoderConfig struct {
	// ForceScreenContentTools is "AV1EncoderConfig.forceScreenContentTools"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ForceScreenContentTools MUST be set to true to make this field effective.
	ForceScreenContentTools bool

	FFI_USE_ForceScreenContentTools bool // for ForceScreenContentTools.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AV1EncoderConfig with all fields set.
func (p AV1EncoderConfig) FromRef(ref js.Ref) AV1EncoderConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AV1EncoderConfig in the application heap.
func (p AV1EncoderConfig) New() js.Ref {
	return bindings.AV1EncoderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AV1EncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.AV1EncoderConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AV1EncoderConfig) Update(ref js.Ref) {
	bindings.AV1EncoderConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type AacBitstreamFormat uint32

const (
	_ AacBitstreamFormat = iota

	AacBitstreamFormat_AAC
	AacBitstreamFormat_ADTS
)

func (AacBitstreamFormat) FromRef(str js.Ref) AacBitstreamFormat {
	return AacBitstreamFormat(bindings.ConstOfAacBitstreamFormat(str))
}

func (x AacBitstreamFormat) String() (string, bool) {
	switch x {
	case AacBitstreamFormat_AAC:
		return "aac", true
	case AacBitstreamFormat_ADTS:
		return "adts", true
	default:
		return "", false
	}
}

type AacEncoderConfig struct {
	// Format is "AacEncoderConfig.format"
	//
	// Optional, defaults to "aac".
	Format AacBitstreamFormat

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AacEncoderConfig with all fields set.
func (p AacEncoderConfig) FromRef(ref js.Ref) AacEncoderConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AacEncoderConfig in the application heap.
func (p AacEncoderConfig) New() js.Ref {
	return bindings.AacEncoderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AacEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.AacEncoderConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AacEncoderConfig) Update(ref js.Ref) {
	bindings.AacEncoderConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type EventHandlerNonNullFunc func(this js.Ref, event Event) js.Ref

func (fn EventHandlerNonNullFunc) Register() js.Func[func(event Event) js.Any] {
	return js.RegisterCallback[func(event Event) js.Any](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EventHandlerNonNullFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		Event{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EventHandlerNonNull[T any] struct {
	Fn  func(arg T, this js.Ref, event Event) js.Ref
	Arg T
}

func (cb *EventHandlerNonNull[T]) Register() js.Func[func(event Event) js.Any] {
	return js.RegisterCallback[func(event Event) js.Any](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EventHandlerNonNull[T]) DispatchCallback(
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

		Event{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

const (
	Event_NONE            uint16 = 0
	Event_CAPTURING_PHASE uint16 = 1
	Event_AT_TARGET       uint16 = 2
	Event_BUBBLING_PHASE  uint16 = 3
)

type EventInit struct {
	// Bubbles is "EventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "EventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "EventInit.composed"
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

// FromRef calls UpdateFrom and returns a EventInit with all fields set.
func (p EventInit) FromRef(ref js.Ref) EventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EventInit in the application heap.
func (p EventInit) New() js.Ref {
	return bindings.EventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p EventInit) UpdateFrom(ref js.Ref) {
	bindings.EventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p EventInit) Update(ref js.Ref) {
	bindings.EventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type EventListenerFunc func(this js.Ref, event Event) js.Ref

func (fn EventListenerFunc) Register() js.Func[func(event Event)] {
	return js.RegisterCallback[func(event Event)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn EventListenerFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		Event{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type EventListener[T any] struct {
	Fn  func(arg T, this js.Ref, event Event) js.Ref
	Arg T
}

func (cb *EventListener[T]) Register() js.Func[func(event Event)] {
	return js.RegisterCallback[func(event Event)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *EventListener[T]) DispatchCallback(
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

		Event{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AddEventListenerOptions struct {
	// Passive is "AddEventListenerOptions.passive"
	//
	// Optional
	//
	// NOTE: FFI_USE_Passive MUST be set to true to make this field effective.
	Passive bool
	// Once is "AddEventListenerOptions.once"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Once MUST be set to true to make this field effective.
	Once bool
	// Signal is "AddEventListenerOptions.signal"
	//
	// Optional
	Signal AbortSignal
	// Capture is "AddEventListenerOptions.capture"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Capture MUST be set to true to make this field effective.
	Capture bool

	FFI_USE_Passive bool // for Passive.
	FFI_USE_Once    bool // for Once.
	FFI_USE_Capture bool // for Capture.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AddEventListenerOptions with all fields set.
func (p AddEventListenerOptions) FromRef(ref js.Ref) AddEventListenerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AddEventListenerOptions in the application heap.
func (p AddEventListenerOptions) New() js.Ref {
	return bindings.AddEventListenerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AddEventListenerOptions) UpdateFrom(ref js.Ref) {
	bindings.AddEventListenerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AddEventListenerOptions) Update(ref js.Ref) {
	bindings.AddEventListenerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_AddEventListenerOptions_Bool struct {
	ref js.Ref
}

func (x OneOf_AddEventListenerOptions_Bool) Ref() js.Ref {
	return x.ref
}

func (x OneOf_AddEventListenerOptions_Bool) Free() {
	x.ref.Free()
}

func (x OneOf_AddEventListenerOptions_Bool) FromRef(ref js.Ref) OneOf_AddEventListenerOptions_Bool {
	return OneOf_AddEventListenerOptions_Bool{
		ref: ref,
	}
}

func (x OneOf_AddEventListenerOptions_Bool) AddEventListenerOptions() AddEventListenerOptions {
	var ret AddEventListenerOptions
	ret.UpdateFrom(x.ref)
	return ret
}

func (x OneOf_AddEventListenerOptions_Bool) Bool() bool {
	return x.ref == js.True
}

type EventListenerOptions struct {
	// Capture is "EventListenerOptions.capture"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Capture MUST be set to true to make this field effective.
	Capture bool

	FFI_USE_Capture bool // for Capture.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EventListenerOptions with all fields set.
func (p EventListenerOptions) FromRef(ref js.Ref) EventListenerOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EventListenerOptions in the application heap.
func (p EventListenerOptions) New() js.Ref {
	return bindings.EventListenerOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p EventListenerOptions) UpdateFrom(ref js.Ref) {
	bindings.EventListenerOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p EventListenerOptions) Update(ref js.Ref) {
	bindings.EventListenerOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_EventListenerOptions_Bool struct {
	ref js.Ref
}

func (x OneOf_EventListenerOptions_Bool) Ref() js.Ref {
	return x.ref
}

func (x OneOf_EventListenerOptions_Bool) Free() {
	x.ref.Free()
}

func (x OneOf_EventListenerOptions_Bool) FromRef(ref js.Ref) OneOf_EventListenerOptions_Bool {
	return OneOf_EventListenerOptions_Bool{
		ref: ref,
	}
}

func (x OneOf_EventListenerOptions_Bool) EventListenerOptions() EventListenerOptions {
	var ret EventListenerOptions
	ret.UpdateFrom(x.ref)
	return ret
}

func (x OneOf_EventListenerOptions_Bool) Bool() bool {
	return x.ref == js.True
}

type EventTarget struct {
	ref js.Ref
}

func (this EventTarget) Once() EventTarget {
	this.Ref().Once()
	return this
}

func (this EventTarget) Ref() js.Ref {
	return this.ref
}

func (this EventTarget) FromRef(ref js.Ref) EventTarget {
	this.ref = ref
	return this
}

func (this EventTarget) Free() {
	this.Ref().Free()
}

// AddEventListener calls the method "EventTarget.addEventListener".
//
// The returned bool will be false if there is no such method.
func (this EventTarget) AddEventListener(typ js.String, callback js.Func[func(event Event)], options OneOf_AddEventListenerOptions_Bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEventTargetAddEventListener(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		callback.Ref(),
		options.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddEventListenerFunc returns the method "EventTarget.addEventListener".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EventTarget) AddEventListenerFunc() (fn js.Func[func(typ js.String, callback js.Func[func(event Event)], options OneOf_AddEventListenerOptions_Bool)]) {
	return fn.FromRef(
		bindings.EventTargetAddEventListenerFunc(
			this.Ref(),
		),
	)
}

// AddEventListener1 calls the method "EventTarget.addEventListener".
//
// The returned bool will be false if there is no such method.
func (this EventTarget) AddEventListener1(typ js.String, callback js.Func[func(event Event)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEventTargetAddEventListener1(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		callback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddEventListener1Func returns the method "EventTarget.addEventListener".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EventTarget) AddEventListener1Func() (fn js.Func[func(typ js.String, callback js.Func[func(event Event)])]) {
	return fn.FromRef(
		bindings.EventTargetAddEventListener1Func(
			this.Ref(),
		),
	)
}

// RemoveEventListener calls the method "EventTarget.removeEventListener".
//
// The returned bool will be false if there is no such method.
func (this EventTarget) RemoveEventListener(typ js.String, callback js.Func[func(event Event)], options OneOf_EventListenerOptions_Bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEventTargetRemoveEventListener(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		callback.Ref(),
		options.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveEventListenerFunc returns the method "EventTarget.removeEventListener".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EventTarget) RemoveEventListenerFunc() (fn js.Func[func(typ js.String, callback js.Func[func(event Event)], options OneOf_EventListenerOptions_Bool)]) {
	return fn.FromRef(
		bindings.EventTargetRemoveEventListenerFunc(
			this.Ref(),
		),
	)
}

// RemoveEventListener1 calls the method "EventTarget.removeEventListener".
//
// The returned bool will be false if there is no such method.
func (this EventTarget) RemoveEventListener1(typ js.String, callback js.Func[func(event Event)]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEventTargetRemoveEventListener1(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		callback.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveEventListener1Func returns the method "EventTarget.removeEventListener".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EventTarget) RemoveEventListener1Func() (fn js.Func[func(typ js.String, callback js.Func[func(event Event)])]) {
	return fn.FromRef(
		bindings.EventTargetRemoveEventListener1Func(
			this.Ref(),
		),
	)
}

// DispatchEvent calls the method "EventTarget.dispatchEvent".
//
// The returned bool will be false if there is no such method.
func (this EventTarget) DispatchEvent(event Event) (bool, bool) {
	var _ok bool
	_ret := bindings.CallEventTargetDispatchEvent(
		this.Ref(), js.Pointer(&_ok),
		event.Ref(),
	)

	return _ret == js.True, _ok
}

// DispatchEventFunc returns the method "EventTarget.dispatchEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this EventTarget) DispatchEventFunc() (fn js.Func[func(event Event) bool]) {
	return fn.FromRef(
		bindings.EventTargetDispatchEventFunc(
			this.Ref(),
		),
	)
}

type DOMHighResTimeStamp float64

func NewEvent(typ js.String, eventInitDict EventInit) Event {
	return Event{}.FromRef(
		bindings.NewEventByEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewEventByEvent1(typ js.String) Event {
	return Event{}.FromRef(
		bindings.NewEventByEvent1(
			typ.Ref()),
	)
}

type Event struct {
	ref js.Ref
}

func (this Event) Once() Event {
	this.Ref().Once()
	return this
}

func (this Event) Ref() js.Ref {
	return this.ref
}

func (this Event) FromRef(ref js.Ref) Event {
	this.ref = ref
	return this
}

func (this Event) Free() {
	this.Ref().Free()
}

// Type returns the value of property "Event.type".
//
// The returned bool will be false if there is no such property.
func (this Event) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetEventType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Target returns the value of property "Event.target".
//
// The returned bool will be false if there is no such property.
func (this Event) Target() (EventTarget, bool) {
	var _ok bool
	_ret := bindings.GetEventTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return EventTarget{}.FromRef(_ret), _ok
}

// SrcElement returns the value of property "Event.srcElement".
//
// The returned bool will be false if there is no such property.
func (this Event) SrcElement() (EventTarget, bool) {
	var _ok bool
	_ret := bindings.GetEventSrcElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return EventTarget{}.FromRef(_ret), _ok
}

// CurrentTarget returns the value of property "Event.currentTarget".
//
// The returned bool will be false if there is no such property.
func (this Event) CurrentTarget() (EventTarget, bool) {
	var _ok bool
	_ret := bindings.GetEventCurrentTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return EventTarget{}.FromRef(_ret), _ok
}

// EventPhase returns the value of property "Event.eventPhase".
//
// The returned bool will be false if there is no such property.
func (this Event) EventPhase() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetEventEventPhase(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// CancelBubble returns the value of property "Event.cancelBubble".
//
// The returned bool will be false if there is no such property.
func (this Event) CancelBubble() (bool, bool) {
	var _ok bool
	_ret := bindings.GetEventCancelBubble(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// CancelBubble sets the value of property "Event.cancelBubble" to val.
//
// It returns false if the property cannot be set.
func (this Event) SetCancelBubble(val bool) bool {
	return js.True == bindings.SetEventCancelBubble(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Bubbles returns the value of property "Event.bubbles".
//
// The returned bool will be false if there is no such property.
func (this Event) Bubbles() (bool, bool) {
	var _ok bool
	_ret := bindings.GetEventBubbles(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Cancelable returns the value of property "Event.cancelable".
//
// The returned bool will be false if there is no such property.
func (this Event) Cancelable() (bool, bool) {
	var _ok bool
	_ret := bindings.GetEventCancelable(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ReturnValue returns the value of property "Event.returnValue".
//
// The returned bool will be false if there is no such property.
func (this Event) ReturnValue() (bool, bool) {
	var _ok bool
	_ret := bindings.GetEventReturnValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ReturnValue sets the value of property "Event.returnValue" to val.
//
// It returns false if the property cannot be set.
func (this Event) SetReturnValue(val bool) bool {
	return js.True == bindings.SetEventReturnValue(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// DefaultPrevented returns the value of property "Event.defaultPrevented".
//
// The returned bool will be false if there is no such property.
func (this Event) DefaultPrevented() (bool, bool) {
	var _ok bool
	_ret := bindings.GetEventDefaultPrevented(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Composed returns the value of property "Event.composed".
//
// The returned bool will be false if there is no such property.
func (this Event) Composed() (bool, bool) {
	var _ok bool
	_ret := bindings.GetEventComposed(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IsTrusted returns the value of property "Event.isTrusted".
//
// The returned bool will be false if there is no such property.
func (this Event) IsTrusted() (bool, bool) {
	var _ok bool
	_ret := bindings.GetEventIsTrusted(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// TimeStamp returns the value of property "Event.timeStamp".
//
// The returned bool will be false if there is no such property.
func (this Event) TimeStamp() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetEventTimeStamp(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// ComposedPath calls the method "Event.composedPath".
//
// The returned bool will be false if there is no such method.
func (this Event) ComposedPath() (js.Array[EventTarget], bool) {
	var _ok bool
	_ret := bindings.CallEventComposedPath(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[EventTarget]{}.FromRef(_ret), _ok
}

// ComposedPathFunc returns the method "Event.composedPath".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Event) ComposedPathFunc() (fn js.Func[func() js.Array[EventTarget]]) {
	return fn.FromRef(
		bindings.EventComposedPathFunc(
			this.Ref(),
		),
	)
}

// StopPropagation calls the method "Event.stopPropagation".
//
// The returned bool will be false if there is no such method.
func (this Event) StopPropagation() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEventStopPropagation(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StopPropagationFunc returns the method "Event.stopPropagation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Event) StopPropagationFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.EventStopPropagationFunc(
			this.Ref(),
		),
	)
}

// StopImmediatePropagation calls the method "Event.stopImmediatePropagation".
//
// The returned bool will be false if there is no such method.
func (this Event) StopImmediatePropagation() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEventStopImmediatePropagation(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StopImmediatePropagationFunc returns the method "Event.stopImmediatePropagation".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Event) StopImmediatePropagationFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.EventStopImmediatePropagationFunc(
			this.Ref(),
		),
	)
}

// PreventDefault calls the method "Event.preventDefault".
//
// The returned bool will be false if there is no such method.
func (this Event) PreventDefault() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEventPreventDefault(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PreventDefaultFunc returns the method "Event.preventDefault".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Event) PreventDefaultFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.EventPreventDefaultFunc(
			this.Ref(),
		),
	)
}

// InitEvent calls the method "Event.initEvent".
//
// The returned bool will be false if there is no such method.
func (this Event) InitEvent(typ js.String, bubbles bool, cancelable bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEventInitEvent(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitEventFunc returns the method "Event.initEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Event) InitEventFunc() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool)]) {
	return fn.FromRef(
		bindings.EventInitEventFunc(
			this.Ref(),
		),
	)
}

// InitEvent1 calls the method "Event.initEvent".
//
// The returned bool will be false if there is no such method.
func (this Event) InitEvent1(typ js.String, bubbles bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEventInitEvent1(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitEvent1Func returns the method "Event.initEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Event) InitEvent1Func() (fn js.Func[func(typ js.String, bubbles bool)]) {
	return fn.FromRef(
		bindings.EventInitEvent1Func(
			this.Ref(),
		),
	)
}

// InitEvent2 calls the method "Event.initEvent".
//
// The returned bool will be false if there is no such method.
func (this Event) InitEvent2(typ js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallEventInitEvent2(
		this.Ref(), js.Pointer(&_ok),
		typ.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitEvent2Func returns the method "Event.initEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Event) InitEvent2Func() (fn js.Func[func(typ js.String)]) {
	return fn.FromRef(
		bindings.EventInitEvent2Func(
			this.Ref(),
		),
	)
}

type EventHandler = js.Func[func(event Event) js.Any]

type AbortSignal struct {
	EventTarget
}

func (this AbortSignal) Once() AbortSignal {
	this.Ref().Once()
	return this
}

func (this AbortSignal) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this AbortSignal) FromRef(ref js.Ref) AbortSignal {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this AbortSignal) Free() {
	this.Ref().Free()
}

// Aborted returns the value of property "AbortSignal.aborted".
//
// The returned bool will be false if there is no such property.
func (this AbortSignal) Aborted() (bool, bool) {
	var _ok bool
	_ret := bindings.GetAbortSignalAborted(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Reason returns the value of property "AbortSignal.reason".
//
// The returned bool will be false if there is no such property.
func (this AbortSignal) Reason() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetAbortSignalReason(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Abort calls the staticmethod "AbortSignal.abort".
//
// The returned bool will be false if there is no such method.
func (this AbortSignal) Abort(reason js.Any) (AbortSignal, bool) {
	var _ok bool
	_ret := bindings.CallAbortSignalAbort(
		this.Ref(), js.Pointer(&_ok),
		reason.Ref(),
	)

	return AbortSignal{}.FromRef(_ret), _ok
}

// AbortFunc returns the staticmethod "AbortSignal.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AbortSignal) AbortFunc() (fn js.Func[func(reason js.Any) AbortSignal]) {
	return fn.FromRef(
		bindings.AbortSignalAbortFunc(
			this.Ref(),
		),
	)
}

// Abort1 calls the staticmethod "AbortSignal.abort".
//
// The returned bool will be false if there is no such method.
func (this AbortSignal) Abort1() (AbortSignal, bool) {
	var _ok bool
	_ret := bindings.CallAbortSignalAbort1(
		this.Ref(), js.Pointer(&_ok),
	)

	return AbortSignal{}.FromRef(_ret), _ok
}

// Abort1Func returns the staticmethod "AbortSignal.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AbortSignal) Abort1Func() (fn js.Func[func() AbortSignal]) {
	return fn.FromRef(
		bindings.AbortSignalAbort1Func(
			this.Ref(),
		),
	)
}

// Timeout calls the staticmethod "AbortSignal.timeout".
//
// The returned bool will be false if there is no such method.
func (this AbortSignal) Timeout(milliseconds uint64) (AbortSignal, bool) {
	var _ok bool
	_ret := bindings.CallAbortSignalTimeout(
		this.Ref(), js.Pointer(&_ok),
		float64(milliseconds),
	)

	return AbortSignal{}.FromRef(_ret), _ok
}

// TimeoutFunc returns the staticmethod "AbortSignal.timeout".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AbortSignal) TimeoutFunc() (fn js.Func[func(milliseconds uint64) AbortSignal]) {
	return fn.FromRef(
		bindings.AbortSignalTimeoutFunc(
			this.Ref(),
		),
	)
}

// Any calls the staticmethod "AbortSignal.any".
//
// The returned bool will be false if there is no such method.
func (this AbortSignal) Any(signals js.Array[AbortSignal]) (AbortSignal, bool) {
	var _ok bool
	_ret := bindings.CallAbortSignalAny(
		this.Ref(), js.Pointer(&_ok),
		signals.Ref(),
	)

	return AbortSignal{}.FromRef(_ret), _ok
}

// AnyFunc returns the staticmethod "AbortSignal.any".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AbortSignal) AnyFunc() (fn js.Func[func(signals js.Array[AbortSignal]) AbortSignal]) {
	return fn.FromRef(
		bindings.AbortSignalAnyFunc(
			this.Ref(),
		),
	)
}

// ThrowIfAborted calls the method "AbortSignal.throwIfAborted".
//
// The returned bool will be false if there is no such method.
func (this AbortSignal) ThrowIfAborted() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAbortSignalThrowIfAborted(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ThrowIfAbortedFunc returns the method "AbortSignal.throwIfAborted".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AbortSignal) ThrowIfAbortedFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AbortSignalThrowIfAbortedFunc(
			this.Ref(),
		),
	)
}

type AbortController struct {
	ref js.Ref
}

func (this AbortController) Once() AbortController {
	this.Ref().Once()
	return this
}

func (this AbortController) Ref() js.Ref {
	return this.ref
}

func (this AbortController) FromRef(ref js.Ref) AbortController {
	this.ref = ref
	return this
}

func (this AbortController) Free() {
	this.Ref().Free()
}

// Signal returns the value of property "AbortController.signal".
//
// The returned bool will be false if there is no such property.
func (this AbortController) Signal() (AbortSignal, bool) {
	var _ok bool
	_ret := bindings.GetAbortControllerSignal(
		this.Ref(), js.Pointer(&_ok),
	)
	return AbortSignal{}.FromRef(_ret), _ok
}

// Abort calls the method "AbortController.abort".
//
// The returned bool will be false if there is no such method.
func (this AbortController) Abort(reason js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAbortControllerAbort(
		this.Ref(), js.Pointer(&_ok),
		reason.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AbortFunc returns the method "AbortController.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AbortController) AbortFunc() (fn js.Func[func(reason js.Any)]) {
	return fn.FromRef(
		bindings.AbortControllerAbortFunc(
			this.Ref(),
		),
	)
}

// Abort1 calls the method "AbortController.abort".
//
// The returned bool will be false if there is no such method.
func (this AbortController) Abort1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallAbortControllerAbort1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Abort1Func returns the method "AbortController.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this AbortController) Abort1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.AbortControllerAbort1Func(
			this.Ref(),
		),
	)
}

type AbsoluteOrientationReadingValues struct {
	// Quaternion is "AbsoluteOrientationReadingValues.quaternion"
	//
	// Required
	Quaternion js.FrozenArray[float64]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AbsoluteOrientationReadingValues with all fields set.
func (p AbsoluteOrientationReadingValues) FromRef(ref js.Ref) AbsoluteOrientationReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AbsoluteOrientationReadingValues in the application heap.
func (p AbsoluteOrientationReadingValues) New() js.Ref {
	return bindings.AbsoluteOrientationReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p AbsoluteOrientationReadingValues) UpdateFrom(ref js.Ref) {
	bindings.AbsoluteOrientationReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p AbsoluteOrientationReadingValues) Update(ref js.Ref) {
	bindings.AbsoluteOrientationReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OrientationSensorLocalCoordinateSystem uint32

const (
	_ OrientationSensorLocalCoordinateSystem = iota

	OrientationSensorLocalCoordinateSystem_DEVICE
	OrientationSensorLocalCoordinateSystem_SCREEN
)

func (OrientationSensorLocalCoordinateSystem) FromRef(str js.Ref) OrientationSensorLocalCoordinateSystem {
	return OrientationSensorLocalCoordinateSystem(bindings.ConstOfOrientationSensorLocalCoordinateSystem(str))
}

func (x OrientationSensorLocalCoordinateSystem) String() (string, bool) {
	switch x {
	case OrientationSensorLocalCoordinateSystem_DEVICE:
		return "device", true
	case OrientationSensorLocalCoordinateSystem_SCREEN:
		return "screen", true
	default:
		return "", false
	}
}

type OrientationSensorOptions struct {
	// ReferenceFrame is "OrientationSensorOptions.referenceFrame"
	//
	// Optional, defaults to "device".
	ReferenceFrame OrientationSensorLocalCoordinateSystem
	// Frequency is "OrientationSensorOptions.frequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OrientationSensorOptions with all fields set.
func (p OrientationSensorOptions) FromRef(ref js.Ref) OrientationSensorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OrientationSensorOptions in the application heap.
func (p OrientationSensorOptions) New() js.Ref {
	return bindings.OrientationSensorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p OrientationSensorOptions) UpdateFrom(ref js.Ref) {
	bindings.OrientationSensorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p OrientationSensorOptions) Update(ref js.Ref) {
	bindings.OrientationSensorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewAbsoluteOrientationSensor(sensorOptions OrientationSensorOptions) AbsoluteOrientationSensor {
	return AbsoluteOrientationSensor{}.FromRef(
		bindings.NewAbsoluteOrientationSensorByAbsoluteOrientationSensor(
			js.Pointer(&sensorOptions)),
	)
}

func NewAbsoluteOrientationSensorByAbsoluteOrientationSensor1() AbsoluteOrientationSensor {
	return AbsoluteOrientationSensor{}.FromRef(
		bindings.NewAbsoluteOrientationSensorByAbsoluteOrientationSensor1(),
	)
}

type AbsoluteOrientationSensor struct {
	OrientationSensor
}

func (this AbsoluteOrientationSensor) Once() AbsoluteOrientationSensor {
	this.Ref().Once()
	return this
}

func (this AbsoluteOrientationSensor) Ref() js.Ref {
	return this.OrientationSensor.Ref()
}

func (this AbsoluteOrientationSensor) FromRef(ref js.Ref) AbsoluteOrientationSensor {
	this.OrientationSensor = this.OrientationSensor.FromRef(ref)
	return this
}

func (this AbsoluteOrientationSensor) Free() {
	this.Ref().Free()
}

const (
	Node_ELEMENT_NODE                              uint16 = 1
	Node_ATTRIBUTE_NODE                            uint16 = 2
	Node_TEXT_NODE                                 uint16 = 3
	Node_CDATA_SECTION_NODE                        uint16 = 4
	Node_ENTITY_REFERENCE_NODE                     uint16 = 5
	Node_ENTITY_NODE                               uint16 = 6
	Node_PROCESSING_INSTRUCTION_NODE               uint16 = 7
	Node_COMMENT_NODE                              uint16 = 8
	Node_DOCUMENT_NODE                             uint16 = 9
	Node_DOCUMENT_TYPE_NODE                        uint16 = 10
	Node_DOCUMENT_FRAGMENT_NODE                    uint16 = 11
	Node_NOTATION_NODE                             uint16 = 12
	Node_DOCUMENT_POSITION_DISCONNECTED            uint16 = 0x01
	Node_DOCUMENT_POSITION_PRECEDING               uint16 = 0x02
	Node_DOCUMENT_POSITION_FOLLOWING               uint16 = 0x04
	Node_DOCUMENT_POSITION_CONTAINS                uint16 = 0x08
	Node_DOCUMENT_POSITION_CONTAINED_BY            uint16 = 0x10
	Node_DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC uint16 = 0x20
)

type GetRootNodeOptions struct {
	// Composed is "GetRootNodeOptions.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Composed bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GetRootNodeOptions with all fields set.
func (p GetRootNodeOptions) FromRef(ref js.Ref) GetRootNodeOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GetRootNodeOptions in the application heap.
func (p GetRootNodeOptions) New() js.Ref {
	return bindings.GetRootNodeOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GetRootNodeOptions) UpdateFrom(ref js.Ref) {
	bindings.GetRootNodeOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GetRootNodeOptions) Update(ref js.Ref) {
	bindings.GetRootNodeOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type Attr struct {
	Node
}

func (this Attr) Once() Attr {
	this.Ref().Once()
	return this
}

func (this Attr) Ref() js.Ref {
	return this.Node.Ref()
}

func (this Attr) FromRef(ref js.Ref) Attr {
	this.Node = this.Node.FromRef(ref)
	return this
}

func (this Attr) Free() {
	this.Ref().Free()
}

// NamespaceURI returns the value of property "Attr.namespaceURI".
//
// The returned bool will be false if there is no such property.
func (this Attr) NamespaceURI() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAttrNamespaceURI(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Prefix returns the value of property "Attr.prefix".
//
// The returned bool will be false if there is no such property.
func (this Attr) Prefix() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAttrPrefix(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LocalName returns the value of property "Attr.localName".
//
// The returned bool will be false if there is no such property.
func (this Attr) LocalName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAttrLocalName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name returns the value of property "Attr.name".
//
// The returned bool will be false if there is no such property.
func (this Attr) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAttrName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value returns the value of property "Attr.value".
//
// The returned bool will be false if there is no such property.
func (this Attr) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetAttrValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "Attr.value" to val.
//
// It returns false if the property cannot be set.
func (this Attr) SetValue(val js.String) bool {
	return js.True == bindings.SetAttrValue(
		this.Ref(),
		val.Ref(),
	)
}

// OwnerElement returns the value of property "Attr.ownerElement".
//
// The returned bool will be false if there is no such property.
func (this Attr) OwnerElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetAttrOwnerElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// Specified returns the value of property "Attr.specified".
//
// The returned bool will be false if there is no such property.
func (this Attr) Specified() (bool, bool) {
	var _ok bool
	_ret := bindings.GetAttrSpecified(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

type FillMode uint32

const (
	_ FillMode = iota

	FillMode_NONE
	FillMode_FORWARDS
	FillMode_BACKWARDS
	FillMode_BOTH
	FillMode_AUTO
)

func (FillMode) FromRef(str js.Ref) FillMode {
	return FillMode(bindings.ConstOfFillMode(str))
}

func (x FillMode) String() (string, bool) {
	switch x {
	case FillMode_NONE:
		return "none", true
	case FillMode_FORWARDS:
		return "forwards", true
	case FillMode_BACKWARDS:
		return "backwards", true
	case FillMode_BOTH:
		return "both", true
	case FillMode_AUTO:
		return "auto", true
	default:
		return "", false
	}
}

type PlaybackDirection uint32

const (
	_ PlaybackDirection = iota

	PlaybackDirection_NORMAL
	PlaybackDirection_REVERSE
	PlaybackDirection_ALTERNATE
	PlaybackDirection_ALTERNATE_REVERSE
)

func (PlaybackDirection) FromRef(str js.Ref) PlaybackDirection {
	return PlaybackDirection(bindings.ConstOfPlaybackDirection(str))
}

func (x PlaybackDirection) String() (string, bool) {
	switch x {
	case PlaybackDirection_NORMAL:
		return "normal", true
	case PlaybackDirection_REVERSE:
		return "reverse", true
	case PlaybackDirection_ALTERNATE:
		return "alternate", true
	case PlaybackDirection_ALTERNATE_REVERSE:
		return "alternate-reverse", true
	default:
		return "", false
	}
}

type OneOf_Float64_CSSNumericValue struct {
	ref js.Ref
}

func (x OneOf_Float64_CSSNumericValue) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_CSSNumericValue) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_CSSNumericValue) FromRef(ref js.Ref) OneOf_Float64_CSSNumericValue {
	return OneOf_Float64_CSSNumericValue{
		ref: ref,
	}
}

func (x OneOf_Float64_CSSNumericValue) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_CSSNumericValue) CSSNumericValue() CSSNumericValue {
	return CSSNumericValue{}.FromRef(x.ref)
}

type CSSNumberish = OneOf_Float64_CSSNumericValue

func NewCSSUnitValue(value float64, unit js.String) CSSUnitValue {
	return CSSUnitValue{}.FromRef(
		bindings.NewCSSUnitValueByCSSUnitValue(
			float64(value),
			unit.Ref()),
	)
}

type CSSUnitValue struct {
	CSSNumericValue
}

func (this CSSUnitValue) Once() CSSUnitValue {
	this.Ref().Once()
	return this
}

func (this CSSUnitValue) Ref() js.Ref {
	return this.CSSNumericValue.Ref()
}

func (this CSSUnitValue) FromRef(ref js.Ref) CSSUnitValue {
	this.CSSNumericValue = this.CSSNumericValue.FromRef(ref)
	return this
}

func (this CSSUnitValue) Free() {
	this.Ref().Free()
}

// Value returns the value of property "CSSUnitValue.value".
//
// The returned bool will be false if there is no such property.
func (this CSSUnitValue) Value() (float64, bool) {
	var _ok bool
	_ret := bindings.GetCSSUnitValueValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Value sets the value of property "CSSUnitValue.value" to val.
//
// It returns false if the property cannot be set.
func (this CSSUnitValue) SetValue(val float64) bool {
	return js.True == bindings.SetCSSUnitValueValue(
		this.Ref(),
		float64(val),
	)
}

// Unit returns the value of property "CSSUnitValue.unit".
//
// The returned bool will be false if there is no such property.
func (this CSSUnitValue) Unit() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetCSSUnitValueUnit(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type CSSNumericArray struct {
	ref js.Ref
}

func (this CSSNumericArray) Once() CSSNumericArray {
	this.Ref().Once()
	return this
}

func (this CSSNumericArray) Ref() js.Ref {
	return this.ref
}

func (this CSSNumericArray) FromRef(ref js.Ref) CSSNumericArray {
	this.ref = ref
	return this
}

func (this CSSNumericArray) Free() {
	this.Ref().Free()
}

// Length returns the value of property "CSSNumericArray.length".
//
// The returned bool will be false if there is no such property.
func (this CSSNumericArray) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetCSSNumericArrayLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Get calls the method "CSSNumericArray.".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericArray) Get(index uint32) (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumericArrayGet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return CSSNumericValue{}.FromRef(_ret), _ok
}

// GetFunc returns the method "CSSNumericArray.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericArray) GetFunc() (fn js.Func[func(index uint32) CSSNumericValue]) {
	return fn.FromRef(
		bindings.CSSNumericArrayGetFunc(
			this.Ref(),
		),
	)
}

func NewCSSMathSum(args ...CSSNumberish) CSSMathSum {
	return CSSMathSum{}.FromRef(
		bindings.NewCSSMathSumByCSSMathSum(
			js.SliceData(args),
			js.SizeU(len(args))),
	)
}

type CSSMathSum struct {
	CSSMathValue
}

func (this CSSMathSum) Once() CSSMathSum {
	this.Ref().Once()
	return this
}

func (this CSSMathSum) Ref() js.Ref {
	return this.CSSMathValue.Ref()
}

func (this CSSMathSum) FromRef(ref js.Ref) CSSMathSum {
	this.CSSMathValue = this.CSSMathValue.FromRef(ref)
	return this
}

func (this CSSMathSum) Free() {
	this.Ref().Free()
}

// Values returns the value of property "CSSMathSum.values".
//
// The returned bool will be false if there is no such property.
func (this CSSMathSum) Values() (CSSNumericArray, bool) {
	var _ok bool
	_ret := bindings.GetCSSMathSumValues(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSNumericArray{}.FromRef(_ret), _ok
}

type CSSNumericBaseType uint32

const (
	_ CSSNumericBaseType = iota

	CSSNumericBaseType_LENGTH
	CSSNumericBaseType_ANGLE
	CSSNumericBaseType_TIME
	CSSNumericBaseType_FREQUENCY
	CSSNumericBaseType_RESOLUTION
	CSSNumericBaseType_FLEX
	CSSNumericBaseType_PERCENT
)

func (CSSNumericBaseType) FromRef(str js.Ref) CSSNumericBaseType {
	return CSSNumericBaseType(bindings.ConstOfCSSNumericBaseType(str))
}

func (x CSSNumericBaseType) String() (string, bool) {
	switch x {
	case CSSNumericBaseType_LENGTH:
		return "length", true
	case CSSNumericBaseType_ANGLE:
		return "angle", true
	case CSSNumericBaseType_TIME:
		return "time", true
	case CSSNumericBaseType_FREQUENCY:
		return "frequency", true
	case CSSNumericBaseType_RESOLUTION:
		return "resolution", true
	case CSSNumericBaseType_FLEX:
		return "flex", true
	case CSSNumericBaseType_PERCENT:
		return "percent", true
	default:
		return "", false
	}
}

type CSSNumericType struct {
	// Length is "CSSNumericType.length"
	//
	// Optional
	//
	// NOTE: FFI_USE_Length MUST be set to true to make this field effective.
	Length int32
	// Angle is "CSSNumericType.angle"
	//
	// Optional
	//
	// NOTE: FFI_USE_Angle MUST be set to true to make this field effective.
	Angle int32
	// Time is "CSSNumericType.time"
	//
	// Optional
	//
	// NOTE: FFI_USE_Time MUST be set to true to make this field effective.
	Time int32
	// Frequency is "CSSNumericType.frequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency int32
	// Resolution is "CSSNumericType.resolution"
	//
	// Optional
	//
	// NOTE: FFI_USE_Resolution MUST be set to true to make this field effective.
	Resolution int32
	// Flex is "CSSNumericType.flex"
	//
	// Optional
	//
	// NOTE: FFI_USE_Flex MUST be set to true to make this field effective.
	Flex int32
	// Percent is "CSSNumericType.percent"
	//
	// Optional
	//
	// NOTE: FFI_USE_Percent MUST be set to true to make this field effective.
	Percent int32
	// PercentHint is "CSSNumericType.percentHint"
	//
	// Optional
	PercentHint CSSNumericBaseType

	FFI_USE_Length     bool // for Length.
	FFI_USE_Angle      bool // for Angle.
	FFI_USE_Time       bool // for Time.
	FFI_USE_Frequency  bool // for Frequency.
	FFI_USE_Resolution bool // for Resolution.
	FFI_USE_Flex       bool // for Flex.
	FFI_USE_Percent    bool // for Percent.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CSSNumericType with all fields set.
func (p CSSNumericType) FromRef(ref js.Ref) CSSNumericType {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CSSNumericType in the application heap.
func (p CSSNumericType) New() js.Ref {
	return bindings.CSSNumericTypeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p CSSNumericType) UpdateFrom(ref js.Ref) {
	bindings.CSSNumericTypeJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p CSSNumericType) Update(ref js.Ref) {
	bindings.CSSNumericTypeJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type CSSNumericValue struct {
	CSSStyleValue
}

func (this CSSNumericValue) Once() CSSNumericValue {
	this.Ref().Once()
	return this
}

func (this CSSNumericValue) Ref() js.Ref {
	return this.CSSStyleValue.Ref()
}

func (this CSSNumericValue) FromRef(ref js.Ref) CSSNumericValue {
	this.CSSStyleValue = this.CSSStyleValue.FromRef(ref)
	return this
}

func (this CSSNumericValue) Free() {
	this.Ref().Free()
}

// Add calls the method "CSSNumericValue.add".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericValue) Add(values ...CSSNumberish) (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumericValueAdd(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return CSSNumericValue{}.FromRef(_ret), _ok
}

// AddFunc returns the method "CSSNumericValue.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericValue) AddFunc() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	return fn.FromRef(
		bindings.CSSNumericValueAddFunc(
			this.Ref(),
		),
	)
}

// Sub calls the method "CSSNumericValue.sub".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericValue) Sub(values ...CSSNumberish) (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumericValueSub(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return CSSNumericValue{}.FromRef(_ret), _ok
}

// SubFunc returns the method "CSSNumericValue.sub".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericValue) SubFunc() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	return fn.FromRef(
		bindings.CSSNumericValueSubFunc(
			this.Ref(),
		),
	)
}

// Mul calls the method "CSSNumericValue.mul".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericValue) Mul(values ...CSSNumberish) (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumericValueMul(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return CSSNumericValue{}.FromRef(_ret), _ok
}

// MulFunc returns the method "CSSNumericValue.mul".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericValue) MulFunc() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	return fn.FromRef(
		bindings.CSSNumericValueMulFunc(
			this.Ref(),
		),
	)
}

// Div calls the method "CSSNumericValue.div".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericValue) Div(values ...CSSNumberish) (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumericValueDiv(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return CSSNumericValue{}.FromRef(_ret), _ok
}

// DivFunc returns the method "CSSNumericValue.div".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericValue) DivFunc() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	return fn.FromRef(
		bindings.CSSNumericValueDivFunc(
			this.Ref(),
		),
	)
}

// Min calls the method "CSSNumericValue.min".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericValue) Min(values ...CSSNumberish) (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumericValueMin(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return CSSNumericValue{}.FromRef(_ret), _ok
}

// MinFunc returns the method "CSSNumericValue.min".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericValue) MinFunc() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	return fn.FromRef(
		bindings.CSSNumericValueMinFunc(
			this.Ref(),
		),
	)
}

// Max calls the method "CSSNumericValue.max".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericValue) Max(values ...CSSNumberish) (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumericValueMax(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return CSSNumericValue{}.FromRef(_ret), _ok
}

// MaxFunc returns the method "CSSNumericValue.max".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericValue) MaxFunc() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	return fn.FromRef(
		bindings.CSSNumericValueMaxFunc(
			this.Ref(),
		),
	)
}

// Equals calls the method "CSSNumericValue.equals".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericValue) Equals(value ...CSSNumberish) (bool, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumericValueEquals(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(value),
		js.SizeU(len(value)),
	)

	return _ret == js.True, _ok
}

// EqualsFunc returns the method "CSSNumericValue.equals".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericValue) EqualsFunc() (fn js.Func[func(value ...CSSNumberish) bool]) {
	return fn.FromRef(
		bindings.CSSNumericValueEqualsFunc(
			this.Ref(),
		),
	)
}

// To calls the method "CSSNumericValue.to".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericValue) To(unit js.String) (CSSUnitValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumericValueTo(
		this.Ref(), js.Pointer(&_ok),
		unit.Ref(),
	)

	return CSSUnitValue{}.FromRef(_ret), _ok
}

// ToFunc returns the method "CSSNumericValue.to".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericValue) ToFunc() (fn js.Func[func(unit js.String) CSSUnitValue]) {
	return fn.FromRef(
		bindings.CSSNumericValueToFunc(
			this.Ref(),
		),
	)
}

// ToSum calls the method "CSSNumericValue.toSum".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericValue) ToSum(units ...js.String) (CSSMathSum, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumericValueToSum(
		this.Ref(), js.Pointer(&_ok),
		js.SliceData(units),
		js.SizeU(len(units)),
	)

	return CSSMathSum{}.FromRef(_ret), _ok
}

// ToSumFunc returns the method "CSSNumericValue.toSum".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericValue) ToSumFunc() (fn js.Func[func(units ...js.String) CSSMathSum]) {
	return fn.FromRef(
		bindings.CSSNumericValueToSumFunc(
			this.Ref(),
		),
	)
}

// Type calls the method "CSSNumericValue.type".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericValue) Type() (CSSNumericType, bool) {
	var _ret CSSNumericType
	_ok := js.True == bindings.CallCSSNumericValueType(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// TypeFunc returns the method "CSSNumericValue.type".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericValue) TypeFunc() (fn js.Func[func() CSSNumericType]) {
	return fn.FromRef(
		bindings.CSSNumericValueTypeFunc(
			this.Ref(),
		),
	)
}

// Parse calls the staticmethod "CSSNumericValue.parse".
//
// The returned bool will be false if there is no such method.
func (this CSSNumericValue) Parse(cssText js.String) (CSSNumericValue, bool) {
	var _ok bool
	_ret := bindings.CallCSSNumericValueParse(
		this.Ref(), js.Pointer(&_ok),
		cssText.Ref(),
	)

	return CSSNumericValue{}.FromRef(_ret), _ok
}

// ParseFunc returns the staticmethod "CSSNumericValue.parse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CSSNumericValue) ParseFunc() (fn js.Func[func(cssText js.String) CSSNumericValue]) {
	return fn.FromRef(
		bindings.CSSNumericValueParseFunc(
			this.Ref(),
		),
	)
}

type OneOf_Float64_CSSNumericValue_String struct {
	ref js.Ref
}

func (x OneOf_Float64_CSSNumericValue_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_CSSNumericValue_String) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_CSSNumericValue_String) FromRef(ref js.Ref) OneOf_Float64_CSSNumericValue_String {
	return OneOf_Float64_CSSNumericValue_String{
		ref: ref,
	}
}

func (x OneOf_Float64_CSSNumericValue_String) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_CSSNumericValue_String) CSSNumericValue() CSSNumericValue {
	return CSSNumericValue{}.FromRef(x.ref)
}

func (x OneOf_Float64_CSSNumericValue_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type EffectTiming struct {
	// Fill is "EffectTiming.fill"
	//
	// Optional, defaults to "auto".
	Fill FillMode
	// IterationStart is "EffectTiming.iterationStart"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_IterationStart MUST be set to true to make this field effective.
	IterationStart float64
	// Iterations is "EffectTiming.iterations"
	//
	// Optional, defaults to 1.0.
	//
	// NOTE: FFI_USE_Iterations MUST be set to true to make this field effective.
	Iterations float64
	// Direction is "EffectTiming.direction"
	//
	// Optional, defaults to "normal".
	Direction PlaybackDirection
	// Easing is "EffectTiming.easing"
	//
	// Optional, defaults to "linear".
	Easing js.String
	// Delay is "EffectTiming.delay"
	//
	// Optional
	//
	// NOTE: FFI_USE_Delay MUST be set to true to make this field effective.
	Delay float64
	// EndDelay is "EffectTiming.endDelay"
	//
	// Optional
	//
	// NOTE: FFI_USE_EndDelay MUST be set to true to make this field effective.
	EndDelay float64
	// PlaybackRate is "EffectTiming.playbackRate"
	//
	// Optional, defaults to 1.0.
	//
	// NOTE: FFI_USE_PlaybackRate MUST be set to true to make this field effective.
	PlaybackRate float64
	// Duration is "EffectTiming.duration"
	//
	// Optional, defaults to "auto".
	Duration OneOf_Float64_CSSNumericValue_String

	FFI_USE_IterationStart bool // for IterationStart.
	FFI_USE_Iterations     bool // for Iterations.
	FFI_USE_Delay          bool // for Delay.
	FFI_USE_EndDelay       bool // for EndDelay.
	FFI_USE_PlaybackRate   bool // for PlaybackRate.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a EffectTiming with all fields set.
func (p EffectTiming) FromRef(ref js.Ref) EffectTiming {
	p.UpdateFrom(ref)
	return p
}

// New creates a new EffectTiming in the application heap.
func (p EffectTiming) New() js.Ref {
	return bindings.EffectTimingJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p EffectTiming) UpdateFrom(ref js.Ref) {
	bindings.EffectTimingJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p EffectTiming) Update(ref js.Ref) {
	bindings.EffectTimingJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ComputedEffectTiming struct {
	// Progress is "ComputedEffectTiming.progress"
	//
	// Optional
	//
	// NOTE: FFI_USE_Progress MUST be set to true to make this field effective.
	Progress float64
	// CurrentIteration is "ComputedEffectTiming.currentIteration"
	//
	// Optional
	//
	// NOTE: FFI_USE_CurrentIteration MUST be set to true to make this field effective.
	CurrentIteration float64
	// Fill is "ComputedEffectTiming.fill"
	//
	// Optional, defaults to "auto".
	Fill FillMode
	// IterationStart is "ComputedEffectTiming.iterationStart"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_IterationStart MUST be set to true to make this field effective.
	IterationStart float64
	// Iterations is "ComputedEffectTiming.iterations"
	//
	// Optional, defaults to 1.0.
	//
	// NOTE: FFI_USE_Iterations MUST be set to true to make this field effective.
	Iterations float64
	// Direction is "ComputedEffectTiming.direction"
	//
	// Optional, defaults to "normal".
	Direction PlaybackDirection
	// Easing is "ComputedEffectTiming.easing"
	//
	// Optional, defaults to "linear".
	Easing js.String
	// StartTime is "ComputedEffectTiming.startTime"
	//
	// Optional
	StartTime CSSNumberish
	// EndTime is "ComputedEffectTiming.endTime"
	//
	// Optional
	EndTime CSSNumberish
	// ActiveDuration is "ComputedEffectTiming.activeDuration"
	//
	// Optional
	ActiveDuration CSSNumberish
	// LocalTime is "ComputedEffectTiming.localTime"
	//
	// Optional
	LocalTime CSSNumberish

	FFI_USE_Progress         bool // for Progress.
	FFI_USE_CurrentIteration bool // for CurrentIteration.
	FFI_USE_IterationStart   bool // for IterationStart.
	FFI_USE_Iterations       bool // for Iterations.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ComputedEffectTiming with all fields set.
func (p ComputedEffectTiming) FromRef(ref js.Ref) ComputedEffectTiming {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ComputedEffectTiming in the application heap.
func (p ComputedEffectTiming) New() js.Ref {
	return bindings.ComputedEffectTimingJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ComputedEffectTiming) UpdateFrom(ref js.Ref) {
	bindings.ComputedEffectTimingJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ComputedEffectTiming) Update(ref js.Ref) {
	bindings.ComputedEffectTimingJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Float64_String struct {
	ref js.Ref
}

func (x OneOf_Float64_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_String) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_String) FromRef(ref js.Ref) OneOf_Float64_String {
	return OneOf_Float64_String{
		ref: ref,
	}
}

func (x OneOf_Float64_String) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

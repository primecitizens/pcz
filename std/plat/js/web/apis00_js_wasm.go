// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncDrawArraysInstancedANGLE returns true if the method "ANGLE_instanced_arrays.drawArraysInstancedANGLE" exists.
func (this ANGLE_instanced_arrays) HasFuncDrawArraysInstancedANGLE() bool {
	return js.True == bindings.HasFuncANGLE_instanced_arraysDrawArraysInstancedANGLE(
		this.ref,
	)
}

// FuncDrawArraysInstancedANGLE returns the method "ANGLE_instanced_arrays.drawArraysInstancedANGLE".
func (this ANGLE_instanced_arrays) FuncDrawArraysInstancedANGLE() (fn js.Func[func(mode GLenum, first GLint, count GLsizei, primcount GLsizei)]) {
	bindings.FuncANGLE_instanced_arraysDrawArraysInstancedANGLE(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawArraysInstancedANGLE calls the method "ANGLE_instanced_arrays.drawArraysInstancedANGLE".
func (this ANGLE_instanced_arrays) DrawArraysInstancedANGLE(mode GLenum, first GLint, count GLsizei, primcount GLsizei) (ret js.Void) {
	bindings.CallANGLE_instanced_arraysDrawArraysInstancedANGLE(
		this.ref, js.Pointer(&ret),
		uint32(mode),
		int32(first),
		int32(count),
		int32(primcount),
	)

	return
}

// TryDrawArraysInstancedANGLE calls the method "ANGLE_instanced_arrays.drawArraysInstancedANGLE"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ANGLE_instanced_arrays) TryDrawArraysInstancedANGLE(mode GLenum, first GLint, count GLsizei, primcount GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryANGLE_instanced_arraysDrawArraysInstancedANGLE(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(first),
		int32(count),
		int32(primcount),
	)

	return
}

// HasFuncDrawElementsInstancedANGLE returns true if the method "ANGLE_instanced_arrays.drawElementsInstancedANGLE" exists.
func (this ANGLE_instanced_arrays) HasFuncDrawElementsInstancedANGLE() bool {
	return js.True == bindings.HasFuncANGLE_instanced_arraysDrawElementsInstancedANGLE(
		this.ref,
	)
}

// FuncDrawElementsInstancedANGLE returns the method "ANGLE_instanced_arrays.drawElementsInstancedANGLE".
func (this ANGLE_instanced_arrays) FuncDrawElementsInstancedANGLE() (fn js.Func[func(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, primcount GLsizei)]) {
	bindings.FuncANGLE_instanced_arraysDrawElementsInstancedANGLE(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DrawElementsInstancedANGLE calls the method "ANGLE_instanced_arrays.drawElementsInstancedANGLE".
func (this ANGLE_instanced_arrays) DrawElementsInstancedANGLE(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, primcount GLsizei) (ret js.Void) {
	bindings.CallANGLE_instanced_arraysDrawElementsInstancedANGLE(
		this.ref, js.Pointer(&ret),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
		int32(primcount),
	)

	return
}

// TryDrawElementsInstancedANGLE calls the method "ANGLE_instanced_arrays.drawElementsInstancedANGLE"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ANGLE_instanced_arrays) TryDrawElementsInstancedANGLE(mode GLenum, count GLsizei, typ GLenum, offset GLintptr, primcount GLsizei) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryANGLE_instanced_arraysDrawElementsInstancedANGLE(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(mode),
		int32(count),
		uint32(typ),
		float64(offset),
		int32(primcount),
	)

	return
}

// HasFuncVertexAttribDivisorANGLE returns true if the method "ANGLE_instanced_arrays.vertexAttribDivisorANGLE" exists.
func (this ANGLE_instanced_arrays) HasFuncVertexAttribDivisorANGLE() bool {
	return js.True == bindings.HasFuncANGLE_instanced_arraysVertexAttribDivisorANGLE(
		this.ref,
	)
}

// FuncVertexAttribDivisorANGLE returns the method "ANGLE_instanced_arrays.vertexAttribDivisorANGLE".
func (this ANGLE_instanced_arrays) FuncVertexAttribDivisorANGLE() (fn js.Func[func(index GLuint, divisor GLuint)]) {
	bindings.FuncANGLE_instanced_arraysVertexAttribDivisorANGLE(
		this.ref, js.Pointer(&fn),
	)
	return
}

// VertexAttribDivisorANGLE calls the method "ANGLE_instanced_arrays.vertexAttribDivisorANGLE".
func (this ANGLE_instanced_arrays) VertexAttribDivisorANGLE(index GLuint, divisor GLuint) (ret js.Void) {
	bindings.CallANGLE_instanced_arraysVertexAttribDivisorANGLE(
		this.ref, js.Pointer(&ret),
		uint32(index),
		uint32(divisor),
	)

	return
}

// TryVertexAttribDivisorANGLE calls the method "ANGLE_instanced_arrays.vertexAttribDivisorANGLE"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ANGLE_instanced_arrays) TryVertexAttribDivisorANGLE(index GLuint, divisor GLuint) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryANGLE_instanced_arraysVertexAttribDivisorANGLE(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		uint32(divisor),
	)

	return
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
func (p *AV1EncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.AV1EncoderConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AV1EncoderConfig) Update(ref js.Ref) {
	bindings.AV1EncoderConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AV1EncoderConfig) FreeMembers(recursive bool) {
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
func (p *AacEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.AacEncoderConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AacEncoderConfig) Update(ref js.Ref) {
	bindings.AacEncoderConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AacEncoderConfig) FreeMembers(recursive bool) {
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
		js.ThrowInvalidCallbackInvocation()
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
func (p *EventInit) UpdateFrom(ref js.Ref) {
	bindings.EventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EventInit) Update(ref js.Ref) {
	bindings.EventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EventInit) FreeMembers(recursive bool) {
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
		js.ThrowInvalidCallbackInvocation()
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
func (p *AddEventListenerOptions) UpdateFrom(ref js.Ref) {
	bindings.AddEventListenerOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AddEventListenerOptions) Update(ref js.Ref) {
	bindings.AddEventListenerOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AddEventListenerOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
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
func (p *EventListenerOptions) UpdateFrom(ref js.Ref) {
	bindings.EventListenerOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EventListenerOptions) Update(ref js.Ref) {
	bindings.EventListenerOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EventListenerOptions) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncAddEventListener returns true if the method "EventTarget.addEventListener" exists.
func (this EventTarget) HasFuncAddEventListener() bool {
	return js.True == bindings.HasFuncEventTargetAddEventListener(
		this.ref,
	)
}

// FuncAddEventListener returns the method "EventTarget.addEventListener".
func (this EventTarget) FuncAddEventListener() (fn js.Func[func(typ js.String, callback js.Func[func(event Event)], options OneOf_AddEventListenerOptions_Bool)]) {
	bindings.FuncEventTargetAddEventListener(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddEventListener calls the method "EventTarget.addEventListener".
func (this EventTarget) AddEventListener(typ js.String, callback js.Func[func(event Event)], options OneOf_AddEventListenerOptions_Bool) (ret js.Void) {
	bindings.CallEventTargetAddEventListener(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		callback.Ref(),
		options.Ref(),
	)

	return
}

// TryAddEventListener calls the method "EventTarget.addEventListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EventTarget) TryAddEventListener(typ js.String, callback js.Func[func(event Event)], options OneOf_AddEventListenerOptions_Bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventTargetAddEventListener(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		callback.Ref(),
		options.Ref(),
	)

	return
}

// HasFuncAddEventListener1 returns true if the method "EventTarget.addEventListener" exists.
func (this EventTarget) HasFuncAddEventListener1() bool {
	return js.True == bindings.HasFuncEventTargetAddEventListener1(
		this.ref,
	)
}

// FuncAddEventListener1 returns the method "EventTarget.addEventListener".
func (this EventTarget) FuncAddEventListener1() (fn js.Func[func(typ js.String, callback js.Func[func(event Event)])]) {
	bindings.FuncEventTargetAddEventListener1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddEventListener1 calls the method "EventTarget.addEventListener".
func (this EventTarget) AddEventListener1(typ js.String, callback js.Func[func(event Event)]) (ret js.Void) {
	bindings.CallEventTargetAddEventListener1(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		callback.Ref(),
	)

	return
}

// TryAddEventListener1 calls the method "EventTarget.addEventListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EventTarget) TryAddEventListener1(typ js.String, callback js.Func[func(event Event)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventTargetAddEventListener1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncRemoveEventListener returns true if the method "EventTarget.removeEventListener" exists.
func (this EventTarget) HasFuncRemoveEventListener() bool {
	return js.True == bindings.HasFuncEventTargetRemoveEventListener(
		this.ref,
	)
}

// FuncRemoveEventListener returns the method "EventTarget.removeEventListener".
func (this EventTarget) FuncRemoveEventListener() (fn js.Func[func(typ js.String, callback js.Func[func(event Event)], options OneOf_EventListenerOptions_Bool)]) {
	bindings.FuncEventTargetRemoveEventListener(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveEventListener calls the method "EventTarget.removeEventListener".
func (this EventTarget) RemoveEventListener(typ js.String, callback js.Func[func(event Event)], options OneOf_EventListenerOptions_Bool) (ret js.Void) {
	bindings.CallEventTargetRemoveEventListener(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		callback.Ref(),
		options.Ref(),
	)

	return
}

// TryRemoveEventListener calls the method "EventTarget.removeEventListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EventTarget) TryRemoveEventListener(typ js.String, callback js.Func[func(event Event)], options OneOf_EventListenerOptions_Bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventTargetRemoveEventListener(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		callback.Ref(),
		options.Ref(),
	)

	return
}

// HasFuncRemoveEventListener1 returns true if the method "EventTarget.removeEventListener" exists.
func (this EventTarget) HasFuncRemoveEventListener1() bool {
	return js.True == bindings.HasFuncEventTargetRemoveEventListener1(
		this.ref,
	)
}

// FuncRemoveEventListener1 returns the method "EventTarget.removeEventListener".
func (this EventTarget) FuncRemoveEventListener1() (fn js.Func[func(typ js.String, callback js.Func[func(event Event)])]) {
	bindings.FuncEventTargetRemoveEventListener1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveEventListener1 calls the method "EventTarget.removeEventListener".
func (this EventTarget) RemoveEventListener1(typ js.String, callback js.Func[func(event Event)]) (ret js.Void) {
	bindings.CallEventTargetRemoveEventListener1(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		callback.Ref(),
	)

	return
}

// TryRemoveEventListener1 calls the method "EventTarget.removeEventListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EventTarget) TryRemoveEventListener1(typ js.String, callback js.Func[func(event Event)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventTargetRemoveEventListener1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncDispatchEvent returns true if the method "EventTarget.dispatchEvent" exists.
func (this EventTarget) HasFuncDispatchEvent() bool {
	return js.True == bindings.HasFuncEventTargetDispatchEvent(
		this.ref,
	)
}

// FuncDispatchEvent returns the method "EventTarget.dispatchEvent".
func (this EventTarget) FuncDispatchEvent() (fn js.Func[func(event Event) bool]) {
	bindings.FuncEventTargetDispatchEvent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DispatchEvent calls the method "EventTarget.dispatchEvent".
func (this EventTarget) DispatchEvent(event Event) (ret bool) {
	bindings.CallEventTargetDispatchEvent(
		this.ref, js.Pointer(&ret),
		event.Ref(),
	)

	return
}

// TryDispatchEvent calls the method "EventTarget.dispatchEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this EventTarget) TryDispatchEvent(event Event) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventTargetDispatchEvent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		event.Ref(),
	)

	return
}

type DOMHighResTimeStamp float64

func NewEvent(typ js.String, eventInitDict EventInit) (ret Event) {
	ret.ref = bindings.NewEventByEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewEventByEvent1(typ js.String) (ret Event) {
	ret.ref = bindings.NewEventByEvent1(
		typ.Ref())
	return
}

type Event struct {
	ref js.Ref
}

func (this Event) Once() Event {
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "Event.type".
//
// It returns ok=false if there is no such property.
func (this Event) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetEventType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Target returns the value of property "Event.target".
//
// It returns ok=false if there is no such property.
func (this Event) Target() (ret EventTarget, ok bool) {
	ok = js.True == bindings.GetEventTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SrcElement returns the value of property "Event.srcElement".
//
// It returns ok=false if there is no such property.
func (this Event) SrcElement() (ret EventTarget, ok bool) {
	ok = js.True == bindings.GetEventSrcElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CurrentTarget returns the value of property "Event.currentTarget".
//
// It returns ok=false if there is no such property.
func (this Event) CurrentTarget() (ret EventTarget, ok bool) {
	ok = js.True == bindings.GetEventCurrentTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// EventPhase returns the value of property "Event.eventPhase".
//
// It returns ok=false if there is no such property.
func (this Event) EventPhase() (ret uint16, ok bool) {
	ok = js.True == bindings.GetEventEventPhase(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CancelBubble returns the value of property "Event.cancelBubble".
//
// It returns ok=false if there is no such property.
func (this Event) CancelBubble() (ret bool, ok bool) {
	ok = js.True == bindings.GetEventCancelBubble(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCancelBubble sets the value of property "Event.cancelBubble" to val.
//
// It returns false if the property cannot be set.
func (this Event) SetCancelBubble(val bool) bool {
	return js.True == bindings.SetEventCancelBubble(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Bubbles returns the value of property "Event.bubbles".
//
// It returns ok=false if there is no such property.
func (this Event) Bubbles() (ret bool, ok bool) {
	ok = js.True == bindings.GetEventBubbles(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Cancelable returns the value of property "Event.cancelable".
//
// It returns ok=false if there is no such property.
func (this Event) Cancelable() (ret bool, ok bool) {
	ok = js.True == bindings.GetEventCancelable(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReturnValue returns the value of property "Event.returnValue".
//
// It returns ok=false if there is no such property.
func (this Event) ReturnValue() (ret bool, ok bool) {
	ok = js.True == bindings.GetEventReturnValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReturnValue sets the value of property "Event.returnValue" to val.
//
// It returns false if the property cannot be set.
func (this Event) SetReturnValue(val bool) bool {
	return js.True == bindings.SetEventReturnValue(
		this.ref,
		js.Bool(bool(val)),
	)
}

// DefaultPrevented returns the value of property "Event.defaultPrevented".
//
// It returns ok=false if there is no such property.
func (this Event) DefaultPrevented() (ret bool, ok bool) {
	ok = js.True == bindings.GetEventDefaultPrevented(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Composed returns the value of property "Event.composed".
//
// It returns ok=false if there is no such property.
func (this Event) Composed() (ret bool, ok bool) {
	ok = js.True == bindings.GetEventComposed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IsTrusted returns the value of property "Event.isTrusted".
//
// It returns ok=false if there is no such property.
func (this Event) IsTrusted() (ret bool, ok bool) {
	ok = js.True == bindings.GetEventIsTrusted(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TimeStamp returns the value of property "Event.timeStamp".
//
// It returns ok=false if there is no such property.
func (this Event) TimeStamp() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetEventTimeStamp(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncComposedPath returns true if the method "Event.composedPath" exists.
func (this Event) HasFuncComposedPath() bool {
	return js.True == bindings.HasFuncEventComposedPath(
		this.ref,
	)
}

// FuncComposedPath returns the method "Event.composedPath".
func (this Event) FuncComposedPath() (fn js.Func[func() js.Array[EventTarget]]) {
	bindings.FuncEventComposedPath(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ComposedPath calls the method "Event.composedPath".
func (this Event) ComposedPath() (ret js.Array[EventTarget]) {
	bindings.CallEventComposedPath(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryComposedPath calls the method "Event.composedPath"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryComposedPath() (ret js.Array[EventTarget], exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventComposedPath(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStopPropagation returns true if the method "Event.stopPropagation" exists.
func (this Event) HasFuncStopPropagation() bool {
	return js.True == bindings.HasFuncEventStopPropagation(
		this.ref,
	)
}

// FuncStopPropagation returns the method "Event.stopPropagation".
func (this Event) FuncStopPropagation() (fn js.Func[func()]) {
	bindings.FuncEventStopPropagation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StopPropagation calls the method "Event.stopPropagation".
func (this Event) StopPropagation() (ret js.Void) {
	bindings.CallEventStopPropagation(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStopPropagation calls the method "Event.stopPropagation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryStopPropagation() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventStopPropagation(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStopImmediatePropagation returns true if the method "Event.stopImmediatePropagation" exists.
func (this Event) HasFuncStopImmediatePropagation() bool {
	return js.True == bindings.HasFuncEventStopImmediatePropagation(
		this.ref,
	)
}

// FuncStopImmediatePropagation returns the method "Event.stopImmediatePropagation".
func (this Event) FuncStopImmediatePropagation() (fn js.Func[func()]) {
	bindings.FuncEventStopImmediatePropagation(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StopImmediatePropagation calls the method "Event.stopImmediatePropagation".
func (this Event) StopImmediatePropagation() (ret js.Void) {
	bindings.CallEventStopImmediatePropagation(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStopImmediatePropagation calls the method "Event.stopImmediatePropagation"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryStopImmediatePropagation() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventStopImmediatePropagation(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPreventDefault returns true if the method "Event.preventDefault" exists.
func (this Event) HasFuncPreventDefault() bool {
	return js.True == bindings.HasFuncEventPreventDefault(
		this.ref,
	)
}

// FuncPreventDefault returns the method "Event.preventDefault".
func (this Event) FuncPreventDefault() (fn js.Func[func()]) {
	bindings.FuncEventPreventDefault(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PreventDefault calls the method "Event.preventDefault".
func (this Event) PreventDefault() (ret js.Void) {
	bindings.CallEventPreventDefault(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryPreventDefault calls the method "Event.preventDefault"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryPreventDefault() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventPreventDefault(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInitEvent returns true if the method "Event.initEvent" exists.
func (this Event) HasFuncInitEvent() bool {
	return js.True == bindings.HasFuncEventInitEvent(
		this.ref,
	)
}

// FuncInitEvent returns the method "Event.initEvent".
func (this Event) FuncInitEvent() (fn js.Func[func(typ js.String, bubbles bool, cancelable bool)]) {
	bindings.FuncEventInitEvent(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitEvent calls the method "Event.initEvent".
func (this Event) InitEvent(typ js.String, bubbles bool, cancelable bool) (ret js.Void) {
	bindings.CallEventInitEvent(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// TryInitEvent calls the method "Event.initEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryInitEvent(typ js.String, bubbles bool, cancelable bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventInitEvent(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
		js.Bool(bool(cancelable)),
	)

	return
}

// HasFuncInitEvent1 returns true if the method "Event.initEvent" exists.
func (this Event) HasFuncInitEvent1() bool {
	return js.True == bindings.HasFuncEventInitEvent1(
		this.ref,
	)
}

// FuncInitEvent1 returns the method "Event.initEvent".
func (this Event) FuncInitEvent1() (fn js.Func[func(typ js.String, bubbles bool)]) {
	bindings.FuncEventInitEvent1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitEvent1 calls the method "Event.initEvent".
func (this Event) InitEvent1(typ js.String, bubbles bool) (ret js.Void) {
	bindings.CallEventInitEvent1(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// TryInitEvent1 calls the method "Event.initEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryInitEvent1(typ js.String, bubbles bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventInitEvent1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
		js.Bool(bool(bubbles)),
	)

	return
}

// HasFuncInitEvent2 returns true if the method "Event.initEvent" exists.
func (this Event) HasFuncInitEvent2() bool {
	return js.True == bindings.HasFuncEventInitEvent2(
		this.ref,
	)
}

// FuncInitEvent2 returns the method "Event.initEvent".
func (this Event) FuncInitEvent2() (fn js.Func[func(typ js.String)]) {
	bindings.FuncEventInitEvent2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InitEvent2 calls the method "Event.initEvent".
func (this Event) InitEvent2(typ js.String) (ret js.Void) {
	bindings.CallEventInitEvent2(
		this.ref, js.Pointer(&ret),
		typ.Ref(),
	)

	return
}

// TryInitEvent2 calls the method "Event.initEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryInitEvent2(typ js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventInitEvent2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		typ.Ref(),
	)

	return
}

type EventHandler = js.Func[func(event Event) js.Any]

type AbortSignal struct {
	EventTarget
}

func (this AbortSignal) Once() AbortSignal {
	this.ref.Once()
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
	this.ref.Free()
}

// Aborted returns the value of property "AbortSignal.aborted".
//
// It returns ok=false if there is no such property.
func (this AbortSignal) Aborted() (ret bool, ok bool) {
	ok = js.True == bindings.GetAbortSignalAborted(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Reason returns the value of property "AbortSignal.reason".
//
// It returns ok=false if there is no such property.
func (this AbortSignal) Reason() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetAbortSignalReason(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncAbort returns true if the static method "AbortSignal.abort" exists.
func (this AbortSignal) HasFuncAbort() bool {
	return js.True == bindings.HasFuncAbortSignalAbort(
		this.ref,
	)
}

// FuncAbort returns the static method "AbortSignal.abort".
func (this AbortSignal) FuncAbort() (fn js.Func[func(reason js.Any) AbortSignal]) {
	bindings.FuncAbortSignalAbort(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort calls the static method "AbortSignal.abort".
func (this AbortSignal) Abort(reason js.Any) (ret AbortSignal) {
	bindings.CallAbortSignalAbort(
		this.ref, js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryAbort calls the static method "AbortSignal.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AbortSignal) TryAbort(reason js.Any) (ret AbortSignal, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAbortSignalAbort(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasFuncAbort1 returns true if the static method "AbortSignal.abort" exists.
func (this AbortSignal) HasFuncAbort1() bool {
	return js.True == bindings.HasFuncAbortSignalAbort1(
		this.ref,
	)
}

// FuncAbort1 returns the static method "AbortSignal.abort".
func (this AbortSignal) FuncAbort1() (fn js.Func[func() AbortSignal]) {
	bindings.FuncAbortSignalAbort1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort1 calls the static method "AbortSignal.abort".
func (this AbortSignal) Abort1() (ret AbortSignal) {
	bindings.CallAbortSignalAbort1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAbort1 calls the static method "AbortSignal.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AbortSignal) TryAbort1() (ret AbortSignal, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAbortSignalAbort1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncTimeout returns true if the static method "AbortSignal.timeout" exists.
func (this AbortSignal) HasFuncTimeout() bool {
	return js.True == bindings.HasFuncAbortSignalTimeout(
		this.ref,
	)
}

// FuncTimeout returns the static method "AbortSignal.timeout".
func (this AbortSignal) FuncTimeout() (fn js.Func[func(milliseconds uint64) AbortSignal]) {
	bindings.FuncAbortSignalTimeout(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Timeout calls the static method "AbortSignal.timeout".
func (this AbortSignal) Timeout(milliseconds uint64) (ret AbortSignal) {
	bindings.CallAbortSignalTimeout(
		this.ref, js.Pointer(&ret),
		float64(milliseconds),
	)

	return
}

// TryTimeout calls the static method "AbortSignal.timeout"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AbortSignal) TryTimeout(milliseconds uint64) (ret AbortSignal, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAbortSignalTimeout(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(milliseconds),
	)

	return
}

// HasFuncAny returns true if the static method "AbortSignal.any" exists.
func (this AbortSignal) HasFuncAny() bool {
	return js.True == bindings.HasFuncAbortSignalAny(
		this.ref,
	)
}

// FuncAny returns the static method "AbortSignal.any".
func (this AbortSignal) FuncAny() (fn js.Func[func(signals js.Array[AbortSignal]) AbortSignal]) {
	bindings.FuncAbortSignalAny(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Any calls the static method "AbortSignal.any".
func (this AbortSignal) Any(signals js.Array[AbortSignal]) (ret AbortSignal) {
	bindings.CallAbortSignalAny(
		this.ref, js.Pointer(&ret),
		signals.Ref(),
	)

	return
}

// TryAny calls the static method "AbortSignal.any"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AbortSignal) TryAny(signals js.Array[AbortSignal]) (ret AbortSignal, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAbortSignalAny(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		signals.Ref(),
	)

	return
}

// HasFuncThrowIfAborted returns true if the method "AbortSignal.throwIfAborted" exists.
func (this AbortSignal) HasFuncThrowIfAborted() bool {
	return js.True == bindings.HasFuncAbortSignalThrowIfAborted(
		this.ref,
	)
}

// FuncThrowIfAborted returns the method "AbortSignal.throwIfAborted".
func (this AbortSignal) FuncThrowIfAborted() (fn js.Func[func()]) {
	bindings.FuncAbortSignalThrowIfAborted(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ThrowIfAborted calls the method "AbortSignal.throwIfAborted".
func (this AbortSignal) ThrowIfAborted() (ret js.Void) {
	bindings.CallAbortSignalThrowIfAborted(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryThrowIfAborted calls the method "AbortSignal.throwIfAborted"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AbortSignal) TryThrowIfAborted() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAbortSignalThrowIfAborted(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type AbortController struct {
	ref js.Ref
}

func (this AbortController) Once() AbortController {
	this.ref.Once()
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
	this.ref.Free()
}

// Signal returns the value of property "AbortController.signal".
//
// It returns ok=false if there is no such property.
func (this AbortController) Signal() (ret AbortSignal, ok bool) {
	ok = js.True == bindings.GetAbortControllerSignal(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncAbort returns true if the method "AbortController.abort" exists.
func (this AbortController) HasFuncAbort() bool {
	return js.True == bindings.HasFuncAbortControllerAbort(
		this.ref,
	)
}

// FuncAbort returns the method "AbortController.abort".
func (this AbortController) FuncAbort() (fn js.Func[func(reason js.Any)]) {
	bindings.FuncAbortControllerAbort(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort calls the method "AbortController.abort".
func (this AbortController) Abort(reason js.Any) (ret js.Void) {
	bindings.CallAbortControllerAbort(
		this.ref, js.Pointer(&ret),
		reason.Ref(),
	)

	return
}

// TryAbort calls the method "AbortController.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AbortController) TryAbort(reason js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAbortControllerAbort(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		reason.Ref(),
	)

	return
}

// HasFuncAbort1 returns true if the method "AbortController.abort" exists.
func (this AbortController) HasFuncAbort1() bool {
	return js.True == bindings.HasFuncAbortControllerAbort1(
		this.ref,
	)
}

// FuncAbort1 returns the method "AbortController.abort".
func (this AbortController) FuncAbort1() (fn js.Func[func()]) {
	bindings.FuncAbortControllerAbort1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort1 calls the method "AbortController.abort".
func (this AbortController) Abort1() (ret js.Void) {
	bindings.CallAbortControllerAbort1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAbort1 calls the method "AbortController.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this AbortController) TryAbort1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAbortControllerAbort1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
func (p *AbsoluteOrientationReadingValues) UpdateFrom(ref js.Ref) {
	bindings.AbsoluteOrientationReadingValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AbsoluteOrientationReadingValues) Update(ref js.Ref) {
	bindings.AbsoluteOrientationReadingValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AbsoluteOrientationReadingValues) FreeMembers(recursive bool) {
	js.Free(
		p.Quaternion.Ref(),
	)
	p.Quaternion = p.Quaternion.FromRef(js.Undefined)
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
func (p *OrientationSensorOptions) UpdateFrom(ref js.Ref) {
	bindings.OrientationSensorOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OrientationSensorOptions) Update(ref js.Ref) {
	bindings.OrientationSensorOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OrientationSensorOptions) FreeMembers(recursive bool) {
}

func NewAbsoluteOrientationSensor(sensorOptions OrientationSensorOptions) (ret AbsoluteOrientationSensor) {
	ret.ref = bindings.NewAbsoluteOrientationSensorByAbsoluteOrientationSensor(
		js.Pointer(&sensorOptions))
	return
}

func NewAbsoluteOrientationSensorByAbsoluteOrientationSensor1() (ret AbsoluteOrientationSensor) {
	ret.ref = bindings.NewAbsoluteOrientationSensorByAbsoluteOrientationSensor1()
	return
}

type AbsoluteOrientationSensor struct {
	OrientationSensor
}

func (this AbsoluteOrientationSensor) Once() AbsoluteOrientationSensor {
	this.ref.Once()
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
	this.ref.Free()
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
func (p *GetRootNodeOptions) UpdateFrom(ref js.Ref) {
	bindings.GetRootNodeOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GetRootNodeOptions) Update(ref js.Ref) {
	bindings.GetRootNodeOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GetRootNodeOptions) FreeMembers(recursive bool) {
}

type Attr struct {
	Node
}

func (this Attr) Once() Attr {
	this.ref.Once()
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
	this.ref.Free()
}

// NamespaceURI returns the value of property "Attr.namespaceURI".
//
// It returns ok=false if there is no such property.
func (this Attr) NamespaceURI() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAttrNamespaceURI(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Prefix returns the value of property "Attr.prefix".
//
// It returns ok=false if there is no such property.
func (this Attr) Prefix() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAttrPrefix(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LocalName returns the value of property "Attr.localName".
//
// It returns ok=false if there is no such property.
func (this Attr) LocalName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAttrLocalName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "Attr.name".
//
// It returns ok=false if there is no such property.
func (this Attr) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAttrName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "Attr.value".
//
// It returns ok=false if there is no such property.
func (this Attr) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetAttrValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "Attr.value" to val.
//
// It returns false if the property cannot be set.
func (this Attr) SetValue(val js.String) bool {
	return js.True == bindings.SetAttrValue(
		this.ref,
		val.Ref(),
	)
}

// OwnerElement returns the value of property "Attr.ownerElement".
//
// It returns ok=false if there is no such property.
func (this Attr) OwnerElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetAttrOwnerElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Specified returns the value of property "Attr.specified".
//
// It returns ok=false if there is no such property.
func (this Attr) Specified() (ret bool, ok bool) {
	ok = js.True == bindings.GetAttrSpecified(
		this.ref, js.Pointer(&ret),
	)
	return
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

func NewCSSUnitValue(value float64, unit js.String) (ret CSSUnitValue) {
	ret.ref = bindings.NewCSSUnitValueByCSSUnitValue(
		float64(value),
		unit.Ref())
	return
}

type CSSUnitValue struct {
	CSSNumericValue
}

func (this CSSUnitValue) Once() CSSUnitValue {
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "CSSUnitValue.value".
//
// It returns ok=false if there is no such property.
func (this CSSUnitValue) Value() (ret float64, ok bool) {
	ok = js.True == bindings.GetCSSUnitValueValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "CSSUnitValue.value" to val.
//
// It returns false if the property cannot be set.
func (this CSSUnitValue) SetValue(val float64) bool {
	return js.True == bindings.SetCSSUnitValueValue(
		this.ref,
		float64(val),
	)
}

// Unit returns the value of property "CSSUnitValue.unit".
//
// It returns ok=false if there is no such property.
func (this CSSUnitValue) Unit() (ret js.String, ok bool) {
	ok = js.True == bindings.GetCSSUnitValueUnit(
		this.ref, js.Pointer(&ret),
	)
	return
}

type CSSNumericArray struct {
	ref js.Ref
}

func (this CSSNumericArray) Once() CSSNumericArray {
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "CSSNumericArray.length".
//
// It returns ok=false if there is no such property.
func (this CSSNumericArray) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetCSSNumericArrayLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "CSSNumericArray." exists.
func (this CSSNumericArray) HasFuncGet() bool {
	return js.True == bindings.HasFuncCSSNumericArrayGet(
		this.ref,
	)
}

// FuncGet returns the method "CSSNumericArray.".
func (this CSSNumericArray) FuncGet() (fn js.Func[func(index uint32) CSSNumericValue]) {
	bindings.FuncCSSNumericArrayGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "CSSNumericArray.".
func (this CSSNumericArray) Get(index uint32) (ret CSSNumericValue) {
	bindings.CallCSSNumericArrayGet(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryGet calls the method "CSSNumericArray."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericArray) TryGet(index uint32) (ret CSSNumericValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericArrayGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

func NewCSSMathSum(args ...CSSNumberish) (ret CSSMathSum) {
	ret.ref = bindings.NewCSSMathSumByCSSMathSum(
		js.SliceData(args),
		js.SizeU(len(args)))
	return
}

type CSSMathSum struct {
	CSSMathValue
}

func (this CSSMathSum) Once() CSSMathSum {
	this.ref.Once()
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
	this.ref.Free()
}

// Values returns the value of property "CSSMathSum.values".
//
// It returns ok=false if there is no such property.
func (this CSSMathSum) Values() (ret CSSNumericArray, ok bool) {
	ok = js.True == bindings.GetCSSMathSumValues(
		this.ref, js.Pointer(&ret),
	)
	return
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
func (p *CSSNumericType) UpdateFrom(ref js.Ref) {
	bindings.CSSNumericTypeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CSSNumericType) Update(ref js.Ref) {
	bindings.CSSNumericTypeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CSSNumericType) FreeMembers(recursive bool) {
}

type CSSNumericValue struct {
	CSSStyleValue
}

func (this CSSNumericValue) Once() CSSNumericValue {
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncAdd returns true if the method "CSSNumericValue.add" exists.
func (this CSSNumericValue) HasFuncAdd() bool {
	return js.True == bindings.HasFuncCSSNumericValueAdd(
		this.ref,
	)
}

// FuncAdd returns the method "CSSNumericValue.add".
func (this CSSNumericValue) FuncAdd() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	bindings.FuncCSSNumericValueAdd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add calls the method "CSSNumericValue.add".
func (this CSSNumericValue) Add(values ...CSSNumberish) (ret CSSNumericValue) {
	bindings.CallCSSNumericValueAdd(
		this.ref, js.Pointer(&ret),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// TryAdd calls the method "CSSNumericValue.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericValue) TryAdd(values ...CSSNumberish) (ret CSSNumericValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericValueAdd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// HasFuncSub returns true if the method "CSSNumericValue.sub" exists.
func (this CSSNumericValue) HasFuncSub() bool {
	return js.True == bindings.HasFuncCSSNumericValueSub(
		this.ref,
	)
}

// FuncSub returns the method "CSSNumericValue.sub".
func (this CSSNumericValue) FuncSub() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	bindings.FuncCSSNumericValueSub(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Sub calls the method "CSSNumericValue.sub".
func (this CSSNumericValue) Sub(values ...CSSNumberish) (ret CSSNumericValue) {
	bindings.CallCSSNumericValueSub(
		this.ref, js.Pointer(&ret),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// TrySub calls the method "CSSNumericValue.sub"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericValue) TrySub(values ...CSSNumberish) (ret CSSNumericValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericValueSub(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// HasFuncMul returns true if the method "CSSNumericValue.mul" exists.
func (this CSSNumericValue) HasFuncMul() bool {
	return js.True == bindings.HasFuncCSSNumericValueMul(
		this.ref,
	)
}

// FuncMul returns the method "CSSNumericValue.mul".
func (this CSSNumericValue) FuncMul() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	bindings.FuncCSSNumericValueMul(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Mul calls the method "CSSNumericValue.mul".
func (this CSSNumericValue) Mul(values ...CSSNumberish) (ret CSSNumericValue) {
	bindings.CallCSSNumericValueMul(
		this.ref, js.Pointer(&ret),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// TryMul calls the method "CSSNumericValue.mul"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericValue) TryMul(values ...CSSNumberish) (ret CSSNumericValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericValueMul(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// HasFuncDiv returns true if the method "CSSNumericValue.div" exists.
func (this CSSNumericValue) HasFuncDiv() bool {
	return js.True == bindings.HasFuncCSSNumericValueDiv(
		this.ref,
	)
}

// FuncDiv returns the method "CSSNumericValue.div".
func (this CSSNumericValue) FuncDiv() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	bindings.FuncCSSNumericValueDiv(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Div calls the method "CSSNumericValue.div".
func (this CSSNumericValue) Div(values ...CSSNumberish) (ret CSSNumericValue) {
	bindings.CallCSSNumericValueDiv(
		this.ref, js.Pointer(&ret),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// TryDiv calls the method "CSSNumericValue.div"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericValue) TryDiv(values ...CSSNumberish) (ret CSSNumericValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericValueDiv(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// HasFuncMin returns true if the method "CSSNumericValue.min" exists.
func (this CSSNumericValue) HasFuncMin() bool {
	return js.True == bindings.HasFuncCSSNumericValueMin(
		this.ref,
	)
}

// FuncMin returns the method "CSSNumericValue.min".
func (this CSSNumericValue) FuncMin() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	bindings.FuncCSSNumericValueMin(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Min calls the method "CSSNumericValue.min".
func (this CSSNumericValue) Min(values ...CSSNumberish) (ret CSSNumericValue) {
	bindings.CallCSSNumericValueMin(
		this.ref, js.Pointer(&ret),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// TryMin calls the method "CSSNumericValue.min"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericValue) TryMin(values ...CSSNumberish) (ret CSSNumericValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericValueMin(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// HasFuncMax returns true if the method "CSSNumericValue.max" exists.
func (this CSSNumericValue) HasFuncMax() bool {
	return js.True == bindings.HasFuncCSSNumericValueMax(
		this.ref,
	)
}

// FuncMax returns the method "CSSNumericValue.max".
func (this CSSNumericValue) FuncMax() (fn js.Func[func(values ...CSSNumberish) CSSNumericValue]) {
	bindings.FuncCSSNumericValueMax(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Max calls the method "CSSNumericValue.max".
func (this CSSNumericValue) Max(values ...CSSNumberish) (ret CSSNumericValue) {
	bindings.CallCSSNumericValueMax(
		this.ref, js.Pointer(&ret),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// TryMax calls the method "CSSNumericValue.max"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericValue) TryMax(values ...CSSNumberish) (ret CSSNumericValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericValueMax(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(values),
		js.SizeU(len(values)),
	)

	return
}

// HasFuncEquals returns true if the method "CSSNumericValue.equals" exists.
func (this CSSNumericValue) HasFuncEquals() bool {
	return js.True == bindings.HasFuncCSSNumericValueEquals(
		this.ref,
	)
}

// FuncEquals returns the method "CSSNumericValue.equals".
func (this CSSNumericValue) FuncEquals() (fn js.Func[func(value ...CSSNumberish) bool]) {
	bindings.FuncCSSNumericValueEquals(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Equals calls the method "CSSNumericValue.equals".
func (this CSSNumericValue) Equals(value ...CSSNumberish) (ret bool) {
	bindings.CallCSSNumericValueEquals(
		this.ref, js.Pointer(&ret),
		js.SliceData(value),
		js.SizeU(len(value)),
	)

	return
}

// TryEquals calls the method "CSSNumericValue.equals"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericValue) TryEquals(value ...CSSNumberish) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericValueEquals(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(value),
		js.SizeU(len(value)),
	)

	return
}

// HasFuncTo returns true if the method "CSSNumericValue.to" exists.
func (this CSSNumericValue) HasFuncTo() bool {
	return js.True == bindings.HasFuncCSSNumericValueTo(
		this.ref,
	)
}

// FuncTo returns the method "CSSNumericValue.to".
func (this CSSNumericValue) FuncTo() (fn js.Func[func(unit js.String) CSSUnitValue]) {
	bindings.FuncCSSNumericValueTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// To calls the method "CSSNumericValue.to".
func (this CSSNumericValue) To(unit js.String) (ret CSSUnitValue) {
	bindings.CallCSSNumericValueTo(
		this.ref, js.Pointer(&ret),
		unit.Ref(),
	)

	return
}

// TryTo calls the method "CSSNumericValue.to"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericValue) TryTo(unit js.String) (ret CSSUnitValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericValueTo(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		unit.Ref(),
	)

	return
}

// HasFuncToSum returns true if the method "CSSNumericValue.toSum" exists.
func (this CSSNumericValue) HasFuncToSum() bool {
	return js.True == bindings.HasFuncCSSNumericValueToSum(
		this.ref,
	)
}

// FuncToSum returns the method "CSSNumericValue.toSum".
func (this CSSNumericValue) FuncToSum() (fn js.Func[func(units ...js.String) CSSMathSum]) {
	bindings.FuncCSSNumericValueToSum(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ToSum calls the method "CSSNumericValue.toSum".
func (this CSSNumericValue) ToSum(units ...js.String) (ret CSSMathSum) {
	bindings.CallCSSNumericValueToSum(
		this.ref, js.Pointer(&ret),
		js.SliceData(units),
		js.SizeU(len(units)),
	)

	return
}

// TryToSum calls the method "CSSNumericValue.toSum"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericValue) TryToSum(units ...js.String) (ret CSSMathSum, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericValueToSum(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.SliceData(units),
		js.SizeU(len(units)),
	)

	return
}

// HasFuncType returns true if the method "CSSNumericValue.type" exists.
func (this CSSNumericValue) HasFuncType() bool {
	return js.True == bindings.HasFuncCSSNumericValueType(
		this.ref,
	)
}

// FuncType returns the method "CSSNumericValue.type".
func (this CSSNumericValue) FuncType() (fn js.Func[func() CSSNumericType]) {
	bindings.FuncCSSNumericValueType(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Type calls the method "CSSNumericValue.type".
func (this CSSNumericValue) Type() (ret CSSNumericType) {
	bindings.CallCSSNumericValueType(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryType calls the method "CSSNumericValue.type"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericValue) TryType() (ret CSSNumericType, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericValueType(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncParse returns true if the static method "CSSNumericValue.parse" exists.
func (this CSSNumericValue) HasFuncParse() bool {
	return js.True == bindings.HasFuncCSSNumericValueParse(
		this.ref,
	)
}

// FuncParse returns the static method "CSSNumericValue.parse".
func (this CSSNumericValue) FuncParse() (fn js.Func[func(cssText js.String) CSSNumericValue]) {
	bindings.FuncCSSNumericValueParse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Parse calls the static method "CSSNumericValue.parse".
func (this CSSNumericValue) Parse(cssText js.String) (ret CSSNumericValue) {
	bindings.CallCSSNumericValueParse(
		this.ref, js.Pointer(&ret),
		cssText.Ref(),
	)

	return
}

// TryParse calls the static method "CSSNumericValue.parse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CSSNumericValue) TryParse(cssText js.String) (ret CSSNumericValue, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCSSNumericValueParse(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		cssText.Ref(),
	)

	return
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
func (p *EffectTiming) UpdateFrom(ref js.Ref) {
	bindings.EffectTimingJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *EffectTiming) Update(ref js.Ref) {
	bindings.EffectTimingJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *EffectTiming) FreeMembers(recursive bool) {
	js.Free(
		p.Easing.Ref(),
		p.Duration.Ref(),
	)
	p.Easing = p.Easing.FromRef(js.Undefined)
	p.Duration = p.Duration.FromRef(js.Undefined)
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
func (p *ComputedEffectTiming) UpdateFrom(ref js.Ref) {
	bindings.ComputedEffectTimingJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ComputedEffectTiming) Update(ref js.Ref) {
	bindings.ComputedEffectTimingJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ComputedEffectTiming) FreeMembers(recursive bool) {
	js.Free(
		p.Easing.Ref(),
		p.StartTime.Ref(),
		p.EndTime.Ref(),
		p.ActiveDuration.Ref(),
		p.LocalTime.Ref(),
	)
	p.Easing = p.Easing.FromRef(js.Undefined)
	p.StartTime = p.StartTime.FromRef(js.Undefined)
	p.EndTime = p.EndTime.FromRef(js.Undefined)
	p.ActiveDuration = p.ActiveDuration.FromRef(js.Undefined)
	p.LocalTime = p.LocalTime.FromRef(js.Undefined)
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

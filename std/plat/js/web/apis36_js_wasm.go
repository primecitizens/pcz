// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

func NewGPUInternalError(message js.String) (ret GPUInternalError) {
	ret.ref = bindings.NewGPUInternalErrorByGPUInternalError(
		message.Ref())
	return
}

type GPUInternalError struct {
	GPUError
}

func (this GPUInternalError) Once() GPUInternalError {
	this.ref.Once()
	return this
}

func (this GPUInternalError) Ref() js.Ref {
	return this.GPUError.Ref()
}

func (this GPUInternalError) FromRef(ref js.Ref) GPUInternalError {
	this.GPUError = this.GPUError.FromRef(ref)
	return this
}

func (this GPUInternalError) Free() {
	this.ref.Free()
}

const (
	GPUMapMode_READ  GPUFlagsConstant = 0x0001
	GPUMapMode_WRITE GPUFlagsConstant = 0x0002
)

type GPUMapMode struct{}

type GPUObjectDescriptorBase struct {
	// Label is "GPUObjectDescriptorBase.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUObjectDescriptorBase with all fields set.
func (p GPUObjectDescriptorBase) FromRef(ref js.Ref) GPUObjectDescriptorBase {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUObjectDescriptorBase in the application heap.
func (p GPUObjectDescriptorBase) New() js.Ref {
	return bindings.GPUObjectDescriptorBaseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUObjectDescriptorBase) UpdateFrom(ref js.Ref) {
	bindings.GPUObjectDescriptorBaseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUObjectDescriptorBase) Update(ref js.Ref) {
	bindings.GPUObjectDescriptorBaseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUObjectDescriptorBase) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
}

func NewGPUOutOfMemoryError(message js.String) (ret GPUOutOfMemoryError) {
	ret.ref = bindings.NewGPUOutOfMemoryErrorByGPUOutOfMemoryError(
		message.Ref())
	return
}

type GPUOutOfMemoryError struct {
	GPUError
}

func (this GPUOutOfMemoryError) Once() GPUOutOfMemoryError {
	this.ref.Once()
	return this
}

func (this GPUOutOfMemoryError) Ref() js.Ref {
	return this.GPUError.Ref()
}

func (this GPUOutOfMemoryError) FromRef(ref js.Ref) GPUOutOfMemoryError {
	this.GPUError = this.GPUError.FromRef(ref)
	return this
}

func (this GPUOutOfMemoryError) Free() {
	this.ref.Free()
}

type GPUPipelineDescriptorBase struct {
	// Layout is "GPUPipelineDescriptorBase.layout"
	//
	// Required
	Layout OneOf_GPUPipelineLayout_GPUAutoLayoutMode
	// Label is "GPUPipelineDescriptorBase.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUPipelineDescriptorBase with all fields set.
func (p GPUPipelineDescriptorBase) FromRef(ref js.Ref) GPUPipelineDescriptorBase {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUPipelineDescriptorBase in the application heap.
func (p GPUPipelineDescriptorBase) New() js.Ref {
	return bindings.GPUPipelineDescriptorBaseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUPipelineDescriptorBase) UpdateFrom(ref js.Ref) {
	bindings.GPUPipelineDescriptorBaseJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUPipelineDescriptorBase) Update(ref js.Ref) {
	bindings.GPUPipelineDescriptorBaseJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUPipelineDescriptorBase) FreeMembers(recursive bool) {
	js.Free(
		p.Layout.Ref(),
		p.Label.Ref(),
	)
	p.Layout = p.Layout.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
}

type GPUPipelineErrorReason uint32

const (
	_ GPUPipelineErrorReason = iota

	GPUPipelineErrorReason_VALIDATION
	GPUPipelineErrorReason_INTERNAL
)

func (GPUPipelineErrorReason) FromRef(str js.Ref) GPUPipelineErrorReason {
	return GPUPipelineErrorReason(bindings.ConstOfGPUPipelineErrorReason(str))
}

func (x GPUPipelineErrorReason) String() (string, bool) {
	switch x {
	case GPUPipelineErrorReason_VALIDATION:
		return "validation", true
	case GPUPipelineErrorReason_INTERNAL:
		return "internal", true
	default:
		return "", false
	}
}

type GPUPipelineErrorInit struct {
	// Reason is "GPUPipelineErrorInit.reason"
	//
	// Required
	Reason GPUPipelineErrorReason

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPUPipelineErrorInit with all fields set.
func (p GPUPipelineErrorInit) FromRef(ref js.Ref) GPUPipelineErrorInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUPipelineErrorInit in the application heap.
func (p GPUPipelineErrorInit) New() js.Ref {
	return bindings.GPUPipelineErrorInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUPipelineErrorInit) UpdateFrom(ref js.Ref) {
	bindings.GPUPipelineErrorInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUPipelineErrorInit) Update(ref js.Ref) {
	bindings.GPUPipelineErrorInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUPipelineErrorInit) FreeMembers(recursive bool) {
}

func NewGPUPipelineError(message js.String, options GPUPipelineErrorInit) (ret GPUPipelineError) {
	ret.ref = bindings.NewGPUPipelineErrorByGPUPipelineError(
		message.Ref(),
		js.Pointer(&options))
	return
}

func NewGPUPipelineErrorByGPUPipelineError1() (ret GPUPipelineError) {
	ret.ref = bindings.NewGPUPipelineErrorByGPUPipelineError1()
	return
}

type GPUPipelineError struct {
	DOMException
}

func (this GPUPipelineError) Once() GPUPipelineError {
	this.ref.Once()
	return this
}

func (this GPUPipelineError) Ref() js.Ref {
	return this.DOMException.Ref()
}

func (this GPUPipelineError) FromRef(ref js.Ref) GPUPipelineError {
	this.DOMException = this.DOMException.FromRef(ref)
	return this
}

func (this GPUPipelineError) Free() {
	this.ref.Free()
}

// Reason returns the value of property "GPUPipelineError.reason".
//
// It returns ok=false if there is no such property.
func (this GPUPipelineError) Reason() (ret GPUPipelineErrorReason, ok bool) {
	ok = js.True == bindings.GetGPUPipelineErrorReason(
		this.ref, js.Pointer(&ret),
	)
	return
}

type GPURenderPassLayout struct {
	// ColorFormats is "GPURenderPassLayout.colorFormats"
	//
	// Required
	ColorFormats js.Array[GPUTextureFormat]
	// DepthStencilFormat is "GPURenderPassLayout.depthStencilFormat"
	//
	// Optional
	DepthStencilFormat GPUTextureFormat
	// SampleCount is "GPURenderPassLayout.sampleCount"
	//
	// Optional, defaults to 1.
	//
	// NOTE: FFI_USE_SampleCount MUST be set to true to make this field effective.
	SampleCount GPUSize32
	// Label is "GPURenderPassLayout.label"
	//
	// Optional, defaults to "".
	Label js.String

	FFI_USE_SampleCount bool // for SampleCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GPURenderPassLayout with all fields set.
func (p GPURenderPassLayout) FromRef(ref js.Ref) GPURenderPassLayout {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPURenderPassLayout in the application heap.
func (p GPURenderPassLayout) New() js.Ref {
	return bindings.GPURenderPassLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPURenderPassLayout) UpdateFrom(ref js.Ref) {
	bindings.GPURenderPassLayoutJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPURenderPassLayout) Update(ref js.Ref) {
	bindings.GPURenderPassLayoutJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPURenderPassLayout) FreeMembers(recursive bool) {
	js.Free(
		p.ColorFormats.Ref(),
		p.Label.Ref(),
	)
	p.ColorFormats = p.ColorFormats.FromRef(js.Undefined)
	p.Label = p.Label.FromRef(js.Undefined)
}

const (
	GPUShaderStage_VERTEX   GPUFlagsConstant = 0x1
	GPUShaderStage_FRAGMENT GPUFlagsConstant = 0x2
	GPUShaderStage_COMPUTE  GPUFlagsConstant = 0x4
)

type GPUShaderStage struct{}

const (
	GPUTextureUsage_COPY_SRC          GPUFlagsConstant = 0x01
	GPUTextureUsage_COPY_DST          GPUFlagsConstant = 0x02
	GPUTextureUsage_TEXTURE_BINDING   GPUFlagsConstant = 0x04
	GPUTextureUsage_STORAGE_BINDING   GPUFlagsConstant = 0x08
	GPUTextureUsage_RENDER_ATTACHMENT GPUFlagsConstant = 0x10
)

type GPUTextureUsage struct{}

type GPUUncapturedErrorEventInit struct {
	// Error is "GPUUncapturedErrorEventInit.error"
	//
	// Required
	Error GPUError
	// Bubbles is "GPUUncapturedErrorEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "GPUUncapturedErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "GPUUncapturedErrorEventInit.composed"
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

// FromRef calls UpdateFrom and returns a GPUUncapturedErrorEventInit with all fields set.
func (p GPUUncapturedErrorEventInit) FromRef(ref js.Ref) GPUUncapturedErrorEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GPUUncapturedErrorEventInit in the application heap.
func (p GPUUncapturedErrorEventInit) New() js.Ref {
	return bindings.GPUUncapturedErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GPUUncapturedErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.GPUUncapturedErrorEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GPUUncapturedErrorEventInit) Update(ref js.Ref) {
	bindings.GPUUncapturedErrorEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GPUUncapturedErrorEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Error.Ref(),
	)
	p.Error = p.Error.FromRef(js.Undefined)
}

func NewGPUUncapturedErrorEvent(typ js.String, gpuUncapturedErrorEventInitDict GPUUncapturedErrorEventInit) (ret GPUUncapturedErrorEvent) {
	ret.ref = bindings.NewGPUUncapturedErrorEventByGPUUncapturedErrorEvent(
		typ.Ref(),
		js.Pointer(&gpuUncapturedErrorEventInitDict))
	return
}

type GPUUncapturedErrorEvent struct {
	Event
}

func (this GPUUncapturedErrorEvent) Once() GPUUncapturedErrorEvent {
	this.ref.Once()
	return this
}

func (this GPUUncapturedErrorEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this GPUUncapturedErrorEvent) FromRef(ref js.Ref) GPUUncapturedErrorEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this GPUUncapturedErrorEvent) Free() {
	this.ref.Free()
}

// Error returns the value of property "GPUUncapturedErrorEvent.error".
//
// It returns ok=false if there is no such property.
func (this GPUUncapturedErrorEvent) Error() (ret GPUError, ok bool) {
	ok = js.True == bindings.GetGPUUncapturedErrorEventError(
		this.ref, js.Pointer(&ret),
	)
	return
}

func NewGPUValidationError(message js.String) (ret GPUValidationError) {
	ret.ref = bindings.NewGPUValidationErrorByGPUValidationError(
		message.Ref())
	return
}

type GPUValidationError struct {
	GPUError
}

func (this GPUValidationError) Once() GPUValidationError {
	this.ref.Once()
	return this
}

func (this GPUValidationError) Ref() js.Ref {
	return this.GPUError.Ref()
}

func (this GPUValidationError) FromRef(ref js.Ref) GPUValidationError {
	this.GPUError = this.GPUError.FromRef(ref)
	return this
}

func (this GPUValidationError) Free() {
	this.ref.Free()
}

type GamepadEventInit struct {
	// Gamepad is "GamepadEventInit.gamepad"
	//
	// Required
	Gamepad Gamepad
	// Bubbles is "GamepadEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "GamepadEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "GamepadEventInit.composed"
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

// FromRef calls UpdateFrom and returns a GamepadEventInit with all fields set.
func (p GamepadEventInit) FromRef(ref js.Ref) GamepadEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GamepadEventInit in the application heap.
func (p GamepadEventInit) New() js.Ref {
	return bindings.GamepadEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GamepadEventInit) UpdateFrom(ref js.Ref) {
	bindings.GamepadEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GamepadEventInit) Update(ref js.Ref) {
	bindings.GamepadEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GamepadEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Gamepad.Ref(),
	)
	p.Gamepad = p.Gamepad.FromRef(js.Undefined)
}

func NewGamepadEvent(typ js.String, eventInitDict GamepadEventInit) (ret GamepadEvent) {
	ret.ref = bindings.NewGamepadEventByGamepadEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type GamepadEvent struct {
	Event
}

func (this GamepadEvent) Once() GamepadEvent {
	this.ref.Once()
	return this
}

func (this GamepadEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this GamepadEvent) FromRef(ref js.Ref) GamepadEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this GamepadEvent) Free() {
	this.ref.Free()
}

// Gamepad returns the value of property "GamepadEvent.gamepad".
//
// It returns ok=false if there is no such property.
func (this GamepadEvent) Gamepad() (ret Gamepad, ok bool) {
	ok = js.True == bindings.GetGamepadEventGamepad(
		this.ref, js.Pointer(&ret),
	)
	return
}

type GenerateAssertionCallbackFunc func(this js.Ref, contents js.String, origin js.String, options *RTCIdentityProviderOptions) js.Ref

func (fn GenerateAssertionCallbackFunc) Register() js.Func[func(contents js.String, origin js.String, options *RTCIdentityProviderOptions) js.Promise[RTCIdentityAssertionResult]] {
	return js.RegisterCallback[func(contents js.String, origin js.String, options *RTCIdentityProviderOptions) js.Promise[RTCIdentityAssertionResult]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GenerateAssertionCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg2 RTCIdentityProviderOptions
	arg2.UpdateFrom(args[2+1])
	defer arg2.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		mark.NoEscape(&arg2),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GenerateAssertionCallback[T any] struct {
	Fn  func(arg T, this js.Ref, contents js.String, origin js.String, options *RTCIdentityProviderOptions) js.Ref
	Arg T
}

func (cb *GenerateAssertionCallback[T]) Register() js.Func[func(contents js.String, origin js.String, options *RTCIdentityProviderOptions) js.Promise[RTCIdentityAssertionResult]] {
	return js.RegisterCallback[func(contents js.String, origin js.String, options *RTCIdentityProviderOptions) js.Promise[RTCIdentityAssertionResult]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GenerateAssertionCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg2 RTCIdentityProviderOptions
	arg2.UpdateFrom(args[2+1])
	defer arg2.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		mark.NoEscape(&arg2),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RTCIdentityProviderDetails struct {
	// Domain is "RTCIdentityProviderDetails.domain"
	//
	// Required
	Domain js.String
	// Protocol is "RTCIdentityProviderDetails.protocol"
	//
	// Optional, defaults to "default".
	Protocol js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIdentityProviderDetails with all fields set.
func (p RTCIdentityProviderDetails) FromRef(ref js.Ref) RTCIdentityProviderDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIdentityProviderDetails in the application heap.
func (p RTCIdentityProviderDetails) New() js.Ref {
	return bindings.RTCIdentityProviderDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCIdentityProviderDetails) UpdateFrom(ref js.Ref) {
	bindings.RTCIdentityProviderDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIdentityProviderDetails) Update(ref js.Ref) {
	bindings.RTCIdentityProviderDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIdentityProviderDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Domain.Ref(),
		p.Protocol.Ref(),
	)
	p.Domain = p.Domain.FromRef(js.Undefined)
	p.Protocol = p.Protocol.FromRef(js.Undefined)
}

type RTCIdentityAssertionResult struct {
	// Idp is "RTCIdentityAssertionResult.idp"
	//
	// Required
	//
	// NOTE: Idp.FFI_USE MUST be set to true to get Idp used.
	Idp RTCIdentityProviderDetails
	// Assertion is "RTCIdentityAssertionResult.assertion"
	//
	// Required
	Assertion js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIdentityAssertionResult with all fields set.
func (p RTCIdentityAssertionResult) FromRef(ref js.Ref) RTCIdentityAssertionResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIdentityAssertionResult in the application heap.
func (p RTCIdentityAssertionResult) New() js.Ref {
	return bindings.RTCIdentityAssertionResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCIdentityAssertionResult) UpdateFrom(ref js.Ref) {
	bindings.RTCIdentityAssertionResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIdentityAssertionResult) Update(ref js.Ref) {
	bindings.RTCIdentityAssertionResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIdentityAssertionResult) FreeMembers(recursive bool) {
	js.Free(
		p.Assertion.Ref(),
	)
	p.Assertion = p.Assertion.FromRef(js.Undefined)
	if recursive {
		p.Idp.FreeMembers(true)
	}
}

type RTCIdentityProviderOptions struct {
	// Protocol is "RTCIdentityProviderOptions.protocol"
	//
	// Optional, defaults to "default".
	Protocol js.String
	// UsernameHint is "RTCIdentityProviderOptions.usernameHint"
	//
	// Optional
	UsernameHint js.String
	// PeerIdentity is "RTCIdentityProviderOptions.peerIdentity"
	//
	// Optional
	PeerIdentity js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RTCIdentityProviderOptions with all fields set.
func (p RTCIdentityProviderOptions) FromRef(ref js.Ref) RTCIdentityProviderOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RTCIdentityProviderOptions in the application heap.
func (p RTCIdentityProviderOptions) New() js.Ref {
	return bindings.RTCIdentityProviderOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RTCIdentityProviderOptions) UpdateFrom(ref js.Ref) {
	bindings.RTCIdentityProviderOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RTCIdentityProviderOptions) Update(ref js.Ref) {
	bindings.RTCIdentityProviderOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RTCIdentityProviderOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Protocol.Ref(),
		p.UsernameHint.Ref(),
		p.PeerIdentity.Ref(),
	)
	p.Protocol = p.Protocol.FromRef(js.Undefined)
	p.UsernameHint = p.UsernameHint.FromRef(js.Undefined)
	p.PeerIdentity = p.PeerIdentity.FromRef(js.Undefined)
}

type GenerateBidInterestGroup struct {
	// Owner is "GenerateBidInterestGroup.owner"
	//
	// Required
	Owner js.String
	// Name is "GenerateBidInterestGroup.name"
	//
	// Required
	Name js.String
	// LifetimeMs is "GenerateBidInterestGroup.lifetimeMs"
	//
	// Required
	LifetimeMs float64
	// EnableBiddingSignalsPrioritization is "GenerateBidInterestGroup.enableBiddingSignalsPrioritization"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_EnableBiddingSignalsPrioritization MUST be set to true to make this field effective.
	EnableBiddingSignalsPrioritization bool
	// PriorityVector is "GenerateBidInterestGroup.priorityVector"
	//
	// Optional
	PriorityVector js.Record[float64]
	// ExecutionMode is "GenerateBidInterestGroup.executionMode"
	//
	// Optional, defaults to "compatibility".
	ExecutionMode js.String
	// BiddingLogicURL is "GenerateBidInterestGroup.biddingLogicURL"
	//
	// Optional
	BiddingLogicURL js.String
	// BiddingWasmHelperURL is "GenerateBidInterestGroup.biddingWasmHelperURL"
	//
	// Optional
	BiddingWasmHelperURL js.String
	// UpdateURL is "GenerateBidInterestGroup.updateURL"
	//
	// Optional
	UpdateURL js.String
	// TrustedBiddingSignalsURL is "GenerateBidInterestGroup.trustedBiddingSignalsURL"
	//
	// Optional
	TrustedBiddingSignalsURL js.String
	// TrustedBiddingSignalsKeys is "GenerateBidInterestGroup.trustedBiddingSignalsKeys"
	//
	// Optional
	TrustedBiddingSignalsKeys js.Array[js.String]
	// UserBiddingSignals is "GenerateBidInterestGroup.userBiddingSignals"
	//
	// Optional
	UserBiddingSignals js.Any
	// Ads is "GenerateBidInterestGroup.ads"
	//
	// Optional
	Ads js.Array[AuctionAd]
	// AdComponents is "GenerateBidInterestGroup.adComponents"
	//
	// Optional
	AdComponents js.Array[AuctionAd]

	FFI_USE_EnableBiddingSignalsPrioritization bool // for EnableBiddingSignalsPrioritization.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GenerateBidInterestGroup with all fields set.
func (p GenerateBidInterestGroup) FromRef(ref js.Ref) GenerateBidInterestGroup {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GenerateBidInterestGroup in the application heap.
func (p GenerateBidInterestGroup) New() js.Ref {
	return bindings.GenerateBidInterestGroupJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GenerateBidInterestGroup) UpdateFrom(ref js.Ref) {
	bindings.GenerateBidInterestGroupJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GenerateBidInterestGroup) Update(ref js.Ref) {
	bindings.GenerateBidInterestGroupJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GenerateBidInterestGroup) FreeMembers(recursive bool) {
	js.Free(
		p.Owner.Ref(),
		p.Name.Ref(),
		p.PriorityVector.Ref(),
		p.ExecutionMode.Ref(),
		p.BiddingLogicURL.Ref(),
		p.BiddingWasmHelperURL.Ref(),
		p.UpdateURL.Ref(),
		p.TrustedBiddingSignalsURL.Ref(),
		p.TrustedBiddingSignalsKeys.Ref(),
		p.UserBiddingSignals.Ref(),
		p.Ads.Ref(),
		p.AdComponents.Ref(),
	)
	p.Owner = p.Owner.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	p.PriorityVector = p.PriorityVector.FromRef(js.Undefined)
	p.ExecutionMode = p.ExecutionMode.FromRef(js.Undefined)
	p.BiddingLogicURL = p.BiddingLogicURL.FromRef(js.Undefined)
	p.BiddingWasmHelperURL = p.BiddingWasmHelperURL.FromRef(js.Undefined)
	p.UpdateURL = p.UpdateURL.FromRef(js.Undefined)
	p.TrustedBiddingSignalsURL = p.TrustedBiddingSignalsURL.FromRef(js.Undefined)
	p.TrustedBiddingSignalsKeys = p.TrustedBiddingSignalsKeys.FromRef(js.Undefined)
	p.UserBiddingSignals = p.UserBiddingSignals.FromRef(js.Undefined)
	p.Ads = p.Ads.FromRef(js.Undefined)
	p.AdComponents = p.AdComponents.FromRef(js.Undefined)
}

type OneOf_String_AdRender struct {
	ref js.Ref
}

func (x OneOf_String_AdRender) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_AdRender) Free() {
	x.ref.Free()
}

func (x OneOf_String_AdRender) FromRef(ref js.Ref) OneOf_String_AdRender {
	return OneOf_String_AdRender{
		ref: ref,
	}
}

func (x OneOf_String_AdRender) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_AdRender) AdRender() AdRender {
	var ret AdRender
	ret.UpdateFrom(x.ref)
	return ret
}

type GenerateBidOutput struct {
	// Bid is "GenerateBidOutput.bid"
	//
	// Optional, defaults to -1.
	//
	// NOTE: FFI_USE_Bid MUST be set to true to make this field effective.
	Bid float64
	// BidCurrency is "GenerateBidOutput.bidCurrency"
	//
	// Optional
	BidCurrency js.String
	// Render is "GenerateBidOutput.render"
	//
	// Optional
	Render OneOf_String_AdRender
	// Ad is "GenerateBidOutput.ad"
	//
	// Optional
	Ad js.Any
	// AdComponents is "GenerateBidOutput.adComponents"
	//
	// Optional
	AdComponents js.Array[OneOf_String_AdRender]
	// AdCost is "GenerateBidOutput.adCost"
	//
	// Optional
	//
	// NOTE: FFI_USE_AdCost MUST be set to true to make this field effective.
	AdCost float64
	// ModelingSignals is "GenerateBidOutput.modelingSignals"
	//
	// Optional
	//
	// NOTE: FFI_USE_ModelingSignals MUST be set to true to make this field effective.
	ModelingSignals float64
	// AllowComponentAuction is "GenerateBidOutput.allowComponentAuction"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AllowComponentAuction MUST be set to true to make this field effective.
	AllowComponentAuction bool

	FFI_USE_Bid                   bool // for Bid.
	FFI_USE_AdCost                bool // for AdCost.
	FFI_USE_ModelingSignals       bool // for ModelingSignals.
	FFI_USE_AllowComponentAuction bool // for AllowComponentAuction.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GenerateBidOutput with all fields set.
func (p GenerateBidOutput) FromRef(ref js.Ref) GenerateBidOutput {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GenerateBidOutput in the application heap.
func (p GenerateBidOutput) New() js.Ref {
	return bindings.GenerateBidOutputJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GenerateBidOutput) UpdateFrom(ref js.Ref) {
	bindings.GenerateBidOutputJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GenerateBidOutput) Update(ref js.Ref) {
	bindings.GenerateBidOutputJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GenerateBidOutput) FreeMembers(recursive bool) {
	js.Free(
		p.BidCurrency.Ref(),
		p.Render.Ref(),
		p.Ad.Ref(),
		p.AdComponents.Ref(),
	)
	p.BidCurrency = p.BidCurrency.FromRef(js.Undefined)
	p.Render = p.Render.FromRef(js.Undefined)
	p.Ad = p.Ad.FromRef(js.Undefined)
	p.AdComponents = p.AdComponents.FromRef(js.Undefined)
}

type GenerateTestReportParameters struct {
	// Message is "GenerateTestReportParameters.message"
	//
	// Required
	Message js.String
	// Group is "GenerateTestReportParameters.group"
	//
	// Optional, defaults to "default".
	Group js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GenerateTestReportParameters with all fields set.
func (p GenerateTestReportParameters) FromRef(ref js.Ref) GenerateTestReportParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GenerateTestReportParameters in the application heap.
func (p GenerateTestReportParameters) New() js.Ref {
	return bindings.GenerateTestReportParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GenerateTestReportParameters) UpdateFrom(ref js.Ref) {
	bindings.GenerateTestReportParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GenerateTestReportParameters) Update(ref js.Ref) {
	bindings.GenerateTestReportParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GenerateTestReportParameters) FreeMembers(recursive bool) {
	js.Free(
		p.Message.Ref(),
		p.Group.Ref(),
	)
	p.Message = p.Message.FromRef(js.Undefined)
	p.Group = p.Group.FromRef(js.Undefined)
}

type GeolocationReadingValues struct {
	// Latitude is "GeolocationReadingValues.latitude"
	//
	// Required
	Latitude float64
	// Longitude is "GeolocationReadingValues.longitude"
	//
	// Required
	Longitude float64
	// Altitude is "GeolocationReadingValues.altitude"
	//
	// Required
	Altitude float64
	// Accuracy is "GeolocationReadingValues.accuracy"
	//
	// Required
	Accuracy float64
	// AltitudeAccuracy is "GeolocationReadingValues.altitudeAccuracy"
	//
	// Required
	AltitudeAccuracy float64
	// Heading is "GeolocationReadingValues.heading"
	//
	// Required
	Heading float64
	// Speed is "GeolocationReadingValues.speed"
	//
	// Required
	Speed float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GeolocationReadingValues with all fields set.
func (p GeolocationReadingValues) FromRef(ref js.Ref) GeolocationReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GeolocationReadingValues in the application heap.
func (p GeolocationReadingValues) New() js.Ref {
	return bindings.GeolocationReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GeolocationReadingValues) UpdateFrom(ref js.Ref) {
	bindings.GeolocationReadingValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GeolocationReadingValues) Update(ref js.Ref) {
	bindings.GeolocationReadingValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GeolocationReadingValues) FreeMembers(recursive bool) {
}

type GeolocationSensorOptions struct {
	// Frequency is "GeolocationSensorOptions.frequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GeolocationSensorOptions with all fields set.
func (p GeolocationSensorOptions) FromRef(ref js.Ref) GeolocationSensorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GeolocationSensorOptions in the application heap.
func (p GeolocationSensorOptions) New() js.Ref {
	return bindings.GeolocationSensorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GeolocationSensorOptions) UpdateFrom(ref js.Ref) {
	bindings.GeolocationSensorOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GeolocationSensorOptions) Update(ref js.Ref) {
	bindings.GeolocationSensorOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GeolocationSensorOptions) FreeMembers(recursive bool) {
}

type GeolocationSensorReading struct {
	// Timestamp is "GeolocationSensorReading.timestamp"
	//
	// Optional
	//
	// NOTE: FFI_USE_Timestamp MUST be set to true to make this field effective.
	Timestamp DOMHighResTimeStamp
	// Latitude is "GeolocationSensorReading.latitude"
	//
	// Optional
	//
	// NOTE: FFI_USE_Latitude MUST be set to true to make this field effective.
	Latitude float64
	// Longitude is "GeolocationSensorReading.longitude"
	//
	// Optional
	//
	// NOTE: FFI_USE_Longitude MUST be set to true to make this field effective.
	Longitude float64
	// Altitude is "GeolocationSensorReading.altitude"
	//
	// Optional
	//
	// NOTE: FFI_USE_Altitude MUST be set to true to make this field effective.
	Altitude float64
	// Accuracy is "GeolocationSensorReading.accuracy"
	//
	// Optional
	//
	// NOTE: FFI_USE_Accuracy MUST be set to true to make this field effective.
	Accuracy float64
	// AltitudeAccuracy is "GeolocationSensorReading.altitudeAccuracy"
	//
	// Optional
	//
	// NOTE: FFI_USE_AltitudeAccuracy MUST be set to true to make this field effective.
	AltitudeAccuracy float64
	// Heading is "GeolocationSensorReading.heading"
	//
	// Optional
	//
	// NOTE: FFI_USE_Heading MUST be set to true to make this field effective.
	Heading float64
	// Speed is "GeolocationSensorReading.speed"
	//
	// Optional
	//
	// NOTE: FFI_USE_Speed MUST be set to true to make this field effective.
	Speed float64

	FFI_USE_Timestamp        bool // for Timestamp.
	FFI_USE_Latitude         bool // for Latitude.
	FFI_USE_Longitude        bool // for Longitude.
	FFI_USE_Altitude         bool // for Altitude.
	FFI_USE_Accuracy         bool // for Accuracy.
	FFI_USE_AltitudeAccuracy bool // for AltitudeAccuracy.
	FFI_USE_Heading          bool // for Heading.
	FFI_USE_Speed            bool // for Speed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GeolocationSensorReading with all fields set.
func (p GeolocationSensorReading) FromRef(ref js.Ref) GeolocationSensorReading {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GeolocationSensorReading in the application heap.
func (p GeolocationSensorReading) New() js.Ref {
	return bindings.GeolocationSensorReadingJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GeolocationSensorReading) UpdateFrom(ref js.Ref) {
	bindings.GeolocationSensorReadingJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GeolocationSensorReading) Update(ref js.Ref) {
	bindings.GeolocationSensorReadingJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GeolocationSensorReading) FreeMembers(recursive bool) {
}

type ReadOptions struct {
	// Signal is "ReadOptions.signal"
	//
	// Optional
	Signal AbortSignal
	// Frequency is "ReadOptions.frequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadOptions with all fields set.
func (p ReadOptions) FromRef(ref js.Ref) ReadOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ReadOptions in the application heap.
func (p ReadOptions) New() js.Ref {
	return bindings.ReadOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ReadOptions) UpdateFrom(ref js.Ref) {
	bindings.ReadOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ReadOptions) Update(ref js.Ref) {
	bindings.ReadOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ReadOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Signal.Ref(),
	)
	p.Signal = p.Signal.FromRef(js.Undefined)
}

func NewGeolocationSensor(options GeolocationSensorOptions) (ret GeolocationSensor) {
	ret.ref = bindings.NewGeolocationSensorByGeolocationSensor(
		js.Pointer(&options))
	return
}

func NewGeolocationSensorByGeolocationSensor1() (ret GeolocationSensor) {
	ret.ref = bindings.NewGeolocationSensorByGeolocationSensor1()
	return
}

type GeolocationSensor struct {
	Sensor
}

func (this GeolocationSensor) Once() GeolocationSensor {
	this.ref.Once()
	return this
}

func (this GeolocationSensor) Ref() js.Ref {
	return this.Sensor.Ref()
}

func (this GeolocationSensor) FromRef(ref js.Ref) GeolocationSensor {
	this.Sensor = this.Sensor.FromRef(ref)
	return this
}

func (this GeolocationSensor) Free() {
	this.ref.Free()
}

// Latitude returns the value of property "GeolocationSensor.latitude".
//
// It returns ok=false if there is no such property.
func (this GeolocationSensor) Latitude() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationSensorLatitude(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Longitude returns the value of property "GeolocationSensor.longitude".
//
// It returns ok=false if there is no such property.
func (this GeolocationSensor) Longitude() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationSensorLongitude(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Altitude returns the value of property "GeolocationSensor.altitude".
//
// It returns ok=false if there is no such property.
func (this GeolocationSensor) Altitude() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationSensorAltitude(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Accuracy returns the value of property "GeolocationSensor.accuracy".
//
// It returns ok=false if there is no such property.
func (this GeolocationSensor) Accuracy() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationSensorAccuracy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AltitudeAccuracy returns the value of property "GeolocationSensor.altitudeAccuracy".
//
// It returns ok=false if there is no such property.
func (this GeolocationSensor) AltitudeAccuracy() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationSensorAltitudeAccuracy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Heading returns the value of property "GeolocationSensor.heading".
//
// It returns ok=false if there is no such property.
func (this GeolocationSensor) Heading() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationSensorHeading(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Speed returns the value of property "GeolocationSensor.speed".
//
// It returns ok=false if there is no such property.
func (this GeolocationSensor) Speed() (ret float64, ok bool) {
	ok = js.True == bindings.GetGeolocationSensorSpeed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncRead returns true if the static method "GeolocationSensor.read" exists.
func (this GeolocationSensor) HasFuncRead() bool {
	return js.True == bindings.HasFuncGeolocationSensorRead(
		this.ref,
	)
}

// FuncRead returns the static method "GeolocationSensor.read".
func (this GeolocationSensor) FuncRead() (fn js.Func[func(readOptions ReadOptions) js.Promise[GeolocationSensorReading]]) {
	bindings.FuncGeolocationSensorRead(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Read calls the static method "GeolocationSensor.read".
func (this GeolocationSensor) Read(readOptions ReadOptions) (ret js.Promise[GeolocationSensorReading]) {
	bindings.CallGeolocationSensorRead(
		this.ref, js.Pointer(&ret),
		js.Pointer(&readOptions),
	)

	return
}

// TryRead calls the static method "GeolocationSensor.read"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GeolocationSensor) TryRead(readOptions ReadOptions) (ret js.Promise[GeolocationSensorReading], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGeolocationSensorRead(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&readOptions),
	)

	return
}

// HasFuncRead1 returns true if the static method "GeolocationSensor.read" exists.
func (this GeolocationSensor) HasFuncRead1() bool {
	return js.True == bindings.HasFuncGeolocationSensorRead1(
		this.ref,
	)
}

// FuncRead1 returns the static method "GeolocationSensor.read".
func (this GeolocationSensor) FuncRead1() (fn js.Func[func() js.Promise[GeolocationSensorReading]]) {
	bindings.FuncGeolocationSensorRead1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Read1 calls the static method "GeolocationSensor.read".
func (this GeolocationSensor) Read1() (ret js.Promise[GeolocationSensorReading]) {
	bindings.CallGeolocationSensorRead1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRead1 calls the static method "GeolocationSensor.read"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GeolocationSensor) TryRead1() (ret js.Promise[GeolocationSensorReading], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGeolocationSensorRead1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ValueType uint32

const (
	_ ValueType = iota

	ValueType_I32
	ValueType_I64
	ValueType_F32
	ValueType_F64
	ValueType_V128
	ValueType_EXTERNREF
	ValueType_ANYFUNC
)

func (ValueType) FromRef(str js.Ref) ValueType {
	return ValueType(bindings.ConstOfValueType(str))
}

func (x ValueType) String() (string, bool) {
	switch x {
	case ValueType_I32:
		return "i32", true
	case ValueType_I64:
		return "i64", true
	case ValueType_F32:
		return "f32", true
	case ValueType_F64:
		return "f64", true
	case ValueType_V128:
		return "v128", true
	case ValueType_EXTERNREF:
		return "externref", true
	case ValueType_ANYFUNC:
		return "anyfunc", true
	default:
		return "", false
	}
}

type GlobalDescriptor struct {
	// Value is "GlobalDescriptor.value"
	//
	// Required
	Value ValueType
	// Mutable is "GlobalDescriptor.mutable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Mutable MUST be set to true to make this field effective.
	Mutable bool

	FFI_USE_Mutable bool // for Mutable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GlobalDescriptor with all fields set.
func (p GlobalDescriptor) FromRef(ref js.Ref) GlobalDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GlobalDescriptor in the application heap.
func (p GlobalDescriptor) New() js.Ref {
	return bindings.GlobalDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GlobalDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GlobalDescriptorJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GlobalDescriptor) Update(ref js.Ref) {
	bindings.GlobalDescriptorJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GlobalDescriptor) FreeMembers(recursive bool) {
}

func NewGlobal(descriptor GlobalDescriptor, v js.Any) (ret Global) {
	ret.ref = bindings.NewGlobalByGlobal(
		js.Pointer(&descriptor),
		v.Ref())
	return
}

func NewGlobalByGlobal1(descriptor GlobalDescriptor) (ret Global) {
	ret.ref = bindings.NewGlobalByGlobal1(
		js.Pointer(&descriptor))
	return
}

type Global struct {
	ref js.Ref
}

func (this Global) Once() Global {
	this.ref.Once()
	return this
}

func (this Global) Ref() js.Ref {
	return this.ref
}

func (this Global) FromRef(ref js.Ref) Global {
	this.ref = ref
	return this
}

func (this Global) Free() {
	this.ref.Free()
}

// Value returns the value of property "Global.value".
//
// It returns ok=false if there is no such property.
func (this Global) Value() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetGlobalValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "Global.value" to val.
//
// It returns false if the property cannot be set.
func (this Global) SetValue(val js.Any) bool {
	return js.True == bindings.SetGlobalValue(
		this.ref,
		val.Ref(),
	)
}

// HasFuncValueOf returns true if the method "Global.valueOf" exists.
func (this Global) HasFuncValueOf() bool {
	return js.True == bindings.HasFuncGlobalValueOf(
		this.ref,
	)
}

// FuncValueOf returns the method "Global.valueOf".
func (this Global) FuncValueOf() (fn js.Func[func() js.Any]) {
	bindings.FuncGlobalValueOf(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ValueOf calls the method "Global.valueOf".
func (this Global) ValueOf() (ret js.Any) {
	bindings.CallGlobalValueOf(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryValueOf calls the method "Global.valueOf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Global) TryValueOf() (ret js.Any, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGlobalValueOf(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type GravityReadingValues struct {
	// X is "GravityReadingValues.x"
	//
	// Required
	X float64
	// Y is "GravityReadingValues.y"
	//
	// Required
	Y float64
	// Z is "GravityReadingValues.z"
	//
	// Required
	Z float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GravityReadingValues with all fields set.
func (p GravityReadingValues) FromRef(ref js.Ref) GravityReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GravityReadingValues in the application heap.
func (p GravityReadingValues) New() js.Ref {
	return bindings.GravityReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GravityReadingValues) UpdateFrom(ref js.Ref) {
	bindings.GravityReadingValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GravityReadingValues) Update(ref js.Ref) {
	bindings.GravityReadingValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GravityReadingValues) FreeMembers(recursive bool) {
}

func NewGravitySensor(options AccelerometerSensorOptions) (ret GravitySensor) {
	ret.ref = bindings.NewGravitySensorByGravitySensor(
		js.Pointer(&options))
	return
}

func NewGravitySensorByGravitySensor1() (ret GravitySensor) {
	ret.ref = bindings.NewGravitySensorByGravitySensor1()
	return
}

type GravitySensor struct {
	Accelerometer
}

func (this GravitySensor) Once() GravitySensor {
	this.ref.Once()
	return this
}

func (this GravitySensor) Ref() js.Ref {
	return this.Accelerometer.Ref()
}

func (this GravitySensor) FromRef(ref js.Ref) GravitySensor {
	this.Accelerometer = this.Accelerometer.FromRef(ref)
	return this
}

func (this GravitySensor) Free() {
	this.ref.Free()
}

type GyroscopeLocalCoordinateSystem uint32

const (
	_ GyroscopeLocalCoordinateSystem = iota

	GyroscopeLocalCoordinateSystem_DEVICE
	GyroscopeLocalCoordinateSystem_SCREEN
)

func (GyroscopeLocalCoordinateSystem) FromRef(str js.Ref) GyroscopeLocalCoordinateSystem {
	return GyroscopeLocalCoordinateSystem(bindings.ConstOfGyroscopeLocalCoordinateSystem(str))
}

func (x GyroscopeLocalCoordinateSystem) String() (string, bool) {
	switch x {
	case GyroscopeLocalCoordinateSystem_DEVICE:
		return "device", true
	case GyroscopeLocalCoordinateSystem_SCREEN:
		return "screen", true
	default:
		return "", false
	}
}

type GyroscopeSensorOptions struct {
	// ReferenceFrame is "GyroscopeSensorOptions.referenceFrame"
	//
	// Optional, defaults to "device".
	ReferenceFrame GyroscopeLocalCoordinateSystem
	// Frequency is "GyroscopeSensorOptions.frequency"
	//
	// Optional
	//
	// NOTE: FFI_USE_Frequency MUST be set to true to make this field effective.
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GyroscopeSensorOptions with all fields set.
func (p GyroscopeSensorOptions) FromRef(ref js.Ref) GyroscopeSensorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GyroscopeSensorOptions in the application heap.
func (p GyroscopeSensorOptions) New() js.Ref {
	return bindings.GyroscopeSensorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GyroscopeSensorOptions) UpdateFrom(ref js.Ref) {
	bindings.GyroscopeSensorOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GyroscopeSensorOptions) Update(ref js.Ref) {
	bindings.GyroscopeSensorOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GyroscopeSensorOptions) FreeMembers(recursive bool) {
}

func NewGyroscope(sensorOptions GyroscopeSensorOptions) (ret Gyroscope) {
	ret.ref = bindings.NewGyroscopeByGyroscope(
		js.Pointer(&sensorOptions))
	return
}

func NewGyroscopeByGyroscope1() (ret Gyroscope) {
	ret.ref = bindings.NewGyroscopeByGyroscope1()
	return
}

type Gyroscope struct {
	Sensor
}

func (this Gyroscope) Once() Gyroscope {
	this.ref.Once()
	return this
}

func (this Gyroscope) Ref() js.Ref {
	return this.Sensor.Ref()
}

func (this Gyroscope) FromRef(ref js.Ref) Gyroscope {
	this.Sensor = this.Sensor.FromRef(ref)
	return this
}

func (this Gyroscope) Free() {
	this.ref.Free()
}

// X returns the value of property "Gyroscope.x".
//
// It returns ok=false if there is no such property.
func (this Gyroscope) X() (ret float64, ok bool) {
	ok = js.True == bindings.GetGyroscopeX(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Y returns the value of property "Gyroscope.y".
//
// It returns ok=false if there is no such property.
func (this Gyroscope) Y() (ret float64, ok bool) {
	ok = js.True == bindings.GetGyroscopeY(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Z returns the value of property "Gyroscope.z".
//
// It returns ok=false if there is no such property.
func (this Gyroscope) Z() (ret float64, ok bool) {
	ok = js.True == bindings.GetGyroscopeZ(
		this.ref, js.Pointer(&ret),
	)
	return
}

type GyroscopeReadingValues struct {
	// X is "GyroscopeReadingValues.x"
	//
	// Required
	X float64
	// Y is "GyroscopeReadingValues.y"
	//
	// Required
	Y float64
	// Z is "GyroscopeReadingValues.z"
	//
	// Required
	Z float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GyroscopeReadingValues with all fields set.
func (p GyroscopeReadingValues) FromRef(ref js.Ref) GyroscopeReadingValues {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GyroscopeReadingValues in the application heap.
func (p GyroscopeReadingValues) New() js.Ref {
	return bindings.GyroscopeReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *GyroscopeReadingValues) UpdateFrom(ref js.Ref) {
	bindings.GyroscopeReadingValuesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GyroscopeReadingValues) Update(ref js.Ref) {
	bindings.GyroscopeReadingValuesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GyroscopeReadingValues) FreeMembers(recursive bool) {
}

type HIDConnectionEventInit struct {
	// Device is "HIDConnectionEventInit.device"
	//
	// Required
	Device HIDDevice
	// Bubbles is "HIDConnectionEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "HIDConnectionEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "HIDConnectionEventInit.composed"
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

// FromRef calls UpdateFrom and returns a HIDConnectionEventInit with all fields set.
func (p HIDConnectionEventInit) FromRef(ref js.Ref) HIDConnectionEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HIDConnectionEventInit in the application heap.
func (p HIDConnectionEventInit) New() js.Ref {
	return bindings.HIDConnectionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HIDConnectionEventInit) UpdateFrom(ref js.Ref) {
	bindings.HIDConnectionEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HIDConnectionEventInit) Update(ref js.Ref) {
	bindings.HIDConnectionEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HIDConnectionEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Device.Ref(),
	)
	p.Device = p.Device.FromRef(js.Undefined)
}

func NewHIDConnectionEvent(typ js.String, eventInitDict HIDConnectionEventInit) (ret HIDConnectionEvent) {
	ret.ref = bindings.NewHIDConnectionEventByHIDConnectionEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type HIDConnectionEvent struct {
	Event
}

func (this HIDConnectionEvent) Once() HIDConnectionEvent {
	this.ref.Once()
	return this
}

func (this HIDConnectionEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this HIDConnectionEvent) FromRef(ref js.Ref) HIDConnectionEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this HIDConnectionEvent) Free() {
	this.ref.Free()
}

// Device returns the value of property "HIDConnectionEvent.device".
//
// It returns ok=false if there is no such property.
func (this HIDConnectionEvent) Device() (ret HIDDevice, ok bool) {
	ok = js.True == bindings.GetHIDConnectionEventDevice(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HIDInputReportEventInit struct {
	// Device is "HIDInputReportEventInit.device"
	//
	// Required
	Device HIDDevice
	// ReportId is "HIDInputReportEventInit.reportId"
	//
	// Required
	ReportId uint8
	// Data is "HIDInputReportEventInit.data"
	//
	// Required
	Data js.DataView
	// Bubbles is "HIDInputReportEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "HIDInputReportEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "HIDInputReportEventInit.composed"
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

// FromRef calls UpdateFrom and returns a HIDInputReportEventInit with all fields set.
func (p HIDInputReportEventInit) FromRef(ref js.Ref) HIDInputReportEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HIDInputReportEventInit in the application heap.
func (p HIDInputReportEventInit) New() js.Ref {
	return bindings.HIDInputReportEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HIDInputReportEventInit) UpdateFrom(ref js.Ref) {
	bindings.HIDInputReportEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HIDInputReportEventInit) Update(ref js.Ref) {
	bindings.HIDInputReportEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HIDInputReportEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.Device.Ref(),
		p.Data.Ref(),
	)
	p.Device = p.Device.FromRef(js.Undefined)
	p.Data = p.Data.FromRef(js.Undefined)
}

func NewHIDInputReportEvent(typ js.String, eventInitDict HIDInputReportEventInit) (ret HIDInputReportEvent) {
	ret.ref = bindings.NewHIDInputReportEventByHIDInputReportEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

type HIDInputReportEvent struct {
	Event
}

func (this HIDInputReportEvent) Once() HIDInputReportEvent {
	this.ref.Once()
	return this
}

func (this HIDInputReportEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this HIDInputReportEvent) FromRef(ref js.Ref) HIDInputReportEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this HIDInputReportEvent) Free() {
	this.ref.Free()
}

// Device returns the value of property "HIDInputReportEvent.device".
//
// It returns ok=false if there is no such property.
func (this HIDInputReportEvent) Device() (ret HIDDevice, ok bool) {
	ok = js.True == bindings.GetHIDInputReportEventDevice(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReportId returns the value of property "HIDInputReportEvent.reportId".
//
// It returns ok=false if there is no such property.
func (this HIDInputReportEvent) ReportId() (ret uint8, ok bool) {
	ok = js.True == bindings.GetHIDInputReportEventReportId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Data returns the value of property "HIDInputReportEvent.data".
//
// It returns ok=false if there is no such property.
func (this HIDInputReportEvent) Data() (ret js.DataView, ok bool) {
	ok = js.True == bindings.GetHIDInputReportEventData(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HTMLAnchorElement struct {
	HTMLElement
}

func (this HTMLAnchorElement) Once() HTMLAnchorElement {
	this.ref.Once()
	return this
}

func (this HTMLAnchorElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLAnchorElement) FromRef(ref js.Ref) HTMLAnchorElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLAnchorElement) Free() {
	this.ref.Free()
}

// Target returns the value of property "HTMLAnchorElement.target".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Target() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTarget sets the value of property "HTMLAnchorElement.target" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetTarget(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementTarget(
		this.ref,
		val.Ref(),
	)
}

// Download returns the value of property "HTMLAnchorElement.download".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Download() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementDownload(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDownload sets the value of property "HTMLAnchorElement.download" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetDownload(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementDownload(
		this.ref,
		val.Ref(),
	)
}

// Ping returns the value of property "HTMLAnchorElement.ping".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Ping() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementPing(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPing sets the value of property "HTMLAnchorElement.ping" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetPing(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementPing(
		this.ref,
		val.Ref(),
	)
}

// Rel returns the value of property "HTMLAnchorElement.rel".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Rel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementRel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRel sets the value of property "HTMLAnchorElement.rel" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetRel(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementRel(
		this.ref,
		val.Ref(),
	)
}

// RelList returns the value of property "HTMLAnchorElement.relList".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) RelList() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementRelList(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Hreflang returns the value of property "HTMLAnchorElement.hreflang".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Hreflang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementHreflang(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHreflang sets the value of property "HTMLAnchorElement.hreflang" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetHreflang(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementHreflang(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLAnchorElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLAnchorElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementType(
		this.ref,
		val.Ref(),
	)
}

// Text returns the value of property "HTMLAnchorElement.text".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "HTMLAnchorElement.text" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetText(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementText(
		this.ref,
		val.Ref(),
	)
}

// ReferrerPolicy returns the value of property "HTMLAnchorElement.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementReferrerPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLAnchorElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementReferrerPolicy(
		this.ref,
		val.Ref(),
	)
}

// AttributionSourceId returns the value of property "HTMLAnchorElement.attributionSourceId".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) AttributionSourceId() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementAttributionSourceId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAttributionSourceId sets the value of property "HTMLAnchorElement.attributionSourceId" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetAttributionSourceId(val uint32) bool {
	return js.True == bindings.SetHTMLAnchorElementAttributionSourceId(
		this.ref,
		uint32(val),
	)
}

// Coords returns the value of property "HTMLAnchorElement.coords".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Coords() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementCoords(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCoords sets the value of property "HTMLAnchorElement.coords" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetCoords(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementCoords(
		this.ref,
		val.Ref(),
	)
}

// Charset returns the value of property "HTMLAnchorElement.charset".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Charset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementCharset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCharset sets the value of property "HTMLAnchorElement.charset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetCharset(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementCharset(
		this.ref,
		val.Ref(),
	)
}

// Name returns the value of property "HTMLAnchorElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLAnchorElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementName(
		this.ref,
		val.Ref(),
	)
}

// Rev returns the value of property "HTMLAnchorElement.rev".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Rev() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementRev(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRev sets the value of property "HTMLAnchorElement.rev" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetRev(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementRev(
		this.ref,
		val.Ref(),
	)
}

// Shape returns the value of property "HTMLAnchorElement.shape".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Shape() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementShape(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetShape sets the value of property "HTMLAnchorElement.shape" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetShape(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementShape(
		this.ref,
		val.Ref(),
	)
}

// Href returns the value of property "HTMLAnchorElement.href".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHref sets the value of property "HTMLAnchorElement.href" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetHref(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementHref(
		this.ref,
		val.Ref(),
	)
}

// Origin returns the value of property "HTMLAnchorElement.origin".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "HTMLAnchorElement.protocol".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetProtocol sets the value of property "HTMLAnchorElement.protocol" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetProtocol(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementProtocol(
		this.ref,
		val.Ref(),
	)
}

// Username returns the value of property "HTMLAnchorElement.username".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Username() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementUsername(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUsername sets the value of property "HTMLAnchorElement.username" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetUsername(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementUsername(
		this.ref,
		val.Ref(),
	)
}

// Password returns the value of property "HTMLAnchorElement.password".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Password() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementPassword(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPassword sets the value of property "HTMLAnchorElement.password" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetPassword(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementPassword(
		this.ref,
		val.Ref(),
	)
}

// Host returns the value of property "HTMLAnchorElement.host".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Host() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementHost(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHost sets the value of property "HTMLAnchorElement.host" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetHost(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementHost(
		this.ref,
		val.Ref(),
	)
}

// Hostname returns the value of property "HTMLAnchorElement.hostname".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Hostname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementHostname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHostname sets the value of property "HTMLAnchorElement.hostname" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetHostname(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementHostname(
		this.ref,
		val.Ref(),
	)
}

// Port returns the value of property "HTMLAnchorElement.port".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Port() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementPort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPort sets the value of property "HTMLAnchorElement.port" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetPort(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementPort(
		this.ref,
		val.Ref(),
	)
}

// Pathname returns the value of property "HTMLAnchorElement.pathname".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Pathname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementPathname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPathname sets the value of property "HTMLAnchorElement.pathname" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetPathname(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementPathname(
		this.ref,
		val.Ref(),
	)
}

// Search returns the value of property "HTMLAnchorElement.search".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Search() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementSearch(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSearch sets the value of property "HTMLAnchorElement.search" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetSearch(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementSearch(
		this.ref,
		val.Ref(),
	)
}

// Hash returns the value of property "HTMLAnchorElement.hash".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) Hash() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementHash(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHash sets the value of property "HTMLAnchorElement.hash" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetHash(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementHash(
		this.ref,
		val.Ref(),
	)
}

// AttributionSrc returns the value of property "HTMLAnchorElement.attributionSrc".
//
// It returns ok=false if there is no such property.
func (this HTMLAnchorElement) AttributionSrc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAnchorElementAttributionSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAttributionSrc sets the value of property "HTMLAnchorElement.attributionSrc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetAttributionSrc(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementAttributionSrc(
		this.ref,
		val.Ref(),
	)
}

type HTMLAreaElement struct {
	HTMLElement
}

func (this HTMLAreaElement) Once() HTMLAreaElement {
	this.ref.Once()
	return this
}

func (this HTMLAreaElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLAreaElement) FromRef(ref js.Ref) HTMLAreaElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLAreaElement) Free() {
	this.ref.Free()
}

// Alt returns the value of property "HTMLAreaElement.alt".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Alt() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementAlt(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlt sets the value of property "HTMLAreaElement.alt" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetAlt(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementAlt(
		this.ref,
		val.Ref(),
	)
}

// Coords returns the value of property "HTMLAreaElement.coords".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Coords() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementCoords(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCoords sets the value of property "HTMLAreaElement.coords" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetCoords(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementCoords(
		this.ref,
		val.Ref(),
	)
}

// Shape returns the value of property "HTMLAreaElement.shape".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Shape() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementShape(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetShape sets the value of property "HTMLAreaElement.shape" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetShape(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementShape(
		this.ref,
		val.Ref(),
	)
}

// Target returns the value of property "HTMLAreaElement.target".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Target() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTarget sets the value of property "HTMLAreaElement.target" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetTarget(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementTarget(
		this.ref,
		val.Ref(),
	)
}

// Download returns the value of property "HTMLAreaElement.download".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Download() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementDownload(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDownload sets the value of property "HTMLAreaElement.download" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetDownload(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementDownload(
		this.ref,
		val.Ref(),
	)
}

// Ping returns the value of property "HTMLAreaElement.ping".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Ping() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementPing(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPing sets the value of property "HTMLAreaElement.ping" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetPing(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementPing(
		this.ref,
		val.Ref(),
	)
}

// Rel returns the value of property "HTMLAreaElement.rel".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Rel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementRel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRel sets the value of property "HTMLAreaElement.rel" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetRel(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementRel(
		this.ref,
		val.Ref(),
	)
}

// RelList returns the value of property "HTMLAreaElement.relList".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) RelList() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementRelList(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ReferrerPolicy returns the value of property "HTMLAreaElement.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementReferrerPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLAreaElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementReferrerPolicy(
		this.ref,
		val.Ref(),
	)
}

// NoHref returns the value of property "HTMLAreaElement.noHref".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) NoHref() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementNoHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNoHref sets the value of property "HTMLAreaElement.noHref" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetNoHref(val bool) bool {
	return js.True == bindings.SetHTMLAreaElementNoHref(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Href returns the value of property "HTMLAreaElement.href".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHref sets the value of property "HTMLAreaElement.href" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetHref(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementHref(
		this.ref,
		val.Ref(),
	)
}

// Origin returns the value of property "HTMLAreaElement.origin".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Protocol returns the value of property "HTMLAreaElement.protocol".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Protocol() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementProtocol(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetProtocol sets the value of property "HTMLAreaElement.protocol" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetProtocol(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementProtocol(
		this.ref,
		val.Ref(),
	)
}

// Username returns the value of property "HTMLAreaElement.username".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Username() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementUsername(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUsername sets the value of property "HTMLAreaElement.username" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetUsername(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementUsername(
		this.ref,
		val.Ref(),
	)
}

// Password returns the value of property "HTMLAreaElement.password".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Password() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementPassword(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPassword sets the value of property "HTMLAreaElement.password" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetPassword(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementPassword(
		this.ref,
		val.Ref(),
	)
}

// Host returns the value of property "HTMLAreaElement.host".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Host() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementHost(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHost sets the value of property "HTMLAreaElement.host" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetHost(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementHost(
		this.ref,
		val.Ref(),
	)
}

// Hostname returns the value of property "HTMLAreaElement.hostname".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Hostname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementHostname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHostname sets the value of property "HTMLAreaElement.hostname" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetHostname(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementHostname(
		this.ref,
		val.Ref(),
	)
}

// Port returns the value of property "HTMLAreaElement.port".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Port() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementPort(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPort sets the value of property "HTMLAreaElement.port" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetPort(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementPort(
		this.ref,
		val.Ref(),
	)
}

// Pathname returns the value of property "HTMLAreaElement.pathname".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Pathname() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementPathname(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPathname sets the value of property "HTMLAreaElement.pathname" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetPathname(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementPathname(
		this.ref,
		val.Ref(),
	)
}

// Search returns the value of property "HTMLAreaElement.search".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Search() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementSearch(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSearch sets the value of property "HTMLAreaElement.search" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetSearch(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementSearch(
		this.ref,
		val.Ref(),
	)
}

// Hash returns the value of property "HTMLAreaElement.hash".
//
// It returns ok=false if there is no such property.
func (this HTMLAreaElement) Hash() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLAreaElementHash(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHash sets the value of property "HTMLAreaElement.hash" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetHash(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementHash(
		this.ref,
		val.Ref(),
	)
}

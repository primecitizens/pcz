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

func NewGPUInternalError(message js.String) GPUInternalError {
	return GPUInternalError{}.FromRef(
		bindings.NewGPUInternalErrorByGPUInternalError(
			message.Ref()),
	)
}

type GPUInternalError struct {
	GPUError
}

func (this GPUInternalError) Once() GPUInternalError {
	this.Ref().Once()
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
	this.Ref().Free()
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

// New creates a new {0x140004cc0e0 GPUObjectDescriptorBase GPUObjectDescriptorBase [// GPUObjectDescriptorBase] [0x14000726fa0] 0x14000d9cbb8 {0 0}} in the application heap.
func (p GPUObjectDescriptorBase) New() js.Ref {
	return bindings.GPUObjectDescriptorBaseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUObjectDescriptorBase) UpdateFrom(ref js.Ref) {
	bindings.GPUObjectDescriptorBaseJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUObjectDescriptorBase) Update(ref js.Ref) {
	bindings.GPUObjectDescriptorBaseJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewGPUOutOfMemoryError(message js.String) GPUOutOfMemoryError {
	return GPUOutOfMemoryError{}.FromRef(
		bindings.NewGPUOutOfMemoryErrorByGPUOutOfMemoryError(
			message.Ref()),
	)
}

type GPUOutOfMemoryError struct {
	GPUError
}

func (this GPUOutOfMemoryError) Once() GPUOutOfMemoryError {
	this.Ref().Once()
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
	this.Ref().Free()
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

// New creates a new {0x140004cc0e0 GPUPipelineDescriptorBase GPUPipelineDescriptorBase [// GPUPipelineDescriptorBase] [0x14000727040 0x140007270e0] 0x14000d9cbd0 {0 0}} in the application heap.
func (p GPUPipelineDescriptorBase) New() js.Ref {
	return bindings.GPUPipelineDescriptorBaseJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUPipelineDescriptorBase) UpdateFrom(ref js.Ref) {
	bindings.GPUPipelineDescriptorBaseJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUPipelineDescriptorBase) Update(ref js.Ref) {
	bindings.GPUPipelineDescriptorBaseJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 GPUPipelineErrorInit GPUPipelineErrorInit [// GPUPipelineErrorInit] [0x14000727180] 0x14003d941f8 {0 0}} in the application heap.
func (p GPUPipelineErrorInit) New() js.Ref {
	return bindings.GPUPipelineErrorInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUPipelineErrorInit) UpdateFrom(ref js.Ref) {
	bindings.GPUPipelineErrorInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUPipelineErrorInit) Update(ref js.Ref) {
	bindings.GPUPipelineErrorInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewGPUPipelineError(message js.String, options GPUPipelineErrorInit) GPUPipelineError {
	return GPUPipelineError{}.FromRef(
		bindings.NewGPUPipelineErrorByGPUPipelineError(
			message.Ref(),
			js.Pointer(&options)),
	)
}

func NewGPUPipelineErrorByGPUPipelineError1() GPUPipelineError {
	return GPUPipelineError{}.FromRef(
		bindings.NewGPUPipelineErrorByGPUPipelineError1(),
	)
}

type GPUPipelineError struct {
	DOMException
}

func (this GPUPipelineError) Once() GPUPipelineError {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Reason returns the value of property "GPUPipelineError.reason".
//
// The returned bool will be false if there is no such property.
func (this GPUPipelineError) Reason() (GPUPipelineErrorReason, bool) {
	var _ok bool
	_ret := bindings.GetGPUPipelineErrorReason(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUPipelineErrorReason(_ret), _ok
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

// New creates a new {0x140004cc0e0 GPURenderPassLayout GPURenderPassLayout [// GPURenderPassLayout] [0x14000727220 0x140007272c0 0x14000727360 0x140007274a0 0x14000727400] 0x14003d949c0 {0 0}} in the application heap.
func (p GPURenderPassLayout) New() js.Ref {
	return bindings.GPURenderPassLayoutJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPURenderPassLayout) UpdateFrom(ref js.Ref) {
	bindings.GPURenderPassLayoutJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPURenderPassLayout) Update(ref js.Ref) {
	bindings.GPURenderPassLayoutJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	Bubbles bool
	// Cancelable is "GPUUncapturedErrorEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "GPUUncapturedErrorEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 GPUUncapturedErrorEventInit GPUUncapturedErrorEventInit [// GPUUncapturedErrorEventInit] [0x14000727540 0x140007275e0 0x14000727720 0x14000727860 0x14000727680 0x140007277c0 0x14000727900] 0x14003d94bd0 {0 0}} in the application heap.
func (p GPUUncapturedErrorEventInit) New() js.Ref {
	return bindings.GPUUncapturedErrorEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GPUUncapturedErrorEventInit) UpdateFrom(ref js.Ref) {
	bindings.GPUUncapturedErrorEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GPUUncapturedErrorEventInit) Update(ref js.Ref) {
	bindings.GPUUncapturedErrorEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewGPUUncapturedErrorEvent(typ js.String, gpuUncapturedErrorEventInitDict GPUUncapturedErrorEventInit) GPUUncapturedErrorEvent {
	return GPUUncapturedErrorEvent{}.FromRef(
		bindings.NewGPUUncapturedErrorEventByGPUUncapturedErrorEvent(
			typ.Ref(),
			js.Pointer(&gpuUncapturedErrorEventInitDict)),
	)
}

type GPUUncapturedErrorEvent struct {
	Event
}

func (this GPUUncapturedErrorEvent) Once() GPUUncapturedErrorEvent {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Error returns the value of property "GPUUncapturedErrorEvent.error".
//
// The returned bool will be false if there is no such property.
func (this GPUUncapturedErrorEvent) Error() (GPUError, bool) {
	var _ok bool
	_ret := bindings.GetGPUUncapturedErrorEventError(
		this.Ref(), js.Pointer(&_ok),
	)
	return GPUError{}.FromRef(_ret), _ok
}

func NewGPUValidationError(message js.String) GPUValidationError {
	return GPUValidationError{}.FromRef(
		bindings.NewGPUValidationErrorByGPUValidationError(
			message.Ref()),
	)
}

type GPUValidationError struct {
	GPUError
}

func (this GPUValidationError) Once() GPUValidationError {
	this.Ref().Once()
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
	this.Ref().Free()
}

type GamepadEventInit struct {
	// Gamepad is "GamepadEventInit.gamepad"
	//
	// Required
	Gamepad Gamepad
	// Bubbles is "GamepadEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "GamepadEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "GamepadEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 GamepadEventInit GamepadEventInit [// GamepadEventInit] [0x140007279a0 0x14000727a40 0x14000727b80 0x14000727cc0 0x14000727ae0 0x14000727c20 0x14000727d60] 0x14003d94df8 {0 0}} in the application heap.
func (p GamepadEventInit) New() js.Ref {
	return bindings.GamepadEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GamepadEventInit) UpdateFrom(ref js.Ref) {
	bindings.GamepadEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GamepadEventInit) Update(ref js.Ref) {
	bindings.GamepadEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewGamepadEvent(typ js.String, eventInitDict GamepadEventInit) GamepadEvent {
	return GamepadEvent{}.FromRef(
		bindings.NewGamepadEventByGamepadEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type GamepadEvent struct {
	Event
}

func (this GamepadEvent) Once() GamepadEvent {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Gamepad returns the value of property "GamepadEvent.gamepad".
//
// The returned bool will be false if there is no such property.
func (this GamepadEvent) Gamepad() (Gamepad, bool) {
	var _ok bool
	_ret := bindings.GetGamepadEventGamepad(
		this.Ref(), js.Pointer(&_ok),
	)
	return Gamepad{}.FromRef(_ret), _ok
}

type GenerateAssertionCallbackFunc func(this js.Ref, contents js.String, origin js.String, options RTCIdentityProviderOptions) js.Ref

func (fn GenerateAssertionCallbackFunc) Register() js.Func[func(contents js.String, origin js.String, options RTCIdentityProviderOptions) js.Promise[RTCIdentityAssertionResult]] {
	return js.RegisterCallback[func(contents js.String, origin js.String, options RTCIdentityProviderOptions) js.Promise[RTCIdentityAssertionResult]](
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

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		RTCIdentityProviderOptions{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GenerateAssertionCallback[T any] struct {
	Fn  func(arg T, this js.Ref, contents js.String, origin js.String, options RTCIdentityProviderOptions) js.Ref
	Arg T
}

func (cb *GenerateAssertionCallback[T]) Register() js.Func[func(contents js.String, origin js.String, options RTCIdentityProviderOptions) js.Promise[RTCIdentityAssertionResult]] {
	return js.RegisterCallback[func(contents js.String, origin js.String, options RTCIdentityProviderOptions) js.Promise[RTCIdentityAssertionResult]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GenerateAssertionCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		assert.Throw("invalid", "callback", "invocation")
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		RTCIdentityProviderOptions{}.FromRef(args[2+1]),
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

// New creates a new {0x140004cc0e0 RTCIdentityProviderDetails RTCIdentityProviderDetails [// RTCIdentityProviderDetails] [0x14000727e00 0x14000727ea0] 0x14003d95bd8 {0 0}} in the application heap.
func (p RTCIdentityProviderDetails) New() js.Ref {
	return bindings.RTCIdentityProviderDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIdentityProviderDetails) UpdateFrom(ref js.Ref) {
	bindings.RTCIdentityProviderDetailsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIdentityProviderDetails) Update(ref js.Ref) {
	bindings.RTCIdentityProviderDetailsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RTCIdentityAssertionResult struct {
	// Idp is "RTCIdentityAssertionResult.idp"
	//
	// Required
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

// New creates a new {0x140004cc0e0 RTCIdentityAssertionResult RTCIdentityAssertionResult [// RTCIdentityAssertionResult] [0x14000727f40 0x14000734000] 0x14003d94e58 {0 0}} in the application heap.
func (p RTCIdentityAssertionResult) New() js.Ref {
	return bindings.RTCIdentityAssertionResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIdentityAssertionResult) UpdateFrom(ref js.Ref) {
	bindings.RTCIdentityAssertionResultJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIdentityAssertionResult) Update(ref js.Ref) {
	bindings.RTCIdentityAssertionResultJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 RTCIdentityProviderOptions RTCIdentityProviderOptions [// RTCIdentityProviderOptions] [0x140007340a0 0x14000734140 0x140007341e0] 0x14003d95c20 {0 0}} in the application heap.
func (p RTCIdentityProviderOptions) New() js.Ref {
	return bindings.RTCIdentityProviderOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RTCIdentityProviderOptions) UpdateFrom(ref js.Ref) {
	bindings.RTCIdentityProviderOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RTCIdentityProviderOptions) Update(ref js.Ref) {
	bindings.RTCIdentityProviderOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 GenerateBidInterestGroup GenerateBidInterestGroup [// GenerateBidInterestGroup] [0x14000734280 0x14000734320 0x140007343c0 0x14000734460 0x140007345a0 0x14000734640 0x140007346e0 0x14000734780 0x14000734820 0x140007348c0 0x14000734960 0x14000734a00 0x14000734aa0 0x14000734b40 0x14000734500] 0x14003d95f98 {0 0}} in the application heap.
func (p GenerateBidInterestGroup) New() js.Ref {
	return bindings.GenerateBidInterestGroupJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GenerateBidInterestGroup) UpdateFrom(ref js.Ref) {
	bindings.GenerateBidInterestGroupJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GenerateBidInterestGroup) Update(ref js.Ref) {
	bindings.GenerateBidInterestGroupJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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
	AdCost float64
	// ModelingSignals is "GenerateBidOutput.modelingSignals"
	//
	// Optional
	ModelingSignals float64
	// AllowComponentAuction is "GenerateBidOutput.allowComponentAuction"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 GenerateBidOutput GenerateBidOutput [// GenerateBidOutput] [0x14000734be0 0x14000734d20 0x14000734dc0 0x14000734e60 0x14000734f00 0x14000734fa0 0x140007350e0 0x14000735220 0x14000734c80 0x14000735040 0x14000735180 0x140007352c0] 0x14000038468 {0 0}} in the application heap.
func (p GenerateBidOutput) New() js.Ref {
	return bindings.GenerateBidOutputJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GenerateBidOutput) UpdateFrom(ref js.Ref) {
	bindings.GenerateBidOutputJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GenerateBidOutput) Update(ref js.Ref) {
	bindings.GenerateBidOutputJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 GenerateTestReportParameters GenerateTestReportParameters [// GenerateTestReportParameters] [0x14000735360 0x14000735400] 0x14000039650 {0 0}} in the application heap.
func (p GenerateTestReportParameters) New() js.Ref {
	return bindings.GenerateTestReportParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GenerateTestReportParameters) UpdateFrom(ref js.Ref) {
	bindings.GenerateTestReportParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GenerateTestReportParameters) Update(ref js.Ref) {
	bindings.GenerateTestReportParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
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

// New creates a new {0x140004cc0e0 GeolocationReadingValues GeolocationReadingValues [// GeolocationReadingValues] [0x140007354a0 0x14000735540 0x140007355e0 0x14000735680 0x14000735720 0x140007357c0 0x14000735860] 0x140000396b0 {0 0}} in the application heap.
func (p GeolocationReadingValues) New() js.Ref {
	return bindings.GeolocationReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GeolocationReadingValues) UpdateFrom(ref js.Ref) {
	bindings.GeolocationReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GeolocationReadingValues) Update(ref js.Ref) {
	bindings.GeolocationReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GeolocationSensorOptions struct {
	// Frequency is "GeolocationSensorOptions.frequency"
	//
	// Optional
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GeolocationSensorOptions with all fields set.
func (p GeolocationSensorOptions) FromRef(ref js.Ref) GeolocationSensorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GeolocationSensorOptions GeolocationSensorOptions [// GeolocationSensorOptions] [0x14000735900 0x140007359a0] 0x14000039818 {0 0}} in the application heap.
func (p GeolocationSensorOptions) New() js.Ref {
	return bindings.GeolocationSensorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GeolocationSensorOptions) UpdateFrom(ref js.Ref) {
	bindings.GeolocationSensorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GeolocationSensorOptions) Update(ref js.Ref) {
	bindings.GeolocationSensorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GeolocationSensorReading struct {
	// Timestamp is "GeolocationSensorReading.timestamp"
	//
	// Optional
	Timestamp DOMHighResTimeStamp
	// Latitude is "GeolocationSensorReading.latitude"
	//
	// Optional
	Latitude float64
	// Longitude is "GeolocationSensorReading.longitude"
	//
	// Optional
	Longitude float64
	// Altitude is "GeolocationSensorReading.altitude"
	//
	// Optional
	Altitude float64
	// Accuracy is "GeolocationSensorReading.accuracy"
	//
	// Optional
	Accuracy float64
	// AltitudeAccuracy is "GeolocationSensorReading.altitudeAccuracy"
	//
	// Optional
	AltitudeAccuracy float64
	// Heading is "GeolocationSensorReading.heading"
	//
	// Optional
	Heading float64
	// Speed is "GeolocationSensorReading.speed"
	//
	// Optional
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

// New creates a new {0x140004cc0e0 GeolocationSensorReading GeolocationSensorReading [// GeolocationSensorReading] [0x14000735a40 0x14000735b80 0x14000735cc0 0x14000735e00 0x14000735f40 0x140007380a0 0x140007381e0 0x14000738320 0x14000735ae0 0x14000735c20 0x14000735d60 0x14000735ea0 0x14000738000 0x14000738140 0x14000738280 0x140007383c0] 0x14000039b48 {0 0}} in the application heap.
func (p GeolocationSensorReading) New() js.Ref {
	return bindings.GeolocationSensorReadingJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GeolocationSensorReading) UpdateFrom(ref js.Ref) {
	bindings.GeolocationSensorReadingJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GeolocationSensorReading) Update(ref js.Ref) {
	bindings.GeolocationSensorReadingJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type ReadOptions struct {
	// Signal is "ReadOptions.signal"
	//
	// Optional
	Signal AbortSignal
	// Frequency is "ReadOptions.frequency"
	//
	// Optional
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ReadOptions with all fields set.
func (p ReadOptions) FromRef(ref js.Ref) ReadOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 ReadOptions ReadOptions [// ReadOptions] [0x14000738460 0x14000738500 0x140007385a0] 0x14000039c08 {0 0}} in the application heap.
func (p ReadOptions) New() js.Ref {
	return bindings.ReadOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ReadOptions) UpdateFrom(ref js.Ref) {
	bindings.ReadOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ReadOptions) Update(ref js.Ref) {
	bindings.ReadOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewGeolocationSensor(options GeolocationSensorOptions) GeolocationSensor {
	return GeolocationSensor{}.FromRef(
		bindings.NewGeolocationSensorByGeolocationSensor(
			js.Pointer(&options)),
	)
}

func NewGeolocationSensorByGeolocationSensor1() GeolocationSensor {
	return GeolocationSensor{}.FromRef(
		bindings.NewGeolocationSensorByGeolocationSensor1(),
	)
}

type GeolocationSensor struct {
	Sensor
}

func (this GeolocationSensor) Once() GeolocationSensor {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Latitude returns the value of property "GeolocationSensor.latitude".
//
// The returned bool will be false if there is no such property.
func (this GeolocationSensor) Latitude() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationSensorLatitude(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Longitude returns the value of property "GeolocationSensor.longitude".
//
// The returned bool will be false if there is no such property.
func (this GeolocationSensor) Longitude() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationSensorLongitude(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Altitude returns the value of property "GeolocationSensor.altitude".
//
// The returned bool will be false if there is no such property.
func (this GeolocationSensor) Altitude() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationSensorAltitude(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Accuracy returns the value of property "GeolocationSensor.accuracy".
//
// The returned bool will be false if there is no such property.
func (this GeolocationSensor) Accuracy() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationSensorAccuracy(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// AltitudeAccuracy returns the value of property "GeolocationSensor.altitudeAccuracy".
//
// The returned bool will be false if there is no such property.
func (this GeolocationSensor) AltitudeAccuracy() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationSensorAltitudeAccuracy(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Heading returns the value of property "GeolocationSensor.heading".
//
// The returned bool will be false if there is no such property.
func (this GeolocationSensor) Heading() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationSensorHeading(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Speed returns the value of property "GeolocationSensor.speed".
//
// The returned bool will be false if there is no such property.
func (this GeolocationSensor) Speed() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGeolocationSensorSpeed(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Read calls the staticmethod "GeolocationSensor.read".
//
// The returned bool will be false if there is no such method.
func (this GeolocationSensor) Read(readOptions ReadOptions) (js.Promise[GeolocationSensorReading], bool) {
	var _ok bool
	_ret := bindings.CallGeolocationSensorRead(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&readOptions),
	)

	return js.Promise[GeolocationSensorReading]{}.FromRef(_ret), _ok
}

// ReadFunc returns the staticmethod "GeolocationSensor.read".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GeolocationSensor) ReadFunc() (fn js.Func[func(readOptions ReadOptions) js.Promise[GeolocationSensorReading]]) {
	return fn.FromRef(
		bindings.GeolocationSensorReadFunc(
			this.Ref(),
		),
	)
}

// Read1 calls the staticmethod "GeolocationSensor.read".
//
// The returned bool will be false if there is no such method.
func (this GeolocationSensor) Read1() (js.Promise[GeolocationSensorReading], bool) {
	var _ok bool
	_ret := bindings.CallGeolocationSensorRead1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[GeolocationSensorReading]{}.FromRef(_ret), _ok
}

// Read1Func returns the staticmethod "GeolocationSensor.read".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GeolocationSensor) Read1Func() (fn js.Func[func() js.Promise[GeolocationSensorReading]]) {
	return fn.FromRef(
		bindings.GeolocationSensorRead1Func(
			this.Ref(),
		),
	)
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
	Mutable bool

	FFI_USE_Mutable bool // for Mutable.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GlobalDescriptor with all fields set.
func (p GlobalDescriptor) FromRef(ref js.Ref) GlobalDescriptor {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GlobalDescriptor GlobalDescriptor [// GlobalDescriptor] [0x140007386e0 0x14000738780 0x14000738820] 0x14000039c80 {0 0}} in the application heap.
func (p GlobalDescriptor) New() js.Ref {
	return bindings.GlobalDescriptorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GlobalDescriptor) UpdateFrom(ref js.Ref) {
	bindings.GlobalDescriptorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GlobalDescriptor) Update(ref js.Ref) {
	bindings.GlobalDescriptorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewGlobal(descriptor GlobalDescriptor, v js.Any) Global {
	return Global{}.FromRef(
		bindings.NewGlobalByGlobal(
			js.Pointer(&descriptor),
			v.Ref()),
	)
}

func NewGlobalByGlobal1(descriptor GlobalDescriptor) Global {
	return Global{}.FromRef(
		bindings.NewGlobalByGlobal1(
			js.Pointer(&descriptor)),
	)
}

type Global struct {
	ref js.Ref
}

func (this Global) Once() Global {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Value returns the value of property "Global.value".
//
// The returned bool will be false if there is no such property.
func (this Global) Value() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetGlobalValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Value sets the value of property "Global.value" to val.
//
// It returns false if the property cannot be set.
func (this Global) SetValue(val js.Any) bool {
	return js.True == bindings.SetGlobalValue(
		this.Ref(),
		val.Ref(),
	)
}

// ValueOf calls the method "Global.valueOf".
//
// The returned bool will be false if there is no such method.
func (this Global) ValueOf() (js.Any, bool) {
	var _ok bool
	_ret := bindings.CallGlobalValueOf(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Any{}.FromRef(_ret), _ok
}

// ValueOfFunc returns the method "Global.valueOf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Global) ValueOfFunc() (fn js.Func[func() js.Any]) {
	return fn.FromRef(
		bindings.GlobalValueOfFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 GravityReadingValues GravityReadingValues [// GravityReadingValues] [0x140007388c0 0x14000738960 0x14000738a00] 0x1400070a4e0 {0 0}} in the application heap.
func (p GravityReadingValues) New() js.Ref {
	return bindings.GravityReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GravityReadingValues) UpdateFrom(ref js.Ref) {
	bindings.GravityReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GravityReadingValues) Update(ref js.Ref) {
	bindings.GravityReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewGravitySensor(options AccelerometerSensorOptions) GravitySensor {
	return GravitySensor{}.FromRef(
		bindings.NewGravitySensorByGravitySensor(
			js.Pointer(&options)),
	)
}

func NewGravitySensorByGravitySensor1() GravitySensor {
	return GravitySensor{}.FromRef(
		bindings.NewGravitySensorByGravitySensor1(),
	)
}

type GravitySensor struct {
	Accelerometer
}

func (this GravitySensor) Once() GravitySensor {
	this.Ref().Once()
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
	this.Ref().Free()
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
	Frequency float64

	FFI_USE_Frequency bool // for Frequency.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GyroscopeSensorOptions with all fields set.
func (p GyroscopeSensorOptions) FromRef(ref js.Ref) GyroscopeSensorOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 GyroscopeSensorOptions GyroscopeSensorOptions [// GyroscopeSensorOptions] [0x14000738aa0 0x14000738b40 0x14000738be0] 0x1400070a8a0 {0 0}} in the application heap.
func (p GyroscopeSensorOptions) New() js.Ref {
	return bindings.GyroscopeSensorOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GyroscopeSensorOptions) UpdateFrom(ref js.Ref) {
	bindings.GyroscopeSensorOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GyroscopeSensorOptions) Update(ref js.Ref) {
	bindings.GyroscopeSensorOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewGyroscope(sensorOptions GyroscopeSensorOptions) Gyroscope {
	return Gyroscope{}.FromRef(
		bindings.NewGyroscopeByGyroscope(
			js.Pointer(&sensorOptions)),
	)
}

func NewGyroscopeByGyroscope1() Gyroscope {
	return Gyroscope{}.FromRef(
		bindings.NewGyroscopeByGyroscope1(),
	)
}

type Gyroscope struct {
	Sensor
}

func (this Gyroscope) Once() Gyroscope {
	this.Ref().Once()
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
	this.Ref().Free()
}

// X returns the value of property "Gyroscope.x".
//
// The returned bool will be false if there is no such property.
func (this Gyroscope) X() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGyroscopeX(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Y returns the value of property "Gyroscope.y".
//
// The returned bool will be false if there is no such property.
func (this Gyroscope) Y() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGyroscopeY(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Z returns the value of property "Gyroscope.z".
//
// The returned bool will be false if there is no such property.
func (this Gyroscope) Z() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGyroscopeZ(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
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

// New creates a new {0x140004cc0e0 GyroscopeReadingValues GyroscopeReadingValues [// GyroscopeReadingValues] [0x14000738d20 0x14000738dc0 0x14000738e60] 0x1400070a900 {0 0}} in the application heap.
func (p GyroscopeReadingValues) New() js.Ref {
	return bindings.GyroscopeReadingValuesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GyroscopeReadingValues) UpdateFrom(ref js.Ref) {
	bindings.GyroscopeReadingValuesJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GyroscopeReadingValues) Update(ref js.Ref) {
	bindings.GyroscopeReadingValuesJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HIDConnectionEventInit struct {
	// Device is "HIDConnectionEventInit.device"
	//
	// Required
	Device HIDDevice
	// Bubbles is "HIDConnectionEventInit.bubbles"
	//
	// Optional, defaults to false.
	Bubbles bool
	// Cancelable is "HIDConnectionEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "HIDConnectionEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 HIDConnectionEventInit HIDConnectionEventInit [// HIDConnectionEventInit] [0x14000738f00 0x14000738fa0 0x140007390e0 0x14000739220 0x14000739040 0x14000739180 0x140007392c0] 0x1400070b200 {0 0}} in the application heap.
func (p HIDConnectionEventInit) New() js.Ref {
	return bindings.HIDConnectionEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HIDConnectionEventInit) UpdateFrom(ref js.Ref) {
	bindings.HIDConnectionEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HIDConnectionEventInit) Update(ref js.Ref) {
	bindings.HIDConnectionEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewHIDConnectionEvent(typ js.String, eventInitDict HIDConnectionEventInit) HIDConnectionEvent {
	return HIDConnectionEvent{}.FromRef(
		bindings.NewHIDConnectionEventByHIDConnectionEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type HIDConnectionEvent struct {
	Event
}

func (this HIDConnectionEvent) Once() HIDConnectionEvent {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Device returns the value of property "HIDConnectionEvent.device".
//
// The returned bool will be false if there is no such property.
func (this HIDConnectionEvent) Device() (HIDDevice, bool) {
	var _ok bool
	_ret := bindings.GetHIDConnectionEventDevice(
		this.Ref(), js.Pointer(&_ok),
	)
	return HIDDevice{}.FromRef(_ret), _ok
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
	Bubbles bool
	// Cancelable is "HIDInputReportEventInit.cancelable"
	//
	// Optional, defaults to false.
	Cancelable bool
	// Composed is "HIDInputReportEventInit.composed"
	//
	// Optional, defaults to false.
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

// New creates a new {0x140004cc0e0 HIDInputReportEventInit HIDInputReportEventInit [// HIDInputReportEventInit] [0x14000739360 0x14000739400 0x140007394a0 0x14000739540 0x14000739680 0x140007397c0 0x140007395e0 0x14000739720 0x14000739860] 0x1400070b230 {0 0}} in the application heap.
func (p HIDInputReportEventInit) New() js.Ref {
	return bindings.HIDInputReportEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HIDInputReportEventInit) UpdateFrom(ref js.Ref) {
	bindings.HIDInputReportEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HIDInputReportEventInit) Update(ref js.Ref) {
	bindings.HIDInputReportEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewHIDInputReportEvent(typ js.String, eventInitDict HIDInputReportEventInit) HIDInputReportEvent {
	return HIDInputReportEvent{}.FromRef(
		bindings.NewHIDInputReportEventByHIDInputReportEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

type HIDInputReportEvent struct {
	Event
}

func (this HIDInputReportEvent) Once() HIDInputReportEvent {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Device returns the value of property "HIDInputReportEvent.device".
//
// The returned bool will be false if there is no such property.
func (this HIDInputReportEvent) Device() (HIDDevice, bool) {
	var _ok bool
	_ret := bindings.GetHIDInputReportEventDevice(
		this.Ref(), js.Pointer(&_ok),
	)
	return HIDDevice{}.FromRef(_ret), _ok
}

// ReportId returns the value of property "HIDInputReportEvent.reportId".
//
// The returned bool will be false if there is no such property.
func (this HIDInputReportEvent) ReportId() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetHIDInputReportEventReportId(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// Data returns the value of property "HIDInputReportEvent.data".
//
// The returned bool will be false if there is no such property.
func (this HIDInputReportEvent) Data() (js.DataView, bool) {
	var _ok bool
	_ret := bindings.GetHIDInputReportEventData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.DataView{}.FromRef(_ret), _ok
}

func NewHTMLAnchorElement() HTMLAnchorElement {
	return HTMLAnchorElement{}.FromRef(
		bindings.NewHTMLAnchorElementByHTMLAnchorElement(),
	)
}

type HTMLAnchorElement struct {
	HTMLElement
}

func (this HTMLAnchorElement) Once() HTMLAnchorElement {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Target returns the value of property "HTMLAnchorElement.target".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Target() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Target sets the value of property "HTMLAnchorElement.target" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetTarget(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementTarget(
		this.Ref(),
		val.Ref(),
	)
}

// Download returns the value of property "HTMLAnchorElement.download".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Download() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementDownload(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Download sets the value of property "HTMLAnchorElement.download" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetDownload(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementDownload(
		this.Ref(),
		val.Ref(),
	)
}

// Ping returns the value of property "HTMLAnchorElement.ping".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Ping() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementPing(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Ping sets the value of property "HTMLAnchorElement.ping" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetPing(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementPing(
		this.Ref(),
		val.Ref(),
	)
}

// Rel returns the value of property "HTMLAnchorElement.rel".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Rel() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementRel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Rel sets the value of property "HTMLAnchorElement.rel" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetRel(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementRel(
		this.Ref(),
		val.Ref(),
	)
}

// RelList returns the value of property "HTMLAnchorElement.relList".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) RelList() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementRelList(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// Hreflang returns the value of property "HTMLAnchorElement.hreflang".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Hreflang() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementHreflang(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hreflang sets the value of property "HTMLAnchorElement.hreflang" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetHreflang(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementHreflang(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLAnchorElement.type".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLAnchorElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementType(
		this.Ref(),
		val.Ref(),
	)
}

// Text returns the value of property "HTMLAnchorElement.text".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Text() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Text sets the value of property "HTMLAnchorElement.text" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetText(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementText(
		this.Ref(),
		val.Ref(),
	)
}

// ReferrerPolicy returns the value of property "HTMLAnchorElement.referrerPolicy".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) ReferrerPolicy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementReferrerPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReferrerPolicy sets the value of property "HTMLAnchorElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementReferrerPolicy(
		this.Ref(),
		val.Ref(),
	)
}

// AttributionSourceId returns the value of property "HTMLAnchorElement.attributionSourceId".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) AttributionSourceId() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementAttributionSourceId(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// AttributionSourceId sets the value of property "HTMLAnchorElement.attributionSourceId" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetAttributionSourceId(val uint32) bool {
	return js.True == bindings.SetHTMLAnchorElementAttributionSourceId(
		this.Ref(),
		uint32(val),
	)
}

// Coords returns the value of property "HTMLAnchorElement.coords".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Coords() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementCoords(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Coords sets the value of property "HTMLAnchorElement.coords" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetCoords(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementCoords(
		this.Ref(),
		val.Ref(),
	)
}

// Charset returns the value of property "HTMLAnchorElement.charset".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Charset() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementCharset(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Charset sets the value of property "HTMLAnchorElement.charset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetCharset(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementCharset(
		this.Ref(),
		val.Ref(),
	)
}

// Name returns the value of property "HTMLAnchorElement.name".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLAnchorElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Rev returns the value of property "HTMLAnchorElement.rev".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Rev() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementRev(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Rev sets the value of property "HTMLAnchorElement.rev" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetRev(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementRev(
		this.Ref(),
		val.Ref(),
	)
}

// Shape returns the value of property "HTMLAnchorElement.shape".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Shape() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementShape(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Shape sets the value of property "HTMLAnchorElement.shape" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetShape(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementShape(
		this.Ref(),
		val.Ref(),
	)
}

// Href returns the value of property "HTMLAnchorElement.href".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Href() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Href sets the value of property "HTMLAnchorElement.href" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetHref(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementHref(
		this.Ref(),
		val.Ref(),
	)
}

// Origin returns the value of property "HTMLAnchorElement.origin".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Origin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol returns the value of property "HTMLAnchorElement.protocol".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Protocol() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol sets the value of property "HTMLAnchorElement.protocol" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetProtocol(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementProtocol(
		this.Ref(),
		val.Ref(),
	)
}

// Username returns the value of property "HTMLAnchorElement.username".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Username() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementUsername(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Username sets the value of property "HTMLAnchorElement.username" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetUsername(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementUsername(
		this.Ref(),
		val.Ref(),
	)
}

// Password returns the value of property "HTMLAnchorElement.password".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Password() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementPassword(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Password sets the value of property "HTMLAnchorElement.password" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetPassword(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementPassword(
		this.Ref(),
		val.Ref(),
	)
}

// Host returns the value of property "HTMLAnchorElement.host".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Host() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementHost(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Host sets the value of property "HTMLAnchorElement.host" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetHost(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementHost(
		this.Ref(),
		val.Ref(),
	)
}

// Hostname returns the value of property "HTMLAnchorElement.hostname".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Hostname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementHostname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hostname sets the value of property "HTMLAnchorElement.hostname" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetHostname(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementHostname(
		this.Ref(),
		val.Ref(),
	)
}

// Port returns the value of property "HTMLAnchorElement.port".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Port() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Port sets the value of property "HTMLAnchorElement.port" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetPort(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementPort(
		this.Ref(),
		val.Ref(),
	)
}

// Pathname returns the value of property "HTMLAnchorElement.pathname".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Pathname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementPathname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Pathname sets the value of property "HTMLAnchorElement.pathname" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetPathname(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementPathname(
		this.Ref(),
		val.Ref(),
	)
}

// Search returns the value of property "HTMLAnchorElement.search".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Search() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementSearch(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Search sets the value of property "HTMLAnchorElement.search" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetSearch(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementSearch(
		this.Ref(),
		val.Ref(),
	)
}

// Hash returns the value of property "HTMLAnchorElement.hash".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) Hash() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementHash(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hash sets the value of property "HTMLAnchorElement.hash" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetHash(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementHash(
		this.Ref(),
		val.Ref(),
	)
}

// AttributionSrc returns the value of property "HTMLAnchorElement.attributionSrc".
//
// The returned bool will be false if there is no such property.
func (this HTMLAnchorElement) AttributionSrc() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAnchorElementAttributionSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// AttributionSrc sets the value of property "HTMLAnchorElement.attributionSrc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAnchorElement) SetAttributionSrc(val js.String) bool {
	return js.True == bindings.SetHTMLAnchorElementAttributionSrc(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLAreaElement() HTMLAreaElement {
	return HTMLAreaElement{}.FromRef(
		bindings.NewHTMLAreaElementByHTMLAreaElement(),
	)
}

type HTMLAreaElement struct {
	HTMLElement
}

func (this HTMLAreaElement) Once() HTMLAreaElement {
	this.Ref().Once()
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
	this.Ref().Free()
}

// Alt returns the value of property "HTMLAreaElement.alt".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Alt() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementAlt(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Alt sets the value of property "HTMLAreaElement.alt" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetAlt(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementAlt(
		this.Ref(),
		val.Ref(),
	)
}

// Coords returns the value of property "HTMLAreaElement.coords".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Coords() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementCoords(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Coords sets the value of property "HTMLAreaElement.coords" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetCoords(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementCoords(
		this.Ref(),
		val.Ref(),
	)
}

// Shape returns the value of property "HTMLAreaElement.shape".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Shape() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementShape(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Shape sets the value of property "HTMLAreaElement.shape" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetShape(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementShape(
		this.Ref(),
		val.Ref(),
	)
}

// Target returns the value of property "HTMLAreaElement.target".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Target() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Target sets the value of property "HTMLAreaElement.target" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetTarget(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementTarget(
		this.Ref(),
		val.Ref(),
	)
}

// Download returns the value of property "HTMLAreaElement.download".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Download() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementDownload(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Download sets the value of property "HTMLAreaElement.download" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetDownload(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementDownload(
		this.Ref(),
		val.Ref(),
	)
}

// Ping returns the value of property "HTMLAreaElement.ping".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Ping() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementPing(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Ping sets the value of property "HTMLAreaElement.ping" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetPing(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementPing(
		this.Ref(),
		val.Ref(),
	)
}

// Rel returns the value of property "HTMLAreaElement.rel".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Rel() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementRel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Rel sets the value of property "HTMLAreaElement.rel" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetRel(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementRel(
		this.Ref(),
		val.Ref(),
	)
}

// RelList returns the value of property "HTMLAreaElement.relList".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) RelList() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementRelList(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// ReferrerPolicy returns the value of property "HTMLAreaElement.referrerPolicy".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) ReferrerPolicy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementReferrerPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReferrerPolicy sets the value of property "HTMLAreaElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementReferrerPolicy(
		this.Ref(),
		val.Ref(),
	)
}

// NoHref returns the value of property "HTMLAreaElement.noHref".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) NoHref() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementNoHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// NoHref sets the value of property "HTMLAreaElement.noHref" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetNoHref(val bool) bool {
	return js.True == bindings.SetHTMLAreaElementNoHref(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Href returns the value of property "HTMLAreaElement.href".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Href() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Href sets the value of property "HTMLAreaElement.href" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetHref(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementHref(
		this.Ref(),
		val.Ref(),
	)
}

// Origin returns the value of property "HTMLAreaElement.origin".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Origin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol returns the value of property "HTMLAreaElement.protocol".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Protocol() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementProtocol(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Protocol sets the value of property "HTMLAreaElement.protocol" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetProtocol(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementProtocol(
		this.Ref(),
		val.Ref(),
	)
}

// Username returns the value of property "HTMLAreaElement.username".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Username() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementUsername(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Username sets the value of property "HTMLAreaElement.username" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetUsername(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementUsername(
		this.Ref(),
		val.Ref(),
	)
}

// Password returns the value of property "HTMLAreaElement.password".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Password() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementPassword(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Password sets the value of property "HTMLAreaElement.password" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetPassword(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementPassword(
		this.Ref(),
		val.Ref(),
	)
}

// Host returns the value of property "HTMLAreaElement.host".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Host() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementHost(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Host sets the value of property "HTMLAreaElement.host" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetHost(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementHost(
		this.Ref(),
		val.Ref(),
	)
}

// Hostname returns the value of property "HTMLAreaElement.hostname".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Hostname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementHostname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hostname sets the value of property "HTMLAreaElement.hostname" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetHostname(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementHostname(
		this.Ref(),
		val.Ref(),
	)
}

// Port returns the value of property "HTMLAreaElement.port".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Port() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementPort(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Port sets the value of property "HTMLAreaElement.port" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetPort(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementPort(
		this.Ref(),
		val.Ref(),
	)
}

// Pathname returns the value of property "HTMLAreaElement.pathname".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Pathname() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementPathname(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Pathname sets the value of property "HTMLAreaElement.pathname" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetPathname(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementPathname(
		this.Ref(),
		val.Ref(),
	)
}

// Search returns the value of property "HTMLAreaElement.search".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Search() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementSearch(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Search sets the value of property "HTMLAreaElement.search" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetSearch(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementSearch(
		this.Ref(),
		val.Ref(),
	)
}

// Hash returns the value of property "HTMLAreaElement.hash".
//
// The returned bool will be false if there is no such property.
func (this HTMLAreaElement) Hash() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLAreaElementHash(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hash sets the value of property "HTMLAreaElement.hash" to val.
//
// It returns false if the property cannot be set.
func (this HTMLAreaElement) SetHash(val js.String) bool {
	return js.True == bindings.SetHTMLAreaElementHash(
		this.Ref(),
		val.Ref(),
	)
}

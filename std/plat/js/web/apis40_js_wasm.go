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

func NewInstance(module Module, importObject js.Object) (ret Instance) {
	ret.ref = bindings.NewInstanceByInstance(
		module.Ref(),
		importObject.Ref())
	return
}

func NewInstanceByInstance1(module Module) (ret Instance) {
	ret.ref = bindings.NewInstanceByInstance1(
		module.Ref())
	return
}

type Instance struct {
	ref js.Ref
}

func (this Instance) Once() Instance {
	this.Ref().Once()
	return this
}

func (this Instance) Ref() js.Ref {
	return this.ref
}

func (this Instance) FromRef(ref js.Ref) Instance {
	this.ref = ref
	return this
}

func (this Instance) Free() {
	this.Ref().Free()
}

// Exports returns the value of property "Instance.exports".
//
// It returns ok=false if there is no such property.
func (this Instance) Exports() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetInstanceExports(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type InterestGroupBiddingScriptRunnerGlobalScope struct {
	InterestGroupScriptRunnerGlobalScope
}

func (this InterestGroupBiddingScriptRunnerGlobalScope) Once() InterestGroupBiddingScriptRunnerGlobalScope {
	this.Ref().Once()
	return this
}

func (this InterestGroupBiddingScriptRunnerGlobalScope) Ref() js.Ref {
	return this.InterestGroupScriptRunnerGlobalScope.Ref()
}

func (this InterestGroupBiddingScriptRunnerGlobalScope) FromRef(ref js.Ref) InterestGroupBiddingScriptRunnerGlobalScope {
	this.InterestGroupScriptRunnerGlobalScope = this.InterestGroupScriptRunnerGlobalScope.FromRef(ref)
	return this
}

func (this InterestGroupBiddingScriptRunnerGlobalScope) Free() {
	this.Ref().Free()
}

// HasSetBid returns true if the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid" exists.
func (this InterestGroupBiddingScriptRunnerGlobalScope) HasSetBid() bool {
	return js.True == bindings.HasInterestGroupBiddingScriptRunnerGlobalScopeSetBid(
		this.Ref(),
	)
}

// SetBidFunc returns the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid".
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetBidFunc() (fn js.Func[func(generateBidOutput GenerateBidOutput) bool]) {
	return fn.FromRef(
		bindings.InterestGroupBiddingScriptRunnerGlobalScopeSetBidFunc(
			this.Ref(),
		),
	)
}

// SetBid calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid".
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetBid(generateBidOutput GenerateBidOutput) (ret bool) {
	bindings.CallInterestGroupBiddingScriptRunnerGlobalScopeSetBid(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&generateBidOutput),
	)

	return
}

// TrySetBid calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InterestGroupBiddingScriptRunnerGlobalScope) TrySetBid(generateBidOutput GenerateBidOutput) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInterestGroupBiddingScriptRunnerGlobalScopeSetBid(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&generateBidOutput),
	)

	return
}

// HasSetBid1 returns true if the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid" exists.
func (this InterestGroupBiddingScriptRunnerGlobalScope) HasSetBid1() bool {
	return js.True == bindings.HasInterestGroupBiddingScriptRunnerGlobalScopeSetBid1(
		this.Ref(),
	)
}

// SetBid1Func returns the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid".
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetBid1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.InterestGroupBiddingScriptRunnerGlobalScopeSetBid1Func(
			this.Ref(),
		),
	)
}

// SetBid1 calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid".
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetBid1() (ret bool) {
	bindings.CallInterestGroupBiddingScriptRunnerGlobalScopeSetBid1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySetBid1 calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InterestGroupBiddingScriptRunnerGlobalScope) TrySetBid1() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInterestGroupBiddingScriptRunnerGlobalScopeSetBid1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetPriority returns true if the method "InterestGroupBiddingScriptRunnerGlobalScope.setPriority" exists.
func (this InterestGroupBiddingScriptRunnerGlobalScope) HasSetPriority() bool {
	return js.True == bindings.HasInterestGroupBiddingScriptRunnerGlobalScopeSetPriority(
		this.Ref(),
	)
}

// SetPriorityFunc returns the method "InterestGroupBiddingScriptRunnerGlobalScope.setPriority".
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPriorityFunc() (fn js.Func[func(priority float64)]) {
	return fn.FromRef(
		bindings.InterestGroupBiddingScriptRunnerGlobalScopeSetPriorityFunc(
			this.Ref(),
		),
	)
}

// SetPriority calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setPriority".
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPriority(priority float64) (ret js.Void) {
	bindings.CallInterestGroupBiddingScriptRunnerGlobalScopeSetPriority(
		this.Ref(), js.Pointer(&ret),
		float64(priority),
	)

	return
}

// TrySetPriority calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setPriority"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InterestGroupBiddingScriptRunnerGlobalScope) TrySetPriority(priority float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInterestGroupBiddingScriptRunnerGlobalScopeSetPriority(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(priority),
	)

	return
}

// HasSetPrioritySignalsOverride returns true if the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride" exists.
func (this InterestGroupBiddingScriptRunnerGlobalScope) HasSetPrioritySignalsOverride() bool {
	return js.True == bindings.HasInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride(
		this.Ref(),
	)
}

// SetPrioritySignalsOverrideFunc returns the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride".
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPrioritySignalsOverrideFunc() (fn js.Func[func(key js.String, priority float64)]) {
	return fn.FromRef(
		bindings.InterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverrideFunc(
			this.Ref(),
		),
	)
}

// SetPrioritySignalsOverride calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride".
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPrioritySignalsOverride(key js.String, priority float64) (ret js.Void) {
	bindings.CallInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
		float64(priority),
	)

	return
}

// TrySetPrioritySignalsOverride calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InterestGroupBiddingScriptRunnerGlobalScope) TrySetPrioritySignalsOverride(key js.String, priority float64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
		float64(priority),
	)

	return
}

// HasSetPrioritySignalsOverride1 returns true if the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride" exists.
func (this InterestGroupBiddingScriptRunnerGlobalScope) HasSetPrioritySignalsOverride1() bool {
	return js.True == bindings.HasInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1(
		this.Ref(),
	)
}

// SetPrioritySignalsOverride1Func returns the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride".
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPrioritySignalsOverride1Func() (fn js.Func[func(key js.String)]) {
	return fn.FromRef(
		bindings.InterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1Func(
			this.Ref(),
		),
	)
}

// SetPrioritySignalsOverride1 calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride".
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPrioritySignalsOverride1(key js.String) (ret js.Void) {
	bindings.CallInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TrySetPrioritySignalsOverride1 calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InterestGroupBiddingScriptRunnerGlobalScope) TrySetPrioritySignalsOverride1(key js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

type InterestGroupReportingScriptRunnerGlobalScope struct {
	InterestGroupScriptRunnerGlobalScope
}

func (this InterestGroupReportingScriptRunnerGlobalScope) Once() InterestGroupReportingScriptRunnerGlobalScope {
	this.Ref().Once()
	return this
}

func (this InterestGroupReportingScriptRunnerGlobalScope) Ref() js.Ref {
	return this.InterestGroupScriptRunnerGlobalScope.Ref()
}

func (this InterestGroupReportingScriptRunnerGlobalScope) FromRef(ref js.Ref) InterestGroupReportingScriptRunnerGlobalScope {
	this.InterestGroupScriptRunnerGlobalScope = this.InterestGroupScriptRunnerGlobalScope.FromRef(ref)
	return this
}

func (this InterestGroupReportingScriptRunnerGlobalScope) Free() {
	this.Ref().Free()
}

// HasSendReportTo returns true if the method "InterestGroupReportingScriptRunnerGlobalScope.sendReportTo" exists.
func (this InterestGroupReportingScriptRunnerGlobalScope) HasSendReportTo() bool {
	return js.True == bindings.HasInterestGroupReportingScriptRunnerGlobalScopeSendReportTo(
		this.Ref(),
	)
}

// SendReportToFunc returns the method "InterestGroupReportingScriptRunnerGlobalScope.sendReportTo".
func (this InterestGroupReportingScriptRunnerGlobalScope) SendReportToFunc() (fn js.Func[func(url js.String)]) {
	return fn.FromRef(
		bindings.InterestGroupReportingScriptRunnerGlobalScopeSendReportToFunc(
			this.Ref(),
		),
	)
}

// SendReportTo calls the method "InterestGroupReportingScriptRunnerGlobalScope.sendReportTo".
func (this InterestGroupReportingScriptRunnerGlobalScope) SendReportTo(url js.String) (ret js.Void) {
	bindings.CallInterestGroupReportingScriptRunnerGlobalScopeSendReportTo(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TrySendReportTo calls the method "InterestGroupReportingScriptRunnerGlobalScope.sendReportTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InterestGroupReportingScriptRunnerGlobalScope) TrySendReportTo(url js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInterestGroupReportingScriptRunnerGlobalScopeSendReportTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasRegisterAdBeacon returns true if the method "InterestGroupReportingScriptRunnerGlobalScope.registerAdBeacon" exists.
func (this InterestGroupReportingScriptRunnerGlobalScope) HasRegisterAdBeacon() bool {
	return js.True == bindings.HasInterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeacon(
		this.Ref(),
	)
}

// RegisterAdBeaconFunc returns the method "InterestGroupReportingScriptRunnerGlobalScope.registerAdBeacon".
func (this InterestGroupReportingScriptRunnerGlobalScope) RegisterAdBeaconFunc() (fn js.Func[func(mapping js.Record[js.String])]) {
	return fn.FromRef(
		bindings.InterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeaconFunc(
			this.Ref(),
		),
	)
}

// RegisterAdBeacon calls the method "InterestGroupReportingScriptRunnerGlobalScope.registerAdBeacon".
func (this InterestGroupReportingScriptRunnerGlobalScope) RegisterAdBeacon(mapping js.Record[js.String]) (ret js.Void) {
	bindings.CallInterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeacon(
		this.Ref(), js.Pointer(&ret),
		mapping.Ref(),
	)

	return
}

// TryRegisterAdBeacon calls the method "InterestGroupReportingScriptRunnerGlobalScope.registerAdBeacon"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InterestGroupReportingScriptRunnerGlobalScope) TryRegisterAdBeacon(mapping js.Record[js.String]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeacon(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		mapping.Ref(),
	)

	return
}

type InterestGroupScoringScriptRunnerGlobalScope struct {
	InterestGroupScriptRunnerGlobalScope
}

func (this InterestGroupScoringScriptRunnerGlobalScope) Once() InterestGroupScoringScriptRunnerGlobalScope {
	this.Ref().Once()
	return this
}

func (this InterestGroupScoringScriptRunnerGlobalScope) Ref() js.Ref {
	return this.InterestGroupScriptRunnerGlobalScope.Ref()
}

func (this InterestGroupScoringScriptRunnerGlobalScope) FromRef(ref js.Ref) InterestGroupScoringScriptRunnerGlobalScope {
	this.InterestGroupScriptRunnerGlobalScope = this.InterestGroupScriptRunnerGlobalScope.FromRef(ref)
	return this
}

func (this InterestGroupScoringScriptRunnerGlobalScope) Free() {
	this.Ref().Free()
}

type InterestGroupScriptRunnerGlobalScope struct {
	ref js.Ref
}

func (this InterestGroupScriptRunnerGlobalScope) Once() InterestGroupScriptRunnerGlobalScope {
	this.Ref().Once()
	return this
}

func (this InterestGroupScriptRunnerGlobalScope) Ref() js.Ref {
	return this.ref
}

func (this InterestGroupScriptRunnerGlobalScope) FromRef(ref js.Ref) InterestGroupScriptRunnerGlobalScope {
	this.ref = ref
	return this
}

func (this InterestGroupScriptRunnerGlobalScope) Free() {
	this.Ref().Free()
}

type IntersectionObserverCallbackFunc func(this js.Ref, entries js.Array[IntersectionObserverEntry], observer IntersectionObserver) js.Ref

func (fn IntersectionObserverCallbackFunc) Register() js.Func[func(entries js.Array[IntersectionObserverEntry], observer IntersectionObserver)] {
	return js.RegisterCallback[func(entries js.Array[IntersectionObserverEntry], observer IntersectionObserver)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IntersectionObserverCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[IntersectionObserverEntry]{}.FromRef(args[0+1]),
		IntersectionObserver{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IntersectionObserverCallback[T any] struct {
	Fn  func(arg T, this js.Ref, entries js.Array[IntersectionObserverEntry], observer IntersectionObserver) js.Ref
	Arg T
}

func (cb *IntersectionObserverCallback[T]) Register() js.Func[func(entries js.Array[IntersectionObserverEntry], observer IntersectionObserver)] {
	return js.RegisterCallback[func(entries js.Array[IntersectionObserverEntry], observer IntersectionObserver)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IntersectionObserverCallback[T]) DispatchCallback(
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

		js.Array[IntersectionObserverEntry]{}.FromRef(args[0+1]),
		IntersectionObserver{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IntersectionObserverEntryInit struct {
	// Time is "IntersectionObserverEntryInit.time"
	//
	// Required
	Time DOMHighResTimeStamp
	// RootBounds is "IntersectionObserverEntryInit.rootBounds"
	//
	// Required
	//
	// NOTE: RootBounds.FFI_USE MUST be set to true to get RootBounds used.
	RootBounds DOMRectInit
	// BoundingClientRect is "IntersectionObserverEntryInit.boundingClientRect"
	//
	// Required
	//
	// NOTE: BoundingClientRect.FFI_USE MUST be set to true to get BoundingClientRect used.
	BoundingClientRect DOMRectInit
	// IntersectionRect is "IntersectionObserverEntryInit.intersectionRect"
	//
	// Required
	//
	// NOTE: IntersectionRect.FFI_USE MUST be set to true to get IntersectionRect used.
	IntersectionRect DOMRectInit
	// IsIntersecting is "IntersectionObserverEntryInit.isIntersecting"
	//
	// Required
	IsIntersecting bool
	// IntersectionRatio is "IntersectionObserverEntryInit.intersectionRatio"
	//
	// Required
	IntersectionRatio float64
	// Target is "IntersectionObserverEntryInit.target"
	//
	// Required
	Target Element

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IntersectionObserverEntryInit with all fields set.
func (p IntersectionObserverEntryInit) FromRef(ref js.Ref) IntersectionObserverEntryInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IntersectionObserverEntryInit in the application heap.
func (p IntersectionObserverEntryInit) New() js.Ref {
	return bindings.IntersectionObserverEntryInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IntersectionObserverEntryInit) UpdateFrom(ref js.Ref) {
	bindings.IntersectionObserverEntryInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IntersectionObserverEntryInit) Update(ref js.Ref) {
	bindings.IntersectionObserverEntryInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewIntersectionObserverEntry(intersectionObserverEntryInit IntersectionObserverEntryInit) (ret IntersectionObserverEntry) {
	ret.ref = bindings.NewIntersectionObserverEntryByIntersectionObserverEntry(
		js.Pointer(&intersectionObserverEntryInit))
	return
}

type IntersectionObserverEntry struct {
	ref js.Ref
}

func (this IntersectionObserverEntry) Once() IntersectionObserverEntry {
	this.Ref().Once()
	return this
}

func (this IntersectionObserverEntry) Ref() js.Ref {
	return this.ref
}

func (this IntersectionObserverEntry) FromRef(ref js.Ref) IntersectionObserverEntry {
	this.ref = ref
	return this
}

func (this IntersectionObserverEntry) Free() {
	this.Ref().Free()
}

// Time returns the value of property "IntersectionObserverEntry.time".
//
// It returns ok=false if there is no such property.
func (this IntersectionObserverEntry) Time() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetIntersectionObserverEntryTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RootBounds returns the value of property "IntersectionObserverEntry.rootBounds".
//
// It returns ok=false if there is no such property.
func (this IntersectionObserverEntry) RootBounds() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetIntersectionObserverEntryRootBounds(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// BoundingClientRect returns the value of property "IntersectionObserverEntry.boundingClientRect".
//
// It returns ok=false if there is no such property.
func (this IntersectionObserverEntry) BoundingClientRect() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetIntersectionObserverEntryBoundingClientRect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IntersectionRect returns the value of property "IntersectionObserverEntry.intersectionRect".
//
// It returns ok=false if there is no such property.
func (this IntersectionObserverEntry) IntersectionRect() (ret DOMRectReadOnly, ok bool) {
	ok = js.True == bindings.GetIntersectionObserverEntryIntersectionRect(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsIntersecting returns the value of property "IntersectionObserverEntry.isIntersecting".
//
// It returns ok=false if there is no such property.
func (this IntersectionObserverEntry) IsIntersecting() (ret bool, ok bool) {
	ok = js.True == bindings.GetIntersectionObserverEntryIsIntersecting(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IntersectionRatio returns the value of property "IntersectionObserverEntry.intersectionRatio".
//
// It returns ok=false if there is no such property.
func (this IntersectionObserverEntry) IntersectionRatio() (ret float64, ok bool) {
	ok = js.True == bindings.GetIntersectionObserverEntryIntersectionRatio(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Target returns the value of property "IntersectionObserverEntry.target".
//
// It returns ok=false if there is no such property.
func (this IntersectionObserverEntry) Target() (ret Element, ok bool) {
	ok = js.True == bindings.GetIntersectionObserverEntryTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type OneOf_Element_Document struct {
	ref js.Ref
}

func (x OneOf_Element_Document) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Element_Document) Free() {
	x.ref.Free()
}

func (x OneOf_Element_Document) FromRef(ref js.Ref) OneOf_Element_Document {
	return OneOf_Element_Document{
		ref: ref,
	}
}

func (x OneOf_Element_Document) Element() Element {
	return Element{}.FromRef(x.ref)
}

func (x OneOf_Element_Document) Document() Document {
	return Document{}.FromRef(x.ref)
}

type OneOf_Float64_ArrayFloat64 struct {
	ref js.Ref
}

func (x OneOf_Float64_ArrayFloat64) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_ArrayFloat64) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_ArrayFloat64) FromRef(ref js.Ref) OneOf_Float64_ArrayFloat64 {
	return OneOf_Float64_ArrayFloat64{
		ref: ref,
	}
}

func (x OneOf_Float64_ArrayFloat64) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_ArrayFloat64) ArrayFloat64() js.Array[float64] {
	return js.Array[float64]{}.FromRef(x.ref)
}

type IntersectionObserverInit struct {
	// Root is "IntersectionObserverInit.root"
	//
	// Optional, defaults to null.
	Root OneOf_Element_Document
	// RootMargin is "IntersectionObserverInit.rootMargin"
	//
	// Optional, defaults to "0px".
	RootMargin js.String
	// Threshold is "IntersectionObserverInit.threshold"
	//
	// Optional, defaults to 0.
	Threshold OneOf_Float64_ArrayFloat64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IntersectionObserverInit with all fields set.
func (p IntersectionObserverInit) FromRef(ref js.Ref) IntersectionObserverInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IntersectionObserverInit in the application heap.
func (p IntersectionObserverInit) New() js.Ref {
	return bindings.IntersectionObserverInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IntersectionObserverInit) UpdateFrom(ref js.Ref) {
	bindings.IntersectionObserverInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IntersectionObserverInit) Update(ref js.Ref) {
	bindings.IntersectionObserverInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewIntersectionObserver(callback js.Func[func(entries js.Array[IntersectionObserverEntry], observer IntersectionObserver)], options IntersectionObserverInit) (ret IntersectionObserver) {
	ret.ref = bindings.NewIntersectionObserverByIntersectionObserver(
		callback.Ref(),
		js.Pointer(&options))
	return
}

func NewIntersectionObserverByIntersectionObserver1(callback js.Func[func(entries js.Array[IntersectionObserverEntry], observer IntersectionObserver)]) (ret IntersectionObserver) {
	ret.ref = bindings.NewIntersectionObserverByIntersectionObserver1(
		callback.Ref())
	return
}

type IntersectionObserver struct {
	ref js.Ref
}

func (this IntersectionObserver) Once() IntersectionObserver {
	this.Ref().Once()
	return this
}

func (this IntersectionObserver) Ref() js.Ref {
	return this.ref
}

func (this IntersectionObserver) FromRef(ref js.Ref) IntersectionObserver {
	this.ref = ref
	return this
}

func (this IntersectionObserver) Free() {
	this.Ref().Free()
}

// Root returns the value of property "IntersectionObserver.root".
//
// It returns ok=false if there is no such property.
func (this IntersectionObserver) Root() (ret OneOf_Element_Document, ok bool) {
	ok = js.True == bindings.GetIntersectionObserverRoot(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// RootMargin returns the value of property "IntersectionObserver.rootMargin".
//
// It returns ok=false if there is no such property.
func (this IntersectionObserver) RootMargin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetIntersectionObserverRootMargin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Thresholds returns the value of property "IntersectionObserver.thresholds".
//
// It returns ok=false if there is no such property.
func (this IntersectionObserver) Thresholds() (ret js.FrozenArray[float64], ok bool) {
	ok = js.True == bindings.GetIntersectionObserverThresholds(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasObserve returns true if the method "IntersectionObserver.observe" exists.
func (this IntersectionObserver) HasObserve() bool {
	return js.True == bindings.HasIntersectionObserverObserve(
		this.Ref(),
	)
}

// ObserveFunc returns the method "IntersectionObserver.observe".
func (this IntersectionObserver) ObserveFunc() (fn js.Func[func(target Element)]) {
	return fn.FromRef(
		bindings.IntersectionObserverObserveFunc(
			this.Ref(),
		),
	)
}

// Observe calls the method "IntersectionObserver.observe".
func (this IntersectionObserver) Observe(target Element) (ret js.Void) {
	bindings.CallIntersectionObserverObserve(
		this.Ref(), js.Pointer(&ret),
		target.Ref(),
	)

	return
}

// TryObserve calls the method "IntersectionObserver.observe"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IntersectionObserver) TryObserve(target Element) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIntersectionObserverObserve(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		target.Ref(),
	)

	return
}

// HasUnobserve returns true if the method "IntersectionObserver.unobserve" exists.
func (this IntersectionObserver) HasUnobserve() bool {
	return js.True == bindings.HasIntersectionObserverUnobserve(
		this.Ref(),
	)
}

// UnobserveFunc returns the method "IntersectionObserver.unobserve".
func (this IntersectionObserver) UnobserveFunc() (fn js.Func[func(target Element)]) {
	return fn.FromRef(
		bindings.IntersectionObserverUnobserveFunc(
			this.Ref(),
		),
	)
}

// Unobserve calls the method "IntersectionObserver.unobserve".
func (this IntersectionObserver) Unobserve(target Element) (ret js.Void) {
	bindings.CallIntersectionObserverUnobserve(
		this.Ref(), js.Pointer(&ret),
		target.Ref(),
	)

	return
}

// TryUnobserve calls the method "IntersectionObserver.unobserve"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IntersectionObserver) TryUnobserve(target Element) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIntersectionObserverUnobserve(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		target.Ref(),
	)

	return
}

// HasDisconnect returns true if the method "IntersectionObserver.disconnect" exists.
func (this IntersectionObserver) HasDisconnect() bool {
	return js.True == bindings.HasIntersectionObserverDisconnect(
		this.Ref(),
	)
}

// DisconnectFunc returns the method "IntersectionObserver.disconnect".
func (this IntersectionObserver) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.IntersectionObserverDisconnectFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "IntersectionObserver.disconnect".
func (this IntersectionObserver) Disconnect() (ret js.Void) {
	bindings.CallIntersectionObserverDisconnect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDisconnect calls the method "IntersectionObserver.disconnect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IntersectionObserver) TryDisconnect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIntersectionObserverDisconnect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasTakeRecords returns true if the method "IntersectionObserver.takeRecords" exists.
func (this IntersectionObserver) HasTakeRecords() bool {
	return js.True == bindings.HasIntersectionObserverTakeRecords(
		this.Ref(),
	)
}

// TakeRecordsFunc returns the method "IntersectionObserver.takeRecords".
func (this IntersectionObserver) TakeRecordsFunc() (fn js.Func[func() js.Array[IntersectionObserverEntry]]) {
	return fn.FromRef(
		bindings.IntersectionObserverTakeRecordsFunc(
			this.Ref(),
		),
	)
}

// TakeRecords calls the method "IntersectionObserver.takeRecords".
func (this IntersectionObserver) TakeRecords() (ret js.Array[IntersectionObserverEntry]) {
	bindings.CallIntersectionObserverTakeRecords(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryTakeRecords calls the method "IntersectionObserver.takeRecords"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IntersectionObserver) TryTakeRecords() (ret js.Array[IntersectionObserverEntry], exception js.Any, ok bool) {
	ok = js.True == bindings.TryIntersectionObserverTakeRecords(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type InterventionReportBody struct {
	ReportBody
}

func (this InterventionReportBody) Once() InterventionReportBody {
	this.Ref().Once()
	return this
}

func (this InterventionReportBody) Ref() js.Ref {
	return this.ReportBody.Ref()
}

func (this InterventionReportBody) FromRef(ref js.Ref) InterventionReportBody {
	this.ReportBody = this.ReportBody.FromRef(ref)
	return this
}

func (this InterventionReportBody) Free() {
	this.Ref().Free()
}

// Id returns the value of property "InterventionReportBody.id".
//
// It returns ok=false if there is no such property.
func (this InterventionReportBody) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetInterventionReportBodyId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Message returns the value of property "InterventionReportBody.message".
//
// It returns ok=false if there is no such property.
func (this InterventionReportBody) Message() (ret js.String, ok bool) {
	ok = js.True == bindings.GetInterventionReportBodyMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SourceFile returns the value of property "InterventionReportBody.sourceFile".
//
// It returns ok=false if there is no such property.
func (this InterventionReportBody) SourceFile() (ret js.String, ok bool) {
	ok = js.True == bindings.GetInterventionReportBodySourceFile(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LineNumber returns the value of property "InterventionReportBody.lineNumber".
//
// It returns ok=false if there is no such property.
func (this InterventionReportBody) LineNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetInterventionReportBodyLineNumber(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ColumnNumber returns the value of property "InterventionReportBody.columnNumber".
//
// It returns ok=false if there is no such property.
func (this InterventionReportBody) ColumnNumber() (ret uint32, ok bool) {
	ok = js.True == bindings.GetInterventionReportBodyColumnNumber(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasToJSON returns true if the method "InterventionReportBody.toJSON" exists.
func (this InterventionReportBody) HasToJSON() bool {
	return js.True == bindings.HasInterventionReportBodyToJSON(
		this.Ref(),
	)
}

// ToJSONFunc returns the method "InterventionReportBody.toJSON".
func (this InterventionReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.InterventionReportBodyToJSONFunc(
			this.Ref(),
		),
	)
}

// ToJSON calls the method "InterventionReportBody.toJSON".
func (this InterventionReportBody) ToJSON() (ret js.Object) {
	bindings.CallInterventionReportBodyToJSON(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryToJSON calls the method "InterventionReportBody.toJSON"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this InterventionReportBody) TryToJSON() (ret js.Object, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInterventionReportBodyToJSON(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type IntrinsicSizesResultOptions struct {
	// MaxContentSize is "IntrinsicSizesResultOptions.maxContentSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxContentSize MUST be set to true to make this field effective.
	MaxContentSize float64
	// MinContentSize is "IntrinsicSizesResultOptions.minContentSize"
	//
	// Optional
	//
	// NOTE: FFI_USE_MinContentSize MUST be set to true to make this field effective.
	MinContentSize float64

	FFI_USE_MaxContentSize bool // for MaxContentSize.
	FFI_USE_MinContentSize bool // for MinContentSize.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IntrinsicSizesResultOptions with all fields set.
func (p IntrinsicSizesResultOptions) FromRef(ref js.Ref) IntrinsicSizesResultOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IntrinsicSizesResultOptions in the application heap.
func (p IntrinsicSizesResultOptions) New() js.Ref {
	return bindings.IntrinsicSizesResultOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IntrinsicSizesResultOptions) UpdateFrom(ref js.Ref) {
	bindings.IntrinsicSizesResultOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IntrinsicSizesResultOptions) Update(ref js.Ref) {
	bindings.IntrinsicSizesResultOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type JsonLd struct {
	ref js.Ref
}

func (this JsonLd) Once() JsonLd {
	this.Ref().Once()
	return this
}

func (this JsonLd) Ref() js.Ref {
	return this.ref
}

func (this JsonLd) FromRef(ref js.Ref) JsonLd {
	this.ref = ref
	return this
}

func (this JsonLd) Free() {
	this.Ref().Free()
}

type OneOf_RecordAny_String struct {
	ref js.Ref
}

func (x OneOf_RecordAny_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_RecordAny_String) Free() {
	x.ref.Free()
}

func (x OneOf_RecordAny_String) FromRef(ref js.Ref) OneOf_RecordAny_String {
	return OneOf_RecordAny_String{
		ref: ref,
	}
}

func (x OneOf_RecordAny_String) RecordAny() js.Record[js.Any] {
	return js.Record[js.Any]{}.FromRef(x.ref)
}

func (x OneOf_RecordAny_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type OneOf_RecordAny_ArrayOneOf_RecordAny_String_String struct {
	ref js.Ref
}

func (x OneOf_RecordAny_ArrayOneOf_RecordAny_String_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_RecordAny_ArrayOneOf_RecordAny_String_String) Free() {
	x.ref.Free()
}

func (x OneOf_RecordAny_ArrayOneOf_RecordAny_String_String) FromRef(ref js.Ref) OneOf_RecordAny_ArrayOneOf_RecordAny_String_String {
	return OneOf_RecordAny_ArrayOneOf_RecordAny_String_String{
		ref: ref,
	}
}

func (x OneOf_RecordAny_ArrayOneOf_RecordAny_String_String) RecordAny() js.Record[js.Any] {
	return js.Record[js.Any]{}.FromRef(x.ref)
}

func (x OneOf_RecordAny_ArrayOneOf_RecordAny_String_String) ArrayOneOf_RecordAny_String() js.Array[OneOf_RecordAny_String] {
	return js.Array[OneOf_RecordAny_String]{}.FromRef(x.ref)
}

func (x OneOf_RecordAny_ArrayOneOf_RecordAny_String_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type JsonLdContext = OneOf_RecordAny_ArrayOneOf_RecordAny_String_String

type JsonLdEmbed uint32

const (
	_ JsonLdEmbed = iota

	JsonLdEmbed__ALWAYS
	JsonLdEmbed__ONCE
	JsonLdEmbed__NEVER
)

func (JsonLdEmbed) FromRef(str js.Ref) JsonLdEmbed {
	return JsonLdEmbed(bindings.ConstOfJsonLdEmbed(str))
}

func (x JsonLdEmbed) String() (string, bool) {
	switch x {
	case JsonLdEmbed__ALWAYS:
		return "@always", true
	case JsonLdEmbed__ONCE:
		return "@once", true
	case JsonLdEmbed__NEVER:
		return "@never", true
	default:
		return "", false
	}
}

type JsonLdErrorCode uint32

const (
	_ JsonLdErrorCode = iota

	JsonLdErrorCode_COLLIDING_KEYWORDS
	JsonLdErrorCode_CONFLICTING_INDEXES
	JsonLdErrorCode_CONTEXT_OVERFLOW
	JsonLdErrorCode_CYCLIC_IRI_MAPPING
	JsonLdErrorCode_INVALID__ID_VALUE
	JsonLdErrorCode_INVALID__IMPORT_VALUE
	JsonLdErrorCode_INVALID__INCLUDED_VALUE
	JsonLdErrorCode_INVALID__INDEX_VALUE
	JsonLdErrorCode_INVALID__NEST_VALUE
	JsonLdErrorCode_INVALID__PREFIX_VALUE
	JsonLdErrorCode_INVALID__PROPAGATE_VALUE
	JsonLdErrorCode_INVALID__PROTECTED_VALUE
	JsonLdErrorCode_INVALID__REVERSE_VALUE
	JsonLdErrorCode_INVALID__VERSION_VALUE
	JsonLdErrorCode_INVALID_BASE_DIRECTION
	JsonLdErrorCode_INVALID_BASE_IRI
	JsonLdErrorCode_INVALID_CONTAINER_MAPPING
	JsonLdErrorCode_INVALID_CONTEXT_ENTRY
	JsonLdErrorCode_INVALID_CONTEXT_NULLIFICATION
	JsonLdErrorCode_INVALID_DEFAULT_LANGUAGE
	JsonLdErrorCode_INVALID_IRI_MAPPING
	JsonLdErrorCode_INVALID_JSON_LITERAL
	JsonLdErrorCode_INVALID_KEYWORD_ALIAS
	JsonLdErrorCode_INVALID_LANGUAGE_MAP_VALUE
	JsonLdErrorCode_INVALID_LANGUAGE_MAPPING
	JsonLdErrorCode_INVALID_LANGUAGE_TAGGED_STRING
	JsonLdErrorCode_INVALID_LANGUAGE_TAGGED_VALUE
	JsonLdErrorCode_INVALID_LOCAL_CONTEXT
	JsonLdErrorCode_INVALID_REMOTE_CONTEXT
	JsonLdErrorCode_INVALID_REVERSE_PROPERTY_MAP
	JsonLdErrorCode_INVALID_REVERSE_PROPERTY_VALUE
	JsonLdErrorCode_INVALID_REVERSE_PROPERTY
	JsonLdErrorCode_INVALID_SCOPED_CONTEXT
	JsonLdErrorCode_INVALID_SCRIPT_ELEMENT
	JsonLdErrorCode_INVALID_SET_OR_LIST_OBJECT
	JsonLdErrorCode_INVALID_TERM_DEFINITION
	JsonLdErrorCode_INVALID_TYPE_MAPPING
	JsonLdErrorCode_INVALID_TYPE_VALUE
	JsonLdErrorCode_INVALID_TYPED_VALUE
	JsonLdErrorCode_INVALID_VALUE_OBJECT_VALUE
	JsonLdErrorCode_INVALID_VALUE_OBJECT
	JsonLdErrorCode_INVALID_VOCAB_MAPPING
	JsonLdErrorCode_IRI_CONFUSED_WITH_PREFIX
	JsonLdErrorCode_KEYWORD_REDEFINITION
	JsonLdErrorCode_LOADING_DOCUMENT_FAILED
	JsonLdErrorCode_LOADING_REMOTE_CONTEXT_FAILED
	JsonLdErrorCode_MULTIPLE_CONTEXT_LINK_HEADERS
	JsonLdErrorCode_PROCESSING_MODE_CONFLICT
	JsonLdErrorCode_PROTECTED_TERM_REDEFINITION
)

func (JsonLdErrorCode) FromRef(str js.Ref) JsonLdErrorCode {
	return JsonLdErrorCode(bindings.ConstOfJsonLdErrorCode(str))
}

func (x JsonLdErrorCode) String() (string, bool) {
	switch x {
	case JsonLdErrorCode_COLLIDING_KEYWORDS:
		return "colliding keywords", true
	case JsonLdErrorCode_CONFLICTING_INDEXES:
		return "conflicting indexes", true
	case JsonLdErrorCode_CONTEXT_OVERFLOW:
		return "context overflow", true
	case JsonLdErrorCode_CYCLIC_IRI_MAPPING:
		return "cyclic IRI mapping", true
	case JsonLdErrorCode_INVALID__ID_VALUE:
		return "invalid @id value", true
	case JsonLdErrorCode_INVALID__IMPORT_VALUE:
		return "invalid @import value", true
	case JsonLdErrorCode_INVALID__INCLUDED_VALUE:
		return "invalid @included value", true
	case JsonLdErrorCode_INVALID__INDEX_VALUE:
		return "invalid @index value", true
	case JsonLdErrorCode_INVALID__NEST_VALUE:
		return "invalid @nest value", true
	case JsonLdErrorCode_INVALID__PREFIX_VALUE:
		return "invalid @prefix value", true
	case JsonLdErrorCode_INVALID__PROPAGATE_VALUE:
		return "invalid @propagate value", true
	case JsonLdErrorCode_INVALID__PROTECTED_VALUE:
		return "invalid @protected value", true
	case JsonLdErrorCode_INVALID__REVERSE_VALUE:
		return "invalid @reverse value", true
	case JsonLdErrorCode_INVALID__VERSION_VALUE:
		return "invalid @version value", true
	case JsonLdErrorCode_INVALID_BASE_DIRECTION:
		return "invalid base direction", true
	case JsonLdErrorCode_INVALID_BASE_IRI:
		return "invalid base IRI", true
	case JsonLdErrorCode_INVALID_CONTAINER_MAPPING:
		return "invalid container mapping", true
	case JsonLdErrorCode_INVALID_CONTEXT_ENTRY:
		return "invalid context entry", true
	case JsonLdErrorCode_INVALID_CONTEXT_NULLIFICATION:
		return "invalid context nullification", true
	case JsonLdErrorCode_INVALID_DEFAULT_LANGUAGE:
		return "invalid default language", true
	case JsonLdErrorCode_INVALID_IRI_MAPPING:
		return "invalid IRI mapping", true
	case JsonLdErrorCode_INVALID_JSON_LITERAL:
		return "invalid JSON literal", true
	case JsonLdErrorCode_INVALID_KEYWORD_ALIAS:
		return "invalid keyword alias", true
	case JsonLdErrorCode_INVALID_LANGUAGE_MAP_VALUE:
		return "invalid language map value", true
	case JsonLdErrorCode_INVALID_LANGUAGE_MAPPING:
		return "invalid language mapping", true
	case JsonLdErrorCode_INVALID_LANGUAGE_TAGGED_STRING:
		return "invalid language-tagged string", true
	case JsonLdErrorCode_INVALID_LANGUAGE_TAGGED_VALUE:
		return "invalid language-tagged value", true
	case JsonLdErrorCode_INVALID_LOCAL_CONTEXT:
		return "invalid local context", true
	case JsonLdErrorCode_INVALID_REMOTE_CONTEXT:
		return "invalid remote context", true
	case JsonLdErrorCode_INVALID_REVERSE_PROPERTY_MAP:
		return "invalid reverse property map", true
	case JsonLdErrorCode_INVALID_REVERSE_PROPERTY_VALUE:
		return "invalid reverse property value", true
	case JsonLdErrorCode_INVALID_REVERSE_PROPERTY:
		return "invalid reverse property", true
	case JsonLdErrorCode_INVALID_SCOPED_CONTEXT:
		return "invalid scoped context", true
	case JsonLdErrorCode_INVALID_SCRIPT_ELEMENT:
		return "invalid script element", true
	case JsonLdErrorCode_INVALID_SET_OR_LIST_OBJECT:
		return "invalid set or list object", true
	case JsonLdErrorCode_INVALID_TERM_DEFINITION:
		return "invalid term definition", true
	case JsonLdErrorCode_INVALID_TYPE_MAPPING:
		return "invalid type mapping", true
	case JsonLdErrorCode_INVALID_TYPE_VALUE:
		return "invalid type value", true
	case JsonLdErrorCode_INVALID_TYPED_VALUE:
		return "invalid typed value", true
	case JsonLdErrorCode_INVALID_VALUE_OBJECT_VALUE:
		return "invalid value object value", true
	case JsonLdErrorCode_INVALID_VALUE_OBJECT:
		return "invalid value object", true
	case JsonLdErrorCode_INVALID_VOCAB_MAPPING:
		return "invalid vocab mapping", true
	case JsonLdErrorCode_IRI_CONFUSED_WITH_PREFIX:
		return "IRI confused with prefix", true
	case JsonLdErrorCode_KEYWORD_REDEFINITION:
		return "keyword redefinition", true
	case JsonLdErrorCode_LOADING_DOCUMENT_FAILED:
		return "loading document failed", true
	case JsonLdErrorCode_LOADING_REMOTE_CONTEXT_FAILED:
		return "loading remote context failed", true
	case JsonLdErrorCode_MULTIPLE_CONTEXT_LINK_HEADERS:
		return "multiple context link headers", true
	case JsonLdErrorCode_PROCESSING_MODE_CONFLICT:
		return "processing mode conflict", true
	case JsonLdErrorCode_PROTECTED_TERM_REDEFINITION:
		return "protected term redefinition", true
	default:
		return "", false
	}
}

type JsonLdError struct {
	// Code is "JsonLdError.code"
	//
	// Optional
	Code JsonLdErrorCode
	// Message is "JsonLdError.message"
	//
	// Optional, defaults to null.
	Message js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a JsonLdError with all fields set.
func (p JsonLdError) FromRef(ref js.Ref) JsonLdError {
	p.UpdateFrom(ref)
	return p
}

// New creates a new JsonLdError in the application heap.
func (p JsonLdError) New() js.Ref {
	return bindings.JsonLdErrorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p JsonLdError) UpdateFrom(ref js.Ref) {
	bindings.JsonLdErrorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p JsonLdError) Update(ref js.Ref) {
	bindings.JsonLdErrorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type JsonLdFramingErrorCode uint32

const (
	_ JsonLdFramingErrorCode = iota

	JsonLdFramingErrorCode_INVALID_FRAME
	JsonLdFramingErrorCode_INVALID__EMBED_VALUE
)

func (JsonLdFramingErrorCode) FromRef(str js.Ref) JsonLdFramingErrorCode {
	return JsonLdFramingErrorCode(bindings.ConstOfJsonLdFramingErrorCode(str))
}

func (x JsonLdFramingErrorCode) String() (string, bool) {
	switch x {
	case JsonLdFramingErrorCode_INVALID_FRAME:
		return "invalid frame", true
	case JsonLdFramingErrorCode_INVALID__EMBED_VALUE:
		return "invalid @embed value", true
	default:
		return "", false
	}
}

type JsonLdFramingError struct {
	// Code is "JsonLdFramingError.code"
	//
	// Optional
	Code JsonLdFramingErrorCode
	// Message is "JsonLdFramingError.message"
	//
	// Optional, defaults to null.
	Message js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a JsonLdFramingError with all fields set.
func (p JsonLdFramingError) FromRef(ref js.Ref) JsonLdFramingError {
	p.UpdateFrom(ref)
	return p
}

// New creates a new JsonLdFramingError in the application heap.
func (p JsonLdFramingError) New() js.Ref {
	return bindings.JsonLdFramingErrorJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p JsonLdFramingError) UpdateFrom(ref js.Ref) {
	bindings.JsonLdFramingErrorJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p JsonLdFramingError) Update(ref js.Ref) {
	bindings.JsonLdFramingErrorJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type JsonLdRecord = js.Record[js.Any]

type RemoteDocument struct {
	ref js.Ref
}

func (this RemoteDocument) Once() RemoteDocument {
	this.Ref().Once()
	return this
}

func (this RemoteDocument) Ref() js.Ref {
	return this.ref
}

func (this RemoteDocument) FromRef(ref js.Ref) RemoteDocument {
	this.ref = ref
	return this
}

func (this RemoteDocument) Free() {
	this.Ref().Free()
}

// ContentType returns the value of property "RemoteDocument.contentType".
//
// It returns ok=false if there is no such property.
func (this RemoteDocument) ContentType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRemoteDocumentContentType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContextUrl returns the value of property "RemoteDocument.contextUrl".
//
// It returns ok=false if there is no such property.
func (this RemoteDocument) ContextUrl() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRemoteDocumentContextUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Document returns the value of property "RemoteDocument.document".
//
// It returns ok=false if there is no such property.
func (this RemoteDocument) Document() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetRemoteDocumentDocument(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDocument sets the value of property "RemoteDocument.document" to val.
//
// It returns false if the property cannot be set.
func (this RemoteDocument) SetDocument(val js.Any) bool {
	return js.True == bindings.SetRemoteDocumentDocument(
		this.Ref(),
		val.Ref(),
	)
}

// DocumentUrl returns the value of property "RemoteDocument.documentUrl".
//
// It returns ok=false if there is no such property.
func (this RemoteDocument) DocumentUrl() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRemoteDocumentDocumentUrl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Profile returns the value of property "RemoteDocument.profile".
//
// It returns ok=false if there is no such property.
func (this RemoteDocument) Profile() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRemoteDocumentProfile(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type OneOf_RecordAny_ArrayJsonLdRecord_String_RemoteDocument struct {
	ref js.Ref
}

func (x OneOf_RecordAny_ArrayJsonLdRecord_String_RemoteDocument) Ref() js.Ref {
	return x.ref
}

func (x OneOf_RecordAny_ArrayJsonLdRecord_String_RemoteDocument) Free() {
	x.ref.Free()
}

func (x OneOf_RecordAny_ArrayJsonLdRecord_String_RemoteDocument) FromRef(ref js.Ref) OneOf_RecordAny_ArrayJsonLdRecord_String_RemoteDocument {
	return OneOf_RecordAny_ArrayJsonLdRecord_String_RemoteDocument{
		ref: ref,
	}
}

func (x OneOf_RecordAny_ArrayJsonLdRecord_String_RemoteDocument) RecordAny() js.Record[js.Any] {
	return js.Record[js.Any]{}.FromRef(x.ref)
}

func (x OneOf_RecordAny_ArrayJsonLdRecord_String_RemoteDocument) ArrayJsonLdRecord() js.Array[JsonLdRecord] {
	return js.Array[JsonLdRecord]{}.FromRef(x.ref)
}

func (x OneOf_RecordAny_ArrayJsonLdRecord_String_RemoteDocument) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_RecordAny_ArrayJsonLdRecord_String_RemoteDocument) RemoteDocument() RemoteDocument {
	return RemoteDocument{}.FromRef(x.ref)
}

type JsonLdInput = OneOf_RecordAny_ArrayJsonLdRecord_String_RemoteDocument

type LoadDocumentCallbackFunc func(this js.Ref, url js.String, options LoadDocumentOptions) js.Ref

func (fn LoadDocumentCallbackFunc) Register() js.Func[func(url js.String, options LoadDocumentOptions) js.Promise[RemoteDocument]] {
	return js.RegisterCallback[func(url js.String, options LoadDocumentOptions) js.Promise[RemoteDocument]](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LoadDocumentCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		LoadDocumentOptions{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LoadDocumentCallback[T any] struct {
	Fn  func(arg T, this js.Ref, url js.String, options LoadDocumentOptions) js.Ref
	Arg T
}

func (cb *LoadDocumentCallback[T]) Register() js.Func[func(url js.String, options LoadDocumentOptions) js.Promise[RemoteDocument]] {
	return js.RegisterCallback[func(url js.String, options LoadDocumentOptions) js.Promise[RemoteDocument]](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LoadDocumentCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
		LoadDocumentOptions{}.FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LoadDocumentOptions struct {
	// ExtractAllScripts is "LoadDocumentOptions.extractAllScripts"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ExtractAllScripts MUST be set to true to make this field effective.
	ExtractAllScripts bool
	// Profile is "LoadDocumentOptions.profile"
	//
	// Optional, defaults to null.
	Profile js.String
	// RequestProfile is "LoadDocumentOptions.requestProfile"
	//
	// Optional, defaults to null.
	RequestProfile OneOf_String_ArrayString

	FFI_USE_ExtractAllScripts bool // for ExtractAllScripts.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a LoadDocumentOptions with all fields set.
func (p LoadDocumentOptions) FromRef(ref js.Ref) LoadDocumentOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new LoadDocumentOptions in the application heap.
func (p LoadDocumentOptions) New() js.Ref {
	return bindings.LoadDocumentOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p LoadDocumentOptions) UpdateFrom(ref js.Ref) {
	bindings.LoadDocumentOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p LoadDocumentOptions) Update(ref js.Ref) {
	bindings.LoadDocumentOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_JsonLdEmbed_Bool struct {
	ref js.Ref
}

func (x OneOf_JsonLdEmbed_Bool) Ref() js.Ref {
	return x.ref
}

func (x OneOf_JsonLdEmbed_Bool) Free() {
	x.ref.Free()
}

func (x OneOf_JsonLdEmbed_Bool) FromRef(ref js.Ref) OneOf_JsonLdEmbed_Bool {
	return OneOf_JsonLdEmbed_Bool{
		ref: ref,
	}
}

func (x OneOf_JsonLdEmbed_Bool) JsonLdEmbed() JsonLdEmbed {
	return JsonLdEmbed(0).FromRef(x.ref)
}

func (x OneOf_JsonLdEmbed_Bool) Bool() bool {
	return x.ref == js.True
}

type JsonLdOptions struct {
	// Base is "JsonLdOptions.base"
	//
	// Optional, defaults to null.
	Base js.String
	// CompactArrays is "JsonLdOptions.compactArrays"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_CompactArrays MUST be set to true to make this field effective.
	CompactArrays bool
	// CompactToRelative is "JsonLdOptions.compactToRelative"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_CompactToRelative MUST be set to true to make this field effective.
	CompactToRelative bool
	// DocumentLoader is "JsonLdOptions.documentLoader"
	//
	// Optional, defaults to null.
	DocumentLoader js.Func[func(url js.String, options LoadDocumentOptions) js.Promise[RemoteDocument]]
	// ExpandContext is "JsonLdOptions.expandContext"
	//
	// Optional, defaults to null.
	ExpandContext OneOf_RecordAny_String
	// ExtractAllScripts is "JsonLdOptions.extractAllScripts"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ExtractAllScripts MUST be set to true to make this field effective.
	ExtractAllScripts bool
	// FrameExpansion is "JsonLdOptions.frameExpansion"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_FrameExpansion MUST be set to true to make this field effective.
	FrameExpansion bool
	// Ordered is "JsonLdOptions.ordered"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Ordered MUST be set to true to make this field effective.
	Ordered bool
	// ProcessingMode is "JsonLdOptions.processingMode"
	//
	// Optional, defaults to "json-ld-1.1".
	ProcessingMode js.String
	// ProduceGeneralizedRdf is "JsonLdOptions.produceGeneralizedRdf"
	//
	// Optional, defaults to true.
	//
	// NOTE: FFI_USE_ProduceGeneralizedRdf MUST be set to true to make this field effective.
	ProduceGeneralizedRdf bool
	// RdfDirection is "JsonLdOptions.rdfDirection"
	//
	// Optional, defaults to null.
	RdfDirection js.String
	// UseNativeTypes is "JsonLdOptions.useNativeTypes"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_UseNativeTypes MUST be set to true to make this field effective.
	UseNativeTypes bool
	// UseRdfType is "JsonLdOptions.useRdfType"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_UseRdfType MUST be set to true to make this field effective.
	UseRdfType bool
	// Embed is "JsonLdOptions.embed"
	//
	// Optional, defaults to "@once".
	Embed OneOf_JsonLdEmbed_Bool
	// Explicit is "JsonLdOptions.explicit"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Explicit MUST be set to true to make this field effective.
	Explicit bool
	// OmitDefault is "JsonLdOptions.omitDefault"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_OmitDefault MUST be set to true to make this field effective.
	OmitDefault bool
	// OmitGraph is "JsonLdOptions.omitGraph"
	//
	// Optional
	//
	// NOTE: FFI_USE_OmitGraph MUST be set to true to make this field effective.
	OmitGraph bool
	// RequireAll is "JsonLdOptions.requireAll"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_RequireAll MUST be set to true to make this field effective.
	RequireAll bool
	// FrameDefault is "JsonLdOptions.frameDefault"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_FrameDefault MUST be set to true to make this field effective.
	FrameDefault bool

	FFI_USE_CompactArrays         bool // for CompactArrays.
	FFI_USE_CompactToRelative     bool // for CompactToRelative.
	FFI_USE_ExtractAllScripts     bool // for ExtractAllScripts.
	FFI_USE_FrameExpansion        bool // for FrameExpansion.
	FFI_USE_Ordered               bool // for Ordered.
	FFI_USE_ProduceGeneralizedRdf bool // for ProduceGeneralizedRdf.
	FFI_USE_UseNativeTypes        bool // for UseNativeTypes.
	FFI_USE_UseRdfType            bool // for UseRdfType.
	FFI_USE_Explicit              bool // for Explicit.
	FFI_USE_OmitDefault           bool // for OmitDefault.
	FFI_USE_OmitGraph             bool // for OmitGraph.
	FFI_USE_RequireAll            bool // for RequireAll.
	FFI_USE_FrameDefault          bool // for FrameDefault.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a JsonLdOptions with all fields set.
func (p JsonLdOptions) FromRef(ref js.Ref) JsonLdOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new JsonLdOptions in the application heap.
func (p JsonLdOptions) New() js.Ref {
	return bindings.JsonLdOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p JsonLdOptions) UpdateFrom(ref js.Ref) {
	bindings.JsonLdOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p JsonLdOptions) Update(ref js.Ref) {
	bindings.JsonLdOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type RdfLiteral struct {
	ref js.Ref
}

func (this RdfLiteral) Once() RdfLiteral {
	this.Ref().Once()
	return this
}

func (this RdfLiteral) Ref() js.Ref {
	return this.ref
}

func (this RdfLiteral) FromRef(ref js.Ref) RdfLiteral {
	this.ref = ref
	return this
}

func (this RdfLiteral) Free() {
	this.Ref().Free()
}

// Value returns the value of property "RdfLiteral.value".
//
// It returns ok=false if there is no such property.
func (this RdfLiteral) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRdfLiteralValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Datatype returns the value of property "RdfLiteral.datatype".
//
// It returns ok=false if there is no such property.
func (this RdfLiteral) Datatype() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRdfLiteralDatatype(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Language returns the value of property "RdfLiteral.language".
//
// It returns ok=false if there is no such property.
func (this RdfLiteral) Language() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRdfLiteralLanguage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type OneOf_String_RdfLiteral struct {
	ref js.Ref
}

func (x OneOf_String_RdfLiteral) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_RdfLiteral) Free() {
	x.ref.Free()
}

func (x OneOf_String_RdfLiteral) FromRef(ref js.Ref) OneOf_String_RdfLiteral {
	return OneOf_String_RdfLiteral{
		ref: ref,
	}
}

func (x OneOf_String_RdfLiteral) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_RdfLiteral) RdfLiteral() RdfLiteral {
	return RdfLiteral{}.FromRef(x.ref)
}

type RdfTriple struct {
	ref js.Ref
}

func (this RdfTriple) Once() RdfTriple {
	this.Ref().Once()
	return this
}

func (this RdfTriple) Ref() js.Ref {
	return this.ref
}

func (this RdfTriple) FromRef(ref js.Ref) RdfTriple {
	this.ref = ref
	return this
}

func (this RdfTriple) Free() {
	this.Ref().Free()
}

// Subject returns the value of property "RdfTriple.subject".
//
// It returns ok=false if there is no such property.
func (this RdfTriple) Subject() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRdfTripleSubject(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Predicate returns the value of property "RdfTriple.predicate".
//
// It returns ok=false if there is no such property.
func (this RdfTriple) Predicate() (ret js.String, ok bool) {
	ok = js.True == bindings.GetRdfTriplePredicate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Object returns the value of property "RdfTriple.object".
//
// It returns ok=false if there is no such property.
func (this RdfTriple) Object() (ret OneOf_String_RdfLiteral, ok bool) {
	ok = js.True == bindings.GetRdfTripleObject(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type RdfGraph struct {
	ref js.Ref
}

func (this RdfGraph) Once() RdfGraph {
	this.Ref().Once()
	return this
}

func (this RdfGraph) Ref() js.Ref {
	return this.ref
}

func (this RdfGraph) FromRef(ref js.Ref) RdfGraph {
	this.ref = ref
	return this
}

func (this RdfGraph) Free() {
	this.Ref().Free()
}

// HasAdd returns true if the method "RdfGraph.add" exists.
func (this RdfGraph) HasAdd() bool {
	return js.True == bindings.HasRdfGraphAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "RdfGraph.add".
func (this RdfGraph) AddFunc() (fn js.Func[func(triple RdfTriple)]) {
	return fn.FromRef(
		bindings.RdfGraphAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "RdfGraph.add".
func (this RdfGraph) Add(triple RdfTriple) (ret js.Void) {
	bindings.CallRdfGraphAdd(
		this.Ref(), js.Pointer(&ret),
		triple.Ref(),
	)

	return
}

// TryAdd calls the method "RdfGraph.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RdfGraph) TryAdd(triple RdfTriple) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRdfGraphAdd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		triple.Ref(),
	)

	return
}

type RdfDataset struct {
	ref js.Ref
}

func (this RdfDataset) Once() RdfDataset {
	this.Ref().Once()
	return this
}

func (this RdfDataset) Ref() js.Ref {
	return this.ref
}

func (this RdfDataset) FromRef(ref js.Ref) RdfDataset {
	this.ref = ref
	return this
}

func (this RdfDataset) Free() {
	this.Ref().Free()
}

// DefaultGraph returns the value of property "RdfDataset.defaultGraph".
//
// It returns ok=false if there is no such property.
func (this RdfDataset) DefaultGraph() (ret RdfGraph, ok bool) {
	ok = js.True == bindings.GetRdfDatasetDefaultGraph(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasAdd returns true if the method "RdfDataset.add" exists.
func (this RdfDataset) HasAdd() bool {
	return js.True == bindings.HasRdfDatasetAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "RdfDataset.add".
func (this RdfDataset) AddFunc() (fn js.Func[func(graphName js.String, graph RdfGraph)]) {
	return fn.FromRef(
		bindings.RdfDatasetAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "RdfDataset.add".
func (this RdfDataset) Add(graphName js.String, graph RdfGraph) (ret js.Void) {
	bindings.CallRdfDatasetAdd(
		this.Ref(), js.Pointer(&ret),
		graphName.Ref(),
		graph.Ref(),
	)

	return
}

// TryAdd calls the method "RdfDataset.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this RdfDataset) TryAdd(graphName js.String, graph RdfGraph) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRdfDatasetAdd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		graphName.Ref(),
		graph.Ref(),
	)

	return
}

type JsonLdProcessor struct {
	ref js.Ref
}

func (this JsonLdProcessor) Once() JsonLdProcessor {
	this.Ref().Once()
	return this
}

func (this JsonLdProcessor) Ref() js.Ref {
	return this.ref
}

func (this JsonLdProcessor) FromRef(ref js.Ref) JsonLdProcessor {
	this.ref = ref
	return this
}

func (this JsonLdProcessor) Free() {
	this.Ref().Free()
}

// HasCompact returns true if the staticmethod "JsonLdProcessor.compact" exists.
func (this JsonLdProcessor) HasCompact() bool {
	return js.True == bindings.HasJsonLdProcessorCompact(
		this.Ref(),
	)
}

// CompactFunc returns the staticmethod "JsonLdProcessor.compact".
func (this JsonLdProcessor) CompactFunc() (fn js.Func[func(input JsonLdInput, context JsonLdContext, options JsonLdOptions) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorCompactFunc(
			this.Ref(),
		),
	)
}

// Compact calls the staticmethod "JsonLdProcessor.compact".
func (this JsonLdProcessor) Compact(input JsonLdInput, context JsonLdContext, options JsonLdOptions) (ret js.Promise[JsonLdRecord]) {
	bindings.CallJsonLdProcessorCompact(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		context.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCompact calls the staticmethod "JsonLdProcessor.compact"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryCompact(input JsonLdInput, context JsonLdContext, options JsonLdOptions) (ret js.Promise[JsonLdRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorCompact(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		context.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasCompact1 returns true if the staticmethod "JsonLdProcessor.compact" exists.
func (this JsonLdProcessor) HasCompact1() bool {
	return js.True == bindings.HasJsonLdProcessorCompact1(
		this.Ref(),
	)
}

// Compact1Func returns the staticmethod "JsonLdProcessor.compact".
func (this JsonLdProcessor) Compact1Func() (fn js.Func[func(input JsonLdInput, context JsonLdContext) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorCompact1Func(
			this.Ref(),
		),
	)
}

// Compact1 calls the staticmethod "JsonLdProcessor.compact".
func (this JsonLdProcessor) Compact1(input JsonLdInput, context JsonLdContext) (ret js.Promise[JsonLdRecord]) {
	bindings.CallJsonLdProcessorCompact1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		context.Ref(),
	)

	return
}

// TryCompact1 calls the staticmethod "JsonLdProcessor.compact"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryCompact1(input JsonLdInput, context JsonLdContext) (ret js.Promise[JsonLdRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorCompact1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		context.Ref(),
	)

	return
}

// HasCompact2 returns true if the staticmethod "JsonLdProcessor.compact" exists.
func (this JsonLdProcessor) HasCompact2() bool {
	return js.True == bindings.HasJsonLdProcessorCompact2(
		this.Ref(),
	)
}

// Compact2Func returns the staticmethod "JsonLdProcessor.compact".
func (this JsonLdProcessor) Compact2Func() (fn js.Func[func(input JsonLdInput) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorCompact2Func(
			this.Ref(),
		),
	)
}

// Compact2 calls the staticmethod "JsonLdProcessor.compact".
func (this JsonLdProcessor) Compact2(input JsonLdInput) (ret js.Promise[JsonLdRecord]) {
	bindings.CallJsonLdProcessorCompact2(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryCompact2 calls the staticmethod "JsonLdProcessor.compact"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryCompact2(input JsonLdInput) (ret js.Promise[JsonLdRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorCompact2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasExpand returns true if the staticmethod "JsonLdProcessor.expand" exists.
func (this JsonLdProcessor) HasExpand() bool {
	return js.True == bindings.HasJsonLdProcessorExpand(
		this.Ref(),
	)
}

// ExpandFunc returns the staticmethod "JsonLdProcessor.expand".
func (this JsonLdProcessor) ExpandFunc() (fn js.Func[func(input JsonLdInput, options JsonLdOptions) js.Promise[js.Array[JsonLdRecord]]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorExpandFunc(
			this.Ref(),
		),
	)
}

// Expand calls the staticmethod "JsonLdProcessor.expand".
func (this JsonLdProcessor) Expand(input JsonLdInput, options JsonLdOptions) (ret js.Promise[js.Array[JsonLdRecord]]) {
	bindings.CallJsonLdProcessorExpand(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryExpand calls the staticmethod "JsonLdProcessor.expand"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryExpand(input JsonLdInput, options JsonLdOptions) (ret js.Promise[js.Array[JsonLdRecord]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorExpand(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasExpand1 returns true if the staticmethod "JsonLdProcessor.expand" exists.
func (this JsonLdProcessor) HasExpand1() bool {
	return js.True == bindings.HasJsonLdProcessorExpand1(
		this.Ref(),
	)
}

// Expand1Func returns the staticmethod "JsonLdProcessor.expand".
func (this JsonLdProcessor) Expand1Func() (fn js.Func[func(input JsonLdInput) js.Promise[js.Array[JsonLdRecord]]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorExpand1Func(
			this.Ref(),
		),
	)
}

// Expand1 calls the staticmethod "JsonLdProcessor.expand".
func (this JsonLdProcessor) Expand1(input JsonLdInput) (ret js.Promise[js.Array[JsonLdRecord]]) {
	bindings.CallJsonLdProcessorExpand1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryExpand1 calls the staticmethod "JsonLdProcessor.expand"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryExpand1(input JsonLdInput) (ret js.Promise[js.Array[JsonLdRecord]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorExpand1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFlatten returns true if the staticmethod "JsonLdProcessor.flatten" exists.
func (this JsonLdProcessor) HasFlatten() bool {
	return js.True == bindings.HasJsonLdProcessorFlatten(
		this.Ref(),
	)
}

// FlattenFunc returns the staticmethod "JsonLdProcessor.flatten".
func (this JsonLdProcessor) FlattenFunc() (fn js.Func[func(input JsonLdInput, context JsonLdContext, options JsonLdOptions) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFlattenFunc(
			this.Ref(),
		),
	)
}

// Flatten calls the staticmethod "JsonLdProcessor.flatten".
func (this JsonLdProcessor) Flatten(input JsonLdInput, context JsonLdContext, options JsonLdOptions) (ret js.Promise[JsonLdRecord]) {
	bindings.CallJsonLdProcessorFlatten(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		context.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryFlatten calls the staticmethod "JsonLdProcessor.flatten"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryFlatten(input JsonLdInput, context JsonLdContext, options JsonLdOptions) (ret js.Promise[JsonLdRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorFlatten(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		context.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFlatten1 returns true if the staticmethod "JsonLdProcessor.flatten" exists.
func (this JsonLdProcessor) HasFlatten1() bool {
	return js.True == bindings.HasJsonLdProcessorFlatten1(
		this.Ref(),
	)
}

// Flatten1Func returns the staticmethod "JsonLdProcessor.flatten".
func (this JsonLdProcessor) Flatten1Func() (fn js.Func[func(input JsonLdInput, context JsonLdContext) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFlatten1Func(
			this.Ref(),
		),
	)
}

// Flatten1 calls the staticmethod "JsonLdProcessor.flatten".
func (this JsonLdProcessor) Flatten1(input JsonLdInput, context JsonLdContext) (ret js.Promise[JsonLdRecord]) {
	bindings.CallJsonLdProcessorFlatten1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		context.Ref(),
	)

	return
}

// TryFlatten1 calls the staticmethod "JsonLdProcessor.flatten"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryFlatten1(input JsonLdInput, context JsonLdContext) (ret js.Promise[JsonLdRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorFlatten1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		context.Ref(),
	)

	return
}

// HasFlatten2 returns true if the staticmethod "JsonLdProcessor.flatten" exists.
func (this JsonLdProcessor) HasFlatten2() bool {
	return js.True == bindings.HasJsonLdProcessorFlatten2(
		this.Ref(),
	)
}

// Flatten2Func returns the staticmethod "JsonLdProcessor.flatten".
func (this JsonLdProcessor) Flatten2Func() (fn js.Func[func(input JsonLdInput) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFlatten2Func(
			this.Ref(),
		),
	)
}

// Flatten2 calls the staticmethod "JsonLdProcessor.flatten".
func (this JsonLdProcessor) Flatten2(input JsonLdInput) (ret js.Promise[JsonLdRecord]) {
	bindings.CallJsonLdProcessorFlatten2(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryFlatten2 calls the staticmethod "JsonLdProcessor.flatten"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryFlatten2(input JsonLdInput) (ret js.Promise[JsonLdRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorFlatten2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFromRdf returns true if the staticmethod "JsonLdProcessor.fromRdf" exists.
func (this JsonLdProcessor) HasFromRdf() bool {
	return js.True == bindings.HasJsonLdProcessorFromRdf(
		this.Ref(),
	)
}

// FromRdfFunc returns the staticmethod "JsonLdProcessor.fromRdf".
func (this JsonLdProcessor) FromRdfFunc() (fn js.Func[func(input RdfDataset, options JsonLdOptions) js.Promise[js.Array[JsonLdRecord]]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFromRdfFunc(
			this.Ref(),
		),
	)
}

// FromRdf calls the staticmethod "JsonLdProcessor.fromRdf".
func (this JsonLdProcessor) FromRdf(input RdfDataset, options JsonLdOptions) (ret js.Promise[js.Array[JsonLdRecord]]) {
	bindings.CallJsonLdProcessorFromRdf(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryFromRdf calls the staticmethod "JsonLdProcessor.fromRdf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryFromRdf(input RdfDataset, options JsonLdOptions) (ret js.Promise[js.Array[JsonLdRecord]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorFromRdf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFromRdf1 returns true if the staticmethod "JsonLdProcessor.fromRdf" exists.
func (this JsonLdProcessor) HasFromRdf1() bool {
	return js.True == bindings.HasJsonLdProcessorFromRdf1(
		this.Ref(),
	)
}

// FromRdf1Func returns the staticmethod "JsonLdProcessor.fromRdf".
func (this JsonLdProcessor) FromRdf1Func() (fn js.Func[func(input RdfDataset) js.Promise[js.Array[JsonLdRecord]]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFromRdf1Func(
			this.Ref(),
		),
	)
}

// FromRdf1 calls the staticmethod "JsonLdProcessor.fromRdf".
func (this JsonLdProcessor) FromRdf1(input RdfDataset) (ret js.Promise[js.Array[JsonLdRecord]]) {
	bindings.CallJsonLdProcessorFromRdf1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryFromRdf1 calls the staticmethod "JsonLdProcessor.fromRdf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryFromRdf1(input RdfDataset) (ret js.Promise[js.Array[JsonLdRecord]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorFromRdf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasToRdf returns true if the staticmethod "JsonLdProcessor.toRdf" exists.
func (this JsonLdProcessor) HasToRdf() bool {
	return js.True == bindings.HasJsonLdProcessorToRdf(
		this.Ref(),
	)
}

// ToRdfFunc returns the staticmethod "JsonLdProcessor.toRdf".
func (this JsonLdProcessor) ToRdfFunc() (fn js.Func[func(input JsonLdInput, options JsonLdOptions) js.Promise[RdfDataset]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorToRdfFunc(
			this.Ref(),
		),
	)
}

// ToRdf calls the staticmethod "JsonLdProcessor.toRdf".
func (this JsonLdProcessor) ToRdf(input JsonLdInput, options JsonLdOptions) (ret js.Promise[RdfDataset]) {
	bindings.CallJsonLdProcessorToRdf(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryToRdf calls the staticmethod "JsonLdProcessor.toRdf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryToRdf(input JsonLdInput, options JsonLdOptions) (ret js.Promise[RdfDataset], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorToRdf(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasToRdf1 returns true if the staticmethod "JsonLdProcessor.toRdf" exists.
func (this JsonLdProcessor) HasToRdf1() bool {
	return js.True == bindings.HasJsonLdProcessorToRdf1(
		this.Ref(),
	)
}

// ToRdf1Func returns the staticmethod "JsonLdProcessor.toRdf".
func (this JsonLdProcessor) ToRdf1Func() (fn js.Func[func(input JsonLdInput) js.Promise[RdfDataset]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorToRdf1Func(
			this.Ref(),
		),
	)
}

// ToRdf1 calls the staticmethod "JsonLdProcessor.toRdf".
func (this JsonLdProcessor) ToRdf1(input JsonLdInput) (ret js.Promise[RdfDataset]) {
	bindings.CallJsonLdProcessorToRdf1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
	)

	return
}

// TryToRdf1 calls the staticmethod "JsonLdProcessor.toRdf"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryToRdf1(input JsonLdInput) (ret js.Promise[RdfDataset], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorToRdf1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
	)

	return
}

// HasFrame returns true if the staticmethod "JsonLdProcessor.frame" exists.
func (this JsonLdProcessor) HasFrame() bool {
	return js.True == bindings.HasJsonLdProcessorFrame(
		this.Ref(),
	)
}

// FrameFunc returns the staticmethod "JsonLdProcessor.frame".
func (this JsonLdProcessor) FrameFunc() (fn js.Func[func(input JsonLdInput, frame JsonLdInput, options JsonLdOptions) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFrameFunc(
			this.Ref(),
		),
	)
}

// Frame calls the staticmethod "JsonLdProcessor.frame".
func (this JsonLdProcessor) Frame(input JsonLdInput, frame JsonLdInput, options JsonLdOptions) (ret js.Promise[JsonLdRecord]) {
	bindings.CallJsonLdProcessorFrame(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		frame.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryFrame calls the staticmethod "JsonLdProcessor.frame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryFrame(input JsonLdInput, frame JsonLdInput, options JsonLdOptions) (ret js.Promise[JsonLdRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorFrame(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		frame.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFrame1 returns true if the staticmethod "JsonLdProcessor.frame" exists.
func (this JsonLdProcessor) HasFrame1() bool {
	return js.True == bindings.HasJsonLdProcessorFrame1(
		this.Ref(),
	)
}

// Frame1Func returns the staticmethod "JsonLdProcessor.frame".
func (this JsonLdProcessor) Frame1Func() (fn js.Func[func(input JsonLdInput, frame JsonLdInput) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFrame1Func(
			this.Ref(),
		),
	)
}

// Frame1 calls the staticmethod "JsonLdProcessor.frame".
func (this JsonLdProcessor) Frame1(input JsonLdInput, frame JsonLdInput) (ret js.Promise[JsonLdRecord]) {
	bindings.CallJsonLdProcessorFrame1(
		this.Ref(), js.Pointer(&ret),
		input.Ref(),
		frame.Ref(),
	)

	return
}

// TryFrame1 calls the staticmethod "JsonLdProcessor.frame"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this JsonLdProcessor) TryFrame1(input JsonLdInput, frame JsonLdInput) (ret js.Promise[JsonLdRecord], exception js.Any, ok bool) {
	ok = js.True == bindings.TryJsonLdProcessorFrame1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		input.Ref(),
		frame.Ref(),
	)

	return
}

const (
	KHR_parallel_shader_compile_COMPLETION_STATUS_KHR GLenum = 0x91B1
)

type KHR_parallel_shader_compile struct {
	ref js.Ref
}

func (this KHR_parallel_shader_compile) Once() KHR_parallel_shader_compile {
	this.Ref().Once()
	return this
}

func (this KHR_parallel_shader_compile) Ref() js.Ref {
	return this.ref
}

func (this KHR_parallel_shader_compile) FromRef(ref js.Ref) KHR_parallel_shader_compile {
	this.ref = ref
	return this
}

func (this KHR_parallel_shader_compile) Free() {
	this.Ref().Free()
}

const (
	KeyboardEvent_DOM_KEY_LOCATION_STANDARD uint32 = 0x00
	KeyboardEvent_DOM_KEY_LOCATION_LEFT     uint32 = 0x01
	KeyboardEvent_DOM_KEY_LOCATION_RIGHT    uint32 = 0x02
	KeyboardEvent_DOM_KEY_LOCATION_NUMPAD   uint32 = 0x03
)

type KeyboardEventInit struct {
	// Key is "KeyboardEventInit.key"
	//
	// Optional, defaults to "".
	Key js.String
	// Code is "KeyboardEventInit.code"
	//
	// Optional, defaults to "".
	Code js.String
	// Location is "KeyboardEventInit.location"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Location MUST be set to true to make this field effective.
	Location uint32
	// Repeat is "KeyboardEventInit.repeat"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Repeat MUST be set to true to make this field effective.
	Repeat bool
	// IsComposing is "KeyboardEventInit.isComposing"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_IsComposing MUST be set to true to make this field effective.
	IsComposing bool
	// CtrlKey is "KeyboardEventInit.ctrlKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_CtrlKey MUST be set to true to make this field effective.
	CtrlKey bool
	// ShiftKey is "KeyboardEventInit.shiftKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ShiftKey MUST be set to true to make this field effective.
	ShiftKey bool
	// AltKey is "KeyboardEventInit.altKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AltKey MUST be set to true to make this field effective.
	AltKey bool
	// MetaKey is "KeyboardEventInit.metaKey"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_MetaKey MUST be set to true to make this field effective.
	MetaKey bool
	// ModifierAltGraph is "KeyboardEventInit.modifierAltGraph"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierAltGraph MUST be set to true to make this field effective.
	ModifierAltGraph bool
	// ModifierCapsLock is "KeyboardEventInit.modifierCapsLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierCapsLock MUST be set to true to make this field effective.
	ModifierCapsLock bool
	// ModifierFn is "KeyboardEventInit.modifierFn"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierFn MUST be set to true to make this field effective.
	ModifierFn bool
	// ModifierFnLock is "KeyboardEventInit.modifierFnLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierFnLock MUST be set to true to make this field effective.
	ModifierFnLock bool
	// ModifierHyper is "KeyboardEventInit.modifierHyper"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierHyper MUST be set to true to make this field effective.
	ModifierHyper bool
	// ModifierNumLock is "KeyboardEventInit.modifierNumLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierNumLock MUST be set to true to make this field effective.
	ModifierNumLock bool
	// ModifierScrollLock is "KeyboardEventInit.modifierScrollLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierScrollLock MUST be set to true to make this field effective.
	ModifierScrollLock bool
	// ModifierSuper is "KeyboardEventInit.modifierSuper"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSuper MUST be set to true to make this field effective.
	ModifierSuper bool
	// ModifierSymbol is "KeyboardEventInit.modifierSymbol"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSymbol MUST be set to true to make this field effective.
	ModifierSymbol bool
	// ModifierSymbolLock is "KeyboardEventInit.modifierSymbolLock"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_ModifierSymbolLock MUST be set to true to make this field effective.
	ModifierSymbolLock bool
	// View is "KeyboardEventInit.view"
	//
	// Optional, defaults to null.
	View Window
	// Detail is "KeyboardEventInit.detail"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_Detail MUST be set to true to make this field effective.
	Detail int32
	// Bubbles is "KeyboardEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "KeyboardEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "KeyboardEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool
	// CharCode is "KeyboardEventInit.charCode"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_CharCode MUST be set to true to make this field effective.
	CharCode uint32
	// KeyCode is "KeyboardEventInit.keyCode"
	//
	// Optional, defaults to 0.
	//
	// NOTE: FFI_USE_KeyCode MUST be set to true to make this field effective.
	KeyCode uint32

	FFI_USE_Location           bool // for Location.
	FFI_USE_Repeat             bool // for Repeat.
	FFI_USE_IsComposing        bool // for IsComposing.
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
	FFI_USE_CharCode           bool // for CharCode.
	FFI_USE_KeyCode            bool // for KeyCode.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a KeyboardEventInit with all fields set.
func (p KeyboardEventInit) FromRef(ref js.Ref) KeyboardEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new KeyboardEventInit in the application heap.
func (p KeyboardEventInit) New() js.Ref {
	return bindings.KeyboardEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p KeyboardEventInit) UpdateFrom(ref js.Ref) {
	bindings.KeyboardEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p KeyboardEventInit) Update(ref js.Ref) {
	bindings.KeyboardEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewKeyboardEvent(typ js.String, eventInitDict KeyboardEventInit) (ret KeyboardEvent) {
	ret.ref = bindings.NewKeyboardEventByKeyboardEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewKeyboardEventByKeyboardEvent1(typ js.String) (ret KeyboardEvent) {
	ret.ref = bindings.NewKeyboardEventByKeyboardEvent1(
		typ.Ref())
	return
}

type KeyboardEvent struct {
	UIEvent
}

func (this KeyboardEvent) Once() KeyboardEvent {
	this.Ref().Once()
	return this
}

func (this KeyboardEvent) Ref() js.Ref {
	return this.UIEvent.Ref()
}

func (this KeyboardEvent) FromRef(ref js.Ref) KeyboardEvent {
	this.UIEvent = this.UIEvent.FromRef(ref)
	return this
}

func (this KeyboardEvent) Free() {
	this.Ref().Free()
}

// Key returns the value of property "KeyboardEvent.key".
//
// It returns ok=false if there is no such property.
func (this KeyboardEvent) Key() (ret js.String, ok bool) {
	ok = js.True == bindings.GetKeyboardEventKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Code returns the value of property "KeyboardEvent.code".
//
// It returns ok=false if there is no such property.
func (this KeyboardEvent) Code() (ret js.String, ok bool) {
	ok = js.True == bindings.GetKeyboardEventCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Location returns the value of property "KeyboardEvent.location".
//
// It returns ok=false if there is no such property.
func (this KeyboardEvent) Location() (ret uint32, ok bool) {
	ok = js.True == bindings.GetKeyboardEventLocation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CtrlKey returns the value of property "KeyboardEvent.ctrlKey".
//
// It returns ok=false if there is no such property.
func (this KeyboardEvent) CtrlKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetKeyboardEventCtrlKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ShiftKey returns the value of property "KeyboardEvent.shiftKey".
//
// It returns ok=false if there is no such property.
func (this KeyboardEvent) ShiftKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetKeyboardEventShiftKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AltKey returns the value of property "KeyboardEvent.altKey".
//
// It returns ok=false if there is no such property.
func (this KeyboardEvent) AltKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetKeyboardEventAltKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MetaKey returns the value of property "KeyboardEvent.metaKey".
//
// It returns ok=false if there is no such property.
func (this KeyboardEvent) MetaKey() (ret bool, ok bool) {
	ok = js.True == bindings.GetKeyboardEventMetaKey(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Repeat returns the value of property "KeyboardEvent.repeat".
//
// It returns ok=false if there is no such property.
func (this KeyboardEvent) Repeat() (ret bool, ok bool) {
	ok = js.True == bindings.GetKeyboardEventRepeat(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IsComposing returns the value of property "KeyboardEvent.isComposing".
//
// It returns ok=false if there is no such property.
func (this KeyboardEvent) IsComposing() (ret bool, ok bool) {
	ok = js.True == bindings.GetKeyboardEventIsComposing(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CharCode returns the value of property "KeyboardEvent.charCode".
//
// It returns ok=false if there is no such property.
func (this KeyboardEvent) CharCode() (ret uint32, ok bool) {
	ok = js.True == bindings.GetKeyboardEventCharCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// KeyCode returns the value of property "KeyboardEvent.keyCode".
//
// It returns ok=false if there is no such property.
func (this KeyboardEvent) KeyCode() (ret uint32, ok bool) {
	ok = js.True == bindings.GetKeyboardEventKeyCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetModifierState returns true if the method "KeyboardEvent.getModifierState" exists.
func (this KeyboardEvent) HasGetModifierState() bool {
	return js.True == bindings.HasKeyboardEventGetModifierState(
		this.Ref(),
	)
}

// GetModifierStateFunc returns the method "KeyboardEvent.getModifierState".
func (this KeyboardEvent) GetModifierStateFunc() (fn js.Func[func(keyArg js.String) bool]) {
	return fn.FromRef(
		bindings.KeyboardEventGetModifierStateFunc(
			this.Ref(),
		),
	)
}

// GetModifierState calls the method "KeyboardEvent.getModifierState".
func (this KeyboardEvent) GetModifierState(keyArg js.String) (ret bool) {
	bindings.CallKeyboardEventGetModifierState(
		this.Ref(), js.Pointer(&ret),
		keyArg.Ref(),
	)

	return
}

// TryGetModifierState calls the method "KeyboardEvent.getModifierState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyboardEvent) TryGetModifierState(keyArg js.String) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardEventGetModifierState(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		keyArg.Ref(),
	)

	return
}

// HasInitKeyboardEvent returns true if the method "KeyboardEvent.initKeyboardEvent" exists.
func (this KeyboardEvent) HasInitKeyboardEvent() bool {
	return js.True == bindings.HasKeyboardEventInitKeyboardEvent(
		this.Ref(),
	)
}

// InitKeyboardEventFunc returns the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEventFunc() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEventFunc(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent calls the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool) (ret js.Void) {
	bindings.CallKeyboardEventInitKeyboardEvent(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
		js.Bool(bool(ctrlKey)),
		js.Bool(bool(altKey)),
		js.Bool(bool(shiftKey)),
		js.Bool(bool(metaKey)),
	)

	return
}

// TryInitKeyboardEvent calls the method "KeyboardEvent.initKeyboardEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyboardEvent) TryInitKeyboardEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardEventInitKeyboardEvent(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
		js.Bool(bool(ctrlKey)),
		js.Bool(bool(altKey)),
		js.Bool(bool(shiftKey)),
		js.Bool(bool(metaKey)),
	)

	return
}

// HasInitKeyboardEvent1 returns true if the method "KeyboardEvent.initKeyboardEvent" exists.
func (this KeyboardEvent) HasInitKeyboardEvent1() bool {
	return js.True == bindings.HasKeyboardEventInitKeyboardEvent1(
		this.Ref(),
	)
}

// InitKeyboardEvent1Func returns the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent1Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool, shiftKey bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent1Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent1 calls the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool, shiftKey bool) (ret js.Void) {
	bindings.CallKeyboardEventInitKeyboardEvent1(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
		js.Bool(bool(ctrlKey)),
		js.Bool(bool(altKey)),
		js.Bool(bool(shiftKey)),
	)

	return
}

// TryInitKeyboardEvent1 calls the method "KeyboardEvent.initKeyboardEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyboardEvent) TryInitKeyboardEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool, shiftKey bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardEventInitKeyboardEvent1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
		js.Bool(bool(ctrlKey)),
		js.Bool(bool(altKey)),
		js.Bool(bool(shiftKey)),
	)

	return
}

// HasInitKeyboardEvent2 returns true if the method "KeyboardEvent.initKeyboardEvent" exists.
func (this KeyboardEvent) HasInitKeyboardEvent2() bool {
	return js.True == bindings.HasKeyboardEventInitKeyboardEvent2(
		this.Ref(),
	)
}

// InitKeyboardEvent2Func returns the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent2Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent2Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent2 calls the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool) (ret js.Void) {
	bindings.CallKeyboardEventInitKeyboardEvent2(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
		js.Bool(bool(ctrlKey)),
		js.Bool(bool(altKey)),
	)

	return
}

// TryInitKeyboardEvent2 calls the method "KeyboardEvent.initKeyboardEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyboardEvent) TryInitKeyboardEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardEventInitKeyboardEvent2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
		js.Bool(bool(ctrlKey)),
		js.Bool(bool(altKey)),
	)

	return
}

// HasInitKeyboardEvent3 returns true if the method "KeyboardEvent.initKeyboardEvent" exists.
func (this KeyboardEvent) HasInitKeyboardEvent3() bool {
	return js.True == bindings.HasKeyboardEventInitKeyboardEvent3(
		this.Ref(),
	)
}

// InitKeyboardEvent3Func returns the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent3Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent3Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent3 calls the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool) (ret js.Void) {
	bindings.CallKeyboardEventInitKeyboardEvent3(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
		js.Bool(bool(ctrlKey)),
	)

	return
}

// TryInitKeyboardEvent3 calls the method "KeyboardEvent.initKeyboardEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyboardEvent) TryInitKeyboardEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardEventInitKeyboardEvent3(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
		js.Bool(bool(ctrlKey)),
	)

	return
}

// HasInitKeyboardEvent4 returns true if the method "KeyboardEvent.initKeyboardEvent" exists.
func (this KeyboardEvent) HasInitKeyboardEvent4() bool {
	return js.True == bindings.HasKeyboardEventInitKeyboardEvent4(
		this.Ref(),
	)
}

// InitKeyboardEvent4Func returns the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent4Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent4Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent4 calls the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32) (ret js.Void) {
	bindings.CallKeyboardEventInitKeyboardEvent4(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
	)

	return
}

// TryInitKeyboardEvent4 calls the method "KeyboardEvent.initKeyboardEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyboardEvent) TryInitKeyboardEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardEventInitKeyboardEvent4(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
	)

	return
}

// HasInitKeyboardEvent5 returns true if the method "KeyboardEvent.initKeyboardEvent" exists.
func (this KeyboardEvent) HasInitKeyboardEvent5() bool {
	return js.True == bindings.HasKeyboardEventInitKeyboardEvent5(
		this.Ref(),
	)
}

// InitKeyboardEvent5Func returns the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent5Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent5Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent5 calls the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String) (ret js.Void) {
	bindings.CallKeyboardEventInitKeyboardEvent5(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
	)

	return
}

// TryInitKeyboardEvent5 calls the method "KeyboardEvent.initKeyboardEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyboardEvent) TryInitKeyboardEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardEventInitKeyboardEvent5(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
	)

	return
}

// HasInitKeyboardEvent6 returns true if the method "KeyboardEvent.initKeyboardEvent" exists.
func (this KeyboardEvent) HasInitKeyboardEvent6() bool {
	return js.True == bindings.HasKeyboardEventInitKeyboardEvent6(
		this.Ref(),
	)
}

// InitKeyboardEvent6Func returns the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent6Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent6Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent6 calls the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent6(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window) (ret js.Void) {
	bindings.CallKeyboardEventInitKeyboardEvent6(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// TryInitKeyboardEvent6 calls the method "KeyboardEvent.initKeyboardEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyboardEvent) TryInitKeyboardEvent6(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardEventInitKeyboardEvent6(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	return
}

// HasInitKeyboardEvent7 returns true if the method "KeyboardEvent.initKeyboardEvent" exists.
func (this KeyboardEvent) HasInitKeyboardEvent7() bool {
	return js.True == bindings.HasKeyboardEventInitKeyboardEvent7(
		this.Ref(),
	)
}

// InitKeyboardEvent7Func returns the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent7Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent7Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent7 calls the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent7(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void) {
	bindings.CallKeyboardEventInitKeyboardEvent7(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// TryInitKeyboardEvent7 calls the method "KeyboardEvent.initKeyboardEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyboardEvent) TryInitKeyboardEvent7(typeArg js.String, bubblesArg bool, cancelableArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardEventInitKeyboardEvent7(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	return
}

// HasInitKeyboardEvent8 returns true if the method "KeyboardEvent.initKeyboardEvent" exists.
func (this KeyboardEvent) HasInitKeyboardEvent8() bool {
	return js.True == bindings.HasKeyboardEventInitKeyboardEvent8(
		this.Ref(),
	)
}

// InitKeyboardEvent8Func returns the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent8Func() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent8Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent8 calls the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent8(typeArg js.String, bubblesArg bool) (ret js.Void) {
	bindings.CallKeyboardEventInitKeyboardEvent8(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// TryInitKeyboardEvent8 calls the method "KeyboardEvent.initKeyboardEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyboardEvent) TryInitKeyboardEvent8(typeArg js.String, bubblesArg bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardEventInitKeyboardEvent8(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	return
}

// HasInitKeyboardEvent9 returns true if the method "KeyboardEvent.initKeyboardEvent" exists.
func (this KeyboardEvent) HasInitKeyboardEvent9() bool {
	return js.True == bindings.HasKeyboardEventInitKeyboardEvent9(
		this.Ref(),
	)
}

// InitKeyboardEvent9Func returns the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent9Func() (fn js.Func[func(typeArg js.String)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent9Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent9 calls the method "KeyboardEvent.initKeyboardEvent".
func (this KeyboardEvent) InitKeyboardEvent9(typeArg js.String) (ret js.Void) {
	bindings.CallKeyboardEventInitKeyboardEvent9(
		this.Ref(), js.Pointer(&ret),
		typeArg.Ref(),
	)

	return
}

// TryInitKeyboardEvent9 calls the method "KeyboardEvent.initKeyboardEvent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyboardEvent) TryInitKeyboardEvent9(typeArg js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyboardEventInitKeyboardEvent9(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		typeArg.Ref(),
	)

	return
}

type KeyframeEffectOptions struct {
	// Composite is "KeyframeEffectOptions.composite"
	//
	// Optional, defaults to "replace".
	Composite CompositeOperation
	// PseudoElement is "KeyframeEffectOptions.pseudoElement"
	//
	// Optional, defaults to null.
	PseudoElement js.String
	// Fill is "KeyframeEffectOptions.fill"
	//
	// Optional, defaults to "auto".
	Fill FillMode
	// IterationStart is "KeyframeEffectOptions.iterationStart"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_IterationStart MUST be set to true to make this field effective.
	IterationStart float64
	// Iterations is "KeyframeEffectOptions.iterations"
	//
	// Optional, defaults to 1.0.
	//
	// NOTE: FFI_USE_Iterations MUST be set to true to make this field effective.
	Iterations float64
	// Direction is "KeyframeEffectOptions.direction"
	//
	// Optional, defaults to "normal".
	Direction PlaybackDirection
	// Easing is "KeyframeEffectOptions.easing"
	//
	// Optional, defaults to "linear".
	Easing js.String
	// IterationComposite is "KeyframeEffectOptions.iterationComposite"
	//
	// Optional, defaults to "replace".
	IterationComposite IterationCompositeOperation

	FFI_USE_IterationStart bool // for IterationStart.
	FFI_USE_Iterations     bool // for Iterations.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a KeyframeEffectOptions with all fields set.
func (p KeyframeEffectOptions) FromRef(ref js.Ref) KeyframeEffectOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new KeyframeEffectOptions in the application heap.
func (p KeyframeEffectOptions) New() js.Ref {
	return bindings.KeyframeEffectOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p KeyframeEffectOptions) UpdateFrom(ref js.Ref) {
	bindings.KeyframeEffectOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p KeyframeEffectOptions) Update(ref js.Ref) {
	bindings.KeyframeEffectOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_Float64_KeyframeEffectOptions struct {
	ref js.Ref
}

func (x OneOf_Float64_KeyframeEffectOptions) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Float64_KeyframeEffectOptions) Free() {
	x.ref.Free()
}

func (x OneOf_Float64_KeyframeEffectOptions) FromRef(ref js.Ref) OneOf_Float64_KeyframeEffectOptions {
	return OneOf_Float64_KeyframeEffectOptions{
		ref: ref,
	}
}

func (x OneOf_Float64_KeyframeEffectOptions) Float64() float64 {
	return js.Number[float64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Float64_KeyframeEffectOptions) KeyframeEffectOptions() KeyframeEffectOptions {
	var ret KeyframeEffectOptions
	ret.UpdateFrom(x.ref)
	return ret
}

func NewKeyframeEffect(target Element, keyframes js.Object, options OneOf_Float64_KeyframeEffectOptions) (ret KeyframeEffect) {
	ret.ref = bindings.NewKeyframeEffectByKeyframeEffect(
		target.Ref(),
		keyframes.Ref(),
		options.Ref())
	return
}

func NewKeyframeEffectByKeyframeEffect1(target Element, keyframes js.Object) (ret KeyframeEffect) {
	ret.ref = bindings.NewKeyframeEffectByKeyframeEffect1(
		target.Ref(),
		keyframes.Ref())
	return
}

func NewKeyframeEffectByKeyframeEffect2(source KeyframeEffect) (ret KeyframeEffect) {
	ret.ref = bindings.NewKeyframeEffectByKeyframeEffect2(
		source.Ref())
	return
}

type KeyframeEffect struct {
	AnimationEffect
}

func (this KeyframeEffect) Once() KeyframeEffect {
	this.Ref().Once()
	return this
}

func (this KeyframeEffect) Ref() js.Ref {
	return this.AnimationEffect.Ref()
}

func (this KeyframeEffect) FromRef(ref js.Ref) KeyframeEffect {
	this.AnimationEffect = this.AnimationEffect.FromRef(ref)
	return this
}

func (this KeyframeEffect) Free() {
	this.Ref().Free()
}

// Target returns the value of property "KeyframeEffect.target".
//
// It returns ok=false if there is no such property.
func (this KeyframeEffect) Target() (ret Element, ok bool) {
	ok = js.True == bindings.GetKeyframeEffectTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTarget sets the value of property "KeyframeEffect.target" to val.
//
// It returns false if the property cannot be set.
func (this KeyframeEffect) SetTarget(val Element) bool {
	return js.True == bindings.SetKeyframeEffectTarget(
		this.Ref(),
		val.Ref(),
	)
}

// PseudoElement returns the value of property "KeyframeEffect.pseudoElement".
//
// It returns ok=false if there is no such property.
func (this KeyframeEffect) PseudoElement() (ret js.String, ok bool) {
	ok = js.True == bindings.GetKeyframeEffectPseudoElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPseudoElement sets the value of property "KeyframeEffect.pseudoElement" to val.
//
// It returns false if the property cannot be set.
func (this KeyframeEffect) SetPseudoElement(val js.String) bool {
	return js.True == bindings.SetKeyframeEffectPseudoElement(
		this.Ref(),
		val.Ref(),
	)
}

// Composite returns the value of property "KeyframeEffect.composite".
//
// It returns ok=false if there is no such property.
func (this KeyframeEffect) Composite() (ret CompositeOperation, ok bool) {
	ok = js.True == bindings.GetKeyframeEffectComposite(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetComposite sets the value of property "KeyframeEffect.composite" to val.
//
// It returns false if the property cannot be set.
func (this KeyframeEffect) SetComposite(val CompositeOperation) bool {
	return js.True == bindings.SetKeyframeEffectComposite(
		this.Ref(),
		uint32(val),
	)
}

// IterationComposite returns the value of property "KeyframeEffect.iterationComposite".
//
// It returns ok=false if there is no such property.
func (this KeyframeEffect) IterationComposite() (ret IterationCompositeOperation, ok bool) {
	ok = js.True == bindings.GetKeyframeEffectIterationComposite(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetIterationComposite sets the value of property "KeyframeEffect.iterationComposite" to val.
//
// It returns false if the property cannot be set.
func (this KeyframeEffect) SetIterationComposite(val IterationCompositeOperation) bool {
	return js.True == bindings.SetKeyframeEffectIterationComposite(
		this.Ref(),
		uint32(val),
	)
}

// HasGetKeyframes returns true if the method "KeyframeEffect.getKeyframes" exists.
func (this KeyframeEffect) HasGetKeyframes() bool {
	return js.True == bindings.HasKeyframeEffectGetKeyframes(
		this.Ref(),
	)
}

// GetKeyframesFunc returns the method "KeyframeEffect.getKeyframes".
func (this KeyframeEffect) GetKeyframesFunc() (fn js.Func[func() js.Array[js.Object]]) {
	return fn.FromRef(
		bindings.KeyframeEffectGetKeyframesFunc(
			this.Ref(),
		),
	)
}

// GetKeyframes calls the method "KeyframeEffect.getKeyframes".
func (this KeyframeEffect) GetKeyframes() (ret js.Array[js.Object]) {
	bindings.CallKeyframeEffectGetKeyframes(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetKeyframes calls the method "KeyframeEffect.getKeyframes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyframeEffect) TryGetKeyframes() (ret js.Array[js.Object], exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyframeEffectGetKeyframes(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetKeyframes returns true if the method "KeyframeEffect.setKeyframes" exists.
func (this KeyframeEffect) HasSetKeyframes() bool {
	return js.True == bindings.HasKeyframeEffectSetKeyframes(
		this.Ref(),
	)
}

// SetKeyframesFunc returns the method "KeyframeEffect.setKeyframes".
func (this KeyframeEffect) SetKeyframesFunc() (fn js.Func[func(keyframes js.Object)]) {
	return fn.FromRef(
		bindings.KeyframeEffectSetKeyframesFunc(
			this.Ref(),
		),
	)
}

// SetKeyframes calls the method "KeyframeEffect.setKeyframes".
func (this KeyframeEffect) SetKeyframes(keyframes js.Object) (ret js.Void) {
	bindings.CallKeyframeEffectSetKeyframes(
		this.Ref(), js.Pointer(&ret),
		keyframes.Ref(),
	)

	return
}

// TrySetKeyframes calls the method "KeyframeEffect.setKeyframes"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this KeyframeEffect) TrySetKeyframes(keyframes js.Object) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryKeyframeEffectSetKeyframes(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		keyframes.Ref(),
	)

	return
}

type LargeBlobSupport uint32

const (
	_ LargeBlobSupport = iota

	LargeBlobSupport_REQUIRED
	LargeBlobSupport_PREFERRED
)

func (LargeBlobSupport) FromRef(str js.Ref) LargeBlobSupport {
	return LargeBlobSupport(bindings.ConstOfLargeBlobSupport(str))
}

func (x LargeBlobSupport) String() (string, bool) {
	switch x {
	case LargeBlobSupport_REQUIRED:
		return "required", true
	case LargeBlobSupport_PREFERRED:
		return "preferred", true
	default:
		return "", false
	}
}

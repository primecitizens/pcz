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

func NewInstance(module Module, importObject js.Object) Instance {
	return Instance{}.FromRef(
		bindings.NewInstanceByInstance(
			module.Ref(),
			importObject.Ref()),
	)
}

func NewInstanceByInstance1(module Module) Instance {
	return Instance{}.FromRef(
		bindings.NewInstanceByInstance1(
			module.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this Instance) Exports() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetInstanceExports(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
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

// SetBid calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid".
//
// The returned bool will be false if there is no such method.
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetBid(generateBidOutput GenerateBidOutput) (bool, bool) {
	var _ok bool
	_ret := bindings.CallInterestGroupBiddingScriptRunnerGlobalScopeSetBid(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&generateBidOutput),
	)

	return _ret == js.True, _ok
}

// SetBidFunc returns the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetBidFunc() (fn js.Func[func(generateBidOutput GenerateBidOutput) bool]) {
	return fn.FromRef(
		bindings.InterestGroupBiddingScriptRunnerGlobalScopeSetBidFunc(
			this.Ref(),
		),
	)
}

// SetBid1 calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid".
//
// The returned bool will be false if there is no such method.
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetBid1() (bool, bool) {
	var _ok bool
	_ret := bindings.CallInterestGroupBiddingScriptRunnerGlobalScopeSetBid1(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// SetBid1Func returns the method "InterestGroupBiddingScriptRunnerGlobalScope.setBid".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetBid1Func() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.InterestGroupBiddingScriptRunnerGlobalScopeSetBid1Func(
			this.Ref(),
		),
	)
}

// SetPriority calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setPriority".
//
// The returned bool will be false if there is no such method.
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPriority(priority float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallInterestGroupBiddingScriptRunnerGlobalScopeSetPriority(
		this.Ref(), js.Pointer(&_ok),
		float64(priority),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPriorityFunc returns the method "InterestGroupBiddingScriptRunnerGlobalScope.setPriority".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPriorityFunc() (fn js.Func[func(priority float64)]) {
	return fn.FromRef(
		bindings.InterestGroupBiddingScriptRunnerGlobalScopeSetPriorityFunc(
			this.Ref(),
		),
	)
}

// SetPrioritySignalsOverride calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride".
//
// The returned bool will be false if there is no such method.
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPrioritySignalsOverride(key js.String, priority float64) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
		float64(priority),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPrioritySignalsOverrideFunc returns the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPrioritySignalsOverrideFunc() (fn js.Func[func(key js.String, priority float64)]) {
	return fn.FromRef(
		bindings.InterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverrideFunc(
			this.Ref(),
		),
	)
}

// SetPrioritySignalsOverride1 calls the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride".
//
// The returned bool will be false if there is no such method.
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPrioritySignalsOverride1(key js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallInterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1(
		this.Ref(), js.Pointer(&_ok),
		key.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetPrioritySignalsOverride1Func returns the method "InterestGroupBiddingScriptRunnerGlobalScope.setPrioritySignalsOverride".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this InterestGroupBiddingScriptRunnerGlobalScope) SetPrioritySignalsOverride1Func() (fn js.Func[func(key js.String)]) {
	return fn.FromRef(
		bindings.InterestGroupBiddingScriptRunnerGlobalScopeSetPrioritySignalsOverride1Func(
			this.Ref(),
		),
	)
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

// SendReportTo calls the method "InterestGroupReportingScriptRunnerGlobalScope.sendReportTo".
//
// The returned bool will be false if there is no such method.
func (this InterestGroupReportingScriptRunnerGlobalScope) SendReportTo(url js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallInterestGroupReportingScriptRunnerGlobalScopeSendReportTo(
		this.Ref(), js.Pointer(&_ok),
		url.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SendReportToFunc returns the method "InterestGroupReportingScriptRunnerGlobalScope.sendReportTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this InterestGroupReportingScriptRunnerGlobalScope) SendReportToFunc() (fn js.Func[func(url js.String)]) {
	return fn.FromRef(
		bindings.InterestGroupReportingScriptRunnerGlobalScopeSendReportToFunc(
			this.Ref(),
		),
	)
}

// RegisterAdBeacon calls the method "InterestGroupReportingScriptRunnerGlobalScope.registerAdBeacon".
//
// The returned bool will be false if there is no such method.
func (this InterestGroupReportingScriptRunnerGlobalScope) RegisterAdBeacon(mapping js.Record[js.String]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallInterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeacon(
		this.Ref(), js.Pointer(&_ok),
		mapping.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RegisterAdBeaconFunc returns the method "InterestGroupReportingScriptRunnerGlobalScope.registerAdBeacon".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this InterestGroupReportingScriptRunnerGlobalScope) RegisterAdBeaconFunc() (fn js.Func[func(mapping js.Record[js.String])]) {
	return fn.FromRef(
		bindings.InterestGroupReportingScriptRunnerGlobalScopeRegisterAdBeaconFunc(
			this.Ref(),
		),
	)
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
	RootBounds DOMRectInit
	// BoundingClientRect is "IntersectionObserverEntryInit.boundingClientRect"
	//
	// Required
	BoundingClientRect DOMRectInit
	// IntersectionRect is "IntersectionObserverEntryInit.intersectionRect"
	//
	// Required
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

func NewIntersectionObserverEntry(intersectionObserverEntryInit IntersectionObserverEntryInit) IntersectionObserverEntry {
	return IntersectionObserverEntry{}.FromRef(
		bindings.NewIntersectionObserverEntryByIntersectionObserverEntry(
			js.Pointer(&intersectionObserverEntryInit)),
	)
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
// The returned bool will be false if there is no such property.
func (this IntersectionObserverEntry) Time() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetIntersectionObserverEntryTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// RootBounds returns the value of property "IntersectionObserverEntry.rootBounds".
//
// The returned bool will be false if there is no such property.
func (this IntersectionObserverEntry) RootBounds() (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetIntersectionObserverEntryRootBounds(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

// BoundingClientRect returns the value of property "IntersectionObserverEntry.boundingClientRect".
//
// The returned bool will be false if there is no such property.
func (this IntersectionObserverEntry) BoundingClientRect() (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetIntersectionObserverEntryBoundingClientRect(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

// IntersectionRect returns the value of property "IntersectionObserverEntry.intersectionRect".
//
// The returned bool will be false if there is no such property.
func (this IntersectionObserverEntry) IntersectionRect() (DOMRectReadOnly, bool) {
	var _ok bool
	_ret := bindings.GetIntersectionObserverEntryIntersectionRect(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMRectReadOnly{}.FromRef(_ret), _ok
}

// IsIntersecting returns the value of property "IntersectionObserverEntry.isIntersecting".
//
// The returned bool will be false if there is no such property.
func (this IntersectionObserverEntry) IsIntersecting() (bool, bool) {
	var _ok bool
	_ret := bindings.GetIntersectionObserverEntryIsIntersecting(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IntersectionRatio returns the value of property "IntersectionObserverEntry.intersectionRatio".
//
// The returned bool will be false if there is no such property.
func (this IntersectionObserverEntry) IntersectionRatio() (float64, bool) {
	var _ok bool
	_ret := bindings.GetIntersectionObserverEntryIntersectionRatio(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Target returns the value of property "IntersectionObserverEntry.target".
//
// The returned bool will be false if there is no such property.
func (this IntersectionObserverEntry) Target() (Element, bool) {
	var _ok bool
	_ret := bindings.GetIntersectionObserverEntryTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
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

func NewIntersectionObserver(callback js.Func[func(entries js.Array[IntersectionObserverEntry], observer IntersectionObserver)], options IntersectionObserverInit) IntersectionObserver {
	return IntersectionObserver{}.FromRef(
		bindings.NewIntersectionObserverByIntersectionObserver(
			callback.Ref(),
			js.Pointer(&options)),
	)
}

func NewIntersectionObserverByIntersectionObserver1(callback js.Func[func(entries js.Array[IntersectionObserverEntry], observer IntersectionObserver)]) IntersectionObserver {
	return IntersectionObserver{}.FromRef(
		bindings.NewIntersectionObserverByIntersectionObserver1(
			callback.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this IntersectionObserver) Root() (OneOf_Element_Document, bool) {
	var _ok bool
	_ret := bindings.GetIntersectionObserverRoot(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_Element_Document{}.FromRef(_ret), _ok
}

// RootMargin returns the value of property "IntersectionObserver.rootMargin".
//
// The returned bool will be false if there is no such property.
func (this IntersectionObserver) RootMargin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetIntersectionObserverRootMargin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Thresholds returns the value of property "IntersectionObserver.thresholds".
//
// The returned bool will be false if there is no such property.
func (this IntersectionObserver) Thresholds() (js.FrozenArray[float64], bool) {
	var _ok bool
	_ret := bindings.GetIntersectionObserverThresholds(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[float64]{}.FromRef(_ret), _ok
}

// Observe calls the method "IntersectionObserver.observe".
//
// The returned bool will be false if there is no such method.
func (this IntersectionObserver) Observe(target Element) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIntersectionObserverObserve(
		this.Ref(), js.Pointer(&_ok),
		target.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ObserveFunc returns the method "IntersectionObserver.observe".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IntersectionObserver) ObserveFunc() (fn js.Func[func(target Element)]) {
	return fn.FromRef(
		bindings.IntersectionObserverObserveFunc(
			this.Ref(),
		),
	)
}

// Unobserve calls the method "IntersectionObserver.unobserve".
//
// The returned bool will be false if there is no such method.
func (this IntersectionObserver) Unobserve(target Element) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIntersectionObserverUnobserve(
		this.Ref(), js.Pointer(&_ok),
		target.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UnobserveFunc returns the method "IntersectionObserver.unobserve".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IntersectionObserver) UnobserveFunc() (fn js.Func[func(target Element)]) {
	return fn.FromRef(
		bindings.IntersectionObserverUnobserveFunc(
			this.Ref(),
		),
	)
}

// Disconnect calls the method "IntersectionObserver.disconnect".
//
// The returned bool will be false if there is no such method.
func (this IntersectionObserver) Disconnect() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIntersectionObserverDisconnect(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DisconnectFunc returns the method "IntersectionObserver.disconnect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IntersectionObserver) DisconnectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.IntersectionObserverDisconnectFunc(
			this.Ref(),
		),
	)
}

// TakeRecords calls the method "IntersectionObserver.takeRecords".
//
// The returned bool will be false if there is no such method.
func (this IntersectionObserver) TakeRecords() (js.Array[IntersectionObserverEntry], bool) {
	var _ok bool
	_ret := bindings.CallIntersectionObserverTakeRecords(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[IntersectionObserverEntry]{}.FromRef(_ret), _ok
}

// TakeRecordsFunc returns the method "IntersectionObserver.takeRecords".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IntersectionObserver) TakeRecordsFunc() (fn js.Func[func() js.Array[IntersectionObserverEntry]]) {
	return fn.FromRef(
		bindings.IntersectionObserverTakeRecordsFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this InterventionReportBody) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetInterventionReportBodyId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Message returns the value of property "InterventionReportBody.message".
//
// The returned bool will be false if there is no such property.
func (this InterventionReportBody) Message() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetInterventionReportBodyMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SourceFile returns the value of property "InterventionReportBody.sourceFile".
//
// The returned bool will be false if there is no such property.
func (this InterventionReportBody) SourceFile() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetInterventionReportBodySourceFile(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LineNumber returns the value of property "InterventionReportBody.lineNumber".
//
// The returned bool will be false if there is no such property.
func (this InterventionReportBody) LineNumber() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetInterventionReportBodyLineNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ColumnNumber returns the value of property "InterventionReportBody.columnNumber".
//
// The returned bool will be false if there is no such property.
func (this InterventionReportBody) ColumnNumber() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetInterventionReportBodyColumnNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ToJSON calls the method "InterventionReportBody.toJSON".
//
// The returned bool will be false if there is no such method.
func (this InterventionReportBody) ToJSON() (js.Object, bool) {
	var _ok bool
	_ret := bindings.CallInterventionReportBodyToJSON(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Object{}.FromRef(_ret), _ok
}

// ToJSONFunc returns the method "InterventionReportBody.toJSON".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this InterventionReportBody) ToJSONFunc() (fn js.Func[func() js.Object]) {
	return fn.FromRef(
		bindings.InterventionReportBodyToJSONFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this RemoteDocument) ContentType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRemoteDocumentContentType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ContextUrl returns the value of property "RemoteDocument.contextUrl".
//
// The returned bool will be false if there is no such property.
func (this RemoteDocument) ContextUrl() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRemoteDocumentContextUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Document returns the value of property "RemoteDocument.document".
//
// The returned bool will be false if there is no such property.
func (this RemoteDocument) Document() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetRemoteDocumentDocument(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// Document sets the value of property "RemoteDocument.document" to val.
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
// The returned bool will be false if there is no such property.
func (this RemoteDocument) DocumentUrl() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRemoteDocumentDocumentUrl(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Profile returns the value of property "RemoteDocument.profile".
//
// The returned bool will be false if there is no such property.
func (this RemoteDocument) Profile() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRemoteDocumentProfile(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this RdfLiteral) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRdfLiteralValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Datatype returns the value of property "RdfLiteral.datatype".
//
// The returned bool will be false if there is no such property.
func (this RdfLiteral) Datatype() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRdfLiteralDatatype(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Language returns the value of property "RdfLiteral.language".
//
// The returned bool will be false if there is no such property.
func (this RdfLiteral) Language() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRdfLiteralLanguage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this RdfTriple) Subject() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRdfTripleSubject(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Predicate returns the value of property "RdfTriple.predicate".
//
// The returned bool will be false if there is no such property.
func (this RdfTriple) Predicate() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetRdfTriplePredicate(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Object returns the value of property "RdfTriple.object".
//
// The returned bool will be false if there is no such property.
func (this RdfTriple) Object() (OneOf_String_RdfLiteral, bool) {
	var _ok bool
	_ret := bindings.GetRdfTripleObject(
		this.Ref(), js.Pointer(&_ok),
	)
	return OneOf_String_RdfLiteral{}.FromRef(_ret), _ok
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

// Add calls the method "RdfGraph.add".
//
// The returned bool will be false if there is no such method.
func (this RdfGraph) Add(triple RdfTriple) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRdfGraphAdd(
		this.Ref(), js.Pointer(&_ok),
		triple.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddFunc returns the method "RdfGraph.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RdfGraph) AddFunc() (fn js.Func[func(triple RdfTriple)]) {
	return fn.FromRef(
		bindings.RdfGraphAddFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this RdfDataset) DefaultGraph() (RdfGraph, bool) {
	var _ok bool
	_ret := bindings.GetRdfDatasetDefaultGraph(
		this.Ref(), js.Pointer(&_ok),
	)
	return RdfGraph{}.FromRef(_ret), _ok
}

// Add calls the method "RdfDataset.add".
//
// The returned bool will be false if there is no such method.
func (this RdfDataset) Add(graphName js.String, graph RdfGraph) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallRdfDatasetAdd(
		this.Ref(), js.Pointer(&_ok),
		graphName.Ref(),
		graph.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddFunc returns the method "RdfDataset.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this RdfDataset) AddFunc() (fn js.Func[func(graphName js.String, graph RdfGraph)]) {
	return fn.FromRef(
		bindings.RdfDatasetAddFunc(
			this.Ref(),
		),
	)
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

// Compact calls the staticmethod "JsonLdProcessor.compact".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) Compact(input JsonLdInput, context JsonLdContext, options JsonLdOptions) (js.Promise[JsonLdRecord], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorCompact(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		context.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[JsonLdRecord]{}.FromRef(_ret), _ok
}

// CompactFunc returns the staticmethod "JsonLdProcessor.compact".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) CompactFunc() (fn js.Func[func(input JsonLdInput, context JsonLdContext, options JsonLdOptions) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorCompactFunc(
			this.Ref(),
		),
	)
}

// Compact1 calls the staticmethod "JsonLdProcessor.compact".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) Compact1(input JsonLdInput, context JsonLdContext) (js.Promise[JsonLdRecord], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorCompact1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		context.Ref(),
	)

	return js.Promise[JsonLdRecord]{}.FromRef(_ret), _ok
}

// Compact1Func returns the staticmethod "JsonLdProcessor.compact".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) Compact1Func() (fn js.Func[func(input JsonLdInput, context JsonLdContext) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorCompact1Func(
			this.Ref(),
		),
	)
}

// Compact2 calls the staticmethod "JsonLdProcessor.compact".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) Compact2(input JsonLdInput) (js.Promise[JsonLdRecord], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorCompact2(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return js.Promise[JsonLdRecord]{}.FromRef(_ret), _ok
}

// Compact2Func returns the staticmethod "JsonLdProcessor.compact".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) Compact2Func() (fn js.Func[func(input JsonLdInput) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorCompact2Func(
			this.Ref(),
		),
	)
}

// Expand calls the staticmethod "JsonLdProcessor.expand".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) Expand(input JsonLdInput, options JsonLdOptions) (js.Promise[js.Array[JsonLdRecord]], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorExpand(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Array[JsonLdRecord]]{}.FromRef(_ret), _ok
}

// ExpandFunc returns the staticmethod "JsonLdProcessor.expand".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) ExpandFunc() (fn js.Func[func(input JsonLdInput, options JsonLdOptions) js.Promise[js.Array[JsonLdRecord]]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorExpandFunc(
			this.Ref(),
		),
	)
}

// Expand1 calls the staticmethod "JsonLdProcessor.expand".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) Expand1(input JsonLdInput) (js.Promise[js.Array[JsonLdRecord]], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorExpand1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return js.Promise[js.Array[JsonLdRecord]]{}.FromRef(_ret), _ok
}

// Expand1Func returns the staticmethod "JsonLdProcessor.expand".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) Expand1Func() (fn js.Func[func(input JsonLdInput) js.Promise[js.Array[JsonLdRecord]]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorExpand1Func(
			this.Ref(),
		),
	)
}

// Flatten calls the staticmethod "JsonLdProcessor.flatten".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) Flatten(input JsonLdInput, context JsonLdContext, options JsonLdOptions) (js.Promise[JsonLdRecord], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorFlatten(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		context.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[JsonLdRecord]{}.FromRef(_ret), _ok
}

// FlattenFunc returns the staticmethod "JsonLdProcessor.flatten".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) FlattenFunc() (fn js.Func[func(input JsonLdInput, context JsonLdContext, options JsonLdOptions) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFlattenFunc(
			this.Ref(),
		),
	)
}

// Flatten1 calls the staticmethod "JsonLdProcessor.flatten".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) Flatten1(input JsonLdInput, context JsonLdContext) (js.Promise[JsonLdRecord], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorFlatten1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		context.Ref(),
	)

	return js.Promise[JsonLdRecord]{}.FromRef(_ret), _ok
}

// Flatten1Func returns the staticmethod "JsonLdProcessor.flatten".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) Flatten1Func() (fn js.Func[func(input JsonLdInput, context JsonLdContext) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFlatten1Func(
			this.Ref(),
		),
	)
}

// Flatten2 calls the staticmethod "JsonLdProcessor.flatten".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) Flatten2(input JsonLdInput) (js.Promise[JsonLdRecord], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorFlatten2(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return js.Promise[JsonLdRecord]{}.FromRef(_ret), _ok
}

// Flatten2Func returns the staticmethod "JsonLdProcessor.flatten".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) Flatten2Func() (fn js.Func[func(input JsonLdInput) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFlatten2Func(
			this.Ref(),
		),
	)
}

// FromRdf calls the staticmethod "JsonLdProcessor.fromRdf".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) FromRdf(input RdfDataset, options JsonLdOptions) (js.Promise[js.Array[JsonLdRecord]], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorFromRdf(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[js.Array[JsonLdRecord]]{}.FromRef(_ret), _ok
}

// FromRdfFunc returns the staticmethod "JsonLdProcessor.fromRdf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) FromRdfFunc() (fn js.Func[func(input RdfDataset, options JsonLdOptions) js.Promise[js.Array[JsonLdRecord]]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFromRdfFunc(
			this.Ref(),
		),
	)
}

// FromRdf1 calls the staticmethod "JsonLdProcessor.fromRdf".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) FromRdf1(input RdfDataset) (js.Promise[js.Array[JsonLdRecord]], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorFromRdf1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return js.Promise[js.Array[JsonLdRecord]]{}.FromRef(_ret), _ok
}

// FromRdf1Func returns the staticmethod "JsonLdProcessor.fromRdf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) FromRdf1Func() (fn js.Func[func(input RdfDataset) js.Promise[js.Array[JsonLdRecord]]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFromRdf1Func(
			this.Ref(),
		),
	)
}

// ToRdf calls the staticmethod "JsonLdProcessor.toRdf".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) ToRdf(input JsonLdInput, options JsonLdOptions) (js.Promise[RdfDataset], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorToRdf(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[RdfDataset]{}.FromRef(_ret), _ok
}

// ToRdfFunc returns the staticmethod "JsonLdProcessor.toRdf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) ToRdfFunc() (fn js.Func[func(input JsonLdInput, options JsonLdOptions) js.Promise[RdfDataset]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorToRdfFunc(
			this.Ref(),
		),
	)
}

// ToRdf1 calls the staticmethod "JsonLdProcessor.toRdf".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) ToRdf1(input JsonLdInput) (js.Promise[RdfDataset], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorToRdf1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
	)

	return js.Promise[RdfDataset]{}.FromRef(_ret), _ok
}

// ToRdf1Func returns the staticmethod "JsonLdProcessor.toRdf".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) ToRdf1Func() (fn js.Func[func(input JsonLdInput) js.Promise[RdfDataset]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorToRdf1Func(
			this.Ref(),
		),
	)
}

// Frame calls the staticmethod "JsonLdProcessor.frame".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) Frame(input JsonLdInput, frame JsonLdInput, options JsonLdOptions) (js.Promise[JsonLdRecord], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorFrame(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		frame.Ref(),
		js.Pointer(&options),
	)

	return js.Promise[JsonLdRecord]{}.FromRef(_ret), _ok
}

// FrameFunc returns the staticmethod "JsonLdProcessor.frame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) FrameFunc() (fn js.Func[func(input JsonLdInput, frame JsonLdInput, options JsonLdOptions) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFrameFunc(
			this.Ref(),
		),
	)
}

// Frame1 calls the staticmethod "JsonLdProcessor.frame".
//
// The returned bool will be false if there is no such method.
func (this JsonLdProcessor) Frame1(input JsonLdInput, frame JsonLdInput) (js.Promise[JsonLdRecord], bool) {
	var _ok bool
	_ret := bindings.CallJsonLdProcessorFrame1(
		this.Ref(), js.Pointer(&_ok),
		input.Ref(),
		frame.Ref(),
	)

	return js.Promise[JsonLdRecord]{}.FromRef(_ret), _ok
}

// Frame1Func returns the staticmethod "JsonLdProcessor.frame".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this JsonLdProcessor) Frame1Func() (fn js.Func[func(input JsonLdInput, frame JsonLdInput) js.Promise[JsonLdRecord]]) {
	return fn.FromRef(
		bindings.JsonLdProcessorFrame1Func(
			this.Ref(),
		),
	)
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

func NewKeyboardEvent(typ js.String, eventInitDict KeyboardEventInit) KeyboardEvent {
	return KeyboardEvent{}.FromRef(
		bindings.NewKeyboardEventByKeyboardEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewKeyboardEventByKeyboardEvent1(typ js.String) KeyboardEvent {
	return KeyboardEvent{}.FromRef(
		bindings.NewKeyboardEventByKeyboardEvent1(
			typ.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this KeyboardEvent) Key() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetKeyboardEventKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Code returns the value of property "KeyboardEvent.code".
//
// The returned bool will be false if there is no such property.
func (this KeyboardEvent) Code() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetKeyboardEventCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Location returns the value of property "KeyboardEvent.location".
//
// The returned bool will be false if there is no such property.
func (this KeyboardEvent) Location() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetKeyboardEventLocation(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// CtrlKey returns the value of property "KeyboardEvent.ctrlKey".
//
// The returned bool will be false if there is no such property.
func (this KeyboardEvent) CtrlKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetKeyboardEventCtrlKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ShiftKey returns the value of property "KeyboardEvent.shiftKey".
//
// The returned bool will be false if there is no such property.
func (this KeyboardEvent) ShiftKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetKeyboardEventShiftKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// AltKey returns the value of property "KeyboardEvent.altKey".
//
// The returned bool will be false if there is no such property.
func (this KeyboardEvent) AltKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetKeyboardEventAltKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// MetaKey returns the value of property "KeyboardEvent.metaKey".
//
// The returned bool will be false if there is no such property.
func (this KeyboardEvent) MetaKey() (bool, bool) {
	var _ok bool
	_ret := bindings.GetKeyboardEventMetaKey(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Repeat returns the value of property "KeyboardEvent.repeat".
//
// The returned bool will be false if there is no such property.
func (this KeyboardEvent) Repeat() (bool, bool) {
	var _ok bool
	_ret := bindings.GetKeyboardEventRepeat(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// IsComposing returns the value of property "KeyboardEvent.isComposing".
//
// The returned bool will be false if there is no such property.
func (this KeyboardEvent) IsComposing() (bool, bool) {
	var _ok bool
	_ret := bindings.GetKeyboardEventIsComposing(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// CharCode returns the value of property "KeyboardEvent.charCode".
//
// The returned bool will be false if there is no such property.
func (this KeyboardEvent) CharCode() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetKeyboardEventCharCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// KeyCode returns the value of property "KeyboardEvent.keyCode".
//
// The returned bool will be false if there is no such property.
func (this KeyboardEvent) KeyCode() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetKeyboardEventKeyCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// GetModifierState calls the method "KeyboardEvent.getModifierState".
//
// The returned bool will be false if there is no such method.
func (this KeyboardEvent) GetModifierState(keyArg js.String) (bool, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardEventGetModifierState(
		this.Ref(), js.Pointer(&_ok),
		keyArg.Ref(),
	)

	return _ret == js.True, _ok
}

// GetModifierStateFunc returns the method "KeyboardEvent.getModifierState".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyboardEvent) GetModifierStateFunc() (fn js.Func[func(keyArg js.String) bool]) {
	return fn.FromRef(
		bindings.KeyboardEventGetModifierStateFunc(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent calls the method "KeyboardEvent.initKeyboardEvent".
//
// The returned bool will be false if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardEventInitKeyboardEvent(
		this.Ref(), js.Pointer(&_ok),
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

	_ = _ret
	return js.Void{}, _ok
}

// InitKeyboardEventFunc returns the method "KeyboardEvent.initKeyboardEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyboardEvent) InitKeyboardEventFunc() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEventFunc(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent1 calls the method "KeyboardEvent.initKeyboardEvent".
//
// The returned bool will be false if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent1(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool, shiftKey bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardEventInitKeyboardEvent1(
		this.Ref(), js.Pointer(&_ok),
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

	_ = _ret
	return js.Void{}, _ok
}

// InitKeyboardEvent1Func returns the method "KeyboardEvent.initKeyboardEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent1Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool, shiftKey bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent1Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent2 calls the method "KeyboardEvent.initKeyboardEvent".
//
// The returned bool will be false if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent2(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardEventInitKeyboardEvent2(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
		js.Bool(bool(ctrlKey)),
		js.Bool(bool(altKey)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitKeyboardEvent2Func returns the method "KeyboardEvent.initKeyboardEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent2Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool, altKey bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent2Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent3 calls the method "KeyboardEvent.initKeyboardEvent".
//
// The returned bool will be false if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent3(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardEventInitKeyboardEvent3(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
		js.Bool(bool(ctrlKey)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitKeyboardEvent3Func returns the method "KeyboardEvent.initKeyboardEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent3Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32, ctrlKey bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent3Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent4 calls the method "KeyboardEvent.initKeyboardEvent".
//
// The returned bool will be false if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent4(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardEventInitKeyboardEvent4(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
		uint32(locationArg),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitKeyboardEvent4Func returns the method "KeyboardEvent.initKeyboardEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent4Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String, locationArg uint32)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent4Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent5 calls the method "KeyboardEvent.initKeyboardEvent".
//
// The returned bool will be false if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent5(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardEventInitKeyboardEvent5(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
		keyArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitKeyboardEvent5Func returns the method "KeyboardEvent.initKeyboardEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent5Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window, keyArg js.String)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent5Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent6 calls the method "KeyboardEvent.initKeyboardEvent".
//
// The returned bool will be false if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent6(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardEventInitKeyboardEvent6(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
		viewArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitKeyboardEvent6Func returns the method "KeyboardEvent.initKeyboardEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent6Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool, viewArg Window)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent6Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent7 calls the method "KeyboardEvent.initKeyboardEvent".
//
// The returned bool will be false if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent7(typeArg js.String, bubblesArg bool, cancelableArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardEventInitKeyboardEvent7(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
		js.Bool(bool(cancelableArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitKeyboardEvent7Func returns the method "KeyboardEvent.initKeyboardEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent7Func() (fn js.Func[func(typeArg js.String, bubblesArg bool, cancelableArg bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent7Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent8 calls the method "KeyboardEvent.initKeyboardEvent".
//
// The returned bool will be false if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent8(typeArg js.String, bubblesArg bool) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardEventInitKeyboardEvent8(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
		js.Bool(bool(bubblesArg)),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitKeyboardEvent8Func returns the method "KeyboardEvent.initKeyboardEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent8Func() (fn js.Func[func(typeArg js.String, bubblesArg bool)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent8Func(
			this.Ref(),
		),
	)
}

// InitKeyboardEvent9 calls the method "KeyboardEvent.initKeyboardEvent".
//
// The returned bool will be false if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent9(typeArg js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyboardEventInitKeyboardEvent9(
		this.Ref(), js.Pointer(&_ok),
		typeArg.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// InitKeyboardEvent9Func returns the method "KeyboardEvent.initKeyboardEvent".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyboardEvent) InitKeyboardEvent9Func() (fn js.Func[func(typeArg js.String)]) {
	return fn.FromRef(
		bindings.KeyboardEventInitKeyboardEvent9Func(
			this.Ref(),
		),
	)
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

func NewKeyframeEffect(target Element, keyframes js.Object, options OneOf_Float64_KeyframeEffectOptions) KeyframeEffect {
	return KeyframeEffect{}.FromRef(
		bindings.NewKeyframeEffectByKeyframeEffect(
			target.Ref(),
			keyframes.Ref(),
			options.Ref()),
	)
}

func NewKeyframeEffectByKeyframeEffect1(target Element, keyframes js.Object) KeyframeEffect {
	return KeyframeEffect{}.FromRef(
		bindings.NewKeyframeEffectByKeyframeEffect1(
			target.Ref(),
			keyframes.Ref()),
	)
}

func NewKeyframeEffectByKeyframeEffect2(source KeyframeEffect) KeyframeEffect {
	return KeyframeEffect{}.FromRef(
		bindings.NewKeyframeEffectByKeyframeEffect2(
			source.Ref()),
	)
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
// The returned bool will be false if there is no such property.
func (this KeyframeEffect) Target() (Element, bool) {
	var _ok bool
	_ret := bindings.GetKeyframeEffectTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// Target sets the value of property "KeyframeEffect.target" to val.
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
// The returned bool will be false if there is no such property.
func (this KeyframeEffect) PseudoElement() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetKeyframeEffectPseudoElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// PseudoElement sets the value of property "KeyframeEffect.pseudoElement" to val.
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
// The returned bool will be false if there is no such property.
func (this KeyframeEffect) Composite() (CompositeOperation, bool) {
	var _ok bool
	_ret := bindings.GetKeyframeEffectComposite(
		this.Ref(), js.Pointer(&_ok),
	)
	return CompositeOperation(_ret), _ok
}

// Composite sets the value of property "KeyframeEffect.composite" to val.
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
// The returned bool will be false if there is no such property.
func (this KeyframeEffect) IterationComposite() (IterationCompositeOperation, bool) {
	var _ok bool
	_ret := bindings.GetKeyframeEffectIterationComposite(
		this.Ref(), js.Pointer(&_ok),
	)
	return IterationCompositeOperation(_ret), _ok
}

// IterationComposite sets the value of property "KeyframeEffect.iterationComposite" to val.
//
// It returns false if the property cannot be set.
func (this KeyframeEffect) SetIterationComposite(val IterationCompositeOperation) bool {
	return js.True == bindings.SetKeyframeEffectIterationComposite(
		this.Ref(),
		uint32(val),
	)
}

// GetKeyframes calls the method "KeyframeEffect.getKeyframes".
//
// The returned bool will be false if there is no such method.
func (this KeyframeEffect) GetKeyframes() (js.Array[js.Object], bool) {
	var _ok bool
	_ret := bindings.CallKeyframeEffectGetKeyframes(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[js.Object]{}.FromRef(_ret), _ok
}

// GetKeyframesFunc returns the method "KeyframeEffect.getKeyframes".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyframeEffect) GetKeyframesFunc() (fn js.Func[func() js.Array[js.Object]]) {
	return fn.FromRef(
		bindings.KeyframeEffectGetKeyframesFunc(
			this.Ref(),
		),
	)
}

// SetKeyframes calls the method "KeyframeEffect.setKeyframes".
//
// The returned bool will be false if there is no such method.
func (this KeyframeEffect) SetKeyframes(keyframes js.Object) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallKeyframeEffectSetKeyframes(
		this.Ref(), js.Pointer(&_ok),
		keyframes.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetKeyframesFunc returns the method "KeyframeEffect.setKeyframes".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this KeyframeEffect) SetKeyframesFunc() (fn js.Func[func(keyframes js.Object)]) {
	return fn.FromRef(
		bindings.KeyframeEffectSetKeyframesFunc(
			this.Ref(),
		),
	)
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

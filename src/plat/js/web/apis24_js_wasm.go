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

type NavigationHistoryBehavior uint32

const (
	_ NavigationHistoryBehavior = iota

	NavigationHistoryBehavior_AUTO
	NavigationHistoryBehavior_PUSH
	NavigationHistoryBehavior_REPLACE
)

func (NavigationHistoryBehavior) FromRef(str js.Ref) NavigationHistoryBehavior {
	return NavigationHistoryBehavior(bindings.ConstOfNavigationHistoryBehavior(str))
}

func (x NavigationHistoryBehavior) String() (string, bool) {
	switch x {
	case NavigationHistoryBehavior_AUTO:
		return "auto", true
	case NavigationHistoryBehavior_PUSH:
		return "push", true
	case NavigationHistoryBehavior_REPLACE:
		return "replace", true
	default:
		return "", false
	}
}

type NavigationNavigateOptions struct {
	// State is "NavigationNavigateOptions.state"
	//
	// Optional
	State js.Any
	// History is "NavigationNavigateOptions.history"
	//
	// Optional, defaults to "auto".
	History NavigationHistoryBehavior
	// Info is "NavigationNavigateOptions.info"
	//
	// Optional
	Info js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NavigationNavigateOptions with all fields set.
func (p NavigationNavigateOptions) FromRef(ref js.Ref) NavigationNavigateOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NavigationNavigateOptions in the application heap.
func (p NavigationNavigateOptions) New() js.Ref {
	return bindings.NavigationNavigateOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NavigationNavigateOptions) UpdateFrom(ref js.Ref) {
	bindings.NavigationNavigateOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NavigationNavigateOptions) Update(ref js.Ref) {
	bindings.NavigationNavigateOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type NavigationReloadOptions struct {
	// State is "NavigationReloadOptions.state"
	//
	// Optional
	State js.Any
	// Info is "NavigationReloadOptions.info"
	//
	// Optional
	Info js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NavigationReloadOptions with all fields set.
func (p NavigationReloadOptions) FromRef(ref js.Ref) NavigationReloadOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NavigationReloadOptions in the application heap.
func (p NavigationReloadOptions) New() js.Ref {
	return bindings.NavigationReloadOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NavigationReloadOptions) UpdateFrom(ref js.Ref) {
	bindings.NavigationReloadOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NavigationReloadOptions) Update(ref js.Ref) {
	bindings.NavigationReloadOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type NavigationOptions struct {
	// Info is "NavigationOptions.info"
	//
	// Optional
	Info js.Any

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a NavigationOptions with all fields set.
func (p NavigationOptions) FromRef(ref js.Ref) NavigationOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new NavigationOptions in the application heap.
func (p NavigationOptions) New() js.Ref {
	return bindings.NavigationOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p NavigationOptions) UpdateFrom(ref js.Ref) {
	bindings.NavigationOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p NavigationOptions) Update(ref js.Ref) {
	bindings.NavigationOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type NavigationType uint32

const (
	_ NavigationType = iota

	NavigationType_PUSH
	NavigationType_REPLACE
	NavigationType_RELOAD
	NavigationType_TRAVERSE
)

func (NavigationType) FromRef(str js.Ref) NavigationType {
	return NavigationType(bindings.ConstOfNavigationType(str))
}

func (x NavigationType) String() (string, bool) {
	switch x {
	case NavigationType_PUSH:
		return "push", true
	case NavigationType_REPLACE:
		return "replace", true
	case NavigationType_RELOAD:
		return "reload", true
	case NavigationType_TRAVERSE:
		return "traverse", true
	default:
		return "", false
	}
}

type NavigationTransition struct {
	ref js.Ref
}

func (this NavigationTransition) Once() NavigationTransition {
	this.Ref().Once()
	return this
}

func (this NavigationTransition) Ref() js.Ref {
	return this.ref
}

func (this NavigationTransition) FromRef(ref js.Ref) NavigationTransition {
	this.ref = ref
	return this
}

func (this NavigationTransition) Free() {
	this.Ref().Free()
}

// NavigationType returns the value of property "NavigationTransition.navigationType".
//
// It returns ok=false if there is no such property.
func (this NavigationTransition) NavigationType() (ret NavigationType, ok bool) {
	ok = js.True == bindings.GetNavigationTransitionNavigationType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// From returns the value of property "NavigationTransition.from".
//
// It returns ok=false if there is no such property.
func (this NavigationTransition) From() (ret NavigationHistoryEntry, ok bool) {
	ok = js.True == bindings.GetNavigationTransitionFrom(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Finished returns the value of property "NavigationTransition.finished".
//
// It returns ok=false if there is no such property.
func (this NavigationTransition) Finished() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetNavigationTransitionFinished(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type Navigation struct {
	EventTarget
}

func (this Navigation) Once() Navigation {
	this.Ref().Once()
	return this
}

func (this Navigation) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this Navigation) FromRef(ref js.Ref) Navigation {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this Navigation) Free() {
	this.Ref().Free()
}

// CurrentEntry returns the value of property "Navigation.currentEntry".
//
// It returns ok=false if there is no such property.
func (this Navigation) CurrentEntry() (ret NavigationHistoryEntry, ok bool) {
	ok = js.True == bindings.GetNavigationCurrentEntry(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Transition returns the value of property "Navigation.transition".
//
// It returns ok=false if there is no such property.
func (this Navigation) Transition() (ret NavigationTransition, ok bool) {
	ok = js.True == bindings.GetNavigationTransition(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CanGoBack returns the value of property "Navigation.canGoBack".
//
// It returns ok=false if there is no such property.
func (this Navigation) CanGoBack() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigationCanGoBack(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// CanGoForward returns the value of property "Navigation.canGoForward".
//
// It returns ok=false if there is no such property.
func (this Navigation) CanGoForward() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigationCanGoForward(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasEntries returns true if the method "Navigation.entries" exists.
func (this Navigation) HasEntries() bool {
	return js.True == bindings.HasNavigationEntries(
		this.Ref(),
	)
}

// EntriesFunc returns the method "Navigation.entries".
func (this Navigation) EntriesFunc() (fn js.Func[func() js.Array[NavigationHistoryEntry]]) {
	return fn.FromRef(
		bindings.NavigationEntriesFunc(
			this.Ref(),
		),
	)
}

// Entries calls the method "Navigation.entries".
func (this Navigation) Entries() (ret js.Array[NavigationHistoryEntry]) {
	bindings.CallNavigationEntries(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryEntries calls the method "Navigation.entries"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryEntries() (ret js.Array[NavigationHistoryEntry], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationEntries(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasUpdateCurrentEntry returns true if the method "Navigation.updateCurrentEntry" exists.
func (this Navigation) HasUpdateCurrentEntry() bool {
	return js.True == bindings.HasNavigationUpdateCurrentEntry(
		this.Ref(),
	)
}

// UpdateCurrentEntryFunc returns the method "Navigation.updateCurrentEntry".
func (this Navigation) UpdateCurrentEntryFunc() (fn js.Func[func(options NavigationUpdateCurrentEntryOptions)]) {
	return fn.FromRef(
		bindings.NavigationUpdateCurrentEntryFunc(
			this.Ref(),
		),
	)
}

// UpdateCurrentEntry calls the method "Navigation.updateCurrentEntry".
func (this Navigation) UpdateCurrentEntry(options NavigationUpdateCurrentEntryOptions) (ret js.Void) {
	bindings.CallNavigationUpdateCurrentEntry(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryUpdateCurrentEntry calls the method "Navigation.updateCurrentEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryUpdateCurrentEntry(options NavigationUpdateCurrentEntryOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationUpdateCurrentEntry(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasNavigate returns true if the method "Navigation.navigate" exists.
func (this Navigation) HasNavigate() bool {
	return js.True == bindings.HasNavigationNavigate(
		this.Ref(),
	)
}

// NavigateFunc returns the method "Navigation.navigate".
func (this Navigation) NavigateFunc() (fn js.Func[func(url js.String, options NavigationNavigateOptions) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationNavigateFunc(
			this.Ref(),
		),
	)
}

// Navigate calls the method "Navigation.navigate".
func (this Navigation) Navigate(url js.String, options NavigationNavigateOptions) (ret NavigationResult) {
	bindings.CallNavigationNavigate(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryNavigate calls the method "Navigation.navigate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryNavigate(url js.String, options NavigationNavigateOptions) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationNavigate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasNavigate1 returns true if the method "Navigation.navigate" exists.
func (this Navigation) HasNavigate1() bool {
	return js.True == bindings.HasNavigationNavigate1(
		this.Ref(),
	)
}

// Navigate1Func returns the method "Navigation.navigate".
func (this Navigation) Navigate1Func() (fn js.Func[func(url js.String) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationNavigate1Func(
			this.Ref(),
		),
	)
}

// Navigate1 calls the method "Navigation.navigate".
func (this Navigation) Navigate1(url js.String) (ret NavigationResult) {
	bindings.CallNavigationNavigate1(
		this.Ref(), js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryNavigate1 calls the method "Navigation.navigate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryNavigate1(url js.String) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationNavigate1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasReload returns true if the method "Navigation.reload" exists.
func (this Navigation) HasReload() bool {
	return js.True == bindings.HasNavigationReload(
		this.Ref(),
	)
}

// ReloadFunc returns the method "Navigation.reload".
func (this Navigation) ReloadFunc() (fn js.Func[func(options NavigationReloadOptions) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationReloadFunc(
			this.Ref(),
		),
	)
}

// Reload calls the method "Navigation.reload".
func (this Navigation) Reload(options NavigationReloadOptions) (ret NavigationResult) {
	bindings.CallNavigationReload(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryReload calls the method "Navigation.reload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryReload(options NavigationReloadOptions) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationReload(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasReload1 returns true if the method "Navigation.reload" exists.
func (this Navigation) HasReload1() bool {
	return js.True == bindings.HasNavigationReload1(
		this.Ref(),
	)
}

// Reload1Func returns the method "Navigation.reload".
func (this Navigation) Reload1Func() (fn js.Func[func() NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationReload1Func(
			this.Ref(),
		),
	)
}

// Reload1 calls the method "Navigation.reload".
func (this Navigation) Reload1() (ret NavigationResult) {
	bindings.CallNavigationReload1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReload1 calls the method "Navigation.reload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryReload1() (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationReload1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasTraverseTo returns true if the method "Navigation.traverseTo" exists.
func (this Navigation) HasTraverseTo() bool {
	return js.True == bindings.HasNavigationTraverseTo(
		this.Ref(),
	)
}

// TraverseToFunc returns the method "Navigation.traverseTo".
func (this Navigation) TraverseToFunc() (fn js.Func[func(key js.String, options NavigationOptions) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationTraverseToFunc(
			this.Ref(),
		),
	)
}

// TraverseTo calls the method "Navigation.traverseTo".
func (this Navigation) TraverseTo(key js.String, options NavigationOptions) (ret NavigationResult) {
	bindings.CallNavigationTraverseTo(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryTraverseTo calls the method "Navigation.traverseTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryTraverseTo(key js.String, options NavigationOptions) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationTraverseTo(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasTraverseTo1 returns true if the method "Navigation.traverseTo" exists.
func (this Navigation) HasTraverseTo1() bool {
	return js.True == bindings.HasNavigationTraverseTo1(
		this.Ref(),
	)
}

// TraverseTo1Func returns the method "Navigation.traverseTo".
func (this Navigation) TraverseTo1Func() (fn js.Func[func(key js.String) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationTraverseTo1Func(
			this.Ref(),
		),
	)
}

// TraverseTo1 calls the method "Navigation.traverseTo".
func (this Navigation) TraverseTo1(key js.String) (ret NavigationResult) {
	bindings.CallNavigationTraverseTo1(
		this.Ref(), js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryTraverseTo1 calls the method "Navigation.traverseTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryTraverseTo1(key js.String) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationTraverseTo1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

// HasBack returns true if the method "Navigation.back" exists.
func (this Navigation) HasBack() bool {
	return js.True == bindings.HasNavigationBack(
		this.Ref(),
	)
}

// BackFunc returns the method "Navigation.back".
func (this Navigation) BackFunc() (fn js.Func[func(options NavigationOptions) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationBackFunc(
			this.Ref(),
		),
	)
}

// Back calls the method "Navigation.back".
func (this Navigation) Back(options NavigationOptions) (ret NavigationResult) {
	bindings.CallNavigationBack(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryBack calls the method "Navigation.back"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryBack(options NavigationOptions) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationBack(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasBack1 returns true if the method "Navigation.back" exists.
func (this Navigation) HasBack1() bool {
	return js.True == bindings.HasNavigationBack1(
		this.Ref(),
	)
}

// Back1Func returns the method "Navigation.back".
func (this Navigation) Back1Func() (fn js.Func[func() NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationBack1Func(
			this.Ref(),
		),
	)
}

// Back1 calls the method "Navigation.back".
func (this Navigation) Back1() (ret NavigationResult) {
	bindings.CallNavigationBack1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryBack1 calls the method "Navigation.back"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryBack1() (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationBack1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasForward returns true if the method "Navigation.forward" exists.
func (this Navigation) HasForward() bool {
	return js.True == bindings.HasNavigationForward(
		this.Ref(),
	)
}

// ForwardFunc returns the method "Navigation.forward".
func (this Navigation) ForwardFunc() (fn js.Func[func(options NavigationOptions) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationForwardFunc(
			this.Ref(),
		),
	)
}

// Forward calls the method "Navigation.forward".
func (this Navigation) Forward(options NavigationOptions) (ret NavigationResult) {
	bindings.CallNavigationForward(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryForward calls the method "Navigation.forward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryForward(options NavigationOptions) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationForward(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasForward1 returns true if the method "Navigation.forward" exists.
func (this Navigation) HasForward1() bool {
	return js.True == bindings.HasNavigationForward1(
		this.Ref(),
	)
}

// Forward1Func returns the method "Navigation.forward".
func (this Navigation) Forward1Func() (fn js.Func[func() NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationForward1Func(
			this.Ref(),
		),
	)
}

// Forward1 calls the method "Navigation.forward".
func (this Navigation) Forward1() (ret NavigationResult) {
	bindings.CallNavigationForward1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryForward1 calls the method "Navigation.forward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryForward1() (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationForward1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type CustomElementConstructorFunc func(this js.Ref) js.Ref

func (fn CustomElementConstructorFunc) Register() js.Func[func() HTMLElement] {
	return js.RegisterCallback[func() HTMLElement](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CustomElementConstructorFunc) DispatchCallback(
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

type CustomElementConstructor[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CustomElementConstructor[T]) Register() js.Func[func() HTMLElement] {
	return js.RegisterCallback[func() HTMLElement](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CustomElementConstructor[T]) DispatchCallback(
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

type ElementDefinitionOptions struct {
	// Extends is "ElementDefinitionOptions.extends"
	//
	// Optional
	Extends js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ElementDefinitionOptions with all fields set.
func (p ElementDefinitionOptions) FromRef(ref js.Ref) ElementDefinitionOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ElementDefinitionOptions in the application heap.
func (p ElementDefinitionOptions) New() js.Ref {
	return bindings.ElementDefinitionOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ElementDefinitionOptions) UpdateFrom(ref js.Ref) {
	bindings.ElementDefinitionOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ElementDefinitionOptions) Update(ref js.Ref) {
	bindings.ElementDefinitionOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type OneOf_FuncCustomElementConstructor_undefined struct {
	ref js.Ref
}

func (x OneOf_FuncCustomElementConstructor_undefined) Ref() js.Ref {
	return x.ref
}

func (x OneOf_FuncCustomElementConstructor_undefined) Free() {
	x.ref.Free()
}

func (x OneOf_FuncCustomElementConstructor_undefined) FromRef(ref js.Ref) OneOf_FuncCustomElementConstructor_undefined {
	return OneOf_FuncCustomElementConstructor_undefined{
		ref: ref,
	}
}

func (x OneOf_FuncCustomElementConstructor_undefined) Undefined() bool {
	return x.ref == js.Undefined
}

func (x OneOf_FuncCustomElementConstructor_undefined) FuncCustomElementConstructor() js.Func[func() HTMLElement] {
	return js.Func[func() HTMLElement]{}.FromRef(x.ref)
}

type CustomElementRegistry struct {
	ref js.Ref
}

func (this CustomElementRegistry) Once() CustomElementRegistry {
	this.Ref().Once()
	return this
}

func (this CustomElementRegistry) Ref() js.Ref {
	return this.ref
}

func (this CustomElementRegistry) FromRef(ref js.Ref) CustomElementRegistry {
	this.ref = ref
	return this
}

func (this CustomElementRegistry) Free() {
	this.Ref().Free()
}

// HasDefine returns true if the method "CustomElementRegistry.define" exists.
func (this CustomElementRegistry) HasDefine() bool {
	return js.True == bindings.HasCustomElementRegistryDefine(
		this.Ref(),
	)
}

// DefineFunc returns the method "CustomElementRegistry.define".
func (this CustomElementRegistry) DefineFunc() (fn js.Func[func(name js.String, constructor js.Func[func() HTMLElement], options ElementDefinitionOptions)]) {
	return fn.FromRef(
		bindings.CustomElementRegistryDefineFunc(
			this.Ref(),
		),
	)
}

// Define calls the method "CustomElementRegistry.define".
func (this CustomElementRegistry) Define(name js.String, constructor js.Func[func() HTMLElement], options ElementDefinitionOptions) (ret js.Void) {
	bindings.CallCustomElementRegistryDefine(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		constructor.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryDefine calls the method "CustomElementRegistry.define"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomElementRegistry) TryDefine(name js.String, constructor js.Func[func() HTMLElement], options ElementDefinitionOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomElementRegistryDefine(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		constructor.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasDefine1 returns true if the method "CustomElementRegistry.define" exists.
func (this CustomElementRegistry) HasDefine1() bool {
	return js.True == bindings.HasCustomElementRegistryDefine1(
		this.Ref(),
	)
}

// Define1Func returns the method "CustomElementRegistry.define".
func (this CustomElementRegistry) Define1Func() (fn js.Func[func(name js.String, constructor js.Func[func() HTMLElement])]) {
	return fn.FromRef(
		bindings.CustomElementRegistryDefine1Func(
			this.Ref(),
		),
	)
}

// Define1 calls the method "CustomElementRegistry.define".
func (this CustomElementRegistry) Define1(name js.String, constructor js.Func[func() HTMLElement]) (ret js.Void) {
	bindings.CallCustomElementRegistryDefine1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		constructor.Ref(),
	)

	return
}

// TryDefine1 calls the method "CustomElementRegistry.define"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomElementRegistry) TryDefine1(name js.String, constructor js.Func[func() HTMLElement]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomElementRegistryDefine1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		constructor.Ref(),
	)

	return
}

// HasGet returns true if the method "CustomElementRegistry.get" exists.
func (this CustomElementRegistry) HasGet() bool {
	return js.True == bindings.HasCustomElementRegistryGet(
		this.Ref(),
	)
}

// GetFunc returns the method "CustomElementRegistry.get".
func (this CustomElementRegistry) GetFunc() (fn js.Func[func(name js.String) OneOf_FuncCustomElementConstructor_undefined]) {
	return fn.FromRef(
		bindings.CustomElementRegistryGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "CustomElementRegistry.get".
func (this CustomElementRegistry) Get(name js.String) (ret OneOf_FuncCustomElementConstructor_undefined) {
	bindings.CallCustomElementRegistryGet(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "CustomElementRegistry.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomElementRegistry) TryGet(name js.String) (ret OneOf_FuncCustomElementConstructor_undefined, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomElementRegistryGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasGetName returns true if the method "CustomElementRegistry.getName" exists.
func (this CustomElementRegistry) HasGetName() bool {
	return js.True == bindings.HasCustomElementRegistryGetName(
		this.Ref(),
	)
}

// GetNameFunc returns the method "CustomElementRegistry.getName".
func (this CustomElementRegistry) GetNameFunc() (fn js.Func[func(constructor js.Func[func() HTMLElement]) js.String]) {
	return fn.FromRef(
		bindings.CustomElementRegistryGetNameFunc(
			this.Ref(),
		),
	)
}

// GetName calls the method "CustomElementRegistry.getName".
func (this CustomElementRegistry) GetName(constructor js.Func[func() HTMLElement]) (ret js.String) {
	bindings.CallCustomElementRegistryGetName(
		this.Ref(), js.Pointer(&ret),
		constructor.Ref(),
	)

	return
}

// TryGetName calls the method "CustomElementRegistry.getName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomElementRegistry) TryGetName(constructor js.Func[func() HTMLElement]) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomElementRegistryGetName(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		constructor.Ref(),
	)

	return
}

// HasWhenDefined returns true if the method "CustomElementRegistry.whenDefined" exists.
func (this CustomElementRegistry) HasWhenDefined() bool {
	return js.True == bindings.HasCustomElementRegistryWhenDefined(
		this.Ref(),
	)
}

// WhenDefinedFunc returns the method "CustomElementRegistry.whenDefined".
func (this CustomElementRegistry) WhenDefinedFunc() (fn js.Func[func(name js.String) js.Promise[js.Func[func() HTMLElement]]]) {
	return fn.FromRef(
		bindings.CustomElementRegistryWhenDefinedFunc(
			this.Ref(),
		),
	)
}

// WhenDefined calls the method "CustomElementRegistry.whenDefined".
func (this CustomElementRegistry) WhenDefined(name js.String) (ret js.Promise[js.Func[func() HTMLElement]]) {
	bindings.CallCustomElementRegistryWhenDefined(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryWhenDefined calls the method "CustomElementRegistry.whenDefined"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomElementRegistry) TryWhenDefined(name js.String) (ret js.Promise[js.Func[func() HTMLElement]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomElementRegistryWhenDefined(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasUpgrade returns true if the method "CustomElementRegistry.upgrade" exists.
func (this CustomElementRegistry) HasUpgrade() bool {
	return js.True == bindings.HasCustomElementRegistryUpgrade(
		this.Ref(),
	)
}

// UpgradeFunc returns the method "CustomElementRegistry.upgrade".
func (this CustomElementRegistry) UpgradeFunc() (fn js.Func[func(root Node)]) {
	return fn.FromRef(
		bindings.CustomElementRegistryUpgradeFunc(
			this.Ref(),
		),
	)
}

// Upgrade calls the method "CustomElementRegistry.upgrade".
func (this CustomElementRegistry) Upgrade(root Node) (ret js.Void) {
	bindings.CallCustomElementRegistryUpgrade(
		this.Ref(), js.Pointer(&ret),
		root.Ref(),
	)

	return
}

// TryUpgrade calls the method "CustomElementRegistry.upgrade"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomElementRegistry) TryUpgrade(root Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomElementRegistryUpgrade(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		root.Ref(),
	)

	return
}

type VibratePattern = OneOf_Uint32_ArrayUint32

type GamepadMappingType uint32

const (
	_ GamepadMappingType = iota

	GamepadMappingType_
	GamepadMappingType_STANDARD
	GamepadMappingType_XR_STANDARD
)

func (GamepadMappingType) FromRef(str js.Ref) GamepadMappingType {
	return GamepadMappingType(bindings.ConstOfGamepadMappingType(str))
}

func (x GamepadMappingType) String() (string, bool) {
	switch x {
	case GamepadMappingType_:
		return "", true
	case GamepadMappingType_STANDARD:
		return "standard", true
	case GamepadMappingType_XR_STANDARD:
		return "xr-standard", true
	default:
		return "", false
	}
}

type GamepadButton struct {
	ref js.Ref
}

func (this GamepadButton) Once() GamepadButton {
	this.Ref().Once()
	return this
}

func (this GamepadButton) Ref() js.Ref {
	return this.ref
}

func (this GamepadButton) FromRef(ref js.Ref) GamepadButton {
	this.ref = ref
	return this
}

func (this GamepadButton) Free() {
	this.Ref().Free()
}

// Pressed returns the value of property "GamepadButton.pressed".
//
// It returns ok=false if there is no such property.
func (this GamepadButton) Pressed() (ret bool, ok bool) {
	ok = js.True == bindings.GetGamepadButtonPressed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Touched returns the value of property "GamepadButton.touched".
//
// It returns ok=false if there is no such property.
func (this GamepadButton) Touched() (ret bool, ok bool) {
	ok = js.True == bindings.GetGamepadButtonTouched(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "GamepadButton.value".
//
// It returns ok=false if there is no such property.
func (this GamepadButton) Value() (ret float64, ok bool) {
	ok = js.True == bindings.GetGamepadButtonValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type GamepadHand uint32

const (
	_ GamepadHand = iota

	GamepadHand_
	GamepadHand_LEFT
	GamepadHand_RIGHT
)

func (GamepadHand) FromRef(str js.Ref) GamepadHand {
	return GamepadHand(bindings.ConstOfGamepadHand(str))
}

func (x GamepadHand) String() (string, bool) {
	switch x {
	case GamepadHand_:
		return "", true
	case GamepadHand_LEFT:
		return "left", true
	case GamepadHand_RIGHT:
		return "right", true
	default:
		return "", false
	}
}

type GamepadHapticEffectType uint32

const (
	_ GamepadHapticEffectType = iota

	GamepadHapticEffectType_DUAL_RUMBLE
)

func (GamepadHapticEffectType) FromRef(str js.Ref) GamepadHapticEffectType {
	return GamepadHapticEffectType(bindings.ConstOfGamepadHapticEffectType(str))
}

func (x GamepadHapticEffectType) String() (string, bool) {
	switch x {
	case GamepadHapticEffectType_DUAL_RUMBLE:
		return "dual-rumble", true
	default:
		return "", false
	}
}

type GamepadHapticsResult uint32

const (
	_ GamepadHapticsResult = iota

	GamepadHapticsResult_COMPLETE
	GamepadHapticsResult_PREEMPTED
)

func (GamepadHapticsResult) FromRef(str js.Ref) GamepadHapticsResult {
	return GamepadHapticsResult(bindings.ConstOfGamepadHapticsResult(str))
}

func (x GamepadHapticsResult) String() (string, bool) {
	switch x {
	case GamepadHapticsResult_COMPLETE:
		return "complete", true
	case GamepadHapticsResult_PREEMPTED:
		return "preempted", true
	default:
		return "", false
	}
}

type GamepadEffectParameters struct {
	// Duration is "GamepadEffectParameters.duration"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_Duration MUST be set to true to make this field effective.
	Duration float64
	// StartDelay is "GamepadEffectParameters.startDelay"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_StartDelay MUST be set to true to make this field effective.
	StartDelay float64
	// StrongMagnitude is "GamepadEffectParameters.strongMagnitude"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_StrongMagnitude MUST be set to true to make this field effective.
	StrongMagnitude float64
	// WeakMagnitude is "GamepadEffectParameters.weakMagnitude"
	//
	// Optional, defaults to 0.0.
	//
	// NOTE: FFI_USE_WeakMagnitude MUST be set to true to make this field effective.
	WeakMagnitude float64

	FFI_USE_Duration        bool // for Duration.
	FFI_USE_StartDelay      bool // for StartDelay.
	FFI_USE_StrongMagnitude bool // for StrongMagnitude.
	FFI_USE_WeakMagnitude   bool // for WeakMagnitude.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a GamepadEffectParameters with all fields set.
func (p GamepadEffectParameters) FromRef(ref js.Ref) GamepadEffectParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new GamepadEffectParameters in the application heap.
func (p GamepadEffectParameters) New() js.Ref {
	return bindings.GamepadEffectParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p GamepadEffectParameters) UpdateFrom(ref js.Ref) {
	bindings.GamepadEffectParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p GamepadEffectParameters) Update(ref js.Ref) {
	bindings.GamepadEffectParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type GamepadHapticActuatorType uint32

const (
	_ GamepadHapticActuatorType = iota

	GamepadHapticActuatorType_VIBRATION
	GamepadHapticActuatorType_DUAL_RUMBLE
)

func (GamepadHapticActuatorType) FromRef(str js.Ref) GamepadHapticActuatorType {
	return GamepadHapticActuatorType(bindings.ConstOfGamepadHapticActuatorType(str))
}

func (x GamepadHapticActuatorType) String() (string, bool) {
	switch x {
	case GamepadHapticActuatorType_VIBRATION:
		return "vibration", true
	case GamepadHapticActuatorType_DUAL_RUMBLE:
		return "dual-rumble", true
	default:
		return "", false
	}
}

type GamepadHapticActuator struct {
	ref js.Ref
}

func (this GamepadHapticActuator) Once() GamepadHapticActuator {
	this.Ref().Once()
	return this
}

func (this GamepadHapticActuator) Ref() js.Ref {
	return this.ref
}

func (this GamepadHapticActuator) FromRef(ref js.Ref) GamepadHapticActuator {
	this.ref = ref
	return this
}

func (this GamepadHapticActuator) Free() {
	this.Ref().Free()
}

// Type returns the value of property "GamepadHapticActuator.type".
//
// It returns ok=false if there is no such property.
func (this GamepadHapticActuator) Type() (ret GamepadHapticActuatorType, ok bool) {
	ok = js.True == bindings.GetGamepadHapticActuatorType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasCanPlayEffectType returns true if the method "GamepadHapticActuator.canPlayEffectType" exists.
func (this GamepadHapticActuator) HasCanPlayEffectType() bool {
	return js.True == bindings.HasGamepadHapticActuatorCanPlayEffectType(
		this.Ref(),
	)
}

// CanPlayEffectTypeFunc returns the method "GamepadHapticActuator.canPlayEffectType".
func (this GamepadHapticActuator) CanPlayEffectTypeFunc() (fn js.Func[func(typ GamepadHapticEffectType) bool]) {
	return fn.FromRef(
		bindings.GamepadHapticActuatorCanPlayEffectTypeFunc(
			this.Ref(),
		),
	)
}

// CanPlayEffectType calls the method "GamepadHapticActuator.canPlayEffectType".
func (this GamepadHapticActuator) CanPlayEffectType(typ GamepadHapticEffectType) (ret bool) {
	bindings.CallGamepadHapticActuatorCanPlayEffectType(
		this.Ref(), js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryCanPlayEffectType calls the method "GamepadHapticActuator.canPlayEffectType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GamepadHapticActuator) TryCanPlayEffectType(typ GamepadHapticEffectType) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGamepadHapticActuatorCanPlayEffectType(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasPlayEffect returns true if the method "GamepadHapticActuator.playEffect" exists.
func (this GamepadHapticActuator) HasPlayEffect() bool {
	return js.True == bindings.HasGamepadHapticActuatorPlayEffect(
		this.Ref(),
	)
}

// PlayEffectFunc returns the method "GamepadHapticActuator.playEffect".
func (this GamepadHapticActuator) PlayEffectFunc() (fn js.Func[func(typ GamepadHapticEffectType, params GamepadEffectParameters) js.Promise[GamepadHapticsResult]]) {
	return fn.FromRef(
		bindings.GamepadHapticActuatorPlayEffectFunc(
			this.Ref(),
		),
	)
}

// PlayEffect calls the method "GamepadHapticActuator.playEffect".
func (this GamepadHapticActuator) PlayEffect(typ GamepadHapticEffectType, params GamepadEffectParameters) (ret js.Promise[GamepadHapticsResult]) {
	bindings.CallGamepadHapticActuatorPlayEffect(
		this.Ref(), js.Pointer(&ret),
		uint32(typ),
		js.Pointer(&params),
	)

	return
}

// TryPlayEffect calls the method "GamepadHapticActuator.playEffect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GamepadHapticActuator) TryPlayEffect(typ GamepadHapticEffectType, params GamepadEffectParameters) (ret js.Promise[GamepadHapticsResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGamepadHapticActuatorPlayEffect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
		js.Pointer(&params),
	)

	return
}

// HasPlayEffect1 returns true if the method "GamepadHapticActuator.playEffect" exists.
func (this GamepadHapticActuator) HasPlayEffect1() bool {
	return js.True == bindings.HasGamepadHapticActuatorPlayEffect1(
		this.Ref(),
	)
}

// PlayEffect1Func returns the method "GamepadHapticActuator.playEffect".
func (this GamepadHapticActuator) PlayEffect1Func() (fn js.Func[func(typ GamepadHapticEffectType) js.Promise[GamepadHapticsResult]]) {
	return fn.FromRef(
		bindings.GamepadHapticActuatorPlayEffect1Func(
			this.Ref(),
		),
	)
}

// PlayEffect1 calls the method "GamepadHapticActuator.playEffect".
func (this GamepadHapticActuator) PlayEffect1(typ GamepadHapticEffectType) (ret js.Promise[GamepadHapticsResult]) {
	bindings.CallGamepadHapticActuatorPlayEffect1(
		this.Ref(), js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryPlayEffect1 calls the method "GamepadHapticActuator.playEffect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GamepadHapticActuator) TryPlayEffect1(typ GamepadHapticEffectType) (ret js.Promise[GamepadHapticsResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGamepadHapticActuatorPlayEffect1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasPulse returns true if the method "GamepadHapticActuator.pulse" exists.
func (this GamepadHapticActuator) HasPulse() bool {
	return js.True == bindings.HasGamepadHapticActuatorPulse(
		this.Ref(),
	)
}

// PulseFunc returns the method "GamepadHapticActuator.pulse".
func (this GamepadHapticActuator) PulseFunc() (fn js.Func[func(value float64, duration float64) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.GamepadHapticActuatorPulseFunc(
			this.Ref(),
		),
	)
}

// Pulse calls the method "GamepadHapticActuator.pulse".
func (this GamepadHapticActuator) Pulse(value float64, duration float64) (ret js.Promise[js.Boolean]) {
	bindings.CallGamepadHapticActuatorPulse(
		this.Ref(), js.Pointer(&ret),
		float64(value),
		float64(duration),
	)

	return
}

// TryPulse calls the method "GamepadHapticActuator.pulse"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GamepadHapticActuator) TryPulse(value float64, duration float64) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGamepadHapticActuatorPulse(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
		float64(duration),
	)

	return
}

// HasReset returns true if the method "GamepadHapticActuator.reset" exists.
func (this GamepadHapticActuator) HasReset() bool {
	return js.True == bindings.HasGamepadHapticActuatorReset(
		this.Ref(),
	)
}

// ResetFunc returns the method "GamepadHapticActuator.reset".
func (this GamepadHapticActuator) ResetFunc() (fn js.Func[func() js.Promise[GamepadHapticsResult]]) {
	return fn.FromRef(
		bindings.GamepadHapticActuatorResetFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "GamepadHapticActuator.reset".
func (this GamepadHapticActuator) Reset() (ret js.Promise[GamepadHapticsResult]) {
	bindings.CallGamepadHapticActuatorReset(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "GamepadHapticActuator.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GamepadHapticActuator) TryReset() (ret js.Promise[GamepadHapticsResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGamepadHapticActuatorReset(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type GamepadPose struct {
	ref js.Ref
}

func (this GamepadPose) Once() GamepadPose {
	this.Ref().Once()
	return this
}

func (this GamepadPose) Ref() js.Ref {
	return this.ref
}

func (this GamepadPose) FromRef(ref js.Ref) GamepadPose {
	this.ref = ref
	return this
}

func (this GamepadPose) Free() {
	this.Ref().Free()
}

// HasOrientation returns the value of property "GamepadPose.hasOrientation".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) HasOrientation() (ret bool, ok bool) {
	ok = js.True == bindings.GetGamepadPoseHasOrientation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasPosition returns the value of property "GamepadPose.hasPosition".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) HasPosition() (ret bool, ok bool) {
	ok = js.True == bindings.GetGamepadPoseHasPosition(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Position returns the value of property "GamepadPose.position".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) Position() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPosePosition(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LinearVelocity returns the value of property "GamepadPose.linearVelocity".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) LinearVelocity() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPoseLinearVelocity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// LinearAcceleration returns the value of property "GamepadPose.linearAcceleration".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) LinearAcceleration() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPoseLinearAcceleration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Orientation returns the value of property "GamepadPose.orientation".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) Orientation() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPoseOrientation(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AngularVelocity returns the value of property "GamepadPose.angularVelocity".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) AngularVelocity() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPoseAngularVelocity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AngularAcceleration returns the value of property "GamepadPose.angularAcceleration".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) AngularAcceleration() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPoseAngularAcceleration(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type GamepadTouch struct {
	ref js.Ref
}

func (this GamepadTouch) Once() GamepadTouch {
	this.Ref().Once()
	return this
}

func (this GamepadTouch) Ref() js.Ref {
	return this.ref
}

func (this GamepadTouch) FromRef(ref js.Ref) GamepadTouch {
	this.ref = ref
	return this
}

func (this GamepadTouch) Free() {
	this.Ref().Free()
}

// TouchId returns the value of property "GamepadTouch.touchId".
//
// It returns ok=false if there is no such property.
func (this GamepadTouch) TouchId() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGamepadTouchTouchId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SurfaceId returns the value of property "GamepadTouch.surfaceId".
//
// It returns ok=false if there is no such property.
func (this GamepadTouch) SurfaceId() (ret uint8, ok bool) {
	ok = js.True == bindings.GetGamepadTouchSurfaceId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Position returns the value of property "GamepadTouch.position".
//
// It returns ok=false if there is no such property.
func (this GamepadTouch) Position() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadTouchPosition(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SurfaceDimensions returns the value of property "GamepadTouch.surfaceDimensions".
//
// It returns ok=false if there is no such property.
func (this GamepadTouch) SurfaceDimensions() (ret js.TypedArray[uint32], ok bool) {
	ok = js.True == bindings.GetGamepadTouchSurfaceDimensions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type Gamepad struct {
	ref js.Ref
}

func (this Gamepad) Once() Gamepad {
	this.Ref().Once()
	return this
}

func (this Gamepad) Ref() js.Ref {
	return this.ref
}

func (this Gamepad) FromRef(ref js.Ref) Gamepad {
	this.ref = ref
	return this
}

func (this Gamepad) Free() {
	this.Ref().Free()
}

// Id returns the value of property "Gamepad.id".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGamepadId(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Index returns the value of property "Gamepad.index".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Index() (ret int32, ok bool) {
	ok = js.True == bindings.GetGamepadIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Connected returns the value of property "Gamepad.connected".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Connected() (ret bool, ok bool) {
	ok = js.True == bindings.GetGamepadConnected(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Timestamp returns the value of property "Gamepad.timestamp".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Timestamp() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetGamepadTimestamp(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Mapping returns the value of property "Gamepad.mapping".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Mapping() (ret GamepadMappingType, ok bool) {
	ok = js.True == bindings.GetGamepadMapping(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Axes returns the value of property "Gamepad.axes".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Axes() (ret js.FrozenArray[float64], ok bool) {
	ok = js.True == bindings.GetGamepadAxes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Buttons returns the value of property "Gamepad.buttons".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Buttons() (ret js.FrozenArray[GamepadButton], ok bool) {
	ok = js.True == bindings.GetGamepadButtons(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Hand returns the value of property "Gamepad.hand".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Hand() (ret GamepadHand, ok bool) {
	ok = js.True == bindings.GetGamepadHand(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HapticActuators returns the value of property "Gamepad.hapticActuators".
//
// It returns ok=false if there is no such property.
func (this Gamepad) HapticActuators() (ret js.FrozenArray[GamepadHapticActuator], ok bool) {
	ok = js.True == bindings.GetGamepadHapticActuators(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Pose returns the value of property "Gamepad.pose".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Pose() (ret GamepadPose, ok bool) {
	ok = js.True == bindings.GetGamepadPose(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// TouchEvents returns the value of property "Gamepad.touchEvents".
//
// It returns ok=false if there is no such property.
func (this Gamepad) TouchEvents() (ret js.FrozenArray[GamepadTouch], ok bool) {
	ok = js.True == bindings.GetGamepadTouchEvents(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// VibrationActuator returns the value of property "Gamepad.vibrationActuator".
//
// It returns ok=false if there is no such property.
func (this Gamepad) VibrationActuator() (ret GamepadHapticActuator, ok bool) {
	ok = js.True == bindings.GetGamepadVibrationActuator(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type RelatedApplication struct {
	// Platform is "RelatedApplication.platform"
	//
	// Required
	Platform js.String
	// Url is "RelatedApplication.url"
	//
	// Optional
	Url js.String
	// Id is "RelatedApplication.id"
	//
	// Optional
	Id js.String
	// Version is "RelatedApplication.version"
	//
	// Optional
	Version js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RelatedApplication with all fields set.
func (p RelatedApplication) FromRef(ref js.Ref) RelatedApplication {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RelatedApplication in the application heap.
func (p RelatedApplication) New() js.Ref {
	return bindings.RelatedApplicationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p RelatedApplication) UpdateFrom(ref js.Ref) {
	bindings.RelatedApplicationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p RelatedApplication) Update(ref js.Ref) {
	bindings.RelatedApplicationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaKeySystemMediaCapability struct {
	// ContentType is "MediaKeySystemMediaCapability.contentType"
	//
	// Optional, defaults to "".
	ContentType js.String
	// EncryptionScheme is "MediaKeySystemMediaCapability.encryptionScheme"
	//
	// Optional, defaults to null.
	EncryptionScheme js.String
	// Robustness is "MediaKeySystemMediaCapability.robustness"
	//
	// Optional, defaults to "".
	Robustness js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaKeySystemMediaCapability with all fields set.
func (p MediaKeySystemMediaCapability) FromRef(ref js.Ref) MediaKeySystemMediaCapability {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaKeySystemMediaCapability in the application heap.
func (p MediaKeySystemMediaCapability) New() js.Ref {
	return bindings.MediaKeySystemMediaCapabilityJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaKeySystemMediaCapability) UpdateFrom(ref js.Ref) {
	bindings.MediaKeySystemMediaCapabilityJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaKeySystemMediaCapability) Update(ref js.Ref) {
	bindings.MediaKeySystemMediaCapabilityJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaKeysRequirement uint32

const (
	_ MediaKeysRequirement = iota

	MediaKeysRequirement_REQUIRED
	MediaKeysRequirement_OPTIONAL
	MediaKeysRequirement_NOT_ALLOWED
)

func (MediaKeysRequirement) FromRef(str js.Ref) MediaKeysRequirement {
	return MediaKeysRequirement(bindings.ConstOfMediaKeysRequirement(str))
}

func (x MediaKeysRequirement) String() (string, bool) {
	switch x {
	case MediaKeysRequirement_REQUIRED:
		return "required", true
	case MediaKeysRequirement_OPTIONAL:
		return "optional", true
	case MediaKeysRequirement_NOT_ALLOWED:
		return "not-allowed", true
	default:
		return "", false
	}
}

type MediaKeySystemConfiguration struct {
	// Label is "MediaKeySystemConfiguration.label"
	//
	// Optional, defaults to "".
	Label js.String
	// InitDataTypes is "MediaKeySystemConfiguration.initDataTypes"
	//
	// Optional, defaults to [].
	InitDataTypes js.Array[js.String]
	// AudioCapabilities is "MediaKeySystemConfiguration.audioCapabilities"
	//
	// Optional, defaults to [].
	AudioCapabilities js.Array[MediaKeySystemMediaCapability]
	// VideoCapabilities is "MediaKeySystemConfiguration.videoCapabilities"
	//
	// Optional, defaults to [].
	VideoCapabilities js.Array[MediaKeySystemMediaCapability]
	// DistinctiveIdentifier is "MediaKeySystemConfiguration.distinctiveIdentifier"
	//
	// Optional, defaults to "optional".
	DistinctiveIdentifier MediaKeysRequirement
	// PersistentState is "MediaKeySystemConfiguration.persistentState"
	//
	// Optional, defaults to "optional".
	PersistentState MediaKeysRequirement
	// SessionTypes is "MediaKeySystemConfiguration.sessionTypes"
	//
	// Optional
	SessionTypes js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaKeySystemConfiguration with all fields set.
func (p MediaKeySystemConfiguration) FromRef(ref js.Ref) MediaKeySystemConfiguration {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaKeySystemConfiguration in the application heap.
func (p MediaKeySystemConfiguration) New() js.Ref {
	return bindings.MediaKeySystemConfigurationJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaKeySystemConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MediaKeySystemConfigurationJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaKeySystemConfiguration) Update(ref js.Ref) {
	bindings.MediaKeySystemConfigurationJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MediaKeySystemAccess struct {
	ref js.Ref
}

func (this MediaKeySystemAccess) Once() MediaKeySystemAccess {
	this.Ref().Once()
	return this
}

func (this MediaKeySystemAccess) Ref() js.Ref {
	return this.ref
}

func (this MediaKeySystemAccess) FromRef(ref js.Ref) MediaKeySystemAccess {
	this.ref = ref
	return this
}

func (this MediaKeySystemAccess) Free() {
	this.Ref().Free()
}

// KeySystem returns the value of property "MediaKeySystemAccess.keySystem".
//
// It returns ok=false if there is no such property.
func (this MediaKeySystemAccess) KeySystem() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaKeySystemAccessKeySystem(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGetConfiguration returns true if the method "MediaKeySystemAccess.getConfiguration" exists.
func (this MediaKeySystemAccess) HasGetConfiguration() bool {
	return js.True == bindings.HasMediaKeySystemAccessGetConfiguration(
		this.Ref(),
	)
}

// GetConfigurationFunc returns the method "MediaKeySystemAccess.getConfiguration".
func (this MediaKeySystemAccess) GetConfigurationFunc() (fn js.Func[func() MediaKeySystemConfiguration]) {
	return fn.FromRef(
		bindings.MediaKeySystemAccessGetConfigurationFunc(
			this.Ref(),
		),
	)
}

// GetConfiguration calls the method "MediaKeySystemAccess.getConfiguration".
func (this MediaKeySystemAccess) GetConfiguration() (ret MediaKeySystemConfiguration) {
	bindings.CallMediaKeySystemAccessGetConfiguration(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetConfiguration calls the method "MediaKeySystemAccess.getConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeySystemAccess) TryGetConfiguration() (ret MediaKeySystemConfiguration, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeySystemAccessGetConfiguration(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateMediaKeys returns true if the method "MediaKeySystemAccess.createMediaKeys" exists.
func (this MediaKeySystemAccess) HasCreateMediaKeys() bool {
	return js.True == bindings.HasMediaKeySystemAccessCreateMediaKeys(
		this.Ref(),
	)
}

// CreateMediaKeysFunc returns the method "MediaKeySystemAccess.createMediaKeys".
func (this MediaKeySystemAccess) CreateMediaKeysFunc() (fn js.Func[func() js.Promise[MediaKeys]]) {
	return fn.FromRef(
		bindings.MediaKeySystemAccessCreateMediaKeysFunc(
			this.Ref(),
		),
	)
}

// CreateMediaKeys calls the method "MediaKeySystemAccess.createMediaKeys".
func (this MediaKeySystemAccess) CreateMediaKeys() (ret js.Promise[MediaKeys]) {
	bindings.CallMediaKeySystemAccessCreateMediaKeys(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateMediaKeys calls the method "MediaKeySystemAccess.createMediaKeys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeySystemAccess) TryCreateMediaKeys() (ret js.Promise[MediaKeys], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeySystemAccessCreateMediaKeys(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OneOf_Bool_MediaTrackConstraints struct {
	ref js.Ref
}

func (x OneOf_Bool_MediaTrackConstraints) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Bool_MediaTrackConstraints) Free() {
	x.ref.Free()
}

func (x OneOf_Bool_MediaTrackConstraints) FromRef(ref js.Ref) OneOf_Bool_MediaTrackConstraints {
	return OneOf_Bool_MediaTrackConstraints{
		ref: ref,
	}
}

func (x OneOf_Bool_MediaTrackConstraints) Bool() bool {
	return x.ref == js.True
}

func (x OneOf_Bool_MediaTrackConstraints) MediaTrackConstraints() MediaTrackConstraints {
	var ret MediaTrackConstraints
	ret.UpdateFrom(x.ref)
	return ret
}

type MediaStreamConstraints struct {
	// Video is "MediaStreamConstraints.video"
	//
	// Optional, defaults to false.
	Video OneOf_Bool_MediaTrackConstraints
	// Audio is "MediaStreamConstraints.audio"
	//
	// Optional, defaults to false.
	Audio OneOf_Bool_MediaTrackConstraints
	// PeerIdentity is "MediaStreamConstraints.peerIdentity"
	//
	// Optional
	PeerIdentity js.String
	// PreferCurrentTab is "MediaStreamConstraints.preferCurrentTab"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_PreferCurrentTab MUST be set to true to make this field effective.
	PreferCurrentTab bool

	FFI_USE_PreferCurrentTab bool // for PreferCurrentTab.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaStreamConstraints with all fields set.
func (p MediaStreamConstraints) FromRef(ref js.Ref) MediaStreamConstraints {
	p.UpdateFrom(ref)
	return p
}

// New creates a new MediaStreamConstraints in the application heap.
func (p MediaStreamConstraints) New() js.Ref {
	return bindings.MediaStreamConstraintsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p MediaStreamConstraints) UpdateFrom(ref js.Ref) {
	bindings.MediaStreamConstraintsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p MediaStreamConstraints) Update(ref js.Ref) {
	bindings.MediaStreamConstraintsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type NavigatorUserMediaSuccessCallbackFunc func(this js.Ref, stream MediaStream) js.Ref

func (fn NavigatorUserMediaSuccessCallbackFunc) Register() js.Func[func(stream MediaStream)] {
	return js.RegisterCallback[func(stream MediaStream)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn NavigatorUserMediaSuccessCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		MediaStream{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NavigatorUserMediaSuccessCallback[T any] struct {
	Fn  func(arg T, this js.Ref, stream MediaStream) js.Ref
	Arg T
}

func (cb *NavigatorUserMediaSuccessCallback[T]) Register() js.Func[func(stream MediaStream)] {
	return js.RegisterCallback[func(stream MediaStream)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *NavigatorUserMediaSuccessCallback[T]) DispatchCallback(
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

		MediaStream{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NavigatorUserMediaErrorCallbackFunc func(this js.Ref, err DOMException) js.Ref

func (fn NavigatorUserMediaErrorCallbackFunc) Register() js.Func[func(err DOMException)] {
	return js.RegisterCallback[func(err DOMException)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn NavigatorUserMediaErrorCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		DOMException{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NavigatorUserMediaErrorCallback[T any] struct {
	Fn  func(arg T, this js.Ref, err DOMException) js.Ref
	Arg T
}

func (cb *NavigatorUserMediaErrorCallback[T]) Register() js.Func[func(err DOMException)] {
	return js.RegisterCallback[func(err DOMException)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *NavigatorUserMediaErrorCallback[T]) DispatchCallback(
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

		DOMException{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OpaqueProperty uint32

const (
	_ OpaqueProperty = iota

	OpaqueProperty_OPAQUE
)

func (OpaqueProperty) FromRef(str js.Ref) OpaqueProperty {
	return OpaqueProperty(bindings.ConstOfOpaqueProperty(str))
}

func (x OpaqueProperty) String() (string, bool) {
	switch x {
	case OpaqueProperty_OPAQUE:
		return "opaque", true
	default:
		return "", false
	}
}

type OneOf_Uint32_OpaqueProperty struct {
	ref js.Ref
}

func (x OneOf_Uint32_OpaqueProperty) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Uint32_OpaqueProperty) Free() {
	x.ref.Free()
}

func (x OneOf_Uint32_OpaqueProperty) FromRef(ref js.Ref) OneOf_Uint32_OpaqueProperty {
	return OneOf_Uint32_OpaqueProperty{
		ref: ref,
	}
}

func (x OneOf_Uint32_OpaqueProperty) Uint32() uint32 {
	return js.Number[uint32]{}.FromRef(x.ref).Get()
}

func (x OneOf_Uint32_OpaqueProperty) OpaqueProperty() OpaqueProperty {
	return OpaqueProperty(0).FromRef(x.ref)
}

type FencedFrameConfigSize = OneOf_Uint32_OpaqueProperty

type FencedFrameConfig struct {
	ref js.Ref
}

func (this FencedFrameConfig) Once() FencedFrameConfig {
	this.Ref().Once()
	return this
}

func (this FencedFrameConfig) Ref() js.Ref {
	return this.ref
}

func (this FencedFrameConfig) FromRef(ref js.Ref) FencedFrameConfig {
	this.ref = ref
	return this
}

func (this FencedFrameConfig) Free() {
	this.Ref().Free()
}

// ContainerWidth returns the value of property "FencedFrameConfig.containerWidth".
//
// It returns ok=false if there is no such property.
func (this FencedFrameConfig) ContainerWidth() (ret FencedFrameConfigSize, ok bool) {
	ok = js.True == bindings.GetFencedFrameConfigContainerWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContainerHeight returns the value of property "FencedFrameConfig.containerHeight".
//
// It returns ok=false if there is no such property.
func (this FencedFrameConfig) ContainerHeight() (ret FencedFrameConfigSize, ok bool) {
	ok = js.True == bindings.GetFencedFrameConfigContainerHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContentWidth returns the value of property "FencedFrameConfig.contentWidth".
//
// It returns ok=false if there is no such property.
func (this FencedFrameConfig) ContentWidth() (ret FencedFrameConfigSize, ok bool) {
	ok = js.True == bindings.GetFencedFrameConfigContentWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContentHeight returns the value of property "FencedFrameConfig.contentHeight".
//
// It returns ok=false if there is no such property.
func (this FencedFrameConfig) ContentHeight() (ret FencedFrameConfigSize, ok bool) {
	ok = js.True == bindings.GetFencedFrameConfigContentHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasSetSharedStorageContext returns true if the method "FencedFrameConfig.setSharedStorageContext" exists.
func (this FencedFrameConfig) HasSetSharedStorageContext() bool {
	return js.True == bindings.HasFencedFrameConfigSetSharedStorageContext(
		this.Ref(),
	)
}

// SetSharedStorageContextFunc returns the method "FencedFrameConfig.setSharedStorageContext".
func (this FencedFrameConfig) SetSharedStorageContextFunc() (fn js.Func[func(contextString js.String)]) {
	return fn.FromRef(
		bindings.FencedFrameConfigSetSharedStorageContextFunc(
			this.Ref(),
		),
	)
}

// SetSharedStorageContext calls the method "FencedFrameConfig.setSharedStorageContext".
func (this FencedFrameConfig) SetSharedStorageContext(contextString js.String) (ret js.Void) {
	bindings.CallFencedFrameConfigSetSharedStorageContext(
		this.Ref(), js.Pointer(&ret),
		contextString.Ref(),
	)

	return
}

// TrySetSharedStorageContext calls the method "FencedFrameConfig.setSharedStorageContext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FencedFrameConfig) TrySetSharedStorageContext(contextString js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFencedFrameConfigSetSharedStorageContext(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		contextString.Ref(),
	)

	return
}

type OneOf_String_FencedFrameConfig struct {
	ref js.Ref
}

func (x OneOf_String_FencedFrameConfig) Ref() js.Ref {
	return x.ref
}

func (x OneOf_String_FencedFrameConfig) Free() {
	x.ref.Free()
}

func (x OneOf_String_FencedFrameConfig) FromRef(ref js.Ref) OneOf_String_FencedFrameConfig {
	return OneOf_String_FencedFrameConfig{
		ref: ref,
	}
}

func (x OneOf_String_FencedFrameConfig) String() js.String {
	return js.String{}.FromRef(x.ref)
}

func (x OneOf_String_FencedFrameConfig) FencedFrameConfig() FencedFrameConfig {
	return FencedFrameConfig{}.FromRef(x.ref)
}

type ShareData struct {
	// Files is "ShareData.files"
	//
	// Optional
	Files js.Array[File]
	// Title is "ShareData.title"
	//
	// Optional
	Title js.String
	// Text is "ShareData.text"
	//
	// Optional
	Text js.String
	// Url is "ShareData.url"
	//
	// Optional
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ShareData with all fields set.
func (p ShareData) FromRef(ref js.Ref) ShareData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ShareData in the application heap.
func (p ShareData) New() js.Ref {
	return bindings.ShareDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p ShareData) UpdateFrom(ref js.Ref) {
	bindings.ShareDataJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p ShareData) Update(ref js.Ref) {
	bindings.ShareDataJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type MIDIInputMap struct {
	ref js.Ref
}

func (this MIDIInputMap) Once() MIDIInputMap {
	this.Ref().Once()
	return this
}

func (this MIDIInputMap) Ref() js.Ref {
	return this.ref
}

func (this MIDIInputMap) FromRef(ref js.Ref) MIDIInputMap {
	this.ref = ref
	return this
}

func (this MIDIInputMap) Free() {
	this.Ref().Free()
}

type MIDIOutputMap struct {
	ref js.Ref
}

func (this MIDIOutputMap) Once() MIDIOutputMap {
	this.Ref().Once()
	return this
}

func (this MIDIOutputMap) Ref() js.Ref {
	return this.ref
}

func (this MIDIOutputMap) FromRef(ref js.Ref) MIDIOutputMap {
	this.ref = ref
	return this
}

func (this MIDIOutputMap) Free() {
	this.Ref().Free()
}

type MIDIAccess struct {
	EventTarget
}

func (this MIDIAccess) Once() MIDIAccess {
	this.Ref().Once()
	return this
}

func (this MIDIAccess) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this MIDIAccess) FromRef(ref js.Ref) MIDIAccess {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this MIDIAccess) Free() {
	this.Ref().Free()
}

// Inputs returns the value of property "MIDIAccess.inputs".
//
// It returns ok=false if there is no such property.
func (this MIDIAccess) Inputs() (ret MIDIInputMap, ok bool) {
	ok = js.True == bindings.GetMIDIAccessInputs(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Outputs returns the value of property "MIDIAccess.outputs".
//
// It returns ok=false if there is no such property.
func (this MIDIAccess) Outputs() (ret MIDIOutputMap, ok bool) {
	ok = js.True == bindings.GetMIDIAccessOutputs(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SysexEnabled returns the value of property "MIDIAccess.sysexEnabled".
//
// It returns ok=false if there is no such property.
func (this MIDIAccess) SysexEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetMIDIAccessSysexEnabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

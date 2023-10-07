// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

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
func (p *NavigationNavigateOptions) UpdateFrom(ref js.Ref) {
	bindings.NavigationNavigateOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NavigationNavigateOptions) Update(ref js.Ref) {
	bindings.NavigationNavigateOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NavigationNavigateOptions) FreeMembers(recursive bool) {
	js.Free(
		p.State.Ref(),
		p.Info.Ref(),
	)
	p.State = p.State.FromRef(js.Undefined)
	p.Info = p.Info.FromRef(js.Undefined)
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
func (p *NavigationReloadOptions) UpdateFrom(ref js.Ref) {
	bindings.NavigationReloadOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NavigationReloadOptions) Update(ref js.Ref) {
	bindings.NavigationReloadOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NavigationReloadOptions) FreeMembers(recursive bool) {
	js.Free(
		p.State.Ref(),
		p.Info.Ref(),
	)
	p.State = p.State.FromRef(js.Undefined)
	p.Info = p.Info.FromRef(js.Undefined)
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
func (p *NavigationOptions) UpdateFrom(ref js.Ref) {
	bindings.NavigationOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *NavigationOptions) Update(ref js.Ref) {
	bindings.NavigationOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *NavigationOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Info.Ref(),
	)
	p.Info = p.Info.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// NavigationType returns the value of property "NavigationTransition.navigationType".
//
// It returns ok=false if there is no such property.
func (this NavigationTransition) NavigationType() (ret NavigationType, ok bool) {
	ok = js.True == bindings.GetNavigationTransitionNavigationType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// From returns the value of property "NavigationTransition.from".
//
// It returns ok=false if there is no such property.
func (this NavigationTransition) From() (ret NavigationHistoryEntry, ok bool) {
	ok = js.True == bindings.GetNavigationTransitionFrom(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Finished returns the value of property "NavigationTransition.finished".
//
// It returns ok=false if there is no such property.
func (this NavigationTransition) Finished() (ret js.Promise[js.Void], ok bool) {
	ok = js.True == bindings.GetNavigationTransitionFinished(
		this.ref, js.Pointer(&ret),
	)
	return
}

type Navigation struct {
	EventTarget
}

func (this Navigation) Once() Navigation {
	this.ref.Once()
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
	this.ref.Free()
}

// CurrentEntry returns the value of property "Navigation.currentEntry".
//
// It returns ok=false if there is no such property.
func (this Navigation) CurrentEntry() (ret NavigationHistoryEntry, ok bool) {
	ok = js.True == bindings.GetNavigationCurrentEntry(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Transition returns the value of property "Navigation.transition".
//
// It returns ok=false if there is no such property.
func (this Navigation) Transition() (ret NavigationTransition, ok bool) {
	ok = js.True == bindings.GetNavigationTransition(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CanGoBack returns the value of property "Navigation.canGoBack".
//
// It returns ok=false if there is no such property.
func (this Navigation) CanGoBack() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigationCanGoBack(
		this.ref, js.Pointer(&ret),
	)
	return
}

// CanGoForward returns the value of property "Navigation.canGoForward".
//
// It returns ok=false if there is no such property.
func (this Navigation) CanGoForward() (ret bool, ok bool) {
	ok = js.True == bindings.GetNavigationCanGoForward(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncEntries returns true if the method "Navigation.entries" exists.
func (this Navigation) HasFuncEntries() bool {
	return js.True == bindings.HasFuncNavigationEntries(
		this.ref,
	)
}

// FuncEntries returns the method "Navigation.entries".
func (this Navigation) FuncEntries() (fn js.Func[func() js.Array[NavigationHistoryEntry]]) {
	bindings.FuncNavigationEntries(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Entries calls the method "Navigation.entries".
func (this Navigation) Entries() (ret js.Array[NavigationHistoryEntry]) {
	bindings.CallNavigationEntries(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryEntries calls the method "Navigation.entries"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryEntries() (ret js.Array[NavigationHistoryEntry], exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationEntries(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncUpdateCurrentEntry returns true if the method "Navigation.updateCurrentEntry" exists.
func (this Navigation) HasFuncUpdateCurrentEntry() bool {
	return js.True == bindings.HasFuncNavigationUpdateCurrentEntry(
		this.ref,
	)
}

// FuncUpdateCurrentEntry returns the method "Navigation.updateCurrentEntry".
func (this Navigation) FuncUpdateCurrentEntry() (fn js.Func[func(options NavigationUpdateCurrentEntryOptions)]) {
	bindings.FuncNavigationUpdateCurrentEntry(
		this.ref, js.Pointer(&fn),
	)
	return
}

// UpdateCurrentEntry calls the method "Navigation.updateCurrentEntry".
func (this Navigation) UpdateCurrentEntry(options NavigationUpdateCurrentEntryOptions) (ret js.Void) {
	bindings.CallNavigationUpdateCurrentEntry(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryUpdateCurrentEntry calls the method "Navigation.updateCurrentEntry"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryUpdateCurrentEntry(options NavigationUpdateCurrentEntryOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationUpdateCurrentEntry(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncNavigate returns true if the method "Navigation.navigate" exists.
func (this Navigation) HasFuncNavigate() bool {
	return js.True == bindings.HasFuncNavigationNavigate(
		this.ref,
	)
}

// FuncNavigate returns the method "Navigation.navigate".
func (this Navigation) FuncNavigate() (fn js.Func[func(url js.String, options NavigationNavigateOptions) NavigationResult]) {
	bindings.FuncNavigationNavigate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Navigate calls the method "Navigation.navigate".
func (this Navigation) Navigate(url js.String, options NavigationNavigateOptions) (ret NavigationResult) {
	bindings.CallNavigationNavigate(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncNavigate1 returns true if the method "Navigation.navigate" exists.
func (this Navigation) HasFuncNavigate1() bool {
	return js.True == bindings.HasFuncNavigationNavigate1(
		this.ref,
	)
}

// FuncNavigate1 returns the method "Navigation.navigate".
func (this Navigation) FuncNavigate1() (fn js.Func[func(url js.String) NavigationResult]) {
	bindings.FuncNavigationNavigate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Navigate1 calls the method "Navigation.navigate".
func (this Navigation) Navigate1(url js.String) (ret NavigationResult) {
	bindings.CallNavigationNavigate1(
		this.ref, js.Pointer(&ret),
		url.Ref(),
	)

	return
}

// TryNavigate1 calls the method "Navigation.navigate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryNavigate1(url js.String) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationNavigate1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		url.Ref(),
	)

	return
}

// HasFuncReload returns true if the method "Navigation.reload" exists.
func (this Navigation) HasFuncReload() bool {
	return js.True == bindings.HasFuncNavigationReload(
		this.ref,
	)
}

// FuncReload returns the method "Navigation.reload".
func (this Navigation) FuncReload() (fn js.Func[func(options NavigationReloadOptions) NavigationResult]) {
	bindings.FuncNavigationReload(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reload calls the method "Navigation.reload".
func (this Navigation) Reload(options NavigationReloadOptions) (ret NavigationResult) {
	bindings.CallNavigationReload(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryReload calls the method "Navigation.reload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryReload(options NavigationReloadOptions) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationReload(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncReload1 returns true if the method "Navigation.reload" exists.
func (this Navigation) HasFuncReload1() bool {
	return js.True == bindings.HasFuncNavigationReload1(
		this.ref,
	)
}

// FuncReload1 returns the method "Navigation.reload".
func (this Navigation) FuncReload1() (fn js.Func[func() NavigationResult]) {
	bindings.FuncNavigationReload1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reload1 calls the method "Navigation.reload".
func (this Navigation) Reload1() (ret NavigationResult) {
	bindings.CallNavigationReload1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReload1 calls the method "Navigation.reload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryReload1() (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationReload1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncTraverseTo returns true if the method "Navigation.traverseTo" exists.
func (this Navigation) HasFuncTraverseTo() bool {
	return js.True == bindings.HasFuncNavigationTraverseTo(
		this.ref,
	)
}

// FuncTraverseTo returns the method "Navigation.traverseTo".
func (this Navigation) FuncTraverseTo() (fn js.Func[func(key js.String, options NavigationOptions) NavigationResult]) {
	bindings.FuncNavigationTraverseTo(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TraverseTo calls the method "Navigation.traverseTo".
func (this Navigation) TraverseTo(key js.String, options NavigationOptions) (ret NavigationResult) {
	bindings.CallNavigationTraverseTo(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncTraverseTo1 returns true if the method "Navigation.traverseTo" exists.
func (this Navigation) HasFuncTraverseTo1() bool {
	return js.True == bindings.HasFuncNavigationTraverseTo1(
		this.ref,
	)
}

// FuncTraverseTo1 returns the method "Navigation.traverseTo".
func (this Navigation) FuncTraverseTo1() (fn js.Func[func(key js.String) NavigationResult]) {
	bindings.FuncNavigationTraverseTo1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// TraverseTo1 calls the method "Navigation.traverseTo".
func (this Navigation) TraverseTo1(key js.String) (ret NavigationResult) {
	bindings.CallNavigationTraverseTo1(
		this.ref, js.Pointer(&ret),
		key.Ref(),
	)

	return
}

// TryTraverseTo1 calls the method "Navigation.traverseTo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryTraverseTo1(key js.String) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationTraverseTo1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		key.Ref(),
	)

	return
}

// HasFuncBack returns true if the method "Navigation.back" exists.
func (this Navigation) HasFuncBack() bool {
	return js.True == bindings.HasFuncNavigationBack(
		this.ref,
	)
}

// FuncBack returns the method "Navigation.back".
func (this Navigation) FuncBack() (fn js.Func[func(options NavigationOptions) NavigationResult]) {
	bindings.FuncNavigationBack(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Back calls the method "Navigation.back".
func (this Navigation) Back(options NavigationOptions) (ret NavigationResult) {
	bindings.CallNavigationBack(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryBack calls the method "Navigation.back"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryBack(options NavigationOptions) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationBack(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncBack1 returns true if the method "Navigation.back" exists.
func (this Navigation) HasFuncBack1() bool {
	return js.True == bindings.HasFuncNavigationBack1(
		this.ref,
	)
}

// FuncBack1 returns the method "Navigation.back".
func (this Navigation) FuncBack1() (fn js.Func[func() NavigationResult]) {
	bindings.FuncNavigationBack1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Back1 calls the method "Navigation.back".
func (this Navigation) Back1() (ret NavigationResult) {
	bindings.CallNavigationBack1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryBack1 calls the method "Navigation.back"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryBack1() (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationBack1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncForward returns true if the method "Navigation.forward" exists.
func (this Navigation) HasFuncForward() bool {
	return js.True == bindings.HasFuncNavigationForward(
		this.ref,
	)
}

// FuncForward returns the method "Navigation.forward".
func (this Navigation) FuncForward() (fn js.Func[func(options NavigationOptions) NavigationResult]) {
	bindings.FuncNavigationForward(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Forward calls the method "Navigation.forward".
func (this Navigation) Forward(options NavigationOptions) (ret NavigationResult) {
	bindings.CallNavigationForward(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryForward calls the method "Navigation.forward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryForward(options NavigationOptions) (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationForward(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncForward1 returns true if the method "Navigation.forward" exists.
func (this Navigation) HasFuncForward1() bool {
	return js.True == bindings.HasFuncNavigationForward1(
		this.ref,
	)
}

// FuncForward1 returns the method "Navigation.forward".
func (this Navigation) FuncForward1() (fn js.Func[func() NavigationResult]) {
	bindings.FuncNavigationForward1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Forward1 calls the method "Navigation.forward".
func (this Navigation) Forward1() (ret NavigationResult) {
	bindings.CallNavigationForward1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryForward1 calls the method "Navigation.forward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Navigation) TryForward1() (ret NavigationResult, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigationForward1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
		js.ThrowInvalidCallbackInvocation()
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
func (p *ElementDefinitionOptions) UpdateFrom(ref js.Ref) {
	bindings.ElementDefinitionOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ElementDefinitionOptions) Update(ref js.Ref) {
	bindings.ElementDefinitionOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ElementDefinitionOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Extends.Ref(),
	)
	p.Extends = p.Extends.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// HasFuncDefine returns true if the method "CustomElementRegistry.define" exists.
func (this CustomElementRegistry) HasFuncDefine() bool {
	return js.True == bindings.HasFuncCustomElementRegistryDefine(
		this.ref,
	)
}

// FuncDefine returns the method "CustomElementRegistry.define".
func (this CustomElementRegistry) FuncDefine() (fn js.Func[func(name js.String, constructor js.Func[func() HTMLElement], options ElementDefinitionOptions)]) {
	bindings.FuncCustomElementRegistryDefine(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Define calls the method "CustomElementRegistry.define".
func (this CustomElementRegistry) Define(name js.String, constructor js.Func[func() HTMLElement], options ElementDefinitionOptions) (ret js.Void) {
	bindings.CallCustomElementRegistryDefine(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		constructor.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncDefine1 returns true if the method "CustomElementRegistry.define" exists.
func (this CustomElementRegistry) HasFuncDefine1() bool {
	return js.True == bindings.HasFuncCustomElementRegistryDefine1(
		this.ref,
	)
}

// FuncDefine1 returns the method "CustomElementRegistry.define".
func (this CustomElementRegistry) FuncDefine1() (fn js.Func[func(name js.String, constructor js.Func[func() HTMLElement])]) {
	bindings.FuncCustomElementRegistryDefine1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Define1 calls the method "CustomElementRegistry.define".
func (this CustomElementRegistry) Define1(name js.String, constructor js.Func[func() HTMLElement]) (ret js.Void) {
	bindings.CallCustomElementRegistryDefine1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		constructor.Ref(),
	)

	return
}

// HasFuncGet returns true if the method "CustomElementRegistry.get" exists.
func (this CustomElementRegistry) HasFuncGet() bool {
	return js.True == bindings.HasFuncCustomElementRegistryGet(
		this.ref,
	)
}

// FuncGet returns the method "CustomElementRegistry.get".
func (this CustomElementRegistry) FuncGet() (fn js.Func[func(name js.String) OneOf_FuncCustomElementConstructor_undefined]) {
	bindings.FuncCustomElementRegistryGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "CustomElementRegistry.get".
func (this CustomElementRegistry) Get(name js.String) (ret OneOf_FuncCustomElementConstructor_undefined) {
	bindings.CallCustomElementRegistryGet(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryGet calls the method "CustomElementRegistry.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomElementRegistry) TryGet(name js.String) (ret OneOf_FuncCustomElementConstructor_undefined, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomElementRegistryGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncGetName returns true if the method "CustomElementRegistry.getName" exists.
func (this CustomElementRegistry) HasFuncGetName() bool {
	return js.True == bindings.HasFuncCustomElementRegistryGetName(
		this.ref,
	)
}

// FuncGetName returns the method "CustomElementRegistry.getName".
func (this CustomElementRegistry) FuncGetName() (fn js.Func[func(constructor js.Func[func() HTMLElement]) js.String]) {
	bindings.FuncCustomElementRegistryGetName(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetName calls the method "CustomElementRegistry.getName".
func (this CustomElementRegistry) GetName(constructor js.Func[func() HTMLElement]) (ret js.String) {
	bindings.CallCustomElementRegistryGetName(
		this.ref, js.Pointer(&ret),
		constructor.Ref(),
	)

	return
}

// TryGetName calls the method "CustomElementRegistry.getName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomElementRegistry) TryGetName(constructor js.Func[func() HTMLElement]) (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomElementRegistryGetName(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		constructor.Ref(),
	)

	return
}

// HasFuncWhenDefined returns true if the method "CustomElementRegistry.whenDefined" exists.
func (this CustomElementRegistry) HasFuncWhenDefined() bool {
	return js.True == bindings.HasFuncCustomElementRegistryWhenDefined(
		this.ref,
	)
}

// FuncWhenDefined returns the method "CustomElementRegistry.whenDefined".
func (this CustomElementRegistry) FuncWhenDefined() (fn js.Func[func(name js.String) js.Promise[js.Func[func() HTMLElement]]]) {
	bindings.FuncCustomElementRegistryWhenDefined(
		this.ref, js.Pointer(&fn),
	)
	return
}

// WhenDefined calls the method "CustomElementRegistry.whenDefined".
func (this CustomElementRegistry) WhenDefined(name js.String) (ret js.Promise[js.Func[func() HTMLElement]]) {
	bindings.CallCustomElementRegistryWhenDefined(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryWhenDefined calls the method "CustomElementRegistry.whenDefined"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomElementRegistry) TryWhenDefined(name js.String) (ret js.Promise[js.Func[func() HTMLElement]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomElementRegistryWhenDefined(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncUpgrade returns true if the method "CustomElementRegistry.upgrade" exists.
func (this CustomElementRegistry) HasFuncUpgrade() bool {
	return js.True == bindings.HasFuncCustomElementRegistryUpgrade(
		this.ref,
	)
}

// FuncUpgrade returns the method "CustomElementRegistry.upgrade".
func (this CustomElementRegistry) FuncUpgrade() (fn js.Func[func(root Node)]) {
	bindings.FuncCustomElementRegistryUpgrade(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Upgrade calls the method "CustomElementRegistry.upgrade".
func (this CustomElementRegistry) Upgrade(root Node) (ret js.Void) {
	bindings.CallCustomElementRegistryUpgrade(
		this.ref, js.Pointer(&ret),
		root.Ref(),
	)

	return
}

// TryUpgrade calls the method "CustomElementRegistry.upgrade"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this CustomElementRegistry) TryUpgrade(root Node) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCustomElementRegistryUpgrade(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Pressed returns the value of property "GamepadButton.pressed".
//
// It returns ok=false if there is no such property.
func (this GamepadButton) Pressed() (ret bool, ok bool) {
	ok = js.True == bindings.GetGamepadButtonPressed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Touched returns the value of property "GamepadButton.touched".
//
// It returns ok=false if there is no such property.
func (this GamepadButton) Touched() (ret bool, ok bool) {
	ok = js.True == bindings.GetGamepadButtonTouched(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Value returns the value of property "GamepadButton.value".
//
// It returns ok=false if there is no such property.
func (this GamepadButton) Value() (ret float64, ok bool) {
	ok = js.True == bindings.GetGamepadButtonValue(
		this.ref, js.Pointer(&ret),
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
func (p *GamepadEffectParameters) UpdateFrom(ref js.Ref) {
	bindings.GamepadEffectParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *GamepadEffectParameters) Update(ref js.Ref) {
	bindings.GamepadEffectParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *GamepadEffectParameters) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Type returns the value of property "GamepadHapticActuator.type".
//
// It returns ok=false if there is no such property.
func (this GamepadHapticActuator) Type() (ret GamepadHapticActuatorType, ok bool) {
	ok = js.True == bindings.GetGamepadHapticActuatorType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncCanPlayEffectType returns true if the method "GamepadHapticActuator.canPlayEffectType" exists.
func (this GamepadHapticActuator) HasFuncCanPlayEffectType() bool {
	return js.True == bindings.HasFuncGamepadHapticActuatorCanPlayEffectType(
		this.ref,
	)
}

// FuncCanPlayEffectType returns the method "GamepadHapticActuator.canPlayEffectType".
func (this GamepadHapticActuator) FuncCanPlayEffectType() (fn js.Func[func(typ GamepadHapticEffectType) bool]) {
	bindings.FuncGamepadHapticActuatorCanPlayEffectType(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CanPlayEffectType calls the method "GamepadHapticActuator.canPlayEffectType".
func (this GamepadHapticActuator) CanPlayEffectType(typ GamepadHapticEffectType) (ret bool) {
	bindings.CallGamepadHapticActuatorCanPlayEffectType(
		this.ref, js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryCanPlayEffectType calls the method "GamepadHapticActuator.canPlayEffectType"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GamepadHapticActuator) TryCanPlayEffectType(typ GamepadHapticEffectType) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGamepadHapticActuatorCanPlayEffectType(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasFuncPlayEffect returns true if the method "GamepadHapticActuator.playEffect" exists.
func (this GamepadHapticActuator) HasFuncPlayEffect() bool {
	return js.True == bindings.HasFuncGamepadHapticActuatorPlayEffect(
		this.ref,
	)
}

// FuncPlayEffect returns the method "GamepadHapticActuator.playEffect".
func (this GamepadHapticActuator) FuncPlayEffect() (fn js.Func[func(typ GamepadHapticEffectType, params GamepadEffectParameters) js.Promise[GamepadHapticsResult]]) {
	bindings.FuncGamepadHapticActuatorPlayEffect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PlayEffect calls the method "GamepadHapticActuator.playEffect".
func (this GamepadHapticActuator) PlayEffect(typ GamepadHapticEffectType, params GamepadEffectParameters) (ret js.Promise[GamepadHapticsResult]) {
	bindings.CallGamepadHapticActuatorPlayEffect(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
		js.Pointer(&params),
	)

	return
}

// HasFuncPlayEffect1 returns true if the method "GamepadHapticActuator.playEffect" exists.
func (this GamepadHapticActuator) HasFuncPlayEffect1() bool {
	return js.True == bindings.HasFuncGamepadHapticActuatorPlayEffect1(
		this.ref,
	)
}

// FuncPlayEffect1 returns the method "GamepadHapticActuator.playEffect".
func (this GamepadHapticActuator) FuncPlayEffect1() (fn js.Func[func(typ GamepadHapticEffectType) js.Promise[GamepadHapticsResult]]) {
	bindings.FuncGamepadHapticActuatorPlayEffect1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PlayEffect1 calls the method "GamepadHapticActuator.playEffect".
func (this GamepadHapticActuator) PlayEffect1(typ GamepadHapticEffectType) (ret js.Promise[GamepadHapticsResult]) {
	bindings.CallGamepadHapticActuatorPlayEffect1(
		this.ref, js.Pointer(&ret),
		uint32(typ),
	)

	return
}

// TryPlayEffect1 calls the method "GamepadHapticActuator.playEffect"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GamepadHapticActuator) TryPlayEffect1(typ GamepadHapticEffectType) (ret js.Promise[GamepadHapticsResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGamepadHapticActuatorPlayEffect1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(typ),
	)

	return
}

// HasFuncPulse returns true if the method "GamepadHapticActuator.pulse" exists.
func (this GamepadHapticActuator) HasFuncPulse() bool {
	return js.True == bindings.HasFuncGamepadHapticActuatorPulse(
		this.ref,
	)
}

// FuncPulse returns the method "GamepadHapticActuator.pulse".
func (this GamepadHapticActuator) FuncPulse() (fn js.Func[func(value float64, duration float64) js.Promise[js.Boolean]]) {
	bindings.FuncGamepadHapticActuatorPulse(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Pulse calls the method "GamepadHapticActuator.pulse".
func (this GamepadHapticActuator) Pulse(value float64, duration float64) (ret js.Promise[js.Boolean]) {
	bindings.CallGamepadHapticActuatorPulse(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		float64(value),
		float64(duration),
	)

	return
}

// HasFuncReset returns true if the method "GamepadHapticActuator.reset" exists.
func (this GamepadHapticActuator) HasFuncReset() bool {
	return js.True == bindings.HasFuncGamepadHapticActuatorReset(
		this.ref,
	)
}

// FuncReset returns the method "GamepadHapticActuator.reset".
func (this GamepadHapticActuator) FuncReset() (fn js.Func[func() js.Promise[GamepadHapticsResult]]) {
	bindings.FuncGamepadHapticActuatorReset(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Reset calls the method "GamepadHapticActuator.reset".
func (this GamepadHapticActuator) Reset() (ret js.Promise[GamepadHapticsResult]) {
	bindings.CallGamepadHapticActuatorReset(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReset calls the method "GamepadHapticActuator.reset"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GamepadHapticActuator) TryReset() (ret js.Promise[GamepadHapticsResult], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGamepadHapticActuatorReset(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type GamepadPose struct {
	ref js.Ref
}

func (this GamepadPose) Once() GamepadPose {
	this.ref.Once()
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
	this.ref.Free()
}

// HasOrientation returns the value of property "GamepadPose.hasOrientation".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) HasOrientation() (ret bool, ok bool) {
	ok = js.True == bindings.GetGamepadPoseHasOrientation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasPosition returns the value of property "GamepadPose.hasPosition".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) HasPosition() (ret bool, ok bool) {
	ok = js.True == bindings.GetGamepadPoseHasPosition(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Position returns the value of property "GamepadPose.position".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) Position() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPosePosition(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LinearVelocity returns the value of property "GamepadPose.linearVelocity".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) LinearVelocity() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPoseLinearVelocity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// LinearAcceleration returns the value of property "GamepadPose.linearAcceleration".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) LinearAcceleration() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPoseLinearAcceleration(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Orientation returns the value of property "GamepadPose.orientation".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) Orientation() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPoseOrientation(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AngularVelocity returns the value of property "GamepadPose.angularVelocity".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) AngularVelocity() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPoseAngularVelocity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AngularAcceleration returns the value of property "GamepadPose.angularAcceleration".
//
// It returns ok=false if there is no such property.
func (this GamepadPose) AngularAcceleration() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadPoseAngularAcceleration(
		this.ref, js.Pointer(&ret),
	)
	return
}

type GamepadTouch struct {
	ref js.Ref
}

func (this GamepadTouch) Once() GamepadTouch {
	this.ref.Once()
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
	this.ref.Free()
}

// TouchId returns the value of property "GamepadTouch.touchId".
//
// It returns ok=false if there is no such property.
func (this GamepadTouch) TouchId() (ret uint32, ok bool) {
	ok = js.True == bindings.GetGamepadTouchTouchId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SurfaceId returns the value of property "GamepadTouch.surfaceId".
//
// It returns ok=false if there is no such property.
func (this GamepadTouch) SurfaceId() (ret uint8, ok bool) {
	ok = js.True == bindings.GetGamepadTouchSurfaceId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Position returns the value of property "GamepadTouch.position".
//
// It returns ok=false if there is no such property.
func (this GamepadTouch) Position() (ret js.TypedArray[float32], ok bool) {
	ok = js.True == bindings.GetGamepadTouchPosition(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SurfaceDimensions returns the value of property "GamepadTouch.surfaceDimensions".
//
// It returns ok=false if there is no such property.
func (this GamepadTouch) SurfaceDimensions() (ret js.TypedArray[uint32], ok bool) {
	ok = js.True == bindings.GetGamepadTouchSurfaceDimensions(
		this.ref, js.Pointer(&ret),
	)
	return
}

type Gamepad struct {
	ref js.Ref
}

func (this Gamepad) Once() Gamepad {
	this.ref.Once()
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
	this.ref.Free()
}

// Id returns the value of property "Gamepad.id".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Id() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGamepadId(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Index returns the value of property "Gamepad.index".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Index() (ret int32, ok bool) {
	ok = js.True == bindings.GetGamepadIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Connected returns the value of property "Gamepad.connected".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Connected() (ret bool, ok bool) {
	ok = js.True == bindings.GetGamepadConnected(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Timestamp returns the value of property "Gamepad.timestamp".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Timestamp() (ret DOMHighResTimeStamp, ok bool) {
	ok = js.True == bindings.GetGamepadTimestamp(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Mapping returns the value of property "Gamepad.mapping".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Mapping() (ret GamepadMappingType, ok bool) {
	ok = js.True == bindings.GetGamepadMapping(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Axes returns the value of property "Gamepad.axes".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Axes() (ret js.FrozenArray[float64], ok bool) {
	ok = js.True == bindings.GetGamepadAxes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Buttons returns the value of property "Gamepad.buttons".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Buttons() (ret js.FrozenArray[GamepadButton], ok bool) {
	ok = js.True == bindings.GetGamepadButtons(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Hand returns the value of property "Gamepad.hand".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Hand() (ret GamepadHand, ok bool) {
	ok = js.True == bindings.GetGamepadHand(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HapticActuators returns the value of property "Gamepad.hapticActuators".
//
// It returns ok=false if there is no such property.
func (this Gamepad) HapticActuators() (ret js.FrozenArray[GamepadHapticActuator], ok bool) {
	ok = js.True == bindings.GetGamepadHapticActuators(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Pose returns the value of property "Gamepad.pose".
//
// It returns ok=false if there is no such property.
func (this Gamepad) Pose() (ret GamepadPose, ok bool) {
	ok = js.True == bindings.GetGamepadPose(
		this.ref, js.Pointer(&ret),
	)
	return
}

// TouchEvents returns the value of property "Gamepad.touchEvents".
//
// It returns ok=false if there is no such property.
func (this Gamepad) TouchEvents() (ret js.FrozenArray[GamepadTouch], ok bool) {
	ok = js.True == bindings.GetGamepadTouchEvents(
		this.ref, js.Pointer(&ret),
	)
	return
}

// VibrationActuator returns the value of property "Gamepad.vibrationActuator".
//
// It returns ok=false if there is no such property.
func (this Gamepad) VibrationActuator() (ret GamepadHapticActuator, ok bool) {
	ok = js.True == bindings.GetGamepadVibrationActuator(
		this.ref, js.Pointer(&ret),
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
func (p *RelatedApplication) UpdateFrom(ref js.Ref) {
	bindings.RelatedApplicationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RelatedApplication) Update(ref js.Ref) {
	bindings.RelatedApplicationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RelatedApplication) FreeMembers(recursive bool) {
	js.Free(
		p.Platform.Ref(),
		p.Url.Ref(),
		p.Id.Ref(),
		p.Version.Ref(),
	)
	p.Platform = p.Platform.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Version = p.Version.FromRef(js.Undefined)
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
func (p *MediaKeySystemMediaCapability) UpdateFrom(ref js.Ref) {
	bindings.MediaKeySystemMediaCapabilityJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaKeySystemMediaCapability) Update(ref js.Ref) {
	bindings.MediaKeySystemMediaCapabilityJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaKeySystemMediaCapability) FreeMembers(recursive bool) {
	js.Free(
		p.ContentType.Ref(),
		p.EncryptionScheme.Ref(),
		p.Robustness.Ref(),
	)
	p.ContentType = p.ContentType.FromRef(js.Undefined)
	p.EncryptionScheme = p.EncryptionScheme.FromRef(js.Undefined)
	p.Robustness = p.Robustness.FromRef(js.Undefined)
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
func (p *MediaKeySystemConfiguration) UpdateFrom(ref js.Ref) {
	bindings.MediaKeySystemConfigurationJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaKeySystemConfiguration) Update(ref js.Ref) {
	bindings.MediaKeySystemConfigurationJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaKeySystemConfiguration) FreeMembers(recursive bool) {
	js.Free(
		p.Label.Ref(),
		p.InitDataTypes.Ref(),
		p.AudioCapabilities.Ref(),
		p.VideoCapabilities.Ref(),
		p.SessionTypes.Ref(),
	)
	p.Label = p.Label.FromRef(js.Undefined)
	p.InitDataTypes = p.InitDataTypes.FromRef(js.Undefined)
	p.AudioCapabilities = p.AudioCapabilities.FromRef(js.Undefined)
	p.VideoCapabilities = p.VideoCapabilities.FromRef(js.Undefined)
	p.SessionTypes = p.SessionTypes.FromRef(js.Undefined)
}

type MediaKeySystemAccess struct {
	ref js.Ref
}

func (this MediaKeySystemAccess) Once() MediaKeySystemAccess {
	this.ref.Once()
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
	this.ref.Free()
}

// KeySystem returns the value of property "MediaKeySystemAccess.keySystem".
//
// It returns ok=false if there is no such property.
func (this MediaKeySystemAccess) KeySystem() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaKeySystemAccessKeySystem(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGetConfiguration returns true if the method "MediaKeySystemAccess.getConfiguration" exists.
func (this MediaKeySystemAccess) HasFuncGetConfiguration() bool {
	return js.True == bindings.HasFuncMediaKeySystemAccessGetConfiguration(
		this.ref,
	)
}

// FuncGetConfiguration returns the method "MediaKeySystemAccess.getConfiguration".
func (this MediaKeySystemAccess) FuncGetConfiguration() (fn js.Func[func() MediaKeySystemConfiguration]) {
	bindings.FuncMediaKeySystemAccessGetConfiguration(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetConfiguration calls the method "MediaKeySystemAccess.getConfiguration".
func (this MediaKeySystemAccess) GetConfiguration() (ret MediaKeySystemConfiguration) {
	bindings.CallMediaKeySystemAccessGetConfiguration(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetConfiguration calls the method "MediaKeySystemAccess.getConfiguration"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeySystemAccess) TryGetConfiguration() (ret MediaKeySystemConfiguration, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeySystemAccessGetConfiguration(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateMediaKeys returns true if the method "MediaKeySystemAccess.createMediaKeys" exists.
func (this MediaKeySystemAccess) HasFuncCreateMediaKeys() bool {
	return js.True == bindings.HasFuncMediaKeySystemAccessCreateMediaKeys(
		this.ref,
	)
}

// FuncCreateMediaKeys returns the method "MediaKeySystemAccess.createMediaKeys".
func (this MediaKeySystemAccess) FuncCreateMediaKeys() (fn js.Func[func() js.Promise[MediaKeys]]) {
	bindings.FuncMediaKeySystemAccessCreateMediaKeys(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateMediaKeys calls the method "MediaKeySystemAccess.createMediaKeys".
func (this MediaKeySystemAccess) CreateMediaKeys() (ret js.Promise[MediaKeys]) {
	bindings.CallMediaKeySystemAccessCreateMediaKeys(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateMediaKeys calls the method "MediaKeySystemAccess.createMediaKeys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaKeySystemAccess) TryCreateMediaKeys() (ret js.Promise[MediaKeys], exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaKeySystemAccessCreateMediaKeys(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *MediaStreamConstraints) UpdateFrom(ref js.Ref) {
	bindings.MediaStreamConstraintsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *MediaStreamConstraints) Update(ref js.Ref) {
	bindings.MediaStreamConstraintsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *MediaStreamConstraints) FreeMembers(recursive bool) {
	js.Free(
		p.Video.Ref(),
		p.Audio.Ref(),
		p.PeerIdentity.Ref(),
	)
	p.Video = p.Video.FromRef(js.Undefined)
	p.Audio = p.Audio.FromRef(js.Undefined)
	p.PeerIdentity = p.PeerIdentity.FromRef(js.Undefined)
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
		js.ThrowInvalidCallbackInvocation()
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
		js.ThrowInvalidCallbackInvocation()
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
	this.ref.Once()
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
	this.ref.Free()
}

// ContainerWidth returns the value of property "FencedFrameConfig.containerWidth".
//
// It returns ok=false if there is no such property.
func (this FencedFrameConfig) ContainerWidth() (ret FencedFrameConfigSize, ok bool) {
	ok = js.True == bindings.GetFencedFrameConfigContainerWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ContainerHeight returns the value of property "FencedFrameConfig.containerHeight".
//
// It returns ok=false if there is no such property.
func (this FencedFrameConfig) ContainerHeight() (ret FencedFrameConfigSize, ok bool) {
	ok = js.True == bindings.GetFencedFrameConfigContainerHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ContentWidth returns the value of property "FencedFrameConfig.contentWidth".
//
// It returns ok=false if there is no such property.
func (this FencedFrameConfig) ContentWidth() (ret FencedFrameConfigSize, ok bool) {
	ok = js.True == bindings.GetFencedFrameConfigContentWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ContentHeight returns the value of property "FencedFrameConfig.contentHeight".
//
// It returns ok=false if there is no such property.
func (this FencedFrameConfig) ContentHeight() (ret FencedFrameConfigSize, ok bool) {
	ok = js.True == bindings.GetFencedFrameConfigContentHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncSetSharedStorageContext returns true if the method "FencedFrameConfig.setSharedStorageContext" exists.
func (this FencedFrameConfig) HasFuncSetSharedStorageContext() bool {
	return js.True == bindings.HasFuncFencedFrameConfigSetSharedStorageContext(
		this.ref,
	)
}

// FuncSetSharedStorageContext returns the method "FencedFrameConfig.setSharedStorageContext".
func (this FencedFrameConfig) FuncSetSharedStorageContext() (fn js.Func[func(contextString js.String)]) {
	bindings.FuncFencedFrameConfigSetSharedStorageContext(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetSharedStorageContext calls the method "FencedFrameConfig.setSharedStorageContext".
func (this FencedFrameConfig) SetSharedStorageContext(contextString js.String) (ret js.Void) {
	bindings.CallFencedFrameConfigSetSharedStorageContext(
		this.ref, js.Pointer(&ret),
		contextString.Ref(),
	)

	return
}

// TrySetSharedStorageContext calls the method "FencedFrameConfig.setSharedStorageContext"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FencedFrameConfig) TrySetSharedStorageContext(contextString js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFencedFrameConfigSetSharedStorageContext(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *ShareData) UpdateFrom(ref js.Ref) {
	bindings.ShareDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ShareData) Update(ref js.Ref) {
	bindings.ShareDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ShareData) FreeMembers(recursive bool) {
	js.Free(
		p.Files.Ref(),
		p.Title.Ref(),
		p.Text.Ref(),
		p.Url.Ref(),
	)
	p.Files = p.Files.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Text = p.Text.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type MIDIInputMap struct {
	ref js.Ref
}

func (this MIDIInputMap) Once() MIDIInputMap {
	this.ref.Once()
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
	this.ref.Free()
}

type MIDIOutputMap struct {
	ref js.Ref
}

func (this MIDIOutputMap) Once() MIDIOutputMap {
	this.ref.Once()
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
	this.ref.Free()
}

type MIDIAccess struct {
	EventTarget
}

func (this MIDIAccess) Once() MIDIAccess {
	this.ref.Once()
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
	this.ref.Free()
}

// Inputs returns the value of property "MIDIAccess.inputs".
//
// It returns ok=false if there is no such property.
func (this MIDIAccess) Inputs() (ret MIDIInputMap, ok bool) {
	ok = js.True == bindings.GetMIDIAccessInputs(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Outputs returns the value of property "MIDIAccess.outputs".
//
// It returns ok=false if there is no such property.
func (this MIDIAccess) Outputs() (ret MIDIOutputMap, ok bool) {
	ok = js.True == bindings.GetMIDIAccessOutputs(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SysexEnabled returns the value of property "MIDIAccess.sysexEnabled".
//
// It returns ok=false if there is no such property.
func (this MIDIAccess) SysexEnabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetMIDIAccessSysexEnabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

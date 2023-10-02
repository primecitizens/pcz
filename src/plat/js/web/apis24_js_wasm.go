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

// New creates a new {0x140004cc0e0 NavigationNavigateOptions NavigationNavigateOptions [// NavigationNavigateOptions] [0x1400069cc80 0x1400069cd20 0x1400069cdc0] 0x14000575470 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 NavigationReloadOptions NavigationReloadOptions [// NavigationReloadOptions] [0x1400069ce60 0x1400069cf00] 0x140005754b8 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 NavigationOptions NavigationOptions [// NavigationOptions] [0x1400069cfa0] 0x14000575518 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this NavigationTransition) NavigationType() (NavigationType, bool) {
	var _ok bool
	_ret := bindings.GetNavigationTransitionNavigationType(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigationType(_ret), _ok
}

// From returns the value of property "NavigationTransition.from".
//
// The returned bool will be false if there is no such property.
func (this NavigationTransition) From() (NavigationHistoryEntry, bool) {
	var _ok bool
	_ret := bindings.GetNavigationTransitionFrom(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigationHistoryEntry{}.FromRef(_ret), _ok
}

// Finished returns the value of property "NavigationTransition.finished".
//
// The returned bool will be false if there is no such property.
func (this NavigationTransition) Finished() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.GetNavigationTransitionFinished(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Promise[js.Void]{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this Navigation) CurrentEntry() (NavigationHistoryEntry, bool) {
	var _ok bool
	_ret := bindings.GetNavigationCurrentEntry(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigationHistoryEntry{}.FromRef(_ret), _ok
}

// Transition returns the value of property "Navigation.transition".
//
// The returned bool will be false if there is no such property.
func (this Navigation) Transition() (NavigationTransition, bool) {
	var _ok bool
	_ret := bindings.GetNavigationTransition(
		this.Ref(), js.Pointer(&_ok),
	)
	return NavigationTransition{}.FromRef(_ret), _ok
}

// CanGoBack returns the value of property "Navigation.canGoBack".
//
// The returned bool will be false if there is no such property.
func (this Navigation) CanGoBack() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigationCanGoBack(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// CanGoForward returns the value of property "Navigation.canGoForward".
//
// The returned bool will be false if there is no such property.
func (this Navigation) CanGoForward() (bool, bool) {
	var _ok bool
	_ret := bindings.GetNavigationCanGoForward(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Entries calls the method "Navigation.entries".
//
// The returned bool will be false if there is no such method.
func (this Navigation) Entries() (js.Array[NavigationHistoryEntry], bool) {
	var _ok bool
	_ret := bindings.CallNavigationEntries(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Array[NavigationHistoryEntry]{}.FromRef(_ret), _ok
}

// EntriesFunc returns the method "Navigation.entries".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) EntriesFunc() (fn js.Func[func() js.Array[NavigationHistoryEntry]]) {
	return fn.FromRef(
		bindings.NavigationEntriesFunc(
			this.Ref(),
		),
	)
}

// UpdateCurrentEntry calls the method "Navigation.updateCurrentEntry".
//
// The returned bool will be false if there is no such method.
func (this Navigation) UpdateCurrentEntry(options NavigationUpdateCurrentEntryOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallNavigationUpdateCurrentEntry(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpdateCurrentEntryFunc returns the method "Navigation.updateCurrentEntry".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) UpdateCurrentEntryFunc() (fn js.Func[func(options NavigationUpdateCurrentEntryOptions)]) {
	return fn.FromRef(
		bindings.NavigationUpdateCurrentEntryFunc(
			this.Ref(),
		),
	)
}

// Navigate calls the method "Navigation.navigate".
//
// The returned bool will be false if there is no such method.
func (this Navigation) Navigate(url js.String, options NavigationNavigateOptions) (NavigationResult, bool) {
	var _ret NavigationResult
	_ok := js.True == bindings.CallNavigationNavigate(
		this.Ref(), js.Pointer(&_ret),
		url.Ref(),
		js.Pointer(&options),
	)

	return _ret, _ok
}

// NavigateFunc returns the method "Navigation.navigate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) NavigateFunc() (fn js.Func[func(url js.String, options NavigationNavigateOptions) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationNavigateFunc(
			this.Ref(),
		),
	)
}

// Navigate1 calls the method "Navigation.navigate".
//
// The returned bool will be false if there is no such method.
func (this Navigation) Navigate1(url js.String) (NavigationResult, bool) {
	var _ret NavigationResult
	_ok := js.True == bindings.CallNavigationNavigate1(
		this.Ref(), js.Pointer(&_ret),
		url.Ref(),
	)

	return _ret, _ok
}

// Navigate1Func returns the method "Navigation.navigate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) Navigate1Func() (fn js.Func[func(url js.String) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationNavigate1Func(
			this.Ref(),
		),
	)
}

// Reload calls the method "Navigation.reload".
//
// The returned bool will be false if there is no such method.
func (this Navigation) Reload(options NavigationReloadOptions) (NavigationResult, bool) {
	var _ret NavigationResult
	_ok := js.True == bindings.CallNavigationReload(
		this.Ref(), js.Pointer(&_ret),
		js.Pointer(&options),
	)

	return _ret, _ok
}

// ReloadFunc returns the method "Navigation.reload".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) ReloadFunc() (fn js.Func[func(options NavigationReloadOptions) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationReloadFunc(
			this.Ref(),
		),
	)
}

// Reload1 calls the method "Navigation.reload".
//
// The returned bool will be false if there is no such method.
func (this Navigation) Reload1() (NavigationResult, bool) {
	var _ret NavigationResult
	_ok := js.True == bindings.CallNavigationReload1(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// Reload1Func returns the method "Navigation.reload".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) Reload1Func() (fn js.Func[func() NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationReload1Func(
			this.Ref(),
		),
	)
}

// TraverseTo calls the method "Navigation.traverseTo".
//
// The returned bool will be false if there is no such method.
func (this Navigation) TraverseTo(key js.String, options NavigationOptions) (NavigationResult, bool) {
	var _ret NavigationResult
	_ok := js.True == bindings.CallNavigationTraverseTo(
		this.Ref(), js.Pointer(&_ret),
		key.Ref(),
		js.Pointer(&options),
	)

	return _ret, _ok
}

// TraverseToFunc returns the method "Navigation.traverseTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) TraverseToFunc() (fn js.Func[func(key js.String, options NavigationOptions) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationTraverseToFunc(
			this.Ref(),
		),
	)
}

// TraverseTo1 calls the method "Navigation.traverseTo".
//
// The returned bool will be false if there is no such method.
func (this Navigation) TraverseTo1(key js.String) (NavigationResult, bool) {
	var _ret NavigationResult
	_ok := js.True == bindings.CallNavigationTraverseTo1(
		this.Ref(), js.Pointer(&_ret),
		key.Ref(),
	)

	return _ret, _ok
}

// TraverseTo1Func returns the method "Navigation.traverseTo".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) TraverseTo1Func() (fn js.Func[func(key js.String) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationTraverseTo1Func(
			this.Ref(),
		),
	)
}

// Back calls the method "Navigation.back".
//
// The returned bool will be false if there is no such method.
func (this Navigation) Back(options NavigationOptions) (NavigationResult, bool) {
	var _ret NavigationResult
	_ok := js.True == bindings.CallNavigationBack(
		this.Ref(), js.Pointer(&_ret),
		js.Pointer(&options),
	)

	return _ret, _ok
}

// BackFunc returns the method "Navigation.back".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) BackFunc() (fn js.Func[func(options NavigationOptions) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationBackFunc(
			this.Ref(),
		),
	)
}

// Back1 calls the method "Navigation.back".
//
// The returned bool will be false if there is no such method.
func (this Navigation) Back1() (NavigationResult, bool) {
	var _ret NavigationResult
	_ok := js.True == bindings.CallNavigationBack1(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// Back1Func returns the method "Navigation.back".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) Back1Func() (fn js.Func[func() NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationBack1Func(
			this.Ref(),
		),
	)
}

// Forward calls the method "Navigation.forward".
//
// The returned bool will be false if there is no such method.
func (this Navigation) Forward(options NavigationOptions) (NavigationResult, bool) {
	var _ret NavigationResult
	_ok := js.True == bindings.CallNavigationForward(
		this.Ref(), js.Pointer(&_ret),
		js.Pointer(&options),
	)

	return _ret, _ok
}

// ForwardFunc returns the method "Navigation.forward".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) ForwardFunc() (fn js.Func[func(options NavigationOptions) NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationForwardFunc(
			this.Ref(),
		),
	)
}

// Forward1 calls the method "Navigation.forward".
//
// The returned bool will be false if there is no such method.
func (this Navigation) Forward1() (NavigationResult, bool) {
	var _ret NavigationResult
	_ok := js.True == bindings.CallNavigationForward1(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// Forward1Func returns the method "Navigation.forward".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this Navigation) Forward1Func() (fn js.Func[func() NavigationResult]) {
	return fn.FromRef(
		bindings.NavigationForward1Func(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 ElementDefinitionOptions ElementDefinitionOptions [// ElementDefinitionOptions] [0x1400069d180] 0x140005755a8 {0 0}} in the application heap.
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

// Define calls the method "CustomElementRegistry.define".
//
// The returned bool will be false if there is no such method.
func (this CustomElementRegistry) Define(name js.String, constructor js.Func[func() HTMLElement], options ElementDefinitionOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCustomElementRegistryDefine(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		constructor.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DefineFunc returns the method "CustomElementRegistry.define".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CustomElementRegistry) DefineFunc() (fn js.Func[func(name js.String, constructor js.Func[func() HTMLElement], options ElementDefinitionOptions)]) {
	return fn.FromRef(
		bindings.CustomElementRegistryDefineFunc(
			this.Ref(),
		),
	)
}

// Define1 calls the method "CustomElementRegistry.define".
//
// The returned bool will be false if there is no such method.
func (this CustomElementRegistry) Define1(name js.String, constructor js.Func[func() HTMLElement]) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCustomElementRegistryDefine1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		constructor.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Define1Func returns the method "CustomElementRegistry.define".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CustomElementRegistry) Define1Func() (fn js.Func[func(name js.String, constructor js.Func[func() HTMLElement])]) {
	return fn.FromRef(
		bindings.CustomElementRegistryDefine1Func(
			this.Ref(),
		),
	)
}

// Get calls the method "CustomElementRegistry.get".
//
// The returned bool will be false if there is no such method.
func (this CustomElementRegistry) Get(name js.String) (OneOf_FuncCustomElementConstructor_undefined, bool) {
	var _ok bool
	_ret := bindings.CallCustomElementRegistryGet(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return OneOf_FuncCustomElementConstructor_undefined{}.FromRef(_ret), _ok
}

// GetFunc returns the method "CustomElementRegistry.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CustomElementRegistry) GetFunc() (fn js.Func[func(name js.String) OneOf_FuncCustomElementConstructor_undefined]) {
	return fn.FromRef(
		bindings.CustomElementRegistryGetFunc(
			this.Ref(),
		),
	)
}

// GetName calls the method "CustomElementRegistry.getName".
//
// The returned bool will be false if there is no such method.
func (this CustomElementRegistry) GetName(constructor js.Func[func() HTMLElement]) (js.String, bool) {
	var _ok bool
	_ret := bindings.CallCustomElementRegistryGetName(
		this.Ref(), js.Pointer(&_ok),
		constructor.Ref(),
	)

	return js.String{}.FromRef(_ret), _ok
}

// GetNameFunc returns the method "CustomElementRegistry.getName".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CustomElementRegistry) GetNameFunc() (fn js.Func[func(constructor js.Func[func() HTMLElement]) js.String]) {
	return fn.FromRef(
		bindings.CustomElementRegistryGetNameFunc(
			this.Ref(),
		),
	)
}

// WhenDefined calls the method "CustomElementRegistry.whenDefined".
//
// The returned bool will be false if there is no such method.
func (this CustomElementRegistry) WhenDefined(name js.String) (js.Promise[js.Func[func() HTMLElement]], bool) {
	var _ok bool
	_ret := bindings.CallCustomElementRegistryWhenDefined(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return js.Promise[js.Func[func() HTMLElement]]{}.FromRef(_ret), _ok
}

// WhenDefinedFunc returns the method "CustomElementRegistry.whenDefined".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CustomElementRegistry) WhenDefinedFunc() (fn js.Func[func(name js.String) js.Promise[js.Func[func() HTMLElement]]]) {
	return fn.FromRef(
		bindings.CustomElementRegistryWhenDefinedFunc(
			this.Ref(),
		),
	)
}

// Upgrade calls the method "CustomElementRegistry.upgrade".
//
// The returned bool will be false if there is no such method.
func (this CustomElementRegistry) Upgrade(root Node) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallCustomElementRegistryUpgrade(
		this.Ref(), js.Pointer(&_ok),
		root.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// UpgradeFunc returns the method "CustomElementRegistry.upgrade".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this CustomElementRegistry) UpgradeFunc() (fn js.Func[func(root Node)]) {
	return fn.FromRef(
		bindings.CustomElementRegistryUpgradeFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this GamepadButton) Pressed() (bool, bool) {
	var _ok bool
	_ret := bindings.GetGamepadButtonPressed(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Touched returns the value of property "GamepadButton.touched".
//
// The returned bool will be false if there is no such property.
func (this GamepadButton) Touched() (bool, bool) {
	var _ok bool
	_ret := bindings.GetGamepadButtonTouched(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Value returns the value of property "GamepadButton.value".
//
// The returned bool will be false if there is no such property.
func (this GamepadButton) Value() (float64, bool) {
	var _ok bool
	_ret := bindings.GetGamepadButtonValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
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
	Duration float64
	// StartDelay is "GamepadEffectParameters.startDelay"
	//
	// Optional, defaults to 0.0.
	StartDelay float64
	// StrongMagnitude is "GamepadEffectParameters.strongMagnitude"
	//
	// Optional, defaults to 0.0.
	StrongMagnitude float64
	// WeakMagnitude is "GamepadEffectParameters.weakMagnitude"
	//
	// Optional, defaults to 0.0.
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

// New creates a new {0x140004cc0e0 GamepadEffectParameters GamepadEffectParameters [// GamepadEffectParameters] [0x1400069d360 0x1400069d4a0 0x1400069d5e0 0x1400069d720 0x1400069d400 0x1400069d540 0x1400069d680 0x1400069d7c0] 0x14000575620 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this GamepadHapticActuator) Type() (GamepadHapticActuatorType, bool) {
	var _ok bool
	_ret := bindings.GetGamepadHapticActuatorType(
		this.Ref(), js.Pointer(&_ok),
	)
	return GamepadHapticActuatorType(_ret), _ok
}

// CanPlayEffectType calls the method "GamepadHapticActuator.canPlayEffectType".
//
// The returned bool will be false if there is no such method.
func (this GamepadHapticActuator) CanPlayEffectType(typ GamepadHapticEffectType) (bool, bool) {
	var _ok bool
	_ret := bindings.CallGamepadHapticActuatorCanPlayEffectType(
		this.Ref(), js.Pointer(&_ok),
		uint32(typ),
	)

	return _ret == js.True, _ok
}

// CanPlayEffectTypeFunc returns the method "GamepadHapticActuator.canPlayEffectType".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GamepadHapticActuator) CanPlayEffectTypeFunc() (fn js.Func[func(typ GamepadHapticEffectType) bool]) {
	return fn.FromRef(
		bindings.GamepadHapticActuatorCanPlayEffectTypeFunc(
			this.Ref(),
		),
	)
}

// PlayEffect calls the method "GamepadHapticActuator.playEffect".
//
// The returned bool will be false if there is no such method.
func (this GamepadHapticActuator) PlayEffect(typ GamepadHapticEffectType, params GamepadEffectParameters) (js.Promise[GamepadHapticsResult], bool) {
	var _ok bool
	_ret := bindings.CallGamepadHapticActuatorPlayEffect(
		this.Ref(), js.Pointer(&_ok),
		uint32(typ),
		js.Pointer(&params),
	)

	return js.Promise[GamepadHapticsResult]{}.FromRef(_ret), _ok
}

// PlayEffectFunc returns the method "GamepadHapticActuator.playEffect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GamepadHapticActuator) PlayEffectFunc() (fn js.Func[func(typ GamepadHapticEffectType, params GamepadEffectParameters) js.Promise[GamepadHapticsResult]]) {
	return fn.FromRef(
		bindings.GamepadHapticActuatorPlayEffectFunc(
			this.Ref(),
		),
	)
}

// PlayEffect1 calls the method "GamepadHapticActuator.playEffect".
//
// The returned bool will be false if there is no such method.
func (this GamepadHapticActuator) PlayEffect1(typ GamepadHapticEffectType) (js.Promise[GamepadHapticsResult], bool) {
	var _ok bool
	_ret := bindings.CallGamepadHapticActuatorPlayEffect1(
		this.Ref(), js.Pointer(&_ok),
		uint32(typ),
	)

	return js.Promise[GamepadHapticsResult]{}.FromRef(_ret), _ok
}

// PlayEffect1Func returns the method "GamepadHapticActuator.playEffect".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GamepadHapticActuator) PlayEffect1Func() (fn js.Func[func(typ GamepadHapticEffectType) js.Promise[GamepadHapticsResult]]) {
	return fn.FromRef(
		bindings.GamepadHapticActuatorPlayEffect1Func(
			this.Ref(),
		),
	)
}

// Pulse calls the method "GamepadHapticActuator.pulse".
//
// The returned bool will be false if there is no such method.
func (this GamepadHapticActuator) Pulse(value float64, duration float64) (js.Promise[js.Boolean], bool) {
	var _ok bool
	_ret := bindings.CallGamepadHapticActuatorPulse(
		this.Ref(), js.Pointer(&_ok),
		float64(value),
		float64(duration),
	)

	return js.Promise[js.Boolean]{}.FromRef(_ret), _ok
}

// PulseFunc returns the method "GamepadHapticActuator.pulse".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GamepadHapticActuator) PulseFunc() (fn js.Func[func(value float64, duration float64) js.Promise[js.Boolean]]) {
	return fn.FromRef(
		bindings.GamepadHapticActuatorPulseFunc(
			this.Ref(),
		),
	)
}

// Reset calls the method "GamepadHapticActuator.reset".
//
// The returned bool will be false if there is no such method.
func (this GamepadHapticActuator) Reset() (js.Promise[GamepadHapticsResult], bool) {
	var _ok bool
	_ret := bindings.CallGamepadHapticActuatorReset(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[GamepadHapticsResult]{}.FromRef(_ret), _ok
}

// ResetFunc returns the method "GamepadHapticActuator.reset".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this GamepadHapticActuator) ResetFunc() (fn js.Func[func() js.Promise[GamepadHapticsResult]]) {
	return fn.FromRef(
		bindings.GamepadHapticActuatorResetFunc(
			this.Ref(),
		),
	)
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
// The returned bool will be false if there is no such property.
func (this GamepadPose) HasOrientation() (bool, bool) {
	var _ok bool
	_ret := bindings.GetGamepadPoseHasOrientation(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// HasPosition returns the value of property "GamepadPose.hasPosition".
//
// The returned bool will be false if there is no such property.
func (this GamepadPose) HasPosition() (bool, bool) {
	var _ok bool
	_ret := bindings.GetGamepadPoseHasPosition(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Position returns the value of property "GamepadPose.position".
//
// The returned bool will be false if there is no such property.
func (this GamepadPose) Position() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetGamepadPosePosition(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// LinearVelocity returns the value of property "GamepadPose.linearVelocity".
//
// The returned bool will be false if there is no such property.
func (this GamepadPose) LinearVelocity() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetGamepadPoseLinearVelocity(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// LinearAcceleration returns the value of property "GamepadPose.linearAcceleration".
//
// The returned bool will be false if there is no such property.
func (this GamepadPose) LinearAcceleration() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetGamepadPoseLinearAcceleration(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// Orientation returns the value of property "GamepadPose.orientation".
//
// The returned bool will be false if there is no such property.
func (this GamepadPose) Orientation() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetGamepadPoseOrientation(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// AngularVelocity returns the value of property "GamepadPose.angularVelocity".
//
// The returned bool will be false if there is no such property.
func (this GamepadPose) AngularVelocity() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetGamepadPoseAngularVelocity(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// AngularAcceleration returns the value of property "GamepadPose.angularAcceleration".
//
// The returned bool will be false if there is no such property.
func (this GamepadPose) AngularAcceleration() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetGamepadPoseAngularAcceleration(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this GamepadTouch) TouchId() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetGamepadTouchTouchId(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// SurfaceId returns the value of property "GamepadTouch.surfaceId".
//
// The returned bool will be false if there is no such property.
func (this GamepadTouch) SurfaceId() (uint8, bool) {
	var _ok bool
	_ret := bindings.GetGamepadTouchSurfaceId(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint8(_ret), _ok
}

// Position returns the value of property "GamepadTouch.position".
//
// The returned bool will be false if there is no such property.
func (this GamepadTouch) Position() (js.TypedArray[float32], bool) {
	var _ok bool
	_ret := bindings.GetGamepadTouchPosition(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[float32]{}.FromRef(_ret), _ok
}

// SurfaceDimensions returns the value of property "GamepadTouch.surfaceDimensions".
//
// The returned bool will be false if there is no such property.
func (this GamepadTouch) SurfaceDimensions() (js.TypedArray[uint32], bool) {
	var _ok bool
	_ret := bindings.GetGamepadTouchSurfaceDimensions(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.TypedArray[uint32]{}.FromRef(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this Gamepad) Id() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetGamepadId(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Index returns the value of property "Gamepad.index".
//
// The returned bool will be false if there is no such property.
func (this Gamepad) Index() (int32, bool) {
	var _ok bool
	_ret := bindings.GetGamepadIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Connected returns the value of property "Gamepad.connected".
//
// The returned bool will be false if there is no such property.
func (this Gamepad) Connected() (bool, bool) {
	var _ok bool
	_ret := bindings.GetGamepadConnected(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Timestamp returns the value of property "Gamepad.timestamp".
//
// The returned bool will be false if there is no such property.
func (this Gamepad) Timestamp() (DOMHighResTimeStamp, bool) {
	var _ok bool
	_ret := bindings.GetGamepadTimestamp(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMHighResTimeStamp(_ret), _ok
}

// Mapping returns the value of property "Gamepad.mapping".
//
// The returned bool will be false if there is no such property.
func (this Gamepad) Mapping() (GamepadMappingType, bool) {
	var _ok bool
	_ret := bindings.GetGamepadMapping(
		this.Ref(), js.Pointer(&_ok),
	)
	return GamepadMappingType(_ret), _ok
}

// Axes returns the value of property "Gamepad.axes".
//
// The returned bool will be false if there is no such property.
func (this Gamepad) Axes() (js.FrozenArray[float64], bool) {
	var _ok bool
	_ret := bindings.GetGamepadAxes(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[float64]{}.FromRef(_ret), _ok
}

// Buttons returns the value of property "Gamepad.buttons".
//
// The returned bool will be false if there is no such property.
func (this Gamepad) Buttons() (js.FrozenArray[GamepadButton], bool) {
	var _ok bool
	_ret := bindings.GetGamepadButtons(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[GamepadButton]{}.FromRef(_ret), _ok
}

// Hand returns the value of property "Gamepad.hand".
//
// The returned bool will be false if there is no such property.
func (this Gamepad) Hand() (GamepadHand, bool) {
	var _ok bool
	_ret := bindings.GetGamepadHand(
		this.Ref(), js.Pointer(&_ok),
	)
	return GamepadHand(_ret), _ok
}

// HapticActuators returns the value of property "Gamepad.hapticActuators".
//
// The returned bool will be false if there is no such property.
func (this Gamepad) HapticActuators() (js.FrozenArray[GamepadHapticActuator], bool) {
	var _ok bool
	_ret := bindings.GetGamepadHapticActuators(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[GamepadHapticActuator]{}.FromRef(_ret), _ok
}

// Pose returns the value of property "Gamepad.pose".
//
// The returned bool will be false if there is no such property.
func (this Gamepad) Pose() (GamepadPose, bool) {
	var _ok bool
	_ret := bindings.GetGamepadPose(
		this.Ref(), js.Pointer(&_ok),
	)
	return GamepadPose{}.FromRef(_ret), _ok
}

// TouchEvents returns the value of property "Gamepad.touchEvents".
//
// The returned bool will be false if there is no such property.
func (this Gamepad) TouchEvents() (js.FrozenArray[GamepadTouch], bool) {
	var _ok bool
	_ret := bindings.GetGamepadTouchEvents(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[GamepadTouch]{}.FromRef(_ret), _ok
}

// VibrationActuator returns the value of property "Gamepad.vibrationActuator".
//
// The returned bool will be false if there is no such property.
func (this Gamepad) VibrationActuator() (GamepadHapticActuator, bool) {
	var _ok bool
	_ret := bindings.GetGamepadVibrationActuator(
		this.Ref(), js.Pointer(&_ok),
	)
	return GamepadHapticActuator{}.FromRef(_ret), _ok
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

// New creates a new {0x140004cc0e0 RelatedApplication RelatedApplication [// RelatedApplication] [0x1400069d9a0 0x1400069da40 0x1400069dae0 0x1400069db80] 0x14000575698 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MediaKeySystemMediaCapability MediaKeySystemMediaCapability [// MediaKeySystemMediaCapability] [0x1400069dd60 0x1400069de00 0x1400069dea0] 0x140005756f8 {0 0}} in the application heap.
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

// New creates a new {0x140004cc0e0 MediaKeySystemConfiguration MediaKeySystemConfiguration [// MediaKeySystemConfiguration] [0x1400069dc20 0x1400069dcc0 0x1400069df40 0x140006ce000 0x140006ce0a0 0x140006ce140 0x140006ce1e0] 0x140005756e0 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this MediaKeySystemAccess) KeySystem() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetMediaKeySystemAccessKeySystem(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// GetConfiguration calls the method "MediaKeySystemAccess.getConfiguration".
//
// The returned bool will be false if there is no such method.
func (this MediaKeySystemAccess) GetConfiguration() (MediaKeySystemConfiguration, bool) {
	var _ret MediaKeySystemConfiguration
	_ok := js.True == bindings.CallMediaKeySystemAccessGetConfiguration(
		this.Ref(), js.Pointer(&_ret),
	)

	return _ret, _ok
}

// GetConfigurationFunc returns the method "MediaKeySystemAccess.getConfiguration".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeySystemAccess) GetConfigurationFunc() (fn js.Func[func() MediaKeySystemConfiguration]) {
	return fn.FromRef(
		bindings.MediaKeySystemAccessGetConfigurationFunc(
			this.Ref(),
		),
	)
}

// CreateMediaKeys calls the method "MediaKeySystemAccess.createMediaKeys".
//
// The returned bool will be false if there is no such method.
func (this MediaKeySystemAccess) CreateMediaKeys() (js.Promise[MediaKeys], bool) {
	var _ok bool
	_ret := bindings.CallMediaKeySystemAccessCreateMediaKeys(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[MediaKeys]{}.FromRef(_ret), _ok
}

// CreateMediaKeysFunc returns the method "MediaKeySystemAccess.createMediaKeys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this MediaKeySystemAccess) CreateMediaKeysFunc() (fn js.Func[func() js.Promise[MediaKeys]]) {
	return fn.FromRef(
		bindings.MediaKeySystemAccessCreateMediaKeysFunc(
			this.Ref(),
		),
	)
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
	PreferCurrentTab bool

	FFI_USE_PreferCurrentTab bool // for PreferCurrentTab.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a MediaStreamConstraints with all fields set.
func (p MediaStreamConstraints) FromRef(ref js.Ref) MediaStreamConstraints {
	p.UpdateFrom(ref)
	return p
}

// New creates a new {0x140004cc0e0 MediaStreamConstraints MediaStreamConstraints [// MediaStreamConstraints] [0x140006ce280 0x140006ce320 0x140006ce3c0 0x140006ce460 0x140006ce500] 0x14000575710 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this FencedFrameConfig) ContainerWidth() (FencedFrameConfigSize, bool) {
	var _ok bool
	_ret := bindings.GetFencedFrameConfigContainerWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return FencedFrameConfigSize{}.FromRef(_ret), _ok
}

// ContainerHeight returns the value of property "FencedFrameConfig.containerHeight".
//
// The returned bool will be false if there is no such property.
func (this FencedFrameConfig) ContainerHeight() (FencedFrameConfigSize, bool) {
	var _ok bool
	_ret := bindings.GetFencedFrameConfigContainerHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return FencedFrameConfigSize{}.FromRef(_ret), _ok
}

// ContentWidth returns the value of property "FencedFrameConfig.contentWidth".
//
// The returned bool will be false if there is no such property.
func (this FencedFrameConfig) ContentWidth() (FencedFrameConfigSize, bool) {
	var _ok bool
	_ret := bindings.GetFencedFrameConfigContentWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return FencedFrameConfigSize{}.FromRef(_ret), _ok
}

// ContentHeight returns the value of property "FencedFrameConfig.contentHeight".
//
// The returned bool will be false if there is no such property.
func (this FencedFrameConfig) ContentHeight() (FencedFrameConfigSize, bool) {
	var _ok bool
	_ret := bindings.GetFencedFrameConfigContentHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return FencedFrameConfigSize{}.FromRef(_ret), _ok
}

// SetSharedStorageContext calls the method "FencedFrameConfig.setSharedStorageContext".
//
// The returned bool will be false if there is no such method.
func (this FencedFrameConfig) SetSharedStorageContext(contextString js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallFencedFrameConfigSetSharedStorageContext(
		this.Ref(), js.Pointer(&_ok),
		contextString.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetSharedStorageContextFunc returns the method "FencedFrameConfig.setSharedStorageContext".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this FencedFrameConfig) SetSharedStorageContextFunc() (fn js.Func[func(contextString js.String)]) {
	return fn.FromRef(
		bindings.FencedFrameConfigSetSharedStorageContextFunc(
			this.Ref(),
		),
	)
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

// New creates a new {0x140004cc0e0 ShareData ShareData [// ShareData] [0x140006ce640 0x140006ce6e0 0x140006ce780 0x140006ce820] 0x140005757a0 {0 0}} in the application heap.
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
// The returned bool will be false if there is no such property.
func (this MIDIAccess) Inputs() (MIDIInputMap, bool) {
	var _ok bool
	_ret := bindings.GetMIDIAccessInputs(
		this.Ref(), js.Pointer(&_ok),
	)
	return MIDIInputMap{}.FromRef(_ret), _ok
}

// Outputs returns the value of property "MIDIAccess.outputs".
//
// The returned bool will be false if there is no such property.
func (this MIDIAccess) Outputs() (MIDIOutputMap, bool) {
	var _ok bool
	_ret := bindings.GetMIDIAccessOutputs(
		this.Ref(), js.Pointer(&_ok),
	)
	return MIDIOutputMap{}.FromRef(_ret), _ok
}

// SysexEnabled returns the value of property "MIDIAccess.sysexEnabled".
//
// The returned bool will be false if there is no such property.
func (this MIDIAccess) SysexEnabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetMIDIAccessSysexEnabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

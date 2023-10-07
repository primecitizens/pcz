// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package events

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/events/bindings"
)

type AddListenerArgCallbackFunc func(this js.Ref) js.Ref

func (fn AddListenerArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AddListenerArgCallbackFunc) DispatchCallback(
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

type AddListenerArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *AddListenerArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AddListenerArgCallback[T]) DispatchCallback(
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

type AddRulesArgCallbackFunc func(this js.Ref, rules js.Array[Rule]) js.Ref

func (fn AddRulesArgCallbackFunc) Register() js.Func[func(rules js.Array[Rule])] {
	return js.RegisterCallback[func(rules js.Array[Rule])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn AddRulesArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Rule]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type AddRulesArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, rules js.Array[Rule]) js.Ref
	Arg T
}

func (cb *AddRulesArgCallback[T]) Register() js.Func[func(rules js.Array[Rule])] {
	return js.RegisterCallback[func(rules js.Array[Rule])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *AddRulesArgCallback[T]) DispatchCallback(
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

		js.Array[Rule]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type Rule struct {
	// Actions is "Rule.actions"
	//
	// Required
	Actions js.Array[js.Any]
	// Conditions is "Rule.conditions"
	//
	// Required
	Conditions js.Array[js.Any]
	// Id is "Rule.id"
	//
	// Optional
	Id js.String
	// Priority is "Rule.priority"
	//
	// Optional
	//
	// NOTE: FFI_USE_Priority MUST be set to true to make this field effective.
	Priority int64
	// Tags is "Rule.tags"
	//
	// Optional
	Tags js.Array[js.String]

	FFI_USE_Priority bool // for Priority.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a Rule with all fields set.
func (p Rule) FromRef(ref js.Ref) Rule {
	p.UpdateFrom(ref)
	return p
}

// New creates a new Rule in the application heap.
func (p Rule) New() js.Ref {
	return bindings.RuleJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *Rule) UpdateFrom(ref js.Ref) {
	bindings.RuleJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *Rule) Update(ref js.Ref) {
	bindings.RuleJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *Rule) FreeMembers(recursive bool) {
	js.Free(
		p.Actions.Ref(),
		p.Conditions.Ref(),
		p.Id.Ref(),
		p.Tags.Ref(),
	)
	p.Actions = p.Actions.FromRef(js.Undefined)
	p.Conditions = p.Conditions.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Tags = p.Tags.FromRef(js.Undefined)
}

type RemoveListenerArgCallbackFunc func(this js.Ref) js.Ref

func (fn RemoveListenerArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RemoveListenerArgCallbackFunc) DispatchCallback(
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

type RemoveListenerArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RemoveListenerArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RemoveListenerArgCallback[T]) DispatchCallback(
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

type HasListenerArgCallbackFunc func(this js.Ref) js.Ref

func (fn HasListenerArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn HasListenerArgCallbackFunc) DispatchCallback(
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

type HasListenerArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *HasListenerArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *HasListenerArgCallback[T]) DispatchCallback(
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

type GetRulesArgCallbackFunc func(this js.Ref, rules js.Array[Rule]) js.Ref

func (fn GetRulesArgCallbackFunc) Register() js.Func[func(rules js.Array[Rule])] {
	return js.RegisterCallback[func(rules js.Array[Rule])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetRulesArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[Rule]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetRulesArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, rules js.Array[Rule]) js.Ref
	Arg T
}

func (cb *GetRulesArgCallback[T]) Register() js.Func[func(rules js.Array[Rule])] {
	return js.RegisterCallback[func(rules js.Array[Rule])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetRulesArgCallback[T]) DispatchCallback(
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

		js.Array[Rule]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RemoveRulesArgCallbackFunc func(this js.Ref) js.Ref

func (fn RemoveRulesArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RemoveRulesArgCallbackFunc) DispatchCallback(
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

type RemoveRulesArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RemoveRulesArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RemoveRulesArgCallback[T]) DispatchCallback(
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

// HasFuncAddListener returns true if the method "Event.addListener" exists.
func (this Event) HasFuncAddListener() bool {
	return js.True == bindings.HasFuncEventAddListener(
		this.ref,
	)
}

// FuncAddListener returns the method "Event.addListener".
func (this Event) FuncAddListener() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncEventAddListener(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddListener calls the method "Event.addListener".
func (this Event) AddListener(callback js.Func[func()]) (ret js.Void) {
	bindings.CallEventAddListener(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryAddListener calls the method "Event.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryAddListener(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventAddListener(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRemoveListener returns true if the method "Event.removeListener" exists.
func (this Event) HasFuncRemoveListener() bool {
	return js.True == bindings.HasFuncEventRemoveListener(
		this.ref,
	)
}

// FuncRemoveListener returns the method "Event.removeListener".
func (this Event) FuncRemoveListener() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncEventRemoveListener(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveListener calls the method "Event.removeListener".
func (this Event) RemoveListener(callback js.Func[func()]) (ret js.Void) {
	bindings.CallEventRemoveListener(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRemoveListener calls the method "Event.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryRemoveListener(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventRemoveListener(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasListener returns true if the method "Event.hasListener" exists.
func (this Event) HasFuncHasListener() bool {
	return js.True == bindings.HasFuncEventHasListener(
		this.ref,
	)
}

// FuncHasListener returns the method "Event.hasListener".
func (this Event) FuncHasListener() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncEventHasListener(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasListener calls the method "Event.hasListener".
func (this Event) HasListener(callback js.Func[func()]) (ret bool) {
	bindings.CallEventHasListener(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasListener calls the method "Event.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryHasListener(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventHasListener(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasListeners returns true if the method "Event.hasListeners" exists.
func (this Event) HasFuncHasListeners() bool {
	return js.True == bindings.HasFuncEventHasListeners(
		this.ref,
	)
}

// FuncHasListeners returns the method "Event.hasListeners".
func (this Event) FuncHasListeners() (fn js.Func[func() bool]) {
	bindings.FuncEventHasListeners(
		this.ref, js.Pointer(&fn),
	)
	return
}

// HasListeners calls the method "Event.hasListeners".
func (this Event) HasListeners() (ret bool) {
	bindings.CallEventHasListeners(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryHasListeners calls the method "Event.hasListeners"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryHasListeners() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventHasListeners(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAddRules returns true if the method "Event.addRules" exists.
func (this Event) HasFuncAddRules() bool {
	return js.True == bindings.HasFuncEventAddRules(
		this.ref,
	)
}

// FuncAddRules returns the method "Event.addRules".
func (this Event) FuncAddRules() (fn js.Func[func(eventName js.String, webViewInstanceId int64, rules js.Array[Rule], callback js.Func[func(rules js.Array[Rule])])]) {
	bindings.FuncEventAddRules(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddRules calls the method "Event.addRules".
func (this Event) AddRules(eventName js.String, webViewInstanceId int64, rules js.Array[Rule], callback js.Func[func(rules js.Array[Rule])]) (ret js.Void) {
	bindings.CallEventAddRules(
		this.ref, js.Pointer(&ret),
		eventName.Ref(),
		float64(webViewInstanceId),
		rules.Ref(),
		callback.Ref(),
	)

	return
}

// TryAddRules calls the method "Event.addRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryAddRules(eventName js.String, webViewInstanceId int64, rules js.Array[Rule], callback js.Func[func(rules js.Array[Rule])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventAddRules(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		eventName.Ref(),
		float64(webViewInstanceId),
		rules.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncAddRules1 returns true if the method "Event.addRules" exists.
func (this Event) HasFuncAddRules1() bool {
	return js.True == bindings.HasFuncEventAddRules1(
		this.ref,
	)
}

// FuncAddRules1 returns the method "Event.addRules".
func (this Event) FuncAddRules1() (fn js.Func[func(eventName js.String, webViewInstanceId int64, rules js.Array[Rule])]) {
	bindings.FuncEventAddRules1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// AddRules1 calls the method "Event.addRules".
func (this Event) AddRules1(eventName js.String, webViewInstanceId int64, rules js.Array[Rule]) (ret js.Void) {
	bindings.CallEventAddRules1(
		this.ref, js.Pointer(&ret),
		eventName.Ref(),
		float64(webViewInstanceId),
		rules.Ref(),
	)

	return
}

// TryAddRules1 calls the method "Event.addRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryAddRules1(eventName js.String, webViewInstanceId int64, rules js.Array[Rule]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventAddRules1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		eventName.Ref(),
		float64(webViewInstanceId),
		rules.Ref(),
	)

	return
}

// HasFuncGetRules returns true if the method "Event.getRules" exists.
func (this Event) HasFuncGetRules() bool {
	return js.True == bindings.HasFuncEventGetRules(
		this.ref,
	)
}

// FuncGetRules returns the method "Event.getRules".
func (this Event) FuncGetRules() (fn js.Func[func(eventName js.String, webViewInstanceId int64, ruleIdentifiers js.Array[js.String], callback js.Func[func(rules js.Array[Rule])])]) {
	bindings.FuncEventGetRules(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRules calls the method "Event.getRules".
func (this Event) GetRules(eventName js.String, webViewInstanceId int64, ruleIdentifiers js.Array[js.String], callback js.Func[func(rules js.Array[Rule])]) (ret js.Void) {
	bindings.CallEventGetRules(
		this.ref, js.Pointer(&ret),
		eventName.Ref(),
		float64(webViewInstanceId),
		ruleIdentifiers.Ref(),
		callback.Ref(),
	)

	return
}

// TryGetRules calls the method "Event.getRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryGetRules(eventName js.String, webViewInstanceId int64, ruleIdentifiers js.Array[js.String], callback js.Func[func(rules js.Array[Rule])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventGetRules(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		eventName.Ref(),
		float64(webViewInstanceId),
		ruleIdentifiers.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncGetRules1 returns true if the method "Event.getRules" exists.
func (this Event) HasFuncGetRules1() bool {
	return js.True == bindings.HasFuncEventGetRules1(
		this.ref,
	)
}

// FuncGetRules1 returns the method "Event.getRules".
func (this Event) FuncGetRules1() (fn js.Func[func(eventName js.String, webViewInstanceId int64)]) {
	bindings.FuncEventGetRules1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetRules1 calls the method "Event.getRules".
func (this Event) GetRules1(eventName js.String, webViewInstanceId int64) (ret js.Void) {
	bindings.CallEventGetRules1(
		this.ref, js.Pointer(&ret),
		eventName.Ref(),
		float64(webViewInstanceId),
	)

	return
}

// TryGetRules1 calls the method "Event.getRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryGetRules1(eventName js.String, webViewInstanceId int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventGetRules1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		eventName.Ref(),
		float64(webViewInstanceId),
	)

	return
}

// HasFuncRemoveRules returns true if the method "Event.removeRules" exists.
func (this Event) HasFuncRemoveRules() bool {
	return js.True == bindings.HasFuncEventRemoveRules(
		this.ref,
	)
}

// FuncRemoveRules returns the method "Event.removeRules".
func (this Event) FuncRemoveRules() (fn js.Func[func(eventName js.String, webViewInstanceId int64, ruleIdentifiers js.Array[js.String], callback js.Func[func()])]) {
	bindings.FuncEventRemoveRules(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveRules calls the method "Event.removeRules".
func (this Event) RemoveRules(eventName js.String, webViewInstanceId int64, ruleIdentifiers js.Array[js.String], callback js.Func[func()]) (ret js.Void) {
	bindings.CallEventRemoveRules(
		this.ref, js.Pointer(&ret),
		eventName.Ref(),
		float64(webViewInstanceId),
		ruleIdentifiers.Ref(),
		callback.Ref(),
	)

	return
}

// TryRemoveRules calls the method "Event.removeRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryRemoveRules(eventName js.String, webViewInstanceId int64, ruleIdentifiers js.Array[js.String], callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventRemoveRules(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		eventName.Ref(),
		float64(webViewInstanceId),
		ruleIdentifiers.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncRemoveRules1 returns true if the method "Event.removeRules" exists.
func (this Event) HasFuncRemoveRules1() bool {
	return js.True == bindings.HasFuncEventRemoveRules1(
		this.ref,
	)
}

// FuncRemoveRules1 returns the method "Event.removeRules".
func (this Event) FuncRemoveRules1() (fn js.Func[func(eventName js.String, webViewInstanceId int64, ruleIdentifiers js.Array[js.String])]) {
	bindings.FuncEventRemoveRules1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveRules1 calls the method "Event.removeRules".
func (this Event) RemoveRules1(eventName js.String, webViewInstanceId int64, ruleIdentifiers js.Array[js.String]) (ret js.Void) {
	bindings.CallEventRemoveRules1(
		this.ref, js.Pointer(&ret),
		eventName.Ref(),
		float64(webViewInstanceId),
		ruleIdentifiers.Ref(),
	)

	return
}

// TryRemoveRules1 calls the method "Event.removeRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryRemoveRules1(eventName js.String, webViewInstanceId int64, ruleIdentifiers js.Array[js.String]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventRemoveRules1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		eventName.Ref(),
		float64(webViewInstanceId),
		ruleIdentifiers.Ref(),
	)

	return
}

// HasFuncRemoveRules2 returns true if the method "Event.removeRules" exists.
func (this Event) HasFuncRemoveRules2() bool {
	return js.True == bindings.HasFuncEventRemoveRules2(
		this.ref,
	)
}

// FuncRemoveRules2 returns the method "Event.removeRules".
func (this Event) FuncRemoveRules2() (fn js.Func[func(eventName js.String, webViewInstanceId int64)]) {
	bindings.FuncEventRemoveRules2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveRules2 calls the method "Event.removeRules".
func (this Event) RemoveRules2(eventName js.String, webViewInstanceId int64) (ret js.Void) {
	bindings.CallEventRemoveRules2(
		this.ref, js.Pointer(&ret),
		eventName.Ref(),
		float64(webViewInstanceId),
	)

	return
}

// TryRemoveRules2 calls the method "Event.removeRules"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Event) TryRemoveRules2(eventName js.String, webViewInstanceId int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryEventRemoveRules2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		eventName.Ref(),
		float64(webViewInstanceId),
	)

	return
}

type OneOf_Int64_ArrayInt64 struct {
	ref js.Ref
}

func (x OneOf_Int64_ArrayInt64) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Int64_ArrayInt64) Free() {
	x.ref.Free()
}

func (x OneOf_Int64_ArrayInt64) FromRef(ref js.Ref) OneOf_Int64_ArrayInt64 {
	return OneOf_Int64_ArrayInt64{
		ref: ref,
	}
}

func (x OneOf_Int64_ArrayInt64) Int64() int64 {
	return js.BigInt[int64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Int64_ArrayInt64) ArrayInt64() js.Array[int64] {
	return js.Array[int64]{}.FromRef(x.ref)
}

type UrlFilter struct {
	// HostContains is "UrlFilter.hostContains"
	//
	// Optional
	HostContains js.String
	// HostEquals is "UrlFilter.hostEquals"
	//
	// Optional
	HostEquals js.String
	// HostPrefix is "UrlFilter.hostPrefix"
	//
	// Optional
	HostPrefix js.String
	// HostSuffix is "UrlFilter.hostSuffix"
	//
	// Optional
	HostSuffix js.String
	// OriginAndPathMatches is "UrlFilter.originAndPathMatches"
	//
	// Optional
	OriginAndPathMatches js.String
	// PathContains is "UrlFilter.pathContains"
	//
	// Optional
	PathContains js.String
	// PathEquals is "UrlFilter.pathEquals"
	//
	// Optional
	PathEquals js.String
	// PathPrefix is "UrlFilter.pathPrefix"
	//
	// Optional
	PathPrefix js.String
	// PathSuffix is "UrlFilter.pathSuffix"
	//
	// Optional
	PathSuffix js.String
	// Ports is "UrlFilter.ports"
	//
	// Optional
	Ports js.Array[OneOf_Int64_ArrayInt64]
	// QueryContains is "UrlFilter.queryContains"
	//
	// Optional
	QueryContains js.String
	// QueryEquals is "UrlFilter.queryEquals"
	//
	// Optional
	QueryEquals js.String
	// QueryPrefix is "UrlFilter.queryPrefix"
	//
	// Optional
	QueryPrefix js.String
	// QuerySuffix is "UrlFilter.querySuffix"
	//
	// Optional
	QuerySuffix js.String
	// Schemes is "UrlFilter.schemes"
	//
	// Optional
	Schemes js.Array[js.String]
	// UrlContains is "UrlFilter.urlContains"
	//
	// Optional
	UrlContains js.String
	// UrlEquals is "UrlFilter.urlEquals"
	//
	// Optional
	UrlEquals js.String
	// UrlMatches is "UrlFilter.urlMatches"
	//
	// Optional
	UrlMatches js.String
	// UrlPrefix is "UrlFilter.urlPrefix"
	//
	// Optional
	UrlPrefix js.String
	// UrlSuffix is "UrlFilter.urlSuffix"
	//
	// Optional
	UrlSuffix js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UrlFilter with all fields set.
func (p UrlFilter) FromRef(ref js.Ref) UrlFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UrlFilter in the application heap.
func (p UrlFilter) New() js.Ref {
	return bindings.UrlFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UrlFilter) UpdateFrom(ref js.Ref) {
	bindings.UrlFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UrlFilter) Update(ref js.Ref) {
	bindings.UrlFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UrlFilter) FreeMembers(recursive bool) {
	js.Free(
		p.HostContains.Ref(),
		p.HostEquals.Ref(),
		p.HostPrefix.Ref(),
		p.HostSuffix.Ref(),
		p.OriginAndPathMatches.Ref(),
		p.PathContains.Ref(),
		p.PathEquals.Ref(),
		p.PathPrefix.Ref(),
		p.PathSuffix.Ref(),
		p.Ports.Ref(),
		p.QueryContains.Ref(),
		p.QueryEquals.Ref(),
		p.QueryPrefix.Ref(),
		p.QuerySuffix.Ref(),
		p.Schemes.Ref(),
		p.UrlContains.Ref(),
		p.UrlEquals.Ref(),
		p.UrlMatches.Ref(),
		p.UrlPrefix.Ref(),
		p.UrlSuffix.Ref(),
	)
	p.HostContains = p.HostContains.FromRef(js.Undefined)
	p.HostEquals = p.HostEquals.FromRef(js.Undefined)
	p.HostPrefix = p.HostPrefix.FromRef(js.Undefined)
	p.HostSuffix = p.HostSuffix.FromRef(js.Undefined)
	p.OriginAndPathMatches = p.OriginAndPathMatches.FromRef(js.Undefined)
	p.PathContains = p.PathContains.FromRef(js.Undefined)
	p.PathEquals = p.PathEquals.FromRef(js.Undefined)
	p.PathPrefix = p.PathPrefix.FromRef(js.Undefined)
	p.PathSuffix = p.PathSuffix.FromRef(js.Undefined)
	p.Ports = p.Ports.FromRef(js.Undefined)
	p.QueryContains = p.QueryContains.FromRef(js.Undefined)
	p.QueryEquals = p.QueryEquals.FromRef(js.Undefined)
	p.QueryPrefix = p.QueryPrefix.FromRef(js.Undefined)
	p.QuerySuffix = p.QuerySuffix.FromRef(js.Undefined)
	p.Schemes = p.Schemes.FromRef(js.Undefined)
	p.UrlContains = p.UrlContains.FromRef(js.Undefined)
	p.UrlEquals = p.UrlEquals.FromRef(js.Undefined)
	p.UrlMatches = p.UrlMatches.FromRef(js.Undefined)
	p.UrlPrefix = p.UrlPrefix.FromRef(js.Undefined)
	p.UrlSuffix = p.UrlSuffix.FromRef(js.Undefined)
}

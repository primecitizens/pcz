// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package userscripts

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensiontypes"
	"github.com/primecitizens/pcz/std/plat/js/webext/userscripts/bindings"
)

type GetScriptsCallbackFunc func(this js.Ref, scripts js.Array[RegisteredUserScript]) js.Ref

func (fn GetScriptsCallbackFunc) Register() js.Func[func(scripts js.Array[RegisteredUserScript])] {
	return js.RegisterCallback[func(scripts js.Array[RegisteredUserScript])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetScriptsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[RegisteredUserScript]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetScriptsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, scripts js.Array[RegisteredUserScript]) js.Ref
	Arg T
}

func (cb *GetScriptsCallback[T]) Register() js.Func[func(scripts js.Array[RegisteredUserScript])] {
	return js.RegisterCallback[func(scripts js.Array[RegisteredUserScript])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetScriptsCallback[T]) DispatchCallback(
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

		js.Array[RegisteredUserScript]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ScriptSource struct {
	// Code is "ScriptSource.code"
	//
	// Optional
	Code js.String
	// File is "ScriptSource.file"
	//
	// Optional
	File js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScriptSource with all fields set.
func (p ScriptSource) FromRef(ref js.Ref) ScriptSource {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScriptSource in the application heap.
func (p ScriptSource) New() js.Ref {
	return bindings.ScriptSourceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScriptSource) UpdateFrom(ref js.Ref) {
	bindings.ScriptSourceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScriptSource) Update(ref js.Ref) {
	bindings.ScriptSourceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScriptSource) FreeMembers(recursive bool) {
	js.Free(
		p.Code.Ref(),
		p.File.Ref(),
	)
	p.Code = p.Code.FromRef(js.Undefined)
	p.File = p.File.FromRef(js.Undefined)
}

type RegisteredUserScript struct {
	// AllFrames is "RegisteredUserScript.allFrames"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllFrames MUST be set to true to make this field effective.
	AllFrames bool
	// ExcludeMatches is "RegisteredUserScript.excludeMatches"
	//
	// Optional
	ExcludeMatches js.Array[js.String]
	// Id is "RegisteredUserScript.id"
	//
	// Optional
	Id js.String
	// Js is "RegisteredUserScript.js"
	//
	// Optional
	Js js.Array[ScriptSource]
	// Matches is "RegisteredUserScript.matches"
	//
	// Optional
	Matches js.Array[js.String]
	// RunAt is "RegisteredUserScript.runAt"
	//
	// Optional
	RunAt extensiontypes.RunAt

	FFI_USE_AllFrames bool // for AllFrames.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RegisteredUserScript with all fields set.
func (p RegisteredUserScript) FromRef(ref js.Ref) RegisteredUserScript {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RegisteredUserScript in the application heap.
func (p RegisteredUserScript) New() js.Ref {
	return bindings.RegisteredUserScriptJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RegisteredUserScript) UpdateFrom(ref js.Ref) {
	bindings.RegisteredUserScriptJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RegisteredUserScript) Update(ref js.Ref) {
	bindings.RegisteredUserScriptJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RegisteredUserScript) FreeMembers(recursive bool) {
	js.Free(
		p.ExcludeMatches.Ref(),
		p.Id.Ref(),
		p.Js.Ref(),
		p.Matches.Ref(),
	)
	p.ExcludeMatches = p.ExcludeMatches.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Js = p.Js.FromRef(js.Undefined)
	p.Matches = p.Matches.FromRef(js.Undefined)
}

type RegisterCallbackFunc func(this js.Ref) js.Ref

func (fn RegisterCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RegisterCallbackFunc) DispatchCallback(
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

type RegisterCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RegisterCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RegisterCallback[T]) DispatchCallback(
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

type UnregisterCallbackFunc func(this js.Ref) js.Ref

func (fn UnregisterCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UnregisterCallbackFunc) DispatchCallback(
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

type UnregisterCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *UnregisterCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UnregisterCallback[T]) DispatchCallback(
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

type UserScriptFilter struct {
	// Ids is "UserScriptFilter.ids"
	//
	// Optional
	Ids js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UserScriptFilter with all fields set.
func (p UserScriptFilter) FromRef(ref js.Ref) UserScriptFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UserScriptFilter in the application heap.
func (p UserScriptFilter) New() js.Ref {
	return bindings.UserScriptFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UserScriptFilter) UpdateFrom(ref js.Ref) {
	bindings.UserScriptFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UserScriptFilter) Update(ref js.Ref) {
	bindings.UserScriptFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UserScriptFilter) FreeMembers(recursive bool) {
	js.Free(
		p.Ids.Ref(),
	)
	p.Ids = p.Ids.FromRef(js.Undefined)
}

// HasFuncGetScripts returns true if the function "WEBEXT.userScripts.getScripts" exists.
func HasFuncGetScripts() bool {
	return js.True == bindings.HasFuncGetScripts()
}

// FuncGetScripts returns the function "WEBEXT.userScripts.getScripts".
func FuncGetScripts() (fn js.Func[func(filter UserScriptFilter) js.Promise[js.Array[RegisteredUserScript]]]) {
	bindings.FuncGetScripts(
		js.Pointer(&fn),
	)
	return
}

// GetScripts calls the function "WEBEXT.userScripts.getScripts" directly.
func GetScripts(filter UserScriptFilter) (ret js.Promise[js.Array[RegisteredUserScript]]) {
	bindings.CallGetScripts(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetScripts calls the function "WEBEXT.userScripts.getScripts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetScripts(filter UserScriptFilter) (ret js.Promise[js.Array[RegisteredUserScript]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetScripts(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

// HasFuncRegister returns true if the function "WEBEXT.userScripts.register" exists.
func HasFuncRegister() bool {
	return js.True == bindings.HasFuncRegister()
}

// FuncRegister returns the function "WEBEXT.userScripts.register".
func FuncRegister() (fn js.Func[func(scripts js.Array[RegisteredUserScript]) js.Promise[js.Void]]) {
	bindings.FuncRegister(
		js.Pointer(&fn),
	)
	return
}

// Register calls the function "WEBEXT.userScripts.register" directly.
func Register(scripts js.Array[RegisteredUserScript]) (ret js.Promise[js.Void]) {
	bindings.CallRegister(
		js.Pointer(&ret),
		scripts.Ref(),
	)

	return
}

// TryRegister calls the function "WEBEXT.userScripts.register"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRegister(scripts js.Array[RegisteredUserScript]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRegister(
		js.Pointer(&ret), js.Pointer(&exception),
		scripts.Ref(),
	)

	return
}

// HasFuncUnregister returns true if the function "WEBEXT.userScripts.unregister" exists.
func HasFuncUnregister() bool {
	return js.True == bindings.HasFuncUnregister()
}

// FuncUnregister returns the function "WEBEXT.userScripts.unregister".
func FuncUnregister() (fn js.Func[func(filter UserScriptFilter) js.Promise[js.Void]]) {
	bindings.FuncUnregister(
		js.Pointer(&fn),
	)
	return
}

// Unregister calls the function "WEBEXT.userScripts.unregister" directly.
func Unregister(filter UserScriptFilter) (ret js.Promise[js.Void]) {
	bindings.CallUnregister(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryUnregister calls the function "WEBEXT.userScripts.unregister"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnregister(filter UserScriptFilter) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnregister(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

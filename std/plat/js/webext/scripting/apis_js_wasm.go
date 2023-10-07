// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package scripting

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensiontypes"
	"github.com/primecitizens/pcz/std/plat/js/webext/scripting/bindings"
)

type InjectionTarget struct {
	// TabId is "InjectionTarget.tabId"
	//
	// Optional
	//
	// NOTE: FFI_USE_TabId MUST be set to true to make this field effective.
	TabId int32
	// FrameIds is "InjectionTarget.frameIds"
	//
	// Optional
	FrameIds js.Array[int32]
	// DocumentIds is "InjectionTarget.documentIds"
	//
	// Optional
	DocumentIds js.Array[js.String]
	// AllFrames is "InjectionTarget.allFrames"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllFrames MUST be set to true to make this field effective.
	AllFrames bool

	FFI_USE_TabId     bool // for TabId.
	FFI_USE_AllFrames bool // for AllFrames.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InjectionTarget with all fields set.
func (p InjectionTarget) FromRef(ref js.Ref) InjectionTarget {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InjectionTarget in the application heap.
func (p InjectionTarget) New() js.Ref {
	return bindings.InjectionTargetJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InjectionTarget) UpdateFrom(ref js.Ref) {
	bindings.InjectionTargetJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InjectionTarget) Update(ref js.Ref) {
	bindings.InjectionTargetJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InjectionTarget) FreeMembers(recursive bool) {
	js.Free(
		p.FrameIds.Ref(),
		p.DocumentIds.Ref(),
	)
	p.FrameIds = p.FrameIds.FromRef(js.Undefined)
	p.DocumentIds = p.DocumentIds.FromRef(js.Undefined)
}

type StyleOrigin uint32

const (
	_ StyleOrigin = iota

	StyleOrigin_AUTHOR
	StyleOrigin_USER
)

func (StyleOrigin) FromRef(str js.Ref) StyleOrigin {
	return StyleOrigin(bindings.ConstOfStyleOrigin(str))
}

func (x StyleOrigin) String() (string, bool) {
	switch x {
	case StyleOrigin_AUTHOR:
		return "AUTHOR", true
	case StyleOrigin_USER:
		return "USER", true
	default:
		return "", false
	}
}

type CSSInjection struct {
	// Target is "CSSInjection.target"
	//
	// Optional
	//
	// NOTE: Target.FFI_USE MUST be set to true to get Target used.
	Target InjectionTarget
	// Css is "CSSInjection.css"
	//
	// Optional
	Css js.String
	// Files is "CSSInjection.files"
	//
	// Optional
	Files js.Array[js.String]
	// Origin is "CSSInjection.origin"
	//
	// Optional
	Origin StyleOrigin

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a CSSInjection with all fields set.
func (p CSSInjection) FromRef(ref js.Ref) CSSInjection {
	p.UpdateFrom(ref)
	return p
}

// New creates a new CSSInjection in the application heap.
func (p CSSInjection) New() js.Ref {
	return bindings.CSSInjectionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *CSSInjection) UpdateFrom(ref js.Ref) {
	bindings.CSSInjectionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *CSSInjection) Update(ref js.Ref) {
	bindings.CSSInjectionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *CSSInjection) FreeMembers(recursive bool) {
	js.Free(
		p.Css.Ref(),
		p.Files.Ref(),
	)
	p.Css = p.Css.FromRef(js.Undefined)
	p.Files = p.Files.FromRef(js.Undefined)
	if recursive {
		p.Target.FreeMembers(true)
	}
}

type CSSInjectionCallbackFunc func(this js.Ref) js.Ref

func (fn CSSInjectionCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CSSInjectionCallbackFunc) DispatchCallback(
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

type CSSInjectionCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CSSInjectionCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CSSInjectionCallback[T]) DispatchCallback(
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

type ContentScriptFilter struct {
	// Ids is "ContentScriptFilter.ids"
	//
	// Optional
	Ids js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContentScriptFilter with all fields set.
func (p ContentScriptFilter) FromRef(ref js.Ref) ContentScriptFilter {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContentScriptFilter in the application heap.
func (p ContentScriptFilter) New() js.Ref {
	return bindings.ContentScriptFilterJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContentScriptFilter) UpdateFrom(ref js.Ref) {
	bindings.ContentScriptFilterJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContentScriptFilter) Update(ref js.Ref) {
	bindings.ContentScriptFilterJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContentScriptFilter) FreeMembers(recursive bool) {
	js.Free(
		p.Ids.Ref(),
	)
	p.Ids = p.Ids.FromRef(js.Undefined)
}

type ExecutionWorld uint32

const (
	_ ExecutionWorld = iota

	ExecutionWorld_ISOLATED
	ExecutionWorld_MAIN
)

func (ExecutionWorld) FromRef(str js.Ref) ExecutionWorld {
	return ExecutionWorld(bindings.ConstOfExecutionWorld(str))
}

func (x ExecutionWorld) String() (string, bool) {
	switch x {
	case ExecutionWorld_ISOLATED:
		return "ISOLATED", true
	case ExecutionWorld_MAIN:
		return "MAIN", true
	default:
		return "", false
	}
}

type GetRegisteredContentScriptsCallbackFunc func(this js.Ref, scripts js.Array[RegisteredContentScript]) js.Ref

func (fn GetRegisteredContentScriptsCallbackFunc) Register() js.Func[func(scripts js.Array[RegisteredContentScript])] {
	return js.RegisterCallback[func(scripts js.Array[RegisteredContentScript])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetRegisteredContentScriptsCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[RegisteredContentScript]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetRegisteredContentScriptsCallback[T any] struct {
	Fn  func(arg T, this js.Ref, scripts js.Array[RegisteredContentScript]) js.Ref
	Arg T
}

func (cb *GetRegisteredContentScriptsCallback[T]) Register() js.Func[func(scripts js.Array[RegisteredContentScript])] {
	return js.RegisterCallback[func(scripts js.Array[RegisteredContentScript])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetRegisteredContentScriptsCallback[T]) DispatchCallback(
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

		js.Array[RegisteredContentScript]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type RegisteredContentScript struct {
	// Id is "RegisteredContentScript.id"
	//
	// Optional
	Id js.String
	// Matches is "RegisteredContentScript.matches"
	//
	// Optional
	Matches js.Array[js.String]
	// ExcludeMatches is "RegisteredContentScript.excludeMatches"
	//
	// Optional
	ExcludeMatches js.Array[js.String]
	// Css is "RegisteredContentScript.css"
	//
	// Optional
	Css js.Array[js.String]
	// Js is "RegisteredContentScript.js"
	//
	// Optional
	Js js.Array[js.String]
	// AllFrames is "RegisteredContentScript.allFrames"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllFrames MUST be set to true to make this field effective.
	AllFrames bool
	// MatchOriginAsFallback is "RegisteredContentScript.matchOriginAsFallback"
	//
	// Optional
	//
	// NOTE: FFI_USE_MatchOriginAsFallback MUST be set to true to make this field effective.
	MatchOriginAsFallback bool
	// RunAt is "RegisteredContentScript.runAt"
	//
	// Optional
	RunAt extensiontypes.RunAt
	// PersistAcrossSessions is "RegisteredContentScript.persistAcrossSessions"
	//
	// Optional
	//
	// NOTE: FFI_USE_PersistAcrossSessions MUST be set to true to make this field effective.
	PersistAcrossSessions bool
	// World is "RegisteredContentScript.world"
	//
	// Optional
	World ExecutionWorld

	FFI_USE_AllFrames             bool // for AllFrames.
	FFI_USE_MatchOriginAsFallback bool // for MatchOriginAsFallback.
	FFI_USE_PersistAcrossSessions bool // for PersistAcrossSessions.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RegisteredContentScript with all fields set.
func (p RegisteredContentScript) FromRef(ref js.Ref) RegisteredContentScript {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RegisteredContentScript in the application heap.
func (p RegisteredContentScript) New() js.Ref {
	return bindings.RegisteredContentScriptJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RegisteredContentScript) UpdateFrom(ref js.Ref) {
	bindings.RegisteredContentScriptJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RegisteredContentScript) Update(ref js.Ref) {
	bindings.RegisteredContentScriptJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RegisteredContentScript) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Matches.Ref(),
		p.ExcludeMatches.Ref(),
		p.Css.Ref(),
		p.Js.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Matches = p.Matches.FromRef(js.Undefined)
	p.ExcludeMatches = p.ExcludeMatches.FromRef(js.Undefined)
	p.Css = p.Css.FromRef(js.Undefined)
	p.Js = p.Js.FromRef(js.Undefined)
}

type InjectedFunctionFunc func(this js.Ref) js.Ref

func (fn InjectedFunctionFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn InjectedFunctionFunc) DispatchCallback(
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

type InjectedFunction[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *InjectedFunction[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *InjectedFunction[T]) DispatchCallback(
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

type InjectionResult struct {
	// Result is "InjectionResult.result"
	//
	// Optional
	Result js.Any
	// FrameId is "InjectionResult.frameId"
	//
	// Optional
	//
	// NOTE: FFI_USE_FrameId MUST be set to true to make this field effective.
	FrameId int32
	// DocumentId is "InjectionResult.documentId"
	//
	// Optional
	DocumentId js.String

	FFI_USE_FrameId bool // for FrameId.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InjectionResult with all fields set.
func (p InjectionResult) FromRef(ref js.Ref) InjectionResult {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InjectionResult in the application heap.
func (p InjectionResult) New() js.Ref {
	return bindings.InjectionResultJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InjectionResult) UpdateFrom(ref js.Ref) {
	bindings.InjectionResultJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InjectionResult) Update(ref js.Ref) {
	bindings.InjectionResultJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InjectionResult) FreeMembers(recursive bool) {
	js.Free(
		p.Result.Ref(),
		p.DocumentId.Ref(),
	)
	p.Result = p.Result.FromRef(js.Undefined)
	p.DocumentId = p.DocumentId.FromRef(js.Undefined)
}

type Properties struct {
	ref js.Ref
}

func (this Properties) Once() Properties {
	this.ref.Once()
	return this
}

func (this Properties) Ref() js.Ref {
	return this.ref
}

func (this Properties) FromRef(ref js.Ref) Properties {
	this.ref = ref
	return this
}

func (this Properties) Free() {
	this.ref.Free()
}

// HasFuncGlobalParams returns true if the static method "Properties.globalParams" exists.
func (this Properties) HasFuncGlobalParams() bool {
	return js.True == bindings.HasFuncPropertiesGlobalParams(
		this.ref,
	)
}

// FuncGlobalParams returns the static method "Properties.globalParams".
func (this Properties) FuncGlobalParams() (fn js.Func[func() int32]) {
	bindings.FuncPropertiesGlobalParams(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GlobalParams calls the static method "Properties.globalParams".
func (this Properties) GlobalParams() (ret int32) {
	bindings.CallPropertiesGlobalParams(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGlobalParams calls the static method "Properties.globalParams"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this Properties) TryGlobalParams() (ret int32, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPropertiesGlobalParams(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type RegisterContentScriptsCallbackFunc func(this js.Ref) js.Ref

func (fn RegisterContentScriptsCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RegisterContentScriptsCallbackFunc) DispatchCallback(
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

type RegisterContentScriptsCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RegisterContentScriptsCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RegisterContentScriptsCallback[T]) DispatchCallback(
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

type ScriptInjection struct {
	// Func is "ScriptInjection.func"
	//
	// Optional
	Func js.Func[func()]
	// Args is "ScriptInjection.args"
	//
	// Optional
	Args js.Array[js.Any]
	// Function is "ScriptInjection.function"
	//
	// Optional
	Function js.Func[func()]
	// Files is "ScriptInjection.files"
	//
	// Optional
	Files js.Array[js.String]
	// Target is "ScriptInjection.target"
	//
	// Optional
	//
	// NOTE: Target.FFI_USE MUST be set to true to get Target used.
	Target InjectionTarget
	// World is "ScriptInjection.world"
	//
	// Optional
	World ExecutionWorld
	// InjectImmediately is "ScriptInjection.injectImmediately"
	//
	// Optional
	//
	// NOTE: FFI_USE_InjectImmediately MUST be set to true to make this field effective.
	InjectImmediately bool

	FFI_USE_InjectImmediately bool // for InjectImmediately.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ScriptInjection with all fields set.
func (p ScriptInjection) FromRef(ref js.Ref) ScriptInjection {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ScriptInjection in the application heap.
func (p ScriptInjection) New() js.Ref {
	return bindings.ScriptInjectionJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ScriptInjection) UpdateFrom(ref js.Ref) {
	bindings.ScriptInjectionJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ScriptInjection) Update(ref js.Ref) {
	bindings.ScriptInjectionJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ScriptInjection) FreeMembers(recursive bool) {
	js.Free(
		p.Func.Ref(),
		p.Args.Ref(),
		p.Function.Ref(),
		p.Files.Ref(),
	)
	p.Func = p.Func.FromRef(js.Undefined)
	p.Args = p.Args.FromRef(js.Undefined)
	p.Function = p.Function.FromRef(js.Undefined)
	p.Files = p.Files.FromRef(js.Undefined)
	if recursive {
		p.Target.FreeMembers(true)
	}
}

type ScriptInjectionCallbackFunc func(this js.Ref, results js.Array[InjectionResult]) js.Ref

func (fn ScriptInjectionCallbackFunc) Register() js.Func[func(results js.Array[InjectionResult])] {
	return js.RegisterCallback[func(results js.Array[InjectionResult])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ScriptInjectionCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[InjectionResult]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ScriptInjectionCallback[T any] struct {
	Fn  func(arg T, this js.Ref, results js.Array[InjectionResult]) js.Ref
	Arg T
}

func (cb *ScriptInjectionCallback[T]) Register() js.Func[func(results js.Array[InjectionResult])] {
	return js.RegisterCallback[func(results js.Array[InjectionResult])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ScriptInjectionCallback[T]) DispatchCallback(
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

		js.Array[InjectionResult]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnregisterContentScriptsCallbackFunc func(this js.Ref) js.Ref

func (fn UnregisterContentScriptsCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UnregisterContentScriptsCallbackFunc) DispatchCallback(
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

type UnregisterContentScriptsCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *UnregisterContentScriptsCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UnregisterContentScriptsCallback[T]) DispatchCallback(
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

type UpdateContentScriptsCallbackFunc func(this js.Ref) js.Ref

func (fn UpdateContentScriptsCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UpdateContentScriptsCallbackFunc) DispatchCallback(
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

type UpdateContentScriptsCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *UpdateContentScriptsCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UpdateContentScriptsCallback[T]) DispatchCallback(
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

// HasFuncExecuteScript returns true if the function "WEBEXT.scripting.executeScript" exists.
func HasFuncExecuteScript() bool {
	return js.True == bindings.HasFuncExecuteScript()
}

// FuncExecuteScript returns the function "WEBEXT.scripting.executeScript".
func FuncExecuteScript() (fn js.Func[func(injection ScriptInjection) js.Promise[js.Array[InjectionResult]]]) {
	bindings.FuncExecuteScript(
		js.Pointer(&fn),
	)
	return
}

// ExecuteScript calls the function "WEBEXT.scripting.executeScript" directly.
func ExecuteScript(injection ScriptInjection) (ret js.Promise[js.Array[InjectionResult]]) {
	bindings.CallExecuteScript(
		js.Pointer(&ret),
		js.Pointer(&injection),
	)

	return
}

// TryExecuteScript calls the function "WEBEXT.scripting.executeScript"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExecuteScript(injection ScriptInjection) (ret js.Promise[js.Array[InjectionResult]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryExecuteScript(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&injection),
	)

	return
}

// HasFuncGetRegisteredContentScripts returns true if the function "WEBEXT.scripting.getRegisteredContentScripts" exists.
func HasFuncGetRegisteredContentScripts() bool {
	return js.True == bindings.HasFuncGetRegisteredContentScripts()
}

// FuncGetRegisteredContentScripts returns the function "WEBEXT.scripting.getRegisteredContentScripts".
func FuncGetRegisteredContentScripts() (fn js.Func[func(filter ContentScriptFilter) js.Promise[js.Array[RegisteredContentScript]]]) {
	bindings.FuncGetRegisteredContentScripts(
		js.Pointer(&fn),
	)
	return
}

// GetRegisteredContentScripts calls the function "WEBEXT.scripting.getRegisteredContentScripts" directly.
func GetRegisteredContentScripts(filter ContentScriptFilter) (ret js.Promise[js.Array[RegisteredContentScript]]) {
	bindings.CallGetRegisteredContentScripts(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryGetRegisteredContentScripts calls the function "WEBEXT.scripting.getRegisteredContentScripts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetRegisteredContentScripts(filter ContentScriptFilter) (ret js.Promise[js.Array[RegisteredContentScript]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetRegisteredContentScripts(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

// HasFuncInsertCSS returns true if the function "WEBEXT.scripting.insertCSS" exists.
func HasFuncInsertCSS() bool {
	return js.True == bindings.HasFuncInsertCSS()
}

// FuncInsertCSS returns the function "WEBEXT.scripting.insertCSS".
func FuncInsertCSS() (fn js.Func[func(injection CSSInjection) js.Promise[js.Void]]) {
	bindings.FuncInsertCSS(
		js.Pointer(&fn),
	)
	return
}

// InsertCSS calls the function "WEBEXT.scripting.insertCSS" directly.
func InsertCSS(injection CSSInjection) (ret js.Promise[js.Void]) {
	bindings.CallInsertCSS(
		js.Pointer(&ret),
		js.Pointer(&injection),
	)

	return
}

// TryInsertCSS calls the function "WEBEXT.scripting.insertCSS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInsertCSS(injection CSSInjection) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryInsertCSS(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&injection),
	)

	return
}

// HasFuncRegisterContentScripts returns true if the function "WEBEXT.scripting.registerContentScripts" exists.
func HasFuncRegisterContentScripts() bool {
	return js.True == bindings.HasFuncRegisterContentScripts()
}

// FuncRegisterContentScripts returns the function "WEBEXT.scripting.registerContentScripts".
func FuncRegisterContentScripts() (fn js.Func[func(scripts js.Array[RegisteredContentScript]) js.Promise[js.Void]]) {
	bindings.FuncRegisterContentScripts(
		js.Pointer(&fn),
	)
	return
}

// RegisterContentScripts calls the function "WEBEXT.scripting.registerContentScripts" directly.
func RegisterContentScripts(scripts js.Array[RegisteredContentScript]) (ret js.Promise[js.Void]) {
	bindings.CallRegisterContentScripts(
		js.Pointer(&ret),
		scripts.Ref(),
	)

	return
}

// TryRegisterContentScripts calls the function "WEBEXT.scripting.registerContentScripts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRegisterContentScripts(scripts js.Array[RegisteredContentScript]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRegisterContentScripts(
		js.Pointer(&ret), js.Pointer(&exception),
		scripts.Ref(),
	)

	return
}

// HasFuncRemoveCSS returns true if the function "WEBEXT.scripting.removeCSS" exists.
func HasFuncRemoveCSS() bool {
	return js.True == bindings.HasFuncRemoveCSS()
}

// FuncRemoveCSS returns the function "WEBEXT.scripting.removeCSS".
func FuncRemoveCSS() (fn js.Func[func(injection CSSInjection) js.Promise[js.Void]]) {
	bindings.FuncRemoveCSS(
		js.Pointer(&fn),
	)
	return
}

// RemoveCSS calls the function "WEBEXT.scripting.removeCSS" directly.
func RemoveCSS(injection CSSInjection) (ret js.Promise[js.Void]) {
	bindings.CallRemoveCSS(
		js.Pointer(&ret),
		js.Pointer(&injection),
	)

	return
}

// TryRemoveCSS calls the function "WEBEXT.scripting.removeCSS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveCSS(injection CSSInjection) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveCSS(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&injection),
	)

	return
}

// HasFuncUnregisterContentScripts returns true if the function "WEBEXT.scripting.unregisterContentScripts" exists.
func HasFuncUnregisterContentScripts() bool {
	return js.True == bindings.HasFuncUnregisterContentScripts()
}

// FuncUnregisterContentScripts returns the function "WEBEXT.scripting.unregisterContentScripts".
func FuncUnregisterContentScripts() (fn js.Func[func(filter ContentScriptFilter) js.Promise[js.Void]]) {
	bindings.FuncUnregisterContentScripts(
		js.Pointer(&fn),
	)
	return
}

// UnregisterContentScripts calls the function "WEBEXT.scripting.unregisterContentScripts" directly.
func UnregisterContentScripts(filter ContentScriptFilter) (ret js.Promise[js.Void]) {
	bindings.CallUnregisterContentScripts(
		js.Pointer(&ret),
		js.Pointer(&filter),
	)

	return
}

// TryUnregisterContentScripts calls the function "WEBEXT.scripting.unregisterContentScripts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUnregisterContentScripts(filter ContentScriptFilter) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUnregisterContentScripts(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&filter),
	)

	return
}

// HasFuncUpdateContentScripts returns true if the function "WEBEXT.scripting.updateContentScripts" exists.
func HasFuncUpdateContentScripts() bool {
	return js.True == bindings.HasFuncUpdateContentScripts()
}

// FuncUpdateContentScripts returns the function "WEBEXT.scripting.updateContentScripts".
func FuncUpdateContentScripts() (fn js.Func[func(scripts js.Array[RegisteredContentScript]) js.Promise[js.Void]]) {
	bindings.FuncUpdateContentScripts(
		js.Pointer(&fn),
	)
	return
}

// UpdateContentScripts calls the function "WEBEXT.scripting.updateContentScripts" directly.
func UpdateContentScripts(scripts js.Array[RegisteredContentScript]) (ret js.Promise[js.Void]) {
	bindings.CallUpdateContentScripts(
		js.Pointer(&ret),
		scripts.Ref(),
	)

	return
}

// TryUpdateContentScripts calls the function "WEBEXT.scripting.updateContentScripts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUpdateContentScripts(scripts js.Array[RegisteredContentScript]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryUpdateContentScripts(
		js.Pointer(&ret), js.Pointer(&exception),
		scripts.Ref(),
	)

	return
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webviewtag

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/contextmenus"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensiontypes"
	"github.com/primecitizens/pcz/std/plat/js/webext/webviewtag/bindings"
)

type AttachArgWebview struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a AttachArgWebview with all fields set.
func (p AttachArgWebview) FromRef(ref js.Ref) AttachArgWebview {
	p.UpdateFrom(ref)
	return p
}

// New creates a new AttachArgWebview in the application heap.
func (p AttachArgWebview) New() js.Ref {
	return bindings.AttachArgWebviewJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *AttachArgWebview) UpdateFrom(ref js.Ref) {
	bindings.AttachArgWebviewJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *AttachArgWebview) Update(ref js.Ref) {
	bindings.AttachArgWebviewJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *AttachArgWebview) FreeMembers(recursive bool) {
}

type BackArgCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn BackArgCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn BackArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type BackArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *BackArgCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *BackArgCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CaptureVisibleRegionArgCallbackFunc func(this js.Ref, dataUrl js.String) js.Ref

func (fn CaptureVisibleRegionArgCallbackFunc) Register() js.Func[func(dataUrl js.String)] {
	return js.RegisterCallback[func(dataUrl js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CaptureVisibleRegionArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type CaptureVisibleRegionArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, dataUrl js.String) js.Ref
	Arg T
}

func (cb *CaptureVisibleRegionArgCallback[T]) Register() js.Func[func(dataUrl js.String)] {
	return js.RegisterCallback[func(dataUrl js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CaptureVisibleRegionArgCallback[T]) DispatchCallback(
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

		js.String{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ClearDataArgCallbackFunc func(this js.Ref) js.Ref

func (fn ClearDataArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ClearDataArgCallbackFunc) DispatchCallback(
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

type ClearDataArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ClearDataArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ClearDataArgCallback[T]) DispatchCallback(
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

type ClearDataOptions struct {
	// Since is "ClearDataOptions.since"
	//
	// Optional
	//
	// NOTE: FFI_USE_Since MUST be set to true to make this field effective.
	Since float64

	FFI_USE_Since bool // for Since.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClearDataOptions with all fields set.
func (p ClearDataOptions) FromRef(ref js.Ref) ClearDataOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClearDataOptions in the application heap.
func (p ClearDataOptions) New() js.Ref {
	return bindings.ClearDataOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClearDataOptions) UpdateFrom(ref js.Ref) {
	bindings.ClearDataOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClearDataOptions) Update(ref js.Ref) {
	bindings.ClearDataOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClearDataOptions) FreeMembers(recursive bool) {
}

type ClearDataTypeSet struct {
	// Appcache is "ClearDataTypeSet.appcache"
	//
	// Optional
	//
	// NOTE: FFI_USE_Appcache MUST be set to true to make this field effective.
	Appcache bool
	// Cache is "ClearDataTypeSet.cache"
	//
	// Optional
	//
	// NOTE: FFI_USE_Cache MUST be set to true to make this field effective.
	Cache bool
	// Cookies is "ClearDataTypeSet.cookies"
	//
	// Optional
	//
	// NOTE: FFI_USE_Cookies MUST be set to true to make this field effective.
	Cookies bool
	// FileSystems is "ClearDataTypeSet.fileSystems"
	//
	// Optional
	//
	// NOTE: FFI_USE_FileSystems MUST be set to true to make this field effective.
	FileSystems bool
	// IndexedDB is "ClearDataTypeSet.indexedDB"
	//
	// Optional
	//
	// NOTE: FFI_USE_IndexedDB MUST be set to true to make this field effective.
	IndexedDB bool
	// LocalStorage is "ClearDataTypeSet.localStorage"
	//
	// Optional
	//
	// NOTE: FFI_USE_LocalStorage MUST be set to true to make this field effective.
	LocalStorage bool
	// PersistentCookies is "ClearDataTypeSet.persistentCookies"
	//
	// Optional
	//
	// NOTE: FFI_USE_PersistentCookies MUST be set to true to make this field effective.
	PersistentCookies bool
	// SessionCookies is "ClearDataTypeSet.sessionCookies"
	//
	// Optional
	//
	// NOTE: FFI_USE_SessionCookies MUST be set to true to make this field effective.
	SessionCookies bool
	// WebSQL is "ClearDataTypeSet.webSQL"
	//
	// Optional
	//
	// NOTE: FFI_USE_WebSQL MUST be set to true to make this field effective.
	WebSQL bool

	FFI_USE_Appcache          bool // for Appcache.
	FFI_USE_Cache             bool // for Cache.
	FFI_USE_Cookies           bool // for Cookies.
	FFI_USE_FileSystems       bool // for FileSystems.
	FFI_USE_IndexedDB         bool // for IndexedDB.
	FFI_USE_LocalStorage      bool // for LocalStorage.
	FFI_USE_PersistentCookies bool // for PersistentCookies.
	FFI_USE_SessionCookies    bool // for SessionCookies.
	FFI_USE_WebSQL            bool // for WebSQL.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ClearDataTypeSet with all fields set.
func (p ClearDataTypeSet) FromRef(ref js.Ref) ClearDataTypeSet {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ClearDataTypeSet in the application heap.
func (p ClearDataTypeSet) New() js.Ref {
	return bindings.ClearDataTypeSetJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ClearDataTypeSet) UpdateFrom(ref js.Ref) {
	bindings.ClearDataTypeSetJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ClearDataTypeSet) Update(ref js.Ref) {
	bindings.ClearDataTypeSetJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ClearDataTypeSet) FreeMembers(recursive bool) {
}

type InjectionItems struct {
	// Code is "InjectionItems.code"
	//
	// Optional
	Code js.String
	// Files is "InjectionItems.files"
	//
	// Optional
	Files js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InjectionItems with all fields set.
func (p InjectionItems) FromRef(ref js.Ref) InjectionItems {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InjectionItems in the application heap.
func (p InjectionItems) New() js.Ref {
	return bindings.InjectionItemsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InjectionItems) UpdateFrom(ref js.Ref) {
	bindings.InjectionItemsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InjectionItems) Update(ref js.Ref) {
	bindings.InjectionItemsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InjectionItems) FreeMembers(recursive bool) {
	js.Free(
		p.Code.Ref(),
		p.Files.Ref(),
	)
	p.Code = p.Code.FromRef(js.Undefined)
	p.Files = p.Files.FromRef(js.Undefined)
}

type ContentScriptDetails struct {
	// AllFrames is "ContentScriptDetails.all_frames"
	//
	// Optional
	//
	// NOTE: FFI_USE_AllFrames MUST be set to true to make this field effective.
	AllFrames bool
	// Css is "ContentScriptDetails.css"
	//
	// Optional
	//
	// NOTE: Css.FFI_USE MUST be set to true to get Css used.
	Css InjectionItems
	// ExcludeGlobs is "ContentScriptDetails.exclude_globs"
	//
	// Optional
	ExcludeGlobs js.Array[js.String]
	// ExcludeMatches is "ContentScriptDetails.exclude_matches"
	//
	// Optional
	ExcludeMatches js.Array[js.String]
	// IncludeGlobs is "ContentScriptDetails.include_globs"
	//
	// Optional
	IncludeGlobs js.Array[js.String]
	// Js is "ContentScriptDetails.js"
	//
	// Optional
	//
	// NOTE: Js.FFI_USE MUST be set to true to get Js used.
	Js InjectionItems
	// MatchAboutBlank is "ContentScriptDetails.match_about_blank"
	//
	// Optional
	//
	// NOTE: FFI_USE_MatchAboutBlank MUST be set to true to make this field effective.
	MatchAboutBlank bool
	// Matches is "ContentScriptDetails.matches"
	//
	// Required
	Matches js.Array[js.String]
	// Name is "ContentScriptDetails.name"
	//
	// Required
	Name js.String
	// RunAt is "ContentScriptDetails.run_at"
	//
	// Optional
	RunAt extensiontypes.RunAt

	FFI_USE_AllFrames       bool // for AllFrames.
	FFI_USE_MatchAboutBlank bool // for MatchAboutBlank.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContentScriptDetails with all fields set.
func (p ContentScriptDetails) FromRef(ref js.Ref) ContentScriptDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContentScriptDetails in the application heap.
func (p ContentScriptDetails) New() js.Ref {
	return bindings.ContentScriptDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContentScriptDetails) UpdateFrom(ref js.Ref) {
	bindings.ContentScriptDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContentScriptDetails) Update(ref js.Ref) {
	bindings.ContentScriptDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContentScriptDetails) FreeMembers(recursive bool) {
	js.Free(
		p.ExcludeGlobs.Ref(),
		p.ExcludeMatches.Ref(),
		p.IncludeGlobs.Ref(),
		p.Matches.Ref(),
		p.Name.Ref(),
	)
	p.ExcludeGlobs = p.ExcludeGlobs.FromRef(js.Undefined)
	p.ExcludeMatches = p.ExcludeMatches.FromRef(js.Undefined)
	p.IncludeGlobs = p.IncludeGlobs.FromRef(js.Undefined)
	p.Matches = p.Matches.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
	if recursive {
		p.Css.FreeMembers(true)
		p.Js.FreeMembers(true)
	}
}

type ContentWindowType struct {
	ref js.Ref
}

func (this ContentWindowType) Once() ContentWindowType {
	this.ref.Once()
	return this
}

func (this ContentWindowType) Ref() js.Ref {
	return this.ref
}

func (this ContentWindowType) FromRef(ref js.Ref) ContentWindowType {
	this.ref = ref
	return this
}

func (this ContentWindowType) Free() {
	this.ref.Free()
}

// HasFuncPostMessage returns true if the method "ContentWindowType.postMessage" exists.
func (this ContentWindowType) HasFuncPostMessage() bool {
	return js.True == bindings.HasFuncContentWindowTypePostMessage(
		this.ref,
	)
}

// FuncPostMessage returns the method "ContentWindowType.postMessage".
func (this ContentWindowType) FuncPostMessage() (fn js.Func[func(message js.Any, targetOrigin js.String)]) {
	bindings.FuncContentWindowTypePostMessage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage calls the method "ContentWindowType.postMessage".
func (this ContentWindowType) PostMessage(message js.Any, targetOrigin js.String) (ret js.Void) {
	bindings.CallContentWindowTypePostMessage(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		targetOrigin.Ref(),
	)

	return
}

// TryPostMessage calls the method "ContentWindowType.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContentWindowType) TryPostMessage(message js.Any, targetOrigin js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContentWindowTypePostMessage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		targetOrigin.Ref(),
	)

	return
}

type ContextType uint32

const (
	_ ContextType = iota

	ContextType_ALL
	ContextType_PAGE
	ContextType_FRAME
	ContextType_SELECTION
	ContextType_LINK
	ContextType_EDITABLE
	ContextType_IMAGE
	ContextType_VIDEO
	ContextType_AUDIO
)

func (ContextType) FromRef(str js.Ref) ContextType {
	return ContextType(bindings.ConstOfContextType(str))
}

func (x ContextType) String() (string, bool) {
	switch x {
	case ContextType_ALL:
		return "all", true
	case ContextType_PAGE:
		return "page", true
	case ContextType_FRAME:
		return "frame", true
	case ContextType_SELECTION:
		return "selection", true
	case ContextType_LINK:
		return "link", true
	case ContextType_EDITABLE:
		return "editable", true
	case ContextType_IMAGE:
		return "image", true
	case ContextType_VIDEO:
		return "video", true
	case ContextType_AUDIO:
		return "audio", true
	default:
		return "", false
	}
}

type ContextMenuCreatePropertiesFieldOnclickFunc func(this js.Ref, info *contextmenus.OnClickData) js.Ref

func (fn ContextMenuCreatePropertiesFieldOnclickFunc) Register() js.Func[func(info *contextmenus.OnClickData)] {
	return js.RegisterCallback[func(info *contextmenus.OnClickData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ContextMenuCreatePropertiesFieldOnclickFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 contextmenus.OnClickData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ContextMenuCreatePropertiesFieldOnclick[T any] struct {
	Fn  func(arg T, this js.Ref, info *contextmenus.OnClickData) js.Ref
	Arg T
}

func (cb *ContextMenuCreatePropertiesFieldOnclick[T]) Register() js.Func[func(info *contextmenus.OnClickData)] {
	return js.RegisterCallback[func(info *contextmenus.OnClickData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ContextMenuCreatePropertiesFieldOnclick[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 contextmenus.OnClickData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type OneOf_Int64_String struct {
	ref js.Ref
}

func (x OneOf_Int64_String) Ref() js.Ref {
	return x.ref
}

func (x OneOf_Int64_String) Free() {
	x.ref.Free()
}

func (x OneOf_Int64_String) FromRef(ref js.Ref) OneOf_Int64_String {
	return OneOf_Int64_String{
		ref: ref,
	}
}

func (x OneOf_Int64_String) Int64() int64 {
	return js.BigInt[int64]{}.FromRef(x.ref).Get()
}

func (x OneOf_Int64_String) String() js.String {
	return js.String{}.FromRef(x.ref)
}

type ContextMenuCreateProperties struct {
	// Checked is "ContextMenuCreateProperties.checked"
	//
	// Optional
	//
	// NOTE: FFI_USE_Checked MUST be set to true to make this field effective.
	Checked bool
	// Contexts is "ContextMenuCreateProperties.contexts"
	//
	// Optional
	Contexts js.Array[ContextType]
	// DocumentUrlPatterns is "ContextMenuCreateProperties.documentUrlPatterns"
	//
	// Optional
	DocumentUrlPatterns js.Array[js.String]
	// Enabled is "ContextMenuCreateProperties.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// Id is "ContextMenuCreateProperties.id"
	//
	// Optional
	Id js.String
	// Onclick is "ContextMenuCreateProperties.onclick"
	//
	// Optional
	Onclick js.Func[func(info *contextmenus.OnClickData)]
	// ParentId is "ContextMenuCreateProperties.parentId"
	//
	// Optional
	ParentId OneOf_Int64_String
	// TargetUrlPatterns is "ContextMenuCreateProperties.targetUrlPatterns"
	//
	// Optional
	TargetUrlPatterns js.Array[js.String]
	// Title is "ContextMenuCreateProperties.title"
	//
	// Optional
	Title js.String
	// Type is "ContextMenuCreateProperties.type"
	//
	// Optional
	Type contextmenus.ItemType

	FFI_USE_Checked bool // for Checked.
	FFI_USE_Enabled bool // for Enabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContextMenuCreateProperties with all fields set.
func (p ContextMenuCreateProperties) FromRef(ref js.Ref) ContextMenuCreateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContextMenuCreateProperties in the application heap.
func (p ContextMenuCreateProperties) New() js.Ref {
	return bindings.ContextMenuCreatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContextMenuCreateProperties) UpdateFrom(ref js.Ref) {
	bindings.ContextMenuCreatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContextMenuCreateProperties) Update(ref js.Ref) {
	bindings.ContextMenuCreatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContextMenuCreateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Contexts.Ref(),
		p.DocumentUrlPatterns.Ref(),
		p.Id.Ref(),
		p.Onclick.Ref(),
		p.ParentId.Ref(),
		p.TargetUrlPatterns.Ref(),
		p.Title.Ref(),
	)
	p.Contexts = p.Contexts.FromRef(js.Undefined)
	p.DocumentUrlPatterns = p.DocumentUrlPatterns.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Onclick = p.Onclick.FromRef(js.Undefined)
	p.ParentId = p.ParentId.FromRef(js.Undefined)
	p.TargetUrlPatterns = p.TargetUrlPatterns.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type ContextMenuUpdatePropertiesFieldOnclickFunc func(this js.Ref, info *contextmenus.OnClickData) js.Ref

func (fn ContextMenuUpdatePropertiesFieldOnclickFunc) Register() js.Func[func(info *contextmenus.OnClickData)] {
	return js.RegisterCallback[func(info *contextmenus.OnClickData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ContextMenuUpdatePropertiesFieldOnclickFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 contextmenus.OnClickData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ContextMenuUpdatePropertiesFieldOnclick[T any] struct {
	Fn  func(arg T, this js.Ref, info *contextmenus.OnClickData) js.Ref
	Arg T
}

func (cb *ContextMenuUpdatePropertiesFieldOnclick[T]) Register() js.Func[func(info *contextmenus.OnClickData)] {
	return js.RegisterCallback[func(info *contextmenus.OnClickData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ContextMenuUpdatePropertiesFieldOnclick[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 contextmenus.OnClickData
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ContextMenuUpdateProperties struct {
	// Checked is "ContextMenuUpdateProperties.checked"
	//
	// Optional
	//
	// NOTE: FFI_USE_Checked MUST be set to true to make this field effective.
	Checked bool
	// Contexts is "ContextMenuUpdateProperties.contexts"
	//
	// Optional
	Contexts js.Array[ContextType]
	// DocumentUrlPatterns is "ContextMenuUpdateProperties.documentUrlPatterns"
	//
	// Optional
	DocumentUrlPatterns js.Array[js.String]
	// Enabled is "ContextMenuUpdateProperties.enabled"
	//
	// Optional
	//
	// NOTE: FFI_USE_Enabled MUST be set to true to make this field effective.
	Enabled bool
	// Onclick is "ContextMenuUpdateProperties.onclick"
	//
	// Optional
	Onclick js.Func[func(info *contextmenus.OnClickData)]
	// ParentId is "ContextMenuUpdateProperties.parentId"
	//
	// Optional
	ParentId OneOf_Int64_String
	// TargetUrlPatterns is "ContextMenuUpdateProperties.targetUrlPatterns"
	//
	// Optional
	TargetUrlPatterns js.Array[js.String]
	// Title is "ContextMenuUpdateProperties.title"
	//
	// Optional
	Title js.String
	// Type is "ContextMenuUpdateProperties.type"
	//
	// Optional
	Type contextmenus.ItemType

	FFI_USE_Checked bool // for Checked.
	FFI_USE_Enabled bool // for Enabled.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a ContextMenuUpdateProperties with all fields set.
func (p ContextMenuUpdateProperties) FromRef(ref js.Ref) ContextMenuUpdateProperties {
	p.UpdateFrom(ref)
	return p
}

// New creates a new ContextMenuUpdateProperties in the application heap.
func (p ContextMenuUpdateProperties) New() js.Ref {
	return bindings.ContextMenuUpdatePropertiesJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *ContextMenuUpdateProperties) UpdateFrom(ref js.Ref) {
	bindings.ContextMenuUpdatePropertiesJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *ContextMenuUpdateProperties) Update(ref js.Ref) {
	bindings.ContextMenuUpdatePropertiesJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *ContextMenuUpdateProperties) FreeMembers(recursive bool) {
	js.Free(
		p.Contexts.Ref(),
		p.DocumentUrlPatterns.Ref(),
		p.Onclick.Ref(),
		p.ParentId.Ref(),
		p.TargetUrlPatterns.Ref(),
		p.Title.Ref(),
	)
	p.Contexts = p.Contexts.FromRef(js.Undefined)
	p.DocumentUrlPatterns = p.DocumentUrlPatterns.FromRef(js.Undefined)
	p.Onclick = p.Onclick.FromRef(js.Undefined)
	p.ParentId = p.ParentId.FromRef(js.Undefined)
	p.TargetUrlPatterns = p.TargetUrlPatterns.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
}

type CreateArgCallbackFunc func(this js.Ref) js.Ref

func (fn CreateArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CreateArgCallbackFunc) DispatchCallback(
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

type CreateArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CreateArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CreateArgCallback[T]) DispatchCallback(
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

type UpdateArgCallbackFunc func(this js.Ref) js.Ref

func (fn UpdateArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UpdateArgCallbackFunc) DispatchCallback(
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

type UpdateArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *UpdateArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UpdateArgCallback[T]) DispatchCallback(
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

type RemoveArgCallbackFunc func(this js.Ref) js.Ref

func (fn RemoveArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RemoveArgCallbackFunc) DispatchCallback(
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

type RemoveArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RemoveArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RemoveArgCallback[T]) DispatchCallback(
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

type RemoveAllArgCallbackFunc func(this js.Ref) js.Ref

func (fn RemoveAllArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn RemoveAllArgCallbackFunc) DispatchCallback(
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

type RemoveAllArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *RemoveAllArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *RemoveAllArgCallback[T]) DispatchCallback(
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

type ContextMenusType struct {
	ref js.Ref
}

func (this ContextMenusType) Once() ContextMenusType {
	this.ref.Once()
	return this
}

func (this ContextMenusType) Ref() js.Ref {
	return this.ref
}

func (this ContextMenusType) FromRef(ref js.Ref) ContextMenusType {
	this.ref = ref
	return this
}

func (this ContextMenusType) Free() {
	this.ref.Free()
}

// HasFuncCreate returns true if the method "ContextMenusType.create" exists.
func (this ContextMenusType) HasFuncCreate() bool {
	return js.True == bindings.HasFuncContextMenusTypeCreate(
		this.ref,
	)
}

// FuncCreate returns the method "ContextMenusType.create".
func (this ContextMenusType) FuncCreate() (fn js.Func[func(createProperties ContextMenuCreateProperties, callback js.Func[func()]) OneOf_Int64_String]) {
	bindings.FuncContextMenusTypeCreate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Create calls the method "ContextMenusType.create".
func (this ContextMenusType) Create(createProperties ContextMenuCreateProperties, callback js.Func[func()]) (ret OneOf_Int64_String) {
	bindings.CallContextMenusTypeCreate(
		this.ref, js.Pointer(&ret),
		js.Pointer(&createProperties),
		callback.Ref(),
	)

	return
}

// TryCreate calls the method "ContextMenusType.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContextMenusType) TryCreate(createProperties ContextMenuCreateProperties, callback js.Func[func()]) (ret OneOf_Int64_String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusTypeCreate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&createProperties),
		callback.Ref(),
	)

	return
}

// HasFuncCreate1 returns true if the method "ContextMenusType.create" exists.
func (this ContextMenusType) HasFuncCreate1() bool {
	return js.True == bindings.HasFuncContextMenusTypeCreate1(
		this.ref,
	)
}

// FuncCreate1 returns the method "ContextMenusType.create".
func (this ContextMenusType) FuncCreate1() (fn js.Func[func(createProperties ContextMenuCreateProperties) OneOf_Int64_String]) {
	bindings.FuncContextMenusTypeCreate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Create1 calls the method "ContextMenusType.create".
func (this ContextMenusType) Create1(createProperties ContextMenuCreateProperties) (ret OneOf_Int64_String) {
	bindings.CallContextMenusTypeCreate1(
		this.ref, js.Pointer(&ret),
		js.Pointer(&createProperties),
	)

	return
}

// TryCreate1 calls the method "ContextMenusType.create"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContextMenusType) TryCreate1(createProperties ContextMenuCreateProperties) (ret OneOf_Int64_String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusTypeCreate1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&createProperties),
	)

	return
}

// HasFuncUpdate returns true if the method "ContextMenusType.update" exists.
func (this ContextMenusType) HasFuncUpdate() bool {
	return js.True == bindings.HasFuncContextMenusTypeUpdate(
		this.ref,
	)
}

// FuncUpdate returns the method "ContextMenusType.update".
func (this ContextMenusType) FuncUpdate() (fn js.Func[func(id OneOf_Int64_String, updateProperties ContextMenuUpdateProperties, callback js.Func[func()])]) {
	bindings.FuncContextMenusTypeUpdate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Update calls the method "ContextMenusType.update".
func (this ContextMenusType) Update(id OneOf_Int64_String, updateProperties ContextMenuUpdateProperties, callback js.Func[func()]) (ret js.Void) {
	bindings.CallContextMenusTypeUpdate(
		this.ref, js.Pointer(&ret),
		id.Ref(),
		js.Pointer(&updateProperties),
		callback.Ref(),
	)

	return
}

// TryUpdate calls the method "ContextMenusType.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContextMenusType) TryUpdate(id OneOf_Int64_String, updateProperties ContextMenuUpdateProperties, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusTypeUpdate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Pointer(&updateProperties),
		callback.Ref(),
	)

	return
}

// HasFuncUpdate1 returns true if the method "ContextMenusType.update" exists.
func (this ContextMenusType) HasFuncUpdate1() bool {
	return js.True == bindings.HasFuncContextMenusTypeUpdate1(
		this.ref,
	)
}

// FuncUpdate1 returns the method "ContextMenusType.update".
func (this ContextMenusType) FuncUpdate1() (fn js.Func[func(id OneOf_Int64_String, updateProperties ContextMenuUpdateProperties)]) {
	bindings.FuncContextMenusTypeUpdate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Update1 calls the method "ContextMenusType.update".
func (this ContextMenusType) Update1(id OneOf_Int64_String, updateProperties ContextMenuUpdateProperties) (ret js.Void) {
	bindings.CallContextMenusTypeUpdate1(
		this.ref, js.Pointer(&ret),
		id.Ref(),
		js.Pointer(&updateProperties),
	)

	return
}

// TryUpdate1 calls the method "ContextMenusType.update"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContextMenusType) TryUpdate1(id OneOf_Int64_String, updateProperties ContextMenuUpdateProperties) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusTypeUpdate1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Pointer(&updateProperties),
	)

	return
}

// HasFuncRemove returns true if the method "ContextMenusType.remove" exists.
func (this ContextMenusType) HasFuncRemove() bool {
	return js.True == bindings.HasFuncContextMenusTypeRemove(
		this.ref,
	)
}

// FuncRemove returns the method "ContextMenusType.remove".
func (this ContextMenusType) FuncRemove() (fn js.Func[func(menuItemId OneOf_Int64_String, callback js.Func[func()])]) {
	bindings.FuncContextMenusTypeRemove(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove calls the method "ContextMenusType.remove".
func (this ContextMenusType) Remove(menuItemId OneOf_Int64_String, callback js.Func[func()]) (ret js.Void) {
	bindings.CallContextMenusTypeRemove(
		this.ref, js.Pointer(&ret),
		menuItemId.Ref(),
		callback.Ref(),
	)

	return
}

// TryRemove calls the method "ContextMenusType.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContextMenusType) TryRemove(menuItemId OneOf_Int64_String, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusTypeRemove(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		menuItemId.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncRemove1 returns true if the method "ContextMenusType.remove" exists.
func (this ContextMenusType) HasFuncRemove1() bool {
	return js.True == bindings.HasFuncContextMenusTypeRemove1(
		this.ref,
	)
}

// FuncRemove1 returns the method "ContextMenusType.remove".
func (this ContextMenusType) FuncRemove1() (fn js.Func[func(menuItemId OneOf_Int64_String)]) {
	bindings.FuncContextMenusTypeRemove1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove1 calls the method "ContextMenusType.remove".
func (this ContextMenusType) Remove1(menuItemId OneOf_Int64_String) (ret js.Void) {
	bindings.CallContextMenusTypeRemove1(
		this.ref, js.Pointer(&ret),
		menuItemId.Ref(),
	)

	return
}

// TryRemove1 calls the method "ContextMenusType.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContextMenusType) TryRemove1(menuItemId OneOf_Int64_String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusTypeRemove1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		menuItemId.Ref(),
	)

	return
}

// HasFuncRemoveAll returns true if the method "ContextMenusType.removeAll" exists.
func (this ContextMenusType) HasFuncRemoveAll() bool {
	return js.True == bindings.HasFuncContextMenusTypeRemoveAll(
		this.ref,
	)
}

// FuncRemoveAll returns the method "ContextMenusType.removeAll".
func (this ContextMenusType) FuncRemoveAll() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncContextMenusTypeRemoveAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveAll calls the method "ContextMenusType.removeAll".
func (this ContextMenusType) RemoveAll(callback js.Func[func()]) (ret js.Void) {
	bindings.CallContextMenusTypeRemoveAll(
		this.ref, js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryRemoveAll calls the method "ContextMenusType.removeAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContextMenusType) TryRemoveAll(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusTypeRemoveAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncRemoveAll1 returns true if the method "ContextMenusType.removeAll" exists.
func (this ContextMenusType) HasFuncRemoveAll1() bool {
	return js.True == bindings.HasFuncContextMenusTypeRemoveAll1(
		this.ref,
	)
}

// FuncRemoveAll1 returns the method "ContextMenusType.removeAll".
func (this ContextMenusType) FuncRemoveAll1() (fn js.Func[func()]) {
	bindings.FuncContextMenusTypeRemoveAll1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// RemoveAll1 calls the method "ContextMenusType.removeAll".
func (this ContextMenusType) RemoveAll1() (ret js.Void) {
	bindings.CallContextMenusTypeRemoveAll1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRemoveAll1 calls the method "ContextMenusType.removeAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this ContextMenusType) TryRemoveAll1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryContextMenusTypeRemoveAll1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type DialogArgMessageType uint32

const (
	_ DialogArgMessageType = iota

	DialogArgMessageType_ALERT
	DialogArgMessageType_CONFIRM
	DialogArgMessageType_PROMPT
)

func (DialogArgMessageType) FromRef(str js.Ref) DialogArgMessageType {
	return DialogArgMessageType(bindings.ConstOfDialogArgMessageType(str))
}

func (x DialogArgMessageType) String() (string, bool) {
	switch x {
	case DialogArgMessageType_ALERT:
		return "alert", true
	case DialogArgMessageType_CONFIRM:
		return "confirm", true
	case DialogArgMessageType_PROMPT:
		return "prompt", true
	default:
		return "", false
	}
}

type DialogController struct {
	ref js.Ref
}

func (this DialogController) Once() DialogController {
	this.ref.Once()
	return this
}

func (this DialogController) Ref() js.Ref {
	return this.ref
}

func (this DialogController) FromRef(ref js.Ref) DialogController {
	this.ref = ref
	return this
}

func (this DialogController) Free() {
	this.ref.Free()
}

// HasFuncOk returns true if the method "DialogController.ok" exists.
func (this DialogController) HasFuncOk() bool {
	return js.True == bindings.HasFuncDialogControllerOk(
		this.ref,
	)
}

// FuncOk returns the method "DialogController.ok".
func (this DialogController) FuncOk() (fn js.Func[func(response js.String)]) {
	bindings.FuncDialogControllerOk(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Ok calls the method "DialogController.ok".
func (this DialogController) Ok(response js.String) (ret js.Void) {
	bindings.CallDialogControllerOk(
		this.ref, js.Pointer(&ret),
		response.Ref(),
	)

	return
}

// TryOk calls the method "DialogController.ok"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DialogController) TryOk(response js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDialogControllerOk(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		response.Ref(),
	)

	return
}

// HasFuncOk1 returns true if the method "DialogController.ok" exists.
func (this DialogController) HasFuncOk1() bool {
	return js.True == bindings.HasFuncDialogControllerOk1(
		this.ref,
	)
}

// FuncOk1 returns the method "DialogController.ok".
func (this DialogController) FuncOk1() (fn js.Func[func()]) {
	bindings.FuncDialogControllerOk1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Ok1 calls the method "DialogController.ok".
func (this DialogController) Ok1() (ret js.Void) {
	bindings.CallDialogControllerOk1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryOk1 calls the method "DialogController.ok"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DialogController) TryOk1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDialogControllerOk1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCancel returns true if the method "DialogController.cancel" exists.
func (this DialogController) HasFuncCancel() bool {
	return js.True == bindings.HasFuncDialogControllerCancel(
		this.ref,
	)
}

// FuncCancel returns the method "DialogController.cancel".
func (this DialogController) FuncCancel() (fn js.Func[func()]) {
	bindings.FuncDialogControllerCancel(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Cancel calls the method "DialogController.cancel".
func (this DialogController) Cancel() (ret js.Void) {
	bindings.CallDialogControllerCancel(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCancel calls the method "DialogController.cancel"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DialogController) TryCancel() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDialogControllerCancel(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type DownloadPermissionRequest struct {
	ref js.Ref
}

func (this DownloadPermissionRequest) Once() DownloadPermissionRequest {
	this.ref.Once()
	return this
}

func (this DownloadPermissionRequest) Ref() js.Ref {
	return this.ref
}

func (this DownloadPermissionRequest) FromRef(ref js.Ref) DownloadPermissionRequest {
	this.ref = ref
	return this
}

func (this DownloadPermissionRequest) Free() {
	this.ref.Free()
}

// RequestMethod returns the value of property "DownloadPermissionRequest.requestMethod".
//
// It returns ok=false if there is no such property.
func (this DownloadPermissionRequest) RequestMethod() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDownloadPermissionRequestRequestMethod(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRequestMethod sets the value of property "DownloadPermissionRequest.requestMethod" to val.
//
// It returns false if the property cannot be set.
func (this DownloadPermissionRequest) SetRequestMethod(val js.String) bool {
	return js.True == bindings.SetDownloadPermissionRequestRequestMethod(
		this.ref,
		val.Ref(),
	)
}

// Url returns the value of property "DownloadPermissionRequest.url".
//
// It returns ok=false if there is no such property.
func (this DownloadPermissionRequest) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetDownloadPermissionRequestUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUrl sets the value of property "DownloadPermissionRequest.url" to val.
//
// It returns false if the property cannot be set.
func (this DownloadPermissionRequest) SetUrl(val js.String) bool {
	return js.True == bindings.SetDownloadPermissionRequestUrl(
		this.ref,
		val.Ref(),
	)
}

// HasFuncAllow returns true if the method "DownloadPermissionRequest.allow" exists.
func (this DownloadPermissionRequest) HasFuncAllow() bool {
	return js.True == bindings.HasFuncDownloadPermissionRequestAllow(
		this.ref,
	)
}

// FuncAllow returns the method "DownloadPermissionRequest.allow".
func (this DownloadPermissionRequest) FuncAllow() (fn js.Func[func()]) {
	bindings.FuncDownloadPermissionRequestAllow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Allow calls the method "DownloadPermissionRequest.allow".
func (this DownloadPermissionRequest) Allow() (ret js.Void) {
	bindings.CallDownloadPermissionRequestAllow(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAllow calls the method "DownloadPermissionRequest.allow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DownloadPermissionRequest) TryAllow() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDownloadPermissionRequestAllow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeny returns true if the method "DownloadPermissionRequest.deny" exists.
func (this DownloadPermissionRequest) HasFuncDeny() bool {
	return js.True == bindings.HasFuncDownloadPermissionRequestDeny(
		this.ref,
	)
}

// FuncDeny returns the method "DownloadPermissionRequest.deny".
func (this DownloadPermissionRequest) FuncDeny() (fn js.Func[func()]) {
	bindings.FuncDownloadPermissionRequestDeny(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Deny calls the method "DownloadPermissionRequest.deny".
func (this DownloadPermissionRequest) Deny() (ret js.Void) {
	bindings.CallDownloadPermissionRequestDeny(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeny calls the method "DownloadPermissionRequest.deny"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this DownloadPermissionRequest) TryDeny() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryDownloadPermissionRequestDeny(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type ExecuteScriptArgCallbackFunc func(this js.Ref, result js.Array[js.Any]) js.Ref

func (fn ExecuteScriptArgCallbackFunc) Register() js.Func[func(result js.Array[js.Any])] {
	return js.RegisterCallback[func(result js.Array[js.Any])](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ExecuteScriptArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Array[js.Any]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ExecuteScriptArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result js.Array[js.Any]) js.Ref
	Arg T
}

func (cb *ExecuteScriptArgCallback[T]) Register() js.Func[func(result js.Array[js.Any])] {
	return js.RegisterCallback[func(result js.Array[js.Any])](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ExecuteScriptArgCallback[T]) DispatchCallback(
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

		js.Array[js.Any]{}.FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ExitArgReason uint32

const (
	_ ExitArgReason = iota

	ExitArgReason_NORMAL
	ExitArgReason_ABNORMAL
	ExitArgReason_CRASH
	ExitArgReason_KILL
)

func (ExitArgReason) FromRef(str js.Ref) ExitArgReason {
	return ExitArgReason(bindings.ConstOfExitArgReason(str))
}

func (x ExitArgReason) String() (string, bool) {
	switch x {
	case ExitArgReason_NORMAL:
		return "normal", true
	case ExitArgReason_ABNORMAL:
		return "abnormal", true
	case ExitArgReason_CRASH:
		return "crash", true
	case ExitArgReason_KILL:
		return "kill", true
	default:
		return "", false
	}
}

type FileSystemPermissionRequest struct {
	ref js.Ref
}

func (this FileSystemPermissionRequest) Once() FileSystemPermissionRequest {
	this.ref.Once()
	return this
}

func (this FileSystemPermissionRequest) Ref() js.Ref {
	return this.ref
}

func (this FileSystemPermissionRequest) FromRef(ref js.Ref) FileSystemPermissionRequest {
	this.ref = ref
	return this
}

func (this FileSystemPermissionRequest) Free() {
	this.ref.Free()
}

// Url returns the value of property "FileSystemPermissionRequest.url".
//
// It returns ok=false if there is no such property.
func (this FileSystemPermissionRequest) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFileSystemPermissionRequestUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUrl sets the value of property "FileSystemPermissionRequest.url" to val.
//
// It returns false if the property cannot be set.
func (this FileSystemPermissionRequest) SetUrl(val js.String) bool {
	return js.True == bindings.SetFileSystemPermissionRequestUrl(
		this.ref,
		val.Ref(),
	)
}

// HasFuncAllow returns true if the method "FileSystemPermissionRequest.allow" exists.
func (this FileSystemPermissionRequest) HasFuncAllow() bool {
	return js.True == bindings.HasFuncFileSystemPermissionRequestAllow(
		this.ref,
	)
}

// FuncAllow returns the method "FileSystemPermissionRequest.allow".
func (this FileSystemPermissionRequest) FuncAllow() (fn js.Func[func()]) {
	bindings.FuncFileSystemPermissionRequestAllow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Allow calls the method "FileSystemPermissionRequest.allow".
func (this FileSystemPermissionRequest) Allow() (ret js.Void) {
	bindings.CallFileSystemPermissionRequestAllow(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAllow calls the method "FileSystemPermissionRequest.allow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemPermissionRequest) TryAllow() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemPermissionRequestAllow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeny returns true if the method "FileSystemPermissionRequest.deny" exists.
func (this FileSystemPermissionRequest) HasFuncDeny() bool {
	return js.True == bindings.HasFuncFileSystemPermissionRequestDeny(
		this.ref,
	)
}

// FuncDeny returns the method "FileSystemPermissionRequest.deny".
func (this FileSystemPermissionRequest) FuncDeny() (fn js.Func[func()]) {
	bindings.FuncFileSystemPermissionRequestDeny(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Deny calls the method "FileSystemPermissionRequest.deny".
func (this FileSystemPermissionRequest) Deny() (ret js.Void) {
	bindings.CallFileSystemPermissionRequestDeny(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeny calls the method "FileSystemPermissionRequest.deny"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FileSystemPermissionRequest) TryDeny() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFileSystemPermissionRequestDeny(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type FindArgCallbackFunc func(this js.Ref, results *FindCallbackResults) js.Ref

func (fn FindArgCallbackFunc) Register() js.Func[func(results *FindCallbackResults)] {
	return js.RegisterCallback[func(results *FindCallbackResults)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FindArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FindCallbackResults
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FindArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, results *FindCallbackResults) js.Ref
	Arg T
}

func (cb *FindArgCallback[T]) Register() js.Func[func(results *FindCallbackResults)] {
	return js.RegisterCallback[func(results *FindCallbackResults)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FindArgCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 FindCallbackResults
	arg0.UpdateFrom(args[0+1])
	defer arg0.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		mark.NoEscape(&arg0),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SelectionRect struct {
	// Height is "SelectionRect.height"
	//
	// Required
	Height int64
	// Left is "SelectionRect.left"
	//
	// Required
	Left int64
	// Top is "SelectionRect.top"
	//
	// Required
	Top int64
	// Width is "SelectionRect.width"
	//
	// Required
	Width int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SelectionRect with all fields set.
func (p SelectionRect) FromRef(ref js.Ref) SelectionRect {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SelectionRect in the application heap.
func (p SelectionRect) New() js.Ref {
	return bindings.SelectionRectJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SelectionRect) UpdateFrom(ref js.Ref) {
	bindings.SelectionRectJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SelectionRect) Update(ref js.Ref) {
	bindings.SelectionRectJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SelectionRect) FreeMembers(recursive bool) {
}

type FindCallbackResults struct {
	// ActiveMatchOrdinal is "FindCallbackResults.activeMatchOrdinal"
	//
	// Required
	ActiveMatchOrdinal int64
	// Canceled is "FindCallbackResults.canceled"
	//
	// Required
	Canceled bool
	// NumberOfMatches is "FindCallbackResults.numberOfMatches"
	//
	// Required
	NumberOfMatches int64
	// SelectionRect is "FindCallbackResults.selectionRect"
	//
	// Required
	//
	// NOTE: SelectionRect.FFI_USE MUST be set to true to get SelectionRect used.
	SelectionRect SelectionRect

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FindCallbackResults with all fields set.
func (p FindCallbackResults) FromRef(ref js.Ref) FindCallbackResults {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FindCallbackResults in the application heap.
func (p FindCallbackResults) New() js.Ref {
	return bindings.FindCallbackResultsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FindCallbackResults) UpdateFrom(ref js.Ref) {
	bindings.FindCallbackResultsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FindCallbackResults) Update(ref js.Ref) {
	bindings.FindCallbackResultsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FindCallbackResults) FreeMembers(recursive bool) {
	if recursive {
		p.SelectionRect.FreeMembers(true)
	}
}

type FindOptions struct {
	// Backward is "FindOptions.backward"
	//
	// Optional
	//
	// NOTE: FFI_USE_Backward MUST be set to true to make this field effective.
	Backward bool
	// MatchCase is "FindOptions.matchCase"
	//
	// Optional
	//
	// NOTE: FFI_USE_MatchCase MUST be set to true to make this field effective.
	MatchCase bool

	FFI_USE_Backward  bool // for Backward.
	FFI_USE_MatchCase bool // for MatchCase.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FindOptions with all fields set.
func (p FindOptions) FromRef(ref js.Ref) FindOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FindOptions in the application heap.
func (p FindOptions) New() js.Ref {
	return bindings.FindOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FindOptions) UpdateFrom(ref js.Ref) {
	bindings.FindOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FindOptions) Update(ref js.Ref) {
	bindings.FindOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FindOptions) FreeMembers(recursive bool) {
}

type ForwardArgCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn ForwardArgCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ForwardArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ForwardArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *ForwardArgCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ForwardArgCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FullscreenPermissionRequest struct {
	ref js.Ref
}

func (this FullscreenPermissionRequest) Once() FullscreenPermissionRequest {
	this.ref.Once()
	return this
}

func (this FullscreenPermissionRequest) Ref() js.Ref {
	return this.ref
}

func (this FullscreenPermissionRequest) FromRef(ref js.Ref) FullscreenPermissionRequest {
	this.ref = ref
	return this
}

func (this FullscreenPermissionRequest) Free() {
	this.ref.Free()
}

// Origin returns the value of property "FullscreenPermissionRequest.origin".
//
// It returns ok=false if there is no such property.
func (this FullscreenPermissionRequest) Origin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetFullscreenPermissionRequestOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOrigin sets the value of property "FullscreenPermissionRequest.origin" to val.
//
// It returns false if the property cannot be set.
func (this FullscreenPermissionRequest) SetOrigin(val js.String) bool {
	return js.True == bindings.SetFullscreenPermissionRequestOrigin(
		this.ref,
		val.Ref(),
	)
}

// HasFuncAllow returns true if the method "FullscreenPermissionRequest.allow" exists.
func (this FullscreenPermissionRequest) HasFuncAllow() bool {
	return js.True == bindings.HasFuncFullscreenPermissionRequestAllow(
		this.ref,
	)
}

// FuncAllow returns the method "FullscreenPermissionRequest.allow".
func (this FullscreenPermissionRequest) FuncAllow() (fn js.Func[func()]) {
	bindings.FuncFullscreenPermissionRequestAllow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Allow calls the method "FullscreenPermissionRequest.allow".
func (this FullscreenPermissionRequest) Allow() (ret js.Void) {
	bindings.CallFullscreenPermissionRequestAllow(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAllow calls the method "FullscreenPermissionRequest.allow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FullscreenPermissionRequest) TryAllow() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFullscreenPermissionRequestAllow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeny returns true if the method "FullscreenPermissionRequest.deny" exists.
func (this FullscreenPermissionRequest) HasFuncDeny() bool {
	return js.True == bindings.HasFuncFullscreenPermissionRequestDeny(
		this.ref,
	)
}

// FuncDeny returns the method "FullscreenPermissionRequest.deny".
func (this FullscreenPermissionRequest) FuncDeny() (fn js.Func[func()]) {
	bindings.FuncFullscreenPermissionRequestDeny(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Deny calls the method "FullscreenPermissionRequest.deny".
func (this FullscreenPermissionRequest) Deny() (ret js.Void) {
	bindings.CallFullscreenPermissionRequestDeny(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeny calls the method "FullscreenPermissionRequest.deny"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this FullscreenPermissionRequest) TryDeny() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFullscreenPermissionRequestDeny(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type GeolocationPermissionRequest struct {
	ref js.Ref
}

func (this GeolocationPermissionRequest) Once() GeolocationPermissionRequest {
	this.ref.Once()
	return this
}

func (this GeolocationPermissionRequest) Ref() js.Ref {
	return this.ref
}

func (this GeolocationPermissionRequest) FromRef(ref js.Ref) GeolocationPermissionRequest {
	this.ref = ref
	return this
}

func (this GeolocationPermissionRequest) Free() {
	this.ref.Free()
}

// Url returns the value of property "GeolocationPermissionRequest.url".
//
// It returns ok=false if there is no such property.
func (this GeolocationPermissionRequest) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetGeolocationPermissionRequestUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUrl sets the value of property "GeolocationPermissionRequest.url" to val.
//
// It returns false if the property cannot be set.
func (this GeolocationPermissionRequest) SetUrl(val js.String) bool {
	return js.True == bindings.SetGeolocationPermissionRequestUrl(
		this.ref,
		val.Ref(),
	)
}

// HasFuncAllow returns true if the method "GeolocationPermissionRequest.allow" exists.
func (this GeolocationPermissionRequest) HasFuncAllow() bool {
	return js.True == bindings.HasFuncGeolocationPermissionRequestAllow(
		this.ref,
	)
}

// FuncAllow returns the method "GeolocationPermissionRequest.allow".
func (this GeolocationPermissionRequest) FuncAllow() (fn js.Func[func()]) {
	bindings.FuncGeolocationPermissionRequestAllow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Allow calls the method "GeolocationPermissionRequest.allow".
func (this GeolocationPermissionRequest) Allow() (ret js.Void) {
	bindings.CallGeolocationPermissionRequestAllow(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAllow calls the method "GeolocationPermissionRequest.allow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GeolocationPermissionRequest) TryAllow() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGeolocationPermissionRequestAllow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeny returns true if the method "GeolocationPermissionRequest.deny" exists.
func (this GeolocationPermissionRequest) HasFuncDeny() bool {
	return js.True == bindings.HasFuncGeolocationPermissionRequestDeny(
		this.ref,
	)
}

// FuncDeny returns the method "GeolocationPermissionRequest.deny".
func (this GeolocationPermissionRequest) FuncDeny() (fn js.Func[func()]) {
	bindings.FuncGeolocationPermissionRequestDeny(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Deny calls the method "GeolocationPermissionRequest.deny".
func (this GeolocationPermissionRequest) Deny() (ret js.Void) {
	bindings.CallGeolocationPermissionRequestDeny(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeny calls the method "GeolocationPermissionRequest.deny"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this GeolocationPermissionRequest) TryDeny() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGeolocationPermissionRequestDeny(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type GetAudioStateArgCallbackFunc func(this js.Ref, audible bool) js.Ref

func (fn GetAudioStateArgCallbackFunc) Register() js.Func[func(audible bool)] {
	return js.RegisterCallback[func(audible bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetAudioStateArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetAudioStateArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, audible bool) js.Ref
	Arg T
}

func (cb *GetAudioStateArgCallback[T]) Register() js.Func[func(audible bool)] {
	return js.RegisterCallback[func(audible bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetAudioStateArgCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetZoomArgCallbackFunc func(this js.Ref, zoomFactor float64) js.Ref

func (fn GetZoomArgCallbackFunc) Register() js.Func[func(zoomFactor float64)] {
	return js.RegisterCallback[func(zoomFactor float64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetZoomArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetZoomArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, zoomFactor float64) js.Ref
	Arg T
}

func (cb *GetZoomArgCallback[T]) Register() js.Func[func(zoomFactor float64)] {
	return js.RegisterCallback[func(zoomFactor float64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetZoomArgCallback[T]) DispatchCallback(
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

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetZoomModeArgCallbackFunc func(this js.Ref, ZoomMode ZoomMode) js.Ref

func (fn GetZoomModeArgCallbackFunc) Register() js.Func[func(ZoomMode ZoomMode)] {
	return js.RegisterCallback[func(ZoomMode ZoomMode)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GetZoomModeArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		ZoomMode(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GetZoomModeArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, ZoomMode ZoomMode) js.Ref
	Arg T
}

func (cb *GetZoomModeArgCallback[T]) Register() js.Func[func(ZoomMode ZoomMode)] {
	return js.RegisterCallback[func(ZoomMode ZoomMode)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GetZoomModeArgCallback[T]) DispatchCallback(
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

		ZoomMode(0).FromRef(args[0+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ZoomMode uint32

const (
	_ ZoomMode = iota

	ZoomMode_PER_ORIGIN
	ZoomMode_PER_VIEW
	ZoomMode_DISABLED
)

func (ZoomMode) FromRef(str js.Ref) ZoomMode {
	return ZoomMode(bindings.ConstOfZoomMode(str))
}

func (x ZoomMode) String() (string, bool) {
	switch x {
	case ZoomMode_PER_ORIGIN:
		return "per-origin", true
	case ZoomMode_PER_VIEW:
		return "per-view", true
	case ZoomMode_DISABLED:
		return "disabled", true
	default:
		return "", false
	}
}

type GoArgCallbackFunc func(this js.Ref, success bool) js.Ref

func (fn GoArgCallbackFunc) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn GoArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type GoArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, success bool) js.Ref
	Arg T
}

func (cb *GoArgCallback[T]) Register() js.Func[func(success bool)] {
	return js.RegisterCallback[func(success bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *GoArgCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type InjectDetails struct {
	// Code is "InjectDetails.code"
	//
	// Optional
	Code js.String
	// File is "InjectDetails.file"
	//
	// Optional
	File js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a InjectDetails with all fields set.
func (p InjectDetails) FromRef(ref js.Ref) InjectDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new InjectDetails in the application heap.
func (p InjectDetails) New() js.Ref {
	return bindings.InjectDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *InjectDetails) UpdateFrom(ref js.Ref) {
	bindings.InjectDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *InjectDetails) Update(ref js.Ref) {
	bindings.InjectDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *InjectDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Code.Ref(),
		p.File.Ref(),
	)
	p.Code = p.Code.FromRef(js.Undefined)
	p.File = p.File.FromRef(js.Undefined)
}

type InsertCSSArgCallbackFunc func(this js.Ref) js.Ref

func (fn InsertCSSArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn InsertCSSArgCallbackFunc) DispatchCallback(
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

type InsertCSSArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *InsertCSSArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *InsertCSSArgCallback[T]) DispatchCallback(
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

type IsAudioMutedArgCallbackFunc func(this js.Ref, muted bool) js.Ref

func (fn IsAudioMutedArgCallbackFunc) Register() js.Func[func(muted bool)] {
	return js.RegisterCallback[func(muted bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsAudioMutedArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IsAudioMutedArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, muted bool) js.Ref
	Arg T
}

func (cb *IsAudioMutedArgCallback[T]) Register() js.Func[func(muted bool)] {
	return js.RegisterCallback[func(muted bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsAudioMutedArgCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IsSpatialNavigationEnabledArgCallbackFunc func(this js.Ref, enabled bool) js.Ref

func (fn IsSpatialNavigationEnabledArgCallbackFunc) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn IsSpatialNavigationEnabledArgCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type IsSpatialNavigationEnabledArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, enabled bool) js.Ref
	Arg T
}

func (cb *IsSpatialNavigationEnabledArgCallback[T]) Register() js.Func[func(enabled bool)] {
	return js.RegisterCallback[func(enabled bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *IsSpatialNavigationEnabledArgCallback[T]) DispatchCallback(
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

		args[0+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LoadPluginPermissionRequest struct {
	ref js.Ref
}

func (this LoadPluginPermissionRequest) Once() LoadPluginPermissionRequest {
	this.ref.Once()
	return this
}

func (this LoadPluginPermissionRequest) Ref() js.Ref {
	return this.ref
}

func (this LoadPluginPermissionRequest) FromRef(ref js.Ref) LoadPluginPermissionRequest {
	this.ref = ref
	return this
}

func (this LoadPluginPermissionRequest) Free() {
	this.ref.Free()
}

// Identifier returns the value of property "LoadPluginPermissionRequest.identifier".
//
// It returns ok=false if there is no such property.
func (this LoadPluginPermissionRequest) Identifier() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLoadPluginPermissionRequestIdentifier(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIdentifier sets the value of property "LoadPluginPermissionRequest.identifier" to val.
//
// It returns false if the property cannot be set.
func (this LoadPluginPermissionRequest) SetIdentifier(val js.String) bool {
	return js.True == bindings.SetLoadPluginPermissionRequestIdentifier(
		this.ref,
		val.Ref(),
	)
}

// Name returns the value of property "LoadPluginPermissionRequest.name".
//
// It returns ok=false if there is no such property.
func (this LoadPluginPermissionRequest) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetLoadPluginPermissionRequestName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "LoadPluginPermissionRequest.name" to val.
//
// It returns false if the property cannot be set.
func (this LoadPluginPermissionRequest) SetName(val js.String) bool {
	return js.True == bindings.SetLoadPluginPermissionRequestName(
		this.ref,
		val.Ref(),
	)
}

// HasFuncAllow returns true if the method "LoadPluginPermissionRequest.allow" exists.
func (this LoadPluginPermissionRequest) HasFuncAllow() bool {
	return js.True == bindings.HasFuncLoadPluginPermissionRequestAllow(
		this.ref,
	)
}

// FuncAllow returns the method "LoadPluginPermissionRequest.allow".
func (this LoadPluginPermissionRequest) FuncAllow() (fn js.Func[func()]) {
	bindings.FuncLoadPluginPermissionRequestAllow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Allow calls the method "LoadPluginPermissionRequest.allow".
func (this LoadPluginPermissionRequest) Allow() (ret js.Void) {
	bindings.CallLoadPluginPermissionRequestAllow(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAllow calls the method "LoadPluginPermissionRequest.allow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LoadPluginPermissionRequest) TryAllow() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLoadPluginPermissionRequestAllow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeny returns true if the method "LoadPluginPermissionRequest.deny" exists.
func (this LoadPluginPermissionRequest) HasFuncDeny() bool {
	return js.True == bindings.HasFuncLoadPluginPermissionRequestDeny(
		this.ref,
	)
}

// FuncDeny returns the method "LoadPluginPermissionRequest.deny".
func (this LoadPluginPermissionRequest) FuncDeny() (fn js.Func[func()]) {
	bindings.FuncLoadPluginPermissionRequestDeny(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Deny calls the method "LoadPluginPermissionRequest.deny".
func (this LoadPluginPermissionRequest) Deny() (ret js.Void) {
	bindings.CallLoadPluginPermissionRequestDeny(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeny calls the method "LoadPluginPermissionRequest.deny"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this LoadPluginPermissionRequest) TryDeny() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLoadPluginPermissionRequestDeny(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type LoadabortArgReason uint32

const (
	_ LoadabortArgReason = iota

	LoadabortArgReason_ERR_ABORTED
	LoadabortArgReason_ERR_INVALID_URL
	LoadabortArgReason_ERR_DISALLOWED_URL_SCHEME
	LoadabortArgReason_ERR_BLOCKED_BY_CLIENT
	LoadabortArgReason_ERR_ADDRESS_UNREACHABLE
	LoadabortArgReason_ERR_EMPTY_RESPONSE
	LoadabortArgReason_ERR_FILE_NOT_FOUND
	LoadabortArgReason_ERR_UNKNOWN_URL_SCHEME
)

func (LoadabortArgReason) FromRef(str js.Ref) LoadabortArgReason {
	return LoadabortArgReason(bindings.ConstOfLoadabortArgReason(str))
}

func (x LoadabortArgReason) String() (string, bool) {
	switch x {
	case LoadabortArgReason_ERR_ABORTED:
		return "ERR_ABORTED", true
	case LoadabortArgReason_ERR_INVALID_URL:
		return "ERR_INVALID_URL", true
	case LoadabortArgReason_ERR_DISALLOWED_URL_SCHEME:
		return "ERR_DISALLOWED_URL_SCHEME", true
	case LoadabortArgReason_ERR_BLOCKED_BY_CLIENT:
		return "ERR_BLOCKED_BY_CLIENT", true
	case LoadabortArgReason_ERR_ADDRESS_UNREACHABLE:
		return "ERR_ADDRESS_UNREACHABLE", true
	case LoadabortArgReason_ERR_EMPTY_RESPONSE:
		return "ERR_EMPTY_RESPONSE", true
	case LoadabortArgReason_ERR_FILE_NOT_FOUND:
		return "ERR_FILE_NOT_FOUND", true
	case LoadabortArgReason_ERR_UNKNOWN_URL_SCHEME:
		return "ERR_UNKNOWN_URL_SCHEME", true
	default:
		return "", false
	}
}

type MediaPermissionRequest struct {
	ref js.Ref
}

func (this MediaPermissionRequest) Once() MediaPermissionRequest {
	this.ref.Once()
	return this
}

func (this MediaPermissionRequest) Ref() js.Ref {
	return this.ref
}

func (this MediaPermissionRequest) FromRef(ref js.Ref) MediaPermissionRequest {
	this.ref = ref
	return this
}

func (this MediaPermissionRequest) Free() {
	this.ref.Free()
}

// Url returns the value of property "MediaPermissionRequest.url".
//
// It returns ok=false if there is no such property.
func (this MediaPermissionRequest) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMediaPermissionRequestUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUrl sets the value of property "MediaPermissionRequest.url" to val.
//
// It returns false if the property cannot be set.
func (this MediaPermissionRequest) SetUrl(val js.String) bool {
	return js.True == bindings.SetMediaPermissionRequestUrl(
		this.ref,
		val.Ref(),
	)
}

// HasFuncAllow returns true if the method "MediaPermissionRequest.allow" exists.
func (this MediaPermissionRequest) HasFuncAllow() bool {
	return js.True == bindings.HasFuncMediaPermissionRequestAllow(
		this.ref,
	)
}

// FuncAllow returns the method "MediaPermissionRequest.allow".
func (this MediaPermissionRequest) FuncAllow() (fn js.Func[func()]) {
	bindings.FuncMediaPermissionRequestAllow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Allow calls the method "MediaPermissionRequest.allow".
func (this MediaPermissionRequest) Allow() (ret js.Void) {
	bindings.CallMediaPermissionRequestAllow(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAllow calls the method "MediaPermissionRequest.allow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaPermissionRequest) TryAllow() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaPermissionRequestAllow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeny returns true if the method "MediaPermissionRequest.deny" exists.
func (this MediaPermissionRequest) HasFuncDeny() bool {
	return js.True == bindings.HasFuncMediaPermissionRequestDeny(
		this.ref,
	)
}

// FuncDeny returns the method "MediaPermissionRequest.deny".
func (this MediaPermissionRequest) FuncDeny() (fn js.Func[func()]) {
	bindings.FuncMediaPermissionRequestDeny(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Deny calls the method "MediaPermissionRequest.deny".
func (this MediaPermissionRequest) Deny() (ret js.Void) {
	bindings.CallMediaPermissionRequestDeny(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeny calls the method "MediaPermissionRequest.deny"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this MediaPermissionRequest) TryDeny() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryMediaPermissionRequestDeny(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type NewWindow struct {
	ref js.Ref
}

func (this NewWindow) Once() NewWindow {
	this.ref.Once()
	return this
}

func (this NewWindow) Ref() js.Ref {
	return this.ref
}

func (this NewWindow) FromRef(ref js.Ref) NewWindow {
	this.ref = ref
	return this
}

func (this NewWindow) Free() {
	this.ref.Free()
}

// HasFuncAttach returns true if the method "NewWindow.attach" exists.
func (this NewWindow) HasFuncAttach() bool {
	return js.True == bindings.HasFuncNewWindowAttach(
		this.ref,
	)
}

// FuncAttach returns the method "NewWindow.attach".
func (this NewWindow) FuncAttach() (fn js.Func[func(webview AttachArgWebview)]) {
	bindings.FuncNewWindowAttach(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Attach calls the method "NewWindow.attach".
func (this NewWindow) Attach(webview AttachArgWebview) (ret js.Void) {
	bindings.CallNewWindowAttach(
		this.ref, js.Pointer(&ret),
		js.Pointer(&webview),
	)

	return
}

// TryAttach calls the method "NewWindow.attach"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NewWindow) TryAttach(webview AttachArgWebview) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNewWindowAttach(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&webview),
	)

	return
}

// HasFuncDiscard returns true if the method "NewWindow.discard" exists.
func (this NewWindow) HasFuncDiscard() bool {
	return js.True == bindings.HasFuncNewWindowDiscard(
		this.ref,
	)
}

// FuncDiscard returns the method "NewWindow.discard".
func (this NewWindow) FuncDiscard() (fn js.Func[func()]) {
	bindings.FuncNewWindowDiscard(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Discard calls the method "NewWindow.discard".
func (this NewWindow) Discard() (ret js.Void) {
	bindings.CallNewWindowDiscard(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDiscard calls the method "NewWindow.discard"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this NewWindow) TryDiscard() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNewWindowDiscard(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type NewwindowArgWindowOpenDisposition uint32

const (
	_ NewwindowArgWindowOpenDisposition = iota

	NewwindowArgWindowOpenDisposition_IGNORE
	NewwindowArgWindowOpenDisposition_SAVE_TO_DISK
	NewwindowArgWindowOpenDisposition_CURRENT_TAB
	NewwindowArgWindowOpenDisposition_NEW_BACKGROUND_TAB
	NewwindowArgWindowOpenDisposition_NEW_FOREGROUND_TAB
	NewwindowArgWindowOpenDisposition_NEW_WINDOW
	NewwindowArgWindowOpenDisposition_NEW_POPUP
)

func (NewwindowArgWindowOpenDisposition) FromRef(str js.Ref) NewwindowArgWindowOpenDisposition {
	return NewwindowArgWindowOpenDisposition(bindings.ConstOfNewwindowArgWindowOpenDisposition(str))
}

func (x NewwindowArgWindowOpenDisposition) String() (string, bool) {
	switch x {
	case NewwindowArgWindowOpenDisposition_IGNORE:
		return "ignore", true
	case NewwindowArgWindowOpenDisposition_SAVE_TO_DISK:
		return "save_to_disk", true
	case NewwindowArgWindowOpenDisposition_CURRENT_TAB:
		return "current_tab", true
	case NewwindowArgWindowOpenDisposition_NEW_BACKGROUND_TAB:
		return "new_background_tab", true
	case NewwindowArgWindowOpenDisposition_NEW_FOREGROUND_TAB:
		return "new_foreground_tab", true
	case NewwindowArgWindowOpenDisposition_NEW_WINDOW:
		return "new_window", true
	case NewwindowArgWindowOpenDisposition_NEW_POPUP:
		return "new_popup", true
	default:
		return "", false
	}
}

type PermissionrequestArgPermission uint32

const (
	_ PermissionrequestArgPermission = iota

	PermissionrequestArgPermission_MEDIA
	PermissionrequestArgPermission_GEOLOCATION
	PermissionrequestArgPermission_POINTER_LOCK
	PermissionrequestArgPermission_DOWNLOAD
	PermissionrequestArgPermission_LOADPLUGIN
	PermissionrequestArgPermission_FILESYSTEM
	PermissionrequestArgPermission_FULLSCREEN
)

func (PermissionrequestArgPermission) FromRef(str js.Ref) PermissionrequestArgPermission {
	return PermissionrequestArgPermission(bindings.ConstOfPermissionrequestArgPermission(str))
}

func (x PermissionrequestArgPermission) String() (string, bool) {
	switch x {
	case PermissionrequestArgPermission_MEDIA:
		return "media", true
	case PermissionrequestArgPermission_GEOLOCATION:
		return "geolocation", true
	case PermissionrequestArgPermission_POINTER_LOCK:
		return "pointerLock", true
	case PermissionrequestArgPermission_DOWNLOAD:
		return "download", true
	case PermissionrequestArgPermission_LOADPLUGIN:
		return "loadplugin", true
	case PermissionrequestArgPermission_FILESYSTEM:
		return "filesystem", true
	case PermissionrequestArgPermission_FULLSCREEN:
		return "fullscreen", true
	default:
		return "", false
	}
}

type PermissionrequestArgRequest struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PermissionrequestArgRequest with all fields set.
func (p PermissionrequestArgRequest) FromRef(ref js.Ref) PermissionrequestArgRequest {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PermissionrequestArgRequest in the application heap.
func (p PermissionrequestArgRequest) New() js.Ref {
	return bindings.PermissionrequestArgRequestJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *PermissionrequestArgRequest) UpdateFrom(ref js.Ref) {
	bindings.PermissionrequestArgRequestJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PermissionrequestArgRequest) Update(ref js.Ref) {
	bindings.PermissionrequestArgRequestJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PermissionrequestArgRequest) FreeMembers(recursive bool) {
}

type PointerLockPermissionRequest struct {
	ref js.Ref
}

func (this PointerLockPermissionRequest) Once() PointerLockPermissionRequest {
	this.ref.Once()
	return this
}

func (this PointerLockPermissionRequest) Ref() js.Ref {
	return this.ref
}

func (this PointerLockPermissionRequest) FromRef(ref js.Ref) PointerLockPermissionRequest {
	this.ref = ref
	return this
}

func (this PointerLockPermissionRequest) Free() {
	this.ref.Free()
}

// LastUnlockedBySelf returns the value of property "PointerLockPermissionRequest.lastUnlockedBySelf".
//
// It returns ok=false if there is no such property.
func (this PointerLockPermissionRequest) LastUnlockedBySelf() (ret bool, ok bool) {
	ok = js.True == bindings.GetPointerLockPermissionRequestLastUnlockedBySelf(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLastUnlockedBySelf sets the value of property "PointerLockPermissionRequest.lastUnlockedBySelf" to val.
//
// It returns false if the property cannot be set.
func (this PointerLockPermissionRequest) SetLastUnlockedBySelf(val bool) bool {
	return js.True == bindings.SetPointerLockPermissionRequestLastUnlockedBySelf(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Url returns the value of property "PointerLockPermissionRequest.url".
//
// It returns ok=false if there is no such property.
func (this PointerLockPermissionRequest) Url() (ret js.String, ok bool) {
	ok = js.True == bindings.GetPointerLockPermissionRequestUrl(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUrl sets the value of property "PointerLockPermissionRequest.url" to val.
//
// It returns false if the property cannot be set.
func (this PointerLockPermissionRequest) SetUrl(val js.String) bool {
	return js.True == bindings.SetPointerLockPermissionRequestUrl(
		this.ref,
		val.Ref(),
	)
}

// UserGesture returns the value of property "PointerLockPermissionRequest.userGesture".
//
// It returns ok=false if there is no such property.
func (this PointerLockPermissionRequest) UserGesture() (ret bool, ok bool) {
	ok = js.True == bindings.GetPointerLockPermissionRequestUserGesture(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUserGesture sets the value of property "PointerLockPermissionRequest.userGesture" to val.
//
// It returns false if the property cannot be set.
func (this PointerLockPermissionRequest) SetUserGesture(val bool) bool {
	return js.True == bindings.SetPointerLockPermissionRequestUserGesture(
		this.ref,
		js.Bool(bool(val)),
	)
}

// HasFuncAllow returns true if the method "PointerLockPermissionRequest.allow" exists.
func (this PointerLockPermissionRequest) HasFuncAllow() bool {
	return js.True == bindings.HasFuncPointerLockPermissionRequestAllow(
		this.ref,
	)
}

// FuncAllow returns the method "PointerLockPermissionRequest.allow".
func (this PointerLockPermissionRequest) FuncAllow() (fn js.Func[func()]) {
	bindings.FuncPointerLockPermissionRequestAllow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Allow calls the method "PointerLockPermissionRequest.allow".
func (this PointerLockPermissionRequest) Allow() (ret js.Void) {
	bindings.CallPointerLockPermissionRequestAllow(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAllow calls the method "PointerLockPermissionRequest.allow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PointerLockPermissionRequest) TryAllow() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPointerLockPermissionRequestAllow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeny returns true if the method "PointerLockPermissionRequest.deny" exists.
func (this PointerLockPermissionRequest) HasFuncDeny() bool {
	return js.True == bindings.HasFuncPointerLockPermissionRequestDeny(
		this.ref,
	)
}

// FuncDeny returns the method "PointerLockPermissionRequest.deny".
func (this PointerLockPermissionRequest) FuncDeny() (fn js.Func[func()]) {
	bindings.FuncPointerLockPermissionRequestDeny(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Deny calls the method "PointerLockPermissionRequest.deny".
func (this PointerLockPermissionRequest) Deny() (ret js.Void) {
	bindings.CallPointerLockPermissionRequestDeny(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeny calls the method "PointerLockPermissionRequest.deny"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this PointerLockPermissionRequest) TryDeny() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPointerLockPermissionRequestDeny(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SetZoomArgCallbackFunc func(this js.Ref) js.Ref

func (fn SetZoomArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetZoomArgCallbackFunc) DispatchCallback(
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

type SetZoomArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SetZoomArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetZoomArgCallback[T]) DispatchCallback(
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

type SetZoomModeArgCallbackFunc func(this js.Ref) js.Ref

func (fn SetZoomModeArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetZoomModeArgCallbackFunc) DispatchCallback(
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

type SetZoomModeArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *SetZoomModeArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetZoomModeArgCallback[T]) DispatchCallback(
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

type StopFindingArgAction uint32

const (
	_ StopFindingArgAction = iota

	StopFindingArgAction_CLEAR
	StopFindingArgAction_KEEP
	StopFindingArgAction_ACTIVATE
)

func (StopFindingArgAction) FromRef(str js.Ref) StopFindingArgAction {
	return StopFindingArgAction(bindings.ConstOfStopFindingArgAction(str))
}

func (x StopFindingArgAction) String() (string, bool) {
	switch x {
	case StopFindingArgAction_CLEAR:
		return "clear", true
	case StopFindingArgAction_KEEP:
		return "keep", true
	case StopFindingArgAction_ACTIVATE:
		return "activate", true
	default:
		return "", false
	}
}

type WebRequestEventInterface struct {
	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a WebRequestEventInterface with all fields set.
func (p WebRequestEventInterface) FromRef(ref js.Ref) WebRequestEventInterface {
	p.UpdateFrom(ref)
	return p
}

// New creates a new WebRequestEventInterface in the application heap.
func (p WebRequestEventInterface) New() js.Ref {
	return bindings.WebRequestEventInterfaceJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *WebRequestEventInterface) UpdateFrom(ref js.Ref) {
	bindings.WebRequestEventInterfaceJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *WebRequestEventInterface) Update(ref js.Ref) {
	bindings.WebRequestEventInterfaceJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *WebRequestEventInterface) FreeMembers(recursive bool) {
}

// HasFuncAddContentScripts returns true if the function "WEBEXT.webviewTag.addContentScripts" exists.
func HasFuncAddContentScripts() bool {
	return js.True == bindings.HasFuncAddContentScripts()
}

// FuncAddContentScripts returns the function "WEBEXT.webviewTag.addContentScripts".
func FuncAddContentScripts() (fn js.Func[func(contentScriptList js.Array[ContentScriptDetails])]) {
	bindings.FuncAddContentScripts(
		js.Pointer(&fn),
	)
	return
}

// AddContentScripts calls the function "WEBEXT.webviewTag.addContentScripts" directly.
func AddContentScripts(contentScriptList js.Array[ContentScriptDetails]) (ret js.Void) {
	bindings.CallAddContentScripts(
		js.Pointer(&ret),
		contentScriptList.Ref(),
	)

	return
}

// TryAddContentScripts calls the function "WEBEXT.webviewTag.addContentScripts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddContentScripts(contentScriptList js.Array[ContentScriptDetails]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddContentScripts(
		js.Pointer(&ret), js.Pointer(&exception),
		contentScriptList.Ref(),
	)

	return
}

// HasFuncBack returns true if the function "WEBEXT.webviewTag.back" exists.
func HasFuncBack() bool {
	return js.True == bindings.HasFuncBack()
}

// FuncBack returns the function "WEBEXT.webviewTag.back".
func FuncBack() (fn js.Func[func(callback js.Func[func(success bool)])]) {
	bindings.FuncBack(
		js.Pointer(&fn),
	)
	return
}

// Back calls the function "WEBEXT.webviewTag.back" directly.
func Back(callback js.Func[func(success bool)]) (ret js.Void) {
	bindings.CallBack(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryBack calls the function "WEBEXT.webviewTag.back"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryBack(callback js.Func[func(success bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryBack(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncCanGoBack returns true if the function "WEBEXT.webviewTag.canGoBack" exists.
func HasFuncCanGoBack() bool {
	return js.True == bindings.HasFuncCanGoBack()
}

// FuncCanGoBack returns the function "WEBEXT.webviewTag.canGoBack".
func FuncCanGoBack() (fn js.Func[func() bool]) {
	bindings.FuncCanGoBack(
		js.Pointer(&fn),
	)
	return
}

// CanGoBack calls the function "WEBEXT.webviewTag.canGoBack" directly.
func CanGoBack() (ret bool) {
	bindings.CallCanGoBack(
		js.Pointer(&ret),
	)

	return
}

// TryCanGoBack calls the function "WEBEXT.webviewTag.canGoBack"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCanGoBack() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanGoBack(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCanGoForward returns true if the function "WEBEXT.webviewTag.canGoForward" exists.
func HasFuncCanGoForward() bool {
	return js.True == bindings.HasFuncCanGoForward()
}

// FuncCanGoForward returns the function "WEBEXT.webviewTag.canGoForward".
func FuncCanGoForward() (fn js.Func[func() bool]) {
	bindings.FuncCanGoForward(
		js.Pointer(&fn),
	)
	return
}

// CanGoForward calls the function "WEBEXT.webviewTag.canGoForward" directly.
func CanGoForward() (ret bool) {
	bindings.CallCanGoForward(
		js.Pointer(&ret),
	)

	return
}

// TryCanGoForward calls the function "WEBEXT.webviewTag.canGoForward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCanGoForward() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanGoForward(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCaptureVisibleRegion returns true if the function "WEBEXT.webviewTag.captureVisibleRegion" exists.
func HasFuncCaptureVisibleRegion() bool {
	return js.True == bindings.HasFuncCaptureVisibleRegion()
}

// FuncCaptureVisibleRegion returns the function "WEBEXT.webviewTag.captureVisibleRegion".
func FuncCaptureVisibleRegion() (fn js.Func[func(options extensiontypes.ImageDetails, callback js.Func[func(dataUrl js.String)])]) {
	bindings.FuncCaptureVisibleRegion(
		js.Pointer(&fn),
	)
	return
}

// CaptureVisibleRegion calls the function "WEBEXT.webviewTag.captureVisibleRegion" directly.
func CaptureVisibleRegion(options extensiontypes.ImageDetails, callback js.Func[func(dataUrl js.String)]) (ret js.Void) {
	bindings.CallCaptureVisibleRegion(
		js.Pointer(&ret),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TryCaptureVisibleRegion calls the function "WEBEXT.webviewTag.captureVisibleRegion"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCaptureVisibleRegion(options extensiontypes.ImageDetails, callback js.Func[func(dataUrl js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCaptureVisibleRegion(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// HasFuncClearData returns true if the function "WEBEXT.webviewTag.clearData" exists.
func HasFuncClearData() bool {
	return js.True == bindings.HasFuncClearData()
}

// FuncClearData returns the function "WEBEXT.webviewTag.clearData".
func FuncClearData() (fn js.Func[func(options ClearDataOptions, types ClearDataTypeSet, callback js.Func[func()])]) {
	bindings.FuncClearData(
		js.Pointer(&fn),
	)
	return
}

// ClearData calls the function "WEBEXT.webviewTag.clearData" directly.
func ClearData(options ClearDataOptions, types ClearDataTypeSet, callback js.Func[func()]) (ret js.Void) {
	bindings.CallClearData(
		js.Pointer(&ret),
		js.Pointer(&options),
		js.Pointer(&types),
		callback.Ref(),
	)

	return
}

// TryClearData calls the function "WEBEXT.webviewTag.clearData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearData(options ClearDataOptions, types ClearDataTypeSet, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearData(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
		js.Pointer(&types),
		callback.Ref(),
	)

	return
}

type CloseEventCallbackFunc func(this js.Ref) js.Ref

func (fn CloseEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn CloseEventCallbackFunc) DispatchCallback(
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

type CloseEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *CloseEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *CloseEventCallback[T]) DispatchCallback(
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

// HasFuncOnClose returns true if the function "WEBEXT.webviewTag.close.addListener" exists.
func HasFuncOnClose() bool {
	return js.True == bindings.HasFuncOnClose()
}

// FuncOnClose returns the function "WEBEXT.webviewTag.close.addListener".
func FuncOnClose() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnClose(
		js.Pointer(&fn),
	)
	return
}

// OnClose calls the function "WEBEXT.webviewTag.close.addListener" directly.
func OnClose(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnClose(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnClose calls the function "WEBEXT.webviewTag.close.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnClose(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnClose(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffClose returns true if the function "WEBEXT.webviewTag.close.removeListener" exists.
func HasFuncOffClose() bool {
	return js.True == bindings.HasFuncOffClose()
}

// FuncOffClose returns the function "WEBEXT.webviewTag.close.removeListener".
func FuncOffClose() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffClose(
		js.Pointer(&fn),
	)
	return
}

// OffClose calls the function "WEBEXT.webviewTag.close.removeListener" directly.
func OffClose(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffClose(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffClose calls the function "WEBEXT.webviewTag.close.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffClose(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffClose(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnClose returns true if the function "WEBEXT.webviewTag.close.hasListener" exists.
func HasFuncHasOnClose() bool {
	return js.True == bindings.HasFuncHasOnClose()
}

// FuncHasOnClose returns the function "WEBEXT.webviewTag.close.hasListener".
func FuncHasOnClose() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnClose(
		js.Pointer(&fn),
	)
	return
}

// HasOnClose calls the function "WEBEXT.webviewTag.close.hasListener" directly.
func HasOnClose(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnClose(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnClose calls the function "WEBEXT.webviewTag.close.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnClose(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnClose(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type ConsolemessageEventCallbackFunc func(this js.Ref, level int64, message js.String, line int64, sourceId js.String) js.Ref

func (fn ConsolemessageEventCallbackFunc) Register() js.Func[func(level int64, message js.String, line int64, sourceId js.String)] {
	return js.RegisterCallback[func(level int64, message js.String, line int64, sourceId js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ConsolemessageEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		js.String{}.FromRef(args[1+1]),
		js.BigInt[int64]{}.FromRef(args[2+1]).Get(),
		js.String{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ConsolemessageEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, level int64, message js.String, line int64, sourceId js.String) js.Ref
	Arg T
}

func (cb *ConsolemessageEventCallback[T]) Register() js.Func[func(level int64, message js.String, line int64, sourceId js.String)] {
	return js.RegisterCallback[func(level int64, message js.String, line int64, sourceId js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ConsolemessageEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		js.String{}.FromRef(args[1+1]),
		js.BigInt[int64]{}.FromRef(args[2+1]).Get(),
		js.String{}.FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnConsolemessage returns true if the function "WEBEXT.webviewTag.consolemessage.addListener" exists.
func HasFuncOnConsolemessage() bool {
	return js.True == bindings.HasFuncOnConsolemessage()
}

// FuncOnConsolemessage returns the function "WEBEXT.webviewTag.consolemessage.addListener".
func FuncOnConsolemessage() (fn js.Func[func(callback js.Func[func(level int64, message js.String, line int64, sourceId js.String)])]) {
	bindings.FuncOnConsolemessage(
		js.Pointer(&fn),
	)
	return
}

// OnConsolemessage calls the function "WEBEXT.webviewTag.consolemessage.addListener" directly.
func OnConsolemessage(callback js.Func[func(level int64, message js.String, line int64, sourceId js.String)]) (ret js.Void) {
	bindings.CallOnConsolemessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnConsolemessage calls the function "WEBEXT.webviewTag.consolemessage.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnConsolemessage(callback js.Func[func(level int64, message js.String, line int64, sourceId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnConsolemessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffConsolemessage returns true if the function "WEBEXT.webviewTag.consolemessage.removeListener" exists.
func HasFuncOffConsolemessage() bool {
	return js.True == bindings.HasFuncOffConsolemessage()
}

// FuncOffConsolemessage returns the function "WEBEXT.webviewTag.consolemessage.removeListener".
func FuncOffConsolemessage() (fn js.Func[func(callback js.Func[func(level int64, message js.String, line int64, sourceId js.String)])]) {
	bindings.FuncOffConsolemessage(
		js.Pointer(&fn),
	)
	return
}

// OffConsolemessage calls the function "WEBEXT.webviewTag.consolemessage.removeListener" directly.
func OffConsolemessage(callback js.Func[func(level int64, message js.String, line int64, sourceId js.String)]) (ret js.Void) {
	bindings.CallOffConsolemessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffConsolemessage calls the function "WEBEXT.webviewTag.consolemessage.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffConsolemessage(callback js.Func[func(level int64, message js.String, line int64, sourceId js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffConsolemessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnConsolemessage returns true if the function "WEBEXT.webviewTag.consolemessage.hasListener" exists.
func HasFuncHasOnConsolemessage() bool {
	return js.True == bindings.HasFuncHasOnConsolemessage()
}

// FuncHasOnConsolemessage returns the function "WEBEXT.webviewTag.consolemessage.hasListener".
func FuncHasOnConsolemessage() (fn js.Func[func(callback js.Func[func(level int64, message js.String, line int64, sourceId js.String)]) bool]) {
	bindings.FuncHasOnConsolemessage(
		js.Pointer(&fn),
	)
	return
}

// HasOnConsolemessage calls the function "WEBEXT.webviewTag.consolemessage.hasListener" directly.
func HasOnConsolemessage(callback js.Func[func(level int64, message js.String, line int64, sourceId js.String)]) (ret bool) {
	bindings.CallHasOnConsolemessage(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnConsolemessage calls the function "WEBEXT.webviewTag.consolemessage.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnConsolemessage(callback js.Func[func(level int64, message js.String, line int64, sourceId js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnConsolemessage(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// ContentWindow returns the value of property "WEBEXT.webviewTag.contentWindow".
//
// The returned bool will be false if there is no such property.
func ContentWindow() (ret ContentWindowType, ok bool) {
	ok = js.True == bindings.GetContentWindow(
		js.Pointer(&ret),
	)

	return
}

// SetContentWindow sets the value of property "WEBEXT.webviewTag.contentWindow" to val.
//
// It returns false if the property cannot be set.
func SetContentWindow(val ContentWindowType) bool {
	return js.True == bindings.SetContentWindow(
		val.Ref())
}

type ContentloadEventCallbackFunc func(this js.Ref) js.Ref

func (fn ContentloadEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ContentloadEventCallbackFunc) DispatchCallback(
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

type ContentloadEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *ContentloadEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ContentloadEventCallback[T]) DispatchCallback(
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

// HasFuncOnContentload returns true if the function "WEBEXT.webviewTag.contentload.addListener" exists.
func HasFuncOnContentload() bool {
	return js.True == bindings.HasFuncOnContentload()
}

// FuncOnContentload returns the function "WEBEXT.webviewTag.contentload.addListener".
func FuncOnContentload() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnContentload(
		js.Pointer(&fn),
	)
	return
}

// OnContentload calls the function "WEBEXT.webviewTag.contentload.addListener" directly.
func OnContentload(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnContentload(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnContentload calls the function "WEBEXT.webviewTag.contentload.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnContentload(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnContentload(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffContentload returns true if the function "WEBEXT.webviewTag.contentload.removeListener" exists.
func HasFuncOffContentload() bool {
	return js.True == bindings.HasFuncOffContentload()
}

// FuncOffContentload returns the function "WEBEXT.webviewTag.contentload.removeListener".
func FuncOffContentload() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffContentload(
		js.Pointer(&fn),
	)
	return
}

// OffContentload calls the function "WEBEXT.webviewTag.contentload.removeListener" directly.
func OffContentload(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffContentload(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffContentload calls the function "WEBEXT.webviewTag.contentload.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffContentload(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffContentload(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnContentload returns true if the function "WEBEXT.webviewTag.contentload.hasListener" exists.
func HasFuncHasOnContentload() bool {
	return js.True == bindings.HasFuncHasOnContentload()
}

// FuncHasOnContentload returns the function "WEBEXT.webviewTag.contentload.hasListener".
func FuncHasOnContentload() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnContentload(
		js.Pointer(&fn),
	)
	return
}

// HasOnContentload calls the function "WEBEXT.webviewTag.contentload.hasListener" directly.
func HasOnContentload(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnContentload(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnContentload calls the function "WEBEXT.webviewTag.contentload.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnContentload(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnContentload(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// ContextMenus returns the value of property "WEBEXT.webviewTag.contextMenus".
//
// The returned bool will be false if there is no such property.
func ContextMenus() (ret ContextMenusType, ok bool) {
	ok = js.True == bindings.GetContextMenus(
		js.Pointer(&ret),
	)

	return
}

// SetContextMenus sets the value of property "WEBEXT.webviewTag.contextMenus" to val.
//
// It returns false if the property cannot be set.
func SetContextMenus(val ContextMenusType) bool {
	return js.True == bindings.SetContextMenus(
		val.Ref())
}

type DialogEventCallbackFunc func(this js.Ref, messageType DialogArgMessageType, messageText js.String, dialog DialogController) js.Ref

func (fn DialogEventCallbackFunc) Register() js.Func[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)] {
	return js.RegisterCallback[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn DialogEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		DialogArgMessageType(0).FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		DialogController{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type DialogEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, messageType DialogArgMessageType, messageText js.String, dialog DialogController) js.Ref
	Arg T
}

func (cb *DialogEventCallback[T]) Register() js.Func[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)] {
	return js.RegisterCallback[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *DialogEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		DialogArgMessageType(0).FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		DialogController{}.FromRef(args[2+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnDialog returns true if the function "WEBEXT.webviewTag.dialog.addListener" exists.
func HasFuncOnDialog() bool {
	return js.True == bindings.HasFuncOnDialog()
}

// FuncOnDialog returns the function "WEBEXT.webviewTag.dialog.addListener".
func FuncOnDialog() (fn js.Func[func(callback js.Func[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)])]) {
	bindings.FuncOnDialog(
		js.Pointer(&fn),
	)
	return
}

// OnDialog calls the function "WEBEXT.webviewTag.dialog.addListener" directly.
func OnDialog(callback js.Func[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)]) (ret js.Void) {
	bindings.CallOnDialog(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDialog calls the function "WEBEXT.webviewTag.dialog.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDialog(callback js.Func[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDialog(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDialog returns true if the function "WEBEXT.webviewTag.dialog.removeListener" exists.
func HasFuncOffDialog() bool {
	return js.True == bindings.HasFuncOffDialog()
}

// FuncOffDialog returns the function "WEBEXT.webviewTag.dialog.removeListener".
func FuncOffDialog() (fn js.Func[func(callback js.Func[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)])]) {
	bindings.FuncOffDialog(
		js.Pointer(&fn),
	)
	return
}

// OffDialog calls the function "WEBEXT.webviewTag.dialog.removeListener" directly.
func OffDialog(callback js.Func[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)]) (ret js.Void) {
	bindings.CallOffDialog(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDialog calls the function "WEBEXT.webviewTag.dialog.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDialog(callback js.Func[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDialog(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDialog returns true if the function "WEBEXT.webviewTag.dialog.hasListener" exists.
func HasFuncHasOnDialog() bool {
	return js.True == bindings.HasFuncHasOnDialog()
}

// FuncHasOnDialog returns the function "WEBEXT.webviewTag.dialog.hasListener".
func FuncHasOnDialog() (fn js.Func[func(callback js.Func[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)]) bool]) {
	bindings.FuncHasOnDialog(
		js.Pointer(&fn),
	)
	return
}

// HasOnDialog calls the function "WEBEXT.webviewTag.dialog.hasListener" directly.
func HasOnDialog(callback js.Func[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)]) (ret bool) {
	bindings.CallHasOnDialog(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDialog calls the function "WEBEXT.webviewTag.dialog.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDialog(callback js.Func[func(messageType DialogArgMessageType, messageText js.String, dialog DialogController)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDialog(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncExecuteScript returns true if the function "WEBEXT.webviewTag.executeScript" exists.
func HasFuncExecuteScript() bool {
	return js.True == bindings.HasFuncExecuteScript()
}

// FuncExecuteScript returns the function "WEBEXT.webviewTag.executeScript".
func FuncExecuteScript() (fn js.Func[func(details InjectDetails, callback js.Func[func(result js.Array[js.Any])])]) {
	bindings.FuncExecuteScript(
		js.Pointer(&fn),
	)
	return
}

// ExecuteScript calls the function "WEBEXT.webviewTag.executeScript" directly.
func ExecuteScript(details InjectDetails, callback js.Func[func(result js.Array[js.Any])]) (ret js.Void) {
	bindings.CallExecuteScript(
		js.Pointer(&ret),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// TryExecuteScript calls the function "WEBEXT.webviewTag.executeScript"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExecuteScript(details InjectDetails, callback js.Func[func(result js.Array[js.Any])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryExecuteScript(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

type ExitEventCallbackFunc func(this js.Ref, processID int64, reason ExitArgReason) js.Ref

func (fn ExitEventCallbackFunc) Register() js.Func[func(processID int64, reason ExitArgReason)] {
	return js.RegisterCallback[func(processID int64, reason ExitArgReason)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ExitEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		ExitArgReason(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ExitEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, processID int64, reason ExitArgReason) js.Ref
	Arg T
}

func (cb *ExitEventCallback[T]) Register() js.Func[func(processID int64, reason ExitArgReason)] {
	return js.RegisterCallback[func(processID int64, reason ExitArgReason)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ExitEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
		ExitArgReason(0).FromRef(args[1+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnExit returns true if the function "WEBEXT.webviewTag.exit.addListener" exists.
func HasFuncOnExit() bool {
	return js.True == bindings.HasFuncOnExit()
}

// FuncOnExit returns the function "WEBEXT.webviewTag.exit.addListener".
func FuncOnExit() (fn js.Func[func(callback js.Func[func(processID int64, reason ExitArgReason)])]) {
	bindings.FuncOnExit(
		js.Pointer(&fn),
	)
	return
}

// OnExit calls the function "WEBEXT.webviewTag.exit.addListener" directly.
func OnExit(callback js.Func[func(processID int64, reason ExitArgReason)]) (ret js.Void) {
	bindings.CallOnExit(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnExit calls the function "WEBEXT.webviewTag.exit.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnExit(callback js.Func[func(processID int64, reason ExitArgReason)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnExit(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffExit returns true if the function "WEBEXT.webviewTag.exit.removeListener" exists.
func HasFuncOffExit() bool {
	return js.True == bindings.HasFuncOffExit()
}

// FuncOffExit returns the function "WEBEXT.webviewTag.exit.removeListener".
func FuncOffExit() (fn js.Func[func(callback js.Func[func(processID int64, reason ExitArgReason)])]) {
	bindings.FuncOffExit(
		js.Pointer(&fn),
	)
	return
}

// OffExit calls the function "WEBEXT.webviewTag.exit.removeListener" directly.
func OffExit(callback js.Func[func(processID int64, reason ExitArgReason)]) (ret js.Void) {
	bindings.CallOffExit(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffExit calls the function "WEBEXT.webviewTag.exit.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffExit(callback js.Func[func(processID int64, reason ExitArgReason)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffExit(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnExit returns true if the function "WEBEXT.webviewTag.exit.hasListener" exists.
func HasFuncHasOnExit() bool {
	return js.True == bindings.HasFuncHasOnExit()
}

// FuncHasOnExit returns the function "WEBEXT.webviewTag.exit.hasListener".
func FuncHasOnExit() (fn js.Func[func(callback js.Func[func(processID int64, reason ExitArgReason)]) bool]) {
	bindings.FuncHasOnExit(
		js.Pointer(&fn),
	)
	return
}

// HasOnExit calls the function "WEBEXT.webviewTag.exit.hasListener" directly.
func HasOnExit(callback js.Func[func(processID int64, reason ExitArgReason)]) (ret bool) {
	bindings.CallHasOnExit(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnExit calls the function "WEBEXT.webviewTag.exit.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnExit(callback js.Func[func(processID int64, reason ExitArgReason)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnExit(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncFind returns true if the function "WEBEXT.webviewTag.find" exists.
func HasFuncFind() bool {
	return js.True == bindings.HasFuncFind()
}

// FuncFind returns the function "WEBEXT.webviewTag.find".
func FuncFind() (fn js.Func[func(searchText js.String, options FindOptions, callback js.Func[func(results *FindCallbackResults)])]) {
	bindings.FuncFind(
		js.Pointer(&fn),
	)
	return
}

// Find calls the function "WEBEXT.webviewTag.find" directly.
func Find(searchText js.String, options FindOptions, callback js.Func[func(results *FindCallbackResults)]) (ret js.Void) {
	bindings.CallFind(
		js.Pointer(&ret),
		searchText.Ref(),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TryFind calls the function "WEBEXT.webviewTag.find"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFind(searchText js.String, options FindOptions, callback js.Func[func(results *FindCallbackResults)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFind(
		js.Pointer(&ret), js.Pointer(&exception),
		searchText.Ref(),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

type FindupdateEventCallbackFunc func(this js.Ref, searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String) js.Ref

func (fn FindupdateEventCallbackFunc) Register() js.Func[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)] {
	return js.RegisterCallback[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn FindupdateEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 6+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg3 SelectionRect
	arg3.UpdateFrom(args[3+1])
	defer arg3.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.BigInt[int64]{}.FromRef(args[1+1]).Get(),
		js.BigInt[int64]{}.FromRef(args[2+1]).Get(),
		mark.NoEscape(&arg3),
		args[4+1] == js.True,
		js.String{}.FromRef(args[5+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type FindupdateEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String) js.Ref
	Arg T
}

func (cb *FindupdateEventCallback[T]) Register() js.Func[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)] {
	return js.RegisterCallback[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *FindupdateEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 6+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg3 SelectionRect
	arg3.UpdateFrom(args[3+1])
	defer arg3.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.BigInt[int64]{}.FromRef(args[1+1]).Get(),
		js.BigInt[int64]{}.FromRef(args[2+1]).Get(),
		mark.NoEscape(&arg3),
		args[4+1] == js.True,
		js.String{}.FromRef(args[5+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnFindupdate returns true if the function "WEBEXT.webviewTag.findupdate.addListener" exists.
func HasFuncOnFindupdate() bool {
	return js.True == bindings.HasFuncOnFindupdate()
}

// FuncOnFindupdate returns the function "WEBEXT.webviewTag.findupdate.addListener".
func FuncOnFindupdate() (fn js.Func[func(callback js.Func[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)])]) {
	bindings.FuncOnFindupdate(
		js.Pointer(&fn),
	)
	return
}

// OnFindupdate calls the function "WEBEXT.webviewTag.findupdate.addListener" directly.
func OnFindupdate(callback js.Func[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)]) (ret js.Void) {
	bindings.CallOnFindupdate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnFindupdate calls the function "WEBEXT.webviewTag.findupdate.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnFindupdate(callback js.Func[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnFindupdate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffFindupdate returns true if the function "WEBEXT.webviewTag.findupdate.removeListener" exists.
func HasFuncOffFindupdate() bool {
	return js.True == bindings.HasFuncOffFindupdate()
}

// FuncOffFindupdate returns the function "WEBEXT.webviewTag.findupdate.removeListener".
func FuncOffFindupdate() (fn js.Func[func(callback js.Func[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)])]) {
	bindings.FuncOffFindupdate(
		js.Pointer(&fn),
	)
	return
}

// OffFindupdate calls the function "WEBEXT.webviewTag.findupdate.removeListener" directly.
func OffFindupdate(callback js.Func[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)]) (ret js.Void) {
	bindings.CallOffFindupdate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffFindupdate calls the function "WEBEXT.webviewTag.findupdate.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffFindupdate(callback js.Func[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffFindupdate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnFindupdate returns true if the function "WEBEXT.webviewTag.findupdate.hasListener" exists.
func HasFuncHasOnFindupdate() bool {
	return js.True == bindings.HasFuncHasOnFindupdate()
}

// FuncHasOnFindupdate returns the function "WEBEXT.webviewTag.findupdate.hasListener".
func FuncHasOnFindupdate() (fn js.Func[func(callback js.Func[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)]) bool]) {
	bindings.FuncHasOnFindupdate(
		js.Pointer(&fn),
	)
	return
}

// HasOnFindupdate calls the function "WEBEXT.webviewTag.findupdate.hasListener" directly.
func HasOnFindupdate(callback js.Func[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)]) (ret bool) {
	bindings.CallHasOnFindupdate(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnFindupdate calls the function "WEBEXT.webviewTag.findupdate.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnFindupdate(callback js.Func[func(searchText js.String, numberOfMatches int64, activeMatchOrdinal int64, selectionRect *SelectionRect, canceled bool, finalUpdate js.String)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnFindupdate(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncForward returns true if the function "WEBEXT.webviewTag.forward" exists.
func HasFuncForward() bool {
	return js.True == bindings.HasFuncForward()
}

// FuncForward returns the function "WEBEXT.webviewTag.forward".
func FuncForward() (fn js.Func[func(callback js.Func[func(success bool)])]) {
	bindings.FuncForward(
		js.Pointer(&fn),
	)
	return
}

// Forward calls the function "WEBEXT.webviewTag.forward" directly.
func Forward(callback js.Func[func(success bool)]) (ret js.Void) {
	bindings.CallForward(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryForward calls the function "WEBEXT.webviewTag.forward"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryForward(callback js.Func[func(success bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryForward(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetAudioState returns true if the function "WEBEXT.webviewTag.getAudioState" exists.
func HasFuncGetAudioState() bool {
	return js.True == bindings.HasFuncGetAudioState()
}

// FuncGetAudioState returns the function "WEBEXT.webviewTag.getAudioState".
func FuncGetAudioState() (fn js.Func[func(callback js.Func[func(audible bool)])]) {
	bindings.FuncGetAudioState(
		js.Pointer(&fn),
	)
	return
}

// GetAudioState calls the function "WEBEXT.webviewTag.getAudioState" directly.
func GetAudioState(callback js.Func[func(audible bool)]) (ret js.Void) {
	bindings.CallGetAudioState(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetAudioState calls the function "WEBEXT.webviewTag.getAudioState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAudioState(callback js.Func[func(audible bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAudioState(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetProcessId returns true if the function "WEBEXT.webviewTag.getProcessId" exists.
func HasFuncGetProcessId() bool {
	return js.True == bindings.HasFuncGetProcessId()
}

// FuncGetProcessId returns the function "WEBEXT.webviewTag.getProcessId".
func FuncGetProcessId() (fn js.Func[func() int64]) {
	bindings.FuncGetProcessId(
		js.Pointer(&fn),
	)
	return
}

// GetProcessId calls the function "WEBEXT.webviewTag.getProcessId" directly.
func GetProcessId() (ret int64) {
	bindings.CallGetProcessId(
		js.Pointer(&ret),
	)

	return
}

// TryGetProcessId calls the function "WEBEXT.webviewTag.getProcessId"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetProcessId() (ret int64, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetProcessId(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetUserAgent returns true if the function "WEBEXT.webviewTag.getUserAgent" exists.
func HasFuncGetUserAgent() bool {
	return js.True == bindings.HasFuncGetUserAgent()
}

// FuncGetUserAgent returns the function "WEBEXT.webviewTag.getUserAgent".
func FuncGetUserAgent() (fn js.Func[func() js.String]) {
	bindings.FuncGetUserAgent(
		js.Pointer(&fn),
	)
	return
}

// GetUserAgent calls the function "WEBEXT.webviewTag.getUserAgent" directly.
func GetUserAgent() (ret js.String) {
	bindings.CallGetUserAgent(
		js.Pointer(&ret),
	)

	return
}

// TryGetUserAgent calls the function "WEBEXT.webviewTag.getUserAgent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetUserAgent() (ret js.String, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetUserAgent(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetZoom returns true if the function "WEBEXT.webviewTag.getZoom" exists.
func HasFuncGetZoom() bool {
	return js.True == bindings.HasFuncGetZoom()
}

// FuncGetZoom returns the function "WEBEXT.webviewTag.getZoom".
func FuncGetZoom() (fn js.Func[func(callback js.Func[func(zoomFactor float64)])]) {
	bindings.FuncGetZoom(
		js.Pointer(&fn),
	)
	return
}

// GetZoom calls the function "WEBEXT.webviewTag.getZoom" directly.
func GetZoom(callback js.Func[func(zoomFactor float64)]) (ret js.Void) {
	bindings.CallGetZoom(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetZoom calls the function "WEBEXT.webviewTag.getZoom"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetZoom(callback js.Func[func(zoomFactor float64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetZoom(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGetZoomMode returns true if the function "WEBEXT.webviewTag.getZoomMode" exists.
func HasFuncGetZoomMode() bool {
	return js.True == bindings.HasFuncGetZoomMode()
}

// FuncGetZoomMode returns the function "WEBEXT.webviewTag.getZoomMode".
func FuncGetZoomMode() (fn js.Func[func(callback js.Func[func(ZoomMode ZoomMode)])]) {
	bindings.FuncGetZoomMode(
		js.Pointer(&fn),
	)
	return
}

// GetZoomMode calls the function "WEBEXT.webviewTag.getZoomMode" directly.
func GetZoomMode(callback js.Func[func(ZoomMode ZoomMode)]) (ret js.Void) {
	bindings.CallGetZoomMode(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryGetZoomMode calls the function "WEBEXT.webviewTag.getZoomMode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetZoomMode(callback js.Func[func(ZoomMode ZoomMode)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetZoomMode(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncGo returns true if the function "WEBEXT.webviewTag.go" exists.
func HasFuncGo() bool {
	return js.True == bindings.HasFuncGo()
}

// FuncGo returns the function "WEBEXT.webviewTag.go".
func FuncGo() (fn js.Func[func(relativeIndex int64, callback js.Func[func(success bool)])]) {
	bindings.FuncGo(
		js.Pointer(&fn),
	)
	return
}

// Go calls the function "WEBEXT.webviewTag.go" directly.
func Go(relativeIndex int64, callback js.Func[func(success bool)]) (ret js.Void) {
	bindings.CallGo(
		js.Pointer(&ret),
		float64(relativeIndex),
		callback.Ref(),
	)

	return
}

// TryGo calls the function "WEBEXT.webviewTag.go"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGo(relativeIndex int64, callback js.Func[func(success bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGo(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(relativeIndex),
		callback.Ref(),
	)

	return
}

// HasFuncInsertCSS returns true if the function "WEBEXT.webviewTag.insertCSS" exists.
func HasFuncInsertCSS() bool {
	return js.True == bindings.HasFuncInsertCSS()
}

// FuncInsertCSS returns the function "WEBEXT.webviewTag.insertCSS".
func FuncInsertCSS() (fn js.Func[func(details InjectDetails, callback js.Func[func()])]) {
	bindings.FuncInsertCSS(
		js.Pointer(&fn),
	)
	return
}

// InsertCSS calls the function "WEBEXT.webviewTag.insertCSS" directly.
func InsertCSS(details InjectDetails, callback js.Func[func()]) (ret js.Void) {
	bindings.CallInsertCSS(
		js.Pointer(&ret),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// TryInsertCSS calls the function "WEBEXT.webviewTag.insertCSS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInsertCSS(details InjectDetails, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInsertCSS(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// HasFuncIsAudioMuted returns true if the function "WEBEXT.webviewTag.isAudioMuted" exists.
func HasFuncIsAudioMuted() bool {
	return js.True == bindings.HasFuncIsAudioMuted()
}

// FuncIsAudioMuted returns the function "WEBEXT.webviewTag.isAudioMuted".
func FuncIsAudioMuted() (fn js.Func[func(callback js.Func[func(muted bool)])]) {
	bindings.FuncIsAudioMuted(
		js.Pointer(&fn),
	)
	return
}

// IsAudioMuted calls the function "WEBEXT.webviewTag.isAudioMuted" directly.
func IsAudioMuted(callback js.Func[func(muted bool)]) (ret js.Void) {
	bindings.CallIsAudioMuted(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryIsAudioMuted calls the function "WEBEXT.webviewTag.isAudioMuted"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsAudioMuted(callback js.Func[func(muted bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsAudioMuted(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncIsSpatialNavigationEnabled returns true if the function "WEBEXT.webviewTag.isSpatialNavigationEnabled" exists.
func HasFuncIsSpatialNavigationEnabled() bool {
	return js.True == bindings.HasFuncIsSpatialNavigationEnabled()
}

// FuncIsSpatialNavigationEnabled returns the function "WEBEXT.webviewTag.isSpatialNavigationEnabled".
func FuncIsSpatialNavigationEnabled() (fn js.Func[func(callback js.Func[func(enabled bool)])]) {
	bindings.FuncIsSpatialNavigationEnabled(
		js.Pointer(&fn),
	)
	return
}

// IsSpatialNavigationEnabled calls the function "WEBEXT.webviewTag.isSpatialNavigationEnabled" directly.
func IsSpatialNavigationEnabled(callback js.Func[func(enabled bool)]) (ret js.Void) {
	bindings.CallIsSpatialNavigationEnabled(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryIsSpatialNavigationEnabled calls the function "WEBEXT.webviewTag.isSpatialNavigationEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsSpatialNavigationEnabled(callback js.Func[func(enabled bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsSpatialNavigationEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncIsUserAgentOverridden returns true if the function "WEBEXT.webviewTag.isUserAgentOverridden" exists.
func HasFuncIsUserAgentOverridden() bool {
	return js.True == bindings.HasFuncIsUserAgentOverridden()
}

// FuncIsUserAgentOverridden returns the function "WEBEXT.webviewTag.isUserAgentOverridden".
func FuncIsUserAgentOverridden() (fn js.Func[func()]) {
	bindings.FuncIsUserAgentOverridden(
		js.Pointer(&fn),
	)
	return
}

// IsUserAgentOverridden calls the function "WEBEXT.webviewTag.isUserAgentOverridden" directly.
func IsUserAgentOverridden() (ret js.Void) {
	bindings.CallIsUserAgentOverridden(
		js.Pointer(&ret),
	)

	return
}

// TryIsUserAgentOverridden calls the function "WEBEXT.webviewTag.isUserAgentOverridden"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsUserAgentOverridden() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsUserAgentOverridden(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncLoadDataWithBaseUrl returns true if the function "WEBEXT.webviewTag.loadDataWithBaseUrl" exists.
func HasFuncLoadDataWithBaseUrl() bool {
	return js.True == bindings.HasFuncLoadDataWithBaseUrl()
}

// FuncLoadDataWithBaseUrl returns the function "WEBEXT.webviewTag.loadDataWithBaseUrl".
func FuncLoadDataWithBaseUrl() (fn js.Func[func(dataUrl js.String, baseUrl js.String, virtualUrl js.String)]) {
	bindings.FuncLoadDataWithBaseUrl(
		js.Pointer(&fn),
	)
	return
}

// LoadDataWithBaseUrl calls the function "WEBEXT.webviewTag.loadDataWithBaseUrl" directly.
func LoadDataWithBaseUrl(dataUrl js.String, baseUrl js.String, virtualUrl js.String) (ret js.Void) {
	bindings.CallLoadDataWithBaseUrl(
		js.Pointer(&ret),
		dataUrl.Ref(),
		baseUrl.Ref(),
		virtualUrl.Ref(),
	)

	return
}

// TryLoadDataWithBaseUrl calls the function "WEBEXT.webviewTag.loadDataWithBaseUrl"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLoadDataWithBaseUrl(dataUrl js.String, baseUrl js.String, virtualUrl js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLoadDataWithBaseUrl(
		js.Pointer(&ret), js.Pointer(&exception),
		dataUrl.Ref(),
		baseUrl.Ref(),
		virtualUrl.Ref(),
	)

	return
}

type LoadabortEventCallbackFunc func(this js.Ref, url js.String, isTopLevel bool, code int64, reason LoadabortArgReason) js.Ref

func (fn LoadabortEventCallbackFunc) Register() js.Func[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)] {
	return js.RegisterCallback[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LoadabortEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.String{}.FromRef(args[0+1]),
		args[1+1] == js.True,
		js.BigInt[int64]{}.FromRef(args[2+1]).Get(),
		LoadabortArgReason(0).FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LoadabortEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, url js.String, isTopLevel bool, code int64, reason LoadabortArgReason) js.Ref
	Arg T
}

func (cb *LoadabortEventCallback[T]) Register() js.Func[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)] {
	return js.RegisterCallback[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LoadabortEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		args[1+1] == js.True,
		js.BigInt[int64]{}.FromRef(args[2+1]).Get(),
		LoadabortArgReason(0).FromRef(args[3+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnLoadabort returns true if the function "WEBEXT.webviewTag.loadabort.addListener" exists.
func HasFuncOnLoadabort() bool {
	return js.True == bindings.HasFuncOnLoadabort()
}

// FuncOnLoadabort returns the function "WEBEXT.webviewTag.loadabort.addListener".
func FuncOnLoadabort() (fn js.Func[func(callback js.Func[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)])]) {
	bindings.FuncOnLoadabort(
		js.Pointer(&fn),
	)
	return
}

// OnLoadabort calls the function "WEBEXT.webviewTag.loadabort.addListener" directly.
func OnLoadabort(callback js.Func[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)]) (ret js.Void) {
	bindings.CallOnLoadabort(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnLoadabort calls the function "WEBEXT.webviewTag.loadabort.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnLoadabort(callback js.Func[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnLoadabort(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffLoadabort returns true if the function "WEBEXT.webviewTag.loadabort.removeListener" exists.
func HasFuncOffLoadabort() bool {
	return js.True == bindings.HasFuncOffLoadabort()
}

// FuncOffLoadabort returns the function "WEBEXT.webviewTag.loadabort.removeListener".
func FuncOffLoadabort() (fn js.Func[func(callback js.Func[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)])]) {
	bindings.FuncOffLoadabort(
		js.Pointer(&fn),
	)
	return
}

// OffLoadabort calls the function "WEBEXT.webviewTag.loadabort.removeListener" directly.
func OffLoadabort(callback js.Func[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)]) (ret js.Void) {
	bindings.CallOffLoadabort(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffLoadabort calls the function "WEBEXT.webviewTag.loadabort.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffLoadabort(callback js.Func[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffLoadabort(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnLoadabort returns true if the function "WEBEXT.webviewTag.loadabort.hasListener" exists.
func HasFuncHasOnLoadabort() bool {
	return js.True == bindings.HasFuncHasOnLoadabort()
}

// FuncHasOnLoadabort returns the function "WEBEXT.webviewTag.loadabort.hasListener".
func FuncHasOnLoadabort() (fn js.Func[func(callback js.Func[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)]) bool]) {
	bindings.FuncHasOnLoadabort(
		js.Pointer(&fn),
	)
	return
}

// HasOnLoadabort calls the function "WEBEXT.webviewTag.loadabort.hasListener" directly.
func HasOnLoadabort(callback js.Func[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)]) (ret bool) {
	bindings.CallHasOnLoadabort(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnLoadabort calls the function "WEBEXT.webviewTag.loadabort.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnLoadabort(callback js.Func[func(url js.String, isTopLevel bool, code int64, reason LoadabortArgReason)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnLoadabort(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type LoadcommitEventCallbackFunc func(this js.Ref, url js.String, isTopLevel bool) js.Ref

func (fn LoadcommitEventCallbackFunc) Register() js.Func[func(url js.String, isTopLevel bool)] {
	return js.RegisterCallback[func(url js.String, isTopLevel bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LoadcommitEventCallbackFunc) DispatchCallback(
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
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LoadcommitEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, url js.String, isTopLevel bool) js.Ref
	Arg T
}

func (cb *LoadcommitEventCallback[T]) Register() js.Func[func(url js.String, isTopLevel bool)] {
	return js.RegisterCallback[func(url js.String, isTopLevel bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LoadcommitEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnLoadcommit returns true if the function "WEBEXT.webviewTag.loadcommit.addListener" exists.
func HasFuncOnLoadcommit() bool {
	return js.True == bindings.HasFuncOnLoadcommit()
}

// FuncOnLoadcommit returns the function "WEBEXT.webviewTag.loadcommit.addListener".
func FuncOnLoadcommit() (fn js.Func[func(callback js.Func[func(url js.String, isTopLevel bool)])]) {
	bindings.FuncOnLoadcommit(
		js.Pointer(&fn),
	)
	return
}

// OnLoadcommit calls the function "WEBEXT.webviewTag.loadcommit.addListener" directly.
func OnLoadcommit(callback js.Func[func(url js.String, isTopLevel bool)]) (ret js.Void) {
	bindings.CallOnLoadcommit(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnLoadcommit calls the function "WEBEXT.webviewTag.loadcommit.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnLoadcommit(callback js.Func[func(url js.String, isTopLevel bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnLoadcommit(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffLoadcommit returns true if the function "WEBEXT.webviewTag.loadcommit.removeListener" exists.
func HasFuncOffLoadcommit() bool {
	return js.True == bindings.HasFuncOffLoadcommit()
}

// FuncOffLoadcommit returns the function "WEBEXT.webviewTag.loadcommit.removeListener".
func FuncOffLoadcommit() (fn js.Func[func(callback js.Func[func(url js.String, isTopLevel bool)])]) {
	bindings.FuncOffLoadcommit(
		js.Pointer(&fn),
	)
	return
}

// OffLoadcommit calls the function "WEBEXT.webviewTag.loadcommit.removeListener" directly.
func OffLoadcommit(callback js.Func[func(url js.String, isTopLevel bool)]) (ret js.Void) {
	bindings.CallOffLoadcommit(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffLoadcommit calls the function "WEBEXT.webviewTag.loadcommit.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffLoadcommit(callback js.Func[func(url js.String, isTopLevel bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffLoadcommit(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnLoadcommit returns true if the function "WEBEXT.webviewTag.loadcommit.hasListener" exists.
func HasFuncHasOnLoadcommit() bool {
	return js.True == bindings.HasFuncHasOnLoadcommit()
}

// FuncHasOnLoadcommit returns the function "WEBEXT.webviewTag.loadcommit.hasListener".
func FuncHasOnLoadcommit() (fn js.Func[func(callback js.Func[func(url js.String, isTopLevel bool)]) bool]) {
	bindings.FuncHasOnLoadcommit(
		js.Pointer(&fn),
	)
	return
}

// HasOnLoadcommit calls the function "WEBEXT.webviewTag.loadcommit.hasListener" directly.
func HasOnLoadcommit(callback js.Func[func(url js.String, isTopLevel bool)]) (ret bool) {
	bindings.CallHasOnLoadcommit(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnLoadcommit calls the function "WEBEXT.webviewTag.loadcommit.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnLoadcommit(callback js.Func[func(url js.String, isTopLevel bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnLoadcommit(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type LoadredirectEventCallbackFunc func(this js.Ref, oldUrl js.String, newUrl js.String, isTopLevel bool) js.Ref

func (fn LoadredirectEventCallbackFunc) Register() js.Func[func(oldUrl js.String, newUrl js.String, isTopLevel bool)] {
	return js.RegisterCallback[func(oldUrl js.String, newUrl js.String, isTopLevel bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LoadredirectEventCallbackFunc) DispatchCallback(
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
		args[2+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LoadredirectEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, oldUrl js.String, newUrl js.String, isTopLevel bool) js.Ref
	Arg T
}

func (cb *LoadredirectEventCallback[T]) Register() js.Func[func(oldUrl js.String, newUrl js.String, isTopLevel bool)] {
	return js.RegisterCallback[func(oldUrl js.String, newUrl js.String, isTopLevel bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LoadredirectEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 3+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		args[2+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnLoadredirect returns true if the function "WEBEXT.webviewTag.loadredirect.addListener" exists.
func HasFuncOnLoadredirect() bool {
	return js.True == bindings.HasFuncOnLoadredirect()
}

// FuncOnLoadredirect returns the function "WEBEXT.webviewTag.loadredirect.addListener".
func FuncOnLoadredirect() (fn js.Func[func(callback js.Func[func(oldUrl js.String, newUrl js.String, isTopLevel bool)])]) {
	bindings.FuncOnLoadredirect(
		js.Pointer(&fn),
	)
	return
}

// OnLoadredirect calls the function "WEBEXT.webviewTag.loadredirect.addListener" directly.
func OnLoadredirect(callback js.Func[func(oldUrl js.String, newUrl js.String, isTopLevel bool)]) (ret js.Void) {
	bindings.CallOnLoadredirect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnLoadredirect calls the function "WEBEXT.webviewTag.loadredirect.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnLoadredirect(callback js.Func[func(oldUrl js.String, newUrl js.String, isTopLevel bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnLoadredirect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffLoadredirect returns true if the function "WEBEXT.webviewTag.loadredirect.removeListener" exists.
func HasFuncOffLoadredirect() bool {
	return js.True == bindings.HasFuncOffLoadredirect()
}

// FuncOffLoadredirect returns the function "WEBEXT.webviewTag.loadredirect.removeListener".
func FuncOffLoadredirect() (fn js.Func[func(callback js.Func[func(oldUrl js.String, newUrl js.String, isTopLevel bool)])]) {
	bindings.FuncOffLoadredirect(
		js.Pointer(&fn),
	)
	return
}

// OffLoadredirect calls the function "WEBEXT.webviewTag.loadredirect.removeListener" directly.
func OffLoadredirect(callback js.Func[func(oldUrl js.String, newUrl js.String, isTopLevel bool)]) (ret js.Void) {
	bindings.CallOffLoadredirect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffLoadredirect calls the function "WEBEXT.webviewTag.loadredirect.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffLoadredirect(callback js.Func[func(oldUrl js.String, newUrl js.String, isTopLevel bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffLoadredirect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnLoadredirect returns true if the function "WEBEXT.webviewTag.loadredirect.hasListener" exists.
func HasFuncHasOnLoadredirect() bool {
	return js.True == bindings.HasFuncHasOnLoadredirect()
}

// FuncHasOnLoadredirect returns the function "WEBEXT.webviewTag.loadredirect.hasListener".
func FuncHasOnLoadredirect() (fn js.Func[func(callback js.Func[func(oldUrl js.String, newUrl js.String, isTopLevel bool)]) bool]) {
	bindings.FuncHasOnLoadredirect(
		js.Pointer(&fn),
	)
	return
}

// HasOnLoadredirect calls the function "WEBEXT.webviewTag.loadredirect.hasListener" directly.
func HasOnLoadredirect(callback js.Func[func(oldUrl js.String, newUrl js.String, isTopLevel bool)]) (ret bool) {
	bindings.CallHasOnLoadredirect(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnLoadredirect calls the function "WEBEXT.webviewTag.loadredirect.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnLoadredirect(callback js.Func[func(oldUrl js.String, newUrl js.String, isTopLevel bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnLoadredirect(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type LoadstartEventCallbackFunc func(this js.Ref, url js.String, isTopLevel bool) js.Ref

func (fn LoadstartEventCallbackFunc) Register() js.Func[func(url js.String, isTopLevel bool)] {
	return js.RegisterCallback[func(url js.String, isTopLevel bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LoadstartEventCallbackFunc) DispatchCallback(
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
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type LoadstartEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, url js.String, isTopLevel bool) js.Ref
	Arg T
}

func (cb *LoadstartEventCallback[T]) Register() js.Func[func(url js.String, isTopLevel bool)] {
	return js.RegisterCallback[func(url js.String, isTopLevel bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LoadstartEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.String{}.FromRef(args[0+1]),
		args[1+1] == js.True,
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnLoadstart returns true if the function "WEBEXT.webviewTag.loadstart.addListener" exists.
func HasFuncOnLoadstart() bool {
	return js.True == bindings.HasFuncOnLoadstart()
}

// FuncOnLoadstart returns the function "WEBEXT.webviewTag.loadstart.addListener".
func FuncOnLoadstart() (fn js.Func[func(callback js.Func[func(url js.String, isTopLevel bool)])]) {
	bindings.FuncOnLoadstart(
		js.Pointer(&fn),
	)
	return
}

// OnLoadstart calls the function "WEBEXT.webviewTag.loadstart.addListener" directly.
func OnLoadstart(callback js.Func[func(url js.String, isTopLevel bool)]) (ret js.Void) {
	bindings.CallOnLoadstart(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnLoadstart calls the function "WEBEXT.webviewTag.loadstart.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnLoadstart(callback js.Func[func(url js.String, isTopLevel bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnLoadstart(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffLoadstart returns true if the function "WEBEXT.webviewTag.loadstart.removeListener" exists.
func HasFuncOffLoadstart() bool {
	return js.True == bindings.HasFuncOffLoadstart()
}

// FuncOffLoadstart returns the function "WEBEXT.webviewTag.loadstart.removeListener".
func FuncOffLoadstart() (fn js.Func[func(callback js.Func[func(url js.String, isTopLevel bool)])]) {
	bindings.FuncOffLoadstart(
		js.Pointer(&fn),
	)
	return
}

// OffLoadstart calls the function "WEBEXT.webviewTag.loadstart.removeListener" directly.
func OffLoadstart(callback js.Func[func(url js.String, isTopLevel bool)]) (ret js.Void) {
	bindings.CallOffLoadstart(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffLoadstart calls the function "WEBEXT.webviewTag.loadstart.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffLoadstart(callback js.Func[func(url js.String, isTopLevel bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffLoadstart(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnLoadstart returns true if the function "WEBEXT.webviewTag.loadstart.hasListener" exists.
func HasFuncHasOnLoadstart() bool {
	return js.True == bindings.HasFuncHasOnLoadstart()
}

// FuncHasOnLoadstart returns the function "WEBEXT.webviewTag.loadstart.hasListener".
func FuncHasOnLoadstart() (fn js.Func[func(callback js.Func[func(url js.String, isTopLevel bool)]) bool]) {
	bindings.FuncHasOnLoadstart(
		js.Pointer(&fn),
	)
	return
}

// HasOnLoadstart calls the function "WEBEXT.webviewTag.loadstart.hasListener" directly.
func HasOnLoadstart(callback js.Func[func(url js.String, isTopLevel bool)]) (ret bool) {
	bindings.CallHasOnLoadstart(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnLoadstart calls the function "WEBEXT.webviewTag.loadstart.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnLoadstart(callback js.Func[func(url js.String, isTopLevel bool)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnLoadstart(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type LoadstopEventCallbackFunc func(this js.Ref) js.Ref

func (fn LoadstopEventCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LoadstopEventCallbackFunc) DispatchCallback(
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

type LoadstopEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *LoadstopEventCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LoadstopEventCallback[T]) DispatchCallback(
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

// HasFuncOnLoadstop returns true if the function "WEBEXT.webviewTag.loadstop.addListener" exists.
func HasFuncOnLoadstop() bool {
	return js.True == bindings.HasFuncOnLoadstop()
}

// FuncOnLoadstop returns the function "WEBEXT.webviewTag.loadstop.addListener".
func FuncOnLoadstop() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOnLoadstop(
		js.Pointer(&fn),
	)
	return
}

// OnLoadstop calls the function "WEBEXT.webviewTag.loadstop.addListener" directly.
func OnLoadstop(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOnLoadstop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnLoadstop calls the function "WEBEXT.webviewTag.loadstop.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnLoadstop(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnLoadstop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffLoadstop returns true if the function "WEBEXT.webviewTag.loadstop.removeListener" exists.
func HasFuncOffLoadstop() bool {
	return js.True == bindings.HasFuncOffLoadstop()
}

// FuncOffLoadstop returns the function "WEBEXT.webviewTag.loadstop.removeListener".
func FuncOffLoadstop() (fn js.Func[func(callback js.Func[func()])]) {
	bindings.FuncOffLoadstop(
		js.Pointer(&fn),
	)
	return
}

// OffLoadstop calls the function "WEBEXT.webviewTag.loadstop.removeListener" directly.
func OffLoadstop(callback js.Func[func()]) (ret js.Void) {
	bindings.CallOffLoadstop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffLoadstop calls the function "WEBEXT.webviewTag.loadstop.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffLoadstop(callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffLoadstop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnLoadstop returns true if the function "WEBEXT.webviewTag.loadstop.hasListener" exists.
func HasFuncHasOnLoadstop() bool {
	return js.True == bindings.HasFuncHasOnLoadstop()
}

// FuncHasOnLoadstop returns the function "WEBEXT.webviewTag.loadstop.hasListener".
func FuncHasOnLoadstop() (fn js.Func[func(callback js.Func[func()]) bool]) {
	bindings.FuncHasOnLoadstop(
		js.Pointer(&fn),
	)
	return
}

// HasOnLoadstop calls the function "WEBEXT.webviewTag.loadstop.hasListener" directly.
func HasOnLoadstop(callback js.Func[func()]) (ret bool) {
	bindings.CallHasOnLoadstop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnLoadstop calls the function "WEBEXT.webviewTag.loadstop.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnLoadstop(callback js.Func[func()]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnLoadstop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type NewwindowEventCallbackFunc func(this js.Ref, window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition) js.Ref

func (fn NewwindowEventCallbackFunc) Register() js.Func[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)] {
	return js.RegisterCallback[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn NewwindowEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 6+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		NewWindow{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		js.Number[float64]{}.FromRef(args[2+1]).Get(),
		js.Number[float64]{}.FromRef(args[3+1]).Get(),
		js.String{}.FromRef(args[4+1]),
		NewwindowArgWindowOpenDisposition(0).FromRef(args[5+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type NewwindowEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition) js.Ref
	Arg T
}

func (cb *NewwindowEventCallback[T]) Register() js.Func[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)] {
	return js.RegisterCallback[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *NewwindowEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 6+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		NewWindow{}.FromRef(args[0+1]),
		js.String{}.FromRef(args[1+1]),
		js.Number[float64]{}.FromRef(args[2+1]).Get(),
		js.Number[float64]{}.FromRef(args[3+1]).Get(),
		js.String{}.FromRef(args[4+1]),
		NewwindowArgWindowOpenDisposition(0).FromRef(args[5+1]),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnNewwindow returns true if the function "WEBEXT.webviewTag.newwindow.addListener" exists.
func HasFuncOnNewwindow() bool {
	return js.True == bindings.HasFuncOnNewwindow()
}

// FuncOnNewwindow returns the function "WEBEXT.webviewTag.newwindow.addListener".
func FuncOnNewwindow() (fn js.Func[func(callback js.Func[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)])]) {
	bindings.FuncOnNewwindow(
		js.Pointer(&fn),
	)
	return
}

// OnNewwindow calls the function "WEBEXT.webviewTag.newwindow.addListener" directly.
func OnNewwindow(callback js.Func[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)]) (ret js.Void) {
	bindings.CallOnNewwindow(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnNewwindow calls the function "WEBEXT.webviewTag.newwindow.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnNewwindow(callback js.Func[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnNewwindow(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffNewwindow returns true if the function "WEBEXT.webviewTag.newwindow.removeListener" exists.
func HasFuncOffNewwindow() bool {
	return js.True == bindings.HasFuncOffNewwindow()
}

// FuncOffNewwindow returns the function "WEBEXT.webviewTag.newwindow.removeListener".
func FuncOffNewwindow() (fn js.Func[func(callback js.Func[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)])]) {
	bindings.FuncOffNewwindow(
		js.Pointer(&fn),
	)
	return
}

// OffNewwindow calls the function "WEBEXT.webviewTag.newwindow.removeListener" directly.
func OffNewwindow(callback js.Func[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)]) (ret js.Void) {
	bindings.CallOffNewwindow(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffNewwindow calls the function "WEBEXT.webviewTag.newwindow.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffNewwindow(callback js.Func[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffNewwindow(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnNewwindow returns true if the function "WEBEXT.webviewTag.newwindow.hasListener" exists.
func HasFuncHasOnNewwindow() bool {
	return js.True == bindings.HasFuncHasOnNewwindow()
}

// FuncHasOnNewwindow returns the function "WEBEXT.webviewTag.newwindow.hasListener".
func FuncHasOnNewwindow() (fn js.Func[func(callback js.Func[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)]) bool]) {
	bindings.FuncHasOnNewwindow(
		js.Pointer(&fn),
	)
	return
}

// HasOnNewwindow calls the function "WEBEXT.webviewTag.newwindow.hasListener" directly.
func HasOnNewwindow(callback js.Func[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)]) (ret bool) {
	bindings.CallHasOnNewwindow(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnNewwindow calls the function "WEBEXT.webviewTag.newwindow.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnNewwindow(callback js.Func[func(window NewWindow, targetUrl js.String, initialWidth float64, initialHeight float64, name js.String, windowOpenDisposition NewwindowArgWindowOpenDisposition)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnNewwindow(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type PermissionrequestEventCallbackFunc func(this js.Ref, permission PermissionrequestArgPermission, request *PermissionrequestArgRequest) js.Ref

func (fn PermissionrequestEventCallbackFunc) Register() js.Func[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)] {
	return js.RegisterCallback[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn PermissionrequestEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 PermissionrequestArgRequest
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(fn(
		args[0],

		PermissionrequestArgPermission(0).FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type PermissionrequestEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, permission PermissionrequestArgPermission, request *PermissionrequestArgRequest) js.Ref
	Arg T
}

func (cb *PermissionrequestEventCallback[T]) Register() js.Func[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)] {
	return js.RegisterCallback[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *PermissionrequestEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg1 PermissionrequestArgRequest
	arg1.UpdateFrom(args[1+1])
	defer arg1.FreeMembers(true)

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		PermissionrequestArgPermission(0).FromRef(args[0+1]),
		mark.NoEscape(&arg1),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnPermissionrequest returns true if the function "WEBEXT.webviewTag.permissionrequest.addListener" exists.
func HasFuncOnPermissionrequest() bool {
	return js.True == bindings.HasFuncOnPermissionrequest()
}

// FuncOnPermissionrequest returns the function "WEBEXT.webviewTag.permissionrequest.addListener".
func FuncOnPermissionrequest() (fn js.Func[func(callback js.Func[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)])]) {
	bindings.FuncOnPermissionrequest(
		js.Pointer(&fn),
	)
	return
}

// OnPermissionrequest calls the function "WEBEXT.webviewTag.permissionrequest.addListener" directly.
func OnPermissionrequest(callback js.Func[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)]) (ret js.Void) {
	bindings.CallOnPermissionrequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnPermissionrequest calls the function "WEBEXT.webviewTag.permissionrequest.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnPermissionrequest(callback js.Func[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnPermissionrequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffPermissionrequest returns true if the function "WEBEXT.webviewTag.permissionrequest.removeListener" exists.
func HasFuncOffPermissionrequest() bool {
	return js.True == bindings.HasFuncOffPermissionrequest()
}

// FuncOffPermissionrequest returns the function "WEBEXT.webviewTag.permissionrequest.removeListener".
func FuncOffPermissionrequest() (fn js.Func[func(callback js.Func[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)])]) {
	bindings.FuncOffPermissionrequest(
		js.Pointer(&fn),
	)
	return
}

// OffPermissionrequest calls the function "WEBEXT.webviewTag.permissionrequest.removeListener" directly.
func OffPermissionrequest(callback js.Func[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)]) (ret js.Void) {
	bindings.CallOffPermissionrequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffPermissionrequest calls the function "WEBEXT.webviewTag.permissionrequest.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffPermissionrequest(callback js.Func[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffPermissionrequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnPermissionrequest returns true if the function "WEBEXT.webviewTag.permissionrequest.hasListener" exists.
func HasFuncHasOnPermissionrequest() bool {
	return js.True == bindings.HasFuncHasOnPermissionrequest()
}

// FuncHasOnPermissionrequest returns the function "WEBEXT.webviewTag.permissionrequest.hasListener".
func FuncHasOnPermissionrequest() (fn js.Func[func(callback js.Func[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)]) bool]) {
	bindings.FuncHasOnPermissionrequest(
		js.Pointer(&fn),
	)
	return
}

// HasOnPermissionrequest calls the function "WEBEXT.webviewTag.permissionrequest.hasListener" directly.
func HasOnPermissionrequest(callback js.Func[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)]) (ret bool) {
	bindings.CallHasOnPermissionrequest(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnPermissionrequest calls the function "WEBEXT.webviewTag.permissionrequest.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnPermissionrequest(callback js.Func[func(permission PermissionrequestArgPermission, request *PermissionrequestArgRequest)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnPermissionrequest(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncPrint returns true if the function "WEBEXT.webviewTag.print" exists.
func HasFuncPrint() bool {
	return js.True == bindings.HasFuncPrint()
}

// FuncPrint returns the function "WEBEXT.webviewTag.print".
func FuncPrint() (fn js.Func[func()]) {
	bindings.FuncPrint(
		js.Pointer(&fn),
	)
	return
}

// Print calls the function "WEBEXT.webviewTag.print" directly.
func Print() (ret js.Void) {
	bindings.CallPrint(
		js.Pointer(&ret),
	)

	return
}

// TryPrint calls the function "WEBEXT.webviewTag.print"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPrint() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryPrint(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReload returns true if the function "WEBEXT.webviewTag.reload" exists.
func HasFuncReload() bool {
	return js.True == bindings.HasFuncReload()
}

// FuncReload returns the function "WEBEXT.webviewTag.reload".
func FuncReload() (fn js.Func[func()]) {
	bindings.FuncReload(
		js.Pointer(&fn),
	)
	return
}

// Reload calls the function "WEBEXT.webviewTag.reload" directly.
func Reload() (ret js.Void) {
	bindings.CallReload(
		js.Pointer(&ret),
	)

	return
}

// TryReload calls the function "WEBEXT.webviewTag.reload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReload() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReload(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemoveContentScripts returns true if the function "WEBEXT.webviewTag.removeContentScripts" exists.
func HasFuncRemoveContentScripts() bool {
	return js.True == bindings.HasFuncRemoveContentScripts()
}

// FuncRemoveContentScripts returns the function "WEBEXT.webviewTag.removeContentScripts".
func FuncRemoveContentScripts() (fn js.Func[func(scriptNameList js.Array[js.String])]) {
	bindings.FuncRemoveContentScripts(
		js.Pointer(&fn),
	)
	return
}

// RemoveContentScripts calls the function "WEBEXT.webviewTag.removeContentScripts" directly.
func RemoveContentScripts(scriptNameList js.Array[js.String]) (ret js.Void) {
	bindings.CallRemoveContentScripts(
		js.Pointer(&ret),
		scriptNameList.Ref(),
	)

	return
}

// TryRemoveContentScripts calls the function "WEBEXT.webviewTag.removeContentScripts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveContentScripts(scriptNameList js.Array[js.String]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveContentScripts(
		js.Pointer(&ret), js.Pointer(&exception),
		scriptNameList.Ref(),
	)

	return
}

// Request returns the value of property "WEBEXT.webviewTag.request".
//
// The returned bool will be false if there is no such property.
func Request() (ret WebRequestEventInterface, ok bool) {
	ok = js.True == bindings.GetRequest(
		js.Pointer(&ret),
	)

	return
}

// SetRequest sets the value of property "WEBEXT.webviewTag.request" to val.
//
// It returns false if the property cannot be set.
func SetRequest(val WebRequestEventInterface) bool {
	return js.True == bindings.SetRequest(
		js.Pointer(&val))
}

type ResponsiveEventCallbackFunc func(this js.Ref, processID int64) js.Ref

func (fn ResponsiveEventCallbackFunc) Register() js.Func[func(processID int64)] {
	return js.RegisterCallback[func(processID int64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ResponsiveEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ResponsiveEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, processID int64) js.Ref
	Arg T
}

func (cb *ResponsiveEventCallback[T]) Register() js.Func[func(processID int64)] {
	return js.RegisterCallback[func(processID int64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ResponsiveEventCallback[T]) DispatchCallback(
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

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnResponsive returns true if the function "WEBEXT.webviewTag.responsive.addListener" exists.
func HasFuncOnResponsive() bool {
	return js.True == bindings.HasFuncOnResponsive()
}

// FuncOnResponsive returns the function "WEBEXT.webviewTag.responsive.addListener".
func FuncOnResponsive() (fn js.Func[func(callback js.Func[func(processID int64)])]) {
	bindings.FuncOnResponsive(
		js.Pointer(&fn),
	)
	return
}

// OnResponsive calls the function "WEBEXT.webviewTag.responsive.addListener" directly.
func OnResponsive(callback js.Func[func(processID int64)]) (ret js.Void) {
	bindings.CallOnResponsive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnResponsive calls the function "WEBEXT.webviewTag.responsive.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnResponsive(callback js.Func[func(processID int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnResponsive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffResponsive returns true if the function "WEBEXT.webviewTag.responsive.removeListener" exists.
func HasFuncOffResponsive() bool {
	return js.True == bindings.HasFuncOffResponsive()
}

// FuncOffResponsive returns the function "WEBEXT.webviewTag.responsive.removeListener".
func FuncOffResponsive() (fn js.Func[func(callback js.Func[func(processID int64)])]) {
	bindings.FuncOffResponsive(
		js.Pointer(&fn),
	)
	return
}

// OffResponsive calls the function "WEBEXT.webviewTag.responsive.removeListener" directly.
func OffResponsive(callback js.Func[func(processID int64)]) (ret js.Void) {
	bindings.CallOffResponsive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffResponsive calls the function "WEBEXT.webviewTag.responsive.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffResponsive(callback js.Func[func(processID int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffResponsive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnResponsive returns true if the function "WEBEXT.webviewTag.responsive.hasListener" exists.
func HasFuncHasOnResponsive() bool {
	return js.True == bindings.HasFuncHasOnResponsive()
}

// FuncHasOnResponsive returns the function "WEBEXT.webviewTag.responsive.hasListener".
func FuncHasOnResponsive() (fn js.Func[func(callback js.Func[func(processID int64)]) bool]) {
	bindings.FuncHasOnResponsive(
		js.Pointer(&fn),
	)
	return
}

// HasOnResponsive calls the function "WEBEXT.webviewTag.responsive.hasListener" directly.
func HasOnResponsive(callback js.Func[func(processID int64)]) (ret bool) {
	bindings.CallHasOnResponsive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnResponsive calls the function "WEBEXT.webviewTag.responsive.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnResponsive(callback js.Func[func(processID int64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnResponsive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSetAudioMuted returns true if the function "WEBEXT.webviewTag.setAudioMuted" exists.
func HasFuncSetAudioMuted() bool {
	return js.True == bindings.HasFuncSetAudioMuted()
}

// FuncSetAudioMuted returns the function "WEBEXT.webviewTag.setAudioMuted".
func FuncSetAudioMuted() (fn js.Func[func(mute bool)]) {
	bindings.FuncSetAudioMuted(
		js.Pointer(&fn),
	)
	return
}

// SetAudioMuted calls the function "WEBEXT.webviewTag.setAudioMuted" directly.
func SetAudioMuted(mute bool) (ret js.Void) {
	bindings.CallSetAudioMuted(
		js.Pointer(&ret),
		js.Bool(bool(mute)),
	)

	return
}

// TrySetAudioMuted calls the function "WEBEXT.webviewTag.setAudioMuted"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAudioMuted(mute bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAudioMuted(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(mute)),
	)

	return
}

// HasFuncSetSpatialNavigationEnabled returns true if the function "WEBEXT.webviewTag.setSpatialNavigationEnabled" exists.
func HasFuncSetSpatialNavigationEnabled() bool {
	return js.True == bindings.HasFuncSetSpatialNavigationEnabled()
}

// FuncSetSpatialNavigationEnabled returns the function "WEBEXT.webviewTag.setSpatialNavigationEnabled".
func FuncSetSpatialNavigationEnabled() (fn js.Func[func(enabled bool)]) {
	bindings.FuncSetSpatialNavigationEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetSpatialNavigationEnabled calls the function "WEBEXT.webviewTag.setSpatialNavigationEnabled" directly.
func SetSpatialNavigationEnabled(enabled bool) (ret js.Void) {
	bindings.CallSetSpatialNavigationEnabled(
		js.Pointer(&ret),
		js.Bool(bool(enabled)),
	)

	return
}

// TrySetSpatialNavigationEnabled calls the function "WEBEXT.webviewTag.setSpatialNavigationEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetSpatialNavigationEnabled(enabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetSpatialNavigationEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Bool(bool(enabled)),
	)

	return
}

// HasFuncSetUserAgentOverride returns true if the function "WEBEXT.webviewTag.setUserAgentOverride" exists.
func HasFuncSetUserAgentOverride() bool {
	return js.True == bindings.HasFuncSetUserAgentOverride()
}

// FuncSetUserAgentOverride returns the function "WEBEXT.webviewTag.setUserAgentOverride".
func FuncSetUserAgentOverride() (fn js.Func[func(userAgent js.String)]) {
	bindings.FuncSetUserAgentOverride(
		js.Pointer(&fn),
	)
	return
}

// SetUserAgentOverride calls the function "WEBEXT.webviewTag.setUserAgentOverride" directly.
func SetUserAgentOverride(userAgent js.String) (ret js.Void) {
	bindings.CallSetUserAgentOverride(
		js.Pointer(&ret),
		userAgent.Ref(),
	)

	return
}

// TrySetUserAgentOverride calls the function "WEBEXT.webviewTag.setUserAgentOverride"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetUserAgentOverride(userAgent js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetUserAgentOverride(
		js.Pointer(&ret), js.Pointer(&exception),
		userAgent.Ref(),
	)

	return
}

// HasFuncSetZoom returns true if the function "WEBEXT.webviewTag.setZoom" exists.
func HasFuncSetZoom() bool {
	return js.True == bindings.HasFuncSetZoom()
}

// FuncSetZoom returns the function "WEBEXT.webviewTag.setZoom".
func FuncSetZoom() (fn js.Func[func(zoomFactor float64, callback js.Func[func()])]) {
	bindings.FuncSetZoom(
		js.Pointer(&fn),
	)
	return
}

// SetZoom calls the function "WEBEXT.webviewTag.setZoom" directly.
func SetZoom(zoomFactor float64, callback js.Func[func()]) (ret js.Void) {
	bindings.CallSetZoom(
		js.Pointer(&ret),
		float64(zoomFactor),
		callback.Ref(),
	)

	return
}

// TrySetZoom calls the function "WEBEXT.webviewTag.setZoom"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetZoom(zoomFactor float64, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetZoom(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(zoomFactor),
		callback.Ref(),
	)

	return
}

// HasFuncSetZoomMode returns true if the function "WEBEXT.webviewTag.setZoomMode" exists.
func HasFuncSetZoomMode() bool {
	return js.True == bindings.HasFuncSetZoomMode()
}

// FuncSetZoomMode returns the function "WEBEXT.webviewTag.setZoomMode".
func FuncSetZoomMode() (fn js.Func[func(ZoomMode ZoomMode, callback js.Func[func()])]) {
	bindings.FuncSetZoomMode(
		js.Pointer(&fn),
	)
	return
}

// SetZoomMode calls the function "WEBEXT.webviewTag.setZoomMode" directly.
func SetZoomMode(ZoomMode ZoomMode, callback js.Func[func()]) (ret js.Void) {
	bindings.CallSetZoomMode(
		js.Pointer(&ret),
		uint32(ZoomMode),
		callback.Ref(),
	)

	return
}

// TrySetZoomMode calls the function "WEBEXT.webviewTag.setZoomMode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetZoomMode(ZoomMode ZoomMode, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetZoomMode(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(ZoomMode),
		callback.Ref(),
	)

	return
}

type SizechangedEventCallbackFunc func(this js.Ref, oldWidth float64, oldHeight float64, newWidth float64, newHeight float64) js.Ref

func (fn SizechangedEventCallbackFunc) Register() js.Func[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)] {
	return js.RegisterCallback[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SizechangedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
		js.Number[float64]{}.FromRef(args[1+1]).Get(),
		js.Number[float64]{}.FromRef(args[2+1]).Get(),
		js.Number[float64]{}.FromRef(args[3+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type SizechangedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, oldWidth float64, oldHeight float64, newWidth float64, newHeight float64) js.Ref
	Arg T
}

func (cb *SizechangedEventCallback[T]) Register() js.Func[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)] {
	return js.RegisterCallback[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SizechangedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 4+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
		js.Number[float64]{}.FromRef(args[1+1]).Get(),
		js.Number[float64]{}.FromRef(args[2+1]).Get(),
		js.Number[float64]{}.FromRef(args[3+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnSizechanged returns true if the function "WEBEXT.webviewTag.sizechanged.addListener" exists.
func HasFuncOnSizechanged() bool {
	return js.True == bindings.HasFuncOnSizechanged()
}

// FuncOnSizechanged returns the function "WEBEXT.webviewTag.sizechanged.addListener".
func FuncOnSizechanged() (fn js.Func[func(callback js.Func[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)])]) {
	bindings.FuncOnSizechanged(
		js.Pointer(&fn),
	)
	return
}

// OnSizechanged calls the function "WEBEXT.webviewTag.sizechanged.addListener" directly.
func OnSizechanged(callback js.Func[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)]) (ret js.Void) {
	bindings.CallOnSizechanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnSizechanged calls the function "WEBEXT.webviewTag.sizechanged.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnSizechanged(callback js.Func[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnSizechanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffSizechanged returns true if the function "WEBEXT.webviewTag.sizechanged.removeListener" exists.
func HasFuncOffSizechanged() bool {
	return js.True == bindings.HasFuncOffSizechanged()
}

// FuncOffSizechanged returns the function "WEBEXT.webviewTag.sizechanged.removeListener".
func FuncOffSizechanged() (fn js.Func[func(callback js.Func[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)])]) {
	bindings.FuncOffSizechanged(
		js.Pointer(&fn),
	)
	return
}

// OffSizechanged calls the function "WEBEXT.webviewTag.sizechanged.removeListener" directly.
func OffSizechanged(callback js.Func[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)]) (ret js.Void) {
	bindings.CallOffSizechanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffSizechanged calls the function "WEBEXT.webviewTag.sizechanged.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffSizechanged(callback js.Func[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffSizechanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnSizechanged returns true if the function "WEBEXT.webviewTag.sizechanged.hasListener" exists.
func HasFuncHasOnSizechanged() bool {
	return js.True == bindings.HasFuncHasOnSizechanged()
}

// FuncHasOnSizechanged returns the function "WEBEXT.webviewTag.sizechanged.hasListener".
func FuncHasOnSizechanged() (fn js.Func[func(callback js.Func[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)]) bool]) {
	bindings.FuncHasOnSizechanged(
		js.Pointer(&fn),
	)
	return
}

// HasOnSizechanged calls the function "WEBEXT.webviewTag.sizechanged.hasListener" directly.
func HasOnSizechanged(callback js.Func[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)]) (ret bool) {
	bindings.CallHasOnSizechanged(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnSizechanged calls the function "WEBEXT.webviewTag.sizechanged.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnSizechanged(callback js.Func[func(oldWidth float64, oldHeight float64, newWidth float64, newHeight float64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnSizechanged(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncStop returns true if the function "WEBEXT.webviewTag.stop" exists.
func HasFuncStop() bool {
	return js.True == bindings.HasFuncStop()
}

// FuncStop returns the function "WEBEXT.webviewTag.stop".
func FuncStop() (fn js.Func[func()]) {
	bindings.FuncStop(
		js.Pointer(&fn),
	)
	return
}

// Stop calls the function "WEBEXT.webviewTag.stop" directly.
func Stop() (ret js.Void) {
	bindings.CallStop(
		js.Pointer(&ret),
	)

	return
}

// TryStop calls the function "WEBEXT.webviewTag.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStop(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStopFinding returns true if the function "WEBEXT.webviewTag.stopFinding" exists.
func HasFuncStopFinding() bool {
	return js.True == bindings.HasFuncStopFinding()
}

// FuncStopFinding returns the function "WEBEXT.webviewTag.stopFinding".
func FuncStopFinding() (fn js.Func[func(action StopFindingArgAction)]) {
	bindings.FuncStopFinding(
		js.Pointer(&fn),
	)
	return
}

// StopFinding calls the function "WEBEXT.webviewTag.stopFinding" directly.
func StopFinding(action StopFindingArgAction) (ret js.Void) {
	bindings.CallStopFinding(
		js.Pointer(&ret),
		uint32(action),
	)

	return
}

// TryStopFinding calls the function "WEBEXT.webviewTag.stopFinding"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopFinding(action StopFindingArgAction) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopFinding(
		js.Pointer(&ret), js.Pointer(&exception),
		uint32(action),
	)

	return
}

// HasFuncTerminate returns true if the function "WEBEXT.webviewTag.terminate" exists.
func HasFuncTerminate() bool {
	return js.True == bindings.HasFuncTerminate()
}

// FuncTerminate returns the function "WEBEXT.webviewTag.terminate".
func FuncTerminate() (fn js.Func[func()]) {
	bindings.FuncTerminate(
		js.Pointer(&fn),
	)
	return
}

// Terminate calls the function "WEBEXT.webviewTag.terminate" directly.
func Terminate() (ret js.Void) {
	bindings.CallTerminate(
		js.Pointer(&ret),
	)

	return
}

// TryTerminate calls the function "WEBEXT.webviewTag.terminate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryTerminate() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTerminate(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type UnresponsiveEventCallbackFunc func(this js.Ref, processID int64) js.Ref

func (fn UnresponsiveEventCallbackFunc) Register() js.Func[func(processID int64)] {
	return js.RegisterCallback[func(processID int64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn UnresponsiveEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type UnresponsiveEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, processID int64) js.Ref
	Arg T
}

func (cb *UnresponsiveEventCallback[T]) Register() js.Func[func(processID int64)] {
	return js.RegisterCallback[func(processID int64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *UnresponsiveEventCallback[T]) DispatchCallback(
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

		js.BigInt[int64]{}.FromRef(args[0+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnUnresponsive returns true if the function "WEBEXT.webviewTag.unresponsive.addListener" exists.
func HasFuncOnUnresponsive() bool {
	return js.True == bindings.HasFuncOnUnresponsive()
}

// FuncOnUnresponsive returns the function "WEBEXT.webviewTag.unresponsive.addListener".
func FuncOnUnresponsive() (fn js.Func[func(callback js.Func[func(processID int64)])]) {
	bindings.FuncOnUnresponsive(
		js.Pointer(&fn),
	)
	return
}

// OnUnresponsive calls the function "WEBEXT.webviewTag.unresponsive.addListener" directly.
func OnUnresponsive(callback js.Func[func(processID int64)]) (ret js.Void) {
	bindings.CallOnUnresponsive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnUnresponsive calls the function "WEBEXT.webviewTag.unresponsive.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnUnresponsive(callback js.Func[func(processID int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnUnresponsive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffUnresponsive returns true if the function "WEBEXT.webviewTag.unresponsive.removeListener" exists.
func HasFuncOffUnresponsive() bool {
	return js.True == bindings.HasFuncOffUnresponsive()
}

// FuncOffUnresponsive returns the function "WEBEXT.webviewTag.unresponsive.removeListener".
func FuncOffUnresponsive() (fn js.Func[func(callback js.Func[func(processID int64)])]) {
	bindings.FuncOffUnresponsive(
		js.Pointer(&fn),
	)
	return
}

// OffUnresponsive calls the function "WEBEXT.webviewTag.unresponsive.removeListener" directly.
func OffUnresponsive(callback js.Func[func(processID int64)]) (ret js.Void) {
	bindings.CallOffUnresponsive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffUnresponsive calls the function "WEBEXT.webviewTag.unresponsive.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffUnresponsive(callback js.Func[func(processID int64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffUnresponsive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnUnresponsive returns true if the function "WEBEXT.webviewTag.unresponsive.hasListener" exists.
func HasFuncHasOnUnresponsive() bool {
	return js.True == bindings.HasFuncHasOnUnresponsive()
}

// FuncHasOnUnresponsive returns the function "WEBEXT.webviewTag.unresponsive.hasListener".
func FuncHasOnUnresponsive() (fn js.Func[func(callback js.Func[func(processID int64)]) bool]) {
	bindings.FuncHasOnUnresponsive(
		js.Pointer(&fn),
	)
	return
}

// HasOnUnresponsive calls the function "WEBEXT.webviewTag.unresponsive.hasListener" directly.
func HasOnUnresponsive(callback js.Func[func(processID int64)]) (ret bool) {
	bindings.CallHasOnUnresponsive(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnUnresponsive calls the function "WEBEXT.webviewTag.unresponsive.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnUnresponsive(callback js.Func[func(processID int64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnUnresponsive(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type ZoomchangeEventCallbackFunc func(this js.Ref, oldZoomFactor float64, newZoomFactor float64) js.Ref

func (fn ZoomchangeEventCallbackFunc) Register() js.Func[func(oldZoomFactor float64, newZoomFactor float64)] {
	return js.RegisterCallback[func(oldZoomFactor float64, newZoomFactor float64)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn ZoomchangeEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(fn(
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
		js.Number[float64]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

type ZoomchangeEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, oldZoomFactor float64, newZoomFactor float64) js.Ref
	Arg T
}

func (cb *ZoomchangeEventCallback[T]) Register() js.Func[func(oldZoomFactor float64, newZoomFactor float64)] {
	return js.RegisterCallback[func(oldZoomFactor float64, newZoomFactor float64)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *ZoomchangeEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 2+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}

	if ctx.Return(cb.Fn(
		cb.Arg,
		args[0],

		js.Number[float64]{}.FromRef(args[0+1]).Get(),
		js.Number[float64]{}.FromRef(args[1+1]).Get(),
	)) {
		return
	}

	js.ThrowCallbackValueNotReturned()
}

// HasFuncOnZoomchange returns true if the function "WEBEXT.webviewTag.zoomchange.addListener" exists.
func HasFuncOnZoomchange() bool {
	return js.True == bindings.HasFuncOnZoomchange()
}

// FuncOnZoomchange returns the function "WEBEXT.webviewTag.zoomchange.addListener".
func FuncOnZoomchange() (fn js.Func[func(callback js.Func[func(oldZoomFactor float64, newZoomFactor float64)])]) {
	bindings.FuncOnZoomchange(
		js.Pointer(&fn),
	)
	return
}

// OnZoomchange calls the function "WEBEXT.webviewTag.zoomchange.addListener" directly.
func OnZoomchange(callback js.Func[func(oldZoomFactor float64, newZoomFactor float64)]) (ret js.Void) {
	bindings.CallOnZoomchange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnZoomchange calls the function "WEBEXT.webviewTag.zoomchange.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnZoomchange(callback js.Func[func(oldZoomFactor float64, newZoomFactor float64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnZoomchange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffZoomchange returns true if the function "WEBEXT.webviewTag.zoomchange.removeListener" exists.
func HasFuncOffZoomchange() bool {
	return js.True == bindings.HasFuncOffZoomchange()
}

// FuncOffZoomchange returns the function "WEBEXT.webviewTag.zoomchange.removeListener".
func FuncOffZoomchange() (fn js.Func[func(callback js.Func[func(oldZoomFactor float64, newZoomFactor float64)])]) {
	bindings.FuncOffZoomchange(
		js.Pointer(&fn),
	)
	return
}

// OffZoomchange calls the function "WEBEXT.webviewTag.zoomchange.removeListener" directly.
func OffZoomchange(callback js.Func[func(oldZoomFactor float64, newZoomFactor float64)]) (ret js.Void) {
	bindings.CallOffZoomchange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffZoomchange calls the function "WEBEXT.webviewTag.zoomchange.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffZoomchange(callback js.Func[func(oldZoomFactor float64, newZoomFactor float64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffZoomchange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnZoomchange returns true if the function "WEBEXT.webviewTag.zoomchange.hasListener" exists.
func HasFuncHasOnZoomchange() bool {
	return js.True == bindings.HasFuncHasOnZoomchange()
}

// FuncHasOnZoomchange returns the function "WEBEXT.webviewTag.zoomchange.hasListener".
func FuncHasOnZoomchange() (fn js.Func[func(callback js.Func[func(oldZoomFactor float64, newZoomFactor float64)]) bool]) {
	bindings.FuncHasOnZoomchange(
		js.Pointer(&fn),
	)
	return
}

// HasOnZoomchange calls the function "WEBEXT.webviewTag.zoomchange.hasListener" directly.
func HasOnZoomchange(callback js.Func[func(oldZoomFactor float64, newZoomFactor float64)]) (ret bool) {
	bindings.CallHasOnZoomchange(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnZoomchange calls the function "WEBEXT.webviewTag.zoomchange.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnZoomchange(callback js.Func[func(oldZoomFactor float64, newZoomFactor float64)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnZoomchange(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

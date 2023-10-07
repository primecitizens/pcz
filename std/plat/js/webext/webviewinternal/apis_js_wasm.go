// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package webviewinternal

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/extensiontypes"
	"github.com/primecitizens/pcz/std/plat/js/webext/webviewinternal/bindings"
)

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

type DataTypeSet struct {
	// Appcache is "DataTypeSet.appcache"
	//
	// Optional
	//
	// NOTE: FFI_USE_Appcache MUST be set to true to make this field effective.
	Appcache bool
	// Cache is "DataTypeSet.cache"
	//
	// Optional
	//
	// NOTE: FFI_USE_Cache MUST be set to true to make this field effective.
	Cache bool
	// Cookies is "DataTypeSet.cookies"
	//
	// Optional
	//
	// NOTE: FFI_USE_Cookies MUST be set to true to make this field effective.
	Cookies bool
	// FileSystems is "DataTypeSet.fileSystems"
	//
	// Optional
	//
	// NOTE: FFI_USE_FileSystems MUST be set to true to make this field effective.
	FileSystems bool
	// IndexedDB is "DataTypeSet.indexedDB"
	//
	// Optional
	//
	// NOTE: FFI_USE_IndexedDB MUST be set to true to make this field effective.
	IndexedDB bool
	// LocalStorage is "DataTypeSet.localStorage"
	//
	// Optional
	//
	// NOTE: FFI_USE_LocalStorage MUST be set to true to make this field effective.
	LocalStorage bool
	// PersistentCookies is "DataTypeSet.persistentCookies"
	//
	// Optional
	//
	// NOTE: FFI_USE_PersistentCookies MUST be set to true to make this field effective.
	PersistentCookies bool
	// SessionCookies is "DataTypeSet.sessionCookies"
	//
	// Optional
	//
	// NOTE: FFI_USE_SessionCookies MUST be set to true to make this field effective.
	SessionCookies bool
	// WebSQL is "DataTypeSet.webSQL"
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

// FromRef calls UpdateFrom and returns a DataTypeSet with all fields set.
func (p DataTypeSet) FromRef(ref js.Ref) DataTypeSet {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DataTypeSet in the application heap.
func (p DataTypeSet) New() js.Ref {
	return bindings.DataTypeSetJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DataTypeSet) UpdateFrom(ref js.Ref) {
	bindings.DataTypeSetJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DataTypeSet) Update(ref js.Ref) {
	bindings.DataTypeSetJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DataTypeSet) FreeMembers(recursive bool) {
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

type FindArgCallbackFunc func(this js.Ref, results *FindArgCallbackArgResults) js.Ref

func (fn FindArgCallbackFunc) Register() js.Func[func(results *FindArgCallbackArgResults)] {
	return js.RegisterCallback[func(results *FindArgCallbackArgResults)](
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
	var arg0 FindArgCallbackArgResults
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
	Fn  func(arg T, this js.Ref, results *FindArgCallbackArgResults) js.Ref
	Arg T
}

func (cb *FindArgCallback[T]) Register() js.Func[func(results *FindArgCallbackArgResults)] {
	return js.RegisterCallback[func(results *FindArgCallbackArgResults)](
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
	var arg0 FindArgCallbackArgResults
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

type FindArgCallbackArgResultsFieldSelectionRect struct {
	// Height is "FindArgCallbackArgResultsFieldSelectionRect.height"
	//
	// Required
	Height int64
	// Left is "FindArgCallbackArgResultsFieldSelectionRect.left"
	//
	// Required
	Left int64
	// Top is "FindArgCallbackArgResultsFieldSelectionRect.top"
	//
	// Required
	Top int64
	// Width is "FindArgCallbackArgResultsFieldSelectionRect.width"
	//
	// Required
	Width int64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FindArgCallbackArgResultsFieldSelectionRect with all fields set.
func (p FindArgCallbackArgResultsFieldSelectionRect) FromRef(ref js.Ref) FindArgCallbackArgResultsFieldSelectionRect {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FindArgCallbackArgResultsFieldSelectionRect in the application heap.
func (p FindArgCallbackArgResultsFieldSelectionRect) New() js.Ref {
	return bindings.FindArgCallbackArgResultsFieldSelectionRectJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FindArgCallbackArgResultsFieldSelectionRect) UpdateFrom(ref js.Ref) {
	bindings.FindArgCallbackArgResultsFieldSelectionRectJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FindArgCallbackArgResultsFieldSelectionRect) Update(ref js.Ref) {
	bindings.FindArgCallbackArgResultsFieldSelectionRectJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FindArgCallbackArgResultsFieldSelectionRect) FreeMembers(recursive bool) {
}

type FindArgCallbackArgResults struct {
	// ActiveMatchOrdinal is "FindArgCallbackArgResults.activeMatchOrdinal"
	//
	// Required
	ActiveMatchOrdinal int64
	// Canceled is "FindArgCallbackArgResults.canceled"
	//
	// Required
	Canceled bool
	// NumberOfMatches is "FindArgCallbackArgResults.numberOfMatches"
	//
	// Required
	NumberOfMatches int64
	// SelectionRect is "FindArgCallbackArgResults.selectionRect"
	//
	// Required
	//
	// NOTE: SelectionRect.FFI_USE MUST be set to true to get SelectionRect used.
	SelectionRect FindArgCallbackArgResultsFieldSelectionRect

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FindArgCallbackArgResults with all fields set.
func (p FindArgCallbackArgResults) FromRef(ref js.Ref) FindArgCallbackArgResults {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FindArgCallbackArgResults in the application heap.
func (p FindArgCallbackArgResults) New() js.Ref {
	return bindings.FindArgCallbackArgResultsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FindArgCallbackArgResults) UpdateFrom(ref js.Ref) {
	bindings.FindArgCallbackArgResultsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FindArgCallbackArgResults) Update(ref js.Ref) {
	bindings.FindArgCallbackArgResultsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FindArgCallbackArgResults) FreeMembers(recursive bool) {
	if recursive {
		p.SelectionRect.FreeMembers(true)
	}
}

type FindArgOptions struct {
	// Backward is "FindArgOptions.backward"
	//
	// Optional
	//
	// NOTE: FFI_USE_Backward MUST be set to true to make this field effective.
	Backward bool
	// MatchCase is "FindArgOptions.matchCase"
	//
	// Optional
	//
	// NOTE: FFI_USE_MatchCase MUST be set to true to make this field effective.
	MatchCase bool

	FFI_USE_Backward  bool // for Backward.
	FFI_USE_MatchCase bool // for MatchCase.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a FindArgOptions with all fields set.
func (p FindArgOptions) FromRef(ref js.Ref) FindArgOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new FindArgOptions in the application heap.
func (p FindArgOptions) New() js.Ref {
	return bindings.FindArgOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *FindArgOptions) UpdateFrom(ref js.Ref) {
	bindings.FindArgOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *FindArgOptions) Update(ref js.Ref) {
	bindings.FindArgOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *FindArgOptions) FreeMembers(recursive bool) {
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

type IsSpatialNavigationEnabledArgCallbackFunc func(this js.Ref, spatialNavEnabled bool) js.Ref

func (fn IsSpatialNavigationEnabledArgCallbackFunc) Register() js.Func[func(spatialNavEnabled bool)] {
	return js.RegisterCallback[func(spatialNavEnabled bool)](
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
	Fn  func(arg T, this js.Ref, spatialNavEnabled bool) js.Ref
	Arg T
}

func (cb *IsSpatialNavigationEnabledArgCallback[T]) Register() js.Func[func(spatialNavEnabled bool)] {
	return js.RegisterCallback[func(spatialNavEnabled bool)](
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

type LoadDataWithBaseUrlArgCallbackFunc func(this js.Ref) js.Ref

func (fn LoadDataWithBaseUrlArgCallbackFunc) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn LoadDataWithBaseUrlArgCallbackFunc) DispatchCallback(
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

type LoadDataWithBaseUrlArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref) js.Ref
	Arg T
}

func (cb *LoadDataWithBaseUrlArgCallback[T]) Register() js.Func[func()] {
	return js.RegisterCallback[func()](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *LoadDataWithBaseUrlArgCallback[T]) DispatchCallback(
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

// MAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND returns the value of property "WEBEXT.webViewInternal.MAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND".
//
// The returned bool will be false if there is no such property.
func MAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND() (ret js.String, ok bool) {
	ok = js.True == bindings.GetMAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND(
		js.Pointer(&ret),
	)

	return
}

// SetMAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND sets the value of property "WEBEXT.webViewInternal.MAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND" to val.
//
// It returns false if the property cannot be set.
func SetMAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND(val js.String) bool {
	return js.True == bindings.SetMAX_CAPTURE_VISIBLE_REGION_CALLS_PER_SECOND(
		val.Ref())
}

type RemovalOptions struct {
	// Since is "RemovalOptions.since"
	//
	// Optional
	//
	// NOTE: FFI_USE_Since MUST be set to true to make this field effective.
	Since float64

	FFI_USE_Since bool // for Since.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a RemovalOptions with all fields set.
func (p RemovalOptions) FromRef(ref js.Ref) RemovalOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new RemovalOptions in the application heap.
func (p RemovalOptions) New() js.Ref {
	return bindings.RemovalOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *RemovalOptions) UpdateFrom(ref js.Ref) {
	bindings.RemovalOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *RemovalOptions) Update(ref js.Ref) {
	bindings.RemovalOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *RemovalOptions) FreeMembers(recursive bool) {
}

type SetPermissionAction uint32

const (
	_ SetPermissionAction = iota

	SetPermissionAction_ALLOW
	SetPermissionAction_DENY
	SetPermissionAction_DEFAULT
)

func (SetPermissionAction) FromRef(str js.Ref) SetPermissionAction {
	return SetPermissionAction(bindings.ConstOfSetPermissionAction(str))
}

func (x SetPermissionAction) String() (string, bool) {
	switch x {
	case SetPermissionAction_ALLOW:
		return "allow", true
	case SetPermissionAction_DENY:
		return "deny", true
	case SetPermissionAction_DEFAULT:
		return "default", true
	default:
		return "", false
	}
}

type SetPermissionArgCallbackFunc func(this js.Ref, allowed bool) js.Ref

func (fn SetPermissionArgCallbackFunc) Register() js.Func[func(allowed bool)] {
	return js.RegisterCallback[func(allowed bool)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn SetPermissionArgCallbackFunc) DispatchCallback(
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

type SetPermissionArgCallback[T any] struct {
	Fn  func(arg T, this js.Ref, allowed bool) js.Ref
	Arg T
}

func (cb *SetPermissionArgCallback[T]) Register() js.Func[func(allowed bool)] {
	return js.RegisterCallback[func(allowed bool)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *SetPermissionArgCallback[T]) DispatchCallback(
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

type StopFindingAction uint32

const (
	_ StopFindingAction = iota

	StopFindingAction_CLEAR
	StopFindingAction_KEEP
	StopFindingAction_ACTIVATE
)

func (StopFindingAction) FromRef(str js.Ref) StopFindingAction {
	return StopFindingAction(bindings.ConstOfStopFindingAction(str))
}

func (x StopFindingAction) String() (string, bool) {
	switch x {
	case StopFindingAction_CLEAR:
		return "clear", true
	case StopFindingAction_KEEP:
		return "keep", true
	case StopFindingAction_ACTIVATE:
		return "activate", true
	default:
		return "", false
	}
}

// HasFuncAddContentScripts returns true if the function "WEBEXT.webViewInternal.addContentScripts" exists.
func HasFuncAddContentScripts() bool {
	return js.True == bindings.HasFuncAddContentScripts()
}

// FuncAddContentScripts returns the function "WEBEXT.webViewInternal.addContentScripts".
func FuncAddContentScripts() (fn js.Func[func(instanceId int64, contentScriptList js.Array[ContentScriptDetails])]) {
	bindings.FuncAddContentScripts(
		js.Pointer(&fn),
	)
	return
}

// AddContentScripts calls the function "WEBEXT.webViewInternal.addContentScripts" directly.
func AddContentScripts(instanceId int64, contentScriptList js.Array[ContentScriptDetails]) (ret js.Void) {
	bindings.CallAddContentScripts(
		js.Pointer(&ret),
		float64(instanceId),
		contentScriptList.Ref(),
	)

	return
}

// TryAddContentScripts calls the function "WEBEXT.webViewInternal.addContentScripts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddContentScripts(instanceId int64, contentScriptList js.Array[ContentScriptDetails]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddContentScripts(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		contentScriptList.Ref(),
	)

	return
}

// HasFuncCaptureVisibleRegion returns true if the function "WEBEXT.webViewInternal.captureVisibleRegion" exists.
func HasFuncCaptureVisibleRegion() bool {
	return js.True == bindings.HasFuncCaptureVisibleRegion()
}

// FuncCaptureVisibleRegion returns the function "WEBEXT.webViewInternal.captureVisibleRegion".
func FuncCaptureVisibleRegion() (fn js.Func[func(instanceId int64, options extensiontypes.ImageDetails, callback js.Func[func(dataUrl js.String)])]) {
	bindings.FuncCaptureVisibleRegion(
		js.Pointer(&fn),
	)
	return
}

// CaptureVisibleRegion calls the function "WEBEXT.webViewInternal.captureVisibleRegion" directly.
func CaptureVisibleRegion(instanceId int64, options extensiontypes.ImageDetails, callback js.Func[func(dataUrl js.String)]) (ret js.Void) {
	bindings.CallCaptureVisibleRegion(
		js.Pointer(&ret),
		float64(instanceId),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TryCaptureVisibleRegion calls the function "WEBEXT.webViewInternal.captureVisibleRegion"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCaptureVisibleRegion(instanceId int64, options extensiontypes.ImageDetails, callback js.Func[func(dataUrl js.String)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryCaptureVisibleRegion(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// HasFuncClearData returns true if the function "WEBEXT.webViewInternal.clearData" exists.
func HasFuncClearData() bool {
	return js.True == bindings.HasFuncClearData()
}

// FuncClearData returns the function "WEBEXT.webViewInternal.clearData".
func FuncClearData() (fn js.Func[func(instanceId int64, options RemovalOptions, dataToRemove DataTypeSet, callback js.Func[func()])]) {
	bindings.FuncClearData(
		js.Pointer(&fn),
	)
	return
}

// ClearData calls the function "WEBEXT.webViewInternal.clearData" directly.
func ClearData(instanceId int64, options RemovalOptions, dataToRemove DataTypeSet, callback js.Func[func()]) (ret js.Void) {
	bindings.CallClearData(
		js.Pointer(&ret),
		float64(instanceId),
		js.Pointer(&options),
		js.Pointer(&dataToRemove),
		callback.Ref(),
	)

	return
}

// TryClearData calls the function "WEBEXT.webViewInternal.clearData"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryClearData(instanceId int64, options RemovalOptions, dataToRemove DataTypeSet, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryClearData(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		js.Pointer(&options),
		js.Pointer(&dataToRemove),
		callback.Ref(),
	)

	return
}

// HasFuncExecuteScript returns true if the function "WEBEXT.webViewInternal.executeScript" exists.
func HasFuncExecuteScript() bool {
	return js.True == bindings.HasFuncExecuteScript()
}

// FuncExecuteScript returns the function "WEBEXT.webViewInternal.executeScript".
func FuncExecuteScript() (fn js.Func[func(instanceId int64, src js.String, details extensiontypes.InjectDetails, callback js.Func[func(result js.Array[js.Any])])]) {
	bindings.FuncExecuteScript(
		js.Pointer(&fn),
	)
	return
}

// ExecuteScript calls the function "WEBEXT.webViewInternal.executeScript" directly.
func ExecuteScript(instanceId int64, src js.String, details extensiontypes.InjectDetails, callback js.Func[func(result js.Array[js.Any])]) (ret js.Void) {
	bindings.CallExecuteScript(
		js.Pointer(&ret),
		float64(instanceId),
		src.Ref(),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// TryExecuteScript calls the function "WEBEXT.webViewInternal.executeScript"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExecuteScript(instanceId int64, src js.String, details extensiontypes.InjectDetails, callback js.Func[func(result js.Array[js.Any])]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryExecuteScript(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		src.Ref(),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// HasFuncFind returns true if the function "WEBEXT.webViewInternal.find" exists.
func HasFuncFind() bool {
	return js.True == bindings.HasFuncFind()
}

// FuncFind returns the function "WEBEXT.webViewInternal.find".
func FuncFind() (fn js.Func[func(instanceId int64, searchText js.String, options FindArgOptions, callback js.Func[func(results *FindArgCallbackArgResults)])]) {
	bindings.FuncFind(
		js.Pointer(&fn),
	)
	return
}

// Find calls the function "WEBEXT.webViewInternal.find" directly.
func Find(instanceId int64, searchText js.String, options FindArgOptions, callback js.Func[func(results *FindArgCallbackArgResults)]) (ret js.Void) {
	bindings.CallFind(
		js.Pointer(&ret),
		float64(instanceId),
		searchText.Ref(),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// TryFind calls the function "WEBEXT.webViewInternal.find"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryFind(instanceId int64, searchText js.String, options FindArgOptions, callback js.Func[func(results *FindArgCallbackArgResults)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryFind(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		searchText.Ref(),
		js.Pointer(&options),
		callback.Ref(),
	)

	return
}

// HasFuncGetAudioState returns true if the function "WEBEXT.webViewInternal.getAudioState" exists.
func HasFuncGetAudioState() bool {
	return js.True == bindings.HasFuncGetAudioState()
}

// FuncGetAudioState returns the function "WEBEXT.webViewInternal.getAudioState".
func FuncGetAudioState() (fn js.Func[func(instanceId int64, callback js.Func[func(audible bool)])]) {
	bindings.FuncGetAudioState(
		js.Pointer(&fn),
	)
	return
}

// GetAudioState calls the function "WEBEXT.webViewInternal.getAudioState" directly.
func GetAudioState(instanceId int64, callback js.Func[func(audible bool)]) (ret js.Void) {
	bindings.CallGetAudioState(
		js.Pointer(&ret),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// TryGetAudioState calls the function "WEBEXT.webViewInternal.getAudioState"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetAudioState(instanceId int64, callback js.Func[func(audible bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetAudioState(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// HasFuncGetZoom returns true if the function "WEBEXT.webViewInternal.getZoom" exists.
func HasFuncGetZoom() bool {
	return js.True == bindings.HasFuncGetZoom()
}

// FuncGetZoom returns the function "WEBEXT.webViewInternal.getZoom".
func FuncGetZoom() (fn js.Func[func(instanceId int64, callback js.Func[func(zoomFactor float64)])]) {
	bindings.FuncGetZoom(
		js.Pointer(&fn),
	)
	return
}

// GetZoom calls the function "WEBEXT.webViewInternal.getZoom" directly.
func GetZoom(instanceId int64, callback js.Func[func(zoomFactor float64)]) (ret js.Void) {
	bindings.CallGetZoom(
		js.Pointer(&ret),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// TryGetZoom calls the function "WEBEXT.webViewInternal.getZoom"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetZoom(instanceId int64, callback js.Func[func(zoomFactor float64)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetZoom(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// HasFuncGetZoomMode returns true if the function "WEBEXT.webViewInternal.getZoomMode" exists.
func HasFuncGetZoomMode() bool {
	return js.True == bindings.HasFuncGetZoomMode()
}

// FuncGetZoomMode returns the function "WEBEXT.webViewInternal.getZoomMode".
func FuncGetZoomMode() (fn js.Func[func(instanceId int64, callback js.Func[func(ZoomMode ZoomMode)])]) {
	bindings.FuncGetZoomMode(
		js.Pointer(&fn),
	)
	return
}

// GetZoomMode calls the function "WEBEXT.webViewInternal.getZoomMode" directly.
func GetZoomMode(instanceId int64, callback js.Func[func(ZoomMode ZoomMode)]) (ret js.Void) {
	bindings.CallGetZoomMode(
		js.Pointer(&ret),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// TryGetZoomMode calls the function "WEBEXT.webViewInternal.getZoomMode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetZoomMode(instanceId int64, callback js.Func[func(ZoomMode ZoomMode)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetZoomMode(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// HasFuncGo returns true if the function "WEBEXT.webViewInternal.go" exists.
func HasFuncGo() bool {
	return js.True == bindings.HasFuncGo()
}

// FuncGo returns the function "WEBEXT.webViewInternal.go".
func FuncGo() (fn js.Func[func(instanceId int64, relativeIndex int64, callback js.Func[func(success bool)])]) {
	bindings.FuncGo(
		js.Pointer(&fn),
	)
	return
}

// Go calls the function "WEBEXT.webViewInternal.go" directly.
func Go(instanceId int64, relativeIndex int64, callback js.Func[func(success bool)]) (ret js.Void) {
	bindings.CallGo(
		js.Pointer(&ret),
		float64(instanceId),
		float64(relativeIndex),
		callback.Ref(),
	)

	return
}

// TryGo calls the function "WEBEXT.webViewInternal.go"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGo(instanceId int64, relativeIndex int64, callback js.Func[func(success bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryGo(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		float64(relativeIndex),
		callback.Ref(),
	)

	return
}

// HasFuncInsertCSS returns true if the function "WEBEXT.webViewInternal.insertCSS" exists.
func HasFuncInsertCSS() bool {
	return js.True == bindings.HasFuncInsertCSS()
}

// FuncInsertCSS returns the function "WEBEXT.webViewInternal.insertCSS".
func FuncInsertCSS() (fn js.Func[func(instanceId int64, src js.String, details extensiontypes.InjectDetails, callback js.Func[func()])]) {
	bindings.FuncInsertCSS(
		js.Pointer(&fn),
	)
	return
}

// InsertCSS calls the function "WEBEXT.webViewInternal.insertCSS" directly.
func InsertCSS(instanceId int64, src js.String, details extensiontypes.InjectDetails, callback js.Func[func()]) (ret js.Void) {
	bindings.CallInsertCSS(
		js.Pointer(&ret),
		float64(instanceId),
		src.Ref(),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// TryInsertCSS calls the function "WEBEXT.webViewInternal.insertCSS"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryInsertCSS(instanceId int64, src js.String, details extensiontypes.InjectDetails, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryInsertCSS(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		src.Ref(),
		js.Pointer(&details),
		callback.Ref(),
	)

	return
}

// HasFuncIsAudioMuted returns true if the function "WEBEXT.webViewInternal.isAudioMuted" exists.
func HasFuncIsAudioMuted() bool {
	return js.True == bindings.HasFuncIsAudioMuted()
}

// FuncIsAudioMuted returns the function "WEBEXT.webViewInternal.isAudioMuted".
func FuncIsAudioMuted() (fn js.Func[func(instanceId int64, callback js.Func[func(muted bool)])]) {
	bindings.FuncIsAudioMuted(
		js.Pointer(&fn),
	)
	return
}

// IsAudioMuted calls the function "WEBEXT.webViewInternal.isAudioMuted" directly.
func IsAudioMuted(instanceId int64, callback js.Func[func(muted bool)]) (ret js.Void) {
	bindings.CallIsAudioMuted(
		js.Pointer(&ret),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// TryIsAudioMuted calls the function "WEBEXT.webViewInternal.isAudioMuted"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsAudioMuted(instanceId int64, callback js.Func[func(muted bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsAudioMuted(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// HasFuncIsSpatialNavigationEnabled returns true if the function "WEBEXT.webViewInternal.isSpatialNavigationEnabled" exists.
func HasFuncIsSpatialNavigationEnabled() bool {
	return js.True == bindings.HasFuncIsSpatialNavigationEnabled()
}

// FuncIsSpatialNavigationEnabled returns the function "WEBEXT.webViewInternal.isSpatialNavigationEnabled".
func FuncIsSpatialNavigationEnabled() (fn js.Func[func(instanceId int64, callback js.Func[func(spatialNavEnabled bool)])]) {
	bindings.FuncIsSpatialNavigationEnabled(
		js.Pointer(&fn),
	)
	return
}

// IsSpatialNavigationEnabled calls the function "WEBEXT.webViewInternal.isSpatialNavigationEnabled" directly.
func IsSpatialNavigationEnabled(instanceId int64, callback js.Func[func(spatialNavEnabled bool)]) (ret js.Void) {
	bindings.CallIsSpatialNavigationEnabled(
		js.Pointer(&ret),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// TryIsSpatialNavigationEnabled calls the function "WEBEXT.webViewInternal.isSpatialNavigationEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryIsSpatialNavigationEnabled(instanceId int64, callback js.Func[func(spatialNavEnabled bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIsSpatialNavigationEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		callback.Ref(),
	)

	return
}

// HasFuncLoadDataWithBaseUrl returns true if the function "WEBEXT.webViewInternal.loadDataWithBaseUrl" exists.
func HasFuncLoadDataWithBaseUrl() bool {
	return js.True == bindings.HasFuncLoadDataWithBaseUrl()
}

// FuncLoadDataWithBaseUrl returns the function "WEBEXT.webViewInternal.loadDataWithBaseUrl".
func FuncLoadDataWithBaseUrl() (fn js.Func[func(instanceId int64, dataUrl js.String, baseUrl js.String, virtualUrl js.String, callback js.Func[func()])]) {
	bindings.FuncLoadDataWithBaseUrl(
		js.Pointer(&fn),
	)
	return
}

// LoadDataWithBaseUrl calls the function "WEBEXT.webViewInternal.loadDataWithBaseUrl" directly.
func LoadDataWithBaseUrl(instanceId int64, dataUrl js.String, baseUrl js.String, virtualUrl js.String, callback js.Func[func()]) (ret js.Void) {
	bindings.CallLoadDataWithBaseUrl(
		js.Pointer(&ret),
		float64(instanceId),
		dataUrl.Ref(),
		baseUrl.Ref(),
		virtualUrl.Ref(),
		callback.Ref(),
	)

	return
}

// TryLoadDataWithBaseUrl calls the function "WEBEXT.webViewInternal.loadDataWithBaseUrl"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryLoadDataWithBaseUrl(instanceId int64, dataUrl js.String, baseUrl js.String, virtualUrl js.String, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryLoadDataWithBaseUrl(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		dataUrl.Ref(),
		baseUrl.Ref(),
		virtualUrl.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncNavigate returns true if the function "WEBEXT.webViewInternal.navigate" exists.
func HasFuncNavigate() bool {
	return js.True == bindings.HasFuncNavigate()
}

// FuncNavigate returns the function "WEBEXT.webViewInternal.navigate".
func FuncNavigate() (fn js.Func[func(instanceId int64, src js.String)]) {
	bindings.FuncNavigate(
		js.Pointer(&fn),
	)
	return
}

// Navigate calls the function "WEBEXT.webViewInternal.navigate" directly.
func Navigate(instanceId int64, src js.String) (ret js.Void) {
	bindings.CallNavigate(
		js.Pointer(&ret),
		float64(instanceId),
		src.Ref(),
	)

	return
}

// TryNavigate calls the function "WEBEXT.webViewInternal.navigate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryNavigate(instanceId int64, src js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryNavigate(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		src.Ref(),
	)

	return
}

// HasFuncOverrideUserAgent returns true if the function "WEBEXT.webViewInternal.overrideUserAgent" exists.
func HasFuncOverrideUserAgent() bool {
	return js.True == bindings.HasFuncOverrideUserAgent()
}

// FuncOverrideUserAgent returns the function "WEBEXT.webViewInternal.overrideUserAgent".
func FuncOverrideUserAgent() (fn js.Func[func(instanceId int64, userAgentOverride js.String)]) {
	bindings.FuncOverrideUserAgent(
		js.Pointer(&fn),
	)
	return
}

// OverrideUserAgent calls the function "WEBEXT.webViewInternal.overrideUserAgent" directly.
func OverrideUserAgent(instanceId int64, userAgentOverride js.String) (ret js.Void) {
	bindings.CallOverrideUserAgent(
		js.Pointer(&ret),
		float64(instanceId),
		userAgentOverride.Ref(),
	)

	return
}

// TryOverrideUserAgent calls the function "WEBEXT.webViewInternal.overrideUserAgent"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOverrideUserAgent(instanceId int64, userAgentOverride js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOverrideUserAgent(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		userAgentOverride.Ref(),
	)

	return
}

// HasFuncReload returns true if the function "WEBEXT.webViewInternal.reload" exists.
func HasFuncReload() bool {
	return js.True == bindings.HasFuncReload()
}

// FuncReload returns the function "WEBEXT.webViewInternal.reload".
func FuncReload() (fn js.Func[func(instanceId int64)]) {
	bindings.FuncReload(
		js.Pointer(&fn),
	)
	return
}

// Reload calls the function "WEBEXT.webViewInternal.reload" directly.
func Reload(instanceId int64) (ret js.Void) {
	bindings.CallReload(
		js.Pointer(&ret),
		float64(instanceId),
	)

	return
}

// TryReload calls the function "WEBEXT.webViewInternal.reload"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryReload(instanceId int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryReload(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
	)

	return
}

// HasFuncRemoveContentScripts returns true if the function "WEBEXT.webViewInternal.removeContentScripts" exists.
func HasFuncRemoveContentScripts() bool {
	return js.True == bindings.HasFuncRemoveContentScripts()
}

// FuncRemoveContentScripts returns the function "WEBEXT.webViewInternal.removeContentScripts".
func FuncRemoveContentScripts() (fn js.Func[func(instanceId int64, scriptNameList js.Array[js.String])]) {
	bindings.FuncRemoveContentScripts(
		js.Pointer(&fn),
	)
	return
}

// RemoveContentScripts calls the function "WEBEXT.webViewInternal.removeContentScripts" directly.
func RemoveContentScripts(instanceId int64, scriptNameList js.Array[js.String]) (ret js.Void) {
	bindings.CallRemoveContentScripts(
		js.Pointer(&ret),
		float64(instanceId),
		scriptNameList.Ref(),
	)

	return
}

// TryRemoveContentScripts calls the function "WEBEXT.webViewInternal.removeContentScripts"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveContentScripts(instanceId int64, scriptNameList js.Array[js.String]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveContentScripts(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		scriptNameList.Ref(),
	)

	return
}

// HasFuncSetAllowScaling returns true if the function "WEBEXT.webViewInternal.setAllowScaling" exists.
func HasFuncSetAllowScaling() bool {
	return js.True == bindings.HasFuncSetAllowScaling()
}

// FuncSetAllowScaling returns the function "WEBEXT.webViewInternal.setAllowScaling".
func FuncSetAllowScaling() (fn js.Func[func(instanceId int64, allow bool)]) {
	bindings.FuncSetAllowScaling(
		js.Pointer(&fn),
	)
	return
}

// SetAllowScaling calls the function "WEBEXT.webViewInternal.setAllowScaling" directly.
func SetAllowScaling(instanceId int64, allow bool) (ret js.Void) {
	bindings.CallSetAllowScaling(
		js.Pointer(&ret),
		float64(instanceId),
		js.Bool(bool(allow)),
	)

	return
}

// TrySetAllowScaling calls the function "WEBEXT.webViewInternal.setAllowScaling"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAllowScaling(instanceId int64, allow bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAllowScaling(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		js.Bool(bool(allow)),
	)

	return
}

// HasFuncSetAllowTransparency returns true if the function "WEBEXT.webViewInternal.setAllowTransparency" exists.
func HasFuncSetAllowTransparency() bool {
	return js.True == bindings.HasFuncSetAllowTransparency()
}

// FuncSetAllowTransparency returns the function "WEBEXT.webViewInternal.setAllowTransparency".
func FuncSetAllowTransparency() (fn js.Func[func(instanceId int64, allow bool)]) {
	bindings.FuncSetAllowTransparency(
		js.Pointer(&fn),
	)
	return
}

// SetAllowTransparency calls the function "WEBEXT.webViewInternal.setAllowTransparency" directly.
func SetAllowTransparency(instanceId int64, allow bool) (ret js.Void) {
	bindings.CallSetAllowTransparency(
		js.Pointer(&ret),
		float64(instanceId),
		js.Bool(bool(allow)),
	)

	return
}

// TrySetAllowTransparency calls the function "WEBEXT.webViewInternal.setAllowTransparency"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAllowTransparency(instanceId int64, allow bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAllowTransparency(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		js.Bool(bool(allow)),
	)

	return
}

// HasFuncSetAudioMuted returns true if the function "WEBEXT.webViewInternal.setAudioMuted" exists.
func HasFuncSetAudioMuted() bool {
	return js.True == bindings.HasFuncSetAudioMuted()
}

// FuncSetAudioMuted returns the function "WEBEXT.webViewInternal.setAudioMuted".
func FuncSetAudioMuted() (fn js.Func[func(instanceId int64, mute bool)]) {
	bindings.FuncSetAudioMuted(
		js.Pointer(&fn),
	)
	return
}

// SetAudioMuted calls the function "WEBEXT.webViewInternal.setAudioMuted" directly.
func SetAudioMuted(instanceId int64, mute bool) (ret js.Void) {
	bindings.CallSetAudioMuted(
		js.Pointer(&ret),
		float64(instanceId),
		js.Bool(bool(mute)),
	)

	return
}

// TrySetAudioMuted calls the function "WEBEXT.webViewInternal.setAudioMuted"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetAudioMuted(instanceId int64, mute bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetAudioMuted(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		js.Bool(bool(mute)),
	)

	return
}

// HasFuncSetName returns true if the function "WEBEXT.webViewInternal.setName" exists.
func HasFuncSetName() bool {
	return js.True == bindings.HasFuncSetName()
}

// FuncSetName returns the function "WEBEXT.webViewInternal.setName".
func FuncSetName() (fn js.Func[func(instanceId int64, frameName js.String)]) {
	bindings.FuncSetName(
		js.Pointer(&fn),
	)
	return
}

// SetName calls the function "WEBEXT.webViewInternal.setName" directly.
func SetName(instanceId int64, frameName js.String) (ret js.Void) {
	bindings.CallSetName(
		js.Pointer(&ret),
		float64(instanceId),
		frameName.Ref(),
	)

	return
}

// TrySetName calls the function "WEBEXT.webViewInternal.setName"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetName(instanceId int64, frameName js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetName(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		frameName.Ref(),
	)

	return
}

// HasFuncSetPermission returns true if the function "WEBEXT.webViewInternal.setPermission" exists.
func HasFuncSetPermission() bool {
	return js.True == bindings.HasFuncSetPermission()
}

// FuncSetPermission returns the function "WEBEXT.webViewInternal.setPermission".
func FuncSetPermission() (fn js.Func[func(instanceId int64, requestId int64, action SetPermissionAction, userInput js.String, callback js.Func[func(allowed bool)])]) {
	bindings.FuncSetPermission(
		js.Pointer(&fn),
	)
	return
}

// SetPermission calls the function "WEBEXT.webViewInternal.setPermission" directly.
func SetPermission(instanceId int64, requestId int64, action SetPermissionAction, userInput js.String, callback js.Func[func(allowed bool)]) (ret js.Void) {
	bindings.CallSetPermission(
		js.Pointer(&ret),
		float64(instanceId),
		float64(requestId),
		uint32(action),
		userInput.Ref(),
		callback.Ref(),
	)

	return
}

// TrySetPermission calls the function "WEBEXT.webViewInternal.setPermission"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetPermission(instanceId int64, requestId int64, action SetPermissionAction, userInput js.String, callback js.Func[func(allowed bool)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetPermission(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		float64(requestId),
		uint32(action),
		userInput.Ref(),
		callback.Ref(),
	)

	return
}

// HasFuncSetSpatialNavigationEnabled returns true if the function "WEBEXT.webViewInternal.setSpatialNavigationEnabled" exists.
func HasFuncSetSpatialNavigationEnabled() bool {
	return js.True == bindings.HasFuncSetSpatialNavigationEnabled()
}

// FuncSetSpatialNavigationEnabled returns the function "WEBEXT.webViewInternal.setSpatialNavigationEnabled".
func FuncSetSpatialNavigationEnabled() (fn js.Func[func(instanceId int64, spatialNavEnabled bool)]) {
	bindings.FuncSetSpatialNavigationEnabled(
		js.Pointer(&fn),
	)
	return
}

// SetSpatialNavigationEnabled calls the function "WEBEXT.webViewInternal.setSpatialNavigationEnabled" directly.
func SetSpatialNavigationEnabled(instanceId int64, spatialNavEnabled bool) (ret js.Void) {
	bindings.CallSetSpatialNavigationEnabled(
		js.Pointer(&ret),
		float64(instanceId),
		js.Bool(bool(spatialNavEnabled)),
	)

	return
}

// TrySetSpatialNavigationEnabled calls the function "WEBEXT.webViewInternal.setSpatialNavigationEnabled"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetSpatialNavigationEnabled(instanceId int64, spatialNavEnabled bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetSpatialNavigationEnabled(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		js.Bool(bool(spatialNavEnabled)),
	)

	return
}

// HasFuncSetZoom returns true if the function "WEBEXT.webViewInternal.setZoom" exists.
func HasFuncSetZoom() bool {
	return js.True == bindings.HasFuncSetZoom()
}

// FuncSetZoom returns the function "WEBEXT.webViewInternal.setZoom".
func FuncSetZoom() (fn js.Func[func(instanceId int64, zoomFactor float64, callback js.Func[func()])]) {
	bindings.FuncSetZoom(
		js.Pointer(&fn),
	)
	return
}

// SetZoom calls the function "WEBEXT.webViewInternal.setZoom" directly.
func SetZoom(instanceId int64, zoomFactor float64, callback js.Func[func()]) (ret js.Void) {
	bindings.CallSetZoom(
		js.Pointer(&ret),
		float64(instanceId),
		float64(zoomFactor),
		callback.Ref(),
	)

	return
}

// TrySetZoom calls the function "WEBEXT.webViewInternal.setZoom"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetZoom(instanceId int64, zoomFactor float64, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetZoom(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		float64(zoomFactor),
		callback.Ref(),
	)

	return
}

// HasFuncSetZoomMode returns true if the function "WEBEXT.webViewInternal.setZoomMode" exists.
func HasFuncSetZoomMode() bool {
	return js.True == bindings.HasFuncSetZoomMode()
}

// FuncSetZoomMode returns the function "WEBEXT.webViewInternal.setZoomMode".
func FuncSetZoomMode() (fn js.Func[func(instanceId int64, ZoomMode ZoomMode, callback js.Func[func()])]) {
	bindings.FuncSetZoomMode(
		js.Pointer(&fn),
	)
	return
}

// SetZoomMode calls the function "WEBEXT.webViewInternal.setZoomMode" directly.
func SetZoomMode(instanceId int64, ZoomMode ZoomMode, callback js.Func[func()]) (ret js.Void) {
	bindings.CallSetZoomMode(
		js.Pointer(&ret),
		float64(instanceId),
		uint32(ZoomMode),
		callback.Ref(),
	)

	return
}

// TrySetZoomMode calls the function "WEBEXT.webViewInternal.setZoomMode"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySetZoomMode(instanceId int64, ZoomMode ZoomMode, callback js.Func[func()]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySetZoomMode(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		uint32(ZoomMode),
		callback.Ref(),
	)

	return
}

// HasFuncStop returns true if the function "WEBEXT.webViewInternal.stop" exists.
func HasFuncStop() bool {
	return js.True == bindings.HasFuncStop()
}

// FuncStop returns the function "WEBEXT.webViewInternal.stop".
func FuncStop() (fn js.Func[func(instanceId int64)]) {
	bindings.FuncStop(
		js.Pointer(&fn),
	)
	return
}

// Stop calls the function "WEBEXT.webViewInternal.stop" directly.
func Stop(instanceId int64) (ret js.Void) {
	bindings.CallStop(
		js.Pointer(&ret),
		float64(instanceId),
	)

	return
}

// TryStop calls the function "WEBEXT.webViewInternal.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStop(instanceId int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStop(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
	)

	return
}

// HasFuncStopFinding returns true if the function "WEBEXT.webViewInternal.stopFinding" exists.
func HasFuncStopFinding() bool {
	return js.True == bindings.HasFuncStopFinding()
}

// FuncStopFinding returns the function "WEBEXT.webViewInternal.stopFinding".
func FuncStopFinding() (fn js.Func[func(instanceId int64, action StopFindingAction)]) {
	bindings.FuncStopFinding(
		js.Pointer(&fn),
	)
	return
}

// StopFinding calls the function "WEBEXT.webViewInternal.stopFinding" directly.
func StopFinding(instanceId int64, action StopFindingAction) (ret js.Void) {
	bindings.CallStopFinding(
		js.Pointer(&ret),
		float64(instanceId),
		uint32(action),
	)

	return
}

// TryStopFinding calls the function "WEBEXT.webViewInternal.stopFinding"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStopFinding(instanceId int64, action StopFindingAction) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStopFinding(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
		uint32(action),
	)

	return
}

// HasFuncTerminate returns true if the function "WEBEXT.webViewInternal.terminate" exists.
func HasFuncTerminate() bool {
	return js.True == bindings.HasFuncTerminate()
}

// FuncTerminate returns the function "WEBEXT.webViewInternal.terminate".
func FuncTerminate() (fn js.Func[func(instanceId int64)]) {
	bindings.FuncTerminate(
		js.Pointer(&fn),
	)
	return
}

// Terminate calls the function "WEBEXT.webViewInternal.terminate" directly.
func Terminate(instanceId int64) (ret js.Void) {
	bindings.CallTerminate(
		js.Pointer(&ret),
		float64(instanceId),
	)

	return
}

// TryTerminate calls the function "WEBEXT.webViewInternal.terminate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryTerminate(instanceId int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryTerminate(
		js.Pointer(&ret), js.Pointer(&exception),
		float64(instanceId),
	)

	return
}

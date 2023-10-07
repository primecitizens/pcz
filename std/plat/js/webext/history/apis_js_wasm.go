// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package history

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/history/bindings"
)

type DeleteRangeArgRange struct {
	// EndTime is "DeleteRangeArgRange.endTime"
	//
	// Required
	EndTime float64
	// StartTime is "DeleteRangeArgRange.startTime"
	//
	// Required
	StartTime float64

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a DeleteRangeArgRange with all fields set.
func (p DeleteRangeArgRange) FromRef(ref js.Ref) DeleteRangeArgRange {
	p.UpdateFrom(ref)
	return p
}

// New creates a new DeleteRangeArgRange in the application heap.
func (p DeleteRangeArgRange) New() js.Ref {
	return bindings.DeleteRangeArgRangeJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *DeleteRangeArgRange) UpdateFrom(ref js.Ref) {
	bindings.DeleteRangeArgRangeJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *DeleteRangeArgRange) Update(ref js.Ref) {
	bindings.DeleteRangeArgRangeJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *DeleteRangeArgRange) FreeMembers(recursive bool) {
}

type HistoryItem struct {
	// Id is "HistoryItem.id"
	//
	// Required
	Id js.String
	// LastVisitTime is "HistoryItem.lastVisitTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_LastVisitTime MUST be set to true to make this field effective.
	LastVisitTime float64
	// Title is "HistoryItem.title"
	//
	// Optional
	Title js.String
	// TypedCount is "HistoryItem.typedCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_TypedCount MUST be set to true to make this field effective.
	TypedCount int64
	// Url is "HistoryItem.url"
	//
	// Optional
	Url js.String
	// VisitCount is "HistoryItem.visitCount"
	//
	// Optional
	//
	// NOTE: FFI_USE_VisitCount MUST be set to true to make this field effective.
	VisitCount int64

	FFI_USE_LastVisitTime bool // for LastVisitTime.
	FFI_USE_TypedCount    bool // for TypedCount.
	FFI_USE_VisitCount    bool // for VisitCount.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HistoryItem with all fields set.
func (p HistoryItem) FromRef(ref js.Ref) HistoryItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HistoryItem in the application heap.
func (p HistoryItem) New() js.Ref {
	return bindings.HistoryItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *HistoryItem) UpdateFrom(ref js.Ref) {
	bindings.HistoryItemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HistoryItem) Update(ref js.Ref) {
	bindings.HistoryItemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HistoryItem) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type OnVisitRemovedArgRemoved struct {
	// AllHistory is "OnVisitRemovedArgRemoved.allHistory"
	//
	// Required
	AllHistory bool
	// Urls is "OnVisitRemovedArgRemoved.urls"
	//
	// Optional
	Urls js.Array[js.String]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a OnVisitRemovedArgRemoved with all fields set.
func (p OnVisitRemovedArgRemoved) FromRef(ref js.Ref) OnVisitRemovedArgRemoved {
	p.UpdateFrom(ref)
	return p
}

// New creates a new OnVisitRemovedArgRemoved in the application heap.
func (p OnVisitRemovedArgRemoved) New() js.Ref {
	return bindings.OnVisitRemovedArgRemovedJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *OnVisitRemovedArgRemoved) UpdateFrom(ref js.Ref) {
	bindings.OnVisitRemovedArgRemovedJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *OnVisitRemovedArgRemoved) Update(ref js.Ref) {
	bindings.OnVisitRemovedArgRemovedJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *OnVisitRemovedArgRemoved) FreeMembers(recursive bool) {
	js.Free(
		p.Urls.Ref(),
	)
	p.Urls = p.Urls.FromRef(js.Undefined)
}

type SearchArgQuery struct {
	// EndTime is "SearchArgQuery.endTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_EndTime MUST be set to true to make this field effective.
	EndTime float64
	// MaxResults is "SearchArgQuery.maxResults"
	//
	// Optional
	//
	// NOTE: FFI_USE_MaxResults MUST be set to true to make this field effective.
	MaxResults int64
	// StartTime is "SearchArgQuery.startTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_StartTime MUST be set to true to make this field effective.
	StartTime float64
	// Text is "SearchArgQuery.text"
	//
	// Required
	Text js.String

	FFI_USE_EndTime    bool // for EndTime.
	FFI_USE_MaxResults bool // for MaxResults.
	FFI_USE_StartTime  bool // for StartTime.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a SearchArgQuery with all fields set.
func (p SearchArgQuery) FromRef(ref js.Ref) SearchArgQuery {
	p.UpdateFrom(ref)
	return p
}

// New creates a new SearchArgQuery in the application heap.
func (p SearchArgQuery) New() js.Ref {
	return bindings.SearchArgQueryJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *SearchArgQuery) UpdateFrom(ref js.Ref) {
	bindings.SearchArgQueryJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *SearchArgQuery) Update(ref js.Ref) {
	bindings.SearchArgQueryJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *SearchArgQuery) FreeMembers(recursive bool) {
	js.Free(
		p.Text.Ref(),
	)
	p.Text = p.Text.FromRef(js.Undefined)
}

type TransitionType uint32

const (
	_ TransitionType = iota

	TransitionType_LINK
	TransitionType_TYPED
	TransitionType_AUTO_BOOKMARK
	TransitionType_AUTO_SUBFRAME
	TransitionType_MANUAL_SUBFRAME
	TransitionType_GENERATED
	TransitionType_AUTO_TOPLEVEL
	TransitionType_FORM_SUBMIT
	TransitionType_RELOAD
	TransitionType_KEYWORD
	TransitionType_KEYWORD_GENERATED
)

func (TransitionType) FromRef(str js.Ref) TransitionType {
	return TransitionType(bindings.ConstOfTransitionType(str))
}

func (x TransitionType) String() (string, bool) {
	switch x {
	case TransitionType_LINK:
		return "link", true
	case TransitionType_TYPED:
		return "typed", true
	case TransitionType_AUTO_BOOKMARK:
		return "auto_bookmark", true
	case TransitionType_AUTO_SUBFRAME:
		return "auto_subframe", true
	case TransitionType_MANUAL_SUBFRAME:
		return "manual_subframe", true
	case TransitionType_GENERATED:
		return "generated", true
	case TransitionType_AUTO_TOPLEVEL:
		return "auto_toplevel", true
	case TransitionType_FORM_SUBMIT:
		return "form_submit", true
	case TransitionType_RELOAD:
		return "reload", true
	case TransitionType_KEYWORD:
		return "keyword", true
	case TransitionType_KEYWORD_GENERATED:
		return "keyword_generated", true
	default:
		return "", false
	}
}

type UrlDetails struct {
	// Url is "UrlDetails.url"
	//
	// Required
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a UrlDetails with all fields set.
func (p UrlDetails) FromRef(ref js.Ref) UrlDetails {
	p.UpdateFrom(ref)
	return p
}

// New creates a new UrlDetails in the application heap.
func (p UrlDetails) New() js.Ref {
	return bindings.UrlDetailsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *UrlDetails) UpdateFrom(ref js.Ref) {
	bindings.UrlDetailsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *UrlDetails) Update(ref js.Ref) {
	bindings.UrlDetailsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *UrlDetails) FreeMembers(recursive bool) {
	js.Free(
		p.Url.Ref(),
	)
	p.Url = p.Url.FromRef(js.Undefined)
}

type VisitItem struct {
	// Id is "VisitItem.id"
	//
	// Required
	Id js.String
	// IsLocal is "VisitItem.isLocal"
	//
	// Required
	IsLocal bool
	// ReferringVisitId is "VisitItem.referringVisitId"
	//
	// Required
	ReferringVisitId js.String
	// Transition is "VisitItem.transition"
	//
	// Required
	Transition TransitionType
	// VisitId is "VisitItem.visitId"
	//
	// Required
	VisitId js.String
	// VisitTime is "VisitItem.visitTime"
	//
	// Optional
	//
	// NOTE: FFI_USE_VisitTime MUST be set to true to make this field effective.
	VisitTime float64

	FFI_USE_VisitTime bool // for VisitTime.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a VisitItem with all fields set.
func (p VisitItem) FromRef(ref js.Ref) VisitItem {
	p.UpdateFrom(ref)
	return p
}

// New creates a new VisitItem in the application heap.
func (p VisitItem) New() js.Ref {
	return bindings.VisitItemJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *VisitItem) UpdateFrom(ref js.Ref) {
	bindings.VisitItemJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *VisitItem) Update(ref js.Ref) {
	bindings.VisitItemJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *VisitItem) FreeMembers(recursive bool) {
	js.Free(
		p.Id.Ref(),
		p.ReferringVisitId.Ref(),
		p.VisitId.Ref(),
	)
	p.Id = p.Id.FromRef(js.Undefined)
	p.ReferringVisitId = p.ReferringVisitId.FromRef(js.Undefined)
	p.VisitId = p.VisitId.FromRef(js.Undefined)
}

// HasFuncAddUrl returns true if the function "WEBEXT.history.addUrl" exists.
func HasFuncAddUrl() bool {
	return js.True == bindings.HasFuncAddUrl()
}

// FuncAddUrl returns the function "WEBEXT.history.addUrl".
func FuncAddUrl() (fn js.Func[func(details UrlDetails) js.Promise[js.Void]]) {
	bindings.FuncAddUrl(
		js.Pointer(&fn),
	)
	return
}

// AddUrl calls the function "WEBEXT.history.addUrl" directly.
func AddUrl(details UrlDetails) (ret js.Promise[js.Void]) {
	bindings.CallAddUrl(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryAddUrl calls the function "WEBEXT.history.addUrl"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryAddUrl(details UrlDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryAddUrl(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncDeleteAll returns true if the function "WEBEXT.history.deleteAll" exists.
func HasFuncDeleteAll() bool {
	return js.True == bindings.HasFuncDeleteAll()
}

// FuncDeleteAll returns the function "WEBEXT.history.deleteAll".
func FuncDeleteAll() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncDeleteAll(
		js.Pointer(&fn),
	)
	return
}

// DeleteAll calls the function "WEBEXT.history.deleteAll" directly.
func DeleteAll() (ret js.Promise[js.Void]) {
	bindings.CallDeleteAll(
		js.Pointer(&ret),
	)

	return
}

// TryDeleteAll calls the function "WEBEXT.history.deleteAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteAll() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteAll(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteRange returns true if the function "WEBEXT.history.deleteRange" exists.
func HasFuncDeleteRange() bool {
	return js.True == bindings.HasFuncDeleteRange()
}

// FuncDeleteRange returns the function "WEBEXT.history.deleteRange".
func FuncDeleteRange() (fn js.Func[func(rng DeleteRangeArgRange) js.Promise[js.Void]]) {
	bindings.FuncDeleteRange(
		js.Pointer(&fn),
	)
	return
}

// DeleteRange calls the function "WEBEXT.history.deleteRange" directly.
func DeleteRange(rng DeleteRangeArgRange) (ret js.Promise[js.Void]) {
	bindings.CallDeleteRange(
		js.Pointer(&ret),
		js.Pointer(&rng),
	)

	return
}

// TryDeleteRange calls the function "WEBEXT.history.deleteRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteRange(rng DeleteRangeArgRange) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteRange(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&rng),
	)

	return
}

// HasFuncDeleteUrl returns true if the function "WEBEXT.history.deleteUrl" exists.
func HasFuncDeleteUrl() bool {
	return js.True == bindings.HasFuncDeleteUrl()
}

// FuncDeleteUrl returns the function "WEBEXT.history.deleteUrl".
func FuncDeleteUrl() (fn js.Func[func(details UrlDetails) js.Promise[js.Void]]) {
	bindings.FuncDeleteUrl(
		js.Pointer(&fn),
	)
	return
}

// DeleteUrl calls the function "WEBEXT.history.deleteUrl" directly.
func DeleteUrl(details UrlDetails) (ret js.Promise[js.Void]) {
	bindings.CallDeleteUrl(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryDeleteUrl calls the function "WEBEXT.history.deleteUrl"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDeleteUrl(details UrlDetails) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDeleteUrl(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

// HasFuncGetVisits returns true if the function "WEBEXT.history.getVisits" exists.
func HasFuncGetVisits() bool {
	return js.True == bindings.HasFuncGetVisits()
}

// FuncGetVisits returns the function "WEBEXT.history.getVisits".
func FuncGetVisits() (fn js.Func[func(details UrlDetails) js.Promise[js.Array[VisitItem]]]) {
	bindings.FuncGetVisits(
		js.Pointer(&fn),
	)
	return
}

// GetVisits calls the function "WEBEXT.history.getVisits" directly.
func GetVisits(details UrlDetails) (ret js.Promise[js.Array[VisitItem]]) {
	bindings.CallGetVisits(
		js.Pointer(&ret),
		js.Pointer(&details),
	)

	return
}

// TryGetVisits calls the function "WEBEXT.history.getVisits"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetVisits(details UrlDetails) (ret js.Promise[js.Array[VisitItem]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetVisits(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&details),
	)

	return
}

type OnVisitRemovedEventCallbackFunc func(this js.Ref, removed *OnVisitRemovedArgRemoved) js.Ref

func (fn OnVisitRemovedEventCallbackFunc) Register() js.Func[func(removed *OnVisitRemovedArgRemoved)] {
	return js.RegisterCallback[func(removed *OnVisitRemovedArgRemoved)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnVisitRemovedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnVisitRemovedArgRemoved
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

type OnVisitRemovedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, removed *OnVisitRemovedArgRemoved) js.Ref
	Arg T
}

func (cb *OnVisitRemovedEventCallback[T]) Register() js.Func[func(removed *OnVisitRemovedArgRemoved)] {
	return js.RegisterCallback[func(removed *OnVisitRemovedArgRemoved)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnVisitRemovedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 OnVisitRemovedArgRemoved
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

// HasFuncOnVisitRemoved returns true if the function "WEBEXT.history.onVisitRemoved.addListener" exists.
func HasFuncOnVisitRemoved() bool {
	return js.True == bindings.HasFuncOnVisitRemoved()
}

// FuncOnVisitRemoved returns the function "WEBEXT.history.onVisitRemoved.addListener".
func FuncOnVisitRemoved() (fn js.Func[func(callback js.Func[func(removed *OnVisitRemovedArgRemoved)])]) {
	bindings.FuncOnVisitRemoved(
		js.Pointer(&fn),
	)
	return
}

// OnVisitRemoved calls the function "WEBEXT.history.onVisitRemoved.addListener" directly.
func OnVisitRemoved(callback js.Func[func(removed *OnVisitRemovedArgRemoved)]) (ret js.Void) {
	bindings.CallOnVisitRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnVisitRemoved calls the function "WEBEXT.history.onVisitRemoved.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnVisitRemoved(callback js.Func[func(removed *OnVisitRemovedArgRemoved)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnVisitRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffVisitRemoved returns true if the function "WEBEXT.history.onVisitRemoved.removeListener" exists.
func HasFuncOffVisitRemoved() bool {
	return js.True == bindings.HasFuncOffVisitRemoved()
}

// FuncOffVisitRemoved returns the function "WEBEXT.history.onVisitRemoved.removeListener".
func FuncOffVisitRemoved() (fn js.Func[func(callback js.Func[func(removed *OnVisitRemovedArgRemoved)])]) {
	bindings.FuncOffVisitRemoved(
		js.Pointer(&fn),
	)
	return
}

// OffVisitRemoved calls the function "WEBEXT.history.onVisitRemoved.removeListener" directly.
func OffVisitRemoved(callback js.Func[func(removed *OnVisitRemovedArgRemoved)]) (ret js.Void) {
	bindings.CallOffVisitRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffVisitRemoved calls the function "WEBEXT.history.onVisitRemoved.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffVisitRemoved(callback js.Func[func(removed *OnVisitRemovedArgRemoved)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffVisitRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnVisitRemoved returns true if the function "WEBEXT.history.onVisitRemoved.hasListener" exists.
func HasFuncHasOnVisitRemoved() bool {
	return js.True == bindings.HasFuncHasOnVisitRemoved()
}

// FuncHasOnVisitRemoved returns the function "WEBEXT.history.onVisitRemoved.hasListener".
func FuncHasOnVisitRemoved() (fn js.Func[func(callback js.Func[func(removed *OnVisitRemovedArgRemoved)]) bool]) {
	bindings.FuncHasOnVisitRemoved(
		js.Pointer(&fn),
	)
	return
}

// HasOnVisitRemoved calls the function "WEBEXT.history.onVisitRemoved.hasListener" directly.
func HasOnVisitRemoved(callback js.Func[func(removed *OnVisitRemovedArgRemoved)]) (ret bool) {
	bindings.CallHasOnVisitRemoved(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnVisitRemoved calls the function "WEBEXT.history.onVisitRemoved.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnVisitRemoved(callback js.Func[func(removed *OnVisitRemovedArgRemoved)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnVisitRemoved(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnVisitedEventCallbackFunc func(this js.Ref, result *HistoryItem) js.Ref

func (fn OnVisitedEventCallbackFunc) Register() js.Func[func(result *HistoryItem)] {
	return js.RegisterCallback[func(result *HistoryItem)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnVisitedEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 HistoryItem
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

type OnVisitedEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, result *HistoryItem) js.Ref
	Arg T
}

func (cb *OnVisitedEventCallback[T]) Register() js.Func[func(result *HistoryItem)] {
	return js.RegisterCallback[func(result *HistoryItem)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnVisitedEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 HistoryItem
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

// HasFuncOnVisited returns true if the function "WEBEXT.history.onVisited.addListener" exists.
func HasFuncOnVisited() bool {
	return js.True == bindings.HasFuncOnVisited()
}

// FuncOnVisited returns the function "WEBEXT.history.onVisited.addListener".
func FuncOnVisited() (fn js.Func[func(callback js.Func[func(result *HistoryItem)])]) {
	bindings.FuncOnVisited(
		js.Pointer(&fn),
	)
	return
}

// OnVisited calls the function "WEBEXT.history.onVisited.addListener" directly.
func OnVisited(callback js.Func[func(result *HistoryItem)]) (ret js.Void) {
	bindings.CallOnVisited(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnVisited calls the function "WEBEXT.history.onVisited.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnVisited(callback js.Func[func(result *HistoryItem)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnVisited(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffVisited returns true if the function "WEBEXT.history.onVisited.removeListener" exists.
func HasFuncOffVisited() bool {
	return js.True == bindings.HasFuncOffVisited()
}

// FuncOffVisited returns the function "WEBEXT.history.onVisited.removeListener".
func FuncOffVisited() (fn js.Func[func(callback js.Func[func(result *HistoryItem)])]) {
	bindings.FuncOffVisited(
		js.Pointer(&fn),
	)
	return
}

// OffVisited calls the function "WEBEXT.history.onVisited.removeListener" directly.
func OffVisited(callback js.Func[func(result *HistoryItem)]) (ret js.Void) {
	bindings.CallOffVisited(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffVisited calls the function "WEBEXT.history.onVisited.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffVisited(callback js.Func[func(result *HistoryItem)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffVisited(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnVisited returns true if the function "WEBEXT.history.onVisited.hasListener" exists.
func HasFuncHasOnVisited() bool {
	return js.True == bindings.HasFuncHasOnVisited()
}

// FuncHasOnVisited returns the function "WEBEXT.history.onVisited.hasListener".
func FuncHasOnVisited() (fn js.Func[func(callback js.Func[func(result *HistoryItem)]) bool]) {
	bindings.FuncHasOnVisited(
		js.Pointer(&fn),
	)
	return
}

// HasOnVisited calls the function "WEBEXT.history.onVisited.hasListener" directly.
func HasOnVisited(callback js.Func[func(result *HistoryItem)]) (ret bool) {
	bindings.CallHasOnVisited(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnVisited calls the function "WEBEXT.history.onVisited.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnVisited(callback js.Func[func(result *HistoryItem)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnVisited(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncSearch returns true if the function "WEBEXT.history.search" exists.
func HasFuncSearch() bool {
	return js.True == bindings.HasFuncSearch()
}

// FuncSearch returns the function "WEBEXT.history.search".
func FuncSearch() (fn js.Func[func(query SearchArgQuery) js.Promise[js.Array[HistoryItem]]]) {
	bindings.FuncSearch(
		js.Pointer(&fn),
	)
	return
}

// Search calls the function "WEBEXT.history.search" directly.
func Search(query SearchArgQuery) (ret js.Promise[js.Array[HistoryItem]]) {
	bindings.CallSearch(
		js.Pointer(&ret),
		js.Pointer(&query),
	)

	return
}

// TrySearch calls the function "WEBEXT.history.search"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySearch(query SearchArgQuery) (ret js.Promise[js.Array[HistoryItem]], exception js.Any, ok bool) {
	ok = js.True == bindings.TrySearch(
		js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&query),
	)

	return
}

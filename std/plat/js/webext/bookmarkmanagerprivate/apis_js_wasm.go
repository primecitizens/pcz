// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bookmarkmanagerprivate

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/mark"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/webext/bookmarkmanagerprivate/bindings"
	"github.com/primecitizens/pcz/std/plat/js/webext/bookmarks"
)

type BookmarkNodeDataElement struct {
	// Children is "BookmarkNodeDataElement.children"
	//
	// Required
	Children js.Array[BookmarkNodeDataElement]
	// Id is "BookmarkNodeDataElement.id"
	//
	// Optional
	Id js.String
	// ParentId is "BookmarkNodeDataElement.parentId"
	//
	// Optional
	ParentId js.String
	// Title is "BookmarkNodeDataElement.title"
	//
	// Required
	Title js.String
	// Url is "BookmarkNodeDataElement.url"
	//
	// Optional
	Url js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BookmarkNodeDataElement with all fields set.
func (p BookmarkNodeDataElement) FromRef(ref js.Ref) BookmarkNodeDataElement {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BookmarkNodeDataElement in the application heap.
func (p BookmarkNodeDataElement) New() js.Ref {
	return bindings.BookmarkNodeDataElementJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BookmarkNodeDataElement) UpdateFrom(ref js.Ref) {
	bindings.BookmarkNodeDataElementJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BookmarkNodeDataElement) Update(ref js.Ref) {
	bindings.BookmarkNodeDataElementJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BookmarkNodeDataElement) FreeMembers(recursive bool) {
	js.Free(
		p.Children.Ref(),
		p.Id.Ref(),
		p.ParentId.Ref(),
		p.Title.Ref(),
		p.Url.Ref(),
	)
	p.Children = p.Children.FromRef(js.Undefined)
	p.Id = p.Id.FromRef(js.Undefined)
	p.ParentId = p.ParentId.FromRef(js.Undefined)
	p.Title = p.Title.FromRef(js.Undefined)
	p.Url = p.Url.FromRef(js.Undefined)
}

type BookmarkNodeData struct {
	// Elements is "BookmarkNodeData.elements"
	//
	// Required
	Elements js.Array[BookmarkNodeDataElement]
	// SameProfile is "BookmarkNodeData.sameProfile"
	//
	// Required
	SameProfile bool

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a BookmarkNodeData with all fields set.
func (p BookmarkNodeData) FromRef(ref js.Ref) BookmarkNodeData {
	p.UpdateFrom(ref)
	return p
}

// New creates a new BookmarkNodeData in the application heap.
func (p BookmarkNodeData) New() js.Ref {
	return bindings.BookmarkNodeDataJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p *BookmarkNodeData) UpdateFrom(ref js.Ref) {
	bindings.BookmarkNodeDataJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *BookmarkNodeData) Update(ref js.Ref) {
	bindings.BookmarkNodeDataJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *BookmarkNodeData) FreeMembers(recursive bool) {
	js.Free(
		p.Elements.Ref(),
	)
	p.Elements = p.Elements.FromRef(js.Undefined)
}

// HasFuncCanPaste returns true if the function "WEBEXT.bookmarkManagerPrivate.canPaste" exists.
func HasFuncCanPaste() bool {
	return js.True == bindings.HasFuncCanPaste()
}

// FuncCanPaste returns the function "WEBEXT.bookmarkManagerPrivate.canPaste".
func FuncCanPaste() (fn js.Func[func(parentId js.String) js.Promise[js.Boolean]]) {
	bindings.FuncCanPaste(
		js.Pointer(&fn),
	)
	return
}

// CanPaste calls the function "WEBEXT.bookmarkManagerPrivate.canPaste" directly.
func CanPaste(parentId js.String) (ret js.Promise[js.Boolean]) {
	bindings.CallCanPaste(
		js.Pointer(&ret),
		parentId.Ref(),
	)

	return
}

// TryCanPaste calls the function "WEBEXT.bookmarkManagerPrivate.canPaste"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCanPaste(parentId js.String) (ret js.Promise[js.Boolean], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCanPaste(
		js.Pointer(&ret), js.Pointer(&exception),
		parentId.Ref(),
	)

	return
}

// HasFuncCopy returns true if the function "WEBEXT.bookmarkManagerPrivate.copy" exists.
func HasFuncCopy() bool {
	return js.True == bindings.HasFuncCopy()
}

// FuncCopy returns the function "WEBEXT.bookmarkManagerPrivate.copy".
func FuncCopy() (fn js.Func[func(idList js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncCopy(
		js.Pointer(&fn),
	)
	return
}

// Copy calls the function "WEBEXT.bookmarkManagerPrivate.copy" directly.
func Copy(idList js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallCopy(
		js.Pointer(&ret),
		idList.Ref(),
	)

	return
}

// TryCopy calls the function "WEBEXT.bookmarkManagerPrivate.copy"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCopy(idList js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCopy(
		js.Pointer(&ret), js.Pointer(&exception),
		idList.Ref(),
	)

	return
}

// HasFuncCut returns true if the function "WEBEXT.bookmarkManagerPrivate.cut" exists.
func HasFuncCut() bool {
	return js.True == bindings.HasFuncCut()
}

// FuncCut returns the function "WEBEXT.bookmarkManagerPrivate.cut".
func FuncCut() (fn js.Func[func(idList js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncCut(
		js.Pointer(&fn),
	)
	return
}

// Cut calls the function "WEBEXT.bookmarkManagerPrivate.cut" directly.
func Cut(idList js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallCut(
		js.Pointer(&ret),
		idList.Ref(),
	)

	return
}

// TryCut calls the function "WEBEXT.bookmarkManagerPrivate.cut"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryCut(idList js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryCut(
		js.Pointer(&ret), js.Pointer(&exception),
		idList.Ref(),
	)

	return
}

// HasFuncDrop returns true if the function "WEBEXT.bookmarkManagerPrivate.drop" exists.
func HasFuncDrop() bool {
	return js.True == bindings.HasFuncDrop()
}

// FuncDrop returns the function "WEBEXT.bookmarkManagerPrivate.drop".
func FuncDrop() (fn js.Func[func(parentId js.String, index int64) js.Promise[js.Void]]) {
	bindings.FuncDrop(
		js.Pointer(&fn),
	)
	return
}

// Drop calls the function "WEBEXT.bookmarkManagerPrivate.drop" directly.
func Drop(parentId js.String, index int64) (ret js.Promise[js.Void]) {
	bindings.CallDrop(
		js.Pointer(&ret),
		parentId.Ref(),
		float64(index),
	)

	return
}

// TryDrop calls the function "WEBEXT.bookmarkManagerPrivate.drop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryDrop(parentId js.String, index int64) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryDrop(
		js.Pointer(&ret), js.Pointer(&exception),
		parentId.Ref(),
		float64(index),
	)

	return
}

// HasFuncExport returns true if the function "WEBEXT.bookmarkManagerPrivate.export" exists.
func HasFuncExport() bool {
	return js.True == bindings.HasFuncExport()
}

// FuncExport returns the function "WEBEXT.bookmarkManagerPrivate.export".
func FuncExport() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncExport(
		js.Pointer(&fn),
	)
	return
}

// Export calls the function "WEBEXT.bookmarkManagerPrivate.export" directly.
func Export() (ret js.Promise[js.Void]) {
	bindings.CallExport(
		js.Pointer(&ret),
	)

	return
}

// TryExport calls the function "WEBEXT.bookmarkManagerPrivate.export"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryExport() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryExport(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetSubtree returns true if the function "WEBEXT.bookmarkManagerPrivate.getSubtree" exists.
func HasFuncGetSubtree() bool {
	return js.True == bindings.HasFuncGetSubtree()
}

// FuncGetSubtree returns the function "WEBEXT.bookmarkManagerPrivate.getSubtree".
func FuncGetSubtree() (fn js.Func[func(id js.String, foldersOnly bool) js.Promise[js.Array[bookmarks.BookmarkTreeNode]]]) {
	bindings.FuncGetSubtree(
		js.Pointer(&fn),
	)
	return
}

// GetSubtree calls the function "WEBEXT.bookmarkManagerPrivate.getSubtree" directly.
func GetSubtree(id js.String, foldersOnly bool) (ret js.Promise[js.Array[bookmarks.BookmarkTreeNode]]) {
	bindings.CallGetSubtree(
		js.Pointer(&ret),
		id.Ref(),
		js.Bool(bool(foldersOnly)),
	)

	return
}

// TryGetSubtree calls the function "WEBEXT.bookmarkManagerPrivate.getSubtree"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryGetSubtree(id js.String, foldersOnly bool) (ret js.Promise[js.Array[bookmarks.BookmarkTreeNode]], exception js.Any, ok bool) {
	ok = js.True == bindings.TryGetSubtree(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Bool(bool(foldersOnly)),
	)

	return
}

// HasFuncImport returns true if the function "WEBEXT.bookmarkManagerPrivate.import" exists.
func HasFuncImport() bool {
	return js.True == bindings.HasFuncImport()
}

// FuncImport returns the function "WEBEXT.bookmarkManagerPrivate.import".
func FuncImport() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncImport(
		js.Pointer(&fn),
	)
	return
}

// Import calls the function "WEBEXT.bookmarkManagerPrivate.import" directly.
func Import() (ret js.Promise[js.Void]) {
	bindings.CallImport(
		js.Pointer(&ret),
	)

	return
}

// TryImport calls the function "WEBEXT.bookmarkManagerPrivate.import"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryImport() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryImport(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type OnDragEnterEventCallbackFunc func(this js.Ref, bookmarkNodeData *BookmarkNodeData) js.Ref

func (fn OnDragEnterEventCallbackFunc) Register() js.Func[func(bookmarkNodeData *BookmarkNodeData)] {
	return js.RegisterCallback[func(bookmarkNodeData *BookmarkNodeData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDragEnterEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BookmarkNodeData
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

type OnDragEnterEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, bookmarkNodeData *BookmarkNodeData) js.Ref
	Arg T
}

func (cb *OnDragEnterEventCallback[T]) Register() js.Func[func(bookmarkNodeData *BookmarkNodeData)] {
	return js.RegisterCallback[func(bookmarkNodeData *BookmarkNodeData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDragEnterEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BookmarkNodeData
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

// HasFuncOnDragEnter returns true if the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.addListener" exists.
func HasFuncOnDragEnter() bool {
	return js.True == bindings.HasFuncOnDragEnter()
}

// FuncOnDragEnter returns the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.addListener".
func FuncOnDragEnter() (fn js.Func[func(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)])]) {
	bindings.FuncOnDragEnter(
		js.Pointer(&fn),
	)
	return
}

// OnDragEnter calls the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.addListener" directly.
func OnDragEnter(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void) {
	bindings.CallOnDragEnter(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDragEnter calls the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDragEnter(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDragEnter(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDragEnter returns true if the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.removeListener" exists.
func HasFuncOffDragEnter() bool {
	return js.True == bindings.HasFuncOffDragEnter()
}

// FuncOffDragEnter returns the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.removeListener".
func FuncOffDragEnter() (fn js.Func[func(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)])]) {
	bindings.FuncOffDragEnter(
		js.Pointer(&fn),
	)
	return
}

// OffDragEnter calls the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.removeListener" directly.
func OffDragEnter(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void) {
	bindings.CallOffDragEnter(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDragEnter calls the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDragEnter(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDragEnter(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDragEnter returns true if the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.hasListener" exists.
func HasFuncHasOnDragEnter() bool {
	return js.True == bindings.HasFuncHasOnDragEnter()
}

// FuncHasOnDragEnter returns the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.hasListener".
func FuncHasOnDragEnter() (fn js.Func[func(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) bool]) {
	bindings.FuncHasOnDragEnter(
		js.Pointer(&fn),
	)
	return
}

// HasOnDragEnter calls the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.hasListener" directly.
func HasOnDragEnter(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret bool) {
	bindings.CallHasOnDragEnter(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDragEnter calls the function "WEBEXT.bookmarkManagerPrivate.onDragEnter.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDragEnter(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDragEnter(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDragLeaveEventCallbackFunc func(this js.Ref, bookmarkNodeData *BookmarkNodeData) js.Ref

func (fn OnDragLeaveEventCallbackFunc) Register() js.Func[func(bookmarkNodeData *BookmarkNodeData)] {
	return js.RegisterCallback[func(bookmarkNodeData *BookmarkNodeData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDragLeaveEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BookmarkNodeData
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

type OnDragLeaveEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, bookmarkNodeData *BookmarkNodeData) js.Ref
	Arg T
}

func (cb *OnDragLeaveEventCallback[T]) Register() js.Func[func(bookmarkNodeData *BookmarkNodeData)] {
	return js.RegisterCallback[func(bookmarkNodeData *BookmarkNodeData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDragLeaveEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BookmarkNodeData
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

// HasFuncOnDragLeave returns true if the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.addListener" exists.
func HasFuncOnDragLeave() bool {
	return js.True == bindings.HasFuncOnDragLeave()
}

// FuncOnDragLeave returns the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.addListener".
func FuncOnDragLeave() (fn js.Func[func(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)])]) {
	bindings.FuncOnDragLeave(
		js.Pointer(&fn),
	)
	return
}

// OnDragLeave calls the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.addListener" directly.
func OnDragLeave(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void) {
	bindings.CallOnDragLeave(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDragLeave calls the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDragLeave(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDragLeave(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDragLeave returns true if the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.removeListener" exists.
func HasFuncOffDragLeave() bool {
	return js.True == bindings.HasFuncOffDragLeave()
}

// FuncOffDragLeave returns the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.removeListener".
func FuncOffDragLeave() (fn js.Func[func(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)])]) {
	bindings.FuncOffDragLeave(
		js.Pointer(&fn),
	)
	return
}

// OffDragLeave calls the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.removeListener" directly.
func OffDragLeave(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void) {
	bindings.CallOffDragLeave(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDragLeave calls the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDragLeave(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDragLeave(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDragLeave returns true if the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.hasListener" exists.
func HasFuncHasOnDragLeave() bool {
	return js.True == bindings.HasFuncHasOnDragLeave()
}

// FuncHasOnDragLeave returns the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.hasListener".
func FuncHasOnDragLeave() (fn js.Func[func(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) bool]) {
	bindings.FuncHasOnDragLeave(
		js.Pointer(&fn),
	)
	return
}

// HasOnDragLeave calls the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.hasListener" directly.
func HasOnDragLeave(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret bool) {
	bindings.CallHasOnDragLeave(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDragLeave calls the function "WEBEXT.bookmarkManagerPrivate.onDragLeave.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDragLeave(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDragLeave(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

type OnDropEventCallbackFunc func(this js.Ref, bookmarkNodeData *BookmarkNodeData) js.Ref

func (fn OnDropEventCallbackFunc) Register() js.Func[func(bookmarkNodeData *BookmarkNodeData)] {
	return js.RegisterCallback[func(bookmarkNodeData *BookmarkNodeData)](
		fn, abi.FuncPCABIInternal(fn),
	)
}

func (fn OnDropEventCallbackFunc) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BookmarkNodeData
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

type OnDropEventCallback[T any] struct {
	Fn  func(arg T, this js.Ref, bookmarkNodeData *BookmarkNodeData) js.Ref
	Arg T
}

func (cb *OnDropEventCallback[T]) Register() js.Func[func(bookmarkNodeData *BookmarkNodeData)] {
	return js.RegisterCallback[func(bookmarkNodeData *BookmarkNodeData)](
		cb, abi.FuncPCABIInternal(cb.Fn),
	)
}

func (cb *OnDropEventCallback[T]) DispatchCallback(
	targetPC uintptr, ctx *js.CallbackContext,
) {
	args := ctx.Args()
	if len(args) != 1+1 /* js this */ ||
		targetPC != uintptr(abi.FuncPCABIInternal(cb.Fn)) {
		js.ThrowInvalidCallbackInvocation()
	}
	var arg0 BookmarkNodeData
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

// HasFuncOnDrop returns true if the function "WEBEXT.bookmarkManagerPrivate.onDrop.addListener" exists.
func HasFuncOnDrop() bool {
	return js.True == bindings.HasFuncOnDrop()
}

// FuncOnDrop returns the function "WEBEXT.bookmarkManagerPrivate.onDrop.addListener".
func FuncOnDrop() (fn js.Func[func(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)])]) {
	bindings.FuncOnDrop(
		js.Pointer(&fn),
	)
	return
}

// OnDrop calls the function "WEBEXT.bookmarkManagerPrivate.onDrop.addListener" directly.
func OnDrop(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void) {
	bindings.CallOnDrop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOnDrop calls the function "WEBEXT.bookmarkManagerPrivate.onDrop.addListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOnDrop(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOnDrop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOffDrop returns true if the function "WEBEXT.bookmarkManagerPrivate.onDrop.removeListener" exists.
func HasFuncOffDrop() bool {
	return js.True == bindings.HasFuncOffDrop()
}

// FuncOffDrop returns the function "WEBEXT.bookmarkManagerPrivate.onDrop.removeListener".
func FuncOffDrop() (fn js.Func[func(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)])]) {
	bindings.FuncOffDrop(
		js.Pointer(&fn),
	)
	return
}

// OffDrop calls the function "WEBEXT.bookmarkManagerPrivate.onDrop.removeListener" directly.
func OffDrop(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void) {
	bindings.CallOffDrop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryOffDrop calls the function "WEBEXT.bookmarkManagerPrivate.onDrop.removeListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOffDrop(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOffDrop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncHasOnDrop returns true if the function "WEBEXT.bookmarkManagerPrivate.onDrop.hasListener" exists.
func HasFuncHasOnDrop() bool {
	return js.True == bindings.HasFuncHasOnDrop()
}

// FuncHasOnDrop returns the function "WEBEXT.bookmarkManagerPrivate.onDrop.hasListener".
func FuncHasOnDrop() (fn js.Func[func(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) bool]) {
	bindings.FuncHasOnDrop(
		js.Pointer(&fn),
	)
	return
}

// HasOnDrop calls the function "WEBEXT.bookmarkManagerPrivate.onDrop.hasListener" directly.
func HasOnDrop(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret bool) {
	bindings.CallHasOnDrop(
		js.Pointer(&ret),
		callback.Ref(),
	)

	return
}

// TryHasOnDrop calls the function "WEBEXT.bookmarkManagerPrivate.onDrop.hasListener"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryHasOnDrop(callback js.Func[func(bookmarkNodeData *BookmarkNodeData)]) (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHasOnDrop(
		js.Pointer(&ret), js.Pointer(&exception),
		callback.Ref(),
	)

	return
}

// HasFuncOpenInNewTab returns true if the function "WEBEXT.bookmarkManagerPrivate.openInNewTab" exists.
func HasFuncOpenInNewTab() bool {
	return js.True == bindings.HasFuncOpenInNewTab()
}

// FuncOpenInNewTab returns the function "WEBEXT.bookmarkManagerPrivate.openInNewTab".
func FuncOpenInNewTab() (fn js.Func[func(id js.String, active bool)]) {
	bindings.FuncOpenInNewTab(
		js.Pointer(&fn),
	)
	return
}

// OpenInNewTab calls the function "WEBEXT.bookmarkManagerPrivate.openInNewTab" directly.
func OpenInNewTab(id js.String, active bool) (ret js.Void) {
	bindings.CallOpenInNewTab(
		js.Pointer(&ret),
		id.Ref(),
		js.Bool(bool(active)),
	)

	return
}

// TryOpenInNewTab calls the function "WEBEXT.bookmarkManagerPrivate.openInNewTab"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenInNewTab(id js.String, active bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenInNewTab(
		js.Pointer(&ret), js.Pointer(&exception),
		id.Ref(),
		js.Bool(bool(active)),
	)

	return
}

// HasFuncOpenInNewWindow returns true if the function "WEBEXT.bookmarkManagerPrivate.openInNewWindow" exists.
func HasFuncOpenInNewWindow() bool {
	return js.True == bindings.HasFuncOpenInNewWindow()
}

// FuncOpenInNewWindow returns the function "WEBEXT.bookmarkManagerPrivate.openInNewWindow".
func FuncOpenInNewWindow() (fn js.Func[func(idList js.Array[js.String], incognito bool)]) {
	bindings.FuncOpenInNewWindow(
		js.Pointer(&fn),
	)
	return
}

// OpenInNewWindow calls the function "WEBEXT.bookmarkManagerPrivate.openInNewWindow" directly.
func OpenInNewWindow(idList js.Array[js.String], incognito bool) (ret js.Void) {
	bindings.CallOpenInNewWindow(
		js.Pointer(&ret),
		idList.Ref(),
		js.Bool(bool(incognito)),
	)

	return
}

// TryOpenInNewWindow calls the function "WEBEXT.bookmarkManagerPrivate.openInNewWindow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryOpenInNewWindow(idList js.Array[js.String], incognito bool) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryOpenInNewWindow(
		js.Pointer(&ret), js.Pointer(&exception),
		idList.Ref(),
		js.Bool(bool(incognito)),
	)

	return
}

// HasFuncPaste returns true if the function "WEBEXT.bookmarkManagerPrivate.paste" exists.
func HasFuncPaste() bool {
	return js.True == bindings.HasFuncPaste()
}

// FuncPaste returns the function "WEBEXT.bookmarkManagerPrivate.paste".
func FuncPaste() (fn js.Func[func(parentId js.String, selectedIdList js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncPaste(
		js.Pointer(&fn),
	)
	return
}

// Paste calls the function "WEBEXT.bookmarkManagerPrivate.paste" directly.
func Paste(parentId js.String, selectedIdList js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallPaste(
		js.Pointer(&ret),
		parentId.Ref(),
		selectedIdList.Ref(),
	)

	return
}

// TryPaste calls the function "WEBEXT.bookmarkManagerPrivate.paste"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryPaste(parentId js.String, selectedIdList js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryPaste(
		js.Pointer(&ret), js.Pointer(&exception),
		parentId.Ref(),
		selectedIdList.Ref(),
	)

	return
}

// HasFuncRedo returns true if the function "WEBEXT.bookmarkManagerPrivate.redo" exists.
func HasFuncRedo() bool {
	return js.True == bindings.HasFuncRedo()
}

// FuncRedo returns the function "WEBEXT.bookmarkManagerPrivate.redo".
func FuncRedo() (fn js.Func[func()]) {
	bindings.FuncRedo(
		js.Pointer(&fn),
	)
	return
}

// Redo calls the function "WEBEXT.bookmarkManagerPrivate.redo" directly.
func Redo() (ret js.Void) {
	bindings.CallRedo(
		js.Pointer(&ret),
	)

	return
}

// TryRedo calls the function "WEBEXT.bookmarkManagerPrivate.redo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRedo() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryRedo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemoveTrees returns true if the function "WEBEXT.bookmarkManagerPrivate.removeTrees" exists.
func HasFuncRemoveTrees() bool {
	return js.True == bindings.HasFuncRemoveTrees()
}

// FuncRemoveTrees returns the function "WEBEXT.bookmarkManagerPrivate.removeTrees".
func FuncRemoveTrees() (fn js.Func[func(idList js.Array[js.String]) js.Promise[js.Void]]) {
	bindings.FuncRemoveTrees(
		js.Pointer(&fn),
	)
	return
}

// RemoveTrees calls the function "WEBEXT.bookmarkManagerPrivate.removeTrees" directly.
func RemoveTrees(idList js.Array[js.String]) (ret js.Promise[js.Void]) {
	bindings.CallRemoveTrees(
		js.Pointer(&ret),
		idList.Ref(),
	)

	return
}

// TryRemoveTrees calls the function "WEBEXT.bookmarkManagerPrivate.removeTrees"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryRemoveTrees(idList js.Array[js.String]) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryRemoveTrees(
		js.Pointer(&ret), js.Pointer(&exception),
		idList.Ref(),
	)

	return
}

// HasFuncSortChildren returns true if the function "WEBEXT.bookmarkManagerPrivate.sortChildren" exists.
func HasFuncSortChildren() bool {
	return js.True == bindings.HasFuncSortChildren()
}

// FuncSortChildren returns the function "WEBEXT.bookmarkManagerPrivate.sortChildren".
func FuncSortChildren() (fn js.Func[func(parentId js.String)]) {
	bindings.FuncSortChildren(
		js.Pointer(&fn),
	)
	return
}

// SortChildren calls the function "WEBEXT.bookmarkManagerPrivate.sortChildren" directly.
func SortChildren(parentId js.String) (ret js.Void) {
	bindings.CallSortChildren(
		js.Pointer(&ret),
		parentId.Ref(),
	)

	return
}

// TrySortChildren calls the function "WEBEXT.bookmarkManagerPrivate.sortChildren"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TrySortChildren(parentId js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TrySortChildren(
		js.Pointer(&ret), js.Pointer(&exception),
		parentId.Ref(),
	)

	return
}

// HasFuncStartDrag returns true if the function "WEBEXT.bookmarkManagerPrivate.startDrag" exists.
func HasFuncStartDrag() bool {
	return js.True == bindings.HasFuncStartDrag()
}

// FuncStartDrag returns the function "WEBEXT.bookmarkManagerPrivate.startDrag".
func FuncStartDrag() (fn js.Func[func(idList js.Array[js.String], dragNodeIndex int64, isFromTouch bool, x int64, y int64)]) {
	bindings.FuncStartDrag(
		js.Pointer(&fn),
	)
	return
}

// StartDrag calls the function "WEBEXT.bookmarkManagerPrivate.startDrag" directly.
func StartDrag(idList js.Array[js.String], dragNodeIndex int64, isFromTouch bool, x int64, y int64) (ret js.Void) {
	bindings.CallStartDrag(
		js.Pointer(&ret),
		idList.Ref(),
		float64(dragNodeIndex),
		js.Bool(bool(isFromTouch)),
		float64(x),
		float64(y),
	)

	return
}

// TryStartDrag calls the function "WEBEXT.bookmarkManagerPrivate.startDrag"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryStartDrag(idList js.Array[js.String], dragNodeIndex int64, isFromTouch bool, x int64, y int64) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryStartDrag(
		js.Pointer(&ret), js.Pointer(&exception),
		idList.Ref(),
		float64(dragNodeIndex),
		js.Bool(bool(isFromTouch)),
		float64(x),
		float64(y),
	)

	return
}

// HasFuncUndo returns true if the function "WEBEXT.bookmarkManagerPrivate.undo" exists.
func HasFuncUndo() bool {
	return js.True == bindings.HasFuncUndo()
}

// FuncUndo returns the function "WEBEXT.bookmarkManagerPrivate.undo".
func FuncUndo() (fn js.Func[func()]) {
	bindings.FuncUndo(
		js.Pointer(&fn),
	)
	return
}

// Undo calls the function "WEBEXT.bookmarkManagerPrivate.undo" directly.
func Undo() (ret js.Void) {
	bindings.CallUndo(
		js.Pointer(&ret),
	)

	return
}

// TryUndo calls the function "WEBEXT.bookmarkManagerPrivate.undo"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func TryUndo() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryUndo(
		js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

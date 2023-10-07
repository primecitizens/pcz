// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

type (
	_ unsafe.Pointer
	_ js.Ref
)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate store_BookmarkNodeDataElement
//go:noescape
func BookmarkNodeDataElementJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate load_BookmarkNodeDataElement
//go:noescape
func BookmarkNodeDataElementJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate store_BookmarkNodeData
//go:noescape
func BookmarkNodeDataJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate load_BookmarkNodeData
//go:noescape
func BookmarkNodeDataJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_CanPaste
//go:noescape
func HasFuncCanPaste() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_CanPaste
//go:noescape
func FuncCanPaste(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_CanPaste
//go:noescape
func CallCanPaste(
	retPtr unsafe.Pointer,
	parentId js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_CanPaste
//go:noescape
func TryCanPaste(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parentId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_Copy
//go:noescape
func HasFuncCopy() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_Copy
//go:noescape
func FuncCopy(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_Copy
//go:noescape
func CallCopy(
	retPtr unsafe.Pointer,
	idList js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_Copy
//go:noescape
func TryCopy(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	idList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_Cut
//go:noescape
func HasFuncCut() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_Cut
//go:noescape
func FuncCut(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_Cut
//go:noescape
func CallCut(
	retPtr unsafe.Pointer,
	idList js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_Cut
//go:noescape
func TryCut(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	idList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_Drop
//go:noescape
func HasFuncDrop() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_Drop
//go:noescape
func FuncDrop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_Drop
//go:noescape
func CallDrop(
	retPtr unsafe.Pointer,
	parentId js.Ref,
	index float64)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_Drop
//go:noescape
func TryDrop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parentId js.Ref,
	index float64) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_Export
//go:noescape
func HasFuncExport() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_Export
//go:noescape
func FuncExport(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_Export
//go:noescape
func CallExport(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_Export
//go:noescape
func TryExport(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_GetSubtree
//go:noescape
func HasFuncGetSubtree() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_GetSubtree
//go:noescape
func FuncGetSubtree(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_GetSubtree
//go:noescape
func CallGetSubtree(
	retPtr unsafe.Pointer,
	id js.Ref,
	foldersOnly js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_GetSubtree
//go:noescape
func TryGetSubtree(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	foldersOnly js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_Import
//go:noescape
func HasFuncImport() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_Import
//go:noescape
func FuncImport(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_Import
//go:noescape
func CallImport(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_Import
//go:noescape
func TryImport(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_OnDragEnter
//go:noescape
func HasFuncOnDragEnter() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_OnDragEnter
//go:noescape
func FuncOnDragEnter(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_OnDragEnter
//go:noescape
func CallOnDragEnter(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_OnDragEnter
//go:noescape
func TryOnDragEnter(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_OffDragEnter
//go:noescape
func HasFuncOffDragEnter() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_OffDragEnter
//go:noescape
func FuncOffDragEnter(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_OffDragEnter
//go:noescape
func CallOffDragEnter(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_OffDragEnter
//go:noescape
func TryOffDragEnter(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_HasOnDragEnter
//go:noescape
func HasFuncHasOnDragEnter() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_HasOnDragEnter
//go:noescape
func FuncHasOnDragEnter(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_HasOnDragEnter
//go:noescape
func CallHasOnDragEnter(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_HasOnDragEnter
//go:noescape
func TryHasOnDragEnter(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_OnDragLeave
//go:noescape
func HasFuncOnDragLeave() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_OnDragLeave
//go:noescape
func FuncOnDragLeave(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_OnDragLeave
//go:noescape
func CallOnDragLeave(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_OnDragLeave
//go:noescape
func TryOnDragLeave(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_OffDragLeave
//go:noescape
func HasFuncOffDragLeave() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_OffDragLeave
//go:noescape
func FuncOffDragLeave(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_OffDragLeave
//go:noescape
func CallOffDragLeave(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_OffDragLeave
//go:noescape
func TryOffDragLeave(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_HasOnDragLeave
//go:noescape
func HasFuncHasOnDragLeave() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_HasOnDragLeave
//go:noescape
func FuncHasOnDragLeave(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_HasOnDragLeave
//go:noescape
func CallHasOnDragLeave(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_HasOnDragLeave
//go:noescape
func TryHasOnDragLeave(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_OnDrop
//go:noescape
func HasFuncOnDrop() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_OnDrop
//go:noescape
func FuncOnDrop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_OnDrop
//go:noescape
func CallOnDrop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_OnDrop
//go:noescape
func TryOnDrop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_OffDrop
//go:noescape
func HasFuncOffDrop() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_OffDrop
//go:noescape
func FuncOffDrop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_OffDrop
//go:noescape
func CallOffDrop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_OffDrop
//go:noescape
func TryOffDrop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_HasOnDrop
//go:noescape
func HasFuncHasOnDrop() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_HasOnDrop
//go:noescape
func FuncHasOnDrop(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_HasOnDrop
//go:noescape
func CallHasOnDrop(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_HasOnDrop
//go:noescape
func TryHasOnDrop(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_OpenInNewTab
//go:noescape
func HasFuncOpenInNewTab() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_OpenInNewTab
//go:noescape
func FuncOpenInNewTab(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_OpenInNewTab
//go:noescape
func CallOpenInNewTab(
	retPtr unsafe.Pointer,
	id js.Ref,
	active js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_OpenInNewTab
//go:noescape
func TryOpenInNewTab(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	active js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_OpenInNewWindow
//go:noescape
func HasFuncOpenInNewWindow() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_OpenInNewWindow
//go:noescape
func FuncOpenInNewWindow(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_OpenInNewWindow
//go:noescape
func CallOpenInNewWindow(
	retPtr unsafe.Pointer,
	idList js.Ref,
	incognito js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_OpenInNewWindow
//go:noescape
func TryOpenInNewWindow(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	idList js.Ref,
	incognito js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_Paste
//go:noescape
func HasFuncPaste() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_Paste
//go:noescape
func FuncPaste(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_Paste
//go:noescape
func CallPaste(
	retPtr unsafe.Pointer,
	parentId js.Ref,
	selectedIdList js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_Paste
//go:noescape
func TryPaste(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parentId js.Ref,
	selectedIdList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_Redo
//go:noescape
func HasFuncRedo() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_Redo
//go:noescape
func FuncRedo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_Redo
//go:noescape
func CallRedo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_Redo
//go:noescape
func TryRedo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_RemoveTrees
//go:noescape
func HasFuncRemoveTrees() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_RemoveTrees
//go:noescape
func FuncRemoveTrees(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_RemoveTrees
//go:noescape
func CallRemoveTrees(
	retPtr unsafe.Pointer,
	idList js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_RemoveTrees
//go:noescape
func TryRemoveTrees(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	idList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_SortChildren
//go:noescape
func HasFuncSortChildren() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_SortChildren
//go:noescape
func FuncSortChildren(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_SortChildren
//go:noescape
func CallSortChildren(
	retPtr unsafe.Pointer,
	parentId js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_SortChildren
//go:noescape
func TrySortChildren(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	parentId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_StartDrag
//go:noescape
func HasFuncStartDrag() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_StartDrag
//go:noescape
func FuncStartDrag(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_StartDrag
//go:noescape
func CallStartDrag(
	retPtr unsafe.Pointer,
	idList js.Ref,
	dragNodeIndex float64,
	isFromTouch js.Ref,
	x float64,
	y float64)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_StartDrag
//go:noescape
func TryStartDrag(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	idList js.Ref,
	dragNodeIndex float64,
	isFromTouch js.Ref,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate has_Undo
//go:noescape
func HasFuncUndo() js.Ref

//go:wasmimport plat/js/webext/bookmarkmanagerprivate func_Undo
//go:noescape
func FuncUndo(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate call_Undo
//go:noescape
func CallUndo(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarkmanagerprivate try_Undo
//go:noescape
func TryUndo(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

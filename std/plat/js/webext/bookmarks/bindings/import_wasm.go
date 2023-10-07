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

//go:wasmimport plat/js/webext/bookmarks constof_BookmarkTreeNodeUnmodifiable
//go:noescape
func ConstOfBookmarkTreeNodeUnmodifiable(str js.Ref) uint32

//go:wasmimport plat/js/webext/bookmarks store_BookmarkTreeNode
//go:noescape
func BookmarkTreeNodeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bookmarks load_BookmarkTreeNode
//go:noescape
func BookmarkTreeNodeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarks store_CreateDetails
//go:noescape
func CreateDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bookmarks load_CreateDetails
//go:noescape
func CreateDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarks get_MAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE
//go:noescape
func GetMAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/bookmarks set_MAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE
//go:noescape
func SetMAX_SUSTAINED_WRITE_OPERATIONS_PER_MINUTE(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarks get_MAX_WRITE_OPERATIONS_PER_HOUR
//go:noescape
func GetMAX_WRITE_OPERATIONS_PER_HOUR(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/bookmarks set_MAX_WRITE_OPERATIONS_PER_HOUR
//go:noescape
func SetMAX_WRITE_OPERATIONS_PER_HOUR(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarks store_MoveArgDestination
//go:noescape
func MoveArgDestinationJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bookmarks load_MoveArgDestination
//go:noescape
func MoveArgDestinationJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarks store_OnChangedArgChangeInfo
//go:noescape
func OnChangedArgChangeInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bookmarks load_OnChangedArgChangeInfo
//go:noescape
func OnChangedArgChangeInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarks store_OnChildrenReorderedArgReorderInfo
//go:noescape
func OnChildrenReorderedArgReorderInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bookmarks load_OnChildrenReorderedArgReorderInfo
//go:noescape
func OnChildrenReorderedArgReorderInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarks store_OnMovedArgMoveInfo
//go:noescape
func OnMovedArgMoveInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bookmarks load_OnMovedArgMoveInfo
//go:noescape
func OnMovedArgMoveInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarks store_OnRemovedArgRemoveInfo
//go:noescape
func OnRemovedArgRemoveInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bookmarks load_OnRemovedArgRemoveInfo
//go:noescape
func OnRemovedArgRemoveInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarks store_SearchArgQueryChoice1
//go:noescape
func SearchArgQueryChoice1JSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bookmarks load_SearchArgQueryChoice1
//go:noescape
func SearchArgQueryChoice1JSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarks store_UpdateArgChanges
//go:noescape
func UpdateArgChangesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/bookmarks load_UpdateArgChanges
//go:noescape
func UpdateArgChangesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/bookmarks has_Create
//go:noescape
func HasFuncCreate() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_Create
//go:noescape
func FuncCreate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_Create
//go:noescape
func CallCreate(
	retPtr unsafe.Pointer,
	bookmark unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks try_Create
//go:noescape
func TryCreate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	bookmark unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer,
	idOrIdList js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	idOrIdList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_GetChildren
//go:noescape
func HasFuncGetChildren() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_GetChildren
//go:noescape
func FuncGetChildren(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_GetChildren
//go:noescape
func CallGetChildren(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_GetChildren
//go:noescape
func TryGetChildren(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_GetRecent
//go:noescape
func HasFuncGetRecent() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_GetRecent
//go:noescape
func FuncGetRecent(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_GetRecent
//go:noescape
func CallGetRecent(
	retPtr unsafe.Pointer,
	numberOfItems float64)

//go:wasmimport plat/js/webext/bookmarks try_GetRecent
//go:noescape
func TryGetRecent(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	numberOfItems float64) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_GetSubTree
//go:noescape
func HasFuncGetSubTree() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_GetSubTree
//go:noescape
func FuncGetSubTree(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_GetSubTree
//go:noescape
func CallGetSubTree(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_GetSubTree
//go:noescape
func TryGetSubTree(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_GetTree
//go:noescape
func HasFuncGetTree() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_GetTree
//go:noescape
func FuncGetTree(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_GetTree
//go:noescape
func CallGetTree(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks try_GetTree
//go:noescape
func TryGetTree(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_Move
//go:noescape
func HasFuncMove() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_Move
//go:noescape
func FuncMove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_Move
//go:noescape
func CallMove(
	retPtr unsafe.Pointer,
	id js.Ref,
	destination unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks try_Move
//go:noescape
func TryMove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	destination unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OnChanged
//go:noescape
func HasFuncOnChanged() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OnChanged
//go:noescape
func FuncOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OnChanged
//go:noescape
func CallOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OnChanged
//go:noescape
func TryOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OffChanged
//go:noescape
func HasFuncOffChanged() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OffChanged
//go:noescape
func FuncOffChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OffChanged
//go:noescape
func CallOffChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OffChanged
//go:noescape
func TryOffChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_HasOnChanged
//go:noescape
func HasFuncHasOnChanged() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_HasOnChanged
//go:noescape
func FuncHasOnChanged(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_HasOnChanged
//go:noescape
func CallHasOnChanged(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_HasOnChanged
//go:noescape
func TryHasOnChanged(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OnChildrenReordered
//go:noescape
func HasFuncOnChildrenReordered() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OnChildrenReordered
//go:noescape
func FuncOnChildrenReordered(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OnChildrenReordered
//go:noescape
func CallOnChildrenReordered(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OnChildrenReordered
//go:noescape
func TryOnChildrenReordered(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OffChildrenReordered
//go:noescape
func HasFuncOffChildrenReordered() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OffChildrenReordered
//go:noescape
func FuncOffChildrenReordered(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OffChildrenReordered
//go:noescape
func CallOffChildrenReordered(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OffChildrenReordered
//go:noescape
func TryOffChildrenReordered(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_HasOnChildrenReordered
//go:noescape
func HasFuncHasOnChildrenReordered() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_HasOnChildrenReordered
//go:noescape
func FuncHasOnChildrenReordered(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_HasOnChildrenReordered
//go:noescape
func CallHasOnChildrenReordered(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_HasOnChildrenReordered
//go:noescape
func TryHasOnChildrenReordered(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OnCreated
//go:noescape
func HasFuncOnCreated() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OnCreated
//go:noescape
func FuncOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OnCreated
//go:noescape
func CallOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OnCreated
//go:noescape
func TryOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OffCreated
//go:noescape
func HasFuncOffCreated() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OffCreated
//go:noescape
func FuncOffCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OffCreated
//go:noescape
func CallOffCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OffCreated
//go:noescape
func TryOffCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_HasOnCreated
//go:noescape
func HasFuncHasOnCreated() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_HasOnCreated
//go:noescape
func FuncHasOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_HasOnCreated
//go:noescape
func CallHasOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_HasOnCreated
//go:noescape
func TryHasOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OnImportBegan
//go:noescape
func HasFuncOnImportBegan() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OnImportBegan
//go:noescape
func FuncOnImportBegan(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OnImportBegan
//go:noescape
func CallOnImportBegan(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OnImportBegan
//go:noescape
func TryOnImportBegan(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OffImportBegan
//go:noescape
func HasFuncOffImportBegan() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OffImportBegan
//go:noescape
func FuncOffImportBegan(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OffImportBegan
//go:noescape
func CallOffImportBegan(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OffImportBegan
//go:noescape
func TryOffImportBegan(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_HasOnImportBegan
//go:noescape
func HasFuncHasOnImportBegan() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_HasOnImportBegan
//go:noescape
func FuncHasOnImportBegan(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_HasOnImportBegan
//go:noescape
func CallHasOnImportBegan(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_HasOnImportBegan
//go:noescape
func TryHasOnImportBegan(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OnImportEnded
//go:noescape
func HasFuncOnImportEnded() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OnImportEnded
//go:noescape
func FuncOnImportEnded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OnImportEnded
//go:noescape
func CallOnImportEnded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OnImportEnded
//go:noescape
func TryOnImportEnded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OffImportEnded
//go:noescape
func HasFuncOffImportEnded() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OffImportEnded
//go:noescape
func FuncOffImportEnded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OffImportEnded
//go:noescape
func CallOffImportEnded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OffImportEnded
//go:noescape
func TryOffImportEnded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_HasOnImportEnded
//go:noescape
func HasFuncHasOnImportEnded() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_HasOnImportEnded
//go:noescape
func FuncHasOnImportEnded(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_HasOnImportEnded
//go:noescape
func CallHasOnImportEnded(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_HasOnImportEnded
//go:noescape
func TryHasOnImportEnded(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OnMoved
//go:noescape
func HasFuncOnMoved() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OnMoved
//go:noescape
func FuncOnMoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OnMoved
//go:noescape
func CallOnMoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OnMoved
//go:noescape
func TryOnMoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OffMoved
//go:noescape
func HasFuncOffMoved() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OffMoved
//go:noescape
func FuncOffMoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OffMoved
//go:noescape
func CallOffMoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OffMoved
//go:noescape
func TryOffMoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_HasOnMoved
//go:noescape
func HasFuncHasOnMoved() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_HasOnMoved
//go:noescape
func FuncHasOnMoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_HasOnMoved
//go:noescape
func CallHasOnMoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_HasOnMoved
//go:noescape
func TryHasOnMoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OnRemoved
//go:noescape
func HasFuncOnRemoved() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OnRemoved
//go:noescape
func FuncOnRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OnRemoved
//go:noescape
func CallOnRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OnRemoved
//go:noescape
func TryOnRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_OffRemoved
//go:noescape
func HasFuncOffRemoved() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_OffRemoved
//go:noescape
func FuncOffRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_OffRemoved
//go:noescape
func CallOffRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_OffRemoved
//go:noescape
func TryOffRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_HasOnRemoved
//go:noescape
func HasFuncHasOnRemoved() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_HasOnRemoved
//go:noescape
func FuncHasOnRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_HasOnRemoved
//go:noescape
func CallHasOnRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_HasOnRemoved
//go:noescape
func TryHasOnRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_Remove
//go:noescape
func HasFuncRemove() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_Remove
//go:noescape
func FuncRemove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_Remove
//go:noescape
func CallRemove(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_Remove
//go:noescape
func TryRemove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_RemoveTree
//go:noescape
func HasFuncRemoveTree() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_RemoveTree
//go:noescape
func FuncRemoveTree(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_RemoveTree
//go:noescape
func CallRemoveTree(
	retPtr unsafe.Pointer,
	id js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_RemoveTree
//go:noescape
func TryRemoveTree(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_Search
//go:noescape
func HasFuncSearch() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_Search
//go:noescape
func FuncSearch(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_Search
//go:noescape
func CallSearch(
	retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/webext/bookmarks try_Search
//go:noescape
func TrySearch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/bookmarks has_Update
//go:noescape
func HasFuncUpdate() js.Ref

//go:wasmimport plat/js/webext/bookmarks func_Update
//go:noescape
func FuncUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks call_Update
//go:noescape
func CallUpdate(
	retPtr unsafe.Pointer,
	id js.Ref,
	changes unsafe.Pointer)

//go:wasmimport plat/js/webext/bookmarks try_Update
//go:noescape
func TryUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	id js.Ref,
	changes unsafe.Pointer) (ok js.Ref)

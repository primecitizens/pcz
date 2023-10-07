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

//go:wasmimport plat/js/webext/tabgroups constof_Color
//go:noescape
func ConstOfColor(str js.Ref) uint32

//go:wasmimport plat/js/webext/tabgroups store_MoveArgMoveProperties
//go:noescape
func MoveArgMovePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabgroups load_MoveArgMoveProperties
//go:noescape
func MoveArgMovePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabgroups store_QueryArgQueryInfo
//go:noescape
func QueryArgQueryInfoJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabgroups load_QueryArgQueryInfo
//go:noescape
func QueryArgQueryInfoJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabgroups get_TAB_GROUP_ID_NONE
//go:noescape
func GetTAB_GROUP_ID_NONE(retPtr unsafe.Pointer) js.Ref

//go:wasmimport plat/js/webext/tabgroups set_TAB_GROUP_ID_NONE
//go:noescape
func SetTAB_GROUP_ID_NONE(
	val js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabgroups store_TabGroup
//go:noescape
func TabGroupJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabgroups load_TabGroup
//go:noescape
func TabGroupJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabgroups store_UpdateArgUpdateProperties
//go:noescape
func UpdateArgUpdatePropertiesJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/tabgroups load_UpdateArgUpdateProperties
//go:noescape
func UpdateArgUpdatePropertiesJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/tabgroups has_Get
//go:noescape
func HasFuncGet() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_Get
//go:noescape
func FuncGet(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_Get
//go:noescape
func CallGet(
	retPtr unsafe.Pointer,
	groupId float64)

//go:wasmimport plat/js/webext/tabgroups try_Get
//go:noescape
func TryGet(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	groupId float64) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_Move
//go:noescape
func HasFuncMove() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_Move
//go:noescape
func FuncMove(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_Move
//go:noescape
func CallMove(
	retPtr unsafe.Pointer,
	groupId float64,
	moveProperties unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups try_Move
//go:noescape
func TryMove(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	groupId float64,
	moveProperties unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_OnCreated
//go:noescape
func HasFuncOnCreated() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_OnCreated
//go:noescape
func FuncOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_OnCreated
//go:noescape
func CallOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_OnCreated
//go:noescape
func TryOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_OffCreated
//go:noescape
func HasFuncOffCreated() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_OffCreated
//go:noescape
func FuncOffCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_OffCreated
//go:noescape
func CallOffCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_OffCreated
//go:noescape
func TryOffCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_HasOnCreated
//go:noescape
func HasFuncHasOnCreated() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_HasOnCreated
//go:noescape
func FuncHasOnCreated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_HasOnCreated
//go:noescape
func CallHasOnCreated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_HasOnCreated
//go:noescape
func TryHasOnCreated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_OnMoved
//go:noescape
func HasFuncOnMoved() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_OnMoved
//go:noescape
func FuncOnMoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_OnMoved
//go:noescape
func CallOnMoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_OnMoved
//go:noescape
func TryOnMoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_OffMoved
//go:noescape
func HasFuncOffMoved() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_OffMoved
//go:noescape
func FuncOffMoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_OffMoved
//go:noescape
func CallOffMoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_OffMoved
//go:noescape
func TryOffMoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_HasOnMoved
//go:noescape
func HasFuncHasOnMoved() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_HasOnMoved
//go:noescape
func FuncHasOnMoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_HasOnMoved
//go:noescape
func CallHasOnMoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_HasOnMoved
//go:noescape
func TryHasOnMoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_OnRemoved
//go:noescape
func HasFuncOnRemoved() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_OnRemoved
//go:noescape
func FuncOnRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_OnRemoved
//go:noescape
func CallOnRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_OnRemoved
//go:noescape
func TryOnRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_OffRemoved
//go:noescape
func HasFuncOffRemoved() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_OffRemoved
//go:noescape
func FuncOffRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_OffRemoved
//go:noescape
func CallOffRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_OffRemoved
//go:noescape
func TryOffRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_HasOnRemoved
//go:noescape
func HasFuncHasOnRemoved() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_HasOnRemoved
//go:noescape
func FuncHasOnRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_HasOnRemoved
//go:noescape
func CallHasOnRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_HasOnRemoved
//go:noescape
func TryHasOnRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_OnUpdated
//go:noescape
func HasFuncOnUpdated() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_OnUpdated
//go:noescape
func FuncOnUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_OnUpdated
//go:noescape
func CallOnUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_OnUpdated
//go:noescape
func TryOnUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_OffUpdated
//go:noescape
func HasFuncOffUpdated() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_OffUpdated
//go:noescape
func FuncOffUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_OffUpdated
//go:noescape
func CallOffUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_OffUpdated
//go:noescape
func TryOffUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_HasOnUpdated
//go:noescape
func HasFuncHasOnUpdated() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_HasOnUpdated
//go:noescape
func FuncHasOnUpdated(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_HasOnUpdated
//go:noescape
func CallHasOnUpdated(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/tabgroups try_HasOnUpdated
//go:noescape
func TryHasOnUpdated(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_Query
//go:noescape
func HasFuncQuery() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_Query
//go:noescape
func FuncQuery(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_Query
//go:noescape
func CallQuery(
	retPtr unsafe.Pointer,
	queryInfo unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups try_Query
//go:noescape
func TryQuery(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	queryInfo unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/tabgroups has_Update
//go:noescape
func HasFuncUpdate() js.Ref

//go:wasmimport plat/js/webext/tabgroups func_Update
//go:noescape
func FuncUpdate(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups call_Update
//go:noescape
func CallUpdate(
	retPtr unsafe.Pointer,
	groupId float64,
	updateProperties unsafe.Pointer)

//go:wasmimport plat/js/webext/tabgroups try_Update
//go:noescape
func TryUpdate(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	groupId float64,
	updateProperties unsafe.Pointer) (ok js.Ref)

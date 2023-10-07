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

//go:wasmimport plat/js/webext/history store_DeleteRangeArgRange
//go:noescape
func DeleteRangeArgRangeJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/history load_DeleteRangeArgRange
//go:noescape
func DeleteRangeArgRangeJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/history store_HistoryItem
//go:noescape
func HistoryItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/history load_HistoryItem
//go:noescape
func HistoryItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/history store_OnVisitRemovedArgRemoved
//go:noescape
func OnVisitRemovedArgRemovedJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/history load_OnVisitRemovedArgRemoved
//go:noescape
func OnVisitRemovedArgRemovedJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/history store_SearchArgQuery
//go:noescape
func SearchArgQueryJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/history load_SearchArgQuery
//go:noescape
func SearchArgQueryJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/history constof_TransitionType
//go:noescape
func ConstOfTransitionType(str js.Ref) uint32

//go:wasmimport plat/js/webext/history store_UrlDetails
//go:noescape
func UrlDetailsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/history load_UrlDetails
//go:noescape
func UrlDetailsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/history store_VisitItem
//go:noescape
func VisitItemJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/webext/history load_VisitItem
//go:noescape
func VisitItemJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/webext/history has_AddUrl
//go:noescape
func HasFuncAddUrl() js.Ref

//go:wasmimport plat/js/webext/history func_AddUrl
//go:noescape
func FuncAddUrl(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_AddUrl
//go:noescape
func CallAddUrl(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/history try_AddUrl
//go:noescape
func TryAddUrl(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/history has_DeleteAll
//go:noescape
func HasFuncDeleteAll() js.Ref

//go:wasmimport plat/js/webext/history func_DeleteAll
//go:noescape
func FuncDeleteAll(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_DeleteAll
//go:noescape
func CallDeleteAll(
	retPtr unsafe.Pointer)

//go:wasmimport plat/js/webext/history try_DeleteAll
//go:noescape
func TryDeleteAll(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/history has_DeleteRange
//go:noescape
func HasFuncDeleteRange() js.Ref

//go:wasmimport plat/js/webext/history func_DeleteRange
//go:noescape
func FuncDeleteRange(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_DeleteRange
//go:noescape
func CallDeleteRange(
	retPtr unsafe.Pointer,
	rng unsafe.Pointer)

//go:wasmimport plat/js/webext/history try_DeleteRange
//go:noescape
func TryDeleteRange(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rng unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/history has_DeleteUrl
//go:noescape
func HasFuncDeleteUrl() js.Ref

//go:wasmimport plat/js/webext/history func_DeleteUrl
//go:noescape
func FuncDeleteUrl(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_DeleteUrl
//go:noescape
func CallDeleteUrl(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/history try_DeleteUrl
//go:noescape
func TryDeleteUrl(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/history has_GetVisits
//go:noescape
func HasFuncGetVisits() js.Ref

//go:wasmimport plat/js/webext/history func_GetVisits
//go:noescape
func FuncGetVisits(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_GetVisits
//go:noescape
func CallGetVisits(
	retPtr unsafe.Pointer,
	details unsafe.Pointer)

//go:wasmimport plat/js/webext/history try_GetVisits
//go:noescape
func TryGetVisits(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	details unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/webext/history has_OnVisitRemoved
//go:noescape
func HasFuncOnVisitRemoved() js.Ref

//go:wasmimport plat/js/webext/history func_OnVisitRemoved
//go:noescape
func FuncOnVisitRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_OnVisitRemoved
//go:noescape
func CallOnVisitRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/history try_OnVisitRemoved
//go:noescape
func TryOnVisitRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/history has_OffVisitRemoved
//go:noescape
func HasFuncOffVisitRemoved() js.Ref

//go:wasmimport plat/js/webext/history func_OffVisitRemoved
//go:noescape
func FuncOffVisitRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_OffVisitRemoved
//go:noescape
func CallOffVisitRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/history try_OffVisitRemoved
//go:noescape
func TryOffVisitRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/history has_HasOnVisitRemoved
//go:noescape
func HasFuncHasOnVisitRemoved() js.Ref

//go:wasmimport plat/js/webext/history func_HasOnVisitRemoved
//go:noescape
func FuncHasOnVisitRemoved(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_HasOnVisitRemoved
//go:noescape
func CallHasOnVisitRemoved(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/history try_HasOnVisitRemoved
//go:noescape
func TryHasOnVisitRemoved(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/history has_OnVisited
//go:noescape
func HasFuncOnVisited() js.Ref

//go:wasmimport plat/js/webext/history func_OnVisited
//go:noescape
func FuncOnVisited(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_OnVisited
//go:noescape
func CallOnVisited(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/history try_OnVisited
//go:noescape
func TryOnVisited(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/history has_OffVisited
//go:noescape
func HasFuncOffVisited() js.Ref

//go:wasmimport plat/js/webext/history func_OffVisited
//go:noescape
func FuncOffVisited(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_OffVisited
//go:noescape
func CallOffVisited(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/history try_OffVisited
//go:noescape
func TryOffVisited(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/history has_HasOnVisited
//go:noescape
func HasFuncHasOnVisited() js.Ref

//go:wasmimport plat/js/webext/history func_HasOnVisited
//go:noescape
func FuncHasOnVisited(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_HasOnVisited
//go:noescape
func CallHasOnVisited(
	retPtr unsafe.Pointer,
	callback js.Ref)

//go:wasmimport plat/js/webext/history try_HasOnVisited
//go:noescape
func TryHasOnVisited(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	callback js.Ref) (ok js.Ref)

//go:wasmimport plat/js/webext/history has_Search
//go:noescape
func HasFuncSearch() js.Ref

//go:wasmimport plat/js/webext/history func_Search
//go:noescape
func FuncSearch(fn unsafe.Pointer)

//go:wasmimport plat/js/webext/history call_Search
//go:noescape
func CallSearch(
	retPtr unsafe.Pointer,
	query unsafe.Pointer)

//go:wasmimport plat/js/webext/history try_Search
//go:noescape
func TrySearch(
	retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query unsafe.Pointer) (ok js.Ref)

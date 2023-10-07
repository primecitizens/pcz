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

//go:wasmimport plat/js/web get_SVGAnimatedEnumeration_BaseVal
//go:noescape
func GetSVGAnimatedEnumerationBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAnimatedEnumeration_BaseVal
//go:noescape
func SetSVGAnimatedEnumerationBaseVal(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedEnumeration_AnimVal
//go:noescape
func GetSVGAnimatedEnumerationAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedInteger_BaseVal
//go:noescape
func GetSVGAnimatedIntegerBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAnimatedInteger_BaseVal
//go:noescape
func SetSVGAnimatedIntegerBaseVal(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedInteger_AnimVal
//go:noescape
func GetSVGAnimatedIntegerAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLengthList_Length
//go:noescape
func GetSVGLengthListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLengthList_NumberOfItems
//go:noescape
func GetSVGLengthListNumberOfItems(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGLengthList_Clear
//go:noescape
func HasFuncSVGLengthListClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_Clear
//go:noescape
func FuncSVGLengthListClear(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGLengthList_Clear
//go:noescape
func CallSVGLengthListClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGLengthList_Clear
//go:noescape
func TrySVGLengthListClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGLengthList_Initialize
//go:noescape
func HasFuncSVGLengthListInitialize(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_Initialize
//go:noescape
func FuncSVGLengthListInitialize(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGLengthList_Initialize
//go:noescape
func CallSVGLengthListInitialize(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGLengthList_Initialize
//go:noescape
func TrySVGLengthListInitialize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGLengthList_GetItem
//go:noescape
func HasFuncSVGLengthListGetItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_GetItem
//go:noescape
func FuncSVGLengthListGetItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGLengthList_GetItem
//go:noescape
func CallSVGLengthListGetItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_SVGLengthList_GetItem
//go:noescape
func TrySVGLengthListGetItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGLengthList_InsertItemBefore
//go:noescape
func HasFuncSVGLengthListInsertItemBefore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_InsertItemBefore
//go:noescape
func FuncSVGLengthListInsertItemBefore(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGLengthList_InsertItemBefore
//go:noescape
func CallSVGLengthListInsertItemBefore(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32)

//go:wasmimport plat/js/web try_SVGLengthList_InsertItemBefore
//go:noescape
func TrySVGLengthListInsertItemBefore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGLengthList_ReplaceItem
//go:noescape
func HasFuncSVGLengthListReplaceItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_ReplaceItem
//go:noescape
func FuncSVGLengthListReplaceItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGLengthList_ReplaceItem
//go:noescape
func CallSVGLengthListReplaceItem(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32)

//go:wasmimport plat/js/web try_SVGLengthList_ReplaceItem
//go:noescape
func TrySVGLengthListReplaceItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGLengthList_RemoveItem
//go:noescape
func HasFuncSVGLengthListRemoveItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_RemoveItem
//go:noescape
func FuncSVGLengthListRemoveItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGLengthList_RemoveItem
//go:noescape
func CallSVGLengthListRemoveItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_SVGLengthList_RemoveItem
//go:noescape
func TrySVGLengthListRemoveItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGLengthList_AppendItem
//go:noescape
func HasFuncSVGLengthListAppendItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_AppendItem
//go:noescape
func FuncSVGLengthListAppendItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGLengthList_AppendItem
//go:noescape
func CallSVGLengthListAppendItem(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGLengthList_AppendItem
//go:noescape
func TrySVGLengthListAppendItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGLengthList_Set
//go:noescape
func HasFuncSVGLengthListSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_Set
//go:noescape
func FuncSVGLengthListSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGLengthList_Set
//go:noescape
func CallSVGLengthListSet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGLengthList_Set
//go:noescape
func TrySVGLengthListSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedLengthList_BaseVal
//go:noescape
func GetSVGAnimatedLengthListBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedLengthList_AnimVal
//go:noescape
func GetSVGAnimatedLengthListAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedNumber_BaseVal
//go:noescape
func GetSVGAnimatedNumberBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAnimatedNumber_BaseVal
//go:noescape
func SetSVGAnimatedNumberBaseVal(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedNumber_AnimVal
//go:noescape
func GetSVGAnimatedNumberAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGNumberList_Length
//go:noescape
func GetSVGNumberListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGNumberList_NumberOfItems
//go:noescape
func GetSVGNumberListNumberOfItems(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGNumberList_Clear
//go:noescape
func HasFuncSVGNumberListClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_Clear
//go:noescape
func FuncSVGNumberListClear(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGNumberList_Clear
//go:noescape
func CallSVGNumberListClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGNumberList_Clear
//go:noescape
func TrySVGNumberListClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGNumberList_Initialize
//go:noescape
func HasFuncSVGNumberListInitialize(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_Initialize
//go:noescape
func FuncSVGNumberListInitialize(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGNumberList_Initialize
//go:noescape
func CallSVGNumberListInitialize(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGNumberList_Initialize
//go:noescape
func TrySVGNumberListInitialize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGNumberList_GetItem
//go:noescape
func HasFuncSVGNumberListGetItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_GetItem
//go:noescape
func FuncSVGNumberListGetItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGNumberList_GetItem
//go:noescape
func CallSVGNumberListGetItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_SVGNumberList_GetItem
//go:noescape
func TrySVGNumberListGetItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGNumberList_InsertItemBefore
//go:noescape
func HasFuncSVGNumberListInsertItemBefore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_InsertItemBefore
//go:noescape
func FuncSVGNumberListInsertItemBefore(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGNumberList_InsertItemBefore
//go:noescape
func CallSVGNumberListInsertItemBefore(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32)

//go:wasmimport plat/js/web try_SVGNumberList_InsertItemBefore
//go:noescape
func TrySVGNumberListInsertItemBefore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGNumberList_ReplaceItem
//go:noescape
func HasFuncSVGNumberListReplaceItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_ReplaceItem
//go:noescape
func FuncSVGNumberListReplaceItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGNumberList_ReplaceItem
//go:noescape
func CallSVGNumberListReplaceItem(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32)

//go:wasmimport plat/js/web try_SVGNumberList_ReplaceItem
//go:noescape
func TrySVGNumberListReplaceItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGNumberList_RemoveItem
//go:noescape
func HasFuncSVGNumberListRemoveItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_RemoveItem
//go:noescape
func FuncSVGNumberListRemoveItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGNumberList_RemoveItem
//go:noescape
func CallSVGNumberListRemoveItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_SVGNumberList_RemoveItem
//go:noescape
func TrySVGNumberListRemoveItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGNumberList_AppendItem
//go:noescape
func HasFuncSVGNumberListAppendItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_AppendItem
//go:noescape
func FuncSVGNumberListAppendItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGNumberList_AppendItem
//go:noescape
func CallSVGNumberListAppendItem(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGNumberList_AppendItem
//go:noescape
func TrySVGNumberListAppendItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGNumberList_Set
//go:noescape
func HasFuncSVGNumberListSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_Set
//go:noescape
func FuncSVGNumberListSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGNumberList_Set
//go:noescape
func CallSVGNumberListSet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGNumberList_Set
//go:noescape
func TrySVGNumberListSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedNumberList_BaseVal
//go:noescape
func GetSVGAnimatedNumberListBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedNumberList_AnimVal
//go:noescape
func GetSVGAnimatedNumberListAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTransformList_Length
//go:noescape
func GetSVGTransformListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTransformList_NumberOfItems
//go:noescape
func GetSVGTransformListNumberOfItems(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransformList_Clear
//go:noescape
func HasFuncSVGTransformListClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_Clear
//go:noescape
func FuncSVGTransformListClear(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTransformList_Clear
//go:noescape
func CallSVGTransformListClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGTransformList_Clear
//go:noescape
func TrySVGTransformListClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransformList_Initialize
//go:noescape
func HasFuncSVGTransformListInitialize(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_Initialize
//go:noescape
func FuncSVGTransformListInitialize(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTransformList_Initialize
//go:noescape
func CallSVGTransformListInitialize(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGTransformList_Initialize
//go:noescape
func TrySVGTransformListInitialize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransformList_GetItem
//go:noescape
func HasFuncSVGTransformListGetItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_GetItem
//go:noescape
func FuncSVGTransformListGetItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTransformList_GetItem
//go:noescape
func CallSVGTransformListGetItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_SVGTransformList_GetItem
//go:noescape
func TrySVGTransformListGetItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransformList_InsertItemBefore
//go:noescape
func HasFuncSVGTransformListInsertItemBefore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_InsertItemBefore
//go:noescape
func FuncSVGTransformListInsertItemBefore(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTransformList_InsertItemBefore
//go:noescape
func CallSVGTransformListInsertItemBefore(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32)

//go:wasmimport plat/js/web try_SVGTransformList_InsertItemBefore
//go:noescape
func TrySVGTransformListInsertItemBefore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransformList_ReplaceItem
//go:noescape
func HasFuncSVGTransformListReplaceItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_ReplaceItem
//go:noescape
func FuncSVGTransformListReplaceItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTransformList_ReplaceItem
//go:noescape
func CallSVGTransformListReplaceItem(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32)

//go:wasmimport plat/js/web try_SVGTransformList_ReplaceItem
//go:noescape
func TrySVGTransformListReplaceItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransformList_RemoveItem
//go:noescape
func HasFuncSVGTransformListRemoveItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_RemoveItem
//go:noescape
func FuncSVGTransformListRemoveItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTransformList_RemoveItem
//go:noescape
func CallSVGTransformListRemoveItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_SVGTransformList_RemoveItem
//go:noescape
func TrySVGTransformListRemoveItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransformList_AppendItem
//go:noescape
func HasFuncSVGTransformListAppendItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_AppendItem
//go:noescape
func FuncSVGTransformListAppendItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTransformList_AppendItem
//go:noescape
func CallSVGTransformListAppendItem(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGTransformList_AppendItem
//go:noescape
func TrySVGTransformListAppendItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransformList_Set
//go:noescape
func HasFuncSVGTransformListSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_Set
//go:noescape
func FuncSVGTransformListSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTransformList_Set
//go:noescape
func CallSVGTransformListSet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGTransformList_Set
//go:noescape
func TrySVGTransformListSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransformList_CreateSVGTransformFromMatrix
//go:noescape
func HasFuncSVGTransformListCreateSVGTransformFromMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_CreateSVGTransformFromMatrix
//go:noescape
func FuncSVGTransformListCreateSVGTransformFromMatrix(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTransformList_CreateSVGTransformFromMatrix
//go:noescape
func CallSVGTransformListCreateSVGTransformFromMatrix(
	this js.Ref, retPtr unsafe.Pointer,
	matrix unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGTransformList_CreateSVGTransformFromMatrix
//go:noescape
func TrySVGTransformListCreateSVGTransformFromMatrix(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	matrix unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransformList_CreateSVGTransformFromMatrix1
//go:noescape
func HasFuncSVGTransformListCreateSVGTransformFromMatrix1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_CreateSVGTransformFromMatrix1
//go:noescape
func FuncSVGTransformListCreateSVGTransformFromMatrix1(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTransformList_CreateSVGTransformFromMatrix1
//go:noescape
func CallSVGTransformListCreateSVGTransformFromMatrix1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGTransformList_CreateSVGTransformFromMatrix1
//go:noescape
func TrySVGTransformListCreateSVGTransformFromMatrix1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransformList_Consolidate
//go:noescape
func HasFuncSVGTransformListConsolidate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_Consolidate
//go:noescape
func FuncSVGTransformListConsolidate(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGTransformList_Consolidate
//go:noescape
func CallSVGTransformListConsolidate(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGTransformList_Consolidate
//go:noescape
func TrySVGTransformListConsolidate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedTransformList_BaseVal
//go:noescape
func GetSVGAnimatedTransformListBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedTransformList_AnimVal
//go:noescape
func GetSVGAnimatedTransformListAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGStringList_Length
//go:noescape
func GetSVGStringListLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGStringList_NumberOfItems
//go:noescape
func GetSVGStringListNumberOfItems(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGStringList_Clear
//go:noescape
func HasFuncSVGStringListClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_Clear
//go:noescape
func FuncSVGStringListClear(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGStringList_Clear
//go:noescape
func CallSVGStringListClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGStringList_Clear
//go:noescape
func TrySVGStringListClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGStringList_Initialize
//go:noescape
func HasFuncSVGStringListInitialize(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_Initialize
//go:noescape
func FuncSVGStringListInitialize(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGStringList_Initialize
//go:noescape
func CallSVGStringListInitialize(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGStringList_Initialize
//go:noescape
func TrySVGStringListInitialize(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGStringList_GetItem
//go:noescape
func HasFuncSVGStringListGetItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_GetItem
//go:noescape
func FuncSVGStringListGetItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGStringList_GetItem
//go:noescape
func CallSVGStringListGetItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_SVGStringList_GetItem
//go:noescape
func TrySVGStringListGetItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGStringList_InsertItemBefore
//go:noescape
func HasFuncSVGStringListInsertItemBefore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_InsertItemBefore
//go:noescape
func FuncSVGStringListInsertItemBefore(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGStringList_InsertItemBefore
//go:noescape
func CallSVGStringListInsertItemBefore(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32)

//go:wasmimport plat/js/web try_SVGStringList_InsertItemBefore
//go:noescape
func TrySVGStringListInsertItemBefore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGStringList_ReplaceItem
//go:noescape
func HasFuncSVGStringListReplaceItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_ReplaceItem
//go:noescape
func FuncSVGStringListReplaceItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGStringList_ReplaceItem
//go:noescape
func CallSVGStringListReplaceItem(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32)

//go:wasmimport plat/js/web try_SVGStringList_ReplaceItem
//go:noescape
func TrySVGStringListReplaceItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGStringList_RemoveItem
//go:noescape
func HasFuncSVGStringListRemoveItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_RemoveItem
//go:noescape
func FuncSVGStringListRemoveItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGStringList_RemoveItem
//go:noescape
func CallSVGStringListRemoveItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_SVGStringList_RemoveItem
//go:noescape
func TrySVGStringListRemoveItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGStringList_AppendItem
//go:noescape
func HasFuncSVGStringListAppendItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_AppendItem
//go:noescape
func FuncSVGStringListAppendItem(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGStringList_AppendItem
//go:noescape
func CallSVGStringListAppendItem(
	this js.Ref, retPtr unsafe.Pointer,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGStringList_AppendItem
//go:noescape
func TrySVGStringListAppendItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGStringList_Set
//go:noescape
func HasFuncSVGStringListSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_Set
//go:noescape
func FuncSVGStringListSet(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGStringList_Set
//go:noescape
func CallSVGStringListSet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	newItem js.Ref)

//go:wasmimport plat/js/web try_SVGStringList_Set
//go:noescape
func TrySVGStringListSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	newItem js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimationElement_TargetElement
//go:noescape
func GetSVGAnimationElementTargetElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimationElement_RequiredExtensions
//go:noescape
func GetSVGAnimationElementRequiredExtensions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimationElement_SystemLanguage
//go:noescape
func GetSVGAnimationElementSystemLanguage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGAnimationElement_GetStartTime
//go:noescape
func HasFuncSVGAnimationElementGetStartTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGAnimationElement_GetStartTime
//go:noescape
func FuncSVGAnimationElementGetStartTime(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGAnimationElement_GetStartTime
//go:noescape
func CallSVGAnimationElementGetStartTime(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGAnimationElement_GetStartTime
//go:noescape
func TrySVGAnimationElementGetStartTime(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGAnimationElement_GetCurrentTime
//go:noescape
func HasFuncSVGAnimationElementGetCurrentTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGAnimationElement_GetCurrentTime
//go:noescape
func FuncSVGAnimationElementGetCurrentTime(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGAnimationElement_GetCurrentTime
//go:noescape
func CallSVGAnimationElementGetCurrentTime(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGAnimationElement_GetCurrentTime
//go:noescape
func TrySVGAnimationElementGetCurrentTime(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGAnimationElement_GetSimpleDuration
//go:noescape
func HasFuncSVGAnimationElementGetSimpleDuration(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGAnimationElement_GetSimpleDuration
//go:noescape
func FuncSVGAnimationElementGetSimpleDuration(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGAnimationElement_GetSimpleDuration
//go:noescape
func CallSVGAnimationElementGetSimpleDuration(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGAnimationElement_GetSimpleDuration
//go:noescape
func TrySVGAnimationElementGetSimpleDuration(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGAnimationElement_BeginElement
//go:noescape
func HasFuncSVGAnimationElementBeginElement(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGAnimationElement_BeginElement
//go:noescape
func FuncSVGAnimationElementBeginElement(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGAnimationElement_BeginElement
//go:noescape
func CallSVGAnimationElementBeginElement(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGAnimationElement_BeginElement
//go:noescape
func TrySVGAnimationElementBeginElement(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGAnimationElement_BeginElementAt
//go:noescape
func HasFuncSVGAnimationElementBeginElementAt(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGAnimationElement_BeginElementAt
//go:noescape
func FuncSVGAnimationElementBeginElementAt(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGAnimationElement_BeginElementAt
//go:noescape
func CallSVGAnimationElementBeginElementAt(
	this js.Ref, retPtr unsafe.Pointer,
	offset float32)

//go:wasmimport plat/js/web try_SVGAnimationElement_BeginElementAt
//go:noescape
func TrySVGAnimationElementBeginElementAt(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	offset float32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGAnimationElement_EndElement
//go:noescape
func HasFuncSVGAnimationElementEndElement(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGAnimationElement_EndElement
//go:noescape
func FuncSVGAnimationElementEndElement(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGAnimationElement_EndElement
//go:noescape
func CallSVGAnimationElementEndElement(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGAnimationElement_EndElement
//go:noescape
func TrySVGAnimationElementEndElement(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGAnimationElement_EndElementAt
//go:noescape
func HasFuncSVGAnimationElementEndElementAt(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGAnimationElement_EndElementAt
//go:noescape
func FuncSVGAnimationElementEndElementAt(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGAnimationElement_EndElementAt
//go:noescape
func CallSVGAnimationElementEndElementAt(
	this js.Ref, retPtr unsafe.Pointer,
	offset float32)

//go:wasmimport plat/js/web try_SVGAnimationElement_EndElementAt
//go:noescape
func TrySVGAnimationElementEndElementAt(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	offset float32) (ok js.Ref)

//go:wasmimport plat/js/web store_SVGBoundingBoxOptions
//go:noescape
func SVGBoundingBoxOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SVGBoundingBoxOptions
//go:noescape
func SVGBoundingBoxOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGCircleElement_Cx
//go:noescape
func GetSVGCircleElementCx(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGCircleElement_Cy
//go:noescape
func GetSVGCircleElementCy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGCircleElement_R
//go:noescape
func GetSVGCircleElementR(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGClipPathElement_ClipPathUnits
//go:noescape
func GetSVGClipPathElementClipPathUnits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGClipPathElement_Transform
//go:noescape
func GetSVGClipPathElementTransform(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Type
//go:noescape
func GetSVGComponentTransferFunctionElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_TableValues
//go:noescape
func GetSVGComponentTransferFunctionElementTableValues(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Slope
//go:noescape
func GetSVGComponentTransferFunctionElementSlope(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Intercept
//go:noescape
func GetSVGComponentTransferFunctionElementIntercept(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Amplitude
//go:noescape
func GetSVGComponentTransferFunctionElementAmplitude(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Exponent
//go:noescape
func GetSVGComponentTransferFunctionElementExponent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Offset
//go:noescape
func GetSVGComponentTransferFunctionElementOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGEllipseElement_Cx
//go:noescape
func GetSVGEllipseElementCx(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGEllipseElement_Cy
//go:noescape
func GetSVGEllipseElementCy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGEllipseElement_Rx
//go:noescape
func GetSVGEllipseElementRx(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGEllipseElement_Ry
//go:noescape
func GetSVGEllipseElementRy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEBlendElement_In1
//go:noescape
func GetSVGFEBlendElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEBlendElement_In2
//go:noescape
func GetSVGFEBlendElementIn2(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEBlendElement_Mode
//go:noescape
func GetSVGFEBlendElementMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEBlendElement_X
//go:noescape
func GetSVGFEBlendElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEBlendElement_Y
//go:noescape
func GetSVGFEBlendElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEBlendElement_Width
//go:noescape
func GetSVGFEBlendElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEBlendElement_Height
//go:noescape
func GetSVGFEBlendElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEBlendElement_Result
//go:noescape
func GetSVGFEBlendElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_In1
//go:noescape
func GetSVGFEColorMatrixElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Type
//go:noescape
func GetSVGFEColorMatrixElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Values
//go:noescape
func GetSVGFEColorMatrixElementValues(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_X
//go:noescape
func GetSVGFEColorMatrixElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Y
//go:noescape
func GetSVGFEColorMatrixElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Width
//go:noescape
func GetSVGFEColorMatrixElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Height
//go:noescape
func GetSVGFEColorMatrixElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Result
//go:noescape
func GetSVGFEColorMatrixElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_In1
//go:noescape
func GetSVGFEComponentTransferElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_X
//go:noescape
func GetSVGFEComponentTransferElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_Y
//go:noescape
func GetSVGFEComponentTransferElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_Width
//go:noescape
func GetSVGFEComponentTransferElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_Height
//go:noescape
func GetSVGFEComponentTransferElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_Result
//go:noescape
func GetSVGFEComponentTransferElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_In1
//go:noescape
func GetSVGFECompositeElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_In2
//go:noescape
func GetSVGFECompositeElementIn2(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_Operator
//go:noescape
func GetSVGFECompositeElementOperator(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_K1
//go:noescape
func GetSVGFECompositeElementK1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_K2
//go:noescape
func GetSVGFECompositeElementK2(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_K3
//go:noescape
func GetSVGFECompositeElementK3(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_K4
//go:noescape
func GetSVGFECompositeElementK4(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_X
//go:noescape
func GetSVGFECompositeElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_Y
//go:noescape
func GetSVGFECompositeElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_Width
//go:noescape
func GetSVGFECompositeElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_Height
//go:noescape
func GetSVGFECompositeElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFECompositeElement_Result
//go:noescape
func GetSVGFECompositeElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_In1
//go:noescape
func GetSVGFEConvolveMatrixElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_OrderX
//go:noescape
func GetSVGFEConvolveMatrixElementOrderX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_OrderY
//go:noescape
func GetSVGFEConvolveMatrixElementOrderY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_KernelMatrix
//go:noescape
func GetSVGFEConvolveMatrixElementKernelMatrix(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Divisor
//go:noescape
func GetSVGFEConvolveMatrixElementDivisor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Bias
//go:noescape
func GetSVGFEConvolveMatrixElementBias(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_TargetX
//go:noescape
func GetSVGFEConvolveMatrixElementTargetX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_TargetY
//go:noescape
func GetSVGFEConvolveMatrixElementTargetY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_EdgeMode
//go:noescape
func GetSVGFEConvolveMatrixElementEdgeMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_KernelUnitLengthX
//go:noescape
func GetSVGFEConvolveMatrixElementKernelUnitLengthX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_KernelUnitLengthY
//go:noescape
func GetSVGFEConvolveMatrixElementKernelUnitLengthY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_PreserveAlpha
//go:noescape
func GetSVGFEConvolveMatrixElementPreserveAlpha(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_X
//go:noescape
func GetSVGFEConvolveMatrixElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Y
//go:noescape
func GetSVGFEConvolveMatrixElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Width
//go:noescape
func GetSVGFEConvolveMatrixElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Height
//go:noescape
func GetSVGFEConvolveMatrixElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Result
//go:noescape
func GetSVGFEConvolveMatrixElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_In1
//go:noescape
func GetSVGFEDiffuseLightingElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_SurfaceScale
//go:noescape
func GetSVGFEDiffuseLightingElementSurfaceScale(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_DiffuseConstant
//go:noescape
func GetSVGFEDiffuseLightingElementDiffuseConstant(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_KernelUnitLengthX
//go:noescape
func GetSVGFEDiffuseLightingElementKernelUnitLengthX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_KernelUnitLengthY
//go:noescape
func GetSVGFEDiffuseLightingElementKernelUnitLengthY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_X
//go:noescape
func GetSVGFEDiffuseLightingElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_Y
//go:noescape
func GetSVGFEDiffuseLightingElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_Width
//go:noescape
func GetSVGFEDiffuseLightingElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_Height
//go:noescape
func GetSVGFEDiffuseLightingElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_Result
//go:noescape
func GetSVGFEDiffuseLightingElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_In1
//go:noescape
func GetSVGFEDisplacementMapElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_In2
//go:noescape
func GetSVGFEDisplacementMapElementIn2(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_Scale
//go:noescape
func GetSVGFEDisplacementMapElementScale(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_XChannelSelector
//go:noescape
func GetSVGFEDisplacementMapElementXChannelSelector(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_YChannelSelector
//go:noescape
func GetSVGFEDisplacementMapElementYChannelSelector(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_X
//go:noescape
func GetSVGFEDisplacementMapElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_Y
//go:noescape
func GetSVGFEDisplacementMapElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_Width
//go:noescape
func GetSVGFEDisplacementMapElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_Height
//go:noescape
func GetSVGFEDisplacementMapElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_Result
//go:noescape
func GetSVGFEDisplacementMapElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDistantLightElement_Azimuth
//go:noescape
func GetSVGFEDistantLightElementAzimuth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDistantLightElement_Elevation
//go:noescape
func GetSVGFEDistantLightElementElevation(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_In1
//go:noescape
func GetSVGFEDropShadowElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Dx
//go:noescape
func GetSVGFEDropShadowElementDx(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Dy
//go:noescape
func GetSVGFEDropShadowElementDy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_StdDeviationX
//go:noescape
func GetSVGFEDropShadowElementStdDeviationX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_StdDeviationY
//go:noescape
func GetSVGFEDropShadowElementStdDeviationY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_X
//go:noescape
func GetSVGFEDropShadowElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Y
//go:noescape
func GetSVGFEDropShadowElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Width
//go:noescape
func GetSVGFEDropShadowElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Height
//go:noescape
func GetSVGFEDropShadowElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Result
//go:noescape
func GetSVGFEDropShadowElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGFEDropShadowElement_SetStdDeviation
//go:noescape
func HasFuncSVGFEDropShadowElementSetStdDeviation(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGFEDropShadowElement_SetStdDeviation
//go:noescape
func FuncSVGFEDropShadowElementSetStdDeviation(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGFEDropShadowElement_SetStdDeviation
//go:noescape
func CallSVGFEDropShadowElementSetStdDeviation(
	this js.Ref, retPtr unsafe.Pointer,
	stdDeviationX float32,
	stdDeviationY float32)

//go:wasmimport plat/js/web try_SVGFEDropShadowElement_SetStdDeviation
//go:noescape
func TrySVGFEDropShadowElementSetStdDeviation(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	stdDeviationX float32,
	stdDeviationY float32) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEFloodElement_X
//go:noescape
func GetSVGFEFloodElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEFloodElement_Y
//go:noescape
func GetSVGFEFloodElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEFloodElement_Width
//go:noescape
func GetSVGFEFloodElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEFloodElement_Height
//go:noescape
func GetSVGFEFloodElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEFloodElement_Result
//go:noescape
func GetSVGFEFloodElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_In1
//go:noescape
func GetSVGFEGaussianBlurElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_StdDeviationX
//go:noescape
func GetSVGFEGaussianBlurElementStdDeviationX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_StdDeviationY
//go:noescape
func GetSVGFEGaussianBlurElementStdDeviationY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_EdgeMode
//go:noescape
func GetSVGFEGaussianBlurElementEdgeMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_X
//go:noescape
func GetSVGFEGaussianBlurElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_Y
//go:noescape
func GetSVGFEGaussianBlurElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_Width
//go:noescape
func GetSVGFEGaussianBlurElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_Height
//go:noescape
func GetSVGFEGaussianBlurElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_Result
//go:noescape
func GetSVGFEGaussianBlurElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGFEGaussianBlurElement_SetStdDeviation
//go:noescape
func HasFuncSVGFEGaussianBlurElementSetStdDeviation(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGFEGaussianBlurElement_SetStdDeviation
//go:noescape
func FuncSVGFEGaussianBlurElementSetStdDeviation(this js.Ref, fn unsafe.Pointer)

//go:wasmimport plat/js/web call_SVGFEGaussianBlurElement_SetStdDeviation
//go:noescape
func CallSVGFEGaussianBlurElementSetStdDeviation(
	this js.Ref, retPtr unsafe.Pointer,
	stdDeviationX float32,
	stdDeviationY float32)

//go:wasmimport plat/js/web try_SVGFEGaussianBlurElement_SetStdDeviation
//go:noescape
func TrySVGFEGaussianBlurElementSetStdDeviation(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	stdDeviationX float32,
	stdDeviationY float32) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEImageElement_PreserveAspectRatio
//go:noescape
func GetSVGFEImageElementPreserveAspectRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEImageElement_CrossOrigin
//go:noescape
func GetSVGFEImageElementCrossOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEImageElement_X
//go:noescape
func GetSVGFEImageElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEImageElement_Y
//go:noescape
func GetSVGFEImageElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEImageElement_Width
//go:noescape
func GetSVGFEImageElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEImageElement_Height
//go:noescape
func GetSVGFEImageElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEImageElement_Result
//go:noescape
func GetSVGFEImageElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEImageElement_Href
//go:noescape
func GetSVGFEImageElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMergeElement_X
//go:noescape
func GetSVGFEMergeElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMergeElement_Y
//go:noescape
func GetSVGFEMergeElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMergeElement_Width
//go:noescape
func GetSVGFEMergeElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMergeElement_Height
//go:noescape
func GetSVGFEMergeElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMergeElement_Result
//go:noescape
func GetSVGFEMergeElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMergeNodeElement_In1
//go:noescape
func GetSVGFEMergeNodeElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_In1
//go:noescape
func GetSVGFEMorphologyElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_Operator
//go:noescape
func GetSVGFEMorphologyElementOperator(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_RadiusX
//go:noescape
func GetSVGFEMorphologyElementRadiusX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_RadiusY
//go:noescape
func GetSVGFEMorphologyElementRadiusY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_X
//go:noescape
func GetSVGFEMorphologyElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_Y
//go:noescape
func GetSVGFEMorphologyElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_Width
//go:noescape
func GetSVGFEMorphologyElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_Height
//go:noescape
func GetSVGFEMorphologyElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_Result
//go:noescape
func GetSVGFEMorphologyElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEOffsetElement_In1
//go:noescape
func GetSVGFEOffsetElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Dx
//go:noescape
func GetSVGFEOffsetElementDx(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Dy
//go:noescape
func GetSVGFEOffsetElementDy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEOffsetElement_X
//go:noescape
func GetSVGFEOffsetElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Y
//go:noescape
func GetSVGFEOffsetElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Width
//go:noescape
func GetSVGFEOffsetElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Height
//go:noescape
func GetSVGFEOffsetElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Result
//go:noescape
func GetSVGFEOffsetElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEPointLightElement_X
//go:noescape
func GetSVGFEPointLightElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEPointLightElement_Y
//go:noescape
func GetSVGFEPointLightElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFEPointLightElement_Z
//go:noescape
func GetSVGFEPointLightElementZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_In1
//go:noescape
func GetSVGFESpecularLightingElementIn1(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_SurfaceScale
//go:noescape
func GetSVGFESpecularLightingElementSurfaceScale(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_SpecularConstant
//go:noescape
func GetSVGFESpecularLightingElementSpecularConstant(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_SpecularExponent
//go:noescape
func GetSVGFESpecularLightingElementSpecularExponent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_KernelUnitLengthX
//go:noescape
func GetSVGFESpecularLightingElementKernelUnitLengthX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_KernelUnitLengthY
//go:noescape
func GetSVGFESpecularLightingElementKernelUnitLengthY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_X
//go:noescape
func GetSVGFESpecularLightingElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_Y
//go:noescape
func GetSVGFESpecularLightingElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_Width
//go:noescape
func GetSVGFESpecularLightingElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_Height
//go:noescape
func GetSVGFESpecularLightingElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_Result
//go:noescape
func GetSVGFESpecularLightingElementResult(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpotLightElement_X
//go:noescape
func GetSVGFESpotLightElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpotLightElement_Y
//go:noescape
func GetSVGFESpotLightElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpotLightElement_Z
//go:noescape
func GetSVGFESpotLightElementZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpotLightElement_PointsAtX
//go:noescape
func GetSVGFESpotLightElementPointsAtX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpotLightElement_PointsAtY
//go:noescape
func GetSVGFESpotLightElementPointsAtY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpotLightElement_PointsAtZ
//go:noescape
func GetSVGFESpotLightElementPointsAtZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpotLightElement_SpecularExponent
//go:noescape
func GetSVGFESpotLightElementSpecularExponent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGFESpotLightElement_LimitingConeAngle
//go:noescape
func GetSVGFESpotLightElementLimitingConeAngle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

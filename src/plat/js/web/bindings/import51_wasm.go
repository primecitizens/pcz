// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package bindings

import (
	"unsafe"

	"github.com/primecitizens/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web get_SVGAnimatedEnumeration_BaseVal
//go:noescape

func GetSVGAnimatedEnumerationBaseVal(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_SVGAnimatedEnumeration_BaseVal
//go:noescape

func SetSVGAnimatedEnumerationBaseVal(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedEnumeration_AnimVal
//go:noescape

func GetSVGAnimatedEnumerationAnimVal(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SVGAnimatedInteger_BaseVal
//go:noescape

func GetSVGAnimatedIntegerBaseVal(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web set_SVGAnimatedInteger_BaseVal
//go:noescape

func SetSVGAnimatedIntegerBaseVal(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedInteger_AnimVal
//go:noescape

func GetSVGAnimatedIntegerAnimVal(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_SVGLengthList_Length
//go:noescape

func GetSVGLengthListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SVGLengthList_NumberOfItems
//go:noescape

func GetSVGLengthListNumberOfItems(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_SVGLengthList_Clear
//go:noescape

func CallSVGLengthListClear(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_Clear
//go:noescape

func SVGLengthListClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGLengthList_Initialize
//go:noescape

func CallSVGLengthListInitialize(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_Initialize
//go:noescape

func SVGLengthListInitializeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGLengthList_GetItem
//go:noescape

func CallSVGLengthListGetItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_GetItem
//go:noescape

func SVGLengthListGetItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGLengthList_InsertItemBefore
//go:noescape

func CallSVGLengthListInsertItemBefore(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_InsertItemBefore
//go:noescape

func SVGLengthListInsertItemBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGLengthList_ReplaceItem
//go:noescape

func CallSVGLengthListReplaceItem(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_ReplaceItem
//go:noescape

func SVGLengthListReplaceItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGLengthList_RemoveItem
//go:noescape

func CallSVGLengthListRemoveItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_RemoveItem
//go:noescape

func SVGLengthListRemoveItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGLengthList_AppendItem
//go:noescape

func CallSVGLengthListAppendItem(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_AppendItem
//go:noescape

func SVGLengthListAppendItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGLengthList_Set
//go:noescape

func CallSVGLengthListSet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGLengthList_Set
//go:noescape

func SVGLengthListSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedLengthList_BaseVal
//go:noescape

func GetSVGAnimatedLengthListBaseVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedLengthList_AnimVal
//go:noescape

func GetSVGAnimatedLengthListAnimVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedNumber_BaseVal
//go:noescape

func GetSVGAnimatedNumberBaseVal(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_SVGAnimatedNumber_BaseVal
//go:noescape

func SetSVGAnimatedNumberBaseVal(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedNumber_AnimVal
//go:noescape

func GetSVGAnimatedNumberAnimVal(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web get_SVGNumberList_Length
//go:noescape

func GetSVGNumberListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SVGNumberList_NumberOfItems
//go:noescape

func GetSVGNumberListNumberOfItems(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_SVGNumberList_Clear
//go:noescape

func CallSVGNumberListClear(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_Clear
//go:noescape

func SVGNumberListClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGNumberList_Initialize
//go:noescape

func CallSVGNumberListInitialize(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_Initialize
//go:noescape

func SVGNumberListInitializeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGNumberList_GetItem
//go:noescape

func CallSVGNumberListGetItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_GetItem
//go:noescape

func SVGNumberListGetItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGNumberList_InsertItemBefore
//go:noescape

func CallSVGNumberListInsertItemBefore(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_InsertItemBefore
//go:noescape

func SVGNumberListInsertItemBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGNumberList_ReplaceItem
//go:noescape

func CallSVGNumberListReplaceItem(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_ReplaceItem
//go:noescape

func SVGNumberListReplaceItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGNumberList_RemoveItem
//go:noescape

func CallSVGNumberListRemoveItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_RemoveItem
//go:noescape

func SVGNumberListRemoveItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGNumberList_AppendItem
//go:noescape

func CallSVGNumberListAppendItem(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_AppendItem
//go:noescape

func SVGNumberListAppendItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGNumberList_Set
//go:noescape

func CallSVGNumberListSet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGNumberList_Set
//go:noescape

func SVGNumberListSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedNumberList_BaseVal
//go:noescape

func GetSVGAnimatedNumberListBaseVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedNumberList_AnimVal
//go:noescape

func GetSVGAnimatedNumberListAnimVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGTransformList_Length
//go:noescape

func GetSVGTransformListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SVGTransformList_NumberOfItems
//go:noescape

func GetSVGTransformListNumberOfItems(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_SVGTransformList_Clear
//go:noescape

func CallSVGTransformListClear(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_Clear
//go:noescape

func SVGTransformListClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransformList_Initialize
//go:noescape

func CallSVGTransformListInitialize(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_Initialize
//go:noescape

func SVGTransformListInitializeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransformList_GetItem
//go:noescape

func CallSVGTransformListGetItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_GetItem
//go:noescape

func SVGTransformListGetItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransformList_InsertItemBefore
//go:noescape

func CallSVGTransformListInsertItemBefore(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_InsertItemBefore
//go:noescape

func SVGTransformListInsertItemBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransformList_ReplaceItem
//go:noescape

func CallSVGTransformListReplaceItem(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_ReplaceItem
//go:noescape

func SVGTransformListReplaceItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransformList_RemoveItem
//go:noescape

func CallSVGTransformListRemoveItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_RemoveItem
//go:noescape

func SVGTransformListRemoveItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransformList_AppendItem
//go:noescape

func CallSVGTransformListAppendItem(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_AppendItem
//go:noescape

func SVGTransformListAppendItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransformList_Set
//go:noescape

func CallSVGTransformListSet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_Set
//go:noescape

func SVGTransformListSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransformList_CreateSVGTransformFromMatrix
//go:noescape

func CallSVGTransformListCreateSVGTransformFromMatrix(
	this js.Ref,
	ptr unsafe.Pointer,

	matrix unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_CreateSVGTransformFromMatrix
//go:noescape

func SVGTransformListCreateSVGTransformFromMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransformList_CreateSVGTransformFromMatrix1
//go:noescape

func CallSVGTransformListCreateSVGTransformFromMatrix1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_CreateSVGTransformFromMatrix1
//go:noescape

func SVGTransformListCreateSVGTransformFromMatrix1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransformList_Consolidate
//go:noescape

func CallSVGTransformListConsolidate(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGTransformList_Consolidate
//go:noescape

func SVGTransformListConsolidateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedTransformList_BaseVal
//go:noescape

func GetSVGAnimatedTransformListBaseVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedTransformList_AnimVal
//go:noescape

func GetSVGAnimatedTransformListAnimVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGStringList_Length
//go:noescape

func GetSVGStringListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SVGStringList_NumberOfItems
//go:noescape

func GetSVGStringListNumberOfItems(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_SVGStringList_Clear
//go:noescape

func CallSVGStringListClear(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_Clear
//go:noescape

func SVGStringListClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGStringList_Initialize
//go:noescape

func CallSVGStringListInitialize(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_Initialize
//go:noescape

func SVGStringListInitializeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGStringList_GetItem
//go:noescape

func CallSVGStringListGetItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_GetItem
//go:noescape

func SVGStringListGetItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGStringList_InsertItemBefore
//go:noescape

func CallSVGStringListInsertItemBefore(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_InsertItemBefore
//go:noescape

func SVGStringListInsertItemBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGStringList_ReplaceItem
//go:noescape

func CallSVGStringListReplaceItem(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_ReplaceItem
//go:noescape

func SVGStringListReplaceItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGStringList_RemoveItem
//go:noescape

func CallSVGStringListRemoveItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_RemoveItem
//go:noescape

func SVGStringListRemoveItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGStringList_AppendItem
//go:noescape

func CallSVGStringListAppendItem(
	this js.Ref,
	ptr unsafe.Pointer,

	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_AppendItem
//go:noescape

func SVGStringListAppendItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGStringList_Set
//go:noescape

func CallSVGStringListSet(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
	newItem js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGStringList_Set
//go:noescape

func SVGStringListSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGAnimationElement_TargetElement
//go:noescape

func GetSVGAnimationElementTargetElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimationElement_RequiredExtensions
//go:noescape

func GetSVGAnimationElementRequiredExtensions(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimationElement_SystemLanguage
//go:noescape

func GetSVGAnimationElementSystemLanguage(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_SVGAnimationElement_GetStartTime
//go:noescape

func CallSVGAnimationElementGetStartTime(
	this js.Ref,
	ptr unsafe.Pointer,

) float32

//go:wasmimport plat/js/web func_SVGAnimationElement_GetStartTime
//go:noescape

func SVGAnimationElementGetStartTimeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGAnimationElement_GetCurrentTime
//go:noescape

func CallSVGAnimationElementGetCurrentTime(
	this js.Ref,
	ptr unsafe.Pointer,

) float32

//go:wasmimport plat/js/web func_SVGAnimationElement_GetCurrentTime
//go:noescape

func SVGAnimationElementGetCurrentTimeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGAnimationElement_GetSimpleDuration
//go:noescape

func CallSVGAnimationElementGetSimpleDuration(
	this js.Ref,
	ptr unsafe.Pointer,

) float32

//go:wasmimport plat/js/web func_SVGAnimationElement_GetSimpleDuration
//go:noescape

func SVGAnimationElementGetSimpleDurationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGAnimationElement_BeginElement
//go:noescape

func CallSVGAnimationElementBeginElement(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGAnimationElement_BeginElement
//go:noescape

func SVGAnimationElementBeginElementFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGAnimationElement_BeginElementAt
//go:noescape

func CallSVGAnimationElementBeginElementAt(
	this js.Ref,
	ptr unsafe.Pointer,

	offset float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGAnimationElement_BeginElementAt
//go:noescape

func SVGAnimationElementBeginElementAtFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGAnimationElement_EndElement
//go:noescape

func CallSVGAnimationElementEndElement(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGAnimationElement_EndElement
//go:noescape

func SVGAnimationElementEndElementFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGAnimationElement_EndElementAt
//go:noescape

func CallSVGAnimationElementEndElementAt(
	this js.Ref,
	ptr unsafe.Pointer,

	offset float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGAnimationElement_EndElementAt
//go:noescape

func SVGAnimationElementEndElementAtFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGCircleElement_Cy
//go:noescape

func GetSVGCircleElementCy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGCircleElement_R
//go:noescape

func GetSVGCircleElementR(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGClipPathElement_ClipPathUnits
//go:noescape

func GetSVGClipPathElementClipPathUnits(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGClipPathElement_Transform
//go:noescape

func GetSVGClipPathElementTransform(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Type
//go:noescape

func GetSVGComponentTransferFunctionElementType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_TableValues
//go:noescape

func GetSVGComponentTransferFunctionElementTableValues(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Slope
//go:noescape

func GetSVGComponentTransferFunctionElementSlope(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Intercept
//go:noescape

func GetSVGComponentTransferFunctionElementIntercept(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Amplitude
//go:noescape

func GetSVGComponentTransferFunctionElementAmplitude(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Exponent
//go:noescape

func GetSVGComponentTransferFunctionElementExponent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGComponentTransferFunctionElement_Offset
//go:noescape

func GetSVGComponentTransferFunctionElementOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGEllipseElement_Cx
//go:noescape

func GetSVGEllipseElementCx(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGEllipseElement_Cy
//go:noescape

func GetSVGEllipseElementCy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGEllipseElement_Rx
//go:noescape

func GetSVGEllipseElementRx(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGEllipseElement_Ry
//go:noescape

func GetSVGEllipseElementRy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEBlendElement_In1
//go:noescape

func GetSVGFEBlendElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEBlendElement_In2
//go:noescape

func GetSVGFEBlendElementIn2(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEBlendElement_Mode
//go:noescape

func GetSVGFEBlendElementMode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEBlendElement_X
//go:noescape

func GetSVGFEBlendElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEBlendElement_Y
//go:noescape

func GetSVGFEBlendElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEBlendElement_Width
//go:noescape

func GetSVGFEBlendElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEBlendElement_Height
//go:noescape

func GetSVGFEBlendElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEBlendElement_Result
//go:noescape

func GetSVGFEBlendElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_In1
//go:noescape

func GetSVGFEColorMatrixElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Type
//go:noescape

func GetSVGFEColorMatrixElementType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Values
//go:noescape

func GetSVGFEColorMatrixElementValues(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_X
//go:noescape

func GetSVGFEColorMatrixElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Y
//go:noescape

func GetSVGFEColorMatrixElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Width
//go:noescape

func GetSVGFEColorMatrixElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Height
//go:noescape

func GetSVGFEColorMatrixElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEColorMatrixElement_Result
//go:noescape

func GetSVGFEColorMatrixElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_In1
//go:noescape

func GetSVGFEComponentTransferElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_X
//go:noescape

func GetSVGFEComponentTransferElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_Y
//go:noescape

func GetSVGFEComponentTransferElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_Width
//go:noescape

func GetSVGFEComponentTransferElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_Height
//go:noescape

func GetSVGFEComponentTransferElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEComponentTransferElement_Result
//go:noescape

func GetSVGFEComponentTransferElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_In1
//go:noescape

func GetSVGFECompositeElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_In2
//go:noescape

func GetSVGFECompositeElementIn2(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_Operator
//go:noescape

func GetSVGFECompositeElementOperator(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_K1
//go:noescape

func GetSVGFECompositeElementK1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_K2
//go:noescape

func GetSVGFECompositeElementK2(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_K3
//go:noescape

func GetSVGFECompositeElementK3(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_K4
//go:noescape

func GetSVGFECompositeElementK4(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_X
//go:noescape

func GetSVGFECompositeElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_Y
//go:noescape

func GetSVGFECompositeElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_Width
//go:noescape

func GetSVGFECompositeElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_Height
//go:noescape

func GetSVGFECompositeElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFECompositeElement_Result
//go:noescape

func GetSVGFECompositeElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_In1
//go:noescape

func GetSVGFEConvolveMatrixElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_OrderX
//go:noescape

func GetSVGFEConvolveMatrixElementOrderX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_OrderY
//go:noescape

func GetSVGFEConvolveMatrixElementOrderY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_KernelMatrix
//go:noescape

func GetSVGFEConvolveMatrixElementKernelMatrix(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Divisor
//go:noescape

func GetSVGFEConvolveMatrixElementDivisor(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Bias
//go:noescape

func GetSVGFEConvolveMatrixElementBias(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_TargetX
//go:noescape

func GetSVGFEConvolveMatrixElementTargetX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_TargetY
//go:noescape

func GetSVGFEConvolveMatrixElementTargetY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_EdgeMode
//go:noescape

func GetSVGFEConvolveMatrixElementEdgeMode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_KernelUnitLengthX
//go:noescape

func GetSVGFEConvolveMatrixElementKernelUnitLengthX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_KernelUnitLengthY
//go:noescape

func GetSVGFEConvolveMatrixElementKernelUnitLengthY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_PreserveAlpha
//go:noescape

func GetSVGFEConvolveMatrixElementPreserveAlpha(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_X
//go:noescape

func GetSVGFEConvolveMatrixElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Y
//go:noescape

func GetSVGFEConvolveMatrixElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Width
//go:noescape

func GetSVGFEConvolveMatrixElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Height
//go:noescape

func GetSVGFEConvolveMatrixElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEConvolveMatrixElement_Result
//go:noescape

func GetSVGFEConvolveMatrixElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_In1
//go:noescape

func GetSVGFEDiffuseLightingElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_SurfaceScale
//go:noescape

func GetSVGFEDiffuseLightingElementSurfaceScale(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_DiffuseConstant
//go:noescape

func GetSVGFEDiffuseLightingElementDiffuseConstant(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_KernelUnitLengthX
//go:noescape

func GetSVGFEDiffuseLightingElementKernelUnitLengthX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_KernelUnitLengthY
//go:noescape

func GetSVGFEDiffuseLightingElementKernelUnitLengthY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_X
//go:noescape

func GetSVGFEDiffuseLightingElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_Y
//go:noescape

func GetSVGFEDiffuseLightingElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_Width
//go:noescape

func GetSVGFEDiffuseLightingElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_Height
//go:noescape

func GetSVGFEDiffuseLightingElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDiffuseLightingElement_Result
//go:noescape

func GetSVGFEDiffuseLightingElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_In1
//go:noescape

func GetSVGFEDisplacementMapElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_In2
//go:noescape

func GetSVGFEDisplacementMapElementIn2(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_Scale
//go:noescape

func GetSVGFEDisplacementMapElementScale(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_XChannelSelector
//go:noescape

func GetSVGFEDisplacementMapElementXChannelSelector(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_YChannelSelector
//go:noescape

func GetSVGFEDisplacementMapElementYChannelSelector(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_X
//go:noescape

func GetSVGFEDisplacementMapElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_Y
//go:noescape

func GetSVGFEDisplacementMapElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_Width
//go:noescape

func GetSVGFEDisplacementMapElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_Height
//go:noescape

func GetSVGFEDisplacementMapElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDisplacementMapElement_Result
//go:noescape

func GetSVGFEDisplacementMapElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDistantLightElement_Azimuth
//go:noescape

func GetSVGFEDistantLightElementAzimuth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDistantLightElement_Elevation
//go:noescape

func GetSVGFEDistantLightElementElevation(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_In1
//go:noescape

func GetSVGFEDropShadowElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Dx
//go:noescape

func GetSVGFEDropShadowElementDx(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Dy
//go:noescape

func GetSVGFEDropShadowElementDy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_StdDeviationX
//go:noescape

func GetSVGFEDropShadowElementStdDeviationX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_StdDeviationY
//go:noescape

func GetSVGFEDropShadowElementStdDeviationY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_X
//go:noescape

func GetSVGFEDropShadowElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Y
//go:noescape

func GetSVGFEDropShadowElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Width
//go:noescape

func GetSVGFEDropShadowElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Height
//go:noescape

func GetSVGFEDropShadowElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEDropShadowElement_Result
//go:noescape

func GetSVGFEDropShadowElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_SVGFEDropShadowElement_SetStdDeviation
//go:noescape

func CallSVGFEDropShadowElementSetStdDeviation(
	this js.Ref,
	ptr unsafe.Pointer,

	stdDeviationX float32,
	stdDeviationY float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGFEDropShadowElement_SetStdDeviation
//go:noescape

func SVGFEDropShadowElementSetStdDeviationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGFEFloodElement_X
//go:noescape

func GetSVGFEFloodElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEFloodElement_Y
//go:noescape

func GetSVGFEFloodElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEFloodElement_Width
//go:noescape

func GetSVGFEFloodElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEFloodElement_Height
//go:noescape

func GetSVGFEFloodElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEFloodElement_Result
//go:noescape

func GetSVGFEFloodElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_In1
//go:noescape

func GetSVGFEGaussianBlurElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_StdDeviationX
//go:noescape

func GetSVGFEGaussianBlurElementStdDeviationX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_StdDeviationY
//go:noescape

func GetSVGFEGaussianBlurElementStdDeviationY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_EdgeMode
//go:noescape

func GetSVGFEGaussianBlurElementEdgeMode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_X
//go:noescape

func GetSVGFEGaussianBlurElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_Y
//go:noescape

func GetSVGFEGaussianBlurElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_Width
//go:noescape

func GetSVGFEGaussianBlurElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_Height
//go:noescape

func GetSVGFEGaussianBlurElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEGaussianBlurElement_Result
//go:noescape

func GetSVGFEGaussianBlurElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_SVGFEGaussianBlurElement_SetStdDeviation
//go:noescape

func CallSVGFEGaussianBlurElementSetStdDeviation(
	this js.Ref,
	ptr unsafe.Pointer,

	stdDeviationX float32,
	stdDeviationY float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGFEGaussianBlurElement_SetStdDeviation
//go:noescape

func SVGFEGaussianBlurElementSetStdDeviationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGFEImageElement_PreserveAspectRatio
//go:noescape

func GetSVGFEImageElementPreserveAspectRatio(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEImageElement_CrossOrigin
//go:noescape

func GetSVGFEImageElementCrossOrigin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEImageElement_X
//go:noescape

func GetSVGFEImageElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEImageElement_Y
//go:noescape

func GetSVGFEImageElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEImageElement_Width
//go:noescape

func GetSVGFEImageElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEImageElement_Height
//go:noescape

func GetSVGFEImageElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEImageElement_Result
//go:noescape

func GetSVGFEImageElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEImageElement_Href
//go:noescape

func GetSVGFEImageElementHref(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMergeElement_X
//go:noescape

func GetSVGFEMergeElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMergeElement_Y
//go:noescape

func GetSVGFEMergeElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMergeElement_Width
//go:noescape

func GetSVGFEMergeElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMergeElement_Height
//go:noescape

func GetSVGFEMergeElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMergeElement_Result
//go:noescape

func GetSVGFEMergeElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMergeNodeElement_In1
//go:noescape

func GetSVGFEMergeNodeElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_In1
//go:noescape

func GetSVGFEMorphologyElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_Operator
//go:noescape

func GetSVGFEMorphologyElementOperator(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_RadiusX
//go:noescape

func GetSVGFEMorphologyElementRadiusX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_RadiusY
//go:noescape

func GetSVGFEMorphologyElementRadiusY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_X
//go:noescape

func GetSVGFEMorphologyElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_Y
//go:noescape

func GetSVGFEMorphologyElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_Width
//go:noescape

func GetSVGFEMorphologyElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_Height
//go:noescape

func GetSVGFEMorphologyElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEMorphologyElement_Result
//go:noescape

func GetSVGFEMorphologyElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEOffsetElement_In1
//go:noescape

func GetSVGFEOffsetElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Dx
//go:noescape

func GetSVGFEOffsetElementDx(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Dy
//go:noescape

func GetSVGFEOffsetElementDy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEOffsetElement_X
//go:noescape

func GetSVGFEOffsetElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Y
//go:noescape

func GetSVGFEOffsetElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Width
//go:noescape

func GetSVGFEOffsetElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Height
//go:noescape

func GetSVGFEOffsetElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEOffsetElement_Result
//go:noescape

func GetSVGFEOffsetElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEPointLightElement_X
//go:noescape

func GetSVGFEPointLightElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEPointLightElement_Y
//go:noescape

func GetSVGFEPointLightElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFEPointLightElement_Z
//go:noescape

func GetSVGFEPointLightElementZ(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_In1
//go:noescape

func GetSVGFESpecularLightingElementIn1(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_SurfaceScale
//go:noescape

func GetSVGFESpecularLightingElementSurfaceScale(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_SpecularConstant
//go:noescape

func GetSVGFESpecularLightingElementSpecularConstant(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_SpecularExponent
//go:noescape

func GetSVGFESpecularLightingElementSpecularExponent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_KernelUnitLengthX
//go:noescape

func GetSVGFESpecularLightingElementKernelUnitLengthX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_KernelUnitLengthY
//go:noescape

func GetSVGFESpecularLightingElementKernelUnitLengthY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_X
//go:noescape

func GetSVGFESpecularLightingElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_Y
//go:noescape

func GetSVGFESpecularLightingElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_Width
//go:noescape

func GetSVGFESpecularLightingElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_Height
//go:noescape

func GetSVGFESpecularLightingElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpecularLightingElement_Result
//go:noescape

func GetSVGFESpecularLightingElementResult(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpotLightElement_X
//go:noescape

func GetSVGFESpotLightElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpotLightElement_Y
//go:noescape

func GetSVGFESpotLightElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpotLightElement_Z
//go:noescape

func GetSVGFESpotLightElementZ(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpotLightElement_PointsAtX
//go:noescape

func GetSVGFESpotLightElementPointsAtX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpotLightElement_PointsAtY
//go:noescape

func GetSVGFESpotLightElementPointsAtY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpotLightElement_PointsAtZ
//go:noescape

func GetSVGFESpotLightElementPointsAtZ(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpotLightElement_SpecularExponent
//go:noescape

func GetSVGFESpotLightElementSpecularExponent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGFESpotLightElement_LimitingConeAngle
//go:noescape

func GetSVGFESpotLightElementLimitingConeAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

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

//go:wasmimport plat/js/web get_Baseline_Name
//go:noescape
func GetBaselineName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Baseline_Value
//go:noescape
func GetBaselineValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Font_Name
//go:noescape
func GetFontName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Font_GlyphsRendered
//go:noescape
func GetFontGlyphsRendered(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_Width
//go:noescape
func GetFontMetricsWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_Advances
//go:noescape
func GetFontMetricsAdvances(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_BoundingBoxLeft
//go:noescape
func GetFontMetricsBoundingBoxLeft(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_BoundingBoxRight
//go:noescape
func GetFontMetricsBoundingBoxRight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_Height
//go:noescape
func GetFontMetricsHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_EmHeightAscent
//go:noescape
func GetFontMetricsEmHeightAscent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_EmHeightDescent
//go:noescape
func GetFontMetricsEmHeightDescent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_BoundingBoxAscent
//go:noescape
func GetFontMetricsBoundingBoxAscent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_BoundingBoxDescent
//go:noescape
func GetFontMetricsBoundingBoxDescent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_FontBoundingBoxAscent
//go:noescape
func GetFontMetricsFontBoundingBoxAscent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_FontBoundingBoxDescent
//go:noescape
func GetFontMetricsFontBoundingBoxDescent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_DominantBaseline
//go:noescape
func GetFontMetricsDominantBaseline(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_Baselines
//go:noescape
func GetFontMetricsBaselines(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_FontMetrics_Fonts
//go:noescape
func GetFontMetricsFonts(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_StaticRangeInit
//go:noescape
func StaticRangeInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_StaticRangeInit
//go:noescape
func StaticRangeInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_StaticRange_StaticRange
//go:noescape
func NewStaticRangeByStaticRange(
	init unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web get_Selection_AnchorNode
//go:noescape
func GetSelectionAnchorNode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Selection_AnchorOffset
//go:noescape
func GetSelectionAnchorOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Selection_FocusNode
//go:noescape
func GetSelectionFocusNode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Selection_FocusOffset
//go:noescape
func GetSelectionFocusOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Selection_IsCollapsed
//go:noescape
func GetSelectionIsCollapsed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Selection_RangeCount
//go:noescape
func GetSelectionRangeCount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Selection_Type
//go:noescape
func GetSelectionType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_Selection_Direction
//go:noescape
func GetSelectionDirection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_GetRangeAt
//go:noescape
func HasSelectionGetRangeAt(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_GetRangeAt
//go:noescape
func SelectionGetRangeAtFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_GetRangeAt
//go:noescape
func CallSelectionGetRangeAt(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_Selection_GetRangeAt
//go:noescape
func TrySelectionGetRangeAt(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_AddRange
//go:noescape
func HasSelectionAddRange(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_AddRange
//go:noescape
func SelectionAddRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_AddRange
//go:noescape
func CallSelectionAddRange(
	this js.Ref, retPtr unsafe.Pointer,
	rng js.Ref)

//go:wasmimport plat/js/web try_Selection_AddRange
//go:noescape
func TrySelectionAddRange(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rng js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_RemoveRange
//go:noescape
func HasSelectionRemoveRange(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_RemoveRange
//go:noescape
func SelectionRemoveRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_RemoveRange
//go:noescape
func CallSelectionRemoveRange(
	this js.Ref, retPtr unsafe.Pointer,
	rng js.Ref)

//go:wasmimport plat/js/web try_Selection_RemoveRange
//go:noescape
func TrySelectionRemoveRange(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rng js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_RemoveAllRanges
//go:noescape
func HasSelectionRemoveAllRanges(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_RemoveAllRanges
//go:noescape
func SelectionRemoveAllRangesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_RemoveAllRanges
//go:noescape
func CallSelectionRemoveAllRanges(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Selection_RemoveAllRanges
//go:noescape
func TrySelectionRemoveAllRanges(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_Empty
//go:noescape
func HasSelectionEmpty(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Empty
//go:noescape
func SelectionEmptyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Empty
//go:noescape
func CallSelectionEmpty(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Selection_Empty
//go:noescape
func TrySelectionEmpty(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_GetComposedRanges
//go:noescape
func HasSelectionGetComposedRanges(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_GetComposedRanges
//go:noescape
func SelectionGetComposedRangesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_GetComposedRanges
//go:noescape
func CallSelectionGetComposedRanges(
	this js.Ref, retPtr unsafe.Pointer,
	shadowRootsPtr unsafe.Pointer,
	shadowRootsCount uint32)

//go:wasmimport plat/js/web try_Selection_GetComposedRanges
//go:noescape
func TrySelectionGetComposedRanges(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	shadowRootsPtr unsafe.Pointer,
	shadowRootsCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_Collapse
//go:noescape
func HasSelectionCollapse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Collapse
//go:noescape
func SelectionCollapseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Collapse
//go:noescape
func CallSelectionCollapse(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref,
	offset uint32)

//go:wasmimport plat/js/web try_Selection_Collapse
//go:noescape
func TrySelectionCollapse(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref,
	offset uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_Collapse1
//go:noescape
func HasSelectionCollapse1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Collapse1
//go:noescape
func SelectionCollapse1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Collapse1
//go:noescape
func CallSelectionCollapse1(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref)

//go:wasmimport plat/js/web try_Selection_Collapse1
//go:noescape
func TrySelectionCollapse1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_SetPosition
//go:noescape
func HasSelectionSetPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_SetPosition
//go:noescape
func SelectionSetPositionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_SetPosition
//go:noescape
func CallSelectionSetPosition(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref,
	offset uint32)

//go:wasmimport plat/js/web try_Selection_SetPosition
//go:noescape
func TrySelectionSetPosition(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref,
	offset uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_SetPosition1
//go:noescape
func HasSelectionSetPosition1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_SetPosition1
//go:noescape
func SelectionSetPosition1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_SetPosition1
//go:noescape
func CallSelectionSetPosition1(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref)

//go:wasmimport plat/js/web try_Selection_SetPosition1
//go:noescape
func TrySelectionSetPosition1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_CollapseToStart
//go:noescape
func HasSelectionCollapseToStart(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_CollapseToStart
//go:noescape
func SelectionCollapseToStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_CollapseToStart
//go:noescape
func CallSelectionCollapseToStart(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Selection_CollapseToStart
//go:noescape
func TrySelectionCollapseToStart(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_CollapseToEnd
//go:noescape
func HasSelectionCollapseToEnd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_CollapseToEnd
//go:noescape
func SelectionCollapseToEndFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_CollapseToEnd
//go:noescape
func CallSelectionCollapseToEnd(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Selection_CollapseToEnd
//go:noescape
func TrySelectionCollapseToEnd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_Extend
//go:noescape
func HasSelectionExtend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Extend
//go:noescape
func SelectionExtendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Extend
//go:noescape
func CallSelectionExtend(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref,
	offset uint32)

//go:wasmimport plat/js/web try_Selection_Extend
//go:noescape
func TrySelectionExtend(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref,
	offset uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_Extend1
//go:noescape
func HasSelectionExtend1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Extend1
//go:noescape
func SelectionExtend1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Extend1
//go:noescape
func CallSelectionExtend1(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref)

//go:wasmimport plat/js/web try_Selection_Extend1
//go:noescape
func TrySelectionExtend1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_SetBaseAndExtent
//go:noescape
func HasSelectionSetBaseAndExtent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_SetBaseAndExtent
//go:noescape
func SelectionSetBaseAndExtentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_SetBaseAndExtent
//go:noescape
func CallSelectionSetBaseAndExtent(
	this js.Ref, retPtr unsafe.Pointer,
	anchorNode js.Ref,
	anchorOffset uint32,
	focusNode js.Ref,
	focusOffset uint32)

//go:wasmimport plat/js/web try_Selection_SetBaseAndExtent
//go:noescape
func TrySelectionSetBaseAndExtent(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	anchorNode js.Ref,
	anchorOffset uint32,
	focusNode js.Ref,
	focusOffset uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_SelectAllChildren
//go:noescape
func HasSelectionSelectAllChildren(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_SelectAllChildren
//go:noescape
func SelectionSelectAllChildrenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_SelectAllChildren
//go:noescape
func CallSelectionSelectAllChildren(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref)

//go:wasmimport plat/js/web try_Selection_SelectAllChildren
//go:noescape
func TrySelectionSelectAllChildren(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_Modify
//go:noescape
func HasSelectionModify(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify
//go:noescape
func SelectionModifyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Modify
//go:noescape
func CallSelectionModify(
	this js.Ref, retPtr unsafe.Pointer,
	alter js.Ref,
	direction js.Ref,
	granularity js.Ref)

//go:wasmimport plat/js/web try_Selection_Modify
//go:noescape
func TrySelectionModify(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	alter js.Ref,
	direction js.Ref,
	granularity js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_Modify1
//go:noescape
func HasSelectionModify1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify1
//go:noescape
func SelectionModify1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Modify1
//go:noescape
func CallSelectionModify1(
	this js.Ref, retPtr unsafe.Pointer,
	alter js.Ref,
	direction js.Ref)

//go:wasmimport plat/js/web try_Selection_Modify1
//go:noescape
func TrySelectionModify1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	alter js.Ref,
	direction js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_Modify2
//go:noescape
func HasSelectionModify2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify2
//go:noescape
func SelectionModify2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Modify2
//go:noescape
func CallSelectionModify2(
	this js.Ref, retPtr unsafe.Pointer,
	alter js.Ref)

//go:wasmimport plat/js/web try_Selection_Modify2
//go:noescape
func TrySelectionModify2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	alter js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_Modify3
//go:noescape
func HasSelectionModify3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify3
//go:noescape
func SelectionModify3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Modify3
//go:noescape
func CallSelectionModify3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Selection_Modify3
//go:noescape
func TrySelectionModify3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_DeleteFromDocument
//go:noescape
func HasSelectionDeleteFromDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_DeleteFromDocument
//go:noescape
func SelectionDeleteFromDocumentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_DeleteFromDocument
//go:noescape
func CallSelectionDeleteFromDocument(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Selection_DeleteFromDocument
//go:noescape
func TrySelectionDeleteFromDocument(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_ContainsNode
//go:noescape
func HasSelectionContainsNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_ContainsNode
//go:noescape
func SelectionContainsNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_ContainsNode
//go:noescape
func CallSelectionContainsNode(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref,
	allowPartialContainment js.Ref)

//go:wasmimport plat/js/web try_Selection_ContainsNode
//go:noescape
func TrySelectionContainsNode(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref,
	allowPartialContainment js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_ContainsNode1
//go:noescape
func HasSelectionContainsNode1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_ContainsNode1
//go:noescape
func SelectionContainsNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_ContainsNode1
//go:noescape
func CallSelectionContainsNode1(
	this js.Ref, retPtr unsafe.Pointer,
	node js.Ref)

//go:wasmimport plat/js/web try_Selection_ContainsNode1
//go:noescape
func TrySelectionContainsNode1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	node js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_Selection_ToString
//go:noescape
func HasSelectionToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_ToString
//go:noescape
func SelectionToStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_ToString
//go:noescape
func CallSelectionToString(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_Selection_ToString
//go:noescape
func TrySelectionToString(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CaretPosition_OffsetNode
//go:noescape
func GetCaretPositionOffsetNode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CaretPosition_Offset
//go:noescape
func GetCaretPositionOffset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_CaretPosition_GetClientRect
//go:noescape
func HasCaretPositionGetClientRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CaretPosition_GetClientRect
//go:noescape
func CaretPositionGetClientRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CaretPosition_GetClientRect
//go:noescape
func CallCaretPositionGetClientRect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_CaretPosition_GetClientRect
//go:noescape
func TryCaretPositionGetClientRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XPathResult_ResultType
//go:noescape
func GetXPathResultResultType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XPathResult_NumberValue
//go:noescape
func GetXPathResultNumberValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XPathResult_StringValue
//go:noescape
func GetXPathResultStringValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XPathResult_BooleanValue
//go:noescape
func GetXPathResultBooleanValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XPathResult_SingleNodeValue
//go:noescape
func GetXPathResultSingleNodeValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XPathResult_InvalidIteratorState
//go:noescape
func GetXPathResultInvalidIteratorState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_XPathResult_SnapshotLength
//go:noescape
func GetXPathResultSnapshotLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathResult_IterateNext
//go:noescape
func HasXPathResultIterateNext(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathResult_IterateNext
//go:noescape
func XPathResultIterateNextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathResult_IterateNext
//go:noescape
func CallXPathResultIterateNext(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_XPathResult_IterateNext
//go:noescape
func TryXPathResultIterateNext(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathResult_SnapshotItem
//go:noescape
func HasXPathResultSnapshotItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathResult_SnapshotItem
//go:noescape
func XPathResultSnapshotItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathResult_SnapshotItem
//go:noescape
func CallXPathResultSnapshotItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_XPathResult_SnapshotItem
//go:noescape
func TryXPathResultSnapshotItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathExpression_Evaluate
//go:noescape
func HasXPathExpressionEvaluate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathExpression_Evaluate
//go:noescape
func XPathExpressionEvaluateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathExpression_Evaluate
//go:noescape
func CallXPathExpressionEvaluate(
	this js.Ref, retPtr unsafe.Pointer,
	contextNode js.Ref,
	typ uint32,
	result js.Ref)

//go:wasmimport plat/js/web try_XPathExpression_Evaluate
//go:noescape
func TryXPathExpressionEvaluate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	contextNode js.Ref,
	typ uint32,
	result js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathExpression_Evaluate1
//go:noescape
func HasXPathExpressionEvaluate1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathExpression_Evaluate1
//go:noescape
func XPathExpressionEvaluate1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathExpression_Evaluate1
//go:noescape
func CallXPathExpressionEvaluate1(
	this js.Ref, retPtr unsafe.Pointer,
	contextNode js.Ref,
	typ uint32)

//go:wasmimport plat/js/web try_XPathExpression_Evaluate1
//go:noescape
func TryXPathExpressionEvaluate1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	contextNode js.Ref,
	typ uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_XPathExpression_Evaluate2
//go:noescape
func HasXPathExpressionEvaluate2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathExpression_Evaluate2
//go:noescape
func XPathExpressionEvaluate2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathExpression_Evaluate2
//go:noescape
func CallXPathExpressionEvaluate2(
	this js.Ref, retPtr unsafe.Pointer,
	contextNode js.Ref)

//go:wasmimport plat/js/web try_XPathExpression_Evaluate2
//go:noescape
func TryXPathExpressionEvaluate2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	contextNode js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_DocumentType_Name
//go:noescape
func GetDocumentTypeName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DocumentType_PublicId
//go:noescape
func GetDocumentTypePublicId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DocumentType_SystemId
//go:noescape
func GetDocumentTypeSystemId(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DocumentType_Before
//go:noescape
func HasDocumentTypeBefore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DocumentType_Before
//go:noescape
func DocumentTypeBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentType_Before
//go:noescape
func CallDocumentTypeBefore(
	this js.Ref, retPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32)

//go:wasmimport plat/js/web try_DocumentType_Before
//go:noescape
func TryDocumentTypeBefore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_DocumentType_After
//go:noescape
func HasDocumentTypeAfter(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DocumentType_After
//go:noescape
func DocumentTypeAfterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentType_After
//go:noescape
func CallDocumentTypeAfter(
	this js.Ref, retPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32)

//go:wasmimport plat/js/web try_DocumentType_After
//go:noescape
func TryDocumentTypeAfter(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_DocumentType_ReplaceWith
//go:noescape
func HasDocumentTypeReplaceWith(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DocumentType_ReplaceWith
//go:noescape
func DocumentTypeReplaceWithFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentType_ReplaceWith
//go:noescape
func CallDocumentTypeReplaceWith(
	this js.Ref, retPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32)

//go:wasmimport plat/js/web try_DocumentType_ReplaceWith
//go:noescape
func TryDocumentTypeReplaceWith(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	nodesPtr unsafe.Pointer,
	nodesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_DocumentType_Remove
//go:noescape
func HasDocumentTypeRemove(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DocumentType_Remove
//go:noescape
func DocumentTypeRemoveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentType_Remove
//go:noescape
func CallDocumentTypeRemove(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DocumentType_Remove
//go:noescape
func TryDocumentTypeRemove(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMImplementation_CreateDocumentType
//go:noescape
func HasDOMImplementationCreateDocumentType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateDocumentType
//go:noescape
func DOMImplementationCreateDocumentTypeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_CreateDocumentType
//go:noescape
func CallDOMImplementationCreateDocumentType(
	this js.Ref, retPtr unsafe.Pointer,
	qualifiedName js.Ref,
	publicId js.Ref,
	systemId js.Ref)

//go:wasmimport plat/js/web try_DOMImplementation_CreateDocumentType
//go:noescape
func TryDOMImplementationCreateDocumentType(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	qualifiedName js.Ref,
	publicId js.Ref,
	systemId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMImplementation_CreateDocument
//go:noescape
func HasDOMImplementationCreateDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateDocument
//go:noescape
func DOMImplementationCreateDocumentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_CreateDocument
//go:noescape
func CallDOMImplementationCreateDocument(
	this js.Ref, retPtr unsafe.Pointer,
	namespace js.Ref,
	qualifiedName js.Ref,
	doctype js.Ref)

//go:wasmimport plat/js/web try_DOMImplementation_CreateDocument
//go:noescape
func TryDOMImplementationCreateDocument(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	namespace js.Ref,
	qualifiedName js.Ref,
	doctype js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMImplementation_CreateDocument1
//go:noescape
func HasDOMImplementationCreateDocument1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateDocument1
//go:noescape
func DOMImplementationCreateDocument1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_CreateDocument1
//go:noescape
func CallDOMImplementationCreateDocument1(
	this js.Ref, retPtr unsafe.Pointer,
	namespace js.Ref,
	qualifiedName js.Ref)

//go:wasmimport plat/js/web try_DOMImplementation_CreateDocument1
//go:noescape
func TryDOMImplementationCreateDocument1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	namespace js.Ref,
	qualifiedName js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMImplementation_CreateHTMLDocument
//go:noescape
func HasDOMImplementationCreateHTMLDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateHTMLDocument
//go:noescape
func DOMImplementationCreateHTMLDocumentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_CreateHTMLDocument
//go:noescape
func CallDOMImplementationCreateHTMLDocument(
	this js.Ref, retPtr unsafe.Pointer,
	title js.Ref)

//go:wasmimport plat/js/web try_DOMImplementation_CreateHTMLDocument
//go:noescape
func TryDOMImplementationCreateHTMLDocument(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	title js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMImplementation_CreateHTMLDocument1
//go:noescape
func HasDOMImplementationCreateHTMLDocument1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateHTMLDocument1
//go:noescape
func DOMImplementationCreateHTMLDocument1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_CreateHTMLDocument1
//go:noescape
func CallDOMImplementationCreateHTMLDocument1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMImplementation_CreateHTMLDocument1
//go:noescape
func TryDOMImplementationCreateHTMLDocument1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMImplementation_HasFeature
//go:noescape
func HasDOMImplementationHasFeature(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_HasFeature
//go:noescape
func DOMImplementationHasFeatureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_HasFeature
//go:noescape
func CallDOMImplementationHasFeature(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMImplementation_HasFeature
//go:noescape
func TryDOMImplementationHasFeature(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PermissionsPolicy_AllowsFeature
//go:noescape
func HasPermissionsPolicyAllowsFeature(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_AllowsFeature
//go:noescape
func PermissionsPolicyAllowsFeatureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PermissionsPolicy_AllowsFeature
//go:noescape
func CallPermissionsPolicyAllowsFeature(
	this js.Ref, retPtr unsafe.Pointer,
	feature js.Ref,
	origin js.Ref)

//go:wasmimport plat/js/web try_PermissionsPolicy_AllowsFeature
//go:noescape
func TryPermissionsPolicyAllowsFeature(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	feature js.Ref,
	origin js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_PermissionsPolicy_AllowsFeature1
//go:noescape
func HasPermissionsPolicyAllowsFeature1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_AllowsFeature1
//go:noescape
func PermissionsPolicyAllowsFeature1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PermissionsPolicy_AllowsFeature1
//go:noescape
func CallPermissionsPolicyAllowsFeature1(
	this js.Ref, retPtr unsafe.Pointer,
	feature js.Ref)

//go:wasmimport plat/js/web try_PermissionsPolicy_AllowsFeature1
//go:noescape
func TryPermissionsPolicyAllowsFeature1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	feature js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_PermissionsPolicy_Features
//go:noescape
func HasPermissionsPolicyFeatures(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_Features
//go:noescape
func PermissionsPolicyFeaturesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PermissionsPolicy_Features
//go:noescape
func CallPermissionsPolicyFeatures(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PermissionsPolicy_Features
//go:noescape
func TryPermissionsPolicyFeatures(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PermissionsPolicy_AllowedFeatures
//go:noescape
func HasPermissionsPolicyAllowedFeatures(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_AllowedFeatures
//go:noescape
func PermissionsPolicyAllowedFeaturesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PermissionsPolicy_AllowedFeatures
//go:noescape
func CallPermissionsPolicyAllowedFeatures(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_PermissionsPolicy_AllowedFeatures
//go:noescape
func TryPermissionsPolicyAllowedFeatures(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_PermissionsPolicy_GetAllowlistForFeature
//go:noescape
func HasPermissionsPolicyGetAllowlistForFeature(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_GetAllowlistForFeature
//go:noescape
func PermissionsPolicyGetAllowlistForFeatureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PermissionsPolicy_GetAllowlistForFeature
//go:noescape
func CallPermissionsPolicyGetAllowlistForFeature(
	this js.Ref, retPtr unsafe.Pointer,
	feature js.Ref)

//go:wasmimport plat/js/web try_PermissionsPolicy_GetAllowlistForFeature
//go:noescape
func TryPermissionsPolicyGetAllowlistForFeature(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	feature js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_DocumentTimelineOptions
//go:noescape
func DocumentTimelineOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DocumentTimelineOptions
//go:noescape
func DocumentTimelineOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_DocumentTimeline_DocumentTimeline
//go:noescape
func NewDocumentTimelineByDocumentTimeline(
	options unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_DocumentTimeline_DocumentTimeline1
//go:noescape
func NewDocumentTimelineByDocumentTimeline1() js.Ref

//go:wasmimport plat/js/web store_FocusOptions
//go:noescape
func FocusOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_FocusOptions
//go:noescape
func FocusOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedString_BaseVal
//go:noescape
func GetSVGAnimatedStringBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAnimatedString_BaseVal
//go:noescape
func SetSVGAnimatedStringBaseVal(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedString_AnimVal
//go:noescape
func GetSVGAnimatedStringAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSStyleDeclaration_CssText
//go:noescape
func GetCSSStyleDeclarationCssText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSStyleDeclaration_CssText
//go:noescape
func SetCSSStyleDeclarationCssText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSStyleDeclaration_Length
//go:noescape
func GetCSSStyleDeclarationLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSStyleDeclaration_ParentRule
//go:noescape
func GetCSSStyleDeclarationParentRule(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_CSSStyleDeclaration_CssFloat
//go:noescape
func GetCSSStyleDeclarationCssFloat(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_CSSStyleDeclaration_CssFloat
//go:noescape
func SetCSSStyleDeclarationCssFloat(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_CSSStyleDeclaration_Item
//go:noescape
func HasCSSStyleDeclarationItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_Item
//go:noescape
func CSSStyleDeclarationItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_Item
//go:noescape
func CallCSSStyleDeclarationItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_CSSStyleDeclaration_Item
//go:noescape
func TryCSSStyleDeclarationItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSStyleDeclaration_GetPropertyValue
//go:noescape
func HasCSSStyleDeclarationGetPropertyValue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_GetPropertyValue
//go:noescape
func CSSStyleDeclarationGetPropertyValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_GetPropertyValue
//go:noescape
func CallCSSStyleDeclarationGetPropertyValue(
	this js.Ref, retPtr unsafe.Pointer,
	property js.Ref)

//go:wasmimport plat/js/web try_CSSStyleDeclaration_GetPropertyValue
//go:noescape
func TryCSSStyleDeclarationGetPropertyValue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	property js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSStyleDeclaration_GetPropertyPriority
//go:noescape
func HasCSSStyleDeclarationGetPropertyPriority(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_GetPropertyPriority
//go:noescape
func CSSStyleDeclarationGetPropertyPriorityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_GetPropertyPriority
//go:noescape
func CallCSSStyleDeclarationGetPropertyPriority(
	this js.Ref, retPtr unsafe.Pointer,
	property js.Ref)

//go:wasmimport plat/js/web try_CSSStyleDeclaration_GetPropertyPriority
//go:noescape
func TryCSSStyleDeclarationGetPropertyPriority(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	property js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSStyleDeclaration_SetProperty
//go:noescape
func HasCSSStyleDeclarationSetProperty(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_SetProperty
//go:noescape
func CSSStyleDeclarationSetPropertyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_SetProperty
//go:noescape
func CallCSSStyleDeclarationSetProperty(
	this js.Ref, retPtr unsafe.Pointer,
	property js.Ref,
	value js.Ref,
	priority js.Ref)

//go:wasmimport plat/js/web try_CSSStyleDeclaration_SetProperty
//go:noescape
func TryCSSStyleDeclarationSetProperty(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	property js.Ref,
	value js.Ref,
	priority js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSStyleDeclaration_SetProperty1
//go:noescape
func HasCSSStyleDeclarationSetProperty1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_SetProperty1
//go:noescape
func CSSStyleDeclarationSetProperty1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_SetProperty1
//go:noescape
func CallCSSStyleDeclarationSetProperty1(
	this js.Ref, retPtr unsafe.Pointer,
	property js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_CSSStyleDeclaration_SetProperty1
//go:noescape
func TryCSSStyleDeclarationSetProperty1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	property js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_CSSStyleDeclaration_RemoveProperty
//go:noescape
func HasCSSStyleDeclarationRemoveProperty(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_RemoveProperty
//go:noescape
func CSSStyleDeclarationRemovePropertyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_RemoveProperty
//go:noescape
func CallCSSStyleDeclarationRemoveProperty(
	this js.Ref, retPtr unsafe.Pointer,
	property js.Ref)

//go:wasmimport plat/js/web try_CSSStyleDeclaration_RemoveProperty
//go:noescape
func TryCSSStyleDeclarationRemoveProperty(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	property js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_StylePropertyMap_Set
//go:noescape
func HasStylePropertyMapSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Set
//go:noescape
func StylePropertyMapSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StylePropertyMap_Set
//go:noescape
func CallStylePropertyMapSet(
	this js.Ref, retPtr unsafe.Pointer,
	property js.Ref,
	valuesPtr unsafe.Pointer,
	valuesCount uint32)

//go:wasmimport plat/js/web try_StylePropertyMap_Set
//go:noescape
func TryStylePropertyMapSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	property js.Ref,
	valuesPtr unsafe.Pointer,
	valuesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_StylePropertyMap_Append
//go:noescape
func HasStylePropertyMapAppend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Append
//go:noescape
func StylePropertyMapAppendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StylePropertyMap_Append
//go:noescape
func CallStylePropertyMapAppend(
	this js.Ref, retPtr unsafe.Pointer,
	property js.Ref,
	valuesPtr unsafe.Pointer,
	valuesCount uint32)

//go:wasmimport plat/js/web try_StylePropertyMap_Append
//go:noescape
func TryStylePropertyMapAppend(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	property js.Ref,
	valuesPtr unsafe.Pointer,
	valuesCount uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_StylePropertyMap_Delete
//go:noescape
func HasStylePropertyMapDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Delete
//go:noescape
func StylePropertyMapDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StylePropertyMap_Delete
//go:noescape
func CallStylePropertyMapDelete(
	this js.Ref, retPtr unsafe.Pointer,
	property js.Ref)

//go:wasmimport plat/js/web try_StylePropertyMap_Delete
//go:noescape
func TryStylePropertyMapDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	property js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_StylePropertyMap_Clear
//go:noescape
func HasStylePropertyMapClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Clear
//go:noescape
func StylePropertyMapClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StylePropertyMap_Clear
//go:noescape
func CallStylePropertyMapClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_StylePropertyMap_Clear
//go:noescape
func TryStylePropertyMapClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLength_UnitType
//go:noescape
func GetSVGLengthUnitType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGLength_Value
//go:noescape
func GetSVGLengthValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGLength_Value
//go:noescape
func SetSVGLengthValue(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGLength_ValueInSpecifiedUnits
//go:noescape
func GetSVGLengthValueInSpecifiedUnits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGLength_ValueInSpecifiedUnits
//go:noescape
func SetSVGLengthValueInSpecifiedUnits(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGLength_ValueAsString
//go:noescape
func GetSVGLengthValueAsString(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGLength_ValueAsString
//go:noescape
func SetSVGLengthValueAsString(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_SVGLength_NewValueSpecifiedUnits
//go:noescape
func HasSVGLengthNewValueSpecifiedUnits(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLength_NewValueSpecifiedUnits
//go:noescape
func SVGLengthNewValueSpecifiedUnitsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGLength_NewValueSpecifiedUnits
//go:noescape
func CallSVGLengthNewValueSpecifiedUnits(
	this js.Ref, retPtr unsafe.Pointer,
	unitType uint32,
	valueInSpecifiedUnits float32)

//go:wasmimport plat/js/web try_SVGLength_NewValueSpecifiedUnits
//go:noescape
func TrySVGLengthNewValueSpecifiedUnits(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	unitType uint32,
	valueInSpecifiedUnits float32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGLength_ConvertToSpecifiedUnits
//go:noescape
func HasSVGLengthConvertToSpecifiedUnits(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLength_ConvertToSpecifiedUnits
//go:noescape
func SVGLengthConvertToSpecifiedUnitsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGLength_ConvertToSpecifiedUnits
//go:noescape
func CallSVGLengthConvertToSpecifiedUnits(
	this js.Ref, retPtr unsafe.Pointer,
	unitType uint32)

//go:wasmimport plat/js/web try_SVGLength_ConvertToSpecifiedUnits
//go:noescape
func TrySVGLengthConvertToSpecifiedUnits(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	unitType uint32) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedLength_BaseVal
//go:noescape
func GetSVGAnimatedLengthBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedLength_AnimVal
//go:noescape
func GetSVGAnimatedLengthAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGUseElement_X
//go:noescape
func GetSVGUseElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGUseElement_Y
//go:noescape
func GetSVGUseElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGUseElement_Width
//go:noescape
func GetSVGUseElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGUseElement_Height
//go:noescape
func GetSVGUseElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGUseElement_InstanceRoot
//go:noescape
func GetSVGUseElementInstanceRoot(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGUseElement_AnimatedInstanceRoot
//go:noescape
func GetSVGUseElementAnimatedInstanceRoot(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGUseElement_Href
//go:noescape
func GetSVGUseElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMStringMap_Get
//go:noescape
func HasDOMStringMapGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMStringMap_Get
//go:noescape
func DOMStringMapGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMStringMap_Get
//go:noescape
func CallDOMStringMapGet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_DOMStringMap_Get
//go:noescape
func TryDOMStringMapGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMStringMap_Set
//go:noescape
func HasDOMStringMapSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMStringMap_Set
//go:noescape
func DOMStringMapSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMStringMap_Set
//go:noescape
func CallDOMStringMapSet(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref)

//go:wasmimport plat/js/web try_DOMStringMap_Set
//go:noescape
func TryDOMStringMapSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMStringMap_Delete
//go:noescape
func HasDOMStringMapDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMStringMap_Delete
//go:noescape
func DOMStringMapDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMStringMap_Delete
//go:noescape
func CallDOMStringMapDelete(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_DOMStringMap_Delete
//go:noescape
func TryDOMStringMapDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGElement_ClassName
//go:noescape
func GetSVGElementClassName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGElement_OwnerSVGElement
//go:noescape
func GetSVGElementOwnerSVGElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGElement_ViewportElement
//go:noescape
func GetSVGElementViewportElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGElement_Style
//go:noescape
func GetSVGElementStyle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGElement_AttributeStyleMap
//go:noescape
func GetSVGElementAttributeStyleMap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGElement_CorrespondingElement
//go:noescape
func GetSVGElementCorrespondingElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGElement_CorrespondingUseElement
//go:noescape
func GetSVGElementCorrespondingUseElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGElement_Dataset
//go:noescape
func GetSVGElementDataset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGElement_Nonce
//go:noescape
func GetSVGElementNonce(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGElement_Nonce
//go:noescape
func SetSVGElementNonce(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_Autofocus
//go:noescape
func GetSVGElementAutofocus(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGElement_Autofocus
//go:noescape
func SetSVGElementAutofocus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_TabIndex
//go:noescape
func GetSVGElementTabIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGElement_TabIndex
//go:noescape
func SetSVGElementTabIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web has_SVGElement_Focus
//go:noescape
func HasSVGElementFocus(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGElement_Focus
//go:noescape
func SVGElementFocusFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGElement_Focus
//go:noescape
func CallSVGElementFocus(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGElement_Focus
//go:noescape
func TrySVGElementFocus(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGElement_Focus1
//go:noescape
func HasSVGElementFocus1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGElement_Focus1
//go:noescape
func SVGElementFocus1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGElement_Focus1
//go:noescape
func CallSVGElementFocus1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGElement_Focus1
//go:noescape
func TrySVGElementFocus1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGElement_Blur
//go:noescape
func HasSVGElementBlur(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGElement_Blur
//go:noescape
func SVGElementBlurFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGElement_Blur
//go:noescape
func CallSVGElementBlur(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGElement_Blur
//go:noescape
func TrySVGElementBlur(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGNumber_Value
//go:noescape
func GetSVGNumberValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGNumber_Value
//go:noescape
func SetSVGNumberValue(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAngle_UnitType
//go:noescape
func GetSVGAngleUnitType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAngle_Value
//go:noescape
func GetSVGAngleValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAngle_Value
//go:noescape
func SetSVGAngleValue(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAngle_ValueInSpecifiedUnits
//go:noescape
func GetSVGAngleValueInSpecifiedUnits(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAngle_ValueInSpecifiedUnits
//go:noescape
func SetSVGAngleValueInSpecifiedUnits(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAngle_ValueAsString
//go:noescape
func GetSVGAngleValueAsString(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGAngle_ValueAsString
//go:noescape
func SetSVGAngleValueAsString(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_SVGAngle_NewValueSpecifiedUnits
//go:noescape
func HasSVGAngleNewValueSpecifiedUnits(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGAngle_NewValueSpecifiedUnits
//go:noescape
func SVGAngleNewValueSpecifiedUnitsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGAngle_NewValueSpecifiedUnits
//go:noescape
func CallSVGAngleNewValueSpecifiedUnits(
	this js.Ref, retPtr unsafe.Pointer,
	unitType uint32,
	valueInSpecifiedUnits float32)

//go:wasmimport plat/js/web try_SVGAngle_NewValueSpecifiedUnits
//go:noescape
func TrySVGAngleNewValueSpecifiedUnits(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	unitType uint32,
	valueInSpecifiedUnits float32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGAngle_ConvertToSpecifiedUnits
//go:noescape
func HasSVGAngleConvertToSpecifiedUnits(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGAngle_ConvertToSpecifiedUnits
//go:noescape
func SVGAngleConvertToSpecifiedUnitsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGAngle_ConvertToSpecifiedUnits
//go:noescape
func CallSVGAngleConvertToSpecifiedUnits(
	this js.Ref, retPtr unsafe.Pointer,
	unitType uint32)

//go:wasmimport plat/js/web try_SVGAngle_ConvertToSpecifiedUnits
//go:noescape
func TrySVGAngleConvertToSpecifiedUnits(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	unitType uint32) (ok js.Ref)

//go:wasmimport plat/js/web store_DOMMatrixInit
//go:noescape
func DOMMatrixInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DOMMatrixInit
//go:noescape
func DOMMatrixInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_DOMMatrix_DOMMatrix
//go:noescape
func NewDOMMatrixByDOMMatrix(
	init js.Ref) js.Ref

//go:wasmimport plat/js/web new_DOMMatrix_DOMMatrix1
//go:noescape
func NewDOMMatrixByDOMMatrix1() js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_A
//go:noescape
func GetDOMMatrixA(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_A
//go:noescape
func SetDOMMatrixA(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_B
//go:noescape
func GetDOMMatrixB(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_B
//go:noescape
func SetDOMMatrixB(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_C
//go:noescape
func GetDOMMatrixC(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_C
//go:noescape
func SetDOMMatrixC(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_D
//go:noescape
func GetDOMMatrixD(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_D
//go:noescape
func SetDOMMatrixD(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_E
//go:noescape
func GetDOMMatrixE(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_E
//go:noescape
func SetDOMMatrixE(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_F
//go:noescape
func GetDOMMatrixF(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_F
//go:noescape
func SetDOMMatrixF(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M11
//go:noescape
func GetDOMMatrixM11(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M11
//go:noescape
func SetDOMMatrixM11(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M12
//go:noescape
func GetDOMMatrixM12(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M12
//go:noescape
func SetDOMMatrixM12(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M13
//go:noescape
func GetDOMMatrixM13(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M13
//go:noescape
func SetDOMMatrixM13(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M14
//go:noescape
func GetDOMMatrixM14(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M14
//go:noescape
func SetDOMMatrixM14(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M21
//go:noescape
func GetDOMMatrixM21(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M21
//go:noescape
func SetDOMMatrixM21(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M22
//go:noescape
func GetDOMMatrixM22(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M22
//go:noescape
func SetDOMMatrixM22(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M23
//go:noescape
func GetDOMMatrixM23(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M23
//go:noescape
func SetDOMMatrixM23(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M24
//go:noescape
func GetDOMMatrixM24(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M24
//go:noescape
func SetDOMMatrixM24(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M31
//go:noescape
func GetDOMMatrixM31(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M31
//go:noescape
func SetDOMMatrixM31(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M32
//go:noescape
func GetDOMMatrixM32(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M32
//go:noescape
func SetDOMMatrixM32(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M33
//go:noescape
func GetDOMMatrixM33(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M33
//go:noescape
func SetDOMMatrixM33(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M34
//go:noescape
func GetDOMMatrixM34(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M34
//go:noescape
func SetDOMMatrixM34(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M41
//go:noescape
func GetDOMMatrixM41(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M41
//go:noescape
func SetDOMMatrixM41(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M42
//go:noescape
func GetDOMMatrixM42(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M42
//go:noescape
func SetDOMMatrixM42(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M43
//go:noescape
func GetDOMMatrixM43(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M43
//go:noescape
func SetDOMMatrixM43(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M44
//go:noescape
func GetDOMMatrixM44(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_DOMMatrix_M44
//go:noescape
func SetDOMMatrixM44(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web has_DOMMatrix_FromMatrix
//go:noescape
func HasDOMMatrixFromMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromMatrix
//go:noescape
func DOMMatrixFromMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_FromMatrix
//go:noescape
func CallDOMMatrixFromMatrix(
	this js.Ref, retPtr unsafe.Pointer,
	other unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_FromMatrix
//go:noescape
func TryDOMMatrixFromMatrix(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	other unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_FromMatrix1
//go:noescape
func HasDOMMatrixFromMatrix1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromMatrix1
//go:noescape
func DOMMatrixFromMatrix1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_FromMatrix1
//go:noescape
func CallDOMMatrixFromMatrix1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_FromMatrix1
//go:noescape
func TryDOMMatrixFromMatrix1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_FromFloat32Array
//go:noescape
func HasDOMMatrixFromFloat32Array(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromFloat32Array
//go:noescape
func DOMMatrixFromFloat32ArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_FromFloat32Array
//go:noescape
func CallDOMMatrixFromFloat32Array(
	this js.Ref, retPtr unsafe.Pointer,
	array32 js.Ref)

//go:wasmimport plat/js/web try_DOMMatrix_FromFloat32Array
//go:noescape
func TryDOMMatrixFromFloat32Array(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	array32 js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_FromFloat64Array
//go:noescape
func HasDOMMatrixFromFloat64Array(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromFloat64Array
//go:noescape
func DOMMatrixFromFloat64ArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_FromFloat64Array
//go:noescape
func CallDOMMatrixFromFloat64Array(
	this js.Ref, retPtr unsafe.Pointer,
	array64 js.Ref)

//go:wasmimport plat/js/web try_DOMMatrix_FromFloat64Array
//go:noescape
func TryDOMMatrixFromFloat64Array(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	array64 js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_MultiplySelf
//go:noescape
func HasDOMMatrixMultiplySelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_MultiplySelf
//go:noescape
func DOMMatrixMultiplySelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_MultiplySelf
//go:noescape
func CallDOMMatrixMultiplySelf(
	this js.Ref, retPtr unsafe.Pointer,
	other unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_MultiplySelf
//go:noescape
func TryDOMMatrixMultiplySelf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	other unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_MultiplySelf1
//go:noescape
func HasDOMMatrixMultiplySelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_MultiplySelf1
//go:noescape
func DOMMatrixMultiplySelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_MultiplySelf1
//go:noescape
func CallDOMMatrixMultiplySelf1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_MultiplySelf1
//go:noescape
func TryDOMMatrixMultiplySelf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_PreMultiplySelf
//go:noescape
func HasDOMMatrixPreMultiplySelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_PreMultiplySelf
//go:noescape
func DOMMatrixPreMultiplySelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_PreMultiplySelf
//go:noescape
func CallDOMMatrixPreMultiplySelf(
	this js.Ref, retPtr unsafe.Pointer,
	other unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_PreMultiplySelf
//go:noescape
func TryDOMMatrixPreMultiplySelf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	other unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_PreMultiplySelf1
//go:noescape
func HasDOMMatrixPreMultiplySelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_PreMultiplySelf1
//go:noescape
func DOMMatrixPreMultiplySelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_PreMultiplySelf1
//go:noescape
func CallDOMMatrixPreMultiplySelf1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_PreMultiplySelf1
//go:noescape
func TryDOMMatrixPreMultiplySelf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_TranslateSelf
//go:noescape
func HasDOMMatrixTranslateSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf
//go:noescape
func DOMMatrixTranslateSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_TranslateSelf
//go:noescape
func CallDOMMatrixTranslateSelf(
	this js.Ref, retPtr unsafe.Pointer,
	tx float64,
	ty float64,
	tz float64)

//go:wasmimport plat/js/web try_DOMMatrix_TranslateSelf
//go:noescape
func TryDOMMatrixTranslateSelf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tx float64,
	ty float64,
	tz float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_TranslateSelf1
//go:noescape
func HasDOMMatrixTranslateSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf1
//go:noescape
func DOMMatrixTranslateSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_TranslateSelf1
//go:noescape
func CallDOMMatrixTranslateSelf1(
	this js.Ref, retPtr unsafe.Pointer,
	tx float64,
	ty float64)

//go:wasmimport plat/js/web try_DOMMatrix_TranslateSelf1
//go:noescape
func TryDOMMatrixTranslateSelf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tx float64,
	ty float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_TranslateSelf2
//go:noescape
func HasDOMMatrixTranslateSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf2
//go:noescape
func DOMMatrixTranslateSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_TranslateSelf2
//go:noescape
func CallDOMMatrixTranslateSelf2(
	this js.Ref, retPtr unsafe.Pointer,
	tx float64)

//go:wasmimport plat/js/web try_DOMMatrix_TranslateSelf2
//go:noescape
func TryDOMMatrixTranslateSelf2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tx float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_TranslateSelf3
//go:noescape
func HasDOMMatrixTranslateSelf3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf3
//go:noescape
func DOMMatrixTranslateSelf3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_TranslateSelf3
//go:noescape
func CallDOMMatrixTranslateSelf3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_TranslateSelf3
//go:noescape
func TryDOMMatrixTranslateSelf3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_ScaleSelf
//go:noescape
func HasDOMMatrixScaleSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf
//go:noescape
func DOMMatrixScaleSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf
//go:noescape
func CallDOMMatrixScaleSelf(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64,
	originY float64,
	originZ float64)

//go:wasmimport plat/js/web try_DOMMatrix_ScaleSelf
//go:noescape
func TryDOMMatrixScaleSelf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64,
	originY float64,
	originZ float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_ScaleSelf1
//go:noescape
func HasDOMMatrixScaleSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf1
//go:noescape
func DOMMatrixScaleSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf1
//go:noescape
func CallDOMMatrixScaleSelf1(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64,
	originY float64)

//go:wasmimport plat/js/web try_DOMMatrix_ScaleSelf1
//go:noescape
func TryDOMMatrixScaleSelf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64,
	originY float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_ScaleSelf2
//go:noescape
func HasDOMMatrixScaleSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf2
//go:noescape
func DOMMatrixScaleSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf2
//go:noescape
func CallDOMMatrixScaleSelf2(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64)

//go:wasmimport plat/js/web try_DOMMatrix_ScaleSelf2
//go:noescape
func TryDOMMatrixScaleSelf2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_ScaleSelf3
//go:noescape
func HasDOMMatrixScaleSelf3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf3
//go:noescape
func DOMMatrixScaleSelf3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf3
//go:noescape
func CallDOMMatrixScaleSelf3(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64)

//go:wasmimport plat/js/web try_DOMMatrix_ScaleSelf3
//go:noescape
func TryDOMMatrixScaleSelf3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64,
	scaleZ float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_ScaleSelf4
//go:noescape
func HasDOMMatrixScaleSelf4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf4
//go:noescape
func DOMMatrixScaleSelf4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf4
//go:noescape
func CallDOMMatrixScaleSelf4(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64)

//go:wasmimport plat/js/web try_DOMMatrix_ScaleSelf4
//go:noescape
func TryDOMMatrixScaleSelf4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64,
	scaleY float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_ScaleSelf5
//go:noescape
func HasDOMMatrixScaleSelf5(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf5
//go:noescape
func DOMMatrixScaleSelf5Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf5
//go:noescape
func CallDOMMatrixScaleSelf5(
	this js.Ref, retPtr unsafe.Pointer,
	scaleX float64)

//go:wasmimport plat/js/web try_DOMMatrix_ScaleSelf5
//go:noescape
func TryDOMMatrixScaleSelf5(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scaleX float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_ScaleSelf6
//go:noescape
func HasDOMMatrixScaleSelf6(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf6
//go:noescape
func DOMMatrixScaleSelf6Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf6
//go:noescape
func CallDOMMatrixScaleSelf6(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_ScaleSelf6
//go:noescape
func TryDOMMatrixScaleSelf6(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_Scale3dSelf
//go:noescape
func HasDOMMatrixScale3dSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf
//go:noescape
func DOMMatrixScale3dSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_Scale3dSelf
//go:noescape
func CallDOMMatrixScale3dSelf(
	this js.Ref, retPtr unsafe.Pointer,
	scale float64,
	originX float64,
	originY float64,
	originZ float64)

//go:wasmimport plat/js/web try_DOMMatrix_Scale3dSelf
//go:noescape
func TryDOMMatrixScale3dSelf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scale float64,
	originX float64,
	originY float64,
	originZ float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_Scale3dSelf1
//go:noescape
func HasDOMMatrixScale3dSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf1
//go:noescape
func DOMMatrixScale3dSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_Scale3dSelf1
//go:noescape
func CallDOMMatrixScale3dSelf1(
	this js.Ref, retPtr unsafe.Pointer,
	scale float64,
	originX float64,
	originY float64)

//go:wasmimport plat/js/web try_DOMMatrix_Scale3dSelf1
//go:noescape
func TryDOMMatrixScale3dSelf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scale float64,
	originX float64,
	originY float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_Scale3dSelf2
//go:noescape
func HasDOMMatrixScale3dSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf2
//go:noescape
func DOMMatrixScale3dSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_Scale3dSelf2
//go:noescape
func CallDOMMatrixScale3dSelf2(
	this js.Ref, retPtr unsafe.Pointer,
	scale float64,
	originX float64)

//go:wasmimport plat/js/web try_DOMMatrix_Scale3dSelf2
//go:noescape
func TryDOMMatrixScale3dSelf2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scale float64,
	originX float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_Scale3dSelf3
//go:noescape
func HasDOMMatrixScale3dSelf3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf3
//go:noescape
func DOMMatrixScale3dSelf3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_Scale3dSelf3
//go:noescape
func CallDOMMatrixScale3dSelf3(
	this js.Ref, retPtr unsafe.Pointer,
	scale float64)

//go:wasmimport plat/js/web try_DOMMatrix_Scale3dSelf3
//go:noescape
func TryDOMMatrixScale3dSelf3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	scale float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_Scale3dSelf4
//go:noescape
func HasDOMMatrixScale3dSelf4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf4
//go:noescape
func DOMMatrixScale3dSelf4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_Scale3dSelf4
//go:noescape
func CallDOMMatrixScale3dSelf4(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_Scale3dSelf4
//go:noescape
func TryDOMMatrixScale3dSelf4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateSelf
//go:noescape
func HasDOMMatrixRotateSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf
//go:noescape
func DOMMatrixRotateSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateSelf
//go:noescape
func CallDOMMatrixRotateSelf(
	this js.Ref, retPtr unsafe.Pointer,
	rotX float64,
	rotY float64,
	rotZ float64)

//go:wasmimport plat/js/web try_DOMMatrix_RotateSelf
//go:noescape
func TryDOMMatrixRotateSelf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rotX float64,
	rotY float64,
	rotZ float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateSelf1
//go:noescape
func HasDOMMatrixRotateSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf1
//go:noescape
func DOMMatrixRotateSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateSelf1
//go:noescape
func CallDOMMatrixRotateSelf1(
	this js.Ref, retPtr unsafe.Pointer,
	rotX float64,
	rotY float64)

//go:wasmimport plat/js/web try_DOMMatrix_RotateSelf1
//go:noescape
func TryDOMMatrixRotateSelf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rotX float64,
	rotY float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateSelf2
//go:noescape
func HasDOMMatrixRotateSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf2
//go:noescape
func DOMMatrixRotateSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateSelf2
//go:noescape
func CallDOMMatrixRotateSelf2(
	this js.Ref, retPtr unsafe.Pointer,
	rotX float64)

//go:wasmimport plat/js/web try_DOMMatrix_RotateSelf2
//go:noescape
func TryDOMMatrixRotateSelf2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rotX float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateSelf3
//go:noescape
func HasDOMMatrixRotateSelf3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf3
//go:noescape
func DOMMatrixRotateSelf3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateSelf3
//go:noescape
func CallDOMMatrixRotateSelf3(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_RotateSelf3
//go:noescape
func TryDOMMatrixRotateSelf3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateFromVectorSelf
//go:noescape
func HasDOMMatrixRotateFromVectorSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateFromVectorSelf
//go:noescape
func DOMMatrixRotateFromVectorSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateFromVectorSelf
//go:noescape
func CallDOMMatrixRotateFromVectorSelf(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_DOMMatrix_RotateFromVectorSelf
//go:noescape
func TryDOMMatrixRotateFromVectorSelf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateFromVectorSelf1
//go:noescape
func HasDOMMatrixRotateFromVectorSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateFromVectorSelf1
//go:noescape
func DOMMatrixRotateFromVectorSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateFromVectorSelf1
//go:noescape
func CallDOMMatrixRotateFromVectorSelf1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64)

//go:wasmimport plat/js/web try_DOMMatrix_RotateFromVectorSelf1
//go:noescape
func TryDOMMatrixRotateFromVectorSelf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateFromVectorSelf2
//go:noescape
func HasDOMMatrixRotateFromVectorSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateFromVectorSelf2
//go:noescape
func DOMMatrixRotateFromVectorSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateFromVectorSelf2
//go:noescape
func CallDOMMatrixRotateFromVectorSelf2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_RotateFromVectorSelf2
//go:noescape
func TryDOMMatrixRotateFromVectorSelf2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateAxisAngleSelf
//go:noescape
func HasDOMMatrixRotateAxisAngleSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf
//go:noescape
func DOMMatrixRotateAxisAngleSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateAxisAngleSelf
//go:noescape
func CallDOMMatrixRotateAxisAngleSelf(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	z float64,
	angle float64)

//go:wasmimport plat/js/web try_DOMMatrix_RotateAxisAngleSelf
//go:noescape
func TryDOMMatrixRotateAxisAngleSelf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	z float64,
	angle float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateAxisAngleSelf1
//go:noescape
func HasDOMMatrixRotateAxisAngleSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf1
//go:noescape
func DOMMatrixRotateAxisAngleSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateAxisAngleSelf1
//go:noescape
func CallDOMMatrixRotateAxisAngleSelf1(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64,
	z float64)

//go:wasmimport plat/js/web try_DOMMatrix_RotateAxisAngleSelf1
//go:noescape
func TryDOMMatrixRotateAxisAngleSelf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64,
	z float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateAxisAngleSelf2
//go:noescape
func HasDOMMatrixRotateAxisAngleSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf2
//go:noescape
func DOMMatrixRotateAxisAngleSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateAxisAngleSelf2
//go:noescape
func CallDOMMatrixRotateAxisAngleSelf2(
	this js.Ref, retPtr unsafe.Pointer,
	x float64,
	y float64)

//go:wasmimport plat/js/web try_DOMMatrix_RotateAxisAngleSelf2
//go:noescape
func TryDOMMatrixRotateAxisAngleSelf2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64,
	y float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateAxisAngleSelf3
//go:noescape
func HasDOMMatrixRotateAxisAngleSelf3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf3
//go:noescape
func DOMMatrixRotateAxisAngleSelf3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateAxisAngleSelf3
//go:noescape
func CallDOMMatrixRotateAxisAngleSelf3(
	this js.Ref, retPtr unsafe.Pointer,
	x float64)

//go:wasmimport plat/js/web try_DOMMatrix_RotateAxisAngleSelf3
//go:noescape
func TryDOMMatrixRotateAxisAngleSelf3(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	x float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_RotateAxisAngleSelf4
//go:noescape
func HasDOMMatrixRotateAxisAngleSelf4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf4
//go:noescape
func DOMMatrixRotateAxisAngleSelf4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateAxisAngleSelf4
//go:noescape
func CallDOMMatrixRotateAxisAngleSelf4(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_RotateAxisAngleSelf4
//go:noescape
func TryDOMMatrixRotateAxisAngleSelf4(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_SkewXSelf
//go:noescape
func HasDOMMatrixSkewXSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewXSelf
//go:noescape
func DOMMatrixSkewXSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_SkewXSelf
//go:noescape
func CallDOMMatrixSkewXSelf(
	this js.Ref, retPtr unsafe.Pointer,
	sx float64)

//go:wasmimport plat/js/web try_DOMMatrix_SkewXSelf
//go:noescape
func TryDOMMatrixSkewXSelf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sx float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_SkewXSelf1
//go:noescape
func HasDOMMatrixSkewXSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewXSelf1
//go:noescape
func DOMMatrixSkewXSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_SkewXSelf1
//go:noescape
func CallDOMMatrixSkewXSelf1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_SkewXSelf1
//go:noescape
func TryDOMMatrixSkewXSelf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_SkewYSelf
//go:noescape
func HasDOMMatrixSkewYSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewYSelf
//go:noescape
func DOMMatrixSkewYSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_SkewYSelf
//go:noescape
func CallDOMMatrixSkewYSelf(
	this js.Ref, retPtr unsafe.Pointer,
	sy float64)

//go:wasmimport plat/js/web try_DOMMatrix_SkewYSelf
//go:noescape
func TryDOMMatrixSkewYSelf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sy float64) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_SkewYSelf1
//go:noescape
func HasDOMMatrixSkewYSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewYSelf1
//go:noescape
func DOMMatrixSkewYSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_SkewYSelf1
//go:noescape
func CallDOMMatrixSkewYSelf1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_SkewYSelf1
//go:noescape
func TryDOMMatrixSkewYSelf1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_InvertSelf
//go:noescape
func HasDOMMatrixInvertSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_InvertSelf
//go:noescape
func DOMMatrixInvertSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_InvertSelf
//go:noescape
func CallDOMMatrixInvertSelf(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMMatrix_InvertSelf
//go:noescape
func TryDOMMatrixInvertSelf(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMMatrix_SetMatrixValue
//go:noescape
func HasDOMMatrixSetMatrixValue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SetMatrixValue
//go:noescape
func DOMMatrixSetMatrixValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_SetMatrixValue
//go:noescape
func CallDOMMatrixSetMatrixValue(
	this js.Ref, retPtr unsafe.Pointer,
	transformList js.Ref)

//go:wasmimport plat/js/web try_DOMMatrix_SetMatrixValue
//go:noescape
func TryDOMMatrixSetMatrixValue(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	transformList js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web store_DOMMatrix2DInit
//go:noescape
func DOMMatrix2DInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_DOMMatrix2DInit
//go:noescape
func DOMMatrix2DInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGTransform_Type
//go:noescape
func GetSVGTransformType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTransform_Matrix
//go:noescape
func GetSVGTransformMatrix(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGTransform_Angle
//go:noescape
func GetSVGTransformAngle(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransform_SetMatrix
//go:noescape
func HasSVGTransformSetMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetMatrix
//go:noescape
func SVGTransformSetMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetMatrix
//go:noescape
func CallSVGTransformSetMatrix(
	this js.Ref, retPtr unsafe.Pointer,
	matrix unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGTransform_SetMatrix
//go:noescape
func TrySVGTransformSetMatrix(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	matrix unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransform_SetMatrix1
//go:noescape
func HasSVGTransformSetMatrix1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetMatrix1
//go:noescape
func SVGTransformSetMatrix1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetMatrix1
//go:noescape
func CallSVGTransformSetMatrix1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGTransform_SetMatrix1
//go:noescape
func TrySVGTransformSetMatrix1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransform_SetTranslate
//go:noescape
func HasSVGTransformSetTranslate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetTranslate
//go:noescape
func SVGTransformSetTranslateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetTranslate
//go:noescape
func CallSVGTransformSetTranslate(
	this js.Ref, retPtr unsafe.Pointer,
	tx float32,
	ty float32)

//go:wasmimport plat/js/web try_SVGTransform_SetTranslate
//go:noescape
func TrySVGTransformSetTranslate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	tx float32,
	ty float32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransform_SetScale
//go:noescape
func HasSVGTransformSetScale(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetScale
//go:noescape
func SVGTransformSetScaleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetScale
//go:noescape
func CallSVGTransformSetScale(
	this js.Ref, retPtr unsafe.Pointer,
	sx float32,
	sy float32)

//go:wasmimport plat/js/web try_SVGTransform_SetScale
//go:noescape
func TrySVGTransformSetScale(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	sx float32,
	sy float32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransform_SetRotate
//go:noescape
func HasSVGTransformSetRotate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetRotate
//go:noescape
func SVGTransformSetRotateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetRotate
//go:noescape
func CallSVGTransformSetRotate(
	this js.Ref, retPtr unsafe.Pointer,
	angle float32,
	cx float32,
	cy float32)

//go:wasmimport plat/js/web try_SVGTransform_SetRotate
//go:noescape
func TrySVGTransformSetRotate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	angle float32,
	cx float32,
	cy float32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransform_SetSkewX
//go:noescape
func HasSVGTransformSetSkewX(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetSkewX
//go:noescape
func SVGTransformSetSkewXFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetSkewX
//go:noescape
func CallSVGTransformSetSkewX(
	this js.Ref, retPtr unsafe.Pointer,
	angle float32)

//go:wasmimport plat/js/web try_SVGTransform_SetSkewX
//go:noescape
func TrySVGTransformSetSkewX(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	angle float32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGTransform_SetSkewY
//go:noescape
func HasSVGTransformSetSkewY(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetSkewY
//go:noescape
func SVGTransformSetSkewYFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetSkewY
//go:noescape
func CallSVGTransformSetSkewY(
	this js.Ref, retPtr unsafe.Pointer,
	angle float32)

//go:wasmimport plat/js/web try_SVGTransform_SetSkewY
//go:noescape
func TrySVGTransformSetSkewY(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	angle float32) (ok js.Ref)

//go:wasmimport plat/js/web new_DOMPointReadOnly_DOMPointReadOnly
//go:noescape
func NewDOMPointReadOnlyByDOMPointReadOnly(
	x float64,
	y float64,
	z float64,
	w float64) js.Ref

//go:wasmimport plat/js/web new_DOMPointReadOnly_DOMPointReadOnly1
//go:noescape
func NewDOMPointReadOnlyByDOMPointReadOnly1(
	x float64,
	y float64,
	z float64) js.Ref

//go:wasmimport plat/js/web new_DOMPointReadOnly_DOMPointReadOnly2
//go:noescape
func NewDOMPointReadOnlyByDOMPointReadOnly2(
	x float64,
	y float64) js.Ref

//go:wasmimport plat/js/web new_DOMPointReadOnly_DOMPointReadOnly3
//go:noescape
func NewDOMPointReadOnlyByDOMPointReadOnly3(
	x float64) js.Ref

//go:wasmimport plat/js/web new_DOMPointReadOnly_DOMPointReadOnly4
//go:noescape
func NewDOMPointReadOnlyByDOMPointReadOnly4() js.Ref

//go:wasmimport plat/js/web get_DOMPointReadOnly_X
//go:noescape
func GetDOMPointReadOnlyX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMPointReadOnly_Y
//go:noescape
func GetDOMPointReadOnlyY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMPointReadOnly_Z
//go:noescape
func GetDOMPointReadOnlyZ(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_DOMPointReadOnly_W
//go:noescape
func GetDOMPointReadOnlyW(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMPointReadOnly_FromPoint
//go:noescape
func HasDOMPointReadOnlyFromPoint(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_FromPoint
//go:noescape
func DOMPointReadOnlyFromPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMPointReadOnly_FromPoint
//go:noescape
func CallDOMPointReadOnlyFromPoint(
	this js.Ref, retPtr unsafe.Pointer,
	other unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMPointReadOnly_FromPoint
//go:noescape
func TryDOMPointReadOnlyFromPoint(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	other unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMPointReadOnly_FromPoint1
//go:noescape
func HasDOMPointReadOnlyFromPoint1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_FromPoint1
//go:noescape
func DOMPointReadOnlyFromPoint1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMPointReadOnly_FromPoint1
//go:noescape
func CallDOMPointReadOnlyFromPoint1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMPointReadOnly_FromPoint1
//go:noescape
func TryDOMPointReadOnlyFromPoint1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMPointReadOnly_MatrixTransform
//go:noescape
func HasDOMPointReadOnlyMatrixTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_MatrixTransform
//go:noescape
func DOMPointReadOnlyMatrixTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMPointReadOnly_MatrixTransform
//go:noescape
func CallDOMPointReadOnlyMatrixTransform(
	this js.Ref, retPtr unsafe.Pointer,
	matrix unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMPointReadOnly_MatrixTransform
//go:noescape
func TryDOMPointReadOnlyMatrixTransform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	matrix unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMPointReadOnly_MatrixTransform1
//go:noescape
func HasDOMPointReadOnlyMatrixTransform1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_MatrixTransform1
//go:noescape
func DOMPointReadOnlyMatrixTransform1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMPointReadOnly_MatrixTransform1
//go:noescape
func CallDOMPointReadOnlyMatrixTransform1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMPointReadOnly_MatrixTransform1
//go:noescape
func TryDOMPointReadOnlyMatrixTransform1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_DOMPointReadOnly_ToJSON
//go:noescape
func HasDOMPointReadOnlyToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_ToJSON
//go:noescape
func DOMPointReadOnlyToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMPointReadOnly_ToJSON
//go:noescape
func CallDOMPointReadOnlyToJSON(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_DOMPointReadOnly_ToJSON
//go:noescape
func TryDOMPointReadOnlyToJSON(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedRect_BaseVal
//go:noescape
func GetSVGAnimatedRectBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedRect_AnimVal
//go:noescape
func GetSVGAnimatedRectAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGPreserveAspectRatio_Align
//go:noescape
func GetSVGPreserveAspectRatioAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGPreserveAspectRatio_Align
//go:noescape
func SetSVGPreserveAspectRatioAlign(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_SVGPreserveAspectRatio_MeetOrSlice
//go:noescape
func GetSVGPreserveAspectRatioMeetOrSlice(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGPreserveAspectRatio_MeetOrSlice
//go:noescape
func SetSVGPreserveAspectRatioMeetOrSlice(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedPreserveAspectRatio_BaseVal
//go:noescape
func GetSVGAnimatedPreserveAspectRatioBaseVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGAnimatedPreserveAspectRatio_AnimVal
//go:noescape
func GetSVGAnimatedPreserveAspectRatioAnimVal(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGSVGElement_X
//go:noescape
func GetSVGSVGElementX(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGSVGElement_Y
//go:noescape
func GetSVGSVGElementY(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGSVGElement_Width
//go:noescape
func GetSVGSVGElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGSVGElement_Height
//go:noescape
func GetSVGSVGElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGSVGElement_CurrentScale
//go:noescape
func GetSVGSVGElementCurrentScale(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_SVGSVGElement_CurrentScale
//go:noescape
func SetSVGSVGElementCurrentScale(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGSVGElement_CurrentTranslate
//go:noescape
func GetSVGSVGElementCurrentTranslate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGSVGElement_ViewBox
//go:noescape
func GetSVGSVGElementViewBox(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_SVGSVGElement_PreserveAspectRatio
//go:noescape
func GetSVGSVGElementPreserveAspectRatio(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_GetIntersectionList
//go:noescape
func HasSVGSVGElementGetIntersectionList(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_GetIntersectionList
//go:noescape
func SVGSVGElementGetIntersectionListFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_GetIntersectionList
//go:noescape
func CallSVGSVGElementGetIntersectionList(
	this js.Ref, retPtr unsafe.Pointer,
	rect js.Ref,
	referenceElement js.Ref)

//go:wasmimport plat/js/web try_SVGSVGElement_GetIntersectionList
//go:noescape
func TrySVGSVGElementGetIntersectionList(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rect js.Ref,
	referenceElement js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_GetEnclosureList
//go:noescape
func HasSVGSVGElementGetEnclosureList(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_GetEnclosureList
//go:noescape
func SVGSVGElementGetEnclosureListFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_GetEnclosureList
//go:noescape
func CallSVGSVGElementGetEnclosureList(
	this js.Ref, retPtr unsafe.Pointer,
	rect js.Ref,
	referenceElement js.Ref)

//go:wasmimport plat/js/web try_SVGSVGElement_GetEnclosureList
//go:noescape
func TrySVGSVGElementGetEnclosureList(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	rect js.Ref,
	referenceElement js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_CheckIntersection
//go:noescape
func HasSVGSVGElementCheckIntersection(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CheckIntersection
//go:noescape
func SVGSVGElementCheckIntersectionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CheckIntersection
//go:noescape
func CallSVGSVGElementCheckIntersection(
	this js.Ref, retPtr unsafe.Pointer,
	element js.Ref,
	rect js.Ref)

//go:wasmimport plat/js/web try_SVGSVGElement_CheckIntersection
//go:noescape
func TrySVGSVGElementCheckIntersection(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref,
	rect js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_CheckEnclosure
//go:noescape
func HasSVGSVGElementCheckEnclosure(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CheckEnclosure
//go:noescape
func SVGSVGElementCheckEnclosureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CheckEnclosure
//go:noescape
func CallSVGSVGElementCheckEnclosure(
	this js.Ref, retPtr unsafe.Pointer,
	element js.Ref,
	rect js.Ref)

//go:wasmimport plat/js/web try_SVGSVGElement_CheckEnclosure
//go:noescape
func TrySVGSVGElementCheckEnclosure(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref,
	rect js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_DeselectAll
//go:noescape
func HasSVGSVGElementDeselectAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_DeselectAll
//go:noescape
func SVGSVGElementDeselectAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_DeselectAll
//go:noescape
func CallSVGSVGElementDeselectAll(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_DeselectAll
//go:noescape
func TrySVGSVGElementDeselectAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_CreateSVGNumber
//go:noescape
func HasSVGSVGElementCreateSVGNumber(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGNumber
//go:noescape
func SVGSVGElementCreateSVGNumberFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGNumber
//go:noescape
func CallSVGSVGElementCreateSVGNumber(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_CreateSVGNumber
//go:noescape
func TrySVGSVGElementCreateSVGNumber(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_CreateSVGLength
//go:noescape
func HasSVGSVGElementCreateSVGLength(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGLength
//go:noescape
func SVGSVGElementCreateSVGLengthFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGLength
//go:noescape
func CallSVGSVGElementCreateSVGLength(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_CreateSVGLength
//go:noescape
func TrySVGSVGElementCreateSVGLength(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_CreateSVGAngle
//go:noescape
func HasSVGSVGElementCreateSVGAngle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGAngle
//go:noescape
func SVGSVGElementCreateSVGAngleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGAngle
//go:noescape
func CallSVGSVGElementCreateSVGAngle(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_CreateSVGAngle
//go:noescape
func TrySVGSVGElementCreateSVGAngle(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_CreateSVGPoint
//go:noescape
func HasSVGSVGElementCreateSVGPoint(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGPoint
//go:noescape
func SVGSVGElementCreateSVGPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGPoint
//go:noescape
func CallSVGSVGElementCreateSVGPoint(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_CreateSVGPoint
//go:noescape
func TrySVGSVGElementCreateSVGPoint(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_CreateSVGMatrix
//go:noescape
func HasSVGSVGElementCreateSVGMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGMatrix
//go:noescape
func SVGSVGElementCreateSVGMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGMatrix
//go:noescape
func CallSVGSVGElementCreateSVGMatrix(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_CreateSVGMatrix
//go:noescape
func TrySVGSVGElementCreateSVGMatrix(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_CreateSVGRect
//go:noescape
func HasSVGSVGElementCreateSVGRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGRect
//go:noescape
func SVGSVGElementCreateSVGRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGRect
//go:noescape
func CallSVGSVGElementCreateSVGRect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_CreateSVGRect
//go:noescape
func TrySVGSVGElementCreateSVGRect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_CreateSVGTransform
//go:noescape
func HasSVGSVGElementCreateSVGTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGTransform
//go:noescape
func SVGSVGElementCreateSVGTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGTransform
//go:noescape
func CallSVGSVGElementCreateSVGTransform(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_CreateSVGTransform
//go:noescape
func TrySVGSVGElementCreateSVGTransform(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_CreateSVGTransformFromMatrix
//go:noescape
func HasSVGSVGElementCreateSVGTransformFromMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGTransformFromMatrix
//go:noescape
func SVGSVGElementCreateSVGTransformFromMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGTransformFromMatrix
//go:noescape
func CallSVGSVGElementCreateSVGTransformFromMatrix(
	this js.Ref, retPtr unsafe.Pointer,
	matrix unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_CreateSVGTransformFromMatrix
//go:noescape
func TrySVGSVGElementCreateSVGTransformFromMatrix(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	matrix unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_CreateSVGTransformFromMatrix1
//go:noescape
func HasSVGSVGElementCreateSVGTransformFromMatrix1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGTransformFromMatrix1
//go:noescape
func SVGSVGElementCreateSVGTransformFromMatrix1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGTransformFromMatrix1
//go:noescape
func CallSVGSVGElementCreateSVGTransformFromMatrix1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_CreateSVGTransformFromMatrix1
//go:noescape
func TrySVGSVGElementCreateSVGTransformFromMatrix1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_GetElementById
//go:noescape
func HasSVGSVGElementGetElementById(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_GetElementById
//go:noescape
func SVGSVGElementGetElementByIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_GetElementById
//go:noescape
func CallSVGSVGElementGetElementById(
	this js.Ref, retPtr unsafe.Pointer,
	elementId js.Ref)

//go:wasmimport plat/js/web try_SVGSVGElement_GetElementById
//go:noescape
func TrySVGSVGElementGetElementById(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	elementId js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_SuspendRedraw
//go:noescape
func HasSVGSVGElementSuspendRedraw(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_SuspendRedraw
//go:noescape
func SVGSVGElementSuspendRedrawFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_SuspendRedraw
//go:noescape
func CallSVGSVGElementSuspendRedraw(
	this js.Ref, retPtr unsafe.Pointer,
	maxWaitMilliseconds uint32)

//go:wasmimport plat/js/web try_SVGSVGElement_SuspendRedraw
//go:noescape
func TrySVGSVGElementSuspendRedraw(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	maxWaitMilliseconds uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_UnsuspendRedraw
//go:noescape
func HasSVGSVGElementUnsuspendRedraw(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_UnsuspendRedraw
//go:noescape
func SVGSVGElementUnsuspendRedrawFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_UnsuspendRedraw
//go:noescape
func CallSVGSVGElementUnsuspendRedraw(
	this js.Ref, retPtr unsafe.Pointer,
	suspendHandleID uint32)

//go:wasmimport plat/js/web try_SVGSVGElement_UnsuspendRedraw
//go:noescape
func TrySVGSVGElementUnsuspendRedraw(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	suspendHandleID uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_UnsuspendRedrawAll
//go:noescape
func HasSVGSVGElementUnsuspendRedrawAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_UnsuspendRedrawAll
//go:noescape
func SVGSVGElementUnsuspendRedrawAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_UnsuspendRedrawAll
//go:noescape
func CallSVGSVGElementUnsuspendRedrawAll(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_UnsuspendRedrawAll
//go:noescape
func TrySVGSVGElementUnsuspendRedrawAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_ForceRedraw
//go:noescape
func HasSVGSVGElementForceRedraw(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_ForceRedraw
//go:noescape
func SVGSVGElementForceRedrawFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_ForceRedraw
//go:noescape
func CallSVGSVGElementForceRedraw(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_ForceRedraw
//go:noescape
func TrySVGSVGElementForceRedraw(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_PauseAnimations
//go:noescape
func HasSVGSVGElementPauseAnimations(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_PauseAnimations
//go:noescape
func SVGSVGElementPauseAnimationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_PauseAnimations
//go:noescape
func CallSVGSVGElementPauseAnimations(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_PauseAnimations
//go:noescape
func TrySVGSVGElementPauseAnimations(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_UnpauseAnimations
//go:noescape
func HasSVGSVGElementUnpauseAnimations(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_UnpauseAnimations
//go:noescape
func SVGSVGElementUnpauseAnimationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_UnpauseAnimations
//go:noescape
func CallSVGSVGElementUnpauseAnimations(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_UnpauseAnimations
//go:noescape
func TrySVGSVGElementUnpauseAnimations(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_AnimationsPaused
//go:noescape
func HasSVGSVGElementAnimationsPaused(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_AnimationsPaused
//go:noescape
func SVGSVGElementAnimationsPausedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_AnimationsPaused
//go:noescape
func CallSVGSVGElementAnimationsPaused(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_AnimationsPaused
//go:noescape
func TrySVGSVGElementAnimationsPaused(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_GetCurrentTime
//go:noescape
func HasSVGSVGElementGetCurrentTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_GetCurrentTime
//go:noescape
func SVGSVGElementGetCurrentTimeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_GetCurrentTime
//go:noescape
func CallSVGSVGElementGetCurrentTime(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_SVGSVGElement_GetCurrentTime
//go:noescape
func TrySVGSVGElementGetCurrentTime(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_SVGSVGElement_SetCurrentTime
//go:noescape
func HasSVGSVGElementSetCurrentTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_SetCurrentTime
//go:noescape
func SVGSVGElementSetCurrentTimeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_SetCurrentTime
//go:noescape
func CallSVGSVGElementSetCurrentTime(
	this js.Ref, retPtr unsafe.Pointer,
	seconds float32)

//go:wasmimport plat/js/web try_SVGSVGElement_SetCurrentTime
//go:noescape
func TrySVGSVGElementSetCurrentTime(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	seconds float32) (ok js.Ref)

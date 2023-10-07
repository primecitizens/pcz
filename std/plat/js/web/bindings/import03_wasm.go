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
func HasFuncSelectionGetRangeAt(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_GetRangeAt
//go:noescape
func FuncSelectionGetRangeAt(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionAddRange(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_AddRange
//go:noescape
func FuncSelectionAddRange(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionRemoveRange(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_RemoveRange
//go:noescape
func FuncSelectionRemoveRange(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionRemoveAllRanges(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_RemoveAllRanges
//go:noescape
func FuncSelectionRemoveAllRanges(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionEmpty(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Empty
//go:noescape
func FuncSelectionEmpty(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionGetComposedRanges(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_GetComposedRanges
//go:noescape
func FuncSelectionGetComposedRanges(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionCollapse(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Collapse
//go:noescape
func FuncSelectionCollapse(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionCollapse1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Collapse1
//go:noescape
func FuncSelectionCollapse1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionSetPosition(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_SetPosition
//go:noescape
func FuncSelectionSetPosition(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionSetPosition1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_SetPosition1
//go:noescape
func FuncSelectionSetPosition1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionCollapseToStart(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_CollapseToStart
//go:noescape
func FuncSelectionCollapseToStart(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionCollapseToEnd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_CollapseToEnd
//go:noescape
func FuncSelectionCollapseToEnd(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionExtend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Extend
//go:noescape
func FuncSelectionExtend(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionExtend1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Extend1
//go:noescape
func FuncSelectionExtend1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionSetBaseAndExtent(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_SetBaseAndExtent
//go:noescape
func FuncSelectionSetBaseAndExtent(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionSelectAllChildren(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_SelectAllChildren
//go:noescape
func FuncSelectionSelectAllChildren(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionModify(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify
//go:noescape
func FuncSelectionModify(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionModify1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify1
//go:noescape
func FuncSelectionModify1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionModify2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify2
//go:noescape
func FuncSelectionModify2(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionModify3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify3
//go:noescape
func FuncSelectionModify3(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionDeleteFromDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_DeleteFromDocument
//go:noescape
func FuncSelectionDeleteFromDocument(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionContainsNode(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_ContainsNode
//go:noescape
func FuncSelectionContainsNode(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionContainsNode1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_ContainsNode1
//go:noescape
func FuncSelectionContainsNode1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSelectionToString(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_Selection_ToString
//go:noescape
func FuncSelectionToString(this js.Ref, fn unsafe.Pointer)

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
func HasFuncCaretPositionGetClientRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CaretPosition_GetClientRect
//go:noescape
func FuncCaretPositionGetClientRect(this js.Ref, fn unsafe.Pointer)

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
func HasFuncXPathResultIterateNext(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathResult_IterateNext
//go:noescape
func FuncXPathResultIterateNext(this js.Ref, fn unsafe.Pointer)

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
func HasFuncXPathResultSnapshotItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathResult_SnapshotItem
//go:noescape
func FuncXPathResultSnapshotItem(this js.Ref, fn unsafe.Pointer)

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
func HasFuncXPathExpressionEvaluate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathExpression_Evaluate
//go:noescape
func FuncXPathExpressionEvaluate(this js.Ref, fn unsafe.Pointer)

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
func HasFuncXPathExpressionEvaluate1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathExpression_Evaluate1
//go:noescape
func FuncXPathExpressionEvaluate1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncXPathExpressionEvaluate2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_XPathExpression_Evaluate2
//go:noescape
func FuncXPathExpressionEvaluate2(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDocumentTypeBefore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DocumentType_Before
//go:noescape
func FuncDocumentTypeBefore(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDocumentTypeAfter(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DocumentType_After
//go:noescape
func FuncDocumentTypeAfter(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDocumentTypeReplaceWith(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DocumentType_ReplaceWith
//go:noescape
func FuncDocumentTypeReplaceWith(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDocumentTypeRemove(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DocumentType_Remove
//go:noescape
func FuncDocumentTypeRemove(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMImplementationCreateDocumentType(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateDocumentType
//go:noescape
func FuncDOMImplementationCreateDocumentType(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMImplementationCreateDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateDocument
//go:noescape
func FuncDOMImplementationCreateDocument(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMImplementationCreateDocument1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateDocument1
//go:noescape
func FuncDOMImplementationCreateDocument1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMImplementationCreateHTMLDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateHTMLDocument
//go:noescape
func FuncDOMImplementationCreateHTMLDocument(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMImplementationCreateHTMLDocument1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateHTMLDocument1
//go:noescape
func FuncDOMImplementationCreateHTMLDocument1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMImplementationHasFeature(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_HasFeature
//go:noescape
func FuncDOMImplementationHasFeature(this js.Ref, fn unsafe.Pointer)

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
func HasFuncPermissionsPolicyAllowsFeature(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_AllowsFeature
//go:noescape
func FuncPermissionsPolicyAllowsFeature(this js.Ref, fn unsafe.Pointer)

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
func HasFuncPermissionsPolicyAllowsFeature1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_AllowsFeature1
//go:noescape
func FuncPermissionsPolicyAllowsFeature1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncPermissionsPolicyFeatures(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_Features
//go:noescape
func FuncPermissionsPolicyFeatures(this js.Ref, fn unsafe.Pointer)

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
func HasFuncPermissionsPolicyAllowedFeatures(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_AllowedFeatures
//go:noescape
func FuncPermissionsPolicyAllowedFeatures(this js.Ref, fn unsafe.Pointer)

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
func HasFuncPermissionsPolicyGetAllowlistForFeature(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_GetAllowlistForFeature
//go:noescape
func FuncPermissionsPolicyGetAllowlistForFeature(this js.Ref, fn unsafe.Pointer)

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
func HasFuncCSSStyleDeclarationItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_Item
//go:noescape
func FuncCSSStyleDeclarationItem(this js.Ref, fn unsafe.Pointer)

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
func HasFuncCSSStyleDeclarationGetPropertyValue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_GetPropertyValue
//go:noescape
func FuncCSSStyleDeclarationGetPropertyValue(this js.Ref, fn unsafe.Pointer)

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
func HasFuncCSSStyleDeclarationGetPropertyPriority(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_GetPropertyPriority
//go:noescape
func FuncCSSStyleDeclarationGetPropertyPriority(this js.Ref, fn unsafe.Pointer)

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
func HasFuncCSSStyleDeclarationSetProperty(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_SetProperty
//go:noescape
func FuncCSSStyleDeclarationSetProperty(this js.Ref, fn unsafe.Pointer)

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
func HasFuncCSSStyleDeclarationSetProperty1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_SetProperty1
//go:noescape
func FuncCSSStyleDeclarationSetProperty1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncCSSStyleDeclarationRemoveProperty(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_RemoveProperty
//go:noescape
func FuncCSSStyleDeclarationRemoveProperty(this js.Ref, fn unsafe.Pointer)

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
func HasFuncStylePropertyMapSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Set
//go:noescape
func FuncStylePropertyMapSet(this js.Ref, fn unsafe.Pointer)

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
func HasFuncStylePropertyMapAppend(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Append
//go:noescape
func FuncStylePropertyMapAppend(this js.Ref, fn unsafe.Pointer)

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
func HasFuncStylePropertyMapDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Delete
//go:noescape
func FuncStylePropertyMapDelete(this js.Ref, fn unsafe.Pointer)

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
func HasFuncStylePropertyMapClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Clear
//go:noescape
func FuncStylePropertyMapClear(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGLengthNewValueSpecifiedUnits(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLength_NewValueSpecifiedUnits
//go:noescape
func FuncSVGLengthNewValueSpecifiedUnits(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGLengthConvertToSpecifiedUnits(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGLength_ConvertToSpecifiedUnits
//go:noescape
func FuncSVGLengthConvertToSpecifiedUnits(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMStringMapGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMStringMap_Get
//go:noescape
func FuncDOMStringMapGet(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMStringMapSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMStringMap_Set
//go:noescape
func FuncDOMStringMapSet(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMStringMapDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMStringMap_Delete
//go:noescape
func FuncDOMStringMapDelete(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGElementFocus(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGElement_Focus
//go:noescape
func FuncSVGElementFocus(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGElementFocus1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGElement_Focus1
//go:noescape
func FuncSVGElementFocus1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGElementBlur(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGElement_Blur
//go:noescape
func FuncSVGElementBlur(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGAngleNewValueSpecifiedUnits(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGAngle_NewValueSpecifiedUnits
//go:noescape
func FuncSVGAngleNewValueSpecifiedUnits(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGAngleConvertToSpecifiedUnits(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGAngle_ConvertToSpecifiedUnits
//go:noescape
func FuncSVGAngleConvertToSpecifiedUnits(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixFromMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromMatrix
//go:noescape
func FuncDOMMatrixFromMatrix(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixFromMatrix1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromMatrix1
//go:noescape
func FuncDOMMatrixFromMatrix1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixFromFloat32Array(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromFloat32Array
//go:noescape
func FuncDOMMatrixFromFloat32Array(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixFromFloat64Array(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromFloat64Array
//go:noescape
func FuncDOMMatrixFromFloat64Array(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixMultiplySelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_MultiplySelf
//go:noescape
func FuncDOMMatrixMultiplySelf(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixMultiplySelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_MultiplySelf1
//go:noescape
func FuncDOMMatrixMultiplySelf1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixPreMultiplySelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_PreMultiplySelf
//go:noescape
func FuncDOMMatrixPreMultiplySelf(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixPreMultiplySelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_PreMultiplySelf1
//go:noescape
func FuncDOMMatrixPreMultiplySelf1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixTranslateSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf
//go:noescape
func FuncDOMMatrixTranslateSelf(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixTranslateSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf1
//go:noescape
func FuncDOMMatrixTranslateSelf1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixTranslateSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf2
//go:noescape
func FuncDOMMatrixTranslateSelf2(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixTranslateSelf3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf3
//go:noescape
func FuncDOMMatrixTranslateSelf3(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScaleSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf
//go:noescape
func FuncDOMMatrixScaleSelf(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScaleSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf1
//go:noescape
func FuncDOMMatrixScaleSelf1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScaleSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf2
//go:noescape
func FuncDOMMatrixScaleSelf2(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScaleSelf3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf3
//go:noescape
func FuncDOMMatrixScaleSelf3(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScaleSelf4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf4
//go:noescape
func FuncDOMMatrixScaleSelf4(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScaleSelf5(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf5
//go:noescape
func FuncDOMMatrixScaleSelf5(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScaleSelf6(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf6
//go:noescape
func FuncDOMMatrixScaleSelf6(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScale3dSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf
//go:noescape
func FuncDOMMatrixScale3dSelf(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScale3dSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf1
//go:noescape
func FuncDOMMatrixScale3dSelf1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScale3dSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf2
//go:noescape
func FuncDOMMatrixScale3dSelf2(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScale3dSelf3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf3
//go:noescape
func FuncDOMMatrixScale3dSelf3(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixScale3dSelf4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf4
//go:noescape
func FuncDOMMatrixScale3dSelf4(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf
//go:noescape
func FuncDOMMatrixRotateSelf(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf1
//go:noescape
func FuncDOMMatrixRotateSelf1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf2
//go:noescape
func FuncDOMMatrixRotateSelf2(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateSelf3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf3
//go:noescape
func FuncDOMMatrixRotateSelf3(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateFromVectorSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateFromVectorSelf
//go:noescape
func FuncDOMMatrixRotateFromVectorSelf(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateFromVectorSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateFromVectorSelf1
//go:noescape
func FuncDOMMatrixRotateFromVectorSelf1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateFromVectorSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateFromVectorSelf2
//go:noescape
func FuncDOMMatrixRotateFromVectorSelf2(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateAxisAngleSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf
//go:noescape
func FuncDOMMatrixRotateAxisAngleSelf(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateAxisAngleSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf1
//go:noescape
func FuncDOMMatrixRotateAxisAngleSelf1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateAxisAngleSelf2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf2
//go:noescape
func FuncDOMMatrixRotateAxisAngleSelf2(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateAxisAngleSelf3(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf3
//go:noescape
func FuncDOMMatrixRotateAxisAngleSelf3(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixRotateAxisAngleSelf4(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf4
//go:noescape
func FuncDOMMatrixRotateAxisAngleSelf4(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixSkewXSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewXSelf
//go:noescape
func FuncDOMMatrixSkewXSelf(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixSkewXSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewXSelf1
//go:noescape
func FuncDOMMatrixSkewXSelf1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixSkewYSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewYSelf
//go:noescape
func FuncDOMMatrixSkewYSelf(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixSkewYSelf1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewYSelf1
//go:noescape
func FuncDOMMatrixSkewYSelf1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixInvertSelf(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_InvertSelf
//go:noescape
func FuncDOMMatrixInvertSelf(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMMatrixSetMatrixValue(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SetMatrixValue
//go:noescape
func FuncDOMMatrixSetMatrixValue(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGTransformSetMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetMatrix
//go:noescape
func FuncSVGTransformSetMatrix(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGTransformSetMatrix1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetMatrix1
//go:noescape
func FuncSVGTransformSetMatrix1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGTransformSetTranslate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetTranslate
//go:noescape
func FuncSVGTransformSetTranslate(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGTransformSetScale(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetScale
//go:noescape
func FuncSVGTransformSetScale(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGTransformSetRotate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetRotate
//go:noescape
func FuncSVGTransformSetRotate(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGTransformSetSkewX(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetSkewX
//go:noescape
func FuncSVGTransformSetSkewX(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGTransformSetSkewY(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetSkewY
//go:noescape
func FuncSVGTransformSetSkewY(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMPointReadOnlyFromPoint(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_FromPoint
//go:noescape
func FuncDOMPointReadOnlyFromPoint(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMPointReadOnlyFromPoint1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_FromPoint1
//go:noescape
func FuncDOMPointReadOnlyFromPoint1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMPointReadOnlyMatrixTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_MatrixTransform
//go:noescape
func FuncDOMPointReadOnlyMatrixTransform(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMPointReadOnlyMatrixTransform1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_MatrixTransform1
//go:noescape
func FuncDOMPointReadOnlyMatrixTransform1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncDOMPointReadOnlyToJSON(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_ToJSON
//go:noescape
func FuncDOMPointReadOnlyToJSON(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementGetIntersectionList(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_GetIntersectionList
//go:noescape
func FuncSVGSVGElementGetIntersectionList(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementGetEnclosureList(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_GetEnclosureList
//go:noescape
func FuncSVGSVGElementGetEnclosureList(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementCheckIntersection(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CheckIntersection
//go:noescape
func FuncSVGSVGElementCheckIntersection(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementCheckEnclosure(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CheckEnclosure
//go:noescape
func FuncSVGSVGElementCheckEnclosure(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementDeselectAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_DeselectAll
//go:noescape
func FuncSVGSVGElementDeselectAll(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementCreateSVGNumber(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGNumber
//go:noescape
func FuncSVGSVGElementCreateSVGNumber(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementCreateSVGLength(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGLength
//go:noescape
func FuncSVGSVGElementCreateSVGLength(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementCreateSVGAngle(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGAngle
//go:noescape
func FuncSVGSVGElementCreateSVGAngle(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementCreateSVGPoint(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGPoint
//go:noescape
func FuncSVGSVGElementCreateSVGPoint(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementCreateSVGMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGMatrix
//go:noescape
func FuncSVGSVGElementCreateSVGMatrix(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementCreateSVGRect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGRect
//go:noescape
func FuncSVGSVGElementCreateSVGRect(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementCreateSVGTransform(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGTransform
//go:noescape
func FuncSVGSVGElementCreateSVGTransform(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementCreateSVGTransformFromMatrix(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGTransformFromMatrix
//go:noescape
func FuncSVGSVGElementCreateSVGTransformFromMatrix(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementCreateSVGTransformFromMatrix1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGTransformFromMatrix1
//go:noescape
func FuncSVGSVGElementCreateSVGTransformFromMatrix1(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementGetElementById(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_GetElementById
//go:noescape
func FuncSVGSVGElementGetElementById(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementSuspendRedraw(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_SuspendRedraw
//go:noescape
func FuncSVGSVGElementSuspendRedraw(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementUnsuspendRedraw(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_UnsuspendRedraw
//go:noescape
func FuncSVGSVGElementUnsuspendRedraw(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementUnsuspendRedrawAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_UnsuspendRedrawAll
//go:noescape
func FuncSVGSVGElementUnsuspendRedrawAll(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementForceRedraw(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_ForceRedraw
//go:noescape
func FuncSVGSVGElementForceRedraw(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementPauseAnimations(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_PauseAnimations
//go:noescape
func FuncSVGSVGElementPauseAnimations(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementUnpauseAnimations(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_UnpauseAnimations
//go:noescape
func FuncSVGSVGElementUnpauseAnimations(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementAnimationsPaused(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_AnimationsPaused
//go:noescape
func FuncSVGSVGElementAnimationsPaused(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementGetCurrentTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_GetCurrentTime
//go:noescape
func FuncSVGSVGElementGetCurrentTime(this js.Ref, fn unsafe.Pointer)

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
func HasFuncSVGSVGElementSetCurrentTime(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_SetCurrentTime
//go:noescape
func FuncSVGSVGElementSetCurrentTime(this js.Ref, fn unsafe.Pointer)

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

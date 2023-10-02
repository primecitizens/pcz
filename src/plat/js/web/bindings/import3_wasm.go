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

//go:wasmimport plat/js/web get_Baseline_Name
//go:noescape

func GetBaselineName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Baseline_Value
//go:noescape

func GetBaselineValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_Font_Name
//go:noescape

func GetFontName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Font_GlyphsRendered
//go:noescape

func GetFontGlyphsRendered(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_FontMetrics_Width
//go:noescape

func GetFontMetricsWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontMetrics_Advances
//go:noescape

func GetFontMetricsAdvances(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FontMetrics_BoundingBoxLeft
//go:noescape

func GetFontMetricsBoundingBoxLeft(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontMetrics_BoundingBoxRight
//go:noescape

func GetFontMetricsBoundingBoxRight(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontMetrics_Height
//go:noescape

func GetFontMetricsHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontMetrics_EmHeightAscent
//go:noescape

func GetFontMetricsEmHeightAscent(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontMetrics_EmHeightDescent
//go:noescape

func GetFontMetricsEmHeightDescent(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontMetrics_BoundingBoxAscent
//go:noescape

func GetFontMetricsBoundingBoxAscent(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontMetrics_BoundingBoxDescent
//go:noescape

func GetFontMetricsBoundingBoxDescent(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontMetrics_FontBoundingBoxAscent
//go:noescape

func GetFontMetricsFontBoundingBoxAscent(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontMetrics_FontBoundingBoxDescent
//go:noescape

func GetFontMetricsFontBoundingBoxDescent(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_FontMetrics_DominantBaseline
//go:noescape

func GetFontMetricsDominantBaseline(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FontMetrics_Baselines
//go:noescape

func GetFontMetricsBaselines(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_FontMetrics_Fonts
//go:noescape

func GetFontMetricsFonts(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Selection_AnchorOffset
//go:noescape

func GetSelectionAnchorOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Selection_FocusNode
//go:noescape

func GetSelectionFocusNode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Selection_FocusOffset
//go:noescape

func GetSelectionFocusOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Selection_IsCollapsed
//go:noescape

func GetSelectionIsCollapsed(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Selection_RangeCount
//go:noescape

func GetSelectionRangeCount(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Selection_Type
//go:noescape

func GetSelectionType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Selection_Direction
//go:noescape

func GetSelectionDirection(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Selection_GetRangeAt
//go:noescape

func CallSelectionGetRangeAt(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_Selection_GetRangeAt
//go:noescape

func SelectionGetRangeAtFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_AddRange
//go:noescape

func CallSelectionAddRange(
	this js.Ref,
	ptr unsafe.Pointer,

	rng js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Selection_AddRange
//go:noescape

func SelectionAddRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_RemoveRange
//go:noescape

func CallSelectionRemoveRange(
	this js.Ref,
	ptr unsafe.Pointer,

	rng js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Selection_RemoveRange
//go:noescape

func SelectionRemoveRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_RemoveAllRanges
//go:noescape

func CallSelectionRemoveAllRanges(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Selection_RemoveAllRanges
//go:noescape

func SelectionRemoveAllRangesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Empty
//go:noescape

func CallSelectionEmpty(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Selection_Empty
//go:noescape

func SelectionEmptyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_GetComposedRanges
//go:noescape

func CallSelectionGetComposedRanges(
	this js.Ref,
	ptr unsafe.Pointer,

	shadowRootsPtr unsafe.Pointer,
	shadowRootsCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Selection_GetComposedRanges
//go:noescape

func SelectionGetComposedRangesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Collapse
//go:noescape

func CallSelectionCollapse(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
	offset uint32,
) js.Ref

//go:wasmimport plat/js/web func_Selection_Collapse
//go:noescape

func SelectionCollapseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Collapse1
//go:noescape

func CallSelectionCollapse1(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Selection_Collapse1
//go:noescape

func SelectionCollapse1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_SetPosition
//go:noescape

func CallSelectionSetPosition(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
	offset uint32,
) js.Ref

//go:wasmimport plat/js/web func_Selection_SetPosition
//go:noescape

func SelectionSetPositionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_SetPosition1
//go:noescape

func CallSelectionSetPosition1(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Selection_SetPosition1
//go:noescape

func SelectionSetPosition1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_CollapseToStart
//go:noescape

func CallSelectionCollapseToStart(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Selection_CollapseToStart
//go:noescape

func SelectionCollapseToStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_CollapseToEnd
//go:noescape

func CallSelectionCollapseToEnd(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Selection_CollapseToEnd
//go:noescape

func SelectionCollapseToEndFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Extend
//go:noescape

func CallSelectionExtend(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
	offset uint32,
) js.Ref

//go:wasmimport plat/js/web func_Selection_Extend
//go:noescape

func SelectionExtendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Extend1
//go:noescape

func CallSelectionExtend1(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Selection_Extend1
//go:noescape

func SelectionExtend1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_SetBaseAndExtent
//go:noescape

func CallSelectionSetBaseAndExtent(
	this js.Ref,
	ptr unsafe.Pointer,

	anchorNode js.Ref,
	anchorOffset uint32,
	focusNode js.Ref,
	focusOffset uint32,
) js.Ref

//go:wasmimport plat/js/web func_Selection_SetBaseAndExtent
//go:noescape

func SelectionSetBaseAndExtentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_SelectAllChildren
//go:noescape

func CallSelectionSelectAllChildren(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Selection_SelectAllChildren
//go:noescape

func SelectionSelectAllChildrenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Modify
//go:noescape

func CallSelectionModify(
	this js.Ref,
	ptr unsafe.Pointer,

	alter js.Ref,
	direction js.Ref,
	granularity js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify
//go:noescape

func SelectionModifyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Modify1
//go:noescape

func CallSelectionModify1(
	this js.Ref,
	ptr unsafe.Pointer,

	alter js.Ref,
	direction js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify1
//go:noescape

func SelectionModify1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Modify2
//go:noescape

func CallSelectionModify2(
	this js.Ref,
	ptr unsafe.Pointer,

	alter js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify2
//go:noescape

func SelectionModify2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_Modify3
//go:noescape

func CallSelectionModify3(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Selection_Modify3
//go:noescape

func SelectionModify3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_DeleteFromDocument
//go:noescape

func CallSelectionDeleteFromDocument(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Selection_DeleteFromDocument
//go:noescape

func SelectionDeleteFromDocumentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_ContainsNode
//go:noescape

func CallSelectionContainsNode(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
	allowPartialContainment js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Selection_ContainsNode
//go:noescape

func SelectionContainsNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_ContainsNode1
//go:noescape

func CallSelectionContainsNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Selection_ContainsNode1
//go:noescape

func SelectionContainsNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Selection_ToString
//go:noescape

func CallSelectionToString(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Selection_ToString
//go:noescape

func SelectionToStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_CaretPosition_OffsetNode
//go:noescape

func GetCaretPositionOffsetNode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CaretPosition_Offset
//go:noescape

func GetCaretPositionOffset(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_CaretPosition_GetClientRect
//go:noescape

func CallCaretPositionGetClientRect(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_CaretPosition_GetClientRect
//go:noescape

func CaretPositionGetClientRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_XPathResult_ResultType
//go:noescape

func GetXPathResultResultType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_XPathResult_NumberValue
//go:noescape

func GetXPathResultNumberValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_XPathResult_StringValue
//go:noescape

func GetXPathResultStringValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XPathResult_BooleanValue
//go:noescape

func GetXPathResultBooleanValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XPathResult_SingleNodeValue
//go:noescape

func GetXPathResultSingleNodeValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XPathResult_InvalidIteratorState
//go:noescape

func GetXPathResultInvalidIteratorState(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_XPathResult_SnapshotLength
//go:noescape

func GetXPathResultSnapshotLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_XPathResult_IterateNext
//go:noescape

func CallXPathResultIterateNext(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_XPathResult_IterateNext
//go:noescape

func XPathResultIterateNextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathResult_SnapshotItem
//go:noescape

func CallXPathResultSnapshotItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_XPathResult_SnapshotItem
//go:noescape

func XPathResultSnapshotItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathExpression_Evaluate
//go:noescape

func CallXPathExpressionEvaluate(
	this js.Ref,
	ptr unsafe.Pointer,

	contextNode js.Ref,
	typ uint32,
	result js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XPathExpression_Evaluate
//go:noescape

func XPathExpressionEvaluateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathExpression_Evaluate1
//go:noescape

func CallXPathExpressionEvaluate1(
	this js.Ref,
	ptr unsafe.Pointer,

	contextNode js.Ref,
	typ uint32,
) js.Ref

//go:wasmimport plat/js/web func_XPathExpression_Evaluate1
//go:noescape

func XPathExpressionEvaluate1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_XPathExpression_Evaluate2
//go:noescape

func CallXPathExpressionEvaluate2(
	this js.Ref,
	ptr unsafe.Pointer,

	contextNode js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_XPathExpression_Evaluate2
//go:noescape

func XPathExpressionEvaluate2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_DocumentType_Name
//go:noescape

func GetDocumentTypeName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DocumentType_PublicId
//go:noescape

func GetDocumentTypePublicId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DocumentType_SystemId
//go:noescape

func GetDocumentTypeSystemId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_DocumentType_Before
//go:noescape

func CallDocumentTypeBefore(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_DocumentType_Before
//go:noescape

func DocumentTypeBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentType_After
//go:noescape

func CallDocumentTypeAfter(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_DocumentType_After
//go:noescape

func DocumentTypeAfterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentType_ReplaceWith
//go:noescape

func CallDocumentTypeReplaceWith(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_DocumentType_ReplaceWith
//go:noescape

func DocumentTypeReplaceWithFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentType_Remove
//go:noescape

func CallDocumentTypeRemove(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DocumentType_Remove
//go:noescape

func DocumentTypeRemoveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_CreateDocumentType
//go:noescape

func CallDOMImplementationCreateDocumentType(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
	publicId js.Ref,
	systemId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateDocumentType
//go:noescape

func DOMImplementationCreateDocumentTypeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_CreateDocument
//go:noescape

func CallDOMImplementationCreateDocument(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	qualifiedName js.Ref,
	doctype js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateDocument
//go:noescape

func DOMImplementationCreateDocumentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_CreateDocument1
//go:noescape

func CallDOMImplementationCreateDocument1(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateDocument1
//go:noescape

func DOMImplementationCreateDocument1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_CreateHTMLDocument
//go:noescape

func CallDOMImplementationCreateHTMLDocument(
	this js.Ref,
	ptr unsafe.Pointer,

	title js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateHTMLDocument
//go:noescape

func DOMImplementationCreateHTMLDocumentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_CreateHTMLDocument1
//go:noescape

func CallDOMImplementationCreateHTMLDocument1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_CreateHTMLDocument1
//go:noescape

func DOMImplementationCreateHTMLDocument1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMImplementation_HasFeature
//go:noescape

func CallDOMImplementationHasFeature(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMImplementation_HasFeature
//go:noescape

func DOMImplementationHasFeatureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PermissionsPolicy_AllowsFeature
//go:noescape

func CallPermissionsPolicyAllowsFeature(
	this js.Ref,
	ptr unsafe.Pointer,

	feature js.Ref,
	origin js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_AllowsFeature
//go:noescape

func PermissionsPolicyAllowsFeatureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PermissionsPolicy_AllowsFeature1
//go:noescape

func CallPermissionsPolicyAllowsFeature1(
	this js.Ref,
	ptr unsafe.Pointer,

	feature js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_AllowsFeature1
//go:noescape

func PermissionsPolicyAllowsFeature1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PermissionsPolicy_Features
//go:noescape

func CallPermissionsPolicyFeatures(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_Features
//go:noescape

func PermissionsPolicyFeaturesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PermissionsPolicy_AllowedFeatures
//go:noescape

func CallPermissionsPolicyAllowedFeatures(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_AllowedFeatures
//go:noescape

func PermissionsPolicyAllowedFeaturesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_PermissionsPolicy_GetAllowlistForFeature
//go:noescape

func CallPermissionsPolicyGetAllowlistForFeature(
	this js.Ref,
	ptr unsafe.Pointer,

	feature js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_PermissionsPolicy_GetAllowlistForFeature
//go:noescape

func PermissionsPolicyGetAllowlistForFeatureFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SVGAnimatedString_BaseVal
//go:noescape

func SetSVGAnimatedStringBaseVal(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedString_AnimVal
//go:noescape

func GetSVGAnimatedStringAnimVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSSStyleDeclaration_CssText
//go:noescape

func GetCSSStyleDeclarationCssText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_CSSStyleDeclaration_CssText
//go:noescape

func SetCSSStyleDeclarationCssText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_CSSStyleDeclaration_Length
//go:noescape

func GetCSSStyleDeclarationLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_CSSStyleDeclaration_ParentRule
//go:noescape

func GetCSSStyleDeclarationParentRule(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSSStyleDeclaration_CssFloat
//go:noescape

func GetCSSStyleDeclarationCssFloat(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_CSSStyleDeclaration_CssFloat
//go:noescape

func SetCSSStyleDeclarationCssFloat(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_Item
//go:noescape

func CallCSSStyleDeclarationItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_Item
//go:noescape

func CSSStyleDeclarationItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_GetPropertyValue
//go:noescape

func CallCSSStyleDeclarationGetPropertyValue(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_GetPropertyValue
//go:noescape

func CSSStyleDeclarationGetPropertyValueFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_GetPropertyPriority
//go:noescape

func CallCSSStyleDeclarationGetPropertyPriority(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_GetPropertyPriority
//go:noescape

func CSSStyleDeclarationGetPropertyPriorityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_SetProperty
//go:noescape

func CallCSSStyleDeclarationSetProperty(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
	value js.Ref,
	priority js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_SetProperty
//go:noescape

func CSSStyleDeclarationSetPropertyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_SetProperty1
//go:noescape

func CallCSSStyleDeclarationSetProperty1(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_SetProperty1
//go:noescape

func CSSStyleDeclarationSetProperty1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSStyleDeclaration_RemoveProperty
//go:noescape

func CallCSSStyleDeclarationRemoveProperty(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSStyleDeclaration_RemoveProperty
//go:noescape

func CSSStyleDeclarationRemovePropertyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StylePropertyMap_Set
//go:noescape

func CallStylePropertyMapSet(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
	valuesPtr unsafe.Pointer,
	valuesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Set
//go:noescape

func StylePropertyMapSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StylePropertyMap_Append
//go:noescape

func CallStylePropertyMapAppend(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
	valuesPtr unsafe.Pointer,
	valuesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Append
//go:noescape

func StylePropertyMapAppendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StylePropertyMap_Delete
//go:noescape

func CallStylePropertyMapDelete(
	this js.Ref,
	ptr unsafe.Pointer,

	property js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Delete
//go:noescape

func StylePropertyMapDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_StylePropertyMap_Clear
//go:noescape

func CallStylePropertyMapClear(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_StylePropertyMap_Clear
//go:noescape

func StylePropertyMapClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGLength_UnitType
//go:noescape

func GetSVGLengthUnitType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SVGLength_Value
//go:noescape

func GetSVGLengthValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_SVGLength_Value
//go:noescape

func SetSVGLengthValue(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGLength_ValueInSpecifiedUnits
//go:noescape

func GetSVGLengthValueInSpecifiedUnits(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_SVGLength_ValueInSpecifiedUnits
//go:noescape

func SetSVGLengthValueInSpecifiedUnits(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGLength_ValueAsString
//go:noescape

func GetSVGLengthValueAsString(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SVGLength_ValueAsString
//go:noescape

func SetSVGLengthValueAsString(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_SVGLength_NewValueSpecifiedUnits
//go:noescape

func CallSVGLengthNewValueSpecifiedUnits(
	this js.Ref,
	ptr unsafe.Pointer,

	unitType uint32,
	valueInSpecifiedUnits float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGLength_NewValueSpecifiedUnits
//go:noescape

func SVGLengthNewValueSpecifiedUnitsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGLength_ConvertToSpecifiedUnits
//go:noescape

func CallSVGLengthConvertToSpecifiedUnits(
	this js.Ref,
	ptr unsafe.Pointer,

	unitType uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGLength_ConvertToSpecifiedUnits
//go:noescape

func SVGLengthConvertToSpecifiedUnitsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedLength_BaseVal
//go:noescape

func GetSVGAnimatedLengthBaseVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedLength_AnimVal
//go:noescape

func GetSVGAnimatedLengthAnimVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGUseElement_X
//go:noescape

func GetSVGUseElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGUseElement_Y
//go:noescape

func GetSVGUseElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGUseElement_Width
//go:noescape

func GetSVGUseElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGUseElement_Height
//go:noescape

func GetSVGUseElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGUseElement_InstanceRoot
//go:noescape

func GetSVGUseElementInstanceRoot(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGUseElement_AnimatedInstanceRoot
//go:noescape

func GetSVGUseElementAnimatedInstanceRoot(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGUseElement_Href
//go:noescape

func GetSVGUseElementHref(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_DOMStringMap_Get
//go:noescape

func CallDOMStringMapGet(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMStringMap_Get
//go:noescape

func DOMStringMapGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMStringMap_Set
//go:noescape

func CallDOMStringMapSet(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMStringMap_Set
//go:noescape

func DOMStringMapSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMStringMap_Delete
//go:noescape

func CallDOMStringMapDelete(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMStringMap_Delete
//go:noescape

func DOMStringMapDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGElement_ClassName
//go:noescape

func GetSVGElementClassName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_OwnerSVGElement
//go:noescape

func GetSVGElementOwnerSVGElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_ViewportElement
//go:noescape

func GetSVGElementViewportElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_Style
//go:noescape

func GetSVGElementStyle(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_AttributeStyleMap
//go:noescape

func GetSVGElementAttributeStyleMap(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_CorrespondingElement
//go:noescape

func GetSVGElementCorrespondingElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_CorrespondingUseElement
//go:noescape

func GetSVGElementCorrespondingUseElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_Dataset
//go:noescape

func GetSVGElementDataset(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_Nonce
//go:noescape

func GetSVGElementNonce(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SVGElement_Nonce
//go:noescape

func SetSVGElementNonce(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_Autofocus
//go:noescape

func GetSVGElementAutofocus(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SVGElement_Autofocus
//go:noescape

func SetSVGElementAutofocus(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_SVGElement_TabIndex
//go:noescape

func GetSVGElementTabIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web set_SVGElement_TabIndex
//go:noescape

func SetSVGElementTabIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web call_SVGElement_Focus
//go:noescape

func CallSVGElementFocus(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_SVGElement_Focus
//go:noescape

func SVGElementFocusFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGElement_Focus1
//go:noescape

func CallSVGElementFocus1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGElement_Focus1
//go:noescape

func SVGElementFocus1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGElement_Blur
//go:noescape

func CallSVGElementBlur(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGElement_Blur
//go:noescape

func SVGElementBlurFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGNumber_Value
//go:noescape

func GetSVGNumberValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_SVGNumber_Value
//go:noescape

func SetSVGNumberValue(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAngle_UnitType
//go:noescape

func GetSVGAngleUnitType(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SVGAngle_Value
//go:noescape

func GetSVGAngleValue(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_SVGAngle_Value
//go:noescape

func SetSVGAngleValue(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAngle_ValueInSpecifiedUnits
//go:noescape

func GetSVGAngleValueInSpecifiedUnits(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_SVGAngle_ValueInSpecifiedUnits
//go:noescape

func SetSVGAngleValueInSpecifiedUnits(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAngle_ValueAsString
//go:noescape

func GetSVGAngleValueAsString(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_SVGAngle_ValueAsString
//go:noescape

func SetSVGAngleValueAsString(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_SVGAngle_NewValueSpecifiedUnits
//go:noescape

func CallSVGAngleNewValueSpecifiedUnits(
	this js.Ref,
	ptr unsafe.Pointer,

	unitType uint32,
	valueInSpecifiedUnits float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGAngle_NewValueSpecifiedUnits
//go:noescape

func SVGAngleNewValueSpecifiedUnitsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGAngle_ConvertToSpecifiedUnits
//go:noescape

func CallSVGAngleConvertToSpecifiedUnits(
	this js.Ref,
	ptr unsafe.Pointer,

	unitType uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGAngle_ConvertToSpecifiedUnits
//go:noescape

func SVGAngleConvertToSpecifiedUnitsFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_A
//go:noescape

func SetDOMMatrixA(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_B
//go:noescape

func GetDOMMatrixB(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_B
//go:noescape

func SetDOMMatrixB(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_C
//go:noescape

func GetDOMMatrixC(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_C
//go:noescape

func SetDOMMatrixC(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_D
//go:noescape

func GetDOMMatrixD(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_D
//go:noescape

func SetDOMMatrixD(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_E
//go:noescape

func GetDOMMatrixE(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_E
//go:noescape

func SetDOMMatrixE(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_F
//go:noescape

func GetDOMMatrixF(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_F
//go:noescape

func SetDOMMatrixF(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M11
//go:noescape

func GetDOMMatrixM11(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M11
//go:noescape

func SetDOMMatrixM11(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M12
//go:noescape

func GetDOMMatrixM12(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M12
//go:noescape

func SetDOMMatrixM12(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M13
//go:noescape

func GetDOMMatrixM13(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M13
//go:noescape

func SetDOMMatrixM13(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M14
//go:noescape

func GetDOMMatrixM14(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M14
//go:noescape

func SetDOMMatrixM14(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M21
//go:noescape

func GetDOMMatrixM21(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M21
//go:noescape

func SetDOMMatrixM21(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M22
//go:noescape

func GetDOMMatrixM22(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M22
//go:noescape

func SetDOMMatrixM22(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M23
//go:noescape

func GetDOMMatrixM23(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M23
//go:noescape

func SetDOMMatrixM23(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M24
//go:noescape

func GetDOMMatrixM24(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M24
//go:noescape

func SetDOMMatrixM24(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M31
//go:noescape

func GetDOMMatrixM31(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M31
//go:noescape

func SetDOMMatrixM31(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M32
//go:noescape

func GetDOMMatrixM32(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M32
//go:noescape

func SetDOMMatrixM32(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M33
//go:noescape

func GetDOMMatrixM33(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M33
//go:noescape

func SetDOMMatrixM33(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M34
//go:noescape

func GetDOMMatrixM34(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M34
//go:noescape

func SetDOMMatrixM34(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M41
//go:noescape

func GetDOMMatrixM41(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M41
//go:noescape

func SetDOMMatrixM41(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M42
//go:noescape

func GetDOMMatrixM42(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M42
//go:noescape

func SetDOMMatrixM42(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M43
//go:noescape

func GetDOMMatrixM43(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M43
//go:noescape

func SetDOMMatrixM43(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_DOMMatrix_M44
//go:noescape

func GetDOMMatrixM44(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_DOMMatrix_M44
//go:noescape

func SetDOMMatrixM44(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_FromMatrix
//go:noescape

func CallDOMMatrixFromMatrix(
	this js.Ref,
	ptr unsafe.Pointer,

	other unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromMatrix
//go:noescape

func DOMMatrixFromMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_FromMatrix1
//go:noescape

func CallDOMMatrixFromMatrix1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromMatrix1
//go:noescape

func DOMMatrixFromMatrix1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_FromFloat32Array
//go:noescape

func CallDOMMatrixFromFloat32Array(
	this js.Ref,
	ptr unsafe.Pointer,

	array32 js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromFloat32Array
//go:noescape

func DOMMatrixFromFloat32ArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_FromFloat64Array
//go:noescape

func CallDOMMatrixFromFloat64Array(
	this js.Ref,
	ptr unsafe.Pointer,

	array64 js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_FromFloat64Array
//go:noescape

func DOMMatrixFromFloat64ArrayFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_MultiplySelf
//go:noescape

func CallDOMMatrixMultiplySelf(
	this js.Ref,
	ptr unsafe.Pointer,

	other unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_MultiplySelf
//go:noescape

func DOMMatrixMultiplySelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_MultiplySelf1
//go:noescape

func CallDOMMatrixMultiplySelf1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_MultiplySelf1
//go:noescape

func DOMMatrixMultiplySelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_PreMultiplySelf
//go:noescape

func CallDOMMatrixPreMultiplySelf(
	this js.Ref,
	ptr unsafe.Pointer,

	other unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_PreMultiplySelf
//go:noescape

func DOMMatrixPreMultiplySelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_PreMultiplySelf1
//go:noescape

func CallDOMMatrixPreMultiplySelf1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_PreMultiplySelf1
//go:noescape

func DOMMatrixPreMultiplySelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_TranslateSelf
//go:noescape

func CallDOMMatrixTranslateSelf(
	this js.Ref,
	ptr unsafe.Pointer,

	tx float64,
	ty float64,
	tz float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf
//go:noescape

func DOMMatrixTranslateSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_TranslateSelf1
//go:noescape

func CallDOMMatrixTranslateSelf1(
	this js.Ref,
	ptr unsafe.Pointer,

	tx float64,
	ty float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf1
//go:noescape

func DOMMatrixTranslateSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_TranslateSelf2
//go:noescape

func CallDOMMatrixTranslateSelf2(
	this js.Ref,
	ptr unsafe.Pointer,

	tx float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf2
//go:noescape

func DOMMatrixTranslateSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_TranslateSelf3
//go:noescape

func CallDOMMatrixTranslateSelf3(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_TranslateSelf3
//go:noescape

func DOMMatrixTranslateSelf3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf
//go:noescape

func CallDOMMatrixScaleSelf(
	this js.Ref,
	ptr unsafe.Pointer,

	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64,
	originY float64,
	originZ float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf
//go:noescape

func DOMMatrixScaleSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf1
//go:noescape

func CallDOMMatrixScaleSelf1(
	this js.Ref,
	ptr unsafe.Pointer,

	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64,
	originY float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf1
//go:noescape

func DOMMatrixScaleSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf2
//go:noescape

func CallDOMMatrixScaleSelf2(
	this js.Ref,
	ptr unsafe.Pointer,

	scaleX float64,
	scaleY float64,
	scaleZ float64,
	originX float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf2
//go:noescape

func DOMMatrixScaleSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf3
//go:noescape

func CallDOMMatrixScaleSelf3(
	this js.Ref,
	ptr unsafe.Pointer,

	scaleX float64,
	scaleY float64,
	scaleZ float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf3
//go:noescape

func DOMMatrixScaleSelf3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf4
//go:noescape

func CallDOMMatrixScaleSelf4(
	this js.Ref,
	ptr unsafe.Pointer,

	scaleX float64,
	scaleY float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf4
//go:noescape

func DOMMatrixScaleSelf4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf5
//go:noescape

func CallDOMMatrixScaleSelf5(
	this js.Ref,
	ptr unsafe.Pointer,

	scaleX float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf5
//go:noescape

func DOMMatrixScaleSelf5Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_ScaleSelf6
//go:noescape

func CallDOMMatrixScaleSelf6(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_ScaleSelf6
//go:noescape

func DOMMatrixScaleSelf6Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_Scale3dSelf
//go:noescape

func CallDOMMatrixScale3dSelf(
	this js.Ref,
	ptr unsafe.Pointer,

	scale float64,
	originX float64,
	originY float64,
	originZ float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf
//go:noescape

func DOMMatrixScale3dSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_Scale3dSelf1
//go:noescape

func CallDOMMatrixScale3dSelf1(
	this js.Ref,
	ptr unsafe.Pointer,

	scale float64,
	originX float64,
	originY float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf1
//go:noescape

func DOMMatrixScale3dSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_Scale3dSelf2
//go:noescape

func CallDOMMatrixScale3dSelf2(
	this js.Ref,
	ptr unsafe.Pointer,

	scale float64,
	originX float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf2
//go:noescape

func DOMMatrixScale3dSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_Scale3dSelf3
//go:noescape

func CallDOMMatrixScale3dSelf3(
	this js.Ref,
	ptr unsafe.Pointer,

	scale float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf3
//go:noescape

func DOMMatrixScale3dSelf3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_Scale3dSelf4
//go:noescape

func CallDOMMatrixScale3dSelf4(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_Scale3dSelf4
//go:noescape

func DOMMatrixScale3dSelf4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateSelf
//go:noescape

func CallDOMMatrixRotateSelf(
	this js.Ref,
	ptr unsafe.Pointer,

	rotX float64,
	rotY float64,
	rotZ float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf
//go:noescape

func DOMMatrixRotateSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateSelf1
//go:noescape

func CallDOMMatrixRotateSelf1(
	this js.Ref,
	ptr unsafe.Pointer,

	rotX float64,
	rotY float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf1
//go:noescape

func DOMMatrixRotateSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateSelf2
//go:noescape

func CallDOMMatrixRotateSelf2(
	this js.Ref,
	ptr unsafe.Pointer,

	rotX float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf2
//go:noescape

func DOMMatrixRotateSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateSelf3
//go:noescape

func CallDOMMatrixRotateSelf3(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateSelf3
//go:noescape

func DOMMatrixRotateSelf3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateFromVectorSelf
//go:noescape

func CallDOMMatrixRotateFromVectorSelf(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateFromVectorSelf
//go:noescape

func DOMMatrixRotateFromVectorSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateFromVectorSelf1
//go:noescape

func CallDOMMatrixRotateFromVectorSelf1(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateFromVectorSelf1
//go:noescape

func DOMMatrixRotateFromVectorSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateFromVectorSelf2
//go:noescape

func CallDOMMatrixRotateFromVectorSelf2(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateFromVectorSelf2
//go:noescape

func DOMMatrixRotateFromVectorSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateAxisAngleSelf
//go:noescape

func CallDOMMatrixRotateAxisAngleSelf(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	z float64,
	angle float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf
//go:noescape

func DOMMatrixRotateAxisAngleSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateAxisAngleSelf1
//go:noescape

func CallDOMMatrixRotateAxisAngleSelf1(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
	z float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf1
//go:noescape

func DOMMatrixRotateAxisAngleSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateAxisAngleSelf2
//go:noescape

func CallDOMMatrixRotateAxisAngleSelf2(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf2
//go:noescape

func DOMMatrixRotateAxisAngleSelf2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateAxisAngleSelf3
//go:noescape

func CallDOMMatrixRotateAxisAngleSelf3(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf3
//go:noescape

func DOMMatrixRotateAxisAngleSelf3Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_RotateAxisAngleSelf4
//go:noescape

func CallDOMMatrixRotateAxisAngleSelf4(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_RotateAxisAngleSelf4
//go:noescape

func DOMMatrixRotateAxisAngleSelf4Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_SkewXSelf
//go:noescape

func CallDOMMatrixSkewXSelf(
	this js.Ref,
	ptr unsafe.Pointer,

	sx float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewXSelf
//go:noescape

func DOMMatrixSkewXSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_SkewXSelf1
//go:noescape

func CallDOMMatrixSkewXSelf1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewXSelf1
//go:noescape

func DOMMatrixSkewXSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_SkewYSelf
//go:noescape

func CallDOMMatrixSkewYSelf(
	this js.Ref,
	ptr unsafe.Pointer,

	sy float64,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewYSelf
//go:noescape

func DOMMatrixSkewYSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_SkewYSelf1
//go:noescape

func CallDOMMatrixSkewYSelf1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SkewYSelf1
//go:noescape

func DOMMatrixSkewYSelf1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_InvertSelf
//go:noescape

func CallDOMMatrixInvertSelf(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_InvertSelf
//go:noescape

func DOMMatrixInvertSelfFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMMatrix_SetMatrixValue
//go:noescape

func CallDOMMatrixSetMatrixValue(
	this js.Ref,
	ptr unsafe.Pointer,

	transformList js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMMatrix_SetMatrixValue
//go:noescape

func DOMMatrixSetMatrixValueFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_SVGTransform_Matrix
//go:noescape

func GetSVGTransformMatrix(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGTransform_Angle
//go:noescape

func GetSVGTransformAngle(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web call_SVGTransform_SetMatrix
//go:noescape

func CallSVGTransformSetMatrix(
	this js.Ref,
	ptr unsafe.Pointer,

	matrix unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetMatrix
//go:noescape

func SVGTransformSetMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetMatrix1
//go:noescape

func CallSVGTransformSetMatrix1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetMatrix1
//go:noescape

func SVGTransformSetMatrix1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetTranslate
//go:noescape

func CallSVGTransformSetTranslate(
	this js.Ref,
	ptr unsafe.Pointer,

	tx float32,
	ty float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetTranslate
//go:noescape

func SVGTransformSetTranslateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetScale
//go:noescape

func CallSVGTransformSetScale(
	this js.Ref,
	ptr unsafe.Pointer,

	sx float32,
	sy float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetScale
//go:noescape

func SVGTransformSetScaleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetRotate
//go:noescape

func CallSVGTransformSetRotate(
	this js.Ref,
	ptr unsafe.Pointer,

	angle float32,
	cx float32,
	cy float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetRotate
//go:noescape

func SVGTransformSetRotateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetSkewX
//go:noescape

func CallSVGTransformSetSkewX(
	this js.Ref,
	ptr unsafe.Pointer,

	angle float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetSkewX
//go:noescape

func SVGTransformSetSkewXFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGTransform_SetSkewY
//go:noescape

func CallSVGTransformSetSkewY(
	this js.Ref,
	ptr unsafe.Pointer,

	angle float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGTransform_SetSkewY
//go:noescape

func SVGTransformSetSkewYFunc(this js.Ref) js.Ref

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
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_DOMPointReadOnly_Y
//go:noescape

func GetDOMPointReadOnlyY(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_DOMPointReadOnly_Z
//go:noescape

func GetDOMPointReadOnlyZ(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web get_DOMPointReadOnly_W
//go:noescape

func GetDOMPointReadOnlyW(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web call_DOMPointReadOnly_FromPoint
//go:noescape

func CallDOMPointReadOnlyFromPoint(
	this js.Ref,
	ptr unsafe.Pointer,

	other unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_FromPoint
//go:noescape

func DOMPointReadOnlyFromPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMPointReadOnly_FromPoint1
//go:noescape

func CallDOMPointReadOnlyFromPoint1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_FromPoint1
//go:noescape

func DOMPointReadOnlyFromPoint1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMPointReadOnly_MatrixTransform
//go:noescape

func CallDOMPointReadOnlyMatrixTransform(
	this js.Ref,
	ptr unsafe.Pointer,

	matrix unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_MatrixTransform
//go:noescape

func DOMPointReadOnlyMatrixTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMPointReadOnly_MatrixTransform1
//go:noescape

func CallDOMPointReadOnlyMatrixTransform1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_MatrixTransform1
//go:noescape

func DOMPointReadOnlyMatrixTransform1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMPointReadOnly_ToJSON
//go:noescape

func CallDOMPointReadOnlyToJSON(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_DOMPointReadOnly_ToJSON
//go:noescape

func DOMPointReadOnlyToJSONFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedRect_BaseVal
//go:noescape

func GetSVGAnimatedRectBaseVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedRect_AnimVal
//go:noescape

func GetSVGAnimatedRectAnimVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGPreserveAspectRatio_Align
//go:noescape

func GetSVGPreserveAspectRatioAlign(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_SVGPreserveAspectRatio_Align
//go:noescape

func SetSVGPreserveAspectRatioAlign(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_SVGPreserveAspectRatio_MeetOrSlice
//go:noescape

func GetSVGPreserveAspectRatioMeetOrSlice(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web set_SVGPreserveAspectRatio_MeetOrSlice
//go:noescape

func SetSVGPreserveAspectRatioMeetOrSlice(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedPreserveAspectRatio_BaseVal
//go:noescape

func GetSVGAnimatedPreserveAspectRatioBaseVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGAnimatedPreserveAspectRatio_AnimVal
//go:noescape

func GetSVGAnimatedPreserveAspectRatioAnimVal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGSVGElement_X
//go:noescape

func GetSVGSVGElementX(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGSVGElement_Y
//go:noescape

func GetSVGSVGElementY(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGSVGElement_Width
//go:noescape

func GetSVGSVGElementWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGSVGElement_Height
//go:noescape

func GetSVGSVGElementHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGSVGElement_CurrentScale
//go:noescape

func GetSVGSVGElementCurrentScale(
	this js.Ref,
	ptr unsafe.Pointer,
) float32

//go:wasmimport plat/js/web set_SVGSVGElement_CurrentScale
//go:noescape

func SetSVGSVGElementCurrentScale(
	this js.Ref,
	val float32,
) js.Ref

//go:wasmimport plat/js/web get_SVGSVGElement_CurrentTranslate
//go:noescape

func GetSVGSVGElementCurrentTranslate(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGSVGElement_ViewBox
//go:noescape

func GetSVGSVGElementViewBox(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_SVGSVGElement_PreserveAspectRatio
//go:noescape

func GetSVGSVGElementPreserveAspectRatio(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_GetIntersectionList
//go:noescape

func CallSVGSVGElementGetIntersectionList(
	this js.Ref,
	ptr unsafe.Pointer,

	rect js.Ref,
	referenceElement js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_GetIntersectionList
//go:noescape

func SVGSVGElementGetIntersectionListFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_GetEnclosureList
//go:noescape

func CallSVGSVGElementGetEnclosureList(
	this js.Ref,
	ptr unsafe.Pointer,

	rect js.Ref,
	referenceElement js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_GetEnclosureList
//go:noescape

func SVGSVGElementGetEnclosureListFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CheckIntersection
//go:noescape

func CallSVGSVGElementCheckIntersection(
	this js.Ref,
	ptr unsafe.Pointer,

	element js.Ref,
	rect js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CheckIntersection
//go:noescape

func SVGSVGElementCheckIntersectionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CheckEnclosure
//go:noescape

func CallSVGSVGElementCheckEnclosure(
	this js.Ref,
	ptr unsafe.Pointer,

	element js.Ref,
	rect js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CheckEnclosure
//go:noescape

func SVGSVGElementCheckEnclosureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_DeselectAll
//go:noescape

func CallSVGSVGElementDeselectAll(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_DeselectAll
//go:noescape

func SVGSVGElementDeselectAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGNumber
//go:noescape

func CallSVGSVGElementCreateSVGNumber(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGNumber
//go:noescape

func SVGSVGElementCreateSVGNumberFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGLength
//go:noescape

func CallSVGSVGElementCreateSVGLength(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGLength
//go:noescape

func SVGSVGElementCreateSVGLengthFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGAngle
//go:noescape

func CallSVGSVGElementCreateSVGAngle(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGAngle
//go:noescape

func SVGSVGElementCreateSVGAngleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGPoint
//go:noescape

func CallSVGSVGElementCreateSVGPoint(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGPoint
//go:noescape

func SVGSVGElementCreateSVGPointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGMatrix
//go:noescape

func CallSVGSVGElementCreateSVGMatrix(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGMatrix
//go:noescape

func SVGSVGElementCreateSVGMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGRect
//go:noescape

func CallSVGSVGElementCreateSVGRect(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGRect
//go:noescape

func SVGSVGElementCreateSVGRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGTransform
//go:noescape

func CallSVGSVGElementCreateSVGTransform(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGTransform
//go:noescape

func SVGSVGElementCreateSVGTransformFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGTransformFromMatrix
//go:noescape

func CallSVGSVGElementCreateSVGTransformFromMatrix(
	this js.Ref,
	ptr unsafe.Pointer,

	matrix unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGTransformFromMatrix
//go:noescape

func SVGSVGElementCreateSVGTransformFromMatrixFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_CreateSVGTransformFromMatrix1
//go:noescape

func CallSVGSVGElementCreateSVGTransformFromMatrix1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_CreateSVGTransformFromMatrix1
//go:noescape

func SVGSVGElementCreateSVGTransformFromMatrix1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_GetElementById
//go:noescape

func CallSVGSVGElementGetElementById(
	this js.Ref,
	ptr unsafe.Pointer,

	elementId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_GetElementById
//go:noescape

func SVGSVGElementGetElementByIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_SuspendRedraw
//go:noescape

func CallSVGSVGElementSuspendRedraw(
	this js.Ref,
	ptr unsafe.Pointer,

	maxWaitMilliseconds uint32,
) uint32

//go:wasmimport plat/js/web func_SVGSVGElement_SuspendRedraw
//go:noescape

func SVGSVGElementSuspendRedrawFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_UnsuspendRedraw
//go:noescape

func CallSVGSVGElementUnsuspendRedraw(
	this js.Ref,
	ptr unsafe.Pointer,

	suspendHandleID uint32,
) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_UnsuspendRedraw
//go:noescape

func SVGSVGElementUnsuspendRedrawFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_UnsuspendRedrawAll
//go:noescape

func CallSVGSVGElementUnsuspendRedrawAll(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_UnsuspendRedrawAll
//go:noescape

func SVGSVGElementUnsuspendRedrawAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_ForceRedraw
//go:noescape

func CallSVGSVGElementForceRedraw(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_ForceRedraw
//go:noescape

func SVGSVGElementForceRedrawFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_PauseAnimations
//go:noescape

func CallSVGSVGElementPauseAnimations(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_PauseAnimations
//go:noescape

func SVGSVGElementPauseAnimationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_UnpauseAnimations
//go:noescape

func CallSVGSVGElementUnpauseAnimations(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_UnpauseAnimations
//go:noescape

func SVGSVGElementUnpauseAnimationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_AnimationsPaused
//go:noescape

func CallSVGSVGElementAnimationsPaused(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_AnimationsPaused
//go:noescape

func SVGSVGElementAnimationsPausedFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_GetCurrentTime
//go:noescape

func CallSVGSVGElementGetCurrentTime(
	this js.Ref,
	ptr unsafe.Pointer,

) float32

//go:wasmimport plat/js/web func_SVGSVGElement_GetCurrentTime
//go:noescape

func SVGSVGElementGetCurrentTimeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_SVGSVGElement_SetCurrentTime
//go:noescape

func CallSVGSVGElementSetCurrentTime(
	this js.Ref,
	ptr unsafe.Pointer,

	seconds float32,
) js.Ref

//go:wasmimport plat/js/web func_SVGSVGElement_SetCurrentTime
//go:noescape

func SVGSVGElementSetCurrentTimeFunc(this js.Ref) js.Ref

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

//go:wasmimport plat/js/web new_HTMLSlotElement_HTMLSlotElement
//go:noescape

func NewHTMLSlotElementByHTMLSlotElement() js.Ref

//go:wasmimport plat/js/web get_HTMLSlotElement_Name
//go:noescape

func GetHTMLSlotElementName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_HTMLSlotElement_Name
//go:noescape

func SetHTMLSlotElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_HTMLSlotElement_AssignedNodes
//go:noescape

func CallHTMLSlotElementAssignedNodes(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_HTMLSlotElement_AssignedNodes
//go:noescape

func HTMLSlotElementAssignedNodesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSlotElement_AssignedNodes1
//go:noescape

func CallHTMLSlotElementAssignedNodes1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLSlotElement_AssignedNodes1
//go:noescape

func HTMLSlotElementAssignedNodes1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSlotElement_AssignedElements
//go:noescape

func CallHTMLSlotElementAssignedElements(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_HTMLSlotElement_AssignedElements
//go:noescape

func HTMLSlotElementAssignedElementsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSlotElement_AssignedElements1
//go:noescape

func CallHTMLSlotElementAssignedElements1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_HTMLSlotElement_AssignedElements1
//go:noescape

func HTMLSlotElementAssignedElements1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSlotElement_Assign
//go:noescape

func CallHTMLSlotElementAssign(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_HTMLSlotElement_Assign
//go:noescape

func HTMLSlotElementAssignFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_Text_Text
//go:noescape

func NewTextByText(
	data js.Ref) js.Ref

//go:wasmimport plat/js/web new_Text_Text1
//go:noescape

func NewTextByText1() js.Ref

//go:wasmimport plat/js/web get_Text_WholeText
//go:noescape

func GetTextWholeText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Text_AssignedSlot
//go:noescape

func GetTextAssignedSlot(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Text_SplitText
//go:noescape

func CallTextSplitText(
	this js.Ref,
	ptr unsafe.Pointer,

	offset uint32,
) js.Ref

//go:wasmimport plat/js/web func_Text_SplitText
//go:noescape

func TextSplitTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Text_GetBoxQuads
//go:noescape

func CallTextGetBoxQuads(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Text_GetBoxQuads
//go:noescape

func TextGetBoxQuadsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Text_GetBoxQuads1
//go:noescape

func CallTextGetBoxQuads1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Text_GetBoxQuads1
//go:noescape

func TextGetBoxQuads1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Text_ConvertQuadFromNode
//go:noescape

func CallTextConvertQuadFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	quad unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Text_ConvertQuadFromNode
//go:noescape

func TextConvertQuadFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Text_ConvertQuadFromNode1
//go:noescape

func CallTextConvertQuadFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	quad unsafe.Pointer,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Text_ConvertQuadFromNode1
//go:noescape

func TextConvertQuadFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Text_ConvertRectFromNode
//go:noescape

func CallTextConvertRectFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	rect js.Ref,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Text_ConvertRectFromNode
//go:noescape

func TextConvertRectFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Text_ConvertRectFromNode1
//go:noescape

func CallTextConvertRectFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	rect js.Ref,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Text_ConvertRectFromNode1
//go:noescape

func TextConvertRectFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Text_ConvertPointFromNode
//go:noescape

func CallTextConvertPointFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	point unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Text_ConvertPointFromNode
//go:noescape

func TextConvertPointFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Text_ConvertPointFromNode1
//go:noescape

func CallTextConvertPointFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	point unsafe.Pointer,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Text_ConvertPointFromNode1
//go:noescape

func TextConvertPointFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_BoxQuadOptions
//go:noescape

func BoxQuadOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_BoxQuadOptions
//go:noescape

func BoxQuadOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSPseudoElement_Type
//go:noescape

func GetCSSPseudoElementType(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSSPseudoElement_Element
//go:noescape

func GetCSSPseudoElementElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_CSSPseudoElement_Parent
//go:noescape

func GetCSSPseudoElementParent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_CSSPseudoElement_Pseudo
//go:noescape

func CallCSSPseudoElementPseudo(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSPseudoElement_Pseudo
//go:noescape

func CSSPseudoElementPseudoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSPseudoElement_GetBoxQuads
//go:noescape

func CallCSSPseudoElementGetBoxQuads(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CSSPseudoElement_GetBoxQuads
//go:noescape

func CSSPseudoElementGetBoxQuadsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSPseudoElement_GetBoxQuads1
//go:noescape

func CallCSSPseudoElementGetBoxQuads1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_CSSPseudoElement_GetBoxQuads1
//go:noescape

func CSSPseudoElementGetBoxQuads1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSPseudoElement_ConvertQuadFromNode
//go:noescape

func CallCSSPseudoElementConvertQuadFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	quad unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CSSPseudoElement_ConvertQuadFromNode
//go:noescape

func CSSPseudoElementConvertQuadFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSPseudoElement_ConvertQuadFromNode1
//go:noescape

func CallCSSPseudoElementConvertQuadFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	quad unsafe.Pointer,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSPseudoElement_ConvertQuadFromNode1
//go:noescape

func CSSPseudoElementConvertQuadFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSPseudoElement_ConvertRectFromNode
//go:noescape

func CallCSSPseudoElementConvertRectFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	rect js.Ref,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CSSPseudoElement_ConvertRectFromNode
//go:noescape

func CSSPseudoElementConvertRectFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSPseudoElement_ConvertRectFromNode1
//go:noescape

func CallCSSPseudoElementConvertRectFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	rect js.Ref,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSPseudoElement_ConvertRectFromNode1
//go:noescape

func CSSPseudoElementConvertRectFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSPseudoElement_ConvertPointFromNode
//go:noescape

func CallCSSPseudoElementConvertPointFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	point unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_CSSPseudoElement_ConvertPointFromNode
//go:noescape

func CSSPseudoElementConvertPointFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_CSSPseudoElement_ConvertPointFromNode1
//go:noescape

func CallCSSPseudoElementConvertPointFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	point unsafe.Pointer,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_CSSPseudoElement_ConvertPointFromNode1
//go:noescape

func CSSPseudoElementConvertPointFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_SanitizerConfig
//go:noescape

func SanitizerConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SanitizerConfig
//go:noescape

func SanitizerConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_NodeList_Length
//go:noescape

func GetNodeListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_NodeList_Item
//go:noescape

func CallNodeListItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_NodeList_Item
//go:noescape

func NodeListItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_DocumentFragment_Children
//go:noescape

func GetDocumentFragmentChildren(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DocumentFragment_FirstElementChild
//go:noescape

func GetDocumentFragmentFirstElementChild(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DocumentFragment_LastElementChild
//go:noescape

func GetDocumentFragmentLastElementChild(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_DocumentFragment_ChildElementCount
//go:noescape

func GetDocumentFragmentChildElementCount(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_DocumentFragment_GetElementById
//go:noescape

func CallDocumentFragmentGetElementById(
	this js.Ref,
	ptr unsafe.Pointer,

	elementId js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DocumentFragment_GetElementById
//go:noescape

func DocumentFragmentGetElementByIdFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentFragment_Prepend
//go:noescape

func CallDocumentFragmentPrepend(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_DocumentFragment_Prepend
//go:noescape

func DocumentFragmentPrependFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentFragment_Append
//go:noescape

func CallDocumentFragmentAppend(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_DocumentFragment_Append
//go:noescape

func DocumentFragmentAppendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentFragment_ReplaceChildren
//go:noescape

func CallDocumentFragmentReplaceChildren(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_DocumentFragment_ReplaceChildren
//go:noescape

func DocumentFragmentReplaceChildrenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentFragment_QuerySelector
//go:noescape

func CallDocumentFragmentQuerySelector(
	this js.Ref,
	ptr unsafe.Pointer,

	selectors js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DocumentFragment_QuerySelector
//go:noescape

func DocumentFragmentQuerySelectorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DocumentFragment_QuerySelectorAll
//go:noescape

func CallDocumentFragmentQuerySelectorAll(
	this js.Ref,
	ptr unsafe.Pointer,

	selectors js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DocumentFragment_QuerySelectorAll
//go:noescape

func DocumentFragmentQuerySelectorAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web new_Sanitizer_Sanitizer
//go:noescape

func NewSanitizerBySanitizer(
	config unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_Sanitizer_Sanitizer1
//go:noescape

func NewSanitizerBySanitizer1() js.Ref

//go:wasmimport plat/js/web call_Sanitizer_Sanitize
//go:noescape

func CallSanitizerSanitize(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Sanitizer_Sanitize
//go:noescape

func SanitizerSanitizeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Sanitizer_SanitizeFor
//go:noescape

func CallSanitizerSanitizeFor(
	this js.Ref,
	ptr unsafe.Pointer,

	element js.Ref,
	input js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Sanitizer_SanitizeFor
//go:noescape

func SanitizerSanitizeForFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Sanitizer_GetConfiguration
//go:noescape

func CallSanitizerGetConfiguration(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Sanitizer_GetConfiguration
//go:noescape

func SanitizerGetConfigurationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Sanitizer_GetDefaultConfiguration
//go:noescape

func CallSanitizerGetDefaultConfiguration(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Sanitizer_GetDefaultConfiguration
//go:noescape

func SanitizerGetDefaultConfigurationFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_SetHTMLOptions
//go:noescape

func SetHTMLOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_SetHTMLOptions
//go:noescape

func SetHTMLOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_DOMRectList_Length
//go:noescape

func GetDOMRectListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_DOMRectList_Item
//go:noescape

func CallDOMRectListItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_DOMRectList_Item
//go:noescape

func DOMRectListItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_CheckVisibilityOptions
//go:noescape

func CheckVisibilityOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_CheckVisibilityOptions
//go:noescape

func CheckVisibilityOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_ScrollLogicalPosition
//go:noescape

func ConstOfScrollLogicalPosition(str js.Ref) uint32

//go:wasmimport plat/js/web constof_ScrollBehavior
//go:noescape

func ConstOfScrollBehavior(str js.Ref) uint32

//go:wasmimport plat/js/web store_ScrollIntoViewOptions
//go:noescape

func ScrollIntoViewOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ScrollIntoViewOptions
//go:noescape

func ScrollIntoViewOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_ScrollToOptions
//go:noescape

func ScrollToOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ScrollToOptions
//go:noescape

func ScrollToOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_IterationCompositeOperation
//go:noescape

func ConstOfIterationCompositeOperation(str js.Ref) uint32

//go:wasmimport plat/js/web store_TimelineRangeOffset
//go:noescape

func TimelineRangeOffsetJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_TimelineRangeOffset
//go:noescape

func TimelineRangeOffsetJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CSSKeywordValue_CSSKeywordValue
//go:noescape

func NewCSSKeywordValueByCSSKeywordValue(
	value js.Ref) js.Ref

//go:wasmimport plat/js/web get_CSSKeywordValue_Value
//go:noescape

func GetCSSKeywordValueValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_CSSKeywordValue_Value
//go:noescape

func SetCSSKeywordValueValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_KeyframeAnimationOptions
//go:noescape

func KeyframeAnimationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_KeyframeAnimationOptions
//go:noescape

func KeyframeAnimationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_GetAnimationsOptions
//go:noescape

func GetAnimationsOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_GetAnimationsOptions
//go:noescape

func GetAnimationsOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_Range_CommonAncestorContainer
//go:noescape

func GetRangeCommonAncestorContainer(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Range_SetStart
//go:noescape

func CallRangeSetStart(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
	offset uint32,
) js.Ref

//go:wasmimport plat/js/web func_Range_SetStart
//go:noescape

func RangeSetStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_SetEnd
//go:noescape

func CallRangeSetEnd(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
	offset uint32,
) js.Ref

//go:wasmimport plat/js/web func_Range_SetEnd
//go:noescape

func RangeSetEndFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_SetStartBefore
//go:noescape

func CallRangeSetStartBefore(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Range_SetStartBefore
//go:noescape

func RangeSetStartBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_SetStartAfter
//go:noescape

func CallRangeSetStartAfter(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Range_SetStartAfter
//go:noescape

func RangeSetStartAfterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_SetEndBefore
//go:noescape

func CallRangeSetEndBefore(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Range_SetEndBefore
//go:noescape

func RangeSetEndBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_SetEndAfter
//go:noescape

func CallRangeSetEndAfter(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Range_SetEndAfter
//go:noescape

func RangeSetEndAfterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_Collapse
//go:noescape

func CallRangeCollapse(
	this js.Ref,
	ptr unsafe.Pointer,

	toStart js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Range_Collapse
//go:noescape

func RangeCollapseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_Collapse1
//go:noescape

func CallRangeCollapse1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Range_Collapse1
//go:noescape

func RangeCollapse1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_SelectNode
//go:noescape

func CallRangeSelectNode(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Range_SelectNode
//go:noescape

func RangeSelectNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_SelectNodeContents
//go:noescape

func CallRangeSelectNodeContents(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Range_SelectNodeContents
//go:noescape

func RangeSelectNodeContentsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_CompareBoundaryPoints
//go:noescape

func CallRangeCompareBoundaryPoints(
	this js.Ref,
	ptr unsafe.Pointer,

	how uint32,
	sourceRange js.Ref,
) int32

//go:wasmimport plat/js/web func_Range_CompareBoundaryPoints
//go:noescape

func RangeCompareBoundaryPointsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_DeleteContents
//go:noescape

func CallRangeDeleteContents(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Range_DeleteContents
//go:noescape

func RangeDeleteContentsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_ExtractContents
//go:noescape

func CallRangeExtractContents(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Range_ExtractContents
//go:noescape

func RangeExtractContentsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_CloneContents
//go:noescape

func CallRangeCloneContents(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Range_CloneContents
//go:noescape

func RangeCloneContentsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_InsertNode
//go:noescape

func CallRangeInsertNode(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Range_InsertNode
//go:noescape

func RangeInsertNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_SurroundContents
//go:noescape

func CallRangeSurroundContents(
	this js.Ref,
	ptr unsafe.Pointer,

	newParent js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Range_SurroundContents
//go:noescape

func RangeSurroundContentsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_CloneRange
//go:noescape

func CallRangeCloneRange(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Range_CloneRange
//go:noescape

func RangeCloneRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_Detach
//go:noescape

func CallRangeDetach(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Range_Detach
//go:noescape

func RangeDetachFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_IsPointInRange
//go:noescape

func CallRangeIsPointInRange(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
	offset uint32,
) js.Ref

//go:wasmimport plat/js/web func_Range_IsPointInRange
//go:noescape

func RangeIsPointInRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_ComparePoint
//go:noescape

func CallRangeComparePoint(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
	offset uint32,
) int32

//go:wasmimport plat/js/web func_Range_ComparePoint
//go:noescape

func RangeComparePointFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_IntersectsNode
//go:noescape

func CallRangeIntersectsNode(
	this js.Ref,
	ptr unsafe.Pointer,

	node js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Range_IntersectsNode
//go:noescape

func RangeIntersectsNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_ToString
//go:noescape

func CallRangeToString(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Range_ToString
//go:noescape

func RangeToStringFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_GetClientRects
//go:noescape

func CallRangeGetClientRects(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Range_GetClientRects
//go:noescape

func RangeGetClientRectsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_GetBoundingClientRect
//go:noescape

func CallRangeGetBoundingClientRect(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Range_GetBoundingClientRect
//go:noescape

func RangeGetBoundingClientRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Range_CreateContextualFragment
//go:noescape

func CallRangeCreateContextualFragment(
	this js.Ref,
	ptr unsafe.Pointer,

	fragment js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Range_CreateContextualFragment
//go:noescape

func RangeCreateContextualFragmentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_DOMTokenList_Length
//go:noescape

func GetDOMTokenListLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_DOMTokenList_Value
//go:noescape

func GetDOMTokenListValue(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_DOMTokenList_Value
//go:noescape

func SetDOMTokenListValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_DOMTokenList_Item
//go:noescape

func CallDOMTokenListItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_DOMTokenList_Item
//go:noescape

func DOMTokenListItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMTokenList_Contains
//go:noescape

func CallDOMTokenListContains(
	this js.Ref,
	ptr unsafe.Pointer,

	token js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMTokenList_Contains
//go:noescape

func DOMTokenListContainsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMTokenList_Add
//go:noescape

func CallDOMTokenListAdd(
	this js.Ref,
	ptr unsafe.Pointer,

	tokensPtr unsafe.Pointer,
	tokensCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_DOMTokenList_Add
//go:noescape

func DOMTokenListAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMTokenList_Remove
//go:noescape

func CallDOMTokenListRemove(
	this js.Ref,
	ptr unsafe.Pointer,

	tokensPtr unsafe.Pointer,
	tokensCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_DOMTokenList_Remove
//go:noescape

func DOMTokenListRemoveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMTokenList_Toggle
//go:noescape

func CallDOMTokenListToggle(
	this js.Ref,
	ptr unsafe.Pointer,

	token js.Ref,
	force js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMTokenList_Toggle
//go:noescape

func DOMTokenListToggleFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMTokenList_Toggle1
//go:noescape

func CallDOMTokenListToggle1(
	this js.Ref,
	ptr unsafe.Pointer,

	token js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMTokenList_Toggle1
//go:noescape

func DOMTokenListToggle1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMTokenList_Replace
//go:noescape

func CallDOMTokenListReplace(
	this js.Ref,
	ptr unsafe.Pointer,

	token js.Ref,
	newToken js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMTokenList_Replace
//go:noescape

func DOMTokenListReplaceFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_DOMTokenList_Supports
//go:noescape

func CallDOMTokenListSupports(
	this js.Ref,
	ptr unsafe.Pointer,

	token js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_DOMTokenList_Supports
//go:noescape

func DOMTokenListSupportsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_NamedNodeMap_Length
//go:noescape

func GetNamedNodeMapLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_NamedNodeMap_Item
//go:noescape

func CallNamedNodeMapItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_NamedNodeMap_Item
//go:noescape

func NamedNodeMapItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NamedNodeMap_GetNamedItem
//go:noescape

func CallNamedNodeMapGetNamedItem(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_NamedNodeMap_GetNamedItem
//go:noescape

func NamedNodeMapGetNamedItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NamedNodeMap_GetNamedItemNS
//go:noescape

func CallNamedNodeMapGetNamedItemNS(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_NamedNodeMap_GetNamedItemNS
//go:noescape

func NamedNodeMapGetNamedItemNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NamedNodeMap_SetNamedItem
//go:noescape

func CallNamedNodeMapSetNamedItem(
	this js.Ref,
	ptr unsafe.Pointer,

	attr js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_NamedNodeMap_SetNamedItem
//go:noescape

func NamedNodeMapSetNamedItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NamedNodeMap_SetNamedItemNS
//go:noescape

func CallNamedNodeMapSetNamedItemNS(
	this js.Ref,
	ptr unsafe.Pointer,

	attr js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_NamedNodeMap_SetNamedItemNS
//go:noescape

func NamedNodeMapSetNamedItemNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NamedNodeMap_RemoveNamedItem
//go:noescape

func CallNamedNodeMapRemoveNamedItem(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_NamedNodeMap_RemoveNamedItem
//go:noescape

func NamedNodeMapRemoveNamedItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NamedNodeMap_RemoveNamedItemNS
//go:noescape

func CallNamedNodeMapRemoveNamedItemNS(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_NamedNodeMap_RemoveNamedItemNS
//go:noescape

func NamedNodeMapRemoveNamedItemNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_Element_NamespaceURI
//go:noescape

func GetElementNamespaceURI(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_Prefix
//go:noescape

func GetElementPrefix(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_LocalName
//go:noescape

func GetElementLocalName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_TagName
//go:noescape

func GetElementTagName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_Id
//go:noescape

func GetElementId(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_Id
//go:noescape

func SetElementId(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_ClassName
//go:noescape

func GetElementClassName(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_ClassName
//go:noescape

func SetElementClassName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_ClassList
//go:noescape

func GetElementClassList(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_Slot
//go:noescape

func GetElementSlot(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_Slot
//go:noescape

func SetElementSlot(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_Attributes
//go:noescape

func GetElementAttributes(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_ShadowRoot
//go:noescape

func GetElementShadowRoot(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_ElementTiming
//go:noescape

func GetElementElementTiming(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_ElementTiming
//go:noescape

func SetElementElementTiming(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_Part
//go:noescape

func GetElementPart(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_OuterHTML
//go:noescape

func GetElementOuterHTML(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_OuterHTML
//go:noescape

func SetElementOuterHTML(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_ScrollTop
//go:noescape

func GetElementScrollTop(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_Element_ScrollTop
//go:noescape

func SetElementScrollTop(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_Element_ScrollLeft
//go:noescape

func GetElementScrollLeft(
	this js.Ref,
	ptr unsafe.Pointer,
) float64

//go:wasmimport plat/js/web set_Element_ScrollLeft
//go:noescape

func SetElementScrollLeft(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_Element_ScrollWidth
//go:noescape

func GetElementScrollWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_Element_ScrollHeight
//go:noescape

func GetElementScrollHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_Element_ClientTop
//go:noescape

func GetElementClientTop(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_Element_ClientLeft
//go:noescape

func GetElementClientLeft(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_Element_ClientWidth
//go:noescape

func GetElementClientWidth(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_Element_ClientHeight
//go:noescape

func GetElementClientHeight(
	this js.Ref,
	ptr unsafe.Pointer,
) int32

//go:wasmimport plat/js/web get_Element_Role
//go:noescape

func GetElementRole(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_Role
//go:noescape

func SetElementRole(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaActiveDescendantElement
//go:noescape

func GetElementAriaActiveDescendantElement(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaActiveDescendantElement
//go:noescape

func SetElementAriaActiveDescendantElement(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaAtomic
//go:noescape

func GetElementAriaAtomic(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaAtomic
//go:noescape

func SetElementAriaAtomic(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaAutoComplete
//go:noescape

func GetElementAriaAutoComplete(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaAutoComplete
//go:noescape

func SetElementAriaAutoComplete(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaBusy
//go:noescape

func GetElementAriaBusy(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaBusy
//go:noescape

func SetElementAriaBusy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaChecked
//go:noescape

func GetElementAriaChecked(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaChecked
//go:noescape

func SetElementAriaChecked(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaColCount
//go:noescape

func GetElementAriaColCount(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaColCount
//go:noescape

func SetElementAriaColCount(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaColIndex
//go:noescape

func GetElementAriaColIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaColIndex
//go:noescape

func SetElementAriaColIndex(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaColIndexText
//go:noescape

func GetElementAriaColIndexText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaColIndexText
//go:noescape

func SetElementAriaColIndexText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaColSpan
//go:noescape

func GetElementAriaColSpan(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaColSpan
//go:noescape

func SetElementAriaColSpan(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaControlsElements
//go:noescape

func GetElementAriaControlsElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaControlsElements
//go:noescape

func SetElementAriaControlsElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaCurrent
//go:noescape

func GetElementAriaCurrent(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaCurrent
//go:noescape

func SetElementAriaCurrent(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaDescribedByElements
//go:noescape

func GetElementAriaDescribedByElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaDescribedByElements
//go:noescape

func SetElementAriaDescribedByElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaDescription
//go:noescape

func GetElementAriaDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaDescription
//go:noescape

func SetElementAriaDescription(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaDetailsElements
//go:noescape

func GetElementAriaDetailsElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaDetailsElements
//go:noescape

func SetElementAriaDetailsElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaDisabled
//go:noescape

func GetElementAriaDisabled(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaDisabled
//go:noescape

func SetElementAriaDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaErrorMessageElements
//go:noescape

func GetElementAriaErrorMessageElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaErrorMessageElements
//go:noescape

func SetElementAriaErrorMessageElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaExpanded
//go:noescape

func GetElementAriaExpanded(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaExpanded
//go:noescape

func SetElementAriaExpanded(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaFlowToElements
//go:noescape

func GetElementAriaFlowToElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaFlowToElements
//go:noescape

func SetElementAriaFlowToElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaHasPopup
//go:noescape

func GetElementAriaHasPopup(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaHasPopup
//go:noescape

func SetElementAriaHasPopup(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaHidden
//go:noescape

func GetElementAriaHidden(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaHidden
//go:noescape

func SetElementAriaHidden(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaInvalid
//go:noescape

func GetElementAriaInvalid(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaInvalid
//go:noescape

func SetElementAriaInvalid(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaKeyShortcuts
//go:noescape

func GetElementAriaKeyShortcuts(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaKeyShortcuts
//go:noescape

func SetElementAriaKeyShortcuts(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaLabel
//go:noescape

func GetElementAriaLabel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaLabel
//go:noescape

func SetElementAriaLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaLabelledByElements
//go:noescape

func GetElementAriaLabelledByElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaLabelledByElements
//go:noescape

func SetElementAriaLabelledByElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaLevel
//go:noescape

func GetElementAriaLevel(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaLevel
//go:noescape

func SetElementAriaLevel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaLive
//go:noescape

func GetElementAriaLive(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaLive
//go:noescape

func SetElementAriaLive(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaModal
//go:noescape

func GetElementAriaModal(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaModal
//go:noescape

func SetElementAriaModal(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaMultiLine
//go:noescape

func GetElementAriaMultiLine(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaMultiLine
//go:noescape

func SetElementAriaMultiLine(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaMultiSelectable
//go:noescape

func GetElementAriaMultiSelectable(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaMultiSelectable
//go:noescape

func SetElementAriaMultiSelectable(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaOrientation
//go:noescape

func GetElementAriaOrientation(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaOrientation
//go:noescape

func SetElementAriaOrientation(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaOwnsElements
//go:noescape

func GetElementAriaOwnsElements(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaOwnsElements
//go:noescape

func SetElementAriaOwnsElements(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaPlaceholder
//go:noescape

func GetElementAriaPlaceholder(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaPlaceholder
//go:noescape

func SetElementAriaPlaceholder(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaPosInSet
//go:noescape

func GetElementAriaPosInSet(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaPosInSet
//go:noescape

func SetElementAriaPosInSet(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaPressed
//go:noescape

func GetElementAriaPressed(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaPressed
//go:noescape

func SetElementAriaPressed(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaReadOnly
//go:noescape

func GetElementAriaReadOnly(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaReadOnly
//go:noescape

func SetElementAriaReadOnly(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaRequired
//go:noescape

func GetElementAriaRequired(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaRequired
//go:noescape

func SetElementAriaRequired(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaRoleDescription
//go:noescape

func GetElementAriaRoleDescription(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaRoleDescription
//go:noescape

func SetElementAriaRoleDescription(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaRowCount
//go:noescape

func GetElementAriaRowCount(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaRowCount
//go:noescape

func SetElementAriaRowCount(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaRowIndex
//go:noescape

func GetElementAriaRowIndex(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaRowIndex
//go:noescape

func SetElementAriaRowIndex(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaRowIndexText
//go:noescape

func GetElementAriaRowIndexText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaRowIndexText
//go:noescape

func SetElementAriaRowIndexText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaRowSpan
//go:noescape

func GetElementAriaRowSpan(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaRowSpan
//go:noescape

func SetElementAriaRowSpan(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaSelected
//go:noescape

func GetElementAriaSelected(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaSelected
//go:noescape

func SetElementAriaSelected(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaSetSize
//go:noescape

func GetElementAriaSetSize(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaSetSize
//go:noescape

func SetElementAriaSetSize(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaSort
//go:noescape

func GetElementAriaSort(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaSort
//go:noescape

func SetElementAriaSort(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaValueMax
//go:noescape

func GetElementAriaValueMax(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaValueMax
//go:noescape

func SetElementAriaValueMax(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaValueMin
//go:noescape

func GetElementAriaValueMin(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaValueMin
//go:noescape

func SetElementAriaValueMin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaValueNow
//go:noescape

func GetElementAriaValueNow(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaValueNow
//go:noescape

func SetElementAriaValueNow(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_AriaValueText
//go:noescape

func GetElementAriaValueText(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_AriaValueText
//go:noescape

func SetElementAriaValueText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_InnerHTML
//go:noescape

func GetElementInnerHTML(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_Element_InnerHTML
//go:noescape

func SetElementInnerHTML(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_Element_Children
//go:noescape

func GetElementChildren(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_FirstElementChild
//go:noescape

func GetElementFirstElementChild(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_LastElementChild
//go:noescape

func GetElementLastElementChild(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_ChildElementCount
//go:noescape

func GetElementChildElementCount(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_Element_PreviousElementSibling
//go:noescape

func GetElementPreviousElementSibling(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_NextElementSibling
//go:noescape

func GetElementNextElementSibling(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_AssignedSlot
//go:noescape

func GetElementAssignedSlot(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_Element_RegionOverset
//go:noescape

func GetElementRegionOverset(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_Element_HasAttributes
//go:noescape

func CallElementHasAttributes(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_HasAttributes
//go:noescape

func ElementHasAttributesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetAttributeNames
//go:noescape

func CallElementGetAttributeNames(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_GetAttributeNames
//go:noescape

func ElementGetAttributeNamesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetAttribute
//go:noescape

func CallElementGetAttribute(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_GetAttribute
//go:noescape

func ElementGetAttributeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetAttributeNS
//go:noescape

func CallElementGetAttributeNS(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_GetAttributeNS
//go:noescape

func ElementGetAttributeNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_SetAttribute
//go:noescape

func CallElementSetAttribute(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_SetAttribute
//go:noescape

func ElementSetAttributeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_SetAttributeNS
//go:noescape

func CallElementSetAttributeNS(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	qualifiedName js.Ref,
	value js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_SetAttributeNS
//go:noescape

func ElementSetAttributeNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_RemoveAttribute
//go:noescape

func CallElementRemoveAttribute(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_RemoveAttribute
//go:noescape

func ElementRemoveAttributeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_RemoveAttributeNS
//go:noescape

func CallElementRemoveAttributeNS(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_RemoveAttributeNS
//go:noescape

func ElementRemoveAttributeNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ToggleAttribute
//go:noescape

func CallElementToggleAttribute(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
	force js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_ToggleAttribute
//go:noescape

func ElementToggleAttributeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ToggleAttribute1
//go:noescape

func CallElementToggleAttribute1(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_ToggleAttribute1
//go:noescape

func ElementToggleAttribute1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_HasAttribute
//go:noescape

func CallElementHasAttribute(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_HasAttribute
//go:noescape

func ElementHasAttributeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_HasAttributeNS
//go:noescape

func CallElementHasAttributeNS(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_HasAttributeNS
//go:noescape

func ElementHasAttributeNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetAttributeNode
//go:noescape

func CallElementGetAttributeNode(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_GetAttributeNode
//go:noescape

func ElementGetAttributeNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetAttributeNodeNS
//go:noescape

func CallElementGetAttributeNodeNS(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_GetAttributeNodeNS
//go:noescape

func ElementGetAttributeNodeNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_SetAttributeNode
//go:noescape

func CallElementSetAttributeNode(
	this js.Ref,
	ptr unsafe.Pointer,

	attr js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_SetAttributeNode
//go:noescape

func ElementSetAttributeNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_SetAttributeNodeNS
//go:noescape

func CallElementSetAttributeNodeNS(
	this js.Ref,
	ptr unsafe.Pointer,

	attr js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_SetAttributeNodeNS
//go:noescape

func ElementSetAttributeNodeNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_RemoveAttributeNode
//go:noescape

func CallElementRemoveAttributeNode(
	this js.Ref,
	ptr unsafe.Pointer,

	attr js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_RemoveAttributeNode
//go:noescape

func ElementRemoveAttributeNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_AttachShadow
//go:noescape

func CallElementAttachShadow(
	this js.Ref,
	ptr unsafe.Pointer,

	init unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_AttachShadow
//go:noescape

func ElementAttachShadowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Closest
//go:noescape

func CallElementClosest(
	this js.Ref,
	ptr unsafe.Pointer,

	selectors js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_Closest
//go:noescape

func ElementClosestFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Matches
//go:noescape

func CallElementMatches(
	this js.Ref,
	ptr unsafe.Pointer,

	selectors js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_Matches
//go:noescape

func ElementMatchesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_WebkitMatchesSelector
//go:noescape

func CallElementWebkitMatchesSelector(
	this js.Ref,
	ptr unsafe.Pointer,

	selectors js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_WebkitMatchesSelector
//go:noescape

func ElementWebkitMatchesSelectorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetElementsByTagName
//go:noescape

func CallElementGetElementsByTagName(
	this js.Ref,
	ptr unsafe.Pointer,

	qualifiedName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_GetElementsByTagName
//go:noescape

func ElementGetElementsByTagNameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetElementsByTagNameNS
//go:noescape

func CallElementGetElementsByTagNameNS(
	this js.Ref,
	ptr unsafe.Pointer,

	namespace js.Ref,
	localName js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_GetElementsByTagNameNS
//go:noescape

func ElementGetElementsByTagNameNSFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetElementsByClassName
//go:noescape

func CallElementGetElementsByClassName(
	this js.Ref,
	ptr unsafe.Pointer,

	classNames js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_GetElementsByClassName
//go:noescape

func ElementGetElementsByClassNameFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_InsertAdjacentElement
//go:noescape

func CallElementInsertAdjacentElement(
	this js.Ref,
	ptr unsafe.Pointer,

	where js.Ref,
	element js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_InsertAdjacentElement
//go:noescape

func ElementInsertAdjacentElementFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_InsertAdjacentText
//go:noescape

func CallElementInsertAdjacentText(
	this js.Ref,
	ptr unsafe.Pointer,

	where js.Ref,
	data js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_InsertAdjacentText
//go:noescape

func ElementInsertAdjacentTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_RequestFullscreen
//go:noescape

func CallElementRequestFullscreen(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_RequestFullscreen
//go:noescape

func ElementRequestFullscreenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_RequestFullscreen1
//go:noescape

func CallElementRequestFullscreen1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_RequestFullscreen1
//go:noescape

func ElementRequestFullscreen1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_RequestPointerLock
//go:noescape

func CallElementRequestPointerLock(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_RequestPointerLock
//go:noescape

func ElementRequestPointerLockFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ComputedStyleMap
//go:noescape

func CallElementComputedStyleMap(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_ComputedStyleMap
//go:noescape

func ElementComputedStyleMapFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetSpatialNavigationContainer
//go:noescape

func CallElementGetSpatialNavigationContainer(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_GetSpatialNavigationContainer
//go:noescape

func ElementGetSpatialNavigationContainerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_FocusableAreas
//go:noescape

func CallElementFocusableAreas(
	this js.Ref,
	ptr unsafe.Pointer,

	option unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_FocusableAreas
//go:noescape

func ElementFocusableAreasFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_FocusableAreas1
//go:noescape

func CallElementFocusableAreas1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_FocusableAreas1
//go:noescape

func ElementFocusableAreas1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_SpatialNavigationSearch
//go:noescape

func CallElementSpatialNavigationSearch(
	this js.Ref,
	ptr unsafe.Pointer,

	dir uint32,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_SpatialNavigationSearch
//go:noescape

func ElementSpatialNavigationSearchFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_SpatialNavigationSearch1
//go:noescape

func CallElementSpatialNavigationSearch1(
	this js.Ref,
	ptr unsafe.Pointer,

	dir uint32,
) js.Ref

//go:wasmimport plat/js/web func_Element_SpatialNavigationSearch1
//go:noescape

func ElementSpatialNavigationSearch1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_SetPointerCapture
//go:noescape

func CallElementSetPointerCapture(
	this js.Ref,
	ptr unsafe.Pointer,

	pointerId int32,
) js.Ref

//go:wasmimport plat/js/web func_Element_SetPointerCapture
//go:noescape

func ElementSetPointerCaptureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ReleasePointerCapture
//go:noescape

func CallElementReleasePointerCapture(
	this js.Ref,
	ptr unsafe.Pointer,

	pointerId int32,
) js.Ref

//go:wasmimport plat/js/web func_Element_ReleasePointerCapture
//go:noescape

func ElementReleasePointerCaptureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_HasPointerCapture
//go:noescape

func CallElementHasPointerCapture(
	this js.Ref,
	ptr unsafe.Pointer,

	pointerId int32,
) js.Ref

//go:wasmimport plat/js/web func_Element_HasPointerCapture
//go:noescape

func ElementHasPointerCaptureFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Pseudo
//go:noescape

func CallElementPseudo(
	this js.Ref,
	ptr unsafe.Pointer,

	typ js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_Pseudo
//go:noescape

func ElementPseudoFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_InsertAdjacentHTML
//go:noescape

func CallElementInsertAdjacentHTML(
	this js.Ref,
	ptr unsafe.Pointer,

	position js.Ref,
	text js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_InsertAdjacentHTML
//go:noescape

func ElementInsertAdjacentHTMLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_SetHTML
//go:noescape

func CallElementSetHTML(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_SetHTML
//go:noescape

func ElementSetHTMLFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_SetHTML1
//go:noescape

func CallElementSetHTML1(
	this js.Ref,
	ptr unsafe.Pointer,

	input js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_SetHTML1
//go:noescape

func ElementSetHTML1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetClientRects
//go:noescape

func CallElementGetClientRects(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_GetClientRects
//go:noescape

func ElementGetClientRectsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetBoundingClientRect
//go:noescape

func CallElementGetBoundingClientRect(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_GetBoundingClientRect
//go:noescape

func ElementGetBoundingClientRectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_CheckVisibility
//go:noescape

func CallElementCheckVisibility(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_CheckVisibility
//go:noescape

func ElementCheckVisibilityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_CheckVisibility1
//go:noescape

func CallElementCheckVisibility1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_CheckVisibility1
//go:noescape

func ElementCheckVisibility1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ScrollIntoView
//go:noescape

func CallElementScrollIntoView(
	this js.Ref,
	ptr unsafe.Pointer,

	arg js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_ScrollIntoView
//go:noescape

func ElementScrollIntoViewFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ScrollIntoView1
//go:noescape

func CallElementScrollIntoView1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_ScrollIntoView1
//go:noescape

func ElementScrollIntoView1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Scroll
//go:noescape

func CallElementScroll(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_Scroll
//go:noescape

func ElementScrollFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Scroll1
//go:noescape

func CallElementScroll1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_Scroll1
//go:noescape

func ElementScroll1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Scroll2
//go:noescape

func CallElementScroll2(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_Element_Scroll2
//go:noescape

func ElementScroll2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ScrollTo
//go:noescape

func CallElementScrollTo(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_ScrollTo
//go:noescape

func ElementScrollToFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ScrollTo1
//go:noescape

func CallElementScrollTo1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_ScrollTo1
//go:noescape

func ElementScrollTo1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ScrollTo2
//go:noescape

func CallElementScrollTo2(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_Element_ScrollTo2
//go:noescape

func ElementScrollTo2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ScrollBy
//go:noescape

func CallElementScrollBy(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_ScrollBy
//go:noescape

func ElementScrollByFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ScrollBy1
//go:noescape

func CallElementScrollBy1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_ScrollBy1
//go:noescape

func ElementScrollBy1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ScrollBy2
//go:noescape

func CallElementScrollBy2(
	this js.Ref,
	ptr unsafe.Pointer,

	x float64,
	y float64,
) js.Ref

//go:wasmimport plat/js/web func_Element_ScrollBy2
//go:noescape

func ElementScrollBy2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetBoxQuads
//go:noescape

func CallElementGetBoxQuads(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_GetBoxQuads
//go:noescape

func ElementGetBoxQuadsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetBoxQuads1
//go:noescape

func CallElementGetBoxQuads1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_GetBoxQuads1
//go:noescape

func ElementGetBoxQuads1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ConvertQuadFromNode
//go:noescape

func CallElementConvertQuadFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	quad unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_ConvertQuadFromNode
//go:noescape

func ElementConvertQuadFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ConvertQuadFromNode1
//go:noescape

func CallElementConvertQuadFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	quad unsafe.Pointer,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_ConvertQuadFromNode1
//go:noescape

func ElementConvertQuadFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ConvertRectFromNode
//go:noescape

func CallElementConvertRectFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	rect js.Ref,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_ConvertRectFromNode
//go:noescape

func ElementConvertRectFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ConvertRectFromNode1
//go:noescape

func CallElementConvertRectFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	rect js.Ref,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_ConvertRectFromNode1
//go:noescape

func ElementConvertRectFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ConvertPointFromNode
//go:noescape

func CallElementConvertPointFromNode(
	this js.Ref,
	ptr unsafe.Pointer,

	point unsafe.Pointer,
	from js.Ref,
	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_ConvertPointFromNode
//go:noescape

func ElementConvertPointFromNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ConvertPointFromNode1
//go:noescape

func CallElementConvertPointFromNode1(
	this js.Ref,
	ptr unsafe.Pointer,

	point unsafe.Pointer,
	from js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_ConvertPointFromNode1
//go:noescape

func ElementConvertPointFromNode1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Prepend
//go:noescape

func CallElementPrepend(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Element_Prepend
//go:noescape

func ElementPrependFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Append
//go:noescape

func CallElementAppend(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Element_Append
//go:noescape

func ElementAppendFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ReplaceChildren
//go:noescape

func CallElementReplaceChildren(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Element_ReplaceChildren
//go:noescape

func ElementReplaceChildrenFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_QuerySelector
//go:noescape

func CallElementQuerySelector(
	this js.Ref,
	ptr unsafe.Pointer,

	selectors js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_QuerySelector
//go:noescape

func ElementQuerySelectorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_QuerySelectorAll
//go:noescape

func CallElementQuerySelectorAll(
	this js.Ref,
	ptr unsafe.Pointer,

	selectors js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_QuerySelectorAll
//go:noescape

func ElementQuerySelectorAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Before
//go:noescape

func CallElementBefore(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Element_Before
//go:noescape

func ElementBeforeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_After
//go:noescape

func CallElementAfter(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Element_After
//go:noescape

func ElementAfterFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_ReplaceWith
//go:noescape

func CallElementReplaceWith(
	this js.Ref,
	ptr unsafe.Pointer,

	nodesPtr unsafe.Pointer,
	nodesCount uint32,
) js.Ref

//go:wasmimport plat/js/web func_Element_ReplaceWith
//go:noescape

func ElementReplaceWithFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Remove
//go:noescape

func CallElementRemove(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_Remove
//go:noescape

func ElementRemoveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Animate
//go:noescape

func CallElementAnimate(
	this js.Ref,
	ptr unsafe.Pointer,

	keyframes js.Ref,
	options js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_Animate
//go:noescape

func ElementAnimateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_Animate1
//go:noescape

func CallElementAnimate1(
	this js.Ref,
	ptr unsafe.Pointer,

	keyframes js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_Element_Animate1
//go:noescape

func ElementAnimate1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetAnimations
//go:noescape

func CallElementGetAnimations(
	this js.Ref,
	ptr unsafe.Pointer,

	options unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web func_Element_GetAnimations
//go:noescape

func ElementGetAnimationsFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetAnimations1
//go:noescape

func CallElementGetAnimations1(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_GetAnimations1
//go:noescape

func ElementGetAnimations1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_Element_GetRegionFlowRanges
//go:noescape

func CallElementGetRegionFlowRanges(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_Element_GetRegionFlowRanges
//go:noescape

func ElementGetRegionFlowRangesFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_HTMLCollection_Length
//go:noescape

func GetHTMLCollectionLength(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web call_HTMLCollection_Item
//go:noescape

func CallHTMLCollectionItem(
	this js.Ref,
	ptr unsafe.Pointer,

	index uint32,
) js.Ref

//go:wasmimport plat/js/web func_HTMLCollection_Item
//go:noescape

func HTMLCollectionItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLCollection_NamedItem
//go:noescape

func CallHTMLCollectionNamedItem(
	this js.Ref,
	ptr unsafe.Pointer,

	name js.Ref,
) js.Ref

//go:wasmimport plat/js/web func_HTMLCollection_NamedItem
//go:noescape

func HTMLCollectionNamedItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web store_ElementCreationOptions
//go:noescape

func ElementCreationOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_ElementCreationOptions
//go:noescape

func ElementCreationOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_CDATASection_CDATASection
//go:noescape

func NewCDATASectionByCDATASection(
	data js.Ref) js.Ref

//go:wasmimport plat/js/web new_CDATASection_CDATASection1
//go:noescape

func NewCDATASectionByCDATASection1() js.Ref

//go:wasmimport plat/js/web new_Comment_Comment
//go:noescape

func NewCommentByComment(
	data js.Ref) js.Ref

//go:wasmimport plat/js/web new_Comment_Comment1
//go:noescape

func NewCommentByComment1() js.Ref

//go:wasmimport plat/js/web get_ProcessingInstruction_Target
//go:noescape

func GetProcessingInstructionTarget(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ProcessingInstruction_Sheet
//go:noescape

func GetProcessingInstructionSheet(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NodeIterator_Root
//go:noescape

func GetNodeIteratorRoot(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NodeIterator_ReferenceNode
//go:noescape

func GetNodeIteratorReferenceNode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NodeIterator_PointerBeforeReferenceNode
//go:noescape

func GetNodeIteratorPointerBeforeReferenceNode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_NodeIterator_WhatToShow
//go:noescape

func GetNodeIteratorWhatToShow(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_NodeIterator_Filter
//go:noescape

func GetNodeIteratorFilter(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_NodeIterator_NextNode
//go:noescape

func CallNodeIteratorNextNode(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NodeIterator_NextNode
//go:noescape

func NodeIteratorNextNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NodeIterator_PreviousNode
//go:noescape

func CallNodeIteratorPreviousNode(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NodeIterator_PreviousNode
//go:noescape

func NodeIteratorPreviousNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_NodeIterator_Detach
//go:noescape

func CallNodeIteratorDetach(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_NodeIterator_Detach
//go:noescape

func NodeIteratorDetachFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_TreeWalker_Root
//go:noescape

func GetTreeWalkerRoot(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TreeWalker_WhatToShow
//go:noescape

func GetTreeWalkerWhatToShow(
	this js.Ref,
	ptr unsafe.Pointer,
) uint32

//go:wasmimport plat/js/web get_TreeWalker_Filter
//go:noescape

func GetTreeWalkerFilter(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_TreeWalker_CurrentNode
//go:noescape

func GetTreeWalkerCurrentNode(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web set_TreeWalker_CurrentNode
//go:noescape

func SetTreeWalkerCurrentNode(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web call_TreeWalker_ParentNode
//go:noescape

func CallTreeWalkerParentNode(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TreeWalker_ParentNode
//go:noescape

func TreeWalkerParentNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TreeWalker_FirstChild
//go:noescape

func CallTreeWalkerFirstChild(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TreeWalker_FirstChild
//go:noescape

func TreeWalkerFirstChildFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TreeWalker_LastChild
//go:noescape

func CallTreeWalkerLastChild(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TreeWalker_LastChild
//go:noescape

func TreeWalkerLastChildFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TreeWalker_PreviousSibling
//go:noescape

func CallTreeWalkerPreviousSibling(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TreeWalker_PreviousSibling
//go:noescape

func TreeWalkerPreviousSiblingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TreeWalker_NextSibling
//go:noescape

func CallTreeWalkerNextSibling(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TreeWalker_NextSibling
//go:noescape

func TreeWalkerNextSiblingFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TreeWalker_PreviousNode
//go:noescape

func CallTreeWalkerPreviousNode(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TreeWalker_PreviousNode
//go:noescape

func TreeWalkerPreviousNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_TreeWalker_NextNode
//go:noescape

func CallTreeWalkerNextNode(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_TreeWalker_NextNode
//go:noescape

func TreeWalkerNextNodeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web get_ViewTransition_UpdateCallbackDone
//go:noescape

func GetViewTransitionUpdateCallbackDone(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ViewTransition_Ready
//go:noescape

func GetViewTransitionReady(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web get_ViewTransition_Finished
//go:noescape

func GetViewTransitionFinished(
	this js.Ref,
	ptr unsafe.Pointer,
) js.Ref

//go:wasmimport plat/js/web call_ViewTransition_SkipTransition
//go:noescape

func CallViewTransitionSkipTransition(
	this js.Ref,
	ptr unsafe.Pointer,

) js.Ref

//go:wasmimport plat/js/web func_ViewTransition_SkipTransition
//go:noescape

func ViewTransitionSkipTransitionFunc(this js.Ref) js.Ref

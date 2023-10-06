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

//go:wasmimport plat/js/web new_HTMLPortalElement_HTMLPortalElement
//go:noescape
func NewHTMLPortalElementByHTMLPortalElement() js.Ref

//go:wasmimport plat/js/web get_HTMLPortalElement_Src
//go:noescape
func GetHTMLPortalElementSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLPortalElement_Src
//go:noescape
func SetHTMLPortalElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLPortalElement_ReferrerPolicy
//go:noescape
func GetHTMLPortalElementReferrerPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLPortalElement_ReferrerPolicy
//go:noescape
func SetHTMLPortalElementReferrerPolicy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLPortalElement_Activate
//go:noescape
func HasHTMLPortalElementActivate(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLPortalElement_Activate
//go:noescape
func HTMLPortalElementActivateFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLPortalElement_Activate
//go:noescape
func CallHTMLPortalElementActivate(
	this js.Ref, retPtr unsafe.Pointer,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLPortalElement_Activate
//go:noescape
func TryHTMLPortalElementActivate(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLPortalElement_Activate1
//go:noescape
func HasHTMLPortalElementActivate1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLPortalElement_Activate1
//go:noescape
func HTMLPortalElementActivate1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLPortalElement_Activate1
//go:noescape
func CallHTMLPortalElementActivate1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLPortalElement_Activate1
//go:noescape
func TryHTMLPortalElementActivate1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLPortalElement_PostMessage
//go:noescape
func HasHTMLPortalElementPostMessage(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLPortalElement_PostMessage
//go:noescape
func HTMLPortalElementPostMessageFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLPortalElement_PostMessage
//go:noescape
func CallHTMLPortalElementPostMessage(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLPortalElement_PostMessage
//go:noescape
func TryHTMLPortalElementPostMessage(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLPortalElement_PostMessage1
//go:noescape
func HasHTMLPortalElementPostMessage1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLPortalElement_PostMessage1
//go:noescape
func HTMLPortalElementPostMessage1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLPortalElement_PostMessage1
//go:noescape
func CallHTMLPortalElementPostMessage1(
	this js.Ref, retPtr unsafe.Pointer,
	message js.Ref)

//go:wasmimport plat/js/web try_HTMLPortalElement_PostMessage1
//go:noescape
func TryHTMLPortalElementPostMessage1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	message js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLPreElement_HTMLPreElement
//go:noescape
func NewHTMLPreElementByHTMLPreElement() js.Ref

//go:wasmimport plat/js/web get_HTMLPreElement_Width
//go:noescape
func GetHTMLPreElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLPreElement_Width
//go:noescape
func SetHTMLPreElementWidth(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web new_HTMLProgressElement_HTMLProgressElement
//go:noescape
func NewHTMLProgressElementByHTMLProgressElement() js.Ref

//go:wasmimport plat/js/web get_HTMLProgressElement_Value
//go:noescape
func GetHTMLProgressElementValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLProgressElement_Value
//go:noescape
func SetHTMLProgressElementValue(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLProgressElement_Max
//go:noescape
func GetHTMLProgressElementMax(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLProgressElement_Max
//go:noescape
func SetHTMLProgressElementMax(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLProgressElement_Position
//go:noescape
func GetHTMLProgressElementPosition(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLProgressElement_Labels
//go:noescape
func GetHTMLProgressElementLabels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLQuoteElement_HTMLQuoteElement
//go:noescape
func NewHTMLQuoteElementByHTMLQuoteElement() js.Ref

//go:wasmimport plat/js/web get_HTMLQuoteElement_Cite
//go:noescape
func GetHTMLQuoteElementCite(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLQuoteElement_Cite
//go:noescape
func SetHTMLQuoteElementCite(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLSelectElement_HTMLSelectElement
//go:noescape
func NewHTMLSelectElementByHTMLSelectElement() js.Ref

//go:wasmimport plat/js/web get_HTMLSelectElement_Autocomplete
//go:noescape
func GetHTMLSelectElementAutocomplete(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSelectElement_Autocomplete
//go:noescape
func SetHTMLSelectElementAutocomplete(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSelectElement_Disabled
//go:noescape
func GetHTMLSelectElementDisabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSelectElement_Disabled
//go:noescape
func SetHTMLSelectElementDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSelectElement_Form
//go:noescape
func GetHTMLSelectElementForm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLSelectElement_Multiple
//go:noescape
func GetHTMLSelectElementMultiple(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSelectElement_Multiple
//go:noescape
func SetHTMLSelectElementMultiple(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSelectElement_Name
//go:noescape
func GetHTMLSelectElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSelectElement_Name
//go:noescape
func SetHTMLSelectElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSelectElement_Required
//go:noescape
func GetHTMLSelectElementRequired(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSelectElement_Required
//go:noescape
func SetHTMLSelectElementRequired(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSelectElement_Size
//go:noescape
func GetHTMLSelectElementSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSelectElement_Size
//go:noescape
func SetHTMLSelectElementSize(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSelectElement_Type
//go:noescape
func GetHTMLSelectElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLSelectElement_Options
//go:noescape
func GetHTMLSelectElementOptions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLSelectElement_Length
//go:noescape
func GetHTMLSelectElementLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSelectElement_Length
//go:noescape
func SetHTMLSelectElementLength(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSelectElement_SelectedOptions
//go:noescape
func GetHTMLSelectElementSelectedOptions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLSelectElement_SelectedIndex
//go:noescape
func GetHTMLSelectElementSelectedIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSelectElement_SelectedIndex
//go:noescape
func SetHTMLSelectElementSelectedIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSelectElement_Value
//go:noescape
func GetHTMLSelectElementValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSelectElement_Value
//go:noescape
func SetHTMLSelectElementValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSelectElement_WillValidate
//go:noescape
func GetHTMLSelectElementWillValidate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLSelectElement_Validity
//go:noescape
func GetHTMLSelectElementValidity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLSelectElement_ValidationMessage
//go:noescape
func GetHTMLSelectElementValidationMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLSelectElement_Labels
//go:noescape
func GetHTMLSelectElementLabels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLSelectElement_Item
//go:noescape
func HasHTMLSelectElementItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLSelectElement_Item
//go:noescape
func HTMLSelectElementItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSelectElement_Item
//go:noescape
func CallHTMLSelectElementItem(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32)

//go:wasmimport plat/js/web try_HTMLSelectElement_Item
//go:noescape
func TryHTMLSelectElementItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLSelectElement_NamedItem
//go:noescape
func HasHTMLSelectElementNamedItem(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLSelectElement_NamedItem
//go:noescape
func HTMLSelectElementNamedItemFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSelectElement_NamedItem
//go:noescape
func CallHTMLSelectElementNamedItem(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_HTMLSelectElement_NamedItem
//go:noescape
func TryHTMLSelectElementNamedItem(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLSelectElement_Add
//go:noescape
func HasHTMLSelectElementAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLSelectElement_Add
//go:noescape
func HTMLSelectElementAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSelectElement_Add
//go:noescape
func CallHTMLSelectElementAdd(
	this js.Ref, retPtr unsafe.Pointer,
	element js.Ref,
	before js.Ref)

//go:wasmimport plat/js/web try_HTMLSelectElement_Add
//go:noescape
func TryHTMLSelectElementAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref,
	before js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLSelectElement_Add1
//go:noescape
func HasHTMLSelectElementAdd1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLSelectElement_Add1
//go:noescape
func HTMLSelectElementAdd1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSelectElement_Add1
//go:noescape
func CallHTMLSelectElementAdd1(
	this js.Ref, retPtr unsafe.Pointer,
	element js.Ref)

//go:wasmimport plat/js/web try_HTMLSelectElement_Add1
//go:noescape
func TryHTMLSelectElementAdd1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLSelectElement_Remove
//go:noescape
func HasHTMLSelectElementRemove(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLSelectElement_Remove
//go:noescape
func HTMLSelectElementRemoveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSelectElement_Remove
//go:noescape
func CallHTMLSelectElementRemove(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLSelectElement_Remove
//go:noescape
func TryHTMLSelectElementRemove(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLSelectElement_Remove1
//go:noescape
func HasHTMLSelectElementRemove1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLSelectElement_Remove1
//go:noescape
func HTMLSelectElementRemove1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSelectElement_Remove1
//go:noescape
func CallHTMLSelectElementRemove1(
	this js.Ref, retPtr unsafe.Pointer,
	index int32)

//go:wasmimport plat/js/web try_HTMLSelectElement_Remove1
//go:noescape
func TryHTMLSelectElementRemove1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index int32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLSelectElement_Set
//go:noescape
func HasHTMLSelectElementSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLSelectElement_Set
//go:noescape
func HTMLSelectElementSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSelectElement_Set
//go:noescape
func CallHTMLSelectElementSet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	option js.Ref)

//go:wasmimport plat/js/web try_HTMLSelectElement_Set
//go:noescape
func TryHTMLSelectElementSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	option js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLSelectElement_CheckValidity
//go:noescape
func HasHTMLSelectElementCheckValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLSelectElement_CheckValidity
//go:noescape
func HTMLSelectElementCheckValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSelectElement_CheckValidity
//go:noescape
func CallHTMLSelectElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLSelectElement_CheckValidity
//go:noescape
func TryHTMLSelectElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLSelectElement_ReportValidity
//go:noescape
func HasHTMLSelectElementReportValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLSelectElement_ReportValidity
//go:noescape
func HTMLSelectElementReportValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSelectElement_ReportValidity
//go:noescape
func CallHTMLSelectElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLSelectElement_ReportValidity
//go:noescape
func TryHTMLSelectElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLSelectElement_SetCustomValidity
//go:noescape
func HasHTMLSelectElementSetCustomValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLSelectElement_SetCustomValidity
//go:noescape
func HTMLSelectElementSetCustomValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLSelectElement_SetCustomValidity
//go:noescape
func CallHTMLSelectElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer,
	err js.Ref)

//go:wasmimport plat/js/web try_HTMLSelectElement_SetCustomValidity
//go:noescape
func TryHTMLSelectElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	err js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLSourceElement_HTMLSourceElement
//go:noescape
func NewHTMLSourceElementByHTMLSourceElement() js.Ref

//go:wasmimport plat/js/web get_HTMLSourceElement_Src
//go:noescape
func GetHTMLSourceElementSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSourceElement_Src
//go:noescape
func SetHTMLSourceElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSourceElement_Type
//go:noescape
func GetHTMLSourceElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSourceElement_Type
//go:noescape
func SetHTMLSourceElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSourceElement_Srcset
//go:noescape
func GetHTMLSourceElementSrcset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSourceElement_Srcset
//go:noescape
func SetHTMLSourceElementSrcset(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSourceElement_Sizes
//go:noescape
func GetHTMLSourceElementSizes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSourceElement_Sizes
//go:noescape
func SetHTMLSourceElementSizes(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSourceElement_Media
//go:noescape
func GetHTMLSourceElementMedia(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSourceElement_Media
//go:noescape
func SetHTMLSourceElementMedia(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSourceElement_Width
//go:noescape
func GetHTMLSourceElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSourceElement_Width
//go:noescape
func SetHTMLSourceElementWidth(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLSourceElement_Height
//go:noescape
func GetHTMLSourceElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLSourceElement_Height
//go:noescape
func SetHTMLSourceElementHeight(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web new_HTMLSpanElement_HTMLSpanElement
//go:noescape
func NewHTMLSpanElementByHTMLSpanElement() js.Ref

//go:wasmimport plat/js/web new_HTMLStyleElement_HTMLStyleElement
//go:noescape
func NewHTMLStyleElementByHTMLStyleElement() js.Ref

//go:wasmimport plat/js/web get_HTMLStyleElement_Disabled
//go:noescape
func GetHTMLStyleElementDisabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLStyleElement_Disabled
//go:noescape
func SetHTMLStyleElementDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLStyleElement_Media
//go:noescape
func GetHTMLStyleElementMedia(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLStyleElement_Media
//go:noescape
func SetHTMLStyleElementMedia(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLStyleElement_Blocking
//go:noescape
func GetHTMLStyleElementBlocking(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLStyleElement_Type
//go:noescape
func GetHTMLStyleElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLStyleElement_Type
//go:noescape
func SetHTMLStyleElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLStyleElement_Sheet
//go:noescape
func GetHTMLStyleElementSheet(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLTableCaptionElement_HTMLTableCaptionElement
//go:noescape
func NewHTMLTableCaptionElementByHTMLTableCaptionElement() js.Ref

//go:wasmimport plat/js/web get_HTMLTableCaptionElement_Align
//go:noescape
func GetHTMLTableCaptionElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCaptionElement_Align
//go:noescape
func SetHTMLTableCaptionElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLTableCellElement_HTMLTableCellElement
//go:noescape
func NewHTMLTableCellElementByHTMLTableCellElement() js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_ColSpan
//go:noescape
func GetHTMLTableCellElementColSpan(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_ColSpan
//go:noescape
func SetHTMLTableCellElementColSpan(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_RowSpan
//go:noescape
func GetHTMLTableCellElementRowSpan(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_RowSpan
//go:noescape
func SetHTMLTableCellElementRowSpan(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_Headers
//go:noescape
func GetHTMLTableCellElementHeaders(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_Headers
//go:noescape
func SetHTMLTableCellElementHeaders(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_CellIndex
//go:noescape
func GetHTMLTableCellElementCellIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTableCellElement_Scope
//go:noescape
func GetHTMLTableCellElementScope(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_Scope
//go:noescape
func SetHTMLTableCellElementScope(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_Abbr
//go:noescape
func GetHTMLTableCellElementAbbr(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_Abbr
//go:noescape
func SetHTMLTableCellElementAbbr(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_Align
//go:noescape
func GetHTMLTableCellElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_Align
//go:noescape
func SetHTMLTableCellElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_Axis
//go:noescape
func GetHTMLTableCellElementAxis(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_Axis
//go:noescape
func SetHTMLTableCellElementAxis(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_Height
//go:noescape
func GetHTMLTableCellElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_Height
//go:noescape
func SetHTMLTableCellElementHeight(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_Width
//go:noescape
func GetHTMLTableCellElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_Width
//go:noescape
func SetHTMLTableCellElementWidth(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_Ch
//go:noescape
func GetHTMLTableCellElementCh(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_Ch
//go:noescape
func SetHTMLTableCellElementCh(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_ChOff
//go:noescape
func GetHTMLTableCellElementChOff(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_ChOff
//go:noescape
func SetHTMLTableCellElementChOff(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_NoWrap
//go:noescape
func GetHTMLTableCellElementNoWrap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_NoWrap
//go:noescape
func SetHTMLTableCellElementNoWrap(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_VAlign
//go:noescape
func GetHTMLTableCellElementVAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_VAlign
//go:noescape
func SetHTMLTableCellElementVAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableCellElement_BgColor
//go:noescape
func GetHTMLTableCellElementBgColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableCellElement_BgColor
//go:noescape
func SetHTMLTableCellElementBgColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLTableColElement_HTMLTableColElement
//go:noescape
func NewHTMLTableColElementByHTMLTableColElement() js.Ref

//go:wasmimport plat/js/web get_HTMLTableColElement_Span
//go:noescape
func GetHTMLTableColElementSpan(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableColElement_Span
//go:noescape
func SetHTMLTableColElementSpan(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableColElement_Align
//go:noescape
func GetHTMLTableColElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableColElement_Align
//go:noescape
func SetHTMLTableColElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableColElement_Ch
//go:noescape
func GetHTMLTableColElementCh(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableColElement_Ch
//go:noescape
func SetHTMLTableColElementCh(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableColElement_ChOff
//go:noescape
func GetHTMLTableColElementChOff(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableColElement_ChOff
//go:noescape
func SetHTMLTableColElementChOff(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableColElement_VAlign
//go:noescape
func GetHTMLTableColElementVAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableColElement_VAlign
//go:noescape
func SetHTMLTableColElementVAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableColElement_Width
//go:noescape
func GetHTMLTableColElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableColElement_Width
//go:noescape
func SetHTMLTableColElementWidth(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLTableRowElement_HTMLTableRowElement
//go:noescape
func NewHTMLTableRowElementByHTMLTableRowElement() js.Ref

//go:wasmimport plat/js/web get_HTMLTableRowElement_RowIndex
//go:noescape
func GetHTMLTableRowElementRowIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTableRowElement_SectionRowIndex
//go:noescape
func GetHTMLTableRowElementSectionRowIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTableRowElement_Cells
//go:noescape
func GetHTMLTableRowElementCells(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTableRowElement_Align
//go:noescape
func GetHTMLTableRowElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableRowElement_Align
//go:noescape
func SetHTMLTableRowElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableRowElement_Ch
//go:noescape
func GetHTMLTableRowElementCh(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableRowElement_Ch
//go:noescape
func SetHTMLTableRowElementCh(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableRowElement_ChOff
//go:noescape
func GetHTMLTableRowElementChOff(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableRowElement_ChOff
//go:noescape
func SetHTMLTableRowElementChOff(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableRowElement_VAlign
//go:noescape
func GetHTMLTableRowElementVAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableRowElement_VAlign
//go:noescape
func SetHTMLTableRowElementVAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableRowElement_BgColor
//go:noescape
func GetHTMLTableRowElementBgColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableRowElement_BgColor
//go:noescape
func SetHTMLTableRowElementBgColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLTableRowElement_InsertCell
//go:noescape
func HasHTMLTableRowElementInsertCell(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableRowElement_InsertCell
//go:noescape
func HTMLTableRowElementInsertCellFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableRowElement_InsertCell
//go:noescape
func CallHTMLTableRowElementInsertCell(
	this js.Ref, retPtr unsafe.Pointer,
	index int32)

//go:wasmimport plat/js/web try_HTMLTableRowElement_InsertCell
//go:noescape
func TryHTMLTableRowElementInsertCell(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index int32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableRowElement_InsertCell1
//go:noescape
func HasHTMLTableRowElementInsertCell1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableRowElement_InsertCell1
//go:noescape
func HTMLTableRowElementInsertCell1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableRowElement_InsertCell1
//go:noescape
func CallHTMLTableRowElementInsertCell1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTableRowElement_InsertCell1
//go:noescape
func TryHTMLTableRowElementInsertCell1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableRowElement_DeleteCell
//go:noescape
func HasHTMLTableRowElementDeleteCell(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableRowElement_DeleteCell
//go:noescape
func HTMLTableRowElementDeleteCellFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableRowElement_DeleteCell
//go:noescape
func CallHTMLTableRowElementDeleteCell(
	this js.Ref, retPtr unsafe.Pointer,
	index int32)

//go:wasmimport plat/js/web try_HTMLTableRowElement_DeleteCell
//go:noescape
func TryHTMLTableRowElementDeleteCell(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index int32) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLTableSectionElement_HTMLTableSectionElement
//go:noescape
func NewHTMLTableSectionElementByHTMLTableSectionElement() js.Ref

//go:wasmimport plat/js/web get_HTMLTableSectionElement_Rows
//go:noescape
func GetHTMLTableSectionElementRows(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTableSectionElement_Align
//go:noescape
func GetHTMLTableSectionElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableSectionElement_Align
//go:noescape
func SetHTMLTableSectionElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableSectionElement_Ch
//go:noescape
func GetHTMLTableSectionElementCh(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableSectionElement_Ch
//go:noescape
func SetHTMLTableSectionElementCh(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableSectionElement_ChOff
//go:noescape
func GetHTMLTableSectionElementChOff(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableSectionElement_ChOff
//go:noescape
func SetHTMLTableSectionElementChOff(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableSectionElement_VAlign
//go:noescape
func GetHTMLTableSectionElementVAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableSectionElement_VAlign
//go:noescape
func SetHTMLTableSectionElementVAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLTableSectionElement_InsertRow
//go:noescape
func HasHTMLTableSectionElementInsertRow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableSectionElement_InsertRow
//go:noescape
func HTMLTableSectionElementInsertRowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableSectionElement_InsertRow
//go:noescape
func CallHTMLTableSectionElementInsertRow(
	this js.Ref, retPtr unsafe.Pointer,
	index int32)

//go:wasmimport plat/js/web try_HTMLTableSectionElement_InsertRow
//go:noescape
func TryHTMLTableSectionElementInsertRow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index int32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableSectionElement_InsertRow1
//go:noescape
func HasHTMLTableSectionElementInsertRow1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableSectionElement_InsertRow1
//go:noescape
func HTMLTableSectionElementInsertRow1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableSectionElement_InsertRow1
//go:noescape
func CallHTMLTableSectionElementInsertRow1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTableSectionElement_InsertRow1
//go:noescape
func TryHTMLTableSectionElementInsertRow1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableSectionElement_DeleteRow
//go:noescape
func HasHTMLTableSectionElementDeleteRow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableSectionElement_DeleteRow
//go:noescape
func HTMLTableSectionElementDeleteRowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableSectionElement_DeleteRow
//go:noescape
func CallHTMLTableSectionElementDeleteRow(
	this js.Ref, retPtr unsafe.Pointer,
	index int32)

//go:wasmimport plat/js/web try_HTMLTableSectionElement_DeleteRow
//go:noescape
func TryHTMLTableSectionElementDeleteRow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index int32) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLTableElement_HTMLTableElement
//go:noescape
func NewHTMLTableElementByHTMLTableElement() js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_Caption
//go:noescape
func GetHTMLTableElementCaption(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_Caption
//go:noescape
func SetHTMLTableElementCaption(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_THead
//go:noescape
func GetHTMLTableElementTHead(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_THead
//go:noescape
func SetHTMLTableElementTHead(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_TFoot
//go:noescape
func GetHTMLTableElementTFoot(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_TFoot
//go:noescape
func SetHTMLTableElementTFoot(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_TBodies
//go:noescape
func GetHTMLTableElementTBodies(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTableElement_Rows
//go:noescape
func GetHTMLTableElementRows(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTableElement_Align
//go:noescape
func GetHTMLTableElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_Align
//go:noescape
func SetHTMLTableElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_Border
//go:noescape
func GetHTMLTableElementBorder(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_Border
//go:noescape
func SetHTMLTableElementBorder(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_Frame
//go:noescape
func GetHTMLTableElementFrame(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_Frame
//go:noescape
func SetHTMLTableElementFrame(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_Rules
//go:noescape
func GetHTMLTableElementRules(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_Rules
//go:noescape
func SetHTMLTableElementRules(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_Summary
//go:noescape
func GetHTMLTableElementSummary(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_Summary
//go:noescape
func SetHTMLTableElementSummary(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_Width
//go:noescape
func GetHTMLTableElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_Width
//go:noescape
func SetHTMLTableElementWidth(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_BgColor
//go:noescape
func GetHTMLTableElementBgColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_BgColor
//go:noescape
func SetHTMLTableElementBgColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_CellPadding
//go:noescape
func GetHTMLTableElementCellPadding(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_CellPadding
//go:noescape
func SetHTMLTableElementCellPadding(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTableElement_CellSpacing
//go:noescape
func GetHTMLTableElementCellSpacing(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTableElement_CellSpacing
//go:noescape
func SetHTMLTableElementCellSpacing(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLTableElement_CreateCaption
//go:noescape
func HasHTMLTableElementCreateCaption(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableElement_CreateCaption
//go:noescape
func HTMLTableElementCreateCaptionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableElement_CreateCaption
//go:noescape
func CallHTMLTableElementCreateCaption(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTableElement_CreateCaption
//go:noescape
func TryHTMLTableElementCreateCaption(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableElement_DeleteCaption
//go:noescape
func HasHTMLTableElementDeleteCaption(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableElement_DeleteCaption
//go:noescape
func HTMLTableElementDeleteCaptionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableElement_DeleteCaption
//go:noescape
func CallHTMLTableElementDeleteCaption(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTableElement_DeleteCaption
//go:noescape
func TryHTMLTableElementDeleteCaption(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableElement_CreateTHead
//go:noescape
func HasHTMLTableElementCreateTHead(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableElement_CreateTHead
//go:noescape
func HTMLTableElementCreateTHeadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableElement_CreateTHead
//go:noescape
func CallHTMLTableElementCreateTHead(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTableElement_CreateTHead
//go:noescape
func TryHTMLTableElementCreateTHead(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableElement_DeleteTHead
//go:noescape
func HasHTMLTableElementDeleteTHead(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableElement_DeleteTHead
//go:noescape
func HTMLTableElementDeleteTHeadFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableElement_DeleteTHead
//go:noescape
func CallHTMLTableElementDeleteTHead(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTableElement_DeleteTHead
//go:noescape
func TryHTMLTableElementDeleteTHead(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableElement_CreateTFoot
//go:noescape
func HasHTMLTableElementCreateTFoot(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableElement_CreateTFoot
//go:noescape
func HTMLTableElementCreateTFootFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableElement_CreateTFoot
//go:noescape
func CallHTMLTableElementCreateTFoot(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTableElement_CreateTFoot
//go:noescape
func TryHTMLTableElementCreateTFoot(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableElement_DeleteTFoot
//go:noescape
func HasHTMLTableElementDeleteTFoot(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableElement_DeleteTFoot
//go:noescape
func HTMLTableElementDeleteTFootFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableElement_DeleteTFoot
//go:noescape
func CallHTMLTableElementDeleteTFoot(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTableElement_DeleteTFoot
//go:noescape
func TryHTMLTableElementDeleteTFoot(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableElement_CreateTBody
//go:noescape
func HasHTMLTableElementCreateTBody(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableElement_CreateTBody
//go:noescape
func HTMLTableElementCreateTBodyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableElement_CreateTBody
//go:noescape
func CallHTMLTableElementCreateTBody(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTableElement_CreateTBody
//go:noescape
func TryHTMLTableElementCreateTBody(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableElement_InsertRow
//go:noescape
func HasHTMLTableElementInsertRow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableElement_InsertRow
//go:noescape
func HTMLTableElementInsertRowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableElement_InsertRow
//go:noescape
func CallHTMLTableElementInsertRow(
	this js.Ref, retPtr unsafe.Pointer,
	index int32)

//go:wasmimport plat/js/web try_HTMLTableElement_InsertRow
//go:noescape
func TryHTMLTableElementInsertRow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index int32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableElement_InsertRow1
//go:noescape
func HasHTMLTableElementInsertRow1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableElement_InsertRow1
//go:noescape
func HTMLTableElementInsertRow1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableElement_InsertRow1
//go:noescape
func CallHTMLTableElementInsertRow1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTableElement_InsertRow1
//go:noescape
func TryHTMLTableElementInsertRow1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTableElement_DeleteRow
//go:noescape
func HasHTMLTableElementDeleteRow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTableElement_DeleteRow
//go:noescape
func HTMLTableElementDeleteRowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTableElement_DeleteRow
//go:noescape
func CallHTMLTableElementDeleteRow(
	this js.Ref, retPtr unsafe.Pointer,
	index int32)

//go:wasmimport plat/js/web try_HTMLTableElement_DeleteRow
//go:noescape
func TryHTMLTableElementDeleteRow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index int32) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLTemplateElement_HTMLTemplateElement
//go:noescape
func NewHTMLTemplateElementByHTMLTemplateElement() js.Ref

//go:wasmimport plat/js/web get_HTMLTemplateElement_Content
//go:noescape
func GetHTMLTemplateElementContent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLTextAreaElement_HTMLTextAreaElement
//go:noescape
func NewHTMLTextAreaElementByHTMLTextAreaElement() js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Autocomplete
//go:noescape
func GetHTMLTextAreaElementAutocomplete(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_Autocomplete
//go:noescape
func SetHTMLTextAreaElementAutocomplete(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Cols
//go:noescape
func GetHTMLTextAreaElementCols(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_Cols
//go:noescape
func SetHTMLTextAreaElementCols(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_DirName
//go:noescape
func GetHTMLTextAreaElementDirName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_DirName
//go:noescape
func SetHTMLTextAreaElementDirName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Disabled
//go:noescape
func GetHTMLTextAreaElementDisabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_Disabled
//go:noescape
func SetHTMLTextAreaElementDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Form
//go:noescape
func GetHTMLTextAreaElementForm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTextAreaElement_MaxLength
//go:noescape
func GetHTMLTextAreaElementMaxLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_MaxLength
//go:noescape
func SetHTMLTextAreaElementMaxLength(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_MinLength
//go:noescape
func GetHTMLTextAreaElementMinLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_MinLength
//go:noescape
func SetHTMLTextAreaElementMinLength(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Name
//go:noescape
func GetHTMLTextAreaElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_Name
//go:noescape
func SetHTMLTextAreaElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Placeholder
//go:noescape
func GetHTMLTextAreaElementPlaceholder(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_Placeholder
//go:noescape
func SetHTMLTextAreaElementPlaceholder(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_ReadOnly
//go:noescape
func GetHTMLTextAreaElementReadOnly(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_ReadOnly
//go:noescape
func SetHTMLTextAreaElementReadOnly(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Required
//go:noescape
func GetHTMLTextAreaElementRequired(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_Required
//go:noescape
func SetHTMLTextAreaElementRequired(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Rows
//go:noescape
func GetHTMLTextAreaElementRows(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_Rows
//go:noescape
func SetHTMLTextAreaElementRows(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Wrap
//go:noescape
func GetHTMLTextAreaElementWrap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_Wrap
//go:noescape
func SetHTMLTextAreaElementWrap(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Type
//go:noescape
func GetHTMLTextAreaElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTextAreaElement_DefaultValue
//go:noescape
func GetHTMLTextAreaElementDefaultValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_DefaultValue
//go:noescape
func SetHTMLTextAreaElementDefaultValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Value
//go:noescape
func GetHTMLTextAreaElementValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_Value
//go:noescape
func SetHTMLTextAreaElementValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_TextLength
//go:noescape
func GetHTMLTextAreaElementTextLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTextAreaElement_WillValidate
//go:noescape
func GetHTMLTextAreaElementWillValidate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Validity
//go:noescape
func GetHTMLTextAreaElementValidity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTextAreaElement_ValidationMessage
//go:noescape
func GetHTMLTextAreaElementValidationMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTextAreaElement_Labels
//go:noescape
func GetHTMLTextAreaElementLabels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTextAreaElement_SelectionStart
//go:noescape
func GetHTMLTextAreaElementSelectionStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_SelectionStart
//go:noescape
func SetHTMLTextAreaElementSelectionStart(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_SelectionEnd
//go:noescape
func GetHTMLTextAreaElementSelectionEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_SelectionEnd
//go:noescape
func SetHTMLTextAreaElementSelectionEnd(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTextAreaElement_SelectionDirection
//go:noescape
func GetHTMLTextAreaElementSelectionDirection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTextAreaElement_SelectionDirection
//go:noescape
func SetHTMLTextAreaElementSelectionDirection(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLTextAreaElement_CheckValidity
//go:noescape
func HasHTMLTextAreaElementCheckValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTextAreaElement_CheckValidity
//go:noescape
func HTMLTextAreaElementCheckValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTextAreaElement_CheckValidity
//go:noescape
func CallHTMLTextAreaElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTextAreaElement_CheckValidity
//go:noescape
func TryHTMLTextAreaElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTextAreaElement_ReportValidity
//go:noescape
func HasHTMLTextAreaElementReportValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTextAreaElement_ReportValidity
//go:noescape
func HTMLTextAreaElementReportValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTextAreaElement_ReportValidity
//go:noescape
func CallHTMLTextAreaElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTextAreaElement_ReportValidity
//go:noescape
func TryHTMLTextAreaElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTextAreaElement_SetCustomValidity
//go:noescape
func HasHTMLTextAreaElementSetCustomValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTextAreaElement_SetCustomValidity
//go:noescape
func HTMLTextAreaElementSetCustomValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTextAreaElement_SetCustomValidity
//go:noescape
func CallHTMLTextAreaElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer,
	err js.Ref)

//go:wasmimport plat/js/web try_HTMLTextAreaElement_SetCustomValidity
//go:noescape
func TryHTMLTextAreaElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	err js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTextAreaElement_Select
//go:noescape
func HasHTMLTextAreaElementSelect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTextAreaElement_Select
//go:noescape
func HTMLTextAreaElementSelectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTextAreaElement_Select
//go:noescape
func CallHTMLTextAreaElementSelect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLTextAreaElement_Select
//go:noescape
func TryHTMLTextAreaElementSelect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTextAreaElement_SetRangeText
//go:noescape
func HasHTMLTextAreaElementSetRangeText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTextAreaElement_SetRangeText
//go:noescape
func HTMLTextAreaElementSetRangeTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTextAreaElement_SetRangeText
//go:noescape
func CallHTMLTextAreaElementSetRangeText(
	this js.Ref, retPtr unsafe.Pointer,
	replacement js.Ref)

//go:wasmimport plat/js/web try_HTMLTextAreaElement_SetRangeText
//go:noescape
func TryHTMLTextAreaElementSetRangeText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	replacement js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTextAreaElement_SetRangeText1
//go:noescape
func HasHTMLTextAreaElementSetRangeText1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTextAreaElement_SetRangeText1
//go:noescape
func HTMLTextAreaElementSetRangeText1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTextAreaElement_SetRangeText1
//go:noescape
func CallHTMLTextAreaElementSetRangeText1(
	this js.Ref, retPtr unsafe.Pointer,
	replacement js.Ref,
	start uint32,
	end uint32,
	selectionMode uint32)

//go:wasmimport plat/js/web try_HTMLTextAreaElement_SetRangeText1
//go:noescape
func TryHTMLTextAreaElementSetRangeText1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	replacement js.Ref,
	start uint32,
	end uint32,
	selectionMode uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTextAreaElement_SetRangeText2
//go:noescape
func HasHTMLTextAreaElementSetRangeText2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTextAreaElement_SetRangeText2
//go:noescape
func HTMLTextAreaElementSetRangeText2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTextAreaElement_SetRangeText2
//go:noescape
func CallHTMLTextAreaElementSetRangeText2(
	this js.Ref, retPtr unsafe.Pointer,
	replacement js.Ref,
	start uint32,
	end uint32)

//go:wasmimport plat/js/web try_HTMLTextAreaElement_SetRangeText2
//go:noescape
func TryHTMLTextAreaElementSetRangeText2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	replacement js.Ref,
	start uint32,
	end uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTextAreaElement_SetSelectionRange
//go:noescape
func HasHTMLTextAreaElementSetSelectionRange(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTextAreaElement_SetSelectionRange
//go:noescape
func HTMLTextAreaElementSetSelectionRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTextAreaElement_SetSelectionRange
//go:noescape
func CallHTMLTextAreaElementSetSelectionRange(
	this js.Ref, retPtr unsafe.Pointer,
	start uint32,
	end uint32,
	direction js.Ref)

//go:wasmimport plat/js/web try_HTMLTextAreaElement_SetSelectionRange
//go:noescape
func TryHTMLTextAreaElementSetSelectionRange(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	start uint32,
	end uint32,
	direction js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLTextAreaElement_SetSelectionRange1
//go:noescape
func HasHTMLTextAreaElementSetSelectionRange1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLTextAreaElement_SetSelectionRange1
//go:noescape
func HTMLTextAreaElementSetSelectionRange1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLTextAreaElement_SetSelectionRange1
//go:noescape
func CallHTMLTextAreaElementSetSelectionRange1(
	this js.Ref, retPtr unsafe.Pointer,
	start uint32,
	end uint32)

//go:wasmimport plat/js/web try_HTMLTextAreaElement_SetSelectionRange1
//go:noescape
func TryHTMLTextAreaElementSetSelectionRange1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	start uint32,
	end uint32) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLTimeElement_HTMLTimeElement
//go:noescape
func NewHTMLTimeElementByHTMLTimeElement() js.Ref

//go:wasmimport plat/js/web get_HTMLTimeElement_DateTime
//go:noescape
func GetHTMLTimeElementDateTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTimeElement_DateTime
//go:noescape
func SetHTMLTimeElementDateTime(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLTitleElement_HTMLTitleElement
//go:noescape
func NewHTMLTitleElementByHTMLTitleElement() js.Ref

//go:wasmimport plat/js/web get_HTMLTitleElement_Text
//go:noescape
func GetHTMLTitleElementText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTitleElement_Text
//go:noescape
func SetHTMLTitleElementText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLTrackElement_HTMLTrackElement
//go:noescape
func NewHTMLTrackElementByHTMLTrackElement() js.Ref

//go:wasmimport plat/js/web get_HTMLTrackElement_Kind
//go:noescape
func GetHTMLTrackElementKind(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTrackElement_Kind
//go:noescape
func SetHTMLTrackElementKind(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTrackElement_Src
//go:noescape
func GetHTMLTrackElementSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTrackElement_Src
//go:noescape
func SetHTMLTrackElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTrackElement_Srclang
//go:noescape
func GetHTMLTrackElementSrclang(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTrackElement_Srclang
//go:noescape
func SetHTMLTrackElementSrclang(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTrackElement_Label
//go:noescape
func GetHTMLTrackElementLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTrackElement_Label
//go:noescape
func SetHTMLTrackElementLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTrackElement_Default
//go:noescape
func GetHTMLTrackElementDefault(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLTrackElement_Default
//go:noescape
func SetHTMLTrackElementDefault(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLTrackElement_ReadyState
//go:noescape
func GetHTMLTrackElementReadyState(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLTrackElement_Track
//go:noescape
func GetHTMLTrackElementTrack(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLUListElement_HTMLUListElement
//go:noescape
func NewHTMLUListElementByHTMLUListElement() js.Ref

//go:wasmimport plat/js/web get_HTMLUListElement_Compact
//go:noescape
func GetHTMLUListElementCompact(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLUListElement_Compact
//go:noescape
func SetHTMLUListElementCompact(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLUListElement_Type
//go:noescape
func GetHTMLUListElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLUListElement_Type
//go:noescape
func SetHTMLUListElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web store_HashChangeEventInit
//go:noescape
func HashChangeEventInitJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HashChangeEventInit
//go:noescape
func HashChangeEventInitJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web new_HashChangeEvent_HashChangeEvent
//go:noescape
func NewHashChangeEventByHashChangeEvent(
	typ js.Ref,
	eventInitDict unsafe.Pointer) js.Ref

//go:wasmimport plat/js/web new_HashChangeEvent_HashChangeEvent1
//go:noescape
func NewHashChangeEventByHashChangeEvent1(
	typ js.Ref) js.Ref

//go:wasmimport plat/js/web get_HashChangeEvent_OldURL
//go:noescape
func GetHashChangeEventOldURL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HashChangeEvent_NewURL
//go:noescape
func GetHashChangeEventNewURL(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_HevcBitstreamFormat
//go:noescape
func ConstOfHevcBitstreamFormat(str js.Ref) uint32

//go:wasmimport plat/js/web store_HevcEncoderConfig
//go:noescape
func HevcEncoderConfigJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HevcEncoderConfig
//go:noescape
func HevcEncoderConfigJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_HighlightType
//go:noescape
func ConstOfHighlightType(str js.Ref) uint32

//go:wasmimport plat/js/web new_Highlight_Highlight
//go:noescape
func NewHighlightByHighlight(
	initialRangesPtr unsafe.Pointer,
	initialRangesCount uint32) js.Ref

//go:wasmimport plat/js/web get_Highlight_Priority
//go:noescape
func GetHighlightPriority(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Highlight_Priority
//go:noescape
func SetHighlightPriority(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_Highlight_Type
//go:noescape
func GetHighlightType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_Highlight_Type
//go:noescape
func SetHighlightType(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web store_HkdfParams
//go:noescape
func HkdfParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HkdfParams
//go:noescape
func HkdfParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_HmacImportParams
//go:noescape
func HmacImportParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HmacImportParams
//go:noescape
func HmacImportParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_KeyAlgorithm
//go:noescape
func KeyAlgorithmJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_KeyAlgorithm
//go:noescape
func KeyAlgorithmJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_HmacKeyAlgorithm
//go:noescape
func HmacKeyAlgorithmJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HmacKeyAlgorithm
//go:noescape
func HmacKeyAlgorithmJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_HmacKeyGenParams
//go:noescape
func HmacKeyGenParamsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_HmacKeyGenParams
//go:noescape
func HmacKeyGenParamsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_IDBCursorDirection
//go:noescape
func ConstOfIDBCursorDirection(str js.Ref) uint32

//go:wasmimport plat/js/web get_IDBIndex_Name
//go:noescape
func GetIDBIndexName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_IDBIndex_Name
//go:noescape
func SetIDBIndexName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_IDBIndex_ObjectStore
//go:noescape
func GetIDBIndexObjectStore(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBIndex_KeyPath
//go:noescape
func GetIDBIndexKeyPath(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBIndex_MultiEntry
//go:noescape
func GetIDBIndexMultiEntry(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBIndex_Unique
//go:noescape
func GetIDBIndexUnique(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_Get
//go:noescape
func HasIDBIndexGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_Get
//go:noescape
func IDBIndexGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_Get
//go:noescape
func CallIDBIndexGet(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBIndex_Get
//go:noescape
func TryIDBIndexGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_GetKey
//go:noescape
func HasIDBIndexGetKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_GetKey
//go:noescape
func IDBIndexGetKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_GetKey
//go:noescape
func CallIDBIndexGetKey(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBIndex_GetKey
//go:noescape
func TryIDBIndexGetKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_GetAll
//go:noescape
func HasIDBIndexGetAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_GetAll
//go:noescape
func IDBIndexGetAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_GetAll
//go:noescape
func CallIDBIndexGetAll(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref,
	count uint32)

//go:wasmimport plat/js/web try_IDBIndex_GetAll
//go:noescape
func TryIDBIndexGetAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	count uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_GetAll1
//go:noescape
func HasIDBIndexGetAll1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_GetAll1
//go:noescape
func IDBIndexGetAll1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_GetAll1
//go:noescape
func CallIDBIndexGetAll1(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBIndex_GetAll1
//go:noescape
func TryIDBIndexGetAll1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_GetAll2
//go:noescape
func HasIDBIndexGetAll2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_GetAll2
//go:noescape
func IDBIndexGetAll2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_GetAll2
//go:noescape
func CallIDBIndexGetAll2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBIndex_GetAll2
//go:noescape
func TryIDBIndexGetAll2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_GetAllKeys
//go:noescape
func HasIDBIndexGetAllKeys(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_GetAllKeys
//go:noescape
func IDBIndexGetAllKeysFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_GetAllKeys
//go:noescape
func CallIDBIndexGetAllKeys(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref,
	count uint32)

//go:wasmimport plat/js/web try_IDBIndex_GetAllKeys
//go:noescape
func TryIDBIndexGetAllKeys(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	count uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_GetAllKeys1
//go:noescape
func HasIDBIndexGetAllKeys1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_GetAllKeys1
//go:noescape
func IDBIndexGetAllKeys1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_GetAllKeys1
//go:noescape
func CallIDBIndexGetAllKeys1(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBIndex_GetAllKeys1
//go:noescape
func TryIDBIndexGetAllKeys1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_GetAllKeys2
//go:noescape
func HasIDBIndexGetAllKeys2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_GetAllKeys2
//go:noescape
func IDBIndexGetAllKeys2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_GetAllKeys2
//go:noescape
func CallIDBIndexGetAllKeys2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBIndex_GetAllKeys2
//go:noescape
func TryIDBIndexGetAllKeys2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_Count
//go:noescape
func HasIDBIndexCount(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_Count
//go:noescape
func IDBIndexCountFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_Count
//go:noescape
func CallIDBIndexCount(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBIndex_Count
//go:noescape
func TryIDBIndexCount(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_Count1
//go:noescape
func HasIDBIndexCount1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_Count1
//go:noescape
func IDBIndexCount1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_Count1
//go:noescape
func CallIDBIndexCount1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBIndex_Count1
//go:noescape
func TryIDBIndexCount1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_OpenCursor
//go:noescape
func HasIDBIndexOpenCursor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_OpenCursor
//go:noescape
func IDBIndexOpenCursorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_OpenCursor
//go:noescape
func CallIDBIndexOpenCursor(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref,
	direction uint32)

//go:wasmimport plat/js/web try_IDBIndex_OpenCursor
//go:noescape
func TryIDBIndexOpenCursor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	direction uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_OpenCursor1
//go:noescape
func HasIDBIndexOpenCursor1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_OpenCursor1
//go:noescape
func IDBIndexOpenCursor1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_OpenCursor1
//go:noescape
func CallIDBIndexOpenCursor1(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBIndex_OpenCursor1
//go:noescape
func TryIDBIndexOpenCursor1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_OpenCursor2
//go:noescape
func HasIDBIndexOpenCursor2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_OpenCursor2
//go:noescape
func IDBIndexOpenCursor2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_OpenCursor2
//go:noescape
func CallIDBIndexOpenCursor2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBIndex_OpenCursor2
//go:noescape
func TryIDBIndexOpenCursor2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_OpenKeyCursor
//go:noescape
func HasIDBIndexOpenKeyCursor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_OpenKeyCursor
//go:noescape
func IDBIndexOpenKeyCursorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_OpenKeyCursor
//go:noescape
func CallIDBIndexOpenKeyCursor(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref,
	direction uint32)

//go:wasmimport plat/js/web try_IDBIndex_OpenKeyCursor
//go:noescape
func TryIDBIndexOpenKeyCursor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	direction uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_OpenKeyCursor1
//go:noescape
func HasIDBIndexOpenKeyCursor1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_OpenKeyCursor1
//go:noescape
func IDBIndexOpenKeyCursor1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_OpenKeyCursor1
//go:noescape
func CallIDBIndexOpenKeyCursor1(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBIndex_OpenKeyCursor1
//go:noescape
func TryIDBIndexOpenKeyCursor1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBIndex_OpenKeyCursor2
//go:noescape
func HasIDBIndexOpenKeyCursor2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBIndex_OpenKeyCursor2
//go:noescape
func IDBIndexOpenKeyCursor2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBIndex_OpenKeyCursor2
//go:noescape
func CallIDBIndexOpenKeyCursor2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBIndex_OpenKeyCursor2
//go:noescape
func TryIDBIndexOpenKeyCursor2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web store_IDBIndexParameters
//go:noescape
func IDBIndexParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IDBIndexParameters
//go:noescape
func IDBIndexParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web constof_IDBTransactionMode
//go:noescape
func ConstOfIDBTransactionMode(str js.Ref) uint32

//go:wasmimport plat/js/web constof_IDBTransactionDurability
//go:noescape
func ConstOfIDBTransactionDurability(str js.Ref) uint32

//go:wasmimport plat/js/web store_IDBTransactionOptions
//go:noescape
func IDBTransactionOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IDBTransactionOptions
//go:noescape
func IDBTransactionOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web store_IDBObjectStoreParameters
//go:noescape
func IDBObjectStoreParametersJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_IDBObjectStoreParameters
//go:noescape
func IDBObjectStoreParametersJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

//go:wasmimport plat/js/web get_IDBDatabase_Name
//go:noescape
func GetIDBDatabaseName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBDatabase_Version
//go:noescape
func GetIDBDatabaseVersion(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBDatabase_ObjectStoreNames
//go:noescape
func GetIDBDatabaseObjectStoreNames(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBDatabase_Transaction
//go:noescape
func HasIDBDatabaseTransaction(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBDatabase_Transaction
//go:noescape
func IDBDatabaseTransactionFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBDatabase_Transaction
//go:noescape
func CallIDBDatabaseTransaction(
	this js.Ref, retPtr unsafe.Pointer,
	storeNames js.Ref,
	mode uint32,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBDatabase_Transaction
//go:noescape
func TryIDBDatabaseTransaction(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	storeNames js.Ref,
	mode uint32,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBDatabase_Transaction1
//go:noescape
func HasIDBDatabaseTransaction1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBDatabase_Transaction1
//go:noescape
func IDBDatabaseTransaction1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBDatabase_Transaction1
//go:noescape
func CallIDBDatabaseTransaction1(
	this js.Ref, retPtr unsafe.Pointer,
	storeNames js.Ref,
	mode uint32)

//go:wasmimport plat/js/web try_IDBDatabase_Transaction1
//go:noescape
func TryIDBDatabaseTransaction1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	storeNames js.Ref,
	mode uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBDatabase_Transaction2
//go:noescape
func HasIDBDatabaseTransaction2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBDatabase_Transaction2
//go:noescape
func IDBDatabaseTransaction2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBDatabase_Transaction2
//go:noescape
func CallIDBDatabaseTransaction2(
	this js.Ref, retPtr unsafe.Pointer,
	storeNames js.Ref)

//go:wasmimport plat/js/web try_IDBDatabase_Transaction2
//go:noescape
func TryIDBDatabaseTransaction2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	storeNames js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBDatabase_Close
//go:noescape
func HasIDBDatabaseClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBDatabase_Close
//go:noescape
func IDBDatabaseCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBDatabase_Close
//go:noescape
func CallIDBDatabaseClose(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBDatabase_Close
//go:noescape
func TryIDBDatabaseClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBDatabase_CreateObjectStore
//go:noescape
func HasIDBDatabaseCreateObjectStore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBDatabase_CreateObjectStore
//go:noescape
func IDBDatabaseCreateObjectStoreFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBDatabase_CreateObjectStore
//go:noescape
func CallIDBDatabaseCreateObjectStore(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBDatabase_CreateObjectStore
//go:noescape
func TryIDBDatabaseCreateObjectStore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBDatabase_CreateObjectStore1
//go:noescape
func HasIDBDatabaseCreateObjectStore1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBDatabase_CreateObjectStore1
//go:noescape
func IDBDatabaseCreateObjectStore1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBDatabase_CreateObjectStore1
//go:noescape
func CallIDBDatabaseCreateObjectStore1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_IDBDatabase_CreateObjectStore1
//go:noescape
func TryIDBDatabaseCreateObjectStore1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBDatabase_DeleteObjectStore
//go:noescape
func HasIDBDatabaseDeleteObjectStore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBDatabase_DeleteObjectStore
//go:noescape
func IDBDatabaseDeleteObjectStoreFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBDatabase_DeleteObjectStore
//go:noescape
func CallIDBDatabaseDeleteObjectStore(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_IDBDatabase_DeleteObjectStore
//go:noescape
func TryIDBDatabaseDeleteObjectStore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBTransaction_ObjectStoreNames
//go:noescape
func GetIDBTransactionObjectStoreNames(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBTransaction_Mode
//go:noescape
func GetIDBTransactionMode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBTransaction_Durability
//go:noescape
func GetIDBTransactionDurability(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBTransaction_Db
//go:noescape
func GetIDBTransactionDb(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBTransaction_Error
//go:noescape
func GetIDBTransactionError(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBTransaction_ObjectStore
//go:noescape
func HasIDBTransactionObjectStore(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBTransaction_ObjectStore
//go:noescape
func IDBTransactionObjectStoreFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBTransaction_ObjectStore
//go:noescape
func CallIDBTransactionObjectStore(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_IDBTransaction_ObjectStore
//go:noescape
func TryIDBTransactionObjectStore(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBTransaction_Commit
//go:noescape
func HasIDBTransactionCommit(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBTransaction_Commit
//go:noescape
func IDBTransactionCommitFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBTransaction_Commit
//go:noescape
func CallIDBTransactionCommit(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBTransaction_Commit
//go:noescape
func TryIDBTransactionCommit(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBTransaction_Abort
//go:noescape
func HasIDBTransactionAbort(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBTransaction_Abort
//go:noescape
func IDBTransactionAbortFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBTransaction_Abort
//go:noescape
func CallIDBTransactionAbort(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBTransaction_Abort
//go:noescape
func TryIDBTransactionAbort(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBObjectStore_Name
//go:noescape
func GetIDBObjectStoreName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_IDBObjectStore_Name
//go:noescape
func SetIDBObjectStoreName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_IDBObjectStore_KeyPath
//go:noescape
func GetIDBObjectStoreKeyPath(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBObjectStore_IndexNames
//go:noescape
func GetIDBObjectStoreIndexNames(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBObjectStore_Transaction
//go:noescape
func GetIDBObjectStoreTransaction(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_IDBObjectStore_AutoIncrement
//go:noescape
func GetIDBObjectStoreAutoIncrement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_Put
//go:noescape
func HasIDBObjectStorePut(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_Put
//go:noescape
func IDBObjectStorePutFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_Put
//go:noescape
func CallIDBObjectStorePut(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref,
	key js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_Put
//go:noescape
func TryIDBObjectStorePut(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref,
	key js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_Put1
//go:noescape
func HasIDBObjectStorePut1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_Put1
//go:noescape
func IDBObjectStorePut1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_Put1
//go:noescape
func CallIDBObjectStorePut1(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_Put1
//go:noescape
func TryIDBObjectStorePut1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_Add
//go:noescape
func HasIDBObjectStoreAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_Add
//go:noescape
func IDBObjectStoreAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_Add
//go:noescape
func CallIDBObjectStoreAdd(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref,
	key js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_Add
//go:noescape
func TryIDBObjectStoreAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref,
	key js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_Add1
//go:noescape
func HasIDBObjectStoreAdd1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_Add1
//go:noescape
func IDBObjectStoreAdd1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_Add1
//go:noescape
func CallIDBObjectStoreAdd1(
	this js.Ref, retPtr unsafe.Pointer,
	value js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_Add1
//go:noescape
func TryIDBObjectStoreAdd1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	value js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_Delete
//go:noescape
func HasIDBObjectStoreDelete(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_Delete
//go:noescape
func IDBObjectStoreDeleteFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_Delete
//go:noescape
func CallIDBObjectStoreDelete(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_Delete
//go:noescape
func TryIDBObjectStoreDelete(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_Clear
//go:noescape
func HasIDBObjectStoreClear(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_Clear
//go:noescape
func IDBObjectStoreClearFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_Clear
//go:noescape
func CallIDBObjectStoreClear(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBObjectStore_Clear
//go:noescape
func TryIDBObjectStoreClear(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_Get
//go:noescape
func HasIDBObjectStoreGet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_Get
//go:noescape
func IDBObjectStoreGetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_Get
//go:noescape
func CallIDBObjectStoreGet(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_Get
//go:noescape
func TryIDBObjectStoreGet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_GetKey
//go:noescape
func HasIDBObjectStoreGetKey(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_GetKey
//go:noescape
func IDBObjectStoreGetKeyFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_GetKey
//go:noescape
func CallIDBObjectStoreGetKey(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_GetKey
//go:noescape
func TryIDBObjectStoreGetKey(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_GetAll
//go:noescape
func HasIDBObjectStoreGetAll(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_GetAll
//go:noescape
func IDBObjectStoreGetAllFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_GetAll
//go:noescape
func CallIDBObjectStoreGetAll(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref,
	count uint32)

//go:wasmimport plat/js/web try_IDBObjectStore_GetAll
//go:noescape
func TryIDBObjectStoreGetAll(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	count uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_GetAll1
//go:noescape
func HasIDBObjectStoreGetAll1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_GetAll1
//go:noescape
func IDBObjectStoreGetAll1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_GetAll1
//go:noescape
func CallIDBObjectStoreGetAll1(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_GetAll1
//go:noescape
func TryIDBObjectStoreGetAll1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_GetAll2
//go:noescape
func HasIDBObjectStoreGetAll2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_GetAll2
//go:noescape
func IDBObjectStoreGetAll2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_GetAll2
//go:noescape
func CallIDBObjectStoreGetAll2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBObjectStore_GetAll2
//go:noescape
func TryIDBObjectStoreGetAll2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_GetAllKeys
//go:noescape
func HasIDBObjectStoreGetAllKeys(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_GetAllKeys
//go:noescape
func IDBObjectStoreGetAllKeysFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_GetAllKeys
//go:noescape
func CallIDBObjectStoreGetAllKeys(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref,
	count uint32)

//go:wasmimport plat/js/web try_IDBObjectStore_GetAllKeys
//go:noescape
func TryIDBObjectStoreGetAllKeys(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	count uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_GetAllKeys1
//go:noescape
func HasIDBObjectStoreGetAllKeys1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_GetAllKeys1
//go:noescape
func IDBObjectStoreGetAllKeys1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_GetAllKeys1
//go:noescape
func CallIDBObjectStoreGetAllKeys1(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_GetAllKeys1
//go:noescape
func TryIDBObjectStoreGetAllKeys1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_GetAllKeys2
//go:noescape
func HasIDBObjectStoreGetAllKeys2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_GetAllKeys2
//go:noescape
func IDBObjectStoreGetAllKeys2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_GetAllKeys2
//go:noescape
func CallIDBObjectStoreGetAllKeys2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBObjectStore_GetAllKeys2
//go:noescape
func TryIDBObjectStoreGetAllKeys2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_Count
//go:noescape
func HasIDBObjectStoreCount(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_Count
//go:noescape
func IDBObjectStoreCountFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_Count
//go:noescape
func CallIDBObjectStoreCount(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_Count
//go:noescape
func TryIDBObjectStoreCount(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_Count1
//go:noescape
func HasIDBObjectStoreCount1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_Count1
//go:noescape
func IDBObjectStoreCount1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_Count1
//go:noescape
func CallIDBObjectStoreCount1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBObjectStore_Count1
//go:noescape
func TryIDBObjectStoreCount1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_OpenCursor
//go:noescape
func HasIDBObjectStoreOpenCursor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_OpenCursor
//go:noescape
func IDBObjectStoreOpenCursorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_OpenCursor
//go:noescape
func CallIDBObjectStoreOpenCursor(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref,
	direction uint32)

//go:wasmimport plat/js/web try_IDBObjectStore_OpenCursor
//go:noescape
func TryIDBObjectStoreOpenCursor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	direction uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_OpenCursor1
//go:noescape
func HasIDBObjectStoreOpenCursor1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_OpenCursor1
//go:noescape
func IDBObjectStoreOpenCursor1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_OpenCursor1
//go:noescape
func CallIDBObjectStoreOpenCursor1(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_OpenCursor1
//go:noescape
func TryIDBObjectStoreOpenCursor1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_OpenCursor2
//go:noescape
func HasIDBObjectStoreOpenCursor2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_OpenCursor2
//go:noescape
func IDBObjectStoreOpenCursor2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_OpenCursor2
//go:noescape
func CallIDBObjectStoreOpenCursor2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBObjectStore_OpenCursor2
//go:noescape
func TryIDBObjectStoreOpenCursor2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_OpenKeyCursor
//go:noescape
func HasIDBObjectStoreOpenKeyCursor(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_OpenKeyCursor
//go:noescape
func IDBObjectStoreOpenKeyCursorFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_OpenKeyCursor
//go:noescape
func CallIDBObjectStoreOpenKeyCursor(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref,
	direction uint32)

//go:wasmimport plat/js/web try_IDBObjectStore_OpenKeyCursor
//go:noescape
func TryIDBObjectStoreOpenKeyCursor(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref,
	direction uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_OpenKeyCursor1
//go:noescape
func HasIDBObjectStoreOpenKeyCursor1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_OpenKeyCursor1
//go:noescape
func IDBObjectStoreOpenKeyCursor1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_OpenKeyCursor1
//go:noescape
func CallIDBObjectStoreOpenKeyCursor1(
	this js.Ref, retPtr unsafe.Pointer,
	query js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_OpenKeyCursor1
//go:noescape
func TryIDBObjectStoreOpenKeyCursor1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	query js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_OpenKeyCursor2
//go:noescape
func HasIDBObjectStoreOpenKeyCursor2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_OpenKeyCursor2
//go:noescape
func IDBObjectStoreOpenKeyCursor2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_OpenKeyCursor2
//go:noescape
func CallIDBObjectStoreOpenKeyCursor2(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBObjectStore_OpenKeyCursor2
//go:noescape
func TryIDBObjectStoreOpenKeyCursor2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_Index
//go:noescape
func HasIDBObjectStoreIndex(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_Index
//go:noescape
func IDBObjectStoreIndexFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_Index
//go:noescape
func CallIDBObjectStoreIndex(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_Index
//go:noescape
func TryIDBObjectStoreIndex(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_CreateIndex
//go:noescape
func HasIDBObjectStoreCreateIndex(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_CreateIndex
//go:noescape
func IDBObjectStoreCreateIndexFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_CreateIndex
//go:noescape
func CallIDBObjectStoreCreateIndex(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	keyPath js.Ref,
	options unsafe.Pointer)

//go:wasmimport plat/js/web try_IDBObjectStore_CreateIndex
//go:noescape
func TryIDBObjectStoreCreateIndex(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	keyPath js.Ref,
	options unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_CreateIndex1
//go:noescape
func HasIDBObjectStoreCreateIndex1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_CreateIndex1
//go:noescape
func IDBObjectStoreCreateIndex1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_CreateIndex1
//go:noescape
func CallIDBObjectStoreCreateIndex1(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref,
	keyPath js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_CreateIndex1
//go:noescape
func TryIDBObjectStoreCreateIndex1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref,
	keyPath js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_IDBObjectStore_DeleteIndex
//go:noescape
func HasIDBObjectStoreDeleteIndex(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_IDBObjectStore_DeleteIndex
//go:noescape
func IDBObjectStoreDeleteIndexFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_IDBObjectStore_DeleteIndex
//go:noescape
func CallIDBObjectStoreDeleteIndex(
	this js.Ref, retPtr unsafe.Pointer,
	name js.Ref)

//go:wasmimport plat/js/web try_IDBObjectStore_DeleteIndex
//go:noescape
func TryIDBObjectStoreDeleteIndex(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	name js.Ref) (ok js.Ref)

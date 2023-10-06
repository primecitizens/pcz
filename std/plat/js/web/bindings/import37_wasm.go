// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

//go:build wasm

package bindings

import (
	"unsafe"

	"github.com/primecitizens/pcz/std/ffi/js"
)

func _() {
	var (
		_ js.Void
		_ unsafe.Pointer
	)
}

//go:wasmimport plat/js/web new_HTMLAudioElement_HTMLAudioElement
//go:noescape
func NewHTMLAudioElementByHTMLAudioElement() js.Ref

//go:wasmimport plat/js/web new_HTMLBRElement_HTMLBRElement
//go:noescape
func NewHTMLBRElementByHTMLBRElement() js.Ref

//go:wasmimport plat/js/web get_HTMLBRElement_Clear
//go:noescape
func GetHTMLBRElementClear(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLBRElement_Clear
//go:noescape
func SetHTMLBRElementClear(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLBaseElement_HTMLBaseElement
//go:noescape
func NewHTMLBaseElementByHTMLBaseElement() js.Ref

//go:wasmimport plat/js/web get_HTMLBaseElement_Href
//go:noescape
func GetHTMLBaseElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLBaseElement_Href
//go:noescape
func SetHTMLBaseElementHref(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLBaseElement_Target
//go:noescape
func GetHTMLBaseElementTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLBaseElement_Target
//go:noescape
func SetHTMLBaseElementTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLBodyElement_HTMLBodyElement
//go:noescape
func NewHTMLBodyElementByHTMLBodyElement() js.Ref

//go:wasmimport plat/js/web get_HTMLBodyElement_Text
//go:noescape
func GetHTMLBodyElementText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLBodyElement_Text
//go:noescape
func SetHTMLBodyElementText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLBodyElement_Link
//go:noescape
func GetHTMLBodyElementLink(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLBodyElement_Link
//go:noescape
func SetHTMLBodyElementLink(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLBodyElement_VLink
//go:noescape
func GetHTMLBodyElementVLink(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLBodyElement_VLink
//go:noescape
func SetHTMLBodyElementVLink(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLBodyElement_ALink
//go:noescape
func GetHTMLBodyElementALink(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLBodyElement_ALink
//go:noescape
func SetHTMLBodyElementALink(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLBodyElement_BgColor
//go:noescape
func GetHTMLBodyElementBgColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLBodyElement_BgColor
//go:noescape
func SetHTMLBodyElementBgColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLBodyElement_Background
//go:noescape
func GetHTMLBodyElementBackground(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLBodyElement_Background
//go:noescape
func SetHTMLBodyElementBackground(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLButtonElement_HTMLButtonElement
//go:noescape
func NewHTMLButtonElementByHTMLButtonElement() js.Ref

//go:wasmimport plat/js/web get_HTMLButtonElement_Disabled
//go:noescape
func GetHTMLButtonElementDisabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLButtonElement_Disabled
//go:noescape
func SetHTMLButtonElementDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLButtonElement_Form
//go:noescape
func GetHTMLButtonElementForm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLButtonElement_FormAction
//go:noescape
func GetHTMLButtonElementFormAction(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLButtonElement_FormAction
//go:noescape
func SetHTMLButtonElementFormAction(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLButtonElement_FormEnctype
//go:noescape
func GetHTMLButtonElementFormEnctype(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLButtonElement_FormEnctype
//go:noescape
func SetHTMLButtonElementFormEnctype(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLButtonElement_FormMethod
//go:noescape
func GetHTMLButtonElementFormMethod(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLButtonElement_FormMethod
//go:noescape
func SetHTMLButtonElementFormMethod(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLButtonElement_FormNoValidate
//go:noescape
func GetHTMLButtonElementFormNoValidate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLButtonElement_FormNoValidate
//go:noescape
func SetHTMLButtonElementFormNoValidate(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLButtonElement_FormTarget
//go:noescape
func GetHTMLButtonElementFormTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLButtonElement_FormTarget
//go:noescape
func SetHTMLButtonElementFormTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLButtonElement_Name
//go:noescape
func GetHTMLButtonElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLButtonElement_Name
//go:noescape
func SetHTMLButtonElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLButtonElement_Type
//go:noescape
func GetHTMLButtonElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLButtonElement_Type
//go:noescape
func SetHTMLButtonElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLButtonElement_Value
//go:noescape
func GetHTMLButtonElementValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLButtonElement_Value
//go:noescape
func SetHTMLButtonElementValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLButtonElement_WillValidate
//go:noescape
func GetHTMLButtonElementWillValidate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLButtonElement_Validity
//go:noescape
func GetHTMLButtonElementValidity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLButtonElement_ValidationMessage
//go:noescape
func GetHTMLButtonElementValidationMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLButtonElement_Labels
//go:noescape
func GetHTMLButtonElementLabels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLButtonElement_PopoverTargetElement
//go:noescape
func GetHTMLButtonElementPopoverTargetElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLButtonElement_PopoverTargetElement
//go:noescape
func SetHTMLButtonElementPopoverTargetElement(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLButtonElement_PopoverTargetAction
//go:noescape
func GetHTMLButtonElementPopoverTargetAction(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLButtonElement_PopoverTargetAction
//go:noescape
func SetHTMLButtonElementPopoverTargetAction(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLButtonElement_CheckValidity
//go:noescape
func HasHTMLButtonElementCheckValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLButtonElement_CheckValidity
//go:noescape
func HTMLButtonElementCheckValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLButtonElement_CheckValidity
//go:noescape
func CallHTMLButtonElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLButtonElement_CheckValidity
//go:noescape
func TryHTMLButtonElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLButtonElement_ReportValidity
//go:noescape
func HasHTMLButtonElementReportValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLButtonElement_ReportValidity
//go:noescape
func HTMLButtonElementReportValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLButtonElement_ReportValidity
//go:noescape
func CallHTMLButtonElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLButtonElement_ReportValidity
//go:noescape
func TryHTMLButtonElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLButtonElement_SetCustomValidity
//go:noescape
func HasHTMLButtonElementSetCustomValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLButtonElement_SetCustomValidity
//go:noescape
func HTMLButtonElementSetCustomValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLButtonElement_SetCustomValidity
//go:noescape
func CallHTMLButtonElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer,
	err js.Ref)

//go:wasmimport plat/js/web try_HTMLButtonElement_SetCustomValidity
//go:noescape
func TryHTMLButtonElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	err js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLDListElement_HTMLDListElement
//go:noescape
func NewHTMLDListElementByHTMLDListElement() js.Ref

//go:wasmimport plat/js/web get_HTMLDListElement_Compact
//go:noescape
func GetHTMLDListElementCompact(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLDListElement_Compact
//go:noescape
func SetHTMLDListElementCompact(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLDataElement_HTMLDataElement
//go:noescape
func NewHTMLDataElementByHTMLDataElement() js.Ref

//go:wasmimport plat/js/web get_HTMLDataElement_Value
//go:noescape
func GetHTMLDataElementValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLDataElement_Value
//go:noescape
func SetHTMLDataElementValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLDataListElement_HTMLDataListElement
//go:noescape
func NewHTMLDataListElementByHTMLDataListElement() js.Ref

//go:wasmimport plat/js/web get_HTMLDataListElement_Options
//go:noescape
func GetHTMLDataListElementOptions(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLDetailsElement_HTMLDetailsElement
//go:noescape
func NewHTMLDetailsElementByHTMLDetailsElement() js.Ref

//go:wasmimport plat/js/web get_HTMLDetailsElement_Open
//go:noescape
func GetHTMLDetailsElementOpen(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLDetailsElement_Open
//go:noescape
func SetHTMLDetailsElementOpen(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLDialogElement_HTMLDialogElement
//go:noescape
func NewHTMLDialogElementByHTMLDialogElement() js.Ref

//go:wasmimport plat/js/web get_HTMLDialogElement_Open
//go:noescape
func GetHTMLDialogElementOpen(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLDialogElement_Open
//go:noescape
func SetHTMLDialogElementOpen(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLDialogElement_ReturnValue
//go:noescape
func GetHTMLDialogElementReturnValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLDialogElement_ReturnValue
//go:noescape
func SetHTMLDialogElementReturnValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLDialogElement_Show
//go:noescape
func HasHTMLDialogElementShow(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLDialogElement_Show
//go:noescape
func HTMLDialogElementShowFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLDialogElement_Show
//go:noescape
func CallHTMLDialogElementShow(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLDialogElement_Show
//go:noescape
func TryHTMLDialogElementShow(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLDialogElement_ShowModal
//go:noescape
func HasHTMLDialogElementShowModal(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLDialogElement_ShowModal
//go:noescape
func HTMLDialogElementShowModalFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLDialogElement_ShowModal
//go:noescape
func CallHTMLDialogElementShowModal(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLDialogElement_ShowModal
//go:noescape
func TryHTMLDialogElementShowModal(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLDialogElement_Close
//go:noescape
func HasHTMLDialogElementClose(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLDialogElement_Close
//go:noescape
func HTMLDialogElementCloseFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLDialogElement_Close
//go:noescape
func CallHTMLDialogElementClose(
	this js.Ref, retPtr unsafe.Pointer,
	returnValue js.Ref)

//go:wasmimport plat/js/web try_HTMLDialogElement_Close
//go:noescape
func TryHTMLDialogElementClose(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	returnValue js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLDialogElement_Close1
//go:noescape
func HasHTMLDialogElementClose1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLDialogElement_Close1
//go:noescape
func HTMLDialogElementClose1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLDialogElement_Close1
//go:noescape
func CallHTMLDialogElementClose1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLDialogElement_Close1
//go:noescape
func TryHTMLDialogElementClose1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLDirectoryElement_HTMLDirectoryElement
//go:noescape
func NewHTMLDirectoryElementByHTMLDirectoryElement() js.Ref

//go:wasmimport plat/js/web get_HTMLDirectoryElement_Compact
//go:noescape
func GetHTMLDirectoryElementCompact(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLDirectoryElement_Compact
//go:noescape
func SetHTMLDirectoryElementCompact(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLDivElement_HTMLDivElement
//go:noescape
func NewHTMLDivElementByHTMLDivElement() js.Ref

//go:wasmimport plat/js/web get_HTMLDivElement_Align
//go:noescape
func GetHTMLDivElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLDivElement_Align
//go:noescape
func SetHTMLDivElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLEmbedElement_HTMLEmbedElement
//go:noescape
func NewHTMLEmbedElementByHTMLEmbedElement() js.Ref

//go:wasmimport plat/js/web get_HTMLEmbedElement_Src
//go:noescape
func GetHTMLEmbedElementSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLEmbedElement_Src
//go:noescape
func SetHTMLEmbedElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLEmbedElement_Type
//go:noescape
func GetHTMLEmbedElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLEmbedElement_Type
//go:noescape
func SetHTMLEmbedElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLEmbedElement_Width
//go:noescape
func GetHTMLEmbedElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLEmbedElement_Width
//go:noescape
func SetHTMLEmbedElementWidth(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLEmbedElement_Height
//go:noescape
func GetHTMLEmbedElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLEmbedElement_Height
//go:noescape
func SetHTMLEmbedElementHeight(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLEmbedElement_Align
//go:noescape
func GetHTMLEmbedElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLEmbedElement_Align
//go:noescape
func SetHTMLEmbedElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLEmbedElement_Name
//go:noescape
func GetHTMLEmbedElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLEmbedElement_Name
//go:noescape
func SetHTMLEmbedElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLEmbedElement_GetSVGDocument
//go:noescape
func HasHTMLEmbedElementGetSVGDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLEmbedElement_GetSVGDocument
//go:noescape
func HTMLEmbedElementGetSVGDocumentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLEmbedElement_GetSVGDocument
//go:noescape
func CallHTMLEmbedElementGetSVGDocument(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLEmbedElement_GetSVGDocument
//go:noescape
func TryHTMLEmbedElementGetSVGDocument(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLFencedFrameElement_HTMLFencedFrameElement
//go:noescape
func NewHTMLFencedFrameElementByHTMLFencedFrameElement() js.Ref

//go:wasmimport plat/js/web get_HTMLFencedFrameElement_Config
//go:noescape
func GetHTMLFencedFrameElementConfig(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFencedFrameElement_Config
//go:noescape
func SetHTMLFencedFrameElementConfig(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFencedFrameElement_Width
//go:noescape
func GetHTMLFencedFrameElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFencedFrameElement_Width
//go:noescape
func SetHTMLFencedFrameElementWidth(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFencedFrameElement_Height
//go:noescape
func GetHTMLFencedFrameElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFencedFrameElement_Height
//go:noescape
func SetHTMLFencedFrameElementHeight(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFencedFrameElement_Allow
//go:noescape
func GetHTMLFencedFrameElementAllow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFencedFrameElement_Allow
//go:noescape
func SetHTMLFencedFrameElementAllow(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLFieldSetElement_HTMLFieldSetElement
//go:noescape
func NewHTMLFieldSetElementByHTMLFieldSetElement() js.Ref

//go:wasmimport plat/js/web get_HTMLFieldSetElement_Disabled
//go:noescape
func GetHTMLFieldSetElementDisabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFieldSetElement_Disabled
//go:noescape
func SetHTMLFieldSetElementDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFieldSetElement_Form
//go:noescape
func GetHTMLFieldSetElementForm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLFieldSetElement_Name
//go:noescape
func GetHTMLFieldSetElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFieldSetElement_Name
//go:noescape
func SetHTMLFieldSetElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFieldSetElement_Type
//go:noescape
func GetHTMLFieldSetElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLFieldSetElement_Elements
//go:noescape
func GetHTMLFieldSetElementElements(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLFieldSetElement_WillValidate
//go:noescape
func GetHTMLFieldSetElementWillValidate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLFieldSetElement_Validity
//go:noescape
func GetHTMLFieldSetElementValidity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLFieldSetElement_ValidationMessage
//go:noescape
func GetHTMLFieldSetElementValidationMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLFieldSetElement_CheckValidity
//go:noescape
func HasHTMLFieldSetElementCheckValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFieldSetElement_CheckValidity
//go:noescape
func HTMLFieldSetElementCheckValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLFieldSetElement_CheckValidity
//go:noescape
func CallHTMLFieldSetElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLFieldSetElement_CheckValidity
//go:noescape
func TryHTMLFieldSetElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLFieldSetElement_ReportValidity
//go:noescape
func HasHTMLFieldSetElementReportValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFieldSetElement_ReportValidity
//go:noescape
func HTMLFieldSetElementReportValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLFieldSetElement_ReportValidity
//go:noescape
func CallHTMLFieldSetElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLFieldSetElement_ReportValidity
//go:noescape
func TryHTMLFieldSetElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLFieldSetElement_SetCustomValidity
//go:noescape
func HasHTMLFieldSetElementSetCustomValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLFieldSetElement_SetCustomValidity
//go:noescape
func HTMLFieldSetElementSetCustomValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLFieldSetElement_SetCustomValidity
//go:noescape
func CallHTMLFieldSetElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer,
	err js.Ref)

//go:wasmimport plat/js/web try_HTMLFieldSetElement_SetCustomValidity
//go:noescape
func TryHTMLFieldSetElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	err js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLFontElement_HTMLFontElement
//go:noescape
func NewHTMLFontElementByHTMLFontElement() js.Ref

//go:wasmimport plat/js/web get_HTMLFontElement_Color
//go:noescape
func GetHTMLFontElementColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFontElement_Color
//go:noescape
func SetHTMLFontElementColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFontElement_Face
//go:noescape
func GetHTMLFontElementFace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFontElement_Face
//go:noescape
func SetHTMLFontElementFace(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFontElement_Size
//go:noescape
func GetHTMLFontElementSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFontElement_Size
//go:noescape
func SetHTMLFontElementSize(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLFrameElement_HTMLFrameElement
//go:noescape
func NewHTMLFrameElementByHTMLFrameElement() js.Ref

//go:wasmimport plat/js/web get_HTMLFrameElement_Name
//go:noescape
func GetHTMLFrameElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFrameElement_Name
//go:noescape
func SetHTMLFrameElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFrameElement_Scrolling
//go:noescape
func GetHTMLFrameElementScrolling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFrameElement_Scrolling
//go:noescape
func SetHTMLFrameElementScrolling(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFrameElement_Src
//go:noescape
func GetHTMLFrameElementSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFrameElement_Src
//go:noescape
func SetHTMLFrameElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFrameElement_FrameBorder
//go:noescape
func GetHTMLFrameElementFrameBorder(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFrameElement_FrameBorder
//go:noescape
func SetHTMLFrameElementFrameBorder(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFrameElement_LongDesc
//go:noescape
func GetHTMLFrameElementLongDesc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFrameElement_LongDesc
//go:noescape
func SetHTMLFrameElementLongDesc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFrameElement_NoResize
//go:noescape
func GetHTMLFrameElementNoResize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFrameElement_NoResize
//go:noescape
func SetHTMLFrameElementNoResize(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFrameElement_ContentDocument
//go:noescape
func GetHTMLFrameElementContentDocument(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLFrameElement_ContentWindow
//go:noescape
func GetHTMLFrameElementContentWindow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLFrameElement_MarginHeight
//go:noescape
func GetHTMLFrameElementMarginHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFrameElement_MarginHeight
//go:noescape
func SetHTMLFrameElementMarginHeight(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFrameElement_MarginWidth
//go:noescape
func GetHTMLFrameElementMarginWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFrameElement_MarginWidth
//go:noescape
func SetHTMLFrameElementMarginWidth(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLFrameSetElement_HTMLFrameSetElement
//go:noescape
func NewHTMLFrameSetElementByHTMLFrameSetElement() js.Ref

//go:wasmimport plat/js/web get_HTMLFrameSetElement_Cols
//go:noescape
func GetHTMLFrameSetElementCols(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFrameSetElement_Cols
//go:noescape
func SetHTMLFrameSetElementCols(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLFrameSetElement_Rows
//go:noescape
func GetHTMLFrameSetElementRows(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLFrameSetElement_Rows
//go:noescape
func SetHTMLFrameSetElementRows(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLHRElement_HTMLHRElement
//go:noescape
func NewHTMLHRElementByHTMLHRElement() js.Ref

//go:wasmimport plat/js/web get_HTMLHRElement_Align
//go:noescape
func GetHTMLHRElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLHRElement_Align
//go:noescape
func SetHTMLHRElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLHRElement_Color
//go:noescape
func GetHTMLHRElementColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLHRElement_Color
//go:noescape
func SetHTMLHRElementColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLHRElement_NoShade
//go:noescape
func GetHTMLHRElementNoShade(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLHRElement_NoShade
//go:noescape
func SetHTMLHRElementNoShade(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLHRElement_Size
//go:noescape
func GetHTMLHRElementSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLHRElement_Size
//go:noescape
func SetHTMLHRElementSize(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLHRElement_Width
//go:noescape
func GetHTMLHRElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLHRElement_Width
//go:noescape
func SetHTMLHRElementWidth(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLHeadingElement_HTMLHeadingElement
//go:noescape
func NewHTMLHeadingElementByHTMLHeadingElement() js.Ref

//go:wasmimport plat/js/web get_HTMLHeadingElement_Align
//go:noescape
func GetHTMLHeadingElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLHeadingElement_Align
//go:noescape
func SetHTMLHeadingElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLHtmlElement_HTMLHtmlElement
//go:noescape
func NewHTMLHtmlElementByHTMLHtmlElement() js.Ref

//go:wasmimport plat/js/web get_HTMLHtmlElement_Version
//go:noescape
func GetHTMLHtmlElementVersion(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLHtmlElement_Version
//go:noescape
func SetHTMLHtmlElementVersion(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLIFrameElement_HTMLIFrameElement
//go:noescape
func NewHTMLIFrameElementByHTMLIFrameElement() js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_Src
//go:noescape
func GetHTMLIFrameElementSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_Src
//go:noescape
func SetHTMLIFrameElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_Srcdoc
//go:noescape
func GetHTMLIFrameElementSrcdoc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_Srcdoc
//go:noescape
func SetHTMLIFrameElementSrcdoc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_Name
//go:noescape
func GetHTMLIFrameElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_Name
//go:noescape
func SetHTMLIFrameElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_Sandbox
//go:noescape
func GetHTMLIFrameElementSandbox(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLIFrameElement_Allow
//go:noescape
func GetHTMLIFrameElementAllow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_Allow
//go:noescape
func SetHTMLIFrameElementAllow(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_AllowFullscreen
//go:noescape
func GetHTMLIFrameElementAllowFullscreen(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_AllowFullscreen
//go:noescape
func SetHTMLIFrameElementAllowFullscreen(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_Width
//go:noescape
func GetHTMLIFrameElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_Width
//go:noescape
func SetHTMLIFrameElementWidth(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_Height
//go:noescape
func GetHTMLIFrameElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_Height
//go:noescape
func SetHTMLIFrameElementHeight(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_ReferrerPolicy
//go:noescape
func GetHTMLIFrameElementReferrerPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_ReferrerPolicy
//go:noescape
func SetHTMLIFrameElementReferrerPolicy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_Loading
//go:noescape
func GetHTMLIFrameElementLoading(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_Loading
//go:noescape
func SetHTMLIFrameElementLoading(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_ContentDocument
//go:noescape
func GetHTMLIFrameElementContentDocument(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLIFrameElement_ContentWindow
//go:noescape
func GetHTMLIFrameElementContentWindow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLIFrameElement_PermissionsPolicy
//go:noescape
func GetHTMLIFrameElementPermissionsPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLIFrameElement_Csp
//go:noescape
func GetHTMLIFrameElementCsp(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_Csp
//go:noescape
func SetHTMLIFrameElementCsp(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_PrivateToken
//go:noescape
func GetHTMLIFrameElementPrivateToken(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_PrivateToken
//go:noescape
func SetHTMLIFrameElementPrivateToken(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_Align
//go:noescape
func GetHTMLIFrameElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_Align
//go:noescape
func SetHTMLIFrameElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_Scrolling
//go:noescape
func GetHTMLIFrameElementScrolling(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_Scrolling
//go:noescape
func SetHTMLIFrameElementScrolling(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_FrameBorder
//go:noescape
func GetHTMLIFrameElementFrameBorder(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_FrameBorder
//go:noescape
func SetHTMLIFrameElementFrameBorder(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_LongDesc
//go:noescape
func GetHTMLIFrameElementLongDesc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_LongDesc
//go:noescape
func SetHTMLIFrameElementLongDesc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_MarginHeight
//go:noescape
func GetHTMLIFrameElementMarginHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_MarginHeight
//go:noescape
func SetHTMLIFrameElementMarginHeight(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLIFrameElement_MarginWidth
//go:noescape
func GetHTMLIFrameElementMarginWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLIFrameElement_MarginWidth
//go:noescape
func SetHTMLIFrameElementMarginWidth(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLIFrameElement_GetSVGDocument
//go:noescape
func HasHTMLIFrameElementGetSVGDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLIFrameElement_GetSVGDocument
//go:noescape
func HTMLIFrameElementGetSVGDocumentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLIFrameElement_GetSVGDocument
//go:noescape
func CallHTMLIFrameElementGetSVGDocument(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLIFrameElement_GetSVGDocument
//go:noescape
func TryHTMLIFrameElementGetSVGDocument(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web constof_SelectionMode
//go:noescape
func ConstOfSelectionMode(str js.Ref) uint32

//go:wasmimport plat/js/web new_HTMLInputElement_HTMLInputElement
//go:noescape
func NewHTMLInputElementByHTMLInputElement() js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Accept
//go:noescape
func GetHTMLInputElementAccept(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Accept
//go:noescape
func SetHTMLInputElementAccept(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Alt
//go:noescape
func GetHTMLInputElementAlt(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Alt
//go:noescape
func SetHTMLInputElementAlt(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Autocomplete
//go:noescape
func GetHTMLInputElementAutocomplete(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Autocomplete
//go:noescape
func SetHTMLInputElementAutocomplete(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_DefaultChecked
//go:noescape
func GetHTMLInputElementDefaultChecked(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_DefaultChecked
//go:noescape
func SetHTMLInputElementDefaultChecked(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Checked
//go:noescape
func GetHTMLInputElementChecked(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Checked
//go:noescape
func SetHTMLInputElementChecked(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_DirName
//go:noescape
func GetHTMLInputElementDirName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_DirName
//go:noescape
func SetHTMLInputElementDirName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Disabled
//go:noescape
func GetHTMLInputElementDisabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Disabled
//go:noescape
func SetHTMLInputElementDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Form
//go:noescape
func GetHTMLInputElementForm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLInputElement_Files
//go:noescape
func GetHTMLInputElementFiles(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Files
//go:noescape
func SetHTMLInputElementFiles(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_FormAction
//go:noescape
func GetHTMLInputElementFormAction(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_FormAction
//go:noescape
func SetHTMLInputElementFormAction(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_FormEnctype
//go:noescape
func GetHTMLInputElementFormEnctype(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_FormEnctype
//go:noescape
func SetHTMLInputElementFormEnctype(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_FormMethod
//go:noescape
func GetHTMLInputElementFormMethod(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_FormMethod
//go:noescape
func SetHTMLInputElementFormMethod(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_FormNoValidate
//go:noescape
func GetHTMLInputElementFormNoValidate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_FormNoValidate
//go:noescape
func SetHTMLInputElementFormNoValidate(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_FormTarget
//go:noescape
func GetHTMLInputElementFormTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_FormTarget
//go:noescape
func SetHTMLInputElementFormTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Height
//go:noescape
func GetHTMLInputElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Height
//go:noescape
func SetHTMLInputElementHeight(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Indeterminate
//go:noescape
func GetHTMLInputElementIndeterminate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Indeterminate
//go:noescape
func SetHTMLInputElementIndeterminate(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_List
//go:noescape
func GetHTMLInputElementList(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLInputElement_Max
//go:noescape
func GetHTMLInputElementMax(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Max
//go:noescape
func SetHTMLInputElementMax(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_MaxLength
//go:noescape
func GetHTMLInputElementMaxLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_MaxLength
//go:noescape
func SetHTMLInputElementMaxLength(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Min
//go:noescape
func GetHTMLInputElementMin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Min
//go:noescape
func SetHTMLInputElementMin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_MinLength
//go:noescape
func GetHTMLInputElementMinLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_MinLength
//go:noescape
func SetHTMLInputElementMinLength(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Multiple
//go:noescape
func GetHTMLInputElementMultiple(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Multiple
//go:noescape
func SetHTMLInputElementMultiple(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Name
//go:noescape
func GetHTMLInputElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Name
//go:noescape
func SetHTMLInputElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Pattern
//go:noescape
func GetHTMLInputElementPattern(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Pattern
//go:noescape
func SetHTMLInputElementPattern(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Placeholder
//go:noescape
func GetHTMLInputElementPlaceholder(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Placeholder
//go:noescape
func SetHTMLInputElementPlaceholder(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_ReadOnly
//go:noescape
func GetHTMLInputElementReadOnly(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_ReadOnly
//go:noescape
func SetHTMLInputElementReadOnly(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Required
//go:noescape
func GetHTMLInputElementRequired(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Required
//go:noescape
func SetHTMLInputElementRequired(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Size
//go:noescape
func GetHTMLInputElementSize(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Size
//go:noescape
func SetHTMLInputElementSize(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Src
//go:noescape
func GetHTMLInputElementSrc(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Src
//go:noescape
func SetHTMLInputElementSrc(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Step
//go:noescape
func GetHTMLInputElementStep(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Step
//go:noescape
func SetHTMLInputElementStep(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Type
//go:noescape
func GetHTMLInputElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Type
//go:noescape
func SetHTMLInputElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_DefaultValue
//go:noescape
func GetHTMLInputElementDefaultValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_DefaultValue
//go:noescape
func SetHTMLInputElementDefaultValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Value
//go:noescape
func GetHTMLInputElementValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Value
//go:noescape
func SetHTMLInputElementValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_ValueAsDate
//go:noescape
func GetHTMLInputElementValueAsDate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_ValueAsDate
//go:noescape
func SetHTMLInputElementValueAsDate(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_ValueAsNumber
//go:noescape
func GetHTMLInputElementValueAsNumber(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_ValueAsNumber
//go:noescape
func SetHTMLInputElementValueAsNumber(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Width
//go:noescape
func GetHTMLInputElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Width
//go:noescape
func SetHTMLInputElementWidth(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_WillValidate
//go:noescape
func GetHTMLInputElementWillValidate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLInputElement_Validity
//go:noescape
func GetHTMLInputElementValidity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLInputElement_ValidationMessage
//go:noescape
func GetHTMLInputElementValidationMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLInputElement_Labels
//go:noescape
func GetHTMLInputElementLabels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLInputElement_SelectionStart
//go:noescape
func GetHTMLInputElementSelectionStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_SelectionStart
//go:noescape
func SetHTMLInputElementSelectionStart(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_SelectionEnd
//go:noescape
func GetHTMLInputElementSelectionEnd(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_SelectionEnd
//go:noescape
func SetHTMLInputElementSelectionEnd(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_SelectionDirection
//go:noescape
func GetHTMLInputElementSelectionDirection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_SelectionDirection
//go:noescape
func SetHTMLInputElementSelectionDirection(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Align
//go:noescape
func GetHTMLInputElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Align
//go:noescape
func SetHTMLInputElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_UseMap
//go:noescape
func GetHTMLInputElementUseMap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_UseMap
//go:noescape
func SetHTMLInputElementUseMap(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_Webkitdirectory
//go:noescape
func GetHTMLInputElementWebkitdirectory(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Webkitdirectory
//go:noescape
func SetHTMLInputElementWebkitdirectory(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_WebkitEntries
//go:noescape
func GetHTMLInputElementWebkitEntries(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLInputElement_Capture
//go:noescape
func GetHTMLInputElementCapture(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_Capture
//go:noescape
func SetHTMLInputElementCapture(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_PopoverTargetElement
//go:noescape
func GetHTMLInputElementPopoverTargetElement(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_PopoverTargetElement
//go:noescape
func SetHTMLInputElementPopoverTargetElement(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLInputElement_PopoverTargetAction
//go:noescape
func GetHTMLInputElementPopoverTargetAction(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLInputElement_PopoverTargetAction
//go:noescape
func SetHTMLInputElementPopoverTargetAction(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLInputElement_StepUp
//go:noescape
func HasHTMLInputElementStepUp(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_StepUp
//go:noescape
func HTMLInputElementStepUpFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_StepUp
//go:noescape
func CallHTMLInputElementStepUp(
	this js.Ref, retPtr unsafe.Pointer,
	n int32)

//go:wasmimport plat/js/web try_HTMLInputElement_StepUp
//go:noescape
func TryHTMLInputElementStepUp(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	n int32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_StepUp1
//go:noescape
func HasHTMLInputElementStepUp1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_StepUp1
//go:noescape
func HTMLInputElementStepUp1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_StepUp1
//go:noescape
func CallHTMLInputElementStepUp1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLInputElement_StepUp1
//go:noescape
func TryHTMLInputElementStepUp1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_StepDown
//go:noescape
func HasHTMLInputElementStepDown(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_StepDown
//go:noescape
func HTMLInputElementStepDownFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_StepDown
//go:noescape
func CallHTMLInputElementStepDown(
	this js.Ref, retPtr unsafe.Pointer,
	n int32)

//go:wasmimport plat/js/web try_HTMLInputElement_StepDown
//go:noescape
func TryHTMLInputElementStepDown(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	n int32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_StepDown1
//go:noescape
func HasHTMLInputElementStepDown1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_StepDown1
//go:noescape
func HTMLInputElementStepDown1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_StepDown1
//go:noescape
func CallHTMLInputElementStepDown1(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLInputElement_StepDown1
//go:noescape
func TryHTMLInputElementStepDown1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_CheckValidity
//go:noescape
func HasHTMLInputElementCheckValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_CheckValidity
//go:noescape
func HTMLInputElementCheckValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_CheckValidity
//go:noescape
func CallHTMLInputElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLInputElement_CheckValidity
//go:noescape
func TryHTMLInputElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_ReportValidity
//go:noescape
func HasHTMLInputElementReportValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_ReportValidity
//go:noescape
func HTMLInputElementReportValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_ReportValidity
//go:noescape
func CallHTMLInputElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLInputElement_ReportValidity
//go:noescape
func TryHTMLInputElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_SetCustomValidity
//go:noescape
func HasHTMLInputElementSetCustomValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_SetCustomValidity
//go:noescape
func HTMLInputElementSetCustomValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_SetCustomValidity
//go:noescape
func CallHTMLInputElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer,
	err js.Ref)

//go:wasmimport plat/js/web try_HTMLInputElement_SetCustomValidity
//go:noescape
func TryHTMLInputElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	err js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_Select
//go:noescape
func HasHTMLInputElementSelect(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_Select
//go:noescape
func HTMLInputElementSelectFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_Select
//go:noescape
func CallHTMLInputElementSelect(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLInputElement_Select
//go:noescape
func TryHTMLInputElementSelect(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_SetRangeText
//go:noescape
func HasHTMLInputElementSetRangeText(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_SetRangeText
//go:noescape
func HTMLInputElementSetRangeTextFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_SetRangeText
//go:noescape
func CallHTMLInputElementSetRangeText(
	this js.Ref, retPtr unsafe.Pointer,
	replacement js.Ref)

//go:wasmimport plat/js/web try_HTMLInputElement_SetRangeText
//go:noescape
func TryHTMLInputElementSetRangeText(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	replacement js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_SetRangeText1
//go:noescape
func HasHTMLInputElementSetRangeText1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_SetRangeText1
//go:noescape
func HTMLInputElementSetRangeText1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_SetRangeText1
//go:noescape
func CallHTMLInputElementSetRangeText1(
	this js.Ref, retPtr unsafe.Pointer,
	replacement js.Ref,
	start uint32,
	end uint32,
	selectionMode uint32)

//go:wasmimport plat/js/web try_HTMLInputElement_SetRangeText1
//go:noescape
func TryHTMLInputElementSetRangeText1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	replacement js.Ref,
	start uint32,
	end uint32,
	selectionMode uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_SetRangeText2
//go:noescape
func HasHTMLInputElementSetRangeText2(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_SetRangeText2
//go:noescape
func HTMLInputElementSetRangeText2Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_SetRangeText2
//go:noescape
func CallHTMLInputElementSetRangeText2(
	this js.Ref, retPtr unsafe.Pointer,
	replacement js.Ref,
	start uint32,
	end uint32)

//go:wasmimport plat/js/web try_HTMLInputElement_SetRangeText2
//go:noescape
func TryHTMLInputElementSetRangeText2(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	replacement js.Ref,
	start uint32,
	end uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_SetSelectionRange
//go:noescape
func HasHTMLInputElementSetSelectionRange(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_SetSelectionRange
//go:noescape
func HTMLInputElementSetSelectionRangeFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_SetSelectionRange
//go:noescape
func CallHTMLInputElementSetSelectionRange(
	this js.Ref, retPtr unsafe.Pointer,
	start uint32,
	end uint32,
	direction js.Ref)

//go:wasmimport plat/js/web try_HTMLInputElement_SetSelectionRange
//go:noescape
func TryHTMLInputElementSetSelectionRange(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	start uint32,
	end uint32,
	direction js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_SetSelectionRange1
//go:noescape
func HasHTMLInputElementSetSelectionRange1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_SetSelectionRange1
//go:noescape
func HTMLInputElementSetSelectionRange1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_SetSelectionRange1
//go:noescape
func CallHTMLInputElementSetSelectionRange1(
	this js.Ref, retPtr unsafe.Pointer,
	start uint32,
	end uint32)

//go:wasmimport plat/js/web try_HTMLInputElement_SetSelectionRange1
//go:noescape
func TryHTMLInputElementSetSelectionRange1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	start uint32,
	end uint32) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLInputElement_ShowPicker
//go:noescape
func HasHTMLInputElementShowPicker(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLInputElement_ShowPicker
//go:noescape
func HTMLInputElementShowPickerFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLInputElement_ShowPicker
//go:noescape
func CallHTMLInputElementShowPicker(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLInputElement_ShowPicker
//go:noescape
func TryHTMLInputElementShowPicker(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLLIElement_HTMLLIElement
//go:noescape
func NewHTMLLIElementByHTMLLIElement() js.Ref

//go:wasmimport plat/js/web get_HTMLLIElement_Value
//go:noescape
func GetHTMLLIElementValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLIElement_Value
//go:noescape
func SetHTMLLIElementValue(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLIElement_Type
//go:noescape
func GetHTMLLIElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLIElement_Type
//go:noescape
func SetHTMLLIElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLLabelElement_HTMLLabelElement
//go:noescape
func NewHTMLLabelElementByHTMLLabelElement() js.Ref

//go:wasmimport plat/js/web get_HTMLLabelElement_Form
//go:noescape
func GetHTMLLabelElementForm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLLabelElement_HtmlFor
//go:noescape
func GetHTMLLabelElementHtmlFor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLabelElement_HtmlFor
//go:noescape
func SetHTMLLabelElementHtmlFor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLabelElement_Control
//go:noescape
func GetHTMLLabelElementControl(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLLegendElement_HTMLLegendElement
//go:noescape
func NewHTMLLegendElementByHTMLLegendElement() js.Ref

//go:wasmimport plat/js/web get_HTMLLegendElement_Form
//go:noescape
func GetHTMLLegendElementForm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLLegendElement_Align
//go:noescape
func GetHTMLLegendElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLegendElement_Align
//go:noescape
func SetHTMLLegendElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLLinkElement_HTMLLinkElement
//go:noescape
func NewHTMLLinkElementByHTMLLinkElement() js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_Href
//go:noescape
func GetHTMLLinkElementHref(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_Href
//go:noescape
func SetHTMLLinkElementHref(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_CrossOrigin
//go:noescape
func GetHTMLLinkElementCrossOrigin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_CrossOrigin
//go:noescape
func SetHTMLLinkElementCrossOrigin(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_Rel
//go:noescape
func GetHTMLLinkElementRel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_Rel
//go:noescape
func SetHTMLLinkElementRel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_As
//go:noescape
func GetHTMLLinkElementAs(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_As
//go:noescape
func SetHTMLLinkElementAs(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_RelList
//go:noescape
func GetHTMLLinkElementRelList(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLLinkElement_Media
//go:noescape
func GetHTMLLinkElementMedia(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_Media
//go:noescape
func SetHTMLLinkElementMedia(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_Integrity
//go:noescape
func GetHTMLLinkElementIntegrity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_Integrity
//go:noescape
func SetHTMLLinkElementIntegrity(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_Hreflang
//go:noescape
func GetHTMLLinkElementHreflang(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_Hreflang
//go:noescape
func SetHTMLLinkElementHreflang(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_Type
//go:noescape
func GetHTMLLinkElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_Type
//go:noescape
func SetHTMLLinkElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_Sizes
//go:noescape
func GetHTMLLinkElementSizes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLLinkElement_ImageSrcset
//go:noescape
func GetHTMLLinkElementImageSrcset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_ImageSrcset
//go:noescape
func SetHTMLLinkElementImageSrcset(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_ImageSizes
//go:noescape
func GetHTMLLinkElementImageSizes(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_ImageSizes
//go:noescape
func SetHTMLLinkElementImageSizes(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_ReferrerPolicy
//go:noescape
func GetHTMLLinkElementReferrerPolicy(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_ReferrerPolicy
//go:noescape
func SetHTMLLinkElementReferrerPolicy(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_Blocking
//go:noescape
func GetHTMLLinkElementBlocking(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLLinkElement_Disabled
//go:noescape
func GetHTMLLinkElementDisabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_Disabled
//go:noescape
func SetHTMLLinkElementDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_FetchPriority
//go:noescape
func GetHTMLLinkElementFetchPriority(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_FetchPriority
//go:noescape
func SetHTMLLinkElementFetchPriority(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_Charset
//go:noescape
func GetHTMLLinkElementCharset(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_Charset
//go:noescape
func SetHTMLLinkElementCharset(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_Rev
//go:noescape
func GetHTMLLinkElementRev(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_Rev
//go:noescape
func SetHTMLLinkElementRev(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_Target
//go:noescape
func GetHTMLLinkElementTarget(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLLinkElement_Target
//go:noescape
func SetHTMLLinkElementTarget(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLLinkElement_Sheet
//go:noescape
func GetHTMLLinkElementSheet(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLMapElement_HTMLMapElement
//go:noescape
func NewHTMLMapElementByHTMLMapElement() js.Ref

//go:wasmimport plat/js/web get_HTMLMapElement_Name
//go:noescape
func GetHTMLMapElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMapElement_Name
//go:noescape
func SetHTMLMapElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMapElement_Areas
//go:noescape
func GetHTMLMapElementAreas(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLMarqueeElement_HTMLMarqueeElement
//go:noescape
func NewHTMLMarqueeElementByHTMLMarqueeElement() js.Ref

//go:wasmimport plat/js/web get_HTMLMarqueeElement_Behavior
//go:noescape
func GetHTMLMarqueeElementBehavior(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMarqueeElement_Behavior
//go:noescape
func SetHTMLMarqueeElementBehavior(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMarqueeElement_BgColor
//go:noescape
func GetHTMLMarqueeElementBgColor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMarqueeElement_BgColor
//go:noescape
func SetHTMLMarqueeElementBgColor(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMarqueeElement_Direction
//go:noescape
func GetHTMLMarqueeElementDirection(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMarqueeElement_Direction
//go:noescape
func SetHTMLMarqueeElementDirection(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMarqueeElement_Height
//go:noescape
func GetHTMLMarqueeElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMarqueeElement_Height
//go:noescape
func SetHTMLMarqueeElementHeight(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMarqueeElement_Hspace
//go:noescape
func GetHTMLMarqueeElementHspace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMarqueeElement_Hspace
//go:noescape
func SetHTMLMarqueeElementHspace(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMarqueeElement_Loop
//go:noescape
func GetHTMLMarqueeElementLoop(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMarqueeElement_Loop
//go:noescape
func SetHTMLMarqueeElementLoop(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMarqueeElement_ScrollAmount
//go:noescape
func GetHTMLMarqueeElementScrollAmount(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMarqueeElement_ScrollAmount
//go:noescape
func SetHTMLMarqueeElementScrollAmount(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMarqueeElement_ScrollDelay
//go:noescape
func GetHTMLMarqueeElementScrollDelay(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMarqueeElement_ScrollDelay
//go:noescape
func SetHTMLMarqueeElementScrollDelay(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMarqueeElement_TrueSpeed
//go:noescape
func GetHTMLMarqueeElementTrueSpeed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMarqueeElement_TrueSpeed
//go:noescape
func SetHTMLMarqueeElementTrueSpeed(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMarqueeElement_Vspace
//go:noescape
func GetHTMLMarqueeElementVspace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMarqueeElement_Vspace
//go:noescape
func SetHTMLMarqueeElementVspace(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMarqueeElement_Width
//go:noescape
func GetHTMLMarqueeElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMarqueeElement_Width
//go:noescape
func SetHTMLMarqueeElementWidth(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLMarqueeElement_Start
//go:noescape
func HasHTMLMarqueeElementStart(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMarqueeElement_Start
//go:noescape
func HTMLMarqueeElementStartFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMarqueeElement_Start
//go:noescape
func CallHTMLMarqueeElementStart(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLMarqueeElement_Start
//go:noescape
func TryHTMLMarqueeElementStart(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLMarqueeElement_Stop
//go:noescape
func HasHTMLMarqueeElementStop(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLMarqueeElement_Stop
//go:noescape
func HTMLMarqueeElementStopFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLMarqueeElement_Stop
//go:noescape
func CallHTMLMarqueeElementStop(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLMarqueeElement_Stop
//go:noescape
func TryHTMLMarqueeElementStop(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLMenuElement_HTMLMenuElement
//go:noescape
func NewHTMLMenuElementByHTMLMenuElement() js.Ref

//go:wasmimport plat/js/web get_HTMLMenuElement_Compact
//go:noescape
func GetHTMLMenuElementCompact(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMenuElement_Compact
//go:noescape
func SetHTMLMenuElementCompact(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLMetaElement_HTMLMetaElement
//go:noescape
func NewHTMLMetaElementByHTMLMetaElement() js.Ref

//go:wasmimport plat/js/web get_HTMLMetaElement_Name
//go:noescape
func GetHTMLMetaElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMetaElement_Name
//go:noescape
func SetHTMLMetaElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMetaElement_HttpEquiv
//go:noescape
func GetHTMLMetaElementHttpEquiv(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMetaElement_HttpEquiv
//go:noescape
func SetHTMLMetaElementHttpEquiv(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMetaElement_Content
//go:noescape
func GetHTMLMetaElementContent(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMetaElement_Content
//go:noescape
func SetHTMLMetaElementContent(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMetaElement_Media
//go:noescape
func GetHTMLMetaElementMedia(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMetaElement_Media
//go:noescape
func SetHTMLMetaElementMedia(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMetaElement_Scheme
//go:noescape
func GetHTMLMetaElementScheme(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMetaElement_Scheme
//go:noescape
func SetHTMLMetaElementScheme(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLMeterElement_HTMLMeterElement
//go:noescape
func NewHTMLMeterElementByHTMLMeterElement() js.Ref

//go:wasmimport plat/js/web get_HTMLMeterElement_Value
//go:noescape
func GetHTMLMeterElementValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMeterElement_Value
//go:noescape
func SetHTMLMeterElementValue(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMeterElement_Min
//go:noescape
func GetHTMLMeterElementMin(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMeterElement_Min
//go:noescape
func SetHTMLMeterElementMin(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMeterElement_Max
//go:noescape
func GetHTMLMeterElementMax(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMeterElement_Max
//go:noescape
func SetHTMLMeterElementMax(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMeterElement_Low
//go:noescape
func GetHTMLMeterElementLow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMeterElement_Low
//go:noescape
func SetHTMLMeterElementLow(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMeterElement_High
//go:noescape
func GetHTMLMeterElementHigh(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMeterElement_High
//go:noescape
func SetHTMLMeterElementHigh(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMeterElement_Optimum
//go:noescape
func GetHTMLMeterElementOptimum(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLMeterElement_Optimum
//go:noescape
func SetHTMLMeterElementOptimum(
	this js.Ref,
	val float64,
) js.Ref

//go:wasmimport plat/js/web get_HTMLMeterElement_Labels
//go:noescape
func GetHTMLMeterElementLabels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLModElement_HTMLModElement
//go:noescape
func NewHTMLModElementByHTMLModElement() js.Ref

//go:wasmimport plat/js/web get_HTMLModElement_Cite
//go:noescape
func GetHTMLModElementCite(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLModElement_Cite
//go:noescape
func SetHTMLModElementCite(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLModElement_DateTime
//go:noescape
func GetHTMLModElementDateTime(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLModElement_DateTime
//go:noescape
func SetHTMLModElementDateTime(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLOListElement_HTMLOListElement
//go:noescape
func NewHTMLOListElementByHTMLOListElement() js.Ref

//go:wasmimport plat/js/web get_HTMLOListElement_Reversed
//go:noescape
func GetHTMLOListElementReversed(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOListElement_Reversed
//go:noescape
func SetHTMLOListElementReversed(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOListElement_Start
//go:noescape
func GetHTMLOListElementStart(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOListElement_Start
//go:noescape
func SetHTMLOListElementStart(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOListElement_Type
//go:noescape
func GetHTMLOListElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOListElement_Type
//go:noescape
func SetHTMLOListElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOListElement_Compact
//go:noescape
func GetHTMLOListElementCompact(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOListElement_Compact
//go:noescape
func SetHTMLOListElementCompact(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLObjectElement_HTMLObjectElement
//go:noescape
func NewHTMLObjectElementByHTMLObjectElement() js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Data
//go:noescape
func GetHTMLObjectElementData(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Data
//go:noescape
func SetHTMLObjectElementData(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Type
//go:noescape
func GetHTMLObjectElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Type
//go:noescape
func SetHTMLObjectElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Name
//go:noescape
func GetHTMLObjectElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Name
//go:noescape
func SetHTMLObjectElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Form
//go:noescape
func GetHTMLObjectElementForm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLObjectElement_Width
//go:noescape
func GetHTMLObjectElementWidth(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Width
//go:noescape
func SetHTMLObjectElementWidth(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Height
//go:noescape
func GetHTMLObjectElementHeight(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Height
//go:noescape
func SetHTMLObjectElementHeight(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_ContentDocument
//go:noescape
func GetHTMLObjectElementContentDocument(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLObjectElement_ContentWindow
//go:noescape
func GetHTMLObjectElementContentWindow(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLObjectElement_WillValidate
//go:noescape
func GetHTMLObjectElementWillValidate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLObjectElement_Validity
//go:noescape
func GetHTMLObjectElementValidity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLObjectElement_ValidationMessage
//go:noescape
func GetHTMLObjectElementValidationMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLObjectElement_Align
//go:noescape
func GetHTMLObjectElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Align
//go:noescape
func SetHTMLObjectElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Archive
//go:noescape
func GetHTMLObjectElementArchive(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Archive
//go:noescape
func SetHTMLObjectElementArchive(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Code
//go:noescape
func GetHTMLObjectElementCode(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Code
//go:noescape
func SetHTMLObjectElementCode(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Declare
//go:noescape
func GetHTMLObjectElementDeclare(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Declare
//go:noescape
func SetHTMLObjectElementDeclare(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Hspace
//go:noescape
func GetHTMLObjectElementHspace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Hspace
//go:noescape
func SetHTMLObjectElementHspace(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Standby
//go:noescape
func GetHTMLObjectElementStandby(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Standby
//go:noescape
func SetHTMLObjectElementStandby(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Vspace
//go:noescape
func GetHTMLObjectElementVspace(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Vspace
//go:noescape
func SetHTMLObjectElementVspace(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_CodeBase
//go:noescape
func GetHTMLObjectElementCodeBase(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_CodeBase
//go:noescape
func SetHTMLObjectElementCodeBase(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_CodeType
//go:noescape
func GetHTMLObjectElementCodeType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_CodeType
//go:noescape
func SetHTMLObjectElementCodeType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_UseMap
//go:noescape
func GetHTMLObjectElementUseMap(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_UseMap
//go:noescape
func SetHTMLObjectElementUseMap(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLObjectElement_Border
//go:noescape
func GetHTMLObjectElementBorder(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLObjectElement_Border
//go:noescape
func SetHTMLObjectElementBorder(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web has_HTMLObjectElement_GetSVGDocument
//go:noescape
func HasHTMLObjectElementGetSVGDocument(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLObjectElement_GetSVGDocument
//go:noescape
func HTMLObjectElementGetSVGDocumentFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLObjectElement_GetSVGDocument
//go:noescape
func CallHTMLObjectElementGetSVGDocument(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLObjectElement_GetSVGDocument
//go:noescape
func TryHTMLObjectElementGetSVGDocument(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLObjectElement_CheckValidity
//go:noescape
func HasHTMLObjectElementCheckValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLObjectElement_CheckValidity
//go:noescape
func HTMLObjectElementCheckValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLObjectElement_CheckValidity
//go:noescape
func CallHTMLObjectElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLObjectElement_CheckValidity
//go:noescape
func TryHTMLObjectElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLObjectElement_ReportValidity
//go:noescape
func HasHTMLObjectElementReportValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLObjectElement_ReportValidity
//go:noescape
func HTMLObjectElementReportValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLObjectElement_ReportValidity
//go:noescape
func CallHTMLObjectElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLObjectElement_ReportValidity
//go:noescape
func TryHTMLObjectElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLObjectElement_SetCustomValidity
//go:noescape
func HasHTMLObjectElementSetCustomValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLObjectElement_SetCustomValidity
//go:noescape
func HTMLObjectElementSetCustomValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLObjectElement_SetCustomValidity
//go:noescape
func CallHTMLObjectElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer,
	err js.Ref)

//go:wasmimport plat/js/web try_HTMLObjectElement_SetCustomValidity
//go:noescape
func TryHTMLObjectElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	err js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLOptGroupElement_HTMLOptGroupElement
//go:noescape
func NewHTMLOptGroupElementByHTMLOptGroupElement() js.Ref

//go:wasmimport plat/js/web get_HTMLOptGroupElement_Disabled
//go:noescape
func GetHTMLOptGroupElementDisabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOptGroupElement_Disabled
//go:noescape
func SetHTMLOptGroupElementDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOptGroupElement_Label
//go:noescape
func GetHTMLOptGroupElementLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOptGroupElement_Label
//go:noescape
func SetHTMLOptGroupElementLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLOptionElement_HTMLOptionElement
//go:noescape
func NewHTMLOptionElementByHTMLOptionElement() js.Ref

//go:wasmimport plat/js/web get_HTMLOptionElement_Disabled
//go:noescape
func GetHTMLOptionElementDisabled(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOptionElement_Disabled
//go:noescape
func SetHTMLOptionElementDisabled(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOptionElement_Form
//go:noescape
func GetHTMLOptionElementForm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLOptionElement_Label
//go:noescape
func GetHTMLOptionElementLabel(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOptionElement_Label
//go:noescape
func SetHTMLOptionElementLabel(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOptionElement_DefaultSelected
//go:noescape
func GetHTMLOptionElementDefaultSelected(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOptionElement_DefaultSelected
//go:noescape
func SetHTMLOptionElementDefaultSelected(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOptionElement_Selected
//go:noescape
func GetHTMLOptionElementSelected(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOptionElement_Selected
//go:noescape
func SetHTMLOptionElementSelected(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOptionElement_Value
//go:noescape
func GetHTMLOptionElementValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOptionElement_Value
//go:noescape
func SetHTMLOptionElementValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOptionElement_Text
//go:noescape
func GetHTMLOptionElementText(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOptionElement_Text
//go:noescape
func SetHTMLOptionElementText(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOptionElement_Index
//go:noescape
func GetHTMLOptionElementIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLOptionsCollection_Length
//go:noescape
func GetHTMLOptionsCollectionLength(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOptionsCollection_Length
//go:noescape
func SetHTMLOptionsCollectionLength(
	this js.Ref,
	val uint32,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOptionsCollection_SelectedIndex
//go:noescape
func GetHTMLOptionsCollectionSelectedIndex(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOptionsCollection_SelectedIndex
//go:noescape
func SetHTMLOptionsCollectionSelectedIndex(
	this js.Ref,
	val int32,
) js.Ref

//go:wasmimport plat/js/web has_HTMLOptionsCollection_Set
//go:noescape
func HasHTMLOptionsCollectionSet(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLOptionsCollection_Set
//go:noescape
func HTMLOptionsCollectionSetFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLOptionsCollection_Set
//go:noescape
func CallHTMLOptionsCollectionSet(
	this js.Ref, retPtr unsafe.Pointer,
	index uint32,
	option js.Ref)

//go:wasmimport plat/js/web try_HTMLOptionsCollection_Set
//go:noescape
func TryHTMLOptionsCollectionSet(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index uint32,
	option js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLOptionsCollection_Add
//go:noescape
func HasHTMLOptionsCollectionAdd(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLOptionsCollection_Add
//go:noescape
func HTMLOptionsCollectionAddFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLOptionsCollection_Add
//go:noescape
func CallHTMLOptionsCollectionAdd(
	this js.Ref, retPtr unsafe.Pointer,
	element js.Ref,
	before js.Ref)

//go:wasmimport plat/js/web try_HTMLOptionsCollection_Add
//go:noescape
func TryHTMLOptionsCollectionAdd(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref,
	before js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLOptionsCollection_Add1
//go:noescape
func HasHTMLOptionsCollectionAdd1(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLOptionsCollection_Add1
//go:noescape
func HTMLOptionsCollectionAdd1Func(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLOptionsCollection_Add1
//go:noescape
func CallHTMLOptionsCollectionAdd1(
	this js.Ref, retPtr unsafe.Pointer,
	element js.Ref)

//go:wasmimport plat/js/web try_HTMLOptionsCollection_Add1
//go:noescape
func TryHTMLOptionsCollectionAdd1(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	element js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLOptionsCollection_Remove
//go:noescape
func HasHTMLOptionsCollectionRemove(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLOptionsCollection_Remove
//go:noescape
func HTMLOptionsCollectionRemoveFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLOptionsCollection_Remove
//go:noescape
func CallHTMLOptionsCollectionRemove(
	this js.Ref, retPtr unsafe.Pointer,
	index int32)

//go:wasmimport plat/js/web try_HTMLOptionsCollection_Remove
//go:noescape
func TryHTMLOptionsCollectionRemove(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	index int32) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLOutputElement_HTMLOutputElement
//go:noescape
func NewHTMLOutputElementByHTMLOutputElement() js.Ref

//go:wasmimport plat/js/web get_HTMLOutputElement_HtmlFor
//go:noescape
func GetHTMLOutputElementHtmlFor(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLOutputElement_Form
//go:noescape
func GetHTMLOutputElementForm(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLOutputElement_Name
//go:noescape
func GetHTMLOutputElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOutputElement_Name
//go:noescape
func SetHTMLOutputElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOutputElement_Type
//go:noescape
func GetHTMLOutputElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLOutputElement_DefaultValue
//go:noescape
func GetHTMLOutputElementDefaultValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOutputElement_DefaultValue
//go:noescape
func SetHTMLOutputElementDefaultValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOutputElement_Value
//go:noescape
func GetHTMLOutputElementValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLOutputElement_Value
//go:noescape
func SetHTMLOutputElementValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLOutputElement_WillValidate
//go:noescape
func GetHTMLOutputElementWillValidate(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLOutputElement_Validity
//go:noescape
func GetHTMLOutputElementValidity(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLOutputElement_ValidationMessage
//go:noescape
func GetHTMLOutputElementValidationMessage(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web get_HTMLOutputElement_Labels
//go:noescape
func GetHTMLOutputElementLabels(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLOutputElement_CheckValidity
//go:noescape
func HasHTMLOutputElementCheckValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLOutputElement_CheckValidity
//go:noescape
func HTMLOutputElementCheckValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLOutputElement_CheckValidity
//go:noescape
func CallHTMLOutputElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLOutputElement_CheckValidity
//go:noescape
func TryHTMLOutputElementCheckValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLOutputElement_ReportValidity
//go:noescape
func HasHTMLOutputElementReportValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLOutputElement_ReportValidity
//go:noescape
func HTMLOutputElementReportValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLOutputElement_ReportValidity
//go:noescape
func CallHTMLOutputElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer)

//go:wasmimport plat/js/web try_HTMLOutputElement_ReportValidity
//go:noescape
func TryHTMLOutputElementReportValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web has_HTMLOutputElement_SetCustomValidity
//go:noescape
func HasHTMLOutputElementSetCustomValidity(this js.Ref) js.Ref

//go:wasmimport plat/js/web func_HTMLOutputElement_SetCustomValidity
//go:noescape
func HTMLOutputElementSetCustomValidityFunc(this js.Ref) js.Ref

//go:wasmimport plat/js/web call_HTMLOutputElement_SetCustomValidity
//go:noescape
func CallHTMLOutputElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer,
	err js.Ref)

//go:wasmimport plat/js/web try_HTMLOutputElement_SetCustomValidity
//go:noescape
func TryHTMLOutputElementSetCustomValidity(
	this js.Ref, retPtr unsafe.Pointer, errPtr unsafe.Pointer,
	err js.Ref) (ok js.Ref)

//go:wasmimport plat/js/web new_HTMLParagraphElement_HTMLParagraphElement
//go:noescape
func NewHTMLParagraphElementByHTMLParagraphElement() js.Ref

//go:wasmimport plat/js/web get_HTMLParagraphElement_Align
//go:noescape
func GetHTMLParagraphElementAlign(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLParagraphElement_Align
//go:noescape
func SetHTMLParagraphElementAlign(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLParamElement_HTMLParamElement
//go:noescape
func NewHTMLParamElementByHTMLParamElement() js.Ref

//go:wasmimport plat/js/web get_HTMLParamElement_Name
//go:noescape
func GetHTMLParamElementName(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLParamElement_Name
//go:noescape
func SetHTMLParamElementName(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLParamElement_Value
//go:noescape
func GetHTMLParamElementValue(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLParamElement_Value
//go:noescape
func SetHTMLParamElementValue(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLParamElement_Type
//go:noescape
func GetHTMLParamElementType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLParamElement_Type
//go:noescape
func SetHTMLParamElementType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web get_HTMLParamElement_ValueType
//go:noescape
func GetHTMLParamElementValueType(
	this js.Ref, retPtr unsafe.Pointer) (ok js.Ref)

//go:wasmimport plat/js/web set_HTMLParamElement_ValueType
//go:noescape
func SetHTMLParamElementValueType(
	this js.Ref,
	val js.Ref,
) js.Ref

//go:wasmimport plat/js/web new_HTMLPictureElement_HTMLPictureElement
//go:noescape
func NewHTMLPictureElementByHTMLPictureElement() js.Ref

//go:wasmimport plat/js/web store_PortalActivateOptions
//go:noescape
func PortalActivateOptionsJSStore(
	ptr unsafe.Pointer, ref js.Ref)

//go:wasmimport plat/js/web load_PortalActivateOptions
//go:noescape
func PortalActivateOptionsJSLoad(
	ptr unsafe.Pointer, create, ref js.Ref) js.Ref

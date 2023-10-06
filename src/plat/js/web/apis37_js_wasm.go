// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/core/abi"
	"github.com/primecitizens/pcz/std/core/assert"
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

func NewHTMLAudioElement() (ret HTMLAudioElement) {
	ret.ref = bindings.NewHTMLAudioElementByHTMLAudioElement()
	return
}

type HTMLAudioElement struct {
	HTMLMediaElement
}

func (this HTMLAudioElement) Once() HTMLAudioElement {
	this.Ref().Once()
	return this
}

func (this HTMLAudioElement) Ref() js.Ref {
	return this.HTMLMediaElement.Ref()
}

func (this HTMLAudioElement) FromRef(ref js.Ref) HTMLAudioElement {
	this.HTMLMediaElement = this.HTMLMediaElement.FromRef(ref)
	return this
}

func (this HTMLAudioElement) Free() {
	this.Ref().Free()
}

func NewHTMLBRElement() (ret HTMLBRElement) {
	ret.ref = bindings.NewHTMLBRElementByHTMLBRElement()
	return
}

type HTMLBRElement struct {
	HTMLElement
}

func (this HTMLBRElement) Once() HTMLBRElement {
	this.Ref().Once()
	return this
}

func (this HTMLBRElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLBRElement) FromRef(ref js.Ref) HTMLBRElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLBRElement) Free() {
	this.Ref().Free()
}

// Clear returns the value of property "HTMLBRElement.clear".
//
// It returns ok=false if there is no such property.
func (this HTMLBRElement) Clear() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBRElementClear(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetClear sets the value of property "HTMLBRElement.clear" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBRElement) SetClear(val js.String) bool {
	return js.True == bindings.SetHTMLBRElementClear(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLBaseElement() (ret HTMLBaseElement) {
	ret.ref = bindings.NewHTMLBaseElementByHTMLBaseElement()
	return
}

type HTMLBaseElement struct {
	HTMLElement
}

func (this HTMLBaseElement) Once() HTMLBaseElement {
	this.Ref().Once()
	return this
}

func (this HTMLBaseElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLBaseElement) FromRef(ref js.Ref) HTMLBaseElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLBaseElement) Free() {
	this.Ref().Free()
}

// Href returns the value of property "HTMLBaseElement.href".
//
// It returns ok=false if there is no such property.
func (this HTMLBaseElement) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBaseElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHref sets the value of property "HTMLBaseElement.href" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBaseElement) SetHref(val js.String) bool {
	return js.True == bindings.SetHTMLBaseElementHref(
		this.Ref(),
		val.Ref(),
	)
}

// Target returns the value of property "HTMLBaseElement.target".
//
// It returns ok=false if there is no such property.
func (this HTMLBaseElement) Target() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBaseElementTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTarget sets the value of property "HTMLBaseElement.target" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBaseElement) SetTarget(val js.String) bool {
	return js.True == bindings.SetHTMLBaseElementTarget(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLBodyElement() (ret HTMLBodyElement) {
	ret.ref = bindings.NewHTMLBodyElementByHTMLBodyElement()
	return
}

type HTMLBodyElement struct {
	HTMLElement
}

func (this HTMLBodyElement) Once() HTMLBodyElement {
	this.Ref().Once()
	return this
}

func (this HTMLBodyElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLBodyElement) FromRef(ref js.Ref) HTMLBodyElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLBodyElement) Free() {
	this.Ref().Free()
}

// Text returns the value of property "HTMLBodyElement.text".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "HTMLBodyElement.text" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetText(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementText(
		this.Ref(),
		val.Ref(),
	)
}

// Link returns the value of property "HTMLBodyElement.link".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) Link() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementLink(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLink sets the value of property "HTMLBodyElement.link" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetLink(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementLink(
		this.Ref(),
		val.Ref(),
	)
}

// VLink returns the value of property "HTMLBodyElement.vLink".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) VLink() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementVLink(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVLink sets the value of property "HTMLBodyElement.vLink" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetVLink(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementVLink(
		this.Ref(),
		val.Ref(),
	)
}

// ALink returns the value of property "HTMLBodyElement.aLink".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) ALink() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementALink(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetALink sets the value of property "HTMLBodyElement.aLink" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetALink(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementALink(
		this.Ref(),
		val.Ref(),
	)
}

// BgColor returns the value of property "HTMLBodyElement.bgColor".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementBgColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "HTMLBodyElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementBgColor(
		this.Ref(),
		val.Ref(),
	)
}

// Background returns the value of property "HTMLBodyElement.background".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) Background() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementBackground(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBackground sets the value of property "HTMLBodyElement.background" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetBackground(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementBackground(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLButtonElement() (ret HTMLButtonElement) {
	ret.ref = bindings.NewHTMLButtonElementByHTMLButtonElement()
	return
}

type HTMLButtonElement struct {
	HTMLElement
}

func (this HTMLButtonElement) Once() HTMLButtonElement {
	this.Ref().Once()
	return this
}

func (this HTMLButtonElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLButtonElement) FromRef(ref js.Ref) HTMLButtonElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLButtonElement) Free() {
	this.Ref().Free()
}

// Disabled returns the value of property "HTMLButtonElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementDisabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLButtonElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLButtonElementDisabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLButtonElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementForm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// FormAction returns the value of property "HTMLButtonElement.formAction".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) FormAction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementFormAction(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFormAction sets the value of property "HTMLButtonElement.formAction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetFormAction(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementFormAction(
		this.Ref(),
		val.Ref(),
	)
}

// FormEnctype returns the value of property "HTMLButtonElement.formEnctype".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) FormEnctype() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementFormEnctype(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFormEnctype sets the value of property "HTMLButtonElement.formEnctype" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetFormEnctype(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementFormEnctype(
		this.Ref(),
		val.Ref(),
	)
}

// FormMethod returns the value of property "HTMLButtonElement.formMethod".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) FormMethod() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementFormMethod(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFormMethod sets the value of property "HTMLButtonElement.formMethod" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetFormMethod(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementFormMethod(
		this.Ref(),
		val.Ref(),
	)
}

// FormNoValidate returns the value of property "HTMLButtonElement.formNoValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) FormNoValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementFormNoValidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFormNoValidate sets the value of property "HTMLButtonElement.formNoValidate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetFormNoValidate(val bool) bool {
	return js.True == bindings.SetHTMLButtonElementFormNoValidate(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// FormTarget returns the value of property "HTMLButtonElement.formTarget".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) FormTarget() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementFormTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFormTarget sets the value of property "HTMLButtonElement.formTarget" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetFormTarget(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementFormTarget(
		this.Ref(),
		val.Ref(),
	)
}

// Name returns the value of property "HTMLButtonElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLButtonElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLButtonElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLButtonElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementType(
		this.Ref(),
		val.Ref(),
	)
}

// Value returns the value of property "HTMLButtonElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLButtonElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementValue(
		this.Ref(),
		val.Ref(),
	)
}

// WillValidate returns the value of property "HTMLButtonElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementWillValidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLButtonElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementValidity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLButtonElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementValidationMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLButtonElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementLabels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PopoverTargetElement returns the value of property "HTMLButtonElement.popoverTargetElement".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) PopoverTargetElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementPopoverTargetElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPopoverTargetElement sets the value of property "HTMLButtonElement.popoverTargetElement" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetPopoverTargetElement(val Element) bool {
	return js.True == bindings.SetHTMLButtonElementPopoverTargetElement(
		this.Ref(),
		val.Ref(),
	)
}

// PopoverTargetAction returns the value of property "HTMLButtonElement.popoverTargetAction".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) PopoverTargetAction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementPopoverTargetAction(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPopoverTargetAction sets the value of property "HTMLButtonElement.popoverTargetAction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetPopoverTargetAction(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementPopoverTargetAction(
		this.Ref(),
		val.Ref(),
	)
}

// HasCheckValidity returns true if the method "HTMLButtonElement.checkValidity" exists.
func (this HTMLButtonElement) HasCheckValidity() bool {
	return js.True == bindings.HasHTMLButtonElementCheckValidity(
		this.Ref(),
	)
}

// CheckValidityFunc returns the method "HTMLButtonElement.checkValidity".
func (this HTMLButtonElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLButtonElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLButtonElement.checkValidity".
func (this HTMLButtonElement) CheckValidity() (ret bool) {
	bindings.CallHTMLButtonElementCheckValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLButtonElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLButtonElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLButtonElementCheckValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReportValidity returns true if the method "HTMLButtonElement.reportValidity" exists.
func (this HTMLButtonElement) HasReportValidity() bool {
	return js.True == bindings.HasHTMLButtonElementReportValidity(
		this.Ref(),
	)
}

// ReportValidityFunc returns the method "HTMLButtonElement.reportValidity".
func (this HTMLButtonElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLButtonElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLButtonElement.reportValidity".
func (this HTMLButtonElement) ReportValidity() (ret bool) {
	bindings.CallHTMLButtonElementReportValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLButtonElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLButtonElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLButtonElementReportValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetCustomValidity returns true if the method "HTMLButtonElement.setCustomValidity" exists.
func (this HTMLButtonElement) HasSetCustomValidity() bool {
	return js.True == bindings.HasHTMLButtonElementSetCustomValidity(
		this.Ref(),
	)
}

// SetCustomValidityFunc returns the method "HTMLButtonElement.setCustomValidity".
func (this HTMLButtonElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLButtonElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLButtonElement.setCustomValidity".
func (this HTMLButtonElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLButtonElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLButtonElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLButtonElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLButtonElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

func NewHTMLDListElement() (ret HTMLDListElement) {
	ret.ref = bindings.NewHTMLDListElementByHTMLDListElement()
	return
}

type HTMLDListElement struct {
	HTMLElement
}

func (this HTMLDListElement) Once() HTMLDListElement {
	this.Ref().Once()
	return this
}

func (this HTMLDListElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLDListElement) FromRef(ref js.Ref) HTMLDListElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLDListElement) Free() {
	this.Ref().Free()
}

// Compact returns the value of property "HTMLDListElement.compact".
//
// It returns ok=false if there is no such property.
func (this HTMLDListElement) Compact() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLDListElementCompact(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCompact sets the value of property "HTMLDListElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDListElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLDListElementCompact(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

func NewHTMLDataElement() (ret HTMLDataElement) {
	ret.ref = bindings.NewHTMLDataElementByHTMLDataElement()
	return
}

type HTMLDataElement struct {
	HTMLElement
}

func (this HTMLDataElement) Once() HTMLDataElement {
	this.Ref().Once()
	return this
}

func (this HTMLDataElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLDataElement) FromRef(ref js.Ref) HTMLDataElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLDataElement) Free() {
	this.Ref().Free()
}

// Value returns the value of property "HTMLDataElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLDataElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLDataElementValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLDataElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDataElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLDataElementValue(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLDataListElement() (ret HTMLDataListElement) {
	ret.ref = bindings.NewHTMLDataListElementByHTMLDataListElement()
	return
}

type HTMLDataListElement struct {
	HTMLElement
}

func (this HTMLDataListElement) Once() HTMLDataListElement {
	this.Ref().Once()
	return this
}

func (this HTMLDataListElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLDataListElement) FromRef(ref js.Ref) HTMLDataListElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLDataListElement) Free() {
	this.Ref().Free()
}

// Options returns the value of property "HTMLDataListElement.options".
//
// It returns ok=false if there is no such property.
func (this HTMLDataListElement) Options() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLDataListElementOptions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewHTMLDetailsElement() (ret HTMLDetailsElement) {
	ret.ref = bindings.NewHTMLDetailsElementByHTMLDetailsElement()
	return
}

type HTMLDetailsElement struct {
	HTMLElement
}

func (this HTMLDetailsElement) Once() HTMLDetailsElement {
	this.Ref().Once()
	return this
}

func (this HTMLDetailsElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLDetailsElement) FromRef(ref js.Ref) HTMLDetailsElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLDetailsElement) Free() {
	this.Ref().Free()
}

// Open returns the value of property "HTMLDetailsElement.open".
//
// It returns ok=false if there is no such property.
func (this HTMLDetailsElement) Open() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLDetailsElementOpen(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetOpen sets the value of property "HTMLDetailsElement.open" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDetailsElement) SetOpen(val bool) bool {
	return js.True == bindings.SetHTMLDetailsElementOpen(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

func NewHTMLDialogElement() (ret HTMLDialogElement) {
	ret.ref = bindings.NewHTMLDialogElementByHTMLDialogElement()
	return
}

type HTMLDialogElement struct {
	HTMLElement
}

func (this HTMLDialogElement) Once() HTMLDialogElement {
	this.Ref().Once()
	return this
}

func (this HTMLDialogElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLDialogElement) FromRef(ref js.Ref) HTMLDialogElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLDialogElement) Free() {
	this.Ref().Free()
}

// Open returns the value of property "HTMLDialogElement.open".
//
// It returns ok=false if there is no such property.
func (this HTMLDialogElement) Open() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLDialogElementOpen(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetOpen sets the value of property "HTMLDialogElement.open" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDialogElement) SetOpen(val bool) bool {
	return js.True == bindings.SetHTMLDialogElementOpen(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// ReturnValue returns the value of property "HTMLDialogElement.returnValue".
//
// It returns ok=false if there is no such property.
func (this HTMLDialogElement) ReturnValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLDialogElementReturnValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReturnValue sets the value of property "HTMLDialogElement.returnValue" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDialogElement) SetReturnValue(val js.String) bool {
	return js.True == bindings.SetHTMLDialogElementReturnValue(
		this.Ref(),
		val.Ref(),
	)
}

// HasShow returns true if the method "HTMLDialogElement.show" exists.
func (this HTMLDialogElement) HasShow() bool {
	return js.True == bindings.HasHTMLDialogElementShow(
		this.Ref(),
	)
}

// ShowFunc returns the method "HTMLDialogElement.show".
func (this HTMLDialogElement) ShowFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLDialogElementShowFunc(
			this.Ref(),
		),
	)
}

// Show calls the method "HTMLDialogElement.show".
func (this HTMLDialogElement) Show() (ret js.Void) {
	bindings.CallHTMLDialogElementShow(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryShow calls the method "HTMLDialogElement.show"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLDialogElement) TryShow() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLDialogElementShow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasShowModal returns true if the method "HTMLDialogElement.showModal" exists.
func (this HTMLDialogElement) HasShowModal() bool {
	return js.True == bindings.HasHTMLDialogElementShowModal(
		this.Ref(),
	)
}

// ShowModalFunc returns the method "HTMLDialogElement.showModal".
func (this HTMLDialogElement) ShowModalFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLDialogElementShowModalFunc(
			this.Ref(),
		),
	)
}

// ShowModal calls the method "HTMLDialogElement.showModal".
func (this HTMLDialogElement) ShowModal() (ret js.Void) {
	bindings.CallHTMLDialogElementShowModal(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryShowModal calls the method "HTMLDialogElement.showModal"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLDialogElement) TryShowModal() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLDialogElementShowModal(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasClose returns true if the method "HTMLDialogElement.close" exists.
func (this HTMLDialogElement) HasClose() bool {
	return js.True == bindings.HasHTMLDialogElementClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "HTMLDialogElement.close".
func (this HTMLDialogElement) CloseFunc() (fn js.Func[func(returnValue js.String)]) {
	return fn.FromRef(
		bindings.HTMLDialogElementCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "HTMLDialogElement.close".
func (this HTMLDialogElement) Close(returnValue js.String) (ret js.Void) {
	bindings.CallHTMLDialogElementClose(
		this.Ref(), js.Pointer(&ret),
		returnValue.Ref(),
	)

	return
}

// TryClose calls the method "HTMLDialogElement.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLDialogElement) TryClose(returnValue js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLDialogElementClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		returnValue.Ref(),
	)

	return
}

// HasClose1 returns true if the method "HTMLDialogElement.close" exists.
func (this HTMLDialogElement) HasClose1() bool {
	return js.True == bindings.HasHTMLDialogElementClose1(
		this.Ref(),
	)
}

// Close1Func returns the method "HTMLDialogElement.close".
func (this HTMLDialogElement) Close1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLDialogElementClose1Func(
			this.Ref(),
		),
	)
}

// Close1 calls the method "HTMLDialogElement.close".
func (this HTMLDialogElement) Close1() (ret js.Void) {
	bindings.CallHTMLDialogElementClose1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose1 calls the method "HTMLDialogElement.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLDialogElement) TryClose1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLDialogElementClose1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewHTMLDirectoryElement() (ret HTMLDirectoryElement) {
	ret.ref = bindings.NewHTMLDirectoryElementByHTMLDirectoryElement()
	return
}

type HTMLDirectoryElement struct {
	HTMLElement
}

func (this HTMLDirectoryElement) Once() HTMLDirectoryElement {
	this.Ref().Once()
	return this
}

func (this HTMLDirectoryElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLDirectoryElement) FromRef(ref js.Ref) HTMLDirectoryElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLDirectoryElement) Free() {
	this.Ref().Free()
}

// Compact returns the value of property "HTMLDirectoryElement.compact".
//
// It returns ok=false if there is no such property.
func (this HTMLDirectoryElement) Compact() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLDirectoryElementCompact(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCompact sets the value of property "HTMLDirectoryElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDirectoryElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLDirectoryElementCompact(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

func NewHTMLDivElement() (ret HTMLDivElement) {
	ret.ref = bindings.NewHTMLDivElementByHTMLDivElement()
	return
}

type HTMLDivElement struct {
	HTMLElement
}

func (this HTMLDivElement) Once() HTMLDivElement {
	this.Ref().Once()
	return this
}

func (this HTMLDivElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLDivElement) FromRef(ref js.Ref) HTMLDivElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLDivElement) Free() {
	this.Ref().Free()
}

// Align returns the value of property "HTMLDivElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLDivElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLDivElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLDivElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDivElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLDivElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLEmbedElement() (ret HTMLEmbedElement) {
	ret.ref = bindings.NewHTMLEmbedElementByHTMLEmbedElement()
	return
}

type HTMLEmbedElement struct {
	HTMLElement
}

func (this HTMLEmbedElement) Once() HTMLEmbedElement {
	this.Ref().Once()
	return this
}

func (this HTMLEmbedElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLEmbedElement) FromRef(ref js.Ref) HTMLEmbedElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLEmbedElement) Free() {
	this.Ref().Free()
}

// Src returns the value of property "HTMLEmbedElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLEmbedElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementSrc(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLEmbedElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLEmbedElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementType(
		this.Ref(),
		val.Ref(),
	)
}

// Width returns the value of property "HTMLEmbedElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLEmbedElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

// Height returns the value of property "HTMLEmbedElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLEmbedElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementHeight(
		this.Ref(),
		val.Ref(),
	)
}

// Align returns the value of property "HTMLEmbedElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLEmbedElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

// Name returns the value of property "HTMLEmbedElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLEmbedElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementName(
		this.Ref(),
		val.Ref(),
	)
}

// HasGetSVGDocument returns true if the method "HTMLEmbedElement.getSVGDocument" exists.
func (this HTMLEmbedElement) HasGetSVGDocument() bool {
	return js.True == bindings.HasHTMLEmbedElementGetSVGDocument(
		this.Ref(),
	)
}

// GetSVGDocumentFunc returns the method "HTMLEmbedElement.getSVGDocument".
func (this HTMLEmbedElement) GetSVGDocumentFunc() (fn js.Func[func() Document]) {
	return fn.FromRef(
		bindings.HTMLEmbedElementGetSVGDocumentFunc(
			this.Ref(),
		),
	)
}

// GetSVGDocument calls the method "HTMLEmbedElement.getSVGDocument".
func (this HTMLEmbedElement) GetSVGDocument() (ret Document) {
	bindings.CallHTMLEmbedElementGetSVGDocument(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSVGDocument calls the method "HTMLEmbedElement.getSVGDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLEmbedElement) TryGetSVGDocument() (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLEmbedElementGetSVGDocument(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewHTMLFencedFrameElement() (ret HTMLFencedFrameElement) {
	ret.ref = bindings.NewHTMLFencedFrameElementByHTMLFencedFrameElement()
	return
}

type HTMLFencedFrameElement struct {
	HTMLElement
}

func (this HTMLFencedFrameElement) Once() HTMLFencedFrameElement {
	this.Ref().Once()
	return this
}

func (this HTMLFencedFrameElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLFencedFrameElement) FromRef(ref js.Ref) HTMLFencedFrameElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLFencedFrameElement) Free() {
	this.Ref().Free()
}

// Config returns the value of property "HTMLFencedFrameElement.config".
//
// It returns ok=false if there is no such property.
func (this HTMLFencedFrameElement) Config() (ret FencedFrameConfig, ok bool) {
	ok = js.True == bindings.GetHTMLFencedFrameElementConfig(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetConfig sets the value of property "HTMLFencedFrameElement.config" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFencedFrameElement) SetConfig(val FencedFrameConfig) bool {
	return js.True == bindings.SetHTMLFencedFrameElementConfig(
		this.Ref(),
		val.Ref(),
	)
}

// Width returns the value of property "HTMLFencedFrameElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLFencedFrameElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFencedFrameElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLFencedFrameElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFencedFrameElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLFencedFrameElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

// Height returns the value of property "HTMLFencedFrameElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLFencedFrameElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFencedFrameElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLFencedFrameElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFencedFrameElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLFencedFrameElementHeight(
		this.Ref(),
		val.Ref(),
	)
}

// Allow returns the value of property "HTMLFencedFrameElement.allow".
//
// It returns ok=false if there is no such property.
func (this HTMLFencedFrameElement) Allow() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFencedFrameElementAllow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAllow sets the value of property "HTMLFencedFrameElement.allow" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFencedFrameElement) SetAllow(val js.String) bool {
	return js.True == bindings.SetHTMLFencedFrameElementAllow(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLFieldSetElement() (ret HTMLFieldSetElement) {
	ret.ref = bindings.NewHTMLFieldSetElementByHTMLFieldSetElement()
	return
}

type HTMLFieldSetElement struct {
	HTMLElement
}

func (this HTMLFieldSetElement) Once() HTMLFieldSetElement {
	this.Ref().Once()
	return this
}

func (this HTMLFieldSetElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLFieldSetElement) FromRef(ref js.Ref) HTMLFieldSetElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLFieldSetElement) Free() {
	this.Ref().Free()
}

// Disabled returns the value of property "HTMLFieldSetElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementDisabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLFieldSetElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFieldSetElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLFieldSetElementDisabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLFieldSetElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementForm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "HTMLFieldSetElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLFieldSetElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFieldSetElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLFieldSetElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLFieldSetElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Elements returns the value of property "HTMLFieldSetElement.elements".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Elements() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementElements(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WillValidate returns the value of property "HTMLFieldSetElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementWillValidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLFieldSetElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementValidity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLFieldSetElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementValidationMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasCheckValidity returns true if the method "HTMLFieldSetElement.checkValidity" exists.
func (this HTMLFieldSetElement) HasCheckValidity() bool {
	return js.True == bindings.HasHTMLFieldSetElementCheckValidity(
		this.Ref(),
	)
}

// CheckValidityFunc returns the method "HTMLFieldSetElement.checkValidity".
func (this HTMLFieldSetElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLFieldSetElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLFieldSetElement.checkValidity".
func (this HTMLFieldSetElement) CheckValidity() (ret bool) {
	bindings.CallHTMLFieldSetElementCheckValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLFieldSetElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFieldSetElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFieldSetElementCheckValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReportValidity returns true if the method "HTMLFieldSetElement.reportValidity" exists.
func (this HTMLFieldSetElement) HasReportValidity() bool {
	return js.True == bindings.HasHTMLFieldSetElementReportValidity(
		this.Ref(),
	)
}

// ReportValidityFunc returns the method "HTMLFieldSetElement.reportValidity".
func (this HTMLFieldSetElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLFieldSetElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLFieldSetElement.reportValidity".
func (this HTMLFieldSetElement) ReportValidity() (ret bool) {
	bindings.CallHTMLFieldSetElementReportValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLFieldSetElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFieldSetElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFieldSetElementReportValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetCustomValidity returns true if the method "HTMLFieldSetElement.setCustomValidity" exists.
func (this HTMLFieldSetElement) HasSetCustomValidity() bool {
	return js.True == bindings.HasHTMLFieldSetElementSetCustomValidity(
		this.Ref(),
	)
}

// SetCustomValidityFunc returns the method "HTMLFieldSetElement.setCustomValidity".
func (this HTMLFieldSetElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLFieldSetElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLFieldSetElement.setCustomValidity".
func (this HTMLFieldSetElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLFieldSetElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLFieldSetElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFieldSetElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFieldSetElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

func NewHTMLFontElement() (ret HTMLFontElement) {
	ret.ref = bindings.NewHTMLFontElementByHTMLFontElement()
	return
}

type HTMLFontElement struct {
	HTMLElement
}

func (this HTMLFontElement) Once() HTMLFontElement {
	this.Ref().Once()
	return this
}

func (this HTMLFontElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLFontElement) FromRef(ref js.Ref) HTMLFontElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLFontElement) Free() {
	this.Ref().Free()
}

// Color returns the value of property "HTMLFontElement.color".
//
// It returns ok=false if there is no such property.
func (this HTMLFontElement) Color() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFontElementColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetColor sets the value of property "HTMLFontElement.color" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFontElement) SetColor(val js.String) bool {
	return js.True == bindings.SetHTMLFontElementColor(
		this.Ref(),
		val.Ref(),
	)
}

// Face returns the value of property "HTMLFontElement.face".
//
// It returns ok=false if there is no such property.
func (this HTMLFontElement) Face() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFontElementFace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFace sets the value of property "HTMLFontElement.face" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFontElement) SetFace(val js.String) bool {
	return js.True == bindings.SetHTMLFontElementFace(
		this.Ref(),
		val.Ref(),
	)
}

// Size returns the value of property "HTMLFontElement.size".
//
// It returns ok=false if there is no such property.
func (this HTMLFontElement) Size() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFontElementSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSize sets the value of property "HTMLFontElement.size" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFontElement) SetSize(val js.String) bool {
	return js.True == bindings.SetHTMLFontElementSize(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLFrameElement() (ret HTMLFrameElement) {
	ret.ref = bindings.NewHTMLFrameElementByHTMLFrameElement()
	return
}

type HTMLFrameElement struct {
	HTMLElement
}

func (this HTMLFrameElement) Once() HTMLFrameElement {
	this.Ref().Once()
	return this
}

func (this HTMLFrameElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLFrameElement) FromRef(ref js.Ref) HTMLFrameElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLFrameElement) Free() {
	this.Ref().Free()
}

// Name returns the value of property "HTMLFrameElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLFrameElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Scrolling returns the value of property "HTMLFrameElement.scrolling".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) Scrolling() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementScrolling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetScrolling sets the value of property "HTMLFrameElement.scrolling" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetScrolling(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementScrolling(
		this.Ref(),
		val.Ref(),
	)
}

// Src returns the value of property "HTMLFrameElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLFrameElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementSrc(
		this.Ref(),
		val.Ref(),
	)
}

// FrameBorder returns the value of property "HTMLFrameElement.frameBorder".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) FrameBorder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementFrameBorder(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFrameBorder sets the value of property "HTMLFrameElement.frameBorder" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetFrameBorder(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementFrameBorder(
		this.Ref(),
		val.Ref(),
	)
}

// LongDesc returns the value of property "HTMLFrameElement.longDesc".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) LongDesc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementLongDesc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLongDesc sets the value of property "HTMLFrameElement.longDesc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetLongDesc(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementLongDesc(
		this.Ref(),
		val.Ref(),
	)
}

// NoResize returns the value of property "HTMLFrameElement.noResize".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) NoResize() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementNoResize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetNoResize sets the value of property "HTMLFrameElement.noResize" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetNoResize(val bool) bool {
	return js.True == bindings.SetHTMLFrameElementNoResize(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// ContentDocument returns the value of property "HTMLFrameElement.contentDocument".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) ContentDocument() (ret Document, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementContentDocument(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContentWindow returns the value of property "HTMLFrameElement.contentWindow".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) ContentWindow() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementContentWindow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MarginHeight returns the value of property "HTMLFrameElement.marginHeight".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) MarginHeight() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementMarginHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMarginHeight sets the value of property "HTMLFrameElement.marginHeight" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetMarginHeight(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementMarginHeight(
		this.Ref(),
		val.Ref(),
	)
}

// MarginWidth returns the value of property "HTMLFrameElement.marginWidth".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) MarginWidth() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementMarginWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMarginWidth sets the value of property "HTMLFrameElement.marginWidth" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetMarginWidth(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementMarginWidth(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLFrameSetElement() (ret HTMLFrameSetElement) {
	ret.ref = bindings.NewHTMLFrameSetElementByHTMLFrameSetElement()
	return
}

type HTMLFrameSetElement struct {
	HTMLElement
}

func (this HTMLFrameSetElement) Once() HTMLFrameSetElement {
	this.Ref().Once()
	return this
}

func (this HTMLFrameSetElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLFrameSetElement) FromRef(ref js.Ref) HTMLFrameSetElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLFrameSetElement) Free() {
	this.Ref().Free()
}

// Cols returns the value of property "HTMLFrameSetElement.cols".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameSetElement) Cols() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameSetElementCols(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCols sets the value of property "HTMLFrameSetElement.cols" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameSetElement) SetCols(val js.String) bool {
	return js.True == bindings.SetHTMLFrameSetElementCols(
		this.Ref(),
		val.Ref(),
	)
}

// Rows returns the value of property "HTMLFrameSetElement.rows".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameSetElement) Rows() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameSetElementRows(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRows sets the value of property "HTMLFrameSetElement.rows" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameSetElement) SetRows(val js.String) bool {
	return js.True == bindings.SetHTMLFrameSetElementRows(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLHRElement() (ret HTMLHRElement) {
	ret.ref = bindings.NewHTMLHRElementByHTMLHRElement()
	return
}

type HTMLHRElement struct {
	HTMLElement
}

func (this HTMLHRElement) Once() HTMLHRElement {
	this.Ref().Once()
	return this
}

func (this HTMLHRElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLHRElement) FromRef(ref js.Ref) HTMLHRElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLHRElement) Free() {
	this.Ref().Free()
}

// Align returns the value of property "HTMLHRElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLHRElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHRElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLHRElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHRElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLHRElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

// Color returns the value of property "HTMLHRElement.color".
//
// It returns ok=false if there is no such property.
func (this HTMLHRElement) Color() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHRElementColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetColor sets the value of property "HTMLHRElement.color" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHRElement) SetColor(val js.String) bool {
	return js.True == bindings.SetHTMLHRElementColor(
		this.Ref(),
		val.Ref(),
	)
}

// NoShade returns the value of property "HTMLHRElement.noShade".
//
// It returns ok=false if there is no such property.
func (this HTMLHRElement) NoShade() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLHRElementNoShade(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetNoShade sets the value of property "HTMLHRElement.noShade" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHRElement) SetNoShade(val bool) bool {
	return js.True == bindings.SetHTMLHRElementNoShade(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Size returns the value of property "HTMLHRElement.size".
//
// It returns ok=false if there is no such property.
func (this HTMLHRElement) Size() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHRElementSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSize sets the value of property "HTMLHRElement.size" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHRElement) SetSize(val js.String) bool {
	return js.True == bindings.SetHTMLHRElementSize(
		this.Ref(),
		val.Ref(),
	)
}

// Width returns the value of property "HTMLHRElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLHRElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHRElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLHRElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHRElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLHRElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLHeadingElement() (ret HTMLHeadingElement) {
	ret.ref = bindings.NewHTMLHeadingElementByHTMLHeadingElement()
	return
}

type HTMLHeadingElement struct {
	HTMLElement
}

func (this HTMLHeadingElement) Once() HTMLHeadingElement {
	this.Ref().Once()
	return this
}

func (this HTMLHeadingElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLHeadingElement) FromRef(ref js.Ref) HTMLHeadingElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLHeadingElement) Free() {
	this.Ref().Free()
}

// Align returns the value of property "HTMLHeadingElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLHeadingElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHeadingElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLHeadingElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHeadingElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLHeadingElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLHtmlElement() (ret HTMLHtmlElement) {
	ret.ref = bindings.NewHTMLHtmlElementByHTMLHtmlElement()
	return
}

type HTMLHtmlElement struct {
	HTMLElement
}

func (this HTMLHtmlElement) Once() HTMLHtmlElement {
	this.Ref().Once()
	return this
}

func (this HTMLHtmlElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLHtmlElement) FromRef(ref js.Ref) HTMLHtmlElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLHtmlElement) Free() {
	this.Ref().Free()
}

// Version returns the value of property "HTMLHtmlElement.version".
//
// It returns ok=false if there is no such property.
func (this HTMLHtmlElement) Version() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHtmlElementVersion(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVersion sets the value of property "HTMLHtmlElement.version" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHtmlElement) SetVersion(val js.String) bool {
	return js.True == bindings.SetHTMLHtmlElementVersion(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLIFrameElement() (ret HTMLIFrameElement) {
	ret.ref = bindings.NewHTMLIFrameElementByHTMLIFrameElement()
	return
}

type HTMLIFrameElement struct {
	HTMLElement
}

func (this HTMLIFrameElement) Once() HTMLIFrameElement {
	this.Ref().Once()
	return this
}

func (this HTMLIFrameElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLIFrameElement) FromRef(ref js.Ref) HTMLIFrameElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLIFrameElement) Free() {
	this.Ref().Free()
}

// Src returns the value of property "HTMLIFrameElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLIFrameElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementSrc(
		this.Ref(),
		val.Ref(),
	)
}

// Srcdoc returns the value of property "HTMLIFrameElement.srcdoc".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Srcdoc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementSrcdoc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrcdoc sets the value of property "HTMLIFrameElement.srcdoc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetSrcdoc(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementSrcdoc(
		this.Ref(),
		val.Ref(),
	)
}

// Name returns the value of property "HTMLIFrameElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLIFrameElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Sandbox returns the value of property "HTMLIFrameElement.sandbox".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Sandbox() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementSandbox(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Allow returns the value of property "HTMLIFrameElement.allow".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Allow() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementAllow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAllow sets the value of property "HTMLIFrameElement.allow" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetAllow(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementAllow(
		this.Ref(),
		val.Ref(),
	)
}

// AllowFullscreen returns the value of property "HTMLIFrameElement.allowFullscreen".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) AllowFullscreen() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementAllowFullscreen(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAllowFullscreen sets the value of property "HTMLIFrameElement.allowFullscreen" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetAllowFullscreen(val bool) bool {
	return js.True == bindings.SetHTMLIFrameElementAllowFullscreen(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Width returns the value of property "HTMLIFrameElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLIFrameElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

// Height returns the value of property "HTMLIFrameElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLIFrameElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementHeight(
		this.Ref(),
		val.Ref(),
	)
}

// ReferrerPolicy returns the value of property "HTMLIFrameElement.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementReferrerPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLIFrameElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementReferrerPolicy(
		this.Ref(),
		val.Ref(),
	)
}

// Loading returns the value of property "HTMLIFrameElement.loading".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Loading() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementLoading(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLoading sets the value of property "HTMLIFrameElement.loading" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetLoading(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementLoading(
		this.Ref(),
		val.Ref(),
	)
}

// ContentDocument returns the value of property "HTMLIFrameElement.contentDocument".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) ContentDocument() (ret Document, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementContentDocument(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContentWindow returns the value of property "HTMLIFrameElement.contentWindow".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) ContentWindow() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementContentWindow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// PermissionsPolicy returns the value of property "HTMLIFrameElement.permissionsPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) PermissionsPolicy() (ret PermissionsPolicy, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementPermissionsPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Csp returns the value of property "HTMLIFrameElement.csp".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Csp() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementCsp(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCsp sets the value of property "HTMLIFrameElement.csp" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetCsp(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementCsp(
		this.Ref(),
		val.Ref(),
	)
}

// PrivateToken returns the value of property "HTMLIFrameElement.privateToken".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) PrivateToken() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementPrivateToken(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPrivateToken sets the value of property "HTMLIFrameElement.privateToken" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetPrivateToken(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementPrivateToken(
		this.Ref(),
		val.Ref(),
	)
}

// Align returns the value of property "HTMLIFrameElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLIFrameElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

// Scrolling returns the value of property "HTMLIFrameElement.scrolling".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Scrolling() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementScrolling(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetScrolling sets the value of property "HTMLIFrameElement.scrolling" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetScrolling(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementScrolling(
		this.Ref(),
		val.Ref(),
	)
}

// FrameBorder returns the value of property "HTMLIFrameElement.frameBorder".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) FrameBorder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementFrameBorder(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFrameBorder sets the value of property "HTMLIFrameElement.frameBorder" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetFrameBorder(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementFrameBorder(
		this.Ref(),
		val.Ref(),
	)
}

// LongDesc returns the value of property "HTMLIFrameElement.longDesc".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) LongDesc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementLongDesc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLongDesc sets the value of property "HTMLIFrameElement.longDesc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetLongDesc(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementLongDesc(
		this.Ref(),
		val.Ref(),
	)
}

// MarginHeight returns the value of property "HTMLIFrameElement.marginHeight".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) MarginHeight() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementMarginHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMarginHeight sets the value of property "HTMLIFrameElement.marginHeight" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetMarginHeight(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementMarginHeight(
		this.Ref(),
		val.Ref(),
	)
}

// MarginWidth returns the value of property "HTMLIFrameElement.marginWidth".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) MarginWidth() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementMarginWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMarginWidth sets the value of property "HTMLIFrameElement.marginWidth" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetMarginWidth(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementMarginWidth(
		this.Ref(),
		val.Ref(),
	)
}

// HasGetSVGDocument returns true if the method "HTMLIFrameElement.getSVGDocument" exists.
func (this HTMLIFrameElement) HasGetSVGDocument() bool {
	return js.True == bindings.HasHTMLIFrameElementGetSVGDocument(
		this.Ref(),
	)
}

// GetSVGDocumentFunc returns the method "HTMLIFrameElement.getSVGDocument".
func (this HTMLIFrameElement) GetSVGDocumentFunc() (fn js.Func[func() Document]) {
	return fn.FromRef(
		bindings.HTMLIFrameElementGetSVGDocumentFunc(
			this.Ref(),
		),
	)
}

// GetSVGDocument calls the method "HTMLIFrameElement.getSVGDocument".
func (this HTMLIFrameElement) GetSVGDocument() (ret Document) {
	bindings.CallHTMLIFrameElementGetSVGDocument(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSVGDocument calls the method "HTMLIFrameElement.getSVGDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLIFrameElement) TryGetSVGDocument() (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLIFrameElementGetSVGDocument(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type SelectionMode uint32

const (
	_ SelectionMode = iota

	SelectionMode_SELECT
	SelectionMode_START
	SelectionMode_END
	SelectionMode_PRESERVE
)

func (SelectionMode) FromRef(str js.Ref) SelectionMode {
	return SelectionMode(bindings.ConstOfSelectionMode(str))
}

func (x SelectionMode) String() (string, bool) {
	switch x {
	case SelectionMode_SELECT:
		return "select", true
	case SelectionMode_START:
		return "start", true
	case SelectionMode_END:
		return "end", true
	case SelectionMode_PRESERVE:
		return "preserve", true
	default:
		return "", false
	}
}

func NewHTMLInputElement() (ret HTMLInputElement) {
	ret.ref = bindings.NewHTMLInputElementByHTMLInputElement()
	return
}

type HTMLInputElement struct {
	HTMLElement
}

func (this HTMLInputElement) Once() HTMLInputElement {
	this.Ref().Once()
	return this
}

func (this HTMLInputElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLInputElement) FromRef(ref js.Ref) HTMLInputElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLInputElement) Free() {
	this.Ref().Free()
}

// Accept returns the value of property "HTMLInputElement.accept".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Accept() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementAccept(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAccept sets the value of property "HTMLInputElement.accept" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetAccept(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementAccept(
		this.Ref(),
		val.Ref(),
	)
}

// Alt returns the value of property "HTMLInputElement.alt".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Alt() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementAlt(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlt sets the value of property "HTMLInputElement.alt" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetAlt(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementAlt(
		this.Ref(),
		val.Ref(),
	)
}

// Autocomplete returns the value of property "HTMLInputElement.autocomplete".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Autocomplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementAutocomplete(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAutocomplete sets the value of property "HTMLInputElement.autocomplete" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetAutocomplete(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementAutocomplete(
		this.Ref(),
		val.Ref(),
	)
}

// DefaultChecked returns the value of property "HTMLInputElement.defaultChecked".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) DefaultChecked() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementDefaultChecked(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDefaultChecked sets the value of property "HTMLInputElement.defaultChecked" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetDefaultChecked(val bool) bool {
	return js.True == bindings.SetHTMLInputElementDefaultChecked(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Checked returns the value of property "HTMLInputElement.checked".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Checked() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementChecked(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetChecked sets the value of property "HTMLInputElement.checked" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetChecked(val bool) bool {
	return js.True == bindings.SetHTMLInputElementChecked(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// DirName returns the value of property "HTMLInputElement.dirName".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) DirName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementDirName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDirName sets the value of property "HTMLInputElement.dirName" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetDirName(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementDirName(
		this.Ref(),
		val.Ref(),
	)
}

// Disabled returns the value of property "HTMLInputElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementDisabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLInputElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLInputElementDisabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLInputElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementForm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Files returns the value of property "HTMLInputElement.files".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Files() (ret FileList, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFiles(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFiles sets the value of property "HTMLInputElement.files" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFiles(val FileList) bool {
	return js.True == bindings.SetHTMLInputElementFiles(
		this.Ref(),
		val.Ref(),
	)
}

// FormAction returns the value of property "HTMLInputElement.formAction".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) FormAction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFormAction(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFormAction sets the value of property "HTMLInputElement.formAction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFormAction(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementFormAction(
		this.Ref(),
		val.Ref(),
	)
}

// FormEnctype returns the value of property "HTMLInputElement.formEnctype".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) FormEnctype() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFormEnctype(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFormEnctype sets the value of property "HTMLInputElement.formEnctype" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFormEnctype(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementFormEnctype(
		this.Ref(),
		val.Ref(),
	)
}

// FormMethod returns the value of property "HTMLInputElement.formMethod".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) FormMethod() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFormMethod(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFormMethod sets the value of property "HTMLInputElement.formMethod" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFormMethod(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementFormMethod(
		this.Ref(),
		val.Ref(),
	)
}

// FormNoValidate returns the value of property "HTMLInputElement.formNoValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) FormNoValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFormNoValidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFormNoValidate sets the value of property "HTMLInputElement.formNoValidate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFormNoValidate(val bool) bool {
	return js.True == bindings.SetHTMLInputElementFormNoValidate(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// FormTarget returns the value of property "HTMLInputElement.formTarget".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) FormTarget() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFormTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFormTarget sets the value of property "HTMLInputElement.formTarget" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFormTarget(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementFormTarget(
		this.Ref(),
		val.Ref(),
	)
}

// Height returns the value of property "HTMLInputElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLInputElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLInputElementHeight(
		this.Ref(),
		uint32(val),
	)
}

// Indeterminate returns the value of property "HTMLInputElement.indeterminate".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Indeterminate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementIndeterminate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetIndeterminate sets the value of property "HTMLInputElement.indeterminate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetIndeterminate(val bool) bool {
	return js.True == bindings.SetHTMLInputElementIndeterminate(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// List returns the value of property "HTMLInputElement.list".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) List() (ret HTMLDataListElement, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementList(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Max returns the value of property "HTMLInputElement.max".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Max() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementMax(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMax sets the value of property "HTMLInputElement.max" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetMax(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementMax(
		this.Ref(),
		val.Ref(),
	)
}

// MaxLength returns the value of property "HTMLInputElement.maxLength".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) MaxLength() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementMaxLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMaxLength sets the value of property "HTMLInputElement.maxLength" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetMaxLength(val int32) bool {
	return js.True == bindings.SetHTMLInputElementMaxLength(
		this.Ref(),
		int32(val),
	)
}

// Min returns the value of property "HTMLInputElement.min".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Min() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementMin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMin sets the value of property "HTMLInputElement.min" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetMin(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementMin(
		this.Ref(),
		val.Ref(),
	)
}

// MinLength returns the value of property "HTMLInputElement.minLength".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) MinLength() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementMinLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMinLength sets the value of property "HTMLInputElement.minLength" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetMinLength(val int32) bool {
	return js.True == bindings.SetHTMLInputElementMinLength(
		this.Ref(),
		int32(val),
	)
}

// Multiple returns the value of property "HTMLInputElement.multiple".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Multiple() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementMultiple(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMultiple sets the value of property "HTMLInputElement.multiple" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetMultiple(val bool) bool {
	return js.True == bindings.SetHTMLInputElementMultiple(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Name returns the value of property "HTMLInputElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLInputElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Pattern returns the value of property "HTMLInputElement.pattern".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Pattern() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementPattern(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPattern sets the value of property "HTMLInputElement.pattern" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetPattern(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementPattern(
		this.Ref(),
		val.Ref(),
	)
}

// Placeholder returns the value of property "HTMLInputElement.placeholder".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Placeholder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementPlaceholder(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPlaceholder sets the value of property "HTMLInputElement.placeholder" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetPlaceholder(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementPlaceholder(
		this.Ref(),
		val.Ref(),
	)
}

// ReadOnly returns the value of property "HTMLInputElement.readOnly".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) ReadOnly() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementReadOnly(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReadOnly sets the value of property "HTMLInputElement.readOnly" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetReadOnly(val bool) bool {
	return js.True == bindings.SetHTMLInputElementReadOnly(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Required returns the value of property "HTMLInputElement.required".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Required() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementRequired(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRequired sets the value of property "HTMLInputElement.required" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetRequired(val bool) bool {
	return js.True == bindings.SetHTMLInputElementRequired(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Size returns the value of property "HTMLInputElement.size".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Size() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSize sets the value of property "HTMLInputElement.size" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetSize(val uint32) bool {
	return js.True == bindings.SetHTMLInputElementSize(
		this.Ref(),
		uint32(val),
	)
}

// Src returns the value of property "HTMLInputElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLInputElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementSrc(
		this.Ref(),
		val.Ref(),
	)
}

// Step returns the value of property "HTMLInputElement.step".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Step() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementStep(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetStep sets the value of property "HTMLInputElement.step" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetStep(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementStep(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLInputElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLInputElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementType(
		this.Ref(),
		val.Ref(),
	)
}

// DefaultValue returns the value of property "HTMLInputElement.defaultValue".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) DefaultValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementDefaultValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDefaultValue sets the value of property "HTMLInputElement.defaultValue" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetDefaultValue(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementDefaultValue(
		this.Ref(),
		val.Ref(),
	)
}

// Value returns the value of property "HTMLInputElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLInputElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementValue(
		this.Ref(),
		val.Ref(),
	)
}

// ValueAsDate returns the value of property "HTMLInputElement.valueAsDate".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) ValueAsDate() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementValueAsDate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValueAsDate sets the value of property "HTMLInputElement.valueAsDate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetValueAsDate(val js.Object) bool {
	return js.True == bindings.SetHTMLInputElementValueAsDate(
		this.Ref(),
		val.Ref(),
	)
}

// ValueAsNumber returns the value of property "HTMLInputElement.valueAsNumber".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) ValueAsNumber() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementValueAsNumber(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValueAsNumber sets the value of property "HTMLInputElement.valueAsNumber" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetValueAsNumber(val float64) bool {
	return js.True == bindings.SetHTMLInputElementValueAsNumber(
		this.Ref(),
		float64(val),
	)
}

// Width returns the value of property "HTMLInputElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLInputElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetWidth(val uint32) bool {
	return js.True == bindings.SetHTMLInputElementWidth(
		this.Ref(),
		uint32(val),
	)
}

// WillValidate returns the value of property "HTMLInputElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementWillValidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLInputElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementValidity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLInputElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementValidationMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLInputElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementLabels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SelectionStart returns the value of property "HTMLInputElement.selectionStart".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) SelectionStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementSelectionStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelectionStart sets the value of property "HTMLInputElement.selectionStart" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetSelectionStart(val uint32) bool {
	return js.True == bindings.SetHTMLInputElementSelectionStart(
		this.Ref(),
		uint32(val),
	)
}

// SelectionEnd returns the value of property "HTMLInputElement.selectionEnd".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) SelectionEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementSelectionEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelectionEnd sets the value of property "HTMLInputElement.selectionEnd" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetSelectionEnd(val uint32) bool {
	return js.True == bindings.SetHTMLInputElementSelectionEnd(
		this.Ref(),
		uint32(val),
	)
}

// SelectionDirection returns the value of property "HTMLInputElement.selectionDirection".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) SelectionDirection() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementSelectionDirection(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelectionDirection sets the value of property "HTMLInputElement.selectionDirection" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetSelectionDirection(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementSelectionDirection(
		this.Ref(),
		val.Ref(),
	)
}

// Align returns the value of property "HTMLInputElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLInputElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

// UseMap returns the value of property "HTMLInputElement.useMap".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) UseMap() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementUseMap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetUseMap sets the value of property "HTMLInputElement.useMap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetUseMap(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementUseMap(
		this.Ref(),
		val.Ref(),
	)
}

// Webkitdirectory returns the value of property "HTMLInputElement.webkitdirectory".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Webkitdirectory() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementWebkitdirectory(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWebkitdirectory sets the value of property "HTMLInputElement.webkitdirectory" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetWebkitdirectory(val bool) bool {
	return js.True == bindings.SetHTMLInputElementWebkitdirectory(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// WebkitEntries returns the value of property "HTMLInputElement.webkitEntries".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) WebkitEntries() (ret js.FrozenArray[FileSystemEntry], ok bool) {
	ok = js.True == bindings.GetHTMLInputElementWebkitEntries(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Capture returns the value of property "HTMLInputElement.capture".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Capture() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementCapture(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCapture sets the value of property "HTMLInputElement.capture" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetCapture(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementCapture(
		this.Ref(),
		val.Ref(),
	)
}

// PopoverTargetElement returns the value of property "HTMLInputElement.popoverTargetElement".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) PopoverTargetElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementPopoverTargetElement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPopoverTargetElement sets the value of property "HTMLInputElement.popoverTargetElement" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetPopoverTargetElement(val Element) bool {
	return js.True == bindings.SetHTMLInputElementPopoverTargetElement(
		this.Ref(),
		val.Ref(),
	)
}

// PopoverTargetAction returns the value of property "HTMLInputElement.popoverTargetAction".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) PopoverTargetAction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementPopoverTargetAction(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPopoverTargetAction sets the value of property "HTMLInputElement.popoverTargetAction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetPopoverTargetAction(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementPopoverTargetAction(
		this.Ref(),
		val.Ref(),
	)
}

// HasStepUp returns true if the method "HTMLInputElement.stepUp" exists.
func (this HTMLInputElement) HasStepUp() bool {
	return js.True == bindings.HasHTMLInputElementStepUp(
		this.Ref(),
	)
}

// StepUpFunc returns the method "HTMLInputElement.stepUp".
func (this HTMLInputElement) StepUpFunc() (fn js.Func[func(n int32)]) {
	return fn.FromRef(
		bindings.HTMLInputElementStepUpFunc(
			this.Ref(),
		),
	)
}

// StepUp calls the method "HTMLInputElement.stepUp".
func (this HTMLInputElement) StepUp(n int32) (ret js.Void) {
	bindings.CallHTMLInputElementStepUp(
		this.Ref(), js.Pointer(&ret),
		int32(n),
	)

	return
}

// TryStepUp calls the method "HTMLInputElement.stepUp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryStepUp(n int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementStepUp(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(n),
	)

	return
}

// HasStepUp1 returns true if the method "HTMLInputElement.stepUp" exists.
func (this HTMLInputElement) HasStepUp1() bool {
	return js.True == bindings.HasHTMLInputElementStepUp1(
		this.Ref(),
	)
}

// StepUp1Func returns the method "HTMLInputElement.stepUp".
func (this HTMLInputElement) StepUp1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLInputElementStepUp1Func(
			this.Ref(),
		),
	)
}

// StepUp1 calls the method "HTMLInputElement.stepUp".
func (this HTMLInputElement) StepUp1() (ret js.Void) {
	bindings.CallHTMLInputElementStepUp1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStepUp1 calls the method "HTMLInputElement.stepUp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryStepUp1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementStepUp1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStepDown returns true if the method "HTMLInputElement.stepDown" exists.
func (this HTMLInputElement) HasStepDown() bool {
	return js.True == bindings.HasHTMLInputElementStepDown(
		this.Ref(),
	)
}

// StepDownFunc returns the method "HTMLInputElement.stepDown".
func (this HTMLInputElement) StepDownFunc() (fn js.Func[func(n int32)]) {
	return fn.FromRef(
		bindings.HTMLInputElementStepDownFunc(
			this.Ref(),
		),
	)
}

// StepDown calls the method "HTMLInputElement.stepDown".
func (this HTMLInputElement) StepDown(n int32) (ret js.Void) {
	bindings.CallHTMLInputElementStepDown(
		this.Ref(), js.Pointer(&ret),
		int32(n),
	)

	return
}

// TryStepDown calls the method "HTMLInputElement.stepDown"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryStepDown(n int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementStepDown(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(n),
	)

	return
}

// HasStepDown1 returns true if the method "HTMLInputElement.stepDown" exists.
func (this HTMLInputElement) HasStepDown1() bool {
	return js.True == bindings.HasHTMLInputElementStepDown1(
		this.Ref(),
	)
}

// StepDown1Func returns the method "HTMLInputElement.stepDown".
func (this HTMLInputElement) StepDown1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLInputElementStepDown1Func(
			this.Ref(),
		),
	)
}

// StepDown1 calls the method "HTMLInputElement.stepDown".
func (this HTMLInputElement) StepDown1() (ret js.Void) {
	bindings.CallHTMLInputElementStepDown1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStepDown1 calls the method "HTMLInputElement.stepDown"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryStepDown1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementStepDown1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCheckValidity returns true if the method "HTMLInputElement.checkValidity" exists.
func (this HTMLInputElement) HasCheckValidity() bool {
	return js.True == bindings.HasHTMLInputElementCheckValidity(
		this.Ref(),
	)
}

// CheckValidityFunc returns the method "HTMLInputElement.checkValidity".
func (this HTMLInputElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLInputElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLInputElement.checkValidity".
func (this HTMLInputElement) CheckValidity() (ret bool) {
	bindings.CallHTMLInputElementCheckValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLInputElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementCheckValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReportValidity returns true if the method "HTMLInputElement.reportValidity" exists.
func (this HTMLInputElement) HasReportValidity() bool {
	return js.True == bindings.HasHTMLInputElementReportValidity(
		this.Ref(),
	)
}

// ReportValidityFunc returns the method "HTMLInputElement.reportValidity".
func (this HTMLInputElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLInputElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLInputElement.reportValidity".
func (this HTMLInputElement) ReportValidity() (ret bool) {
	bindings.CallHTMLInputElementReportValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLInputElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementReportValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetCustomValidity returns true if the method "HTMLInputElement.setCustomValidity" exists.
func (this HTMLInputElement) HasSetCustomValidity() bool {
	return js.True == bindings.HasHTMLInputElementSetCustomValidity(
		this.Ref(),
	)
}

// SetCustomValidityFunc returns the method "HTMLInputElement.setCustomValidity".
func (this HTMLInputElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLInputElement.setCustomValidity".
func (this HTMLInputElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLInputElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLInputElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

// HasSelect returns true if the method "HTMLInputElement.select" exists.
func (this HTMLInputElement) HasSelect() bool {
	return js.True == bindings.HasHTMLInputElementSelect(
		this.Ref(),
	)
}

// SelectFunc returns the method "HTMLInputElement.select".
func (this HTMLInputElement) SelectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLInputElementSelectFunc(
			this.Ref(),
		),
	)
}

// Select calls the method "HTMLInputElement.select".
func (this HTMLInputElement) Select() (ret js.Void) {
	bindings.CallHTMLInputElementSelect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySelect calls the method "HTMLInputElement.select"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TrySelect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementSelect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetRangeText returns true if the method "HTMLInputElement.setRangeText" exists.
func (this HTMLInputElement) HasSetRangeText() bool {
	return js.True == bindings.HasHTMLInputElementSetRangeText(
		this.Ref(),
	)
}

// SetRangeTextFunc returns the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) SetRangeTextFunc() (fn js.Func[func(replacement js.String)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetRangeTextFunc(
			this.Ref(),
		),
	)
}

// SetRangeText calls the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) SetRangeText(replacement js.String) (ret js.Void) {
	bindings.CallHTMLInputElementSetRangeText(
		this.Ref(), js.Pointer(&ret),
		replacement.Ref(),
	)

	return
}

// TrySetRangeText calls the method "HTMLInputElement.setRangeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TrySetRangeText(replacement js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementSetRangeText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
	)

	return
}

// HasSetRangeText1 returns true if the method "HTMLInputElement.setRangeText" exists.
func (this HTMLInputElement) HasSetRangeText1() bool {
	return js.True == bindings.HasHTMLInputElementSetRangeText1(
		this.Ref(),
	)
}

// SetRangeText1Func returns the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) SetRangeText1Func() (fn js.Func[func(replacement js.String, start uint32, end uint32, selectionMode SelectionMode)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetRangeText1Func(
			this.Ref(),
		),
	)
}

// SetRangeText1 calls the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) SetRangeText1(replacement js.String, start uint32, end uint32, selectionMode SelectionMode) (ret js.Void) {
	bindings.CallHTMLInputElementSetRangeText1(
		this.Ref(), js.Pointer(&ret),
		replacement.Ref(),
		uint32(start),
		uint32(end),
		uint32(selectionMode),
	)

	return
}

// TrySetRangeText1 calls the method "HTMLInputElement.setRangeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TrySetRangeText1(replacement js.String, start uint32, end uint32, selectionMode SelectionMode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementSetRangeText1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
		uint32(start),
		uint32(end),
		uint32(selectionMode),
	)

	return
}

// HasSetRangeText2 returns true if the method "HTMLInputElement.setRangeText" exists.
func (this HTMLInputElement) HasSetRangeText2() bool {
	return js.True == bindings.HasHTMLInputElementSetRangeText2(
		this.Ref(),
	)
}

// SetRangeText2Func returns the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) SetRangeText2Func() (fn js.Func[func(replacement js.String, start uint32, end uint32)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetRangeText2Func(
			this.Ref(),
		),
	)
}

// SetRangeText2 calls the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) SetRangeText2(replacement js.String, start uint32, end uint32) (ret js.Void) {
	bindings.CallHTMLInputElementSetRangeText2(
		this.Ref(), js.Pointer(&ret),
		replacement.Ref(),
		uint32(start),
		uint32(end),
	)

	return
}

// TrySetRangeText2 calls the method "HTMLInputElement.setRangeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TrySetRangeText2(replacement js.String, start uint32, end uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementSetRangeText2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
		uint32(start),
		uint32(end),
	)

	return
}

// HasSetSelectionRange returns true if the method "HTMLInputElement.setSelectionRange" exists.
func (this HTMLInputElement) HasSetSelectionRange() bool {
	return js.True == bindings.HasHTMLInputElementSetSelectionRange(
		this.Ref(),
	)
}

// SetSelectionRangeFunc returns the method "HTMLInputElement.setSelectionRange".
func (this HTMLInputElement) SetSelectionRangeFunc() (fn js.Func[func(start uint32, end uint32, direction js.String)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetSelectionRangeFunc(
			this.Ref(),
		),
	)
}

// SetSelectionRange calls the method "HTMLInputElement.setSelectionRange".
func (this HTMLInputElement) SetSelectionRange(start uint32, end uint32, direction js.String) (ret js.Void) {
	bindings.CallHTMLInputElementSetSelectionRange(
		this.Ref(), js.Pointer(&ret),
		uint32(start),
		uint32(end),
		direction.Ref(),
	)

	return
}

// TrySetSelectionRange calls the method "HTMLInputElement.setSelectionRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TrySetSelectionRange(start uint32, end uint32, direction js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementSetSelectionRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(start),
		uint32(end),
		direction.Ref(),
	)

	return
}

// HasSetSelectionRange1 returns true if the method "HTMLInputElement.setSelectionRange" exists.
func (this HTMLInputElement) HasSetSelectionRange1() bool {
	return js.True == bindings.HasHTMLInputElementSetSelectionRange1(
		this.Ref(),
	)
}

// SetSelectionRange1Func returns the method "HTMLInputElement.setSelectionRange".
func (this HTMLInputElement) SetSelectionRange1Func() (fn js.Func[func(start uint32, end uint32)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetSelectionRange1Func(
			this.Ref(),
		),
	)
}

// SetSelectionRange1 calls the method "HTMLInputElement.setSelectionRange".
func (this HTMLInputElement) SetSelectionRange1(start uint32, end uint32) (ret js.Void) {
	bindings.CallHTMLInputElementSetSelectionRange1(
		this.Ref(), js.Pointer(&ret),
		uint32(start),
		uint32(end),
	)

	return
}

// TrySetSelectionRange1 calls the method "HTMLInputElement.setSelectionRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TrySetSelectionRange1(start uint32, end uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementSetSelectionRange1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(start),
		uint32(end),
	)

	return
}

// HasShowPicker returns true if the method "HTMLInputElement.showPicker" exists.
func (this HTMLInputElement) HasShowPicker() bool {
	return js.True == bindings.HasHTMLInputElementShowPicker(
		this.Ref(),
	)
}

// ShowPickerFunc returns the method "HTMLInputElement.showPicker".
func (this HTMLInputElement) ShowPickerFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLInputElementShowPickerFunc(
			this.Ref(),
		),
	)
}

// ShowPicker calls the method "HTMLInputElement.showPicker".
func (this HTMLInputElement) ShowPicker() (ret js.Void) {
	bindings.CallHTMLInputElementShowPicker(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryShowPicker calls the method "HTMLInputElement.showPicker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryShowPicker() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementShowPicker(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewHTMLLIElement() (ret HTMLLIElement) {
	ret.ref = bindings.NewHTMLLIElementByHTMLLIElement()
	return
}

type HTMLLIElement struct {
	HTMLElement
}

func (this HTMLLIElement) Once() HTMLLIElement {
	this.Ref().Once()
	return this
}

func (this HTMLLIElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLLIElement) FromRef(ref js.Ref) HTMLLIElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLLIElement) Free() {
	this.Ref().Free()
}

// Value returns the value of property "HTMLLIElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLLIElement) Value() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLLIElementValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLLIElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLIElement) SetValue(val int32) bool {
	return js.True == bindings.SetHTMLLIElementValue(
		this.Ref(),
		int32(val),
	)
}

// Type returns the value of property "HTMLLIElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLLIElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLIElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLLIElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLIElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLLIElementType(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLLabelElement() (ret HTMLLabelElement) {
	ret.ref = bindings.NewHTMLLabelElementByHTMLLabelElement()
	return
}

type HTMLLabelElement struct {
	HTMLElement
}

func (this HTMLLabelElement) Once() HTMLLabelElement {
	this.Ref().Once()
	return this
}

func (this HTMLLabelElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLLabelElement) FromRef(ref js.Ref) HTMLLabelElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLLabelElement) Free() {
	this.Ref().Free()
}

// Form returns the value of property "HTMLLabelElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLLabelElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLLabelElementForm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HtmlFor returns the value of property "HTMLLabelElement.htmlFor".
//
// It returns ok=false if there is no such property.
func (this HTMLLabelElement) HtmlFor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLabelElementHtmlFor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHtmlFor sets the value of property "HTMLLabelElement.htmlFor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLabelElement) SetHtmlFor(val js.String) bool {
	return js.True == bindings.SetHTMLLabelElementHtmlFor(
		this.Ref(),
		val.Ref(),
	)
}

// Control returns the value of property "HTMLLabelElement.control".
//
// It returns ok=false if there is no such property.
func (this HTMLLabelElement) Control() (ret HTMLElement, ok bool) {
	ok = js.True == bindings.GetHTMLLabelElementControl(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewHTMLLegendElement() (ret HTMLLegendElement) {
	ret.ref = bindings.NewHTMLLegendElementByHTMLLegendElement()
	return
}

type HTMLLegendElement struct {
	HTMLElement
}

func (this HTMLLegendElement) Once() HTMLLegendElement {
	this.Ref().Once()
	return this
}

func (this HTMLLegendElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLLegendElement) FromRef(ref js.Ref) HTMLLegendElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLLegendElement) Free() {
	this.Ref().Free()
}

// Form returns the value of property "HTMLLegendElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLLegendElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLLegendElementForm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Align returns the value of property "HTMLLegendElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLLegendElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLegendElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLLegendElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLegendElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLLegendElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLLinkElement() (ret HTMLLinkElement) {
	ret.ref = bindings.NewHTMLLinkElementByHTMLLinkElement()
	return
}

type HTMLLinkElement struct {
	HTMLElement
}

func (this HTMLLinkElement) Once() HTMLLinkElement {
	this.Ref().Once()
	return this
}

func (this HTMLLinkElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLLinkElement) FromRef(ref js.Ref) HTMLLinkElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLLinkElement) Free() {
	this.Ref().Free()
}

// Href returns the value of property "HTMLLinkElement.href".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementHref(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHref sets the value of property "HTMLLinkElement.href" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetHref(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementHref(
		this.Ref(),
		val.Ref(),
	)
}

// CrossOrigin returns the value of property "HTMLLinkElement.crossOrigin".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) CrossOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementCrossOrigin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCrossOrigin sets the value of property "HTMLLinkElement.crossOrigin" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetCrossOrigin(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementCrossOrigin(
		this.Ref(),
		val.Ref(),
	)
}

// Rel returns the value of property "HTMLLinkElement.rel".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Rel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementRel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRel sets the value of property "HTMLLinkElement.rel" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetRel(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementRel(
		this.Ref(),
		val.Ref(),
	)
}

// As returns the value of property "HTMLLinkElement.as".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) As() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementAs(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAs sets the value of property "HTMLLinkElement.as" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetAs(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementAs(
		this.Ref(),
		val.Ref(),
	)
}

// RelList returns the value of property "HTMLLinkElement.relList".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) RelList() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementRelList(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Media returns the value of property "HTMLLinkElement.media".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementMedia(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMedia sets the value of property "HTMLLinkElement.media" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetMedia(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementMedia(
		this.Ref(),
		val.Ref(),
	)
}

// Integrity returns the value of property "HTMLLinkElement.integrity".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Integrity() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementIntegrity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetIntegrity sets the value of property "HTMLLinkElement.integrity" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetIntegrity(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementIntegrity(
		this.Ref(),
		val.Ref(),
	)
}

// Hreflang returns the value of property "HTMLLinkElement.hreflang".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Hreflang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementHreflang(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHreflang sets the value of property "HTMLLinkElement.hreflang" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetHreflang(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementHreflang(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLLinkElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLLinkElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementType(
		this.Ref(),
		val.Ref(),
	)
}

// Sizes returns the value of property "HTMLLinkElement.sizes".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Sizes() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementSizes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ImageSrcset returns the value of property "HTMLLinkElement.imageSrcset".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) ImageSrcset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementImageSrcset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetImageSrcset sets the value of property "HTMLLinkElement.imageSrcset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetImageSrcset(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementImageSrcset(
		this.Ref(),
		val.Ref(),
	)
}

// ImageSizes returns the value of property "HTMLLinkElement.imageSizes".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) ImageSizes() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementImageSizes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetImageSizes sets the value of property "HTMLLinkElement.imageSizes" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetImageSizes(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementImageSizes(
		this.Ref(),
		val.Ref(),
	)
}

// ReferrerPolicy returns the value of property "HTMLLinkElement.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementReferrerPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLLinkElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementReferrerPolicy(
		this.Ref(),
		val.Ref(),
	)
}

// Blocking returns the value of property "HTMLLinkElement.blocking".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Blocking() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementBlocking(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Disabled returns the value of property "HTMLLinkElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementDisabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLLinkElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLLinkElementDisabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// FetchPriority returns the value of property "HTMLLinkElement.fetchPriority".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) FetchPriority() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementFetchPriority(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFetchPriority sets the value of property "HTMLLinkElement.fetchPriority" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetFetchPriority(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementFetchPriority(
		this.Ref(),
		val.Ref(),
	)
}

// Charset returns the value of property "HTMLLinkElement.charset".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Charset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementCharset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCharset sets the value of property "HTMLLinkElement.charset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetCharset(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementCharset(
		this.Ref(),
		val.Ref(),
	)
}

// Rev returns the value of property "HTMLLinkElement.rev".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Rev() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementRev(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRev sets the value of property "HTMLLinkElement.rev" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetRev(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementRev(
		this.Ref(),
		val.Ref(),
	)
}

// Target returns the value of property "HTMLLinkElement.target".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Target() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementTarget(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTarget sets the value of property "HTMLLinkElement.target" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetTarget(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementTarget(
		this.Ref(),
		val.Ref(),
	)
}

// Sheet returns the value of property "HTMLLinkElement.sheet".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Sheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementSheet(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewHTMLMapElement() (ret HTMLMapElement) {
	ret.ref = bindings.NewHTMLMapElementByHTMLMapElement()
	return
}

type HTMLMapElement struct {
	HTMLElement
}

func (this HTMLMapElement) Once() HTMLMapElement {
	this.Ref().Once()
	return this
}

func (this HTMLMapElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLMapElement) FromRef(ref js.Ref) HTMLMapElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLMapElement) Free() {
	this.Ref().Free()
}

// Name returns the value of property "HTMLMapElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLMapElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMapElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLMapElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMapElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLMapElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Areas returns the value of property "HTMLMapElement.areas".
//
// It returns ok=false if there is no such property.
func (this HTMLMapElement) Areas() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLMapElementAreas(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewHTMLMarqueeElement() (ret HTMLMarqueeElement) {
	ret.ref = bindings.NewHTMLMarqueeElementByHTMLMarqueeElement()
	return
}

type HTMLMarqueeElement struct {
	HTMLElement
}

func (this HTMLMarqueeElement) Once() HTMLMarqueeElement {
	this.Ref().Once()
	return this
}

func (this HTMLMarqueeElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLMarqueeElement) FromRef(ref js.Ref) HTMLMarqueeElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLMarqueeElement) Free() {
	this.Ref().Free()
}

// Behavior returns the value of property "HTMLMarqueeElement.behavior".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Behavior() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementBehavior(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBehavior sets the value of property "HTMLMarqueeElement.behavior" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetBehavior(val js.String) bool {
	return js.True == bindings.SetHTMLMarqueeElementBehavior(
		this.Ref(),
		val.Ref(),
	)
}

// BgColor returns the value of property "HTMLMarqueeElement.bgColor".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementBgColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "HTMLMarqueeElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLMarqueeElementBgColor(
		this.Ref(),
		val.Ref(),
	)
}

// Direction returns the value of property "HTMLMarqueeElement.direction".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Direction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementDirection(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDirection sets the value of property "HTMLMarqueeElement.direction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetDirection(val js.String) bool {
	return js.True == bindings.SetHTMLMarqueeElementDirection(
		this.Ref(),
		val.Ref(),
	)
}

// Height returns the value of property "HTMLMarqueeElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLMarqueeElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLMarqueeElementHeight(
		this.Ref(),
		val.Ref(),
	)
}

// Hspace returns the value of property "HTMLMarqueeElement.hspace".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Hspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementHspace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHspace sets the value of property "HTMLMarqueeElement.hspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetHspace(val uint32) bool {
	return js.True == bindings.SetHTMLMarqueeElementHspace(
		this.Ref(),
		uint32(val),
	)
}

// Loop returns the value of property "HTMLMarqueeElement.loop".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Loop() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementLoop(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLoop sets the value of property "HTMLMarqueeElement.loop" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetLoop(val int32) bool {
	return js.True == bindings.SetHTMLMarqueeElementLoop(
		this.Ref(),
		int32(val),
	)
}

// ScrollAmount returns the value of property "HTMLMarqueeElement.scrollAmount".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) ScrollAmount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementScrollAmount(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetScrollAmount sets the value of property "HTMLMarqueeElement.scrollAmount" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetScrollAmount(val uint32) bool {
	return js.True == bindings.SetHTMLMarqueeElementScrollAmount(
		this.Ref(),
		uint32(val),
	)
}

// ScrollDelay returns the value of property "HTMLMarqueeElement.scrollDelay".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) ScrollDelay() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementScrollDelay(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetScrollDelay sets the value of property "HTMLMarqueeElement.scrollDelay" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetScrollDelay(val uint32) bool {
	return js.True == bindings.SetHTMLMarqueeElementScrollDelay(
		this.Ref(),
		uint32(val),
	)
}

// TrueSpeed returns the value of property "HTMLMarqueeElement.trueSpeed".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) TrueSpeed() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementTrueSpeed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTrueSpeed sets the value of property "HTMLMarqueeElement.trueSpeed" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetTrueSpeed(val bool) bool {
	return js.True == bindings.SetHTMLMarqueeElementTrueSpeed(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Vspace returns the value of property "HTMLMarqueeElement.vspace".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Vspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementVspace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVspace sets the value of property "HTMLMarqueeElement.vspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetVspace(val uint32) bool {
	return js.True == bindings.SetHTMLMarqueeElementVspace(
		this.Ref(),
		uint32(val),
	)
}

// Width returns the value of property "HTMLMarqueeElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLMarqueeElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLMarqueeElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

// HasStart returns true if the method "HTMLMarqueeElement.start" exists.
func (this HTMLMarqueeElement) HasStart() bool {
	return js.True == bindings.HasHTMLMarqueeElementStart(
		this.Ref(),
	)
}

// StartFunc returns the method "HTMLMarqueeElement.start".
func (this HTMLMarqueeElement) StartFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLMarqueeElementStartFunc(
			this.Ref(),
		),
	)
}

// Start calls the method "HTMLMarqueeElement.start".
func (this HTMLMarqueeElement) Start() (ret js.Void) {
	bindings.CallHTMLMarqueeElementStart(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStart calls the method "HTMLMarqueeElement.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLMarqueeElement) TryStart() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMarqueeElementStart(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasStop returns true if the method "HTMLMarqueeElement.stop" exists.
func (this HTMLMarqueeElement) HasStop() bool {
	return js.True == bindings.HasHTMLMarqueeElementStop(
		this.Ref(),
	)
}

// StopFunc returns the method "HTMLMarqueeElement.stop".
func (this HTMLMarqueeElement) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLMarqueeElementStopFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "HTMLMarqueeElement.stop".
func (this HTMLMarqueeElement) Stop() (ret js.Void) {
	bindings.CallHTMLMarqueeElementStop(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "HTMLMarqueeElement.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLMarqueeElement) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMarqueeElementStop(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

func NewHTMLMenuElement() (ret HTMLMenuElement) {
	ret.ref = bindings.NewHTMLMenuElementByHTMLMenuElement()
	return
}

type HTMLMenuElement struct {
	HTMLElement
}

func (this HTMLMenuElement) Once() HTMLMenuElement {
	this.Ref().Once()
	return this
}

func (this HTMLMenuElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLMenuElement) FromRef(ref js.Ref) HTMLMenuElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLMenuElement) Free() {
	this.Ref().Free()
}

// Compact returns the value of property "HTMLMenuElement.compact".
//
// It returns ok=false if there is no such property.
func (this HTMLMenuElement) Compact() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMenuElementCompact(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCompact sets the value of property "HTMLMenuElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMenuElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLMenuElementCompact(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

func NewHTMLMetaElement() (ret HTMLMetaElement) {
	ret.ref = bindings.NewHTMLMetaElementByHTMLMetaElement()
	return
}

type HTMLMetaElement struct {
	HTMLElement
}

func (this HTMLMetaElement) Once() HTMLMetaElement {
	this.Ref().Once()
	return this
}

func (this HTMLMetaElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLMetaElement) FromRef(ref js.Ref) HTMLMetaElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLMetaElement) Free() {
	this.Ref().Free()
}

// Name returns the value of property "HTMLMetaElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLMetaElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMetaElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLMetaElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMetaElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLMetaElementName(
		this.Ref(),
		val.Ref(),
	)
}

// HttpEquiv returns the value of property "HTMLMetaElement.httpEquiv".
//
// It returns ok=false if there is no such property.
func (this HTMLMetaElement) HttpEquiv() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMetaElementHttpEquiv(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHttpEquiv sets the value of property "HTMLMetaElement.httpEquiv" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMetaElement) SetHttpEquiv(val js.String) bool {
	return js.True == bindings.SetHTMLMetaElementHttpEquiv(
		this.Ref(),
		val.Ref(),
	)
}

// Content returns the value of property "HTMLMetaElement.content".
//
// It returns ok=false if there is no such property.
func (this HTMLMetaElement) Content() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMetaElementContent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetContent sets the value of property "HTMLMetaElement.content" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMetaElement) SetContent(val js.String) bool {
	return js.True == bindings.SetHTMLMetaElementContent(
		this.Ref(),
		val.Ref(),
	)
}

// Media returns the value of property "HTMLMetaElement.media".
//
// It returns ok=false if there is no such property.
func (this HTMLMetaElement) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMetaElementMedia(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMedia sets the value of property "HTMLMetaElement.media" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMetaElement) SetMedia(val js.String) bool {
	return js.True == bindings.SetHTMLMetaElementMedia(
		this.Ref(),
		val.Ref(),
	)
}

// Scheme returns the value of property "HTMLMetaElement.scheme".
//
// It returns ok=false if there is no such property.
func (this HTMLMetaElement) Scheme() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMetaElementScheme(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetScheme sets the value of property "HTMLMetaElement.scheme" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMetaElement) SetScheme(val js.String) bool {
	return js.True == bindings.SetHTMLMetaElementScheme(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLMeterElement() (ret HTMLMeterElement) {
	ret.ref = bindings.NewHTMLMeterElementByHTMLMeterElement()
	return
}

type HTMLMeterElement struct {
	HTMLElement
}

func (this HTMLMeterElement) Once() HTMLMeterElement {
	this.Ref().Once()
	return this
}

func (this HTMLMeterElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLMeterElement) FromRef(ref js.Ref) HTMLMeterElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLMeterElement) Free() {
	this.Ref().Free()
}

// Value returns the value of property "HTMLMeterElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Value() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLMeterElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetValue(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementValue(
		this.Ref(),
		float64(val),
	)
}

// Min returns the value of property "HTMLMeterElement.min".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Min() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementMin(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMin sets the value of property "HTMLMeterElement.min" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetMin(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementMin(
		this.Ref(),
		float64(val),
	)
}

// Max returns the value of property "HTMLMeterElement.max".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Max() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementMax(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMax sets the value of property "HTMLMeterElement.max" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetMax(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementMax(
		this.Ref(),
		float64(val),
	)
}

// Low returns the value of property "HTMLMeterElement.low".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Low() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementLow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLow sets the value of property "HTMLMeterElement.low" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetLow(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementLow(
		this.Ref(),
		float64(val),
	)
}

// High returns the value of property "HTMLMeterElement.high".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) High() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementHigh(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHigh sets the value of property "HTMLMeterElement.high" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetHigh(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementHigh(
		this.Ref(),
		float64(val),
	)
}

// Optimum returns the value of property "HTMLMeterElement.optimum".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Optimum() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementOptimum(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetOptimum sets the value of property "HTMLMeterElement.optimum" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetOptimum(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementOptimum(
		this.Ref(),
		float64(val),
	)
}

// Labels returns the value of property "HTMLMeterElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementLabels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewHTMLModElement() (ret HTMLModElement) {
	ret.ref = bindings.NewHTMLModElementByHTMLModElement()
	return
}

type HTMLModElement struct {
	HTMLElement
}

func (this HTMLModElement) Once() HTMLModElement {
	this.Ref().Once()
	return this
}

func (this HTMLModElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLModElement) FromRef(ref js.Ref) HTMLModElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLModElement) Free() {
	this.Ref().Free()
}

// Cite returns the value of property "HTMLModElement.cite".
//
// It returns ok=false if there is no such property.
func (this HTMLModElement) Cite() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLModElementCite(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCite sets the value of property "HTMLModElement.cite" to val.
//
// It returns false if the property cannot be set.
func (this HTMLModElement) SetCite(val js.String) bool {
	return js.True == bindings.SetHTMLModElementCite(
		this.Ref(),
		val.Ref(),
	)
}

// DateTime returns the value of property "HTMLModElement.dateTime".
//
// It returns ok=false if there is no such property.
func (this HTMLModElement) DateTime() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLModElementDateTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDateTime sets the value of property "HTMLModElement.dateTime" to val.
//
// It returns false if the property cannot be set.
func (this HTMLModElement) SetDateTime(val js.String) bool {
	return js.True == bindings.SetHTMLModElementDateTime(
		this.Ref(),
		val.Ref(),
	)
}

type HTMLModelElement struct {
	HTMLElement
}

func (this HTMLModelElement) Once() HTMLModelElement {
	this.Ref().Once()
	return this
}

func (this HTMLModelElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLModelElement) FromRef(ref js.Ref) HTMLModelElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLModelElement) Free() {
	this.Ref().Free()
}

func NewHTMLOListElement() (ret HTMLOListElement) {
	ret.ref = bindings.NewHTMLOListElementByHTMLOListElement()
	return
}

type HTMLOListElement struct {
	HTMLElement
}

func (this HTMLOListElement) Once() HTMLOListElement {
	this.Ref().Once()
	return this
}

func (this HTMLOListElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLOListElement) FromRef(ref js.Ref) HTMLOListElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLOListElement) Free() {
	this.Ref().Free()
}

// Reversed returns the value of property "HTMLOListElement.reversed".
//
// It returns ok=false if there is no such property.
func (this HTMLOListElement) Reversed() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOListElementReversed(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReversed sets the value of property "HTMLOListElement.reversed" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOListElement) SetReversed(val bool) bool {
	return js.True == bindings.SetHTMLOListElementReversed(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Start returns the value of property "HTMLOListElement.start".
//
// It returns ok=false if there is no such property.
func (this HTMLOListElement) Start() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLOListElementStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetStart sets the value of property "HTMLOListElement.start" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOListElement) SetStart(val int32) bool {
	return js.True == bindings.SetHTMLOListElementStart(
		this.Ref(),
		int32(val),
	)
}

// Type returns the value of property "HTMLOListElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLOListElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOListElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLOListElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOListElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLOListElementType(
		this.Ref(),
		val.Ref(),
	)
}

// Compact returns the value of property "HTMLOListElement.compact".
//
// It returns ok=false if there is no such property.
func (this HTMLOListElement) Compact() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOListElementCompact(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCompact sets the value of property "HTMLOListElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOListElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLOListElementCompact(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

func NewHTMLObjectElement() (ret HTMLObjectElement) {
	ret.ref = bindings.NewHTMLObjectElementByHTMLObjectElement()
	return
}

type HTMLObjectElement struct {
	HTMLElement
}

func (this HTMLObjectElement) Once() HTMLObjectElement {
	this.Ref().Once()
	return this
}

func (this HTMLObjectElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLObjectElement) FromRef(ref js.Ref) HTMLObjectElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLObjectElement) Free() {
	this.Ref().Free()
}

// Data returns the value of property "HTMLObjectElement.data".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Data() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementData(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetData sets the value of property "HTMLObjectElement.data" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetData(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementData(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLObjectElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLObjectElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementType(
		this.Ref(),
		val.Ref(),
	)
}

// Name returns the value of property "HTMLObjectElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLObjectElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Form returns the value of property "HTMLObjectElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementForm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "HTMLObjectElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLObjectElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

// Height returns the value of property "HTMLObjectElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLObjectElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementHeight(
		this.Ref(),
		val.Ref(),
	)
}

// ContentDocument returns the value of property "HTMLObjectElement.contentDocument".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) ContentDocument() (ret Document, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementContentDocument(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ContentWindow returns the value of property "HTMLObjectElement.contentWindow".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) ContentWindow() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementContentWindow(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WillValidate returns the value of property "HTMLObjectElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementWillValidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLObjectElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementValidity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLObjectElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementValidationMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Align returns the value of property "HTMLObjectElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLObjectElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

// Archive returns the value of property "HTMLObjectElement.archive".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Archive() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementArchive(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetArchive sets the value of property "HTMLObjectElement.archive" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetArchive(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementArchive(
		this.Ref(),
		val.Ref(),
	)
}

// Code returns the value of property "HTMLObjectElement.code".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Code() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementCode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCode sets the value of property "HTMLObjectElement.code" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetCode(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementCode(
		this.Ref(),
		val.Ref(),
	)
}

// Declare returns the value of property "HTMLObjectElement.declare".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Declare() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementDeclare(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDeclare sets the value of property "HTMLObjectElement.declare" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetDeclare(val bool) bool {
	return js.True == bindings.SetHTMLObjectElementDeclare(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Hspace returns the value of property "HTMLObjectElement.hspace".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Hspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementHspace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHspace sets the value of property "HTMLObjectElement.hspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetHspace(val uint32) bool {
	return js.True == bindings.SetHTMLObjectElementHspace(
		this.Ref(),
		uint32(val),
	)
}

// Standby returns the value of property "HTMLObjectElement.standby".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Standby() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementStandby(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetStandby sets the value of property "HTMLObjectElement.standby" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetStandby(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementStandby(
		this.Ref(),
		val.Ref(),
	)
}

// Vspace returns the value of property "HTMLObjectElement.vspace".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Vspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementVspace(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVspace sets the value of property "HTMLObjectElement.vspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetVspace(val uint32) bool {
	return js.True == bindings.SetHTMLObjectElementVspace(
		this.Ref(),
		uint32(val),
	)
}

// CodeBase returns the value of property "HTMLObjectElement.codeBase".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) CodeBase() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementCodeBase(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCodeBase sets the value of property "HTMLObjectElement.codeBase" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetCodeBase(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementCodeBase(
		this.Ref(),
		val.Ref(),
	)
}

// CodeType returns the value of property "HTMLObjectElement.codeType".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) CodeType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementCodeType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCodeType sets the value of property "HTMLObjectElement.codeType" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetCodeType(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementCodeType(
		this.Ref(),
		val.Ref(),
	)
}

// UseMap returns the value of property "HTMLObjectElement.useMap".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) UseMap() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementUseMap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetUseMap sets the value of property "HTMLObjectElement.useMap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetUseMap(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementUseMap(
		this.Ref(),
		val.Ref(),
	)
}

// Border returns the value of property "HTMLObjectElement.border".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Border() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementBorder(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBorder sets the value of property "HTMLObjectElement.border" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetBorder(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementBorder(
		this.Ref(),
		val.Ref(),
	)
}

// HasGetSVGDocument returns true if the method "HTMLObjectElement.getSVGDocument" exists.
func (this HTMLObjectElement) HasGetSVGDocument() bool {
	return js.True == bindings.HasHTMLObjectElementGetSVGDocument(
		this.Ref(),
	)
}

// GetSVGDocumentFunc returns the method "HTMLObjectElement.getSVGDocument".
func (this HTMLObjectElement) GetSVGDocumentFunc() (fn js.Func[func() Document]) {
	return fn.FromRef(
		bindings.HTMLObjectElementGetSVGDocumentFunc(
			this.Ref(),
		),
	)
}

// GetSVGDocument calls the method "HTMLObjectElement.getSVGDocument".
func (this HTMLObjectElement) GetSVGDocument() (ret Document) {
	bindings.CallHTMLObjectElementGetSVGDocument(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetSVGDocument calls the method "HTMLObjectElement.getSVGDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLObjectElement) TryGetSVGDocument() (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLObjectElementGetSVGDocument(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCheckValidity returns true if the method "HTMLObjectElement.checkValidity" exists.
func (this HTMLObjectElement) HasCheckValidity() bool {
	return js.True == bindings.HasHTMLObjectElementCheckValidity(
		this.Ref(),
	)
}

// CheckValidityFunc returns the method "HTMLObjectElement.checkValidity".
func (this HTMLObjectElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLObjectElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLObjectElement.checkValidity".
func (this HTMLObjectElement) CheckValidity() (ret bool) {
	bindings.CallHTMLObjectElementCheckValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLObjectElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLObjectElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLObjectElementCheckValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReportValidity returns true if the method "HTMLObjectElement.reportValidity" exists.
func (this HTMLObjectElement) HasReportValidity() bool {
	return js.True == bindings.HasHTMLObjectElementReportValidity(
		this.Ref(),
	)
}

// ReportValidityFunc returns the method "HTMLObjectElement.reportValidity".
func (this HTMLObjectElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLObjectElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLObjectElement.reportValidity".
func (this HTMLObjectElement) ReportValidity() (ret bool) {
	bindings.CallHTMLObjectElementReportValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLObjectElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLObjectElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLObjectElementReportValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetCustomValidity returns true if the method "HTMLObjectElement.setCustomValidity" exists.
func (this HTMLObjectElement) HasSetCustomValidity() bool {
	return js.True == bindings.HasHTMLObjectElementSetCustomValidity(
		this.Ref(),
	)
}

// SetCustomValidityFunc returns the method "HTMLObjectElement.setCustomValidity".
func (this HTMLObjectElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLObjectElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLObjectElement.setCustomValidity".
func (this HTMLObjectElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLObjectElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLObjectElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLObjectElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLObjectElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

func NewHTMLOptGroupElement() (ret HTMLOptGroupElement) {
	ret.ref = bindings.NewHTMLOptGroupElementByHTMLOptGroupElement()
	return
}

type HTMLOptGroupElement struct {
	HTMLElement
}

func (this HTMLOptGroupElement) Once() HTMLOptGroupElement {
	this.Ref().Once()
	return this
}

func (this HTMLOptGroupElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLOptGroupElement) FromRef(ref js.Ref) HTMLOptGroupElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLOptGroupElement) Free() {
	this.Ref().Free()
}

// Disabled returns the value of property "HTMLOptGroupElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLOptGroupElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOptGroupElementDisabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLOptGroupElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptGroupElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLOptGroupElementDisabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Label returns the value of property "HTMLOptGroupElement.label".
//
// It returns ok=false if there is no such property.
func (this HTMLOptGroupElement) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOptGroupElementLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "HTMLOptGroupElement.label" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptGroupElement) SetLabel(val js.String) bool {
	return js.True == bindings.SetHTMLOptGroupElementLabel(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLOptionElement() (ret HTMLOptionElement) {
	ret.ref = bindings.NewHTMLOptionElementByHTMLOptionElement()
	return
}

type HTMLOptionElement struct {
	HTMLElement
}

func (this HTMLOptionElement) Once() HTMLOptionElement {
	this.Ref().Once()
	return this
}

func (this HTMLOptionElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLOptionElement) FromRef(ref js.Ref) HTMLOptionElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLOptionElement) Free() {
	this.Ref().Free()
}

// Disabled returns the value of property "HTMLOptionElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementDisabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLOptionElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLOptionElementDisabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLOptionElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementForm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "HTMLOptionElement.label".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "HTMLOptionElement.label" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetLabel(val js.String) bool {
	return js.True == bindings.SetHTMLOptionElementLabel(
		this.Ref(),
		val.Ref(),
	)
}

// DefaultSelected returns the value of property "HTMLOptionElement.defaultSelected".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) DefaultSelected() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementDefaultSelected(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDefaultSelected sets the value of property "HTMLOptionElement.defaultSelected" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetDefaultSelected(val bool) bool {
	return js.True == bindings.SetHTMLOptionElementDefaultSelected(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Selected returns the value of property "HTMLOptionElement.selected".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Selected() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementSelected(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelected sets the value of property "HTMLOptionElement.selected" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetSelected(val bool) bool {
	return js.True == bindings.SetHTMLOptionElementSelected(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Value returns the value of property "HTMLOptionElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLOptionElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLOptionElementValue(
		this.Ref(),
		val.Ref(),
	)
}

// Text returns the value of property "HTMLOptionElement.text".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "HTMLOptionElement.text" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetText(val js.String) bool {
	return js.True == bindings.SetHTMLOptionElementText(
		this.Ref(),
		val.Ref(),
	)
}

// Index returns the value of property "HTMLOptionElement.index".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Index() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

type OneOf_HTMLOptionElement_HTMLOptGroupElement struct {
	ref js.Ref
}

func (x OneOf_HTMLOptionElement_HTMLOptGroupElement) Ref() js.Ref {
	return x.ref
}

func (x OneOf_HTMLOptionElement_HTMLOptGroupElement) Free() {
	x.ref.Free()
}

func (x OneOf_HTMLOptionElement_HTMLOptGroupElement) FromRef(ref js.Ref) OneOf_HTMLOptionElement_HTMLOptGroupElement {
	return OneOf_HTMLOptionElement_HTMLOptGroupElement{
		ref: ref,
	}
}

func (x OneOf_HTMLOptionElement_HTMLOptGroupElement) HTMLOptionElement() HTMLOptionElement {
	return HTMLOptionElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLOptionElement_HTMLOptGroupElement) HTMLOptGroupElement() HTMLOptGroupElement {
	return HTMLOptGroupElement{}.FromRef(x.ref)
}

type OneOf_HTMLElement_Int32 struct {
	ref js.Ref
}

func (x OneOf_HTMLElement_Int32) Ref() js.Ref {
	return x.ref
}

func (x OneOf_HTMLElement_Int32) Free() {
	x.ref.Free()
}

func (x OneOf_HTMLElement_Int32) FromRef(ref js.Ref) OneOf_HTMLElement_Int32 {
	return OneOf_HTMLElement_Int32{
		ref: ref,
	}
}

func (x OneOf_HTMLElement_Int32) HTMLElement() HTMLElement {
	return HTMLElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLElement_Int32) Int32() int32 {
	return js.Number[int32]{}.FromRef(x.ref).Get()
}

type HTMLOptionsCollection struct {
	HTMLCollection
}

func (this HTMLOptionsCollection) Once() HTMLOptionsCollection {
	this.Ref().Once()
	return this
}

func (this HTMLOptionsCollection) Ref() js.Ref {
	return this.HTMLCollection.Ref()
}

func (this HTMLOptionsCollection) FromRef(ref js.Ref) HTMLOptionsCollection {
	this.HTMLCollection = this.HTMLCollection.FromRef(ref)
	return this
}

func (this HTMLOptionsCollection) Free() {
	this.Ref().Free()
}

// Length returns the value of property "HTMLOptionsCollection.length".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionsCollection) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLOptionsCollectionLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLength sets the value of property "HTMLOptionsCollection.length" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionsCollection) SetLength(val uint32) bool {
	return js.True == bindings.SetHTMLOptionsCollectionLength(
		this.Ref(),
		uint32(val),
	)
}

// SelectedIndex returns the value of property "HTMLOptionsCollection.selectedIndex".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionsCollection) SelectedIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLOptionsCollectionSelectedIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelectedIndex sets the value of property "HTMLOptionsCollection.selectedIndex" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionsCollection) SetSelectedIndex(val int32) bool {
	return js.True == bindings.SetHTMLOptionsCollectionSelectedIndex(
		this.Ref(),
		int32(val),
	)
}

// HasSet returns true if the method "HTMLOptionsCollection." exists.
func (this HTMLOptionsCollection) HasSet() bool {
	return js.True == bindings.HasHTMLOptionsCollectionSet(
		this.Ref(),
	)
}

// SetFunc returns the method "HTMLOptionsCollection.".
func (this HTMLOptionsCollection) SetFunc() (fn js.Func[func(index uint32, option HTMLOptionElement)]) {
	return fn.FromRef(
		bindings.HTMLOptionsCollectionSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "HTMLOptionsCollection.".
func (this HTMLOptionsCollection) Set(index uint32, option HTMLOptionElement) (ret js.Void) {
	bindings.CallHTMLOptionsCollectionSet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		option.Ref(),
	)

	return
}

// TrySet calls the method "HTMLOptionsCollection."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOptionsCollection) TrySet(index uint32, option HTMLOptionElement) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOptionsCollectionSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		option.Ref(),
	)

	return
}

// HasAdd returns true if the method "HTMLOptionsCollection.add" exists.
func (this HTMLOptionsCollection) HasAdd() bool {
	return js.True == bindings.HasHTMLOptionsCollectionAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "HTMLOptionsCollection.add".
func (this HTMLOptionsCollection) AddFunc() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32)]) {
	return fn.FromRef(
		bindings.HTMLOptionsCollectionAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "HTMLOptionsCollection.add".
func (this HTMLOptionsCollection) Add(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32) (ret js.Void) {
	bindings.CallHTMLOptionsCollectionAdd(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
		before.Ref(),
	)

	return
}

// TryAdd calls the method "HTMLOptionsCollection.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOptionsCollection) TryAdd(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOptionsCollectionAdd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
		before.Ref(),
	)

	return
}

// HasAdd1 returns true if the method "HTMLOptionsCollection.add" exists.
func (this HTMLOptionsCollection) HasAdd1() bool {
	return js.True == bindings.HasHTMLOptionsCollectionAdd1(
		this.Ref(),
	)
}

// Add1Func returns the method "HTMLOptionsCollection.add".
func (this HTMLOptionsCollection) Add1Func() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement)]) {
	return fn.FromRef(
		bindings.HTMLOptionsCollectionAdd1Func(
			this.Ref(),
		),
	)
}

// Add1 calls the method "HTMLOptionsCollection.add".
func (this HTMLOptionsCollection) Add1(element OneOf_HTMLOptionElement_HTMLOptGroupElement) (ret js.Void) {
	bindings.CallHTMLOptionsCollectionAdd1(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryAdd1 calls the method "HTMLOptionsCollection.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOptionsCollection) TryAdd1(element OneOf_HTMLOptionElement_HTMLOptGroupElement) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOptionsCollectionAdd1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

// HasRemove returns true if the method "HTMLOptionsCollection.remove" exists.
func (this HTMLOptionsCollection) HasRemove() bool {
	return js.True == bindings.HasHTMLOptionsCollectionRemove(
		this.Ref(),
	)
}

// RemoveFunc returns the method "HTMLOptionsCollection.remove".
func (this HTMLOptionsCollection) RemoveFunc() (fn js.Func[func(index int32)]) {
	return fn.FromRef(
		bindings.HTMLOptionsCollectionRemoveFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "HTMLOptionsCollection.remove".
func (this HTMLOptionsCollection) Remove(index int32) (ret js.Void) {
	bindings.CallHTMLOptionsCollectionRemove(
		this.Ref(), js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryRemove calls the method "HTMLOptionsCollection.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOptionsCollection) TryRemove(index int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOptionsCollectionRemove(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

type OneOf_HTMLImageElement_SVGImageElement struct {
	ref js.Ref
}

func (x OneOf_HTMLImageElement_SVGImageElement) Ref() js.Ref {
	return x.ref
}

func (x OneOf_HTMLImageElement_SVGImageElement) Free() {
	x.ref.Free()
}

func (x OneOf_HTMLImageElement_SVGImageElement) FromRef(ref js.Ref) OneOf_HTMLImageElement_SVGImageElement {
	return OneOf_HTMLImageElement_SVGImageElement{
		ref: ref,
	}
}

func (x OneOf_HTMLImageElement_SVGImageElement) HTMLImageElement() HTMLImageElement {
	return HTMLImageElement{}.FromRef(x.ref)
}

func (x OneOf_HTMLImageElement_SVGImageElement) SVGImageElement() SVGImageElement {
	return SVGImageElement{}.FromRef(x.ref)
}

type HTMLOrSVGImageElement = OneOf_HTMLImageElement_SVGImageElement

func NewHTMLOutputElement() (ret HTMLOutputElement) {
	ret.ref = bindings.NewHTMLOutputElementByHTMLOutputElement()
	return
}

type HTMLOutputElement struct {
	HTMLElement
}

func (this HTMLOutputElement) Once() HTMLOutputElement {
	this.Ref().Once()
	return this
}

func (this HTMLOutputElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLOutputElement) FromRef(ref js.Ref) HTMLOutputElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLOutputElement) Free() {
	this.Ref().Free()
}

// HtmlFor returns the value of property "HTMLOutputElement.htmlFor".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) HtmlFor() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementHtmlFor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Form returns the value of property "HTMLOutputElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementForm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "HTMLOutputElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLOutputElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOutputElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLOutputElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLOutputElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DefaultValue returns the value of property "HTMLOutputElement.defaultValue".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) DefaultValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementDefaultValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDefaultValue sets the value of property "HTMLOutputElement.defaultValue" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOutputElement) SetDefaultValue(val js.String) bool {
	return js.True == bindings.SetHTMLOutputElementDefaultValue(
		this.Ref(),
		val.Ref(),
	)
}

// Value returns the value of property "HTMLOutputElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLOutputElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOutputElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLOutputElementValue(
		this.Ref(),
		val.Ref(),
	)
}

// WillValidate returns the value of property "HTMLOutputElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementWillValidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLOutputElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementValidity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLOutputElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementValidationMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLOutputElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementLabels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasCheckValidity returns true if the method "HTMLOutputElement.checkValidity" exists.
func (this HTMLOutputElement) HasCheckValidity() bool {
	return js.True == bindings.HasHTMLOutputElementCheckValidity(
		this.Ref(),
	)
}

// CheckValidityFunc returns the method "HTMLOutputElement.checkValidity".
func (this HTMLOutputElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLOutputElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLOutputElement.checkValidity".
func (this HTMLOutputElement) CheckValidity() (ret bool) {
	bindings.CallHTMLOutputElementCheckValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLOutputElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOutputElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOutputElementCheckValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReportValidity returns true if the method "HTMLOutputElement.reportValidity" exists.
func (this HTMLOutputElement) HasReportValidity() bool {
	return js.True == bindings.HasHTMLOutputElementReportValidity(
		this.Ref(),
	)
}

// ReportValidityFunc returns the method "HTMLOutputElement.reportValidity".
func (this HTMLOutputElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLOutputElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLOutputElement.reportValidity".
func (this HTMLOutputElement) ReportValidity() (ret bool) {
	bindings.CallHTMLOutputElementReportValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLOutputElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOutputElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOutputElementReportValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetCustomValidity returns true if the method "HTMLOutputElement.setCustomValidity" exists.
func (this HTMLOutputElement) HasSetCustomValidity() bool {
	return js.True == bindings.HasHTMLOutputElementSetCustomValidity(
		this.Ref(),
	)
}

// SetCustomValidityFunc returns the method "HTMLOutputElement.setCustomValidity".
func (this HTMLOutputElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLOutputElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLOutputElement.setCustomValidity".
func (this HTMLOutputElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLOutputElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLOutputElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOutputElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOutputElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

func NewHTMLParagraphElement() (ret HTMLParagraphElement) {
	ret.ref = bindings.NewHTMLParagraphElementByHTMLParagraphElement()
	return
}

type HTMLParagraphElement struct {
	HTMLElement
}

func (this HTMLParagraphElement) Once() HTMLParagraphElement {
	this.Ref().Once()
	return this
}

func (this HTMLParagraphElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLParagraphElement) FromRef(ref js.Ref) HTMLParagraphElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLParagraphElement) Free() {
	this.Ref().Free()
}

// Align returns the value of property "HTMLParagraphElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLParagraphElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLParagraphElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLParagraphElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParagraphElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLParagraphElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLParamElement() (ret HTMLParamElement) {
	ret.ref = bindings.NewHTMLParamElementByHTMLParamElement()
	return
}

type HTMLParamElement struct {
	HTMLElement
}

func (this HTMLParamElement) Once() HTMLParamElement {
	this.Ref().Once()
	return this
}

func (this HTMLParamElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLParamElement) FromRef(ref js.Ref) HTMLParamElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLParamElement) Free() {
	this.Ref().Free()
}

// Name returns the value of property "HTMLParamElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLParamElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLParamElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLParamElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParamElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLParamElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Value returns the value of property "HTMLParamElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLParamElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLParamElementValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLParamElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParamElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLParamElementValue(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLParamElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLParamElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLParamElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLParamElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParamElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLParamElementType(
		this.Ref(),
		val.Ref(),
	)
}

// ValueType returns the value of property "HTMLParamElement.valueType".
//
// It returns ok=false if there is no such property.
func (this HTMLParamElement) ValueType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLParamElementValueType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValueType sets the value of property "HTMLParamElement.valueType" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParamElement) SetValueType(val js.String) bool {
	return js.True == bindings.SetHTMLParamElementValueType(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLPictureElement() (ret HTMLPictureElement) {
	ret.ref = bindings.NewHTMLPictureElementByHTMLPictureElement()
	return
}

type HTMLPictureElement struct {
	HTMLElement
}

func (this HTMLPictureElement) Once() HTMLPictureElement {
	this.Ref().Once()
	return this
}

func (this HTMLPictureElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLPictureElement) FromRef(ref js.Ref) HTMLPictureElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLPictureElement) Free() {
	this.Ref().Free()
}

type PortalActivateOptions struct {
	// Data is "PortalActivateOptions.data"
	//
	// Optional
	Data js.Any
	// Transfer is "PortalActivateOptions.transfer"
	//
	// Optional, defaults to [].
	Transfer js.Array[js.Object]

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a PortalActivateOptions with all fields set.
func (p PortalActivateOptions) FromRef(ref js.Ref) PortalActivateOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new PortalActivateOptions in the application heap.
func (p PortalActivateOptions) New() js.Ref {
	return bindings.PortalActivateOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p PortalActivateOptions) UpdateFrom(ref js.Ref) {
	bindings.PortalActivateOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p PortalActivateOptions) Update(ref js.Ref) {
	bindings.PortalActivateOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

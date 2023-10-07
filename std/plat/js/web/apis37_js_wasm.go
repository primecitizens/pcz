// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type HTMLAudioElement struct {
	HTMLMediaElement
}

func (this HTMLAudioElement) Once() HTMLAudioElement {
	this.ref.Once()
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
	this.ref.Free()
}

type HTMLBRElement struct {
	HTMLElement
}

func (this HTMLBRElement) Once() HTMLBRElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Clear returns the value of property "HTMLBRElement.clear".
//
// It returns ok=false if there is no such property.
func (this HTMLBRElement) Clear() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBRElementClear(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetClear sets the value of property "HTMLBRElement.clear" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBRElement) SetClear(val js.String) bool {
	return js.True == bindings.SetHTMLBRElementClear(
		this.ref,
		val.Ref(),
	)
}

type HTMLBaseElement struct {
	HTMLElement
}

func (this HTMLBaseElement) Once() HTMLBaseElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Href returns the value of property "HTMLBaseElement.href".
//
// It returns ok=false if there is no such property.
func (this HTMLBaseElement) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBaseElementHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHref sets the value of property "HTMLBaseElement.href" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBaseElement) SetHref(val js.String) bool {
	return js.True == bindings.SetHTMLBaseElementHref(
		this.ref,
		val.Ref(),
	)
}

// Target returns the value of property "HTMLBaseElement.target".
//
// It returns ok=false if there is no such property.
func (this HTMLBaseElement) Target() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBaseElementTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTarget sets the value of property "HTMLBaseElement.target" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBaseElement) SetTarget(val js.String) bool {
	return js.True == bindings.SetHTMLBaseElementTarget(
		this.ref,
		val.Ref(),
	)
}

type HTMLBodyElement struct {
	HTMLElement
}

func (this HTMLBodyElement) Once() HTMLBodyElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Text returns the value of property "HTMLBodyElement.text".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "HTMLBodyElement.text" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetText(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementText(
		this.ref,
		val.Ref(),
	)
}

// Link returns the value of property "HTMLBodyElement.link".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) Link() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementLink(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLink sets the value of property "HTMLBodyElement.link" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetLink(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementLink(
		this.ref,
		val.Ref(),
	)
}

// VLink returns the value of property "HTMLBodyElement.vLink".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) VLink() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementVLink(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVLink sets the value of property "HTMLBodyElement.vLink" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetVLink(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementVLink(
		this.ref,
		val.Ref(),
	)
}

// ALink returns the value of property "HTMLBodyElement.aLink".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) ALink() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementALink(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetALink sets the value of property "HTMLBodyElement.aLink" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetALink(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementALink(
		this.ref,
		val.Ref(),
	)
}

// BgColor returns the value of property "HTMLBodyElement.bgColor".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementBgColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "HTMLBodyElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementBgColor(
		this.ref,
		val.Ref(),
	)
}

// Background returns the value of property "HTMLBodyElement.background".
//
// It returns ok=false if there is no such property.
func (this HTMLBodyElement) Background() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLBodyElementBackground(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBackground sets the value of property "HTMLBodyElement.background" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetBackground(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementBackground(
		this.ref,
		val.Ref(),
	)
}

type HTMLButtonElement struct {
	HTMLElement
}

func (this HTMLButtonElement) Once() HTMLButtonElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Disabled returns the value of property "HTMLButtonElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLButtonElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLButtonElementDisabled(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLButtonElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementForm(
		this.ref, js.Pointer(&ret),
	)
	return
}

// FormAction returns the value of property "HTMLButtonElement.formAction".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) FormAction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementFormAction(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFormAction sets the value of property "HTMLButtonElement.formAction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetFormAction(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementFormAction(
		this.ref,
		val.Ref(),
	)
}

// FormEnctype returns the value of property "HTMLButtonElement.formEnctype".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) FormEnctype() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementFormEnctype(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFormEnctype sets the value of property "HTMLButtonElement.formEnctype" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetFormEnctype(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementFormEnctype(
		this.ref,
		val.Ref(),
	)
}

// FormMethod returns the value of property "HTMLButtonElement.formMethod".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) FormMethod() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementFormMethod(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFormMethod sets the value of property "HTMLButtonElement.formMethod" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetFormMethod(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementFormMethod(
		this.ref,
		val.Ref(),
	)
}

// FormNoValidate returns the value of property "HTMLButtonElement.formNoValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) FormNoValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementFormNoValidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFormNoValidate sets the value of property "HTMLButtonElement.formNoValidate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetFormNoValidate(val bool) bool {
	return js.True == bindings.SetHTMLButtonElementFormNoValidate(
		this.ref,
		js.Bool(bool(val)),
	)
}

// FormTarget returns the value of property "HTMLButtonElement.formTarget".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) FormTarget() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementFormTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFormTarget sets the value of property "HTMLButtonElement.formTarget" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetFormTarget(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementFormTarget(
		this.ref,
		val.Ref(),
	)
}

// Name returns the value of property "HTMLButtonElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLButtonElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementName(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLButtonElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLButtonElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementType(
		this.ref,
		val.Ref(),
	)
}

// Value returns the value of property "HTMLButtonElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLButtonElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementValue(
		this.ref,
		val.Ref(),
	)
}

// WillValidate returns the value of property "HTMLButtonElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementWillValidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLButtonElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementValidity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLButtonElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementValidationMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLButtonElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementLabels(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PopoverTargetElement returns the value of property "HTMLButtonElement.popoverTargetElement".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) PopoverTargetElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementPopoverTargetElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPopoverTargetElement sets the value of property "HTMLButtonElement.popoverTargetElement" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetPopoverTargetElement(val Element) bool {
	return js.True == bindings.SetHTMLButtonElementPopoverTargetElement(
		this.ref,
		val.Ref(),
	)
}

// PopoverTargetAction returns the value of property "HTMLButtonElement.popoverTargetAction".
//
// It returns ok=false if there is no such property.
func (this HTMLButtonElement) PopoverTargetAction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLButtonElementPopoverTargetAction(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPopoverTargetAction sets the value of property "HTMLButtonElement.popoverTargetAction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetPopoverTargetAction(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementPopoverTargetAction(
		this.ref,
		val.Ref(),
	)
}

// HasFuncCheckValidity returns true if the method "HTMLButtonElement.checkValidity" exists.
func (this HTMLButtonElement) HasFuncCheckValidity() bool {
	return js.True == bindings.HasFuncHTMLButtonElementCheckValidity(
		this.ref,
	)
}

// FuncCheckValidity returns the method "HTMLButtonElement.checkValidity".
func (this HTMLButtonElement) FuncCheckValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLButtonElementCheckValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckValidity calls the method "HTMLButtonElement.checkValidity".
func (this HTMLButtonElement) CheckValidity() (ret bool) {
	bindings.CallHTMLButtonElementCheckValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLButtonElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLButtonElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLButtonElementCheckValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReportValidity returns true if the method "HTMLButtonElement.reportValidity" exists.
func (this HTMLButtonElement) HasFuncReportValidity() bool {
	return js.True == bindings.HasFuncHTMLButtonElementReportValidity(
		this.ref,
	)
}

// FuncReportValidity returns the method "HTMLButtonElement.reportValidity".
func (this HTMLButtonElement) FuncReportValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLButtonElementReportValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportValidity calls the method "HTMLButtonElement.reportValidity".
func (this HTMLButtonElement) ReportValidity() (ret bool) {
	bindings.CallHTMLButtonElementReportValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLButtonElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLButtonElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLButtonElementReportValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetCustomValidity returns true if the method "HTMLButtonElement.setCustomValidity" exists.
func (this HTMLButtonElement) HasFuncSetCustomValidity() bool {
	return js.True == bindings.HasFuncHTMLButtonElementSetCustomValidity(
		this.ref,
	)
}

// FuncSetCustomValidity returns the method "HTMLButtonElement.setCustomValidity".
func (this HTMLButtonElement) FuncSetCustomValidity() (fn js.Func[func(err js.String)]) {
	bindings.FuncHTMLButtonElementSetCustomValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetCustomValidity calls the method "HTMLButtonElement.setCustomValidity".
func (this HTMLButtonElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLButtonElementSetCustomValidity(
		this.ref, js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLButtonElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLButtonElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLButtonElementSetCustomValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

type HTMLDListElement struct {
	HTMLElement
}

func (this HTMLDListElement) Once() HTMLDListElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Compact returns the value of property "HTMLDListElement.compact".
//
// It returns ok=false if there is no such property.
func (this HTMLDListElement) Compact() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLDListElementCompact(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCompact sets the value of property "HTMLDListElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDListElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLDListElementCompact(
		this.ref,
		js.Bool(bool(val)),
	)
}

type HTMLDataElement struct {
	HTMLElement
}

func (this HTMLDataElement) Once() HTMLDataElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "HTMLDataElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLDataElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLDataElementValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLDataElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDataElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLDataElementValue(
		this.ref,
		val.Ref(),
	)
}

type HTMLDataListElement struct {
	HTMLElement
}

func (this HTMLDataListElement) Once() HTMLDataListElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Options returns the value of property "HTMLDataListElement.options".
//
// It returns ok=false if there is no such property.
func (this HTMLDataListElement) Options() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLDataListElementOptions(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HTMLDetailsElement struct {
	HTMLElement
}

func (this HTMLDetailsElement) Once() HTMLDetailsElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Open returns the value of property "HTMLDetailsElement.open".
//
// It returns ok=false if there is no such property.
func (this HTMLDetailsElement) Open() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLDetailsElementOpen(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOpen sets the value of property "HTMLDetailsElement.open" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDetailsElement) SetOpen(val bool) bool {
	return js.True == bindings.SetHTMLDetailsElementOpen(
		this.ref,
		js.Bool(bool(val)),
	)
}

type HTMLDialogElement struct {
	HTMLElement
}

func (this HTMLDialogElement) Once() HTMLDialogElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Open returns the value of property "HTMLDialogElement.open".
//
// It returns ok=false if there is no such property.
func (this HTMLDialogElement) Open() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLDialogElementOpen(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOpen sets the value of property "HTMLDialogElement.open" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDialogElement) SetOpen(val bool) bool {
	return js.True == bindings.SetHTMLDialogElementOpen(
		this.ref,
		js.Bool(bool(val)),
	)
}

// ReturnValue returns the value of property "HTMLDialogElement.returnValue".
//
// It returns ok=false if there is no such property.
func (this HTMLDialogElement) ReturnValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLDialogElementReturnValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReturnValue sets the value of property "HTMLDialogElement.returnValue" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDialogElement) SetReturnValue(val js.String) bool {
	return js.True == bindings.SetHTMLDialogElementReturnValue(
		this.ref,
		val.Ref(),
	)
}

// HasFuncShow returns true if the method "HTMLDialogElement.show" exists.
func (this HTMLDialogElement) HasFuncShow() bool {
	return js.True == bindings.HasFuncHTMLDialogElementShow(
		this.ref,
	)
}

// FuncShow returns the method "HTMLDialogElement.show".
func (this HTMLDialogElement) FuncShow() (fn js.Func[func()]) {
	bindings.FuncHTMLDialogElementShow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Show calls the method "HTMLDialogElement.show".
func (this HTMLDialogElement) Show() (ret js.Void) {
	bindings.CallHTMLDialogElementShow(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryShow calls the method "HTMLDialogElement.show"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLDialogElement) TryShow() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLDialogElementShow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncShowModal returns true if the method "HTMLDialogElement.showModal" exists.
func (this HTMLDialogElement) HasFuncShowModal() bool {
	return js.True == bindings.HasFuncHTMLDialogElementShowModal(
		this.ref,
	)
}

// FuncShowModal returns the method "HTMLDialogElement.showModal".
func (this HTMLDialogElement) FuncShowModal() (fn js.Func[func()]) {
	bindings.FuncHTMLDialogElementShowModal(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShowModal calls the method "HTMLDialogElement.showModal".
func (this HTMLDialogElement) ShowModal() (ret js.Void) {
	bindings.CallHTMLDialogElementShowModal(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryShowModal calls the method "HTMLDialogElement.showModal"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLDialogElement) TryShowModal() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLDialogElementShowModal(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncClose returns true if the method "HTMLDialogElement.close" exists.
func (this HTMLDialogElement) HasFuncClose() bool {
	return js.True == bindings.HasFuncHTMLDialogElementClose(
		this.ref,
	)
}

// FuncClose returns the method "HTMLDialogElement.close".
func (this HTMLDialogElement) FuncClose() (fn js.Func[func(returnValue js.String)]) {
	bindings.FuncHTMLDialogElementClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "HTMLDialogElement.close".
func (this HTMLDialogElement) Close(returnValue js.String) (ret js.Void) {
	bindings.CallHTMLDialogElementClose(
		this.ref, js.Pointer(&ret),
		returnValue.Ref(),
	)

	return
}

// TryClose calls the method "HTMLDialogElement.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLDialogElement) TryClose(returnValue js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLDialogElementClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		returnValue.Ref(),
	)

	return
}

// HasFuncClose1 returns true if the method "HTMLDialogElement.close" exists.
func (this HTMLDialogElement) HasFuncClose1() bool {
	return js.True == bindings.HasFuncHTMLDialogElementClose1(
		this.ref,
	)
}

// FuncClose1 returns the method "HTMLDialogElement.close".
func (this HTMLDialogElement) FuncClose1() (fn js.Func[func()]) {
	bindings.FuncHTMLDialogElementClose1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close1 calls the method "HTMLDialogElement.close".
func (this HTMLDialogElement) Close1() (ret js.Void) {
	bindings.CallHTMLDialogElementClose1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose1 calls the method "HTMLDialogElement.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLDialogElement) TryClose1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLDialogElementClose1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type HTMLDirectoryElement struct {
	HTMLElement
}

func (this HTMLDirectoryElement) Once() HTMLDirectoryElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Compact returns the value of property "HTMLDirectoryElement.compact".
//
// It returns ok=false if there is no such property.
func (this HTMLDirectoryElement) Compact() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLDirectoryElementCompact(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCompact sets the value of property "HTMLDirectoryElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDirectoryElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLDirectoryElementCompact(
		this.ref,
		js.Bool(bool(val)),
	)
}

type HTMLDivElement struct {
	HTMLElement
}

func (this HTMLDivElement) Once() HTMLDivElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Align returns the value of property "HTMLDivElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLDivElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLDivElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLDivElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDivElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLDivElementAlign(
		this.ref,
		val.Ref(),
	)
}

type HTMLEmbedElement struct {
	HTMLElement
}

func (this HTMLEmbedElement) Once() HTMLEmbedElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Src returns the value of property "HTMLEmbedElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLEmbedElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementSrc(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLEmbedElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLEmbedElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementType(
		this.ref,
		val.Ref(),
	)
}

// Width returns the value of property "HTMLEmbedElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLEmbedElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementWidth(
		this.ref,
		val.Ref(),
	)
}

// Height returns the value of property "HTMLEmbedElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLEmbedElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementHeight(
		this.ref,
		val.Ref(),
	)
}

// Align returns the value of property "HTMLEmbedElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLEmbedElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementAlign(
		this.ref,
		val.Ref(),
	)
}

// Name returns the value of property "HTMLEmbedElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLEmbedElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLEmbedElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLEmbedElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementName(
		this.ref,
		val.Ref(),
	)
}

// HasFuncGetSVGDocument returns true if the method "HTMLEmbedElement.getSVGDocument" exists.
func (this HTMLEmbedElement) HasFuncGetSVGDocument() bool {
	return js.True == bindings.HasFuncHTMLEmbedElementGetSVGDocument(
		this.ref,
	)
}

// FuncGetSVGDocument returns the method "HTMLEmbedElement.getSVGDocument".
func (this HTMLEmbedElement) FuncGetSVGDocument() (fn js.Func[func() Document]) {
	bindings.FuncHTMLEmbedElementGetSVGDocument(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSVGDocument calls the method "HTMLEmbedElement.getSVGDocument".
func (this HTMLEmbedElement) GetSVGDocument() (ret Document) {
	bindings.CallHTMLEmbedElementGetSVGDocument(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSVGDocument calls the method "HTMLEmbedElement.getSVGDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLEmbedElement) TryGetSVGDocument() (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLEmbedElementGetSVGDocument(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type HTMLFencedFrameElement struct {
	HTMLElement
}

func (this HTMLFencedFrameElement) Once() HTMLFencedFrameElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Config returns the value of property "HTMLFencedFrameElement.config".
//
// It returns ok=false if there is no such property.
func (this HTMLFencedFrameElement) Config() (ret FencedFrameConfig, ok bool) {
	ok = js.True == bindings.GetHTMLFencedFrameElementConfig(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetConfig sets the value of property "HTMLFencedFrameElement.config" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFencedFrameElement) SetConfig(val FencedFrameConfig) bool {
	return js.True == bindings.SetHTMLFencedFrameElementConfig(
		this.ref,
		val.Ref(),
	)
}

// Width returns the value of property "HTMLFencedFrameElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLFencedFrameElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFencedFrameElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLFencedFrameElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFencedFrameElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLFencedFrameElementWidth(
		this.ref,
		val.Ref(),
	)
}

// Height returns the value of property "HTMLFencedFrameElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLFencedFrameElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFencedFrameElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLFencedFrameElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFencedFrameElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLFencedFrameElementHeight(
		this.ref,
		val.Ref(),
	)
}

// Allow returns the value of property "HTMLFencedFrameElement.allow".
//
// It returns ok=false if there is no such property.
func (this HTMLFencedFrameElement) Allow() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFencedFrameElementAllow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAllow sets the value of property "HTMLFencedFrameElement.allow" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFencedFrameElement) SetAllow(val js.String) bool {
	return js.True == bindings.SetHTMLFencedFrameElementAllow(
		this.ref,
		val.Ref(),
	)
}

type HTMLFieldSetElement struct {
	HTMLElement
}

func (this HTMLFieldSetElement) Once() HTMLFieldSetElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Disabled returns the value of property "HTMLFieldSetElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLFieldSetElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFieldSetElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLFieldSetElementDisabled(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLFieldSetElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementForm(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "HTMLFieldSetElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLFieldSetElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFieldSetElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLFieldSetElementName(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLFieldSetElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Elements returns the value of property "HTMLFieldSetElement.elements".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Elements() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementElements(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WillValidate returns the value of property "HTMLFieldSetElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementWillValidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLFieldSetElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementValidity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLFieldSetElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLFieldSetElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFieldSetElementValidationMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncCheckValidity returns true if the method "HTMLFieldSetElement.checkValidity" exists.
func (this HTMLFieldSetElement) HasFuncCheckValidity() bool {
	return js.True == bindings.HasFuncHTMLFieldSetElementCheckValidity(
		this.ref,
	)
}

// FuncCheckValidity returns the method "HTMLFieldSetElement.checkValidity".
func (this HTMLFieldSetElement) FuncCheckValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLFieldSetElementCheckValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckValidity calls the method "HTMLFieldSetElement.checkValidity".
func (this HTMLFieldSetElement) CheckValidity() (ret bool) {
	bindings.CallHTMLFieldSetElementCheckValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLFieldSetElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFieldSetElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFieldSetElementCheckValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReportValidity returns true if the method "HTMLFieldSetElement.reportValidity" exists.
func (this HTMLFieldSetElement) HasFuncReportValidity() bool {
	return js.True == bindings.HasFuncHTMLFieldSetElementReportValidity(
		this.ref,
	)
}

// FuncReportValidity returns the method "HTMLFieldSetElement.reportValidity".
func (this HTMLFieldSetElement) FuncReportValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLFieldSetElementReportValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportValidity calls the method "HTMLFieldSetElement.reportValidity".
func (this HTMLFieldSetElement) ReportValidity() (ret bool) {
	bindings.CallHTMLFieldSetElementReportValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLFieldSetElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFieldSetElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFieldSetElementReportValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetCustomValidity returns true if the method "HTMLFieldSetElement.setCustomValidity" exists.
func (this HTMLFieldSetElement) HasFuncSetCustomValidity() bool {
	return js.True == bindings.HasFuncHTMLFieldSetElementSetCustomValidity(
		this.ref,
	)
}

// FuncSetCustomValidity returns the method "HTMLFieldSetElement.setCustomValidity".
func (this HTMLFieldSetElement) FuncSetCustomValidity() (fn js.Func[func(err js.String)]) {
	bindings.FuncHTMLFieldSetElementSetCustomValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetCustomValidity calls the method "HTMLFieldSetElement.setCustomValidity".
func (this HTMLFieldSetElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLFieldSetElementSetCustomValidity(
		this.ref, js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLFieldSetElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLFieldSetElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLFieldSetElementSetCustomValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

type HTMLFontElement struct {
	HTMLElement
}

func (this HTMLFontElement) Once() HTMLFontElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Color returns the value of property "HTMLFontElement.color".
//
// It returns ok=false if there is no such property.
func (this HTMLFontElement) Color() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFontElementColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetColor sets the value of property "HTMLFontElement.color" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFontElement) SetColor(val js.String) bool {
	return js.True == bindings.SetHTMLFontElementColor(
		this.ref,
		val.Ref(),
	)
}

// Face returns the value of property "HTMLFontElement.face".
//
// It returns ok=false if there is no such property.
func (this HTMLFontElement) Face() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFontElementFace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFace sets the value of property "HTMLFontElement.face" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFontElement) SetFace(val js.String) bool {
	return js.True == bindings.SetHTMLFontElementFace(
		this.ref,
		val.Ref(),
	)
}

// Size returns the value of property "HTMLFontElement.size".
//
// It returns ok=false if there is no such property.
func (this HTMLFontElement) Size() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFontElementSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSize sets the value of property "HTMLFontElement.size" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFontElement) SetSize(val js.String) bool {
	return js.True == bindings.SetHTMLFontElementSize(
		this.ref,
		val.Ref(),
	)
}

type HTMLFrameElement struct {
	HTMLElement
}

func (this HTMLFrameElement) Once() HTMLFrameElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "HTMLFrameElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLFrameElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementName(
		this.ref,
		val.Ref(),
	)
}

// Scrolling returns the value of property "HTMLFrameElement.scrolling".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) Scrolling() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementScrolling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrolling sets the value of property "HTMLFrameElement.scrolling" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetScrolling(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementScrolling(
		this.ref,
		val.Ref(),
	)
}

// Src returns the value of property "HTMLFrameElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLFrameElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementSrc(
		this.ref,
		val.Ref(),
	)
}

// FrameBorder returns the value of property "HTMLFrameElement.frameBorder".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) FrameBorder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementFrameBorder(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFrameBorder sets the value of property "HTMLFrameElement.frameBorder" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetFrameBorder(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementFrameBorder(
		this.ref,
		val.Ref(),
	)
}

// LongDesc returns the value of property "HTMLFrameElement.longDesc".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) LongDesc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementLongDesc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLongDesc sets the value of property "HTMLFrameElement.longDesc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetLongDesc(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementLongDesc(
		this.ref,
		val.Ref(),
	)
}

// NoResize returns the value of property "HTMLFrameElement.noResize".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) NoResize() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementNoResize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNoResize sets the value of property "HTMLFrameElement.noResize" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetNoResize(val bool) bool {
	return js.True == bindings.SetHTMLFrameElementNoResize(
		this.ref,
		js.Bool(bool(val)),
	)
}

// ContentDocument returns the value of property "HTMLFrameElement.contentDocument".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) ContentDocument() (ret Document, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementContentDocument(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ContentWindow returns the value of property "HTMLFrameElement.contentWindow".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) ContentWindow() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementContentWindow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MarginHeight returns the value of property "HTMLFrameElement.marginHeight".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) MarginHeight() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementMarginHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMarginHeight sets the value of property "HTMLFrameElement.marginHeight" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetMarginHeight(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementMarginHeight(
		this.ref,
		val.Ref(),
	)
}

// MarginWidth returns the value of property "HTMLFrameElement.marginWidth".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameElement) MarginWidth() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameElementMarginWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMarginWidth sets the value of property "HTMLFrameElement.marginWidth" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetMarginWidth(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementMarginWidth(
		this.ref,
		val.Ref(),
	)
}

type HTMLFrameSetElement struct {
	HTMLElement
}

func (this HTMLFrameSetElement) Once() HTMLFrameSetElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Cols returns the value of property "HTMLFrameSetElement.cols".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameSetElement) Cols() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameSetElementCols(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCols sets the value of property "HTMLFrameSetElement.cols" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameSetElement) SetCols(val js.String) bool {
	return js.True == bindings.SetHTMLFrameSetElementCols(
		this.ref,
		val.Ref(),
	)
}

// Rows returns the value of property "HTMLFrameSetElement.rows".
//
// It returns ok=false if there is no such property.
func (this HTMLFrameSetElement) Rows() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLFrameSetElementRows(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRows sets the value of property "HTMLFrameSetElement.rows" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameSetElement) SetRows(val js.String) bool {
	return js.True == bindings.SetHTMLFrameSetElementRows(
		this.ref,
		val.Ref(),
	)
}

type HTMLHRElement struct {
	HTMLElement
}

func (this HTMLHRElement) Once() HTMLHRElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Align returns the value of property "HTMLHRElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLHRElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHRElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLHRElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHRElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLHRElementAlign(
		this.ref,
		val.Ref(),
	)
}

// Color returns the value of property "HTMLHRElement.color".
//
// It returns ok=false if there is no such property.
func (this HTMLHRElement) Color() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHRElementColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetColor sets the value of property "HTMLHRElement.color" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHRElement) SetColor(val js.String) bool {
	return js.True == bindings.SetHTMLHRElementColor(
		this.ref,
		val.Ref(),
	)
}

// NoShade returns the value of property "HTMLHRElement.noShade".
//
// It returns ok=false if there is no such property.
func (this HTMLHRElement) NoShade() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLHRElementNoShade(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNoShade sets the value of property "HTMLHRElement.noShade" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHRElement) SetNoShade(val bool) bool {
	return js.True == bindings.SetHTMLHRElementNoShade(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Size returns the value of property "HTMLHRElement.size".
//
// It returns ok=false if there is no such property.
func (this HTMLHRElement) Size() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHRElementSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSize sets the value of property "HTMLHRElement.size" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHRElement) SetSize(val js.String) bool {
	return js.True == bindings.SetHTMLHRElementSize(
		this.ref,
		val.Ref(),
	)
}

// Width returns the value of property "HTMLHRElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLHRElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHRElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLHRElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHRElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLHRElementWidth(
		this.ref,
		val.Ref(),
	)
}

type HTMLHeadingElement struct {
	HTMLElement
}

func (this HTMLHeadingElement) Once() HTMLHeadingElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Align returns the value of property "HTMLHeadingElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLHeadingElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHeadingElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLHeadingElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHeadingElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLHeadingElementAlign(
		this.ref,
		val.Ref(),
	)
}

type HTMLHtmlElement struct {
	HTMLElement
}

func (this HTMLHtmlElement) Once() HTMLHtmlElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Version returns the value of property "HTMLHtmlElement.version".
//
// It returns ok=false if there is no such property.
func (this HTMLHtmlElement) Version() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLHtmlElementVersion(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVersion sets the value of property "HTMLHtmlElement.version" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHtmlElement) SetVersion(val js.String) bool {
	return js.True == bindings.SetHTMLHtmlElementVersion(
		this.ref,
		val.Ref(),
	)
}

type HTMLIFrameElement struct {
	HTMLElement
}

func (this HTMLIFrameElement) Once() HTMLIFrameElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Src returns the value of property "HTMLIFrameElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLIFrameElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementSrc(
		this.ref,
		val.Ref(),
	)
}

// Srcdoc returns the value of property "HTMLIFrameElement.srcdoc".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Srcdoc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementSrcdoc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrcdoc sets the value of property "HTMLIFrameElement.srcdoc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetSrcdoc(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementSrcdoc(
		this.ref,
		val.Ref(),
	)
}

// Name returns the value of property "HTMLIFrameElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLIFrameElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementName(
		this.ref,
		val.Ref(),
	)
}

// Sandbox returns the value of property "HTMLIFrameElement.sandbox".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Sandbox() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementSandbox(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Allow returns the value of property "HTMLIFrameElement.allow".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Allow() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementAllow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAllow sets the value of property "HTMLIFrameElement.allow" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetAllow(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementAllow(
		this.ref,
		val.Ref(),
	)
}

// AllowFullscreen returns the value of property "HTMLIFrameElement.allowFullscreen".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) AllowFullscreen() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementAllowFullscreen(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAllowFullscreen sets the value of property "HTMLIFrameElement.allowFullscreen" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetAllowFullscreen(val bool) bool {
	return js.True == bindings.SetHTMLIFrameElementAllowFullscreen(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Width returns the value of property "HTMLIFrameElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLIFrameElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementWidth(
		this.ref,
		val.Ref(),
	)
}

// Height returns the value of property "HTMLIFrameElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLIFrameElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementHeight(
		this.ref,
		val.Ref(),
	)
}

// ReferrerPolicy returns the value of property "HTMLIFrameElement.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementReferrerPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLIFrameElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementReferrerPolicy(
		this.ref,
		val.Ref(),
	)
}

// Loading returns the value of property "HTMLIFrameElement.loading".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Loading() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementLoading(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLoading sets the value of property "HTMLIFrameElement.loading" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetLoading(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementLoading(
		this.ref,
		val.Ref(),
	)
}

// ContentDocument returns the value of property "HTMLIFrameElement.contentDocument".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) ContentDocument() (ret Document, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementContentDocument(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ContentWindow returns the value of property "HTMLIFrameElement.contentWindow".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) ContentWindow() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementContentWindow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// PermissionsPolicy returns the value of property "HTMLIFrameElement.permissionsPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) PermissionsPolicy() (ret PermissionsPolicy, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementPermissionsPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Csp returns the value of property "HTMLIFrameElement.csp".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Csp() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementCsp(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCsp sets the value of property "HTMLIFrameElement.csp" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetCsp(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementCsp(
		this.ref,
		val.Ref(),
	)
}

// PrivateToken returns the value of property "HTMLIFrameElement.privateToken".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) PrivateToken() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementPrivateToken(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPrivateToken sets the value of property "HTMLIFrameElement.privateToken" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetPrivateToken(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementPrivateToken(
		this.ref,
		val.Ref(),
	)
}

// Align returns the value of property "HTMLIFrameElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLIFrameElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementAlign(
		this.ref,
		val.Ref(),
	)
}

// Scrolling returns the value of property "HTMLIFrameElement.scrolling".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) Scrolling() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementScrolling(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrolling sets the value of property "HTMLIFrameElement.scrolling" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetScrolling(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementScrolling(
		this.ref,
		val.Ref(),
	)
}

// FrameBorder returns the value of property "HTMLIFrameElement.frameBorder".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) FrameBorder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementFrameBorder(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFrameBorder sets the value of property "HTMLIFrameElement.frameBorder" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetFrameBorder(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementFrameBorder(
		this.ref,
		val.Ref(),
	)
}

// LongDesc returns the value of property "HTMLIFrameElement.longDesc".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) LongDesc() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementLongDesc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLongDesc sets the value of property "HTMLIFrameElement.longDesc" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetLongDesc(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementLongDesc(
		this.ref,
		val.Ref(),
	)
}

// MarginHeight returns the value of property "HTMLIFrameElement.marginHeight".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) MarginHeight() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementMarginHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMarginHeight sets the value of property "HTMLIFrameElement.marginHeight" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetMarginHeight(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementMarginHeight(
		this.ref,
		val.Ref(),
	)
}

// MarginWidth returns the value of property "HTMLIFrameElement.marginWidth".
//
// It returns ok=false if there is no such property.
func (this HTMLIFrameElement) MarginWidth() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLIFrameElementMarginWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMarginWidth sets the value of property "HTMLIFrameElement.marginWidth" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetMarginWidth(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementMarginWidth(
		this.ref,
		val.Ref(),
	)
}

// HasFuncGetSVGDocument returns true if the method "HTMLIFrameElement.getSVGDocument" exists.
func (this HTMLIFrameElement) HasFuncGetSVGDocument() bool {
	return js.True == bindings.HasFuncHTMLIFrameElementGetSVGDocument(
		this.ref,
	)
}

// FuncGetSVGDocument returns the method "HTMLIFrameElement.getSVGDocument".
func (this HTMLIFrameElement) FuncGetSVGDocument() (fn js.Func[func() Document]) {
	bindings.FuncHTMLIFrameElementGetSVGDocument(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSVGDocument calls the method "HTMLIFrameElement.getSVGDocument".
func (this HTMLIFrameElement) GetSVGDocument() (ret Document) {
	bindings.CallHTMLIFrameElementGetSVGDocument(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSVGDocument calls the method "HTMLIFrameElement.getSVGDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLIFrameElement) TryGetSVGDocument() (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLIFrameElementGetSVGDocument(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

type HTMLInputElement struct {
	HTMLElement
}

func (this HTMLInputElement) Once() HTMLInputElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Accept returns the value of property "HTMLInputElement.accept".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Accept() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementAccept(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAccept sets the value of property "HTMLInputElement.accept" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetAccept(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementAccept(
		this.ref,
		val.Ref(),
	)
}

// Alt returns the value of property "HTMLInputElement.alt".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Alt() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementAlt(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlt sets the value of property "HTMLInputElement.alt" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetAlt(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementAlt(
		this.ref,
		val.Ref(),
	)
}

// Autocomplete returns the value of property "HTMLInputElement.autocomplete".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Autocomplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementAutocomplete(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAutocomplete sets the value of property "HTMLInputElement.autocomplete" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetAutocomplete(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementAutocomplete(
		this.ref,
		val.Ref(),
	)
}

// DefaultChecked returns the value of property "HTMLInputElement.defaultChecked".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) DefaultChecked() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementDefaultChecked(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDefaultChecked sets the value of property "HTMLInputElement.defaultChecked" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetDefaultChecked(val bool) bool {
	return js.True == bindings.SetHTMLInputElementDefaultChecked(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Checked returns the value of property "HTMLInputElement.checked".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Checked() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementChecked(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChecked sets the value of property "HTMLInputElement.checked" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetChecked(val bool) bool {
	return js.True == bindings.SetHTMLInputElementChecked(
		this.ref,
		js.Bool(bool(val)),
	)
}

// DirName returns the value of property "HTMLInputElement.dirName".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) DirName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementDirName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDirName sets the value of property "HTMLInputElement.dirName" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetDirName(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementDirName(
		this.ref,
		val.Ref(),
	)
}

// Disabled returns the value of property "HTMLInputElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLInputElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLInputElementDisabled(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLInputElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementForm(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Files returns the value of property "HTMLInputElement.files".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Files() (ret FileList, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFiles(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFiles sets the value of property "HTMLInputElement.files" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFiles(val FileList) bool {
	return js.True == bindings.SetHTMLInputElementFiles(
		this.ref,
		val.Ref(),
	)
}

// FormAction returns the value of property "HTMLInputElement.formAction".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) FormAction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFormAction(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFormAction sets the value of property "HTMLInputElement.formAction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFormAction(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementFormAction(
		this.ref,
		val.Ref(),
	)
}

// FormEnctype returns the value of property "HTMLInputElement.formEnctype".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) FormEnctype() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFormEnctype(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFormEnctype sets the value of property "HTMLInputElement.formEnctype" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFormEnctype(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementFormEnctype(
		this.ref,
		val.Ref(),
	)
}

// FormMethod returns the value of property "HTMLInputElement.formMethod".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) FormMethod() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFormMethod(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFormMethod sets the value of property "HTMLInputElement.formMethod" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFormMethod(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementFormMethod(
		this.ref,
		val.Ref(),
	)
}

// FormNoValidate returns the value of property "HTMLInputElement.formNoValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) FormNoValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFormNoValidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFormNoValidate sets the value of property "HTMLInputElement.formNoValidate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFormNoValidate(val bool) bool {
	return js.True == bindings.SetHTMLInputElementFormNoValidate(
		this.ref,
		js.Bool(bool(val)),
	)
}

// FormTarget returns the value of property "HTMLInputElement.formTarget".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) FormTarget() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementFormTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFormTarget sets the value of property "HTMLInputElement.formTarget" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetFormTarget(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementFormTarget(
		this.ref,
		val.Ref(),
	)
}

// Height returns the value of property "HTMLInputElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLInputElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLInputElementHeight(
		this.ref,
		uint32(val),
	)
}

// Indeterminate returns the value of property "HTMLInputElement.indeterminate".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Indeterminate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementIndeterminate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIndeterminate sets the value of property "HTMLInputElement.indeterminate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetIndeterminate(val bool) bool {
	return js.True == bindings.SetHTMLInputElementIndeterminate(
		this.ref,
		js.Bool(bool(val)),
	)
}

// List returns the value of property "HTMLInputElement.list".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) List() (ret HTMLDataListElement, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementList(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Max returns the value of property "HTMLInputElement.max".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Max() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementMax(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMax sets the value of property "HTMLInputElement.max" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetMax(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementMax(
		this.ref,
		val.Ref(),
	)
}

// MaxLength returns the value of property "HTMLInputElement.maxLength".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) MaxLength() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementMaxLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMaxLength sets the value of property "HTMLInputElement.maxLength" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetMaxLength(val int32) bool {
	return js.True == bindings.SetHTMLInputElementMaxLength(
		this.ref,
		int32(val),
	)
}

// Min returns the value of property "HTMLInputElement.min".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Min() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementMin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMin sets the value of property "HTMLInputElement.min" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetMin(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementMin(
		this.ref,
		val.Ref(),
	)
}

// MinLength returns the value of property "HTMLInputElement.minLength".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) MinLength() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementMinLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMinLength sets the value of property "HTMLInputElement.minLength" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetMinLength(val int32) bool {
	return js.True == bindings.SetHTMLInputElementMinLength(
		this.ref,
		int32(val),
	)
}

// Multiple returns the value of property "HTMLInputElement.multiple".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Multiple() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementMultiple(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMultiple sets the value of property "HTMLInputElement.multiple" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetMultiple(val bool) bool {
	return js.True == bindings.SetHTMLInputElementMultiple(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Name returns the value of property "HTMLInputElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLInputElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementName(
		this.ref,
		val.Ref(),
	)
}

// Pattern returns the value of property "HTMLInputElement.pattern".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Pattern() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementPattern(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPattern sets the value of property "HTMLInputElement.pattern" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetPattern(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementPattern(
		this.ref,
		val.Ref(),
	)
}

// Placeholder returns the value of property "HTMLInputElement.placeholder".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Placeholder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementPlaceholder(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPlaceholder sets the value of property "HTMLInputElement.placeholder" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetPlaceholder(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementPlaceholder(
		this.ref,
		val.Ref(),
	)
}

// ReadOnly returns the value of property "HTMLInputElement.readOnly".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) ReadOnly() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementReadOnly(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReadOnly sets the value of property "HTMLInputElement.readOnly" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetReadOnly(val bool) bool {
	return js.True == bindings.SetHTMLInputElementReadOnly(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Required returns the value of property "HTMLInputElement.required".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Required() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementRequired(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRequired sets the value of property "HTMLInputElement.required" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetRequired(val bool) bool {
	return js.True == bindings.SetHTMLInputElementRequired(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Size returns the value of property "HTMLInputElement.size".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Size() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSize sets the value of property "HTMLInputElement.size" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetSize(val uint32) bool {
	return js.True == bindings.SetHTMLInputElementSize(
		this.ref,
		uint32(val),
	)
}

// Src returns the value of property "HTMLInputElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLInputElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementSrc(
		this.ref,
		val.Ref(),
	)
}

// Step returns the value of property "HTMLInputElement.step".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Step() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementStep(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetStep sets the value of property "HTMLInputElement.step" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetStep(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementStep(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLInputElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLInputElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementType(
		this.ref,
		val.Ref(),
	)
}

// DefaultValue returns the value of property "HTMLInputElement.defaultValue".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) DefaultValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementDefaultValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDefaultValue sets the value of property "HTMLInputElement.defaultValue" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetDefaultValue(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementDefaultValue(
		this.ref,
		val.Ref(),
	)
}

// Value returns the value of property "HTMLInputElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLInputElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementValue(
		this.ref,
		val.Ref(),
	)
}

// ValueAsDate returns the value of property "HTMLInputElement.valueAsDate".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) ValueAsDate() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementValueAsDate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValueAsDate sets the value of property "HTMLInputElement.valueAsDate" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetValueAsDate(val js.Object) bool {
	return js.True == bindings.SetHTMLInputElementValueAsDate(
		this.ref,
		val.Ref(),
	)
}

// ValueAsNumber returns the value of property "HTMLInputElement.valueAsNumber".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) ValueAsNumber() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementValueAsNumber(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValueAsNumber sets the value of property "HTMLInputElement.valueAsNumber" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetValueAsNumber(val float64) bool {
	return js.True == bindings.SetHTMLInputElementValueAsNumber(
		this.ref,
		float64(val),
	)
}

// Width returns the value of property "HTMLInputElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLInputElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetWidth(val uint32) bool {
	return js.True == bindings.SetHTMLInputElementWidth(
		this.ref,
		uint32(val),
	)
}

// WillValidate returns the value of property "HTMLInputElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementWillValidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLInputElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementValidity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLInputElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementValidationMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLInputElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementLabels(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SelectionStart returns the value of property "HTMLInputElement.selectionStart".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) SelectionStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementSelectionStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionStart sets the value of property "HTMLInputElement.selectionStart" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetSelectionStart(val uint32) bool {
	return js.True == bindings.SetHTMLInputElementSelectionStart(
		this.ref,
		uint32(val),
	)
}

// SelectionEnd returns the value of property "HTMLInputElement.selectionEnd".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) SelectionEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementSelectionEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionEnd sets the value of property "HTMLInputElement.selectionEnd" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetSelectionEnd(val uint32) bool {
	return js.True == bindings.SetHTMLInputElementSelectionEnd(
		this.ref,
		uint32(val),
	)
}

// SelectionDirection returns the value of property "HTMLInputElement.selectionDirection".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) SelectionDirection() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementSelectionDirection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionDirection sets the value of property "HTMLInputElement.selectionDirection" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetSelectionDirection(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementSelectionDirection(
		this.ref,
		val.Ref(),
	)
}

// Align returns the value of property "HTMLInputElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLInputElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementAlign(
		this.ref,
		val.Ref(),
	)
}

// UseMap returns the value of property "HTMLInputElement.useMap".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) UseMap() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementUseMap(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUseMap sets the value of property "HTMLInputElement.useMap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetUseMap(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementUseMap(
		this.ref,
		val.Ref(),
	)
}

// Webkitdirectory returns the value of property "HTMLInputElement.webkitdirectory".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Webkitdirectory() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementWebkitdirectory(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWebkitdirectory sets the value of property "HTMLInputElement.webkitdirectory" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetWebkitdirectory(val bool) bool {
	return js.True == bindings.SetHTMLInputElementWebkitdirectory(
		this.ref,
		js.Bool(bool(val)),
	)
}

// WebkitEntries returns the value of property "HTMLInputElement.webkitEntries".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) WebkitEntries() (ret js.FrozenArray[FileSystemEntry], ok bool) {
	ok = js.True == bindings.GetHTMLInputElementWebkitEntries(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Capture returns the value of property "HTMLInputElement.capture".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) Capture() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementCapture(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCapture sets the value of property "HTMLInputElement.capture" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetCapture(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementCapture(
		this.ref,
		val.Ref(),
	)
}

// PopoverTargetElement returns the value of property "HTMLInputElement.popoverTargetElement".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) PopoverTargetElement() (ret Element, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementPopoverTargetElement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPopoverTargetElement sets the value of property "HTMLInputElement.popoverTargetElement" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetPopoverTargetElement(val Element) bool {
	return js.True == bindings.SetHTMLInputElementPopoverTargetElement(
		this.ref,
		val.Ref(),
	)
}

// PopoverTargetAction returns the value of property "HTMLInputElement.popoverTargetAction".
//
// It returns ok=false if there is no such property.
func (this HTMLInputElement) PopoverTargetAction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLInputElementPopoverTargetAction(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPopoverTargetAction sets the value of property "HTMLInputElement.popoverTargetAction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetPopoverTargetAction(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementPopoverTargetAction(
		this.ref,
		val.Ref(),
	)
}

// HasFuncStepUp returns true if the method "HTMLInputElement.stepUp" exists.
func (this HTMLInputElement) HasFuncStepUp() bool {
	return js.True == bindings.HasFuncHTMLInputElementStepUp(
		this.ref,
	)
}

// FuncStepUp returns the method "HTMLInputElement.stepUp".
func (this HTMLInputElement) FuncStepUp() (fn js.Func[func(n int32)]) {
	bindings.FuncHTMLInputElementStepUp(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StepUp calls the method "HTMLInputElement.stepUp".
func (this HTMLInputElement) StepUp(n int32) (ret js.Void) {
	bindings.CallHTMLInputElementStepUp(
		this.ref, js.Pointer(&ret),
		int32(n),
	)

	return
}

// TryStepUp calls the method "HTMLInputElement.stepUp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryStepUp(n int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementStepUp(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(n),
	)

	return
}

// HasFuncStepUp1 returns true if the method "HTMLInputElement.stepUp" exists.
func (this HTMLInputElement) HasFuncStepUp1() bool {
	return js.True == bindings.HasFuncHTMLInputElementStepUp1(
		this.ref,
	)
}

// FuncStepUp1 returns the method "HTMLInputElement.stepUp".
func (this HTMLInputElement) FuncStepUp1() (fn js.Func[func()]) {
	bindings.FuncHTMLInputElementStepUp1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StepUp1 calls the method "HTMLInputElement.stepUp".
func (this HTMLInputElement) StepUp1() (ret js.Void) {
	bindings.CallHTMLInputElementStepUp1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStepUp1 calls the method "HTMLInputElement.stepUp"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryStepUp1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementStepUp1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStepDown returns true if the method "HTMLInputElement.stepDown" exists.
func (this HTMLInputElement) HasFuncStepDown() bool {
	return js.True == bindings.HasFuncHTMLInputElementStepDown(
		this.ref,
	)
}

// FuncStepDown returns the method "HTMLInputElement.stepDown".
func (this HTMLInputElement) FuncStepDown() (fn js.Func[func(n int32)]) {
	bindings.FuncHTMLInputElementStepDown(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StepDown calls the method "HTMLInputElement.stepDown".
func (this HTMLInputElement) StepDown(n int32) (ret js.Void) {
	bindings.CallHTMLInputElementStepDown(
		this.ref, js.Pointer(&ret),
		int32(n),
	)

	return
}

// TryStepDown calls the method "HTMLInputElement.stepDown"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryStepDown(n int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementStepDown(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(n),
	)

	return
}

// HasFuncStepDown1 returns true if the method "HTMLInputElement.stepDown" exists.
func (this HTMLInputElement) HasFuncStepDown1() bool {
	return js.True == bindings.HasFuncHTMLInputElementStepDown1(
		this.ref,
	)
}

// FuncStepDown1 returns the method "HTMLInputElement.stepDown".
func (this HTMLInputElement) FuncStepDown1() (fn js.Func[func()]) {
	bindings.FuncHTMLInputElementStepDown1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// StepDown1 calls the method "HTMLInputElement.stepDown".
func (this HTMLInputElement) StepDown1() (ret js.Void) {
	bindings.CallHTMLInputElementStepDown1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStepDown1 calls the method "HTMLInputElement.stepDown"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryStepDown1() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementStepDown1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCheckValidity returns true if the method "HTMLInputElement.checkValidity" exists.
func (this HTMLInputElement) HasFuncCheckValidity() bool {
	return js.True == bindings.HasFuncHTMLInputElementCheckValidity(
		this.ref,
	)
}

// FuncCheckValidity returns the method "HTMLInputElement.checkValidity".
func (this HTMLInputElement) FuncCheckValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLInputElementCheckValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckValidity calls the method "HTMLInputElement.checkValidity".
func (this HTMLInputElement) CheckValidity() (ret bool) {
	bindings.CallHTMLInputElementCheckValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLInputElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementCheckValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReportValidity returns true if the method "HTMLInputElement.reportValidity" exists.
func (this HTMLInputElement) HasFuncReportValidity() bool {
	return js.True == bindings.HasFuncHTMLInputElementReportValidity(
		this.ref,
	)
}

// FuncReportValidity returns the method "HTMLInputElement.reportValidity".
func (this HTMLInputElement) FuncReportValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLInputElementReportValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportValidity calls the method "HTMLInputElement.reportValidity".
func (this HTMLInputElement) ReportValidity() (ret bool) {
	bindings.CallHTMLInputElementReportValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLInputElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementReportValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetCustomValidity returns true if the method "HTMLInputElement.setCustomValidity" exists.
func (this HTMLInputElement) HasFuncSetCustomValidity() bool {
	return js.True == bindings.HasFuncHTMLInputElementSetCustomValidity(
		this.ref,
	)
}

// FuncSetCustomValidity returns the method "HTMLInputElement.setCustomValidity".
func (this HTMLInputElement) FuncSetCustomValidity() (fn js.Func[func(err js.String)]) {
	bindings.FuncHTMLInputElementSetCustomValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetCustomValidity calls the method "HTMLInputElement.setCustomValidity".
func (this HTMLInputElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLInputElementSetCustomValidity(
		this.ref, js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLInputElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementSetCustomValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

// HasFuncSelect returns true if the method "HTMLInputElement.select" exists.
func (this HTMLInputElement) HasFuncSelect() bool {
	return js.True == bindings.HasFuncHTMLInputElementSelect(
		this.ref,
	)
}

// FuncSelect returns the method "HTMLInputElement.select".
func (this HTMLInputElement) FuncSelect() (fn js.Func[func()]) {
	bindings.FuncHTMLInputElementSelect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Select calls the method "HTMLInputElement.select".
func (this HTMLInputElement) Select() (ret js.Void) {
	bindings.CallHTMLInputElementSelect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySelect calls the method "HTMLInputElement.select"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TrySelect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementSelect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetRangeText returns true if the method "HTMLInputElement.setRangeText" exists.
func (this HTMLInputElement) HasFuncSetRangeText() bool {
	return js.True == bindings.HasFuncHTMLInputElementSetRangeText(
		this.ref,
	)
}

// FuncSetRangeText returns the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) FuncSetRangeText() (fn js.Func[func(replacement js.String)]) {
	bindings.FuncHTMLInputElementSetRangeText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetRangeText calls the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) SetRangeText(replacement js.String) (ret js.Void) {
	bindings.CallHTMLInputElementSetRangeText(
		this.ref, js.Pointer(&ret),
		replacement.Ref(),
	)

	return
}

// TrySetRangeText calls the method "HTMLInputElement.setRangeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TrySetRangeText(replacement js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementSetRangeText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
	)

	return
}

// HasFuncSetRangeText1 returns true if the method "HTMLInputElement.setRangeText" exists.
func (this HTMLInputElement) HasFuncSetRangeText1() bool {
	return js.True == bindings.HasFuncHTMLInputElementSetRangeText1(
		this.ref,
	)
}

// FuncSetRangeText1 returns the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) FuncSetRangeText1() (fn js.Func[func(replacement js.String, start uint32, end uint32, selectionMode SelectionMode)]) {
	bindings.FuncHTMLInputElementSetRangeText1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetRangeText1 calls the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) SetRangeText1(replacement js.String, start uint32, end uint32, selectionMode SelectionMode) (ret js.Void) {
	bindings.CallHTMLInputElementSetRangeText1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
		uint32(start),
		uint32(end),
		uint32(selectionMode),
	)

	return
}

// HasFuncSetRangeText2 returns true if the method "HTMLInputElement.setRangeText" exists.
func (this HTMLInputElement) HasFuncSetRangeText2() bool {
	return js.True == bindings.HasFuncHTMLInputElementSetRangeText2(
		this.ref,
	)
}

// FuncSetRangeText2 returns the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) FuncSetRangeText2() (fn js.Func[func(replacement js.String, start uint32, end uint32)]) {
	bindings.FuncHTMLInputElementSetRangeText2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetRangeText2 calls the method "HTMLInputElement.setRangeText".
func (this HTMLInputElement) SetRangeText2(replacement js.String, start uint32, end uint32) (ret js.Void) {
	bindings.CallHTMLInputElementSetRangeText2(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
		uint32(start),
		uint32(end),
	)

	return
}

// HasFuncSetSelectionRange returns true if the method "HTMLInputElement.setSelectionRange" exists.
func (this HTMLInputElement) HasFuncSetSelectionRange() bool {
	return js.True == bindings.HasFuncHTMLInputElementSetSelectionRange(
		this.ref,
	)
}

// FuncSetSelectionRange returns the method "HTMLInputElement.setSelectionRange".
func (this HTMLInputElement) FuncSetSelectionRange() (fn js.Func[func(start uint32, end uint32, direction js.String)]) {
	bindings.FuncHTMLInputElementSetSelectionRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetSelectionRange calls the method "HTMLInputElement.setSelectionRange".
func (this HTMLInputElement) SetSelectionRange(start uint32, end uint32, direction js.String) (ret js.Void) {
	bindings.CallHTMLInputElementSetSelectionRange(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(start),
		uint32(end),
		direction.Ref(),
	)

	return
}

// HasFuncSetSelectionRange1 returns true if the method "HTMLInputElement.setSelectionRange" exists.
func (this HTMLInputElement) HasFuncSetSelectionRange1() bool {
	return js.True == bindings.HasFuncHTMLInputElementSetSelectionRange1(
		this.ref,
	)
}

// FuncSetSelectionRange1 returns the method "HTMLInputElement.setSelectionRange".
func (this HTMLInputElement) FuncSetSelectionRange1() (fn js.Func[func(start uint32, end uint32)]) {
	bindings.FuncHTMLInputElementSetSelectionRange1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetSelectionRange1 calls the method "HTMLInputElement.setSelectionRange".
func (this HTMLInputElement) SetSelectionRange1(start uint32, end uint32) (ret js.Void) {
	bindings.CallHTMLInputElementSetSelectionRange1(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(start),
		uint32(end),
	)

	return
}

// HasFuncShowPicker returns true if the method "HTMLInputElement.showPicker" exists.
func (this HTMLInputElement) HasFuncShowPicker() bool {
	return js.True == bindings.HasFuncHTMLInputElementShowPicker(
		this.ref,
	)
}

// FuncShowPicker returns the method "HTMLInputElement.showPicker".
func (this HTMLInputElement) FuncShowPicker() (fn js.Func[func()]) {
	bindings.FuncHTMLInputElementShowPicker(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ShowPicker calls the method "HTMLInputElement.showPicker".
func (this HTMLInputElement) ShowPicker() (ret js.Void) {
	bindings.CallHTMLInputElementShowPicker(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryShowPicker calls the method "HTMLInputElement.showPicker"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLInputElement) TryShowPicker() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLInputElementShowPicker(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type HTMLLIElement struct {
	HTMLElement
}

func (this HTMLLIElement) Once() HTMLLIElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "HTMLLIElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLLIElement) Value() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLLIElementValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLLIElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLIElement) SetValue(val int32) bool {
	return js.True == bindings.SetHTMLLIElementValue(
		this.ref,
		int32(val),
	)
}

// Type returns the value of property "HTMLLIElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLLIElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLIElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLLIElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLIElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLLIElementType(
		this.ref,
		val.Ref(),
	)
}

type HTMLLabelElement struct {
	HTMLElement
}

func (this HTMLLabelElement) Once() HTMLLabelElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Form returns the value of property "HTMLLabelElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLLabelElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLLabelElementForm(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HtmlFor returns the value of property "HTMLLabelElement.htmlFor".
//
// It returns ok=false if there is no such property.
func (this HTMLLabelElement) HtmlFor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLabelElementHtmlFor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHtmlFor sets the value of property "HTMLLabelElement.htmlFor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLabelElement) SetHtmlFor(val js.String) bool {
	return js.True == bindings.SetHTMLLabelElementHtmlFor(
		this.ref,
		val.Ref(),
	)
}

// Control returns the value of property "HTMLLabelElement.control".
//
// It returns ok=false if there is no such property.
func (this HTMLLabelElement) Control() (ret HTMLElement, ok bool) {
	ok = js.True == bindings.GetHTMLLabelElementControl(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HTMLLegendElement struct {
	HTMLElement
}

func (this HTMLLegendElement) Once() HTMLLegendElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Form returns the value of property "HTMLLegendElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLLegendElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLLegendElementForm(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Align returns the value of property "HTMLLegendElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLLegendElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLegendElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLLegendElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLegendElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLLegendElementAlign(
		this.ref,
		val.Ref(),
	)
}

type HTMLLinkElement struct {
	HTMLElement
}

func (this HTMLLinkElement) Once() HTMLLinkElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Href returns the value of property "HTMLLinkElement.href".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Href() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementHref(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHref sets the value of property "HTMLLinkElement.href" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetHref(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementHref(
		this.ref,
		val.Ref(),
	)
}

// CrossOrigin returns the value of property "HTMLLinkElement.crossOrigin".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) CrossOrigin() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementCrossOrigin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCrossOrigin sets the value of property "HTMLLinkElement.crossOrigin" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetCrossOrigin(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementCrossOrigin(
		this.ref,
		val.Ref(),
	)
}

// Rel returns the value of property "HTMLLinkElement.rel".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Rel() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementRel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRel sets the value of property "HTMLLinkElement.rel" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetRel(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementRel(
		this.ref,
		val.Ref(),
	)
}

// As returns the value of property "HTMLLinkElement.as".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) As() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementAs(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAs sets the value of property "HTMLLinkElement.as" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetAs(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementAs(
		this.ref,
		val.Ref(),
	)
}

// RelList returns the value of property "HTMLLinkElement.relList".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) RelList() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementRelList(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Media returns the value of property "HTMLLinkElement.media".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementMedia(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMedia sets the value of property "HTMLLinkElement.media" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetMedia(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementMedia(
		this.ref,
		val.Ref(),
	)
}

// Integrity returns the value of property "HTMLLinkElement.integrity".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Integrity() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementIntegrity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetIntegrity sets the value of property "HTMLLinkElement.integrity" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetIntegrity(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementIntegrity(
		this.ref,
		val.Ref(),
	)
}

// Hreflang returns the value of property "HTMLLinkElement.hreflang".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Hreflang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementHreflang(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHreflang sets the value of property "HTMLLinkElement.hreflang" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetHreflang(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementHreflang(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLLinkElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLLinkElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementType(
		this.ref,
		val.Ref(),
	)
}

// Sizes returns the value of property "HTMLLinkElement.sizes".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Sizes() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementSizes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ImageSrcset returns the value of property "HTMLLinkElement.imageSrcset".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) ImageSrcset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementImageSrcset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetImageSrcset sets the value of property "HTMLLinkElement.imageSrcset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetImageSrcset(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementImageSrcset(
		this.ref,
		val.Ref(),
	)
}

// ImageSizes returns the value of property "HTMLLinkElement.imageSizes".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) ImageSizes() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementImageSizes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetImageSizes sets the value of property "HTMLLinkElement.imageSizes" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetImageSizes(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementImageSizes(
		this.ref,
		val.Ref(),
	)
}

// ReferrerPolicy returns the value of property "HTMLLinkElement.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementReferrerPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLLinkElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementReferrerPolicy(
		this.ref,
		val.Ref(),
	)
}

// Blocking returns the value of property "HTMLLinkElement.blocking".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Blocking() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementBlocking(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Disabled returns the value of property "HTMLLinkElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLLinkElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLLinkElementDisabled(
		this.ref,
		js.Bool(bool(val)),
	)
}

// FetchPriority returns the value of property "HTMLLinkElement.fetchPriority".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) FetchPriority() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementFetchPriority(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFetchPriority sets the value of property "HTMLLinkElement.fetchPriority" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetFetchPriority(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementFetchPriority(
		this.ref,
		val.Ref(),
	)
}

// Charset returns the value of property "HTMLLinkElement.charset".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Charset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementCharset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCharset sets the value of property "HTMLLinkElement.charset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetCharset(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementCharset(
		this.ref,
		val.Ref(),
	)
}

// Rev returns the value of property "HTMLLinkElement.rev".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Rev() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementRev(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRev sets the value of property "HTMLLinkElement.rev" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetRev(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementRev(
		this.ref,
		val.Ref(),
	)
}

// Target returns the value of property "HTMLLinkElement.target".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Target() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementTarget(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTarget sets the value of property "HTMLLinkElement.target" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLinkElement) SetTarget(val js.String) bool {
	return js.True == bindings.SetHTMLLinkElementTarget(
		this.ref,
		val.Ref(),
	)
}

// Sheet returns the value of property "HTMLLinkElement.sheet".
//
// It returns ok=false if there is no such property.
func (this HTMLLinkElement) Sheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetHTMLLinkElementSheet(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HTMLMapElement struct {
	HTMLElement
}

func (this HTMLMapElement) Once() HTMLMapElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "HTMLMapElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLMapElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMapElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLMapElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMapElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLMapElementName(
		this.ref,
		val.Ref(),
	)
}

// Areas returns the value of property "HTMLMapElement.areas".
//
// It returns ok=false if there is no such property.
func (this HTMLMapElement) Areas() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLMapElementAreas(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HTMLMarqueeElement struct {
	HTMLElement
}

func (this HTMLMarqueeElement) Once() HTMLMarqueeElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Behavior returns the value of property "HTMLMarqueeElement.behavior".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Behavior() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementBehavior(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBehavior sets the value of property "HTMLMarqueeElement.behavior" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetBehavior(val js.String) bool {
	return js.True == bindings.SetHTMLMarqueeElementBehavior(
		this.ref,
		val.Ref(),
	)
}

// BgColor returns the value of property "HTMLMarqueeElement.bgColor".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementBgColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "HTMLMarqueeElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLMarqueeElementBgColor(
		this.ref,
		val.Ref(),
	)
}

// Direction returns the value of property "HTMLMarqueeElement.direction".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Direction() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementDirection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDirection sets the value of property "HTMLMarqueeElement.direction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetDirection(val js.String) bool {
	return js.True == bindings.SetHTMLMarqueeElementDirection(
		this.ref,
		val.Ref(),
	)
}

// Height returns the value of property "HTMLMarqueeElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLMarqueeElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLMarqueeElementHeight(
		this.ref,
		val.Ref(),
	)
}

// Hspace returns the value of property "HTMLMarqueeElement.hspace".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Hspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementHspace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHspace sets the value of property "HTMLMarqueeElement.hspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetHspace(val uint32) bool {
	return js.True == bindings.SetHTMLMarqueeElementHspace(
		this.ref,
		uint32(val),
	)
}

// Loop returns the value of property "HTMLMarqueeElement.loop".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Loop() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementLoop(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLoop sets the value of property "HTMLMarqueeElement.loop" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetLoop(val int32) bool {
	return js.True == bindings.SetHTMLMarqueeElementLoop(
		this.ref,
		int32(val),
	)
}

// ScrollAmount returns the value of property "HTMLMarqueeElement.scrollAmount".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) ScrollAmount() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementScrollAmount(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollAmount sets the value of property "HTMLMarqueeElement.scrollAmount" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetScrollAmount(val uint32) bool {
	return js.True == bindings.SetHTMLMarqueeElementScrollAmount(
		this.ref,
		uint32(val),
	)
}

// ScrollDelay returns the value of property "HTMLMarqueeElement.scrollDelay".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) ScrollDelay() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementScrollDelay(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScrollDelay sets the value of property "HTMLMarqueeElement.scrollDelay" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetScrollDelay(val uint32) bool {
	return js.True == bindings.SetHTMLMarqueeElementScrollDelay(
		this.ref,
		uint32(val),
	)
}

// TrueSpeed returns the value of property "HTMLMarqueeElement.trueSpeed".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) TrueSpeed() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementTrueSpeed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTrueSpeed sets the value of property "HTMLMarqueeElement.trueSpeed" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetTrueSpeed(val bool) bool {
	return js.True == bindings.SetHTMLMarqueeElementTrueSpeed(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Vspace returns the value of property "HTMLMarqueeElement.vspace".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Vspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementVspace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVspace sets the value of property "HTMLMarqueeElement.vspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetVspace(val uint32) bool {
	return js.True == bindings.SetHTMLMarqueeElementVspace(
		this.ref,
		uint32(val),
	)
}

// Width returns the value of property "HTMLMarqueeElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLMarqueeElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMarqueeElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLMarqueeElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLMarqueeElementWidth(
		this.ref,
		val.Ref(),
	)
}

// HasFuncStart returns true if the method "HTMLMarqueeElement.start" exists.
func (this HTMLMarqueeElement) HasFuncStart() bool {
	return js.True == bindings.HasFuncHTMLMarqueeElementStart(
		this.ref,
	)
}

// FuncStart returns the method "HTMLMarqueeElement.start".
func (this HTMLMarqueeElement) FuncStart() (fn js.Func[func()]) {
	bindings.FuncHTMLMarqueeElementStart(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Start calls the method "HTMLMarqueeElement.start".
func (this HTMLMarqueeElement) Start() (ret js.Void) {
	bindings.CallHTMLMarqueeElementStart(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStart calls the method "HTMLMarqueeElement.start"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLMarqueeElement) TryStart() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMarqueeElementStart(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncStop returns true if the method "HTMLMarqueeElement.stop" exists.
func (this HTMLMarqueeElement) HasFuncStop() bool {
	return js.True == bindings.HasFuncHTMLMarqueeElementStop(
		this.ref,
	)
}

// FuncStop returns the method "HTMLMarqueeElement.stop".
func (this HTMLMarqueeElement) FuncStop() (fn js.Func[func()]) {
	bindings.FuncHTMLMarqueeElementStop(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Stop calls the method "HTMLMarqueeElement.stop".
func (this HTMLMarqueeElement) Stop() (ret js.Void) {
	bindings.CallHTMLMarqueeElementStop(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryStop calls the method "HTMLMarqueeElement.stop"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLMarqueeElement) TryStop() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLMarqueeElementStop(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type HTMLMenuElement struct {
	HTMLElement
}

func (this HTMLMenuElement) Once() HTMLMenuElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Compact returns the value of property "HTMLMenuElement.compact".
//
// It returns ok=false if there is no such property.
func (this HTMLMenuElement) Compact() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLMenuElementCompact(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCompact sets the value of property "HTMLMenuElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMenuElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLMenuElementCompact(
		this.ref,
		js.Bool(bool(val)),
	)
}

type HTMLMetaElement struct {
	HTMLElement
}

func (this HTMLMetaElement) Once() HTMLMetaElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "HTMLMetaElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLMetaElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMetaElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLMetaElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMetaElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLMetaElementName(
		this.ref,
		val.Ref(),
	)
}

// HttpEquiv returns the value of property "HTMLMetaElement.httpEquiv".
//
// It returns ok=false if there is no such property.
func (this HTMLMetaElement) HttpEquiv() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMetaElementHttpEquiv(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHttpEquiv sets the value of property "HTMLMetaElement.httpEquiv" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMetaElement) SetHttpEquiv(val js.String) bool {
	return js.True == bindings.SetHTMLMetaElementHttpEquiv(
		this.ref,
		val.Ref(),
	)
}

// Content returns the value of property "HTMLMetaElement.content".
//
// It returns ok=false if there is no such property.
func (this HTMLMetaElement) Content() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMetaElementContent(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetContent sets the value of property "HTMLMetaElement.content" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMetaElement) SetContent(val js.String) bool {
	return js.True == bindings.SetHTMLMetaElementContent(
		this.ref,
		val.Ref(),
	)
}

// Media returns the value of property "HTMLMetaElement.media".
//
// It returns ok=false if there is no such property.
func (this HTMLMetaElement) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMetaElementMedia(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMedia sets the value of property "HTMLMetaElement.media" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMetaElement) SetMedia(val js.String) bool {
	return js.True == bindings.SetHTMLMetaElementMedia(
		this.ref,
		val.Ref(),
	)
}

// Scheme returns the value of property "HTMLMetaElement.scheme".
//
// It returns ok=false if there is no such property.
func (this HTMLMetaElement) Scheme() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLMetaElementScheme(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScheme sets the value of property "HTMLMetaElement.scheme" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMetaElement) SetScheme(val js.String) bool {
	return js.True == bindings.SetHTMLMetaElementScheme(
		this.ref,
		val.Ref(),
	)
}

type HTMLMeterElement struct {
	HTMLElement
}

func (this HTMLMeterElement) Once() HTMLMeterElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "HTMLMeterElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Value() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLMeterElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetValue(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementValue(
		this.ref,
		float64(val),
	)
}

// Min returns the value of property "HTMLMeterElement.min".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Min() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementMin(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMin sets the value of property "HTMLMeterElement.min" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetMin(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementMin(
		this.ref,
		float64(val),
	)
}

// Max returns the value of property "HTMLMeterElement.max".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Max() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementMax(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMax sets the value of property "HTMLMeterElement.max" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetMax(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementMax(
		this.ref,
		float64(val),
	)
}

// Low returns the value of property "HTMLMeterElement.low".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Low() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementLow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLow sets the value of property "HTMLMeterElement.low" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetLow(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementLow(
		this.ref,
		float64(val),
	)
}

// High returns the value of property "HTMLMeterElement.high".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) High() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementHigh(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHigh sets the value of property "HTMLMeterElement.high" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetHigh(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementHigh(
		this.ref,
		float64(val),
	)
}

// Optimum returns the value of property "HTMLMeterElement.optimum".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Optimum() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementOptimum(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetOptimum sets the value of property "HTMLMeterElement.optimum" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMeterElement) SetOptimum(val float64) bool {
	return js.True == bindings.SetHTMLMeterElementOptimum(
		this.ref,
		float64(val),
	)
}

// Labels returns the value of property "HTMLMeterElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLMeterElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLMeterElementLabels(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HTMLModElement struct {
	HTMLElement
}

func (this HTMLModElement) Once() HTMLModElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Cite returns the value of property "HTMLModElement.cite".
//
// It returns ok=false if there is no such property.
func (this HTMLModElement) Cite() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLModElementCite(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCite sets the value of property "HTMLModElement.cite" to val.
//
// It returns false if the property cannot be set.
func (this HTMLModElement) SetCite(val js.String) bool {
	return js.True == bindings.SetHTMLModElementCite(
		this.ref,
		val.Ref(),
	)
}

// DateTime returns the value of property "HTMLModElement.dateTime".
//
// It returns ok=false if there is no such property.
func (this HTMLModElement) DateTime() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLModElementDateTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDateTime sets the value of property "HTMLModElement.dateTime" to val.
//
// It returns false if the property cannot be set.
func (this HTMLModElement) SetDateTime(val js.String) bool {
	return js.True == bindings.SetHTMLModElementDateTime(
		this.ref,
		val.Ref(),
	)
}

type HTMLModelElement struct {
	HTMLElement
}

func (this HTMLModelElement) Once() HTMLModelElement {
	this.ref.Once()
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
	this.ref.Free()
}

type HTMLOListElement struct {
	HTMLElement
}

func (this HTMLOListElement) Once() HTMLOListElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Reversed returns the value of property "HTMLOListElement.reversed".
//
// It returns ok=false if there is no such property.
func (this HTMLOListElement) Reversed() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOListElementReversed(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReversed sets the value of property "HTMLOListElement.reversed" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOListElement) SetReversed(val bool) bool {
	return js.True == bindings.SetHTMLOListElementReversed(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Start returns the value of property "HTMLOListElement.start".
//
// It returns ok=false if there is no such property.
func (this HTMLOListElement) Start() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLOListElementStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetStart sets the value of property "HTMLOListElement.start" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOListElement) SetStart(val int32) bool {
	return js.True == bindings.SetHTMLOListElementStart(
		this.ref,
		int32(val),
	)
}

// Type returns the value of property "HTMLOListElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLOListElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOListElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLOListElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOListElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLOListElementType(
		this.ref,
		val.Ref(),
	)
}

// Compact returns the value of property "HTMLOListElement.compact".
//
// It returns ok=false if there is no such property.
func (this HTMLOListElement) Compact() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOListElementCompact(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCompact sets the value of property "HTMLOListElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOListElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLOListElementCompact(
		this.ref,
		js.Bool(bool(val)),
	)
}

type HTMLObjectElement struct {
	HTMLElement
}

func (this HTMLObjectElement) Once() HTMLObjectElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Data returns the value of property "HTMLObjectElement.data".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Data() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementData(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetData sets the value of property "HTMLObjectElement.data" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetData(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementData(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLObjectElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLObjectElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementType(
		this.ref,
		val.Ref(),
	)
}

// Name returns the value of property "HTMLObjectElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLObjectElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementName(
		this.ref,
		val.Ref(),
	)
}

// Form returns the value of property "HTMLObjectElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementForm(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Width returns the value of property "HTMLObjectElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLObjectElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementWidth(
		this.ref,
		val.Ref(),
	)
}

// Height returns the value of property "HTMLObjectElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLObjectElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementHeight(
		this.ref,
		val.Ref(),
	)
}

// ContentDocument returns the value of property "HTMLObjectElement.contentDocument".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) ContentDocument() (ret Document, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementContentDocument(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ContentWindow returns the value of property "HTMLObjectElement.contentWindow".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) ContentWindow() (ret js.Object, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementContentWindow(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WillValidate returns the value of property "HTMLObjectElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementWillValidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLObjectElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementValidity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLObjectElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementValidationMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Align returns the value of property "HTMLObjectElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLObjectElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementAlign(
		this.ref,
		val.Ref(),
	)
}

// Archive returns the value of property "HTMLObjectElement.archive".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Archive() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementArchive(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetArchive sets the value of property "HTMLObjectElement.archive" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetArchive(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementArchive(
		this.ref,
		val.Ref(),
	)
}

// Code returns the value of property "HTMLObjectElement.code".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Code() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementCode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCode sets the value of property "HTMLObjectElement.code" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetCode(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementCode(
		this.ref,
		val.Ref(),
	)
}

// Declare returns the value of property "HTMLObjectElement.declare".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Declare() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementDeclare(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDeclare sets the value of property "HTMLObjectElement.declare" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetDeclare(val bool) bool {
	return js.True == bindings.SetHTMLObjectElementDeclare(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Hspace returns the value of property "HTMLObjectElement.hspace".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Hspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementHspace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHspace sets the value of property "HTMLObjectElement.hspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetHspace(val uint32) bool {
	return js.True == bindings.SetHTMLObjectElementHspace(
		this.ref,
		uint32(val),
	)
}

// Standby returns the value of property "HTMLObjectElement.standby".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Standby() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementStandby(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetStandby sets the value of property "HTMLObjectElement.standby" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetStandby(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementStandby(
		this.ref,
		val.Ref(),
	)
}

// Vspace returns the value of property "HTMLObjectElement.vspace".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Vspace() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementVspace(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVspace sets the value of property "HTMLObjectElement.vspace" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetVspace(val uint32) bool {
	return js.True == bindings.SetHTMLObjectElementVspace(
		this.ref,
		uint32(val),
	)
}

// CodeBase returns the value of property "HTMLObjectElement.codeBase".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) CodeBase() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementCodeBase(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCodeBase sets the value of property "HTMLObjectElement.codeBase" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetCodeBase(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementCodeBase(
		this.ref,
		val.Ref(),
	)
}

// CodeType returns the value of property "HTMLObjectElement.codeType".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) CodeType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementCodeType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCodeType sets the value of property "HTMLObjectElement.codeType" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetCodeType(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementCodeType(
		this.ref,
		val.Ref(),
	)
}

// UseMap returns the value of property "HTMLObjectElement.useMap".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) UseMap() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementUseMap(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetUseMap sets the value of property "HTMLObjectElement.useMap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetUseMap(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementUseMap(
		this.ref,
		val.Ref(),
	)
}

// Border returns the value of property "HTMLObjectElement.border".
//
// It returns ok=false if there is no such property.
func (this HTMLObjectElement) Border() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLObjectElementBorder(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBorder sets the value of property "HTMLObjectElement.border" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetBorder(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementBorder(
		this.ref,
		val.Ref(),
	)
}

// HasFuncGetSVGDocument returns true if the method "HTMLObjectElement.getSVGDocument" exists.
func (this HTMLObjectElement) HasFuncGetSVGDocument() bool {
	return js.True == bindings.HasFuncHTMLObjectElementGetSVGDocument(
		this.ref,
	)
}

// FuncGetSVGDocument returns the method "HTMLObjectElement.getSVGDocument".
func (this HTMLObjectElement) FuncGetSVGDocument() (fn js.Func[func() Document]) {
	bindings.FuncHTMLObjectElementGetSVGDocument(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetSVGDocument calls the method "HTMLObjectElement.getSVGDocument".
func (this HTMLObjectElement) GetSVGDocument() (ret Document) {
	bindings.CallHTMLObjectElementGetSVGDocument(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetSVGDocument calls the method "HTMLObjectElement.getSVGDocument"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLObjectElement) TryGetSVGDocument() (ret Document, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLObjectElementGetSVGDocument(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCheckValidity returns true if the method "HTMLObjectElement.checkValidity" exists.
func (this HTMLObjectElement) HasFuncCheckValidity() bool {
	return js.True == bindings.HasFuncHTMLObjectElementCheckValidity(
		this.ref,
	)
}

// FuncCheckValidity returns the method "HTMLObjectElement.checkValidity".
func (this HTMLObjectElement) FuncCheckValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLObjectElementCheckValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckValidity calls the method "HTMLObjectElement.checkValidity".
func (this HTMLObjectElement) CheckValidity() (ret bool) {
	bindings.CallHTMLObjectElementCheckValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLObjectElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLObjectElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLObjectElementCheckValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReportValidity returns true if the method "HTMLObjectElement.reportValidity" exists.
func (this HTMLObjectElement) HasFuncReportValidity() bool {
	return js.True == bindings.HasFuncHTMLObjectElementReportValidity(
		this.ref,
	)
}

// FuncReportValidity returns the method "HTMLObjectElement.reportValidity".
func (this HTMLObjectElement) FuncReportValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLObjectElementReportValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportValidity calls the method "HTMLObjectElement.reportValidity".
func (this HTMLObjectElement) ReportValidity() (ret bool) {
	bindings.CallHTMLObjectElementReportValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLObjectElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLObjectElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLObjectElementReportValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetCustomValidity returns true if the method "HTMLObjectElement.setCustomValidity" exists.
func (this HTMLObjectElement) HasFuncSetCustomValidity() bool {
	return js.True == bindings.HasFuncHTMLObjectElementSetCustomValidity(
		this.ref,
	)
}

// FuncSetCustomValidity returns the method "HTMLObjectElement.setCustomValidity".
func (this HTMLObjectElement) FuncSetCustomValidity() (fn js.Func[func(err js.String)]) {
	bindings.FuncHTMLObjectElementSetCustomValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetCustomValidity calls the method "HTMLObjectElement.setCustomValidity".
func (this HTMLObjectElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLObjectElementSetCustomValidity(
		this.ref, js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLObjectElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLObjectElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLObjectElementSetCustomValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

type HTMLOptGroupElement struct {
	HTMLElement
}

func (this HTMLOptGroupElement) Once() HTMLOptGroupElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Disabled returns the value of property "HTMLOptGroupElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLOptGroupElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOptGroupElementDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLOptGroupElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptGroupElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLOptGroupElementDisabled(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Label returns the value of property "HTMLOptGroupElement.label".
//
// It returns ok=false if there is no such property.
func (this HTMLOptGroupElement) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOptGroupElementLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "HTMLOptGroupElement.label" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptGroupElement) SetLabel(val js.String) bool {
	return js.True == bindings.SetHTMLOptGroupElementLabel(
		this.ref,
		val.Ref(),
	)
}

type HTMLOptionElement struct {
	HTMLElement
}

func (this HTMLOptionElement) Once() HTMLOptionElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Disabled returns the value of property "HTMLOptionElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLOptionElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLOptionElementDisabled(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLOptionElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementForm(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Label returns the value of property "HTMLOptionElement.label".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "HTMLOptionElement.label" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetLabel(val js.String) bool {
	return js.True == bindings.SetHTMLOptionElementLabel(
		this.ref,
		val.Ref(),
	)
}

// DefaultSelected returns the value of property "HTMLOptionElement.defaultSelected".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) DefaultSelected() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementDefaultSelected(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDefaultSelected sets the value of property "HTMLOptionElement.defaultSelected" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetDefaultSelected(val bool) bool {
	return js.True == bindings.SetHTMLOptionElementDefaultSelected(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Selected returns the value of property "HTMLOptionElement.selected".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Selected() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementSelected(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelected sets the value of property "HTMLOptionElement.selected" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetSelected(val bool) bool {
	return js.True == bindings.SetHTMLOptionElementSelected(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Value returns the value of property "HTMLOptionElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLOptionElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLOptionElementValue(
		this.ref,
		val.Ref(),
	)
}

// Text returns the value of property "HTMLOptionElement.text".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "HTMLOptionElement.text" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionElement) SetText(val js.String) bool {
	return js.True == bindings.SetHTMLOptionElementText(
		this.ref,
		val.Ref(),
	)
}

// Index returns the value of property "HTMLOptionElement.index".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionElement) Index() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLOptionElementIndex(
		this.ref, js.Pointer(&ret),
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
	this.ref.Once()
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
	this.ref.Free()
}

// Length returns the value of property "HTMLOptionsCollection.length".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionsCollection) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLOptionsCollectionLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLength sets the value of property "HTMLOptionsCollection.length" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionsCollection) SetLength(val uint32) bool {
	return js.True == bindings.SetHTMLOptionsCollectionLength(
		this.ref,
		uint32(val),
	)
}

// SelectedIndex returns the value of property "HTMLOptionsCollection.selectedIndex".
//
// It returns ok=false if there is no such property.
func (this HTMLOptionsCollection) SelectedIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLOptionsCollectionSelectedIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectedIndex sets the value of property "HTMLOptionsCollection.selectedIndex" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionsCollection) SetSelectedIndex(val int32) bool {
	return js.True == bindings.SetHTMLOptionsCollectionSelectedIndex(
		this.ref,
		int32(val),
	)
}

// HasFuncSet returns true if the method "HTMLOptionsCollection." exists.
func (this HTMLOptionsCollection) HasFuncSet() bool {
	return js.True == bindings.HasFuncHTMLOptionsCollectionSet(
		this.ref,
	)
}

// FuncSet returns the method "HTMLOptionsCollection.".
func (this HTMLOptionsCollection) FuncSet() (fn js.Func[func(index uint32, option HTMLOptionElement)]) {
	bindings.FuncHTMLOptionsCollectionSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "HTMLOptionsCollection.".
func (this HTMLOptionsCollection) Set(index uint32, option HTMLOptionElement) (ret js.Void) {
	bindings.CallHTMLOptionsCollectionSet(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		option.Ref(),
	)

	return
}

// HasFuncAdd returns true if the method "HTMLOptionsCollection.add" exists.
func (this HTMLOptionsCollection) HasFuncAdd() bool {
	return js.True == bindings.HasFuncHTMLOptionsCollectionAdd(
		this.ref,
	)
}

// FuncAdd returns the method "HTMLOptionsCollection.add".
func (this HTMLOptionsCollection) FuncAdd() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32)]) {
	bindings.FuncHTMLOptionsCollectionAdd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add calls the method "HTMLOptionsCollection.add".
func (this HTMLOptionsCollection) Add(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32) (ret js.Void) {
	bindings.CallHTMLOptionsCollectionAdd(
		this.ref, js.Pointer(&ret),
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
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
		before.Ref(),
	)

	return
}

// HasFuncAdd1 returns true if the method "HTMLOptionsCollection.add" exists.
func (this HTMLOptionsCollection) HasFuncAdd1() bool {
	return js.True == bindings.HasFuncHTMLOptionsCollectionAdd1(
		this.ref,
	)
}

// FuncAdd1 returns the method "HTMLOptionsCollection.add".
func (this HTMLOptionsCollection) FuncAdd1() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement)]) {
	bindings.FuncHTMLOptionsCollectionAdd1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add1 calls the method "HTMLOptionsCollection.add".
func (this HTMLOptionsCollection) Add1(element OneOf_HTMLOptionElement_HTMLOptGroupElement) (ret js.Void) {
	bindings.CallHTMLOptionsCollectionAdd1(
		this.ref, js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryAdd1 calls the method "HTMLOptionsCollection.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOptionsCollection) TryAdd1(element OneOf_HTMLOptionElement_HTMLOptGroupElement) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOptionsCollectionAdd1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

// HasFuncRemove returns true if the method "HTMLOptionsCollection.remove" exists.
func (this HTMLOptionsCollection) HasFuncRemove() bool {
	return js.True == bindings.HasFuncHTMLOptionsCollectionRemove(
		this.ref,
	)
}

// FuncRemove returns the method "HTMLOptionsCollection.remove".
func (this HTMLOptionsCollection) FuncRemove() (fn js.Func[func(index int32)]) {
	bindings.FuncHTMLOptionsCollectionRemove(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove calls the method "HTMLOptionsCollection.remove".
func (this HTMLOptionsCollection) Remove(index int32) (ret js.Void) {
	bindings.CallHTMLOptionsCollectionRemove(
		this.ref, js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryRemove calls the method "HTMLOptionsCollection.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOptionsCollection) TryRemove(index int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOptionsCollectionRemove(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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

type HTMLOutputElement struct {
	HTMLElement
}

func (this HTMLOutputElement) Once() HTMLOutputElement {
	this.ref.Once()
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
	this.ref.Free()
}

// HtmlFor returns the value of property "HTMLOutputElement.htmlFor".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) HtmlFor() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementHtmlFor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Form returns the value of property "HTMLOutputElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementForm(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Name returns the value of property "HTMLOutputElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLOutputElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOutputElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLOutputElementName(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLOutputElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DefaultValue returns the value of property "HTMLOutputElement.defaultValue".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) DefaultValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementDefaultValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDefaultValue sets the value of property "HTMLOutputElement.defaultValue" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOutputElement) SetDefaultValue(val js.String) bool {
	return js.True == bindings.SetHTMLOutputElementDefaultValue(
		this.ref,
		val.Ref(),
	)
}

// Value returns the value of property "HTMLOutputElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLOutputElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOutputElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLOutputElementValue(
		this.ref,
		val.Ref(),
	)
}

// WillValidate returns the value of property "HTMLOutputElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementWillValidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLOutputElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementValidity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLOutputElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementValidationMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLOutputElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLOutputElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLOutputElementLabels(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncCheckValidity returns true if the method "HTMLOutputElement.checkValidity" exists.
func (this HTMLOutputElement) HasFuncCheckValidity() bool {
	return js.True == bindings.HasFuncHTMLOutputElementCheckValidity(
		this.ref,
	)
}

// FuncCheckValidity returns the method "HTMLOutputElement.checkValidity".
func (this HTMLOutputElement) FuncCheckValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLOutputElementCheckValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckValidity calls the method "HTMLOutputElement.checkValidity".
func (this HTMLOutputElement) CheckValidity() (ret bool) {
	bindings.CallHTMLOutputElementCheckValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLOutputElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOutputElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOutputElementCheckValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReportValidity returns true if the method "HTMLOutputElement.reportValidity" exists.
func (this HTMLOutputElement) HasFuncReportValidity() bool {
	return js.True == bindings.HasFuncHTMLOutputElementReportValidity(
		this.ref,
	)
}

// FuncReportValidity returns the method "HTMLOutputElement.reportValidity".
func (this HTMLOutputElement) FuncReportValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLOutputElementReportValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportValidity calls the method "HTMLOutputElement.reportValidity".
func (this HTMLOutputElement) ReportValidity() (ret bool) {
	bindings.CallHTMLOutputElementReportValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLOutputElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOutputElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOutputElementReportValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetCustomValidity returns true if the method "HTMLOutputElement.setCustomValidity" exists.
func (this HTMLOutputElement) HasFuncSetCustomValidity() bool {
	return js.True == bindings.HasFuncHTMLOutputElementSetCustomValidity(
		this.ref,
	)
}

// FuncSetCustomValidity returns the method "HTMLOutputElement.setCustomValidity".
func (this HTMLOutputElement) FuncSetCustomValidity() (fn js.Func[func(err js.String)]) {
	bindings.FuncHTMLOutputElementSetCustomValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetCustomValidity calls the method "HTMLOutputElement.setCustomValidity".
func (this HTMLOutputElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLOutputElementSetCustomValidity(
		this.ref, js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLOutputElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLOutputElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLOutputElementSetCustomValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

type HTMLParagraphElement struct {
	HTMLElement
}

func (this HTMLParagraphElement) Once() HTMLParagraphElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Align returns the value of property "HTMLParagraphElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLParagraphElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLParagraphElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLParagraphElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParagraphElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLParagraphElementAlign(
		this.ref,
		val.Ref(),
	)
}

type HTMLParamElement struct {
	HTMLElement
}

func (this HTMLParamElement) Once() HTMLParamElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "HTMLParamElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLParamElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLParamElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLParamElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParamElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLParamElementName(
		this.ref,
		val.Ref(),
	)
}

// Value returns the value of property "HTMLParamElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLParamElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLParamElementValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLParamElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParamElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLParamElementValue(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLParamElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLParamElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLParamElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLParamElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParamElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLParamElementType(
		this.ref,
		val.Ref(),
	)
}

// ValueType returns the value of property "HTMLParamElement.valueType".
//
// It returns ok=false if there is no such property.
func (this HTMLParamElement) ValueType() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLParamElementValueType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValueType sets the value of property "HTMLParamElement.valueType" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParamElement) SetValueType(val js.String) bool {
	return js.True == bindings.SetHTMLParamElementValueType(
		this.ref,
		val.Ref(),
	)
}

type HTMLPictureElement struct {
	HTMLElement
}

func (this HTMLPictureElement) Once() HTMLPictureElement {
	this.ref.Once()
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
	this.ref.Free()
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
func (p *PortalActivateOptions) UpdateFrom(ref js.Ref) {
	bindings.PortalActivateOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *PortalActivateOptions) Update(ref js.Ref) {
	bindings.PortalActivateOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *PortalActivateOptions) FreeMembers(recursive bool) {
	js.Free(
		p.Data.Ref(),
		p.Transfer.Ref(),
	)
	p.Data = p.Data.FromRef(js.Undefined)
	p.Transfer = p.Transfer.FromRef(js.Undefined)
}

// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/std/core/abi"
	"github.com/primecitizens/std/core/assert"
	"github.com/primecitizens/std/ffi/js"
	"github.com/primecitizens/std/plat/js/web/bindings"
)

func _() {
	var (
		_ abi.FuncID
	)
	assert.TODO()
}

func NewHTMLAudioElement() HTMLAudioElement {
	return HTMLAudioElement{}.FromRef(
		bindings.NewHTMLAudioElementByHTMLAudioElement(),
	)
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

func NewHTMLBRElement() HTMLBRElement {
	return HTMLBRElement{}.FromRef(
		bindings.NewHTMLBRElementByHTMLBRElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLBRElement) Clear() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLBRElementClear(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Clear sets the value of property "HTMLBRElement.clear" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBRElement) SetClear(val js.String) bool {
	return js.True == bindings.SetHTMLBRElementClear(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLBaseElement() HTMLBaseElement {
	return HTMLBaseElement{}.FromRef(
		bindings.NewHTMLBaseElementByHTMLBaseElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLBaseElement) Href() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLBaseElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Href sets the value of property "HTMLBaseElement.href" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLBaseElement) Target() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLBaseElementTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Target sets the value of property "HTMLBaseElement.target" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBaseElement) SetTarget(val js.String) bool {
	return js.True == bindings.SetHTMLBaseElementTarget(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLBodyElement() HTMLBodyElement {
	return HTMLBodyElement{}.FromRef(
		bindings.NewHTMLBodyElementByHTMLBodyElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLBodyElement) Text() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLBodyElementText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Text sets the value of property "HTMLBodyElement.text" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLBodyElement) Link() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLBodyElementLink(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Link sets the value of property "HTMLBodyElement.link" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLBodyElement) VLink() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLBodyElementVLink(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// VLink sets the value of property "HTMLBodyElement.vLink" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLBodyElement) ALink() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLBodyElementALink(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ALink sets the value of property "HTMLBodyElement.aLink" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLBodyElement) BgColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLBodyElementBgColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BgColor sets the value of property "HTMLBodyElement.bgColor" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLBodyElement) Background() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLBodyElementBackground(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Background sets the value of property "HTMLBodyElement.background" to val.
//
// It returns false if the property cannot be set.
func (this HTMLBodyElement) SetBackground(val js.String) bool {
	return js.True == bindings.SetHTMLBodyElementBackground(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLButtonElement() HTMLButtonElement {
	return HTMLButtonElement{}.FromRef(
		bindings.NewHTMLButtonElementByHTMLButtonElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) Disabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Disabled sets the value of property "HTMLButtonElement.disabled" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) Form() (HTMLFormElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementForm(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormElement{}.FromRef(_ret), _ok
}

// FormAction returns the value of property "HTMLButtonElement.formAction".
//
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) FormAction() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementFormAction(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FormAction sets the value of property "HTMLButtonElement.formAction" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) FormEnctype() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementFormEnctype(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FormEnctype sets the value of property "HTMLButtonElement.formEnctype" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) FormMethod() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementFormMethod(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FormMethod sets the value of property "HTMLButtonElement.formMethod" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) FormNoValidate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementFormNoValidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// FormNoValidate sets the value of property "HTMLButtonElement.formNoValidate" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) FormTarget() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementFormTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FormTarget sets the value of property "HTMLButtonElement.formTarget" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLButtonElement.name" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLButtonElement.type" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "HTMLButtonElement.value" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) WillValidate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementWillValidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Validity returns the value of property "HTMLButtonElement.validity".
//
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) Validity() (ValidityState, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementValidity(
		this.Ref(), js.Pointer(&_ok),
	)
	return ValidityState{}.FromRef(_ret), _ok
}

// ValidationMessage returns the value of property "HTMLButtonElement.validationMessage".
//
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) ValidationMessage() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementValidationMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Labels returns the value of property "HTMLButtonElement.labels".
//
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) Labels() (NodeList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementLabels(
		this.Ref(), js.Pointer(&_ok),
	)
	return NodeList{}.FromRef(_ret), _ok
}

// PopoverTargetElement returns the value of property "HTMLButtonElement.popoverTargetElement".
//
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) PopoverTargetElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementPopoverTargetElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// PopoverTargetElement sets the value of property "HTMLButtonElement.popoverTargetElement" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLButtonElement) PopoverTargetAction() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLButtonElementPopoverTargetAction(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// PopoverTargetAction sets the value of property "HTMLButtonElement.popoverTargetAction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLButtonElement) SetPopoverTargetAction(val js.String) bool {
	return js.True == bindings.SetHTMLButtonElementPopoverTargetAction(
		this.Ref(),
		val.Ref(),
	)
}

// CheckValidity calls the method "HTMLButtonElement.checkValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLButtonElement) CheckValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLButtonElementCheckValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// CheckValidityFunc returns the method "HTMLButtonElement.checkValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLButtonElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLButtonElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLButtonElement.reportValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLButtonElement) ReportValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLButtonElementReportValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// ReportValidityFunc returns the method "HTMLButtonElement.reportValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLButtonElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLButtonElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLButtonElement.setCustomValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLButtonElement) SetCustomValidity(err js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLButtonElementSetCustomValidity(
		this.Ref(), js.Pointer(&_ok),
		err.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCustomValidityFunc returns the method "HTMLButtonElement.setCustomValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLButtonElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLButtonElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

func NewHTMLDListElement() HTMLDListElement {
	return HTMLDListElement{}.FromRef(
		bindings.NewHTMLDListElementByHTMLDListElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLDListElement) Compact() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLDListElementCompact(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Compact sets the value of property "HTMLDListElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDListElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLDListElementCompact(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

func NewHTMLDataElement() HTMLDataElement {
	return HTMLDataElement{}.FromRef(
		bindings.NewHTMLDataElementByHTMLDataElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLDataElement) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLDataElementValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "HTMLDataElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDataElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLDataElementValue(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLDataListElement() HTMLDataListElement {
	return HTMLDataListElement{}.FromRef(
		bindings.NewHTMLDataListElementByHTMLDataListElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLDataListElement) Options() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetHTMLDataListElementOptions(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

func NewHTMLDetailsElement() HTMLDetailsElement {
	return HTMLDetailsElement{}.FromRef(
		bindings.NewHTMLDetailsElementByHTMLDetailsElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLDetailsElement) Open() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLDetailsElementOpen(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Open sets the value of property "HTMLDetailsElement.open" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDetailsElement) SetOpen(val bool) bool {
	return js.True == bindings.SetHTMLDetailsElementOpen(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

func NewHTMLDialogElement() HTMLDialogElement {
	return HTMLDialogElement{}.FromRef(
		bindings.NewHTMLDialogElementByHTMLDialogElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLDialogElement) Open() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLDialogElementOpen(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Open sets the value of property "HTMLDialogElement.open" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLDialogElement) ReturnValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLDialogElementReturnValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReturnValue sets the value of property "HTMLDialogElement.returnValue" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDialogElement) SetReturnValue(val js.String) bool {
	return js.True == bindings.SetHTMLDialogElementReturnValue(
		this.Ref(),
		val.Ref(),
	)
}

// Show calls the method "HTMLDialogElement.show".
//
// The returned bool will be false if there is no such method.
func (this HTMLDialogElement) Show() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLDialogElementShow(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ShowFunc returns the method "HTMLDialogElement.show".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLDialogElement) ShowFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLDialogElementShowFunc(
			this.Ref(),
		),
	)
}

// ShowModal calls the method "HTMLDialogElement.showModal".
//
// The returned bool will be false if there is no such method.
func (this HTMLDialogElement) ShowModal() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLDialogElementShowModal(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ShowModalFunc returns the method "HTMLDialogElement.showModal".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLDialogElement) ShowModalFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLDialogElementShowModalFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "HTMLDialogElement.close".
//
// The returned bool will be false if there is no such method.
func (this HTMLDialogElement) Close(returnValue js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLDialogElementClose(
		this.Ref(), js.Pointer(&_ok),
		returnValue.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "HTMLDialogElement.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLDialogElement) CloseFunc() (fn js.Func[func(returnValue js.String)]) {
	return fn.FromRef(
		bindings.HTMLDialogElementCloseFunc(
			this.Ref(),
		),
	)
}

// Close1 calls the method "HTMLDialogElement.close".
//
// The returned bool will be false if there is no such method.
func (this HTMLDialogElement) Close1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLDialogElementClose1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Close1Func returns the method "HTMLDialogElement.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLDialogElement) Close1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLDialogElementClose1Func(
			this.Ref(),
		),
	)
}

func NewHTMLDirectoryElement() HTMLDirectoryElement {
	return HTMLDirectoryElement{}.FromRef(
		bindings.NewHTMLDirectoryElementByHTMLDirectoryElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLDirectoryElement) Compact() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLDirectoryElementCompact(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Compact sets the value of property "HTMLDirectoryElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDirectoryElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLDirectoryElementCompact(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

func NewHTMLDivElement() HTMLDivElement {
	return HTMLDivElement{}.FromRef(
		bindings.NewHTMLDivElementByHTMLDivElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLDivElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLDivElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLDivElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLDivElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLDivElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLEmbedElement() HTMLEmbedElement {
	return HTMLEmbedElement{}.FromRef(
		bindings.NewHTMLEmbedElementByHTMLEmbedElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLEmbedElement) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLEmbedElementSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src sets the value of property "HTMLEmbedElement.src" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLEmbedElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLEmbedElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLEmbedElement.type" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLEmbedElement) Width() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLEmbedElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Width sets the value of property "HTMLEmbedElement.width" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLEmbedElement) Height() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLEmbedElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Height sets the value of property "HTMLEmbedElement.height" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLEmbedElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLEmbedElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLEmbedElement.align" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLEmbedElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLEmbedElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLEmbedElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLEmbedElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLEmbedElementName(
		this.Ref(),
		val.Ref(),
	)
}

// GetSVGDocument calls the method "HTMLEmbedElement.getSVGDocument".
//
// The returned bool will be false if there is no such method.
func (this HTMLEmbedElement) GetSVGDocument() (Document, bool) {
	var _ok bool
	_ret := bindings.CallHTMLEmbedElementGetSVGDocument(
		this.Ref(), js.Pointer(&_ok),
	)

	return Document{}.FromRef(_ret), _ok
}

// GetSVGDocumentFunc returns the method "HTMLEmbedElement.getSVGDocument".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLEmbedElement) GetSVGDocumentFunc() (fn js.Func[func() Document]) {
	return fn.FromRef(
		bindings.HTMLEmbedElementGetSVGDocumentFunc(
			this.Ref(),
		),
	)
}

func NewHTMLFencedFrameElement() HTMLFencedFrameElement {
	return HTMLFencedFrameElement{}.FromRef(
		bindings.NewHTMLFencedFrameElementByHTMLFencedFrameElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLFencedFrameElement) Config() (FencedFrameConfig, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFencedFrameElementConfig(
		this.Ref(), js.Pointer(&_ok),
	)
	return FencedFrameConfig{}.FromRef(_ret), _ok
}

// Config sets the value of property "HTMLFencedFrameElement.config" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFencedFrameElement) Width() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFencedFrameElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Width sets the value of property "HTMLFencedFrameElement.width" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFencedFrameElement) Height() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFencedFrameElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Height sets the value of property "HTMLFencedFrameElement.height" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFencedFrameElement) Allow() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFencedFrameElementAllow(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Allow sets the value of property "HTMLFencedFrameElement.allow" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFencedFrameElement) SetAllow(val js.String) bool {
	return js.True == bindings.SetHTMLFencedFrameElementAllow(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLFieldSetElement() HTMLFieldSetElement {
	return HTMLFieldSetElement{}.FromRef(
		bindings.NewHTMLFieldSetElementByHTMLFieldSetElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLFieldSetElement) Disabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFieldSetElementDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Disabled sets the value of property "HTMLFieldSetElement.disabled" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFieldSetElement) Form() (HTMLFormElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFieldSetElementForm(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormElement{}.FromRef(_ret), _ok
}

// Name returns the value of property "HTMLFieldSetElement.name".
//
// The returned bool will be false if there is no such property.
func (this HTMLFieldSetElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFieldSetElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLFieldSetElement.name" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFieldSetElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFieldSetElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Elements returns the value of property "HTMLFieldSetElement.elements".
//
// The returned bool will be false if there is no such property.
func (this HTMLFieldSetElement) Elements() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFieldSetElementElements(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// WillValidate returns the value of property "HTMLFieldSetElement.willValidate".
//
// The returned bool will be false if there is no such property.
func (this HTMLFieldSetElement) WillValidate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFieldSetElementWillValidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Validity returns the value of property "HTMLFieldSetElement.validity".
//
// The returned bool will be false if there is no such property.
func (this HTMLFieldSetElement) Validity() (ValidityState, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFieldSetElementValidity(
		this.Ref(), js.Pointer(&_ok),
	)
	return ValidityState{}.FromRef(_ret), _ok
}

// ValidationMessage returns the value of property "HTMLFieldSetElement.validationMessage".
//
// The returned bool will be false if there is no such property.
func (this HTMLFieldSetElement) ValidationMessage() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFieldSetElementValidationMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CheckValidity calls the method "HTMLFieldSetElement.checkValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLFieldSetElement) CheckValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFieldSetElementCheckValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// CheckValidityFunc returns the method "HTMLFieldSetElement.checkValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFieldSetElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLFieldSetElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLFieldSetElement.reportValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLFieldSetElement) ReportValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFieldSetElementReportValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// ReportValidityFunc returns the method "HTMLFieldSetElement.reportValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFieldSetElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLFieldSetElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLFieldSetElement.setCustomValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLFieldSetElement) SetCustomValidity(err js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLFieldSetElementSetCustomValidity(
		this.Ref(), js.Pointer(&_ok),
		err.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCustomValidityFunc returns the method "HTMLFieldSetElement.setCustomValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLFieldSetElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLFieldSetElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

func NewHTMLFontElement() HTMLFontElement {
	return HTMLFontElement{}.FromRef(
		bindings.NewHTMLFontElementByHTMLFontElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLFontElement) Color() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFontElementColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Color sets the value of property "HTMLFontElement.color" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFontElement) Face() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFontElementFace(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Face sets the value of property "HTMLFontElement.face" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFontElement) Size() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFontElementSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Size sets the value of property "HTMLFontElement.size" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFontElement) SetSize(val js.String) bool {
	return js.True == bindings.SetHTMLFontElementSize(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLFrameElement() HTMLFrameElement {
	return HTMLFrameElement{}.FromRef(
		bindings.NewHTMLFrameElementByHTMLFrameElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLFrameElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLFrameElement.name" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFrameElement) Scrolling() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameElementScrolling(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Scrolling sets the value of property "HTMLFrameElement.scrolling" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFrameElement) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameElementSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src sets the value of property "HTMLFrameElement.src" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFrameElement) FrameBorder() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameElementFrameBorder(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FrameBorder sets the value of property "HTMLFrameElement.frameBorder" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFrameElement) LongDesc() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameElementLongDesc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LongDesc sets the value of property "HTMLFrameElement.longDesc" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFrameElement) NoResize() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameElementNoResize(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// NoResize sets the value of property "HTMLFrameElement.noResize" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFrameElement) ContentDocument() (Document, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameElementContentDocument(
		this.Ref(), js.Pointer(&_ok),
	)
	return Document{}.FromRef(_ret), _ok
}

// ContentWindow returns the value of property "HTMLFrameElement.contentWindow".
//
// The returned bool will be false if there is no such property.
func (this HTMLFrameElement) ContentWindow() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameElementContentWindow(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// MarginHeight returns the value of property "HTMLFrameElement.marginHeight".
//
// The returned bool will be false if there is no such property.
func (this HTMLFrameElement) MarginHeight() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameElementMarginHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// MarginHeight sets the value of property "HTMLFrameElement.marginHeight" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFrameElement) MarginWidth() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameElementMarginWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// MarginWidth sets the value of property "HTMLFrameElement.marginWidth" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameElement) SetMarginWidth(val js.String) bool {
	return js.True == bindings.SetHTMLFrameElementMarginWidth(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLFrameSetElement() HTMLFrameSetElement {
	return HTMLFrameSetElement{}.FromRef(
		bindings.NewHTMLFrameSetElementByHTMLFrameSetElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLFrameSetElement) Cols() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameSetElementCols(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Cols sets the value of property "HTMLFrameSetElement.cols" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLFrameSetElement) Rows() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLFrameSetElementRows(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Rows sets the value of property "HTMLFrameSetElement.rows" to val.
//
// It returns false if the property cannot be set.
func (this HTMLFrameSetElement) SetRows(val js.String) bool {
	return js.True == bindings.SetHTMLFrameSetElementRows(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLHRElement() HTMLHRElement {
	return HTMLHRElement{}.FromRef(
		bindings.NewHTMLHRElementByHTMLHRElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLHRElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLHRElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLHRElement.align" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLHRElement) Color() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLHRElementColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Color sets the value of property "HTMLHRElement.color" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLHRElement) NoShade() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLHRElementNoShade(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// NoShade sets the value of property "HTMLHRElement.noShade" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLHRElement) Size() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLHRElementSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Size sets the value of property "HTMLHRElement.size" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLHRElement) Width() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLHRElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Width sets the value of property "HTMLHRElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHRElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLHRElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLHeadingElement() HTMLHeadingElement {
	return HTMLHeadingElement{}.FromRef(
		bindings.NewHTMLHeadingElementByHTMLHeadingElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLHeadingElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLHeadingElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLHeadingElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHeadingElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLHeadingElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLHtmlElement() HTMLHtmlElement {
	return HTMLHtmlElement{}.FromRef(
		bindings.NewHTMLHtmlElementByHTMLHtmlElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLHtmlElement) Version() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLHtmlElementVersion(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Version sets the value of property "HTMLHtmlElement.version" to val.
//
// It returns false if the property cannot be set.
func (this HTMLHtmlElement) SetVersion(val js.String) bool {
	return js.True == bindings.SetHTMLHtmlElementVersion(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLIFrameElement() HTMLIFrameElement {
	return HTMLIFrameElement{}.FromRef(
		bindings.NewHTMLIFrameElementByHTMLIFrameElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src sets the value of property "HTMLIFrameElement.src" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) Srcdoc() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementSrcdoc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Srcdoc sets the value of property "HTMLIFrameElement.srcdoc" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLIFrameElement.name" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) Sandbox() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementSandbox(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// Allow returns the value of property "HTMLIFrameElement.allow".
//
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) Allow() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementAllow(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Allow sets the value of property "HTMLIFrameElement.allow" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) AllowFullscreen() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementAllowFullscreen(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// AllowFullscreen sets the value of property "HTMLIFrameElement.allowFullscreen" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) Width() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Width sets the value of property "HTMLIFrameElement.width" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) Height() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Height sets the value of property "HTMLIFrameElement.height" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) ReferrerPolicy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementReferrerPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReferrerPolicy sets the value of property "HTMLIFrameElement.referrerPolicy" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) Loading() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementLoading(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Loading sets the value of property "HTMLIFrameElement.loading" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) ContentDocument() (Document, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementContentDocument(
		this.Ref(), js.Pointer(&_ok),
	)
	return Document{}.FromRef(_ret), _ok
}

// ContentWindow returns the value of property "HTMLIFrameElement.contentWindow".
//
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) ContentWindow() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementContentWindow(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// PermissionsPolicy returns the value of property "HTMLIFrameElement.permissionsPolicy".
//
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) PermissionsPolicy() (PermissionsPolicy, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementPermissionsPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return PermissionsPolicy{}.FromRef(_ret), _ok
}

// Csp returns the value of property "HTMLIFrameElement.csp".
//
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) Csp() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementCsp(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Csp sets the value of property "HTMLIFrameElement.csp" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) PrivateToken() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementPrivateToken(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// PrivateToken sets the value of property "HTMLIFrameElement.privateToken" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLIFrameElement.align" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) Scrolling() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementScrolling(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Scrolling sets the value of property "HTMLIFrameElement.scrolling" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) FrameBorder() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementFrameBorder(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FrameBorder sets the value of property "HTMLIFrameElement.frameBorder" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) LongDesc() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementLongDesc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// LongDesc sets the value of property "HTMLIFrameElement.longDesc" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) MarginHeight() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementMarginHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// MarginHeight sets the value of property "HTMLIFrameElement.marginHeight" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLIFrameElement) MarginWidth() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLIFrameElementMarginWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// MarginWidth sets the value of property "HTMLIFrameElement.marginWidth" to val.
//
// It returns false if the property cannot be set.
func (this HTMLIFrameElement) SetMarginWidth(val js.String) bool {
	return js.True == bindings.SetHTMLIFrameElementMarginWidth(
		this.Ref(),
		val.Ref(),
	)
}

// GetSVGDocument calls the method "HTMLIFrameElement.getSVGDocument".
//
// The returned bool will be false if there is no such method.
func (this HTMLIFrameElement) GetSVGDocument() (Document, bool) {
	var _ok bool
	_ret := bindings.CallHTMLIFrameElementGetSVGDocument(
		this.Ref(), js.Pointer(&_ok),
	)

	return Document{}.FromRef(_ret), _ok
}

// GetSVGDocumentFunc returns the method "HTMLIFrameElement.getSVGDocument".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLIFrameElement) GetSVGDocumentFunc() (fn js.Func[func() Document]) {
	return fn.FromRef(
		bindings.HTMLIFrameElementGetSVGDocumentFunc(
			this.Ref(),
		),
	)
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

func NewHTMLInputElement() HTMLInputElement {
	return HTMLInputElement{}.FromRef(
		bindings.NewHTMLInputElementByHTMLInputElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Accept() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementAccept(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Accept sets the value of property "HTMLInputElement.accept" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Alt() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementAlt(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Alt sets the value of property "HTMLInputElement.alt" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Autocomplete() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementAutocomplete(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Autocomplete sets the value of property "HTMLInputElement.autocomplete" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) DefaultChecked() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementDefaultChecked(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// DefaultChecked sets the value of property "HTMLInputElement.defaultChecked" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Checked() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementChecked(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Checked sets the value of property "HTMLInputElement.checked" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) DirName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementDirName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DirName sets the value of property "HTMLInputElement.dirName" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Disabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Disabled sets the value of property "HTMLInputElement.disabled" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Form() (HTMLFormElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementForm(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormElement{}.FromRef(_ret), _ok
}

// Files returns the value of property "HTMLInputElement.files".
//
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Files() (FileList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementFiles(
		this.Ref(), js.Pointer(&_ok),
	)
	return FileList{}.FromRef(_ret), _ok
}

// Files sets the value of property "HTMLInputElement.files" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) FormAction() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementFormAction(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FormAction sets the value of property "HTMLInputElement.formAction" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) FormEnctype() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementFormEnctype(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FormEnctype sets the value of property "HTMLInputElement.formEnctype" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) FormMethod() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementFormMethod(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FormMethod sets the value of property "HTMLInputElement.formMethod" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) FormNoValidate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementFormNoValidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// FormNoValidate sets the value of property "HTMLInputElement.formNoValidate" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) FormTarget() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementFormTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FormTarget sets the value of property "HTMLInputElement.formTarget" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Height() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Height sets the value of property "HTMLInputElement.height" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Indeterminate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementIndeterminate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Indeterminate sets the value of property "HTMLInputElement.indeterminate" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) List() (HTMLDataListElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementList(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLDataListElement{}.FromRef(_ret), _ok
}

// Max returns the value of property "HTMLInputElement.max".
//
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Max() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementMax(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Max sets the value of property "HTMLInputElement.max" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) MaxLength() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementMaxLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// MaxLength sets the value of property "HTMLInputElement.maxLength" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Min() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementMin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Min sets the value of property "HTMLInputElement.min" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) MinLength() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementMinLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// MinLength sets the value of property "HTMLInputElement.minLength" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Multiple() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementMultiple(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Multiple sets the value of property "HTMLInputElement.multiple" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLInputElement.name" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Pattern() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementPattern(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Pattern sets the value of property "HTMLInputElement.pattern" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Placeholder() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementPlaceholder(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Placeholder sets the value of property "HTMLInputElement.placeholder" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) ReadOnly() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementReadOnly(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ReadOnly sets the value of property "HTMLInputElement.readOnly" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Required() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementRequired(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Required sets the value of property "HTMLInputElement.required" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Size() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Size sets the value of property "HTMLInputElement.size" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src sets the value of property "HTMLInputElement.src" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Step() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementStep(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Step sets the value of property "HTMLInputElement.step" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLInputElement.type" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) DefaultValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementDefaultValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DefaultValue sets the value of property "HTMLInputElement.defaultValue" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "HTMLInputElement.value" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) ValueAsDate() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementValueAsDate(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// ValueAsDate sets the value of property "HTMLInputElement.valueAsDate" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) ValueAsNumber() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementValueAsNumber(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// ValueAsNumber sets the value of property "HTMLInputElement.valueAsNumber" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Width() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Width sets the value of property "HTMLInputElement.width" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) WillValidate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementWillValidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Validity returns the value of property "HTMLInputElement.validity".
//
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Validity() (ValidityState, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementValidity(
		this.Ref(), js.Pointer(&_ok),
	)
	return ValidityState{}.FromRef(_ret), _ok
}

// ValidationMessage returns the value of property "HTMLInputElement.validationMessage".
//
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) ValidationMessage() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementValidationMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Labels returns the value of property "HTMLInputElement.labels".
//
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Labels() (NodeList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementLabels(
		this.Ref(), js.Pointer(&_ok),
	)
	return NodeList{}.FromRef(_ret), _ok
}

// SelectionStart returns the value of property "HTMLInputElement.selectionStart".
//
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) SelectionStart() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementSelectionStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// SelectionStart sets the value of property "HTMLInputElement.selectionStart" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) SelectionEnd() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementSelectionEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// SelectionEnd sets the value of property "HTMLInputElement.selectionEnd" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) SelectionDirection() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementSelectionDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SelectionDirection sets the value of property "HTMLInputElement.selectionDirection" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLInputElement.align" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) UseMap() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementUseMap(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// UseMap sets the value of property "HTMLInputElement.useMap" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Webkitdirectory() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementWebkitdirectory(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Webkitdirectory sets the value of property "HTMLInputElement.webkitdirectory" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) WebkitEntries() (js.FrozenArray[FileSystemEntry], bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementWebkitEntries(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.FrozenArray[FileSystemEntry]{}.FromRef(_ret), _ok
}

// Capture returns the value of property "HTMLInputElement.capture".
//
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) Capture() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementCapture(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Capture sets the value of property "HTMLInputElement.capture" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) PopoverTargetElement() (Element, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementPopoverTargetElement(
		this.Ref(), js.Pointer(&_ok),
	)
	return Element{}.FromRef(_ret), _ok
}

// PopoverTargetElement sets the value of property "HTMLInputElement.popoverTargetElement" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLInputElement) PopoverTargetAction() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLInputElementPopoverTargetAction(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// PopoverTargetAction sets the value of property "HTMLInputElement.popoverTargetAction" to val.
//
// It returns false if the property cannot be set.
func (this HTMLInputElement) SetPopoverTargetAction(val js.String) bool {
	return js.True == bindings.SetHTMLInputElementPopoverTargetAction(
		this.Ref(),
		val.Ref(),
	)
}

// StepUp calls the method "HTMLInputElement.stepUp".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) StepUp(n int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementStepUp(
		this.Ref(), js.Pointer(&_ok),
		int32(n),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StepUpFunc returns the method "HTMLInputElement.stepUp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) StepUpFunc() (fn js.Func[func(n int32)]) {
	return fn.FromRef(
		bindings.HTMLInputElementStepUpFunc(
			this.Ref(),
		),
	)
}

// StepUp1 calls the method "HTMLInputElement.stepUp".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) StepUp1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementStepUp1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StepUp1Func returns the method "HTMLInputElement.stepUp".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) StepUp1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLInputElementStepUp1Func(
			this.Ref(),
		),
	)
}

// StepDown calls the method "HTMLInputElement.stepDown".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) StepDown(n int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementStepDown(
		this.Ref(), js.Pointer(&_ok),
		int32(n),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StepDownFunc returns the method "HTMLInputElement.stepDown".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) StepDownFunc() (fn js.Func[func(n int32)]) {
	return fn.FromRef(
		bindings.HTMLInputElementStepDownFunc(
			this.Ref(),
		),
	)
}

// StepDown1 calls the method "HTMLInputElement.stepDown".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) StepDown1() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementStepDown1(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StepDown1Func returns the method "HTMLInputElement.stepDown".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) StepDown1Func() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLInputElementStepDown1Func(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLInputElement.checkValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) CheckValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementCheckValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// CheckValidityFunc returns the method "HTMLInputElement.checkValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLInputElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLInputElement.reportValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) ReportValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementReportValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// ReportValidityFunc returns the method "HTMLInputElement.reportValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLInputElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLInputElement.setCustomValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) SetCustomValidity(err js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementSetCustomValidity(
		this.Ref(), js.Pointer(&_ok),
		err.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCustomValidityFunc returns the method "HTMLInputElement.setCustomValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

// Select calls the method "HTMLInputElement.select".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) Select() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementSelect(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SelectFunc returns the method "HTMLInputElement.select".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) SelectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLInputElementSelectFunc(
			this.Ref(),
		),
	)
}

// SetRangeText calls the method "HTMLInputElement.setRangeText".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) SetRangeText(replacement js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementSetRangeText(
		this.Ref(), js.Pointer(&_ok),
		replacement.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetRangeTextFunc returns the method "HTMLInputElement.setRangeText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) SetRangeTextFunc() (fn js.Func[func(replacement js.String)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetRangeTextFunc(
			this.Ref(),
		),
	)
}

// SetRangeText1 calls the method "HTMLInputElement.setRangeText".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) SetRangeText1(replacement js.String, start uint32, end uint32, selectionMode SelectionMode) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementSetRangeText1(
		this.Ref(), js.Pointer(&_ok),
		replacement.Ref(),
		uint32(start),
		uint32(end),
		uint32(selectionMode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetRangeText1Func returns the method "HTMLInputElement.setRangeText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) SetRangeText1Func() (fn js.Func[func(replacement js.String, start uint32, end uint32, selectionMode SelectionMode)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetRangeText1Func(
			this.Ref(),
		),
	)
}

// SetRangeText2 calls the method "HTMLInputElement.setRangeText".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) SetRangeText2(replacement js.String, start uint32, end uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementSetRangeText2(
		this.Ref(), js.Pointer(&_ok),
		replacement.Ref(),
		uint32(start),
		uint32(end),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetRangeText2Func returns the method "HTMLInputElement.setRangeText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) SetRangeText2Func() (fn js.Func[func(replacement js.String, start uint32, end uint32)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetRangeText2Func(
			this.Ref(),
		),
	)
}

// SetSelectionRange calls the method "HTMLInputElement.setSelectionRange".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) SetSelectionRange(start uint32, end uint32, direction js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementSetSelectionRange(
		this.Ref(), js.Pointer(&_ok),
		uint32(start),
		uint32(end),
		direction.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetSelectionRangeFunc returns the method "HTMLInputElement.setSelectionRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) SetSelectionRangeFunc() (fn js.Func[func(start uint32, end uint32, direction js.String)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetSelectionRangeFunc(
			this.Ref(),
		),
	)
}

// SetSelectionRange1 calls the method "HTMLInputElement.setSelectionRange".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) SetSelectionRange1(start uint32, end uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementSetSelectionRange1(
		this.Ref(), js.Pointer(&_ok),
		uint32(start),
		uint32(end),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetSelectionRange1Func returns the method "HTMLInputElement.setSelectionRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) SetSelectionRange1Func() (fn js.Func[func(start uint32, end uint32)]) {
	return fn.FromRef(
		bindings.HTMLInputElementSetSelectionRange1Func(
			this.Ref(),
		),
	)
}

// ShowPicker calls the method "HTMLInputElement.showPicker".
//
// The returned bool will be false if there is no such method.
func (this HTMLInputElement) ShowPicker() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLInputElementShowPicker(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// ShowPickerFunc returns the method "HTMLInputElement.showPicker".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLInputElement) ShowPickerFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLInputElementShowPickerFunc(
			this.Ref(),
		),
	)
}

func NewHTMLLIElement() HTMLLIElement {
	return HTMLLIElement{}.FromRef(
		bindings.NewHTMLLIElementByHTMLLIElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLLIElement) Value() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLIElementValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Value sets the value of property "HTMLLIElement.value" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLIElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLIElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLLIElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLIElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLLIElementType(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLLabelElement() HTMLLabelElement {
	return HTMLLabelElement{}.FromRef(
		bindings.NewHTMLLabelElementByHTMLLabelElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLLabelElement) Form() (HTMLFormElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLabelElementForm(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormElement{}.FromRef(_ret), _ok
}

// HtmlFor returns the value of property "HTMLLabelElement.htmlFor".
//
// The returned bool will be false if there is no such property.
func (this HTMLLabelElement) HtmlFor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLabelElementHtmlFor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// HtmlFor sets the value of property "HTMLLabelElement.htmlFor" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLabelElement) Control() (HTMLElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLabelElementControl(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLElement{}.FromRef(_ret), _ok
}

func NewHTMLLegendElement() HTMLLegendElement {
	return HTMLLegendElement{}.FromRef(
		bindings.NewHTMLLegendElementByHTMLLegendElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLLegendElement) Form() (HTMLFormElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLegendElementForm(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormElement{}.FromRef(_ret), _ok
}

// Align returns the value of property "HTMLLegendElement.align".
//
// The returned bool will be false if there is no such property.
func (this HTMLLegendElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLegendElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLLegendElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLLegendElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLLegendElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLLinkElement() HTMLLinkElement {
	return HTMLLinkElement{}.FromRef(
		bindings.NewHTMLLinkElementByHTMLLinkElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Href() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementHref(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Href sets the value of property "HTMLLinkElement.href" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) CrossOrigin() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementCrossOrigin(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CrossOrigin sets the value of property "HTMLLinkElement.crossOrigin" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Rel() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementRel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Rel sets the value of property "HTMLLinkElement.rel" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) As() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementAs(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// As sets the value of property "HTMLLinkElement.as" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) RelList() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementRelList(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// Media returns the value of property "HTMLLinkElement.media".
//
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Media() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementMedia(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Media sets the value of property "HTMLLinkElement.media" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Integrity() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementIntegrity(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Integrity sets the value of property "HTMLLinkElement.integrity" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Hreflang() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementHreflang(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Hreflang sets the value of property "HTMLLinkElement.hreflang" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLLinkElement.type" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Sizes() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementSizes(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// ImageSrcset returns the value of property "HTMLLinkElement.imageSrcset".
//
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) ImageSrcset() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementImageSrcset(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ImageSrcset sets the value of property "HTMLLinkElement.imageSrcset" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) ImageSizes() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementImageSizes(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ImageSizes sets the value of property "HTMLLinkElement.imageSizes" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) ReferrerPolicy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementReferrerPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReferrerPolicy sets the value of property "HTMLLinkElement.referrerPolicy" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Blocking() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementBlocking(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// Disabled returns the value of property "HTMLLinkElement.disabled".
//
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Disabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Disabled sets the value of property "HTMLLinkElement.disabled" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) FetchPriority() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementFetchPriority(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// FetchPriority sets the value of property "HTMLLinkElement.fetchPriority" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Charset() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementCharset(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Charset sets the value of property "HTMLLinkElement.charset" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Rev() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementRev(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Rev sets the value of property "HTMLLinkElement.rev" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Target() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementTarget(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Target sets the value of property "HTMLLinkElement.target" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLLinkElement) Sheet() (CSSStyleSheet, bool) {
	var _ok bool
	_ret := bindings.GetHTMLLinkElementSheet(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleSheet{}.FromRef(_ret), _ok
}

func NewHTMLMapElement() HTMLMapElement {
	return HTMLMapElement{}.FromRef(
		bindings.NewHTMLMapElementByHTMLMapElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLMapElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMapElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLMapElement.name" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMapElement) Areas() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMapElementAreas(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

func NewHTMLMarqueeElement() HTMLMarqueeElement {
	return HTMLMarqueeElement{}.FromRef(
		bindings.NewHTMLMarqueeElementByHTMLMarqueeElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLMarqueeElement) Behavior() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMarqueeElementBehavior(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Behavior sets the value of property "HTMLMarqueeElement.behavior" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMarqueeElement) BgColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMarqueeElementBgColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BgColor sets the value of property "HTMLMarqueeElement.bgColor" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMarqueeElement) Direction() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMarqueeElementDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Direction sets the value of property "HTMLMarqueeElement.direction" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMarqueeElement) Height() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMarqueeElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Height sets the value of property "HTMLMarqueeElement.height" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMarqueeElement) Hspace() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMarqueeElementHspace(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Hspace sets the value of property "HTMLMarqueeElement.hspace" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMarqueeElement) Loop() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMarqueeElementLoop(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Loop sets the value of property "HTMLMarqueeElement.loop" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMarqueeElement) ScrollAmount() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMarqueeElementScrollAmount(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ScrollAmount sets the value of property "HTMLMarqueeElement.scrollAmount" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMarqueeElement) ScrollDelay() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMarqueeElementScrollDelay(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ScrollDelay sets the value of property "HTMLMarqueeElement.scrollDelay" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMarqueeElement) TrueSpeed() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMarqueeElementTrueSpeed(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// TrueSpeed sets the value of property "HTMLMarqueeElement.trueSpeed" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMarqueeElement) Vspace() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMarqueeElementVspace(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Vspace sets the value of property "HTMLMarqueeElement.vspace" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMarqueeElement) Width() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMarqueeElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Width sets the value of property "HTMLMarqueeElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMarqueeElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLMarqueeElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

// Start calls the method "HTMLMarqueeElement.start".
//
// The returned bool will be false if there is no such method.
func (this HTMLMarqueeElement) Start() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLMarqueeElementStart(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StartFunc returns the method "HTMLMarqueeElement.start".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMarqueeElement) StartFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLMarqueeElementStartFunc(
			this.Ref(),
		),
	)
}

// Stop calls the method "HTMLMarqueeElement.stop".
//
// The returned bool will be false if there is no such method.
func (this HTMLMarqueeElement) Stop() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLMarqueeElementStop(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// StopFunc returns the method "HTMLMarqueeElement.stop".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLMarqueeElement) StopFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLMarqueeElementStopFunc(
			this.Ref(),
		),
	)
}

func NewHTMLMenuElement() HTMLMenuElement {
	return HTMLMenuElement{}.FromRef(
		bindings.NewHTMLMenuElementByHTMLMenuElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLMenuElement) Compact() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMenuElementCompact(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Compact sets the value of property "HTMLMenuElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMenuElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLMenuElementCompact(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

func NewHTMLMetaElement() HTMLMetaElement {
	return HTMLMetaElement{}.FromRef(
		bindings.NewHTMLMetaElementByHTMLMetaElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLMetaElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMetaElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLMetaElement.name" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMetaElement) HttpEquiv() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMetaElementHttpEquiv(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// HttpEquiv sets the value of property "HTMLMetaElement.httpEquiv" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMetaElement) Content() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMetaElementContent(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Content sets the value of property "HTMLMetaElement.content" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMetaElement) Media() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMetaElementMedia(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Media sets the value of property "HTMLMetaElement.media" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMetaElement) Scheme() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMetaElementScheme(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Scheme sets the value of property "HTMLMetaElement.scheme" to val.
//
// It returns false if the property cannot be set.
func (this HTMLMetaElement) SetScheme(val js.String) bool {
	return js.True == bindings.SetHTMLMetaElementScheme(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLMeterElement() HTMLMeterElement {
	return HTMLMeterElement{}.FromRef(
		bindings.NewHTMLMeterElementByHTMLMeterElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLMeterElement) Value() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMeterElementValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Value sets the value of property "HTMLMeterElement.value" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMeterElement) Min() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMeterElementMin(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Min sets the value of property "HTMLMeterElement.min" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMeterElement) Max() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMeterElementMax(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Max sets the value of property "HTMLMeterElement.max" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMeterElement) Low() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMeterElementLow(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Low sets the value of property "HTMLMeterElement.low" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMeterElement) High() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMeterElementHigh(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// High sets the value of property "HTMLMeterElement.high" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMeterElement) Optimum() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMeterElementOptimum(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Optimum sets the value of property "HTMLMeterElement.optimum" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLMeterElement) Labels() (NodeList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLMeterElementLabels(
		this.Ref(), js.Pointer(&_ok),
	)
	return NodeList{}.FromRef(_ret), _ok
}

func NewHTMLModElement() HTMLModElement {
	return HTMLModElement{}.FromRef(
		bindings.NewHTMLModElementByHTMLModElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLModElement) Cite() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLModElementCite(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Cite sets the value of property "HTMLModElement.cite" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLModElement) DateTime() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLModElementDateTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DateTime sets the value of property "HTMLModElement.dateTime" to val.
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

func NewHTMLOListElement() HTMLOListElement {
	return HTMLOListElement{}.FromRef(
		bindings.NewHTMLOListElementByHTMLOListElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLOListElement) Reversed() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOListElementReversed(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Reversed sets the value of property "HTMLOListElement.reversed" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOListElement) Start() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOListElementStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Start sets the value of property "HTMLOListElement.start" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOListElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOListElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLOListElement.type" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOListElement) Compact() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOListElementCompact(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Compact sets the value of property "HTMLOListElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOListElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLOListElementCompact(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

func NewHTMLObjectElement() HTMLObjectElement {
	return HTMLObjectElement{}.FromRef(
		bindings.NewHTMLObjectElementByHTMLObjectElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Data() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementData(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Data sets the value of property "HTMLObjectElement.data" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLObjectElement.type" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLObjectElement.name" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Form() (HTMLFormElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementForm(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormElement{}.FromRef(_ret), _ok
}

// Width returns the value of property "HTMLObjectElement.width".
//
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Width() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Width sets the value of property "HTMLObjectElement.width" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Height() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Height sets the value of property "HTMLObjectElement.height" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) ContentDocument() (Document, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementContentDocument(
		this.Ref(), js.Pointer(&_ok),
	)
	return Document{}.FromRef(_ret), _ok
}

// ContentWindow returns the value of property "HTMLObjectElement.contentWindow".
//
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) ContentWindow() (js.Object, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementContentWindow(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Object{}.FromRef(_ret), _ok
}

// WillValidate returns the value of property "HTMLObjectElement.willValidate".
//
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) WillValidate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementWillValidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Validity returns the value of property "HTMLObjectElement.validity".
//
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Validity() (ValidityState, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementValidity(
		this.Ref(), js.Pointer(&_ok),
	)
	return ValidityState{}.FromRef(_ret), _ok
}

// ValidationMessage returns the value of property "HTMLObjectElement.validationMessage".
//
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) ValidationMessage() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementValidationMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align returns the value of property "HTMLObjectElement.align".
//
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLObjectElement.align" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Archive() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementArchive(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Archive sets the value of property "HTMLObjectElement.archive" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Code() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementCode(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Code sets the value of property "HTMLObjectElement.code" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Declare() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementDeclare(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Declare sets the value of property "HTMLObjectElement.declare" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Hspace() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementHspace(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Hspace sets the value of property "HTMLObjectElement.hspace" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Standby() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementStandby(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Standby sets the value of property "HTMLObjectElement.standby" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Vspace() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementVspace(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Vspace sets the value of property "HTMLObjectElement.vspace" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) CodeBase() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementCodeBase(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CodeBase sets the value of property "HTMLObjectElement.codeBase" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) CodeType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementCodeType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CodeType sets the value of property "HTMLObjectElement.codeType" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) UseMap() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementUseMap(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// UseMap sets the value of property "HTMLObjectElement.useMap" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLObjectElement) Border() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLObjectElementBorder(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Border sets the value of property "HTMLObjectElement.border" to val.
//
// It returns false if the property cannot be set.
func (this HTMLObjectElement) SetBorder(val js.String) bool {
	return js.True == bindings.SetHTMLObjectElementBorder(
		this.Ref(),
		val.Ref(),
	)
}

// GetSVGDocument calls the method "HTMLObjectElement.getSVGDocument".
//
// The returned bool will be false if there is no such method.
func (this HTMLObjectElement) GetSVGDocument() (Document, bool) {
	var _ok bool
	_ret := bindings.CallHTMLObjectElementGetSVGDocument(
		this.Ref(), js.Pointer(&_ok),
	)

	return Document{}.FromRef(_ret), _ok
}

// GetSVGDocumentFunc returns the method "HTMLObjectElement.getSVGDocument".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLObjectElement) GetSVGDocumentFunc() (fn js.Func[func() Document]) {
	return fn.FromRef(
		bindings.HTMLObjectElementGetSVGDocumentFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLObjectElement.checkValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLObjectElement) CheckValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLObjectElementCheckValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// CheckValidityFunc returns the method "HTMLObjectElement.checkValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLObjectElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLObjectElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLObjectElement.reportValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLObjectElement) ReportValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLObjectElementReportValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// ReportValidityFunc returns the method "HTMLObjectElement.reportValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLObjectElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLObjectElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLObjectElement.setCustomValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLObjectElement) SetCustomValidity(err js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLObjectElementSetCustomValidity(
		this.Ref(), js.Pointer(&_ok),
		err.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCustomValidityFunc returns the method "HTMLObjectElement.setCustomValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLObjectElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLObjectElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

func NewHTMLOptGroupElement() HTMLOptGroupElement {
	return HTMLOptGroupElement{}.FromRef(
		bindings.NewHTMLOptGroupElementByHTMLOptGroupElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLOptGroupElement) Disabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptGroupElementDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Disabled sets the value of property "HTMLOptGroupElement.disabled" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOptGroupElement) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptGroupElementLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "HTMLOptGroupElement.label" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptGroupElement) SetLabel(val js.String) bool {
	return js.True == bindings.SetHTMLOptGroupElementLabel(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLOptionElement() HTMLOptionElement {
	return HTMLOptionElement{}.FromRef(
		bindings.NewHTMLOptionElementByHTMLOptionElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLOptionElement) Disabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptionElementDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Disabled sets the value of property "HTMLOptionElement.disabled" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOptionElement) Form() (HTMLFormElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptionElementForm(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormElement{}.FromRef(_ret), _ok
}

// Label returns the value of property "HTMLOptionElement.label".
//
// The returned bool will be false if there is no such property.
func (this HTMLOptionElement) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptionElementLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "HTMLOptionElement.label" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOptionElement) DefaultSelected() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptionElementDefaultSelected(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// DefaultSelected sets the value of property "HTMLOptionElement.defaultSelected" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOptionElement) Selected() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptionElementSelected(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Selected sets the value of property "HTMLOptionElement.selected" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOptionElement) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptionElementValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "HTMLOptionElement.value" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOptionElement) Text() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptionElementText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Text sets the value of property "HTMLOptionElement.text" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOptionElement) Index() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptionElementIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
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
// The returned bool will be false if there is no such property.
func (this HTMLOptionsCollection) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptionsCollectionLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Length sets the value of property "HTMLOptionsCollection.length" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOptionsCollection) SelectedIndex() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOptionsCollectionSelectedIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// SelectedIndex sets the value of property "HTMLOptionsCollection.selectedIndex" to val.
//
// It returns false if the property cannot be set.
func (this HTMLOptionsCollection) SetSelectedIndex(val int32) bool {
	return js.True == bindings.SetHTMLOptionsCollectionSelectedIndex(
		this.Ref(),
		int32(val),
	)
}

// Set calls the method "HTMLOptionsCollection.".
//
// The returned bool will be false if there is no such method.
func (this HTMLOptionsCollection) Set(index uint32, option HTMLOptionElement) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLOptionsCollectionSet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		option.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "HTMLOptionsCollection.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLOptionsCollection) SetFunc() (fn js.Func[func(index uint32, option HTMLOptionElement)]) {
	return fn.FromRef(
		bindings.HTMLOptionsCollectionSetFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "HTMLOptionsCollection.add".
//
// The returned bool will be false if there is no such method.
func (this HTMLOptionsCollection) Add(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLOptionsCollectionAdd(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
		before.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddFunc returns the method "HTMLOptionsCollection.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLOptionsCollection) AddFunc() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32)]) {
	return fn.FromRef(
		bindings.HTMLOptionsCollectionAddFunc(
			this.Ref(),
		),
	)
}

// Add1 calls the method "HTMLOptionsCollection.add".
//
// The returned bool will be false if there is no such method.
func (this HTMLOptionsCollection) Add1(element OneOf_HTMLOptionElement_HTMLOptGroupElement) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLOptionsCollectionAdd1(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Add1Func returns the method "HTMLOptionsCollection.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLOptionsCollection) Add1Func() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement)]) {
	return fn.FromRef(
		bindings.HTMLOptionsCollectionAdd1Func(
			this.Ref(),
		),
	)
}

// Remove calls the method "HTMLOptionsCollection.remove".
//
// The returned bool will be false if there is no such method.
func (this HTMLOptionsCollection) Remove(index int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLOptionsCollectionRemove(
		this.Ref(), js.Pointer(&_ok),
		int32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveFunc returns the method "HTMLOptionsCollection.remove".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLOptionsCollection) RemoveFunc() (fn js.Func[func(index int32)]) {
	return fn.FromRef(
		bindings.HTMLOptionsCollectionRemoveFunc(
			this.Ref(),
		),
	)
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

func NewHTMLOutputElement() HTMLOutputElement {
	return HTMLOutputElement{}.FromRef(
		bindings.NewHTMLOutputElementByHTMLOutputElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLOutputElement) HtmlFor() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOutputElementHtmlFor(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// Form returns the value of property "HTMLOutputElement.form".
//
// The returned bool will be false if there is no such property.
func (this HTMLOutputElement) Form() (HTMLFormElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOutputElementForm(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormElement{}.FromRef(_ret), _ok
}

// Name returns the value of property "HTMLOutputElement.name".
//
// The returned bool will be false if there is no such property.
func (this HTMLOutputElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOutputElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLOutputElement.name" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOutputElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOutputElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DefaultValue returns the value of property "HTMLOutputElement.defaultValue".
//
// The returned bool will be false if there is no such property.
func (this HTMLOutputElement) DefaultValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOutputElementDefaultValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DefaultValue sets the value of property "HTMLOutputElement.defaultValue" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOutputElement) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOutputElementValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "HTMLOutputElement.value" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLOutputElement) WillValidate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOutputElementWillValidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Validity returns the value of property "HTMLOutputElement.validity".
//
// The returned bool will be false if there is no such property.
func (this HTMLOutputElement) Validity() (ValidityState, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOutputElementValidity(
		this.Ref(), js.Pointer(&_ok),
	)
	return ValidityState{}.FromRef(_ret), _ok
}

// ValidationMessage returns the value of property "HTMLOutputElement.validationMessage".
//
// The returned bool will be false if there is no such property.
func (this HTMLOutputElement) ValidationMessage() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOutputElementValidationMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Labels returns the value of property "HTMLOutputElement.labels".
//
// The returned bool will be false if there is no such property.
func (this HTMLOutputElement) Labels() (NodeList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLOutputElementLabels(
		this.Ref(), js.Pointer(&_ok),
	)
	return NodeList{}.FromRef(_ret), _ok
}

// CheckValidity calls the method "HTMLOutputElement.checkValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLOutputElement) CheckValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLOutputElementCheckValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// CheckValidityFunc returns the method "HTMLOutputElement.checkValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLOutputElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLOutputElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLOutputElement.reportValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLOutputElement) ReportValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLOutputElementReportValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// ReportValidityFunc returns the method "HTMLOutputElement.reportValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLOutputElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLOutputElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLOutputElement.setCustomValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLOutputElement) SetCustomValidity(err js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLOutputElementSetCustomValidity(
		this.Ref(), js.Pointer(&_ok),
		err.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCustomValidityFunc returns the method "HTMLOutputElement.setCustomValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLOutputElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLOutputElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

func NewHTMLParagraphElement() HTMLParagraphElement {
	return HTMLParagraphElement{}.FromRef(
		bindings.NewHTMLParagraphElementByHTMLParagraphElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLParagraphElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLParagraphElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLParagraphElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParagraphElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLParagraphElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLParamElement() HTMLParamElement {
	return HTMLParamElement{}.FromRef(
		bindings.NewHTMLParamElementByHTMLParamElement(),
	)
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
// The returned bool will be false if there is no such property.
func (this HTMLParamElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLParamElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLParamElement.name" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLParamElement) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLParamElementValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "HTMLParamElement.value" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLParamElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLParamElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLParamElement.type" to val.
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
// The returned bool will be false if there is no such property.
func (this HTMLParamElement) ValueType() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLParamElementValueType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ValueType sets the value of property "HTMLParamElement.valueType" to val.
//
// It returns false if the property cannot be set.
func (this HTMLParamElement) SetValueType(val js.String) bool {
	return js.True == bindings.SetHTMLParamElementValueType(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLPictureElement() HTMLPictureElement {
	return HTMLPictureElement{}.FromRef(
		bindings.NewHTMLPictureElementByHTMLPictureElement(),
	)
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

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

func NewHTMLPortalElement() HTMLPortalElement {
	return HTMLPortalElement{}.FromRef(
		bindings.NewHTMLPortalElementByHTMLPortalElement(),
	)
}

type HTMLPortalElement struct {
	HTMLElement
}

func (this HTMLPortalElement) Once() HTMLPortalElement {
	this.Ref().Once()
	return this
}

func (this HTMLPortalElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLPortalElement) FromRef(ref js.Ref) HTMLPortalElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLPortalElement) Free() {
	this.Ref().Free()
}

// Src returns the value of property "HTMLPortalElement.src".
//
// The returned bool will be false if there is no such property.
func (this HTMLPortalElement) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLPortalElementSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src sets the value of property "HTMLPortalElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLPortalElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLPortalElementSrc(
		this.Ref(),
		val.Ref(),
	)
}

// ReferrerPolicy returns the value of property "HTMLPortalElement.referrerPolicy".
//
// The returned bool will be false if there is no such property.
func (this HTMLPortalElement) ReferrerPolicy() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLPortalElementReferrerPolicy(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ReferrerPolicy sets the value of property "HTMLPortalElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLPortalElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLPortalElementReferrerPolicy(
		this.Ref(),
		val.Ref(),
	)
}

// Activate calls the method "HTMLPortalElement.activate".
//
// The returned bool will be false if there is no such method.
func (this HTMLPortalElement) Activate(options PortalActivateOptions) (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallHTMLPortalElementActivate(
		this.Ref(), js.Pointer(&_ok),
		js.Pointer(&options),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// ActivateFunc returns the method "HTMLPortalElement.activate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLPortalElement) ActivateFunc() (fn js.Func[func(options PortalActivateOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLPortalElementActivateFunc(
			this.Ref(),
		),
	)
}

// Activate1 calls the method "HTMLPortalElement.activate".
//
// The returned bool will be false if there is no such method.
func (this HTMLPortalElement) Activate1() (js.Promise[js.Void], bool) {
	var _ok bool
	_ret := bindings.CallHTMLPortalElementActivate1(
		this.Ref(), js.Pointer(&_ok),
	)

	return js.Promise[js.Void]{}.FromRef(_ret), _ok
}

// Activate1Func returns the method "HTMLPortalElement.activate".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLPortalElement) Activate1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLPortalElementActivate1Func(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "HTMLPortalElement.postMessage".
//
// The returned bool will be false if there is no such method.
func (this HTMLPortalElement) PostMessage(message js.Any, options StructuredSerializeOptions) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLPortalElementPostMessage(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
		js.Pointer(&options),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessageFunc returns the method "HTMLPortalElement.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLPortalElement) PostMessageFunc() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.HTMLPortalElementPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "HTMLPortalElement.postMessage".
//
// The returned bool will be false if there is no such method.
func (this HTMLPortalElement) PostMessage1(message js.Any) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLPortalElementPostMessage1(
		this.Ref(), js.Pointer(&_ok),
		message.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// PostMessage1Func returns the method "HTMLPortalElement.postMessage".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLPortalElement) PostMessage1Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.HTMLPortalElementPostMessage1Func(
			this.Ref(),
		),
	)
}

func NewHTMLPreElement() HTMLPreElement {
	return HTMLPreElement{}.FromRef(
		bindings.NewHTMLPreElementByHTMLPreElement(),
	)
}

type HTMLPreElement struct {
	HTMLElement
}

func (this HTMLPreElement) Once() HTMLPreElement {
	this.Ref().Once()
	return this
}

func (this HTMLPreElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLPreElement) FromRef(ref js.Ref) HTMLPreElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLPreElement) Free() {
	this.Ref().Free()
}

// Width returns the value of property "HTMLPreElement.width".
//
// The returned bool will be false if there is no such property.
func (this HTMLPreElement) Width() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLPreElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Width sets the value of property "HTMLPreElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLPreElement) SetWidth(val int32) bool {
	return js.True == bindings.SetHTMLPreElementWidth(
		this.Ref(),
		int32(val),
	)
}

func NewHTMLProgressElement() HTMLProgressElement {
	return HTMLProgressElement{}.FromRef(
		bindings.NewHTMLProgressElementByHTMLProgressElement(),
	)
}

type HTMLProgressElement struct {
	HTMLElement
}

func (this HTMLProgressElement) Once() HTMLProgressElement {
	this.Ref().Once()
	return this
}

func (this HTMLProgressElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLProgressElement) FromRef(ref js.Ref) HTMLProgressElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLProgressElement) Free() {
	this.Ref().Free()
}

// Value returns the value of property "HTMLProgressElement.value".
//
// The returned bool will be false if there is no such property.
func (this HTMLProgressElement) Value() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLProgressElementValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Value sets the value of property "HTMLProgressElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLProgressElement) SetValue(val float64) bool {
	return js.True == bindings.SetHTMLProgressElementValue(
		this.Ref(),
		float64(val),
	)
}

// Max returns the value of property "HTMLProgressElement.max".
//
// The returned bool will be false if there is no such property.
func (this HTMLProgressElement) Max() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLProgressElementMax(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Max sets the value of property "HTMLProgressElement.max" to val.
//
// It returns false if the property cannot be set.
func (this HTMLProgressElement) SetMax(val float64) bool {
	return js.True == bindings.SetHTMLProgressElementMax(
		this.Ref(),
		float64(val),
	)
}

// Position returns the value of property "HTMLProgressElement.position".
//
// The returned bool will be false if there is no such property.
func (this HTMLProgressElement) Position() (float64, bool) {
	var _ok bool
	_ret := bindings.GetHTMLProgressElementPosition(
		this.Ref(), js.Pointer(&_ok),
	)
	return float64(_ret), _ok
}

// Labels returns the value of property "HTMLProgressElement.labels".
//
// The returned bool will be false if there is no such property.
func (this HTMLProgressElement) Labels() (NodeList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLProgressElementLabels(
		this.Ref(), js.Pointer(&_ok),
	)
	return NodeList{}.FromRef(_ret), _ok
}

func NewHTMLQuoteElement() HTMLQuoteElement {
	return HTMLQuoteElement{}.FromRef(
		bindings.NewHTMLQuoteElementByHTMLQuoteElement(),
	)
}

type HTMLQuoteElement struct {
	HTMLElement
}

func (this HTMLQuoteElement) Once() HTMLQuoteElement {
	this.Ref().Once()
	return this
}

func (this HTMLQuoteElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLQuoteElement) FromRef(ref js.Ref) HTMLQuoteElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLQuoteElement) Free() {
	this.Ref().Free()
}

// Cite returns the value of property "HTMLQuoteElement.cite".
//
// The returned bool will be false if there is no such property.
func (this HTMLQuoteElement) Cite() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLQuoteElementCite(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Cite sets the value of property "HTMLQuoteElement.cite" to val.
//
// It returns false if the property cannot be set.
func (this HTMLQuoteElement) SetCite(val js.String) bool {
	return js.True == bindings.SetHTMLQuoteElementCite(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLSelectElement() HTMLSelectElement {
	return HTMLSelectElement{}.FromRef(
		bindings.NewHTMLSelectElementByHTMLSelectElement(),
	)
}

type HTMLSelectElement struct {
	HTMLElement
}

func (this HTMLSelectElement) Once() HTMLSelectElement {
	this.Ref().Once()
	return this
}

func (this HTMLSelectElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLSelectElement) FromRef(ref js.Ref) HTMLSelectElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLSelectElement) Free() {
	this.Ref().Free()
}

// Autocomplete returns the value of property "HTMLSelectElement.autocomplete".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Autocomplete() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementAutocomplete(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Autocomplete sets the value of property "HTMLSelectElement.autocomplete" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetAutocomplete(val js.String) bool {
	return js.True == bindings.SetHTMLSelectElementAutocomplete(
		this.Ref(),
		val.Ref(),
	)
}

// Disabled returns the value of property "HTMLSelectElement.disabled".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Disabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Disabled sets the value of property "HTMLSelectElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLSelectElementDisabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLSelectElement.form".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Form() (HTMLFormElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementForm(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormElement{}.FromRef(_ret), _ok
}

// Multiple returns the value of property "HTMLSelectElement.multiple".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Multiple() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementMultiple(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Multiple sets the value of property "HTMLSelectElement.multiple" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetMultiple(val bool) bool {
	return js.True == bindings.SetHTMLSelectElementMultiple(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Name returns the value of property "HTMLSelectElement.name".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLSelectElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLSelectElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Required returns the value of property "HTMLSelectElement.required".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Required() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementRequired(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Required sets the value of property "HTMLSelectElement.required" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetRequired(val bool) bool {
	return js.True == bindings.SetHTMLSelectElementRequired(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Size returns the value of property "HTMLSelectElement.size".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Size() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementSize(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Size sets the value of property "HTMLSelectElement.size" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetSize(val uint32) bool {
	return js.True == bindings.SetHTMLSelectElementSize(
		this.Ref(),
		uint32(val),
	)
}

// Type returns the value of property "HTMLSelectElement.type".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Options returns the value of property "HTMLSelectElement.options".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Options() (HTMLOptionsCollection, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementOptions(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLOptionsCollection{}.FromRef(_ret), _ok
}

// Length returns the value of property "HTMLSelectElement.length".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Length() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Length sets the value of property "HTMLSelectElement.length" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetLength(val uint32) bool {
	return js.True == bindings.SetHTMLSelectElementLength(
		this.Ref(),
		uint32(val),
	)
}

// SelectedOptions returns the value of property "HTMLSelectElement.selectedOptions".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) SelectedOptions() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementSelectedOptions(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// SelectedIndex returns the value of property "HTMLSelectElement.selectedIndex".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) SelectedIndex() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementSelectedIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// SelectedIndex sets the value of property "HTMLSelectElement.selectedIndex" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetSelectedIndex(val int32) bool {
	return js.True == bindings.SetHTMLSelectElementSelectedIndex(
		this.Ref(),
		int32(val),
	)
}

// Value returns the value of property "HTMLSelectElement.value".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "HTMLSelectElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLSelectElementValue(
		this.Ref(),
		val.Ref(),
	)
}

// WillValidate returns the value of property "HTMLSelectElement.willValidate".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) WillValidate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementWillValidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Validity returns the value of property "HTMLSelectElement.validity".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Validity() (ValidityState, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementValidity(
		this.Ref(), js.Pointer(&_ok),
	)
	return ValidityState{}.FromRef(_ret), _ok
}

// ValidationMessage returns the value of property "HTMLSelectElement.validationMessage".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) ValidationMessage() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementValidationMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Labels returns the value of property "HTMLSelectElement.labels".
//
// The returned bool will be false if there is no such property.
func (this HTMLSelectElement) Labels() (NodeList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSelectElementLabels(
		this.Ref(), js.Pointer(&_ok),
	)
	return NodeList{}.FromRef(_ret), _ok
}

// Item calls the method "HTMLSelectElement.item".
//
// The returned bool will be false if there is no such method.
func (this HTMLSelectElement) Item(index uint32) (HTMLOptionElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLSelectElementItem(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
	)

	return HTMLOptionElement{}.FromRef(_ret), _ok
}

// ItemFunc returns the method "HTMLSelectElement.item".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSelectElement) ItemFunc() (fn js.Func[func(index uint32) HTMLOptionElement]) {
	return fn.FromRef(
		bindings.HTMLSelectElementItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "HTMLSelectElement.namedItem".
//
// The returned bool will be false if there is no such method.
func (this HTMLSelectElement) NamedItem(name js.String) (HTMLOptionElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLSelectElementNamedItem(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return HTMLOptionElement{}.FromRef(_ret), _ok
}

// NamedItemFunc returns the method "HTMLSelectElement.namedItem".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSelectElement) NamedItemFunc() (fn js.Func[func(name js.String) HTMLOptionElement]) {
	return fn.FromRef(
		bindings.HTMLSelectElementNamedItemFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "HTMLSelectElement.add".
//
// The returned bool will be false if there is no such method.
func (this HTMLSelectElement) Add(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLSelectElementAdd(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
		before.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AddFunc returns the method "HTMLSelectElement.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSelectElement) AddFunc() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32)]) {
	return fn.FromRef(
		bindings.HTMLSelectElementAddFunc(
			this.Ref(),
		),
	)
}

// Add1 calls the method "HTMLSelectElement.add".
//
// The returned bool will be false if there is no such method.
func (this HTMLSelectElement) Add1(element OneOf_HTMLOptionElement_HTMLOptGroupElement) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLSelectElementAdd1(
		this.Ref(), js.Pointer(&_ok),
		element.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Add1Func returns the method "HTMLSelectElement.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSelectElement) Add1Func() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement)]) {
	return fn.FromRef(
		bindings.HTMLSelectElementAdd1Func(
			this.Ref(),
		),
	)
}

// Remove calls the method "HTMLSelectElement.remove".
//
// The returned bool will be false if there is no such method.
func (this HTMLSelectElement) Remove() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLSelectElementRemove(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// RemoveFunc returns the method "HTMLSelectElement.remove".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSelectElement) RemoveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLSelectElementRemoveFunc(
			this.Ref(),
		),
	)
}

// Remove1 calls the method "HTMLSelectElement.remove".
//
// The returned bool will be false if there is no such method.
func (this HTMLSelectElement) Remove1(index int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLSelectElementRemove1(
		this.Ref(), js.Pointer(&_ok),
		int32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// Remove1Func returns the method "HTMLSelectElement.remove".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSelectElement) Remove1Func() (fn js.Func[func(index int32)]) {
	return fn.FromRef(
		bindings.HTMLSelectElementRemove1Func(
			this.Ref(),
		),
	)
}

// Set calls the method "HTMLSelectElement.".
//
// The returned bool will be false if there is no such method.
func (this HTMLSelectElement) Set(index uint32, option HTMLOptionElement) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLSelectElementSet(
		this.Ref(), js.Pointer(&_ok),
		uint32(index),
		option.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetFunc returns the method "HTMLSelectElement.".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSelectElement) SetFunc() (fn js.Func[func(index uint32, option HTMLOptionElement)]) {
	return fn.FromRef(
		bindings.HTMLSelectElementSetFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLSelectElement.checkValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLSelectElement) CheckValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLSelectElementCheckValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// CheckValidityFunc returns the method "HTMLSelectElement.checkValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSelectElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLSelectElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLSelectElement.reportValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLSelectElement) ReportValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLSelectElementReportValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// ReportValidityFunc returns the method "HTMLSelectElement.reportValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSelectElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLSelectElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLSelectElement.setCustomValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLSelectElement) SetCustomValidity(err js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLSelectElementSetCustomValidity(
		this.Ref(), js.Pointer(&_ok),
		err.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCustomValidityFunc returns the method "HTMLSelectElement.setCustomValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLSelectElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLSelectElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

func NewHTMLSourceElement() HTMLSourceElement {
	return HTMLSourceElement{}.FromRef(
		bindings.NewHTMLSourceElementByHTMLSourceElement(),
	)
}

type HTMLSourceElement struct {
	HTMLElement
}

func (this HTMLSourceElement) Once() HTMLSourceElement {
	this.Ref().Once()
	return this
}

func (this HTMLSourceElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLSourceElement) FromRef(ref js.Ref) HTMLSourceElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLSourceElement) Free() {
	this.Ref().Free()
}

// Src returns the value of property "HTMLSourceElement.src".
//
// The returned bool will be false if there is no such property.
func (this HTMLSourceElement) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSourceElementSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src sets the value of property "HTMLSourceElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLSourceElementSrc(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLSourceElement.type".
//
// The returned bool will be false if there is no such property.
func (this HTMLSourceElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSourceElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLSourceElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLSourceElementType(
		this.Ref(),
		val.Ref(),
	)
}

// Srcset returns the value of property "HTMLSourceElement.srcset".
//
// The returned bool will be false if there is no such property.
func (this HTMLSourceElement) Srcset() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSourceElementSrcset(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Srcset sets the value of property "HTMLSourceElement.srcset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetSrcset(val js.String) bool {
	return js.True == bindings.SetHTMLSourceElementSrcset(
		this.Ref(),
		val.Ref(),
	)
}

// Sizes returns the value of property "HTMLSourceElement.sizes".
//
// The returned bool will be false if there is no such property.
func (this HTMLSourceElement) Sizes() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSourceElementSizes(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Sizes sets the value of property "HTMLSourceElement.sizes" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetSizes(val js.String) bool {
	return js.True == bindings.SetHTMLSourceElementSizes(
		this.Ref(),
		val.Ref(),
	)
}

// Media returns the value of property "HTMLSourceElement.media".
//
// The returned bool will be false if there is no such property.
func (this HTMLSourceElement) Media() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSourceElementMedia(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Media sets the value of property "HTMLSourceElement.media" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetMedia(val js.String) bool {
	return js.True == bindings.SetHTMLSourceElementMedia(
		this.Ref(),
		val.Ref(),
	)
}

// Width returns the value of property "HTMLSourceElement.width".
//
// The returned bool will be false if there is no such property.
func (this HTMLSourceElement) Width() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSourceElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Width sets the value of property "HTMLSourceElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetWidth(val uint32) bool {
	return js.True == bindings.SetHTMLSourceElementWidth(
		this.Ref(),
		uint32(val),
	)
}

// Height returns the value of property "HTMLSourceElement.height".
//
// The returned bool will be false if there is no such property.
func (this HTMLSourceElement) Height() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLSourceElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Height sets the value of property "HTMLSourceElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLSourceElementHeight(
		this.Ref(),
		uint32(val),
	)
}

func NewHTMLSpanElement() HTMLSpanElement {
	return HTMLSpanElement{}.FromRef(
		bindings.NewHTMLSpanElementByHTMLSpanElement(),
	)
}

type HTMLSpanElement struct {
	HTMLElement
}

func (this HTMLSpanElement) Once() HTMLSpanElement {
	this.Ref().Once()
	return this
}

func (this HTMLSpanElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLSpanElement) FromRef(ref js.Ref) HTMLSpanElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLSpanElement) Free() {
	this.Ref().Free()
}

type HTMLString = js.String

func NewHTMLStyleElement() HTMLStyleElement {
	return HTMLStyleElement{}.FromRef(
		bindings.NewHTMLStyleElementByHTMLStyleElement(),
	)
}

type HTMLStyleElement struct {
	HTMLElement
}

func (this HTMLStyleElement) Once() HTMLStyleElement {
	this.Ref().Once()
	return this
}

func (this HTMLStyleElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLStyleElement) FromRef(ref js.Ref) HTMLStyleElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLStyleElement) Free() {
	this.Ref().Free()
}

// Disabled returns the value of property "HTMLStyleElement.disabled".
//
// The returned bool will be false if there is no such property.
func (this HTMLStyleElement) Disabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLStyleElementDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Disabled sets the value of property "HTMLStyleElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLStyleElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLStyleElementDisabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Media returns the value of property "HTMLStyleElement.media".
//
// The returned bool will be false if there is no such property.
func (this HTMLStyleElement) Media() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLStyleElementMedia(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Media sets the value of property "HTMLStyleElement.media" to val.
//
// It returns false if the property cannot be set.
func (this HTMLStyleElement) SetMedia(val js.String) bool {
	return js.True == bindings.SetHTMLStyleElementMedia(
		this.Ref(),
		val.Ref(),
	)
}

// Blocking returns the value of property "HTMLStyleElement.blocking".
//
// The returned bool will be false if there is no such property.
func (this HTMLStyleElement) Blocking() (DOMTokenList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLStyleElementBlocking(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMTokenList{}.FromRef(_ret), _ok
}

// Type returns the value of property "HTMLStyleElement.type".
//
// The returned bool will be false if there is no such property.
func (this HTMLStyleElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLStyleElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLStyleElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLStyleElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLStyleElementType(
		this.Ref(),
		val.Ref(),
	)
}

// Sheet returns the value of property "HTMLStyleElement.sheet".
//
// The returned bool will be false if there is no such property.
func (this HTMLStyleElement) Sheet() (CSSStyleSheet, bool) {
	var _ok bool
	_ret := bindings.GetHTMLStyleElementSheet(
		this.Ref(), js.Pointer(&_ok),
	)
	return CSSStyleSheet{}.FromRef(_ret), _ok
}

func NewHTMLTableCaptionElement() HTMLTableCaptionElement {
	return HTMLTableCaptionElement{}.FromRef(
		bindings.NewHTMLTableCaptionElementByHTMLTableCaptionElement(),
	)
}

type HTMLTableCaptionElement struct {
	HTMLElement
}

func (this HTMLTableCaptionElement) Once() HTMLTableCaptionElement {
	this.Ref().Once()
	return this
}

func (this HTMLTableCaptionElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLTableCaptionElement) FromRef(ref js.Ref) HTMLTableCaptionElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLTableCaptionElement) Free() {
	this.Ref().Free()
}

// Align returns the value of property "HTMLTableCaptionElement.align".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCaptionElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCaptionElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLTableCaptionElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCaptionElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableCaptionElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLTableCellElement() HTMLTableCellElement {
	return HTMLTableCellElement{}.FromRef(
		bindings.NewHTMLTableCellElementByHTMLTableCellElement(),
	)
}

type HTMLTableCellElement struct {
	HTMLElement
}

func (this HTMLTableCellElement) Once() HTMLTableCellElement {
	this.Ref().Once()
	return this
}

func (this HTMLTableCellElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLTableCellElement) FromRef(ref js.Ref) HTMLTableCellElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLTableCellElement) Free() {
	this.Ref().Free()
}

// ColSpan returns the value of property "HTMLTableCellElement.colSpan".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) ColSpan() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementColSpan(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// ColSpan sets the value of property "HTMLTableCellElement.colSpan" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetColSpan(val uint32) bool {
	return js.True == bindings.SetHTMLTableCellElementColSpan(
		this.Ref(),
		uint32(val),
	)
}

// RowSpan returns the value of property "HTMLTableCellElement.rowSpan".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) RowSpan() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementRowSpan(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// RowSpan sets the value of property "HTMLTableCellElement.rowSpan" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetRowSpan(val uint32) bool {
	return js.True == bindings.SetHTMLTableCellElementRowSpan(
		this.Ref(),
		uint32(val),
	)
}

// Headers returns the value of property "HTMLTableCellElement.headers".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) Headers() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementHeaders(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Headers sets the value of property "HTMLTableCellElement.headers" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetHeaders(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementHeaders(
		this.Ref(),
		val.Ref(),
	)
}

// CellIndex returns the value of property "HTMLTableCellElement.cellIndex".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) CellIndex() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementCellIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Scope returns the value of property "HTMLTableCellElement.scope".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) Scope() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementScope(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Scope sets the value of property "HTMLTableCellElement.scope" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetScope(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementScope(
		this.Ref(),
		val.Ref(),
	)
}

// Abbr returns the value of property "HTMLTableCellElement.abbr".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) Abbr() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementAbbr(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Abbr sets the value of property "HTMLTableCellElement.abbr" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetAbbr(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementAbbr(
		this.Ref(),
		val.Ref(),
	)
}

// Align returns the value of property "HTMLTableCellElement.align".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLTableCellElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

// Axis returns the value of property "HTMLTableCellElement.axis".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) Axis() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementAxis(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Axis sets the value of property "HTMLTableCellElement.axis" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetAxis(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementAxis(
		this.Ref(),
		val.Ref(),
	)
}

// Height returns the value of property "HTMLTableCellElement.height".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) Height() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementHeight(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Height sets the value of property "HTMLTableCellElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementHeight(
		this.Ref(),
		val.Ref(),
	)
}

// Width returns the value of property "HTMLTableCellElement.width".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) Width() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Width sets the value of property "HTMLTableCellElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

// Ch returns the value of property "HTMLTableCellElement.ch".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) Ch() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementCh(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Ch sets the value of property "HTMLTableCellElement.ch" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetCh(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementCh(
		this.Ref(),
		val.Ref(),
	)
}

// ChOff returns the value of property "HTMLTableCellElement.chOff".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) ChOff() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementChOff(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ChOff sets the value of property "HTMLTableCellElement.chOff" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetChOff(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementChOff(
		this.Ref(),
		val.Ref(),
	)
}

// NoWrap returns the value of property "HTMLTableCellElement.noWrap".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) NoWrap() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementNoWrap(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// NoWrap sets the value of property "HTMLTableCellElement.noWrap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetNoWrap(val bool) bool {
	return js.True == bindings.SetHTMLTableCellElementNoWrap(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// VAlign returns the value of property "HTMLTableCellElement.vAlign".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) VAlign() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementVAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// VAlign sets the value of property "HTMLTableCellElement.vAlign" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetVAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementVAlign(
		this.Ref(),
		val.Ref(),
	)
}

// BgColor returns the value of property "HTMLTableCellElement.bgColor".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableCellElement) BgColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableCellElementBgColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BgColor sets the value of property "HTMLTableCellElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementBgColor(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLTableColElement() HTMLTableColElement {
	return HTMLTableColElement{}.FromRef(
		bindings.NewHTMLTableColElementByHTMLTableColElement(),
	)
}

type HTMLTableColElement struct {
	HTMLElement
}

func (this HTMLTableColElement) Once() HTMLTableColElement {
	this.Ref().Once()
	return this
}

func (this HTMLTableColElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLTableColElement) FromRef(ref js.Ref) HTMLTableColElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLTableColElement) Free() {
	this.Ref().Free()
}

// Span returns the value of property "HTMLTableColElement.span".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableColElement) Span() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableColElementSpan(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Span sets the value of property "HTMLTableColElement.span" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetSpan(val uint32) bool {
	return js.True == bindings.SetHTMLTableColElementSpan(
		this.Ref(),
		uint32(val),
	)
}

// Align returns the value of property "HTMLTableColElement.align".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableColElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableColElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLTableColElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableColElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

// Ch returns the value of property "HTMLTableColElement.ch".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableColElement) Ch() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableColElementCh(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Ch sets the value of property "HTMLTableColElement.ch" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetCh(val js.String) bool {
	return js.True == bindings.SetHTMLTableColElementCh(
		this.Ref(),
		val.Ref(),
	)
}

// ChOff returns the value of property "HTMLTableColElement.chOff".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableColElement) ChOff() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableColElementChOff(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ChOff sets the value of property "HTMLTableColElement.chOff" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetChOff(val js.String) bool {
	return js.True == bindings.SetHTMLTableColElementChOff(
		this.Ref(),
		val.Ref(),
	)
}

// VAlign returns the value of property "HTMLTableColElement.vAlign".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableColElement) VAlign() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableColElementVAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// VAlign sets the value of property "HTMLTableColElement.vAlign" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetVAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableColElementVAlign(
		this.Ref(),
		val.Ref(),
	)
}

// Width returns the value of property "HTMLTableColElement.width".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableColElement) Width() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableColElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Width sets the value of property "HTMLTableColElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLTableColElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLTableRowElement() HTMLTableRowElement {
	return HTMLTableRowElement{}.FromRef(
		bindings.NewHTMLTableRowElementByHTMLTableRowElement(),
	)
}

type HTMLTableRowElement struct {
	HTMLElement
}

func (this HTMLTableRowElement) Once() HTMLTableRowElement {
	this.Ref().Once()
	return this
}

func (this HTMLTableRowElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLTableRowElement) FromRef(ref js.Ref) HTMLTableRowElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLTableRowElement) Free() {
	this.Ref().Free()
}

// RowIndex returns the value of property "HTMLTableRowElement.rowIndex".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableRowElement) RowIndex() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableRowElementRowIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// SectionRowIndex returns the value of property "HTMLTableRowElement.sectionRowIndex".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableRowElement) SectionRowIndex() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableRowElementSectionRowIndex(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Cells returns the value of property "HTMLTableRowElement.cells".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableRowElement) Cells() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableRowElementCells(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// Align returns the value of property "HTMLTableRowElement.align".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableRowElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableRowElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLTableRowElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableRowElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableRowElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

// Ch returns the value of property "HTMLTableRowElement.ch".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableRowElement) Ch() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableRowElementCh(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Ch sets the value of property "HTMLTableRowElement.ch" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableRowElement) SetCh(val js.String) bool {
	return js.True == bindings.SetHTMLTableRowElementCh(
		this.Ref(),
		val.Ref(),
	)
}

// ChOff returns the value of property "HTMLTableRowElement.chOff".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableRowElement) ChOff() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableRowElementChOff(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ChOff sets the value of property "HTMLTableRowElement.chOff" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableRowElement) SetChOff(val js.String) bool {
	return js.True == bindings.SetHTMLTableRowElementChOff(
		this.Ref(),
		val.Ref(),
	)
}

// VAlign returns the value of property "HTMLTableRowElement.vAlign".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableRowElement) VAlign() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableRowElementVAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// VAlign sets the value of property "HTMLTableRowElement.vAlign" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableRowElement) SetVAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableRowElementVAlign(
		this.Ref(),
		val.Ref(),
	)
}

// BgColor returns the value of property "HTMLTableRowElement.bgColor".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableRowElement) BgColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableRowElementBgColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BgColor sets the value of property "HTMLTableRowElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableRowElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLTableRowElementBgColor(
		this.Ref(),
		val.Ref(),
	)
}

// InsertCell calls the method "HTMLTableRowElement.insertCell".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableRowElement) InsertCell(index int32) (HTMLTableCellElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableRowElementInsertCell(
		this.Ref(), js.Pointer(&_ok),
		int32(index),
	)

	return HTMLTableCellElement{}.FromRef(_ret), _ok
}

// InsertCellFunc returns the method "HTMLTableRowElement.insertCell".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableRowElement) InsertCellFunc() (fn js.Func[func(index int32) HTMLTableCellElement]) {
	return fn.FromRef(
		bindings.HTMLTableRowElementInsertCellFunc(
			this.Ref(),
		),
	)
}

// InsertCell1 calls the method "HTMLTableRowElement.insertCell".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableRowElement) InsertCell1() (HTMLTableCellElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableRowElementInsertCell1(
		this.Ref(), js.Pointer(&_ok),
	)

	return HTMLTableCellElement{}.FromRef(_ret), _ok
}

// InsertCell1Func returns the method "HTMLTableRowElement.insertCell".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableRowElement) InsertCell1Func() (fn js.Func[func() HTMLTableCellElement]) {
	return fn.FromRef(
		bindings.HTMLTableRowElementInsertCell1Func(
			this.Ref(),
		),
	)
}

// DeleteCell calls the method "HTMLTableRowElement.deleteCell".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableRowElement) DeleteCell(index int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableRowElementDeleteCell(
		this.Ref(), js.Pointer(&_ok),
		int32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteCellFunc returns the method "HTMLTableRowElement.deleteCell".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableRowElement) DeleteCellFunc() (fn js.Func[func(index int32)]) {
	return fn.FromRef(
		bindings.HTMLTableRowElementDeleteCellFunc(
			this.Ref(),
		),
	)
}

func NewHTMLTableSectionElement() HTMLTableSectionElement {
	return HTMLTableSectionElement{}.FromRef(
		bindings.NewHTMLTableSectionElementByHTMLTableSectionElement(),
	)
}

type HTMLTableSectionElement struct {
	HTMLElement
}

func (this HTMLTableSectionElement) Once() HTMLTableSectionElement {
	this.Ref().Once()
	return this
}

func (this HTMLTableSectionElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLTableSectionElement) FromRef(ref js.Ref) HTMLTableSectionElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLTableSectionElement) Free() {
	this.Ref().Free()
}

// Rows returns the value of property "HTMLTableSectionElement.rows".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableSectionElement) Rows() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableSectionElementRows(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// Align returns the value of property "HTMLTableSectionElement.align".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableSectionElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableSectionElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLTableSectionElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableSectionElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableSectionElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

// Ch returns the value of property "HTMLTableSectionElement.ch".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableSectionElement) Ch() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableSectionElementCh(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Ch sets the value of property "HTMLTableSectionElement.ch" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableSectionElement) SetCh(val js.String) bool {
	return js.True == bindings.SetHTMLTableSectionElementCh(
		this.Ref(),
		val.Ref(),
	)
}

// ChOff returns the value of property "HTMLTableSectionElement.chOff".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableSectionElement) ChOff() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableSectionElementChOff(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// ChOff sets the value of property "HTMLTableSectionElement.chOff" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableSectionElement) SetChOff(val js.String) bool {
	return js.True == bindings.SetHTMLTableSectionElementChOff(
		this.Ref(),
		val.Ref(),
	)
}

// VAlign returns the value of property "HTMLTableSectionElement.vAlign".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableSectionElement) VAlign() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableSectionElementVAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// VAlign sets the value of property "HTMLTableSectionElement.vAlign" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableSectionElement) SetVAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableSectionElementVAlign(
		this.Ref(),
		val.Ref(),
	)
}

// InsertRow calls the method "HTMLTableSectionElement.insertRow".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableSectionElement) InsertRow(index int32) (HTMLTableRowElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableSectionElementInsertRow(
		this.Ref(), js.Pointer(&_ok),
		int32(index),
	)

	return HTMLTableRowElement{}.FromRef(_ret), _ok
}

// InsertRowFunc returns the method "HTMLTableSectionElement.insertRow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableSectionElement) InsertRowFunc() (fn js.Func[func(index int32) HTMLTableRowElement]) {
	return fn.FromRef(
		bindings.HTMLTableSectionElementInsertRowFunc(
			this.Ref(),
		),
	)
}

// InsertRow1 calls the method "HTMLTableSectionElement.insertRow".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableSectionElement) InsertRow1() (HTMLTableRowElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableSectionElementInsertRow1(
		this.Ref(), js.Pointer(&_ok),
	)

	return HTMLTableRowElement{}.FromRef(_ret), _ok
}

// InsertRow1Func returns the method "HTMLTableSectionElement.insertRow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableSectionElement) InsertRow1Func() (fn js.Func[func() HTMLTableRowElement]) {
	return fn.FromRef(
		bindings.HTMLTableSectionElementInsertRow1Func(
			this.Ref(),
		),
	)
}

// DeleteRow calls the method "HTMLTableSectionElement.deleteRow".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableSectionElement) DeleteRow(index int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableSectionElementDeleteRow(
		this.Ref(), js.Pointer(&_ok),
		int32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteRowFunc returns the method "HTMLTableSectionElement.deleteRow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableSectionElement) DeleteRowFunc() (fn js.Func[func(index int32)]) {
	return fn.FromRef(
		bindings.HTMLTableSectionElementDeleteRowFunc(
			this.Ref(),
		),
	)
}

func NewHTMLTableElement() HTMLTableElement {
	return HTMLTableElement{}.FromRef(
		bindings.NewHTMLTableElementByHTMLTableElement(),
	)
}

type HTMLTableElement struct {
	HTMLElement
}

func (this HTMLTableElement) Once() HTMLTableElement {
	this.Ref().Once()
	return this
}

func (this HTMLTableElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLTableElement) FromRef(ref js.Ref) HTMLTableElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLTableElement) Free() {
	this.Ref().Free()
}

// Caption returns the value of property "HTMLTableElement.caption".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) Caption() (HTMLTableCaptionElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementCaption(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLTableCaptionElement{}.FromRef(_ret), _ok
}

// Caption sets the value of property "HTMLTableElement.caption" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetCaption(val HTMLTableCaptionElement) bool {
	return js.True == bindings.SetHTMLTableElementCaption(
		this.Ref(),
		val.Ref(),
	)
}

// THead returns the value of property "HTMLTableElement.tHead".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) THead() (HTMLTableSectionElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementTHead(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLTableSectionElement{}.FromRef(_ret), _ok
}

// THead sets the value of property "HTMLTableElement.tHead" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetTHead(val HTMLTableSectionElement) bool {
	return js.True == bindings.SetHTMLTableElementTHead(
		this.Ref(),
		val.Ref(),
	)
}

// TFoot returns the value of property "HTMLTableElement.tFoot".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) TFoot() (HTMLTableSectionElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementTFoot(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLTableSectionElement{}.FromRef(_ret), _ok
}

// TFoot sets the value of property "HTMLTableElement.tFoot" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetTFoot(val HTMLTableSectionElement) bool {
	return js.True == bindings.SetHTMLTableElementTFoot(
		this.Ref(),
		val.Ref(),
	)
}

// TBodies returns the value of property "HTMLTableElement.tBodies".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) TBodies() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementTBodies(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// Rows returns the value of property "HTMLTableElement.rows".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) Rows() (HTMLCollection, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementRows(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLCollection{}.FromRef(_ret), _ok
}

// Align returns the value of property "HTMLTableElement.align".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) Align() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementAlign(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Align sets the value of property "HTMLTableElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

// Border returns the value of property "HTMLTableElement.border".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) Border() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementBorder(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Border sets the value of property "HTMLTableElement.border" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetBorder(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementBorder(
		this.Ref(),
		val.Ref(),
	)
}

// Frame returns the value of property "HTMLTableElement.frame".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) Frame() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementFrame(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Frame sets the value of property "HTMLTableElement.frame" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetFrame(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementFrame(
		this.Ref(),
		val.Ref(),
	)
}

// Rules returns the value of property "HTMLTableElement.rules".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) Rules() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementRules(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Rules sets the value of property "HTMLTableElement.rules" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetRules(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementRules(
		this.Ref(),
		val.Ref(),
	)
}

// Summary returns the value of property "HTMLTableElement.summary".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) Summary() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementSummary(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Summary sets the value of property "HTMLTableElement.summary" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetSummary(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementSummary(
		this.Ref(),
		val.Ref(),
	)
}

// Width returns the value of property "HTMLTableElement.width".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) Width() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementWidth(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Width sets the value of property "HTMLTableElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

// BgColor returns the value of property "HTMLTableElement.bgColor".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) BgColor() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementBgColor(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// BgColor sets the value of property "HTMLTableElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementBgColor(
		this.Ref(),
		val.Ref(),
	)
}

// CellPadding returns the value of property "HTMLTableElement.cellPadding".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) CellPadding() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementCellPadding(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CellPadding sets the value of property "HTMLTableElement.cellPadding" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetCellPadding(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementCellPadding(
		this.Ref(),
		val.Ref(),
	)
}

// CellSpacing returns the value of property "HTMLTableElement.cellSpacing".
//
// The returned bool will be false if there is no such property.
func (this HTMLTableElement) CellSpacing() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTableElementCellSpacing(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// CellSpacing sets the value of property "HTMLTableElement.cellSpacing" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetCellSpacing(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementCellSpacing(
		this.Ref(),
		val.Ref(),
	)
}

// CreateCaption calls the method "HTMLTableElement.createCaption".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableElement) CreateCaption() (HTMLTableCaptionElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableElementCreateCaption(
		this.Ref(), js.Pointer(&_ok),
	)

	return HTMLTableCaptionElement{}.FromRef(_ret), _ok
}

// CreateCaptionFunc returns the method "HTMLTableElement.createCaption".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableElement) CreateCaptionFunc() (fn js.Func[func() HTMLTableCaptionElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementCreateCaptionFunc(
			this.Ref(),
		),
	)
}

// DeleteCaption calls the method "HTMLTableElement.deleteCaption".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableElement) DeleteCaption() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableElementDeleteCaption(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteCaptionFunc returns the method "HTMLTableElement.deleteCaption".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableElement) DeleteCaptionFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLTableElementDeleteCaptionFunc(
			this.Ref(),
		),
	)
}

// CreateTHead calls the method "HTMLTableElement.createTHead".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableElement) CreateTHead() (HTMLTableSectionElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableElementCreateTHead(
		this.Ref(), js.Pointer(&_ok),
	)

	return HTMLTableSectionElement{}.FromRef(_ret), _ok
}

// CreateTHeadFunc returns the method "HTMLTableElement.createTHead".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableElement) CreateTHeadFunc() (fn js.Func[func() HTMLTableSectionElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementCreateTHeadFunc(
			this.Ref(),
		),
	)
}

// DeleteTHead calls the method "HTMLTableElement.deleteTHead".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableElement) DeleteTHead() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableElementDeleteTHead(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteTHeadFunc returns the method "HTMLTableElement.deleteTHead".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableElement) DeleteTHeadFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLTableElementDeleteTHeadFunc(
			this.Ref(),
		),
	)
}

// CreateTFoot calls the method "HTMLTableElement.createTFoot".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableElement) CreateTFoot() (HTMLTableSectionElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableElementCreateTFoot(
		this.Ref(), js.Pointer(&_ok),
	)

	return HTMLTableSectionElement{}.FromRef(_ret), _ok
}

// CreateTFootFunc returns the method "HTMLTableElement.createTFoot".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableElement) CreateTFootFunc() (fn js.Func[func() HTMLTableSectionElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementCreateTFootFunc(
			this.Ref(),
		),
	)
}

// DeleteTFoot calls the method "HTMLTableElement.deleteTFoot".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableElement) DeleteTFoot() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableElementDeleteTFoot(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteTFootFunc returns the method "HTMLTableElement.deleteTFoot".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableElement) DeleteTFootFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLTableElementDeleteTFootFunc(
			this.Ref(),
		),
	)
}

// CreateTBody calls the method "HTMLTableElement.createTBody".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableElement) CreateTBody() (HTMLTableSectionElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableElementCreateTBody(
		this.Ref(), js.Pointer(&_ok),
	)

	return HTMLTableSectionElement{}.FromRef(_ret), _ok
}

// CreateTBodyFunc returns the method "HTMLTableElement.createTBody".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableElement) CreateTBodyFunc() (fn js.Func[func() HTMLTableSectionElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementCreateTBodyFunc(
			this.Ref(),
		),
	)
}

// InsertRow calls the method "HTMLTableElement.insertRow".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableElement) InsertRow(index int32) (HTMLTableRowElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableElementInsertRow(
		this.Ref(), js.Pointer(&_ok),
		int32(index),
	)

	return HTMLTableRowElement{}.FromRef(_ret), _ok
}

// InsertRowFunc returns the method "HTMLTableElement.insertRow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableElement) InsertRowFunc() (fn js.Func[func(index int32) HTMLTableRowElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementInsertRowFunc(
			this.Ref(),
		),
	)
}

// InsertRow1 calls the method "HTMLTableElement.insertRow".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableElement) InsertRow1() (HTMLTableRowElement, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableElementInsertRow1(
		this.Ref(), js.Pointer(&_ok),
	)

	return HTMLTableRowElement{}.FromRef(_ret), _ok
}

// InsertRow1Func returns the method "HTMLTableElement.insertRow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableElement) InsertRow1Func() (fn js.Func[func() HTMLTableRowElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementInsertRow1Func(
			this.Ref(),
		),
	)
}

// DeleteRow calls the method "HTMLTableElement.deleteRow".
//
// The returned bool will be false if there is no such method.
func (this HTMLTableElement) DeleteRow(index int32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTableElementDeleteRow(
		this.Ref(), js.Pointer(&_ok),
		int32(index),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteRowFunc returns the method "HTMLTableElement.deleteRow".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTableElement) DeleteRowFunc() (fn js.Func[func(index int32)]) {
	return fn.FromRef(
		bindings.HTMLTableElementDeleteRowFunc(
			this.Ref(),
		),
	)
}

func NewHTMLTemplateElement() HTMLTemplateElement {
	return HTMLTemplateElement{}.FromRef(
		bindings.NewHTMLTemplateElementByHTMLTemplateElement(),
	)
}

type HTMLTemplateElement struct {
	HTMLElement
}

func (this HTMLTemplateElement) Once() HTMLTemplateElement {
	this.Ref().Once()
	return this
}

func (this HTMLTemplateElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLTemplateElement) FromRef(ref js.Ref) HTMLTemplateElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLTemplateElement) Free() {
	this.Ref().Free()
}

// Content returns the value of property "HTMLTemplateElement.content".
//
// The returned bool will be false if there is no such property.
func (this HTMLTemplateElement) Content() (DocumentFragment, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTemplateElementContent(
		this.Ref(), js.Pointer(&_ok),
	)
	return DocumentFragment{}.FromRef(_ret), _ok
}

func NewHTMLTextAreaElement() HTMLTextAreaElement {
	return HTMLTextAreaElement{}.FromRef(
		bindings.NewHTMLTextAreaElementByHTMLTextAreaElement(),
	)
}

type HTMLTextAreaElement struct {
	HTMLElement
}

func (this HTMLTextAreaElement) Once() HTMLTextAreaElement {
	this.Ref().Once()
	return this
}

func (this HTMLTextAreaElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLTextAreaElement) FromRef(ref js.Ref) HTMLTextAreaElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLTextAreaElement) Free() {
	this.Ref().Free()
}

// Autocomplete returns the value of property "HTMLTextAreaElement.autocomplete".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Autocomplete() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementAutocomplete(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Autocomplete sets the value of property "HTMLTextAreaElement.autocomplete" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetAutocomplete(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementAutocomplete(
		this.Ref(),
		val.Ref(),
	)
}

// Cols returns the value of property "HTMLTextAreaElement.cols".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Cols() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementCols(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Cols sets the value of property "HTMLTextAreaElement.cols" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetCols(val uint32) bool {
	return js.True == bindings.SetHTMLTextAreaElementCols(
		this.Ref(),
		uint32(val),
	)
}

// DirName returns the value of property "HTMLTextAreaElement.dirName".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) DirName() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementDirName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DirName sets the value of property "HTMLTextAreaElement.dirName" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetDirName(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementDirName(
		this.Ref(),
		val.Ref(),
	)
}

// Disabled returns the value of property "HTMLTextAreaElement.disabled".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Disabled() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementDisabled(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Disabled sets the value of property "HTMLTextAreaElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLTextAreaElementDisabled(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLTextAreaElement.form".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Form() (HTMLFormElement, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementForm(
		this.Ref(), js.Pointer(&_ok),
	)
	return HTMLFormElement{}.FromRef(_ret), _ok
}

// MaxLength returns the value of property "HTMLTextAreaElement.maxLength".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) MaxLength() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementMaxLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// MaxLength sets the value of property "HTMLTextAreaElement.maxLength" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetMaxLength(val int32) bool {
	return js.True == bindings.SetHTMLTextAreaElementMaxLength(
		this.Ref(),
		int32(val),
	)
}

// MinLength returns the value of property "HTMLTextAreaElement.minLength".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) MinLength() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementMinLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// MinLength sets the value of property "HTMLTextAreaElement.minLength" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetMinLength(val int32) bool {
	return js.True == bindings.SetHTMLTextAreaElementMinLength(
		this.Ref(),
		int32(val),
	)
}

// Name returns the value of property "HTMLTextAreaElement.name".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "HTMLTextAreaElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementName(
		this.Ref(),
		val.Ref(),
	)
}

// Placeholder returns the value of property "HTMLTextAreaElement.placeholder".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Placeholder() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementPlaceholder(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Placeholder sets the value of property "HTMLTextAreaElement.placeholder" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetPlaceholder(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementPlaceholder(
		this.Ref(),
		val.Ref(),
	)
}

// ReadOnly returns the value of property "HTMLTextAreaElement.readOnly".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) ReadOnly() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementReadOnly(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// ReadOnly sets the value of property "HTMLTextAreaElement.readOnly" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetReadOnly(val bool) bool {
	return js.True == bindings.SetHTMLTextAreaElementReadOnly(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Required returns the value of property "HTMLTextAreaElement.required".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Required() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementRequired(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Required sets the value of property "HTMLTextAreaElement.required" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetRequired(val bool) bool {
	return js.True == bindings.SetHTMLTextAreaElementRequired(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Rows returns the value of property "HTMLTextAreaElement.rows".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Rows() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementRows(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// Rows sets the value of property "HTMLTextAreaElement.rows" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetRows(val uint32) bool {
	return js.True == bindings.SetHTMLTextAreaElementRows(
		this.Ref(),
		uint32(val),
	)
}

// Wrap returns the value of property "HTMLTextAreaElement.wrap".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Wrap() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementWrap(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Wrap sets the value of property "HTMLTextAreaElement.wrap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetWrap(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementWrap(
		this.Ref(),
		val.Ref(),
	)
}

// Type returns the value of property "HTMLTextAreaElement.type".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DefaultValue returns the value of property "HTMLTextAreaElement.defaultValue".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) DefaultValue() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementDefaultValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DefaultValue sets the value of property "HTMLTextAreaElement.defaultValue" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetDefaultValue(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementDefaultValue(
		this.Ref(),
		val.Ref(),
	)
}

// Value returns the value of property "HTMLTextAreaElement.value".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Value() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementValue(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Value sets the value of property "HTMLTextAreaElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementValue(
		this.Ref(),
		val.Ref(),
	)
}

// TextLength returns the value of property "HTMLTextAreaElement.textLength".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) TextLength() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementTextLength(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// WillValidate returns the value of property "HTMLTextAreaElement.willValidate".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) WillValidate() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementWillValidate(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Validity returns the value of property "HTMLTextAreaElement.validity".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Validity() (ValidityState, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementValidity(
		this.Ref(), js.Pointer(&_ok),
	)
	return ValidityState{}.FromRef(_ret), _ok
}

// ValidationMessage returns the value of property "HTMLTextAreaElement.validationMessage".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) ValidationMessage() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementValidationMessage(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Labels returns the value of property "HTMLTextAreaElement.labels".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) Labels() (NodeList, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementLabels(
		this.Ref(), js.Pointer(&_ok),
	)
	return NodeList{}.FromRef(_ret), _ok
}

// SelectionStart returns the value of property "HTMLTextAreaElement.selectionStart".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) SelectionStart() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementSelectionStart(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// SelectionStart sets the value of property "HTMLTextAreaElement.selectionStart" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetSelectionStart(val uint32) bool {
	return js.True == bindings.SetHTMLTextAreaElementSelectionStart(
		this.Ref(),
		uint32(val),
	)
}

// SelectionEnd returns the value of property "HTMLTextAreaElement.selectionEnd".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) SelectionEnd() (uint32, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementSelectionEnd(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint32(_ret), _ok
}

// SelectionEnd sets the value of property "HTMLTextAreaElement.selectionEnd" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetSelectionEnd(val uint32) bool {
	return js.True == bindings.SetHTMLTextAreaElementSelectionEnd(
		this.Ref(),
		uint32(val),
	)
}

// SelectionDirection returns the value of property "HTMLTextAreaElement.selectionDirection".
//
// The returned bool will be false if there is no such property.
func (this HTMLTextAreaElement) SelectionDirection() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTextAreaElementSelectionDirection(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// SelectionDirection sets the value of property "HTMLTextAreaElement.selectionDirection" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetSelectionDirection(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementSelectionDirection(
		this.Ref(),
		val.Ref(),
	)
}

// CheckValidity calls the method "HTMLTextAreaElement.checkValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLTextAreaElement) CheckValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTextAreaElementCheckValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// CheckValidityFunc returns the method "HTMLTextAreaElement.checkValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTextAreaElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLTextAreaElement.reportValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLTextAreaElement) ReportValidity() (bool, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTextAreaElementReportValidity(
		this.Ref(), js.Pointer(&_ok),
	)

	return _ret == js.True, _ok
}

// ReportValidityFunc returns the method "HTMLTextAreaElement.reportValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTextAreaElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLTextAreaElement.setCustomValidity".
//
// The returned bool will be false if there is no such method.
func (this HTMLTextAreaElement) SetCustomValidity(err js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTextAreaElementSetCustomValidity(
		this.Ref(), js.Pointer(&_ok),
		err.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetCustomValidityFunc returns the method "HTMLTextAreaElement.setCustomValidity".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTextAreaElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

// Select calls the method "HTMLTextAreaElement.select".
//
// The returned bool will be false if there is no such method.
func (this HTMLTextAreaElement) Select() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTextAreaElementSelect(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SelectFunc returns the method "HTMLTextAreaElement.select".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTextAreaElement) SelectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSelectFunc(
			this.Ref(),
		),
	)
}

// SetRangeText calls the method "HTMLTextAreaElement.setRangeText".
//
// The returned bool will be false if there is no such method.
func (this HTMLTextAreaElement) SetRangeText(replacement js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTextAreaElementSetRangeText(
		this.Ref(), js.Pointer(&_ok),
		replacement.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetRangeTextFunc returns the method "HTMLTextAreaElement.setRangeText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTextAreaElement) SetRangeTextFunc() (fn js.Func[func(replacement js.String)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetRangeTextFunc(
			this.Ref(),
		),
	)
}

// SetRangeText1 calls the method "HTMLTextAreaElement.setRangeText".
//
// The returned bool will be false if there is no such method.
func (this HTMLTextAreaElement) SetRangeText1(replacement js.String, start uint32, end uint32, selectionMode SelectionMode) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTextAreaElementSetRangeText1(
		this.Ref(), js.Pointer(&_ok),
		replacement.Ref(),
		uint32(start),
		uint32(end),
		uint32(selectionMode),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetRangeText1Func returns the method "HTMLTextAreaElement.setRangeText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTextAreaElement) SetRangeText1Func() (fn js.Func[func(replacement js.String, start uint32, end uint32, selectionMode SelectionMode)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetRangeText1Func(
			this.Ref(),
		),
	)
}

// SetRangeText2 calls the method "HTMLTextAreaElement.setRangeText".
//
// The returned bool will be false if there is no such method.
func (this HTMLTextAreaElement) SetRangeText2(replacement js.String, start uint32, end uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTextAreaElementSetRangeText2(
		this.Ref(), js.Pointer(&_ok),
		replacement.Ref(),
		uint32(start),
		uint32(end),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetRangeText2Func returns the method "HTMLTextAreaElement.setRangeText".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTextAreaElement) SetRangeText2Func() (fn js.Func[func(replacement js.String, start uint32, end uint32)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetRangeText2Func(
			this.Ref(),
		),
	)
}

// SetSelectionRange calls the method "HTMLTextAreaElement.setSelectionRange".
//
// The returned bool will be false if there is no such method.
func (this HTMLTextAreaElement) SetSelectionRange(start uint32, end uint32, direction js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTextAreaElementSetSelectionRange(
		this.Ref(), js.Pointer(&_ok),
		uint32(start),
		uint32(end),
		direction.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetSelectionRangeFunc returns the method "HTMLTextAreaElement.setSelectionRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTextAreaElement) SetSelectionRangeFunc() (fn js.Func[func(start uint32, end uint32, direction js.String)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetSelectionRangeFunc(
			this.Ref(),
		),
	)
}

// SetSelectionRange1 calls the method "HTMLTextAreaElement.setSelectionRange".
//
// The returned bool will be false if there is no such method.
func (this HTMLTextAreaElement) SetSelectionRange1(start uint32, end uint32) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallHTMLTextAreaElementSetSelectionRange1(
		this.Ref(), js.Pointer(&_ok),
		uint32(start),
		uint32(end),
	)

	_ = _ret
	return js.Void{}, _ok
}

// SetSelectionRange1Func returns the method "HTMLTextAreaElement.setSelectionRange".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this HTMLTextAreaElement) SetSelectionRange1Func() (fn js.Func[func(start uint32, end uint32)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetSelectionRange1Func(
			this.Ref(),
		),
	)
}

func NewHTMLTimeElement() HTMLTimeElement {
	return HTMLTimeElement{}.FromRef(
		bindings.NewHTMLTimeElementByHTMLTimeElement(),
	)
}

type HTMLTimeElement struct {
	HTMLElement
}

func (this HTMLTimeElement) Once() HTMLTimeElement {
	this.Ref().Once()
	return this
}

func (this HTMLTimeElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLTimeElement) FromRef(ref js.Ref) HTMLTimeElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLTimeElement) Free() {
	this.Ref().Free()
}

// DateTime returns the value of property "HTMLTimeElement.dateTime".
//
// The returned bool will be false if there is no such property.
func (this HTMLTimeElement) DateTime() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTimeElementDateTime(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// DateTime sets the value of property "HTMLTimeElement.dateTime" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTimeElement) SetDateTime(val js.String) bool {
	return js.True == bindings.SetHTMLTimeElementDateTime(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLTitleElement() HTMLTitleElement {
	return HTMLTitleElement{}.FromRef(
		bindings.NewHTMLTitleElementByHTMLTitleElement(),
	)
}

type HTMLTitleElement struct {
	HTMLElement
}

func (this HTMLTitleElement) Once() HTMLTitleElement {
	this.Ref().Once()
	return this
}

func (this HTMLTitleElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLTitleElement) FromRef(ref js.Ref) HTMLTitleElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLTitleElement) Free() {
	this.Ref().Free()
}

// Text returns the value of property "HTMLTitleElement.text".
//
// The returned bool will be false if there is no such property.
func (this HTMLTitleElement) Text() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTitleElementText(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Text sets the value of property "HTMLTitleElement.text" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTitleElement) SetText(val js.String) bool {
	return js.True == bindings.SetHTMLTitleElementText(
		this.Ref(),
		val.Ref(),
	)
}

const (
	HTMLTrackElement_NONE    uint16 = 0
	HTMLTrackElement_LOADING uint16 = 1
	HTMLTrackElement_LOADED  uint16 = 2
	HTMLTrackElement_ERROR   uint16 = 3
)

func NewHTMLTrackElement() HTMLTrackElement {
	return HTMLTrackElement{}.FromRef(
		bindings.NewHTMLTrackElementByHTMLTrackElement(),
	)
}

type HTMLTrackElement struct {
	HTMLElement
}

func (this HTMLTrackElement) Once() HTMLTrackElement {
	this.Ref().Once()
	return this
}

func (this HTMLTrackElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLTrackElement) FromRef(ref js.Ref) HTMLTrackElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLTrackElement) Free() {
	this.Ref().Free()
}

// Kind returns the value of property "HTMLTrackElement.kind".
//
// The returned bool will be false if there is no such property.
func (this HTMLTrackElement) Kind() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTrackElementKind(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Kind sets the value of property "HTMLTrackElement.kind" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTrackElement) SetKind(val js.String) bool {
	return js.True == bindings.SetHTMLTrackElementKind(
		this.Ref(),
		val.Ref(),
	)
}

// Src returns the value of property "HTMLTrackElement.src".
//
// The returned bool will be false if there is no such property.
func (this HTMLTrackElement) Src() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTrackElementSrc(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Src sets the value of property "HTMLTrackElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTrackElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLTrackElementSrc(
		this.Ref(),
		val.Ref(),
	)
}

// Srclang returns the value of property "HTMLTrackElement.srclang".
//
// The returned bool will be false if there is no such property.
func (this HTMLTrackElement) Srclang() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTrackElementSrclang(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Srclang sets the value of property "HTMLTrackElement.srclang" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTrackElement) SetSrclang(val js.String) bool {
	return js.True == bindings.SetHTMLTrackElementSrclang(
		this.Ref(),
		val.Ref(),
	)
}

// Label returns the value of property "HTMLTrackElement.label".
//
// The returned bool will be false if there is no such property.
func (this HTMLTrackElement) Label() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTrackElementLabel(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Label sets the value of property "HTMLTrackElement.label" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTrackElement) SetLabel(val js.String) bool {
	return js.True == bindings.SetHTMLTrackElementLabel(
		this.Ref(),
		val.Ref(),
	)
}

// Default returns the value of property "HTMLTrackElement.default".
//
// The returned bool will be false if there is no such property.
func (this HTMLTrackElement) Default() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTrackElementDefault(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Default sets the value of property "HTMLTrackElement.default" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTrackElement) SetDefault(val bool) bool {
	return js.True == bindings.SetHTMLTrackElementDefault(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// ReadyState returns the value of property "HTMLTrackElement.readyState".
//
// The returned bool will be false if there is no such property.
func (this HTMLTrackElement) ReadyState() (uint16, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTrackElementReadyState(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint16(_ret), _ok
}

// Track returns the value of property "HTMLTrackElement.track".
//
// The returned bool will be false if there is no such property.
func (this HTMLTrackElement) Track() (TextTrack, bool) {
	var _ok bool
	_ret := bindings.GetHTMLTrackElementTrack(
		this.Ref(), js.Pointer(&_ok),
	)
	return TextTrack{}.FromRef(_ret), _ok
}

func NewHTMLUListElement() HTMLUListElement {
	return HTMLUListElement{}.FromRef(
		bindings.NewHTMLUListElementByHTMLUListElement(),
	)
}

type HTMLUListElement struct {
	HTMLElement
}

func (this HTMLUListElement) Once() HTMLUListElement {
	this.Ref().Once()
	return this
}

func (this HTMLUListElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLUListElement) FromRef(ref js.Ref) HTMLUListElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLUListElement) Free() {
	this.Ref().Free()
}

// Compact returns the value of property "HTMLUListElement.compact".
//
// The returned bool will be false if there is no such property.
func (this HTMLUListElement) Compact() (bool, bool) {
	var _ok bool
	_ret := bindings.GetHTMLUListElementCompact(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Compact sets the value of property "HTMLUListElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLUListElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLUListElementCompact(
		this.Ref(),
		js.Bool(bool(val)),
	)
}

// Type returns the value of property "HTMLUListElement.type".
//
// The returned bool will be false if there is no such property.
func (this HTMLUListElement) Type() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHTMLUListElementType(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Type sets the value of property "HTMLUListElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLUListElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLUListElementType(
		this.Ref(),
		val.Ref(),
	)
}

type HTMLUnknownElement struct {
	HTMLElement
}

func (this HTMLUnknownElement) Once() HTMLUnknownElement {
	this.Ref().Once()
	return this
}

func (this HTMLUnknownElement) Ref() js.Ref {
	return this.HTMLElement.Ref()
}

func (this HTMLUnknownElement) FromRef(ref js.Ref) HTMLUnknownElement {
	this.HTMLElement = this.HTMLElement.FromRef(ref)
	return this
}

func (this HTMLUnknownElement) Free() {
	this.Ref().Free()
}

type HashChangeEventInit struct {
	// OldURL is "HashChangeEventInit.oldURL"
	//
	// Optional, defaults to "".
	OldURL js.String
	// NewURL is "HashChangeEventInit.newURL"
	//
	// Optional, defaults to "".
	NewURL js.String
	// Bubbles is "HashChangeEventInit.bubbles"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Bubbles MUST be set to true to make this field effective.
	Bubbles bool
	// Cancelable is "HashChangeEventInit.cancelable"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Cancelable MUST be set to true to make this field effective.
	Cancelable bool
	// Composed is "HashChangeEventInit.composed"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Composed MUST be set to true to make this field effective.
	Composed bool

	FFI_USE_Bubbles    bool // for Bubbles.
	FFI_USE_Cancelable bool // for Cancelable.
	FFI_USE_Composed   bool // for Composed.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HashChangeEventInit with all fields set.
func (p HashChangeEventInit) FromRef(ref js.Ref) HashChangeEventInit {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HashChangeEventInit in the application heap.
func (p HashChangeEventInit) New() js.Ref {
	return bindings.HashChangeEventInitJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HashChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.HashChangeEventInitJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HashChangeEventInit) Update(ref js.Ref) {
	bindings.HashChangeEventInitJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

func NewHashChangeEvent(typ js.String, eventInitDict HashChangeEventInit) HashChangeEvent {
	return HashChangeEvent{}.FromRef(
		bindings.NewHashChangeEventByHashChangeEvent(
			typ.Ref(),
			js.Pointer(&eventInitDict)),
	)
}

func NewHashChangeEventByHashChangeEvent1(typ js.String) HashChangeEvent {
	return HashChangeEvent{}.FromRef(
		bindings.NewHashChangeEventByHashChangeEvent1(
			typ.Ref()),
	)
}

type HashChangeEvent struct {
	Event
}

func (this HashChangeEvent) Once() HashChangeEvent {
	this.Ref().Once()
	return this
}

func (this HashChangeEvent) Ref() js.Ref {
	return this.Event.Ref()
}

func (this HashChangeEvent) FromRef(ref js.Ref) HashChangeEvent {
	this.Event = this.Event.FromRef(ref)
	return this
}

func (this HashChangeEvent) Free() {
	this.Ref().Free()
}

// OldURL returns the value of property "HashChangeEvent.oldURL".
//
// The returned bool will be false if there is no such property.
func (this HashChangeEvent) OldURL() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHashChangeEventOldURL(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// NewURL returns the value of property "HashChangeEvent.newURL".
//
// The returned bool will be false if there is no such property.
func (this HashChangeEvent) NewURL() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetHashChangeEventNewURL(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

type HevcBitstreamFormat uint32

const (
	_ HevcBitstreamFormat = iota

	HevcBitstreamFormat_ANNEXB
	HevcBitstreamFormat_HEVC
)

func (HevcBitstreamFormat) FromRef(str js.Ref) HevcBitstreamFormat {
	return HevcBitstreamFormat(bindings.ConstOfHevcBitstreamFormat(str))
}

func (x HevcBitstreamFormat) String() (string, bool) {
	switch x {
	case HevcBitstreamFormat_ANNEXB:
		return "annexb", true
	case HevcBitstreamFormat_HEVC:
		return "hevc", true
	default:
		return "", false
	}
}

type HevcEncoderConfig struct {
	// Format is "HevcEncoderConfig.format"
	//
	// Optional, defaults to "hevc".
	Format HevcBitstreamFormat

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HevcEncoderConfig with all fields set.
func (p HevcEncoderConfig) FromRef(ref js.Ref) HevcEncoderConfig {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HevcEncoderConfig in the application heap.
func (p HevcEncoderConfig) New() js.Ref {
	return bindings.HevcEncoderConfigJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HevcEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.HevcEncoderConfigJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HevcEncoderConfig) Update(ref js.Ref) {
	bindings.HevcEncoderConfigJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HighlightType uint32

const (
	_ HighlightType = iota

	HighlightType_HIGHLIGHT
	HighlightType_SPELLING_ERROR
	HighlightType_GRAMMAR_ERROR
)

func (HighlightType) FromRef(str js.Ref) HighlightType {
	return HighlightType(bindings.ConstOfHighlightType(str))
}

func (x HighlightType) String() (string, bool) {
	switch x {
	case HighlightType_HIGHLIGHT:
		return "highlight", true
	case HighlightType_SPELLING_ERROR:
		return "spelling-error", true
	case HighlightType_GRAMMAR_ERROR:
		return "grammar-error", true
	default:
		return "", false
	}
}

func NewHighlight(initialRanges ...AbstractRange) Highlight {
	return Highlight{}.FromRef(
		bindings.NewHighlightByHighlight(
			js.SliceData(initialRanges),
			js.SizeU(len(initialRanges))),
	)
}

type Highlight struct {
	ref js.Ref
}

func (this Highlight) Once() Highlight {
	this.Ref().Once()
	return this
}

func (this Highlight) Ref() js.Ref {
	return this.ref
}

func (this Highlight) FromRef(ref js.Ref) Highlight {
	this.ref = ref
	return this
}

func (this Highlight) Free() {
	this.Ref().Free()
}

// Priority returns the value of property "Highlight.priority".
//
// The returned bool will be false if there is no such property.
func (this Highlight) Priority() (int32, bool) {
	var _ok bool
	_ret := bindings.GetHighlightPriority(
		this.Ref(), js.Pointer(&_ok),
	)
	return int32(_ret), _ok
}

// Priority sets the value of property "Highlight.priority" to val.
//
// It returns false if the property cannot be set.
func (this Highlight) SetPriority(val int32) bool {
	return js.True == bindings.SetHighlightPriority(
		this.Ref(),
		int32(val),
	)
}

// Type returns the value of property "Highlight.type".
//
// The returned bool will be false if there is no such property.
func (this Highlight) Type() (HighlightType, bool) {
	var _ok bool
	_ret := bindings.GetHighlightType(
		this.Ref(), js.Pointer(&_ok),
	)
	return HighlightType(_ret), _ok
}

// Type sets the value of property "Highlight.type" to val.
//
// It returns false if the property cannot be set.
func (this Highlight) SetType(val HighlightType) bool {
	return js.True == bindings.SetHighlightType(
		this.Ref(),
		uint32(val),
	)
}

type HkdfParams struct {
	// Hash is "HkdfParams.hash"
	//
	// Required
	Hash HashAlgorithmIdentifier
	// Salt is "HkdfParams.salt"
	//
	// Required
	Salt BufferSource
	// Info is "HkdfParams.info"
	//
	// Required
	Info BufferSource
	// Name is "HkdfParams.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HkdfParams with all fields set.
func (p HkdfParams) FromRef(ref js.Ref) HkdfParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HkdfParams in the application heap.
func (p HkdfParams) New() js.Ref {
	return bindings.HkdfParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HkdfParams) UpdateFrom(ref js.Ref) {
	bindings.HkdfParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HkdfParams) Update(ref js.Ref) {
	bindings.HkdfParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HmacImportParams struct {
	// Hash is "HmacImportParams.hash"
	//
	// Required
	Hash HashAlgorithmIdentifier
	// Length is "HmacImportParams.length"
	//
	// Optional
	//
	// NOTE: FFI_USE_Length MUST be set to true to make this field effective.
	Length uint32
	// Name is "HmacImportParams.name"
	//
	// Required
	Name js.String

	FFI_USE_Length bool // for Length.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HmacImportParams with all fields set.
func (p HmacImportParams) FromRef(ref js.Ref) HmacImportParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HmacImportParams in the application heap.
func (p HmacImportParams) New() js.Ref {
	return bindings.HmacImportParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HmacImportParams) UpdateFrom(ref js.Ref) {
	bindings.HmacImportParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HmacImportParams) Update(ref js.Ref) {
	bindings.HmacImportParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type KeyAlgorithm struct {
	// Name is "KeyAlgorithm.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a KeyAlgorithm with all fields set.
func (p KeyAlgorithm) FromRef(ref js.Ref) KeyAlgorithm {
	p.UpdateFrom(ref)
	return p
}

// New creates a new KeyAlgorithm in the application heap.
func (p KeyAlgorithm) New() js.Ref {
	return bindings.KeyAlgorithmJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p KeyAlgorithm) UpdateFrom(ref js.Ref) {
	bindings.KeyAlgorithmJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p KeyAlgorithm) Update(ref js.Ref) {
	bindings.KeyAlgorithmJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HmacKeyAlgorithm struct {
	// Hash is "HmacKeyAlgorithm.hash"
	//
	// Required
	Hash KeyAlgorithm
	// Length is "HmacKeyAlgorithm.length"
	//
	// Required
	Length uint32
	// Name is "HmacKeyAlgorithm.name"
	//
	// Required
	Name js.String

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HmacKeyAlgorithm with all fields set.
func (p HmacKeyAlgorithm) FromRef(ref js.Ref) HmacKeyAlgorithm {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HmacKeyAlgorithm in the application heap.
func (p HmacKeyAlgorithm) New() js.Ref {
	return bindings.HmacKeyAlgorithmJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HmacKeyAlgorithm) UpdateFrom(ref js.Ref) {
	bindings.HmacKeyAlgorithmJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HmacKeyAlgorithm) Update(ref js.Ref) {
	bindings.HmacKeyAlgorithmJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type HmacKeyGenParams struct {
	// Hash is "HmacKeyGenParams.hash"
	//
	// Required
	Hash HashAlgorithmIdentifier
	// Length is "HmacKeyGenParams.length"
	//
	// Optional
	//
	// NOTE: FFI_USE_Length MUST be set to true to make this field effective.
	Length uint32
	// Name is "HmacKeyGenParams.name"
	//
	// Required
	Name js.String

	FFI_USE_Length bool // for Length.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a HmacKeyGenParams with all fields set.
func (p HmacKeyGenParams) FromRef(ref js.Ref) HmacKeyGenParams {
	p.UpdateFrom(ref)
	return p
}

// New creates a new HmacKeyGenParams in the application heap.
func (p HmacKeyGenParams) New() js.Ref {
	return bindings.HmacKeyGenParamsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p HmacKeyGenParams) UpdateFrom(ref js.Ref) {
	bindings.HmacKeyGenParamsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p HmacKeyGenParams) Update(ref js.Ref) {
	bindings.HmacKeyGenParamsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IDBCursorDirection uint32

const (
	_ IDBCursorDirection = iota

	IDBCursorDirection_NEXT
	IDBCursorDirection_NEXTUNIQUE
	IDBCursorDirection_PREV
	IDBCursorDirection_PREVUNIQUE
)

func (IDBCursorDirection) FromRef(str js.Ref) IDBCursorDirection {
	return IDBCursorDirection(bindings.ConstOfIDBCursorDirection(str))
}

func (x IDBCursorDirection) String() (string, bool) {
	switch x {
	case IDBCursorDirection_NEXT:
		return "next", true
	case IDBCursorDirection_NEXTUNIQUE:
		return "nextunique", true
	case IDBCursorDirection_PREV:
		return "prev", true
	case IDBCursorDirection_PREVUNIQUE:
		return "prevunique", true
	default:
		return "", false
	}
}

type IDBIndex struct {
	ref js.Ref
}

func (this IDBIndex) Once() IDBIndex {
	this.Ref().Once()
	return this
}

func (this IDBIndex) Ref() js.Ref {
	return this.ref
}

func (this IDBIndex) FromRef(ref js.Ref) IDBIndex {
	this.ref = ref
	return this
}

func (this IDBIndex) Free() {
	this.Ref().Free()
}

// Name returns the value of property "IDBIndex.name".
//
// The returned bool will be false if there is no such property.
func (this IDBIndex) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetIDBIndexName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "IDBIndex.name" to val.
//
// It returns false if the property cannot be set.
func (this IDBIndex) SetName(val js.String) bool {
	return js.True == bindings.SetIDBIndexName(
		this.Ref(),
		val.Ref(),
	)
}

// ObjectStore returns the value of property "IDBIndex.objectStore".
//
// The returned bool will be false if there is no such property.
func (this IDBIndex) ObjectStore() (IDBObjectStore, bool) {
	var _ok bool
	_ret := bindings.GetIDBIndexObjectStore(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBObjectStore{}.FromRef(_ret), _ok
}

// KeyPath returns the value of property "IDBIndex.keyPath".
//
// The returned bool will be false if there is no such property.
func (this IDBIndex) KeyPath() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetIDBIndexKeyPath(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// MultiEntry returns the value of property "IDBIndex.multiEntry".
//
// The returned bool will be false if there is no such property.
func (this IDBIndex) MultiEntry() (bool, bool) {
	var _ok bool
	_ret := bindings.GetIDBIndexMultiEntry(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Unique returns the value of property "IDBIndex.unique".
//
// The returned bool will be false if there is no such property.
func (this IDBIndex) Unique() (bool, bool) {
	var _ok bool
	_ret := bindings.GetIDBIndexUnique(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Get calls the method "IDBIndex.get".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) Get(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexGet(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetFunc returns the method "IDBIndex.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) GetFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetFunc(
			this.Ref(),
		),
	)
}

// GetKey calls the method "IDBIndex.getKey".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) GetKey(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexGetKey(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetKeyFunc returns the method "IDBIndex.getKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) GetKeyFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetKeyFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "IDBIndex.getAll".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) GetAll(query js.Any, count uint32) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexGetAll(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
		uint32(count),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAllFunc returns the method "IDBIndex.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) GetAllFunc() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAllFunc(
			this.Ref(),
		),
	)
}

// GetAll1 calls the method "IDBIndex.getAll".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) GetAll1(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexGetAll1(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAll1Func returns the method "IDBIndex.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) GetAll1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAll1Func(
			this.Ref(),
		),
	)
}

// GetAll2 calls the method "IDBIndex.getAll".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) GetAll2() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexGetAll2(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAll2Func returns the method "IDBIndex.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) GetAll2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAll2Func(
			this.Ref(),
		),
	)
}

// GetAllKeys calls the method "IDBIndex.getAllKeys".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) GetAllKeys(query js.Any, count uint32) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexGetAllKeys(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
		uint32(count),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAllKeysFunc returns the method "IDBIndex.getAllKeys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) GetAllKeysFunc() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAllKeysFunc(
			this.Ref(),
		),
	)
}

// GetAllKeys1 calls the method "IDBIndex.getAllKeys".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) GetAllKeys1(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexGetAllKeys1(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAllKeys1Func returns the method "IDBIndex.getAllKeys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) GetAllKeys1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAllKeys1Func(
			this.Ref(),
		),
	)
}

// GetAllKeys2 calls the method "IDBIndex.getAllKeys".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) GetAllKeys2() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexGetAllKeys2(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAllKeys2Func returns the method "IDBIndex.getAllKeys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) GetAllKeys2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAllKeys2Func(
			this.Ref(),
		),
	)
}

// Count calls the method "IDBIndex.count".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) Count(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexCount(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// CountFunc returns the method "IDBIndex.count".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) CountFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexCountFunc(
			this.Ref(),
		),
	)
}

// Count1 calls the method "IDBIndex.count".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) Count1() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexCount1(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// Count1Func returns the method "IDBIndex.count".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) Count1Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexCount1Func(
			this.Ref(),
		),
	)
}

// OpenCursor calls the method "IDBIndex.openCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) OpenCursor(query js.Any, direction IDBCursorDirection) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexOpenCursor(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
		uint32(direction),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenCursorFunc returns the method "IDBIndex.openCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) OpenCursorFunc() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenCursorFunc(
			this.Ref(),
		),
	)
}

// OpenCursor1 calls the method "IDBIndex.openCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) OpenCursor1(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexOpenCursor1(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenCursor1Func returns the method "IDBIndex.openCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) OpenCursor1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenCursor1Func(
			this.Ref(),
		),
	)
}

// OpenCursor2 calls the method "IDBIndex.openCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) OpenCursor2() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexOpenCursor2(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenCursor2Func returns the method "IDBIndex.openCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) OpenCursor2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenCursor2Func(
			this.Ref(),
		),
	)
}

// OpenKeyCursor calls the method "IDBIndex.openKeyCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) OpenKeyCursor(query js.Any, direction IDBCursorDirection) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexOpenKeyCursor(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
		uint32(direction),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenKeyCursorFunc returns the method "IDBIndex.openKeyCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) OpenKeyCursorFunc() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenKeyCursorFunc(
			this.Ref(),
		),
	)
}

// OpenKeyCursor1 calls the method "IDBIndex.openKeyCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) OpenKeyCursor1(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexOpenKeyCursor1(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenKeyCursor1Func returns the method "IDBIndex.openKeyCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) OpenKeyCursor1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenKeyCursor1Func(
			this.Ref(),
		),
	)
}

// OpenKeyCursor2 calls the method "IDBIndex.openKeyCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBIndex) OpenKeyCursor2() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBIndexOpenKeyCursor2(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenKeyCursor2Func returns the method "IDBIndex.openKeyCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBIndex) OpenKeyCursor2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenKeyCursor2Func(
			this.Ref(),
		),
	)
}

type IDBIndexParameters struct {
	// Unique is "IDBIndexParameters.unique"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_Unique MUST be set to true to make this field effective.
	Unique bool
	// MultiEntry is "IDBIndexParameters.multiEntry"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_MultiEntry MUST be set to true to make this field effective.
	MultiEntry bool

	FFI_USE_Unique     bool // for Unique.
	FFI_USE_MultiEntry bool // for MultiEntry.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IDBIndexParameters with all fields set.
func (p IDBIndexParameters) FromRef(ref js.Ref) IDBIndexParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IDBIndexParameters in the application heap.
func (p IDBIndexParameters) New() js.Ref {
	return bindings.IDBIndexParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IDBIndexParameters) UpdateFrom(ref js.Ref) {
	bindings.IDBIndexParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IDBIndexParameters) Update(ref js.Ref) {
	bindings.IDBIndexParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IDBTransactionMode uint32

const (
	_ IDBTransactionMode = iota

	IDBTransactionMode_READONLY
	IDBTransactionMode_READWRITE
	IDBTransactionMode_VERSIONCHANGE
)

func (IDBTransactionMode) FromRef(str js.Ref) IDBTransactionMode {
	return IDBTransactionMode(bindings.ConstOfIDBTransactionMode(str))
}

func (x IDBTransactionMode) String() (string, bool) {
	switch x {
	case IDBTransactionMode_READONLY:
		return "readonly", true
	case IDBTransactionMode_READWRITE:
		return "readwrite", true
	case IDBTransactionMode_VERSIONCHANGE:
		return "versionchange", true
	default:
		return "", false
	}
}

type IDBTransactionDurability uint32

const (
	_ IDBTransactionDurability = iota

	IDBTransactionDurability_DEFAULT
	IDBTransactionDurability_STRICT
	IDBTransactionDurability_RELAXED
)

func (IDBTransactionDurability) FromRef(str js.Ref) IDBTransactionDurability {
	return IDBTransactionDurability(bindings.ConstOfIDBTransactionDurability(str))
}

func (x IDBTransactionDurability) String() (string, bool) {
	switch x {
	case IDBTransactionDurability_DEFAULT:
		return "default", true
	case IDBTransactionDurability_STRICT:
		return "strict", true
	case IDBTransactionDurability_RELAXED:
		return "relaxed", true
	default:
		return "", false
	}
}

type IDBTransactionOptions struct {
	// Durability is "IDBTransactionOptions.durability"
	//
	// Optional, defaults to "default".
	Durability IDBTransactionDurability

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IDBTransactionOptions with all fields set.
func (p IDBTransactionOptions) FromRef(ref js.Ref) IDBTransactionOptions {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IDBTransactionOptions in the application heap.
func (p IDBTransactionOptions) New() js.Ref {
	return bindings.IDBTransactionOptionsJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IDBTransactionOptions) UpdateFrom(ref js.Ref) {
	bindings.IDBTransactionOptionsJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IDBTransactionOptions) Update(ref js.Ref) {
	bindings.IDBTransactionOptionsJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IDBObjectStoreParameters struct {
	// KeyPath is "IDBObjectStoreParameters.keyPath"
	//
	// Optional, defaults to null.
	KeyPath OneOf_String_ArrayString
	// AutoIncrement is "IDBObjectStoreParameters.autoIncrement"
	//
	// Optional, defaults to false.
	//
	// NOTE: FFI_USE_AutoIncrement MUST be set to true to make this field effective.
	AutoIncrement bool

	FFI_USE_AutoIncrement bool // for AutoIncrement.

	FFI_USE bool
}

// FromRef calls UpdateFrom and returns a IDBObjectStoreParameters with all fields set.
func (p IDBObjectStoreParameters) FromRef(ref js.Ref) IDBObjectStoreParameters {
	p.UpdateFrom(ref)
	return p
}

// New creates a new IDBObjectStoreParameters in the application heap.
func (p IDBObjectStoreParameters) New() js.Ref {
	return bindings.IDBObjectStoreParametersJSLoad(
		js.Pointer(&p), js.True, 0,
	)
}

// UpdateFrom copies value of all fields of the heap object to p.
func (p IDBObjectStoreParameters) UpdateFrom(ref js.Ref) {
	bindings.IDBObjectStoreParametersJSStore(
		js.Pointer(&p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p IDBObjectStoreParameters) Update(ref js.Ref) {
	bindings.IDBObjectStoreParametersJSLoad(
		js.Pointer(&p), js.False, ref,
	)
}

type IDBDatabase struct {
	EventTarget
}

func (this IDBDatabase) Once() IDBDatabase {
	this.Ref().Once()
	return this
}

func (this IDBDatabase) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this IDBDatabase) FromRef(ref js.Ref) IDBDatabase {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this IDBDatabase) Free() {
	this.Ref().Free()
}

// Name returns the value of property "IDBDatabase.name".
//
// The returned bool will be false if there is no such property.
func (this IDBDatabase) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetIDBDatabaseName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Version returns the value of property "IDBDatabase.version".
//
// The returned bool will be false if there is no such property.
func (this IDBDatabase) Version() (uint64, bool) {
	var _ok bool
	_ret := bindings.GetIDBDatabaseVersion(
		this.Ref(), js.Pointer(&_ok),
	)
	return uint64(_ret), _ok
}

// ObjectStoreNames returns the value of property "IDBDatabase.objectStoreNames".
//
// The returned bool will be false if there is no such property.
func (this IDBDatabase) ObjectStoreNames() (DOMStringList, bool) {
	var _ok bool
	_ret := bindings.GetIDBDatabaseObjectStoreNames(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMStringList{}.FromRef(_ret), _ok
}

// Transaction calls the method "IDBDatabase.transaction".
//
// The returned bool will be false if there is no such method.
func (this IDBDatabase) Transaction(storeNames OneOf_String_ArrayString, mode IDBTransactionMode, options IDBTransactionOptions) (IDBTransaction, bool) {
	var _ok bool
	_ret := bindings.CallIDBDatabaseTransaction(
		this.Ref(), js.Pointer(&_ok),
		storeNames.Ref(),
		uint32(mode),
		js.Pointer(&options),
	)

	return IDBTransaction{}.FromRef(_ret), _ok
}

// TransactionFunc returns the method "IDBDatabase.transaction".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBDatabase) TransactionFunc() (fn js.Func[func(storeNames OneOf_String_ArrayString, mode IDBTransactionMode, options IDBTransactionOptions) IDBTransaction]) {
	return fn.FromRef(
		bindings.IDBDatabaseTransactionFunc(
			this.Ref(),
		),
	)
}

// Transaction1 calls the method "IDBDatabase.transaction".
//
// The returned bool will be false if there is no such method.
func (this IDBDatabase) Transaction1(storeNames OneOf_String_ArrayString, mode IDBTransactionMode) (IDBTransaction, bool) {
	var _ok bool
	_ret := bindings.CallIDBDatabaseTransaction1(
		this.Ref(), js.Pointer(&_ok),
		storeNames.Ref(),
		uint32(mode),
	)

	return IDBTransaction{}.FromRef(_ret), _ok
}

// Transaction1Func returns the method "IDBDatabase.transaction".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBDatabase) Transaction1Func() (fn js.Func[func(storeNames OneOf_String_ArrayString, mode IDBTransactionMode) IDBTransaction]) {
	return fn.FromRef(
		bindings.IDBDatabaseTransaction1Func(
			this.Ref(),
		),
	)
}

// Transaction2 calls the method "IDBDatabase.transaction".
//
// The returned bool will be false if there is no such method.
func (this IDBDatabase) Transaction2(storeNames OneOf_String_ArrayString) (IDBTransaction, bool) {
	var _ok bool
	_ret := bindings.CallIDBDatabaseTransaction2(
		this.Ref(), js.Pointer(&_ok),
		storeNames.Ref(),
	)

	return IDBTransaction{}.FromRef(_ret), _ok
}

// Transaction2Func returns the method "IDBDatabase.transaction".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBDatabase) Transaction2Func() (fn js.Func[func(storeNames OneOf_String_ArrayString) IDBTransaction]) {
	return fn.FromRef(
		bindings.IDBDatabaseTransaction2Func(
			this.Ref(),
		),
	)
}

// Close calls the method "IDBDatabase.close".
//
// The returned bool will be false if there is no such method.
func (this IDBDatabase) Close() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIDBDatabaseClose(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CloseFunc returns the method "IDBDatabase.close".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBDatabase) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.IDBDatabaseCloseFunc(
			this.Ref(),
		),
	)
}

// CreateObjectStore calls the method "IDBDatabase.createObjectStore".
//
// The returned bool will be false if there is no such method.
func (this IDBDatabase) CreateObjectStore(name js.String, options IDBObjectStoreParameters) (IDBObjectStore, bool) {
	var _ok bool
	_ret := bindings.CallIDBDatabaseCreateObjectStore(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		js.Pointer(&options),
	)

	return IDBObjectStore{}.FromRef(_ret), _ok
}

// CreateObjectStoreFunc returns the method "IDBDatabase.createObjectStore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBDatabase) CreateObjectStoreFunc() (fn js.Func[func(name js.String, options IDBObjectStoreParameters) IDBObjectStore]) {
	return fn.FromRef(
		bindings.IDBDatabaseCreateObjectStoreFunc(
			this.Ref(),
		),
	)
}

// CreateObjectStore1 calls the method "IDBDatabase.createObjectStore".
//
// The returned bool will be false if there is no such method.
func (this IDBDatabase) CreateObjectStore1(name js.String) (IDBObjectStore, bool) {
	var _ok bool
	_ret := bindings.CallIDBDatabaseCreateObjectStore1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return IDBObjectStore{}.FromRef(_ret), _ok
}

// CreateObjectStore1Func returns the method "IDBDatabase.createObjectStore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBDatabase) CreateObjectStore1Func() (fn js.Func[func(name js.String) IDBObjectStore]) {
	return fn.FromRef(
		bindings.IDBDatabaseCreateObjectStore1Func(
			this.Ref(),
		),
	)
}

// DeleteObjectStore calls the method "IDBDatabase.deleteObjectStore".
//
// The returned bool will be false if there is no such method.
func (this IDBDatabase) DeleteObjectStore(name js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIDBDatabaseDeleteObjectStore(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteObjectStoreFunc returns the method "IDBDatabase.deleteObjectStore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBDatabase) DeleteObjectStoreFunc() (fn js.Func[func(name js.String)]) {
	return fn.FromRef(
		bindings.IDBDatabaseDeleteObjectStoreFunc(
			this.Ref(),
		),
	)
}

type IDBTransaction struct {
	EventTarget
}

func (this IDBTransaction) Once() IDBTransaction {
	this.Ref().Once()
	return this
}

func (this IDBTransaction) Ref() js.Ref {
	return this.EventTarget.Ref()
}

func (this IDBTransaction) FromRef(ref js.Ref) IDBTransaction {
	this.EventTarget = this.EventTarget.FromRef(ref)
	return this
}

func (this IDBTransaction) Free() {
	this.Ref().Free()
}

// ObjectStoreNames returns the value of property "IDBTransaction.objectStoreNames".
//
// The returned bool will be false if there is no such property.
func (this IDBTransaction) ObjectStoreNames() (DOMStringList, bool) {
	var _ok bool
	_ret := bindings.GetIDBTransactionObjectStoreNames(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMStringList{}.FromRef(_ret), _ok
}

// Mode returns the value of property "IDBTransaction.mode".
//
// The returned bool will be false if there is no such property.
func (this IDBTransaction) Mode() (IDBTransactionMode, bool) {
	var _ok bool
	_ret := bindings.GetIDBTransactionMode(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBTransactionMode(_ret), _ok
}

// Durability returns the value of property "IDBTransaction.durability".
//
// The returned bool will be false if there is no such property.
func (this IDBTransaction) Durability() (IDBTransactionDurability, bool) {
	var _ok bool
	_ret := bindings.GetIDBTransactionDurability(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBTransactionDurability(_ret), _ok
}

// Db returns the value of property "IDBTransaction.db".
//
// The returned bool will be false if there is no such property.
func (this IDBTransaction) Db() (IDBDatabase, bool) {
	var _ok bool
	_ret := bindings.GetIDBTransactionDb(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBDatabase{}.FromRef(_ret), _ok
}

// Error returns the value of property "IDBTransaction.error".
//
// The returned bool will be false if there is no such property.
func (this IDBTransaction) Error() (DOMException, bool) {
	var _ok bool
	_ret := bindings.GetIDBTransactionError(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMException{}.FromRef(_ret), _ok
}

// ObjectStore calls the method "IDBTransaction.objectStore".
//
// The returned bool will be false if there is no such method.
func (this IDBTransaction) ObjectStore(name js.String) (IDBObjectStore, bool) {
	var _ok bool
	_ret := bindings.CallIDBTransactionObjectStore(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return IDBObjectStore{}.FromRef(_ret), _ok
}

// ObjectStoreFunc returns the method "IDBTransaction.objectStore".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBTransaction) ObjectStoreFunc() (fn js.Func[func(name js.String) IDBObjectStore]) {
	return fn.FromRef(
		bindings.IDBTransactionObjectStoreFunc(
			this.Ref(),
		),
	)
}

// Commit calls the method "IDBTransaction.commit".
//
// The returned bool will be false if there is no such method.
func (this IDBTransaction) Commit() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIDBTransactionCommit(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// CommitFunc returns the method "IDBTransaction.commit".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBTransaction) CommitFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.IDBTransactionCommitFunc(
			this.Ref(),
		),
	)
}

// Abort calls the method "IDBTransaction.abort".
//
// The returned bool will be false if there is no such method.
func (this IDBTransaction) Abort() (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIDBTransactionAbort(
		this.Ref(), js.Pointer(&_ok),
	)

	_ = _ret
	return js.Void{}, _ok
}

// AbortFunc returns the method "IDBTransaction.abort".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBTransaction) AbortFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.IDBTransactionAbortFunc(
			this.Ref(),
		),
	)
}

type IDBObjectStore struct {
	ref js.Ref
}

func (this IDBObjectStore) Once() IDBObjectStore {
	this.Ref().Once()
	return this
}

func (this IDBObjectStore) Ref() js.Ref {
	return this.ref
}

func (this IDBObjectStore) FromRef(ref js.Ref) IDBObjectStore {
	this.ref = ref
	return this
}

func (this IDBObjectStore) Free() {
	this.Ref().Free()
}

// Name returns the value of property "IDBObjectStore.name".
//
// The returned bool will be false if there is no such property.
func (this IDBObjectStore) Name() (js.String, bool) {
	var _ok bool
	_ret := bindings.GetIDBObjectStoreName(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.String{}.FromRef(_ret), _ok
}

// Name sets the value of property "IDBObjectStore.name" to val.
//
// It returns false if the property cannot be set.
func (this IDBObjectStore) SetName(val js.String) bool {
	return js.True == bindings.SetIDBObjectStoreName(
		this.Ref(),
		val.Ref(),
	)
}

// KeyPath returns the value of property "IDBObjectStore.keyPath".
//
// The returned bool will be false if there is no such property.
func (this IDBObjectStore) KeyPath() (js.Any, bool) {
	var _ok bool
	_ret := bindings.GetIDBObjectStoreKeyPath(
		this.Ref(), js.Pointer(&_ok),
	)
	return js.Any{}.FromRef(_ret), _ok
}

// IndexNames returns the value of property "IDBObjectStore.indexNames".
//
// The returned bool will be false if there is no such property.
func (this IDBObjectStore) IndexNames() (DOMStringList, bool) {
	var _ok bool
	_ret := bindings.GetIDBObjectStoreIndexNames(
		this.Ref(), js.Pointer(&_ok),
	)
	return DOMStringList{}.FromRef(_ret), _ok
}

// Transaction returns the value of property "IDBObjectStore.transaction".
//
// The returned bool will be false if there is no such property.
func (this IDBObjectStore) Transaction() (IDBTransaction, bool) {
	var _ok bool
	_ret := bindings.GetIDBObjectStoreTransaction(
		this.Ref(), js.Pointer(&_ok),
	)
	return IDBTransaction{}.FromRef(_ret), _ok
}

// AutoIncrement returns the value of property "IDBObjectStore.autoIncrement".
//
// The returned bool will be false if there is no such property.
func (this IDBObjectStore) AutoIncrement() (bool, bool) {
	var _ok bool
	_ret := bindings.GetIDBObjectStoreAutoIncrement(
		this.Ref(), js.Pointer(&_ok),
	)
	return _ret == js.True, _ok
}

// Put calls the method "IDBObjectStore.put".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) Put(value js.Any, key js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStorePut(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
		key.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// PutFunc returns the method "IDBObjectStore.put".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) PutFunc() (fn js.Func[func(value js.Any, key js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStorePutFunc(
			this.Ref(),
		),
	)
}

// Put1 calls the method "IDBObjectStore.put".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) Put1(value js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStorePut1(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// Put1Func returns the method "IDBObjectStore.put".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) Put1Func() (fn js.Func[func(value js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStorePut1Func(
			this.Ref(),
		),
	)
}

// Add calls the method "IDBObjectStore.add".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) Add(value js.Any, key js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreAdd(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
		key.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// AddFunc returns the method "IDBObjectStore.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) AddFunc() (fn js.Func[func(value js.Any, key js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreAddFunc(
			this.Ref(),
		),
	)
}

// Add1 calls the method "IDBObjectStore.add".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) Add1(value js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreAdd1(
		this.Ref(), js.Pointer(&_ok),
		value.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// Add1Func returns the method "IDBObjectStore.add".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) Add1Func() (fn js.Func[func(value js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreAdd1Func(
			this.Ref(),
		),
	)
}

// Delete calls the method "IDBObjectStore.delete".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) Delete(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreDelete(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// DeleteFunc returns the method "IDBObjectStore.delete".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) DeleteFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreDeleteFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "IDBObjectStore.clear".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) Clear() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreClear(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// ClearFunc returns the method "IDBObjectStore.clear".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) ClearFunc() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreClearFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "IDBObjectStore.get".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) Get(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreGet(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetFunc returns the method "IDBObjectStore.get".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) GetFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetFunc(
			this.Ref(),
		),
	)
}

// GetKey calls the method "IDBObjectStore.getKey".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) GetKey(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreGetKey(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetKeyFunc returns the method "IDBObjectStore.getKey".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) GetKeyFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetKeyFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "IDBObjectStore.getAll".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) GetAll(query js.Any, count uint32) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreGetAll(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
		uint32(count),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAllFunc returns the method "IDBObjectStore.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) GetAllFunc() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAllFunc(
			this.Ref(),
		),
	)
}

// GetAll1 calls the method "IDBObjectStore.getAll".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) GetAll1(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreGetAll1(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAll1Func returns the method "IDBObjectStore.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) GetAll1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAll1Func(
			this.Ref(),
		),
	)
}

// GetAll2 calls the method "IDBObjectStore.getAll".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) GetAll2() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreGetAll2(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAll2Func returns the method "IDBObjectStore.getAll".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) GetAll2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAll2Func(
			this.Ref(),
		),
	)
}

// GetAllKeys calls the method "IDBObjectStore.getAllKeys".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) GetAllKeys(query js.Any, count uint32) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreGetAllKeys(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
		uint32(count),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAllKeysFunc returns the method "IDBObjectStore.getAllKeys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) GetAllKeysFunc() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAllKeysFunc(
			this.Ref(),
		),
	)
}

// GetAllKeys1 calls the method "IDBObjectStore.getAllKeys".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) GetAllKeys1(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreGetAllKeys1(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAllKeys1Func returns the method "IDBObjectStore.getAllKeys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) GetAllKeys1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAllKeys1Func(
			this.Ref(),
		),
	)
}

// GetAllKeys2 calls the method "IDBObjectStore.getAllKeys".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) GetAllKeys2() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreGetAllKeys2(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// GetAllKeys2Func returns the method "IDBObjectStore.getAllKeys".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) GetAllKeys2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAllKeys2Func(
			this.Ref(),
		),
	)
}

// Count calls the method "IDBObjectStore.count".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) Count(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreCount(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// CountFunc returns the method "IDBObjectStore.count".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) CountFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreCountFunc(
			this.Ref(),
		),
	)
}

// Count1 calls the method "IDBObjectStore.count".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) Count1() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreCount1(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// Count1Func returns the method "IDBObjectStore.count".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) Count1Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreCount1Func(
			this.Ref(),
		),
	)
}

// OpenCursor calls the method "IDBObjectStore.openCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) OpenCursor(query js.Any, direction IDBCursorDirection) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreOpenCursor(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
		uint32(direction),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenCursorFunc returns the method "IDBObjectStore.openCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) OpenCursorFunc() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenCursorFunc(
			this.Ref(),
		),
	)
}

// OpenCursor1 calls the method "IDBObjectStore.openCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) OpenCursor1(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreOpenCursor1(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenCursor1Func returns the method "IDBObjectStore.openCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) OpenCursor1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenCursor1Func(
			this.Ref(),
		),
	)
}

// OpenCursor2 calls the method "IDBObjectStore.openCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) OpenCursor2() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreOpenCursor2(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenCursor2Func returns the method "IDBObjectStore.openCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) OpenCursor2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenCursor2Func(
			this.Ref(),
		),
	)
}

// OpenKeyCursor calls the method "IDBObjectStore.openKeyCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) OpenKeyCursor(query js.Any, direction IDBCursorDirection) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreOpenKeyCursor(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
		uint32(direction),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenKeyCursorFunc returns the method "IDBObjectStore.openKeyCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) OpenKeyCursorFunc() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenKeyCursorFunc(
			this.Ref(),
		),
	)
}

// OpenKeyCursor1 calls the method "IDBObjectStore.openKeyCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) OpenKeyCursor1(query js.Any) (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreOpenKeyCursor1(
		this.Ref(), js.Pointer(&_ok),
		query.Ref(),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenKeyCursor1Func returns the method "IDBObjectStore.openKeyCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) OpenKeyCursor1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenKeyCursor1Func(
			this.Ref(),
		),
	)
}

// OpenKeyCursor2 calls the method "IDBObjectStore.openKeyCursor".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) OpenKeyCursor2() (IDBRequest, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreOpenKeyCursor2(
		this.Ref(), js.Pointer(&_ok),
	)

	return IDBRequest{}.FromRef(_ret), _ok
}

// OpenKeyCursor2Func returns the method "IDBObjectStore.openKeyCursor".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) OpenKeyCursor2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenKeyCursor2Func(
			this.Ref(),
		),
	)
}

// Index calls the method "IDBObjectStore.index".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) Index(name js.String) (IDBIndex, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreIndex(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	return IDBIndex{}.FromRef(_ret), _ok
}

// IndexFunc returns the method "IDBObjectStore.index".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) IndexFunc() (fn js.Func[func(name js.String) IDBIndex]) {
	return fn.FromRef(
		bindings.IDBObjectStoreIndexFunc(
			this.Ref(),
		),
	)
}

// CreateIndex calls the method "IDBObjectStore.createIndex".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) CreateIndex(name js.String, keyPath OneOf_String_ArrayString, options IDBIndexParameters) (IDBIndex, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreCreateIndex(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		keyPath.Ref(),
		js.Pointer(&options),
	)

	return IDBIndex{}.FromRef(_ret), _ok
}

// CreateIndexFunc returns the method "IDBObjectStore.createIndex".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) CreateIndexFunc() (fn js.Func[func(name js.String, keyPath OneOf_String_ArrayString, options IDBIndexParameters) IDBIndex]) {
	return fn.FromRef(
		bindings.IDBObjectStoreCreateIndexFunc(
			this.Ref(),
		),
	)
}

// CreateIndex1 calls the method "IDBObjectStore.createIndex".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) CreateIndex1(name js.String, keyPath OneOf_String_ArrayString) (IDBIndex, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreCreateIndex1(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
		keyPath.Ref(),
	)

	return IDBIndex{}.FromRef(_ret), _ok
}

// CreateIndex1Func returns the method "IDBObjectStore.createIndex".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) CreateIndex1Func() (fn js.Func[func(name js.String, keyPath OneOf_String_ArrayString) IDBIndex]) {
	return fn.FromRef(
		bindings.IDBObjectStoreCreateIndex1Func(
			this.Ref(),
		),
	)
}

// DeleteIndex calls the method "IDBObjectStore.deleteIndex".
//
// The returned bool will be false if there is no such method.
func (this IDBObjectStore) DeleteIndex(name js.String) (js.Void, bool) {
	var _ok bool
	_ret := bindings.CallIDBObjectStoreDeleteIndex(
		this.Ref(), js.Pointer(&_ok),
		name.Ref(),
	)

	_ = _ret
	return js.Void{}, _ok
}

// DeleteIndexFunc returns the method "IDBObjectStore.deleteIndex".
//
// The ref value of the returned js.Func will be js.Undefined if there is no such method.
func (this IDBObjectStore) DeleteIndexFunc() (fn js.Func[func(name js.String)]) {
	return fn.FromRef(
		bindings.IDBObjectStoreDeleteIndexFunc(
			this.Ref(),
		),
	)
}

type OneOf_IDBObjectStore_IDBIndex_IDBCursor struct {
	ref js.Ref
}

func (x OneOf_IDBObjectStore_IDBIndex_IDBCursor) Ref() js.Ref {
	return x.ref
}

func (x OneOf_IDBObjectStore_IDBIndex_IDBCursor) Free() {
	x.ref.Free()
}

func (x OneOf_IDBObjectStore_IDBIndex_IDBCursor) FromRef(ref js.Ref) OneOf_IDBObjectStore_IDBIndex_IDBCursor {
	return OneOf_IDBObjectStore_IDBIndex_IDBCursor{
		ref: ref,
	}
}

func (x OneOf_IDBObjectStore_IDBIndex_IDBCursor) IDBObjectStore() IDBObjectStore {
	return IDBObjectStore{}.FromRef(x.ref)
}

func (x OneOf_IDBObjectStore_IDBIndex_IDBCursor) IDBIndex() IDBIndex {
	return IDBIndex{}.FromRef(x.ref)
}

func (x OneOf_IDBObjectStore_IDBIndex_IDBCursor) IDBCursor() IDBCursor {
	return IDBCursor{}.FromRef(x.ref)
}

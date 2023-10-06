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

func NewHTMLPortalElement() (ret HTMLPortalElement) {
	ret.ref = bindings.NewHTMLPortalElementByHTMLPortalElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLPortalElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLPortalElementSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLPortalElement.src" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLPortalElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLPortalElementReferrerPolicy(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLPortalElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLPortalElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLPortalElementReferrerPolicy(
		this.Ref(),
		val.Ref(),
	)
}

// HasActivate returns true if the method "HTMLPortalElement.activate" exists.
func (this HTMLPortalElement) HasActivate() bool {
	return js.True == bindings.HasHTMLPortalElementActivate(
		this.Ref(),
	)
}

// ActivateFunc returns the method "HTMLPortalElement.activate".
func (this HTMLPortalElement) ActivateFunc() (fn js.Func[func(options PortalActivateOptions) js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLPortalElementActivateFunc(
			this.Ref(),
		),
	)
}

// Activate calls the method "HTMLPortalElement.activate".
func (this HTMLPortalElement) Activate(options PortalActivateOptions) (ret js.Promise[js.Void]) {
	bindings.CallHTMLPortalElementActivate(
		this.Ref(), js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryActivate calls the method "HTMLPortalElement.activate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLPortalElement) TryActivate(options PortalActivateOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLPortalElementActivate(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasActivate1 returns true if the method "HTMLPortalElement.activate" exists.
func (this HTMLPortalElement) HasActivate1() bool {
	return js.True == bindings.HasHTMLPortalElementActivate1(
		this.Ref(),
	)
}

// Activate1Func returns the method "HTMLPortalElement.activate".
func (this HTMLPortalElement) Activate1Func() (fn js.Func[func() js.Promise[js.Void]]) {
	return fn.FromRef(
		bindings.HTMLPortalElementActivate1Func(
			this.Ref(),
		),
	)
}

// Activate1 calls the method "HTMLPortalElement.activate".
func (this HTMLPortalElement) Activate1() (ret js.Promise[js.Void]) {
	bindings.CallHTMLPortalElementActivate1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryActivate1 calls the method "HTMLPortalElement.activate"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLPortalElement) TryActivate1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLPortalElementActivate1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasPostMessage returns true if the method "HTMLPortalElement.postMessage" exists.
func (this HTMLPortalElement) HasPostMessage() bool {
	return js.True == bindings.HasHTMLPortalElementPostMessage(
		this.Ref(),
	)
}

// PostMessageFunc returns the method "HTMLPortalElement.postMessage".
func (this HTMLPortalElement) PostMessageFunc() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	return fn.FromRef(
		bindings.HTMLPortalElementPostMessageFunc(
			this.Ref(),
		),
	)
}

// PostMessage calls the method "HTMLPortalElement.postMessage".
func (this HTMLPortalElement) PostMessage(message js.Any, options StructuredSerializeOptions) (ret js.Void) {
	bindings.CallHTMLPortalElementPostMessage(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage calls the method "HTMLPortalElement.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLPortalElement) TryPostMessage(message js.Any, options StructuredSerializeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLPortalElementPostMessage(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasPostMessage1 returns true if the method "HTMLPortalElement.postMessage" exists.
func (this HTMLPortalElement) HasPostMessage1() bool {
	return js.True == bindings.HasHTMLPortalElementPostMessage1(
		this.Ref(),
	)
}

// PostMessage1Func returns the method "HTMLPortalElement.postMessage".
func (this HTMLPortalElement) PostMessage1Func() (fn js.Func[func(message js.Any)]) {
	return fn.FromRef(
		bindings.HTMLPortalElementPostMessage1Func(
			this.Ref(),
		),
	)
}

// PostMessage1 calls the method "HTMLPortalElement.postMessage".
func (this HTMLPortalElement) PostMessage1(message js.Any) (ret js.Void) {
	bindings.CallHTMLPortalElementPostMessage1(
		this.Ref(), js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage1 calls the method "HTMLPortalElement.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLPortalElement) TryPostMessage1(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLPortalElementPostMessage1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

func NewHTMLPreElement() (ret HTMLPreElement) {
	ret.ref = bindings.NewHTMLPreElementByHTMLPreElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLPreElement) Width() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLPreElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLPreElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLPreElement) SetWidth(val int32) bool {
	return js.True == bindings.SetHTMLPreElementWidth(
		this.Ref(),
		int32(val),
	)
}

func NewHTMLProgressElement() (ret HTMLProgressElement) {
	ret.ref = bindings.NewHTMLProgressElementByHTMLProgressElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLProgressElement) Value() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLProgressElementValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLProgressElement.value" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLProgressElement) Max() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLProgressElementMax(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMax sets the value of property "HTMLProgressElement.max" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLProgressElement) Position() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLProgressElementPosition(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLProgressElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLProgressElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLProgressElementLabels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewHTMLQuoteElement() (ret HTMLQuoteElement) {
	ret.ref = bindings.NewHTMLQuoteElementByHTMLQuoteElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLQuoteElement) Cite() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLQuoteElementCite(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCite sets the value of property "HTMLQuoteElement.cite" to val.
//
// It returns false if the property cannot be set.
func (this HTMLQuoteElement) SetCite(val js.String) bool {
	return js.True == bindings.SetHTMLQuoteElementCite(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLSelectElement() (ret HTMLSelectElement) {
	ret.ref = bindings.NewHTMLSelectElementByHTMLSelectElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Autocomplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementAutocomplete(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAutocomplete sets the value of property "HTMLSelectElement.autocomplete" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementDisabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLSelectElement.disabled" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementForm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Multiple returns the value of property "HTMLSelectElement.multiple".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Multiple() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementMultiple(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMultiple sets the value of property "HTMLSelectElement.multiple" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLSelectElement.name" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Required() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementRequired(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRequired sets the value of property "HTMLSelectElement.required" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Size() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementSize(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSize sets the value of property "HTMLSelectElement.size" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Options returns the value of property "HTMLSelectElement.options".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Options() (ret HTMLOptionsCollection, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementOptions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "HTMLSelectElement.length".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLength sets the value of property "HTMLSelectElement.length" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) SelectedOptions() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementSelectedOptions(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SelectedIndex returns the value of property "HTMLSelectElement.selectedIndex".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) SelectedIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementSelectedIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelectedIndex sets the value of property "HTMLSelectElement.selectedIndex" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLSelectElement.value" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementWillValidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLSelectElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementValidity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLSelectElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementValidationMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLSelectElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementLabels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasItem returns true if the method "HTMLSelectElement.item" exists.
func (this HTMLSelectElement) HasItem() bool {
	return js.True == bindings.HasHTMLSelectElementItem(
		this.Ref(),
	)
}

// ItemFunc returns the method "HTMLSelectElement.item".
func (this HTMLSelectElement) ItemFunc() (fn js.Func[func(index uint32) HTMLOptionElement]) {
	return fn.FromRef(
		bindings.HTMLSelectElementItemFunc(
			this.Ref(),
		),
	)
}

// Item calls the method "HTMLSelectElement.item".
func (this HTMLSelectElement) Item(index uint32) (ret HTMLOptionElement) {
	bindings.CallHTMLSelectElementItem(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "HTMLSelectElement.item"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSelectElement) TryItem(index uint32) (ret HTMLOptionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasNamedItem returns true if the method "HTMLSelectElement.namedItem" exists.
func (this HTMLSelectElement) HasNamedItem() bool {
	return js.True == bindings.HasHTMLSelectElementNamedItem(
		this.Ref(),
	)
}

// NamedItemFunc returns the method "HTMLSelectElement.namedItem".
func (this HTMLSelectElement) NamedItemFunc() (fn js.Func[func(name js.String) HTMLOptionElement]) {
	return fn.FromRef(
		bindings.HTMLSelectElementNamedItemFunc(
			this.Ref(),
		),
	)
}

// NamedItem calls the method "HTMLSelectElement.namedItem".
func (this HTMLSelectElement) NamedItem(name js.String) (ret HTMLOptionElement) {
	bindings.CallHTMLSelectElementNamedItem(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "HTMLSelectElement.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSelectElement) TryNamedItem(name js.String) (ret HTMLOptionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementNamedItem(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasAdd returns true if the method "HTMLSelectElement.add" exists.
func (this HTMLSelectElement) HasAdd() bool {
	return js.True == bindings.HasHTMLSelectElementAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "HTMLSelectElement.add".
func (this HTMLSelectElement) AddFunc() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32)]) {
	return fn.FromRef(
		bindings.HTMLSelectElementAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "HTMLSelectElement.add".
func (this HTMLSelectElement) Add(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32) (ret js.Void) {
	bindings.CallHTMLSelectElementAdd(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
		before.Ref(),
	)

	return
}

// TryAdd calls the method "HTMLSelectElement.add"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSelectElement) TryAdd(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementAdd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
		before.Ref(),
	)

	return
}

// HasAdd1 returns true if the method "HTMLSelectElement.add" exists.
func (this HTMLSelectElement) HasAdd1() bool {
	return js.True == bindings.HasHTMLSelectElementAdd1(
		this.Ref(),
	)
}

// Add1Func returns the method "HTMLSelectElement.add".
func (this HTMLSelectElement) Add1Func() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement)]) {
	return fn.FromRef(
		bindings.HTMLSelectElementAdd1Func(
			this.Ref(),
		),
	)
}

// Add1 calls the method "HTMLSelectElement.add".
func (this HTMLSelectElement) Add1(element OneOf_HTMLOptionElement_HTMLOptGroupElement) (ret js.Void) {
	bindings.CallHTMLSelectElementAdd1(
		this.Ref(), js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryAdd1 calls the method "HTMLSelectElement.add"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSelectElement) TryAdd1(element OneOf_HTMLOptionElement_HTMLOptGroupElement) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementAdd1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

// HasRemove returns true if the method "HTMLSelectElement.remove" exists.
func (this HTMLSelectElement) HasRemove() bool {
	return js.True == bindings.HasHTMLSelectElementRemove(
		this.Ref(),
	)
}

// RemoveFunc returns the method "HTMLSelectElement.remove".
func (this HTMLSelectElement) RemoveFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLSelectElementRemoveFunc(
			this.Ref(),
		),
	)
}

// Remove calls the method "HTMLSelectElement.remove".
func (this HTMLSelectElement) Remove() (ret js.Void) {
	bindings.CallHTMLSelectElementRemove(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryRemove calls the method "HTMLSelectElement.remove"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSelectElement) TryRemove() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementRemove(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasRemove1 returns true if the method "HTMLSelectElement.remove" exists.
func (this HTMLSelectElement) HasRemove1() bool {
	return js.True == bindings.HasHTMLSelectElementRemove1(
		this.Ref(),
	)
}

// Remove1Func returns the method "HTMLSelectElement.remove".
func (this HTMLSelectElement) Remove1Func() (fn js.Func[func(index int32)]) {
	return fn.FromRef(
		bindings.HTMLSelectElementRemove1Func(
			this.Ref(),
		),
	)
}

// Remove1 calls the method "HTMLSelectElement.remove".
func (this HTMLSelectElement) Remove1(index int32) (ret js.Void) {
	bindings.CallHTMLSelectElementRemove1(
		this.Ref(), js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryRemove1 calls the method "HTMLSelectElement.remove"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSelectElement) TryRemove1(index int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementRemove1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

// HasSet returns true if the method "HTMLSelectElement." exists.
func (this HTMLSelectElement) HasSet() bool {
	return js.True == bindings.HasHTMLSelectElementSet(
		this.Ref(),
	)
}

// SetFunc returns the method "HTMLSelectElement.".
func (this HTMLSelectElement) SetFunc() (fn js.Func[func(index uint32, option HTMLOptionElement)]) {
	return fn.FromRef(
		bindings.HTMLSelectElementSetFunc(
			this.Ref(),
		),
	)
}

// Set calls the method "HTMLSelectElement.".
func (this HTMLSelectElement) Set(index uint32, option HTMLOptionElement) (ret js.Void) {
	bindings.CallHTMLSelectElementSet(
		this.Ref(), js.Pointer(&ret),
		uint32(index),
		option.Ref(),
	)

	return
}

// TrySet calls the method "HTMLSelectElement."
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSelectElement) TrySet(index uint32, option HTMLOptionElement) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementSet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		option.Ref(),
	)

	return
}

// HasCheckValidity returns true if the method "HTMLSelectElement.checkValidity" exists.
func (this HTMLSelectElement) HasCheckValidity() bool {
	return js.True == bindings.HasHTMLSelectElementCheckValidity(
		this.Ref(),
	)
}

// CheckValidityFunc returns the method "HTMLSelectElement.checkValidity".
func (this HTMLSelectElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLSelectElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLSelectElement.checkValidity".
func (this HTMLSelectElement) CheckValidity() (ret bool) {
	bindings.CallHTMLSelectElementCheckValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLSelectElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSelectElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementCheckValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReportValidity returns true if the method "HTMLSelectElement.reportValidity" exists.
func (this HTMLSelectElement) HasReportValidity() bool {
	return js.True == bindings.HasHTMLSelectElementReportValidity(
		this.Ref(),
	)
}

// ReportValidityFunc returns the method "HTMLSelectElement.reportValidity".
func (this HTMLSelectElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLSelectElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLSelectElement.reportValidity".
func (this HTMLSelectElement) ReportValidity() (ret bool) {
	bindings.CallHTMLSelectElementReportValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLSelectElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSelectElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementReportValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetCustomValidity returns true if the method "HTMLSelectElement.setCustomValidity" exists.
func (this HTMLSelectElement) HasSetCustomValidity() bool {
	return js.True == bindings.HasHTMLSelectElementSetCustomValidity(
		this.Ref(),
	)
}

// SetCustomValidityFunc returns the method "HTMLSelectElement.setCustomValidity".
func (this HTMLSelectElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLSelectElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLSelectElement.setCustomValidity".
func (this HTMLSelectElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLSelectElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLSelectElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLSelectElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

func NewHTMLSourceElement() (ret HTMLSourceElement) {
	ret.ref = bindings.NewHTMLSourceElementByHTMLSourceElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLSourceElement.src" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLSourceElement.type" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Srcset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementSrcset(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrcset sets the value of property "HTMLSourceElement.srcset" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Sizes() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementSizes(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSizes sets the value of property "HTMLSourceElement.sizes" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementMedia(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMedia sets the value of property "HTMLSourceElement.media" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLSourceElement.width" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLSourceElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLSourceElementHeight(
		this.Ref(),
		uint32(val),
	)
}

func NewHTMLSpanElement() (ret HTMLSpanElement) {
	ret.ref = bindings.NewHTMLSpanElementByHTMLSpanElement()
	return
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

func NewHTMLStyleElement() (ret HTMLStyleElement) {
	ret.ref = bindings.NewHTMLStyleElementByHTMLStyleElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLStyleElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLStyleElementDisabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLStyleElement.disabled" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLStyleElement) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLStyleElementMedia(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMedia sets the value of property "HTMLStyleElement.media" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLStyleElement) Blocking() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLStyleElementBlocking(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "HTMLStyleElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLStyleElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLStyleElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLStyleElement.type" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLStyleElement) Sheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetHTMLStyleElementSheet(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewHTMLTableCaptionElement() (ret HTMLTableCaptionElement) {
	ret.ref = bindings.NewHTMLTableCaptionElementByHTMLTableCaptionElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLTableCaptionElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCaptionElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableCaptionElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCaptionElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableCaptionElementAlign(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLTableCellElement() (ret HTMLTableCellElement) {
	ret.ref = bindings.NewHTMLTableCellElementByHTMLTableCellElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) ColSpan() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementColSpan(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetColSpan sets the value of property "HTMLTableCellElement.colSpan" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) RowSpan() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementRowSpan(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRowSpan sets the value of property "HTMLTableCellElement.rowSpan" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Headers() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementHeaders(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeaders sets the value of property "HTMLTableCellElement.headers" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) CellIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementCellIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Scope returns the value of property "HTMLTableCellElement.scope".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Scope() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementScope(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetScope sets the value of property "HTMLTableCellElement.scope" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Abbr() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementAbbr(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAbbr sets the value of property "HTMLTableCellElement.abbr" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableCellElement.align" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Axis() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementAxis(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAxis sets the value of property "HTMLTableCellElement.axis" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementHeight(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLTableCellElement.height" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLTableCellElement.width" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Ch() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementCh(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCh sets the value of property "HTMLTableCellElement.ch" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) ChOff() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementChOff(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetChOff sets the value of property "HTMLTableCellElement.chOff" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) NoWrap() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementNoWrap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetNoWrap sets the value of property "HTMLTableCellElement.noWrap" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) VAlign() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementVAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVAlign sets the value of property "HTMLTableCellElement.vAlign" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementBgColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "HTMLTableCellElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementBgColor(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLTableColElement() (ret HTMLTableColElement) {
	ret.ref = bindings.NewHTMLTableColElementByHTMLTableColElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) Span() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementSpan(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSpan sets the value of property "HTMLTableColElement.span" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableColElement.align" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) Ch() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementCh(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCh sets the value of property "HTMLTableColElement.ch" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) ChOff() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementChOff(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetChOff sets the value of property "HTMLTableColElement.chOff" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) VAlign() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementVAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVAlign sets the value of property "HTMLTableColElement.vAlign" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLTableColElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLTableColElementWidth(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLTableRowElement() (ret HTMLTableRowElement) {
	ret.ref = bindings.NewHTMLTableRowElementByHTMLTableRowElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) RowIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementRowIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SectionRowIndex returns the value of property "HTMLTableRowElement.sectionRowIndex".
//
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) SectionRowIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementSectionRowIndex(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Cells returns the value of property "HTMLTableRowElement.cells".
//
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) Cells() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementCells(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Align returns the value of property "HTMLTableRowElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableRowElement.align" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) Ch() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementCh(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCh sets the value of property "HTMLTableRowElement.ch" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) ChOff() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementChOff(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetChOff sets the value of property "HTMLTableRowElement.chOff" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) VAlign() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementVAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVAlign sets the value of property "HTMLTableRowElement.vAlign" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementBgColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "HTMLTableRowElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableRowElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLTableRowElementBgColor(
		this.Ref(),
		val.Ref(),
	)
}

// HasInsertCell returns true if the method "HTMLTableRowElement.insertCell" exists.
func (this HTMLTableRowElement) HasInsertCell() bool {
	return js.True == bindings.HasHTMLTableRowElementInsertCell(
		this.Ref(),
	)
}

// InsertCellFunc returns the method "HTMLTableRowElement.insertCell".
func (this HTMLTableRowElement) InsertCellFunc() (fn js.Func[func(index int32) HTMLTableCellElement]) {
	return fn.FromRef(
		bindings.HTMLTableRowElementInsertCellFunc(
			this.Ref(),
		),
	)
}

// InsertCell calls the method "HTMLTableRowElement.insertCell".
func (this HTMLTableRowElement) InsertCell(index int32) (ret HTMLTableCellElement) {
	bindings.CallHTMLTableRowElementInsertCell(
		this.Ref(), js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryInsertCell calls the method "HTMLTableRowElement.insertCell"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableRowElement) TryInsertCell(index int32) (ret HTMLTableCellElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableRowElementInsertCell(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

// HasInsertCell1 returns true if the method "HTMLTableRowElement.insertCell" exists.
func (this HTMLTableRowElement) HasInsertCell1() bool {
	return js.True == bindings.HasHTMLTableRowElementInsertCell1(
		this.Ref(),
	)
}

// InsertCell1Func returns the method "HTMLTableRowElement.insertCell".
func (this HTMLTableRowElement) InsertCell1Func() (fn js.Func[func() HTMLTableCellElement]) {
	return fn.FromRef(
		bindings.HTMLTableRowElementInsertCell1Func(
			this.Ref(),
		),
	)
}

// InsertCell1 calls the method "HTMLTableRowElement.insertCell".
func (this HTMLTableRowElement) InsertCell1() (ret HTMLTableCellElement) {
	bindings.CallHTMLTableRowElementInsertCell1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryInsertCell1 calls the method "HTMLTableRowElement.insertCell"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableRowElement) TryInsertCell1() (ret HTMLTableCellElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableRowElementInsertCell1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteCell returns true if the method "HTMLTableRowElement.deleteCell" exists.
func (this HTMLTableRowElement) HasDeleteCell() bool {
	return js.True == bindings.HasHTMLTableRowElementDeleteCell(
		this.Ref(),
	)
}

// DeleteCellFunc returns the method "HTMLTableRowElement.deleteCell".
func (this HTMLTableRowElement) DeleteCellFunc() (fn js.Func[func(index int32)]) {
	return fn.FromRef(
		bindings.HTMLTableRowElementDeleteCellFunc(
			this.Ref(),
		),
	)
}

// DeleteCell calls the method "HTMLTableRowElement.deleteCell".
func (this HTMLTableRowElement) DeleteCell(index int32) (ret js.Void) {
	bindings.CallHTMLTableRowElementDeleteCell(
		this.Ref(), js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryDeleteCell calls the method "HTMLTableRowElement.deleteCell"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableRowElement) TryDeleteCell(index int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableRowElementDeleteCell(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

func NewHTMLTableSectionElement() (ret HTMLTableSectionElement) {
	ret.ref = bindings.NewHTMLTableSectionElementByHTMLTableSectionElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLTableSectionElement) Rows() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLTableSectionElementRows(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Align returns the value of property "HTMLTableSectionElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLTableSectionElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableSectionElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableSectionElement.align" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableSectionElement) Ch() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableSectionElementCh(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCh sets the value of property "HTMLTableSectionElement.ch" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableSectionElement) ChOff() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableSectionElementChOff(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetChOff sets the value of property "HTMLTableSectionElement.chOff" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableSectionElement) VAlign() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableSectionElementVAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetVAlign sets the value of property "HTMLTableSectionElement.vAlign" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableSectionElement) SetVAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableSectionElementVAlign(
		this.Ref(),
		val.Ref(),
	)
}

// HasInsertRow returns true if the method "HTMLTableSectionElement.insertRow" exists.
func (this HTMLTableSectionElement) HasInsertRow() bool {
	return js.True == bindings.HasHTMLTableSectionElementInsertRow(
		this.Ref(),
	)
}

// InsertRowFunc returns the method "HTMLTableSectionElement.insertRow".
func (this HTMLTableSectionElement) InsertRowFunc() (fn js.Func[func(index int32) HTMLTableRowElement]) {
	return fn.FromRef(
		bindings.HTMLTableSectionElementInsertRowFunc(
			this.Ref(),
		),
	)
}

// InsertRow calls the method "HTMLTableSectionElement.insertRow".
func (this HTMLTableSectionElement) InsertRow(index int32) (ret HTMLTableRowElement) {
	bindings.CallHTMLTableSectionElementInsertRow(
		this.Ref(), js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryInsertRow calls the method "HTMLTableSectionElement.insertRow"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableSectionElement) TryInsertRow(index int32) (ret HTMLTableRowElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableSectionElementInsertRow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

// HasInsertRow1 returns true if the method "HTMLTableSectionElement.insertRow" exists.
func (this HTMLTableSectionElement) HasInsertRow1() bool {
	return js.True == bindings.HasHTMLTableSectionElementInsertRow1(
		this.Ref(),
	)
}

// InsertRow1Func returns the method "HTMLTableSectionElement.insertRow".
func (this HTMLTableSectionElement) InsertRow1Func() (fn js.Func[func() HTMLTableRowElement]) {
	return fn.FromRef(
		bindings.HTMLTableSectionElementInsertRow1Func(
			this.Ref(),
		),
	)
}

// InsertRow1 calls the method "HTMLTableSectionElement.insertRow".
func (this HTMLTableSectionElement) InsertRow1() (ret HTMLTableRowElement) {
	bindings.CallHTMLTableSectionElementInsertRow1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryInsertRow1 calls the method "HTMLTableSectionElement.insertRow"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableSectionElement) TryInsertRow1() (ret HTMLTableRowElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableSectionElementInsertRow1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteRow returns true if the method "HTMLTableSectionElement.deleteRow" exists.
func (this HTMLTableSectionElement) HasDeleteRow() bool {
	return js.True == bindings.HasHTMLTableSectionElementDeleteRow(
		this.Ref(),
	)
}

// DeleteRowFunc returns the method "HTMLTableSectionElement.deleteRow".
func (this HTMLTableSectionElement) DeleteRowFunc() (fn js.Func[func(index int32)]) {
	return fn.FromRef(
		bindings.HTMLTableSectionElementDeleteRowFunc(
			this.Ref(),
		),
	)
}

// DeleteRow calls the method "HTMLTableSectionElement.deleteRow".
func (this HTMLTableSectionElement) DeleteRow(index int32) (ret js.Void) {
	bindings.CallHTMLTableSectionElementDeleteRow(
		this.Ref(), js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryDeleteRow calls the method "HTMLTableSectionElement.deleteRow"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableSectionElement) TryDeleteRow(index int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableSectionElementDeleteRow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

func NewHTMLTableElement() (ret HTMLTableElement) {
	ret.ref = bindings.NewHTMLTableElementByHTMLTableElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Caption() (ret HTMLTableCaptionElement, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementCaption(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCaption sets the value of property "HTMLTableElement.caption" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) THead() (ret HTMLTableSectionElement, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementTHead(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTHead sets the value of property "HTMLTableElement.tHead" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) TFoot() (ret HTMLTableSectionElement, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementTFoot(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetTFoot sets the value of property "HTMLTableElement.tFoot" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) TBodies() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementTBodies(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Rows returns the value of property "HTMLTableElement.rows".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Rows() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementRows(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Align returns the value of property "HTMLTableElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementAlign(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableElement.align" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Border() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementBorder(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBorder sets the value of property "HTMLTableElement.border" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Frame() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementFrame(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetFrame sets the value of property "HTMLTableElement.frame" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Rules() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementRules(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRules sets the value of property "HTMLTableElement.rules" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Summary() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementSummary(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSummary sets the value of property "HTMLTableElement.summary" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementWidth(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLTableElement.width" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementBgColor(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "HTMLTableElement.bgColor" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) CellPadding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementCellPadding(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCellPadding sets the value of property "HTMLTableElement.cellPadding" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTableElement) CellSpacing() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementCellSpacing(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCellSpacing sets the value of property "HTMLTableElement.cellSpacing" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetCellSpacing(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementCellSpacing(
		this.Ref(),
		val.Ref(),
	)
}

// HasCreateCaption returns true if the method "HTMLTableElement.createCaption" exists.
func (this HTMLTableElement) HasCreateCaption() bool {
	return js.True == bindings.HasHTMLTableElementCreateCaption(
		this.Ref(),
	)
}

// CreateCaptionFunc returns the method "HTMLTableElement.createCaption".
func (this HTMLTableElement) CreateCaptionFunc() (fn js.Func[func() HTMLTableCaptionElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementCreateCaptionFunc(
			this.Ref(),
		),
	)
}

// CreateCaption calls the method "HTMLTableElement.createCaption".
func (this HTMLTableElement) CreateCaption() (ret HTMLTableCaptionElement) {
	bindings.CallHTMLTableElementCreateCaption(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateCaption calls the method "HTMLTableElement.createCaption"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableElement) TryCreateCaption() (ret HTMLTableCaptionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementCreateCaption(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteCaption returns true if the method "HTMLTableElement.deleteCaption" exists.
func (this HTMLTableElement) HasDeleteCaption() bool {
	return js.True == bindings.HasHTMLTableElementDeleteCaption(
		this.Ref(),
	)
}

// DeleteCaptionFunc returns the method "HTMLTableElement.deleteCaption".
func (this HTMLTableElement) DeleteCaptionFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLTableElementDeleteCaptionFunc(
			this.Ref(),
		),
	)
}

// DeleteCaption calls the method "HTMLTableElement.deleteCaption".
func (this HTMLTableElement) DeleteCaption() (ret js.Void) {
	bindings.CallHTMLTableElementDeleteCaption(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDeleteCaption calls the method "HTMLTableElement.deleteCaption"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableElement) TryDeleteCaption() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementDeleteCaption(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateTHead returns true if the method "HTMLTableElement.createTHead" exists.
func (this HTMLTableElement) HasCreateTHead() bool {
	return js.True == bindings.HasHTMLTableElementCreateTHead(
		this.Ref(),
	)
}

// CreateTHeadFunc returns the method "HTMLTableElement.createTHead".
func (this HTMLTableElement) CreateTHeadFunc() (fn js.Func[func() HTMLTableSectionElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementCreateTHeadFunc(
			this.Ref(),
		),
	)
}

// CreateTHead calls the method "HTMLTableElement.createTHead".
func (this HTMLTableElement) CreateTHead() (ret HTMLTableSectionElement) {
	bindings.CallHTMLTableElementCreateTHead(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateTHead calls the method "HTMLTableElement.createTHead"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableElement) TryCreateTHead() (ret HTMLTableSectionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementCreateTHead(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteTHead returns true if the method "HTMLTableElement.deleteTHead" exists.
func (this HTMLTableElement) HasDeleteTHead() bool {
	return js.True == bindings.HasHTMLTableElementDeleteTHead(
		this.Ref(),
	)
}

// DeleteTHeadFunc returns the method "HTMLTableElement.deleteTHead".
func (this HTMLTableElement) DeleteTHeadFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLTableElementDeleteTHeadFunc(
			this.Ref(),
		),
	)
}

// DeleteTHead calls the method "HTMLTableElement.deleteTHead".
func (this HTMLTableElement) DeleteTHead() (ret js.Void) {
	bindings.CallHTMLTableElementDeleteTHead(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDeleteTHead calls the method "HTMLTableElement.deleteTHead"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableElement) TryDeleteTHead() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementDeleteTHead(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateTFoot returns true if the method "HTMLTableElement.createTFoot" exists.
func (this HTMLTableElement) HasCreateTFoot() bool {
	return js.True == bindings.HasHTMLTableElementCreateTFoot(
		this.Ref(),
	)
}

// CreateTFootFunc returns the method "HTMLTableElement.createTFoot".
func (this HTMLTableElement) CreateTFootFunc() (fn js.Func[func() HTMLTableSectionElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementCreateTFootFunc(
			this.Ref(),
		),
	)
}

// CreateTFoot calls the method "HTMLTableElement.createTFoot".
func (this HTMLTableElement) CreateTFoot() (ret HTMLTableSectionElement) {
	bindings.CallHTMLTableElementCreateTFoot(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateTFoot calls the method "HTMLTableElement.createTFoot"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableElement) TryCreateTFoot() (ret HTMLTableSectionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementCreateTFoot(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteTFoot returns true if the method "HTMLTableElement.deleteTFoot" exists.
func (this HTMLTableElement) HasDeleteTFoot() bool {
	return js.True == bindings.HasHTMLTableElementDeleteTFoot(
		this.Ref(),
	)
}

// DeleteTFootFunc returns the method "HTMLTableElement.deleteTFoot".
func (this HTMLTableElement) DeleteTFootFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLTableElementDeleteTFootFunc(
			this.Ref(),
		),
	)
}

// DeleteTFoot calls the method "HTMLTableElement.deleteTFoot".
func (this HTMLTableElement) DeleteTFoot() (ret js.Void) {
	bindings.CallHTMLTableElementDeleteTFoot(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryDeleteTFoot calls the method "HTMLTableElement.deleteTFoot"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableElement) TryDeleteTFoot() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementDeleteTFoot(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateTBody returns true if the method "HTMLTableElement.createTBody" exists.
func (this HTMLTableElement) HasCreateTBody() bool {
	return js.True == bindings.HasHTMLTableElementCreateTBody(
		this.Ref(),
	)
}

// CreateTBodyFunc returns the method "HTMLTableElement.createTBody".
func (this HTMLTableElement) CreateTBodyFunc() (fn js.Func[func() HTMLTableSectionElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementCreateTBodyFunc(
			this.Ref(),
		),
	)
}

// CreateTBody calls the method "HTMLTableElement.createTBody".
func (this HTMLTableElement) CreateTBody() (ret HTMLTableSectionElement) {
	bindings.CallHTMLTableElementCreateTBody(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCreateTBody calls the method "HTMLTableElement.createTBody"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableElement) TryCreateTBody() (ret HTMLTableSectionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementCreateTBody(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasInsertRow returns true if the method "HTMLTableElement.insertRow" exists.
func (this HTMLTableElement) HasInsertRow() bool {
	return js.True == bindings.HasHTMLTableElementInsertRow(
		this.Ref(),
	)
}

// InsertRowFunc returns the method "HTMLTableElement.insertRow".
func (this HTMLTableElement) InsertRowFunc() (fn js.Func[func(index int32) HTMLTableRowElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementInsertRowFunc(
			this.Ref(),
		),
	)
}

// InsertRow calls the method "HTMLTableElement.insertRow".
func (this HTMLTableElement) InsertRow(index int32) (ret HTMLTableRowElement) {
	bindings.CallHTMLTableElementInsertRow(
		this.Ref(), js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryInsertRow calls the method "HTMLTableElement.insertRow"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableElement) TryInsertRow(index int32) (ret HTMLTableRowElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementInsertRow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

// HasInsertRow1 returns true if the method "HTMLTableElement.insertRow" exists.
func (this HTMLTableElement) HasInsertRow1() bool {
	return js.True == bindings.HasHTMLTableElementInsertRow1(
		this.Ref(),
	)
}

// InsertRow1Func returns the method "HTMLTableElement.insertRow".
func (this HTMLTableElement) InsertRow1Func() (fn js.Func[func() HTMLTableRowElement]) {
	return fn.FromRef(
		bindings.HTMLTableElementInsertRow1Func(
			this.Ref(),
		),
	)
}

// InsertRow1 calls the method "HTMLTableElement.insertRow".
func (this HTMLTableElement) InsertRow1() (ret HTMLTableRowElement) {
	bindings.CallHTMLTableElementInsertRow1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryInsertRow1 calls the method "HTMLTableElement.insertRow"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableElement) TryInsertRow1() (ret HTMLTableRowElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementInsertRow1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasDeleteRow returns true if the method "HTMLTableElement.deleteRow" exists.
func (this HTMLTableElement) HasDeleteRow() bool {
	return js.True == bindings.HasHTMLTableElementDeleteRow(
		this.Ref(),
	)
}

// DeleteRowFunc returns the method "HTMLTableElement.deleteRow".
func (this HTMLTableElement) DeleteRowFunc() (fn js.Func[func(index int32)]) {
	return fn.FromRef(
		bindings.HTMLTableElementDeleteRowFunc(
			this.Ref(),
		),
	)
}

// DeleteRow calls the method "HTMLTableElement.deleteRow".
func (this HTMLTableElement) DeleteRow(index int32) (ret js.Void) {
	bindings.CallHTMLTableElementDeleteRow(
		this.Ref(), js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryDeleteRow calls the method "HTMLTableElement.deleteRow"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTableElement) TryDeleteRow(index int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementDeleteRow(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

func NewHTMLTemplateElement() (ret HTMLTemplateElement) {
	ret.ref = bindings.NewHTMLTemplateElementByHTMLTemplateElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLTemplateElement) Content() (ret DocumentFragment, ok bool) {
	ok = js.True == bindings.GetHTMLTemplateElementContent(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewHTMLTextAreaElement() (ret HTMLTextAreaElement) {
	ret.ref = bindings.NewHTMLTextAreaElementByHTMLTextAreaElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Autocomplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementAutocomplete(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetAutocomplete sets the value of property "HTMLTextAreaElement.autocomplete" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Cols() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementCols(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCols sets the value of property "HTMLTextAreaElement.cols" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) DirName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementDirName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDirName sets the value of property "HTMLTextAreaElement.dirName" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementDisabled(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLTextAreaElement.disabled" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementForm(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MaxLength returns the value of property "HTMLTextAreaElement.maxLength".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) MaxLength() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementMaxLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMaxLength sets the value of property "HTMLTextAreaElement.maxLength" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) MinLength() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementMinLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetMinLength sets the value of property "HTMLTextAreaElement.minLength" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLTextAreaElement.name" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Placeholder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementPlaceholder(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPlaceholder sets the value of property "HTMLTextAreaElement.placeholder" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) ReadOnly() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementReadOnly(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetReadOnly sets the value of property "HTMLTextAreaElement.readOnly" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Required() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementRequired(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRequired sets the value of property "HTMLTextAreaElement.required" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Rows() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementRows(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetRows sets the value of property "HTMLTextAreaElement.rows" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Wrap() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementWrap(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetWrap sets the value of property "HTMLTextAreaElement.wrap" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// DefaultValue returns the value of property "HTMLTextAreaElement.defaultValue".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) DefaultValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementDefaultValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDefaultValue sets the value of property "HTMLTextAreaElement.defaultValue" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementValue(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLTextAreaElement.value" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) TextLength() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementTextLength(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// WillValidate returns the value of property "HTMLTextAreaElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementWillValidate(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLTextAreaElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementValidity(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLTextAreaElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementValidationMessage(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLTextAreaElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementLabels(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SelectionStart returns the value of property "HTMLTextAreaElement.selectionStart".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) SelectionStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementSelectionStart(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelectionStart sets the value of property "HTMLTextAreaElement.selectionStart" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) SelectionEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementSelectionEnd(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelectionEnd sets the value of property "HTMLTextAreaElement.selectionEnd" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) SelectionDirection() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementSelectionDirection(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSelectionDirection sets the value of property "HTMLTextAreaElement.selectionDirection" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetSelectionDirection(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementSelectionDirection(
		this.Ref(),
		val.Ref(),
	)
}

// HasCheckValidity returns true if the method "HTMLTextAreaElement.checkValidity" exists.
func (this HTMLTextAreaElement) HasCheckValidity() bool {
	return js.True == bindings.HasHTMLTextAreaElementCheckValidity(
		this.Ref(),
	)
}

// CheckValidityFunc returns the method "HTMLTextAreaElement.checkValidity".
func (this HTMLTextAreaElement) CheckValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementCheckValidityFunc(
			this.Ref(),
		),
	)
}

// CheckValidity calls the method "HTMLTextAreaElement.checkValidity".
func (this HTMLTextAreaElement) CheckValidity() (ret bool) {
	bindings.CallHTMLTextAreaElementCheckValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLTextAreaElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTextAreaElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementCheckValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasReportValidity returns true if the method "HTMLTextAreaElement.reportValidity" exists.
func (this HTMLTextAreaElement) HasReportValidity() bool {
	return js.True == bindings.HasHTMLTextAreaElementReportValidity(
		this.Ref(),
	)
}

// ReportValidityFunc returns the method "HTMLTextAreaElement.reportValidity".
func (this HTMLTextAreaElement) ReportValidityFunc() (fn js.Func[func() bool]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementReportValidityFunc(
			this.Ref(),
		),
	)
}

// ReportValidity calls the method "HTMLTextAreaElement.reportValidity".
func (this HTMLTextAreaElement) ReportValidity() (ret bool) {
	bindings.CallHTMLTextAreaElementReportValidity(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLTextAreaElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTextAreaElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementReportValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetCustomValidity returns true if the method "HTMLTextAreaElement.setCustomValidity" exists.
func (this HTMLTextAreaElement) HasSetCustomValidity() bool {
	return js.True == bindings.HasHTMLTextAreaElementSetCustomValidity(
		this.Ref(),
	)
}

// SetCustomValidityFunc returns the method "HTMLTextAreaElement.setCustomValidity".
func (this HTMLTextAreaElement) SetCustomValidityFunc() (fn js.Func[func(err js.String)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetCustomValidityFunc(
			this.Ref(),
		),
	)
}

// SetCustomValidity calls the method "HTMLTextAreaElement.setCustomValidity".
func (this HTMLTextAreaElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLTextAreaElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTextAreaElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetCustomValidity(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

// HasSelect returns true if the method "HTMLTextAreaElement.select" exists.
func (this HTMLTextAreaElement) HasSelect() bool {
	return js.True == bindings.HasHTMLTextAreaElementSelect(
		this.Ref(),
	)
}

// SelectFunc returns the method "HTMLTextAreaElement.select".
func (this HTMLTextAreaElement) SelectFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSelectFunc(
			this.Ref(),
		),
	)
}

// Select calls the method "HTMLTextAreaElement.select".
func (this HTMLTextAreaElement) Select() (ret js.Void) {
	bindings.CallHTMLTextAreaElementSelect(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TrySelect calls the method "HTMLTextAreaElement.select"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTextAreaElement) TrySelect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSelect(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasSetRangeText returns true if the method "HTMLTextAreaElement.setRangeText" exists.
func (this HTMLTextAreaElement) HasSetRangeText() bool {
	return js.True == bindings.HasHTMLTextAreaElementSetRangeText(
		this.Ref(),
	)
}

// SetRangeTextFunc returns the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) SetRangeTextFunc() (fn js.Func[func(replacement js.String)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetRangeTextFunc(
			this.Ref(),
		),
	)
}

// SetRangeText calls the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) SetRangeText(replacement js.String) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetRangeText(
		this.Ref(), js.Pointer(&ret),
		replacement.Ref(),
	)

	return
}

// TrySetRangeText calls the method "HTMLTextAreaElement.setRangeText"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTextAreaElement) TrySetRangeText(replacement js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetRangeText(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
	)

	return
}

// HasSetRangeText1 returns true if the method "HTMLTextAreaElement.setRangeText" exists.
func (this HTMLTextAreaElement) HasSetRangeText1() bool {
	return js.True == bindings.HasHTMLTextAreaElementSetRangeText1(
		this.Ref(),
	)
}

// SetRangeText1Func returns the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) SetRangeText1Func() (fn js.Func[func(replacement js.String, start uint32, end uint32, selectionMode SelectionMode)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetRangeText1Func(
			this.Ref(),
		),
	)
}

// SetRangeText1 calls the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) SetRangeText1(replacement js.String, start uint32, end uint32, selectionMode SelectionMode) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetRangeText1(
		this.Ref(), js.Pointer(&ret),
		replacement.Ref(),
		uint32(start),
		uint32(end),
		uint32(selectionMode),
	)

	return
}

// TrySetRangeText1 calls the method "HTMLTextAreaElement.setRangeText"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTextAreaElement) TrySetRangeText1(replacement js.String, start uint32, end uint32, selectionMode SelectionMode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetRangeText1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
		uint32(start),
		uint32(end),
		uint32(selectionMode),
	)

	return
}

// HasSetRangeText2 returns true if the method "HTMLTextAreaElement.setRangeText" exists.
func (this HTMLTextAreaElement) HasSetRangeText2() bool {
	return js.True == bindings.HasHTMLTextAreaElementSetRangeText2(
		this.Ref(),
	)
}

// SetRangeText2Func returns the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) SetRangeText2Func() (fn js.Func[func(replacement js.String, start uint32, end uint32)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetRangeText2Func(
			this.Ref(),
		),
	)
}

// SetRangeText2 calls the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) SetRangeText2(replacement js.String, start uint32, end uint32) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetRangeText2(
		this.Ref(), js.Pointer(&ret),
		replacement.Ref(),
		uint32(start),
		uint32(end),
	)

	return
}

// TrySetRangeText2 calls the method "HTMLTextAreaElement.setRangeText"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTextAreaElement) TrySetRangeText2(replacement js.String, start uint32, end uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetRangeText2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
		uint32(start),
		uint32(end),
	)

	return
}

// HasSetSelectionRange returns true if the method "HTMLTextAreaElement.setSelectionRange" exists.
func (this HTMLTextAreaElement) HasSetSelectionRange() bool {
	return js.True == bindings.HasHTMLTextAreaElementSetSelectionRange(
		this.Ref(),
	)
}

// SetSelectionRangeFunc returns the method "HTMLTextAreaElement.setSelectionRange".
func (this HTMLTextAreaElement) SetSelectionRangeFunc() (fn js.Func[func(start uint32, end uint32, direction js.String)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetSelectionRangeFunc(
			this.Ref(),
		),
	)
}

// SetSelectionRange calls the method "HTMLTextAreaElement.setSelectionRange".
func (this HTMLTextAreaElement) SetSelectionRange(start uint32, end uint32, direction js.String) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetSelectionRange(
		this.Ref(), js.Pointer(&ret),
		uint32(start),
		uint32(end),
		direction.Ref(),
	)

	return
}

// TrySetSelectionRange calls the method "HTMLTextAreaElement.setSelectionRange"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTextAreaElement) TrySetSelectionRange(start uint32, end uint32, direction js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetSelectionRange(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(start),
		uint32(end),
		direction.Ref(),
	)

	return
}

// HasSetSelectionRange1 returns true if the method "HTMLTextAreaElement.setSelectionRange" exists.
func (this HTMLTextAreaElement) HasSetSelectionRange1() bool {
	return js.True == bindings.HasHTMLTextAreaElementSetSelectionRange1(
		this.Ref(),
	)
}

// SetSelectionRange1Func returns the method "HTMLTextAreaElement.setSelectionRange".
func (this HTMLTextAreaElement) SetSelectionRange1Func() (fn js.Func[func(start uint32, end uint32)]) {
	return fn.FromRef(
		bindings.HTMLTextAreaElementSetSelectionRange1Func(
			this.Ref(),
		),
	)
}

// SetSelectionRange1 calls the method "HTMLTextAreaElement.setSelectionRange".
func (this HTMLTextAreaElement) SetSelectionRange1(start uint32, end uint32) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetSelectionRange1(
		this.Ref(), js.Pointer(&ret),
		uint32(start),
		uint32(end),
	)

	return
}

// TrySetSelectionRange1 calls the method "HTMLTextAreaElement.setSelectionRange"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this HTMLTextAreaElement) TrySetSelectionRange1(start uint32, end uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetSelectionRange1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		uint32(start),
		uint32(end),
	)

	return
}

func NewHTMLTimeElement() (ret HTMLTimeElement) {
	ret.ref = bindings.NewHTMLTimeElementByHTMLTimeElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLTimeElement) DateTime() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTimeElementDateTime(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDateTime sets the value of property "HTMLTimeElement.dateTime" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTimeElement) SetDateTime(val js.String) bool {
	return js.True == bindings.SetHTMLTimeElementDateTime(
		this.Ref(),
		val.Ref(),
	)
}

func NewHTMLTitleElement() (ret HTMLTitleElement) {
	ret.ref = bindings.NewHTMLTitleElementByHTMLTitleElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLTitleElement) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTitleElementText(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "HTMLTitleElement.text" to val.
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

func NewHTMLTrackElement() (ret HTMLTrackElement) {
	ret.ref = bindings.NewHTMLTrackElementByHTMLTrackElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Kind() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementKind(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetKind sets the value of property "HTMLTrackElement.kind" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementSrc(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLTrackElement.src" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Srclang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementSrclang(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetSrclang sets the value of property "HTMLTrackElement.srclang" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementLabel(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "HTMLTrackElement.label" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Default() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementDefault(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetDefault sets the value of property "HTMLTrackElement.default" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) ReadyState() (ret uint16, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementReadyState(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Track returns the value of property "HTMLTrackElement.track".
//
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Track() (ret TextTrack, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementTrack(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

func NewHTMLUListElement() (ret HTMLUListElement) {
	ret.ref = bindings.NewHTMLUListElementByHTMLUListElement()
	return
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
// It returns ok=false if there is no such property.
func (this HTMLUListElement) Compact() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLUListElementCompact(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetCompact sets the value of property "HTMLUListElement.compact" to val.
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
// It returns ok=false if there is no such property.
func (this HTMLUListElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLUListElementType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLUListElement.type" to val.
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

func NewHashChangeEvent(typ js.String, eventInitDict HashChangeEventInit) (ret HashChangeEvent) {
	ret.ref = bindings.NewHashChangeEventByHashChangeEvent(
		typ.Ref(),
		js.Pointer(&eventInitDict))
	return
}

func NewHashChangeEventByHashChangeEvent1(typ js.String) (ret HashChangeEvent) {
	ret.ref = bindings.NewHashChangeEventByHashChangeEvent1(
		typ.Ref())
	return
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
// It returns ok=false if there is no such property.
func (this HashChangeEvent) OldURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHashChangeEventOldURL(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// NewURL returns the value of property "HashChangeEvent.newURL".
//
// It returns ok=false if there is no such property.
func (this HashChangeEvent) NewURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHashChangeEventNewURL(
		this.Ref(), js.Pointer(&ret),
	)
	return
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

func NewHighlight(initialRanges ...AbstractRange) (ret Highlight) {
	ret.ref = bindings.NewHighlightByHighlight(
		js.SliceData(initialRanges),
		js.SizeU(len(initialRanges)))
	return
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
// It returns ok=false if there is no such property.
func (this Highlight) Priority() (ret int32, ok bool) {
	ok = js.True == bindings.GetHighlightPriority(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetPriority sets the value of property "Highlight.priority" to val.
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
// It returns ok=false if there is no such property.
func (this Highlight) Type() (ret HighlightType, ok bool) {
	ok = js.True == bindings.GetHighlightType(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "Highlight.type" to val.
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
	//
	// NOTE: Hash.FFI_USE MUST be set to true to get Hash used.
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
// It returns ok=false if there is no such property.
func (this IDBIndex) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetIDBIndexName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "IDBIndex.name" to val.
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
// It returns ok=false if there is no such property.
func (this IDBIndex) ObjectStore() (ret IDBObjectStore, ok bool) {
	ok = js.True == bindings.GetIDBIndexObjectStore(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// KeyPath returns the value of property "IDBIndex.keyPath".
//
// It returns ok=false if there is no such property.
func (this IDBIndex) KeyPath() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBIndexKeyPath(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// MultiEntry returns the value of property "IDBIndex.multiEntry".
//
// It returns ok=false if there is no such property.
func (this IDBIndex) MultiEntry() (ret bool, ok bool) {
	ok = js.True == bindings.GetIDBIndexMultiEntry(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Unique returns the value of property "IDBIndex.unique".
//
// It returns ok=false if there is no such property.
func (this IDBIndex) Unique() (ret bool, ok bool) {
	ok = js.True == bindings.GetIDBIndexUnique(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasGet returns true if the method "IDBIndex.get" exists.
func (this IDBIndex) HasGet() bool {
	return js.True == bindings.HasIDBIndexGet(
		this.Ref(),
	)
}

// GetFunc returns the method "IDBIndex.get".
func (this IDBIndex) GetFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "IDBIndex.get".
func (this IDBIndex) Get(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexGet(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGet calls the method "IDBIndex.get"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryGet(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasGetKey returns true if the method "IDBIndex.getKey" exists.
func (this IDBIndex) HasGetKey() bool {
	return js.True == bindings.HasIDBIndexGetKey(
		this.Ref(),
	)
}

// GetKeyFunc returns the method "IDBIndex.getKey".
func (this IDBIndex) GetKeyFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetKeyFunc(
			this.Ref(),
		),
	)
}

// GetKey calls the method "IDBIndex.getKey".
func (this IDBIndex) GetKey(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexGetKey(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetKey calls the method "IDBIndex.getKey"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryGetKey(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasGetAll returns true if the method "IDBIndex.getAll" exists.
func (this IDBIndex) HasGetAll() bool {
	return js.True == bindings.HasIDBIndexGetAll(
		this.Ref(),
	)
}

// GetAllFunc returns the method "IDBIndex.getAll".
func (this IDBIndex) GetAllFunc() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAllFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "IDBIndex.getAll".
func (this IDBIndex) GetAll(query js.Any, count uint32) (ret IDBRequest) {
	bindings.CallIDBIndexGetAll(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
		uint32(count),
	)

	return
}

// TryGetAll calls the method "IDBIndex.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryGetAll(query js.Any, count uint32) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(count),
	)

	return
}

// HasGetAll1 returns true if the method "IDBIndex.getAll" exists.
func (this IDBIndex) HasGetAll1() bool {
	return js.True == bindings.HasIDBIndexGetAll1(
		this.Ref(),
	)
}

// GetAll1Func returns the method "IDBIndex.getAll".
func (this IDBIndex) GetAll1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAll1Func(
			this.Ref(),
		),
	)
}

// GetAll1 calls the method "IDBIndex.getAll".
func (this IDBIndex) GetAll1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexGetAll1(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetAll1 calls the method "IDBIndex.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryGetAll1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAll1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasGetAll2 returns true if the method "IDBIndex.getAll" exists.
func (this IDBIndex) HasGetAll2() bool {
	return js.True == bindings.HasIDBIndexGetAll2(
		this.Ref(),
	)
}

// GetAll2Func returns the method "IDBIndex.getAll".
func (this IDBIndex) GetAll2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAll2Func(
			this.Ref(),
		),
	)
}

// GetAll2 calls the method "IDBIndex.getAll".
func (this IDBIndex) GetAll2() (ret IDBRequest) {
	bindings.CallIDBIndexGetAll2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAll2 calls the method "IDBIndex.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryGetAll2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAll2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetAllKeys returns true if the method "IDBIndex.getAllKeys" exists.
func (this IDBIndex) HasGetAllKeys() bool {
	return js.True == bindings.HasIDBIndexGetAllKeys(
		this.Ref(),
	)
}

// GetAllKeysFunc returns the method "IDBIndex.getAllKeys".
func (this IDBIndex) GetAllKeysFunc() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAllKeysFunc(
			this.Ref(),
		),
	)
}

// GetAllKeys calls the method "IDBIndex.getAllKeys".
func (this IDBIndex) GetAllKeys(query js.Any, count uint32) (ret IDBRequest) {
	bindings.CallIDBIndexGetAllKeys(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
		uint32(count),
	)

	return
}

// TryGetAllKeys calls the method "IDBIndex.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryGetAllKeys(query js.Any, count uint32) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAllKeys(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(count),
	)

	return
}

// HasGetAllKeys1 returns true if the method "IDBIndex.getAllKeys" exists.
func (this IDBIndex) HasGetAllKeys1() bool {
	return js.True == bindings.HasIDBIndexGetAllKeys1(
		this.Ref(),
	)
}

// GetAllKeys1Func returns the method "IDBIndex.getAllKeys".
func (this IDBIndex) GetAllKeys1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAllKeys1Func(
			this.Ref(),
		),
	)
}

// GetAllKeys1 calls the method "IDBIndex.getAllKeys".
func (this IDBIndex) GetAllKeys1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexGetAllKeys1(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetAllKeys1 calls the method "IDBIndex.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryGetAllKeys1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAllKeys1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasGetAllKeys2 returns true if the method "IDBIndex.getAllKeys" exists.
func (this IDBIndex) HasGetAllKeys2() bool {
	return js.True == bindings.HasIDBIndexGetAllKeys2(
		this.Ref(),
	)
}

// GetAllKeys2Func returns the method "IDBIndex.getAllKeys".
func (this IDBIndex) GetAllKeys2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexGetAllKeys2Func(
			this.Ref(),
		),
	)
}

// GetAllKeys2 calls the method "IDBIndex.getAllKeys".
func (this IDBIndex) GetAllKeys2() (ret IDBRequest) {
	bindings.CallIDBIndexGetAllKeys2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAllKeys2 calls the method "IDBIndex.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryGetAllKeys2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAllKeys2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCount returns true if the method "IDBIndex.count" exists.
func (this IDBIndex) HasCount() bool {
	return js.True == bindings.HasIDBIndexCount(
		this.Ref(),
	)
}

// CountFunc returns the method "IDBIndex.count".
func (this IDBIndex) CountFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexCountFunc(
			this.Ref(),
		),
	)
}

// Count calls the method "IDBIndex.count".
func (this IDBIndex) Count(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexCount(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryCount calls the method "IDBIndex.count"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryCount(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexCount(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasCount1 returns true if the method "IDBIndex.count" exists.
func (this IDBIndex) HasCount1() bool {
	return js.True == bindings.HasIDBIndexCount1(
		this.Ref(),
	)
}

// Count1Func returns the method "IDBIndex.count".
func (this IDBIndex) Count1Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexCount1Func(
			this.Ref(),
		),
	)
}

// Count1 calls the method "IDBIndex.count".
func (this IDBIndex) Count1() (ret IDBRequest) {
	bindings.CallIDBIndexCount1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCount1 calls the method "IDBIndex.count"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryCount1() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexCount1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasOpenCursor returns true if the method "IDBIndex.openCursor" exists.
func (this IDBIndex) HasOpenCursor() bool {
	return js.True == bindings.HasIDBIndexOpenCursor(
		this.Ref(),
	)
}

// OpenCursorFunc returns the method "IDBIndex.openCursor".
func (this IDBIndex) OpenCursorFunc() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenCursorFunc(
			this.Ref(),
		),
	)
}

// OpenCursor calls the method "IDBIndex.openCursor".
func (this IDBIndex) OpenCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest) {
	bindings.CallIDBIndexOpenCursor(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
		uint32(direction),
	)

	return
}

// TryOpenCursor calls the method "IDBIndex.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryOpenCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenCursor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(direction),
	)

	return
}

// HasOpenCursor1 returns true if the method "IDBIndex.openCursor" exists.
func (this IDBIndex) HasOpenCursor1() bool {
	return js.True == bindings.HasIDBIndexOpenCursor1(
		this.Ref(),
	)
}

// OpenCursor1Func returns the method "IDBIndex.openCursor".
func (this IDBIndex) OpenCursor1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenCursor1Func(
			this.Ref(),
		),
	)
}

// OpenCursor1 calls the method "IDBIndex.openCursor".
func (this IDBIndex) OpenCursor1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexOpenCursor1(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryOpenCursor1 calls the method "IDBIndex.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryOpenCursor1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenCursor1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasOpenCursor2 returns true if the method "IDBIndex.openCursor" exists.
func (this IDBIndex) HasOpenCursor2() bool {
	return js.True == bindings.HasIDBIndexOpenCursor2(
		this.Ref(),
	)
}

// OpenCursor2Func returns the method "IDBIndex.openCursor".
func (this IDBIndex) OpenCursor2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenCursor2Func(
			this.Ref(),
		),
	)
}

// OpenCursor2 calls the method "IDBIndex.openCursor".
func (this IDBIndex) OpenCursor2() (ret IDBRequest) {
	bindings.CallIDBIndexOpenCursor2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryOpenCursor2 calls the method "IDBIndex.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryOpenCursor2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenCursor2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasOpenKeyCursor returns true if the method "IDBIndex.openKeyCursor" exists.
func (this IDBIndex) HasOpenKeyCursor() bool {
	return js.True == bindings.HasIDBIndexOpenKeyCursor(
		this.Ref(),
	)
}

// OpenKeyCursorFunc returns the method "IDBIndex.openKeyCursor".
func (this IDBIndex) OpenKeyCursorFunc() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenKeyCursorFunc(
			this.Ref(),
		),
	)
}

// OpenKeyCursor calls the method "IDBIndex.openKeyCursor".
func (this IDBIndex) OpenKeyCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest) {
	bindings.CallIDBIndexOpenKeyCursor(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
		uint32(direction),
	)

	return
}

// TryOpenKeyCursor calls the method "IDBIndex.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryOpenKeyCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenKeyCursor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(direction),
	)

	return
}

// HasOpenKeyCursor1 returns true if the method "IDBIndex.openKeyCursor" exists.
func (this IDBIndex) HasOpenKeyCursor1() bool {
	return js.True == bindings.HasIDBIndexOpenKeyCursor1(
		this.Ref(),
	)
}

// OpenKeyCursor1Func returns the method "IDBIndex.openKeyCursor".
func (this IDBIndex) OpenKeyCursor1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenKeyCursor1Func(
			this.Ref(),
		),
	)
}

// OpenKeyCursor1 calls the method "IDBIndex.openKeyCursor".
func (this IDBIndex) OpenKeyCursor1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexOpenKeyCursor1(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryOpenKeyCursor1 calls the method "IDBIndex.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryOpenKeyCursor1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenKeyCursor1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasOpenKeyCursor2 returns true if the method "IDBIndex.openKeyCursor" exists.
func (this IDBIndex) HasOpenKeyCursor2() bool {
	return js.True == bindings.HasIDBIndexOpenKeyCursor2(
		this.Ref(),
	)
}

// OpenKeyCursor2Func returns the method "IDBIndex.openKeyCursor".
func (this IDBIndex) OpenKeyCursor2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBIndexOpenKeyCursor2Func(
			this.Ref(),
		),
	)
}

// OpenKeyCursor2 calls the method "IDBIndex.openKeyCursor".
func (this IDBIndex) OpenKeyCursor2() (ret IDBRequest) {
	bindings.CallIDBIndexOpenKeyCursor2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryOpenKeyCursor2 calls the method "IDBIndex.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBIndex) TryOpenKeyCursor2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenKeyCursor2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this IDBDatabase) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetIDBDatabaseName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Version returns the value of property "IDBDatabase.version".
//
// It returns ok=false if there is no such property.
func (this IDBDatabase) Version() (ret uint64, ok bool) {
	ok = js.True == bindings.GetIDBDatabaseVersion(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// ObjectStoreNames returns the value of property "IDBDatabase.objectStoreNames".
//
// It returns ok=false if there is no such property.
func (this IDBDatabase) ObjectStoreNames() (ret DOMStringList, ok bool) {
	ok = js.True == bindings.GetIDBDatabaseObjectStoreNames(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasTransaction returns true if the method "IDBDatabase.transaction" exists.
func (this IDBDatabase) HasTransaction() bool {
	return js.True == bindings.HasIDBDatabaseTransaction(
		this.Ref(),
	)
}

// TransactionFunc returns the method "IDBDatabase.transaction".
func (this IDBDatabase) TransactionFunc() (fn js.Func[func(storeNames OneOf_String_ArrayString, mode IDBTransactionMode, options IDBTransactionOptions) IDBTransaction]) {
	return fn.FromRef(
		bindings.IDBDatabaseTransactionFunc(
			this.Ref(),
		),
	)
}

// Transaction calls the method "IDBDatabase.transaction".
func (this IDBDatabase) Transaction(storeNames OneOf_String_ArrayString, mode IDBTransactionMode, options IDBTransactionOptions) (ret IDBTransaction) {
	bindings.CallIDBDatabaseTransaction(
		this.Ref(), js.Pointer(&ret),
		storeNames.Ref(),
		uint32(mode),
		js.Pointer(&options),
	)

	return
}

// TryTransaction calls the method "IDBDatabase.transaction"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBDatabase) TryTransaction(storeNames OneOf_String_ArrayString, mode IDBTransactionMode, options IDBTransactionOptions) (ret IDBTransaction, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseTransaction(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		storeNames.Ref(),
		uint32(mode),
		js.Pointer(&options),
	)

	return
}

// HasTransaction1 returns true if the method "IDBDatabase.transaction" exists.
func (this IDBDatabase) HasTransaction1() bool {
	return js.True == bindings.HasIDBDatabaseTransaction1(
		this.Ref(),
	)
}

// Transaction1Func returns the method "IDBDatabase.transaction".
func (this IDBDatabase) Transaction1Func() (fn js.Func[func(storeNames OneOf_String_ArrayString, mode IDBTransactionMode) IDBTransaction]) {
	return fn.FromRef(
		bindings.IDBDatabaseTransaction1Func(
			this.Ref(),
		),
	)
}

// Transaction1 calls the method "IDBDatabase.transaction".
func (this IDBDatabase) Transaction1(storeNames OneOf_String_ArrayString, mode IDBTransactionMode) (ret IDBTransaction) {
	bindings.CallIDBDatabaseTransaction1(
		this.Ref(), js.Pointer(&ret),
		storeNames.Ref(),
		uint32(mode),
	)

	return
}

// TryTransaction1 calls the method "IDBDatabase.transaction"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBDatabase) TryTransaction1(storeNames OneOf_String_ArrayString, mode IDBTransactionMode) (ret IDBTransaction, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseTransaction1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		storeNames.Ref(),
		uint32(mode),
	)

	return
}

// HasTransaction2 returns true if the method "IDBDatabase.transaction" exists.
func (this IDBDatabase) HasTransaction2() bool {
	return js.True == bindings.HasIDBDatabaseTransaction2(
		this.Ref(),
	)
}

// Transaction2Func returns the method "IDBDatabase.transaction".
func (this IDBDatabase) Transaction2Func() (fn js.Func[func(storeNames OneOf_String_ArrayString) IDBTransaction]) {
	return fn.FromRef(
		bindings.IDBDatabaseTransaction2Func(
			this.Ref(),
		),
	)
}

// Transaction2 calls the method "IDBDatabase.transaction".
func (this IDBDatabase) Transaction2(storeNames OneOf_String_ArrayString) (ret IDBTransaction) {
	bindings.CallIDBDatabaseTransaction2(
		this.Ref(), js.Pointer(&ret),
		storeNames.Ref(),
	)

	return
}

// TryTransaction2 calls the method "IDBDatabase.transaction"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBDatabase) TryTransaction2(storeNames OneOf_String_ArrayString) (ret IDBTransaction, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseTransaction2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		storeNames.Ref(),
	)

	return
}

// HasClose returns true if the method "IDBDatabase.close" exists.
func (this IDBDatabase) HasClose() bool {
	return js.True == bindings.HasIDBDatabaseClose(
		this.Ref(),
	)
}

// CloseFunc returns the method "IDBDatabase.close".
func (this IDBDatabase) CloseFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.IDBDatabaseCloseFunc(
			this.Ref(),
		),
	)
}

// Close calls the method "IDBDatabase.close".
func (this IDBDatabase) Close() (ret js.Void) {
	bindings.CallIDBDatabaseClose(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "IDBDatabase.close"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBDatabase) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseClose(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCreateObjectStore returns true if the method "IDBDatabase.createObjectStore" exists.
func (this IDBDatabase) HasCreateObjectStore() bool {
	return js.True == bindings.HasIDBDatabaseCreateObjectStore(
		this.Ref(),
	)
}

// CreateObjectStoreFunc returns the method "IDBDatabase.createObjectStore".
func (this IDBDatabase) CreateObjectStoreFunc() (fn js.Func[func(name js.String, options IDBObjectStoreParameters) IDBObjectStore]) {
	return fn.FromRef(
		bindings.IDBDatabaseCreateObjectStoreFunc(
			this.Ref(),
		),
	)
}

// CreateObjectStore calls the method "IDBDatabase.createObjectStore".
func (this IDBDatabase) CreateObjectStore(name js.String, options IDBObjectStoreParameters) (ret IDBObjectStore) {
	bindings.CallIDBDatabaseCreateObjectStore(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCreateObjectStore calls the method "IDBDatabase.createObjectStore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBDatabase) TryCreateObjectStore(name js.String, options IDBObjectStoreParameters) (ret IDBObjectStore, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseCreateObjectStore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasCreateObjectStore1 returns true if the method "IDBDatabase.createObjectStore" exists.
func (this IDBDatabase) HasCreateObjectStore1() bool {
	return js.True == bindings.HasIDBDatabaseCreateObjectStore1(
		this.Ref(),
	)
}

// CreateObjectStore1Func returns the method "IDBDatabase.createObjectStore".
func (this IDBDatabase) CreateObjectStore1Func() (fn js.Func[func(name js.String) IDBObjectStore]) {
	return fn.FromRef(
		bindings.IDBDatabaseCreateObjectStore1Func(
			this.Ref(),
		),
	)
}

// CreateObjectStore1 calls the method "IDBDatabase.createObjectStore".
func (this IDBDatabase) CreateObjectStore1(name js.String) (ret IDBObjectStore) {
	bindings.CallIDBDatabaseCreateObjectStore1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryCreateObjectStore1 calls the method "IDBDatabase.createObjectStore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBDatabase) TryCreateObjectStore1(name js.String) (ret IDBObjectStore, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseCreateObjectStore1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasDeleteObjectStore returns true if the method "IDBDatabase.deleteObjectStore" exists.
func (this IDBDatabase) HasDeleteObjectStore() bool {
	return js.True == bindings.HasIDBDatabaseDeleteObjectStore(
		this.Ref(),
	)
}

// DeleteObjectStoreFunc returns the method "IDBDatabase.deleteObjectStore".
func (this IDBDatabase) DeleteObjectStoreFunc() (fn js.Func[func(name js.String)]) {
	return fn.FromRef(
		bindings.IDBDatabaseDeleteObjectStoreFunc(
			this.Ref(),
		),
	)
}

// DeleteObjectStore calls the method "IDBDatabase.deleteObjectStore".
func (this IDBDatabase) DeleteObjectStore(name js.String) (ret js.Void) {
	bindings.CallIDBDatabaseDeleteObjectStore(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDeleteObjectStore calls the method "IDBDatabase.deleteObjectStore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBDatabase) TryDeleteObjectStore(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseDeleteObjectStore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
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
// It returns ok=false if there is no such property.
func (this IDBTransaction) ObjectStoreNames() (ret DOMStringList, ok bool) {
	ok = js.True == bindings.GetIDBTransactionObjectStoreNames(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Mode returns the value of property "IDBTransaction.mode".
//
// It returns ok=false if there is no such property.
func (this IDBTransaction) Mode() (ret IDBTransactionMode, ok bool) {
	ok = js.True == bindings.GetIDBTransactionMode(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Durability returns the value of property "IDBTransaction.durability".
//
// It returns ok=false if there is no such property.
func (this IDBTransaction) Durability() (ret IDBTransactionDurability, ok bool) {
	ok = js.True == bindings.GetIDBTransactionDurability(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Db returns the value of property "IDBTransaction.db".
//
// It returns ok=false if there is no such property.
func (this IDBTransaction) Db() (ret IDBDatabase, ok bool) {
	ok = js.True == bindings.GetIDBTransactionDb(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Error returns the value of property "IDBTransaction.error".
//
// It returns ok=false if there is no such property.
func (this IDBTransaction) Error() (ret DOMException, ok bool) {
	ok = js.True == bindings.GetIDBTransactionError(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasObjectStore returns true if the method "IDBTransaction.objectStore" exists.
func (this IDBTransaction) HasObjectStore() bool {
	return js.True == bindings.HasIDBTransactionObjectStore(
		this.Ref(),
	)
}

// ObjectStoreFunc returns the method "IDBTransaction.objectStore".
func (this IDBTransaction) ObjectStoreFunc() (fn js.Func[func(name js.String) IDBObjectStore]) {
	return fn.FromRef(
		bindings.IDBTransactionObjectStoreFunc(
			this.Ref(),
		),
	)
}

// ObjectStore calls the method "IDBTransaction.objectStore".
func (this IDBTransaction) ObjectStore(name js.String) (ret IDBObjectStore) {
	bindings.CallIDBTransactionObjectStore(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryObjectStore calls the method "IDBTransaction.objectStore"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBTransaction) TryObjectStore(name js.String) (ret IDBObjectStore, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBTransactionObjectStore(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasCommit returns true if the method "IDBTransaction.commit" exists.
func (this IDBTransaction) HasCommit() bool {
	return js.True == bindings.HasIDBTransactionCommit(
		this.Ref(),
	)
}

// CommitFunc returns the method "IDBTransaction.commit".
func (this IDBTransaction) CommitFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.IDBTransactionCommitFunc(
			this.Ref(),
		),
	)
}

// Commit calls the method "IDBTransaction.commit".
func (this IDBTransaction) Commit() (ret js.Void) {
	bindings.CallIDBTransactionCommit(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCommit calls the method "IDBTransaction.commit"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBTransaction) TryCommit() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBTransactionCommit(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasAbort returns true if the method "IDBTransaction.abort" exists.
func (this IDBTransaction) HasAbort() bool {
	return js.True == bindings.HasIDBTransactionAbort(
		this.Ref(),
	)
}

// AbortFunc returns the method "IDBTransaction.abort".
func (this IDBTransaction) AbortFunc() (fn js.Func[func()]) {
	return fn.FromRef(
		bindings.IDBTransactionAbortFunc(
			this.Ref(),
		),
	)
}

// Abort calls the method "IDBTransaction.abort".
func (this IDBTransaction) Abort() (ret js.Void) {
	bindings.CallIDBTransactionAbort(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryAbort calls the method "IDBTransaction.abort"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBTransaction) TryAbort() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBTransactionAbort(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
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
// It returns ok=false if there is no such property.
func (this IDBObjectStore) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetIDBObjectStoreName(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "IDBObjectStore.name" to val.
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
// It returns ok=false if there is no such property.
func (this IDBObjectStore) KeyPath() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBObjectStoreKeyPath(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// IndexNames returns the value of property "IDBObjectStore.indexNames".
//
// It returns ok=false if there is no such property.
func (this IDBObjectStore) IndexNames() (ret DOMStringList, ok bool) {
	ok = js.True == bindings.GetIDBObjectStoreIndexNames(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// Transaction returns the value of property "IDBObjectStore.transaction".
//
// It returns ok=false if there is no such property.
func (this IDBObjectStore) Transaction() (ret IDBTransaction, ok bool) {
	ok = js.True == bindings.GetIDBObjectStoreTransaction(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// AutoIncrement returns the value of property "IDBObjectStore.autoIncrement".
//
// It returns ok=false if there is no such property.
func (this IDBObjectStore) AutoIncrement() (ret bool, ok bool) {
	ok = js.True == bindings.GetIDBObjectStoreAutoIncrement(
		this.Ref(), js.Pointer(&ret),
	)
	return
}

// HasPut returns true if the method "IDBObjectStore.put" exists.
func (this IDBObjectStore) HasPut() bool {
	return js.True == bindings.HasIDBObjectStorePut(
		this.Ref(),
	)
}

// PutFunc returns the method "IDBObjectStore.put".
func (this IDBObjectStore) PutFunc() (fn js.Func[func(value js.Any, key js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStorePutFunc(
			this.Ref(),
		),
	)
}

// Put calls the method "IDBObjectStore.put".
func (this IDBObjectStore) Put(value js.Any, key js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStorePut(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
		key.Ref(),
	)

	return
}

// TryPut calls the method "IDBObjectStore.put"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryPut(value js.Any, key js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStorePut(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
		key.Ref(),
	)

	return
}

// HasPut1 returns true if the method "IDBObjectStore.put" exists.
func (this IDBObjectStore) HasPut1() bool {
	return js.True == bindings.HasIDBObjectStorePut1(
		this.Ref(),
	)
}

// Put1Func returns the method "IDBObjectStore.put".
func (this IDBObjectStore) Put1Func() (fn js.Func[func(value js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStorePut1Func(
			this.Ref(),
		),
	)
}

// Put1 calls the method "IDBObjectStore.put".
func (this IDBObjectStore) Put1(value js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStorePut1(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryPut1 calls the method "IDBObjectStore.put"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryPut1(value js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStorePut1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasAdd returns true if the method "IDBObjectStore.add" exists.
func (this IDBObjectStore) HasAdd() bool {
	return js.True == bindings.HasIDBObjectStoreAdd(
		this.Ref(),
	)
}

// AddFunc returns the method "IDBObjectStore.add".
func (this IDBObjectStore) AddFunc() (fn js.Func[func(value js.Any, key js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreAddFunc(
			this.Ref(),
		),
	)
}

// Add calls the method "IDBObjectStore.add".
func (this IDBObjectStore) Add(value js.Any, key js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreAdd(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
		key.Ref(),
	)

	return
}

// TryAdd calls the method "IDBObjectStore.add"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryAdd(value js.Any, key js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreAdd(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
		key.Ref(),
	)

	return
}

// HasAdd1 returns true if the method "IDBObjectStore.add" exists.
func (this IDBObjectStore) HasAdd1() bool {
	return js.True == bindings.HasIDBObjectStoreAdd1(
		this.Ref(),
	)
}

// Add1Func returns the method "IDBObjectStore.add".
func (this IDBObjectStore) Add1Func() (fn js.Func[func(value js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreAdd1Func(
			this.Ref(),
		),
	)
}

// Add1 calls the method "IDBObjectStore.add".
func (this IDBObjectStore) Add1(value js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreAdd1(
		this.Ref(), js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryAdd1 calls the method "IDBObjectStore.add"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryAdd1(value js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreAdd1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasDelete returns true if the method "IDBObjectStore.delete" exists.
func (this IDBObjectStore) HasDelete() bool {
	return js.True == bindings.HasIDBObjectStoreDelete(
		this.Ref(),
	)
}

// DeleteFunc returns the method "IDBObjectStore.delete".
func (this IDBObjectStore) DeleteFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreDeleteFunc(
			this.Ref(),
		),
	)
}

// Delete calls the method "IDBObjectStore.delete".
func (this IDBObjectStore) Delete(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreDelete(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryDelete calls the method "IDBObjectStore.delete"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryDelete(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreDelete(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasClear returns true if the method "IDBObjectStore.clear" exists.
func (this IDBObjectStore) HasClear() bool {
	return js.True == bindings.HasIDBObjectStoreClear(
		this.Ref(),
	)
}

// ClearFunc returns the method "IDBObjectStore.clear".
func (this IDBObjectStore) ClearFunc() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreClearFunc(
			this.Ref(),
		),
	)
}

// Clear calls the method "IDBObjectStore.clear".
func (this IDBObjectStore) Clear() (ret IDBRequest) {
	bindings.CallIDBObjectStoreClear(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "IDBObjectStore.clear"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryClear() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreClear(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGet returns true if the method "IDBObjectStore.get" exists.
func (this IDBObjectStore) HasGet() bool {
	return js.True == bindings.HasIDBObjectStoreGet(
		this.Ref(),
	)
}

// GetFunc returns the method "IDBObjectStore.get".
func (this IDBObjectStore) GetFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetFunc(
			this.Ref(),
		),
	)
}

// Get calls the method "IDBObjectStore.get".
func (this IDBObjectStore) Get(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGet(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGet calls the method "IDBObjectStore.get"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryGet(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGet(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasGetKey returns true if the method "IDBObjectStore.getKey" exists.
func (this IDBObjectStore) HasGetKey() bool {
	return js.True == bindings.HasIDBObjectStoreGetKey(
		this.Ref(),
	)
}

// GetKeyFunc returns the method "IDBObjectStore.getKey".
func (this IDBObjectStore) GetKeyFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetKeyFunc(
			this.Ref(),
		),
	)
}

// GetKey calls the method "IDBObjectStore.getKey".
func (this IDBObjectStore) GetKey(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetKey(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetKey calls the method "IDBObjectStore.getKey"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryGetKey(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetKey(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasGetAll returns true if the method "IDBObjectStore.getAll" exists.
func (this IDBObjectStore) HasGetAll() bool {
	return js.True == bindings.HasIDBObjectStoreGetAll(
		this.Ref(),
	)
}

// GetAllFunc returns the method "IDBObjectStore.getAll".
func (this IDBObjectStore) GetAllFunc() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAllFunc(
			this.Ref(),
		),
	)
}

// GetAll calls the method "IDBObjectStore.getAll".
func (this IDBObjectStore) GetAll(query js.Any, count uint32) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAll(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
		uint32(count),
	)

	return
}

// TryGetAll calls the method "IDBObjectStore.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryGetAll(query js.Any, count uint32) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAll(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(count),
	)

	return
}

// HasGetAll1 returns true if the method "IDBObjectStore.getAll" exists.
func (this IDBObjectStore) HasGetAll1() bool {
	return js.True == bindings.HasIDBObjectStoreGetAll1(
		this.Ref(),
	)
}

// GetAll1Func returns the method "IDBObjectStore.getAll".
func (this IDBObjectStore) GetAll1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAll1Func(
			this.Ref(),
		),
	)
}

// GetAll1 calls the method "IDBObjectStore.getAll".
func (this IDBObjectStore) GetAll1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAll1(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetAll1 calls the method "IDBObjectStore.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryGetAll1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAll1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasGetAll2 returns true if the method "IDBObjectStore.getAll" exists.
func (this IDBObjectStore) HasGetAll2() bool {
	return js.True == bindings.HasIDBObjectStoreGetAll2(
		this.Ref(),
	)
}

// GetAll2Func returns the method "IDBObjectStore.getAll".
func (this IDBObjectStore) GetAll2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAll2Func(
			this.Ref(),
		),
	)
}

// GetAll2 calls the method "IDBObjectStore.getAll".
func (this IDBObjectStore) GetAll2() (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAll2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAll2 calls the method "IDBObjectStore.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryGetAll2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAll2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasGetAllKeys returns true if the method "IDBObjectStore.getAllKeys" exists.
func (this IDBObjectStore) HasGetAllKeys() bool {
	return js.True == bindings.HasIDBObjectStoreGetAllKeys(
		this.Ref(),
	)
}

// GetAllKeysFunc returns the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) GetAllKeysFunc() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAllKeysFunc(
			this.Ref(),
		),
	)
}

// GetAllKeys calls the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) GetAllKeys(query js.Any, count uint32) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAllKeys(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
		uint32(count),
	)

	return
}

// TryGetAllKeys calls the method "IDBObjectStore.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryGetAllKeys(query js.Any, count uint32) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAllKeys(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(count),
	)

	return
}

// HasGetAllKeys1 returns true if the method "IDBObjectStore.getAllKeys" exists.
func (this IDBObjectStore) HasGetAllKeys1() bool {
	return js.True == bindings.HasIDBObjectStoreGetAllKeys1(
		this.Ref(),
	)
}

// GetAllKeys1Func returns the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) GetAllKeys1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAllKeys1Func(
			this.Ref(),
		),
	)
}

// GetAllKeys1 calls the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) GetAllKeys1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAllKeys1(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetAllKeys1 calls the method "IDBObjectStore.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryGetAllKeys1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAllKeys1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasGetAllKeys2 returns true if the method "IDBObjectStore.getAllKeys" exists.
func (this IDBObjectStore) HasGetAllKeys2() bool {
	return js.True == bindings.HasIDBObjectStoreGetAllKeys2(
		this.Ref(),
	)
}

// GetAllKeys2Func returns the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) GetAllKeys2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreGetAllKeys2Func(
			this.Ref(),
		),
	)
}

// GetAllKeys2 calls the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) GetAllKeys2() (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAllKeys2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryGetAllKeys2 calls the method "IDBObjectStore.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryGetAllKeys2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAllKeys2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasCount returns true if the method "IDBObjectStore.count" exists.
func (this IDBObjectStore) HasCount() bool {
	return js.True == bindings.HasIDBObjectStoreCount(
		this.Ref(),
	)
}

// CountFunc returns the method "IDBObjectStore.count".
func (this IDBObjectStore) CountFunc() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreCountFunc(
			this.Ref(),
		),
	)
}

// Count calls the method "IDBObjectStore.count".
func (this IDBObjectStore) Count(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreCount(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryCount calls the method "IDBObjectStore.count"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryCount(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreCount(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasCount1 returns true if the method "IDBObjectStore.count" exists.
func (this IDBObjectStore) HasCount1() bool {
	return js.True == bindings.HasIDBObjectStoreCount1(
		this.Ref(),
	)
}

// Count1Func returns the method "IDBObjectStore.count".
func (this IDBObjectStore) Count1Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreCount1Func(
			this.Ref(),
		),
	)
}

// Count1 calls the method "IDBObjectStore.count".
func (this IDBObjectStore) Count1() (ret IDBRequest) {
	bindings.CallIDBObjectStoreCount1(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryCount1 calls the method "IDBObjectStore.count"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryCount1() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreCount1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasOpenCursor returns true if the method "IDBObjectStore.openCursor" exists.
func (this IDBObjectStore) HasOpenCursor() bool {
	return js.True == bindings.HasIDBObjectStoreOpenCursor(
		this.Ref(),
	)
}

// OpenCursorFunc returns the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) OpenCursorFunc() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenCursorFunc(
			this.Ref(),
		),
	)
}

// OpenCursor calls the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) OpenCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenCursor(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
		uint32(direction),
	)

	return
}

// TryOpenCursor calls the method "IDBObjectStore.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryOpenCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenCursor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(direction),
	)

	return
}

// HasOpenCursor1 returns true if the method "IDBObjectStore.openCursor" exists.
func (this IDBObjectStore) HasOpenCursor1() bool {
	return js.True == bindings.HasIDBObjectStoreOpenCursor1(
		this.Ref(),
	)
}

// OpenCursor1Func returns the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) OpenCursor1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenCursor1Func(
			this.Ref(),
		),
	)
}

// OpenCursor1 calls the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) OpenCursor1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenCursor1(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryOpenCursor1 calls the method "IDBObjectStore.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryOpenCursor1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenCursor1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasOpenCursor2 returns true if the method "IDBObjectStore.openCursor" exists.
func (this IDBObjectStore) HasOpenCursor2() bool {
	return js.True == bindings.HasIDBObjectStoreOpenCursor2(
		this.Ref(),
	)
}

// OpenCursor2Func returns the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) OpenCursor2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenCursor2Func(
			this.Ref(),
		),
	)
}

// OpenCursor2 calls the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) OpenCursor2() (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenCursor2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryOpenCursor2 calls the method "IDBObjectStore.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryOpenCursor2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenCursor2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasOpenKeyCursor returns true if the method "IDBObjectStore.openKeyCursor" exists.
func (this IDBObjectStore) HasOpenKeyCursor() bool {
	return js.True == bindings.HasIDBObjectStoreOpenKeyCursor(
		this.Ref(),
	)
}

// OpenKeyCursorFunc returns the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) OpenKeyCursorFunc() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenKeyCursorFunc(
			this.Ref(),
		),
	)
}

// OpenKeyCursor calls the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) OpenKeyCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenKeyCursor(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
		uint32(direction),
	)

	return
}

// TryOpenKeyCursor calls the method "IDBObjectStore.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryOpenKeyCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenKeyCursor(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(direction),
	)

	return
}

// HasOpenKeyCursor1 returns true if the method "IDBObjectStore.openKeyCursor" exists.
func (this IDBObjectStore) HasOpenKeyCursor1() bool {
	return js.True == bindings.HasIDBObjectStoreOpenKeyCursor1(
		this.Ref(),
	)
}

// OpenKeyCursor1Func returns the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) OpenKeyCursor1Func() (fn js.Func[func(query js.Any) IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenKeyCursor1Func(
			this.Ref(),
		),
	)
}

// OpenKeyCursor1 calls the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) OpenKeyCursor1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenKeyCursor1(
		this.Ref(), js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryOpenKeyCursor1 calls the method "IDBObjectStore.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryOpenKeyCursor1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenKeyCursor1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasOpenKeyCursor2 returns true if the method "IDBObjectStore.openKeyCursor" exists.
func (this IDBObjectStore) HasOpenKeyCursor2() bool {
	return js.True == bindings.HasIDBObjectStoreOpenKeyCursor2(
		this.Ref(),
	)
}

// OpenKeyCursor2Func returns the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) OpenKeyCursor2Func() (fn js.Func[func() IDBRequest]) {
	return fn.FromRef(
		bindings.IDBObjectStoreOpenKeyCursor2Func(
			this.Ref(),
		),
	)
}

// OpenKeyCursor2 calls the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) OpenKeyCursor2() (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenKeyCursor2(
		this.Ref(), js.Pointer(&ret),
	)

	return
}

// TryOpenKeyCursor2 calls the method "IDBObjectStore.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryOpenKeyCursor2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenKeyCursor2(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasIndex returns true if the method "IDBObjectStore.index" exists.
func (this IDBObjectStore) HasIndex() bool {
	return js.True == bindings.HasIDBObjectStoreIndex(
		this.Ref(),
	)
}

// IndexFunc returns the method "IDBObjectStore.index".
func (this IDBObjectStore) IndexFunc() (fn js.Func[func(name js.String) IDBIndex]) {
	return fn.FromRef(
		bindings.IDBObjectStoreIndexFunc(
			this.Ref(),
		),
	)
}

// Index calls the method "IDBObjectStore.index".
func (this IDBObjectStore) Index(name js.String) (ret IDBIndex) {
	bindings.CallIDBObjectStoreIndex(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryIndex calls the method "IDBObjectStore.index"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryIndex(name js.String) (ret IDBIndex, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreIndex(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasCreateIndex returns true if the method "IDBObjectStore.createIndex" exists.
func (this IDBObjectStore) HasCreateIndex() bool {
	return js.True == bindings.HasIDBObjectStoreCreateIndex(
		this.Ref(),
	)
}

// CreateIndexFunc returns the method "IDBObjectStore.createIndex".
func (this IDBObjectStore) CreateIndexFunc() (fn js.Func[func(name js.String, keyPath OneOf_String_ArrayString, options IDBIndexParameters) IDBIndex]) {
	return fn.FromRef(
		bindings.IDBObjectStoreCreateIndexFunc(
			this.Ref(),
		),
	)
}

// CreateIndex calls the method "IDBObjectStore.createIndex".
func (this IDBObjectStore) CreateIndex(name js.String, keyPath OneOf_String_ArrayString, options IDBIndexParameters) (ret IDBIndex) {
	bindings.CallIDBObjectStoreCreateIndex(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		keyPath.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCreateIndex calls the method "IDBObjectStore.createIndex"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryCreateIndex(name js.String, keyPath OneOf_String_ArrayString, options IDBIndexParameters) (ret IDBIndex, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreCreateIndex(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		keyPath.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasCreateIndex1 returns true if the method "IDBObjectStore.createIndex" exists.
func (this IDBObjectStore) HasCreateIndex1() bool {
	return js.True == bindings.HasIDBObjectStoreCreateIndex1(
		this.Ref(),
	)
}

// CreateIndex1Func returns the method "IDBObjectStore.createIndex".
func (this IDBObjectStore) CreateIndex1Func() (fn js.Func[func(name js.String, keyPath OneOf_String_ArrayString) IDBIndex]) {
	return fn.FromRef(
		bindings.IDBObjectStoreCreateIndex1Func(
			this.Ref(),
		),
	)
}

// CreateIndex1 calls the method "IDBObjectStore.createIndex".
func (this IDBObjectStore) CreateIndex1(name js.String, keyPath OneOf_String_ArrayString) (ret IDBIndex) {
	bindings.CallIDBObjectStoreCreateIndex1(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
		keyPath.Ref(),
	)

	return
}

// TryCreateIndex1 calls the method "IDBObjectStore.createIndex"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryCreateIndex1(name js.String, keyPath OneOf_String_ArrayString) (ret IDBIndex, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreCreateIndex1(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		keyPath.Ref(),
	)

	return
}

// HasDeleteIndex returns true if the method "IDBObjectStore.deleteIndex" exists.
func (this IDBObjectStore) HasDeleteIndex() bool {
	return js.True == bindings.HasIDBObjectStoreDeleteIndex(
		this.Ref(),
	)
}

// DeleteIndexFunc returns the method "IDBObjectStore.deleteIndex".
func (this IDBObjectStore) DeleteIndexFunc() (fn js.Func[func(name js.String)]) {
	return fn.FromRef(
		bindings.IDBObjectStoreDeleteIndexFunc(
			this.Ref(),
		),
	)
}

// DeleteIndex calls the method "IDBObjectStore.deleteIndex".
func (this IDBObjectStore) DeleteIndex(name js.String) (ret js.Void) {
	bindings.CallIDBObjectStoreDeleteIndex(
		this.Ref(), js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDeleteIndex calls the method "IDBObjectStore.deleteIndex"
// in a try/catch block and returns (_, err, ok = false) when it went though
// the catch clause.
func (this IDBObjectStore) TryDeleteIndex(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreDeleteIndex(
		this.Ref(), js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
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

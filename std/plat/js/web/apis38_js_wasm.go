// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 The Prime Citizens

package web

import (
	"github.com/primecitizens/pcz/std/ffi/js"
	"github.com/primecitizens/pcz/std/plat/js/web/bindings"
)

type HTMLPortalElement struct {
	HTMLElement
}

func (this HTMLPortalElement) Once() HTMLPortalElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Src returns the value of property "HTMLPortalElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLPortalElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLPortalElementSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLPortalElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLPortalElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLPortalElementSrc(
		this.ref,
		val.Ref(),
	)
}

// ReferrerPolicy returns the value of property "HTMLPortalElement.referrerPolicy".
//
// It returns ok=false if there is no such property.
func (this HTMLPortalElement) ReferrerPolicy() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLPortalElementReferrerPolicy(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReferrerPolicy sets the value of property "HTMLPortalElement.referrerPolicy" to val.
//
// It returns false if the property cannot be set.
func (this HTMLPortalElement) SetReferrerPolicy(val js.String) bool {
	return js.True == bindings.SetHTMLPortalElementReferrerPolicy(
		this.ref,
		val.Ref(),
	)
}

// HasFuncActivate returns true if the method "HTMLPortalElement.activate" exists.
func (this HTMLPortalElement) HasFuncActivate() bool {
	return js.True == bindings.HasFuncHTMLPortalElementActivate(
		this.ref,
	)
}

// FuncActivate returns the method "HTMLPortalElement.activate".
func (this HTMLPortalElement) FuncActivate() (fn js.Func[func(options PortalActivateOptions) js.Promise[js.Void]]) {
	bindings.FuncHTMLPortalElementActivate(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Activate calls the method "HTMLPortalElement.activate".
func (this HTMLPortalElement) Activate(options PortalActivateOptions) (ret js.Promise[js.Void]) {
	bindings.CallHTMLPortalElementActivate(
		this.ref, js.Pointer(&ret),
		js.Pointer(&options),
	)

	return
}

// TryActivate calls the method "HTMLPortalElement.activate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLPortalElement) TryActivate(options PortalActivateOptions) (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLPortalElementActivate(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		js.Pointer(&options),
	)

	return
}

// HasFuncActivate1 returns true if the method "HTMLPortalElement.activate" exists.
func (this HTMLPortalElement) HasFuncActivate1() bool {
	return js.True == bindings.HasFuncHTMLPortalElementActivate1(
		this.ref,
	)
}

// FuncActivate1 returns the method "HTMLPortalElement.activate".
func (this HTMLPortalElement) FuncActivate1() (fn js.Func[func() js.Promise[js.Void]]) {
	bindings.FuncHTMLPortalElementActivate1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Activate1 calls the method "HTMLPortalElement.activate".
func (this HTMLPortalElement) Activate1() (ret js.Promise[js.Void]) {
	bindings.CallHTMLPortalElementActivate1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryActivate1 calls the method "HTMLPortalElement.activate"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLPortalElement) TryActivate1() (ret js.Promise[js.Void], exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLPortalElementActivate1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncPostMessage returns true if the method "HTMLPortalElement.postMessage" exists.
func (this HTMLPortalElement) HasFuncPostMessage() bool {
	return js.True == bindings.HasFuncHTMLPortalElementPostMessage(
		this.ref,
	)
}

// FuncPostMessage returns the method "HTMLPortalElement.postMessage".
func (this HTMLPortalElement) FuncPostMessage() (fn js.Func[func(message js.Any, options StructuredSerializeOptions)]) {
	bindings.FuncHTMLPortalElementPostMessage(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage calls the method "HTMLPortalElement.postMessage".
func (this HTMLPortalElement) PostMessage(message js.Any, options StructuredSerializeOptions) (ret js.Void) {
	bindings.CallHTMLPortalElementPostMessage(
		this.ref, js.Pointer(&ret),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryPostMessage calls the method "HTMLPortalElement.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLPortalElement) TryPostMessage(message js.Any, options StructuredSerializeOptions) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLPortalElementPostMessage(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncPostMessage1 returns true if the method "HTMLPortalElement.postMessage" exists.
func (this HTMLPortalElement) HasFuncPostMessage1() bool {
	return js.True == bindings.HasFuncHTMLPortalElementPostMessage1(
		this.ref,
	)
}

// FuncPostMessage1 returns the method "HTMLPortalElement.postMessage".
func (this HTMLPortalElement) FuncPostMessage1() (fn js.Func[func(message js.Any)]) {
	bindings.FuncHTMLPortalElementPostMessage1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// PostMessage1 calls the method "HTMLPortalElement.postMessage".
func (this HTMLPortalElement) PostMessage1(message js.Any) (ret js.Void) {
	bindings.CallHTMLPortalElementPostMessage1(
		this.ref, js.Pointer(&ret),
		message.Ref(),
	)

	return
}

// TryPostMessage1 calls the method "HTMLPortalElement.postMessage"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLPortalElement) TryPostMessage1(message js.Any) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLPortalElementPostMessage1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		message.Ref(),
	)

	return
}

type HTMLPreElement struct {
	HTMLElement
}

func (this HTMLPreElement) Once() HTMLPreElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Width returns the value of property "HTMLPreElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLPreElement) Width() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLPreElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLPreElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLPreElement) SetWidth(val int32) bool {
	return js.True == bindings.SetHTMLPreElementWidth(
		this.ref,
		int32(val),
	)
}

type HTMLProgressElement struct {
	HTMLElement
}

func (this HTMLProgressElement) Once() HTMLProgressElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Value returns the value of property "HTMLProgressElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLProgressElement) Value() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLProgressElementValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLProgressElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLProgressElement) SetValue(val float64) bool {
	return js.True == bindings.SetHTMLProgressElementValue(
		this.ref,
		float64(val),
	)
}

// Max returns the value of property "HTMLProgressElement.max".
//
// It returns ok=false if there is no such property.
func (this HTMLProgressElement) Max() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLProgressElementMax(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMax sets the value of property "HTMLProgressElement.max" to val.
//
// It returns false if the property cannot be set.
func (this HTMLProgressElement) SetMax(val float64) bool {
	return js.True == bindings.SetHTMLProgressElementMax(
		this.ref,
		float64(val),
	)
}

// Position returns the value of property "HTMLProgressElement.position".
//
// It returns ok=false if there is no such property.
func (this HTMLProgressElement) Position() (ret float64, ok bool) {
	ok = js.True == bindings.GetHTMLProgressElementPosition(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLProgressElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLProgressElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLProgressElementLabels(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HTMLQuoteElement struct {
	HTMLElement
}

func (this HTMLQuoteElement) Once() HTMLQuoteElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Cite returns the value of property "HTMLQuoteElement.cite".
//
// It returns ok=false if there is no such property.
func (this HTMLQuoteElement) Cite() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLQuoteElementCite(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCite sets the value of property "HTMLQuoteElement.cite" to val.
//
// It returns false if the property cannot be set.
func (this HTMLQuoteElement) SetCite(val js.String) bool {
	return js.True == bindings.SetHTMLQuoteElementCite(
		this.ref,
		val.Ref(),
	)
}

type HTMLSelectElement struct {
	HTMLElement
}

func (this HTMLSelectElement) Once() HTMLSelectElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Autocomplete returns the value of property "HTMLSelectElement.autocomplete".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Autocomplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementAutocomplete(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAutocomplete sets the value of property "HTMLSelectElement.autocomplete" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetAutocomplete(val js.String) bool {
	return js.True == bindings.SetHTMLSelectElementAutocomplete(
		this.ref,
		val.Ref(),
	)
}

// Disabled returns the value of property "HTMLSelectElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLSelectElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLSelectElementDisabled(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLSelectElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementForm(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Multiple returns the value of property "HTMLSelectElement.multiple".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Multiple() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementMultiple(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMultiple sets the value of property "HTMLSelectElement.multiple" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetMultiple(val bool) bool {
	return js.True == bindings.SetHTMLSelectElementMultiple(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Name returns the value of property "HTMLSelectElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLSelectElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLSelectElementName(
		this.ref,
		val.Ref(),
	)
}

// Required returns the value of property "HTMLSelectElement.required".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Required() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementRequired(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRequired sets the value of property "HTMLSelectElement.required" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetRequired(val bool) bool {
	return js.True == bindings.SetHTMLSelectElementRequired(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Size returns the value of property "HTMLSelectElement.size".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Size() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementSize(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSize sets the value of property "HTMLSelectElement.size" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetSize(val uint32) bool {
	return js.True == bindings.SetHTMLSelectElementSize(
		this.ref,
		uint32(val),
	)
}

// Type returns the value of property "HTMLSelectElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Options returns the value of property "HTMLSelectElement.options".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Options() (ret HTMLOptionsCollection, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementOptions(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Length returns the value of property "HTMLSelectElement.length".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Length() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLength sets the value of property "HTMLSelectElement.length" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetLength(val uint32) bool {
	return js.True == bindings.SetHTMLSelectElementLength(
		this.ref,
		uint32(val),
	)
}

// SelectedOptions returns the value of property "HTMLSelectElement.selectedOptions".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) SelectedOptions() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementSelectedOptions(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SelectedIndex returns the value of property "HTMLSelectElement.selectedIndex".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) SelectedIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementSelectedIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectedIndex sets the value of property "HTMLSelectElement.selectedIndex" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetSelectedIndex(val int32) bool {
	return js.True == bindings.SetHTMLSelectElementSelectedIndex(
		this.ref,
		int32(val),
	)
}

// Value returns the value of property "HTMLSelectElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLSelectElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSelectElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLSelectElementValue(
		this.ref,
		val.Ref(),
	)
}

// WillValidate returns the value of property "HTMLSelectElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementWillValidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLSelectElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementValidity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLSelectElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementValidationMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLSelectElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLSelectElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLSelectElementLabels(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncItem returns true if the method "HTMLSelectElement.item" exists.
func (this HTMLSelectElement) HasFuncItem() bool {
	return js.True == bindings.HasFuncHTMLSelectElementItem(
		this.ref,
	)
}

// FuncItem returns the method "HTMLSelectElement.item".
func (this HTMLSelectElement) FuncItem() (fn js.Func[func(index uint32) HTMLOptionElement]) {
	bindings.FuncHTMLSelectElementItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Item calls the method "HTMLSelectElement.item".
func (this HTMLSelectElement) Item(index uint32) (ret HTMLOptionElement) {
	bindings.CallHTMLSelectElementItem(
		this.ref, js.Pointer(&ret),
		uint32(index),
	)

	return
}

// TryItem calls the method "HTMLSelectElement.item"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSelectElement) TryItem(index uint32) (ret HTMLOptionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
	)

	return
}

// HasFuncNamedItem returns true if the method "HTMLSelectElement.namedItem" exists.
func (this HTMLSelectElement) HasFuncNamedItem() bool {
	return js.True == bindings.HasFuncHTMLSelectElementNamedItem(
		this.ref,
	)
}

// FuncNamedItem returns the method "HTMLSelectElement.namedItem".
func (this HTMLSelectElement) FuncNamedItem() (fn js.Func[func(name js.String) HTMLOptionElement]) {
	bindings.FuncHTMLSelectElementNamedItem(
		this.ref, js.Pointer(&fn),
	)
	return
}

// NamedItem calls the method "HTMLSelectElement.namedItem".
func (this HTMLSelectElement) NamedItem(name js.String) (ret HTMLOptionElement) {
	bindings.CallHTMLSelectElementNamedItem(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryNamedItem calls the method "HTMLSelectElement.namedItem"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSelectElement) TryNamedItem(name js.String) (ret HTMLOptionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementNamedItem(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncAdd returns true if the method "HTMLSelectElement.add" exists.
func (this HTMLSelectElement) HasFuncAdd() bool {
	return js.True == bindings.HasFuncHTMLSelectElementAdd(
		this.ref,
	)
}

// FuncAdd returns the method "HTMLSelectElement.add".
func (this HTMLSelectElement) FuncAdd() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32)]) {
	bindings.FuncHTMLSelectElementAdd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add calls the method "HTMLSelectElement.add".
func (this HTMLSelectElement) Add(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32) (ret js.Void) {
	bindings.CallHTMLSelectElementAdd(
		this.ref, js.Pointer(&ret),
		element.Ref(),
		before.Ref(),
	)

	return
}

// TryAdd calls the method "HTMLSelectElement.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSelectElement) TryAdd(element OneOf_HTMLOptionElement_HTMLOptGroupElement, before OneOf_HTMLElement_Int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementAdd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
		before.Ref(),
	)

	return
}

// HasFuncAdd1 returns true if the method "HTMLSelectElement.add" exists.
func (this HTMLSelectElement) HasFuncAdd1() bool {
	return js.True == bindings.HasFuncHTMLSelectElementAdd1(
		this.ref,
	)
}

// FuncAdd1 returns the method "HTMLSelectElement.add".
func (this HTMLSelectElement) FuncAdd1() (fn js.Func[func(element OneOf_HTMLOptionElement_HTMLOptGroupElement)]) {
	bindings.FuncHTMLSelectElementAdd1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add1 calls the method "HTMLSelectElement.add".
func (this HTMLSelectElement) Add1(element OneOf_HTMLOptionElement_HTMLOptGroupElement) (ret js.Void) {
	bindings.CallHTMLSelectElementAdd1(
		this.ref, js.Pointer(&ret),
		element.Ref(),
	)

	return
}

// TryAdd1 calls the method "HTMLSelectElement.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSelectElement) TryAdd1(element OneOf_HTMLOptionElement_HTMLOptGroupElement) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementAdd1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		element.Ref(),
	)

	return
}

// HasFuncRemove returns true if the method "HTMLSelectElement.remove" exists.
func (this HTMLSelectElement) HasFuncRemove() bool {
	return js.True == bindings.HasFuncHTMLSelectElementRemove(
		this.ref,
	)
}

// FuncRemove returns the method "HTMLSelectElement.remove".
func (this HTMLSelectElement) FuncRemove() (fn js.Func[func()]) {
	bindings.FuncHTMLSelectElementRemove(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove calls the method "HTMLSelectElement.remove".
func (this HTMLSelectElement) Remove() (ret js.Void) {
	bindings.CallHTMLSelectElementRemove(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryRemove calls the method "HTMLSelectElement.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSelectElement) TryRemove() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementRemove(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncRemove1 returns true if the method "HTMLSelectElement.remove" exists.
func (this HTMLSelectElement) HasFuncRemove1() bool {
	return js.True == bindings.HasFuncHTMLSelectElementRemove1(
		this.ref,
	)
}

// FuncRemove1 returns the method "HTMLSelectElement.remove".
func (this HTMLSelectElement) FuncRemove1() (fn js.Func[func(index int32)]) {
	bindings.FuncHTMLSelectElementRemove1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Remove1 calls the method "HTMLSelectElement.remove".
func (this HTMLSelectElement) Remove1(index int32) (ret js.Void) {
	bindings.CallHTMLSelectElementRemove1(
		this.ref, js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryRemove1 calls the method "HTMLSelectElement.remove"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSelectElement) TryRemove1(index int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementRemove1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

// HasFuncSet returns true if the method "HTMLSelectElement." exists.
func (this HTMLSelectElement) HasFuncSet() bool {
	return js.True == bindings.HasFuncHTMLSelectElementSet(
		this.ref,
	)
}

// FuncSet returns the method "HTMLSelectElement.".
func (this HTMLSelectElement) FuncSet() (fn js.Func[func(index uint32, option HTMLOptionElement)]) {
	bindings.FuncHTMLSelectElementSet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Set calls the method "HTMLSelectElement.".
func (this HTMLSelectElement) Set(index uint32, option HTMLOptionElement) (ret js.Void) {
	bindings.CallHTMLSelectElementSet(
		this.ref, js.Pointer(&ret),
		uint32(index),
		option.Ref(),
	)

	return
}

// TrySet calls the method "HTMLSelectElement."
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSelectElement) TrySet(index uint32, option HTMLOptionElement) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementSet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(index),
		option.Ref(),
	)

	return
}

// HasFuncCheckValidity returns true if the method "HTMLSelectElement.checkValidity" exists.
func (this HTMLSelectElement) HasFuncCheckValidity() bool {
	return js.True == bindings.HasFuncHTMLSelectElementCheckValidity(
		this.ref,
	)
}

// FuncCheckValidity returns the method "HTMLSelectElement.checkValidity".
func (this HTMLSelectElement) FuncCheckValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLSelectElementCheckValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckValidity calls the method "HTMLSelectElement.checkValidity".
func (this HTMLSelectElement) CheckValidity() (ret bool) {
	bindings.CallHTMLSelectElementCheckValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLSelectElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSelectElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementCheckValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReportValidity returns true if the method "HTMLSelectElement.reportValidity" exists.
func (this HTMLSelectElement) HasFuncReportValidity() bool {
	return js.True == bindings.HasFuncHTMLSelectElementReportValidity(
		this.ref,
	)
}

// FuncReportValidity returns the method "HTMLSelectElement.reportValidity".
func (this HTMLSelectElement) FuncReportValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLSelectElementReportValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportValidity calls the method "HTMLSelectElement.reportValidity".
func (this HTMLSelectElement) ReportValidity() (ret bool) {
	bindings.CallHTMLSelectElementReportValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLSelectElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSelectElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementReportValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetCustomValidity returns true if the method "HTMLSelectElement.setCustomValidity" exists.
func (this HTMLSelectElement) HasFuncSetCustomValidity() bool {
	return js.True == bindings.HasFuncHTMLSelectElementSetCustomValidity(
		this.ref,
	)
}

// FuncSetCustomValidity returns the method "HTMLSelectElement.setCustomValidity".
func (this HTMLSelectElement) FuncSetCustomValidity() (fn js.Func[func(err js.String)]) {
	bindings.FuncHTMLSelectElementSetCustomValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetCustomValidity calls the method "HTMLSelectElement.setCustomValidity".
func (this HTMLSelectElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLSelectElementSetCustomValidity(
		this.ref, js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLSelectElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLSelectElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLSelectElementSetCustomValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

type HTMLSourceElement struct {
	HTMLElement
}

func (this HTMLSourceElement) Once() HTMLSourceElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Src returns the value of property "HTMLSourceElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLSourceElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLSourceElementSrc(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLSourceElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLSourceElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLSourceElementType(
		this.ref,
		val.Ref(),
	)
}

// Srcset returns the value of property "HTMLSourceElement.srcset".
//
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Srcset() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementSrcset(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrcset sets the value of property "HTMLSourceElement.srcset" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetSrcset(val js.String) bool {
	return js.True == bindings.SetHTMLSourceElementSrcset(
		this.ref,
		val.Ref(),
	)
}

// Sizes returns the value of property "HTMLSourceElement.sizes".
//
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Sizes() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementSizes(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSizes sets the value of property "HTMLSourceElement.sizes" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetSizes(val js.String) bool {
	return js.True == bindings.SetHTMLSourceElementSizes(
		this.ref,
		val.Ref(),
	)
}

// Media returns the value of property "HTMLSourceElement.media".
//
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementMedia(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMedia sets the value of property "HTMLSourceElement.media" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetMedia(val js.String) bool {
	return js.True == bindings.SetHTMLSourceElementMedia(
		this.ref,
		val.Ref(),
	)
}

// Width returns the value of property "HTMLSourceElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Width() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLSourceElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetWidth(val uint32) bool {
	return js.True == bindings.SetHTMLSourceElementWidth(
		this.ref,
		uint32(val),
	)
}

// Height returns the value of property "HTMLSourceElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLSourceElement) Height() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLSourceElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLSourceElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLSourceElement) SetHeight(val uint32) bool {
	return js.True == bindings.SetHTMLSourceElementHeight(
		this.ref,
		uint32(val),
	)
}

type HTMLSpanElement struct {
	HTMLElement
}

func (this HTMLSpanElement) Once() HTMLSpanElement {
	this.ref.Once()
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
	this.ref.Free()
}

type HTMLString = js.String

type HTMLStyleElement struct {
	HTMLElement
}

func (this HTMLStyleElement) Once() HTMLStyleElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Disabled returns the value of property "HTMLStyleElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLStyleElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLStyleElementDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLStyleElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLStyleElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLStyleElementDisabled(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Media returns the value of property "HTMLStyleElement.media".
//
// It returns ok=false if there is no such property.
func (this HTMLStyleElement) Media() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLStyleElementMedia(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMedia sets the value of property "HTMLStyleElement.media" to val.
//
// It returns false if the property cannot be set.
func (this HTMLStyleElement) SetMedia(val js.String) bool {
	return js.True == bindings.SetHTMLStyleElementMedia(
		this.ref,
		val.Ref(),
	)
}

// Blocking returns the value of property "HTMLStyleElement.blocking".
//
// It returns ok=false if there is no such property.
func (this HTMLStyleElement) Blocking() (ret DOMTokenList, ok bool) {
	ok = js.True == bindings.GetHTMLStyleElementBlocking(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Type returns the value of property "HTMLStyleElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLStyleElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLStyleElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLStyleElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLStyleElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLStyleElementType(
		this.ref,
		val.Ref(),
	)
}

// Sheet returns the value of property "HTMLStyleElement.sheet".
//
// It returns ok=false if there is no such property.
func (this HTMLStyleElement) Sheet() (ret CSSStyleSheet, ok bool) {
	ok = js.True == bindings.GetHTMLStyleElementSheet(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HTMLTableCaptionElement struct {
	HTMLElement
}

func (this HTMLTableCaptionElement) Once() HTMLTableCaptionElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Align returns the value of property "HTMLTableCaptionElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCaptionElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCaptionElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableCaptionElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCaptionElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableCaptionElementAlign(
		this.ref,
		val.Ref(),
	)
}

type HTMLTableCellElement struct {
	HTMLElement
}

func (this HTMLTableCellElement) Once() HTMLTableCellElement {
	this.ref.Once()
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
	this.ref.Free()
}

// ColSpan returns the value of property "HTMLTableCellElement.colSpan".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) ColSpan() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementColSpan(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetColSpan sets the value of property "HTMLTableCellElement.colSpan" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetColSpan(val uint32) bool {
	return js.True == bindings.SetHTMLTableCellElementColSpan(
		this.ref,
		uint32(val),
	)
}

// RowSpan returns the value of property "HTMLTableCellElement.rowSpan".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) RowSpan() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementRowSpan(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRowSpan sets the value of property "HTMLTableCellElement.rowSpan" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetRowSpan(val uint32) bool {
	return js.True == bindings.SetHTMLTableCellElementRowSpan(
		this.ref,
		uint32(val),
	)
}

// Headers returns the value of property "HTMLTableCellElement.headers".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Headers() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementHeaders(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeaders sets the value of property "HTMLTableCellElement.headers" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetHeaders(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementHeaders(
		this.ref,
		val.Ref(),
	)
}

// CellIndex returns the value of property "HTMLTableCellElement.cellIndex".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) CellIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementCellIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Scope returns the value of property "HTMLTableCellElement.scope".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Scope() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementScope(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetScope sets the value of property "HTMLTableCellElement.scope" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetScope(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementScope(
		this.ref,
		val.Ref(),
	)
}

// Abbr returns the value of property "HTMLTableCellElement.abbr".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Abbr() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementAbbr(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAbbr sets the value of property "HTMLTableCellElement.abbr" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetAbbr(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementAbbr(
		this.ref,
		val.Ref(),
	)
}

// Align returns the value of property "HTMLTableCellElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableCellElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementAlign(
		this.ref,
		val.Ref(),
	)
}

// Axis returns the value of property "HTMLTableCellElement.axis".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Axis() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementAxis(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAxis sets the value of property "HTMLTableCellElement.axis" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetAxis(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementAxis(
		this.ref,
		val.Ref(),
	)
}

// Height returns the value of property "HTMLTableCellElement.height".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Height() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementHeight(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetHeight sets the value of property "HTMLTableCellElement.height" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetHeight(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementHeight(
		this.ref,
		val.Ref(),
	)
}

// Width returns the value of property "HTMLTableCellElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLTableCellElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementWidth(
		this.ref,
		val.Ref(),
	)
}

// Ch returns the value of property "HTMLTableCellElement.ch".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) Ch() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementCh(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCh sets the value of property "HTMLTableCellElement.ch" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetCh(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementCh(
		this.ref,
		val.Ref(),
	)
}

// ChOff returns the value of property "HTMLTableCellElement.chOff".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) ChOff() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementChOff(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChOff sets the value of property "HTMLTableCellElement.chOff" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetChOff(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementChOff(
		this.ref,
		val.Ref(),
	)
}

// NoWrap returns the value of property "HTMLTableCellElement.noWrap".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) NoWrap() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementNoWrap(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetNoWrap sets the value of property "HTMLTableCellElement.noWrap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetNoWrap(val bool) bool {
	return js.True == bindings.SetHTMLTableCellElementNoWrap(
		this.ref,
		js.Bool(bool(val)),
	)
}

// VAlign returns the value of property "HTMLTableCellElement.vAlign".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) VAlign() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementVAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVAlign sets the value of property "HTMLTableCellElement.vAlign" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetVAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementVAlign(
		this.ref,
		val.Ref(),
	)
}

// BgColor returns the value of property "HTMLTableCellElement.bgColor".
//
// It returns ok=false if there is no such property.
func (this HTMLTableCellElement) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableCellElementBgColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "HTMLTableCellElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableCellElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLTableCellElementBgColor(
		this.ref,
		val.Ref(),
	)
}

type HTMLTableColElement struct {
	HTMLElement
}

func (this HTMLTableColElement) Once() HTMLTableColElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Span returns the value of property "HTMLTableColElement.span".
//
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) Span() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementSpan(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSpan sets the value of property "HTMLTableColElement.span" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetSpan(val uint32) bool {
	return js.True == bindings.SetHTMLTableColElementSpan(
		this.ref,
		uint32(val),
	)
}

// Align returns the value of property "HTMLTableColElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableColElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableColElementAlign(
		this.ref,
		val.Ref(),
	)
}

// Ch returns the value of property "HTMLTableColElement.ch".
//
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) Ch() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementCh(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCh sets the value of property "HTMLTableColElement.ch" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetCh(val js.String) bool {
	return js.True == bindings.SetHTMLTableColElementCh(
		this.ref,
		val.Ref(),
	)
}

// ChOff returns the value of property "HTMLTableColElement.chOff".
//
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) ChOff() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementChOff(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChOff sets the value of property "HTMLTableColElement.chOff" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetChOff(val js.String) bool {
	return js.True == bindings.SetHTMLTableColElementChOff(
		this.ref,
		val.Ref(),
	)
}

// VAlign returns the value of property "HTMLTableColElement.vAlign".
//
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) VAlign() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementVAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVAlign sets the value of property "HTMLTableColElement.vAlign" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetVAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableColElementVAlign(
		this.ref,
		val.Ref(),
	)
}

// Width returns the value of property "HTMLTableColElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLTableColElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableColElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLTableColElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableColElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLTableColElementWidth(
		this.ref,
		val.Ref(),
	)
}

type HTMLTableRowElement struct {
	HTMLElement
}

func (this HTMLTableRowElement) Once() HTMLTableRowElement {
	this.ref.Once()
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
	this.ref.Free()
}

// RowIndex returns the value of property "HTMLTableRowElement.rowIndex".
//
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) RowIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementRowIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SectionRowIndex returns the value of property "HTMLTableRowElement.sectionRowIndex".
//
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) SectionRowIndex() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementSectionRowIndex(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Cells returns the value of property "HTMLTableRowElement.cells".
//
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) Cells() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementCells(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Align returns the value of property "HTMLTableRowElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableRowElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableRowElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableRowElementAlign(
		this.ref,
		val.Ref(),
	)
}

// Ch returns the value of property "HTMLTableRowElement.ch".
//
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) Ch() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementCh(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCh sets the value of property "HTMLTableRowElement.ch" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableRowElement) SetCh(val js.String) bool {
	return js.True == bindings.SetHTMLTableRowElementCh(
		this.ref,
		val.Ref(),
	)
}

// ChOff returns the value of property "HTMLTableRowElement.chOff".
//
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) ChOff() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementChOff(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChOff sets the value of property "HTMLTableRowElement.chOff" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableRowElement) SetChOff(val js.String) bool {
	return js.True == bindings.SetHTMLTableRowElementChOff(
		this.ref,
		val.Ref(),
	)
}

// VAlign returns the value of property "HTMLTableRowElement.vAlign".
//
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) VAlign() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementVAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVAlign sets the value of property "HTMLTableRowElement.vAlign" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableRowElement) SetVAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableRowElementVAlign(
		this.ref,
		val.Ref(),
	)
}

// BgColor returns the value of property "HTMLTableRowElement.bgColor".
//
// It returns ok=false if there is no such property.
func (this HTMLTableRowElement) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableRowElementBgColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "HTMLTableRowElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableRowElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLTableRowElementBgColor(
		this.ref,
		val.Ref(),
	)
}

// HasFuncInsertCell returns true if the method "HTMLTableRowElement.insertCell" exists.
func (this HTMLTableRowElement) HasFuncInsertCell() bool {
	return js.True == bindings.HasFuncHTMLTableRowElementInsertCell(
		this.ref,
	)
}

// FuncInsertCell returns the method "HTMLTableRowElement.insertCell".
func (this HTMLTableRowElement) FuncInsertCell() (fn js.Func[func(index int32) HTMLTableCellElement]) {
	bindings.FuncHTMLTableRowElementInsertCell(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertCell calls the method "HTMLTableRowElement.insertCell".
func (this HTMLTableRowElement) InsertCell(index int32) (ret HTMLTableCellElement) {
	bindings.CallHTMLTableRowElementInsertCell(
		this.ref, js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryInsertCell calls the method "HTMLTableRowElement.insertCell"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableRowElement) TryInsertCell(index int32) (ret HTMLTableCellElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableRowElementInsertCell(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

// HasFuncInsertCell1 returns true if the method "HTMLTableRowElement.insertCell" exists.
func (this HTMLTableRowElement) HasFuncInsertCell1() bool {
	return js.True == bindings.HasFuncHTMLTableRowElementInsertCell1(
		this.ref,
	)
}

// FuncInsertCell1 returns the method "HTMLTableRowElement.insertCell".
func (this HTMLTableRowElement) FuncInsertCell1() (fn js.Func[func() HTMLTableCellElement]) {
	bindings.FuncHTMLTableRowElementInsertCell1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertCell1 calls the method "HTMLTableRowElement.insertCell".
func (this HTMLTableRowElement) InsertCell1() (ret HTMLTableCellElement) {
	bindings.CallHTMLTableRowElementInsertCell1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryInsertCell1 calls the method "HTMLTableRowElement.insertCell"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableRowElement) TryInsertCell1() (ret HTMLTableCellElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableRowElementInsertCell1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteCell returns true if the method "HTMLTableRowElement.deleteCell" exists.
func (this HTMLTableRowElement) HasFuncDeleteCell() bool {
	return js.True == bindings.HasFuncHTMLTableRowElementDeleteCell(
		this.ref,
	)
}

// FuncDeleteCell returns the method "HTMLTableRowElement.deleteCell".
func (this HTMLTableRowElement) FuncDeleteCell() (fn js.Func[func(index int32)]) {
	bindings.FuncHTMLTableRowElementDeleteCell(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteCell calls the method "HTMLTableRowElement.deleteCell".
func (this HTMLTableRowElement) DeleteCell(index int32) (ret js.Void) {
	bindings.CallHTMLTableRowElementDeleteCell(
		this.ref, js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryDeleteCell calls the method "HTMLTableRowElement.deleteCell"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableRowElement) TryDeleteCell(index int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableRowElementDeleteCell(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

type HTMLTableSectionElement struct {
	HTMLElement
}

func (this HTMLTableSectionElement) Once() HTMLTableSectionElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Rows returns the value of property "HTMLTableSectionElement.rows".
//
// It returns ok=false if there is no such property.
func (this HTMLTableSectionElement) Rows() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLTableSectionElementRows(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Align returns the value of property "HTMLTableSectionElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLTableSectionElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableSectionElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableSectionElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableSectionElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableSectionElementAlign(
		this.ref,
		val.Ref(),
	)
}

// Ch returns the value of property "HTMLTableSectionElement.ch".
//
// It returns ok=false if there is no such property.
func (this HTMLTableSectionElement) Ch() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableSectionElementCh(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCh sets the value of property "HTMLTableSectionElement.ch" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableSectionElement) SetCh(val js.String) bool {
	return js.True == bindings.SetHTMLTableSectionElementCh(
		this.ref,
		val.Ref(),
	)
}

// ChOff returns the value of property "HTMLTableSectionElement.chOff".
//
// It returns ok=false if there is no such property.
func (this HTMLTableSectionElement) ChOff() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableSectionElementChOff(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetChOff sets the value of property "HTMLTableSectionElement.chOff" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableSectionElement) SetChOff(val js.String) bool {
	return js.True == bindings.SetHTMLTableSectionElementChOff(
		this.ref,
		val.Ref(),
	)
}

// VAlign returns the value of property "HTMLTableSectionElement.vAlign".
//
// It returns ok=false if there is no such property.
func (this HTMLTableSectionElement) VAlign() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableSectionElementVAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetVAlign sets the value of property "HTMLTableSectionElement.vAlign" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableSectionElement) SetVAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableSectionElementVAlign(
		this.ref,
		val.Ref(),
	)
}

// HasFuncInsertRow returns true if the method "HTMLTableSectionElement.insertRow" exists.
func (this HTMLTableSectionElement) HasFuncInsertRow() bool {
	return js.True == bindings.HasFuncHTMLTableSectionElementInsertRow(
		this.ref,
	)
}

// FuncInsertRow returns the method "HTMLTableSectionElement.insertRow".
func (this HTMLTableSectionElement) FuncInsertRow() (fn js.Func[func(index int32) HTMLTableRowElement]) {
	bindings.FuncHTMLTableSectionElementInsertRow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertRow calls the method "HTMLTableSectionElement.insertRow".
func (this HTMLTableSectionElement) InsertRow(index int32) (ret HTMLTableRowElement) {
	bindings.CallHTMLTableSectionElementInsertRow(
		this.ref, js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryInsertRow calls the method "HTMLTableSectionElement.insertRow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableSectionElement) TryInsertRow(index int32) (ret HTMLTableRowElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableSectionElementInsertRow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

// HasFuncInsertRow1 returns true if the method "HTMLTableSectionElement.insertRow" exists.
func (this HTMLTableSectionElement) HasFuncInsertRow1() bool {
	return js.True == bindings.HasFuncHTMLTableSectionElementInsertRow1(
		this.ref,
	)
}

// FuncInsertRow1 returns the method "HTMLTableSectionElement.insertRow".
func (this HTMLTableSectionElement) FuncInsertRow1() (fn js.Func[func() HTMLTableRowElement]) {
	bindings.FuncHTMLTableSectionElementInsertRow1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertRow1 calls the method "HTMLTableSectionElement.insertRow".
func (this HTMLTableSectionElement) InsertRow1() (ret HTMLTableRowElement) {
	bindings.CallHTMLTableSectionElementInsertRow1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryInsertRow1 calls the method "HTMLTableSectionElement.insertRow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableSectionElement) TryInsertRow1() (ret HTMLTableRowElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableSectionElementInsertRow1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteRow returns true if the method "HTMLTableSectionElement.deleteRow" exists.
func (this HTMLTableSectionElement) HasFuncDeleteRow() bool {
	return js.True == bindings.HasFuncHTMLTableSectionElementDeleteRow(
		this.ref,
	)
}

// FuncDeleteRow returns the method "HTMLTableSectionElement.deleteRow".
func (this HTMLTableSectionElement) FuncDeleteRow() (fn js.Func[func(index int32)]) {
	bindings.FuncHTMLTableSectionElementDeleteRow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteRow calls the method "HTMLTableSectionElement.deleteRow".
func (this HTMLTableSectionElement) DeleteRow(index int32) (ret js.Void) {
	bindings.CallHTMLTableSectionElementDeleteRow(
		this.ref, js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryDeleteRow calls the method "HTMLTableSectionElement.deleteRow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableSectionElement) TryDeleteRow(index int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableSectionElementDeleteRow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

type HTMLTableElement struct {
	HTMLElement
}

func (this HTMLTableElement) Once() HTMLTableElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Caption returns the value of property "HTMLTableElement.caption".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Caption() (ret HTMLTableCaptionElement, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementCaption(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCaption sets the value of property "HTMLTableElement.caption" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetCaption(val HTMLTableCaptionElement) bool {
	return js.True == bindings.SetHTMLTableElementCaption(
		this.ref,
		val.Ref(),
	)
}

// THead returns the value of property "HTMLTableElement.tHead".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) THead() (ret HTMLTableSectionElement, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementTHead(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTHead sets the value of property "HTMLTableElement.tHead" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetTHead(val HTMLTableSectionElement) bool {
	return js.True == bindings.SetHTMLTableElementTHead(
		this.ref,
		val.Ref(),
	)
}

// TFoot returns the value of property "HTMLTableElement.tFoot".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) TFoot() (ret HTMLTableSectionElement, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementTFoot(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetTFoot sets the value of property "HTMLTableElement.tFoot" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetTFoot(val HTMLTableSectionElement) bool {
	return js.True == bindings.SetHTMLTableElementTFoot(
		this.ref,
		val.Ref(),
	)
}

// TBodies returns the value of property "HTMLTableElement.tBodies".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) TBodies() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementTBodies(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Rows returns the value of property "HTMLTableElement.rows".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Rows() (ret HTMLCollection, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementRows(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Align returns the value of property "HTMLTableElement.align".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Align() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementAlign(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAlign sets the value of property "HTMLTableElement.align" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetAlign(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementAlign(
		this.ref,
		val.Ref(),
	)
}

// Border returns the value of property "HTMLTableElement.border".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Border() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementBorder(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBorder sets the value of property "HTMLTableElement.border" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetBorder(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementBorder(
		this.ref,
		val.Ref(),
	)
}

// Frame returns the value of property "HTMLTableElement.frame".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Frame() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementFrame(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetFrame sets the value of property "HTMLTableElement.frame" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetFrame(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementFrame(
		this.ref,
		val.Ref(),
	)
}

// Rules returns the value of property "HTMLTableElement.rules".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Rules() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementRules(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRules sets the value of property "HTMLTableElement.rules" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetRules(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementRules(
		this.ref,
		val.Ref(),
	)
}

// Summary returns the value of property "HTMLTableElement.summary".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Summary() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementSummary(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSummary sets the value of property "HTMLTableElement.summary" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetSummary(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementSummary(
		this.ref,
		val.Ref(),
	)
}

// Width returns the value of property "HTMLTableElement.width".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) Width() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementWidth(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWidth sets the value of property "HTMLTableElement.width" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetWidth(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementWidth(
		this.ref,
		val.Ref(),
	)
}

// BgColor returns the value of property "HTMLTableElement.bgColor".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) BgColor() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementBgColor(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetBgColor sets the value of property "HTMLTableElement.bgColor" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetBgColor(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementBgColor(
		this.ref,
		val.Ref(),
	)
}

// CellPadding returns the value of property "HTMLTableElement.cellPadding".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) CellPadding() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementCellPadding(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCellPadding sets the value of property "HTMLTableElement.cellPadding" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetCellPadding(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementCellPadding(
		this.ref,
		val.Ref(),
	)
}

// CellSpacing returns the value of property "HTMLTableElement.cellSpacing".
//
// It returns ok=false if there is no such property.
func (this HTMLTableElement) CellSpacing() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTableElementCellSpacing(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCellSpacing sets the value of property "HTMLTableElement.cellSpacing" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTableElement) SetCellSpacing(val js.String) bool {
	return js.True == bindings.SetHTMLTableElementCellSpacing(
		this.ref,
		val.Ref(),
	)
}

// HasFuncCreateCaption returns true if the method "HTMLTableElement.createCaption" exists.
func (this HTMLTableElement) HasFuncCreateCaption() bool {
	return js.True == bindings.HasFuncHTMLTableElementCreateCaption(
		this.ref,
	)
}

// FuncCreateCaption returns the method "HTMLTableElement.createCaption".
func (this HTMLTableElement) FuncCreateCaption() (fn js.Func[func() HTMLTableCaptionElement]) {
	bindings.FuncHTMLTableElementCreateCaption(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateCaption calls the method "HTMLTableElement.createCaption".
func (this HTMLTableElement) CreateCaption() (ret HTMLTableCaptionElement) {
	bindings.CallHTMLTableElementCreateCaption(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateCaption calls the method "HTMLTableElement.createCaption"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableElement) TryCreateCaption() (ret HTMLTableCaptionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementCreateCaption(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteCaption returns true if the method "HTMLTableElement.deleteCaption" exists.
func (this HTMLTableElement) HasFuncDeleteCaption() bool {
	return js.True == bindings.HasFuncHTMLTableElementDeleteCaption(
		this.ref,
	)
}

// FuncDeleteCaption returns the method "HTMLTableElement.deleteCaption".
func (this HTMLTableElement) FuncDeleteCaption() (fn js.Func[func()]) {
	bindings.FuncHTMLTableElementDeleteCaption(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteCaption calls the method "HTMLTableElement.deleteCaption".
func (this HTMLTableElement) DeleteCaption() (ret js.Void) {
	bindings.CallHTMLTableElementDeleteCaption(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeleteCaption calls the method "HTMLTableElement.deleteCaption"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableElement) TryDeleteCaption() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementDeleteCaption(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateTHead returns true if the method "HTMLTableElement.createTHead" exists.
func (this HTMLTableElement) HasFuncCreateTHead() bool {
	return js.True == bindings.HasFuncHTMLTableElementCreateTHead(
		this.ref,
	)
}

// FuncCreateTHead returns the method "HTMLTableElement.createTHead".
func (this HTMLTableElement) FuncCreateTHead() (fn js.Func[func() HTMLTableSectionElement]) {
	bindings.FuncHTMLTableElementCreateTHead(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateTHead calls the method "HTMLTableElement.createTHead".
func (this HTMLTableElement) CreateTHead() (ret HTMLTableSectionElement) {
	bindings.CallHTMLTableElementCreateTHead(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateTHead calls the method "HTMLTableElement.createTHead"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableElement) TryCreateTHead() (ret HTMLTableSectionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementCreateTHead(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteTHead returns true if the method "HTMLTableElement.deleteTHead" exists.
func (this HTMLTableElement) HasFuncDeleteTHead() bool {
	return js.True == bindings.HasFuncHTMLTableElementDeleteTHead(
		this.ref,
	)
}

// FuncDeleteTHead returns the method "HTMLTableElement.deleteTHead".
func (this HTMLTableElement) FuncDeleteTHead() (fn js.Func[func()]) {
	bindings.FuncHTMLTableElementDeleteTHead(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteTHead calls the method "HTMLTableElement.deleteTHead".
func (this HTMLTableElement) DeleteTHead() (ret js.Void) {
	bindings.CallHTMLTableElementDeleteTHead(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeleteTHead calls the method "HTMLTableElement.deleteTHead"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableElement) TryDeleteTHead() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementDeleteTHead(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateTFoot returns true if the method "HTMLTableElement.createTFoot" exists.
func (this HTMLTableElement) HasFuncCreateTFoot() bool {
	return js.True == bindings.HasFuncHTMLTableElementCreateTFoot(
		this.ref,
	)
}

// FuncCreateTFoot returns the method "HTMLTableElement.createTFoot".
func (this HTMLTableElement) FuncCreateTFoot() (fn js.Func[func() HTMLTableSectionElement]) {
	bindings.FuncHTMLTableElementCreateTFoot(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateTFoot calls the method "HTMLTableElement.createTFoot".
func (this HTMLTableElement) CreateTFoot() (ret HTMLTableSectionElement) {
	bindings.CallHTMLTableElementCreateTFoot(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateTFoot calls the method "HTMLTableElement.createTFoot"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableElement) TryCreateTFoot() (ret HTMLTableSectionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementCreateTFoot(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteTFoot returns true if the method "HTMLTableElement.deleteTFoot" exists.
func (this HTMLTableElement) HasFuncDeleteTFoot() bool {
	return js.True == bindings.HasFuncHTMLTableElementDeleteTFoot(
		this.ref,
	)
}

// FuncDeleteTFoot returns the method "HTMLTableElement.deleteTFoot".
func (this HTMLTableElement) FuncDeleteTFoot() (fn js.Func[func()]) {
	bindings.FuncHTMLTableElementDeleteTFoot(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteTFoot calls the method "HTMLTableElement.deleteTFoot".
func (this HTMLTableElement) DeleteTFoot() (ret js.Void) {
	bindings.CallHTMLTableElementDeleteTFoot(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryDeleteTFoot calls the method "HTMLTableElement.deleteTFoot"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableElement) TryDeleteTFoot() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementDeleteTFoot(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateTBody returns true if the method "HTMLTableElement.createTBody" exists.
func (this HTMLTableElement) HasFuncCreateTBody() bool {
	return js.True == bindings.HasFuncHTMLTableElementCreateTBody(
		this.ref,
	)
}

// FuncCreateTBody returns the method "HTMLTableElement.createTBody".
func (this HTMLTableElement) FuncCreateTBody() (fn js.Func[func() HTMLTableSectionElement]) {
	bindings.FuncHTMLTableElementCreateTBody(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateTBody calls the method "HTMLTableElement.createTBody".
func (this HTMLTableElement) CreateTBody() (ret HTMLTableSectionElement) {
	bindings.CallHTMLTableElementCreateTBody(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCreateTBody calls the method "HTMLTableElement.createTBody"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableElement) TryCreateTBody() (ret HTMLTableSectionElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementCreateTBody(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncInsertRow returns true if the method "HTMLTableElement.insertRow" exists.
func (this HTMLTableElement) HasFuncInsertRow() bool {
	return js.True == bindings.HasFuncHTMLTableElementInsertRow(
		this.ref,
	)
}

// FuncInsertRow returns the method "HTMLTableElement.insertRow".
func (this HTMLTableElement) FuncInsertRow() (fn js.Func[func(index int32) HTMLTableRowElement]) {
	bindings.FuncHTMLTableElementInsertRow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertRow calls the method "HTMLTableElement.insertRow".
func (this HTMLTableElement) InsertRow(index int32) (ret HTMLTableRowElement) {
	bindings.CallHTMLTableElementInsertRow(
		this.ref, js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryInsertRow calls the method "HTMLTableElement.insertRow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableElement) TryInsertRow(index int32) (ret HTMLTableRowElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementInsertRow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

// HasFuncInsertRow1 returns true if the method "HTMLTableElement.insertRow" exists.
func (this HTMLTableElement) HasFuncInsertRow1() bool {
	return js.True == bindings.HasFuncHTMLTableElementInsertRow1(
		this.ref,
	)
}

// FuncInsertRow1 returns the method "HTMLTableElement.insertRow".
func (this HTMLTableElement) FuncInsertRow1() (fn js.Func[func() HTMLTableRowElement]) {
	bindings.FuncHTMLTableElementInsertRow1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// InsertRow1 calls the method "HTMLTableElement.insertRow".
func (this HTMLTableElement) InsertRow1() (ret HTMLTableRowElement) {
	bindings.CallHTMLTableElementInsertRow1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryInsertRow1 calls the method "HTMLTableElement.insertRow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableElement) TryInsertRow1() (ret HTMLTableRowElement, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementInsertRow1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncDeleteRow returns true if the method "HTMLTableElement.deleteRow" exists.
func (this HTMLTableElement) HasFuncDeleteRow() bool {
	return js.True == bindings.HasFuncHTMLTableElementDeleteRow(
		this.ref,
	)
}

// FuncDeleteRow returns the method "HTMLTableElement.deleteRow".
func (this HTMLTableElement) FuncDeleteRow() (fn js.Func[func(index int32)]) {
	bindings.FuncHTMLTableElementDeleteRow(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteRow calls the method "HTMLTableElement.deleteRow".
func (this HTMLTableElement) DeleteRow(index int32) (ret js.Void) {
	bindings.CallHTMLTableElementDeleteRow(
		this.ref, js.Pointer(&ret),
		int32(index),
	)

	return
}

// TryDeleteRow calls the method "HTMLTableElement.deleteRow"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTableElement) TryDeleteRow(index int32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTableElementDeleteRow(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		int32(index),
	)

	return
}

type HTMLTemplateElement struct {
	HTMLElement
}

func (this HTMLTemplateElement) Once() HTMLTemplateElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Content returns the value of property "HTMLTemplateElement.content".
//
// It returns ok=false if there is no such property.
func (this HTMLTemplateElement) Content() (ret DocumentFragment, ok bool) {
	ok = js.True == bindings.GetHTMLTemplateElementContent(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HTMLTextAreaElement struct {
	HTMLElement
}

func (this HTMLTextAreaElement) Once() HTMLTextAreaElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Autocomplete returns the value of property "HTMLTextAreaElement.autocomplete".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Autocomplete() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementAutocomplete(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetAutocomplete sets the value of property "HTMLTextAreaElement.autocomplete" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetAutocomplete(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementAutocomplete(
		this.ref,
		val.Ref(),
	)
}

// Cols returns the value of property "HTMLTextAreaElement.cols".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Cols() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementCols(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCols sets the value of property "HTMLTextAreaElement.cols" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetCols(val uint32) bool {
	return js.True == bindings.SetHTMLTextAreaElementCols(
		this.ref,
		uint32(val),
	)
}

// DirName returns the value of property "HTMLTextAreaElement.dirName".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) DirName() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementDirName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDirName sets the value of property "HTMLTextAreaElement.dirName" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetDirName(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementDirName(
		this.ref,
		val.Ref(),
	)
}

// Disabled returns the value of property "HTMLTextAreaElement.disabled".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Disabled() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementDisabled(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDisabled sets the value of property "HTMLTextAreaElement.disabled" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetDisabled(val bool) bool {
	return js.True == bindings.SetHTMLTextAreaElementDisabled(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Form returns the value of property "HTMLTextAreaElement.form".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Form() (ret HTMLFormElement, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementForm(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MaxLength returns the value of property "HTMLTextAreaElement.maxLength".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) MaxLength() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementMaxLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMaxLength sets the value of property "HTMLTextAreaElement.maxLength" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetMaxLength(val int32) bool {
	return js.True == bindings.SetHTMLTextAreaElementMaxLength(
		this.ref,
		int32(val),
	)
}

// MinLength returns the value of property "HTMLTextAreaElement.minLength".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) MinLength() (ret int32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementMinLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetMinLength sets the value of property "HTMLTextAreaElement.minLength" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetMinLength(val int32) bool {
	return js.True == bindings.SetHTMLTextAreaElementMinLength(
		this.ref,
		int32(val),
	)
}

// Name returns the value of property "HTMLTextAreaElement.name".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "HTMLTextAreaElement.name" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetName(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementName(
		this.ref,
		val.Ref(),
	)
}

// Placeholder returns the value of property "HTMLTextAreaElement.placeholder".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Placeholder() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementPlaceholder(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPlaceholder sets the value of property "HTMLTextAreaElement.placeholder" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetPlaceholder(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementPlaceholder(
		this.ref,
		val.Ref(),
	)
}

// ReadOnly returns the value of property "HTMLTextAreaElement.readOnly".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) ReadOnly() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementReadOnly(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetReadOnly sets the value of property "HTMLTextAreaElement.readOnly" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetReadOnly(val bool) bool {
	return js.True == bindings.SetHTMLTextAreaElementReadOnly(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Required returns the value of property "HTMLTextAreaElement.required".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Required() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementRequired(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRequired sets the value of property "HTMLTextAreaElement.required" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetRequired(val bool) bool {
	return js.True == bindings.SetHTMLTextAreaElementRequired(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Rows returns the value of property "HTMLTextAreaElement.rows".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Rows() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementRows(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetRows sets the value of property "HTMLTextAreaElement.rows" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetRows(val uint32) bool {
	return js.True == bindings.SetHTMLTextAreaElementRows(
		this.ref,
		uint32(val),
	)
}

// Wrap returns the value of property "HTMLTextAreaElement.wrap".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Wrap() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementWrap(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetWrap sets the value of property "HTMLTextAreaElement.wrap" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetWrap(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementWrap(
		this.ref,
		val.Ref(),
	)
}

// Type returns the value of property "HTMLTextAreaElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// DefaultValue returns the value of property "HTMLTextAreaElement.defaultValue".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) DefaultValue() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementDefaultValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDefaultValue sets the value of property "HTMLTextAreaElement.defaultValue" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetDefaultValue(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementDefaultValue(
		this.ref,
		val.Ref(),
	)
}

// Value returns the value of property "HTMLTextAreaElement.value".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Value() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementValue(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetValue sets the value of property "HTMLTextAreaElement.value" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetValue(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementValue(
		this.ref,
		val.Ref(),
	)
}

// TextLength returns the value of property "HTMLTextAreaElement.textLength".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) TextLength() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementTextLength(
		this.ref, js.Pointer(&ret),
	)
	return
}

// WillValidate returns the value of property "HTMLTextAreaElement.willValidate".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) WillValidate() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementWillValidate(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Validity returns the value of property "HTMLTextAreaElement.validity".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Validity() (ret ValidityState, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementValidity(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ValidationMessage returns the value of property "HTMLTextAreaElement.validationMessage".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) ValidationMessage() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementValidationMessage(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Labels returns the value of property "HTMLTextAreaElement.labels".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) Labels() (ret NodeList, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementLabels(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SelectionStart returns the value of property "HTMLTextAreaElement.selectionStart".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) SelectionStart() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementSelectionStart(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionStart sets the value of property "HTMLTextAreaElement.selectionStart" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetSelectionStart(val uint32) bool {
	return js.True == bindings.SetHTMLTextAreaElementSelectionStart(
		this.ref,
		uint32(val),
	)
}

// SelectionEnd returns the value of property "HTMLTextAreaElement.selectionEnd".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) SelectionEnd() (ret uint32, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementSelectionEnd(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionEnd sets the value of property "HTMLTextAreaElement.selectionEnd" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetSelectionEnd(val uint32) bool {
	return js.True == bindings.SetHTMLTextAreaElementSelectionEnd(
		this.ref,
		uint32(val),
	)
}

// SelectionDirection returns the value of property "HTMLTextAreaElement.selectionDirection".
//
// It returns ok=false if there is no such property.
func (this HTMLTextAreaElement) SelectionDirection() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTextAreaElementSelectionDirection(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSelectionDirection sets the value of property "HTMLTextAreaElement.selectionDirection" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTextAreaElement) SetSelectionDirection(val js.String) bool {
	return js.True == bindings.SetHTMLTextAreaElementSelectionDirection(
		this.ref,
		val.Ref(),
	)
}

// HasFuncCheckValidity returns true if the method "HTMLTextAreaElement.checkValidity" exists.
func (this HTMLTextAreaElement) HasFuncCheckValidity() bool {
	return js.True == bindings.HasFuncHTMLTextAreaElementCheckValidity(
		this.ref,
	)
}

// FuncCheckValidity returns the method "HTMLTextAreaElement.checkValidity".
func (this HTMLTextAreaElement) FuncCheckValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLTextAreaElementCheckValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CheckValidity calls the method "HTMLTextAreaElement.checkValidity".
func (this HTMLTextAreaElement) CheckValidity() (ret bool) {
	bindings.CallHTMLTextAreaElementCheckValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCheckValidity calls the method "HTMLTextAreaElement.checkValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTextAreaElement) TryCheckValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementCheckValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncReportValidity returns true if the method "HTMLTextAreaElement.reportValidity" exists.
func (this HTMLTextAreaElement) HasFuncReportValidity() bool {
	return js.True == bindings.HasFuncHTMLTextAreaElementReportValidity(
		this.ref,
	)
}

// FuncReportValidity returns the method "HTMLTextAreaElement.reportValidity".
func (this HTMLTextAreaElement) FuncReportValidity() (fn js.Func[func() bool]) {
	bindings.FuncHTMLTextAreaElementReportValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ReportValidity calls the method "HTMLTextAreaElement.reportValidity".
func (this HTMLTextAreaElement) ReportValidity() (ret bool) {
	bindings.CallHTMLTextAreaElementReportValidity(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryReportValidity calls the method "HTMLTextAreaElement.reportValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTextAreaElement) TryReportValidity() (ret bool, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementReportValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetCustomValidity returns true if the method "HTMLTextAreaElement.setCustomValidity" exists.
func (this HTMLTextAreaElement) HasFuncSetCustomValidity() bool {
	return js.True == bindings.HasFuncHTMLTextAreaElementSetCustomValidity(
		this.ref,
	)
}

// FuncSetCustomValidity returns the method "HTMLTextAreaElement.setCustomValidity".
func (this HTMLTextAreaElement) FuncSetCustomValidity() (fn js.Func[func(err js.String)]) {
	bindings.FuncHTMLTextAreaElementSetCustomValidity(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetCustomValidity calls the method "HTMLTextAreaElement.setCustomValidity".
func (this HTMLTextAreaElement) SetCustomValidity(err js.String) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetCustomValidity(
		this.ref, js.Pointer(&ret),
		err.Ref(),
	)

	return
}

// TrySetCustomValidity calls the method "HTMLTextAreaElement.setCustomValidity"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTextAreaElement) TrySetCustomValidity(err js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetCustomValidity(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		err.Ref(),
	)

	return
}

// HasFuncSelect returns true if the method "HTMLTextAreaElement.select" exists.
func (this HTMLTextAreaElement) HasFuncSelect() bool {
	return js.True == bindings.HasFuncHTMLTextAreaElementSelect(
		this.ref,
	)
}

// FuncSelect returns the method "HTMLTextAreaElement.select".
func (this HTMLTextAreaElement) FuncSelect() (fn js.Func[func()]) {
	bindings.FuncHTMLTextAreaElementSelect(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Select calls the method "HTMLTextAreaElement.select".
func (this HTMLTextAreaElement) Select() (ret js.Void) {
	bindings.CallHTMLTextAreaElementSelect(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TrySelect calls the method "HTMLTextAreaElement.select"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTextAreaElement) TrySelect() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSelect(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncSetRangeText returns true if the method "HTMLTextAreaElement.setRangeText" exists.
func (this HTMLTextAreaElement) HasFuncSetRangeText() bool {
	return js.True == bindings.HasFuncHTMLTextAreaElementSetRangeText(
		this.ref,
	)
}

// FuncSetRangeText returns the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) FuncSetRangeText() (fn js.Func[func(replacement js.String)]) {
	bindings.FuncHTMLTextAreaElementSetRangeText(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetRangeText calls the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) SetRangeText(replacement js.String) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetRangeText(
		this.ref, js.Pointer(&ret),
		replacement.Ref(),
	)

	return
}

// TrySetRangeText calls the method "HTMLTextAreaElement.setRangeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTextAreaElement) TrySetRangeText(replacement js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetRangeText(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
	)

	return
}

// HasFuncSetRangeText1 returns true if the method "HTMLTextAreaElement.setRangeText" exists.
func (this HTMLTextAreaElement) HasFuncSetRangeText1() bool {
	return js.True == bindings.HasFuncHTMLTextAreaElementSetRangeText1(
		this.ref,
	)
}

// FuncSetRangeText1 returns the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) FuncSetRangeText1() (fn js.Func[func(replacement js.String, start uint32, end uint32, selectionMode SelectionMode)]) {
	bindings.FuncHTMLTextAreaElementSetRangeText1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetRangeText1 calls the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) SetRangeText1(replacement js.String, start uint32, end uint32, selectionMode SelectionMode) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetRangeText1(
		this.ref, js.Pointer(&ret),
		replacement.Ref(),
		uint32(start),
		uint32(end),
		uint32(selectionMode),
	)

	return
}

// TrySetRangeText1 calls the method "HTMLTextAreaElement.setRangeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTextAreaElement) TrySetRangeText1(replacement js.String, start uint32, end uint32, selectionMode SelectionMode) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetRangeText1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
		uint32(start),
		uint32(end),
		uint32(selectionMode),
	)

	return
}

// HasFuncSetRangeText2 returns true if the method "HTMLTextAreaElement.setRangeText" exists.
func (this HTMLTextAreaElement) HasFuncSetRangeText2() bool {
	return js.True == bindings.HasFuncHTMLTextAreaElementSetRangeText2(
		this.ref,
	)
}

// FuncSetRangeText2 returns the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) FuncSetRangeText2() (fn js.Func[func(replacement js.String, start uint32, end uint32)]) {
	bindings.FuncHTMLTextAreaElementSetRangeText2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetRangeText2 calls the method "HTMLTextAreaElement.setRangeText".
func (this HTMLTextAreaElement) SetRangeText2(replacement js.String, start uint32, end uint32) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetRangeText2(
		this.ref, js.Pointer(&ret),
		replacement.Ref(),
		uint32(start),
		uint32(end),
	)

	return
}

// TrySetRangeText2 calls the method "HTMLTextAreaElement.setRangeText"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTextAreaElement) TrySetRangeText2(replacement js.String, start uint32, end uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetRangeText2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		replacement.Ref(),
		uint32(start),
		uint32(end),
	)

	return
}

// HasFuncSetSelectionRange returns true if the method "HTMLTextAreaElement.setSelectionRange" exists.
func (this HTMLTextAreaElement) HasFuncSetSelectionRange() bool {
	return js.True == bindings.HasFuncHTMLTextAreaElementSetSelectionRange(
		this.ref,
	)
}

// FuncSetSelectionRange returns the method "HTMLTextAreaElement.setSelectionRange".
func (this HTMLTextAreaElement) FuncSetSelectionRange() (fn js.Func[func(start uint32, end uint32, direction js.String)]) {
	bindings.FuncHTMLTextAreaElementSetSelectionRange(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetSelectionRange calls the method "HTMLTextAreaElement.setSelectionRange".
func (this HTMLTextAreaElement) SetSelectionRange(start uint32, end uint32, direction js.String) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetSelectionRange(
		this.ref, js.Pointer(&ret),
		uint32(start),
		uint32(end),
		direction.Ref(),
	)

	return
}

// TrySetSelectionRange calls the method "HTMLTextAreaElement.setSelectionRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTextAreaElement) TrySetSelectionRange(start uint32, end uint32, direction js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetSelectionRange(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(start),
		uint32(end),
		direction.Ref(),
	)

	return
}

// HasFuncSetSelectionRange1 returns true if the method "HTMLTextAreaElement.setSelectionRange" exists.
func (this HTMLTextAreaElement) HasFuncSetSelectionRange1() bool {
	return js.True == bindings.HasFuncHTMLTextAreaElementSetSelectionRange1(
		this.ref,
	)
}

// FuncSetSelectionRange1 returns the method "HTMLTextAreaElement.setSelectionRange".
func (this HTMLTextAreaElement) FuncSetSelectionRange1() (fn js.Func[func(start uint32, end uint32)]) {
	bindings.FuncHTMLTextAreaElementSetSelectionRange1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// SetSelectionRange1 calls the method "HTMLTextAreaElement.setSelectionRange".
func (this HTMLTextAreaElement) SetSelectionRange1(start uint32, end uint32) (ret js.Void) {
	bindings.CallHTMLTextAreaElementSetSelectionRange1(
		this.ref, js.Pointer(&ret),
		uint32(start),
		uint32(end),
	)

	return
}

// TrySetSelectionRange1 calls the method "HTMLTextAreaElement.setSelectionRange"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this HTMLTextAreaElement) TrySetSelectionRange1(start uint32, end uint32) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryHTMLTextAreaElementSetSelectionRange1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		uint32(start),
		uint32(end),
	)

	return
}

type HTMLTimeElement struct {
	HTMLElement
}

func (this HTMLTimeElement) Once() HTMLTimeElement {
	this.ref.Once()
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
	this.ref.Free()
}

// DateTime returns the value of property "HTMLTimeElement.dateTime".
//
// It returns ok=false if there is no such property.
func (this HTMLTimeElement) DateTime() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTimeElementDateTime(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDateTime sets the value of property "HTMLTimeElement.dateTime" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTimeElement) SetDateTime(val js.String) bool {
	return js.True == bindings.SetHTMLTimeElementDateTime(
		this.ref,
		val.Ref(),
	)
}

type HTMLTitleElement struct {
	HTMLElement
}

func (this HTMLTitleElement) Once() HTMLTitleElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Text returns the value of property "HTMLTitleElement.text".
//
// It returns ok=false if there is no such property.
func (this HTMLTitleElement) Text() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTitleElementText(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetText sets the value of property "HTMLTitleElement.text" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTitleElement) SetText(val js.String) bool {
	return js.True == bindings.SetHTMLTitleElementText(
		this.ref,
		val.Ref(),
	)
}

const (
	HTMLTrackElement_NONE    uint16 = 0
	HTMLTrackElement_LOADING uint16 = 1
	HTMLTrackElement_LOADED  uint16 = 2
	HTMLTrackElement_ERROR   uint16 = 3
)

type HTMLTrackElement struct {
	HTMLElement
}

func (this HTMLTrackElement) Once() HTMLTrackElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Kind returns the value of property "HTMLTrackElement.kind".
//
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Kind() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementKind(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetKind sets the value of property "HTMLTrackElement.kind" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTrackElement) SetKind(val js.String) bool {
	return js.True == bindings.SetHTMLTrackElementKind(
		this.ref,
		val.Ref(),
	)
}

// Src returns the value of property "HTMLTrackElement.src".
//
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Src() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementSrc(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrc sets the value of property "HTMLTrackElement.src" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTrackElement) SetSrc(val js.String) bool {
	return js.True == bindings.SetHTMLTrackElementSrc(
		this.ref,
		val.Ref(),
	)
}

// Srclang returns the value of property "HTMLTrackElement.srclang".
//
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Srclang() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementSrclang(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetSrclang sets the value of property "HTMLTrackElement.srclang" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTrackElement) SetSrclang(val js.String) bool {
	return js.True == bindings.SetHTMLTrackElementSrclang(
		this.ref,
		val.Ref(),
	)
}

// Label returns the value of property "HTMLTrackElement.label".
//
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Label() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementLabel(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetLabel sets the value of property "HTMLTrackElement.label" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTrackElement) SetLabel(val js.String) bool {
	return js.True == bindings.SetHTMLTrackElementLabel(
		this.ref,
		val.Ref(),
	)
}

// Default returns the value of property "HTMLTrackElement.default".
//
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Default() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementDefault(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetDefault sets the value of property "HTMLTrackElement.default" to val.
//
// It returns false if the property cannot be set.
func (this HTMLTrackElement) SetDefault(val bool) bool {
	return js.True == bindings.SetHTMLTrackElementDefault(
		this.ref,
		js.Bool(bool(val)),
	)
}

// ReadyState returns the value of property "HTMLTrackElement.readyState".
//
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) ReadyState() (ret uint16, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementReadyState(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Track returns the value of property "HTMLTrackElement.track".
//
// It returns ok=false if there is no such property.
func (this HTMLTrackElement) Track() (ret TextTrack, ok bool) {
	ok = js.True == bindings.GetHTMLTrackElementTrack(
		this.ref, js.Pointer(&ret),
	)
	return
}

type HTMLUListElement struct {
	HTMLElement
}

func (this HTMLUListElement) Once() HTMLUListElement {
	this.ref.Once()
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
	this.ref.Free()
}

// Compact returns the value of property "HTMLUListElement.compact".
//
// It returns ok=false if there is no such property.
func (this HTMLUListElement) Compact() (ret bool, ok bool) {
	ok = js.True == bindings.GetHTMLUListElementCompact(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetCompact sets the value of property "HTMLUListElement.compact" to val.
//
// It returns false if the property cannot be set.
func (this HTMLUListElement) SetCompact(val bool) bool {
	return js.True == bindings.SetHTMLUListElementCompact(
		this.ref,
		js.Bool(bool(val)),
	)
}

// Type returns the value of property "HTMLUListElement.type".
//
// It returns ok=false if there is no such property.
func (this HTMLUListElement) Type() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHTMLUListElementType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "HTMLUListElement.type" to val.
//
// It returns false if the property cannot be set.
func (this HTMLUListElement) SetType(val js.String) bool {
	return js.True == bindings.SetHTMLUListElementType(
		this.ref,
		val.Ref(),
	)
}

type HTMLUnknownElement struct {
	HTMLElement
}

func (this HTMLUnknownElement) Once() HTMLUnknownElement {
	this.ref.Once()
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
	this.ref.Free()
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
func (p *HashChangeEventInit) UpdateFrom(ref js.Ref) {
	bindings.HashChangeEventInitJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HashChangeEventInit) Update(ref js.Ref) {
	bindings.HashChangeEventInitJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HashChangeEventInit) FreeMembers(recursive bool) {
	js.Free(
		p.OldURL.Ref(),
		p.NewURL.Ref(),
	)
	p.OldURL = p.OldURL.FromRef(js.Undefined)
	p.NewURL = p.NewURL.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// OldURL returns the value of property "HashChangeEvent.oldURL".
//
// It returns ok=false if there is no such property.
func (this HashChangeEvent) OldURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHashChangeEventOldURL(
		this.ref, js.Pointer(&ret),
	)
	return
}

// NewURL returns the value of property "HashChangeEvent.newURL".
//
// It returns ok=false if there is no such property.
func (this HashChangeEvent) NewURL() (ret js.String, ok bool) {
	ok = js.True == bindings.GetHashChangeEventNewURL(
		this.ref, js.Pointer(&ret),
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
func (p *HevcEncoderConfig) UpdateFrom(ref js.Ref) {
	bindings.HevcEncoderConfigJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HevcEncoderConfig) Update(ref js.Ref) {
	bindings.HevcEncoderConfigJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HevcEncoderConfig) FreeMembers(recursive bool) {
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
	this.ref.Once()
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
	this.ref.Free()
}

// Priority returns the value of property "Highlight.priority".
//
// It returns ok=false if there is no such property.
func (this Highlight) Priority() (ret int32, ok bool) {
	ok = js.True == bindings.GetHighlightPriority(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetPriority sets the value of property "Highlight.priority" to val.
//
// It returns false if the property cannot be set.
func (this Highlight) SetPriority(val int32) bool {
	return js.True == bindings.SetHighlightPriority(
		this.ref,
		int32(val),
	)
}

// Type returns the value of property "Highlight.type".
//
// It returns ok=false if there is no such property.
func (this Highlight) Type() (ret HighlightType, ok bool) {
	ok = js.True == bindings.GetHighlightType(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetType sets the value of property "Highlight.type" to val.
//
// It returns false if the property cannot be set.
func (this Highlight) SetType(val HighlightType) bool {
	return js.True == bindings.SetHighlightType(
		this.ref,
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
func (p *HkdfParams) UpdateFrom(ref js.Ref) {
	bindings.HkdfParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HkdfParams) Update(ref js.Ref) {
	bindings.HkdfParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HkdfParams) FreeMembers(recursive bool) {
	js.Free(
		p.Hash.Ref(),
		p.Salt.Ref(),
		p.Info.Ref(),
		p.Name.Ref(),
	)
	p.Hash = p.Hash.FromRef(js.Undefined)
	p.Salt = p.Salt.FromRef(js.Undefined)
	p.Info = p.Info.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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
func (p *HmacImportParams) UpdateFrom(ref js.Ref) {
	bindings.HmacImportParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HmacImportParams) Update(ref js.Ref) {
	bindings.HmacImportParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HmacImportParams) FreeMembers(recursive bool) {
	js.Free(
		p.Hash.Ref(),
		p.Name.Ref(),
	)
	p.Hash = p.Hash.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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
func (p *KeyAlgorithm) UpdateFrom(ref js.Ref) {
	bindings.KeyAlgorithmJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *KeyAlgorithm) Update(ref js.Ref) {
	bindings.KeyAlgorithmJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *KeyAlgorithm) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
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
func (p *HmacKeyAlgorithm) UpdateFrom(ref js.Ref) {
	bindings.HmacKeyAlgorithmJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HmacKeyAlgorithm) Update(ref js.Ref) {
	bindings.HmacKeyAlgorithmJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HmacKeyAlgorithm) FreeMembers(recursive bool) {
	js.Free(
		p.Name.Ref(),
	)
	p.Name = p.Name.FromRef(js.Undefined)
	if recursive {
		p.Hash.FreeMembers(true)
	}
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
func (p *HmacKeyGenParams) UpdateFrom(ref js.Ref) {
	bindings.HmacKeyGenParamsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *HmacKeyGenParams) Update(ref js.Ref) {
	bindings.HmacKeyGenParamsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *HmacKeyGenParams) FreeMembers(recursive bool) {
	js.Free(
		p.Hash.Ref(),
		p.Name.Ref(),
	)
	p.Hash = p.Hash.FromRef(js.Undefined)
	p.Name = p.Name.FromRef(js.Undefined)
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
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "IDBIndex.name".
//
// It returns ok=false if there is no such property.
func (this IDBIndex) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetIDBIndexName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "IDBIndex.name" to val.
//
// It returns false if the property cannot be set.
func (this IDBIndex) SetName(val js.String) bool {
	return js.True == bindings.SetIDBIndexName(
		this.ref,
		val.Ref(),
	)
}

// ObjectStore returns the value of property "IDBIndex.objectStore".
//
// It returns ok=false if there is no such property.
func (this IDBIndex) ObjectStore() (ret IDBObjectStore, ok bool) {
	ok = js.True == bindings.GetIDBIndexObjectStore(
		this.ref, js.Pointer(&ret),
	)
	return
}

// KeyPath returns the value of property "IDBIndex.keyPath".
//
// It returns ok=false if there is no such property.
func (this IDBIndex) KeyPath() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBIndexKeyPath(
		this.ref, js.Pointer(&ret),
	)
	return
}

// MultiEntry returns the value of property "IDBIndex.multiEntry".
//
// It returns ok=false if there is no such property.
func (this IDBIndex) MultiEntry() (ret bool, ok bool) {
	ok = js.True == bindings.GetIDBIndexMultiEntry(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Unique returns the value of property "IDBIndex.unique".
//
// It returns ok=false if there is no such property.
func (this IDBIndex) Unique() (ret bool, ok bool) {
	ok = js.True == bindings.GetIDBIndexUnique(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncGet returns true if the method "IDBIndex.get" exists.
func (this IDBIndex) HasFuncGet() bool {
	return js.True == bindings.HasFuncIDBIndexGet(
		this.ref,
	)
}

// FuncGet returns the method "IDBIndex.get".
func (this IDBIndex) FuncGet() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBIndexGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "IDBIndex.get".
func (this IDBIndex) Get(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexGet(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGet calls the method "IDBIndex.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryGet(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncGetKey returns true if the method "IDBIndex.getKey" exists.
func (this IDBIndex) HasFuncGetKey() bool {
	return js.True == bindings.HasFuncIDBIndexGetKey(
		this.ref,
	)
}

// FuncGetKey returns the method "IDBIndex.getKey".
func (this IDBIndex) FuncGetKey() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBIndexGetKey(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetKey calls the method "IDBIndex.getKey".
func (this IDBIndex) GetKey(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexGetKey(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetKey calls the method "IDBIndex.getKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryGetKey(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetKey(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncGetAll returns true if the method "IDBIndex.getAll" exists.
func (this IDBIndex) HasFuncGetAll() bool {
	return js.True == bindings.HasFuncIDBIndexGetAll(
		this.ref,
	)
}

// FuncGetAll returns the method "IDBIndex.getAll".
func (this IDBIndex) FuncGetAll() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	bindings.FuncIDBIndexGetAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAll calls the method "IDBIndex.getAll".
func (this IDBIndex) GetAll(query js.Any, count uint32) (ret IDBRequest) {
	bindings.CallIDBIndexGetAll(
		this.ref, js.Pointer(&ret),
		query.Ref(),
		uint32(count),
	)

	return
}

// TryGetAll calls the method "IDBIndex.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryGetAll(query js.Any, count uint32) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(count),
	)

	return
}

// HasFuncGetAll1 returns true if the method "IDBIndex.getAll" exists.
func (this IDBIndex) HasFuncGetAll1() bool {
	return js.True == bindings.HasFuncIDBIndexGetAll1(
		this.ref,
	)
}

// FuncGetAll1 returns the method "IDBIndex.getAll".
func (this IDBIndex) FuncGetAll1() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBIndexGetAll1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAll1 calls the method "IDBIndex.getAll".
func (this IDBIndex) GetAll1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexGetAll1(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetAll1 calls the method "IDBIndex.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryGetAll1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAll1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncGetAll2 returns true if the method "IDBIndex.getAll" exists.
func (this IDBIndex) HasFuncGetAll2() bool {
	return js.True == bindings.HasFuncIDBIndexGetAll2(
		this.ref,
	)
}

// FuncGetAll2 returns the method "IDBIndex.getAll".
func (this IDBIndex) FuncGetAll2() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBIndexGetAll2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAll2 calls the method "IDBIndex.getAll".
func (this IDBIndex) GetAll2() (ret IDBRequest) {
	bindings.CallIDBIndexGetAll2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAll2 calls the method "IDBIndex.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryGetAll2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAll2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAllKeys returns true if the method "IDBIndex.getAllKeys" exists.
func (this IDBIndex) HasFuncGetAllKeys() bool {
	return js.True == bindings.HasFuncIDBIndexGetAllKeys(
		this.ref,
	)
}

// FuncGetAllKeys returns the method "IDBIndex.getAllKeys".
func (this IDBIndex) FuncGetAllKeys() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	bindings.FuncIDBIndexGetAllKeys(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAllKeys calls the method "IDBIndex.getAllKeys".
func (this IDBIndex) GetAllKeys(query js.Any, count uint32) (ret IDBRequest) {
	bindings.CallIDBIndexGetAllKeys(
		this.ref, js.Pointer(&ret),
		query.Ref(),
		uint32(count),
	)

	return
}

// TryGetAllKeys calls the method "IDBIndex.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryGetAllKeys(query js.Any, count uint32) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAllKeys(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(count),
	)

	return
}

// HasFuncGetAllKeys1 returns true if the method "IDBIndex.getAllKeys" exists.
func (this IDBIndex) HasFuncGetAllKeys1() bool {
	return js.True == bindings.HasFuncIDBIndexGetAllKeys1(
		this.ref,
	)
}

// FuncGetAllKeys1 returns the method "IDBIndex.getAllKeys".
func (this IDBIndex) FuncGetAllKeys1() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBIndexGetAllKeys1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAllKeys1 calls the method "IDBIndex.getAllKeys".
func (this IDBIndex) GetAllKeys1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexGetAllKeys1(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetAllKeys1 calls the method "IDBIndex.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryGetAllKeys1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAllKeys1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncGetAllKeys2 returns true if the method "IDBIndex.getAllKeys" exists.
func (this IDBIndex) HasFuncGetAllKeys2() bool {
	return js.True == bindings.HasFuncIDBIndexGetAllKeys2(
		this.ref,
	)
}

// FuncGetAllKeys2 returns the method "IDBIndex.getAllKeys".
func (this IDBIndex) FuncGetAllKeys2() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBIndexGetAllKeys2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAllKeys2 calls the method "IDBIndex.getAllKeys".
func (this IDBIndex) GetAllKeys2() (ret IDBRequest) {
	bindings.CallIDBIndexGetAllKeys2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAllKeys2 calls the method "IDBIndex.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryGetAllKeys2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexGetAllKeys2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCount returns true if the method "IDBIndex.count" exists.
func (this IDBIndex) HasFuncCount() bool {
	return js.True == bindings.HasFuncIDBIndexCount(
		this.ref,
	)
}

// FuncCount returns the method "IDBIndex.count".
func (this IDBIndex) FuncCount() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBIndexCount(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Count calls the method "IDBIndex.count".
func (this IDBIndex) Count(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexCount(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryCount calls the method "IDBIndex.count"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryCount(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexCount(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncCount1 returns true if the method "IDBIndex.count" exists.
func (this IDBIndex) HasFuncCount1() bool {
	return js.True == bindings.HasFuncIDBIndexCount1(
		this.ref,
	)
}

// FuncCount1 returns the method "IDBIndex.count".
func (this IDBIndex) FuncCount1() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBIndexCount1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Count1 calls the method "IDBIndex.count".
func (this IDBIndex) Count1() (ret IDBRequest) {
	bindings.CallIDBIndexCount1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCount1 calls the method "IDBIndex.count"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryCount1() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexCount1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOpenCursor returns true if the method "IDBIndex.openCursor" exists.
func (this IDBIndex) HasFuncOpenCursor() bool {
	return js.True == bindings.HasFuncIDBIndexOpenCursor(
		this.ref,
	)
}

// FuncOpenCursor returns the method "IDBIndex.openCursor".
func (this IDBIndex) FuncOpenCursor() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	bindings.FuncIDBIndexOpenCursor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenCursor calls the method "IDBIndex.openCursor".
func (this IDBIndex) OpenCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest) {
	bindings.CallIDBIndexOpenCursor(
		this.ref, js.Pointer(&ret),
		query.Ref(),
		uint32(direction),
	)

	return
}

// TryOpenCursor calls the method "IDBIndex.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryOpenCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenCursor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(direction),
	)

	return
}

// HasFuncOpenCursor1 returns true if the method "IDBIndex.openCursor" exists.
func (this IDBIndex) HasFuncOpenCursor1() bool {
	return js.True == bindings.HasFuncIDBIndexOpenCursor1(
		this.ref,
	)
}

// FuncOpenCursor1 returns the method "IDBIndex.openCursor".
func (this IDBIndex) FuncOpenCursor1() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBIndexOpenCursor1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenCursor1 calls the method "IDBIndex.openCursor".
func (this IDBIndex) OpenCursor1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexOpenCursor1(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryOpenCursor1 calls the method "IDBIndex.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryOpenCursor1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenCursor1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncOpenCursor2 returns true if the method "IDBIndex.openCursor" exists.
func (this IDBIndex) HasFuncOpenCursor2() bool {
	return js.True == bindings.HasFuncIDBIndexOpenCursor2(
		this.ref,
	)
}

// FuncOpenCursor2 returns the method "IDBIndex.openCursor".
func (this IDBIndex) FuncOpenCursor2() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBIndexOpenCursor2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenCursor2 calls the method "IDBIndex.openCursor".
func (this IDBIndex) OpenCursor2() (ret IDBRequest) {
	bindings.CallIDBIndexOpenCursor2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryOpenCursor2 calls the method "IDBIndex.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryOpenCursor2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenCursor2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOpenKeyCursor returns true if the method "IDBIndex.openKeyCursor" exists.
func (this IDBIndex) HasFuncOpenKeyCursor() bool {
	return js.True == bindings.HasFuncIDBIndexOpenKeyCursor(
		this.ref,
	)
}

// FuncOpenKeyCursor returns the method "IDBIndex.openKeyCursor".
func (this IDBIndex) FuncOpenKeyCursor() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	bindings.FuncIDBIndexOpenKeyCursor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenKeyCursor calls the method "IDBIndex.openKeyCursor".
func (this IDBIndex) OpenKeyCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest) {
	bindings.CallIDBIndexOpenKeyCursor(
		this.ref, js.Pointer(&ret),
		query.Ref(),
		uint32(direction),
	)

	return
}

// TryOpenKeyCursor calls the method "IDBIndex.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryOpenKeyCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenKeyCursor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(direction),
	)

	return
}

// HasFuncOpenKeyCursor1 returns true if the method "IDBIndex.openKeyCursor" exists.
func (this IDBIndex) HasFuncOpenKeyCursor1() bool {
	return js.True == bindings.HasFuncIDBIndexOpenKeyCursor1(
		this.ref,
	)
}

// FuncOpenKeyCursor1 returns the method "IDBIndex.openKeyCursor".
func (this IDBIndex) FuncOpenKeyCursor1() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBIndexOpenKeyCursor1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenKeyCursor1 calls the method "IDBIndex.openKeyCursor".
func (this IDBIndex) OpenKeyCursor1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBIndexOpenKeyCursor1(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryOpenKeyCursor1 calls the method "IDBIndex.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryOpenKeyCursor1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenKeyCursor1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncOpenKeyCursor2 returns true if the method "IDBIndex.openKeyCursor" exists.
func (this IDBIndex) HasFuncOpenKeyCursor2() bool {
	return js.True == bindings.HasFuncIDBIndexOpenKeyCursor2(
		this.ref,
	)
}

// FuncOpenKeyCursor2 returns the method "IDBIndex.openKeyCursor".
func (this IDBIndex) FuncOpenKeyCursor2() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBIndexOpenKeyCursor2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenKeyCursor2 calls the method "IDBIndex.openKeyCursor".
func (this IDBIndex) OpenKeyCursor2() (ret IDBRequest) {
	bindings.CallIDBIndexOpenKeyCursor2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryOpenKeyCursor2 calls the method "IDBIndex.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBIndex) TryOpenKeyCursor2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBIndexOpenKeyCursor2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
func (p *IDBIndexParameters) UpdateFrom(ref js.Ref) {
	bindings.IDBIndexParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IDBIndexParameters) Update(ref js.Ref) {
	bindings.IDBIndexParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IDBIndexParameters) FreeMembers(recursive bool) {
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
func (p *IDBTransactionOptions) UpdateFrom(ref js.Ref) {
	bindings.IDBTransactionOptionsJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IDBTransactionOptions) Update(ref js.Ref) {
	bindings.IDBTransactionOptionsJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IDBTransactionOptions) FreeMembers(recursive bool) {
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
func (p *IDBObjectStoreParameters) UpdateFrom(ref js.Ref) {
	bindings.IDBObjectStoreParametersJSStore(
		js.Pointer(p), ref,
	)
}

// Update writes all fields of the p to the heap object referenced by ref.
func (p *IDBObjectStoreParameters) Update(ref js.Ref) {
	bindings.IDBObjectStoreParametersJSLoad(
		js.Pointer(p), js.False, ref,
	)
}

// FreeMembers frees fields with heap reference, if recursive is true
// free all heap references reachable from p.
func (p *IDBObjectStoreParameters) FreeMembers(recursive bool) {
	js.Free(
		p.KeyPath.Ref(),
	)
	p.KeyPath = p.KeyPath.FromRef(js.Undefined)
}

type IDBDatabase struct {
	EventTarget
}

func (this IDBDatabase) Once() IDBDatabase {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "IDBDatabase.name".
//
// It returns ok=false if there is no such property.
func (this IDBDatabase) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetIDBDatabaseName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Version returns the value of property "IDBDatabase.version".
//
// It returns ok=false if there is no such property.
func (this IDBDatabase) Version() (ret uint64, ok bool) {
	ok = js.True == bindings.GetIDBDatabaseVersion(
		this.ref, js.Pointer(&ret),
	)
	return
}

// ObjectStoreNames returns the value of property "IDBDatabase.objectStoreNames".
//
// It returns ok=false if there is no such property.
func (this IDBDatabase) ObjectStoreNames() (ret DOMStringList, ok bool) {
	ok = js.True == bindings.GetIDBDatabaseObjectStoreNames(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncTransaction returns true if the method "IDBDatabase.transaction" exists.
func (this IDBDatabase) HasFuncTransaction() bool {
	return js.True == bindings.HasFuncIDBDatabaseTransaction(
		this.ref,
	)
}

// FuncTransaction returns the method "IDBDatabase.transaction".
func (this IDBDatabase) FuncTransaction() (fn js.Func[func(storeNames OneOf_String_ArrayString, mode IDBTransactionMode, options IDBTransactionOptions) IDBTransaction]) {
	bindings.FuncIDBDatabaseTransaction(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Transaction calls the method "IDBDatabase.transaction".
func (this IDBDatabase) Transaction(storeNames OneOf_String_ArrayString, mode IDBTransactionMode, options IDBTransactionOptions) (ret IDBTransaction) {
	bindings.CallIDBDatabaseTransaction(
		this.ref, js.Pointer(&ret),
		storeNames.Ref(),
		uint32(mode),
		js.Pointer(&options),
	)

	return
}

// TryTransaction calls the method "IDBDatabase.transaction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBDatabase) TryTransaction(storeNames OneOf_String_ArrayString, mode IDBTransactionMode, options IDBTransactionOptions) (ret IDBTransaction, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseTransaction(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		storeNames.Ref(),
		uint32(mode),
		js.Pointer(&options),
	)

	return
}

// HasFuncTransaction1 returns true if the method "IDBDatabase.transaction" exists.
func (this IDBDatabase) HasFuncTransaction1() bool {
	return js.True == bindings.HasFuncIDBDatabaseTransaction1(
		this.ref,
	)
}

// FuncTransaction1 returns the method "IDBDatabase.transaction".
func (this IDBDatabase) FuncTransaction1() (fn js.Func[func(storeNames OneOf_String_ArrayString, mode IDBTransactionMode) IDBTransaction]) {
	bindings.FuncIDBDatabaseTransaction1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Transaction1 calls the method "IDBDatabase.transaction".
func (this IDBDatabase) Transaction1(storeNames OneOf_String_ArrayString, mode IDBTransactionMode) (ret IDBTransaction) {
	bindings.CallIDBDatabaseTransaction1(
		this.ref, js.Pointer(&ret),
		storeNames.Ref(),
		uint32(mode),
	)

	return
}

// TryTransaction1 calls the method "IDBDatabase.transaction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBDatabase) TryTransaction1(storeNames OneOf_String_ArrayString, mode IDBTransactionMode) (ret IDBTransaction, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseTransaction1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		storeNames.Ref(),
		uint32(mode),
	)

	return
}

// HasFuncTransaction2 returns true if the method "IDBDatabase.transaction" exists.
func (this IDBDatabase) HasFuncTransaction2() bool {
	return js.True == bindings.HasFuncIDBDatabaseTransaction2(
		this.ref,
	)
}

// FuncTransaction2 returns the method "IDBDatabase.transaction".
func (this IDBDatabase) FuncTransaction2() (fn js.Func[func(storeNames OneOf_String_ArrayString) IDBTransaction]) {
	bindings.FuncIDBDatabaseTransaction2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Transaction2 calls the method "IDBDatabase.transaction".
func (this IDBDatabase) Transaction2(storeNames OneOf_String_ArrayString) (ret IDBTransaction) {
	bindings.CallIDBDatabaseTransaction2(
		this.ref, js.Pointer(&ret),
		storeNames.Ref(),
	)

	return
}

// TryTransaction2 calls the method "IDBDatabase.transaction"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBDatabase) TryTransaction2(storeNames OneOf_String_ArrayString) (ret IDBTransaction, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseTransaction2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		storeNames.Ref(),
	)

	return
}

// HasFuncClose returns true if the method "IDBDatabase.close" exists.
func (this IDBDatabase) HasFuncClose() bool {
	return js.True == bindings.HasFuncIDBDatabaseClose(
		this.ref,
	)
}

// FuncClose returns the method "IDBDatabase.close".
func (this IDBDatabase) FuncClose() (fn js.Func[func()]) {
	bindings.FuncIDBDatabaseClose(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Close calls the method "IDBDatabase.close".
func (this IDBDatabase) Close() (ret js.Void) {
	bindings.CallIDBDatabaseClose(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClose calls the method "IDBDatabase.close"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBDatabase) TryClose() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseClose(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCreateObjectStore returns true if the method "IDBDatabase.createObjectStore" exists.
func (this IDBDatabase) HasFuncCreateObjectStore() bool {
	return js.True == bindings.HasFuncIDBDatabaseCreateObjectStore(
		this.ref,
	)
}

// FuncCreateObjectStore returns the method "IDBDatabase.createObjectStore".
func (this IDBDatabase) FuncCreateObjectStore() (fn js.Func[func(name js.String, options IDBObjectStoreParameters) IDBObjectStore]) {
	bindings.FuncIDBDatabaseCreateObjectStore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateObjectStore calls the method "IDBDatabase.createObjectStore".
func (this IDBDatabase) CreateObjectStore(name js.String, options IDBObjectStoreParameters) (ret IDBObjectStore) {
	bindings.CallIDBDatabaseCreateObjectStore(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCreateObjectStore calls the method "IDBDatabase.createObjectStore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBDatabase) TryCreateObjectStore(name js.String, options IDBObjectStoreParameters) (ret IDBObjectStore, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseCreateObjectStore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreateObjectStore1 returns true if the method "IDBDatabase.createObjectStore" exists.
func (this IDBDatabase) HasFuncCreateObjectStore1() bool {
	return js.True == bindings.HasFuncIDBDatabaseCreateObjectStore1(
		this.ref,
	)
}

// FuncCreateObjectStore1 returns the method "IDBDatabase.createObjectStore".
func (this IDBDatabase) FuncCreateObjectStore1() (fn js.Func[func(name js.String) IDBObjectStore]) {
	bindings.FuncIDBDatabaseCreateObjectStore1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateObjectStore1 calls the method "IDBDatabase.createObjectStore".
func (this IDBDatabase) CreateObjectStore1(name js.String) (ret IDBObjectStore) {
	bindings.CallIDBDatabaseCreateObjectStore1(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryCreateObjectStore1 calls the method "IDBDatabase.createObjectStore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBDatabase) TryCreateObjectStore1(name js.String) (ret IDBObjectStore, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseCreateObjectStore1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncDeleteObjectStore returns true if the method "IDBDatabase.deleteObjectStore" exists.
func (this IDBDatabase) HasFuncDeleteObjectStore() bool {
	return js.True == bindings.HasFuncIDBDatabaseDeleteObjectStore(
		this.ref,
	)
}

// FuncDeleteObjectStore returns the method "IDBDatabase.deleteObjectStore".
func (this IDBDatabase) FuncDeleteObjectStore() (fn js.Func[func(name js.String)]) {
	bindings.FuncIDBDatabaseDeleteObjectStore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteObjectStore calls the method "IDBDatabase.deleteObjectStore".
func (this IDBDatabase) DeleteObjectStore(name js.String) (ret js.Void) {
	bindings.CallIDBDatabaseDeleteObjectStore(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDeleteObjectStore calls the method "IDBDatabase.deleteObjectStore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBDatabase) TryDeleteObjectStore(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBDatabaseDeleteObjectStore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

type IDBTransaction struct {
	EventTarget
}

func (this IDBTransaction) Once() IDBTransaction {
	this.ref.Once()
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
	this.ref.Free()
}

// ObjectStoreNames returns the value of property "IDBTransaction.objectStoreNames".
//
// It returns ok=false if there is no such property.
func (this IDBTransaction) ObjectStoreNames() (ret DOMStringList, ok bool) {
	ok = js.True == bindings.GetIDBTransactionObjectStoreNames(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Mode returns the value of property "IDBTransaction.mode".
//
// It returns ok=false if there is no such property.
func (this IDBTransaction) Mode() (ret IDBTransactionMode, ok bool) {
	ok = js.True == bindings.GetIDBTransactionMode(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Durability returns the value of property "IDBTransaction.durability".
//
// It returns ok=false if there is no such property.
func (this IDBTransaction) Durability() (ret IDBTransactionDurability, ok bool) {
	ok = js.True == bindings.GetIDBTransactionDurability(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Db returns the value of property "IDBTransaction.db".
//
// It returns ok=false if there is no such property.
func (this IDBTransaction) Db() (ret IDBDatabase, ok bool) {
	ok = js.True == bindings.GetIDBTransactionDb(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Error returns the value of property "IDBTransaction.error".
//
// It returns ok=false if there is no such property.
func (this IDBTransaction) Error() (ret DOMException, ok bool) {
	ok = js.True == bindings.GetIDBTransactionError(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncObjectStore returns true if the method "IDBTransaction.objectStore" exists.
func (this IDBTransaction) HasFuncObjectStore() bool {
	return js.True == bindings.HasFuncIDBTransactionObjectStore(
		this.ref,
	)
}

// FuncObjectStore returns the method "IDBTransaction.objectStore".
func (this IDBTransaction) FuncObjectStore() (fn js.Func[func(name js.String) IDBObjectStore]) {
	bindings.FuncIDBTransactionObjectStore(
		this.ref, js.Pointer(&fn),
	)
	return
}

// ObjectStore calls the method "IDBTransaction.objectStore".
func (this IDBTransaction) ObjectStore(name js.String) (ret IDBObjectStore) {
	bindings.CallIDBTransactionObjectStore(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryObjectStore calls the method "IDBTransaction.objectStore"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBTransaction) TryObjectStore(name js.String) (ret IDBObjectStore, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBTransactionObjectStore(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncCommit returns true if the method "IDBTransaction.commit" exists.
func (this IDBTransaction) HasFuncCommit() bool {
	return js.True == bindings.HasFuncIDBTransactionCommit(
		this.ref,
	)
}

// FuncCommit returns the method "IDBTransaction.commit".
func (this IDBTransaction) FuncCommit() (fn js.Func[func()]) {
	bindings.FuncIDBTransactionCommit(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Commit calls the method "IDBTransaction.commit".
func (this IDBTransaction) Commit() (ret js.Void) {
	bindings.CallIDBTransactionCommit(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCommit calls the method "IDBTransaction.commit"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBTransaction) TryCommit() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBTransactionCommit(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncAbort returns true if the method "IDBTransaction.abort" exists.
func (this IDBTransaction) HasFuncAbort() bool {
	return js.True == bindings.HasFuncIDBTransactionAbort(
		this.ref,
	)
}

// FuncAbort returns the method "IDBTransaction.abort".
func (this IDBTransaction) FuncAbort() (fn js.Func[func()]) {
	bindings.FuncIDBTransactionAbort(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Abort calls the method "IDBTransaction.abort".
func (this IDBTransaction) Abort() (ret js.Void) {
	bindings.CallIDBTransactionAbort(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryAbort calls the method "IDBTransaction.abort"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBTransaction) TryAbort() (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBTransactionAbort(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

type IDBObjectStore struct {
	ref js.Ref
}

func (this IDBObjectStore) Once() IDBObjectStore {
	this.ref.Once()
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
	this.ref.Free()
}

// Name returns the value of property "IDBObjectStore.name".
//
// It returns ok=false if there is no such property.
func (this IDBObjectStore) Name() (ret js.String, ok bool) {
	ok = js.True == bindings.GetIDBObjectStoreName(
		this.ref, js.Pointer(&ret),
	)
	return
}

// SetName sets the value of property "IDBObjectStore.name" to val.
//
// It returns false if the property cannot be set.
func (this IDBObjectStore) SetName(val js.String) bool {
	return js.True == bindings.SetIDBObjectStoreName(
		this.ref,
		val.Ref(),
	)
}

// KeyPath returns the value of property "IDBObjectStore.keyPath".
//
// It returns ok=false if there is no such property.
func (this IDBObjectStore) KeyPath() (ret js.Any, ok bool) {
	ok = js.True == bindings.GetIDBObjectStoreKeyPath(
		this.ref, js.Pointer(&ret),
	)
	return
}

// IndexNames returns the value of property "IDBObjectStore.indexNames".
//
// It returns ok=false if there is no such property.
func (this IDBObjectStore) IndexNames() (ret DOMStringList, ok bool) {
	ok = js.True == bindings.GetIDBObjectStoreIndexNames(
		this.ref, js.Pointer(&ret),
	)
	return
}

// Transaction returns the value of property "IDBObjectStore.transaction".
//
// It returns ok=false if there is no such property.
func (this IDBObjectStore) Transaction() (ret IDBTransaction, ok bool) {
	ok = js.True == bindings.GetIDBObjectStoreTransaction(
		this.ref, js.Pointer(&ret),
	)
	return
}

// AutoIncrement returns the value of property "IDBObjectStore.autoIncrement".
//
// It returns ok=false if there is no such property.
func (this IDBObjectStore) AutoIncrement() (ret bool, ok bool) {
	ok = js.True == bindings.GetIDBObjectStoreAutoIncrement(
		this.ref, js.Pointer(&ret),
	)
	return
}

// HasFuncPut returns true if the method "IDBObjectStore.put" exists.
func (this IDBObjectStore) HasFuncPut() bool {
	return js.True == bindings.HasFuncIDBObjectStorePut(
		this.ref,
	)
}

// FuncPut returns the method "IDBObjectStore.put".
func (this IDBObjectStore) FuncPut() (fn js.Func[func(value js.Any, key js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStorePut(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Put calls the method "IDBObjectStore.put".
func (this IDBObjectStore) Put(value js.Any, key js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStorePut(
		this.ref, js.Pointer(&ret),
		value.Ref(),
		key.Ref(),
	)

	return
}

// TryPut calls the method "IDBObjectStore.put"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryPut(value js.Any, key js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStorePut(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
		key.Ref(),
	)

	return
}

// HasFuncPut1 returns true if the method "IDBObjectStore.put" exists.
func (this IDBObjectStore) HasFuncPut1() bool {
	return js.True == bindings.HasFuncIDBObjectStorePut1(
		this.ref,
	)
}

// FuncPut1 returns the method "IDBObjectStore.put".
func (this IDBObjectStore) FuncPut1() (fn js.Func[func(value js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStorePut1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Put1 calls the method "IDBObjectStore.put".
func (this IDBObjectStore) Put1(value js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStorePut1(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryPut1 calls the method "IDBObjectStore.put"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryPut1(value js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStorePut1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncAdd returns true if the method "IDBObjectStore.add" exists.
func (this IDBObjectStore) HasFuncAdd() bool {
	return js.True == bindings.HasFuncIDBObjectStoreAdd(
		this.ref,
	)
}

// FuncAdd returns the method "IDBObjectStore.add".
func (this IDBObjectStore) FuncAdd() (fn js.Func[func(value js.Any, key js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStoreAdd(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add calls the method "IDBObjectStore.add".
func (this IDBObjectStore) Add(value js.Any, key js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreAdd(
		this.ref, js.Pointer(&ret),
		value.Ref(),
		key.Ref(),
	)

	return
}

// TryAdd calls the method "IDBObjectStore.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryAdd(value js.Any, key js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreAdd(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
		key.Ref(),
	)

	return
}

// HasFuncAdd1 returns true if the method "IDBObjectStore.add" exists.
func (this IDBObjectStore) HasFuncAdd1() bool {
	return js.True == bindings.HasFuncIDBObjectStoreAdd1(
		this.ref,
	)
}

// FuncAdd1 returns the method "IDBObjectStore.add".
func (this IDBObjectStore) FuncAdd1() (fn js.Func[func(value js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStoreAdd1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Add1 calls the method "IDBObjectStore.add".
func (this IDBObjectStore) Add1(value js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreAdd1(
		this.ref, js.Pointer(&ret),
		value.Ref(),
	)

	return
}

// TryAdd1 calls the method "IDBObjectStore.add"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryAdd1(value js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreAdd1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		value.Ref(),
	)

	return
}

// HasFuncDelete returns true if the method "IDBObjectStore.delete" exists.
func (this IDBObjectStore) HasFuncDelete() bool {
	return js.True == bindings.HasFuncIDBObjectStoreDelete(
		this.ref,
	)
}

// FuncDelete returns the method "IDBObjectStore.delete".
func (this IDBObjectStore) FuncDelete() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStoreDelete(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Delete calls the method "IDBObjectStore.delete".
func (this IDBObjectStore) Delete(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreDelete(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryDelete calls the method "IDBObjectStore.delete"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryDelete(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreDelete(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncClear returns true if the method "IDBObjectStore.clear" exists.
func (this IDBObjectStore) HasFuncClear() bool {
	return js.True == bindings.HasFuncIDBObjectStoreClear(
		this.ref,
	)
}

// FuncClear returns the method "IDBObjectStore.clear".
func (this IDBObjectStore) FuncClear() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBObjectStoreClear(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Clear calls the method "IDBObjectStore.clear".
func (this IDBObjectStore) Clear() (ret IDBRequest) {
	bindings.CallIDBObjectStoreClear(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryClear calls the method "IDBObjectStore.clear"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryClear() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreClear(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGet returns true if the method "IDBObjectStore.get" exists.
func (this IDBObjectStore) HasFuncGet() bool {
	return js.True == bindings.HasFuncIDBObjectStoreGet(
		this.ref,
	)
}

// FuncGet returns the method "IDBObjectStore.get".
func (this IDBObjectStore) FuncGet() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStoreGet(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Get calls the method "IDBObjectStore.get".
func (this IDBObjectStore) Get(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGet(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGet calls the method "IDBObjectStore.get"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryGet(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGet(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncGetKey returns true if the method "IDBObjectStore.getKey" exists.
func (this IDBObjectStore) HasFuncGetKey() bool {
	return js.True == bindings.HasFuncIDBObjectStoreGetKey(
		this.ref,
	)
}

// FuncGetKey returns the method "IDBObjectStore.getKey".
func (this IDBObjectStore) FuncGetKey() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStoreGetKey(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetKey calls the method "IDBObjectStore.getKey".
func (this IDBObjectStore) GetKey(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetKey(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetKey calls the method "IDBObjectStore.getKey"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryGetKey(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetKey(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncGetAll returns true if the method "IDBObjectStore.getAll" exists.
func (this IDBObjectStore) HasFuncGetAll() bool {
	return js.True == bindings.HasFuncIDBObjectStoreGetAll(
		this.ref,
	)
}

// FuncGetAll returns the method "IDBObjectStore.getAll".
func (this IDBObjectStore) FuncGetAll() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	bindings.FuncIDBObjectStoreGetAll(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAll calls the method "IDBObjectStore.getAll".
func (this IDBObjectStore) GetAll(query js.Any, count uint32) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAll(
		this.ref, js.Pointer(&ret),
		query.Ref(),
		uint32(count),
	)

	return
}

// TryGetAll calls the method "IDBObjectStore.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryGetAll(query js.Any, count uint32) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAll(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(count),
	)

	return
}

// HasFuncGetAll1 returns true if the method "IDBObjectStore.getAll" exists.
func (this IDBObjectStore) HasFuncGetAll1() bool {
	return js.True == bindings.HasFuncIDBObjectStoreGetAll1(
		this.ref,
	)
}

// FuncGetAll1 returns the method "IDBObjectStore.getAll".
func (this IDBObjectStore) FuncGetAll1() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStoreGetAll1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAll1 calls the method "IDBObjectStore.getAll".
func (this IDBObjectStore) GetAll1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAll1(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetAll1 calls the method "IDBObjectStore.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryGetAll1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAll1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncGetAll2 returns true if the method "IDBObjectStore.getAll" exists.
func (this IDBObjectStore) HasFuncGetAll2() bool {
	return js.True == bindings.HasFuncIDBObjectStoreGetAll2(
		this.ref,
	)
}

// FuncGetAll2 returns the method "IDBObjectStore.getAll".
func (this IDBObjectStore) FuncGetAll2() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBObjectStoreGetAll2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAll2 calls the method "IDBObjectStore.getAll".
func (this IDBObjectStore) GetAll2() (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAll2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAll2 calls the method "IDBObjectStore.getAll"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryGetAll2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAll2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncGetAllKeys returns true if the method "IDBObjectStore.getAllKeys" exists.
func (this IDBObjectStore) HasFuncGetAllKeys() bool {
	return js.True == bindings.HasFuncIDBObjectStoreGetAllKeys(
		this.ref,
	)
}

// FuncGetAllKeys returns the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) FuncGetAllKeys() (fn js.Func[func(query js.Any, count uint32) IDBRequest]) {
	bindings.FuncIDBObjectStoreGetAllKeys(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAllKeys calls the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) GetAllKeys(query js.Any, count uint32) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAllKeys(
		this.ref, js.Pointer(&ret),
		query.Ref(),
		uint32(count),
	)

	return
}

// TryGetAllKeys calls the method "IDBObjectStore.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryGetAllKeys(query js.Any, count uint32) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAllKeys(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(count),
	)

	return
}

// HasFuncGetAllKeys1 returns true if the method "IDBObjectStore.getAllKeys" exists.
func (this IDBObjectStore) HasFuncGetAllKeys1() bool {
	return js.True == bindings.HasFuncIDBObjectStoreGetAllKeys1(
		this.ref,
	)
}

// FuncGetAllKeys1 returns the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) FuncGetAllKeys1() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStoreGetAllKeys1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAllKeys1 calls the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) GetAllKeys1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAllKeys1(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryGetAllKeys1 calls the method "IDBObjectStore.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryGetAllKeys1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAllKeys1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncGetAllKeys2 returns true if the method "IDBObjectStore.getAllKeys" exists.
func (this IDBObjectStore) HasFuncGetAllKeys2() bool {
	return js.True == bindings.HasFuncIDBObjectStoreGetAllKeys2(
		this.ref,
	)
}

// FuncGetAllKeys2 returns the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) FuncGetAllKeys2() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBObjectStoreGetAllKeys2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// GetAllKeys2 calls the method "IDBObjectStore.getAllKeys".
func (this IDBObjectStore) GetAllKeys2() (ret IDBRequest) {
	bindings.CallIDBObjectStoreGetAllKeys2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryGetAllKeys2 calls the method "IDBObjectStore.getAllKeys"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryGetAllKeys2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreGetAllKeys2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncCount returns true if the method "IDBObjectStore.count" exists.
func (this IDBObjectStore) HasFuncCount() bool {
	return js.True == bindings.HasFuncIDBObjectStoreCount(
		this.ref,
	)
}

// FuncCount returns the method "IDBObjectStore.count".
func (this IDBObjectStore) FuncCount() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStoreCount(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Count calls the method "IDBObjectStore.count".
func (this IDBObjectStore) Count(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreCount(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryCount calls the method "IDBObjectStore.count"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryCount(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreCount(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncCount1 returns true if the method "IDBObjectStore.count" exists.
func (this IDBObjectStore) HasFuncCount1() bool {
	return js.True == bindings.HasFuncIDBObjectStoreCount1(
		this.ref,
	)
}

// FuncCount1 returns the method "IDBObjectStore.count".
func (this IDBObjectStore) FuncCount1() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBObjectStoreCount1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Count1 calls the method "IDBObjectStore.count".
func (this IDBObjectStore) Count1() (ret IDBRequest) {
	bindings.CallIDBObjectStoreCount1(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryCount1 calls the method "IDBObjectStore.count"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryCount1() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreCount1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOpenCursor returns true if the method "IDBObjectStore.openCursor" exists.
func (this IDBObjectStore) HasFuncOpenCursor() bool {
	return js.True == bindings.HasFuncIDBObjectStoreOpenCursor(
		this.ref,
	)
}

// FuncOpenCursor returns the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) FuncOpenCursor() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	bindings.FuncIDBObjectStoreOpenCursor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenCursor calls the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) OpenCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenCursor(
		this.ref, js.Pointer(&ret),
		query.Ref(),
		uint32(direction),
	)

	return
}

// TryOpenCursor calls the method "IDBObjectStore.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryOpenCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenCursor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(direction),
	)

	return
}

// HasFuncOpenCursor1 returns true if the method "IDBObjectStore.openCursor" exists.
func (this IDBObjectStore) HasFuncOpenCursor1() bool {
	return js.True == bindings.HasFuncIDBObjectStoreOpenCursor1(
		this.ref,
	)
}

// FuncOpenCursor1 returns the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) FuncOpenCursor1() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStoreOpenCursor1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenCursor1 calls the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) OpenCursor1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenCursor1(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryOpenCursor1 calls the method "IDBObjectStore.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryOpenCursor1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenCursor1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncOpenCursor2 returns true if the method "IDBObjectStore.openCursor" exists.
func (this IDBObjectStore) HasFuncOpenCursor2() bool {
	return js.True == bindings.HasFuncIDBObjectStoreOpenCursor2(
		this.ref,
	)
}

// FuncOpenCursor2 returns the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) FuncOpenCursor2() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBObjectStoreOpenCursor2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenCursor2 calls the method "IDBObjectStore.openCursor".
func (this IDBObjectStore) OpenCursor2() (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenCursor2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryOpenCursor2 calls the method "IDBObjectStore.openCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryOpenCursor2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenCursor2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncOpenKeyCursor returns true if the method "IDBObjectStore.openKeyCursor" exists.
func (this IDBObjectStore) HasFuncOpenKeyCursor() bool {
	return js.True == bindings.HasFuncIDBObjectStoreOpenKeyCursor(
		this.ref,
	)
}

// FuncOpenKeyCursor returns the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) FuncOpenKeyCursor() (fn js.Func[func(query js.Any, direction IDBCursorDirection) IDBRequest]) {
	bindings.FuncIDBObjectStoreOpenKeyCursor(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenKeyCursor calls the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) OpenKeyCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenKeyCursor(
		this.ref, js.Pointer(&ret),
		query.Ref(),
		uint32(direction),
	)

	return
}

// TryOpenKeyCursor calls the method "IDBObjectStore.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryOpenKeyCursor(query js.Any, direction IDBCursorDirection) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenKeyCursor(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
		uint32(direction),
	)

	return
}

// HasFuncOpenKeyCursor1 returns true if the method "IDBObjectStore.openKeyCursor" exists.
func (this IDBObjectStore) HasFuncOpenKeyCursor1() bool {
	return js.True == bindings.HasFuncIDBObjectStoreOpenKeyCursor1(
		this.ref,
	)
}

// FuncOpenKeyCursor1 returns the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) FuncOpenKeyCursor1() (fn js.Func[func(query js.Any) IDBRequest]) {
	bindings.FuncIDBObjectStoreOpenKeyCursor1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenKeyCursor1 calls the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) OpenKeyCursor1(query js.Any) (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenKeyCursor1(
		this.ref, js.Pointer(&ret),
		query.Ref(),
	)

	return
}

// TryOpenKeyCursor1 calls the method "IDBObjectStore.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryOpenKeyCursor1(query js.Any) (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenKeyCursor1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		query.Ref(),
	)

	return
}

// HasFuncOpenKeyCursor2 returns true if the method "IDBObjectStore.openKeyCursor" exists.
func (this IDBObjectStore) HasFuncOpenKeyCursor2() bool {
	return js.True == bindings.HasFuncIDBObjectStoreOpenKeyCursor2(
		this.ref,
	)
}

// FuncOpenKeyCursor2 returns the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) FuncOpenKeyCursor2() (fn js.Func[func() IDBRequest]) {
	bindings.FuncIDBObjectStoreOpenKeyCursor2(
		this.ref, js.Pointer(&fn),
	)
	return
}

// OpenKeyCursor2 calls the method "IDBObjectStore.openKeyCursor".
func (this IDBObjectStore) OpenKeyCursor2() (ret IDBRequest) {
	bindings.CallIDBObjectStoreOpenKeyCursor2(
		this.ref, js.Pointer(&ret),
	)

	return
}

// TryOpenKeyCursor2 calls the method "IDBObjectStore.openKeyCursor"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryOpenKeyCursor2() (ret IDBRequest, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreOpenKeyCursor2(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
	)

	return
}

// HasFuncIndex returns true if the method "IDBObjectStore.index" exists.
func (this IDBObjectStore) HasFuncIndex() bool {
	return js.True == bindings.HasFuncIDBObjectStoreIndex(
		this.ref,
	)
}

// FuncIndex returns the method "IDBObjectStore.index".
func (this IDBObjectStore) FuncIndex() (fn js.Func[func(name js.String) IDBIndex]) {
	bindings.FuncIDBObjectStoreIndex(
		this.ref, js.Pointer(&fn),
	)
	return
}

// Index calls the method "IDBObjectStore.index".
func (this IDBObjectStore) Index(name js.String) (ret IDBIndex) {
	bindings.CallIDBObjectStoreIndex(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryIndex calls the method "IDBObjectStore.index"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryIndex(name js.String) (ret IDBIndex, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreIndex(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
	)

	return
}

// HasFuncCreateIndex returns true if the method "IDBObjectStore.createIndex" exists.
func (this IDBObjectStore) HasFuncCreateIndex() bool {
	return js.True == bindings.HasFuncIDBObjectStoreCreateIndex(
		this.ref,
	)
}

// FuncCreateIndex returns the method "IDBObjectStore.createIndex".
func (this IDBObjectStore) FuncCreateIndex() (fn js.Func[func(name js.String, keyPath OneOf_String_ArrayString, options IDBIndexParameters) IDBIndex]) {
	bindings.FuncIDBObjectStoreCreateIndex(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateIndex calls the method "IDBObjectStore.createIndex".
func (this IDBObjectStore) CreateIndex(name js.String, keyPath OneOf_String_ArrayString, options IDBIndexParameters) (ret IDBIndex) {
	bindings.CallIDBObjectStoreCreateIndex(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		keyPath.Ref(),
		js.Pointer(&options),
	)

	return
}

// TryCreateIndex calls the method "IDBObjectStore.createIndex"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryCreateIndex(name js.String, keyPath OneOf_String_ArrayString, options IDBIndexParameters) (ret IDBIndex, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreCreateIndex(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		keyPath.Ref(),
		js.Pointer(&options),
	)

	return
}

// HasFuncCreateIndex1 returns true if the method "IDBObjectStore.createIndex" exists.
func (this IDBObjectStore) HasFuncCreateIndex1() bool {
	return js.True == bindings.HasFuncIDBObjectStoreCreateIndex1(
		this.ref,
	)
}

// FuncCreateIndex1 returns the method "IDBObjectStore.createIndex".
func (this IDBObjectStore) FuncCreateIndex1() (fn js.Func[func(name js.String, keyPath OneOf_String_ArrayString) IDBIndex]) {
	bindings.FuncIDBObjectStoreCreateIndex1(
		this.ref, js.Pointer(&fn),
	)
	return
}

// CreateIndex1 calls the method "IDBObjectStore.createIndex".
func (this IDBObjectStore) CreateIndex1(name js.String, keyPath OneOf_String_ArrayString) (ret IDBIndex) {
	bindings.CallIDBObjectStoreCreateIndex1(
		this.ref, js.Pointer(&ret),
		name.Ref(),
		keyPath.Ref(),
	)

	return
}

// TryCreateIndex1 calls the method "IDBObjectStore.createIndex"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryCreateIndex1(name js.String, keyPath OneOf_String_ArrayString) (ret IDBIndex, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreCreateIndex1(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
		name.Ref(),
		keyPath.Ref(),
	)

	return
}

// HasFuncDeleteIndex returns true if the method "IDBObjectStore.deleteIndex" exists.
func (this IDBObjectStore) HasFuncDeleteIndex() bool {
	return js.True == bindings.HasFuncIDBObjectStoreDeleteIndex(
		this.ref,
	)
}

// FuncDeleteIndex returns the method "IDBObjectStore.deleteIndex".
func (this IDBObjectStore) FuncDeleteIndex() (fn js.Func[func(name js.String)]) {
	bindings.FuncIDBObjectStoreDeleteIndex(
		this.ref, js.Pointer(&fn),
	)
	return
}

// DeleteIndex calls the method "IDBObjectStore.deleteIndex".
func (this IDBObjectStore) DeleteIndex(name js.String) (ret js.Void) {
	bindings.CallIDBObjectStoreDeleteIndex(
		this.ref, js.Pointer(&ret),
		name.Ref(),
	)

	return
}

// TryDeleteIndex calls the method "IDBObjectStore.deleteIndex"
// in a try/catch block and returns (_, err, ok = false) when it went through
// the catch clause.
func (this IDBObjectStore) TryDeleteIndex(name js.String) (ret js.Void, exception js.Any, ok bool) {
	ok = js.True == bindings.TryIDBObjectStoreDeleteIndex(
		this.ref, js.Pointer(&ret), js.Pointer(&exception),
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
